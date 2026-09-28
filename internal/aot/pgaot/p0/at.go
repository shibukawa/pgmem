package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATExecAddConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v325 int64
	_ = v325
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v576 int32
	_ = v576
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v698 int32
	_ = v698
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v783 int32
	_ = v783
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int64
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v909 int32
	_ = v909
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v979 int32
	_ = v979
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1037 int32
	_ = v1037
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1207 int32
	_ = v1207
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1263 int64
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1289 int32
	_ = v1289
	var v1306 int32
	_ = v1306
	var v1320 int32
	_ = v1320
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1459 int32
	_ = v1459
	var v1465 int32
	_ = v1465
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1586 int32
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var v1646 int32
	_ = v1646
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1684 int32
	_ = v1684
	var v1706 int32
	_ = v1706
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1778 int32
	_ = v1778
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1827 int32
	_ = v1827
	var v1830 int32
	_ = v1830
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1929 int32
	_ = v1929
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2028 int32
	_ = v2028
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2152 int32
	_ = v2152
	var v2175 int32
	_ = v2175
	var v2185 int32
	_ = v2185
	var v2193 int32
	_ = v2193
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2222 int32
	_ = v2222
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2228 int64
	_ = v2228
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2274 int32
	_ = v2274
	var v2279 int32
	_ = v2279
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2299 int32
	_ = v2299
	var v2302 int32
	_ = v2302
	var v2309 int32
	_ = v2309
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2328 int32
	_ = v2328
	var v2333 int32
	_ = v2333
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2359 int32
	_ = v2359
	var v2364 int32
	_ = v2364
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2394 int32
	_ = v2394
	var v2398 int32
	_ = v2398
	var v2407 int32
	_ = v2407
	var v2412 int32
	_ = v2412
	var v2416 int32
	_ = v2416
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2456 int32
	_ = v2456
	var v2460 int32
	_ = v2460
	var v2464 int32
	_ = v2464
	var v2469 int32
	_ = v2469
	var v2473 int32
	_ = v2473
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2513 int32
	_ = v2513
	var v2517 int32
	_ = v2517
	var v2524 int32
	_ = v2524
	var v2529 int32
	_ = v2529
	var v2533 int32
	_ = v2533
	var v2540 int32
	_ = v2540
	var v2545 int32
	_ = v2545
	var v2577 int32
	_ = v2577
	var v2580 int32
	_ = v2580
	var v2584 int32
	_ = v2584
	var v2589 int32
	_ = v2589
	var v2621 int32
	_ = v2621
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2633 int32
	_ = v2633
	var v2638 int32
	_ = v2638
	var v2644 int32
	_ = v2644
	var v2667 int32
	_ = v2667
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2685 int32
	_ = v2685
	var v2690 int32
	_ = v2690
	v9 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(1632)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
	v37 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v39 - int32(1) {
	case 0, 4:
		goto L21
	default:
		goto L20
	case 8:
		goto L22
	}
L1:
	;
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v550)+12))
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v2667+v2644<<(uint(int32(2))%32))))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2671)+4))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L27
	} else {
		goto L514
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L27
	} else {
		goto L510
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L27
	} else {
		goto L506
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2533 = m.ExcPending
	if v2533 != 0 {
		goto L27
	} else {
		goto L503
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L27
	} else {
		goto L500
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L27
	} else {
		goto L493
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L27
	} else {
		goto L490
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L27
	} else {
		goto L483
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L27
	} else {
		goto L480
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L27
	} else {
		goto L470
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L27
	} else {
		goto L467
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L27
	} else {
		goto L463
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L27
	} else {
		goto L459
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L27
	} else {
		goto L455
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L27
	} else {
		goto L451
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L27
	} else {
		goto L447
	}
L17:
	;
	m.G0 = v31 + int32(1632)
	return
L18:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v424)+56))
	v1666 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[2]))
	v1668 = F_pg_class_aclcheck(m, v1664, v1666, int64(32))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L27
	} else {
		goto L312
	}
L19:
	;
	if v536 == int32(0) {
		goto L3
	} else {
		goto L311
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L27
	} else {
		goto L308
	}
L21:
	;
	F_ATAddCheckNNConstraint(m, l0, l1, l2, l3, l4, l5, int32(0), l6, l7)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L27
	} else {
		goto L307
	}
L22:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v42 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v325 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1624)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1616)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1608)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1600)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1592)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1584)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1576)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1568)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1560)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1552)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1544)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1536)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1528)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1520)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1512)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+1504)) = v325
	v359 = int32(0)
	v360 = int32(128)
	base.MemoryFill(m, v31+int32(1376), v359, v360)
	base.MemoryFill(m, v31+int32(1248), v359, v360)
	base.MemoryFill(m, v31+int32(1120), v359, v360)
	base.MemoryFill(m, v31+int32(992), v359, v360)
	base.MemoryFill(m, v31+int32(864), v359, v360)
	base.MemoryFill(m, v31+int32(736), v359, v360)
	base.MemoryFill(m, v31+int32(608), v359, v360)
	base.MemoryFill(m, v31+int32(480), v359, v360)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+472)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+464)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+456)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+448)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+440)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+432)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+424)) = v325
	*(*int64)(unsafe.Add(mBase, uint32(v31)+416)) = v325
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	if v413 != 0 {
		goto L77
	} else {
		goto L78
	}
L24:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v45 = F_ConstraintNameIsUsed(m, int32(0), v44, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1376)) = uint8(v74)
	if v72 == v74 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	return
L28:
	;
	if v45 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+384)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+388)) = v56 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_1), v31+int32(384))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_3), int32(_a_F_ATExecAddConstraint_4))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L27
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
	v288 = F_pstrdup(m, v31+int32(1376))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L27
	} else {
		goto L75
	}
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v80 <= int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v90 = int32(0)
	v93 = v9
	goto L37
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v93<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if int32(0) < v90 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L34
L39:
	;
	v123 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(1376)+v90))) = uint8(v123)
	v127 = v90 + int32(1)
	goto L41
L40:
	;
	v127 = v90
	goto L41
L41:
	;
	v130 = v31 + int32(1376) + v127
	goto L45
L42:
	;
	v250 = F_strlen(m, v130)
	mBase = m.M
	v251 = v250 + v127
	if int32(64) <= v251 {
		goto L34
	} else {
		goto L73
	}
L43:
	;
	v247 = F_strlen(m, v236)
	mBase = m.M
	goto L42
L45:
	;
	goto L46
L46:
	;
	v137 = int32(63)
	if (v130^v117)&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v240)
	goto L43
L48:
	;
	v221 = v216
	v222 = v217
	v223 = v218
	goto L69
L49:
	;
	if v211 == int32(0) {
		v236 = v209
		v237 = v210
		goto L47
	} else {
		goto L68
	}
L50:
	;
	v209 = v117
	v210 = v130
	v211 = v137
	goto L49
L51:
	;
	goto L52
L52:
	;
	v141 = int32(0)
	if base.B2i32(v117&int32(3) == v141)|int32(0) == v141 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v177 == int32(0) {
		v236 = v174
		v237 = v175
		goto L47
	} else {
		goto L62
	}
L54:
	;
	v153 = v117
	v154 = v130
	v155 = v137
	goto L57
L55:
	;
	goto L56
L56:
	;
	v174 = v117
	v175 = v130
	v176 = v137
	v177 = int32(1)
	goto L53
L57:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v157)
	if v157 == int32(0) {
		v216 = v153
		v217 = v154
		v218 = v155
		goto L48
	} else {
		goto L59
	}
L58:
	;
	v174 = v168
	v175 = v162
	v176 = v164
	v177 = v166
	goto L53
L59:
	;
	v161 = int32(1)
	v162 = v154 + v161
	v164 = v155 - v161
	v165 = int32(0)
	v166 = base.B2i32(v164 != v165)
	v168 = v153 + v161
	if v168&int32(3) == v165 {
		v174 = v168
		v175 = v162
		v176 = v164
		v177 = v166
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v164 != 0 {
		v153 = v168
		v154 = v162
		v155 = v164
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if base.B2i32(v180 == int32(0))|base.B2i32(base.Ui32(v176) < base.Ui32(int32(4))) != 0 {
		v209 = v174
		v210 = v175
		v211 = v176
		goto L49
	} else {
		goto L63
	}
L63:
	;
	v187 = v174
	v188 = v175
	v189 = v176
	goto L64
L64:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v195 = int32(-2139062144)
	if (int32(16843008)-v192|v192)&v195 != v195 {
		v216 = v187
		v217 = v188
		v218 = v189
		goto L48
	} else {
		goto L66
	}
L65:
	;
	v209 = v203
	v210 = v201
	v211 = v205
	goto L49
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v192
	v200 = int32(4)
	v201 = v188 + v200
	v203 = v187 + v200
	v205 = v189 - v200
	if base.Ui32(int32(3)) < base.Ui32(v205) {
		v187 = v203
		v188 = v201
		v189 = v205
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v216 = v209
	v217 = v210
	v218 = v211
	goto L48
L69:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	*(*uint8)(unsafe.Add(mBase, uint32(v222))) = uint8(v225)
	if v225 == int32(0) {
		v236 = v221
		v237 = v222
		goto L47
	} else {
		goto L71
	}
L70:
	;
	v236 = v232
	v237 = v230
	goto L47
L71:
	;
	v229 = int32(1)
	v230 = v222 + v229
	v232 = v221 + v229
	v234 = v223 - v229
	if v234 != 0 {
		v221 = v232
		v222 = v230
		v223 = v234
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v255 = v93 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v255 < v256 {
		v90 = v251
		v93 = v255
		goto L37
	} else {
		goto L74
	}
L74:
	;
	goto L38
L75:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+68))
	v294 = F_ChooseConstraintName(m, v73+int32(4), v288, int32(_a_F_ATExecAddConstraint_5), v292, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L27
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v294
	goto L23
L77:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+12))
	v415 = v414
	goto L79
L78:
	;
	v415 = v9
	goto L79
L79:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l4)+100))
	if v416 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if l5 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v418 = F_table_open(m, v416, int32(6))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L27
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l4)+72))
	v422 = F_table_openrv(m, v420, int32(6))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L27
	} else {
		goto L85
	}
L84:
	;
	v424 = v418
	goto L80
L85:
	;
	v424 = v422
	goto L80
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L27
	} else {
		goto L303
	}
L87:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+119)))
	if v428 == int32(112) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+119)))
	switch v432 - int32(112) {
	case 0, 2:
		goto L91
	default:
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[3])))
	if v457 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L27
	} else {
		goto L93
	}
L93:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L27
	} else {
		goto L94
	}
L94:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v442 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_6), v31+int32(16))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L27
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_7), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L27
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L27
	} else {
		goto L299
	}
L98:
	;
	v461 = int32(1)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v424)+56))
	if base.Ui32(v462) < base.Ui32(int32(_a_F_ATExecAddConstraint_9)) {
		v471 = v461
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+118)))
	switch v473 - int32(112) {
	case 0:
		goto L111
	default:
		goto L108
	case 4:
		goto L109
	case 5:
		goto L110
	}
L101:
	;
	if v471 != 0 {
		goto L97
	} else {
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+68))
	if v466 == int32(99) {
		v471 = v461
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v469 = F_isTempToastNamespace(m, v466)
	mBase = m.M
	v471 = v469
	goto L102
L105:
	;
	goto L100
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L27
	} else {
		goto L295
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L27
	} else {
		goto L291
	}
L108:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v534 = F_transformColumnNameList(m, v526, v527, v31+int32(1504), v31+int32(1248), v31+int32(992))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L27
	} else {
		goto L125
	}
L109:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+118)))
	if v517 != int32(116) {
		goto L106
	} else {
		goto L122
	}
L110:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+118)))
	switch v497 - int32(112) {
	case 0, 5:
		goto L108
	default:
		goto L117
	}
L111:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+118)))
	if v477 == int32(112) {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L27
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L27
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_10), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L27
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_11), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L27
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L27
	} else {
		goto L118
	}
L118:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L27
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_12), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L27
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_13), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L27
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424)+24)))
	if v520 != int32(1) {
		goto L107
	} else {
		goto L123
	}
L123:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)))
	if v523 == int32(0) {
		goto L107
	} else {
		goto L124
	}
L124:
	;
	goto L108
L125:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v536 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L27
	} else {
		goto L287
	}
L127:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+85)))
	if v539 == int32(1) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v546 = int32(0)
	v548 = F_transformColumnNameList(m, v542, v543, v31+int32(416), v546, v546)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L27
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	if v548 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v551 = int32(0)
	if v534 == v551 {
		v2644 = v551
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v757 = v9
	goto L134
L134:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	if v763 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L135:
	;
	v559 = v551
	v576 = v9
	goto L136
L136:
	;
	v588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(416)+v559<<(uint(int32(1))%32)))))
	v595 = int32(0)
	goto L138
L137:
	;
	v757 = v726
	goto L134
L138:
	;
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(1504)+v595<<(uint(int32(1))%32)))))
	if v588 != v622 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v627 = int32(0)
	if v627 < v576 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	v625 = v595 + int32(1)
	if v534 != v625 {
		v595 = v625
		goto L138
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	v2644 = v559
	goto L1
L144:
	;
	v733 = v559 + int32(1)
	if v733 != v548 {
		v559 = v733
		v576 = v726
		goto L136
	} else {
		goto L152
	}
L145:
	;
	v636 = v627
	goto L148
L146:
	;
	goto L147
L147:
	;
	v698 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(416)+v576<<(uint(v698)%32)))) = uint16(v588)
	v726 = v576 + v698
	goto L144
L148:
	;
	v663 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(416)+v636<<(uint(int32(1))%32)))))
	if v663 == v588 {
		v726 = v576
		goto L144
	} else {
		goto L150
	}
L149:
	;
	goto L147
L150:
	;
	v666 = v636 + int32(1)
	if v666 != v576 {
		v636 = v666
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	goto L137
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L27
	} else {
		goto L284
	}
L154:
	;
	v766 = F_RelationGetIndexList(m, v424)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L27
	} else {
		goto L159
	}
L155:
	;
	goto L156
L156:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v424)+56))
	v1026 = F_transformColumnNameList(m, v1019, v763, v31+int32(1568), v31+int32(1376), v31+int32(1120))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L27
	} else {
		goto L199
	}
L157:
	;
	F_list_free(m, v766)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L27
	} else {
		goto L177
	}
L158:
	;
	F_list_free(m, v766)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L27
	} else {
		goto L176
	}
L159:
	;
	if v766 == int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	if v770 <= int32(0) {
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v783 = int32(0)
	goto L162
L162:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v766)+12))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v803+v783<<(uint(int32(2))%32))))
	v809 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(v807))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L27
	} else {
		goto L164
	}
L163:
	;
	goto L158
L164:
	;
	if v809 == int32(0) {
		goto L153
	} else {
		goto L165
	}
L165:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v809)+16))
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813)+22)))
	v815 = v813 + v814
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+14)))
	if v816 != int32(1) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_ReleaseCatCache(m, v809)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L27
	} else {
		goto L174
	}
L167:
	;
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+18)))
	if v819 != int32(1) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+16)))
	if v822 != 0 {
		goto L157
	} else {
		goto L169
	}
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L27
	} else {
		goto L170
	}
L170:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L27
	} else {
		goto L171
	}
L171:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v830 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_14), v31+int32(272))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L27
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_15), int32(_a_F_ATExecAddConstraint_16))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L27
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	v847 = v783 + int32(1)
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	if v847 < v848 {
		v783 = v847
		goto L162
	} else {
		goto L175
	}
L175:
	;
	goto L163
L176:
	;
	goto L2
L177:
	;
	if v807 == int32(0) {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	v884 = int32(0)
	v887 = F_SysCacheGetAttrNotNull(m, int32(34), v809, int32(18))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L27
	} else {
		goto L179
	}
L179:
	;
	v889 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v889
	v891 = int32(*(*int16)(unsafe.Add(mBase, uint32(v815)+10)))
	if v889 < v891 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v909 = v884
	goto L183
L181:
	;
	v979 = v884
	goto L182
L182:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+15)))
	F_ReleaseCatCache(m, v809)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L27
	} else {
		goto L192
	}
L183:
	;
	v928 = v909 << (uint(int32(1)) % 32)
	v933 = int32(*(*int16)(unsafe.Add(mBase, uint32(v928+(v815+int32(48))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v928+(v31+int32(1568))))) = uint16(v933)
	v936 = v909 << (uint(int32(2)) % 32)
	v940 = F_attnumTypeId(m, v424, v933)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L27
	} else {
		goto L185
	}
L184:
	;
	v979 = v966
	goto L182
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v936+(v31+int32(1376))))) = v940
	v946 = F_attnumCollationId(m, v424, v933)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L27
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(1120)+v936))) = v946
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v936+(base.I32_wrap_i64(v887)+int32(24)))))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(864)+v936))) = v953
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v956 = F_attnumAttName(m, v424, v933)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L27
	} else {
		goto L187
	}
L187:
	;
	v958 = F_pstrdup(m, v956)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L27
	} else {
		goto L188
	}
L188:
	;
	v960 = F_makeString(m, v958)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L27
	} else {
		goto L189
	}
L189:
	;
	v962 = F_lappend(m, v955, v960)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L27
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v962
	v966 = v909 + int32(1)
	v967 = int32(*(*int16)(unsafe.Add(mBase, uint32(v815)+10)))
	if v966 < v967 {
		v909 = v966
		goto L183
	} else {
		goto L191
	}
L191:
	;
	goto L184
L192:
	;
	if v997 != int32(1) {
		v1646 = v979
		v1659 = v807
		goto L18
	} else {
		goto L193
	}
L193:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v1002 != 0 {
		goto L19
	} else {
		goto L194
	}
L194:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L27
	} else {
		goto L195
	}
L195:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L27
	} else {
		goto L196
	}
L196:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_18), int32(0))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L27
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_19), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L27
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	if v536 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L27
	} else {
		goto L280
	}
L201:
	;
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+85)))
	if v1028 == int32(0) {
		goto L200
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if v1026 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L203
L205:
	;
	v1037 = int32(0)
	goto L208
L206:
	;
	goto L207
L207:
	;
	v1180 = F_RelationGetIndexList(m, v424)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L27
	} else {
		goto L227
	}
L208:
	;
	v1061 = v1037 + int32(1)
	if base.Ui32(v1026) <= base.Ui32(v1061) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	goto L207
L210:
	;
	if v1061 != v1026 {
		v1037 = v1061
		goto L208
	} else {
		goto L222
	}
L211:
	;
	v1068 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(1568)+v1037<<(uint(int32(1))%32)))))
	v1075 = v1061
	goto L212
L212:
	;
	v1102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(1568)+v1075<<(uint(int32(1))%32)))))
	if v1102 != v1068 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L27
	} else {
		goto L218
	}
L214:
	;
	v1105 = v1075 + int32(1)
	if v1026 != v1105 {
		v1075 = v1105
		goto L212
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	goto L213
L217:
	;
	goto L210
L218:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L27
	} else {
		goto L219
	}
L219:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_20), int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L27
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_21), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L27
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	goto L209
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L27
	} else {
		goto L276
	}
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L27
	} else {
		goto L273
	}
L225:
	;
	v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+15)))
	F_ReleaseCatCache(m, v1233)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L27
	} else {
		goto L270
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L27
	} else {
		goto L266
	}
L227:
	;
	if v1180 == int32(0) {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1180)+4))
	if v1184 <= int32(0) {
		goto L226
	} else {
		goto L229
	}
L229:
	;
	v1187 = int32(0)
	v1190 = int32(1)
	v1193 = (v1026 - v1190) << (uint(v1190) % 32)
	v1207 = v1187
	v1224 = v9
	goto L230
L230:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1180)+12))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1227+v1207<<(uint(int32(2))%32))))
	v1233 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(v1231))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L27
	} else {
		goto L232
	}
L231:
	;
	if v1389 != 0 {
		goto L223
	} else {
		goto L265
	}
L232:
	;
	if v1233 == int32(0) {
		goto L224
	} else {
		goto L233
	}
L233:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+16))
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237)+22)))
	v1239 = v1237 + v1238
	v1240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1239)+10)))
	if v1026 != v1240 {
		v1389 = v1224
		goto L234
	} else {
		goto L235
	}
L234:
	;
	F_ReleaseCatCache(m, v1233)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L27
	} else {
		goto L263
	}
L235:
	;
	if v536 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+18)))
	if v1246 != int32(1) {
		v1389 = v1224
		goto L234
	} else {
		goto L242
	}
L237:
	;
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+15)))
	if v1242 != 0 {
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+12)))
	if v1243 != int32(1) {
		v1389 = v1224
		goto L234
	} else {
		goto L241
	}
L240:
	;
	v1389 = v1224
	goto L234
L241:
	;
	goto L236
L242:
	;
	v1251 = F_heap_attisnull(m, v1233, int32(21), int32(0))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L27
	} else {
		goto L243
	}
L243:
	;
	if v1251 == int32(0) {
		v1389 = v1224
		goto L234
	} else {
		goto L244
	}
L244:
	;
	v1257 = F_heap_attisnull(m, v1233, int32(20), int32(0))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L27
	} else {
		goto L245
	}
L245:
	;
	if v1257 == int32(0) {
		v1389 = v1224
		goto L234
	} else {
		goto L246
	}
L246:
	;
	v1263 = F_SysCacheGetAttrNotNull(m, int32(34), v1233, int32(18))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L27
	} else {
		goto L247
	}
L247:
	;
	if v1026 == int32(0) {
		v1389 = v1224
		goto L234
	} else {
		goto L248
	}
L248:
	;
	v1268 = v1239 + int32(48)
	v1289 = int32(0)
	goto L249
L249:
	;
	v1306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(1568)+v1289<<(uint(int32(1))%32)))))
	v1320 = int32(0)
	goto L251
L250:
	;
	if base.B2i32(v1026 != v1187)&v536 != 0 {
		goto L258
	} else {
		goto L259
	}
L251:
	;
	v1339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1268+v1320<<(uint(int32(1))%32)))))
	if v1339 != v1306 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1346 = int32(2)
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1263)+int32(24)+v1320<<(uint(v1346)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(864)+v1289<<(uint(v1346)%32)))) = v1352
	v1355 = v1289 + int32(1)
	if v1355 != v1026 {
		v1289 = v1355
		goto L249
	} else {
		goto L257
	}
L253:
	;
	v1342 = v1320 + int32(1)
	if v1342 != v1026 {
		v1320 = v1342
		goto L251
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	goto L252
L256:
	;
	v1389 = v1224
	goto L234
L257:
	;
	goto L250
L258:
	;
	v1357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1193+(v31+int32(1568))))))
	v1359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1193+v1268))))
	if v1357 != v1359 {
		v1389 = v1224
		goto L234
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1239)+16)))
	if v1362 != 0 {
		goto L225
	} else {
		goto L262
	}
L261:
	;
	goto L260
L262:
	;
	v1389 = int32(1)
	goto L234
L263:
	;
	v1394 = v1207 + int32(1)
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1180)+4))
	if v1394 < v1395 {
		v1207 = v1394
		v1224 = v1389
		goto L230
	} else {
		goto L264
	}
L264:
	;
	goto L231
L265:
	;
	goto L226
L266:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L27
	} else {
		goto L267
	}
L267:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+288)) = v1432 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_23), v31+int32(288))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L27
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_24), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L27
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_list_free(m, v1180)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L27
	} else {
		goto L271
	}
L271:
	;
	if v536|base.B2i32(v1446&int32(1) == int32(0)) != 0 {
		v1646 = v1026
		v1659 = v1231
		goto L18
	} else {
		goto L272
	}
L272:
	;
	goto L3
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v1231
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_25), v31+int32(304))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L27
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_26), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L27
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L27
	} else {
		goto L277
	}
L277:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+320)) = v1478 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_27), v31+int32(320))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L27
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_28), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L27
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L27
	} else {
		goto L281
	}
L281:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_29), int32(0))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L27
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_30), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L27
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v807
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_25), v31+int32(48))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L27
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_31), int32(_a_F_ATExecAddConstraint_16))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L27
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L27
	} else {
		goto L288
	}
L288:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_18), int32(0))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L27
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_32), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L27
	} else {
		goto L290
	}
L290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L291:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L27
	} else {
		goto L292
	}
L292:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_33), int32(0))
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L27
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_34), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L27
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L27
	} else {
		goto L296
	}
L296:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_35), int32(0))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L27
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_36), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L27
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L27
	} else {
		goto L300
	}
L300:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+352)) = v1578 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_37), v31+int32(352))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L27
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_38), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L27
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L27
	} else {
		goto L304
	}
L304:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	v1601 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+372)) = v1600 + v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+368)) = v1599 + v1601
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_39), v31+int32(368))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L27
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_40), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L27
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	goto L17
L308:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v1624
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_41), v31)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L27
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_42), int32(_a_F_ATExecAddConstraint_4))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L27
	} else {
		goto L310
	}
L310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L311:
	;
	v1646 = v979
	v1659 = v807
	goto L18
L312:
	;
	v1670 = int32(0)
	if base.B2i32(v1668 == v1670)|base.B2i32(v1646 <= v1670) == v1670 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1684 = int32(0)
	goto L316
L314:
	;
	goto L315
L315:
	;
	if v534 != 0 {
		goto L331
	} else {
		goto L332
	}
L316:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v424)+56))
	v1712 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31+int32(1568)+v1684<<(uint(int32(1))%32)))))
	v1714 = F_pg_attribute_aclcheck(m, v1706, v1712, v1666, int64(32))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L27
	} else {
		goto L318
	}
L317:
	;
	goto L315
L318:
	;
	if v1714 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	v1717 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1716)+119)))
	switch v1717 - int32(73) {
	case 0, 32:
		v1727 = int32(20)
		goto L323
	default:
		goto L324
	case 10:
		goto L328
	case 29:
		goto L325
	case 36:
		goto L326
	case 45:
		goto L327
	}
L320:
	;
	goto L321
L321:
	;
	v1736 = v1684 + int32(1)
	if v1736 != v1646 {
		v1684 = v1736
		goto L316
	} else {
		goto L330
	}
L322:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	F_aclcheck_error(m, v1714, v1729, v1730+int32(4))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L27
	} else {
		goto L329
	}
L323:
	;
	v1729 = v1727
	goto L322
L324:
	;
	v1727 = int32(42)
	goto L323
L325:
	;
	v1729 = int32(18)
	goto L322
L326:
	;
	v1729 = int32(23)
	goto L322
L327:
	;
	v1729 = int32(52)
	goto L322
L328:
	;
	v1729 = int32(38)
	goto L322
L329:
	;
	goto L321
L330:
	;
	goto L317
L331:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1766)))
	v1778 = int32(0)
	goto L334
L332:
	;
	goto L333
L333:
	;
	v1877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v1877 != int32(1) {
		goto L351
	} else {
		goto L352
	}
L334:
	;
	v1805 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31+int32(1504)+v1778<<(uint(int32(1))%32)))))
	v1809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766+v1767<<(uint(int32(3))%32)+v1805*int32(100))+18)))
	if v1809 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	goto L333
L336:
	;
	v1810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+87)))
	v1812 = v1810 - int32(99)
	if int32(1)<<(uint(v1812)%32)&int32(2051) != 0 {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	goto L338
L338:
	;
	v1847 = v1778 + int32(1)
	if v1847 != v534 {
		v1778 = v1847
		goto L334
	} else {
		goto L350
	}
L339:
	;
	v1820 = base.B2i32(base.Ui32(v1812) <= base.Ui32(int32(11)))
	goto L341
L340:
	;
	v1820 = int32(0)
	goto L341
L341:
	;
	if v1820 != 0 {
		goto L16
	} else {
		goto L342
	}
L342:
	;
	v1821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+88)))
	switch v1821 - int32(100) {
	case 0, 10:
		goto L344
	default:
		goto L343
	}
L343:
	;
	if v1809 == int32(118) {
		goto L15
	} else {
		goto L349
	}
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L27
	} else {
		goto L345
	}
L345:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L27
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = int32(_a_F_ATExecAddConstraint_43)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_44), v31+int32(240))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L27
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_45), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L27
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	goto L338
L350:
	;
	goto L335
L351:
	;
	if v1646 != v534 {
		goto L12
	} else {
		goto L359
	}
L352:
	;
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+87)))
	v1882 = v1880 - int32(99)
	if int32(1)<<(uint(v1882)%32)&int32(_a_F_ATExecAddConstraint_46) != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1890 = base.B2i32(base.Ui32(v1882) <= base.Ui32(int32(15)))
	goto L355
L354:
	;
	v1890 = int32(0)
	goto L355
L355:
	;
	if v1890 != 0 {
		goto L14
	} else {
		goto L356
	}
L356:
	;
	v1891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+88)))
	v1893 = v1891 - int32(99)
	if base.Ui32(int32(15)) < base.Ui32(v1893) {
		goto L351
	} else {
		goto L357
	}
L357:
	;
	if int32(1)<<(uint(v1893)%32)&int32(_a_F_ATExecAddConstraint_46) != 0 {
		goto L13
	} else {
		goto L358
	}
L358:
	;
	goto L351
L359:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	v1904 = base.B2i32(v1902 != int32(0))
	if v534 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1918 = int32(0)
	v1921 = v415
	v1929 = v1904
	goto L363
L361:
	;
	v2175 = v1904
	goto L362
L362:
	;
	if v536 != 0 {
		goto L439
	} else {
		goto L440
	}
L363:
	;
	v1937 = v1918 << (uint(int32(2)) % 32)
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1937+(v31+int32(992)))))
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(1120)+v1937)))
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(1248)+v1937)))
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(1376)+v1937)))
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(864)+v1937)))
	v1960 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(v1958))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L27
	} else {
		goto L365
	}
L364:
	;
	v2175 = v2137
	goto L362
L365:
	;
	if v1960 == int32(0) {
		goto L11
	} else {
		goto L366
	}
L366:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1960)+16))
	v1965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1964)+22)))
	v1966 = v1964 + v1965
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1966)+84))
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1966)+80))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1966)+4))
	F_ReleaseCatCache(m, v1960)
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L27
	} else {
		goto L367
	}
L367:
	;
	v1975 = v536 & base.B2i32(v1918 == v534-int32(1))
	if v1975 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1976 = int32(7)
	goto L370
L369:
	;
	v1976 = int32(3)
	goto L370
L370:
	;
	v1978 = F_IndexAmTranslateCompareType(m, v1976, v1969, v1968, int32(1))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L27
	} else {
		goto L371
	}
L371:
	;
	if v1978 == int32(0) {
		goto L10
	} else {
		goto L372
	}
L372:
	;
	v1982 = base.I32_extend16_s(v1978)
	v1983 = F_get_opfamily_member(m, v1968, v1967, v1967, v1982)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L27
	} else {
		goto L373
	}
L373:
	;
	if v1983 == int32(0) {
		goto L9
	} else {
		goto L374
	}
L374:
	;
	v1987 = F_getBaseType(m, v1949)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L27
	} else {
		goto L377
	}
L375:
	;
	v2018 = int32(0)
	if base.B2i32(v2016 == v2018)|base.B2i32(v2017 == v2018) != 0 {
		goto L8
	} else {
		goto L394
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+396)) = v1949
	*(*int32)(unsafe.Add(mBase, uint32(v31)+392)) = v1953
	*(*int32)(unsafe.Add(mBase, uint32(v31)+412)) = v1967
	*(*int32)(unsafe.Add(mBase, uint32(v31)+408)) = v1967
	v2008 = F_can_coerce_type(m, int32(2), v31+int32(392), v31+int32(408), int32(0))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L27
	} else {
		goto L384
	}
L377:
	;
	v1989 = F_get_opfamily_member(m, v1968, v1967, v1987, v1982)
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L27
	} else {
		goto L378
	}
L378:
	;
	if v1989 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1996 = int32(0)
	goto L376
L380:
	;
	goto L381
L381:
	;
	v1994 = F_get_opfamily_member(m, v1968, v1987, v1987, v1982)
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L27
	} else {
		goto L382
	}
L382:
	;
	if v1994 != 0 {
		v2014 = v1987
		v2016 = v1989
		v2017 = v1994
		goto L375
	} else {
		goto L383
	}
L383:
	;
	v1996 = v1987
	goto L376
L384:
	;
	if v2008 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v2010 = v1983
	goto L387
L386:
	;
	v2010 = v1989
	goto L387
L387:
	;
	if v2008 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2012 = v1983
	goto L390
L389:
	;
	v2012 = int32(0)
	goto L390
L390:
	;
	if v2008 != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2013 = v1967
	goto L393
L392:
	;
	v2013 = v1996
	goto L393
L393:
	;
	v2014 = v2013
	v2016 = v2010
	v2017 = v2012
	goto L375
L394:
	;
	v2023 = int32(0)
	if base.B2i32(v1945 != v2023) != base.B2i32(v1941 != v2023) {
		goto L7
	} else {
		goto L395
	}
L395:
	;
	v2028 = int32(0)
	if base.B2i32(v1945 == v2028)|base.B2i32(v1941 == v2028) != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2041 = int32(0)
	if v1929&int32(1) == v2041 {
		goto L403
	} else {
		goto L404
	}
L397:
	;
	v2033 = F_get_collation_isdeterministic(m, v1945)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L27
	} else {
		goto L398
	}
L398:
	;
	v2035 = F_get_collation_isdeterministic(m, v1941)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L27
	} else {
		goto L399
	}
L399:
	;
	if v2033&v2035 != 0 {
		goto L396
	} else {
		goto L400
	}
L400:
	;
	if v1941 != v1945 {
		goto L6
	} else {
		goto L401
	}
L401:
	;
	goto L396
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(608)+v1937))) = v1983
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(736)+v1937))) = v2016
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(480)+v1937))) = v2017
	v2152 = v1918 + int32(1)
	if v2152 != v534 {
		v1918 = v2152
		v1921 = v2135
		v1929 = v2137
		goto L363
	} else {
		goto L438
	}
L403:
	;
	v2135 = v1921
	v2137 = v2041
	goto L402
L404:
	;
	goto L405
L405:
	;
	v2045 = v1921 + int32(4)
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+12))
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+4))
	if base.Ui32(v2045) < base.Ui32(v2048+v2049<<(uint(int32(2))%32)) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v2054 = v2045
	goto L408
L407:
	;
	v2054 = int32(0)
	goto L408
L408:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v1921)))
	if v2016 != v2055 {
		v2135 = v2054
		v2137 = v2041
		goto L402
	} else {
		goto L409
	}
L409:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v2057)))
	v2067 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31+int32(1504)+v1918<<(uint(int32(1))%32)))))
	v2072 = v2057 + v2058<<(uint(int32(3))%32) + v2067*int32(100) - int32(72)
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+68))
	if v2073 == v2014 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	if v2014 == v1949 {
		goto L417
	} else {
		goto L418
	}
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+392)) = int32(0)
	v2085 = int32(2)
	goto L410
L412:
	;
	goto L413
L413:
	;
	v2081 = F_find_coercion_pathway(m, v2014, v2073, int32(0), v31+int32(392))
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L27
	} else {
		goto L414
	}
L414:
	;
	if v2081 == int32(0) {
		goto L5
	} else {
		goto L415
	}
L415:
	;
	v2085 = v2081
	goto L410
L416:
	;
	if v2085 != v2097 {
		v2135 = v2054
		v2137 = v2041
		goto L402
	} else {
		goto L422
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+408)) = int32(0)
	v2097 = int32(2)
	goto L416
L418:
	;
	goto L419
L419:
	;
	v2093 = F_find_coercion_pathway(m, v2014, v1949, int32(0), v31+int32(408))
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L27
	} else {
		goto L420
	}
L420:
	;
	if v2093 == int32(0) {
		goto L4
	} else {
		goto L421
	}
L421:
	;
	v2097 = v2093
	goto L416
L422:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v31)+408))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v31)+392))
	if v2099 != v2100 {
		v2135 = v2054
		v2137 = v2041
		goto L402
	} else {
		goto L423
	}
L423:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+96))
	if v2014 <= int32(3830) {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	if v2102 == v1941 {
		v2135 = v2054
		v2137 = int32(1)
		goto L402
	} else {
		goto L434
	}
L425:
	;
	if v2073 != v1949 {
		v2135 = v2054
		v2137 = v2041
		goto L402
	} else {
		goto L433
	}
L426:
	;
	switch v2014 - int32(2277) {
	case 0, 6:
		goto L425
	case 1, 2, 3, 4, 5:
		goto L424
	default:
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	if base.B2i32(base.Ui32(v2014-int32(_a_F_ATExecAddConstraint_47)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v2014-int32(_a_F_ATExecAddConstraint_48)) < base.Ui32(int32(2))) != 0 {
		goto L425
	} else {
		goto L431
	}
L429:
	;
	if base.B2i32(v2014 == int32(2776))|base.B2i32(v2014 == int32(3500)) != 0 {
		goto L425
	} else {
		goto L430
	}
L430:
	;
	goto L424
L431:
	;
	if v2014 != int32(3831) {
		goto L424
	} else {
		goto L432
	}
L432:
	;
	goto L425
L433:
	;
	goto L424
L434:
	;
	v2127 = F_get_collation_isdeterministic(m, v2102)
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		goto L27
	} else {
		goto L435
	}
L435:
	;
	if v2127 == int32(0) {
		v2135 = v2054
		v2137 = int32(0)
		goto L402
	} else {
		goto L436
	}
L436:
	;
	v2131 = F_get_collation_isdeterministic(m, v1941)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L27
	} else {
		goto L437
	}
L437:
	;
	v2135 = v2054
	v2137 = v2131
	goto L402
L438:
	;
	goto L364
L439:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v534<<(uint(int32(2))%32)+v31)+860))
	F_FindFKPeriodOpers(m, v2185, v31+int32(392), v31+int32(408), v31+int32(404))
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L27
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v2198 = int32(0)
	v2200 = v31 + int32(1568)
	v2202 = v31 + int32(1504)
	v2204 = v31 + int32(736)
	v2206 = v31 + int32(608)
	v2208 = v31 + int32(480)
	v2210 = v31 + int32(416)
	F_addFkConstraint(m, v31+int32(392), int32(2), v2197, l4, l3, v424, v1659, v2198, v534, v2200, v2202, v2204, v2206, v2208, v757, v2210, v2198, v536)
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L27
	} else {
		goto L443
	}
L442:
	;
	goto L441
L443:
	;
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v31)+396))
	v2215 = int32(0)
	F_addFkRecurseReferenced(m, l4, l3, v424, v1659, v2214, v534, v2200, v2202, v2204, v2206, v2208, v757, v2210, v2215, v2215, v536)
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L27
	} else {
		goto L444
	}
L444:
	;
	v2219 = int32(0)
	F_addFkRecurseReferencing(m, l1, l4, l3, v424, v1659, v2214, v534, v2200, v2202, v2204, v2206, v2208, v757, v2210, v2175, l7, v2219, v2219, v536)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L27
	} else {
		goto L445
	}
L445:
	;
	F_relation_close(m, v424, int32(0))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L27
	} else {
		goto L446
	}
L446:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v31)+400))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2226
	v2228 = *(*int64)(unsafe.Add(mBase, uint32(v31)+392))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v2228
	goto L17
L447:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L27
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = int32(_a_F_ATExecAddConstraint_49)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_44), v31+int32(256))
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L27
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_50), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L27
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L27
	} else {
		goto L452
	}
L452:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_51), int32(0))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L27
	} else {
		goto L453
	}
L453:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_52), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L27
	} else {
		goto L454
	}
L454:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L455:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L27
	} else {
		goto L456
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = int32(_a_F_ATExecAddConstraint_49)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_53), v31+int32(224))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L27
	} else {
		goto L457
	}
L457:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_54), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L27
	} else {
		goto L458
	}
L458:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L459:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L27
	} else {
		goto L460
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+208)) = int32(_a_F_ATExecAddConstraint_43)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_53), v31+int32(208))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L27
	} else {
		goto L461
	}
L461:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_55), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2333 = m.ExcPending
	if v2333 != 0 {
		goto L27
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
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L27
	} else {
		goto L464
	}
L464:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_56), int32(0))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L27
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_57), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L27
	} else {
		goto L466
	}
L466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v1958
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_58), v31-int32(-64))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L27
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_59), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L27
	} else {
		goto L469
	}
L469:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L470:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L27
	} else {
		goto L471
	}
L471:
	;
	if v1975 != 0 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2374 = int32(_a_F_ATExecAddConstraint_60)
	goto L474
L473:
	;
	v2374 = int32(_a_F_ATExecAddConstraint_61)
	goto L474
L474:
	;
	F_errmsg(m, v2374, int32(0))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L27
	} else {
		goto L475
	}
L475:
	;
	v2378 = F_get_opfamily_name(m, v1968)
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L27
	} else {
		goto L476
	}
L476:
	;
	v2380 = F_get_am_name(m, v1969)
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L27
	} else {
		goto L477
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = v2380
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = v2378
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1976
	v2388 = F_errdetail(m, int32(_a_F_ATExecAddConstraint_62), v31+int32(80))
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L27
	} else {
		goto L478
	}
L478:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_63), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L27
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+108)) = v1968
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v1967
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = v1967
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v1982
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_64), v31+int32(96))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L27
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_65), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L27
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L27
	} else {
		goto L484
	}
L484:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+192)) = v2420
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_66), v31+int32(192))
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L27
	} else {
		goto L485
	}
L485:
	;
	v2428 = v1918 << (uint(int32(2)) % 32)
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v2429)+12))
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v2428+v2430)))
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v2432)+4))
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2434)+12))
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2435+v2428)))
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2437)+4))
	v2439 = F_format_type_be(m, v1949)
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L27
	} else {
		goto L486
	}
L486:
	;
	v2441 = F_format_type_be(m, v1953)
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L27
	} else {
		goto L487
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+188)) = v2441
	*(*int32)(unsafe.Add(mBase, uint32(v31)+184)) = v2439
	*(*int32)(unsafe.Add(mBase, uint32(v31)+180)) = v2438
	*(*int32)(unsafe.Add(mBase, uint32(v31)+176)) = v2433
	v2450 = F_errdetail(m, int32(_a_F_ATExecAddConstraint_67), v31+int32(176))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L27
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_68), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L27
	} else {
		goto L489
	}
L489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L490:
	;
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_69), int32(0))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L27
	} else {
		goto L491
	}
L491:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_70), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L27
	} else {
		goto L492
	}
L492:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L493:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v2476 = m.ExcPending
	if v2476 != 0 {
		goto L27
	} else {
		goto L494
	}
L494:
	;
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+160)) = v2477
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_66), v31+int32(160))
	mBase = m.M
	v2483 = m.ExcPending
	if v2483 != 0 {
		goto L27
	} else {
		goto L495
	}
L495:
	;
	v2485 = v1918 << (uint(int32(2)) % 32)
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v2486)+12))
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v2485+v2487)))
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v2489)+4))
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v2491)+12))
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v2492+v2485)))
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v2494)+4))
	v2496 = F_get_collation_name(m, v1941)
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L27
	} else {
		goto L496
	}
L496:
	;
	v2498 = F_get_collation_name(m, v1945)
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L27
	} else {
		goto L497
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+156)) = v2498
	*(*int32)(unsafe.Add(mBase, uint32(v31)+152)) = v2496
	*(*int32)(unsafe.Add(mBase, uint32(v31)+148)) = v2495
	*(*int32)(unsafe.Add(mBase, uint32(v31)+144)) = v2490
	v2507 = F_errdetail(m, int32(_a_F_ATExecAddConstraint_71), v31+int32(144))
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L27
	} else {
		goto L498
	}
L498:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_72), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L27
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+132)) = v2014
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v2073
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_73), v31+int32(128))
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L27
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_74), int32(_a_F_ATExecAddConstraint_75))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L27
	} else {
		goto L502
	}
L502:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = v2014
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v1949
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_73), v31+int32(112))
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L27
	} else {
		goto L504
	}
L504:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_74), int32(_a_F_ATExecAddConstraint_75))
	mBase = m.M
	v2545 = m.ExcPending
	if v2545 != 0 {
		goto L27
	} else {
		goto L505
	}
L505:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L506:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L27
	} else {
		goto L507
	}
L507:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_76), int32(0))
	mBase = m.M
	v2584 = m.ExcPending
	if v2584 != 0 {
		goto L27
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_77), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L27
	} else {
		goto L509
	}
L509:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L510:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L27
	} else {
		goto L511
	}
L511:
	;
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v424)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v2625 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_78), v31+int32(32))
	mBase = m.M
	v2633 = m.ExcPending
	if v2633 != 0 {
		goto L27
	} else {
		goto L512
	}
L512:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_79), int32(_a_F_ATExecAddConstraint_16))
	mBase = m.M
	v2638 = m.ExcPending
	if v2638 != 0 {
		goto L27
	} else {
		goto L513
	}
L513:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L514:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_80))
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L27
	} else {
		goto L515
	}
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+336)) = v2672
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_81), v31+int32(336))
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L27
	} else {
		goto L516
	}
L516:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_82), int32(_a_F_ATExecAddConstraint_83))
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L27
	} else {
		goto L517
	}
L517:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecSetIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+119)))
	if base.B2i32(l5 == v8)&base.B2i32(v21 == int32(112)) == v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v178 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L18
	} else {
		goto L50
	}
L2:
	;
	if l6 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L45
	}
L5:
	;
	v67 = int32(0)
	v70 = v8
	goto L24
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+131)))
	if v29&int32(1) != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v171 = int32(0)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v36 <= v35 {
		v171 = v35
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v39 = int32(0)
	if v39 < v36 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = v36
	goto L17
L16:
	;
	v42 = v39
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	goto L5
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_ATExecSetIdentity_0), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_ATExecSetIdentity_1), int32(_a_F_ATExecSetIdentity_2), int32(_a_F_ATExecSetIdentity_3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L18
	} else {
		goto L42
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v43+v67<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v79 = int32(_a_F_ATExecSetIdentity_4)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATExecSetIdentity[0])))
	if base.B2i32(v82 == int32(0))|base.B2i32(v82 != v85) != 0 {
		v103 = v82
		v104 = v85
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L38
	}
L26:
	;
	if v103-v104 != 0 {
		goto L23
	} else {
		goto L33
	}
L27:
	;
	goto L26
L28:
	;
	v88 = v78
	v89 = v79
	goto L29
L29:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v93 == int32(0) {
		v103 = v93
		v104 = v92
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v103 = v93
	v104 = v92
	goto L27
L31:
	;
	v96 = int32(1)
	if v93 == v92 {
		v88 = v88 + v96
		v89 = v89 + v96
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if v70 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v109 = v67 + int32(1)
	if v109 == v42 {
		v171 = v77
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L25
L37:
	;
	v67 = v109
	v70 = v77
	goto L24
L38:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_ATExecSetIdentity_5), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_ATExecSetIdentity_1), int32(_a_F_ATExecSetIdentity_6), int32(_a_F_ATExecSetIdentity_3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v131
	F_errmsg_internal(m, int32(_a_F_ATExecSetIdentity_7), v14+int32(-16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_ATExecSetIdentity_1), int32(_a_F_ATExecSetIdentity_8), int32(_a_F_ATExecSetIdentity_3))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_ATExecSetIdentity_9), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	F_errhint(m, int32(_a_F_ATExecSetIdentity_10), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ATExecSetIdentity_1), int32(_a_F_ATExecSetIdentity_11), int32(_a_F_ATExecSetIdentity_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v181 = F_SearchSysCacheCopyAttName(m, v180, l2)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	if v181 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L18
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+22)))
	v207 = v205 + v206
	v208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v207)+74)))
	if int32(0) < v208 {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v192 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecSetIdentity_12), v16)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_ATExecSetIdentity_1), int32(_a_F_ATExecSetIdentity_13), int32(_a_F_ATExecSetIdentity_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L18
	} else {
		goto L91
	}
L60:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+89)))
	if v211 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L18
	} else {
		goto L87
	}
L63:
	;
	if v171 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	F_pfree(m, v181)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L18
	} else {
		goto L74
	}
L65:
	;
	v214 = F_defGetInt32(m, v171)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L18
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetIdentity[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v236
	v239 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecSetIdentity[2]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v239
	goto L64
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+89)) = uint8(v214)
	F_CatalogTupleUpdate(m, v178, v181+int32(4), v181)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetIdentity[3]))
	if v222 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v225 = int32(*(*int16)(unsafe.Add(mBase, uint32(v207)+74)))
	v226 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v224, v225, v226, v226)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L18
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v232
	goto L64
L73:
	;
	goto L72
L74:
	;
	F_relation_close(m, v178, int32(3))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	v247 = int32(0)
	if base.B2i32(v171 == v247)|(base.B2i32(l5 == v247)|base.B2i32(v21 != int32(112))) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	m.G0 = v16 - int32(-64)
	return
L77:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v256 = F_find_inheritance_children(m, v255, l4)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	if v256 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v260 <= int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v270 = int32(0)
	goto L81
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v270<<(uint(int32(2))%32))))
	v285 = F_table_open(m, v283, int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L18
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v287 = int32(1)
	F_ATExecSetIdentity(m, v14+int32(-12), v285, l2, l3, l4, v287, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	F_relation_close(m, v285, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L18
	} else {
		goto L85
	}
L85:
	;
	v295 = v270 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v295 < v296 {
		v270 = v295
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L82
L87:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L18
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	F_errmsg(m, int32(_a_F_ATExecSetIdentity_14), v14+int32(-48))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L18
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_ATExecSetIdentity_1), int32(_a_F_ATExecSetIdentity_15), int32(_a_F_ATExecSetIdentity_3))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L18
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L18
	} else {
		goto L92
	}
L92:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v339 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecSetIdentity_16), v14+int32(-32))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L18
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_ATExecSetIdentity_1), int32(_a_F_ATExecSetIdentity_17), int32(_a_F_ATExecSetIdentity_3))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L18
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecSetRowSecurity(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v14 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v19 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(v11), int64(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				*(*uint8)(unsafe.Add(mBase, uint32(v21+v22)+127)) = uint8(v2)
				F_CatalogTupleUpdate(m, v14, v19+int32(4), v19)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetRowSecurity[0]))
					if v30 != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v33 = int32(0)
						F_RunObjectPostAlterHook(m, int32(1259), v32, v33, v33, v33)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_relation_close(m, v14, int32(3))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_pfree(m, v19)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									m.G0 = v9 + int32(16)
									return
								}
							}
						}
					} else {
						F_relation_close(m, v14, int32(3))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_pfree(m, v19)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					F_errmsg_internal(m, int32(_a_F_ATExecSetRowSecurity_0), v9)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ATExecSetRowSecurity_1), int32(_a_F_ATExecSetRowSecurity_2), int32(_a_F_ATExecSetRowSecurity_3))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
