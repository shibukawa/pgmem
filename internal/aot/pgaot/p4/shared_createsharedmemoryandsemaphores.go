package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateSharedMemoryAndSemaphores(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v692 int32
	_ = v692
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
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
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1262 int32
	_ = v1262
	var v1267 int64
	_ = v1267
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1316 int64
	_ = v1316
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1352 int32
	_ = v1352
	var v1354 int64
	_ = v1354
	var v1378 int32
	_ = v1378
	var v1460 int32
	_ = v1460
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1601 int32
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
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
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1748 int32
	_ = v1748
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1867 int32
	_ = v1867
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1910 int32
	_ = v1910
	var v1933 int32
	_ = v1933
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1987 int64
	_ = v1987
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v2001 int32
	_ = v2001
	var v2008 int32
	_ = v2008
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2097 int32
	_ = v2097
	var v2102 int32
	_ = v2102
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2136 int32
	_ = v2136
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int64
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2187 int32
	_ = v2187
	var v2192 int32
	_ = v2192
	var v2203 int32
	_ = v2203
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2244 int32
	_ = v2244
	var v2245 int64
	_ = v2245
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2258 int32
	_ = v2258
	var v2259 int64
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2272 int32
	_ = v2272
	var v2273 int64
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2286 int32
	_ = v2286
	var v2287 int64
	_ = v2287
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2300 int32
	_ = v2300
	var v2301 int64
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2314 int32
	_ = v2314
	var v2315 int64
	_ = v2315
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2328 int32
	_ = v2328
	var v2329 int64
	_ = v2329
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2342 int32
	_ = v2342
	var v2343 int64
	_ = v2343
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2364 int64
	_ = v2364
	var v2370 int32
	_ = v2370
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2425 int64
	_ = v2425
	var v2426 int64
	_ = v2426
	var v2437 int32
	_ = v2437
	var v2438 int64
	_ = v2438
	var v2451 int32
	_ = v2451
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2475 int32
	_ = v2475
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2500 int32
	_ = v2500
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2539 int32
	_ = v2539
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2557 int32
	_ = v2557
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2565 int32
	_ = v2565
	var v2567 int32
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2575 int32
	_ = v2575
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2594 int32
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2599 int32
	_ = v2599
	var v2601 int32
	_ = v2601
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2607 int32
	_ = v2607
	var v2610 int32
	_ = v2610
	var v2613 int32
	_ = v2613
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2619 int32
	_ = v2619
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2629 int32
	_ = v2629
	var v2634 int32
	_ = v2634
	var v2637 int32
	_ = v2637
	var v2639 int32
	_ = v2639
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2656 int32
	_ = v2656
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2676 int32
	_ = v2676
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2697 int32
	_ = v2697
	var v2700 int32
	_ = v2700
	var v2702 int32
	_ = v2702
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2714 int32
	_ = v2714
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2726 int32
	_ = v2726
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2742 int32
	_ = v2742
	var v2747 int32
	_ = v2747
	var v2750 int32
	_ = v2750
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2769 int32
	_ = v2769
	var v2772 int32
	_ = v2772
	var v2775 int32
	_ = v2775
	var v2778 int32
	_ = v2778
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2789 int32
	_ = v2789
	var v2793 int32
	_ = v2793
	var v2795 int32
	_ = v2795
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2814 int32
	_ = v2814
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2826 int32
	_ = v2826
	var v2832 int32
	_ = v2832
	var v2834 int32
	_ = v2834
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
	var v2843 int32
	_ = v2843
	var v2846 int32
	_ = v2846
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2882 int32
	_ = v2882
	var v2892 int32
	_ = v2892
	var v2894 int32
	_ = v2894
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2903 int32
	_ = v2903
	var v2908 int32
	_ = v2908
	var v2912 int32
	_ = v2912
	var v2914 int32
	_ = v2914
	var v2919 int32
	_ = v2919
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2930 int32
	_ = v2930
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2947 int32
	_ = v2947
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2958 int32
	_ = v2958
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2970 int32
	_ = v2970
	var v2973 int32
	_ = v2973
	var v2977 int32
	_ = v2977
	var v2982 int32
	_ = v2982
	var v2997 int32
	_ = v2997
	var v3000 int32
	_ = v3000
	var v3003 int32
	_ = v3003
	var v3017 int32
	_ = v3017
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3058 int32
	_ = v3058
	var v3082 int32
	_ = v3082
	var v3084 int32
	_ = v3084
	var v3087 int32
	_ = v3087
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3092 int32
	_ = v3092
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3118 int32
	_ = v3118
	var v3123 int32
	_ = v3123
	var v3139 int32
	_ = v3139
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3168 int32
	_ = v3168
	var v3170 int32
	_ = v3170
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3183 int32
	_ = v3183
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3196 int32
	_ = v3196
	var v3197 int32
	_ = v3197
	var v3199 int32
	_ = v3199
	var v3204 int32
	_ = v3204
	var v3207 int32
	_ = v3207
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3235 int32
	_ = v3235
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3251 int32
	_ = v3251
	var v3269 int32
	_ = v3269
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3280 int32
	_ = v3280
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
	var v3289 int32
	_ = v3289
	var v3290 int32
	_ = v3290
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3296 int32
	_ = v3296
	var v3299 int32
	_ = v3299
	var v3302 int64
	_ = v3302
	var v3306 int32
	_ = v3306
	var v3311 int32
	_ = v3311
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3363 int32
	_ = v3363
	var v3368 int32
	_ = v3368
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3379 int32
	_ = v3379
	var v3382 int32
	_ = v3382
	var v3383 int32
	_ = v3383
	var v3391 int32
	_ = v3391
	var v3395 int32
	_ = v3395
	var v3396 int32
	_ = v3396
	var v3412 int32
	_ = v3412
	var v3414 int32
	_ = v3414
	var v3415 int32
	_ = v3415
	var v3417 int32
	_ = v3417
	var v3418 int64
	_ = v3418
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3428 int32
	_ = v3428
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3436 int32
	_ = v3436
	var v3438 int32
	_ = v3438
	var v3441 int32
	_ = v3441
	var v3444 int32
	_ = v3444
	var v3446 int32
	_ = v3446
	var v3449 int32
	_ = v3449
	var v3452 int32
	_ = v3452
	var v3455 int32
	_ = v3455
	var v3458 int32
	_ = v3458
	var v3461 int32
	_ = v3461
	var v3464 int32
	_ = v3464
	var v3467 int32
	_ = v3467
	var v3485 int32
	_ = v3485
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3500 int32
	_ = v3500
	var v3502 int32
	_ = v3502
	var v3503 int32
	_ = v3503
	var v3505 int32
	_ = v3505
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3511 int32
	_ = v3511
	var v3523 int32
	_ = v3523
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3541 int32
	_ = v3541
	var v3547 int32
	_ = v3547
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3554 int32
	_ = v3554
	var v3555 int32
	_ = v3555
	var v3561 int32
	_ = v3561
	var v3566 int32
	_ = v3566
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3592 int32
	_ = v3592
	var v3603 int32
	_ = v3603
	var v3604 int32
	_ = v3604
	var v3611 int32
	_ = v3611
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3620 int32
	_ = v3620
	var v3624 int32
	_ = v3624
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3631 int32
	_ = v3631
	var v3637 int32
	_ = v3637
	var v3641 int32
	_ = v3641
	var v3647 int32
	_ = v3647
	var v3650 int32
	_ = v3650
	var v3652 int32
	_ = v3652
	var v3655 int32
	_ = v3655
	var v3657 int32
	_ = v3657
	var v3662 int32
	_ = v3662
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3666 int32
	_ = v3666
	var v3669 int32
	_ = v3669
	var v3673 int32
	_ = v3673
	var v3677 int32
	_ = v3677
	var v3681 int32
	_ = v3681
	var v3692 int32
	_ = v3692
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3698 int32
	_ = v3698
	var v3712 int32
	_ = v3712
	var v3714 int32
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3722 int32
	_ = v3722
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3734 int32
	_ = v3734
	var v3736 int32
	_ = v3736
	var v3739 int32
	_ = v3739
	var v3742 int32
	_ = v3742
	var v3749 int32
	_ = v3749
	var v3753 int32
	_ = v3753
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3762 int32
	_ = v3762
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3770 int32
	_ = v3770
	var v3773 int32
	_ = v3773
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3789 int32
	_ = v3789
	var v3791 int32
	_ = v3791
	var v3793 int32
	_ = v3793
	var v3799 int32
	_ = v3799
	var v3814 int32
	_ = v3814
	var v3815 int32
	_ = v3815
	var v3830 int32
	_ = v3830
	var v3831 int32
	_ = v3831
	var v3835 int32
	_ = v3835
	var v3840 int32
	_ = v3840
	var v3842 int32
	_ = v3842
	var v3845 int32
	_ = v3845
	var v3849 int32
	_ = v3849
	var v3851 int32
	_ = v3851
	var v3858 int32
	_ = v3858
	var v3862 int32
	_ = v3862
	var v3867 int32
	_ = v3867
	var v3870 int32
	_ = v3870
	var v3875 int32
	_ = v3875
	var v3878 int32
	_ = v3878
	var v3879 int32
	_ = v3879
	var v3887 int32
	_ = v3887
	var v3890 int32
	_ = v3890
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3899 int32
	_ = v3899
	var v3905 int32
	_ = v3905
	var v3908 int32
	_ = v3908
	var v3911 int32
	_ = v3911
	var v3913 int32
	_ = v3913
	var v3914 int32
	_ = v3914
	var v3920 int32
	_ = v3920
	var v3926 int32
	_ = v3926
	var v3930 int32
	_ = v3930
	var v3932 int32
	_ = v3932
	var v3933 int32
	_ = v3933
	var v3939 int32
	_ = v3939
	var v3945 int32
	_ = v3945
	var v3948 int32
	_ = v3948
	var v3950 int32
	_ = v3950
	var v3951 int32
	_ = v3951
	var v3957 int32
	_ = v3957
	var v3964 int32
	_ = v3964
	var v3966 int32
	_ = v3966
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3981 int32
	_ = v3981
	var v3985 int32
	_ = v3985
	var v3989 int32
	_ = v3989
	var v3993 int32
	_ = v3993
	var v3997 int32
	_ = v3997
	var v4001 int32
	_ = v4001
	var v4005 int32
	_ = v4005
	var v4009 int32
	_ = v4009
	var v4013 int32
	_ = v4013
	var v4017 int32
	_ = v4017
	var v4021 int32
	_ = v4021
	var v4025 int32
	_ = v4025
	var v4029 int32
	_ = v4029
	var v4033 int32
	_ = v4033
	var v4037 int32
	_ = v4037
	var v4040 int32
	_ = v4040
	var v4047 int32
	_ = v4047
	var v4067 int32
	_ = v4067
	var v4070 int32
	_ = v4070
	var v4081 int32
	_ = v4081
	var v4082 int32
	_ = v4082
	var v4105 int32
	_ = v4105
	var v4107 int32
	_ = v4107
	var v4114 int32
	_ = v4114
	var v4116 int32
	_ = v4116
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4124 int32
	_ = v4124
	var v4125 int32
	_ = v4125
	var v4127 int32
	_ = v4127
	var v4133 int32
	_ = v4133
	var v4135 int32
	_ = v4135
	var v4136 int64
	_ = v4136
	var v4142 int32
	_ = v4142
	var v4148 int32
	_ = v4148
	var v4155 int32
	_ = v4155
	var v4156 int32
	_ = v4156
	var v4159 int32
	_ = v4159
	var v4166 int32
	_ = v4166
	var v4168 int32
	_ = v4168
	var v4172 int32
	_ = v4172
	var v4173 int32
	_ = v4173
	var v4175 int32
	_ = v4175
	var v4176 int32
	_ = v4176
	var v4177 int32
	_ = v4177
	var v4183 int32
	_ = v4183
	var v4185 int32
	_ = v4185
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4191 int32
	_ = v4191
	var v4192 int32
	_ = v4192
	var v4195 int32
	_ = v4195
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4206 int32
	_ = v4206
	var v4209 int32
	_ = v4209
	var v4210 int32
	_ = v4210
	var v4213 int32
	_ = v4213
	var v4214 int32
	_ = v4214
	var v4216 int32
	_ = v4216
	var v4217 int32
	_ = v4217
	var v4231 int32
	_ = v4231
	var v4233 int32
	_ = v4233
	var v4235 int32
	_ = v4235
	var v4241 int32
	_ = v4241
	var v4253 int32
	_ = v4253
	var v4256 int32
	_ = v4256
	var v4257 int32
	_ = v4257
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4263 int32
	_ = v4263
	var v4267 int32
	_ = v4267
	var v4270 int32
	_ = v4270
	var v4284 int32
	_ = v4284
	var v4286 int32
	_ = v4286
	var v4288 int32
	_ = v4288
	var v4294 int32
	_ = v4294
	var v4303 int32
	_ = v4303
	var v4305 int32
	_ = v4305
	var v4309 int32
	_ = v4309
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4334 int32
	_ = v4334
	var v4336 int32
	_ = v4336
	var v4338 int32
	_ = v4338
	var v4342 int32
	_ = v4342
	var v4359 int32
	_ = v4359
	var v4360 int32
	_ = v4360
	var v4363 int32
	_ = v4363
	var v4364 int32
	_ = v4364
	var v4366 int32
	_ = v4366
	var v4370 int32
	_ = v4370
	var v4373 int32
	_ = v4373
	var v4387 int32
	_ = v4387
	var v4389 int32
	_ = v4389
	var v4391 int32
	_ = v4391
	var v4397 int32
	_ = v4397
	var v4406 int32
	_ = v4406
	var v4408 int32
	_ = v4408
	var v4412 int32
	_ = v4412
	var v4414 int32
	_ = v4414
	var v4415 int32
	_ = v4415
	var v4437 int32
	_ = v4437
	var v4439 int32
	_ = v4439
	var v4441 int32
	_ = v4441
	var v4445 int32
	_ = v4445
	var v4461 int32
	_ = v4461
	var v4462 int32
	_ = v4462
	var v4463 int32
	_ = v4463
	var v4469 int32
	_ = v4469
	var v4470 int32
	_ = v4470
	var v4472 int32
	_ = v4472
	var v4473 int32
	_ = v4473
	var v4476 int32
	_ = v4476
	var v4489 int32
	_ = v4489
	var v4491 int32
	_ = v4491
	var v4493 int32
	_ = v4493
	var v4499 int32
	_ = v4499
	var v4508 int32
	_ = v4508
	var v4514 int32
	_ = v4514
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4537 int32
	_ = v4537
	var v4540 int32
	_ = v4540
	var v4542 int32
	_ = v4542
	var v4562 int32
	_ = v4562
	var v4566 int32
	_ = v4566
	var v4568 int32
	_ = v4568
	var v4574 int32
	_ = v4574
	var v4576 int32
	_ = v4576
	var v4577 int32
	_ = v4577
	var v4578 int32
	_ = v4578
	var v4579 int32
	_ = v4579
	var v4585 int32
	_ = v4585
	var v4587 int32
	_ = v4587
	var v4588 int32
	_ = v4588
	var v4589 int32
	_ = v4589
	var v4590 int32
	_ = v4590
	var v4593 int32
	_ = v4593
	var v4594 int32
	_ = v4594
	var v4597 int32
	_ = v4597
	var v4601 int32
	_ = v4601
	var v4613 int32
	_ = v4613
	var v4614 int32
	_ = v4614
	var v4630 int32
	_ = v4630
	var v4634 int32
	_ = v4634
	var v4635 int32
	_ = v4635
	var v4639 int32
	_ = v4639
	var v4640 int32
	_ = v4640
	var v4643 int32
	_ = v4643
	var v4646 int32
	_ = v4646
	var v4648 int32
	_ = v4648
	var v4666 int32
	_ = v4666
	var v4669 int32
	_ = v4669
	var v4670 int32
	_ = v4670
	var v4672 int32
	_ = v4672
	var v4678 int32
	_ = v4678
	var v4680 int32
	_ = v4680
	var v4681 int32
	_ = v4681
	var v4682 int32
	_ = v4682
	var v4683 int32
	_ = v4683
	var v4686 int32
	_ = v4686
	var v4687 int32
	_ = v4687
	var v4690 int32
	_ = v4690
	var v4692 int32
	_ = v4692
	var v4697 int32
	_ = v4697
	var v4698 int32
	_ = v4698
	var v4705 int32
	_ = v4705
	var v4706 int32
	_ = v4706
	var v4722 int32
	_ = v4722
	var v4723 int32
	_ = v4723
	var v4725 int32
	_ = v4725
	var v4730 int32
	_ = v4730
	var v4746 int32
	_ = v4746
	var v4747 int32
	_ = v4747
	var v4751 int32
	_ = v4751
	var v4753 int32
	_ = v4753
	var v4754 int32
	_ = v4754
	var v4770 int32
	_ = v4770
	var v4774 int32
	_ = v4774
	var v4792 int32
	_ = v4792
	var v4795 int32
	_ = v4795
	var v4797 int32
	_ = v4797
	var v4815 int32
	_ = v4815
	var v4818 int32
	_ = v4818
	var v4819 int32
	_ = v4819
	var v4821 int32
	_ = v4821
	var v4828 int32
	_ = v4828
	var v4831 int32
	_ = v4831
	var v4832 int32
	_ = v4832
	var v4833 int32
	_ = v4833
	var v4834 int32
	_ = v4834
	var v4837 int32
	_ = v4837
	var v4840 int32
	_ = v4840
	var v4841 int32
	_ = v4841
	var v4842 int32
	_ = v4842
	var v4843 int32
	_ = v4843
	var v4846 int32
	_ = v4846
	var v4847 int32
	_ = v4847
	var v4849 int32
	_ = v4849
	var v4857 int32
	_ = v4857
	var v4866 int32
	_ = v4866
	var v4881 int32
	_ = v4881
	var v4883 int32
	_ = v4883
	var v4885 int32
	_ = v4885
	var v4899 int32
	_ = v4899
	var v4901 int32
	_ = v4901
	var v4907 int32
	_ = v4907
	var v4945 int32
	_ = v4945
	var v4948 int32
	_ = v4948
	var v4950 int32
	_ = v4950
	var v4955 int32
	_ = v4955
	var v4956 int32
	_ = v4956
	var v4958 int32
	_ = v4958
	var v4959 int32
	_ = v4959
	var v4960 int32
	_ = v4960
	var v4961 int32
	_ = v4961
	var v4964 int32
	_ = v4964
	var v4965 int32
	_ = v4965
	var v4967 int32
	_ = v4967
	var v4973 int32
	_ = v4973
	var v4974 int32
	_ = v4974
	var v4976 int32
	_ = v4976
	var v4977 int32
	_ = v4977
	var v4978 int32
	_ = v4978
	var v4979 int32
	_ = v4979
	var v4992 int32
	_ = v4992
	var v4994 int32
	_ = v4994
	var v4996 int32
	_ = v4996
	var v5002 int32
	_ = v5002
	var v5010 int32
	_ = v5010
	var v5011 int32
	_ = v5011
	var v5013 int32
	_ = v5013
	var v5017 int32
	_ = v5017
	var v5021 int32
	_ = v5021
	var v5023 int32
	_ = v5023
	var v5028 int32
	_ = v5028
	var v5032 int32
	_ = v5032
	var v5033 int32
	_ = v5033
	var v5035 int32
	_ = v5035
	var v5036 int32
	_ = v5036
	var v5039 int32
	_ = v5039
	var v5040 int32
	_ = v5040
	var v5042 int32
	_ = v5042
	var v5046 int32
	_ = v5046
	var v5052 int32
	_ = v5052
	var v5068 int32
	_ = v5068
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5080 int64
	_ = v5080
	var v5099 int32
	_ = v5099
	var v5105 int32
	_ = v5105
	var v5107 int32
	_ = v5107
	var v5127 int32
	_ = v5127
	var v5130 int32
	_ = v5130
	var v5132 int32
	_ = v5132
	var v5137 int32
	_ = v5137
	var v5139 int32
	_ = v5139
	var v5142 int32
	_ = v5142
	var v5144 int32
	_ = v5144
	var v5145 int32
	_ = v5145
	var v5146 int32
	_ = v5146
	var v5147 int32
	_ = v5147
	var v5150 int32
	_ = v5150
	var v5151 int32
	_ = v5151
	var v5153 int32
	_ = v5153
	var v5156 int32
	_ = v5156
	var v5170 int32
	_ = v5170
	var v5172 int32
	_ = v5172
	var v5174 int32
	_ = v5174
	var v5180 int32
	_ = v5180
	var v5190 int32
	_ = v5190
	var v5192 int32
	_ = v5192
	var v5195 int32
	_ = v5195
	var v5198 int32
	_ = v5198
	var v5204 int32
	_ = v5204
	var v5206 int32
	_ = v5206
	var v5213 int32
	_ = v5213
	var v5216 int32
	_ = v5216
	var v5218 int32
	_ = v5218
	var v5220 int32
	_ = v5220
	var v5226 int32
	_ = v5226
	var v5228 int32
	_ = v5228
	var v5229 int32
	_ = v5229
	var v5230 int32
	_ = v5230
	var v5231 int32
	_ = v5231
	var v5234 int32
	_ = v5234
	var v5235 int32
	_ = v5235
	var v5238 int32
	_ = v5238
	var v5241 int32
	_ = v5241
	var v5246 int32
	_ = v5246
	var v5250 int32
	_ = v5250
	var v5259 int32
	_ = v5259
	var v5264 int32
	_ = v5264
	var v5267 int32
	_ = v5267
	var v5269 int32
	_ = v5269
	var v5282 int32
	_ = v5282
	var v5286 int32
	_ = v5286
	var v5287 int32
	_ = v5287
	var v5293 int32
	_ = v5293
	var v5295 int32
	_ = v5295
	var v5331 int32
	_ = v5331
	var v5334 int32
	_ = v5334
	var v5336 int32
	_ = v5336
	var v5339 int32
	_ = v5339
	var v5346 int32
	_ = v5346
	var v5347 int32
	_ = v5347
	var v5348 int32
	_ = v5348
	var v5349 int32
	_ = v5349
	var v5352 int32
	_ = v5352
	var v5353 int32
	_ = v5353
	var v5355 int32
	_ = v5355
	var v5358 int32
	_ = v5358
	var v5361 int32
	_ = v5361
	var v5362 int32
	_ = v5362
	var v5363 int32
	_ = v5363
	var v5364 int32
	_ = v5364
	var v5365 int32
	_ = v5365
	var v5366 int32
	_ = v5366
	var v5380 int32
	_ = v5380
	var v5382 int32
	_ = v5382
	var v5384 int32
	_ = v5384
	var v5390 int32
	_ = v5390
	var v5398 int32
	_ = v5398
	var v5400 int32
	_ = v5400
	var v5403 int32
	_ = v5403
	var v5420 int32
	_ = v5420
	var v5423 int32
	_ = v5423
	var v5427 int32
	_ = v5427
	var v5428 int32
	_ = v5428
	var v5435 int32
	_ = v5435
	var v5441 int32
	_ = v5441
	var v5443 int32
	_ = v5443
	var v5461 int32
	_ = v5461
	var v5464 int32
	_ = v5464
	var v5465 int32
	_ = v5465
	var v5467 int32
	_ = v5467
	var v5470 int32
	_ = v5470
	var v5477 int32
	_ = v5477
	var v5478 int32
	_ = v5478
	var v5480 int32
	_ = v5480
	var v5482 int32
	_ = v5482
	var v5483 int32
	_ = v5483
	var v5484 int32
	_ = v5484
	var v5485 int32
	_ = v5485
	var v5488 int32
	_ = v5488
	var v5489 int32
	_ = v5489
	var v5495 int32
	_ = v5495
	var v5497 int32
	_ = v5497
	var v5500 int32
	_ = v5500
	var v5501 int32
	_ = v5501
	var v5503 int32
	_ = v5503
	var v5505 int32
	_ = v5505
	var v5506 int32
	_ = v5506
	var v5507 int32
	_ = v5507
	var v5508 int32
	_ = v5508
	var v5509 int32
	_ = v5509
	var v5510 int32
	_ = v5510
	var v5524 int32
	_ = v5524
	var v5526 int32
	_ = v5526
	var v5528 int32
	_ = v5528
	var v5534 int32
	_ = v5534
	var v5542 int32
	_ = v5542
	var v5544 int32
	_ = v5544
	var v5548 int32
	_ = v5548
	var v5552 int32
	_ = v5552
	var v5568 int32
	_ = v5568
	var v5570 int32
	_ = v5570
	var v5573 int32
	_ = v5573
	var v5575 int32
	_ = v5575
	var v5576 int32
	_ = v5576
	var v5583 int32
	_ = v5583
	var v5586 int32
	_ = v5586
	var v5592 int32
	_ = v5592
	var v5594 int32
	_ = v5594
	var v5612 int32
	_ = v5612
	var v5615 int32
	_ = v5615
	var v5617 int32
	_ = v5617
	var v5623 int32
	_ = v5623
	var v5625 int32
	_ = v5625
	var v5626 int32
	_ = v5626
	var v5627 int32
	_ = v5627
	var v5628 int32
	_ = v5628
	var v5631 int32
	_ = v5631
	var v5632 int32
	_ = v5632
	var v5634 int32
	_ = v5634
	var v5641 int32
	_ = v5641
	var v5643 int32
	_ = v5643
	var v5644 int32
	_ = v5644
	var v5645 int32
	_ = v5645
	var v5646 int32
	_ = v5646
	var v5659 int32
	_ = v5659
	var v5661 int32
	_ = v5661
	var v5663 int32
	_ = v5663
	var v5669 int32
	_ = v5669
	var v5677 int32
	_ = v5677
	var v5679 int32
	_ = v5679
	var v5681 int32
	_ = v5681
	var v5685 int32
	_ = v5685
	var v5691 int32
	_ = v5691
	var v5695 int32
	_ = v5695
	var v5711 int32
	_ = v5711
	var v5718 int32
	_ = v5718
	var v5720 int32
	_ = v5720
	var v5723 int32
	_ = v5723
	var v5740 int32
	_ = v5740
	var v5742 int32
	_ = v5742
	var v5748 int32
	_ = v5748
	var v5750 int32
	_ = v5750
	var v5756 int32
	_ = v5756
	var v5758 int32
	_ = v5758
	var v5779 int32
	_ = v5779
	var v5782 int32
	_ = v5782
	var v5784 int32
	_ = v5784
	var v5790 int32
	_ = v5790
	var v5791 int32
	_ = v5791
	var v5794 int32
	_ = v5794
	var v5795 int32
	_ = v5795
	var v5797 int32
	_ = v5797
	var v5804 int32
	_ = v5804
	var v5805 int32
	_ = v5805
	var v5818 int32
	_ = v5818
	var v5820 int32
	_ = v5820
	var v5822 int32
	_ = v5822
	var v5828 int32
	_ = v5828
	var v5837 int32
	_ = v5837
	var v5841 int32
	_ = v5841
	var v5847 int32
	_ = v5847
	var v5857 int32
	_ = v5857
	var v5860 int32
	_ = v5860
	var v5862 int32
	_ = v5862
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5872 int32
	_ = v5872
	var v5875 int64
	_ = v5875
	var v5877 int32
	_ = v5877
	var v5879 int32
	_ = v5879
	var v5888 int32
	_ = v5888
	var v5893 int32
	_ = v5893
	var v5896 int32
	_ = v5896
	var v5898 int32
	_ = v5898
	var v5904 int32
	_ = v5904
	var v5905 int32
	_ = v5905
	var v5908 int32
	_ = v5908
	var v5909 int32
	_ = v5909
	var v5911 int32
	_ = v5911
	var v5918 int32
	_ = v5918
	var v5919 int32
	_ = v5919
	var v5932 int32
	_ = v5932
	var v5934 int32
	_ = v5934
	var v5936 int32
	_ = v5936
	var v5942 int32
	_ = v5942
	var v5951 int32
	_ = v5951
	var v5958 int32
	_ = v5958
	var v5961 int32
	_ = v5961
	var v5963 int32
	_ = v5963
	var v5969 int32
	_ = v5969
	var v5971 int32
	_ = v5971
	var v5972 int32
	_ = v5972
	var v5973 int32
	_ = v5973
	var v5974 int32
	_ = v5974
	var v5977 int32
	_ = v5977
	var v5978 int32
	_ = v5978
	var v5980 int32
	_ = v5980
	var v5983 int32
	_ = v5983
	var v5985 int32
	_ = v5985
	var v5986 int32
	_ = v5986
	var v5987 int32
	_ = v5987
	var v5988 int32
	_ = v5988
	var v5992 int32
	_ = v5992
	var v5996 int32
	_ = v5996
	var v6002 int32
	_ = v6002
	var v6018 int32
	_ = v6018
	var v6020 int32
	_ = v6020
	var v6021 int32
	_ = v6021
	var v6027 int32
	_ = v6027
	var v6029 int32
	_ = v6029
	var v6047 int32
	_ = v6047
	var v6050 int32
	_ = v6050
	var v6052 int32
	_ = v6052
	var v6059 int32
	_ = v6059
	var v6060 int32
	_ = v6060
	var v6062 int32
	_ = v6062
	var v6065 int64
	_ = v6065
	var v6075 int32
	_ = v6075
	var v6078 int32
	_ = v6078
	var v6080 int32
	_ = v6080
	var v6084 int32
	_ = v6084
	var v6086 int32
	_ = v6086
	var v6088 int32
	_ = v6088
	var v6089 int32
	_ = v6089
	var v6090 int32
	_ = v6090
	var v6091 int32
	_ = v6091
	var v6094 int32
	_ = v6094
	var v6095 int32
	_ = v6095
	var v6098 int32
	_ = v6098
	var v6101 int64
	_ = v6101
	var v6103 int32
	_ = v6103
	var v6108 int32
	_ = v6108
	var v6112 int32
	_ = v6112
	var v6115 int32
	_ = v6115
	var v6117 int32
	_ = v6117
	var v6124 int32
	_ = v6124
	var v6125 int32
	_ = v6125
	var v6128 int32
	_ = v6128
	var v6141 int32
	_ = v6141
	var v6147 int32
	_ = v6147
	var v6150 int32
	_ = v6150
	var v6165 int32
	_ = v6165
	var v6166 int32
	_ = v6166
	var v6167 int32
	_ = v6167
	var v6178 int32
	_ = v6178
	var v6182 int32
	_ = v6182
	var v6183 int32
	_ = v6183
	var v6186 int64
	_ = v6186
	var v6208 int32
	_ = v6208
	var v6212 int32
	_ = v6212
	var v6214 int32
	_ = v6214
	var v6219 int32
	_ = v6219
	var v6221 int32
	_ = v6221
	var v6222 int32
	_ = v6222
	var v6224 int32
	_ = v6224
	var v6225 int32
	_ = v6225
	var v6228 int32
	_ = v6228
	var v6229 int32
	_ = v6229
	var v6231 int32
	_ = v6231
	var v6232 int64
	_ = v6232
	var v6238 int32
	_ = v6238
	var v6247 int32
	_ = v6247
	var v6255 int32
	_ = v6255
	var v6269 int32
	_ = v6269
	var v6273 int32
	_ = v6273
	var v6281 int32
	_ = v6281
	var v6283 int32
	_ = v6283
	var v6307 int32
	_ = v6307
	var v6315 int32
	_ = v6315
	var v6316 int32
	_ = v6316
	var v6322 int32
	_ = v6322
	var v6323 int32
	_ = v6323
	var v6324 int32
	_ = v6324
	var v6327 int32
	_ = v6327
	var v6329 int32
	_ = v6329
	var v6333 int32
	_ = v6333
	var v6334 int32
	_ = v6334
	var v6337 int32
	_ = v6337
	var v6338 int32
	_ = v6338
	var v6341 int32
	_ = v6341
	var v6345 int32
	_ = v6345
	var v6350 int32
	_ = v6350
	var v6351 int32
	_ = v6351
	var v6353 int32
	_ = v6353
	var v6356 int32
	_ = v6356
	var v6359 int32
	_ = v6359
	var v6360 int32
	_ = v6360
	var v6361 int32
	_ = v6361
	var v6362 int32
	_ = v6362
	var v6366 int32
	_ = v6366
	var v6368 int32
	_ = v6368
	var v6370 int32
	_ = v6370
	var v6376 int32
	_ = v6376
	var v6406 int32
	_ = v6406
	var v6408 int32
	_ = v6408
	var v6416 int32
	_ = v6416
	var v6418 int32
	_ = v6418
	var v6420 int32
	_ = v6420
	var v6423 int32
	_ = v6423
	var v6430 int32
	_ = v6430
	var v6435 int32
	_ = v6435
	var v6436 int32
	_ = v6436
	var v6437 int32
	_ = v6437
	var v6439 int32
	_ = v6439
	var v6440 int32
	_ = v6440
	var v6442 int32
	_ = v6442
	var v6445 int32
	_ = v6445
	var v6467 int32
	_ = v6467
	var v6469 int32
	_ = v6469
	var v6476 int32
	_ = v6476
	var v6477 int32
	_ = v6477
	var v6479 int32
	_ = v6479
	var v6491 int32
	_ = v6491
	var v6493 int32
	_ = v6493
	var v6494 int32
	_ = v6494
	var v6503 int32
	_ = v6503
	var v6504 int32
	_ = v6504
	var v6509 int32
	_ = v6509
	var v6511 int32
	_ = v6511
	var v6513 int32
	_ = v6513
	var v6516 int32
	_ = v6516
	var v6518 int32
	_ = v6518
	var v6523 int32
	_ = v6523
	var v6524 int32
	_ = v6524
	var v6525 int32
	_ = v6525
	var v6527 int32
	_ = v6527
	var v6530 int64
	_ = v6530
	var v6537 int32
	_ = v6537
	var v6539 int32
	_ = v6539
	var v6541 int32
	_ = v6541
	var v6543 int32
	_ = v6543
	var v6551 int32
	_ = v6551
	var v6552 int32
	_ = v6552
	var v6553 int32
	_ = v6553
	var v6554 int32
	_ = v6554
	var v6556 int32
	_ = v6556
	var v6560 int32
	_ = v6560
	var v6564 int32
	_ = v6564
	var v6566 int32
	_ = v6566
	var v6567 int32
	_ = v6567
	var v6568 int32
	_ = v6568
	var v6569 int32
	_ = v6569
	var v6570 int32
	_ = v6570
	var v6571 int32
	_ = v6571
	var v6573 int32
	_ = v6573
	var v6578 int32
	_ = v6578
	var v6580 int32
	_ = v6580
	var v6583 int32
	_ = v6583
	var v6584 int32
	_ = v6584
	var v6586 int32
	_ = v6586
	var v6587 int32
	_ = v6587
	var v6588 int32
	_ = v6588
	var v6589 int32
	_ = v6589
	var v6590 int32
	_ = v6590
	var v6591 int32
	_ = v6591
	var v6592 int32
	_ = v6592
	var v6594 int32
	_ = v6594
	var v6599 int32
	_ = v6599
	var v6601 int32
	_ = v6601
	var v6604 int32
	_ = v6604
	var v6605 int32
	_ = v6605
	var v6607 int32
	_ = v6607
	var v6608 int32
	_ = v6608
	var v6609 int32
	_ = v6609
	var v6610 int32
	_ = v6610
	var v6611 int32
	_ = v6611
	var v6612 int32
	_ = v6612
	var v6613 int32
	_ = v6613
	var v6615 int32
	_ = v6615
	var v6618 int32
	_ = v6618
	var v6625 int32
	_ = v6625
	var v6626 int32
	_ = v6626
	var v6627 int32
	_ = v6627
	var v6639 int32
	_ = v6639
	var v6640 int32
	_ = v6640
	var v6643 int32
	_ = v6643
	var v6645 int32
	_ = v6645
	var v6646 int32
	_ = v6646
	var v6647 int32
	_ = v6647
	var v6650 int32
	_ = v6650
	var v6661 int32
	_ = v6661
	var v6666 int32
	_ = v6666
	var v6672 int32
	_ = v6672
	var v6675 int32
	_ = v6675
	var v6686 int32
	_ = v6686
	var v6687 int32
	_ = v6687
	var v6688 int32
	_ = v6688
	var v6689 int32
	_ = v6689
	var v6694 int32
	_ = v6694
	var v6699 int32
	_ = v6699
	var v6707 int32
	_ = v6707
	var v6712 int32
	_ = v6712
	var v6717 int32
	_ = v6717
	var v6725 int32
	_ = v6725
	var v6728 int32
	_ = v6728
	var v6731 int32
	_ = v6731
	var v6732 int32
	_ = v6732
	var v6736 int32
	_ = v6736
	var v6737 int32
	_ = v6737
	var v6739 int32
	_ = v6739
	var v6741 int32
	_ = v6741
	var v6746 int32
	_ = v6746
	var v6761 int32
	_ = v6761
	var v6763 int32
	_ = v6763
	var v6784 int32
	_ = v6784
	var v6785 int32
	_ = v6785
	var v6786 int32
	_ = v6786
	var v6792 int32
	_ = v6792
	var v6796 int32
	_ = v6796
	var v6797 int32
	_ = v6797
	var v6799 int32
	_ = v6799
	var v6804 int32
	_ = v6804
	var v6808 int32
	_ = v6808
	var v6809 int32
	_ = v6809
	var v6811 int32
	_ = v6811
	var v6812 int32
	_ = v6812
	var v6814 int32
	_ = v6814
	var v6830 int32
	_ = v6830
	var v6831 int32
	_ = v6831
	var v6838 int32
	_ = v6838
	var v6839 int32
	_ = v6839
	var v6840 int32
	_ = v6840
	var v6841 int32
	_ = v6841
	var v6842 int32
	_ = v6842
	var v6844 int32
	_ = v6844
	var v6850 int32
	_ = v6850
	var v6853 int32
	_ = v6853
	var v6854 int32
	_ = v6854
	var v6855 int32
	_ = v6855
	var v6860 int32
	_ = v6860
	var v6862 int32
	_ = v6862
	var v6865 int32
	_ = v6865
	var v6869 int32
	_ = v6869
	var v6870 int32
	_ = v6870
	var v6882 int32
	_ = v6882
	var v6887 int32
	_ = v6887
	var v6888 int32
	_ = v6888
	var v6891 int32
	_ = v6891
	var v6892 int32
	_ = v6892
	var v6898 int32
	_ = v6898
	var v6903 int32
	_ = v6903
	var v6906 int32
	_ = v6906
	var v6909 int32
	_ = v6909
	var v6910 int32
	_ = v6910
	var v6928 int32
	_ = v6928
	var v6946 int32
	_ = v6946
	var v6950 int32
	_ = v6950
	var v6953 int32
	_ = v6953
	var v6954 int32
	_ = v6954
	var v6960 int32
	_ = v6960
	var v6965 int32
	_ = v6965
	var v6969 int32
	_ = v6969
	var v6988 int32
	_ = v6988
	var v6990 int32
	_ = v6990
	var v7000 int32
	_ = v7000
	var v7001 int32
	_ = v7001
	var v7005 int32
	_ = v7005
	var v7009 int32
	_ = v7009
	var v7012 int32
	_ = v7012
	var v7013 int32
	_ = v7013
	var v7016 int32
	_ = v7016
	var v7020 int32
	_ = v7020
	var v7025 int32
	_ = v7025
	var v7027 int32
	_ = v7027
	var v7030 int32
	_ = v7030
	var v7040 int32
	_ = v7040
	var v7042 int32
	_ = v7042
	var v7050 int32
	_ = v7050
	var v7055 int32
	_ = v7055
	var v7057 int32
	_ = v7057
	var v7059 int32
	_ = v7059
	var v7064 int32
	_ = v7064
	var v7069 int32
	_ = v7069
	var v7074 int32
	_ = v7074
	v1 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v23 = F_CalculateShmemSize(m, v19+int32(8))
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
	v27 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v23
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_0), v19)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v39 = m.G0
	v41 = v39 - int32(352)
	m.G0 = v41
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	v49 = F___fstatat(m, int32(-100), v44, v41+int32(192), int32(0))
	mBase = m.M
	goto L10
L7:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_1), int32(211), int32(_a_F_CreateSharedMemoryAndSemaphores_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1])) = v874
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2])) = v874
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v874)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[3])) = v874 + v992
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v996 = m.G0
	v998 = v996 - int32(112)
	m.G0 = v998
	v1001 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	v1006 = F___fstatat(m, int32(-100), v1001, v998+int32(16), int32(0))
	mBase = m.M
	goto L252
L10:
	;
	if int32(0) <= v49 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[4]))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	if base.B2i32(v55 == int32(1))&base.B2i32(v59 != int32(2)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L248
	}
L14:
	;
	if v59 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L244
	}
L17:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v41)+280))
	v198 = base.I32_wrap_i64(v194)
	goto L56
L18:
	;
	v68 = int32(1)
	if base.Ui32(v68) < base.Ui32(v55-v68) {
		v123 = v23
		v124 = int32(-1)
		v125 = v1
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_3), int32(_a_F_CreateSharedMemoryAndSemaphores_4), int32(0), int32(1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L55
	}
L21:
	;
	if v124 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[6]))
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = v73 << (uint(int32(10)) % 32)
	goto L25
L24:
	;
	v77 = int32(_a_F_CreateSharedMemoryAndSemaphores_5)
	goto L25
L25:
	;
	if v77 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v88 = int32(_a_F_CreateSharedMemoryAndSemaphores_6) - base.I32_wrap_i64(base.I64_clz(base.I64_extend_i32_u(v77)-int64(1)))<<(uint(int32(26))%32)
	goto L28
L27:
	;
	v88 = int32(_a_F_CreateSharedMemoryAndSemaphores_6)
	goto L28
L28:
	;
	v89 = base.I32_rem_u_s(v23, v77)
	if v89 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v92 = v77 - v89
	goto L31
L30:
	;
	v92 = int32(0)
	goto L31
L31:
	;
	v93 = v92 + v23
	v94 = int32(-1)
	v95 = F_mmap(m, v93, v88, v94)
	mBase = m.M
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[4]))
	if base.B2i32(v99 != int32(2))|base.B2i32(v95 != v94) != 0 {
		v123 = v93
		v124 = v95
		v125 = v97
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v105 = int32(-1)
	v108 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v108 == int32(0) {
		v123 = v93
		v124 = v105
		v125 = v97
		goto L21
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+176)) = v93
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_7), v41+int32(176))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(627), int32(_a_F_CreateSharedMemoryAndSemaphores_9))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v123 = v93
	v124 = v105
	v125 = v97
	goto L21
L37:
	;
	v131 = int32(_a_F_CreateSharedMemoryAndSemaphores_4)
	goto L39
L38:
	;
	v131 = int32(_a_F_CreateSharedMemoryAndSemaphores_10)
	goto L39
L39:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_3), v131, int32(0), int32(1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v124 != int32(-1) {
		v147 = v123
		v148 = v124
		v149 = v125
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v148 == int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[4]))
	if v139 == int32(1) {
		v147 = v123
		v148 = v124
		v149 = v125
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v144 = F_mmap(m, v23, int32(33), int32(-1))
	mBase = m.M
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	v147 = v23
	v148 = v144
	v149 = v146
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = v149
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[8])) = v147
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[9])) = v148
	F_on_shmem_exit(m, int32(911), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L54
	}
L47:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v149 == int32(48) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v147
	F_errhint(m, int32(_a_F_CreateSharedMemoryAndSemaphores_12), v41+int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(663), int32(_a_F_CreateSharedMemoryAndSemaphores_9))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v190 = int32(40)
	v191 = v147
	goto L17
L55:
	;
	v190 = v23
	v191 = v23
	goto L17
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10]))
	if v217 != 0 {
		goto L64
	} else {
		goto L65
	}
L58:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
	if v913 == int32(0) {
		v198 = v899
		goto L56
	} else {
		goto L229
	}
L59:
	;
	v899 = v198 + int32(1)
	goto L58
L60:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L226
	}
L61:
	;
	v843 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v484)+16)) = v843
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = int32(679834894)
	*(*int32)(unsafe.Add(mBase, uint32(v484)+4)) = int32(42)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v41)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v484)+32)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v484)+24)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v484)+12)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v484)+8)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(12)))) = v484
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[11])) = v198
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[12])) = v484
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[9]))
	if v861 == v843 {
		goto L223
	} else {
		goto L224
	}
L62:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10]))
	if v510 != 0 {
		goto L151
	} else {
		goto L152
	}
L63:
	;
	if v288 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L64:
	;
	v221 = v217
	goto L67
L65:
	;
	goto L66
L66:
	;
	goto L73
L67:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v198 == v224 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L66
L69:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v288 = v226
	goto L63
L70:
	;
	goto L71
L71:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	if v227 != 0 {
		v221 = v227
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	v240 = int32(_a_F_CreateSharedMemoryAndSemaphores_13)
	goto L76
L76:
	;
	if base.Ui32(v240) < base.Ui32(v190) {
		v240 = v240 << (uint(int32(1)) % 32)
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v246 = F_emscripten_builtin_malloc(m, v240)
	mBase = m.M
	if v246 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(48)
	v288 = int32(-1)
	goto L63
L80:
	;
	goto L81
L81:
	;
	v254 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v254 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_emscripten_builtin_free(m, v246)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(48)
	v288 = int32(-1)
	goto L63
L83:
	;
	goto L84
L84:
	;
	v262 = int32(_a_F_CreateSharedMemoryAndSemaphores_14)
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[13])) = v264 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+16)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v254)+8)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v264
	v273 = int32(_a_F_CreateSharedMemoryAndSemaphores_15)
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+20)) = v274
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10])) = v254
	v288 = v264
	goto L63
L85:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	switch v292 - int32(2) {
	case 0, 18, 22:
		goto L62
	default:
		goto L90
	case 26:
		goto L91
	}
L86:
	;
	goto L87
L87:
	;
	F_on_shmem_exit(m, int32(912), v288)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L134
	}
L88:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(248), int32(_a_F_CreateSharedMemoryAndSemaphores_16))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L133
	}
L89:
	;
	F_errhint(m, v454, int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L132
	}
L90:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L128
	}
L91:
	;
	v295 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10]))
	if v301 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(28)
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L125
	}
L93:
	;
	if v372 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L94:
	;
	v305 = v301
	goto L97
L95:
	;
	goto L96
L96:
	;
	goto L103
L97:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v198 == v308 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L96
L99:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v372 = v310
	goto L93
L100:
	;
	goto L101
L101:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v305)+20))
	if v311 != 0 {
		v305 = v311
		goto L97
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	v324 = int32(_a_F_CreateSharedMemoryAndSemaphores_13)
	goto L106
L106:
	;
	if base.Ui32(v324) < base.Ui32(v295) {
		v324 = v324 << (uint(int32(1)) % 32)
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v330 = F_emscripten_builtin_malloc(m, v324)
	mBase = m.M
	if v330 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(48)
	v372 = int32(-1)
	goto L93
L110:
	;
	goto L111
L111:
	;
	v338 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v338 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_emscripten_builtin_free(m, v330)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(48)
	v372 = int32(-1)
	goto L93
L113:
	;
	goto L114
L114:
	;
	v346 = int32(_a_F_CreateSharedMemoryAndSemaphores_14)
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[13])) = v348 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v338)+16)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v338)+12)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v338)+8)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v338)+4)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v338))) = v348
	v357 = int32(_a_F_CreateSharedMemoryAndSemaphores_15)
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+20)) = v358
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10])) = v338
	v372 = v348
	goto L93
L115:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	if base.B2i32(base.Ui32(int32(24)) < base.Ui32(v376))|base.B2i32(int32(1)<<(uint(v376)%32)&int32(17825796) == int32(0)) != 0 {
		goto L92
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v386 = int32(0)
	v388 = F_pgl_shmctl(m, v372, v386, v386)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	goto L62
L119:
	;
	if int32(0) <= v388 {
		goto L92
	} else {
		goto L120
	}
L120:
	;
	v394 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v394 == int32(0) {
		goto L92
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+132)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+128)) = v372
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_17), v41+int32(128))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(209), int32(_a_F_CreateSharedMemoryAndSemaphores_16))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L92
L125:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_18), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+120)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+116)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v41)+112)) = v198
	F_errdetail(m, int32(_a_F_CreateSharedMemoryAndSemaphores_19), v41+int32(112))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v454 = int32(_a_F_CreateSharedMemoryAndSemaphores_20)
	goto L89
L128:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_18), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v198
	F_errdetail(m, int32(_a_F_CreateSharedMemoryAndSemaphores_19), v41+int32(32))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	switch v292 - int32(48) {
	case 0:
		v454 = int32(_a_F_CreateSharedMemoryAndSemaphores_21)
		goto L89
	default:
		goto L88
	case 3:
		goto L131
	}
L131:
	;
	v454 = int32(_a_F_CreateSharedMemoryAndSemaphores_22)
	goto L89
L132:
	;
	goto L88
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10]))
	if v469 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v484 == int32(-1) {
		goto L60
	} else {
		goto L145
	}
L136:
	;
	v471 = v469
	goto L139
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(28)
	v484 = int32(-1)
	goto L135
L139:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	if v288 == v472 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L138
L141:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v484 = v474
	goto L135
L142:
	;
	goto L143
L143:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v471)+20))
	if v475 != 0 {
		v471 = v475
		goto L139
	} else {
		goto L144
	}
L144:
	;
	goto L140
L145:
	;
	F_on_shmem_exit(m, int32(913), v484)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+160)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v41)+164)) = v288
	v493 = v41 + int32(288)
	v497 = F_pg_sprintf(m, v493, int32(_a_F_CreateSharedMemoryAndSemaphores_23), v41+int32(160))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_AddToDataDirLockFile(m, int32(7), v493)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	if v484 != 0 {
		goto L61
	} else {
		goto L149
	}
L149:
	;
	goto L62
L150:
	;
	if v581 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L151:
	;
	v514 = v510
	goto L154
L152:
	;
	goto L153
L153:
	;
	goto L161
L154:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v198 == v517 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L153
L156:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v581 = v519
	goto L150
L157:
	;
	goto L158
L158:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v514)+20))
	if v520 != 0 {
		v514 = v520
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L155
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(44)
	v581 = int32(-1)
	goto L150
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+188)) = int32(0)
	goto L59
L173:
	;
	goto L174
L174:
	;
	v588 = F_PGSharedMemoryAttach(m, v581, v41+int32(188))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L178
	}
L175:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v41)+188))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+16))
	if v638 != 0 {
		goto L188
	} else {
		goto L189
	}
L176:
	;
	v621 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L184
	}
L177:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	switch v588 - int32(2) {
	case 0:
		goto L176
	case 1:
		goto L59
	case 2:
		goto L175
	default:
		goto L177
	}
L179:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = v198
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_24), v41+int32(80))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = v607
	F_errhint(m, int32(_a_F_CreateSharedMemoryAndSemaphores_25), v41-int32(-64))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(802), int32(_a_F_CreateSharedMemoryAndSemaphores_26))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	if v621 == int32(0) {
		v899 = v198
		goto L58
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+100)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = v198
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_27), v41+int32(96))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(814), int32(_a_F_CreateSharedMemoryAndSemaphores_26))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v899 = v198
	goto L58
L188:
	;
	v639 = m.G0
	v641 = v639 - int32(48)
	m.G0 = v641
	v643 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v641)+44)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v641)+40)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v641)+36)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v641)+32)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v641)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v641)+24)) = v643
	v664 = F_dsm_impl_op(m, int32(1), v638, v643, v641+int32(36), v641+int32(44), v641+int32(28), int32(14))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v836 = int32(0)
	v838 = F_pgl_shmctl(m, v581, v836, v836)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L221
	}
L191:
	;
	if v664 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v666 = int32(2)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v641)+28))
	if base.Ui32(v667) < base.Ui32(int32(12)) {
		v780 = v666
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	m.G0 = v641 + int32(48)
	goto L190
L195:
	;
	v799 = F_dsm_impl_op(m, v780, v638, int32(0), v641+int32(36), v641+int32(44), v641+int32(28), int32(15))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L220
	}
L196:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v641)+44))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	if v671 != int32(-1706017486) {
		v780 = v666
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v670)+8))
	if base.Ui64(base.I64_extend_i32_u(v667)) < base.Ui64(base.I64_extend_i32_u(v675)*int64(24)+int64(12)) {
		v780 = v666
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v670)+4))
	if base.Ui32(v675) < base.Ui32(v682) {
		v780 = v666
		goto L195
	} else {
		goto L199
	}
L199:
	;
	if v682 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v692 = int32(0)
	goto L203
L201:
	;
	goto L202
L202:
	;
	v759 = int32(3)
	v762 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L216
	}
L203:
	;
	v705 = v670 + int32(12) + v692*int32(24)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+4))
	if v706 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L202
L205:
	;
	v741 = v692 + int32(1)
	if v741 != v682 {
		v692 = v741
		goto L203
	} else {
		goto L215
	}
L206:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	if v709&int32(1) != 0 {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v714 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	if v714 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641)+20)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v641)+16)) = v709
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_28), v641+int32(16))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v737 = F_dsm_impl_op(m, int32(3), v709, int32(0), v641+int32(32), v641+int32(40), v641+int32(24), int32(15))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), int32(294), int32(_a_F_CreateSharedMemoryAndSemaphores_30))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	goto L205
L215:
	;
	goto L204
L216:
	;
	if v762 == int32(0) {
		v780 = v759
		goto L195
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641))) = v638
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_31), v641)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), int32(304), int32(_a_F_CreateSharedMemoryAndSemaphores_30))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v780 = v759
	goto L195
L220:
	;
	goto L194
L221:
	;
	v899 = int32(base.Ui32(v838)>>(uint(int32(31))%32)) + v198
	goto L58
L222:
	;
	m.G0 = v41 + int32(352)
	goto L9
L223:
	;
	v874 = v484
	goto L222
L224:
	;
	goto L225
L225:
	;
	v864 = *(*int64)(unsafe.Add(mBase, uint32(v484)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+32)) = v864
	v866 = *(*int64)(unsafe.Add(mBase, uint32(v484)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+24)) = v866
	v868 = *(*int64)(unsafe.Add(mBase, uint32(v484)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+16)) = v868
	v870 = *(*int64)(unsafe.Add(mBase, uint32(v484)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v861)+8)) = v870
	v872 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
	*(*int64)(unsafe.Add(mBase, uint32(v861))) = v872
	v874 = v861
	goto L222
L226:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+148)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+144)) = v288
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_32), v41+int32(144))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(259), int32(_a_F_CreateSharedMemoryAndSemaphores_16))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10]))
	if v918 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if int32(0) <= v935 {
		v198 = v899
		goto L56
	} else {
		goto L239
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(28)
	v935 = int32(-1)
	goto L230
L232:
	;
	v922 = v918
	goto L233
L233:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v922)+12))
	if v913 != v923 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v935 = int32(0)
	goto L230
L235:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v922)+20))
	if v925 != 0 {
		v922 = v925
		goto L233
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	goto L234
L238:
	;
	goto L231
L239:
	;
	v940 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	if v940 == int32(0) {
		v198 = v899
		goto L56
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v913
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_33), v41+int32(48))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(839), int32(_a_F_CreateSharedMemoryAndSemaphores_26))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v198 = v899
	goto L56
L244:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_34), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(732), int32(_a_F_CreateSharedMemoryAndSemaphores_26))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v978 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v978
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_35), v41)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(718), int32(_a_F_CreateSharedMemoryAndSemaphores_26))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	if v1006 < int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1027 = F_mul_size(m, v995, int32(128))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L1
	} else {
		goto L260
	}
L256:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v998))) = v1016
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_35), v998)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_36), int32(210), int32(_a_F_CreateSharedMemoryAndSemaphores_37))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	v1029 = m.G0
	v1031 = v1029 - int32(16)
	m.G0 = v1031
	v1034 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2]))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+12))
	v1039 = (v1027 + int32(7)) & int32(-8)
	v1040 = v1035 + v1039
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+8))
	if base.Ui32(v1041) < base.Ui32(v1040) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_38), v7064)
	mBase = m.M
	v7069 = m.ExcPending
	if v7069 != 0 {
		goto L1
	} else {
		goto L1193
	}
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+12)) = v1040
	v1053 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	m.G0 = v1031 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[14])) = v1053 + v1035
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[15])) = v995
	v1063 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[16])) = v1063
	F_on_shmem_exit(m, int32(910), v1063)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	F_errcode(m, int32(_a_F_CreateSharedMemoryAndSemaphores_39))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1031))) = v1039
	v7064 = v1031
	goto L261
L267:
	;
	m.G0 = v998 + int32(112)
	v1072 = m.G0
	v1074 = v1072 - int32(16)
	m.G0 = v1074
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2]))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+12))
	v1080 = v1078 + int32(8)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+8))
	if base.Ui32(v1081) < base.Ui32(v1080) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1077)+12)) = v1080
	v1095 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	v1096 = v1095 + v1078
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[17])) = v1096
	v1098 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1096))) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v1077)+20)) = v1098
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1077)+12)) = (v1077+v1102+int32(127))&int32(-128) - v1077
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[18])) = v1098
	m.G0 = v1074 + int32(16)
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v1119 == int32(1) {
		goto L275
	} else {
		goto L276
	}
L271:
	;
	F_errcode(m, int32(_a_F_CreateSharedMemoryAndSemaphores_39))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1074))) = int32(8)
	v7064 = v1074
	goto L261
L273:
	;
	v1903 = m.G0
	v1905 = v1903 + int32(-64)
	m.G0 = v1905
	*(*int64)(unsafe.Add(mBase, uint32(v1905)+28)) = int64(257698037808)
	v1910 = int32(0)
	goto L357
L274:
	;
	if v1785 <= int32(0) {
		goto L273
	} else {
		goto L333
	}
L275:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[20]))
	v1785 = v1123
	goto L274
L276:
	;
	goto L277
L277:
	;
	v1125 = F_LWLockShmemSize(m)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v1127 = F_ShmemAlloc(m, v1125)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v1132 = (v1127 + int32(4)) & int32(-128)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21])) = v1132 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+124)) = int32(95)
	v1139 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[20]))
	if v1139 <= int32(0) {
		v1229 = v1098
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	v1246 = v1245
	v1247 = int32(0)
	goto L292
L281:
	;
	v1143 = v1139 & int32(3)
	v1145 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[22]))
	v1146 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1139) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1151 = v1146
	v1153 = v1098
	v1157 = v1
	goto L285
L283:
	;
	v1185 = v1146
	v1187 = v1098
	goto L284
L284:
	;
	v1201 = v1185
	v1203 = v1187
	v1204 = v1098
	goto L289
L285:
	;
	v1169 = v1145 + v1151*int32(68)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+268))
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+200))
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+132))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+64))
	v1177 = v1170 + (v1171 + (v1172 + (v1173 + v1153)))
	v1178 = int32(4)
	v1179 = v1151 + v1178
	v1181 = v1157 + v1178
	if v1181 != v1139&int32(2147483644) {
		v1151 = v1179
		v1153 = v1177
		v1157 = v1181
		goto L285
	} else {
		goto L287
	}
L286:
	;
	if v1143 == int32(0) {
		v1229 = v1177
		goto L280
	} else {
		goto L288
	}
L287:
	;
	goto L286
L288:
	;
	v1185 = v1179
	v1187 = v1177
	goto L284
L289:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1145+v1201*int32(68))+64))
	v1221 = v1220 + v1203
	v1222 = int32(1)
	v1225 = v1204 + v1222
	if v1225 != v1143 {
		v1201 = v1201 + v1222
		v1203 = v1221
		v1204 = v1225
		goto L289
	} else {
		goto L291
	}
L290:
	;
	v1229 = v1221
	goto L280
L291:
	;
	goto L290
L292:
	;
	v1262 = int32(1073741824)
	*(*int32)(unsafe.Add(mBase, uint32(v1246)+4)) = v1262
	*(*int32)(unsafe.Add(mBase, uint32(v1246)+132)) = v1262
	*(*uint16)(unsafe.Add(mBase, uint32(v1246))) = uint16(v1247)
	v1267 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1246)+8)) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v1246)+136)) = v1267
	*(*int64)(unsafe.Add(mBase, uint32(v1246)+264)) = v1267
	v1274 = v1247 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1246)+128)) = uint16(v1274)
	v1277 = v1247 + int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1246)+256)) = uint16(v1277)
	*(*int32)(unsafe.Add(mBase, uint32(v1246)+260)) = v1262
	v1284 = v1247 + int32(3)
	if v1284 != int32(54) {
		v1246 = v1246 + int32(384)
		v1247 = v1284
		goto L292
	} else {
		goto L294
	}
L293:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	v1292 = v1289 + int32(_a_F_CreateSharedMemoryAndSemaphores_40)
	v1293 = int32(0)
	goto L295
L294:
	;
	goto L293
L295:
	;
	v1308 = int32(1073741824)
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+4)) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+132)) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+260)) = v1308
	v1314 = int32(66)
	*(*uint16)(unsafe.Add(mBase, uint32(v1292))) = uint16(v1314)
	v1316 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+8)) = v1316
	*(*uint16)(unsafe.Add(mBase, uint32(v1292)+128)) = uint16(v1314)
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+136)) = v1316
	*(*uint16)(unsafe.Add(mBase, uint32(v1292)+256)) = uint16(v1314)
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+264)) = v1316
	*(*uint16)(unsafe.Add(mBase, uint32(v1292)+384)) = uint16(v1314)
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+392)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+388)) = v1308
	v1335 = v1293 + int32(4)
	if v1335 != int32(128) {
		v1292 = v1292 + int32(512)
		v1293 = v1335
		goto L295
	} else {
		goto L297
	}
L296:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	v1340 = int32(1073741824)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[23]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[24]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[25]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[26]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[27]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[28]))) = v1340
	v1352 = int32(67)
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[29]))) = uint16(v1352)
	v1354 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[30]))) = v1354
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[31]))) = uint16(v1352)
	*(*int64)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[32]))) = v1354
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[33]))) = uint16(v1352)
	*(*int64)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[34]))) = v1354
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[35]))) = uint16(v1352)
	*(*int64)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[36]))) = v1354
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[37]))) = uint16(v1352)
	*(*int64)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[38]))) = v1354
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[39]))) = uint16(v1352)
	*(*int64)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[40]))) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[41]))) = v1340
	v1378 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[42]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[43]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[44]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[45]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[46]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[47]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[48]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[49]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[50]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[51]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[52]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[53]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[54]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[55]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[56]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[57]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[58]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[59]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[60]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[61]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[62]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[63]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[64]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[65]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[66]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[67]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[68]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[69]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[70]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[71]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[72]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[73]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[74]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[75]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[76]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[77]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[78]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[79]))) = uint16(v1352)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[80]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[81]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[82]))) = v1378
	v1460 = int32(68)
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[83]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[84]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[85]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[86]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[87]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[88]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[89]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[90]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[91]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[92]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[93]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[94]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[95]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[96]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[97]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[98]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[99]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[100]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[101]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[102]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[103]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[104]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[105]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[106]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[107]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[108]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[109]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[110]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[111]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[112]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[113]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[114]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[115]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[116]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[117]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[118]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[119]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[120]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[121]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[122]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[123]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[124]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[125]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[126]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[127]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[128]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[129]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[130]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[131]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[132]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[133]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[134]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[135]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[136]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[137]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[138]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[139]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[140]))) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[141]))) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[142]))) = v1378
	*(*uint16)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[143]))) = uint16(v1460)
	*(*int32)(unsafe.Add(mBase, uint32(v1339)+uint32(_c_F_CreateSharedMemoryAndSemaphores[144]))) = v1378
	v1585 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[20]))
	if v1585 <= int32(0) {
		goto L273
	} else {
		goto L298
	}
L297:
	;
	goto L296
L298:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	v1592 = v1590 + int32(_a_F_CreateSharedMemoryAndSemaphores_41)
	v1595 = v1592 + v1229<<(uint(int32(7))%32)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[145])) = v1595
	v1601 = v1592
	v1606 = int32(0)
	v1607 = v1595 + v1585<<(uint(int32(3))%32)
	goto L299
L299:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[145]))
	v1620 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[22]))
	v1623 = v1620 + v1606*int32(68)
	v1624 = F_strlen(m, v1623)
	mBase = m.M
	if (v1623^v1607)&int32(3) != 0 {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	v1785 = v1782
	goto L274
L301:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	v1702 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[17]))
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1702)))
	*(*int32)(unsafe.Add(mBase, uint32(v1702))) = int32(1)
	if v1703 != 0 {
		goto L322
	} else {
		goto L323
	}
L302:
	;
	goto L301
L303:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1679))) = uint8(v1678)
	if v1678&int32(255) == int32(0) {
		goto L302
	} else {
		goto L318
	}
L304:
	;
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1623))))
	v1677 = v1623
	v1678 = v1630
	v1679 = v1607
	goto L303
L305:
	;
	goto L306
L306:
	;
	if v1623&int32(3) != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1634 = v1623
	v1636 = v1607
	goto L310
L308:
	;
	v1648 = v1623
	v1650 = v1607
	goto L309
L309:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1648)))
	v1655 = int32(-2139062144)
	if (int32(16843008)-v1652|v1652)&v1655 != v1655 {
		v1677 = v1648
		v1678 = v1652
		v1679 = v1650
		goto L303
	} else {
		goto L314
	}
L310:
	;
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1636))) = uint8(v1637)
	if v1637 == int32(0) {
		goto L302
	} else {
		goto L312
	}
L311:
	;
	v1648 = v1644
	v1650 = v1642
	goto L309
L312:
	;
	v1641 = int32(1)
	v1642 = v1636 + v1641
	v1644 = v1634 + v1641
	if v1644&int32(3) != 0 {
		v1634 = v1644
		v1636 = v1642
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	v1660 = v1648
	v1661 = v1652
	v1662 = v1650
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1662))) = v1661
	v1664 = int32(4)
	v1665 = v1662 + v1664
	v1667 = v1660 + v1664
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1660)+4))
	v1672 = int32(-2139062144)
	if (int32(16843008)-v1669|v1669)&v1672 == v1672 {
		v1660 = v1667
		v1661 = v1669
		v1662 = v1665
		goto L315
	} else {
		goto L317
	}
L316:
	;
	v1677 = v1667
	v1678 = v1669
	v1679 = v1665
	goto L303
L317:
	;
	goto L316
L318:
	;
	v1686 = v1677
	v1688 = v1679
	goto L319
L319:
	;
	v1689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1686)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1688)+1)) = uint8(v1689)
	v1691 = int32(1)
	if v1689 != 0 {
		v1686 = v1686 + v1691
		v1688 = v1688 + v1691
		goto L319
	} else {
		goto L321
	}
L320:
	;
	goto L302
L321:
	;
	goto L320
L322:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[17]))
	F_s_lock(m, v1707, int32(_a_F_CreateSharedMemoryAndSemaphores_42), int32(622), int32(_a_F_CreateSharedMemoryAndSemaphores_43))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	v1714 = v1700 - int32(4)
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1714)))
	*(*int32)(unsafe.Add(mBase, uint32(v1714))) = v1715 + int32(1)
	v1719 = int32(0)
	v1721 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v1721))) = v1719
	v1726 = v1618 + v1606<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1726)+4)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v1726))) = v1715
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+64))
	if v1719 < v1729 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	goto L324
L326:
	;
	v1732 = v1601
	v1733 = v1719
	goto L329
L327:
	;
	v1760 = v1601
	goto L328
L328:
	;
	v1777 = int32(1)
	v1780 = v1606 + v1777
	v1782 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[20]))
	if v1780 < v1782 {
		v1601 = v1760
		v1606 = v1780
		v1607 = v1624 + v1607 + v1777
		goto L299
	} else {
		goto L332
	}
L329:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1726)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1732))) = uint16(v1748)
	*(*int32)(unsafe.Add(mBase, uint32(v1732)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v1732)+8)) = int64(-1)
	v1755 = v1732 + int32(128)
	v1757 = v1733 + int32(1)
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+64))
	if v1757 < v1758 {
		v1732 = v1755
		v1733 = v1757
		goto L329
	} else {
		goto L331
	}
L330:
	;
	v1760 = v1755
	goto L328
L331:
	;
	goto L330
L332:
	;
	goto L300
L333:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[145]))
	v1805 = int32(0)
	v1806 = v1785
	v1808 = v1803
	goto L334
L334:
	;
	v1823 = v1808 + v1805<<(uint(int32(3))%32)
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1823)))
	if int32(95) <= v1824 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	goto L273
L336:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+4))
	v1829 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[146]))
	v1831 = v1824 - int32(95)
	v1833 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[147]))
	if v1833 <= v1831 {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	v1878 = v1806
	v1880 = v1808
	goto L338
L338:
	;
	v1885 = v1805 + int32(1)
	if v1885 < v1878 {
		v1805 = v1885
		v1806 = v1878
		v1808 = v1880
		goto L334
	} else {
		goto L354
	}
L339:
	;
	v1837 = int32(102)
	if base.Ui32(v1824) <= base.Ui32(v1837) {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	v1867 = v1829
	goto L341
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1867+v1831<<(uint(int32(2))%32)))) = v1827
	v1875 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[145]))
	v1877 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[20]))
	v1878 = v1877
	v1880 = v1875
	goto L338
L342:
	;
	v1840 = v1837
	goto L344
L343:
	;
	v1840 = v1824
	goto L344
L344:
	;
	v1842 = v1840 - int32(94)
	if v1842&(v1840-int32(95)) != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1849 = int32(1) << (uint(int32(32)-base.I32_clz(v1842)) % 32)
	goto L347
L346:
	;
	v1849 = v1842
	goto L347
L347:
	;
	v1851 = v1849 << (uint(int32(2)) % 32)
	if v1829 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[147])) = v1849
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[146])) = v1862
	v1867 = v1862
	goto L341
L349:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[148]))
	v1856 = F_MemoryContextAllocZero(m, v1855, v1851)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L1
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1860 = F_repalloc0(m, v1829, v1833<<(uint(int32(2))%32), v1851)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L353
	}
L352:
	;
	v1862 = v1856
	goto L348
L353:
	;
	v1862 = v1860
	goto L348
L354:
	;
	goto L335
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1905)+48)) = int32(1108)
	*(*int32)(unsafe.Add(mBase, uint32(v1905)+20)) = v1933
	*(*int32)(unsafe.Add(mBase, uint32(v1905)+24)) = v1933
	v1945 = v1903 + int32(-52)
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+8))
	goto L362
L357:
	;
	goto L358
L358:
	;
	v1933 = int32(256)
	goto L359
L359:
	;
	if v1933 < int32(1)<<(uint(v1910-base.I32_clz(int32(base.Ui32(int32(-1)<<(uint(v1910-base.I32_clz(int32(63)))%32)^int32(-1))>>(uint(int32(8))%32))))%32) {
		v1933 = v1933 << (uint(int32(1)) % 32)
		goto L359
	} else {
		goto L361
	}
L360:
	;
	goto L355
L361:
	;
	goto L360
L362:
	;
	v1953 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_44), v1946<<(uint(int32(2))%32)+int32(432), v1903+int32(-1))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1905)+56)) = v1953
	v1961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1905)+63)))
	if v1961 != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1962 = int32(_a_F_CreateSharedMemoryAndSemaphores_45)
	goto L366
L365:
	;
	v1962 = int32(2588)
	goto L366
L366:
	;
	v1963 = F_hash_create(m, int32(_a_F_CreateSharedMemoryAndSemaphores_44), int32(64), v1945, v1962)
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[18])) = v1963
	m.G0 = v1905 - int32(-64)
	v1969 = m.G0
	v1971 = v1969 - int32(16)
	m.G0 = v1971
	v1974 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[149]))
	v1976 = v1974 << (uint(int32(20)) % 32)
	if v1976 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2016 = int32(16)
	m.G0 = v1971 + v2016
	v2019 = m.G0
	v2021 = v2019 - v2016
	m.G0 = v2021
	v2028 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_46), int32(8), v2021+int32(15))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L1
	} else {
		goto L377
	}
L369:
	;
	v1983 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_47), v1976, v1971+int32(15))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[150])) = v1983
	v1986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1971)+15)))
	if v1986 != 0 {
		goto L368
	} else {
		goto L371
	}
L371:
	;
	v1987 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1983)+4)) = v1987
	*(*int64)(unsafe.Add(mBase, uint32(v1983)+12)) = v1987
	*(*int64)(unsafe.Add(mBase, uint32(v1983)+20)) = v1987
	v1993 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1983)+28)) = v1993
	v1995 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1983)+32)) = uint8(v1995)
	if v1983 != 0 {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	v2008 = int32(1)
	F_FreePageManagerPut(m, v1983, v2008, int32(base.Ui32(v1976)>>(uint(int32(12))%32))-v2008)
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L1
	} else {
		goto L376
	}
L373:
	;
	v2001 = v1983 - v1983 + v1995
	goto L375
L374:
	;
	v2001 = v1993
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1983))) = v2001
	base.MemoryFill(m, v1983+int32(36), int32(0), int32(516))
	goto L372
L376:
	;
	goto L368
L377:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[151])) = v2028
	v2031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2021)+15)))
	if v2031 == int32(0) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2028))) = int64(0)
	goto L380
L379:
	;
	goto L380
L380:
	;
	v2036 = int32(16)
	m.G0 = v2021 + v2036
	v2039 = m.G0
	v2041 = v2039 - v2036
	m.G0 = v2041
	v2048 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_48), int32(72), v2041+int32(15))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[152])) = v2048
	v2052 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v2052 == int32(0) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	base.MemoryFill(m, v2048, int32(0), int32(72))
	goto L384
L383:
	;
	goto L384
L384:
	;
	v2058 = int32(16)
	m.G0 = v2041 + v2058
	v2061 = int32(0)
	v2063 = m.G0
	v2065 = v2063 - v2058
	m.G0 = v2065
	v2069 = F_XLOGShmemSize(m)
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	v2073 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_49), v2069, v2065+int32(14))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[153])) = v2073
	v2076 = int32(_a_F_CreateSharedMemoryAndSemaphores_50)
	v2077 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[154]))
	v2083 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_51), int32(296), v2065+int32(15))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[154])) = v2083
	v2087 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[153]))
	v2088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+15)))
	if v2088 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	v2399 = int32(16)
	m.G0 = v2065 + v2399
	v2402 = m.G0
	v2404 = v2402 - v2399
	m.G0 = v2404
	v2411 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_52), int32(72), v2404+int32(15))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L1
	} else {
		goto L423
	}
L389:
	;
	base.MemoryFill(m, v2087, int32(0), int32(448))
	if v2077 != 0 {
		goto L396
	} else {
		goto L397
	}
L390:
	;
	v2091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+14)))
	if v2091&int32(1) == int32(0) {
		goto L389
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2087)+176))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155])) = v2097
	if v2077 == int32(0) {
		goto L388
	} else {
		goto L394
	}
L393:
	;
	goto L392
L394:
	;
	F_pfree(m, v2077)
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	goto L388
L396:
	;
	base.MemoryCopy(m, v2083, v2077, int32(296))
	F_pfree(m, v2077)
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L1
	} else {
		goto L399
	}
L397:
	;
	v2112 = v2087
	goto L398
L398:
	;
	v2114 = v2112 + int32(448)
	*(*int32)(unsafe.Add(mBase, uint32(v2112)+300)) = v2114
	v2117 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[156]))
	if v2117 <= int32(0) {
		goto L400
	} else {
		goto L401
	}
L399:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[153]))
	v2112 = v2111
	goto L398
L400:
	;
	v2231 = (v2114 + v2117<<(uint(int32(3))%32)) & int32(-128)
	v2233 = v2231 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v2112)+176)) = v2233
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155])) = v2233
	v2237 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2233))) = uint16(v2237)
	*(*int32)(unsafe.Add(mBase, uint32(v2233)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2233)+8)) = int64(-1)
	goto L412
L401:
	;
	v2124 = v2117 & int32(3)
	v2125 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v2117) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v2130 = v2125
	v2136 = v2061
	goto L405
L403:
	;
	v2171 = v2125
	goto L404
L404:
	;
	v2187 = v2171
	v2192 = v2061
	goto L409
L405:
	;
	v2147 = v2130 << (uint(int32(3)) % 32)
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+300))
	v2150 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2147+v2148))) = v2150
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+300))
	*(*int64)(unsafe.Add(mBase, uint32(v2152+v2147)+8)) = v2150
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+300))
	*(*int64)(unsafe.Add(mBase, uint32(v2156+v2147)+16)) = v2150
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+300))
	*(*int64)(unsafe.Add(mBase, uint32(v2160+v2147)+24)) = v2150
	v2164 = int32(4)
	v2165 = v2130 + v2164
	v2167 = v2136 + v2164
	if v2167 != v2117&int32(-4) {
		v2130 = v2165
		v2136 = v2167
		goto L405
	} else {
		goto L407
	}
L406:
	;
	if v2124 == int32(0) {
		goto L400
	} else {
		goto L408
	}
L407:
	;
	goto L406
L408:
	;
	v2171 = v2165
	goto L404
L409:
	;
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2112)+300))
	*(*int64)(unsafe.Add(mBase, uint32(v2203+v2187<<(uint(int32(3))%32)))) = int64(0)
	v2209 = int32(1)
	v2212 = v2192 + v2209
	if v2212 != v2124 {
		v2187 = v2187 + v2209
		v2192 = v2212
		goto L409
	} else {
		goto L411
	}
L410:
	;
	goto L400
L411:
	;
	goto L410
L412:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2245 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2244)+24)) = v2245
	*(*int64)(unsafe.Add(mBase, uint32(v2244)+16)) = v2245
	v2250 = v2244 + int32(128)
	v2251 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2250))) = uint16(v2251)
	*(*int32)(unsafe.Add(mBase, uint32(v2250)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2250)+8)) = int64(-1)
	goto L413
L413:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2259 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2258)+152)) = v2259
	*(*int64)(unsafe.Add(mBase, uint32(v2258)+144)) = v2259
	v2264 = v2258 + int32(256)
	v2265 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2264))) = uint16(v2265)
	*(*int32)(unsafe.Add(mBase, uint32(v2264)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2264)+8)) = int64(-1)
	goto L414
L414:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2273 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2272)+280)) = v2273
	*(*int64)(unsafe.Add(mBase, uint32(v2272)+272)) = v2273
	v2278 = v2272 + int32(384)
	v2279 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2278))) = uint16(v2279)
	*(*int32)(unsafe.Add(mBase, uint32(v2278)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2278)+8)) = int64(-1)
	goto L415
L415:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2287 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2286)+408)) = v2287
	*(*int64)(unsafe.Add(mBase, uint32(v2286)+400)) = v2287
	v2292 = v2286 + int32(512)
	v2293 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2292))) = uint16(v2293)
	*(*int32)(unsafe.Add(mBase, uint32(v2292)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2292)+8)) = int64(-1)
	goto L416
L416:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2301 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2300)+536)) = v2301
	*(*int64)(unsafe.Add(mBase, uint32(v2300)+528)) = v2301
	v2306 = v2300 + int32(640)
	v2307 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2306))) = uint16(v2307)
	*(*int32)(unsafe.Add(mBase, uint32(v2306)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2306)+8)) = int64(-1)
	goto L417
L417:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2315 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2314)+664)) = v2315
	*(*int64)(unsafe.Add(mBase, uint32(v2314)+656)) = v2315
	v2320 = v2314 + int32(768)
	v2321 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2320))) = uint16(v2321)
	*(*int32)(unsafe.Add(mBase, uint32(v2320)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2320)+8)) = int64(-1)
	goto L418
L418:
	;
	v2328 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2329 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2328)+792)) = v2329
	*(*int64)(unsafe.Add(mBase, uint32(v2328)+784)) = v2329
	v2334 = v2328 + int32(896)
	v2335 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v2334))) = uint16(v2335)
	*(*int32)(unsafe.Add(mBase, uint32(v2334)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v2334)+8)) = int64(-1)
	goto L419
L419:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[155]))
	v2343 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2342)+920)) = v2343
	*(*int64)(unsafe.Add(mBase, uint32(v2342)+912)) = v2343
	v2348 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[153]))
	v2352 = (v2231 + int32(_a_F_CreateSharedMemoryAndSemaphores_53)) & int32(-8192)
	*(*int32)(unsafe.Add(mBase, uint32(v2348)+296)) = v2352
	v2355 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[156]))
	v2357 = v2355 << (uint(int32(13)) % 32)
	if v2357 != 0 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	base.MemoryFill(m, v2352, int32(0), v2357)
	goto L422
L421:
	;
	goto L422
L422:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[156]))
	v2363 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[153]))
	v2364 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2363)+264)) = v2364
	*(*int64)(unsafe.Add(mBase, uint32(v2363)+272)) = v2364
	*(*int64)(unsafe.Add(mBase, uint32(v2363)+280)) = v2364
	v2370 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2363)+320)) = uint16(v2370)
	*(*int32)(unsafe.Add(mBase, uint32(v2363)+316)) = v2370
	*(*int32)(unsafe.Add(mBase, uint32(v2363)+440)) = v2370
	*(*int32)(unsafe.Add(mBase, uint32(v2363))) = v2370
	*(*int32)(unsafe.Add(mBase, uint32(v2363)+304)) = v2361 - int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v2363)+240)) = v2364
	goto L388
L423:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[157])) = v2411
	v2414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2404)+15)))
	if v2414 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v2420 = m.G0
	v2421 = int32(16)
	v2422 = v2420 - v2421
	m.G0 = v2422
	F_gettimeofday(m, v2422)
	mBase = m.M
	v2425 = *(*int64)(unsafe.Add(mBase, uint32(v2422)))
	v2426 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2422)+8)))
	m.G0 = v2422 + v2421
	goto L427
L425:
	;
	goto L426
L426:
	;
	v2451 = int32(16)
	m.G0 = v2404 + v2451
	v2454 = m.G0
	v2456 = v2454 - v2451
	m.G0 = v2456
	v2463 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_54), int32(104), v2456+int32(15))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L1
	} else {
		goto L428
	}
L427:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2411))) = v2426 + v2425*int64(1000000) - int64(946684800000000)
	v2437 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[157]))
	v2438 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2437)+8)) = v2438
	*(*int64)(unsafe.Add(mBase, uint32(v2437)+16)) = v2438
	*(*int64)(unsafe.Add(mBase, uint32(v2437)+24)) = v2438
	*(*int64)(unsafe.Add(mBase, uint32(v2437)+32)) = v2438
	*(*int64)(unsafe.Add(mBase, uint32(v2437)+40)) = v2438
	*(*int64)(unsafe.Add(mBase, uint32(v2437)+48)) = v2438
	goto L426
L428:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[158])) = v2463
	v2466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2456)+15)))
	if v2466 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v2469 = int32(0)
	base.MemoryFill(m, v2463, v2469, int32(104))
	*(*int32)(unsafe.Add(mBase, uint32(v2463)+96)) = v2469
	v2475 = v2463 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2475)+12)) = v2469
	*(*int64)(unsafe.Add(mBase, uint32(v2475))) = int64(0)
	v2480 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2475)+8)) = uint8(v2480)
	goto L432
L430:
	;
	goto L431
L431:
	;
	m.G0 = v2456 + int32(16)
	v2493 = m.G0
	v2495 = v2493 - int32(48)
	m.G0 = v2495
	v2500 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[159]))
	if v2500 != 0 {
		v2561 = v2500
		goto L436
	} else {
		goto L437
	}
L432:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[158]))
	v2485 = v2483 + int32(84)
	*(*int32)(unsafe.Add(mBase, uint32(v2485)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v2485))) = int64(-4294967296)
	goto L433
L433:
	;
	goto L431
L434:
	;
	F_SimpleLruInit(m, int32(_a_F_CreateSharedMemoryAndSemaphores_55), int32(_a_F_CreateSharedMemoryAndSemaphores_56), v2575, int32(1024), int32(_a_F_CreateSharedMemoryAndSemaphores_57), int32(54), int32(92), int32(1), int32(0))
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L1
	} else {
		goto L463
	}
L435:
	;
	v2567 = int32(16)
	if v2565 <= v2567 {
		goto L457
	} else {
		goto L458
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[160])) = int32(287)
	v2565 = v2561
	goto L435
L437:
	;
	v2503 = int32(16)
	v2505 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2507 = base.I32_div_s(v2505, int32(512))
	v2509 = base.I32_rem_s(v2507, v2503)
	v2510 = v2507 - v2509
	if v2510 <= v2503 {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2495))) = v2516
	v2522 = F_pg_snprintf(m, v2495+int32(16), int32(32), int32(_a_F_CreateSharedMemoryAndSemaphores_58), v2495)
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L1
	} else {
		goto L445
	}
L439:
	;
	v2513 = v2503
	goto L441
L440:
	;
	v2513 = v2510
	goto L441
L441:
	;
	if int32(1024) < v2513 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v2516 = int32(1024)
	goto L444
L443:
	;
	v2516 = v2513
	goto L444
L444:
	;
	goto L438
L445:
	;
	v2527 = int32(1)
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_59), v2495+int32(16), v2527, v2527)
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[159]))
	if v2532 != 0 {
		v2561 = v2532
		goto L436
	} else {
		goto L447
	}
L447:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_59), v2495+int32(16), int32(1), int32(10))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[160])) = int32(287)
	v2544 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[159]))
	if v2544 != 0 {
		v2565 = v2544
		goto L435
	} else {
		goto L449
	}
L449:
	;
	v2547 = int32(16)
	v2549 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2551 = base.I32_div_s(v2549, int32(512))
	v2553 = base.I32_rem_s(v2551, v2547)
	v2554 = v2551 - v2553
	if v2554 <= v2547 {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	v2575 = v2560
	goto L434
L451:
	;
	v2557 = v2547
	goto L453
L452:
	;
	v2557 = v2554
	goto L453
L453:
	;
	if int32(1024) < v2557 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v2560 = int32(1024)
	goto L456
L455:
	;
	v2560 = v2557
	goto L456
L456:
	;
	goto L450
L457:
	;
	v2570 = v2567
	goto L459
L458:
	;
	v2570 = v2565
	goto L459
L459:
	;
	if int32(_a_F_CreateSharedMemoryAndSemaphores_13) <= v2570 {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v2573 = int32(_a_F_CreateSharedMemoryAndSemaphores_13)
	goto L462
L461:
	;
	v2573 = v2570
	goto L462
L462:
	;
	v2575 = v2573
	goto L434
L463:
	;
	v2584 = int32(48)
	m.G0 = v2495 + v2584
	v2587 = m.G0
	v2589 = v2587 - v2584
	m.G0 = v2589
	v2594 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[162]))
	if v2594 != 0 {
		v2651 = v2594
		goto L466
	} else {
		goto L467
	}
L464:
	;
	v2669 = int32(0)
	F_SimpleLruInit(m, int32(_a_F_CreateSharedMemoryAndSemaphores_60), int32(_a_F_CreateSharedMemoryAndSemaphores_61), v2668, v2669, int32(_a_F_CreateSharedMemoryAndSemaphores_62), int32(55), int32(86), int32(2), v2669)
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L1
	} else {
		goto L493
	}
L465:
	;
	v2659 = int32(16)
	if v2656 <= v2659 {
		goto L487
	} else {
		goto L488
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[163])) = int32(289)
	v2656 = v2651
	goto L465
L467:
	;
	v2597 = int32(16)
	v2599 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2601 = base.I32_div_s(v2599, int32(512))
	v2603 = base.I32_rem_s(v2601, v2597)
	v2604 = v2601 - v2603
	if v2604 <= v2597 {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2589))) = v2610
	v2613 = v2589 + int32(16)
	v2616 = F_pg_snprintf(m, v2613, int32(32), int32(_a_F_CreateSharedMemoryAndSemaphores_58), v2589)
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		goto L1
	} else {
		goto L475
	}
L469:
	;
	v2607 = v2597
	goto L471
L470:
	;
	v2607 = v2604
	goto L471
L471:
	;
	if int32(1024) < v2607 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2610 = int32(1024)
	goto L474
L473:
	;
	v2610 = v2607
	goto L474
L474:
	;
	goto L468
L475:
	;
	v2619 = int32(1)
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_63), v2613, v2619, v2619)
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[162]))
	if v2624 != 0 {
		v2651 = v2624
		goto L466
	} else {
		goto L477
	}
L477:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_63), v2613, int32(1), int32(10))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[163])) = int32(289)
	v2634 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[162]))
	if v2634 != 0 {
		v2656 = v2634
		goto L465
	} else {
		goto L479
	}
L479:
	;
	v2637 = int32(16)
	v2639 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2641 = base.I32_div_s(v2639, int32(512))
	v2643 = base.I32_rem_s(v2641, v2637)
	v2644 = v2641 - v2643
	if v2644 <= v2637 {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v2668 = v2650
	goto L464
L481:
	;
	v2647 = v2637
	goto L483
L482:
	;
	v2647 = v2644
	goto L483
L483:
	;
	if int32(1024) < v2647 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2650 = int32(1024)
	goto L486
L485:
	;
	v2650 = v2647
	goto L486
L486:
	;
	goto L480
L487:
	;
	v2662 = v2659
	goto L489
L488:
	;
	v2662 = v2656
	goto L489
L489:
	;
	if int32(_a_F_CreateSharedMemoryAndSemaphores_64) <= v2662 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v2665 = int32(_a_F_CreateSharedMemoryAndSemaphores_64)
	goto L492
L491:
	;
	v2665 = v2662
	goto L492
L492:
	;
	v2668 = v2665
	goto L464
L493:
	;
	v2682 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_65), int32(32), v2589+int32(16))
	mBase = m.M
	v2683 = m.ExcPending
	if v2683 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[164])) = v2682
	v2686 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v2686 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v2689 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2682)+24)) = uint8(v2689)
	*(*uint16)(unsafe.Add(mBase, uint32(v2682)+16)) = uint16(v2689)
	*(*int64)(unsafe.Add(mBase, uint32(v2682)+8)) = int64(-9223372036854775807 - 1)
	*(*int32)(unsafe.Add(mBase, uint32(v2682))) = v2689
	goto L497
L496:
	;
	goto L497
L497:
	;
	v2697 = int32(48)
	m.G0 = v2589 + v2697
	v2700 = m.G0
	v2702 = v2700 - v2697
	m.G0 = v2702
	v2707 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[165]))
	if v2707 != 0 {
		v2764 = v2707
		goto L500
	} else {
		goto L501
	}
L498:
	;
	v2782 = int32(0)
	F_SimpleLruInit(m, int32(_a_F_CreateSharedMemoryAndSemaphores_66), int32(_a_F_CreateSharedMemoryAndSemaphores_67), v2781, v2782, int32(_a_F_CreateSharedMemoryAndSemaphores_68), int32(56), int32(91), int32(5), v2782)
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L1
	} else {
		goto L527
	}
L499:
	;
	v2772 = int32(16)
	if v2769 <= v2772 {
		goto L521
	} else {
		goto L522
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[166])) = int32(392)
	v2769 = v2764
	goto L499
L501:
	;
	v2710 = int32(16)
	v2712 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2714 = base.I32_div_s(v2712, int32(512))
	v2716 = base.I32_rem_s(v2714, v2710)
	v2717 = v2714 - v2716
	if v2717 <= v2710 {
		goto L503
	} else {
		goto L504
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2702))) = v2723
	v2726 = v2702 + int32(16)
	v2729 = F_pg_snprintf(m, v2726, int32(32), int32(_a_F_CreateSharedMemoryAndSemaphores_58), v2702)
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L1
	} else {
		goto L509
	}
L503:
	;
	v2720 = v2710
	goto L505
L504:
	;
	v2720 = v2717
	goto L505
L505:
	;
	if int32(1024) < v2720 {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v2723 = int32(1024)
	goto L508
L507:
	;
	v2723 = v2720
	goto L508
L508:
	;
	goto L502
L509:
	;
	v2732 = int32(1)
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_69), v2726, v2732, v2732)
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	v2737 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[165]))
	if v2737 != 0 {
		v2764 = v2737
		goto L500
	} else {
		goto L511
	}
L511:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_69), v2726, int32(1), int32(10))
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[166])) = int32(392)
	v2747 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[165]))
	if v2747 != 0 {
		v2769 = v2747
		goto L499
	} else {
		goto L513
	}
L513:
	;
	v2750 = int32(16)
	v2752 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2754 = base.I32_div_s(v2752, int32(512))
	v2756 = base.I32_rem_s(v2754, v2750)
	v2757 = v2754 - v2756
	if v2757 <= v2750 {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	v2781 = v2763
	goto L498
L515:
	;
	v2760 = v2750
	goto L517
L516:
	;
	v2760 = v2757
	goto L517
L517:
	;
	if int32(1024) < v2760 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2763 = int32(1024)
	goto L520
L519:
	;
	v2763 = v2760
	goto L520
L520:
	;
	goto L514
L521:
	;
	v2775 = v2772
	goto L523
L522:
	;
	v2775 = v2769
	goto L523
L523:
	;
	if int32(_a_F_CreateSharedMemoryAndSemaphores_64) <= v2775 {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v2778 = int32(_a_F_CreateSharedMemoryAndSemaphores_64)
	goto L526
L525:
	;
	v2778 = v2775
	goto L526
L526:
	;
	v2781 = v2778
	goto L498
L527:
	;
	m.G0 = v2702 + int32(48)
	v2793 = m.G0
	v2795 = v2793 - int32(16)
	m.G0 = v2795
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[167])) = int32(292)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[168])) = int32(293)
	v2806 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[169]))
	v2807 = int32(0)
	F_SimpleLruInit(m, int32(_a_F_CreateSharedMemoryAndSemaphores_70), int32(_a_F_CreateSharedMemoryAndSemaphores_71), v2806, v2807, int32(_a_F_CreateSharedMemoryAndSemaphores_72), int32(57), int32(88), int32(3), v2807)
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	v2818 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[170]))
	v2819 = int32(0)
	F_SimpleLruInit(m, int32(_a_F_CreateSharedMemoryAndSemaphores_73), int32(_a_F_CreateSharedMemoryAndSemaphores_74), v2818, v2819, int32(_a_F_CreateSharedMemoryAndSemaphores_75), int32(58), int32(87), int32(4), v2819)
	mBase = m.M
	v2826 = m.ExcPending
	if v2826 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	v2832 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v2834 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v2836 = F_mul_size(m, int32(8), v2832+v2834)
	mBase = m.M
	v2837 = m.ExcPending
	if v2837 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	v2838 = F_add_size(m, int32(48), v2836)
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	v2842 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_76), v2838, v2795+int32(15))
	mBase = m.M
	v2843 = m.ExcPending
	if v2843 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[173])) = v2842
	v2846 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v2846 != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v2892 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[173]))
	v2894 = v2892 + int32(48)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[174])) = v2894
	v2898 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v2899 = int32(2)
	v2903 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[175])) = v2894 + v2898<<(uint(v2899)%32) + v2903<<(uint(v2899)%32)
	v2908 = int32(16)
	m.G0 = v2795 + v2908
	v2912 = m.G0
	v2914 = v2912 - v2908
	m.G0 = v2914
	v2919 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2924 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_77), v2919<<(uint(int32(6))%32), v2914+int32(14))
	mBase = m.M
	v2925 = m.ExcPending
	if v2925 != 0 {
		goto L1
	} else {
		goto L545
	}
L534:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v2854 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v2856 = F_mul_size(m, int32(8), v2852+v2854)
	mBase = m.M
	v2857 = m.ExcPending
	if v2857 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	v2858 = F_add_size(m, int32(48), v2856)
	mBase = m.M
	v2859 = m.ExcPending
	if v2859 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	if v2842&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v2858))|v2858&int32(3) == int32(0) {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	if v2858 == int32(0) {
		goto L533
	} else {
		goto L540
	}
L538:
	;
	v2882 = v2858
	goto L539
L539:
	;
	if v2882 == int32(0) {
		goto L533
	} else {
		goto L544
	}
L540:
	;
	v2872 = v2858 + v2842
	v2874 = v2842 + int32(4)
	if base.Ui32(v2874) < base.Ui32(v2872) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2876 = v2872
	goto L543
L542:
	;
	v2876 = v2874
	goto L543
L543:
	;
	v2882 = (v2842^int32(-1)+v2876)&int32(-4) + int32(4)
	goto L539
L544:
	;
	base.MemoryFill(m, v2842, int32(0), v2882)
	goto L533
L545:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[176])) = v2924
	v2930 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2937 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_78), v2930<<(uint(int32(13))%32)|int32(_a_F_CreateSharedMemoryAndSemaphores_79), v2914+int32(15))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[177])) = (v2937 + int32(4095)) & int32(-4096)
	v2947 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2952 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_80), v2947<<(uint(int32(4))%32), v2914+int32(13))
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L1
	} else {
		goto L547
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[178])) = v2952
	v2958 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v2963 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_81), v2958*int32(20), v2914+int32(12))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[179])) = v2963
	v2966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2914)+14)))
	if v2966 != 0 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v3082 = m.G0
	v3084 = v3082 - int32(16)
	m.G0 = v3084
	v3087 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	v3088 = int32(128)
	v3089 = v3087 + v3088
	v3090 = m.G0
	v3092 = v3090 - int32(48)
	m.G0 = v3092
	*(*int32)(unsafe.Add(mBase, uint32(v3092))) = v3088
	*(*int64)(unsafe.Add(mBase, uint32(v3092)+16)) = int64(103079215124)
	v3101 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_82), v3089, v3089, v3092, int32(41))
	mBase = m.M
	v3102 = m.ExcPending
	if v3102 != 0 {
		goto L1
	} else {
		goto L563
	}
L550:
	;
	v2967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2914)+15)))
	if v2967&int32(1) != 0 {
		goto L549
	} else {
		goto L551
	}
L551:
	;
	v2970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2914)+13)))
	if v2970&int32(1) != 0 {
		goto L549
	} else {
		goto L552
	}
L552:
	;
	v2973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2914)+12)))
	if v2973&int32(1) != 0 {
		goto L549
	} else {
		goto L553
	}
L553:
	;
	v2977 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	if int32(0) < v2977 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v2982 = int32(0)
	goto L557
L555:
	;
	v3041 = v2977
	goto L556
L556:
	;
	v3058 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[176]))
	*(*int32)(unsafe.Add(mBase, uint32(v3058+v3041<<(uint(int32(6))%32)-int32(32)))) = int32(-1)
	goto L549
L557:
	;
	v2997 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[176]))
	v3000 = v2997 + v2982<<(uint(int32(6))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v3000)+24)) = int32(0)
	v3003 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3000)+16)) = v3003
	*(*int64)(unsafe.Add(mBase, uint32(v3000)+8)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v3000))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3000)+28)) = v3003
	*(*int32)(unsafe.Add(mBase, uint32(v3000)+20)) = v2982
	*(*int32)(unsafe.Add(mBase, uint32(v3000+int32(36)))) = v3003
	goto L559
L558:
	;
	v3041 = v3038
	goto L556
L559:
	;
	v3017 = v2982 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3000)+32)) = v3017
	v3020 = v3000 + int32(48)
	v3021 = int32(62)
	*(*uint16)(unsafe.Add(mBase, uint32(v3020))) = uint16(v3021)
	*(*int32)(unsafe.Add(mBase, uint32(v3020)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v3020)+8)) = int64(-1)
	goto L560
L560:
	;
	v3028 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[178]))
	v3029 = *(*int32)(unsafe.Add(mBase, uint32(v3000)+20))
	v3032 = v3028 + v3029<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v3032)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v3032))) = int64(-4294967296)
	goto L561
L561:
	;
	v3038 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	if v3017 < v3038 {
		v2982 = v3017
		goto L557
	} else {
		goto L562
	}
L562:
	;
	goto L558
L563:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[180])) = v3101
	m.G0 = v3092 + int32(48)
	v3112 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_83), int32(28), v3084+int32(15))
	mBase = m.M
	v3113 = m.ExcPending
	if v3113 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[181])) = v3112
	v3115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3084)+15)))
	if v3115 == int32(0) {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	v3118 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3112)+8)) = v3118
	*(*int32)(unsafe.Add(mBase, uint32(v3112))) = v3118
	v3123 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	*(*int32)(unsafe.Add(mBase, uint32(v3112)+4)) = v3118
	*(*int32)(unsafe.Add(mBase, uint32(v3112)+16)) = v3118
	*(*int32)(unsafe.Add(mBase, uint32(v3112)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3112)+12)) = v3123 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3112)+20)) = v3118
	goto L567
L566:
	;
	goto L567
L567:
	;
	m.G0 = v3084 + int32(16)
	v3139 = int32(_a_F_CreateSharedMemoryAndSemaphores_84)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[182])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[183])) = int32(_a_F_CreateSharedMemoryAndSemaphores_85)
	goto L568
L568:
	;
	m.G0 = v2914 + int32(16)
	v3147 = m.G0
	v3149 = v3147 + int32(-64)
	m.G0 = v3149
	v3152 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[184]))
	v3154 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v3156 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v3157 = F_add_size(m, v3154, v3156)
	mBase = m.M
	v3158 = m.ExcPending
	if v3158 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	v3159 = F_mul_size(m, v3152, v3157)
	mBase = m.M
	v3160 = m.ExcPending
	if v3160 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3149)+16)) = int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v3149)+32)) = int64(566935683088)
	v3168 = base.I32_div_s(v3159, int32(2))
	v3170 = v3147 + int32(-48)
	v3172 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_86), v3168, v3159, v3170, int32(41))
	mBase = m.M
	v3173 = m.ExcPending
	if v3173 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[185])) = v3172
	*(*int32)(unsafe.Add(mBase, uint32(v3149)+40)) = int32(1113)
	*(*int64)(unsafe.Add(mBase, uint32(v3149)+32)) = int64(154618822664)
	*(*int32)(unsafe.Add(mBase, uint32(v3149)+16)) = int32(16)
	v3183 = int32(1)
	v3188 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_87), v3168<<(uint(v3183)%32), v3159<<(uint(v3183)%32), v3170, int32(73))
	mBase = m.M
	v3189 = m.ExcPending
	if v3189 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[186])) = v3188
	v3196 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_88), int32(_a_F_CreateSharedMemoryAndSemaphores_89), v3147+int32(-49))
	mBase = m.M
	v3197 = m.ExcPending
	if v3197 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[187])) = v3196
	v3199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3149)+15)))
	if v3199 == int32(0) {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3196))) = int32(0)
	goto L576
L575:
	;
	goto L576
L576:
	;
	v3204 = int32(-64)
	m.G0 = v3149 - v3204
	v3207 = m.G0
	v3209 = v3207 + v3204
	m.G0 = v3209
	v3212 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[188]))
	v3214 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v3216 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v3217 = F_add_size(m, v3214, v3216)
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v3219 = F_mul_size(m, v3212, v3217)
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+12)) = int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v3209)+28)) = int64(103079215120)
	v3230 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_90), v3219, v3219, v3207+int32(-52), int32(_a_F_CreateSharedMemoryAndSemaphores_91))
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[189])) = v3230
	v3235 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v3235 != 0 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v3244 = v3230
	goto L582
L581:
	;
	v3240 = F_hash_search(m, v3230, int32(_a_F_CreateSharedMemoryAndSemaphores_92), int32(1), v3207+int32(-53))
	mBase = m.M
	v3241 = m.ExcPending
	if v3241 != 0 {
		goto L1
	} else {
		goto L583
	}
L582:
	;
	v3246 = F_get_hash_value(m, v3244, int32(_a_F_CreateSharedMemoryAndSemaphores_92))
	mBase = m.M
	v3247 = m.ExcPending
	if v3247 != 0 {
		goto L1
	} else {
		goto L584
	}
L583:
	;
	v3243 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[189]))
	v3244 = v3243
	goto L582
L584:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[190])) = v3246
	v3251 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[191])) = v3251 + v3246&int32(15)<<(uint(int32(7))%32) + int32(_a_F_CreateSharedMemoryAndSemaphores_93)
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+36)) = int32(1114)
	*(*int64)(unsafe.Add(mBase, uint32(v3209)+28)) = int64(137438953480)
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+12)) = int32(16)
	v3269 = v3219 << (uint(int32(1)) % 32)
	v3273 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_94), v3269, v3269, v3207+int32(-52), int32(_a_F_CreateSharedMemoryAndSemaphores_95))
	mBase = m.M
	v3274 = m.ExcPending
	if v3274 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[192])) = v3273
	v3280 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v3282 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v3283 = v3280 + v3282
	v3285 = v3283 * int32(10)
	v3287 = F_mul_size(m, v3285, int32(120))
	mBase = m.M
	v3288 = m.ExcPending
	if v3288 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v3289 = F_add_size(m, int32(64), v3287)
	mBase = m.M
	v3290 = m.ExcPending
	if v3290 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	v3293 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_96), v3289, v3207+int32(-53))
	mBase = m.M
	v3294 = m.ExcPending
	if v3294 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[193])) = v3293
	v3296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3209)+11)))
	if v3296 == int32(0) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v3299 = int32(0)
	if v3289 != 0 {
		goto L592
	} else {
		goto L593
	}
L590:
	;
	v3467 = v3293
	goto L591
L591:
	;
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3467)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[194])) = v3485
	*(*int64)(unsafe.Add(mBase, uint32(v3209)+28)) = int64(34359738372)
	v3494 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_97), v3285, v3285, v3207+int32(-52), int32(_a_F_CreateSharedMemoryAndSemaphores_98))
	mBase = m.M
	v3495 = m.ExcPending
	if v3495 != 0 {
		goto L1
	} else {
		goto L608
	}
L592:
	;
	base.MemoryFill(m, v3293, int32(0), v3289)
	goto L594
L593:
	;
	goto L594
L594:
	;
	v3302 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3293)+40)) = v3302
	*(*int64)(unsafe.Add(mBase, uint32(v3293)+32)) = int64(1)
	v3306 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3293)+24)) = v3306
	*(*int64)(unsafe.Add(mBase, uint32(v3293)+16)) = v3302
	v3311 = v3293 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v3293)+12)) = v3311
	*(*int32)(unsafe.Add(mBase, uint32(v3293)+8)) = v3311
	*(*int64)(unsafe.Add(mBase, uint32(v3293)+48)) = v3302
	*(*int32)(unsafe.Add(mBase, uint32(v3293)+60)) = v3293 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v3293)+4)) = v3293
	*(*int32)(unsafe.Add(mBase, uint32(v3293))) = v3293
	if v3285 <= v3306 {
		v3395 = v3293
		v3396 = v3299
		goto L595
	} else {
		goto L596
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3395)+56)) = v3396
	v3412 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3396))) = v3412
	v3414 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	v3415 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3414)+4)) = v3415
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	v3418 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3417)+8)) = v3418
	v3420 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v3420)+16)) = v3418
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v3423)+24)) = v3418
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	v3428 = v3426 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v3426)+36)) = v3428
	*(*int32)(unsafe.Add(mBase, uint32(v3426)+32)) = v3428
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	v3433 = v3431 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v3431)+44)) = v3433
	*(*int32)(unsafe.Add(mBase, uint32(v3431)+40)) = v3433
	v3436 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	v3438 = v3436 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v3436)+52)) = v3438
	*(*int32)(unsafe.Add(mBase, uint32(v3436)+48)) = v3438
	v3441 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v3441)+56)) = v3418
	v3444 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	v3446 = v3444 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v3444)+92)) = v3446
	*(*int32)(unsafe.Add(mBase, uint32(v3444)+88)) = v3446
	v3449 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3449)+96)) = v3415
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3452)+100)) = v3415
	v3455 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3455)+104)) = v3415
	v3458 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3458)+108)) = int32(1)
	v3461 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3461)+112)) = v3415
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(v3395)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3464)+116)) = v3412
	v3467 = v3395
	goto L591
L596:
	;
	v3323 = v3293
	v3324 = v3299
	goto L597
L597:
	;
	v3340 = v3324 * int32(120)
	v3341 = *(*int32)(unsafe.Add(mBase, uint32(v3323)+60))
	v3344 = v3340 + v3341 + int32(72)
	v3345 = int32(78)
	*(*uint16)(unsafe.Add(mBase, uint32(v3344))) = uint16(v3345)
	*(*int32)(unsafe.Add(mBase, uint32(v3344)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v3344)+8)) = int64(-1)
	goto L599
L598:
	;
	v3370 = int32(0)
	v3371 = *(*int32)(unsafe.Add(mBase, uint32(v3352)+4))
	if base.B2i32(v3371 == v3370)|base.B2i32(v3352 == v3371) != 0 {
		v3395 = v3352
		v3396 = v3370
		goto L595
	} else {
		goto L604
	}
L599:
	;
	v3352 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[193]))
	v3353 = *(*int32)(unsafe.Add(mBase, uint32(v3352)+60))
	v3354 = v3353 + v3340
	v3356 = v3354 - int32(-64)
	v3357 = *(*int32)(unsafe.Add(mBase, uint32(v3352)+4))
	if v3357 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3352)+4)) = v3352
	*(*int32)(unsafe.Add(mBase, uint32(v3352))) = v3352
	goto L602
L601:
	;
	goto L602
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3354)+68)) = v3352
	v3363 = *(*int32)(unsafe.Add(mBase, uint32(v3352)))
	*(*int32)(unsafe.Add(mBase, uint32(v3354)+64)) = v3363
	*(*int32)(unsafe.Add(mBase, uint32(v3363)+4)) = v3356
	*(*int32)(unsafe.Add(mBase, uint32(v3352))) = v3356
	v3368 = v3324 + int32(1)
	if v3368 != v3285 {
		v3323 = v3352
		v3324 = v3368
		goto L597
	} else {
		goto L603
	}
L603:
	;
	goto L598
L604:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(v3371)))
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(v3371)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3376)+4)) = v3377
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3371)))
	*(*int32)(unsafe.Add(mBase, uint32(v3377))) = v3379
	v3382 = v3352 + int32(8)
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v3352)+12))
	if v3383 == int32(0) {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3352)+12)) = v3382
	*(*int32)(unsafe.Add(mBase, uint32(v3352)+8)) = v3382
	goto L607
L606:
	;
	goto L607
L607:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3371)+4)) = v3382
	v3391 = *(*int32)(unsafe.Add(mBase, uint32(v3382)))
	*(*int32)(unsafe.Add(mBase, uint32(v3371))) = v3391
	*(*int32)(unsafe.Add(mBase, uint32(v3391)+4)) = v3371
	*(*int32)(unsafe.Add(mBase, uint32(v3382))) = v3371
	v3395 = v3352
	v3396 = v3371 + int32(-64)
	goto L595
L608:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[195])) = v3494
	v3500 = v3283 * int32(50)
	v3502 = F_mul_size(m, v3500, int32(24))
	mBase = m.M
	v3503 = m.ExcPending
	if v3503 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	v3505 = v3502 + int32(16)
	v3508 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_99), v3505, v3207+int32(-53))
	mBase = m.M
	v3509 = m.ExcPending
	if v3509 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[196])) = v3508
	v3511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3209)+11)))
	if v3511 != 0 {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v3589 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_100), int32(8), v3207+int32(-53))
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L1
	} else {
		goto L626
	}
L612:
	;
	if v3505 != 0 {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	base.MemoryFill(m, v3508, int32(0), v3505)
	goto L615
L614:
	;
	goto L615
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3508)+8)) = v3508 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v3508)+4)) = v3508
	*(*int32)(unsafe.Add(mBase, uint32(v3508))) = v3508
	if v3500 <= int32(0) {
		goto L611
	} else {
		goto L616
	}
L616:
	;
	v3523 = int32(0)
	goto L617
L617:
	;
	v3538 = v3523 * int32(24)
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v3508)+8))
	v3540 = v3538 + v3539
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(v3508)+4))
	if v3541 == int32(0) {
		goto L619
	} else {
		goto L620
	}
L618:
	;
	goto L611
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3508)+4)) = v3508
	*(*int32)(unsafe.Add(mBase, uint32(v3508))) = v3508
	goto L621
L620:
	;
	goto L621
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3540)+4)) = v3508
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(v3508)))
	*(*int32)(unsafe.Add(mBase, uint32(v3540))) = v3547
	*(*int32)(unsafe.Add(mBase, uint32(v3547)+4)) = v3540
	*(*int32)(unsafe.Add(mBase, uint32(v3508))) = v3540
	v3551 = *(*int32)(unsafe.Add(mBase, uint32(v3508)+8))
	v3552 = v3551 + v3538
	v3554 = v3552 + int32(24)
	v3555 = *(*int32)(unsafe.Add(mBase, uint32(v3508)+4))
	if v3555 == int32(0) {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3508)+4)) = v3508
	*(*int32)(unsafe.Add(mBase, uint32(v3508))) = v3508
	goto L624
L623:
	;
	goto L624
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3552)+28)) = v3508
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v3508)))
	*(*int32)(unsafe.Add(mBase, uint32(v3552)+24)) = v3561
	*(*int32)(unsafe.Add(mBase, uint32(v3561)+4)) = v3554
	*(*int32)(unsafe.Add(mBase, uint32(v3508))) = v3554
	v3566 = v3523 + int32(2)
	if v3566 != v3500 {
		v3523 = v3566
		goto L617
	} else {
		goto L625
	}
L625:
	;
	goto L618
L626:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[197])) = v3589
	v3592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3209)+11)))
	if v3592 == int32(0) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3589)+4)) = v3589
	*(*int32)(unsafe.Add(mBase, uint32(v3589))) = v3589
	goto L629
L628:
	;
	goto L629
L629:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[198])) = int32(1115)
	v3603 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[199]))
	v3604 = int32(0)
	F_SimpleLruInit(m, int32(_a_F_CreateSharedMemoryAndSemaphores_101), int32(_a_F_CreateSharedMemoryAndSemaphores_102), v3603, v3604, int32(_a_F_CreateSharedMemoryAndSemaphores_103), int32(60), int32(90), int32(5), v3604)
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	v3617 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_104), int32(16), v3207+int32(-1))
	mBase = m.M
	v3618 = m.ExcPending
	if v3618 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[200])) = v3617
	v3620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3209)+63)))
	if v3620 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v3624 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	v3628 = F_LWLockAcquire(m, v3624+int32(_a_F_CreateSharedMemoryAndSemaphores_105), int32(0))
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L1
	} else {
		goto L635
	}
L633:
	;
	goto L634
L634:
	;
	m.G0 = v3209 - int32(-64)
	v3647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v3647 == int32(0) {
		goto L637
	} else {
		goto L638
	}
L635:
	;
	v3631 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[200]))
	*(*int64)(unsafe.Add(mBase, uint32(v3631)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3631))) = int64(-1)
	v3637 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	F_LWLockRelease(m, v3637+int32(_a_F_CreateSharedMemoryAndSemaphores_105))
	mBase = m.M
	v3641 = m.ExcPending
	if v3641 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	goto L634
L637:
	;
	v3650 = m.G0
	v3652 = v3650 - int32(16)
	m.G0 = v3652
	v3655 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v3657 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v3662 = v3652 + int32(15)
	v3663 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_106), int32(76), v3662)
	mBase = m.M
	v3664 = m.ExcPending
	if v3664 != 0 {
		goto L1
	} else {
		goto L640
	}
L638:
	;
	goto L639
L639:
	;
	v4105 = m.G0
	v4107 = v4105 - int32(16)
	m.G0 = v4107
	v4114 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v4116 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4118 = F_mul_size(m, int32(4), v4114+v4116)
	mBase = m.M
	v4119 = m.ExcPending
	if v4119 != 0 {
		goto L1
	} else {
		goto L709
	}
L640:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201])) = v3663
	v3666 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+52)) = v3666
	v3669 = v3663 + int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+48)) = v3669
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+44)) = v3669
	v3673 = v3663 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+40)) = v3673
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+36)) = v3673
	v3677 = v3663 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+32)) = v3677
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+28)) = v3677
	v3681 = v3663 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+24)) = v3681
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+20)) = v3681
	*(*int64)(unsafe.Add(mBase, uint32(v3663)+68)) = int64(-4294967196)
	*(*int64)(unsafe.Add(mBase, uint32(v3663)+60)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3663)+56)) = v3666
	v3692 = v3655 + v3657 + int32(38)
	v3694 = F_PGProcShmemSize(m)
	mBase = m.M
	v3695 = m.ExcPending
	if v3695 != 0 {
		goto L1
	} else {
		goto L642
	}
L641:
	;
	v3730 = int32(_a_F_CreateSharedMemoryAndSemaphores_107)
	v3731 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	*(*int32)(unsafe.Add(mBase, uint32(v3731))) = v3696
	v3734 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v3736 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	v3739 = v3696 + v3692*int32(640)
	v3742 = v3739 + v3692<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v3736)+8)) = v3742
	*(*int32)(unsafe.Add(mBase, uint32(v3736)+4)) = v3739
	*(*int32)(unsafe.Add(mBase, uint32(v3736)+12)) = v3742 + v3692<<(uint(int32(1))%32)
	v3749 = int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v3736)+16)) = v3734 + v3749
	v3753 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[202]))
	v3758 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v3759 = F_add_size(m, v3749, v3758)
	mBase = m.M
	v3760 = m.ExcPending
	if v3760 != 0 {
		goto L1
	} else {
		goto L653
	}
L642:
	;
	v3696 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_108), v3694, v3662)
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L1
	} else {
		goto L643
	}
L643:
	;
	v3698 = int32(3)
	if v3696&v3698|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v3694))|v3694&v3698 == int32(0) {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	if v3694 == int32(0) {
		goto L641
	} else {
		goto L647
	}
L645:
	;
	v3722 = v3694
	goto L646
L646:
	;
	if v3722 == int32(0) {
		goto L641
	} else {
		goto L651
	}
L647:
	;
	v3712 = v3694 + v3696
	v3714 = v3696 + int32(4)
	if base.Ui32(v3714) < base.Ui32(v3712) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	v3716 = v3712
	goto L650
L649:
	;
	v3716 = v3714
	goto L650
L650:
	;
	v3722 = (v3696^int32(-1)+v3716)&int32(-4) + int32(4)
	goto L646
L651:
	;
	base.MemoryFill(m, v3696, int32(0), v3722)
	goto L641
L652:
	;
	if v3692 != 0 {
		goto L666
	} else {
		goto L667
	}
L653:
	;
	v3761 = F_add_size(m, v3734, v3759)
	mBase = m.M
	v3762 = m.ExcPending
	if v3762 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	v3764 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[202]))
	v3767 = F_mul_size(m, v3761, v3764*int32(72))
	mBase = m.M
	v3768 = m.ExcPending
	if v3768 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	v3769 = F_add_size(m, int32(0), v3767)
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	v3773 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_109), v3769, v3652+int32(15))
	mBase = m.M
	v3774 = m.ExcPending
	if v3774 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	v3775 = int32(3)
	if v3773&v3775|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v3769))|v3769&v3775 == int32(0) {
		goto L658
	} else {
		goto L659
	}
L658:
	;
	if v3769 == int32(0) {
		goto L652
	} else {
		goto L661
	}
L659:
	;
	v3799 = v3769
	goto L660
L660:
	;
	if v3799 == int32(0) {
		goto L652
	} else {
		goto L665
	}
L661:
	;
	v3789 = v3769 + v3773
	v3791 = v3773 + int32(4)
	if base.Ui32(v3791) < base.Ui32(v3789) {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v3793 = v3789
	goto L664
L663:
	;
	v3793 = v3791
	goto L664
L664:
	;
	v3799 = (v3773^int32(-1)+v3793)&int32(-4) + int32(4)
	goto L660
L665:
	;
	base.MemoryFill(m, v3773, int32(0), v3799)
	goto L652
L666:
	;
	v3814 = v3773
	v3815 = int32(0)
	goto L669
L667:
	;
	goto L668
L668:
	;
	v4067 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4070 = v3696 + v4067*int32(640)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[203])) = v4070
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[204])) = v4070 + int32(_a_F_CreateSharedMemoryAndSemaphores_110)
	v4081 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_111), int32(4), v3652+int32(15))
	mBase = m.M
	v4082 = m.ExcPending
	if v4082 != 0 {
		goto L1
	} else {
		goto L708
	}
L669:
	;
	v3830 = v3696 + v3815*int32(640)
	v3831 = v3814 + v3753<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+604)) = v3831
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+600)) = v3814
	v3835 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v3815 < v3835+int32(38) {
		goto L671
	} else {
		goto L672
	}
L670:
	;
	goto L668
L671:
	;
	v3840 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[16]))
	v3842 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[15]))
	if v3840 < v3842 {
		goto L675
	} else {
		goto L676
	}
L672:
	;
	goto L673
L673:
	;
	v3887 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[205]))
	if v3815 < v3887 {
		goto L685
	} else {
		goto L686
	}
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+12)) = v3845 + v3840<<(uint(int32(7))%32)
	v3870 = v3830 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v3870)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3870))) = int64(0)
	v3875 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3870)+8)) = uint8(v3875)
	goto L681
L675:
	;
	v3845 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[14]))
	v3849 = int32(_a_F_CreateSharedMemoryAndSemaphores_112)
	v3851 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[16])) = v3851 + int32(1)
	goto L674
L676:
	;
	goto L677
L677:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3858 = m.ExcPending
	if v3858 != 0 {
		goto L1
	} else {
		goto L678
	}
L678:
	;
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_113), int32(0))
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
		goto L1
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_36), int32(270), int32(_a_F_CreateSharedMemoryAndSemaphores_114))
	mBase = m.M
	v3867 = m.ExcPending
	if v3867 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L681:
	;
	v3878 = v3830 + int32(584)
	v3879 = int32(65)
	*(*uint16)(unsafe.Add(mBase, uint32(v3878))) = uint16(v3879)
	*(*int32)(unsafe.Add(mBase, uint32(v3878)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v3878)+8)) = int64(-1)
	goto L682
L682:
	;
	goto L673
L683:
	;
	v3973 = v3830 + int32(620)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+624)) = v3973
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+620)) = v3973
	v3977 = v3830 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+272)) = v3977
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+268)) = v3977
	v3981 = v3830 + int32(260)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+264)) = v3981
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+260)) = v3981
	v3985 = v3830 + int32(252)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+256)) = v3985
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+252)) = v3985
	v3989 = v3830 + int32(244)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+248)) = v3989
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+244)) = v3989
	v3993 = v3830 + int32(236)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+240)) = v3993
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+236)) = v3993
	v3997 = v3830 + int32(228)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+232)) = v3997
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+228)) = v3997
	v4001 = v3830 + int32(220)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+224)) = v4001
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+220)) = v4001
	v4005 = v3830 + int32(212)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+216)) = v4005
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+212)) = v4005
	v4009 = v3830 + int32(204)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+208)) = v4009
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+204)) = v4009
	v4013 = v3830 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+200)) = v4013
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+196)) = v4013
	v4017 = v3830 + int32(188)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+192)) = v4017
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+188)) = v4017
	v4021 = v3830 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+184)) = v4021
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+180)) = v4021
	v4025 = v3830 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+176)) = v4025
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+172)) = v4025
	v4029 = v3830 + int32(164)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+168)) = v4029
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+164)) = v4029
	v4033 = v3830 + int32(156)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+160)) = v4033
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+156)) = v4033
	v4037 = v3830 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+152)) = v4037
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+148)) = v4037
	v4040 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+540)) = v4040
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+556)) = v4040
	*(*int64)(unsafe.Add(mBase, uint32(v3830)+112)) = int64(0)
	v4047 = v3815 + int32(1)
	if v4047 != v3692 {
		v3814 = v3753<<(uint(int32(6))%32) + v3831
		v3815 = v4047
		goto L669
	} else {
		goto L707
	}
L684:
	;
	v3966 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+8)) = v3964 + v3966
	goto L683
L685:
	;
	v3890 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	v3892 = v3890 + int32(20)
	v3893 = *(*int32)(unsafe.Add(mBase, uint32(v3890)+24))
	if v3893 == int32(0) {
		goto L688
	} else {
		goto L689
	}
L686:
	;
	goto L687
L687:
	;
	v3905 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[206]))
	v3908 = v3887 + v3905 + int32(2)
	if v3815 < v3908 {
		goto L691
	} else {
		goto L692
	}
L688:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3890)+24)) = v3892
	*(*int32)(unsafe.Add(mBase, uint32(v3890)+20)) = v3892
	goto L690
L689:
	;
	goto L690
L690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+4)) = v3892
	v3899 = *(*int32)(unsafe.Add(mBase, uint32(v3892)))
	*(*int32)(unsafe.Add(mBase, uint32(v3830))) = v3899
	*(*int32)(unsafe.Add(mBase, uint32(v3899)+4)) = v3830
	*(*int32)(unsafe.Add(mBase, uint32(v3892))) = v3830
	v3964 = int32(20)
	goto L684
L691:
	;
	v3911 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	v3913 = v3911 + int32(28)
	v3914 = *(*int32)(unsafe.Add(mBase, uint32(v3911)+32))
	if v3914 == int32(0) {
		goto L694
	} else {
		goto L695
	}
L692:
	;
	goto L693
L693:
	;
	v3926 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[207]))
	if v3815 < v3926+v3908 {
		goto L697
	} else {
		goto L698
	}
L694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3911)+32)) = v3913
	*(*int32)(unsafe.Add(mBase, uint32(v3911)+28)) = v3913
	goto L696
L695:
	;
	goto L696
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+4)) = v3913
	v3920 = *(*int32)(unsafe.Add(mBase, uint32(v3913)))
	*(*int32)(unsafe.Add(mBase, uint32(v3830))) = v3920
	*(*int32)(unsafe.Add(mBase, uint32(v3920)+4)) = v3830
	*(*int32)(unsafe.Add(mBase, uint32(v3913))) = v3830
	v3964 = int32(28)
	goto L684
L697:
	;
	v3930 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	v3932 = v3930 + int32(36)
	v3933 = *(*int32)(unsafe.Add(mBase, uint32(v3930)+40))
	if v3933 == int32(0) {
		goto L700
	} else {
		goto L701
	}
L698:
	;
	goto L699
L699:
	;
	v3945 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v3945 <= v3815 {
		goto L683
	} else {
		goto L703
	}
L700:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3930)+40)) = v3932
	*(*int32)(unsafe.Add(mBase, uint32(v3930)+36)) = v3932
	goto L702
L701:
	;
	goto L702
L702:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+4)) = v3932
	v3939 = *(*int32)(unsafe.Add(mBase, uint32(v3932)))
	*(*int32)(unsafe.Add(mBase, uint32(v3830))) = v3939
	*(*int32)(unsafe.Add(mBase, uint32(v3939)+4)) = v3830
	*(*int32)(unsafe.Add(mBase, uint32(v3932))) = v3830
	v3964 = int32(36)
	goto L684
L703:
	;
	v3948 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	v3950 = v3948 + int32(44)
	v3951 = *(*int32)(unsafe.Add(mBase, uint32(v3948)+48))
	if v3951 == int32(0) {
		goto L704
	} else {
		goto L705
	}
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3948)+48)) = v3950
	*(*int32)(unsafe.Add(mBase, uint32(v3948)+44)) = v3950
	goto L706
L705:
	;
	goto L706
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3830)+4)) = v3950
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v3950)))
	*(*int32)(unsafe.Add(mBase, uint32(v3830))) = v3957
	*(*int32)(unsafe.Add(mBase, uint32(v3957)+4)) = v3830
	*(*int32)(unsafe.Add(mBase, uint32(v3950))) = v3830
	v3964 = int32(44)
	goto L684
L707:
	;
	goto L670
L708:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[208])) = v4081
	*(*int32)(unsafe.Add(mBase, uint32(v4081))) = int32(0)
	m.G0 = v3652 + int32(16)
	goto L639
L709:
	;
	v4120 = F_add_size(m, int32(36), v4118)
	mBase = m.M
	v4121 = m.ExcPending
	if v4121 != 0 {
		goto L1
	} else {
		goto L710
	}
L710:
	;
	v4124 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_115), v4120, v4107+int32(15))
	mBase = m.M
	v4125 = m.ExcPending
	if v4125 != 0 {
		goto L1
	} else {
		goto L711
	}
L711:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[209])) = v4124
	v4127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4107)+15)))
	if v4127 == int32(0) {
		goto L712
	} else {
		goto L713
	}
L712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4124))) = int32(0)
	v4133 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v4135 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4136 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4124)+12)) = v4136
	*(*int64)(unsafe.Add(mBase, uint32(v4124)+20)) = v4136
	*(*int64)(unsafe.Add(mBase, uint32(v4124)+28)) = v4136
	v4142 = v4133 + v4135
	*(*int32)(unsafe.Add(mBase, uint32(v4124)+4)) = v4142
	*(*int32)(unsafe.Add(mBase, uint32(v4124)+8)) = v4142 * int32(65)
	v4148 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[152]))
	*(*int64)(unsafe.Add(mBase, uint32(v4148)+56)) = int64(1)
	goto L714
L713:
	;
	goto L714
L714:
	;
	v4155 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	v4156 = *(*int32)(unsafe.Add(mBase, uint32(v4155)))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[210])) = v4156
	v4159 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[211])))
	if v4159 == int32(1) {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v4166 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v4168 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4172 = F_mul_size(m, int32(4), (v4166+v4168)*int32(65))
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L1
	} else {
		goto L718
	}
L716:
	;
	goto L717
L717:
	;
	v4195 = int32(16)
	m.G0 = v4107 + v4195
	v4198 = m.G0
	v4200 = v4198 - v4195
	m.G0 = v4200
	v4206 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4209 = F_mul_size(m, int32(408), v4206+int32(38))
	mBase = m.M
	v4210 = m.ExcPending
	if v4210 != 0 {
		goto L1
	} else {
		goto L722
	}
L718:
	;
	v4175 = v4107 + int32(15)
	v4176 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_116), v4172, v4175)
	mBase = m.M
	v4177 = m.ExcPending
	if v4177 != 0 {
		goto L1
	} else {
		goto L719
	}
L719:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[212])) = v4176
	v4183 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v4185 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4189 = F_mul_size(m, int32(1), (v4183+v4185)*int32(65))
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	v4191 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_117), v4189, v4175)
	mBase = m.M
	v4192 = m.ExcPending
	if v4192 != 0 {
		goto L1
	} else {
		goto L721
	}
L721:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[213])) = v4191
	goto L717
L722:
	;
	v4213 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_118), v4209, v4200+int32(15))
	mBase = m.M
	v4214 = m.ExcPending
	if v4214 != 0 {
		goto L1
	} else {
		goto L723
	}
L723:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[214])) = v4213
	v4216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4200)+15)))
	if v4216 != 0 {
		goto L724
	} else {
		goto L725
	}
L724:
	;
	v4253 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4256 = F_mul_size(m, int32(64), v4253+int32(38))
	mBase = m.M
	v4257 = m.ExcPending
	if v4257 != 0 {
		goto L1
	} else {
		goto L734
	}
L725:
	;
	v4217 = int32(3)
	if v4209&v4217|(v4213&v4217|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v4209))) == int32(0) {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	if v4209 == int32(0) {
		goto L724
	} else {
		goto L729
	}
L727:
	;
	v4241 = v4209
	goto L728
L728:
	;
	if v4241 == int32(0) {
		goto L724
	} else {
		goto L733
	}
L729:
	;
	v4231 = v4213 + v4209
	v4233 = v4213 + int32(4)
	if base.Ui32(v4233) < base.Ui32(v4231) {
		goto L730
	} else {
		goto L731
	}
L730:
	;
	v4235 = v4231
	goto L732
L731:
	;
	v4235 = v4233
	goto L732
L732:
	;
	v4241 = (v4213^int32(-1)+v4235)&int32(-4) + int32(4)
	goto L728
L733:
	;
	base.MemoryFill(m, v4213, int32(0), v4241)
	goto L724
L734:
	;
	v4260 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_119), v4256, v4200+int32(15))
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[215])) = v4260
	v4263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4200)+15)))
	if v4263 == int32(1) {
		goto L737
	} else {
		goto L738
	}
L736:
	;
	v4359 = F_mul_size(m, int32(64), v4342)
	mBase = m.M
	v4360 = m.ExcPending
	if v4360 != 0 {
		goto L1
	} else {
		goto L753
	}
L737:
	;
	v4267 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4342 = v4267 + int32(38)
	goto L736
L738:
	;
	goto L739
L739:
	;
	v4270 = int32(3)
	if v4256&v4270|(v4260&v4270|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v4256))) == int32(0) {
		goto L741
	} else {
		goto L742
	}
L740:
	;
	v4303 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4305 = v4303 + int32(38)
	if v4305 <= int32(0) {
		v4342 = v4305
		goto L736
	} else {
		goto L749
	}
L741:
	;
	if v4256 == int32(0) {
		goto L740
	} else {
		goto L744
	}
L742:
	;
	v4294 = v4256
	goto L743
L743:
	;
	if v4294 == int32(0) {
		goto L740
	} else {
		goto L748
	}
L744:
	;
	v4284 = v4256 + v4260
	v4286 = v4260 + int32(4)
	if base.Ui32(v4286) < base.Ui32(v4284) {
		goto L745
	} else {
		goto L746
	}
L745:
	;
	v4288 = v4284
	goto L747
L746:
	;
	v4288 = v4286
	goto L747
L747:
	;
	v4294 = (v4260^int32(-1)+v4288)&int32(-4) + int32(4)
	goto L743
L748:
	;
	base.MemoryFill(m, v4260, int32(0), v4294)
	goto L740
L749:
	;
	v4309 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[214]))
	v4311 = int32(0)
	v4312 = v4260
	goto L750
L750:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4309+v4311*int32(408))+212)) = v4312
	v4334 = v4311 + int32(1)
	v4336 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4338 = v4336 + int32(38)
	if v4334 < v4338 {
		v4311 = v4334
		v4312 = v4312 - int32(-64)
		goto L750
	} else {
		goto L752
	}
L751:
	;
	v4342 = v4338
	goto L736
L752:
	;
	goto L751
L753:
	;
	v4363 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_120), v4359, v4200+int32(15))
	mBase = m.M
	v4364 = m.ExcPending
	if v4364 != 0 {
		goto L1
	} else {
		goto L754
	}
L754:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[216])) = v4363
	v4366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4200)+15)))
	if v4366 == int32(1) {
		goto L756
	} else {
		goto L757
	}
L755:
	;
	v4461 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[217]))
	v4462 = F_mul_size(m, v4461, v4445)
	mBase = m.M
	v4463 = m.ExcPending
	if v4463 != 0 {
		goto L1
	} else {
		goto L772
	}
L756:
	;
	v4370 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4445 = v4370 + int32(38)
	goto L755
L757:
	;
	goto L758
L758:
	;
	v4373 = int32(3)
	if v4359&v4373|(v4363&v4373|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v4359))) == int32(0) {
		goto L760
	} else {
		goto L761
	}
L759:
	;
	v4406 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4408 = v4406 + int32(38)
	if v4408 <= int32(0) {
		v4445 = v4408
		goto L755
	} else {
		goto L768
	}
L760:
	;
	if v4359 == int32(0) {
		goto L759
	} else {
		goto L763
	}
L761:
	;
	v4397 = v4359
	goto L762
L762:
	;
	if v4397 == int32(0) {
		goto L759
	} else {
		goto L767
	}
L763:
	;
	v4387 = v4359 + v4363
	v4389 = v4363 + int32(4)
	if base.Ui32(v4389) < base.Ui32(v4387) {
		goto L764
	} else {
		goto L765
	}
L764:
	;
	v4391 = v4387
	goto L766
L765:
	;
	v4391 = v4389
	goto L766
L766:
	;
	v4397 = (v4363^int32(-1)+v4391)&int32(-4) + int32(4)
	goto L762
L767:
	;
	base.MemoryFill(m, v4363, int32(0), v4397)
	goto L759
L768:
	;
	v4412 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[214]))
	v4414 = int32(0)
	v4415 = v4363
	goto L769
L769:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4412+v4414*int32(408))+188)) = v4415
	v4437 = v4414 + int32(1)
	v4439 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4441 = v4439 + int32(38)
	if v4437 < v4441 {
		v4414 = v4437
		v4415 = v4415 - int32(-64)
		goto L769
	} else {
		goto L771
	}
L770:
	;
	v4445 = v4441
	goto L755
L771:
	;
	goto L770
L772:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[218])) = v4462
	v4469 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_121), v4462, v4200+int32(15))
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		goto L1
	} else {
		goto L773
	}
L773:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[219])) = v4469
	v4472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4200)+15)))
	if v4472 != 0 {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v4562 = int32(16)
	m.G0 = v4200 + v4562
	v4566 = m.G0
	v4568 = v4566 - v4562
	m.G0 = v4568
	v4574 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v4576 = F_mul_size(m, v4574, int32(4))
	mBase = m.M
	v4577 = m.ExcPending
	if v4577 != 0 {
		goto L1
	} else {
		goto L789
	}
L775:
	;
	v4473 = int32(3)
	v4476 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[218]))
	if v4469&v4473|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v4476))|v4476&v4473 == int32(0) {
		goto L777
	} else {
		goto L778
	}
L776:
	;
	v4508 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v4508+int32(38) <= int32(0) {
		goto L774
	} else {
		goto L785
	}
L777:
	;
	if v4476 == int32(0) {
		goto L776
	} else {
		goto L780
	}
L778:
	;
	v4499 = v4476
	goto L779
L779:
	;
	if v4499 == int32(0) {
		goto L776
	} else {
		goto L784
	}
L780:
	;
	v4489 = v4476 + v4469
	v4491 = v4469 + int32(4)
	if base.Ui32(v4491) < base.Ui32(v4489) {
		goto L781
	} else {
		goto L782
	}
L781:
	;
	v4493 = v4489
	goto L783
L782:
	;
	v4493 = v4491
	goto L783
L783:
	;
	v4499 = (v4469^int32(-1)+v4493)&int32(-4) + int32(4)
	goto L779
L784:
	;
	base.MemoryFill(m, v4469, int32(0), v4499)
	goto L776
L785:
	;
	v4514 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[214]))
	v4516 = int32(0)
	v4517 = v4469
	goto L786
L786:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4514+v4516*int32(408))+216)) = v4517
	v4537 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[217]))
	v4540 = v4516 + int32(1)
	v4542 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v4540 < v4542+int32(38) {
		v4516 = v4540
		v4517 = v4517 + v4537
		goto L786
	} else {
		goto L788
	}
L787:
	;
	goto L774
L788:
	;
	goto L787
L789:
	;
	v4578 = F_add_size(m, int32(8), v4576)
	mBase = m.M
	v4579 = m.ExcPending
	if v4579 != 0 {
		goto L1
	} else {
		goto L790
	}
L790:
	;
	v4585 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	v4587 = F_mul_size(m, v4585, int32(248))
	mBase = m.M
	v4588 = m.ExcPending
	if v4588 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	v4589 = F_add_size(m, (v4578+int32(7))&int32(-8), v4587)
	mBase = m.M
	v4590 = m.ExcPending
	if v4590 != 0 {
		goto L1
	} else {
		goto L792
	}
L792:
	;
	v4593 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_122), v4589, v4568+int32(15))
	mBase = m.M
	v4594 = m.ExcPending
	if v4594 != 0 {
		goto L1
	} else {
		goto L793
	}
L793:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[220])) = v4593
	v4597 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v4597 != 0 {
		goto L794
	} else {
		goto L795
	}
L794:
	;
	v4666 = int32(16)
	m.G0 = v4568 + v4666
	v4669 = int32(0)
	v4670 = m.G0
	v4672 = v4670 - v4666
	m.G0 = v4672
	v4678 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[207]))
	v4680 = F_mul_size(m, v4678, int32(1480))
	mBase = m.M
	v4681 = m.ExcPending
	if v4681 != 0 {
		goto L1
	} else {
		goto L800
	}
L795:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4593))) = int64(0)
	v4601 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	if v4601 <= int32(0) {
		goto L794
	} else {
		goto L796
	}
L796:
	;
	v4613 = int32(0)
	v4614 = int32(0)
	goto L797
L797:
	;
	v4630 = v4593 + (v4601<<(uint(int32(2))%32)+int32(15))&int32(-8) + v4614*int32(248)
	*(*int32)(unsafe.Add(mBase, uint32(v4630))) = v4613
	*(*int32)(unsafe.Add(mBase, uint32(v4593))) = v4630
	v4634 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[204]))
	v4635 = int32(640)
	v4639 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[201]))
	v4640 = *(*int32)(unsafe.Add(mBase, uint32(v4639)))
	v4643 = base.I32_div_s(v4634+v4614*v4635-v4640, v4635)
	*(*int32)(unsafe.Add(mBase, uint32(v4630)+4)) = v4643
	v4646 = v4614 + int32(1)
	v4648 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[171]))
	if v4646 < v4648 {
		v4613 = v4630
		v4614 = v4646
		goto L797
	} else {
		goto L799
	}
L798:
	;
	goto L794
L799:
	;
	goto L798
L800:
	;
	v4682 = F_add_size(m, v4666, v4680)
	mBase = m.M
	v4683 = m.ExcPending
	if v4683 != 0 {
		goto L1
	} else {
		goto L801
	}
L801:
	;
	v4686 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_123), v4682, v4672+int32(15))
	mBase = m.M
	v4687 = m.ExcPending
	if v4687 != 0 {
		goto L1
	} else {
		goto L802
	}
L802:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[221])) = v4686
	v4690 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v4690 != 0 {
		goto L803
	} else {
		goto L804
	}
L803:
	;
	v4815 = int32(16)
	m.G0 = v4672 + v4815
	v4818 = int32(0)
	v4819 = m.G0
	v4821 = v4819 - v4815
	m.G0 = v4821
	v4828 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4831 = F_mul_size(m, v4815, v4828+int32(38))
	mBase = m.M
	v4832 = m.ExcPending
	if v4832 != 0 {
		goto L1
	} else {
		goto L815
	}
L804:
	;
	v4692 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[207]))
	*(*int64)(unsafe.Add(mBase, uint32(v4686)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4686))) = v4692
	v4697 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[222]))
	v4698 = int32(0)
	if base.B2i32(v4697 == v4698)|base.B2i32(v4697 == int32(_a_F_CreateSharedMemoryAndSemaphores_124)) == v4698 {
		goto L805
	} else {
		goto L806
	}
L805:
	;
	v4705 = v4697
	v4706 = v4669
	goto L808
L806:
	;
	v4753 = v4669
	v4754 = v4692
	goto L807
L807:
	;
	if v4754 <= v4753 {
		goto L803
	} else {
		goto L811
	}
L808:
	;
	v4722 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[221]))
	v4723 = int32(1480)
	v4725 = v4722 + v4706*v4723
	*(*int64)(unsafe.Add(mBase, uint32(v4725)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4725)+20)) = int32(-1)
	v4730 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v4725)+16)) = uint16(v4730)
	*(*int32)(unsafe.Add(mBase, uint32(v4705-int32(24)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4705-int32(8)))) = v4706
	base.MemoryCopy(m, v4725+int32(32), v4705-v4723, int32(1460))
	v4746 = v4706 + v4730
	v4747 = *(*int32)(unsafe.Add(mBase, uint32(v4705)+4))
	if v4747 != int32(_a_F_CreateSharedMemoryAndSemaphores_124) {
		v4705 = v4747
		v4706 = v4746
		goto L808
	} else {
		goto L810
	}
L809:
	;
	v4751 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[207]))
	v4753 = v4746
	v4754 = v4751
	goto L807
L810:
	;
	goto L809
L811:
	;
	v4770 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[221]))
	v4774 = v4753
	goto L812
L812:
	;
	v4792 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4770+int32(16)+v4774*int32(1480)))) = uint8(v4792)
	v4795 = v4774 + int32(1)
	v4797 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[207]))
	if v4795 < v4797 {
		v4774 = v4795
		goto L812
	} else {
		goto L814
	}
L813:
	;
	goto L803
L814:
	;
	goto L813
L815:
	;
	v4833 = F_add_size(m, int32(_a_F_CreateSharedMemoryAndSemaphores_125), v4831)
	mBase = m.M
	v4834 = m.ExcPending
	if v4834 != 0 {
		goto L1
	} else {
		goto L816
	}
L816:
	;
	v4837 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v4840 = F_mul_size(m, int32(4), v4837+int32(38))
	mBase = m.M
	v4841 = m.ExcPending
	if v4841 != 0 {
		goto L1
	} else {
		goto L817
	}
L817:
	;
	v4842 = F_add_size(m, v4833, v4840)
	mBase = m.M
	v4843 = m.ExcPending
	if v4843 != 0 {
		goto L1
	} else {
		goto L818
	}
L818:
	;
	v4846 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_126), v4842, v4821+int32(15))
	mBase = m.M
	v4847 = m.ExcPending
	if v4847 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[223])) = v4846
	v4849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4821)+15)))
	if v4849 == int32(0) {
		goto L820
	} else {
		goto L821
	}
L820:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4846)+8)) = int64(2048)
	*(*int64)(unsafe.Add(mBase, uint32(v4846))) = int64(0)
	v4857 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if int32(0) < v4857+int32(38) {
		goto L823
	} else {
		goto L824
	}
L821:
	;
	goto L822
L822:
	;
	v4945 = int32(16)
	m.G0 = v4821 + v4945
	v4948 = m.G0
	v4950 = v4948 - v4945
	m.G0 = v4950
	v4955 = F_MaxLivePostmasterChildren(m)
	mBase = m.M
	v4956 = m.ExcPending
	if v4956 != 0 {
		goto L1
	} else {
		goto L829
	}
L823:
	;
	v4866 = v4818
	goto L826
L824:
	;
	v4907 = v4818
	goto L825
L825:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4846)+uint32(_c_F_CreateSharedMemoryAndSemaphores[224]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4846)+uint32(_c_F_CreateSharedMemoryAndSemaphores[225]))) = v4846 + v4907<<(uint(int32(4))%32) + int32(_a_F_CreateSharedMemoryAndSemaphores_125)
	goto L822
L826:
	;
	v4881 = v4866 << (uint(int32(4)) % 32)
	v4883 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4846+int32(_a_F_CreateSharedMemoryAndSemaphores_125)+v4881))) = v4883
	v4885 = v4846 + v4881
	*(*int32)(unsafe.Add(mBase, uint32(v4885)+uint32(_c_F_CreateSharedMemoryAndSemaphores[226]))) = v4883
	*(*int32)(unsafe.Add(mBase, uint32(v4885)+uint32(_c_F_CreateSharedMemoryAndSemaphores[227]))) = v4883
	*(*int32)(unsafe.Add(mBase, uint32(v4885)+uint32(_c_F_CreateSharedMemoryAndSemaphores[228]))) = v4883
	v4899 = v4866 + int32(1)
	v4901 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v4899 < v4901+int32(38) {
		v4866 = v4899
		goto L826
	} else {
		goto L828
	}
L827:
	;
	v4907 = v4899
	goto L825
L828:
	;
	goto L827
L829:
	;
	v4958 = F_mul_size(m, v4955, int32(4))
	mBase = m.M
	v4959 = m.ExcPending
	if v4959 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	v4960 = F_add_size(m, int32(48), v4958)
	mBase = m.M
	v4961 = m.ExcPending
	if v4961 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	v4964 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_127), v4960, v4950+int32(15))
	mBase = m.M
	v4965 = m.ExcPending
	if v4965 != 0 {
		goto L1
	} else {
		goto L832
	}
L832:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[229])) = v4964
	v4967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4950)+15)))
	if v4967 == int32(0) {
		goto L833
	} else {
		goto L834
	}
L833:
	;
	v4973 = F_MaxLivePostmasterChildren(m)
	mBase = m.M
	v4974 = m.ExcPending
	if v4974 != 0 {
		goto L1
	} else {
		goto L837
	}
L834:
	;
	goto L835
L835:
	;
	v5017 = int32(16)
	m.G0 = v4950 + v5017
	v5021 = m.G0
	v5023 = v5021 - v5017
	m.G0 = v5023
	v5028 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v5032 = F_mul_size(m, v5028+int32(38), int32(128))
	mBase = m.M
	v5033 = m.ExcPending
	if v5033 != 0 {
		goto L1
	} else {
		goto L849
	}
L836:
	;
	v5010 = F_MaxLivePostmasterChildren(m)
	mBase = m.M
	v5011 = m.ExcPending
	if v5011 != 0 {
		goto L1
	} else {
		goto L848
	}
L837:
	;
	v4976 = F_mul_size(m, v4973, int32(4))
	mBase = m.M
	v4977 = m.ExcPending
	if v4977 != 0 {
		goto L1
	} else {
		goto L838
	}
L838:
	;
	v4978 = F_add_size(m, int32(48), v4976)
	mBase = m.M
	v4979 = m.ExcPending
	if v4979 != 0 {
		goto L1
	} else {
		goto L839
	}
L839:
	;
	if v4964&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v4978))|v4978&int32(3) == int32(0) {
		goto L840
	} else {
		goto L841
	}
L840:
	;
	if v4978 == int32(0) {
		goto L836
	} else {
		goto L843
	}
L841:
	;
	v5002 = v4978
	goto L842
L842:
	;
	if v5002 == int32(0) {
		goto L836
	} else {
		goto L847
	}
L843:
	;
	v4992 = v4978 + v4964
	v4994 = v4964 + int32(4)
	if base.Ui32(v4994) < base.Ui32(v4992) {
		goto L844
	} else {
		goto L845
	}
L844:
	;
	v4996 = v4992
	goto L846
L845:
	;
	v4996 = v4994
	goto L846
L846:
	;
	v5002 = (v4964^int32(-1)+v4996)&int32(-4) + int32(4)
	goto L842
L847:
	;
	base.MemoryFill(m, v4964, int32(0), v5002)
	goto L836
L848:
	;
	v5013 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[229]))
	*(*int32)(unsafe.Add(mBase, uint32(v5013)+44)) = v5010
	goto L835
L849:
	;
	v5035 = F_add_size(m, v5032, int32(8))
	mBase = m.M
	v5036 = m.ExcPending
	if v5036 != 0 {
		goto L1
	} else {
		goto L850
	}
L850:
	;
	v5039 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_128), v5035, v5023+int32(15))
	mBase = m.M
	v5040 = m.ExcPending
	if v5040 != 0 {
		goto L1
	} else {
		goto L851
	}
L851:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[230])) = v5039
	v5042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5023)+15)))
	if v5042 != 0 {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	v5127 = int32(16)
	m.G0 = v5023 + v5127
	v5130 = m.G0
	v5132 = v5130 - v5127
	m.G0 = v5132
	v5137 = int32(_a_F_CreateSharedMemoryAndSemaphores_129)
	v5139 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	if v5137 <= v5139 {
		goto L859
	} else {
		goto L860
	}
L853:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5039))) = int64(0)
	v5046 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v5046+int32(38) <= int32(0) {
		goto L852
	} else {
		goto L854
	}
L854:
	;
	v5052 = int32(0)
	goto L855
L855:
	;
	v5068 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[230]))
	v5071 = v5068 + v5052<<(uint(int32(7))%32)
	v5072 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5071)+8)) = v5072
	*(*int64)(unsafe.Add(mBase, uint32(v5071)+112)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5071)+104)) = v5072
	*(*int32)(unsafe.Add(mBase, uint32(v5071)+12)) = v5072
	v5080 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5071)+96)) = v5080
	*(*int64)(unsafe.Add(mBase, uint32(v5071)+88)) = v5080
	*(*int64)(unsafe.Add(mBase, uint32(v5071)+80)) = v5080
	*(*int64)(unsafe.Add(mBase, uint32(v5071)+72)) = v5080
	*(*int64)(unsafe.Add(mBase, uint32(v5071-int32(-64)))) = v5080
	*(*int64)(unsafe.Add(mBase, uint32(v5071)+56)) = v5080
	*(*int64)(unsafe.Add(mBase, uint32(v5071)+48)) = v5080
	*(*int32)(unsafe.Add(mBase, uint32(v5071)+120)) = v5072
	v5099 = v5071 + int32(124)
	*(*int32)(unsafe.Add(mBase, uint32(v5099)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5099))) = int64(-4294967296)
	goto L857
L856:
	;
	goto L852
L857:
	;
	v5105 = v5052 + int32(1)
	v5107 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v5105 < v5107+int32(38) {
		v5052 = v5105
		goto L855
	} else {
		goto L858
	}
L858:
	;
	goto L856
L859:
	;
	v5142 = v5137
	goto L861
L860:
	;
	v5142 = v5139
	goto L861
L861:
	;
	v5144 = F_mul_size(m, v5142, int32(32))
	mBase = m.M
	v5145 = m.ExcPending
	if v5145 != 0 {
		goto L1
	} else {
		goto L862
	}
L862:
	;
	v5146 = F_add_size(m, int32(56), v5144)
	mBase = m.M
	v5147 = m.ExcPending
	if v5147 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	v5150 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_130), v5146, v5132+int32(15))
	mBase = m.M
	v5151 = m.ExcPending
	if v5151 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[231])) = v5150
	v5153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5132)+15)))
	if v5153 == int32(0) {
		goto L865
	} else {
		goto L866
	}
L865:
	;
	v5156 = int32(3)
	if v5146&v5156|(v5150&v5156|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v5146))) == int32(0) {
		goto L869
	} else {
		goto L870
	}
L866:
	;
	goto L867
L867:
	;
	v5213 = int32(16)
	m.G0 = v5132 + v5213
	v5216 = int32(0)
	v5218 = m.G0
	v5220 = v5218 - v5213
	m.G0 = v5220
	v5226 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[206]))
	v5228 = F_mul_size(m, v5226, int32(40))
	mBase = m.M
	v5229 = m.ExcPending
	if v5229 != 0 {
		goto L1
	} else {
		goto L882
	}
L868:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5150)+4)) = int32(0)
	v5190 = int32(_a_F_CreateSharedMemoryAndSemaphores_129)
	v5192 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[161]))
	if v5190 <= v5192 {
		goto L877
	} else {
		goto L878
	}
L869:
	;
	if v5146 == int32(0) {
		goto L868
	} else {
		goto L872
	}
L870:
	;
	v5180 = v5146
	goto L871
L871:
	;
	if v5180 == int32(0) {
		goto L868
	} else {
		goto L876
	}
L872:
	;
	v5170 = v5146 + v5150
	v5172 = v5150 + int32(4)
	if base.Ui32(v5172) < base.Ui32(v5170) {
		goto L873
	} else {
		goto L874
	}
L873:
	;
	v5174 = v5170
	goto L875
L874:
	;
	v5174 = v5172
	goto L875
L875:
	;
	v5180 = (v5150^int32(-1)+v5174)&int32(-4) + int32(4)
	goto L871
L876:
	;
	base.MemoryFill(m, v5150, int32(0), v5180)
	goto L868
L877:
	;
	v5195 = v5190
	goto L879
L878:
	;
	v5195 = v5192
	goto L879
L879:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5150)+52)) = v5195
	v5198 = v5150 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v5198)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5198))) = int64(-4294967296)
	goto L880
L880:
	;
	v5204 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[231]))
	v5206 = v5204 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v5206)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5206))) = int64(-4294967296)
	goto L881
L881:
	;
	goto L867
L882:
	;
	v5230 = F_add_size(m, int32(_a_F_CreateSharedMemoryAndSemaphores_131), v5228)
	mBase = m.M
	v5231 = m.ExcPending
	if v5231 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	v5234 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_132), v5230, v5220+int32(15))
	mBase = m.M
	v5235 = m.ExcPending
	if v5235 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[232])) = v5234
	v5238 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v5238 == int32(0) {
		goto L885
	} else {
		goto L886
	}
L885:
	;
	v5241 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+20)) = v5241
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+8)) = v5241
	v5246 = v5234 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+28)) = v5246
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+24)) = v5246
	v5250 = v5234 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+16)) = v5250
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+12)) = v5250
	base.MemoryFill(m, v5234+int32(32), v5241, int32(_a_F_CreateSharedMemoryAndSemaphores_133))
	v5259 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[206]))
	if v5241 < v5259 {
		goto L888
	} else {
		goto L889
	}
L886:
	;
	goto L887
L887:
	;
	v5331 = int32(16)
	m.G0 = v5220 + v5331
	v5334 = m.G0
	v5336 = v5334 - v5331
	m.G0 = v5336
	v5339 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[233]))
	if v5339 == int32(0) {
		goto L894
	} else {
		goto L895
	}
L888:
	;
	v5264 = v5250
	v5267 = v5216
	v5269 = v5216
	goto L891
L889:
	;
	goto L890
L890:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+uint32(_c_F_CreateSharedMemoryAndSemaphores[234]))) = int32(0)
	goto L887
L891:
	;
	v5282 = v5234 + int32(_a_F_CreateSharedMemoryAndSemaphores_131) + v5267*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v5282))) = v5250
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+4)) = v5264
	*(*int32)(unsafe.Add(mBase, uint32(v5264))) = v5282
	v5286 = int32(1)
	v5287 = v5269 + v5286
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+20)) = v5287
	*(*int32)(unsafe.Add(mBase, uint32(v5234)+16)) = v5282
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+32)) = int32(0)
	v5293 = v5267 + v5286
	v5295 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[206]))
	if v5293 < v5295 {
		v5264 = v5282
		v5267 = v5293
		v5269 = v5287
		goto L891
	} else {
		goto L893
	}
L892:
	;
	goto L890
L893:
	;
	goto L892
L894:
	;
	v5461 = int32(16)
	m.G0 = v5336 + v5461
	v5464 = int32(0)
	v5465 = m.G0
	v5467 = v5465 - v5461
	m.G0 = v5467
	v5470 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[235]))
	if v5470 == v5464 {
		goto L920
	} else {
		goto L921
	}
L895:
	;
	v5346 = F_mul_size(m, v5339, int32(288))
	mBase = m.M
	v5347 = m.ExcPending
	if v5347 != 0 {
		goto L1
	} else {
		goto L896
	}
L896:
	;
	v5348 = F_add_size(m, int32(0), v5346)
	mBase = m.M
	v5349 = m.ExcPending
	if v5349 != 0 {
		goto L1
	} else {
		goto L897
	}
L897:
	;
	v5352 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_134), v5348, v5336+int32(15))
	mBase = m.M
	v5353 = m.ExcPending
	if v5353 != 0 {
		goto L1
	} else {
		goto L898
	}
L898:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[236])) = v5352
	v5355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5336)+15)))
	if v5355 != 0 {
		goto L894
	} else {
		goto L899
	}
L899:
	;
	v5358 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[233]))
	if v5358 != 0 {
		goto L900
	} else {
		goto L901
	}
L900:
	;
	v5361 = F_mul_size(m, v5358, int32(288))
	mBase = m.M
	v5362 = m.ExcPending
	if v5362 != 0 {
		goto L1
	} else {
		goto L903
	}
L901:
	;
	v5365 = int32(0)
	goto L902
L902:
	;
	v5366 = int32(3)
	if v5365&v5366|(v5352&v5366|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v5365))) == int32(0) {
		goto L906
	} else {
		goto L907
	}
L903:
	;
	v5363 = F_add_size(m, int32(0), v5361)
	mBase = m.M
	v5364 = m.ExcPending
	if v5364 != 0 {
		goto L1
	} else {
		goto L904
	}
L904:
	;
	v5365 = v5363
	goto L902
L905:
	;
	v5398 = int32(0)
	v5400 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[233]))
	if v5400 <= v5398 {
		goto L894
	} else {
		goto L914
	}
L906:
	;
	if v5365 == int32(0) {
		goto L905
	} else {
		goto L909
	}
L907:
	;
	v5390 = v5365
	goto L908
L908:
	;
	if v5390 == int32(0) {
		goto L905
	} else {
		goto L913
	}
L909:
	;
	v5380 = v5365 + v5352
	v5382 = v5352 + int32(4)
	if base.Ui32(v5382) < base.Ui32(v5380) {
		goto L910
	} else {
		goto L911
	}
L910:
	;
	v5384 = v5380
	goto L912
L911:
	;
	v5384 = v5382
	goto L912
L912:
	;
	v5390 = (v5352^int32(-1)+v5384)&int32(-4) + int32(4)
	goto L908
L913:
	;
	base.MemoryFill(m, v5352, int32(0), v5390)
	goto L905
L914:
	;
	v5403 = v5398
	goto L915
L915:
	;
	v5420 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[236]))
	v5423 = v5420 + v5403*int32(288)
	*(*int32)(unsafe.Add(mBase, uint32(v5423))) = int32(0)
	v5427 = v5423 + int32(208)
	v5428 = int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v5427))) = uint16(v5428)
	*(*int32)(unsafe.Add(mBase, uint32(v5427)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v5427)+8)) = int64(-1)
	goto L917
L916:
	;
	goto L894
L917:
	;
	v5435 = v5423 + int32(224)
	*(*int32)(unsafe.Add(mBase, uint32(v5435)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5435))) = int64(-4294967296)
	goto L918
L918:
	;
	v5441 = v5403 + int32(1)
	v5443 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[233]))
	if v5441 < v5443 {
		v5403 = v5441
		goto L915
	} else {
		goto L919
	}
L919:
	;
	goto L916
L920:
	;
	v5612 = int32(16)
	m.G0 = v5467 + v5612
	v5615 = m.G0
	v5617 = v5615 - v5612
	m.G0 = v5617
	v5623 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[237]))
	v5625 = F_mul_size(m, v5623, int32(96))
	mBase = m.M
	v5626 = m.ExcPending
	if v5626 != 0 {
		goto L1
	} else {
		goto L948
	}
L921:
	;
	v5477 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v5478 = m.ExcPending
	if v5478 != 0 {
		goto L1
	} else {
		goto L922
	}
L922:
	;
	v5480 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[235]))
	v5482 = F_mul_size(m, v5480, int32(56))
	mBase = m.M
	v5483 = m.ExcPending
	if v5483 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	v5484 = F_add_size(m, v5477, v5482)
	mBase = m.M
	v5485 = m.ExcPending
	if v5485 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	v5488 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_135), v5484, v5467+int32(15))
	mBase = m.M
	v5489 = m.ExcPending
	if v5489 != 0 {
		goto L1
	} else {
		goto L925
	}
L925:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[238])) = v5488
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[239])) = v5488 + int32(8)
	v5495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5467)+15)))
	if v5495 != 0 {
		goto L920
	} else {
		goto L926
	}
L926:
	;
	v5497 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[235]))
	if v5497 != 0 {
		goto L927
	} else {
		goto L928
	}
L927:
	;
	v5500 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v5501 = m.ExcPending
	if v5501 != 0 {
		goto L1
	} else {
		goto L930
	}
L928:
	;
	v5509 = v5464
	goto L929
L929:
	;
	v5510 = int32(3)
	if v5509&v5510|(v5488&v5510|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v5509))) == int32(0) {
		goto L934
	} else {
		goto L935
	}
L930:
	;
	v5503 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[235]))
	v5505 = F_mul_size(m, v5503, int32(56))
	mBase = m.M
	v5506 = m.ExcPending
	if v5506 != 0 {
		goto L1
	} else {
		goto L931
	}
L931:
	;
	v5507 = F_add_size(m, v5500, v5505)
	mBase = m.M
	v5508 = m.ExcPending
	if v5508 != 0 {
		goto L1
	} else {
		goto L932
	}
L932:
	;
	v5509 = v5507
	goto L929
L933:
	;
	v5542 = int32(0)
	v5544 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[238]))
	*(*int32)(unsafe.Add(mBase, uint32(v5544))) = int32(63)
	v5548 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[235]))
	if v5548 <= v5542 {
		goto L920
	} else {
		goto L942
	}
L934:
	;
	if v5509 == int32(0) {
		goto L933
	} else {
		goto L937
	}
L935:
	;
	v5534 = v5509
	goto L936
L936:
	;
	if v5534 == int32(0) {
		goto L933
	} else {
		goto L941
	}
L937:
	;
	v5524 = v5509 + v5488
	v5526 = v5488 + int32(4)
	if base.Ui32(v5526) < base.Ui32(v5524) {
		goto L938
	} else {
		goto L939
	}
L938:
	;
	v5528 = v5524
	goto L940
L939:
	;
	v5528 = v5526
	goto L940
L940:
	;
	v5534 = (v5488^int32(-1)+v5528)&int32(-4) + int32(4)
	goto L936
L941:
	;
	base.MemoryFill(m, v5488, int32(0), v5534)
	goto L933
L942:
	;
	v5552 = v5542
	goto L943
L943:
	;
	v5568 = v5552 * int32(56)
	v5570 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[239]))
	v5573 = v5568 + v5570 + int32(40)
	v5575 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[238]))
	v5576 = *(*int32)(unsafe.Add(mBase, uint32(v5575)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5573))) = uint16(v5576)
	*(*int32)(unsafe.Add(mBase, uint32(v5573)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v5573)+8)) = int64(-1)
	goto L945
L944:
	;
	goto L920
L945:
	;
	v5583 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[239]))
	v5586 = v5583 + v5568 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v5586)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5586))) = int64(-4294967296)
	goto L946
L946:
	;
	v5592 = v5552 + int32(1)
	v5594 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[235]))
	if v5592 < v5594 {
		v5552 = v5592
		goto L943
	} else {
		goto L947
	}
L947:
	;
	goto L944
L948:
	;
	v5627 = F_add_size(m, int32(88), v5625)
	mBase = m.M
	v5628 = m.ExcPending
	if v5628 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	v5631 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_136), v5627, v5617+int32(15))
	mBase = m.M
	v5632 = m.ExcPending
	if v5632 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[240])) = v5631
	v5634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5617)+15)))
	if v5634 == int32(0) {
		goto L951
	} else {
		goto L952
	}
L951:
	;
	v5641 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[237]))
	v5643 = F_mul_size(m, v5641, int32(96))
	mBase = m.M
	v5644 = m.ExcPending
	if v5644 != 0 {
		goto L1
	} else {
		goto L955
	}
L952:
	;
	goto L953
L953:
	;
	v5779 = int32(16)
	m.G0 = v5617 + v5779
	v5782 = m.G0
	v5784 = v5782 - v5779
	m.G0 = v5784
	v5790 = F_add_size(m, int32(0), int32(1480))
	mBase = m.M
	v5791 = m.ExcPending
	if v5791 != 0 {
		goto L1
	} else {
		goto L974
	}
L954:
	;
	v5677 = int32(0)
	v5679 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[240]))
	v5681 = v5679 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+20)) = v5681
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+16)) = v5681
	v5685 = v5679 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+12)) = v5685
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+8)) = v5685
	*(*int32)(unsafe.Add(mBase, uint32(v5679)+4)) = v5679
	*(*int32)(unsafe.Add(mBase, uint32(v5679))) = v5679
	v5691 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[237]))
	if v5677 < v5691 {
		goto L965
	} else {
		goto L966
	}
L955:
	;
	v5645 = F_add_size(m, int32(88), v5643)
	mBase = m.M
	v5646 = m.ExcPending
	if v5646 != 0 {
		goto L1
	} else {
		goto L956
	}
L956:
	;
	if v5631&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v5645))|v5645&int32(3) == int32(0) {
		goto L957
	} else {
		goto L958
	}
L957:
	;
	if v5645 == int32(0) {
		goto L954
	} else {
		goto L960
	}
L958:
	;
	v5669 = v5645
	goto L959
L959:
	;
	if v5669 == int32(0) {
		goto L954
	} else {
		goto L964
	}
L960:
	;
	v5659 = v5645 + v5631
	v5661 = v5631 + int32(4)
	if base.Ui32(v5661) < base.Ui32(v5659) {
		goto L961
	} else {
		goto L962
	}
L961:
	;
	v5663 = v5659
	goto L963
L962:
	;
	v5663 = v5661
	goto L963
L963:
	;
	v5669 = (v5631^int32(-1)+v5663)&int32(-4) + int32(4)
	goto L959
L964:
	;
	base.MemoryFill(m, v5631, int32(0), v5669)
	goto L954
L965:
	;
	v5695 = v5677
	goto L968
L966:
	;
	v5740 = v5679
	goto L967
L967:
	;
	v5742 = v5740 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v5742)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5742))) = int64(-4294967296)
	goto L971
L968:
	;
	v5711 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[240]))
	*(*int32)(unsafe.Add(mBase, uint32(v5711+v5695*int32(96))+164)) = int32(0)
	v5718 = v5695 + int32(1)
	v5720 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[237]))
	if v5718 < v5720 {
		v5695 = v5718
		goto L968
	} else {
		goto L970
	}
L969:
	;
	v5723 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[240]))
	v5740 = v5723
	goto L967
L970:
	;
	goto L969
L971:
	;
	v5748 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[240]))
	v5750 = v5748 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v5750)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5750))) = int64(-4294967296)
	goto L972
L972:
	;
	v5756 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[240]))
	v5758 = v5756 + int32(76)
	*(*int32)(unsafe.Add(mBase, uint32(v5758)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5758))) = int64(-4294967296)
	goto L973
L973:
	;
	goto L953
L974:
	;
	v5794 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_137), v5790, v5784+int32(15))
	mBase = m.M
	v5795 = m.ExcPending
	if v5795 != 0 {
		goto L1
	} else {
		goto L975
	}
L975:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[241])) = v5794
	v5797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5784)+15)))
	if v5797 == int32(0) {
		goto L976
	} else {
		goto L977
	}
L976:
	;
	v5804 = F_add_size(m, int32(0), int32(1480))
	mBase = m.M
	v5805 = m.ExcPending
	if v5805 != 0 {
		goto L1
	} else {
		goto L980
	}
L977:
	;
	goto L978
L978:
	;
	v5857 = int32(16)
	m.G0 = v5784 + v5857
	v5860 = m.G0
	v5862 = v5860 - v5857
	m.G0 = v5862
	v5869 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_138), int32(48), v5862+int32(15))
	mBase = m.M
	v5870 = m.ExcPending
	if v5870 != 0 {
		goto L1
	} else {
		goto L990
	}
L979:
	;
	v5837 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[241]))
	*(*int32)(unsafe.Add(mBase, uint32(v5837)+8)) = int32(0)
	v5841 = v5837 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v5841)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5841))) = int64(-4294967296)
	goto L989
L980:
	;
	if v5794&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v5804))|v5804&int32(3) == int32(0) {
		goto L981
	} else {
		goto L982
	}
L981:
	;
	if v5804 == int32(0) {
		goto L979
	} else {
		goto L984
	}
L982:
	;
	v5828 = v5804
	goto L983
L983:
	;
	if v5828 == int32(0) {
		goto L979
	} else {
		goto L988
	}
L984:
	;
	v5818 = v5794 + v5804
	v5820 = v5794 + int32(4)
	if base.Ui32(v5820) < base.Ui32(v5818) {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v5822 = v5818
	goto L987
L986:
	;
	v5822 = v5820
	goto L987
L987:
	;
	v5828 = (v5794^int32(-1)+v5822)&int32(-4) + int32(4)
	goto L983
L988:
	;
	base.MemoryFill(m, v5794, int32(0), v5828)
	goto L979
L989:
	;
	v5847 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[241]))
	*(*int64)(unsafe.Add(mBase, uint32(v5847)+1464)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5847)+1456)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5847))) = int32(-1)
	goto L978
L990:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[242])) = v5869
	v5872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5862)+15)))
	if v5872 == int32(0) {
		goto L991
	} else {
		goto L992
	}
L991:
	;
	v5875 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5869)+24)) = v5875
	v5877 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5869)+20)) = v5877
	v5879 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5869)+16)) = uint8(v5879)
	*(*int64)(unsafe.Add(mBase, uint32(v5869)+8)) = v5875
	*(*int32)(unsafe.Add(mBase, uint32(v5869)+4)) = v5879
	*(*uint8)(unsafe.Add(mBase, uint32(v5869))) = uint8(v5879)
	v5888 = v5869 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v5888)+8)) = v5877
	*(*int64)(unsafe.Add(mBase, uint32(v5888))) = int64(-4294967296)
	goto L994
L992:
	;
	goto L993
L993:
	;
	v5893 = int32(16)
	m.G0 = v5862 + v5893
	v5896 = m.G0
	v5898 = v5896 - v5893
	m.G0 = v5898
	v5904 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v5905 = m.ExcPending
	if v5905 != 0 {
		goto L1
	} else {
		goto L995
	}
L994:
	;
	goto L993
L995:
	;
	v5908 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_139), v5904, v5898+int32(15))
	mBase = m.M
	v5909 = m.ExcPending
	if v5909 != 0 {
		goto L1
	} else {
		goto L996
	}
L996:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[243])) = v5908
	v5911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5898)+15)))
	if v5911 == int32(0) {
		goto L997
	} else {
		goto L998
	}
L997:
	;
	v5918 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v5919 = m.ExcPending
	if v5919 != 0 {
		goto L1
	} else {
		goto L1001
	}
L998:
	;
	goto L999
L999:
	;
	v5958 = int32(16)
	m.G0 = v5898 + v5958
	v5961 = m.G0
	v5963 = v5961 - v5958
	m.G0 = v5963
	v5969 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[244]))
	v5971 = F_mul_size(m, v5969, int32(112))
	mBase = m.M
	v5972 = m.ExcPending
	if v5972 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1000:
	;
	v5951 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[243]))
	*(*int32)(unsafe.Add(mBase, uint32(v5951)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5951))) = int32(-1)
	goto L999
L1001:
	;
	if v5908&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v5918))|v5918&int32(3) == int32(0) {
		goto L1002
	} else {
		goto L1003
	}
L1002:
	;
	if v5918 == int32(0) {
		goto L1000
	} else {
		goto L1005
	}
L1003:
	;
	v5942 = v5918
	goto L1004
L1004:
	;
	if v5942 == int32(0) {
		goto L1000
	} else {
		goto L1009
	}
L1005:
	;
	v5932 = v5918 + v5908
	v5934 = v5908 + int32(4)
	if base.Ui32(v5934) < base.Ui32(v5932) {
		goto L1006
	} else {
		goto L1007
	}
L1006:
	;
	v5936 = v5932
	goto L1008
L1007:
	;
	v5936 = v5934
	goto L1008
L1008:
	;
	v5942 = (v5908^int32(-1)+v5936)&int32(-4) + int32(4)
	goto L1004
L1009:
	;
	base.MemoryFill(m, v5908, int32(0), v5942)
	goto L1000
L1010:
	;
	v5973 = F_add_size(m, v5958, v5971)
	mBase = m.M
	v5974 = m.ExcPending
	if v5974 != 0 {
		goto L1
	} else {
		goto L1011
	}
L1011:
	;
	v5977 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_140), v5973, v5963+int32(15))
	mBase = m.M
	v5978 = m.ExcPending
	if v5978 != 0 {
		goto L1
	} else {
		goto L1012
	}
L1012:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[245])) = v5977
	v5980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5963)+15)))
	if v5980 != 0 {
		goto L1013
	} else {
		goto L1014
	}
L1013:
	;
	v6047 = int32(16)
	m.G0 = v5963 + v6047
	v6050 = m.G0
	v6052 = v6050 - v6047
	m.G0 = v6052
	v6059 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_141), int32(24), v6052+int32(15))
	mBase = m.M
	v6060 = m.ExcPending
	if v6060 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1014:
	;
	v5983 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[244]))
	v5985 = F_mul_size(m, v5983, int32(112))
	mBase = m.M
	v5986 = m.ExcPending
	if v5986 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	v5987 = F_add_size(m, int32(16), v5985)
	mBase = m.M
	v5988 = m.ExcPending
	if v5988 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	if v5987 != 0 {
		goto L1017
	} else {
		goto L1018
	}
L1017:
	;
	base.MemoryFill(m, v5977, int32(0), v5987)
	goto L1019
L1018:
	;
	goto L1019
L1019:
	;
	v5992 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[245]))
	*(*int64)(unsafe.Add(mBase, uint32(v5992)+4)) = int64(0)
	v5996 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[244]))
	if v5996 <= int32(0) {
		goto L1013
	} else {
		goto L1020
	}
L1020:
	;
	v6002 = int32(0)
	goto L1021
L1021:
	;
	v6018 = int32(112)
	v6020 = v5992 + int32(16) + v6002*v6018
	v6021 = int32(0)
	base.MemoryFill(m, v6020, v6021, v6018)
	*(*int32)(unsafe.Add(mBase, uint32(v6020)+56)) = v6021
	v6027 = v6002 + int32(1)
	v6029 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[244]))
	if v6027 < v6029 {
		v6002 = v6027
		goto L1021
	} else {
		goto L1023
	}
L1022:
	;
	goto L1013
L1023:
	;
	goto L1022
L1024:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[246])) = v6059
	v6062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6052)+15)))
	if v6062 == int32(0) {
		goto L1025
	} else {
		goto L1026
	}
L1025:
	;
	v6065 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6059)+16)) = v6065
	*(*int64)(unsafe.Add(mBase, uint32(v6059))) = v6065
	*(*int64)(unsafe.Add(mBase, uint32(v6059)+8)) = v6065
	*(*int32)(unsafe.Add(mBase, uint32(v6059)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6059))) = int32(-1)
	goto L1027
L1026:
	;
	goto L1027
L1027:
	;
	v6075 = int32(16)
	m.G0 = v6052 + v6075
	v6078 = m.G0
	v6080 = v6078 - v6075
	m.G0 = v6080
	v6084 = int32(12)
	v6086 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v6088 = F_mul_size(m, v6086, v6084)
	mBase = m.M
	v6089 = m.ExcPending
	if v6089 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1028:
	;
	v6090 = F_add_size(m, v6084, v6088)
	mBase = m.M
	v6091 = m.ExcPending
	if v6091 != 0 {
		goto L1
	} else {
		goto L1029
	}
L1029:
	;
	v6094 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_142), v6090, v6080+int32(15))
	mBase = m.M
	v6095 = m.ExcPending
	if v6095 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[247])) = v6094
	v6098 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v6098 == int32(0) {
		goto L1031
	} else {
		goto L1032
	}
L1031:
	;
	v6101 = F_time(m)
	mBase = m.M
	v6103 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[247]))
	*(*int32)(unsafe.Add(mBase, uint32(v6103)+4)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v6103))) = uint16(v6101)
	v6108 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	*(*int32)(unsafe.Add(mBase, uint32(v6103)+8)) = v6108
	goto L1033
L1032:
	;
	goto L1033
L1033:
	;
	v6112 = int32(16)
	m.G0 = v6080 + v6112
	v6115 = m.G0
	v6117 = v6115 - v6112
	m.G0 = v6117
	v6124 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_143), int32(488), v6117+int32(15))
	mBase = m.M
	v6125 = m.ExcPending
	if v6125 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1034:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[248])) = v6124
	v6128 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v6128 == int32(0) {
		goto L1035
	} else {
		goto L1036
	}
L1035:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6124)+24)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v6124)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6124)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6124)+4)) = v6124 + int32(464)
	v6141 = v6124 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v6124))) = v6141
	*(*int32)(unsafe.Add(mBase, uint32(v6124)+12)) = v6124 + int32(32)
	v6147 = v6124 - int32(16)
	v6150 = int32(1)
	goto L1038
L1036:
	;
	goto L1037
L1037:
	;
	v6208 = int32(16)
	m.G0 = v6117 + v6208
	v6212 = m.G0
	v6214 = v6212 - v6208
	m.G0 = v6214
	v6219 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v6221 = F_mul_size(m, v6219, int32(32))
	mBase = m.M
	v6222 = m.ExcPending
	if v6222 != 0 {
		goto L1
	} else {
		goto L1041
	}
L1038:
	;
	v6165 = int32(24)
	v6166 = v6150 * v6165
	v6167 = v6141 + v6166
	*(*int64)(unsafe.Add(mBase, uint32(v6167)+16)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v6167)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6167))) = v6147 + v6166
	*(*int32)(unsafe.Add(mBase, uint32(v6167)+4)) = v6167 + v6165
	v6178 = v6150 + int32(1)
	if v6178 != int32(19) {
		v6150 = v6178
		goto L1038
	} else {
		goto L1040
	}
L1039:
	;
	v6182 = v6178 * int32(24)
	v6183 = v6141 + v6182
	*(*int32)(unsafe.Add(mBase, uint32(v6183)+20)) = int32(-1)
	v6186 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6183)+12)) = v6186
	*(*int64)(unsafe.Add(mBase, uint32(v6183)+4)) = v6186
	*(*int32)(unsafe.Add(mBase, uint32(v6183))) = v6182 + v6147
	goto L1037
L1040:
	;
	goto L1039
L1041:
	;
	v6224 = F_add_size(m, v6221, int32(56))
	mBase = m.M
	v6225 = m.ExcPending
	if v6225 != 0 {
		goto L1
	} else {
		goto L1042
	}
L1042:
	;
	v6228 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_144), v6224, v6214+int32(15))
	mBase = m.M
	v6229 = m.ExcPending
	if v6229 != 0 {
		goto L1
	} else {
		goto L1043
	}
L1043:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[249])) = v6228
	v6231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6214)+15)))
	if v6231 != 0 {
		goto L1044
	} else {
		goto L1045
	}
L1044:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[250])) = int32(511)
	v6307 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[251]))
	F_SimpleLruInit(m, int32(_a_F_CreateSharedMemoryAndSemaphores_145), int32(_a_F_CreateSharedMemoryAndSemaphores_146), v6307, int32(0), int32(_a_F_CreateSharedMemoryAndSemaphores_147), int32(59), int32(89), int32(5), int32(1))
	mBase = m.M
	v6315 = m.ExcPending
	if v6315 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1045:
	;
	v6232 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6228)+48)) = v6232
	*(*int32)(unsafe.Add(mBase, uint32(v6228)+40)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v6228)+32)) = v6232
	v6238 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6228)+24)) = v6238
	*(*int64)(unsafe.Add(mBase, uint32(v6228)+16)) = v6232
	*(*int32)(unsafe.Add(mBase, uint32(v6228)+8)) = v6238
	*(*int64)(unsafe.Add(mBase, uint32(v6228))) = v6232
	v6247 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v6247 <= v6238 {
		goto L1044
	} else {
		goto L1046
	}
L1046:
	;
	v6255 = int32(0)
	goto L1047
L1047:
	;
	v6269 = v6255 << (uint(int32(5)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v6228+int32(56)+v6269))) = int32(-1)
	v6273 = v6228 + v6269
	*(*int32)(unsafe.Add(mBase, uint32(v6273)+80)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6273)+72)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6273)+60)) = int64(-4294967296)
	v6281 = v6255 + int32(1)
	v6283 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v6281 < v6283 {
		v6255 = v6281
		goto L1047
	} else {
		goto L1049
	}
L1048:
	;
	goto L1044
L1049:
	;
	goto L1048
L1050:
	;
	v6316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6214)+15)))
	if v6316 == int32(0) {
		goto L1051
	} else {
		goto L1052
	}
L1051:
	;
	v6322 = F_SlruScanDirectory(m, int32(_a_F_CreateSharedMemoryAndSemaphores_145), int32(290), int32(0))
	mBase = m.M
	v6323 = m.ExcPending
	if v6323 != 0 {
		goto L1
	} else {
		goto L1054
	}
L1052:
	;
	goto L1053
L1053:
	;
	v6324 = int32(16)
	m.G0 = v6214 + v6324
	v6327 = m.G0
	v6329 = v6327 - v6324
	m.G0 = v6329
	v6333 = F_StatsShmemSize(m)
	mBase = m.M
	v6334 = m.ExcPending
	if v6334 != 0 {
		goto L1
	} else {
		goto L1055
	}
L1054:
	;
	goto L1053
L1055:
	;
	v6337 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_148), v6333, v6329+int32(15))
	mBase = m.M
	v6338 = m.ExcPending
	if v6338 != 0 {
		goto L1
	} else {
		goto L1056
	}
L1056:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[252])) = v6337
	v6341 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])))
	if v6341 == int32(0) {
		goto L1057
	} else {
		goto L1058
	}
L1057:
	;
	v6345 = v6337 + int32(_a_F_CreateSharedMemoryAndSemaphores_149)
	*(*int32)(unsafe.Add(mBase, uint32(v6337))) = v6345
	v6350 = F_dsa_create_in_place_ext(m, v6345, int32(_a_F_CreateSharedMemoryAndSemaphores_150), int32(79), int32(0))
	mBase = m.M
	v6351 = m.ExcPending
	if v6351 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1058:
	;
	goto L1059
L1059:
	;
	m.G0 = v6329 + int32(16)
	v6467 = m.G0
	v6469 = v6467 + int32(-64)
	m.G0 = v6469
	v6476 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_151), int32(8), v6467+int32(-1))
	mBase = m.M
	v6477 = m.ExcPending
	if v6477 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1060:
	;
	F_dsa_pin(m, v6350)
	mBase = m.M
	v6353 = m.ExcPending
	if v6353 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1061:
	;
	F_dsa_set_size_limit(m, v6350, int32(_a_F_CreateSharedMemoryAndSemaphores_150))
	mBase = m.M
	v6356 = m.ExcPending
	if v6356 != 0 {
		goto L1
	} else {
		goto L1062
	}
L1062:
	;
	v6359 = F_dshash_create(m, v6350, int32(_a_F_CreateSharedMemoryAndSemaphores_152), int32(0))
	mBase = m.M
	v6360 = m.ExcPending
	if v6360 != 0 {
		goto L1
	} else {
		goto L1063
	}
L1063:
	;
	v6361 = *(*int32)(unsafe.Add(mBase, uint32(v6359)+32))
	v6362 = *(*int32)(unsafe.Add(mBase, uint32(v6361)))
	goto L1064
L1064:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6337)+4)) = v6362
	F_dsa_set_size_limit(m, v6350, int32(-1))
	mBase = m.M
	v6366 = m.ExcPending
	if v6366 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1065:
	;
	F_pfree(m, v6359)
	mBase = m.M
	v6368 = m.ExcPending
	if v6368 != 0 {
		goto L1
	} else {
		goto L1066
	}
L1066:
	;
	F_dsa_detach(m, v6350)
	mBase = m.M
	v6370 = m.ExcPending
	if v6370 != 0 {
		goto L1
	} else {
		goto L1067
	}
L1067:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6337)+16)) = int64(1)
	v6376 = int32(1)
	goto L1068
L1068:
	;
	if base.Ui32(v6376-int32(1)) <= base.Ui32(int32(11)) {
		goto L1072
	} else {
		goto L1073
	}
L1069:
	;
	goto L1059
L1070:
	;
	v6445 = v6376 + int32(1)
	if v6445 != int32(33) {
		v6376 = v6445
		goto L1068
	} else {
		goto L1086
	}
L1071:
	;
	if v6420 == int32(0) {
		goto L1070
	} else {
		goto L1078
	}
L1072:
	;
	v6420 = v6376*int32(72) + int32(_a_F_CreateSharedMemoryAndSemaphores_153)
	goto L1071
L1073:
	;
	goto L1074
L1074:
	;
	if base.Ui32(int32(8)) < base.Ui32(v6376-int32(24)) {
		v6418 = int32(0)
		goto L1075
	} else {
		goto L1076
	}
L1075:
	;
	v6420 = v6418
	goto L1071
L1076:
	;
	v6406 = int32(0)
	v6408 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[253]))
	if v6408 == v6406 {
		v6418 = v6406
		goto L1075
	} else {
		goto L1077
	}
L1077:
	;
	v6416 = *(*int32)(unsafe.Add(mBase, uint32(v6408+v6376<<(uint(int32(2))%32)-int32(96))))
	v6418 = v6416
	goto L1075
L1078:
	;
	v6423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6420))))
	if v6423&int32(1) == int32(0) {
		goto L1070
	} else {
		goto L1079
	}
L1079:
	;
	if base.Ui32(v6376) <= base.Ui32(int32(12)) {
		goto L1081
	} else {
		goto L1082
	}
L1080:
	;
	v6440 = *(*int32)(unsafe.Add(mBase, uint32(v6420)+52))
	m.T0[v6440].(func(*base.Module, int32))(m, v6439)
	mBase = m.M
	v6442 = m.ExcPending
	if v6442 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1081:
	;
	v6430 = *(*int32)(unsafe.Add(mBase, uint32(v6420)+12))
	v6439 = v6337 + v6430
	goto L1080
L1082:
	;
	goto L1083
L1083:
	;
	v6435 = *(*int32)(unsafe.Add(mBase, uint32(v6420)+4))
	v6436 = F_ShmemAlloc(m, v6435)
	mBase = m.M
	v6437 = m.ExcPending
	if v6437 != 0 {
		goto L1
	} else {
		goto L1084
	}
L1084:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6337+int32(_a_F_CreateSharedMemoryAndSemaphores_154)+v6376<<(uint(int32(2))%32)))) = v6436
	v6439 = v6436
	goto L1080
L1085:
	;
	goto L1070
L1086:
	;
	goto L1069
L1087:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[254])) = v6476
	v6479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6469)+63)))
	if v6479 == int32(0) {
		goto L1088
	} else {
		goto L1089
	}
L1088:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6476))) = int64(1)
	goto L1090
L1089:
	;
	goto L1090
L1090:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6469)+28)) = int64(292057776132)
	v6491 = v6467 + int32(-52)
	v6493 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_155), int32(16), int32(128), v6491, int32(40))
	mBase = m.M
	v6494 = m.ExcPending
	if v6494 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[255])) = v6493
	*(*int64)(unsafe.Add(mBase, uint32(v6469)+28)) = int64(292057776192)
	v6503 = F_ShmemInitHash(m, int32(_a_F_CreateSharedMemoryAndSemaphores_156), int32(16), int32(128), v6491, int32(24))
	mBase = m.M
	v6504 = m.ExcPending
	if v6504 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[256])) = v6503
	m.G0 = v6469 - int32(-64)
	v6509 = int32(0)
	v6511 = m.G0
	v6513 = v6511 - int32(16)
	m.G0 = v6513
	v6516 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	v6518 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[258]))
	v6523 = v6513 + int32(15)
	v6524 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_157), int32(28), v6523)
	mBase = m.M
	v6525 = m.ExcPending
	if v6525 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1093:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259])) = v6524
	v6527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6513)+15)))
	if v6527 != 0 {
		goto L1094
	} else {
		goto L1095
	}
L1094:
	;
	v6784 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[260]))
	v6785 = *(*int32)(unsafe.Add(mBase, uint32(v6784)+8))
	if v6785 != 0 {
		goto L1123
	} else {
		goto L1124
	}
L1095:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6524)+24)) = int32(0)
	v6530 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6524)+16)) = v6530
	*(*int64)(unsafe.Add(mBase, uint32(v6524)+8)) = v6530
	*(*int64)(unsafe.Add(mBase, uint32(v6524))) = v6530
	v6537 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	v6539 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259]))
	v6541 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v6543 = v6541 + int32(38)
	*(*int32)(unsafe.Add(mBase, uint32(v6539)+8)) = v6543 * (v6516 * v6518)
	*(*int32)(unsafe.Add(mBase, uint32(v6539)+20)) = v6543 * v6537
	v6551 = F_mul_size(m, v6543, int32(164))
	mBase = m.M
	v6552 = m.ExcPending
	if v6552 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1096:
	;
	v6553 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_158), v6551, v6523)
	mBase = m.M
	v6554 = m.ExcPending
	if v6554 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1097:
	;
	v6556 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259]))
	*(*int32)(unsafe.Add(mBase, uint32(v6556)+4)) = v6553
	v6560 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v6564 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	v6566 = F_mul_size(m, v6564, int32(128))
	mBase = m.M
	v6567 = m.ExcPending
	if v6567 != 0 {
		goto L1
	} else {
		goto L1098
	}
L1098:
	;
	v6568 = F_mul_size(m, v6560+int32(38), v6566)
	mBase = m.M
	v6569 = m.ExcPending
	if v6569 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1099:
	;
	v6570 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_159), v6568, v6523)
	mBase = m.M
	v6571 = m.ExcPending
	if v6571 != 0 {
		goto L1
	} else {
		goto L1100
	}
L1100:
	;
	v6573 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259]))
	*(*int32)(unsafe.Add(mBase, uint32(v6573)+24)) = v6570
	v6578 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[258]))
	v6580 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v6583 = F_mul_size(m, v6578, v6580+int32(38))
	mBase = m.M
	v6584 = m.ExcPending
	if v6584 != 0 {
		goto L1
	} else {
		goto L1101
	}
L1101:
	;
	v6586 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	v6587 = F_mul_size(m, v6583, v6586)
	mBase = m.M
	v6588 = m.ExcPending
	if v6588 != 0 {
		goto L1
	} else {
		goto L1102
	}
L1102:
	;
	v6589 = F_mul_size(m, int32(8), v6587)
	mBase = m.M
	v6590 = m.ExcPending
	if v6590 != 0 {
		goto L1
	} else {
		goto L1103
	}
L1103:
	;
	v6591 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_160), v6589, v6523)
	mBase = m.M
	v6592 = m.ExcPending
	if v6592 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1104:
	;
	v6594 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259]))
	*(*int32)(unsafe.Add(mBase, uint32(v6594)+12)) = v6591
	v6599 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[258]))
	v6601 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v6604 = F_mul_size(m, v6599, v6601+int32(38))
	mBase = m.M
	v6605 = m.ExcPending
	if v6605 != 0 {
		goto L1
	} else {
		goto L1105
	}
L1105:
	;
	v6607 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	v6608 = F_mul_size(m, v6604, v6607)
	mBase = m.M
	v6609 = m.ExcPending
	if v6609 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1106:
	;
	v6610 = F_mul_size(m, int32(8), v6608)
	mBase = m.M
	v6611 = m.ExcPending
	if v6611 != 0 {
		goto L1
	} else {
		goto L1107
	}
L1107:
	;
	v6612 = F_ShmemInitStruct(m, int32(_a_F_CreateSharedMemoryAndSemaphores_161), v6610, v6523)
	mBase = m.M
	v6613 = m.ExcPending
	if v6613 != 0 {
		goto L1
	} else {
		goto L1108
	}
L1108:
	;
	v6615 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259]))
	*(*int32)(unsafe.Add(mBase, uint32(v6615)+16)) = v6612
	v6618 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if v6618 == int32(-38) {
		goto L1094
	} else {
		goto L1109
	}
L1109:
	;
	v6625 = int32(0)
	v6626 = v6509
	v6627 = v6509
	goto L1110
L1110:
	;
	v6639 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259]))
	v6640 = *(*int32)(unsafe.Add(mBase, uint32(v6639)+4))
	v6643 = v6640 + v6626*int32(164)
	*(*int32)(unsafe.Add(mBase, uint32(v6643))) = v6627
	v6645 = int32(_a_F_CreateSharedMemoryAndSemaphores_162)
	v6646 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	v6647 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+12)) = v6647
	v6650 = v6643 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+8)) = v6650
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+4)) = v6650
	base.MemoryFill(m, v6643+int32(24), v6647, int32(128))
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+160)) = v6647
	v6661 = v6643 + int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+156)) = v6661
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+152)) = v6661
	v6666 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	if v6647 < v6666 {
		goto L1112
	} else {
		goto L1113
	}
L1111:
	;
	goto L1094
L1112:
	;
	v6672 = v6625
	v6675 = v6647
	goto L1115
L1113:
	;
	v6746 = v6625
	goto L1114
L1114:
	;
	v6761 = v6626 + int32(1)
	v6763 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	if base.Ui32(v6761) < base.Ui32(v6763+int32(38)) {
		v6625 = v6746
		v6626 = v6761
		v6627 = v6627 + v6646
		goto L1110
	} else {
		goto L1122
	}
L1115:
	;
	v6686 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[259]))
	v6687 = *(*int32)(unsafe.Add(mBase, uint32(v6686)+24))
	v6688 = *(*int32)(unsafe.Add(mBase, uint32(v6643)))
	v6689 = int32(7)
	v6694 = v6687 + v6688<<(uint(v6689)%32) + v6675<<(uint(v6689)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v6694)+76)) = v6672
	*(*int32)(unsafe.Add(mBase, uint32(v6694)+16)) = v6626
	*(*int64)(unsafe.Add(mBase, uint32(v6694)+48)) = int64(1)
	v6699 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6694)+80)) = v6699
	*(*uint8)(unsafe.Add(mBase, uint32(v6694)+13)) = uint8(v6699)
	*(*int32)(unsafe.Add(mBase, uint32(v6694)+32)) = v6699
	*(*uint16)(unsafe.Add(mBase, uint32(v6694)+3)) = uint16(v6699)
	v6707 = *(*int32)(unsafe.Add(mBase, uint32(v6694)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v6694)+68)) = v6707 & int32(-449)
	v6712 = v6694 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v6712)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v6712))) = int64(-4294967296)
	goto L1117
L1116:
	;
	v6746 = v6737
	goto L1114
L1117:
	;
	v6717 = *(*int32)(unsafe.Add(mBase, uint32(v6643)+8))
	if v6717 == int32(0) {
		goto L1118
	} else {
		goto L1119
	}
L1118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+8)) = v6650
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+4)) = v6650
	goto L1120
L1119:
	;
	goto L1120
L1120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6694)+28)) = v6650
	v6725 = *(*int32)(unsafe.Add(mBase, uint32(v6643)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6694)+24)) = v6725
	v6728 = v6694 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v6725)+4)) = v6728
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+4)) = v6728
	v6731 = *(*int32)(unsafe.Add(mBase, uint32(v6643)+12))
	v6732 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6643)+12)) = v6731 + v6732
	v6736 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[258]))
	v6737 = v6736 + v6672
	v6739 = v6675 + v6732
	v6741 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[257]))
	if v6739 < v6741 {
		v6672 = v6737
		v6675 = v6739
		goto L1115
	} else {
		goto L1121
	}
L1121:
	;
	goto L1116
L1122:
	;
	goto L1111
L1123:
	;
	v6786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6513)+15)))
	m.T0[v6785].(func(*base.Module, int32))(m, (v6786^int32(-1))&int32(1))
	mBase = m.M
	v6792 = m.ExcPending
	if v6792 != 0 {
		goto L1
	} else {
		goto L1126
	}
L1124:
	;
	goto L1125
L1125:
	;
	m.G0 = v6513 + int32(16)
	v6796 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v6797 = m.G0
	v6799 = v6797 - int32(1120)
	m.G0 = v6799
	*(*int32)(unsafe.Add(mBase, uint32(v6799)+76)) = int32(0)
	v6804 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[261]))
	if v6804 == int32(4) {
		goto L1129
	} else {
		goto L1130
	}
L1126:
	;
	goto L1125
L1127:
	;
	v7057 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[262]))
	if v7057 != 0 {
		goto L1189
	} else {
		goto L1190
	}
L1128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7040 = m.ExcPending
	if v7040 != 0 {
		goto L1
	} else {
		goto L1185
	}
L1129:
	;
	v6808 = F_AllocateDir(m, int32(_a_F_CreateSharedMemoryAndSemaphores_163))
	mBase = m.M
	v6809 = m.ExcPending
	if v6809 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1130:
	;
	goto L1131
L1131:
	;
	v6946 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[172]))
	v6950 = v6946*int32(5) - int32(-64)
	v6953 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6954 = m.ExcPending
	if v6954 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1132:
	;
	v6811 = F_ReadDir(m, v6808, int32(_a_F_CreateSharedMemoryAndSemaphores_163))
	mBase = m.M
	v6812 = m.ExcPending
	if v6812 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	if v6811 != 0 {
		goto L1134
	} else {
		goto L1135
	}
L1134:
	;
	v6814 = v6811
	goto L1137
L1135:
	;
	goto L1136
L1136:
	;
	F_FreeDir(m, v6808)
	mBase = m.M
	v6928 = m.ExcPending
	if v6928 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1137:
	;
	v6830 = v6814 + int32(19)
	v6831 = int32(_a_F_CreateSharedMemoryAndSemaphores_164)
	goto L1141
L1138:
	;
	goto L1136
L1139:
	;
	if v6869-v6870 == int32(0) {
		goto L1152
	} else {
		goto L1153
	}
L1141:
	;
	goto L1142
L1142:
	;
	v6838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6830))))
	if v6838 != 0 {
		goto L1143
	} else {
		goto L1144
	}
L1143:
	;
	v6839 = v6830
	v6840 = v6831
	v6841 = int32(5)
	v6842 = v6838
	goto L1147
L1144:
	;
	v6865 = v6831
	v6869 = int32(0)
	goto L1145
L1145:
	;
	v6870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6865))))
	goto L1139
L1146:
	;
	v6865 = v6860
	v6869 = v6862
	goto L1145
L1147:
	;
	v6844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6840))))
	if base.B2i32(v6842 != v6844)|base.B2i32(v6844 == int32(0)) != 0 {
		v6860 = v6840
		v6862 = v6842
		goto L1146
	} else {
		goto L1149
	}
L1148:
	;
	v6860 = v6854
	v6862 = int32(0)
	goto L1146
L1149:
	;
	v6850 = v6841 - int32(1)
	if v6850 == int32(0) {
		v6860 = v6840
		v6862 = v6842
		goto L1146
	} else {
		goto L1150
	}
L1150:
	;
	v6853 = int32(1)
	v6854 = v6840 + v6853
	v6855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6839)+1)))
	if v6855 != 0 {
		v6839 = v6839 + v6853
		v6840 = v6854
		v6841 = v6850
		v6842 = v6855
		goto L1147
	} else {
		goto L1151
	}
L1151:
	;
	goto L1148
L1152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6799)+64)) = v6830
	v6882 = v6799 + int32(80)
	v6887 = F_pg_snprintf(m, v6882, int32(1036), int32(_a_F_CreateSharedMemoryAndSemaphores_165), v6799-int32(-64))
	mBase = m.M
	v6888 = m.ExcPending
	if v6888 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1153:
	;
	goto L1154
L1154:
	;
	v6909 = F_ReadDir(m, v6808, int32(_a_F_CreateSharedMemoryAndSemaphores_163))
	mBase = m.M
	v6910 = m.ExcPending
	if v6910 != 0 {
		goto L1
	} else {
		goto L1163
	}
L1155:
	;
	v6891 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v6892 = m.ExcPending
	if v6892 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1156:
	;
	if v6891 != 0 {
		goto L1157
	} else {
		goto L1158
	}
L1157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6799)+48)) = v6882
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_166), v6799+int32(48))
	mBase = m.M
	v6898 = m.ExcPending
	if v6898 != 0 {
		goto L1
	} else {
		goto L1160
	}
L1158:
	;
	goto L1159
L1159:
	;
	v6906 = F_unlink(m, v6799+int32(80))
	mBase = m.M
	if v6906 != 0 {
		goto L1128
	} else {
		goto L1162
	}
L1160:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), int32(337), int32(_a_F_CreateSharedMemoryAndSemaphores_167))
	mBase = m.M
	v6903 = m.ExcPending
	if v6903 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1161:
	;
	goto L1159
L1162:
	;
	goto L1154
L1163:
	;
	if v6909 != 0 {
		v6814 = v6909
		goto L1137
	} else {
		goto L1164
	}
L1164:
	;
	goto L1138
L1165:
	;
	goto L1131
L1166:
	;
	if v6953 != 0 {
		goto L1167
	} else {
		goto L1168
	}
L1167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6799)+16)) = v6950
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_168), v6799+int32(16))
	mBase = m.M
	v6960 = m.ExcPending
	if v6960 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1168:
	;
	goto L1169
L1169:
	;
	v6969 = v6950*int32(24) + int32(12)
	goto L1172
L1170:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), int32(198), int32(_a_F_CreateSharedMemoryAndSemaphores_169))
	mBase = m.M
	v6965 = m.ExcPending
	if v6965 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1171:
	;
	goto L1169
L1172:
	;
	v6988 = Fn13964(m, int64(32))
	mBase = m.M
	goto L1174
L1173:
	;
	v7005 = *(*int32)(unsafe.Add(mBase, uint32(v6799)+76))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[263])) = v7005
	F_on_shmem_exit(m, int32(1097), v6796)
	mBase = m.M
	v7009 = m.ExcPending
	if v7009 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1174:
	;
	v6990 = v6988 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[264])) = v6990
	if v6990 == int32(0) {
		goto L1172
	} else {
		goto L1175
	}
L1175:
	;
	v7000 = F_dsm_impl_op(m, int32(0), v6990, v6969, int32(_a_F_CreateSharedMemoryAndSemaphores_170), v6799+int32(76), int32(_a_F_CreateSharedMemoryAndSemaphores_171), int32(21))
	mBase = m.M
	v7001 = m.ExcPending
	if v7001 != 0 {
		goto L1
	} else {
		goto L1176
	}
L1176:
	;
	if v7000 == int32(0) {
		goto L1172
	} else {
		goto L1177
	}
L1177:
	;
	goto L1173
L1178:
	;
	v7012 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v7013 = m.ExcPending
	if v7013 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1179:
	;
	if v7012 != 0 {
		goto L1180
	} else {
		goto L1181
	}
L1180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6799)+4)) = v6969
	v7016 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[264]))
	*(*int32)(unsafe.Add(mBase, uint32(v6799))) = v7016
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_172), v6799)
	mBase = m.M
	v7020 = m.ExcPending
	if v7020 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1181:
	;
	goto L1182
L1182:
	;
	v7027 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[264]))
	*(*int32)(unsafe.Add(mBase, uint32(v6796)+16)) = v7027
	v7030 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[263]))
	*(*int32)(unsafe.Add(mBase, uint32(v7030)+8)) = v6950
	*(*int64)(unsafe.Add(mBase, uint32(v7030))) = int64(2588949810)
	m.G0 = v6799 + int32(1120)
	goto L1127
L1183:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), int32(223), int32(_a_F_CreateSharedMemoryAndSemaphores_169))
	mBase = m.M
	v7025 = m.ExcPending
	if v7025 != 0 {
		goto L1
	} else {
		goto L1184
	}
L1184:
	;
	goto L1182
L1185:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v7042 = m.ExcPending
	if v7042 != 0 {
		goto L1
	} else {
		goto L1186
	}
L1186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6799)+32)) = v6799 + int32(80)
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_173), v6799+int32(32))
	mBase = m.M
	v7050 = m.ExcPending
	if v7050 != 0 {
		goto L1
	} else {
		goto L1187
	}
L1187:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), int32(343), int32(_a_F_CreateSharedMemoryAndSemaphores_167))
	mBase = m.M
	v7055 = m.ExcPending
	if v7055 != 0 {
		goto L1
	} else {
		goto L1188
	}
L1188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1189:
	;
	m.T0[v7057].(func(*base.Module))(m)
	mBase = m.M
	v7059 = m.ExcPending
	if v7059 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1190:
	;
	goto L1191
L1191:
	;
	m.G0 = v19 + int32(16)
	return
L1192:
	;
	goto L1191
L1193:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_174), int32(258), int32(_a_F_CreateSharedMemoryAndSemaphores_175))
	mBase = m.M
	v7074 = m.ExcPending
	if v7074 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
