package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_first(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v313 int32
	_ = v313
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v520 int32
	_ = v520
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v853 int64
	_ = v853
	var v855 int64
	_ = v855
	var v857 int64
	_ = v857
	var v859 int64
	_ = v859
	var v861 int64
	_ = v861
	var v863 int64
	_ = v863
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1162 int32
	_ = v1162
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1259 int32
	_ = v1259
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1340 int32
	_ = v1340
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1479 int32
	_ = v1479
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1512 int32
	_ = v1512
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1577 int32
	_ = v1577
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1627 int32
	_ = v1627
	var v1643 int32
	_ = v1643
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1673 int32
	_ = v1673
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1757 int64
	_ = v1757
	var v1759 int64
	_ = v1759
	var v1761 int64
	_ = v1761
	var v1763 int64
	_ = v1763
	var v1765 int64
	_ = v1765
	var v1767 int64
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1779 int64
	_ = v1779
	var v1799 int32
	_ = v1799
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1814 int32
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1868 int32
	_ = v1868
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1890 int32
	_ = v1890
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2116 int32
	_ = v2116
	var v2119 int32
	_ = v2119
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2144 int32
	_ = v2144
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2158 int64
	_ = v2158
	var v2160 int64
	_ = v2160
	var v2162 int64
	_ = v2162
	var v2164 int64
	_ = v2164
	var v2166 int64
	_ = v2166
	var v2168 int64
	_ = v2168
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2201 int64
	_ = v2201
	var v2203 int64
	_ = v2203
	var v2205 int64
	_ = v2205
	var v2207 int64
	_ = v2207
	var v2209 int64
	_ = v2209
	var v2211 int64
	_ = v2211
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2224 int32
	_ = v2224
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2240 int64
	_ = v2240
	var v2242 int64
	_ = v2242
	var v2244 int64
	_ = v2244
	var v2246 int64
	_ = v2246
	var v2248 int64
	_ = v2248
	var v2250 int64
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2272 int64
	_ = v2272
	var v2274 int64
	_ = v2274
	var v2276 int64
	_ = v2276
	var v2278 int64
	_ = v2278
	var v2280 int64
	_ = v2280
	var v2282 int64
	_ = v2282
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2304 int64
	_ = v2304
	var v2306 int64
	_ = v2306
	var v2308 int64
	_ = v2308
	var v2310 int64
	_ = v2310
	var v2312 int64
	_ = v2312
	var v2314 int64
	_ = v2314
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2332 int64
	_ = v2332
	var v2346 int32
	_ = v2346
	var v2351 int32
	_ = v2351
	var v2353 int32
	_ = v2353
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2367 int32
	_ = v2367
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2383 int32
	_ = v2383
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2432 int64
	_ = v2432
	var v2434 int64
	_ = v2434
	var v2436 int64
	_ = v2436
	var v2438 int64
	_ = v2438
	var v2440 int64
	_ = v2440
	var v2442 int64
	_ = v2442
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2469 int32
	_ = v2469
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2497 int32
	_ = v2497
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2533 int32
	_ = v2533
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2566 int32
	_ = v2566
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2572 int64
	_ = v2572
	var v2574 int64
	_ = v2574
	var v2576 int64
	_ = v2576
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2630 int32
	_ = v2630
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2637 int32
	_ = v2637
	var v2639 int32
	_ = v2639
	var v2645 int32
	_ = v2645
	var v2651 int32
	_ = v2651
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2698 int32
	_ = v2698
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2721 int32
	_ = v2721
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2891 int32
	_ = v2891
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2901 int32
	_ = v2901
	var v2906 int32
	_ = v2906
	var v2994 int32
	_ = v2994
	var v2997 int32
	_ = v2997
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3008 int32
	_ = v3008
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3019 int32
	_ = v3019
	var v3021 int32
	_ = v3021
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3053 int32
	_ = v3053
	var v3055 int32
	_ = v3055
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3081 int32
	_ = v3081
	var v3085 int32
	_ = v3085
	var v3093 int32
	_ = v3093
	var v3097 int32
	_ = v3097
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3126 int32
	_ = v3126
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3134 int32
	_ = v3134
	var v3139 int32
	_ = v3139
	var v3151 int32
	_ = v3151
	var v3182 int32
	_ = v3182
	var v3194 int32
	_ = v3194
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3233 int32
	_ = v3233
	var v3237 int32
	_ = v3237
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3243 int32
	_ = v3243
	var v3247 int32
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3253 int32
	_ = v3253
	var v3257 int32
	_ = v3257
	var v3259 int32
	_ = v3259
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3272 int32
	_ = v3272
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3285 int32
	_ = v3285
	var v3287 int32
	_ = v3287
	var v3288 int32
	_ = v3288
	var v3293 int32
	_ = v3293
	var v3312 int32
	_ = v3312
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3343 int32
	_ = v3343
	var v3345 int32
	_ = v3345
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3396 int32
	_ = v3396
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3425 int32
	_ = v3425
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3432 int32
	_ = v3432
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3444 int32
	_ = v3444
	var v3445 int32
	_ = v3445
	var v3448 int32
	_ = v3448
	var v3450 int32
	_ = v3450
	var v3459 int32
	_ = v3459
	var v3461 int32
	_ = v3461
	var v3472 int32
	_ = v3472
	var v3493 int32
	_ = v3493
	var v3496 int32
	_ = v3496
	var v3498 int32
	_ = v3498
	var v3503 int32
	_ = v3503
	var v3504 int64
	_ = v3504
	var v3506 int64
	_ = v3506
	var v3508 int64
	_ = v3508
	var v3510 int64
	_ = v3510
	var v3512 int64
	_ = v3512
	var v3514 int64
	_ = v3514
	var v3516 int32
	_ = v3516
	var v3521 int32
	_ = v3521
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3530 int64
	_ = v3530
	var v3532 int64
	_ = v3532
	var v3534 int64
	_ = v3534
	var v3542 int32
	_ = v3542
	var v3543 int64
	_ = v3543
	var v3545 int64
	_ = v3545
	var v3547 int64
	_ = v3547
	var v3549 int64
	_ = v3549
	var v3551 int64
	_ = v3551
	var v3553 int64
	_ = v3553
	var v3555 int32
	_ = v3555
	var v3559 int32
	_ = v3559
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3566 int32
	_ = v3566
	var v3569 int32
	_ = v3569
	var v3570 int32
	_ = v3570
	var v3572 int64
	_ = v3572
	var v3574 int64
	_ = v3574
	var v3576 int64
	_ = v3576
	var v3580 int32
	_ = v3580
	var v3586 int32
	_ = v3586
	var v3590 int32
	_ = v3590
	var v3627 int32
	_ = v3627
	var v3685 int32
	_ = v3685
	var v3687 int32
	_ = v3687
	var v3720 int32
	_ = v3720
	var v3721 int32
	_ = v3721
	var v3729 int32
	_ = v3729
	var v3731 int32
	_ = v3731
	var v3764 int32
	_ = v3764
	var v3765 int32
	_ = v3765
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3773 int32
	_ = v3773
	var v3775 int32
	_ = v3775
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3780 int32
	_ = v3780
	var v3781 int32
	_ = v3781
	var v3784 int32
	_ = v3784
	var v3785 int32
	_ = v3785
	var v3788 int32
	_ = v3788
	var v3795 int32
	_ = v3795
	var v3832 int32
	_ = v3832
	var v3835 int32
	_ = v3835
	var v3836 int32
	_ = v3836
	var v3840 int32
	_ = v3840
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3861 int32
	_ = v3861
	var v3886 int32
	_ = v3886
	var v3890 int32
	_ = v3890
	var v3892 int32
	_ = v3892
	var v3894 int32
	_ = v3894
	var v3940 int32
	_ = v3940
	var v3944 int32
	_ = v3944
	var v3949 int32
	_ = v3949
	var v3954 int32
	_ = v3954
	var v3960 int32
	_ = v3960
	var v4005 int32
	_ = v4005
	var v4008 int32
	_ = v4008
	var v4016 int32
	_ = v4016
	var v4017 int32
	_ = v4017
	var v4018 int32
	_ = v4018
	var v4021 int32
	_ = v4021
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4027 int32
	_ = v4027
	var v4028 int32
	_ = v4028
	var v4030 int32
	_ = v4030
	var v4031 int32
	_ = v4031
	var v4035 int32
	_ = v4035
	var v4038 int32
	_ = v4038
	var v4039 int32
	_ = v4039
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4045 int32
	_ = v4045
	var v4048 int32
	_ = v4048
	var v4051 int32
	_ = v4051
	var v4054 int32
	_ = v4054
	var v4058 int32
	_ = v4058
	var v4059 int32
	_ = v4059
	var v4060 int32
	_ = v4060
	var v4061 int64
	_ = v4061
	var v4066 int32
	_ = v4066
	var v4067 int64
	_ = v4067
	var v4071 int32
	_ = v4071
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4082 int32
	_ = v4082
	var v4084 int32
	_ = v4084
	var v4087 int32
	_ = v4087
	var v4100 int32
	_ = v4100
	var v4107 int32
	_ = v4107
	var v4108 int32
	_ = v4108
	var v4111 int32
	_ = v4111
	var v4114 int32
	_ = v4114
	var v4123 int32
	_ = v4123
	var v4125 int32
	_ = v4125
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4134 int32
	_ = v4134
	var v4135 int32
	_ = v4135
	var v4155 int32
	_ = v4155
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4183 int32
	_ = v4183
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4189 int32
	_ = v4189
	var v4220 int32
	_ = v4220
	var v4231 int32
	_ = v4231
	var v4232 int32
	_ = v4232
	var v4237 int32
	_ = v4237
	var v4246 int32
	_ = v4246
	var v4251 int32
	_ = v4251
	var v4257 int32
	_ = v4257
	var v4262 int32
	_ = v4262
	var v4291 int32
	_ = v4291
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4318 int32
	_ = v4318
	var v4322 int32
	_ = v4322
	var v4325 int32
	_ = v4325
	var v4326 int32
	_ = v4326
	var v4328 int32
	_ = v4328
	var v4333 int32
	_ = v4333
	var v4337 int32
	_ = v4337
	var v4340 int32
	_ = v4340
	var v4353 int32
	_ = v4353
	var v4360 int32
	_ = v4360
	var v4364 int32
	_ = v4364
	var v4375 int32
	_ = v4375
	var v4376 int32
	_ = v4376
	var v4383 int32
	_ = v4383
	var v4386 int32
	_ = v4386
	var v4389 int32
	_ = v4389
	var v4402 int32
	_ = v4402
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4413 int32
	_ = v4413
	var v4433 int32
	_ = v4433
	var v4453 int32
	_ = v4453
	var v4473 int32
	_ = v4473
	var v4491 int32
	_ = v4491
	var v4516 int32
	_ = v4516
	var v4520 int32
	_ = v4520
	var v4521 int32
	_ = v4521
	var v4526 int32
	_ = v4526
	var v4527 int32
	_ = v4527
	var v4528 int64
	_ = v4528
	var v4530 int64
	_ = v4530
	var v4532 int64
	_ = v4532
	var v4534 int64
	_ = v4534
	var v4536 int64
	_ = v4536
	var v4538 int64
	_ = v4538
	var v4558 int32
	_ = v4558
	var v4565 int32
	_ = v4565
	var v4580 int32
	_ = v4580
	var v4585 int32
	_ = v4585
	var v4587 int32
	_ = v4587
	var v4589 int32
	_ = v4589
	var v4590 int64
	_ = v4590
	var v4592 int64
	_ = v4592
	var v4594 int64
	_ = v4594
	var v4596 int64
	_ = v4596
	var v4598 int64
	_ = v4598
	var v4600 int64
	_ = v4600
	var v4603 int32
	_ = v4603
	var v4604 int32
	_ = v4604
	var v4617 int32
	_ = v4617
	var v4618 int32
	_ = v4618
	var v4620 int32
	_ = v4620
	var v4623 int32
	_ = v4623
	var v4625 int32
	_ = v4625
	var v4626 int32
	_ = v4626
	var v4629 int32
	_ = v4629
	var v4630 int32
	_ = v4630
	var v4631 int32
	_ = v4631
	var v4632 int32
	_ = v4632
	var v4633 int32
	_ = v4633
	var v4634 int32
	_ = v4634
	var v4638 int32
	_ = v4638
	var v4645 int32
	_ = v4645
	var v4647 int32
	_ = v4647
	var v4649 int32
	_ = v4649
	var v4651 int32
	_ = v4651
	var v4652 int32
	_ = v4652
	var v4658 int32
	_ = v4658
	var v4659 int32
	_ = v4659
	var v4661 int32
	_ = v4661
	var v4662 int32
	_ = v4662
	var v4663 int32
	_ = v4663
	var v4665 int32
	_ = v4665
	var v4669 int32
	_ = v4669
	var v4696 int32
	_ = v4696
	var v4712 int32
	_ = v4712
	var v4716 int32
	_ = v4716
	var v4718 int32
	_ = v4718
	var v4723 int32
	_ = v4723
	var v4727 int32
	_ = v4727
	var v4731 int32
	_ = v4731
	var v4733 int32
	_ = v4733
	var v4777 int32
	_ = v4777
	var v4780 int32
	_ = v4780
	var v4784 int32
	_ = v4784
	var v4786 int32
	_ = v4786
	var v4830 int32
	_ = v4830
	var v4834 int32
	_ = v4834
	var v4837 int32
	_ = v4837
	var v4841 int32
	_ = v4841
	var v4843 int32
	_ = v4843
	var v4887 int32
	_ = v4887
	var v4890 int32
	_ = v4890
	var v4894 int32
	_ = v4894
	var v4896 int32
	_ = v4896
	var v4940 int32
	_ = v4940
	var v4945 int32
	_ = v4945
	var v4951 int32
	_ = v4951
	var v4956 int32
	_ = v4956
	var v4957 int32
	_ = v4957
	var v4999 int32
	_ = v4999
	var v5002 int32
	_ = v5002
	var v5004 int32
	_ = v5004
	var v5006 int32
	_ = v5006
	var v5007 int32
	_ = v5007
	var v5009 int32
	_ = v5009
	var v5010 int32
	_ = v5010
	var v5014 int32
	_ = v5014
	var v5017 int32
	_ = v5017
	var v5019 int32
	_ = v5019
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5025 int32
	_ = v5025
	var v5026 int32
	_ = v5026
	var v5029 int32
	_ = v5029
	var v5032 int32
	_ = v5032
	var v5033 int32
	_ = v5033
	var v5034 int32
	_ = v5034
	var v5035 int32
	_ = v5035
	var v5038 int32
	_ = v5038
	var v5041 int32
	_ = v5041
	var v5042 int32
	_ = v5042
	var v5045 int32
	_ = v5045
	var v5046 int32
	_ = v5046
	var v5048 int32
	_ = v5048
	var v5049 int32
	_ = v5049
	var v5052 int32
	_ = v5052
	var v5058 int32
	_ = v5058
	var v5059 int32
	_ = v5059
	var v5063 int32
	_ = v5063
	var v5064 int32
	_ = v5064
	var v5065 int32
	_ = v5065
	var v5066 int32
	_ = v5066
	var v5079 int32
	_ = v5079
	var v5084 int32
	_ = v5084
	var v5125 int32
	_ = v5125
	var v5126 int32
	_ = v5126
	var v5127 int32
	_ = v5127
	var v5131 int32
	_ = v5131
	var v5132 int32
	_ = v5132
	var v5136 int32
	_ = v5136
	var v5138 int32
	_ = v5138
	var v5140 int32
	_ = v5140
	var v5144 int32
	_ = v5144
	var v5150 int32
	_ = v5150
	var v5152 int32
	_ = v5152
	var v5158 int32
	_ = v5158
	var v5163 int32
	_ = v5163
	var v5165 int32
	_ = v5165
	var v5166 int32
	_ = v5166
	var v5169 int32
	_ = v5169
	var v5177 int32
	_ = v5177
	var v5179 int32
	_ = v5179
	var v5182 int32
	_ = v5182
	var v5183 int32
	_ = v5183
	var v5186 int32
	_ = v5186
	var v5189 int32
	_ = v5189
	var v5190 int32
	_ = v5190
	var v5193 int32
	_ = v5193
	var v5194 int32
	_ = v5194
	var v5196 int32
	_ = v5196
	var v5197 int32
	_ = v5197
	var v5200 int32
	_ = v5200
	var v5206 int32
	_ = v5206
	var v5210 int32
	_ = v5210
	var v5215 int32
	_ = v5215
	var v5234 int32
	_ = v5234
	var v5257 int32
	_ = v5257
	var v5276 int32
	_ = v5276
	v3 = int32(0)
	v41 = m.G0
	v43 = v41 - int32(3280)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+60)) = int32(-1)
	v49 = m.G0
	v51 = v49 - int32(240)
	m.G0 = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v3 < v54 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v51 + int32(240)
	v4005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v4005 == int32(0) {
		v5234 = v3
		goto L567
	} else {
		goto L568
	}
L2:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+224))
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v60
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v62)
	if v57 <= v60 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v67 <= int32(0) {
		v256 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v66)+224))
	v287 = int32(0)
	v290 = int32(1)
	v292 = v290
	v294 = v287
	v295 = v290
	v297 = v287
	v298 = v3
	v299 = v287
	v313 = v3
	goto L16
L5:
	;
	v71 = v67 & int32(3)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v67) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v86 = v3
	v88 = v3
	v97 = v3
	goto L9
L7:
	;
	v160 = v3
	v162 = v3
	goto L8
L8:
	;
	v194 = v3
	v200 = v160
	v202 = v162
	goto L13
L9:
	;
	v119 = v72 + v86*int32(48)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+144))
	v121 = int32(5)
	v123 = int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v119)+48))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v119)+96))
	v143 = int32(base.Ui32(v120)>>(uint(v121)%32))&v123 + (int32(base.Ui32(v125)>>(uint(v121)%32))&v123 + v88 + int32(base.Ui32(v131)>>(uint(v121)%32))&v123 + int32(base.Ui32(v137)>>(uint(v121)%32))&v123)
	v144 = int32(4)
	v145 = v86 + v144
	v147 = v97 + v144
	if v147 != v67&int32(2147483644) {
		v86 = v145
		v88 = v143
		v97 = v147
		goto L9
	} else {
		goto L11
	}
L10:
	;
	if v71 == int32(0) {
		v256 = v143
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v160 = v145
	v162 = v143
	goto L8
L13:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v72+v200*int32(48))))
	v237 = int32(1)
	v239 = int32(base.Ui32(v234)>>(uint(int32(5))%32))&v237 + v202
	v243 = v194 + v237
	if v243 != v71 {
		v194 = v243
		v200 = v200 + v237
		v202 = v239
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v256 = v239
	goto L4
L15:
	;
	goto L14
L16:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v333 = base.I32_extend16_s(v295)
	v334 = base.I32_extend16_s(v292)
	if base.B2i32(v334 <= v333) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v285)+16)) = uint8(base.B2i32(int32(0) < v520))
	v557 = v520 + v256
	if v557 != 0 {
		goto L41
	} else {
		goto L42
	}
L18:
	;
	goto L17
L19:
	;
	v347 = v297
	v350 = v333
	goto L22
L20:
	;
	v413 = v295
	v415 = v297
	v418 = v333
	goto L21
L21:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (base.B2i32(v294 == v450)|v313)&int32(1) != 0 {
		v520 = v415
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v385 = v350<<(uint(int32(2))%32) - int32(4)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v66)+208))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389+v385)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v66)+212))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v392+v385)))
	v396 = F_get_opfamily_member(m, v391, v394, v394, int32(3))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v413 = v292
	v415 = v406
	v418 = v404
	goto L21
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385+(v51+int32(80))))) = v396
	if v396 == int32(0) {
		v520 = v298
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v403 = int32(1)
	v404 = v350 + v403
	v406 = v347 + v403
	if (v292+v297-v295)&int32(_a_F__bt_first_0) != v406&int32(_a_F__bt_first_0) {
		v347 = v406
		v350 = v404
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v457 = v332 + v294*int32(48)
	v458 = int32(*(*int16)(unsafe.Add(mBase, uint32(v457)+4)))
	if v458 <= v334 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457)+6)))
	v292 = v494
	v294 = v294 + int32(1)
	v295 = v495
	v297 = v496
	v298 = v415
	v299 = int32(base.Ui32(v498&int32(64))>>(uint(int32(6))%32)) | base.B2i32(v503 == int32(3)) | v497
	v313 = int32(base.Ui32(v498&int32(4)) >> (uint(int32(2)) % 32))
	goto L16
L30:
	;
	v494 = v292
	v495 = v413
	v496 = v415
	v497 = v299
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v299&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v494 = v489
	v495 = v413 + int32(1)
	v496 = v490
	v497 = int32(0)
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418<<(uint(int32(2))%32)+v51)+76)) = int32(0)
	v489 = v458
	v490 = v415
	goto L33
L35:
	;
	goto L36
L36:
	;
	v470 = v418<<(uint(int32(2))%32) - int32(4)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v66)+208))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v474+v470)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v66)+212))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v477+v470)))
	v481 = F_get_opfamily_member(m, v476, v479, v479, int32(3))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470+(v51+int32(80))))) = v481
	if v481 == int32(0) {
		v520 = v415
		goto L18
	} else {
		goto L38
	}
L38:
	;
	v486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457)+4)))
	v489 = v486
	v490 = v415 + int32(1)
	goto L33
L39:
	;
	v1643 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1627)+4)))
	if int32(0) < v1643 {
		goto L175
	} else {
		goto L176
	}
L40:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1601 = int32(0)
	v1614 = v1601
	v1620 = v1577
	v1622 = v1601
	v1627 = v1600
	goto L39
L41:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v559 = v558 + v520
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v285)+28))
	if v560 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v1557 == int32(0) {
		goto L1
	} else {
		goto L172
	}
L44:
	;
	v576 = int32(_a_F__bt_first_1)
	v577 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v575
	v582 = F_palloc(m, v559*int32(48))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L24
	} else {
		goto L50
	}
L45:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	v569 = F_AllocSetContextCreateInternal(m, v564, int32(_a_F__bt_first_2), int32(0), int32(1024), int32(_a_F__bt_first_3))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L24
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_MemoryContextReset(m, v560)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L24
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+28)) = v569
	v575 = v569
	goto L44
L49:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v285)+28))
	v575 = v574
	goto L44
L50:
	;
	v586 = F_palloc(m, v557<<(uint(int32(5))%32))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+20)) = v586
	v591 = F_palloc(m, v559*int32(28))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L24
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+24)) = v591
	v594 = int32(0)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v596 <= v594 {
		v1497 = v594
		v1512 = v594
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+12)) = v1497
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v577
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v1538 == int32(0) {
		goto L1
	} else {
		goto L167
	}
L54:
	;
	v603 = v594
	v607 = v520
	v616 = int32(-1)
	v617 = v3
	v618 = v594
	v620 = v3
	v621 = int32(1)
	v629 = v3
	goto L55
L55:
	;
	v641 = int32(48)
	v643 = v582 + v618*v641
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v647 = v644 + v617*v641
	*(*int32)(unsafe.Add(mBase, uint32(v51)+48)) = v51 + int32(52)
	if v607 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v1497 = v1453
	v1512 = v1468
	goto L53
L57:
	;
	v1492 = v617 + int32(1)
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1492 < v1493 {
		v603 = v1453
		v607 = v819
		v616 = v1466
		v617 = v1492
		v618 = v1468
		v620 = v1470
		v621 = v833
		v629 = v1479
		goto L55
	} else {
		goto L166
	}
L58:
	;
	v1431 = v815 << (uint(int32(5)) % 32)
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1431+v1432))) = v830
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1435+v1431)+4)) = v1437
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1439+v1431)+8)) = v1441
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1443+v1431)+12)) = int32(-1)
	v1447 = int32(1)
	v1453 = v815 + v1447
	v1466 = v1425
	v1468 = v830 + v1447
	v1470 = v1426
	v1479 = v1427
	goto L57
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L24
	} else {
		goto L163
	}
L60:
	;
	v853 = *(*int64)(unsafe.Add(mBase, uint32(v647)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v816)+40)) = v853
	v855 = *(*int64)(unsafe.Add(mBase, uint32(v647)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v816)+32)) = v855
	v857 = *(*int64)(unsafe.Add(mBase, uint32(v647)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v816)+24)) = v857
	v859 = *(*int64)(unsafe.Add(mBase, uint32(v647)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v816)+16)) = v859
	v861 = *(*int64)(unsafe.Add(mBase, uint32(v647)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v816)+8)) = v861
	v863 = *(*int64)(unsafe.Add(mBase, uint32(v647)))
	*(*int64)(unsafe.Add(mBase, uint32(v816))) = v863
	v865 = base.I32_wrap_i64(v863)
	if v865&int32(32) != 0 {
		goto L89
	} else {
		goto L90
	}
L61:
	;
	v815 = v603
	v816 = v643
	v819 = int32(0)
	v830 = v618
	v833 = v621
	goto L60
L62:
	;
	goto L63
L63:
	;
	v654 = v607
	v656 = v603
	v657 = v643
	v671 = v618
	v674 = v621
	goto L64
L64:
	;
	v694 = base.I32_extend16_s(v674)
	v695 = int32(*(*int16)(unsafe.Add(mBase, uint32(v647)+4)))
	if v695 < v694 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v815 = v805
	v816 = v810
	v819 = v722
	v830 = v807
	v833 = v803
	goto L60
L66:
	;
	v815 = v656
	v816 = v657
	v819 = v654
	v830 = v671
	v833 = v674
	goto L60
L67:
	;
	goto L68
L68:
	;
	v698 = v694 - int32(1)
	v700 = v698 << (uint(int32(2)) % 32)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v700+(v51+int32(80)))))
	if v704 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v815 = v656
	v816 = v657
	v819 = v654
	v830 = v671
	v833 = v674 + int32(1)
	goto L60
L70:
	;
	goto L71
L71:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v66)+248))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v709+v700)))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v66)+212))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v712+v700)))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v66)+208))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v715+v700)))
	v718 = F_get_opcode(m, v704)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	if v718 == int32(0) {
		goto L59
	} else {
		goto L73
	}
L73:
	;
	v722 = int32(0)
	F_ScanKeyEntryInitialize(m, v657, int32(_a_F__bt_first_4), v694, int32(3), v722, v711, v718, v722)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L24
	} else {
		goto L74
	}
L74:
	;
	v730 = v656 << (uint(int32(5)) % 32)
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v730+v731))) = v671
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v734+v730)+4)) = int32(-1)
	v738 = int32(1)
	v741 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286+v698<<(uint(v738)%32)))))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v66)+52))
	v747 = v744 + v698<<(uint(int32(4))%32)
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v747)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v742+v730)+16)) = uint16(v748)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v750+v730)+18)) = uint8(v752)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v754+v730)+19)) = uint8(v738)
	v759 = F_get_opfamily_proc(m, v717, v714, v714, int32(6))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L24
	} else {
		goto L75
	}
L75:
	;
	if v759 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v765 = F_palloc(m, int32(16))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L24
	} else {
		goto L79
	}
L77:
	;
	v783 = int32(0)
	goto L78
L78:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v784+v730)+20)) = v783
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v789 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v787+v730)+24)) = v789
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v791+v730)+28)) = v789
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v285)+24))
	F__bt_setup_array_cmp(m, l0, v657, v714, v795+v671*int32(28), v789)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L24
	} else {
		goto L84
	}
L79:
	;
	v767 = F_OidFunctionCall1Coll(m, v759, int32(0), v765)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L24
	} else {
		goto L80
	}
L80:
	;
	if v741&int32(1) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v765)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v765))) = v770
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v765)+12))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+12)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v765)+8)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v765)+4)) = v769
	goto L83
L82:
	;
	goto L83
L83:
	;
	v783 = v765
	goto L78
L84:
	;
	v802 = int32(1)
	v803 = v674 + v802
	v805 = v656 + v802
	v807 = v671 + v802
	v810 = v582 + v807*int32(48)
	v812 = v654 - v802
	if v812 != 0 {
		v654 = v812
		v656 = v805
		v657 = v810
		v671 = v807
		v674 = v803
		goto L64
	} else {
		goto L85
	}
L85:
	;
	goto L65
L86:
	;
	v1404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v1404)
	v1497 = v815
	v1512 = v830
	goto L53
L87:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v285)+24))
	F__bt_setup_array_cmp(m, l0, v816, v975, v1047+v830*int32(28), v51+int32(48))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L24
	} else {
		goto L114
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L24
	} else {
		goto L111
	}
L89:
	;
	if v865&int32(1) != 0 {
		goto L86
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v1453 = v815
	v1466 = v616
	v1468 = v830 + int32(1)
	v1470 = v620
	v1479 = v629
	goto L57
L92:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v816)+44))
	v871 = F_pg_detoast_datum(m, v870)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L24
	} else {
		goto L93
	}
L93:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v871)+12))
	F_get_typlenbyvalalign(m, v873, v51+int32(46), v51+int32(45), v51+int32(44))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v883 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+46)))
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+45)))
	v885 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51)+44)))
	F_deconstruct_array(m, v871, v883, v884, v885, v51+int32(36), v51+int32(32), v51+int32(40))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	v894 = int32(0)
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
	if v896 <= v894 {
		goto L86
	} else {
		goto L96
	}
L96:
	;
	v904 = v894
	v907 = v894
	v910 = v896
	goto L97
L97:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939+v907))))
	if v941 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v957 == int32(0) {
		goto L86
	} else {
		goto L103
	}
L99:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	v945 = int32(2)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v944+v907<<(uint(v945)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v944+v904<<(uint(v945)%32)))) = v951
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
	v957 = v904 + int32(1)
	v958 = v953
	goto L101
L100:
	;
	v957 = v904
	v958 = v910
	goto L101
L101:
	;
	v960 = v907 + int32(1)
	if v960 < v958 {
		v904 = v957
		v907 = v960
		v910 = v958
		goto L97
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v816)+8))
	if v964 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v66)+212))
	v968 = int32(*(*int16)(unsafe.Add(mBase, uint32(v816)+4)))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v967+v968<<(uint(int32(2))%32)-int32(4))))
	v975 = v974
	goto L106
L105:
	;
	v975 = v964
	goto L106
L106:
	;
	v976 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v816)+6)))
	switch v976 - int32(1) {
	case 0, 1:
		goto L108
	case 2:
		goto L87
	case 3, 4:
		goto L107
	default:
		goto L88
	}
L107:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	v988 = F__bt_find_extreme_element(m, l0, v816, v975, int32(1), v987, v957)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L24
	} else {
		goto L110
	}
L108:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	v981 = F__bt_find_extreme_element(m, l0, v816, v975, int32(5), v980, v957)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L24
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+44)) = v981
	v1453 = v815
	v1466 = v616
	v1468 = v830 + int32(1)
	v1470 = v620
	v1479 = v629
	goto L57
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+44)) = v988
	goto L91
L111:
	;
	v1037 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v816)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v1037
	F_errmsg_internal(m, int32(_a_F__bt_first_5), v51)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L24
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(2094), int32(_a_F__bt_first_7))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L24
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v1055 = int32(*(*int16)(unsafe.Add(mBase, uint32(v816)+4)))
	v1056 = int32(1)
	v1059 = int32(2)
	v1061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286+v1055<<(uint(v1056)%32)-v1059))))
	v1063 = v1061 & v1056
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v51)+48))
	if v1059 <= v957 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+212)) = v1064
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v816)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+220)) = uint8(v1063)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+216)) = v1069
	F_qsort_arg(m, v1067, v957, int32(4), int32(210), v51+int32(212))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L24
	} else {
		goto L118
	}
L116:
	;
	v1162 = v957
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+40)) = v1162
	v1198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v816)+4)))
	if v1198 != v620 {
		goto L132
	} else {
		goto L133
	}
L118:
	;
	v1087 = int32(0)
	v1088 = int32(1)
	goto L119
L119:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v51)+212))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v51)+216))
	v1124 = int32(2)
	v1126 = v1067 + v1088<<(uint(v1124)%32)
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1067+v1087<<(uint(v1124)%32))))
	v1132 = F_FunctionCall2Coll(m, v1122, v1123, v1127, v1131)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L24
	} else {
		goto L122
	}
L120:
	;
	v1162 = v1151 + int32(1)
	goto L117
L121:
	;
	v1153 = v1088 + int32(1)
	if v1153 != v957 {
		v1087 = v1151
		v1088 = v1153
		goto L119
	} else {
		goto L131
	}
L122:
	;
	if v1132 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v1137 = int32(1)
	goto L125
L124:
	;
	v1137 = int32(0) - v1132
	goto L125
L125:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+220)))
	if v1138 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v1139 = v1137
	goto L128
L127:
	;
	v1139 = v1132
	goto L128
L128:
	;
	if v1139 == int32(0) {
		v1151 = v1087
		goto L121
	} else {
		goto L129
	}
L129:
	;
	v1143 = v1087 + int32(1)
	if v1143 == v1088 {
		v1151 = v1088
		goto L121
	} else {
		goto L130
	}
L130:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	*(*int32)(unsafe.Add(mBase, uint32(v1067+v1143<<(uint(int32(2))%32)))) = v1148
	v1151 = v1143
	goto L121
L131:
	;
	goto L120
L132:
	;
	v1425 = v815
	v1426 = v1198
	v1427 = v975
	goto L58
L133:
	;
	goto L134
L134:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	v1203 = v1200 + v616<<(uint(int32(5))%32)
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+4))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+8))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	if v975 != v629 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+208))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1210+v620<<(uint(int32(2))%32)-int32(4))))
	v1218 = F_get_opfamily_proc(m, v1216, v629, v975, int32(1))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L24
	} else {
		goto L138
	}
L136:
	;
	v1228 = v1064
	goto L137
L137:
	;
	v1229 = int32(0)
	if base.B2i32(v1162 <= v1229)|base.B2i32(v1204 <= v1229) != 0 {
		v1340 = v1229
		goto L141
	} else {
		goto L142
	}
L138:
	;
	if v1218 == int32(0) {
		v1425 = v616
		v1426 = v620
		v1427 = v629
		goto L58
	} else {
		goto L139
	}
L139:
	;
	v1223 = v51 + int32(212)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+28))
	F_fmgr_info_cxt(m, v1218, v1223, v1224)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L24
	} else {
		goto L140
	}
L140:
	;
	v1228 = v1223
	goto L137
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1203)+4)) = v1340
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
	F_pfree(m, v1360)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L24
	} else {
		goto L161
	}
L142:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v816)+12))
	v1236 = int32(0)
	v1241 = v1236
	v1246 = v1236
	v1259 = v1229
	goto L143
L143:
	;
	v1280 = int32(2)
	v1282 = v1205 + v1241<<(uint(v1280)%32)
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1206+v1246<<(uint(v1280)%32))))
	v1288 = F_FunctionCall2Coll(m, v1228, v1235, v1283, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L24
	} else {
		goto L146
	}
L144:
	;
	v1340 = v1316
	goto L141
L145:
	;
	if v1204 <= v1314 {
		v1340 = v1316
		goto L141
	} else {
		goto L159
	}
L146:
	;
	if v1288 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v1293 = int32(1)
	goto L149
L148:
	;
	v1293 = int32(0) - v1288
	goto L149
L149:
	;
	if v1063 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v1294 = v1293
	goto L152
L151:
	;
	v1294 = v1288
	goto L152
L152:
	;
	if v1294 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	*(*int32)(unsafe.Add(mBase, uint32(v1205+v1259<<(uint(int32(2))%32)))) = v1300
	v1302 = int32(1)
	v1314 = v1241 + v1302
	v1315 = v1246 + v1302
	v1316 = v1259 + v1302
	goto L145
L154:
	;
	goto L155
L155:
	;
	if v1294 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v1314 = v1241 + int32(1)
	v1315 = v1246
	v1316 = v1259
	goto L145
L157:
	;
	goto L158
L158:
	;
	v1314 = v1241
	v1315 = v1246 + int32(1)
	v1316 = v1259
	goto L145
L159:
	;
	if v1315 < v1162 {
		v1241 = v1314
		v1246 = v1315
		v1259 = v1316
		goto L143
	} else {
		goto L160
	}
L160:
	;
	goto L144
L161:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+4))
	if v1363 != 0 {
		v1453 = v815
		v1466 = v616
		v1468 = v830
		v1470 = v620
		v1479 = v629
		goto L57
	} else {
		goto L162
	}
L162:
	;
	goto L86
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v704
	F_errmsg_internal(m, int32(_a_F__bt_first_8), v51+int32(16))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L24
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(1957), int32(_a_F__bt_first_7))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L24
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	goto L56
L167:
	;
	if v582 == int32(0) {
		v1577 = v1512
		goto L40
	} else {
		goto L168
	}
L168:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	v1546 = F_MemoryContextAlloc(m, v1543, v1512<<(uint(int32(2))%32))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L24
	} else {
		goto L169
	}
L169:
	;
	v1548 = int32(1)
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1512 <= v1549 {
		v1614 = v1548
		v1620 = v1512
		v1622 = v1546
		v1627 = v582
		goto L39
	} else {
		goto L170
	}
L170:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v1554 = F_repalloc(m, v1551, v1512*int32(48))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L24
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v1554
	v1614 = v1548
	v1620 = v1512
	v1622 = v1546
	v1627 = v582
	goto L39
L172:
	;
	v1577 = v57
	goto L40
L173:
	;
	v3960 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v3960)
	goto L1
L174:
	;
	v3954 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v3954)
	goto L1
L175:
	;
	if v1620 == int32(1) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3940 = m.ExcPending
	if v3940 != 0 {
		goto L24
	} else {
		goto L563
	}
L178:
	;
	v1651 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1627)+4)))
	v1652 = int32(1)
	v1657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59+v1651<<(uint(v1652)%32)-int32(2)))))
	v1659 = v1657 << (uint(int32(24)) % 32)
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1627)))
	if v1660&v1652 != 0 {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	goto L180
L180:
	;
	v1777 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+136)) = v1777
	v1779 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+128)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v51)+120)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v51)+112)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v51)+104)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v51)+96)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v51)+88)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v51)+80)) = v1779
	v1799 = v1777
	v1806 = v1777
	v1808 = v1777
	v1814 = v1777
	v1819 = int32(1)
	v1820 = v1777
	goto L211
L181:
	;
	if v1751 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L182:
	;
	v1751 = int32(1)
	goto L181
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1627)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1627)+6)) = uint16(v1736)
	goto L182
L184:
	;
	if v1663&int32(33554432) != 0 {
		goto L203
	} else {
		goto L204
	}
L185:
	;
	v1663 = v1660 | v1659
	*(*int32)(unsafe.Add(mBase, uint32(v1627))) = v1663
	if v1660&int32(64) != 0 {
		v1736 = int32(3)
		goto L183
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1673 = int32(0)
	if base.B2i32(v1657&int32(1) == v1673)|v1660&int32(16777216) == v1673 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	if v1660&int32(128) != 0 {
		goto L184
	} else {
		goto L189
	}
L189:
	;
	v1751 = int32(0)
	goto L181
L190:
	;
	v1681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1627)+6)))
	v1682 = int32(6) - v1681
	*(*uint16)(unsafe.Add(mBase, uint32(v1627)+6)) = uint16(v1682)
	goto L192
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1627))) = v1660 | v1659
	if v1660&int32(4) == int32(0) {
		goto L182
	} else {
		goto L193
	}
L193:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1627)+44))
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1690))))
	if v1691&int32(1) != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1751 = int32(0)
	goto L181
L195:
	;
	goto L196
L196:
	;
	v1695 = v1690
	goto L197
L197:
	;
	v1700 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1695)+4)))
	v1701 = int32(1)
	v1706 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59+v1700<<(uint(v1701)%32)-int32(2)))))
	v1711 = int32(0)
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1695)))
	if base.B2i32(v1706&v1701 == v1711)|v1713&int32(16777216) == v1711 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L182
L199:
	;
	v1720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1695)+6)))
	v1721 = int32(6) - v1720
	*(*uint16)(unsafe.Add(mBase, uint32(v1695)+6)) = uint16(v1721)
	goto L201
L200:
	;
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1695))) = v1713 | v1706<<(uint(int32(24))%32)
	if v1713&int32(16) == int32(0) {
		v1695 = v1695 + int32(48)
		goto L197
	} else {
		goto L202
	}
L202:
	;
	goto L198
L203:
	;
	v1735 = int32(5)
	goto L205
L204:
	;
	v1735 = int32(1)
	goto L205
L205:
	;
	v1736 = v1735
	goto L183
L206:
	;
	v1754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v1754)
	goto L208
L207:
	;
	goto L208
L208:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v1757 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1756)+40)) = v1757
	v1759 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1756)+32)) = v1759
	v1761 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1756)+24)) = v1761
	v1763 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1756)+16)) = v1763
	v1765 = *(*int64)(unsafe.Add(mBase, uint32(v1627)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1756)+8)) = v1765
	v1767 = *(*int64)(unsafe.Add(mBase, uint32(v1627)))
	*(*int64)(unsafe.Add(mBase, uint32(v1756))) = v1767
	v1769 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v1769
	v1771 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1627)+4)))
	if v1771 != v1769 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	F__bt_mark_scankey_required(m, v1774)
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L24
	} else {
		goto L210
	}
L210:
	;
	goto L1
L211:
	;
	v1841 = v1627 + v1814*int32(48)
	v1842 = base.B2i32(v1620 <= v1814)
	if v1620 <= v1814 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1949 = base.B2i32(v1814 == v1620)
	if v1949 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L214:
	;
	v1846 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1841)+4)))
	v1847 = int32(1)
	v1852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59+v1846<<(uint(v1847)%32)-int32(2)))))
	v1854 = v1852 << (uint(int32(24)) % 32)
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1841)))
	if v1855&v1847 != 0 {
		goto L219
	} else {
		goto L220
	}
L215:
	;
	if v1946 != 0 {
		goto L213
	} else {
		goto L240
	}
L216:
	;
	v1946 = int32(1)
	goto L215
L217:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1841)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1841)+6)) = uint16(v1931)
	goto L216
L218:
	;
	if v1858&int32(33554432) != 0 {
		goto L237
	} else {
		goto L238
	}
L219:
	;
	v1858 = v1855 | v1854
	*(*int32)(unsafe.Add(mBase, uint32(v1841))) = v1858
	if v1855&int32(64) != 0 {
		v1931 = int32(3)
		goto L217
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v1868 = int32(0)
	if base.B2i32(v1852&int32(1) == v1868)|v1855&int32(16777216) == v1868 {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	if v1855&int32(128) != 0 {
		goto L218
	} else {
		goto L223
	}
L223:
	;
	v1946 = int32(0)
	goto L215
L224:
	;
	v1876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1841)+6)))
	v1877 = int32(6) - v1876
	*(*uint16)(unsafe.Add(mBase, uint32(v1841)+6)) = uint16(v1877)
	goto L226
L225:
	;
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1841))) = v1855 | v1854
	if v1855&int32(4) == int32(0) {
		goto L216
	} else {
		goto L227
	}
L227:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+44))
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1885))))
	if v1886&int32(1) != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1946 = int32(0)
	goto L215
L229:
	;
	goto L230
L230:
	;
	v1890 = v1885
	goto L231
L231:
	;
	v1895 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1890)+4)))
	v1896 = int32(1)
	v1901 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59+v1895<<(uint(v1896)%32)-int32(2)))))
	v1906 = int32(0)
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1890)))
	if base.B2i32(v1901&v1896 == v1906)|v1908&int32(16777216) == v1906 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	goto L216
L233:
	;
	v1915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1890)+6)))
	v1916 = int32(6) - v1915
	*(*uint16)(unsafe.Add(mBase, uint32(v1890)+6)) = uint16(v1916)
	goto L235
L234:
	;
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1890))) = v1908 | v1901<<(uint(int32(24))%32)
	if v1908&int32(16) == int32(0) {
		v1890 = v1890 + int32(48)
		goto L231
	} else {
		goto L236
	}
L236:
	;
	goto L232
L237:
	;
	v1930 = int32(5)
	goto L239
L238:
	;
	v1930 = int32(1)
	goto L239
L239:
	;
	v1931 = v1930
	goto L217
L240:
	;
	v1947 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v1947)
	goto L1
L241:
	;
	v1799 = v2346
	v1806 = v2351
	v1808 = v2353
	v1814 = v1814 + int32(1)
	v1819 = v2355
	v1820 = v2367
	goto L211
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v2328
	if v1614 != 0 {
		goto L378
	} else {
		goto L379
	}
L243:
	;
	v2356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1841)+6)))
	v2358 = v2356 - int32(1)
	if v2356 == int32(3) {
		goto L353
	} else {
		goto L354
	}
L244:
	;
	v1952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1841)+4)))
	if v1952 == v1819&int32(_a_F__bt_first_0) {
		v2346 = v1799
		v2351 = v1806
		v2353 = v1808
		v2355 = v1819
		goto L243
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	if v1842 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L247:
	;
	goto L246
L248:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v51)+104))
	if v2233 == int32(0) {
		v2264 = v2232
		goto L331
	} else {
		goto L332
	}
L249:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v2200 = v2197 + v2196*int32(48)
	v2201 = *(*int64)(unsafe.Add(mBase, uint32(v2192)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2200)+40)) = v2201
	v2203 = *(*int64)(unsafe.Add(mBase, uint32(v2192)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2200)+32)) = v2203
	v2205 = *(*int64)(unsafe.Add(mBase, uint32(v2192)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2200)+24)) = v2205
	v2207 = *(*int64)(unsafe.Add(mBase, uint32(v2192)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2200)+16)) = v2207
	v2209 = *(*int64)(unsafe.Add(mBase, uint32(v2192)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2200)+8)) = v2209
	v2211 = *(*int64)(unsafe.Add(mBase, uint32(v2192)))
	*(*int64)(unsafe.Add(mBase, uint32(v2200))) = v2211
	if v1614 != 0 {
		goto L324
	} else {
		goto L325
	}
L250:
	;
	if v2184 == int32(0) {
		v2229 = v2188
		v2230 = v2185
		v2232 = v2187
		goto L248
	} else {
		goto L323
	}
L251:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v2157 = v2154 + v1808*int32(48)
	v2158 = *(*int64)(unsafe.Add(mBase, uint32(v2098)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2157)+40)) = v2158
	v2160 = *(*int64)(unsafe.Add(mBase, uint32(v2098)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2157)+32)) = v2160
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v2098)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2157)+24)) = v2162
	v2164 = *(*int64)(unsafe.Add(mBase, uint32(v2098)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2157)+16)) = v2164
	v2166 = *(*int64)(unsafe.Add(mBase, uint32(v2098)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2157)+8)) = v2166
	v2168 = *(*int64)(unsafe.Add(mBase, uint32(v2098)))
	*(*int64)(unsafe.Add(mBase, uint32(v2157))) = v2168
	if v1614 != 0 {
		goto L318
	} else {
		goto L319
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+128)) = int32(0)
	v2192 = v2101
	v2193 = base.B2i32(v1806 == base.I32_extend16_s(v1819)-int32(1))
	v2194 = v2097
	v2196 = v1808
	goto L249
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L24
	} else {
		goto L315
	}
L254:
	;
	v1958 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1841)+4)))
	if v1958 < base.I32_extend16_s(v1819) {
		goto L253
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v51)+104))
	if v1961 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L256
L258:
	;
	v2075 = int32(0)
	if base.B2i32(v2069 == v2075)|base.B2i32(v2070 == v2075) != 0 {
		goto L298
	} else {
		goto L299
	}
L259:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v51)+92))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	v2068 = v1799
	v2069 = v1965
	v2070 = v1964
	v2074 = v1806
	goto L258
L260:
	;
	goto L261
L261:
	;
	v1966 = int32(0)
	if v1614 == v1966 {
		v1988 = v1966
		v1989 = v1966
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v51)+128))
	if v1990 != 0 {
		goto L266
	} else {
		goto L267
	}
L263:
	;
	v1970 = int32(0)
	v1971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961))))
	if v1971&int32(32) == v1970 {
		v1988 = v1966
		v1989 = v1970
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v51)+112))
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v51)+108))
	v1988 = v1976 + v1977<<(uint(int32(5))%32) - int32(32)
	v1989 = v1983 + v1984*int32(28)
	goto L262
L265:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v51)+116))
	if v2008 != 0 {
		goto L274
	} else {
		goto L275
	}
L266:
	;
	v1991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961))))
	if v1991&int32(64) != 0 {
		goto L173
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v2007 = v1799
	goto L265
L269:
	;
	v1997 = F__bt_compare_scankey_args(m, l0, v1990, v1961, v1990, v1988, v1989, v51+int32(212))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L24
	} else {
		goto L270
	}
L270:
	;
	if v1997 == int32(0) {
		v2007 = int32(1)
		goto L265
	} else {
		goto L271
	}
L271:
	;
	v2001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+212)))
	if v2001 == int32(0) {
		goto L174
	} else {
		goto L272
	}
L272:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+128)) = int64(-4294967296)
	goto L268
L273:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v51)+92))
	if v2026 != 0 {
		goto L282
	} else {
		goto L283
	}
L274:
	;
	v2009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961))))
	if v2009&int32(64) != 0 {
		goto L173
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v2025 = v2007
	goto L273
L277:
	;
	v2015 = F__bt_compare_scankey_args(m, l0, v2008, v1961, v2008, v1988, v1989, v51+int32(212))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L24
	} else {
		goto L278
	}
L278:
	;
	if v2015 == int32(0) {
		v2025 = int32(1)
		goto L273
	} else {
		goto L279
	}
L279:
	;
	v2019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+212)))
	if v2019 != int32(1) {
		goto L174
	} else {
		goto L280
	}
L280:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+116)) = int64(-4294967296)
	goto L276
L281:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	if v2046 != 0 {
		goto L290
	} else {
		goto L291
	}
L282:
	;
	v2027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961))))
	if v2027&int32(64) != 0 {
		goto L173
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v2044 = int32(0)
	v2045 = v2025
	goto L281
L285:
	;
	v2033 = F__bt_compare_scankey_args(m, l0, v2026, v1961, v2026, v1988, v1989, v51+int32(212))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L24
	} else {
		goto L286
	}
L286:
	;
	if v2033 == int32(0) {
		v2044 = v2026
		v2045 = int32(1)
		goto L281
	} else {
		goto L287
	}
L287:
	;
	v2037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+212)))
	if v2037 != int32(1) {
		goto L174
	} else {
		goto L288
	}
L288:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+92)) = int64(-4294967296)
	goto L284
L289:
	;
	v2068 = v2064
	v2069 = v2065
	v2070 = v2044
	v2074 = v1806 + int32(1)
	goto L258
L290:
	;
	v2047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1961))))
	if v2047&int32(64) != 0 {
		goto L173
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v2064 = v2045
	v2065 = int32(0)
	goto L289
L293:
	;
	v2053 = F__bt_compare_scankey_args(m, l0, v2046, v1961, v2046, v1988, v1989, v51+int32(212))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L24
	} else {
		goto L294
	}
L294:
	;
	if v2053 == int32(0) {
		v2064 = int32(1)
		v2065 = v2046
		goto L289
	} else {
		goto L295
	}
L295:
	;
	v2057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+212)))
	if v2057 != int32(1) {
		goto L174
	} else {
		goto L296
	}
L296:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v51)+80)) = int64(-4294967296)
	goto L292
L297:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v51)+128))
	v2099 = int32(0)
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v51)+116))
	if base.B2i32(v2098 == v2099)|base.B2i32(v2101 == v2099) == v2099 {
		goto L305
	} else {
		goto L306
	}
L298:
	;
	v2097 = v2068
	goto L297
L299:
	;
	v2081 = int32(0)
	v2085 = F__bt_compare_scankey_args(m, l0, v2070, v2069, v2070, v2081, v2081, v51+int32(212))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L24
	} else {
		goto L300
	}
L300:
	;
	if v2085 == int32(0) {
		v2097 = int32(1)
		goto L297
	} else {
		goto L301
	}
L301:
	;
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+212)))
	if v2089 == int32(1) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+92)) = int32(0)
	goto L298
L303:
	;
	goto L304
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+80)) = int32(0)
	goto L298
L305:
	;
	v2107 = int32(0)
	v2111 = F__bt_compare_scankey_args(m, l0, v2101, v2098, v2101, v2107, v2107, v51+int32(212))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L24
	} else {
		goto L309
	}
L306:
	;
	goto L307
L307:
	;
	v2131 = base.B2i32(v1806 == base.I32_extend16_s(v1819)-int32(1))
	if v2098 != 0 {
		v2151 = v2131
		v2152 = v2101
		v2153 = v2097
		goto L251
	} else {
		goto L314
	}
L308:
	;
	v2151 = base.B2i32(v1806 == base.I32_extend16_s(v1819)-int32(1))
	v2152 = v2122
	v2153 = v2123
	goto L251
L309:
	;
	if v2111 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v2122 = v2101
	v2123 = int32(1)
	goto L308
L311:
	;
	goto L312
L312:
	;
	v2116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+212)))
	if v2116 != int32(1) {
		goto L252
	} else {
		goto L313
	}
L313:
	;
	v2119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+116)) = v2119
	v2122 = v2119
	v2123 = v2097
	goto L308
L314:
	;
	v2184 = v2101
	v2185 = v2097
	v2187 = v1808
	v2188 = v2131
	goto L250
L315:
	;
	F_errmsg_internal(m, int32(_a_F__bt_first_9), int32(0))
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L24
	} else {
		goto L316
	}
L316:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(343), int32(_a_F__bt_first_10))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L24
	} else {
		goto L317
	}
L317:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L318:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v51)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v1808<<(uint(int32(2))%32)))) = v2173
	goto L320
L319:
	;
	goto L320
L320:
	;
	v2176 = v1808 + int32(1)
	v2177 = int32(0)
	if v2151 == v2177 {
		v2184 = v2152
		v2185 = v2153
		v2187 = v2176
		v2188 = v2177
		goto L250
	} else {
		goto L321
	}
L321:
	;
	F__bt_mark_scankey_required(m, v2157)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L24
	} else {
		goto L322
	}
L322:
	;
	v2184 = v2152
	v2185 = v2153
	v2187 = v2176
	v2188 = int32(1)
	goto L250
L323:
	;
	v2192 = v2184
	v2193 = v2188
	v2194 = v2185
	v2196 = v2187
	goto L249
L324:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v51)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v2196<<(uint(int32(2))%32)))) = v2216
	goto L326
L325:
	;
	goto L326
L326:
	;
	v2219 = v2196 + int32(1)
	if v2193 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v2229 = int32(0)
	v2230 = v2194
	v2232 = v2219
	goto L248
L328:
	;
	goto L329
L329:
	;
	F__bt_mark_scankey_required(m, v2200)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L24
	} else {
		goto L330
	}
L330:
	;
	v2229 = int32(1)
	v2230 = v2194
	v2232 = v2219
	goto L248
L331:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v51)+92))
	if v2265 == int32(0) {
		v2296 = v2264
		goto L338
	} else {
		goto L339
	}
L332:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v2239 = v2236 + v2232*int32(48)
	v2240 = *(*int64)(unsafe.Add(mBase, uint32(v2233)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2239)+40)) = v2240
	v2242 = *(*int64)(unsafe.Add(mBase, uint32(v2233)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2239)+32)) = v2242
	v2244 = *(*int64)(unsafe.Add(mBase, uint32(v2233)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2239)+24)) = v2244
	v2246 = *(*int64)(unsafe.Add(mBase, uint32(v2233)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2239)+16)) = v2246
	v2248 = *(*int64)(unsafe.Add(mBase, uint32(v2233)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2239)+8)) = v2248
	v2250 = *(*int64)(unsafe.Add(mBase, uint32(v2233)))
	*(*int64)(unsafe.Add(mBase, uint32(v2239))) = v2250
	if v1614 != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v51)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v2232<<(uint(int32(2))%32)))) = v2255
	goto L335
L334:
	;
	goto L335
L335:
	;
	v2258 = v2232 + int32(1)
	if v2229 == int32(0) {
		v2264 = v2258
		goto L331
	} else {
		goto L336
	}
L336:
	;
	F__bt_mark_scankey_required(m, v2239)
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L24
	} else {
		goto L337
	}
L337:
	;
	v2264 = v2258
	goto L331
L338:
	;
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	if v2297 == int32(0) {
		v2328 = v2296
		goto L345
	} else {
		goto L346
	}
L339:
	;
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v2271 = v2268 + v2264*int32(48)
	v2272 = *(*int64)(unsafe.Add(mBase, uint32(v2265)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2271)+40)) = v2272
	v2274 = *(*int64)(unsafe.Add(mBase, uint32(v2265)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2271)+32)) = v2274
	v2276 = *(*int64)(unsafe.Add(mBase, uint32(v2265)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2271)+24)) = v2276
	v2278 = *(*int64)(unsafe.Add(mBase, uint32(v2265)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2271)+16)) = v2278
	v2280 = *(*int64)(unsafe.Add(mBase, uint32(v2265)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2271)+8)) = v2280
	v2282 = *(*int64)(unsafe.Add(mBase, uint32(v2265)))
	*(*int64)(unsafe.Add(mBase, uint32(v2271))) = v2282
	if v1614 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v51)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v2264<<(uint(int32(2))%32)))) = v2287
	goto L342
L341:
	;
	goto L342
L342:
	;
	v2290 = v2264 + int32(1)
	if v2229 == int32(0) {
		v2296 = v2290
		goto L338
	} else {
		goto L343
	}
L343:
	;
	F__bt_mark_scankey_required(m, v2271)
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L24
	} else {
		goto L344
	}
L344:
	;
	v2296 = v2290
	goto L338
L345:
	;
	if v1814 == v1620 {
		goto L242
	} else {
		goto L352
	}
L346:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v2303 = v2300 + v2296*int32(48)
	v2304 = *(*int64)(unsafe.Add(mBase, uint32(v2297)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2303)+40)) = v2304
	v2306 = *(*int64)(unsafe.Add(mBase, uint32(v2297)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2303)+32)) = v2306
	v2308 = *(*int64)(unsafe.Add(mBase, uint32(v2297)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2303)+24)) = v2308
	v2310 = *(*int64)(unsafe.Add(mBase, uint32(v2297)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2303)+16)) = v2310
	v2312 = *(*int64)(unsafe.Add(mBase, uint32(v2297)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2303)+8)) = v2312
	v2314 = *(*int64)(unsafe.Add(mBase, uint32(v2297)))
	*(*int64)(unsafe.Add(mBase, uint32(v2303))) = v2314
	if v1614 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v2296<<(uint(int32(2))%32)))) = v2319
	goto L349
L348:
	;
	goto L349
L349:
	;
	v2322 = v2296 + int32(1)
	if v2229 == int32(0) {
		v2328 = v2322
		goto L345
	} else {
		goto L350
	}
L350:
	;
	F__bt_mark_scankey_required(m, v2303)
	mBase = m.M
	v2326 = m.ExcPending
	if v2326 != 0 {
		goto L24
	} else {
		goto L351
	}
L351:
	;
	v2328 = v2322
	goto L345
L352:
	;
	v2329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1841)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+136)) = int32(0)
	v2332 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+128)) = v2332
	*(*int64)(unsafe.Add(mBase, uint32(v51)+120)) = v2332
	*(*int64)(unsafe.Add(mBase, uint32(v51)+112)) = v2332
	*(*int64)(unsafe.Add(mBase, uint32(v51)+104)) = v2332
	*(*int64)(unsafe.Add(mBase, uint32(v51)+96)) = v2332
	*(*int64)(unsafe.Add(mBase, uint32(v51)+88)) = v2332
	*(*int64)(unsafe.Add(mBase, uint32(v51)+80)) = v2332
	v2346 = v2230
	v2351 = v2074
	v2353 = v2328
	v2355 = v2329
	goto L243
L353:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v1841)))
	v2367 = int32(base.Ui32(v2361)>>(uint(int32(5))%32))&int32(1) + v1820
	goto L355
L354:
	;
	v2367 = v1820
	goto L355
L355:
	;
	v2372 = v51 + int32(80) + v2358*int32(12)
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2372)))
	if v2373 == int32(0) {
		v2458 = v2346
		v2461 = v2353
		goto L357
	} else {
		goto L358
	}
L356:
	;
	if v2358 != int32(2) {
		goto L241
	} else {
		goto L377
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2372)+8)) = v2367
	*(*int32)(unsafe.Add(mBase, uint32(v2372)+4)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v2372))) = v1841
	v1799 = v2458
	v1806 = v2351
	v1808 = v2461
	v1814 = v1814 + int32(1)
	v1819 = v2355
	v1820 = v2367
	goto L211
L358:
	;
	v2376 = int32(0)
	v2379 = base.B2i32(v2358 != int32(2))
	if v2379|(v1614^int32(1)) != 0 {
		v2414 = v2376
		v2415 = v2376
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v2418 = F__bt_compare_scankey_args(m, l0, v1841, v1841, v2373, v2415, v2414, v51+int32(212))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L24
	} else {
		goto L365
	}
L360:
	;
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1841))))
	if v2383&int32(32) != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	v2414 = v2386 + v1814*int32(28)
	v2415 = v2390 + v2367<<(uint(int32(5))%32) - int32(32)
	goto L359
L362:
	;
	goto L363
L363:
	;
	v2396 = int32(0)
	v2397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2373))))
	if v2397&int32(32) == v2396 {
		v2414 = v2376
		v2415 = v2396
		goto L359
	} else {
		goto L364
	}
L364:
	;
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2372)+4))
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v2372)+8))
	v2414 = v2402 + v2403*int32(28)
	v2415 = v2407 + v2408<<(uint(int32(5))%32) - int32(32)
	goto L359
L365:
	;
	if v2418 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v2420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+212)))
	if v2420 != int32(1) {
		goto L356
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v2431 = v2428 + v2353*int32(48)
	v2432 = *(*int64)(unsafe.Add(mBase, uint32(v2373)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2431)+40)) = v2432
	v2434 = *(*int64)(unsafe.Add(mBase, uint32(v2373)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2431)+32)) = v2434
	v2436 = *(*int64)(unsafe.Add(mBase, uint32(v2373)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2431)+24)) = v2436
	v2438 = *(*int64)(unsafe.Add(mBase, uint32(v2373)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2431)+16)) = v2438
	v2440 = *(*int64)(unsafe.Add(mBase, uint32(v2373)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2431)+8)) = v2440
	v2442 = *(*int64)(unsafe.Add(mBase, uint32(v2373)))
	*(*int64)(unsafe.Add(mBase, uint32(v2431))) = v2442
	if v1614 != 0 {
		goto L372
	} else {
		goto L373
	}
L369:
	;
	if v2358 != int32(2) {
		v2458 = v2346
		v2461 = v2353
		goto L357
	} else {
		goto L370
	}
L370:
	;
	v2423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2373))))
	if v2423&int32(32) == int32(0) {
		v2458 = v2346
		v2461 = v2353
		goto L357
	} else {
		goto L371
	}
L371:
	;
	goto L241
L372:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v2372)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v2353<<(uint(int32(2))%32)))) = v2447
	goto L374
L373:
	;
	goto L374
L374:
	;
	v2449 = int32(1)
	v2450 = v2353 + v2449
	if v2351 != base.I32_extend16_s(v2355)-v2449 {
		v2458 = v2449
		v2461 = v2450
		goto L357
	} else {
		goto L375
	}
L375:
	;
	F__bt_mark_scankey_required(m, v2431)
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L24
	} else {
		goto L376
	}
L376:
	;
	v2458 = v2449
	v2461 = v2450
	goto L357
L377:
	;
	v2469 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v2469)
	goto L1
L378:
	;
	v2472 = int32(0)
	v2474 = m.G0
	v2476 = v2474 - int32(16)
	m.G0 = v2476
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+12))
	if v2479 == v2472 {
		goto L381
	} else {
		goto L382
	}
L379:
	;
	goto L380
L380:
	;
	if v2230&int32(1) == int32(0) {
		goto L1
	} else {
		goto L457
	}
L381:
	;
	m.G0 = v2476 + int32(16)
	goto L380
L382:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+4))
	if int32(0) < v2482 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2486 = v2472
	v2497 = v2472
	goto L386
L384:
	;
	goto L385
L385:
	;
	v2882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v2882 == int32(0) {
		goto L381
	} else {
		goto L451
	}
L386:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+8))
	v2529 = v2526 + v2497*int32(48)
	v2530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2529)+6)))
	if v2530 != int32(3) {
		v2798 = v2486
		goto L388
	} else {
		goto L389
	}
L387:
	;
	goto L385
L388:
	;
	v2839 = v2497 + int32(1)
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+4))
	if v2839 < v2840 {
		v2486 = v2798
		v2497 = v2839
		goto L386
	} else {
		goto L450
	}
L389:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v2529)))
	if v2533&int32(32) == int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	if v2533&int32(_a_F__bt_first_11) != int32(_a_F__bt_first_12) {
		v2798 = v2486
		goto L388
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+24))
	v2560 = int32(28)
	v2562 = v2559 + v2497*v2560
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v1622+v2497<<(uint(int32(2))%32))))
	v2569 = v2559 + v2566*v2560
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v2569)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2562)+24)) = v2570
	v2572 = *(*int64)(unsafe.Add(mBase, uint32(v2569)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2562)+16)) = v2572
	v2574 = *(*int64)(unsafe.Add(mBase, uint32(v2569)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2562)+8)) = v2574
	v2576 = *(*int64)(unsafe.Add(mBase, uint32(v2569)))
	*(*int64)(unsafe.Add(mBase, uint32(v2562))) = v2576
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+12))
	if v2578 <= v2486 {
		v2798 = v2486
		goto L388
	} else {
		goto L398
	}
L393:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v2529)+8))
	if v2542 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v2551 = v2542
	goto L396
L395:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2485)+212))
	v2544 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2529)+4)))
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2543+v2544<<(uint(int32(2))%32)-int32(4))))
	v2551 = v2550
	goto L396
L396:
	;
	v2552 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+24))
	F__bt_setup_array_cmp(m, l0, v2529, v2551, v2552+v2497*int32(28), int32(0))
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L24
	} else {
		goto L397
	}
L397:
	;
	v2798 = v2486
	goto L388
L398:
	;
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+20))
	v2581 = v2486
	goto L399
L399:
	;
	v2623 = v2580 + v2581<<(uint(int32(5))%32)
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v2623)))
	if v2566 == v2624 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v2798 = v2578
	goto L388
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2623))) = v2497
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+4))
	switch v2627 + int32(1) {
	case 0:
		goto L405
	default:
		goto L404
	case 2:
		goto L406
	}
L402:
	;
	goto L403
L403:
	;
	v2796 = v2581 + int32(1)
	if v2796 != v2578 {
		v2581 = v2796
		goto L399
	} else {
		goto L449
	}
L404:
	;
	v2798 = v2581 + int32(1)
	goto L388
L405:
	;
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+20))
	if v2651 == int32(0) {
		goto L404
	} else {
		goto L409
	}
L406:
	;
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2529)))
	*(*int32)(unsafe.Add(mBase, uint32(v2529))) = v2630 & int32(-33)
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+8))
	v2635 = *(*int32)(unsafe.Add(mBase, uint32(v2634)))
	*(*int32)(unsafe.Add(mBase, uint32(v2529)+44)) = v2635
	v2637 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+12))
	v2639 = v2637 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2478)+12)) = v2639
	if v2639 == int32(0) {
		goto L381
	} else {
		goto L407
	}
L407:
	;
	v2645 = (v2639 - v2581) << (uint(int32(5)) % 32)
	if v2645 == int32(0) {
		v2798 = v2581
		goto L388
	} else {
		goto L408
	}
L408:
	;
	base.MemoryCopy(m, v2623, v2623+int32(32), v2645)
	v2798 = v2581
	goto L388
L409:
	;
	v2654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2623)+19)))
	if v2654 != 0 {
		goto L404
	} else {
		goto L410
	}
L410:
	;
	v2655 = int32(_a_F__bt_first_1)
	v2656 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v2658)+28))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v2659
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+28))
	if v2661 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+24))
	if v2721 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L412:
	;
	v2664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2661)+6)))
	if v2664 != int32(1) {
		goto L411
	} else {
		goto L413
	}
L413:
	;
	v2667 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2529)+4)))
	v2671 = v2667<<(uint(int32(2))%32) - int32(4)
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v2672)+208))
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v2671+v2673)))
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(v2661)+44))
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(v2661)+8))
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v2672)+212))
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v2678+v2671)))
	if v2677 != 0 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v2683 = base.B2i32(v2677 != v2680)
	goto L416
L415:
	;
	v2683 = int32(0)
	goto L416
L416:
	;
	if v2683 != 0 {
		goto L411
	} else {
		goto L417
	}
L417:
	;
	v2686 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+20))
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v2686)+8))
	v2688 = m.T0[v2687].(func(*base.Module, int32, int32, int32) int32)(m, v2672, v2676, v2476+int32(14))
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L24
	} else {
		goto L418
	}
L418:
	;
	v2690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476)+14)))
	if v2690 == int32(1) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2694 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2693))) = uint8(v2694)
	goto L411
L420:
	;
	goto L421
L421:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v2661)))
	if v2698&int32(16777216) != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v2701 = int32(4)
	goto L424
L423:
	;
	v2701 = int32(2)
	goto L424
L424:
	;
	v2702 = F_get_opfamily_member(m, v2675, v2680, v2680, v2701)
	mBase = m.M
	v2703 = m.ExcPending
	if v2703 != 0 {
		goto L24
	} else {
		goto L425
	}
L425:
	;
	if v2702 == int32(0) {
		goto L411
	} else {
		goto L426
	}
L426:
	;
	v2706 = F_get_opcode(m, v2702)
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		goto L24
	} else {
		goto L427
	}
L427:
	;
	if v2706 == int32(0) {
		goto L411
	} else {
		goto L428
	}
L428:
	;
	F_fmgr_info(m, v2706, v2661+int32(16))
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L24
	} else {
		goto L429
	}
L429:
	;
	v2714 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v2661)+6)) = uint16(v2714)
	*(*int32)(unsafe.Add(mBase, uint32(v2661)+44)) = v2688
	goto L411
L430:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v2656
	goto L404
L431:
	;
	v2724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2721)+6)))
	if v2724 != int32(5) {
		goto L430
	} else {
		goto L432
	}
L432:
	;
	v2727 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2529)+4)))
	v2731 = v2727<<(uint(int32(2))%32) - int32(4)
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2732)+208))
	v2735 = *(*int32)(unsafe.Add(mBase, uint32(v2731+v2733)))
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v2721)+44))
	v2737 = *(*int32)(unsafe.Add(mBase, uint32(v2721)+8))
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2732)+212))
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2738+v2731)))
	if v2737 != 0 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2743 = base.B2i32(v2737 != v2740)
	goto L435
L434:
	;
	v2743 = int32(0)
	goto L435
L435:
	;
	if v2743 != 0 {
		goto L430
	} else {
		goto L436
	}
L436:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2623)+20))
	v2747 = *(*int32)(unsafe.Add(mBase, uint32(v2746)+12))
	v2748 = m.T0[v2747].(func(*base.Module, int32, int32, int32) int32)(m, v2732, v2736, v2476+int32(15))
	mBase = m.M
	v2749 = m.ExcPending
	if v2749 != 0 {
		goto L24
	} else {
		goto L437
	}
L437:
	;
	v2750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476)+15)))
	if v2750 == int32(1) {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2753))) = uint8(v2754)
	goto L430
L439:
	;
	goto L440
L440:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v2721)))
	if v2758&int32(16777216) != 0 {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v2761 = int32(2)
	goto L443
L442:
	;
	v2761 = int32(4)
	goto L443
L443:
	;
	v2762 = F_get_opfamily_member(m, v2735, v2740, v2740, v2761)
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L24
	} else {
		goto L444
	}
L444:
	;
	if v2762 == int32(0) {
		goto L430
	} else {
		goto L445
	}
L445:
	;
	v2766 = F_get_opcode(m, v2762)
	mBase = m.M
	v2767 = m.ExcPending
	if v2767 != 0 {
		goto L24
	} else {
		goto L446
	}
L446:
	;
	if v2766 == int32(0) {
		goto L430
	} else {
		goto L447
	}
L447:
	;
	F_fmgr_info(m, v2766, v2721+int32(16))
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
		goto L24
	} else {
		goto L448
	}
L448:
	;
	v2774 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v2721)+6)) = uint16(v2774)
	*(*int32)(unsafe.Add(mBase, uint32(v2721)+44)) = v2748
	goto L430
L449:
	;
	goto L400
L450:
	;
	goto L387
L451:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+12))
	if v2885 < int32(33) {
		goto L381
	} else {
		goto L452
	}
L452:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L24
	} else {
		goto L453
	}
L453:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
		goto L24
	} else {
		goto L454
	}
L454:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v2478)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2476)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v2476))) = v2895
	F_errmsg_internal(m, int32(_a_F__bt_first_13), v2476)
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L24
	} else {
		goto L455
	}
L455:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(2373), int32(_a_F__bt_first_14))
	mBase = m.M
	v2906 = m.ExcPending
	if v2906 != 0 {
		goto L24
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
	v2994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v2994 != int32(1) {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v2997 = int32(0)
	v3004 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3005 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	v3006 = F_palloc0(m, v3005)
	mBase = m.M
	v3007 = m.ExcPending
	if v3007 != 0 {
		goto L24
	} else {
		goto L459
	}
L459:
	;
	v3008 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	if int32(0) < v3008 {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+8))
	v3012 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3011)+4)))
	v3015 = v2997
	v3016 = v3012
	v3019 = v2997
	v3021 = v2997
	v3028 = v2997
	v3029 = v2997
	v3032 = v2997
	goto L463
L461:
	;
	v3396 = v2997
	goto L462
L462:
	;
	v3423 = F_palloc(m, v3396*int32(48))
	mBase = m.M
	v3424 = m.ExcPending
	if v3424 != 0 {
		goto L24
	} else {
		goto L506
	}
L463:
	;
	v3053 = int32(0)
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+8))
	v3058 = v3055 + v3032*int32(48)
	v3059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3058)+4)))
	if v3059 == v3016&int32(_a_F__bt_first_0) {
		goto L468
	} else {
		goto L469
	}
L464:
	;
	v3396 = v3352
	goto L462
L465:
	;
	v3378 = v3032 + int32(1)
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	if v3378 < v3379 {
		v3015 = v3339
		v3016 = v3340
		v3019 = v3343
		v3021 = v3345
		v3028 = v3352
		v3029 = v3353
		v3032 = v3378
		goto L463
	} else {
		goto L505
	}
L466:
	;
	v3339 = v3066
	v3340 = v3068
	v3343 = v3067
	v3345 = v3070
	v3352 = v3312
	v3353 = int32(1)
	goto L465
L467:
	;
	v3293 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3006+v3032))) = uint8(v3293)
	v3339 = v3282
	v3340 = v3283
	v3343 = v3285
	v3345 = v3287
	v3352 = v3028 + v3293
	v3353 = v3288
	goto L465
L468:
	;
	v3063 = int32(1)
	if v3029&v3063 != 0 {
		v3282 = v3015
		v3283 = v3016
		v3285 = v3019
		v3287 = v3021
		v3288 = v3063
		goto L467
	} else {
		goto L471
	}
L469:
	;
	v3066 = v3032
	v3067 = v3053
	v3068 = v3059
	v3070 = v3053
	goto L470
L470:
	;
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v3058)))
	v3072 = int32(_a_F__bt_first_15)
	if v3071&v3072 == v3072 {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	v3066 = v3015
	v3067 = v3019
	v3068 = v3016
	v3070 = v3021
	goto L470
L472:
	;
	if v3032 <= v3066 {
		v3312 = v3028
		goto L466
	} else {
		goto L475
	}
L473:
	;
	goto L474
L474:
	;
	v3261 = int32(1)
	v3262 = int32(0)
	if (base.B2i32(v3071&int32(_a_F__bt_first_12) == v3262)|v3070)&v3261 == v3262 {
		goto L501
	} else {
		goto L502
	}
L475:
	;
	v3081 = (v3032 - v3066) & int32(3)
	if v3081 != 0 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v3085 = v3066
	v3093 = int32(0)
	v3097 = v3028
	goto L479
L477:
	;
	v3139 = v3066
	v3151 = v3028
	goto L478
L478:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v3066-v3032) {
		v3312 = v3151
		goto L466
	} else {
		goto L485
	}
L479:
	;
	v3122 = v3085 + v3006
	v3123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3122))))
	if v3123 == int32(0) {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v3139 = v3132
	v3151 = v3130
	goto L478
L481:
	;
	v3126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3122))) = uint8(v3126)
	v3130 = v3097 + v3126
	goto L483
L482:
	;
	v3130 = v3097
	goto L483
L483:
	;
	v3131 = int32(1)
	v3132 = v3085 + v3131
	v3134 = v3093 + v3131
	if v3134 != v3081 {
		v3085 = v3132
		v3093 = v3134
		v3097 = v3130
		goto L479
	} else {
		goto L484
	}
L484:
	;
	goto L480
L485:
	;
	v3182 = v3139
	v3194 = v3151
	goto L486
L486:
	;
	v3219 = v3182 + v3006
	v3220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3219))))
	if v3220 == int32(0) {
		goto L488
	} else {
		goto L489
	}
L487:
	;
	v3312 = v3257
	goto L466
L488:
	;
	v3223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3219))) = uint8(v3223)
	v3227 = v3194 + v3223
	goto L490
L489:
	;
	v3227 = v3194
	goto L490
L490:
	;
	v3229 = v3219 + int32(1)
	v3230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3229))))
	if v3230 == int32(0) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v3233 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3229))) = uint8(v3233)
	v3237 = v3227 + v3233
	goto L493
L492:
	;
	v3237 = v3227
	goto L493
L493:
	;
	v3239 = v3219 + int32(2)
	v3240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3239))))
	if v3240 == int32(0) {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v3243 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3239))) = uint8(v3243)
	v3247 = v3237 + v3243
	goto L496
L495:
	;
	v3247 = v3237
	goto L496
L496:
	;
	v3249 = v3219 + int32(3)
	v3250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3249))))
	if v3250 == int32(0) {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v3253 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3249))) = uint8(v3253)
	v3257 = v3247 + v3253
	goto L499
L498:
	;
	v3257 = v3247
	goto L499
L499:
	;
	v3259 = v3182 + int32(4)
	if v3032 != v3259 {
		v3182 = v3259
		v3194 = v3257
		goto L486
	} else {
		goto L500
	}
L500:
	;
	goto L487
L501:
	;
	v3339 = v3066
	v3340 = v3068
	v3343 = v3067
	v3345 = v3261
	v3352 = v3028
	v3353 = v3262
	goto L465
L502:
	;
	goto L503
L503:
	;
	v3272 = int32(0)
	if (v3067|base.B2i32(v3071&int32(_a_F__bt_first_16) == v3272))&int32(1) != 0 {
		v3282 = v3066
		v3283 = v3068
		v3285 = v3067
		v3287 = v3070
		v3288 = v3272
		goto L467
	} else {
		goto L504
	}
L504:
	;
	v3339 = v3066
	v3340 = v3068
	v3343 = int32(1)
	v3345 = v3070
	v3352 = v3028
	v3353 = v3262
	goto L465
L505:
	;
	goto L464
L506:
	;
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	v3429 = F_palloc(m, (v3425-v3396)*int32(48))
	mBase = m.M
	v3430 = m.ExcPending
	if v3430 != 0 {
		goto L24
	} else {
		goto L507
	}
L507:
	;
	v3432 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+12))
	if v3432 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3435 = F_palloc(m, v3396*int32(28))
	mBase = m.M
	v3436 = m.ExcPending
	if v3436 != 0 {
		goto L24
	} else {
		goto L511
	}
L509:
	;
	v3443 = int32(0)
	v3444 = v2997
	goto L510
L510:
	;
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	if v3445 <= int32(0) {
		goto L514
	} else {
		goto L515
	}
L511:
	;
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	v3441 = F_palloc(m, (v3437-v3396)*int32(28))
	mBase = m.M
	v3442 = m.ExcPending
	if v3442 != 0 {
		goto L24
	} else {
		goto L512
	}
L512:
	;
	v3443 = v3435
	v3444 = v3441
	goto L510
L513:
	;
	v3764 = v3731 * int32(48)
	if v3764 != 0 {
		goto L536
	} else {
		goto L537
	}
L514:
	;
	v3448 = int32(0)
	v3729 = v3448
	v3731 = v3448
	goto L513
L515:
	;
	goto L516
L516:
	;
	v3450 = int32(0)
	v3459 = v3450
	v3461 = v3450
	v3472 = v3450
	goto L517
L517:
	;
	v3493 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+8))
	v3496 = v3493 + v3472*int32(48)
	v3498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3006+v3472))))
	if v3498 == int32(0) {
		goto L520
	} else {
		goto L521
	}
L518:
	;
	v3729 = v3685
	v3731 = v3687
	goto L513
L519:
	;
	v3720 = v3472 + int32(1)
	v3721 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	if v3720 < v3721 {
		v3459 = v3685
		v3461 = v3687
		v3472 = v3720
		goto L517
	} else {
		goto L535
	}
L520:
	;
	v3503 = v3429 + v3461*int32(48)
	v3504 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3503)+40)) = v3504
	v3506 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v3503)+32)) = v3506
	v3508 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3503)+24)) = v3508
	v3510 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3503)+16)) = v3510
	v3512 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3503)+8)) = v3512
	v3514 = *(*int64)(unsafe.Add(mBase, uint32(v3496)))
	*(*int64)(unsafe.Add(mBase, uint32(v3503))) = v3514
	v3516 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+12))
	if v3516 != 0 {
		goto L523
	} else {
		goto L524
	}
L521:
	;
	goto L522
L522:
	;
	v3542 = v3423 + v3459*int32(48)
	v3543 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3542)+40)) = v3543
	v3545 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v3542)+32)) = v3545
	v3547 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3542)+24)) = v3547
	v3549 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3542)+16)) = v3549
	v3551 = *(*int64)(unsafe.Add(mBase, uint32(v3496)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3542)+8)) = v3551
	v3553 = *(*int64)(unsafe.Add(mBase, uint32(v3496)))
	*(*int64)(unsafe.Add(mBase, uint32(v3542))) = v3553
	v3555 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+12))
	if v3555 != 0 {
		goto L526
	} else {
		goto L527
	}
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v3472<<(uint(int32(2))%32)))) = v3461
	v3521 = int32(28)
	v3523 = v3444 + v3461*v3521
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+24))
	v3527 = v3524 + v3472*v3521
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3527)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3523)+24)) = v3528
	v3530 = *(*int64)(unsafe.Add(mBase, uint32(v3527)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3523)+16)) = v3530
	v3532 = *(*int64)(unsafe.Add(mBase, uint32(v3527)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3523)+8)) = v3532
	v3534 = *(*int64)(unsafe.Add(mBase, uint32(v3527)))
	*(*int64)(unsafe.Add(mBase, uint32(v3523))) = v3534
	goto L525
L524:
	;
	goto L525
L525:
	;
	v3685 = v3459
	v3687 = v3461 + int32(1)
	goto L519
L526:
	;
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1622+v3472<<(uint(int32(2))%32)))) = v3559 + (v3459 - v3396)
	v3563 = int32(28)
	v3565 = v3443 + v3459*v3563
	v3566 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+24))
	v3569 = v3566 + v3472*v3563
	v3570 = *(*int32)(unsafe.Add(mBase, uint32(v3569)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3565)+24)) = v3570
	v3572 = *(*int64)(unsafe.Add(mBase, uint32(v3569)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3565)+16)) = v3572
	v3574 = *(*int64)(unsafe.Add(mBase, uint32(v3569)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3565)+8)) = v3574
	v3576 = *(*int64)(unsafe.Add(mBase, uint32(v3569)))
	*(*int64)(unsafe.Add(mBase, uint32(v3565))) = v3576
	goto L528
L527:
	;
	goto L528
L528:
	;
	v3580 = *(*int32)(unsafe.Add(mBase, uint32(v3542)))
	*(*int32)(unsafe.Add(mBase, uint32(v3542))) = v3580 & int32(-196609)
	if v3580&int32(4) != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v3586 = *(*int32)(unsafe.Add(mBase, uint32(v3542)+44))
	v3590 = v3586
	goto L532
L530:
	;
	goto L531
L531:
	;
	v3685 = v3459 + int32(1)
	v3687 = v3461
	goto L519
L532:
	;
	v3627 = *(*int32)(unsafe.Add(mBase, uint32(v3590)))
	*(*int32)(unsafe.Add(mBase, uint32(v3590))) = v3627 & int32(-196609)
	if v3627&int32(16) == int32(0) {
		v3590 = v3590 + int32(48)
		goto L532
	} else {
		goto L534
	}
L533:
	;
	goto L531
L534:
	;
	goto L533
L535:
	;
	goto L518
L536:
	;
	v3765 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+8))
	base.MemoryCopy(m, v3765, v3429, v3764)
	goto L538
L537:
	;
	goto L538
L538:
	;
	v3768 = v3729 * int32(48)
	if v3768 != 0 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v3769 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+8))
	base.MemoryCopy(m, v3769+v3764, v3423, v3768)
	goto L541
L540:
	;
	goto L541
L541:
	;
	F_pfree(m, v3006)
	mBase = m.M
	v3773 = m.ExcPending
	if v3773 != 0 {
		goto L24
	} else {
		goto L542
	}
L542:
	;
	F_pfree(m, v3429)
	mBase = m.M
	v3775 = m.ExcPending
	if v3775 != 0 {
		goto L24
	} else {
		goto L543
	}
L543:
	;
	F_pfree(m, v3423)
	mBase = m.M
	v3777 = m.ExcPending
	if v3777 != 0 {
		goto L24
	} else {
		goto L544
	}
L544:
	;
	v3778 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+12))
	if v3778 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v3780 = v3731 * int32(28)
	if v3780 != 0 {
		goto L548
	} else {
		goto L549
	}
L546:
	;
	goto L547
L547:
	;
	goto L1
L548:
	;
	v3781 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+24))
	base.MemoryCopy(m, v3781, v3444, v3780)
	goto L550
L549:
	;
	goto L550
L550:
	;
	v3784 = v3729 * int32(28)
	if v3784 != 0 {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	v3785 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+24))
	base.MemoryCopy(m, v3785+v3780, v3443, v3784)
	goto L553
L552:
	;
	goto L553
L553:
	;
	v3788 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+12))
	if int32(0) < v3788 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v3795 = int32(0)
	goto L557
L555:
	;
	v3861 = v3788
	goto L556
L556:
	;
	v3886 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+20))
	F_pg_qsort(m, v3886, v3861, int32(32), int32(211))
	mBase = m.M
	v3890 = m.ExcPending
	if v3890 != 0 {
		goto L24
	} else {
		goto L560
	}
L557:
	;
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+20))
	v3835 = v3832 + v3795<<(uint(int32(5))%32)
	v3836 = *(*int32)(unsafe.Add(mBase, uint32(v3835)))
	v3840 = *(*int32)(unsafe.Add(mBase, uint32(v1622+v3836<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3835))) = v3840
	v3843 = v3795 + int32(1)
	v3844 = *(*int32)(unsafe.Add(mBase, uint32(v3004)+12))
	if v3843 < v3844 {
		v3795 = v3843
		goto L557
	} else {
		goto L559
	}
L558:
	;
	v3861 = v3844
	goto L556
L559:
	;
	goto L558
L560:
	;
	F_pfree(m, v3443)
	mBase = m.M
	v3892 = m.ExcPending
	if v3892 != 0 {
		goto L24
	} else {
		goto L561
	}
L561:
	;
	F_pfree(m, v3444)
	mBase = m.M
	v3894 = m.ExcPending
	if v3894 != 0 {
		goto L24
	} else {
		goto L562
	}
L562:
	;
	goto L547
L563:
	;
	F_errmsg_internal(m, int32(_a_F__bt_first_9), int32(0))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L24
	} else {
		goto L564
	}
L564:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(267), int32(_a_F__bt_first_10))
	mBase = m.M
	v3949 = m.ExcPending
	if v3949 != 0 {
		goto L24
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
	m.G0 = v43 + int32(3280)
	return v5276
L567:
	;
	F__bt_parallel_done(m, l0)
	mBase = m.M
	v5257 = m.ExcPending
	if v5257 != 0 {
		goto L24
	} else {
		goto L750
	}
L568:
	;
	v4008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v4008 == int32(0) {
		goto L569
	} else {
		goto L570
	}
L569:
	;
	v4018 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v4018 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L570:
	;
	v4016 = F__bt_parallel_seize(m, l0, v43+int32(60), v43+int32(56), int32(1))
	mBase = m.M
	v4017 = m.ExcPending
	if v4017 != 0 {
		goto L24
	} else {
		goto L571
	}
L571:
	;
	if v4016 != 0 {
		goto L569
	} else {
		goto L572
	}
L572:
	;
	v5276 = v3
	goto L566
L573:
	;
	v4024 = *(*int32)(unsafe.Add(mBase, uint32(v43)+60))
	if v4024 != int32(-1) {
		goto L577
	} else {
		goto L578
	}
L574:
	;
	v4021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+17)))
	if v4021 != 0 {
		goto L573
	} else {
		goto L575
	}
L575:
	;
	F__bt_start_array_keys(m, l0, l1)
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		goto L24
	} else {
		goto L576
	}
L576:
	;
	goto L573
L577:
	;
	v4027 = int32(1)
	v4028 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v4030 = F__bt_readnextpage(m, l0, v4024, v4028, l1, v4027)
	mBase = m.M
	v4031 = m.ExcPending
	if v4031 != 0 {
		goto L24
	} else {
		goto L580
	}
L578:
	;
	goto L579
L579:
	;
	v4051 = *(*int32)(unsafe.Add(mBase, uint32(v45)+272))
	if v4051 == int32(0) {
		goto L586
	} else {
		goto L587
	}
L580:
	;
	if v4030 == int32(0) {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v5276 = int32(0)
	goto L566
L582:
	;
	goto L583
L583:
	;
	v4035 = *(*int32)(unsafe.Add(mBase, uint32(v46)+100))
	v4038 = v46 + v4035*int32(10)
	v4039 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4038)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v4039)
	v4042 = v4038 + int32(104)
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v4042)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v4043
	v4045 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v4045 == int32(0) {
		v5276 = v4027
		goto L566
	} else {
		goto L584
	}
L584:
	;
	v4048 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4042)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v4045 + v4048
	v5276 = v4027
	goto L566
L585:
	;
	v4066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4066 != 0 {
		goto L591
	} else {
		goto L592
	}
L586:
	;
	v4054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+268)))
	if v4054 != int32(1) {
		goto L585
	} else {
		goto L589
	}
L587:
	;
	v4060 = v4051
	goto L588
L588:
	;
	v4061 = *(*int64)(unsafe.Add(mBase, uint32(v4060)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4060)+16)) = v4061 + int64(1)
	goto L585
L589:
	;
	F_pgstat_assoc_relation(m, v45)
	mBase = m.M
	v4058 = m.ExcPending
	if v4058 != 0 {
		goto L24
	} else {
		goto L590
	}
L590:
	;
	v4059 = *(*int32)(unsafe.Add(mBase, uint32(v45)+272))
	v4060 = v4059
	goto L588
L591:
	;
	v4067 = *(*int64)(unsafe.Add(mBase, uint32(v4066)))
	*(*int64)(unsafe.Add(mBase, uint32(v4066))) = v4067 + int64(1)
	goto L593
L592:
	;
	goto L593
L593:
	;
	v4071 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v4071 <= int32(0) {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v5125 = int32(0)
	v5126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5131 = F__bt_get_endpoint(m, v5127, v5125, base.B2i32(l1 == int32(-1)))
	mBase = m.M
	v5132 = m.ExcPending
	if v5132 != 0 {
		goto L24
	} else {
		goto L720
	}
L595:
	;
	v4077 = base.B2i32(l1 == int32(1))
	if l1 == int32(1) {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v4078 = int32(24)
	goto L598
L597:
	;
	v4078 = int32(28)
	goto L598
L598:
	;
	v4079 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v4082 = v4079
	v4084 = int32(1)
	v4087 = int32(3)
	v4100 = v4071
	v4107 = v3
	v4108 = v3
	v4111 = v3
	v4114 = v3
	goto L599
L599:
	;
	if v4114 < v4100 {
		goto L605
	} else {
		goto L606
	}
L600:
	;
	if v4453 == int32(0) {
		goto L594
	} else {
		goto L653
	}
L601:
	;
	goto L600
L602:
	;
	v4082 = v4082 + int32(48)
	v4084 = v4386
	v4087 = v4389
	v4100 = v4402
	v4107 = v4409
	v4108 = v4410
	v4111 = v4413
	v4114 = v4114 + int32(1)
	goto L599
L603:
	;
	v4375 = int32(0)
	v4376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4082)+6)))
	switch v4376 - int32(1) {
	case 0, 1:
		goto L647
	case 2:
		goto L646
	case 3, 4:
		goto L645
	default:
		v4386 = v4337
		v4389 = v4340
		v4402 = v4353
		v4409 = v4360
		v4410 = v4375
		v4413 = v4364
		goto L602
	}
L604:
	;
	if v4108 != 0 {
		v4386 = v4084
		v4389 = v4087
		v4402 = v4100
		v4409 = v4107
		v4410 = v4108
		v4413 = v4111
		goto L602
	} else {
		goto L643
	}
L605:
	;
	v4123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4082)+4)))
	if v4123 == v4084 {
		goto L604
	} else {
		goto L608
	}
L606:
	;
	goto L607
L607:
	;
	if v4108 != 0 {
		goto L610
	} else {
		goto L611
	}
L608:
	;
	goto L607
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+int32(1600)+v4107<<(uint(int32(2))%32)))) = v4291
	v4311 = int32(1)
	v4312 = v4107 + v4311
	v4313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4291)+6)))
	if v4313&int32(_a_F__bt_first_17) == v4311 {
		v4433 = v4313
		v4453 = v4312
		goto L601
	} else {
		goto L634
	}
L610:
	;
	v4125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4108)+2)))
	if v4125&int32(24) == int32(0) {
		v4291 = v4108
		goto L609
	} else {
		goto L613
	}
L611:
	;
	v4220 = v4111
	v4231 = int32(0)
	goto L612
L612:
	;
	v4232 = int32(0)
	if v4231|base.B2i32(v4220 == v4232) == v4232 {
		goto L623
	} else {
		goto L624
	}
L613:
	;
	v4130 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v4133 = base.I32_div_s(v4108-v4130, int32(48))
	v4134 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v4135 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v4155 = int32(0)
	goto L614
L614:
	;
	v4179 = v4134 + v4155<<(uint(int32(5))%32)
	v4180 = *(*int32)(unsafe.Add(mBase, uint32(v4179)))
	if v4133 != v4180 {
		goto L616
	} else {
		goto L617
	}
L615:
	;
	v4186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4179)+19)))
	if v4186 != 0 {
		goto L620
	} else {
		goto L621
	}
L616:
	;
	v4183 = v4155 + int32(1)
	if base.Ui32(v4183) < base.Ui32(v4135) {
		v4155 = v4183
		goto L614
	} else {
		goto L619
	}
L617:
	;
	goto L618
L618:
	;
	goto L615
L619:
	;
	goto L618
L620:
	;
	v4187 = v4111
	goto L622
L621:
	;
	v4187 = v4108
	goto L622
L622:
	;
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v4078+v4179)))
	v4220 = v4187
	v4231 = v4189
	goto L612
L623:
	;
	v4237 = *(*int32)(unsafe.Add(mBase, uint32(v4220)))
	if v4237&int32(33554432) != 0 {
		goto L627
	} else {
		goto L628
	}
L624:
	;
	goto L625
L625:
	;
	if v4231 == int32(0) {
		v4433 = v4087
		v4453 = v4107
		goto L601
	} else {
		goto L633
	}
L626:
	;
	v4251 = v43 - int32(-64) + v4107*int32(48)
	v4257 = int32(0)
	F_ScanKeyEntryInitialize(m, v4251, v4237&int32(50331648)|int32(129), base.I32_extend16_s(v4084), v4246, v4257, v4257, v4257, v4257)
	mBase = m.M
	v4262 = m.ExcPending
	if v4262 != 0 {
		goto L24
	} else {
		goto L632
	}
L627:
	;
	if v4077 == int32(0) {
		v4433 = v4087
		v4453 = v4107
		goto L601
	} else {
		goto L630
	}
L628:
	;
	goto L629
L629:
	;
	if l1 != int32(-1) {
		v4433 = v4087
		v4453 = v4107
		goto L601
	} else {
		goto L631
	}
L630:
	;
	v4246 = int32(5)
	goto L626
L631:
	;
	v4246 = int32(1)
	goto L626
L632:
	;
	v4291 = v4251
	goto L609
L633:
	;
	v4291 = v4231
	goto L609
L634:
	;
	v4318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4291)+2)))
	if v4318&int32(96) != 0 {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v4322 = int32(1)
	if l1 == v4322 {
		goto L638
	} else {
		goto L639
	}
L636:
	;
	goto L637
L637:
	;
	v4326 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v4326 <= v4114 {
		v4433 = v4313
		v4453 = v4312
		goto L601
	} else {
		goto L641
	}
L638:
	;
	v4325 = int32(5)
	goto L640
L639:
	;
	v4325 = v4322
	goto L640
L640:
	;
	v4433 = v4325
	v4453 = v4312
	goto L601
L641:
	;
	v4328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4082)+2)))
	if v4328&int32(3) == int32(0) {
		v4433 = v4313
		v4453 = v4312
		goto L601
	} else {
		goto L642
	}
L642:
	;
	v4333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4082)+4)))
	v4337 = v4333
	v4340 = v4313
	v4353 = v4326
	v4360 = v4312
	v4364 = int32(0)
	goto L603
L643:
	;
	v4337 = v4084
	v4340 = v4087
	v4353 = v4100
	v4360 = v4107
	v4364 = v4111
	goto L603
L644:
	;
	if v4364 != 0 {
		goto L650
	} else {
		goto L651
	}
L645:
	;
	if v4077 == int32(0) {
		goto L644
	} else {
		goto L649
	}
L646:
	;
	v4386 = v4337
	v4389 = v4340
	v4402 = v4353
	v4409 = v4360
	v4410 = v4082
	v4413 = v4364
	goto L602
L647:
	;
	if l1 != int32(-1) {
		goto L644
	} else {
		goto L648
	}
L648:
	;
	goto L646
L649:
	;
	v4386 = v4337
	v4389 = v4340
	v4402 = v4353
	v4409 = v4360
	v4410 = v4082
	v4413 = v4364
	goto L602
L650:
	;
	v4383 = v4364
	goto L652
L651:
	;
	v4383 = v4082
	goto L652
L652:
	;
	v4386 = v4337
	v4389 = v4340
	v4402 = v4353
	v4409 = v4360
	v4410 = v4375
	v4413 = v4383
	goto L602
L653:
	;
	if v4453 <= int32(0) {
		v4696 = v4453
		goto L667
	} else {
		goto L668
	}
L654:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5058 = m.ExcPending
	if v5058 != 0 {
		goto L24
	} else {
		goto L717
	}
L655:
	;
	v4999 = int32(0)
	v5002 = v43 + int32(1728)
	v5004 = v46 + int32(56)
	v5006 = F__bt_search(m, v45, v4999, v5002, v5004, int32(1))
	mBase = m.M
	v5007 = m.ExcPending
	if v5007 != 0 {
		goto L24
	} else {
		goto L703
	}
L656:
	;
	v4957 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+1731)) = uint16(v4957)
	goto L655
L657:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4945 = m.ExcPending
	if v4945 != 0 {
		goto L24
	} else {
		goto L700
	}
L658:
	;
	v4940 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+1731)) = uint16(v4940)
	goto L655
L659:
	;
	v4890 = v43 + int32(1728)
	F__bt_metaversion(m, v45, v4890, v4890|int32(1))
	mBase = m.M
	v4894 = m.ExcPending
	if v4894 != 0 {
		goto L24
	} else {
		goto L699
	}
L660:
	;
	v4887 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+1731)) = uint16(v4887)
	goto L655
L661:
	;
	v4837 = v43 + int32(1728)
	F__bt_metaversion(m, v45, v4837, v4837|int32(1))
	mBase = m.M
	v4841 = m.ExcPending
	if v4841 != 0 {
		goto L24
	} else {
		goto L698
	}
L662:
	;
	if l1 != int32(-1) {
		goto L656
	} else {
		goto L697
	}
L663:
	;
	v4830 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+1731)) = uint16(v4830)
	goto L655
L664:
	;
	v4780 = v43 + int32(1728)
	F__bt_metaversion(m, v45, v4780, v4780|int32(1))
	mBase = m.M
	v4784 = m.ExcPending
	if v4784 != 0 {
		goto L24
	} else {
		goto L696
	}
L665:
	;
	v4777 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+1731)) = uint16(v4777)
	goto L655
L666:
	;
	v4727 = v43 + int32(1728)
	F__bt_metaversion(m, v45, v4727, v4727|int32(1))
	mBase = m.M
	v4731 = m.ExcPending
	if v4731 != 0 {
		goto L24
	} else {
		goto L695
	}
L667:
	;
	v4712 = v43 + int32(1728)
	F__bt_metaversion(m, v45, v4712, v4712|int32(1))
	mBase = m.M
	v4716 = m.ExcPending
	if v4716 != 0 {
		goto L24
	} else {
		goto L694
	}
L668:
	;
	v4473 = v43 + int32(1744)
	v4491 = int32(0)
	goto L669
L669:
	;
	v4516 = v4491 << (uint(int32(2)) % 32)
	v4520 = *(*int32)(unsafe.Add(mBase, uint32(v4516+(v43+int32(1600)))))
	v4521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4520))))
	if v4521&int32(4) != 0 {
		goto L671
	} else {
		goto L672
	}
L670:
	;
	v4696 = v4453
	goto L667
L671:
	;
	v4526 = v4473 + v4491*int32(48)
	v4527 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+44))
	v4528 = *(*int64)(unsafe.Add(mBase, uint32(v4527)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4526)+40)) = v4528
	v4530 = *(*int64)(unsafe.Add(mBase, uint32(v4527)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4526)+32)) = v4530
	v4532 = *(*int64)(unsafe.Add(mBase, uint32(v4527)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4526)+24)) = v4532
	v4534 = *(*int64)(unsafe.Add(mBase, uint32(v4527)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4526)+16)) = v4534
	v4536 = *(*int64)(unsafe.Add(mBase, uint32(v4527)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4526)+8)) = v4536
	v4538 = *(*int64)(unsafe.Add(mBase, uint32(v4527)))
	*(*int64)(unsafe.Add(mBase, uint32(v4526))) = v4538
	v4558 = v4527
	v4565 = v4453
	goto L675
L672:
	;
	goto L673
L673:
	;
	v4617 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+8))
	if v4617 != 0 {
		goto L684
	} else {
		goto L685
	}
L674:
	;
	switch v4433&int32(_a_F__bt_first_0) - int32(2) {
	case 0:
		goto L666
	default:
		v4696 = v4565
		goto L667
	case 2:
		goto L659
	}
L675:
	;
	v4580 = *(*int32)(unsafe.Add(mBase, uint32(v4558)+48))
	if v4580&int32(1) != 0 {
		goto L674
	} else {
		goto L677
	}
L676:
	;
	switch v4433&int32(_a_F__bt_first_0) - int32(1) {
	case 0:
		goto L664
	default:
		v4696 = v4565
		goto L667
	case 4:
		goto L661
	}
L677:
	;
	if v4580&int32(_a_F__bt_first_15) != 0 {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	v4585 = int32(48)
	v4587 = v4473 + v4565*v4585
	v4589 = v4558 + v4585
	v4590 = *(*int64)(unsafe.Add(mBase, uint32(v4589)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4587)+40)) = v4590
	v4592 = *(*int64)(unsafe.Add(mBase, uint32(v4589)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4587)+32)) = v4592
	v4594 = *(*int64)(unsafe.Add(mBase, uint32(v4589)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4587)+24)) = v4594
	v4596 = *(*int64)(unsafe.Add(mBase, uint32(v4589)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4587)+16)) = v4596
	v4598 = *(*int64)(unsafe.Add(mBase, uint32(v4589)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4587)+8)) = v4598
	v4600 = *(*int64)(unsafe.Add(mBase, uint32(v4589)))
	*(*int64)(unsafe.Add(mBase, uint32(v4587))) = v4600
	v4603 = v4565 + int32(1)
	v4604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4589))))
	if v4604&int32(16) == int32(0) {
		v4558 = v4589
		v4565 = v4603
		goto L675
	} else {
		goto L681
	}
L679:
	;
	goto L680
L680:
	;
	goto L676
L681:
	;
	v4696 = v4603
	goto L667
L682:
	;
	v4669 = v4491 + int32(1)
	if v4669 != v4453 {
		v4491 = v4669
		goto L669
	} else {
		goto L693
	}
L683:
	;
	v4647 = *(*int32)(unsafe.Add(mBase, uint32(v45)+208))
	v4649 = *(*int32)(unsafe.Add(mBase, uint32(v4647+v4516)))
	v4651 = F_get_opfamily_proc(m, v4649, v4620, v4617, int32(1))
	mBase = m.M
	v4652 = m.ExcPending
	if v4652 != 0 {
		goto L24
	} else {
		goto L690
	}
L684:
	;
	v4618 = *(*int32)(unsafe.Add(mBase, uint32(v45)+212))
	v4620 = *(*int32)(unsafe.Add(mBase, uint32(v4618+v4516)))
	if v4617 != v4620 {
		goto L683
	} else {
		goto L687
	}
L685:
	;
	goto L686
L686:
	;
	v4623 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4520)+4)))
	v4625 = F_index_getprocinfo(m, v45, v4623, int32(1))
	mBase = m.M
	v4626 = m.ExcPending
	if v4626 != 0 {
		goto L24
	} else {
		goto L688
	}
L687:
	;
	goto L686
L688:
	;
	v4629 = v4473 + v4491*int32(48)
	v4630 = *(*int32)(unsafe.Add(mBase, uint32(v4520)))
	v4631 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4520)+4)))
	v4632 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+8))
	v4633 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+12))
	v4634 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v4629)+44)) = v4634
	*(*int32)(unsafe.Add(mBase, uint32(v4629)+12)) = v4633
	*(*int32)(unsafe.Add(mBase, uint32(v4629)+8)) = v4632
	v4638 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4629)+6)) = uint16(v4638)
	*(*uint16)(unsafe.Add(mBase, uint32(v4629)+4)) = uint16(v4631)
	*(*int32)(unsafe.Add(mBase, uint32(v4629))) = v4630
	v4645 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	F_fmgr_info_copy(m, v4629+int32(16), v4625, v4645)
	mBase = m.M
	goto L689
L689:
	;
	goto L682
L690:
	;
	if v4651 == int32(0) {
		goto L654
	} else {
		goto L691
	}
L691:
	;
	v4658 = *(*int32)(unsafe.Add(mBase, uint32(v4520)))
	v4659 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4520)+4)))
	v4661 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+8))
	v4662 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+12))
	v4663 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+44))
	F_ScanKeyEntryInitialize(m, v4473+v4491*int32(48), v4658, v4659, int32(0), v4661, v4662, v4651, v4663)
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		goto L24
	} else {
		goto L692
	}
L692:
	;
	goto L682
L693:
	;
	goto L670
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1740)) = v4696
	v4718 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1736)) = v4718
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+1730)) = uint8(v4718)
	v4723 = v4433 & int32(_a_F__bt_first_0)
	switch v4723 - int32(1) {
	case 0:
		goto L665
	case 1:
		goto L663
	case 2:
		goto L662
	case 3:
		goto L660
	case 4:
		goto L658
	default:
		goto L657
	}
L695:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1740)) = v4565
	v4733 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1736)) = v4733
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+1730)) = uint8(v4733)
	goto L665
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1740)) = v4565
	v4786 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1736)) = v4786
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+1730)) = uint8(v4786)
	goto L663
L697:
	;
	v4834 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+1731)) = uint16(v4834)
	goto L655
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1740)) = v4565
	v4843 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1736)) = v4843
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+1730)) = uint8(v4843)
	goto L660
L699:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1740)) = v4565
	v4896 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+1736)) = v4896
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+1730)) = uint8(v4896)
	goto L658
L700:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v4723
	F_errmsg_internal(m, int32(_a_F__bt_first_18), v43+int32(16))
	mBase = m.M
	v4951 = m.ExcPending
	if v4951 != 0 {
		goto L24
	} else {
		goto L701
	}
L701:
	;
	F_errfinish(m, int32(_a_F__bt_first_19), int32(1511), int32(_a_F__bt_first_20))
	mBase = m.M
	v4956 = m.ExcPending
	if v4956 != 0 {
		goto L24
	} else {
		goto L702
	}
L702:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L703:
	;
	F__bt_freestack(m, v5006)
	mBase = m.M
	v5009 = m.ExcPending
	if v5009 != 0 {
		goto L24
	} else {
		goto L704
	}
L704:
	;
	v5010 = *(*int32)(unsafe.Add(mBase, uint32(v46)+56))
	if v5010 == int32(0) {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v5014 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[1]))
	if v5014 != int32(3) {
		v5234 = v4999
		goto L567
	} else {
		goto L708
	}
L706:
	;
	v5029 = v5010
	goto L707
L707:
	;
	v5032 = F__bt_binsrch(m, v45, v43+int32(1728), v5029)
	mBase = m.M
	v5033 = m.ExcPending
	if v5033 != 0 {
		goto L24
	} else {
		goto L713
	}
L708:
	;
	v5017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockRelation(m, v45, v5017)
	mBase = m.M
	v5019 = m.ExcPending
	if v5019 != 0 {
		goto L24
	} else {
		goto L709
	}
L709:
	;
	v5022 = F__bt_search(m, v45, int32(0), v5002, v5004, int32(1))
	mBase = m.M
	v5023 = m.ExcPending
	if v5023 != 0 {
		goto L24
	} else {
		goto L710
	}
L710:
	;
	F__bt_freestack(m, v5022)
	mBase = m.M
	v5025 = m.ExcPending
	if v5025 != 0 {
		goto L24
	} else {
		goto L711
	}
L711:
	;
	v5026 = *(*int32)(unsafe.Add(mBase, uint32(v5004)))
	if v5026 == int32(0) {
		v5234 = v4999
		goto L567
	} else {
		goto L712
	}
L712:
	;
	v5029 = v5026
	goto L707
L713:
	;
	v5034 = F__bt_readfirstpage(m, l0, v5032, l1)
	mBase = m.M
	v5035 = m.ExcPending
	if v5035 != 0 {
		goto L24
	} else {
		goto L714
	}
L714:
	;
	if v5034 == int32(0) {
		v5276 = v4999
		goto L566
	} else {
		goto L715
	}
L715:
	;
	v5038 = *(*int32)(unsafe.Add(mBase, uint32(v46)+100))
	v5041 = v46 + v5038*int32(10)
	v5042 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5041)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v5042)
	v5045 = v5041 + int32(104)
	v5046 = *(*int32)(unsafe.Add(mBase, uint32(v5045)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v5046
	v5048 = int32(1)
	v5049 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v5049 == int32(0) {
		v5276 = v5048
		goto L566
	} else {
		goto L716
	}
L716:
	;
	v5052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5045)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v5049 + v5052
	v5276 = v5048
	goto L566
L717:
	;
	v5059 = *(*int32)(unsafe.Add(mBase, uint32(v45)+212))
	v5063 = *(*int32)(unsafe.Add(mBase, uint32(v5059+v4491<<(uint(int32(2))%32))))
	v5064 = *(*int32)(unsafe.Add(mBase, uint32(v4520)+8))
	v5065 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4520)+4)))
	v5066 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v5066 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+44)) = v5065
	*(*int32)(unsafe.Add(mBase, uint32(v43)+40)) = v5064
	*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v5063
	*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = int32(1)
	F_errmsg_internal(m, int32(_a_F__bt_first_21), v43+int32(32))
	mBase = m.M
	v5079 = m.ExcPending
	if v5079 != 0 {
		goto L24
	} else {
		goto L718
	}
L718:
	;
	F_errfinish(m, int32(_a_F__bt_first_19), int32(1430), int32(_a_F__bt_first_20))
	mBase = m.M
	v5084 = m.ExcPending
	if v5084 != 0 {
		goto L24
	} else {
		goto L719
	}
L719:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L720:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5126)+56)) = v5131
	if v5131 == int32(0) {
		goto L721
	} else {
		goto L722
	}
L721:
	;
	v5136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockRelation(m, v5127, v5136)
	mBase = m.M
	v5138 = m.ExcPending
	if v5138 != 0 {
		goto L24
	} else {
		goto L724
	}
L722:
	;
	goto L723
L723:
	;
	if v5131 < int32(0) {
		goto L727
	} else {
		goto L728
	}
L724:
	;
	F__bt_parallel_done(m, l0)
	mBase = m.M
	v5140 = m.ExcPending
	if v5140 != 0 {
		goto L24
	} else {
		goto L725
	}
L725:
	;
	v5276 = v5125
	goto L566
L726:
	;
	if l1 == int32(1) {
		goto L732
	} else {
		goto L733
	}
L727:
	;
	v5144 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[2]))
	v5150 = *(*int32)(unsafe.Add(mBase, uint32(v5144+(v5131^int32(-1))<<(uint(int32(2))%32))))
	v5158 = v5150
	goto L726
L728:
	;
	goto L729
L729:
	;
	v5152 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[3]))
	v5158 = v5152 + v5131<<(uint(int32(13))%32) + int32(-8192)
	goto L726
L730:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5206 = m.ExcPending
	if v5206 != 0 {
		goto L24
	} else {
		goto L747
	}
L731:
	;
	v5182 = F__bt_readfirstpage(m, l0, v5179&int32(_a_F__bt_first_0), l1)
	mBase = m.M
	v5183 = m.ExcPending
	if v5183 != 0 {
		goto L24
	} else {
		goto L742
	}
L732:
	;
	v5163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5158)+16)))
	v5165 = *(*int32)(unsafe.Add(mBase, uint32(v5158+v5163)+4))
	if v5165 != 0 {
		goto L735
	} else {
		goto L736
	}
L733:
	;
	goto L734
L734:
	;
	if l1 != int32(-1) {
		goto L730
	} else {
		goto L738
	}
L735:
	;
	v5166 = int32(2)
	goto L737
L736:
	;
	v5166 = int32(1)
	goto L737
L737:
	;
	v5179 = v5166
	goto L731
L738:
	;
	v5169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5158)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v5169) {
		goto L739
	} else {
		goto L740
	}
L739:
	;
	v5177 = int32(base.Ui32(v5169+int32(_a_F__bt_first_22)) >> (uint(int32(2)) % 32))
	goto L741
L740:
	;
	v5177 = int32(0)
	goto L741
L741:
	;
	v5179 = v5177
	goto L731
L742:
	;
	if v5182 == int32(0) {
		goto L743
	} else {
		goto L744
	}
L743:
	;
	v5276 = v5125
	goto L566
L744:
	;
	goto L745
L745:
	;
	v5186 = *(*int32)(unsafe.Add(mBase, uint32(v5126)+100))
	v5189 = v5126 + v5186*int32(10)
	v5190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5189)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v5190)
	v5193 = v5189 + int32(104)
	v5194 = *(*int32)(unsafe.Add(mBase, uint32(v5193)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v5194
	v5196 = int32(1)
	v5197 = *(*int32)(unsafe.Add(mBase, uint32(v5126)+44))
	if v5197 == int32(0) {
		v5276 = v5196
		goto L566
	} else {
		goto L746
	}
L746:
	;
	v5200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5193)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v5197 + v5200
	v5276 = v5196
	goto L566
L747:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = l1
	F_errmsg_internal(m, int32(_a_F__bt_first_23), v43)
	mBase = m.M
	v5210 = m.ExcPending
	if v5210 != 0 {
		goto L24
	} else {
		goto L748
	}
L748:
	;
	F_errfinish(m, int32(_a_F__bt_first_19), int32(2744), int32(_a_F__bt_first_24))
	mBase = m.M
	v5215 = m.ExcPending
	if v5215 != 0 {
		goto L24
	} else {
		goto L749
	}
L749:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L750:
	;
	v5276 = v5234
	goto L566
}
