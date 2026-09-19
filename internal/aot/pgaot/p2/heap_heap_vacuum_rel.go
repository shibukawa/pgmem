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
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v319 int32
	_ = v319
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int64
	_ = v369
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int64
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v418 float64
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 float64
	_ = v473
	var v476 int32
	_ = v476
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int64
	_ = v498
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v597 int32
	_ = v597
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v896 int32
	_ = v896
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v985 int64
	_ = v985
	var v986 int64
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1120 int32
	_ = v1120
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1311 int64
	_ = v1311
	var v1313 int64
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1318 int64
	_ = v1318
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1463 int64
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1498 int32
	_ = v1498
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1566 int32
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int64
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1632 int32
	_ = v1632
	var v1638 int32
	_ = v1638
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1726 int32
	_ = v1726
	var v1732 int32
	_ = v1732
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1869 int32
	_ = v1869
	var v1878 int32
	_ = v1878
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1939 int32
	_ = v1939
	var v1952 int32
	_ = v1952
	var v1965 int32
	_ = v1965
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1993 int32
	_ = v1993
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2065 int64
	_ = v2065
	var v2066 int64
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2071 int64
	_ = v2071
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2100 int32
	_ = v2100
	var v2115 int64
	_ = v2115
	var v2118 int64
	_ = v2118
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2142 int32
	_ = v2142
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int64
	_ = v2148
	var v2149 int64
	_ = v2149
	var v2152 int32
	_ = v2152
	var v2153 int64
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2189 int32
	_ = v2189
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2303 int64
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2338 int64
	_ = v2338
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2379 int64
	_ = v2379
	var v2382 int64
	_ = v2382
	var v2390 int64
	_ = v2390
	var v2393 int64
	_ = v2393
	var v2396 int64
	_ = v2396
	var v2402 int32
	_ = v2402
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2430 int32
	_ = v2430
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int64
	_ = v2507
	var v2511 int32
	_ = v2511
	var v2512 int64
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2545 int32
	_ = v2545
	var v2548 int32
	_ = v2548
	var v2643 int32
	_ = v2643
	var v2646 int32
	_ = v2646
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2662 int64
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2702 int64
	_ = v2702
	var v2703 int64
	_ = v2703
	var v2706 int64
	_ = v2706
	var v2710 int64
	_ = v2710
	var v2714 int64
	_ = v2714
	var v2715 int64
	_ = v2715
	var v2718 int64
	_ = v2718
	var v2719 int64
	_ = v2719
	var v2722 int32
	_ = v2722
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2744 int32
	_ = v2744
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2759 int32
	_ = v2759
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2767 int32
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2782 int32
	_ = v2782
	var v2785 int32
	_ = v2785
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2794 int32
	_ = v2794
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2810 int32
	_ = v2810
	var v2815 int32
	_ = v2815
	var v2817 int32
	_ = v2817
	var v2819 int32
	_ = v2819
	var v2822 int32
	_ = v2822
	var v2824 int32
	_ = v2824
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2840 int32
	_ = v2840
	var v2845 int32
	_ = v2845
	var v2848 int32
	_ = v2848
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2853 int32
	_ = v2853
	var v2856 int32
	_ = v2856
	var v2861 int32
	_ = v2861
	var v2866 int32
	_ = v2866
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2873 int32
	_ = v2873
	var v2879 int32
	_ = v2879
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2899 int32
	_ = v2899
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2912 int32
	_ = v2912
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2928 int32
	_ = v2928
	var v2929 int32
	_ = v2929
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2935 int64
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2944 int32
	_ = v2944
	var v2949 int32
	_ = v2949
	var v2955 int32
	_ = v2955
	var v2963 int32
	_ = v2963
	var v2967 int32
	_ = v2967
	var v2970 int32
	_ = v2970
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3024 int32
	_ = v3024
	var v3026 int32
	_ = v3026
	var v3037 int32
	_ = v3037
	var v3042 int32
	_ = v3042
	var v3051 int32
	_ = v3051
	var v3060 int32
	_ = v3060
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3099 int32
	_ = v3099
	var v3101 int32
	_ = v3101
	var v3103 int32
	_ = v3103
	var v3106 int32
	_ = v3106
	var v3108 int32
	_ = v3108
	var v3112 int32
	_ = v3112
	var v3116 int32
	_ = v3116
	var v3121 int32
	_ = v3121
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3127 int32
	_ = v3127
	var v3135 int32
	_ = v3135
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3147 int64
	_ = v3147
	var v3148 float64
	_ = v3148
	var v3153 int32
	_ = v3153
	var v3154 float32
	_ = v3154
	var v3155 float64
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3169 float64
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3193 float64
	_ = v3193
	var v3198 float64
	_ = v3198
	var v3200 float64
	_ = v3200
	var v3203 float64
	_ = v3203
	var v3204 int64
	_ = v3204
	var v3207 int64
	_ = v3207
	var v3212 int32
	_ = v3212
	var v3213 int32
	_ = v3213
	var v3214 int64
	_ = v3214
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3226 int32
	_ = v3226
	var v3230 int32
	_ = v3230
	var v3235 int32
	_ = v3235
	var v3237 int32
	_ = v3237
	var v3238 int32
	_ = v3238
	var v3241 int32
	_ = v3241
	var v3249 int32
	_ = v3249
	var v3255 int32
	_ = v3255
	var v3259 int32
	_ = v3259
	var v3262 int32
	_ = v3262
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3267 float64
	_ = v3267
	var v3276 int64
	_ = v3276
	var v3294 int32
	_ = v3294
	var v3298 int32
	_ = v3298
	var v3303 int32
	_ = v3303
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3309 int32
	_ = v3309
	var v3404 int32
	_ = v3404
	var v3407 int32
	_ = v3407
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3423 int64
	_ = v3423
	var v3425 int32
	_ = v3425
	var v3428 int32
	_ = v3428
	var v3439 int32
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3443 int32
	_ = v3443
	var v3445 int32
	_ = v3445
	var v3458 int32
	_ = v3458
	var v3461 int32
	_ = v3461
	var v3504 int64
	_ = v3504
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3521 int32
	_ = v3521
	var v3523 int32
	_ = v3523
	var v3525 int32
	_ = v3525
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3535 int32
	_ = v3535
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3544 int32
	_ = v3544
	var v3547 int32
	_ = v3547
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3557 int32
	_ = v3557
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3567 int64
	_ = v3567
	var v3570 int32
	_ = v3570
	var v3574 int32
	_ = v3574
	var v3579 int32
	_ = v3579
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3585 int32
	_ = v3585
	var v3593 int32
	_ = v3593
	var v3599 int32
	_ = v3599
	var v3603 int64
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3610 int32
	_ = v3610
	var v3615 int32
	_ = v3615
	var v3679 int32
	_ = v3679
	var v3683 int32
	_ = v3683
	var v3688 int32
	_ = v3688
	var v3690 int32
	_ = v3690
	var v3691 int32
	_ = v3691
	var v3694 int32
	_ = v3694
	var v3789 int32
	_ = v3789
	var v3792 int32
	_ = v3792
	var v3801 int32
	_ = v3801
	var v3802 int32
	_ = v3802
	var v3808 int64
	_ = v3808
	var v3810 int32
	_ = v3810
	var v3813 int32
	_ = v3813
	var v3824 int32
	_ = v3824
	var v3825 int32
	_ = v3825
	var v3828 int32
	_ = v3828
	var v3830 int32
	_ = v3830
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3906 int32
	_ = v3906
	var v3947 int32
	_ = v3947
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3961 int64
	_ = v3961
	var v3963 int64
	_ = v3963
	var v3965 int64
	_ = v3965
	var v3967 int64
	_ = v3967
	var v3969 int64
	_ = v3969
	var v3978 int32
	_ = v3978
	var v3979 int32
	_ = v3979
	var v4030 int32
	_ = v4030
	var v4032 int32
	_ = v4032
	var v4033 int32
	_ = v4033
	var v4035 int32
	_ = v4035
	var v4038 int32
	_ = v4038
	var v4039 int32
	_ = v4039
	var v4043 int32
	_ = v4043
	var v4045 int32
	_ = v4045
	var v4047 int32
	_ = v4047
	var v4099 int32
	_ = v4099
	var v4100 int32
	_ = v4100
	var v4103 int32
	_ = v4103
	var v4107 int32
	_ = v4107
	var v4111 int32
	_ = v4111
	var v4159 int32
	_ = v4159
	var v4161 int32
	_ = v4161
	var v4164 int32
	_ = v4164
	var v4166 int32
	_ = v4166
	var v4167 int32
	_ = v4167
	var v4168 float64
	_ = v4168
	var v4169 int32
	_ = v4169
	var v4178 int32
	_ = v4178
	var v4180 int32
	_ = v4180
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4190 int32
	_ = v4190
	var v4233 int32
	_ = v4233
	var v4236 int32
	_ = v4236
	var v4237 int32
	_ = v4237
	var v4241 int32
	_ = v4241
	var v4244 int32
	_ = v4244
	var v4245 int32
	_ = v4245
	var v4247 int32
	_ = v4247
	var v4258 int32
	_ = v4258
	var v4262 int32
	_ = v4262
	var v4267 int32
	_ = v4267
	var v4269 int32
	_ = v4269
	var v4270 int32
	_ = v4270
	var v4273 int32
	_ = v4273
	var v4281 int32
	_ = v4281
	var v4287 int32
	_ = v4287
	var v4293 int32
	_ = v4293
	var v4295 int32
	_ = v4295
	var v4304 int32
	_ = v4304
	var v4347 int32
	_ = v4347
	var v4348 int32
	_ = v4348
	var v4349 int32
	_ = v4349
	var v4354 int32
	_ = v4354
	var v4402 int32
	_ = v4402
	var v4404 int32
	_ = v4404
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4416 int32
	_ = v4416
	var v4422 int32
	_ = v4422
	var v4427 int32
	_ = v4427
	var v4429 int32
	_ = v4429
	var v4433 int32
	_ = v4433
	var v4434 int32
	_ = v4434
	var v4436 int32
	_ = v4436
	var v4441 int32
	_ = v4441
	var v4442 int32
	_ = v4442
	var v4443 int32
	_ = v4443
	var v4495 int32
	_ = v4495
	var v4497 int32
	_ = v4497
	var v4498 int32
	_ = v4498
	var v4500 int32
	_ = v4500
	var v4502 int32
	_ = v4502
	var v4507 int32
	_ = v4507
	var v4508 int32
	_ = v4508
	var v4509 int32
	_ = v4509
	var v4511 int64
	_ = v4511
	var v4512 int64
	_ = v4512
	var v4522 int32
	_ = v4522
	var v4530 int32
	_ = v4530
	var v4555 int64
	_ = v4555
	var v4572 int64
	_ = v4572
	var v4573 int64
	_ = v4573
	var v4576 int64
	_ = v4576
	var v4580 int32
	_ = v4580
	var v4581 int32
	_ = v4581
	var v4583 int32
	_ = v4583
	var v4585 int32
	_ = v4585
	var v4587 int32
	_ = v4587
	var v4591 int32
	_ = v4591
	var v4593 int32
	_ = v4593
	var v4595 int32
	_ = v4595
	var v4604 int32
	_ = v4604
	var v4605 int32
	_ = v4605
	var v4608 int64
	_ = v4608
	var v4610 int64
	_ = v4610
	var v4614 int32
	_ = v4614
	var v4616 int32
	_ = v4616
	var v4621 int32
	_ = v4621
	var v4622 int32
	_ = v4622
	var v4623 int64
	_ = v4623
	var v4628 int32
	_ = v4628
	var v4629 int32
	_ = v4629
	var v4632 int32
	_ = v4632
	var v4633 int32
	_ = v4633
	var v4639 int32
	_ = v4639
	var v4644 int32
	_ = v4644
	var v4646 int32
	_ = v4646
	var v4647 int32
	_ = v4647
	var v4654 int32
	_ = v4654
	var v4656 int32
	_ = v4656
	var v4657 int32
	_ = v4657
	var v4658 int32
	_ = v4658
	var v4659 int32
	_ = v4659
	var v4667 int32
	_ = v4667
	var v4670 int32
	_ = v4670
	var v4671 int32
	_ = v4671
	var v4672 int32
	_ = v4672
	var v4673 int32
	_ = v4673
	var v4679 int32
	_ = v4679
	var v4684 int32
	_ = v4684
	var v4686 int32
	_ = v4686
	var v4687 int32
	_ = v4687
	var v4688 int32
	_ = v4688
	var v4689 int32
	_ = v4689
	var v4690 int32
	_ = v4690
	var v4692 int32
	_ = v4692
	var v4696 int32
	_ = v4696
	var v4704 int32
	_ = v4704
	var v4708 int32
	_ = v4708
	var v4713 int32
	_ = v4713
	var v4717 int32
	_ = v4717
	var v4724 int32
	_ = v4724
	var v4729 int32
	_ = v4729
	var v4735 int32
	_ = v4735
	var v4738 int32
	_ = v4738
	var v4739 int32
	_ = v4739
	var v4741 int32
	_ = v4741
	var v4742 int32
	_ = v4742
	var v4745 int32
	_ = v4745
	var v4751 int32
	_ = v4751
	var v4756 int32
	_ = v4756
	var v4763 int64
	_ = v4763
	var v4766 int32
	_ = v4766
	var v4768 int32
	_ = v4768
	var v4770 int32
	_ = v4770
	var v4773 int32
	_ = v4773
	var v4776 int32
	_ = v4776
	var v4825 int32
	_ = v4825
	var v4828 int32
	_ = v4828
	var v4830 int32
	_ = v4830
	var v4832 int32
	_ = v4832
	var v4849 int32
	_ = v4849
	var v4885 int32
	_ = v4885
	var v4886 int32
	_ = v4886
	var v4888 int32
	_ = v4888
	var v4889 int32
	_ = v4889
	var v4890 int32
	_ = v4890
	var v4893 int32
	_ = v4893
	var v4897 int32
	_ = v4897
	var v4903 int32
	_ = v4903
	var v4905 int32
	_ = v4905
	var v4911 int32
	_ = v4911
	var v4912 int32
	_ = v4912
	var v4915 int32
	_ = v4915
	var v4923 int32
	_ = v4923
	var v4931 int32
	_ = v4931
	var v4983 int32
	_ = v4983
	var v4989 int32
	_ = v4989
	var v4994 int32
	_ = v4994
	var v5045 int32
	_ = v5045
	var v5046 int32
	_ = v5046
	var v5056 int32
	_ = v5056
	var v5069 int32
	_ = v5069
	var v5098 int32
	_ = v5098
	var v5101 int32
	_ = v5101
	var v5103 int32
	_ = v5103
	var v5104 int32
	_ = v5104
	var v5106 int32
	_ = v5106
	var v5108 int32
	_ = v5108
	var v5114 int32
	_ = v5114
	var v5115 int32
	_ = v5115
	var v5117 int32
	_ = v5117
	var v5118 int32
	_ = v5118
	var v5119 int32
	_ = v5119
	var v5127 int32
	_ = v5127
	var v5132 int32
	_ = v5132
	var v5134 int32
	_ = v5134
	var v5187 int32
	_ = v5187
	var v5193 int32
	_ = v5193
	var v5197 int32
	_ = v5197
	var v5202 int32
	_ = v5202
	var v5204 int32
	_ = v5204
	var v5205 int32
	_ = v5205
	var v5208 int32
	_ = v5208
	var v5216 int32
	_ = v5216
	var v5222 int32
	_ = v5222
	var v5226 int32
	_ = v5226
	var v5229 int32
	_ = v5229
	var v5233 int32
	_ = v5233
	var v5239 int32
	_ = v5239
	var v5240 int32
	_ = v5240
	var v5243 int32
	_ = v5243
	var v5244 int32
	_ = v5244
	var v5247 int32
	_ = v5247
	var v5248 int32
	_ = v5248
	var v5249 float64
	_ = v5249
	var v5250 int32
	_ = v5250
	var v5253 int32
	_ = v5253
	var v5254 int32
	_ = v5254
	var v5261 int32
	_ = v5261
	var v5262 int32
	_ = v5262
	var v5263 int32
	_ = v5263
	var v5264 int32
	_ = v5264
	var v5265 float64
	_ = v5265
	var v5266 float64
	_ = v5266
	var v5269 float64
	_ = v5269
	var v5271 int64
	_ = v5271
	var v5272 int64
	_ = v5272
	var v5275 int32
	_ = v5275
	var v5279 int32
	_ = v5279
	var v5283 int32
	_ = v5283
	var v5284 int32
	_ = v5284
	var v5285 int32
	_ = v5285
	var v5288 int64
	_ = v5288
	var v5289 int64
	_ = v5289
	var v5297 int64
	_ = v5297
	var v5303 int64
	_ = v5303
	var v5312 int64
	_ = v5312
	var v5315 int32
	_ = v5315
	var v5318 int32
	_ = v5318
	var v5321 int32
	_ = v5321
	var v5322 int32
	_ = v5322
	var v5323 int32
	_ = v5323
	var v5331 int32
	_ = v5331
	var v5333 int32
	_ = v5333
	var v5334 int32
	_ = v5334
	var v5339 int32
	_ = v5339
	var v5340 int32
	_ = v5340
	var v5341 int64
	_ = v5341
	var v5347 int32
	_ = v5347
	var v5348 int32
	_ = v5348
	var v5349 int64
	_ = v5349
	var v5354 int32
	_ = v5354
	var v5357 int32
	_ = v5357
	var v5360 int32
	_ = v5360
	var v5361 int32
	_ = v5361
	var v5370 int32
	_ = v5370
	var v5374 int32
	_ = v5374
	var v5379 int32
	_ = v5379
	var v5382 int32
	_ = v5382
	var v5384 int32
	_ = v5384
	var v5385 int32
	_ = v5385
	var v5388 int32
	_ = v5388
	var v5392 int32
	_ = v5392
	var v5402 int32
	_ = v5402
	var v5411 int32
	_ = v5411
	var v5412 int32
	_ = v5412
	var v5413 int32
	_ = v5413
	var v5416 int64
	_ = v5416
	var v5417 int64
	_ = v5417
	var v5425 int64
	_ = v5425
	var v5426 int32
	_ = v5426
	var v5443 int64
	_ = v5443
	var v5447 int64
	_ = v5447
	var v5448 int64
	_ = v5448
	var v5455 int32
	_ = v5455
	var v5456 int32
	_ = v5456
	var v5459 int64
	_ = v5459
	var v5468 int32
	_ = v5468
	var v5470 int32
	_ = v5470
	var v5471 int64
	_ = v5471
	var v5473 int64
	_ = v5473
	var v5474 int64
	_ = v5474
	var v5478 int64
	_ = v5478
	var v5480 int64
	_ = v5480
	var v5481 int64
	_ = v5481
	var v5485 int64
	_ = v5485
	var v5487 int64
	_ = v5487
	var v5488 int64
	_ = v5488
	var v5492 int64
	_ = v5492
	var v5494 int64
	_ = v5494
	var v5495 int64
	_ = v5495
	var v5500 int32
	_ = v5500
	var v5505 int32
	_ = v5505
	var v5506 int64
	_ = v5506
	var v5508 int64
	_ = v5508
	var v5509 int64
	_ = v5509
	var v5513 int64
	_ = v5513
	var v5515 int64
	_ = v5515
	var v5516 int64
	_ = v5516
	var v5520 int64
	_ = v5520
	var v5522 int64
	_ = v5522
	var v5523 int64
	_ = v5523
	var v5527 int64
	_ = v5527
	var v5529 int64
	_ = v5529
	var v5530 int64
	_ = v5530
	var v5534 int64
	_ = v5534
	var v5536 int64
	_ = v5536
	var v5537 int64
	_ = v5537
	var v5541 int64
	_ = v5541
	var v5543 int64
	_ = v5543
	var v5544 int64
	_ = v5544
	var v5548 int64
	_ = v5548
	var v5550 int64
	_ = v5550
	var v5551 int64
	_ = v5551
	var v5555 int64
	_ = v5555
	var v5557 int64
	_ = v5557
	var v5558 int64
	_ = v5558
	var v5562 int64
	_ = v5562
	var v5564 int64
	_ = v5564
	var v5565 int64
	_ = v5565
	var v5569 int64
	_ = v5569
	var v5571 int64
	_ = v5571
	var v5572 int64
	_ = v5572
	var v5576 int64
	_ = v5576
	var v5578 int64
	_ = v5578
	var v5579 int64
	_ = v5579
	var v5583 int64
	_ = v5583
	var v5585 int64
	_ = v5585
	var v5586 int64
	_ = v5586
	var v5590 int64
	_ = v5590
	var v5592 int64
	_ = v5592
	var v5593 int64
	_ = v5593
	var v5597 int64
	_ = v5597
	var v5599 int64
	_ = v5599
	var v5600 int64
	_ = v5600
	var v5604 int64
	_ = v5604
	var v5606 int64
	_ = v5606
	var v5607 int64
	_ = v5607
	var v5611 int64
	_ = v5611
	var v5613 int64
	_ = v5613
	var v5614 int64
	_ = v5614
	var v5618 int64
	_ = v5618
	var v5619 int64
	_ = v5619
	var v5620 int64
	_ = v5620
	var v5621 int64
	_ = v5621
	var v5622 int64
	_ = v5622
	var v5623 int64
	_ = v5623
	var v5627 int32
	_ = v5627
	var v5631 int32
	_ = v5631
	var v5634 int32
	_ = v5634
	var v5635 int32
	_ = v5635
	var v5642 int32
	_ = v5642
	var v5644 int32
	_ = v5644
	var v5645 int32
	_ = v5645
	var v5646 int64
	_ = v5646
	var v5647 int32
	_ = v5647
	var v5652 int32
	_ = v5652
	var v5656 int32
	_ = v5656
	var v5657 int32
	_ = v5657
	var v5658 int32
	_ = v5658
	var v5659 int32
	_ = v5659
	var v5662 float64
	_ = v5662
	var v5667 float64
	_ = v5667
	var v5676 int32
	_ = v5676
	var v5677 float64
	_ = v5677
	var v5678 int64
	_ = v5678
	var v5679 int64
	_ = v5679
	var v5688 int32
	_ = v5688
	var v5689 int64
	_ = v5689
	var v5692 int32
	_ = v5692
	var v5699 int32
	_ = v5699
	var v5700 int64
	_ = v5700
	var v5701 int32
	_ = v5701
	var v5702 int32
	_ = v5702
	var v5708 int32
	_ = v5708
	var v5713 int32
	_ = v5713
	var v5714 int32
	_ = v5714
	var v5717 int32
	_ = v5717
	var v5718 int32
	_ = v5718
	var v5726 int32
	_ = v5726
	var v5729 int32
	_ = v5729
	var v5732 int32
	_ = v5732
	var v5733 int32
	_ = v5733
	var v5743 int32
	_ = v5743
	var v5746 int32
	_ = v5746
	var v5747 int64
	_ = v5747
	var v5750 float64
	_ = v5750
	var v5755 float64
	_ = v5755
	var v5759 int32
	_ = v5759
	var v5764 int32
	_ = v5764
	var v5765 int32
	_ = v5765
	var v5766 int32
	_ = v5766
	var v5767 int32
	_ = v5767
	var v5776 int32
	_ = v5776
	var v5777 int32
	_ = v5777
	var v5780 int32
	_ = v5780
	var v5782 int32
	_ = v5782
	var v5787 int32
	_ = v5787
	var v5788 int32
	_ = v5788
	var v5793 int32
	_ = v5793
	var v5794 int32
	_ = v5794
	var v5795 int32
	_ = v5795
	var v5796 int32
	_ = v5796
	var v5798 int32
	_ = v5798
	var v5799 int32
	_ = v5799
	var v5800 int64
	_ = v5800
	var v5803 float64
	_ = v5803
	var v5808 float64
	_ = v5808
	var v5816 int32
	_ = v5816
	var v5817 int32
	_ = v5817
	var v5818 int32
	_ = v5818
	var v5828 int32
	_ = v5828
	var v5831 int32
	_ = v5831
	var v5873 int32
	_ = v5873
	var v5874 int32
	_ = v5874
	var v5876 int32
	_ = v5876
	var v5878 int32
	_ = v5878
	var v5879 int64
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5881 int32
	_ = v5881
	var v5892 int32
	_ = v5892
	var v5893 int32
	_ = v5893
	var v5896 int32
	_ = v5896
	var v5899 int32
	_ = v5899
	var v5951 int32
	_ = v5951
	var v5953 int32
	_ = v5953
	var v5954 int64
	_ = v5954
	var v5965 int32
	_ = v5965
	var v5967 int32
	_ = v5967
	var v5971 int64
	_ = v5971
	var v5974 float64
	_ = v5974
	var v5978 int64
	_ = v5978
	var v5990 int32
	_ = v5990
	var v5991 int64
	_ = v5991
	var v5992 int64
	_ = v5992
	var v5994 int32
	_ = v5994
	var v5995 int32
	_ = v5995
	var v6002 float64
	_ = v6002
	var v6004 float64
	_ = v6004
	var v6010 float64
	_ = v6010
	var v6019 float64
	_ = v6019
	var v6020 float64
	_ = v6020
	var v6024 int32
	_ = v6024
	var v6029 int32
	_ = v6029
	var v6037 int32
	_ = v6037
	var v6038 int64
	_ = v6038
	var v6040 int64
	_ = v6040
	var v6042 int64
	_ = v6042
	var v6044 int64
	_ = v6044
	var v6050 int32
	_ = v6050
	var v6053 int32
	_ = v6053
	var v6054 int32
	_ = v6054
	var v6060 int32
	_ = v6060
	var v6063 int32
	_ = v6063
	var v6065 int32
	_ = v6065
	var v6066 int32
	_ = v6066
	var v6067 int32
	_ = v6067
	var v6071 int32
	_ = v6071
	var v6076 int32
	_ = v6076
	var v6077 int32
	_ = v6077
	var v6079 int32
	_ = v6079
	var v6129 int32
	_ = v6129
	var v6134 int32
	_ = v6134
	var v6182 int32
	_ = v6182
	var v6183 int32
	_ = v6183
	var v6185 int32
	_ = v6185
	var v6187 int32
	_ = v6187
	var v6189 int32
	_ = v6189
	var v6191 int32
	_ = v6191
	var v6193 int32
	_ = v6193
	var v6194 int32
	_ = v6194
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
	v165 = F_palloc0(m, int32(256))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v128)+220)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+224)) = v125
	base.MemoryFill(m, v128+int32(232), int32(0), int32(160))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v154 + v140
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v160 - v140
	goto L11
L14:
	;
	return
L15:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v169 = F_get_database_name(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+68)) = v169
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+68))
	v174 = F_get_namespace_name(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+72)) = v174
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v180 = F_pstrdup(m, v177+int32(4))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+96)) = uint8(v75)
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+92)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v165)+80)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v165)+76)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v52)+572)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v52)+568)) = int32(187)
	v191 = int32(_a_F_heap_vacuum_rel_2)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12])) = v52 + int32(564)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+564)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = l0
	v201 = v165 + int32(8)
	v203 = v165 + int32(4)
	F_vac_open_indexes(m, l0, int32(3), v201, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = l2
	if v103 == int32(0) {
		v319 = v4
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])) = uint8(v335)
	v337 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+24)) = uint8(v337)
	v339 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+22)) = uint16(v339)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v342 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+25)) = uint8(base.B2i32(v341 != v342))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	switch v345 - v342 {
	case 0:
		goto L31
	case 1:
		goto L30
	default:
		goto L29
	}
L21:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v209 <= int32(0) {
		v319 = v4
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v214 = F_palloc(m, v209<<(uint(int32(2))%32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v216 <= int32(0) {
		v319 = v214
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v222 = int32(0)
	goto L25
L25:
	;
	v270 = v222 << (uint(int32(2)) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v272+v270)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+48))
	v278 = F_pstrdup(m, v275+int32(4))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L14
	} else {
		goto L27
	}
L26:
	;
	v319 = v214
	goto L20
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214+v270))) = v278
	v282 = v222 + int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v282 < v283 {
		v222 = v282
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v352 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v165)+120)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v165)+112)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v165)+140)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v165)+148)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v165)+156)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v165)+164)) = int32(0)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v367 = F_palloc0(m, v364<<(uint(int32(2))%32))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L14
	} else {
		goto L32
	}
L30:
	;
	v350 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+22)) = uint8(v350)
	goto L29
L31:
	;
	v348 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+23)) = uint16(v348)
	goto L29
L32:
	;
	v369 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v165)+172)) = v369
	v371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+136)) = v371
	*(*int64)(unsafe.Add(mBase, uint32(v165)+128)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v165)+168)) = v367
	*(*int64)(unsafe.Add(mBase, uint32(v165)+180)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v165)+188)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v165)+196)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v165)+204)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v165)+212)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v165)+220)) = v371
	v389 = v165 + int32(28)
	v390 = F_vacuum_get_cutoffs(m, l0, l1, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)) = uint8(v390)
	v394 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+108)) = v394
	v397 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v397<<(uint(int32(2))%32))+uint32(_c_F_heap_vacuum_rel[14])))
	goto L35
L35:
	;
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+64)) = uint8(v401)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v400
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v165)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+56)) = v404
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v408 = v406 & int32(256)
	if v408 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v409 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)) = uint8(v409)
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+21)) = uint8(base.B2i32(v408 == int32(0)))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+248)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v165)+240)) = int64(4294967295)
	v418 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	if base.F64_eq(v418, float64(0)) != 0 {
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
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)))
	if v421 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	if base.Ui32(v422) < base.Ui32(int32(_a_F_heap_vacuum_rel_3)) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	if base.Ui32(int32(3)) <= base.Ui32(v425) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_visibilitymap_count(m, v451, v52+int32(960), v52+int32(528))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L14
	} else {
		goto L55
	}
L44:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v428))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v425)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	if v441 == int32(0) {
		goto L39
	} else {
		goto L52
	}
L47:
	;
	if v440 != 0 {
		goto L43
	} else {
		goto L51
	}
L48:
	;
	v440 = base.B2i32(base.Ui32(v425) < base.Ui32(v428))
	goto L47
L49:
	;
	goto L50
L50:
	;
	v440 = int32(base.Ui32(v425-v428) >> (uint(int32(31)) % 32))
	goto L47
L51:
	;
	goto L46
L52:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v165)+48))
	goto L53
L53:
	;
	if int32(base.Ui32(v441-v444)>>(uint(int32(31))%32)) == int32(0) {
		goto L39
	} else {
		goto L54
	}
L54:
	;
	goto L43
L55:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v52)+960))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v52)+528))
	v464 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_u(v458-v459), float64(0.2)))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+244)) = v464
	if v464 == int32(0) {
		goto L39
	} else {
		goto L56
	}
L56:
	;
	v469 = Fn13966(m, int64(32))
	mBase = m.M
	goto L57
L57:
	;
	v471 = v469 & int32(4095)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+240)) = v471
	v473 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	v476 = base.I32_trunc_sat_f64_u(base.F64_mul(v473, float64(4096)))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+248)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v165)+252)) = base.I32_trunc_sat_f32_u(base.F32_mul(base.F32_add(base.F32_mul(base.F32_convert_i32_u(v471), float32(-0.00024414062)), float32(1)), base.F32_convert_i32_u(v476)))
	goto L39
L58:
	;
	v520 = F_lazy_check_wraparound_failsafe(m, v165)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L14
	} else {
		goto L70
	}
L59:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)))
	v494 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	if v494 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v165)+68))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+520)) = v499
	*(*int64)(unsafe.Add(mBase, uint32(v52)+512)) = v498
	v505 = v491 & int32(1)
	if v505 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v506 = int32(_a_F_heap_vacuum_rel_4)
	goto L64
L63:
	;
	v506 = int32(_a_F_heap_vacuum_rel_5)
	goto L64
L64:
	;
	F_errmsg(m, v506, v52+int32(512))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	if v505 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v514 = int32(817)
	goto L68
L67:
	;
	v514 = int32(822)
	goto L68
L68:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), v514, int32(_a_F_heap_vacuum_rel_7))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	goto L58
L70:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[15]))
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[16]))
	if v523 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v528 = v523
	goto L73
L72:
	;
	v528 = v525
	goto L73
L73:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[4]))
	if v530 == int32(4) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v533 = v528
	goto L76
L75:
	;
	v533 = v525
	goto L76
L76:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v534 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v1293 = v165 + int32(60)
	v1295 = v165 + int32(56)
	v1297 = v165 + int32(172)
	v1301 = v165 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+100)) = v1291
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v165)+244))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	v1305 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+924)) = v1305
	v1308 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+920)) = v1308
	v1311 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[18]))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+912)) = v1311
	v1313 = base.I64_extend_i32_u(v1304)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+536)) = v1313
	*(*int64)(unsafe.Add(mBase, uint32(v52)+528)) = int64(1)
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v1318 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1317))))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+544)) = v1318
	goto L209
L78:
	;
	v1231 = F_palloc(m, int32(16))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L14
	} else {
		goto L205
	}
L79:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v537 < int32(2) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+23)))
	if v540 != int32(1) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+48))
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+118)))
	if v545 == int32(116) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if v1171 == int32(0) {
		goto L78
	} else {
		goto L203
	}
L83:
	;
	if v534 == int32(0) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+96)))
	if v571 != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	v552 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L14
	} else {
		goto L87
	}
L87:
	;
	if v552 == int32(0) {
		goto L82
	} else {
		goto L88
	}
L88:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+496)) = v556
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_8), v52+int32(496))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L14
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(3500), int32(_a_F_heap_vacuum_rel_9))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L14
	} else {
		goto L90
	}
L90:
	;
	goto L82
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v1120
	goto L82
L92:
	;
	v572 = int32(17)
	goto L94
L93:
	;
	v572 = int32(13)
	goto L94
L94:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v574 = F_palloc0(m, v537)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	if v537 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v737 = F_palloc0(m, int32(72))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L14
	} else {
		goto L123
	}
L97:
	;
	F_pfree(m, v574)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L14
	} else {
		goto L122
	}
L98:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[19])))
	if v579&int32(1) == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[20]))
	if v585 == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v597 = v4
	v606 = v4
	v610 = v4
	goto L101
L101:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v568+v606<<(uint(int32(2))%32))))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+204))
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+29)))
	if v642 == int32(0) {
		v662 = v597
		v663 = v610
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v662 < v663 {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	v665 = v606 + int32(1)
	if v665 != v537 {
		v597 = v662
		v606 = v665
		v610 = v663
		goto L101
	} else {
		goto L107
	}
L104:
	;
	v646 = F_RelationGetNumberOfBlocksInFork(m, v640, int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L14
	} else {
		goto L105
	}
L105:
	;
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[21]))
	if base.Ui32(v646) < base.Ui32(v649) {
		v662 = v597
		v663 = v610
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v652 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v606+v574))) = uint8(v652)
	v662 = v597 + base.B2i32(v642&int32(6) != int32(0))
	v663 = v642&v652 + v610
	goto L103
L107:
	;
	goto L102
L108:
	;
	v668 = v663
	goto L110
L109:
	;
	v668 = v662
	goto L110
L110:
	;
	v670 = v668 - int32(1)
	if v670 <= int32(0) {
		goto L97
	} else {
		goto L111
	}
L111:
	;
	if base.Ui32(v534) < base.Ui32(v670) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v674 = v534
	goto L114
L113:
	;
	v674 = v670
	goto L114
L114:
	;
	if int32(0) < v534 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v677 = v674
	goto L117
L116:
	;
	v677 = v670
	goto L117
L117:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[20]))
	if v677 < v679 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v681 = v677
	goto L120
L119:
	;
	v681 = v679
	goto L120
L120:
	;
	if int32(0) < v681 {
		goto L96
	} else {
		goto L121
	}
L121:
	;
	goto L97
L122:
	;
	v1120 = int32(0)
	goto L91
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737)+52)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v737)+36)) = v574
	*(*int32)(unsafe.Add(mBase, uint32(v737)+12)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v737)+8)) = v568
	*(*int32)(unsafe.Add(mBase, uint32(v737)+4)) = v543
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[22]))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v746)+72)) = v747 + int32(1)
	goto L124
L124:
	;
	v753 = F_CreateParallelContext(m, int32(_a_F_heap_vacuum_rel_10), int32(_a_F_heap_vacuum_rel_11), v681)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L14
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = v753
	v757 = F_mul_size(m, int32(48), v537)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L14
	} else {
		goto L126
	}
L126:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v753)+36))
	v764 = F_add_size(m, v759, (v757+int32(31))&int32(-32))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L14
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+36)) = v764
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v753)+40))
	v769 = F_add_size(m, v767, int32(1))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L14
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+40)) = v769
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v753)+36))
	v774 = F_add_size(m, v772, int32(96))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L14
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+36)) = v774
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v753)+40))
	v779 = F_add_size(m, v777, int32(1))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L14
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+40)) = v779
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v753)+36))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v785 = F_mul_size(m, int32(128), v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L14
	} else {
		goto L131
	}
L131:
	;
	v791 = F_add_size(m, v782, (v785+int32(31))&int32(-32))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+36)) = v791
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v753)+40))
	v796 = F_add_size(m, v794, int32(1))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+40)) = v796
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v753)+36))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v802 = F_mul_size(m, int32(32), v801)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L14
	} else {
		goto L134
	}
L134:
	;
	v808 = F_add_size(m, v799, (v802+int32(31))&int32(-32))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L14
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+36)) = v808
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v753)+40))
	v813 = F_add_size(m, v811, int32(1))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L14
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+40)) = v813
	v817 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	if v817 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v753)+36))
	v819 = F_strlen(m, v817)
	mBase = m.M
	v824 = F_add_size(m, v818, v819&int32(-32)+int32(32))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L14
	} else {
		goto L140
	}
L138:
	;
	v832 = v4
	goto L139
L139:
	;
	F_InitializeParallelDSM(m, v753)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L14
	} else {
		goto L142
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+36)) = v824
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v753)+40))
	v829 = F_add_size(m, v827, int32(1))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L14
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753)+40)) = v829
	v832 = v819
	goto L139
L142:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	v838 = F_shm_toc_allocate(m, v837, v757)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L14
	} else {
		goto L144
	}
L143:
	;
	v872 = int32(1)
	if v537 <= v872 {
		goto L154
	} else {
		goto L155
	}
L144:
	;
	if v757&int32(3)|(v838&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v757))) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if v757 == int32(0) {
		goto L143
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if v757 == int32(0) {
		goto L143
	} else {
		goto L153
	}
L148:
	;
	v852 = v757 + v838
	v854 = v838 + int32(4)
	if base.Ui32(v854) < base.Ui32(v852) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v856 = v852
	goto L151
L150:
	;
	v856 = v854
	goto L151
L151:
	;
	v861 = (v838^int32(-1)+v856)&int32(-4) + int32(4)
	if v861 == int32(0) {
		goto L143
	} else {
		goto L152
	}
L152:
	;
	base.MemoryFill(m, v838, int32(0), v861)
	goto L143
L153:
	;
	base.MemoryFill(m, v838, int32(0), v757)
	goto L143
L154:
	;
	v875 = v872
	goto L156
L155:
	;
	v875 = v537
	goto L156
L156:
	;
	v876 = int32(0)
	v883 = v876
	v896 = v876
	goto L157
L157:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v896+v574))))
	if v928 != int32(1) {
		v959 = v883
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	F_shm_toc_insert(m, v964, int64(5), v838)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L14
	} else {
		goto L169
	}
L159:
	;
	v962 = v896 + int32(1)
	if v962 != v875 {
		v883 = v959
		v896 = v962
		goto L157
	} else {
		goto L168
	}
L160:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v568+v896<<(uint(int32(2))%32))))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v934)+204))
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935)+27)))
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935)+29)))
	if v937&int32(1) != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v737)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+40)) = v940 + int32(1)
	goto L163
L162:
	;
	goto L163
L163:
	;
	if v937&int32(4) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v737)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+44)) = v946 + int32(1)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v950 = v936 + v883
	if v937&int32(2) == int32(0) {
		v959 = v950
		goto L159
	} else {
		goto L167
	}
L167:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v737)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+48)) = v955 + int32(1)
	v959 = v950
	goto L159
L168:
	;
	goto L158
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737)+20)) = v838
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	v971 = F_shm_toc_allocate(m, v969, int32(72))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L14
	} else {
		goto L170
	}
L170:
	;
	v973 = int32(0)
	base.MemoryFill(m, v971, v973, int32(72))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v543)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v971)+4)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v971))) = v976
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v981 == v973 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v971)+8)) = v986
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[16]))
	if int32(0) < v959 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v986 = int64(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v981)+392))
	v986 = v985
	goto L171
L175:
	;
	if v681 < v959 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v995 = v989
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v971)+28)) = v995
	v998 = v533 << (uint(int32(10)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v971)+56)) = v998
	v1000 = F_TidStoreCreateShared(m, v998)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L14
	} else {
		goto L181
	}
L178:
	;
	v993 = v681
	goto L180
L179:
	;
	v993 = v959
	goto L180
L180:
	;
	v994 = base.I32_div_s(v989, v993)
	v995 = v994
	goto L177
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737)+24)) = v1000
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+4))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1003)))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v971)+52)) = v1005
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+8))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+28))
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v971)+48)) = v1009
	if v573 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v1016 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v971)+36)) = v1016
	*(*int32)(unsafe.Add(mBase, uint32(v971)+40)) = v1016
	*(*int32)(unsafe.Add(mBase, uint32(v971)+32)) = v1015
	*(*int32)(unsafe.Add(mBase, uint32(v971)+44)) = v1016
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	F_shm_toc_insert(m, v1023, int64(1), v971)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L14
	} else {
		goto L188
	}
L185:
	;
	v1015 = int32(0)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v1015 = v1014
	goto L184
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737)+16)) = v971
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v1031 = F_mul_size(m, int32(128), v1030)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L14
	} else {
		goto L189
	}
L189:
	;
	v1033 = F_shm_toc_allocate(m, v1028, v1031)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L14
	} else {
		goto L190
	}
L190:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	F_shm_toc_insert(m, v1035, int64(3), v1033)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L14
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737)+28)) = v1033
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v1043 = F_mul_size(m, int32(32), v1042)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L14
	} else {
		goto L192
	}
L192:
	;
	v1045 = F_shm_toc_allocate(m, v1040, v1043)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L14
	} else {
		goto L193
	}
L193:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	F_shm_toc_insert(m, v1047, int64(4), v1045)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L14
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737)+32)) = v1045
	v1053 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	if v1053 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	v1056 = v832 + int32(1)
	v1057 = F_shm_toc_allocate(m, v1054, v1056)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L14
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v1120 = v737
	goto L91
L198:
	;
	if v1056 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	base.MemoryCopy(m, v1057, v1060, v1056)
	goto L201
L200:
	;
	goto L201
L201:
	;
	v1063 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1057+v832))) = uint8(v1063)
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	F_shm_toc_insert(m, v1065, int64(2), v1057)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L14
	} else {
		goto L202
	}
L202:
	;
	goto L197
L203:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v165+int32(104)))) = v1176 + int32(56)
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+24))
	goto L204
L204:
	;
	v1291 = v1180
	goto L77
L205:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1231)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1231))) = v533 << (uint(int32(10)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+104)) = v1231
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1231)))
	v1240 = F_TidStoreCreateLocal(m, v1239)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L14
	} else {
		goto L206
	}
L206:
	;
	v1291 = v1240
	goto L77
L207:
	;
	v1498 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+236)) = v1498
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+232)) = uint16(v1498)
	*(*int64)(unsafe.Add(mBase, uint32(v165)+224)) = int64(-1)
	v1505 = v165 + int32(88)
	v1507 = v52 + int32(996)
	v1508 = int32(1)
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v1514 = F_read_stream_begin_relation(m, v1508, v1509, v1510, v1498, int32(188), v165, v1508)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L14
	} else {
		goto L224
	}
L208:
	;
	goto L207
L209:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v1334 == int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v1338&int32(1) == int32(0) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v1343 = int32(_a_F_heap_vacuum_rel_1)
	v1345 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1346 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1345 + v1346
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1334)))
	*(*int32)(unsafe.Add(mBase, uint32(v1334))) = v1349 + v1346
	goto L213
L212:
	;
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1334)))
	v1480 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1334))) = v1479 + v1480
	v1483 = int32(_a_F_heap_vacuum_rel_1)
	v1485 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1485 - v1480
	goto L208
L213:
	;
	goto L215
L215:
	;
	goto L216
L216:
	;
	v1444 = int32(0)
	v1447 = v1305
	goto L221
L221:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(912)+v1447<<(uint(int32(2))%32))))
	v1457 = int32(3)
	v1463 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(528)+v1447<<(uint(v1457)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1334+int32(232)+v1456<<(uint(v1457)%32)))) = v1463
	v1465 = int32(1)
	v1468 = v1444 + v1465
	if v1468 != int32(3) {
		v1444 = v1468
		v1447 = v1447 + v1465
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
	v1522 = int32(0)
	v1525 = v4
	goto L225
L225:
	;
	v1566 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+908)) = v1566
	F_vacuum_delay_point(m, v1566)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L14
	} else {
		goto L227
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+84)) = int32(-1)
	v3106 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	if v3106 != 0 {
		goto L521
	} else {
		goto L522
	}
L227:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1301)))
	v1572 = int32(0)
	if base.B2i32(v1571 == v1572)|v1571&int32(_a_F_heap_vacuum_rel_12) == v1572 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1579 = F_lazy_check_wraparound_failsafe(m, v165)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L14
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v1582 = *(*int64)(unsafe.Add(mBase, uint32(v1581)+8))
	if v1582 <= int64(0) {
		v1643 = v1525
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	v1646 = F_read_stream_next_buffer(m, v1514, v52+int32(908))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L14
	} else {
		goto L246
	}
L233:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v165)+100))
	v1586 = F_TidStoreMemoryUsage(m, v1585)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L14
	} else {
		goto L234
	}
L234:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1588)))
	if base.Ui32(v1586) <= base.Ui32(v1589) {
		v1643 = v1525
		goto L232
	} else {
		goto L235
	}
L235:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	if v1591 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	F_ReleaseBuffer(m, v1591)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L14
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v1596 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+22)) = uint8(v1596)
	F_lazy_vacuum(m, v165)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
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
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_FreeSpaceMapVacuumRange(m, v1600, v1525, v1522+int32(1))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L14
	} else {
		goto L241
	}
L241:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v1609 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1643 = v1522
	goto L232
L243:
	;
	goto L242
L244:
	;
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v1613&int32(1) == int32(0) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v1618 = int32(_a_F_heap_vacuum_rel_1)
	v1620 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1621 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1620 + v1621
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1609)))
	*(*int32)(unsafe.Add(mBase, uint32(v1609))) = v1624 + v1621
	*(*int64)(unsafe.Add(mBase, uint32(v1609+int32(0))+232)) = int64(1)
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1609)))
	*(*int32)(unsafe.Add(mBase, uint32(v1609))) = v1632 + v1621
	v1638 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1638 - v1621
	goto L243
L246:
	;
	if v1646 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v52)+908))
	v1649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648))))
	F_CheckBufferIsPinnedOnce(m, v1646)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
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
	if v1646 < int32(0) {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	if v1646 < int32(0) {
		goto L256
	} else {
		goto L257
	}
L252:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[24]))
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1655+(v1646^int32(-1))<<(uint(int32(2))%32))))
	v1669 = v1661
	goto L251
L253:
	;
	goto L254
L254:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[25]))
	v1669 = v1663 + v1646<<(uint(int32(13))%32) + int32(-8192)
	goto L251
L255:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1301)))
	v1690 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1301))) = v1689 + v1690
	v1694 = v1649 & v1690
	if v1694 != 0 {
		goto L259
	} else {
		goto L260
	}
L256:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[26]))
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1673+(v1646^int32(-1))<<(uint(int32(6))%32))+16))
	v1688 = v1679
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[27]))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1681+v1646<<(uint(int32(6))%32)+int32(-64))+16))
	v1688 = v1687
	goto L255
L259:
	;
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v165)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+116)) = v1695 + int32(1)
	goto L261
L260:
	;
	goto L261
L261:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v1703 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+92)) = int32(1)
	v1738 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+88)) = uint16(v1738)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+84)) = v1688
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_visibilitymap_pin(m, v1741, v1688, v52+int32(924))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L14
	} else {
		goto L266
	}
L263:
	;
	goto L262
L264:
	;
	v1707 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v1707&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1712 = int32(_a_F_heap_vacuum_rel_1)
	v1714 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1715 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1714 + v1715
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1703)))
	*(*int32)(unsafe.Add(mBase, uint32(v1703))) = v1718 + v1715
	*(*int64)(unsafe.Add(mBase, uint32(v1703+int32(16))+232)) = base.I64_extend_i32_u(v1688)
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1703)))
	*(*int32)(unsafe.Add(mBase, uint32(v1703))) = v1726 + v1715
	v1732 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1732 - v1715
	goto L263
L266:
	;
	v1746 = F_ConditionalLockBufferForCleanup(m, v1646)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L14
	} else {
		goto L267
	}
L267:
	;
	if v1746 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	F_LockBuffer(m, v1646, int32(1))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L14
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+14)))
	if v1753 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L271:
	;
	goto L270
L272:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v3010 != 0 {
		goto L492
	} else {
		goto L493
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1608)) = v2430
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if v2475 != 0 {
		goto L388
	} else {
		goto L389
	}
L274:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1608)) = v1909
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1293)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1600)) = v1911
	v1918 = int32(base.Ui32(v1908+int32(_a_F_heap_vacuum_rel_13))>>(uint(int32(2))%32)) & int32(_a_F_heap_vacuum_rel_14)
	if v1918 != 0 {
		goto L330
	} else {
		goto L331
	}
L275:
	;
	if v1746 != 0 {
		v2430 = v1762
		goto L273
	} else {
		goto L325
	}
L276:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_RecordPageWithFreeSpace(m, v1905, v1688, v1902)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L14
	} else {
		goto L324
	}
L277:
	;
	F_UnlockReleaseBuffer(m, v1646)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L14
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	v1763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v1763) {
		goto L275
	} else {
		goto L283
	}
L280:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v1759 = F_GetRecordedFreeSpace(m, v1758, v1688)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L14
	} else {
		goto L281
	}
L281:
	;
	if v1759 != 0 {
		v1522 = v1688
		v1525 = v1643
		goto L225
	} else {
		goto L282
	}
L282:
	;
	v1902 = int32(_a_F_heap_vacuum_rel_15)
	goto L276
L283:
	;
	if v1746 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	F_LockBuffer(m, v1646, int32(0))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L14
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669)+10)))
	if v1778&int32(4) == int32(0) {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	F_LockBuffer(m, v1646, int32(2))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L14
	} else {
		goto L288
	}
L288:
	;
	v1774 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v1774) {
		v1908 = v1774
		goto L274
	} else {
		goto L289
	}
L289:
	;
	goto L286
L290:
	;
	v1783 = int32(_a_F_heap_vacuum_rel_1)
	v1785 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1785 + int32(1)
	F_MarkBufferDirty(m, v1646)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L14
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v1836 = int32(4)
	v1837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+14)))
	v1838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+12)))
	v1839 = v1837 - v1838
	if v1839 <= v1836 {
		goto L305
	} else {
		goto L306
	}
L293:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+48))
	v1793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1792)+118)))
	if v1793 != int32(112) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1808 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)))
	v1810 = v1808 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)) = uint16(v1810)
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v1816 = F_visibilitymap_set(m, v1812, v1688, v1646, int64(0), v1762, int32(0), int32(3))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L14
	} else {
		goto L303
	}
L295:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[28]))
	if v1797 <= int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+32))
	if v1800 != 0 {
		goto L294
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1669)+4))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1669)))
	if v1802|v1803 != 0 {
		goto L294
	} else {
		goto L301
	}
L299:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+40))
	if v1801 != 0 {
		goto L294
	} else {
		goto L300
	}
L300:
	;
	goto L298
L301:
	;
	F_log_newpage_buffer(m, v1646, int32(1))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L14
	} else {
		goto L302
	}
L302:
	;
	goto L294
L303:
	;
	v1818 = int32(_a_F_heap_vacuum_rel_1)
	v1820 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v1821 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v1820 - v1821
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v165)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+128)) = v1824 + v1821
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v165)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+132)) = v1828 + v1821
	goto L292
L304:
	;
	F_UnlockReleaseBuffer(m, v1646)
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L14
	} else {
		goto L323
	}
L305:
	;
	v1842 = v1836
	goto L307
L306:
	;
	v1842 = v1839
	goto L307
L307:
	;
	v1844 = v1842 - int32(4)
	if v1844 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1899 = int32(0)
	goto L304
L309:
	;
	goto L310
L310:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1838) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v1899 = v1844
	goto L304
L312:
	;
	v1855 = int32(base.Ui32(v1838+int32(_a_F_heap_vacuum_rel_13)) >> (uint(int32(2)) % 32))
	goto L314
L313:
	;
	v1855 = int32(0)
	goto L314
L314:
	;
	if base.Ui32(v1855&int32(_a_F_heap_vacuum_rel_14)) < base.Ui32(int32(291)) {
		goto L311
	} else {
		goto L315
	}
L315:
	;
	v1860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669)+10)))
	if v1860&int32(1) == int32(0) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1899 = int32(0)
	goto L304
L317:
	;
	goto L318
L318:
	;
	v1869 = int32(1)
	goto L319
L319:
	;
	v1878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669+int32(20)+v1869&int32(_a_F_heap_vacuum_rel_14)<<(uint(int32(2))%32))+1)))
	if v1878&int32(384) == int32(0) {
		goto L311
	} else {
		goto L321
	}
L320:
	;
	v1899 = int32(0)
	goto L304
L321:
	;
	v1884 = v1869 + int32(1)
	v1885 = int32(_a_F_heap_vacuum_rel_14)
	if base.Ui32(v1884&v1885) <= base.Ui32(v1855&v1885) {
		v1869 = v1884
		goto L319
	} else {
		goto L322
	}
L322:
	;
	goto L320
L323:
	;
	v1902 = v1899
	goto L276
L324:
	;
	v1522 = v1688
	v1525 = v1643
	goto L225
L325:
	;
	v1908 = v1763
	goto L274
L326:
	;
	v2413 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1505))) = uint16(v2413)
	F_LockBuffer(m, v1646, v2413)
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L14
	} else {
		goto L386
	}
L327:
	;
	v2390 = *(*int64)(unsafe.Add(mBase, uint32(v165)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+200)) = v2390 + v2379
	v2393 = *(*int64)(unsafe.Add(mBase, uint32(v165)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+208)) = v2393 + v2382
	v2396 = *(*int64)(unsafe.Add(mBase, uint32(v165)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+216)) = v2396 + base.I64_extend_i32_s(v2347)
	if int32(0) < v2347 {
		goto L380
	} else {
		goto L381
	}
L328:
	;
	if v2052 <= int32(0) {
		goto L358
	} else {
		goto L359
	}
L329:
	;
	v2126 = int32(0)
	v2127 = base.B2i32(v2126 < v2100)
	if v2126 < v2100 {
		goto L355
	} else {
		goto L356
	}
L330:
	;
	v1920 = int32(base.Ui32(v1688) >> (uint(int32(16)) % 32))
	v1923 = int32(0)
	v1931 = int32(1)
	v1935 = v1923
	v1936 = v1923
	v1939 = v1923
	v1952 = v1923
	v1965 = v1923
	goto L333
L331:
	;
	goto L332
L332:
	;
	v2068 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1505))) = uint16(v2068)
	v2071 = int64(0)
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v2076 != 0 {
		v2343 = v2068
		v2347 = v2068
		v2348 = v2068
		v2379 = v2071
		v2382 = v2071
		goto L327
	} else {
		goto L354
	}
L333:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1505))) = uint16(v1931)
	v1983 = v1669 + int32(20) + v1931&int32(_a_F_heap_vacuum_rel_14)<<(uint(int32(2))%32)
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1983)))
	switch int32(base.Ui32(v1984)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L337
	case 1:
		goto L336
	case 2:
		goto L338
	default:
		v2049 = v1935
		v2050 = v1936
		v2051 = v1939
		v2052 = v1952
		v2053 = v1965
		goto L335
	}
L334:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1600))
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2061 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1505))) = uint16(v2061)
	*(*int32)(unsafe.Add(mBase, uint32(v1295))) = v2060
	*(*int32)(unsafe.Add(mBase, uint32(v1293))) = v2059
	v2065 = base.I64_extend_i32_s(v2053)
	v2066 = base.I64_extend_i32_s(v2051)
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v2067 != 0 {
		goto L328
	} else {
		goto L353
	}
L335:
	;
	v2055 = v1931 + int32(1)
	if base.Ui32(v2055&int32(_a_F_heap_vacuum_rel_14)) <= base.Ui32(v1918) {
		v1931 = v2055
		v1935 = v2049
		v1936 = v2050
		v1939 = v2051
		v1952 = v2052
		v1965 = v2053
		goto L333
	} else {
		goto L352
	}
L336:
	;
	v2049 = v1935
	v2050 = int32(1)
	v2051 = v1939
	v2052 = v1952
	v2053 = v1965
	goto L335
L337:
	;
	v2006 = F_heap_tuple_should_freeze(m, v1669+v1984&int32(_a_F_heap_vacuum_rel_16), v389, v52+int32(1608), v52+int32(1600))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L14
	} else {
		goto L339
	}
L338:
	;
	v1993 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(960)+v1952<<(uint(v1993)%32)))) = uint16(v1931)
	v2049 = v1935
	v2050 = v1936
	v2051 = v1939
	v2052 = v1952 + v1993
	v2053 = v1965
	goto L335
L339:
	;
	if v2006 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)))
	if v2008 != 0 {
		goto L326
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+936)) = uint16(v1931)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+934)) = uint16(v1688)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+932)) = uint16(v1920)
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v1983)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+928)) = int32(base.Ui32(v2012) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+944)) = v1669 + v2012&int32(_a_F_heap_vacuum_rel_16)
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v2020)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+940)) = v2021
	v2023 = int32(1)
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v2027 = F_HeapTupleSatisfiesVacuum(m, v52+int32(928), v2026, v1646)
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
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
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L14
	} else {
		goto L349
	}
L345:
	;
	v2049 = v1935
	v2050 = v2023
	v2051 = v1939
	v2052 = v1952
	v2053 = v1965 + int32(1)
	goto L335
L346:
	;
	v2049 = v1935 + int32(1)
	v2050 = v2023
	v2051 = v1939
	v2052 = v1952
	v2053 = v1965
	goto L335
L347:
	;
	v2049 = v1935
	v2050 = v2023
	v2051 = v1939 + int32(1)
	v2052 = v1952
	v2053 = v1965
	goto L335
L348:
	;
	switch v2027 {
	case 0:
		goto L346
	case 1, 4:
		goto L347
	case 2:
		goto L345
	case 3:
		v2049 = v1935
		v2050 = v2023
		v2051 = v1939
		v2052 = v1952
		v2053 = v1965
		goto L335
	default:
		goto L344
	}
L349:
	;
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_17), int32(0))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L14
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(2369), int32(_a_F_heap_vacuum_rel_18))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
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
	v2079 = v2050
	v2083 = v2049
	v2100 = v2052
	v2115 = v2066
	v2118 = v2065
	goto L329
L354:
	;
	v2079 = v2068
	v2083 = v2068
	v2100 = v2068
	v2115 = v2071
	v2118 = v2071
	goto L329
L355:
	;
	v2130 = v2100
	goto L357
L356:
	;
	v2130 = v2126
	goto L357
L357:
	;
	v2343 = v2127
	v2347 = v2130 + v2083
	v2348 = v2079 | v2127
	v2379 = v2115
	v2382 = v2118
	goto L327
L358:
	;
	v2343 = int32(0)
	v2347 = v2049
	v2348 = v2050
	v2379 = v2066
	v2382 = v2065
	goto L327
L359:
	;
	goto L360
L360:
	;
	v2135 = int32(1)
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v165)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+140)) = v2136 + v2135
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1584)) = int64(25769803783)
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v165)+100))
	F_TidStoreSetBlockOffsets(m, v2142, v1688, v52+int32(960), v2052)
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L14
	} else {
		goto L361
	}
L361:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v2148 = base.I64_extend_i32_u(v2052)
	v2149 = *(*int64)(unsafe.Add(mBase, uint32(v2147)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2147)+8)) = v2148 + v2149
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v2153 = *(*int64)(unsafe.Add(mBase, uint32(v2152)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+928)) = v2153
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v165)+100))
	v2156 = F_TidStoreMemoryUsage(m, v2155)
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L14
	} else {
		goto L362
	}
L362:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+936)) = base.I64_extend_i32_u(v2156)
	goto L365
L363:
	;
	v2338 = *(*int64)(unsafe.Add(mBase, uint32(v165)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+192)) = v2338 + v2148
	v2343 = v2135
	v2347 = v2049
	v2348 = v2050
	v2379 = v2066
	v2382 = v2065
	goto L327
L364:
	;
	goto L363
L365:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v2174 == int32(0) {
		goto L364
	} else {
		goto L366
	}
L366:
	;
	v2178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v2178&int32(1) == int32(0) {
		goto L364
	} else {
		goto L367
	}
L367:
	;
	v2183 = int32(_a_F_heap_vacuum_rel_1)
	v2185 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v2186 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2185 + v2186
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v2174)))
	*(*int32)(unsafe.Add(mBase, uint32(v2174))) = v2189 + v2186
	goto L369
L368:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v2174)))
	v2320 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2174))) = v2319 + v2320
	v2323 = int32(_a_F_heap_vacuum_rel_1)
	v2325 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2325 - v2320
	goto L364
L369:
	;
	goto L371
L371:
	;
	goto L372
L372:
	;
	v2284 = int32(0)
	v2287 = int32(0)
	goto L377
L377:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1584)+v2287<<(uint(int32(2))%32))))
	v2297 = int32(3)
	v2303 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(928)+v2287<<(uint(v2297)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2174+int32(232)+v2296<<(uint(v2297)%32)))) = v2303
	v2305 = int32(1)
	v2308 = v2284 + v2305
	if v2308 != int32(2) {
		v2284 = v2308
		v2287 = v2287 + v2305
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
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v165)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+144)) = v2402 + int32(1)
	goto L382
L381:
	;
	goto L382
L382:
	;
	if v2348&int32(1) != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+148)) = v1688 + int32(1)
	goto L385
L384:
	;
	goto L385
L385:
	;
	v2411 = int32(0)
	v2963 = v2343
	v2967 = v2411
	v2970 = v2411
	goto L272
L386:
	;
	F_LockBufferForCleanup(m, v1646)
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L14
	} else {
		goto L387
	}
L387:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v52)+924))
	v2430 = v2420
	goto L273
L388:
	;
	v2476 = int32(2)
	goto L390
L389:
	;
	v2476 = int32(3)
	goto L390
L390:
	;
	F_heap_page_prune_and_freeze(m, v2471, v1646, v2472, v2476, v389, v52+int32(960), int32(1), v1505, v1295, v1293)
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L14
	} else {
		goto L391
	}
L391:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v52)+968))
	if int32(0) < v2482 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v165)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+124)) = v2485 + int32(1)
	goto L394
L393:
	;
	goto L394
L394:
	;
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	if int32(0) < v2489 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v165)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+140)) = v2492 + int32(1)
	F_pg_qsort(m, v1507, v2489, int32(2), int32(189))
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L14
	} else {
		goto L398
	}
L396:
	;
	v2699 = v2489
	v2701 = v2482
	goto L397
L397:
	;
	v2702 = *(*int64)(unsafe.Add(mBase, uint32(v165)+176))
	v2703 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+960)))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+176)) = v2702 + v2703
	v2706 = *(*int64)(unsafe.Add(mBase, uint32(v165)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+184)) = v2706 + base.I64_extend_i32_s(v2701)
	v2710 = *(*int64)(unsafe.Add(mBase, uint32(v165)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+192)) = v2710 + base.I64_extend_i32_s(v2699)
	v2714 = *(*int64)(unsafe.Add(mBase, uint32(v165)+200))
	v2715 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+972)))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+200)) = v2714 + v2715
	v2718 = *(*int64)(unsafe.Add(mBase, uint32(v165)+208))
	v2719 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+976)))
	*(*int64)(unsafe.Add(mBase, uint32(v165)+208)) = v2718 + v2719
	v2722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+988)))
	if v2722 == int32(1) {
		goto L418
	} else {
		goto L419
	}
L398:
	;
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1584)) = int64(25769803783)
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v165)+100))
	F_TidStoreSetBlockOffsets(m, v2503, v1688, v1507, v2500)
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L14
	} else {
		goto L399
	}
L399:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v2507 = *(*int64)(unsafe.Add(mBase, uint32(v2506)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2506)+8)) = v2507 + base.I64_extend_i32_s(v2500)
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v2512 = *(*int64)(unsafe.Add(mBase, uint32(v2511)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+928)) = v2512
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v165)+100))
	v2515 = F_TidStoreMemoryUsage(m, v2514)
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L14
	} else {
		goto L400
	}
L400:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+936)) = base.I64_extend_i32_u(v2515)
	goto L403
L401:
	;
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v52)+968))
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	v2699 = v2698
	v2701 = v2697
	goto L397
L402:
	;
	goto L401
L403:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v2533 == int32(0) {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v2537 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v2537&int32(1) == int32(0) {
		goto L402
	} else {
		goto L405
	}
L405:
	;
	v2542 = int32(_a_F_heap_vacuum_rel_1)
	v2544 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v2545 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2544 + v2545
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v2533)))
	*(*int32)(unsafe.Add(mBase, uint32(v2533))) = v2548 + v2545
	goto L407
L406:
	;
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v2533)))
	v2679 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2533))) = v2678 + v2679
	v2682 = int32(_a_F_heap_vacuum_rel_1)
	v2684 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v2684 - v2679
	goto L402
L407:
	;
	goto L409
L409:
	;
	goto L410
L410:
	;
	v2643 = int32(0)
	v2646 = int32(0)
	goto L415
L415:
	;
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1584)+v2646<<(uint(int32(2))%32))))
	v2656 = int32(3)
	v2662 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(928)+v2646<<(uint(v2656)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2533+int32(232)+v2655<<(uint(v2656)%32)))) = v2662
	v2664 = int32(1)
	v2667 = v2643 + v2664
	if v2667 != int32(2) {
		v2643 = v2667
		v2646 = v2646 + v2664
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
	*(*int32)(unsafe.Add(mBase, uint32(v165)+148)) = v1688 + int32(1)
	goto L420
L419:
	;
	goto L420
L420:
	;
	v2729 = v1649 & int32(2)
	if v2729 != 0 {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v2910 = int32(0)
	v2911 = base.B2i32(v2910 < v2699)
	v2912 = *(*int32)(unsafe.Add(mBase, uint32(v52)+960))
	v2914 = base.B2i32(v2910 < v2912)
	v2915 = int32(1)
	if v1694 == v2910 {
		v2963 = v2911
		v2967 = v2914
		v2970 = v2915
		goto L272
	} else {
		goto L473
	}
L422:
	;
	if v2729 == int32(0) {
		v2791 = v2699
		goto L438
	} else {
		goto L439
	}
L423:
	;
	v2730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+980)))
	if v2730&int32(1) == int32(0) {
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v2735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	v2736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)))
	v2738 = v2736 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)) = uint16(v2738)
	F_MarkBufferDirty(m, v1646)
	mBase = m.M
	v2741 = m.ExcPending
	if v2741 != 0 {
		goto L14
	} else {
		goto L425
	}
L425:
	;
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(v52)+984))
	if v2735 != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2747 = int32(3)
	goto L428
L427:
	;
	v2747 = int32(1)
	goto L428
L428:
	;
	v2748 = F_visibilitymap_set(m, v2742, v1688, v1646, int64(0), v2430, v2744, v2747)
	mBase = m.M
	v2749 = m.ExcPending
	if v2749 != 0 {
		goto L14
	} else {
		goto L429
	}
L429:
	;
	if v2748&int32(1) == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v165)+128))
	v2755 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+128)) = v2754 + v2755
	v2759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	if v2759 != v2755 {
		v2909 = int32(0)
		goto L421
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	v2767 = int32(0)
	if v2748&int32(2) != 0 {
		v2909 = v2767
		goto L421
	} else {
		goto L434
	}
L433:
	;
	v2762 = int32(1)
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v165)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+132)) = v2763 + v2762
	v2909 = v2762
	goto L421
L434:
	;
	v2770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	if v2770&int32(1) == int32(0) {
		v2909 = v2767
		goto L421
	} else {
		goto L435
	}
L435:
	;
	v2775 = int32(1)
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v165)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+136)) = v2776 + v2775
	v2909 = v2775
	goto L421
L436:
	;
	v2853 = int32(0)
	if v2729 == v2853 {
		v2909 = v2853
		goto L421
	} else {
		goto L460
	}
L437:
	;
	v2831 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L14
	} else {
		goto L453
	}
L438:
	;
	if v2791 <= int32(0) {
		goto L436
	} else {
		goto L443
	}
L439:
	;
	v2782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669)+10)))
	if v2782&int32(4) != 0 {
		v2791 = v2699
		goto L438
	} else {
		goto L440
	}
L440:
	;
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v2788 = F_visibilitymap_get_status(m, v2785, v1688, v52+int32(1608))
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L14
	} else {
		goto L441
	}
L441:
	;
	if v2788 != 0 {
		goto L437
	} else {
		goto L442
	}
L442:
	;
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v52)+992))
	v2791 = v2790
	goto L438
L443:
	;
	v2794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669)+10)))
	if v2794&int32(4) == int32(0) {
		goto L436
	} else {
		goto L444
	}
L444:
	;
	v2801 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v2802 = m.ExcPending
	if v2802 != 0 {
		goto L14
	} else {
		goto L445
	}
L445:
	;
	if v2801 != 0 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+468)) = v1688
	*(*int32)(unsafe.Add(mBase, uint32(v52)+464)) = v2803
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_19), v52+int32(464))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L14
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	v2817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)))
	v2819 = v2817 & int32(_a_F_heap_vacuum_rel_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)) = uint16(v2819)
	F_MarkBufferDirty(m, v1646)
	mBase = m.M
	v2822 = m.ExcPending
	if v2822 != 0 {
		goto L14
	} else {
		goto L451
	}
L449:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(2148), int32(_a_F_heap_vacuum_rel_21))
	mBase = m.M
	v2815 = m.ExcPending
	if v2815 != 0 {
		goto L14
	} else {
		goto L450
	}
L450:
	;
	goto L448
L451:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2826 = F_visibilitymap_clear(m, v1688, v2824, int32(3))
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		goto L14
	} else {
		goto L452
	}
L452:
	;
	v2909 = int32(0)
	goto L421
L453:
	;
	if v2831 != 0 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+484)) = v1688
	*(*int32)(unsafe.Add(mBase, uint32(v52)+480)) = v2833
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_22), v52+int32(480))
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L14
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2850 = F_visibilitymap_clear(m, v1688, v2848, int32(3))
	mBase = m.M
	v2851 = m.ExcPending
	if v2851 != 0 {
		goto L14
	} else {
		goto L459
	}
L457:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(2126), int32(_a_F_heap_vacuum_rel_21))
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L14
	} else {
		goto L458
	}
L458:
	;
	goto L456
L459:
	;
	v2909 = int32(0)
	goto L421
L460:
	;
	v2856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+980)))
	if v2856&int32(1) == int32(0) {
		v2909 = v2853
		goto L421
	} else {
		goto L461
	}
L461:
	;
	v2861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+981)))
	if v2861&int32(1) == int32(0) {
		v2909 = v2853
		goto L421
	} else {
		goto L462
	}
L462:
	;
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v2869 = F_visibilitymap_get_status(m, v2866, v1688, v52+int32(1608))
	mBase = m.M
	v2870 = m.ExcPending
	if v2870 != 0 {
		goto L14
	} else {
		goto L463
	}
L463:
	;
	if v2869&int32(2) != 0 {
		v2909 = v2853
		goto L421
	} else {
		goto L464
	}
L464:
	;
	v2873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)))
	if v2873&int32(4) == int32(0) {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v2879 = v2873 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1669)+10)) = uint16(v2879)
	F_MarkBufferDirty(m, v1646)
	mBase = m.M
	v2882 = m.ExcPending
	if v2882 != 0 {
		goto L14
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	v2888 = F_visibilitymap_set(m, v2883, v1688, v1646, int64(0), v2885, int32(0), int32(3))
	mBase = m.M
	v2889 = m.ExcPending
	if v2889 != 0 {
		goto L14
	} else {
		goto L469
	}
L468:
	;
	goto L467
L469:
	;
	if v2888&int32(1) == int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2894 = int32(1)
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v165)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+128)) = v2895 + v2894
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v165)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+132)) = v2899 + v2894
	v2909 = v2894
	goto L421
L471:
	;
	goto L472
L472:
	;
	v2903 = int32(1)
	v2904 = *(*int32)(unsafe.Add(mBase, uint32(v165)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+136)) = v2904 + v2903
	v2909 = v2903
	goto L421
L473:
	;
	if v2909 != 0 {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v165)+244))
	if v2918 != 0 {
		goto L477
	} else {
		goto L478
	}
L475:
	;
	goto L476
L476:
	;
	v2955 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
	if v2955 == int32(0) {
		v2963 = v2911
		v2967 = v2914
		v2970 = v2915
		goto L272
	} else {
		goto L490
	}
L477:
	;
	v2920 = v2918 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+244)) = v2920
	if v2920 != 0 {
		v2963 = v2911
		v2967 = v2914
		v2970 = v2915
		goto L272
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v165)+248))
	if v2923 == int32(0) {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	goto L479
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+240)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v165)+248)) = int64(0)
	v2963 = v2911
	v2967 = v2914
	v2970 = v2915
	goto L272
L482:
	;
	v2928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+96)))
	if v2928 != 0 {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v2929 = int32(17)
	goto L485
L484:
	;
	v2929 = int32(13)
	goto L485
L485:
	;
	v2931 = F_errstart(m, v2929, int32(0))
	mBase = m.M
	v2932 = m.ExcPending
	if v2932 != 0 {
		goto L14
	} else {
		goto L486
	}
L486:
	;
	if v2931 == int32(0) {
		goto L481
	} else {
		goto L487
	}
L487:
	;
	v2935 = *(*int64)(unsafe.Add(mBase, uint32(v165)+68))
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+460)) = v2936
	*(*int64)(unsafe.Add(mBase, uint32(v52)+452)) = v2935
	*(*int32)(unsafe.Add(mBase, uint32(v52)+448)) = v1303
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_23), v52+int32(448))
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L14
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(1435), int32(_a_F_heap_vacuum_rel_24))
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L14
	} else {
		goto L489
	}
L489:
	;
	goto L481
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+252)) = v2955 - int32(1)
	v2963 = v2911
	v2967 = v2914
	v2970 = v2915
	goto L272
L491:
	;
	F_UnlockReleaseBuffer(m, v1646)
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L14
	} else {
		goto L520
	}
L492:
	;
	v3011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+23)))
	if v3011&v2963&int32(1) != 0 {
		goto L491
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	v3018 = int32(4)
	v3019 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+14)))
	v3020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669)+12)))
	v3021 = v3019 - v3020
	if v3021 <= v3018 {
		goto L497
	} else {
		goto L498
	}
L495:
	;
	goto L494
L496:
	;
	F_UnlockReleaseBuffer(m, v1646)
	mBase = m.M
	v3083 = m.ExcPending
	if v3083 != 0 {
		goto L14
	} else {
		goto L515
	}
L497:
	;
	v3024 = v3018
	goto L499
L498:
	;
	v3024 = v3021
	goto L499
L499:
	;
	v3026 = v3024 - int32(4)
	if v3026 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v3081 = int32(0)
	goto L496
L501:
	;
	goto L502
L502:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v3020) {
		goto L504
	} else {
		goto L505
	}
L503:
	;
	v3081 = v3026
	goto L496
L504:
	;
	v3037 = int32(base.Ui32(v3020+int32(_a_F_heap_vacuum_rel_13)) >> (uint(int32(2)) % 32))
	goto L506
L505:
	;
	v3037 = int32(0)
	goto L506
L506:
	;
	if base.Ui32(v3037&int32(_a_F_heap_vacuum_rel_14)) < base.Ui32(int32(291)) {
		goto L503
	} else {
		goto L507
	}
L507:
	;
	v3042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669)+10)))
	if v3042&int32(1) == int32(0) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3081 = int32(0)
	goto L496
L509:
	;
	goto L510
L510:
	;
	v3051 = int32(1)
	goto L511
L511:
	;
	v3060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1669+int32(20)+v3051&int32(_a_F_heap_vacuum_rel_14)<<(uint(int32(2))%32))+1)))
	if v3060&int32(384) == int32(0) {
		goto L503
	} else {
		goto L513
	}
L512:
	;
	v3081 = int32(0)
	goto L496
L513:
	;
	v3066 = v3051 + int32(1)
	v3067 = int32(_a_F_heap_vacuum_rel_14)
	if base.Ui32(v3066&v3067) <= base.Ui32(v3037&v3067) {
		v3051 = v3066
		goto L511
	} else {
		goto L514
	}
L514:
	;
	goto L512
L515:
	;
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_RecordPageWithFreeSpace(m, v3084, v1688, v3081)
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L14
	} else {
		goto L516
	}
L516:
	;
	if v2970 == int32(0) {
		v1522 = v1688
		v1525 = v1643
		goto L225
	} else {
		goto L517
	}
L517:
	;
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v3090 = int32(0)
	if base.B2i32(base.B2i32(v3089 == v3090)&v2967 == v3090)|base.B2i32(base.Ui32(v1688-v1643) < base.Ui32(int32(_a_F_heap_vacuum_rel_25))) != 0 {
		v1522 = v1688
		v1525 = v1643
		goto L225
	} else {
		goto L518
	}
L518:
	;
	v3099 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_FreeSpaceMapVacuumRange(m, v3099, v1643, v1688)
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L14
	} else {
		goto L519
	}
L519:
	;
	v1522 = v1688
	v1525 = v1688
	goto L225
L520:
	;
	v1522 = v1688
	v1525 = v1643
	goto L225
L521:
	;
	F_ReleaseBuffer(m, v3106)
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L14
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	v3112 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3112 == int32(0) {
		goto L526
	} else {
		goto L527
	}
L524:
	;
	goto L523
L525:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	v3147 = *(*int64)(unsafe.Add(mBase, uint32(v165)+200))
	v3148 = base.F64_convert_i64_s(v3147)
	if base.Ui32(v3146) < base.Ui32(v1304) {
		goto L530
	} else {
		goto L531
	}
L526:
	;
	goto L525
L527:
	;
	v3116 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3116&int32(1) == int32(0) {
		goto L526
	} else {
		goto L528
	}
L528:
	;
	v3121 = int32(_a_F_heap_vacuum_rel_1)
	v3123 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3124 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3123 + v3124
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(v3112)))
	*(*int32)(unsafe.Add(mBase, uint32(v3112))) = v3127 + v3124
	*(*int64)(unsafe.Add(mBase, uint32(v3112+int32(16))+232)) = v1313
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v3112)))
	*(*int32)(unsafe.Add(mBase, uint32(v3112))) = v3135 + v3124
	v3141 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3141 - v3124
	goto L526
L529:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v165)+160)) = v3198
	v3200 = float64(0)
	if base.F64_gt(v3198, v3200) != 0 {
		goto L548
	} else {
		goto L549
	}
L530:
	;
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+48))
	v3154 = *(*float32)(unsafe.Add(mBase, uint32(v3153)+100))
	v3155 = base.F64_promote_f32(v3154)
	v3156 = *(*int32)(unsafe.Add(mBase, uint32(v3153)+96))
	if v1304 == v3156 {
		goto L534
	} else {
		goto L535
	}
L531:
	;
	v3193 = v3148
	goto L532
L532:
	;
	v3198 = v3193
	goto L529
L533:
	;
	v3169 = base.F64_convert_i32_u(v1304)
	if v3156 != 0 {
		goto L542
	} else {
		goto L543
	}
L534:
	;
	if base.Ui32(v3146) < base.Ui32(int32(2)) {
		goto L537
	} else {
		goto L538
	}
L535:
	;
	goto L536
L536:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v3146) {
		goto L533
	} else {
		goto L541
	}
L537:
	;
	v3198 = v3155
	goto L529
L538:
	;
	goto L539
L539:
	;
	if base.F64_lt(base.F64_convert_i32_u(v3146), base.F64_mul(base.F64_convert_i32_u(v1304), float64(0.02))) == int32(0) {
		goto L533
	} else {
		goto L540
	}
L540:
	;
	v3198 = v3155
	goto L529
L541:
	;
	v3198 = v3155
	goto L529
L542:
	;
	v3173 = base.F32_lt(v3154, float32(0))
	goto L544
L543:
	;
	v3173 = int32(1)
	goto L544
L544:
	;
	if v3173 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v3198 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v3148, base.F64_convert_i32_u(v3146)), v3169), float64(0.5)))
	goto L529
L546:
	;
	goto L547
L547:
	;
	v3193 = base.F64_floor(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(v3155, base.F64_convert_i32_u(v3156)), base.F64_sub(v3169, base.F64_convert_i32_u(v3146))), v3148), float64(0.5)))
	goto L532
L548:
	;
	v3203 = v3198
	goto L550
L549:
	;
	v3203 = v3200
	goto L550
L550:
	;
	v3204 = *(*int64)(unsafe.Add(mBase, uint32(v165)+208))
	v3207 = *(*int64)(unsafe.Add(mBase, uint32(v165)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v165)+152)) = base.F64_add(base.F64_add(v3203, base.F64_convert_i64_s(v3204)), base.F64_convert_i64_s(v3207))
	F_read_stream_end(m, v1514)
	mBase = m.M
	v3212 = m.ExcPending
	if v3212 != 0 {
		goto L14
	} else {
		goto L551
	}
L551:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v165)+104))
	v3214 = *(*int64)(unsafe.Add(mBase, uint32(v3213)+8))
	if int64(0) < v3214 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	F_lazy_vacuum(m, v165)
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L14
	} else {
		goto L555
	}
L553:
	;
	goto L554
L554:
	;
	if base.Ui32(v1643) < base.Ui32(v1304) {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	goto L554
L556:
	;
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_FreeSpaceMapVacuumRange(m, v3220, v1643, v1304)
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L14
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v3226 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3226 == int32(0) {
		goto L561
	} else {
		goto L562
	}
L559:
	;
	goto L558
L560:
	;
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v3259 <= int32(0) {
		goto L564
	} else {
		goto L565
	}
L561:
	;
	goto L560
L562:
	;
	v3230 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3230&int32(1) == int32(0) {
		goto L561
	} else {
		goto L563
	}
L563:
	;
	v3235 = int32(_a_F_heap_vacuum_rel_1)
	v3237 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3238 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3237 + v3238
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v3226)))
	*(*int32)(unsafe.Add(mBase, uint32(v3226))) = v3241 + v3238
	*(*int64)(unsafe.Add(mBase, uint32(v3226+int32(24))+232)) = v1313
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(v3226)))
	*(*int32)(unsafe.Add(mBase, uint32(v3226))) = v3249 + v3238
	v3255 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3255 - v3238
	goto L561
L564:
	;
	v3892 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if v3892 != 0 {
		goto L617
	} else {
		goto L618
	}
L565:
	;
	v3262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+24)))
	if v3262 != int32(1) {
		goto L564
	} else {
		goto L566
	}
L566:
	;
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	v3267 = *(*float64)(unsafe.Add(mBase, uint32(v165)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1608)) = int64(34359738368)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1600)) = int64(38654705672)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+936)) = base.I64_extend_i32_u(v3259)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+928)) = int64(4)
	v3276 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1592)) = v3276
	*(*int64)(unsafe.Add(mBase, uint32(v52)+1584)) = v3276
	goto L569
L567:
	;
	v3458 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if v3458 == int32(0) {
		goto L585
	} else {
		goto L586
	}
L568:
	;
	goto L567
L569:
	;
	v3294 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3294 == int32(0) {
		goto L568
	} else {
		goto L570
	}
L570:
	;
	v3298 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3298&int32(1) == int32(0) {
		goto L568
	} else {
		goto L571
	}
L571:
	;
	v3303 = int32(_a_F_heap_vacuum_rel_1)
	v3305 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3306 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3305 + v3306
	v3309 = *(*int32)(unsafe.Add(mBase, uint32(v3294)))
	*(*int32)(unsafe.Add(mBase, uint32(v3294))) = v3309 + v3306
	goto L573
L572:
	;
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(v3294)))
	v3440 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3294))) = v3439 + v3440
	v3443 = int32(_a_F_heap_vacuum_rel_1)
	v3445 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3445 - v3440
	goto L568
L573:
	;
	goto L575
L575:
	;
	goto L576
L576:
	;
	v3404 = int32(0)
	v3407 = int32(0)
	goto L581
L581:
	;
	v3416 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1608)+v3407<<(uint(int32(2))%32))))
	v3417 = int32(3)
	v3423 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(928)+v3407<<(uint(v3417)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v3294+int32(232)+v3416<<(uint(v3417)%32)))) = v3423
	v3425 = int32(1)
	v3428 = v3404 + v3425
	if v3428 != int32(2) {
		v3404 = v3428
		v3407 = v3407 + v3425
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
	v3461 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v3461 <= int32(0) {
		goto L584
	} else {
		goto L588
	}
L586:
	;
	goto L587
L587:
	;
	v3605 = *(*int32)(unsafe.Add(mBase, uint32(v1297)))
	v3606 = *(*int32)(unsafe.Add(mBase, uint32(v3458)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v3606)+16)) = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(v3267))
	v3610 = *(*int32)(unsafe.Add(mBase, uint32(v3458)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v3610)+24)) = uint8(base.B2i32(base.Ui32(v3266) < base.Ui32(v3265)))
	F_parallel_vacuum_process_all_indexes(m, v3458, v3605, int32(0))
	mBase = m.M
	v3615 = m.ExcPending
	if v3615 != 0 {
		goto L14
	} else {
		goto L599
	}
L588:
	;
	v3504 = int64(0)
	goto L589
L589:
	;
	v3517 = base.I32_wrap_i64(v3504) << (uint(int32(2)) % 32)
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v165)+168))
	v3520 = *(*int32)(unsafe.Add(mBase, uint32(v3517+v3518)))
	v3521 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v3523 = *(*int32)(unsafe.Add(mBase, uint32(v3521+v3517)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+960)) = v3523
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*float64)(unsafe.Add(mBase, uint32(v52)+976)) = v3267
	*(*int32)(unsafe.Add(mBase, uint32(v52)+972)) = int32(13)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+970)) = uint8(base.B2i32(base.Ui32(v3266) < base.Ui32(v3265)))
	v3530 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+968)) = uint16(v3530)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+964)) = v3525
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+984)) = v3533
	v3535 = *(*int32)(unsafe.Add(mBase, uint32(v3523)+48))
	v3538 = F_pstrdup(m, v3535+int32(4))
	mBase = m.M
	v3539 = m.ExcPending
	if v3539 != 0 {
		goto L14
	} else {
		goto L591
	}
L590:
	;
	goto L584
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+80)) = v3538
	v3541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+88)))
	v3542 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+88)) = uint16(v3542)
	v3544 = *(*int32)(unsafe.Add(mBase, uint32(v165)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+84)) = int32(-1)
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(v165)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+92)) = int32(4)
	v3552 = F_vac_cleanup_one_index(m, v52+int32(960), v3520)
	mBase = m.M
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L14
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+92)) = v3547
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+88)) = uint16(v3541)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+84)) = v3544
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v165)+80))
	F_pfree(m, v3557)
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L14
	} else {
		goto L593
	}
L593:
	;
	v3560 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+80)) = v3560
	v3562 = *(*int32)(unsafe.Add(mBase, uint32(v165)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v3562+v3517))) = v3552
	v3567 = v3504 + int64(1)
	v3570 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3570 == v3560 {
		goto L595
	} else {
		goto L596
	}
L594:
	;
	v3603 = int64(*(*int32)(unsafe.Add(mBase, uint32(v165)+8)))
	if v3567 < v3603 {
		v3504 = v3567
		goto L589
	} else {
		goto L598
	}
L595:
	;
	goto L594
L596:
	;
	v3574 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3574&int32(1) == int32(0) {
		goto L595
	} else {
		goto L597
	}
L597:
	;
	v3579 = int32(_a_F_heap_vacuum_rel_1)
	v3581 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3582 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3581 + v3582
	v3585 = *(*int32)(unsafe.Add(mBase, uint32(v3570)))
	*(*int32)(unsafe.Add(mBase, uint32(v3570))) = v3585 + v3582
	*(*int64)(unsafe.Add(mBase, uint32(v3570+int32(72))+232)) = v3567
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v3570)))
	*(*int32)(unsafe.Add(mBase, uint32(v3570))) = v3593 + v3582
	v3599 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3599 - v3582
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
	v3679 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v3679 == int32(0) {
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v3683 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v3683&int32(1) == int32(0) {
		goto L601
	} else {
		goto L604
	}
L604:
	;
	v3688 = int32(_a_F_heap_vacuum_rel_1)
	v3690 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v3691 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3690 + v3691
	v3694 = *(*int32)(unsafe.Add(mBase, uint32(v3679)))
	*(*int32)(unsafe.Add(mBase, uint32(v3679))) = v3694 + v3691
	goto L606
L605:
	;
	v3824 = *(*int32)(unsafe.Add(mBase, uint32(v3679)))
	v3825 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3679))) = v3824 + v3825
	v3828 = int32(_a_F_heap_vacuum_rel_1)
	v3830 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v3830 - v3825
	goto L601
L606:
	;
	goto L608
L608:
	;
	goto L609
L609:
	;
	v3789 = int32(0)
	v3792 = int32(0)
	goto L614
L614:
	;
	v3801 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(1600)+v3792<<(uint(int32(2))%32))))
	v3802 = int32(3)
	v3808 = *(*int64)(unsafe.Add(mBase, uint32(v52+int32(1584)+v3792<<(uint(v3802)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v3679+int32(232)+v3801<<(uint(v3802)%32)))) = v3808
	v3810 = int32(1)
	v3813 = v3789 + v3810
	if v3813 != int32(2) {
		v3789 = v3813
		v3792 = v3792 + v3810
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
	v3893 = *(*int32)(unsafe.Add(mBase, uint32(v165)+168))
	v3894 = int32(0)
	v3895 = *(*int32)(unsafe.Add(mBase, uint32(v3892)+12))
	if v3894 < v3895 {
		goto L620
	} else {
		goto L621
	}
L618:
	;
	goto L619
L619:
	;
	v4099 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v4100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+24)))
	v4103 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if base.B2i32(v4100 != int32(1))|base.B2i32(v4103 <= int32(0)) != 0 {
		goto L636
	} else {
		goto L637
	}
L620:
	;
	v3906 = v3894
	goto L623
L621:
	;
	goto L622
L622:
	;
	v4030 = *(*int32)(unsafe.Add(mBase, uint32(v3892)+24))
	F_TidStoreDestroy(m, v4030)
	mBase = m.M
	v4032 = m.ExcPending
	if v4032 != 0 {
		goto L14
	} else {
		goto L631
	}
L623:
	;
	v3947 = *(*int32)(unsafe.Add(mBase, uint32(v3892)+20))
	v3950 = v3947 + v3906*int32(48)
	v3951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3950)+5)))
	if v3951 == int32(1) {
		goto L626
	} else {
		goto L627
	}
L624:
	;
	goto L622
L625:
	;
	v3978 = v3906 + int32(1)
	v3979 = *(*int32)(unsafe.Add(mBase, uint32(v3892)+12))
	if v3978 < v3979 {
		v3906 = v3978
		goto L623
	} else {
		goto L630
	}
L626:
	;
	v3958 = F_palloc0(m, int32(40))
	mBase = m.M
	v3959 = m.ExcPending
	if v3959 != 0 {
		goto L14
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3893+v3906<<(uint(int32(2))%32)))) = int32(0)
	goto L625
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3893+v3906<<(uint(int32(2))%32)))) = v3958
	v3961 = *(*int64)(unsafe.Add(mBase, uint32(v3950)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3958)+32)) = v3961
	v3963 = *(*int64)(unsafe.Add(mBase, uint32(v3950)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v3958)+24)) = v3963
	v3965 = *(*int64)(unsafe.Add(mBase, uint32(v3950)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3958)+16)) = v3965
	v3967 = *(*int64)(unsafe.Add(mBase, uint32(v3950)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3958)+8)) = v3967
	v3969 = *(*int64)(unsafe.Add(mBase, uint32(v3950)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3958))) = v3969
	goto L625
L630:
	;
	goto L624
L631:
	;
	v4033 = *(*int32)(unsafe.Add(mBase, uint32(v3892)))
	F_DestroyParallelContext(m, v4033)
	mBase = m.M
	v4035 = m.ExcPending
	if v4035 != 0 {
		goto L14
	} else {
		goto L632
	}
L632:
	;
	v4038 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[22]))
	v4039 = *(*int32)(unsafe.Add(mBase, uint32(v4038)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v4038)+72)) = v4039 - int32(1)
	goto L633
L633:
	;
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v3892)+36))
	F_pfree(m, v4043)
	mBase = m.M
	v4045 = m.ExcPending
	if v4045 != 0 {
		goto L14
	} else {
		goto L634
	}
L634:
	;
	F_pfree(m, v3892)
	mBase = m.M
	v4047 = m.ExcPending
	if v4047 != 0 {
		goto L14
	} else {
		goto L635
	}
L635:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = int32(0)
	goto L619
L636:
	;
	v4190 = v4099
	v4233 = v4103
	goto L638
L637:
	;
	v4107 = *(*int32)(unsafe.Add(mBase, uint32(v165)+168))
	v4111 = int32(0)
	goto L639
L638:
	;
	F_vac_close_indexes(m, v4233, v4190, int32(0))
	mBase = m.M
	v4236 = m.ExcPending
	if v4236 != 0 {
		goto L14
	} else {
		goto L646
	}
L639:
	;
	v4159 = v4111 << (uint(int32(2)) % 32)
	v4161 = *(*int32)(unsafe.Add(mBase, uint32(v4107+v4159)))
	if v4161 == int32(0) {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	v4182 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v4190 = v4182
	v4233 = v4183
	goto L638
L641:
	;
	v4180 = v4111 + int32(1)
	if v4180 != v4103 {
		v4111 = v4180
		goto L639
	} else {
		goto L645
	}
L642:
	;
	v4164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4161)+4)))
	if v4164 != 0 {
		goto L641
	} else {
		goto L643
	}
L643:
	;
	v4166 = *(*int32)(unsafe.Add(mBase, uint32(v4159+v4099)))
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v4161)))
	v4168 = *(*float64)(unsafe.Add(mBase, uint32(v4161)+8))
	v4169 = int32(0)
	F_vac_update_relstats(m, v4166, v4167, v4168, v4169, v4169, v4169, v4169, v4169, v4169, v4169, v4169)
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
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
	v4237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+25)))
	if v4237 != int32(1) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v5187 = *(*int32)(unsafe.Add(mBase, uint32(v52)+564))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12])) = v5187
	v5193 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v5193 == int32(0) {
		goto L789
	} else {
		goto L790
	}
L648:
	;
	v4241 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])))
	if v4241&int32(1) != 0 {
		goto L647
	} else {
		goto L649
	}
L649:
	;
	v4244 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	v4245 = *(*int32)(unsafe.Add(mBase, uint32(v165)+148))
	if v4244 == v4245 {
		goto L647
	} else {
		goto L650
	}
L650:
	;
	v4247 = v4244 - v4245
	if base.B2i32(base.Ui32(v4247) <= base.Ui32(int32(999)))&base.B2i32(base.Ui32(v4247) < base.Ui32(int32(base.Ui32(v4244)>>(uint(int32(4))%32)))) != 0 {
		goto L647
	} else {
		goto L651
	}
L651:
	;
	v4258 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v4258 == int32(0) {
		goto L653
	} else {
		goto L654
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+92)) = int32(5)
	v4293 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+88)) = uint16(v4293)
	v4295 = *(*int32)(unsafe.Add(mBase, uint32(v165)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+84)) = v4295
	v4304 = v4244
	goto L656
L653:
	;
	goto L652
L654:
	;
	v4262 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v4262&int32(1) == int32(0) {
		goto L653
	} else {
		goto L655
	}
L655:
	;
	v4267 = int32(_a_F_heap_vacuum_rel_1)
	v4269 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v4270 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v4269 + v4270
	v4273 = *(*int32)(unsafe.Add(mBase, uint32(v4258)))
	*(*int32)(unsafe.Add(mBase, uint32(v4258))) = v4273 + v4270
	*(*int64)(unsafe.Add(mBase, uint32(v4258+int32(0))+232)) = int64(5)
	v4281 = *(*int32)(unsafe.Add(mBase, uint32(v4258)))
	*(*int32)(unsafe.Add(mBase, uint32(v4258))) = v4281 + v4270
	v4287 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v4287 - v4270
	goto L653
L656:
	;
	v4347 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v4348 = F_ConditionalLockRelation(m, v4347)
	mBase = m.M
	v4349 = m.ExcPending
	if v4349 != 0 {
		goto L14
	} else {
		goto L658
	}
L657:
	;
	goto L647
L658:
	;
	if v4348 == int32(0) {
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v4354 = int32(0)
	goto L662
L660:
	;
	goto L661
L661:
	;
	v4495 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v4497 = F_RelationGetNumberOfBlocksInFork(m, v4495, int32(0))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		goto L14
	} else {
		goto L682
	}
L662:
	;
	v4402 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29]))
	if v4402 != 0 {
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
	v4404 = m.ExcPending
	if v4404 != 0 {
		goto L14
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	if v4354 == int32(100) {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	goto L666
L668:
	;
	v4409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+96)))
	if v4409 != 0 {
		goto L671
	} else {
		goto L672
	}
L669:
	;
	goto L670
L670:
	;
	v4429 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[30]))
	v4433 = F_WaitLatch(m, v4429, int32(41), int32(50), int32(150994952))
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L14
	} else {
		goto L678
	}
L671:
	;
	v4410 = int32(17)
	goto L673
L672:
	;
	v4410 = int32(13)
	goto L673
L673:
	;
	v4412 = F_errstart(m, v4410, int32(0))
	mBase = m.M
	v4413 = m.ExcPending
	if v4413 != 0 {
		goto L14
	} else {
		goto L674
	}
L674:
	;
	if v4412 == int32(0) {
		goto L647
	} else {
		goto L675
	}
L675:
	;
	v4416 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+432)) = v4416
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_26), v52+int32(432))
	mBase = m.M
	v4422 = m.ExcPending
	if v4422 != 0 {
		goto L14
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(3249), int32(_a_F_heap_vacuum_rel_27))
	mBase = m.M
	v4427 = m.ExcPending
	if v4427 != 0 {
		goto L14
	} else {
		goto L677
	}
L677:
	;
	goto L647
L678:
	;
	v4436 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v4436))) = int32(0)
	goto L679
L679:
	;
	v4441 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v4442 = F_ConditionalLockRelation(m, v4441)
	mBase = m.M
	v4443 = m.ExcPending
	if v4443 != 0 {
		goto L14
	} else {
		goto L680
	}
L680:
	;
	if v4442 == int32(0) {
		v4354 = v4354 + int32(1)
		goto L662
	} else {
		goto L681
	}
L681:
	;
	goto L663
L682:
	;
	if v4497 != v4304 {
		goto L683
	} else {
		goto L684
	}
L683:
	;
	v4500 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_UnlockRelation(m, v4500)
	mBase = m.M
	v4502 = m.ExcPending
	if v4502 != 0 {
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
	v4507 = int32(0)
	v4508 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	v4509 = *(*int32)(unsafe.Add(mBase, uint32(v165)+148))
	if base.Ui32(v4508) <= base.Ui32(v4509) {
		v5056 = v4509
		v5069 = v4507
		goto L687
	} else {
		goto L688
	}
L686:
	;
	goto L647
L687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+84)) = v5056
	v5098 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if base.Ui32(v4304) <= base.Ui32(v5056) {
		goto L772
	} else {
		goto L773
	}
L688:
	;
	v4511 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+968)))
	v4512 = *(*int64)(unsafe.Add(mBase, uint32(v52)+960))
	v4522 = v4508
	v4530 = int32(-1)
	v4555 = v4511 + v4512*int64(1000000000)
	goto L689
L689:
	;
	if v4522&int32(31) != 0 {
		v4763 = v4555
		goto L691
	} else {
		goto L692
	}
L690:
	;
	v5056 = v5046
	v5069 = v4507
	goto L687
L691:
	;
	v4766 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29]))
	if v4766 != 0 {
		goto L738
	} else {
		goto L739
	}
L692:
	;
	F___clock_gettime(m, int32(1), v52+int32(960))
	mBase = m.M
	v4572 = int64(*(*int32)(unsafe.Add(mBase, uint32(v52)+968)))
	v4573 = *(*int64)(unsafe.Add(mBase, uint32(v52)+960))
	v4576 = v4572 + v4573*int64(1000000000)
	if v4576-v4555 < int64(20000000) {
		v4763 = v4555
		goto L691
	} else {
		goto L693
	}
L693:
	;
	v4580 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v4581 = m.G0
	v4583 = v4581 - int32(16)
	m.G0 = v4583
	v4585 = *(*int32)(unsafe.Add(mBase, uint32(v4580)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v4583))) = v4585
	v4587 = *(*int32)(unsafe.Add(mBase, uint32(v4580)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v4583)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v4583)+4)) = v4587
	v4591 = m.G0
	v4593 = v4591 - int32(80)
	m.G0 = v4593
	v4595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4583)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v4595-int32(3))&int32(255)) {
		goto L696
	} else {
		goto L697
	}
L694:
	;
	m.G0 = v4583 + int32(16)
	if v4696 == int32(0) {
		v4763 = v4576
		goto L691
	} else {
		goto L730
	}
L695:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4717 = m.ExcPending
	if v4717 != 0 {
		goto L14
	} else {
		goto L727
	}
L696:
	;
	v4604 = *(*int32)(unsafe.Add(mBase, uint32(v4595<<(uint(int32(2))%32))+uint32(_c_F_heap_vacuum_rel[31])))
	v4605 = *(*int32)(unsafe.Add(mBase, uint32(v4604)))
	if v4605 < int32(8) {
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
	v4704 = m.ExcPending
	if v4704 != 0 {
		goto L14
	} else {
		goto L724
	}
L699:
	;
	v4608 = *(*int64)(unsafe.Add(mBase, uint32(v4583)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4593)+64)) = v4608
	v4610 = *(*int64)(unsafe.Add(mBase, uint32(v4583)))
	*(*int64)(unsafe.Add(mBase, uint32(v4593)+56)) = v4610
	*(*int32)(unsafe.Add(mBase, uint32(v4593)+72)) = int32(8)
	v4614 = int32(0)
	v4616 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[32]))
	v4621 = F_hash_search(m, v4616, v4593+int32(56), v4614, v4614)
	mBase = m.M
	v4622 = m.ExcPending
	if v4622 != 0 {
		goto L14
	} else {
		goto L702
	}
L700:
	;
	m.G0 = v4593 + int32(80)
	goto L694
L701:
	;
	v4646 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[33]))
	v4647 = *(*int32)(unsafe.Add(mBase, uint32(v4621)+20))
	v4654 = v4646 + v4647&int32(15)<<(uint(int32(7))%32) + int32(_a_F_heap_vacuum_rel_28)
	v4656 = F_LWLockAcquire(m, v4654, int32(1))
	mBase = m.M
	v4657 = m.ExcPending
	if v4657 != 0 {
		goto L14
	} else {
		goto L711
	}
L702:
	;
	if v4621 != 0 {
		goto L703
	} else {
		goto L704
	}
L703:
	;
	v4623 = *(*int64)(unsafe.Add(mBase, uint32(v4621)+32))
	if int64(0) < v4623 {
		goto L701
	} else {
		goto L706
	}
L704:
	;
	goto L705
L705:
	;
	v4628 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4629 = m.ExcPending
	if v4629 != 0 {
		goto L14
	} else {
		goto L707
	}
L706:
	;
	goto L705
L707:
	;
	if v4628 == int32(0) {
		v4696 = v4614
		goto L700
	} else {
		goto L708
	}
L708:
	;
	v4632 = *(*int32)(unsafe.Add(mBase, uint32(v4604)+8))
	v4633 = *(*int32)(unsafe.Add(mBase, uint32(v4632)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4593)+32)) = v4633
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_29), v4593+int32(32))
	mBase = m.M
	v4639 = m.ExcPending
	if v4639 != 0 {
		goto L14
	} else {
		goto L709
	}
L709:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(736), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		goto L14
	} else {
		goto L710
	}
L710:
	;
	v4696 = v4614
	goto L700
L711:
	;
	v4658 = *(*int32)(unsafe.Add(mBase, uint32(v4621)+28))
	v4659 = *(*int32)(unsafe.Add(mBase, uint32(v4658)+12))
	if int32(base.Ui32(v4659)>>(uint(int32(8))%32))&int32(1) == int32(0) {
		goto L712
	} else {
		goto L713
	}
L712:
	;
	F_LWLockRelease(m, v4654)
	mBase = m.M
	v4667 = m.ExcPending
	if v4667 != 0 {
		goto L14
	} else {
		goto L715
	}
L713:
	;
	goto L714
L714:
	;
	v4687 = *(*int32)(unsafe.Add(mBase, uint32(v4604)+4))
	v4688 = *(*int32)(unsafe.Add(mBase, uint32(v4687)+32))
	v4689 = *(*int32)(unsafe.Add(mBase, uint32(v4621)+24))
	v4690 = *(*int32)(unsafe.Add(mBase, uint32(v4689)+20))
	F_LWLockRelease(m, v4654)
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
		goto L14
	} else {
		goto L723
	}
L715:
	;
	v4670 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4671 = m.ExcPending
	if v4671 != 0 {
		goto L14
	} else {
		goto L716
	}
L716:
	;
	if v4670 != 0 {
		goto L717
	} else {
		goto L718
	}
L717:
	;
	v4672 = *(*int32)(unsafe.Add(mBase, uint32(v4604)+8))
	v4673 = *(*int32)(unsafe.Add(mBase, uint32(v4672)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4593)+48)) = v4673
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_29), v4593+int32(48))
	mBase = m.M
	v4679 = m.ExcPending
	if v4679 != 0 {
		goto L14
	} else {
		goto L720
	}
L718:
	;
	goto L719
L719:
	;
	F_RemoveLocalLock(m, v4621)
	mBase = m.M
	v4686 = m.ExcPending
	if v4686 != 0 {
		goto L14
	} else {
		goto L722
	}
L720:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(766), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v4684 = m.ExcPending
	if v4684 != 0 {
		goto L14
	} else {
		goto L721
	}
L721:
	;
	goto L719
L722:
	;
	v4696 = v4614
	goto L700
L723:
	;
	v4696 = base.B2i32(v4690&v4688 != int32(0))
	goto L700
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4593))) = v4595
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_32), v4593)
	mBase = m.M
	v4708 = m.ExcPending
	if v4708 != 0 {
		goto L14
	} else {
		goto L725
	}
L725:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(707), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v4593)+16)) = int32(8)
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_33), v4593+int32(16))
	mBase = m.M
	v4724 = m.ExcPending
	if v4724 != 0 {
		goto L14
	} else {
		goto L728
	}
L728:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(710), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v4729 = m.ExcPending
	if v4729 != 0 {
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
	v4735 = int32(1)
	v4738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+96)))
	if v4738 != 0 {
		goto L731
	} else {
		goto L732
	}
L731:
	;
	v4739 = int32(17)
	goto L733
L732:
	;
	v4739 = int32(13)
	goto L733
L733:
	;
	v4741 = F_errstart(m, v4739, int32(0))
	mBase = m.M
	v4742 = m.ExcPending
	if v4742 != 0 {
		goto L14
	} else {
		goto L734
	}
L734:
	;
	if v4741 == int32(0) {
		v5056 = v4522
		v5069 = v4735
		goto L687
	} else {
		goto L735
	}
L735:
	;
	v4745 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+416)) = v4745
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_34), v52+int32(416))
	mBase = m.M
	v4751 = m.ExcPending
	if v4751 != 0 {
		goto L14
	} else {
		goto L736
	}
L736:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(3381), int32(_a_F_heap_vacuum_rel_35))
	mBase = m.M
	v4756 = m.ExcPending
	if v4756 != 0 {
		goto L14
	} else {
		goto L737
	}
L737:
	;
	v5056 = v4522
	v5069 = v4735
	goto L687
L738:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4768 = m.ExcPending
	if v4768 != 0 {
		goto L14
	} else {
		goto L741
	}
L739:
	;
	goto L740
L740:
	;
	v4770 = v4522 - int32(1)
	if base.Ui32(v4770) < base.Ui32(v4530) {
		goto L742
	} else {
		goto L743
	}
L741:
	;
	goto L740
L742:
	;
	v4773 = v4770 & int32(-32)
	v4776 = v4773
	goto L745
L743:
	;
	v4849 = v4530
	goto L744
L744:
	;
	v4885 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v4886 = int32(0)
	v4888 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v4889 = F_ReadBufferExtended(m, v4885, v4886, v4770, v4886, v4888)
	mBase = m.M
	v4890 = m.ExcPending
	if v4890 != 0 {
		goto L14
	} else {
		goto L753
	}
L745:
	;
	v4825 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_PrefetchBuffer(m, v52+int32(960), v4825, int32(0), v4776)
	mBase = m.M
	v4828 = m.ExcPending
	if v4828 != 0 {
		goto L14
	} else {
		goto L747
	}
L746:
	;
	v4849 = v4773
	goto L744
L747:
	;
	v4830 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29]))
	if v4830 != 0 {
		goto L748
	} else {
		goto L749
	}
L748:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4832 = m.ExcPending
	if v4832 != 0 {
		goto L14
	} else {
		goto L751
	}
L749:
	;
	goto L750
L750:
	;
	if base.Ui32(v4776) < base.Ui32(v4770) {
		v4776 = v4776 + int32(1)
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
	F_LockBuffer(m, v4889, int32(1))
	mBase = m.M
	v4893 = m.ExcPending
	if v4893 != 0 {
		goto L14
	} else {
		goto L754
	}
L754:
	;
	if v4889 < int32(0) {
		goto L757
	} else {
		goto L758
	}
L755:
	;
	F_UnlockReleaseBuffer(m, v4889)
	mBase = m.M
	v5045 = m.ExcPending
	if v5045 != 0 {
		goto L14
	} else {
		goto L770
	}
L756:
	;
	v4912 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4911)+14)))
	if v4912 == int32(0) {
		goto L755
	} else {
		goto L760
	}
L757:
	;
	v4897 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[24]))
	v4903 = *(*int32)(unsafe.Add(mBase, uint32(v4897+(v4889^int32(-1))<<(uint(int32(2))%32))))
	v4911 = v4903
	goto L756
L758:
	;
	goto L759
L759:
	;
	v4905 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[25]))
	v4911 = v4905 + v4889<<(uint(int32(13))%32) + int32(-8192)
	goto L756
L760:
	;
	v4915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4911)+12)))
	if base.Ui32(v4915) < base.Ui32(int32(25)) {
		goto L755
	} else {
		goto L761
	}
L761:
	;
	v4923 = int32(base.Ui32(v4915+int32(_a_F_heap_vacuum_rel_13))>>(uint(int32(2))%32)) & int32(_a_F_heap_vacuum_rel_14)
	if v4923 == int32(0) {
		goto L755
	} else {
		goto L762
	}
L762:
	;
	v4931 = int32(1)
	goto L763
L763:
	;
	v4983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4911+int32(20)+v4931&int32(_a_F_heap_vacuum_rel_14)<<(uint(int32(2))%32))+1)))
	if v4983&int32(384) == int32(0) {
		goto L765
	} else {
		goto L766
	}
L764:
	;
	F_UnlockReleaseBuffer(m, v4889)
	mBase = m.M
	v4994 = m.ExcPending
	if v4994 != 0 {
		goto L14
	} else {
		goto L769
	}
L765:
	;
	v4989 = v4931 + int32(1)
	if base.Ui32(v4989&int32(_a_F_heap_vacuum_rel_14)) <= base.Ui32(v4923) {
		v4931 = v4989
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
	v5056 = v4522
	v5069 = v4507
	goto L687
L770:
	;
	v5046 = *(*int32)(unsafe.Add(mBase, uint32(v165)+148))
	if base.Ui32(v5046) < base.Ui32(v4770) {
		v4522 = v4770
		v4530 = v4849
		v4555 = v4763
		goto L689
	} else {
		goto L771
	}
L771:
	;
	goto L690
L772:
	;
	F_UnlockRelation(m, v5098)
	mBase = m.M
	v5101 = m.ExcPending
	if v5101 != 0 {
		goto L14
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	F_RelationTruncate(m, v5098, v5056)
	mBase = m.M
	v5103 = m.ExcPending
	if v5103 != 0 {
		goto L14
	} else {
		goto L776
	}
L775:
	;
	goto L647
L776:
	;
	v5104 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	F_UnlockRelation(m, v5104)
	mBase = m.M
	v5106 = m.ExcPending
	if v5106 != 0 {
		goto L14
	} else {
		goto L777
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+108)) = v5056
	v5108 = *(*int32)(unsafe.Add(mBase, uint32(v165)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+120)) = v5108 + (v4304 - v5056)
	v5114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+96)))
	if v5114 != 0 {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v5115 = int32(17)
	goto L780
L779:
	;
	v5115 = int32(13)
	goto L780
L780:
	;
	v5117 = F_errstart(m, v5115, int32(0))
	mBase = m.M
	v5118 = m.ExcPending
	if v5118 != 0 {
		goto L14
	} else {
		goto L781
	}
L781:
	;
	if v5117 != 0 {
		goto L782
	} else {
		goto L783
	}
L782:
	;
	v5119 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+408)) = v5056
	*(*int32)(unsafe.Add(mBase, uint32(v52)+404)) = v4304
	*(*int32)(unsafe.Add(mBase, uint32(v52)+400)) = v5119
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_36), v52+int32(400))
	mBase = m.M
	v5127 = m.ExcPending
	if v5127 != 0 {
		goto L14
	} else {
		goto L785
	}
L783:
	;
	goto L784
L784:
	;
	v5134 = *(*int32)(unsafe.Add(mBase, uint32(v165)+148))
	if v5069&base.B2i32(base.Ui32(v5134) < base.Ui32(v5056)) != 0 {
		v4304 = v5056
		goto L656
	} else {
		goto L787
	}
L785:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(3320), int32(_a_F_heap_vacuum_rel_27))
	mBase = m.M
	v5132 = m.ExcPending
	if v5132 != 0 {
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
	v5226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+64)))
	if v5226 == int32(1) {
		goto L792
	} else {
		goto L793
	}
L789:
	;
	goto L788
L790:
	;
	v5197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v5197&int32(1) == int32(0) {
		goto L789
	} else {
		goto L791
	}
L791:
	;
	v5202 = int32(_a_F_heap_vacuum_rel_1)
	v5204 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v5205 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5204 + v5205
	v5208 = *(*int32)(unsafe.Add(mBase, uint32(v5193)))
	*(*int32)(unsafe.Add(mBase, uint32(v5193))) = v5208 + v5205
	*(*int64)(unsafe.Add(mBase, uint32(v5193+int32(0))+232)) = int64(6)
	v5216 = *(*int32)(unsafe.Add(mBase, uint32(v5193)))
	*(*int32)(unsafe.Add(mBase, uint32(v5193))) = v5216 + v5205
	v5222 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5222 - v5205
	goto L789
L792:
	;
	v5229 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1295))) = v5229
	*(*int32)(unsafe.Add(mBase, uint32(v1293))) = v5229
	goto L794
L793:
	;
	goto L794
L794:
	;
	v5233 = *(*int32)(unsafe.Add(mBase, uint32(v165)+108))
	F_visibilitymap_count(m, l0, v52+int32(1584), v52+int32(912))
	mBase = m.M
	v5239 = m.ExcPending
	if v5239 != 0 {
		goto L14
	} else {
		goto L795
	}
L795:
	;
	v5240 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1584))
	if base.Ui32(v5233) < base.Ui32(v5240) {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+1584)) = v5233
	v5243 = v5233
	goto L798
L797:
	;
	v5243 = v5240
	goto L798
L798:
	;
	v5244 = *(*int32)(unsafe.Add(mBase, uint32(v52)+912))
	if base.Ui32(v5243) < base.Ui32(v5244) {
		goto L799
	} else {
		goto L800
	}
L799:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+912)) = v5243
	v5247 = v5243
	goto L801
L800:
	;
	v5247 = v5244
	goto L801
L801:
	;
	v5248 = int32(0)
	v5249 = *(*float64)(unsafe.Add(mBase, uint32(v165)+160))
	v5250 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v5253 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v5254 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	F_vac_update_relstats(m, l0, v5233, v5249, v5243, v5247, base.B2i32(v5248 < v5250), v5253, v5254, v52+int32(924), v52+int32(908), v5248)
	mBase = m.M
	v5261 = m.ExcPending
	if v5261 != 0 {
		goto L14
	} else {
		goto L802
	}
L802:
	;
	v5262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5263)+117)))
	v5265 = *(*float64)(unsafe.Add(mBase, uint32(v165)+160))
	v5266 = float64(0)
	if base.F64_gt(v5265, v5266) != 0 {
		goto L803
	} else {
		goto L804
	}
L803:
	;
	v5269 = v5265
	goto L805
L804:
	;
	v5269 = v5266
	goto L805
L805:
	;
	v5271 = *(*int64)(unsafe.Add(mBase, uint32(v165)+216))
	v5272 = *(*int64)(unsafe.Add(mBase, uint32(v165)+208))
	v5275 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[34])))
	if v5275 == int32(1) {
		goto L806
	} else {
		goto L807
	}
L806:
	;
	v5279 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v5283 = m.G0
	v5284 = int32(16)
	v5285 = v5283 - v5284
	m.G0 = v5285
	F_gettimeofday(m, v5285)
	mBase = m.M
	v5288 = *(*int64)(unsafe.Add(mBase, uint32(v5285)))
	v5289 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5285)+8)))
	m.G0 = v5285 + v5284
	v5297 = v5289 + v5288*int64(1000000) - int64(946684800000000)
	goto L809
L807:
	;
	goto L808
L808:
	;
	v5370 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	if v5370 == int32(0) {
		goto L831
	} else {
		goto L832
	}
L809:
	;
	if v5297 <= v123 {
		v5315 = int32(0)
		goto L811
	} else {
		goto L812
	}
L810:
	;
	if v5264 != 0 {
		goto L814
	} else {
		goto L815
	}
L811:
	;
	goto L810
L812:
	;
	v5303 = v5297 - v123
	if base.B2i32(int64(0) < v123)^base.B2i32(v5303 < v5297)|base.B2i32(int64(2147483646000) < v5303) != 0 {
		v5315 = int32(2147483647)
		goto L811
	} else {
		goto L813
	}
L813:
	;
	v5312 = base.I64_div_s(v5303+int64(999), int64(1000))
	v5315 = base.I32_wrap_i64(v5312)
	goto L811
L814:
	;
	v5318 = int32(0)
	goto L816
L815:
	;
	v5318 = v5279
	goto L816
L816:
	;
	v5321 = F_pgstat_get_entry_ref_locked(m, int32(2), v5318, base.I64_extend_i32_u(v5262), int32(0))
	mBase = m.M
	v5322 = m.ExcPending
	if v5322 != 0 {
		goto L14
	} else {
		goto L817
	}
L817:
	;
	v5323 = *(*int32)(unsafe.Add(mBase, uint32(v5321)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v5323)+120)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5323)+104)) = v5271 + v5272
	*(*int64)(unsafe.Add(mBase, uint32(v5323)+96)) = base.I64_trunc_sat_f64_s(v5269)
	v5331 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[4]))
	v5333 = base.B2i32(v5331 == int32(4))
	if v5331 == int32(4) {
		goto L818
	} else {
		goto L819
	}
L818:
	;
	v5334 = int32(160)
	goto L820
L819:
	;
	v5334 = int32(144)
	goto L820
L820:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5323+v5334))) = v5297
	if v5331 == int32(4) {
		goto L821
	} else {
		goto L822
	}
L821:
	;
	v5339 = int32(168)
	goto L823
L822:
	;
	v5339 = int32(152)
	goto L823
L823:
	;
	v5340 = v5323 + v5339
	v5341 = *(*int64)(unsafe.Add(mBase, uint32(v5340)))
	*(*int64)(unsafe.Add(mBase, uint32(v5340))) = v5341 + int64(1)
	if v5331 == int32(4) {
		goto L824
	} else {
		goto L825
	}
L824:
	;
	v5347 = int32(216)
	goto L826
L825:
	;
	v5347 = int32(208)
	goto L826
L826:
	;
	v5348 = v5323 + v5347
	v5349 = *(*int64)(unsafe.Add(mBase, uint32(v5348)))
	*(*int64)(unsafe.Add(mBase, uint32(v5348))) = v5349 + base.I64_extend_i32_s(v5315)
	F_pgstat_unlock_entry(m, v5321)
	mBase = m.M
	v5354 = m.ExcPending
	if v5354 != 0 {
		goto L14
	} else {
		goto L827
	}
L827:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v5357 = m.ExcPending
	if v5357 != 0 {
		goto L14
	} else {
		goto L828
	}
L828:
	;
	v5360 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v5361 = m.ExcPending
	if v5361 != 0 {
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
	v5374 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9])))
	if v5374&int32(1) == int32(0) {
		goto L831
	} else {
		goto L833
	}
L833:
	;
	v5379 = *(*int32)(unsafe.Add(mBase, uint32(v5370)+220))
	if v5379 == int32(0) {
		goto L831
	} else {
		goto L834
	}
L834:
	;
	v5382 = int32(_a_F_heap_vacuum_rel_1)
	v5384 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	v5385 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5384 + v5385
	v5388 = *(*int32)(unsafe.Add(mBase, uint32(v5370)))
	*(*int32)(unsafe.Add(mBase, uint32(v5370))) = v5388 + v5385
	v5392 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5370)+220)) = v5392
	*(*int32)(unsafe.Add(mBase, uint32(v5370)+224)) = v5392
	*(*int32)(unsafe.Add(mBase, uint32(v5370))) = v5388 + int32(2)
	v5402 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])) = v5402 - v5385
	goto L831
L835:
	;
	v6129 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if int32(0) < v6129 {
		goto L938
	} else {
		goto L939
	}
L836:
	;
	v5411 = m.G0
	v5412 = int32(16)
	v5413 = v5411 - v5412
	m.G0 = v5413
	F_gettimeofday(m, v5413)
	mBase = m.M
	v5416 = *(*int64)(unsafe.Add(mBase, uint32(v5413)))
	v5417 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5413)+8)))
	m.G0 = v5413 + v5412
	v5425 = v5417 + v5416*int64(1000000) - int64(946684800000000)
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
	v5443 = v5425 - v123
	if v5443 <= int64(0) {
		goto L844
	} else {
		goto L845
	}
L839:
	;
	v5426 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v5426 == int32(0) {
		goto L838
	} else {
		goto L840
	}
L840:
	;
	goto L841
L841:
	;
	if base.B2i32(base.I64_extend_i32_s(v5426)*int64(1000) <= v5425-v123) == int32(0) {
		goto L835
	} else {
		goto L842
	}
L842:
	;
	goto L838
L843:
	;
	v5459 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v52)+552)) = v5459
	*(*int64)(unsafe.Add(mBase, uint32(v52)+544)) = v5459
	*(*int64)(unsafe.Add(mBase, uint32(v52)+536)) = v5459
	*(*int64)(unsafe.Add(mBase, uint32(v52)+528)) = v5459
	v5468 = v52 + int32(528)
	v5470 = v52 + int32(704)
	v5471 = *(*int64)(unsafe.Add(mBase, uint32(v5468)+16))
	v5473 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[1]))
	v5474 = *(*int64)(unsafe.Add(mBase, uint32(v5470)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5468)+16)) = v5471 + (v5473 - v5474)
	v5478 = *(*int64)(unsafe.Add(mBase, uint32(v5468)))
	v5480 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[3]))
	v5481 = *(*int64)(unsafe.Add(mBase, uint32(v5470)))
	*(*int64)(unsafe.Add(mBase, uint32(v5468))) = v5478 + (v5480 - v5481)
	v5485 = *(*int64)(unsafe.Add(mBase, uint32(v5468)+8))
	v5487 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[2]))
	v5488 = *(*int64)(unsafe.Add(mBase, uint32(v5470)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5468)+8)) = v5485 + (v5487 - v5488)
	v5492 = *(*int64)(unsafe.Add(mBase, uint32(v5468)+24))
	v5494 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[0]))
	v5495 = *(*int64)(unsafe.Add(mBase, uint32(v5470)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5468)+24)) = v5492 + (v5494 - v5495)
	goto L847
L844:
	;
	v5455 = int32(0)
	v5456 = int32(0)
	goto L846
L845:
	;
	v5447 = int64(1000000)
	v5448 = base.I64_div_u_s(v5443, v5447)
	v5455 = base.I32_wrap_i64(v5448)
	v5456 = base.I32_wrap_i64(v5443 - v5448*v5447)
	goto L846
L846:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52+int32(1608)))) = v5455
	*(*int32)(unsafe.Add(mBase, uint32(v52+int32(1600)))) = v5456
	goto L843
L847:
	;
	v5500 = v52 + int32(960)
	base.MemoryFill(m, v5500, int32(0), int32(128))
	v5505 = v52 + int32(576)
	v5506 = *(*int64)(unsafe.Add(mBase, uint32(v5500)))
	v5508 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[35]))
	v5509 = *(*int64)(unsafe.Add(mBase, uint32(v5505)))
	*(*int64)(unsafe.Add(mBase, uint32(v5500))) = v5506 + (v5508 - v5509)
	v5513 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+8))
	v5515 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[36]))
	v5516 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+8)) = v5513 + (v5515 - v5516)
	v5520 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+16))
	v5522 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[37]))
	v5523 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+16)) = v5520 + (v5522 - v5523)
	v5527 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+24))
	v5529 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[38]))
	v5530 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+24)) = v5527 + (v5529 - v5530)
	v5534 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+32))
	v5536 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[39]))
	v5537 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+32)) = v5534 + (v5536 - v5537)
	v5541 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+40))
	v5543 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[40]))
	v5544 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+40)) = v5541 + (v5543 - v5544)
	v5548 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+48))
	v5550 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[41]))
	v5551 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+48)) = v5548 + (v5550 - v5551)
	v5555 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+56))
	v5557 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[42]))
	v5558 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+56)) = v5555 + (v5557 - v5558)
	v5562 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+64))
	v5564 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[43]))
	v5565 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+64)) = v5562 + (v5564 - v5565)
	v5569 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+72))
	v5571 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[44]))
	v5572 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+72)) = v5569 + (v5571 - v5572)
	v5576 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+80))
	v5578 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[45]))
	v5579 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+80)) = v5576 + (v5578 - v5579)
	v5583 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+88))
	v5585 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[46]))
	v5586 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+88)) = v5583 + (v5585 - v5586)
	v5590 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+96))
	v5592 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[47]))
	v5593 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+96)) = v5590 + (v5592 - v5593)
	v5597 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+104))
	v5599 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[48]))
	v5600 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+104)) = v5597 + (v5599 - v5600)
	v5604 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+112))
	v5606 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[49]))
	v5607 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+112)) = v5604 + (v5606 - v5607)
	v5611 = *(*int64)(unsafe.Add(mBase, uint32(v5500)+120))
	v5613 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[50]))
	v5614 = *(*int64)(unsafe.Add(mBase, uint32(v5505)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v5500)+120)) = v5611 + (v5613 - v5614)
	goto L848
L848:
	;
	v5618 = *(*int64)(unsafe.Add(mBase, uint32(v52)+976))
	v5619 = *(*int64)(unsafe.Add(mBase, uint32(v52)+1008))
	v5620 = *(*int64)(unsafe.Add(mBase, uint32(v52)+968))
	v5621 = *(*int64)(unsafe.Add(mBase, uint32(v52)+1000))
	v5622 = *(*int64)(unsafe.Add(mBase, uint32(v52)+960))
	v5623 = *(*int64)(unsafe.Add(mBase, uint32(v52)+992))
	F_initStringInfo(m, v52+int32(928))
	mBase = m.M
	v5627 = m.ExcPending
	if v5627 != 0 {
		goto L14
	} else {
		goto L849
	}
L849:
	;
	if v73 != 0 {
		v5644 = int32(_a_F_heap_vacuum_rel_37)
		goto L850
	} else {
		goto L851
	}
L850:
	;
	v5645 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	v5646 = *(*int64)(unsafe.Add(mBase, uint32(v165)+68))
	v5647 = *(*int32)(unsafe.Add(mBase, uint32(v165)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+396)) = v5647
	*(*int64)(unsafe.Add(mBase, uint32(v52)+384)) = v5646
	*(*int32)(unsafe.Add(mBase, uint32(v52)+392)) = v5645
	v5652 = v52 + int32(928)
	F_appendStringInfo(m, v5652, v5644, v52+int32(384))
	mBase = m.M
	v5656 = m.ExcPending
	if v5656 != 0 {
		goto L14
	} else {
		goto L859
	}
L851:
	;
	v5631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)))
	if v5631&int32(1) != 0 {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	v5634 = int32(_a_F_heap_vacuum_rel_38)
	goto L854
L853:
	;
	v5634 = int32(_a_F_heap_vacuum_rel_39)
	goto L854
L854:
	;
	v5635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v5635 == int32(1) {
		v5644 = v5634
		goto L850
	} else {
		goto L855
	}
L855:
	;
	if v5631&int32(1) != 0 {
		goto L856
	} else {
		goto L857
	}
L856:
	;
	v5642 = int32(_a_F_heap_vacuum_rel_40)
	goto L858
L857:
	;
	v5642 = int32(_a_F_heap_vacuum_rel_41)
	goto L858
L858:
	;
	v5644 = v5642
	goto L850
L859:
	;
	v5657 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	v5658 = *(*int32)(unsafe.Add(mBase, uint32(v165)+120))
	v5659 = *(*int32)(unsafe.Add(mBase, uint32(v165)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+376)) = v5659
	v5662 = float64(100)
	if v394 != 0 {
		goto L860
	} else {
		goto L861
	}
L860:
	;
	v5667 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v5657), v5662), base.F64_convert_i32_u(v394))
	goto L862
L861:
	;
	v5667 = v5662
	goto L862
L862:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+368)) = v5667
	*(*int32)(unsafe.Add(mBase, uint32(v52)+360)) = v5657
	*(*int32)(unsafe.Add(mBase, uint32(v52)+356)) = v5233
	*(*int32)(unsafe.Add(mBase, uint32(v52)+352)) = v5658
	F_appendStringInfo(m, v5652, int32(_a_F_heap_vacuum_rel_42), v52+int32(352))
	mBase = m.M
	v5676 = m.ExcPending
	if v5676 != 0 {
		goto L14
	} else {
		goto L863
	}
L863:
	;
	v5677 = *(*float64)(unsafe.Add(mBase, uint32(v165)+152))
	v5678 = *(*int64)(unsafe.Add(mBase, uint32(v165)+176))
	v5679 = *(*int64)(unsafe.Add(mBase, uint32(v165)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+336)) = v5679
	*(*int64)(unsafe.Add(mBase, uint32(v52)+320)) = v5678
	*(*int64)(unsafe.Add(mBase, uint32(v52)+328)) = base.I64_trunc_sat_f64_s(v5677)
	F_appendStringInfo(m, v5652, int32(_a_F_heap_vacuum_rel_43), v52+int32(320))
	mBase = m.M
	v5688 = m.ExcPending
	if v5688 != 0 {
		goto L14
	} else {
		goto L864
	}
L864:
	;
	v5689 = *(*int64)(unsafe.Add(mBase, uint32(v165)+216))
	if int64(0) < v5689 {
		goto L865
	} else {
		goto L866
	}
L865:
	;
	v5692 = *(*int32)(unsafe.Add(mBase, uint32(v165)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+312)) = v5692
	*(*int64)(unsafe.Add(mBase, uint32(v52)+304)) = v5689
	F_appendStringInfo(m, v5652, int32(_a_F_heap_vacuum_rel_44), v52+int32(304))
	mBase = m.M
	v5699 = m.ExcPending
	if v5699 != 0 {
		goto L14
	} else {
		goto L868
	}
L866:
	;
	goto L867
L867:
	;
	v5700 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v5701 = m.ExcPending
	if v5701 != 0 {
		goto L14
	} else {
		goto L869
	}
L868:
	;
	goto L867
L869:
	;
	v5702 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+288)) = v5702
	*(*int32)(unsafe.Add(mBase, uint32(v52)+292)) = base.I32_wrap_i64(v5700) - v5702
	v5708 = v52 + int32(928)
	F_appendStringInfo(m, v5708, int32(_a_F_heap_vacuum_rel_45), v52+int32(288))
	mBase = m.M
	v5713 = m.ExcPending
	if v5713 != 0 {
		goto L14
	} else {
		goto L870
	}
L870:
	;
	v5714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+924)))
	if v5714 == int32(1) {
		goto L871
	} else {
		goto L872
	}
L871:
	;
	v5717 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v5718 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+272)) = v5718
	*(*int32)(unsafe.Add(mBase, uint32(v52)+276)) = v5718 - v5717
	F_appendStringInfo(m, v5708, int32(_a_F_heap_vacuum_rel_46), v52+int32(272))
	mBase = m.M
	v5726 = m.ExcPending
	if v5726 != 0 {
		goto L14
	} else {
		goto L874
	}
L872:
	;
	goto L873
L873:
	;
	v5729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+908)))
	if v5729 == int32(1) {
		goto L875
	} else {
		goto L876
	}
L874:
	;
	goto L873
L875:
	;
	v5732 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v5733 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+256)) = v5733
	*(*int32)(unsafe.Add(mBase, uint32(v52)+260)) = v5733 - v5732
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_47), v52+int32(256))
	mBase = m.M
	v5743 = m.ExcPending
	if v5743 != 0 {
		goto L14
	} else {
		goto L878
	}
L876:
	;
	goto L877
L877:
	;
	v5746 = *(*int32)(unsafe.Add(mBase, uint32(v165)+124))
	v5747 = *(*int64)(unsafe.Add(mBase, uint32(v165)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+240)) = v5747
	v5750 = float64(100)
	if v394 != 0 {
		goto L879
	} else {
		goto L880
	}
L878:
	;
	goto L877
L879:
	;
	v5755 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v5746), v5750), base.F64_convert_i32_u(v394))
	goto L881
L880:
	;
	v5755 = v5750
	goto L881
L881:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+232)) = v5755
	*(*int32)(unsafe.Add(mBase, uint32(v52)+224)) = v5746
	v5759 = v52 + int32(928)
	F_appendStringInfo(m, v5759, int32(_a_F_heap_vacuum_rel_48), v52+int32(224))
	mBase = m.M
	v5764 = m.ExcPending
	if v5764 != 0 {
		goto L14
	} else {
		goto L882
	}
L882:
	;
	v5765 = *(*int32)(unsafe.Add(mBase, uint32(v165)+132))
	v5766 = *(*int32)(unsafe.Add(mBase, uint32(v165)+128))
	v5767 = *(*int32)(unsafe.Add(mBase, uint32(v165)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+216)) = v5767
	*(*int32)(unsafe.Add(mBase, uint32(v52)+208)) = v5766
	*(*int32)(unsafe.Add(mBase, uint32(v52)+212)) = v5767 + v5765
	F_appendStringInfo(m, v5759, int32(_a_F_heap_vacuum_rel_49), v52+int32(208))
	mBase = m.M
	v5776 = m.ExcPending
	if v5776 != 0 {
		goto L14
	} else {
		goto L883
	}
L883:
	;
	v5777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+23)))
	if v5777 == int32(1) {
		goto L885
	} else {
		goto L886
	}
L884:
	;
	F_appendStringInfoString(m, v5759, v5796)
	mBase = m.M
	v5798 = m.ExcPending
	if v5798 != 0 {
		goto L14
	} else {
		goto L895
	}
L885:
	;
	v5780 = int32(_a_F_heap_vacuum_rel_50)
	v5782 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v5782 == int32(0) {
		v5795 = v5780
		v5796 = int32(_a_F_heap_vacuum_rel_51)
		goto L884
	} else {
		goto L888
	}
L886:
	;
	goto L887
L887:
	;
	v5793 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])))
	if v5793 != 0 {
		goto L892
	} else {
		goto L893
	}
L888:
	;
	v5787 = *(*int32)(unsafe.Add(mBase, uint32(v1297)))
	if v5787 != 0 {
		goto L889
	} else {
		goto L890
	}
L889:
	;
	v5788 = int32(_a_F_heap_vacuum_rel_52)
	goto L891
L890:
	;
	v5788 = int32(_a_F_heap_vacuum_rel_51)
	goto L891
L891:
	;
	v5795 = v5780
	v5796 = v5788
	goto L884
L892:
	;
	v5794 = int32(_a_F_heap_vacuum_rel_53)
	goto L894
L893:
	;
	v5794 = int32(_a_F_heap_vacuum_rel_54)
	goto L894
L894:
	;
	v5795 = int32(_a_F_heap_vacuum_rel_55)
	v5796 = v5794
	goto L884
L895:
	;
	v5799 = *(*int32)(unsafe.Add(mBase, uint32(v165+int32(140))))
	v5800 = *(*int64)(unsafe.Add(mBase, uint32(v165)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+192)) = v5800
	v5803 = float64(100)
	if v394 != 0 {
		goto L896
	} else {
		goto L897
	}
L896:
	;
	v5808 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v5799), v5803), base.F64_convert_i32_u(v394))
	goto L898
L897:
	;
	v5808 = v5803
	goto L898
L898:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+184)) = v5808
	*(*int32)(unsafe.Add(mBase, uint32(v52)+176)) = v5799
	F_appendStringInfo(m, v52+int32(928), v5795, v52+int32(176))
	mBase = m.M
	v5816 = m.ExcPending
	if v5816 != 0 {
		goto L14
	} else {
		goto L899
	}
L899:
	;
	v5817 = int32(0)
	v5818 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if v5817 < v5818 {
		goto L900
	} else {
		goto L901
	}
L900:
	;
	v5828 = v5817
	v5831 = v5818
	goto L903
L901:
	;
	goto L902
L902:
	;
	v5951 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[51])))
	if v5951 != 0 {
		goto L910
	} else {
		goto L911
	}
L903:
	;
	v5873 = v5828 << (uint(int32(2)) % 32)
	v5874 = *(*int32)(unsafe.Add(mBase, uint32(v165)+168))
	v5876 = *(*int32)(unsafe.Add(mBase, uint32(v5873+v5874)))
	if v5876 != 0 {
		goto L905
	} else {
		goto L906
	}
L904:
	;
	goto L902
L905:
	;
	v5878 = *(*int32)(unsafe.Add(mBase, uint32(v5873+v319)))
	v5879 = *(*int64)(unsafe.Add(mBase, uint32(v5876)+24))
	v5880 = *(*int32)(unsafe.Add(mBase, uint32(v5876)))
	v5881 = *(*int32)(unsafe.Add(mBase, uint32(v5876)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v52+int32(160)))) = v5881
	*(*int32)(unsafe.Add(mBase, uint32(v52)+148)) = v5880
	*(*int64)(unsafe.Add(mBase, uint32(v52)+152)) = v5879
	*(*int32)(unsafe.Add(mBase, uint32(v52)+144)) = v5878
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_56), v52+int32(144))
	mBase = m.M
	v5892 = m.ExcPending
	if v5892 != 0 {
		goto L14
	} else {
		goto L908
	}
L906:
	;
	v5896 = v5831
	goto L907
L907:
	;
	v5899 = v5828 + int32(1)
	if v5899 < v5896 {
		v5828 = v5899
		v5831 = v5896
		goto L903
	} else {
		goto L909
	}
L908:
	;
	v5893 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v5896 = v5893
	goto L907
L909:
	;
	goto L904
L910:
	;
	v5953 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	v5954 = *(*int64)(unsafe.Add(mBase, uint32(v5953)+312))
	*(*float64)(unsafe.Add(mBase, uint32(v52)+128)) = base.F64_div(base.F64_convert_i64_s(v5954), float64(1e+06))
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_57), v52+int32(128))
	mBase = m.M
	v5965 = m.ExcPending
	if v5965 != 0 {
		goto L14
	} else {
		goto L913
	}
L911:
	;
	goto L912
L912:
	;
	v5967 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5])))
	if v5967 == int32(1) {
		goto L914
	} else {
		goto L915
	}
L913:
	;
	goto L912
L914:
	;
	v5971 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[7]))
	v5974 = float64(1000)
	*(*float64)(unsafe.Add(mBase, uint32(v52)+112)) = base.F64_div(base.F64_convert_i64_s(v5971-v105), v5974)
	v5978 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[6]))
	*(*float64)(unsafe.Add(mBase, uint32(v52)+120)) = base.F64_div(base.F64_convert_i64_s(v5978-v104), v5974)
	F_appendStringInfo(m, v52+int32(928), int32(_a_F_heap_vacuum_rel_58), v52+int32(112))
	mBase = m.M
	v5990 = m.ExcPending
	if v5990 != 0 {
		goto L14
	} else {
		goto L917
	}
L915:
	;
	goto L916
L916:
	;
	v5991 = v5619 + v5618
	v5992 = v5621 + v5620
	v5994 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1600))
	v5995 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1608))
	if v5995 <= int32(0) {
		goto L919
	} else {
		goto L920
	}
L917:
	;
	goto L916
L918:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v52)+104)) = v6019
	*(*float64)(unsafe.Add(mBase, uint32(v52)+96)) = v6020
	v6024 = v52 + int32(928)
	F_appendStringInfo(m, v6024, int32(_a_F_heap_vacuum_rel_59), v52+int32(96))
	mBase = m.M
	v6029 = m.ExcPending
	if v6029 != 0 {
		goto L14
	} else {
		goto L923
	}
L919:
	;
	if v5994 <= int32(0) {
		v6019 = float64(0)
		v6020 = float64(0)
		goto L918
	} else {
		goto L922
	}
L920:
	;
	goto L921
L921:
	;
	v6002 = float64(8192)
	v6004 = float64(9.5367431640625e-07)
	v6010 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v5994), float64(1e+06)), base.F64_convert_i32_s(v5995))
	v6019 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_convert_i64_s(v5991), v6002), v6004), v6010)
	v6020 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_convert_i64_s(v5992), v6002), v6004), v6010)
	goto L918
L922:
	;
	goto L921
L923:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v52)+80)) = v5991
	*(*int64)(unsafe.Add(mBase, uint32(v52)+72)) = v5992
	*(*int64)(unsafe.Add(mBase, uint32(v52)+64)) = v5623 + v5622
	F_appendStringInfo(m, v6024, int32(_a_F_heap_vacuum_rel_60), v52-int32(-64))
	mBase = m.M
	v6037 = m.ExcPending
	if v6037 != 0 {
		goto L14
	} else {
		goto L924
	}
L924:
	;
	v6038 = *(*int64)(unsafe.Add(mBase, uint32(v52)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+48)) = v6038
	v6040 = *(*int64)(unsafe.Add(mBase, uint32(v52)+552))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+56)) = v6040
	v6042 = *(*int64)(unsafe.Add(mBase, uint32(v52)+528))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+32)) = v6042
	v6044 = *(*int64)(unsafe.Add(mBase, uint32(v52)+536))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+40)) = v6044
	F_appendStringInfo(m, v6024, int32(_a_F_heap_vacuum_rel_61), v52+int32(32))
	mBase = m.M
	v6050 = m.ExcPending
	if v6050 != 0 {
		goto L14
	} else {
		goto L925
	}
L925:
	;
	v6053 = F_pg_rusage_show(m, v52+int32(736))
	mBase = m.M
	v6054 = m.ExcPending
	if v6054 != 0 {
		goto L14
	} else {
		goto L926
	}
L926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v6053
	F_appendStringInfo(m, v6024, int32(_a_F_heap_vacuum_rel_62), v52+int32(16))
	mBase = m.M
	v6060 = m.ExcPending
	if v6060 != 0 {
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
	v6063 = int32(17)
	goto L930
L929:
	;
	v6063 = int32(15)
	goto L930
L930:
	;
	v6065 = F_errstart(m, v6063, int32(0))
	mBase = m.M
	v6066 = m.ExcPending
	if v6066 != 0 {
		goto L14
	} else {
		goto L931
	}
L931:
	;
	if v6065 != 0 {
		goto L932
	} else {
		goto L933
	}
L932:
	;
	v6067 = *(*int32)(unsafe.Add(mBase, uint32(v52)+928))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v6067
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_63), v52)
	mBase = m.M
	v6071 = m.ExcPending
	if v6071 != 0 {
		goto L14
	} else {
		goto L935
	}
L933:
	;
	goto L934
L934:
	;
	v6077 = *(*int32)(unsafe.Add(mBase, uint32(v52)+928))
	F_pfree(m, v6077)
	mBase = m.M
	v6079 = m.ExcPending
	if v6079 != 0 {
		goto L14
	} else {
		goto L937
	}
L935:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_6), int32(1147), int32(_a_F_heap_vacuum_rel_7))
	mBase = m.M
	v6076 = m.ExcPending
	if v6076 != 0 {
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
	v6134 = v5248
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
	v6182 = v6134 << (uint(int32(2)) % 32)
	v6183 = *(*int32)(unsafe.Add(mBase, uint32(v165)+168))
	v6185 = *(*int32)(unsafe.Add(mBase, uint32(v6182+v6183)))
	if v6185 != 0 {
		goto L943
	} else {
		goto L944
	}
L942:
	;
	goto L940
L943:
	;
	F_pfree(m, v6185)
	mBase = m.M
	v6187 = m.ExcPending
	if v6187 != 0 {
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
	v6189 = *(*int32)(unsafe.Add(mBase, uint32(v6182+v319)))
	F_pfree(m, v6189)
	mBase = m.M
	v6191 = m.ExcPending
	if v6191 != 0 {
		goto L14
	} else {
		goto L950
	}
L948:
	;
	goto L949
L949:
	;
	v6193 = v6134 + int32(1)
	v6194 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v6193 < v6194 {
		v6134 = v6193
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
