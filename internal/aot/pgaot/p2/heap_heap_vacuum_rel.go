package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heap_vacuum_rel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v39 int64
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v327 int32
	_ = v327
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v379 int32
	_ = v379
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int64
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v426 float64
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 float64
	_ = v481
	var v484 int32
	_ = v484
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int64
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v904 int32
	_ = v904
	var v936 int32
	_ = v936
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v993 int64
	_ = v993
	var v994 int64
	_ = v994
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1128 int32
	_ = v1128
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1319 int64
	_ = v1319
	var v1321 int64
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1475 int64
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1514 int32
	_ = v1514
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int64
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1629 int32
	_ = v1629
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1764 int32
	_ = v1764
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
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
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1887 int32
	_ = v1887
	var v1892 int32
	_ = v1892
	var v1901 int32
	_ = v1901
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1984 int32
	_ = v1984
	var v1997 int32
	_ = v1997
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2025 int32
	_ = v2025
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2044 int32
	_ = v2044
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2097 int64
	_ = v2097
	var v2098 int64
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2103 int64
	_ = v2103
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2132 int32
	_ = v2132
	var v2147 int64
	_ = v2147
	var v2150 int64
	_ = v2150
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int64
	_ = v2180
	var v2181 int64
	_ = v2181
	var v2184 int32
	_ = v2184
	var v2185 int64
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2225 int32
	_ = v2225
	var v2228 int32
	_ = v2228
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2339 int64
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2355 int32
	_ = v2355
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2378 int64
	_ = v2378
	var v2383 int32
	_ = v2383
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2419 int64
	_ = v2419
	var v2422 int64
	_ = v2422
	var v2430 int64
	_ = v2430
	var v2433 int64
	_ = v2433
	var v2436 int64
	_ = v2436
	var v2442 int32
	_ = v2442
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2470 int32
	_ = v2470
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2547 int64
	_ = v2547
	var v2551 int32
	_ = v2551
	var v2552 int64
	_ = v2552
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2582 int32
	_ = v2582
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2595 int32
	_ = v2595
	var v2687 int32
	_ = v2687
	var v2690 int32
	_ = v2690
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2706 int64
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2711 int32
	_ = v2711
	var v2722 int32
	_ = v2722
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2749 int32
	_ = v2749
	var v2750 int64
	_ = v2750
	var v2751 int64
	_ = v2751
	var v2754 int64
	_ = v2754
	var v2758 int64
	_ = v2758
	var v2762 int64
	_ = v2762
	var v2763 int64
	_ = v2763
	var v2766 int64
	_ = v2766
	var v2767 int64
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2815 int32
	_ = v2815
	var v2818 int32
	_ = v2818
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2830 int32
	_ = v2830
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2858 int32
	_ = v2858
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2870 int32
	_ = v2870
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2888 int32
	_ = v2888
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2904 int32
	_ = v2904
	var v2909 int32
	_ = v2909
	var v2914 int32
	_ = v2914
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2921 int32
	_ = v2921
	var v2927 int32
	_ = v2927
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2933 int32
	_ = v2933
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2947 int32
	_ = v2947
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2966 int32
	_ = v2966
	var v2968 int32
	_ = v2968
	var v2971 int32
	_ = v2971
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2983 int64
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2992 int32
	_ = v2992
	var v2997 int32
	_ = v2997
	var v3003 int32
	_ = v3003
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3085 int32
	_ = v3085
	var v3090 int32
	_ = v3090
	var v3099 int32
	_ = v3099
	var v3108 int32
	_ = v3108
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3129 int32
	_ = v3129
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3134 int32
	_ = v3134
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3160 int32
	_ = v3160
	var v3164 int32
	_ = v3164
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3172 int32
	_ = v3172
	var v3175 int32
	_ = v3175
	var v3179 int32
	_ = v3179
	var v3181 int32
	_ = v3181
	var v3182 int32
	_ = v3182
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3197 int32
	_ = v3197
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3203 int64
	_ = v3203
	var v3204 float64
	_ = v3204
	var v3209 int32
	_ = v3209
	var v3210 float32
	_ = v3210
	var v3211 float64
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3225 float64
	_ = v3225
	var v3229 int32
	_ = v3229
	var v3249 float64
	_ = v3249
	var v3254 float64
	_ = v3254
	var v3256 float64
	_ = v3256
	var v3259 float64
	_ = v3259
	var v3260 int64
	_ = v3260
	var v3263 int64
	_ = v3263
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3270 int64
	_ = v3270
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3282 int32
	_ = v3282
	var v3286 int32
	_ = v3286
	var v3291 int32
	_ = v3291
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3301 int32
	_ = v3301
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3312 int32
	_ = v3312
	var v3313 int32
	_ = v3313
	var v3319 int32
	_ = v3319
	var v3323 int32
	_ = v3323
	var v3326 int32
	_ = v3326
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3331 float64
	_ = v3331
	var v3340 int64
	_ = v3340
	var v3358 int32
	_ = v3358
	var v3362 int32
	_ = v3362
	var v3367 int32
	_ = v3367
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3373 int32
	_ = v3373
	var v3377 int32
	_ = v3377
	var v3380 int32
	_ = v3380
	var v3472 int32
	_ = v3472
	var v3475 int32
	_ = v3475
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3491 int64
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3496 int32
	_ = v3496
	var v3507 int32
	_ = v3507
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3515 int32
	_ = v3515
	var v3517 int32
	_ = v3517
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3576 int64
	_ = v3576
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3595 int32
	_ = v3595
	var v3597 int32
	_ = v3597
	var v3602 int32
	_ = v3602
	var v3605 int32
	_ = v3605
	var v3607 int32
	_ = v3607
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3616 int32
	_ = v3616
	var v3619 int32
	_ = v3619
	var v3624 int32
	_ = v3624
	var v3625 int32
	_ = v3625
	var v3629 int32
	_ = v3629
	var v3631 int32
	_ = v3631
	var v3632 int32
	_ = v3632
	var v3634 int32
	_ = v3634
	var v3639 int64
	_ = v3639
	var v3642 int32
	_ = v3642
	var v3646 int32
	_ = v3646
	var v3651 int32
	_ = v3651
	var v3653 int32
	_ = v3653
	var v3654 int32
	_ = v3654
	var v3657 int32
	_ = v3657
	var v3661 int32
	_ = v3661
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3672 int32
	_ = v3672
	var v3673 int32
	_ = v3673
	var v3679 int32
	_ = v3679
	var v3683 int64
	_ = v3683
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3690 int32
	_ = v3690
	var v3695 int32
	_ = v3695
	var v3759 int32
	_ = v3759
	var v3763 int32
	_ = v3763
	var v3768 int32
	_ = v3768
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3774 int32
	_ = v3774
	var v3778 int32
	_ = v3778
	var v3781 int32
	_ = v3781
	var v3873 int32
	_ = v3873
	var v3876 int32
	_ = v3876
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3892 int64
	_ = v3892
	var v3894 int32
	_ = v3894
	var v3897 int32
	_ = v3897
	var v3908 int32
	_ = v3908
	var v3911 int32
	_ = v3911
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3916 int32
	_ = v3916
	var v3918 int32
	_ = v3918
	var v3980 int32
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3983 int32
	_ = v3983
	var v3994 int32
	_ = v3994
	var v4035 int32
	_ = v4035
	var v4038 int32
	_ = v4038
	var v4039 int32
	_ = v4039
	var v4046 int32
	_ = v4046
	var v4047 int32
	_ = v4047
	var v4049 int64
	_ = v4049
	var v4051 int64
	_ = v4051
	var v4053 int64
	_ = v4053
	var v4055 int64
	_ = v4055
	var v4057 int64
	_ = v4057
	var v4066 int32
	_ = v4066
	var v4067 int32
	_ = v4067
	var v4118 int32
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4123 int32
	_ = v4123
	var v4126 int32
	_ = v4126
	var v4127 int32
	_ = v4127
	var v4131 int32
	_ = v4131
	var v4133 int32
	_ = v4133
	var v4135 int32
	_ = v4135
	var v4187 int32
	_ = v4187
	var v4188 int32
	_ = v4188
	var v4191 int32
	_ = v4191
	var v4195 int32
	_ = v4195
	var v4199 int32
	_ = v4199
	var v4247 int32
	_ = v4247
	var v4249 int32
	_ = v4249
	var v4252 int32
	_ = v4252
	var v4254 int32
	_ = v4254
	var v4255 int32
	_ = v4255
	var v4256 float64
	_ = v4256
	var v4257 int32
	_ = v4257
	var v4266 int32
	_ = v4266
	var v4268 int32
	_ = v4268
	var v4270 int32
	_ = v4270
	var v4271 int32
	_ = v4271
	var v4278 int32
	_ = v4278
	var v4321 int32
	_ = v4321
	var v4324 int32
	_ = v4324
	var v4325 int32
	_ = v4325
	var v4329 int32
	_ = v4329
	var v4332 int32
	_ = v4332
	var v4333 int32
	_ = v4333
	var v4335 int32
	_ = v4335
	var v4346 int32
	_ = v4346
	var v4350 int32
	_ = v4350
	var v4355 int32
	_ = v4355
	var v4357 int32
	_ = v4357
	var v4358 int32
	_ = v4358
	var v4361 int32
	_ = v4361
	var v4365 int32
	_ = v4365
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4376 int32
	_ = v4376
	var v4377 int32
	_ = v4377
	var v4383 int32
	_ = v4383
	var v4389 int32
	_ = v4389
	var v4391 int32
	_ = v4391
	var v4400 int32
	_ = v4400
	var v4443 int32
	_ = v4443
	var v4444 int32
	_ = v4444
	var v4445 int32
	_ = v4445
	var v4450 int32
	_ = v4450
	var v4498 int32
	_ = v4498
	var v4500 int32
	_ = v4500
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4508 int32
	_ = v4508
	var v4509 int32
	_ = v4509
	var v4512 int32
	_ = v4512
	var v4518 int32
	_ = v4518
	var v4523 int32
	_ = v4523
	var v4525 int32
	_ = v4525
	var v4529 int32
	_ = v4529
	var v4530 int32
	_ = v4530
	var v4532 int32
	_ = v4532
	var v4533 int32
	_ = v4533
	var v4538 int32
	_ = v4538
	var v4541 int32
	_ = v4541
	var v4542 int32
	_ = v4542
	var v4543 int32
	_ = v4543
	var v4595 int32
	_ = v4595
	var v4597 int32
	_ = v4597
	var v4598 int32
	_ = v4598
	var v4600 int32
	_ = v4600
	var v4602 int32
	_ = v4602
	var v4607 int32
	_ = v4607
	var v4608 int32
	_ = v4608
	var v4609 int32
	_ = v4609
	var v4611 int64
	_ = v4611
	var v4612 int64
	_ = v4612
	var v4622 int32
	_ = v4622
	var v4630 int32
	_ = v4630
	var v4655 int64
	_ = v4655
	var v4672 int64
	_ = v4672
	var v4673 int64
	_ = v4673
	var v4676 int64
	_ = v4676
	var v4680 int32
	_ = v4680
	var v4681 int32
	_ = v4681
	var v4683 int32
	_ = v4683
	var v4685 int32
	_ = v4685
	var v4687 int32
	_ = v4687
	var v4691 int32
	_ = v4691
	var v4693 int32
	_ = v4693
	var v4695 int32
	_ = v4695
	var v4704 int32
	_ = v4704
	var v4705 int32
	_ = v4705
	var v4708 int64
	_ = v4708
	var v4710 int64
	_ = v4710
	var v4714 int32
	_ = v4714
	var v4716 int32
	_ = v4716
	var v4721 int32
	_ = v4721
	var v4722 int32
	_ = v4722
	var v4723 int64
	_ = v4723
	var v4728 int32
	_ = v4728
	var v4729 int32
	_ = v4729
	var v4732 int32
	_ = v4732
	var v4733 int32
	_ = v4733
	var v4739 int32
	_ = v4739
	var v4744 int32
	_ = v4744
	var v4746 int32
	_ = v4746
	var v4747 int32
	_ = v4747
	var v4754 int32
	_ = v4754
	var v4756 int32
	_ = v4756
	var v4757 int32
	_ = v4757
	var v4758 int32
	_ = v4758
	var v4759 int32
	_ = v4759
	var v4767 int32
	_ = v4767
	var v4770 int32
	_ = v4770
	var v4771 int32
	_ = v4771
	var v4772 int32
	_ = v4772
	var v4773 int32
	_ = v4773
	var v4779 int32
	_ = v4779
	var v4784 int32
	_ = v4784
	var v4786 int32
	_ = v4786
	var v4787 int32
	_ = v4787
	var v4788 int32
	_ = v4788
	var v4789 int32
	_ = v4789
	var v4790 int32
	_ = v4790
	var v4792 int32
	_ = v4792
	var v4796 int32
	_ = v4796
	var v4804 int32
	_ = v4804
	var v4808 int32
	_ = v4808
	var v4813 int32
	_ = v4813
	var v4817 int32
	_ = v4817
	var v4824 int32
	_ = v4824
	var v4829 int32
	_ = v4829
	var v4835 int32
	_ = v4835
	var v4838 int32
	_ = v4838
	var v4839 int32
	_ = v4839
	var v4841 int32
	_ = v4841
	var v4842 int32
	_ = v4842
	var v4845 int32
	_ = v4845
	var v4851 int32
	_ = v4851
	var v4856 int32
	_ = v4856
	var v4863 int64
	_ = v4863
	var v4866 int32
	_ = v4866
	var v4868 int32
	_ = v4868
	var v4870 int32
	_ = v4870
	var v4873 int32
	_ = v4873
	var v4876 int32
	_ = v4876
	var v4925 int32
	_ = v4925
	var v4928 int32
	_ = v4928
	var v4930 int32
	_ = v4930
	var v4932 int32
	_ = v4932
	var v4949 int32
	_ = v4949
	var v4985 int32
	_ = v4985
	var v4986 int32
	_ = v4986
	var v4988 int32
	_ = v4988
	var v4989 int32
	_ = v4989
	var v4990 int32
	_ = v4990
	var v4993 int32
	_ = v4993
	var v4997 int32
	_ = v4997
	var v5003 int32
	_ = v5003
	var v5005 int32
	_ = v5005
	var v5011 int32
	_ = v5011
	var v5012 int32
	_ = v5012
	var v5015 int32
	_ = v5015
	var v5023 int32
	_ = v5023
	var v5031 int32
	_ = v5031
	var v5083 int32
	_ = v5083
	var v5089 int32
	_ = v5089
	var v5094 int32
	_ = v5094
	var v5145 int32
	_ = v5145
	var v5146 int32
	_ = v5146
	var v5156 int32
	_ = v5156
	var v5169 int32
	_ = v5169
	var v5198 int32
	_ = v5198
	var v5201 int32
	_ = v5201
	var v5203 int32
	_ = v5203
	var v5204 int32
	_ = v5204
	var v5206 int32
	_ = v5206
	var v5208 int32
	_ = v5208
	var v5214 int32
	_ = v5214
	var v5215 int32
	_ = v5215
	var v5217 int32
	_ = v5217
	var v5218 int32
	_ = v5218
	var v5219 int32
	_ = v5219
	var v5227 int32
	_ = v5227
	var v5232 int32
	_ = v5232
	var v5234 int32
	_ = v5234
	var v5287 int32
	_ = v5287
	var v5293 int32
	_ = v5293
	var v5297 int32
	_ = v5297
	var v5302 int32
	_ = v5302
	var v5304 int32
	_ = v5304
	var v5305 int32
	_ = v5305
	var v5308 int32
	_ = v5308
	var v5312 int32
	_ = v5312
	var v5314 int32
	_ = v5314
	var v5315 int32
	_ = v5315
	var v5323 int32
	_ = v5323
	var v5324 int32
	_ = v5324
	var v5330 int32
	_ = v5330
	var v5334 int32
	_ = v5334
	var v5337 int32
	_ = v5337
	var v5341 int32
	_ = v5341
	var v5347 int32
	_ = v5347
	var v5348 int32
	_ = v5348
	var v5351 int32
	_ = v5351
	var v5352 int32
	_ = v5352
	var v5355 int32
	_ = v5355
	var v5356 int32
	_ = v5356
	var v5357 float64
	_ = v5357
	var v5358 int32
	_ = v5358
	var v5361 int32
	_ = v5361
	var v5362 int32
	_ = v5362
	var v5369 int32
	_ = v5369
	var v5370 int32
	_ = v5370
	var v5371 int32
	_ = v5371
	var v5372 int32
	_ = v5372
	var v5373 float64
	_ = v5373
	var v5374 float64
	_ = v5374
	var v5377 float64
	_ = v5377
	var v5379 int64
	_ = v5379
	var v5380 int64
	_ = v5380
	var v5383 int32
	_ = v5383
	var v5387 int32
	_ = v5387
	var v5391 int32
	_ = v5391
	var v5392 int32
	_ = v5392
	var v5393 int32
	_ = v5393
	var v5396 int64
	_ = v5396
	var v5397 int64
	_ = v5397
	var v5405 int64
	_ = v5405
	var v5411 int64
	_ = v5411
	var v5420 int64
	_ = v5420
	var v5423 int32
	_ = v5423
	var v5426 int32
	_ = v5426
	var v5429 int32
	_ = v5429
	var v5430 int32
	_ = v5430
	var v5431 int32
	_ = v5431
	var v5439 int32
	_ = v5439
	var v5441 int32
	_ = v5441
	var v5442 int32
	_ = v5442
	var v5447 int32
	_ = v5447
	var v5448 int32
	_ = v5448
	var v5449 int64
	_ = v5449
	var v5455 int32
	_ = v5455
	var v5456 int32
	_ = v5456
	var v5457 int64
	_ = v5457
	var v5462 int32
	_ = v5462
	var v5465 int32
	_ = v5465
	var v5468 int32
	_ = v5468
	var v5469 int32
	_ = v5469
	var v5478 int32
	_ = v5478
	var v5482 int32
	_ = v5482
	var v5487 int32
	_ = v5487
	var v5490 int32
	_ = v5490
	var v5492 int32
	_ = v5492
	var v5493 int32
	_ = v5493
	var v5496 int32
	_ = v5496
	var v5500 int32
	_ = v5500
	var v5502 int32
	_ = v5502
	var v5503 int32
	_ = v5503
	var v5511 int32
	_ = v5511
	var v5512 int32
	_ = v5512
	var v5518 int32
	_ = v5518
	var v5527 int32
	_ = v5527
	var v5528 int32
	_ = v5528
	var v5529 int32
	_ = v5529
	var v5532 int64
	_ = v5532
	var v5533 int64
	_ = v5533
	var v5541 int64
	_ = v5541
	var v5542 int32
	_ = v5542
	var v5559 int64
	_ = v5559
	var v5563 int64
	_ = v5563
	var v5564 int64
	_ = v5564
	var v5571 int32
	_ = v5571
	var v5572 int32
	_ = v5572
	var v5575 int64
	_ = v5575
	var v5584 int32
	_ = v5584
	var v5586 int32
	_ = v5586
	var v5587 int64
	_ = v5587
	var v5589 int64
	_ = v5589
	var v5590 int64
	_ = v5590
	var v5594 int64
	_ = v5594
	var v5596 int64
	_ = v5596
	var v5597 int64
	_ = v5597
	var v5601 int64
	_ = v5601
	var v5603 int64
	_ = v5603
	var v5604 int64
	_ = v5604
	var v5608 int64
	_ = v5608
	var v5610 int64
	_ = v5610
	var v5611 int64
	_ = v5611
	var v5616 int32
	_ = v5616
	var v5621 int32
	_ = v5621
	var v5622 int64
	_ = v5622
	var v5624 int64
	_ = v5624
	var v5625 int64
	_ = v5625
	var v5629 int64
	_ = v5629
	var v5631 int64
	_ = v5631
	var v5632 int64
	_ = v5632
	var v5636 int64
	_ = v5636
	var v5638 int64
	_ = v5638
	var v5639 int64
	_ = v5639
	var v5643 int64
	_ = v5643
	var v5645 int64
	_ = v5645
	var v5646 int64
	_ = v5646
	var v5650 int64
	_ = v5650
	var v5652 int64
	_ = v5652
	var v5653 int64
	_ = v5653
	var v5657 int64
	_ = v5657
	var v5659 int64
	_ = v5659
	var v5660 int64
	_ = v5660
	var v5664 int64
	_ = v5664
	var v5666 int64
	_ = v5666
	var v5667 int64
	_ = v5667
	var v5671 int64
	_ = v5671
	var v5673 int64
	_ = v5673
	var v5674 int64
	_ = v5674
	var v5678 int64
	_ = v5678
	var v5680 int64
	_ = v5680
	var v5681 int64
	_ = v5681
	var v5685 int64
	_ = v5685
	var v5687 int64
	_ = v5687
	var v5688 int64
	_ = v5688
	var v5692 int64
	_ = v5692
	var v5694 int64
	_ = v5694
	var v5695 int64
	_ = v5695
	var v5699 int64
	_ = v5699
	var v5701 int64
	_ = v5701
	var v5702 int64
	_ = v5702
	var v5706 int64
	_ = v5706
	var v5708 int64
	_ = v5708
	var v5709 int64
	_ = v5709
	var v5713 int64
	_ = v5713
	var v5715 int64
	_ = v5715
	var v5716 int64
	_ = v5716
	var v5720 int64
	_ = v5720
	var v5722 int64
	_ = v5722
	var v5723 int64
	_ = v5723
	var v5727 int64
	_ = v5727
	var v5729 int64
	_ = v5729
	var v5730 int64
	_ = v5730
	var v5734 int64
	_ = v5734
	var v5735 int64
	_ = v5735
	var v5736 int64
	_ = v5736
	var v5737 int64
	_ = v5737
	var v5738 int64
	_ = v5738
	var v5739 int64
	_ = v5739
	var v5743 int32
	_ = v5743
	var v5747 int32
	_ = v5747
	var v5750 int32
	_ = v5750
	var v5751 int32
	_ = v5751
	var v5758 int32
	_ = v5758
	var v5760 int32
	_ = v5760
	var v5761 int32
	_ = v5761
	var v5762 int64
	_ = v5762
	var v5763 int32
	_ = v5763
	var v5768 int32
	_ = v5768
	var v5772 int32
	_ = v5772
	var v5773 int32
	_ = v5773
	var v5774 int32
	_ = v5774
	var v5775 int32
	_ = v5775
	var v5778 float64
	_ = v5778
	var v5783 float64
	_ = v5783
	var v5792 int32
	_ = v5792
	var v5793 float64
	_ = v5793
	var v5794 int64
	_ = v5794
	var v5795 int64
	_ = v5795
	var v5804 int32
	_ = v5804
	var v5805 int64
	_ = v5805
	var v5808 int32
	_ = v5808
	var v5815 int32
	_ = v5815
	var v5816 int64
	_ = v5816
	var v5817 int32
	_ = v5817
	var v5818 int32
	_ = v5818
	var v5824 int32
	_ = v5824
	var v5829 int32
	_ = v5829
	var v5830 int32
	_ = v5830
	var v5833 int32
	_ = v5833
	var v5834 int32
	_ = v5834
	var v5842 int32
	_ = v5842
	var v5845 int32
	_ = v5845
	var v5848 int32
	_ = v5848
	var v5849 int32
	_ = v5849
	var v5859 int32
	_ = v5859
	var v5862 int32
	_ = v5862
	var v5863 int64
	_ = v5863
	var v5866 float64
	_ = v5866
	var v5871 float64
	_ = v5871
	var v5875 int32
	_ = v5875
	var v5880 int32
	_ = v5880
	var v5881 int32
	_ = v5881
	var v5882 int32
	_ = v5882
	var v5883 int32
	_ = v5883
	var v5892 int32
	_ = v5892
	var v5893 int32
	_ = v5893
	var v5896 int32
	_ = v5896
	var v5898 int32
	_ = v5898
	var v5903 int32
	_ = v5903
	var v5904 int32
	_ = v5904
	var v5909 int32
	_ = v5909
	var v5910 int32
	_ = v5910
	var v5911 int32
	_ = v5911
	var v5912 int32
	_ = v5912
	var v5914 int32
	_ = v5914
	var v5915 int32
	_ = v5915
	var v5916 int64
	_ = v5916
	var v5919 float64
	_ = v5919
	var v5924 float64
	_ = v5924
	var v5932 int32
	_ = v5932
	var v5933 int32
	_ = v5933
	var v5934 int32
	_ = v5934
	var v5944 int32
	_ = v5944
	var v5947 int32
	_ = v5947
	var v5989 int32
	_ = v5989
	var v5990 int32
	_ = v5990
	var v5992 int32
	_ = v5992
	var v5994 int32
	_ = v5994
	var v5995 int64
	_ = v5995
	var v5996 int32
	_ = v5996
	var v5997 int32
	_ = v5997
	var v6008 int32
	_ = v6008
	var v6009 int32
	_ = v6009
	var v6012 int32
	_ = v6012
	var v6015 int32
	_ = v6015
	var v6067 int32
	_ = v6067
	var v6069 int32
	_ = v6069
	var v6070 int64
	_ = v6070
	var v6081 int32
	_ = v6081
	var v6083 int32
	_ = v6083
	var v6087 int64
	_ = v6087
	var v6090 float64
	_ = v6090
	var v6094 int64
	_ = v6094
	var v6106 int32
	_ = v6106
	var v6107 int64
	_ = v6107
	var v6108 int64
	_ = v6108
	var v6110 int32
	_ = v6110
	var v6111 int32
	_ = v6111
	var v6118 float64
	_ = v6118
	var v6120 float64
	_ = v6120
	var v6126 float64
	_ = v6126
	var v6135 float64
	_ = v6135
	var v6136 float64
	_ = v6136
	var v6140 int32
	_ = v6140
	var v6145 int32
	_ = v6145
	var v6153 int32
	_ = v6153
	var v6154 int64
	_ = v6154
	var v6156 int64
	_ = v6156
	var v6158 int64
	_ = v6158
	var v6160 int64
	_ = v6160
	var v6166 int32
	_ = v6166
	var v6169 int32
	_ = v6169
	var v6170 int32
	_ = v6170
	var v6176 int32
	_ = v6176
	var v6179 int32
	_ = v6179
	var v6181 int32
	_ = v6181
	var v6182 int32
	_ = v6182
	var v6183 int32
	_ = v6183
	var v6187 int32
	_ = v6187
	var v6192 int32
	_ = v6192
	var v6193 int32
	_ = v6193
	var v6195 int32
	_ = v6195
	var v6245 int32
	_ = v6245
	var v6250 int32
	_ = v6250
	var v6298 int32
	_ = v6298
	var v6299 int32
	_ = v6299
	var v6301 int32
	_ = v6301
	var v6303 int32
	_ = v6303
	var v6305 int32
	_ = v6305
	var v6307 int32
	_ = v6307
	var v6309 int32
	_ = v6309
	var v6310 int32
	_ = v6310
	v4 = int32(0)
	v39 = int64(0)
	v50 = m.G0
	v52 = v50 - int32(1616)
	m.G0 = v52
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+728)) = v55
	v58 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+720)) = v58
	v61 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+712)) = v61
	v64 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+704)) = v64
	base.MemoryCopy(m, v52+int32(576), int32(_a_F_heap_vacuum_rel_0), int32(128))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = v71 & int32(4)
	v75 = int32(base.Ui32(v73) >> (uint(int32(2)) % 32))
	if v73 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v109 = m.G0
	v110 = int32(16)
	v111 = v109 - v110
	m.G0 = v111
	F_gettimeofday(m, v111)
	mBase = m.M
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
	v115 = int64(*(*int32)(unsafe.Add(mBase, uint32(v111)+8)))
	m.G0 = v111 + v110
	v123 = v115 + v114*int64(1000000) - int64(946684800000000)
	goto L9
L2:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[4]))
	if v80 != int32(4) {
		v103 = v4
		v104 = v39
		v105 = int64(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_getrusage(m, v52+int32(752))
	mBase = m.M
	F_gettimeofday(m, v52+int32(736))
	mBase = m.M
	goto L7
L5:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v84 < int32(0) {
		v103 = v4
		v104 = v39
		v105 = int64(0)
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v93 = int32(1)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5])))
	if v96 != v93 {
		v103 = v93
		v104 = v39
		v105 = int64(0)
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v100 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[6]))
	v102 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[7]))
	v103 = v93
	v104 = v100
	v105 = v102
	goto L1
L9:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v128 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v173 = F_palloc0(m, int32(256))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v132&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v137 = int32(_a_F_heap_vacuum_rel_1)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v140 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v139 + v140
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v143 + v140
	v147 = int32(0)
	v149 = int32(_a_F_heap_vacuum_rel_2)
	v150 = base.AtomicRmwOr32(m, v147, v149, v147)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+220)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+224)) = v125
	base.MemoryFill(m, v128+int32(232), v147, int32(160))
	v161 = base.AtomicRmwOr32(m, v147, v149, v147)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v162 + v140
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v168 - v140
	goto L11
L14:
	;
	return
L15:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v177 = F_get_database_name(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+68)) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+68))
	v182 = F_get_namespace_name(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+72)) = v182
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v188 = F_pstrdup(m, v185+int32(4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+96)) = uint8(v75)
	v191 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+92)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v173)+80)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v173)+76)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v52)+572)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v52)+568)) = int32(187)
	v199 = int32(_a_F_heap_vacuum_rel_3)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12])) = v52 + int32(564)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+564)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = l0
	v209 = v173 + int32(8)
	v211 = v173 + int32(4)
	F_vac_open_indexes(m, l0, int32(3), v209, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = l2
	if v103 == int32(0) {
		v327 = v4
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v343 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])) = uint8(v343)
	v345 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+24)) = uint8(v345)
	v347 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+22)) = uint16(v347)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v350 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+25)) = uint8(base.B2i32(v349 != v350))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	switch v353 - v350 {
	case 0:
		goto L31
	case 1:
		goto L30
	default:
		goto L29
	}
L21:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v217 <= int32(0) {
		v327 = v4
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v222 = F_palloc(m, v217<<(uint(int32(2))%32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v224 <= int32(0) {
		v327 = v222
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v230 = int32(0)
	goto L25
L25:
	;
	v278 = v230 << (uint(int32(2)) % 32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v280+v278)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+48))
	v286 = F_pstrdup(m, v283+int32(4))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L14
	} else {
		goto L27
	}
L26:
	;
	v327 = v222
	goto L20
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222+v278))) = v286
	v290 = v230 + int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v290 < v291 {
		v230 = v290
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v360 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v173)+120)) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v173)+112)) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v173)+140)) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v173)+148)) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v173)+156)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v173)+164)) = int32(0)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v375 = F_palloc0(m, v372<<(uint(int32(2))%32))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L14
	} else {
		goto L32
	}
L30:
	;
	v358 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+22)) = uint8(v358)
	goto L29
L31:
	;
	v356 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+23)) = uint16(v356)
	goto L29
L32:
	;
	v377 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v173)+172)) = v377
	v379 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+136)) = v379
	*(*int64)(unsafe.Add(mBase, uint32(v173)+128)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v173)+168)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v173)+180)) = v377
	*(*int64)(unsafe.Add(mBase, uint32(v173)+188)) = v377
	*(*int64)(unsafe.Add(mBase, uint32(v173)+196)) = v377
	*(*int64)(unsafe.Add(mBase, uint32(v173)+204)) = v377
	*(*int64)(unsafe.Add(mBase, uint32(v173)+212)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v173)+220)) = v379
	v397 = v173 + int32(28)
	v398 = F_vacuum_get_cutoffs(m, l0, l1, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)) = uint8(v398)
	v402 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+108)) = v402
	v405 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405<<(uint(int32(2))%32))+uint32(_c_F_heap_vacuum_rel[14])))
	goto L35
L35:
	;
	v409 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+64)) = uint8(v409)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+52)) = v408
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v173)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+56)) = v412
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v416 = v414 & int32(256)
	if v416 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v417 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)) = uint8(v417)
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+21)) = uint8(base.B2i32(v416 == int32(0)))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+248)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v173)+240)) = int64(4294967295)
	v426 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	if base.F64_eq(v426, float64(0)) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v73 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L40:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)))
	if v429 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v173)+108))
	if base.Ui32(v430) < base.Ui32(int32(_a_F_heap_vacuum_rel_4)) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	if base.Ui32(int32(3)) <= base.Ui32(v433) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_visibilitymap_count(m, v459, v52+int32(960), v52+int32(528))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L14
	} else {
		goto L55
	}
L44:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v173)+44))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v436))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v433)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v173)+32))
	if v449 == int32(0) {
		goto L39
	} else {
		goto L52
	}
L47:
	;
	if v448 != 0 {
		goto L43
	} else {
		goto L51
	}
L48:
	;
	v448 = base.B2i32(base.Ui32(v433) < base.Ui32(v436))
	goto L47
L49:
	;
	goto L50
L50:
	;
	v448 = int32(base.Ui32(v433-v436) >> (uint(int32(31)) % 32))
	goto L47
L51:
	;
	goto L46
L52:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v173)+48))
	goto L53
L53:
	;
	if int32(base.Ui32(v449-v452)>>(uint(int32(31))%32)) == int32(0) {
		goto L39
	} else {
		goto L54
	}
L54:
	;
	goto L43
L55:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v52)+960))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v52)+528))
	v472 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_u(v466-v467), float64(0.2)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+244)) = v472
	if v472 == int32(0) {
		goto L39
	} else {
		goto L56
	}
L56:
	;
	v477 = Fn13966(m, int64(32))
	mBase = m.M
	goto L57
L57:
	;
	v479 = v477 & int32(4095)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+240)) = v479
	v481 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	v484 = base.I32_trunc_sat_f64_u(base.F64_mul(v481, float64(4096)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+248)) = v484
	*(*int32)(unsafe.Add(mBase, uint32(v173)+252)) = base.I32_trunc_sat_f32_u(base.F32_mul(base.F32_add(base.F32_mul(base.F32_convert_i32_u(v479), float32(-0.00024414062)), float32(1)), base.F32_convert_i32_u(v484)))
	goto L39
L58:
	;
	v528 = F_lazy_check_wraparound_failsafe(m, v173)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L14
	} else {
		goto L70
	}
L59:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)))
	v502 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	if v502 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v173)+68))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+520)) = v507
	*(*int64)(unsafe.Add(mBase, uint32(v52)+512)) = v506
	v513 = v499 & int32(1)
	if v513 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v514 = int32(_a_F_heap_vacuum_rel_5)
	goto L64
L63:
	;
	v514 = int32(_a_F_heap_vacuum_rel_6)
	goto L64
L64:
	;
	F_errmsg(m, v514, v52+int32(512))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	if v513 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v522 = int32(817)
	goto L68
L67:
	;
	v522 = int32(822)
	goto L68
L68:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), v522, int32(_a_F_heap_vacuum_rel_8))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	goto L58
L70:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[15]))
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[16]))
	if v531 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v536 = v531
	goto L73
L72:
	;
	v536 = v533
	goto L73
L73:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[4]))
	if v538 == int32(4) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v541 = v536
	goto L76
L75:
	;
	v541 = v533
	goto L76
L76:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v542 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v1301 = v173 + int32(60)
	v1303 = v173 + int32(56)
	v1305 = v173 + int32(172)
	v1309 = v173 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+100)) = v1299
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v173)+244))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v173)+108))
	v1313 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+924)) = v1313
	v1316 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+920)) = v1316
	v1319 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[18]))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+912)) = v1319
	v1321 = base.I64_extend_i32_u(v1312)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+536)) = v1321
	*(*int64)(unsafe.Add(mBase, uint32(v52)+528)) = int64(1)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v1326 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1325))))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+544)) = v1326
	goto L209
L78:
	;
	v1239 = F_palloc(m, int32(16))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L14
	} else {
		goto L205
	}
L79:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v545 < int32(2) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+23)))
	if v548 != int32(1) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+48))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+118)))
	if v553 == int32(116) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	if v1179 == int32(0) {
		goto L78
	} else {
		goto L203
	}
L83:
	;
	if v542 == int32(0) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+96)))
	if v579 != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	v560 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L14
	} else {
		goto L87
	}
L87:
	;
	if v560 == int32(0) {
		goto L82
	} else {
		goto L88
	}
L88:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+496)) = v564
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_9), v52+int32(496))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L14
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3500), int32(_a_F_heap_vacuum_rel_10))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L14
	} else {
		goto L90
	}
L90:
	;
	goto L82
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = v1128
	goto L82
L92:
	;
	v580 = int32(17)
	goto L94
L93:
	;
	v580 = int32(13)
	goto L94
L94:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v582 = F_palloc0(m, v545)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	if v545 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v745 = F_palloc0(m, int32(72))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L14
	} else {
		goto L123
	}
L97:
	;
	F_pfree(m, v582)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L14
	} else {
		goto L122
	}
L98:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[19])))
	if v587&int32(1) == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[20]))
	if v593 == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v605 = v4
	v614 = v4
	v618 = v4
	goto L101
L101:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v576+v614<<(uint(int32(2))%32))))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+204))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+29)))
	if v650 == int32(0) {
		v670 = v605
		v671 = v618
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v670 < v671 {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	v673 = v614 + int32(1)
	if v673 != v545 {
		v605 = v670
		v614 = v673
		v618 = v671
		goto L101
	} else {
		goto L107
	}
L104:
	;
	v654 = F_RelationGetNumberOfBlocksInFork(m, v648, int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L14
	} else {
		goto L105
	}
L105:
	;
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[21]))
	if base.Ui32(v654) < base.Ui32(v657) {
		v670 = v605
		v671 = v618
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v660 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v614+v582))) = uint8(v660)
	v670 = v605 + base.B2i32(v650&int32(6) != int32(0))
	v671 = v650&v660 + v618
	goto L103
L107:
	;
	goto L102
L108:
	;
	v676 = v671
	goto L110
L109:
	;
	v676 = v670
	goto L110
L110:
	;
	v678 = v676 - int32(1)
	if v678 <= int32(0) {
		goto L97
	} else {
		goto L111
	}
L111:
	;
	if base.Ui32(v542) < base.Ui32(v678) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v682 = v542
	goto L114
L113:
	;
	v682 = v678
	goto L114
L114:
	;
	if int32(0) < v542 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v685 = v682
	goto L117
L116:
	;
	v685 = v678
	goto L117
L117:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[20]))
	if v685 < v687 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v689 = v685
	goto L120
L119:
	;
	v689 = v687
	goto L120
L120:
	;
	if int32(0) < v689 {
		goto L96
	} else {
		goto L121
	}
L121:
	;
	goto L97
L122:
	;
	v1128 = int32(0)
	goto L91
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+52)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v745)+36)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v745)+12)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v745)+8)) = v576
	*(*int32)(unsafe.Add(mBase, uint32(v745)+4)) = v551
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[22]))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v754)+72)) = v755 + int32(1)
	goto L124
L124:
	;
	v761 = F_CreateParallelContext(m, int32(_a_F_heap_vacuum_rel_11), int32(_a_F_heap_vacuum_rel_12), v689)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L14
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745))) = v761
	v765 = F_mul_size(m, int32(48), v545)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L14
	} else {
		goto L126
	}
L126:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v761)+36))
	v772 = F_add_size(m, v767, (v765+int32(31))&int32(-32))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L14
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+36)) = v772
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v761)+40))
	v777 = F_add_size(m, v775, int32(1))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L14
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+40)) = v777
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v761)+36))
	v782 = F_add_size(m, v780, int32(96))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L14
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+36)) = v782
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v761)+40))
	v787 = F_add_size(m, v785, int32(1))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L14
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+40)) = v787
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v761)+36))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	v793 = F_mul_size(m, int32(128), v792)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L14
	} else {
		goto L131
	}
L131:
	;
	v799 = F_add_size(m, v790, (v793+int32(31))&int32(-32))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+36)) = v799
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v761)+40))
	v804 = F_add_size(m, v802, int32(1))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+40)) = v804
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v761)+36))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	v810 = F_mul_size(m, int32(32), v809)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L14
	} else {
		goto L134
	}
L134:
	;
	v816 = F_add_size(m, v807, (v810+int32(31))&int32(-32))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L14
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+36)) = v816
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v761)+40))
	v821 = F_add_size(m, v819, int32(1))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L14
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+40)) = v821
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	if v825 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v761)+36))
	v827 = F_strlen(m, v825)
	mBase = m.M
	v832 = F_add_size(m, v826, v827&int32(-32)+int32(32))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L14
	} else {
		goto L140
	}
L138:
	;
	v840 = v4
	goto L139
L139:
	;
	F_InitializeParallelDSM(m, v761)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L14
	} else {
		goto L142
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+36)) = v832
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v761)+40))
	v837 = F_add_size(m, v835, int32(1))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L14
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761)+40)) = v837
	v840 = v827
	goto L139
L142:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	v846 = F_shm_toc_allocate(m, v845, v765)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L14
	} else {
		goto L144
	}
L143:
	;
	v880 = int32(1)
	if v545 <= v880 {
		goto L154
	} else {
		goto L155
	}
L144:
	;
	if v765&int32(3)|(v846&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v765))) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if v765 == int32(0) {
		goto L143
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if v765 == int32(0) {
		goto L143
	} else {
		goto L153
	}
L148:
	;
	v860 = v765 + v846
	v862 = v846 + int32(4)
	if base.Ui32(v862) < base.Ui32(v860) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v864 = v860
	goto L151
L150:
	;
	v864 = v862
	goto L151
L151:
	;
	v869 = (v846^int32(-1)+v864)&int32(-4) + int32(4)
	if v869 == int32(0) {
		goto L143
	} else {
		goto L152
	}
L152:
	;
	base.MemoryFill(m, v846, int32(0), v869)
	goto L143
L153:
	;
	base.MemoryFill(m, v846, int32(0), v765)
	goto L143
L154:
	;
	v883 = v880
	goto L156
L155:
	;
	v883 = v545
	goto L156
L156:
	;
	v884 = int32(0)
	v891 = v884
	v904 = v884
	goto L157
L157:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904+v582))))
	if v936 != int32(1) {
		v967 = v891
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	F_shm_toc_insert(m, v972, int64(5), v846)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L14
	} else {
		goto L169
	}
L159:
	;
	v970 = v904 + int32(1)
	if v970 != v883 {
		v891 = v967
		v904 = v970
		goto L157
	} else {
		goto L168
	}
L160:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v576+v904<<(uint(int32(2))%32))))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v942)+204))
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+27)))
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+29)))
	if v945&int32(1) != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v745)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v745)+40)) = v948 + int32(1)
	goto L163
L162:
	;
	goto L163
L163:
	;
	if v945&int32(4) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v745)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v745)+44)) = v954 + int32(1)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v958 = v944 + v891
	if v945&int32(2) == int32(0) {
		v967 = v958
		goto L159
	} else {
		goto L167
	}
L167:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v745)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v745)+48)) = v963 + int32(1)
	v967 = v958
	goto L159
L168:
	;
	goto L158
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+20)) = v846
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	v979 = F_shm_toc_allocate(m, v977, int32(72))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L14
	} else {
		goto L170
	}
L170:
	;
	v981 = int32(0)
	base.MemoryFill(m, v979, v981, int32(72))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v551)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v979)+4)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v979))) = v984
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v989 == v981 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v979)+8)) = v994
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[16]))
	if int32(0) < v967 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v994 = int64(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v993 = *(*int64)(unsafe.Add(mBase, uint32(v989)+392))
	v994 = v993
	goto L171
L175:
	;
	if v689 < v967 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v1003 = v997
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979)+28)) = v1003
	v1006 = v541 << (uint(int32(10)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v979)+56)) = v1006
	v1008 = F_TidStoreCreateShared(m, v1006)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L14
	} else {
		goto L181
	}
L178:
	;
	v1001 = v689
	goto L180
L179:
	;
	v1001 = v967
	goto L180
L180:
	;
	v1002 = base.I32_div_s(v997, v1001)
	v1003 = v1002
	goto L177
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+24)) = v1008
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+4))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1012)))
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979)+52)) = v1013
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+8))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+28))
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v979)+48)) = v1017
	if v581 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v1024 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v979)+36)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v979)+40)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v979)+32)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v979)+44)) = v1024
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	F_shm_toc_insert(m, v1031, int64(1), v979)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L14
	} else {
		goto L188
	}
L185:
	;
	v1023 = int32(0)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	v1023 = v1022
	goto L184
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+16)) = v979
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	v1039 = F_mul_size(m, int32(128), v1038)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L14
	} else {
		goto L189
	}
L189:
	;
	v1041 = F_shm_toc_allocate(m, v1036, v1039)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L14
	} else {
		goto L190
	}
L190:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	F_shm_toc_insert(m, v1043, int64(3), v1041)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L14
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+28)) = v1041
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	v1051 = F_mul_size(m, int32(32), v1050)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L14
	} else {
		goto L192
	}
L192:
	;
	v1053 = F_shm_toc_allocate(m, v1048, v1051)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L14
	} else {
		goto L193
	}
L193:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	F_shm_toc_insert(m, v1055, int64(4), v1053)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L14
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+32)) = v1053
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	if v1061 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	v1064 = v840 + int32(1)
	v1065 = F_shm_toc_allocate(m, v1062, v1064)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L14
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v1128 = v745
	goto L91
L198:
	;
	if v1064 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	base.MemoryCopy(m, v1065, v1068, v1064)
	goto L201
L200:
	;
	goto L201
L201:
	;
	v1071 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1065+v840))) = uint8(v1071)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v761)+52))
	F_shm_toc_insert(m, v1073, int64(2), v1065)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L14
	} else {
		goto L202
	}
L202:
	;
	goto L197
L203:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v173+int32(104)))) = v1184 + int32(56)
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+24))
	goto L204
L204:
	;
	v1299 = v1188
	goto L77
L205:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1239)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1239))) = v541 << (uint(int32(10)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+104)) = v1239
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1239)))
	v1248 = F_TidStoreCreateLocal(m, v1247)
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L14
	} else {
		goto L206
	}
L206:
	;
	v1299 = v1248
	goto L77
L207:
	;
	v1514 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+236)) = v1514
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+232)) = uint16(v1514)
	*(*int64)(unsafe.Add(mBase, uint32(v173)+224)) = int64(-1)
	v1521 = v173 + int32(88)
	v1523 = v52 + int32(996)
	v1524 = int32(1)
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v1530 = F_read_stream_begin_relation(m, v1524, v1525, v1526, v1514, int32(188), v173, v1524)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L14
	} else {
		goto L224
	}
L208:
	;
	goto L207
L209:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v1342 == int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v1346&int32(1) == int32(0) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v1351 = int32(_a_F_heap_vacuum_rel_1)
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1354 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1353 + v1354
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1342)))
	*(*int32)(unsafe.Add(mBase, uint32(v1342))) = v1357 + v1354
	v1361 = int32(0)
	v1364 = base.AtomicRmwOr32(m, v1361, int32(_a_F_heap_vacuum_rel_2), v1361)
	goto L213
L212:
	;
	v1491 = int32(0)
	v1494 = base.AtomicRmwOr32(m, v1491, int32(_a_F_heap_vacuum_rel_2), v1491)
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1342)))
	v1496 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1342))) = v1495 + v1496
	v1499 = int32(_a_F_heap_vacuum_rel_1)
	v1501 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1501 - v1496
	goto L208
L213:
	;
	goto L215
L215:
	;
	goto L216
L216:
	;
	v1456 = int32(0)
	v1459 = v1313
	goto L221
L221:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(912)+v1459<<(uint(int32(2))%32))))
	v1469 = int32(3)
	v1475 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(528)+v1459<<(uint(v1469)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1342+int32(232)+v1468<<(uint(v1469)%32)))) = v1475
	v1477 = int32(1)
	v1480 = v1456 + v1477
	if v1480 != int32(3) {
		v1456 = v1480
		v1459 = v1459 + v1477
		goto L221
	} else {
		goto L223
	}
L222:
	;
	goto L212
L223:
	;
	goto L222
L224:
	;
	v1538 = int32(0)
	v1541 = v4
	goto L225
L225:
	;
	v1582 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+908)) = v1582
	F_vacuum_delay_point(m, v1582)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L14
	} else {
		goto L227
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+84)) = int32(-1)
	v3154 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	if v3154 != 0 {
		goto L521
	} else {
		goto L522
	}
L227:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1309)))
	v1588 = int32(0)
	if base.B2i32(v1587 == v1588)|v1587&int32(_a_F_heap_vacuum_rel_13) == v1588 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1595 = F_lazy_check_wraparound_failsafe(m, v173)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L14
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v1598 = *(*int64)(unsafe.Add(mBase, uint32(v1597)+8))
	if v1598 <= int64(0) {
		v1667 = v1541
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	v1670 = F_read_stream_next_buffer(m, v1530, v52+int32(908))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L14
	} else {
		goto L246
	}
L233:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v173)+100))
	v1602 = F_TidStoreMemoryUsage(m, v1601)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L14
	} else {
		goto L234
	}
L234:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1604)))
	if base.Ui32(v1602) <= base.Ui32(v1605) {
		v1667 = v1541
		goto L232
	} else {
		goto L235
	}
L235:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	if v1607 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	F_ReleaseBuffer(m, v1607)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L14
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v1612 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+22)) = uint8(v1612)
	F_lazy_vacuum(m, v173)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L14
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+924)) = int32(0)
	goto L238
L240:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_FreeSpaceMapVacuumRange(m, v1616, v1541, v1538+int32(1))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L14
	} else {
		goto L241
	}
L241:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v1625 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1667 = v1538
	goto L232
L243:
	;
	goto L242
L244:
	;
	v1629 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v1629&int32(1) == int32(0) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v1634 = int32(_a_F_heap_vacuum_rel_1)
	v1636 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1637 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1636 + v1637
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1625)))
	*(*int32)(unsafe.Add(mBase, uint32(v1625))) = v1640 + v1637
	v1644 = int32(0)
	v1646 = int32(_a_F_heap_vacuum_rel_2)
	v1647 = base.AtomicRmwOr32(m, v1644, v1646, v1644)
	*(*int64)(unsafe.Add(mBase, uint32(v1625+v1644)+232)) = int64(1)
	v1655 = base.AtomicRmwOr32(m, v1644, v1646, v1644)
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1625)))
	*(*int32)(unsafe.Add(mBase, uint32(v1625))) = v1656 + v1637
	v1662 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1662 - v1637
	goto L243
L246:
	;
	if v1670 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v52)+908))
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672))))
	F_CheckBufferIsPinnedOnce(m, v1670)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L14
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	goto L226
L250:
	;
	if v1670 < int32(0) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	if v1670 < int32(0) {
		goto L256
	} else {
		goto L257
	}
L252:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[24]))
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1679+(v1670^int32(-1))<<(uint(int32(2))%32))))
	v1693 = v1685
	goto L251
L253:
	;
	goto L254
L254:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[25]))
	v1693 = v1687 + v1670<<(uint(int32(13))%32) + int32(-8192)
	goto L251
L255:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1309)))
	v1714 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1309))) = v1713 + v1714
	v1718 = v1673 & v1714
	if v1718 != 0 {
		goto L259
	} else {
		goto L260
	}
L256:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[26]))
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1697+(v1670^int32(-1))<<(uint(int32(6))%32))+16))
	v1712 = v1703
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[27]))
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1705+v1670<<(uint(int32(6))%32)+int32(-64))+16))
	v1712 = v1711
	goto L255
L259:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v173)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+116)) = v1719 + int32(1)
	goto L261
L260:
	;
	goto L261
L261:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v1727 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+92)) = int32(1)
	v1770 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+88)) = uint16(v1770)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+84)) = v1712
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_visibilitymap_pin(m, v1773, v1712, v52+int32(924))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L14
	} else {
		goto L266
	}
L263:
	;
	goto L262
L264:
	;
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v1731&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1736 = int32(_a_F_heap_vacuum_rel_1)
	v1738 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1739 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1738 + v1739
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1727)))
	*(*int32)(unsafe.Add(mBase, uint32(v1727))) = v1742 + v1739
	v1746 = int32(0)
	v1748 = int32(_a_F_heap_vacuum_rel_2)
	v1749 = base.AtomicRmwOr32(m, v1746, v1748, v1746)
	*(*int64)(unsafe.Add(mBase, uint32(v1727+int32(16))+232)) = base.I64_extend_i32_u(v1712)
	v1757 = base.AtomicRmwOr32(m, v1746, v1748, v1746)
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1727)))
	*(*int32)(unsafe.Add(mBase, uint32(v1727))) = v1758 + v1739
	v1764 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1764 - v1739
	goto L263
L266:
	;
	v1778 = F_ConditionalLockBufferForCleanup(m, v1670)
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L14
	} else {
		goto L267
	}
L267:
	;
	if v1778 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	F_LockBuffer(m, v1670, int32(1))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L14
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+14)))
	if v1785 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L271:
	;
	goto L270
L272:
	;
	v3058 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v3058 != 0 {
		goto L492
	} else {
		goto L493
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1608)) = v2470
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v2512 = *(*int32)(unsafe.Add(mBase, uint32(v173)+52))
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if v2515 != 0 {
		goto L388
	} else {
		goto L389
	}
L274:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1303)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1608)) = v1941
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1301)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1600)) = v1943
	v1950 = int32(base.Ui32(v1940+int32(_a_F_heap_vacuum_rel_14))>>(uint(int32(2))%32)) & int32(_a_F_heap_vacuum_rel_15)
	if v1950 != 0 {
		goto L330
	} else {
		goto L331
	}
L275:
	;
	if v1778 != 0 {
		v2470 = v1794
		goto L273
	} else {
		goto L325
	}
L276:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_RecordPageWithFreeSpace(m, v1937, v1712, v1934)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L14
	} else {
		goto L324
	}
L277:
	;
	F_UnlockReleaseBuffer(m, v1670)
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L14
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	v1795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v1795) {
		goto L275
	} else {
		goto L283
	}
L280:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v1791 = F_GetRecordedFreeSpace(m, v1790, v1712)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L14
	} else {
		goto L281
	}
L281:
	;
	if v1791 != 0 {
		v1538 = v1712
		v1541 = v1667
		goto L225
	} else {
		goto L282
	}
L282:
	;
	v1934 = int32(_a_F_heap_vacuum_rel_16)
	goto L276
L283:
	;
	if v1778 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	F_LockBuffer(m, v1670, int32(0))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L14
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693)+10)))
	if v1810&int32(4) == int32(0) {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	F_LockBuffer(m, v1670, int32(2))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L14
	} else {
		goto L288
	}
L288:
	;
	v1806 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v1806) {
		v1940 = v1806
		goto L274
	} else {
		goto L289
	}
L289:
	;
	goto L286
L290:
	;
	v1815 = int32(_a_F_heap_vacuum_rel_1)
	v1817 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1817 + int32(1)
	F_MarkBufferDirty(m, v1670)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L14
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v1868 = int32(4)
	v1869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+14)))
	v1870 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+12)))
	v1871 = v1869 - v1870
	if v1871 <= v1868 {
		goto L305
	} else {
		goto L306
	}
L293:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+48))
	v1825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1824)+118)))
	if v1825 != int32(112) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)))
	v1842 = v1840 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)) = uint16(v1842)
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v1848 = F_visibilitymap_set(m, v1844, v1712, v1670, int64(0), v1794, int32(0), int32(3))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L14
	} else {
		goto L303
	}
L295:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[28]))
	if v1829 <= int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+32))
	if v1832 != 0 {
		goto L294
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1693)+4))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1693)))
	if v1834|v1835 != 0 {
		goto L294
	} else {
		goto L301
	}
L299:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+40))
	if v1833 != 0 {
		goto L294
	} else {
		goto L300
	}
L300:
	;
	goto L298
L301:
	;
	F_log_newpage_buffer(m, v1670, int32(1))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L14
	} else {
		goto L302
	}
L302:
	;
	goto L294
L303:
	;
	v1850 = int32(_a_F_heap_vacuum_rel_1)
	v1852 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1853 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1852 - v1853
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v173)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+128)) = v1856 + v1853
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v173)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+132)) = v1860 + v1853
	goto L292
L304:
	;
	F_UnlockReleaseBuffer(m, v1670)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L14
	} else {
		goto L323
	}
L305:
	;
	v1874 = v1868
	goto L307
L306:
	;
	v1874 = v1871
	goto L307
L307:
	;
	v1876 = v1874 - int32(4)
	if v1876 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1931 = int32(0)
	goto L304
L309:
	;
	goto L310
L310:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1870) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v1931 = v1876
	goto L304
L312:
	;
	v1887 = int32(base.Ui32(v1870+int32(_a_F_heap_vacuum_rel_14)) >> (uint(int32(2)) % 32))
	goto L314
L313:
	;
	v1887 = int32(0)
	goto L314
L314:
	;
	if base.Ui32(v1887&int32(_a_F_heap_vacuum_rel_15)) < base.Ui32(int32(291)) {
		goto L311
	} else {
		goto L315
	}
L315:
	;
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693)+10)))
	if v1892&int32(1) == int32(0) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1931 = int32(0)
	goto L304
L317:
	;
	goto L318
L318:
	;
	v1901 = int32(1)
	goto L319
L319:
	;
	v1910 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693+int32(20)+v1901&int32(_a_F_heap_vacuum_rel_15)<<(uint(int32(2))%32))+1)))
	if v1910&int32(384) == int32(0) {
		goto L311
	} else {
		goto L321
	}
L320:
	;
	v1931 = int32(0)
	goto L304
L321:
	;
	v1916 = v1901 + int32(1)
	v1917 = int32(_a_F_heap_vacuum_rel_15)
	if base.Ui32(v1916&v1917) <= base.Ui32(v1887&v1917) {
		v1901 = v1916
		goto L319
	} else {
		goto L322
	}
L322:
	;
	goto L320
L323:
	;
	v1934 = v1931
	goto L276
L324:
	;
	v1538 = v1712
	v1541 = v1667
	goto L225
L325:
	;
	v1940 = v1795
	goto L274
L326:
	;
	v2453 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1521))) = uint16(v2453)
	F_LockBuffer(m, v1670, v2453)
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L14
	} else {
		goto L386
	}
L327:
	;
	v2430 = *(*int64)(unsafe.Add(mBase, uint32(v173)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+200)) = v2430 + v2419
	v2433 = *(*int64)(unsafe.Add(mBase, uint32(v173)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+208)) = v2433 + v2422
	v2436 = *(*int64)(unsafe.Add(mBase, uint32(v173)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+216)) = v2436 + base.I64_extend_i32_s(v2387)
	if int32(0) < v2387 {
		goto L380
	} else {
		goto L381
	}
L328:
	;
	if v2084 <= int32(0) {
		goto L358
	} else {
		goto L359
	}
L329:
	;
	v2158 = int32(0)
	v2159 = base.B2i32(v2158 < v2132)
	if v2158 < v2132 {
		goto L355
	} else {
		goto L356
	}
L330:
	;
	v1952 = int32(base.Ui32(v1712) >> (uint(int32(16)) % 32))
	v1955 = int32(0)
	v1963 = int32(1)
	v1967 = v1955
	v1968 = v1955
	v1971 = v1955
	v1984 = v1955
	v1997 = v1955
	goto L333
L331:
	;
	goto L332
L332:
	;
	v2100 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1521))) = uint16(v2100)
	v2103 = int64(0)
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v2108 != 0 {
		v2383 = v2100
		v2387 = v2100
		v2388 = v2100
		v2419 = v2103
		v2422 = v2103
		goto L327
	} else {
		goto L354
	}
L333:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1521))) = uint16(v1963)
	v2015 = v1693 + int32(20) + v1963&int32(_a_F_heap_vacuum_rel_15)<<(uint(int32(2))%32)
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v2015)))
	switch int32(base.Ui32(v2016)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L337
	case 1:
		goto L336
	case 2:
		goto L338
	default:
		v2081 = v1967
		v2082 = v1968
		v2083 = v1971
		v2084 = v1984
		v2085 = v1997
		goto L335
	}
L334:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1600))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2093 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1521))) = uint16(v2093)
	*(*int32)(unsafe.Add(mBase, uint32(v1303))) = v2092
	*(*int32)(unsafe.Add(mBase, uint32(v1301))) = v2091
	v2097 = base.I64_extend_i32_s(v2085)
	v2098 = base.I64_extend_i32_s(v2083)
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v2099 != 0 {
		goto L328
	} else {
		goto L353
	}
L335:
	;
	v2087 = v1963 + int32(1)
	if base.Ui32(v2087&int32(_a_F_heap_vacuum_rel_15)) <= base.Ui32(v1950) {
		v1963 = v2087
		v1967 = v2081
		v1968 = v2082
		v1971 = v2083
		v1984 = v2084
		v1997 = v2085
		goto L333
	} else {
		goto L352
	}
L336:
	;
	v2081 = v1967
	v2082 = int32(1)
	v2083 = v1971
	v2084 = v1984
	v2085 = v1997
	goto L335
L337:
	;
	v2038 = F_heap_tuple_should_freeze(m, v1693+v2016&int32(_a_F_heap_vacuum_rel_17), v397, v52+int32(1608), v52+int32(1600))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L14
	} else {
		goto L339
	}
L338:
	;
	v2025 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(960)+v1984<<(uint(v2025)%32)))) = uint16(v1963)
	v2081 = v1967
	v2082 = v1968
	v2083 = v1971
	v2084 = v1984 + v2025
	v2085 = v1997
	goto L335
L339:
	;
	if v2038 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)))
	if v2040 != 0 {
		goto L326
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+936)) = uint16(v1963)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+934)) = uint16(v1712)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+932)) = uint16(v1952)
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v2015)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+928)) = int32(base.Ui32(v2044) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+944)) = v1693 + v2044&int32(_a_F_heap_vacuum_rel_17)
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2052)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+940)) = v2053
	v2055 = int32(1)
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v173)+36))
	v2059 = F_HeapTupleSatisfiesVacuum(m, v52+int32(928), v2058, v1670)
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L14
	} else {
		goto L348
	}
L343:
	;
	goto L342
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L14
	} else {
		goto L349
	}
L345:
	;
	v2081 = v1967
	v2082 = v2055
	v2083 = v1971
	v2084 = v1984
	v2085 = v1997 + int32(1)
	goto L335
L346:
	;
	v2081 = v1967 + int32(1)
	v2082 = v2055
	v2083 = v1971
	v2084 = v1984
	v2085 = v1997
	goto L335
L347:
	;
	v2081 = v1967
	v2082 = v2055
	v2083 = v1971 + int32(1)
	v2084 = v1984
	v2085 = v1997
	goto L335
L348:
	;
	switch v2059 {
	case 0:
		goto L346
	case 1, 4:
		goto L347
	case 2:
		goto L345
	case 3:
		v2081 = v1967
		v2082 = v2055
		v2083 = v1971
		v2084 = v1984
		v2085 = v1997
		goto L335
	default:
		goto L344
	}
L349:
	;
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_18), int32(0))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L14
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(2369), int32(_a_F_heap_vacuum_rel_19))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L14
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	goto L334
L353:
	;
	v2111 = v2082
	v2115 = v2081
	v2132 = v2084
	v2147 = v2098
	v2150 = v2097
	goto L329
L354:
	;
	v2111 = v2100
	v2115 = v2100
	v2132 = v2100
	v2147 = v2103
	v2150 = v2103
	goto L329
L355:
	;
	v2162 = v2132
	goto L357
L356:
	;
	v2162 = v2158
	goto L357
L357:
	;
	v2383 = v2159
	v2387 = v2162 + v2115
	v2388 = v2111 | v2159
	v2419 = v2147
	v2422 = v2150
	goto L327
L358:
	;
	v2383 = int32(0)
	v2387 = v2081
	v2388 = v2082
	v2419 = v2098
	v2422 = v2097
	goto L327
L359:
	;
	goto L360
L360:
	;
	v2167 = int32(1)
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v173)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+140)) = v2168 + v2167
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1584)) = int64(25769803783)
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+100))
	F_TidStoreSetBlockOffsets(m, v2174, v1712, v52+int32(960), v2084)
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L14
	} else {
		goto L361
	}
L361:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v2180 = base.I64_extend_i32_u(v2084)
	v2181 = *(*int64)(unsafe.Add(mBase, uint32(v2179)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2179)+8)) = v2180 + v2181
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v2185 = *(*int64)(unsafe.Add(mBase, uint32(v2184)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+928)) = v2185
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v173)+100))
	v2188 = F_TidStoreMemoryUsage(m, v2187)
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L14
	} else {
		goto L362
	}
L362:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+936)) = base.I64_extend_i32_u(v2188)
	goto L365
L363:
	;
	v2378 = *(*int64)(unsafe.Add(mBase, uint32(v173)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+192)) = v2378 + v2180
	v2383 = v2167
	v2387 = v2081
	v2388 = v2082
	v2419 = v2098
	v2422 = v2097
	goto L327
L364:
	;
	goto L363
L365:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v2206 == int32(0) {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v2210 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v2210&int32(1) == int32(0) {
		goto L364
	} else {
		goto L367
	}
L367:
	;
	v2215 = int32(_a_F_heap_vacuum_rel_1)
	v2217 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v2218 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2217 + v2218
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2206)))
	*(*int32)(unsafe.Add(mBase, uint32(v2206))) = v2221 + v2218
	v2225 = int32(0)
	v2228 = base.AtomicRmwOr32(m, v2225, int32(_a_F_heap_vacuum_rel_2), v2225)
	goto L369
L368:
	;
	v2355 = int32(0)
	v2358 = base.AtomicRmwOr32(m, v2355, int32(_a_F_heap_vacuum_rel_2), v2355)
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v2206)))
	v2360 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2206))) = v2359 + v2360
	v2363 = int32(_a_F_heap_vacuum_rel_1)
	v2365 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2365 - v2360
	goto L364
L369:
	;
	goto L371
L371:
	;
	goto L372
L372:
	;
	v2320 = int32(0)
	v2323 = int32(0)
	goto L377
L377:
	;
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1584)+v2323<<(uint(int32(2))%32))))
	v2333 = int32(3)
	v2339 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(928)+v2323<<(uint(v2333)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2206+int32(232)+v2332<<(uint(v2333)%32)))) = v2339
	v2341 = int32(1)
	v2344 = v2320 + v2341
	if v2344 != int32(2) {
		v2320 = v2344
		v2323 = v2323 + v2341
		goto L377
	} else {
		goto L379
	}
L378:
	;
	goto L368
L379:
	;
	goto L378
L380:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v173)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+144)) = v2442 + int32(1)
	goto L382
L381:
	;
	goto L382
L382:
	;
	if v2388&int32(1) != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+148)) = v1712 + int32(1)
	goto L385
L384:
	;
	goto L385
L385:
	;
	v2451 = int32(0)
	v3011 = v2383
	v3015 = v2451
	v3018 = v2451
	goto L272
L386:
	;
	F_LockBufferForCleanup(m, v1670)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L14
	} else {
		goto L387
	}
L387:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	v2470 = v2460
	goto L273
L388:
	;
	v2516 = int32(2)
	goto L390
L389:
	;
	v2516 = int32(3)
	goto L390
L390:
	;
	F_heap_page_prune_and_freeze(m, v2511, v1670, v2512, v2516, v397, v52+int32(960), int32(1), v1521, v1303, v1301)
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L14
	} else {
		goto L391
	}
L391:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v52)+968))
	if int32(0) < v2522 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v173)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+124)) = v2525 + int32(1)
	goto L394
L393:
	;
	goto L394
L394:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	if int32(0) < v2529 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v173)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+140)) = v2532 + int32(1)
	F_pg_qsort(m, v1523, v2529, int32(2), int32(189))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L14
	} else {
		goto L398
	}
L396:
	;
	v2747 = v2529
	v2749 = v2522
	goto L397
L397:
	;
	v2750 = *(*int64)(unsafe.Add(mBase, uint32(v173)+176))
	v2751 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+960)))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+176)) = v2750 + v2751
	v2754 = *(*int64)(unsafe.Add(mBase, uint32(v173)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+184)) = v2754 + base.I64_extend_i32_s(v2749)
	v2758 = *(*int64)(unsafe.Add(mBase, uint32(v173)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+192)) = v2758 + base.I64_extend_i32_s(v2747)
	v2762 = *(*int64)(unsafe.Add(mBase, uint32(v173)+200))
	v2763 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+972)))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+200)) = v2762 + v2763
	v2766 = *(*int64)(unsafe.Add(mBase, uint32(v173)+208))
	v2767 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+976)))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+208)) = v2766 + v2767
	v2770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+988)))
	if v2770 == int32(1) {
		goto L418
	} else {
		goto L419
	}
L398:
	;
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1584)) = int64(25769803783)
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v173)+100))
	F_TidStoreSetBlockOffsets(m, v2543, v1712, v1523, v2540)
	mBase = m.M
	v2545 = m.ExcPending
	if v2545 != 0 {
		goto L14
	} else {
		goto L399
	}
L399:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v2547 = *(*int64)(unsafe.Add(mBase, uint32(v2546)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2546)+8)) = v2547 + base.I64_extend_i32_s(v2540)
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v2552 = *(*int64)(unsafe.Add(mBase, uint32(v2551)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+928)) = v2552
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(v173)+100))
	v2555 = F_TidStoreMemoryUsage(m, v2554)
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L14
	} else {
		goto L400
	}
L400:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+936)) = base.I64_extend_i32_u(v2555)
	goto L403
L401:
	;
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(v52)+968))
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	v2747 = v2746
	v2749 = v2745
	goto L397
L402:
	;
	goto L401
L403:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v2573 == int32(0) {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v2577 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v2577&int32(1) == int32(0) {
		goto L402
	} else {
		goto L405
	}
L405:
	;
	v2582 = int32(_a_F_heap_vacuum_rel_1)
	v2584 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v2585 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2584 + v2585
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v2573)))
	*(*int32)(unsafe.Add(mBase, uint32(v2573))) = v2588 + v2585
	v2592 = int32(0)
	v2595 = base.AtomicRmwOr32(m, v2592, int32(_a_F_heap_vacuum_rel_2), v2592)
	goto L407
L406:
	;
	v2722 = int32(0)
	v2725 = base.AtomicRmwOr32(m, v2722, int32(_a_F_heap_vacuum_rel_2), v2722)
	v2726 = *(*int32)(unsafe.Add(mBase, uint32(v2573)))
	v2727 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2573))) = v2726 + v2727
	v2730 = int32(_a_F_heap_vacuum_rel_1)
	v2732 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2732 - v2727
	goto L402
L407:
	;
	goto L409
L409:
	;
	goto L410
L410:
	;
	v2687 = int32(0)
	v2690 = int32(0)
	goto L415
L415:
	;
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1584)+v2690<<(uint(int32(2))%32))))
	v2700 = int32(3)
	v2706 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(928)+v2690<<(uint(v2700)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2573+int32(232)+v2699<<(uint(v2700)%32)))) = v2706
	v2708 = int32(1)
	v2711 = v2687 + v2708
	if v2711 != int32(2) {
		v2687 = v2711
		v2690 = v2690 + v2708
		goto L415
	} else {
		goto L417
	}
L416:
	;
	goto L406
L417:
	;
	goto L416
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+148)) = v1712 + int32(1)
	goto L420
L419:
	;
	goto L420
L420:
	;
	v2777 = v1673 & int32(2)
	if v2777 != 0 {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v2958 = int32(0)
	v2959 = base.B2i32(v2958 < v2747)
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(v52)+960))
	v2962 = base.B2i32(v2958 < v2960)
	v2963 = int32(1)
	if v1718 == v2958 {
		v3011 = v2959
		v3015 = v2962
		v3018 = v2963
		goto L272
	} else {
		goto L473
	}
L422:
	;
	if v2777 == int32(0) {
		v2839 = v2747
		goto L438
	} else {
		goto L439
	}
L423:
	;
	v2778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+980)))
	if v2778&int32(1) == int32(0) {
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v2783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	v2784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)))
	v2786 = v2784 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)) = uint16(v2786)
	F_MarkBufferDirty(m, v1670)
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L14
	} else {
		goto L425
	}
L425:
	;
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v52)+984))
	if v2783 != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2795 = int32(3)
	goto L428
L427:
	;
	v2795 = int32(1)
	goto L428
L428:
	;
	v2796 = F_visibilitymap_set(m, v2790, v1712, v1670, int64(0), v2470, v2792, v2795)
	mBase = m.M
	v2797 = m.ExcPending
	if v2797 != 0 {
		goto L14
	} else {
		goto L429
	}
L429:
	;
	if v2796&int32(1) == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(v173)+128))
	v2803 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+128)) = v2802 + v2803
	v2807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	if v2807 != v2803 {
		v2957 = int32(0)
		goto L421
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	v2815 = int32(0)
	if v2796&int32(2) != 0 {
		v2957 = v2815
		goto L421
	} else {
		goto L434
	}
L433:
	;
	v2810 = int32(1)
	v2811 = *(*int32)(unsafe.Add(mBase, uint32(v173)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+132)) = v2811 + v2810
	v2957 = v2810
	goto L421
L434:
	;
	v2818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	if v2818&int32(1) == int32(0) {
		v2957 = v2815
		goto L421
	} else {
		goto L435
	}
L435:
	;
	v2823 = int32(1)
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v173)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+136)) = v2824 + v2823
	v2957 = v2823
	goto L421
L436:
	;
	v2901 = int32(0)
	if v2777 == v2901 {
		v2957 = v2901
		goto L421
	} else {
		goto L460
	}
L437:
	;
	v2879 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2880 = m.ExcPending
	if v2880 != 0 {
		goto L14
	} else {
		goto L453
	}
L438:
	;
	if v2839 <= int32(0) {
		goto L436
	} else {
		goto L443
	}
L439:
	;
	v2830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693)+10)))
	if v2830&int32(4) != 0 {
		v2839 = v2747
		goto L438
	} else {
		goto L440
	}
L440:
	;
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v2836 = F_visibilitymap_get_status(m, v2833, v1712, v52+int32(1608))
	mBase = m.M
	v2837 = m.ExcPending
	if v2837 != 0 {
		goto L14
	} else {
		goto L441
	}
L441:
	;
	if v2836 != 0 {
		goto L437
	} else {
		goto L442
	}
L442:
	;
	v2838 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	v2839 = v2838
	goto L438
L443:
	;
	v2842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693)+10)))
	if v2842&int32(4) == int32(0) {
		goto L436
	} else {
		goto L444
	}
L444:
	;
	v2849 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2850 = m.ExcPending
	if v2850 != 0 {
		goto L14
	} else {
		goto L445
	}
L445:
	;
	if v2849 != 0 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v2851 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+468)) = v1712
	*(*int32)(unsafe.Add(mBase, uint32(v52)+464)) = v2851
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_20), v52+int32(464))
	mBase = m.M
	v2858 = m.ExcPending
	if v2858 != 0 {
		goto L14
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	v2865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)))
	v2867 = v2865 & int32(_a_F_heap_vacuum_rel_21)
	*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)) = uint16(v2867)
	F_MarkBufferDirty(m, v1670)
	mBase = m.M
	v2870 = m.ExcPending
	if v2870 != 0 {
		goto L14
	} else {
		goto L451
	}
L449:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(2148), int32(_a_F_heap_vacuum_rel_22))
	mBase = m.M
	v2863 = m.ExcPending
	if v2863 != 0 {
		goto L14
	} else {
		goto L450
	}
L450:
	;
	goto L448
L451:
	;
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2874 = F_visibilitymap_clear(m, v1712, v2872, int32(3))
	mBase = m.M
	v2875 = m.ExcPending
	if v2875 != 0 {
		goto L14
	} else {
		goto L452
	}
L452:
	;
	v2957 = int32(0)
	goto L421
L453:
	;
	if v2879 != 0 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+484)) = v1712
	*(*int32)(unsafe.Add(mBase, uint32(v52)+480)) = v2881
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_23), v52+int32(480))
	mBase = m.M
	v2888 = m.ExcPending
	if v2888 != 0 {
		goto L14
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2898 = F_visibilitymap_clear(m, v1712, v2896, int32(3))
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L14
	} else {
		goto L459
	}
L457:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(2126), int32(_a_F_heap_vacuum_rel_22))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L14
	} else {
		goto L458
	}
L458:
	;
	goto L456
L459:
	;
	v2957 = int32(0)
	goto L421
L460:
	;
	v2904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+980)))
	if v2904&int32(1) == int32(0) {
		v2957 = v2901
		goto L421
	} else {
		goto L461
	}
L461:
	;
	v2909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	if v2909&int32(1) == int32(0) {
		v2957 = v2901
		goto L421
	} else {
		goto L462
	}
L462:
	;
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v2917 = F_visibilitymap_get_status(m, v2914, v1712, v52+int32(1608))
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L14
	} else {
		goto L463
	}
L463:
	;
	if v2917&int32(2) != 0 {
		v2957 = v2901
		goto L421
	} else {
		goto L464
	}
L464:
	;
	v2921 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)))
	if v2921&int32(4) == int32(0) {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v2927 = v2921 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1693)+10)) = uint16(v2927)
	F_MarkBufferDirty(m, v1670)
	mBase = m.M
	v2930 = m.ExcPending
	if v2930 != 0 {
		goto L14
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2936 = F_visibilitymap_set(m, v2931, v1712, v1670, int64(0), v2933, int32(0), int32(3))
	mBase = m.M
	v2937 = m.ExcPending
	if v2937 != 0 {
		goto L14
	} else {
		goto L469
	}
L468:
	;
	goto L467
L469:
	;
	if v2936&int32(1) == int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2942 = int32(1)
	v2943 = *(*int32)(unsafe.Add(mBase, uint32(v173)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+128)) = v2943 + v2942
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(v173)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+132)) = v2947 + v2942
	v2957 = v2942
	goto L421
L471:
	;
	goto L472
L472:
	;
	v2951 = int32(1)
	v2952 = *(*int32)(unsafe.Add(mBase, uint32(v173)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+136)) = v2952 + v2951
	v2957 = v2951
	goto L421
L473:
	;
	if v2957 != 0 {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v2966 = *(*int32)(unsafe.Add(mBase, uint32(v173)+244))
	if v2966 != 0 {
		goto L477
	} else {
		goto L478
	}
L475:
	;
	goto L476
L476:
	;
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(v173)+252))
	if v3003 == int32(0) {
		v3011 = v2959
		v3015 = v2962
		v3018 = v2963
		goto L272
	} else {
		goto L490
	}
L477:
	;
	v2968 = v2966 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+244)) = v2968
	if v2968 != 0 {
		v3011 = v2959
		v3015 = v2962
		v3018 = v2963
		goto L272
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v173)+248))
	if v2971 == int32(0) {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	goto L479
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+240)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v173)+248)) = int64(0)
	v3011 = v2959
	v3015 = v2962
	v3018 = v2963
	goto L272
L482:
	;
	v2976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+96)))
	if v2976 != 0 {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v2977 = int32(17)
	goto L485
L484:
	;
	v2977 = int32(13)
	goto L485
L485:
	;
	v2979 = F_errstart(m, v2977, int32(0))
	mBase = m.M
	v2980 = m.ExcPending
	if v2980 != 0 {
		goto L14
	} else {
		goto L486
	}
L486:
	;
	if v2979 == int32(0) {
		goto L481
	} else {
		goto L487
	}
L487:
	;
	v2983 = *(*int64)(unsafe.Add(mBase, uint32(v173)+68))
	v2984 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+460)) = v2984
	*(*int64)(unsafe.Add(mBase, uint32(v52)+452)) = v2983
	*(*int32)(unsafe.Add(mBase, uint32(v52)+448)) = v1311
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_24), v52+int32(448))
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L14
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(1435), int32(_a_F_heap_vacuum_rel_25))
	mBase = m.M
	v2997 = m.ExcPending
	if v2997 != 0 {
		goto L14
	} else {
		goto L489
	}
L489:
	;
	goto L481
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+252)) = v3003 - int32(1)
	v3011 = v2959
	v3015 = v2962
	v3018 = v2963
	goto L272
L491:
	;
	F_UnlockReleaseBuffer(m, v1670)
	mBase = m.M
	v3151 = m.ExcPending
	if v3151 != 0 {
		goto L14
	} else {
		goto L520
	}
L492:
	;
	v3059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+23)))
	if v3059&v3011&int32(1) != 0 {
		goto L491
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	v3066 = int32(4)
	v3067 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+14)))
	v3068 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693)+12)))
	v3069 = v3067 - v3068
	if v3069 <= v3066 {
		goto L497
	} else {
		goto L498
	}
L495:
	;
	goto L494
L496:
	;
	F_UnlockReleaseBuffer(m, v1670)
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L14
	} else {
		goto L515
	}
L497:
	;
	v3072 = v3066
	goto L499
L498:
	;
	v3072 = v3069
	goto L499
L499:
	;
	v3074 = v3072 - int32(4)
	if v3074 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v3129 = int32(0)
	goto L496
L501:
	;
	goto L502
L502:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v3068) {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	v3129 = v3074
	goto L496
L504:
	;
	v3085 = int32(base.Ui32(v3068+int32(_a_F_heap_vacuum_rel_14)) >> (uint(int32(2)) % 32))
	goto L506
L505:
	;
	v3085 = int32(0)
	goto L506
L506:
	;
	if base.Ui32(v3085&int32(_a_F_heap_vacuum_rel_15)) < base.Ui32(int32(291)) {
		goto L503
	} else {
		goto L507
	}
L507:
	;
	v3090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693)+10)))
	if v3090&int32(1) == int32(0) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3129 = int32(0)
	goto L496
L509:
	;
	goto L510
L510:
	;
	v3099 = int32(1)
	goto L511
L511:
	;
	v3108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1693+int32(20)+v3099&int32(_a_F_heap_vacuum_rel_15)<<(uint(int32(2))%32))+1)))
	if v3108&int32(384) == int32(0) {
		goto L503
	} else {
		goto L513
	}
L512:
	;
	v3129 = int32(0)
	goto L496
L513:
	;
	v3114 = v3099 + int32(1)
	v3115 = int32(_a_F_heap_vacuum_rel_15)
	if base.Ui32(v3114&v3115) <= base.Ui32(v3085&v3115) {
		v3099 = v3114
		goto L511
	} else {
		goto L514
	}
L514:
	;
	goto L512
L515:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_RecordPageWithFreeSpace(m, v3132, v1712, v3129)
	mBase = m.M
	v3134 = m.ExcPending
	if v3134 != 0 {
		goto L14
	} else {
		goto L516
	}
L516:
	;
	if v3018 == int32(0) {
		v1538 = v1712
		v1541 = v1667
		goto L225
	} else {
		goto L517
	}
L517:
	;
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v3138 = int32(0)
	if base.B2i32(base.B2i32(v3137 == v3138)&v3015 == v3138)|base.B2i32(base.Ui32(v1712-v1667) < base.Ui32(int32(_a_F_heap_vacuum_rel_26))) != 0 {
		v1538 = v1712
		v1541 = v1667
		goto L225
	} else {
		goto L518
	}
L518:
	;
	v3147 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_FreeSpaceMapVacuumRange(m, v3147, v1667, v1712)
	mBase = m.M
	v3149 = m.ExcPending
	if v3149 != 0 {
		goto L14
	} else {
		goto L519
	}
L519:
	;
	v1538 = v1712
	v1541 = v1712
	goto L225
L520:
	;
	v1538 = v1712
	v1541 = v1667
	goto L225
L521:
	;
	F_ReleaseBuffer(m, v3154)
	mBase = m.M
	v3156 = m.ExcPending
	if v3156 != 0 {
		goto L14
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	v3160 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3160 == int32(0) {
		goto L526
	} else {
		goto L527
	}
L524:
	;
	goto L523
L525:
	;
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(v173)+112))
	v3203 = *(*int64)(unsafe.Add(mBase, uint32(v173)+200))
	v3204 = base.F64_convert_i64_s(v3203)
	if base.Ui32(v3202) < base.Ui32(v1312) {
		goto L530
	} else {
		goto L531
	}
L526:
	;
	goto L525
L527:
	;
	v3164 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3164&int32(1) == int32(0) {
		goto L526
	} else {
		goto L528
	}
L528:
	;
	v3169 = int32(_a_F_heap_vacuum_rel_1)
	v3171 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3172 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3171 + v3172
	v3175 = *(*int32)(unsafe.Add(mBase, uint32(v3160)))
	*(*int32)(unsafe.Add(mBase, uint32(v3160))) = v3175 + v3172
	v3179 = int32(0)
	v3181 = int32(_a_F_heap_vacuum_rel_2)
	v3182 = base.AtomicRmwOr32(m, v3179, v3181, v3179)
	*(*int64)(unsafe.Add(mBase, uint32(v3160+int32(16))+232)) = v1321
	v3190 = base.AtomicRmwOr32(m, v3179, v3181, v3179)
	v3191 = *(*int32)(unsafe.Add(mBase, uint32(v3160)))
	*(*int32)(unsafe.Add(mBase, uint32(v3160))) = v3191 + v3172
	v3197 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3197 - v3172
	goto L526
L529:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v173)+160)) = v3254
	v3256 = float64(0)
	if base.F64_gt(v3254, v3256) != 0 {
		goto L548
	} else {
		goto L549
	}
L530:
	;
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3201)+48))
	v3210 = *(*float32)(unsafe.Add(mBase, uint32(v3209)+100))
	v3211 = base.F64_promote_f32(v3210)
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+96))
	if v1312 == v3212 {
		goto L534
	} else {
		goto L535
	}
L531:
	;
	v3249 = v3204
	goto L532
L532:
	;
	v3254 = v3249
	goto L529
L533:
	;
	v3225 = base.F64_convert_i32_u(v1312)
	if v3212 != 0 {
		goto L542
	} else {
		goto L543
	}
L534:
	;
	if base.Ui32(v3202) < base.Ui32(int32(2)) {
		goto L537
	} else {
		goto L538
	}
L535:
	;
	goto L536
L536:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v3202) {
		goto L533
	} else {
		goto L541
	}
L537:
	;
	v3254 = v3211
	goto L529
L538:
	;
	goto L539
L539:
	;
	if base.F64_lt(base.F64_convert_i32_u(v3202), base.F64_mul(base.F64_convert_i32_u(v1312), float64(0.02))) == int32(0) {
		goto L533
	} else {
		goto L540
	}
L540:
	;
	v3254 = v3211
	goto L529
L541:
	;
	v3254 = v3211
	goto L529
L542:
	;
	v3229 = base.F32_lt(v3210, float32(0))
	goto L544
L543:
	;
	v3229 = int32(1)
	goto L544
L544:
	;
	if v3229 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v3254 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v3204, base.F64_convert_i32_u(v3202)), v3225), float64(0.5)))
	goto L529
L546:
	;
	goto L547
L547:
	;
	v3249 = base.F64_floor(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(v3211, base.F64_convert_i32_u(v3212)), base.F64_sub(v3225, base.F64_convert_i32_u(v3202))), v3204), float64(0.5)))
	goto L532
L548:
	;
	v3259 = v3254
	goto L550
L549:
	;
	v3259 = v3256
	goto L550
L550:
	;
	v3260 = *(*int64)(unsafe.Add(mBase, uint32(v173)+208))
	v3263 = *(*int64)(unsafe.Add(mBase, uint32(v173)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v173)+152)) = base.F64_add(base.F64_add(v3259, base.F64_convert_i64_s(v3260)), base.F64_convert_i64_s(v3263))
	F_read_stream_end(m, v1530)
	mBase = m.M
	v3268 = m.ExcPending
	if v3268 != 0 {
		goto L14
	} else {
		goto L551
	}
L551:
	;
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v3270 = *(*int64)(unsafe.Add(mBase, uint32(v3269)+8))
	if int64(0) < v3270 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	F_lazy_vacuum(m, v173)
	mBase = m.M
	v3274 = m.ExcPending
	if v3274 != 0 {
		goto L14
	} else {
		goto L555
	}
L553:
	;
	goto L554
L554:
	;
	if base.Ui32(v1667) < base.Ui32(v1312) {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	goto L554
L556:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_FreeSpaceMapVacuumRange(m, v3276, v1667, v1312)
	mBase = m.M
	v3278 = m.ExcPending
	if v3278 != 0 {
		goto L14
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v3282 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3282 == int32(0) {
		goto L561
	} else {
		goto L562
	}
L559:
	;
	goto L558
L560:
	;
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v3323 <= int32(0) {
		goto L564
	} else {
		goto L565
	}
L561:
	;
	goto L560
L562:
	;
	v3286 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3286&int32(1) == int32(0) {
		goto L561
	} else {
		goto L563
	}
L563:
	;
	v3291 = int32(_a_F_heap_vacuum_rel_1)
	v3293 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3294 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3293 + v3294
	v3297 = *(*int32)(unsafe.Add(mBase, uint32(v3282)))
	*(*int32)(unsafe.Add(mBase, uint32(v3282))) = v3297 + v3294
	v3301 = int32(0)
	v3303 = int32(_a_F_heap_vacuum_rel_2)
	v3304 = base.AtomicRmwOr32(m, v3301, v3303, v3301)
	*(*int64)(unsafe.Add(mBase, uint32(v3282+int32(24))+232)) = v1321
	v3312 = base.AtomicRmwOr32(m, v3301, v3303, v3301)
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(v3282)))
	*(*int32)(unsafe.Add(mBase, uint32(v3282))) = v3313 + v3294
	v3319 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3319 - v3294
	goto L561
L564:
	;
	v3980 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	if v3980 != 0 {
		goto L617
	} else {
		goto L618
	}
L565:
	;
	v3326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+24)))
	if v3326 != int32(1) {
		goto L564
	} else {
		goto L566
	}
L566:
	;
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(v173)+108))
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v173)+112))
	v3331 = *(*float64)(unsafe.Add(mBase, uint32(v173)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1608)) = int64(34359738368)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1600)) = int64(38654705672)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+936)) = base.I64_extend_i32_u(v3323)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+928)) = int64(4)
	v3340 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1592)) = v3340
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1584)) = v3340
	goto L569
L567:
	;
	v3530 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	if v3530 == int32(0) {
		goto L585
	} else {
		goto L586
	}
L568:
	;
	goto L567
L569:
	;
	v3358 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3358 == int32(0) {
		goto L568
	} else {
		goto L570
	}
L570:
	;
	v3362 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3362&int32(1) == int32(0) {
		goto L568
	} else {
		goto L571
	}
L571:
	;
	v3367 = int32(_a_F_heap_vacuum_rel_1)
	v3369 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3370 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3369 + v3370
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(v3358)))
	*(*int32)(unsafe.Add(mBase, uint32(v3358))) = v3373 + v3370
	v3377 = int32(0)
	v3380 = base.AtomicRmwOr32(m, v3377, int32(_a_F_heap_vacuum_rel_2), v3377)
	goto L573
L572:
	;
	v3507 = int32(0)
	v3510 = base.AtomicRmwOr32(m, v3507, int32(_a_F_heap_vacuum_rel_2), v3507)
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v3358)))
	v3512 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3358))) = v3511 + v3512
	v3515 = int32(_a_F_heap_vacuum_rel_1)
	v3517 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3517 - v3512
	goto L568
L573:
	;
	goto L575
L575:
	;
	goto L576
L576:
	;
	v3472 = int32(0)
	v3475 = int32(0)
	goto L581
L581:
	;
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1608)+v3475<<(uint(int32(2))%32))))
	v3485 = int32(3)
	v3491 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(928)+v3475<<(uint(v3485)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v3358+int32(232)+v3484<<(uint(v3485)%32)))) = v3491
	v3493 = int32(1)
	v3496 = v3472 + v3493
	if v3496 != int32(2) {
		v3472 = v3496
		v3475 = v3475 + v3493
		goto L581
	} else {
		goto L583
	}
L582:
	;
	goto L572
L583:
	;
	goto L582
L584:
	;
	goto L602
L585:
	;
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v3533 <= int32(0) {
		goto L584
	} else {
		goto L588
	}
L586:
	;
	goto L587
L587:
	;
	v3685 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	v3686 = *(*int32)(unsafe.Add(mBase, uint32(v3530)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v3686)+16)) = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(v3331))
	v3690 = *(*int32)(unsafe.Add(mBase, uint32(v3530)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v3690)+24)) = uint8(base.B2i32(base.Ui32(v3330) < base.Ui32(v3329)))
	F_parallel_vacuum_process_all_indexes(m, v3530, v3685, int32(0))
	mBase = m.M
	v3695 = m.ExcPending
	if v3695 != 0 {
		goto L14
	} else {
		goto L599
	}
L588:
	;
	v3576 = int64(0)
	goto L589
L589:
	;
	v3589 = base.I32_wrap_i64(v3576) << (uint(int32(2)) % 32)
	v3590 = *(*int32)(unsafe.Add(mBase, uint32(v173)+168))
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(v3589+v3590)))
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v3595 = *(*int32)(unsafe.Add(mBase, uint32(v3593+v3589)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+960)) = v3595
	v3597 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	*(*float64)(unsafe.Add(mBase, uint32(v52)+976)) = v3331
	*(*int32)(unsafe.Add(mBase, uint32(v52)+972)) = int32(13)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+970)) = uint8(base.B2i32(base.Ui32(v3330) < base.Ui32(v3329)))
	v3602 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+968)) = uint16(v3602)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+964)) = v3597
	v3605 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+984)) = v3605
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(v3595)+48))
	v3610 = F_pstrdup(m, v3607+int32(4))
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L14
	} else {
		goto L591
	}
L590:
	;
	goto L584
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+80)) = v3610
	v3613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+88)))
	v3614 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+88)) = uint16(v3614)
	v3616 = *(*int32)(unsafe.Add(mBase, uint32(v173)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+84)) = int32(-1)
	v3619 = *(*int32)(unsafe.Add(mBase, uint32(v173)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+92)) = int32(4)
	v3624 = F_vac_cleanup_one_index(m, v52+int32(960), v3592)
	mBase = m.M
	v3625 = m.ExcPending
	if v3625 != 0 {
		goto L14
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+92)) = v3619
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+88)) = uint16(v3613)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+84)) = v3616
	v3629 = *(*int32)(unsafe.Add(mBase, uint32(v173)+80))
	F_pfree(m, v3629)
	mBase = m.M
	v3631 = m.ExcPending
	if v3631 != 0 {
		goto L14
	} else {
		goto L593
	}
L593:
	;
	v3632 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+80)) = v3632
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v173)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v3634+v3589))) = v3624
	v3639 = v3576 + int64(1)
	v3642 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3642 == v3632 {
		goto L595
	} else {
		goto L596
	}
L594:
	;
	v3683 = int64(*(*int32)(unsafe.Add(mBase, uint32(v173)+8)))
	if v3639 < v3683 {
		v3576 = v3639
		goto L589
	} else {
		goto L598
	}
L595:
	;
	goto L594
L596:
	;
	v3646 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3646&int32(1) == int32(0) {
		goto L595
	} else {
		goto L597
	}
L597:
	;
	v3651 = int32(_a_F_heap_vacuum_rel_1)
	v3653 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3654 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3653 + v3654
	v3657 = *(*int32)(unsafe.Add(mBase, uint32(v3642)))
	*(*int32)(unsafe.Add(mBase, uint32(v3642))) = v3657 + v3654
	v3661 = int32(0)
	v3663 = int32(_a_F_heap_vacuum_rel_2)
	v3664 = base.AtomicRmwOr32(m, v3661, v3663, v3661)
	*(*int64)(unsafe.Add(mBase, uint32(v3642+int32(72))+232)) = v3639
	v3672 = base.AtomicRmwOr32(m, v3661, v3663, v3661)
	v3673 = *(*int32)(unsafe.Add(mBase, uint32(v3642)))
	*(*int32)(unsafe.Add(mBase, uint32(v3642))) = v3673 + v3654
	v3679 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3679 - v3654
	goto L595
L598:
	;
	goto L590
L599:
	;
	goto L584
L600:
	;
	goto L564
L601:
	;
	goto L600
L602:
	;
	v3759 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3759 == int32(0) {
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v3763 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3763&int32(1) == int32(0) {
		goto L601
	} else {
		goto L604
	}
L604:
	;
	v3768 = int32(_a_F_heap_vacuum_rel_1)
	v3770 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3771 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3770 + v3771
	v3774 = *(*int32)(unsafe.Add(mBase, uint32(v3759)))
	*(*int32)(unsafe.Add(mBase, uint32(v3759))) = v3774 + v3771
	v3778 = int32(0)
	v3781 = base.AtomicRmwOr32(m, v3778, int32(_a_F_heap_vacuum_rel_2), v3778)
	goto L606
L605:
	;
	v3908 = int32(0)
	v3911 = base.AtomicRmwOr32(m, v3908, int32(_a_F_heap_vacuum_rel_2), v3908)
	v3912 = *(*int32)(unsafe.Add(mBase, uint32(v3759)))
	v3913 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3759))) = v3912 + v3913
	v3916 = int32(_a_F_heap_vacuum_rel_1)
	v3918 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3918 - v3913
	goto L601
L606:
	;
	goto L608
L608:
	;
	goto L609
L609:
	;
	v3873 = int32(0)
	v3876 = int32(0)
	goto L614
L614:
	;
	v3885 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1600)+v3876<<(uint(int32(2))%32))))
	v3886 = int32(3)
	v3892 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(1584)+v3876<<(uint(v3886)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v3759+int32(232)+v3885<<(uint(v3886)%32)))) = v3892
	v3894 = int32(1)
	v3897 = v3873 + v3894
	if v3897 != int32(2) {
		v3873 = v3897
		v3876 = v3876 + v3894
		goto L614
	} else {
		goto L616
	}
L615:
	;
	goto L605
L616:
	;
	goto L615
L617:
	;
	v3981 = *(*int32)(unsafe.Add(mBase, uint32(v173)+168))
	v3982 = int32(0)
	v3983 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+12))
	if v3982 < v3983 {
		goto L620
	} else {
		goto L621
	}
L618:
	;
	goto L619
L619:
	;
	v4187 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v4188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+24)))
	v4191 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if base.B2i32(v4188 != int32(1))|base.B2i32(v4191 <= int32(0)) != 0 {
		goto L636
	} else {
		goto L637
	}
L620:
	;
	v3994 = v3982
	goto L623
L621:
	;
	goto L622
L622:
	;
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+24))
	F_TidStoreDestroy(m, v4118)
	mBase = m.M
	v4120 = m.ExcPending
	if v4120 != 0 {
		goto L14
	} else {
		goto L631
	}
L623:
	;
	v4035 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+20))
	v4038 = v4035 + v3994*int32(48)
	v4039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4038)+5)))
	if v4039 == int32(1) {
		goto L626
	} else {
		goto L627
	}
L624:
	;
	goto L622
L625:
	;
	v4066 = v3994 + int32(1)
	v4067 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+12))
	if v4066 < v4067 {
		v3994 = v4066
		goto L623
	} else {
		goto L630
	}
L626:
	;
	v4046 = F_palloc0(m, int32(40))
	mBase = m.M
	v4047 = m.ExcPending
	if v4047 != 0 {
		goto L14
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3981+v3994<<(uint(int32(2))%32)))) = int32(0)
	goto L625
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3981+v3994<<(uint(int32(2))%32)))) = v4046
	v4049 = *(*int64)(unsafe.Add(mBase, uint32(v4038)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4046)+32)) = v4049
	v4051 = *(*int64)(unsafe.Add(mBase, uint32(v4038)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4046)+24)) = v4051
	v4053 = *(*int64)(unsafe.Add(mBase, uint32(v4038)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4046)+16)) = v4053
	v4055 = *(*int64)(unsafe.Add(mBase, uint32(v4038)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4046)+8)) = v4055
	v4057 = *(*int64)(unsafe.Add(mBase, uint32(v4038)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4046))) = v4057
	goto L625
L630:
	;
	goto L624
L631:
	;
	v4121 = *(*int32)(unsafe.Add(mBase, uint32(v3980)))
	F_DestroyParallelContext(m, v4121)
	mBase = m.M
	v4123 = m.ExcPending
	if v4123 != 0 {
		goto L14
	} else {
		goto L632
	}
L632:
	;
	v4126 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[22]))
	v4127 = *(*int32)(unsafe.Add(mBase, uint32(v4126)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v4126)+72)) = v4127 - int32(1)
	goto L633
L633:
	;
	v4131 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+36))
	F_pfree(m, v4131)
	mBase = m.M
	v4133 = m.ExcPending
	if v4133 != 0 {
		goto L14
	} else {
		goto L634
	}
L634:
	;
	F_pfree(m, v3980)
	mBase = m.M
	v4135 = m.ExcPending
	if v4135 != 0 {
		goto L14
	} else {
		goto L635
	}
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = int32(0)
	goto L619
L636:
	;
	v4278 = v4187
	v4321 = v4191
	goto L638
L637:
	;
	v4195 = *(*int32)(unsafe.Add(mBase, uint32(v173)+168))
	v4199 = int32(0)
	goto L639
L638:
	;
	F_vac_close_indexes(m, v4321, v4278, int32(0))
	mBase = m.M
	v4324 = m.ExcPending
	if v4324 != 0 {
		goto L14
	} else {
		goto L646
	}
L639:
	;
	v4247 = v4199 << (uint(int32(2)) % 32)
	v4249 = *(*int32)(unsafe.Add(mBase, uint32(v4195+v4247)))
	if v4249 == int32(0) {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	v4270 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v4271 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v4278 = v4270
	v4321 = v4271
	goto L638
L641:
	;
	v4268 = v4199 + int32(1)
	if v4268 != v4191 {
		v4199 = v4268
		goto L639
	} else {
		goto L645
	}
L642:
	;
	v4252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4249)+4)))
	if v4252 != 0 {
		goto L641
	} else {
		goto L643
	}
L643:
	;
	v4254 = *(*int32)(unsafe.Add(mBase, uint32(v4247+v4187)))
	v4255 = *(*int32)(unsafe.Add(mBase, uint32(v4249)))
	v4256 = *(*float64)(unsafe.Add(mBase, uint32(v4249)+8))
	v4257 = int32(0)
	F_vac_update_relstats(m, v4254, v4255, v4256, v4257, v4257, v4257, v4257, v4257, v4257, v4257, v4257)
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L14
	} else {
		goto L644
	}
L644:
	;
	goto L641
L645:
	;
	goto L640
L646:
	;
	v4325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+25)))
	if v4325 != int32(1) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v5287 = *(*int32)(unsafe.Add(mBase, uint32(v52)+564))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12])) = v5287
	v5293 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v5293 == int32(0) {
		goto L789
	} else {
		goto L790
	}
L648:
	;
	v4329 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])))
	if v4329&int32(1) != 0 {
		goto L647
	} else {
		goto L649
	}
L649:
	;
	v4332 = *(*int32)(unsafe.Add(mBase, uint32(v173)+108))
	v4333 = *(*int32)(unsafe.Add(mBase, uint32(v173)+148))
	if v4332 == v4333 {
		goto L647
	} else {
		goto L650
	}
L650:
	;
	v4335 = v4332 - v4333
	if base.B2i32(base.Ui32(v4335) <= base.Ui32(int32(999)))&base.B2i32(base.Ui32(v4335) < base.Ui32(int32(base.Ui32(v4332)>>(uint(int32(4))%32)))) != 0 {
		goto L647
	} else {
		goto L651
	}
L651:
	;
	v4346 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v4346 == int32(0) {
		goto L653
	} else {
		goto L654
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+92)) = int32(5)
	v4389 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+88)) = uint16(v4389)
	v4391 = *(*int32)(unsafe.Add(mBase, uint32(v173)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+84)) = v4391
	v4400 = v4332
	goto L656
L653:
	;
	goto L652
L654:
	;
	v4350 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v4350&int32(1) == int32(0) {
		goto L653
	} else {
		goto L655
	}
L655:
	;
	v4355 = int32(_a_F_heap_vacuum_rel_1)
	v4357 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v4358 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v4357 + v4358
	v4361 = *(*int32)(unsafe.Add(mBase, uint32(v4346)))
	*(*int32)(unsafe.Add(mBase, uint32(v4346))) = v4361 + v4358
	v4365 = int32(0)
	v4367 = int32(_a_F_heap_vacuum_rel_2)
	v4368 = base.AtomicRmwOr32(m, v4365, v4367, v4365)
	*(*int64)(unsafe.Add(mBase, uint32(v4346+v4365)+232)) = int64(5)
	v4376 = base.AtomicRmwOr32(m, v4365, v4367, v4365)
	v4377 = *(*int32)(unsafe.Add(mBase, uint32(v4346)))
	*(*int32)(unsafe.Add(mBase, uint32(v4346))) = v4377 + v4358
	v4383 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v4383 - v4358
	goto L653
L656:
	;
	v4443 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v4444 = F_ConditionalLockRelation(m, v4443)
	mBase = m.M
	v4445 = m.ExcPending
	if v4445 != 0 {
		goto L14
	} else {
		goto L658
	}
L657:
	;
	goto L647
L658:
	;
	if v4444 == int32(0) {
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v4450 = int32(0)
	goto L662
L660:
	;
	goto L661
L661:
	;
	v4595 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v4597 = F_RelationGetNumberOfBlocksInFork(m, v4595, int32(0))
	mBase = m.M
	v4598 = m.ExcPending
	if v4598 != 0 {
		goto L14
	} else {
		goto L682
	}
L662:
	;
	v4498 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29]))
	if v4498 != 0 {
		goto L664
	} else {
		goto L665
	}
L663:
	;
	goto L661
L664:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4500 = m.ExcPending
	if v4500 != 0 {
		goto L14
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	if v4450 == int32(100) {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	goto L666
L668:
	;
	v4505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+96)))
	if v4505 != 0 {
		goto L671
	} else {
		goto L672
	}
L669:
	;
	goto L670
L670:
	;
	v4525 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[30]))
	v4529 = F_WaitLatch(m, v4525, int32(41), int32(50), int32(150994952))
	mBase = m.M
	v4530 = m.ExcPending
	if v4530 != 0 {
		goto L14
	} else {
		goto L678
	}
L671:
	;
	v4506 = int32(17)
	goto L673
L672:
	;
	v4506 = int32(13)
	goto L673
L673:
	;
	v4508 = F_errstart(m, v4506, int32(0))
	mBase = m.M
	v4509 = m.ExcPending
	if v4509 != 0 {
		goto L14
	} else {
		goto L674
	}
L674:
	;
	if v4508 == int32(0) {
		goto L647
	} else {
		goto L675
	}
L675:
	;
	v4512 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+432)) = v4512
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_27), v52+int32(432))
	mBase = m.M
	v4518 = m.ExcPending
	if v4518 != 0 {
		goto L14
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3249), int32(_a_F_heap_vacuum_rel_28))
	mBase = m.M
	v4523 = m.ExcPending
	if v4523 != 0 {
		goto L14
	} else {
		goto L677
	}
L677:
	;
	goto L647
L678:
	;
	v4532 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[30]))
	v4533 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4532))) = v4533
	v4538 = base.AtomicRmwOr32(m, v4533, int32(_a_F_heap_vacuum_rel_29), v4533)
	goto L679
L679:
	;
	v4541 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v4542 = F_ConditionalLockRelation(m, v4541)
	mBase = m.M
	v4543 = m.ExcPending
	if v4543 != 0 {
		goto L14
	} else {
		goto L680
	}
L680:
	;
	if v4542 == int32(0) {
		v4450 = v4450 + int32(1)
		goto L662
	} else {
		goto L681
	}
L681:
	;
	goto L663
L682:
	;
	if v4597 != v4400 {
		goto L683
	} else {
		goto L684
	}
L683:
	;
	v4600 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_UnlockRelation(m, v4600)
	mBase = m.M
	v4602 = m.ExcPending
	if v4602 != 0 {
		goto L14
	} else {
		goto L686
	}
L684:
	;
	goto L685
L685:
	;
	F___clock_gettime(m, int32(1), v52+int32(960))
	mBase = m.M
	v4607 = int32(0)
	v4608 = *(*int32)(unsafe.Add(mBase, uint32(v173)+108))
	v4609 = *(*int32)(unsafe.Add(mBase, uint32(v173)+148))
	if base.Ui32(v4608) <= base.Ui32(v4609) {
		v5156 = v4609
		v5169 = v4607
		goto L687
	} else {
		goto L688
	}
L686:
	;
	goto L647
L687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+84)) = v5156
	v5198 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if base.Ui32(v4400) <= base.Ui32(v5156) {
		goto L772
	} else {
		goto L773
	}
L688:
	;
	v4611 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+968)))
	v4612 = *(*int64)(unsafe.Add(mBase, uint32(v52)+960))
	v4622 = v4608
	v4630 = int32(-1)
	v4655 = v4611 + v4612*int64(1000000000)
	goto L689
L689:
	;
	if v4622&int32(31) != 0 {
		v4863 = v4655
		goto L691
	} else {
		goto L692
	}
L690:
	;
	v5156 = v5146
	v5169 = v4607
	goto L687
L691:
	;
	v4866 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29]))
	if v4866 != 0 {
		goto L738
	} else {
		goto L739
	}
L692:
	;
	F___clock_gettime(m, int32(1), v52+int32(960))
	mBase = m.M
	v4672 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+968)))
	v4673 = *(*int64)(unsafe.Add(mBase, uint32(v52)+960))
	v4676 = v4672 + v4673*int64(1000000000)
	if v4676-v4655 < int64(20000000) {
		v4863 = v4655
		goto L691
	} else {
		goto L693
	}
L693:
	;
	v4680 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v4681 = m.G0
	v4683 = v4681 - int32(16)
	m.G0 = v4683
	v4685 = *(*int32)(unsafe.Add(mBase, uint32(v4680)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v4683))) = v4685
	v4687 = *(*int32)(unsafe.Add(mBase, uint32(v4680)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v4683)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v4683)+4)) = v4687
	v4691 = m.G0
	v4693 = v4691 - int32(80)
	m.G0 = v4693
	v4695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4683)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v4695-int32(3))&int32(255)) {
		goto L696
	} else {
		goto L697
	}
L694:
	;
	m.G0 = v4683 + int32(16)
	if v4796 == int32(0) {
		v4863 = v4676
		goto L691
	} else {
		goto L730
	}
L695:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L14
	} else {
		goto L727
	}
L696:
	;
	v4704 = *(*int32)(unsafe.Add(mBase, uint32(v4695<<(uint(int32(2))%32))+uint32(_c_F_heap_vacuum_rel[31])))
	v4705 = *(*int32)(unsafe.Add(mBase, uint32(v4704)))
	if v4705 < int32(8) {
		goto L695
	} else {
		goto L699
	}
L697:
	;
	goto L698
L698:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4804 = m.ExcPending
	if v4804 != 0 {
		goto L14
	} else {
		goto L724
	}
L699:
	;
	v4708 = *(*int64)(unsafe.Add(mBase, uint32(v4683)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4693)+64)) = v4708
	v4710 = *(*int64)(unsafe.Add(mBase, uint32(v4683)))
	*(*int64)(unsafe.Add(mBase, uint32(v4693)+56)) = v4710
	*(*int32)(unsafe.Add(mBase, uint32(v4693)+72)) = int32(8)
	v4714 = int32(0)
	v4716 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[32]))
	v4721 = F_hash_search(m, v4716, v4693+int32(56), v4714, v4714)
	mBase = m.M
	v4722 = m.ExcPending
	if v4722 != 0 {
		goto L14
	} else {
		goto L702
	}
L700:
	;
	m.G0 = v4693 + int32(80)
	goto L694
L701:
	;
	v4746 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[33]))
	v4747 = *(*int32)(unsafe.Add(mBase, uint32(v4721)+20))
	v4754 = v4746 + v4747&int32(15)<<(uint(int32(7))%32) + int32(_a_F_heap_vacuum_rel_30)
	v4756 = F_LWLockAcquire(m, v4754, int32(1))
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L14
	} else {
		goto L711
	}
L702:
	;
	if v4721 != 0 {
		goto L703
	} else {
		goto L704
	}
L703:
	;
	v4723 = *(*int64)(unsafe.Add(mBase, uint32(v4721)+32))
	if int64(0) < v4723 {
		goto L701
	} else {
		goto L706
	}
L704:
	;
	goto L705
L705:
	;
	v4728 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4729 = m.ExcPending
	if v4729 != 0 {
		goto L14
	} else {
		goto L707
	}
L706:
	;
	goto L705
L707:
	;
	if v4728 == int32(0) {
		v4796 = v4714
		goto L700
	} else {
		goto L708
	}
L708:
	;
	v4732 = *(*int32)(unsafe.Add(mBase, uint32(v4704)+8))
	v4733 = *(*int32)(unsafe.Add(mBase, uint32(v4732)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4693)+32)) = v4733
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_31), v4693+int32(32))
	mBase = m.M
	v4739 = m.ExcPending
	if v4739 != 0 {
		goto L14
	} else {
		goto L709
	}
L709:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_32), int32(736), int32(_a_F_heap_vacuum_rel_33))
	mBase = m.M
	v4744 = m.ExcPending
	if v4744 != 0 {
		goto L14
	} else {
		goto L710
	}
L710:
	;
	v4796 = v4714
	goto L700
L711:
	;
	v4758 = *(*int32)(unsafe.Add(mBase, uint32(v4721)+28))
	v4759 = *(*int32)(unsafe.Add(mBase, uint32(v4758)+12))
	if int32(base.Ui32(v4759)>>(uint(int32(8))%32))&int32(1) == int32(0) {
		goto L712
	} else {
		goto L713
	}
L712:
	;
	F_LWLockRelease(m, v4754)
	mBase = m.M
	v4767 = m.ExcPending
	if v4767 != 0 {
		goto L14
	} else {
		goto L715
	}
L713:
	;
	goto L714
L714:
	;
	v4787 = *(*int32)(unsafe.Add(mBase, uint32(v4704)+4))
	v4788 = *(*int32)(unsafe.Add(mBase, uint32(v4787)+32))
	v4789 = *(*int32)(unsafe.Add(mBase, uint32(v4721)+24))
	v4790 = *(*int32)(unsafe.Add(mBase, uint32(v4789)+20))
	F_LWLockRelease(m, v4754)
	mBase = m.M
	v4792 = m.ExcPending
	if v4792 != 0 {
		goto L14
	} else {
		goto L723
	}
L715:
	;
	v4770 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4771 = m.ExcPending
	if v4771 != 0 {
		goto L14
	} else {
		goto L716
	}
L716:
	;
	if v4770 != 0 {
		goto L717
	} else {
		goto L718
	}
L717:
	;
	v4772 = *(*int32)(unsafe.Add(mBase, uint32(v4704)+8))
	v4773 = *(*int32)(unsafe.Add(mBase, uint32(v4772)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4693)+48)) = v4773
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_31), v4693+int32(48))
	mBase = m.M
	v4779 = m.ExcPending
	if v4779 != 0 {
		goto L14
	} else {
		goto L720
	}
L718:
	;
	goto L719
L719:
	;
	F_RemoveLocalLock(m, v4721)
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		goto L14
	} else {
		goto L722
	}
L720:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_32), int32(766), int32(_a_F_heap_vacuum_rel_33))
	mBase = m.M
	v4784 = m.ExcPending
	if v4784 != 0 {
		goto L14
	} else {
		goto L721
	}
L721:
	;
	goto L719
L722:
	;
	v4796 = v4714
	goto L700
L723:
	;
	v4796 = base.B2i32(v4790&v4788 != int32(0))
	goto L700
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4693))) = v4695
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_34), v4693)
	mBase = m.M
	v4808 = m.ExcPending
	if v4808 != 0 {
		goto L14
	} else {
		goto L725
	}
L725:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_32), int32(707), int32(_a_F_heap_vacuum_rel_33))
	mBase = m.M
	v4813 = m.ExcPending
	if v4813 != 0 {
		goto L14
	} else {
		goto L726
	}
L726:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4693)+16)) = int32(8)
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_35), v4693+int32(16))
	mBase = m.M
	v4824 = m.ExcPending
	if v4824 != 0 {
		goto L14
	} else {
		goto L728
	}
L728:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_32), int32(710), int32(_a_F_heap_vacuum_rel_33))
	mBase = m.M
	v4829 = m.ExcPending
	if v4829 != 0 {
		goto L14
	} else {
		goto L729
	}
L729:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L730:
	;
	v4835 = int32(1)
	v4838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+96)))
	if v4838 != 0 {
		goto L731
	} else {
		goto L732
	}
L731:
	;
	v4839 = int32(17)
	goto L733
L732:
	;
	v4839 = int32(13)
	goto L733
L733:
	;
	v4841 = F_errstart(m, v4839, int32(0))
	mBase = m.M
	v4842 = m.ExcPending
	if v4842 != 0 {
		goto L14
	} else {
		goto L734
	}
L734:
	;
	if v4841 == int32(0) {
		v5156 = v4622
		v5169 = v4835
		goto L687
	} else {
		goto L735
	}
L735:
	;
	v4845 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+416)) = v4845
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_36), v52+int32(416))
	mBase = m.M
	v4851 = m.ExcPending
	if v4851 != 0 {
		goto L14
	} else {
		goto L736
	}
L736:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3381), int32(_a_F_heap_vacuum_rel_37))
	mBase = m.M
	v4856 = m.ExcPending
	if v4856 != 0 {
		goto L14
	} else {
		goto L737
	}
L737:
	;
	v5156 = v4622
	v5169 = v4835
	goto L687
L738:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4868 = m.ExcPending
	if v4868 != 0 {
		goto L14
	} else {
		goto L741
	}
L739:
	;
	goto L740
L740:
	;
	v4870 = v4622 - int32(1)
	if base.Ui32(v4870) < base.Ui32(v4630) {
		goto L742
	} else {
		goto L743
	}
L741:
	;
	goto L740
L742:
	;
	v4873 = v4870 & int32(-32)
	v4876 = v4873
	goto L745
L743:
	;
	v4949 = v4630
	goto L744
L744:
	;
	v4985 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v4986 = int32(0)
	v4988 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v4989 = F_ReadBufferExtended(m, v4985, v4986, v4870, v4986, v4988)
	mBase = m.M
	v4990 = m.ExcPending
	if v4990 != 0 {
		goto L14
	} else {
		goto L753
	}
L745:
	;
	v4925 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_PrefetchBuffer(m, v52+int32(960), v4925, int32(0), v4876)
	mBase = m.M
	v4928 = m.ExcPending
	if v4928 != 0 {
		goto L14
	} else {
		goto L747
	}
L746:
	;
	v4949 = v4873
	goto L744
L747:
	;
	v4930 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29]))
	if v4930 != 0 {
		goto L748
	} else {
		goto L749
	}
L748:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4932 = m.ExcPending
	if v4932 != 0 {
		goto L14
	} else {
		goto L751
	}
L749:
	;
	goto L750
L750:
	;
	if base.Ui32(v4876) < base.Ui32(v4870) {
		v4876 = v4876 + int32(1)
		goto L745
	} else {
		goto L752
	}
L751:
	;
	goto L750
L752:
	;
	goto L746
L753:
	;
	F_LockBuffer(m, v4989, int32(1))
	mBase = m.M
	v4993 = m.ExcPending
	if v4993 != 0 {
		goto L14
	} else {
		goto L754
	}
L754:
	;
	if v4989 < int32(0) {
		goto L757
	} else {
		goto L758
	}
L755:
	;
	F_UnlockReleaseBuffer(m, v4989)
	mBase = m.M
	v5145 = m.ExcPending
	if v5145 != 0 {
		goto L14
	} else {
		goto L770
	}
L756:
	;
	v5012 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5011)+14)))
	if v5012 == int32(0) {
		goto L755
	} else {
		goto L760
	}
L757:
	;
	v4997 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[24]))
	v5003 = *(*int32)(unsafe.Add(mBase, uint32(v4997+(v4989^int32(-1))<<(uint(int32(2))%32))))
	v5011 = v5003
	goto L756
L758:
	;
	goto L759
L759:
	;
	v5005 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[25]))
	v5011 = v5005 + v4989<<(uint(int32(13))%32) + int32(-8192)
	goto L756
L760:
	;
	v5015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5011)+12)))
	if base.Ui32(v5015) < base.Ui32(int32(25)) {
		goto L755
	} else {
		goto L761
	}
L761:
	;
	v5023 = int32(base.Ui32(v5015+int32(_a_F_heap_vacuum_rel_14))>>(uint(int32(2))%32)) & int32(_a_F_heap_vacuum_rel_15)
	if v5023 == int32(0) {
		goto L755
	} else {
		goto L762
	}
L762:
	;
	v5031 = int32(1)
	goto L763
L763:
	;
	v5083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5011+int32(20)+v5031&int32(_a_F_heap_vacuum_rel_15)<<(uint(int32(2))%32))+1)))
	if v5083&int32(384) == int32(0) {
		goto L765
	} else {
		goto L766
	}
L764:
	;
	F_UnlockReleaseBuffer(m, v4989)
	mBase = m.M
	v5094 = m.ExcPending
	if v5094 != 0 {
		goto L14
	} else {
		goto L769
	}
L765:
	;
	v5089 = v5031 + int32(1)
	if base.Ui32(v5089&int32(_a_F_heap_vacuum_rel_15)) <= base.Ui32(v5023) {
		v5031 = v5089
		goto L763
	} else {
		goto L768
	}
L766:
	;
	goto L767
L767:
	;
	goto L764
L768:
	;
	goto L755
L769:
	;
	v5156 = v4622
	v5169 = v4607
	goto L687
L770:
	;
	v5146 = *(*int32)(unsafe.Add(mBase, uint32(v173)+148))
	if base.Ui32(v5146) < base.Ui32(v4870) {
		v4622 = v4870
		v4630 = v4949
		v4655 = v4863
		goto L689
	} else {
		goto L771
	}
L771:
	;
	goto L690
L772:
	;
	F_UnlockRelation(m, v5198)
	mBase = m.M
	v5201 = m.ExcPending
	if v5201 != 0 {
		goto L14
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	F_RelationTruncate(m, v5198, v5156)
	mBase = m.M
	v5203 = m.ExcPending
	if v5203 != 0 {
		goto L14
	} else {
		goto L776
	}
L775:
	;
	goto L647
L776:
	;
	v5204 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	F_UnlockRelation(m, v5204)
	mBase = m.M
	v5206 = m.ExcPending
	if v5206 != 0 {
		goto L14
	} else {
		goto L777
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+108)) = v5156
	v5208 = *(*int32)(unsafe.Add(mBase, uint32(v173)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+120)) = v5208 + (v4400 - v5156)
	v5214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+96)))
	if v5214 != 0 {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v5215 = int32(17)
	goto L780
L779:
	;
	v5215 = int32(13)
	goto L780
L780:
	;
	v5217 = F_errstart(m, v5215, int32(0))
	mBase = m.M
	v5218 = m.ExcPending
	if v5218 != 0 {
		goto L14
	} else {
		goto L781
	}
L781:
	;
	if v5217 != 0 {
		goto L782
	} else {
		goto L783
	}
L782:
	;
	v5219 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+408)) = v5156
	*(*int32)(unsafe.Add(mBase, uint32(v52)+404)) = v4400
	*(*int32)(unsafe.Add(mBase, uint32(v52)+400)) = v5219
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_38), v52+int32(400))
	mBase = m.M
	v5227 = m.ExcPending
	if v5227 != 0 {
		goto L14
	} else {
		goto L785
	}
L783:
	;
	goto L784
L784:
	;
	v5234 = *(*int32)(unsafe.Add(mBase, uint32(v173)+148))
	if v5169&base.B2i32(base.Ui32(v5234) < base.Ui32(v5156)) != 0 {
		v4400 = v5156
		goto L656
	} else {
		goto L787
	}
L785:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3320), int32(_a_F_heap_vacuum_rel_28))
	mBase = m.M
	v5232 = m.ExcPending
	if v5232 != 0 {
		goto L14
	} else {
		goto L786
	}
L786:
	;
	goto L784
L787:
	;
	goto L657
L788:
	;
	v5334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+64)))
	if v5334 == int32(1) {
		goto L792
	} else {
		goto L793
	}
L789:
	;
	goto L788
L790:
	;
	v5297 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v5297&int32(1) == int32(0) {
		goto L789
	} else {
		goto L791
	}
L791:
	;
	v5302 = int32(_a_F_heap_vacuum_rel_1)
	v5304 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v5305 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5304 + v5305
	v5308 = *(*int32)(unsafe.Add(mBase, uint32(v5293)))
	*(*int32)(unsafe.Add(mBase, uint32(v5293))) = v5308 + v5305
	v5312 = int32(0)
	v5314 = int32(_a_F_heap_vacuum_rel_2)
	v5315 = base.AtomicRmwOr32(m, v5312, v5314, v5312)
	*(*int64)(unsafe.Add(mBase, uint32(v5293+v5312)+232)) = int64(6)
	v5323 = base.AtomicRmwOr32(m, v5312, v5314, v5312)
	v5324 = *(*int32)(unsafe.Add(mBase, uint32(v5293)))
	*(*int32)(unsafe.Add(mBase, uint32(v5293))) = v5324 + v5305
	v5330 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5330 - v5305
	goto L789
L792:
	;
	v5337 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1303))) = v5337
	*(*int32)(unsafe.Add(mBase, uint32(v1301))) = v5337
	goto L794
L793:
	;
	goto L794
L794:
	;
	v5341 = *(*int32)(unsafe.Add(mBase, uint32(v173)+108))
	F_visibilitymap_count(m, l0, v52+int32(1584), v52+int32(912))
	mBase = m.M
	v5347 = m.ExcPending
	if v5347 != 0 {
		goto L14
	} else {
		goto L795
	}
L795:
	;
	v5348 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1584))
	if base.Ui32(v5341) < base.Ui32(v5348) {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1584)) = v5341
	v5351 = v5341
	goto L798
L797:
	;
	v5351 = v5348
	goto L798
L798:
	;
	v5352 = *(*int32)(unsafe.Add(mBase, uint32(v52)+912))
	if base.Ui32(v5351) < base.Ui32(v5352) {
		goto L799
	} else {
		goto L800
	}
L799:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+912)) = v5351
	v5355 = v5351
	goto L801
L800:
	;
	v5355 = v5352
	goto L801
L801:
	;
	v5356 = int32(0)
	v5357 = *(*float64)(unsafe.Add(mBase, uint32(v173)+160))
	v5358 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v5361 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
	v5362 = *(*int32)(unsafe.Add(mBase, uint32(v173)+60))
	F_vac_update_relstats(m, l0, v5341, v5357, v5351, v5355, base.B2i32(v5356 < v5358), v5361, v5362, v52+int32(924), v52+int32(908), v5356)
	mBase = m.M
	v5369 = m.ExcPending
	if v5369 != 0 {
		goto L14
	} else {
		goto L802
	}
L802:
	;
	v5370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5371)+117)))
	v5373 = *(*float64)(unsafe.Add(mBase, uint32(v173)+160))
	v5374 = float64(0)
	if base.F64_gt(v5373, v5374) != 0 {
		goto L803
	} else {
		goto L804
	}
L803:
	;
	v5377 = v5373
	goto L805
L804:
	;
	v5377 = v5374
	goto L805
L805:
	;
	v5379 = *(*int64)(unsafe.Add(mBase, uint32(v173)+216))
	v5380 = *(*int64)(unsafe.Add(mBase, uint32(v173)+208))
	v5383 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[34])))
	if v5383 == int32(1) {
		goto L806
	} else {
		goto L807
	}
L806:
	;
	v5387 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v5391 = m.G0
	v5392 = int32(16)
	v5393 = v5391 - v5392
	m.G0 = v5393
	F_gettimeofday(m, v5393)
	mBase = m.M
	v5396 = *(*int64)(unsafe.Add(mBase, uint32(v5393)))
	v5397 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5393)+8)))
	m.G0 = v5393 + v5392
	v5405 = v5397 + v5396*int64(1000000) - int64(946684800000000)
	goto L809
L807:
	;
	goto L808
L808:
	;
	v5478 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v5478 == int32(0) {
		goto L831
	} else {
		goto L832
	}
L809:
	;
	if v5405 <= v123 {
		v5423 = int32(0)
		goto L811
	} else {
		goto L812
	}
L810:
	;
	if v5372 != 0 {
		goto L814
	} else {
		goto L815
	}
L811:
	;
	goto L810
L812:
	;
	v5411 = v5405 - v123
	if base.B2i32(int64(0) < v123)^base.B2i32(v5411 < v5405)|base.B2i32(int64(2147483646000) < v5411) != 0 {
		v5423 = int32(2147483647)
		goto L811
	} else {
		goto L813
	}
L813:
	;
	v5420 = base.I64_div_s(v5411+int64(999), int64(1000))
	v5423 = base.I32_wrap_i64(v5420)
	goto L811
L814:
	;
	v5426 = int32(0)
	goto L816
L815:
	;
	v5426 = v5387
	goto L816
L816:
	;
	v5429 = F_pgstat_get_entry_ref_locked(m, int32(2), v5426, base.I64_extend_i32_u(v5370), int32(0))
	mBase = m.M
	v5430 = m.ExcPending
	if v5430 != 0 {
		goto L14
	} else {
		goto L817
	}
L817:
	;
	v5431 = *(*int32)(unsafe.Add(mBase, uint32(v5429)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v5431)+120)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5431)+104)) = v5379 + v5380
	*(*int64)(unsafe.Add(mBase, uint32(v5431)+96)) = base.I64_trunc_sat_f64_s(v5377)
	v5439 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[4]))
	v5441 = base.B2i32(v5439 == int32(4))
	if v5439 == int32(4) {
		goto L818
	} else {
		goto L819
	}
L818:
	;
	v5442 = int32(160)
	goto L820
L819:
	;
	v5442 = int32(144)
	goto L820
L820:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5431+v5442))) = v5405
	if v5439 == int32(4) {
		goto L821
	} else {
		goto L822
	}
L821:
	;
	v5447 = int32(168)
	goto L823
L822:
	;
	v5447 = int32(152)
	goto L823
L823:
	;
	v5448 = v5431 + v5447
	v5449 = *(*int64)(unsafe.Add(mBase, uint32(v5448)))
	*(*int64)(unsafe.Add(mBase, uint32(v5448))) = v5449 + int64(1)
	if v5439 == int32(4) {
		goto L824
	} else {
		goto L825
	}
L824:
	;
	v5455 = int32(216)
	goto L826
L825:
	;
	v5455 = int32(208)
	goto L826
L826:
	;
	v5456 = v5431 + v5455
	v5457 = *(*int64)(unsafe.Add(mBase, uint32(v5456)))
	*(*int64)(unsafe.Add(mBase, uint32(v5456))) = v5457 + base.I64_extend_i32_s(v5423)
	F_pgstat_unlock_entry(m, v5429)
	mBase = m.M
	v5462 = m.ExcPending
	if v5462 != 0 {
		goto L14
	} else {
		goto L827
	}
L827:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v5465 = m.ExcPending
	if v5465 != 0 {
		goto L14
	} else {
		goto L828
	}
L828:
	;
	v5468 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v5469 = m.ExcPending
	if v5469 != 0 {
		goto L14
	} else {
		goto L829
	}
L829:
	;
	goto L808
L830:
	;
	if v103 == int32(0) {
		goto L835
	} else {
		goto L836
	}
L831:
	;
	goto L830
L832:
	;
	v5482 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v5482&int32(1) == int32(0) {
		goto L831
	} else {
		goto L833
	}
L833:
	;
	v5487 = *(*int32)(unsafe.Add(mBase, uint32(v5478)+220))
	if v5487 == int32(0) {
		goto L831
	} else {
		goto L834
	}
L834:
	;
	v5490 = int32(_a_F_heap_vacuum_rel_1)
	v5492 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v5493 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5492 + v5493
	v5496 = *(*int32)(unsafe.Add(mBase, uint32(v5478)))
	*(*int32)(unsafe.Add(mBase, uint32(v5478))) = v5496 + v5493
	v5500 = int32(0)
	v5502 = int32(_a_F_heap_vacuum_rel_2)
	v5503 = base.AtomicRmwOr32(m, v5500, v5502, v5500)
	*(*int32)(unsafe.Add(mBase, uint32(v5478)+220)) = v5500
	*(*int32)(unsafe.Add(mBase, uint32(v5478)+224)) = v5500
	v5511 = base.AtomicRmwOr32(m, v5500, v5502, v5500)
	v5512 = *(*int32)(unsafe.Add(mBase, uint32(v5478)))
	*(*int32)(unsafe.Add(mBase, uint32(v5478))) = v5512 + v5493
	v5518 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5518 - v5493
	goto L831
L835:
	;
	v6245 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if int32(0) < v6245 {
		goto L938
	} else {
		goto L939
	}
L836:
	;
	v5527 = m.G0
	v5528 = int32(16)
	v5529 = v5527 - v5528
	m.G0 = v5529
	F_gettimeofday(m, v5529)
	mBase = m.M
	v5532 = *(*int64)(unsafe.Add(mBase, uint32(v5529)))
	v5533 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5529)+8)))
	m.G0 = v5529 + v5528
	v5541 = v5533 + v5532*int64(1000000) - int64(946684800000000)
	goto L837
L837:
	;
	if v73 != 0 {
		goto L838
	} else {
		goto L839
	}
L838:
	;
	v5559 = v5541 - v123
	if v5559 <= int64(0) {
		goto L844
	} else {
		goto L845
	}
L839:
	;
	v5542 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v5542 == int32(0) {
		goto L838
	} else {
		goto L840
	}
L840:
	;
	goto L841
L841:
	;
	if base.B2i32(base.I64_extend_i32_s(v5542)*int64(1000) <= v5541-v123) == int32(0) {
		goto L835
	} else {
		goto L842
	}
L842:
	;
	goto L838
L843:
	;
	v5575 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+552)) = v5575
	*(*int64)(unsafe.Add(mBase, uint32(v52)+544)) = v5575
	*(*int64)(unsafe.Add(mBase, uint32(v52)+536)) = v5575
	*(*int64)(unsafe.Add(mBase, uint32(v52)+528)) = v5575
	v5584 = v52 + int32(528)
	v5586 = v52 + int32(704)
	v5587 = *(*int64)(unsafe.Add(mBase, uint32(v5584)+16))
	v5589 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[1]))
	v5590 = *(*int64)(unsafe.Add(mBase, uint32(v5586)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5584)+16)) = v5587 + (v5589 - v5590)
	v5594 = *(*int64)(unsafe.Add(mBase, uint32(v5584)))
	v5596 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[3]))
	v5597 = *(*int64)(unsafe.Add(mBase, uint32(v5586)))
	*(*int64)(unsafe.Add(mBase, uint32(v5584))) = v5594 + (v5596 - v5597)
	v5601 = *(*int64)(unsafe.Add(mBase, uint32(v5584)+8))
	v5603 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[2]))
	v5604 = *(*int64)(unsafe.Add(mBase, uint32(v5586)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5584)+8)) = v5601 + (v5603 - v5604)
	v5608 = *(*int64)(unsafe.Add(mBase, uint32(v5584)+24))
	v5610 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[0]))
	v5611 = *(*int64)(unsafe.Add(mBase, uint32(v5586)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5584)+24)) = v5608 + (v5610 - v5611)
	goto L847
L844:
	;
	v5571 = int32(0)
	v5572 = int32(0)
	goto L846
L845:
	;
	v5563 = int64(1000000)
	v5564 = base.I64_div_u_s(v5559, v5563)
	v5571 = base.I32_wrap_i64(v5564)
	v5572 = base.I32_wrap_i64(v5559 - v5564*v5563)
	goto L846
L846:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52+int32(1608)))) = v5571
	*(*int32)(unsafe.Add(mBase, uint32(v52+int32(1600)))) = v5572
	goto L843
L847:
	;
	v5616 = v52 + int32(960)
	base.MemoryFill(m, v5616, int32(0), int32(128))
	v5621 = v52 + int32(576)
	v5622 = *(*int64)(unsafe.Add(mBase, uint32(v5616)))
	v5624 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[35]))
	v5625 = *(*int64)(unsafe.Add(mBase, uint32(v5621)))
	*(*int64)(unsafe.Add(mBase, uint32(v5616))) = v5622 + (v5624 - v5625)
	v5629 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+8))
	v5631 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[36]))
	v5632 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+8)) = v5629 + (v5631 - v5632)
	v5636 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+16))
	v5638 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[37]))
	v5639 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+16)) = v5636 + (v5638 - v5639)
	v5643 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+24))
	v5645 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[38]))
	v5646 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+24)) = v5643 + (v5645 - v5646)
	v5650 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+32))
	v5652 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[39]))
	v5653 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+32)) = v5650 + (v5652 - v5653)
	v5657 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+40))
	v5659 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[40]))
	v5660 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+40)) = v5657 + (v5659 - v5660)
	v5664 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+48))
	v5666 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[41]))
	v5667 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+48)) = v5664 + (v5666 - v5667)
	v5671 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+56))
	v5673 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[42]))
	v5674 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+56)) = v5671 + (v5673 - v5674)
	v5678 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+64))
	v5680 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[43]))
	v5681 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+64)) = v5678 + (v5680 - v5681)
	v5685 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+72))
	v5687 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[44]))
	v5688 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+72)) = v5685 + (v5687 - v5688)
	v5692 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+80))
	v5694 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[45]))
	v5695 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+80)) = v5692 + (v5694 - v5695)
	v5699 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+88))
	v5701 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[46]))
	v5702 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+88)) = v5699 + (v5701 - v5702)
	v5706 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+96))
	v5708 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[47]))
	v5709 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+96)) = v5706 + (v5708 - v5709)
	v5713 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+104))
	v5715 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[48]))
	v5716 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+104)) = v5713 + (v5715 - v5716)
	v5720 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+112))
	v5722 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[49]))
	v5723 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+112)) = v5720 + (v5722 - v5723)
	v5727 = *(*int64)(unsafe.Add(mBase, uint32(v5616)+120))
	v5729 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[50]))
	v5730 = *(*int64)(unsafe.Add(mBase, uint32(v5621)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v5616)+120)) = v5727 + (v5729 - v5730)
	goto L848
L848:
	;
	v5734 = *(*int64)(unsafe.Add(mBase, uint32(v52)+976))
	v5735 = *(*int64)(unsafe.Add(mBase, uint32(v52)+1008))
	v5736 = *(*int64)(unsafe.Add(mBase, uint32(v52)+968))
	v5737 = *(*int64)(unsafe.Add(mBase, uint32(v52)+1000))
	v5738 = *(*int64)(unsafe.Add(mBase, uint32(v52)+960))
	v5739 = *(*int64)(unsafe.Add(mBase, uint32(v52)+992))
	F_initStringInfo(m, v52+int32(928))
	mBase = m.M
	v5743 = m.ExcPending
	if v5743 != 0 {
		goto L14
	} else {
		goto L849
	}
L849:
	;
	if v73 != 0 {
		v5760 = int32(_a_F_heap_vacuum_rel_39)
		goto L850
	} else {
		goto L851
	}
L850:
	;
	v5761 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	v5762 = *(*int64)(unsafe.Add(mBase, uint32(v173)+68))
	v5763 = *(*int32)(unsafe.Add(mBase, uint32(v173)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+396)) = v5763
	*(*int64)(unsafe.Add(mBase, uint32(v52)+384)) = v5762
	*(*int32)(unsafe.Add(mBase, uint32(v52)+392)) = v5761
	v5768 = v52 + int32(928)
	F_appendStringInfo(m, v5768, v5760, v52+int32(384))
	mBase = m.M
	v5772 = m.ExcPending
	if v5772 != 0 {
		goto L14
	} else {
		goto L859
	}
L851:
	;
	v5747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)))
	if v5747&int32(1) != 0 {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	v5750 = int32(_a_F_heap_vacuum_rel_40)
	goto L854
L853:
	;
	v5750 = int32(_a_F_heap_vacuum_rel_41)
	goto L854
L854:
	;
	v5751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v5751 == int32(1) {
		v5760 = v5750
		goto L850
	} else {
		goto L855
	}
L855:
	;
	if v5747&int32(1) != 0 {
		goto L856
	} else {
		goto L857
	}
L856:
	;
	v5758 = int32(_a_F_heap_vacuum_rel_42)
	goto L858
L857:
	;
	v5758 = int32(_a_F_heap_vacuum_rel_43)
	goto L858
L858:
	;
	v5760 = v5758
	goto L850
L859:
	;
	v5773 = *(*int32)(unsafe.Add(mBase, uint32(v173)+112))
	v5774 = *(*int32)(unsafe.Add(mBase, uint32(v173)+120))
	v5775 = *(*int32)(unsafe.Add(mBase, uint32(v173)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+376)) = v5775
	v5778 = float64(100)
	if v402 != 0 {
		goto L860
	} else {
		goto L861
	}
L860:
	;
	v5783 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v5773), v5778), base.F64_convert_i32_u(v402))
	goto L862
L861:
	;
	v5783 = v5778
	goto L862
L862:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+368)) = v5783
	*(*int32)(unsafe.Add(mBase, uint32(v52)+360)) = v5773
	*(*int32)(unsafe.Add(mBase, uint32(v52)+356)) = v5341
	*(*int32)(unsafe.Add(mBase, uint32(v52)+352)) = v5774
	F_appendStringInfo(m, v5768, int32(_a_F_heap_vacuum_rel_44), v52+int32(352))
	mBase = m.M
	v5792 = m.ExcPending
	if v5792 != 0 {
		goto L14
	} else {
		goto L863
	}
L863:
	;
	v5793 = *(*float64)(unsafe.Add(mBase, uint32(v173)+152))
	v5794 = *(*int64)(unsafe.Add(mBase, uint32(v173)+176))
	v5795 = *(*int64)(unsafe.Add(mBase, uint32(v173)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+336)) = v5795
	*(*int64)(unsafe.Add(mBase, uint32(v52)+320)) = v5794
	*(*int64)(unsafe.Add(mBase, uint32(v52)+328)) = base.I64_trunc_sat_f64_s(v5793)
	F_appendStringInfo(m, v5768, int32(_a_F_heap_vacuum_rel_45), v52+int32(320))
	mBase = m.M
	v5804 = m.ExcPending
	if v5804 != 0 {
		goto L14
	} else {
		goto L864
	}
L864:
	;
	v5805 = *(*int64)(unsafe.Add(mBase, uint32(v173)+216))
	if int64(0) < v5805 {
		goto L865
	} else {
		goto L866
	}
L865:
	;
	v5808 = *(*int32)(unsafe.Add(mBase, uint32(v173)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+312)) = v5808
	*(*int64)(unsafe.Add(mBase, uint32(v52)+304)) = v5805
	F_appendStringInfo(m, v5768, int32(_a_F_heap_vacuum_rel_46), v52+int32(304))
	mBase = m.M
	v5815 = m.ExcPending
	if v5815 != 0 {
		goto L14
	} else {
		goto L868
	}
L866:
	;
	goto L867
L867:
	;
	v5816 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v5817 = m.ExcPending
	if v5817 != 0 {
		goto L14
	} else {
		goto L869
	}
L868:
	;
	goto L867
L869:
	;
	v5818 = *(*int32)(unsafe.Add(mBase, uint32(v173)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+288)) = v5818
	*(*int32)(unsafe.Add(mBase, uint32(v52)+292)) = base.I32_wrap_i64(v5816) - v5818
	v5824 = v52 + int32(928)
	F_appendStringInfo(m, v5824, int32(_a_F_heap_vacuum_rel_47), v52+int32(288))
	mBase = m.M
	v5829 = m.ExcPending
	if v5829 != 0 {
		goto L14
	} else {
		goto L870
	}
L870:
	;
	v5830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+924)))
	if v5830 == int32(1) {
		goto L871
	} else {
		goto L872
	}
L871:
	;
	v5833 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v5834 = *(*int32)(unsafe.Add(mBase, uint32(v1303)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+272)) = v5834
	*(*int32)(unsafe.Add(mBase, uint32(v52)+276)) = v5834 - v5833
	F_appendStringInfo(m, v5824, int32(_a_F_heap_vacuum_rel_48), v52+int32(272))
	mBase = m.M
	v5842 = m.ExcPending
	if v5842 != 0 {
		goto L14
	} else {
		goto L874
	}
L872:
	;
	goto L873
L873:
	;
	v5845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+908)))
	if v5845 == int32(1) {
		goto L875
	} else {
		goto L876
	}
L874:
	;
	goto L873
L875:
	;
	v5848 = *(*int32)(unsafe.Add(mBase, uint32(v173)+32))
	v5849 = *(*int32)(unsafe.Add(mBase, uint32(v173)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+256)) = v5849
	*(*int32)(unsafe.Add(mBase, uint32(v52)+260)) = v5849 - v5848
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_49), v52+int32(256))
	mBase = m.M
	v5859 = m.ExcPending
	if v5859 != 0 {
		goto L14
	} else {
		goto L878
	}
L876:
	;
	goto L877
L877:
	;
	v5862 = *(*int32)(unsafe.Add(mBase, uint32(v173)+124))
	v5863 = *(*int64)(unsafe.Add(mBase, uint32(v173)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+240)) = v5863
	v5866 = float64(100)
	if v402 != 0 {
		goto L879
	} else {
		goto L880
	}
L878:
	;
	goto L877
L879:
	;
	v5871 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v5862), v5866), base.F64_convert_i32_u(v402))
	goto L881
L880:
	;
	v5871 = v5866
	goto L881
L881:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+232)) = v5871
	*(*int32)(unsafe.Add(mBase, uint32(v52)+224)) = v5862
	v5875 = v52 + int32(928)
	F_appendStringInfo(m, v5875, int32(_a_F_heap_vacuum_rel_50), v52+int32(224))
	mBase = m.M
	v5880 = m.ExcPending
	if v5880 != 0 {
		goto L14
	} else {
		goto L882
	}
L882:
	;
	v5881 = *(*int32)(unsafe.Add(mBase, uint32(v173)+132))
	v5882 = *(*int32)(unsafe.Add(mBase, uint32(v173)+128))
	v5883 = *(*int32)(unsafe.Add(mBase, uint32(v173)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+216)) = v5883
	*(*int32)(unsafe.Add(mBase, uint32(v52)+208)) = v5882
	*(*int32)(unsafe.Add(mBase, uint32(v52)+212)) = v5883 + v5881
	F_appendStringInfo(m, v5875, int32(_a_F_heap_vacuum_rel_51), v52+int32(208))
	mBase = m.M
	v5892 = m.ExcPending
	if v5892 != 0 {
		goto L14
	} else {
		goto L883
	}
L883:
	;
	v5893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+23)))
	if v5893 == int32(1) {
		goto L885
	} else {
		goto L886
	}
L884:
	;
	F_appendStringInfoString(m, v5875, v5912)
	mBase = m.M
	v5914 = m.ExcPending
	if v5914 != 0 {
		goto L14
	} else {
		goto L895
	}
L885:
	;
	v5896 = int32(_a_F_heap_vacuum_rel_52)
	v5898 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v5898 == int32(0) {
		v5911 = v5896
		v5912 = int32(_a_F_heap_vacuum_rel_53)
		goto L884
	} else {
		goto L888
	}
L886:
	;
	goto L887
L887:
	;
	v5909 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])))
	if v5909 != 0 {
		goto L892
	} else {
		goto L893
	}
L888:
	;
	v5903 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	if v5903 != 0 {
		goto L889
	} else {
		goto L890
	}
L889:
	;
	v5904 = int32(_a_F_heap_vacuum_rel_54)
	goto L891
L890:
	;
	v5904 = int32(_a_F_heap_vacuum_rel_53)
	goto L891
L891:
	;
	v5911 = v5896
	v5912 = v5904
	goto L884
L892:
	;
	v5910 = int32(_a_F_heap_vacuum_rel_55)
	goto L894
L893:
	;
	v5910 = int32(_a_F_heap_vacuum_rel_56)
	goto L894
L894:
	;
	v5911 = int32(_a_F_heap_vacuum_rel_57)
	v5912 = v5910
	goto L884
L895:
	;
	v5915 = *(*int32)(unsafe.Add(mBase, uint32(v173+int32(140))))
	v5916 = *(*int64)(unsafe.Add(mBase, uint32(v173)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+192)) = v5916
	v5919 = float64(100)
	if v402 != 0 {
		goto L896
	} else {
		goto L897
	}
L896:
	;
	v5924 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v5915), v5919), base.F64_convert_i32_u(v402))
	goto L898
L897:
	;
	v5924 = v5919
	goto L898
L898:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+184)) = v5924
	*(*int32)(unsafe.Add(mBase, uint32(v52)+176)) = v5915
	F_appendStringInfo(m, v52+int32(928), v5911, v52+int32(176))
	mBase = m.M
	v5932 = m.ExcPending
	if v5932 != 0 {
		goto L14
	} else {
		goto L899
	}
L899:
	;
	v5933 = int32(0)
	v5934 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if v5933 < v5934 {
		goto L900
	} else {
		goto L901
	}
L900:
	;
	v5944 = v5933
	v5947 = v5934
	goto L903
L901:
	;
	goto L902
L902:
	;
	v6067 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[51])))
	if v6067 != 0 {
		goto L910
	} else {
		goto L911
	}
L903:
	;
	v5989 = v5944 << (uint(int32(2)) % 32)
	v5990 = *(*int32)(unsafe.Add(mBase, uint32(v173)+168))
	v5992 = *(*int32)(unsafe.Add(mBase, uint32(v5989+v5990)))
	if v5992 != 0 {
		goto L905
	} else {
		goto L906
	}
L904:
	;
	goto L902
L905:
	;
	v5994 = *(*int32)(unsafe.Add(mBase, uint32(v5989+v327)))
	v5995 = *(*int64)(unsafe.Add(mBase, uint32(v5992)+24))
	v5996 = *(*int32)(unsafe.Add(mBase, uint32(v5992)))
	v5997 = *(*int32)(unsafe.Add(mBase, uint32(v5992)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v52+int32(160)))) = v5997
	*(*int32)(unsafe.Add(mBase, uint32(v52)+148)) = v5996
	*(*int64)(unsafe.Add(mBase, uint32(v52)+152)) = v5995
	*(*int32)(unsafe.Add(mBase, uint32(v52)+144)) = v5994
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_58), v52+int32(144))
	mBase = m.M
	v6008 = m.ExcPending
	if v6008 != 0 {
		goto L14
	} else {
		goto L908
	}
L906:
	;
	v6012 = v5947
	goto L907
L907:
	;
	v6015 = v5944 + int32(1)
	if v6015 < v6012 {
		v5944 = v6015
		v5947 = v6012
		goto L903
	} else {
		goto L909
	}
L908:
	;
	v6009 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v6012 = v6009
	goto L907
L909:
	;
	goto L904
L910:
	;
	v6069 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	v6070 = *(*int64)(unsafe.Add(mBase, uint32(v6069)+312))
	*(*float64)(unsafe.Add(mBase, uint32(v52)+128)) = base.F64_div(base.F64_convert_i64_s(v6070), float64(1e+06))
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_59), v52+int32(128))
	mBase = m.M
	v6081 = m.ExcPending
	if v6081 != 0 {
		goto L14
	} else {
		goto L913
	}
L911:
	;
	goto L912
L912:
	;
	v6083 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5])))
	if v6083 == int32(1) {
		goto L914
	} else {
		goto L915
	}
L913:
	;
	goto L912
L914:
	;
	v6087 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[7]))
	v6090 = float64(1000)
	*(*float64)(unsafe.Add(mBase, uint32(v52)+112)) = base.F64_div(base.F64_convert_i64_s(v6087-v105), v6090)
	v6094 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[6]))
	*(*float64)(unsafe.Add(mBase, uint32(v52)+120)) = base.F64_div(base.F64_convert_i64_s(v6094-v104), v6090)
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_60), v52+int32(112))
	mBase = m.M
	v6106 = m.ExcPending
	if v6106 != 0 {
		goto L14
	} else {
		goto L917
	}
L915:
	;
	goto L916
L916:
	;
	v6107 = v5735 + v5734
	v6108 = v5737 + v5736
	v6110 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1600))
	v6111 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	if v6111 <= int32(0) {
		goto L919
	} else {
		goto L920
	}
L917:
	;
	goto L916
L918:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+104)) = v6135
	*(*float64)(unsafe.Add(mBase, uint32(v52)+96)) = v6136
	v6140 = v52 + int32(928)
	F_appendStringInfo(m, v6140, int32(_a_F_heap_vacuum_rel_61), v52+int32(96))
	mBase = m.M
	v6145 = m.ExcPending
	if v6145 != 0 {
		goto L14
	} else {
		goto L923
	}
L919:
	;
	if v6110 <= int32(0) {
		v6135 = float64(0)
		v6136 = float64(0)
		goto L918
	} else {
		goto L922
	}
L920:
	;
	goto L921
L921:
	;
	v6118 = float64(8192)
	v6120 = float64(9.5367431640625e-07)
	v6126 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v6110), float64(1e+06)), base.F64_convert_i32_s(v6111))
	v6135 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_convert_i64_s(v6107), v6118), v6120), v6126)
	v6136 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_convert_i64_s(v6108), v6118), v6120), v6126)
	goto L918
L922:
	;
	goto L921
L923:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+80)) = v6107
	*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = v6108
	*(*int64)(unsafe.Add(mBase, uint32(v52)+64)) = v5739 + v5738
	F_appendStringInfo(m, v6140, int32(_a_F_heap_vacuum_rel_62), v52-int32(-64))
	mBase = m.M
	v6153 = m.ExcPending
	if v6153 != 0 {
		goto L14
	} else {
		goto L924
	}
L924:
	;
	v6154 = *(*int64)(unsafe.Add(mBase, uint32(v52)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+48)) = v6154
	v6156 = *(*int64)(unsafe.Add(mBase, uint32(v52)+552))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+56)) = v6156
	v6158 = *(*int64)(unsafe.Add(mBase, uint32(v52)+528))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+32)) = v6158
	v6160 = *(*int64)(unsafe.Add(mBase, uint32(v52)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+40)) = v6160
	F_appendStringInfo(m, v6140, int32(_a_F_heap_vacuum_rel_63), v52+int32(32))
	mBase = m.M
	v6166 = m.ExcPending
	if v6166 != 0 {
		goto L14
	} else {
		goto L925
	}
L925:
	;
	v6169 = F_pg_rusage_show(m, v52+int32(736))
	mBase = m.M
	v6170 = m.ExcPending
	if v6170 != 0 {
		goto L14
	} else {
		goto L926
	}
L926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v6169
	F_appendStringInfo(m, v6140, int32(_a_F_heap_vacuum_rel_64), v52+int32(16))
	mBase = m.M
	v6176 = m.ExcPending
	if v6176 != 0 {
		goto L14
	} else {
		goto L927
	}
L927:
	;
	if v73 != 0 {
		goto L928
	} else {
		goto L929
	}
L928:
	;
	v6179 = int32(17)
	goto L930
L929:
	;
	v6179 = int32(15)
	goto L930
L930:
	;
	v6181 = F_errstart(m, v6179, int32(0))
	mBase = m.M
	v6182 = m.ExcPending
	if v6182 != 0 {
		goto L14
	} else {
		goto L931
	}
L931:
	;
	if v6181 != 0 {
		goto L932
	} else {
		goto L933
	}
L932:
	;
	v6183 = *(*int32)(unsafe.Add(mBase, uint32(v52)+928))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v6183
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_65), v52)
	mBase = m.M
	v6187 = m.ExcPending
	if v6187 != 0 {
		goto L14
	} else {
		goto L935
	}
L933:
	;
	goto L934
L934:
	;
	v6193 = *(*int32)(unsafe.Add(mBase, uint32(v52)+928))
	F_pfree(m, v6193)
	mBase = m.M
	v6195 = m.ExcPending
	if v6195 != 0 {
		goto L14
	} else {
		goto L937
	}
L935:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(1147), int32(_a_F_heap_vacuum_rel_8))
	mBase = m.M
	v6192 = m.ExcPending
	if v6192 != 0 {
		goto L14
	} else {
		goto L936
	}
L936:
	;
	goto L934
L937:
	;
	goto L835
L938:
	;
	v6250 = v5356
	goto L941
L939:
	;
	goto L940
L940:
	;
	m.G0 = v52 + int32(1616)
	return
L941:
	;
	v6298 = v6250 << (uint(int32(2)) % 32)
	v6299 = *(*int32)(unsafe.Add(mBase, uint32(v173)+168))
	v6301 = *(*int32)(unsafe.Add(mBase, uint32(v6298+v6299)))
	if v6301 != 0 {
		goto L943
	} else {
		goto L944
	}
L942:
	;
	goto L940
L943:
	;
	F_pfree(m, v6301)
	mBase = m.M
	v6303 = m.ExcPending
	if v6303 != 0 {
		goto L14
	} else {
		goto L946
	}
L944:
	;
	goto L945
L945:
	;
	if v103 != 0 {
		goto L947
	} else {
		goto L948
	}
L946:
	;
	goto L945
L947:
	;
	v6305 = *(*int32)(unsafe.Add(mBase, uint32(v6298+v327)))
	F_pfree(m, v6305)
	mBase = m.M
	v6307 = m.ExcPending
	if v6307 != 0 {
		goto L14
	} else {
		goto L950
	}
L948:
	;
	goto L949
L949:
	;
	v6309 = v6250 + int32(1)
	v6310 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v6309 < v6310 {
		v6250 = v6309
		goto L941
	} else {
		goto L951
	}
L950:
	;
	goto L949
L951:
	;
	goto L942
}
