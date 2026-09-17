package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_plpgsql_yyparse(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v765 int32
	_ = v765
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v839 int32
	_ = v839
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v924 int32
	_ = v924
	var v966 int32
	_ = v966
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v994 int32
	_ = v994
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1079 int32
	_ = v1079
	var v1121 int32
	_ = v1121
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1159 int32
	_ = v1159
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1286 int32
	_ = v1286
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1316 int32
	_ = v1316
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1443 int32
	_ = v1443
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1651 int32
	_ = v1651
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1669 int32
	_ = v1669
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1688 int32
	_ = v1688
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1709 int32
	_ = v1709
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1773 int32
	_ = v1773
	var v1815 int32
	_ = v1815
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1934 int32
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v2052 int32
	_ = v2052
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2070 int32
	_ = v2070
	var v2076 int32
	_ = v2076
	var v2078 int32
	_ = v2078
	var v2089 int32
	_ = v2089
	var v2098 int32
	_ = v2098
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2110 int32
	_ = v2110
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2128 int32
	_ = v2128
	var v2174 int32
	_ = v2174
	var v2216 int32
	_ = v2216
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2346 int32
	_ = v2346
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2363 int32
	_ = v2363
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2374 int32
	_ = v2374
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
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2447 int32
	_ = v2447
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2546 int32
	_ = v2546
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2559 int32
	_ = v2559
	var v2566 int32
	_ = v2566
	var v2595 int32
	_ = v2595
	var v2601 int32
	_ = v2601
	var v2615 int32
	_ = v2615
	var v2618 int32
	_ = v2618
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2645 int32
	_ = v2645
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2654 int32
	_ = v2654
	var v2656 int32
	_ = v2656
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2700 int32
	_ = v2700
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2773 int32
	_ = v2773
	var v2783 int32
	_ = v2783
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
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
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2836 int32
	_ = v2836
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2841 int32
	_ = v2841
	var v2843 int32
	_ = v2843
	var v2848 int32
	_ = v2848
	var v2852 int32
	_ = v2852
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2877 int32
	_ = v2877
	var v2879 int32
	_ = v2879
	var v2881 int32
	_ = v2881
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2909 int32
	_ = v2909
	var v2911 int32
	_ = v2911
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2917 int32
	_ = v2917
	var v2919 int32
	_ = v2919
	var v2921 int32
	_ = v2921
	var v2923 int32
	_ = v2923
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2943 int32
	_ = v2943
	var v2948 int32
	_ = v2948
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2965 int32
	_ = v2965
	var v2968 int32
	_ = v2968
	var v2976 int32
	_ = v2976
	var v2980 int32
	_ = v2980
	var v2981 int32
	_ = v2981
	var v2983 int32
	_ = v2983
	var v2988 int32
	_ = v2988
	var v2990 int32
	_ = v2990
	var v2992 int32
	_ = v2992
	var v2994 int32
	_ = v2994
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3006 int32
	_ = v3006
	var v3008 int32
	_ = v3008
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3026 int32
	_ = v3026
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3046 int32
	_ = v3046
	var v3051 int32
	_ = v3051
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3059 int32
	_ = v3059
	var v3062 int32
	_ = v3062
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3068 int32
	_ = v3068
	var v3071 int32
	_ = v3071
	var v3079 int32
	_ = v3079
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3091 int32
	_ = v3091
	var v3093 int32
	_ = v3093
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3098 int32
	_ = v3098
	var v3102 int32
	_ = v3102
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3124 int32
	_ = v3124
	var v3129 int32
	_ = v3129
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3137 int32
	_ = v3137
	var v3140 int32
	_ = v3140
	var v3142 int32
	_ = v3142
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3149 int32
	_ = v3149
	var v3157 int32
	_ = v3157
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3176 int32
	_ = v3176
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3185 int32
	_ = v3185
	var v3188 int32
	_ = v3188
	var v3191 int32
	_ = v3191
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3235 int32
	_ = v3235
	var v3238 int32
	_ = v3238
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3244 int32
	_ = v3244
	var v3247 int32
	_ = v3247
	var v3255 int32
	_ = v3255
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3262 int32
	_ = v3262
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3270 int32
	_ = v3270
	var v3272 int32
	_ = v3272
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3287 int32
	_ = v3287
	var v3288 int32
	_ = v3288
	var v3290 int32
	_ = v3290
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3317 int32
	_ = v3317
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3330 int32
	_ = v3330
	var v3333 int32
	_ = v3333
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3339 int32
	_ = v3339
	var v3342 int32
	_ = v3342
	var v3350 int32
	_ = v3350
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3357 int32
	_ = v3357
	var v3362 int32
	_ = v3362
	var v3366 int32
	_ = v3366
	var v3370 int32
	_ = v3370
	var v3373 int32
	_ = v3373
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3383 int32
	_ = v3383
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3439 int32
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3447 int32
	_ = v3447
	var v3453 int32
	_ = v3453
	var v3456 int32
	_ = v3456
	var v3457 int32
	_ = v3457
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3475 int32
	_ = v3475
	var v3480 int32
	_ = v3480
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3491 int32
	_ = v3491
	var v3496 int32
	_ = v3496
	var v3498 int32
	_ = v3498
	var v3526 int32
	_ = v3526
	var v3528 int32
	_ = v3528
	var v3530 int32
	_ = v3530
	var v3534 int32
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3536 int32
	_ = v3536
	var v3537 int32
	_ = v3537
	var v3539 int32
	_ = v3539
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3549 int32
	_ = v3549
	var v3550 int32
	_ = v3550
	var v3553 int32
	_ = v3553
	var v3554 int32
	_ = v3554
	var v3556 int32
	_ = v3556
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3567 int32
	_ = v3567
	var v3570 int32
	_ = v3570
	var v3573 int32
	_ = v3573
	var v3576 int32
	_ = v3576
	var v3579 int32
	_ = v3579
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3590 int32
	_ = v3590
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3603 int32
	_ = v3603
	var v3606 int32
	_ = v3606
	var v3609 int32
	_ = v3609
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3620 int32
	_ = v3620
	var v3627 int32
	_ = v3627
	var v3628 int32
	_ = v3628
	var v3633 int32
	_ = v3633
	var v3636 int32
	_ = v3636
	var v3639 int32
	_ = v3639
	var v3642 int32
	_ = v3642
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3649 int32
	_ = v3649
	var v3650 int32
	_ = v3650
	var v3653 int32
	_ = v3653
	var v3660 int32
	_ = v3660
	var v3661 int32
	_ = v3661
	var v3666 int32
	_ = v3666
	var v3669 int32
	_ = v3669
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3676 int32
	_ = v3676
	var v3679 int32
	_ = v3679
	var v3680 int32
	_ = v3680
	var v3683 int32
	_ = v3683
	var v3690 int32
	_ = v3690
	var v3691 int32
	_ = v3691
	var v3696 int32
	_ = v3696
	var v3699 int32
	_ = v3699
	var v3702 int32
	_ = v3702
	var v3705 int32
	_ = v3705
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3712 int32
	_ = v3712
	var v3713 int32
	_ = v3713
	var v3716 int32
	_ = v3716
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3729 int32
	_ = v3729
	var v3732 int32
	_ = v3732
	var v3735 int32
	_ = v3735
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3742 int32
	_ = v3742
	var v3743 int32
	_ = v3743
	var v3746 int32
	_ = v3746
	var v3753 int32
	_ = v3753
	var v3754 int32
	_ = v3754
	var v3759 int32
	_ = v3759
	var v3762 int32
	_ = v3762
	var v3765 int32
	_ = v3765
	var v3768 int32
	_ = v3768
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3792 int32
	_ = v3792
	var v3795 int32
	_ = v3795
	var v3798 int32
	_ = v3798
	var v3801 int32
	_ = v3801
	var v3802 int32
	_ = v3802
	var v3805 int32
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3809 int32
	_ = v3809
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3822 int32
	_ = v3822
	var v3825 int32
	_ = v3825
	var v3828 int32
	_ = v3828
	var v3831 int32
	_ = v3831
	var v3834 int32
	_ = v3834
	var v3835 int32
	_ = v3835
	var v3838 int32
	_ = v3838
	var v3839 int32
	_ = v3839
	var v3842 int32
	_ = v3842
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3855 int32
	_ = v3855
	var v3858 int32
	_ = v3858
	var v3861 int32
	_ = v3861
	var v3864 int32
	_ = v3864
	var v3865 int32
	_ = v3865
	var v3868 int32
	_ = v3868
	var v3869 int32
	_ = v3869
	var v3872 int32
	_ = v3872
	var v3879 int32
	_ = v3879
	var v3880 int32
	_ = v3880
	var v3885 int32
	_ = v3885
	var v3888 int32
	_ = v3888
	var v3891 int32
	_ = v3891
	var v3894 int32
	_ = v3894
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3901 int32
	_ = v3901
	var v3902 int32
	_ = v3902
	var v3905 int32
	_ = v3905
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3918 int32
	_ = v3918
	var v3921 int32
	_ = v3921
	var v3924 int32
	_ = v3924
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3931 int32
	_ = v3931
	var v3932 int32
	_ = v3932
	var v3935 int32
	_ = v3935
	var v3942 int32
	_ = v3942
	var v3943 int32
	_ = v3943
	var v3948 int32
	_ = v3948
	var v3951 int32
	_ = v3951
	var v3954 int32
	_ = v3954
	var v3957 int32
	_ = v3957
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3964 int32
	_ = v3964
	var v3965 int32
	_ = v3965
	var v3968 int32
	_ = v3968
	var v3975 int32
	_ = v3975
	var v3976 int32
	_ = v3976
	var v3987 int32
	_ = v3987
	var v3988 int32
	_ = v3988
	var v3989 int32
	_ = v3989
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3998 int32
	_ = v3998
	var v3999 int32
	_ = v3999
	var v4001 int32
	_ = v4001
	var v4002 int32
	_ = v4002
	var v4005 int32
	_ = v4005
	var v4006 int32
	_ = v4006
	var v4011 int32
	_ = v4011
	var v4012 int32
	_ = v4012
	var v4017 int32
	_ = v4017
	var v4018 int32
	_ = v4018
	var v4021 int32
	_ = v4021
	var v4022 int32
	_ = v4022
	var v4024 int32
	_ = v4024
	var v4029 int32
	_ = v4029
	var v4031 int32
	_ = v4031
	var v4032 int32
	_ = v4032
	var v4037 int32
	_ = v4037
	var v4040 int32
	_ = v4040
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4046 int32
	_ = v4046
	var v4049 int32
	_ = v4049
	var v4057 int32
	_ = v4057
	var v4061 int32
	_ = v4061
	var v4062 int32
	_ = v4062
	var v4064 int32
	_ = v4064
	var v4069 int32
	_ = v4069
	var v4073 int32
	_ = v4073
	var v4077 int32
	_ = v4077
	var v4081 int32
	_ = v4081
	var v4087 int32
	_ = v4087
	var v4088 int32
	_ = v4088
	var v4091 int32
	_ = v4091
	var v4092 int32
	_ = v4092
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4101 int32
	_ = v4101
	var v4102 int32
	_ = v4102
	var v4104 int32
	_ = v4104
	var v4109 int32
	_ = v4109
	var v4111 int32
	_ = v4111
	var v4112 int32
	_ = v4112
	var v4117 int32
	_ = v4117
	var v4120 int32
	_ = v4120
	var v4122 int32
	_ = v4122
	var v4123 int32
	_ = v4123
	var v4126 int32
	_ = v4126
	var v4129 int32
	_ = v4129
	var v4137 int32
	_ = v4137
	var v4142 int32
	_ = v4142
	var v4144 int32
	_ = v4144
	var v4148 int32
	_ = v4148
	var v4149 int32
	_ = v4149
	var v4150 int32
	_ = v4150
	var v4154 int32
	_ = v4154
	var v4158 int32
	_ = v4158
	var v4159 int32
	_ = v4159
	var v4161 int32
	_ = v4161
	var v4164 int32
	_ = v4164
	var v4167 int32
	_ = v4167
	var v4168 int32
	_ = v4168
	var v4170 int32
	_ = v4170
	var v4173 int32
	_ = v4173
	var v4174 int32
	_ = v4174
	var v4177 int32
	_ = v4177
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4189 int32
	_ = v4189
	var v4194 int32
	_ = v4194
	var v4196 int32
	_ = v4196
	var v4197 int32
	_ = v4197
	var v4202 int32
	_ = v4202
	var v4205 int32
	_ = v4205
	var v4207 int32
	_ = v4207
	var v4208 int32
	_ = v4208
	var v4211 int32
	_ = v4211
	var v4214 int32
	_ = v4214
	var v4222 int32
	_ = v4222
	var v4226 int32
	_ = v4226
	var v4227 int32
	_ = v4227
	var v4229 int32
	_ = v4229
	var v4231 int32
	_ = v4231
	var v4241 int32
	_ = v4241
	var v4244 int32
	_ = v4244
	var v4245 int32
	_ = v4245
	var v4247 int32
	_ = v4247
	var v4252 int32
	_ = v4252
	var v4255 int32
	_ = v4255
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4263 int32
	_ = v4263
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4277 int32
	_ = v4277
	var v4285 int32
	_ = v4285
	var v4305 int32
	_ = v4305
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4312 int32
	_ = v4312
	var v4314 int32
	_ = v4314
	var v4315 int32
	_ = v4315
	var v4322 int32
	_ = v4322
	var v4323 int32
	_ = v4323
	var v4325 int32
	_ = v4325
	var v4326 int32
	_ = v4326
	var v4327 int32
	_ = v4327
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4333 int32
	_ = v4333
	var v4335 int32
	_ = v4335
	var v4337 int32
	_ = v4337
	var v4338 int32
	_ = v4338
	var v4371 int32
	_ = v4371
	var v4373 int32
	_ = v4373
	var v4374 int32
	_ = v4374
	var v4375 int32
	_ = v4375
	var v4379 int32
	_ = v4379
	var v4381 int32
	_ = v4381
	var v4385 int32
	_ = v4385
	var v4389 int32
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4391 int32
	_ = v4391
	var v4398 int32
	_ = v4398
	var v4402 int32
	_ = v4402
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4407 int32
	_ = v4407
	var v4413 int32
	_ = v4413
	var v4414 int32
	_ = v4414
	var v4417 int32
	_ = v4417
	var v4418 int32
	_ = v4418
	var v4421 int32
	_ = v4421
	var v4422 int32
	_ = v4422
	var v4427 int32
	_ = v4427
	var v4428 int32
	_ = v4428
	var v4431 int32
	_ = v4431
	var v4432 int32
	_ = v4432
	var v4434 int32
	_ = v4434
	var v4439 int32
	_ = v4439
	var v4441 int32
	_ = v4441
	var v4442 int32
	_ = v4442
	var v4447 int32
	_ = v4447
	var v4450 int32
	_ = v4450
	var v4452 int32
	_ = v4452
	var v4453 int32
	_ = v4453
	var v4456 int32
	_ = v4456
	var v4459 int32
	_ = v4459
	var v4467 int32
	_ = v4467
	var v4472 int32
	_ = v4472
	var v4474 int32
	_ = v4474
	var v4479 int32
	_ = v4479
	var v4481 int32
	_ = v4481
	var v4488 int32
	_ = v4488
	var v4489 int32
	_ = v4489
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4494 int32
	_ = v4494
	var v4498 int32
	_ = v4498
	var v4499 int32
	_ = v4499
	var v4504 int32
	_ = v4504
	var v4505 int32
	_ = v4505
	var v4508 int32
	_ = v4508
	var v4509 int32
	_ = v4509
	var v4511 int32
	_ = v4511
	var v4516 int32
	_ = v4516
	var v4518 int32
	_ = v4518
	var v4519 int32
	_ = v4519
	var v4524 int32
	_ = v4524
	var v4527 int32
	_ = v4527
	var v4529 int32
	_ = v4529
	var v4530 int32
	_ = v4530
	var v4533 int32
	_ = v4533
	var v4536 int32
	_ = v4536
	var v4544 int32
	_ = v4544
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4551 int32
	_ = v4551
	var v4555 int32
	_ = v4555
	var v4556 int32
	_ = v4556
	var v4558 int32
	_ = v4558
	var v4560 int32
	_ = v4560
	var v4561 int32
	_ = v4561
	var v4562 int32
	_ = v4562
	var v4564 int32
	_ = v4564
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4569 int32
	_ = v4569
	var v4570 int32
	_ = v4570
	var v4571 int32
	_ = v4571
	var v4572 int32
	_ = v4572
	var v4574 int32
	_ = v4574
	var v4578 int32
	_ = v4578
	var v4579 int32
	_ = v4579
	var v4584 int32
	_ = v4584
	var v4585 int32
	_ = v4585
	var v4590 int32
	_ = v4590
	var v4591 int32
	_ = v4591
	var v4594 int32
	_ = v4594
	var v4595 int32
	_ = v4595
	var v4597 int32
	_ = v4597
	var v4602 int32
	_ = v4602
	var v4604 int32
	_ = v4604
	var v4605 int32
	_ = v4605
	var v4610 int32
	_ = v4610
	var v4613 int32
	_ = v4613
	var v4615 int32
	_ = v4615
	var v4616 int32
	_ = v4616
	var v4619 int32
	_ = v4619
	var v4622 int32
	_ = v4622
	var v4630 int32
	_ = v4630
	var v4634 int32
	_ = v4634
	var v4635 int32
	_ = v4635
	var v4637 int32
	_ = v4637
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4646 int32
	_ = v4646
	var v4648 int32
	_ = v4648
	var v4650 int32
	_ = v4650
	var v4651 int32
	_ = v4651
	var v4652 int32
	_ = v4652
	var v4654 int32
	_ = v4654
	var v4657 int32
	_ = v4657
	var v4658 int32
	_ = v4658
	var v4659 int32
	_ = v4659
	var v4660 int32
	_ = v4660
	var v4661 int32
	_ = v4661
	var v4662 int32
	_ = v4662
	var v4664 int32
	_ = v4664
	var v4669 int32
	_ = v4669
	var v4670 int32
	_ = v4670
	var v4673 int32
	_ = v4673
	var v4674 int32
	_ = v4674
	var v4679 int32
	_ = v4679
	var v4680 int32
	_ = v4680
	var v4683 int32
	_ = v4683
	var v4684 int32
	_ = v4684
	var v4686 int32
	_ = v4686
	var v4691 int32
	_ = v4691
	var v4693 int32
	_ = v4693
	var v4694 int32
	_ = v4694
	var v4699 int32
	_ = v4699
	var v4702 int32
	_ = v4702
	var v4704 int32
	_ = v4704
	var v4705 int32
	_ = v4705
	var v4708 int32
	_ = v4708
	var v4711 int32
	_ = v4711
	var v4719 int32
	_ = v4719
	var v4723 int32
	_ = v4723
	var v4724 int32
	_ = v4724
	var v4730 int32
	_ = v4730
	var v4732 int32
	_ = v4732
	var v4735 int32
	_ = v4735
	var v4736 int32
	_ = v4736
	var v4737 int32
	_ = v4737
	var v4739 int32
	_ = v4739
	var v4742 int32
	_ = v4742
	var v4743 int32
	_ = v4743
	var v4744 int32
	_ = v4744
	var v4745 int32
	_ = v4745
	var v4746 int32
	_ = v4746
	var v4747 int32
	_ = v4747
	var v4749 int32
	_ = v4749
	var v4752 int32
	_ = v4752
	var v4754 int32
	_ = v4754
	var v4755 int32
	_ = v4755
	var v4756 int32
	_ = v4756
	var v4758 int32
	_ = v4758
	var v4765 int32
	_ = v4765
	var v4768 int32
	_ = v4768
	var v4773 int32
	_ = v4773
	var v4774 int32
	_ = v4774
	var v4776 int32
	_ = v4776
	var v4777 int32
	_ = v4777
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4784 int32
	_ = v4784
	var v4787 int32
	_ = v4787
	var v4788 int32
	_ = v4788
	var v4790 int32
	_ = v4790
	var v4791 int32
	_ = v4791
	var v4793 int32
	_ = v4793
	var v4794 int32
	_ = v4794
	var v4796 int32
	_ = v4796
	var v4799 int32
	_ = v4799
	var v4804 int32
	_ = v4804
	var v4807 int32
	_ = v4807
	var v4808 int32
	_ = v4808
	var v4809 int32
	_ = v4809
	var v4810 int32
	_ = v4810
	var v4814 int32
	_ = v4814
	var v4844 int32
	_ = v4844
	var v4847 int32
	_ = v4847
	var v4856 int32
	_ = v4856
	var v4857 int32
	_ = v4857
	var v4858 int32
	_ = v4858
	var v4859 int32
	_ = v4859
	var v4860 int32
	_ = v4860
	var v4862 int32
	_ = v4862
	var v4891 int32
	_ = v4891
	var v4892 int32
	_ = v4892
	var v4893 int32
	_ = v4893
	var v4894 int32
	_ = v4894
	var v4898 int32
	_ = v4898
	var v4899 int32
	_ = v4899
	var v4900 int32
	_ = v4900
	var v4903 int32
	_ = v4903
	var v4904 int32
	_ = v4904
	var v4906 int32
	_ = v4906
	var v4909 int32
	_ = v4909
	var v4913 int32
	_ = v4913
	var v4916 int32
	_ = v4916
	var v4917 int32
	_ = v4917
	var v4925 int32
	_ = v4925
	var v4926 int32
	_ = v4926
	var v4930 int32
	_ = v4930
	var v4933 int32
	_ = v4933
	var v4937 int32
	_ = v4937
	var v4938 int32
	_ = v4938
	var v4941 int32
	_ = v4941
	var v4944 int32
	_ = v4944
	var v4947 int32
	_ = v4947
	var v4950 int32
	_ = v4950
	var v4953 int32
	_ = v4953
	var v4956 int32
	_ = v4956
	var v4957 int32
	_ = v4957
	var v4960 int32
	_ = v4960
	var v4961 int32
	_ = v4961
	var v4964 int32
	_ = v4964
	var v4971 int32
	_ = v4971
	var v4972 int32
	_ = v4972
	var v4978 int32
	_ = v4978
	var v4980 int32
	_ = v4980
	var v4982 int32
	_ = v4982
	var v4983 int32
	_ = v4983
	var v4995 int32
	_ = v4995
	var v4996 int32
	_ = v4996
	var v4997 int32
	_ = v4997
	var v5000 int32
	_ = v5000
	var v5001 int32
	_ = v5001
	var v5002 int32
	_ = v5002
	var v5004 int32
	_ = v5004
	var v5006 int32
	_ = v5006
	var v5007 int32
	_ = v5007
	var v5011 int32
	_ = v5011
	var v5012 int32
	_ = v5012
	var v5014 int32
	_ = v5014
	var v5017 int32
	_ = v5017
	var v5018 int32
	_ = v5018
	var v5020 int32
	_ = v5020
	var v5021 int32
	_ = v5021
	var v5023 int32
	_ = v5023
	var v5024 int32
	_ = v5024
	var v5026 int32
	_ = v5026
	var v5029 int32
	_ = v5029
	var v5034 int32
	_ = v5034
	var v5037 int32
	_ = v5037
	var v5038 int32
	_ = v5038
	var v5039 int32
	_ = v5039
	var v5040 int32
	_ = v5040
	var v5045 int32
	_ = v5045
	var v5046 int32
	_ = v5046
	var v5047 int32
	_ = v5047
	var v5048 int32
	_ = v5048
	var v5049 int32
	_ = v5049
	var v5051 int32
	_ = v5051
	var v5052 int32
	_ = v5052
	var v5057 int32
	_ = v5057
	var v5058 int32
	_ = v5058
	var v5061 int32
	_ = v5061
	var v5062 int32
	_ = v5062
	var v5064 int32
	_ = v5064
	var v5069 int32
	_ = v5069
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5077 int32
	_ = v5077
	var v5080 int32
	_ = v5080
	var v5082 int32
	_ = v5082
	var v5083 int32
	_ = v5083
	var v5086 int32
	_ = v5086
	var v5089 int32
	_ = v5089
	var v5097 int32
	_ = v5097
	var v5100 int32
	_ = v5100
	var v5101 int32
	_ = v5101
	var v5102 int32
	_ = v5102
	var v5108 int32
	_ = v5108
	var v5114 int32
	_ = v5114
	var v5116 int32
	_ = v5116
	var v5117 int32
	_ = v5117
	var v5118 int32
	_ = v5118
	var v5120 int32
	_ = v5120
	var v5123 int32
	_ = v5123
	var v5124 int32
	_ = v5124
	var v5125 int32
	_ = v5125
	var v5126 int32
	_ = v5126
	var v5127 int32
	_ = v5127
	var v5129 int32
	_ = v5129
	var v5131 int32
	_ = v5131
	var v5132 int32
	_ = v5132
	var v5137 int32
	_ = v5137
	var v5138 int32
	_ = v5138
	var v5141 int32
	_ = v5141
	var v5142 int32
	_ = v5142
	var v5144 int32
	_ = v5144
	var v5149 int32
	_ = v5149
	var v5151 int32
	_ = v5151
	var v5152 int32
	_ = v5152
	var v5157 int32
	_ = v5157
	var v5160 int32
	_ = v5160
	var v5162 int32
	_ = v5162
	var v5163 int32
	_ = v5163
	var v5166 int32
	_ = v5166
	var v5169 int32
	_ = v5169
	var v5177 int32
	_ = v5177
	var v5183 int32
	_ = v5183
	var v5185 int32
	_ = v5185
	var v5186 int32
	_ = v5186
	var v5187 int32
	_ = v5187
	var v5189 int32
	_ = v5189
	var v5193 int32
	_ = v5193
	var v5194 int32
	_ = v5194
	var v5199 int32
	_ = v5199
	var v5200 int32
	_ = v5200
	var v5205 int32
	_ = v5205
	var v5206 int32
	_ = v5206
	var v5209 int32
	_ = v5209
	var v5210 int32
	_ = v5210
	var v5212 int32
	_ = v5212
	var v5217 int32
	_ = v5217
	var v5219 int32
	_ = v5219
	var v5220 int32
	_ = v5220
	var v5225 int32
	_ = v5225
	var v5228 int32
	_ = v5228
	var v5230 int32
	_ = v5230
	var v5231 int32
	_ = v5231
	var v5234 int32
	_ = v5234
	var v5237 int32
	_ = v5237
	var v5245 int32
	_ = v5245
	var v5249 int32
	_ = v5249
	var v5250 int32
	_ = v5250
	var v5252 int32
	_ = v5252
	var v5256 int32
	_ = v5256
	var v5257 int32
	_ = v5257
	var v5261 int32
	_ = v5261
	var v5265 int32
	_ = v5265
	var v5267 int32
	_ = v5267
	var v5270 int32
	_ = v5270
	var v5271 int32
	_ = v5271
	var v5275 int32
	_ = v5275
	var v5276 int32
	_ = v5276
	var v5279 int32
	_ = v5279
	var v5280 int32
	_ = v5280
	var v5281 int32
	_ = v5281
	var v5283 int32
	_ = v5283
	var v5286 int32
	_ = v5286
	var v5288 int32
	_ = v5288
	var v5289 int32
	_ = v5289
	var v5290 int32
	_ = v5290
	var v5291 int32
	_ = v5291
	var v5293 int32
	_ = v5293
	var v5296 int32
	_ = v5296
	var v5297 int32
	_ = v5297
	var v5298 int32
	_ = v5298
	var v5299 int32
	_ = v5299
	var v5300 int32
	_ = v5300
	var v5301 int32
	_ = v5301
	var v5303 int32
	_ = v5303
	var v5308 int32
	_ = v5308
	var v5311 int32
	_ = v5311
	var v5312 int32
	_ = v5312
	var v5316 int32
	_ = v5316
	var v5317 int32
	_ = v5317
	var v5319 int32
	_ = v5319
	var v5324 int32
	_ = v5324
	var v5327 int32
	_ = v5327
	var v5328 int32
	_ = v5328
	var v5329 int32
	_ = v5329
	var v5334 int32
	_ = v5334
	var v5335 int32
	_ = v5335
	var v5338 int32
	_ = v5338
	var v5339 int32
	_ = v5339
	var v5341 int32
	_ = v5341
	var v5346 int32
	_ = v5346
	var v5348 int32
	_ = v5348
	var v5349 int32
	_ = v5349
	var v5354 int32
	_ = v5354
	var v5357 int32
	_ = v5357
	var v5359 int32
	_ = v5359
	var v5360 int32
	_ = v5360
	var v5363 int32
	_ = v5363
	var v5366 int32
	_ = v5366
	var v5374 int32
	_ = v5374
	var v5378 int32
	_ = v5378
	var v5379 int32
	_ = v5379
	var v5381 int32
	_ = v5381
	var v5384 int32
	_ = v5384
	var v5385 int32
	_ = v5385
	var v5386 int32
	_ = v5386
	var v5387 int32
	_ = v5387
	var v5389 int32
	_ = v5389
	var v5392 int32
	_ = v5392
	var v5393 int32
	_ = v5393
	var v5399 int32
	_ = v5399
	var v5402 int32
	_ = v5402
	var v5405 int32
	_ = v5405
	var v5409 int32
	_ = v5409
	var v5412 int32
	_ = v5412
	var v5413 int32
	_ = v5413
	var v5419 int32
	_ = v5419
	var v5422 int32
	_ = v5422
	var v5423 int32
	_ = v5423
	var v5424 int32
	_ = v5424
	var v5429 int32
	_ = v5429
	var v5430 int32
	_ = v5430
	var v5431 int32
	_ = v5431
	var v5432 int32
	_ = v5432
	var v5435 int32
	_ = v5435
	var v5439 int32
	_ = v5439
	var v5444 int32
	_ = v5444
	var v5446 int32
	_ = v5446
	var v5452 int32
	_ = v5452
	var v5453 int32
	_ = v5453
	var v5458 int32
	_ = v5458
	var v5461 int32
	_ = v5461
	var v5464 int32
	_ = v5464
	var v5467 int32
	_ = v5467
	var v5470 int32
	_ = v5470
	var v5473 int32
	_ = v5473
	var v5474 int32
	_ = v5474
	var v5477 int32
	_ = v5477
	var v5478 int32
	_ = v5478
	var v5481 int32
	_ = v5481
	var v5488 int32
	_ = v5488
	var v5489 int32
	_ = v5489
	var v5492 int32
	_ = v5492
	var v5494 int32
	_ = v5494
	var v5496 int32
	_ = v5496
	var v5498 int32
	_ = v5498
	var v5499 int32
	_ = v5499
	var v5501 int32
	_ = v5501
	var v5502 int32
	_ = v5502
	var v5505 int32
	_ = v5505
	var v5510 int32
	_ = v5510
	var v5511 int32
	_ = v5511
	var v5514 int32
	_ = v5514
	var v5515 int32
	_ = v5515
	var v5517 int32
	_ = v5517
	var v5522 int32
	_ = v5522
	var v5524 int32
	_ = v5524
	var v5525 int32
	_ = v5525
	var v5530 int32
	_ = v5530
	var v5533 int32
	_ = v5533
	var v5535 int32
	_ = v5535
	var v5536 int32
	_ = v5536
	var v5539 int32
	_ = v5539
	var v5542 int32
	_ = v5542
	var v5550 int32
	_ = v5550
	var v5554 int32
	_ = v5554
	var v5555 int32
	_ = v5555
	var v5557 int32
	_ = v5557
	var v5562 int32
	_ = v5562
	var v5563 int32
	_ = v5563
	var v5564 int32
	_ = v5564
	var v5570 int32
	_ = v5570
	var v5571 int32
	_ = v5571
	var v5575 int32
	_ = v5575
	var v5576 int32
	_ = v5576
	var v5579 int32
	_ = v5579
	var v5580 int32
	_ = v5580
	var v5586 int32
	_ = v5586
	var v5588 int32
	_ = v5588
	var v5589 int32
	_ = v5589
	var v5593 int32
	_ = v5593
	var v5595 int32
	_ = v5595
	var v5599 int32
	_ = v5599
	var v5603 int32
	_ = v5603
	var v5604 int32
	_ = v5604
	var v5609 int32
	_ = v5609
	var v5612 int32
	_ = v5612
	var v5616 int32
	_ = v5616
	var v5617 int32
	_ = v5617
	var v5618 int32
	_ = v5618
	var v5623 int32
	_ = v5623
	var v5627 int32
	_ = v5627
	var v5630 int32
	_ = v5630
	var v5634 int32
	_ = v5634
	var v5635 int32
	_ = v5635
	var v5636 int32
	_ = v5636
	var v5637 int32
	_ = v5637
	var v5642 int32
	_ = v5642
	var v5646 int32
	_ = v5646
	var v5649 int32
	_ = v5649
	var v5652 int32
	_ = v5652
	var v5655 int32
	_ = v5655
	var v5656 int32
	_ = v5656
	var v5659 int32
	_ = v5659
	var v5660 int32
	_ = v5660
	var v5663 int32
	_ = v5663
	var v5670 int32
	_ = v5670
	var v5671 int32
	_ = v5671
	var v5674 int32
	_ = v5674
	var v5675 int32
	_ = v5675
	var v5677 int32
	_ = v5677
	var v5680 int32
	_ = v5680
	var v5681 int32
	_ = v5681
	var v5683 int32
	_ = v5683
	var v5685 int32
	_ = v5685
	var v5687 int32
	_ = v5687
	var v5688 int32
	_ = v5688
	var v5691 int32
	_ = v5691
	var v5696 int32
	_ = v5696
	var v5697 int32
	_ = v5697
	var v5700 int32
	_ = v5700
	var v5701 int32
	_ = v5701
	var v5703 int32
	_ = v5703
	var v5708 int32
	_ = v5708
	var v5710 int32
	_ = v5710
	var v5711 int32
	_ = v5711
	var v5716 int32
	_ = v5716
	var v5719 int32
	_ = v5719
	var v5721 int32
	_ = v5721
	var v5722 int32
	_ = v5722
	var v5725 int32
	_ = v5725
	var v5728 int32
	_ = v5728
	var v5736 int32
	_ = v5736
	var v5740 int32
	_ = v5740
	var v5741 int32
	_ = v5741
	var v5743 int32
	_ = v5743
	var v5746 int32
	_ = v5746
	var v5747 int32
	_ = v5747
	var v5751 int32
	_ = v5751
	var v5753 int32
	_ = v5753
	var v5761 int32
	_ = v5761
	var v5762 int32
	_ = v5762
	var v5766 int32
	_ = v5766
	var v5769 int32
	_ = v5769
	var v5774 int32
	_ = v5774
	var v5775 int32
	_ = v5775
	var v5777 int32
	_ = v5777
	var v5807 int32
	_ = v5807
	var v5810 int32
	_ = v5810
	var v5815 int32
	_ = v5815
	var v5816 int32
	_ = v5816
	var v5817 int32
	_ = v5817
	var v5818 int32
	_ = v5818
	var v5819 int32
	_ = v5819
	var v5821 int32
	_ = v5821
	var v5855 int32
	_ = v5855
	var v5858 int32
	_ = v5858
	var v5862 int32
	_ = v5862
	var v5863 int32
	_ = v5863
	var v5864 int32
	_ = v5864
	var v5869 int32
	_ = v5869
	var v5873 int32
	_ = v5873
	var v5875 int32
	_ = v5875
	var v5877 int32
	_ = v5877
	var v5878 int32
	_ = v5878
	var v5880 int32
	_ = v5880
	var v5881 int32
	_ = v5881
	var v5884 int32
	_ = v5884
	var v5889 int32
	_ = v5889
	var v5890 int32
	_ = v5890
	var v5893 int32
	_ = v5893
	var v5894 int32
	_ = v5894
	var v5896 int32
	_ = v5896
	var v5901 int32
	_ = v5901
	var v5903 int32
	_ = v5903
	var v5904 int32
	_ = v5904
	var v5909 int32
	_ = v5909
	var v5912 int32
	_ = v5912
	var v5914 int32
	_ = v5914
	var v5915 int32
	_ = v5915
	var v5918 int32
	_ = v5918
	var v5921 int32
	_ = v5921
	var v5929 int32
	_ = v5929
	var v5933 int32
	_ = v5933
	var v5934 int32
	_ = v5934
	var v5935 int32
	_ = v5935
	var v5936 int32
	_ = v5936
	var v5941 int32
	_ = v5941
	var v5944 int32
	_ = v5944
	var v5945 int32
	_ = v5945
	var v5951 int32
	_ = v5951
	var v5954 int32
	_ = v5954
	var v5958 int32
	_ = v5958
	var v5962 int32
	_ = v5962
	var v5963 int32
	_ = v5963
	var v5964 int32
	_ = v5964
	var v5965 int32
	_ = v5965
	var v5970 int32
	_ = v5970
	var v5971 int32
	_ = v5971
	var v5974 int32
	_ = v5974
	var v5975 int32
	_ = v5975
	var v5979 int32
	_ = v5979
	var v5980 int32
	_ = v5980
	var v5984 int32
	_ = v5984
	var v5989 int32
	_ = v5989
	var v5993 int32
	_ = v5993
	var v5994 int32
	_ = v5994
	var v5995 int32
	_ = v5995
	var v5996 int32
	_ = v5996
	var v6001 int32
	_ = v6001
	var v6002 int32
	_ = v6002
	var v6003 int32
	_ = v6003
	var v6004 int32
	_ = v6004
	var v6010 int32
	_ = v6010
	var v6011 int32
	_ = v6011
	var v6015 int32
	_ = v6015
	var v6016 int32
	_ = v6016
	var v6019 int32
	_ = v6019
	var v6020 int32
	_ = v6020
	var v6026 int32
	_ = v6026
	var v6028 int32
	_ = v6028
	var v6029 int32
	_ = v6029
	var v6033 int32
	_ = v6033
	var v6035 int32
	_ = v6035
	var v6039 int32
	_ = v6039
	var v6043 int32
	_ = v6043
	var v6044 int32
	_ = v6044
	var v6048 int32
	_ = v6048
	var v6052 int32
	_ = v6052
	var v6053 int32
	_ = v6053
	var v6054 int32
	_ = v6054
	var v6055 int32
	_ = v6055
	var v6060 int32
	_ = v6060
	var v6064 int32
	_ = v6064
	var v6067 int32
	_ = v6067
	var v6071 int32
	_ = v6071
	var v6072 int32
	_ = v6072
	var v6073 int32
	_ = v6073
	var v6074 int32
	_ = v6074
	var v6079 int32
	_ = v6079
	var v6085 int32
	_ = v6085
	var v6086 int32
	_ = v6086
	var v6089 int32
	_ = v6089
	var v6090 int32
	_ = v6090
	var v6095 int32
	_ = v6095
	var v6096 int32
	_ = v6096
	var v6099 int32
	_ = v6099
	var v6100 int32
	_ = v6100
	var v6102 int32
	_ = v6102
	var v6107 int32
	_ = v6107
	var v6109 int32
	_ = v6109
	var v6110 int32
	_ = v6110
	var v6115 int32
	_ = v6115
	var v6118 int32
	_ = v6118
	var v6120 int32
	_ = v6120
	var v6121 int32
	_ = v6121
	var v6124 int32
	_ = v6124
	var v6127 int32
	_ = v6127
	var v6135 int32
	_ = v6135
	var v6139 int32
	_ = v6139
	var v6140 int32
	_ = v6140
	var v6142 int32
	_ = v6142
	var v6144 int64
	_ = v6144
	var v6152 int32
	_ = v6152
	var v6156 int32
	_ = v6156
	var v6157 int32
	_ = v6157
	var v6158 int32
	_ = v6158
	var v6168 int32
	_ = v6168
	var v6173 int32
	_ = v6173
	var v6176 int32
	_ = v6176
	var v6179 int32
	_ = v6179
	var v6182 int32
	_ = v6182
	var v6185 int32
	_ = v6185
	var v6188 int32
	_ = v6188
	var v6189 int32
	_ = v6189
	var v6192 int32
	_ = v6192
	var v6193 int32
	_ = v6193
	var v6196 int32
	_ = v6196
	var v6203 int32
	_ = v6203
	var v6204 int32
	_ = v6204
	var v6213 int32
	_ = v6213
	var v6214 int32
	_ = v6214
	var v6215 int32
	_ = v6215
	var v6218 int32
	_ = v6218
	var v6221 int32
	_ = v6221
	var v6224 int32
	_ = v6224
	var v6225 int32
	_ = v6225
	var v6228 int32
	_ = v6228
	var v6229 int32
	_ = v6229
	var v6232 int32
	_ = v6232
	var v6239 int32
	_ = v6239
	var v6240 int32
	_ = v6240
	var v6249 int32
	_ = v6249
	var v6250 int32
	_ = v6250
	var v6251 int32
	_ = v6251
	var v6254 int32
	_ = v6254
	var v6257 int32
	_ = v6257
	var v6260 int32
	_ = v6260
	var v6263 int32
	_ = v6263
	var v6264 int32
	_ = v6264
	var v6267 int32
	_ = v6267
	var v6268 int32
	_ = v6268
	var v6271 int32
	_ = v6271
	var v6278 int32
	_ = v6278
	var v6279 int32
	_ = v6279
	var v6288 int32
	_ = v6288
	var v6289 int32
	_ = v6289
	var v6290 int32
	_ = v6290
	var v6293 int32
	_ = v6293
	var v6296 int32
	_ = v6296
	var v6299 int32
	_ = v6299
	var v6300 int32
	_ = v6300
	var v6303 int32
	_ = v6303
	var v6304 int32
	_ = v6304
	var v6307 int32
	_ = v6307
	var v6314 int32
	_ = v6314
	var v6315 int32
	_ = v6315
	var v6320 int32
	_ = v6320
	var v6323 int32
	_ = v6323
	var v6326 int32
	_ = v6326
	var v6329 int32
	_ = v6329
	var v6332 int32
	_ = v6332
	var v6333 int32
	_ = v6333
	var v6336 int32
	_ = v6336
	var v6337 int32
	_ = v6337
	var v6340 int32
	_ = v6340
	var v6347 int32
	_ = v6347
	var v6348 int32
	_ = v6348
	var v6353 int32
	_ = v6353
	var v6356 int32
	_ = v6356
	var v6359 int32
	_ = v6359
	var v6362 int32
	_ = v6362
	var v6363 int32
	_ = v6363
	var v6366 int32
	_ = v6366
	var v6367 int32
	_ = v6367
	var v6370 int32
	_ = v6370
	var v6377 int32
	_ = v6377
	var v6378 int32
	_ = v6378
	var v6388 int32
	_ = v6388
	var v6389 int32
	_ = v6389
	var v6391 int32
	_ = v6391
	var v6396 int32
	_ = v6396
	var v6402 int32
	_ = v6402
	var v6403 int32
	_ = v6403
	var v6410 int32
	_ = v6410
	var v6411 int32
	_ = v6411
	var v6414 int32
	_ = v6414
	var v6417 int32
	_ = v6417
	var v6420 int32
	_ = v6420
	var v6423 int32
	_ = v6423
	var v6426 int32
	_ = v6426
	var v6427 int32
	_ = v6427
	var v6430 int32
	_ = v6430
	var v6431 int32
	_ = v6431
	var v6434 int32
	_ = v6434
	var v6441 int32
	_ = v6441
	var v6442 int32
	_ = v6442
	var v6450 int32
	_ = v6450
	var v6451 int32
	_ = v6451
	var v6454 int32
	_ = v6454
	var v6455 int32
	_ = v6455
	var v6458 int32
	_ = v6458
	var v6462 int32
	_ = v6462
	var v6464 int32
	_ = v6464
	var v6465 int64
	_ = v6465
	var v6473 int32
	_ = v6473
	var v6477 int32
	_ = v6477
	var v6481 int32
	_ = v6481
	var v6487 int32
	_ = v6487
	var v6491 int32
	_ = v6491
	var v6492 int32
	_ = v6492
	var v6499 int32
	_ = v6499
	var v6500 int32
	_ = v6500
	var v6501 int32
	_ = v6501
	var v6505 int32
	_ = v6505
	var v6508 int32
	_ = v6508
	var v6512 int32
	_ = v6512
	var v6513 int32
	_ = v6513
	var v6521 int32
	_ = v6521
	var v6527 int32
	_ = v6527
	var v6529 int32
	_ = v6529
	var v6531 int32
	_ = v6531
	var v6541 int32
	_ = v6541
	var v6545 int32
	_ = v6545
	var v6550 int32
	_ = v6550
	var v6556 int32
	_ = v6556
	var v6557 int32
	_ = v6557
	var v6561 int32
	_ = v6561
	var v6565 int32
	_ = v6565
	var v6566 int32
	_ = v6566
	var v6567 int32
	_ = v6567
	var v6568 int32
	_ = v6568
	var v6571 int32
	_ = v6571
	var v6574 int32
	_ = v6574
	var v6575 int32
	_ = v6575
	var v6581 int32
	_ = v6581
	var v6582 int32
	_ = v6582
	var v6583 int32
	_ = v6583
	var v6593 int32
	_ = v6593
	var v6624 int32
	_ = v6624
	var v6633 int32
	_ = v6633
	var v6634 int32
	_ = v6634
	var v6635 int32
	_ = v6635
	var v6636 int32
	_ = v6636
	var v6637 int32
	_ = v6637
	var v6639 int32
	_ = v6639
	var v6646 int32
	_ = v6646
	var v6695 int32
	_ = v6695
	var v6697 int32
	_ = v6697
	var v6699 int32
	_ = v6699
	var v6701 int32
	_ = v6701
	var v6707 int32
	_ = v6707
	var v6728 int32
	_ = v6728
	var v6729 int32
	_ = v6729
	var v6732 int32
	_ = v6732
	var v6734 int32
	_ = v6734
	var v6735 int32
	_ = v6735
	var v6738 int32
	_ = v6738
	var v6739 int32
	_ = v6739
	var v6742 int32
	_ = v6742
	var v6745 int32
	_ = v6745
	var v6748 int32
	_ = v6748
	var v6751 int32
	_ = v6751
	var v6752 int32
	_ = v6752
	var v6755 int32
	_ = v6755
	var v6756 int32
	_ = v6756
	var v6759 int32
	_ = v6759
	var v6766 int32
	_ = v6766
	var v6767 int32
	_ = v6767
	var v6771 int32
	_ = v6771
	var v6774 int32
	_ = v6774
	var v6777 int32
	_ = v6777
	var v6780 int32
	_ = v6780
	var v6781 int32
	_ = v6781
	var v6784 int32
	_ = v6784
	var v6785 int32
	_ = v6785
	var v6788 int32
	_ = v6788
	var v6795 int32
	_ = v6795
	var v6796 int32
	_ = v6796
	var v6800 int32
	_ = v6800
	var v6803 int32
	_ = v6803
	var v6806 int32
	_ = v6806
	var v6809 int32
	_ = v6809
	var v6812 int32
	_ = v6812
	var v6813 int32
	_ = v6813
	var v6816 int32
	_ = v6816
	var v6817 int32
	_ = v6817
	var v6820 int32
	_ = v6820
	var v6827 int32
	_ = v6827
	var v6828 int32
	_ = v6828
	var v6833 int32
	_ = v6833
	var v6836 int32
	_ = v6836
	var v6839 int32
	_ = v6839
	var v6842 int32
	_ = v6842
	var v6843 int32
	_ = v6843
	var v6846 int32
	_ = v6846
	var v6847 int32
	_ = v6847
	var v6850 int32
	_ = v6850
	var v6857 int32
	_ = v6857
	var v6858 int32
	_ = v6858
	var v6863 int32
	_ = v6863
	var v6866 int32
	_ = v6866
	var v6869 int32
	_ = v6869
	var v6872 int32
	_ = v6872
	var v6875 int32
	_ = v6875
	var v6876 int32
	_ = v6876
	var v6879 int32
	_ = v6879
	var v6880 int32
	_ = v6880
	var v6883 int32
	_ = v6883
	var v6890 int32
	_ = v6890
	var v6891 int32
	_ = v6891
	var v6896 int32
	_ = v6896
	var v6899 int32
	_ = v6899
	var v6902 int32
	_ = v6902
	var v6905 int32
	_ = v6905
	var v6906 int32
	_ = v6906
	var v6909 int32
	_ = v6909
	var v6910 int32
	_ = v6910
	var v6913 int32
	_ = v6913
	var v6920 int32
	_ = v6920
	var v6921 int32
	_ = v6921
	var v6926 int32
	_ = v6926
	var v6929 int32
	_ = v6929
	var v6932 int32
	_ = v6932
	var v6935 int32
	_ = v6935
	var v6938 int32
	_ = v6938
	var v6939 int32
	_ = v6939
	var v6942 int32
	_ = v6942
	var v6943 int32
	_ = v6943
	var v6946 int32
	_ = v6946
	var v6953 int32
	_ = v6953
	var v6954 int32
	_ = v6954
	var v6959 int32
	_ = v6959
	var v6962 int32
	_ = v6962
	var v6965 int32
	_ = v6965
	var v6968 int32
	_ = v6968
	var v6969 int32
	_ = v6969
	var v6972 int32
	_ = v6972
	var v6973 int32
	_ = v6973
	var v6976 int32
	_ = v6976
	var v6983 int32
	_ = v6983
	var v6984 int32
	_ = v6984
	var v6989 int32
	_ = v6989
	var v6992 int32
	_ = v6992
	var v6995 int32
	_ = v6995
	var v6998 int32
	_ = v6998
	var v7001 int32
	_ = v7001
	var v7002 int32
	_ = v7002
	var v7005 int32
	_ = v7005
	var v7006 int32
	_ = v7006
	var v7009 int32
	_ = v7009
	var v7016 int32
	_ = v7016
	var v7017 int32
	_ = v7017
	var v7022 int32
	_ = v7022
	var v7025 int32
	_ = v7025
	var v7026 int32
	_ = v7026
	var v7035 int32
	_ = v7035
	var v7038 int32
	_ = v7038
	var v7043 int32
	_ = v7043
	var v7044 int32
	_ = v7044
	var v7046 int32
	_ = v7046
	var v7047 int32
	_ = v7047
	var v7048 int32
	_ = v7048
	var v7057 int32
	_ = v7057
	var v7063 int32
	_ = v7063
	var v7067 int32
	_ = v7067
	var v7095 int32
	_ = v7095
	var v7100 int32
	_ = v7100
	var v7101 int32
	_ = v7101
	var v7121 int32
	_ = v7121
	var v7124 int32
	_ = v7124
	var v7125 int32
	_ = v7125
	var v7130 int32
	_ = v7130
	var v7132 int32
	_ = v7132
	var v7133 int32
	_ = v7133
	var v7135 int32
	_ = v7135
	var v7136 int32
	_ = v7136
	var v7140 int32
	_ = v7140
	var v7142 int32
	_ = v7142
	var v7172 int32
	_ = v7172
	var v7175 int32
	_ = v7175
	var v7179 int32
	_ = v7179
	var v7184 int32
	_ = v7184
	var v7189 int32
	_ = v7189
	var v7192 int32
	_ = v7192
	var v7196 int32
	_ = v7196
	var v7201 int32
	_ = v7201
	var v7204 int32
	_ = v7204
	var v7205 int32
	_ = v7205
	var v7208 int32
	_ = v7208
	var v7209 int32
	_ = v7209
	var v7214 int32
	_ = v7214
	var v7215 int32
	_ = v7215
	var v7218 int32
	_ = v7218
	var v7219 int32
	_ = v7219
	var v7221 int32
	_ = v7221
	var v7226 int32
	_ = v7226
	var v7228 int32
	_ = v7228
	var v7229 int32
	_ = v7229
	var v7234 int32
	_ = v7234
	var v7237 int32
	_ = v7237
	var v7239 int32
	_ = v7239
	var v7240 int32
	_ = v7240
	var v7243 int32
	_ = v7243
	var v7246 int32
	_ = v7246
	var v7254 int32
	_ = v7254
	var v7258 int32
	_ = v7258
	var v7259 int32
	_ = v7259
	var v7260 int32
	_ = v7260
	var v7261 int32
	_ = v7261
	var v7266 int32
	_ = v7266
	var v7275 int32
	_ = v7275
	var v7277 int32
	_ = v7277
	var v7278 int32
	_ = v7278
	var v7279 int32
	_ = v7279
	var v7281 int32
	_ = v7281
	var v7285 int32
	_ = v7285
	var v7289 int32
	_ = v7289
	var v7293 int32
	_ = v7293
	var v7294 int32
	_ = v7294
	var v7296 int32
	_ = v7296
	var v7301 int32
	_ = v7301
	var v7305 int32
	_ = v7305
	var v7309 int32
	_ = v7309
	var v7312 int32
	_ = v7312
	var v7318 int32
	_ = v7318
	var v7319 int32
	_ = v7319
	var v7322 int32
	_ = v7322
	var v7328 int32
	_ = v7328
	var v7329 int32
	_ = v7329
	var v7332 int32
	_ = v7332
	var v7338 int32
	_ = v7338
	var v7339 int32
	_ = v7339
	var v7342 int32
	_ = v7342
	var v7344 int32
	_ = v7344
	var v7345 int32
	_ = v7345
	var v7346 int32
	_ = v7346
	var v7348 int32
	_ = v7348
	var v7349 int32
	_ = v7349
	var v7362 int32
	_ = v7362
	var v7363 int32
	_ = v7363
	var v7366 int32
	_ = v7366
	var v7368 int32
	_ = v7368
	var v7369 int32
	_ = v7369
	var v7370 int32
	_ = v7370
	var v7372 int32
	_ = v7372
	var v7373 int32
	_ = v7373
	var v7387 int32
	_ = v7387
	var v7388 int32
	_ = v7388
	var v7395 int32
	_ = v7395
	var v7404 int32
	_ = v7404
	var v7405 int32
	_ = v7405
	var v7407 int32
	_ = v7407
	var v7408 int32
	_ = v7408
	var v7411 int32
	_ = v7411
	var v7412 int32
	_ = v7412
	var v7417 int32
	_ = v7417
	var v7418 int32
	_ = v7418
	var v7421 int32
	_ = v7421
	var v7422 int32
	_ = v7422
	var v7424 int32
	_ = v7424
	var v7429 int32
	_ = v7429
	var v7431 int32
	_ = v7431
	var v7432 int32
	_ = v7432
	var v7437 int32
	_ = v7437
	var v7440 int32
	_ = v7440
	var v7442 int32
	_ = v7442
	var v7443 int32
	_ = v7443
	var v7446 int32
	_ = v7446
	var v7449 int32
	_ = v7449
	var v7457 int32
	_ = v7457
	var v7461 int32
	_ = v7461
	var v7462 int32
	_ = v7462
	var v7464 int32
	_ = v7464
	var v7468 int32
	_ = v7468
	var v7476 int32
	_ = v7476
	var v7481 int32
	_ = v7481
	var v7508 int32
	_ = v7508
	var v7511 int32
	_ = v7511
	var v7514 int32
	_ = v7514
	var v7516 int32
	_ = v7516
	var v7518 int32
	_ = v7518
	var v7519 int32
	_ = v7519
	var v7520 int32
	_ = v7520
	var v7522 int32
	_ = v7522
	var v7553 int32
	_ = v7553
	var v7562 int32
	_ = v7562
	var v7563 int32
	_ = v7563
	var v7564 int32
	_ = v7564
	var v7565 int32
	_ = v7565
	var v7566 int32
	_ = v7566
	var v7568 int32
	_ = v7568
	var v7573 int32
	_ = v7573
	var v7574 int32
	_ = v7574
	var v7579 int32
	_ = v7579
	var v7580 int32
	_ = v7580
	var v7585 int32
	_ = v7585
	var v7586 int32
	_ = v7586
	var v7589 int32
	_ = v7589
	var v7590 int32
	_ = v7590
	var v7592 int32
	_ = v7592
	var v7597 int32
	_ = v7597
	var v7599 int32
	_ = v7599
	var v7600 int32
	_ = v7600
	var v7605 int32
	_ = v7605
	var v7608 int32
	_ = v7608
	var v7610 int32
	_ = v7610
	var v7611 int32
	_ = v7611
	var v7614 int32
	_ = v7614
	var v7617 int32
	_ = v7617
	var v7625 int32
	_ = v7625
	var v7629 int32
	_ = v7629
	var v7630 int32
	_ = v7630
	var v7632 int32
	_ = v7632
	var v7635 int32
	_ = v7635
	var v7636 int32
	_ = v7636
	var v7640 int32
	_ = v7640
	var v7641 int32
	_ = v7641
	var v7644 int32
	_ = v7644
	var v7649 int32
	_ = v7649
	var v7650 int32
	_ = v7650
	var v7657 int32
	_ = v7657
	var v7660 int32
	_ = v7660
	var v7663 int32
	_ = v7663
	var v7666 int32
	_ = v7666
	var v7669 int32
	_ = v7669
	var v7671 int32
	_ = v7671
	var v7676 int32
	_ = v7676
	var v7677 int32
	_ = v7677
	var v7682 int32
	_ = v7682
	var v7685 int32
	_ = v7685
	var v7688 int32
	_ = v7688
	var v7691 int32
	_ = v7691
	var v7694 int32
	_ = v7694
	var v7697 int32
	_ = v7697
	var v7698 int32
	_ = v7698
	var v7701 int32
	_ = v7701
	var v7702 int32
	_ = v7702
	var v7705 int32
	_ = v7705
	var v7712 int32
	_ = v7712
	var v7713 int32
	_ = v7713
	var v7717 int32
	_ = v7717
	var v7720 int32
	_ = v7720
	var v7721 int32
	_ = v7721
	var v7724 int32
	_ = v7724
	var v7727 int32
	_ = v7727
	var v7730 int32
	_ = v7730
	var v7731 int32
	_ = v7731
	var v7734 int32
	_ = v7734
	var v7735 int32
	_ = v7735
	var v7738 int32
	_ = v7738
	var v7745 int32
	_ = v7745
	var v7746 int32
	_ = v7746
	var v7749 int32
	_ = v7749
	var v7750 int32
	_ = v7750
	var v7757 int32
	_ = v7757
	var v7758 int32
	_ = v7758
	var v7759 int32
	_ = v7759
	var v7764 int32
	_ = v7764
	var v7766 int32
	_ = v7766
	var v7767 int32
	_ = v7767
	var v7768 int32
	_ = v7768
	var v7773 int32
	_ = v7773
	var v7776 int32
	_ = v7776
	var v7781 int32
	_ = v7781
	var v7782 int32
	_ = v7782
	var v7784 int32
	_ = v7784
	var v7814 int32
	_ = v7814
	var v7817 int32
	_ = v7817
	var v7826 int32
	_ = v7826
	var v7827 int32
	_ = v7827
	var v7828 int32
	_ = v7828
	var v7829 int32
	_ = v7829
	var v7830 int32
	_ = v7830
	var v7832 int32
	_ = v7832
	var v7862 int32
	_ = v7862
	var v7864 int32
	_ = v7864
	var v7866 int32
	_ = v7866
	var v7867 int32
	_ = v7867
	var v7868 int32
	_ = v7868
	var v7876 int32
	_ = v7876
	var v7877 int32
	_ = v7877
	var v7882 int32
	_ = v7882
	var v7887 int32
	_ = v7887
	var v7889 int32
	_ = v7889
	var v7891 int32
	_ = v7891
	var v7892 int32
	_ = v7892
	var v7893 int32
	_ = v7893
	var v7896 int32
	_ = v7896
	var v7901 int32
	_ = v7901
	var v7902 int32
	_ = v7902
	var v7907 int32
	_ = v7907
	var v7908 int32
	_ = v7908
	var v7911 int32
	_ = v7911
	var v7912 int32
	_ = v7912
	var v7914 int32
	_ = v7914
	var v7919 int32
	_ = v7919
	var v7921 int32
	_ = v7921
	var v7922 int32
	_ = v7922
	var v7927 int32
	_ = v7927
	var v7930 int32
	_ = v7930
	var v7932 int32
	_ = v7932
	var v7933 int32
	_ = v7933
	var v7936 int32
	_ = v7936
	var v7939 int32
	_ = v7939
	var v7947 int32
	_ = v7947
	var v7950 int32
	_ = v7950
	var v7954 int32
	_ = v7954
	var v7955 int32
	_ = v7955
	var v7956 int32
	_ = v7956
	var v7962 int32
	_ = v7962
	var v7965 int32
	_ = v7965
	var v7966 int32
	_ = v7966
	var v7971 int32
	_ = v7971
	var v7972 int32
	_ = v7972
	var v7975 int32
	_ = v7975
	var v7976 int32
	_ = v7976
	var v7978 int32
	_ = v7978
	var v7983 int32
	_ = v7983
	var v7985 int32
	_ = v7985
	var v7986 int32
	_ = v7986
	var v7991 int32
	_ = v7991
	var v7994 int32
	_ = v7994
	var v7996 int32
	_ = v7996
	var v7997 int32
	_ = v7997
	var v8000 int32
	_ = v8000
	var v8003 int32
	_ = v8003
	var v8011 int32
	_ = v8011
	var v8016 int32
	_ = v8016
	var v8017 int32
	_ = v8017
	var v8018 int32
	_ = v8018
	var v8023 int32
	_ = v8023
	var v8025 int32
	_ = v8025
	var v8026 int32
	_ = v8026
	var v8028 int32
	_ = v8028
	var v8030 int32
	_ = v8030
	var v8033 int32
	_ = v8033
	var v8034 int32
	_ = v8034
	var v8038 int32
	_ = v8038
	var v8039 int32
	_ = v8039
	var v8041 int32
	_ = v8041
	var v8043 int32
	_ = v8043
	var v8050 int32
	_ = v8050
	var v8051 int32
	_ = v8051
	var v8059 int32
	_ = v8059
	var v8060 int32
	_ = v8060
	var v8061 int32
	_ = v8061
	var v8062 int32
	_ = v8062
	var v8065 int32
	_ = v8065
	var v8068 int32
	_ = v8068
	var v8071 int32
	_ = v8071
	var v8072 int32
	_ = v8072
	var v8075 int32
	_ = v8075
	var v8076 int32
	_ = v8076
	var v8079 int32
	_ = v8079
	var v8086 int32
	_ = v8086
	var v8087 int32
	_ = v8087
	var v8091 int32
	_ = v8091
	var v8094 int32
	_ = v8094
	var v8097 int32
	_ = v8097
	var v8100 int32
	_ = v8100
	var v8103 int32
	_ = v8103
	var v8104 int32
	_ = v8104
	var v8107 int32
	_ = v8107
	var v8108 int32
	_ = v8108
	var v8111 int32
	_ = v8111
	var v8118 int32
	_ = v8118
	var v8119 int32
	_ = v8119
	var v8124 int32
	_ = v8124
	var v8127 int32
	_ = v8127
	var v8130 int32
	_ = v8130
	var v8133 int32
	_ = v8133
	var v8136 int32
	_ = v8136
	var v8137 int32
	_ = v8137
	var v8140 int32
	_ = v8140
	var v8141 int32
	_ = v8141
	var v8144 int32
	_ = v8144
	var v8151 int32
	_ = v8151
	var v8152 int32
	_ = v8152
	var v8157 int32
	_ = v8157
	var v8160 int32
	_ = v8160
	var v8163 int32
	_ = v8163
	var v8166 int32
	_ = v8166
	var v8169 int32
	_ = v8169
	var v8170 int32
	_ = v8170
	var v8173 int32
	_ = v8173
	var v8174 int32
	_ = v8174
	var v8177 int32
	_ = v8177
	var v8184 int32
	_ = v8184
	var v8185 int32
	_ = v8185
	var v8190 int32
	_ = v8190
	var v8193 int32
	_ = v8193
	var v8196 int32
	_ = v8196
	var v8199 int32
	_ = v8199
	var v8202 int32
	_ = v8202
	var v8203 int32
	_ = v8203
	var v8206 int32
	_ = v8206
	var v8207 int32
	_ = v8207
	var v8210 int32
	_ = v8210
	var v8217 int32
	_ = v8217
	var v8218 int32
	_ = v8218
	var v8221 int32
	_ = v8221
	var v8225 int32
	_ = v8225
	var v8228 int32
	_ = v8228
	var v8232 int32
	_ = v8232
	var v8233 int32
	_ = v8233
	var v8235 int32
	_ = v8235
	var v8237 int32
	_ = v8237
	var v8240 int32
	_ = v8240
	var v8243 int32
	_ = v8243
	var v8246 int32
	_ = v8246
	var v8249 int32
	_ = v8249
	var v8250 int32
	_ = v8250
	var v8253 int32
	_ = v8253
	var v8254 int32
	_ = v8254
	var v8257 int32
	_ = v8257
	var v8264 int32
	_ = v8264
	var v8265 int32
	_ = v8265
	var v8272 int32
	_ = v8272
	var v8275 int32
	_ = v8275
	var v8279 int32
	_ = v8279
	var v8280 int32
	_ = v8280
	var v8282 int32
	_ = v8282
	var v8288 int32
	_ = v8288
	var v8292 int32
	_ = v8292
	var v8293 int32
	_ = v8293
	var v8296 int32
	_ = v8296
	var v8302 int32
	_ = v8302
	var v8303 int32
	_ = v8303
	var v8310 int32
	_ = v8310
	var v8314 int32
	_ = v8314
	var v8315 int32
	_ = v8315
	var v8318 int32
	_ = v8318
	var v8324 int32
	_ = v8324
	var v8328 int32
	_ = v8328
	var v8332 int32
	_ = v8332
	var v8336 int32
	_ = v8336
	var v8337 int32
	_ = v8337
	var v8340 int32
	_ = v8340
	var v8346 int32
	_ = v8346
	var v8352 int32
	_ = v8352
	var v8355 int32
	_ = v8355
	var v8360 int32
	_ = v8360
	var v8363 int32
	_ = v8363
	var v8366 int32
	_ = v8366
	var v8370 int32
	_ = v8370
	var v8371 int32
	_ = v8371
	var v8372 int32
	_ = v8372
	var v8375 int32
	_ = v8375
	var v8379 int32
	_ = v8379
	var v8380 int32
	_ = v8380
	var v8384 int32
	_ = v8384
	var v8387 int32
	_ = v8387
	var v8388 int32
	_ = v8388
	var v8394 int32
	_ = v8394
	var v8400 int32
	_ = v8400
	var v8401 int32
	_ = v8401
	var v8406 int32
	_ = v8406
	var v8407 int32
	_ = v8407
	var v8412 int32
	_ = v8412
	var v8413 int32
	_ = v8413
	var v8416 int32
	_ = v8416
	var v8417 int32
	_ = v8417
	var v8419 int32
	_ = v8419
	var v8424 int32
	_ = v8424
	var v8426 int32
	_ = v8426
	var v8427 int32
	_ = v8427
	var v8432 int32
	_ = v8432
	var v8435 int32
	_ = v8435
	var v8437 int32
	_ = v8437
	var v8438 int32
	_ = v8438
	var v8441 int32
	_ = v8441
	var v8444 int32
	_ = v8444
	var v8452 int32
	_ = v8452
	var v8456 int32
	_ = v8456
	var v8457 int32
	_ = v8457
	var v8459 int32
	_ = v8459
	var v8464 int32
	_ = v8464
	var v8465 int32
	_ = v8465
	var v8471 int32
	_ = v8471
	var v8472 int32
	_ = v8472
	var v8477 int32
	_ = v8477
	var v8478 int32
	_ = v8478
	var v8483 int32
	_ = v8483
	var v8484 int32
	_ = v8484
	var v8487 int32
	_ = v8487
	var v8488 int32
	_ = v8488
	var v8490 int32
	_ = v8490
	var v8495 int32
	_ = v8495
	var v8497 int32
	_ = v8497
	var v8498 int32
	_ = v8498
	var v8503 int32
	_ = v8503
	var v8506 int32
	_ = v8506
	var v8508 int32
	_ = v8508
	var v8509 int32
	_ = v8509
	var v8512 int32
	_ = v8512
	var v8515 int32
	_ = v8515
	var v8523 int32
	_ = v8523
	var v8527 int32
	_ = v8527
	var v8528 int32
	_ = v8528
	var v8530 int32
	_ = v8530
	var v8535 int32
	_ = v8535
	var v8541 int32
	_ = v8541
	var v8542 int32
	_ = v8542
	var v8547 int32
	_ = v8547
	var v8548 int32
	_ = v8548
	var v8553 int32
	_ = v8553
	var v8554 int32
	_ = v8554
	var v8557 int32
	_ = v8557
	var v8558 int32
	_ = v8558
	var v8560 int32
	_ = v8560
	var v8565 int32
	_ = v8565
	var v8567 int32
	_ = v8567
	var v8568 int32
	_ = v8568
	var v8573 int32
	_ = v8573
	var v8576 int32
	_ = v8576
	var v8578 int32
	_ = v8578
	var v8579 int32
	_ = v8579
	var v8582 int32
	_ = v8582
	var v8585 int32
	_ = v8585
	var v8593 int32
	_ = v8593
	var v8597 int32
	_ = v8597
	var v8598 int32
	_ = v8598
	var v8600 int32
	_ = v8600
	var v8605 int32
	_ = v8605
	var v8616 int32
	_ = v8616
	var v8617 int32
	_ = v8617
	var v8618 int32
	_ = v8618
	var v8619 int32
	_ = v8619
	var v8622 int32
	_ = v8622
	var v8623 int32
	_ = v8623
	var v8624 int32
	_ = v8624
	var v8630 int32
	_ = v8630
	var v8631 int32
	_ = v8631
	var v8636 int32
	_ = v8636
	var v8637 int32
	_ = v8637
	var v8640 int32
	_ = v8640
	var v8641 int32
	_ = v8641
	var v8643 int32
	_ = v8643
	var v8648 int32
	_ = v8648
	var v8650 int32
	_ = v8650
	var v8651 int32
	_ = v8651
	var v8656 int32
	_ = v8656
	var v8659 int32
	_ = v8659
	var v8661 int32
	_ = v8661
	var v8662 int32
	_ = v8662
	var v8665 int32
	_ = v8665
	var v8668 int32
	_ = v8668
	var v8676 int32
	_ = v8676
	var v8679 int32
	_ = v8679
	var v8680 int32
	_ = v8680
	var v8682 int32
	_ = v8682
	var v8683 int32
	_ = v8683
	var v8688 int32
	_ = v8688
	var v8690 int32
	_ = v8690
	var v8691 int32
	_ = v8691
	var v8693 int32
	_ = v8693
	var v8694 int32
	_ = v8694
	var v8695 int32
	_ = v8695
	var v8697 int32
	_ = v8697
	var v8703 int32
	_ = v8703
	var v8704 int32
	_ = v8704
	var v8706 int32
	_ = v8706
	var v8707 int32
	_ = v8707
	var v8709 int32
	_ = v8709
	var v8710 int32
	_ = v8710
	var v8711 int32
	_ = v8711
	var v8713 int32
	_ = v8713
	var v8718 int32
	_ = v8718
	var v8719 int32
	_ = v8719
	var v8724 int32
	_ = v8724
	var v8725 int32
	_ = v8725
	var v8726 int32
	_ = v8726
	var v8727 int32
	_ = v8727
	var v8729 int32
	_ = v8729
	var v8735 int32
	_ = v8735
	var v8736 int32
	_ = v8736
	var v8739 int32
	_ = v8739
	var v8740 int32
	_ = v8740
	var v8743 int32
	_ = v8743
	var v8744 int32
	_ = v8744
	var v8749 int32
	_ = v8749
	var v8750 int32
	_ = v8750
	var v8753 int32
	_ = v8753
	var v8754 int32
	_ = v8754
	var v8756 int32
	_ = v8756
	var v8761 int32
	_ = v8761
	var v8763 int32
	_ = v8763
	var v8764 int32
	_ = v8764
	var v8769 int32
	_ = v8769
	var v8772 int32
	_ = v8772
	var v8774 int32
	_ = v8774
	var v8775 int32
	_ = v8775
	var v8778 int32
	_ = v8778
	var v8781 int32
	_ = v8781
	var v8789 int32
	_ = v8789
	var v8794 int32
	_ = v8794
	var v8796 int32
	_ = v8796
	var v8800 int32
	_ = v8800
	var v8801 int32
	_ = v8801
	var v8806 int32
	_ = v8806
	var v8827 int32
	_ = v8827
	var v8828 int32
	_ = v8828
	var v8830 int32
	_ = v8830
	var v8832 int32
	_ = v8832
	var v8834 int32
	_ = v8834
	var v8835 int32
	_ = v8835
	var v8838 int32
	_ = v8838
	var v8841 int32
	_ = v8841
	var v8844 int32
	_ = v8844
	var v8845 int32
	_ = v8845
	var v8848 int32
	_ = v8848
	var v8849 int32
	_ = v8849
	var v8852 int32
	_ = v8852
	var v8859 int32
	_ = v8859
	var v8860 int32
	_ = v8860
	var v8862 int32
	_ = v8862
	var v8863 int32
	_ = v8863
	var v8869 int32
	_ = v8869
	var v8870 int32
	_ = v8870
	var v8873 int32
	_ = v8873
	var v8874 int32
	_ = v8874
	var v8877 int32
	_ = v8877
	var v8881 int32
	_ = v8881
	var v8883 int32
	_ = v8883
	var v8884 int64
	_ = v8884
	var v8892 int32
	_ = v8892
	var v8896 int32
	_ = v8896
	var v8900 int32
	_ = v8900
	var v8906 int32
	_ = v8906
	var v8910 int32
	_ = v8910
	var v8911 int32
	_ = v8911
	var v8918 int32
	_ = v8918
	var v8919 int32
	_ = v8919
	var v8920 int32
	_ = v8920
	var v8924 int32
	_ = v8924
	var v8927 int32
	_ = v8927
	var v8931 int32
	_ = v8931
	var v8932 int32
	_ = v8932
	var v8940 int32
	_ = v8940
	var v8946 int32
	_ = v8946
	var v8948 int32
	_ = v8948
	var v8950 int32
	_ = v8950
	var v8960 int32
	_ = v8960
	var v8964 int32
	_ = v8964
	var v8965 int32
	_ = v8965
	var v8966 int32
	_ = v8966
	var v8967 int32
	_ = v8967
	var v8968 int32
	_ = v8968
	var v8969 int32
	_ = v8969
	var v8970 int32
	_ = v8970
	var v8974 int32
	_ = v8974
	var v8976 int32
	_ = v8976
	var v9009 int32
	_ = v9009
	var v9013 int32
	_ = v9013
	var v9021 int32
	_ = v9021
	var v9022 int32
	_ = v9022
	var v9025 int32
	_ = v9025
	var v9029 int32
	_ = v9029
	var v9037 int32
	_ = v9037
	var v9038 int32
	_ = v9038
	var v9041 int32
	_ = v9041
	var v9045 int32
	_ = v9045
	var v9053 int32
	_ = v9053
	var v9054 int32
	_ = v9054
	var v9056 int32
	_ = v9056
	var v9059 int32
	_ = v9059
	var v9063 int32
	_ = v9063
	var v9064 int32
	_ = v9064
	var v9067 int32
	_ = v9067
	var v9068 int32
	_ = v9068
	var v9073 int32
	_ = v9073
	var v9077 int32
	_ = v9077
	var v9078 int32
	_ = v9078
	var v9081 int32
	_ = v9081
	var v9082 int32
	_ = v9082
	var v9086 int32
	_ = v9086
	var v9090 int32
	_ = v9090
	var v9092 int32
	_ = v9092
	var v9094 int32
	_ = v9094
	var v9095 int32
	_ = v9095
	var v9096 int32
	_ = v9096
	var v9098 int32
	_ = v9098
	var v9104 int32
	_ = v9104
	var v9117 int32
	_ = v9117
	var v9121 int32
	_ = v9121
	var v9124 int32
	_ = v9124
	var v9125 int32
	_ = v9125
	var v9126 int32
	_ = v9126
	var v9127 int32
	_ = v9127
	var v9128 int32
	_ = v9128
	var v9134 int32
	_ = v9134
	var v9137 int32
	_ = v9137
	var v9138 int32
	_ = v9138
	var v9139 int32
	_ = v9139
	var v9144 int32
	_ = v9144
	var v9148 int32
	_ = v9148
	var v9151 int32
	_ = v9151
	var v9152 int32
	_ = v9152
	var v9158 int32
	_ = v9158
	var v9161 int32
	_ = v9161
	var v9162 int32
	_ = v9162
	var v9163 int32
	_ = v9163
	var v9168 int32
	_ = v9168
	var v9172 int32
	_ = v9172
	var v9175 int32
	_ = v9175
	var v9176 int32
	_ = v9176
	var v9182 int32
	_ = v9182
	var v9183 int32
	_ = v9183
	var v9184 int32
	_ = v9184
	var v9185 int32
	_ = v9185
	var v9190 int32
	_ = v9190
	var v9194 int32
	_ = v9194
	var v9197 int32
	_ = v9197
	var v9198 int32
	_ = v9198
	var v9204 int32
	_ = v9204
	var v9205 int32
	_ = v9205
	var v9206 int32
	_ = v9206
	var v9207 int32
	_ = v9207
	var v9212 int32
	_ = v9212
	var v9217 int32
	_ = v9217
	var v9220 int32
	_ = v9220
	var v9221 int32
	_ = v9221
	var v9222 int32
	_ = v9222
	var v9223 int32
	_ = v9223
	var v9229 int32
	_ = v9229
	var v9230 int32
	_ = v9230
	var v9231 int32
	_ = v9231
	var v9232 int32
	_ = v9232
	var v9237 int32
	_ = v9237
	var v9242 int32
	_ = v9242
	var v9246 int32
	_ = v9246
	var v9251 int32
	_ = v9251
	var v9255 int32
	_ = v9255
	var v9258 int32
	_ = v9258
	var v9259 int32
	_ = v9259
	var v9260 int32
	_ = v9260
	var v9266 int32
	_ = v9266
	var v9267 int32
	_ = v9267
	var v9268 int32
	_ = v9268
	var v9269 int32
	_ = v9269
	var v9274 int32
	_ = v9274
	var v9278 int32
	_ = v9278
	var v9281 int32
	_ = v9281
	var v9285 int32
	_ = v9285
	var v9286 int32
	_ = v9286
	var v9287 int32
	_ = v9287
	var v9288 int32
	_ = v9288
	var v9293 int32
	_ = v9293
	var v9297 int32
	_ = v9297
	var v9300 int32
	_ = v9300
	var v9304 int32
	_ = v9304
	var v9307 int32
	_ = v9307
	var v9308 int32
	_ = v9308
	var v9309 int32
	_ = v9309
	var v9314 int32
	_ = v9314
	var v9318 int32
	_ = v9318
	var v9321 int32
	_ = v9321
	var v9325 int32
	_ = v9325
	var v9326 int32
	_ = v9326
	var v9327 int32
	_ = v9327
	var v9332 int32
	_ = v9332
	var v9336 int32
	_ = v9336
	var v9339 int32
	_ = v9339
	var v9343 int32
	_ = v9343
	var v9344 int32
	_ = v9344
	var v9345 int32
	_ = v9345
	var v9346 int32
	_ = v9346
	var v9351 int32
	_ = v9351
	var v9355 int32
	_ = v9355
	var v9358 int32
	_ = v9358
	var v9362 int32
	_ = v9362
	var v9365 int32
	_ = v9365
	var v9366 int32
	_ = v9366
	var v9367 int32
	_ = v9367
	var v9372 int32
	_ = v9372
	var v9376 int32
	_ = v9376
	var v9379 int32
	_ = v9379
	var v9380 int32
	_ = v9380
	var v9386 int32
	_ = v9386
	var v9389 int32
	_ = v9389
	var v9390 int32
	_ = v9390
	var v9391 int32
	_ = v9391
	var v9396 int32
	_ = v9396
	var v9400 int32
	_ = v9400
	var v9403 int32
	_ = v9403
	var v9406 int32
	_ = v9406
	var v9407 int32
	_ = v9407
	var v9410 int32
	_ = v9410
	var v9411 int32
	_ = v9411
	var v9412 int32
	_ = v9412
	var v9413 int32
	_ = v9413
	var v9418 int32
	_ = v9418
	var v9420 int32
	_ = v9420
	var v9422 int32
	_ = v9422
	var v9426 int32
	_ = v9426
	var v9429 int32
	_ = v9429
	var v9433 int32
	_ = v9433
	var v9436 int32
	_ = v9436
	var v9437 int32
	_ = v9437
	var v9438 int32
	_ = v9438
	var v9443 int32
	_ = v9443
	var v9447 int32
	_ = v9447
	var v9450 int32
	_ = v9450
	var v9454 int32
	_ = v9454
	var v9455 int32
	_ = v9455
	var v9456 int32
	_ = v9456
	var v9457 int32
	_ = v9457
	var v9462 int32
	_ = v9462
	var v9466 int32
	_ = v9466
	var v9469 int32
	_ = v9469
	var v9470 int32
	_ = v9470
	var v9471 int32
	_ = v9471
	var v9477 int32
	_ = v9477
	var v9478 int32
	_ = v9478
	var v9479 int32
	_ = v9479
	var v9480 int32
	_ = v9480
	var v9485 int32
	_ = v9485
	var v9490 int32
	_ = v9490
	var v9492 int32
	_ = v9492
	var v9494 int32
	_ = v9494
	var v9497 int32
	_ = v9497
	var v9499 int32
	_ = v9499
	var v9501 int32
	_ = v9501
	var v9508 int32
	_ = v9508
	var v9510 int32
	_ = v9510
	var v9524 int32
	_ = v9524
	var v9527 int32
	_ = v9527
	var v9530 int32
	_ = v9530
	var v9533 int32
	_ = v9533
	var v9538 int32
	_ = v9538
	var v9546 int32
	_ = v9546
	var v9547 int32
	_ = v9547
	var v9552 int32
	_ = v9552
	var v9555 int32
	_ = v9555
	var v9558 int64
	_ = v9558
	var v9560 int64
	_ = v9560
	var v9575 int32
	_ = v9575
	var v9577 int32
	_ = v9577
	var v9580 int32
	_ = v9580
	var v9593 int32
	_ = v9593
	var v9594 int32
	_ = v9594
	var v9595 int32
	_ = v9595
	var v9601 int32
	_ = v9601
	var v9604 int32
	_ = v9604
	var v9608 int32
	_ = v9608
	var v9609 int32
	_ = v9609
	var v9610 int32
	_ = v9610
	var v9615 int32
	_ = v9615
	var v9616 int32
	_ = v9616
	var v9618 int32
	_ = v9618
	var v9619 int32
	_ = v9619
	var v9621 int32
	_ = v9621
	var v9623 int32
	_ = v9623
	var v9625 int32
	_ = v9625
	var v9626 int32
	_ = v9626
	var v9632 int32
	_ = v9632
	var v9638 int32
	_ = v9638
	var v9640 int32
	_ = v9640
	var v9641 int32
	_ = v9641
	var v9642 int32
	_ = v9642
	var v9643 int32
	_ = v9643
	var v9647 int32
	_ = v9647
	var v9651 int32
	_ = v9651
	var v9655 int32
	_ = v9655
	var v9656 int32
	_ = v9656
	var v9657 int32
	_ = v9657
	var v9660 int32
	_ = v9660
	var v9663 int32
	_ = v9663
	var v9666 int32
	_ = v9666
	var v9669 int32
	_ = v9669
	var v9672 int32
	_ = v9672
	var v9674 int32
	_ = v9674
	var v9675 int32
	_ = v9675
	var v9677 int32
	_ = v9677
	var v9678 int32
	_ = v9678
	var v9680 int32
	_ = v9680
	var v9681 int32
	_ = v9681
	var v9685 int32
	_ = v9685
	var v9686 int32
	_ = v9686
	var v9688 int32
	_ = v9688
	var v9708 int32
	_ = v9708
	var v9723 int32
	_ = v9723
	var v9726 int32
	_ = v9726
	var v9727 int64
	_ = v9727
	var v9729 int64
	_ = v9729
	var v9733 int32
	_ = v9733
	var v9734 int32
	_ = v9734
	var v9736 int32
	_ = v9736
	var v9737 int32
	_ = v9737
	var v9740 int32
	_ = v9740
	var v9744 int32
	_ = v9744
	var v9747 int32
	_ = v9747
	var v9748 int32
	_ = v9748
	var v9752 int32
	_ = v9752
	var v9755 int32
	_ = v9755
	var v9761 int32
	_ = v9761
	var v9765 int32
	_ = v9765
	var v9771 int32
	_ = v9771
	var v9772 int32
	_ = v9772
	var v9777 int32
	_ = v9777
	var v9778 int32
	_ = v9778
	var v9779 int32
	_ = v9779
	var v9785 int32
	_ = v9785
	var v9796 int32
	_ = v9796
	var v9799 int32
	_ = v9799
	var v9803 int32
	_ = v9803
	var v9806 int32
	_ = v9806
	var v9807 int32
	_ = v9807
	var v9808 int32
	_ = v9808
	var v9813 int32
	_ = v9813
	var v9818 int32
	_ = v9818
	var v9828 int32
	_ = v9828
	var v9843 int32
	_ = v9843
	var v9851 int32
	_ = v9851
	var v9853 int32
	_ = v9853
	var v9854 int32
	_ = v9854
	var v9856 int32
	_ = v9856
	var v9866 int32
	_ = v9866
	var v9897 int32
	_ = v9897
	var v9903 int32
	_ = v9903
	var v9913 int32
	_ = v9913
	v3 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(_a_F_plpgsql_yyparse_0)
	m.G0 = v28
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1]))) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2]))) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+288)) = v3
	v42 = v28 + int32(288)
	v44 = v28 + int32(1088)
	v46 = v28 + int32(_a_F_plpgsql_yyparse_1)
	v53 = v3
	v54 = v44
	v59 = int32(-2)
	v60 = v42
	v61 = v46
	v62 = v46
	v65 = int32(200)
	v66 = v44
	v67 = v3
	v69 = v42
	goto L8
L1:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_3))
	mBase = m.M
	v9913 = m.ExcPending
	if v9913 != 0 {
		goto L46
	} else {
		goto L2552
	}
L2:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_4))
	mBase = m.M
	v9903 = m.ExcPending
	if v9903 != 0 {
		goto L46
	} else {
		goto L2551
	}
L3:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_5))
	mBase = m.M
	v9897 = m.ExcPending
	if v9897 != 0 {
		goto L46
	} else {
		goto L2550
	}
L4:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v9866 = m.ExcPending
	if v9866 != 0 {
		goto L46
	} else {
		goto L2549
	}
L5:
	;
	v9854 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_cword_is_not_variable(m, v148, v9854, l1)
	mBase = m.M
	v9856 = m.ExcPending
	if v9856 != 0 {
		goto L46
	} else {
		goto L2548
	}
L6:
	;
	v9851 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_word_is_not_variable(m, v148, v9851, l1)
	mBase = m.M
	v9853 = m.ExcPending
	if v9853 != 0 {
		goto L46
	} else {
		goto L2547
	}
L7:
	;
	if v28+int32(_a_F_plpgsql_yyparse_1) != v9828 {
		goto L2543
	} else {
		goto L2544
	}
L8:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v53)
	v75 = v65 << (uint(int32(1)) % 32)
	if base.Ui32(v61) < base.Ui32(v62+v75-int32(2)) {
		goto L39
	} else {
		goto L40
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9796 = m.ExcPending
	if v9796 != 0 {
		goto L46
	} else {
		goto L2538
	}
L10:
	;
	goto L9
L11:
	;
	v53 = v9771
	v54 = v9772
	v59 = v9777
	v60 = v9778
	v61 = v9779 + int32(2)
	v62 = v153
	v65 = v154
	v66 = v155
	v67 = v9785
	v69 = v156
	goto L8
L12:
	;
	v9723 = int32(0) - v239
	v9726 = v148 + v9723<<(uint(int32(4))%32)
	v9727 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v9726)+24)) = v9727
	v9729 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v9726)+16)) = v9729
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v254
	v9733 = v9726 + int32(16)
	v9734 = int32(1)
	v9736 = v152 + v9723<<(uint(v9734)%32)
	v9737 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9736))))
	v9740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+uint32(_c_F_plpgsql_yyparse[3]))))
	v9744 = (v9740 - int32(137)) << (uint(v9734) % 32)
	v9747 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9744)+uint32(_c_F_plpgsql_yyparse[4]))))
	v9748 = v9737 + v9747
	if base.Ui32(int32(1293)) < base.Ui32(v9748) {
		goto L2535
	} else {
		goto L2536
	}
L13:
	;
	v9619 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v9616)+4)) = v9619
	v9621 = *(*int32)(unsafe.Add(mBase, uint32(v9616)))
	v9623 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	F_check_sql_expr(m, v9621, v9619, v9623, l1)
	mBase = m.M
	v9625 = m.ExcPending
	if v9625 != 0 {
		goto L46
	} else {
		goto L2522
	}
L14:
	;
	v9577 = int32(1)
	v9580 = int32(0)
	v9593 = F_read_sql_construct(m, int32(269), int32(336), v9580, int32(_a_F_plpgsql_yyparse_8), v9580, v9577, v9580, v28+int32(240), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9594 = m.ExcPending
	if v9594 != 0 {
		goto L46
	} else {
		goto L2515
	}
L15:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_10))
	mBase = m.M
	v9575 = m.ExcPending
	if v9575 != 0 {
		goto L46
	} else {
		goto L2514
	}
L16:
	;
	v9497 = v164
	v9499 = v147
	v9501 = v148
	v9508 = v152
	v9510 = v9494
	goto L2503
L17:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), l0, l1, int32(_a_F_plpgsql_yyparse_5))
	mBase = m.M
	v9490 = m.ExcPending
	if v9490 != 0 {
		goto L46
	} else {
		goto L2502
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9466 = m.ExcPending
	if v9466 != 0 {
		goto L46
	} else {
		goto L2497
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9447 = m.ExcPending
	if v9447 != 0 {
		goto L46
	} else {
		goto L2492
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9426 = m.ExcPending
	if v9426 != 0 {
		goto L46
	} else {
		goto L2487
	}
L21:
	;
	F_cword_is_not_variable(m, v148, v7373, l1)
	mBase = m.M
	v9422 = m.ExcPending
	if v9422 != 0 {
		goto L46
	} else {
		goto L2486
	}
L22:
	;
	F_word_is_not_variable(m, v148, v7349, l1)
	mBase = m.M
	v9420 = m.ExcPending
	if v9420 != 0 {
		goto L46
	} else {
		goto L2485
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9400 = m.ExcPending
	if v9400 != 0 {
		goto L46
	} else {
		goto L2477
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9376 = m.ExcPending
	if v9376 != 0 {
		goto L46
	} else {
		goto L2472
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9355 = m.ExcPending
	if v9355 != 0 {
		goto L46
	} else {
		goto L2467
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9336 = m.ExcPending
	if v9336 != 0 {
		goto L46
	} else {
		goto L2462
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9318 = m.ExcPending
	if v9318 != 0 {
		goto L46
	} else {
		goto L2457
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9297 = m.ExcPending
	if v9297 != 0 {
		goto L46
	} else {
		goto L2452
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9278 = m.ExcPending
	if v9278 != 0 {
		goto L46
	} else {
		goto L2447
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9255 = m.ExcPending
	if v9255 != 0 {
		goto L46
	} else {
		goto L2441
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9242 = m.ExcPending
	if v9242 != 0 {
		goto L46
	} else {
		goto L2438
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9217 = m.ExcPending
	if v9217 != 0 {
		goto L46
	} else {
		goto L2432
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9194 = m.ExcPending
	if v9194 != 0 {
		goto L46
	} else {
		goto L2427
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9172 = m.ExcPending
	if v9172 != 0 {
		goto L46
	} else {
		goto L2422
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9148 = m.ExcPending
	if v9148 != 0 {
		goto L46
	} else {
		goto L2417
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9121 = m.ExcPending
	if v9121 != 0 {
		goto L46
	} else {
		goto L2411
	}
L37:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), l0, l1, int32(_a_F_plpgsql_yyparse_11))
	mBase = m.M
	v9117 = m.ExcPending
	if v9117 != 0 {
		goto L46
	} else {
		goto L2410
	}
L38:
	;
	if v53 == int32(3) {
		goto L65
	} else {
		goto L66
	}
L39:
	;
	v147 = v60
	v148 = v54
	v152 = v61
	v153 = v62
	v154 = v65
	v155 = v66
	v156 = v69
	goto L38
L40:
	;
	goto L41
L41:
	;
	if int32(_a_F_plpgsql_yyparse_12) < v65 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v82 = int32(_a_F_plpgsql_yyparse_13)
	if v82 <= v75 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v85 = v82
	goto L45
L44:
	;
	v85 = v75
	goto L45
L45:
	;
	v90 = F_palloc(m, v85*int32(22)+int32(30))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return int32(0)
L47:
	;
	if v90 == int32(0) {
		goto L37
	} else {
		goto L48
	}
L48:
	;
	v97 = int32(1)
	v100 = (v61-v62)>>(uint(v97)%32) + v97
	v102 = v100 << (uint(v97) % 32)
	if v102 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	base.MemoryCopy(m, v90, v62, v102)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v109 = base.I32_div_s(v85<<(uint(int32(1))%32)+int32(15), int32(16))
	v110 = int32(4)
	v112 = v90 + v109<<(uint(v110)%32)
	v114 = v100 << (uint(v110) % 32)
	if v114 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	base.MemoryCopy(m, v112, v66, v114)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v116 = int32(4)
	v121 = base.I32_div_s(v85<<(uint(v116)%32)|int32(15), int32(16))
	v124 = v112 + v121<<(uint(v116)%32)
	v126 = v100 << (uint(int32(2)) % 32)
	if v126 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, v124, v69, v126)
	goto L57
L56:
	;
	goto L57
L57:
	;
	if v28+int32(_a_F_plpgsql_yyparse_1) != v62 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_pfree(m, v62)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L46
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v85 <= v100 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v9818 = int32(1)
	v9828 = v90
	goto L7
L63:
	;
	goto L64
L64:
	;
	v147 = v126 + v124 - int32(4)
	v148 = v114 + v112 - int32(16)
	v152 = v90 + v100<<(uint(int32(1))%32) - int32(2)
	v153 = v90
	v154 = v85
	v155 = v112
	v156 = v124
	goto L38
L65:
	;
	v9818 = int32(0)
	v9828 = v153
	goto L7
L66:
	;
	goto L67
L67:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[5]))))
	if v164 == int32(_a_F_plpgsql_yyparse_14) {
		v224 = v59
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v67 == int32(0) {
		goto L17
	} else {
		goto L2404
	}
L69:
	;
	v239 = int32(*(*int8)(unsafe.Add(mBase, uint32(v235)+uint32(_c_F_plpgsql_yyparse[6]))))
	v241 = int32(4)
	v243 = v148 + (int32(1)-v239)<<(uint(v241)%32)
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+280)) = v244
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v243)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+272)) = v246
	v252 = v147 - v239<<(uint(int32(2))%32) + v241
	if v239 != 0 {
		goto L90
	} else {
		goto L91
	}
L70:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_plpgsql_yyparse[7]))))
	if v228 == int32(0) {
		goto L68
	} else {
		goto L89
	}
L71:
	;
	if v59 == int32(-2) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v193 = v191 + base.I32_extend16_s(v164)
	if base.Ui32(int32(1293)) < base.Ui32(v193) {
		v224 = v190
		goto L70
	} else {
		goto L84
	}
L73:
	;
	v173 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L46
	} else {
		goto L76
	}
L74:
	;
	v175 = v59
	goto L75
L75:
	;
	if v175 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v175 = v173
	goto L75
L77:
	;
	v178 = int32(0)
	v190 = v178
	v191 = v178
	goto L72
L78:
	;
	goto L79
L79:
	;
	if v175 == int32(256) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v9492 = int32(257)
	v9494 = v183
	goto L16
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(385)) < base.Ui32(v175) {
		v190 = v175
		v191 = int32(2)
		goto L72
	} else {
		goto L83
	}
L83:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+uint32(_c_F_plpgsql_yyparse[8]))))
	v190 = v175
	v191 = v189
	goto L72
L84:
	;
	v197 = v193 << (uint(int32(1)) % 32)
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_plpgsql_yyparse[9]))))
	if v191 != v200 {
		v224 = v190
		goto L70
	} else {
		goto L85
	}
L85:
	;
	v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_plpgsql_yyparse[10]))))
	if int32(0) < v204 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+24)) = v207
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+16)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v211
	v9771 = v204
	v9772 = v148 + int32(16)
	v9777 = int32(-2)
	v9778 = v147 + int32(4)
	v9779 = v152
	v9785 = v67 - base.B2i32(v67 != int32(0))
	goto L11
L87:
	;
	goto L88
L88:
	;
	v233 = v190
	v235 = int32(0) - v204
	goto L69
L89:
	;
	v233 = v224
	v235 = v228
	goto L69
L90:
	;
	v253 = v252
	goto L92
L91:
	;
	v253 = v147
	goto L92
L92:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	switch v235 - int32(2) {
	case 0:
		goto L243
	default:
		v9708 = v233
		goto L12
	case 3:
		goto L242
	case 4:
		goto L241
	case 5:
		goto L240
	case 6:
		goto L239
	case 7:
		goto L238
	case 8:
		goto L237
	case 9:
		goto L236
	case 12:
		goto L235
	case 13:
		goto L234
	case 14:
		goto L233
	case 15:
		goto L232
	case 16:
		goto L231
	case 21:
		goto L230
	case 22:
		goto L229
	case 23:
		goto L228
	case 24:
		goto L227
	case 25:
		goto L226
	case 26:
		goto L225
	case 27:
		goto L224
	case 28:
		goto L223
	case 29:
		goto L222
	case 30:
		goto L221
	case 31:
		goto L220
	case 32:
		goto L219
	case 33:
		goto L218
	case 34:
		goto L217
	case 37:
		goto L216
	case 38:
		goto L215
	case 39:
		goto L214
	case 40:
		goto L213
	case 41:
		goto L212
	case 42:
		goto L211
	case 43:
		goto L210
	case 44:
		goto L209
	case 45:
		goto L208
	case 46:
		goto L207
	case 47:
		goto L206
	case 48:
		goto L205
	case 49:
		goto L204
	case 50:
		goto L203
	case 51:
		goto L202
	case 52:
		goto L201
	case 57:
		goto L200
	case 58:
		goto L199
	case 59:
		goto L198
	case 60:
		goto L197
	case 61:
		goto L196
	case 62:
		goto L195
	case 63:
		goto L194
	case 64:
		goto L193
	case 65:
		goto L192
	case 66:
		goto L191
	case 67:
		goto L190
	case 68:
		goto L189
	case 69:
		goto L188
	case 70:
		goto L187
	case 71:
		goto L186
	case 72:
		goto L185
	case 73:
		goto L184
	case 74:
		goto L183
	case 75:
		goto L182
	case 76:
		goto L181
	case 77:
		goto L180
	case 78:
		goto L179
	case 79:
		goto L178
	case 80:
		goto L177
	case 81:
		goto L176
	case 82:
		goto L175
	case 83:
		goto L174
	case 84:
		goto L173
	case 85:
		goto L172
	case 86:
		goto L171
	case 87:
		goto L170
	case 88:
		goto L169
	case 89:
		goto L168
	case 90:
		goto L167
	case 91:
		goto L166
	case 92:
		goto L165
	case 93:
		goto L164
	case 94:
		goto L163
	case 95:
		goto L162
	case 96, 145:
		goto L6
	case 97, 116, 146:
		goto L5
	case 98:
		goto L161
	case 99:
		goto L160
	case 100:
		goto L159
	case 101:
		goto L158
	case 102:
		goto L157
	case 103:
		goto L156
	case 104:
		goto L155
	case 105:
		goto L154
	case 106:
		goto L153
	case 107:
		goto L152
	case 108:
		goto L151
	case 109:
		goto L150
	case 110:
		goto L149
	case 111:
		goto L148
	case 112:
		goto L147
	case 113:
		goto L146
	case 114:
		goto L145
	case 115:
		goto L144
	case 117:
		goto L143
	case 118:
		goto L142
	case 119:
		goto L141
	case 120:
		goto L140
	case 121:
		goto L139
	case 122:
		goto L138
	case 123:
		goto L137
	case 124:
		goto L136
	case 125:
		goto L135
	case 126:
		goto L134
	case 127:
		goto L133
	case 128:
		goto L132
	case 129:
		goto L131
	case 130:
		goto L130
	case 131:
		goto L129
	case 132:
		goto L128
	case 133:
		goto L127
	case 134:
		goto L126
	case 135:
		goto L125
	case 136:
		goto L124
	case 137:
		goto L123
	case 138:
		goto L122
	case 139:
		goto L121
	case 140:
		goto L120
	case 141:
		goto L119
	case 142:
		goto L118
	case 143:
		goto L117
	case 144:
		goto L116
	case 147:
		goto L115
	case 148:
		goto L114
	case 149:
		goto L113
	case 150:
		goto L112
	case 151:
		goto L111
	case 152:
		goto L110
	case 153:
		goto L109
	case 154:
		goto L108
	case 155:
		goto L107
	case 156:
		goto L106
	case 157:
		goto L105
	case 158:
		goto L104
	case 159:
		goto L103
	case 160:
		goto L102
	case 161:
		goto L101
	case 162:
		goto L100
	case 163:
		goto L99
	case 164:
		goto L98
	case 165:
		goto L97
	case 166:
		goto L96
	case 167:
		goto L95
	case 168:
		goto L94
	case 169:
		goto L93
	}
L93:
	;
	v9098 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v9098 == int32(0) {
		goto L3
	} else {
		goto L2403
	}
L94:
	;
	v9094 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v9095 = F_pstrdup(m, v9094)
	mBase = m.M
	v9096 = m.ExcPending
	if v9096 != 0 {
		goto L46
	} else {
		goto L2402
	}
L95:
	;
	v9092 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9092
	v9708 = v233
	goto L12
L96:
	;
	v9090 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9090
	v9708 = v233
	goto L12
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L98:
	;
	v9086 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9086
	v9708 = v233
	goto L12
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L100:
	;
	v9077 = v148 - int32(16)
	v9078 = *(*int32)(unsafe.Add(mBase, uint32(v9077)))
	F_plpgsql_ns_push(m, v9078, int32(1))
	mBase = m.M
	v9081 = m.ExcPending
	if v9081 != 0 {
		goto L46
	} else {
		goto L2401
	}
L101:
	;
	F_plpgsql_ns_push(m, int32(0), int32(1))
	mBase = m.M
	v9073 = m.ExcPending
	if v9073 != 0 {
		goto L46
	} else {
		goto L2400
	}
L102:
	;
	v9063 = v148 - int32(16)
	v9064 = *(*int32)(unsafe.Add(mBase, uint32(v9063)))
	F_plpgsql_ns_push(m, v9064, int32(0))
	mBase = m.M
	v9067 = m.ExcPending
	if v9067 != 0 {
		goto L46
	} else {
		goto L2399
	}
L103:
	;
	v9056 = int32(0)
	F_plpgsql_ns_push(m, v9056, v9056)
	mBase = m.M
	v9059 = m.ExcPending
	if v9059 != 0 {
		goto L46
	} else {
		goto L2398
	}
L104:
	;
	v9041 = int32(0)
	v9045 = int32(1)
	v9053 = F_read_sql_construct(m, int32(336), v9041, v9041, int32(_a_F_plpgsql_yyparse_8), int32(2), v9045, v9045, v9041, v9041, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9054 = m.ExcPending
	if v9054 != 0 {
		goto L46
	} else {
		goto L2397
	}
L105:
	;
	v9025 = int32(0)
	v9029 = int32(1)
	v9037 = F_read_sql_construct(m, int32(376), v9025, v9025, int32(_a_F_plpgsql_yyparse_15), int32(2), v9029, v9029, v9025, v9025, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9038 = m.ExcPending
	if v9038 != 0 {
		goto L46
	} else {
		goto L2396
	}
L106:
	;
	v9009 = int32(0)
	v9013 = int32(1)
	v9021 = F_read_sql_construct(m, int32(59), v9009, v9009, int32(_a_F_plpgsql_yyparse_16), int32(2), v9013, v9013, v9009, v9009, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9022 = m.ExcPending
	if v9022 != 0 {
		goto L46
	} else {
		goto L2395
	}
L107:
	;
	v8834 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8835 = int32(_a_F_plpgsql_yyparse_17)
	v8838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8834))))
	v8841 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[11])))
	if base.B2i32(v8838 == int32(0))|base.B2i32(v8838 != v8841) != 0 {
		v8859 = v8838
		v8860 = v8841
		goto L2361
	} else {
		goto L2362
	}
L108:
	;
	v8832 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8832
	v9708 = v233
	goto L12
L109:
	;
	v8800 = v148 - int32(32)
	v8801 = *(*int32)(unsafe.Add(mBase, uint32(v8800)))
	v8806 = v8801
	goto L2357
L110:
	;
	v8739 = F_palloc0(m, int32(12))
	mBase = m.M
	v8740 = m.ExcPending
	if v8740 != 0 {
		goto L46
	} else {
		goto L2343
	}
L111:
	;
	v8729 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+232)) = v8729
	*(*int32)(unsafe.Add(mBase, uint32(v28)+236)) = v8729
	v8735 = F_list_make1_impl(m, int32(1), v28+int32(232))
	mBase = m.M
	v8736 = m.ExcPending
	if v8736 != 0 {
		goto L46
	} else {
		goto L2342
	}
L112:
	;
	v8724 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8725 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8726 = F_lappend(m, v8724, v8725)
	mBase = m.M
	v8727 = m.ExcPending
	if v8727 != 0 {
		goto L46
	} else {
		goto L2341
	}
L113:
	;
	v8718 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8719 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v8718)+8)) = v8719
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8718
	v9708 = v233
	goto L12
L114:
	;
	v8630 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v8631 = int32(0)
	if v8630 < v8631 {
		v8676 = v8631
		goto L2324
	} else {
		goto L2325
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L116:
	;
	v8616 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8617 = *(*int32)(unsafe.Add(mBase, uint32(v8616)))
	if v8617 != 0 {
		goto L19
	} else {
		goto L2319
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(1)
	v9708 = v233
	goto L12
L120:
	;
	v8541 = F_palloc(m, int32(16))
	mBase = m.M
	v8542 = m.ExcPending
	if v8542 != 0 {
		goto L46
	} else {
		goto L2305
	}
L121:
	;
	v8471 = F_palloc(m, int32(16))
	mBase = m.M
	v8472 = m.ExcPending
	if v8472 != 0 {
		goto L46
	} else {
		goto L2291
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L123:
	;
	v8400 = F_palloc(m, int32(16))
	mBase = m.M
	v8401 = m.ExcPending
	if v8401 != 0 {
		goto L46
	} else {
		goto L2277
	}
L124:
	;
	v8023 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v8025 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v8026 = m.G0
	v8028 = v8026 - int32(16)
	m.G0 = v8028
	v8030 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8028)+15)) = uint8(v8030)
	v8033 = F_palloc0(m, int32(36))
	mBase = m.M
	v8034 = m.ExcPending
	if v8034 != 0 {
		goto L46
	} else {
		goto L2152
	}
L125:
	;
	v7962 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v7965 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v7966 = int32(0)
	if v7965 < v7966 {
		v8011 = v7966
		goto L2140
	} else {
		goto L2141
	}
L126:
	;
	v7882 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v7887 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v7889 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_read_into_target(m, v28+int32(256), int32(0), v7887, v7889, l1)
	mBase = m.M
	v7891 = m.ExcPending
	if v7891 != 0 {
		goto L46
	} else {
		goto L2122
	}
L127:
	;
	v7573 = F_palloc0(m, int32(36))
	mBase = m.M
	v7574 = m.ExcPending
	if v7574 != 0 {
		goto L46
	} else {
		goto L2050
	}
L128:
	;
	v7395 = int32(1)
	v7404 = F_read_sql_construct(m, int32(332), int32(381), int32(59), int32(_a_F_plpgsql_yyparse_18), int32(2), v7395, v7395, int32(0), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7405 = m.ExcPending
	if v7405 != 0 {
		goto L46
	} else {
		goto L2017
	}
L129:
	;
	v7366 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v7368 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v7369 = F_plpgsql_yylex(m, v7366, v7368, l1)
	mBase = m.M
	v7370 = m.ExcPending
	if v7370 != 0 {
		goto L46
	} else {
		goto L2011
	}
L130:
	;
	v7342 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v7344 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v7345 = F_plpgsql_yylex(m, v7342, v7344, l1)
	mBase = m.M
	v7346 = m.ExcPending
	if v7346 != 0 {
		goto L46
	} else {
		goto L2005
	}
L131:
	;
	v7332 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7338 = F_make_execsql_stmt(m, int32(337), v7332, int32(0), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7339 = m.ExcPending
	if v7339 != 0 {
		goto L46
	} else {
		goto L2004
	}
L132:
	;
	v7322 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7328 = F_make_execsql_stmt(m, int32(331), v7322, int32(0), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7329 = m.ExcPending
	if v7329 != 0 {
		goto L46
	} else {
		goto L2003
	}
L133:
	;
	v7312 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7318 = F_make_execsql_stmt(m, int32(328), v7312, int32(0), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7319 = m.ExcPending
	if v7319 != 0 {
		goto L46
	} else {
		goto L2002
	}
L134:
	;
	v7301 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7301
	v7305 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v7305
	v7309 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+280)) = v7309
	v9708 = v233
	goto L12
L135:
	;
	v7204 = F_palloc(m, int32(20))
	mBase = m.M
	v7205 = m.ExcPending
	if v7205 != 0 {
		goto L46
	} else {
		goto L1983
	}
L136:
	;
	v6085 = F_palloc(m, int32(32))
	mBase = m.M
	v6086 = m.ExcPending
	if v6086 != 0 {
		goto L46
	} else {
		goto L1675
	}
L137:
	;
	v5452 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v5453 = m.ExcPending
	if v5453 != 0 {
		goto L46
	} else {
		goto L1511
	}
L138:
	;
	v5446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v5446)
	v9708 = v233
	goto L12
L139:
	;
	v5444 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v5444)
	v9708 = v233
	goto L12
L140:
	;
	v5311 = F_palloc0(m, int32(24))
	mBase = m.M
	v5312 = m.ExcPending
	if v5312 != 0 {
		goto L46
	} else {
		goto L1459
	}
L141:
	;
	v5308 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5308
	v9708 = v233
	goto L12
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L143:
	;
	v5193 = F_palloc0(m, int32(32))
	mBase = m.M
	v5194 = m.ExcPending
	if v5194 != 0 {
		goto L46
	} else {
		goto L1432
	}
L144:
	;
	v5129 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5129
	v5131 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5132 = int32(0)
	if v5131 < v5132 {
		v5177 = v5132
		goto L1417
	} else {
		goto L1418
	}
L145:
	;
	v5045 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v5045 != 0 {
		goto L1392
	} else {
		goto L1393
	}
L146:
	;
	v4752 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v4754 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v4755 = F_plpgsql_yylex(m, v4752, v4754, l1)
	mBase = m.M
	v4756 = m.ExcPending
	if v4756 != 0 {
		goto L46
	} else {
		goto L1334
	}
L147:
	;
	v4669 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4670 = *(*int32)(unsafe.Add(mBase, uint32(v4669)))
	v4673 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4674 = int32(0)
	if v4673 < v4674 {
		v4719 = v4674
		goto L1311
	} else {
		goto L1312
	}
L148:
	;
	v4578 = F_palloc0(m, int32(24))
	mBase = m.M
	v4579 = m.ExcPending
	if v4579 != 0 {
		goto L46
	} else {
		goto L1288
	}
L149:
	;
	v4492 = F_palloc0(m, int32(20))
	mBase = m.M
	v4493 = m.ExcPending
	if v4493 != 0 {
		goto L46
	} else {
		goto L1266
	}
L150:
	;
	v4479 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v4479 != 0 {
		goto L1262
	} else {
		goto L1263
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L152:
	;
	v4417 = F_palloc(m, int32(12))
	mBase = m.M
	v4418 = m.ExcPending
	if v4418 != 0 {
		goto L46
	} else {
		goto L1248
	}
L153:
	;
	v4407 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+184)) = v4407
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = v4407
	v4413 = F_list_make1_impl(m, int32(1), v28+int32(184))
	mBase = m.M
	v4414 = m.ExcPending
	if v4414 != 0 {
		goto L46
	} else {
		goto L1247
	}
L154:
	;
	v4402 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4403 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v4404 = F_lappend(m, v4402, v4403)
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
		goto L46
	} else {
		goto L1246
	}
L155:
	;
	v4371 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v4373 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v4374 = F_plpgsql_yylex(m, v4371, v4373, l1)
	mBase = m.M
	v4375 = m.ExcPending
	if v4375 != 0 {
		goto L46
	} else {
		goto L1239
	}
L156:
	;
	v4158 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(24))))
	v4159 = int32(80)
	v4161 = *(*int32)(unsafe.Add(mBase, uint32(v148-v4159)))
	v4164 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	v4168 = m.G0
	v4170 = v4168 - v4159
	m.G0 = v4170
	v4173 = F_palloc(m, int32(32))
	mBase = m.M
	v4174 = m.ExcPending
	if v4174 != 0 {
		goto L46
	} else {
		goto L1206
	}
L157:
	;
	v4154 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4154
	v9708 = v233
	goto L12
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L159:
	;
	v4087 = F_palloc0(m, int32(12))
	mBase = m.M
	v4088 = m.ExcPending
	if v4088 != 0 {
		goto L46
	} else {
		goto L1191
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L161:
	;
	v4005 = F_palloc0(m, int32(28))
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L46
	} else {
		goto L1177
	}
L162:
	;
	v3988 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3989 = *(*int32)(unsafe.Add(mBase, uint32(v3988)))
	if base.Ui32(v3989-int32(1)) < base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L1173
	}
L163:
	;
	v3563 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L46
	} else {
		goto L1059
	}
L164:
	;
	v3549 = F_palloc(m, int32(8))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		goto L46
	} else {
		goto L1031
	}
L165:
	;
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+172)) = v3539
	*(*int32)(unsafe.Add(mBase, uint32(v28)+252)) = v3539
	v3545 = F_list_make1_impl(m, int32(1), v28+int32(172))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L46
	} else {
		goto L1030
	}
L166:
	;
	v3534 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v3535 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3536 = F_lappend(m, v3534, v3535)
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
		goto L46
	} else {
		goto L1029
	}
L167:
	;
	v3530 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v3530)
	v9708 = v233
	goto L12
L168:
	;
	v3528 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v3528)
	v9708 = v233
	goto L12
L169:
	;
	v3526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v3526)
	v9708 = v233
	goto L12
L170:
	;
	v3298 = F_palloc0(m, int32(20))
	mBase = m.M
	v3299 = m.ExcPending
	if v3299 != 0 {
		goto L46
	} else {
		goto L977
	}
L171:
	;
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v3185 == int32(0) {
		goto L952
	} else {
		goto L953
	}
L172:
	;
	v3107 = F_palloc0(m, int32(24))
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L46
	} else {
		goto L936
	}
L173:
	;
	v3029 = F_palloc0(m, int32(24))
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		goto L46
	} else {
		goto L920
	}
L174:
	;
	v2926 = F_palloc0(m, int32(16))
	mBase = m.M
	v2927 = m.ExcPending
	if v2927 != 0 {
		goto L46
	} else {
		goto L900
	}
L175:
	;
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2923
	v9708 = v233
	goto L12
L176:
	;
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2921
	v9708 = v233
	goto L12
L177:
	;
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2919
	v9708 = v233
	goto L12
L178:
	;
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2917
	v9708 = v233
	goto L12
L179:
	;
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2915
	v9708 = v233
	goto L12
L180:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2913
	v9708 = v233
	goto L12
L181:
	;
	v2911 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2911
	v9708 = v233
	goto L12
L182:
	;
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2909
	v9708 = v233
	goto L12
L183:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2907
	v9708 = v233
	goto L12
L184:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2905
	v9708 = v233
	goto L12
L185:
	;
	v2903 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2903
	v9708 = v233
	goto L12
L186:
	;
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2901
	v9708 = v233
	goto L12
L187:
	;
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2899
	v9708 = v233
	goto L12
L188:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2897
	v9708 = v233
	goto L12
L189:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2895
	v9708 = v233
	goto L12
L190:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2893
	v9708 = v233
	goto L12
L191:
	;
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2891
	v9708 = v233
	goto L12
L192:
	;
	v2889 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2889
	v9708 = v233
	goto L12
L193:
	;
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2887
	v9708 = v233
	goto L12
L194:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2885
	v9708 = v233
	goto L12
L195:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2883
	v9708 = v233
	goto L12
L196:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2881
	v9708 = v233
	goto L12
L197:
	;
	v2879 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2879
	v9708 = v233
	goto L12
L198:
	;
	v2877 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2877
	v9708 = v233
	goto L12
L199:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v2868 == int32(0) {
		goto L896
	} else {
		goto L897
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L201:
	;
	v2848 = int32(0)
	v2852 = int32(1)
	v2860 = F_read_sql_construct(m, int32(59), v2848, v2848, int32(_a_F_plpgsql_yyparse_16), int32(2), v2852, v2852, v2848, v2848, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v2861 = m.ExcPending
	if v2861 != 0 {
		goto L46
	} else {
		goto L895
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L203:
	;
	v2843 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2843)
	v9708 = v233
	goto L12
L204:
	;
	v2841 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2841)
	v9708 = v233
	goto L12
L205:
	;
	v2836 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v2838 = F_get_collation_oid(m, v2836, int32(0))
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L46
	} else {
		goto L894
	}
L206:
	;
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v2821 = F_pstrdup(m, v2820)
	mBase = m.M
	v2822 = m.ExcPending
	if v2822 != 0 {
		goto L46
	} else {
		goto L890
	}
L207:
	;
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v2807 = F_makeString(m, v2806)
	mBase = m.M
	v2808 = m.ExcPending
	if v2808 != 0 {
		goto L46
	} else {
		goto L887
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L209:
	;
	v2264 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v2266 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v2267 = m.G0
	v2269 = v2267 - int32(48)
	m.G0 = v2269
	if v233 == int32(-2) {
		goto L696
	} else {
		goto L697
	}
L210:
	;
	v2261 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2261)
	v9708 = v233
	goto L12
L211:
	;
	v2259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2259)
	v9708 = v233
	goto L12
L212:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v1859 = F_pstrdup(m, v1858)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L46
	} else {
		goto L598
	}
L213:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1459
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v1462 = int32(0)
	if v1461 < v1462 {
		v1507 = v1462
		goto L502
	} else {
		goto L503
	}
L214:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1135 == int32(0) {
		goto L32
	} else {
		goto L424
	}
L215:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v981 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L216:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v826 == int32(0) {
		goto L354
	} else {
		goto L355
	}
L217:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v822 = F_plpgsql_build_variable(m, v816, v819, v820, int32(1))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L46
	} else {
		goto L349
	}
L218:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v811 = F_lappend(m, v809, v810)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L46
	} else {
		goto L348
	}
L219:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v798
	*(*int32)(unsafe.Add(mBase, uint32(v28)+268)) = v798
	v804 = F_list_make1_impl(m, int32(1), v28+int32(44))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L46
	} else {
		goto L347
	}
L220:
	;
	v617 = F_palloc0(m, int32(40))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L46
	} else {
		goto L317
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L222:
	;
	v611 = F_read_sql_stmt(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L46
	} else {
		goto L316
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(2)
	v9708 = v233
	goto L12
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(4)
	v9708 = v233
	goto L12
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L226:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	if v563 != 0 {
		goto L305
	} else {
		goto L306
	}
L227:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	F_plpgsql_ns_push(m, v556, int32(2))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L46
	} else {
		goto L303
	}
L228:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	F_plpgsql_ns_additem(m, v547, v548, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L46
	} else {
		goto L302
	}
L229:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	if v496 != 0 {
		goto L291
	} else {
		goto L292
	}
L230:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L46
	} else {
		goto L286
	}
L231:
	;
	v468 = F_plpgsql_add_initdatums(m, int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L46
	} else {
		goto L285
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v462
	v464 = F_plpgsql_add_initdatums(m, v28+int32(280))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L46
	} else {
		goto L284
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+276)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v453
	v9708 = v233
	goto L12
L234:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = int32(0)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+276)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v444
	v9708 = v233
	goto L12
L235:
	;
	v342 = F_palloc0(m, int32(32))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L46
	} else {
		goto L262
	}
L236:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v338 = F_pstrdup(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L46
	} else {
		goto L261
	}
L237:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v335
	v9708 = v233
	goto L12
L238:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+488)) = int32(2)
	v9708 = v233
	goto L12
L239:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+488)) = int32(1)
	v9708 = v233
	goto L12
L240:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+488)) = int32(0)
	v9708 = v233
	goto L12
L241:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v265 != int32(111) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	v262 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15])) = uint8(v262)
	v9708 = v233
	goto L12
L243:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v259
	v9708 = v233
	goto L12
L244:
	;
	v276 = int32(_a_F_plpgsql_yyparse_19)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[16])))
	if base.B2i32(v279 == int32(0))|base.B2i32(v279 != v282) != 0 {
		v300 = v279
		v301 = v282
		goto L249
	} else {
		goto L250
	}
L245:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)))
	if v268 != int32(110) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+2)))
	if v271 != 0 {
		goto L244
	} else {
		goto L247
	}
L247:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v273)+492)) = uint8(v274)
	v9708 = v233
	goto L12
L248:
	;
	if v300-v301 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L249:
	;
	goto L248
L250:
	;
	v285 = v264
	v286 = v276
	goto L251
L251:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	if v290 == int32(0) {
		v300 = v290
		v301 = v289
		goto L249
	} else {
		goto L253
	}
L252:
	;
	v300 = v290
	v301 = v289
	goto L249
L253:
	;
	v293 = int32(1)
	if v290 == v289 {
		v285 = v285 + v293
		v286 = v286 + v293
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+492)) = uint8(v307)
	v9708 = v233
	goto L12
L256:
	;
	goto L257
L257:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L46
	} else {
		goto L258
	}
L258:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v313
	F_errmsg_internal(m, int32(_a_F_plpgsql_yyparse_20), v28)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L46
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(396), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L46
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
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v338
	v9708 = v233
	goto L12
L262:
	;
	v344 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v344
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(16))))
	if v348 < v344 {
		v394 = v344
		goto L264
	} else {
		goto L265
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+4)) = v394
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+520))
	v401 = v399 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v398)+520)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v342)+8)) = v401
	v405 = v148 - int32(80)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+12)) = v406
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(76))))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+20)) = v410
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(72))))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+24)) = v414
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+16)) = v418
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+28)) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_check_labels(m, v424, v425, v426, l1)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L46
	} else {
		goto L276
	}
L264:
	;
	goto L263
L265:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+60))
	if v355 == int32(0) {
		v394 = v344
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v358 = v348 + v355
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v354)+188))
	if base.Ui32(v359) <= base.Ui32(v358) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v354)+196))
	if base.B2i32(v368 == int32(0))|base.B2i32(base.Ui32(v358) <= base.Ui32(v368)) != 0 {
		v394 = v369
		goto L264
	} else {
		goto L271
	}
L268:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v354)+192))
	v368 = v361
	goto L267
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v354)+188)) = v355
	v366 = F_strchr(m, v355, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v354)+192)) = v366
	v368 = v366
	goto L267
L271:
	;
	v374 = v368
	v377 = v369
	goto L272
L272:
	;
	v379 = int32(1)
	v380 = v377 + v379
	*(*int32)(unsafe.Add(mBase, uint32(v354)+196)) = v380
	v383 = v374 + v379
	*(*int32)(unsafe.Add(mBase, uint32(v354)+188)) = v383
	v386 = F_strchr(m, v383, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v354)+192)) = v386
	if v386 == int32(0) {
		v394 = v380
		goto L264
	} else {
		goto L274
	}
L273:
	;
	v394 = v380
	goto L264
L274:
	;
	if base.Ui32(v386) < base.Ui32(v358) {
		v374 = v386
		v377 = v380
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	if v432 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v342
	v9708 = v233
	goto L12
L278:
	;
	v433 = v431
	goto L281
L279:
	;
	v436 = v431
	goto L280
L280:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])) = v438
	goto L277
L281:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+8))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	if v435 != 0 {
		v433 = v434
		goto L281
	} else {
		goto L283
	}
L282:
	;
	v436 = v434
	goto L280
L283:
	;
	goto L282
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v464
	v9708 = v233
	goto L12
L285:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = int32(1)
	v9708 = v233
	goto L12
L286:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L46
	} else {
		goto L287
	}
L287:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_23), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L46
	} else {
		goto L288
	}
L288:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v487 = F_plpgsql_scanner_errposition(m, v486, l1)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L46
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(502), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L46
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
	v498 = v148 - int32(48)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+16))
	if v500 == int32(0) {
		goto L36
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(80))))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(76))))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	v516 = F_plpgsql_build_variable(m, v508, v511, v514, int32(1))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L46
	} else {
		goto L295
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499)+16)) = v496
	goto L293
L295:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+int32(-64)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v516)+16)) = uint8(v520)
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148-int32(16)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v516)+17)) = uint8(v524)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v516)+20)) = v526
	if base.B2i32(v526 == int32(0))&base.B2i32(v524 == int32(1)) != 0 {
		goto L35
	} else {
		goto L296
	}
L296:
	;
	if v526 == int32(0) {
		v9708 = v233
		goto L12
	} else {
		goto L297
	}
L297:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	if v535 != 0 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+20)) = uint8(v540)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+16)) = v541
	v9708 = v233
	goto L12
L299:
	;
	v540 = int32(0)
	v541 = int32(-1)
	goto L298
L300:
	;
	goto L301
L301:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v540 = int32(1)
	v541 = v539
	goto L298
L302:
	;
	v9708 = v233
	goto L12
L303:
	;
	v9708 = v233
	goto L12
L304:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(96))))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(92))))
	v579 = int32(0)
	v581 = F_plpgsql_build_datatype(m, int32(1790), int32(-1), v579, v579)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L46
	} else {
		goto L311
	}
L305:
	;
	v564 = v562
	goto L308
L306:
	;
	v567 = v562
	goto L307
L307:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])) = v569
	goto L304
L308:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+8))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	if v566 != 0 {
		v564 = v565
		goto L308
	} else {
		goto L310
	}
L309:
	;
	v567 = v565
	goto L307
L310:
	;
	goto L309
L311:
	;
	v584 = F_plpgsql_build_variable(m, v573, v576, v581, int32(1))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L46
	} else {
		goto L312
	}
L312:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+28)) = v586
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	if v590 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	v593 = v591
	goto L315
L314:
	;
	v593 = int32(-1)
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584)+32)) = v593
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(80))))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+36)) = v597 | int32(256)
	v9708 = v233
	goto L12
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v611
	v9708 = v233
	goto L12
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+8)) = int32(_a_F_plpgsql_yyparse_24)
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = int32(1)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v626 = int32(0)
	if v625 < v626 {
		v671 = v626
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v673 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v617)+24)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v617)+12)) = v671
	v678 = v148 - int32(16)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	if v679 != 0 {
		goto L331
	} else {
		goto L332
	}
L319:
	;
	goto L318
L320:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+60))
	if v632 == int32(0) {
		v671 = v626
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v635 = v625 + v632
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v631)+188))
	if base.Ui32(v636) <= base.Ui32(v635) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v631)+196))
	if base.B2i32(v645 == int32(0))|base.B2i32(base.Ui32(v635) <= base.Ui32(v645)) != 0 {
		v671 = v646
		goto L319
	} else {
		goto L326
	}
L323:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v631)+192))
	v645 = v638
	goto L322
L324:
	;
	goto L325
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v631)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v631)+188)) = v632
	v643 = F_strchr(m, v632, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v631)+192)) = v643
	v645 = v643
	goto L322
L326:
	;
	v651 = v645
	v654 = v646
	goto L327
L327:
	;
	v656 = int32(1)
	v657 = v654 + v656
	*(*int32)(unsafe.Add(mBase, uint32(v631)+196)) = v657
	v660 = v651 + v656
	*(*int32)(unsafe.Add(mBase, uint32(v631)+188)) = v660
	v663 = F_strchr(m, v660, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v631)+192)) = v663
	if v663 == int32(0) {
		v671 = v657
		goto L319
	} else {
		goto L329
	}
L328:
	;
	v671 = v657
	goto L319
L329:
	;
	if base.Ui32(v663) < base.Ui32(v635) {
		v651 = v663
		v654 = v657
		goto L327
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)+4))
	v681 = v680
	goto L333
L332:
	;
	v681 = v673
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+28)) = v681
	v685 = F_palloc(m, v681<<(uint(int32(2))%32))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L46
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+32)) = v685
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v617)+28))
	v691 = F_palloc(m, v688<<(uint(int32(2))%32))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L46
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+36)) = v691
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	if v694 != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if int32(0) < v695 {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	v792 = int32(0)
	goto L338
L338:
	;
	F_list_free(m, v792)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L46
	} else {
		goto L345
	}
L339:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v617)+32))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v694)+12))
	v705 = int32(0)
	goto L342
L340:
	;
	goto L341
L341:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v792 = v765
	goto L338
L342:
	;
	v727 = v705 << (uint(int32(2)) % 32)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v727+v699)))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v698+v727))) = v731
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v730)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v727+v691))) = v734
	v737 = v705 + int32(1)
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v737 < v738 {
		v705 = v737
		goto L342
	} else {
		goto L344
	}
L343:
	;
	goto L341
L344:
	;
	goto L343
L345:
	;
	F_plpgsql_adddatum(m, v617)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L46
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v617
	v9708 = v233
	goto L12
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v804
	v9708 = v233
	goto L12
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v811
	v9708 = v233
	goto L12
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v822
	v9708 = v233
	goto L12
L350:
	;
	if v966 == int32(0) {
		goto L34
	} else {
		goto L386
	}
L351:
	;
	goto L350
L352:
	;
	v966 = v853
	goto L351
L354:
	;
	v966 = int32(0)
	goto L351
L355:
	;
	v839 = v826
	goto L356
L356:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	if v848 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L354
L358:
	;
	v853 = v839
	v855 = v848
	goto L361
L359:
	;
	v878 = v839
	goto L360
L360:
	;
	goto L368
L361:
	;
	v860 = F_strcmp(m, v853+int32(12), v828)
	mBase = m.M
	if v860|base.B2i32(v855 == int32(1))&int32(0) == int32(0) {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v878 = v872
	goto L360
L363:
	;
	goto L352
L364:
	;
	goto L365
L365:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v853)+8))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	if v873 != 0 {
		v853 = v872
		v855 = v873
		goto L361
	} else {
		goto L367
	}
L367:
	;
	goto L362
L368:
	;
	goto L383
L383:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v878)+8))
	if v924 != 0 {
		v839 = v924
		goto L356
	} else {
		goto L384
	}
L384:
	;
	goto L357
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v966
	v9708 = v233
	goto L12
L387:
	;
	if v1121 == int32(0) {
		goto L33
	} else {
		goto L423
	}
L388:
	;
	goto L387
L389:
	;
	v1121 = v1008
	goto L388
L391:
	;
	v1121 = int32(0)
	goto L388
L392:
	;
	v994 = v981
	goto L393
L393:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v994)))
	if v1003 != 0 {
		goto L395
	} else {
		goto L396
	}
L394:
	;
	goto L391
L395:
	;
	v1008 = v994
	v1010 = v1003
	goto L398
L396:
	;
	v1033 = v994
	goto L397
L397:
	;
	goto L405
L398:
	;
	v1015 = F_strcmp(m, v1008+int32(12), v983)
	mBase = m.M
	if v1015|base.B2i32(v1010 == int32(1))&int32(0) == int32(0) {
		goto L400
	} else {
		goto L401
	}
L399:
	;
	v1033 = v1027
	goto L397
L400:
	;
	goto L389
L401:
	;
	goto L402
L402:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+8))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	if v1028 != 0 {
		v1008 = v1027
		v1010 = v1028
		goto L398
	} else {
		goto L404
	}
L404:
	;
	goto L399
L405:
	;
	goto L420
L420:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+8))
	if v1079 != 0 {
		v994 = v1079
		goto L393
	} else {
		goto L421
	}
L421:
	;
	goto L394
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1121
	v9708 = v233
	goto L12
L424:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+4))
	switch v1138 - int32(2) {
	case 0:
		goto L427
	case 1:
		goto L426
	default:
		goto L32
	}
L425:
	;
	if v1455 == int32(0) {
		goto L32
	} else {
		goto L500
	}
L426:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+12))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1301)))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+4))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+4))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+4))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+8))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1306)+4))
	if v1298 == int32(0) {
		goto L468
	} else {
		goto L469
	}
L427:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+12))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+4))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+4))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+4))
	if v1142 == int32(0) {
		goto L432
	} else {
		goto L433
	}
L428:
	;
	v1455 = v1296
	goto L425
L429:
	;
	v1296 = v1286
	goto L428
L430:
	;
	v1286 = v1173
	goto L429
L432:
	;
	v1286 = int32(0)
	goto L429
L433:
	;
	v1159 = v1142
	goto L434
L434:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	if v1168 != 0 {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	goto L432
L436:
	;
	v1173 = v1159
	v1175 = v1168
	goto L439
L437:
	;
	v1198 = v1159
	goto L438
L438:
	;
	if v1149 == int32(0) {
		v1239 = v1198
		goto L446
	} else {
		goto L447
	}
L439:
	;
	v1180 = F_strcmp(m, v1173+int32(12), v1147)
	mBase = m.M
	v1183 = int32(0)
	if v1180|base.B2i32(v1175 == int32(1))&base.B2i32(v1149 != v1183) == v1183 {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	v1198 = v1192
	goto L438
L441:
	;
	goto L430
L442:
	;
	goto L443
L443:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+8))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1192)))
	if v1193 != 0 {
		v1173 = v1192
		v1175 = v1193
		goto L439
	} else {
		goto L445
	}
L445:
	;
	goto L440
L446:
	;
	goto L461
L447:
	;
	v1207 = F_strcmp(m, v1198+int32(12), v1147)
	mBase = m.M
	if v1207 != 0 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v1208 = v1198
	goto L450
L449:
	;
	v1208 = v1159
	goto L450
L450:
	;
	if v1207|base.B2i32(v1168 == int32(0)) != 0 {
		v1239 = v1208
		goto L446
	} else {
		goto L451
	}
L451:
	;
	v1212 = v1159
	v1219 = v1168
	goto L452
L452:
	;
	v1223 = F_strcmp(m, v1212+int32(12), v1149)
	mBase = m.M
	if v1223|int32(0)&base.B2i32(v1219 == int32(1)) == int32(0) {
		goto L454
	} else {
		goto L455
	}
L453:
	;
	v1239 = v1233
	goto L446
L454:
	;
	goto L457
L455:
	;
	goto L456
L456:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+8))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1233)))
	if v1234 != 0 {
		v1212 = v1233
		v1219 = v1234
		goto L452
	} else {
		goto L460
	}
L457:
	;
	v1296 = v1212
	goto L428
L460:
	;
	goto L453
L461:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+8))
	if v1244 != 0 {
		v1159 = v1244
		goto L434
	} else {
		goto L462
	}
L462:
	;
	goto L435
L464:
	;
	v1455 = v1453
	goto L425
L465:
	;
	v1453 = v1443
	goto L464
L466:
	;
	v1443 = v1330
	goto L465
L468:
	;
	v1443 = int32(0)
	goto L465
L469:
	;
	v1316 = v1298
	goto L470
L470:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1316)))
	if v1325 != 0 {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	goto L468
L472:
	;
	v1330 = v1316
	v1332 = v1325
	goto L475
L473:
	;
	v1355 = v1316
	goto L474
L474:
	;
	if v1305 == int32(0) {
		v1396 = v1355
		goto L482
	} else {
		goto L483
	}
L475:
	;
	v1337 = F_strcmp(m, v1330+int32(12), v1303)
	mBase = m.M
	v1340 = int32(0)
	if v1337|base.B2i32(v1332 == int32(1))&base.B2i32(v1305 != v1340) == v1340 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v1355 = v1349
	goto L474
L477:
	;
	goto L466
L478:
	;
	goto L479
L479:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+8))
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	if v1350 != 0 {
		v1330 = v1349
		v1332 = v1350
		goto L475
	} else {
		goto L481
	}
L481:
	;
	goto L476
L482:
	;
	goto L497
L483:
	;
	v1364 = F_strcmp(m, v1355+int32(12), v1303)
	mBase = m.M
	if v1364 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1365 = v1355
	goto L486
L485:
	;
	v1365 = v1316
	goto L486
L486:
	;
	if v1364|base.B2i32(v1325 == int32(0)) != 0 {
		v1396 = v1365
		goto L482
	} else {
		goto L487
	}
L487:
	;
	v1369 = v1316
	v1376 = v1325
	goto L488
L488:
	;
	v1380 = F_strcmp(m, v1369+int32(12), v1305)
	mBase = m.M
	if v1380|base.B2i32(v1307 != int32(0))&base.B2i32(v1376 == int32(1)) == int32(0) {
		goto L490
	} else {
		goto L491
	}
L489:
	;
	v1396 = v1390
	goto L482
L490:
	;
	goto L493
L491:
	;
	goto L492
L492:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+8))
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1390)))
	if v1391 != 0 {
		v1369 = v1390
		v1376 = v1391
		goto L488
	} else {
		goto L496
	}
L493:
	;
	v1453 = v1369
	goto L464
L496:
	;
	goto L489
L497:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+8))
	if v1401 != 0 {
		v1316 = v1401
		goto L470
	} else {
		goto L498
	}
L498:
	;
	goto L471
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1455
	v9708 = v233
	goto L12
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v1507
	v1511 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1511 == int32(0) {
		goto L518
	} else {
		goto L519
	}
L502:
	;
	goto L501
L503:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+60))
	if v1468 == int32(0) {
		v1507 = v1462
		goto L502
	} else {
		goto L504
	}
L504:
	;
	v1471 = v1461 + v1468
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+188))
	if base.Ui32(v1472) <= base.Ui32(v1471) {
		goto L506
	} else {
		goto L507
	}
L505:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+196))
	if base.B2i32(v1481 == int32(0))|base.B2i32(base.Ui32(v1471) <= base.Ui32(v1481)) != 0 {
		v1507 = v1482
		goto L502
	} else {
		goto L509
	}
L506:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1467)+192))
	v1481 = v1474
	goto L505
L507:
	;
	goto L508
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+188)) = v1468
	v1479 = F_strchr(m, v1468, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+192)) = v1479
	v1481 = v1479
	goto L505
L509:
	;
	v1487 = v1481
	v1490 = v1482
	goto L510
L510:
	;
	v1492 = int32(1)
	v1493 = v1490 + v1492
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+196)) = v1493
	v1496 = v1487 + v1492
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+188)) = v1496
	v1499 = F_strchr(m, v1496, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1467)+192)) = v1499
	if v1499 == int32(0) {
		v1507 = v1493
		goto L502
	} else {
		goto L512
	}
L511:
	;
	v1507 = v1493
	goto L502
L512:
	;
	if base.Ui32(v1499) < base.Ui32(v1471) {
		v1487 = v1499
		v1490 = v1493
		goto L510
	} else {
		goto L513
	}
L513:
	;
	goto L511
L514:
	;
	if v1651 != 0 {
		goto L2
	} else {
		goto L550
	}
L515:
	;
	goto L514
L516:
	;
	v1651 = v1538
	goto L515
L518:
	;
	v1651 = int32(0)
	goto L515
L519:
	;
	goto L520
L520:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	if v1533 != 0 {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v1538 = v1511
	v1540 = v1533
	goto L525
L523:
	;
	goto L524
L524:
	;
	goto L532
L525:
	;
	v1545 = F_strcmp(m, v1538+int32(12), v1513)
	mBase = m.M
	if v1545|base.B2i32(v1540 == int32(1))&int32(0) == int32(0) {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	goto L524
L527:
	;
	goto L516
L528:
	;
	goto L529
L529:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+8))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1557)))
	if v1558 != 0 {
		v1538 = v1557
		v1540 = v1558
		goto L525
	} else {
		goto L531
	}
L531:
	;
	goto L526
L532:
	;
	goto L518
L550:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v1664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1663)+496)))
	if v1664&int32(2) == int32(0) {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	v1669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1663)+500)))
	if v1669&int32(2) == int32(0) {
		v9708 = v233
		goto L12
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1675 == int32(0) {
		goto L559
	} else {
		goto L560
	}
L554:
	;
	goto L553
L555:
	;
	if v1815 == int32(0) {
		v9708 = v233
		goto L12
	} else {
		goto L591
	}
L556:
	;
	goto L555
L557:
	;
	v1815 = v1702
	goto L556
L559:
	;
	v1815 = int32(0)
	goto L556
L560:
	;
	v1688 = v1675
	goto L561
L561:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1688)))
	if v1697 != 0 {
		goto L563
	} else {
		goto L564
	}
L562:
	;
	goto L559
L563:
	;
	v1702 = v1688
	v1704 = v1697
	goto L566
L564:
	;
	v1727 = v1688
	goto L565
L565:
	;
	goto L573
L566:
	;
	v1709 = F_strcmp(m, v1702+int32(12), v1677)
	mBase = m.M
	if v1709|base.B2i32(v1704 == int32(1))&int32(0) == int32(0) {
		goto L568
	} else {
		goto L569
	}
L567:
	;
	v1727 = v1721
	goto L565
L568:
	;
	goto L557
L569:
	;
	goto L570
L570:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+8))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1721)))
	if v1722 != 0 {
		v1702 = v1721
		v1704 = v1722
		goto L566
	} else {
		goto L572
	}
L572:
	;
	goto L567
L573:
	;
	goto L588
L588:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1727)+8))
	if v1773 != 0 {
		v1688 = v1773
		goto L561
	} else {
		goto L589
	}
L589:
	;
	goto L562
L591:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+500))
	v1836 = F_errstart(m, v1830&int32(2)+int32(19), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L46
	} else {
		goto L592
	}
L592:
	;
	if v1836 == int32(0) {
		v9708 = v233
		goto L12
	} else {
		goto L593
	}
L593:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L46
	} else {
		goto L594
	}
L594:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v1843
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_25), v28+int32(96))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L46
	} else {
		goto L595
	}
L595:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v1851 = F_plpgsql_scanner_errposition(m, v1850, l1)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L46
	} else {
		goto L596
	}
L596:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(737), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L46
	} else {
		goto L597
	}
L597:
	;
	v9708 = v233
	goto L12
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1859
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v1863 = int32(0)
	if v1862 < v1863 {
		v1908 = v1863
		goto L600
	} else {
		goto L601
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v1908
	v1912 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1912 == int32(0) {
		goto L616
	} else {
		goto L617
	}
L600:
	;
	goto L599
L601:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1868)+60))
	if v1869 == int32(0) {
		v1908 = v1863
		goto L600
	} else {
		goto L602
	}
L602:
	;
	v1872 = v1862 + v1869
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1868)+188))
	if base.Ui32(v1873) <= base.Ui32(v1872) {
		goto L604
	} else {
		goto L605
	}
L603:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1868)+196))
	if base.B2i32(v1882 == int32(0))|base.B2i32(base.Ui32(v1872) <= base.Ui32(v1882)) != 0 {
		v1908 = v1883
		goto L600
	} else {
		goto L607
	}
L604:
	;
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1868)+192))
	v1882 = v1875
	goto L603
L605:
	;
	goto L606
L606:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1868)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1868)+188)) = v1869
	v1880 = F_strchr(m, v1869, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1868)+192)) = v1880
	v1882 = v1880
	goto L603
L607:
	;
	v1888 = v1882
	v1891 = v1883
	goto L608
L608:
	;
	v1893 = int32(1)
	v1894 = v1891 + v1893
	*(*int32)(unsafe.Add(mBase, uint32(v1868)+196)) = v1894
	v1897 = v1888 + v1893
	*(*int32)(unsafe.Add(mBase, uint32(v1868)+188)) = v1897
	v1900 = F_strchr(m, v1897, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1868)+192)) = v1900
	if v1900 == int32(0) {
		v1908 = v1894
		goto L600
	} else {
		goto L610
	}
L609:
	;
	v1908 = v1894
	goto L600
L610:
	;
	if base.Ui32(v1900) < base.Ui32(v1872) {
		v1888 = v1900
		v1891 = v1894
		goto L608
	} else {
		goto L611
	}
L611:
	;
	goto L609
L612:
	;
	if v2052 != 0 {
		goto L2
	} else {
		goto L648
	}
L613:
	;
	goto L612
L614:
	;
	v2052 = v1939
	goto L613
L616:
	;
	v2052 = int32(0)
	goto L613
L617:
	;
	goto L618
L618:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1912)))
	if v1934 != 0 {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v1939 = v1912
	v1941 = v1934
	goto L623
L621:
	;
	goto L622
L622:
	;
	goto L630
L623:
	;
	v1946 = F_strcmp(m, v1939+int32(12), v1914)
	mBase = m.M
	if v1946|base.B2i32(v1941 == int32(1))&int32(0) == int32(0) {
		goto L625
	} else {
		goto L626
	}
L624:
	;
	goto L622
L625:
	;
	goto L614
L626:
	;
	goto L627
L627:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1939)+8))
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v1958)))
	if v1959 != 0 {
		v1939 = v1958
		v1941 = v1959
		goto L623
	} else {
		goto L629
	}
L629:
	;
	goto L624
L630:
	;
	goto L616
L648:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2064)+496)))
	if v2065&int32(2) == int32(0) {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	v2070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2064)+500)))
	if v2070&int32(2) == int32(0) {
		v9708 = v233
		goto L12
	} else {
		goto L652
	}
L650:
	;
	goto L651
L651:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v2076 == int32(0) {
		goto L657
	} else {
		goto L658
	}
L652:
	;
	goto L651
L653:
	;
	if v2216 == int32(0) {
		v9708 = v233
		goto L12
	} else {
		goto L689
	}
L654:
	;
	goto L653
L655:
	;
	v2216 = v2103
	goto L654
L657:
	;
	v2216 = int32(0)
	goto L654
L658:
	;
	v2089 = v2076
	goto L659
L659:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2089)))
	if v2098 != 0 {
		goto L661
	} else {
		goto L662
	}
L660:
	;
	goto L657
L661:
	;
	v2103 = v2089
	v2105 = v2098
	goto L664
L662:
	;
	v2128 = v2089
	goto L663
L663:
	;
	goto L671
L664:
	;
	v2110 = F_strcmp(m, v2103+int32(12), v2078)
	mBase = m.M
	if v2110|base.B2i32(v2105 == int32(1))&int32(0) == int32(0) {
		goto L666
	} else {
		goto L667
	}
L665:
	;
	v2128 = v2122
	goto L663
L666:
	;
	goto L655
L667:
	;
	goto L668
L668:
	;
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2103)+8))
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2122)))
	if v2123 != 0 {
		v2103 = v2122
		v2105 = v2123
		goto L664
	} else {
		goto L670
	}
L670:
	;
	goto L665
L671:
	;
	goto L686
L686:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+8))
	if v2174 != 0 {
		v2089 = v2174
		goto L659
	} else {
		goto L687
	}
L687:
	;
	goto L660
L689:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+500))
	v2237 = F_errstart(m, v2231&int32(2)+int32(19), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L46
	} else {
		goto L690
	}
L690:
	;
	if v2237 == int32(0) {
		v9708 = v233
		goto L12
	} else {
		goto L691
	}
L691:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L46
	} else {
		goto L692
	}
L692:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v2244
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_25), v28+int32(112))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L46
	} else {
		goto L693
	}
L693:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v2252 = F_plpgsql_scanner_errposition(m, v2251, l1)
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L46
	} else {
		goto L694
	}
L694:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(765), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L46
	} else {
		goto L695
	}
L695:
	;
	v9708 = v233
	goto L12
L696:
	;
	v2273 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L46
	} else {
		goto L699
	}
L697:
	;
	v2275 = v233
	goto L698
L698:
	;
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	if v2275 == int32(275) {
		goto L704
	} else {
		goto L705
	}
L699:
	;
	v2275 = v2273
	goto L698
L700:
	;
	m.G0 = v2269 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2783
	v9708 = int32(-2)
	goto L12
L701:
	;
	v2675 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L46
	} else {
		goto L855
	}
L702:
	;
	v2559 = v2552
	v2566 = int32(0)
	goto L814
L703:
	;
	if v2549 != 0 {
		goto L701
	} else {
		goto L813
	}
L704:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v2280 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L46
	} else {
		goto L707
	}
L705:
	;
	goto L706
L706:
	;
	v2363 = int32(0)
	goto L739
L707:
	;
	if v2280 != int32(37) {
		v2552 = v2280
		goto L702
	} else {
		goto L708
	}
L708:
	;
	v2284 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L46
	} else {
		goto L709
	}
L709:
	;
	switch v2284 - int32(366) {
	case 0:
		goto L710
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v2552 = v2284
		goto L702
	case 12:
		goto L712
	default:
		goto L713
	}
L710:
	;
	v2358 = F_plpgsql_parse_wordrowtype(m, v2279)
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L46
	} else {
		goto L737
	}
L711:
	;
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v2326 == int32(0) {
		v2552 = v2288
		goto L702
	} else {
		goto L728
	}
L712:
	;
	v2324 = F_plpgsql_parse_wordtype(m, v2279)
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L46
	} else {
		goto L727
	}
L713:
	;
	v2288 = int32(277)
	if v2284 != v2288 {
		goto L714
	} else {
		goto L715
	}
L714:
	;
	v2552 = v2284
	goto L702
L715:
	;
	goto L716
L716:
	;
	v2291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2291 != 0 {
		v2552 = v2288
		goto L702
	} else {
		goto L717
	}
L717:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v2292 == int32(0) {
		goto L711
	} else {
		goto L718
	}
L718:
	;
	v2295 = int32(_a_F_plpgsql_yyparse_26)
	v2298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2292))))
	v2301 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[18])))
	if base.B2i32(v2298 == int32(0))|base.B2i32(v2298 != v2301) != 0 {
		v2319 = v2298
		v2320 = v2301
		goto L720
	} else {
		goto L721
	}
L719:
	;
	if v2319-v2320 != 0 {
		goto L711
	} else {
		goto L726
	}
L720:
	;
	goto L719
L721:
	;
	v2304 = v2292
	v2305 = v2295
	goto L722
L722:
	;
	v2308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2305)+1)))
	v2309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2304)+1)))
	if v2309 == int32(0) {
		v2319 = v2309
		v2320 = v2308
		goto L720
	} else {
		goto L724
	}
L723:
	;
	v2319 = v2309
	v2320 = v2308
	goto L720
L724:
	;
	v2312 = int32(1)
	if v2309 == v2308 {
		v2304 = v2304 + v2312
		v2305 = v2305 + v2312
		goto L722
	} else {
		goto L725
	}
L725:
	;
	goto L723
L726:
	;
	goto L712
L727:
	;
	v2549 = v2324
	v2551 = v2284
	goto L703
L728:
	;
	v2329 = int32(_a_F_plpgsql_yyparse_27)
	v2332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2326))))
	v2335 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[19])))
	if base.B2i32(v2332 == int32(0))|base.B2i32(v2332 != v2335) != 0 {
		v2353 = v2332
		v2354 = v2335
		goto L730
	} else {
		goto L731
	}
L729:
	;
	if v2353-v2354 != 0 {
		v2552 = v2288
		goto L702
	} else {
		goto L736
	}
L730:
	;
	goto L729
L731:
	;
	v2338 = v2326
	v2339 = v2329
	goto L732
L732:
	;
	v2342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2339)+1)))
	v2343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2338)+1)))
	if v2343 == int32(0) {
		v2353 = v2343
		v2354 = v2342
		goto L730
	} else {
		goto L734
	}
L733:
	;
	v2353 = v2343
	v2354 = v2342
	goto L730
L734:
	;
	v2346 = int32(1)
	if v2343 == v2342 {
		v2338 = v2338 + v2346
		v2339 = v2339 + v2346
		goto L732
	} else {
		goto L735
	}
L735:
	;
	goto L733
L736:
	;
	goto L710
L737:
	;
	v2549 = v2358
	v2551 = v2284
	goto L703
L738:
	;
	if v2275 == v2369 {
		goto L745
	} else {
		goto L746
	}
L739:
	;
	v2369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2363<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[20]))))
	v2370 = base.B2i32(v2275 == v2369)
	if v2370 == int32(0) {
		goto L741
	} else {
		goto L742
	}
L740:
	;
	goto L738
L741:
	;
	v2374 = v2363 + int32(1)
	if v2374 != int32(83) {
		v2363 = v2374
		goto L739
	} else {
		goto L744
	}
L742:
	;
	goto L743
L743:
	;
	goto L740
L744:
	;
	goto L743
L745:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v2379 = F_pstrdup(m, v2378)
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L46
	} else {
		goto L748
	}
L746:
	;
	goto L747
L747:
	;
	if v2275 != int32(276) {
		v2552 = v2275
		goto L702
	} else {
		goto L780
	}
L748:
	;
	v2381 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L46
	} else {
		goto L749
	}
L749:
	;
	if v2381 != int32(37) {
		v2552 = v2381
		goto L702
	} else {
		goto L750
	}
L750:
	;
	v2385 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L46
	} else {
		goto L751
	}
L751:
	;
	switch v2385 - int32(366) {
	case 0:
		goto L752
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v2552 = v2385
		goto L702
	case 12:
		goto L754
	default:
		goto L755
	}
L752:
	;
	v2459 = F_plpgsql_parse_wordrowtype(m, v2379)
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L46
	} else {
		goto L779
	}
L753:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v2427 == int32(0) {
		v2552 = v2389
		goto L702
	} else {
		goto L770
	}
L754:
	;
	v2425 = F_plpgsql_parse_wordtype(m, v2379)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L46
	} else {
		goto L769
	}
L755:
	;
	v2389 = int32(277)
	if v2385 != v2389 {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v2552 = v2385
	goto L702
L757:
	;
	goto L758
L758:
	;
	v2392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2392 != 0 {
		v2552 = v2389
		goto L702
	} else {
		goto L759
	}
L759:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v2393 == int32(0) {
		goto L753
	} else {
		goto L760
	}
L760:
	;
	v2396 = int32(_a_F_plpgsql_yyparse_26)
	v2399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2393))))
	v2402 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[18])))
	if base.B2i32(v2399 == int32(0))|base.B2i32(v2399 != v2402) != 0 {
		v2420 = v2399
		v2421 = v2402
		goto L762
	} else {
		goto L763
	}
L761:
	;
	if v2420-v2421 != 0 {
		goto L753
	} else {
		goto L768
	}
L762:
	;
	goto L761
L763:
	;
	v2405 = v2393
	v2406 = v2396
	goto L764
L764:
	;
	v2409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2406)+1)))
	v2410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2405)+1)))
	if v2410 == int32(0) {
		v2420 = v2410
		v2421 = v2409
		goto L762
	} else {
		goto L766
	}
L765:
	;
	v2420 = v2410
	v2421 = v2409
	goto L762
L766:
	;
	v2413 = int32(1)
	if v2410 == v2409 {
		v2405 = v2405 + v2413
		v2406 = v2406 + v2413
		goto L764
	} else {
		goto L767
	}
L767:
	;
	goto L765
L768:
	;
	goto L754
L769:
	;
	v2549 = v2425
	v2551 = v2385
	goto L703
L770:
	;
	v2430 = int32(_a_F_plpgsql_yyparse_27)
	v2433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2427))))
	v2436 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[19])))
	if base.B2i32(v2433 == int32(0))|base.B2i32(v2433 != v2436) != 0 {
		v2454 = v2433
		v2455 = v2436
		goto L772
	} else {
		goto L773
	}
L771:
	;
	if v2454-v2455 != 0 {
		v2552 = v2389
		goto L702
	} else {
		goto L778
	}
L772:
	;
	goto L771
L773:
	;
	v2439 = v2427
	v2440 = v2430
	goto L774
L774:
	;
	v2443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2440)+1)))
	v2444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2439)+1)))
	if v2444 == int32(0) {
		v2454 = v2444
		v2455 = v2443
		goto L772
	} else {
		goto L776
	}
L775:
	;
	v2454 = v2444
	v2455 = v2443
	goto L772
L776:
	;
	v2447 = int32(1)
	if v2444 == v2443 {
		v2439 = v2439 + v2447
		v2440 = v2440 + v2447
		goto L774
	} else {
		goto L777
	}
L777:
	;
	goto L775
L778:
	;
	goto L752
L779:
	;
	v2549 = v2459
	v2551 = v2385
	goto L703
L780:
	;
	v2463 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v2464 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L46
	} else {
		goto L781
	}
L781:
	;
	if v2464 != int32(37) {
		v2552 = v2464
		goto L702
	} else {
		goto L782
	}
L782:
	;
	v2468 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L46
	} else {
		goto L783
	}
L783:
	;
	switch v2468 - int32(366) {
	case 0:
		goto L785
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v2552 = v2468
		goto L702
	case 12:
		goto L787
	default:
		goto L788
	}
L784:
	;
	v2549 = v2546
	v2551 = v2468
	goto L703
L785:
	;
	v2542 = F_plpgsql_parse_cwordrowtype(m, v2463)
	mBase = m.M
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L46
	} else {
		goto L812
	}
L786:
	;
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v2510 == int32(0) {
		v2552 = v2472
		goto L702
	} else {
		goto L803
	}
L787:
	;
	v2508 = F_plpgsql_parse_cwordtype(m, v2463)
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
		goto L46
	} else {
		goto L802
	}
L788:
	;
	v2472 = int32(277)
	if v2468 != v2472 {
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v2552 = v2468
	goto L702
L790:
	;
	goto L791
L791:
	;
	v2475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2475 != 0 {
		v2552 = v2472
		goto L702
	} else {
		goto L792
	}
L792:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v2476 == int32(0) {
		goto L786
	} else {
		goto L793
	}
L793:
	;
	v2479 = int32(_a_F_plpgsql_yyparse_26)
	v2482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476))))
	v2485 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[18])))
	if base.B2i32(v2482 == int32(0))|base.B2i32(v2482 != v2485) != 0 {
		v2503 = v2482
		v2504 = v2485
		goto L795
	} else {
		goto L796
	}
L794:
	;
	if v2503-v2504 != 0 {
		goto L786
	} else {
		goto L801
	}
L795:
	;
	goto L794
L796:
	;
	v2488 = v2476
	v2489 = v2479
	goto L797
L797:
	;
	v2492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2489)+1)))
	v2493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2488)+1)))
	if v2493 == int32(0) {
		v2503 = v2493
		v2504 = v2492
		goto L795
	} else {
		goto L799
	}
L798:
	;
	v2503 = v2493
	v2504 = v2492
	goto L795
L799:
	;
	v2496 = int32(1)
	if v2493 == v2492 {
		v2488 = v2488 + v2496
		v2489 = v2489 + v2496
		goto L797
	} else {
		goto L800
	}
L800:
	;
	goto L798
L801:
	;
	goto L787
L802:
	;
	v2546 = v2508
	goto L784
L803:
	;
	v2513 = int32(_a_F_plpgsql_yyparse_27)
	v2516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2510))))
	v2519 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[19])))
	if base.B2i32(v2516 == int32(0))|base.B2i32(v2516 != v2519) != 0 {
		v2537 = v2516
		v2538 = v2519
		goto L805
	} else {
		goto L806
	}
L804:
	;
	if v2537-v2538 != 0 {
		v2552 = v2472
		goto L702
	} else {
		goto L811
	}
L805:
	;
	goto L804
L806:
	;
	v2522 = v2510
	v2523 = v2513
	goto L807
L807:
	;
	v2526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2523)+1)))
	v2527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2522)+1)))
	if v2527 == int32(0) {
		v2537 = v2527
		v2538 = v2526
		goto L805
	} else {
		goto L809
	}
L808:
	;
	v2537 = v2527
	v2538 = v2526
	goto L805
L809:
	;
	v2530 = int32(1)
	if v2527 == v2526 {
		v2522 = v2522 + v2530
		v2523 = v2523 + v2530
		goto L807
	} else {
		goto L810
	}
L810:
	;
	goto L808
L811:
	;
	goto L785
L812:
	;
	v2546 = v2542
	goto L784
L813:
	;
	v2552 = v2551
	goto L702
L814:
	;
	if v2559 <= int32(292) {
		goto L821
	} else {
		goto L822
	}
L815:
	;
	v2623 = v2269 + int32(4)
	F_initStringInfo(m, v2623)
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L46
	} else {
		goto L839
	}
L816:
	;
	goto L815
L817:
	;
	if base.B2i32(v2559 != int32(44))&base.B2i32(v2559 != int32(41))|v2566 == int32(0) {
		goto L816
	} else {
		goto L831
	}
L818:
	;
	F_plpgsql_yyerror(m, v2266, int32(0), l1, int32(_a_F_plpgsql_yyparse_28))
	mBase = m.M
	v2601 = m.ExcPending
	if v2601 != 0 {
		goto L46
	} else {
		goto L830
	}
L819:
	;
	if v2559 != int32(343) {
		goto L817
	} else {
		goto L829
	}
L820:
	;
	if v2566 != 0 {
		goto L818
	} else {
		goto L827
	}
L821:
	;
	switch v2559 - int32(59) {
	case 0, 2:
		goto L816
	case 1:
		goto L817
	default:
		goto L824
	}
L822:
	;
	goto L823
L823:
	;
	switch v2559 - int32(293) {
	case 0, 13:
		goto L816
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L817
	default:
		goto L819
	}
L824:
	;
	if v2559 == int32(270) {
		goto L816
	} else {
		goto L825
	}
L825:
	;
	if v2559 == int32(0) {
		goto L820
	} else {
		goto L826
	}
L826:
	;
	goto L817
L827:
	;
	F_plpgsql_yyerror(m, v2266, int32(0), l1, int32(_a_F_plpgsql_yyparse_29))
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L46
	} else {
		goto L828
	}
L828:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L829:
	;
	goto L816
L830:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L831:
	;
	if v2559 == int32(41) {
		goto L832
	} else {
		goto L833
	}
L832:
	;
	v2615 = int32(-1)
	goto L834
L833:
	;
	v2615 = int32(0)
	goto L834
L834:
	;
	if v2559 == int32(40) {
		goto L835
	} else {
		goto L836
	}
L835:
	;
	v2618 = int32(1)
	goto L837
L836:
	;
	v2618 = v2615
	goto L837
L837:
	;
	v2620 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L46
	} else {
		goto L838
	}
L838:
	;
	v2559 = v2620
	v2566 = v2618 + v2566
	goto L814
L839:
	;
	v2626 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	F_plpgsql_append_source_text(m, v2623, v2276, v2626, l1)
	mBase = m.M
	v2628 = m.ExcPending
	if v2628 != 0 {
		goto L46
	} else {
		goto L840
	}
L840:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+4))
	v2630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2629))))
	if v2630 != 0 {
		goto L841
	} else {
		goto L842
	}
L841:
	;
	v2631 = int32(_a_F_plpgsql_yyparse_30)
	v2632 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[21])) = v2269 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+32)) = v2276
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+24)) = int32(_a_F_plpgsql_yyparse_31)
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+20)) = v2632
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+28)) = v2269 + int32(32)
	v2645 = int32(0)
	v2647 = F_typeStringToTypeName(m, v2629, v2645)
	mBase = m.M
	v2648 = m.ExcPending
	if v2648 != 0 {
		goto L46
	} else {
		goto L844
	}
L842:
	;
	goto L843
L843:
	;
	F_plpgsql_yyerror(m, v2266, int32(0), l1, int32(_a_F_plpgsql_yyparse_32))
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L46
	} else {
		goto L849
	}
L844:
	;
	F_typenameTypeIdAndMod(m, v2645, v2647, v2269+int32(44), v2269+int32(40))
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L46
	} else {
		goto L845
	}
L845:
	;
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[21])) = v2656
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+44))
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+40))
	v2661 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2661)+44))
	v2663 = F_plpgsql_build_datatype(m, v2658, v2659, v2662, v2647)
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L46
	} else {
		goto L846
	}
L846:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+4))
	F_pfree(m, v2665)
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L46
	} else {
		goto L847
	}
L847:
	;
	F_plpgsql_push_back_token(m, v2559, v2264, v2266, l1)
	mBase = m.M
	v2669 = m.ExcPending
	if v2669 != 0 {
		goto L46
	} else {
		goto L848
	}
L848:
	;
	v2783 = v2663
	goto L700
L849:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L850:
	;
	F_plpgsql_yyerror(m, v2266, int32(0), l1, int32(_a_F_plpgsql_yyparse_33))
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
		goto L46
	} else {
		goto L886
	}
L851:
	;
	F_plpgsql_push_back_token(m, int32(277), v2264, v2266, l1)
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L46
	} else {
		goto L885
	}
L852:
	;
	if v2714 == int32(91) {
		goto L867
	} else {
		goto L868
	}
L853:
	;
	v2712 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L46
	} else {
		goto L866
	}
L854:
	;
	v2679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2679 != 0 {
		goto L851
	} else {
		goto L856
	}
L855:
	;
	switch v2675 - int32(277) {
	case 0:
		goto L854
	default:
		v2714 = v2675
		v2715 = int32(0)
		goto L852
	case 7:
		goto L853
	}
L856:
	;
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v2680 == int32(0) {
		goto L851
	} else {
		goto L857
	}
L857:
	;
	v2683 = int32(_a_F_plpgsql_yyparse_34)
	v2686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2680))))
	v2689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[22])))
	if base.B2i32(v2686 == int32(0))|base.B2i32(v2686 != v2689) != 0 {
		v2707 = v2686
		v2708 = v2689
		goto L859
	} else {
		goto L860
	}
L858:
	;
	if v2707-v2708 != 0 {
		goto L851
	} else {
		goto L865
	}
L859:
	;
	goto L858
L860:
	;
	v2692 = v2680
	v2693 = v2683
	goto L861
L861:
	;
	v2696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2693)+1)))
	v2697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2692)+1)))
	if v2697 == int32(0) {
		v2707 = v2697
		v2708 = v2696
		goto L859
	} else {
		goto L863
	}
L862:
	;
	v2707 = v2697
	v2708 = v2696
	goto L859
L863:
	;
	v2700 = int32(1)
	if v2697 == v2696 {
		v2692 = v2692 + v2700
		v2693 = v2693 + v2700
		goto L861
	} else {
		goto L864
	}
L864:
	;
	goto L862
L865:
	;
	goto L853
L866:
	;
	v2714 = v2712
	v2715 = int32(1)
	goto L852
L867:
	;
	goto L870
L868:
	;
	goto L869
L869:
	;
	F_plpgsql_push_back_token(m, v2714, v2264, v2266, l1)
	mBase = m.M
	v2761 = m.ExcPending
	if v2761 != 0 {
		goto L46
	} else {
		goto L882
	}
L870:
	;
	v2743 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2744 = m.ExcPending
	if v2744 != 0 {
		goto L46
	} else {
		goto L872
	}
L871:
	;
	F_plpgsql_push_back_token(m, v2752, v2264, v2266, l1)
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L46
	} else {
		goto L880
	}
L872:
	;
	if v2743 == int32(266) {
		goto L873
	} else {
		goto L874
	}
L873:
	;
	v2747 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L46
	} else {
		goto L876
	}
L874:
	;
	v2749 = v2743
	goto L875
L875:
	;
	if v2749 != int32(93) {
		goto L850
	} else {
		goto L877
	}
L876:
	;
	v2749 = v2747
	goto L875
L877:
	;
	v2752 = F_plpgsql_yylex(m, v2264, v2266, l1)
	mBase = m.M
	v2753 = m.ExcPending
	if v2753 != 0 {
		goto L46
	} else {
		goto L878
	}
L878:
	;
	if v2752 == int32(91) {
		goto L870
	} else {
		goto L879
	}
L879:
	;
	goto L871
L880:
	;
	v2758 = F_plpgsql_build_datatype_arrayof(m, v2549)
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L46
	} else {
		goto L881
	}
L881:
	;
	v2783 = v2758
	goto L700
L882:
	;
	if v2715 == int32(0) {
		v2783 = v2549
		goto L700
	} else {
		goto L883
	}
L883:
	;
	v2764 = F_plpgsql_build_datatype_arrayof(m, v2549)
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L46
	} else {
		goto L884
	}
L884:
	;
	v2783 = v2764
	goto L700
L885:
	;
	v2783 = v2549
	goto L700
L886:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L887:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v2807
	*(*int32)(unsafe.Add(mBase, uint32(v28)+264)) = v2807
	v2814 = F_list_make1_impl(m, int32(1), v28+int32(120))
	mBase = m.M
	v2815 = m.ExcPending
	if v2815 != 0 {
		goto L46
	} else {
		goto L888
	}
L888:
	;
	v2817 = F_get_collation_oid(m, v2814, int32(0))
	mBase = m.M
	v2818 = m.ExcPending
	if v2818 != 0 {
		goto L46
	} else {
		goto L889
	}
L889:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2817
	v9708 = v233
	goto L12
L890:
	;
	v2823 = F_makeString(m, v2821)
	mBase = m.M
	v2824 = m.ExcPending
	if v2824 != 0 {
		goto L46
	} else {
		goto L891
	}
L891:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v2823
	*(*int32)(unsafe.Add(mBase, uint32(v28)+260)) = v2823
	v2830 = F_list_make1_impl(m, int32(1), v28+int32(124))
	mBase = m.M
	v2831 = m.ExcPending
	if v2831 != 0 {
		goto L46
	} else {
		goto L892
	}
L892:
	;
	v2833 = F_get_collation_oid(m, v2830, int32(0))
	mBase = m.M
	v2834 = m.ExcPending
	if v2834 != 0 {
		goto L46
	} else {
		goto L893
	}
L893:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2833
	v9708 = v233
	goto L12
L894:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2838
	v9708 = v233
	goto L12
L895:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2860
	v9708 = v233
	goto L12
L896:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2867
	v9708 = v233
	goto L12
L897:
	;
	goto L898
L898:
	;
	v2872 = F_lappend(m, v2867, v2868)
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L46
	} else {
		goto L899
	}
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2872
	v9708 = v233
	goto L12
L900:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2926))) = int32(23)
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v2931 = int32(0)
	if v2930 < v2931 {
		v2976 = v2931
		goto L902
	} else {
		goto L903
	}
L901:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2926)+4)) = v2976
	v2980 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v2981 = *(*int32)(unsafe.Add(mBase, uint32(v2980)+520))
	v2983 = v2981 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2980)+520)) = v2983
	*(*int32)(unsafe.Add(mBase, uint32(v2926)+8)) = v2983
	v2988 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v2990 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(349), v2988, v2990, l1)
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L46
	} else {
		goto L914
	}
L902:
	;
	goto L901
L903:
	;
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2937 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+60))
	if v2937 == int32(0) {
		v2976 = v2931
		goto L902
	} else {
		goto L904
	}
L904:
	;
	v2940 = v2930 + v2937
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+188))
	if base.Ui32(v2941) <= base.Ui32(v2940) {
		goto L906
	} else {
		goto L907
	}
L905:
	;
	v2951 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+196))
	if base.B2i32(v2950 == int32(0))|base.B2i32(base.Ui32(v2940) <= base.Ui32(v2950)) != 0 {
		v2976 = v2951
		goto L902
	} else {
		goto L909
	}
L906:
	;
	v2943 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+192))
	v2950 = v2943
	goto L905
L907:
	;
	goto L908
L908:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2936)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2936)+188)) = v2937
	v2948 = F_strchr(m, v2937, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2936)+192)) = v2948
	v2950 = v2948
	goto L905
L909:
	;
	v2956 = v2950
	v2959 = v2951
	goto L910
L910:
	;
	v2961 = int32(1)
	v2962 = v2959 + v2961
	*(*int32)(unsafe.Add(mBase, uint32(v2936)+196)) = v2962
	v2965 = v2956 + v2961
	*(*int32)(unsafe.Add(mBase, uint32(v2936)+188)) = v2965
	v2968 = F_strchr(m, v2965, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2936)+192)) = v2968
	if v2968 == int32(0) {
		v2976 = v2962
		goto L902
	} else {
		goto L912
	}
L911:
	;
	v2976 = v2962
	goto L902
L912:
	;
	if base.Ui32(v2968) < base.Ui32(v2940) {
		v2956 = v2968
		v2959 = v2962
		goto L910
	} else {
		goto L913
	}
L913:
	;
	goto L911
L914:
	;
	v2994 = int32(0)
	v3003 = F_read_sql_construct(m, int32(59), v2994, v2994, int32(_a_F_plpgsql_yyparse_16), v2994, v2994, v2994, v28+int32(256), v2994, v2988, v2990, l1)
	mBase = m.M
	v3004 = m.ExcPending
	if v3004 != 0 {
		goto L46
	} else {
		goto L915
	}
L915:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2926)+12)) = v3003
	v3006 = *(*int32)(unsafe.Add(mBase, uint32(v3003)))
	v3008 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v3006)+3)) = v3008
	v3011 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v3006))) = v3011
	v3013 = *(*int32)(unsafe.Add(mBase, uint32(v2926)+12))
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v3013)))
	v3015 = F_strlen(m, v3014)
	mBase = m.M
	if v3015 != 0 {
		goto L916
	} else {
		goto L917
	}
L916:
	;
	base.MemoryCopy(m, v3014, v3014+int32(1), v3015)
	goto L918
L917:
	;
	goto L918
L918:
	;
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v2926)+12))
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v3019)))
	v3021 = *(*int32)(unsafe.Add(mBase, uint32(v3019)+4))
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	F_check_sql_expr(m, v3020, v3021, v3022+int32(1), l1)
	mBase = m.M
	v3026 = m.ExcPending
	if v3026 != 0 {
		goto L46
	} else {
		goto L919
	}
L919:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2926
	v9708 = v233
	goto L12
L920:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3029))) = int32(24)
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v3034 = int32(0)
	if v3033 < v3034 {
		v3079 = v3034
		goto L922
	} else {
		goto L923
	}
L921:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3029)+4)) = v3079
	v3083 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v3083)+520))
	v3086 = v3084 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3083)+520)) = v3086
	*(*int32)(unsafe.Add(mBase, uint32(v3029)+8)) = v3086
	v3091 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v3093 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(289), v3091, v3093, l1)
	mBase = m.M
	v3095 = m.ExcPending
	if v3095 != 0 {
		goto L46
	} else {
		goto L934
	}
L922:
	;
	goto L921
L923:
	;
	v3039 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3040 = *(*int32)(unsafe.Add(mBase, uint32(v3039)+60))
	if v3040 == int32(0) {
		v3079 = v3034
		goto L922
	} else {
		goto L924
	}
L924:
	;
	v3043 = v3033 + v3040
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(v3039)+188))
	if base.Ui32(v3044) <= base.Ui32(v3043) {
		goto L926
	} else {
		goto L927
	}
L925:
	;
	v3054 = *(*int32)(unsafe.Add(mBase, uint32(v3039)+196))
	if base.B2i32(v3053 == int32(0))|base.B2i32(base.Ui32(v3043) <= base.Ui32(v3053)) != 0 {
		v3079 = v3054
		goto L922
	} else {
		goto L929
	}
L926:
	;
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v3039)+192))
	v3053 = v3046
	goto L925
L927:
	;
	goto L928
L928:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3039)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3039)+188)) = v3040
	v3051 = F_strchr(m, v3040, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3039)+192)) = v3051
	v3053 = v3051
	goto L925
L929:
	;
	v3059 = v3053
	v3062 = v3054
	goto L930
L930:
	;
	v3064 = int32(1)
	v3065 = v3062 + v3064
	*(*int32)(unsafe.Add(mBase, uint32(v3039)+196)) = v3065
	v3068 = v3059 + v3064
	*(*int32)(unsafe.Add(mBase, uint32(v3039)+188)) = v3068
	v3071 = F_strchr(m, v3068, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3039)+192)) = v3071
	if v3071 == int32(0) {
		v3079 = v3065
		goto L922
	} else {
		goto L932
	}
L931:
	;
	v3079 = v3065
	goto L922
L932:
	;
	if base.Ui32(v3071) < base.Ui32(v3043) {
		v3059 = v3071
		v3062 = v3065
		goto L930
	} else {
		goto L933
	}
L933:
	;
	goto L931
L934:
	;
	v3096 = F_read_sql_stmt(m, v3091, v3093, l1)
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L46
	} else {
		goto L935
	}
L935:
	;
	v3098 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3029)+16)) = uint8(v3098)
	*(*int32)(unsafe.Add(mBase, uint32(v3029)+12)) = v3096
	v3102 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3102)+524)) = uint8(v3098)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3029
	v9708 = v233
	goto L12
L936:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3107))) = int32(24)
	v3111 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v3112 = int32(0)
	if v3111 < v3112 {
		v3157 = v3112
		goto L938
	} else {
		goto L939
	}
L937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3107)+4)) = v3157
	v3161 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v3161)+520))
	v3164 = v3162 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3161)+520)) = v3164
	*(*int32)(unsafe.Add(mBase, uint32(v3107)+8)) = v3164
	v3169 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v3171 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(309), v3169, v3171, l1)
	mBase = m.M
	v3173 = m.ExcPending
	if v3173 != 0 {
		goto L46
	} else {
		goto L950
	}
L938:
	;
	goto L937
L939:
	;
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3118 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+60))
	if v3118 == int32(0) {
		v3157 = v3112
		goto L938
	} else {
		goto L940
	}
L940:
	;
	v3121 = v3111 + v3118
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+188))
	if base.Ui32(v3122) <= base.Ui32(v3121) {
		goto L942
	} else {
		goto L943
	}
L941:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+196))
	if base.B2i32(v3131 == int32(0))|base.B2i32(base.Ui32(v3121) <= base.Ui32(v3131)) != 0 {
		v3157 = v3132
		goto L938
	} else {
		goto L945
	}
L942:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+192))
	v3131 = v3124
	goto L941
L943:
	;
	goto L944
L944:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+188)) = v3118
	v3129 = F_strchr(m, v3118, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+192)) = v3129
	v3131 = v3129
	goto L941
L945:
	;
	v3137 = v3131
	v3140 = v3132
	goto L946
L946:
	;
	v3142 = int32(1)
	v3143 = v3140 + v3142
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+196)) = v3143
	v3146 = v3137 + v3142
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+188)) = v3146
	v3149 = F_strchr(m, v3146, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3117)+192)) = v3149
	if v3149 == int32(0) {
		v3157 = v3143
		goto L938
	} else {
		goto L948
	}
L947:
	;
	v3157 = v3143
	goto L938
L948:
	;
	if base.Ui32(v3149) < base.Ui32(v3121) {
		v3137 = v3149
		v3140 = v3143
		goto L946
	} else {
		goto L949
	}
L949:
	;
	goto L947
L950:
	;
	v3174 = F_read_sql_stmt(m, v3169, v3171, l1)
	mBase = m.M
	v3175 = m.ExcPending
	if v3175 != 0 {
		goto L46
	} else {
		goto L951
	}
L951:
	;
	v3176 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3107)+16)) = uint8(v3176)
	*(*int32)(unsafe.Add(mBase, uint32(v3107)+12)) = v3174
	v3180 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v3181 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3180)+524)) = uint8(v3181)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3107
	v9708 = v233
	goto L12
L952:
	;
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	if v3188 == int32(0) {
		goto L31
	} else {
		goto L955
	}
L953:
	;
	v3199 = int32(3)
	goto L954
L954:
	;
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_check_assignable(m, v3200, v3201, l1)
	mBase = m.M
	v3203 = m.ExcPending
	if v3203 != 0 {
		goto L46
	} else {
		goto L957
	}
L955:
	;
	v3191 = *(*int32)(unsafe.Add(mBase, uint32(v3188)+4))
	if base.Ui32(int32(3)) <= base.Ui32(v3191-int32(1)) {
		goto L31
	} else {
		goto L956
	}
L956:
	;
	v3199 = v3191 + int32(2)
	goto L954
L957:
	;
	v3205 = F_palloc0(m, int32(20))
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L46
	} else {
		goto L958
	}
L958:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3205))) = int32(1)
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v3210 = int32(0)
	if v3209 < v3210 {
		v3255 = v3210
		goto L960
	} else {
		goto L961
	}
L959:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+4)) = v3255
	v3259 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v3259)+520))
	v3262 = v3260 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3259)+520)) = v3262
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+8)) = v3262
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v3265)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+12)) = v3266
	v3270 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v3272 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(277), v3270, v3272, l1)
	mBase = m.M
	v3274 = m.ExcPending
	if v3274 != 0 {
		goto L46
	} else {
		goto L972
	}
L960:
	;
	goto L959
L961:
	;
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+60))
	if v3216 == int32(0) {
		v3255 = v3210
		goto L960
	} else {
		goto L962
	}
L962:
	;
	v3219 = v3209 + v3216
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+188))
	if base.Ui32(v3220) <= base.Ui32(v3219) {
		goto L964
	} else {
		goto L965
	}
L963:
	;
	v3230 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+196))
	if base.B2i32(v3229 == int32(0))|base.B2i32(base.Ui32(v3219) <= base.Ui32(v3229)) != 0 {
		v3255 = v3230
		goto L960
	} else {
		goto L967
	}
L964:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+192))
	v3229 = v3222
	goto L963
L965:
	;
	goto L966
L966:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+188)) = v3216
	v3227 = F_strchr(m, v3216, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+192)) = v3227
	v3229 = v3227
	goto L963
L967:
	;
	v3235 = v3229
	v3238 = v3230
	goto L968
L968:
	;
	v3240 = int32(1)
	v3241 = v3238 + v3240
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+196)) = v3241
	v3244 = v3235 + v3240
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+188)) = v3244
	v3247 = F_strchr(m, v3244, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+192)) = v3247
	if v3247 == int32(0) {
		v3255 = v3241
		goto L960
	} else {
		goto L970
	}
L969:
	;
	v3255 = v3241
	goto L960
L970:
	;
	if base.Ui32(v3247) < base.Ui32(v3219) {
		v3235 = v3247
		v3238 = v3241
		goto L968
	} else {
		goto L971
	}
L971:
	;
	goto L969
L972:
	;
	v3276 = int32(0)
	v3283 = F_read_sql_construct(m, int32(59), v3276, v3276, int32(_a_F_plpgsql_yyparse_16), v3199, v3276, int32(1), v3276, v3276, v3270, v3272, l1)
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L46
	} else {
		goto L973
	}
L973:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+16)) = v3283
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3288 = *(*int32)(unsafe.Add(mBase, uint32(v3287)))
	if v3288 != 0 {
		goto L974
	} else {
		goto L975
	}
L974:
	;
	v3292 = int32(-1)
	v3293 = int32(0)
	goto L976
L975:
	;
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v3287)+4))
	v3292 = v3290
	v3293 = int32(1)
	goto L976
L976:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3283)+20)) = uint8(v3293)
	*(*int32)(unsafe.Add(mBase, uint32(v3283)+16)) = v3292
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3205
	v9708 = v233
	goto L12
L977:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3298))) = int32(19)
	v3303 = v147 - int32(16)
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(v3303)))
	v3305 = int32(0)
	if v3304 < v3305 {
		v3350 = v3305
		goto L979
	} else {
		goto L980
	}
L978:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3298)+4)) = v3350
	v3354 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3354)+520))
	v3357 = v3355 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3354)+520)) = v3357
	*(*int32)(unsafe.Add(mBase, uint32(v3298)+8)) = v3357
	v3362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148-int32(48)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3298)+12)) = uint8(v3362)
	v3366 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v3298)+16)) = v3366
	if v3366 == int32(0) {
		goto L991
	} else {
		goto L992
	}
L979:
	;
	goto L978
L980:
	;
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v3310)+60))
	if v3311 == int32(0) {
		v3350 = v3305
		goto L979
	} else {
		goto L981
	}
L981:
	;
	v3314 = v3304 + v3311
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(v3310)+188))
	if base.Ui32(v3315) <= base.Ui32(v3314) {
		goto L983
	} else {
		goto L984
	}
L982:
	;
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(v3310)+196))
	if base.B2i32(v3324 == int32(0))|base.B2i32(base.Ui32(v3314) <= base.Ui32(v3324)) != 0 {
		v3350 = v3325
		goto L979
	} else {
		goto L986
	}
L983:
	;
	v3317 = *(*int32)(unsafe.Add(mBase, uint32(v3310)+192))
	v3324 = v3317
	goto L982
L984:
	;
	goto L985
L985:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+188)) = v3311
	v3322 = F_strchr(m, v3311, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+192)) = v3322
	v3324 = v3322
	goto L982
L986:
	;
	v3330 = v3324
	v3333 = v3325
	goto L987
L987:
	;
	v3335 = int32(1)
	v3336 = v3333 + v3335
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+196)) = v3336
	v3339 = v3330 + v3335
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+188)) = v3339
	v3342 = F_strchr(m, v3339, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3310)+192)) = v3342
	if v3342 == int32(0) {
		v3350 = v3336
		goto L979
	} else {
		goto L989
	}
L988:
	;
	v3350 = v3336
	goto L979
L989:
	;
	if base.Ui32(v3342) < base.Ui32(v3314) {
		v3330 = v3342
		v3333 = v3336
		goto L987
	} else {
		goto L990
	}
L990:
	;
	goto L988
L991:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3298
	v9708 = v233
	goto L12
L992:
	;
	v3370 = *(*int32)(unsafe.Add(mBase, uint32(v3366)+4))
	if v3370 <= int32(0) {
		goto L991
	} else {
		goto L993
	}
L993:
	;
	v3373 = int32(0)
	if v3373 < v3370 {
		goto L994
	} else {
		goto L995
	}
L994:
	;
	v3377 = v3370
	goto L996
L995:
	;
	v3377 = v3373
	goto L996
L996:
	;
	v3378 = *(*int32)(unsafe.Add(mBase, uint32(v3366)+12))
	v3383 = v3373
	goto L997
L997:
	;
	v3407 = *(*int32)(unsafe.Add(mBase, uint32(v3378+v3383<<(uint(int32(2))%32))))
	v3408 = *(*int32)(unsafe.Add(mBase, uint32(v3407)))
	if base.Ui32(int32(10)) <= base.Ui32(v3408-int32(3)) {
		goto L1001
	} else {
		goto L1002
	}
L998:
	;
	goto L991
L999:
	;
	v3498 = v3383 + int32(1)
	if v3498 != v3377 {
		v3383 = v3498
		goto L997
	} else {
		goto L1028
	}
L1000:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v3484 = m.ExcPending
	if v3484 != 0 {
		goto L46
	} else {
		goto L1025
	}
L1001:
	;
	switch v3408 {
	case 0, 1:
		goto L1004
	case 2:
		goto L999
	default:
		goto L1000
	}
L1002:
	;
	goto L1003
L1003:
	;
	if v3362&int32(1) != 0 {
		goto L999
	} else {
		goto L1015
	}
L1004:
	;
	if v3362&int32(1) == int32(0) {
		goto L999
	} else {
		goto L1005
	}
L1005:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v3420 = m.ExcPending
	if v3420 != 0 {
		goto L46
	} else {
		goto L1006
	}
L1006:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L46
	} else {
		goto L1007
	}
L1007:
	;
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v3407)))
	if base.Ui32(int32(12)) < base.Ui32(v3424) {
		goto L1009
	} else {
		goto L1010
	}
L1008:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+144)) = v3433
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_35), v28+int32(144))
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L46
	} else {
		goto L1012
	}
L1009:
	;
	v3433 = int32(_a_F_plpgsql_yyparse_36)
	goto L1008
L1010:
	;
	goto L1011
L1011:
	;
	v3432 = *(*int32)(unsafe.Add(mBase, uint32(v3424<<(uint(int32(2))%32))+uint32(_c_F_plpgsql_yyparse[25])))
	v3433 = v3432
	goto L1008
L1012:
	;
	v3440 = *(*int32)(unsafe.Add(mBase, uint32(v3303)))
	v3441 = F_plpgsql_scanner_errposition(m, v3440, l1)
	mBase = m.M
	v3442 = m.ExcPending
	if v3442 != 0 {
		goto L46
	} else {
		goto L1013
	}
L1013:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1042), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L46
	} else {
		goto L1014
	}
L1014:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1015:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v3453 = m.ExcPending
	if v3453 != 0 {
		goto L46
	} else {
		goto L1016
	}
L1016:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3456 = m.ExcPending
	if v3456 != 0 {
		goto L46
	} else {
		goto L1017
	}
L1017:
	;
	v3457 = *(*int32)(unsafe.Add(mBase, uint32(v3407)))
	if base.Ui32(int32(12)) < base.Ui32(v3457) {
		goto L1019
	} else {
		goto L1020
	}
L1018:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v3466
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_37), v28+int32(160))
	mBase = m.M
	v3472 = m.ExcPending
	if v3472 != 0 {
		goto L46
	} else {
		goto L1022
	}
L1019:
	;
	v3466 = int32(_a_F_plpgsql_yyparse_36)
	goto L1018
L1020:
	;
	goto L1021
L1021:
	;
	v3465 = *(*int32)(unsafe.Add(mBase, uint32(v3457<<(uint(int32(2))%32))+uint32(_c_F_plpgsql_yyparse[25])))
	v3466 = v3465
	goto L1018
L1022:
	;
	v3473 = *(*int32)(unsafe.Add(mBase, uint32(v3303)))
	v3474 = F_plpgsql_scanner_errposition(m, v3473, l1)
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L46
	} else {
		goto L1023
	}
L1023:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1060), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v3480 = m.ExcPending
	if v3480 != 0 {
		goto L46
	} else {
		goto L1024
	}
L1024:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1025:
	;
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3407)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v3485
	F_errmsg_internal(m, int32(_a_F_plpgsql_yyparse_38), v28+int32(128))
	mBase = m.M
	v3491 = m.ExcPending
	if v3491 != 0 {
		goto L46
	} else {
		goto L1026
	}
L1026:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1067), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v3496 = m.ExcPending
	if v3496 != 0 {
		goto L46
	} else {
		goto L1027
	}
L1027:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1028:
	;
	goto L998
L1029:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3536
	v9708 = v233
	goto L12
L1030:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3545
	v9708 = v233
	goto L12
L1031:
	;
	v3553 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v3554 = *(*int32)(unsafe.Add(mBase, uint32(v3553)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3549)+4)) = v3554
	v3556 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v3549))) = v3556
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3549
	v9708 = v233
	goto L12
L1032:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_39))
	mBase = m.M
	v3987 = m.ExcPending
	if v3987 != 0 {
		goto L46
	} else {
		goto L1172
	}
L1033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(6)
	v9708 = v233
	goto L12
L1034:
	;
	v3948 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v3948 == int32(0) {
		goto L1032
	} else {
		goto L1163
	}
L1035:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(12)
	v9708 = v233
	goto L12
L1036:
	;
	v3918 = int32(_a_F_plpgsql_yyparse_40)
	v3921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3885))))
	v3924 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[26])))
	if base.B2i32(v3921 == int32(0))|base.B2i32(v3921 != v3924) != 0 {
		v3942 = v3921
		v3943 = v3924
		goto L1156
	} else {
		goto L1157
	}
L1037:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(11)
	v9708 = v233
	goto L12
L1038:
	;
	v3885 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v3885 == int32(0) {
		goto L1032
	} else {
		goto L1146
	}
L1039:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(10)
	v9708 = v233
	goto L12
L1040:
	;
	v3855 = int32(_a_F_plpgsql_yyparse_41)
	v3858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3822))))
	v3861 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[27])))
	if base.B2i32(v3858 == int32(0))|base.B2i32(v3858 != v3861) != 0 {
		v3879 = v3858
		v3880 = v3861
		goto L1139
	} else {
		goto L1140
	}
L1041:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(9)
	v9708 = v233
	goto L12
L1042:
	;
	v3822 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v3822 == int32(0) {
		goto L1032
	} else {
		goto L1129
	}
L1043:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(8)
	v9708 = v233
	goto L12
L1044:
	;
	v3792 = int32(_a_F_plpgsql_yyparse_42)
	v3795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3759))))
	v3798 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[28])))
	if base.B2i32(v3795 == int32(0))|base.B2i32(v3795 != v3798) != 0 {
		v3816 = v3795
		v3817 = v3798
		goto L1122
	} else {
		goto L1123
	}
L1045:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(7)
	v9708 = v233
	goto L12
L1046:
	;
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v3759 == int32(0) {
		goto L1032
	} else {
		goto L1112
	}
L1047:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(3)
	v9708 = v233
	goto L12
L1048:
	;
	v3729 = int32(_a_F_plpgsql_yyparse_43)
	v3732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3696))))
	v3735 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[29])))
	if base.B2i32(v3732 == int32(0))|base.B2i32(v3732 != v3735) != 0 {
		v3753 = v3732
		v3754 = v3735
		goto L1105
	} else {
		goto L1106
	}
L1049:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(5)
	v9708 = v233
	goto L12
L1050:
	;
	v3696 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v3696 == int32(0) {
		goto L1032
	} else {
		goto L1095
	}
L1051:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(4)
	v9708 = v233
	goto L12
L1052:
	;
	v3666 = int32(_a_F_plpgsql_yyparse_44)
	v3669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3633))))
	v3672 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[30])))
	if base.B2i32(v3669 == int32(0))|base.B2i32(v3669 != v3672) != 0 {
		v3690 = v3669
		v3691 = v3672
		goto L1088
	} else {
		goto L1089
	}
L1053:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(2)
	v9708 = v233
	goto L12
L1054:
	;
	v3633 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v3633 == int32(0) {
		goto L1032
	} else {
		goto L1078
	}
L1055:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(1)
	v9708 = v233
	goto L12
L1056:
	;
	v3603 = int32(_a_F_plpgsql_yyparse_45)
	v3606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3570))))
	v3609 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[31])))
	if base.B2i32(v3606 == int32(0))|base.B2i32(v3606 != v3609) != 0 {
		v3627 = v3606
		v3628 = v3609
		goto L1071
	} else {
		goto L1072
	}
L1057:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L1058:
	;
	v3567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v3567&int32(1) != 0 {
		goto L1032
	} else {
		goto L1060
	}
L1059:
	;
	switch v3563 - int32(277) {
	case 0:
		goto L1058
	default:
		goto L1032
	case 18:
		goto L1045
	case 22:
		goto L1043
	case 62:
		goto L1039
	case 73:
		goto L1053
	case 74:
		goto L1041
	case 75:
		goto L1047
	case 76:
		goto L1051
	case 77:
		goto L1049
	case 78:
		goto L1055
	case 85:
		goto L1033
	case 88:
		goto L1057
	case 91:
		goto L1035
	case 98:
		goto L1037
	}
L1060:
	;
	v3570 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v3570 == int32(0) {
		goto L1032
	} else {
		goto L1061
	}
L1061:
	;
	v3573 = int32(_a_F_plpgsql_yyparse_46)
	v3576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3570))))
	v3579 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[32])))
	if base.B2i32(v3576 == int32(0))|base.B2i32(v3576 != v3579) != 0 {
		v3597 = v3576
		v3598 = v3579
		goto L1063
	} else {
		goto L1064
	}
L1062:
	;
	if v3597-v3598 != 0 {
		goto L1056
	} else {
		goto L1069
	}
L1063:
	;
	goto L1062
L1064:
	;
	v3582 = v3570
	v3583 = v3573
	goto L1065
L1065:
	;
	v3586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3583)+1)))
	v3587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3582)+1)))
	if v3587 == int32(0) {
		v3597 = v3587
		v3598 = v3586
		goto L1063
	} else {
		goto L1067
	}
L1066:
	;
	v3597 = v3587
	v3598 = v3586
	goto L1063
L1067:
	;
	v3590 = int32(1)
	if v3587 == v3586 {
		v3582 = v3582 + v3590
		v3583 = v3583 + v3590
		goto L1065
	} else {
		goto L1068
	}
L1068:
	;
	goto L1066
L1069:
	;
	goto L1057
L1070:
	;
	if v3627-v3628 != 0 {
		goto L1054
	} else {
		goto L1077
	}
L1071:
	;
	goto L1070
L1072:
	;
	v3612 = v3570
	v3613 = v3603
	goto L1073
L1073:
	;
	v3616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3613)+1)))
	v3617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3612)+1)))
	if v3617 == int32(0) {
		v3627 = v3617
		v3628 = v3616
		goto L1071
	} else {
		goto L1075
	}
L1074:
	;
	v3627 = v3617
	v3628 = v3616
	goto L1071
L1075:
	;
	v3620 = int32(1)
	if v3617 == v3616 {
		v3612 = v3612 + v3620
		v3613 = v3613 + v3620
		goto L1073
	} else {
		goto L1076
	}
L1076:
	;
	goto L1074
L1077:
	;
	goto L1055
L1078:
	;
	v3636 = int32(_a_F_plpgsql_yyparse_47)
	v3639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3633))))
	v3642 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[33])))
	if base.B2i32(v3639 == int32(0))|base.B2i32(v3639 != v3642) != 0 {
		v3660 = v3639
		v3661 = v3642
		goto L1080
	} else {
		goto L1081
	}
L1079:
	;
	if v3660-v3661 != 0 {
		goto L1052
	} else {
		goto L1086
	}
L1080:
	;
	goto L1079
L1081:
	;
	v3645 = v3633
	v3646 = v3636
	goto L1082
L1082:
	;
	v3649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3646)+1)))
	v3650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3645)+1)))
	if v3650 == int32(0) {
		v3660 = v3650
		v3661 = v3649
		goto L1080
	} else {
		goto L1084
	}
L1083:
	;
	v3660 = v3650
	v3661 = v3649
	goto L1080
L1084:
	;
	v3653 = int32(1)
	if v3650 == v3649 {
		v3645 = v3645 + v3653
		v3646 = v3646 + v3653
		goto L1082
	} else {
		goto L1085
	}
L1085:
	;
	goto L1083
L1086:
	;
	goto L1053
L1087:
	;
	if v3690-v3691 != 0 {
		goto L1050
	} else {
		goto L1094
	}
L1088:
	;
	goto L1087
L1089:
	;
	v3675 = v3633
	v3676 = v3666
	goto L1090
L1090:
	;
	v3679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3676)+1)))
	v3680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3675)+1)))
	if v3680 == int32(0) {
		v3690 = v3680
		v3691 = v3679
		goto L1088
	} else {
		goto L1092
	}
L1091:
	;
	v3690 = v3680
	v3691 = v3679
	goto L1088
L1092:
	;
	v3683 = int32(1)
	if v3680 == v3679 {
		v3675 = v3675 + v3683
		v3676 = v3676 + v3683
		goto L1090
	} else {
		goto L1093
	}
L1093:
	;
	goto L1091
L1094:
	;
	goto L1051
L1095:
	;
	v3699 = int32(_a_F_plpgsql_yyparse_48)
	v3702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3696))))
	v3705 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[34])))
	if base.B2i32(v3702 == int32(0))|base.B2i32(v3702 != v3705) != 0 {
		v3723 = v3702
		v3724 = v3705
		goto L1097
	} else {
		goto L1098
	}
L1096:
	;
	if v3723-v3724 != 0 {
		goto L1048
	} else {
		goto L1103
	}
L1097:
	;
	goto L1096
L1098:
	;
	v3708 = v3696
	v3709 = v3699
	goto L1099
L1099:
	;
	v3712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3709)+1)))
	v3713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3708)+1)))
	if v3713 == int32(0) {
		v3723 = v3713
		v3724 = v3712
		goto L1097
	} else {
		goto L1101
	}
L1100:
	;
	v3723 = v3713
	v3724 = v3712
	goto L1097
L1101:
	;
	v3716 = int32(1)
	if v3713 == v3712 {
		v3708 = v3708 + v3716
		v3709 = v3709 + v3716
		goto L1099
	} else {
		goto L1102
	}
L1102:
	;
	goto L1100
L1103:
	;
	goto L1049
L1104:
	;
	if v3753-v3754 != 0 {
		goto L1046
	} else {
		goto L1111
	}
L1105:
	;
	goto L1104
L1106:
	;
	v3738 = v3696
	v3739 = v3729
	goto L1107
L1107:
	;
	v3742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3739)+1)))
	v3743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3738)+1)))
	if v3743 == int32(0) {
		v3753 = v3743
		v3754 = v3742
		goto L1105
	} else {
		goto L1109
	}
L1108:
	;
	v3753 = v3743
	v3754 = v3742
	goto L1105
L1109:
	;
	v3746 = int32(1)
	if v3743 == v3742 {
		v3738 = v3738 + v3746
		v3739 = v3739 + v3746
		goto L1107
	} else {
		goto L1110
	}
L1110:
	;
	goto L1108
L1111:
	;
	goto L1047
L1112:
	;
	v3762 = int32(_a_F_plpgsql_yyparse_49)
	v3765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3759))))
	v3768 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[35])))
	if base.B2i32(v3765 == int32(0))|base.B2i32(v3765 != v3768) != 0 {
		v3786 = v3765
		v3787 = v3768
		goto L1114
	} else {
		goto L1115
	}
L1113:
	;
	if v3786-v3787 != 0 {
		goto L1044
	} else {
		goto L1120
	}
L1114:
	;
	goto L1113
L1115:
	;
	v3771 = v3759
	v3772 = v3762
	goto L1116
L1116:
	;
	v3775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3772)+1)))
	v3776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3771)+1)))
	if v3776 == int32(0) {
		v3786 = v3776
		v3787 = v3775
		goto L1114
	} else {
		goto L1118
	}
L1117:
	;
	v3786 = v3776
	v3787 = v3775
	goto L1114
L1118:
	;
	v3779 = int32(1)
	if v3776 == v3775 {
		v3771 = v3771 + v3779
		v3772 = v3772 + v3779
		goto L1116
	} else {
		goto L1119
	}
L1119:
	;
	goto L1117
L1120:
	;
	goto L1045
L1121:
	;
	if v3816-v3817 != 0 {
		goto L1042
	} else {
		goto L1128
	}
L1122:
	;
	goto L1121
L1123:
	;
	v3801 = v3759
	v3802 = v3792
	goto L1124
L1124:
	;
	v3805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3802)+1)))
	v3806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3801)+1)))
	if v3806 == int32(0) {
		v3816 = v3806
		v3817 = v3805
		goto L1122
	} else {
		goto L1126
	}
L1125:
	;
	v3816 = v3806
	v3817 = v3805
	goto L1122
L1126:
	;
	v3809 = int32(1)
	if v3806 == v3805 {
		v3801 = v3801 + v3809
		v3802 = v3802 + v3809
		goto L1124
	} else {
		goto L1127
	}
L1127:
	;
	goto L1125
L1128:
	;
	goto L1043
L1129:
	;
	v3825 = int32(_a_F_plpgsql_yyparse_50)
	v3828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3822))))
	v3831 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[36])))
	if base.B2i32(v3828 == int32(0))|base.B2i32(v3828 != v3831) != 0 {
		v3849 = v3828
		v3850 = v3831
		goto L1131
	} else {
		goto L1132
	}
L1130:
	;
	if v3849-v3850 != 0 {
		goto L1040
	} else {
		goto L1137
	}
L1131:
	;
	goto L1130
L1132:
	;
	v3834 = v3822
	v3835 = v3825
	goto L1133
L1133:
	;
	v3838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3835)+1)))
	v3839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3834)+1)))
	if v3839 == int32(0) {
		v3849 = v3839
		v3850 = v3838
		goto L1131
	} else {
		goto L1135
	}
L1134:
	;
	v3849 = v3839
	v3850 = v3838
	goto L1131
L1135:
	;
	v3842 = int32(1)
	if v3839 == v3838 {
		v3834 = v3834 + v3842
		v3835 = v3835 + v3842
		goto L1133
	} else {
		goto L1136
	}
L1136:
	;
	goto L1134
L1137:
	;
	goto L1041
L1138:
	;
	if v3879-v3880 != 0 {
		goto L1038
	} else {
		goto L1145
	}
L1139:
	;
	goto L1138
L1140:
	;
	v3864 = v3822
	v3865 = v3855
	goto L1141
L1141:
	;
	v3868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3865)+1)))
	v3869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3864)+1)))
	if v3869 == int32(0) {
		v3879 = v3869
		v3880 = v3868
		goto L1139
	} else {
		goto L1143
	}
L1142:
	;
	v3879 = v3869
	v3880 = v3868
	goto L1139
L1143:
	;
	v3872 = int32(1)
	if v3869 == v3868 {
		v3864 = v3864 + v3872
		v3865 = v3865 + v3872
		goto L1141
	} else {
		goto L1144
	}
L1144:
	;
	goto L1142
L1145:
	;
	goto L1039
L1146:
	;
	v3888 = int32(_a_F_plpgsql_yyparse_51)
	v3891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3885))))
	v3894 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[37])))
	if base.B2i32(v3891 == int32(0))|base.B2i32(v3891 != v3894) != 0 {
		v3912 = v3891
		v3913 = v3894
		goto L1148
	} else {
		goto L1149
	}
L1147:
	;
	if v3912-v3913 != 0 {
		goto L1036
	} else {
		goto L1154
	}
L1148:
	;
	goto L1147
L1149:
	;
	v3897 = v3885
	v3898 = v3888
	goto L1150
L1150:
	;
	v3901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3898)+1)))
	v3902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3897)+1)))
	if v3902 == int32(0) {
		v3912 = v3902
		v3913 = v3901
		goto L1148
	} else {
		goto L1152
	}
L1151:
	;
	v3912 = v3902
	v3913 = v3901
	goto L1148
L1152:
	;
	v3905 = int32(1)
	if v3902 == v3901 {
		v3897 = v3897 + v3905
		v3898 = v3898 + v3905
		goto L1150
	} else {
		goto L1153
	}
L1153:
	;
	goto L1151
L1154:
	;
	goto L1037
L1155:
	;
	if v3942-v3943 != 0 {
		goto L1034
	} else {
		goto L1162
	}
L1156:
	;
	goto L1155
L1157:
	;
	v3927 = v3885
	v3928 = v3918
	goto L1158
L1158:
	;
	v3931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3928)+1)))
	v3932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927)+1)))
	if v3932 == int32(0) {
		v3942 = v3932
		v3943 = v3931
		goto L1156
	} else {
		goto L1160
	}
L1159:
	;
	v3942 = v3932
	v3943 = v3931
	goto L1156
L1160:
	;
	v3935 = int32(1)
	if v3932 == v3931 {
		v3927 = v3927 + v3935
		v3928 = v3928 + v3935
		goto L1158
	} else {
		goto L1161
	}
L1161:
	;
	goto L1159
L1162:
	;
	goto L1035
L1163:
	;
	v3951 = int32(_a_F_plpgsql_yyparse_52)
	v3954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3948))))
	v3957 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[38])))
	if base.B2i32(v3954 == int32(0))|base.B2i32(v3954 != v3957) != 0 {
		v3975 = v3954
		v3976 = v3957
		goto L1165
	} else {
		goto L1166
	}
L1164:
	;
	if v3975-v3976 != 0 {
		goto L1032
	} else {
		goto L1171
	}
L1165:
	;
	goto L1164
L1166:
	;
	v3960 = v3948
	v3961 = v3951
	goto L1167
L1167:
	;
	v3964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3961)+1)))
	v3965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3960)+1)))
	if v3965 == int32(0) {
		v3975 = v3965
		v3976 = v3964
		goto L1165
	} else {
		goto L1169
	}
L1168:
	;
	v3975 = v3965
	v3976 = v3964
	goto L1165
L1169:
	;
	v3968 = int32(1)
	if v3965 == v3964 {
		v3960 = v3960 + v3968
		v3961 = v3961 + v3968
		goto L1167
	} else {
		goto L1170
	}
L1170:
	;
	goto L1168
L1171:
	;
	goto L1033
L1172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1173:
	;
	v3994 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v3995 = m.ExcPending
	if v3995 != 0 {
		goto L46
	} else {
		goto L1174
	}
L1174:
	;
	if v3994 == int32(91) {
		goto L30
	} else {
		goto L1175
	}
L1175:
	;
	v3998 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3999 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_check_assignable(m, v3998, v3999, l1)
	mBase = m.M
	v4001 = m.ExcPending
	if v4001 != 0 {
		goto L46
	} else {
		goto L1176
	}
L1176:
	;
	v4002 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4002
	v9708 = v233
	goto L12
L1177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4005))) = int32(2)
	v4011 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(28))))
	v4012 = int32(0)
	if v4011 < v4012 {
		v4057 = v4012
		goto L1179
	} else {
		goto L1180
	}
L1178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4005)+4)) = v4057
	v4061 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v4062 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+520))
	v4064 = v4062 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4061)+520)) = v4064
	*(*int32)(unsafe.Add(mBase, uint32(v4005)+8)) = v4064
	v4069 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(96))))
	*(*int32)(unsafe.Add(mBase, uint32(v4005)+12)) = v4069
	v4073 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(80))))
	*(*int32)(unsafe.Add(mBase, uint32(v4005)+16)) = v4073
	v4077 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v4005)+20)) = v4077
	v4081 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	*(*int32)(unsafe.Add(mBase, uint32(v4005)+24)) = v4081
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4005
	v9708 = v233
	goto L12
L1179:
	;
	goto L1178
L1180:
	;
	v4017 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4018 = *(*int32)(unsafe.Add(mBase, uint32(v4017)+60))
	if v4018 == int32(0) {
		v4057 = v4012
		goto L1179
	} else {
		goto L1181
	}
L1181:
	;
	v4021 = v4011 + v4018
	v4022 = *(*int32)(unsafe.Add(mBase, uint32(v4017)+188))
	if base.Ui32(v4022) <= base.Ui32(v4021) {
		goto L1183
	} else {
		goto L1184
	}
L1182:
	;
	v4032 = *(*int32)(unsafe.Add(mBase, uint32(v4017)+196))
	if base.B2i32(v4031 == int32(0))|base.B2i32(base.Ui32(v4021) <= base.Ui32(v4031)) != 0 {
		v4057 = v4032
		goto L1179
	} else {
		goto L1186
	}
L1183:
	;
	v4024 = *(*int32)(unsafe.Add(mBase, uint32(v4017)+192))
	v4031 = v4024
	goto L1182
L1184:
	;
	goto L1185
L1185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4017)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4017)+188)) = v4018
	v4029 = F_strchr(m, v4018, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4017)+192)) = v4029
	v4031 = v4029
	goto L1182
L1186:
	;
	v4037 = v4031
	v4040 = v4032
	goto L1187
L1187:
	;
	v4042 = int32(1)
	v4043 = v4040 + v4042
	*(*int32)(unsafe.Add(mBase, uint32(v4017)+196)) = v4043
	v4046 = v4037 + v4042
	*(*int32)(unsafe.Add(mBase, uint32(v4017)+188)) = v4046
	v4049 = F_strchr(m, v4046, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4017)+192)) = v4049
	if v4049 == int32(0) {
		v4057 = v4043
		goto L1179
	} else {
		goto L1189
	}
L1188:
	;
	v4057 = v4043
	goto L1179
L1189:
	;
	if base.Ui32(v4049) < base.Ui32(v4021) {
		v4037 = v4049
		v4040 = v4043
		goto L1187
	} else {
		goto L1190
	}
L1190:
	;
	goto L1188
L1191:
	;
	v4091 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4092 = int32(0)
	if v4091 < v4092 {
		v4137 = v4092
		goto L1193
	} else {
		goto L1194
	}
L1192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4087))) = v4137
	v4142 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4087)+4)) = v4142
	v4144 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4087)+8)) = v4144
	v4148 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	v4149 = F_lappend(m, v4148, v4087)
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		goto L46
	} else {
		goto L1205
	}
L1193:
	;
	goto L1192
L1194:
	;
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4098 = *(*int32)(unsafe.Add(mBase, uint32(v4097)+60))
	if v4098 == int32(0) {
		v4137 = v4092
		goto L1193
	} else {
		goto L1195
	}
L1195:
	;
	v4101 = v4091 + v4098
	v4102 = *(*int32)(unsafe.Add(mBase, uint32(v4097)+188))
	if base.Ui32(v4102) <= base.Ui32(v4101) {
		goto L1197
	} else {
		goto L1198
	}
L1196:
	;
	v4112 = *(*int32)(unsafe.Add(mBase, uint32(v4097)+196))
	if base.B2i32(v4111 == int32(0))|base.B2i32(base.Ui32(v4101) <= base.Ui32(v4111)) != 0 {
		v4137 = v4112
		goto L1193
	} else {
		goto L1200
	}
L1197:
	;
	v4104 = *(*int32)(unsafe.Add(mBase, uint32(v4097)+192))
	v4111 = v4104
	goto L1196
L1198:
	;
	goto L1199
L1199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+188)) = v4098
	v4109 = F_strchr(m, v4098, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+192)) = v4109
	v4111 = v4109
	goto L1196
L1200:
	;
	v4117 = v4111
	v4120 = v4112
	goto L1201
L1201:
	;
	v4122 = int32(1)
	v4123 = v4120 + v4122
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+196)) = v4123
	v4126 = v4117 + v4122
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+188)) = v4126
	v4129 = F_strchr(m, v4126, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4097)+192)) = v4129
	if v4129 == int32(0) {
		v4137 = v4123
		goto L1193
	} else {
		goto L1203
	}
L1202:
	;
	v4137 = v4123
	goto L1193
L1203:
	;
	if base.Ui32(v4129) < base.Ui32(v4101) {
		v4117 = v4129
		v4120 = v4123
		goto L1201
	} else {
		goto L1204
	}
L1204:
	;
	goto L1202
L1205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4149
	v9708 = v233
	goto L12
L1206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4173))) = int32(3)
	v4177 = int32(0)
	if v4158 < v4177 {
		v4222 = v4177
		goto L1208
	} else {
		goto L1209
	}
L1207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4173)+4)) = v4222
	v4226 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v4227 = *(*int32)(unsafe.Add(mBase, uint32(v4226)+520))
	v4229 = v4227 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4226)+520)) = v4229
	v4231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4173)+24)) = uint8(base.B2i32(v4167 != v4231))
	*(*int32)(unsafe.Add(mBase, uint32(v4173)+20)) = v4164
	*(*int32)(unsafe.Add(mBase, uint32(v4173)+16)) = v4231
	*(*int32)(unsafe.Add(mBase, uint32(v4173)+12)) = v4161
	*(*int32)(unsafe.Add(mBase, uint32(v4173)+8)) = v4229
	if v4167 == v4231 {
		v4247 = v4167
		goto L1220
	} else {
		goto L1221
	}
L1208:
	;
	goto L1207
L1209:
	;
	v4182 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v4182)+60))
	if v4183 == int32(0) {
		v4222 = v4177
		goto L1208
	} else {
		goto L1210
	}
L1210:
	;
	v4186 = v4158 + v4183
	v4187 = *(*int32)(unsafe.Add(mBase, uint32(v4182)+188))
	if base.Ui32(v4187) <= base.Ui32(v4186) {
		goto L1212
	} else {
		goto L1213
	}
L1211:
	;
	v4197 = *(*int32)(unsafe.Add(mBase, uint32(v4182)+196))
	if base.B2i32(v4196 == int32(0))|base.B2i32(base.Ui32(v4186) <= base.Ui32(v4196)) != 0 {
		v4222 = v4197
		goto L1208
	} else {
		goto L1215
	}
L1212:
	;
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v4182)+192))
	v4196 = v4189
	goto L1211
L1213:
	;
	goto L1214
L1214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4182)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4182)+188)) = v4183
	v4194 = F_strchr(m, v4183, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4182)+192)) = v4194
	v4196 = v4194
	goto L1211
L1215:
	;
	v4202 = v4196
	v4205 = v4197
	goto L1216
L1216:
	;
	v4207 = int32(1)
	v4208 = v4205 + v4207
	*(*int32)(unsafe.Add(mBase, uint32(v4182)+196)) = v4208
	v4211 = v4202 + v4207
	*(*int32)(unsafe.Add(mBase, uint32(v4182)+188)) = v4211
	v4214 = F_strchr(m, v4211, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4182)+192)) = v4214
	if v4214 == int32(0) {
		v4222 = v4208
		goto L1208
	} else {
		goto L1218
	}
L1217:
	;
	v4222 = v4208
	goto L1208
L1218:
	;
	if base.Ui32(v4214) < base.Ui32(v4186) {
		v4202 = v4214
		v4205 = v4208
		goto L1216
	} else {
		goto L1219
	}
L1219:
	;
	goto L1217
L1220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4173)+28)) = v4247
	if v4161 == int32(0) {
		goto L1224
	} else {
		goto L1225
	}
L1221:
	;
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+4))
	if v4241 != int32(1) {
		v4247 = v4167
		goto L1220
	} else {
		goto L1222
	}
L1222:
	;
	v4244 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+12))
	v4245 = *(*int32)(unsafe.Add(mBase, uint32(v4244)))
	if v4245 != 0 {
		v4247 = v4167
		goto L1220
	} else {
		goto L1223
	}
L1223:
	;
	v4247 = int32(0)
	goto L1220
L1224:
	;
	m.G0 = v4170 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4173
	v9708 = v233
	goto L12
L1225:
	;
	v4252 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v4170)+16)) = v4252
	v4255 = v4170 + int32(48)
	v4260 = F_pg_snprintf(m, v4255, int32(32), int32(_a_F_plpgsql_yyparse_53), v4170+int32(16))
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		goto L46
	} else {
		goto L1226
	}
L1226:
	;
	v4262 = int32(0)
	v4263 = *(*int32)(unsafe.Add(mBase, uint32(v4173)+4))
	v4268 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v4262, v4262)
	mBase = m.M
	v4269 = m.ExcPending
	if v4269 != 0 {
		goto L46
	} else {
		goto L1227
	}
L1227:
	;
	v4271 = F_plpgsql_build_variable(m, v4255, v4263, v4268, int32(1))
	mBase = m.M
	v4272 = m.ExcPending
	if v4272 != 0 {
		goto L46
	} else {
		goto L1228
	}
L1228:
	;
	v4273 = *(*int32)(unsafe.Add(mBase, uint32(v4271)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4173)+16)) = v4273
	if v4164 == int32(0) {
		goto L1224
	} else {
		goto L1229
	}
L1229:
	;
	v4277 = *(*int32)(unsafe.Add(mBase, uint32(v4164)+4))
	if v4277 <= int32(0) {
		goto L1224
	} else {
		goto L1230
	}
L1230:
	;
	v4285 = v4262
	goto L1231
L1231:
	;
	v4305 = *(*int32)(unsafe.Add(mBase, uint32(v4164)+12))
	v4309 = *(*int32)(unsafe.Add(mBase, uint32(v4305+v4285<<(uint(int32(2))%32))))
	v4310 = *(*int32)(unsafe.Add(mBase, uint32(v4309)+4))
	v4312 = v4170 + int32(32)
	F_initStringInfo(m, v4312)
	mBase = m.M
	v4314 = m.ExcPending
	if v4314 != 0 {
		goto L46
	} else {
		goto L1233
	}
L1232:
	;
	goto L1224
L1233:
	;
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(v4310)))
	*(*int32)(unsafe.Add(mBase, uint32(v4170)+4)) = v4315
	*(*int32)(unsafe.Add(mBase, uint32(v4170))) = v4170 + int32(48)
	F_appendStringInfo(m, v4312, int32(_a_F_plpgsql_yyparse_54), v4170)
	mBase = m.M
	v4322 = m.ExcPending
	if v4322 != 0 {
		goto L46
	} else {
		goto L1234
	}
L1234:
	;
	v4323 = *(*int32)(unsafe.Add(mBase, uint32(v4310)))
	F_pfree(m, v4323)
	mBase = m.M
	v4325 = m.ExcPending
	if v4325 != 0 {
		goto L46
	} else {
		goto L1235
	}
L1235:
	;
	v4326 = *(*int32)(unsafe.Add(mBase, uint32(v4170)+32))
	v4327 = F_pstrdup(m, v4326)
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L46
	} else {
		goto L1236
	}
L1236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4310))) = v4327
	v4331 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v4310)+12)) = v4331
	v4333 = *(*int32)(unsafe.Add(mBase, uint32(v4170)+32))
	F_pfree(m, v4333)
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L46
	} else {
		goto L1237
	}
L1237:
	;
	v4337 = v4285 + int32(1)
	v4338 = *(*int32)(unsafe.Add(mBase, uint32(v4164)+4))
	if v4337 < v4338 {
		v4285 = v4337
		goto L1231
	} else {
		goto L1238
	}
L1238:
	;
	goto L1232
L1239:
	;
	if v4374 != int32(384) {
		goto L1240
	} else {
		goto L1241
	}
L1240:
	;
	F_plpgsql_push_back_token(m, v4374, v4371, v4373, l1)
	mBase = m.M
	v4379 = m.ExcPending
	if v4379 != 0 {
		goto L46
	} else {
		goto L1243
	}
L1241:
	;
	v4391 = int32(0)
	goto L1242
L1242:
	;
	F_plpgsql_push_back_token(m, int32(384), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4398 = m.ExcPending
	if v4398 != 0 {
		goto L46
	} else {
		goto L1245
	}
L1243:
	;
	v4381 = int32(0)
	v4385 = int32(1)
	v4389 = F_read_sql_construct(m, int32(384), v4381, v4381, int32(_a_F_plpgsql_yyparse_55), int32(2), v4385, v4385, v4381, v4381, v4371, v4373, l1)
	mBase = m.M
	v4390 = m.ExcPending
	if v4390 != 0 {
		goto L46
	} else {
		goto L1244
	}
L1244:
	;
	v4391 = v4389
	goto L1242
L1245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4391
	v9708 = v233
	goto L12
L1246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4404
	v9708 = v233
	goto L12
L1247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4413
	v9708 = v233
	goto L12
L1248:
	;
	v4421 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4422 = int32(0)
	if v4421 < v4422 {
		v4467 = v4422
		goto L1250
	} else {
		goto L1251
	}
L1249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4417))) = v4467
	v4472 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4417)+4)) = v4472
	v4474 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4417)+8)) = v4474
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4417
	v9708 = v233
	goto L12
L1250:
	;
	goto L1249
L1251:
	;
	v4427 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4428 = *(*int32)(unsafe.Add(mBase, uint32(v4427)+60))
	if v4428 == int32(0) {
		v4467 = v4422
		goto L1250
	} else {
		goto L1252
	}
L1252:
	;
	v4431 = v4421 + v4428
	v4432 = *(*int32)(unsafe.Add(mBase, uint32(v4427)+188))
	if base.Ui32(v4432) <= base.Ui32(v4431) {
		goto L1254
	} else {
		goto L1255
	}
L1253:
	;
	v4442 = *(*int32)(unsafe.Add(mBase, uint32(v4427)+196))
	if base.B2i32(v4441 == int32(0))|base.B2i32(base.Ui32(v4431) <= base.Ui32(v4441)) != 0 {
		v4467 = v4442
		goto L1250
	} else {
		goto L1257
	}
L1254:
	;
	v4434 = *(*int32)(unsafe.Add(mBase, uint32(v4427)+192))
	v4441 = v4434
	goto L1253
L1255:
	;
	goto L1256
L1256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4427)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4427)+188)) = v4428
	v4439 = F_strchr(m, v4428, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4427)+192)) = v4439
	v4441 = v4439
	goto L1253
L1257:
	;
	v4447 = v4441
	v4450 = v4442
	goto L1258
L1258:
	;
	v4452 = int32(1)
	v4453 = v4450 + v4452
	*(*int32)(unsafe.Add(mBase, uint32(v4427)+196)) = v4453
	v4456 = v4447 + v4452
	*(*int32)(unsafe.Add(mBase, uint32(v4427)+188)) = v4456
	v4459 = F_strchr(m, v4456, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4427)+192)) = v4459
	if v4459 == int32(0) {
		v4467 = v4453
		goto L1250
	} else {
		goto L1260
	}
L1259:
	;
	v4467 = v4453
	goto L1250
L1260:
	;
	if base.Ui32(v4459) < base.Ui32(v4431) {
		v4447 = v4459
		v4450 = v4453
		goto L1258
	} else {
		goto L1261
	}
L1261:
	;
	goto L1259
L1262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4479
	v9708 = v233
	goto L12
L1263:
	;
	goto L1264
L1264:
	;
	v4481 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+188)) = v4481
	*(*int32)(unsafe.Add(mBase, uint32(v28)+244)) = v4481
	v4488 = F_list_make1_impl(m, int32(1), v28+int32(188))
	mBase = m.M
	v4489 = m.ExcPending
	if v4489 != 0 {
		goto L46
	} else {
		goto L1265
	}
L1265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4488
	v9708 = v233
	goto L12
L1266:
	;
	v4494 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4492))) = v4494
	v4498 = *(*int32)(unsafe.Add(mBase, uint32(v147-v4494)))
	v4499 = int32(0)
	if v4498 < v4499 {
		v4544 = v4499
		goto L1268
	} else {
		goto L1269
	}
L1267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4492)+4)) = v4544
	v4548 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v4549 = *(*int32)(unsafe.Add(mBase, uint32(v4548)+520))
	v4551 = v4549 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4548)+520)) = v4551
	*(*int32)(unsafe.Add(mBase, uint32(v4492)+8)) = v4551
	v4555 = v148 - int32(32)
	v4556 = *(*int32)(unsafe.Add(mBase, uint32(v4555)))
	*(*int32)(unsafe.Add(mBase, uint32(v4492)+12)) = v4556
	v4558 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4492)+16)) = v4558
	v4560 = *(*int32)(unsafe.Add(mBase, uint32(v4555)))
	v4561 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v4562 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v4560, v4561, v4562, l1)
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L46
	} else {
		goto L1280
	}
L1268:
	;
	goto L1267
L1269:
	;
	v4504 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4505 = *(*int32)(unsafe.Add(mBase, uint32(v4504)+60))
	if v4505 == int32(0) {
		v4544 = v4499
		goto L1268
	} else {
		goto L1270
	}
L1270:
	;
	v4508 = v4498 + v4505
	v4509 = *(*int32)(unsafe.Add(mBase, uint32(v4504)+188))
	if base.Ui32(v4509) <= base.Ui32(v4508) {
		goto L1272
	} else {
		goto L1273
	}
L1271:
	;
	v4519 = *(*int32)(unsafe.Add(mBase, uint32(v4504)+196))
	if base.B2i32(v4518 == int32(0))|base.B2i32(base.Ui32(v4508) <= base.Ui32(v4518)) != 0 {
		v4544 = v4519
		goto L1268
	} else {
		goto L1275
	}
L1272:
	;
	v4511 = *(*int32)(unsafe.Add(mBase, uint32(v4504)+192))
	v4518 = v4511
	goto L1271
L1273:
	;
	goto L1274
L1274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4504)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4504)+188)) = v4505
	v4516 = F_strchr(m, v4505, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4504)+192)) = v4516
	v4518 = v4516
	goto L1271
L1275:
	;
	v4524 = v4518
	v4527 = v4519
	goto L1276
L1276:
	;
	v4529 = int32(1)
	v4530 = v4527 + v4529
	*(*int32)(unsafe.Add(mBase, uint32(v4504)+196)) = v4530
	v4533 = v4524 + v4529
	*(*int32)(unsafe.Add(mBase, uint32(v4504)+188)) = v4533
	v4536 = F_strchr(m, v4533, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4504)+192)) = v4536
	if v4536 == int32(0) {
		v4544 = v4530
		goto L1268
	} else {
		goto L1278
	}
L1277:
	;
	v4544 = v4530
	goto L1268
L1278:
	;
	if base.Ui32(v4536) < base.Ui32(v4508) {
		v4524 = v4536
		v4527 = v4530
		goto L1276
	} else {
		goto L1279
	}
L1279:
	;
	goto L1277
L1280:
	;
	v4567 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v4568 = *(*int32)(unsafe.Add(mBase, uint32(v4567)))
	if v4568 != 0 {
		goto L1282
	} else {
		goto L1283
	}
L1281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4492
	v9708 = v233
	goto L12
L1282:
	;
	v4569 = v4567
	goto L1285
L1283:
	;
	v4572 = v4567
	goto L1284
L1284:
	;
	v4574 = *(*int32)(unsafe.Add(mBase, uint32(v4572)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])) = v4574
	goto L1281
L1285:
	;
	v4570 = *(*int32)(unsafe.Add(mBase, uint32(v4569)+8))
	v4571 = *(*int32)(unsafe.Add(mBase, uint32(v4570)))
	if v4571 != 0 {
		v4569 = v4570
		goto L1285
	} else {
		goto L1287
	}
L1286:
	;
	v4572 = v4570
	goto L1284
L1287:
	;
	goto L1286
L1288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4578))) = int32(5)
	v4584 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4585 = int32(0)
	if v4584 < v4585 {
		v4630 = v4585
		goto L1290
	} else {
		goto L1291
	}
L1289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4578)+4)) = v4630
	v4634 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v4635 = *(*int32)(unsafe.Add(mBase, uint32(v4634)+520))
	v4637 = v4635 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4634)+520)) = v4637
	*(*int32)(unsafe.Add(mBase, uint32(v4578)+8)) = v4637
	v4641 = v148 - int32(48)
	v4642 = *(*int32)(unsafe.Add(mBase, uint32(v4641)))
	*(*int32)(unsafe.Add(mBase, uint32(v4578)+12)) = v4642
	v4646 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4578)+16)) = v4646
	v4648 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4578)+20)) = v4648
	v4650 = *(*int32)(unsafe.Add(mBase, uint32(v4641)))
	v4651 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v4652 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v4650, v4651, v4652, l1)
	mBase = m.M
	v4654 = m.ExcPending
	if v4654 != 0 {
		goto L46
	} else {
		goto L1302
	}
L1290:
	;
	goto L1289
L1291:
	;
	v4590 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4591 = *(*int32)(unsafe.Add(mBase, uint32(v4590)+60))
	if v4591 == int32(0) {
		v4630 = v4585
		goto L1290
	} else {
		goto L1292
	}
L1292:
	;
	v4594 = v4584 + v4591
	v4595 = *(*int32)(unsafe.Add(mBase, uint32(v4590)+188))
	if base.Ui32(v4595) <= base.Ui32(v4594) {
		goto L1294
	} else {
		goto L1295
	}
L1293:
	;
	v4605 = *(*int32)(unsafe.Add(mBase, uint32(v4590)+196))
	if base.B2i32(v4604 == int32(0))|base.B2i32(base.Ui32(v4594) <= base.Ui32(v4604)) != 0 {
		v4630 = v4605
		goto L1290
	} else {
		goto L1297
	}
L1294:
	;
	v4597 = *(*int32)(unsafe.Add(mBase, uint32(v4590)+192))
	v4604 = v4597
	goto L1293
L1295:
	;
	goto L1296
L1296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4590)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4590)+188)) = v4591
	v4602 = F_strchr(m, v4591, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4590)+192)) = v4602
	v4604 = v4602
	goto L1293
L1297:
	;
	v4610 = v4604
	v4613 = v4605
	goto L1298
L1298:
	;
	v4615 = int32(1)
	v4616 = v4613 + v4615
	*(*int32)(unsafe.Add(mBase, uint32(v4590)+196)) = v4616
	v4619 = v4610 + v4615
	*(*int32)(unsafe.Add(mBase, uint32(v4590)+188)) = v4619
	v4622 = F_strchr(m, v4619, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4590)+192)) = v4622
	if v4622 == int32(0) {
		v4630 = v4616
		goto L1290
	} else {
		goto L1300
	}
L1299:
	;
	v4630 = v4616
	goto L1290
L1300:
	;
	if base.Ui32(v4622) < base.Ui32(v4594) {
		v4610 = v4622
		v4613 = v4616
		goto L1298
	} else {
		goto L1301
	}
L1301:
	;
	goto L1299
L1302:
	;
	v4657 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v4658 = *(*int32)(unsafe.Add(mBase, uint32(v4657)))
	if v4658 != 0 {
		goto L1304
	} else {
		goto L1305
	}
L1303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4578
	v9708 = v233
	goto L12
L1304:
	;
	v4659 = v4657
	goto L1307
L1305:
	;
	v4662 = v4657
	goto L1306
L1306:
	;
	v4664 = *(*int32)(unsafe.Add(mBase, uint32(v4662)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])) = v4664
	goto L1303
L1307:
	;
	v4660 = *(*int32)(unsafe.Add(mBase, uint32(v4659)+8))
	v4661 = *(*int32)(unsafe.Add(mBase, uint32(v4660)))
	if v4661 != 0 {
		v4659 = v4660
		goto L1307
	} else {
		goto L1309
	}
L1308:
	;
	v4662 = v4660
	goto L1306
L1309:
	;
	goto L1308
L1310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4669)+4)) = v4719
	v4723 = v148 - int32(48)
	v4724 = *(*int32)(unsafe.Add(mBase, uint32(v4723)))
	*(*int32)(unsafe.Add(mBase, uint32(v4669)+12)) = v4724
	if v4670 == int32(6) {
		goto L1323
	} else {
		goto L1324
	}
L1311:
	;
	goto L1310
L1312:
	;
	v4679 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4680 = *(*int32)(unsafe.Add(mBase, uint32(v4679)+60))
	if v4680 == int32(0) {
		v4719 = v4674
		goto L1311
	} else {
		goto L1313
	}
L1313:
	;
	v4683 = v4673 + v4680
	v4684 = *(*int32)(unsafe.Add(mBase, uint32(v4679)+188))
	if base.Ui32(v4684) <= base.Ui32(v4683) {
		goto L1315
	} else {
		goto L1316
	}
L1314:
	;
	v4694 = *(*int32)(unsafe.Add(mBase, uint32(v4679)+196))
	if base.B2i32(v4693 == int32(0))|base.B2i32(base.Ui32(v4683) <= base.Ui32(v4693)) != 0 {
		v4719 = v4694
		goto L1311
	} else {
		goto L1318
	}
L1315:
	;
	v4686 = *(*int32)(unsafe.Add(mBase, uint32(v4679)+192))
	v4693 = v4686
	goto L1314
L1316:
	;
	goto L1317
L1317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+188)) = v4680
	v4691 = F_strchr(m, v4680, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+192)) = v4691
	v4693 = v4691
	goto L1314
L1318:
	;
	v4699 = v4693
	v4702 = v4694
	goto L1319
L1319:
	;
	v4704 = int32(1)
	v4705 = v4702 + v4704
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+196)) = v4705
	v4708 = v4699 + v4704
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+188)) = v4708
	v4711 = F_strchr(m, v4708, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+192)) = v4711
	if v4711 == int32(0) {
		v4719 = v4705
		goto L1311
	} else {
		goto L1321
	}
L1320:
	;
	v4719 = v4705
	goto L1311
L1321:
	;
	if base.Ui32(v4711) < base.Ui32(v4683) {
		v4699 = v4711
		v4702 = v4705
		goto L1319
	} else {
		goto L1322
	}
L1322:
	;
	goto L1320
L1323:
	;
	v4730 = int32(36)
	goto L1325
L1324:
	;
	v4730 = int32(20)
	goto L1325
L1325:
	;
	v4732 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4669+v4730))) = v4732
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4669
	v4735 = *(*int32)(unsafe.Add(mBase, uint32(v4723)))
	v4736 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v4737 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v4735, v4736, v4737, l1)
	mBase = m.M
	v4739 = m.ExcPending
	if v4739 != 0 {
		goto L46
	} else {
		goto L1326
	}
L1326:
	;
	v4742 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v4743 = *(*int32)(unsafe.Add(mBase, uint32(v4742)))
	if v4743 != 0 {
		goto L1328
	} else {
		goto L1329
	}
L1327:
	;
	v9708 = v233
	goto L12
L1328:
	;
	v4744 = v4742
	goto L1331
L1329:
	;
	v4747 = v4742
	goto L1330
L1330:
	;
	v4749 = *(*int32)(unsafe.Add(mBase, uint32(v4747)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])) = v4749
	goto L1327
L1331:
	;
	v4745 = *(*int32)(unsafe.Add(mBase, uint32(v4744)+8))
	v4746 = *(*int32)(unsafe.Add(mBase, uint32(v4745)))
	if v4746 != 0 {
		v4744 = v4745
		goto L1331
	} else {
		goto L1333
	}
L1332:
	;
	v4747 = v4745
	goto L1330
L1333:
	;
	goto L1332
L1334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v4755
	v4758 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	if v4755 != int32(277) {
		goto L1337
	} else {
		goto L1338
	}
L1335:
	;
	v4978 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v4980 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, v4755, v4978, v4980, l1)
	mBase = m.M
	v4982 = m.ExcPending
	if v4982 != 0 {
		goto L46
	} else {
		goto L1380
	}
L1336:
	;
	if v4755 == int32(363) {
		goto L14
	} else {
		goto L1379
	}
L1337:
	;
	if v4755 != int32(317) {
		goto L1336
	} else {
		goto L1340
	}
L1338:
	;
	goto L1339
L1339:
	;
	v4891 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v4892 = *(*int32)(unsafe.Add(mBase, uint32(v4891)))
	if v4892 != 0 {
		goto L1358
	} else {
		goto L1359
	}
L1340:
	;
	v4765 = int32(0)
	v4768 = int32(1)
	v4773 = F_read_sql_construct(m, int32(336), int32(381), v4765, int32(_a_F_plpgsql_yyparse_56), int32(2), v4768, v4768, v4765, v28+int32(240), v4752, v4754, l1)
	mBase = m.M
	v4774 = m.ExcPending
	if v4774 != 0 {
		goto L46
	} else {
		goto L1341
	}
L1341:
	;
	v4776 = F_palloc0(m, int32(32))
	mBase = m.M
	v4777 = m.ExcPending
	if v4777 != 0 {
		goto L46
	} else {
		goto L1342
	}
L1342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4776))) = int32(18)
	v4781 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v4782 = *(*int32)(unsafe.Add(mBase, uint32(v4781)+520))
	v4784 = v4782 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4781)+520)) = v4784
	*(*int32)(unsafe.Add(mBase, uint32(v4776)+8)) = v4784
	v4787 = int32(4)
	v4788 = v147 - v4787
	v4790 = v148 - v4787
	v4791 = *(*int32)(unsafe.Add(mBase, uint32(v4790)))
	if v4791 != 0 {
		goto L1344
	} else {
		goto L1345
	}
L1343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4776)+24)) = v4773
	v4814 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	if v4814 == int32(381) {
		goto L1350
	} else {
		goto L1351
	}
L1344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4776)+16)) = v4791
	v4793 = *(*int32)(unsafe.Add(mBase, uint32(v4790)))
	v4794 = *(*int32)(unsafe.Add(mBase, uint32(v4788)))
	F_check_assignable(m, v4793, v4794, l1)
	mBase = m.M
	v4796 = m.ExcPending
	if v4796 != 0 {
		goto L46
	} else {
		goto L1347
	}
L1345:
	;
	goto L1346
L1346:
	;
	v4799 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(8))))
	if v4799 == int32(0) {
		goto L29
	} else {
		goto L1348
	}
L1347:
	;
	goto L1343
L1348:
	;
	v4804 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4807 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v4808 = *(*int32)(unsafe.Add(mBase, uint32(v4788)))
	v4809 = F_make_scalar_list1(m, v4804, v4799, v4807, v4808, l1)
	mBase = m.M
	v4810 = m.ExcPending
	if v4810 != 0 {
		goto L46
	} else {
		goto L1349
	}
L1349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4776)+16)) = v4809
	goto L1343
L1350:
	;
	goto L1353
L1351:
	;
	goto L1352
L1352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4776
	v9708 = v233
	goto L12
L1353:
	;
	v4844 = int32(0)
	v4847 = int32(1)
	v4856 = F_read_sql_construct(m, int32(44), int32(336), v4844, int32(_a_F_plpgsql_yyparse_57), int32(2), v4847, v4847, v4844, v28+int32(240), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4857 = m.ExcPending
	if v4857 != 0 {
		goto L46
	} else {
		goto L1355
	}
L1354:
	;
	goto L1352
L1355:
	;
	v4858 = *(*int32)(unsafe.Add(mBase, uint32(v4776)+28))
	v4859 = F_lappend(m, v4858, v4856)
	mBase = m.M
	v4860 = m.ExcPending
	if v4860 != 0 {
		goto L46
	} else {
		goto L1356
	}
L1356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4776)+28)) = v4859
	v4862 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	if v4862 == int32(44) {
		goto L1353
	} else {
		goto L1357
	}
L1357:
	;
	goto L1354
L1358:
	;
	v4941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v4941&int32(1) != 0 {
		goto L1335
	} else {
		goto L1369
	}
L1359:
	;
	v4893 = *(*int32)(unsafe.Add(mBase, uint32(v4891)+24))
	v4894 = *(*int32)(unsafe.Add(mBase, uint32(v4893)+4))
	if v4894 != int32(1790) {
		goto L1358
	} else {
		goto L1360
	}
L1360:
	;
	v4898 = F_palloc0(m, int32(32))
	mBase = m.M
	v4899 = m.ExcPending
	if v4899 != 0 {
		goto L46
	} else {
		goto L1361
	}
L1361:
	;
	v4900 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4898))) = v4900
	v4903 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v4904 = *(*int32)(unsafe.Add(mBase, uint32(v4903)+520))
	v4906 = v4904 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4903)+520)) = v4906
	*(*int32)(unsafe.Add(mBase, uint32(v4898)+8)) = v4906
	v4909 = *(*int32)(unsafe.Add(mBase, uint32(v4891)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4898)+24)) = v4909
	v4913 = *(*int32)(unsafe.Add(mBase, uint32(v148-v4900)))
	if v4913 != 0 {
		goto L1362
	} else {
		goto L1363
	}
L1362:
	;
	v4916 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(4))))
	if v4916 != 0 {
		goto L28
	} else {
		goto L1365
	}
L1363:
	;
	goto L1364
L1364:
	;
	v4917 = *(*int32)(unsafe.Add(mBase, uint32(v4891)+28))
	if v4917 == int32(0) {
		goto L27
	} else {
		goto L1366
	}
L1365:
	;
	goto L1364
L1366:
	;
	v4925 = F_read_cursor_args(m, v4891, int32(336), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4926 = m.ExcPending
	if v4926 != 0 {
		goto L46
	} else {
		goto L1367
	}
L1367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4898)+28)) = v4925
	v4930 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4933 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v4937 = F_plpgsql_build_record(m, v4930, v4933, int32(0), int32(2249), int32(1))
	mBase = m.M
	v4938 = m.ExcPending
	if v4938 != 0 {
		goto L46
	} else {
		goto L1368
	}
L1368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4898)+16)) = v4937
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4898
	v9708 = v233
	goto L12
L1369:
	;
	v4944 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v4944 == int32(0) {
		goto L1335
	} else {
		goto L1370
	}
L1370:
	;
	v4947 = int32(_a_F_plpgsql_yyparse_58)
	v4950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4944))))
	v4953 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[40])))
	if base.B2i32(v4950 == int32(0))|base.B2i32(v4950 != v4953) != 0 {
		v4971 = v4950
		v4972 = v4953
		goto L1372
	} else {
		goto L1373
	}
L1371:
	;
	if v4971-v4972 != 0 {
		goto L1335
	} else {
		goto L1378
	}
L1372:
	;
	goto L1371
L1373:
	;
	v4956 = v4944
	v4957 = v4947
	goto L1374
L1374:
	;
	v4960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4957)+1)))
	v4961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4956)+1)))
	if v4961 == int32(0) {
		v4971 = v4961
		v4972 = v4960
		goto L1372
	} else {
		goto L1376
	}
L1375:
	;
	v4971 = v4961
	v4972 = v4960
	goto L1372
L1376:
	;
	v4964 = int32(1)
	if v4961 == v4960 {
		v4956 = v4956 + v4964
		v4957 = v4957 + v4964
		goto L1374
	} else {
		goto L1377
	}
L1377:
	;
	goto L1375
L1378:
	;
	goto L14
L1379:
	;
	goto L1335
L1380:
	;
	v4983 = int32(0)
	v4995 = F_read_sql_construct(m, int32(269), int32(336), v4983, int32(_a_F_plpgsql_yyparse_8), v4983, int32(1), v4983, v28+int32(240), v28+int32(256), v4978, v4980, l1)
	mBase = m.M
	v4996 = m.ExcPending
	if v4996 != 0 {
		goto L46
	} else {
		goto L1381
	}
L1381:
	;
	v4997 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v4997 == int32(269) {
		v9616 = v4995
		v9618 = v4983
		goto L13
	} else {
		goto L1382
	}
L1382:
	;
	v5000 = *(*int32)(unsafe.Add(mBase, uint32(v4995)))
	v5001 = *(*int32)(unsafe.Add(mBase, uint32(v4995)+4))
	v5002 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	F_check_sql_expr(m, v5000, v5001, v5002, l1)
	mBase = m.M
	v5004 = m.ExcPending
	if v5004 != 0 {
		goto L46
	} else {
		goto L1383
	}
L1383:
	;
	v5006 = F_palloc0(m, int32(28))
	mBase = m.M
	v5007 = m.ExcPending
	if v5007 != 0 {
		goto L46
	} else {
		goto L1384
	}
L1384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5006))) = int32(7)
	v5011 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5012 = *(*int32)(unsafe.Add(mBase, uint32(v5011)+520))
	v5014 = v5012 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5011)+520)) = v5014
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+8)) = v5014
	v5017 = int32(4)
	v5018 = v147 - v5017
	v5020 = v148 - v5017
	v5021 = *(*int32)(unsafe.Add(mBase, uint32(v5020)))
	if v5021 != 0 {
		goto L1386
	} else {
		goto L1387
	}
L1385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+24)) = v4995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5006
	v9708 = v233
	goto L12
L1386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+16)) = v5021
	v5023 = *(*int32)(unsafe.Add(mBase, uint32(v5020)))
	v5024 = *(*int32)(unsafe.Add(mBase, uint32(v5018)))
	F_check_assignable(m, v5023, v5024, l1)
	mBase = m.M
	v5026 = m.ExcPending
	if v5026 != 0 {
		goto L46
	} else {
		goto L1389
	}
L1387:
	;
	goto L1388
L1388:
	;
	v5029 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(8))))
	if v5029 == int32(0) {
		goto L26
	} else {
		goto L1390
	}
L1389:
	;
	goto L1385
L1390:
	;
	v5034 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v5037 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v5038 = *(*int32)(unsafe.Add(mBase, uint32(v5018)))
	v5039 = F_make_scalar_list1(m, v5034, v5029, v5037, v5038, l1)
	mBase = m.M
	v5040 = m.ExcPending
	if v5040 != 0 {
		goto L46
	} else {
		goto L1391
	}
L1391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+16)) = v5039
	goto L1385
L1392:
	;
	v5049 = v5045
	goto L1394
L1393:
	;
	v5046 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v5047 = F_NameListToString(m, v5046)
	mBase = m.M
	v5048 = m.ExcPending
	if v5048 != 0 {
		goto L46
	} else {
		goto L1395
	}
L1394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5049
	v5051 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5052 = int32(0)
	if v5051 < v5052 {
		v5097 = v5052
		goto L1397
	} else {
		goto L1398
	}
L1395:
	;
	v5049 = v5047
	goto L1394
L1396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v5097
	v5100 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v5101 = *(*int32)(unsafe.Add(mBase, uint32(v5100)))
	v5102 = int32(1)
	if base.Ui32(v5101-v5102) <= base.Ui32(v5102) {
		goto L1409
	} else {
		goto L1410
	}
L1397:
	;
	goto L1396
L1398:
	;
	v5057 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5058 = *(*int32)(unsafe.Add(mBase, uint32(v5057)+60))
	if v5058 == int32(0) {
		v5097 = v5052
		goto L1397
	} else {
		goto L1399
	}
L1399:
	;
	v5061 = v5051 + v5058
	v5062 = *(*int32)(unsafe.Add(mBase, uint32(v5057)+188))
	if base.Ui32(v5062) <= base.Ui32(v5061) {
		goto L1401
	} else {
		goto L1402
	}
L1400:
	;
	v5072 = *(*int32)(unsafe.Add(mBase, uint32(v5057)+196))
	if base.B2i32(v5071 == int32(0))|base.B2i32(base.Ui32(v5061) <= base.Ui32(v5071)) != 0 {
		v5097 = v5072
		goto L1397
	} else {
		goto L1404
	}
L1401:
	;
	v5064 = *(*int32)(unsafe.Add(mBase, uint32(v5057)+192))
	v5071 = v5064
	goto L1400
L1402:
	;
	goto L1403
L1403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5057)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5057)+188)) = v5058
	v5069 = F_strchr(m, v5058, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5057)+192)) = v5069
	v5071 = v5069
	goto L1400
L1404:
	;
	v5077 = v5071
	v5080 = v5072
	goto L1405
L1405:
	;
	v5082 = int32(1)
	v5083 = v5080 + v5082
	*(*int32)(unsafe.Add(mBase, uint32(v5057)+196)) = v5083
	v5086 = v5077 + v5082
	*(*int32)(unsafe.Add(mBase, uint32(v5057)+188)) = v5086
	v5089 = F_strchr(m, v5086, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5057)+192)) = v5089
	if v5089 == int32(0) {
		v5097 = v5083
		goto L1397
	} else {
		goto L1407
	}
L1406:
	;
	v5097 = v5083
	goto L1397
L1407:
	;
	if base.Ui32(v5089) < base.Ui32(v5061) {
		v5077 = v5089
		v5080 = v5083
		goto L1405
	} else {
		goto L1408
	}
L1408:
	;
	goto L1406
L1409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+280)) = int32(0)
	v5108 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+284)) = v5108
	v9708 = v233
	goto L12
L1410:
	;
	goto L1411
L1411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+284)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+280)) = v5100
	v5114 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v5116 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5117 = F_plpgsql_yylex(m, v5114, v5116, l1)
	mBase = m.M
	v5118 = m.ExcPending
	if v5118 != 0 {
		goto L46
	} else {
		goto L1412
	}
L1412:
	;
	F_plpgsql_push_back_token(m, v5117, v5114, v5116, l1)
	mBase = m.M
	v5120 = m.ExcPending
	if v5120 != 0 {
		goto L46
	} else {
		goto L1413
	}
L1413:
	;
	if v5117 != int32(44) {
		v9708 = v233
		goto L12
	} else {
		goto L1414
	}
L1414:
	;
	v5123 = *(*int32)(unsafe.Add(mBase, uint32(v28)+272))
	v5124 = *(*int32)(unsafe.Add(mBase, uint32(v28)+280))
	v5125 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5126 = F_read_into_scalar_list(m, v5123, v5124, v5125, v5114, v5116, l1)
	mBase = m.M
	v5127 = m.ExcPending
	if v5127 != 0 {
		goto L46
	} else {
		goto L1415
	}
L1415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+284)) = v5126
	v9708 = v233
	goto L12
L1416:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+280)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v5177
	v5183 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v5185 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5186 = F_plpgsql_yylex(m, v5183, v5185, l1)
	mBase = m.M
	v5187 = m.ExcPending
	if v5187 != 0 {
		goto L46
	} else {
		goto L1429
	}
L1417:
	;
	goto L1416
L1418:
	;
	v5137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5138 = *(*int32)(unsafe.Add(mBase, uint32(v5137)+60))
	if v5138 == int32(0) {
		v5177 = v5132
		goto L1417
	} else {
		goto L1419
	}
L1419:
	;
	v5141 = v5131 + v5138
	v5142 = *(*int32)(unsafe.Add(mBase, uint32(v5137)+188))
	if base.Ui32(v5142) <= base.Ui32(v5141) {
		goto L1421
	} else {
		goto L1422
	}
L1420:
	;
	v5152 = *(*int32)(unsafe.Add(mBase, uint32(v5137)+196))
	if base.B2i32(v5151 == int32(0))|base.B2i32(base.Ui32(v5141) <= base.Ui32(v5151)) != 0 {
		v5177 = v5152
		goto L1417
	} else {
		goto L1424
	}
L1421:
	;
	v5144 = *(*int32)(unsafe.Add(mBase, uint32(v5137)+192))
	v5151 = v5144
	goto L1420
L1422:
	;
	goto L1423
L1423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5137)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5137)+188)) = v5138
	v5149 = F_strchr(m, v5138, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5137)+192)) = v5149
	v5151 = v5149
	goto L1420
L1424:
	;
	v5157 = v5151
	v5160 = v5152
	goto L1425
L1425:
	;
	v5162 = int32(1)
	v5163 = v5160 + v5162
	*(*int32)(unsafe.Add(mBase, uint32(v5137)+196)) = v5163
	v5166 = v5157 + v5162
	*(*int32)(unsafe.Add(mBase, uint32(v5137)+188)) = v5166
	v5169 = F_strchr(m, v5166, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5137)+192)) = v5169
	if v5169 == int32(0) {
		v5177 = v5163
		goto L1417
	} else {
		goto L1427
	}
L1426:
	;
	v5177 = v5163
	goto L1417
L1427:
	;
	if base.Ui32(v5169) < base.Ui32(v5141) {
		v5157 = v5169
		v5160 = v5163
		goto L1425
	} else {
		goto L1428
	}
L1428:
	;
	goto L1426
L1429:
	;
	F_plpgsql_push_back_token(m, v5186, v5183, v5185, l1)
	mBase = m.M
	v5189 = m.ExcPending
	if v5189 != 0 {
		goto L46
	} else {
		goto L1430
	}
L1430:
	;
	if v5186 != int32(44) {
		v9708 = v233
		goto L12
	} else {
		goto L1431
	}
L1431:
	;
	goto L6
L1432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5193))) = int32(9)
	v5199 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(24))))
	v5200 = int32(0)
	if v5199 < v5200 {
		v5245 = v5200
		goto L1434
	} else {
		goto L1435
	}
L1433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5193)+4)) = v5245
	v5249 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5250 = *(*int32)(unsafe.Add(mBase, uint32(v5249)+520))
	v5252 = v5250 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5249)+520)) = v5252
	*(*int32)(unsafe.Add(mBase, uint32(v5193)+8)) = v5252
	v5256 = v148 - int32(112)
	v5257 = *(*int32)(unsafe.Add(mBase, uint32(v5256)))
	*(*int32)(unsafe.Add(mBase, uint32(v5193)+12)) = v5257
	v5261 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v5193)+20)) = v5261
	v5265 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v5193)+24)) = v5265
	v5267 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v5193)+28)) = v5267
	v5270 = v148 - int32(68)
	v5271 = *(*int32)(unsafe.Add(mBase, uint32(v5270)))
	if v5271 == int32(0) {
		goto L1446
	} else {
		goto L1447
	}
L1434:
	;
	goto L1433
L1435:
	;
	v5205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5206 = *(*int32)(unsafe.Add(mBase, uint32(v5205)+60))
	if v5206 == int32(0) {
		v5245 = v5200
		goto L1434
	} else {
		goto L1436
	}
L1436:
	;
	v5209 = v5199 + v5206
	v5210 = *(*int32)(unsafe.Add(mBase, uint32(v5205)+188))
	if base.Ui32(v5210) <= base.Ui32(v5209) {
		goto L1438
	} else {
		goto L1439
	}
L1437:
	;
	v5220 = *(*int32)(unsafe.Add(mBase, uint32(v5205)+196))
	if base.B2i32(v5219 == int32(0))|base.B2i32(base.Ui32(v5209) <= base.Ui32(v5219)) != 0 {
		v5245 = v5220
		goto L1434
	} else {
		goto L1441
	}
L1438:
	;
	v5212 = *(*int32)(unsafe.Add(mBase, uint32(v5205)+192))
	v5219 = v5212
	goto L1437
L1439:
	;
	goto L1440
L1440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5205)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5205)+188)) = v5206
	v5217 = F_strchr(m, v5206, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5205)+192)) = v5217
	v5219 = v5217
	goto L1437
L1441:
	;
	v5225 = v5219
	v5228 = v5220
	goto L1442
L1442:
	;
	v5230 = int32(1)
	v5231 = v5228 + v5230
	*(*int32)(unsafe.Add(mBase, uint32(v5205)+196)) = v5231
	v5234 = v5225 + v5230
	*(*int32)(unsafe.Add(mBase, uint32(v5205)+188)) = v5234
	v5237 = F_strchr(m, v5234, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5205)+192)) = v5237
	if v5237 == int32(0) {
		v5245 = v5231
		goto L1434
	} else {
		goto L1444
	}
L1443:
	;
	v5245 = v5231
	goto L1434
L1444:
	;
	if base.Ui32(v5237) < base.Ui32(v5209) {
		v5225 = v5237
		v5228 = v5231
		goto L1442
	} else {
		goto L1445
	}
L1445:
	;
	goto L1443
L1446:
	;
	v5275 = v148 - int32(72)
	v5276 = *(*int32)(unsafe.Add(mBase, uint32(v5275)))
	if v5276 == int32(0) {
		goto L25
	} else {
		goto L1449
	}
L1447:
	;
	v5279 = v5271
	v5280 = v5270
	goto L1448
L1448:
	;
	v5281 = *(*int32)(unsafe.Add(mBase, uint32(v5279)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5193)+16)) = v5281
	v5283 = *(*int32)(unsafe.Add(mBase, uint32(v5280)))
	v5286 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(20))))
	F_check_assignable(m, v5283, v5286, l1)
	mBase = m.M
	v5288 = m.ExcPending
	if v5288 != 0 {
		goto L46
	} else {
		goto L1450
	}
L1449:
	;
	v5279 = v5276
	v5280 = v5275
	goto L1448
L1450:
	;
	v5289 = *(*int32)(unsafe.Add(mBase, uint32(v5256)))
	v5290 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v5291 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v5289, v5290, v5291, l1)
	mBase = m.M
	v5293 = m.ExcPending
	if v5293 != 0 {
		goto L46
	} else {
		goto L1451
	}
L1451:
	;
	v5296 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v5297 = *(*int32)(unsafe.Add(mBase, uint32(v5296)))
	if v5297 != 0 {
		goto L1453
	} else {
		goto L1454
	}
L1452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5193
	v9708 = v233
	goto L12
L1453:
	;
	v5298 = v5296
	goto L1456
L1454:
	;
	v5301 = v5296
	goto L1455
L1455:
	;
	v5303 = *(*int32)(unsafe.Add(mBase, uint32(v5301)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])) = v5303
	goto L1452
L1456:
	;
	v5299 = *(*int32)(unsafe.Add(mBase, uint32(v5298)+8))
	v5300 = *(*int32)(unsafe.Add(mBase, uint32(v5299)))
	if v5300 != 0 {
		v5298 = v5299
		goto L1456
	} else {
		goto L1458
	}
L1457:
	;
	v5301 = v5299
	goto L1455
L1458:
	;
	goto L1457
L1459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5311))) = int32(10)
	v5316 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5317 = *(*int32)(unsafe.Add(mBase, uint32(v5316)+520))
	v5319 = v5317 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5316)+520)) = v5319
	*(*int32)(unsafe.Add(mBase, uint32(v5311)+8)) = v5319
	v5324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148-int32(32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5311)+12)) = uint8(v5324)
	v5327 = v147 - int32(8)
	v5328 = *(*int32)(unsafe.Add(mBase, uint32(v5327)))
	v5329 = int32(0)
	if v5328 < v5329 {
		v5374 = v5329
		goto L1461
	} else {
		goto L1462
	}
L1460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5311)+4)) = v5374
	v5378 = v148 - int32(16)
	v5379 = *(*int32)(unsafe.Add(mBase, uint32(v5378)))
	*(*int32)(unsafe.Add(mBase, uint32(v5311)+16)) = v5379
	v5381 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v5311)+20)) = v5381
	v5384 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12]))
	v5385 = *(*int32)(unsafe.Add(mBase, uint32(v5378)))
	if v5385 != 0 {
		goto L1474
	} else {
		goto L1475
	}
L1461:
	;
	goto L1460
L1462:
	;
	v5334 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5335 = *(*int32)(unsafe.Add(mBase, uint32(v5334)+60))
	if v5335 == int32(0) {
		v5374 = v5329
		goto L1461
	} else {
		goto L1463
	}
L1463:
	;
	v5338 = v5328 + v5335
	v5339 = *(*int32)(unsafe.Add(mBase, uint32(v5334)+188))
	if base.Ui32(v5339) <= base.Ui32(v5338) {
		goto L1465
	} else {
		goto L1466
	}
L1464:
	;
	v5349 = *(*int32)(unsafe.Add(mBase, uint32(v5334)+196))
	if base.B2i32(v5348 == int32(0))|base.B2i32(base.Ui32(v5338) <= base.Ui32(v5348)) != 0 {
		v5374 = v5349
		goto L1461
	} else {
		goto L1468
	}
L1465:
	;
	v5341 = *(*int32)(unsafe.Add(mBase, uint32(v5334)+192))
	v5348 = v5341
	goto L1464
L1466:
	;
	goto L1467
L1467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5334)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5334)+188)) = v5335
	v5346 = F_strchr(m, v5335, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5334)+192)) = v5346
	v5348 = v5346
	goto L1464
L1468:
	;
	v5354 = v5348
	v5357 = v5349
	goto L1469
L1469:
	;
	v5359 = int32(1)
	v5360 = v5357 + v5359
	*(*int32)(unsafe.Add(mBase, uint32(v5334)+196)) = v5360
	v5363 = v5354 + v5359
	*(*int32)(unsafe.Add(mBase, uint32(v5334)+188)) = v5363
	v5366 = F_strchr(m, v5363, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5334)+192)) = v5366
	if v5366 == int32(0) {
		v5374 = v5360
		goto L1461
	} else {
		goto L1471
	}
L1470:
	;
	v5374 = v5360
	goto L1461
L1471:
	;
	if base.Ui32(v5366) < base.Ui32(v5338) {
		v5354 = v5366
		v5357 = v5360
		goto L1469
	} else {
		goto L1472
	}
L1472:
	;
	goto L1470
L1473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5311
	v9708 = v233
	goto L12
L1474:
	;
	v5386 = *(*int32)(unsafe.Add(mBase, uint32(v5378)))
	if v5384 != 0 {
		goto L1478
	} else {
		goto L1479
	}
L1475:
	;
	goto L1476
L1476:
	;
	if v5384 != 0 {
		goto L1496
	} else {
		goto L1497
	}
L1477:
	;
	if v5399 == int32(0) {
		goto L24
	} else {
		goto L1487
	}
L1478:
	;
	v5387 = v5384
	goto L1481
L1479:
	;
	goto L1480
L1480:
	;
	v5399 = int32(0)
	goto L1477
L1481:
	;
	v5389 = *(*int32)(unsafe.Add(mBase, uint32(v5387)))
	if v5389 != 0 {
		goto L1483
	} else {
		goto L1484
	}
L1482:
	;
	goto L1480
L1483:
	;
	v5393 = *(*int32)(unsafe.Add(mBase, uint32(v5387)+8))
	if v5393 != 0 {
		v5387 = v5393
		goto L1481
	} else {
		goto L1486
	}
L1484:
	;
	v5392 = F_strcmp(m, v5387+int32(12), v5386)
	mBase = m.M
	if v5392 != 0 {
		goto L1483
	} else {
		goto L1485
	}
L1485:
	;
	v5399 = v5387
	goto L1477
L1486:
	;
	goto L1482
L1487:
	;
	v5402 = *(*int32)(unsafe.Add(mBase, uint32(v5399)+4))
	if v5402 == int32(1) {
		goto L1473
	} else {
		goto L1488
	}
L1488:
	;
	v5405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5311)+12)))
	if v5405 != 0 {
		goto L1473
	} else {
		goto L1489
	}
L1489:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5409 = m.ExcPending
	if v5409 != 0 {
		goto L46
	} else {
		goto L1490
	}
L1490:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v5412 = m.ExcPending
	if v5412 != 0 {
		goto L46
	} else {
		goto L1491
	}
L1491:
	;
	v5413 = *(*int32)(unsafe.Add(mBase, uint32(v5378)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+208)) = v5413
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_59), v28+int32(208))
	mBase = m.M
	v5419 = m.ExcPending
	if v5419 != 0 {
		goto L46
	} else {
		goto L1492
	}
L1492:
	;
	v5422 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v5423 = F_plpgsql_scanner_errposition(m, v5422, l1)
	mBase = m.M
	v5424 = m.ExcPending
	if v5424 != 0 {
		goto L46
	} else {
		goto L1493
	}
L1493:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1752), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v5429 = m.ExcPending
	if v5429 != 0 {
		goto L46
	} else {
		goto L1494
	}
L1494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1495:
	;
	if v5439 == int32(0) {
		goto L23
	} else {
		goto L1505
	}
L1496:
	;
	v5430 = v5384
	goto L1499
L1497:
	;
	goto L1498
L1498:
	;
	v5439 = int32(0)
	goto L1495
L1499:
	;
	v5431 = *(*int32)(unsafe.Add(mBase, uint32(v5430)))
	if v5431 != 0 {
		goto L1501
	} else {
		goto L1502
	}
L1500:
	;
	goto L1498
L1501:
	;
	v5435 = *(*int32)(unsafe.Add(mBase, uint32(v5430)+8))
	if v5435 != 0 {
		v5430 = v5435
		goto L1499
	} else {
		goto L1504
	}
L1502:
	;
	v5432 = *(*int32)(unsafe.Add(mBase, uint32(v5430)+4))
	if v5432 != int32(1) {
		goto L1501
	} else {
		goto L1503
	}
L1503:
	;
	v5439 = v5430
	goto L1495
L1504:
	;
	goto L1500
L1505:
	;
	goto L1473
L1506:
	;
	v5873 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v5875 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, v5452, v5873, v5875, l1)
	mBase = m.M
	v5877 = m.ExcPending
	if v5877 != 0 {
		goto L46
	} else {
		goto L1613
	}
L1507:
	;
	v5674 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5675 = m.G0
	v5677 = v5675 - int32(16)
	m.G0 = v5677
	v5680 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5680)+63)))
	if v5681 != 0 {
		goto L1577
	} else {
		goto L1578
	}
L1508:
	;
	v5646 = int32(_a_F_plpgsql_yyparse_60)
	v5649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5461))))
	v5652 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[41])))
	if base.B2i32(v5649 == int32(0))|base.B2i32(v5649 != v5652) != 0 {
		v5670 = v5649
		v5671 = v5652
		goto L1569
	} else {
		goto L1570
	}
L1509:
	;
	v5492 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5494 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v5496 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5498 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5498)+63)))
	if v5499 != 0 {
		goto L1528
	} else {
		goto L1529
	}
L1510:
	;
	if v5452 != int32(277) {
		goto L1512
	} else {
		goto L1513
	}
L1511:
	;
	switch v5452 - int32(341) {
	case 0:
		goto L1509
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		goto L1506
	case 17:
		goto L1507
	default:
		goto L1510
	}
L1512:
	;
	if v5452 != 0 {
		goto L1506
	} else {
		goto L1515
	}
L1513:
	;
	goto L1514
L1514:
	;
	v5458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v5458&int32(1) != 0 {
		goto L1506
	} else {
		goto L1516
	}
L1515:
	;
	goto L4
L1516:
	;
	v5461 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v5461 == int32(0) {
		goto L1506
	} else {
		goto L1517
	}
L1517:
	;
	v5464 = int32(_a_F_plpgsql_yyparse_61)
	v5467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5461))))
	v5470 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[42])))
	if base.B2i32(v5467 == int32(0))|base.B2i32(v5467 != v5470) != 0 {
		v5488 = v5467
		v5489 = v5470
		goto L1519
	} else {
		goto L1520
	}
L1518:
	;
	if v5488-v5489 != 0 {
		goto L1508
	} else {
		goto L1525
	}
L1519:
	;
	goto L1518
L1520:
	;
	v5473 = v5461
	v5474 = v5464
	goto L1521
L1521:
	;
	v5477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5474)+1)))
	v5478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5473)+1)))
	if v5478 == int32(0) {
		v5488 = v5478
		v5489 = v5477
		goto L1519
	} else {
		goto L1523
	}
L1522:
	;
	v5488 = v5478
	v5489 = v5477
	goto L1519
L1523:
	;
	v5481 = int32(1)
	if v5478 == v5477 {
		v5473 = v5473 + v5481
		v5474 = v5474 + v5481
		goto L1521
	} else {
		goto L1524
	}
L1524:
	;
	goto L1522
L1525:
	;
	goto L1509
L1526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5501
	v9708 = v233
	goto L12
L1527:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5627 = m.ExcPending
	if v5627 != 0 {
		goto L46
	} else {
		goto L1563
	}
L1528:
	;
	v5501 = F_palloc0(m, int32(20))
	mBase = m.M
	v5502 = m.ExcPending
	if v5502 != 0 {
		goto L46
	} else {
		goto L1531
	}
L1529:
	;
	goto L1530
L1530:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5609 = m.ExcPending
	if v5609 != 0 {
		goto L46
	} else {
		goto L1558
	}
L1531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5501))) = int32(12)
	v5505 = int32(0)
	if v5492 < v5505 {
		v5550 = v5505
		goto L1533
	} else {
		goto L1534
	}
L1532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5501)+4)) = v5550
	v5554 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5555 = *(*int32)(unsafe.Add(mBase, uint32(v5554)+520))
	v5557 = v5555 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5554)+520)) = v5557
	*(*int64)(unsafe.Add(mBase, uint32(v5501)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5501)+8)) = v5557
	v5562 = *(*int32)(unsafe.Add(mBase, uint32(v5554)+472))
	v5563 = F_plpgsql_yylex(m, v5494, v5496, l1)
	mBase = m.M
	v5564 = m.ExcPending
	if v5564 != 0 {
		goto L46
	} else {
		goto L1545
	}
L1533:
	;
	goto L1532
L1534:
	;
	v5510 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5511 = *(*int32)(unsafe.Add(mBase, uint32(v5510)+60))
	if v5511 == int32(0) {
		v5550 = v5505
		goto L1533
	} else {
		goto L1535
	}
L1535:
	;
	v5514 = v5492 + v5511
	v5515 = *(*int32)(unsafe.Add(mBase, uint32(v5510)+188))
	if base.Ui32(v5515) <= base.Ui32(v5514) {
		goto L1537
	} else {
		goto L1538
	}
L1536:
	;
	v5525 = *(*int32)(unsafe.Add(mBase, uint32(v5510)+196))
	if base.B2i32(v5524 == int32(0))|base.B2i32(base.Ui32(v5514) <= base.Ui32(v5524)) != 0 {
		v5550 = v5525
		goto L1533
	} else {
		goto L1540
	}
L1537:
	;
	v5517 = *(*int32)(unsafe.Add(mBase, uint32(v5510)+192))
	v5524 = v5517
	goto L1536
L1538:
	;
	goto L1539
L1539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5510)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5510)+188)) = v5511
	v5522 = F_strchr(m, v5511, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5510)+192)) = v5522
	v5524 = v5522
	goto L1536
L1540:
	;
	v5530 = v5524
	v5533 = v5525
	goto L1541
L1541:
	;
	v5535 = int32(1)
	v5536 = v5533 + v5535
	*(*int32)(unsafe.Add(mBase, uint32(v5510)+196)) = v5536
	v5539 = v5530 + v5535
	*(*int32)(unsafe.Add(mBase, uint32(v5510)+188)) = v5539
	v5542 = F_strchr(m, v5539, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5510)+192)) = v5542
	if v5542 == int32(0) {
		v5550 = v5536
		goto L1533
	} else {
		goto L1543
	}
L1542:
	;
	v5550 = v5536
	goto L1533
L1543:
	;
	if base.Ui32(v5542) < base.Ui32(v5514) {
		v5530 = v5542
		v5533 = v5536
		goto L1541
	} else {
		goto L1544
	}
L1544:
	;
	goto L1542
L1545:
	;
	if int32(0) <= v5562 {
		goto L1546
	} else {
		goto L1547
	}
L1546:
	;
	if v5563 != int32(59) {
		goto L1527
	} else {
		goto L1549
	}
L1547:
	;
	goto L1548
L1548:
	;
	if v5563 != int32(277) {
		goto L1550
	} else {
		goto L1551
	}
L1549:
	;
	v5570 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5571 = *(*int32)(unsafe.Add(mBase, uint32(v5570)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v5501)+16)) = v5571
	goto L1526
L1550:
	;
	F_plpgsql_push_back_token(m, v5563, v5494, v5496, l1)
	mBase = m.M
	v5593 = m.ExcPending
	if v5593 != 0 {
		goto L46
	} else {
		goto L1556
	}
L1551:
	;
	v5575 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v5576 = m.ExcPending
	if v5576 != 0 {
		goto L46
	} else {
		goto L1552
	}
L1552:
	;
	if v5575 != int32(59) {
		goto L1550
	} else {
		goto L1553
	}
L1553:
	;
	v5579 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v5580 = *(*int32)(unsafe.Add(mBase, uint32(v5579)))
	if base.B2i32(base.Ui32(int32(4)) < base.Ui32(v5580))|base.B2i32(v5580 == int32(3)) != 0 {
		goto L1550
	} else {
		goto L1554
	}
L1554:
	;
	v5586 = *(*int32)(unsafe.Add(mBase, uint32(v5579)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5501)+16)) = v5586
	v5588 = F_plpgsql_yylex(m, v5494, v5496, l1)
	mBase = m.M
	v5589 = m.ExcPending
	if v5589 != 0 {
		goto L46
	} else {
		goto L1555
	}
L1555:
	;
	goto L1526
L1556:
	;
	v5595 = int32(0)
	v5599 = int32(1)
	v5603 = F_read_sql_construct(m, int32(59), v5595, v5595, int32(_a_F_plpgsql_yyparse_16), int32(2), v5599, v5599, v5595, v5595, v5494, v5496, l1)
	mBase = m.M
	v5604 = m.ExcPending
	if v5604 != 0 {
		goto L46
	} else {
		goto L1557
	}
L1557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5501)+12)) = v5603
	goto L1526
L1558:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5612 = m.ExcPending
	if v5612 != 0 {
		goto L46
	} else {
		goto L1559
	}
L1559:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_62), int32(0))
	mBase = m.M
	v5616 = m.ExcPending
	if v5616 != 0 {
		goto L46
	} else {
		goto L1560
	}
L1560:
	;
	v5617 = F_plpgsql_scanner_errposition(m, v5492, l1)
	mBase = m.M
	v5618 = m.ExcPending
	if v5618 != 0 {
		goto L46
	} else {
		goto L1561
	}
L1561:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3445), int32(_a_F_plpgsql_yyparse_63))
	mBase = m.M
	v5623 = m.ExcPending
	if v5623 != 0 {
		goto L46
	} else {
		goto L1562
	}
L1562:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1563:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5630 = m.ExcPending
	if v5630 != 0 {
		goto L46
	} else {
		goto L1564
	}
L1564:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_64), int32(0))
	mBase = m.M
	v5634 = m.ExcPending
	if v5634 != 0 {
		goto L46
	} else {
		goto L1565
	}
L1565:
	;
	v5635 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v5636 = F_plpgsql_scanner_errposition(m, v5635, l1)
	mBase = m.M
	v5637 = m.ExcPending
	if v5637 != 0 {
		goto L46
	} else {
		goto L1566
	}
L1566:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3460), int32(_a_F_plpgsql_yyparse_63))
	mBase = m.M
	v5642 = m.ExcPending
	if v5642 != 0 {
		goto L46
	} else {
		goto L1567
	}
L1567:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1568:
	;
	if v5670-v5671 != 0 {
		goto L1506
	} else {
		goto L1575
	}
L1569:
	;
	goto L1568
L1570:
	;
	v5655 = v5461
	v5656 = v5646
	goto L1571
L1571:
	;
	v5659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5656)+1)))
	v5660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5655)+1)))
	if v5660 == int32(0) {
		v5670 = v5660
		v5671 = v5659
		goto L1569
	} else {
		goto L1573
	}
L1572:
	;
	v5670 = v5660
	v5671 = v5659
	goto L1569
L1573:
	;
	v5663 = int32(1)
	if v5660 == v5659 {
		v5655 = v5655 + v5663
		v5656 = v5656 + v5663
		goto L1571
	} else {
		goto L1574
	}
L1574:
	;
	goto L1572
L1575:
	;
	goto L1507
L1576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5687
	v9708 = v233
	goto L12
L1577:
	;
	v5683 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v5685 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5687 = F_palloc0(m, int32(24))
	mBase = m.M
	v5688 = m.ExcPending
	if v5688 != 0 {
		goto L46
	} else {
		goto L1580
	}
L1578:
	;
	goto L1579
L1579:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5855 = m.ExcPending
	if v5855 != 0 {
		goto L46
	} else {
		goto L1608
	}
L1580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5687))) = int32(13)
	v5691 = int32(0)
	if v5674 < v5691 {
		v5736 = v5691
		goto L1582
	} else {
		goto L1583
	}
L1581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5687)+4)) = v5736
	v5740 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5741 = *(*int32)(unsafe.Add(mBase, uint32(v5740)+520))
	v5743 = v5741 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5740)+520)) = v5743
	*(*int32)(unsafe.Add(mBase, uint32(v5687)+8)) = v5743
	v5746 = F_plpgsql_yylex(m, v5683, v5685, l1)
	mBase = m.M
	v5747 = m.ExcPending
	if v5747 != 0 {
		goto L46
	} else {
		goto L1595
	}
L1582:
	;
	goto L1581
L1583:
	;
	v5696 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5697 = *(*int32)(unsafe.Add(mBase, uint32(v5696)+60))
	if v5697 == int32(0) {
		v5736 = v5691
		goto L1582
	} else {
		goto L1584
	}
L1584:
	;
	v5700 = v5674 + v5697
	v5701 = *(*int32)(unsafe.Add(mBase, uint32(v5696)+188))
	if base.Ui32(v5701) <= base.Ui32(v5700) {
		goto L1586
	} else {
		goto L1587
	}
L1585:
	;
	v5711 = *(*int32)(unsafe.Add(mBase, uint32(v5696)+196))
	if base.B2i32(v5710 == int32(0))|base.B2i32(base.Ui32(v5700) <= base.Ui32(v5710)) != 0 {
		v5736 = v5711
		goto L1582
	} else {
		goto L1589
	}
L1586:
	;
	v5703 = *(*int32)(unsafe.Add(mBase, uint32(v5696)+192))
	v5710 = v5703
	goto L1585
L1587:
	;
	goto L1588
L1588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5696)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5696)+188)) = v5697
	v5708 = F_strchr(m, v5697, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5696)+192)) = v5708
	v5710 = v5708
	goto L1585
L1589:
	;
	v5716 = v5710
	v5719 = v5711
	goto L1590
L1590:
	;
	v5721 = int32(1)
	v5722 = v5719 + v5721
	*(*int32)(unsafe.Add(mBase, uint32(v5696)+196)) = v5722
	v5725 = v5716 + v5721
	*(*int32)(unsafe.Add(mBase, uint32(v5696)+188)) = v5725
	v5728 = F_strchr(m, v5725, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5696)+192)) = v5728
	if v5728 == int32(0) {
		v5736 = v5722
		goto L1582
	} else {
		goto L1592
	}
L1591:
	;
	v5736 = v5722
	goto L1582
L1592:
	;
	if base.Ui32(v5728) < base.Ui32(v5700) {
		v5716 = v5728
		v5719 = v5722
		goto L1590
	} else {
		goto L1593
	}
L1593:
	;
	goto L1591
L1594:
	;
	m.G0 = v5677 + int32(16)
	goto L1576
L1595:
	;
	if v5746 != int32(317) {
		goto L1596
	} else {
		goto L1597
	}
L1596:
	;
	F_plpgsql_push_back_token(m, v5746, v5683, v5685, l1)
	mBase = m.M
	v5751 = m.ExcPending
	if v5751 != 0 {
		goto L46
	} else {
		goto L1599
	}
L1597:
	;
	goto L1598
L1598:
	;
	v5766 = int32(0)
	v5769 = int32(1)
	v5774 = F_read_sql_construct(m, int32(59), int32(381), v5766, int32(_a_F_plpgsql_yyparse_65), int32(2), v5769, v5769, v5766, v5677+int32(12), v5683, v5685, l1)
	mBase = m.M
	v5775 = m.ExcPending
	if v5775 != 0 {
		goto L46
	} else {
		goto L1601
	}
L1599:
	;
	v5753 = int32(0)
	v5761 = F_read_sql_construct(m, int32(59), v5753, v5753, int32(_a_F_plpgsql_yyparse_16), v5753, v5753, int32(1), v5753, v5753, v5683, v5685, l1)
	mBase = m.M
	v5762 = m.ExcPending
	if v5762 != 0 {
		goto L46
	} else {
		goto L1600
	}
L1600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5687)+12)) = v5761
	goto L1594
L1601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5687)+16)) = v5774
	v5777 = *(*int32)(unsafe.Add(mBase, uint32(v5677)+12))
	if v5777 != int32(381) {
		goto L1594
	} else {
		goto L1602
	}
L1602:
	;
	goto L1603
L1603:
	;
	v5807 = int32(0)
	v5810 = int32(1)
	v5815 = F_read_sql_construct(m, int32(44), int32(59), v5807, int32(_a_F_plpgsql_yyparse_66), int32(2), v5810, v5810, v5807, v5677+int32(12), v5683, v5685, l1)
	mBase = m.M
	v5816 = m.ExcPending
	if v5816 != 0 {
		goto L46
	} else {
		goto L1605
	}
L1604:
	;
	goto L1594
L1605:
	;
	v5817 = *(*int32)(unsafe.Add(mBase, uint32(v5687)+20))
	v5818 = F_lappend(m, v5817, v5815)
	mBase = m.M
	v5819 = m.ExcPending
	if v5819 != 0 {
		goto L46
	} else {
		goto L1606
	}
L1606:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5687)+20)) = v5818
	v5821 = *(*int32)(unsafe.Add(mBase, uint32(v5677)+12))
	if v5821 == int32(44) {
		goto L1603
	} else {
		goto L1607
	}
L1607:
	;
	goto L1604
L1608:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5858 = m.ExcPending
	if v5858 != 0 {
		goto L46
	} else {
		goto L1609
	}
L1609:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_67), int32(0))
	mBase = m.M
	v5862 = m.ExcPending
	if v5862 != 0 {
		goto L46
	} else {
		goto L1610
	}
L1610:
	;
	v5863 = F_plpgsql_scanner_errposition(m, v5674, l1)
	mBase = m.M
	v5864 = m.ExcPending
	if v5864 != 0 {
		goto L46
	} else {
		goto L1611
	}
L1611:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3509), int32(_a_F_plpgsql_yyparse_68))
	mBase = m.M
	v5869 = m.ExcPending
	if v5869 != 0 {
		goto L46
	} else {
		goto L1612
	}
L1612:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1613:
	;
	v5878 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5880 = F_palloc0(m, int32(20))
	mBase = m.M
	v5881 = m.ExcPending
	if v5881 != 0 {
		goto L46
	} else {
		goto L1614
	}
L1614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5880))) = int32(11)
	v5884 = int32(0)
	if v5878 < v5884 {
		v5929 = v5884
		goto L1616
	} else {
		goto L1617
	}
L1615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5880)+4)) = v5929
	v5933 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5934 = *(*int32)(unsafe.Add(mBase, uint32(v5933)+520))
	v5935 = int32(1)
	v5936 = v5934 + v5935
	*(*int32)(unsafe.Add(mBase, uint32(v5933)+520)) = v5936
	*(*int64)(unsafe.Add(mBase, uint32(v5880)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5880)+8)) = v5936
	v5941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5933)+63)))
	if v5941 == v5935 {
		goto L1629
	} else {
		goto L1630
	}
L1616:
	;
	goto L1615
L1617:
	;
	v5889 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5890 = *(*int32)(unsafe.Add(mBase, uint32(v5889)+60))
	if v5890 == int32(0) {
		v5929 = v5884
		goto L1616
	} else {
		goto L1618
	}
L1618:
	;
	v5893 = v5878 + v5890
	v5894 = *(*int32)(unsafe.Add(mBase, uint32(v5889)+188))
	if base.Ui32(v5894) <= base.Ui32(v5893) {
		goto L1620
	} else {
		goto L1621
	}
L1619:
	;
	v5904 = *(*int32)(unsafe.Add(mBase, uint32(v5889)+196))
	if base.B2i32(v5903 == int32(0))|base.B2i32(base.Ui32(v5893) <= base.Ui32(v5903)) != 0 {
		v5929 = v5904
		goto L1616
	} else {
		goto L1623
	}
L1620:
	;
	v5896 = *(*int32)(unsafe.Add(mBase, uint32(v5889)+192))
	v5903 = v5896
	goto L1619
L1621:
	;
	goto L1622
L1622:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5889)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5889)+188)) = v5890
	v5901 = F_strchr(m, v5890, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5889)+192)) = v5901
	v5903 = v5901
	goto L1619
L1623:
	;
	v5909 = v5903
	v5912 = v5904
	goto L1624
L1624:
	;
	v5914 = int32(1)
	v5915 = v5912 + v5914
	*(*int32)(unsafe.Add(mBase, uint32(v5889)+196)) = v5915
	v5918 = v5909 + v5914
	*(*int32)(unsafe.Add(mBase, uint32(v5889)+188)) = v5918
	v5921 = F_strchr(m, v5918, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5889)+192)) = v5921
	if v5921 == int32(0) {
		v5929 = v5915
		goto L1616
	} else {
		goto L1626
	}
L1625:
	;
	v5929 = v5915
	goto L1616
L1626:
	;
	if base.Ui32(v5921) < base.Ui32(v5893) {
		v5909 = v5921
		v5912 = v5915
		goto L1624
	} else {
		goto L1627
	}
L1627:
	;
	goto L1625
L1628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5880
	v9708 = v233
	goto L12
L1629:
	;
	v5944 = F_plpgsql_yylex(m, v5873, v5875, l1)
	mBase = m.M
	v5945 = m.ExcPending
	if v5945 != 0 {
		goto L46
	} else {
		goto L1632
	}
L1630:
	;
	goto L1631
L1631:
	;
	v5971 = *(*int32)(unsafe.Add(mBase, uint32(v5933)+52))
	if v5971 == int32(2278) {
		goto L1642
	} else {
		goto L1643
	}
L1632:
	;
	if v5944 == int32(59) {
		goto L1628
	} else {
		goto L1633
	}
L1633:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5951 = m.ExcPending
	if v5951 != 0 {
		goto L46
	} else {
		goto L1634
	}
L1634:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5954 = m.ExcPending
	if v5954 != 0 {
		goto L46
	} else {
		goto L1635
	}
L1635:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_69), int32(0))
	mBase = m.M
	v5958 = m.ExcPending
	if v5958 != 0 {
		goto L46
	} else {
		goto L1636
	}
L1636:
	;
	F_errhint(m, int32(_a_F_plpgsql_yyparse_70), int32(0))
	mBase = m.M
	v5962 = m.ExcPending
	if v5962 != 0 {
		goto L46
	} else {
		goto L1637
	}
L1637:
	;
	v5963 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v5964 = F_plpgsql_scanner_errposition(m, v5963, l1)
	mBase = m.M
	v5965 = m.ExcPending
	if v5965 != 0 {
		goto L46
	} else {
		goto L1638
	}
L1638:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3373), int32(_a_F_plpgsql_yyparse_71))
	mBase = m.M
	v5970 = m.ExcPending
	if v5970 != 0 {
		goto L46
	} else {
		goto L1639
	}
L1639:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1640:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v6064 = m.ExcPending
	if v6064 != 0 {
		goto L46
	} else {
		goto L1670
	}
L1641:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6048 = m.ExcPending
	if v6048 != 0 {
		goto L46
	} else {
		goto L1666
	}
L1642:
	;
	v5974 = F_plpgsql_yylex(m, v5873, v5875, l1)
	mBase = m.M
	v5975 = m.ExcPending
	if v5975 != 0 {
		goto L46
	} else {
		goto L1645
	}
L1643:
	;
	goto L1644
L1644:
	;
	v6002 = *(*int32)(unsafe.Add(mBase, uint32(v5933)+472))
	v6003 = F_plpgsql_yylex(m, v5873, v5875, l1)
	mBase = m.M
	v6004 = m.ExcPending
	if v6004 != 0 {
		goto L46
	} else {
		goto L1653
	}
L1645:
	;
	if v5974 == int32(59) {
		goto L1628
	} else {
		goto L1646
	}
L1646:
	;
	v5979 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v5980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5979)+65)))
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5984 = m.ExcPending
	if v5984 != 0 {
		goto L46
	} else {
		goto L1647
	}
L1647:
	;
	if v5980 == int32(112) {
		goto L1641
	} else {
		goto L1648
	}
L1648:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5989 = m.ExcPending
	if v5989 != 0 {
		goto L46
	} else {
		goto L1649
	}
L1649:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_72), int32(0))
	mBase = m.M
	v5993 = m.ExcPending
	if v5993 != 0 {
		goto L46
	} else {
		goto L1650
	}
L1650:
	;
	v5994 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v5995 = F_plpgsql_scanner_errposition(m, v5994, l1)
	mBase = m.M
	v5996 = m.ExcPending
	if v5996 != 0 {
		goto L46
	} else {
		goto L1651
	}
L1651:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3388), int32(_a_F_plpgsql_yyparse_71))
	mBase = m.M
	v6001 = m.ExcPending
	if v6001 != 0 {
		goto L46
	} else {
		goto L1652
	}
L1652:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1653:
	;
	if int32(0) <= v6002 {
		goto L1654
	} else {
		goto L1655
	}
L1654:
	;
	if v6003 != int32(59) {
		goto L1640
	} else {
		goto L1657
	}
L1655:
	;
	goto L1656
L1656:
	;
	if v6003 != int32(277) {
		goto L1658
	} else {
		goto L1659
	}
L1657:
	;
	v6010 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v6011 = *(*int32)(unsafe.Add(mBase, uint32(v6010)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v5880)+16)) = v6011
	goto L1628
L1658:
	;
	F_plpgsql_push_back_token(m, v6003, v5873, v5875, l1)
	mBase = m.M
	v6033 = m.ExcPending
	if v6033 != 0 {
		goto L46
	} else {
		goto L1664
	}
L1659:
	;
	v6015 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v6016 = m.ExcPending
	if v6016 != 0 {
		goto L46
	} else {
		goto L1660
	}
L1660:
	;
	if v6015 != int32(59) {
		goto L1658
	} else {
		goto L1661
	}
L1661:
	;
	v6019 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6020 = *(*int32)(unsafe.Add(mBase, uint32(v6019)))
	if base.B2i32(base.Ui32(int32(4)) < base.Ui32(v6020))|base.B2i32(v6020 == int32(3)) != 0 {
		goto L1658
	} else {
		goto L1662
	}
L1662:
	;
	v6026 = *(*int32)(unsafe.Add(mBase, uint32(v6019)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5880)+16)) = v6026
	v6028 = F_plpgsql_yylex(m, v5873, v5875, l1)
	mBase = m.M
	v6029 = m.ExcPending
	if v6029 != 0 {
		goto L46
	} else {
		goto L1663
	}
L1663:
	;
	goto L1628
L1664:
	;
	v6035 = int32(0)
	v6039 = int32(1)
	v6043 = F_read_sql_construct(m, int32(59), v6035, v6035, int32(_a_F_plpgsql_yyparse_16), int32(2), v6039, v6039, v6035, v6035, v5873, v5875, l1)
	mBase = m.M
	v6044 = m.ExcPending
	if v6044 != 0 {
		goto L46
	} else {
		goto L1665
	}
L1665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5880)+12)) = v6043
	goto L1628
L1666:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_73), int32(0))
	mBase = m.M
	v6052 = m.ExcPending
	if v6052 != 0 {
		goto L46
	} else {
		goto L1667
	}
L1667:
	;
	v6053 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v6054 = F_plpgsql_scanner_errposition(m, v6053, l1)
	mBase = m.M
	v6055 = m.ExcPending
	if v6055 != 0 {
		goto L46
	} else {
		goto L1668
	}
L1668:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3383), int32(_a_F_plpgsql_yyparse_71))
	mBase = m.M
	v6060 = m.ExcPending
	if v6060 != 0 {
		goto L46
	} else {
		goto L1669
	}
L1669:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1670:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v6067 = m.ExcPending
	if v6067 != 0 {
		goto L46
	} else {
		goto L1671
	}
L1671:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_74), int32(0))
	mBase = m.M
	v6071 = m.ExcPending
	if v6071 != 0 {
		goto L46
	} else {
		goto L1672
	}
L1672:
	;
	v6072 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v6073 = F_plpgsql_scanner_errposition(m, v6072, l1)
	mBase = m.M
	v6074 = m.ExcPending
	if v6074 != 0 {
		goto L46
	} else {
		goto L1673
	}
L1673:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3397), int32(_a_F_plpgsql_yyparse_71))
	mBase = m.M
	v6079 = m.ExcPending
	if v6079 != 0 {
		goto L46
	} else {
		goto L1674
	}
L1674:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085))) = int32(14)
	v6089 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v6090 = int32(0)
	if v6089 < v6090 {
		v6135 = v6090
		goto L1677
	} else {
		goto L1678
	}
L1676:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+4)) = v6135
	v6139 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v6140 = *(*int32)(unsafe.Add(mBase, uint32(v6139)+520))
	v6142 = v6140 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6139)+520)) = v6142
	v6144 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6085)+16)) = v6144
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+12)) = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+8)) = v6142
	*(*int64)(unsafe.Add(mBase, uint32(v6085)+24)) = v6144
	v6152 = v6085 + int32(16)
	v6156 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v6157 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v6156, l1)
	mBase = m.M
	v6158 = m.ExcPending
	if v6158 != 0 {
		goto L46
	} else {
		goto L1713
	}
L1677:
	;
	goto L1676
L1678:
	;
	v6095 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6096 = *(*int32)(unsafe.Add(mBase, uint32(v6095)+60))
	if v6096 == int32(0) {
		v6135 = v6090
		goto L1677
	} else {
		goto L1679
	}
L1679:
	;
	v6099 = v6089 + v6096
	v6100 = *(*int32)(unsafe.Add(mBase, uint32(v6095)+188))
	if base.Ui32(v6100) <= base.Ui32(v6099) {
		goto L1681
	} else {
		goto L1682
	}
L1680:
	;
	v6110 = *(*int32)(unsafe.Add(mBase, uint32(v6095)+196))
	if base.B2i32(v6109 == int32(0))|base.B2i32(base.Ui32(v6099) <= base.Ui32(v6109)) != 0 {
		v6135 = v6110
		goto L1677
	} else {
		goto L1684
	}
L1681:
	;
	v6102 = *(*int32)(unsafe.Add(mBase, uint32(v6095)+192))
	v6109 = v6102
	goto L1680
L1682:
	;
	goto L1683
L1683:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6095)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6095)+188)) = v6096
	v6107 = F_strchr(m, v6096, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6095)+192)) = v6107
	v6109 = v6107
	goto L1680
L1684:
	;
	v6115 = v6109
	v6118 = v6110
	goto L1685
L1685:
	;
	v6120 = int32(1)
	v6121 = v6118 + v6120
	*(*int32)(unsafe.Add(mBase, uint32(v6095)+196)) = v6121
	v6124 = v6115 + v6120
	*(*int32)(unsafe.Add(mBase, uint32(v6095)+188)) = v6124
	v6127 = F_strchr(m, v6124, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6095)+192)) = v6127
	if v6127 == int32(0) {
		v6135 = v6121
		goto L1677
	} else {
		goto L1687
	}
L1686:
	;
	v6135 = v6121
	goto L1677
L1687:
	;
	if base.Ui32(v6127) < base.Ui32(v6099) {
		v6115 = v6127
		v6118 = v6121
		goto L1685
	} else {
		goto L1688
	}
L1688:
	;
	goto L1686
L1689:
	;
	v7095 = *(*int32)(unsafe.Add(mBase, uint32(v6085)+20))
	if v7095 != 0 {
		goto L1956
	} else {
		goto L1957
	}
L1690:
	;
	v6695 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v6697 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v6699 = m.G0
	v6701 = v6699 - int32(16)
	m.G0 = v6701
	v6707 = int32(0)
	goto L1845
L1691:
	;
	if v6646 != int32(381) {
		goto L1689
	} else {
		goto L1840
	}
L1692:
	;
	goto L1835
L1693:
	;
	v6581 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v6582 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v6581, l1)
	mBase = m.M
	v6583 = m.ExcPending
	if v6583 != 0 {
		goto L46
	} else {
		goto L1832
	}
L1694:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6152))) = v6571
	v6574 = F_plpgsql_recognize_err_condition(m, v6571, int32(0))
	mBase = m.M
	v6575 = m.ExcPending
	if v6575 != 0 {
		goto L46
	} else {
		goto L1831
	}
L1695:
	;
	v6568 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6571 = v6568
	goto L1694
L1696:
	;
	v6550 = int32(0)
	goto L1821
L1697:
	;
	v6450 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6451 = m.ExcPending
	if v6451 != 0 {
		goto L46
	} else {
		goto L1797
	}
L1698:
	;
	v6410 = int32(277)
	v6411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v6411&int32(1) != 0 {
		v6545 = v6410
		goto L1696
	} else {
		goto L1787
	}
L1699:
	;
	switch v6391 - int32(261) {
	case 0:
		goto L1779
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15:
		v6545 = v6391
		goto L1696
	case 14:
		goto L1695
	case 16:
		goto L1698
	default:
		goto L1780
	}
L1700:
	;
	v6388 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6389 = m.ExcPending
	if v6389 != 0 {
		goto L46
	} else {
		goto L1778
	}
L1701:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+12)) = int32(14)
	goto L1700
L1702:
	;
	v6353 = int32(_a_F_plpgsql_yyparse_75)
	v6356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6320))))
	v6359 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[43])))
	if base.B2i32(v6356 == int32(0))|base.B2i32(v6356 != v6359) != 0 {
		v6377 = v6356
		v6378 = v6359
		goto L1771
	} else {
		goto L1772
	}
L1703:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+12)) = int32(15)
	goto L1700
L1704:
	;
	v6320 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6320 == int32(0) {
		goto L1698
	} else {
		goto L1761
	}
L1705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+12)) = int32(17)
	goto L1700
L1706:
	;
	v6290 = int32(_a_F_plpgsql_yyparse_76)
	v6293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6251))))
	v6296 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[44])))
	if base.B2i32(v6293 == int32(0))|base.B2i32(v6293 != v6296) != 0 {
		v6314 = v6293
		v6315 = v6296
		goto L1754
	} else {
		goto L1755
	}
L1707:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+12)) = int32(18)
	v6288 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6289 = m.ExcPending
	if v6289 != 0 {
		goto L46
	} else {
		goto L1752
	}
L1708:
	;
	v6251 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6251 == int32(0) {
		goto L1698
	} else {
		goto L1743
	}
L1709:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+12)) = int32(19)
	v6249 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6250 = m.ExcPending
	if v6250 != 0 {
		goto L46
	} else {
		goto L1742
	}
L1710:
	;
	v6215 = int32(_a_F_plpgsql_yyparse_77)
	v6218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6176))))
	v6221 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[45])))
	if base.B2i32(v6218 == int32(0))|base.B2i32(v6218 != v6221) != 0 {
		v6239 = v6218
		v6240 = v6221
		goto L1735
	} else {
		goto L1736
	}
L1711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+12)) = int32(21)
	v6213 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6214 = m.ExcPending
	if v6214 != 0 {
		goto L46
	} else {
		goto L1733
	}
L1712:
	;
	v6173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v6173&int32(1) != 0 {
		goto L1698
	} else {
		goto L1723
	}
L1713:
	;
	if v6157 <= int32(303) {
		goto L1714
	} else {
		goto L1715
	}
L1714:
	;
	if v6157 == int32(59) {
		goto L1689
	} else {
		goto L1717
	}
L1715:
	;
	goto L1716
L1716:
	;
	switch v6157 - int32(304) {
	case 0:
		goto L1701
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30, 32, 33, 34, 35, 36, 37, 38, 39:
		v6391 = v6157
		goto L1699
	case 12:
		goto L1711
	case 26:
		goto L1705
	case 31:
		goto L1703
	case 40:
		goto L1707
	default:
		goto L1721
	}
L1717:
	;
	if v6157 == int32(277) {
		goto L1712
	} else {
		goto L1718
	}
L1718:
	;
	if v6157 != 0 {
		v6391 = v6157
		goto L1699
	} else {
		goto L1719
	}
L1719:
	;
	F_plpgsql_yyerror(m, v6156, int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v6168 = m.ExcPending
	if v6168 != 0 {
		goto L46
	} else {
		goto L1720
	}
L1720:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1721:
	;
	if v6157 == int32(383) {
		goto L1709
	} else {
		goto L1722
	}
L1722:
	;
	v6391 = v6157
	goto L1699
L1723:
	;
	v6176 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6176 == int32(0) {
		goto L1698
	} else {
		goto L1724
	}
L1724:
	;
	v6179 = int32(_a_F_plpgsql_yyparse_78)
	v6182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6176))))
	v6185 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[46])))
	if base.B2i32(v6182 == int32(0))|base.B2i32(v6182 != v6185) != 0 {
		v6203 = v6182
		v6204 = v6185
		goto L1726
	} else {
		goto L1727
	}
L1725:
	;
	if v6203-v6204 != 0 {
		goto L1710
	} else {
		goto L1732
	}
L1726:
	;
	goto L1725
L1727:
	;
	v6188 = v6176
	v6189 = v6179
	goto L1728
L1728:
	;
	v6192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6189)+1)))
	v6193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6188)+1)))
	if v6193 == int32(0) {
		v6203 = v6193
		v6204 = v6192
		goto L1726
	} else {
		goto L1730
	}
L1729:
	;
	v6203 = v6193
	v6204 = v6192
	goto L1726
L1730:
	;
	v6196 = int32(1)
	if v6193 == v6192 {
		v6188 = v6188 + v6196
		v6189 = v6189 + v6196
		goto L1728
	} else {
		goto L1731
	}
L1731:
	;
	goto L1729
L1732:
	;
	goto L1711
L1733:
	;
	v6391 = v6213
	goto L1699
L1734:
	;
	if v6239-v6240 != 0 {
		goto L1708
	} else {
		goto L1741
	}
L1735:
	;
	goto L1734
L1736:
	;
	v6224 = v6176
	v6225 = v6215
	goto L1737
L1737:
	;
	v6228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6225)+1)))
	v6229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6224)+1)))
	if v6229 == int32(0) {
		v6239 = v6229
		v6240 = v6228
		goto L1735
	} else {
		goto L1739
	}
L1738:
	;
	v6239 = v6229
	v6240 = v6228
	goto L1735
L1739:
	;
	v6232 = int32(1)
	if v6229 == v6228 {
		v6224 = v6224 + v6232
		v6225 = v6225 + v6232
		goto L1737
	} else {
		goto L1740
	}
L1740:
	;
	goto L1738
L1741:
	;
	goto L1709
L1742:
	;
	v6391 = v6249
	goto L1699
L1743:
	;
	v6254 = int32(_a_F_plpgsql_yyparse_79)
	v6257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6251))))
	v6260 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[47])))
	if base.B2i32(v6257 == int32(0))|base.B2i32(v6257 != v6260) != 0 {
		v6278 = v6257
		v6279 = v6260
		goto L1745
	} else {
		goto L1746
	}
L1744:
	;
	if v6278-v6279 != 0 {
		goto L1706
	} else {
		goto L1751
	}
L1745:
	;
	goto L1744
L1746:
	;
	v6263 = v6251
	v6264 = v6254
	goto L1747
L1747:
	;
	v6267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6264)+1)))
	v6268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6263)+1)))
	if v6268 == int32(0) {
		v6278 = v6268
		v6279 = v6267
		goto L1745
	} else {
		goto L1749
	}
L1748:
	;
	v6278 = v6268
	v6279 = v6267
	goto L1745
L1749:
	;
	v6271 = int32(1)
	if v6268 == v6267 {
		v6263 = v6263 + v6271
		v6264 = v6264 + v6271
		goto L1747
	} else {
		goto L1750
	}
L1750:
	;
	goto L1748
L1751:
	;
	goto L1707
L1752:
	;
	v6391 = v6288
	goto L1699
L1753:
	;
	if v6314-v6315 != 0 {
		goto L1704
	} else {
		goto L1760
	}
L1754:
	;
	goto L1753
L1755:
	;
	v6299 = v6251
	v6300 = v6290
	goto L1756
L1756:
	;
	v6303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6300)+1)))
	v6304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6299)+1)))
	if v6304 == int32(0) {
		v6314 = v6304
		v6315 = v6303
		goto L1754
	} else {
		goto L1758
	}
L1757:
	;
	v6314 = v6304
	v6315 = v6303
	goto L1754
L1758:
	;
	v6307 = int32(1)
	if v6304 == v6303 {
		v6299 = v6299 + v6307
		v6300 = v6300 + v6307
		goto L1756
	} else {
		goto L1759
	}
L1759:
	;
	goto L1757
L1760:
	;
	goto L1705
L1761:
	;
	v6323 = int32(_a_F_plpgsql_yyparse_80)
	v6326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6320))))
	v6329 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[48])))
	if base.B2i32(v6326 == int32(0))|base.B2i32(v6326 != v6329) != 0 {
		v6347 = v6326
		v6348 = v6329
		goto L1763
	} else {
		goto L1764
	}
L1762:
	;
	if v6347-v6348 != 0 {
		goto L1702
	} else {
		goto L1769
	}
L1763:
	;
	goto L1762
L1764:
	;
	v6332 = v6320
	v6333 = v6323
	goto L1765
L1765:
	;
	v6336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6333)+1)))
	v6337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6332)+1)))
	if v6337 == int32(0) {
		v6347 = v6337
		v6348 = v6336
		goto L1763
	} else {
		goto L1767
	}
L1766:
	;
	v6347 = v6337
	v6348 = v6336
	goto L1763
L1767:
	;
	v6340 = int32(1)
	if v6337 == v6336 {
		v6332 = v6332 + v6340
		v6333 = v6333 + v6340
		goto L1765
	} else {
		goto L1768
	}
L1768:
	;
	goto L1766
L1769:
	;
	goto L1703
L1770:
	;
	if v6377-v6378 != 0 {
		goto L1698
	} else {
		goto L1777
	}
L1771:
	;
	goto L1770
L1772:
	;
	v6362 = v6320
	v6363 = v6353
	goto L1773
L1773:
	;
	v6366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6363)+1)))
	v6367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6362)+1)))
	if v6367 == int32(0) {
		v6377 = v6367
		v6378 = v6366
		goto L1771
	} else {
		goto L1775
	}
L1774:
	;
	v6377 = v6367
	v6378 = v6366
	goto L1771
L1775:
	;
	v6370 = int32(1)
	if v6367 == v6366 {
		v6362 = v6362 + v6370
		v6363 = v6363 + v6370
		goto L1773
	} else {
		goto L1776
	}
L1776:
	;
	goto L1774
L1777:
	;
	goto L1701
L1778:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v6388
	v6391 = v6388
	goto L1699
L1779:
	;
	v6396 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+20)) = v6396
	v6402 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6403 = m.ExcPending
	if v6403 != 0 {
		goto L46
	} else {
		goto L1783
	}
L1780:
	;
	switch v6391 - int32(371) {
	case 0:
		goto L1697
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v6545 = v6391
		goto L1696
	case 10:
		goto L1690
	default:
		goto L1781
	}
L1781:
	;
	if v6391 != 0 {
		v6545 = v6391
		goto L1696
	} else {
		goto L1782
	}
L1782:
	;
	goto L4
L1783:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v6402
	switch v6402 - int32(44) {
	case 0:
		goto L1692
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L1784
	case 15:
		v6646 = v6402
		goto L1691
	default:
		goto L1785
	}
L1784:
	;
	goto L3
L1785:
	;
	if v6402 == int32(381) {
		v6646 = v6402
		goto L1691
	} else {
		goto L1786
	}
L1786:
	;
	goto L1784
L1787:
	;
	v6414 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6414 == int32(0) {
		v6545 = v6410
		goto L1696
	} else {
		goto L1788
	}
L1788:
	;
	v6417 = int32(_a_F_plpgsql_yyparse_17)
	v6420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6414))))
	v6423 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[11])))
	if base.B2i32(v6420 == int32(0))|base.B2i32(v6420 != v6423) != 0 {
		v6441 = v6420
		v6442 = v6423
		goto L1790
	} else {
		goto L1791
	}
L1789:
	;
	if v6441-v6442 != 0 {
		v6545 = v6410
		goto L1696
	} else {
		goto L1796
	}
L1790:
	;
	goto L1789
L1791:
	;
	v6426 = v6414
	v6427 = v6417
	goto L1792
L1792:
	;
	v6430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6427)+1)))
	v6431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6426)+1)))
	if v6431 == int32(0) {
		v6441 = v6431
		v6442 = v6430
		goto L1790
	} else {
		goto L1794
	}
L1793:
	;
	v6441 = v6431
	v6442 = v6430
	goto L1790
L1794:
	;
	v6434 = int32(1)
	if v6431 == v6430 {
		v6426 = v6426 + v6434
		v6427 = v6427 + v6434
		goto L1792
	} else {
		goto L1795
	}
L1795:
	;
	goto L1793
L1796:
	;
	goto L1697
L1797:
	;
	if v6450 != int32(261) {
		goto L3
	} else {
		goto L1798
	}
L1798:
	;
	v6454 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6455 = F_strlen(m, v6454)
	mBase = m.M
	if v6455 != int32(5) {
		goto L1
	} else {
		goto L1799
	}
L1799:
	;
	v6458 = int32(_a_F_plpgsql_yyparse_81)
	v6462 = m.G0
	v6464 = v6462 - int32(32)
	v6465 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6464)+24)) = v6465
	*(*int64)(unsafe.Add(mBase, uint32(v6464)+16)) = v6465
	*(*int64)(unsafe.Add(mBase, uint32(v6464)+8)) = v6465
	*(*int64)(unsafe.Add(mBase, uint32(v6464))) = v6465
	v6473 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[49])))
	if v6473 == int32(0) {
		goto L1801
	} else {
		goto L1802
	}
L1800:
	;
	if v6541 != int32(5) {
		goto L1
	} else {
		goto L1819
	}
L1801:
	;
	v6541 = int32(0)
	goto L1800
L1802:
	;
	goto L1803
L1803:
	;
	v6477 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[50])))
	if v6477 == int32(0) {
		goto L1804
	} else {
		goto L1805
	}
L1804:
	;
	v6481 = v6454
	goto L1807
L1805:
	;
	goto L1806
L1806:
	;
	v6491 = v6458
	v6492 = v6473
	goto L1810
L1807:
	;
	v6487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6481))))
	if v6487 == v6473 {
		v6481 = v6481 + int32(1)
		goto L1807
	} else {
		goto L1809
	}
L1808:
	;
	v6541 = v6481 - v6454
	goto L1800
L1809:
	;
	goto L1808
L1810:
	;
	v6499 = v6464 + int32(base.Ui32(v6492)>>(uint(int32(3))%32))&int32(28)
	v6500 = *(*int32)(unsafe.Add(mBase, uint32(v6499)))
	v6501 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6499))) = v6500 | v6501<<(uint(v6492)%32)
	v6505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6491)+1)))
	if v6505 != 0 {
		v6491 = v6491 + v6501
		v6492 = v6505
		goto L1810
	} else {
		goto L1812
	}
L1811:
	;
	v6508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6454))))
	if v6508 == int32(0) {
		v6531 = v6454
		goto L1813
	} else {
		goto L1814
	}
L1812:
	;
	goto L1811
L1813:
	;
	v6541 = v6531 - v6454
	goto L1800
L1814:
	;
	v6512 = v6454
	v6513 = v6508
	goto L1815
L1815:
	;
	v6521 = *(*int32)(unsafe.Add(mBase, uint32(v6464+int32(base.Ui32(v6513)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v6521)>>(uint(v6513)%32))&int32(1) == int32(0) {
		v6531 = v6512
		goto L1813
	} else {
		goto L1817
	}
L1816:
	;
	v6531 = v6529
	goto L1813
L1817:
	;
	v6527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6512)+1)))
	v6529 = v6512 + int32(1)
	if v6527 != 0 {
		v6512 = v6529
		v6513 = v6527
		goto L1815
	} else {
		goto L1818
	}
L1818:
	;
	goto L1816
L1819:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6152))) = v6454
	goto L1693
L1820:
	;
	if v6545 == v6556 {
		goto L1827
	} else {
		goto L1828
	}
L1821:
	;
	v6556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6550<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[20]))))
	v6557 = base.B2i32(v6545 == v6556)
	if v6557 == int32(0) {
		goto L1823
	} else {
		goto L1824
	}
L1822:
	;
	goto L1820
L1823:
	;
	v6561 = v6550 + int32(1)
	if v6561 != int32(83) {
		v6550 = v6561
		goto L1821
	} else {
		goto L1826
	}
L1824:
	;
	goto L1825
L1825:
	;
	goto L1822
L1826:
	;
	goto L1825
L1827:
	;
	v6565 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6566 = F_pstrdup(m, v6565)
	mBase = m.M
	v6567 = m.ExcPending
	if v6567 != 0 {
		goto L46
	} else {
		goto L1830
	}
L1828:
	;
	goto L1829
L1829:
	;
	goto L3
L1830:
	;
	v6571 = v6566
	goto L1694
L1831:
	;
	goto L1693
L1832:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v6582
	if base.B2i32(v6582 == int32(59))|base.B2i32(v6582 == int32(381)) != 0 {
		v6646 = v6582
		goto L1691
	} else {
		goto L1833
	}
L1833:
	;
	F_plpgsql_yyerror(m, v6581, int32(0), l1, int32(_a_F_plpgsql_yyparse_5))
	mBase = m.M
	v6593 = m.ExcPending
	if v6593 != 0 {
		goto L46
	} else {
		goto L1834
	}
L1834:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1835:
	;
	v6624 = int32(1)
	v6633 = F_read_sql_construct(m, int32(44), int32(59), int32(381), int32(_a_F_plpgsql_yyparse_82), int32(2), v6624, v6624, int32(0), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6634 = m.ExcPending
	if v6634 != 0 {
		goto L46
	} else {
		goto L1837
	}
L1836:
	;
	v6646 = v6639
	goto L1691
L1837:
	;
	v6635 = *(*int32)(unsafe.Add(mBase, uint32(v6085)+24))
	v6636 = F_lappend(m, v6635, v6633)
	mBase = m.M
	v6637 = m.ExcPending
	if v6637 != 0 {
		goto L46
	} else {
		goto L1838
	}
L1838:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+24)) = v6636
	v6639 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v6639 == int32(44) {
		goto L1835
	} else {
		goto L1839
	}
L1839:
	;
	goto L1836
L1840:
	;
	goto L1690
L1841:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6085)+28)) = v7046
	goto L1689
L1842:
	;
	F_plpgsql_yyerror(m, v6697, int32(0), l1, int32(_a_F_plpgsql_yyparse_83))
	mBase = m.M
	v7067 = m.ExcPending
	if v7067 != 0 {
		goto L46
	} else {
		goto L1952
	}
L1843:
	;
	F_plpgsql_yyerror(m, v6697, int32(0), l1, int32(_a_F_plpgsql_yyparse_84))
	mBase = m.M
	v7063 = m.ExcPending
	if v7063 != 0 {
		goto L46
	} else {
		goto L1951
	}
L1844:
	;
	F_plpgsql_yyerror(m, v6697, int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v7057 = m.ExcPending
	if v7057 != 0 {
		goto L46
	} else {
		goto L1950
	}
L1845:
	;
	v6728 = F_plpgsql_yylex(m, v6695, v6697, l1)
	mBase = m.M
	v6729 = m.ExcPending
	if v6729 != 0 {
		goto L46
	} else {
		goto L1847
	}
L1846:
	;
	m.G0 = v6701 + int32(16)
	goto L1841
L1847:
	;
	if v6728 == int32(0) {
		goto L1844
	} else {
		goto L1848
	}
L1848:
	;
	v6732 = int32(0)
	v6734 = F_palloc(m, int32(8))
	mBase = m.M
	v6735 = m.ExcPending
	if v6735 != 0 {
		goto L46
	} else {
		goto L1849
	}
L1849:
	;
	switch v6728 - int32(277) {
	case 0:
		goto L1866
	default:
		goto L1843
	case 17:
		goto L1859
	case 21:
		goto L1857
	case 26:
		goto L1855
	case 30:
		goto L1863
	case 37:
		v7022 = v6732
		goto L1850
	case 49:
		goto L1861
	case 61:
		goto L1865
	case 90:
		goto L1851
	case 97:
		goto L1853
	}
L1850:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6734))) = v7022
	v7025 = F_plpgsql_yylex(m, v6695, v6697, l1)
	mBase = m.M
	v7026 = m.ExcPending
	if v7026 != 0 {
		goto L46
	} else {
		goto L1945
	}
L1851:
	;
	v7022 = int32(8)
	goto L1850
L1852:
	;
	v6989 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6989 == int32(0) {
		goto L1843
	} else {
		goto L1936
	}
L1853:
	;
	v7022 = int32(7)
	goto L1850
L1854:
	;
	v6959 = int32(_a_F_plpgsql_yyparse_85)
	v6962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6926))))
	v6965 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[51])))
	if base.B2i32(v6962 == int32(0))|base.B2i32(v6962 != v6965) != 0 {
		v6983 = v6962
		v6984 = v6965
		goto L1929
	} else {
		goto L1930
	}
L1855:
	;
	v7022 = int32(6)
	goto L1850
L1856:
	;
	v6926 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6926 == int32(0) {
		goto L1843
	} else {
		goto L1919
	}
L1857:
	;
	v7022 = int32(5)
	goto L1850
L1858:
	;
	v6896 = int32(_a_F_plpgsql_yyparse_86)
	v6899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6863))))
	v6902 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[52])))
	if base.B2i32(v6899 == int32(0))|base.B2i32(v6899 != v6902) != 0 {
		v6920 = v6899
		v6921 = v6902
		goto L1912
	} else {
		goto L1913
	}
L1859:
	;
	v7022 = int32(4)
	goto L1850
L1860:
	;
	v6863 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6863 == int32(0) {
		goto L1843
	} else {
		goto L1902
	}
L1861:
	;
	v7022 = int32(3)
	goto L1850
L1862:
	;
	v6833 = int32(_a_F_plpgsql_yyparse_87)
	v6836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6800))))
	v6839 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[53])))
	if base.B2i32(v6836 == int32(0))|base.B2i32(v6836 != v6839) != 0 {
		v6857 = v6836
		v6858 = v6839
		goto L1895
	} else {
		goto L1896
	}
L1863:
	;
	v7022 = int32(2)
	goto L1850
L1864:
	;
	v6800 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6800 == int32(0) {
		goto L1843
	} else {
		goto L1885
	}
L1865:
	;
	v7022 = int32(1)
	goto L1850
L1866:
	;
	v6738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v6738 != 0 {
		goto L1843
	} else {
		goto L1867
	}
L1867:
	;
	v6739 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v6739 == int32(0) {
		goto L1843
	} else {
		goto L1868
	}
L1868:
	;
	v6742 = int32(_a_F_plpgsql_yyparse_88)
	v6745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6739))))
	v6748 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[54])))
	if base.B2i32(v6745 == int32(0))|base.B2i32(v6745 != v6748) != 0 {
		v6766 = v6745
		v6767 = v6748
		goto L1870
	} else {
		goto L1871
	}
L1869:
	;
	if v6766-v6767 == int32(0) {
		v7022 = v6732
		goto L1850
	} else {
		goto L1876
	}
L1870:
	;
	goto L1869
L1871:
	;
	v6751 = v6739
	v6752 = v6742
	goto L1872
L1872:
	;
	v6755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6752)+1)))
	v6756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6751)+1)))
	if v6756 == int32(0) {
		v6766 = v6756
		v6767 = v6755
		goto L1870
	} else {
		goto L1874
	}
L1873:
	;
	v6766 = v6756
	v6767 = v6755
	goto L1870
L1874:
	;
	v6759 = int32(1)
	if v6756 == v6755 {
		v6751 = v6751 + v6759
		v6752 = v6752 + v6759
		goto L1872
	} else {
		goto L1875
	}
L1875:
	;
	goto L1873
L1876:
	;
	v6771 = int32(_a_F_plpgsql_yyparse_89)
	v6774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6739))))
	v6777 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[55])))
	if base.B2i32(v6774 == int32(0))|base.B2i32(v6774 != v6777) != 0 {
		v6795 = v6774
		v6796 = v6777
		goto L1878
	} else {
		goto L1879
	}
L1877:
	;
	if v6795-v6796 != 0 {
		goto L1864
	} else {
		goto L1884
	}
L1878:
	;
	goto L1877
L1879:
	;
	v6780 = v6739
	v6781 = v6771
	goto L1880
L1880:
	;
	v6784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6781)+1)))
	v6785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6780)+1)))
	if v6785 == int32(0) {
		v6795 = v6785
		v6796 = v6784
		goto L1878
	} else {
		goto L1882
	}
L1881:
	;
	v6795 = v6785
	v6796 = v6784
	goto L1878
L1882:
	;
	v6788 = int32(1)
	if v6785 == v6784 {
		v6780 = v6780 + v6788
		v6781 = v6781 + v6788
		goto L1880
	} else {
		goto L1883
	}
L1883:
	;
	goto L1881
L1884:
	;
	goto L1865
L1885:
	;
	v6803 = int32(_a_F_plpgsql_yyparse_90)
	v6806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6800))))
	v6809 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[56])))
	if base.B2i32(v6806 == int32(0))|base.B2i32(v6806 != v6809) != 0 {
		v6827 = v6806
		v6828 = v6809
		goto L1887
	} else {
		goto L1888
	}
L1886:
	;
	if v6827-v6828 != 0 {
		goto L1862
	} else {
		goto L1893
	}
L1887:
	;
	goto L1886
L1888:
	;
	v6812 = v6800
	v6813 = v6803
	goto L1889
L1889:
	;
	v6816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6813)+1)))
	v6817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6812)+1)))
	if v6817 == int32(0) {
		v6827 = v6817
		v6828 = v6816
		goto L1887
	} else {
		goto L1891
	}
L1890:
	;
	v6827 = v6817
	v6828 = v6816
	goto L1887
L1891:
	;
	v6820 = int32(1)
	if v6817 == v6816 {
		v6812 = v6812 + v6820
		v6813 = v6813 + v6820
		goto L1889
	} else {
		goto L1892
	}
L1892:
	;
	goto L1890
L1893:
	;
	goto L1863
L1894:
	;
	if v6857-v6858 != 0 {
		goto L1860
	} else {
		goto L1901
	}
L1895:
	;
	goto L1894
L1896:
	;
	v6842 = v6800
	v6843 = v6833
	goto L1897
L1897:
	;
	v6846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6843)+1)))
	v6847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6842)+1)))
	if v6847 == int32(0) {
		v6857 = v6847
		v6858 = v6846
		goto L1895
	} else {
		goto L1899
	}
L1898:
	;
	v6857 = v6847
	v6858 = v6846
	goto L1895
L1899:
	;
	v6850 = int32(1)
	if v6847 == v6846 {
		v6842 = v6842 + v6850
		v6843 = v6843 + v6850
		goto L1897
	} else {
		goto L1900
	}
L1900:
	;
	goto L1898
L1901:
	;
	goto L1861
L1902:
	;
	v6866 = int32(_a_F_plpgsql_yyparse_91)
	v6869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6863))))
	v6872 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[57])))
	if base.B2i32(v6869 == int32(0))|base.B2i32(v6869 != v6872) != 0 {
		v6890 = v6869
		v6891 = v6872
		goto L1904
	} else {
		goto L1905
	}
L1903:
	;
	if v6890-v6891 != 0 {
		goto L1858
	} else {
		goto L1910
	}
L1904:
	;
	goto L1903
L1905:
	;
	v6875 = v6863
	v6876 = v6866
	goto L1906
L1906:
	;
	v6879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6876)+1)))
	v6880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6875)+1)))
	if v6880 == int32(0) {
		v6890 = v6880
		v6891 = v6879
		goto L1904
	} else {
		goto L1908
	}
L1907:
	;
	v6890 = v6880
	v6891 = v6879
	goto L1904
L1908:
	;
	v6883 = int32(1)
	if v6880 == v6879 {
		v6875 = v6875 + v6883
		v6876 = v6876 + v6883
		goto L1906
	} else {
		goto L1909
	}
L1909:
	;
	goto L1907
L1910:
	;
	goto L1859
L1911:
	;
	if v6920-v6921 != 0 {
		goto L1856
	} else {
		goto L1918
	}
L1912:
	;
	goto L1911
L1913:
	;
	v6905 = v6863
	v6906 = v6896
	goto L1914
L1914:
	;
	v6909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6906)+1)))
	v6910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6905)+1)))
	if v6910 == int32(0) {
		v6920 = v6910
		v6921 = v6909
		goto L1912
	} else {
		goto L1916
	}
L1915:
	;
	v6920 = v6910
	v6921 = v6909
	goto L1912
L1916:
	;
	v6913 = int32(1)
	if v6910 == v6909 {
		v6905 = v6905 + v6913
		v6906 = v6906 + v6913
		goto L1914
	} else {
		goto L1917
	}
L1917:
	;
	goto L1915
L1918:
	;
	goto L1857
L1919:
	;
	v6929 = int32(_a_F_plpgsql_yyparse_92)
	v6932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6926))))
	v6935 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[58])))
	if base.B2i32(v6932 == int32(0))|base.B2i32(v6932 != v6935) != 0 {
		v6953 = v6932
		v6954 = v6935
		goto L1921
	} else {
		goto L1922
	}
L1920:
	;
	if v6953-v6954 != 0 {
		goto L1854
	} else {
		goto L1927
	}
L1921:
	;
	goto L1920
L1922:
	;
	v6938 = v6926
	v6939 = v6929
	goto L1923
L1923:
	;
	v6942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6939)+1)))
	v6943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6938)+1)))
	if v6943 == int32(0) {
		v6953 = v6943
		v6954 = v6942
		goto L1921
	} else {
		goto L1925
	}
L1924:
	;
	v6953 = v6943
	v6954 = v6942
	goto L1921
L1925:
	;
	v6946 = int32(1)
	if v6943 == v6942 {
		v6938 = v6938 + v6946
		v6939 = v6939 + v6946
		goto L1923
	} else {
		goto L1926
	}
L1926:
	;
	goto L1924
L1927:
	;
	goto L1855
L1928:
	;
	if v6983-v6984 != 0 {
		goto L1852
	} else {
		goto L1935
	}
L1929:
	;
	goto L1928
L1930:
	;
	v6968 = v6926
	v6969 = v6959
	goto L1931
L1931:
	;
	v6972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6969)+1)))
	v6973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6968)+1)))
	if v6973 == int32(0) {
		v6983 = v6973
		v6984 = v6972
		goto L1929
	} else {
		goto L1933
	}
L1932:
	;
	v6983 = v6973
	v6984 = v6972
	goto L1929
L1933:
	;
	v6976 = int32(1)
	if v6973 == v6972 {
		v6968 = v6968 + v6976
		v6969 = v6969 + v6976
		goto L1931
	} else {
		goto L1934
	}
L1934:
	;
	goto L1932
L1935:
	;
	goto L1853
L1936:
	;
	v6992 = int32(_a_F_plpgsql_yyparse_93)
	v6995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6989))))
	v6998 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[59])))
	if base.B2i32(v6995 == int32(0))|base.B2i32(v6995 != v6998) != 0 {
		v7016 = v6995
		v7017 = v6998
		goto L1938
	} else {
		goto L1939
	}
L1937:
	;
	if v7016-v7017 != 0 {
		goto L1843
	} else {
		goto L1944
	}
L1938:
	;
	goto L1937
L1939:
	;
	v7001 = v6989
	v7002 = v6992
	goto L1940
L1940:
	;
	v7005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7002)+1)))
	v7006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7001)+1)))
	if v7006 == int32(0) {
		v7016 = v7006
		v7017 = v7005
		goto L1938
	} else {
		goto L1942
	}
L1941:
	;
	v7016 = v7006
	v7017 = v7005
	goto L1938
L1942:
	;
	v7009 = int32(1)
	if v7006 == v7005 {
		v7001 = v7001 + v7009
		v7002 = v7002 + v7009
		goto L1940
	} else {
		goto L1943
	}
L1943:
	;
	goto L1941
L1944:
	;
	goto L1851
L1945:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6701)+12)) = v7025
	if base.B2i32(v7025 != int32(61))&base.B2i32(v7025 != int32(270)) != 0 {
		goto L1842
	} else {
		goto L1946
	}
L1946:
	;
	v7035 = int32(0)
	v7038 = int32(1)
	v7043 = F_read_sql_construct(m, int32(44), int32(59), v7035, int32(_a_F_plpgsql_yyparse_66), int32(2), v7038, v7038, v7035, v6701+int32(12), v6695, v6697, l1)
	mBase = m.M
	v7044 = m.ExcPending
	if v7044 != 0 {
		goto L46
	} else {
		goto L1947
	}
L1947:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6734)+4)) = v7043
	v7046 = F_lappend(m, v6707, v6734)
	mBase = m.M
	v7047 = m.ExcPending
	if v7047 != 0 {
		goto L46
	} else {
		goto L1948
	}
L1948:
	;
	v7048 = *(*int32)(unsafe.Add(mBase, uint32(v6701)+12))
	if v7048 != int32(59) {
		v6707 = v7046
		goto L1845
	} else {
		goto L1949
	}
L1949:
	;
	goto L1846
L1950:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1951:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1952:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v6085
	v9708 = v233
	goto L12
L1954:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v7189 = m.ExcPending
	if v7189 != 0 {
		goto L46
	} else {
		goto L1979
	}
L1955:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v7172 = m.ExcPending
	if v7172 != 0 {
		goto L46
	} else {
		goto L1975
	}
L1956:
	;
	v7100 = v7095
	v7101 = int32(0)
	goto L1960
L1957:
	;
	goto L1958
L1958:
	;
	goto L1953
L1959:
	;
	if v7142 < v7101 {
		goto L1955
	} else {
		goto L1974
	}
L1960:
	;
	v7121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7100))))
	if v7121 != int32(37) {
		goto L1964
	} else {
		goto L1965
	}
L1961:
	;
	v7140 = *(*int32)(unsafe.Add(mBase, uint32(v7124)+4))
	if v7101 < v7140 {
		goto L1954
	} else {
		goto L1973
	}
L1962:
	;
	goto L1961
L1963:
	;
	v7100 = v7135 + int32(1)
	v7101 = v7136
	goto L1960
L1964:
	;
	if v7121 != 0 {
		v7135 = v7100
		v7136 = v7101
		goto L1963
	} else {
		goto L1967
	}
L1965:
	;
	goto L1966
L1966:
	;
	v7130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7100)+1)))
	v7132 = base.B2i32(v7130 != int32(37))
	if v7130 != int32(37) {
		goto L1970
	} else {
		goto L1971
	}
L1967:
	;
	v7124 = *(*int32)(unsafe.Add(mBase, uint32(v6085)+24))
	if v7124 != 0 {
		goto L1962
	} else {
		goto L1968
	}
L1968:
	;
	v7125 = int32(0)
	if v7125 <= v7101 {
		v7142 = v7125
		goto L1959
	} else {
		goto L1969
	}
L1969:
	;
	goto L1954
L1970:
	;
	v7133 = v7100
	goto L1972
L1971:
	;
	v7133 = v7100 + int32(1)
	goto L1972
L1972:
	;
	v7135 = v7133
	v7136 = v7101 + v7132
	goto L1963
L1973:
	;
	v7142 = v7140
	goto L1959
L1974:
	;
	goto L1958
L1975:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7175 = m.ExcPending
	if v7175 != 0 {
		goto L46
	} else {
		goto L1976
	}
L1976:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_94), int32(0))
	mBase = m.M
	v7179 = m.ExcPending
	if v7179 != 0 {
		goto L46
	} else {
		goto L1977
	}
L1977:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(_a_F_plpgsql_yyparse_95), int32(_a_F_plpgsql_yyparse_96))
	mBase = m.M
	v7184 = m.ExcPending
	if v7184 != 0 {
		goto L46
	} else {
		goto L1978
	}
L1978:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1979:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7192 = m.ExcPending
	if v7192 != 0 {
		goto L46
	} else {
		goto L1980
	}
L1980:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_97), int32(0))
	mBase = m.M
	v7196 = m.ExcPending
	if v7196 != 0 {
		goto L46
	} else {
		goto L1981
	}
L1981:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(_a_F_plpgsql_yyparse_98), int32(_a_F_plpgsql_yyparse_96))
	mBase = m.M
	v7201 = m.ExcPending
	if v7201 != 0 {
		goto L46
	} else {
		goto L1982
	}
L1982:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1983:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7204))) = int32(15)
	v7208 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7209 = int32(0)
	if v7208 < v7209 {
		v7254 = v7209
		goto L1985
	} else {
		goto L1986
	}
L1984:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7204)+4)) = v7254
	v7258 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v7259 = *(*int32)(unsafe.Add(mBase, uint32(v7258)+520))
	v7260 = int32(1)
	v7261 = v7259 + v7260
	*(*int32)(unsafe.Add(mBase, uint32(v7258)+520)) = v7261
	*(*int32)(unsafe.Add(mBase, uint32(v7204)+8)) = v7261
	v7266 = int32(0)
	v7275 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v7277 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v7278 = F_read_sql_construct(m, int32(44), int32(59), v7266, int32(_a_F_plpgsql_yyparse_66), int32(2), v7260, v7260, v7266, v28+int32(256), v7275, v7277, l1)
	mBase = m.M
	v7279 = m.ExcPending
	if v7279 != 0 {
		goto L46
	} else {
		goto L1997
	}
L1985:
	;
	goto L1984
L1986:
	;
	v7214 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7215 = *(*int32)(unsafe.Add(mBase, uint32(v7214)+60))
	if v7215 == int32(0) {
		v7254 = v7209
		goto L1985
	} else {
		goto L1987
	}
L1987:
	;
	v7218 = v7208 + v7215
	v7219 = *(*int32)(unsafe.Add(mBase, uint32(v7214)+188))
	if base.Ui32(v7219) <= base.Ui32(v7218) {
		goto L1989
	} else {
		goto L1990
	}
L1988:
	;
	v7229 = *(*int32)(unsafe.Add(mBase, uint32(v7214)+196))
	if base.B2i32(v7228 == int32(0))|base.B2i32(base.Ui32(v7218) <= base.Ui32(v7228)) != 0 {
		v7254 = v7229
		goto L1985
	} else {
		goto L1992
	}
L1989:
	;
	v7221 = *(*int32)(unsafe.Add(mBase, uint32(v7214)+192))
	v7228 = v7221
	goto L1988
L1990:
	;
	goto L1991
L1991:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7214)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7214)+188)) = v7215
	v7226 = F_strchr(m, v7215, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7214)+192)) = v7226
	v7228 = v7226
	goto L1988
L1992:
	;
	v7234 = v7228
	v7237 = v7229
	goto L1993
L1993:
	;
	v7239 = int32(1)
	v7240 = v7237 + v7239
	*(*int32)(unsafe.Add(mBase, uint32(v7214)+196)) = v7240
	v7243 = v7234 + v7239
	*(*int32)(unsafe.Add(mBase, uint32(v7214)+188)) = v7243
	v7246 = F_strchr(m, v7243, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7214)+192)) = v7246
	if v7246 == int32(0) {
		v7254 = v7240
		goto L1985
	} else {
		goto L1995
	}
L1994:
	;
	v7254 = v7240
	goto L1985
L1995:
	;
	if base.Ui32(v7246) < base.Ui32(v7218) {
		v7234 = v7246
		v7237 = v7240
		goto L1993
	} else {
		goto L1996
	}
L1996:
	;
	goto L1994
L1997:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7204)+12)) = v7278
	v7281 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7281 == int32(44) {
		goto L1998
	} else {
		goto L1999
	}
L1998:
	;
	v7285 = int32(0)
	v7289 = int32(1)
	v7293 = F_read_sql_construct(m, int32(59), v7285, v7285, int32(_a_F_plpgsql_yyparse_16), int32(2), v7289, v7289, v7285, v7285, v7275, v7277, l1)
	mBase = m.M
	v7294 = m.ExcPending
	if v7294 != 0 {
		goto L46
	} else {
		goto L2001
	}
L1999:
	;
	v7296 = int32(0)
	goto L2000
L2000:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7204)+16)) = v7296
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7204
	v9708 = v233
	goto L12
L2001:
	;
	v7296 = v7293
	goto L2000
L2002:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7318
	v9708 = v233
	goto L12
L2003:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7328
	v9708 = v233
	goto L12
L2004:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7338
	v9708 = v233
	goto L12
L2005:
	;
	F_plpgsql_push_back_token(m, v7345, v7342, v7344, l1)
	mBase = m.M
	v7348 = m.ExcPending
	if v7348 != 0 {
		goto L46
	} else {
		goto L2006
	}
L2006:
	;
	v7349 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	switch v7345 - int32(46) {
	case 0, 15:
		goto L22
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2007
	default:
		goto L2008
	}
L2007:
	;
	v7362 = F_make_execsql_stmt(m, int32(275), v7349, v148, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7363 = m.ExcPending
	if v7363 != 0 {
		goto L46
	} else {
		goto L2010
	}
L2008:
	;
	if base.B2i32(v7345 == int32(270))|base.B2i32(v7345 == int32(91)) != 0 {
		goto L22
	} else {
		goto L2009
	}
L2009:
	;
	goto L2007
L2010:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7362
	v9708 = v233
	goto L12
L2011:
	;
	F_plpgsql_push_back_token(m, v7369, v7366, v7368, l1)
	mBase = m.M
	v7372 = m.ExcPending
	if v7372 != 0 {
		goto L46
	} else {
		goto L2012
	}
L2012:
	;
	v7373 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	switch v7369 - int32(46) {
	case 0, 15:
		goto L21
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2013
	default:
		goto L2014
	}
L2013:
	;
	v7387 = F_make_execsql_stmt(m, int32(276), v7373, int32(0), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7388 = m.ExcPending
	if v7388 != 0 {
		goto L46
	} else {
		goto L2016
	}
L2014:
	;
	if base.B2i32(v7369 == int32(270))|base.B2i32(v7369 == int32(91)) != 0 {
		goto L21
	} else {
		goto L2015
	}
L2015:
	;
	goto L2013
L2016:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7387
	v9708 = v233
	goto L12
L2017:
	;
	v7407 = F_palloc(m, int32(28))
	mBase = m.M
	v7408 = m.ExcPending
	if v7408 != 0 {
		goto L46
	} else {
		goto L2018
	}
L2018:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7407))) = int32(17)
	v7411 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7412 = int32(0)
	if v7411 < v7412 {
		v7457 = v7412
		goto L2020
	} else {
		goto L2021
	}
L2019:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7407)+4)) = v7457
	v7461 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v7462 = *(*int32)(unsafe.Add(mBase, uint32(v7461)+520))
	v7464 = v7462 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7461)+520)) = v7464
	*(*int64)(unsafe.Add(mBase, uint32(v7407)+20)) = int64(0)
	v7468 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7407)+16)) = uint16(v7468)
	*(*int32)(unsafe.Add(mBase, uint32(v7407)+12)) = v7404
	*(*int32)(unsafe.Add(mBase, uint32(v7407)+8)) = v7464
	v7476 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	v7481 = v7476
	goto L2032
L2020:
	;
	goto L2019
L2021:
	;
	v7417 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7418 = *(*int32)(unsafe.Add(mBase, uint32(v7417)+60))
	if v7418 == int32(0) {
		v7457 = v7412
		goto L2020
	} else {
		goto L2022
	}
L2022:
	;
	v7421 = v7411 + v7418
	v7422 = *(*int32)(unsafe.Add(mBase, uint32(v7417)+188))
	if base.Ui32(v7422) <= base.Ui32(v7421) {
		goto L2024
	} else {
		goto L2025
	}
L2023:
	;
	v7432 = *(*int32)(unsafe.Add(mBase, uint32(v7417)+196))
	if base.B2i32(v7431 == int32(0))|base.B2i32(base.Ui32(v7421) <= base.Ui32(v7431)) != 0 {
		v7457 = v7432
		goto L2020
	} else {
		goto L2027
	}
L2024:
	;
	v7424 = *(*int32)(unsafe.Add(mBase, uint32(v7417)+192))
	v7431 = v7424
	goto L2023
L2025:
	;
	goto L2026
L2026:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7417)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7417)+188)) = v7418
	v7429 = F_strchr(m, v7418, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7417)+192)) = v7429
	v7431 = v7429
	goto L2023
L2027:
	;
	v7437 = v7431
	v7440 = v7432
	goto L2028
L2028:
	;
	v7442 = int32(1)
	v7443 = v7440 + v7442
	*(*int32)(unsafe.Add(mBase, uint32(v7417)+196)) = v7443
	v7446 = v7437 + v7442
	*(*int32)(unsafe.Add(mBase, uint32(v7417)+188)) = v7446
	v7449 = F_strchr(m, v7446, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7417)+192)) = v7449
	if v7449 == int32(0) {
		v7457 = v7443
		goto L2020
	} else {
		goto L2030
	}
L2029:
	;
	v7457 = v7443
	goto L2020
L2030:
	;
	if base.Ui32(v7449) < base.Ui32(v7421) {
		v7437 = v7449
		v7440 = v7443
		goto L2028
	} else {
		goto L2031
	}
L2031:
	;
	goto L2029
L2032:
	;
	if v7481 != int32(332) {
		goto L2036
	} else {
		goto L2037
	}
L2033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7407
	v9708 = v233
	goto L12
L2034:
	;
	goto L2033
L2035:
	;
	v7522 = *(*int32)(unsafe.Add(mBase, uint32(v7407)+24))
	if v7522 != 0 {
		goto L3
	} else {
		goto L2044
	}
L2036:
	;
	if v7481 == int32(381) {
		goto L2035
	} else {
		goto L2039
	}
L2037:
	;
	goto L2038
L2038:
	;
	v7508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7407)+16)))
	if v7508 == int32(1) {
		goto L3
	} else {
		goto L2041
	}
L2039:
	;
	if v7481 == int32(59) {
		goto L2034
	} else {
		goto L2040
	}
L2040:
	;
	goto L3
L2041:
	;
	v7511 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7407)+16)) = uint8(v7511)
	v7514 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v7516 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_read_into_target(m, v7407+int32(20), v7407+int32(17), v7514, v7516, l1)
	mBase = m.M
	v7518 = m.ExcPending
	if v7518 != 0 {
		goto L46
	} else {
		goto L2042
	}
L2042:
	;
	v7519 = F_plpgsql_yylex(m, v7514, v7516, l1)
	mBase = m.M
	v7520 = m.ExcPending
	if v7520 != 0 {
		goto L46
	} else {
		goto L2043
	}
L2043:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v7519
	v7481 = v7519
	goto L2032
L2044:
	;
	goto L2045
L2045:
	;
	v7553 = int32(1)
	v7562 = F_read_sql_construct(m, int32(44), int32(59), int32(332), int32(_a_F_plpgsql_yyparse_99), int32(2), v7553, v7553, int32(0), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7563 = m.ExcPending
	if v7563 != 0 {
		goto L46
	} else {
		goto L2047
	}
L2046:
	;
	v7481 = v7568
	goto L2032
L2047:
	;
	v7564 = *(*int32)(unsafe.Add(mBase, uint32(v7407)+24))
	v7565 = F_lappend(m, v7564, v7562)
	mBase = m.M
	v7566 = m.ExcPending
	if v7566 != 0 {
		goto L46
	} else {
		goto L2048
	}
L2048:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7407)+24)) = v7565
	v7568 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7568 == int32(44) {
		goto L2045
	} else {
		goto L2049
	}
L2049:
	;
	goto L2046
L2050:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7573))) = int32(20)
	v7579 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v7580 = int32(0)
	if v7579 < v7580 {
		v7625 = v7580
		goto L2052
	} else {
		goto L2053
	}
L2051:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+4)) = v7625
	v7629 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v7630 = *(*int32)(unsafe.Add(mBase, uint32(v7629)+520))
	v7632 = v7630 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7629)+520)) = v7632
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+8)) = v7632
	v7635 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v7636 = *(*int32)(unsafe.Add(mBase, uint32(v7635)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+16)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+12)) = v7636
	v7640 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v7641 = *(*int32)(unsafe.Add(mBase, uint32(v7640)+28))
	if v7641 == int32(0) {
		goto L2064
	} else {
		goto L2065
	}
L2052:
	;
	goto L2051
L2053:
	;
	v7585 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7586 = *(*int32)(unsafe.Add(mBase, uint32(v7585)+60))
	if v7586 == int32(0) {
		v7625 = v7580
		goto L2052
	} else {
		goto L2054
	}
L2054:
	;
	v7589 = v7579 + v7586
	v7590 = *(*int32)(unsafe.Add(mBase, uint32(v7585)+188))
	if base.Ui32(v7590) <= base.Ui32(v7589) {
		goto L2056
	} else {
		goto L2057
	}
L2055:
	;
	v7600 = *(*int32)(unsafe.Add(mBase, uint32(v7585)+196))
	if base.B2i32(v7599 == int32(0))|base.B2i32(base.Ui32(v7589) <= base.Ui32(v7599)) != 0 {
		v7625 = v7600
		goto L2052
	} else {
		goto L2059
	}
L2056:
	;
	v7592 = *(*int32)(unsafe.Add(mBase, uint32(v7585)+192))
	v7599 = v7592
	goto L2055
L2057:
	;
	goto L2058
L2058:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7585)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7585)+188)) = v7586
	v7597 = F_strchr(m, v7586, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7585)+192)) = v7597
	v7599 = v7597
	goto L2055
L2059:
	;
	v7605 = v7599
	v7608 = v7600
	goto L2060
L2060:
	;
	v7610 = int32(1)
	v7611 = v7608 + v7610
	*(*int32)(unsafe.Add(mBase, uint32(v7585)+196)) = v7611
	v7614 = v7605 + v7610
	*(*int32)(unsafe.Add(mBase, uint32(v7585)+188)) = v7614
	v7617 = F_strchr(m, v7614, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7585)+192)) = v7617
	if v7617 == int32(0) {
		v7625 = v7611
		goto L2052
	} else {
		goto L2062
	}
L2061:
	;
	v7625 = v7611
	goto L2052
L2062:
	;
	if base.Ui32(v7617) < base.Ui32(v7589) {
		v7605 = v7617
		v7608 = v7611
		goto L2060
	} else {
		goto L2063
	}
L2063:
	;
	goto L2061
L2064:
	;
	v7644 = int32(2)
	v7649 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7650 = m.ExcPending
	if v7650 != 0 {
		goto L46
	} else {
		goto L2069
	}
L2065:
	;
	goto L2066
L2066:
	;
	v7876 = F_read_cursor_args(m, v7640, int32(59), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7877 = m.ExcPending
	if v7877 != 0 {
		goto L46
	} else {
		goto L2121
	}
L2067:
	;
	if v7759 != int32(321) {
		goto L15
	} else {
		goto L2105
	}
L2068:
	;
	v7750 = *(*int32)(unsafe.Add(mBase, uint32(v7573)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+16)) = v7750 | v7749
	v7757 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7758 = m.ExcPending
	if v7758 != 0 {
		goto L46
	} else {
		goto L2104
	}
L2069:
	;
	if v7649 == int32(369) {
		v7749 = v7644
		goto L2068
	} else {
		goto L2070
	}
L2070:
	;
	if v7649 != int32(342) {
		goto L2073
	} else {
		goto L2074
	}
L2071:
	;
	v7721 = int32(_a_F_plpgsql_yyparse_100)
	v7724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7720))))
	v7727 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[60])))
	if base.B2i32(v7724 == int32(0))|base.B2i32(v7724 != v7727) != 0 {
		v7745 = v7724
		v7746 = v7727
		goto L2097
	} else {
		goto L2098
	}
L2072:
	;
	v7717 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v7717 == int32(0) {
		goto L15
	} else {
		goto L2095
	}
L2073:
	;
	if v7649 != int32(277) {
		v7759 = v7649
		goto L2067
	} else {
		goto L2076
	}
L2074:
	;
	goto L2075
L2075:
	;
	v7671 = int32(4)
	v7676 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7677 = m.ExcPending
	if v7677 != 0 {
		goto L46
	} else {
		goto L2082
	}
L2076:
	;
	v7657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v7657&int32(1) != 0 {
		goto L15
	} else {
		goto L2077
	}
L2077:
	;
	v7660 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v7660 == int32(0) {
		goto L15
	} else {
		goto L2078
	}
L2078:
	;
	v7663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7660))))
	if v7663 != int32(110) {
		v7720 = v7660
		goto L2071
	} else {
		goto L2079
	}
L2079:
	;
	v7666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7660)+1)))
	if v7666 != int32(111) {
		goto L2072
	} else {
		goto L2080
	}
L2080:
	;
	v7669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7660)+2)))
	if v7669 != 0 {
		goto L2072
	} else {
		goto L2081
	}
L2081:
	;
	goto L2075
L2082:
	;
	if v7676 == int32(369) {
		v7749 = v7671
		goto L2068
	} else {
		goto L2083
	}
L2083:
	;
	if v7676 != int32(277) {
		v7759 = v7676
		goto L2067
	} else {
		goto L2084
	}
L2084:
	;
	v7682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v7682&int32(1) != 0 {
		goto L15
	} else {
		goto L2085
	}
L2085:
	;
	v7685 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v7685 == int32(0) {
		goto L15
	} else {
		goto L2086
	}
L2086:
	;
	v7688 = int32(_a_F_plpgsql_yyparse_100)
	v7691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7685))))
	v7694 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[60])))
	if base.B2i32(v7691 == int32(0))|base.B2i32(v7691 != v7694) != 0 {
		v7712 = v7691
		v7713 = v7694
		goto L2088
	} else {
		goto L2089
	}
L2087:
	;
	if v7712-v7713 == int32(0) {
		v7749 = v7671
		goto L2068
	} else {
		goto L2094
	}
L2088:
	;
	goto L2087
L2089:
	;
	v7697 = v7685
	v7698 = v7688
	goto L2090
L2090:
	;
	v7701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7698)+1)))
	v7702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7697)+1)))
	if v7702 == int32(0) {
		v7712 = v7702
		v7713 = v7701
		goto L2088
	} else {
		goto L2092
	}
L2091:
	;
	v7712 = v7702
	v7713 = v7701
	goto L2088
L2092:
	;
	v7705 = int32(1)
	if v7702 == v7701 {
		v7697 = v7697 + v7705
		v7698 = v7698 + v7705
		goto L2090
	} else {
		goto L2093
	}
L2093:
	;
	goto L2091
L2094:
	;
	goto L15
L2095:
	;
	v7720 = v7717
	goto L2071
L2096:
	;
	if v7745-v7746 != 0 {
		goto L15
	} else {
		goto L2103
	}
L2097:
	;
	goto L2096
L2098:
	;
	v7730 = v7720
	v7731 = v7721
	goto L2099
L2099:
	;
	v7734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7731)+1)))
	v7735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7730)+1)))
	if v7735 == int32(0) {
		v7745 = v7735
		v7746 = v7734
		goto L2097
	} else {
		goto L2101
	}
L2100:
	;
	v7745 = v7735
	v7746 = v7734
	goto L2097
L2101:
	;
	v7738 = int32(1)
	if v7735 == v7734 {
		v7730 = v7730 + v7738
		v7731 = v7731 + v7738
		goto L2099
	} else {
		goto L2102
	}
L2102:
	;
	goto L2100
L2103:
	;
	v7749 = v7644
	goto L2068
L2104:
	;
	v7759 = v7757
	goto L2067
L2105:
	;
	v7764 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v7766 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v7767 = F_plpgsql_yylex(m, v7764, v7766, l1)
	mBase = m.M
	v7768 = m.ExcPending
	if v7768 != 0 {
		goto L46
	} else {
		goto L2106
	}
L2106:
	;
	if v7767 == int32(317) {
		goto L2107
	} else {
		goto L2108
	}
L2107:
	;
	v7773 = int32(0)
	v7776 = int32(1)
	v7781 = F_read_sql_construct(m, int32(381), int32(59), v7773, int32(_a_F_plpgsql_yyparse_101), int32(2), v7776, v7776, v7773, v28+int32(256), v7764, v7766, l1)
	mBase = m.M
	v7782 = m.ExcPending
	if v7782 != 0 {
		goto L46
	} else {
		goto L2110
	}
L2108:
	;
	goto L2109
L2109:
	;
	v7862 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v7864 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, v7767, v7862, v7864, l1)
	mBase = m.M
	v7866 = m.ExcPending
	if v7866 != 0 {
		goto L46
	} else {
		goto L2119
	}
L2110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+28)) = v7781
	v7784 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7784 == int32(381) {
		goto L2111
	} else {
		goto L2112
	}
L2111:
	;
	goto L2114
L2112:
	;
	goto L2113
L2113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7573
	v9708 = v233
	goto L12
L2114:
	;
	v7814 = int32(0)
	v7817 = int32(1)
	v7826 = F_read_sql_construct(m, int32(44), int32(59), v7814, int32(_a_F_plpgsql_yyparse_66), int32(2), v7817, v7817, v7814, v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7827 = m.ExcPending
	if v7827 != 0 {
		goto L46
	} else {
		goto L2116
	}
L2115:
	;
	goto L2113
L2116:
	;
	v7828 = *(*int32)(unsafe.Add(mBase, uint32(v7573)+32))
	v7829 = F_lappend(m, v7828, v7826)
	mBase = m.M
	v7830 = m.ExcPending
	if v7830 != 0 {
		goto L46
	} else {
		goto L2117
	}
L2117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+32)) = v7829
	v7832 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7832 == int32(44) {
		goto L2114
	} else {
		goto L2118
	}
L2118:
	;
	goto L2115
L2119:
	;
	v7867 = F_read_sql_stmt(m, v7862, v7864, l1)
	mBase = m.M
	v7868 = m.ExcPending
	if v7868 != 0 {
		goto L46
	} else {
		goto L2120
	}
L2120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+24)) = v7867
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7573
	v9708 = v233
	goto L12
L2121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7573)+20)) = v7876
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7573
	v9708 = v233
	goto L12
L2122:
	;
	v7892 = F_plpgsql_yylex(m, v7887, v7889, l1)
	mBase = m.M
	v7893 = m.ExcPending
	if v7893 != 0 {
		goto L46
	} else {
		goto L2123
	}
L2123:
	;
	if v7892 != int32(59) {
		goto L3
	} else {
		goto L2124
	}
L2124:
	;
	v7896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7882)+33)))
	if v7896 == int32(1) {
		goto L20
	} else {
		goto L2125
	}
L2125:
	;
	v7901 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v7902 = int32(0)
	if v7901 < v7902 {
		v7947 = v7902
		goto L2127
	} else {
		goto L2128
	}
L2126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7882)+4)) = v7947
	v7950 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	*(*int32)(unsafe.Add(mBase, uint32(v7882)+12)) = v7950
	v7954 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v7955 = *(*int32)(unsafe.Add(mBase, uint32(v7954)+4))
	v7956 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7882)+32)) = uint8(v7956)
	*(*int32)(unsafe.Add(mBase, uint32(v7882)+16)) = v7955
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7882
	v9708 = v233
	goto L12
L2127:
	;
	goto L2126
L2128:
	;
	v7907 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7908 = *(*int32)(unsafe.Add(mBase, uint32(v7907)+60))
	if v7908 == int32(0) {
		v7947 = v7902
		goto L2127
	} else {
		goto L2129
	}
L2129:
	;
	v7911 = v7901 + v7908
	v7912 = *(*int32)(unsafe.Add(mBase, uint32(v7907)+188))
	if base.Ui32(v7912) <= base.Ui32(v7911) {
		goto L2131
	} else {
		goto L2132
	}
L2130:
	;
	v7922 = *(*int32)(unsafe.Add(mBase, uint32(v7907)+196))
	if base.B2i32(v7921 == int32(0))|base.B2i32(base.Ui32(v7911) <= base.Ui32(v7921)) != 0 {
		v7947 = v7922
		goto L2127
	} else {
		goto L2134
	}
L2131:
	;
	v7914 = *(*int32)(unsafe.Add(mBase, uint32(v7907)+192))
	v7921 = v7914
	goto L2130
L2132:
	;
	goto L2133
L2133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7907)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7907)+188)) = v7908
	v7919 = F_strchr(m, v7908, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7907)+192)) = v7919
	v7921 = v7919
	goto L2130
L2134:
	;
	v7927 = v7921
	v7930 = v7922
	goto L2135
L2135:
	;
	v7932 = int32(1)
	v7933 = v7930 + v7932
	*(*int32)(unsafe.Add(mBase, uint32(v7907)+196)) = v7933
	v7936 = v7927 + v7932
	*(*int32)(unsafe.Add(mBase, uint32(v7907)+188)) = v7936
	v7939 = F_strchr(m, v7936, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7907)+192)) = v7939
	if v7939 == int32(0) {
		v7947 = v7933
		goto L2127
	} else {
		goto L2137
	}
L2136:
	;
	v7947 = v7933
	goto L2127
L2137:
	;
	if base.Ui32(v7939) < base.Ui32(v7911) {
		v7927 = v7939
		v7930 = v7933
		goto L2135
	} else {
		goto L2138
	}
L2138:
	;
	goto L2136
L2139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7962)+4)) = v8011
	v8016 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8017 = *(*int32)(unsafe.Add(mBase, uint32(v8016)+4))
	v8018 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7962)+32)) = uint8(v8018)
	*(*int32)(unsafe.Add(mBase, uint32(v7962)+16)) = v8017
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7962
	v9708 = v233
	goto L12
L2140:
	;
	goto L2139
L2141:
	;
	v7971 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7972 = *(*int32)(unsafe.Add(mBase, uint32(v7971)+60))
	if v7972 == int32(0) {
		v8011 = v7966
		goto L2140
	} else {
		goto L2142
	}
L2142:
	;
	v7975 = v7965 + v7972
	v7976 = *(*int32)(unsafe.Add(mBase, uint32(v7971)+188))
	if base.Ui32(v7976) <= base.Ui32(v7975) {
		goto L2144
	} else {
		goto L2145
	}
L2143:
	;
	v7986 = *(*int32)(unsafe.Add(mBase, uint32(v7971)+196))
	if base.B2i32(v7985 == int32(0))|base.B2i32(base.Ui32(v7975) <= base.Ui32(v7985)) != 0 {
		v8011 = v7986
		goto L2140
	} else {
		goto L2147
	}
L2144:
	;
	v7978 = *(*int32)(unsafe.Add(mBase, uint32(v7971)+192))
	v7985 = v7978
	goto L2143
L2145:
	;
	goto L2146
L2146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7971)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7971)+188)) = v7972
	v7983 = F_strchr(m, v7972, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7971)+192)) = v7983
	v7985 = v7983
	goto L2143
L2147:
	;
	v7991 = v7985
	v7994 = v7986
	goto L2148
L2148:
	;
	v7996 = int32(1)
	v7997 = v7994 + v7996
	*(*int32)(unsafe.Add(mBase, uint32(v7971)+196)) = v7997
	v8000 = v7991 + v7996
	*(*int32)(unsafe.Add(mBase, uint32(v7971)+188)) = v8000
	v8003 = F_strchr(m, v8000, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7971)+192)) = v8003
	if v8003 == int32(0) {
		v8011 = v7997
		goto L2140
	} else {
		goto L2150
	}
L2149:
	;
	v8011 = v7997
	goto L2140
L2150:
	;
	if base.Ui32(v8003) < base.Ui32(v7975) {
		v7991 = v8003
		v7994 = v7997
		goto L2148
	} else {
		goto L2151
	}
L2151:
	;
	goto L2149
L2152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8033))) = int32(21)
	v8038 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v8039 = *(*int32)(unsafe.Add(mBase, uint32(v8038)+520))
	v8041 = v8039 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8038)+520)) = v8041
	v8043 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8033)+33)) = uint8(v8043)
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+28)) = v8043
	*(*int64)(unsafe.Add(mBase, uint32(v8033)+20)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+8)) = v8041
	v8050 = F_plpgsql_yylex(m, v8023, v8025, l1)
	mBase = m.M
	v8051 = m.ExcPending
	if v8051 != 0 {
		goto L46
	} else {
		goto L2165
	}
L2153:
	;
	v8384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8028)+15)))
	if v8384 != int32(1) {
		goto L2272
	} else {
		goto L2273
	}
L2154:
	;
	v8288 = int32(1)
	if v8050 == int32(282) {
		v8302 = v8288
		goto L2230
	} else {
		goto L2231
	}
L2155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+20)) = int32(3)
	v8272 = int32(0)
	v8275 = int32(1)
	v8279 = F_read_sql_construct(m, int32(324), int32(329), v8272, int32(_a_F_plpgsql_yyparse_102), int32(2), v8275, v8275, v8272, v8272, v8023, v8025, l1)
	mBase = m.M
	v8280 = m.ExcPending
	if v8280 != 0 {
		goto L46
	} else {
		goto L2228
	}
L2156:
	;
	v8237 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8237 == int32(0) {
		goto L2154
	} else {
		goto L2219
	}
L2157:
	;
	v8221 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+20)) = v8221
	v8225 = int32(0)
	v8228 = int32(1)
	v8232 = F_read_sql_construct(m, int32(324), int32(329), v8225, int32(_a_F_plpgsql_yyparse_102), v8221, v8228, v8228, v8225, v8225, v8023, v8025, l1)
	mBase = m.M
	v8233 = m.ExcPending
	if v8233 != 0 {
		goto L46
	} else {
		goto L2218
	}
L2158:
	;
	v8190 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8190 == int32(0) {
		goto L2156
	} else {
		goto L2209
	}
L2159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8033)+20)) = int64(-4294967294)
	goto L2153
L2160:
	;
	v8157 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8157 == int32(0) {
		goto L2158
	} else {
		goto L2200
	}
L2161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+20)) = int32(2)
	goto L2153
L2162:
	;
	v8124 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8124 == int32(0) {
		goto L2160
	} else {
		goto L2191
	}
L2163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+20)) = int32(1)
	goto L2153
L2164:
	;
	switch v8050 - int32(277) {
	case 0:
		goto L2166
	case 1, 2:
		goto L2154
	case 3:
		goto L2157
	default:
		goto L2167
	}
L2165:
	;
	switch v8050 - int32(320) {
	case 0:
		goto L2161
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 38, 39:
		goto L2154
	case 14:
		goto L2159
	case 21:
		goto L2153
	case 37:
		goto L2163
	case 40:
		goto L2155
	default:
		goto L2164
	}
L2166:
	;
	v8060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8060 != 0 {
		goto L2154
	} else {
		goto L2170
	}
L2167:
	;
	if v8050 != 0 {
		goto L2154
	} else {
		goto L2168
	}
L2168:
	;
	F_plpgsql_yyerror(m, v8025, int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v8059 = m.ExcPending
	if v8059 != 0 {
		goto L46
	} else {
		goto L2169
	}
L2169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2170:
	;
	v8061 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8061 != 0 {
		goto L2171
	} else {
		goto L2172
	}
L2171:
	;
	v8062 = int32(_a_F_plpgsql_yyparse_61)
	v8065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8061))))
	v8068 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[42])))
	if base.B2i32(v8065 == int32(0))|base.B2i32(v8065 != v8068) != 0 {
		v8086 = v8065
		v8087 = v8068
		goto L2175
	} else {
		goto L2176
	}
L2172:
	;
	goto L2173
L2173:
	;
	v8091 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8091 == int32(0) {
		goto L2162
	} else {
		goto L2182
	}
L2174:
	;
	if v8086-v8087 == int32(0) {
		goto L2153
	} else {
		goto L2181
	}
L2175:
	;
	goto L2174
L2176:
	;
	v8071 = v8061
	v8072 = v8062
	goto L2177
L2177:
	;
	v8075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8072)+1)))
	v8076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8071)+1)))
	if v8076 == int32(0) {
		v8086 = v8076
		v8087 = v8075
		goto L2175
	} else {
		goto L2179
	}
L2178:
	;
	v8086 = v8076
	v8087 = v8075
	goto L2175
L2179:
	;
	v8079 = int32(1)
	if v8076 == v8075 {
		v8071 = v8071 + v8079
		v8072 = v8072 + v8079
		goto L2177
	} else {
		goto L2180
	}
L2180:
	;
	goto L2178
L2181:
	;
	goto L2173
L2182:
	;
	v8094 = int32(_a_F_plpgsql_yyparse_103)
	v8097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8091))))
	v8100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[61])))
	if base.B2i32(v8097 == int32(0))|base.B2i32(v8097 != v8100) != 0 {
		v8118 = v8097
		v8119 = v8100
		goto L2184
	} else {
		goto L2185
	}
L2183:
	;
	if v8118-v8119 != 0 {
		goto L2162
	} else {
		goto L2190
	}
L2184:
	;
	goto L2183
L2185:
	;
	v8103 = v8091
	v8104 = v8094
	goto L2186
L2186:
	;
	v8107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8104)+1)))
	v8108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8103)+1)))
	if v8108 == int32(0) {
		v8118 = v8108
		v8119 = v8107
		goto L2184
	} else {
		goto L2188
	}
L2187:
	;
	v8118 = v8108
	v8119 = v8107
	goto L2184
L2188:
	;
	v8111 = int32(1)
	if v8108 == v8107 {
		v8103 = v8103 + v8111
		v8104 = v8104 + v8111
		goto L2186
	} else {
		goto L2189
	}
L2189:
	;
	goto L2187
L2190:
	;
	goto L2163
L2191:
	;
	v8127 = int32(_a_F_plpgsql_yyparse_104)
	v8130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8124))))
	v8133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[62])))
	if base.B2i32(v8130 == int32(0))|base.B2i32(v8130 != v8133) != 0 {
		v8151 = v8130
		v8152 = v8133
		goto L2193
	} else {
		goto L2194
	}
L2192:
	;
	if v8151-v8152 != 0 {
		goto L2160
	} else {
		goto L2199
	}
L2193:
	;
	goto L2192
L2194:
	;
	v8136 = v8124
	v8137 = v8127
	goto L2195
L2195:
	;
	v8140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8137)+1)))
	v8141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8136)+1)))
	if v8141 == int32(0) {
		v8151 = v8141
		v8152 = v8140
		goto L2193
	} else {
		goto L2197
	}
L2196:
	;
	v8151 = v8141
	v8152 = v8140
	goto L2193
L2197:
	;
	v8144 = int32(1)
	if v8141 == v8140 {
		v8136 = v8136 + v8144
		v8137 = v8137 + v8144
		goto L2195
	} else {
		goto L2198
	}
L2198:
	;
	goto L2196
L2199:
	;
	goto L2161
L2200:
	;
	v8160 = int32(_a_F_plpgsql_yyparse_105)
	v8163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8157))))
	v8166 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[63])))
	if base.B2i32(v8163 == int32(0))|base.B2i32(v8163 != v8166) != 0 {
		v8184 = v8163
		v8185 = v8166
		goto L2202
	} else {
		goto L2203
	}
L2201:
	;
	if v8184-v8185 != 0 {
		goto L2158
	} else {
		goto L2208
	}
L2202:
	;
	goto L2201
L2203:
	;
	v8169 = v8157
	v8170 = v8160
	goto L2204
L2204:
	;
	v8173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8170)+1)))
	v8174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8169)+1)))
	if v8174 == int32(0) {
		v8184 = v8174
		v8185 = v8173
		goto L2202
	} else {
		goto L2206
	}
L2205:
	;
	v8184 = v8174
	v8185 = v8173
	goto L2202
L2206:
	;
	v8177 = int32(1)
	if v8174 == v8173 {
		v8169 = v8169 + v8177
		v8170 = v8170 + v8177
		goto L2204
	} else {
		goto L2207
	}
L2207:
	;
	goto L2205
L2208:
	;
	goto L2159
L2209:
	;
	v8193 = int32(_a_F_plpgsql_yyparse_106)
	v8196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8190))))
	v8199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[64])))
	if base.B2i32(v8196 == int32(0))|base.B2i32(v8196 != v8199) != 0 {
		v8217 = v8196
		v8218 = v8199
		goto L2211
	} else {
		goto L2212
	}
L2210:
	;
	if v8217-v8218 != 0 {
		goto L2156
	} else {
		goto L2217
	}
L2211:
	;
	goto L2210
L2212:
	;
	v8202 = v8190
	v8203 = v8193
	goto L2213
L2213:
	;
	v8206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8203)+1)))
	v8207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8202)+1)))
	if v8207 == int32(0) {
		v8217 = v8207
		v8218 = v8206
		goto L2211
	} else {
		goto L2215
	}
L2214:
	;
	v8217 = v8207
	v8218 = v8206
	goto L2211
L2215:
	;
	v8210 = int32(1)
	if v8207 == v8206 {
		v8202 = v8202 + v8210
		v8203 = v8203 + v8210
		goto L2213
	} else {
		goto L2216
	}
L2216:
	;
	goto L2214
L2217:
	;
	goto L2157
L2218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+28)) = v8232
	v8235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8028)+15)) = uint8(v8235)
	goto L2153
L2219:
	;
	v8240 = int32(_a_F_plpgsql_yyparse_107)
	v8243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8237))))
	v8246 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[65])))
	if base.B2i32(v8243 == int32(0))|base.B2i32(v8243 != v8246) != 0 {
		v8264 = v8243
		v8265 = v8246
		goto L2221
	} else {
		goto L2222
	}
L2220:
	;
	if v8264-v8265 != 0 {
		goto L2154
	} else {
		goto L2227
	}
L2221:
	;
	goto L2220
L2222:
	;
	v8249 = v8237
	v8250 = v8240
	goto L2223
L2223:
	;
	v8253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8250)+1)))
	v8254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8249)+1)))
	if v8254 == int32(0) {
		v8264 = v8254
		v8265 = v8253
		goto L2221
	} else {
		goto L2225
	}
L2224:
	;
	v8264 = v8254
	v8265 = v8253
	goto L2221
L2225:
	;
	v8257 = int32(1)
	if v8254 == v8253 {
		v8249 = v8249 + v8257
		v8250 = v8250 + v8257
		goto L2223
	} else {
		goto L2226
	}
L2226:
	;
	goto L2224
L2227:
	;
	goto L2155
L2228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+28)) = v8279
	v8282 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8028)+15)) = uint8(v8282)
	goto L2153
L2229:
	;
	if v8302 != 0 {
		goto L2237
	} else {
		goto L2238
	}
L2230:
	;
	goto L2229
L2231:
	;
	if v8050 != int32(277) {
		goto L2232
	} else {
		goto L2233
	}
L2232:
	;
	v8302 = int32(0)
	goto L2230
L2233:
	;
	v8292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8292 != 0 {
		goto L2232
	} else {
		goto L2234
	}
L2234:
	;
	v8293 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8293 == int32(0) {
		goto L2232
	} else {
		goto L2235
	}
L2235:
	;
	v8296 = F_strcmp(m, v8293, int32(_a_F_plpgsql_yyparse_108))
	mBase = m.M
	if v8296 == int32(0) {
		v8302 = v8288
		goto L2230
	} else {
		goto L2236
	}
L2236:
	;
	goto L2232
L2237:
	;
	v8303 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8033)+33)) = uint8(v8303)
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+24)) = int32(2147483647)
	goto L2153
L2238:
	;
	goto L2239
L2239:
	;
	v8310 = int32(1)
	if v8050 == int32(323) {
		v8324 = v8310
		goto L2241
	} else {
		goto L2242
	}
L2240:
	;
	if v8324 != 0 {
		goto L2248
	} else {
		goto L2249
	}
L2241:
	;
	goto L2240
L2242:
	;
	if v8050 != int32(277) {
		goto L2243
	} else {
		goto L2244
	}
L2243:
	;
	v8324 = int32(0)
	goto L2241
L2244:
	;
	v8314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8314 != 0 {
		goto L2243
	} else {
		goto L2245
	}
L2245:
	;
	v8315 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8315 == int32(0) {
		goto L2243
	} else {
		goto L2246
	}
L2246:
	;
	v8318 = F_strcmp(m, v8315, int32(_a_F_plpgsql_yyparse_109))
	mBase = m.M
	if v8318 == int32(0) {
		v8324 = v8310
		goto L2241
	} else {
		goto L2247
	}
L2247:
	;
	goto L2243
L2248:
	;
	F_complete_direction(m, v8033, v8028+int32(15), v8023, v8025, l1)
	mBase = m.M
	v8328 = m.ExcPending
	if v8328 != 0 {
		goto L46
	} else {
		goto L2251
	}
L2249:
	;
	goto L2250
L2250:
	;
	v8332 = int32(1)
	if v8050 == int32(286) {
		v8346 = v8332
		goto L2253
	} else {
		goto L2254
	}
L2251:
	;
	goto L2153
L2252:
	;
	if v8346 != 0 {
		goto L2260
	} else {
		goto L2261
	}
L2253:
	;
	goto L2252
L2254:
	;
	if v8050 != int32(277) {
		goto L2255
	} else {
		goto L2256
	}
L2255:
	;
	v8346 = int32(0)
	goto L2253
L2256:
	;
	v8336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8336 != 0 {
		goto L2255
	} else {
		goto L2257
	}
L2257:
	;
	v8337 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[17])))
	if v8337 == int32(0) {
		goto L2255
	} else {
		goto L2258
	}
L2258:
	;
	v8340 = F_strcmp(m, v8337, int32(_a_F_plpgsql_yyparse_110))
	mBase = m.M
	if v8340 == int32(0) {
		v8346 = v8332
		goto L2253
	} else {
		goto L2259
	}
L2259:
	;
	goto L2255
L2260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+20)) = int32(1)
	F_complete_direction(m, v8033, v8028+int32(15), v8023, v8025, l1)
	mBase = m.M
	v8352 = m.ExcPending
	if v8352 != 0 {
		goto L46
	} else {
		goto L2263
	}
L2261:
	;
	goto L2262
L2262:
	;
	switch v8050 - int32(324) {
	case 0, 5:
		goto L2267
	case 1, 2, 3, 4:
		goto L2265
	default:
		goto L2266
	}
L2263:
	;
	goto L2153
L2264:
	;
	F_plpgsql_push_back_token(m, int32(277), v8023, v8025, l1)
	mBase = m.M
	v8379 = m.ExcPending
	if v8379 != 0 {
		goto L46
	} else {
		goto L2271
	}
L2265:
	;
	F_plpgsql_push_back_token(m, v8050, v8023, v8025, l1)
	mBase = m.M
	v8360 = m.ExcPending
	if v8360 != 0 {
		goto L46
	} else {
		goto L2269
	}
L2266:
	;
	if v8050 == int32(277) {
		goto L2264
	} else {
		goto L2268
	}
L2267:
	;
	v8355 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8028)+15)) = uint8(v8355)
	goto L2153
L2268:
	;
	goto L2265
L2269:
	;
	v8363 = int32(0)
	v8366 = int32(1)
	v8370 = F_read_sql_construct(m, int32(324), int32(329), v8363, int32(_a_F_plpgsql_yyparse_102), int32(2), v8366, v8366, v8363, v8363, v8023, v8025, l1)
	mBase = m.M
	v8371 = m.ExcPending
	if v8371 != 0 {
		goto L46
	} else {
		goto L2270
	}
L2270:
	;
	v8372 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8033)+33)) = uint8(v8372)
	*(*int32)(unsafe.Add(mBase, uint32(v8033)+28)) = v8370
	v8375 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8028)+15)) = uint8(v8375)
	goto L2153
L2271:
	;
	v8380 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8028)+15)) = uint8(v8380)
	goto L2153
L2272:
	;
	m.G0 = v8028 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8033
	v9708 = v233
	goto L12
L2273:
	;
	v8387 = F_plpgsql_yylex(m, v8023, v8025, l1)
	mBase = m.M
	v8388 = m.ExcPending
	if v8388 != 0 {
		goto L46
	} else {
		goto L2275
	}
L2274:
	;
	F_plpgsql_yyerror(m, v8025, int32(0), l1, int32(_a_F_plpgsql_yyparse_111))
	mBase = m.M
	v8394 = m.ExcPending
	if v8394 != 0 {
		goto L46
	} else {
		goto L2276
	}
L2275:
	;
	switch v8387 - int32(324) {
	case 0, 5:
		goto L2272
	default:
		goto L2274
	}
L2276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8400))) = int32(22)
	v8406 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v8407 = int32(0)
	if v8406 < v8407 {
		v8452 = v8407
		goto L2279
	} else {
		goto L2280
	}
L2278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8400)+4)) = v8452
	v8456 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v8457 = *(*int32)(unsafe.Add(mBase, uint32(v8456)+520))
	v8459 = v8457 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8456)+520)) = v8459
	*(*int32)(unsafe.Add(mBase, uint32(v8400)+8)) = v8459
	v8464 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8465 = *(*int32)(unsafe.Add(mBase, uint32(v8464)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8400)+12)) = v8465
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8400
	v9708 = v233
	goto L12
L2279:
	;
	goto L2278
L2280:
	;
	v8412 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8413 = *(*int32)(unsafe.Add(mBase, uint32(v8412)+60))
	if v8413 == int32(0) {
		v8452 = v8407
		goto L2279
	} else {
		goto L2281
	}
L2281:
	;
	v8416 = v8406 + v8413
	v8417 = *(*int32)(unsafe.Add(mBase, uint32(v8412)+188))
	if base.Ui32(v8417) <= base.Ui32(v8416) {
		goto L2283
	} else {
		goto L2284
	}
L2282:
	;
	v8427 = *(*int32)(unsafe.Add(mBase, uint32(v8412)+196))
	if base.B2i32(v8426 == int32(0))|base.B2i32(base.Ui32(v8416) <= base.Ui32(v8426)) != 0 {
		v8452 = v8427
		goto L2279
	} else {
		goto L2286
	}
L2283:
	;
	v8419 = *(*int32)(unsafe.Add(mBase, uint32(v8412)+192))
	v8426 = v8419
	goto L2282
L2284:
	;
	goto L2285
L2285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+188)) = v8413
	v8424 = F_strchr(m, v8413, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+192)) = v8424
	v8426 = v8424
	goto L2282
L2286:
	;
	v8432 = v8426
	v8435 = v8427
	goto L2287
L2287:
	;
	v8437 = int32(1)
	v8438 = v8435 + v8437
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+196)) = v8438
	v8441 = v8432 + v8437
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+188)) = v8441
	v8444 = F_strchr(m, v8441, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+192)) = v8444
	if v8444 == int32(0) {
		v8452 = v8438
		goto L2279
	} else {
		goto L2289
	}
L2288:
	;
	v8452 = v8438
	goto L2279
L2289:
	;
	if base.Ui32(v8444) < base.Ui32(v8416) {
		v8432 = v8444
		v8435 = v8438
		goto L2287
	} else {
		goto L2290
	}
L2290:
	;
	goto L2288
L2291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8471))) = int32(25)
	v8477 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v8478 = int32(0)
	if v8477 < v8478 {
		v8523 = v8478
		goto L2293
	} else {
		goto L2294
	}
L2292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8471)+4)) = v8523
	v8527 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v8528 = *(*int32)(unsafe.Add(mBase, uint32(v8527)+520))
	v8530 = v8528 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8527)+520)) = v8530
	*(*int32)(unsafe.Add(mBase, uint32(v8471)+8)) = v8530
	v8535 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8471)+12)) = uint8(base.B2i32(v8535 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8471
	v9708 = v233
	goto L12
L2293:
	;
	goto L2292
L2294:
	;
	v8483 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8484 = *(*int32)(unsafe.Add(mBase, uint32(v8483)+60))
	if v8484 == int32(0) {
		v8523 = v8478
		goto L2293
	} else {
		goto L2295
	}
L2295:
	;
	v8487 = v8477 + v8484
	v8488 = *(*int32)(unsafe.Add(mBase, uint32(v8483)+188))
	if base.Ui32(v8488) <= base.Ui32(v8487) {
		goto L2297
	} else {
		goto L2298
	}
L2296:
	;
	v8498 = *(*int32)(unsafe.Add(mBase, uint32(v8483)+196))
	if base.B2i32(v8497 == int32(0))|base.B2i32(base.Ui32(v8487) <= base.Ui32(v8497)) != 0 {
		v8523 = v8498
		goto L2293
	} else {
		goto L2300
	}
L2297:
	;
	v8490 = *(*int32)(unsafe.Add(mBase, uint32(v8483)+192))
	v8497 = v8490
	goto L2296
L2298:
	;
	goto L2299
L2299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8483)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8483)+188)) = v8484
	v8495 = F_strchr(m, v8484, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8483)+192)) = v8495
	v8497 = v8495
	goto L2296
L2300:
	;
	v8503 = v8497
	v8506 = v8498
	goto L2301
L2301:
	;
	v8508 = int32(1)
	v8509 = v8506 + v8508
	*(*int32)(unsafe.Add(mBase, uint32(v8483)+196)) = v8509
	v8512 = v8503 + v8508
	*(*int32)(unsafe.Add(mBase, uint32(v8483)+188)) = v8512
	v8515 = F_strchr(m, v8512, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8483)+192)) = v8515
	if v8515 == int32(0) {
		v8523 = v8509
		goto L2293
	} else {
		goto L2303
	}
L2302:
	;
	v8523 = v8509
	goto L2293
L2303:
	;
	if base.Ui32(v8515) < base.Ui32(v8487) {
		v8503 = v8515
		v8506 = v8509
		goto L2301
	} else {
		goto L2304
	}
L2304:
	;
	goto L2302
L2305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8541))) = int32(26)
	v8547 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v8548 = int32(0)
	if v8547 < v8548 {
		v8593 = v8548
		goto L2307
	} else {
		goto L2308
	}
L2306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8541)+4)) = v8593
	v8597 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v8598 = *(*int32)(unsafe.Add(mBase, uint32(v8597)+520))
	v8600 = v8598 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8597)+520)) = v8600
	*(*int32)(unsafe.Add(mBase, uint32(v8541)+8)) = v8600
	v8605 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8541)+12)) = uint8(base.B2i32(v8605 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8541
	v9708 = v233
	goto L12
L2307:
	;
	goto L2306
L2308:
	;
	v8553 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8554 = *(*int32)(unsafe.Add(mBase, uint32(v8553)+60))
	if v8554 == int32(0) {
		v8593 = v8548
		goto L2307
	} else {
		goto L2309
	}
L2309:
	;
	v8557 = v8547 + v8554
	v8558 = *(*int32)(unsafe.Add(mBase, uint32(v8553)+188))
	if base.Ui32(v8558) <= base.Ui32(v8557) {
		goto L2311
	} else {
		goto L2312
	}
L2310:
	;
	v8568 = *(*int32)(unsafe.Add(mBase, uint32(v8553)+196))
	if base.B2i32(v8567 == int32(0))|base.B2i32(base.Ui32(v8557) <= base.Ui32(v8567)) != 0 {
		v8593 = v8568
		goto L2307
	} else {
		goto L2314
	}
L2311:
	;
	v8560 = *(*int32)(unsafe.Add(mBase, uint32(v8553)+192))
	v8567 = v8560
	goto L2310
L2312:
	;
	goto L2313
L2313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8553)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8553)+188)) = v8554
	v8565 = F_strchr(m, v8554, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8553)+192)) = v8565
	v8567 = v8565
	goto L2310
L2314:
	;
	v8573 = v8567
	v8576 = v8568
	goto L2315
L2315:
	;
	v8578 = int32(1)
	v8579 = v8576 + v8578
	*(*int32)(unsafe.Add(mBase, uint32(v8553)+196)) = v8579
	v8582 = v8573 + v8578
	*(*int32)(unsafe.Add(mBase, uint32(v8553)+188)) = v8582
	v8585 = F_strchr(m, v8582, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8553)+192)) = v8585
	if v8585 == int32(0) {
		v8593 = v8579
		goto L2307
	} else {
		goto L2317
	}
L2316:
	;
	v8593 = v8579
	goto L2307
L2317:
	;
	if base.Ui32(v8585) < base.Ui32(v8557) {
		v8573 = v8585
		v8576 = v8579
		goto L2315
	} else {
		goto L2318
	}
L2318:
	;
	goto L2316
L2319:
	;
	v8618 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v8619 = m.ExcPending
	if v8619 != 0 {
		goto L46
	} else {
		goto L2320
	}
L2320:
	;
	if v8618 == int32(91) {
		goto L19
	} else {
		goto L2321
	}
L2321:
	;
	v8622 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8623 = *(*int32)(unsafe.Add(mBase, uint32(v8622)+24))
	v8624 = *(*int32)(unsafe.Add(mBase, uint32(v8623)+4))
	if v8624 != int32(1790) {
		goto L18
	} else {
		goto L2322
	}
L2322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8622
	v9708 = v233
	goto L12
L2323:
	;
	v8679 = F_palloc(m, int32(12))
	mBase = m.M
	v8680 = m.ExcPending
	if v8680 != 0 {
		goto L46
	} else {
		goto L2336
	}
L2324:
	;
	goto L2323
L2325:
	;
	v8636 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8637 = *(*int32)(unsafe.Add(mBase, uint32(v8636)+60))
	if v8637 == int32(0) {
		v8676 = v8631
		goto L2324
	} else {
		goto L2326
	}
L2326:
	;
	v8640 = v8630 + v8637
	v8641 = *(*int32)(unsafe.Add(mBase, uint32(v8636)+188))
	if base.Ui32(v8641) <= base.Ui32(v8640) {
		goto L2328
	} else {
		goto L2329
	}
L2327:
	;
	v8651 = *(*int32)(unsafe.Add(mBase, uint32(v8636)+196))
	if base.B2i32(v8650 == int32(0))|base.B2i32(base.Ui32(v8640) <= base.Ui32(v8650)) != 0 {
		v8676 = v8651
		goto L2324
	} else {
		goto L2331
	}
L2328:
	;
	v8643 = *(*int32)(unsafe.Add(mBase, uint32(v8636)+192))
	v8650 = v8643
	goto L2327
L2329:
	;
	goto L2330
L2330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8636)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8636)+188)) = v8637
	v8648 = F_strchr(m, v8637, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8636)+192)) = v8648
	v8650 = v8648
	goto L2327
L2331:
	;
	v8656 = v8650
	v8659 = v8651
	goto L2332
L2332:
	;
	v8661 = int32(1)
	v8662 = v8659 + v8661
	*(*int32)(unsafe.Add(mBase, uint32(v8636)+196)) = v8662
	v8665 = v8656 + v8661
	*(*int32)(unsafe.Add(mBase, uint32(v8636)+188)) = v8665
	v8668 = F_strchr(m, v8665, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8636)+192)) = v8668
	if v8668 == int32(0) {
		v8676 = v8662
		goto L2324
	} else {
		goto L2334
	}
L2333:
	;
	v8676 = v8662
	goto L2324
L2334:
	;
	if base.Ui32(v8668) < base.Ui32(v8640) {
		v8656 = v8668
		v8659 = v8662
		goto L2332
	} else {
		goto L2335
	}
L2335:
	;
	goto L2333
L2336:
	;
	v8682 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v8683 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8682)+525)) = uint8(v8683)
	v8688 = *(*int32)(unsafe.Add(mBase, uint32(v8682)+44))
	v8690 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8688, int32(0))
	mBase = m.M
	v8691 = m.ExcPending
	if v8691 != 0 {
		goto L46
	} else {
		goto L2337
	}
L2337:
	;
	v8693 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_yyparse_17), v8676, v8690, int32(1))
	mBase = m.M
	v8694 = m.ExcPending
	if v8694 != 0 {
		goto L46
	} else {
		goto L2338
	}
L2338:
	;
	v8695 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8693)+16)) = uint8(v8695)
	v8697 = *(*int32)(unsafe.Add(mBase, uint32(v8693)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8679))) = v8697
	v8703 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v8704 = *(*int32)(unsafe.Add(mBase, uint32(v8703)+44))
	v8706 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8704, int32(0))
	mBase = m.M
	v8707 = m.ExcPending
	if v8707 != 0 {
		goto L46
	} else {
		goto L2339
	}
L2339:
	;
	v8709 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_yyparse_112), v8676, v8706, int32(1))
	mBase = m.M
	v8710 = m.ExcPending
	if v8710 != 0 {
		goto L46
	} else {
		goto L2340
	}
L2340:
	;
	v8711 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8709)+16)) = uint8(v8711)
	v8713 = *(*int32)(unsafe.Add(mBase, uint32(v8709)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8679)+4)) = v8713
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8679
	v9708 = v233
	goto L12
L2341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8726
	v9708 = v233
	goto L12
L2342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8735
	v9708 = v233
	goto L12
L2343:
	;
	v8743 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v8744 = int32(0)
	if v8743 < v8744 {
		v8789 = v8744
		goto L2345
	} else {
		goto L2346
	}
L2344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8739))) = v8789
	v8794 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8739)+4)) = v8794
	v8796 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v8739)+8)) = v8796
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8739
	v9708 = v233
	goto L12
L2345:
	;
	goto L2344
L2346:
	;
	v8749 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8750 = *(*int32)(unsafe.Add(mBase, uint32(v8749)+60))
	if v8750 == int32(0) {
		v8789 = v8744
		goto L2345
	} else {
		goto L2347
	}
L2347:
	;
	v8753 = v8743 + v8750
	v8754 = *(*int32)(unsafe.Add(mBase, uint32(v8749)+188))
	if base.Ui32(v8754) <= base.Ui32(v8753) {
		goto L2349
	} else {
		goto L2350
	}
L2348:
	;
	v8764 = *(*int32)(unsafe.Add(mBase, uint32(v8749)+196))
	if base.B2i32(v8763 == int32(0))|base.B2i32(base.Ui32(v8753) <= base.Ui32(v8763)) != 0 {
		v8789 = v8764
		goto L2345
	} else {
		goto L2352
	}
L2349:
	;
	v8756 = *(*int32)(unsafe.Add(mBase, uint32(v8749)+192))
	v8763 = v8756
	goto L2348
L2350:
	;
	goto L2351
L2351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8749)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8749)+188)) = v8750
	v8761 = F_strchr(m, v8750, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8749)+192)) = v8761
	v8763 = v8761
	goto L2348
L2352:
	;
	v8769 = v8763
	v8772 = v8764
	goto L2353
L2353:
	;
	v8774 = int32(1)
	v8775 = v8772 + v8774
	*(*int32)(unsafe.Add(mBase, uint32(v8749)+196)) = v8775
	v8778 = v8769 + v8774
	*(*int32)(unsafe.Add(mBase, uint32(v8749)+188)) = v8778
	v8781 = F_strchr(m, v8778, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8749)+192)) = v8781
	if v8781 == int32(0) {
		v8789 = v8775
		goto L2345
	} else {
		goto L2355
	}
L2354:
	;
	v8789 = v8775
	goto L2345
L2355:
	;
	if base.Ui32(v8781) < base.Ui32(v8753) {
		v8769 = v8781
		v8772 = v8775
		goto L2353
	} else {
		goto L2356
	}
L2356:
	;
	goto L2354
L2357:
	;
	v8827 = *(*int32)(unsafe.Add(mBase, uint32(v8806)+8))
	if v8827 != 0 {
		v8806 = v8827
		goto L2357
	} else {
		goto L2359
	}
L2358:
	;
	v8828 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v8806)+8)) = v8828
	v8830 = *(*int32)(unsafe.Add(mBase, uint32(v8800)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8830
	v9708 = v233
	goto L12
L2359:
	;
	goto L2358
L2360:
	;
	if v8859-v8860 != 0 {
		goto L2367
	} else {
		goto L2368
	}
L2361:
	;
	goto L2360
L2362:
	;
	v8844 = v8834
	v8845 = v8835
	goto L2363
L2363:
	;
	v8848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8845)+1)))
	v8849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8844)+1)))
	if v8849 == int32(0) {
		v8859 = v8849
		v8860 = v8848
		goto L2361
	} else {
		goto L2365
	}
L2364:
	;
	v8859 = v8849
	v8860 = v8848
	goto L2361
L2365:
	;
	v8852 = int32(1)
	if v8849 == v8848 {
		v8844 = v8844 + v8852
		v8845 = v8845 + v8852
		goto L2363
	} else {
		goto L2366
	}
L2366:
	;
	goto L2364
L2367:
	;
	v8862 = F_plpgsql_parse_err_condition(m, v8834)
	mBase = m.M
	v8863 = m.ExcPending
	if v8863 != 0 {
		goto L46
	} else {
		goto L2370
	}
L2368:
	;
	goto L2369
L2369:
	;
	v8869 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_9), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v8870 = m.ExcPending
	if v8870 != 0 {
		goto L46
	} else {
		goto L2371
	}
L2370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8862
	v9708 = v233
	goto L12
L2371:
	;
	if v8869 != int32(261) {
		goto L3
	} else {
		goto L2372
	}
L2372:
	;
	v8873 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v8874 = F_strlen(m, v8873)
	mBase = m.M
	if v8874 != int32(5) {
		goto L1
	} else {
		goto L2373
	}
L2373:
	;
	v8877 = int32(_a_F_plpgsql_yyparse_81)
	v8881 = m.G0
	v8883 = v8881 - int32(32)
	v8884 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8883)+24)) = v8884
	*(*int64)(unsafe.Add(mBase, uint32(v8883)+16)) = v8884
	*(*int64)(unsafe.Add(mBase, uint32(v8883)+8)) = v8884
	*(*int64)(unsafe.Add(mBase, uint32(v8883))) = v8884
	v8892 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[49])))
	if v8892 == int32(0) {
		goto L2375
	} else {
		goto L2376
	}
L2374:
	;
	if v8960 != int32(5) {
		goto L1
	} else {
		goto L2393
	}
L2375:
	;
	v8960 = int32(0)
	goto L2374
L2376:
	;
	goto L2377
L2377:
	;
	v8896 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[50])))
	if v8896 == int32(0) {
		goto L2378
	} else {
		goto L2379
	}
L2378:
	;
	v8900 = v8873
	goto L2381
L2379:
	;
	goto L2380
L2380:
	;
	v8910 = v8877
	v8911 = v8892
	goto L2384
L2381:
	;
	v8906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8900))))
	if v8906 == v8892 {
		v8900 = v8900 + int32(1)
		goto L2381
	} else {
		goto L2383
	}
L2382:
	;
	v8960 = v8900 - v8873
	goto L2374
L2383:
	;
	goto L2382
L2384:
	;
	v8918 = v8883 + int32(base.Ui32(v8911)>>(uint(int32(3))%32))&int32(28)
	v8919 = *(*int32)(unsafe.Add(mBase, uint32(v8918)))
	v8920 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8918))) = v8919 | v8920<<(uint(v8911)%32)
	v8924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8910)+1)))
	if v8924 != 0 {
		v8910 = v8910 + v8920
		v8911 = v8924
		goto L2384
	} else {
		goto L2386
	}
L2385:
	;
	v8927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8873))))
	if v8927 == int32(0) {
		v8950 = v8873
		goto L2387
	} else {
		goto L2388
	}
L2386:
	;
	goto L2385
L2387:
	;
	v8960 = v8950 - v8873
	goto L2374
L2388:
	;
	v8931 = v8873
	v8932 = v8927
	goto L2389
L2389:
	;
	v8940 = *(*int32)(unsafe.Add(mBase, uint32(v8883+int32(base.Ui32(v8932)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v8940)>>(uint(v8932)%32))&int32(1) == int32(0) {
		v8950 = v8931
		goto L2387
	} else {
		goto L2391
	}
L2390:
	;
	v8950 = v8948
	goto L2387
L2391:
	;
	v8946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8931)+1)))
	v8948 = v8931 + int32(1)
	if v8946 != 0 {
		v8931 = v8948
		v8932 = v8946
		goto L2389
	} else {
		goto L2392
	}
L2392:
	;
	goto L2390
L2393:
	;
	v8964 = F_palloc(m, int32(12))
	mBase = m.M
	v8965 = m.ExcPending
	if v8965 != 0 {
		goto L46
	} else {
		goto L2394
	}
L2394:
	;
	v8966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8873)+4)))
	v8967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8873)+3)))
	v8968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8873)+2)))
	v8969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8873)+1)))
	v8970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8873))))
	*(*int32)(unsafe.Add(mBase, uint32(v8964)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8964)+4)) = v8873
	v8974 = int32(16)
	v8976 = int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v8964))) = (v8970+v8974)&v8976 | (v8969+v8974)&v8976<<(uint(int32(6))%32) | (v8968+v8974)&v8976<<(uint(int32(12))%32) | (v8967+v8974)&v8976<<(uint(int32(18))%32) | (v8966+v8974)&v8976<<(uint(int32(24))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8964
	v9708 = v233
	goto L12
L2395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9021
	v9708 = v233
	goto L12
L2396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9037
	v9708 = v233
	goto L12
L2397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9053
	v9708 = v233
	goto L12
L2398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L2399:
	;
	v9068 = *(*int32)(unsafe.Add(mBase, uint32(v9063)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9068
	v9708 = v233
	goto L12
L2400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9708 = v233
	goto L12
L2401:
	;
	v9082 = *(*int32)(unsafe.Add(mBase, uint32(v9077)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9082
	v9708 = v233
	goto L12
L2402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9095
	v9708 = v233
	goto L12
L2403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9098
	v9708 = v233
	goto L12
L2404:
	;
	v9104 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	if v67 != int32(3) {
		v9492 = v224
		v9494 = v9104
		goto L16
	} else {
		goto L2405
	}
L2405:
	;
	if int32(0) < v224 {
		goto L2406
	} else {
		goto L2407
	}
L2406:
	;
	v9492 = int32(-2)
	v9494 = v9104
	goto L16
L2407:
	;
	goto L2408
L2408:
	;
	if v224 != 0 {
		v9492 = v224
		v9494 = v9104
		goto L16
	} else {
		goto L2409
	}
L2409:
	;
	v9818 = int32(1)
	v9828 = v153
	goto L7
L2410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2411:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9124 = m.ExcPending
	if v9124 != 0 {
		goto L46
	} else {
		goto L2412
	}
L2412:
	;
	v9125 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v9126 = *(*int32)(unsafe.Add(mBase, uint32(v9125)+4))
	v9127 = F_format_type_be(m, v9126)
	mBase = m.M
	v9128 = m.ExcPending
	if v9128 != 0 {
		goto L46
	} else {
		goto L2413
	}
L2413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v9127
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_113), v28+int32(32))
	mBase = m.M
	v9134 = m.ExcPending
	if v9134 != 0 {
		goto L46
	} else {
		goto L2414
	}
L2414:
	;
	v9137 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v9138 = F_plpgsql_scanner_errposition(m, v9137, l1)
	mBase = m.M
	v9139 = m.ExcPending
	if v9139 != 0 {
		goto L46
	} else {
		goto L2415
	}
L2415:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(523), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9144 = m.ExcPending
	if v9144 != 0 {
		goto L46
	} else {
		goto L2416
	}
L2416:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2417:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v9151 = m.ExcPending
	if v9151 != 0 {
		goto L46
	} else {
		goto L2418
	}
L2418:
	;
	v9152 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v9152
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_114), v28+int32(16))
	mBase = m.M
	v9158 = m.ExcPending
	if v9158 != 0 {
		goto L46
	} else {
		goto L2419
	}
L2419:
	;
	v9161 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9162 = F_plpgsql_scanner_errposition(m, v9161, l1)
	mBase = m.M
	v9163 = m.ExcPending
	if v9163 != 0 {
		goto L46
	} else {
		goto L2420
	}
L2420:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(542), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9168 = m.ExcPending
	if v9168 != 0 {
		goto L46
	} else {
		goto L2421
	}
L2421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2422:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9175 = m.ExcPending
	if v9175 != 0 {
		goto L46
	} else {
		goto L2423
	}
L2423:
	;
	v9176 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v9176
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_115), v28+int32(48))
	mBase = m.M
	v9182 = m.ExcPending
	if v9182 != 0 {
		goto L46
	} else {
		goto L2424
	}
L2424:
	;
	v9183 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9184 = F_plpgsql_scanner_errposition(m, v9183, l1)
	mBase = m.M
	v9185 = m.ExcPending
	if v9185 != 0 {
		goto L46
	} else {
		goto L2425
	}
L2425:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(667), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9190 = m.ExcPending
	if v9190 != 0 {
		goto L46
	} else {
		goto L2426
	}
L2426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2427:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9197 = m.ExcPending
	if v9197 != 0 {
		goto L46
	} else {
		goto L2428
	}
L2428:
	;
	v9198 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v9198
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_115), v28-int32(-64))
	mBase = m.M
	v9204 = m.ExcPending
	if v9204 != 0 {
		goto L46
	} else {
		goto L2429
	}
L2429:
	;
	v9205 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9206 = F_plpgsql_scanner_errposition(m, v9205, l1)
	mBase = m.M
	v9207 = m.ExcPending
	if v9207 != 0 {
		goto L46
	} else {
		goto L2430
	}
L2430:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(682), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9212 = m.ExcPending
	if v9212 != 0 {
		goto L46
	} else {
		goto L2431
	}
L2431:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2432:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9220 = m.ExcPending
	if v9220 != 0 {
		goto L46
	} else {
		goto L2433
	}
L2433:
	;
	v9221 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v9222 = F_NameListToString(m, v9221)
	mBase = m.M
	v9223 = m.ExcPending
	if v9223 != 0 {
		goto L46
	} else {
		goto L2434
	}
L2434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v9222
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_115), v28+int32(80))
	mBase = m.M
	v9229 = m.ExcPending
	if v9229 != 0 {
		goto L46
	} else {
		goto L2435
	}
L2435:
	;
	v9230 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9231 = F_plpgsql_scanner_errposition(m, v9230, l1)
	mBase = m.M
	v9232 = m.ExcPending
	if v9232 != 0 {
		goto L46
	} else {
		goto L2436
	}
L2436:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(708), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9237 = m.ExcPending
	if v9237 != 0 {
		goto L46
	} else {
		goto L2437
	}
L2437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2438:
	;
	F_errmsg_internal(m, int32(_a_F_plpgsql_yyparse_116), int32(0))
	mBase = m.M
	v9246 = m.ExcPending
	if v9246 != 0 {
		goto L46
	} else {
		goto L2439
	}
L2439:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(990), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9251 = m.ExcPending
	if v9251 != 0 {
		goto L46
	} else {
		goto L2440
	}
L2440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2441:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9258 = m.ExcPending
	if v9258 != 0 {
		goto L46
	} else {
		goto L2442
	}
L2442:
	;
	v9259 = F_NameOfDatum(m, v148)
	mBase = m.M
	v9260 = m.ExcPending
	if v9260 != 0 {
		goto L46
	} else {
		goto L2443
	}
L2443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = v9259
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_117), v28+int32(176))
	mBase = m.M
	v9266 = m.ExcPending
	if v9266 != 0 {
		goto L46
	} else {
		goto L2444
	}
L2444:
	;
	v9267 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9268 = F_plpgsql_scanner_errposition(m, v9267, l1)
	mBase = m.M
	v9269 = m.ExcPending
	if v9269 != 0 {
		goto L46
	} else {
		goto L2445
	}
L2445:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1174), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9274 = m.ExcPending
	if v9274 != 0 {
		goto L46
	} else {
		goto L2446
	}
L2446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2447:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9281 = m.ExcPending
	if v9281 != 0 {
		goto L46
	} else {
		goto L2448
	}
L2448:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_118), int32(0))
	mBase = m.M
	v9285 = m.ExcPending
	if v9285 != 0 {
		goto L46
	} else {
		goto L2449
	}
L2449:
	;
	v9286 = *(*int32)(unsafe.Add(mBase, uint32(v4788)))
	v9287 = F_plpgsql_scanner_errposition(m, v9286, l1)
	mBase = m.M
	v9288 = m.ExcPending
	if v9288 != 0 {
		goto L46
	} else {
		goto L2450
	}
L2450:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1403), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9293 = m.ExcPending
	if v9293 != 0 {
		goto L46
	} else {
		goto L2451
	}
L2451:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2452:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9300 = m.ExcPending
	if v9300 != 0 {
		goto L46
	} else {
		goto L2453
	}
L2453:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_119), int32(0))
	mBase = m.M
	v9304 = m.ExcPending
	if v9304 != 0 {
		goto L46
	} else {
		goto L2454
	}
L2454:
	;
	v9307 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9308 = F_plpgsql_scanner_errposition(m, v9307, l1)
	mBase = m.M
	v9309 = m.ExcPending
	if v9309 != 0 {
		goto L46
	} else {
		goto L2455
	}
L2455:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1438), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9314 = m.ExcPending
	if v9314 != 0 {
		goto L46
	} else {
		goto L2456
	}
L2456:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2457:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9321 = m.ExcPending
	if v9321 != 0 {
		goto L46
	} else {
		goto L2458
	}
L2458:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_120), int32(0))
	mBase = m.M
	v9325 = m.ExcPending
	if v9325 != 0 {
		goto L46
	} else {
		goto L2459
	}
L2459:
	;
	v9326 = F_plpgsql_scanner_errposition(m, v4758, l1)
	mBase = m.M
	v9327 = m.ExcPending
	if v9327 != 0 {
		goto L46
	} else {
		goto L2460
	}
L2460:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1445), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9332 = m.ExcPending
	if v9332 != 0 {
		goto L46
	} else {
		goto L2461
	}
L2461:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2462:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9339 = m.ExcPending
	if v9339 != 0 {
		goto L46
	} else {
		goto L2463
	}
L2463:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_118), int32(0))
	mBase = m.M
	v9343 = m.ExcPending
	if v9343 != 0 {
		goto L46
	} else {
		goto L2464
	}
L2464:
	;
	v9344 = *(*int32)(unsafe.Add(mBase, uint32(v5018)))
	v9345 = F_plpgsql_scanner_errposition(m, v9344, l1)
	mBase = m.M
	v9346 = m.ExcPending
	if v9346 != 0 {
		goto L46
	} else {
		goto L2465
	}
L2465:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1596), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9351 = m.ExcPending
	if v9351 != 0 {
		goto L46
	} else {
		goto L2466
	}
L2466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2467:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9358 = m.ExcPending
	if v9358 != 0 {
		goto L46
	} else {
		goto L2468
	}
L2468:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_121), int32(0))
	mBase = m.M
	v9362 = m.ExcPending
	if v9362 != 0 {
		goto L46
	} else {
		goto L2469
	}
L2469:
	;
	v9365 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(20))))
	v9366 = F_plpgsql_scanner_errposition(m, v9365, l1)
	mBase = m.M
	v9367 = m.ExcPending
	if v9367 != 0 {
		goto L46
	} else {
		goto L2470
	}
L2470:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1701), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9372 = m.ExcPending
	if v9372 != 0 {
		goto L46
	} else {
		goto L2471
	}
L2471:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2472:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9379 = m.ExcPending
	if v9379 != 0 {
		goto L46
	} else {
		goto L2473
	}
L2473:
	;
	v9380 = *(*int32)(unsafe.Add(mBase, uint32(v5378)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+192)) = v9380
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_122), v28+int32(192))
	mBase = m.M
	v9386 = m.ExcPending
	if v9386 != 0 {
		goto L46
	} else {
		goto L2474
	}
L2474:
	;
	v9389 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9390 = F_plpgsql_scanner_errposition(m, v9389, l1)
	mBase = m.M
	v9391 = m.ExcPending
	if v9391 != 0 {
		goto L46
	} else {
		goto L2475
	}
L2475:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1745), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9396 = m.ExcPending
	if v9396 != 0 {
		goto L46
	} else {
		goto L2476
	}
L2476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2477:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9403 = m.ExcPending
	if v9403 != 0 {
		goto L46
	} else {
		goto L2478
	}
L2478:
	;
	v9406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5311)+12)))
	if v9406 != 0 {
		goto L2479
	} else {
		goto L2480
	}
L2479:
	;
	v9407 = int32(_a_F_plpgsql_yyparse_123)
	goto L2481
L2480:
	;
	v9407 = int32(_a_F_plpgsql_yyparse_124)
	goto L2481
L2481:
	;
	F_errmsg(m, v9407, int32(0))
	mBase = m.M
	v9410 = m.ExcPending
	if v9410 != 0 {
		goto L46
	} else {
		goto L2482
	}
L2482:
	;
	v9411 = *(*int32)(unsafe.Add(mBase, uint32(v5327)))
	v9412 = F_plpgsql_scanner_errposition(m, v9411, l1)
	mBase = m.M
	v9413 = m.ExcPending
	if v9413 != 0 {
		goto L46
	} else {
		goto L2483
	}
L2483:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1767), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9418 = m.ExcPending
	if v9418 != 0 {
		goto L46
	} else {
		goto L2484
	}
L2484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2485:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2486:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2487:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v9429 = m.ExcPending
	if v9429 != 0 {
		goto L46
	} else {
		goto L2488
	}
L2488:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_125), int32(0))
	mBase = m.M
	v9433 = m.ExcPending
	if v9433 != 0 {
		goto L46
	} else {
		goto L2489
	}
L2489:
	;
	v9436 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v9437 = F_plpgsql_scanner_errposition(m, v9436, l1)
	mBase = m.M
	v9438 = m.ExcPending
	if v9438 != 0 {
		goto L46
	} else {
		goto L2490
	}
L2490:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(2197), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9443 = m.ExcPending
	if v9443 != 0 {
		goto L46
	} else {
		goto L2491
	}
L2491:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2492:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9450 = m.ExcPending
	if v9450 != 0 {
		goto L46
	} else {
		goto L2493
	}
L2493:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_126), int32(0))
	mBase = m.M
	v9454 = m.ExcPending
	if v9454 != 0 {
		goto L46
	} else {
		goto L2494
	}
L2494:
	;
	v9455 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9456 = F_plpgsql_scanner_errposition(m, v9455, l1)
	mBase = m.M
	v9457 = m.ExcPending
	if v9457 != 0 {
		goto L46
	} else {
		goto L2495
	}
L2495:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(2294), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9462 = m.ExcPending
	if v9462 != 0 {
		goto L46
	} else {
		goto L2496
	}
L2496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2497:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9469 = m.ExcPending
	if v9469 != 0 {
		goto L46
	} else {
		goto L2498
	}
L2498:
	;
	v9470 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v9471 = *(*int32)(unsafe.Add(mBase, uint32(v9470)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+224)) = v9471
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_127), v28+int32(224))
	mBase = m.M
	v9477 = m.ExcPending
	if v9477 != 0 {
		goto L46
	} else {
		goto L2499
	}
L2499:
	;
	v9478 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9479 = F_plpgsql_scanner_errposition(m, v9478, l1)
	mBase = m.M
	v9480 = m.ExcPending
	if v9480 != 0 {
		goto L46
	} else {
		goto L2500
	}
L2500:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(2301), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9485 = m.ExcPending
	if v9485 != 0 {
		goto L46
	} else {
		goto L2501
	}
L2501:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2502:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2503:
	;
	if v9497&int32(_a_F_plpgsql_yyparse_128) == int32(_a_F_plpgsql_yyparse_14) {
		goto L2506
	} else {
		goto L2507
	}
L2504:
	;
	v9558 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v9501)+24)) = v9558
	v9560 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v9501)+16)) = v9560
	*(*int32)(unsafe.Add(mBase, uint32(v9499)+4)) = v9510
	v9771 = v9538
	v9772 = v9501 + int32(16)
	v9777 = v9492
	v9778 = v9499 + int32(4)
	v9779 = v9508
	v9785 = int32(3)
	goto L11
L2505:
	;
	goto L2504
L2506:
	;
	if v9508 == v153 {
		goto L2511
	} else {
		goto L2512
	}
L2507:
	;
	v9524 = base.I32_extend16_s(v9497)
	if v9524 < int32(-1) {
		goto L2506
	} else {
		goto L2508
	}
L2508:
	;
	v9527 = int32(1)
	v9530 = (v9524 + v9527) << (uint(v9527) % 32)
	v9533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9530)+uint32(_c_F_plpgsql_yyparse[9]))))
	if v9533 != v9527 {
		goto L2506
	} else {
		goto L2509
	}
L2509:
	;
	v9538 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9530)+uint32(_c_F_plpgsql_yyparse[10]))))
	if int32(0) < v9538 {
		goto L2505
	} else {
		goto L2510
	}
L2510:
	;
	goto L2506
L2511:
	;
	v9818 = int32(1)
	v9828 = v153
	goto L7
L2512:
	;
	v9546 = v9508 - int32(2)
	v9547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9546))))
	v9552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9547<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[5]))))
	v9555 = *(*int32)(unsafe.Add(mBase, uint32(v9499)))
	v9497 = v9552
	v9499 = v9499 - int32(4)
	v9501 = v9501 - int32(16)
	v9508 = v9546
	v9510 = v9555
	goto L2503
L2514:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2515:
	;
	v9595 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v9595 == int32(269) {
		v9616 = v9593
		v9618 = v9577
		goto L13
	} else {
		goto L2516
	}
L2516:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9601 = m.ExcPending
	if v9601 != 0 {
		goto L46
	} else {
		goto L2517
	}
L2517:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9604 = m.ExcPending
	if v9604 != 0 {
		goto L46
	} else {
		goto L2518
	}
L2518:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_129), int32(0))
	mBase = m.M
	v9608 = m.ExcPending
	if v9608 != 0 {
		goto L46
	} else {
		goto L2519
	}
L2519:
	;
	v9609 = F_plpgsql_scanner_errposition(m, v4758, l1)
	mBase = m.M
	v9610 = m.ExcPending
	if v9610 != 0 {
		goto L46
	} else {
		goto L2520
	}
L2520:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1569), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9615 = m.ExcPending
	if v9615 != 0 {
		goto L46
	} else {
		goto L2521
	}
L2521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2522:
	;
	v9626 = int32(0)
	v9632 = int32(1)
	v9638 = v28 + int32(_a_F_plpgsql_yyparse_9)
	v9640 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v9641 = F_read_sql_construct(m, int32(336), int32(288), v9626, int32(_a_F_plpgsql_yyparse_8), int32(2), v9632, v9632, v9626, v28+int32(256), v9638, v9640, l1)
	mBase = m.M
	v9642 = m.ExcPending
	if v9642 != 0 {
		goto L46
	} else {
		goto L2523
	}
L2523:
	;
	v9643 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v9643 == int32(288) {
		goto L2524
	} else {
		goto L2525
	}
L2524:
	;
	v9647 = int32(0)
	v9651 = int32(1)
	v9655 = F_read_sql_construct(m, int32(336), v9647, v9647, int32(_a_F_plpgsql_yyparse_8), int32(2), v9651, v9651, v9647, v9647, v9638, v9640, l1)
	mBase = m.M
	v9656 = m.ExcPending
	if v9656 != 0 {
		goto L46
	} else {
		goto L2527
	}
L2525:
	;
	v9657 = v9626
	goto L2526
L2526:
	;
	v9660 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(8))))
	if v9660 != 0 {
		goto L2528
	} else {
		goto L2529
	}
L2527:
	;
	v9657 = v9655
	goto L2526
L2528:
	;
	v9663 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(4))))
	if v9663 != 0 {
		goto L10
	} else {
		goto L2531
	}
L2529:
	;
	goto L2530
L2530:
	;
	v9666 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v9669 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v9672 = int32(0)
	v9674 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v9672, v9672)
	mBase = m.M
	v9675 = m.ExcPending
	if v9675 != 0 {
		goto L46
	} else {
		goto L2532
	}
L2531:
	;
	goto L2530
L2532:
	;
	v9677 = F_plpgsql_build_variable(m, v9666, v9669, v9674, int32(1))
	mBase = m.M
	v9678 = m.ExcPending
	if v9678 != 0 {
		goto L46
	} else {
		goto L2533
	}
L2533:
	;
	v9680 = F_palloc0(m, int32(40))
	mBase = m.M
	v9681 = m.ExcPending
	if v9681 != 0 {
		goto L46
	} else {
		goto L2534
	}
L2534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9680))) = int32(6)
	v9685 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14]))
	v9686 = *(*int32)(unsafe.Add(mBase, uint32(v9685)+520))
	v9688 = v9686 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9685)+520)) = v9688
	*(*int32)(unsafe.Add(mBase, uint32(v9680)+32)) = v9618
	*(*int32)(unsafe.Add(mBase, uint32(v9680)+16)) = v9677
	*(*int32)(unsafe.Add(mBase, uint32(v9680)+8)) = v9688
	*(*int32)(unsafe.Add(mBase, uint32(v9680)+28)) = v9657
	*(*int32)(unsafe.Add(mBase, uint32(v9680)+24)) = v9641
	*(*int32)(unsafe.Add(mBase, uint32(v9680)+20)) = v9616
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9680
	v9708 = v233
	goto L12
L2535:
	;
	v9765 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9744)+uint32(_c_F_plpgsql_yyparse[66]))))
	v9771 = v9765
	v9772 = v9733
	v9777 = v9708
	v9778 = v252
	v9779 = v9736
	v9785 = v67
	goto L11
L2536:
	;
	v9752 = v9748 << (uint(int32(1)) % 32)
	v9755 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9752)+uint32(_c_F_plpgsql_yyparse[9]))))
	if v9755 != v9737&int32(_a_F_plpgsql_yyparse_128) {
		goto L2535
	} else {
		goto L2537
	}
L2537:
	;
	v9761 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9752)+uint32(_c_F_plpgsql_yyparse[10]))))
	v9771 = v9761
	v9772 = v9733
	v9777 = v9708
	v9778 = v252
	v9779 = v9736
	v9785 = v67
	goto L11
L2538:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9799 = m.ExcPending
	if v9799 != 0 {
		goto L46
	} else {
		goto L2539
	}
L2539:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_130), int32(0))
	mBase = m.M
	v9803 = m.ExcPending
	if v9803 != 0 {
		goto L46
	} else {
		goto L2540
	}
L2540:
	;
	v9806 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9807 = F_plpgsql_scanner_errposition(m, v9806, l1)
	mBase = m.M
	v9808 = m.ExcPending
	if v9808 != 0 {
		goto L46
	} else {
		goto L2541
	}
L2541:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1535), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9813 = m.ExcPending
	if v9813 != 0 {
		goto L46
	} else {
		goto L2542
	}
L2542:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2543:
	;
	F_pfree(m, v9828)
	mBase = m.M
	v9843 = m.ExcPending
	if v9843 != 0 {
		goto L46
	} else {
		goto L2546
	}
L2544:
	;
	goto L2545
L2545:
	;
	m.G0 = v28 + int32(_a_F_plpgsql_yyparse_0)
	return v9818
L2546:
	;
	goto L2545
L2547:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2548:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2549:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2550:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2551:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2552:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
