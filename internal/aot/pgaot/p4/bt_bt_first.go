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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v269 int32
	_ = v269
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v527 int32
	_ = v527
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int64
	_ = v779
	var v780 int32
	_ = v780
	var v781 int64
	_ = v781
	var v782 int64
	_ = v782
	var v785 int64
	_ = v785
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v864 int64
	_ = v864
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v870 int64
	_ = v870
	var v872 int64
	_ = v872
	var v874 int64
	_ = v874
	var v876 int64
	_ = v876
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v972 int64
	_ = v972
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v989 int64
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1050 int64
	_ = v1050
	var v1057 int32
	_ = v1057
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1113 int64
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1119 int64
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int64
	_ = v1260
	var v1264 int64
	_ = v1264
	var v1265 int64
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1282 int64
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1421 int64
	_ = v1421
	var v1425 int64
	_ = v1425
	var v1426 int64
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1439 int64
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1465 int32
	_ = v1465
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1545 int32
	_ = v1545
	var v1550 int32
	_ = v1550
	var v1556 int32
	_ = v1556
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1649 int32
	_ = v1649
	var v1656 int32
	_ = v1656
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1715 int32
	_ = v1715
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1788 int32
	_ = v1788
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1818 int32
	_ = v1818
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int64
	_ = v1902
	var v1904 int64
	_ = v1904
	var v1906 int64
	_ = v1906
	var v1908 int64
	_ = v1908
	var v1910 int64
	_ = v1910
	var v1912 int64
	_ = v1912
	var v1914 int64
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1926 int64
	_ = v1926
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1965 int32
	_ = v1965
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2016 int32
	_ = v2016
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2049 int32
	_ = v2049
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2185 int32
	_ = v2185
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2229 int32
	_ = v2229
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2279 int32
	_ = v2279
	var v2283 int32
	_ = v2283
	var v2287 int32
	_ = v2287
	var v2292 int32
	_ = v2292
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
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
	var v2316 int64
	_ = v2316
	var v2318 int64
	_ = v2318
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2350 int64
	_ = v2350
	var v2352 int64
	_ = v2352
	var v2354 int64
	_ = v2354
	var v2356 int64
	_ = v2356
	var v2358 int64
	_ = v2358
	var v2360 int64
	_ = v2360
	var v2362 int64
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2375 int32
	_ = v2375
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2389 int64
	_ = v2389
	var v2391 int64
	_ = v2391
	var v2393 int64
	_ = v2393
	var v2395 int64
	_ = v2395
	var v2397 int64
	_ = v2397
	var v2399 int64
	_ = v2399
	var v2401 int64
	_ = v2401
	var v2406 int32
	_ = v2406
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2419 int32
	_ = v2419
	var v2422 int32
	_ = v2422
	var v2423 int64
	_ = v2423
	var v2425 int64
	_ = v2425
	var v2427 int64
	_ = v2427
	var v2429 int64
	_ = v2429
	var v2431 int64
	_ = v2431
	var v2433 int64
	_ = v2433
	var v2435 int64
	_ = v2435
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2456 int32
	_ = v2456
	var v2457 int64
	_ = v2457
	var v2459 int64
	_ = v2459
	var v2461 int64
	_ = v2461
	var v2463 int64
	_ = v2463
	var v2465 int64
	_ = v2465
	var v2467 int64
	_ = v2467
	var v2469 int64
	_ = v2469
	var v2474 int32
	_ = v2474
	var v2477 int32
	_ = v2477
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2487 int64
	_ = v2487
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2505 int32
	_ = v2505
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2522 int32
	_ = v2522
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2531 int32
	_ = v2531
	var v2534 int32
	_ = v2534
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2545 int32
	_ = v2545
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2578 int32
	_ = v2578
	var v2583 int32
	_ = v2583
	var v2586 int32
	_ = v2586
	var v2587 int64
	_ = v2587
	var v2589 int64
	_ = v2589
	var v2591 int64
	_ = v2591
	var v2593 int64
	_ = v2593
	var v2595 int64
	_ = v2595
	var v2597 int64
	_ = v2597
	var v2599 int64
	_ = v2599
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2626 int32
	_ = v2626
	var v2629 int32
	_ = v2629
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2639 int32
	_ = v2639
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2652 int32
	_ = v2652
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2691 int32
	_ = v2691
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2730 int64
	_ = v2730
	var v2732 int64
	_ = v2732
	var v2734 int64
	_ = v2734
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2786 int32
	_ = v2786
	var v2789 int32
	_ = v2789
	var v2793 int32
	_ = v2793
	var v2794 int64
	_ = v2794
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2804 int32
	_ = v2804
	var v2810 int32
	_ = v2810
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2817 int32
	_ = v2817
	var v2818 int32
	_ = v2818
	var v2820 int32
	_ = v2820
	var v2823 int32
	_ = v2823
	var v2826 int32
	_ = v2826
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2834 int32
	_ = v2834
	var v2835 int64
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2839 int32
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2847 int64
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2857 int32
	_ = v2857
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2880 int32
	_ = v2880
	var v2883 int32
	_ = v2883
	var v2886 int32
	_ = v2886
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2894 int32
	_ = v2894
	var v2895 int64
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2907 int64
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2912 int32
	_ = v2912
	var v2913 int32
	_ = v2913
	var v2917 int32
	_ = v2917
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2955 int32
	_ = v2955
	var v2957 int32
	_ = v2957
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3043 int32
	_ = v3043
	var v3046 int32
	_ = v3046
	var v3052 int32
	_ = v3052
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3062 int32
	_ = v3062
	var v3067 int32
	_ = v3067
	var v3157 int32
	_ = v3157
	var v3160 int32
	_ = v3160
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3181 int32
	_ = v3181
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3186 int32
	_ = v3186
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3199 int32
	_ = v3199
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3227 int32
	_ = v3227
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3245 int32
	_ = v3245
	var v3254 int32
	_ = v3254
	var v3260 int32
	_ = v3260
	var v3265 int32
	_ = v3265
	var v3287 int32
	_ = v3287
	var v3288 int32
	_ = v3288
	var v3291 int32
	_ = v3291
	var v3295 int32
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3299 int32
	_ = v3299
	var v3309 int32
	_ = v3309
	var v3315 int32
	_ = v3315
	var v3353 int32
	_ = v3353
	var v3359 int32
	_ = v3359
	var v3386 int32
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3390 int32
	_ = v3390
	var v3394 int32
	_ = v3394
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3400 int32
	_ = v3400
	var v3404 int32
	_ = v3404
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3410 int32
	_ = v3410
	var v3414 int32
	_ = v3414
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3420 int32
	_ = v3420
	var v3424 int32
	_ = v3424
	var v3426 int32
	_ = v3426
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3439 int32
	_ = v3439
	var v3451 int32
	_ = v3451
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3456 int32
	_ = v3456
	var v3460 int32
	_ = v3460
	var v3472 int32
	_ = v3472
	var v3510 int32
	_ = v3510
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3515 int32
	_ = v3515
	var v3519 int32
	_ = v3519
	var v3528 int32
	_ = v3528
	var v3547 int32
	_ = v3547
	var v3548 int32
	_ = v3548
	var v3558 int32
	_ = v3558
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3595 int32
	_ = v3595
	var v3599 int32
	_ = v3599
	var v3600 int32
	_ = v3600
	var v3602 int32
	_ = v3602
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3618 int32
	_ = v3618
	var v3620 int32
	_ = v3620
	var v3628 int32
	_ = v3628
	var v3630 int32
	_ = v3630
	var v3638 int32
	_ = v3638
	var v3664 int32
	_ = v3664
	var v3667 int32
	_ = v3667
	var v3669 int32
	_ = v3669
	var v3674 int32
	_ = v3674
	var v3675 int64
	_ = v3675
	var v3677 int64
	_ = v3677
	var v3679 int64
	_ = v3679
	var v3681 int64
	_ = v3681
	var v3683 int64
	_ = v3683
	var v3685 int64
	_ = v3685
	var v3687 int64
	_ = v3687
	var v3689 int32
	_ = v3689
	var v3694 int32
	_ = v3694
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3703 int64
	_ = v3703
	var v3705 int64
	_ = v3705
	var v3707 int64
	_ = v3707
	var v3715 int32
	_ = v3715
	var v3716 int64
	_ = v3716
	var v3718 int64
	_ = v3718
	var v3720 int64
	_ = v3720
	var v3722 int64
	_ = v3722
	var v3724 int64
	_ = v3724
	var v3726 int64
	_ = v3726
	var v3728 int64
	_ = v3728
	var v3730 int32
	_ = v3730
	var v3734 int32
	_ = v3734
	var v3738 int32
	_ = v3738
	var v3740 int32
	_ = v3740
	var v3741 int32
	_ = v3741
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3747 int64
	_ = v3747
	var v3749 int64
	_ = v3749
	var v3751 int64
	_ = v3751
	var v3755 int32
	_ = v3755
	var v3761 int32
	_ = v3761
	var v3776 int32
	_ = v3776
	var v3803 int32
	_ = v3803
	var v3861 int32
	_ = v3861
	var v3863 int32
	_ = v3863
	var v3898 int32
	_ = v3898
	var v3899 int32
	_ = v3899
	var v3906 int32
	_ = v3906
	var v3908 int32
	_ = v3908
	var v3943 int32
	_ = v3943
	var v3944 int32
	_ = v3944
	var v3947 int32
	_ = v3947
	var v3948 int32
	_ = v3948
	var v3952 int32
	_ = v3952
	var v3954 int32
	_ = v3954
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3959 int32
	_ = v3959
	var v3960 int32
	_ = v3960
	var v3963 int32
	_ = v3963
	var v3964 int32
	_ = v3964
	var v3967 int32
	_ = v3967
	var v3985 int32
	_ = v3985
	var v4012 int32
	_ = v4012
	var v4015 int32
	_ = v4015
	var v4016 int32
	_ = v4016
	var v4020 int32
	_ = v4020
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4034 int32
	_ = v4034
	var v4067 int32
	_ = v4067
	var v4071 int32
	_ = v4071
	var v4073 int32
	_ = v4073
	var v4075 int32
	_ = v4075
	var v4122 int32
	_ = v4122
	var v4126 int32
	_ = v4126
	var v4131 int32
	_ = v4131
	var v4136 int32
	_ = v4136
	var v4142 int32
	_ = v4142
	var v4145 int32
	_ = v4145
	var v4148 int32
	_ = v4148
	var v4155 int32
	_ = v4155
	var v4157 int32
	_ = v4157
	var v4160 int32
	_ = v4160
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4173 int32
	_ = v4173
	var v4176 int32
	_ = v4176
	var v4177 int32
	_ = v4177
	var v4188 int32
	_ = v4188
	var v4191 int32
	_ = v4191
	var v4199 int32
	_ = v4199
	var v4200 int32
	_ = v4200
	var v4201 int32
	_ = v4201
	var v4204 int32
	_ = v4204
	var v4206 int32
	_ = v4206
	var v4207 int32
	_ = v4207
	var v4210 int32
	_ = v4210
	var v4211 int32
	_ = v4211
	var v4213 int32
	_ = v4213
	var v4214 int32
	_ = v4214
	var v4218 int32
	_ = v4218
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4225 int32
	_ = v4225
	var v4226 int32
	_ = v4226
	var v4228 int32
	_ = v4228
	var v4231 int32
	_ = v4231
	var v4234 int32
	_ = v4234
	var v4237 int32
	_ = v4237
	var v4241 int32
	_ = v4241
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4244 int64
	_ = v4244
	var v4249 int32
	_ = v4249
	var v4250 int64
	_ = v4250
	var v4254 int32
	_ = v4254
	var v4258 int32
	_ = v4258
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4267 int32
	_ = v4267
	var v4273 int32
	_ = v4273
	var v4283 int32
	_ = v4283
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4296 int32
	_ = v4296
	var v4299 int32
	_ = v4299
	var v4300 int32
	_ = v4300
	var v4309 int32
	_ = v4309
	var v4311 int32
	_ = v4311
	var v4316 int32
	_ = v4316
	var v4319 int32
	_ = v4319
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4339 int32
	_ = v4339
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4370 int32
	_ = v4370
	var v4373 int32
	_ = v4373
	var v4374 int32
	_ = v4374
	var v4376 int32
	_ = v4376
	var v4407 int32
	_ = v4407
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4425 int32
	_ = v4425
	var v4431 int32
	_ = v4431
	var v4437 int32
	_ = v4437
	var v4442 int32
	_ = v4442
	var v4471 int32
	_ = v4471
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4494 int32
	_ = v4494
	var v4499 int32
	_ = v4499
	var v4502 int32
	_ = v4502
	var v4504 int32
	_ = v4504
	var v4509 int32
	_ = v4509
	var v4517 int32
	_ = v4517
	var v4527 int32
	_ = v4527
	var v4538 int32
	_ = v4538
	var v4540 int32
	_ = v4540
	var v4543 int32
	_ = v4543
	var v4552 int32
	_ = v4552
	var v4553 int32
	_ = v4553
	var v4560 int32
	_ = v4560
	var v4567 int32
	_ = v4567
	var v4577 int32
	_ = v4577
	var v4587 int32
	_ = v4587
	var v4588 int32
	_ = v4588
	var v4590 int32
	_ = v4590
	var v4593 int32
	_ = v4593
	var v4633 int32
	_ = v4633
	var v4638 int32
	_ = v4638
	var v4652 int32
	_ = v4652
	var v4656 int32
	_ = v4656
	var v4696 int32
	_ = v4696
	var v4700 int32
	_ = v4700
	var v4701 int32
	_ = v4701
	var v4706 int32
	_ = v4706
	var v4707 int32
	_ = v4707
	var v4708 int64
	_ = v4708
	var v4710 int64
	_ = v4710
	var v4712 int64
	_ = v4712
	var v4714 int64
	_ = v4714
	var v4716 int64
	_ = v4716
	var v4718 int64
	_ = v4718
	var v4720 int64
	_ = v4720
	var v4738 int32
	_ = v4738
	var v4749 int32
	_ = v4749
	var v4763 int32
	_ = v4763
	var v4768 int32
	_ = v4768
	var v4770 int32
	_ = v4770
	var v4772 int32
	_ = v4772
	var v4773 int64
	_ = v4773
	var v4775 int64
	_ = v4775
	var v4777 int64
	_ = v4777
	var v4779 int64
	_ = v4779
	var v4781 int64
	_ = v4781
	var v4783 int64
	_ = v4783
	var v4785 int64
	_ = v4785
	var v4788 int32
	_ = v4788
	var v4789 int32
	_ = v4789
	var v4802 int32
	_ = v4802
	var v4803 int32
	_ = v4803
	var v4805 int32
	_ = v4805
	var v4808 int32
	_ = v4808
	var v4810 int32
	_ = v4810
	var v4811 int32
	_ = v4811
	var v4814 int32
	_ = v4814
	var v4815 int32
	_ = v4815
	var v4816 int32
	_ = v4816
	var v4817 int32
	_ = v4817
	var v4818 int32
	_ = v4818
	var v4819 int64
	_ = v4819
	var v4823 int32
	_ = v4823
	var v4830 int32
	_ = v4830
	var v4832 int32
	_ = v4832
	var v4834 int32
	_ = v4834
	var v4836 int32
	_ = v4836
	var v4837 int32
	_ = v4837
	var v4843 int32
	_ = v4843
	var v4844 int32
	_ = v4844
	var v4846 int32
	_ = v4846
	var v4847 int32
	_ = v4847
	var v4848 int64
	_ = v4848
	var v4850 int32
	_ = v4850
	var v4854 int32
	_ = v4854
	var v4883 int32
	_ = v4883
	var v4898 int32
	_ = v4898
	var v4902 int32
	_ = v4902
	var v4904 int32
	_ = v4904
	var v4909 int32
	_ = v4909
	var v4913 int32
	_ = v4913
	var v4917 int32
	_ = v4917
	var v4919 int32
	_ = v4919
	var v4964 int32
	_ = v4964
	var v4967 int32
	_ = v4967
	var v4971 int32
	_ = v4971
	var v4973 int32
	_ = v4973
	var v5018 int32
	_ = v5018
	var v5022 int32
	_ = v5022
	var v5025 int32
	_ = v5025
	var v5029 int32
	_ = v5029
	var v5031 int32
	_ = v5031
	var v5076 int32
	_ = v5076
	var v5079 int32
	_ = v5079
	var v5083 int32
	_ = v5083
	var v5085 int32
	_ = v5085
	var v5130 int32
	_ = v5130
	var v5135 int32
	_ = v5135
	var v5141 int32
	_ = v5141
	var v5146 int32
	_ = v5146
	var v5147 int32
	_ = v5147
	var v5190 int32
	_ = v5190
	var v5193 int32
	_ = v5193
	var v5195 int32
	_ = v5195
	var v5198 int32
	_ = v5198
	var v5199 int32
	_ = v5199
	var v5200 int32
	_ = v5200
	var v5204 int32
	_ = v5204
	var v5207 int32
	_ = v5207
	var v5209 int32
	_ = v5209
	var v5210 int32
	_ = v5210
	var v5213 int32
	_ = v5213
	var v5214 int32
	_ = v5214
	var v5215 int32
	_ = v5215
	var v5218 int32
	_ = v5218
	var v5221 int32
	_ = v5221
	var v5222 int32
	_ = v5222
	var v5223 int32
	_ = v5223
	var v5224 int32
	_ = v5224
	var v5227 int32
	_ = v5227
	var v5230 int32
	_ = v5230
	var v5231 int32
	_ = v5231
	var v5234 int32
	_ = v5234
	var v5235 int32
	_ = v5235
	var v5237 int32
	_ = v5237
	var v5238 int32
	_ = v5238
	var v5241 int32
	_ = v5241
	var v5247 int32
	_ = v5247
	var v5248 int32
	_ = v5248
	var v5252 int32
	_ = v5252
	var v5253 int32
	_ = v5253
	var v5254 int32
	_ = v5254
	var v5255 int32
	_ = v5255
	var v5268 int32
	_ = v5268
	var v5273 int32
	_ = v5273
	var v5315 int32
	_ = v5315
	var v5316 int32
	_ = v5316
	var v5319 int32
	_ = v5319
	var v5320 int32
	_ = v5320
	var v5321 int32
	_ = v5321
	var v5326 int32
	_ = v5326
	var v5329 int32
	_ = v5329
	var v5331 int32
	_ = v5331
	var v5333 int32
	_ = v5333
	var v5334 int32
	_ = v5334
	var v5338 int32
	_ = v5338
	var v5342 int32
	_ = v5342
	var v5348 int32
	_ = v5348
	var v5350 int32
	_ = v5350
	var v5356 int32
	_ = v5356
	var v5361 int32
	_ = v5361
	var v5363 int32
	_ = v5363
	var v5364 int32
	_ = v5364
	var v5367 int32
	_ = v5367
	var v5375 int32
	_ = v5375
	var v5377 int32
	_ = v5377
	var v5380 int32
	_ = v5380
	var v5381 int32
	_ = v5381
	var v5385 int32
	_ = v5385
	var v5388 int32
	_ = v5388
	var v5389 int32
	_ = v5389
	var v5392 int32
	_ = v5392
	var v5393 int32
	_ = v5393
	var v5395 int32
	_ = v5395
	var v5396 int32
	_ = v5396
	var v5399 int32
	_ = v5399
	var v5405 int32
	_ = v5405
	var v5409 int32
	_ = v5409
	var v5414 int32
	_ = v5414
	var v5457 int32
	_ = v5457
	var v5475 int32
	_ = v5475
	v3 = int32(0)
	v42 = m.G0
	v44 = v42 - int32(2064)
	m.G0 = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+68)) = int32(-1)
	v50 = m.G0
	v52 = v50 - int32(240)
	m.G0 = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v3 < v55 {
		v4145 = l1
		v4148 = v52
		v4155 = l0
		v4157 = v44
		v4160 = v3
		v4165 = v46
		v4166 = v47
		v4170 = v3
		v4171 = v3
		v4173 = v3
		v4176 = v3
		v4177 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v4148 + int32(240)
	v4188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4166))))
	if v4188 == int32(0) {
		goto L576
	} else {
		goto L577
	}
L2:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+224))
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v61
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v63)
	if v58 <= v61 {
		v4145 = l1
		v4148 = v52
		v4155 = l0
		v4157 = v44
		v4160 = v3
		v4165 = v46
		v4166 = v47
		v4170 = v3
		v4171 = v3
		v4173 = v3
		v4176 = v3
		v4177 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v68 <= int32(0) {
		v269 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v67)+224))
	v292 = int32(0)
	v295 = int32(1)
	v297 = v295
	v299 = v292
	v300 = v292
	v302 = v3
	v304 = v3
	v306 = v295
	v307 = v292
	goto L16
L5:
	;
	v72 = v68 & int32(3)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v68) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v84 = v3
	v97 = v3
	v98 = v3
	goto L9
L7:
	;
	v159 = v3
	v173 = v3
	goto L8
L8:
	;
	v200 = v159
	v203 = v3
	v214 = v173
	goto L13
L9:
	;
	v121 = v73 + v84*int32(56)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+168))
	v123 = int32(5)
	v125 = int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v121)+112))
	v145 = int32(base.Ui32(v122)>>(uint(v123)%32))&v125 + (int32(base.Ui32(v127)>>(uint(v123)%32))&v125 + v98 + int32(base.Ui32(v133)>>(uint(v123)%32))&v125 + int32(base.Ui32(v139)>>(uint(v123)%32))&v125)
	v146 = int32(4)
	v147 = v84 + v146
	v149 = v97 + v146
	if v149 != v68&int32(2147483644) {
		v84 = v147
		v97 = v149
		v98 = v145
		goto L9
	} else {
		goto L11
	}
L10:
	;
	if v72 == int32(0) {
		v269 = v145
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v159 = v147
	v173 = v145
	goto L8
L13:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v73+v200*int32(56))))
	v241 = int32(1)
	v243 = int32(base.Ui32(v238)>>(uint(int32(5))%32))&v241 + v214
	v247 = v203 + v241
	if v247 != v72 {
		v200 = v200 + v241
		v203 = v247
		v214 = v243
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v269 = v243
	goto L4
L15:
	;
	goto L14
L16:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v339 = base.I32_extend16_s(v306)
	v340 = base.I32_extend16_s(v297)
	if base.B2i32(v340 <= v339) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+16)) = uint8(base.B2i32(int32(0) < v527))
	v566 = v527 + v269
	if v566 != 0 {
		goto L41
	} else {
		goto L42
	}
L18:
	;
	goto L17
L19:
	;
	v351 = v300
	v354 = v339
	goto L22
L20:
	;
	v420 = v300
	v423 = v339
	v426 = v306
	goto L21
L21:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (base.B2i32(v307 == v458)|v304)&int32(1) != 0 {
		v527 = v420
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v392 = v354<<(uint(int32(2))%32) - int32(4)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v67)+208))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v396+v392)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v67)+212))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399+v392)))
	v403 = F_get_opfamily_member(m, v398, v401, v401, int32(3))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v420 = v413
	v423 = v411
	v426 = v297
	goto L21
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392+(v52+int32(80))))) = v403
	if v403 == int32(0) {
		v527 = v302
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v410 = int32(1)
	v411 = v354 + v410
	v413 = v351 + v410
	if (v297+v300-v306)&int32(_a_F__bt_first_0) != v413&int32(_a_F__bt_first_0) {
		v351 = v413
		v354 = v411
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v465 = v338 + v307*int32(56)
	v466 = int32(*(*int16)(unsafe.Add(mBase, uint32(v465)+4)))
	if v466 <= v340 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465)+6)))
	v297 = v502
	v299 = int32(base.Ui32(v506&int32(64))>>(uint(int32(6))%32)) | base.B2i32(v511 == int32(3)) | v503
	v300 = v504
	v302 = v420
	v304 = int32(base.Ui32(v506&int32(4)) >> (uint(int32(2)) % 32))
	v306 = v505
	v307 = v307 + int32(1)
	goto L16
L30:
	;
	v502 = v297
	v503 = v299
	v504 = v420
	v505 = v426
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
	v502 = v497
	v503 = int32(0)
	v504 = v498
	v505 = v426 + int32(1)
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423<<(uint(int32(2))%32)+v52)+76)) = int32(0)
	v497 = v466
	v498 = v420
	goto L33
L35:
	;
	goto L36
L36:
	;
	v478 = v423<<(uint(int32(2))%32) - int32(4)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v67)+208))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v482+v478)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v67)+212))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v485+v478)))
	v489 = F_get_opfamily_member(m, v484, v487, v487, int32(3))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v478+(v52+int32(80))))) = v489
	if v489 == int32(0) {
		v527 = v420
		goto L18
	} else {
		goto L38
	}
L38:
	;
	v494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465)+4)))
	v497 = v494
	v498 = v420 + int32(1)
	goto L33
L39:
	;
	v1788 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1772)+4)))
	if int32(0) < v1788 {
		goto L184
	} else {
		goto L185
	}
L40:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1745 = int32(0)
	v1748 = l1
	v1751 = v52
	v1758 = l0
	v1759 = v1715
	v1760 = v44
	v1763 = v3
	v1767 = v1745
	v1768 = v46
	v1769 = v47
	v1770 = v1745
	v1772 = v1744
	v1773 = v3
	v1774 = v3
	v1776 = v3
	v1779 = v3
	v1780 = v3
	goto L39
L41:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v568 = v567 + v527
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v290)+28))
	if v569 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v1700 == int32(0) {
		v4145 = l1
		v4148 = v52
		v4155 = l0
		v4157 = v44
		v4160 = v3
		v4165 = v46
		v4166 = v47
		v4170 = v3
		v4171 = v3
		v4173 = v3
		v4176 = v3
		v4177 = v3
		goto L1
	} else {
		goto L181
	}
L44:
	;
	v585 = int32(_a_F__bt_first_1)
	v586 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v584
	v591 = F_palloc(m, v568*int32(56))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L24
	} else {
		goto L50
	}
L45:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	v578 = F_AllocSetContextCreateInternal(m, v573, int32(_a_F__bt_first_2), int32(0), int32(1024), int32(_a_F__bt_first_3))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L24
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_MemoryContextReset(m, v569)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L24
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v578
	v584 = v578
	goto L44
L49:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v290)+28))
	v584 = v583
	goto L44
L50:
	;
	v595 = F_palloc(m, v566<<(uint(int32(5))%32))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+20)) = v595
	v600 = F_palloc(m, v568*int32(28))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L24
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+24)) = v600
	v603 = int32(0)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v605 <= v603 {
		v1649 = v603
		v1656 = v603
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+12)) = v1656
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v586
	v1681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v1681 == int32(0) {
		v4145 = l1
		v4148 = v52
		v4155 = l0
		v4157 = v44
		v4160 = v3
		v4165 = v46
		v4166 = v47
		v4170 = v3
		v4171 = v3
		v4173 = v3
		v4176 = v3
		v4177 = v3
		goto L1
	} else {
		goto L176
	}
L54:
	;
	v615 = v527
	v620 = int32(-1)
	v622 = v603
	v624 = v3
	v629 = v603
	v630 = int32(1)
	v633 = v3
	v634 = v3
	goto L55
L55:
	;
	v651 = int32(56)
	v653 = v591 + v622*v651
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v657 = v654 + v624*v651
	*(*int32)(unsafe.Add(mBase, uint32(v52)+48)) = v52 + int32(52)
	if v615 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v1649 = v1604
	v1656 = v1611
	goto L53
L57:
	;
	v1634 = v624 + int32(1)
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1634 < v1635 {
		v615 = v828
		v620 = v1602
		v622 = v1604
		v624 = v1634
		v629 = v1611
		v630 = v843
		v633 = v1615
		v634 = v1616
		goto L55
	} else {
		goto L175
	}
L58:
	;
	v1572 = v842 << (uint(int32(5)) % 32)
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1572+v1573))) = v835
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1576+v1572)+4)) = v1578
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1580+v1572)+8)) = v1582
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1584+v1572)+12)) = int32(-1)
	v1588 = int32(1)
	v1602 = v1566
	v1604 = v835 + v1588
	v1611 = v842 + v1588
	v1615 = v1567
	v1616 = v1568
	goto L57
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L24
	} else {
		goto L172
	}
L60:
	;
	v864 = *(*int64)(unsafe.Add(mBase, uint32(v657)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v832)+48)) = v864
	v866 = *(*int64)(unsafe.Add(mBase, uint32(v657)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v832)+40)) = v866
	v868 = *(*int64)(unsafe.Add(mBase, uint32(v657)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v832)+32)) = v868
	v870 = *(*int64)(unsafe.Add(mBase, uint32(v657)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v832)+24)) = v870
	v872 = *(*int64)(unsafe.Add(mBase, uint32(v657)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v832)+16)) = v872
	v874 = *(*int64)(unsafe.Add(mBase, uint32(v657)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v832)+8)) = v874
	v876 = *(*int64)(unsafe.Add(mBase, uint32(v657)))
	*(*int64)(unsafe.Add(mBase, uint32(v832))) = v876
	v878 = base.I32_wrap_i64(v876)
	if v878&int32(32) != 0 {
		goto L89
	} else {
		goto L90
	}
L61:
	;
	v828 = int32(0)
	v832 = v653
	v835 = v622
	v842 = v629
	v843 = v630
	goto L60
L62:
	;
	goto L63
L63:
	;
	v671 = v615
	v673 = v653
	v676 = v622
	v683 = v629
	v684 = v630
	goto L64
L64:
	;
	v705 = base.I32_extend16_s(v684)
	v706 = int32(*(*int16)(unsafe.Add(mBase, uint32(v657)+4)))
	if v706 < v705 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v828 = v733
	v832 = v820
	v835 = v817
	v842 = v815
	v843 = v813
	goto L60
L66:
	;
	v828 = v671
	v832 = v673
	v835 = v676
	v842 = v683
	v843 = v684
	goto L60
L67:
	;
	goto L68
L68:
	;
	v709 = v705 - int32(1)
	v711 = v709 << (uint(int32(2)) % 32)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v711+(v52+int32(80)))))
	if v715 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v828 = v671
	v832 = v673
	v835 = v676
	v842 = v683
	v843 = v684 + int32(1)
	goto L60
L70:
	;
	goto L71
L71:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v67)+248))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v720+v711)))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v67)+212))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v723+v711)))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v67)+208))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v726+v711)))
	v729 = F_get_opcode(m, v715)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	if v729 == int32(0) {
		goto L59
	} else {
		goto L73
	}
L73:
	;
	v733 = int32(0)
	F_ScanKeyEntryInitialize(m, v673, int32(_a_F__bt_first_4), v705, int32(3), v733, v722, v729, int64(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L24
	} else {
		goto L74
	}
L74:
	;
	v741 = v683 << (uint(int32(5)) % 32)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v741+v742))) = v676
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v745+v741)+4)) = int32(-1)
	v749 = int32(1)
	v752 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291+v709<<(uint(v749)%32)))))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
	v758 = v755 + v709<<(uint(int32(3))%32)
	v759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v758)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v753+v741)+16)) = uint16(v759)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758)+32)))
	*(*uint8)(unsafe.Add(mBase, uint32(v761+v741)+18)) = uint8(v763)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v765+v741)+19)) = uint8(v749)
	v770 = F_get_opfamily_proc(m, v728, v725, v725, int32(6))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L24
	} else {
		goto L75
	}
L75:
	;
	if v770 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v776 = F_palloc(m, int32(24))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L24
	} else {
		goto L79
	}
L77:
	;
	v793 = int32(0)
	goto L78
L78:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v794+v741)+20)) = v793
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v799 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v797+v741)+24)) = v799
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v801+v741)+28)) = v799
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v290)+24))
	F__bt_setup_array_cmp(m, l0, v673, v725, v805+v676*int32(28), v799)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L24
	} else {
		goto L84
	}
L79:
	;
	v779 = F_OidFunctionCall1Coll(m, v770, int32(0), base.I64_extend_i32_u(v776))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L24
	} else {
		goto L80
	}
L80:
	;
	if v752&int32(1) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v781 = *(*int64)(unsafe.Add(mBase, uint32(v776)))
	v782 = *(*int64)(unsafe.Add(mBase, uint32(v776)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v776))) = v782
	*(*int64)(unsafe.Add(mBase, uint32(v776)+8)) = v781
	v785 = *(*int64)(unsafe.Add(mBase, uint32(v776)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v776)+16)) = base.I64_rotl(v785, int64(32))
	goto L83
L82:
	;
	goto L83
L83:
	;
	v793 = v776
	goto L78
L84:
	;
	v812 = int32(1)
	v813 = v684 + v812
	v815 = v683 + v812
	v817 = v676 + v812
	v820 = v591 + v817*int32(56)
	v822 = v671 - v812
	if v822 != 0 {
		v671 = v822
		v673 = v820
		v676 = v817
		v683 = v815
		v684 = v813
		goto L64
	} else {
		goto L85
	}
L85:
	;
	goto L65
L86:
	;
	v1545 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v1545)
	v1649 = v835
	v1656 = v842
	goto L53
L87:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v290)+24))
	F__bt_setup_array_cmp(m, l0, v832, v1108, v1179+v835*int32(28), v52+int32(48))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L24
	} else {
		goto L123
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L24
	} else {
		goto L120
	}
L89:
	;
	if v878&int32(1) != 0 {
		goto L86
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v1602 = v620
	v1604 = v835 + int32(1)
	v1611 = v842
	v1615 = v633
	v1616 = v634
	goto L57
L92:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v832)+48))
	v884 = F_pg_detoast_datum(m, v883)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L24
	} else {
		goto L93
	}
L93:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v884)+12))
	F_get_typlenbyvalalign(m, v886, v52+int32(46), v52+int32(45), v52+int32(44))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v896 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+46)))
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+45)))
	v898 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52)+44)))
	F_deconstruct_array(m, v884, v896, v897, v898, v52+int32(36), v52+int32(32), v52+int32(40))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
	if v907 <= int32(0) {
		goto L86
	} else {
		goto L96
	}
L96:
	;
	v910 = int32(0)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	if v907 != int32(1) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v1057 == int32(0) {
		goto L86
	} else {
		goto L112
	}
L98:
	;
	v921 = int32(0)
	v924 = v910
	v927 = v910
	goto L101
L99:
	;
	v1004 = v910
	v1007 = v910
	goto L100
L100:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007+v912))))
	if v1043 != 0 {
		v1057 = v1004
		goto L97
	} else {
		goto L111
	}
L101:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927+v912))))
	if v963 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v907&int32(1) == int32(0) {
		v1057 = v993
		goto L97
	} else {
		goto L110
	}
L103:
	;
	v966 = int32(3)
	v972 = *(*int64)(unsafe.Add(mBase, uint32(v911+v927<<(uint(v966)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v911+v924<<(uint(v966)%32)))) = v972
	v976 = v924 + int32(1)
	goto L105
L104:
	;
	v976 = v924
	goto L105
L105:
	;
	v978 = v927 | int32(1)
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912+v978))))
	if v980 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v983 = int32(3)
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v911+v978<<(uint(v983)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v911+v976<<(uint(v983)%32)))) = v989
	v993 = v976 + int32(1)
	goto L108
L107:
	;
	v993 = v976
	goto L108
L108:
	;
	v994 = int32(2)
	v995 = v927 + v994
	v997 = v921 + v994
	if v997 != v907&int32(2147483646) {
		v921 = v997
		v924 = v993
		v927 = v995
		goto L101
	} else {
		goto L109
	}
L109:
	;
	goto L102
L110:
	;
	v1004 = v993
	v1007 = v995
	goto L100
L111:
	;
	v1044 = int32(3)
	v1050 = *(*int64)(unsafe.Add(mBase, uint32(v911+v1007<<(uint(v1044)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v911+v1004<<(uint(v1044)%32)))) = v1050
	v1057 = v1004 + int32(1)
	goto L97
L112:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v832)+8))
	if v1097 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v67)+212))
	v1101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v832)+4)))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1100+v1101<<(uint(int32(2))%32)-int32(4))))
	v1108 = v1107
	goto L115
L114:
	;
	v1108 = v1097
	goto L115
L115:
	;
	v1109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+6)))
	switch v1109 - int32(1) {
	case 0, 1:
		goto L117
	case 2:
		goto L87
	case 3, 4:
		goto L116
	default:
		goto L88
	}
L116:
	;
	v1119 = F__bt_find_extreme_element(m, l0, v832, v1108, int32(1), v911, v1057)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L24
	} else {
		goto L119
	}
L117:
	;
	v1113 = F__bt_find_extreme_element(m, l0, v832, v1108, int32(5), v911, v1057)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L24
	} else {
		goto L118
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v832)+48)) = v1113
	v1602 = v620
	v1604 = v835 + int32(1)
	v1611 = v842
	v1615 = v633
	v1616 = v634
	goto L57
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v832)+48)) = v1119
	goto L91
L120:
	;
	v1169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v1169
	F_errmsg_internal(m, int32(_a_F__bt_first_5), v52)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L24
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(2096), int32(_a_F__bt_first_7))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L24
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	v1187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v832)+4)))
	v1188 = int32(1)
	v1191 = int32(2)
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291+v1187<<(uint(v1188)%32)-v1191))))
	v1195 = v1193 & v1188
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	if v1191 <= v1057 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+212)) = v1196
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v832)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+220)) = uint8(v1195)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+216)) = v1201
	F_qsort_arg(m, v1199, v1057, int32(8), int32(215), v52+int32(212))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L24
	} else {
		goto L127
	}
L125:
	;
	v1294 = v1057
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+40)) = v1294
	v1333 = int32(*(*int16)(unsafe.Add(mBase, uint32(v832)+4)))
	if v1333 != v633 {
		goto L141
	} else {
		goto L142
	}
L127:
	;
	v1214 = int32(0)
	v1218 = int32(1)
	goto L128
L128:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v52)+212))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v52)+216))
	v1257 = int32(3)
	v1259 = v1199 + v1218<<(uint(v1257)%32)
	v1260 = *(*int64)(unsafe.Add(mBase, uint32(v1259)))
	v1264 = *(*int64)(unsafe.Add(mBase, uint32(v1199+v1214<<(uint(v1257)%32))))
	v1265 = F_FunctionCall2Coll(m, v1255, v1256, v1260, v1264)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L24
	} else {
		goto L131
	}
L129:
	;
	v1294 = v1285 + int32(1)
	goto L126
L130:
	;
	v1287 = v1218 + int32(1)
	if v1287 != v1057 {
		v1214 = v1285
		v1218 = v1287
		goto L128
	} else {
		goto L140
	}
L131:
	;
	v1267 = base.I32_wrap_i64(v1265)
	if v1267 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v1271 = int32(1)
	goto L134
L133:
	;
	v1271 = int32(0) - v1267
	goto L134
L134:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+220)))
	if v1272 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v1273 = v1271
	goto L137
L136:
	;
	v1273 = v1267
	goto L137
L137:
	;
	if v1273 == int32(0) {
		v1285 = v1214
		goto L130
	} else {
		goto L138
	}
L138:
	;
	v1277 = v1214 + int32(1)
	if v1277 == v1218 {
		v1285 = v1218
		goto L130
	} else {
		goto L139
	}
L139:
	;
	v1282 = *(*int64)(unsafe.Add(mBase, uint32(v1259)))
	*(*int64)(unsafe.Add(mBase, uint32(v1199+v1277<<(uint(int32(3))%32)))) = v1282
	v1285 = v1277
	goto L130
L140:
	;
	goto L129
L141:
	;
	v1566 = v842
	v1567 = v1333
	v1568 = v1108
	goto L58
L142:
	;
	goto L143
L143:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v1338 = v1335 + v620<<(uint(int32(5))%32)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+8))
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
	if v1108 != v634 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+208))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1345+v633<<(uint(int32(2))%32)-int32(4))))
	v1353 = F_get_opfamily_proc(m, v1351, v634, v1108, int32(1))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L24
	} else {
		goto L147
	}
L145:
	;
	v1365 = v1196
	goto L146
L146:
	;
	v1366 = int32(0)
	if base.B2i32(v1294 <= v1366)|base.B2i32(v1339 <= v1366) != 0 {
		v1465 = v1366
		goto L150
	} else {
		goto L151
	}
L147:
	;
	if v1353 == int32(0) {
		v1566 = v620
		v1567 = v633
		v1568 = v634
		goto L58
	} else {
		goto L148
	}
L148:
	;
	v1358 = v52 + int32(212)
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+28))
	F_fmgr_info_cxt(m, v1353, v1358, v1359)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L24
	} else {
		goto L149
	}
L149:
	;
	v1365 = v1358
	goto L146
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+4)) = v1465
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
	F_pfree(m, v1500)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L24
	} else {
		goto L170
	}
L151:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v832)+12))
	v1373 = int32(0)
	v1381 = v1373
	v1382 = v1366
	v1384 = v1373
	goto L152
L152:
	;
	v1418 = int32(3)
	v1420 = v1340 + v1384<<(uint(v1418)%32)
	v1421 = *(*int64)(unsafe.Add(mBase, uint32(v1420)))
	v1425 = *(*int64)(unsafe.Add(mBase, uint32(v1341+v1381<<(uint(v1418)%32))))
	v1426 = F_FunctionCall2Coll(m, v1365, v1372, v1421, v1425)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L24
	} else {
		goto L155
	}
L153:
	;
	v1465 = v1454
	goto L150
L154:
	;
	if v1339 <= v1455 {
		v1465 = v1454
		goto L150
	} else {
		goto L168
	}
L155:
	;
	v1428 = base.I32_wrap_i64(v1426)
	if v1428 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v1432 = int32(1)
	goto L158
L157:
	;
	v1432 = int32(0) - v1428
	goto L158
L158:
	;
	if v1195 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1433 = v1432
	goto L161
L160:
	;
	v1433 = v1428
	goto L161
L161:
	;
	if v1433 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1439 = *(*int64)(unsafe.Add(mBase, uint32(v1420)))
	*(*int64)(unsafe.Add(mBase, uint32(v1340+v1382<<(uint(int32(3))%32)))) = v1439
	v1441 = int32(1)
	v1453 = v1381 + v1441
	v1454 = v1382 + v1441
	v1455 = v1384 + v1441
	goto L154
L163:
	;
	goto L164
L164:
	;
	if v1433 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1453 = v1381
	v1454 = v1382
	v1455 = v1384 + int32(1)
	goto L154
L166:
	;
	goto L167
L167:
	;
	v1453 = v1381 + int32(1)
	v1454 = v1382
	v1455 = v1384
	goto L154
L168:
	;
	if v1453 < v1294 {
		v1381 = v1453
		v1382 = v1454
		v1384 = v1455
		goto L152
	} else {
		goto L169
	}
L169:
	;
	goto L153
L170:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+4))
	if v1503 != 0 {
		v1602 = v620
		v1604 = v835
		v1611 = v842
		v1615 = v633
		v1616 = v634
		goto L57
	} else {
		goto L171
	}
L171:
	;
	goto L86
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v715
	F_errmsg_internal(m, int32(_a_F__bt_first_8), v52+int32(16))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L24
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(1959), int32(_a_F__bt_first_7))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L24
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
	goto L56
L176:
	;
	if v591 == int32(0) {
		v1715 = v1649
		goto L40
	} else {
		goto L177
	}
L177:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	v1689 = F_MemoryContextAlloc(m, v1686, v1649<<(uint(int32(2))%32))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L24
	} else {
		goto L178
	}
L178:
	;
	v1691 = int32(1)
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1649 <= v1692 {
		v1748 = l1
		v1751 = v52
		v1758 = l0
		v1759 = v1649
		v1760 = v44
		v1763 = v3
		v1767 = v1691
		v1768 = v46
		v1769 = v47
		v1770 = v1689
		v1772 = v591
		v1773 = v3
		v1774 = v3
		v1776 = v3
		v1779 = v3
		v1780 = v3
		goto L39
	} else {
		goto L179
	}
L179:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v1697 = F_repalloc(m, v1694, v1649*int32(56))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L24
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v1697
	v1748 = l1
	v1751 = v52
	v1758 = l0
	v1759 = v1649
	v1760 = v44
	v1763 = v3
	v1767 = v1691
	v1768 = v46
	v1769 = v47
	v1770 = v1689
	v1772 = v591
	v1773 = v3
	v1774 = v3
	v1776 = v3
	v1779 = v3
	v1780 = v3
	goto L39
L181:
	;
	v1715 = v58
	goto L40
L182:
	;
	v4142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v4142)
	v4145 = v1748
	v4148 = v1751
	v4155 = v1758
	v4157 = v1760
	v4160 = v1763
	v4165 = v1768
	v4166 = v1769
	v4170 = v1773
	v4171 = v1774
	v4173 = v1776
	v4176 = v1779
	v4177 = v1780
	goto L1
L183:
	;
	v4136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v4136)
	v4145 = v1748
	v4148 = v1751
	v4155 = v1758
	v4157 = v1760
	v4160 = v1763
	v4165 = v1768
	v4166 = v1769
	v4170 = v1773
	v4171 = v1774
	v4173 = v1776
	v4176 = v1779
	v4177 = v1780
	goto L1
L184:
	;
	if v1759 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4122 = m.ExcPending
	if v4122 != 0 {
		goto L24
	} else {
		goto L572
	}
L187:
	;
	v1796 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1772)+4)))
	v1797 = int32(1)
	v1802 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60+v1796<<(uint(v1797)%32)-int32(2)))))
	v1804 = v1802 << (uint(int32(24)) % 32)
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1772)))
	if v1805&v1797 != 0 {
		goto L194
	} else {
		goto L195
	}
L188:
	;
	goto L189
L189:
	;
	v1924 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1751)+136)) = v1924
	v1926 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+128)) = v1926
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+120)) = v1926
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+112)) = v1926
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+104)) = v1926
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+96)) = v1926
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+88)) = v1926
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+80)) = v1926
	v1946 = v1924
	v1948 = v1924
	v1952 = v1924
	v1953 = v1924
	v1955 = v1924
	v1965 = int32(1)
	goto L220
L190:
	;
	if v1896 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L191:
	;
	v1896 = int32(1)
	goto L190
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1772)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1772)+6)) = uint16(v1881)
	goto L191
L193:
	;
	if v1808&int32(33554432) != 0 {
		goto L212
	} else {
		goto L213
	}
L194:
	;
	v1808 = v1805 | v1804
	*(*int32)(unsafe.Add(mBase, uint32(v1772))) = v1808
	if v1805&int32(64) != 0 {
		v1881 = int32(3)
		goto L192
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v1818 = int32(0)
	if base.B2i32(v1802&int32(1) == v1818)|v1805&int32(16777216) == v1818 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	if v1805&int32(128) != 0 {
		goto L193
	} else {
		goto L198
	}
L198:
	;
	v1896 = int32(0)
	goto L190
L199:
	;
	v1826 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1772)+6)))
	v1827 = int32(6) - v1826
	*(*uint16)(unsafe.Add(mBase, uint32(v1772)+6)) = uint16(v1827)
	goto L201
L200:
	;
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1772))) = v1805 | v1804
	if v1805&int32(4) == int32(0) {
		goto L191
	} else {
		goto L202
	}
L202:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1772)+48))
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1835))))
	if v1836&int32(1) != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1896 = int32(0)
	goto L190
L204:
	;
	goto L205
L205:
	;
	v1840 = v1835
	goto L206
L206:
	;
	v1845 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1840)+4)))
	v1846 = int32(1)
	v1851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60+v1845<<(uint(v1846)%32)-int32(2)))))
	v1856 = int32(0)
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1840)))
	if base.B2i32(v1851&v1846 == v1856)|v1858&int32(16777216) == v1856 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	goto L191
L208:
	;
	v1865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1840)+6)))
	v1866 = int32(6) - v1865
	*(*uint16)(unsafe.Add(mBase, uint32(v1840)+6)) = uint16(v1866)
	goto L210
L209:
	;
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1840))) = v1858 | v1851<<(uint(int32(24))%32)
	if v1858&int32(16) == int32(0) {
		v1840 = v1840 + int32(56)
		goto L206
	} else {
		goto L211
	}
L211:
	;
	goto L207
L212:
	;
	v1880 = int32(5)
	goto L214
L213:
	;
	v1880 = int32(1)
	goto L214
L214:
	;
	v1881 = v1880
	goto L192
L215:
	;
	v1899 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v1899)
	goto L217
L216:
	;
	goto L217
L217:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v1902 = *(*int64)(unsafe.Add(mBase, uint32(v1772)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1901)+48)) = v1902
	v1904 = *(*int64)(unsafe.Add(mBase, uint32(v1772)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1901)+40)) = v1904
	v1906 = *(*int64)(unsafe.Add(mBase, uint32(v1772)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1901)+32)) = v1906
	v1908 = *(*int64)(unsafe.Add(mBase, uint32(v1772)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1901)+24)) = v1908
	v1910 = *(*int64)(unsafe.Add(mBase, uint32(v1772)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1901)+16)) = v1910
	v1912 = *(*int64)(unsafe.Add(mBase, uint32(v1772)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1901)+8)) = v1912
	v1914 = *(*int64)(unsafe.Add(mBase, uint32(v1772)))
	*(*int64)(unsafe.Add(mBase, uint32(v1901))) = v1914
	v1916 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v1916
	v1918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1772)+4)))
	if v1918 != v1916 {
		v4145 = v1748
		v4148 = v1751
		v4155 = v1758
		v4157 = v1760
		v4160 = v1763
		v4165 = v1768
		v4166 = v1769
		v4170 = v1773
		v4171 = v1774
		v4173 = v1776
		v4176 = v1779
		v4177 = v1780
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	F__bt_mark_scankey_required(m, v1921)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L24
	} else {
		goto L219
	}
L219:
	;
	v4145 = v1748
	v4148 = v1751
	v4155 = v1758
	v4157 = v1760
	v4160 = v1763
	v4165 = v1768
	v4166 = v1769
	v4170 = v1773
	v4171 = v1774
	v4173 = v1776
	v4176 = v1779
	v4177 = v1780
	goto L1
L220:
	;
	v1989 = v1772 + v1955*int32(56)
	v1990 = base.B2i32(v1759 <= v1955)
	if v1759 <= v1955 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v2097 = base.B2i32(v1955 == v1759)
	if v2097 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L223:
	;
	v1994 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1989)+4)))
	v1995 = int32(1)
	v2000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60+v1994<<(uint(v1995)%32)-int32(2)))))
	v2002 = v2000 << (uint(int32(24)) % 32)
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1989)))
	if v2003&v1995 != 0 {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	if v2094 != 0 {
		goto L222
	} else {
		goto L249
	}
L225:
	;
	v2094 = int32(1)
	goto L224
L226:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1989)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1989)+6)) = uint16(v2079)
	goto L225
L227:
	;
	if v2006&int32(33554432) != 0 {
		goto L246
	} else {
		goto L247
	}
L228:
	;
	v2006 = v2003 | v2002
	*(*int32)(unsafe.Add(mBase, uint32(v1989))) = v2006
	if v2003&int32(64) != 0 {
		v2079 = int32(3)
		goto L226
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v2016 = int32(0)
	if base.B2i32(v2000&int32(1) == v2016)|v2003&int32(16777216) == v2016 {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	if v2003&int32(128) != 0 {
		goto L227
	} else {
		goto L232
	}
L232:
	;
	v2094 = int32(0)
	goto L224
L233:
	;
	v2024 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1989)+6)))
	v2025 = int32(6) - v2024
	*(*uint16)(unsafe.Add(mBase, uint32(v1989)+6)) = uint16(v2025)
	goto L235
L234:
	;
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1989))) = v2003 | v2002
	if v2003&int32(4) == int32(0) {
		goto L225
	} else {
		goto L236
	}
L236:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+48))
	v2034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2033))))
	if v2034&int32(1) != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v2094 = int32(0)
	goto L224
L238:
	;
	goto L239
L239:
	;
	v2038 = v2033
	goto L240
L240:
	;
	v2043 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2038)+4)))
	v2044 = int32(1)
	v2049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60+v2043<<(uint(v2044)%32)-int32(2)))))
	v2054 = int32(0)
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2038)))
	if base.B2i32(v2049&v2044 == v2054)|v2056&int32(16777216) == v2054 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	goto L225
L242:
	;
	v2063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2038)+6)))
	v2064 = int32(6) - v2063
	*(*uint16)(unsafe.Add(mBase, uint32(v2038)+6)) = uint16(v2064)
	goto L244
L243:
	;
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2038))) = v2056 | v2049<<(uint(int32(24))%32)
	if v2056&int32(16) == int32(0) {
		v2038 = v2038 + int32(56)
		goto L240
	} else {
		goto L245
	}
L245:
	;
	goto L241
L246:
	;
	v2078 = int32(5)
	goto L248
L247:
	;
	v2078 = int32(1)
	goto L248
L248:
	;
	v2079 = v2078
	goto L226
L249:
	;
	v2095 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v2095)
	v4145 = v1748
	v4148 = v1751
	v4155 = v1758
	v4157 = v1760
	v4160 = v1763
	v4165 = v1768
	v4166 = v1769
	v4170 = v1773
	v4171 = v1774
	v4173 = v1776
	v4176 = v1779
	v4177 = v1780
	goto L1
L250:
	;
	v1946 = v2501
	v1948 = v2502
	v1952 = v2505
	v1953 = v2522
	v1955 = v1955 + int32(1)
	v1965 = v2509
	goto L220
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v2483
	if v1767 != 0 {
		goto L387
	} else {
		goto L388
	}
L252:
	;
	v2511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1989)+6)))
	v2513 = v2511 - int32(1)
	if v2511 == int32(3) {
		goto L362
	} else {
		goto L363
	}
L253:
	;
	v2100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1989)+4)))
	if v2100 == v1965&int32(_a_F__bt_first_0) {
		v2501 = v1946
		v2502 = v1948
		v2505 = v1952
		v2509 = v1965
		goto L252
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	if v1990 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L256:
	;
	goto L255
L257:
	;
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+104))
	if v2382 == int32(0) {
		v2415 = v2380
		goto L340
	} else {
		goto L341
	}
L258:
	;
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v2349 = v2346 + v2344*int32(56)
	v2350 = *(*int64)(unsafe.Add(mBase, uint32(v2341)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2349)+48)) = v2350
	v2352 = *(*int64)(unsafe.Add(mBase, uint32(v2341)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2349)+40)) = v2352
	v2354 = *(*int64)(unsafe.Add(mBase, uint32(v2341)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2349)+32)) = v2354
	v2356 = *(*int64)(unsafe.Add(mBase, uint32(v2341)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2349)+24)) = v2356
	v2358 = *(*int64)(unsafe.Add(mBase, uint32(v2341)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2349)+16)) = v2358
	v2360 = *(*int64)(unsafe.Add(mBase, uint32(v2341)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2349)+8)) = v2360
	v2362 = *(*int64)(unsafe.Add(mBase, uint32(v2341)))
	*(*int64)(unsafe.Add(mBase, uint32(v2349))) = v2362
	if v1767 != 0 {
		goto L333
	} else {
		goto L334
	}
L259:
	;
	if v2333 == int32(0) {
		v2379 = v2335
		v2380 = v2336
		v2381 = v2338
		goto L257
	} else {
		goto L332
	}
L260:
	;
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v2305 = v2302 + v1952*int32(56)
	v2306 = *(*int64)(unsafe.Add(mBase, uint32(v2246)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2305)+48)) = v2306
	v2308 = *(*int64)(unsafe.Add(mBase, uint32(v2246)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2305)+40)) = v2308
	v2310 = *(*int64)(unsafe.Add(mBase, uint32(v2246)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2305)+32)) = v2310
	v2312 = *(*int64)(unsafe.Add(mBase, uint32(v2246)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2305)+24)) = v2312
	v2314 = *(*int64)(unsafe.Add(mBase, uint32(v2246)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2305)+16)) = v2314
	v2316 = *(*int64)(unsafe.Add(mBase, uint32(v2246)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2305)+8)) = v2316
	v2318 = *(*int64)(unsafe.Add(mBase, uint32(v2246)))
	*(*int64)(unsafe.Add(mBase, uint32(v2305))) = v2318
	if v1767 != 0 {
		goto L327
	} else {
		goto L328
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1751)+128)) = int32(0)
	v2341 = v2249
	v2343 = v2245
	v2344 = v1952
	v2345 = base.B2i32(v1948 == base.I32_extend16_s(v1965)-int32(1))
	goto L258
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L24
	} else {
		goto L324
	}
L263:
	;
	v2106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1989)+4)))
	if v2106 < base.I32_extend16_s(v1965) {
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+104))
	if v2109 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L265
L267:
	;
	v2223 = int32(0)
	if base.B2i32(v2219 == v2223)|base.B2i32(v2217 == v2223) != 0 {
		goto L307
	} else {
		goto L308
	}
L268:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+92))
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+80))
	v2216 = v1946
	v2217 = v2112
	v2219 = v2113
	v2222 = v1948
	goto L267
L269:
	;
	goto L270
L270:
	;
	v2114 = int32(0)
	if v1767 == v2114 {
		v2136 = v2114
		v2137 = v2114
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+128))
	if v2138 != 0 {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	v2118 = int32(0)
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109))))
	if v2119&int32(32) == v2118 {
		v2136 = v2114
		v2137 = v2118
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+112))
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+108))
	v2136 = v2124 + v2125<<(uint(int32(5))%32) - int32(32)
	v2137 = v2131 + v2132*int32(28)
	goto L271
L274:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+116))
	if v2156 != 0 {
		goto L283
	} else {
		goto L284
	}
L275:
	;
	v2139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109))))
	if v2139&int32(64) != 0 {
		goto L182
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v2155 = v1946
	goto L274
L278:
	;
	v2145 = F__bt_compare_scankey_args(m, v1758, v2138, v2109, v2138, v2136, v2137, v1751+int32(212))
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L24
	} else {
		goto L279
	}
L279:
	;
	if v2145 == int32(0) {
		v2155 = int32(1)
		goto L274
	} else {
		goto L280
	}
L280:
	;
	v2149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+212)))
	if v2149 == int32(0) {
		goto L183
	} else {
		goto L281
	}
L281:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+128)) = int64(-4294967296)
	goto L277
L282:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+92))
	if v2174 != 0 {
		goto L291
	} else {
		goto L292
	}
L283:
	;
	v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109))))
	if v2157&int32(64) != 0 {
		goto L182
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v2173 = v2155
	goto L282
L286:
	;
	v2163 = F__bt_compare_scankey_args(m, v1758, v2156, v2109, v2156, v2136, v2137, v1751+int32(212))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L24
	} else {
		goto L287
	}
L287:
	;
	if v2163 == int32(0) {
		v2173 = int32(1)
		goto L282
	} else {
		goto L288
	}
L288:
	;
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+212)))
	if v2167 != int32(1) {
		goto L183
	} else {
		goto L289
	}
L289:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+116)) = int64(-4294967296)
	goto L285
L290:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+80))
	if v2194 != 0 {
		goto L299
	} else {
		goto L300
	}
L291:
	;
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109))))
	if v2175&int32(64) != 0 {
		goto L182
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v2192 = int32(0)
	v2193 = v2173
	goto L290
L294:
	;
	v2181 = F__bt_compare_scankey_args(m, v1758, v2174, v2109, v2174, v2136, v2137, v1751+int32(212))
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L24
	} else {
		goto L295
	}
L295:
	;
	if v2181 == int32(0) {
		v2192 = v2174
		v2193 = int32(1)
		goto L290
	} else {
		goto L296
	}
L296:
	;
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+212)))
	if v2185 != int32(1) {
		goto L183
	} else {
		goto L297
	}
L297:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+92)) = int64(-4294967296)
	goto L293
L298:
	;
	v2216 = v2212
	v2217 = v2192
	v2219 = v2213
	v2222 = v1948 + int32(1)
	goto L267
L299:
	;
	v2195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109))))
	if v2195&int32(64) != 0 {
		goto L182
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v2212 = v2193
	v2213 = int32(0)
	goto L298
L302:
	;
	v2201 = F__bt_compare_scankey_args(m, v1758, v2194, v2109, v2194, v2136, v2137, v1751+int32(212))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L24
	} else {
		goto L303
	}
L303:
	;
	if v2201 == int32(0) {
		v2212 = int32(1)
		v2213 = v2194
		goto L298
	} else {
		goto L304
	}
L304:
	;
	v2205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+212)))
	if v2205 != int32(1) {
		goto L183
	} else {
		goto L305
	}
L305:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+80)) = int64(-4294967296)
	goto L301
L306:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+128))
	v2247 = int32(0)
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+116))
	if base.B2i32(v2246 == v2247)|base.B2i32(v2249 == v2247) == v2247 {
		goto L314
	} else {
		goto L315
	}
L307:
	;
	v2245 = v2216
	goto L306
L308:
	;
	v2229 = int32(0)
	v2233 = F__bt_compare_scankey_args(m, v1758, v2217, v2219, v2217, v2229, v2229, v1751+int32(212))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L24
	} else {
		goto L309
	}
L309:
	;
	if v2233 == int32(0) {
		v2245 = int32(1)
		goto L306
	} else {
		goto L310
	}
L310:
	;
	v2237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+212)))
	if v2237 == int32(1) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1751)+92)) = int32(0)
	goto L307
L312:
	;
	goto L313
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1751)+80)) = int32(0)
	goto L307
L314:
	;
	v2255 = int32(0)
	v2259 = F__bt_compare_scankey_args(m, v1758, v2249, v2246, v2249, v2255, v2255, v1751+int32(212))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L24
	} else {
		goto L318
	}
L315:
	;
	goto L316
L316:
	;
	v2279 = base.B2i32(v1948 == base.I32_extend16_s(v1965)-int32(1))
	if v2246 != 0 {
		v2299 = v2249
		v2300 = v2279
		v2301 = v2245
		goto L260
	} else {
		goto L323
	}
L317:
	;
	v2299 = v2270
	v2300 = base.B2i32(v1948 == base.I32_extend16_s(v1965)-int32(1))
	v2301 = v2271
	goto L260
L318:
	;
	if v2259 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v2270 = v2249
	v2271 = int32(1)
	goto L317
L320:
	;
	goto L321
L321:
	;
	v2264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+212)))
	if v2264 != int32(1) {
		goto L261
	} else {
		goto L322
	}
L322:
	;
	v2267 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1751)+116)) = v2267
	v2270 = v2267
	v2271 = v2245
	goto L317
L323:
	;
	v2333 = v2249
	v2335 = v2245
	v2336 = v1952
	v2338 = v2279
	goto L259
L324:
	;
	F_errmsg_internal(m, int32(_a_F__bt_first_9), int32(0))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L24
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(345), int32(_a_F__bt_first_10))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L24
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v1952<<(uint(int32(2))%32)))) = v2323
	goto L329
L328:
	;
	goto L329
L329:
	;
	v2326 = v1952 + int32(1)
	v2327 = int32(0)
	if v2300 == v2327 {
		v2333 = v2299
		v2335 = v2301
		v2336 = v2326
		v2338 = v2327
		goto L259
	} else {
		goto L330
	}
L330:
	;
	F__bt_mark_scankey_required(m, v2305)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L24
	} else {
		goto L331
	}
L331:
	;
	v2333 = v2299
	v2335 = v2301
	v2336 = v2326
	v2338 = int32(1)
	goto L259
L332:
	;
	v2341 = v2333
	v2343 = v2335
	v2344 = v2336
	v2345 = v2338
	goto L258
L333:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v2344<<(uint(int32(2))%32)))) = v2367
	goto L335
L334:
	;
	goto L335
L335:
	;
	v2370 = v2344 + int32(1)
	if v2345 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2379 = v2343
	v2380 = v2370
	v2381 = int32(0)
	goto L257
L337:
	;
	goto L338
L338:
	;
	F__bt_mark_scankey_required(m, v2349)
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L24
	} else {
		goto L339
	}
L339:
	;
	v2379 = v2343
	v2380 = v2370
	v2381 = int32(1)
	goto L257
L340:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+92))
	if v2416 == int32(0) {
		v2449 = v2415
		goto L347
	} else {
		goto L348
	}
L341:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v2388 = v2385 + v2380*int32(56)
	v2389 = *(*int64)(unsafe.Add(mBase, uint32(v2382)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2388)+48)) = v2389
	v2391 = *(*int64)(unsafe.Add(mBase, uint32(v2382)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2388)+40)) = v2391
	v2393 = *(*int64)(unsafe.Add(mBase, uint32(v2382)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2388)+32)) = v2393
	v2395 = *(*int64)(unsafe.Add(mBase, uint32(v2382)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2388)+24)) = v2395
	v2397 = *(*int64)(unsafe.Add(mBase, uint32(v2382)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2388)+16)) = v2397
	v2399 = *(*int64)(unsafe.Add(mBase, uint32(v2382)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2388)+8)) = v2399
	v2401 = *(*int64)(unsafe.Add(mBase, uint32(v2382)))
	*(*int64)(unsafe.Add(mBase, uint32(v2388))) = v2401
	if v1767 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v2380<<(uint(int32(2))%32)))) = v2406
	goto L344
L343:
	;
	goto L344
L344:
	;
	v2409 = v2380 + int32(1)
	if v2381 == int32(0) {
		v2415 = v2409
		goto L340
	} else {
		goto L345
	}
L345:
	;
	F__bt_mark_scankey_required(m, v2388)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L24
	} else {
		goto L346
	}
L346:
	;
	v2415 = v2409
	goto L340
L347:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+80))
	if v2450 == int32(0) {
		v2483 = v2449
		goto L354
	} else {
		goto L355
	}
L348:
	;
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v2422 = v2419 + v2415*int32(56)
	v2423 = *(*int64)(unsafe.Add(mBase, uint32(v2416)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2422)+48)) = v2423
	v2425 = *(*int64)(unsafe.Add(mBase, uint32(v2416)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2422)+40)) = v2425
	v2427 = *(*int64)(unsafe.Add(mBase, uint32(v2416)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2422)+32)) = v2427
	v2429 = *(*int64)(unsafe.Add(mBase, uint32(v2416)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2422)+24)) = v2429
	v2431 = *(*int64)(unsafe.Add(mBase, uint32(v2416)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2422)+16)) = v2431
	v2433 = *(*int64)(unsafe.Add(mBase, uint32(v2416)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2422)+8)) = v2433
	v2435 = *(*int64)(unsafe.Add(mBase, uint32(v2416)))
	*(*int64)(unsafe.Add(mBase, uint32(v2422))) = v2435
	if v1767 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v2415<<(uint(int32(2))%32)))) = v2440
	goto L351
L350:
	;
	goto L351
L351:
	;
	v2443 = v2415 + int32(1)
	if v2381 == int32(0) {
		v2449 = v2443
		goto L347
	} else {
		goto L352
	}
L352:
	;
	F__bt_mark_scankey_required(m, v2422)
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L24
	} else {
		goto L353
	}
L353:
	;
	v2449 = v2443
	goto L347
L354:
	;
	if v1955 == v1759 {
		goto L251
	} else {
		goto L361
	}
L355:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v2456 = v2453 + v2449*int32(56)
	v2457 = *(*int64)(unsafe.Add(mBase, uint32(v2450)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2456)+48)) = v2457
	v2459 = *(*int64)(unsafe.Add(mBase, uint32(v2450)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2456)+40)) = v2459
	v2461 = *(*int64)(unsafe.Add(mBase, uint32(v2450)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2456)+32)) = v2461
	v2463 = *(*int64)(unsafe.Add(mBase, uint32(v2450)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2456)+24)) = v2463
	v2465 = *(*int64)(unsafe.Add(mBase, uint32(v2450)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2456)+16)) = v2465
	v2467 = *(*int64)(unsafe.Add(mBase, uint32(v2450)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2456)+8)) = v2467
	v2469 = *(*int64)(unsafe.Add(mBase, uint32(v2450)))
	*(*int64)(unsafe.Add(mBase, uint32(v2456))) = v2469
	if v1767 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v2449<<(uint(int32(2))%32)))) = v2474
	goto L358
L357:
	;
	goto L358
L358:
	;
	v2477 = v2449 + int32(1)
	if v2381 == int32(0) {
		v2483 = v2477
		goto L354
	} else {
		goto L359
	}
L359:
	;
	F__bt_mark_scankey_required(m, v2456)
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L24
	} else {
		goto L360
	}
L360:
	;
	v2483 = v2477
	goto L354
L361:
	;
	v2484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1989)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1751)+136)) = int32(0)
	v2487 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+128)) = v2487
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+120)) = v2487
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+112)) = v2487
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+104)) = v2487
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+96)) = v2487
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+88)) = v2487
	*(*int64)(unsafe.Add(mBase, uint32(v1751)+80)) = v2487
	v2501 = v2379
	v2502 = v2222
	v2505 = v2483
	v2509 = v2484
	goto L252
L362:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v1989)))
	v2522 = int32(base.Ui32(v2516)>>(uint(int32(5))%32))&int32(1) + v1953
	goto L364
L363:
	;
	v2522 = v1953
	goto L364
L364:
	;
	v2527 = v1751 + int32(80) + v2513*int32(12)
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2527)))
	if v2528 == int32(0) {
		v2615 = v2501
		v2617 = v2505
		goto L366
	} else {
		goto L367
	}
L365:
	;
	if v2513 != int32(2) {
		goto L250
	} else {
		goto L386
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2527)+8)) = v2522
	*(*int32)(unsafe.Add(mBase, uint32(v2527)+4)) = v1955
	*(*int32)(unsafe.Add(mBase, uint32(v2527))) = v1989
	v1946 = v2615
	v1948 = v2502
	v1952 = v2617
	v1953 = v2522
	v1955 = v1955 + int32(1)
	v1965 = v2509
	goto L220
L367:
	;
	v2531 = int32(0)
	v2534 = base.B2i32(v2513 != int32(2))
	if v2534|(v1767^int32(1)) != 0 {
		v2569 = v2531
		v2570 = v2531
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2573 = F__bt_compare_scankey_args(m, v1758, v1989, v1989, v2528, v2570, v2569, v1751+int32(212))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L24
	} else {
		goto L374
	}
L369:
	;
	v2538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1989))))
	if v2538&int32(32) != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v2569 = v2541 + v1955*int32(28)
	v2570 = v2545 + v2522<<(uint(int32(5))%32) - int32(32)
	goto L368
L371:
	;
	goto L372
L372:
	;
	v2551 = int32(0)
	v2552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2528))))
	if v2552&int32(32) == v2551 {
		v2569 = v2531
		v2570 = v2551
		goto L368
	} else {
		goto L373
	}
L373:
	;
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+4))
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+8))
	v2569 = v2557 + v2558*int32(28)
	v2570 = v2562 + v2563<<(uint(int32(5))%32) - int32(32)
	goto L368
L374:
	;
	if v2573 != 0 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v2575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+212)))
	if v2575 != int32(1) {
		goto L365
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v2586 = v2583 + v2505*int32(56)
	v2587 = *(*int64)(unsafe.Add(mBase, uint32(v2528)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2586)+48)) = v2587
	v2589 = *(*int64)(unsafe.Add(mBase, uint32(v2528)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2586)+40)) = v2589
	v2591 = *(*int64)(unsafe.Add(mBase, uint32(v2528)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2586)+32)) = v2591
	v2593 = *(*int64)(unsafe.Add(mBase, uint32(v2528)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2586)+24)) = v2593
	v2595 = *(*int64)(unsafe.Add(mBase, uint32(v2528)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2586)+16)) = v2595
	v2597 = *(*int64)(unsafe.Add(mBase, uint32(v2528)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2586)+8)) = v2597
	v2599 = *(*int64)(unsafe.Add(mBase, uint32(v2528)))
	*(*int64)(unsafe.Add(mBase, uint32(v2586))) = v2599
	if v1767 != 0 {
		goto L381
	} else {
		goto L382
	}
L378:
	;
	if v2513 != int32(2) {
		v2615 = v2501
		v2617 = v2505
		goto L366
	} else {
		goto L379
	}
L379:
	;
	v2578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2528))))
	if v2578&int32(32) == int32(0) {
		v2615 = v2501
		v2617 = v2505
		goto L366
	} else {
		goto L380
	}
L380:
	;
	goto L250
L381:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v2505<<(uint(int32(2))%32)))) = v2604
	goto L383
L382:
	;
	goto L383
L383:
	;
	v2606 = int32(1)
	v2607 = v2505 + v2606
	if v2502 != base.I32_extend16_s(v2509)-v2606 {
		v2615 = v2606
		v2617 = v2607
		goto L366
	} else {
		goto L384
	}
L384:
	;
	F__bt_mark_scankey_required(m, v2586)
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		goto L24
	} else {
		goto L385
	}
L385:
	;
	v2615 = v2606
	v2617 = v2607
	goto L366
L386:
	;
	v2626 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v2626)
	v4145 = v1748
	v4148 = v1751
	v4155 = v1758
	v4157 = v1760
	v4160 = v1763
	v4165 = v1768
	v4166 = v1769
	v4170 = v1773
	v4171 = v1774
	v4173 = v1776
	v4176 = v1779
	v4177 = v1780
	goto L1
L387:
	;
	v2629 = int32(0)
	v2631 = m.G0
	v2633 = v2631 - int32(16)
	m.G0 = v2633
	v2635 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+36))
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+12))
	if v2636 == v2629 {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	goto L389
L389:
	;
	if v2379&int32(1) == int32(0) {
		v4145 = v1748
		v4148 = v1751
		v4155 = v1758
		v4157 = v1760
		v4160 = v1763
		v4165 = v1768
		v4166 = v1769
		v4170 = v1773
		v4171 = v1774
		v4173 = v1776
		v4176 = v1779
		v4177 = v1780
		goto L1
	} else {
		goto L466
	}
L390:
	;
	m.G0 = v2633 + int32(16)
	goto L389
L391:
	;
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+4))
	if int32(0) < v2639 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v2642 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+4))
	v2643 = v2629
	v2652 = v2629
	goto L395
L393:
	;
	goto L394
L394:
	;
	v3043 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+88))
	if v3043 == int32(0) {
		goto L390
	} else {
		goto L460
	}
L395:
	;
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+8))
	v2687 = v2684 + v2652*int32(56)
	v2688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2687)+6)))
	if v2688 != int32(3) {
		v2957 = v2643
		goto L397
	} else {
		goto L398
	}
L396:
	;
	goto L394
L397:
	;
	v2999 = v2652 + int32(1)
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+4))
	if v2999 < v3000 {
		v2643 = v2957
		v2652 = v2999
		goto L395
	} else {
		goto L459
	}
L398:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2687)))
	if v2691&int32(32) == int32(0) {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	if v2691&int32(_a_F__bt_first_11) != int32(_a_F__bt_first_12) {
		v2957 = v2643
		goto L397
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+24))
	v2718 = int32(28)
	v2720 = v2717 + v2652*v2718
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v1770+v2652<<(uint(int32(2))%32))))
	v2727 = v2717 + v2724*v2718
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v2727)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2720)+24)) = v2728
	v2730 = *(*int64)(unsafe.Add(mBase, uint32(v2727)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2720)+16)) = v2730
	v2732 = *(*int64)(unsafe.Add(mBase, uint32(v2727)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2720)+8)) = v2732
	v2734 = *(*int64)(unsafe.Add(mBase, uint32(v2727)))
	*(*int64)(unsafe.Add(mBase, uint32(v2720))) = v2734
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+12))
	if v2736 <= v2643 {
		v2957 = v2643
		goto L397
	} else {
		goto L407
	}
L402:
	;
	v2700 = *(*int32)(unsafe.Add(mBase, uint32(v2687)+8))
	if v2700 != 0 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v2709 = v2700
	goto L405
L404:
	;
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(v2642)+212))
	v2702 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2687)+4)))
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v2701+v2702<<(uint(int32(2))%32)-int32(4))))
	v2709 = v2708
	goto L405
L405:
	;
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+24))
	F__bt_setup_array_cmp(m, v1758, v2687, v2709, v2710+v2652*int32(28), int32(0))
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L24
	} else {
		goto L406
	}
L406:
	;
	v2957 = v2643
	goto L397
L407:
	;
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+20))
	v2739 = v2643
	goto L408
L408:
	;
	v2782 = v2738 + v2739<<(uint(int32(5))%32)
	v2783 = *(*int32)(unsafe.Add(mBase, uint32(v2782)))
	if v2724 == v2783 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v2957 = v2736
	goto L397
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2782))) = v2652
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2782)+4))
	switch v2786 + int32(1) {
	case 0:
		goto L414
	default:
		goto L413
	case 2:
		goto L415
	}
L411:
	;
	goto L412
L412:
	;
	v2955 = v2739 + int32(1)
	if v2955 != v2736 {
		v2739 = v2955
		goto L408
	} else {
		goto L458
	}
L413:
	;
	v2957 = v2739 + int32(1)
	goto L397
L414:
	;
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2782)+20))
	if v2810 == int32(0) {
		goto L413
	} else {
		goto L418
	}
L415:
	;
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2687)))
	*(*int32)(unsafe.Add(mBase, uint32(v2687))) = v2789 & int32(-33)
	v2793 = *(*int32)(unsafe.Add(mBase, uint32(v2782)+8))
	v2794 = *(*int64)(unsafe.Add(mBase, uint32(v2793)))
	*(*int64)(unsafe.Add(mBase, uint32(v2687)+48)) = v2794
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+12))
	v2798 = v2796 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2635)+12)) = v2798
	if v2798 == int32(0) {
		goto L390
	} else {
		goto L416
	}
L416:
	;
	v2804 = (v2798 - v2739) << (uint(int32(5)) % 32)
	if v2804 == int32(0) {
		v2957 = v2739
		goto L397
	} else {
		goto L417
	}
L417:
	;
	base.MemoryCopy(m, v2782, v2782+int32(32), v2804)
	v2957 = v2739
	goto L397
L418:
	;
	v2813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2782)+19)))
	if v2813 != 0 {
		goto L413
	} else {
		goto L419
	}
L419:
	;
	v2814 = int32(_a_F__bt_first_1)
	v2815 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	v2817 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+36))
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(v2817)+28))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v2818
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v2782)+28))
	if v2820 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v2782)+24))
	if v2880 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L421:
	;
	v2823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2820)+6)))
	if v2823 != int32(1) {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	v2826 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2687)+4)))
	v2830 = v2826<<(uint(int32(2))%32) - int32(4)
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+4))
	v2832 = *(*int32)(unsafe.Add(mBase, uint32(v2831)+208))
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(v2830+v2832)))
	v2835 = *(*int64)(unsafe.Add(mBase, uint32(v2820)+48))
	v2836 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+8))
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2831)+212))
	v2839 = *(*int32)(unsafe.Add(mBase, uint32(v2837+v2830)))
	if v2836 != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v2842 = base.B2i32(v2836 != v2839)
	goto L425
L424:
	;
	v2842 = int32(0)
	goto L425
L425:
	;
	if v2842 != 0 {
		goto L420
	} else {
		goto L426
	}
L426:
	;
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2782)+20))
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v2845)+16))
	v2847 = m.T0[v2846].(func(*base.Module, int32, int64, int32) int64)(m, v2831, v2835, v2633+int32(14))
	mBase = m.M
	v2848 = m.ExcPending
	if v2848 != 0 {
		goto L24
	} else {
		goto L427
	}
L427:
	;
	v2849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2633)+14)))
	if v2849 == int32(1) {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+36))
	v2853 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2852))) = uint8(v2853)
	goto L420
L429:
	;
	goto L430
L430:
	;
	v2857 = *(*int32)(unsafe.Add(mBase, uint32(v2820)))
	if v2857&int32(16777216) != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v2860 = int32(4)
	goto L433
L432:
	;
	v2860 = int32(2)
	goto L433
L433:
	;
	v2861 = F_get_opfamily_member(m, v2834, v2839, v2839, v2860)
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L24
	} else {
		goto L434
	}
L434:
	;
	if v2861 == int32(0) {
		goto L420
	} else {
		goto L435
	}
L435:
	;
	v2865 = F_get_opcode(m, v2861)
	mBase = m.M
	v2866 = m.ExcPending
	if v2866 != 0 {
		goto L24
	} else {
		goto L436
	}
L436:
	;
	if v2865 == int32(0) {
		goto L420
	} else {
		goto L437
	}
L437:
	;
	F_fmgr_info(m, v2865, v2820+int32(16))
	mBase = m.M
	v2872 = m.ExcPending
	if v2872 != 0 {
		goto L24
	} else {
		goto L438
	}
L438:
	;
	v2873 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v2820)+6)) = uint16(v2873)
	*(*int64)(unsafe.Add(mBase, uint32(v2820)+48)) = v2847
	goto L420
L439:
	;
	*(*int32)(unsafe.Add(mBase, _c_F__bt_first[0])) = v2815
	goto L413
L440:
	;
	v2883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2880)+6)))
	if v2883 != int32(5) {
		goto L439
	} else {
		goto L441
	}
L441:
	;
	v2886 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2687)+4)))
	v2890 = v2886<<(uint(int32(2))%32) - int32(4)
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+4))
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+208))
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v2890+v2892)))
	v2895 = *(*int64)(unsafe.Add(mBase, uint32(v2880)+48))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+8))
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v2891)+212))
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v2897+v2890)))
	if v2896 != 0 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v2902 = base.B2i32(v2896 != v2899)
	goto L444
L443:
	;
	v2902 = int32(0)
	goto L444
L444:
	;
	if v2902 != 0 {
		goto L439
	} else {
		goto L445
	}
L445:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2782)+20))
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v2905)+20))
	v2907 = m.T0[v2906].(func(*base.Module, int32, int64, int32) int64)(m, v2891, v2895, v2633+int32(15))
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L24
	} else {
		goto L446
	}
L446:
	;
	v2909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2633)+15)))
	if v2909 == int32(1) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2912 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+36))
	v2913 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2912))) = uint8(v2913)
	goto L439
L448:
	;
	goto L449
L449:
	;
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v2880)))
	if v2917&int32(16777216) != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2920 = int32(2)
	goto L452
L451:
	;
	v2920 = int32(4)
	goto L452
L452:
	;
	v2921 = F_get_opfamily_member(m, v2894, v2899, v2899, v2920)
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L24
	} else {
		goto L453
	}
L453:
	;
	if v2921 == int32(0) {
		goto L439
	} else {
		goto L454
	}
L454:
	;
	v2925 = F_get_opcode(m, v2921)
	mBase = m.M
	v2926 = m.ExcPending
	if v2926 != 0 {
		goto L24
	} else {
		goto L455
	}
L455:
	;
	if v2925 == int32(0) {
		goto L439
	} else {
		goto L456
	}
L456:
	;
	F_fmgr_info(m, v2925, v2880+int32(16))
	mBase = m.M
	v2932 = m.ExcPending
	if v2932 != 0 {
		goto L24
	} else {
		goto L457
	}
L457:
	;
	v2933 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v2880)+6)) = uint16(v2933)
	*(*int64)(unsafe.Add(mBase, uint32(v2880)+48)) = v2907
	goto L439
L458:
	;
	goto L409
L459:
	;
	goto L396
L460:
	;
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+12))
	if v3046 < int32(33) {
		goto L390
	} else {
		goto L461
	}
L461:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3052 = m.ExcPending
	if v3052 != 0 {
		goto L24
	} else {
		goto L462
	}
L462:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L24
	} else {
		goto L463
	}
L463:
	;
	v3056 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2633)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v2633))) = v3056
	F_errmsg_internal(m, int32(_a_F__bt_first_13), v2633)
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L24
	} else {
		goto L464
	}
L464:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(2375), int32(_a_F__bt_first_14))
	mBase = m.M
	v3067 = m.ExcPending
	if v3067 != 0 {
		goto L24
	} else {
		goto L465
	}
L465:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L466:
	;
	v3157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v3157 != int32(1) {
		v4145 = v1748
		v4148 = v1751
		v4155 = v1758
		v4157 = v1760
		v4160 = v1763
		v4165 = v1768
		v4166 = v1769
		v4170 = v1773
		v4171 = v1774
		v4173 = v1776
		v4176 = v1779
		v4177 = v1780
		goto L1
	} else {
		goto L467
	}
L467:
	;
	v3160 = int32(0)
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+36))
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	v3169 = F_palloc0(m, v3168)
	mBase = m.M
	v3170 = m.ExcPending
	if v3170 != 0 {
		goto L24
	} else {
		goto L468
	}
L468:
	;
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	if int32(0) < v3171 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+8))
	v3175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3174)+4)))
	v3181 = v3160
	v3183 = v3160
	v3184 = v3160
	v3186 = v3160
	v3190 = v3175
	v3191 = v3160
	v3199 = v3160
	goto L472
L470:
	;
	v3558 = v3160
	goto L471
L471:
	;
	v3593 = F_palloc(m, v3558*int32(56))
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L24
	} else {
		goto L515
	}
L472:
	;
	v3217 = int32(0)
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+8))
	v3222 = v3219 + v3191*int32(56)
	v3223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3222)+4)))
	if v3223 == v3190&int32(_a_F__bt_first_0) {
		goto L477
	} else {
		goto L478
	}
L473:
	;
	v3558 = v3513
	goto L471
L474:
	;
	v3547 = v3191 + int32(1)
	v3548 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	if v3547 < v3548 {
		v3181 = v3510
		v3183 = v3512
		v3184 = v3513
		v3186 = v3515
		v3190 = v3519
		v3191 = v3547
		v3199 = v3528
		goto L472
	} else {
		goto L514
	}
L475:
	;
	v3510 = v3232
	v3512 = v3233
	v3513 = v3472
	v3515 = v3230
	v3519 = v3231
	v3528 = int32(1)
	goto L474
L476:
	;
	v3460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3191+v3169))) = uint8(v3460)
	v3510 = v3451
	v3512 = v3453
	v3513 = v3184 + v3460
	v3515 = v3455
	v3519 = v3456
	v3528 = v3454
	goto L474
L477:
	;
	v3227 = int32(1)
	if v3199&v3227 != 0 {
		v3451 = v3181
		v3453 = v3183
		v3454 = v3227
		v3455 = v3186
		v3456 = v3190
		goto L476
	} else {
		goto L480
	}
L478:
	;
	v3230 = v3191
	v3231 = v3223
	v3232 = v3217
	v3233 = v3217
	goto L479
L479:
	;
	v3235 = *(*int32)(unsafe.Add(mBase, uint32(v3222)))
	v3236 = int32(_a_F__bt_first_15)
	if v3235&v3236 == v3236 {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v3230 = v3186
	v3231 = v3190
	v3232 = v3181
	v3233 = v3183
	goto L479
L481:
	;
	if v3191 <= v3230 {
		v3472 = v3184
		goto L475
	} else {
		goto L484
	}
L482:
	;
	goto L483
L483:
	;
	v3428 = int32(1)
	v3429 = int32(0)
	if (base.B2i32(v3235&int32(_a_F__bt_first_12) == v3429)|v3233)&v3428 == v3429 {
		goto L510
	} else {
		goto L511
	}
L484:
	;
	v3245 = (v3191 - v3230) & int32(3)
	if v3245 != 0 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v3254 = v3184
	v3260 = v3230
	v3265 = int32(0)
	goto L488
L486:
	;
	v3309 = v3184
	v3315 = v3230
	goto L487
L487:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v3230-v3191) {
		v3472 = v3309
		goto L475
	} else {
		goto L494
	}
L488:
	;
	v3287 = v3260 + v3169
	v3288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3287))))
	if v3288 == int32(0) {
		goto L490
	} else {
		goto L491
	}
L489:
	;
	v3309 = v3295
	v3315 = v3297
	goto L487
L490:
	;
	v3291 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3287))) = uint8(v3291)
	v3295 = v3254 + v3291
	goto L492
L491:
	;
	v3295 = v3254
	goto L492
L492:
	;
	v3296 = int32(1)
	v3297 = v3260 + v3296
	v3299 = v3265 + v3296
	if v3299 != v3245 {
		v3254 = v3295
		v3260 = v3297
		v3265 = v3299
		goto L488
	} else {
		goto L493
	}
L493:
	;
	goto L489
L494:
	;
	v3353 = v3309
	v3359 = v3315
	goto L495
L495:
	;
	v3386 = v3359 + v3169
	v3387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3386))))
	if v3387 == int32(0) {
		goto L497
	} else {
		goto L498
	}
L496:
	;
	v3472 = v3424
	goto L475
L497:
	;
	v3390 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3386))) = uint8(v3390)
	v3394 = v3353 + v3390
	goto L499
L498:
	;
	v3394 = v3353
	goto L499
L499:
	;
	v3396 = v3386 + int32(1)
	v3397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3396))))
	if v3397 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v3400 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3396))) = uint8(v3400)
	v3404 = v3394 + v3400
	goto L502
L501:
	;
	v3404 = v3394
	goto L502
L502:
	;
	v3406 = v3386 + int32(2)
	v3407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3406))))
	if v3407 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v3410 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3406))) = uint8(v3410)
	v3414 = v3404 + v3410
	goto L505
L504:
	;
	v3414 = v3404
	goto L505
L505:
	;
	v3416 = v3386 + int32(3)
	v3417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3416))))
	if v3417 == int32(0) {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v3420 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3416))) = uint8(v3420)
	v3424 = v3414 + v3420
	goto L508
L507:
	;
	v3424 = v3414
	goto L508
L508:
	;
	v3426 = v3359 + int32(4)
	if v3191 != v3426 {
		v3353 = v3424
		v3359 = v3426
		goto L495
	} else {
		goto L509
	}
L509:
	;
	goto L496
L510:
	;
	v3510 = v3232
	v3512 = v3428
	v3513 = v3184
	v3515 = v3230
	v3519 = v3231
	v3528 = v3429
	goto L474
L511:
	;
	goto L512
L512:
	;
	v3439 = int32(0)
	if (v3232|base.B2i32(v3235&int32(_a_F__bt_first_16) == v3439))&int32(1) != 0 {
		v3451 = v3232
		v3453 = v3233
		v3454 = v3439
		v3455 = v3230
		v3456 = v3231
		goto L476
	} else {
		goto L513
	}
L513:
	;
	v3510 = int32(1)
	v3512 = v3233
	v3513 = v3184
	v3515 = v3230
	v3519 = v3231
	v3528 = v3429
	goto L474
L514:
	;
	goto L473
L515:
	;
	v3595 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	v3599 = F_palloc(m, (v3595-v3558)*int32(56))
	mBase = m.M
	v3600 = m.ExcPending
	if v3600 != 0 {
		goto L24
	} else {
		goto L516
	}
L516:
	;
	v3602 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+12))
	if v3602 != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v3605 = F_palloc(m, v3558*int32(28))
	mBase = m.M
	v3606 = m.ExcPending
	if v3606 != 0 {
		goto L24
	} else {
		goto L520
	}
L518:
	;
	v3613 = int32(0)
	v3614 = v3160
	goto L519
L519:
	;
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	if v3615 <= int32(0) {
		goto L523
	} else {
		goto L524
	}
L520:
	;
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	v3611 = F_palloc(m, (v3607-v3558)*int32(28))
	mBase = m.M
	v3612 = m.ExcPending
	if v3612 != 0 {
		goto L24
	} else {
		goto L521
	}
L521:
	;
	v3613 = v3605
	v3614 = v3611
	goto L519
L522:
	;
	v3943 = v3908 * int32(56)
	if v3943 != 0 {
		goto L545
	} else {
		goto L546
	}
L523:
	;
	v3618 = int32(0)
	v3906 = v3618
	v3908 = v3618
	goto L522
L524:
	;
	goto L525
L525:
	;
	v3620 = int32(0)
	v3628 = v3620
	v3630 = v3620
	v3638 = v3620
	goto L526
L526:
	;
	v3664 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+8))
	v3667 = v3664 + v3638*int32(56)
	v3669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3638+v3169))))
	if v3669 == int32(0) {
		goto L529
	} else {
		goto L530
	}
L527:
	;
	v3906 = v3861
	v3908 = v3863
	goto L522
L528:
	;
	v3898 = v3638 + int32(1)
	v3899 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	if v3898 < v3899 {
		v3628 = v3861
		v3630 = v3863
		v3638 = v3898
		goto L526
	} else {
		goto L544
	}
L529:
	;
	v3674 = v3599 + v3630*int32(56)
	v3675 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v3674)+48)) = v3675
	v3677 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3674)+40)) = v3677
	v3679 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v3674)+32)) = v3679
	v3681 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3674)+24)) = v3681
	v3683 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3674)+16)) = v3683
	v3685 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3674)+8)) = v3685
	v3687 = *(*int64)(unsafe.Add(mBase, uint32(v3667)))
	*(*int64)(unsafe.Add(mBase, uint32(v3674))) = v3687
	v3689 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+12))
	if v3689 != 0 {
		goto L532
	} else {
		goto L533
	}
L530:
	;
	goto L531
L531:
	;
	v3715 = v3593 + v3628*int32(56)
	v3716 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v3715)+48)) = v3716
	v3718 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3715)+40)) = v3718
	v3720 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v3715)+32)) = v3720
	v3722 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3715)+24)) = v3722
	v3724 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3715)+16)) = v3724
	v3726 = *(*int64)(unsafe.Add(mBase, uint32(v3667)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3715)+8)) = v3726
	v3728 = *(*int64)(unsafe.Add(mBase, uint32(v3667)))
	*(*int64)(unsafe.Add(mBase, uint32(v3715))) = v3728
	v3730 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+12))
	if v3730 != 0 {
		goto L535
	} else {
		goto L536
	}
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v3638<<(uint(int32(2))%32)))) = v3630
	v3694 = int32(28)
	v3696 = v3614 + v3630*v3694
	v3697 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+24))
	v3700 = v3697 + v3638*v3694
	v3701 = *(*int32)(unsafe.Add(mBase, uint32(v3700)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3696)+24)) = v3701
	v3703 = *(*int64)(unsafe.Add(mBase, uint32(v3700)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3696)+16)) = v3703
	v3705 = *(*int64)(unsafe.Add(mBase, uint32(v3700)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3696)+8)) = v3705
	v3707 = *(*int64)(unsafe.Add(mBase, uint32(v3700)))
	*(*int64)(unsafe.Add(mBase, uint32(v3696))) = v3707
	goto L534
L533:
	;
	goto L534
L534:
	;
	v3861 = v3628
	v3863 = v3630 + int32(1)
	goto L528
L535:
	;
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1770+v3638<<(uint(int32(2))%32)))) = v3734 + (v3628 - v3558)
	v3738 = int32(28)
	v3740 = v3613 + v3628*v3738
	v3741 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+24))
	v3744 = v3741 + v3638*v3738
	v3745 = *(*int32)(unsafe.Add(mBase, uint32(v3744)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3740)+24)) = v3745
	v3747 = *(*int64)(unsafe.Add(mBase, uint32(v3744)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3740)+16)) = v3747
	v3749 = *(*int64)(unsafe.Add(mBase, uint32(v3744)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3740)+8)) = v3749
	v3751 = *(*int64)(unsafe.Add(mBase, uint32(v3744)))
	*(*int64)(unsafe.Add(mBase, uint32(v3740))) = v3751
	goto L537
L536:
	;
	goto L537
L537:
	;
	v3755 = *(*int32)(unsafe.Add(mBase, uint32(v3715)))
	*(*int32)(unsafe.Add(mBase, uint32(v3715))) = v3755 & int32(-196609)
	if v3755&int32(4) != 0 {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v3761 = *(*int32)(unsafe.Add(mBase, uint32(v3715)+48))
	v3776 = v3761
	goto L541
L539:
	;
	goto L540
L540:
	;
	v3861 = v3628 + int32(1)
	v3863 = v3630
	goto L528
L541:
	;
	v3803 = *(*int32)(unsafe.Add(mBase, uint32(v3776)))
	*(*int32)(unsafe.Add(mBase, uint32(v3776))) = v3803 & int32(-196609)
	if v3803&int32(16) == int32(0) {
		v3776 = v3776 + int32(56)
		goto L541
	} else {
		goto L543
	}
L542:
	;
	goto L540
L543:
	;
	goto L542
L544:
	;
	goto L527
L545:
	;
	v3944 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+8))
	base.MemoryCopy(m, v3944, v3599, v3943)
	goto L547
L546:
	;
	goto L547
L547:
	;
	v3947 = v3906 * int32(56)
	if v3947 != 0 {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	v3948 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+8))
	base.MemoryCopy(m, v3948+v3943, v3593, v3947)
	goto L550
L549:
	;
	goto L550
L550:
	;
	F_pfree(m, v3169)
	mBase = m.M
	v3952 = m.ExcPending
	if v3952 != 0 {
		goto L24
	} else {
		goto L551
	}
L551:
	;
	F_pfree(m, v3599)
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
		goto L24
	} else {
		goto L552
	}
L552:
	;
	F_pfree(m, v3593)
	mBase = m.M
	v3956 = m.ExcPending
	if v3956 != 0 {
		goto L24
	} else {
		goto L553
	}
L553:
	;
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+12))
	if v3957 != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v3959 = v3908 * int32(28)
	if v3959 != 0 {
		goto L557
	} else {
		goto L558
	}
L555:
	;
	goto L556
L556:
	;
	v4145 = v1748
	v4148 = v1751
	v4155 = v1758
	v4157 = v1760
	v4160 = v1763
	v4165 = v1768
	v4166 = v1769
	v4170 = v1773
	v4171 = v1774
	v4173 = v1776
	v4176 = v1779
	v4177 = v1780
	goto L1
L557:
	;
	v3960 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+24))
	base.MemoryCopy(m, v3960, v3614, v3959)
	goto L559
L558:
	;
	goto L559
L559:
	;
	v3963 = v3906 * int32(28)
	if v3963 != 0 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v3964 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+24))
	base.MemoryCopy(m, v3964+v3959, v3613, v3963)
	goto L562
L561:
	;
	goto L562
L562:
	;
	v3967 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+12))
	if int32(0) < v3967 {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	v3985 = int32(0)
	goto L566
L564:
	;
	v4034 = v3967
	goto L565
L565:
	;
	v4067 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+20))
	F_pg_qsort(m, v4067, v4034, int32(32), int32(216))
	mBase = m.M
	v4071 = m.ExcPending
	if v4071 != 0 {
		goto L24
	} else {
		goto L569
	}
L566:
	;
	v4012 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+20))
	v4015 = v4012 + v3985<<(uint(int32(5))%32)
	v4016 = *(*int32)(unsafe.Add(mBase, uint32(v4015)))
	v4020 = *(*int32)(unsafe.Add(mBase, uint32(v1770+v4016<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4015))) = v4020
	v4023 = v3985 + int32(1)
	v4024 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+12))
	if v4023 < v4024 {
		v3985 = v4023
		goto L566
	} else {
		goto L568
	}
L567:
	;
	v4034 = v4024
	goto L565
L568:
	;
	goto L567
L569:
	;
	F_pfree(m, v3613)
	mBase = m.M
	v4073 = m.ExcPending
	if v4073 != 0 {
		goto L24
	} else {
		goto L570
	}
L570:
	;
	F_pfree(m, v3614)
	mBase = m.M
	v4075 = m.ExcPending
	if v4075 != 0 {
		goto L24
	} else {
		goto L571
	}
L571:
	;
	goto L556
L572:
	;
	F_errmsg_internal(m, int32(_a_F__bt_first_9), int32(0))
	mBase = m.M
	v4126 = m.ExcPending
	if v4126 != 0 {
		goto L24
	} else {
		goto L573
	}
L573:
	;
	F_errfinish(m, int32(_a_F__bt_first_6), int32(269), int32(_a_F__bt_first_10))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L24
	} else {
		goto L574
	}
L574:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L575:
	;
	m.G0 = v4157 + int32(2064)
	return v5475
L576:
	;
	F__bt_parallel_done(m, v4155)
	mBase = m.M
	v5457 = m.ExcPending
	if v5457 != 0 {
		goto L24
	} else {
		goto L759
	}
L577:
	;
	v4191 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+88))
	if v4191 == int32(0) {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v4201 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+12))
	if v4201 == int32(0) {
		goto L582
	} else {
		goto L583
	}
L579:
	;
	v4199 = F__bt_parallel_seize(m, v4155, v4157+int32(68), v4157-int32(-64), int32(1))
	mBase = m.M
	v4200 = m.ExcPending
	if v4200 != 0 {
		goto L24
	} else {
		goto L580
	}
L580:
	;
	if v4199 != 0 {
		goto L578
	} else {
		goto L581
	}
L581:
	;
	v5475 = v4160
	goto L575
L582:
	;
	v4207 = *(*int32)(unsafe.Add(mBase, uint32(v4157)+68))
	if v4207 != int32(-1) {
		goto L586
	} else {
		goto L587
	}
L583:
	;
	v4204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4166)+17)))
	if v4204 != 0 {
		goto L582
	} else {
		goto L584
	}
L584:
	;
	F__bt_start_array_keys(m, v4155, v4145)
	mBase = m.M
	v4206 = m.ExcPending
	if v4206 != 0 {
		goto L24
	} else {
		goto L585
	}
L585:
	;
	goto L582
L586:
	;
	v4210 = int32(1)
	v4211 = *(*int32)(unsafe.Add(mBase, uint32(v4157)+64))
	v4213 = F__bt_readnextpage(m, v4155, v4207, v4211, v4145, v4210)
	mBase = m.M
	v4214 = m.ExcPending
	if v4214 != 0 {
		goto L24
	} else {
		goto L589
	}
L587:
	;
	goto L588
L588:
	;
	v4234 = *(*int32)(unsafe.Add(mBase, uint32(v4165)+272))
	if v4234 == int32(0) {
		goto L595
	} else {
		goto L596
	}
L589:
	;
	if v4213 == int32(0) {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	v5475 = int32(0)
	goto L575
L591:
	;
	goto L592
L592:
	;
	v4218 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+100))
	v4221 = v4166 + v4218*int32(10)
	v4222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4221)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4155)+64)) = uint16(v4222)
	v4225 = v4221 + int32(104)
	v4226 = *(*int32)(unsafe.Add(mBase, uint32(v4225)))
	*(*int32)(unsafe.Add(mBase, uint32(v4155)+60)) = v4226
	v4228 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+44))
	if v4228 == int32(0) {
		v5475 = v4210
		goto L575
	} else {
		goto L593
	}
L593:
	;
	v4231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4225)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v4155)+44)) = v4228 + v4231
	v5475 = v4210
	goto L575
L594:
	;
	v4249 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+40))
	if v4249 != 0 {
		goto L600
	} else {
		goto L601
	}
L595:
	;
	v4237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4165)+268)))
	if v4237 != int32(1) {
		goto L594
	} else {
		goto L598
	}
L596:
	;
	v4243 = v4234
	goto L597
L597:
	;
	v4244 = *(*int64)(unsafe.Add(mBase, uint32(v4243)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4243)+16)) = v4244 + int64(1)
	goto L594
L598:
	;
	F_pgstat_assoc_relation(m, v4165)
	mBase = m.M
	v4241 = m.ExcPending
	if v4241 != 0 {
		goto L24
	} else {
		goto L599
	}
L599:
	;
	v4242 = *(*int32)(unsafe.Add(mBase, uint32(v4165)+272))
	v4243 = v4242
	goto L597
L600:
	;
	v4250 = *(*int64)(unsafe.Add(mBase, uint32(v4249)))
	*(*int64)(unsafe.Add(mBase, uint32(v4249))) = v4250 + int64(1)
	goto L602
L601:
	;
	goto L602
L602:
	;
	v4254 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+4))
	if v4254 <= int32(0) {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v5315 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+36))
	v5316 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+4))
	v5319 = base.B2i32(v4145 == int32(-1))
	v5320 = F__bt_get_endpoint(m, v5316, int32(0), v5319)
	mBase = m.M
	v5321 = m.ExcPending
	if v5321 != 0 {
		goto L24
	} else {
		goto L727
	}
L604:
	;
	v4258 = int32(1)
	v4260 = base.B2i32(v4145 == v4258)
	if v4145 == v4258 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	v4261 = int32(5)
	goto L607
L606:
	;
	v4261 = v4258
	goto L607
L607:
	;
	if v4145 == v4258 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v4264 = int32(24)
	goto L610
L609:
	;
	v4264 = int32(28)
	goto L610
L610:
	;
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+8))
	v4267 = v4265
	v4273 = int32(1)
	v4283 = v4254
	v4293 = v4170
	v4294 = v4171
	v4296 = v4173
	v4299 = v4176
	v4300 = v4177
	goto L611
L611:
	;
	if v4300 < v4283 {
		goto L617
	} else {
		goto L618
	}
L612:
	;
	if v4633 == int32(0) {
		goto L603
	} else {
		goto L662
	}
L613:
	;
	goto L612
L614:
	;
	v4267 = v4267 + int32(56)
	v4273 = v4567
	v4283 = v4577
	v4293 = v4587
	v4294 = v4588
	v4296 = v4590
	v4299 = v4593
	v4300 = v4300 + int32(1)
	goto L611
L615:
	;
	v4552 = int32(0)
	v4553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4267)+6)))
	switch v4553 - int32(1) {
	case 0, 1:
		goto L656
	case 2:
		goto L655
	case 3, 4:
		goto L654
	default:
		v4567 = v4517
		v4577 = v4527
		v4587 = v4552
		v4588 = v4538
		v4590 = v4540
		v4593 = v4543
		goto L614
	}
L616:
	;
	if v4293 != 0 {
		v4567 = v4273
		v4577 = v4283
		v4587 = v4293
		v4588 = v4294
		v4590 = v4296
		v4593 = v4299
		goto L614
	} else {
		goto L652
	}
L617:
	;
	v4309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4267)+4)))
	if v4309 == v4273 {
		goto L616
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	if v4293 != 0 {
		goto L622
	} else {
		goto L623
	}
L620:
	;
	goto L619
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157+int32(128)+v4294<<(uint(int32(2))%32)))) = v4471
	v4492 = int32(1)
	v4493 = v4294 + v4492
	v4494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4471)+6)))
	if v4494&int32(_a_F__bt_first_17) == v4492 {
		v4633 = v4493
		v4638 = v4494
		goto L613
	} else {
		goto L646
	}
L622:
	;
	v4311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4293)+2)))
	if v4311&int32(24) == int32(0) {
		v4471 = v4293
		goto L621
	} else {
		goto L625
	}
L623:
	;
	v4407 = v4296
	v4419 = int32(0)
	goto L624
L624:
	;
	v4420 = int32(0)
	if v4419|base.B2i32(v4407 == v4420) == v4420 {
		goto L635
	} else {
		goto L636
	}
L625:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+8))
	v4319 = base.I32_div_s(v4293-v4316, int32(56))
	v4320 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+20))
	v4321 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+12))
	v4339 = int32(0)
	goto L626
L626:
	;
	v4366 = v4320 + v4339<<(uint(int32(5))%32)
	v4367 = *(*int32)(unsafe.Add(mBase, uint32(v4366)))
	if v4319 != v4367 {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	v4373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4366)+19)))
	if v4373 != 0 {
		goto L632
	} else {
		goto L633
	}
L628:
	;
	v4370 = v4339 + int32(1)
	if base.Ui32(v4370) < base.Ui32(v4321) {
		v4339 = v4370
		goto L626
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	goto L627
L631:
	;
	goto L630
L632:
	;
	v4374 = v4296
	goto L634
L633:
	;
	v4374 = v4293
	goto L634
L634:
	;
	v4376 = *(*int32)(unsafe.Add(mBase, uint32(v4366+v4264)))
	v4407 = v4374
	v4419 = v4376
	goto L624
L635:
	;
	v4425 = *(*int32)(unsafe.Add(mBase, uint32(v4407)))
	if v4425&int32(33554432) != 0 {
		goto L639
	} else {
		goto L640
	}
L636:
	;
	goto L637
L637:
	;
	if v4419 == int32(0) {
		v4633 = v4294
		v4638 = v4299
		goto L613
	} else {
		goto L645
	}
L638:
	;
	v4431 = v4157 + int32(72)
	v4437 = int32(0)
	F_ScanKeyEntryInitialize(m, v4431, v4425&int32(50331648)|int32(129), base.I32_extend16_s(v4273), v4261, v4437, v4437, v4437, int64(0))
	mBase = m.M
	v4442 = m.ExcPending
	if v4442 != 0 {
		goto L24
	} else {
		goto L644
	}
L639:
	;
	if v4145 == v4258 {
		goto L638
	} else {
		goto L642
	}
L640:
	;
	goto L641
L641:
	;
	if v4145 != int32(-1) {
		v4633 = v4294
		v4638 = v4299
		goto L613
	} else {
		goto L643
	}
L642:
	;
	v4633 = v4294
	v4638 = v4299
	goto L613
L643:
	;
	goto L638
L644:
	;
	v4471 = v4431
	goto L621
L645:
	;
	v4471 = v4419
	goto L621
L646:
	;
	v4499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4471)+2)))
	if v4499&int32(96) != 0 {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v4633 = v4493
	v4638 = v4261
	goto L613
L648:
	;
	goto L649
L649:
	;
	v4502 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+4))
	if v4502 <= v4300 {
		v4633 = v4493
		v4638 = v4494
		goto L613
	} else {
		goto L650
	}
L650:
	;
	v4504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4267)+2)))
	if v4504&int32(3) == int32(0) {
		v4633 = v4493
		v4638 = v4494
		goto L613
	} else {
		goto L651
	}
L651:
	;
	v4509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4267)+4)))
	v4517 = v4509
	v4527 = v4502
	v4538 = v4493
	v4540 = int32(0)
	v4543 = v4494
	goto L615
L652:
	;
	v4517 = v4273
	v4527 = v4283
	v4538 = v4294
	v4540 = v4296
	v4543 = v4299
	goto L615
L653:
	;
	if v4540 != 0 {
		goto L659
	} else {
		goto L660
	}
L654:
	;
	if v4260 == int32(0) {
		goto L653
	} else {
		goto L658
	}
L655:
	;
	v4567 = v4517
	v4577 = v4527
	v4587 = v4267
	v4588 = v4538
	v4590 = v4540
	v4593 = v4543
	goto L614
L656:
	;
	if v4145 != int32(-1) {
		goto L653
	} else {
		goto L657
	}
L657:
	;
	goto L655
L658:
	;
	v4567 = v4517
	v4577 = v4527
	v4587 = v4267
	v4588 = v4538
	v4590 = v4540
	v4593 = v4543
	goto L614
L659:
	;
	v4560 = v4540
	goto L661
L660:
	;
	v4560 = v4267
	goto L661
L661:
	;
	v4567 = v4517
	v4577 = v4527
	v4587 = v4552
	v4588 = v4538
	v4590 = v4560
	v4593 = v4543
	goto L614
L662:
	;
	if v4633 <= int32(0) {
		v4883 = v4633
		goto L676
	} else {
		goto L677
	}
L663:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5247 = m.ExcPending
	if v5247 != 0 {
		goto L24
	} else {
		goto L724
	}
L664:
	;
	v5190 = int32(0)
	v5193 = v4157 + int32(256)
	v5195 = v4166 + int32(56)
	v5198 = F__bt_search(m, v4165, v5190, v5193, v5195, int32(1), v5190)
	mBase = m.M
	v5199 = m.ExcPending
	if v5199 != 0 {
		goto L24
	} else {
		goto L712
	}
L665:
	;
	v5147 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4157)+259)) = uint16(v5147)
	goto L664
L666:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5135 = m.ExcPending
	if v5135 != 0 {
		goto L24
	} else {
		goto L709
	}
L667:
	;
	v5130 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v4157)+259)) = uint16(v5130)
	goto L664
L668:
	;
	v5079 = v4157 + int32(256)
	F__bt_metaversion(m, v4165, v5079, v5079|int32(1))
	mBase = m.M
	v5083 = m.ExcPending
	if v5083 != 0 {
		goto L24
	} else {
		goto L708
	}
L669:
	;
	v5076 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4157)+259)) = uint16(v5076)
	goto L664
L670:
	;
	v5025 = v4157 + int32(256)
	F__bt_metaversion(m, v4165, v5025, v5025|int32(1))
	mBase = m.M
	v5029 = m.ExcPending
	if v5029 != 0 {
		goto L24
	} else {
		goto L707
	}
L671:
	;
	if v4145 != int32(-1) {
		goto L665
	} else {
		goto L706
	}
L672:
	;
	v5018 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v4157)+259)) = uint16(v5018)
	goto L664
L673:
	;
	v4967 = v4157 + int32(256)
	F__bt_metaversion(m, v4165, v4967, v4967|int32(1))
	mBase = m.M
	v4971 = m.ExcPending
	if v4971 != 0 {
		goto L24
	} else {
		goto L705
	}
L674:
	;
	v4964 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v4157)+259)) = uint16(v4964)
	goto L664
L675:
	;
	v4913 = v4157 + int32(256)
	F__bt_metaversion(m, v4165, v4913, v4913|int32(1))
	mBase = m.M
	v4917 = m.ExcPending
	if v4917 != 0 {
		goto L24
	} else {
		goto L704
	}
L676:
	;
	v4898 = v4157 + int32(256)
	F__bt_metaversion(m, v4165, v4898, v4898|int32(1))
	mBase = m.M
	v4902 = m.ExcPending
	if v4902 != 0 {
		goto L24
	} else {
		goto L703
	}
L677:
	;
	v4652 = v4157 + int32(272)
	v4656 = int32(0)
	goto L678
L678:
	;
	v4696 = v4656 << (uint(int32(2)) % 32)
	v4700 = *(*int32)(unsafe.Add(mBase, uint32(v4696+(v4157+int32(128)))))
	v4701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4700))))
	if v4701&int32(4) != 0 {
		goto L680
	} else {
		goto L681
	}
L679:
	;
	v4883 = v4633
	goto L676
L680:
	;
	v4706 = v4652 + v4656*int32(56)
	v4707 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+48))
	v4708 = *(*int64)(unsafe.Add(mBase, uint32(v4707)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v4706)+48)) = v4708
	v4710 = *(*int64)(unsafe.Add(mBase, uint32(v4707)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4706)+40)) = v4710
	v4712 = *(*int64)(unsafe.Add(mBase, uint32(v4707)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4706)+32)) = v4712
	v4714 = *(*int64)(unsafe.Add(mBase, uint32(v4707)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4706)+24)) = v4714
	v4716 = *(*int64)(unsafe.Add(mBase, uint32(v4707)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4706)+16)) = v4716
	v4718 = *(*int64)(unsafe.Add(mBase, uint32(v4707)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4706)+8)) = v4718
	v4720 = *(*int64)(unsafe.Add(mBase, uint32(v4707)))
	*(*int64)(unsafe.Add(mBase, uint32(v4706))) = v4720
	v4738 = v4707
	v4749 = v4633
	goto L684
L681:
	;
	goto L682
L682:
	;
	v4802 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+8))
	if v4802 != 0 {
		goto L693
	} else {
		goto L694
	}
L683:
	;
	switch v4638&int32(_a_F__bt_first_0) - int32(2) {
	case 0:
		goto L675
	default:
		v4883 = v4749
		goto L676
	case 2:
		goto L668
	}
L684:
	;
	v4763 = *(*int32)(unsafe.Add(mBase, uint32(v4738)+56))
	if v4763&int32(1) != 0 {
		goto L683
	} else {
		goto L686
	}
L685:
	;
	switch v4638&int32(_a_F__bt_first_0) - int32(1) {
	case 0:
		goto L673
	default:
		v4883 = v4749
		goto L676
	case 4:
		goto L670
	}
L686:
	;
	if v4763&int32(_a_F__bt_first_15) != 0 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v4768 = int32(56)
	v4770 = v4652 + v4749*v4768
	v4772 = v4738 + v4768
	v4773 = *(*int64)(unsafe.Add(mBase, uint32(v4772)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v4770)+48)) = v4773
	v4775 = *(*int64)(unsafe.Add(mBase, uint32(v4772)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4770)+40)) = v4775
	v4777 = *(*int64)(unsafe.Add(mBase, uint32(v4772)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4770)+32)) = v4777
	v4779 = *(*int64)(unsafe.Add(mBase, uint32(v4772)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4770)+24)) = v4779
	v4781 = *(*int64)(unsafe.Add(mBase, uint32(v4772)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4770)+16)) = v4781
	v4783 = *(*int64)(unsafe.Add(mBase, uint32(v4772)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4770)+8)) = v4783
	v4785 = *(*int64)(unsafe.Add(mBase, uint32(v4772)))
	*(*int64)(unsafe.Add(mBase, uint32(v4770))) = v4785
	v4788 = v4749 + int32(1)
	v4789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4772))))
	if v4789&int32(16) == int32(0) {
		v4738 = v4772
		v4749 = v4788
		goto L684
	} else {
		goto L690
	}
L688:
	;
	goto L689
L689:
	;
	goto L685
L690:
	;
	v4883 = v4788
	goto L676
L691:
	;
	v4854 = v4656 + int32(1)
	if v4854 != v4633 {
		v4656 = v4854
		goto L678
	} else {
		goto L702
	}
L692:
	;
	v4832 = *(*int32)(unsafe.Add(mBase, uint32(v4165)+208))
	v4834 = *(*int32)(unsafe.Add(mBase, uint32(v4832+v4696)))
	v4836 = F_get_opfamily_proc(m, v4834, v4805, v4802, int32(1))
	mBase = m.M
	v4837 = m.ExcPending
	if v4837 != 0 {
		goto L24
	} else {
		goto L699
	}
L693:
	;
	v4803 = *(*int32)(unsafe.Add(mBase, uint32(v4165)+212))
	v4805 = *(*int32)(unsafe.Add(mBase, uint32(v4803+v4696)))
	if v4802 != v4805 {
		goto L692
	} else {
		goto L696
	}
L694:
	;
	goto L695
L695:
	;
	v4808 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4700)+4)))
	v4810 = F_index_getprocinfo(m, v4165, v4808, int32(1))
	mBase = m.M
	v4811 = m.ExcPending
	if v4811 != 0 {
		goto L24
	} else {
		goto L697
	}
L696:
	;
	goto L695
L697:
	;
	v4814 = v4652 + v4656*int32(56)
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(v4700)))
	v4816 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4700)+4)))
	v4817 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+8))
	v4818 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+12))
	v4819 = *(*int64)(unsafe.Add(mBase, uint32(v4700)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v4814)+48)) = v4819
	*(*int32)(unsafe.Add(mBase, uint32(v4814)+12)) = v4818
	*(*int32)(unsafe.Add(mBase, uint32(v4814)+8)) = v4817
	v4823 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4814)+6)) = uint16(v4823)
	*(*uint16)(unsafe.Add(mBase, uint32(v4814)+4)) = uint16(v4816)
	*(*int32)(unsafe.Add(mBase, uint32(v4814))) = v4815
	v4830 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[0]))
	F_fmgr_info_copy(m, v4814+int32(16), v4810, v4830)
	mBase = m.M
	goto L698
L698:
	;
	goto L691
L699:
	;
	if v4836 == int32(0) {
		goto L663
	} else {
		goto L700
	}
L700:
	;
	v4843 = *(*int32)(unsafe.Add(mBase, uint32(v4700)))
	v4844 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4700)+4)))
	v4846 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+8))
	v4847 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+12))
	v4848 = *(*int64)(unsafe.Add(mBase, uint32(v4700)+48))
	F_ScanKeyEntryInitialize(m, v4652+v4656*int32(56), v4843, v4844, int32(0), v4846, v4847, v4836, v4848)
	mBase = m.M
	v4850 = m.ExcPending
	if v4850 != 0 {
		goto L24
	} else {
		goto L701
	}
L701:
	;
	goto L691
L702:
	;
	goto L679
L703:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+268)) = v4883
	v4904 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+264)) = v4904
	*(*uint8)(unsafe.Add(mBase, uint32(v4157)+258)) = uint8(v4904)
	v4909 = v4638 & int32(_a_F__bt_first_0)
	switch v4909 - int32(1) {
	case 0:
		goto L674
	case 1:
		goto L672
	case 2:
		goto L671
	case 3:
		goto L669
	case 4:
		goto L667
	default:
		goto L666
	}
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+268)) = v4749
	v4919 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+264)) = v4919
	*(*uint8)(unsafe.Add(mBase, uint32(v4157)+258)) = uint8(v4919)
	goto L674
L705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+268)) = v4749
	v4973 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+264)) = v4973
	*(*uint8)(unsafe.Add(mBase, uint32(v4157)+258)) = uint8(v4973)
	goto L672
L706:
	;
	v5022 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v4157)+259)) = uint16(v5022)
	goto L664
L707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+268)) = v4749
	v5031 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+264)) = v5031
	*(*uint8)(unsafe.Add(mBase, uint32(v4157)+258)) = uint8(v5031)
	goto L669
L708:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+268)) = v4749
	v5085 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+264)) = v5085
	*(*uint8)(unsafe.Add(mBase, uint32(v4157)+258)) = uint8(v5085)
	goto L667
L709:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+16)) = v4909
	F_errmsg_internal(m, int32(_a_F__bt_first_18), v4157+int32(16))
	mBase = m.M
	v5141 = m.ExcPending
	if v5141 != 0 {
		goto L24
	} else {
		goto L710
	}
L710:
	;
	F_errfinish(m, int32(_a_F__bt_first_19), int32(1504), int32(_a_F__bt_first_20))
	mBase = m.M
	v5146 = m.ExcPending
	if v5146 != 0 {
		goto L24
	} else {
		goto L711
	}
L711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L712:
	;
	v5200 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+56))
	if v5200 == int32(0) {
		goto L713
	} else {
		goto L714
	}
L713:
	;
	v5204 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[1]))
	if v5204 != int32(3) {
		goto L576
	} else {
		goto L716
	}
L714:
	;
	v5218 = v5200
	goto L715
L715:
	;
	v5221 = F__bt_binsrch(m, v4165, v4157+int32(256), v5218)
	mBase = m.M
	v5222 = m.ExcPending
	if v5222 != 0 {
		goto L24
	} else {
		goto L720
	}
L716:
	;
	v5207 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+8))
	F_PredicateLockRelation(m, v4165, v5207)
	mBase = m.M
	v5209 = m.ExcPending
	if v5209 != 0 {
		goto L24
	} else {
		goto L717
	}
L717:
	;
	v5210 = int32(0)
	v5213 = F__bt_search(m, v4165, v5210, v5193, v5195, int32(1), v5210)
	mBase = m.M
	v5214 = m.ExcPending
	if v5214 != 0 {
		goto L24
	} else {
		goto L718
	}
L718:
	;
	v5215 = *(*int32)(unsafe.Add(mBase, uint32(v5195)))
	if v5215 == int32(0) {
		goto L576
	} else {
		goto L719
	}
L719:
	;
	v5218 = v5215
	goto L715
L720:
	;
	v5223 = F__bt_readfirstpage(m, v4155, v5221, v4145)
	mBase = m.M
	v5224 = m.ExcPending
	if v5224 != 0 {
		goto L24
	} else {
		goto L721
	}
L721:
	;
	if v5223 == int32(0) {
		v5475 = v5190
		goto L575
	} else {
		goto L722
	}
L722:
	;
	v5227 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+100))
	v5230 = v4166 + v5227*int32(10)
	v5231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5230)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4155)+64)) = uint16(v5231)
	v5234 = v5230 + int32(104)
	v5235 = *(*int32)(unsafe.Add(mBase, uint32(v5234)))
	*(*int32)(unsafe.Add(mBase, uint32(v4155)+60)) = v5235
	v5237 = int32(1)
	v5238 = *(*int32)(unsafe.Add(mBase, uint32(v4166)+44))
	if v5238 == int32(0) {
		v5475 = v5237
		goto L575
	} else {
		goto L723
	}
L723:
	;
	v5241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5234)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v4155)+44)) = v5238 + v5241
	v5475 = v5237
	goto L575
L724:
	;
	v5248 = *(*int32)(unsafe.Add(mBase, uint32(v4165)+212))
	v5252 = *(*int32)(unsafe.Add(mBase, uint32(v5248+v4656<<(uint(int32(2))%32))))
	v5253 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+8))
	v5254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4700)+4)))
	v5255 = *(*int32)(unsafe.Add(mBase, uint32(v4165)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+48)) = v5255 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+44)) = v5254
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+40)) = v5253
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+36)) = v5252
	*(*int32)(unsafe.Add(mBase, uint32(v4157)+32)) = int32(1)
	F_errmsg_internal(m, int32(_a_F__bt_first_21), v4157+int32(32))
	mBase = m.M
	v5268 = m.ExcPending
	if v5268 != 0 {
		goto L24
	} else {
		goto L725
	}
L725:
	;
	F_errfinish(m, int32(_a_F__bt_first_19), int32(1423), int32(_a_F__bt_first_20))
	mBase = m.M
	v5273 = m.ExcPending
	if v5273 != 0 {
		goto L24
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
	*(*int32)(unsafe.Add(mBase, uint32(v5315)+56)) = v5320
	if v5320 == int32(0) {
		goto L728
	} else {
		goto L729
	}
L728:
	;
	v5326 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[1]))
	if v5326 != int32(3) {
		goto L576
	} else {
		goto L731
	}
L729:
	;
	v5338 = v5320
	goto L730
L730:
	;
	if v5338 < int32(0) {
		goto L736
	} else {
		goto L737
	}
L731:
	;
	v5329 = *(*int32)(unsafe.Add(mBase, uint32(v4155)+8))
	F_PredicateLockRelation(m, v5316, v5329)
	mBase = m.M
	v5331 = m.ExcPending
	if v5331 != 0 {
		goto L24
	} else {
		goto L732
	}
L732:
	;
	v5333 = F__bt_get_endpoint(m, v5316, int32(0), v5319)
	mBase = m.M
	v5334 = m.ExcPending
	if v5334 != 0 {
		goto L24
	} else {
		goto L733
	}
L733:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5315)+56)) = v5333
	if v5333 == int32(0) {
		goto L576
	} else {
		goto L734
	}
L734:
	;
	v5338 = v5333
	goto L730
L735:
	;
	if v4145 == int32(1) {
		goto L741
	} else {
		goto L742
	}
L736:
	;
	v5342 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[2]))
	v5348 = *(*int32)(unsafe.Add(mBase, uint32(v5342+(v5338^int32(-1))<<(uint(int32(2))%32))))
	v5356 = v5348
	goto L735
L737:
	;
	goto L738
L738:
	;
	v5350 = *(*int32)(unsafe.Add(mBase, _c_F__bt_first[3]))
	v5356 = v5350 + v5338<<(uint(int32(13))%32) + int32(-8192)
	goto L735
L739:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5405 = m.ExcPending
	if v5405 != 0 {
		goto L24
	} else {
		goto L756
	}
L740:
	;
	v5380 = F__bt_readfirstpage(m, v4155, v5377&int32(_a_F__bt_first_0), v4145)
	mBase = m.M
	v5381 = m.ExcPending
	if v5381 != 0 {
		goto L24
	} else {
		goto L751
	}
L741:
	;
	v5361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5356)+16)))
	v5363 = *(*int32)(unsafe.Add(mBase, uint32(v5356+v5361)+4))
	if v5363 != 0 {
		goto L744
	} else {
		goto L745
	}
L742:
	;
	goto L743
L743:
	;
	if v4145 != int32(-1) {
		goto L739
	} else {
		goto L747
	}
L744:
	;
	v5364 = int32(2)
	goto L746
L745:
	;
	v5364 = int32(1)
	goto L746
L746:
	;
	v5377 = v5364
	goto L740
L747:
	;
	v5367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5356)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v5367) {
		goto L748
	} else {
		goto L749
	}
L748:
	;
	v5375 = int32(base.Ui32(v5367+int32(_a_F__bt_first_22)) >> (uint(int32(2)) % 32))
	goto L750
L749:
	;
	v5375 = int32(0)
	goto L750
L750:
	;
	v5377 = v5375
	goto L740
L751:
	;
	if v5380 == int32(0) {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	v5475 = int32(0)
	goto L575
L753:
	;
	goto L754
L754:
	;
	v5385 = *(*int32)(unsafe.Add(mBase, uint32(v5315)+100))
	v5388 = v5315 + v5385*int32(10)
	v5389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5388)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4155)+64)) = uint16(v5389)
	v5392 = v5388 + int32(104)
	v5393 = *(*int32)(unsafe.Add(mBase, uint32(v5392)))
	*(*int32)(unsafe.Add(mBase, uint32(v4155)+60)) = v5393
	v5395 = int32(1)
	v5396 = *(*int32)(unsafe.Add(mBase, uint32(v5315)+44))
	if v5396 == int32(0) {
		v5475 = v5395
		goto L575
	} else {
		goto L755
	}
L755:
	;
	v5399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5392)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v4155)+44)) = v5396 + v5399
	v5475 = v5395
	goto L575
L756:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4157))) = v4145
	F_errmsg_internal(m, int32(_a_F__bt_first_23), v4157)
	mBase = m.M
	v5409 = m.ExcPending
	if v5409 != 0 {
		goto L24
	} else {
		goto L757
	}
L757:
	;
	F_errfinish(m, int32(_a_F__bt_first_19), int32(2234), int32(_a_F__bt_first_24))
	mBase = m.M
	v5414 = m.ExcPending
	if v5414 != 0 {
		goto L24
	} else {
		goto L758
	}
L758:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L759:
	;
	v5475 = int32(0)
	goto L575
}
