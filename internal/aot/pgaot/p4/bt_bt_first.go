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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v283 int32
	_ = v283
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v550 int32
	_ = v550
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v878 int32
	_ = v878
	var v899 int64
	_ = v899
	var v901 int64
	_ = v901
	var v903 int64
	_ = v903
	var v905 int64
	_ = v905
	var v907 int64
	_ = v907
	var v909 int64
	_ = v909
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1124 int32
	_ = v1124
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1315 int32
	_ = v1315
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1398 int32
	_ = v1398
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1475 int32
	_ = v1475
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1521 int32
	_ = v1521
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1542 int32
	_ = v1542
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1567 int32
	_ = v1567
	var v1578 int32
	_ = v1578
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1645 int32
	_ = v1645
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1697 int32
	_ = v1697
	var v1711 int32
	_ = v1711
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1819 int64
	_ = v1819
	var v1821 int64
	_ = v1821
	var v1823 int64
	_ = v1823
	var v1825 int64
	_ = v1825
	var v1827 int64
	_ = v1827
	var v1829 int64
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1845 int64
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2022 int32
	_ = v2022
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2107 int32
	_ = v2107
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2150 int32
	_ = v2150
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2173 int32
	_ = v2173
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2182 int32
	_ = v2182
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2197 int32
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2210 int32
	_ = v2210
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2223 int32
	_ = v2223
	var v2224 int64
	_ = v2224
	var v2226 int64
	_ = v2226
	var v2228 int64
	_ = v2228
	var v2230 int64
	_ = v2230
	var v2232 int64
	_ = v2232
	var v2234 int64
	_ = v2234
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2263 int32
	_ = v2263
	var v2266 int32
	_ = v2266
	var v2267 int64
	_ = v2267
	var v2269 int64
	_ = v2269
	var v2271 int64
	_ = v2271
	var v2273 int64
	_ = v2273
	var v2275 int64
	_ = v2275
	var v2277 int64
	_ = v2277
	var v2282 int32
	_ = v2282
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int64
	_ = v2305
	var v2307 int64
	_ = v2307
	var v2309 int64
	_ = v2309
	var v2311 int64
	_ = v2311
	var v2313 int64
	_ = v2313
	var v2315 int64
	_ = v2315
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2336 int32
	_ = v2336
	var v2337 int64
	_ = v2337
	var v2339 int64
	_ = v2339
	var v2341 int64
	_ = v2341
	var v2343 int64
	_ = v2343
	var v2345 int64
	_ = v2345
	var v2347 int64
	_ = v2347
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2369 int64
	_ = v2369
	var v2371 int64
	_ = v2371
	var v2373 int64
	_ = v2373
	var v2375 int64
	_ = v2375
	var v2377 int64
	_ = v2377
	var v2379 int64
	_ = v2379
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2397 int64
	_ = v2397
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2423 int32
	_ = v2423
	var v2426 int32
	_ = v2426
	var v2432 int32
	_ = v2432
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2494 int32
	_ = v2494
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int64
	_ = v2499
	var v2501 int64
	_ = v2501
	var v2503 int64
	_ = v2503
	var v2505 int64
	_ = v2505
	var v2507 int64
	_ = v2507
	var v2509 int64
	_ = v2509
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2563 int32
	_ = v2563
	var v2570 int32
	_ = v2570
	var v2597 int32
	_ = v2597
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2604 int32
	_ = v2604
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2641 int64
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2645 int64
	_ = v2645
	var v2647 int64
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2660 int32
	_ = v2660
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2700 int32
	_ = v2700
	var v2703 int32
	_ = v2703
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2724 int32
	_ = v2724
	var v2731 int32
	_ = v2731
	var v2734 int32
	_ = v2734
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2747 int32
	_ = v2747
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2761 int32
	_ = v2761
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2776 int32
	_ = v2776
	var v2783 int32
	_ = v2783
	var v2787 int32
	_ = v2787
	var v2790 int32
	_ = v2790
	var v2796 int32
	_ = v2796
	var v2803 int32
	_ = v2803
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2824 int32
	_ = v2824
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2829 int32
	_ = v2829
	var v2831 int32
	_ = v2831
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2853 int32
	_ = v2853
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2881 int32
	_ = v2881
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2894 int32
	_ = v2894
	var v2897 int32
	_ = v2897
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2912 int32
	_ = v2912
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2936 int32
	_ = v2936
	var v2939 int32
	_ = v2939
	var v2942 int32
	_ = v2942
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2965 int32
	_ = v2965
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2973 int32
	_ = v2973
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v3012 int32
	_ = v3012
	var v3022 int32
	_ = v3022
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3102 int32
	_ = v3102
	var v3105 int32
	_ = v3105
	var v3111 int32
	_ = v3111
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3121 int32
	_ = v3121
	var v3126 int32
	_ = v3126
	var v3218 int32
	_ = v3218
	var v3221 int32
	_ = v3221
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3245 int32
	_ = v3245
	var v3248 int32
	_ = v3248
	var v3250 int32
	_ = v3250
	var v3253 int32
	_ = v3253
	var v3259 int32
	_ = v3259
	var v3261 int32
	_ = v3261
	var v3265 int32
	_ = v3265
	var v3285 int32
	_ = v3285
	var v3287 int32
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3295 int32
	_ = v3295
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3313 int32
	_ = v3313
	var v3321 int32
	_ = v3321
	var v3329 int32
	_ = v3329
	var v3332 int32
	_ = v3332
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3360 int32
	_ = v3360
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3368 int32
	_ = v3368
	var v3377 int32
	_ = v3377
	var v3388 int32
	_ = v3388
	var v3422 int32
	_ = v3422
	var v3433 int32
	_ = v3433
	var v3457 int32
	_ = v3457
	var v3458 int32
	_ = v3458
	var v3461 int32
	_ = v3461
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3470 int32
	_ = v3470
	var v3474 int32
	_ = v3474
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3479 int32
	_ = v3479
	var v3483 int32
	_ = v3483
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3488 int32
	_ = v3488
	var v3492 int32
	_ = v3492
	var v3494 int32
	_ = v3494
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3507 int32
	_ = v3507
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3521 int32
	_ = v3521
	var v3523 int32
	_ = v3523
	var v3528 int32
	_ = v3528
	var v3550 int32
	_ = v3550
	var v3576 int32
	_ = v3576
	var v3579 int32
	_ = v3579
	var v3581 int32
	_ = v3581
	var v3584 int32
	_ = v3584
	var v3592 int32
	_ = v3592
	var v3596 int32
	_ = v3596
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3638 int32
	_ = v3638
	var v3664 int32
	_ = v3664
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3670 int32
	_ = v3670
	var v3671 int32
	_ = v3671
	var v3673 int32
	_ = v3673
	var v3676 int32
	_ = v3676
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3689 int32
	_ = v3689
	var v3691 int32
	_ = v3691
	var v3696 int32
	_ = v3696
	var v3699 int32
	_ = v3699
	var v3709 int32
	_ = v3709
	var v3736 int32
	_ = v3736
	var v3739 int32
	_ = v3739
	var v3741 int32
	_ = v3741
	var v3746 int32
	_ = v3746
	var v3747 int64
	_ = v3747
	var v3749 int64
	_ = v3749
	var v3751 int64
	_ = v3751
	var v3753 int64
	_ = v3753
	var v3755 int64
	_ = v3755
	var v3757 int64
	_ = v3757
	var v3759 int32
	_ = v3759
	var v3764 int32
	_ = v3764
	var v3766 int32
	_ = v3766
	var v3767 int32
	_ = v3767
	var v3770 int32
	_ = v3770
	var v3771 int64
	_ = v3771
	var v3773 int32
	_ = v3773
	var v3775 int64
	_ = v3775
	var v3777 int64
	_ = v3777
	var v3785 int32
	_ = v3785
	var v3786 int64
	_ = v3786
	var v3788 int64
	_ = v3788
	var v3790 int64
	_ = v3790
	var v3792 int64
	_ = v3792
	var v3794 int64
	_ = v3794
	var v3796 int64
	_ = v3796
	var v3798 int32
	_ = v3798
	var v3802 int32
	_ = v3802
	var v3806 int32
	_ = v3806
	var v3808 int32
	_ = v3808
	var v3809 int32
	_ = v3809
	var v3812 int32
	_ = v3812
	var v3813 int64
	_ = v3813
	var v3815 int32
	_ = v3815
	var v3817 int64
	_ = v3817
	var v3819 int64
	_ = v3819
	var v3823 int32
	_ = v3823
	var v3829 int32
	_ = v3829
	var v3837 int32
	_ = v3837
	var v3872 int32
	_ = v3872
	var v3928 int32
	_ = v3928
	var v3931 int32
	_ = v3931
	var v3969 int32
	_ = v3969
	var v3970 int32
	_ = v3970
	var v3974 int32
	_ = v3974
	var v3977 int32
	_ = v3977
	var v4014 int32
	_ = v4014
	var v4016 int32
	_ = v4016
	var v4017 int32
	_ = v4017
	var v4019 int32
	_ = v4019
	var v4022 int32
	_ = v4022
	var v4023 int32
	_ = v4023
	var v4026 int32
	_ = v4026
	var v4028 int32
	_ = v4028
	var v4030 int32
	_ = v4030
	var v4031 int32
	_ = v4031
	var v4032 int32
	_ = v4032
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4037 int32
	_ = v4037
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4054 int32
	_ = v4054
	var v4089 int32
	_ = v4089
	var v4092 int32
	_ = v4092
	var v4093 int32
	_ = v4093
	var v4097 int32
	_ = v4097
	var v4100 int32
	_ = v4100
	var v4101 int32
	_ = v4101
	var v4121 int32
	_ = v4121
	var v4145 int32
	_ = v4145
	var v4149 int32
	_ = v4149
	var v4151 int32
	_ = v4151
	var v4153 int32
	_ = v4153
	var v4201 int32
	_ = v4201
	var v4205 int32
	_ = v4205
	var v4210 int32
	_ = v4210
	var v4215 int32
	_ = v4215
	var v4221 int32
	_ = v4221
	var v4268 int32
	_ = v4268
	var v4271 int32
	_ = v4271
	var v4279 int32
	_ = v4279
	var v4280 int32
	_ = v4280
	var v4281 int32
	_ = v4281
	var v4284 int32
	_ = v4284
	var v4286 int32
	_ = v4286
	var v4287 int32
	_ = v4287
	var v4290 int32
	_ = v4290
	var v4291 int32
	_ = v4291
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4300 int32
	_ = v4300
	var v4303 int32
	_ = v4303
	var v4304 int32
	_ = v4304
	var v4307 int32
	_ = v4307
	var v4308 int32
	_ = v4308
	var v4310 int32
	_ = v4310
	var v4313 int32
	_ = v4313
	var v4316 int32
	_ = v4316
	var v4319 int32
	_ = v4319
	var v4323 int32
	_ = v4323
	var v4324 int32
	_ = v4324
	var v4325 int32
	_ = v4325
	var v4326 int64
	_ = v4326
	var v4331 int32
	_ = v4331
	var v4332 int64
	_ = v4332
	var v4336 int32
	_ = v4336
	var v4342 int32
	_ = v4342
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4349 int32
	_ = v4349
	var v4352 int32
	_ = v4352
	var v4354 int32
	_ = v4354
	var v4359 int32
	_ = v4359
	var v4372 int32
	_ = v4372
	var v4376 int32
	_ = v4376
	var v4379 int32
	_ = v4379
	var v4385 int32
	_ = v4385
	var v4390 int32
	_ = v4390
	var v4392 int32
	_ = v4392
	var v4397 int32
	_ = v4397
	var v4400 int32
	_ = v4400
	var v4401 int32
	_ = v4401
	var v4402 int32
	_ = v4402
	var v4406 int32
	_ = v4406
	var v4448 int32
	_ = v4448
	var v4449 int32
	_ = v4449
	var v4452 int32
	_ = v4452
	var v4455 int32
	_ = v4455
	var v4456 int32
	_ = v4456
	var v4458 int32
	_ = v4458
	var v4492 int32
	_ = v4492
	var v4502 int32
	_ = v4502
	var v4505 int32
	_ = v4505
	var v4514 int32
	_ = v4514
	var v4519 int32
	_ = v4519
	var v4525 int32
	_ = v4525
	var v4530 int32
	_ = v4530
	var v4533 int32
	_ = v4533
	var v4561 int32
	_ = v4561
	var v4584 int32
	_ = v4584
	var v4585 int32
	_ = v4585
	var v4586 int32
	_ = v4586
	var v4591 int32
	_ = v4591
	var v4595 int32
	_ = v4595
	var v4598 int32
	_ = v4598
	var v4599 int32
	_ = v4599
	var v4601 int32
	_ = v4601
	var v4606 int32
	_ = v4606
	var v4610 int32
	_ = v4610
	var v4613 int32
	_ = v4613
	var v4615 int32
	_ = v4615
	var v4637 int32
	_ = v4637
	var v4640 int32
	_ = v4640
	var v4650 int32
	_ = v4650
	var v4651 int32
	_ = v4651
	var v4658 int32
	_ = v4658
	var v4661 int32
	_ = v4661
	var v4664 int32
	_ = v4664
	var v4666 int32
	_ = v4666
	var v4684 int32
	_ = v4684
	var v4688 int32
	_ = v4688
	var v4691 int32
	_ = v4691
	var v4712 int32
	_ = v4712
	var v4734 int32
	_ = v4734
	var v4752 int32
	_ = v4752
	var v4762 int32
	_ = v4762
	var v4797 int32
	_ = v4797
	var v4801 int32
	_ = v4801
	var v4802 int32
	_ = v4802
	var v4807 int32
	_ = v4807
	var v4808 int32
	_ = v4808
	var v4809 int64
	_ = v4809
	var v4811 int64
	_ = v4811
	var v4813 int64
	_ = v4813
	var v4815 int64
	_ = v4815
	var v4817 int64
	_ = v4817
	var v4819 int64
	_ = v4819
	var v4823 int32
	_ = v4823
	var v4850 int32
	_ = v4850
	var v4864 int32
	_ = v4864
	var v4865 int32
	_ = v4865
	var v4872 int32
	_ = v4872
	var v4873 int64
	_ = v4873
	var v4875 int64
	_ = v4875
	var v4877 int64
	_ = v4877
	var v4879 int64
	_ = v4879
	var v4881 int64
	_ = v4881
	var v4883 int64
	_ = v4883
	var v4886 int32
	_ = v4886
	var v4887 int32
	_ = v4887
	var v4900 int32
	_ = v4900
	var v4901 int32
	_ = v4901
	var v4903 int32
	_ = v4903
	var v4906 int32
	_ = v4906
	var v4908 int32
	_ = v4908
	var v4909 int32
	_ = v4909
	var v4912 int32
	_ = v4912
	var v4913 int32
	_ = v4913
	var v4914 int32
	_ = v4914
	var v4915 int32
	_ = v4915
	var v4916 int32
	_ = v4916
	var v4917 int32
	_ = v4917
	var v4921 int32
	_ = v4921
	var v4928 int32
	_ = v4928
	var v4930 int32
	_ = v4930
	var v4932 int32
	_ = v4932
	var v4934 int32
	_ = v4934
	var v4935 int32
	_ = v4935
	var v4941 int32
	_ = v4941
	var v4942 int32
	_ = v4942
	var v4944 int32
	_ = v4944
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4948 int32
	_ = v4948
	var v4952 int32
	_ = v4952
	var v4983 int32
	_ = v4983
	var v4997 int32
	_ = v4997
	var v5003 int32
	_ = v5003
	var v5005 int32
	_ = v5005
	var v5010 int32
	_ = v5010
	var v5014 int32
	_ = v5014
	var v5020 int32
	_ = v5020
	var v5022 int32
	_ = v5022
	var v5068 int32
	_ = v5068
	var v5071 int32
	_ = v5071
	var v5077 int32
	_ = v5077
	var v5079 int32
	_ = v5079
	var v5125 int32
	_ = v5125
	var v5129 int32
	_ = v5129
	var v5132 int32
	_ = v5132
	var v5138 int32
	_ = v5138
	var v5140 int32
	_ = v5140
	var v5186 int32
	_ = v5186
	var v5189 int32
	_ = v5189
	var v5195 int32
	_ = v5195
	var v5197 int32
	_ = v5197
	var v5243 int32
	_ = v5243
	var v5248 int32
	_ = v5248
	var v5254 int32
	_ = v5254
	var v5259 int32
	_ = v5259
	var v5260 int32
	_ = v5260
	var v5308 int32
	_ = v5308
	var v5310 int32
	_ = v5310
	var v5311 int32
	_ = v5311
	var v5313 int32
	_ = v5313
	var v5314 int32
	_ = v5314
	var v5318 int32
	_ = v5318
	var v5321 int32
	_ = v5321
	var v5323 int32
	_ = v5323
	var v5328 int32
	_ = v5328
	var v5329 int32
	_ = v5329
	var v5331 int32
	_ = v5331
	var v5332 int32
	_ = v5332
	var v5335 int32
	_ = v5335
	var v5336 int32
	_ = v5336
	var v5340 int32
	_ = v5340
	var v5346 int32
	_ = v5346
	var v5348 int32
	_ = v5348
	var v5354 int32
	_ = v5354
	var v5355 int32
	_ = v5355
	var v5363 int32
	_ = v5363
	var v5368 int32
	_ = v5368
	var v5369 int32
	_ = v5369
	var v5370 int32
	_ = v5370
	var v5371 int32
	_ = v5371
	var v5373 int32
	_ = v5373
	var v5375 int32
	_ = v5375
	var v5380 int32
	_ = v5380
	var v5381 int32
	_ = v5381
	var v5425 int32
	_ = v5425
	var v5430 int32
	_ = v5430
	var v5431 int32
	_ = v5431
	var v5432 int32
	_ = v5432
	var v5433 int32
	_ = v5433
	var v5438 int32
	_ = v5438
	var v5442 int32
	_ = v5442
	var v5445 int32
	_ = v5445
	var v5451 int32
	_ = v5451
	var v5493 int32
	_ = v5493
	var v5494 int32
	_ = v5494
	var v5499 int32
	_ = v5499
	var v5502 int32
	_ = v5502
	var v5503 int32
	_ = v5503
	var v5506 int32
	_ = v5506
	var v5507 int32
	_ = v5507
	var v5509 int32
	_ = v5509
	var v5510 int32
	_ = v5510
	var v5513 int32
	_ = v5513
	var v5519 int32
	_ = v5519
	var v5520 int32
	_ = v5520
	var v5524 int32
	_ = v5524
	var v5525 int32
	_ = v5525
	var v5526 int32
	_ = v5526
	var v5527 int32
	_ = v5527
	var v5540 int32
	_ = v5540
	var v5545 int32
	_ = v5545
	var v5588 int32
	_ = v5588
	var v5589 int32
	_ = v5589
	var v5590 int32
	_ = v5590
	var v5594 int32
	_ = v5594
	var v5595 int32
	_ = v5595
	var v5599 int32
	_ = v5599
	var v5601 int32
	_ = v5601
	var v5603 int32
	_ = v5603
	var v5607 int32
	_ = v5607
	var v5613 int32
	_ = v5613
	var v5615 int32
	_ = v5615
	var v5621 int32
	_ = v5621
	var v5626 int32
	_ = v5626
	var v5628 int32
	_ = v5628
	var v5629 int32
	_ = v5629
	var v5632 int32
	_ = v5632
	var v5640 int32
	_ = v5640
	var v5642 int32
	_ = v5642
	var v5645 int32
	_ = v5645
	var v5646 int32
	_ = v5646
	var v5651 int32
	_ = v5651
	var v5654 int32
	_ = v5654
	var v5655 int32
	_ = v5655
	var v5658 int32
	_ = v5658
	var v5659 int32
	_ = v5659
	var v5661 int32
	_ = v5661
	var v5662 int32
	_ = v5662
	var v5665 int32
	_ = v5665
	var v5671 int32
	_ = v5671
	var v5675 int32
	_ = v5675
	var v5680 int32
	_ = v5680
	var v5724 int32
	_ = v5724
	var v5756 int32
	_ = v5756
	v3 = int32(0)
	v43 = m.G0
	v45 = v43 - int32(3280)
	m.G0 = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+60)) = int32(-1)
	v51 = m.G0
	v53 = v51 - int32(240)
	m.G0 = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v3 < v56 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v53 + int32(240)
	v4268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v4268 == int32(0) {
		goto L618
	} else {
		goto L619
	}
L2:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+224))
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v62
	v64 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v64)
	if v59 <= v62 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v69 <= int32(0) {
		v283 = v3
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v68)+224))
	v315 = int32(0)
	v317 = int32(1)
	v321 = v3
	v322 = v315
	v323 = v317
	v327 = v3
	v329 = v315
	v331 = v317
	v342 = v3
	goto L16
L5:
	;
	v73 = v69 & int32(3)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v69) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v83 = int32(0)
	v89 = v3
	v94 = v3
	goto L9
L7:
	;
	v178 = v3
	v183 = v3
	goto L8
L8:
	;
	if v73 == int32(0) {
		v283 = v183
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v122 = int32(48)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v74+v89*v122)))
	v126 = int32(5)
	v128 = int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v74+(v89|v128)*v122)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v74+(v89|int32(2))*v122)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v74+(v89|int32(3))*v122)))
	v163 = int32(base.Ui32(v125)>>(uint(v126)%32))&v128 + v94 + int32(base.Ui32(v136)>>(uint(v126)%32))&v128 + int32(base.Ui32(v147)>>(uint(v126)%32))&v128 + int32(base.Ui32(v158)>>(uint(v126)%32))&v128
	v164 = int32(4)
	v165 = v89 + v164
	v167 = v83 + v164
	if v167 != v69&int32(2147483644) {
		v83 = v167
		v89 = v165
		v94 = v163
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v178 = v165
	v183 = v163
	goto L8
L11:
	;
	goto L10
L12:
	;
	v217 = v3
	v222 = v178
	v227 = v183
	goto L13
L13:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v74+v222*int32(48))))
	v261 = int32(1)
	v263 = int32(base.Ui32(v258)>>(uint(int32(5))%32))&v261 + v227
	v267 = v217 + v261
	if v267 != v73 {
		v217 = v267
		v222 = v222 + v261
		v227 = v263
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v283 = v263
	goto L4
L15:
	;
	goto L14
L16:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v362 = base.I32_extend16_s(v323)
	v363 = base.I32_extend16_s(v331)
	if v363 <= v362 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v311)+16)) = uint8(base.B2i32(int32(0) < v550))
	v593 = v550 + v283
	if v593 != 0 {
		goto L41
	} else {
		goto L42
	}
L18:
	;
	goto L17
L19:
	;
	v443 = v323
	v444 = v362
	v481 = v322
	goto L21
L20:
	;
	v372 = v322
	v374 = v362
	goto L22
L21:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (base.B2i32(v329 == v482)|v342)&int32(1) != 0 {
		v550 = v481
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v414 = v374<<(uint(int32(2))%32) - int32(4)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v68)+208))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v418+v414)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v68)+212))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421+v414)))
	v425 = F_get_opfamily_member(m, v420, v423, v423, int32(3))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v443 = v331
	v444 = v433
	v481 = v435
	goto L21
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414+(v53+int32(80))))) = v425
	if v425 == int32(0) {
		v550 = v321
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v432 = int32(1)
	v433 = v374 + v432
	v435 = v372 + v432
	if (v322+v331-v323)&int32(65535) != v435&int32(65535) {
		v372 = v435
		v374 = v433
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v489 = v361 + v329*int32(48)
	v490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v489)+4)))
	if v490 <= v363 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v489)+6)))
	v321 = v481
	v322 = v527
	v323 = v528
	v327 = int32(base.Ui32(v532&int32(64))>>(uint(int32(6))%32)) | base.B2i32(v537 == int32(3)) | v530
	v329 = v329 + int32(1)
	v331 = v531
	v342 = int32(base.Ui32(v532&int32(4)) >> (uint(int32(2)) % 32))
	goto L16
L30:
	;
	v527 = v481
	v528 = v443
	v530 = v327
	v531 = v331
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v327&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v527 = v523
	v528 = v443 + int32(1)
	v530 = int32(0)
	v531 = v522
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53+int32(76)+v444<<(uint(int32(2))%32)))) = int32(0)
	v522 = v490
	v523 = v481
	goto L33
L35:
	;
	goto L36
L36:
	;
	v502 = v444<<(uint(int32(2))%32) - int32(4)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v68)+208))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v506+v502)))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v68)+212))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v509+v502)))
	v513 = F_get_opfamily_member(m, v508, v511, v511, int32(3))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502+(v53+int32(80))))) = v513
	if v513 == int32(0) {
		v550 = v481
		goto L18
	} else {
		goto L38
	}
L38:
	;
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v489)+4)))
	v522 = v518
	v523 = v481 + int32(1)
	goto L33
L39:
	;
	v1711 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1697)+4)))
	if int32(0) < v1711 {
		goto L176
	} else {
		goto L177
	}
L40:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1667 = int32(0)
	v1683 = v1667
	v1685 = v1667
	v1690 = v1645
	v1697 = v1666
	goto L39
L41:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v595 = v594 + v550
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v311)+28))
	if v596 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v1621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v1621 == int32(0) {
		goto L1
	} else {
		goto L173
	}
L44:
	;
	v612 = int32(4554240)
	v613 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v611
	v618 = F_palloc(m, v595*int32(48))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L24
	} else {
		goto L50
	}
L45:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v605 = F_AllocSetContextCreateInternal(m, v600, int32(65754), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L24
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_MemoryContextReset(m, v596)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L24
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+28)) = v605
	v611 = v605
	goto L44
L49:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v311)+28))
	v611 = v610
	goto L44
L50:
	;
	v622 = F_palloc(m, v593<<(uint(int32(5))%32))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+20)) = v622
	v627 = F_palloc(m, v595*int32(28))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L24
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+24)) = v627
	v630 = int32(0)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v632 <= v630 {
		v1567 = v630
		v1578 = v630
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+12)) = v1567
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v613
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v1602 == int32(0) {
		goto L1
	} else {
		goto L168
	}
L54:
	;
	v641 = v550
	v649 = v630
	v650 = int32(1)
	v657 = v3
	v659 = v3
	v660 = v630
	v661 = v3
	v670 = int32(-1)
	goto L55
L55:
	;
	v681 = int32(48)
	v683 = v618 + v660*v681
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v687 = v684 + v661*v681
	*(*int32)(unsafe.Add(mBase, uint32(v53)+48)) = v53 + int32(52)
	if v641 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v1567 = v1521
	v1578 = v1532
	goto L53
L57:
	;
	v1554 = v661 + int32(1)
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1554 < v1555 {
		v641 = v859
		v649 = v1521
		v650 = v868
		v657 = v1529
		v659 = v1531
		v660 = v1532
		v661 = v1554
		v670 = v1542
		goto L55
	} else {
		goto L167
	}
L58:
	;
	v1491 = v867 << (uint(int32(5)) % 32)
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1491+v1492))) = v878
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v53)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1495+v1491)+4)) = v1497
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1499+v1491)+8)) = v1501
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1503+v1491)+12)) = int32(-1)
	v1507 = int32(1)
	v1521 = v867 + v1507
	v1529 = v1484
	v1531 = v1485
	v1532 = v878 + v1507
	v1542 = v1486
	goto L57
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L24
	} else {
		goto L164
	}
L60:
	;
	v899 = *(*int64)(unsafe.Add(mBase, uint32(v687)))
	*(*int64)(unsafe.Add(mBase, uint32(v861))) = v899
	v901 = *(*int64)(unsafe.Add(mBase, uint32(v687)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+40)) = v901
	v903 = *(*int64)(unsafe.Add(mBase, uint32(v687)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+32)) = v903
	v905 = *(*int64)(unsafe.Add(mBase, uint32(v687)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+24)) = v905
	v907 = *(*int64)(unsafe.Add(mBase, uint32(v687)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+16)) = v907
	v909 = *(*int64)(unsafe.Add(mBase, uint32(v687)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+8)) = v909
	v911 = base.I32_wrap_i64(v899)
	if v911&int32(32) != 0 {
		goto L88
	} else {
		goto L89
	}
L61:
	;
	v859 = int32(0)
	v861 = v683
	v867 = v649
	v868 = v650
	v878 = v660
	goto L60
L62:
	;
	goto L63
L63:
	;
	v698 = v683
	v704 = v649
	v705 = v650
	v706 = v641
	v715 = v660
	goto L64
L64:
	;
	v736 = base.I32_extend16_s(v705)
	v737 = int32(*(*int16)(unsafe.Add(mBase, uint32(v687)+4)))
	if v737 < v736 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v859 = v764
	v861 = v854
	v867 = v849
	v868 = v847
	v878 = v851
	goto L60
L66:
	;
	v859 = v706
	v861 = v698
	v867 = v704
	v868 = v705
	v878 = v715
	goto L60
L67:
	;
	goto L68
L68:
	;
	v740 = v736 - int32(1)
	v742 = v740 << (uint(int32(2)) % 32)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v742+(v53+int32(80)))))
	if v746 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v859 = v706
	v861 = v698
	v867 = v704
	v868 = v705 + int32(1)
	v878 = v715
	goto L60
L70:
	;
	goto L71
L71:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v68)+248))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v751+v742)))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v68)+212))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v754+v742)))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v68)+208))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v757+v742)))
	v760 = F_get_opcode(m, v746)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	if v760 == int32(0) {
		goto L59
	} else {
		goto L73
	}
L73:
	;
	v764 = int32(0)
	F_ScanKeyEntryInitialize(m, v698, int32(262176), v736, int32(3), v764, v753, v760, v764)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L24
	} else {
		goto L74
	}
L74:
	;
	v772 = v704 << (uint(int32(5)) % 32)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v772+v773))) = v715
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v776+v772)+4)) = int32(-1)
	v780 = int32(1)
	v783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312+v740<<(uint(v780)%32)))))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v68)+52))
	v789 = v786 + v740<<(uint(int32(4))%32)
	v790 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v789)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v784+v772)+16)) = uint16(v790)
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v789)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v792+v772)+18)) = uint8(v794)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v796+v772)+19)) = uint8(v780)
	v804 = F_get_opfamily_proc(m, v759, v756, v756, int32(6))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L24
	} else {
		goto L76
	}
L75:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v828+v772)+20)) = v827
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	v833 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v831+v772)+24)) = v833
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v835+v772)+28)) = v833
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v311)+24))
	F__bt_setup_array_cmp(m, l0, v698, v756, v839+v715*int32(28), v833)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L24
	} else {
		goto L83
	}
L76:
	;
	if v804 == int32(0) {
		v827 = int32(0)
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v810 = F_palloc(m, int32(16))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L24
	} else {
		goto L78
	}
L78:
	;
	v812 = F_OidFunctionCall1Coll(m, v804, int32(0), v810)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	if v783&v780 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v810)))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v815
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v810)+12))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v810)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v810)+12)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v810)+8)) = v817
	*(*int32)(unsafe.Add(mBase, uint32(v810)+4)) = v814
	goto L82
L81:
	;
	goto L82
L82:
	;
	v827 = v810
	goto L75
L83:
	;
	v846 = int32(1)
	v847 = v705 + v846
	v849 = v704 + v846
	v851 = v715 + v846
	v854 = v618 + v851*int32(48)
	v856 = v706 - v846
	if v856 != 0 {
		v698 = v854
		v704 = v849
		v705 = v847
		v706 = v856
		v715 = v851
		goto L64
	} else {
		goto L84
	}
L84:
	;
	goto L65
L85:
	;
	v1464 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v1464)
	v1567 = v867
	v1578 = v878
	goto L53
L86:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v311)+24))
	F__bt_setup_array_cmp(m, l0, v861, v1022, v1096+v878*int32(28), v53+int32(48))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L24
	} else {
		goto L113
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L24
	} else {
		goto L110
	}
L88:
	;
	if v911&int32(1) != 0 {
		goto L85
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v1521 = v867
	v1529 = v657
	v1531 = v659
	v1532 = v878 + int32(1)
	v1542 = v670
	goto L57
L91:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v861)+44))
	v917 = F_pg_detoast_datum(m, v916)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L24
	} else {
		goto L92
	}
L92:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v917)+12))
	F_get_typlenbyvalalign(m, v919, v53+int32(46), v53+int32(45), v53+int32(44))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L24
	} else {
		goto L93
	}
L93:
	;
	v929 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+46)))
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+45)))
	v931 = int32(*(*int8)(unsafe.Add(mBase, uint32(v53)+44)))
	F_deconstruct_array(m, v917, v929, v930, v931, v53+int32(36), v53+int32(32), v53+int32(40))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v940 = int32(0)
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v53)+40))
	if v942 <= v940 {
		goto L85
	} else {
		goto L95
	}
L95:
	;
	v948 = v940
	v950 = v940
	v959 = v942
	goto L96
L96:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v53)+32))
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987+v950))))
	if v989 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v1004 == int32(0) {
		goto L85
	} else {
		goto L102
	}
L98:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	v993 = int32(2)
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v992+v950<<(uint(v993)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v992+v948<<(uint(v993)%32)))) = v999
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v53)+40))
	v1004 = v948 + int32(1)
	v1005 = v1001
	goto L100
L99:
	;
	v1004 = v948
	v1005 = v959
	goto L100
L100:
	;
	v1007 = v950 + int32(1)
	if v1007 < v1005 {
		v948 = v1004
		v950 = v1007
		v959 = v1005
		goto L96
	} else {
		goto L101
	}
L101:
	;
	goto L97
L102:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v861)+8))
	if v1011 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v68)+212))
	v1015 = int32(*(*int16)(unsafe.Add(mBase, uint32(v861)+4)))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1014+v1015<<(uint(int32(2))%32)-int32(4))))
	v1022 = v1021
	goto L105
L104:
	;
	v1022 = v1011
	goto L105
L105:
	;
	v1023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v861)+6)))
	switch v1023 - int32(1) {
	case 0, 1:
		goto L107
	case 2:
		goto L86
	case 3, 4:
		goto L106
	default:
		goto L87
	}
L106:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	v1035 = F__bt_find_extreme_element(m, l0, v861, v1022, int32(1), v1034, v1004)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L24
	} else {
		goto L109
	}
L107:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	v1028 = F__bt_find_extreme_element(m, l0, v861, v1022, int32(5), v1027, v1004)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L24
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v861)+44)) = v1028
	v1521 = v867
	v1529 = v657
	v1531 = v659
	v1532 = v878 + int32(1)
	v1542 = v670
	goto L57
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v861)+44)) = v1035
	goto L90
L110:
	;
	v1086 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v861)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v1086
	F_errmsg_internal(m, int32(504200), v53)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L24
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(515985), int32(2094), int32(120364))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L24
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v1104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v861)+4)))
	v1105 = int32(1)
	v1108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312-int32(2)+v1104<<(uint(v1105)%32)))))
	v1110 = v1108 & v1105
	if int32(2) <= v1004 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+212)) = v1114
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v861)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+220)) = uint8(v1110)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+216)) = v1116
	F_qsort_arg(m, v1113, v1004, int32(4), int32(210), v53+int32(212))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L24
	} else {
		goto L117
	}
L115:
	;
	v1209 = v1004
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+40)) = v1209
	v1249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v861)+4)))
	if v1249 != v659 {
		goto L131
	} else {
		goto L132
	}
L117:
	;
	v1132 = int32(1)
	v1135 = int32(0)
	goto L118
L118:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v53)+212))
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v53)+216))
	v1173 = int32(2)
	v1175 = v1113 + v1132<<(uint(v1173)%32)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1175)))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1113+v1135<<(uint(v1173)%32))))
	v1181 = F_FunctionCall2Coll(m, v1171, v1172, v1176, v1180)
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L24
	} else {
		goto L121
	}
L119:
	;
	v1209 = v1199 + int32(1)
	goto L116
L120:
	;
	v1202 = v1132 + int32(1)
	if v1202 != v1004 {
		v1132 = v1202
		v1135 = v1199
		goto L118
	} else {
		goto L130
	}
L121:
	;
	if v1181 < int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v1186 = int32(1)
	goto L124
L123:
	;
	v1186 = int32(0) - v1181
	goto L124
L124:
	;
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+220)))
	if v1187 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v1188 = v1186
	goto L127
L126:
	;
	v1188 = v1181
	goto L127
L127:
	;
	if v1188 == int32(0) {
		v1199 = v1135
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v1192 = v1135 + int32(1)
	if v1132 == v1192 {
		v1199 = v1132
		goto L120
	} else {
		goto L129
	}
L129:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1175)))
	*(*int32)(unsafe.Add(mBase, uint32(v1113+v1192<<(uint(int32(2))%32)))) = v1197
	v1199 = v1192
	goto L120
L130:
	;
	goto L119
L131:
	;
	v1484 = v1022
	v1485 = v1249
	v1486 = v867
	goto L58
L132:
	;
	goto L133
L133:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v311)+20))
	v1254 = v1251 + v670<<(uint(int32(5))%32)
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+4))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+8))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	if v1022 == v657 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v1284 = int32(0)
	if v1209 <= v1284 {
		v1398 = v1284
		goto L141
	} else {
		goto L142
	}
L135:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
	v1283 = v1259
	goto L134
L136:
	;
	goto L137
L137:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+208))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1262+v659<<(uint(int32(2))%32)-int32(4))))
	v1270 = F_get_opfamily_proc(m, v1268, v657, v1022, int32(1))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L24
	} else {
		goto L138
	}
L138:
	;
	if v1270 == int32(0) {
		v1484 = v657
		v1485 = v659
		v1486 = v670
		goto L58
	} else {
		goto L139
	}
L139:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+28))
	F_fmgr_info_cxt(m, v1270, v53+int32(212), v1276)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L24
	} else {
		goto L140
	}
L140:
	;
	v1283 = v53 + int32(212)
	goto L134
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+4)) = v1398
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	F_pfree(m, v1418)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L24
	} else {
		goto L162
	}
L142:
	;
	if v1255 <= int32(0) {
		v1398 = v1284
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v861)+12))
	v1290 = int32(0)
	v1296 = v1290
	v1297 = v1290
	v1315 = v1284
	goto L144
L144:
	;
	v1336 = int32(2)
	v1338 = v1256 + v1296<<(uint(v1336)%32)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1257+v1297<<(uint(v1336)%32))))
	v1344 = F_FunctionCall2Coll(m, v1283, v1289, v1339, v1343)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L24
	} else {
		goto L147
	}
L145:
	;
	v1398 = v1372
	goto L141
L146:
	;
	if v1255 <= v1370 {
		v1398 = v1372
		goto L141
	} else {
		goto L160
	}
L147:
	;
	if v1344 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1349 = int32(1)
	goto L150
L149:
	;
	v1349 = int32(0) - v1344
	goto L150
L150:
	;
	if v1110 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v1350 = v1349
	goto L153
L152:
	;
	v1350 = v1344
	goto L153
L153:
	;
	if v1350 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	*(*int32)(unsafe.Add(mBase, uint32(v1256+v1315<<(uint(int32(2))%32)))) = v1356
	v1358 = int32(1)
	v1370 = v1296 + v1358
	v1371 = v1297 + v1358
	v1372 = v1315 + v1358
	goto L146
L155:
	;
	goto L156
L156:
	;
	if v1350 < int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1370 = v1296 + int32(1)
	v1371 = v1297
	v1372 = v1315
	goto L146
L158:
	;
	goto L159
L159:
	;
	v1370 = v1296
	v1371 = v1297 + int32(1)
	v1372 = v1315
	goto L146
L160:
	;
	if v1371 < v1209 {
		v1296 = v1370
		v1297 = v1371
		v1315 = v1372
		goto L144
	} else {
		goto L161
	}
L161:
	;
	goto L145
L162:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+4))
	if v1421 != 0 {
		v1521 = v867
		v1529 = v657
		v1531 = v659
		v1532 = v878
		v1542 = v670
		goto L57
	} else {
		goto L163
	}
L163:
	;
	goto L85
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v746
	F_errmsg_internal(m, int32(46511), v53+int32(16))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L24
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(515985), int32(1957), int32(120364))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L24
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	goto L56
L168:
	;
	if v618 == int32(0) {
		v1645 = v1578
		goto L40
	} else {
		goto L169
	}
L169:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v1610 = F_MemoryContextAlloc(m, v1607, v1578<<(uint(int32(2))%32))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L24
	} else {
		goto L170
	}
L170:
	;
	v1612 = int32(1)
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1578 <= v1613 {
		v1683 = v1612
		v1685 = v1610
		v1690 = v1578
		v1697 = v618
		goto L39
	} else {
		goto L171
	}
L171:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v1618 = F_repalloc(m, v1615, v1578*int32(48))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L24
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v1618
	v1683 = v1612
	v1685 = v1610
	v1690 = v1578
	v1697 = v618
	goto L39
L173:
	;
	v1645 = v59
	goto L40
L174:
	;
	v4221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v4221)
	goto L1
L175:
	;
	v4215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v4215)
	goto L1
L176:
	;
	if v1690 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4201 = m.ExcPending
	if v4201 != 0 {
		goto L24
	} else {
		goto L614
	}
L179:
	;
	v1719 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1697)+4)))
	v1720 = int32(1)
	v1725 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+v1719<<(uint(v1720)%32)-int32(2)))))
	v1727 = v1725 << (uint(int32(24)) % 32)
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	if v1728&v1720 != 0 {
		goto L186
	} else {
		goto L187
	}
L180:
	;
	goto L181
L181:
	;
	v1840 = v53 + int32(136)
	v1841 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1840))) = v1841
	v1844 = v53 + int32(128)
	v1845 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1844))) = v1845
	v1848 = v53 + int32(120)
	*(*int64)(unsafe.Add(mBase, uint32(v1848))) = v1845
	v1852 = v53 + int32(112)
	*(*int64)(unsafe.Add(mBase, uint32(v1852))) = v1845
	v1856 = v53 + int32(104)
	*(*int64)(unsafe.Add(mBase, uint32(v1856))) = v1845
	v1860 = v53 + int32(96)
	*(*int64)(unsafe.Add(mBase, uint32(v1860))) = v1845
	*(*int64)(unsafe.Add(mBase, uint32(v53)+88)) = v1845
	*(*int64)(unsafe.Add(mBase, uint32(v53)+80)) = v1845
	v1881 = v1841
	v1882 = v1841
	v1884 = int32(1)
	v1885 = v1841
	v1896 = v1841
	v1897 = v1841
	goto L212
L182:
	;
	if v1813 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L183:
	;
	v1813 = int32(1)
	goto L182
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1697)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1697)+6)) = uint16(v1798)
	goto L183
L185:
	;
	if v1731&int32(33554432) != 0 {
		goto L204
	} else {
		goto L205
	}
L186:
	;
	v1731 = v1728 | v1727
	*(*int32)(unsafe.Add(mBase, uint32(v1697))) = v1731
	if v1728&int32(64) != 0 {
		v1798 = int32(3)
		goto L184
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	if v1725&int32(1) == int32(0) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	if v1728&int32(128) != 0 {
		goto L185
	} else {
		goto L190
	}
L190:
	;
	v1813 = int32(0)
	goto L182
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1697))) = v1728 | v1727
	if v1728&int32(4) == int32(0) {
		goto L183
	} else {
		goto L194
	}
L192:
	;
	if v1728&int32(16777216) != 0 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v1746 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1697)+6)))
	v1747 = int32(6) - v1746
	*(*uint16)(unsafe.Add(mBase, uint32(v1697)+6)) = uint16(v1747)
	goto L191
L194:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1697)+44))
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1755))))
	if v1756&int32(1) != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1813 = int32(0)
	goto L182
L196:
	;
	goto L197
L197:
	;
	v1762 = v1755
	goto L198
L198:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1762)))
	v1768 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1762)+4)))
	v1769 = int32(1)
	v1772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61-int32(2)+v1768<<(uint(v1769)%32)))))
	if v1772&v1769 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	goto L183
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1762))) = v1767 | v1772<<(uint(int32(24))%32)
	if v1767&int32(16) == int32(0) {
		v1762 = v1762 + int32(48)
		goto L198
	} else {
		goto L203
	}
L201:
	;
	if v1767&int32(16777216) != 0 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v1782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1762)+6)))
	v1783 = int32(6) - v1782
	*(*uint16)(unsafe.Add(mBase, uint32(v1762)+6)) = uint16(v1783)
	goto L200
L203:
	;
	goto L199
L204:
	;
	v1797 = int32(5)
	goto L206
L205:
	;
	v1797 = int32(1)
	goto L206
L206:
	;
	v1798 = v1797
	goto L184
L207:
	;
	v1816 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v1816)
	goto L209
L208:
	;
	goto L209
L209:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v1819 = *(*int64)(unsafe.Add(mBase, uint32(v1697)))
	*(*int64)(unsafe.Add(mBase, uint32(v1818))) = v1819
	v1821 = *(*int64)(unsafe.Add(mBase, uint32(v1697)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+40)) = v1821
	v1823 = *(*int64)(unsafe.Add(mBase, uint32(v1697)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+32)) = v1823
	v1825 = *(*int64)(unsafe.Add(mBase, uint32(v1697)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+24)) = v1825
	v1827 = *(*int64)(unsafe.Add(mBase, uint32(v1697)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+16)) = v1827
	v1829 = *(*int64)(unsafe.Add(mBase, uint32(v1697)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+8)) = v1829
	v1831 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v1831
	v1833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1697)+4)))
	if v1833 != v1831 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	F__bt_mark_scankey_required(m, v1836)
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L24
	} else {
		goto L211
	}
L211:
	;
	goto L1
L212:
	;
	v1917 = v1697 + v1897*int32(48)
	v1918 = base.B2i32(v1690 <= v1897)
	if v1690 <= v1897 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v2019 = base.B2i32(v1690 == v1897)
	if v2019 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L215:
	;
	v1922 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1917)+4)))
	v1923 = int32(1)
	v1928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+v1922<<(uint(v1923)%32)-int32(2)))))
	v1930 = v1928 << (uint(int32(24)) % 32)
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	if v1931&v1923 != 0 {
		goto L220
	} else {
		goto L221
	}
L216:
	;
	if v2016 != 0 {
		goto L214
	} else {
		goto L241
	}
L217:
	;
	v2016 = int32(1)
	goto L216
L218:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1917)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1917)+6)) = uint16(v2001)
	goto L217
L219:
	;
	if v1934&int32(33554432) != 0 {
		goto L238
	} else {
		goto L239
	}
L220:
	;
	v1934 = v1931 | v1930
	*(*int32)(unsafe.Add(mBase, uint32(v1917))) = v1934
	if v1931&int32(64) != 0 {
		v2001 = int32(3)
		goto L218
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	if v1928&int32(1) == int32(0) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	if v1931&int32(128) != 0 {
		goto L219
	} else {
		goto L224
	}
L224:
	;
	v2016 = int32(0)
	goto L216
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1917))) = v1931 | v1930
	if v1931&int32(4) == int32(0) {
		goto L217
	} else {
		goto L228
	}
L226:
	;
	if v1931&int32(16777216) != 0 {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v1949 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1917)+6)))
	v1950 = int32(6) - v1949
	*(*uint16)(unsafe.Add(mBase, uint32(v1917)+6)) = uint16(v1950)
	goto L225
L228:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+44))
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1958))))
	if v1959&int32(1) != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v2016 = int32(0)
	goto L216
L230:
	;
	goto L231
L231:
	;
	v1965 = v1958
	goto L232
L232:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1965)))
	v1971 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1965)+4)))
	v1972 = int32(1)
	v1975 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61-int32(2)+v1971<<(uint(v1972)%32)))))
	if v1975&v1972 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L217
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1965))) = v1970 | v1975<<(uint(int32(24))%32)
	if v1970&int32(16) == int32(0) {
		v1965 = v1965 + int32(48)
		goto L232
	} else {
		goto L237
	}
L235:
	;
	if v1970&int32(16777216) != 0 {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1985 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1965)+6)))
	v1986 = int32(6) - v1985
	*(*uint16)(unsafe.Add(mBase, uint32(v1965)+6)) = uint16(v1986)
	goto L234
L237:
	;
	goto L233
L238:
	;
	v2000 = int32(5)
	goto L240
L239:
	;
	v2000 = int32(1)
	goto L240
L240:
	;
	v2001 = v2000
	goto L218
L241:
	;
	v2017 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v2017)
	goto L1
L242:
	;
	v1881 = v2414
	v1882 = v2415
	v1884 = v2417
	v1885 = v2418
	v1896 = v2432
	v1897 = v1897 + int32(1)
	goto L212
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v2392
	if v1683 != 0 {
		goto L380
	} else {
		goto L381
	}
L244:
	;
	v2421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1917)+6)))
	v2423 = v2421 - int32(1)
	if v2421 == int32(3) {
		goto L355
	} else {
		goto L356
	}
L245:
	;
	v2022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1917)+4)))
	if v2022 == v1884&int32(65535) {
		v2414 = v1881
		v2415 = v1882
		v2417 = v1884
		v2418 = v1885
		goto L244
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	if v1918 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L248:
	;
	goto L247
L249:
	;
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v53)+104))
	if v2298 == int32(0) {
		v2328 = v2295
		goto L333
	} else {
		goto L334
	}
L250:
	;
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v2266 = v2263 + v2260*int32(48)
	v2267 = *(*int64)(unsafe.Add(mBase, uint32(v2259)))
	*(*int64)(unsafe.Add(mBase, uint32(v2266))) = v2267
	v2269 = *(*int64)(unsafe.Add(mBase, uint32(v2259)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2266)+40)) = v2269
	v2271 = *(*int64)(unsafe.Add(mBase, uint32(v2259)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2266)+32)) = v2271
	v2273 = *(*int64)(unsafe.Add(mBase, uint32(v2259)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2266)+24)) = v2273
	v2275 = *(*int64)(unsafe.Add(mBase, uint32(v2259)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2266)+16)) = v2275
	v2277 = *(*int64)(unsafe.Add(mBase, uint32(v2259)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2266)+8)) = v2277
	if v1683 != 0 {
		goto L326
	} else {
		goto L327
	}
L251:
	;
	if v2250 == int32(0) {
		v2292 = v2249
		v2293 = v2254
		v2295 = v2251
		goto L249
	} else {
		goto L325
	}
L252:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v2223 = v2220 + v1882*int32(48)
	v2224 = *(*int64)(unsafe.Add(mBase, uint32(v2168)))
	*(*int64)(unsafe.Add(mBase, uint32(v2223))) = v2224
	v2226 = *(*int64)(unsafe.Add(mBase, uint32(v2168)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2223)+40)) = v2226
	v2228 = *(*int64)(unsafe.Add(mBase, uint32(v2168)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2223)+32)) = v2228
	v2230 = *(*int64)(unsafe.Add(mBase, uint32(v2168)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2223)+24)) = v2230
	v2232 = *(*int64)(unsafe.Add(mBase, uint32(v2168)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2223)+16)) = v2232
	v2234 = *(*int64)(unsafe.Add(mBase, uint32(v2168)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2223)+8)) = v2234
	if v1683 != 0 {
		goto L320
	} else {
		goto L321
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+128)) = int32(0)
	v2257 = v2166
	v2258 = base.B2i32(v1881 == base.I32_extend16_s(v1884)-int32(1))
	v2259 = v2167
	v2260 = v1882
	goto L250
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L24
	} else {
		goto L317
	}
L255:
	;
	v2028 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1917)+4)))
	if v2028 < base.I32_extend16_s(v1884) {
		goto L254
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v53)+104))
	if v2031 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L257
L259:
	;
	if v2140 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L260:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v53)+92))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	v2139 = v2034
	v2140 = v2035
	v2141 = v1885
	v2144 = v1881
	goto L259
L261:
	;
	goto L262
L262:
	;
	v2036 = int32(0)
	if v1683 == v2036 {
		v2058 = v2036
		v2059 = v2036
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v53)+128))
	if v2060 != 0 {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v2040 = int32(0)
	v2041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031))))
	if v2041&int32(32) == v2040 {
		v2058 = v2036
		v2059 = v2040
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v53)+112))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v53)+108))
	v2058 = v2046 + v2047<<(uint(int32(5))%32) - int32(32)
	v2059 = v2053 + v2054*int32(28)
	goto L263
L266:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v2078 != 0 {
		goto L275
	} else {
		goto L276
	}
L267:
	;
	v2061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031))))
	if v2061&int32(64) != 0 {
		goto L174
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v2077 = v1885
	goto L266
L270:
	;
	v2067 = F__bt_compare_scankey_args(m, l0, v2060, v2031, v2060, v2058, v2059, v53+int32(212))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L24
	} else {
		goto L271
	}
L271:
	;
	if v2067 == int32(0) {
		v2077 = int32(1)
		goto L266
	} else {
		goto L272
	}
L272:
	;
	v2071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+212)))
	if v2071 == int32(0) {
		goto L175
	} else {
		goto L273
	}
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+128)) = int64(-4294967296)
	goto L269
L274:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v53)+92))
	if v2096 != 0 {
		goto L283
	} else {
		goto L284
	}
L275:
	;
	v2079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031))))
	if v2079&int32(64) != 0 {
		goto L174
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v2095 = v2077
	goto L274
L278:
	;
	v2085 = F__bt_compare_scankey_args(m, l0, v2078, v2031, v2078, v2058, v2059, v53+int32(212))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L24
	} else {
		goto L279
	}
L279:
	;
	if v2085 == int32(0) {
		v2095 = int32(1)
		goto L274
	} else {
		goto L280
	}
L280:
	;
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+212)))
	if v2089 != int32(1) {
		goto L175
	} else {
		goto L281
	}
L281:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+116)) = int64(-4294967296)
	goto L277
L282:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	if v2116 != 0 {
		goto L291
	} else {
		goto L292
	}
L283:
	;
	v2097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031))))
	if v2097&int32(64) != 0 {
		goto L174
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v2114 = int32(0)
	v2115 = v2095
	goto L282
L286:
	;
	v2103 = F__bt_compare_scankey_args(m, l0, v2096, v2031, v2096, v2058, v2059, v53+int32(212))
	mBase = m.M
	v2104 = m.ExcPending
	if v2104 != 0 {
		goto L24
	} else {
		goto L287
	}
L287:
	;
	if v2103 == int32(0) {
		v2114 = v2096
		v2115 = int32(1)
		goto L282
	} else {
		goto L288
	}
L288:
	;
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+212)))
	if v2107 != int32(1) {
		goto L175
	} else {
		goto L289
	}
L289:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+92)) = int64(-4294967296)
	goto L285
L290:
	;
	v2139 = v2114
	v2140 = v2134
	v2141 = v2135
	v2144 = v1881 + int32(1)
	goto L259
L291:
	;
	v2117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031))))
	if v2117&int32(64) != 0 {
		goto L174
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v2134 = int32(0)
	v2135 = v2115
	goto L290
L294:
	;
	v2123 = F__bt_compare_scankey_args(m, l0, v2116, v2031, v2116, v2058, v2059, v53+int32(212))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L24
	} else {
		goto L295
	}
L295:
	;
	if v2123 == int32(0) {
		v2134 = v2116
		v2135 = int32(1)
		goto L290
	} else {
		goto L296
	}
L296:
	;
	v2127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+212)))
	if v2127 != int32(1) {
		goto L175
	} else {
		goto L297
	}
L297:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+80)) = int64(-4294967296)
	goto L293
L298:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v53)+128))
	if v2168 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L299:
	;
	v2166 = v2141
	goto L298
L300:
	;
	if v2139 == int32(0) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v2150 = int32(0)
	v2154 = F__bt_compare_scankey_args(m, l0, v2139, v2140, v2139, v2150, v2150, v53+int32(212))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L24
	} else {
		goto L302
	}
L302:
	;
	if v2154 == int32(0) {
		v2166 = int32(1)
		goto L298
	} else {
		goto L303
	}
L303:
	;
	v2158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+212)))
	if v2158 == int32(1) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+92)) = int32(0)
	goto L299
L305:
	;
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+80)) = int32(0)
	goto L299
L307:
	;
	v2197 = base.B2i32(v1881 == base.I32_extend16_s(v1884)-int32(1))
	if v2168 != 0 {
		v2217 = v2166
		v2218 = v2167
		v2219 = v2197
		goto L252
	} else {
		goto L316
	}
L308:
	;
	if v2167 == int32(0) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v2173 = int32(0)
	v2177 = F__bt_compare_scankey_args(m, l0, v2167, v2168, v2167, v2173, v2173, v53+int32(212))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L24
	} else {
		goto L311
	}
L310:
	;
	v2217 = v2188
	v2218 = v2189
	v2219 = base.B2i32(v1881 == base.I32_extend16_s(v1884)-int32(1))
	goto L252
L311:
	;
	if v2177 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v2188 = int32(1)
	v2189 = v2167
	goto L310
L313:
	;
	goto L314
L314:
	;
	v2182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+212)))
	if v2182 != int32(1) {
		goto L253
	} else {
		goto L315
	}
L315:
	;
	v2185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+116)) = v2185
	v2188 = v2166
	v2189 = v2185
	goto L310
L316:
	;
	v2249 = v2166
	v2250 = v2167
	v2251 = v1882
	v2254 = v2197
	goto L251
L317:
	;
	F_errmsg_internal(m, int32(365738), int32(0))
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L24
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(515985), int32(343), int32(120480))
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L24
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v53)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v1882<<(uint(int32(2))%32)))) = v2239
	goto L322
L321:
	;
	goto L322
L322:
	;
	v2242 = v1882 + int32(1)
	v2243 = int32(0)
	if v2219 == v2243 {
		v2249 = v2217
		v2250 = v2218
		v2251 = v2242
		v2254 = v2243
		goto L251
	} else {
		goto L323
	}
L323:
	;
	F__bt_mark_scankey_required(m, v2223)
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L24
	} else {
		goto L324
	}
L324:
	;
	v2249 = v2217
	v2250 = v2218
	v2251 = v2242
	v2254 = int32(1)
	goto L251
L325:
	;
	v2257 = v2249
	v2258 = v2254
	v2259 = v2250
	v2260 = v2251
	goto L250
L326:
	;
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v53)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v2260<<(uint(int32(2))%32)))) = v2282
	goto L328
L327:
	;
	goto L328
L328:
	;
	v2285 = v2260 + int32(1)
	if v2258 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v2292 = v2257
	v2293 = int32(0)
	v2295 = v2285
	goto L249
L330:
	;
	goto L331
L331:
	;
	F__bt_mark_scankey_required(m, v2266)
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L24
	} else {
		goto L332
	}
L332:
	;
	v2292 = v2257
	v2293 = int32(1)
	v2295 = v2285
	goto L249
L333:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v53)+92))
	if v2330 == int32(0) {
		v2360 = v2328
		goto L340
	} else {
		goto L341
	}
L334:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v2304 = v2301 + v2295*int32(48)
	v2305 = *(*int64)(unsafe.Add(mBase, uint32(v2298)))
	*(*int64)(unsafe.Add(mBase, uint32(v2304))) = v2305
	v2307 = *(*int64)(unsafe.Add(mBase, uint32(v2298)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2304)+40)) = v2307
	v2309 = *(*int64)(unsafe.Add(mBase, uint32(v2298)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2304)+32)) = v2309
	v2311 = *(*int64)(unsafe.Add(mBase, uint32(v2298)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2304)+24)) = v2311
	v2313 = *(*int64)(unsafe.Add(mBase, uint32(v2298)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2304)+16)) = v2313
	v2315 = *(*int64)(unsafe.Add(mBase, uint32(v2298)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2304)+8)) = v2315
	if v1683 != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v53)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v2295<<(uint(int32(2))%32)))) = v2320
	goto L337
L336:
	;
	goto L337
L337:
	;
	v2323 = v2295 + int32(1)
	if v2293 == int32(0) {
		v2328 = v2323
		goto L333
	} else {
		goto L338
	}
L338:
	;
	F__bt_mark_scankey_required(m, v2304)
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L24
	} else {
		goto L339
	}
L339:
	;
	v2328 = v2323
	goto L333
L340:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	if v2362 == int32(0) {
		v2392 = v2360
		goto L347
	} else {
		goto L348
	}
L341:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v2336 = v2333 + v2328*int32(48)
	v2337 = *(*int64)(unsafe.Add(mBase, uint32(v2330)))
	*(*int64)(unsafe.Add(mBase, uint32(v2336))) = v2337
	v2339 = *(*int64)(unsafe.Add(mBase, uint32(v2330)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2336)+40)) = v2339
	v2341 = *(*int64)(unsafe.Add(mBase, uint32(v2330)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2336)+32)) = v2341
	v2343 = *(*int64)(unsafe.Add(mBase, uint32(v2330)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2336)+24)) = v2343
	v2345 = *(*int64)(unsafe.Add(mBase, uint32(v2330)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2336)+16)) = v2345
	v2347 = *(*int64)(unsafe.Add(mBase, uint32(v2330)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2336)+8)) = v2347
	if v1683 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v2328<<(uint(int32(2))%32)))) = v2352
	goto L344
L343:
	;
	goto L344
L344:
	;
	v2355 = v2328 + int32(1)
	if v2293 == int32(0) {
		v2360 = v2355
		goto L340
	} else {
		goto L345
	}
L345:
	;
	F__bt_mark_scankey_required(m, v2336)
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L24
	} else {
		goto L346
	}
L346:
	;
	v2360 = v2355
	goto L340
L347:
	;
	if v1690 == v1897 {
		goto L243
	} else {
		goto L354
	}
L348:
	;
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v2368 = v2365 + v2360*int32(48)
	v2369 = *(*int64)(unsafe.Add(mBase, uint32(v2362)))
	*(*int64)(unsafe.Add(mBase, uint32(v2368))) = v2369
	v2371 = *(*int64)(unsafe.Add(mBase, uint32(v2362)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2368)+40)) = v2371
	v2373 = *(*int64)(unsafe.Add(mBase, uint32(v2362)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2368)+32)) = v2373
	v2375 = *(*int64)(unsafe.Add(mBase, uint32(v2362)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2368)+24)) = v2375
	v2377 = *(*int64)(unsafe.Add(mBase, uint32(v2362)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2368)+16)) = v2377
	v2379 = *(*int64)(unsafe.Add(mBase, uint32(v2362)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2368)+8)) = v2379
	if v1683 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v53)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v2360<<(uint(int32(2))%32)))) = v2384
	goto L351
L350:
	;
	goto L351
L351:
	;
	v2387 = v2360 + int32(1)
	if v2293 == int32(0) {
		v2392 = v2387
		goto L347
	} else {
		goto L352
	}
L352:
	;
	F__bt_mark_scankey_required(m, v2368)
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L24
	} else {
		goto L353
	}
L353:
	;
	v2392 = v2387
	goto L347
L354:
	;
	v2394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1917)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1840))) = int32(0)
	v2397 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1844))) = v2397
	*(*int64)(unsafe.Add(mBase, uint32(v1848))) = v2397
	*(*int64)(unsafe.Add(mBase, uint32(v1852))) = v2397
	*(*int64)(unsafe.Add(mBase, uint32(v1856))) = v2397
	*(*int64)(unsafe.Add(mBase, uint32(v1860))) = v2397
	*(*int64)(unsafe.Add(mBase, uint32(v53)+88)) = v2397
	*(*int64)(unsafe.Add(mBase, uint32(v53)+80)) = v2397
	v2414 = v2144
	v2415 = v2392
	v2417 = v2394
	v2418 = v2292
	goto L244
L355:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	v2432 = int32(base.Ui32(v2426)>>(uint(int32(5))%32))&int32(1) + v1896
	goto L357
L356:
	;
	v2432 = v1896
	goto L357
L357:
	;
	v2437 = v53 + int32(80) + v2423*int32(12)
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2437)))
	if v2438 == int32(0) {
		v2527 = v2415
		v2529 = v2418
		goto L359
	} else {
		goto L360
	}
L358:
	;
	if v2423 != int32(2) {
		goto L242
	} else {
		goto L379
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2437)+8)) = v2432
	*(*int32)(unsafe.Add(mBase, uint32(v2437)+4)) = v1897
	*(*int32)(unsafe.Add(mBase, uint32(v2437))) = v1917
	v1881 = v2414
	v1882 = v2527
	v1884 = v2417
	v1885 = v2529
	v1896 = v2432
	v1897 = v1897 + int32(1)
	goto L212
L360:
	;
	v2441 = int32(0)
	v2444 = base.B2i32(v2423 != int32(2))
	if v2444|(v1683^int32(1)) != 0 {
		v2479 = v2441
		v2480 = v2441
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v2483 = F__bt_compare_scankey_args(m, l0, v1917, v1917, v2438, v2480, v2479, v53+int32(212))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L24
	} else {
		goto L367
	}
L362:
	;
	v2448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1917))))
	if v2448&int32(32) != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	v2479 = v2451 + v1897*int32(28)
	v2480 = v2455 + v2432<<(uint(int32(5))%32) - int32(32)
	goto L361
L364:
	;
	goto L365
L365:
	;
	v2461 = int32(0)
	v2462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2438))))
	if v2462&int32(32) == v2461 {
		v2479 = v2441
		v2480 = v2461
		goto L361
	} else {
		goto L366
	}
L366:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+8))
	v2479 = v2467 + v2468*int32(28)
	v2480 = v2472 + v2473<<(uint(int32(5))%32) - int32(32)
	goto L361
L367:
	;
	if v2483 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+212)))
	if v2485 != int32(1) {
		goto L358
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v2497 = v2494 + v2415*int32(48)
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v2437)))
	v2499 = *(*int64)(unsafe.Add(mBase, uint32(v2498)))
	*(*int64)(unsafe.Add(mBase, uint32(v2497))) = v2499
	v2501 = *(*int64)(unsafe.Add(mBase, uint32(v2498)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2497)+40)) = v2501
	v2503 = *(*int64)(unsafe.Add(mBase, uint32(v2498)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2497)+32)) = v2503
	v2505 = *(*int64)(unsafe.Add(mBase, uint32(v2498)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2497)+24)) = v2505
	v2507 = *(*int64)(unsafe.Add(mBase, uint32(v2498)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2497)+16)) = v2507
	v2509 = *(*int64)(unsafe.Add(mBase, uint32(v2498)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2497)+8)) = v2509
	if v1683 != 0 {
		goto L374
	} else {
		goto L375
	}
L371:
	;
	if v2423 != int32(2) {
		v2527 = v2415
		v2529 = v2418
		goto L359
	} else {
		goto L372
	}
L372:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2437)))
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2488))))
	if v2489&int32(32) == int32(0) {
		v2527 = v2415
		v2529 = v2418
		goto L359
	} else {
		goto L373
	}
L373:
	;
	goto L242
L374:
	;
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v2415<<(uint(int32(2))%32)))) = v2514
	goto L376
L375:
	;
	goto L376
L376:
	;
	v2516 = int32(1)
	v2517 = v2415 + v2516
	if v2414 != base.I32_extend16_s(v2417)-v2516 {
		v2527 = v2517
		v2529 = v2516
		goto L359
	} else {
		goto L377
	}
L377:
	;
	F__bt_mark_scankey_required(m, v2497)
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L24
	} else {
		goto L378
	}
L378:
	;
	v2527 = v2517
	v2529 = v2516
	goto L359
L379:
	;
	v2538 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v2538)
	goto L1
L380:
	;
	v2541 = int32(0)
	v2543 = m.G0
	v2545 = v2543 - int32(16)
	m.G0 = v2545
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+12))
	if v2548 == v2541 {
		goto L383
	} else {
		goto L384
	}
L381:
	;
	goto L382
L382:
	;
	if v2292&int32(1) == int32(0) {
		goto L1
	} else {
		goto L504
	}
L383:
	;
	m.G0 = v2545 + int32(16)
	goto L382
L384:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+4))
	if int32(0) < v2551 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2563 = v2541
	v2570 = v2541
	goto L388
L386:
	;
	goto L387
L387:
	;
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v3102 == int32(0) {
		goto L383
	} else {
		goto L498
	}
L388:
	;
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+8))
	v2600 = v2597 + v2570*int32(48)
	v2601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2600)+6)))
	if v2601 != int32(3) {
		v3022 = v2563
		goto L390
	} else {
		goto L391
	}
L389:
	;
	goto L387
L390:
	;
	v3057 = v2570 + int32(1)
	v3058 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+4))
	if v3057 < v3058 {
		v2563 = v3022
		v2570 = v3057
		goto L388
	} else {
		goto L497
	}
L391:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v2600)))
	if v2604&int32(32) == int32(0) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	if v2604&int32(65600) != int32(65536) {
		v3022 = v2563
		goto L390
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+24))
	v2631 = int32(28)
	v2633 = v2630 + v2570*v2631
	v2637 = *(*int32)(unsafe.Add(mBase, uint32(v1685+v2570<<(uint(int32(2))%32))))
	v2640 = v2630 + v2637*v2631
	v2641 = *(*int64)(unsafe.Add(mBase, uint32(v2640)))
	*(*int64)(unsafe.Add(mBase, uint32(v2633))) = v2641
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2633)+24)) = v2643
	v2645 = *(*int64)(unsafe.Add(mBase, uint32(v2640)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2633)+16)) = v2645
	v2647 = *(*int64)(unsafe.Add(mBase, uint32(v2640)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2633)+8)) = v2647
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+12))
	if v2649 <= v2563 {
		v3022 = v2563
		goto L390
	} else {
		goto L400
	}
L395:
	;
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v2600)+8))
	if v2613 != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2622 = v2613
	goto L398
L397:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v2554)+212))
	v2615 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2600)+4)))
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2614+v2615<<(uint(int32(2))%32)-int32(4))))
	v2622 = v2621
	goto L398
L398:
	;
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+24))
	F__bt_setup_array_cmp(m, l0, v2600, v2622, v2623+v2570*int32(28), int32(0))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L24
	} else {
		goto L399
	}
L399:
	;
	v3022 = v2563
	goto L390
L400:
	;
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+20))
	v2660 = v2563
	goto L401
L401:
	;
	v2696 = v2651 + v2660<<(uint(int32(5))%32)
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v2696)))
	if v2637 == v2697 {
		goto L403
	} else {
		goto L404
	}
L402:
	;
	v3022 = v2649
	goto L390
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2696))) = v2570
	v2700 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+4))
	switch v2700 + int32(1) {
	case 0:
		goto L407
	default:
		goto L406
	case 2:
		goto L408
	}
L404:
	;
	goto L405
L405:
	;
	v3012 = v2660 + int32(1)
	if v3012 != v2649 {
		v2660 = v3012
		goto L401
	} else {
		goto L496
	}
L406:
	;
	v3022 = v2660 + int32(1)
	goto L390
L407:
	;
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+20))
	if v2865 == int32(0) {
		goto L406
	} else {
		goto L456
	}
L408:
	;
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v2600)))
	*(*int32)(unsafe.Add(mBase, uint32(v2600))) = v2703 & int32(-33)
	v2707 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+8))
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v2707)))
	*(*int32)(unsafe.Add(mBase, uint32(v2600)+44)) = v2708
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+12))
	v2712 = v2710 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2547)+12)) = v2712
	if v2712 == int32(0) {
		goto L383
	} else {
		goto L409
	}
L409:
	;
	v2717 = v2696 + int32(32)
	v2720 = (v2712 - v2660) << (uint(int32(5)) % 32)
	if v2696 == v2717 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	v3022 = v2660
	goto L390
L411:
	;
	goto L410
L412:
	;
	v2724 = v2696 + v2720
	if base.Ui32(v2717-v2724) <= base.Ui32(int32(0)-v2720<<(uint(int32(1))%32)) {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v2731 = F___memcpy(m, v2696, v2717, v2720)
	mBase = m.M
	goto L410
L414:
	;
	goto L415
L415:
	;
	v2734 = (v2696 ^ v2717) & int32(3)
	if base.Ui32(v2696) < base.Ui32(v2717) {
		goto L418
	} else {
		goto L419
	}
L416:
	;
	if v2836 == int32(0) {
		goto L411
	} else {
		goto L452
	}
L417:
	;
	if base.Ui32(v2814) <= base.Ui32(int32(3)) {
		v2835 = v2813
		v2836 = v2814
		v2837 = v2815
		goto L416
	} else {
		goto L448
	}
L418:
	;
	if v2734 != 0 {
		goto L421
	} else {
		goto L422
	}
L419:
	;
	goto L420
L420:
	;
	if v2734 != 0 {
		v2796 = v2720
		goto L431
	} else {
		goto L432
	}
L421:
	;
	v2835 = v2717
	v2836 = v2720
	v2837 = v2696
	goto L416
L422:
	;
	goto L423
L423:
	;
	if v2696&int32(3) == int32(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v2813 = v2717
	v2814 = v2720
	v2815 = v2696
	goto L417
L425:
	;
	goto L426
L426:
	;
	v2741 = v2717
	v2742 = v2720
	v2743 = v2696
	goto L427
L427:
	;
	if v2742 == int32(0) {
		goto L411
	} else {
		goto L429
	}
L428:
	;
	v2813 = v2750
	v2814 = v2752
	v2815 = v2754
	goto L417
L429:
	;
	v2747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2741))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2743))) = uint8(v2747)
	v2749 = int32(1)
	v2750 = v2741 + v2749
	v2752 = v2742 - v2749
	v2754 = v2743 + v2749
	if v2754&int32(3) != 0 {
		v2741 = v2750
		v2742 = v2752
		v2743 = v2754
		goto L427
	} else {
		goto L430
	}
L430:
	;
	goto L428
L431:
	;
	if v2796 == int32(0) {
		goto L411
	} else {
		goto L444
	}
L432:
	;
	if v2724&int32(3) != 0 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2761 = v2720
	goto L436
L434:
	;
	v2776 = v2720
	goto L435
L435:
	;
	if base.Ui32(v2776) <= base.Ui32(int32(3)) {
		v2796 = v2776
		goto L431
	} else {
		goto L440
	}
L436:
	;
	if v2761 == int32(0) {
		goto L411
	} else {
		goto L438
	}
L437:
	;
	v2776 = v2767
	goto L435
L438:
	;
	v2767 = v2761 - int32(1)
	v2768 = v2696 + v2767
	v2770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2717+v2767))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2768))) = uint8(v2770)
	if v2768&int32(3) != 0 {
		v2761 = v2767
		goto L436
	} else {
		goto L439
	}
L439:
	;
	goto L437
L440:
	;
	v2783 = v2776
	goto L441
L441:
	;
	v2787 = v2783 - int32(4)
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v2717+v2787)))
	*(*int32)(unsafe.Add(mBase, uint32(v2696+v2787))) = v2790
	if base.Ui32(int32(3)) < base.Ui32(v2787) {
		v2783 = v2787
		goto L441
	} else {
		goto L443
	}
L442:
	;
	v2796 = v2787
	goto L431
L443:
	;
	goto L442
L444:
	;
	v2803 = v2796
	goto L445
L445:
	;
	v2807 = v2803 - int32(1)
	v2810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2717+v2807))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2696+v2807))) = uint8(v2810)
	if v2807 != 0 {
		v2803 = v2807
		goto L445
	} else {
		goto L447
	}
L446:
	;
	goto L411
L447:
	;
	goto L446
L448:
	;
	v2820 = v2813
	v2821 = v2814
	v2822 = v2815
	goto L449
L449:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v2820)))
	*(*int32)(unsafe.Add(mBase, uint32(v2822))) = v2824
	v2826 = int32(4)
	v2827 = v2820 + v2826
	v2829 = v2822 + v2826
	v2831 = v2821 - v2826
	if base.Ui32(int32(3)) < base.Ui32(v2831) {
		v2820 = v2827
		v2821 = v2831
		v2822 = v2829
		goto L449
	} else {
		goto L451
	}
L450:
	;
	v2835 = v2827
	v2836 = v2831
	v2837 = v2829
	goto L416
L451:
	;
	goto L450
L452:
	;
	v2842 = v2835
	v2843 = v2836
	v2844 = v2837
	goto L453
L453:
	;
	v2846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2842))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2844))) = uint8(v2846)
	v2848 = int32(1)
	v2853 = v2843 - v2848
	if v2853 != 0 {
		v2842 = v2842 + v2848
		v2843 = v2853
		v2844 = v2844 + v2848
		goto L453
	} else {
		goto L455
	}
L454:
	;
	goto L411
L455:
	;
	goto L454
L456:
	;
	v2868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2696)+19)))
	if v2868 != 0 {
		goto L406
	} else {
		goto L457
	}
L457:
	;
	v2869 = int32(4554240)
	v2870 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2873 = *(*int32)(unsafe.Add(mBase, uint32(v2872)+28))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v2873
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+28))
	if v2875 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+24))
	if v2936 == int32(0) {
		goto L477
	} else {
		goto L478
	}
L459:
	;
	v2878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2875)+6)))
	if v2878 != int32(1) {
		goto L458
	} else {
		goto L460
	}
L460:
	;
	v2881 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2600)+4)))
	v2885 = v2881<<(uint(int32(2))%32) - int32(4)
	v2886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(v2886)+208))
	v2889 = *(*int32)(unsafe.Add(mBase, uint32(v2885+v2887)))
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+44))
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+8))
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(v2886)+212))
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v2892+v2885)))
	if v2891 != 0 {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2897 = base.B2i32(v2891 != v2894)
	goto L463
L462:
	;
	v2897 = int32(0)
	goto L463
L463:
	;
	if v2897 != 0 {
		goto L458
	} else {
		goto L464
	}
L464:
	;
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+20))
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v2900)+8))
	v2902 = m.T0[v2901].(func(*base.Module, int32, int32, int32) int32)(m, v2886, v2890, v2545+int32(14))
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L24
	} else {
		goto L465
	}
L465:
	;
	v2904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+14)))
	if v2904 == int32(1) {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2908 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2907))) = uint8(v2908)
	goto L458
L467:
	;
	goto L468
L468:
	;
	v2912 = *(*int32)(unsafe.Add(mBase, uint32(v2875)))
	if v2912&int32(16777216) != 0 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v2915 = int32(4)
	goto L471
L470:
	;
	v2915 = int32(2)
	goto L471
L471:
	;
	v2916 = F_get_opfamily_member(m, v2889, v2894, v2894, v2915)
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L24
	} else {
		goto L472
	}
L472:
	;
	if v2916 == int32(0) {
		goto L458
	} else {
		goto L473
	}
L473:
	;
	v2920 = F_get_opcode(m, v2916)
	mBase = m.M
	v2921 = m.ExcPending
	if v2921 != 0 {
		goto L24
	} else {
		goto L474
	}
L474:
	;
	if v2920 == int32(0) {
		goto L458
	} else {
		goto L475
	}
L475:
	;
	F_fmgr_info(m, v2920, v2875+int32(16))
	mBase = m.M
	v2927 = m.ExcPending
	if v2927 != 0 {
		goto L24
	} else {
		goto L476
	}
L476:
	;
	v2928 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v2875)+6)) = uint16(v2928)
	*(*int32)(unsafe.Add(mBase, uint32(v2875)+44)) = v2902
	goto L458
L477:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v2870
	goto L406
L478:
	;
	v2939 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2936)+6)))
	if v2939 != int32(5) {
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v2942 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2600)+4)))
	v2946 = v2942<<(uint(int32(2))%32) - int32(4)
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2948 = *(*int32)(unsafe.Add(mBase, uint32(v2947)+208))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v2946+v2948)))
	v2951 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+44))
	v2952 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+8))
	v2953 = *(*int32)(unsafe.Add(mBase, uint32(v2947)+212))
	v2955 = *(*int32)(unsafe.Add(mBase, uint32(v2953+v2946)))
	if v2952 != 0 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v2958 = base.B2i32(v2952 != v2955)
	goto L482
L481:
	;
	v2958 = int32(0)
	goto L482
L482:
	;
	if v2958 != 0 {
		goto L477
	} else {
		goto L483
	}
L483:
	;
	v2961 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+20))
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(v2961)+12))
	v2963 = m.T0[v2962].(func(*base.Module, int32, int32, int32) int32)(m, v2947, v2951, v2545+int32(15))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L24
	} else {
		goto L484
	}
L484:
	;
	v2965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2545)+15)))
	if v2965 == int32(1) {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2969 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2968))) = uint8(v2969)
	goto L477
L486:
	;
	goto L487
L487:
	;
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v2936)))
	if v2973&int32(16777216) != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v2976 = int32(2)
	goto L490
L489:
	;
	v2976 = int32(4)
	goto L490
L490:
	;
	v2977 = F_get_opfamily_member(m, v2950, v2955, v2955, v2976)
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L24
	} else {
		goto L491
	}
L491:
	;
	if v2977 == int32(0) {
		goto L477
	} else {
		goto L492
	}
L492:
	;
	v2981 = F_get_opcode(m, v2977)
	mBase = m.M
	v2982 = m.ExcPending
	if v2982 != 0 {
		goto L24
	} else {
		goto L493
	}
L493:
	;
	if v2981 == int32(0) {
		goto L477
	} else {
		goto L494
	}
L494:
	;
	F_fmgr_info(m, v2981, v2936+int32(16))
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L24
	} else {
		goto L495
	}
L495:
	;
	v2989 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v2936)+6)) = uint16(v2989)
	*(*int32)(unsafe.Add(mBase, uint32(v2936)+44)) = v2963
	goto L477
L496:
	;
	goto L402
L497:
	;
	goto L389
L498:
	;
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+12))
	if v3105 < int32(33) {
		goto L383
	} else {
		goto L499
	}
L499:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L24
	} else {
		goto L500
	}
L500:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L24
	} else {
		goto L501
	}
L501:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2545)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v2545))) = v3115
	F_errmsg_internal(m, int32(709181), v2545)
	mBase = m.M
	v3121 = m.ExcPending
	if v3121 != 0 {
		goto L24
	} else {
		goto L502
	}
L502:
	;
	F_errfinish(m, int32(515985), int32(2373), int32(327721))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L24
	} else {
		goto L503
	}
L503:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L504:
	;
	v3218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v3218 != int32(1) {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	v3221 = int32(0)
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	v3230 = F_palloc0(m, v3229)
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L24
	} else {
		goto L506
	}
L506:
	;
	v3232 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	if int32(0) < v3232 {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+8))
	v3242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3241)+4)))
	v3245 = v3221
	v3248 = v3221
	v3250 = v3242
	v3253 = v3221
	v3259 = v3221
	v3261 = v3221
	v3265 = v3221
	goto L510
L508:
	;
	v3638 = v3221
	goto L509
L509:
	;
	v3664 = F_palloc(m, v3638*int32(48))
	mBase = m.M
	v3665 = m.ExcPending
	if v3665 != 0 {
		goto L24
	} else {
		goto L553
	}
L510:
	;
	v3285 = int32(0)
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+8))
	v3290 = v3287 + v3259*int32(48)
	v3291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3290)+4)))
	if v3291 == v3250&int32(65535) {
		goto L515
	} else {
		goto L516
	}
L511:
	;
	v3638 = v3592
	goto L509
L512:
	;
	v3617 = v3259 + int32(1)
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	if v3617 < v3618 {
		v3245 = v3576
		v3248 = v3579
		v3250 = v3581
		v3253 = v3584
		v3259 = v3617
		v3261 = v3592
		v3265 = v3596
		goto L510
	} else {
		goto L552
	}
L513:
	;
	v3576 = v3299
	v3579 = v3302
	v3581 = v3298
	v3584 = v3300
	v3592 = v3550
	v3596 = int32(1)
	goto L512
L514:
	;
	v3528 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3259+v3230))) = uint8(v3528)
	v3576 = v3516
	v3579 = v3518
	v3581 = v3519
	v3584 = v3521
	v3592 = v3261 + v3528
	v3596 = v3523
	goto L512
L515:
	;
	v3295 = int32(1)
	if v3265&v3295 != 0 {
		v3516 = v3245
		v3518 = v3248
		v3519 = v3250
		v3521 = v3253
		v3523 = v3295
		goto L514
	} else {
		goto L518
	}
L516:
	;
	v3298 = v3291
	v3299 = v3285
	v3300 = v3259
	v3302 = v3285
	goto L517
L517:
	;
	v3303 = *(*int32)(unsafe.Add(mBase, uint32(v3290)))
	v3304 = int32(196608)
	if v3303&v3304 == v3304 {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	v3298 = v3250
	v3299 = v3245
	v3300 = v3253
	v3302 = v3248
	goto L517
L519:
	;
	if v3259 <= v3300 {
		v3550 = v3261
		goto L513
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	v3496 = int32(1)
	v3497 = int32(0)
	if (base.B2i32(v3303&int32(65536) == v3497)|v3302)&v3496 == v3497 {
		goto L548
	} else {
		goto L549
	}
L522:
	;
	v3313 = (v3259 - v3300) & int32(3)
	if v3313 != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v3321 = v3300
	v3329 = int32(0)
	v3332 = v3261
	goto L526
L524:
	;
	v3377 = v3300
	v3388 = v3261
	goto L525
L525:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v3300-v3259) {
		v3550 = v3388
		goto L513
	} else {
		goto L532
	}
L526:
	;
	v3356 = v3321 + v3230
	v3357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3356))))
	if v3357 == int32(0) {
		goto L528
	} else {
		goto L529
	}
L527:
	;
	v3377 = v3366
	v3388 = v3364
	goto L525
L528:
	;
	v3360 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3356))) = uint8(v3360)
	v3364 = v3332 + v3360
	goto L530
L529:
	;
	v3364 = v3332
	goto L530
L530:
	;
	v3365 = int32(1)
	v3366 = v3321 + v3365
	v3368 = v3329 + v3365
	if v3368 != v3313 {
		v3321 = v3366
		v3329 = v3368
		v3332 = v3364
		goto L526
	} else {
		goto L531
	}
L531:
	;
	goto L527
L532:
	;
	v3422 = v3377
	v3433 = v3388
	goto L533
L533:
	;
	v3457 = v3422 + v3230
	v3458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3457))))
	if v3458 == int32(0) {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	v3550 = v3492
	goto L513
L535:
	;
	v3461 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3457))) = uint8(v3461)
	v3465 = v3433 + v3461
	goto L537
L536:
	;
	v3465 = v3433
	goto L537
L537:
	;
	v3466 = v3422 + (v3230 + int32(1))
	v3467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3466))))
	if v3467 == int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v3470 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3466))) = uint8(v3470)
	v3474 = v3465 + v3470
	goto L540
L539:
	;
	v3474 = v3465
	goto L540
L540:
	;
	v3475 = v3422 + (v3230 + int32(2))
	v3476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3475))))
	if v3476 == int32(0) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v3479 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3475))) = uint8(v3479)
	v3483 = v3474 + v3479
	goto L543
L542:
	;
	v3483 = v3474
	goto L543
L543:
	;
	v3484 = v3422 + (v3230 + int32(3))
	v3485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3484))))
	if v3485 == int32(0) {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v3488 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3484))) = uint8(v3488)
	v3492 = v3483 + v3488
	goto L546
L545:
	;
	v3492 = v3483
	goto L546
L546:
	;
	v3494 = v3422 + int32(4)
	if v3259 != v3494 {
		v3422 = v3494
		v3433 = v3492
		goto L533
	} else {
		goto L547
	}
L547:
	;
	goto L534
L548:
	;
	v3576 = v3299
	v3579 = v3496
	v3581 = v3298
	v3584 = v3300
	v3592 = v3261
	v3596 = v3497
	goto L512
L549:
	;
	goto L550
L550:
	;
	v3507 = int32(0)
	if (v3299|base.B2i32(v3303&int32(131072) == v3507))&int32(1) != 0 {
		v3516 = v3299
		v3518 = v3302
		v3519 = v3298
		v3521 = v3300
		v3523 = v3507
		goto L514
	} else {
		goto L551
	}
L551:
	;
	v3576 = int32(1)
	v3579 = v3302
	v3581 = v3298
	v3584 = v3300
	v3592 = v3261
	v3596 = v3497
	goto L512
L552:
	;
	goto L511
L553:
	;
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	v3670 = F_palloc(m, (v3666-v3638)*int32(48))
	mBase = m.M
	v3671 = m.ExcPending
	if v3671 != 0 {
		goto L24
	} else {
		goto L554
	}
L554:
	;
	v3673 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+12))
	if v3673 != 0 {
		goto L555
	} else {
		goto L556
	}
L555:
	;
	v3676 = F_palloc(m, v3638*int32(28))
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L24
	} else {
		goto L558
	}
L556:
	;
	v3684 = int32(0)
	v3685 = v3221
	goto L557
L557:
	;
	v3686 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	if v3686 <= int32(0) {
		goto L561
	} else {
		goto L562
	}
L558:
	;
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	v3682 = F_palloc(m, (v3678-v3638)*int32(28))
	mBase = m.M
	v3683 = m.ExcPending
	if v3683 != 0 {
		goto L24
	} else {
		goto L559
	}
L559:
	;
	v3684 = v3676
	v3685 = v3682
	goto L557
L560:
	;
	v4014 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+8))
	v4016 = v3977 * int32(48)
	if v4016 != 0 {
		goto L584
	} else {
		goto L585
	}
L561:
	;
	v3689 = int32(0)
	v3974 = v3689
	v3977 = v3689
	goto L560
L562:
	;
	goto L563
L563:
	;
	v3691 = int32(0)
	v3696 = v3691
	v3699 = v3691
	v3709 = v3691
	goto L564
L564:
	;
	v3736 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+8))
	v3739 = v3736 + v3709*int32(48)
	v3741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3709+v3230))))
	if v3741 == int32(0) {
		goto L567
	} else {
		goto L568
	}
L565:
	;
	v3974 = v3928
	v3977 = v3931
	goto L560
L566:
	;
	v3969 = v3709 + int32(1)
	v3970 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	if v3969 < v3970 {
		v3696 = v3928
		v3699 = v3931
		v3709 = v3969
		goto L564
	} else {
		goto L582
	}
L567:
	;
	v3746 = v3670 + v3699*int32(48)
	v3747 = *(*int64)(unsafe.Add(mBase, uint32(v3739)))
	*(*int64)(unsafe.Add(mBase, uint32(v3746))) = v3747
	v3749 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3746)+40)) = v3749
	v3751 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v3746)+32)) = v3751
	v3753 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3746)+24)) = v3753
	v3755 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3746)+16)) = v3755
	v3757 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3746)+8)) = v3757
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+12))
	if v3759 != 0 {
		goto L570
	} else {
		goto L571
	}
L568:
	;
	goto L569
L569:
	;
	v3785 = v3664 + v3696*int32(48)
	v3786 = *(*int64)(unsafe.Add(mBase, uint32(v3739)))
	*(*int64)(unsafe.Add(mBase, uint32(v3785))) = v3786
	v3788 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v3785)+40)) = v3788
	v3790 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v3785)+32)) = v3790
	v3792 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3785)+24)) = v3792
	v3794 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3785)+16)) = v3794
	v3796 = *(*int64)(unsafe.Add(mBase, uint32(v3739)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3785)+8)) = v3796
	v3798 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+12))
	if v3798 != 0 {
		goto L573
	} else {
		goto L574
	}
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v3709<<(uint(int32(2))%32)))) = v3699
	v3764 = int32(28)
	v3766 = v3685 + v3699*v3764
	v3767 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+24))
	v3770 = v3767 + v3709*v3764
	v3771 = *(*int64)(unsafe.Add(mBase, uint32(v3770)))
	*(*int64)(unsafe.Add(mBase, uint32(v3766))) = v3771
	v3773 = *(*int32)(unsafe.Add(mBase, uint32(v3770)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3766)+24)) = v3773
	v3775 = *(*int64)(unsafe.Add(mBase, uint32(v3770)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3766)+16)) = v3775
	v3777 = *(*int64)(unsafe.Add(mBase, uint32(v3770)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3766)+8)) = v3777
	goto L572
L571:
	;
	goto L572
L572:
	;
	v3928 = v3696
	v3931 = v3699 + int32(1)
	goto L566
L573:
	;
	v3802 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1685+v3709<<(uint(int32(2))%32)))) = v3802 + (v3696 - v3638)
	v3806 = int32(28)
	v3808 = v3684 + v3696*v3806
	v3809 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+24))
	v3812 = v3809 + v3709*v3806
	v3813 = *(*int64)(unsafe.Add(mBase, uint32(v3812)))
	*(*int64)(unsafe.Add(mBase, uint32(v3808))) = v3813
	v3815 = *(*int32)(unsafe.Add(mBase, uint32(v3812)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3808)+24)) = v3815
	v3817 = *(*int64)(unsafe.Add(mBase, uint32(v3812)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3808)+16)) = v3817
	v3819 = *(*int64)(unsafe.Add(mBase, uint32(v3812)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3808)+8)) = v3819
	goto L575
L574:
	;
	goto L575
L575:
	;
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v3785)))
	*(*int32)(unsafe.Add(mBase, uint32(v3785))) = v3823 & int32(-196609)
	if v3823&int32(4) != 0 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v3829 = *(*int32)(unsafe.Add(mBase, uint32(v3785)+44))
	v3837 = v3829
	goto L579
L577:
	;
	goto L578
L578:
	;
	v3928 = v3696 + int32(1)
	v3931 = v3699
	goto L566
L579:
	;
	v3872 = *(*int32)(unsafe.Add(mBase, uint32(v3837)))
	*(*int32)(unsafe.Add(mBase, uint32(v3837))) = v3872 & int32(-196609)
	if v3872&int32(16) == int32(0) {
		v3837 = v3837 + int32(48)
		goto L579
	} else {
		goto L581
	}
L580:
	;
	goto L578
L581:
	;
	goto L580
L582:
	;
	goto L565
L583:
	;
	v4019 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+8))
	v4022 = v3974 * int32(48)
	if v4022 != 0 {
		goto L588
	} else {
		goto L589
	}
L584:
	;
	v4017 = F__emscripten_memcpy_bulkmem(m, v4014, v3670, v4016)
	mBase = m.M
	goto L586
L585:
	;
	goto L586
L586:
	;
	goto L583
L587:
	;
	F_pfree(m, v3230)
	mBase = m.M
	v4026 = m.ExcPending
	if v4026 != 0 {
		goto L24
	} else {
		goto L591
	}
L588:
	;
	v4023 = F__emscripten_memcpy_bulkmem(m, v4019+v4016, v3664, v4022)
	mBase = m.M
	goto L590
L589:
	;
	goto L590
L590:
	;
	goto L587
L591:
	;
	F_pfree(m, v3670)
	mBase = m.M
	v4028 = m.ExcPending
	if v4028 != 0 {
		goto L24
	} else {
		goto L592
	}
L592:
	;
	F_pfree(m, v3664)
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L24
	} else {
		goto L593
	}
L593:
	;
	v4031 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+12))
	if v4031 != 0 {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v4032 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+24))
	v4034 = v3977 * int32(28)
	if v4034 != 0 {
		goto L598
	} else {
		goto L599
	}
L595:
	;
	goto L596
L596:
	;
	goto L1
L597:
	;
	v4037 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+24))
	v4040 = v3974 * int32(28)
	if v4040 != 0 {
		goto L602
	} else {
		goto L603
	}
L598:
	;
	v4035 = F__emscripten_memcpy_bulkmem(m, v4032, v3685, v4034)
	mBase = m.M
	goto L600
L599:
	;
	goto L600
L600:
	;
	goto L597
L601:
	;
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+12))
	if int32(0) < v4043 {
		goto L605
	} else {
		goto L606
	}
L602:
	;
	v4041 = F__emscripten_memcpy_bulkmem(m, v4037+v4034, v3684, v4040)
	mBase = m.M
	goto L604
L603:
	;
	goto L604
L604:
	;
	goto L601
L605:
	;
	v4054 = int32(0)
	goto L608
L606:
	;
	v4121 = v4043
	goto L607
L607:
	;
	v4145 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+20))
	F_pg_qsort(m, v4145, v4121, int32(32), int32(211))
	mBase = m.M
	v4149 = m.ExcPending
	if v4149 != 0 {
		goto L24
	} else {
		goto L611
	}
L608:
	;
	v4089 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+20))
	v4092 = v4089 + v4054<<(uint(int32(5))%32)
	v4093 = *(*int32)(unsafe.Add(mBase, uint32(v4092)))
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(v1685+v4093<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4092))) = v4097
	v4100 = v4054 + int32(1)
	v4101 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+12))
	if v4100 < v4101 {
		v4054 = v4100
		goto L608
	} else {
		goto L610
	}
L609:
	;
	v4121 = v4101
	goto L607
L610:
	;
	goto L609
L611:
	;
	F_pfree(m, v3684)
	mBase = m.M
	v4151 = m.ExcPending
	if v4151 != 0 {
		goto L24
	} else {
		goto L612
	}
L612:
	;
	F_pfree(m, v3685)
	mBase = m.M
	v4153 = m.ExcPending
	if v4153 != 0 {
		goto L24
	} else {
		goto L613
	}
L613:
	;
	goto L596
L614:
	;
	F_errmsg_internal(m, int32(365738), int32(0))
	mBase = m.M
	v4205 = m.ExcPending
	if v4205 != 0 {
		goto L24
	} else {
		goto L615
	}
L615:
	;
	F_errfinish(m, int32(515985), int32(267), int32(120480))
	mBase = m.M
	v4210 = m.ExcPending
	if v4210 != 0 {
		goto L24
	} else {
		goto L616
	}
L616:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L617:
	;
	m.G0 = v45 + int32(3280)
	return v5756
L618:
	;
	F__bt_parallel_done(m, l0)
	mBase = m.M
	v5724 = m.ExcPending
	if v5724 != 0 {
		goto L24
	} else {
		goto L823
	}
L619:
	;
	v4271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v4271 == int32(0) {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v4281 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v4281 == int32(0) {
		goto L624
	} else {
		goto L625
	}
L621:
	;
	v4279 = F__bt_parallel_seize(m, l0, v45+int32(60), v45+int32(56), int32(1))
	mBase = m.M
	v4280 = m.ExcPending
	if v4280 != 0 {
		goto L24
	} else {
		goto L622
	}
L622:
	;
	if v4279 != 0 {
		goto L620
	} else {
		goto L623
	}
L623:
	;
	v5756 = v3
	goto L617
L624:
	;
	v4287 = *(*int32)(unsafe.Add(mBase, uint32(v45)+60))
	if v4287 != int32(-1) {
		goto L628
	} else {
		goto L629
	}
L625:
	;
	v4284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+17)))
	if v4284 != 0 {
		goto L624
	} else {
		goto L626
	}
L626:
	;
	F__bt_start_array_keys(m, l0, l1)
	mBase = m.M
	v4286 = m.ExcPending
	if v4286 != 0 {
		goto L24
	} else {
		goto L627
	}
L627:
	;
	goto L624
L628:
	;
	v4290 = int32(1)
	v4291 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	v4293 = F__bt_readnextpage(m, l0, v4287, v4291, l1, v4290)
	mBase = m.M
	v4294 = m.ExcPending
	if v4294 != 0 {
		goto L24
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(v47)+272))
	if v4316 == int32(0) {
		goto L637
	} else {
		goto L638
	}
L631:
	;
	if v4293 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v5756 = int32(0)
	goto L617
L633:
	;
	goto L634
L634:
	;
	v4300 = *(*int32)(unsafe.Add(mBase, uint32(v48)+100))
	v4303 = v48 + v4300*int32(10)
	v4304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4303)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v4304)
	v4307 = v4303 + int32(104)
	v4308 = *(*int32)(unsafe.Add(mBase, uint32(v4307)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v4308
	v4310 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
	if v4310 == int32(0) {
		v5756 = v4290
		goto L617
	} else {
		goto L635
	}
L635:
	;
	v4313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4307)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v4310 + v4313
	v5756 = v4290
	goto L617
L636:
	;
	v4331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4331 != 0 {
		goto L642
	} else {
		goto L643
	}
L637:
	;
	v4319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+268)))
	if v4319 != int32(1) {
		goto L636
	} else {
		goto L640
	}
L638:
	;
	v4325 = v4316
	goto L639
L639:
	;
	v4326 = *(*int64)(unsafe.Add(mBase, uint32(v4325)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4325)+16)) = v4326 + int64(1)
	goto L636
L640:
	;
	F_pgstat_assoc_relation(m, v47)
	mBase = m.M
	v4323 = m.ExcPending
	if v4323 != 0 {
		goto L24
	} else {
		goto L641
	}
L641:
	;
	v4324 = *(*int32)(unsafe.Add(mBase, uint32(v47)+272))
	v4325 = v4324
	goto L639
L642:
	;
	v4332 = *(*int64)(unsafe.Add(mBase, uint32(v4331)))
	*(*int64)(unsafe.Add(mBase, uint32(v4331))) = v4332 + int64(1)
	goto L644
L643:
	;
	goto L644
L644:
	;
	v4336 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v4336 <= int32(0) {
		goto L645
	} else {
		goto L646
	}
L645:
	;
	v5588 = int32(0)
	v5589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5594 = F__bt_get_endpoint(m, v5590, v5588, base.B2i32(l1 == int32(-1)))
	mBase = m.M
	v5595 = m.ExcPending
	if v5595 != 0 {
		goto L24
	} else {
		goto L795
	}
L646:
	;
	v4342 = base.B2i32(l1 == int32(1))
	if l1 == int32(1) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v4343 = int32(24)
	goto L649
L648:
	;
	v4343 = int32(28)
	goto L649
L649:
	;
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v4349 = v4336
	v4352 = int32(1)
	v4354 = int32(3)
	v4359 = v4344
	v4372 = v3
	v4376 = v3
	v4379 = v3
	v4385 = v3
	goto L650
L650:
	;
	if v4385 < v4349 {
		goto L656
	} else {
		goto L657
	}
L651:
	;
	if v4734 == int32(0) {
		goto L645
	} else {
		goto L704
	}
L652:
	;
	goto L651
L653:
	;
	v4349 = v4661
	v4352 = v4664
	v4354 = v4666
	v4359 = v4359 + int32(48)
	v4372 = v4684
	v4376 = v4688
	v4379 = v4691
	v4385 = v4385 + int32(1)
	goto L650
L654:
	;
	v4650 = int32(0)
	v4651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4359)+6)))
	switch v4651 - int32(1) {
	case 0, 1:
		goto L698
	case 2:
		goto L697
	case 3, 4:
		goto L696
	default:
		v4661 = v4610
		v4664 = v4613
		v4666 = v4615
		v4684 = v4650
		v4688 = v4637
		v4691 = v4640
		goto L653
	}
L655:
	;
	if v4372 != 0 {
		v4661 = v4349
		v4664 = v4352
		v4666 = v4354
		v4684 = v4372
		v4688 = v4376
		v4691 = v4379
		goto L653
	} else {
		goto L694
	}
L656:
	;
	v4390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4359)+4)))
	if v4390 == v4352 {
		goto L655
	} else {
		goto L659
	}
L657:
	;
	goto L658
L658:
	;
	if v4372 != 0 {
		goto L662
	} else {
		goto L663
	}
L659:
	;
	goto L658
L660:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(1600)+v4376<<(uint(int32(2))%32)))) = v4561
	v4584 = int32(1)
	v4585 = v4376 + v4584
	v4586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4561)+6)))
	if v4586&int32(65531) == v4584 {
		v4712 = v4586
		v4734 = v4585
		goto L652
	} else {
		goto L685
	}
L661:
	;
	if v4533 == int32(0) {
		v4712 = v4354
		v4734 = v4376
		goto L652
	} else {
		goto L684
	}
L662:
	;
	v4392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4372)+2)))
	if v4392&int32(24) == int32(0) {
		v4561 = v4372
		goto L660
	} else {
		goto L665
	}
L663:
	;
	v4492 = v4379
	v4502 = int32(0)
	goto L664
L664:
	;
	if v4502 != 0 {
		v4533 = v4502
		goto L661
	} else {
		goto L675
	}
L665:
	;
	v4397 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v4400 = base.I32_div_s(v4372-v4397, int32(48))
	v4401 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v4402 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v4406 = int32(0)
	goto L666
L666:
	;
	v4448 = v4401 + v4406<<(uint(int32(5))%32)
	v4449 = *(*int32)(unsafe.Add(mBase, uint32(v4448)))
	if v4400 != v4449 {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	v4455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4448)+19)))
	if v4455 != 0 {
		goto L672
	} else {
		goto L673
	}
L668:
	;
	v4452 = v4406 + int32(1)
	if base.Ui32(v4452) < base.Ui32(v4402) {
		v4406 = v4452
		goto L666
	} else {
		goto L671
	}
L669:
	;
	goto L670
L670:
	;
	goto L667
L671:
	;
	goto L670
L672:
	;
	v4456 = v4379
	goto L674
L673:
	;
	v4456 = v4372
	goto L674
L674:
	;
	v4458 = *(*int32)(unsafe.Add(mBase, uint32(v4448+v4343)))
	v4492 = v4456
	v4502 = v4458
	goto L664
L675:
	;
	if v4492 == int32(0) {
		v4533 = v4502
		goto L661
	} else {
		goto L676
	}
L676:
	;
	v4505 = *(*int32)(unsafe.Add(mBase, uint32(v4492)))
	if v4505&int32(33554432) != 0 {
		goto L678
	} else {
		goto L679
	}
L677:
	;
	v4519 = v45 - int32(-64) + v4376*int32(48)
	v4525 = int32(0)
	F_ScanKeyEntryInitialize(m, v4519, v4505&int32(50331648)|int32(129), base.I32_extend16_s(v4352), v4514, v4525, v4525, v4525, v4525)
	mBase = m.M
	v4530 = m.ExcPending
	if v4530 != 0 {
		goto L24
	} else {
		goto L683
	}
L678:
	;
	if v4342 == int32(0) {
		v4712 = v4354
		v4734 = v4376
		goto L652
	} else {
		goto L681
	}
L679:
	;
	goto L680
L680:
	;
	if l1 != int32(-1) {
		v4712 = v4354
		v4734 = v4376
		goto L652
	} else {
		goto L682
	}
L681:
	;
	v4514 = int32(5)
	goto L677
L682:
	;
	v4514 = int32(1)
	goto L677
L683:
	;
	v4533 = v4519
	goto L661
L684:
	;
	v4561 = v4533
	goto L660
L685:
	;
	v4591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4561)+2)))
	if v4591&int32(96) != 0 {
		goto L686
	} else {
		goto L687
	}
L686:
	;
	v4595 = int32(1)
	if l1 == v4595 {
		goto L689
	} else {
		goto L690
	}
L687:
	;
	goto L688
L688:
	;
	v4599 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v4599 <= v4385 {
		v4712 = v4586
		v4734 = v4585
		goto L652
	} else {
		goto L692
	}
L689:
	;
	v4598 = int32(5)
	goto L691
L690:
	;
	v4598 = v4595
	goto L691
L691:
	;
	v4712 = v4598
	v4734 = v4585
	goto L652
L692:
	;
	v4601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4359)+2)))
	if v4601&int32(3) == int32(0) {
		v4712 = v4586
		v4734 = v4585
		goto L652
	} else {
		goto L693
	}
L693:
	;
	v4606 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4359)+4)))
	v4610 = v4599
	v4613 = v4606
	v4615 = v4586
	v4637 = v4585
	v4640 = int32(0)
	goto L654
L694:
	;
	v4610 = v4349
	v4613 = v4352
	v4615 = v4354
	v4637 = v4376
	v4640 = v4379
	goto L654
L695:
	;
	if v4640 != 0 {
		goto L701
	} else {
		goto L702
	}
L696:
	;
	if v4342 == int32(0) {
		goto L695
	} else {
		goto L700
	}
L697:
	;
	v4661 = v4610
	v4664 = v4613
	v4666 = v4615
	v4684 = v4359
	v4688 = v4637
	v4691 = v4640
	goto L653
L698:
	;
	if l1 != int32(-1) {
		goto L695
	} else {
		goto L699
	}
L699:
	;
	goto L697
L700:
	;
	v4661 = v4610
	v4664 = v4613
	v4666 = v4615
	v4684 = v4359
	v4688 = v4637
	v4691 = v4640
	goto L653
L701:
	;
	v4658 = v4640
	goto L703
L702:
	;
	v4658 = v4359
	goto L703
L703:
	;
	v4661 = v4610
	v4664 = v4613
	v4666 = v4615
	v4684 = v4650
	v4688 = v4637
	v4691 = v4658
	goto L653
L704:
	;
	if v4734 <= int32(0) {
		v4983 = v4734
		goto L718
	} else {
		goto L719
	}
L705:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5519 = m.ExcPending
	if v5519 != 0 {
		goto L24
	} else {
		goto L792
	}
L706:
	;
	v5308 = v48 + int32(56)
	v5310 = F__bt_search(m, v47, int32(0), v45+int32(1728), v5308, int32(1))
	mBase = m.M
	v5311 = m.ExcPending
	if v5311 != 0 {
		goto L24
	} else {
		goto L754
	}
L707:
	;
	v5260 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+1731)) = uint16(v5260)
	goto L706
L708:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5248 = m.ExcPending
	if v5248 != 0 {
		goto L24
	} else {
		goto L751
	}
L709:
	;
	v5243 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+1731)) = uint16(v5243)
	goto L706
L710:
	;
	v5189 = v45 + int32(1728)
	F__bt_metaversion(m, v47, v5189, v5189|int32(1))
	mBase = m.M
	v5195 = m.ExcPending
	if v5195 != 0 {
		goto L24
	} else {
		goto L750
	}
L711:
	;
	v5186 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+1731)) = uint16(v5186)
	goto L706
L712:
	;
	v5132 = v45 + int32(1728)
	F__bt_metaversion(m, v47, v5132, v5132|int32(1))
	mBase = m.M
	v5138 = m.ExcPending
	if v5138 != 0 {
		goto L24
	} else {
		goto L749
	}
L713:
	;
	if l1 != int32(-1) {
		goto L707
	} else {
		goto L748
	}
L714:
	;
	v5125 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+1731)) = uint16(v5125)
	goto L706
L715:
	;
	v5071 = v45 + int32(1728)
	F__bt_metaversion(m, v47, v5071, v5071|int32(1))
	mBase = m.M
	v5077 = m.ExcPending
	if v5077 != 0 {
		goto L24
	} else {
		goto L747
	}
L716:
	;
	v5068 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+1731)) = uint16(v5068)
	goto L706
L717:
	;
	v5014 = v45 + int32(1728)
	F__bt_metaversion(m, v47, v5014, v5014|int32(1))
	mBase = m.M
	v5020 = m.ExcPending
	if v5020 != 0 {
		goto L24
	} else {
		goto L746
	}
L718:
	;
	v4997 = v45 + int32(1728)
	F__bt_metaversion(m, v47, v4997, v4997|int32(1))
	mBase = m.M
	v5003 = m.ExcPending
	if v5003 != 0 {
		goto L24
	} else {
		goto L745
	}
L719:
	;
	v4752 = v45 + int32(1744)
	v4762 = int32(0)
	goto L720
L720:
	;
	v4797 = v4762 << (uint(int32(2)) % 32)
	v4801 = *(*int32)(unsafe.Add(mBase, uint32(v4797+(v45+int32(1600)))))
	v4802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4801))))
	if v4802&int32(4) != 0 {
		goto L722
	} else {
		goto L723
	}
L721:
	;
	v4983 = v4734
	goto L718
L722:
	;
	v4807 = v4752 + v4762*int32(48)
	v4808 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+44))
	v4809 = *(*int64)(unsafe.Add(mBase, uint32(v4808)))
	*(*int64)(unsafe.Add(mBase, uint32(v4807))) = v4809
	v4811 = *(*int64)(unsafe.Add(mBase, uint32(v4808)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4807)+40)) = v4811
	v4813 = *(*int64)(unsafe.Add(mBase, uint32(v4808)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4807)+32)) = v4813
	v4815 = *(*int64)(unsafe.Add(mBase, uint32(v4808)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4807)+24)) = v4815
	v4817 = *(*int64)(unsafe.Add(mBase, uint32(v4808)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4807)+16)) = v4817
	v4819 = *(*int64)(unsafe.Add(mBase, uint32(v4808)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4807)+8)) = v4819
	v4823 = v4808
	v4850 = v4734
	goto L726
L723:
	;
	goto L724
L724:
	;
	v4900 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+8))
	if v4900 != 0 {
		goto L735
	} else {
		goto L736
	}
L725:
	;
	switch v4712&int32(65535) - int32(2) {
	case 0:
		goto L717
	default:
		v4983 = v4850
		goto L718
	case 2:
		goto L710
	}
L726:
	;
	v4864 = v4823 + int32(48)
	v4865 = *(*int32)(unsafe.Add(mBase, uint32(v4864)))
	if v4865&int32(1) != 0 {
		goto L725
	} else {
		goto L728
	}
L727:
	;
	switch v4712&int32(65535) - int32(1) {
	case 0:
		goto L715
	default:
		v4983 = v4850
		goto L718
	case 4:
		goto L712
	}
L728:
	;
	if v4865&int32(196608) != 0 {
		goto L729
	} else {
		goto L730
	}
L729:
	;
	v4872 = v4752 + v4850*int32(48)
	v4873 = *(*int64)(unsafe.Add(mBase, uint32(v4864)))
	*(*int64)(unsafe.Add(mBase, uint32(v4872))) = v4873
	v4875 = *(*int64)(unsafe.Add(mBase, uint32(v4864)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4872)+40)) = v4875
	v4877 = *(*int64)(unsafe.Add(mBase, uint32(v4864)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4872)+32)) = v4877
	v4879 = *(*int64)(unsafe.Add(mBase, uint32(v4864)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4872)+24)) = v4879
	v4881 = *(*int64)(unsafe.Add(mBase, uint32(v4864)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4872)+16)) = v4881
	v4883 = *(*int64)(unsafe.Add(mBase, uint32(v4864)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4872)+8)) = v4883
	v4886 = v4850 + int32(1)
	v4887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4864))))
	if v4887&int32(16) == int32(0) {
		v4823 = v4864
		v4850 = v4886
		goto L726
	} else {
		goto L732
	}
L730:
	;
	goto L731
L731:
	;
	goto L727
L732:
	;
	v4983 = v4886
	goto L718
L733:
	;
	v4952 = v4762 + int32(1)
	if v4952 != v4734 {
		v4762 = v4952
		goto L720
	} else {
		goto L744
	}
L734:
	;
	v4930 = *(*int32)(unsafe.Add(mBase, uint32(v47)+208))
	v4932 = *(*int32)(unsafe.Add(mBase, uint32(v4930+v4797)))
	v4934 = F_get_opfamily_proc(m, v4932, v4903, v4900, int32(1))
	mBase = m.M
	v4935 = m.ExcPending
	if v4935 != 0 {
		goto L24
	} else {
		goto L741
	}
L735:
	;
	v4901 = *(*int32)(unsafe.Add(mBase, uint32(v47)+212))
	v4903 = *(*int32)(unsafe.Add(mBase, uint32(v4901+v4797)))
	if v4900 != v4903 {
		goto L734
	} else {
		goto L738
	}
L736:
	;
	goto L737
L737:
	;
	v4906 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4801)+4)))
	v4908 = F_index_getprocinfo(m, v47, v4906, int32(1))
	mBase = m.M
	v4909 = m.ExcPending
	if v4909 != 0 {
		goto L24
	} else {
		goto L739
	}
L738:
	;
	goto L737
L739:
	;
	v4912 = v4752 + v4762*int32(48)
	v4913 = *(*int32)(unsafe.Add(mBase, uint32(v4801)))
	v4914 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4801)+4)))
	v4915 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+8))
	v4916 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+12))
	v4917 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v4912)+44)) = v4917
	*(*int32)(unsafe.Add(mBase, uint32(v4912)+12)) = v4916
	*(*int32)(unsafe.Add(mBase, uint32(v4912)+8)) = v4915
	v4921 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v4912)+6)) = uint16(v4921)
	*(*uint16)(unsafe.Add(mBase, uint32(v4912)+4)) = uint16(v4914)
	*(*int32)(unsafe.Add(mBase, uint32(v4912))) = v4913
	v4928 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	F_fmgr_info_copy(m, v4912+int32(16), v4908, v4928)
	mBase = m.M
	goto L740
L740:
	;
	goto L733
L741:
	;
	if v4934 == int32(0) {
		goto L705
	} else {
		goto L742
	}
L742:
	;
	v4941 = *(*int32)(unsafe.Add(mBase, uint32(v4801)))
	v4942 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4801)+4)))
	v4944 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+8))
	v4945 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+12))
	v4946 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+44))
	F_ScanKeyEntryInitialize(m, v4752+v4762*int32(48), v4941, v4942, int32(0), v4944, v4945, v4934, v4946)
	mBase = m.M
	v4948 = m.ExcPending
	if v4948 != 0 {
		goto L24
	} else {
		goto L743
	}
L743:
	;
	goto L733
L744:
	;
	goto L721
L745:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1740)) = v4983
	v5005 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1736)) = v5005
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+1730)) = uint8(v5005)
	v5010 = v4712 & int32(65535)
	switch v5010 - int32(1) {
	case 0:
		goto L716
	case 1:
		goto L714
	case 2:
		goto L713
	case 3:
		goto L711
	case 4:
		goto L709
	default:
		goto L708
	}
L746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1740)) = v4850
	v5022 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1736)) = v5022
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+1730)) = uint8(v5022)
	goto L716
L747:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1740)) = v4850
	v5079 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1736)) = v5079
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+1730)) = uint8(v5079)
	goto L714
L748:
	;
	v5129 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+1731)) = uint16(v5129)
	goto L706
L749:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1740)) = v4850
	v5140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1736)) = v5140
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+1730)) = uint8(v5140)
	goto L711
L750:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1740)) = v4850
	v5197 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+1736)) = v5197
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+1730)) = uint8(v5197)
	goto L709
L751:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v5010
	F_errmsg_internal(m, int32(504902), v45+int32(16))
	mBase = m.M
	v5254 = m.ExcPending
	if v5254 != 0 {
		goto L24
	} else {
		goto L752
	}
L752:
	;
	F_errfinish(m, int32(521556), int32(1511), int32(73755))
	mBase = m.M
	v5259 = m.ExcPending
	if v5259 != 0 {
		goto L24
	} else {
		goto L753
	}
L753:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L754:
	;
	F__bt_freestack(m, v5310)
	mBase = m.M
	v5313 = m.ExcPending
	if v5313 != 0 {
		goto L24
	} else {
		goto L755
	}
L755:
	;
	v5314 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v5314 == int32(0) {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v5318 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v5318 != int32(3) {
		goto L618
	} else {
		goto L759
	}
L757:
	;
	v5335 = v5314
	goto L758
L758:
	;
	v5336 = int32(0)
	if v5335 < v5336 {
		goto L766
	} else {
		goto L767
	}
L759:
	;
	v5321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockRelation(m, v47, v5321)
	mBase = m.M
	v5323 = m.ExcPending
	if v5323 != 0 {
		goto L24
	} else {
		goto L760
	}
L760:
	;
	v5328 = F__bt_search(m, v47, int32(0), v45+int32(1728), v5308, int32(1))
	mBase = m.M
	v5329 = m.ExcPending
	if v5329 != 0 {
		goto L24
	} else {
		goto L761
	}
L761:
	;
	F__bt_freestack(m, v5328)
	mBase = m.M
	v5331 = m.ExcPending
	if v5331 != 0 {
		goto L24
	} else {
		goto L762
	}
L762:
	;
	v5332 = *(*int32)(unsafe.Add(mBase, uint32(v5308)))
	if v5332 == int32(0) {
		goto L618
	} else {
		goto L763
	}
L763:
	;
	v5335 = v5332
	goto L758
L764:
	;
	v5493 = F__bt_readfirstpage(m, l0, v5451&int32(65535), l1)
	mBase = m.M
	v5494 = m.ExcPending
	if v5494 != 0 {
		goto L24
	} else {
		goto L789
	}
L765:
	;
	v5355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5354)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v5355) {
		goto L769
	} else {
		goto L770
	}
L766:
	;
	v5340 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v5346 = *(*int32)(unsafe.Add(mBase, uint32(v5340+(v5335^int32(-1))<<(uint(int32(2))%32))))
	v5354 = v5346
	goto L765
L767:
	;
	goto L768
L768:
	;
	v5348 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v5354 = v5348 + v5335<<(uint(int32(13))%32) + int32(-8192)
	goto L765
L769:
	;
	v5363 = int32(base.Ui32(v5355+int32(262120)) >> (uint(int32(2)) % 32))
	goto L771
L770:
	;
	v5363 = int32(0)
	goto L771
L771:
	;
	v5368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5354)+16)))
	v5369 = v5354 + v5368
	v5370 = *(*int32)(unsafe.Add(mBase, uint32(v5369)+4))
	if v5370 != 0 {
		goto L772
	} else {
		goto L773
	}
L772:
	;
	v5371 = int32(2)
	goto L774
L773:
	;
	v5371 = int32(1)
	goto L774
L774:
	;
	if base.Ui32(v5363&int32(65535)) < base.Ui32(v5371) {
		v5451 = v5371
		goto L764
	} else {
		goto L775
	}
L775:
	;
	v5373 = int32(1)
	v5375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1731)))
	v5380 = v5371
	v5381 = v5363 + v5373
	goto L776
L776:
	;
	v5425 = int32(base.Ui32((v5381-v5380)&int32(65534))>>(uint(int32(1))%32)) + v5380
	v5430 = F__bt_compare(m, v47, v45+int32(1728), v5354, v5425&int32(65535))
	mBase = m.M
	v5431 = m.ExcPending
	if v5431 != 0 {
		goto L24
	} else {
		goto L778
	}
L777:
	;
	v5442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5369)+12)))
	if v5442&int32(1) != 0 {
		goto L786
	} else {
		goto L787
	}
L778:
	;
	v5432 = base.B2i32(v5430 < v5375^v5373)
	if v5430 < v5375^v5373 {
		goto L779
	} else {
		goto L780
	}
L779:
	;
	v5433 = v5425
	goto L781
L780:
	;
	v5433 = v5381
	goto L781
L781:
	;
	if v5430 < v5375^v5373 {
		goto L782
	} else {
		goto L783
	}
L782:
	;
	v5438 = v5380
	goto L784
L783:
	;
	v5438 = v5425 + int32(1)
	goto L784
L784:
	;
	if base.Ui32(v5438&int32(65535)) < base.Ui32(v5433&int32(65535)) {
		v5380 = v5438
		v5381 = v5433
		goto L776
	} else {
		goto L785
	}
L785:
	;
	goto L777
L786:
	;
	v5445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1732)))
	v5451 = v5438 - v5445
	goto L764
L787:
	;
	goto L788
L788:
	;
	v5451 = v5438 - int32(1)
	goto L764
L789:
	;
	if v5493 == int32(0) {
		v5756 = v5336
		goto L617
	} else {
		goto L790
	}
L790:
	;
	v5499 = *(*int32)(unsafe.Add(mBase, uint32(v48)+100))
	v5502 = v48 + v5499*int32(10)
	v5503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5502)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v5503)
	v5506 = v5502 + int32(104)
	v5507 = *(*int32)(unsafe.Add(mBase, uint32(v5506)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v5507
	v5509 = int32(1)
	v5510 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
	if v5510 == int32(0) {
		v5756 = v5509
		goto L617
	} else {
		goto L791
	}
L791:
	;
	v5513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5506)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v5510 + v5513
	v5756 = v5509
	goto L617
L792:
	;
	v5520 = *(*int32)(unsafe.Add(mBase, uint32(v47)+212))
	v5524 = *(*int32)(unsafe.Add(mBase, uint32(v5520+v4762<<(uint(int32(2))%32))))
	v5525 = *(*int32)(unsafe.Add(mBase, uint32(v4801)+8))
	v5526 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4801)+4)))
	v5527 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+48)) = v5527 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+44)) = v5526
	*(*int32)(unsafe.Add(mBase, uint32(v45)+40)) = v5525
	*(*int32)(unsafe.Add(mBase, uint32(v45)+36)) = v5524
	*(*int32)(unsafe.Add(mBase, uint32(v45)+32)) = int32(1)
	F_errmsg_internal(m, int32(728794), v45+int32(32))
	mBase = m.M
	v5540 = m.ExcPending
	if v5540 != 0 {
		goto L24
	} else {
		goto L793
	}
L793:
	;
	F_errfinish(m, int32(521556), int32(1430), int32(73755))
	mBase = m.M
	v5545 = m.ExcPending
	if v5545 != 0 {
		goto L24
	} else {
		goto L794
	}
L794:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L795:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5589)+56)) = v5594
	if v5594 == int32(0) {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v5599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockRelation(m, v5590, v5599)
	mBase = m.M
	v5601 = m.ExcPending
	if v5601 != 0 {
		goto L24
	} else {
		goto L799
	}
L797:
	;
	goto L798
L798:
	;
	if v5594 < int32(0) {
		goto L802
	} else {
		goto L803
	}
L799:
	;
	F__bt_parallel_done(m, l0)
	mBase = m.M
	v5603 = m.ExcPending
	if v5603 != 0 {
		goto L24
	} else {
		goto L800
	}
L800:
	;
	v5756 = v5588
	goto L617
L801:
	;
	if l1 == int32(1) {
		goto L807
	} else {
		goto L808
	}
L802:
	;
	v5607 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v5613 = *(*int32)(unsafe.Add(mBase, uint32(v5607+(v5594^int32(-1))<<(uint(int32(2))%32))))
	v5621 = v5613
	goto L801
L803:
	;
	goto L804
L804:
	;
	v5615 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v5621 = v5615 + v5594<<(uint(int32(13))%32) + int32(-8192)
	goto L801
L805:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5671 = m.ExcPending
	if v5671 != 0 {
		goto L24
	} else {
		goto L820
	}
L806:
	;
	v5645 = F__bt_readfirstpage(m, l0, v5642&int32(65535), l1)
	mBase = m.M
	v5646 = m.ExcPending
	if v5646 != 0 {
		goto L24
	} else {
		goto L817
	}
L807:
	;
	v5626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5621)+16)))
	v5628 = *(*int32)(unsafe.Add(mBase, uint32(v5621+v5626)+4))
	if v5628 != 0 {
		goto L810
	} else {
		goto L811
	}
L808:
	;
	goto L809
L809:
	;
	if l1 != int32(-1) {
		goto L805
	} else {
		goto L813
	}
L810:
	;
	v5629 = int32(2)
	goto L812
L811:
	;
	v5629 = int32(1)
	goto L812
L812:
	;
	v5642 = v5629
	goto L806
L813:
	;
	v5632 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5621)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v5632) {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	v5640 = int32(base.Ui32(v5632+int32(262120)) >> (uint(int32(2)) % 32))
	goto L816
L815:
	;
	v5640 = int32(0)
	goto L816
L816:
	;
	v5642 = v5640
	goto L806
L817:
	;
	if v5645 == int32(0) {
		v5756 = v5588
		goto L617
	} else {
		goto L818
	}
L818:
	;
	v5651 = *(*int32)(unsafe.Add(mBase, uint32(v5589)+100))
	v5654 = v5589 + v5651*int32(10)
	v5655 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5654)+108)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v5655)
	v5658 = v5654 + int32(104)
	v5659 = *(*int32)(unsafe.Add(mBase, uint32(v5658)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v5659
	v5661 = int32(1)
	v5662 = *(*int32)(unsafe.Add(mBase, uint32(v5589)+44))
	if v5662 == int32(0) {
		v5756 = v5661
		goto L617
	} else {
		goto L819
	}
L819:
	;
	v5665 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5658)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v5662 + v5665
	v5756 = v5661
	goto L617
L820:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = l1
	F_errmsg_internal(m, int32(504626), v45)
	mBase = m.M
	v5675 = m.ExcPending
	if v5675 != 0 {
		goto L24
	} else {
		goto L821
	}
L821:
	;
	F_errfinish(m, int32(521556), int32(2744), int32(94784))
	mBase = m.M
	v5680 = m.ExcPending
	if v5680 != 0 {
		goto L24
	} else {
		goto L822
	}
L822:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L823:
	;
	v5756 = int32(0)
	goto L617
}
