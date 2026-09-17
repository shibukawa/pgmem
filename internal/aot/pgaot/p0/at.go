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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v321 int64
	_ = v321
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
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
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v572 int32
	_ = v572
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v630 int32
	_ = v630
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v690 int32
	_ = v690
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v773 int32
	_ = v773
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v895 int32
	_ = v895
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v964 int32
	_ = v964
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1021 int32
	_ = v1021
	var v1044 int32
	_ = v1044
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1187 int32
	_ = v1187
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1266 int32
	_ = v1266
	var v1282 int32
	_ = v1282
	var v1296 int32
	_ = v1296
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1584 int32
	_ = v1584
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1619 int32
	_ = v1619
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1656 int32
	_ = v1656
	var v1677 int32
	_ = v1677
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1748 int32
	_ = v1748
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1796 int32
	_ = v1796
	var v1799 int32
	_ = v1799
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1816 int32
	_ = v1816
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1897 int32
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1916 int32
	_ = v1916
	var v1920 int32
	_ = v1920
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1989 int32
	_ = v1989
	var v1994 int32
	_ = v1994
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2033 int32
	_ = v2033
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2118 int32
	_ = v2118
	var v2141 int32
	_ = v2141
	var v2150 int32
	_ = v2150
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2193 int64
	_ = v2193
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2238 int32
	_ = v2238
	var v2243 int32
	_ = v2243
	var v2247 int32
	_ = v2247
	var v2250 int32
	_ = v2250
	var v2254 int32
	_ = v2254
	var v2259 int32
	_ = v2259
	var v2263 int32
	_ = v2263
	var v2266 int32
	_ = v2266
	var v2273 int32
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2285 int32
	_ = v2285
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
	var v2317 int32
	_ = v2317
	var v2323 int32
	_ = v2323
	var v2328 int32
	_ = v2328
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2353 int32
	_ = v2353
	var v2358 int32
	_ = v2358
	var v2362 int32
	_ = v2362
	var v2371 int32
	_ = v2371
	var v2376 int32
	_ = v2376
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2415 int32
	_ = v2415
	var v2420 int32
	_ = v2420
	var v2424 int32
	_ = v2424
	var v2428 int32
	_ = v2428
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2472 int32
	_ = v2472
	var v2477 int32
	_ = v2477
	var v2481 int32
	_ = v2481
	var v2488 int32
	_ = v2488
	var v2493 int32
	_ = v2493
	var v2497 int32
	_ = v2497
	var v2504 int32
	_ = v2504
	var v2509 int32
	_ = v2509
	var v2540 int32
	_ = v2540
	var v2543 int32
	_ = v2543
	var v2547 int32
	_ = v2547
	var v2552 int32
	_ = v2552
	var v2583 int32
	_ = v2583
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2595 int32
	_ = v2595
	var v2600 int32
	_ = v2600
	var v2606 int32
	_ = v2606
	var v2628 int32
	_ = v2628
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2646 int32
	_ = v2646
	var v2651 int32
	_ = v2651
	v9 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(1632)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
	v36 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v38 - int32(1) {
	case 0, 4:
		goto L21
	default:
		goto L20
	case 8:
		goto L22
	}
L1:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v546)+12))
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2628+v2606<<(uint(int32(2))%32))))
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v2632)+4))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		goto L27
	} else {
		goto L514
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L27
	} else {
		goto L510
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L27
	} else {
		goto L506
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L27
	} else {
		goto L503
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L27
	} else {
		goto L500
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L27
	} else {
		goto L493
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L27
	} else {
		goto L490
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L27
	} else {
		goto L483
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L27
	} else {
		goto L480
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L27
	} else {
		goto L470
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L27
	} else {
		goto L467
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L27
	} else {
		goto L463
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L27
	} else {
		goto L459
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L27
	} else {
		goto L455
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L27
	} else {
		goto L451
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L27
	} else {
		goto L447
	}
L17:
	;
	m.G0 = v30 + int32(1632)
	return
L18:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v420)+56))
	v1638 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[2]))
	v1640 = F_pg_class_aclcheck(m, v1636, v1638, int64(32))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L27
	} else {
		goto L312
	}
L19:
	;
	if v532 == int32(0) {
		goto L3
	} else {
		goto L311
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L27
	} else {
		goto L308
	}
L21:
	;
	F_ATAddCheckNNConstraint(m, l0, l1, l2, l3, l4, l5, int32(0), l6, l7)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L27
	} else {
		goto L307
	}
L22:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v41 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v321 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1624)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1616)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1608)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1600)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1592)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1584)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1576)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1568)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1560)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1552)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1544)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1536)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1528)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1520)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1512)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1504)) = v321
	v355 = int32(0)
	v356 = int32(128)
	base.MemoryFill(m, v30+int32(1376), v355, v356)
	base.MemoryFill(m, v30+int32(1248), v355, v356)
	base.MemoryFill(m, v30+int32(1120), v355, v356)
	base.MemoryFill(m, v30+int32(992), v355, v356)
	base.MemoryFill(m, v30+int32(864), v355, v356)
	base.MemoryFill(m, v30+int32(736), v355, v356)
	base.MemoryFill(m, v30+int32(608), v355, v356)
	base.MemoryFill(m, v30+int32(480), v355, v356)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+472)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+464)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+456)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+448)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+440)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+432)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+424)) = v321
	*(*int64)(unsafe.Add(mBase, uint32(v30)+416)) = v321
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	if v409 != 0 {
		goto L77
	} else {
		goto L78
	}
L24:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v44 = F_ConstraintNameIsUsed(m, int32(0), v43, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+1376)) = uint8(v73)
	if v71 == v73 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	return
L28:
	;
	if v44 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v55 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_1), v30+int32(384))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_3), int32(_a_F_ATExecAddConstraint_4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
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
	v285 = F_pstrdup(m, v30+int32(1376))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L27
	} else {
		goto L75
	}
L35:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v79 <= int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v89 = int32(0)
	v92 = v9
	goto L37
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v92<<(uint(int32(2))%32))))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if int32(0) < v89 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L34
L39:
	;
	v121 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(1376)+v89))) = uint8(v121)
	v125 = v89 + int32(1)
	goto L41
L40:
	;
	v125 = v89
	goto L41
L41:
	;
	v128 = v30 + int32(1376) + v125
	goto L45
L42:
	;
	v248 = F_strlen(m, v128)
	mBase = m.M
	v249 = v248 + v125
	if int32(64) <= v249 {
		goto L34
	} else {
		goto L73
	}
L43:
	;
	v245 = F_strlen(m, v234)
	mBase = m.M
	goto L42
L45:
	;
	goto L46
L46:
	;
	v135 = int32(63)
	if (v128^v115)&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v238)
	goto L43
L48:
	;
	v219 = v214
	v220 = v215
	v221 = v216
	goto L69
L49:
	;
	if v209 == int32(0) {
		v234 = v207
		v235 = v208
		goto L47
	} else {
		goto L68
	}
L50:
	;
	v207 = v115
	v208 = v128
	v209 = v135
	goto L49
L51:
	;
	goto L52
L52:
	;
	v139 = int32(0)
	if base.B2i32(v115&int32(3) == v139)|int32(0) == v139 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v175 == int32(0) {
		v234 = v172
		v235 = v173
		goto L47
	} else {
		goto L62
	}
L54:
	;
	v151 = v115
	v152 = v128
	v153 = v135
	goto L57
L55:
	;
	goto L56
L56:
	;
	v172 = v115
	v173 = v128
	v174 = v135
	v175 = int32(1)
	goto L53
L57:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v155)
	if v155 == int32(0) {
		v214 = v151
		v215 = v152
		v216 = v153
		goto L48
	} else {
		goto L59
	}
L58:
	;
	v172 = v166
	v173 = v160
	v174 = v162
	v175 = v164
	goto L53
L59:
	;
	v159 = int32(1)
	v160 = v152 + v159
	v162 = v153 - v159
	v163 = int32(0)
	v164 = base.B2i32(v162 != v163)
	v166 = v151 + v159
	if v166&int32(3) == v163 {
		v172 = v166
		v173 = v160
		v174 = v162
		v175 = v164
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v162 != 0 {
		v151 = v166
		v152 = v160
		v153 = v162
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if base.B2i32(v178 == int32(0))|base.B2i32(base.Ui32(v174) < base.Ui32(int32(4))) != 0 {
		v207 = v172
		v208 = v173
		v209 = v174
		goto L49
	} else {
		goto L63
	}
L63:
	;
	v185 = v172
	v186 = v173
	v187 = v174
	goto L64
L64:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v193 = int32(-2139062144)
	if (int32(16843008)-v190|v190)&v193 != v193 {
		v214 = v185
		v215 = v186
		v216 = v187
		goto L48
	} else {
		goto L66
	}
L65:
	;
	v207 = v201
	v208 = v199
	v209 = v203
	goto L49
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v190
	v198 = int32(4)
	v199 = v186 + v198
	v201 = v185 + v198
	v203 = v187 - v198
	if base.Ui32(int32(3)) < base.Ui32(v203) {
		v185 = v201
		v186 = v199
		v187 = v203
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v214 = v207
	v215 = v208
	v216 = v209
	goto L48
L69:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	*(*uint8)(unsafe.Add(mBase, uint32(v220))) = uint8(v223)
	if v223 == int32(0) {
		v234 = v219
		v235 = v220
		goto L47
	} else {
		goto L71
	}
L70:
	;
	v234 = v230
	v235 = v228
	goto L47
L71:
	;
	v227 = int32(1)
	v228 = v220 + v227
	v230 = v219 + v227
	v232 = v221 - v227
	if v232 != 0 {
		v219 = v230
		v220 = v228
		v221 = v232
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v253 = v92 + int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v253 < v254 {
		v89 = v249
		v92 = v253
		goto L37
	} else {
		goto L74
	}
L74:
	;
	goto L38
L75:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+68))
	v291 = F_ChooseConstraintName(m, v72+int32(4), v285, int32(_a_F_ATExecAddConstraint_5), v289, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L27
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v291
	goto L23
L77:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+12))
	v411 = v410
	goto L79
L78:
	;
	v411 = v9
	goto L79
L79:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l4)+100))
	if v412 != 0 {
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
	v414 = F_table_open(m, v412, int32(6))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L27
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l4)+72))
	v418 = F_table_openrv(m, v416, int32(6))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L27
	} else {
		goto L85
	}
L84:
	;
	v420 = v414
	goto L80
L85:
	;
	v420 = v418
	goto L80
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L27
	} else {
		goto L303
	}
L87:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+119)))
	if v424 == int32(112) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+119)))
	switch v428 - int32(112) {
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
	v453 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATExecAddConstraint[3])))
	if v453 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L27
	} else {
		goto L93
	}
L93:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L27
	} else {
		goto L94
	}
L94:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v438 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_6), v30+int32(16))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L27
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_7), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
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
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L27
	} else {
		goto L299
	}
L98:
	;
	v457 = int32(1)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v420)+56))
	if base.Ui32(v458) < base.Ui32(int32(_a_F_ATExecAddConstraint_9)) {
		v467 = v457
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+118)))
	switch v469 - int32(112) {
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
	if v467 != 0 {
		goto L97
	} else {
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+68))
	if v462 == int32(99) {
		v467 = v457
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v465 = F_isTempToastNamespace(m, v462)
	mBase = m.M
	v467 = v465
	goto L102
L105:
	;
	goto L100
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L27
	} else {
		goto L295
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L27
	} else {
		goto L291
	}
L108:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v530 = F_transformColumnNameList(m, v522, v523, v30+int32(1504), v30+int32(1248), v30+int32(992))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L27
	} else {
		goto L125
	}
L109:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+118)))
	if v513 != int32(116) {
		goto L106
	} else {
		goto L122
	}
L110:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+118)))
	switch v493 - int32(112) {
	case 0, 5:
		goto L108
	default:
		goto L117
	}
L111:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+118)))
	if v473 == int32(112) {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L27
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L27
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_10), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L27
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_11), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
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
	v499 = m.ExcPending
	if v499 != 0 {
		goto L27
	} else {
		goto L118
	}
L118:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L27
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_12), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L27
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_13), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
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
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+24)))
	if v516 != int32(1) {
		goto L107
	} else {
		goto L123
	}
L123:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)))
	if v519 == int32(0) {
		goto L107
	} else {
		goto L124
	}
L124:
	;
	goto L108
L125:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v532 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L27
	} else {
		goto L287
	}
L127:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+85)))
	if v535 == int32(1) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v542 = int32(0)
	v544 = F_transformColumnNameList(m, v538, v539, v30+int32(416), v542, v542)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L27
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	if v544 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v547 = int32(0)
	if v530 == v547 {
		v2606 = v547
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v748 = v9
	goto L134
L134:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	if v753 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L135:
	;
	v555 = v547
	v572 = v9
	goto L136
L136:
	;
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(416)+v555<<(uint(int32(1))%32)))))
	v590 = int32(0)
	goto L138
L137:
	;
	v748 = v718
	goto L134
L138:
	;
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1504)+v590<<(uint(int32(1))%32)))))
	if v583 != v616 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v621 = int32(0)
	if v621 < v572 {
		goto L145
	} else {
		goto L146
	}
L140:
	;
	v619 = v590 + int32(1)
	if v530 != v619 {
		v590 = v619
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
	v2606 = v555
	goto L1
L144:
	;
	v724 = v555 + int32(1)
	if v724 != v544 {
		v555 = v724
		v572 = v718
		goto L136
	} else {
		goto L152
	}
L145:
	;
	v630 = v621
	goto L148
L146:
	;
	goto L147
L147:
	;
	v690 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(416)+v572<<(uint(v690)%32)))) = uint16(v583)
	v718 = v572 + v690
	goto L144
L148:
	;
	v656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(416)+v630<<(uint(int32(1))%32)))))
	if v656 == v583 {
		v718 = v572
		goto L144
	} else {
		goto L150
	}
L149:
	;
	goto L147
L150:
	;
	v659 = v630 + int32(1)
	if v659 != v572 {
		v630 = v659
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
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L27
	} else {
		goto L284
	}
L154:
	;
	v756 = F_RelationGetIndexList(m, v420)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L27
	} else {
		goto L159
	}
L155:
	;
	goto L156
L156:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v420)+56))
	v1010 = F_transformColumnNameList(m, v1003, v753, v30+int32(1568), v30+int32(1376), v30+int32(1120))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L27
	} else {
		goto L199
	}
L157:
	;
	F_list_free(m, v756)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L27
	} else {
		goto L177
	}
L158:
	;
	F_list_free(m, v756)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L27
	} else {
		goto L176
	}
L159:
	;
	if v756 == int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	if v760 <= int32(0) {
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v773 = int32(0)
	goto L162
L162:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v756)+12))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v792+v773<<(uint(int32(2))%32))))
	v797 = F_SearchSysCache1(m, int32(34), v796)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L27
	} else {
		goto L164
	}
L163:
	;
	goto L158
L164:
	;
	if v797 == int32(0) {
		goto L153
	} else {
		goto L165
	}
L165:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v797)+16))
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801)+22)))
	v803 = v801 + v802
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+14)))
	if v804 != int32(1) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_ReleaseCatCache(m, v797)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L27
	} else {
		goto L174
	}
L167:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+18)))
	if v807 != int32(1) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+16)))
	if v810 != 0 {
		goto L157
	} else {
		goto L169
	}
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L27
	} else {
		goto L170
	}
L170:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L27
	} else {
		goto L171
	}
L171:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v818 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_14), v30+int32(272))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L27
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_15), int32(_a_F_ATExecAddConstraint_16))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
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
	v835 = v773 + int32(1)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	if v835 < v836 {
		v773 = v835
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
	if v796 == int32(0) {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	v871 = int32(0)
	v874 = F_SysCacheGetAttrNotNull(m, int32(34), v797, int32(18))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L27
	} else {
		goto L179
	}
L179:
	;
	v876 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v876
	v878 = int32(*(*int16)(unsafe.Add(mBase, uint32(v803)+10)))
	if v876 < v878 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v895 = v871
	goto L183
L181:
	;
	v964 = v871
	goto L182
L182:
	;
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+15)))
	F_ReleaseCatCache(m, v797)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L27
	} else {
		goto L192
	}
L183:
	;
	v913 = v895 << (uint(int32(1)) % 32)
	v918 = int32(*(*int16)(unsafe.Add(mBase, uint32(v913+(v803+int32(48))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v913+(v30+int32(1568))))) = uint16(v918)
	v921 = v895 << (uint(int32(2)) % 32)
	v925 = F_attnumTypeId(m, v420, v918)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L27
	} else {
		goto L185
	}
L184:
	;
	v964 = v951
	goto L182
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v921+(v30+int32(1376))))) = v925
	v931 = F_attnumCollationId(m, v420, v918)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L27
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(1120)+v921))) = v931
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v921+(v874+int32(24)))))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(864)+v921))) = v938
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v941 = F_attnumAttName(m, v420, v918)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L27
	} else {
		goto L187
	}
L187:
	;
	v943 = F_pstrdup(m, v941)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L27
	} else {
		goto L188
	}
L188:
	;
	v945 = F_makeString(m, v943)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L27
	} else {
		goto L189
	}
L189:
	;
	v947 = F_lappend(m, v940, v945)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L27
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v947
	v951 = v895 + int32(1)
	v952 = int32(*(*int16)(unsafe.Add(mBase, uint32(v803)+10)))
	if v951 < v952 {
		v895 = v951
		goto L183
	} else {
		goto L191
	}
L191:
	;
	goto L184
L192:
	;
	if v981 != int32(1) {
		v1619 = v964
		v1632 = v796
		goto L18
	} else {
		goto L193
	}
L193:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v986 != 0 {
		goto L19
	} else {
		goto L194
	}
L194:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L27
	} else {
		goto L195
	}
L195:
	;
	F_errcode(m, int32(_a_F_ATExecAddConstraint_17))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L27
	} else {
		goto L196
	}
L196:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_18), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L27
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_19), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
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
	if v532 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L27
	} else {
		goto L280
	}
L201:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+85)))
	if v1012 == int32(0) {
		goto L200
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if v1010 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L203
L205:
	;
	v1021 = int32(0)
	goto L208
L206:
	;
	goto L207
L207:
	;
	v1160 = F_RelationGetIndexList(m, v420)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L27
	} else {
		goto L227
	}
L208:
	;
	v1044 = v1021 + int32(1)
	if base.Ui32(v1010) <= base.Ui32(v1044) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	goto L207
L210:
	;
	if v1044 != v1010 {
		v1021 = v1044
		goto L208
	} else {
		goto L222
	}
L211:
	;
	v1051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1021<<(uint(int32(1))%32)))))
	v1058 = v1044
	goto L212
L212:
	;
	v1084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1058<<(uint(int32(1))%32)))))
	if v1084 != v1051 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L27
	} else {
		goto L218
	}
L214:
	;
	v1087 = v1058 + int32(1)
	if v1010 != v1087 {
		v1058 = v1087
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
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L27
	} else {
		goto L219
	}
L219:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_20), int32(0))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L27
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_21), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
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
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L27
	} else {
		goto L276
	}
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L27
	} else {
		goto L273
	}
L225:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+15)))
	F_ReleaseCatCache(m, v1211)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L27
	} else {
		goto L270
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L27
	} else {
		goto L266
	}
L227:
	;
	if v1160 == int32(0) {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+4))
	if v1164 <= int32(0) {
		goto L226
	} else {
		goto L229
	}
L229:
	;
	v1167 = int32(0)
	v1170 = int32(1)
	v1173 = (v1010 - v1170) << (uint(v1170) % 32)
	v1187 = v1167
	v1204 = v9
	goto L230
L230:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+12))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1206+v1187<<(uint(int32(2))%32))))
	v1211 = F_SearchSysCache1(m, int32(34), v1210)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L27
	} else {
		goto L232
	}
L231:
	;
	if v1364 != 0 {
		goto L223
	} else {
		goto L265
	}
L232:
	;
	if v1211 == int32(0) {
		goto L224
	} else {
		goto L233
	}
L233:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+16))
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1215)+22)))
	v1217 = v1215 + v1216
	v1218 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1217)+10)))
	if v1010 != v1218 {
		v1364 = v1204
		goto L234
	} else {
		goto L235
	}
L234:
	;
	F_ReleaseCatCache(m, v1211)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L27
	} else {
		goto L263
	}
L235:
	;
	if v532 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+18)))
	if v1224 != int32(1) {
		v1364 = v1204
		goto L234
	} else {
		goto L242
	}
L237:
	;
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+15)))
	if v1220 != 0 {
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+12)))
	if v1221 != int32(1) {
		v1364 = v1204
		goto L234
	} else {
		goto L241
	}
L240:
	;
	v1364 = v1204
	goto L234
L241:
	;
	goto L236
L242:
	;
	v1229 = F_heap_attisnull(m, v1211, int32(21), int32(0))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L27
	} else {
		goto L243
	}
L243:
	;
	if v1229 == int32(0) {
		v1364 = v1204
		goto L234
	} else {
		goto L244
	}
L244:
	;
	v1235 = F_heap_attisnull(m, v1211, int32(20), int32(0))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L27
	} else {
		goto L245
	}
L245:
	;
	if v1235 == int32(0) {
		v1364 = v1204
		goto L234
	} else {
		goto L246
	}
L246:
	;
	v1241 = F_SysCacheGetAttrNotNull(m, int32(34), v1211, int32(18))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L27
	} else {
		goto L247
	}
L247:
	;
	if v1010 == int32(0) {
		v1364 = v1204
		goto L234
	} else {
		goto L248
	}
L248:
	;
	v1248 = v1217 + int32(48)
	v1266 = int32(0)
	goto L249
L249:
	;
	v1282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1266<<(uint(int32(1))%32)))))
	v1296 = int32(0)
	goto L251
L250:
	;
	if base.B2i32(v1010 != v1167)&v532 != 0 {
		goto L258
	} else {
		goto L259
	}
L251:
	;
	v1314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248+v1296<<(uint(int32(1))%32)))))
	if v1314 != v1282 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1321 = int32(2)
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1241+int32(24)+v1296<<(uint(v1321)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(864)+v1266<<(uint(v1321)%32)))) = v1327
	v1330 = v1266 + int32(1)
	if v1330 != v1010 {
		v1266 = v1330
		goto L249
	} else {
		goto L257
	}
L253:
	;
	v1317 = v1296 + int32(1)
	if v1317 != v1010 {
		v1296 = v1317
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
	v1364 = v1204
	goto L234
L257:
	;
	goto L250
L258:
	;
	v1332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173+(v30+int32(1568))))))
	v1334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173+v1248))))
	if v1332 != v1334 {
		v1364 = v1204
		goto L234
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+16)))
	if v1337 != 0 {
		goto L225
	} else {
		goto L262
	}
L261:
	;
	goto L260
L262:
	;
	v1364 = int32(1)
	goto L234
L263:
	;
	v1368 = v1187 + int32(1)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+4))
	if v1368 < v1369 {
		v1187 = v1368
		v1204 = v1364
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
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L27
	} else {
		goto L267
	}
L267:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+288)) = v1405 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_23), v30+int32(288))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L27
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_24), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
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
	F_list_free(m, v1160)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L27
	} else {
		goto L271
	}
L271:
	;
	if v532|base.B2i32(v1419&int32(1) == int32(0)) != 0 {
		v1619 = v1010
		v1632 = v1210
		goto L18
	} else {
		goto L272
	}
L272:
	;
	goto L3
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+304)) = v1210
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_25), v30+int32(304))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L27
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_26), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
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
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L27
	} else {
		goto L277
	}
L277:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+320)) = v1451 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_27), v30+int32(320))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L27
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_28), int32(_a_F_ATExecAddConstraint_22))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
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
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L27
	} else {
		goto L281
	}
L281:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_29), int32(0))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L27
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_30), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v796
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_25), v30+int32(48))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L27
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_31), int32(_a_F_ATExecAddConstraint_16))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
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
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L27
	} else {
		goto L288
	}
L288:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_18), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L27
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_32), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
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
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L27
	} else {
		goto L292
	}
L292:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_33), int32(0))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L27
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_34), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
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
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L27
	} else {
		goto L296
	}
L296:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_35), int32(0))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L27
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_36), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
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
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L27
	} else {
		goto L300
	}
L300:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+352)) = v1551 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_37), v30+int32(352))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L27
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_38), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
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
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L27
	} else {
		goto L304
	}
L304:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v1574 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v1573 + v1574
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v1572 + v1574
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_39), v30+int32(368))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L27
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_40), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
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
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1597
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_41), v30)
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L27
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_42), int32(_a_F_ATExecAddConstraint_4))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
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
	v1619 = v964
	v1632 = v796
	goto L18
L312:
	;
	v1642 = int32(0)
	if base.B2i32(v1640 == v1642)|base.B2i32(v1619 <= v1642) == v1642 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1656 = int32(0)
	goto L316
L314:
	;
	goto L315
L315:
	;
	if v530 != 0 {
		goto L331
	} else {
		goto L332
	}
L316:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v420)+56))
	v1683 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1656<<(uint(int32(1))%32)))))
	v1685 = F_pg_attribute_aclcheck(m, v1677, v1683, v1638, int64(32))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L27
	} else {
		goto L318
	}
L317:
	;
	goto L315
L318:
	;
	if v1685 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	v1688 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1687)+119)))
	switch v1688 - int32(73) {
	case 0, 32:
		v1698 = int32(20)
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
	v1707 = v1656 + int32(1)
	if v1707 != v1619 {
		v1656 = v1707
		goto L316
	} else {
		goto L330
	}
L322:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	F_aclcheck_error(m, v1685, v1700, v1701+int32(4))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L27
	} else {
		goto L329
	}
L323:
	;
	v1700 = v1698
	goto L322
L324:
	;
	v1698 = int32(41)
	goto L323
L325:
	;
	v1700 = int32(18)
	goto L322
L326:
	;
	v1700 = int32(23)
	goto L322
L327:
	;
	v1700 = int32(51)
	goto L322
L328:
	;
	v1700 = int32(37)
	goto L322
L329:
	;
	goto L321
L330:
	;
	goto L317
L331:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1736)))
	v1748 = int32(0)
	goto L334
L332:
	;
	goto L333
L333:
	;
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v1845 != int32(1) {
		goto L351
	} else {
		goto L352
	}
L334:
	;
	v1774 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(1504)+v1748<<(uint(int32(1))%32)))))
	v1778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1736+v1737<<(uint(int32(4))%32)+v1774*int32(100))+10)))
	if v1778 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	goto L333
L336:
	;
	v1779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+87)))
	v1781 = v1779 - int32(99)
	if int32(1)<<(uint(v1781)%32)&int32(2051) != 0 {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	goto L338
L338:
	;
	v1816 = v1748 + int32(1)
	if v1816 != v530 {
		v1748 = v1816
		goto L334
	} else {
		goto L350
	}
L339:
	;
	v1789 = base.B2i32(base.Ui32(v1781) <= base.Ui32(int32(11)))
	goto L341
L340:
	;
	v1789 = int32(0)
	goto L341
L341:
	;
	if v1789 != 0 {
		goto L16
	} else {
		goto L342
	}
L342:
	;
	v1790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+88)))
	switch v1790 - int32(100) {
	case 0, 10:
		goto L344
	default:
		goto L343
	}
L343:
	;
	if v1778 == int32(118) {
		goto L15
	} else {
		goto L349
	}
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L27
	} else {
		goto L345
	}
L345:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L27
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+240)) = int32(_a_F_ATExecAddConstraint_43)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_44), v30+int32(240))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L27
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_45), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
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
	if v1619 != v530 {
		goto L12
	} else {
		goto L359
	}
L352:
	;
	v1848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+87)))
	v1850 = v1848 - int32(99)
	if int32(1)<<(uint(v1850)%32)&int32(_a_F_ATExecAddConstraint_46) != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1858 = base.B2i32(base.Ui32(v1850) <= base.Ui32(int32(15)))
	goto L355
L354:
	;
	v1858 = int32(0)
	goto L355
L355:
	;
	if v1858 != 0 {
		goto L14
	} else {
		goto L356
	}
L356:
	;
	v1859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+88)))
	v1861 = v1859 - int32(99)
	if base.Ui32(int32(15)) < base.Ui32(v1861) {
		goto L351
	} else {
		goto L357
	}
L357:
	;
	if int32(1)<<(uint(v1861)%32)&int32(_a_F_ATExecAddConstraint_46) != 0 {
		goto L13
	} else {
		goto L358
	}
L358:
	;
	goto L351
L359:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	v1872 = base.B2i32(v1870 != int32(0))
	if v530 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1886 = int32(0)
	v1889 = v411
	v1897 = v1872
	goto L363
L361:
	;
	v2141 = v1872
	goto L362
L362:
	;
	if v532 != 0 {
		goto L439
	} else {
		goto L440
	}
L363:
	;
	v1904 = v1886 << (uint(int32(2)) % 32)
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1904+(v30+int32(992)))))
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(1120)+v1904)))
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(1248)+v1904)))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(1376)+v1904)))
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(864)+v1904)))
	v1926 = F_SearchSysCache1(m, int32(14), v1925)
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L27
	} else {
		goto L365
	}
L364:
	;
	v2141 = v2103
	goto L362
L365:
	;
	if v1926 == int32(0) {
		goto L11
	} else {
		goto L366
	}
L366:
	;
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+16))
	v1931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1930)+22)))
	v1932 = v1930 + v1931
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+84))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+80))
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+4))
	F_ReleaseCatCache(m, v1926)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L27
	} else {
		goto L367
	}
L367:
	;
	v1941 = v532 & base.B2i32(v1886 == v530-int32(1))
	if v1941 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1942 = int32(7)
	goto L370
L369:
	;
	v1942 = int32(3)
	goto L370
L370:
	;
	v1944 = F_IndexAmTranslateCompareType(m, v1942, v1935, v1934, int32(1))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L27
	} else {
		goto L371
	}
L371:
	;
	if v1944 == int32(0) {
		goto L10
	} else {
		goto L372
	}
L372:
	;
	v1948 = base.I32_extend16_s(v1944)
	v1949 = F_get_opfamily_member(m, v1934, v1933, v1933, v1948)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L27
	} else {
		goto L373
	}
L373:
	;
	if v1949 == int32(0) {
		goto L9
	} else {
		goto L374
	}
L374:
	;
	v1953 = F_getBaseType(m, v1916)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L27
	} else {
		goto L377
	}
L375:
	;
	v1984 = int32(0)
	if base.B2i32(v1982 == v1984)|base.B2i32(v1983 == v1984) != 0 {
		goto L8
	} else {
		goto L394
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v30)+412)) = v1933
	*(*int32)(unsafe.Add(mBase, uint32(v30)+408)) = v1933
	v1974 = F_can_coerce_type(m, int32(2), v30+int32(392), v30+int32(408), int32(0))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L27
	} else {
		goto L384
	}
L377:
	;
	v1955 = F_get_opfamily_member(m, v1934, v1933, v1953, v1948)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L27
	} else {
		goto L378
	}
L378:
	;
	if v1955 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1962 = int32(0)
	goto L376
L380:
	;
	goto L381
L381:
	;
	v1960 = F_get_opfamily_member(m, v1934, v1953, v1953, v1948)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L27
	} else {
		goto L382
	}
L382:
	;
	if v1960 != 0 {
		v1980 = v1953
		v1982 = v1955
		v1983 = v1960
		goto L375
	} else {
		goto L383
	}
L383:
	;
	v1962 = v1953
	goto L376
L384:
	;
	if v1974 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1976 = v1949
	goto L387
L386:
	;
	v1976 = v1955
	goto L387
L387:
	;
	if v1974 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v1978 = v1949
	goto L390
L389:
	;
	v1978 = int32(0)
	goto L390
L390:
	;
	if v1974 != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1979 = v1933
	goto L393
L392:
	;
	v1979 = v1962
	goto L393
L393:
	;
	v1980 = v1979
	v1982 = v1976
	v1983 = v1978
	goto L375
L394:
	;
	v1989 = int32(0)
	if base.B2i32(v1912 != v1989) != base.B2i32(v1908 != v1989) {
		goto L7
	} else {
		goto L395
	}
L395:
	;
	v1994 = int32(0)
	if base.B2i32(v1912 == v1994)|base.B2i32(v1908 == v1994) != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2007 = int32(0)
	if v1897&int32(1) == v2007 {
		goto L403
	} else {
		goto L404
	}
L397:
	;
	v1999 = F_get_collation_isdeterministic(m, v1912)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L27
	} else {
		goto L398
	}
L398:
	;
	v2001 = F_get_collation_isdeterministic(m, v1908)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L27
	} else {
		goto L399
	}
L399:
	;
	if v1999&v2001 != 0 {
		goto L396
	} else {
		goto L400
	}
L400:
	;
	if v1908 != v1912 {
		goto L6
	} else {
		goto L401
	}
L401:
	;
	goto L396
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(608)+v1904))) = v1949
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(736)+v1904))) = v1982
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(480)+v1904))) = v1983
	v2118 = v1886 + int32(1)
	if v2118 != v530 {
		v1886 = v2118
		v1889 = v2101
		v1897 = v2103
		goto L363
	} else {
		goto L438
	}
L403:
	;
	v2101 = v1889
	v2103 = v2007
	goto L402
L404:
	;
	goto L405
L405:
	;
	v2011 = v1889 + int32(4)
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2013)+12))
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v2013)+4))
	if base.Ui32(v2011) < base.Ui32(v2014+v2015<<(uint(int32(2))%32)) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v2020 = v2011
	goto L408
L407:
	;
	v2020 = int32(0)
	goto L408
L408:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v1889)))
	if v1982 != v2021 {
		v2101 = v2020
		v2103 = v2007
		goto L402
	} else {
		goto L409
	}
L409:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v2023)))
	v2033 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(1504)+v1886<<(uint(int32(1))%32)))))
	v2038 = v2023 + v2024<<(uint(int32(4))%32) + v2033*int32(100) - int32(80)
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+68))
	if v2039 == v1980 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	if v1980 == v1916 {
		goto L417
	} else {
		goto L418
	}
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = int32(0)
	v2051 = int32(2)
	goto L410
L412:
	;
	goto L413
L413:
	;
	v2047 = F_find_coercion_pathway(m, v1980, v2039, int32(0), v30+int32(392))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L27
	} else {
		goto L414
	}
L414:
	;
	if v2047 == int32(0) {
		goto L5
	} else {
		goto L415
	}
L415:
	;
	v2051 = v2047
	goto L410
L416:
	;
	if v2051 != v2063 {
		v2101 = v2020
		v2103 = v2007
		goto L402
	} else {
		goto L422
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+408)) = int32(0)
	v2063 = int32(2)
	goto L416
L418:
	;
	goto L419
L419:
	;
	v2059 = F_find_coercion_pathway(m, v1980, v1916, int32(0), v30+int32(408))
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L27
	} else {
		goto L420
	}
L420:
	;
	if v2059 == int32(0) {
		goto L4
	} else {
		goto L421
	}
L421:
	;
	v2063 = v2059
	goto L416
L422:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v30)+408))
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v30)+392))
	if v2065 != v2066 {
		v2101 = v2020
		v2103 = v2007
		goto L402
	} else {
		goto L423
	}
L423:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+96))
	if v1980 <= int32(3830) {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	if v2068 == v1908 {
		v2101 = v2020
		v2103 = int32(1)
		goto L402
	} else {
		goto L434
	}
L425:
	;
	if v2039 != v1916 {
		v2101 = v2020
		v2103 = v2007
		goto L402
	} else {
		goto L433
	}
L426:
	;
	switch v1980 - int32(2277) {
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
	if base.B2i32(base.Ui32(v1980-int32(_a_F_ATExecAddConstraint_47)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v1980-int32(_a_F_ATExecAddConstraint_48)) < base.Ui32(int32(2))) != 0 {
		goto L425
	} else {
		goto L431
	}
L429:
	;
	if base.B2i32(v1980 == int32(2776))|base.B2i32(v1980 == int32(3500)) != 0 {
		goto L425
	} else {
		goto L430
	}
L430:
	;
	goto L424
L431:
	;
	if v1980 != int32(3831) {
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
	v2093 = F_get_collation_isdeterministic(m, v2068)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L27
	} else {
		goto L435
	}
L435:
	;
	if v2093 == int32(0) {
		v2101 = v2020
		v2103 = int32(0)
		goto L402
	} else {
		goto L436
	}
L436:
	;
	v2097 = F_get_collation_isdeterministic(m, v1908)
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L27
	} else {
		goto L437
	}
L437:
	;
	v2101 = v2020
	v2103 = v2097
	goto L402
L438:
	;
	goto L364
L439:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v530<<(uint(int32(2))%32)+v30)+860))
	F_FindFKPeriodOpers(m, v2150, v30+int32(392), v30+int32(408), v30+int32(404))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L27
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v2163 = int32(0)
	v2165 = v30 + int32(1568)
	v2167 = v30 + int32(1504)
	v2169 = v30 + int32(736)
	v2171 = v30 + int32(608)
	v2173 = v30 + int32(480)
	v2175 = v30 + int32(416)
	F_addFkConstraint(m, v30+int32(392), int32(2), v2162, l4, l3, v420, v1632, v2163, v530, v2165, v2167, v2169, v2171, v2173, v748, v2175, v2163, v532)
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L27
	} else {
		goto L443
	}
L442:
	;
	goto L441
L443:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v30)+396))
	v2180 = int32(0)
	F_addFkRecurseReferenced(m, l4, l3, v420, v1632, v2179, v530, v2165, v2167, v2169, v2171, v2173, v748, v2175, v2180, v2180, v532)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L27
	} else {
		goto L444
	}
L444:
	;
	v2184 = int32(0)
	F_addFkRecurseReferencing(m, l1, l4, l3, v420, v1632, v2179, v530, v2165, v2167, v2169, v2171, v2173, v748, v2175, v2141, l7, v2184, v2184, v532)
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L27
	} else {
		goto L445
	}
L445:
	;
	F_relation_close(m, v420, int32(0))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L27
	} else {
		goto L446
	}
L446:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v30)+400))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2191
	v2193 = *(*int64)(unsafe.Add(mBase, uint32(v30)+392))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v2193
	goto L17
L447:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L27
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = int32(_a_F_ATExecAddConstraint_49)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_44), v30+int32(256))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L27
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_50), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
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
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L27
	} else {
		goto L452
	}
L452:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_51), int32(0))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L27
	} else {
		goto L453
	}
L453:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_52), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
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
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L27
	} else {
		goto L456
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+224)) = int32(_a_F_ATExecAddConstraint_49)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_53), v30+int32(224))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L27
	} else {
		goto L457
	}
L457:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_54), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
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
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L27
	} else {
		goto L460
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+208)) = int32(_a_F_ATExecAddConstraint_43)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_53), v30+int32(208))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L27
	} else {
		goto L461
	}
L461:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_55), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
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
	v2304 = m.ExcPending
	if v2304 != 0 {
		goto L27
	} else {
		goto L464
	}
L464:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_56), int32(0))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L27
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_57), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v1925
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_58), v30-int32(-64))
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L27
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_59), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
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
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L27
	} else {
		goto L471
	}
L471:
	;
	if v1941 != 0 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2338 = int32(_a_F_ATExecAddConstraint_60)
	goto L474
L473:
	;
	v2338 = int32(_a_F_ATExecAddConstraint_61)
	goto L474
L474:
	;
	F_errmsg(m, v2338, int32(0))
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L27
	} else {
		goto L475
	}
L475:
	;
	v2342 = F_get_opfamily_name(m, v1934)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L27
	} else {
		goto L476
	}
L476:
	;
	v2344 = F_get_am_name(m, v1935)
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L27
	} else {
		goto L477
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+88)) = v2344
	*(*int32)(unsafe.Add(mBase, uint32(v30)+84)) = v2342
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v1942
	F_errdetail(m, int32(_a_F_ATExecAddConstraint_62), v30+int32(80))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L27
	} else {
		goto L478
	}
L478:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_63), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = v1934
	*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = v1933
	*(*int32)(unsafe.Add(mBase, uint32(v30)+100)) = v1933
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v1948
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_64), v30+int32(96))
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L27
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_65), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
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
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L27
	} else {
		goto L484
	}
L484:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+192)) = v2384
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_66), v30+int32(192))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L27
	} else {
		goto L485
	}
L485:
	;
	v2392 = v1886 << (uint(int32(2)) % 32)
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+12))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2392+v2394)))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+4))
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v2398)+12))
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2399+v2392)))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v2401)+4))
	v2403 = F_format_type_be(m, v1916)
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L27
	} else {
		goto L486
	}
L486:
	;
	v2405 = F_format_type_be(m, v1920)
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L27
	} else {
		goto L487
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+188)) = v2405
	*(*int32)(unsafe.Add(mBase, uint32(v30)+184)) = v2403
	*(*int32)(unsafe.Add(mBase, uint32(v30)+180)) = v2402
	*(*int32)(unsafe.Add(mBase, uint32(v30)+176)) = v2397
	F_errdetail(m, int32(_a_F_ATExecAddConstraint_67), v30+int32(176))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L27
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_68), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2420 = m.ExcPending
	if v2420 != 0 {
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
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L27
	} else {
		goto L491
	}
L491:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_70), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
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
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L27
	} else {
		goto L494
	}
L494:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v2441
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_66), v30+int32(160))
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L27
	} else {
		goto L495
	}
L495:
	;
	v2449 = v1886 << (uint(int32(2)) % 32)
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v2450)+12))
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2449+v2451)))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2453)+4))
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2455)+12))
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v2456+v2449)))
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2458)+4))
	v2460 = F_get_collation_name(m, v1908)
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L27
	} else {
		goto L496
	}
L496:
	;
	v2462 = F_get_collation_name(m, v1912)
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L27
	} else {
		goto L497
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+156)) = v2462
	*(*int32)(unsafe.Add(mBase, uint32(v30)+152)) = v2460
	*(*int32)(unsafe.Add(mBase, uint32(v30)+148)) = v2459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = v2454
	F_errdetail(m, int32(_a_F_ATExecAddConstraint_71), v30+int32(144))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L27
	} else {
		goto L498
	}
L498:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_72), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+132)) = v1980
	*(*int32)(unsafe.Add(mBase, uint32(v30)+128)) = v2039
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_73), v30+int32(128))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L27
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_74), int32(_a_F_ATExecAddConstraint_75))
	mBase = m.M
	v2493 = m.ExcPending
	if v2493 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+116)) = v1980
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v1916
	F_errmsg_internal(m, int32(_a_F_ATExecAddConstraint_73), v30+int32(112))
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L27
	} else {
		goto L504
	}
L504:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_74), int32(_a_F_ATExecAddConstraint_75))
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
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
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L27
	} else {
		goto L507
	}
L507:
	;
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_76), int32(0))
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		goto L27
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_77), int32(_a_F_ATExecAddConstraint_8))
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
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
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L27
	} else {
		goto L511
	}
L511:
	;
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v420)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v2587 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_78), v30+int32(32))
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L27
	} else {
		goto L512
	}
L512:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_79), int32(_a_F_ATExecAddConstraint_16))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
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
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L27
	} else {
		goto L515
	}
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+336)) = v2633
	F_errmsg(m, int32(_a_F_ATExecAddConstraint_81), v30+int32(336))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L27
	} else {
		goto L516
	}
L516:
	;
	F_errfinish(m, int32(_a_F_ATExecAddConstraint_2), int32(_a_F_ATExecAddConstraint_82), int32(_a_F_ATExecAddConstraint_83))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
		v18 = F_SearchSysCacheCopy(m, int32(57), v11, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
				*(*uint8)(unsafe.Add(mBase, uint32(v20+v21)+127)) = uint8(v2)
				F_CatalogTupleUpdate(m, v14, v18+int32(4), v18)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetRowSecurity[0]))
					if v29 != 0 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v32 = int32(0)
						F_RunObjectPostAlterHook(m, int32(1259), v31, v32, v32, v32)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_relation_close(m, v14, int32(3))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
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
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_pfree(m, v18)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					F_errmsg_internal(m, int32(_a_F_ATExecSetRowSecurity_0), v9)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ATExecSetRowSecurity_1), int32(_a_F_ATExecSetRowSecurity_2), int32(_a_F_ATExecSetRowSecurity_3))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
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
