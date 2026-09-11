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
	var v33 int64
	_ = v33
	var v36 int32
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v375 int64
	_ = v375
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v674 int32
	_ = v674
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v714 int32
	_ = v714
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v774 int32
	_ = v774
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v857 int32
	_ = v857
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
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
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v980 int32
	_ = v980
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1049 int32
	_ = v1049
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1106 int32
	_ = v1106
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1143 int32
	_ = v1143
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1272 int32
	_ = v1272
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1353 int32
	_ = v1353
	var v1367 int32
	_ = v1367
	var v1379 int32
	_ = v1379
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1406 int32
	_ = v1406
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1445 int32
	_ = v1445
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1500 int32
	_ = v1500
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1524 int32
	_ = v1524
	var v1532 int32
	_ = v1532
	var v1540 int32
	_ = v1540
	var v1556 int32
	_ = v1556
	var v1564 int32
	_ = v1564
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1589 int32
	_ = v1589
	var v1610 int32
	_ = v1610
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1683 int32
	_ = v1683
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1741 int32
	_ = v1741
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1822 int32
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1834 int32
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1923 int32
	_ = v1923
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1966 int32
	_ = v1966
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1984 int32
	_ = v1984
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2049 int32
	_ = v2049
	var v2074 int32
	_ = v2074
	var v2081 int32
	_ = v2081
	var v2089 int32
	_ = v2089
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2139 int32
	_ = v2139
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2148 int64
	_ = v2148
	var v2153 int32
	_ = v2153
	var v2159 int32
	_ = v2159
	var v2164 int32
	_ = v2164
	var v2168 int32
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2180 int32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2215 int32
	_ = v2215
	var v2220 int32
	_ = v2220
	var v2224 int32
	_ = v2224
	var v2227 int32
	_ = v2227
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2262 int32
	_ = v2262
	var v2269 int32
	_ = v2269
	var v2274 int32
	_ = v2274
	var v2278 int32
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2300 int32
	_ = v2300
	var v2305 int32
	_ = v2305
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2330 int32
	_ = v2330
	var v2335 int32
	_ = v2335
	var v2339 int32
	_ = v2339
	var v2348 int32
	_ = v2348
	var v2353 int32
	_ = v2353
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2367 int32
	_ = v2367
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2392 int32
	_ = v2392
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2405 int32
	_ = v2405
	var v2410 int32
	_ = v2410
	var v2414 int32
	_ = v2414
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2424 int32
	_ = v2424
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2449 int32
	_ = v2449
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2465 int32
	_ = v2465
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2481 int32
	_ = v2481
	var v2486 int32
	_ = v2486
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2497 int32
	_ = v2497
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2512 int32
	_ = v2512
	var v2517 int32
	_ = v2517
	var v2521 int32
	_ = v2521
	var v2524 int32
	_ = v2524
	var v2528 int32
	_ = v2528
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2565 int32
	_ = v2565
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2581 int32
	_ = v2581
	var v2586 int32
	_ = v2586
	var v2590 int32
	_ = v2590
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2606 int32
	_ = v2606
	var v2611 int32
	_ = v2611
	var v2614 int32
	_ = v2614
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2658 int32
	_ = v2658
	var v2689 int32
	_ = v2689
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2701 int32
	_ = v2701
	var v2706 int32
	_ = v2706
	var v2712 int32
	_ = v2712
	var v2734 int32
	_ = v2734
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2743 int32
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2752 int32
	_ = v2752
	var v2757 int32
	_ = v2757
	v9 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(1632)
	m.G0 = v30
	v33 = *(*int64)(unsafe.Add(mBase, _consts[224]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v33
	v36 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v38 - int32(1) {
	case 0, 4:
		goto L5
	default:
		goto L3
	case 8:
		goto L6
	}
L1:
	;
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2734+v2712<<(uint(int32(2))%32))))
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(v2738)+4))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2743 = m.ExcPending
	if v2743 != 0 {
		goto L11
	} else {
		goto L544
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L11
	} else {
		goto L540
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2648 = m.ExcPending
	if v2648 != 0 {
		goto L11
	} else {
		goto L537
	}
L4:
	;
	m.G0 = v30 + int32(1632)
	return
L5:
	;
	F_ATAddCheckNNConstraint(m, l0, l1, l2, l3, l4, l5, int32(0), l6, l7)
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		goto L11
	} else {
		goto L536
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v375 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1624)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1616)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1608)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1600)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1592)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1584)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1576)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1568)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1560)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1552)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1544)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1536)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1528)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30+int32(1520)))) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1512)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v30)+1504)) = v375
	v434 = F__emscripten_memset_bulkmem(m, v30+int32(1376), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L79
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v44 = F_ConstraintNameIsUsed(m, int32(0), v43, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+1376)) = uint8(v73)
	if v71 == v73 {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	return
L12:
	;
	if v44 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v55 + int32(4)
	F_errmsg(m, int32(107256), v30+int32(384))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(464155), int32(9833), int32(85475))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	v337 = F_pstrdup(m, v30+int32(1376))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L11
	} else {
		goto L77
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v79 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v89 = int32(0)
	v92 = v9
	goto L21
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v92<<(uint(int32(2))%32))))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if int32(0) < v89 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L18
L23:
	;
	v121 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(1376)+v89))) = uint8(v121)
	v125 = v89 + int32(1)
	goto L25
L24:
	;
	v125 = v89
	goto L25
L25:
	;
	v128 = v30 + int32(1376) + v125
	goto L29
L26:
	;
	if v128&int32(3) == int32(0) {
		v267 = v128
		goto L60
	} else {
		goto L61
	}
L27:
	;
	v241 = F_strlen(m, v230)
	mBase = m.M
	goto L26
L29:
	;
	goto L30
L30:
	;
	v135 = int32(63)
	if (v128^v115)&int32(3) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231))) = uint8(v234)
	goto L27
L32:
	;
	v215 = v210
	v216 = v211
	v217 = v212
	goto L54
L33:
	;
	if v205 == int32(0) {
		v230 = v203
		v231 = v204
		goto L31
	} else {
		goto L53
	}
L34:
	;
	v203 = v115
	v204 = v128
	v205 = v135
	goto L33
L35:
	;
	goto L36
L36:
	;
	if v115&int32(3) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v172 == int32(0) {
		v230 = v169
		v231 = v170
		goto L31
	} else {
		goto L46
	}
L38:
	;
	v169 = v115
	v170 = v128
	v171 = v135
	v172 = int32(1)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v148 = v115
	v149 = v128
	v150 = v135
	goto L41
L41:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v152)
	if v152 == int32(0) {
		v210 = v148
		v211 = v149
		v212 = v150
		goto L32
	} else {
		goto L43
	}
L42:
	;
	v169 = v163
	v170 = v157
	v171 = v159
	v172 = v161
	goto L37
L43:
	;
	v156 = int32(1)
	v157 = v149 + v156
	v159 = v150 - v156
	v160 = int32(0)
	v161 = base.B2i32(v159 != v160)
	v163 = v148 + v156
	if v163&int32(3) == v160 {
		v169 = v163
		v170 = v157
		v171 = v159
		v172 = v161
		goto L37
	} else {
		goto L44
	}
L44:
	;
	if v159 != 0 {
		v148 = v163
		v149 = v157
		v150 = v159
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v175 == int32(0) {
		v203 = v169
		v204 = v170
		v205 = v171
		goto L33
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(v171) < base.Ui32(int32(4)) {
		v203 = v169
		v204 = v170
		v205 = v171
		goto L33
	} else {
		goto L48
	}
L48:
	;
	v181 = v169
	v182 = v170
	v183 = v171
	goto L49
L49:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v189 = int32(-2139062144)
	if (int32(16843008)-v186|v186)&v189 != v189 {
		v210 = v181
		v211 = v182
		v212 = v183
		goto L32
	} else {
		goto L51
	}
L50:
	;
	v203 = v197
	v204 = v195
	v205 = v199
	goto L33
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v186
	v194 = int32(4)
	v195 = v182 + v194
	v197 = v181 + v194
	v199 = v183 - v194
	if base.Ui32(int32(3)) < base.Ui32(v199) {
		v181 = v197
		v182 = v195
		v183 = v199
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v210 = v203
	v211 = v204
	v212 = v205
	goto L32
L54:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v219)
	if v219 == int32(0) {
		v230 = v215
		v231 = v216
		goto L31
	} else {
		goto L56
	}
L55:
	;
	v230 = v226
	v231 = v224
	goto L31
L56:
	;
	v223 = int32(1)
	v224 = v216 + v223
	v226 = v215 + v223
	v228 = v217 - v223
	if v228 != 0 {
		v215 = v226
		v216 = v224
		v217 = v228
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v301 = v300 + v125
	if int32(64) <= v301 {
		goto L18
	} else {
		goto L75
	}
L59:
	;
	v300 = v292 - v128
	goto L58
L60:
	;
	v271 = v267
	goto L69
L61:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v251 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v300 = int32(0)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v256 = v128
	goto L65
L65:
	;
	v260 = v256 + int32(1)
	if v260&int32(3) == int32(0) {
		v267 = v260
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v292 = v260
	goto L59
L67:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v265 != 0 {
		v256 = v260
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v280 = int32(-2139062144)
	if (int32(16843008)-v277|v277)&v280 == v280 {
		v271 = v271 + int32(4)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v286 = v271
	goto L72
L71:
	;
	goto L70
L72:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v290 != 0 {
		v286 = v286 + int32(1)
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v292 = v286
	goto L59
L74:
	;
	goto L73
L75:
	;
	v305 = v92 + int32(1)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v305 < v306 {
		v89 = v301
		v92 = v305
		goto L21
	} else {
		goto L76
	}
L76:
	;
	goto L22
L77:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+68))
	v343 = F_ChooseConstraintName(m, v72+int32(4), v337, int32(19535), v341, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v343
	goto L7
L79:
	;
	v440 = F__emscripten_memset_bulkmem(m, v30+int32(1248), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L80
L80:
	;
	v446 = F__emscripten_memset_bulkmem(m, v30+int32(1120), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L81
L81:
	;
	v452 = F__emscripten_memset_bulkmem(m, v30+int32(992), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L82
L82:
	;
	v458 = F__emscripten_memset_bulkmem(m, v30+int32(864), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L83
L83:
	;
	v464 = F__emscripten_memset_bulkmem(m, v30+int32(736), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L84
L84:
	;
	v470 = F__emscripten_memset_bulkmem(m, v30+int32(608), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L85
L85:
	;
	v476 = F__emscripten_memset_bulkmem(m, v30+int32(480), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L86
L86:
	;
	v477 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+472)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v30)+464)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v30)+456)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v30)+448)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v30)+440)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v30)+432)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v30)+424)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v30)+416)) = v477
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	if v493 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	v495 = v494
	goto L89
L88:
	;
	v495 = v9
	goto L89
L89:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l4)+100))
	if v496 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if l5 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v498 = F_table_open(m, v496, int32(6))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L11
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l4)+72))
	v502 = F_table_openrv(m, v500, int32(6))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L11
	} else {
		goto L95
	}
L94:
	;
	v504 = v498
	goto L90
L95:
	;
	v504 = v502
	goto L90
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L11
	} else {
		goto L532
	}
L97:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+119)))
	if v508 == int32(112) {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+119)))
	switch v512 - int32(112) {
	case 0, 2:
		goto L101
	default:
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
	if v537 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v522 + int32(4)
	F_errmsg(m, int32(371325), v30+int32(16))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(464155), int32(10119), int32(84984))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L11
	} else {
		goto L528
	}
L108:
	;
	v541 = int32(1)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v504)+56))
	if base.Ui32(v542) < base.Ui32(int32(12000)) {
		v551 = v541
		goto L112
	} else {
		goto L113
	}
L109:
	;
	goto L110
L110:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+118)))
	switch v553 - int32(112) {
	case 0:
		goto L121
	default:
		goto L118
	case 4:
		goto L119
	case 5:
		goto L120
	}
L111:
	;
	if v551 != 0 {
		goto L107
	} else {
		goto L115
	}
L112:
	;
	goto L111
L113:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+68))
	if v546 == int32(99) {
		v551 = v541
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v549 = F_isTempToastNamespace(m, v546)
	mBase = m.M
	v551 = v549
	goto L112
L115:
	;
	goto L110
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L11
	} else {
		goto L524
	}
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L11
	} else {
		goto L520
	}
L118:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v614 = F_transformColumnNameList(m, v606, v607, v30+int32(1504), v30+int32(1248), v30+int32(992))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L11
	} else {
		goto L135
	}
L119:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+118)))
	if v597 != int32(116) {
		goto L116
	} else {
		goto L132
	}
L120:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+118)))
	switch v577 - int32(112) {
	case 0, 5:
		goto L118
	default:
		goto L127
	}
L121:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+118)))
	if v557 == int32(112) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(153877), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(464155), int32(10141), int32(84984))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L11
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(155048), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(464155), int32(10148), int32(84984))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+24)))
	if v600 != int32(1) {
		goto L117
	} else {
		goto L133
	}
L133:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+24)))
	if v603 == int32(0) {
		goto L117
	} else {
		goto L134
	}
L134:
	;
	goto L118
L135:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v616 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L11
	} else {
		goto L516
	}
L137:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+85)))
	if v619 == int32(1) {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v626 = int32(0)
	v628 = F_transformColumnNameList(m, v622, v623, v30+int32(416), v626, v626)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L11
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	if v628 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v631 = int32(0)
	if v614 == v631 {
		v2712 = v631
		goto L1
	} else {
		goto L145
	}
L143:
	;
	v834 = v9
	goto L144
L144:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	if v837 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L145:
	;
	v639 = v631
	v658 = v9
	goto L146
L146:
	;
	v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(416)+v639<<(uint(int32(1))%32)))))
	v674 = int32(0)
	goto L148
L147:
	;
	v834 = v804
	goto L144
L148:
	;
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1504)+v674<<(uint(int32(1))%32)))))
	if v667 != v700 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v705 = int32(0)
	if v705 < v658 {
		goto L155
	} else {
		goto L156
	}
L150:
	;
	v703 = v674 + int32(1)
	if v614 != v703 {
		v674 = v703
		goto L148
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	goto L149
L153:
	;
	v2712 = v639
	goto L1
L154:
	;
	v808 = v639 + int32(1)
	if v808 != v628 {
		v639 = v808
		v658 = v804
		goto L146
	} else {
		goto L162
	}
L155:
	;
	v714 = v705
	goto L158
L156:
	;
	goto L157
L157:
	;
	v774 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(416)+v658<<(uint(v774)%32)))) = uint16(v667)
	v804 = v658 + v774
	goto L154
L158:
	;
	v740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(416)+v714<<(uint(int32(1))%32)))))
	if v740 == v667 {
		v804 = v658
		goto L154
	} else {
		goto L160
	}
L159:
	;
	goto L157
L160:
	;
	v743 = v714 + int32(1)
	if v743 != v658 {
		v714 = v743
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	goto L147
L163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		goto L11
	} else {
		goto L513
	}
L164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L11
	} else {
		goto L509
	}
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L11
	} else {
		goto L506
	}
L166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2458 = m.ExcPending
	if v2458 != 0 {
		goto L11
	} else {
		goto L503
	}
L167:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L11
	} else {
		goto L496
	}
L168:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2401 = m.ExcPending
	if v2401 != 0 {
		goto L11
	} else {
		goto L493
	}
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L11
	} else {
		goto L486
	}
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L11
	} else {
		goto L483
	}
L171:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L11
	} else {
		goto L473
	}
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L11
	} else {
		goto L470
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L11
	} else {
		goto L466
	}
L174:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L11
	} else {
		goto L462
	}
L175:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L11
	} else {
		goto L458
	}
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L11
	} else {
		goto L454
	}
L177:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L11
	} else {
		goto L450
	}
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L11
	} else {
		goto L446
	}
L179:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L11
	} else {
		goto L442
	}
L180:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L11
	} else {
		goto L439
	}
L181:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v504)+56))
	v1574 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v1576 = F_pg_class_aclcheck(m, v1572, v1574, int64(32))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L11
	} else {
		goto L300
	}
L182:
	;
	if v1540&int32(1) != 0 {
		v1556 = v1524
		v1564 = v1532
		goto L181
	} else {
		goto L298
	}
L183:
	;
	v840 = F_RelationGetIndexList(m, v504)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L11
	} else {
		goto L188
	}
L184:
	;
	goto L185
L185:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v504)+56))
	v1095 = F_transformColumnNameList(m, v1088, v837, v30+int32(1568), v30+int32(1376), v30+int32(1120))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L11
	} else {
		goto L228
	}
L186:
	;
	F_list_free(m, v840)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L11
	} else {
		goto L206
	}
L187:
	;
	F_list_free(m, v840)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L11
	} else {
		goto L205
	}
L188:
	;
	if v840 == int32(0) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if v844 <= int32(0) {
		goto L187
	} else {
		goto L190
	}
L190:
	;
	v857 = int32(0)
	goto L191
L191:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v876+v857<<(uint(int32(2))%32))))
	v881 = F_SearchSysCache1(m, int32(34), v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L11
	} else {
		goto L193
	}
L192:
	;
	goto L187
L193:
	;
	if v881 == int32(0) {
		goto L163
	} else {
		goto L194
	}
L194:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v881)+16))
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+22)))
	v887 = v885 + v886
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+14)))
	if v888 != int32(1) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	F_ReleaseCatCache(m, v881)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L11
	} else {
		goto L203
	}
L196:
	;
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+18)))
	if v891 != int32(1) {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+16)))
	if v894 != 0 {
		goto L186
	} else {
		goto L198
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L11
	} else {
		goto L199
	}
L199:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L11
	} else {
		goto L200
	}
L200:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v902 + int32(4)
	F_errmsg(m, int32(670691), v30+int32(272))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L11
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(464155), int32(13423), int32(20568))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L11
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v919 = v857 + int32(1)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if v919 < v920 {
		v857 = v919
		goto L191
	} else {
		goto L204
	}
L204:
	;
	goto L192
L205:
	;
	goto L2
L206:
	;
	if v880 == int32(0) {
		goto L2
	} else {
		goto L207
	}
L207:
	;
	v955 = int32(0)
	v958 = F_SysCacheGetAttrNotNull(m, int32(34), v881, int32(18))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L11
	} else {
		goto L208
	}
L208:
	;
	v960 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v960
	v962 = int32(*(*int16)(unsafe.Add(mBase, uint32(v887)+10)))
	if v960 < v962 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v980 = v955
	goto L212
L210:
	;
	v1049 = v955
	goto L211
L211:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+15)))
	F_ReleaseCatCache(m, v881)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L11
	} else {
		goto L221
	}
L212:
	;
	v997 = v980 << (uint(int32(1)) % 32)
	v1002 = int32(*(*int16)(unsafe.Add(mBase, uint32(v997+(v887+int32(48))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v997+(v30+int32(1568))))) = uint16(v1002)
	v1005 = v980 << (uint(int32(2)) % 32)
	v1009 = F_attnumTypeId(m, v504, v1002)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L11
	} else {
		goto L214
	}
L213:
	;
	v1049 = v1035
	goto L211
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1005+(v30+int32(1376))))) = v1009
	v1015 = F_attnumCollationId(m, v504, v1002)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L11
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(1120)+v1005))) = v1015
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1005+(v958+int32(24)))))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(864)+v1005))) = v1022
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v1025 = F_attnumAttName(m, v504, v1002)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L11
	} else {
		goto L216
	}
L216:
	;
	v1027 = F_pstrdup(m, v1025)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L11
	} else {
		goto L217
	}
L217:
	;
	v1029 = F_makeString(m, v1027)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L11
	} else {
		goto L218
	}
L218:
	;
	v1031 = F_lappend(m, v1024, v1029)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L11
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v1031
	v1035 = v980 + int32(1)
	v1036 = int32(*(*int16)(unsafe.Add(mBase, uint32(v887)+10)))
	if v1035 < v1036 {
		v980 = v1035
		goto L212
	} else {
		goto L220
	}
L220:
	;
	goto L213
L221:
	;
	if v1065 != int32(1) {
		v1556 = v1049
		v1564 = v880
		goto L181
	} else {
		goto L222
	}
L222:
	;
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v1071 != 0 {
		v1524 = v1049
		v1532 = v880
		v1540 = int32(0)
		goto L182
	} else {
		goto L223
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L11
	} else {
		goto L224
	}
L224:
	;
	F_errcode(m, int32(819332))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L11
	} else {
		goto L225
	}
L225:
	;
	F_errmsg(m, int32(370301), int32(0))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L11
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(464155), int32(10201), int32(84984))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L11
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	if v616 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+85)))
	if v1097 == int32(0) {
		goto L164
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	if v1095 != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	goto L231
L233:
	;
	v1106 = int32(0)
	goto L236
L234:
	;
	goto L235
L235:
	;
	v1245 = F_RelationGetIndexList(m, v504)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L11
	} else {
		goto L253
	}
L236:
	;
	v1129 = v1106 + int32(1)
	if base.Ui32(v1095) <= base.Ui32(v1129) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	goto L235
L238:
	;
	if v1129 != v1095 {
		v1106 = v1129
		goto L236
	} else {
		goto L250
	}
L239:
	;
	v1136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1106<<(uint(int32(1))%32)))))
	v1143 = v1129
	goto L240
L240:
	;
	v1169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1143<<(uint(int32(1))%32)))))
	if v1169 != v1136 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L11
	} else {
		goto L246
	}
L242:
	;
	v1172 = v1143 + int32(1)
	if v1095 != v1172 {
		v1143 = v1172
		goto L240
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	goto L241
L245:
	;
	goto L238
L246:
	;
	F_errcode(m, int32(819332))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L11
	} else {
		goto L247
	}
L247:
	;
	F_errmsg(m, int32(149225), int32(0))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L11
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(464155), int32(13512), int32(120611))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L11
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	goto L237
L251:
	;
	v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302)+15)))
	F_ReleaseCatCache(m, v1296)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L11
	} else {
		goto L296
	}
L252:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L11
	} else {
		goto L292
	}
L253:
	;
	if v1245 == int32(0) {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	if v1249 <= int32(0) {
		goto L252
	} else {
		goto L255
	}
L255:
	;
	v1252 = int32(0)
	v1256 = int32(1)
	v1259 = (v1095 - v1256) << (uint(v1256) % 32)
	v1272 = v1252
	v1285 = v9
	goto L256
L256:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+12))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1291+v1272<<(uint(int32(2))%32))))
	v1296 = F_SearchSysCache1(m, int32(34), v1295)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L11
	} else {
		goto L258
	}
L257:
	;
	if v1445&int32(1) != 0 {
		goto L179
	} else {
		goto L291
	}
L258:
	;
	if v1296 == int32(0) {
		goto L180
	} else {
		goto L259
	}
L259:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+16))
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300)+22)))
	v1302 = v1300 + v1301
	v1303 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1302)+10)))
	if v1095 != v1303 {
		v1445 = v1285
		goto L260
	} else {
		goto L261
	}
L260:
	;
	F_ReleaseCatCache(m, v1296)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L11
	} else {
		goto L289
	}
L261:
	;
	if v616 != 0 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302)+18)))
	if v1309 != int32(1) {
		v1445 = v1285
		goto L260
	} else {
		goto L268
	}
L263:
	;
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302)+15)))
	if v1305 != 0 {
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302)+12)))
	if v1306 != int32(1) {
		v1445 = v1285
		goto L260
	} else {
		goto L267
	}
L266:
	;
	v1445 = v1285
	goto L260
L267:
	;
	goto L262
L268:
	;
	v1314 = F_heap_attisnull(m, v1296, int32(21), int32(0))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L11
	} else {
		goto L269
	}
L269:
	;
	if v1314 == int32(0) {
		v1445 = v1285
		goto L260
	} else {
		goto L270
	}
L270:
	;
	v1320 = F_heap_attisnull(m, v1296, int32(20), int32(0))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L11
	} else {
		goto L271
	}
L271:
	;
	if v1320 == int32(0) {
		v1445 = v1285
		goto L260
	} else {
		goto L272
	}
L272:
	;
	v1326 = F_SysCacheGetAttrNotNull(m, int32(34), v1296, int32(18))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L11
	} else {
		goto L273
	}
L273:
	;
	if v1095 == int32(0) {
		v1445 = v1285
		goto L260
	} else {
		goto L274
	}
L274:
	;
	v1333 = v1302 + int32(48)
	v1353 = int32(0)
	goto L275
L275:
	;
	v1367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1353<<(uint(int32(1))%32)))))
	v1379 = int32(0)
	goto L277
L276:
	;
	if base.B2i32(v1095 != v1252)&v616 != 0 {
		goto L284
	} else {
		goto L285
	}
L277:
	;
	v1399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1333+v1379<<(uint(int32(1))%32)))))
	if v1399 != v1367 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1406 = int32(2)
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1326+int32(24)+v1379<<(uint(v1406)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(864)+v1353<<(uint(v1406)%32)))) = v1412
	v1415 = v1353 + int32(1)
	if v1415 != v1095 {
		v1353 = v1415
		goto L275
	} else {
		goto L283
	}
L279:
	;
	v1402 = v1379 + int32(1)
	if v1402 != v1095 {
		v1379 = v1402
		goto L277
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	goto L278
L282:
	;
	v1445 = v1285
	goto L260
L283:
	;
	goto L276
L284:
	;
	v1417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1259+(v30+int32(1568))))))
	v1419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1333+v1259))))
	if v1417 != v1419 {
		v1445 = v1285
		goto L260
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302)+16)))
	if v1422 != 0 {
		goto L251
	} else {
		goto L288
	}
L287:
	;
	goto L286
L288:
	;
	v1445 = int32(1)
	goto L260
L289:
	;
	v1453 = v1272 + int32(1)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	if v1453 < v1454 {
		v1272 = v1453
		v1285 = v1445
		goto L256
	} else {
		goto L290
	}
L290:
	;
	goto L257
L291:
	;
	goto L252
L292:
	;
	F_errcode(m, int32(819332))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L11
	} else {
		goto L293
	}
L293:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+288)) = v1492 + int32(4)
	F_errmsg(m, int32(670821), v30+int32(288))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L11
	} else {
		goto L294
	}
L294:
	;
	F_errfinish(m, int32(464155), int32(13621), int32(120611))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L11
	} else {
		goto L295
	}
L295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L296:
	;
	F_list_free(m, v1245)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L11
	} else {
		goto L297
	}
L297:
	;
	v1524 = v1095
	v1532 = v1295
	v1540 = v1506 ^ int32(1)
	goto L182
L298:
	;
	if v616 == int32(0) {
		goto L178
	} else {
		goto L299
	}
L299:
	;
	v1556 = v1524
	v1564 = v1532
	goto L181
L300:
	;
	if v1556 <= int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	if v614 != 0 {
		goto L319
	} else {
		goto L320
	}
L302:
	;
	if v1576 == int32(0) {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1589 = int32(0)
	goto L304
L304:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v504)+56))
	v1616 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(1568)+v1589<<(uint(int32(1))%32)))))
	v1618 = F_pg_attribute_aclcheck(m, v1610, v1616, v1574, int64(32))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L11
	} else {
		goto L306
	}
L305:
	;
	goto L301
L306:
	;
	if v1618 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	v1621 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1620)+119)))
	switch v1621 - int32(73) {
	case 0, 32:
		goto L316
	default:
		v1631 = int32(41)
		goto L311
	case 10:
		goto L315
	case 29:
		goto L312
	case 36:
		goto L313
	case 45:
		goto L314
	}
L308:
	;
	goto L309
L309:
	;
	v1640 = v1589 + int32(1)
	if v1640 != v1556 {
		v1589 = v1640
		goto L304
	} else {
		goto L318
	}
L310:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	F_aclcheck_error(m, v1618, v1633, v1634+int32(4))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L11
	} else {
		goto L317
	}
L311:
	;
	v1633 = v1631
	goto L310
L312:
	;
	v1631 = int32(18)
	goto L311
L313:
	;
	v1633 = int32(23)
	goto L310
L314:
	;
	v1633 = int32(51)
	goto L310
L315:
	;
	v1633 = int32(37)
	goto L310
L316:
	;
	v1633 = int32(20)
	goto L310
L317:
	;
	goto L309
L318:
	;
	goto L305
L319:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1669)))
	v1683 = int32(0)
	goto L322
L320:
	;
	goto L321
L321:
	;
	v1780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+84)))
	if v1780 != int32(1) {
		goto L339
	} else {
		goto L340
	}
L322:
	;
	v1709 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(1504)+v1683<<(uint(int32(1))%32)))))
	v1713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669+v1670<<(uint(int32(4))%32)+int32(10)+v1709*int32(100)))))
	if v1713 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	goto L321
L324:
	;
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+87)))
	v1716 = v1714 - int32(99)
	if int32(1)<<(uint(v1716)%32)&int32(2051) != 0 {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	goto L326
L326:
	;
	v1751 = v1683 + int32(1)
	if v1751 != v614 {
		v1683 = v1751
		goto L322
	} else {
		goto L338
	}
L327:
	;
	v1724 = base.B2i32(base.Ui32(v1716) <= base.Ui32(int32(11)))
	goto L329
L328:
	;
	v1724 = int32(0)
	goto L329
L329:
	;
	if v1724 != 0 {
		goto L177
	} else {
		goto L330
	}
L330:
	;
	v1725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+88)))
	switch v1725 - int32(100) {
	case 0, 10:
		goto L332
	default:
		goto L331
	}
L331:
	;
	if v1713 == int32(118) {
		goto L176
	} else {
		goto L337
	}
L332:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L11
	} else {
		goto L333
	}
L333:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L11
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+240)) = int32(506160)
	F_errmsg(m, int32(257706), v30+int32(240))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L11
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(464155), int32(10258), int32(84984))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L11
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	goto L326
L338:
	;
	goto L323
L339:
	;
	if v1556 != v614 {
		goto L173
	} else {
		goto L347
	}
L340:
	;
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+87)))
	v1785 = v1783 - int32(99)
	if int32(1)<<(uint(v1785)%32)&int32(34819) != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1793 = base.B2i32(base.Ui32(v1785) <= base.Ui32(int32(15)))
	goto L343
L342:
	;
	v1793 = int32(0)
	goto L343
L343:
	;
	if v1793 != 0 {
		goto L175
	} else {
		goto L344
	}
L344:
	;
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+88)))
	v1796 = v1794 - int32(99)
	if base.Ui32(int32(15)) < base.Ui32(v1796) {
		goto L339
	} else {
		goto L345
	}
L345:
	;
	if int32(1)<<(uint(v1796)%32)&int32(34819) != 0 {
		goto L174
	} else {
		goto L346
	}
L346:
	;
	goto L339
L347:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	v1807 = base.B2i32(v1805 != int32(0))
	if v614 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1822 = int32(0)
	v1826 = v495
	v1834 = v1807
	goto L351
L349:
	;
	v2074 = v1807
	goto L350
L350:
	;
	if v616 != 0 {
		goto L431
	} else {
		goto L432
	}
L351:
	;
	v1839 = v1822 << (uint(int32(2)) % 32)
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1839+(v30+int32(992)))))
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(1120)+v1839)))
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(1248)+v1839)))
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(1376)+v1839)))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(864)+v1839)))
	v1861 = F_SearchSysCache1(m, int32(14), v1860)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L11
	} else {
		goto L353
	}
L352:
	;
	v2074 = v2034
	goto L350
L353:
	;
	if v1861 == int32(0) {
		goto L172
	} else {
		goto L354
	}
L354:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1861)+16))
	v1866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865)+22)))
	v1867 = v1865 + v1866
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v1867)+84))
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1867)+80))
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1867)+4))
	F_ReleaseCatCache(m, v1861)
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L11
	} else {
		goto L355
	}
L355:
	;
	v1876 = v616 & base.B2i32(v1822 == v614-int32(1))
	if v1876 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1877 = int32(7)
	goto L358
L357:
	;
	v1877 = int32(3)
	goto L358
L358:
	;
	v1879 = F_IndexAmTranslateCompareType(m, v1877, v1870, v1869, int32(1))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L11
	} else {
		goto L359
	}
L359:
	;
	if v1879 == int32(0) {
		goto L171
	} else {
		goto L360
	}
L360:
	;
	v1883 = base.I32_extend16_s(v1879)
	v1884 = F_get_opfamily_member(m, v1869, v1868, v1868, v1883)
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L11
	} else {
		goto L361
	}
L361:
	;
	if v1884 == int32(0) {
		goto L170
	} else {
		goto L362
	}
L362:
	;
	v1888 = F_getBaseType(m, v1851)
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L11
	} else {
		goto L365
	}
L363:
	;
	if v1916 == int32(0) {
		goto L169
	} else {
		goto L382
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v1851
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v1855
	*(*int32)(unsafe.Add(mBase, uint32(v30)+412)) = v1868
	*(*int32)(unsafe.Add(mBase, uint32(v30)+408)) = v1868
	v1909 = F_can_coerce_type(m, int32(2), v30+int32(392), v30+int32(408), int32(0))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L11
	} else {
		goto L372
	}
L365:
	;
	v1890 = F_get_opfamily_member(m, v1869, v1868, v1888, v1883)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L11
	} else {
		goto L366
	}
L366:
	;
	if v1890 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1897 = int32(0)
	goto L364
L368:
	;
	goto L369
L369:
	;
	v1895 = F_get_opfamily_member(m, v1869, v1888, v1888, v1883)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L11
	} else {
		goto L370
	}
L370:
	;
	if v1895 != 0 {
		v1915 = v1888
		v1916 = v1890
		v1917 = v1895
		goto L363
	} else {
		goto L371
	}
L371:
	;
	v1897 = v1888
	goto L364
L372:
	;
	if v1909 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1911 = v1884
	goto L375
L374:
	;
	v1911 = v1890
	goto L375
L375:
	;
	if v1909 != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1913 = v1884
	goto L378
L377:
	;
	v1913 = int32(0)
	goto L378
L378:
	;
	if v1909 != 0 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1914 = v1868
	goto L381
L380:
	;
	v1914 = v1897
	goto L381
L381:
	;
	v1915 = v1914
	v1916 = v1911
	v1917 = v1913
	goto L363
L382:
	;
	if v1917 == int32(0) {
		goto L169
	} else {
		goto L383
	}
L383:
	;
	v1923 = int32(0)
	if base.B2i32(v1847 != v1923) != base.B2i32(v1843 != v1923) {
		goto L168
	} else {
		goto L384
	}
L384:
	;
	if v1847 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1940 = int32(0)
	if v1834&int32(1) == v1940 {
		goto L393
	} else {
		goto L394
	}
L386:
	;
	if v1843 == int32(0) {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1932 = F_get_collation_isdeterministic(m, v1847)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L11
	} else {
		goto L388
	}
L388:
	;
	v1934 = F_get_collation_isdeterministic(m, v1843)
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L11
	} else {
		goto L389
	}
L389:
	;
	if v1932&v1934 != 0 {
		goto L385
	} else {
		goto L390
	}
L390:
	;
	if v1847 != v1843 {
		goto L167
	} else {
		goto L391
	}
L391:
	;
	goto L385
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(608)+v1839))) = v1884
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(736)+v1839))) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v30+int32(480)+v1839))) = v1917
	v2049 = v1822 + int32(1)
	if v2049 != v614 {
		v1822 = v2049
		v1826 = v2032
		v1834 = v2034
		goto L351
	} else {
		goto L430
	}
L393:
	;
	v2032 = v1826
	v2034 = v1940
	goto L392
L394:
	;
	goto L395
L395:
	;
	v1944 = v1826 + int32(4)
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l4)+96))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+12))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1946)+4))
	if base.Ui32(v1944) < base.Ui32(v1947+v1948<<(uint(int32(2))%32)) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1953 = v1944
	goto L398
L397:
	;
	v1953 = int32(0)
	goto L398
L398:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1826)))
	if v1916 != v1954 {
		v2032 = v1953
		v2034 = v1940
		goto L392
	} else {
		goto L399
	}
L399:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1956)))
	v1966 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(1504)+v1822<<(uint(int32(1))%32)))))
	v1971 = v1956 + v1957<<(uint(int32(4))%32) + v1966*int32(100) - int32(80)
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+68))
	if v1972 == v1915 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	if v1915 == v1851 {
		goto L407
	} else {
		goto L408
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = int32(0)
	v1984 = int32(2)
	goto L400
L402:
	;
	goto L403
L403:
	;
	v1980 = F_find_coercion_pathway(m, v1915, v1972, int32(0), v30+int32(392))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L11
	} else {
		goto L404
	}
L404:
	;
	if v1980 == int32(0) {
		goto L166
	} else {
		goto L405
	}
L405:
	;
	v1984 = v1980
	goto L400
L406:
	;
	if v1984 != v1996 {
		v2032 = v1953
		v2034 = v1940
		goto L392
	} else {
		goto L412
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+408)) = int32(0)
	v1996 = int32(2)
	goto L406
L408:
	;
	goto L409
L409:
	;
	v1992 = F_find_coercion_pathway(m, v1915, v1851, int32(0), v30+int32(408))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L11
	} else {
		goto L410
	}
L410:
	;
	if v1992 == int32(0) {
		goto L165
	} else {
		goto L411
	}
L411:
	;
	v1996 = v1992
	goto L406
L412:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v30)+408))
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v30)+392))
	if v1998 != v1999 {
		v2032 = v1953
		v2034 = v1940
		goto L392
	} else {
		goto L413
	}
L413:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+96))
	if v1915 <= int32(3830) {
		goto L416
	} else {
		goto L417
	}
L414:
	;
	if v2001 == v1843 {
		v2032 = v1953
		v2034 = int32(1)
		goto L392
	} else {
		goto L426
	}
L415:
	;
	if v1851 != v1972 {
		v2032 = v1953
		v2034 = v1940
		goto L392
	} else {
		goto L425
	}
L416:
	;
	switch v1915 - int32(2277) {
	case 0, 6:
		goto L415
	case 1, 2, 3, 4, 5:
		goto L414
	default:
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	if base.Ui32(v1915-int32(5077)) < base.Ui32(int32(4)) {
		goto L415
	} else {
		goto L422
	}
L419:
	;
	if v1915 == int32(2776) {
		goto L415
	} else {
		goto L420
	}
L420:
	;
	if v1915 == int32(3500) {
		goto L415
	} else {
		goto L421
	}
L421:
	;
	goto L414
L422:
	;
	if base.Ui32(v1915-int32(4537)) < base.Ui32(int32(2)) {
		goto L415
	} else {
		goto L423
	}
L423:
	;
	if v1915 != int32(3831) {
		goto L414
	} else {
		goto L424
	}
L424:
	;
	goto L415
L425:
	;
	goto L414
L426:
	;
	v2024 = F_get_collation_isdeterministic(m, v2001)
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L11
	} else {
		goto L427
	}
L427:
	;
	if v2024 == int32(0) {
		v2032 = v1953
		v2034 = int32(0)
		goto L392
	} else {
		goto L428
	}
L428:
	;
	v2028 = F_get_collation_isdeterministic(m, v1843)
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L11
	} else {
		goto L429
	}
L429:
	;
	v2032 = v1953
	v2034 = v2028
	goto L392
L430:
	;
	goto L352
L431:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v614<<(uint(int32(2))%32)+v30)+860))
	F_FindFKPeriodOpers(m, v2081, v30+int32(392), v30+int32(408), v30+int32(404))
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L11
	} else {
		goto L434
	}
L432:
	;
	goto L433
L433:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v2094 = int32(0)
	F_addFkConstraint(m, v30+int32(392), int32(2), v2093, l4, l3, v504, v1564, v2094, v614, v30+int32(1568), v30+int32(1504), v30+int32(736), v30+int32(608), v30+int32(480), v834, v30+int32(416), v2094, v616)
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L11
	} else {
		goto L435
	}
L434:
	;
	goto L433
L435:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v30)+396))
	v2123 = int32(0)
	F_addFkRecurseReferenced(m, l4, l3, v504, v1564, v2110, v614, v30+int32(1568), v30+int32(1504), v30+int32(736), v30+int32(608), v30+int32(480), v834, v30+int32(416), v2123, v2123, v616)
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L11
	} else {
		goto L436
	}
L436:
	;
	v2139 = int32(0)
	F_addFkRecurseReferencing(m, l1, l4, l3, v504, v1564, v2110, v614, v30+int32(1568), v30+int32(1504), v30+int32(736), v30+int32(608), v30+int32(480), v834, v30+int32(416), v2074, l7, v2139, v2139, v616)
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L11
	} else {
		goto L437
	}
L437:
	;
	F_sequence_close(m, v504, int32(0))
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L11
	} else {
		goto L438
	}
L438:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v30)+400))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2146
	v2148 = *(*int64)(unsafe.Add(mBase, uint32(v30)+392))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v2148
	goto L4
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+304)) = v1295
	F_errmsg_internal(m, int32(37172), v30+int32(304))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L11
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(464155), int32(13531), int32(120611))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L11
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
	F_errcode(m, int32(325))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L11
	} else {
		goto L443
	}
L443:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+320)) = v2172 + int32(4)
	F_errmsg(m, int32(670753), v30+int32(320))
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L11
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(464155), int32(13616), int32(120611))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L11
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	F_errcode(m, int32(819332))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L11
	} else {
		goto L447
	}
L447:
	;
	F_errmsg(m, int32(491288), int32(0))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L11
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(464155), int32(10227), int32(84984))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L11
	} else {
		goto L449
	}
L449:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L450:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L11
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = int32(506875)
	F_errmsg(m, int32(257706), v30+int32(256))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L11
	} else {
		goto L452
	}
L452:
	;
	F_errfinish(m, int32(464155), int32(10252), int32(84984))
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L11
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L11
	} else {
		goto L455
	}
L455:
	;
	F_errmsg(m, int32(416905), int32(0))
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L11
	} else {
		goto L456
	}
L456:
	;
	F_errfinish(m, int32(464155), int32(10272), int32(84984))
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L11
	} else {
		goto L457
	}
L457:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L458:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L11
	} else {
		goto L459
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+224)) = int32(506875)
	F_errmsg(m, int32(510445), v30+int32(224))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L11
	} else {
		goto L460
	}
L460:
	;
	F_errfinish(m, int32(464155), int32(10287), int32(84984))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L11
	} else {
		goto L461
	}
L461:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L462:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L11
	} else {
		goto L463
	}
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+208)) = int32(506160)
	F_errmsg(m, int32(510445), v30+int32(208))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L11
	} else {
		goto L464
	}
L464:
	;
	F_errfinish(m, int32(464155), int32(10296), int32(84984))
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L11
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
	F_errcode(m, int32(819332))
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L11
	} else {
		goto L467
	}
L467:
	;
	F_errmsg(m, int32(385914), int32(0))
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L11
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(464155), int32(10310), int32(84984))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L11
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v1860
	F_errmsg_internal(m, int32(39335), v30-int32(-64))
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L11
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(464155), int32(10342), int32(84984))
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L11
	} else {
		goto L472
	}
L472:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L473:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L11
	} else {
		goto L474
	}
L474:
	;
	if v1876 != 0 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v2315 = int32(20057)
	goto L477
L476:
	;
	v2315 = int32(20001)
	goto L477
L477:
	;
	F_errmsg(m, v2315, int32(0))
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L11
	} else {
		goto L478
	}
L478:
	;
	v2319 = F_get_opfamily_name(m, v1869)
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L11
	} else {
		goto L479
	}
L479:
	;
	v2321 = F_get_am_name(m, v1870)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L11
	} else {
		goto L480
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+88)) = v2321
	*(*int32)(unsafe.Add(mBase, uint32(v30)+84)) = v2319
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v1877
	F_errdetail(m, int32(617868), v30+int32(80))
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L11
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(464155), int32(10369), int32(84984))
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L11
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+108)) = v1869
	*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = v1868
	*(*int32)(unsafe.Add(mBase, uint32(v30)+100)) = v1868
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v1883
	F_errmsg_internal(m, int32(36732), v30+int32(96))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L11
	} else {
		goto L484
	}
L484:
	;
	F_errfinish(m, int32(464155), int32(10380), int32(84984))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L11
	} else {
		goto L485
	}
L485:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L486:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L11
	} else {
		goto L487
	}
L487:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+192)) = v2361
	F_errmsg(m, int32(419808), v30+int32(192))
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L11
	} else {
		goto L488
	}
L488:
	;
	v2369 = v1822 << (uint(int32(2)) % 32)
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+12))
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2369+v2371)))
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(v2373)+4))
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2375)+12))
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v2376+v2369)))
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2378)+4))
	v2380 = F_format_type_be(m, v1851)
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L11
	} else {
		goto L489
	}
L489:
	;
	v2382 = F_format_type_be(m, v1855)
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L11
	} else {
		goto L490
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+188)) = v2382
	*(*int32)(unsafe.Add(mBase, uint32(v30)+184)) = v2380
	*(*int32)(unsafe.Add(mBase, uint32(v30)+180)) = v2379
	*(*int32)(unsafe.Add(mBase, uint32(v30)+176)) = v2374
	F_errdetail(m, int32(557190), v30+int32(176))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L11
	} else {
		goto L491
	}
L491:
	;
	F_errfinish(m, int32(464155), int32(10439), int32(84984))
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L11
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
	F_errmsg_internal(m, int32(368128), int32(0))
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L11
	} else {
		goto L494
	}
L494:
	;
	F_errfinish(m, int32(464155), int32(10446), int32(84984))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L11
	} else {
		goto L495
	}
L495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L496:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L11
	} else {
		goto L497
	}
L497:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v2418
	F_errmsg(m, int32(419808), v30+int32(160))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L11
	} else {
		goto L498
	}
L498:
	;
	v2426 = v1822 << (uint(int32(2)) % 32)
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(v2427)+12))
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v2426+v2428)))
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2430)+4))
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v2432)+12))
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2433+v2426)))
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2435)+4))
	v2437 = F_get_collation_name(m, v1843)
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L11
	} else {
		goto L499
	}
L499:
	;
	v2439 = F_get_collation_name(m, v1847)
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L11
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+156)) = v2439
	*(*int32)(unsafe.Add(mBase, uint32(v30)+152)) = v2437
	*(*int32)(unsafe.Add(mBase, uint32(v30)+148)) = v2436
	*(*int32)(unsafe.Add(mBase, uint32(v30)+144)) = v2431
	F_errdetail(m, int32(586272), v30+int32(144))
	mBase = m.M
	v2449 = m.ExcPending
	if v2449 != 0 {
		goto L11
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(464155), int32(10473), int32(84984))
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L11
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+132)) = v1915
	*(*int32)(unsafe.Add(mBase, uint32(v30)+128)) = v1972
	F_errmsg_internal(m, int32(41323), v30+int32(128))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L11
	} else {
		goto L504
	}
L504:
	;
	F_errfinish(m, int32(464155), int32(13652), int32(73850))
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		goto L11
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
	*(*int32)(unsafe.Add(mBase, uint32(v30)+116)) = v1915
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v1851
	F_errmsg_internal(m, int32(41323), v30+int32(112))
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L11
	} else {
		goto L507
	}
L507:
	;
	F_errfinish(m, int32(464155), int32(13652), int32(73850))
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L11
	} else {
		goto L508
	}
L508:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L509:
	;
	F_errcode(m, int32(819332))
	mBase = m.M
	v2493 = m.ExcPending
	if v2493 != 0 {
		goto L11
	} else {
		goto L510
	}
L510:
	;
	F_errmsg(m, int32(371247), int32(0))
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L11
	} else {
		goto L511
	}
L511:
	;
	F_errfinish(m, int32(464155), int32(10213), int32(84984))
	mBase = m.M
	v2502 = m.ExcPending
	if v2502 != 0 {
		goto L11
	} else {
		goto L512
	}
L512:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v880
	F_errmsg_internal(m, int32(37172), v30+int32(48))
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L11
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(464155), int32(13410), int32(20568))
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L11
	} else {
		goto L515
	}
L515:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L516:
	;
	F_errcode(m, int32(819332))
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L11
	} else {
		goto L517
	}
L517:
	;
	F_errmsg(m, int32(370301), int32(0))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L11
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(464155), int32(10173), int32(84984))
	mBase = m.M
	v2533 = m.ExcPending
	if v2533 != 0 {
		goto L11
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L11
	} else {
		goto L521
	}
L521:
	;
	F_errmsg(m, int32(252885), int32(0))
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L11
	} else {
		goto L522
	}
L522:
	;
	F_errfinish(m, int32(464155), int32(10158), int32(84984))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L11
	} else {
		goto L523
	}
L523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L524:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L11
	} else {
		goto L525
	}
L525:
	;
	F_errmsg(m, int32(153762), int32(0))
	mBase = m.M
	v2560 = m.ExcPending
	if v2560 != 0 {
		goto L11
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(464155), int32(10154), int32(84984))
	mBase = m.M
	v2565 = m.ExcPending
	if v2565 != 0 {
		goto L11
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L11
	} else {
		goto L529
	}
L529:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+352)) = v2573 + int32(4)
	F_errmsg(m, int32(306910), v30+int32(352))
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L11
	} else {
		goto L530
	}
L530:
	;
	F_errfinish(m, int32(464155), int32(10125), int32(84984))
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L11
	} else {
		goto L531
	}
L531:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L532:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L11
	} else {
		goto L533
	}
L533:
	;
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	v2596 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v2595 + v2596
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v2594 + v2596
	F_errmsg(m, int32(656563), v30+int32(368))
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L11
	} else {
		goto L534
	}
L534:
	;
	F_errfinish(m, int32(464155), int32(10112), int32(84984))
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L11
	} else {
		goto L535
	}
L535:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L536:
	;
	goto L4
L537:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v2649
	F_errmsg_internal(m, int32(455045), v30)
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L11
	} else {
		goto L538
	}
L538:
	;
	F_errfinish(m, int32(464155), int32(9851), int32(85475))
	mBase = m.M
	v2658 = m.ExcPending
	if v2658 != 0 {
		goto L11
	} else {
		goto L539
	}
L539:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L540:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L11
	} else {
		goto L541
	}
L541:
	;
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(v504)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v2693 + int32(4)
	F_errmsg(m, int32(670641), v30+int32(32))
	mBase = m.M
	v2701 = m.ExcPending
	if v2701 != 0 {
		goto L11
	} else {
		goto L542
	}
L542:
	;
	F_errfinish(m, int32(464155), int32(13440), int32(20568))
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L11
	} else {
		goto L543
	}
L543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L544:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v2746 = m.ExcPending
	if v2746 != 0 {
		goto L11
	} else {
		goto L545
	}
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+336)) = v2739
	F_errmsg(m, int32(20113), v30+int32(336))
	mBase = m.M
	v2752 = m.ExcPending
	if v2752 != 0 {
		goto L11
	} else {
		goto L546
	}
L546:
	;
	F_errfinish(m, int32(464155), int32(10673), int32(138074))
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L11
	} else {
		goto L547
	}
L547:
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int64
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+119)))
	if base.B2i32(l5 == v8)&base.B2i32(v21 == int32(112)) == v8 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L35
	} else {
		goto L91
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L35
	} else {
		goto L87
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L35
	} else {
		goto L83
	}
L4:
	;
	if l6 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L35
	} else {
		goto L78
	}
L7:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+131)))
	if v29&int32(1) != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l3 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L9
L11:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+22)))
	v166 = v164 + v165
	v167 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166)+74)))
	if v167 <= int32(0) {
		goto L3
	} else {
		goto L51
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L35
	} else {
		goto L48
	}
L13:
	;
	v123 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L35
	} else {
		goto L41
	}
L14:
	;
	v115 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v36 <= v35 {
		v115 = v35
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v39 = int32(0)
	if v39 < v36 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v42 = v36
	goto L20
L19:
	;
	v42 = v39
	goto L20
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v51 = int32(0)
	v54 = v8
	goto L21
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v43+v51<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v63 = int32(421044)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[246])))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v67 == int32(0) {
		v86 = v66
		v87 = v67
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v115 = v61
	goto L13
L23:
	;
	if v87-v86 != 0 {
		goto L12
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	if v66 != v67 {
		v86 = v66
		v87 = v67
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v71 = v62
	v72 = v63
	goto L27
L27:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v75
		v87 = v76
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v86 = v75
	v87 = v76
	goto L24
L29:
	;
	v79 = int32(1)
	if v75 == v76 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v54 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v106 = v51 + int32(1)
	if v106 != v42 {
		v51 = v106
		v54 = v61
		goto L21
	} else {
		goto L40
	}
L35:
	;
	return
L36:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(126665), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(464155), int32(8404), int32(9746))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	goto L22
L41:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v126 = F_SearchSysCacheCopyAttName(m, v125, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	if v126 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v135 + int32(4)
	F_errmsg(m, int32(67610), v16)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(464155), int32(8424), int32(9746))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L35
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v152
	F_errmsg_internal(m, int32(411533), v14+int32(-16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(464155), int32(8409), int32(9746))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L35
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+89)))
	if v170 == int32(0) {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	if v115 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_pfree(m, v126)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L35
	} else {
		goto L63
	}
L54:
	;
	v173 = F_defGetInt32(m, v115)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L35
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v195 = *(*int64)(unsafe.Add(mBase, _consts[224]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v195
	v198 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
	goto L53
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+89)) = uint8(v173)
	F_CatalogTupleUpdate(m, v123, v126+int32(4), v126)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L35
	} else {
		goto L58
	}
L58:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v181 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166)+74)))
	v185 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v183, v184, v185, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L35
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v191
	goto L53
L62:
	;
	goto L61
L63:
	;
	F_sequence_close(m, v123, int32(3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L35
	} else {
		goto L64
	}
L64:
	;
	if v21 != int32(112) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	m.G0 = v16 - int32(-64)
	return
L66:
	;
	if v115 == int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	if l5 == int32(0) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v213 = F_find_inheritance_children(m, v212, l4)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L35
	} else {
		goto L69
	}
L69:
	;
	if v213 == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v217 <= int32(0) {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v227 = int32(0)
	goto L72
L72:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236+v227<<(uint(int32(2))%32))))
	v242 = F_table_open(m, v240, int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L35
	} else {
		goto L74
	}
L73:
	;
	goto L65
L74:
	;
	v244 = int32(1)
	F_ATExecSetIdentity(m, v14+int32(-12), v242, l2, l3, l4, v244, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L35
	} else {
		goto L75
	}
L75:
	;
	F_sequence_close(m, v242, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L35
	} else {
		goto L76
	}
L76:
	;
	v252 = v227 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v252 < v253 {
		v227 = v252
		goto L72
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L35
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(370851), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L35
	} else {
		goto L80
	}
L80:
	;
	F_errhint(m, int32(595173), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L35
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(464155), int32(8388), int32(9746))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L35
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L35
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	F_errmsg(m, int32(660745), v14+int32(-48))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L35
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(464155), int32(8433), int32(9746))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L35
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L35
	} else {
		goto L88
	}
L88:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v316 + int32(4)
	F_errmsg(m, int32(256724), v14+int32(-32))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L35
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(464155), int32(8439), int32(9746))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L35
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L35
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(233559), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L35
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(464155), int32(8393), int32(9746))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L35
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
					v29 = *(*int32)(unsafe.Add(mBase, _consts[206]))
					if v29 != 0 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v32 = int32(0)
						F_RunObjectPostAlterHook(m, int32(1259), v31, v32, v32, v32)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_sequence_close(m, v14, int32(3))
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
						F_sequence_close(m, v14, int32(3))
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
					F_errmsg_internal(m, int32(43320), v9)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errfinish(m, int32(464155), int32(18618), int32(9996))
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
