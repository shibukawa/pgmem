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
	var v58 int32
	_ = v58
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
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v763 int32
	_ = v763
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v837 int32
	_ = v837
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v922 int32
	_ = v922
	var v964 int32
	_ = v964
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v992 int32
	_ = v992
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1077 int32
	_ = v1077
	var v1119 int32
	_ = v1119
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1157 int32
	_ = v1157
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1284 int32
	_ = v1284
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
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
	var v1314 int32
	_ = v1314
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1441 int32
	_ = v1441
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1649 int32
	_ = v1649
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1686 int32
	_ = v1686
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1771 int32
	_ = v1771
	var v1813 int32
	_ = v1813
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1932 int32
	_ = v1932
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1944 int32
	_ = v1944
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v2050 int32
	_ = v2050
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2068 int32
	_ = v2068
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2087 int32
	_ = v2087
	var v2096 int32
	_ = v2096
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2108 int32
	_ = v2108
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2126 int32
	_ = v2126
	var v2172 int32
	_ = v2172
	var v2214 int32
	_ = v2214
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2299 int32
	_ = v2299
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2361 int32
	_ = v2361
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2372 int32
	_ = v2372
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2397 int32
	_ = v2397
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2511 int32
	_ = v2511
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2528 int32
	_ = v2528
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2557 int32
	_ = v2557
	var v2565 int32
	_ = v2565
	var v2593 int32
	_ = v2593
	var v2599 int32
	_ = v2599
	var v2613 int32
	_ = v2613
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2658 int32
	_ = v2658
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2692 int32
	_ = v2692
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2706 int32
	_ = v2706
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2749 int32
	_ = v2749
	var v2750 int32
	_ = v2750
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2767 int32
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2775 int32
	_ = v2775
	var v2779 int32
	_ = v2779
	var v2790 int32
	_ = v2790
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2830 int32
	_ = v2830
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2849 int32
	_ = v2849
	var v2854 int32
	_ = v2854
	var v2858 int32
	_ = v2858
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
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
	var v2925 int32
	_ = v2925
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2932 int32
	_ = v2932
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
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2954 int32
	_ = v2954
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2962 int32
	_ = v2962
	var v2965 int32
	_ = v2965
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2971 int32
	_ = v2971
	var v2974 int32
	_ = v2974
	var v2982 int32
	_ = v2982
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2989 int32
	_ = v2989
	var v2994 int32
	_ = v2994
	var v2996 int32
	_ = v2996
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3012 int32
	_ = v3012
	var v3014 int32
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3032 int32
	_ = v3032
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3057 int32
	_ = v3057
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3065 int32
	_ = v3065
	var v3068 int32
	_ = v3068
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3074 int32
	_ = v3074
	var v3077 int32
	_ = v3077
	var v3085 int32
	_ = v3085
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3092 int32
	_ = v3092
	var v3097 int32
	_ = v3097
	var v3099 int32
	_ = v3099
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3104 int32
	_ = v3104
	var v3108 int32
	_ = v3108
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3130 int32
	_ = v3130
	var v3135 int32
	_ = v3135
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3148 int32
	_ = v3148
	var v3149 int32
	_ = v3149
	var v3152 int32
	_ = v3152
	var v3155 int32
	_ = v3155
	var v3163 int32
	_ = v3163
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3170 int32
	_ = v3170
	var v3175 int32
	_ = v3175
	var v3177 int32
	_ = v3177
	var v3179 int32
	_ = v3179
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3182 int32
	_ = v3182
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3191 int32
	_ = v3191
	var v3194 int32
	_ = v3194
	var v3197 int32
	_ = v3197
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3241 int32
	_ = v3241
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3253 int32
	_ = v3253
	var v3261 int32
	_ = v3261
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3268 int32
	_ = v3268
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3280 int32
	_ = v3280
	var v3282 int32
	_ = v3282
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
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3316 int32
	_ = v3316
	var v3317 int32
	_ = v3317
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3323 int32
	_ = v3323
	var v3328 int32
	_ = v3328
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3336 int32
	_ = v3336
	var v3339 int32
	_ = v3339
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3345 int32
	_ = v3345
	var v3348 int32
	_ = v3348
	var v3356 int32
	_ = v3356
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3363 int32
	_ = v3363
	var v3368 int32
	_ = v3368
	var v3372 int32
	_ = v3372
	var v3376 int32
	_ = v3376
	var v3379 int32
	_ = v3379
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3389 int32
	_ = v3389
	var v3413 int32
	_ = v3413
	var v3414 int32
	_ = v3414
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3445 int32
	_ = v3445
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3453 int32
	_ = v3453
	var v3459 int32
	_ = v3459
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3478 int32
	_ = v3478
	var v3479 int32
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3481 int32
	_ = v3481
	var v3486 int32
	_ = v3486
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3497 int32
	_ = v3497
	var v3502 int32
	_ = v3502
	var v3504 int32
	_ = v3504
	var v3532 int32
	_ = v3532
	var v3534 int32
	_ = v3534
	var v3536 int32
	_ = v3536
	var v3540 int32
	_ = v3540
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3545 int32
	_ = v3545
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3555 int32
	_ = v3555
	var v3556 int32
	_ = v3556
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3569 int32
	_ = v3569
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
	var v3585 int32
	_ = v3585
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3596 int32
	_ = v3596
	var v3603 int32
	_ = v3603
	var v3604 int32
	_ = v3604
	var v3609 int32
	_ = v3609
	var v3612 int32
	_ = v3612
	var v3615 int32
	_ = v3615
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3626 int32
	_ = v3626
	var v3633 int32
	_ = v3633
	var v3634 int32
	_ = v3634
	var v3639 int32
	_ = v3639
	var v3642 int32
	_ = v3642
	var v3645 int32
	_ = v3645
	var v3648 int32
	_ = v3648
	var v3651 int32
	_ = v3651
	var v3652 int32
	_ = v3652
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3659 int32
	_ = v3659
	var v3666 int32
	_ = v3666
	var v3667 int32
	_ = v3667
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3689 int32
	_ = v3689
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3702 int32
	_ = v3702
	var v3705 int32
	_ = v3705
	var v3708 int32
	_ = v3708
	var v3711 int32
	_ = v3711
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3718 int32
	_ = v3718
	var v3719 int32
	_ = v3719
	var v3722 int32
	_ = v3722
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3735 int32
	_ = v3735
	var v3738 int32
	_ = v3738
	var v3741 int32
	_ = v3741
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3752 int32
	_ = v3752
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3765 int32
	_ = v3765
	var v3768 int32
	_ = v3768
	var v3771 int32
	_ = v3771
	var v3774 int32
	_ = v3774
	var v3777 int32
	_ = v3777
	var v3778 int32
	_ = v3778
	var v3781 int32
	_ = v3781
	var v3782 int32
	_ = v3782
	var v3785 int32
	_ = v3785
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3798 int32
	_ = v3798
	var v3801 int32
	_ = v3801
	var v3804 int32
	_ = v3804
	var v3807 int32
	_ = v3807
	var v3808 int32
	_ = v3808
	var v3811 int32
	_ = v3811
	var v3812 int32
	_ = v3812
	var v3815 int32
	_ = v3815
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3828 int32
	_ = v3828
	var v3831 int32
	_ = v3831
	var v3834 int32
	_ = v3834
	var v3837 int32
	_ = v3837
	var v3840 int32
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3844 int32
	_ = v3844
	var v3845 int32
	_ = v3845
	var v3848 int32
	_ = v3848
	var v3855 int32
	_ = v3855
	var v3856 int32
	_ = v3856
	var v3861 int32
	_ = v3861
	var v3864 int32
	_ = v3864
	var v3867 int32
	_ = v3867
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3874 int32
	_ = v3874
	var v3875 int32
	_ = v3875
	var v3878 int32
	_ = v3878
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3891 int32
	_ = v3891
	var v3894 int32
	_ = v3894
	var v3897 int32
	_ = v3897
	var v3900 int32
	_ = v3900
	var v3903 int32
	_ = v3903
	var v3904 int32
	_ = v3904
	var v3907 int32
	_ = v3907
	var v3908 int32
	_ = v3908
	var v3911 int32
	_ = v3911
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3924 int32
	_ = v3924
	var v3927 int32
	_ = v3927
	var v3930 int32
	_ = v3930
	var v3933 int32
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3937 int32
	_ = v3937
	var v3938 int32
	_ = v3938
	var v3941 int32
	_ = v3941
	var v3948 int32
	_ = v3948
	var v3949 int32
	_ = v3949
	var v3954 int32
	_ = v3954
	var v3957 int32
	_ = v3957
	var v3960 int32
	_ = v3960
	var v3963 int32
	_ = v3963
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3970 int32
	_ = v3970
	var v3971 int32
	_ = v3971
	var v3974 int32
	_ = v3974
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v4000 int32
	_ = v4000
	var v4001 int32
	_ = v4001
	var v4004 int32
	_ = v4004
	var v4005 int32
	_ = v4005
	var v4007 int32
	_ = v4007
	var v4008 int32
	_ = v4008
	var v4011 int32
	_ = v4011
	var v4012 int32
	_ = v4012
	var v4017 int32
	_ = v4017
	var v4018 int32
	_ = v4018
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
	var v4035 int32
	_ = v4035
	var v4037 int32
	_ = v4037
	var v4038 int32
	_ = v4038
	var v4043 int32
	_ = v4043
	var v4046 int32
	_ = v4046
	var v4048 int32
	_ = v4048
	var v4049 int32
	_ = v4049
	var v4052 int32
	_ = v4052
	var v4055 int32
	_ = v4055
	var v4063 int32
	_ = v4063
	var v4067 int32
	_ = v4067
	var v4068 int32
	_ = v4068
	var v4070 int32
	_ = v4070
	var v4075 int32
	_ = v4075
	var v4079 int32
	_ = v4079
	var v4083 int32
	_ = v4083
	var v4087 int32
	_ = v4087
	var v4093 int32
	_ = v4093
	var v4094 int32
	_ = v4094
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4103 int32
	_ = v4103
	var v4104 int32
	_ = v4104
	var v4107 int32
	_ = v4107
	var v4108 int32
	_ = v4108
	var v4110 int32
	_ = v4110
	var v4115 int32
	_ = v4115
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4123 int32
	_ = v4123
	var v4126 int32
	_ = v4126
	var v4128 int32
	_ = v4128
	var v4129 int32
	_ = v4129
	var v4132 int32
	_ = v4132
	var v4135 int32
	_ = v4135
	var v4143 int32
	_ = v4143
	var v4148 int32
	_ = v4148
	var v4150 int32
	_ = v4150
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4156 int32
	_ = v4156
	var v4160 int32
	_ = v4160
	var v4164 int32
	_ = v4164
	var v4165 int32
	_ = v4165
	var v4167 int32
	_ = v4167
	var v4170 int32
	_ = v4170
	var v4173 int32
	_ = v4173
	var v4174 int32
	_ = v4174
	var v4176 int32
	_ = v4176
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4183 int32
	_ = v4183
	var v4188 int32
	_ = v4188
	var v4189 int32
	_ = v4189
	var v4192 int32
	_ = v4192
	var v4193 int32
	_ = v4193
	var v4195 int32
	_ = v4195
	var v4200 int32
	_ = v4200
	var v4202 int32
	_ = v4202
	var v4203 int32
	_ = v4203
	var v4208 int32
	_ = v4208
	var v4211 int32
	_ = v4211
	var v4213 int32
	_ = v4213
	var v4214 int32
	_ = v4214
	var v4217 int32
	_ = v4217
	var v4220 int32
	_ = v4220
	var v4228 int32
	_ = v4228
	var v4232 int32
	_ = v4232
	var v4233 int32
	_ = v4233
	var v4235 int32
	_ = v4235
	var v4237 int32
	_ = v4237
	var v4247 int32
	_ = v4247
	var v4250 int32
	_ = v4250
	var v4251 int32
	_ = v4251
	var v4253 int32
	_ = v4253
	var v4258 int32
	_ = v4258
	var v4261 int32
	_ = v4261
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4274 int32
	_ = v4274
	var v4275 int32
	_ = v4275
	var v4277 int32
	_ = v4277
	var v4278 int32
	_ = v4278
	var v4279 int32
	_ = v4279
	var v4283 int32
	_ = v4283
	var v4290 int32
	_ = v4290
	var v4311 int32
	_ = v4311
	var v4315 int32
	_ = v4315
	var v4316 int32
	_ = v4316
	var v4318 int32
	_ = v4318
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4328 int32
	_ = v4328
	var v4329 int32
	_ = v4329
	var v4331 int32
	_ = v4331
	var v4332 int32
	_ = v4332
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4337 int32
	_ = v4337
	var v4339 int32
	_ = v4339
	var v4341 int32
	_ = v4341
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4377 int32
	_ = v4377
	var v4379 int32
	_ = v4379
	var v4380 int32
	_ = v4380
	var v4381 int32
	_ = v4381
	var v4385 int32
	_ = v4385
	var v4387 int32
	_ = v4387
	var v4391 int32
	_ = v4391
	var v4395 int32
	_ = v4395
	var v4396 int32
	_ = v4396
	var v4397 int32
	_ = v4397
	var v4404 int32
	_ = v4404
	var v4408 int32
	_ = v4408
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4411 int32
	_ = v4411
	var v4413 int32
	_ = v4413
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4423 int32
	_ = v4423
	var v4424 int32
	_ = v4424
	var v4427 int32
	_ = v4427
	var v4428 int32
	_ = v4428
	var v4433 int32
	_ = v4433
	var v4434 int32
	_ = v4434
	var v4437 int32
	_ = v4437
	var v4438 int32
	_ = v4438
	var v4440 int32
	_ = v4440
	var v4445 int32
	_ = v4445
	var v4447 int32
	_ = v4447
	var v4448 int32
	_ = v4448
	var v4453 int32
	_ = v4453
	var v4456 int32
	_ = v4456
	var v4458 int32
	_ = v4458
	var v4459 int32
	_ = v4459
	var v4462 int32
	_ = v4462
	var v4465 int32
	_ = v4465
	var v4473 int32
	_ = v4473
	var v4478 int32
	_ = v4478
	var v4480 int32
	_ = v4480
	var v4485 int32
	_ = v4485
	var v4487 int32
	_ = v4487
	var v4494 int32
	_ = v4494
	var v4495 int32
	_ = v4495
	var v4498 int32
	_ = v4498
	var v4499 int32
	_ = v4499
	var v4500 int32
	_ = v4500
	var v4504 int32
	_ = v4504
	var v4505 int32
	_ = v4505
	var v4510 int32
	_ = v4510
	var v4511 int32
	_ = v4511
	var v4514 int32
	_ = v4514
	var v4515 int32
	_ = v4515
	var v4517 int32
	_ = v4517
	var v4522 int32
	_ = v4522
	var v4524 int32
	_ = v4524
	var v4525 int32
	_ = v4525
	var v4530 int32
	_ = v4530
	var v4533 int32
	_ = v4533
	var v4535 int32
	_ = v4535
	var v4536 int32
	_ = v4536
	var v4539 int32
	_ = v4539
	var v4542 int32
	_ = v4542
	var v4550 int32
	_ = v4550
	var v4554 int32
	_ = v4554
	var v4555 int32
	_ = v4555
	var v4557 int32
	_ = v4557
	var v4561 int32
	_ = v4561
	var v4562 int32
	_ = v4562
	var v4564 int32
	_ = v4564
	var v4566 int32
	_ = v4566
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4570 int32
	_ = v4570
	var v4573 int32
	_ = v4573
	var v4574 int32
	_ = v4574
	var v4575 int32
	_ = v4575
	var v4576 int32
	_ = v4576
	var v4577 int32
	_ = v4577
	var v4578 int32
	_ = v4578
	var v4580 int32
	_ = v4580
	var v4584 int32
	_ = v4584
	var v4585 int32
	_ = v4585
	var v4590 int32
	_ = v4590
	var v4591 int32
	_ = v4591
	var v4596 int32
	_ = v4596
	var v4597 int32
	_ = v4597
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4603 int32
	_ = v4603
	var v4608 int32
	_ = v4608
	var v4610 int32
	_ = v4610
	var v4611 int32
	_ = v4611
	var v4616 int32
	_ = v4616
	var v4619 int32
	_ = v4619
	var v4621 int32
	_ = v4621
	var v4622 int32
	_ = v4622
	var v4625 int32
	_ = v4625
	var v4628 int32
	_ = v4628
	var v4636 int32
	_ = v4636
	var v4640 int32
	_ = v4640
	var v4641 int32
	_ = v4641
	var v4643 int32
	_ = v4643
	var v4647 int32
	_ = v4647
	var v4648 int32
	_ = v4648
	var v4652 int32
	_ = v4652
	var v4654 int32
	_ = v4654
	var v4656 int32
	_ = v4656
	var v4657 int32
	_ = v4657
	var v4658 int32
	_ = v4658
	var v4660 int32
	_ = v4660
	var v4663 int32
	_ = v4663
	var v4664 int32
	_ = v4664
	var v4665 int32
	_ = v4665
	var v4666 int32
	_ = v4666
	var v4667 int32
	_ = v4667
	var v4668 int32
	_ = v4668
	var v4670 int32
	_ = v4670
	var v4675 int32
	_ = v4675
	var v4676 int32
	_ = v4676
	var v4679 int32
	_ = v4679
	var v4680 int32
	_ = v4680
	var v4685 int32
	_ = v4685
	var v4686 int32
	_ = v4686
	var v4689 int32
	_ = v4689
	var v4690 int32
	_ = v4690
	var v4692 int32
	_ = v4692
	var v4697 int32
	_ = v4697
	var v4699 int32
	_ = v4699
	var v4700 int32
	_ = v4700
	var v4705 int32
	_ = v4705
	var v4708 int32
	_ = v4708
	var v4710 int32
	_ = v4710
	var v4711 int32
	_ = v4711
	var v4714 int32
	_ = v4714
	var v4717 int32
	_ = v4717
	var v4725 int32
	_ = v4725
	var v4729 int32
	_ = v4729
	var v4730 int32
	_ = v4730
	var v4736 int32
	_ = v4736
	var v4738 int32
	_ = v4738
	var v4741 int32
	_ = v4741
	var v4742 int32
	_ = v4742
	var v4743 int32
	_ = v4743
	var v4745 int32
	_ = v4745
	var v4748 int32
	_ = v4748
	var v4749 int32
	_ = v4749
	var v4750 int32
	_ = v4750
	var v4751 int32
	_ = v4751
	var v4752 int32
	_ = v4752
	var v4753 int32
	_ = v4753
	var v4755 int32
	_ = v4755
	var v4761 int32
	_ = v4761
	var v4762 int32
	_ = v4762
	var v4766 int32
	_ = v4766
	var v4769 int32
	_ = v4769
	var v4772 int32
	_ = v4772
	var v4775 int32
	_ = v4775
	var v4778 int32
	_ = v4778
	var v4781 int32
	_ = v4781
	var v4784 int32
	_ = v4784
	var v4785 int32
	_ = v4785
	var v4788 int32
	_ = v4788
	var v4789 int32
	_ = v4789
	var v4792 int32
	_ = v4792
	var v4799 int32
	_ = v4799
	var v4800 int32
	_ = v4800
	var v4807 int32
	_ = v4807
	var v4810 int32
	_ = v4810
	var v4819 int32
	_ = v4819
	var v4820 int32
	_ = v4820
	var v4822 int32
	_ = v4822
	var v4823 int32
	_ = v4823
	var v4827 int32
	_ = v4827
	var v4828 int32
	_ = v4828
	var v4830 int32
	_ = v4830
	var v4833 int32
	_ = v4833
	var v4834 int32
	_ = v4834
	var v4836 int32
	_ = v4836
	var v4837 int32
	_ = v4837
	var v4839 int32
	_ = v4839
	var v4840 int32
	_ = v4840
	var v4842 int32
	_ = v4842
	var v4845 int32
	_ = v4845
	var v4850 int32
	_ = v4850
	var v4853 int32
	_ = v4853
	var v4854 int32
	_ = v4854
	var v4855 int32
	_ = v4855
	var v4856 int32
	_ = v4856
	var v4860 int32
	_ = v4860
	var v4890 int32
	_ = v4890
	var v4893 int32
	_ = v4893
	var v4902 int32
	_ = v4902
	var v4903 int32
	_ = v4903
	var v4904 int32
	_ = v4904
	var v4905 int32
	_ = v4905
	var v4906 int32
	_ = v4906
	var v4908 int32
	_ = v4908
	var v4938 int32
	_ = v4938
	var v4939 int32
	_ = v4939
	var v4940 int32
	_ = v4940
	var v4941 int32
	_ = v4941
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4947 int32
	_ = v4947
	var v4950 int32
	_ = v4950
	var v4951 int32
	_ = v4951
	var v4953 int32
	_ = v4953
	var v4956 int32
	_ = v4956
	var v4960 int32
	_ = v4960
	var v4963 int32
	_ = v4963
	var v4964 int32
	_ = v4964
	var v4972 int32
	_ = v4972
	var v4973 int32
	_ = v4973
	var v4977 int32
	_ = v4977
	var v4980 int32
	_ = v4980
	var v4984 int32
	_ = v4984
	var v4985 int32
	_ = v4985
	var v4994 int32
	_ = v4994
	var v4995 int32
	_ = v4995
	var v4996 int32
	_ = v4996
	var v4997 int32
	_ = v4997
	var v4998 int32
	_ = v4998
	var v5000 int32
	_ = v5000
	var v5001 int32
	_ = v5001
	var v5006 int32
	_ = v5006
	var v5007 int32
	_ = v5007
	var v5010 int32
	_ = v5010
	var v5011 int32
	_ = v5011
	var v5013 int32
	_ = v5013
	var v5018 int32
	_ = v5018
	var v5020 int32
	_ = v5020
	var v5021 int32
	_ = v5021
	var v5026 int32
	_ = v5026
	var v5029 int32
	_ = v5029
	var v5031 int32
	_ = v5031
	var v5032 int32
	_ = v5032
	var v5035 int32
	_ = v5035
	var v5038 int32
	_ = v5038
	var v5046 int32
	_ = v5046
	var v5049 int32
	_ = v5049
	var v5050 int32
	_ = v5050
	var v5051 int32
	_ = v5051
	var v5057 int32
	_ = v5057
	var v5063 int32
	_ = v5063
	var v5065 int32
	_ = v5065
	var v5066 int32
	_ = v5066
	var v5067 int32
	_ = v5067
	var v5069 int32
	_ = v5069
	var v5072 int32
	_ = v5072
	var v5073 int32
	_ = v5073
	var v5074 int32
	_ = v5074
	var v5075 int32
	_ = v5075
	var v5076 int32
	_ = v5076
	var v5078 int32
	_ = v5078
	var v5080 int32
	_ = v5080
	var v5081 int32
	_ = v5081
	var v5086 int32
	_ = v5086
	var v5087 int32
	_ = v5087
	var v5090 int32
	_ = v5090
	var v5091 int32
	_ = v5091
	var v5093 int32
	_ = v5093
	var v5098 int32
	_ = v5098
	var v5100 int32
	_ = v5100
	var v5101 int32
	_ = v5101
	var v5106 int32
	_ = v5106
	var v5109 int32
	_ = v5109
	var v5111 int32
	_ = v5111
	var v5112 int32
	_ = v5112
	var v5115 int32
	_ = v5115
	var v5118 int32
	_ = v5118
	var v5126 int32
	_ = v5126
	var v5132 int32
	_ = v5132
	var v5134 int32
	_ = v5134
	var v5135 int32
	_ = v5135
	var v5136 int32
	_ = v5136
	var v5138 int32
	_ = v5138
	var v5142 int32
	_ = v5142
	var v5143 int32
	_ = v5143
	var v5148 int32
	_ = v5148
	var v5149 int32
	_ = v5149
	var v5154 int32
	_ = v5154
	var v5155 int32
	_ = v5155
	var v5158 int32
	_ = v5158
	var v5159 int32
	_ = v5159
	var v5161 int32
	_ = v5161
	var v5166 int32
	_ = v5166
	var v5168 int32
	_ = v5168
	var v5169 int32
	_ = v5169
	var v5174 int32
	_ = v5174
	var v5177 int32
	_ = v5177
	var v5179 int32
	_ = v5179
	var v5180 int32
	_ = v5180
	var v5183 int32
	_ = v5183
	var v5186 int32
	_ = v5186
	var v5194 int32
	_ = v5194
	var v5198 int32
	_ = v5198
	var v5199 int32
	_ = v5199
	var v5201 int32
	_ = v5201
	var v5205 int32
	_ = v5205
	var v5206 int32
	_ = v5206
	var v5210 int32
	_ = v5210
	var v5214 int32
	_ = v5214
	var v5216 int32
	_ = v5216
	var v5219 int32
	_ = v5219
	var v5220 int32
	_ = v5220
	var v5224 int32
	_ = v5224
	var v5225 int32
	_ = v5225
	var v5228 int32
	_ = v5228
	var v5229 int32
	_ = v5229
	var v5230 int32
	_ = v5230
	var v5232 int32
	_ = v5232
	var v5235 int32
	_ = v5235
	var v5237 int32
	_ = v5237
	var v5238 int32
	_ = v5238
	var v5239 int32
	_ = v5239
	var v5240 int32
	_ = v5240
	var v5242 int32
	_ = v5242
	var v5245 int32
	_ = v5245
	var v5246 int32
	_ = v5246
	var v5247 int32
	_ = v5247
	var v5248 int32
	_ = v5248
	var v5249 int32
	_ = v5249
	var v5250 int32
	_ = v5250
	var v5252 int32
	_ = v5252
	var v5257 int32
	_ = v5257
	var v5260 int32
	_ = v5260
	var v5261 int32
	_ = v5261
	var v5265 int32
	_ = v5265
	var v5266 int32
	_ = v5266
	var v5268 int32
	_ = v5268
	var v5273 int32
	_ = v5273
	var v5276 int32
	_ = v5276
	var v5277 int32
	_ = v5277
	var v5278 int32
	_ = v5278
	var v5283 int32
	_ = v5283
	var v5284 int32
	_ = v5284
	var v5287 int32
	_ = v5287
	var v5288 int32
	_ = v5288
	var v5290 int32
	_ = v5290
	var v5295 int32
	_ = v5295
	var v5297 int32
	_ = v5297
	var v5298 int32
	_ = v5298
	var v5303 int32
	_ = v5303
	var v5306 int32
	_ = v5306
	var v5308 int32
	_ = v5308
	var v5309 int32
	_ = v5309
	var v5312 int32
	_ = v5312
	var v5315 int32
	_ = v5315
	var v5323 int32
	_ = v5323
	var v5327 int32
	_ = v5327
	var v5328 int32
	_ = v5328
	var v5330 int32
	_ = v5330
	var v5333 int32
	_ = v5333
	var v5334 int32
	_ = v5334
	var v5335 int32
	_ = v5335
	var v5336 int32
	_ = v5336
	var v5338 int32
	_ = v5338
	var v5341 int32
	_ = v5341
	var v5342 int32
	_ = v5342
	var v5348 int32
	_ = v5348
	var v5351 int32
	_ = v5351
	var v5354 int32
	_ = v5354
	var v5358 int32
	_ = v5358
	var v5361 int32
	_ = v5361
	var v5362 int32
	_ = v5362
	var v5368 int32
	_ = v5368
	var v5371 int32
	_ = v5371
	var v5372 int32
	_ = v5372
	var v5373 int32
	_ = v5373
	var v5378 int32
	_ = v5378
	var v5379 int32
	_ = v5379
	var v5380 int32
	_ = v5380
	var v5381 int32
	_ = v5381
	var v5384 int32
	_ = v5384
	var v5388 int32
	_ = v5388
	var v5393 int32
	_ = v5393
	var v5395 int32
	_ = v5395
	var v5401 int32
	_ = v5401
	var v5402 int32
	_ = v5402
	var v5407 int32
	_ = v5407
	var v5410 int32
	_ = v5410
	var v5413 int32
	_ = v5413
	var v5416 int32
	_ = v5416
	var v5419 int32
	_ = v5419
	var v5422 int32
	_ = v5422
	var v5423 int32
	_ = v5423
	var v5426 int32
	_ = v5426
	var v5427 int32
	_ = v5427
	var v5430 int32
	_ = v5430
	var v5437 int32
	_ = v5437
	var v5438 int32
	_ = v5438
	var v5441 int32
	_ = v5441
	var v5443 int32
	_ = v5443
	var v5445 int32
	_ = v5445
	var v5447 int32
	_ = v5447
	var v5448 int32
	_ = v5448
	var v5450 int32
	_ = v5450
	var v5451 int32
	_ = v5451
	var v5454 int32
	_ = v5454
	var v5459 int32
	_ = v5459
	var v5460 int32
	_ = v5460
	var v5463 int32
	_ = v5463
	var v5464 int32
	_ = v5464
	var v5466 int32
	_ = v5466
	var v5471 int32
	_ = v5471
	var v5473 int32
	_ = v5473
	var v5474 int32
	_ = v5474
	var v5479 int32
	_ = v5479
	var v5482 int32
	_ = v5482
	var v5484 int32
	_ = v5484
	var v5485 int32
	_ = v5485
	var v5488 int32
	_ = v5488
	var v5491 int32
	_ = v5491
	var v5499 int32
	_ = v5499
	var v5503 int32
	_ = v5503
	var v5504 int32
	_ = v5504
	var v5506 int32
	_ = v5506
	var v5511 int32
	_ = v5511
	var v5512 int32
	_ = v5512
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
	var v5528 int32
	_ = v5528
	var v5529 int32
	_ = v5529
	var v5535 int32
	_ = v5535
	var v5537 int32
	_ = v5537
	var v5538 int32
	_ = v5538
	var v5542 int32
	_ = v5542
	var v5544 int32
	_ = v5544
	var v5548 int32
	_ = v5548
	var v5552 int32
	_ = v5552
	var v5553 int32
	_ = v5553
	var v5558 int32
	_ = v5558
	var v5561 int32
	_ = v5561
	var v5565 int32
	_ = v5565
	var v5566 int32
	_ = v5566
	var v5567 int32
	_ = v5567
	var v5572 int32
	_ = v5572
	var v5576 int32
	_ = v5576
	var v5579 int32
	_ = v5579
	var v5583 int32
	_ = v5583
	var v5584 int32
	_ = v5584
	var v5585 int32
	_ = v5585
	var v5586 int32
	_ = v5586
	var v5591 int32
	_ = v5591
	var v5595 int32
	_ = v5595
	var v5598 int32
	_ = v5598
	var v5601 int32
	_ = v5601
	var v5604 int32
	_ = v5604
	var v5605 int32
	_ = v5605
	var v5608 int32
	_ = v5608
	var v5609 int32
	_ = v5609
	var v5612 int32
	_ = v5612
	var v5619 int32
	_ = v5619
	var v5620 int32
	_ = v5620
	var v5623 int32
	_ = v5623
	var v5624 int32
	_ = v5624
	var v5626 int32
	_ = v5626
	var v5629 int32
	_ = v5629
	var v5630 int32
	_ = v5630
	var v5632 int32
	_ = v5632
	var v5634 int32
	_ = v5634
	var v5636 int32
	_ = v5636
	var v5637 int32
	_ = v5637
	var v5640 int32
	_ = v5640
	var v5645 int32
	_ = v5645
	var v5646 int32
	_ = v5646
	var v5649 int32
	_ = v5649
	var v5650 int32
	_ = v5650
	var v5652 int32
	_ = v5652
	var v5657 int32
	_ = v5657
	var v5659 int32
	_ = v5659
	var v5660 int32
	_ = v5660
	var v5665 int32
	_ = v5665
	var v5668 int32
	_ = v5668
	var v5670 int32
	_ = v5670
	var v5671 int32
	_ = v5671
	var v5674 int32
	_ = v5674
	var v5677 int32
	_ = v5677
	var v5685 int32
	_ = v5685
	var v5689 int32
	_ = v5689
	var v5690 int32
	_ = v5690
	var v5692 int32
	_ = v5692
	var v5695 int32
	_ = v5695
	var v5696 int32
	_ = v5696
	var v5701 int32
	_ = v5701
	var v5702 int32
	_ = v5702
	var v5705 int32
	_ = v5705
	var v5708 int32
	_ = v5708
	var v5711 int32
	_ = v5711
	var v5714 int32
	_ = v5714
	var v5715 int32
	_ = v5715
	var v5718 int32
	_ = v5718
	var v5719 int32
	_ = v5719
	var v5722 int32
	_ = v5722
	var v5729 int32
	_ = v5729
	var v5730 int32
	_ = v5730
	var v5736 int32
	_ = v5736
	var v5738 int32
	_ = v5738
	var v5746 int32
	_ = v5746
	var v5747 int32
	_ = v5747
	var v5752 int32
	_ = v5752
	var v5755 int32
	_ = v5755
	var v5760 int32
	_ = v5760
	var v5761 int32
	_ = v5761
	var v5763 int32
	_ = v5763
	var v5793 int32
	_ = v5793
	var v5796 int32
	_ = v5796
	var v5801 int32
	_ = v5801
	var v5802 int32
	_ = v5802
	var v5803 int32
	_ = v5803
	var v5804 int32
	_ = v5804
	var v5805 int32
	_ = v5805
	var v5807 int32
	_ = v5807
	var v5841 int32
	_ = v5841
	var v5844 int32
	_ = v5844
	var v5848 int32
	_ = v5848
	var v5849 int32
	_ = v5849
	var v5850 int32
	_ = v5850
	var v5855 int32
	_ = v5855
	var v5859 int32
	_ = v5859
	var v5861 int32
	_ = v5861
	var v5863 int32
	_ = v5863
	var v5864 int32
	_ = v5864
	var v5866 int32
	_ = v5866
	var v5867 int32
	_ = v5867
	var v5870 int32
	_ = v5870
	var v5875 int32
	_ = v5875
	var v5876 int32
	_ = v5876
	var v5879 int32
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5882 int32
	_ = v5882
	var v5887 int32
	_ = v5887
	var v5889 int32
	_ = v5889
	var v5890 int32
	_ = v5890
	var v5895 int32
	_ = v5895
	var v5898 int32
	_ = v5898
	var v5900 int32
	_ = v5900
	var v5901 int32
	_ = v5901
	var v5904 int32
	_ = v5904
	var v5907 int32
	_ = v5907
	var v5915 int32
	_ = v5915
	var v5919 int32
	_ = v5919
	var v5920 int32
	_ = v5920
	var v5921 int32
	_ = v5921
	var v5922 int32
	_ = v5922
	var v5927 int32
	_ = v5927
	var v5930 int32
	_ = v5930
	var v5931 int32
	_ = v5931
	var v5937 int32
	_ = v5937
	var v5940 int32
	_ = v5940
	var v5944 int32
	_ = v5944
	var v5948 int32
	_ = v5948
	var v5949 int32
	_ = v5949
	var v5950 int32
	_ = v5950
	var v5951 int32
	_ = v5951
	var v5956 int32
	_ = v5956
	var v5957 int32
	_ = v5957
	var v5960 int32
	_ = v5960
	var v5961 int32
	_ = v5961
	var v5965 int32
	_ = v5965
	var v5966 int32
	_ = v5966
	var v5970 int32
	_ = v5970
	var v5975 int32
	_ = v5975
	var v5979 int32
	_ = v5979
	var v5980 int32
	_ = v5980
	var v5981 int32
	_ = v5981
	var v5982 int32
	_ = v5982
	var v5987 int32
	_ = v5987
	var v5988 int32
	_ = v5988
	var v5989 int32
	_ = v5989
	var v5990 int32
	_ = v5990
	var v5996 int32
	_ = v5996
	var v5997 int32
	_ = v5997
	var v6001 int32
	_ = v6001
	var v6002 int32
	_ = v6002
	var v6005 int32
	_ = v6005
	var v6006 int32
	_ = v6006
	var v6012 int32
	_ = v6012
	var v6014 int32
	_ = v6014
	var v6015 int32
	_ = v6015
	var v6019 int32
	_ = v6019
	var v6021 int32
	_ = v6021
	var v6025 int32
	_ = v6025
	var v6029 int32
	_ = v6029
	var v6030 int32
	_ = v6030
	var v6034 int32
	_ = v6034
	var v6038 int32
	_ = v6038
	var v6039 int32
	_ = v6039
	var v6040 int32
	_ = v6040
	var v6041 int32
	_ = v6041
	var v6046 int32
	_ = v6046
	var v6050 int32
	_ = v6050
	var v6053 int32
	_ = v6053
	var v6057 int32
	_ = v6057
	var v6058 int32
	_ = v6058
	var v6059 int32
	_ = v6059
	var v6060 int32
	_ = v6060
	var v6065 int32
	_ = v6065
	var v6071 int32
	_ = v6071
	var v6072 int32
	_ = v6072
	var v6075 int32
	_ = v6075
	var v6076 int32
	_ = v6076
	var v6081 int32
	_ = v6081
	var v6082 int32
	_ = v6082
	var v6085 int32
	_ = v6085
	var v6086 int32
	_ = v6086
	var v6088 int32
	_ = v6088
	var v6093 int32
	_ = v6093
	var v6095 int32
	_ = v6095
	var v6096 int32
	_ = v6096
	var v6101 int32
	_ = v6101
	var v6104 int32
	_ = v6104
	var v6106 int32
	_ = v6106
	var v6107 int32
	_ = v6107
	var v6110 int32
	_ = v6110
	var v6113 int32
	_ = v6113
	var v6121 int32
	_ = v6121
	var v6125 int32
	_ = v6125
	var v6126 int32
	_ = v6126
	var v6128 int32
	_ = v6128
	var v6130 int64
	_ = v6130
	var v6138 int32
	_ = v6138
	var v6142 int32
	_ = v6142
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6154 int32
	_ = v6154
	var v6159 int32
	_ = v6159
	var v6162 int32
	_ = v6162
	var v6165 int32
	_ = v6165
	var v6168 int32
	_ = v6168
	var v6171 int32
	_ = v6171
	var v6174 int32
	_ = v6174
	var v6175 int32
	_ = v6175
	var v6178 int32
	_ = v6178
	var v6179 int32
	_ = v6179
	var v6182 int32
	_ = v6182
	var v6189 int32
	_ = v6189
	var v6190 int32
	_ = v6190
	var v6199 int32
	_ = v6199
	var v6200 int32
	_ = v6200
	var v6201 int32
	_ = v6201
	var v6204 int32
	_ = v6204
	var v6207 int32
	_ = v6207
	var v6210 int32
	_ = v6210
	var v6211 int32
	_ = v6211
	var v6214 int32
	_ = v6214
	var v6215 int32
	_ = v6215
	var v6218 int32
	_ = v6218
	var v6225 int32
	_ = v6225
	var v6226 int32
	_ = v6226
	var v6235 int32
	_ = v6235
	var v6236 int32
	_ = v6236
	var v6237 int32
	_ = v6237
	var v6240 int32
	_ = v6240
	var v6243 int32
	_ = v6243
	var v6246 int32
	_ = v6246
	var v6249 int32
	_ = v6249
	var v6250 int32
	_ = v6250
	var v6253 int32
	_ = v6253
	var v6254 int32
	_ = v6254
	var v6257 int32
	_ = v6257
	var v6264 int32
	_ = v6264
	var v6265 int32
	_ = v6265
	var v6274 int32
	_ = v6274
	var v6275 int32
	_ = v6275
	var v6276 int32
	_ = v6276
	var v6279 int32
	_ = v6279
	var v6282 int32
	_ = v6282
	var v6285 int32
	_ = v6285
	var v6286 int32
	_ = v6286
	var v6289 int32
	_ = v6289
	var v6290 int32
	_ = v6290
	var v6293 int32
	_ = v6293
	var v6300 int32
	_ = v6300
	var v6301 int32
	_ = v6301
	var v6306 int32
	_ = v6306
	var v6309 int32
	_ = v6309
	var v6312 int32
	_ = v6312
	var v6315 int32
	_ = v6315
	var v6318 int32
	_ = v6318
	var v6319 int32
	_ = v6319
	var v6322 int32
	_ = v6322
	var v6323 int32
	_ = v6323
	var v6326 int32
	_ = v6326
	var v6333 int32
	_ = v6333
	var v6334 int32
	_ = v6334
	var v6339 int32
	_ = v6339
	var v6342 int32
	_ = v6342
	var v6345 int32
	_ = v6345
	var v6348 int32
	_ = v6348
	var v6349 int32
	_ = v6349
	var v6352 int32
	_ = v6352
	var v6353 int32
	_ = v6353
	var v6356 int32
	_ = v6356
	var v6363 int32
	_ = v6363
	var v6364 int32
	_ = v6364
	var v6374 int32
	_ = v6374
	var v6375 int32
	_ = v6375
	var v6377 int32
	_ = v6377
	var v6382 int32
	_ = v6382
	var v6388 int32
	_ = v6388
	var v6389 int32
	_ = v6389
	var v6396 int32
	_ = v6396
	var v6397 int32
	_ = v6397
	var v6400 int32
	_ = v6400
	var v6403 int32
	_ = v6403
	var v6406 int32
	_ = v6406
	var v6409 int32
	_ = v6409
	var v6412 int32
	_ = v6412
	var v6413 int32
	_ = v6413
	var v6416 int32
	_ = v6416
	var v6417 int32
	_ = v6417
	var v6420 int32
	_ = v6420
	var v6427 int32
	_ = v6427
	var v6428 int32
	_ = v6428
	var v6436 int32
	_ = v6436
	var v6437 int32
	_ = v6437
	var v6440 int32
	_ = v6440
	var v6441 int32
	_ = v6441
	var v6444 int32
	_ = v6444
	var v6448 int32
	_ = v6448
	var v6450 int32
	_ = v6450
	var v6451 int64
	_ = v6451
	var v6459 int32
	_ = v6459
	var v6463 int32
	_ = v6463
	var v6467 int32
	_ = v6467
	var v6473 int32
	_ = v6473
	var v6477 int32
	_ = v6477
	var v6478 int32
	_ = v6478
	var v6485 int32
	_ = v6485
	var v6486 int32
	_ = v6486
	var v6487 int32
	_ = v6487
	var v6491 int32
	_ = v6491
	var v6494 int32
	_ = v6494
	var v6498 int32
	_ = v6498
	var v6499 int32
	_ = v6499
	var v6507 int32
	_ = v6507
	var v6513 int32
	_ = v6513
	var v6515 int32
	_ = v6515
	var v6517 int32
	_ = v6517
	var v6527 int32
	_ = v6527
	var v6531 int32
	_ = v6531
	var v6536 int32
	_ = v6536
	var v6542 int32
	_ = v6542
	var v6543 int32
	_ = v6543
	var v6547 int32
	_ = v6547
	var v6551 int32
	_ = v6551
	var v6552 int32
	_ = v6552
	var v6553 int32
	_ = v6553
	var v6554 int32
	_ = v6554
	var v6557 int32
	_ = v6557
	var v6560 int32
	_ = v6560
	var v6561 int32
	_ = v6561
	var v6567 int32
	_ = v6567
	var v6568 int32
	_ = v6568
	var v6569 int32
	_ = v6569
	var v6579 int32
	_ = v6579
	var v6610 int32
	_ = v6610
	var v6619 int32
	_ = v6619
	var v6620 int32
	_ = v6620
	var v6621 int32
	_ = v6621
	var v6622 int32
	_ = v6622
	var v6623 int32
	_ = v6623
	var v6625 int32
	_ = v6625
	var v6632 int32
	_ = v6632
	var v6681 int32
	_ = v6681
	var v6683 int32
	_ = v6683
	var v6685 int32
	_ = v6685
	var v6687 int32
	_ = v6687
	var v6693 int32
	_ = v6693
	var v6714 int32
	_ = v6714
	var v6715 int32
	_ = v6715
	var v6718 int32
	_ = v6718
	var v6720 int32
	_ = v6720
	var v6721 int32
	_ = v6721
	var v6724 int32
	_ = v6724
	var v6725 int32
	_ = v6725
	var v6728 int32
	_ = v6728
	var v6731 int32
	_ = v6731
	var v6734 int32
	_ = v6734
	var v6737 int32
	_ = v6737
	var v6738 int32
	_ = v6738
	var v6741 int32
	_ = v6741
	var v6742 int32
	_ = v6742
	var v6745 int32
	_ = v6745
	var v6752 int32
	_ = v6752
	var v6753 int32
	_ = v6753
	var v6757 int32
	_ = v6757
	var v6760 int32
	_ = v6760
	var v6763 int32
	_ = v6763
	var v6766 int32
	_ = v6766
	var v6767 int32
	_ = v6767
	var v6770 int32
	_ = v6770
	var v6771 int32
	_ = v6771
	var v6774 int32
	_ = v6774
	var v6781 int32
	_ = v6781
	var v6782 int32
	_ = v6782
	var v6786 int32
	_ = v6786
	var v6789 int32
	_ = v6789
	var v6792 int32
	_ = v6792
	var v6795 int32
	_ = v6795
	var v6798 int32
	_ = v6798
	var v6799 int32
	_ = v6799
	var v6802 int32
	_ = v6802
	var v6803 int32
	_ = v6803
	var v6806 int32
	_ = v6806
	var v6813 int32
	_ = v6813
	var v6814 int32
	_ = v6814
	var v6819 int32
	_ = v6819
	var v6822 int32
	_ = v6822
	var v6825 int32
	_ = v6825
	var v6828 int32
	_ = v6828
	var v6829 int32
	_ = v6829
	var v6832 int32
	_ = v6832
	var v6833 int32
	_ = v6833
	var v6836 int32
	_ = v6836
	var v6843 int32
	_ = v6843
	var v6844 int32
	_ = v6844
	var v6849 int32
	_ = v6849
	var v6852 int32
	_ = v6852
	var v6855 int32
	_ = v6855
	var v6858 int32
	_ = v6858
	var v6861 int32
	_ = v6861
	var v6862 int32
	_ = v6862
	var v6865 int32
	_ = v6865
	var v6866 int32
	_ = v6866
	var v6869 int32
	_ = v6869
	var v6876 int32
	_ = v6876
	var v6877 int32
	_ = v6877
	var v6882 int32
	_ = v6882
	var v6885 int32
	_ = v6885
	var v6888 int32
	_ = v6888
	var v6891 int32
	_ = v6891
	var v6892 int32
	_ = v6892
	var v6895 int32
	_ = v6895
	var v6896 int32
	_ = v6896
	var v6899 int32
	_ = v6899
	var v6906 int32
	_ = v6906
	var v6907 int32
	_ = v6907
	var v6912 int32
	_ = v6912
	var v6915 int32
	_ = v6915
	var v6918 int32
	_ = v6918
	var v6921 int32
	_ = v6921
	var v6924 int32
	_ = v6924
	var v6925 int32
	_ = v6925
	var v6928 int32
	_ = v6928
	var v6929 int32
	_ = v6929
	var v6932 int32
	_ = v6932
	var v6939 int32
	_ = v6939
	var v6940 int32
	_ = v6940
	var v6945 int32
	_ = v6945
	var v6948 int32
	_ = v6948
	var v6951 int32
	_ = v6951
	var v6954 int32
	_ = v6954
	var v6955 int32
	_ = v6955
	var v6958 int32
	_ = v6958
	var v6959 int32
	_ = v6959
	var v6962 int32
	_ = v6962
	var v6969 int32
	_ = v6969
	var v6970 int32
	_ = v6970
	var v6975 int32
	_ = v6975
	var v6978 int32
	_ = v6978
	var v6981 int32
	_ = v6981
	var v6984 int32
	_ = v6984
	var v6987 int32
	_ = v6987
	var v6988 int32
	_ = v6988
	var v6991 int32
	_ = v6991
	var v6992 int32
	_ = v6992
	var v6995 int32
	_ = v6995
	var v7002 int32
	_ = v7002
	var v7003 int32
	_ = v7003
	var v7008 int32
	_ = v7008
	var v7011 int32
	_ = v7011
	var v7012 int32
	_ = v7012
	var v7021 int32
	_ = v7021
	var v7024 int32
	_ = v7024
	var v7029 int32
	_ = v7029
	var v7030 int32
	_ = v7030
	var v7032 int32
	_ = v7032
	var v7033 int32
	_ = v7033
	var v7034 int32
	_ = v7034
	var v7043 int32
	_ = v7043
	var v7049 int32
	_ = v7049
	var v7053 int32
	_ = v7053
	var v7081 int32
	_ = v7081
	var v7086 int32
	_ = v7086
	var v7087 int32
	_ = v7087
	var v7107 int32
	_ = v7107
	var v7110 int32
	_ = v7110
	var v7111 int32
	_ = v7111
	var v7116 int32
	_ = v7116
	var v7118 int32
	_ = v7118
	var v7119 int32
	_ = v7119
	var v7121 int32
	_ = v7121
	var v7122 int32
	_ = v7122
	var v7126 int32
	_ = v7126
	var v7128 int32
	_ = v7128
	var v7158 int32
	_ = v7158
	var v7161 int32
	_ = v7161
	var v7165 int32
	_ = v7165
	var v7170 int32
	_ = v7170
	var v7175 int32
	_ = v7175
	var v7178 int32
	_ = v7178
	var v7182 int32
	_ = v7182
	var v7187 int32
	_ = v7187
	var v7190 int32
	_ = v7190
	var v7191 int32
	_ = v7191
	var v7194 int32
	_ = v7194
	var v7195 int32
	_ = v7195
	var v7200 int32
	_ = v7200
	var v7201 int32
	_ = v7201
	var v7204 int32
	_ = v7204
	var v7205 int32
	_ = v7205
	var v7207 int32
	_ = v7207
	var v7212 int32
	_ = v7212
	var v7214 int32
	_ = v7214
	var v7215 int32
	_ = v7215
	var v7220 int32
	_ = v7220
	var v7223 int32
	_ = v7223
	var v7225 int32
	_ = v7225
	var v7226 int32
	_ = v7226
	var v7229 int32
	_ = v7229
	var v7232 int32
	_ = v7232
	var v7240 int32
	_ = v7240
	var v7244 int32
	_ = v7244
	var v7245 int32
	_ = v7245
	var v7246 int32
	_ = v7246
	var v7247 int32
	_ = v7247
	var v7252 int32
	_ = v7252
	var v7261 int32
	_ = v7261
	var v7263 int32
	_ = v7263
	var v7264 int32
	_ = v7264
	var v7265 int32
	_ = v7265
	var v7267 int32
	_ = v7267
	var v7271 int32
	_ = v7271
	var v7275 int32
	_ = v7275
	var v7279 int32
	_ = v7279
	var v7280 int32
	_ = v7280
	var v7282 int32
	_ = v7282
	var v7287 int32
	_ = v7287
	var v7291 int32
	_ = v7291
	var v7295 int32
	_ = v7295
	var v7298 int32
	_ = v7298
	var v7304 int32
	_ = v7304
	var v7305 int32
	_ = v7305
	var v7308 int32
	_ = v7308
	var v7314 int32
	_ = v7314
	var v7315 int32
	_ = v7315
	var v7318 int32
	_ = v7318
	var v7324 int32
	_ = v7324
	var v7325 int32
	_ = v7325
	var v7328 int32
	_ = v7328
	var v7330 int32
	_ = v7330
	var v7331 int32
	_ = v7331
	var v7332 int32
	_ = v7332
	var v7334 int32
	_ = v7334
	var v7335 int32
	_ = v7335
	var v7348 int32
	_ = v7348
	var v7349 int32
	_ = v7349
	var v7352 int32
	_ = v7352
	var v7354 int32
	_ = v7354
	var v7355 int32
	_ = v7355
	var v7356 int32
	_ = v7356
	var v7358 int32
	_ = v7358
	var v7359 int32
	_ = v7359
	var v7373 int32
	_ = v7373
	var v7374 int32
	_ = v7374
	var v7381 int32
	_ = v7381
	var v7390 int32
	_ = v7390
	var v7391 int32
	_ = v7391
	var v7393 int32
	_ = v7393
	var v7394 int32
	_ = v7394
	var v7397 int32
	_ = v7397
	var v7398 int32
	_ = v7398
	var v7403 int32
	_ = v7403
	var v7404 int32
	_ = v7404
	var v7407 int32
	_ = v7407
	var v7408 int32
	_ = v7408
	var v7410 int32
	_ = v7410
	var v7415 int32
	_ = v7415
	var v7417 int32
	_ = v7417
	var v7418 int32
	_ = v7418
	var v7423 int32
	_ = v7423
	var v7426 int32
	_ = v7426
	var v7428 int32
	_ = v7428
	var v7429 int32
	_ = v7429
	var v7432 int32
	_ = v7432
	var v7435 int32
	_ = v7435
	var v7443 int32
	_ = v7443
	var v7447 int32
	_ = v7447
	var v7448 int32
	_ = v7448
	var v7450 int32
	_ = v7450
	var v7454 int32
	_ = v7454
	var v7462 int32
	_ = v7462
	var v7467 int32
	_ = v7467
	var v7494 int32
	_ = v7494
	var v7497 int32
	_ = v7497
	var v7500 int32
	_ = v7500
	var v7502 int32
	_ = v7502
	var v7504 int32
	_ = v7504
	var v7505 int32
	_ = v7505
	var v7506 int32
	_ = v7506
	var v7508 int32
	_ = v7508
	var v7539 int32
	_ = v7539
	var v7548 int32
	_ = v7548
	var v7549 int32
	_ = v7549
	var v7550 int32
	_ = v7550
	var v7551 int32
	_ = v7551
	var v7552 int32
	_ = v7552
	var v7554 int32
	_ = v7554
	var v7559 int32
	_ = v7559
	var v7560 int32
	_ = v7560
	var v7565 int32
	_ = v7565
	var v7566 int32
	_ = v7566
	var v7571 int32
	_ = v7571
	var v7572 int32
	_ = v7572
	var v7575 int32
	_ = v7575
	var v7576 int32
	_ = v7576
	var v7578 int32
	_ = v7578
	var v7583 int32
	_ = v7583
	var v7585 int32
	_ = v7585
	var v7586 int32
	_ = v7586
	var v7591 int32
	_ = v7591
	var v7594 int32
	_ = v7594
	var v7596 int32
	_ = v7596
	var v7597 int32
	_ = v7597
	var v7600 int32
	_ = v7600
	var v7603 int32
	_ = v7603
	var v7611 int32
	_ = v7611
	var v7615 int32
	_ = v7615
	var v7616 int32
	_ = v7616
	var v7618 int32
	_ = v7618
	var v7621 int32
	_ = v7621
	var v7622 int32
	_ = v7622
	var v7626 int32
	_ = v7626
	var v7627 int32
	_ = v7627
	var v7630 int32
	_ = v7630
	var v7635 int32
	_ = v7635
	var v7636 int32
	_ = v7636
	var v7643 int32
	_ = v7643
	var v7646 int32
	_ = v7646
	var v7649 int32
	_ = v7649
	var v7652 int32
	_ = v7652
	var v7655 int32
	_ = v7655
	var v7657 int32
	_ = v7657
	var v7662 int32
	_ = v7662
	var v7663 int32
	_ = v7663
	var v7668 int32
	_ = v7668
	var v7671 int32
	_ = v7671
	var v7674 int32
	_ = v7674
	var v7677 int32
	_ = v7677
	var v7680 int32
	_ = v7680
	var v7683 int32
	_ = v7683
	var v7684 int32
	_ = v7684
	var v7687 int32
	_ = v7687
	var v7688 int32
	_ = v7688
	var v7691 int32
	_ = v7691
	var v7698 int32
	_ = v7698
	var v7699 int32
	_ = v7699
	var v7703 int32
	_ = v7703
	var v7706 int32
	_ = v7706
	var v7707 int32
	_ = v7707
	var v7710 int32
	_ = v7710
	var v7713 int32
	_ = v7713
	var v7716 int32
	_ = v7716
	var v7717 int32
	_ = v7717
	var v7720 int32
	_ = v7720
	var v7721 int32
	_ = v7721
	var v7724 int32
	_ = v7724
	var v7731 int32
	_ = v7731
	var v7732 int32
	_ = v7732
	var v7735 int32
	_ = v7735
	var v7736 int32
	_ = v7736
	var v7743 int32
	_ = v7743
	var v7744 int32
	_ = v7744
	var v7745 int32
	_ = v7745
	var v7753 int32
	_ = v7753
	var v7754 int32
	_ = v7754
	var v7759 int32
	_ = v7759
	var v7762 int32
	_ = v7762
	var v7765 int32
	_ = v7765
	var v7768 int32
	_ = v7768
	var v7771 int32
	_ = v7771
	var v7774 int32
	_ = v7774
	var v7775 int32
	_ = v7775
	var v7778 int32
	_ = v7778
	var v7779 int32
	_ = v7779
	var v7782 int32
	_ = v7782
	var v7789 int32
	_ = v7789
	var v7790 int32
	_ = v7790
	var v7795 int32
	_ = v7795
	var v7798 int32
	_ = v7798
	var v7807 int32
	_ = v7807
	var v7808 int32
	_ = v7808
	var v7810 int32
	_ = v7810
	var v7840 int32
	_ = v7840
	var v7843 int32
	_ = v7843
	var v7852 int32
	_ = v7852
	var v7853 int32
	_ = v7853
	var v7854 int32
	_ = v7854
	var v7855 int32
	_ = v7855
	var v7856 int32
	_ = v7856
	var v7858 int32
	_ = v7858
	var v7889 int32
	_ = v7889
	var v7891 int32
	_ = v7891
	var v7893 int32
	_ = v7893
	var v7894 int32
	_ = v7894
	var v7895 int32
	_ = v7895
	var v7903 int32
	_ = v7903
	var v7904 int32
	_ = v7904
	var v7909 int32
	_ = v7909
	var v7914 int32
	_ = v7914
	var v7916 int32
	_ = v7916
	var v7918 int32
	_ = v7918
	var v7919 int32
	_ = v7919
	var v7920 int32
	_ = v7920
	var v7923 int32
	_ = v7923
	var v7928 int32
	_ = v7928
	var v7929 int32
	_ = v7929
	var v7934 int32
	_ = v7934
	var v7935 int32
	_ = v7935
	var v7938 int32
	_ = v7938
	var v7939 int32
	_ = v7939
	var v7941 int32
	_ = v7941
	var v7946 int32
	_ = v7946
	var v7948 int32
	_ = v7948
	var v7949 int32
	_ = v7949
	var v7954 int32
	_ = v7954
	var v7957 int32
	_ = v7957
	var v7959 int32
	_ = v7959
	var v7960 int32
	_ = v7960
	var v7963 int32
	_ = v7963
	var v7966 int32
	_ = v7966
	var v7974 int32
	_ = v7974
	var v7977 int32
	_ = v7977
	var v7981 int32
	_ = v7981
	var v7982 int32
	_ = v7982
	var v7983 int32
	_ = v7983
	var v7989 int32
	_ = v7989
	var v7992 int32
	_ = v7992
	var v7993 int32
	_ = v7993
	var v7998 int32
	_ = v7998
	var v7999 int32
	_ = v7999
	var v8002 int32
	_ = v8002
	var v8003 int32
	_ = v8003
	var v8005 int32
	_ = v8005
	var v8010 int32
	_ = v8010
	var v8012 int32
	_ = v8012
	var v8013 int32
	_ = v8013
	var v8018 int32
	_ = v8018
	var v8021 int32
	_ = v8021
	var v8023 int32
	_ = v8023
	var v8024 int32
	_ = v8024
	var v8027 int32
	_ = v8027
	var v8030 int32
	_ = v8030
	var v8038 int32
	_ = v8038
	var v8043 int32
	_ = v8043
	var v8044 int32
	_ = v8044
	var v8045 int32
	_ = v8045
	var v8050 int32
	_ = v8050
	var v8052 int32
	_ = v8052
	var v8053 int32
	_ = v8053
	var v8055 int32
	_ = v8055
	var v8057 int32
	_ = v8057
	var v8060 int32
	_ = v8060
	var v8061 int32
	_ = v8061
	var v8065 int32
	_ = v8065
	var v8066 int32
	_ = v8066
	var v8068 int32
	_ = v8068
	var v8070 int32
	_ = v8070
	var v8077 int32
	_ = v8077
	var v8078 int32
	_ = v8078
	var v8086 int32
	_ = v8086
	var v8087 int32
	_ = v8087
	var v8088 int32
	_ = v8088
	var v8089 int32
	_ = v8089
	var v8092 int32
	_ = v8092
	var v8095 int32
	_ = v8095
	var v8098 int32
	_ = v8098
	var v8099 int32
	_ = v8099
	var v8102 int32
	_ = v8102
	var v8103 int32
	_ = v8103
	var v8106 int32
	_ = v8106
	var v8113 int32
	_ = v8113
	var v8114 int32
	_ = v8114
	var v8118 int32
	_ = v8118
	var v8121 int32
	_ = v8121
	var v8124 int32
	_ = v8124
	var v8127 int32
	_ = v8127
	var v8130 int32
	_ = v8130
	var v8131 int32
	_ = v8131
	var v8134 int32
	_ = v8134
	var v8135 int32
	_ = v8135
	var v8138 int32
	_ = v8138
	var v8145 int32
	_ = v8145
	var v8146 int32
	_ = v8146
	var v8151 int32
	_ = v8151
	var v8154 int32
	_ = v8154
	var v8157 int32
	_ = v8157
	var v8160 int32
	_ = v8160
	var v8163 int32
	_ = v8163
	var v8164 int32
	_ = v8164
	var v8167 int32
	_ = v8167
	var v8168 int32
	_ = v8168
	var v8171 int32
	_ = v8171
	var v8178 int32
	_ = v8178
	var v8179 int32
	_ = v8179
	var v8184 int32
	_ = v8184
	var v8187 int32
	_ = v8187
	var v8190 int32
	_ = v8190
	var v8193 int32
	_ = v8193
	var v8196 int32
	_ = v8196
	var v8197 int32
	_ = v8197
	var v8200 int32
	_ = v8200
	var v8201 int32
	_ = v8201
	var v8204 int32
	_ = v8204
	var v8211 int32
	_ = v8211
	var v8212 int32
	_ = v8212
	var v8217 int32
	_ = v8217
	var v8220 int32
	_ = v8220
	var v8223 int32
	_ = v8223
	var v8226 int32
	_ = v8226
	var v8229 int32
	_ = v8229
	var v8230 int32
	_ = v8230
	var v8233 int32
	_ = v8233
	var v8234 int32
	_ = v8234
	var v8237 int32
	_ = v8237
	var v8244 int32
	_ = v8244
	var v8245 int32
	_ = v8245
	var v8248 int32
	_ = v8248
	var v8252 int32
	_ = v8252
	var v8255 int32
	_ = v8255
	var v8259 int32
	_ = v8259
	var v8260 int32
	_ = v8260
	var v8262 int32
	_ = v8262
	var v8264 int32
	_ = v8264
	var v8267 int32
	_ = v8267
	var v8270 int32
	_ = v8270
	var v8273 int32
	_ = v8273
	var v8276 int32
	_ = v8276
	var v8277 int32
	_ = v8277
	var v8280 int32
	_ = v8280
	var v8281 int32
	_ = v8281
	var v8284 int32
	_ = v8284
	var v8291 int32
	_ = v8291
	var v8292 int32
	_ = v8292
	var v8299 int32
	_ = v8299
	var v8302 int32
	_ = v8302
	var v8306 int32
	_ = v8306
	var v8307 int32
	_ = v8307
	var v8309 int32
	_ = v8309
	var v8315 int32
	_ = v8315
	var v8319 int32
	_ = v8319
	var v8320 int32
	_ = v8320
	var v8323 int32
	_ = v8323
	var v8329 int32
	_ = v8329
	var v8330 int32
	_ = v8330
	var v8337 int32
	_ = v8337
	var v8341 int32
	_ = v8341
	var v8342 int32
	_ = v8342
	var v8345 int32
	_ = v8345
	var v8351 int32
	_ = v8351
	var v8355 int32
	_ = v8355
	var v8359 int32
	_ = v8359
	var v8363 int32
	_ = v8363
	var v8364 int32
	_ = v8364
	var v8367 int32
	_ = v8367
	var v8373 int32
	_ = v8373
	var v8379 int32
	_ = v8379
	var v8382 int32
	_ = v8382
	var v8387 int32
	_ = v8387
	var v8390 int32
	_ = v8390
	var v8393 int32
	_ = v8393
	var v8397 int32
	_ = v8397
	var v8398 int32
	_ = v8398
	var v8399 int32
	_ = v8399
	var v8402 int32
	_ = v8402
	var v8406 int32
	_ = v8406
	var v8407 int32
	_ = v8407
	var v8411 int32
	_ = v8411
	var v8414 int32
	_ = v8414
	var v8415 int32
	_ = v8415
	var v8421 int32
	_ = v8421
	var v8427 int32
	_ = v8427
	var v8428 int32
	_ = v8428
	var v8433 int32
	_ = v8433
	var v8434 int32
	_ = v8434
	var v8439 int32
	_ = v8439
	var v8440 int32
	_ = v8440
	var v8443 int32
	_ = v8443
	var v8444 int32
	_ = v8444
	var v8446 int32
	_ = v8446
	var v8451 int32
	_ = v8451
	var v8453 int32
	_ = v8453
	var v8454 int32
	_ = v8454
	var v8459 int32
	_ = v8459
	var v8462 int32
	_ = v8462
	var v8464 int32
	_ = v8464
	var v8465 int32
	_ = v8465
	var v8468 int32
	_ = v8468
	var v8471 int32
	_ = v8471
	var v8479 int32
	_ = v8479
	var v8483 int32
	_ = v8483
	var v8484 int32
	_ = v8484
	var v8486 int32
	_ = v8486
	var v8491 int32
	_ = v8491
	var v8492 int32
	_ = v8492
	var v8498 int32
	_ = v8498
	var v8499 int32
	_ = v8499
	var v8504 int32
	_ = v8504
	var v8505 int32
	_ = v8505
	var v8510 int32
	_ = v8510
	var v8511 int32
	_ = v8511
	var v8514 int32
	_ = v8514
	var v8515 int32
	_ = v8515
	var v8517 int32
	_ = v8517
	var v8522 int32
	_ = v8522
	var v8524 int32
	_ = v8524
	var v8525 int32
	_ = v8525
	var v8530 int32
	_ = v8530
	var v8533 int32
	_ = v8533
	var v8535 int32
	_ = v8535
	var v8536 int32
	_ = v8536
	var v8539 int32
	_ = v8539
	var v8542 int32
	_ = v8542
	var v8550 int32
	_ = v8550
	var v8554 int32
	_ = v8554
	var v8555 int32
	_ = v8555
	var v8557 int32
	_ = v8557
	var v8562 int32
	_ = v8562
	var v8568 int32
	_ = v8568
	var v8569 int32
	_ = v8569
	var v8574 int32
	_ = v8574
	var v8575 int32
	_ = v8575
	var v8580 int32
	_ = v8580
	var v8581 int32
	_ = v8581
	var v8584 int32
	_ = v8584
	var v8585 int32
	_ = v8585
	var v8587 int32
	_ = v8587
	var v8592 int32
	_ = v8592
	var v8594 int32
	_ = v8594
	var v8595 int32
	_ = v8595
	var v8600 int32
	_ = v8600
	var v8603 int32
	_ = v8603
	var v8605 int32
	_ = v8605
	var v8606 int32
	_ = v8606
	var v8609 int32
	_ = v8609
	var v8612 int32
	_ = v8612
	var v8620 int32
	_ = v8620
	var v8624 int32
	_ = v8624
	var v8625 int32
	_ = v8625
	var v8627 int32
	_ = v8627
	var v8632 int32
	_ = v8632
	var v8643 int32
	_ = v8643
	var v8644 int32
	_ = v8644
	var v8645 int32
	_ = v8645
	var v8646 int32
	_ = v8646
	var v8649 int32
	_ = v8649
	var v8650 int32
	_ = v8650
	var v8651 int32
	_ = v8651
	var v8657 int32
	_ = v8657
	var v8658 int32
	_ = v8658
	var v8663 int32
	_ = v8663
	var v8664 int32
	_ = v8664
	var v8667 int32
	_ = v8667
	var v8668 int32
	_ = v8668
	var v8670 int32
	_ = v8670
	var v8675 int32
	_ = v8675
	var v8677 int32
	_ = v8677
	var v8678 int32
	_ = v8678
	var v8683 int32
	_ = v8683
	var v8686 int32
	_ = v8686
	var v8688 int32
	_ = v8688
	var v8689 int32
	_ = v8689
	var v8692 int32
	_ = v8692
	var v8695 int32
	_ = v8695
	var v8703 int32
	_ = v8703
	var v8706 int32
	_ = v8706
	var v8707 int32
	_ = v8707
	var v8709 int32
	_ = v8709
	var v8710 int32
	_ = v8710
	var v8715 int32
	_ = v8715
	var v8717 int32
	_ = v8717
	var v8718 int32
	_ = v8718
	var v8720 int32
	_ = v8720
	var v8721 int32
	_ = v8721
	var v8722 int32
	_ = v8722
	var v8724 int32
	_ = v8724
	var v8730 int32
	_ = v8730
	var v8731 int32
	_ = v8731
	var v8733 int32
	_ = v8733
	var v8734 int32
	_ = v8734
	var v8736 int32
	_ = v8736
	var v8737 int32
	_ = v8737
	var v8738 int32
	_ = v8738
	var v8740 int32
	_ = v8740
	var v8745 int32
	_ = v8745
	var v8746 int32
	_ = v8746
	var v8751 int32
	_ = v8751
	var v8752 int32
	_ = v8752
	var v8753 int32
	_ = v8753
	var v8754 int32
	_ = v8754
	var v8756 int32
	_ = v8756
	var v8762 int32
	_ = v8762
	var v8763 int32
	_ = v8763
	var v8766 int32
	_ = v8766
	var v8767 int32
	_ = v8767
	var v8770 int32
	_ = v8770
	var v8771 int32
	_ = v8771
	var v8776 int32
	_ = v8776
	var v8777 int32
	_ = v8777
	var v8780 int32
	_ = v8780
	var v8781 int32
	_ = v8781
	var v8783 int32
	_ = v8783
	var v8788 int32
	_ = v8788
	var v8790 int32
	_ = v8790
	var v8791 int32
	_ = v8791
	var v8796 int32
	_ = v8796
	var v8799 int32
	_ = v8799
	var v8801 int32
	_ = v8801
	var v8802 int32
	_ = v8802
	var v8805 int32
	_ = v8805
	var v8808 int32
	_ = v8808
	var v8816 int32
	_ = v8816
	var v8821 int32
	_ = v8821
	var v8823 int32
	_ = v8823
	var v8827 int32
	_ = v8827
	var v8828 int32
	_ = v8828
	var v8833 int32
	_ = v8833
	var v8854 int32
	_ = v8854
	var v8855 int32
	_ = v8855
	var v8857 int32
	_ = v8857
	var v8859 int32
	_ = v8859
	var v8861 int32
	_ = v8861
	var v8862 int32
	_ = v8862
	var v8865 int32
	_ = v8865
	var v8868 int32
	_ = v8868
	var v8871 int32
	_ = v8871
	var v8872 int32
	_ = v8872
	var v8875 int32
	_ = v8875
	var v8876 int32
	_ = v8876
	var v8879 int32
	_ = v8879
	var v8886 int32
	_ = v8886
	var v8887 int32
	_ = v8887
	var v8889 int32
	_ = v8889
	var v8890 int32
	_ = v8890
	var v8896 int32
	_ = v8896
	var v8897 int32
	_ = v8897
	var v8900 int32
	_ = v8900
	var v8901 int32
	_ = v8901
	var v8904 int32
	_ = v8904
	var v8908 int32
	_ = v8908
	var v8910 int32
	_ = v8910
	var v8911 int64
	_ = v8911
	var v8919 int32
	_ = v8919
	var v8923 int32
	_ = v8923
	var v8927 int32
	_ = v8927
	var v8933 int32
	_ = v8933
	var v8937 int32
	_ = v8937
	var v8938 int32
	_ = v8938
	var v8945 int32
	_ = v8945
	var v8946 int32
	_ = v8946
	var v8947 int32
	_ = v8947
	var v8951 int32
	_ = v8951
	var v8954 int32
	_ = v8954
	var v8958 int32
	_ = v8958
	var v8959 int32
	_ = v8959
	var v8967 int32
	_ = v8967
	var v8973 int32
	_ = v8973
	var v8975 int32
	_ = v8975
	var v8977 int32
	_ = v8977
	var v8987 int32
	_ = v8987
	var v8991 int32
	_ = v8991
	var v8992 int32
	_ = v8992
	var v8993 int32
	_ = v8993
	var v8994 int32
	_ = v8994
	var v8995 int32
	_ = v8995
	var v8996 int32
	_ = v8996
	var v8997 int32
	_ = v8997
	var v9001 int32
	_ = v9001
	var v9003 int32
	_ = v9003
	var v9036 int32
	_ = v9036
	var v9040 int32
	_ = v9040
	var v9048 int32
	_ = v9048
	var v9049 int32
	_ = v9049
	var v9052 int32
	_ = v9052
	var v9056 int32
	_ = v9056
	var v9064 int32
	_ = v9064
	var v9065 int32
	_ = v9065
	var v9068 int32
	_ = v9068
	var v9072 int32
	_ = v9072
	var v9080 int32
	_ = v9080
	var v9081 int32
	_ = v9081
	var v9083 int32
	_ = v9083
	var v9086 int32
	_ = v9086
	var v9090 int32
	_ = v9090
	var v9091 int32
	_ = v9091
	var v9094 int32
	_ = v9094
	var v9095 int32
	_ = v9095
	var v9100 int32
	_ = v9100
	var v9104 int32
	_ = v9104
	var v9105 int32
	_ = v9105
	var v9108 int32
	_ = v9108
	var v9109 int32
	_ = v9109
	var v9113 int32
	_ = v9113
	var v9117 int32
	_ = v9117
	var v9119 int32
	_ = v9119
	var v9121 int32
	_ = v9121
	var v9122 int32
	_ = v9122
	var v9123 int32
	_ = v9123
	var v9125 int32
	_ = v9125
	var v9131 int32
	_ = v9131
	var v9144 int32
	_ = v9144
	var v9148 int32
	_ = v9148
	var v9151 int32
	_ = v9151
	var v9152 int32
	_ = v9152
	var v9153 int32
	_ = v9153
	var v9154 int32
	_ = v9154
	var v9155 int32
	_ = v9155
	var v9161 int32
	_ = v9161
	var v9164 int32
	_ = v9164
	var v9165 int32
	_ = v9165
	var v9166 int32
	_ = v9166
	var v9171 int32
	_ = v9171
	var v9175 int32
	_ = v9175
	var v9178 int32
	_ = v9178
	var v9179 int32
	_ = v9179
	var v9185 int32
	_ = v9185
	var v9188 int32
	_ = v9188
	var v9189 int32
	_ = v9189
	var v9190 int32
	_ = v9190
	var v9195 int32
	_ = v9195
	var v9199 int32
	_ = v9199
	var v9202 int32
	_ = v9202
	var v9203 int32
	_ = v9203
	var v9209 int32
	_ = v9209
	var v9210 int32
	_ = v9210
	var v9211 int32
	_ = v9211
	var v9212 int32
	_ = v9212
	var v9217 int32
	_ = v9217
	var v9221 int32
	_ = v9221
	var v9224 int32
	_ = v9224
	var v9225 int32
	_ = v9225
	var v9231 int32
	_ = v9231
	var v9232 int32
	_ = v9232
	var v9233 int32
	_ = v9233
	var v9234 int32
	_ = v9234
	var v9239 int32
	_ = v9239
	var v9244 int32
	_ = v9244
	var v9247 int32
	_ = v9247
	var v9248 int32
	_ = v9248
	var v9249 int32
	_ = v9249
	var v9250 int32
	_ = v9250
	var v9256 int32
	_ = v9256
	var v9257 int32
	_ = v9257
	var v9258 int32
	_ = v9258
	var v9259 int32
	_ = v9259
	var v9264 int32
	_ = v9264
	var v9269 int32
	_ = v9269
	var v9273 int32
	_ = v9273
	var v9278 int32
	_ = v9278
	var v9282 int32
	_ = v9282
	var v9285 int32
	_ = v9285
	var v9286 int32
	_ = v9286
	var v9287 int32
	_ = v9287
	var v9293 int32
	_ = v9293
	var v9294 int32
	_ = v9294
	var v9295 int32
	_ = v9295
	var v9296 int32
	_ = v9296
	var v9301 int32
	_ = v9301
	var v9305 int32
	_ = v9305
	var v9308 int32
	_ = v9308
	var v9312 int32
	_ = v9312
	var v9313 int32
	_ = v9313
	var v9314 int32
	_ = v9314
	var v9315 int32
	_ = v9315
	var v9320 int32
	_ = v9320
	var v9324 int32
	_ = v9324
	var v9327 int32
	_ = v9327
	var v9331 int32
	_ = v9331
	var v9334 int32
	_ = v9334
	var v9335 int32
	_ = v9335
	var v9336 int32
	_ = v9336
	var v9341 int32
	_ = v9341
	var v9345 int32
	_ = v9345
	var v9348 int32
	_ = v9348
	var v9352 int32
	_ = v9352
	var v9353 int32
	_ = v9353
	var v9354 int32
	_ = v9354
	var v9359 int32
	_ = v9359
	var v9363 int32
	_ = v9363
	var v9366 int32
	_ = v9366
	var v9370 int32
	_ = v9370
	var v9373 int32
	_ = v9373
	var v9374 int32
	_ = v9374
	var v9375 int32
	_ = v9375
	var v9380 int32
	_ = v9380
	var v9384 int32
	_ = v9384
	var v9387 int32
	_ = v9387
	var v9388 int32
	_ = v9388
	var v9394 int32
	_ = v9394
	var v9397 int32
	_ = v9397
	var v9398 int32
	_ = v9398
	var v9399 int32
	_ = v9399
	var v9404 int32
	_ = v9404
	var v9408 int32
	_ = v9408
	var v9411 int32
	_ = v9411
	var v9414 int32
	_ = v9414
	var v9415 int32
	_ = v9415
	var v9418 int32
	_ = v9418
	var v9419 int32
	_ = v9419
	var v9420 int32
	_ = v9420
	var v9421 int32
	_ = v9421
	var v9426 int32
	_ = v9426
	var v9428 int32
	_ = v9428
	var v9430 int32
	_ = v9430
	var v9434 int32
	_ = v9434
	var v9437 int32
	_ = v9437
	var v9441 int32
	_ = v9441
	var v9444 int32
	_ = v9444
	var v9445 int32
	_ = v9445
	var v9446 int32
	_ = v9446
	var v9451 int32
	_ = v9451
	var v9455 int32
	_ = v9455
	var v9458 int32
	_ = v9458
	var v9462 int32
	_ = v9462
	var v9463 int32
	_ = v9463
	var v9464 int32
	_ = v9464
	var v9465 int32
	_ = v9465
	var v9470 int32
	_ = v9470
	var v9474 int32
	_ = v9474
	var v9477 int32
	_ = v9477
	var v9478 int32
	_ = v9478
	var v9479 int32
	_ = v9479
	var v9485 int32
	_ = v9485
	var v9486 int32
	_ = v9486
	var v9487 int32
	_ = v9487
	var v9488 int32
	_ = v9488
	var v9493 int32
	_ = v9493
	var v9498 int32
	_ = v9498
	var v9500 int32
	_ = v9500
	var v9502 int32
	_ = v9502
	var v9505 int32
	_ = v9505
	var v9507 int32
	_ = v9507
	var v9509 int32
	_ = v9509
	var v9516 int32
	_ = v9516
	var v9518 int32
	_ = v9518
	var v9532 int32
	_ = v9532
	var v9535 int32
	_ = v9535
	var v9538 int32
	_ = v9538
	var v9541 int32
	_ = v9541
	var v9546 int32
	_ = v9546
	var v9554 int32
	_ = v9554
	var v9555 int32
	_ = v9555
	var v9560 int32
	_ = v9560
	var v9563 int32
	_ = v9563
	var v9566 int64
	_ = v9566
	var v9568 int64
	_ = v9568
	var v9583 int32
	_ = v9583
	var v9584 int32
	_ = v9584
	var v9587 int32
	_ = v9587
	var v9590 int32
	_ = v9590
	var v9593 int32
	_ = v9593
	var v9596 int32
	_ = v9596
	var v9597 int32
	_ = v9597
	var v9600 int32
	_ = v9600
	var v9601 int32
	_ = v9601
	var v9604 int32
	_ = v9604
	var v9611 int32
	_ = v9611
	var v9612 int32
	_ = v9612
	var v9616 int32
	_ = v9616
	var v9619 int32
	_ = v9619
	var v9632 int32
	_ = v9632
	var v9633 int32
	_ = v9633
	var v9634 int32
	_ = v9634
	var v9640 int32
	_ = v9640
	var v9643 int32
	_ = v9643
	var v9647 int32
	_ = v9647
	var v9648 int32
	_ = v9648
	var v9649 int32
	_ = v9649
	var v9654 int32
	_ = v9654
	var v9658 int32
	_ = v9658
	var v9660 int32
	_ = v9660
	var v9662 int32
	_ = v9662
	var v9663 int32
	_ = v9663
	var v9675 int32
	_ = v9675
	var v9676 int32
	_ = v9676
	var v9677 int32
	_ = v9677
	var v9680 int32
	_ = v9680
	var v9681 int32
	_ = v9681
	var v9684 int32
	_ = v9684
	var v9686 int32
	_ = v9686
	var v9688 int32
	_ = v9688
	var v9690 int32
	_ = v9690
	var v9691 int32
	_ = v9691
	var v9697 int32
	_ = v9697
	var v9703 int32
	_ = v9703
	var v9705 int32
	_ = v9705
	var v9706 int32
	_ = v9706
	var v9707 int32
	_ = v9707
	var v9708 int32
	_ = v9708
	var v9712 int32
	_ = v9712
	var v9716 int32
	_ = v9716
	var v9720 int32
	_ = v9720
	var v9721 int32
	_ = v9721
	var v9722 int32
	_ = v9722
	var v9725 int32
	_ = v9725
	var v9728 int32
	_ = v9728
	var v9731 int32
	_ = v9731
	var v9734 int32
	_ = v9734
	var v9737 int32
	_ = v9737
	var v9739 int32
	_ = v9739
	var v9740 int32
	_ = v9740
	var v9742 int32
	_ = v9742
	var v9743 int32
	_ = v9743
	var v9745 int32
	_ = v9745
	var v9746 int32
	_ = v9746
	var v9750 int32
	_ = v9750
	var v9751 int32
	_ = v9751
	var v9753 int32
	_ = v9753
	var v9762 int32
	_ = v9762
	var v9763 int32
	_ = v9763
	var v9764 int32
	_ = v9764
	var v9766 int32
	_ = v9766
	var v9768 int32
	_ = v9768
	var v9769 int32
	_ = v9769
	var v9773 int32
	_ = v9773
	var v9774 int32
	_ = v9774
	var v9776 int32
	_ = v9776
	var v9779 int32
	_ = v9779
	var v9780 int32
	_ = v9780
	var v9782 int32
	_ = v9782
	var v9783 int32
	_ = v9783
	var v9785 int32
	_ = v9785
	var v9786 int32
	_ = v9786
	var v9788 int32
	_ = v9788
	var v9791 int32
	_ = v9791
	var v9796 int32
	_ = v9796
	var v9799 int32
	_ = v9799
	var v9800 int32
	_ = v9800
	var v9801 int32
	_ = v9801
	var v9802 int32
	_ = v9802
	var v9817 int32
	_ = v9817
	var v9833 int32
	_ = v9833
	var v9836 int32
	_ = v9836
	var v9837 int64
	_ = v9837
	var v9839 int64
	_ = v9839
	var v9843 int32
	_ = v9843
	var v9844 int32
	_ = v9844
	var v9846 int32
	_ = v9846
	var v9847 int32
	_ = v9847
	var v9850 int32
	_ = v9850
	var v9854 int32
	_ = v9854
	var v9857 int32
	_ = v9857
	var v9858 int32
	_ = v9858
	var v9862 int32
	_ = v9862
	var v9865 int32
	_ = v9865
	var v9871 int32
	_ = v9871
	var v9875 int32
	_ = v9875
	var v9881 int32
	_ = v9881
	var v9882 int32
	_ = v9882
	var v9886 int32
	_ = v9886
	var v9888 int32
	_ = v9888
	var v9889 int32
	_ = v9889
	var v9895 int32
	_ = v9895
	var v9906 int32
	_ = v9906
	var v9909 int32
	_ = v9909
	var v9913 int32
	_ = v9913
	var v9916 int32
	_ = v9916
	var v9917 int32
	_ = v9917
	var v9918 int32
	_ = v9918
	var v9923 int32
	_ = v9923
	var v9927 int32
	_ = v9927
	var v9930 int32
	_ = v9930
	var v9934 int32
	_ = v9934
	var v9935 int32
	_ = v9935
	var v9936 int32
	_ = v9936
	var v9937 int32
	_ = v9937
	var v9942 int32
	_ = v9942
	var v9947 int32
	_ = v9947
	var v9957 int32
	_ = v9957
	var v9972 int32
	_ = v9972
	var v9980 int32
	_ = v9980
	var v9982 int32
	_ = v9982
	var v9983 int32
	_ = v9983
	var v9985 int32
	_ = v9985
	var v9995 int32
	_ = v9995
	var v10026 int32
	_ = v10026
	var v10032 int32
	_ = v10032
	var v10042 int32
	_ = v10042
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
	v58 = int32(-2)
	v60 = v42
	v61 = v46
	v62 = v46
	v65 = int32(200)
	v66 = v44
	v67 = v3
	v69 = v42
	goto L9
L1:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_3))
	mBase = m.M
	v10042 = m.ExcPending
	if v10042 != 0 {
		goto L49
	} else {
		goto L2588
	}
L2:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_4))
	mBase = m.M
	v10032 = m.ExcPending
	if v10032 != 0 {
		goto L49
	} else {
		goto L2587
	}
L3:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_5))
	mBase = m.M
	v10026 = m.ExcPending
	if v10026 != 0 {
		goto L49
	} else {
		goto L2586
	}
L4:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v9995 = m.ExcPending
	if v9995 != 0 {
		goto L49
	} else {
		goto L2585
	}
L5:
	;
	v9983 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_cword_is_not_variable(m, v148, v9983, l1)
	mBase = m.M
	v9985 = m.ExcPending
	if v9985 != 0 {
		goto L49
	} else {
		goto L2584
	}
L6:
	;
	v9980 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_word_is_not_variable(m, v148, v9980, l1)
	mBase = m.M
	v9982 = m.ExcPending
	if v9982 != 0 {
		goto L49
	} else {
		goto L2583
	}
L7:
	;
	if v28+int32(_a_F_plpgsql_yyparse_1) != v9957 {
		goto L2579
	} else {
		goto L2580
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9927 = m.ExcPending
	if v9927 != 0 {
		goto L49
	} else {
		goto L2574
	}
L9:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v53)
	v75 = v65 << (uint(int32(1)) % 32)
	if base.Ui32(v61) < base.Ui32(v62+v75-int32(2)) {
		goto L42
	} else {
		goto L43
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9906 = m.ExcPending
	if v9906 != 0 {
		goto L49
	} else {
		goto L2569
	}
L11:
	;
	goto L10
L12:
	;
	v53 = v9881
	v54 = v9882
	v58 = v9886
	v60 = v9888
	v61 = v9889 + int32(2)
	v62 = v153
	v65 = v154
	v66 = v155
	v67 = v9895
	v69 = v156
	goto L9
L13:
	;
	v9833 = int32(0) - v239
	v9836 = v148 + v9833<<(uint(int32(4))%32)
	v9837 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v9836)+24)) = v9837
	v9839 = *(*int64)(unsafe.Add(mBase, uint32(v28)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v9836)+16)) = v9839
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v254
	v9843 = v9836 + int32(16)
	v9844 = int32(1)
	v9846 = v152 + v9833<<(uint(v9844)%32)
	v9847 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9846))))
	v9850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+uint32(_c_F_plpgsql_yyparse[3]))))
	v9854 = (v9850 - int32(137)) << (uint(v9844) % 32)
	v9857 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9854)+uint32(_c_F_plpgsql_yyparse[4]))))
	v9858 = v9847 + v9857
	if base.Ui32(int32(1305)) < base.Ui32(v9858) {
		goto L2566
	} else {
		goto L2567
	}
L14:
	;
	v9762 = *(*int32)(unsafe.Add(mBase, uint32(v9675)))
	v9763 = *(*int32)(unsafe.Add(mBase, uint32(v9675)+4))
	v9764 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	F_check_sql_expr(m, v9762, v9763, v9764, l1)
	mBase = m.M
	v9766 = m.ExcPending
	if v9766 != 0 {
		goto L49
	} else {
		goto L2557
	}
L15:
	;
	v9684 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v9680)+4)) = v9684
	v9686 = *(*int32)(unsafe.Add(mBase, uint32(v9680)))
	v9688 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	F_check_sql_expr(m, v9686, v9684, v9688, l1)
	mBase = m.M
	v9690 = m.ExcPending
	if v9690 != 0 {
		goto L49
	} else {
		goto L2544
	}
L16:
	;
	v9658 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v9660 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, v4761, v9658, v9660, l1)
	mBase = m.M
	v9662 = m.ExcPending
	if v9662 != 0 {
		goto L49
	} else {
		goto L2541
	}
L17:
	;
	v9616 = int32(1)
	v9619 = int32(0)
	v9632 = F_read_sql_construct(m, int32(269), int32(336), v9619, int32(_a_F_plpgsql_yyparse_9), v9619, v9616, v9619, v28+int32(240), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9633 = m.ExcPending
	if v9633 != 0 {
		goto L49
	} else {
		goto L2534
	}
L18:
	;
	v9584 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v9584 == int32(0) {
		goto L16
	} else {
		goto L2525
	}
L19:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_10))
	mBase = m.M
	v9583 = m.ExcPending
	if v9583 != 0 {
		goto L49
	} else {
		goto L2524
	}
L20:
	;
	v9505 = v164
	v9507 = v147
	v9509 = v148
	v9516 = v152
	v9518 = v9502
	goto L2513
L21:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), l0, l1, int32(_a_F_plpgsql_yyparse_5))
	mBase = m.M
	v9498 = m.ExcPending
	if v9498 != 0 {
		goto L49
	} else {
		goto L2512
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9474 = m.ExcPending
	if v9474 != 0 {
		goto L49
	} else {
		goto L2507
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9455 = m.ExcPending
	if v9455 != 0 {
		goto L49
	} else {
		goto L2502
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9434 = m.ExcPending
	if v9434 != 0 {
		goto L49
	} else {
		goto L2497
	}
L25:
	;
	F_cword_is_not_variable(m, v148, v7359, l1)
	mBase = m.M
	v9430 = m.ExcPending
	if v9430 != 0 {
		goto L49
	} else {
		goto L2496
	}
L26:
	;
	F_word_is_not_variable(m, v148, v7335, l1)
	mBase = m.M
	v9428 = m.ExcPending
	if v9428 != 0 {
		goto L49
	} else {
		goto L2495
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9408 = m.ExcPending
	if v9408 != 0 {
		goto L49
	} else {
		goto L2487
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9384 = m.ExcPending
	if v9384 != 0 {
		goto L49
	} else {
		goto L2482
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9363 = m.ExcPending
	if v9363 != 0 {
		goto L49
	} else {
		goto L2477
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9345 = m.ExcPending
	if v9345 != 0 {
		goto L49
	} else {
		goto L2472
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9324 = m.ExcPending
	if v9324 != 0 {
		goto L49
	} else {
		goto L2467
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9305 = m.ExcPending
	if v9305 != 0 {
		goto L49
	} else {
		goto L2462
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9282 = m.ExcPending
	if v9282 != 0 {
		goto L49
	} else {
		goto L2456
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9269 = m.ExcPending
	if v9269 != 0 {
		goto L49
	} else {
		goto L2453
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9244 = m.ExcPending
	if v9244 != 0 {
		goto L49
	} else {
		goto L2447
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9221 = m.ExcPending
	if v9221 != 0 {
		goto L49
	} else {
		goto L2442
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9199 = m.ExcPending
	if v9199 != 0 {
		goto L49
	} else {
		goto L2437
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9175 = m.ExcPending
	if v9175 != 0 {
		goto L49
	} else {
		goto L2432
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9148 = m.ExcPending
	if v9148 != 0 {
		goto L49
	} else {
		goto L2426
	}
L40:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), l0, l1, int32(_a_F_plpgsql_yyparse_11))
	mBase = m.M
	v9144 = m.ExcPending
	if v9144 != 0 {
		goto L49
	} else {
		goto L2425
	}
L41:
	;
	if v53 == int32(3) {
		goto L68
	} else {
		goto L69
	}
L42:
	;
	v147 = v60
	v148 = v54
	v152 = v61
	v153 = v62
	v154 = v65
	v155 = v66
	v156 = v69
	goto L41
L43:
	;
	goto L44
L44:
	;
	if int32(_a_F_plpgsql_yyparse_12) < v65 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v82 = int32(_a_F_plpgsql_yyparse_13)
	if v82 <= v75 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v85 = v82
	goto L48
L47:
	;
	v85 = v75
	goto L48
L48:
	;
	v90 = F_palloc(m, v85*int32(22)+int32(30))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	if v90 == int32(0) {
		goto L40
	} else {
		goto L51
	}
L51:
	;
	v97 = int32(1)
	v100 = (v61-v62)>>(uint(v97)%32) + v97
	v102 = v100 << (uint(v97) % 32)
	if v102 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	base.MemoryCopy(m, v90, v62, v102)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v109 = base.I32_div_s(v85<<(uint(int32(1))%32)+int32(15), int32(16))
	v110 = int32(4)
	v112 = v90 + v109<<(uint(v110)%32)
	v114 = v100 << (uint(v110) % 32)
	if v114 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, v112, v66, v114)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v116 = int32(4)
	v121 = base.I32_div_s(v85<<(uint(v116)%32)|int32(15), int32(16))
	v124 = v112 + v121<<(uint(v116)%32)
	v126 = v100 << (uint(int32(2)) % 32)
	if v126 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	base.MemoryCopy(m, v124, v69, v126)
	goto L60
L59:
	;
	goto L60
L60:
	;
	if v28+int32(_a_F_plpgsql_yyparse_1) != v62 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_pfree(m, v62)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L49
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v85 <= v100 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v9947 = int32(1)
	v9957 = v90
	goto L7
L66:
	;
	goto L67
L67:
	;
	v147 = v126 + v124 - int32(4)
	v148 = v114 + v112 - int32(16)
	v152 = v90 + v100<<(uint(int32(1))%32) - int32(2)
	v153 = v90
	v154 = v85
	v155 = v112
	v156 = v124
	goto L41
L68:
	;
	v9947 = int32(0)
	v9957 = v153
	goto L7
L69:
	;
	goto L70
L70:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[6]))))
	if v164 == int32(_a_F_plpgsql_yyparse_14) {
		v224 = v58
		goto L73
	} else {
		goto L74
	}
L71:
	;
	if v67 == int32(0) {
		goto L21
	} else {
		goto L2419
	}
L72:
	;
	v239 = int32(*(*int8)(unsafe.Add(mBase, uint32(v235)+uint32(_c_F_plpgsql_yyparse[7]))))
	v241 = int32(4)
	v243 = v148 + (int32(1)-v239)<<(uint(v241)%32)
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+280)) = v244
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v243)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+272)) = v246
	v252 = v147 - v239<<(uint(int32(2))%32) + v241
	if v239 != 0 {
		goto L93
	} else {
		goto L94
	}
L73:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_plpgsql_yyparse[8]))))
	if v228 == int32(0) {
		goto L71
	} else {
		goto L92
	}
L74:
	;
	if v58 == int32(-2) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v193 = v191 + base.I32_extend16_s(v164)
	if base.Ui32(int32(1305)) < base.Ui32(v193) {
		v224 = v190
		goto L73
	} else {
		goto L87
	}
L76:
	;
	v173 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L49
	} else {
		goto L79
	}
L77:
	;
	v175 = v58
	goto L78
L78:
	;
	if v175 <= int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v175 = v173
	goto L78
L80:
	;
	v178 = int32(0)
	v190 = v178
	v191 = v178
	goto L75
L81:
	;
	goto L82
L82:
	;
	if v175 == int32(256) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v9500 = int32(257)
	v9502 = v183
	goto L20
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(385)) < base.Ui32(v175) {
		v190 = v175
		v191 = int32(2)
		goto L75
	} else {
		goto L86
	}
L86:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+uint32(_c_F_plpgsql_yyparse[9]))))
	v190 = v175
	v191 = v189
	goto L75
L87:
	;
	v197 = v193 << (uint(int32(1)) % 32)
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_plpgsql_yyparse[10]))))
	if v191 != v200 {
		v224 = v190
		goto L73
	} else {
		goto L88
	}
L88:
	;
	v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v197)+uint32(_c_F_plpgsql_yyparse[11]))))
	if int32(0) < v204 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+24)) = v207
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+16)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v211
	v9881 = v204
	v9882 = v148 + int32(16)
	v9886 = int32(-2)
	v9888 = v147 + int32(4)
	v9889 = v152
	v9895 = v67 - base.B2i32(v67 != int32(0))
	goto L12
L90:
	;
	goto L91
L91:
	;
	v233 = v190
	v235 = int32(0) - v204
	goto L72
L92:
	;
	v233 = v224
	v235 = v228
	goto L72
L93:
	;
	v253 = v252
	goto L95
L94:
	;
	v253 = v147
	goto L95
L95:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	switch v235 - int32(2) {
	case 0:
		goto L246
	default:
		v9817 = v233
		goto L13
	case 3:
		goto L245
	case 4:
		goto L244
	case 5:
		goto L243
	case 6:
		goto L242
	case 7:
		goto L241
	case 8:
		goto L240
	case 9:
		goto L239
	case 12:
		goto L238
	case 13:
		goto L237
	case 14:
		goto L236
	case 15:
		goto L235
	case 16:
		goto L234
	case 21:
		goto L233
	case 22:
		goto L232
	case 23:
		goto L231
	case 24:
		goto L230
	case 25:
		goto L229
	case 26:
		goto L228
	case 27:
		goto L227
	case 28:
		goto L226
	case 29:
		goto L225
	case 30:
		goto L224
	case 31:
		goto L223
	case 32:
		goto L222
	case 33:
		goto L221
	case 34:
		goto L220
	case 37:
		goto L219
	case 38:
		goto L218
	case 39:
		goto L217
	case 40:
		goto L216
	case 41:
		goto L215
	case 42:
		goto L214
	case 43:
		goto L213
	case 44:
		goto L212
	case 45:
		goto L211
	case 46:
		goto L210
	case 47:
		goto L209
	case 48:
		goto L208
	case 49:
		goto L207
	case 50:
		goto L206
	case 51:
		goto L205
	case 52:
		goto L204
	case 57:
		goto L203
	case 58:
		goto L202
	case 59:
		goto L201
	case 60:
		goto L200
	case 61:
		goto L199
	case 62:
		goto L198
	case 63:
		goto L197
	case 64:
		goto L196
	case 65:
		goto L195
	case 66:
		goto L194
	case 67:
		goto L193
	case 68:
		goto L192
	case 69:
		goto L191
	case 70:
		goto L190
	case 71:
		goto L189
	case 72:
		goto L188
	case 73:
		goto L187
	case 74:
		goto L186
	case 75:
		goto L185
	case 76:
		goto L184
	case 77:
		goto L183
	case 78:
		goto L182
	case 79:
		goto L181
	case 80:
		goto L180
	case 81:
		goto L179
	case 82:
		goto L178
	case 83:
		goto L177
	case 84:
		goto L176
	case 85:
		goto L175
	case 86:
		goto L174
	case 87:
		goto L173
	case 88:
		goto L172
	case 89:
		goto L171
	case 90:
		goto L170
	case 91:
		goto L169
	case 92:
		goto L168
	case 93:
		goto L167
	case 94:
		goto L166
	case 95:
		goto L165
	case 96, 145:
		goto L6
	case 97, 116, 146:
		goto L5
	case 98:
		goto L164
	case 99:
		goto L163
	case 100:
		goto L162
	case 101:
		goto L161
	case 102:
		goto L160
	case 103:
		goto L159
	case 104:
		goto L158
	case 105:
		goto L157
	case 106:
		goto L156
	case 107:
		goto L155
	case 108:
		goto L154
	case 109:
		goto L153
	case 110:
		goto L152
	case 111:
		goto L151
	case 112:
		goto L150
	case 113:
		goto L149
	case 114:
		goto L148
	case 115:
		goto L147
	case 117:
		goto L146
	case 118:
		goto L145
	case 119:
		goto L144
	case 120:
		goto L143
	case 121:
		goto L142
	case 122:
		goto L141
	case 123:
		goto L140
	case 124:
		goto L139
	case 125:
		goto L138
	case 126:
		goto L137
	case 127:
		goto L136
	case 128:
		goto L135
	case 129:
		goto L134
	case 130:
		goto L133
	case 131:
		goto L132
	case 132:
		goto L131
	case 133:
		goto L130
	case 134:
		goto L129
	case 135:
		goto L128
	case 136:
		goto L127
	case 137:
		goto L126
	case 138:
		goto L125
	case 139:
		goto L124
	case 140:
		goto L123
	case 141:
		goto L122
	case 142:
		goto L121
	case 143:
		goto L120
	case 144:
		goto L119
	case 147:
		goto L118
	case 148:
		goto L117
	case 149:
		goto L116
	case 150:
		goto L115
	case 151:
		goto L114
	case 152:
		goto L113
	case 153:
		goto L112
	case 154:
		goto L111
	case 155:
		goto L110
	case 156:
		goto L109
	case 157:
		goto L108
	case 158:
		goto L107
	case 159:
		goto L106
	case 160:
		goto L105
	case 161:
		goto L104
	case 162:
		goto L103
	case 163:
		goto L102
	case 164:
		goto L101
	case 165:
		goto L100
	case 166:
		goto L99
	case 167:
		goto L98
	case 168:
		goto L97
	case 169:
		goto L96
	}
L96:
	;
	v9125 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v9125 == int32(0) {
		goto L3
	} else {
		goto L2418
	}
L97:
	;
	v9121 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v9122 = F_pstrdup(m, v9121)
	mBase = m.M
	v9123 = m.ExcPending
	if v9123 != 0 {
		goto L49
	} else {
		goto L2417
	}
L98:
	;
	v9119 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9119
	v9817 = v233
	goto L13
L99:
	;
	v9117 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9117
	v9817 = v233
	goto L13
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L101:
	;
	v9113 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9113
	v9817 = v233
	goto L13
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L103:
	;
	v9104 = v148 - int32(16)
	v9105 = *(*int32)(unsafe.Add(mBase, uint32(v9104)))
	F_plpgsql_ns_push(m, v9105, int32(1))
	mBase = m.M
	v9108 = m.ExcPending
	if v9108 != 0 {
		goto L49
	} else {
		goto L2416
	}
L104:
	;
	F_plpgsql_ns_push(m, int32(0), int32(1))
	mBase = m.M
	v9100 = m.ExcPending
	if v9100 != 0 {
		goto L49
	} else {
		goto L2415
	}
L105:
	;
	v9090 = v148 - int32(16)
	v9091 = *(*int32)(unsafe.Add(mBase, uint32(v9090)))
	F_plpgsql_ns_push(m, v9091, int32(0))
	mBase = m.M
	v9094 = m.ExcPending
	if v9094 != 0 {
		goto L49
	} else {
		goto L2414
	}
L106:
	;
	v9083 = int32(0)
	F_plpgsql_ns_push(m, v9083, v9083)
	mBase = m.M
	v9086 = m.ExcPending
	if v9086 != 0 {
		goto L49
	} else {
		goto L2413
	}
L107:
	;
	v9068 = int32(0)
	v9072 = int32(1)
	v9080 = F_read_sql_construct(m, int32(336), v9068, v9068, int32(_a_F_plpgsql_yyparse_9), int32(2), v9072, v9072, v9068, v9068, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9081 = m.ExcPending
	if v9081 != 0 {
		goto L49
	} else {
		goto L2412
	}
L108:
	;
	v9052 = int32(0)
	v9056 = int32(1)
	v9064 = F_read_sql_construct(m, int32(376), v9052, v9052, int32(_a_F_plpgsql_yyparse_15), int32(2), v9056, v9056, v9052, v9052, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9065 = m.ExcPending
	if v9065 != 0 {
		goto L49
	} else {
		goto L2411
	}
L109:
	;
	v9036 = int32(0)
	v9040 = int32(1)
	v9048 = F_read_sql_construct(m, int32(59), v9036, v9036, int32(_a_F_plpgsql_yyparse_16), int32(2), v9040, v9040, v9036, v9036, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v9049 = m.ExcPending
	if v9049 != 0 {
		goto L49
	} else {
		goto L2410
	}
L110:
	;
	v8861 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8862 = int32(_a_F_plpgsql_yyparse_17)
	v8865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8861))))
	v8868 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])))
	if base.B2i32(v8865 == int32(0))|base.B2i32(v8865 != v8868) != 0 {
		v8886 = v8865
		v8887 = v8868
		goto L2376
	} else {
		goto L2377
	}
L111:
	;
	v8859 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8859
	v9817 = v233
	goto L13
L112:
	;
	v8827 = v148 - int32(32)
	v8828 = *(*int32)(unsafe.Add(mBase, uint32(v8827)))
	v8833 = v8828
	goto L2372
L113:
	;
	v8766 = F_palloc0(m, int32(12))
	mBase = m.M
	v8767 = m.ExcPending
	if v8767 != 0 {
		goto L49
	} else {
		goto L2358
	}
L114:
	;
	v8756 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+232)) = v8756
	*(*int32)(unsafe.Add(mBase, uint32(v28)+236)) = v8756
	v8762 = F_list_make1_impl(m, int32(1), v28+int32(232))
	mBase = m.M
	v8763 = m.ExcPending
	if v8763 != 0 {
		goto L49
	} else {
		goto L2357
	}
L115:
	;
	v8751 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8752 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8753 = F_lappend(m, v8751, v8752)
	mBase = m.M
	v8754 = m.ExcPending
	if v8754 != 0 {
		goto L49
	} else {
		goto L2356
	}
L116:
	;
	v8745 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8746 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v8745)+8)) = v8746
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8745
	v9817 = v233
	goto L13
L117:
	;
	v8657 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v8658 = int32(0)
	if v8657 < v8658 {
		v8703 = v8658
		goto L2339
	} else {
		goto L2340
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L119:
	;
	v8643 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8644 = *(*int32)(unsafe.Add(mBase, uint32(v8643)))
	if v8644 != 0 {
		goto L23
	} else {
		goto L2334
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(1)
	v9817 = v233
	goto L13
L123:
	;
	v8568 = F_palloc(m, int32(16))
	mBase = m.M
	v8569 = m.ExcPending
	if v8569 != 0 {
		goto L49
	} else {
		goto L2320
	}
L124:
	;
	v8498 = F_palloc(m, int32(16))
	mBase = m.M
	v8499 = m.ExcPending
	if v8499 != 0 {
		goto L49
	} else {
		goto L2306
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L126:
	;
	v8427 = F_palloc(m, int32(16))
	mBase = m.M
	v8428 = m.ExcPending
	if v8428 != 0 {
		goto L49
	} else {
		goto L2292
	}
L127:
	;
	v8050 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v8052 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v8053 = m.G0
	v8055 = v8053 - int32(16)
	m.G0 = v8055
	v8057 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8055)+15)) = uint8(v8057)
	v8060 = F_palloc0(m, int32(36))
	mBase = m.M
	v8061 = m.ExcPending
	if v8061 != 0 {
		goto L49
	} else {
		goto L2167
	}
L128:
	;
	v7989 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v7992 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v7993 = int32(0)
	if v7992 < v7993 {
		v8038 = v7993
		goto L2155
	} else {
		goto L2156
	}
L129:
	;
	v7909 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v7914 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v7916 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_read_into_target(m, v28+int32(256), int32(0), v7914, v7916, l1)
	mBase = m.M
	v7918 = m.ExcPending
	if v7918 != 0 {
		goto L49
	} else {
		goto L2137
	}
L130:
	;
	v7559 = F_palloc0(m, int32(36))
	mBase = m.M
	v7560 = m.ExcPending
	if v7560 != 0 {
		goto L49
	} else {
		goto L2053
	}
L131:
	;
	v7381 = int32(1)
	v7390 = F_read_sql_construct(m, int32(332), int32(381), int32(59), int32(_a_F_plpgsql_yyparse_18), int32(2), v7381, v7381, int32(0), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7391 = m.ExcPending
	if v7391 != 0 {
		goto L49
	} else {
		goto L2020
	}
L132:
	;
	v7352 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v7354 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v7355 = F_plpgsql_yylex(m, v7352, v7354, l1)
	mBase = m.M
	v7356 = m.ExcPending
	if v7356 != 0 {
		goto L49
	} else {
		goto L2014
	}
L133:
	;
	v7328 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v7330 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v7331 = F_plpgsql_yylex(m, v7328, v7330, l1)
	mBase = m.M
	v7332 = m.ExcPending
	if v7332 != 0 {
		goto L49
	} else {
		goto L2008
	}
L134:
	;
	v7318 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7324 = F_make_execsql_stmt(m, int32(337), v7318, int32(0), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7325 = m.ExcPending
	if v7325 != 0 {
		goto L49
	} else {
		goto L2007
	}
L135:
	;
	v7308 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7314 = F_make_execsql_stmt(m, int32(331), v7308, int32(0), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7315 = m.ExcPending
	if v7315 != 0 {
		goto L49
	} else {
		goto L2006
	}
L136:
	;
	v7298 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7304 = F_make_execsql_stmt(m, int32(328), v7298, int32(0), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7305 = m.ExcPending
	if v7305 != 0 {
		goto L49
	} else {
		goto L2005
	}
L137:
	;
	v7287 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7287
	v7291 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v7291
	v7295 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+280)) = v7295
	v9817 = v233
	goto L13
L138:
	;
	v7190 = F_palloc(m, int32(20))
	mBase = m.M
	v7191 = m.ExcPending
	if v7191 != 0 {
		goto L49
	} else {
		goto L1986
	}
L139:
	;
	v6071 = F_palloc(m, int32(32))
	mBase = m.M
	v6072 = m.ExcPending
	if v6072 != 0 {
		goto L49
	} else {
		goto L1678
	}
L140:
	;
	v5401 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v5402 = m.ExcPending
	if v5402 != 0 {
		goto L49
	} else {
		goto L1503
	}
L141:
	;
	v5395 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v5395)
	v9817 = v233
	goto L13
L142:
	;
	v5393 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v5393)
	v9817 = v233
	goto L13
L143:
	;
	v5260 = F_palloc0(m, int32(24))
	mBase = m.M
	v5261 = m.ExcPending
	if v5261 != 0 {
		goto L49
	} else {
		goto L1451
	}
L144:
	;
	v5257 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5257
	v9817 = v233
	goto L13
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L146:
	;
	v5142 = F_palloc0(m, int32(32))
	mBase = m.M
	v5143 = m.ExcPending
	if v5143 != 0 {
		goto L49
	} else {
		goto L1424
	}
L147:
	;
	v5078 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5078
	v5080 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5081 = int32(0)
	if v5080 < v5081 {
		v5126 = v5081
		goto L1409
	} else {
		goto L1410
	}
L148:
	;
	v4994 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v4994 != 0 {
		goto L1384
	} else {
		goto L1385
	}
L149:
	;
	v4761 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4762 = m.ExcPending
	if v4762 != 0 {
		goto L49
	} else {
		goto L1337
	}
L150:
	;
	v4675 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4676 = *(*int32)(unsafe.Add(mBase, uint32(v4675)))
	v4679 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4680 = int32(0)
	if v4679 < v4680 {
		v4725 = v4680
		goto L1314
	} else {
		goto L1315
	}
L151:
	;
	v4584 = F_palloc0(m, int32(24))
	mBase = m.M
	v4585 = m.ExcPending
	if v4585 != 0 {
		goto L49
	} else {
		goto L1291
	}
L152:
	;
	v4498 = F_palloc0(m, int32(20))
	mBase = m.M
	v4499 = m.ExcPending
	if v4499 != 0 {
		goto L49
	} else {
		goto L1269
	}
L153:
	;
	v4485 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v4485 != 0 {
		goto L1265
	} else {
		goto L1266
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L155:
	;
	v4423 = F_palloc(m, int32(12))
	mBase = m.M
	v4424 = m.ExcPending
	if v4424 != 0 {
		goto L49
	} else {
		goto L1251
	}
L156:
	;
	v4413 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+184)) = v4413
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = v4413
	v4419 = F_list_make1_impl(m, int32(1), v28+int32(184))
	mBase = m.M
	v4420 = m.ExcPending
	if v4420 != 0 {
		goto L49
	} else {
		goto L1250
	}
L157:
	;
	v4408 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4409 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v4410 = F_lappend(m, v4408, v4409)
	mBase = m.M
	v4411 = m.ExcPending
	if v4411 != 0 {
		goto L49
	} else {
		goto L1249
	}
L158:
	;
	v4377 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v4379 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v4380 = F_plpgsql_yylex(m, v4377, v4379, l1)
	mBase = m.M
	v4381 = m.ExcPending
	if v4381 != 0 {
		goto L49
	} else {
		goto L1242
	}
L159:
	;
	v4164 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(24))))
	v4165 = int32(80)
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v148-v4165)))
	v4170 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	v4173 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	v4174 = m.G0
	v4176 = v4174 - v4165
	m.G0 = v4176
	v4179 = F_palloc(m, int32(32))
	mBase = m.M
	v4180 = m.ExcPending
	if v4180 != 0 {
		goto L49
	} else {
		goto L1209
	}
L160:
	;
	v4160 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4160
	v9817 = v233
	goto L13
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L162:
	;
	v4093 = F_palloc0(m, int32(12))
	mBase = m.M
	v4094 = m.ExcPending
	if v4094 != 0 {
		goto L49
	} else {
		goto L1194
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L164:
	;
	v4011 = F_palloc0(m, int32(28))
	mBase = m.M
	v4012 = m.ExcPending
	if v4012 != 0 {
		goto L49
	} else {
		goto L1180
	}
L165:
	;
	v3994 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3995 = *(*int32)(unsafe.Add(mBase, uint32(v3994)))
	if base.Ui32(v3995-int32(1)) < base.Ui32(int32(2)) {
		goto L33
	} else {
		goto L1176
	}
L166:
	;
	v3569 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v3570 = m.ExcPending
	if v3570 != 0 {
		goto L49
	} else {
		goto L1062
	}
L167:
	;
	v3555 = F_palloc(m, int32(8))
	mBase = m.M
	v3556 = m.ExcPending
	if v3556 != 0 {
		goto L49
	} else {
		goto L1034
	}
L168:
	;
	v3545 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+172)) = v3545
	*(*int32)(unsafe.Add(mBase, uint32(v28)+252)) = v3545
	v3551 = F_list_make1_impl(m, int32(1), v28+int32(172))
	mBase = m.M
	v3552 = m.ExcPending
	if v3552 != 0 {
		goto L49
	} else {
		goto L1033
	}
L169:
	;
	v3540 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3542 = F_lappend(m, v3540, v3541)
	mBase = m.M
	v3543 = m.ExcPending
	if v3543 != 0 {
		goto L49
	} else {
		goto L1032
	}
L170:
	;
	v3536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v3536)
	v9817 = v233
	goto L13
L171:
	;
	v3534 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v3534)
	v9817 = v233
	goto L13
L172:
	;
	v3532 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v3532)
	v9817 = v233
	goto L13
L173:
	;
	v3304 = F_palloc0(m, int32(20))
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L49
	} else {
		goto L980
	}
L174:
	;
	v3191 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v3191 == int32(0) {
		goto L955
	} else {
		goto L956
	}
L175:
	;
	v3113 = F_palloc0(m, int32(24))
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L49
	} else {
		goto L939
	}
L176:
	;
	v3035 = F_palloc0(m, int32(24))
	mBase = m.M
	v3036 = m.ExcPending
	if v3036 != 0 {
		goto L49
	} else {
		goto L923
	}
L177:
	;
	v2932 = F_palloc0(m, int32(16))
	mBase = m.M
	v2933 = m.ExcPending
	if v2933 != 0 {
		goto L49
	} else {
		goto L903
	}
L178:
	;
	v2929 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2929
	v9817 = v233
	goto L13
L179:
	;
	v2927 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2927
	v9817 = v233
	goto L13
L180:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2925
	v9817 = v233
	goto L13
L181:
	;
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2923
	v9817 = v233
	goto L13
L182:
	;
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2921
	v9817 = v233
	goto L13
L183:
	;
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2919
	v9817 = v233
	goto L13
L184:
	;
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2917
	v9817 = v233
	goto L13
L185:
	;
	v2915 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2915
	v9817 = v233
	goto L13
L186:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2913
	v9817 = v233
	goto L13
L187:
	;
	v2911 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2911
	v9817 = v233
	goto L13
L188:
	;
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2909
	v9817 = v233
	goto L13
L189:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2907
	v9817 = v233
	goto L13
L190:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2905
	v9817 = v233
	goto L13
L191:
	;
	v2903 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2903
	v9817 = v233
	goto L13
L192:
	;
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2901
	v9817 = v233
	goto L13
L193:
	;
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2899
	v9817 = v233
	goto L13
L194:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2897
	v9817 = v233
	goto L13
L195:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2895
	v9817 = v233
	goto L13
L196:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2893
	v9817 = v233
	goto L13
L197:
	;
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2891
	v9817 = v233
	goto L13
L198:
	;
	v2889 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2889
	v9817 = v233
	goto L13
L199:
	;
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2887
	v9817 = v233
	goto L13
L200:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2885
	v9817 = v233
	goto L13
L201:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2883
	v9817 = v233
	goto L13
L202:
	;
	v2873 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v2874 == int32(0) {
		goto L899
	} else {
		goto L900
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L204:
	;
	v2854 = int32(0)
	v2858 = int32(1)
	v2866 = F_read_sql_construct(m, int32(59), v2854, v2854, int32(_a_F_plpgsql_yyparse_16), int32(2), v2858, v2858, v2854, v2854, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v2867 = m.ExcPending
	if v2867 != 0 {
		goto L49
	} else {
		goto L898
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L206:
	;
	v2849 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2849)
	v9817 = v233
	goto L13
L207:
	;
	v2847 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2847)
	v9817 = v233
	goto L13
L208:
	;
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v2844 = F_get_collation_oid(m, v2842, int32(0))
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L49
	} else {
		goto L897
	}
L209:
	;
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v2827 = F_pstrdup(m, v2826)
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L49
	} else {
		goto L893
	}
L210:
	;
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v2813 = F_makeString(m, v2812)
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L49
	} else {
		goto L890
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L212:
	;
	v2262 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v2264 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v2265 = m.G0
	v2267 = v2265 - int32(48)
	m.G0 = v2267
	if v233 == int32(-2) {
		goto L699
	} else {
		goto L700
	}
L213:
	;
	v2259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2259)
	v9817 = v233
	goto L13
L214:
	;
	v2257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+272)) = uint8(v2257)
	v9817 = v233
	goto L13
L215:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v1857 = F_pstrdup(m, v1856)
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L49
	} else {
		goto L601
	}
L216:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1457
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v1460 = int32(0)
	if v1459 < v1460 {
		v1505 = v1460
		goto L505
	} else {
		goto L506
	}
L217:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1133 == int32(0) {
		goto L35
	} else {
		goto L427
	}
L218:
	;
	v979 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v979 == int32(0) {
		goto L394
	} else {
		goto L395
	}
L219:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v824 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L220:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v820 = F_plpgsql_build_variable(m, v814, v817, v818, int32(1))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L49
	} else {
		goto L352
	}
L221:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v809 = F_lappend(m, v807, v808)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L49
	} else {
		goto L351
	}
L222:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v796
	*(*int32)(unsafe.Add(mBase, uint32(v28)+268)) = v796
	v802 = F_list_make1_impl(m, int32(1), v28+int32(44))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L49
	} else {
		goto L350
	}
L223:
	;
	v617 = F_palloc0(m, int32(40))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L49
	} else {
		goto L320
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L225:
	;
	v611 = F_read_sql_stmt(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L49
	} else {
		goto L319
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(2)
	v9817 = v233
	goto L13
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(4)
	v9817 = v233
	goto L13
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L229:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	if v563 != 0 {
		goto L308
	} else {
		goto L309
	}
L230:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	F_plpgsql_ns_push(m, v556, int32(2))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L49
	} else {
		goto L306
	}
L231:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	F_plpgsql_ns_additem(m, v547, v548, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L49
	} else {
		goto L305
	}
L232:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	if v496 != 0 {
		goto L294
	} else {
		goto L295
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L49
	} else {
		goto L289
	}
L234:
	;
	v468 = F_plpgsql_add_initdatums(m, int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L49
	} else {
		goto L288
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14])) = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v462
	v464 = F_plpgsql_add_initdatums(m, v28+int32(280))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L49
	} else {
		goto L287
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14])) = int32(0)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+276)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v453
	v9817 = v233
	goto L13
L237:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14])) = int32(0)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+276)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v444
	v9817 = v233
	goto L13
L238:
	;
	v342 = F_palloc0(m, int32(32))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L49
	} else {
		goto L265
	}
L239:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v338 = F_pstrdup(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L49
	} else {
		goto L264
	}
L240:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v335
	v9817 = v233
	goto L13
L241:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+488)) = int32(2)
	v9817 = v233
	goto L13
L242:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+488)) = int32(1)
	v9817 = v233
	goto L13
L243:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+488)) = int32(0)
	v9817 = v233
	goto L13
L244:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v265 != int32(111) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	v262 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[16])) = uint8(v262)
	v9817 = v233
	goto L13
L246:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v259
	v9817 = v233
	goto L13
L247:
	;
	v276 = int32(_a_F_plpgsql_yyparse_19)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[17])))
	if base.B2i32(v279 == int32(0))|base.B2i32(v279 != v282) != 0 {
		v300 = v279
		v301 = v282
		goto L252
	} else {
		goto L253
	}
L248:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)))
	if v268 != int32(110) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+2)))
	if v271 != 0 {
		goto L247
	} else {
		goto L250
	}
L250:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v273)+492)) = uint8(v274)
	v9817 = v233
	goto L13
L251:
	;
	if v300-v301 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L252:
	;
	goto L251
L253:
	;
	v285 = v264
	v286 = v276
	goto L254
L254:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	if v290 == int32(0) {
		v300 = v290
		v301 = v289
		goto L252
	} else {
		goto L256
	}
L255:
	;
	v300 = v290
	v301 = v289
	goto L252
L256:
	;
	v293 = int32(1)
	if v290 == v289 {
		v285 = v285 + v293
		v286 = v286 + v293
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+492)) = uint8(v307)
	v9817 = v233
	goto L13
L259:
	;
	goto L260
L260:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L49
	} else {
		goto L261
	}
L261:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v313
	F_errmsg_internal(m, int32(_a_F_plpgsql_yyparse_20), v28)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L49
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(396), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L49
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v338
	v9817 = v233
	goto L13
L265:
	;
	v344 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v344
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(16))))
	if v348 < v344 {
		v394 = v344
		goto L267
	} else {
		goto L268
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+4)) = v394
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
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
		goto L49
	} else {
		goto L279
	}
L267:
	;
	goto L266
L268:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+56))
	if v355 == int32(0) {
		v394 = v344
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v358 = v348 + v355
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v354)+184))
	if base.Ui32(v359) <= base.Ui32(v358) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v354)+192))
	if base.B2i32(v368 == int32(0))|base.B2i32(base.Ui32(v358) <= base.Ui32(v368)) != 0 {
		v394 = v369
		goto L267
	} else {
		goto L274
	}
L271:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v354)+188))
	v368 = v361
	goto L270
L272:
	;
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v354)+184)) = v355
	v366 = F_strchr(m, v355, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v354)+188)) = v366
	v368 = v366
	goto L270
L274:
	;
	v374 = v368
	v377 = v369
	goto L275
L275:
	;
	v379 = int32(1)
	v380 = v377 + v379
	*(*int32)(unsafe.Add(mBase, uint32(v354)+192)) = v380
	v383 = v374 + v379
	*(*int32)(unsafe.Add(mBase, uint32(v354)+184)) = v383
	v386 = F_strchr(m, v383, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v354)+188)) = v386
	if v386 == int32(0) {
		v394 = v380
		goto L267
	} else {
		goto L277
	}
L276:
	;
	v394 = v380
	goto L267
L277:
	;
	if base.Ui32(v386) < base.Ui32(v358) {
		v374 = v386
		v377 = v380
		goto L275
	} else {
		goto L278
	}
L278:
	;
	goto L276
L279:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	if v432 != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v342
	v9817 = v233
	goto L13
L281:
	;
	v433 = v431
	goto L284
L282:
	;
	v436 = v431
	goto L283
L283:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = v438
	goto L280
L284:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+8))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	if v435 != 0 {
		v433 = v434
		goto L284
	} else {
		goto L286
	}
L285:
	;
	v436 = v434
	goto L283
L286:
	;
	goto L285
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v464
	v9817 = v233
	goto L13
L288:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[14])) = int32(1)
	v9817 = v233
	goto L13
L289:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L49
	} else {
		goto L290
	}
L290:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_23), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L49
	} else {
		goto L291
	}
L291:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v487 = F_plpgsql_scanner_errposition(m, v486, l1)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L49
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(502), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L49
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	v498 = v148 - int32(48)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+16))
	if v500 == int32(0) {
		goto L39
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(80))))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(76))))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	v516 = F_plpgsql_build_variable(m, v508, v511, v514, int32(1))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L49
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499)+16)) = v496
	goto L296
L298:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+int32(-64)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v516)+16)) = uint8(v520)
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148-int32(16)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v516)+17)) = uint8(v524)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v516)+20)) = v526
	if base.B2i32(v526 == int32(0))&base.B2i32(v524 == int32(1)) != 0 {
		goto L38
	} else {
		goto L299
	}
L299:
	;
	if v526 == int32(0) {
		v9817 = v233
		goto L13
	} else {
		goto L300
	}
L300:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	if v535 != 0 {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+20)) = uint8(v540)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+16)) = v541
	v9817 = v233
	goto L13
L302:
	;
	v540 = int32(0)
	v541 = int32(-1)
	goto L301
L303:
	;
	goto L304
L304:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v540 = int32(1)
	v541 = v539
	goto L301
L305:
	;
	v9817 = v233
	goto L13
L306:
	;
	v9817 = v233
	goto L13
L307:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(96))))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(92))))
	v579 = int32(0)
	v581 = F_plpgsql_build_datatype(m, int32(1790), int32(-1), v579, v579)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L49
	} else {
		goto L314
	}
L308:
	;
	v564 = v562
	goto L311
L309:
	;
	v567 = v562
	goto L310
L310:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = v569
	goto L307
L311:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+8))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	if v566 != 0 {
		v564 = v565
		goto L311
	} else {
		goto L313
	}
L312:
	;
	v567 = v565
	goto L310
L313:
	;
	goto L312
L314:
	;
	v584 = F_plpgsql_build_variable(m, v573, v576, v581, int32(1))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L49
	} else {
		goto L315
	}
L315:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+28)) = v586
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	if v590 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	v593 = v591
	goto L318
L317:
	;
	v593 = int32(-1)
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584)+32)) = v593
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(80))))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+36)) = v597 | int32(256)
	v9817 = v233
	goto L13
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v611
	v9817 = v233
	goto L13
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+8)) = int32(_a_F_plpgsql_yyparse_24)
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = int32(1)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v626 = int32(0)
	if v625 < v626 {
		v671 = v626
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v673 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v617)+24)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v617)+12)) = v671
	v678 = v148 - int32(16)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	if v679 != 0 {
		goto L334
	} else {
		goto L335
	}
L322:
	;
	goto L321
L323:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+56))
	if v632 == int32(0) {
		v671 = v626
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v635 = v625 + v632
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v631)+184))
	if base.Ui32(v636) <= base.Ui32(v635) {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v631)+192))
	if base.B2i32(v645 == int32(0))|base.B2i32(base.Ui32(v635) <= base.Ui32(v645)) != 0 {
		v671 = v646
		goto L322
	} else {
		goto L329
	}
L326:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v631)+188))
	v645 = v638
	goto L325
L327:
	;
	goto L328
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v631)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v631)+184)) = v632
	v643 = F_strchr(m, v632, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v631)+188)) = v643
	v645 = v643
	goto L325
L329:
	;
	v651 = v645
	v654 = v646
	goto L330
L330:
	;
	v656 = int32(1)
	v657 = v654 + v656
	*(*int32)(unsafe.Add(mBase, uint32(v631)+192)) = v657
	v660 = v651 + v656
	*(*int32)(unsafe.Add(mBase, uint32(v631)+184)) = v660
	v663 = F_strchr(m, v660, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v631)+188)) = v663
	if v663 == int32(0) {
		v671 = v657
		goto L322
	} else {
		goto L332
	}
L331:
	;
	v671 = v657
	goto L322
L332:
	;
	if base.Ui32(v663) < base.Ui32(v635) {
		v651 = v663
		v654 = v657
		goto L330
	} else {
		goto L333
	}
L333:
	;
	goto L331
L334:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)+4))
	v681 = v680
	goto L336
L335:
	;
	v681 = v673
	goto L336
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+28)) = v681
	v684 = F_palloc_mul(m, int32(4), v681)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L49
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+32)) = v684
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v617)+28))
	v689 = F_palloc_mul(m, int32(4), v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L49
	} else {
		goto L338
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v617)+36)) = v689
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	if v692 != 0 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
	if int32(0) < v693 {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	v790 = int32(0)
	goto L341
L341:
	;
	F_list_free(m, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L49
	} else {
		goto L348
	}
L342:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v617)+32))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v692)+12))
	v703 = int32(0)
	goto L345
L343:
	;
	goto L344
L344:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v790 = v763
	goto L341
L345:
	;
	v725 = v703 << (uint(int32(2)) % 32)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v725+v697)))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v696+v725))) = v729
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v728)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v725+v689))) = v732
	v735 = v703 + int32(1)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
	if v735 < v736 {
		v703 = v735
		goto L345
	} else {
		goto L347
	}
L346:
	;
	goto L344
L347:
	;
	goto L346
L348:
	;
	F_plpgsql_adddatum(m, v617)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L49
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v617
	v9817 = v233
	goto L13
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v802
	v9817 = v233
	goto L13
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v809
	v9817 = v233
	goto L13
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v820
	v9817 = v233
	goto L13
L353:
	;
	if v964 == int32(0) {
		goto L37
	} else {
		goto L389
	}
L354:
	;
	goto L353
L355:
	;
	v964 = v851
	goto L354
L357:
	;
	v964 = int32(0)
	goto L354
L358:
	;
	v837 = v824
	goto L359
L359:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	if v846 != 0 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	goto L357
L361:
	;
	v851 = v837
	v853 = v846
	goto L364
L362:
	;
	v876 = v837
	goto L363
L363:
	;
	goto L371
L364:
	;
	v858 = F_strcmp(m, v851+int32(12), v826)
	mBase = m.M
	if v858|base.B2i32(v853 == int32(1))&int32(0) == int32(0) {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v876 = v870
	goto L363
L366:
	;
	goto L355
L367:
	;
	goto L368
L368:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v851)+8))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	if v871 != 0 {
		v851 = v870
		v853 = v871
		goto L364
	} else {
		goto L370
	}
L370:
	;
	goto L365
L371:
	;
	goto L386
L386:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v876)+8))
	if v922 != 0 {
		v837 = v922
		goto L359
	} else {
		goto L387
	}
L387:
	;
	goto L360
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v964
	v9817 = v233
	goto L13
L390:
	;
	if v1119 == int32(0) {
		goto L36
	} else {
		goto L426
	}
L391:
	;
	goto L390
L392:
	;
	v1119 = v1006
	goto L391
L394:
	;
	v1119 = int32(0)
	goto L391
L395:
	;
	v992 = v979
	goto L396
L396:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	if v1001 != 0 {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	goto L394
L398:
	;
	v1006 = v992
	v1008 = v1001
	goto L401
L399:
	;
	v1031 = v992
	goto L400
L400:
	;
	goto L408
L401:
	;
	v1013 = F_strcmp(m, v1006+int32(12), v981)
	mBase = m.M
	if v1013|base.B2i32(v1008 == int32(1))&int32(0) == int32(0) {
		goto L403
	} else {
		goto L404
	}
L402:
	;
	v1031 = v1025
	goto L400
L403:
	;
	goto L392
L404:
	;
	goto L405
L405:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+8))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)))
	if v1026 != 0 {
		v1006 = v1025
		v1008 = v1026
		goto L401
	} else {
		goto L407
	}
L407:
	;
	goto L402
L408:
	;
	goto L423
L423:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+8))
	if v1077 != 0 {
		v992 = v1077
		goto L396
	} else {
		goto L424
	}
L424:
	;
	goto L397
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1119
	v9817 = v233
	goto L13
L427:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	switch v1136 - int32(2) {
	case 0:
		goto L430
	case 1:
		goto L429
	default:
		goto L35
	}
L428:
	;
	if v1453 == int32(0) {
		goto L35
	} else {
		goto L503
	}
L429:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1298)+12))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1299)))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1300)+4))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1299)+4))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1302)+4))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1299)+8))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+4))
	if v1296 == int32(0) {
		goto L471
	} else {
		goto L472
	}
L430:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+12))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+4))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+4))
	if v1140 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L431:
	;
	v1453 = v1294
	goto L428
L432:
	;
	v1294 = v1284
	goto L431
L433:
	;
	v1284 = v1171
	goto L432
L435:
	;
	v1284 = int32(0)
	goto L432
L436:
	;
	v1157 = v1140
	goto L437
L437:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1157)))
	if v1166 != 0 {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	goto L435
L439:
	;
	v1171 = v1157
	v1173 = v1166
	goto L442
L440:
	;
	v1196 = v1157
	goto L441
L441:
	;
	if v1147 == int32(0) {
		v1237 = v1196
		goto L449
	} else {
		goto L450
	}
L442:
	;
	v1178 = F_strcmp(m, v1171+int32(12), v1145)
	mBase = m.M
	v1181 = int32(0)
	if v1178|base.B2i32(v1173 == int32(1))&base.B2i32(v1147 != v1181) == v1181 {
		goto L444
	} else {
		goto L445
	}
L443:
	;
	v1196 = v1190
	goto L441
L444:
	;
	goto L433
L445:
	;
	goto L446
L446:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+8))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	if v1191 != 0 {
		v1171 = v1190
		v1173 = v1191
		goto L442
	} else {
		goto L448
	}
L448:
	;
	goto L443
L449:
	;
	goto L464
L450:
	;
	v1205 = F_strcmp(m, v1196+int32(12), v1145)
	mBase = m.M
	if v1205 != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1206 = v1196
	goto L453
L452:
	;
	v1206 = v1157
	goto L453
L453:
	;
	if v1205|base.B2i32(v1166 == int32(0)) != 0 {
		v1237 = v1206
		goto L449
	} else {
		goto L454
	}
L454:
	;
	v1210 = v1157
	v1217 = v1166
	goto L455
L455:
	;
	v1221 = F_strcmp(m, v1210+int32(12), v1147)
	mBase = m.M
	if v1221|int32(0)&base.B2i32(v1217 == int32(1)) == int32(0) {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	v1237 = v1231
	goto L449
L457:
	;
	goto L460
L458:
	;
	goto L459
L459:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+8))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1231)))
	if v1232 != 0 {
		v1210 = v1231
		v1217 = v1232
		goto L455
	} else {
		goto L463
	}
L460:
	;
	v1294 = v1210
	goto L431
L463:
	;
	goto L456
L464:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+8))
	if v1242 != 0 {
		v1157 = v1242
		goto L437
	} else {
		goto L465
	}
L465:
	;
	goto L438
L467:
	;
	v1453 = v1451
	goto L428
L468:
	;
	v1451 = v1441
	goto L467
L469:
	;
	v1441 = v1328
	goto L468
L471:
	;
	v1441 = int32(0)
	goto L468
L472:
	;
	v1314 = v1296
	goto L473
L473:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1314)))
	if v1323 != 0 {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	goto L471
L475:
	;
	v1328 = v1314
	v1330 = v1323
	goto L478
L476:
	;
	v1353 = v1314
	goto L477
L477:
	;
	if v1303 == int32(0) {
		v1394 = v1353
		goto L485
	} else {
		goto L486
	}
L478:
	;
	v1335 = F_strcmp(m, v1328+int32(12), v1301)
	mBase = m.M
	v1338 = int32(0)
	if v1335|base.B2i32(v1330 == int32(1))&base.B2i32(v1303 != v1338) == v1338 {
		goto L480
	} else {
		goto L481
	}
L479:
	;
	v1353 = v1347
	goto L477
L480:
	;
	goto L469
L481:
	;
	goto L482
L482:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+8))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1347)))
	if v1348 != 0 {
		v1328 = v1347
		v1330 = v1348
		goto L478
	} else {
		goto L484
	}
L484:
	;
	goto L479
L485:
	;
	goto L500
L486:
	;
	v1362 = F_strcmp(m, v1353+int32(12), v1301)
	mBase = m.M
	if v1362 != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v1363 = v1353
	goto L489
L488:
	;
	v1363 = v1314
	goto L489
L489:
	;
	if v1362|base.B2i32(v1323 == int32(0)) != 0 {
		v1394 = v1363
		goto L485
	} else {
		goto L490
	}
L490:
	;
	v1367 = v1314
	v1374 = v1323
	goto L491
L491:
	;
	v1378 = F_strcmp(m, v1367+int32(12), v1303)
	mBase = m.M
	if v1378|base.B2i32(v1305 != int32(0))&base.B2i32(v1374 == int32(1)) == int32(0) {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	v1394 = v1388
	goto L485
L493:
	;
	goto L496
L494:
	;
	goto L495
L495:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+8))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1388)))
	if v1389 != 0 {
		v1367 = v1388
		v1374 = v1389
		goto L491
	} else {
		goto L499
	}
L496:
	;
	v1451 = v1367
	goto L467
L499:
	;
	goto L492
L500:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+8))
	if v1399 != 0 {
		v1314 = v1399
		goto L473
	} else {
		goto L501
	}
L501:
	;
	goto L474
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1453
	v9817 = v233
	goto L13
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v1505
	v1509 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1509 == int32(0) {
		goto L521
	} else {
		goto L522
	}
L505:
	;
	goto L504
L506:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+56))
	if v1466 == int32(0) {
		v1505 = v1460
		goto L505
	} else {
		goto L507
	}
L507:
	;
	v1469 = v1459 + v1466
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+184))
	if base.Ui32(v1470) <= base.Ui32(v1469) {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+192))
	if base.B2i32(v1479 == int32(0))|base.B2i32(base.Ui32(v1469) <= base.Ui32(v1479)) != 0 {
		v1505 = v1480
		goto L505
	} else {
		goto L512
	}
L509:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+188))
	v1479 = v1472
	goto L508
L510:
	;
	goto L511
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+184)) = v1466
	v1477 = F_strchr(m, v1466, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+188)) = v1477
	v1479 = v1477
	goto L508
L512:
	;
	v1485 = v1479
	v1488 = v1480
	goto L513
L513:
	;
	v1490 = int32(1)
	v1491 = v1488 + v1490
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+192)) = v1491
	v1494 = v1485 + v1490
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+184)) = v1494
	v1497 = F_strchr(m, v1494, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1465)+188)) = v1497
	if v1497 == int32(0) {
		v1505 = v1491
		goto L505
	} else {
		goto L515
	}
L514:
	;
	v1505 = v1491
	goto L505
L515:
	;
	if base.Ui32(v1497) < base.Ui32(v1469) {
		v1485 = v1497
		v1488 = v1491
		goto L513
	} else {
		goto L516
	}
L516:
	;
	goto L514
L517:
	;
	if v1649 != 0 {
		goto L2
	} else {
		goto L553
	}
L518:
	;
	goto L517
L519:
	;
	v1649 = v1536
	goto L518
L521:
	;
	v1649 = int32(0)
	goto L518
L522:
	;
	goto L523
L523:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1509)))
	if v1531 != 0 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v1536 = v1509
	v1538 = v1531
	goto L528
L526:
	;
	goto L527
L527:
	;
	goto L535
L528:
	;
	v1543 = F_strcmp(m, v1536+int32(12), v1511)
	mBase = m.M
	if v1543|base.B2i32(v1538 == int32(1))&int32(0) == int32(0) {
		goto L530
	} else {
		goto L531
	}
L529:
	;
	goto L527
L530:
	;
	goto L519
L531:
	;
	goto L532
L532:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+8))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1555)))
	if v1556 != 0 {
		v1536 = v1555
		v1538 = v1556
		goto L528
	} else {
		goto L534
	}
L534:
	;
	goto L529
L535:
	;
	goto L521
L553:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v1662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661)+496)))
	if v1662&int32(2) == int32(0) {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1661)+500)))
	if v1667&int32(2) == int32(0) {
		v9817 = v233
		goto L13
	} else {
		goto L557
	}
L555:
	;
	goto L556
L556:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1673 == int32(0) {
		goto L562
	} else {
		goto L563
	}
L557:
	;
	goto L556
L558:
	;
	if v1813 == int32(0) {
		v9817 = v233
		goto L13
	} else {
		goto L594
	}
L559:
	;
	goto L558
L560:
	;
	v1813 = v1700
	goto L559
L562:
	;
	v1813 = int32(0)
	goto L559
L563:
	;
	v1686 = v1673
	goto L564
L564:
	;
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1686)))
	if v1695 != 0 {
		goto L566
	} else {
		goto L567
	}
L565:
	;
	goto L562
L566:
	;
	v1700 = v1686
	v1702 = v1695
	goto L569
L567:
	;
	v1725 = v1686
	goto L568
L568:
	;
	goto L576
L569:
	;
	v1707 = F_strcmp(m, v1700+int32(12), v1675)
	mBase = m.M
	if v1707|base.B2i32(v1702 == int32(1))&int32(0) == int32(0) {
		goto L571
	} else {
		goto L572
	}
L570:
	;
	v1725 = v1719
	goto L568
L571:
	;
	goto L560
L572:
	;
	goto L573
L573:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1700)+8))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)))
	if v1720 != 0 {
		v1700 = v1719
		v1702 = v1720
		goto L569
	} else {
		goto L575
	}
L575:
	;
	goto L570
L576:
	;
	goto L591
L591:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1725)+8))
	if v1771 != 0 {
		v1686 = v1771
		goto L564
	} else {
		goto L592
	}
L592:
	;
	goto L565
L594:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+500))
	v1834 = F_errstart(m, v1828&int32(2)+int32(19), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L49
	} else {
		goto L595
	}
L595:
	;
	if v1834 == int32(0) {
		v9817 = v233
		goto L13
	} else {
		goto L596
	}
L596:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L49
	} else {
		goto L597
	}
L597:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v1841
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_25), v28+int32(96))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L49
	} else {
		goto L598
	}
L598:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v1849 = F_plpgsql_scanner_errposition(m, v1848, l1)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L49
	} else {
		goto L599
	}
L599:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(737), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L49
	} else {
		goto L600
	}
L600:
	;
	v9817 = v233
	goto L13
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v1857
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v1861 = int32(0)
	if v1860 < v1861 {
		v1906 = v1861
		goto L603
	} else {
		goto L604
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v1906
	v1910 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v1910 == int32(0) {
		goto L619
	} else {
		goto L620
	}
L603:
	;
	goto L602
L604:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1866)+56))
	if v1867 == int32(0) {
		v1906 = v1861
		goto L603
	} else {
		goto L605
	}
L605:
	;
	v1870 = v1860 + v1867
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1866)+184))
	if base.Ui32(v1871) <= base.Ui32(v1870) {
		goto L607
	} else {
		goto L608
	}
L606:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1866)+192))
	if base.B2i32(v1880 == int32(0))|base.B2i32(base.Ui32(v1870) <= base.Ui32(v1880)) != 0 {
		v1906 = v1881
		goto L603
	} else {
		goto L610
	}
L607:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1866)+188))
	v1880 = v1873
	goto L606
L608:
	;
	goto L609
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1866)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1866)+184)) = v1867
	v1878 = F_strchr(m, v1867, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1866)+188)) = v1878
	v1880 = v1878
	goto L606
L610:
	;
	v1886 = v1880
	v1889 = v1881
	goto L611
L611:
	;
	v1891 = int32(1)
	v1892 = v1889 + v1891
	*(*int32)(unsafe.Add(mBase, uint32(v1866)+192)) = v1892
	v1895 = v1886 + v1891
	*(*int32)(unsafe.Add(mBase, uint32(v1866)+184)) = v1895
	v1898 = F_strchr(m, v1895, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1866)+188)) = v1898
	if v1898 == int32(0) {
		v1906 = v1892
		goto L603
	} else {
		goto L613
	}
L612:
	;
	v1906 = v1892
	goto L603
L613:
	;
	if base.Ui32(v1898) < base.Ui32(v1870) {
		v1886 = v1898
		v1889 = v1892
		goto L611
	} else {
		goto L614
	}
L614:
	;
	goto L612
L615:
	;
	if v2050 != 0 {
		goto L2
	} else {
		goto L651
	}
L616:
	;
	goto L615
L617:
	;
	v2050 = v1937
	goto L616
L619:
	;
	v2050 = int32(0)
	goto L616
L620:
	;
	goto L621
L621:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1910)))
	if v1932 != 0 {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	v1937 = v1910
	v1939 = v1932
	goto L626
L624:
	;
	goto L625
L625:
	;
	goto L633
L626:
	;
	v1944 = F_strcmp(m, v1937+int32(12), v1912)
	mBase = m.M
	if v1944|base.B2i32(v1939 == int32(1))&int32(0) == int32(0) {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	goto L625
L628:
	;
	goto L617
L629:
	;
	goto L630
L630:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1937)+8))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1956)))
	if v1957 != 0 {
		v1937 = v1956
		v1939 = v1957
		goto L626
	} else {
		goto L632
	}
L632:
	;
	goto L627
L633:
	;
	goto L619
L651:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2062)+496)))
	if v2063&int32(2) == int32(0) {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v2068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2062)+500)))
	if v2068&int32(2) == int32(0) {
		v9817 = v233
		goto L13
	} else {
		goto L655
	}
L653:
	;
	goto L654
L654:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v2074 == int32(0) {
		goto L660
	} else {
		goto L661
	}
L655:
	;
	goto L654
L656:
	;
	if v2214 == int32(0) {
		v9817 = v233
		goto L13
	} else {
		goto L692
	}
L657:
	;
	goto L656
L658:
	;
	v2214 = v2101
	goto L657
L660:
	;
	v2214 = int32(0)
	goto L657
L661:
	;
	v2087 = v2074
	goto L662
L662:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2087)))
	if v2096 != 0 {
		goto L664
	} else {
		goto L665
	}
L663:
	;
	goto L660
L664:
	;
	v2101 = v2087
	v2103 = v2096
	goto L667
L665:
	;
	v2126 = v2087
	goto L666
L666:
	;
	goto L674
L667:
	;
	v2108 = F_strcmp(m, v2101+int32(12), v2076)
	mBase = m.M
	if v2108|base.B2i32(v2103 == int32(1))&int32(0) == int32(0) {
		goto L669
	} else {
		goto L670
	}
L668:
	;
	v2126 = v2120
	goto L666
L669:
	;
	goto L658
L670:
	;
	goto L671
L671:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+8))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2120)))
	if v2121 != 0 {
		v2101 = v2120
		v2103 = v2121
		goto L667
	} else {
		goto L673
	}
L673:
	;
	goto L668
L674:
	;
	goto L689
L689:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2126)+8))
	if v2172 != 0 {
		v2087 = v2172
		goto L662
	} else {
		goto L690
	}
L690:
	;
	goto L663
L692:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+500))
	v2235 = F_errstart(m, v2229&int32(2)+int32(19), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L49
	} else {
		goto L693
	}
L693:
	;
	if v2235 == int32(0) {
		v9817 = v233
		goto L13
	} else {
		goto L694
	}
L694:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L49
	} else {
		goto L695
	}
L695:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v2242
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_25), v28+int32(112))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L49
	} else {
		goto L696
	}
L696:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v2250 = F_plpgsql_scanner_errposition(m, v2249, l1)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L49
	} else {
		goto L697
	}
L697:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(765), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L49
	} else {
		goto L698
	}
L698:
	;
	v9817 = v233
	goto L13
L699:
	;
	v2271 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L49
	} else {
		goto L702
	}
L700:
	;
	v2273 = v233
	goto L701
L701:
	;
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	if v2273 == int32(275) {
		goto L707
	} else {
		goto L708
	}
L702:
	;
	v2273 = v2271
	goto L701
L703:
	;
	m.G0 = v2267 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2790
	v9817 = int32(-2)
	goto L13
L704:
	;
	v2681 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		goto L49
	} else {
		goto L858
	}
L705:
	;
	v2557 = v2550
	v2565 = int32(0)
	goto L817
L706:
	;
	if v2547 != 0 {
		goto L704
	} else {
		goto L816
	}
L707:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v2278 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L49
	} else {
		goto L710
	}
L708:
	;
	goto L709
L709:
	;
	v2361 = int32(0)
	goto L742
L710:
	;
	if v2278 != int32(37) {
		v2550 = v2278
		goto L705
	} else {
		goto L711
	}
L711:
	;
	v2282 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L49
	} else {
		goto L712
	}
L712:
	;
	switch v2282 - int32(366) {
	case 0:
		goto L713
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v2550 = v2282
		goto L705
	case 12:
		goto L715
	default:
		goto L716
	}
L713:
	;
	v2356 = F_plpgsql_parse_wordrowtype(m, v2277)
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L49
	} else {
		goto L740
	}
L714:
	;
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v2324 == int32(0) {
		v2550 = v2286
		goto L705
	} else {
		goto L731
	}
L715:
	;
	v2322 = F_plpgsql_parse_wordtype(m, v2277)
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L49
	} else {
		goto L730
	}
L716:
	;
	v2286 = int32(277)
	if v2282 != v2286 {
		goto L717
	} else {
		goto L718
	}
L717:
	;
	v2550 = v2282
	goto L705
L718:
	;
	goto L719
L719:
	;
	v2289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2289 != 0 {
		v2550 = v2286
		goto L705
	} else {
		goto L720
	}
L720:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v2290 == int32(0) {
		goto L714
	} else {
		goto L721
	}
L721:
	;
	v2293 = int32(_a_F_plpgsql_yyparse_26)
	v2296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2290))))
	v2299 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[18])))
	if base.B2i32(v2296 == int32(0))|base.B2i32(v2296 != v2299) != 0 {
		v2317 = v2296
		v2318 = v2299
		goto L723
	} else {
		goto L724
	}
L722:
	;
	if v2317-v2318 != 0 {
		goto L714
	} else {
		goto L729
	}
L723:
	;
	goto L722
L724:
	;
	v2302 = v2290
	v2303 = v2293
	goto L725
L725:
	;
	v2306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303)+1)))
	v2307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2302)+1)))
	if v2307 == int32(0) {
		v2317 = v2307
		v2318 = v2306
		goto L723
	} else {
		goto L727
	}
L726:
	;
	v2317 = v2307
	v2318 = v2306
	goto L723
L727:
	;
	v2310 = int32(1)
	if v2307 == v2306 {
		v2302 = v2302 + v2310
		v2303 = v2303 + v2310
		goto L725
	} else {
		goto L728
	}
L728:
	;
	goto L726
L729:
	;
	goto L715
L730:
	;
	v2547 = v2322
	v2549 = v2282
	goto L706
L731:
	;
	v2327 = int32(_a_F_plpgsql_yyparse_27)
	v2330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2324))))
	v2333 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[19])))
	if base.B2i32(v2330 == int32(0))|base.B2i32(v2330 != v2333) != 0 {
		v2351 = v2330
		v2352 = v2333
		goto L733
	} else {
		goto L734
	}
L732:
	;
	if v2351-v2352 != 0 {
		v2550 = v2286
		goto L705
	} else {
		goto L739
	}
L733:
	;
	goto L732
L734:
	;
	v2336 = v2324
	v2337 = v2327
	goto L735
L735:
	;
	v2340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2337)+1)))
	v2341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2336)+1)))
	if v2341 == int32(0) {
		v2351 = v2341
		v2352 = v2340
		goto L733
	} else {
		goto L737
	}
L736:
	;
	v2351 = v2341
	v2352 = v2340
	goto L733
L737:
	;
	v2344 = int32(1)
	if v2341 == v2340 {
		v2336 = v2336 + v2344
		v2337 = v2337 + v2344
		goto L735
	} else {
		goto L738
	}
L738:
	;
	goto L736
L739:
	;
	goto L713
L740:
	;
	v2547 = v2356
	v2549 = v2282
	goto L706
L741:
	;
	if v2273 == v2367 {
		goto L748
	} else {
		goto L749
	}
L742:
	;
	v2367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2361<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[20]))))
	v2368 = base.B2i32(v2273 == v2367)
	if v2368 == int32(0) {
		goto L744
	} else {
		goto L745
	}
L743:
	;
	goto L741
L744:
	;
	v2372 = v2361 + int32(1)
	if v2372 != int32(85) {
		v2361 = v2372
		goto L742
	} else {
		goto L747
	}
L745:
	;
	goto L746
L746:
	;
	goto L743
L747:
	;
	goto L746
L748:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v2377 = F_pstrdup(m, v2376)
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L49
	} else {
		goto L751
	}
L749:
	;
	goto L750
L750:
	;
	if v2273 != int32(276) {
		v2550 = v2273
		goto L705
	} else {
		goto L783
	}
L751:
	;
	v2379 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L49
	} else {
		goto L752
	}
L752:
	;
	if v2379 != int32(37) {
		v2550 = v2379
		goto L705
	} else {
		goto L753
	}
L753:
	;
	v2383 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L49
	} else {
		goto L754
	}
L754:
	;
	switch v2383 - int32(366) {
	case 0:
		goto L755
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v2550 = v2383
		goto L705
	case 12:
		goto L757
	default:
		goto L758
	}
L755:
	;
	v2457 = F_plpgsql_parse_wordrowtype(m, v2377)
	mBase = m.M
	v2458 = m.ExcPending
	if v2458 != 0 {
		goto L49
	} else {
		goto L782
	}
L756:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v2425 == int32(0) {
		v2550 = v2387
		goto L705
	} else {
		goto L773
	}
L757:
	;
	v2423 = F_plpgsql_parse_wordtype(m, v2377)
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L49
	} else {
		goto L772
	}
L758:
	;
	v2387 = int32(277)
	if v2383 != v2387 {
		goto L759
	} else {
		goto L760
	}
L759:
	;
	v2550 = v2383
	goto L705
L760:
	;
	goto L761
L761:
	;
	v2390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2390 != 0 {
		v2550 = v2387
		goto L705
	} else {
		goto L762
	}
L762:
	;
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v2391 == int32(0) {
		goto L756
	} else {
		goto L763
	}
L763:
	;
	v2394 = int32(_a_F_plpgsql_yyparse_26)
	v2397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2391))))
	v2400 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[18])))
	if base.B2i32(v2397 == int32(0))|base.B2i32(v2397 != v2400) != 0 {
		v2418 = v2397
		v2419 = v2400
		goto L765
	} else {
		goto L766
	}
L764:
	;
	if v2418-v2419 != 0 {
		goto L756
	} else {
		goto L771
	}
L765:
	;
	goto L764
L766:
	;
	v2403 = v2391
	v2404 = v2394
	goto L767
L767:
	;
	v2407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2404)+1)))
	v2408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2403)+1)))
	if v2408 == int32(0) {
		v2418 = v2408
		v2419 = v2407
		goto L765
	} else {
		goto L769
	}
L768:
	;
	v2418 = v2408
	v2419 = v2407
	goto L765
L769:
	;
	v2411 = int32(1)
	if v2408 == v2407 {
		v2403 = v2403 + v2411
		v2404 = v2404 + v2411
		goto L767
	} else {
		goto L770
	}
L770:
	;
	goto L768
L771:
	;
	goto L757
L772:
	;
	v2547 = v2423
	v2549 = v2383
	goto L706
L773:
	;
	v2428 = int32(_a_F_plpgsql_yyparse_27)
	v2431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2425))))
	v2434 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[19])))
	if base.B2i32(v2431 == int32(0))|base.B2i32(v2431 != v2434) != 0 {
		v2452 = v2431
		v2453 = v2434
		goto L775
	} else {
		goto L776
	}
L774:
	;
	if v2452-v2453 != 0 {
		v2550 = v2387
		goto L705
	} else {
		goto L781
	}
L775:
	;
	goto L774
L776:
	;
	v2437 = v2425
	v2438 = v2428
	goto L777
L777:
	;
	v2441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2438)+1)))
	v2442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2437)+1)))
	if v2442 == int32(0) {
		v2452 = v2442
		v2453 = v2441
		goto L775
	} else {
		goto L779
	}
L778:
	;
	v2452 = v2442
	v2453 = v2441
	goto L775
L779:
	;
	v2445 = int32(1)
	if v2442 == v2441 {
		v2437 = v2437 + v2445
		v2438 = v2438 + v2445
		goto L777
	} else {
		goto L780
	}
L780:
	;
	goto L778
L781:
	;
	goto L755
L782:
	;
	v2547 = v2457
	v2549 = v2383
	goto L706
L783:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v2462 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L49
	} else {
		goto L784
	}
L784:
	;
	if v2462 != int32(37) {
		v2550 = v2462
		goto L705
	} else {
		goto L785
	}
L785:
	;
	v2466 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L49
	} else {
		goto L786
	}
L786:
	;
	switch v2466 - int32(366) {
	case 0:
		goto L788
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v2550 = v2466
		goto L705
	case 12:
		goto L790
	default:
		goto L791
	}
L787:
	;
	v2547 = v2544
	v2549 = v2466
	goto L706
L788:
	;
	v2540 = F_plpgsql_parse_cwordrowtype(m, v2461)
	mBase = m.M
	v2541 = m.ExcPending
	if v2541 != 0 {
		goto L49
	} else {
		goto L815
	}
L789:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v2508 == int32(0) {
		v2550 = v2470
		goto L705
	} else {
		goto L806
	}
L790:
	;
	v2506 = F_plpgsql_parse_cwordtype(m, v2461)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L49
	} else {
		goto L805
	}
L791:
	;
	v2470 = int32(277)
	if v2466 != v2470 {
		goto L792
	} else {
		goto L793
	}
L792:
	;
	v2550 = v2466
	goto L705
L793:
	;
	goto L794
L794:
	;
	v2473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2473 != 0 {
		v2550 = v2470
		goto L705
	} else {
		goto L795
	}
L795:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v2474 == int32(0) {
		goto L789
	} else {
		goto L796
	}
L796:
	;
	v2477 = int32(_a_F_plpgsql_yyparse_26)
	v2480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2474))))
	v2483 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[18])))
	if base.B2i32(v2480 == int32(0))|base.B2i32(v2480 != v2483) != 0 {
		v2501 = v2480
		v2502 = v2483
		goto L798
	} else {
		goto L799
	}
L797:
	;
	if v2501-v2502 != 0 {
		goto L789
	} else {
		goto L804
	}
L798:
	;
	goto L797
L799:
	;
	v2486 = v2474
	v2487 = v2477
	goto L800
L800:
	;
	v2490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2487)+1)))
	v2491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2486)+1)))
	if v2491 == int32(0) {
		v2501 = v2491
		v2502 = v2490
		goto L798
	} else {
		goto L802
	}
L801:
	;
	v2501 = v2491
	v2502 = v2490
	goto L798
L802:
	;
	v2494 = int32(1)
	if v2491 == v2490 {
		v2486 = v2486 + v2494
		v2487 = v2487 + v2494
		goto L800
	} else {
		goto L803
	}
L803:
	;
	goto L801
L804:
	;
	goto L790
L805:
	;
	v2544 = v2506
	goto L787
L806:
	;
	v2511 = int32(_a_F_plpgsql_yyparse_27)
	v2514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2508))))
	v2517 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[19])))
	if base.B2i32(v2514 == int32(0))|base.B2i32(v2514 != v2517) != 0 {
		v2535 = v2514
		v2536 = v2517
		goto L808
	} else {
		goto L809
	}
L807:
	;
	if v2535-v2536 != 0 {
		v2550 = v2470
		goto L705
	} else {
		goto L814
	}
L808:
	;
	goto L807
L809:
	;
	v2520 = v2508
	v2521 = v2511
	goto L810
L810:
	;
	v2524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2521)+1)))
	v2525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2520)+1)))
	if v2525 == int32(0) {
		v2535 = v2525
		v2536 = v2524
		goto L808
	} else {
		goto L812
	}
L811:
	;
	v2535 = v2525
	v2536 = v2524
	goto L808
L812:
	;
	v2528 = int32(1)
	if v2525 == v2524 {
		v2520 = v2520 + v2528
		v2521 = v2521 + v2528
		goto L810
	} else {
		goto L813
	}
L813:
	;
	goto L811
L814:
	;
	goto L788
L815:
	;
	v2544 = v2540
	goto L787
L816:
	;
	v2550 = v2549
	goto L705
L817:
	;
	if v2557 <= int32(292) {
		goto L824
	} else {
		goto L825
	}
L818:
	;
	v2621 = v2267 + int32(4)
	F_initStringInfo(m, v2621)
	mBase = m.M
	v2623 = m.ExcPending
	if v2623 != 0 {
		goto L49
	} else {
		goto L842
	}
L819:
	;
	goto L818
L820:
	;
	if base.B2i32(v2557 != int32(44))&base.B2i32(v2557 != int32(41))|v2565 == int32(0) {
		goto L819
	} else {
		goto L834
	}
L821:
	;
	F_plpgsql_yyerror(m, v2264, int32(0), l1, int32(_a_F_plpgsql_yyparse_28))
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L49
	} else {
		goto L833
	}
L822:
	;
	if v2557 != int32(343) {
		goto L820
	} else {
		goto L832
	}
L823:
	;
	if v2565 != 0 {
		goto L821
	} else {
		goto L830
	}
L824:
	;
	switch v2557 - int32(59) {
	case 0, 2:
		goto L819
	case 1:
		goto L820
	default:
		goto L827
	}
L825:
	;
	goto L826
L826:
	;
	switch v2557 - int32(293) {
	case 0, 13:
		goto L819
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L820
	default:
		goto L822
	}
L827:
	;
	if v2557 == int32(270) {
		goto L819
	} else {
		goto L828
	}
L828:
	;
	if v2557 == int32(0) {
		goto L823
	} else {
		goto L829
	}
L829:
	;
	goto L820
L830:
	;
	F_plpgsql_yyerror(m, v2264, int32(0), l1, int32(_a_F_plpgsql_yyparse_29))
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L49
	} else {
		goto L831
	}
L831:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L832:
	;
	goto L819
L833:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L834:
	;
	if v2557 == int32(41) {
		goto L835
	} else {
		goto L836
	}
L835:
	;
	v2613 = int32(-1)
	goto L837
L836:
	;
	v2613 = int32(0)
	goto L837
L837:
	;
	if v2557 == int32(40) {
		goto L838
	} else {
		goto L839
	}
L838:
	;
	v2616 = int32(1)
	goto L840
L839:
	;
	v2616 = v2613
	goto L840
L840:
	;
	v2618 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2619 = m.ExcPending
	if v2619 != 0 {
		goto L49
	} else {
		goto L841
	}
L841:
	;
	v2557 = v2618
	v2565 = v2616 + v2565
	goto L817
L842:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	F_plpgsql_append_source_text(m, v2621, v2274, v2624, l1)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L49
	} else {
		goto L843
	}
L843:
	;
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+4))
	v2628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2627))))
	if v2628 != 0 {
		goto L844
	} else {
		goto L845
	}
L844:
	;
	v2629 = int32(_a_F_plpgsql_yyparse_30)
	v2630 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[21]))
	v2633 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[21])) = v2633
	v2635 = int32(_a_F_plpgsql_yyparse_31)
	v2636 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[23])) = v2267 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+32)) = v2274
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+24)) = int32(_a_F_plpgsql_yyparse_32)
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+20)) = v2636
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+28)) = v2267 + int32(32)
	v2649 = int32(0)
	v2651 = F_typeStringToTypeName(m, v2627, v2649)
	mBase = m.M
	v2652 = m.ExcPending
	if v2652 != 0 {
		goto L49
	} else {
		goto L847
	}
L845:
	;
	goto L846
L846:
	;
	F_plpgsql_yyerror(m, v2264, int32(0), l1, int32(_a_F_plpgsql_yyparse_33))
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L49
	} else {
		goto L852
	}
L847:
	;
	F_typenameTypeIdAndMod(m, v2649, v2651, v2267+int32(44), v2267+int32(40))
	mBase = m.M
	v2658 = m.ExcPending
	if v2658 != 0 {
		goto L49
	} else {
		goto L848
	}
L848:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[21])) = v2630
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[23])) = v2662
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+44))
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+40))
	v2667 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v2667)+44))
	v2669 = F_plpgsql_build_datatype(m, v2664, v2665, v2668, v2651)
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L49
	} else {
		goto L849
	}
L849:
	;
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+4))
	F_pfree(m, v2671)
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L49
	} else {
		goto L850
	}
L850:
	;
	F_plpgsql_push_back_token(m, v2557, v2262, v2264, l1)
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L49
	} else {
		goto L851
	}
L851:
	;
	v2790 = v2669
	goto L703
L852:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L853:
	;
	F_plpgsql_yyerror(m, v2264, int32(0), l1, int32(_a_F_plpgsql_yyparse_34))
	mBase = m.M
	v2779 = m.ExcPending
	if v2779 != 0 {
		goto L49
	} else {
		goto L889
	}
L854:
	;
	F_plpgsql_push_back_token(m, int32(277), v2262, v2264, l1)
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L49
	} else {
		goto L888
	}
L855:
	;
	if v2720 == int32(91) {
		goto L870
	} else {
		goto L871
	}
L856:
	;
	v2718 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L49
	} else {
		goto L869
	}
L857:
	;
	v2685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v2685 != 0 {
		goto L854
	} else {
		goto L859
	}
L858:
	;
	switch v2681 - int32(277) {
	case 0:
		goto L857
	default:
		v2720 = v2681
		v2721 = int32(0)
		goto L855
	case 7:
		goto L856
	}
L859:
	;
	v2686 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v2686 == int32(0) {
		goto L854
	} else {
		goto L860
	}
L860:
	;
	v2689 = int32(_a_F_plpgsql_yyparse_35)
	v2692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2686))))
	v2695 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[24])))
	if base.B2i32(v2692 == int32(0))|base.B2i32(v2692 != v2695) != 0 {
		v2713 = v2692
		v2714 = v2695
		goto L862
	} else {
		goto L863
	}
L861:
	;
	if v2713-v2714 != 0 {
		goto L854
	} else {
		goto L868
	}
L862:
	;
	goto L861
L863:
	;
	v2698 = v2686
	v2699 = v2689
	goto L864
L864:
	;
	v2702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2699)+1)))
	v2703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2698)+1)))
	if v2703 == int32(0) {
		v2713 = v2703
		v2714 = v2702
		goto L862
	} else {
		goto L866
	}
L865:
	;
	v2713 = v2703
	v2714 = v2702
	goto L862
L866:
	;
	v2706 = int32(1)
	if v2703 == v2702 {
		v2698 = v2698 + v2706
		v2699 = v2699 + v2706
		goto L864
	} else {
		goto L867
	}
L867:
	;
	goto L865
L868:
	;
	goto L856
L869:
	;
	v2720 = v2718
	v2721 = int32(1)
	goto L855
L870:
	;
	goto L873
L871:
	;
	goto L872
L872:
	;
	F_plpgsql_push_back_token(m, v2720, v2262, v2264, l1)
	mBase = m.M
	v2767 = m.ExcPending
	if v2767 != 0 {
		goto L49
	} else {
		goto L885
	}
L873:
	;
	v2749 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2750 = m.ExcPending
	if v2750 != 0 {
		goto L49
	} else {
		goto L875
	}
L874:
	;
	F_plpgsql_push_back_token(m, v2758, v2262, v2264, l1)
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L49
	} else {
		goto L883
	}
L875:
	;
	if v2749 == int32(266) {
		goto L876
	} else {
		goto L877
	}
L876:
	;
	v2753 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L49
	} else {
		goto L879
	}
L877:
	;
	v2755 = v2749
	goto L878
L878:
	;
	if v2755 != int32(93) {
		goto L853
	} else {
		goto L880
	}
L879:
	;
	v2755 = v2753
	goto L878
L880:
	;
	v2758 = F_plpgsql_yylex(m, v2262, v2264, l1)
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L49
	} else {
		goto L881
	}
L881:
	;
	if v2758 == int32(91) {
		goto L873
	} else {
		goto L882
	}
L882:
	;
	goto L874
L883:
	;
	v2764 = F_plpgsql_build_datatype_arrayof(m, v2547)
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L49
	} else {
		goto L884
	}
L884:
	;
	v2790 = v2764
	goto L703
L885:
	;
	if v2721 == int32(0) {
		v2790 = v2547
		goto L703
	} else {
		goto L886
	}
L886:
	;
	v2770 = F_plpgsql_build_datatype_arrayof(m, v2547)
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		goto L49
	} else {
		goto L887
	}
L887:
	;
	v2790 = v2770
	goto L703
L888:
	;
	v2790 = v2547
	goto L703
L889:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L890:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v2813
	*(*int32)(unsafe.Add(mBase, uint32(v28)+264)) = v2813
	v2820 = F_list_make1_impl(m, int32(1), v28+int32(120))
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L49
	} else {
		goto L891
	}
L891:
	;
	v2823 = F_get_collation_oid(m, v2820, int32(0))
	mBase = m.M
	v2824 = m.ExcPending
	if v2824 != 0 {
		goto L49
	} else {
		goto L892
	}
L892:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2823
	v9817 = v233
	goto L13
L893:
	;
	v2829 = F_makeString(m, v2827)
	mBase = m.M
	v2830 = m.ExcPending
	if v2830 != 0 {
		goto L49
	} else {
		goto L894
	}
L894:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v2829
	*(*int32)(unsafe.Add(mBase, uint32(v28)+260)) = v2829
	v2836 = F_list_make1_impl(m, int32(1), v28+int32(124))
	mBase = m.M
	v2837 = m.ExcPending
	if v2837 != 0 {
		goto L49
	} else {
		goto L895
	}
L895:
	;
	v2839 = F_get_collation_oid(m, v2836, int32(0))
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L49
	} else {
		goto L896
	}
L896:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2839
	v9817 = v233
	goto L13
L897:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2844
	v9817 = v233
	goto L13
L898:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2866
	v9817 = v233
	goto L13
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2873
	v9817 = v233
	goto L13
L900:
	;
	goto L901
L901:
	;
	v2878 = F_lappend(m, v2873, v2874)
	mBase = m.M
	v2879 = m.ExcPending
	if v2879 != 0 {
		goto L49
	} else {
		goto L902
	}
L902:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2878
	v9817 = v233
	goto L13
L903:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2932))) = int32(23)
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v2937 = int32(0)
	if v2936 < v2937 {
		v2982 = v2937
		goto L905
	} else {
		goto L906
	}
L904:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2932)+4)) = v2982
	v2986 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v2987 = *(*int32)(unsafe.Add(mBase, uint32(v2986)+520))
	v2989 = v2987 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2986)+520)) = v2989
	*(*int32)(unsafe.Add(mBase, uint32(v2932)+8)) = v2989
	v2994 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v2996 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(349), v2994, v2996, l1)
	mBase = m.M
	v2998 = m.ExcPending
	if v2998 != 0 {
		goto L49
	} else {
		goto L917
	}
L905:
	;
	goto L904
L906:
	;
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2943 = *(*int32)(unsafe.Add(mBase, uint32(v2942)+56))
	if v2943 == int32(0) {
		v2982 = v2937
		goto L905
	} else {
		goto L907
	}
L907:
	;
	v2946 = v2936 + v2943
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(v2942)+184))
	if base.Ui32(v2947) <= base.Ui32(v2946) {
		goto L909
	} else {
		goto L910
	}
L908:
	;
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(v2942)+192))
	if base.B2i32(v2956 == int32(0))|base.B2i32(base.Ui32(v2946) <= base.Ui32(v2956)) != 0 {
		v2982 = v2957
		goto L905
	} else {
		goto L912
	}
L909:
	;
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v2942)+188))
	v2956 = v2949
	goto L908
L910:
	;
	goto L911
L911:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2942)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2942)+184)) = v2943
	v2954 = F_strchr(m, v2943, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2942)+188)) = v2954
	v2956 = v2954
	goto L908
L912:
	;
	v2962 = v2956
	v2965 = v2957
	goto L913
L913:
	;
	v2967 = int32(1)
	v2968 = v2965 + v2967
	*(*int32)(unsafe.Add(mBase, uint32(v2942)+192)) = v2968
	v2971 = v2962 + v2967
	*(*int32)(unsafe.Add(mBase, uint32(v2942)+184)) = v2971
	v2974 = F_strchr(m, v2971, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2942)+188)) = v2974
	if v2974 == int32(0) {
		v2982 = v2968
		goto L905
	} else {
		goto L915
	}
L914:
	;
	v2982 = v2968
	goto L905
L915:
	;
	if base.Ui32(v2974) < base.Ui32(v2946) {
		v2962 = v2974
		v2965 = v2968
		goto L913
	} else {
		goto L916
	}
L916:
	;
	goto L914
L917:
	;
	v3000 = int32(0)
	v3009 = F_read_sql_construct(m, int32(59), v3000, v3000, int32(_a_F_plpgsql_yyparse_16), v3000, v3000, v3000, v28+int32(256), v3000, v2994, v2996, l1)
	mBase = m.M
	v3010 = m.ExcPending
	if v3010 != 0 {
		goto L49
	} else {
		goto L918
	}
L918:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2932)+12)) = v3009
	v3012 = *(*int32)(unsafe.Add(mBase, uint32(v3009)))
	v3014 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v3012)+3)) = v3014
	v3017 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v3012))) = v3017
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v2932)+12))
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v3019)))
	v3021 = F_strlen(m, v3020)
	mBase = m.M
	if v3021 != 0 {
		goto L919
	} else {
		goto L920
	}
L919:
	;
	base.MemoryCopy(m, v3020, v3020+int32(1), v3021)
	goto L921
L920:
	;
	goto L921
L921:
	;
	v3025 = *(*int32)(unsafe.Add(mBase, uint32(v2932)+12))
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(v3025)))
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(v3025)+4))
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	F_check_sql_expr(m, v3026, v3027, v3028+int32(1), l1)
	mBase = m.M
	v3032 = m.ExcPending
	if v3032 != 0 {
		goto L49
	} else {
		goto L922
	}
L922:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v2932
	v9817 = v233
	goto L13
L923:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3035))) = int32(24)
	v3039 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v3040 = int32(0)
	if v3039 < v3040 {
		v3085 = v3040
		goto L925
	} else {
		goto L926
	}
L924:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3035)+4)) = v3085
	v3089 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v3089)+520))
	v3092 = v3090 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3089)+520)) = v3092
	*(*int32)(unsafe.Add(mBase, uint32(v3035)+8)) = v3092
	v3097 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v3099 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(289), v3097, v3099, l1)
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L49
	} else {
		goto L937
	}
L925:
	;
	goto L924
L926:
	;
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v3045)+56))
	if v3046 == int32(0) {
		v3085 = v3040
		goto L925
	} else {
		goto L927
	}
L927:
	;
	v3049 = v3039 + v3046
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(v3045)+184))
	if base.Ui32(v3050) <= base.Ui32(v3049) {
		goto L929
	} else {
		goto L930
	}
L928:
	;
	v3060 = *(*int32)(unsafe.Add(mBase, uint32(v3045)+192))
	if base.B2i32(v3059 == int32(0))|base.B2i32(base.Ui32(v3049) <= base.Ui32(v3059)) != 0 {
		v3085 = v3060
		goto L925
	} else {
		goto L932
	}
L929:
	;
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(v3045)+188))
	v3059 = v3052
	goto L928
L930:
	;
	goto L931
L931:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3045)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3045)+184)) = v3046
	v3057 = F_strchr(m, v3046, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3045)+188)) = v3057
	v3059 = v3057
	goto L928
L932:
	;
	v3065 = v3059
	v3068 = v3060
	goto L933
L933:
	;
	v3070 = int32(1)
	v3071 = v3068 + v3070
	*(*int32)(unsafe.Add(mBase, uint32(v3045)+192)) = v3071
	v3074 = v3065 + v3070
	*(*int32)(unsafe.Add(mBase, uint32(v3045)+184)) = v3074
	v3077 = F_strchr(m, v3074, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3045)+188)) = v3077
	if v3077 == int32(0) {
		v3085 = v3071
		goto L925
	} else {
		goto L935
	}
L934:
	;
	v3085 = v3071
	goto L925
L935:
	;
	if base.Ui32(v3077) < base.Ui32(v3049) {
		v3065 = v3077
		v3068 = v3071
		goto L933
	} else {
		goto L936
	}
L936:
	;
	goto L934
L937:
	;
	v3102 = F_read_sql_stmt(m, v3097, v3099, l1)
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L49
	} else {
		goto L938
	}
L938:
	;
	v3104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3035)+16)) = uint8(v3104)
	*(*int32)(unsafe.Add(mBase, uint32(v3035)+12)) = v3102
	v3108 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3108)+524)) = uint8(v3104)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3035
	v9817 = v233
	goto L13
L939:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3113))) = int32(24)
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v3118 = int32(0)
	if v3117 < v3118 {
		v3163 = v3118
		goto L941
	} else {
		goto L942
	}
L940:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3113)+4)) = v3163
	v3167 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(v3167)+520))
	v3170 = v3168 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3167)+520)) = v3170
	*(*int32)(unsafe.Add(mBase, uint32(v3113)+8)) = v3170
	v3175 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v3177 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(309), v3175, v3177, l1)
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L49
	} else {
		goto L953
	}
L941:
	;
	goto L940
L942:
	;
	v3123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(v3123)+56))
	if v3124 == int32(0) {
		v3163 = v3118
		goto L941
	} else {
		goto L943
	}
L943:
	;
	v3127 = v3117 + v3124
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v3123)+184))
	if base.Ui32(v3128) <= base.Ui32(v3127) {
		goto L945
	} else {
		goto L946
	}
L944:
	;
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v3123)+192))
	if base.B2i32(v3137 == int32(0))|base.B2i32(base.Ui32(v3127) <= base.Ui32(v3137)) != 0 {
		v3163 = v3138
		goto L941
	} else {
		goto L948
	}
L945:
	;
	v3130 = *(*int32)(unsafe.Add(mBase, uint32(v3123)+188))
	v3137 = v3130
	goto L944
L946:
	;
	goto L947
L947:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3123)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3123)+184)) = v3124
	v3135 = F_strchr(m, v3124, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3123)+188)) = v3135
	v3137 = v3135
	goto L944
L948:
	;
	v3143 = v3137
	v3146 = v3138
	goto L949
L949:
	;
	v3148 = int32(1)
	v3149 = v3146 + v3148
	*(*int32)(unsafe.Add(mBase, uint32(v3123)+192)) = v3149
	v3152 = v3143 + v3148
	*(*int32)(unsafe.Add(mBase, uint32(v3123)+184)) = v3152
	v3155 = F_strchr(m, v3152, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3123)+188)) = v3155
	if v3155 == int32(0) {
		v3163 = v3149
		goto L941
	} else {
		goto L951
	}
L950:
	;
	v3163 = v3149
	goto L941
L951:
	;
	if base.Ui32(v3155) < base.Ui32(v3127) {
		v3143 = v3155
		v3146 = v3149
		goto L949
	} else {
		goto L952
	}
L952:
	;
	goto L950
L953:
	;
	v3180 = F_read_sql_stmt(m, v3175, v3177, l1)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L49
	} else {
		goto L954
	}
L954:
	;
	v3182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3113)+16)) = uint8(v3182)
	*(*int32)(unsafe.Add(mBase, uint32(v3113)+12)) = v3180
	v3186 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v3187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3186)+524)) = uint8(v3187)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3113
	v9817 = v233
	goto L13
L955:
	;
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	if v3194 == int32(0) {
		goto L34
	} else {
		goto L958
	}
L956:
	;
	v3205 = int32(3)
	goto L957
L957:
	;
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_check_assignable(m, v3206, v3207, l1)
	mBase = m.M
	v3209 = m.ExcPending
	if v3209 != 0 {
		goto L49
	} else {
		goto L960
	}
L958:
	;
	v3197 = *(*int32)(unsafe.Add(mBase, uint32(v3194)+4))
	if base.Ui32(int32(3)) <= base.Ui32(v3197-int32(1)) {
		goto L34
	} else {
		goto L959
	}
L959:
	;
	v3205 = v3197 + int32(2)
	goto L957
L960:
	;
	v3211 = F_palloc0(m, int32(20))
	mBase = m.M
	v3212 = m.ExcPending
	if v3212 != 0 {
		goto L49
	} else {
		goto L961
	}
L961:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3211))) = int32(1)
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v3216 = int32(0)
	if v3215 < v3216 {
		v3261 = v3216
		goto L963
	} else {
		goto L964
	}
L962:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+4)) = v3261
	v3265 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v3265)+520))
	v3268 = v3266 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3265)+520)) = v3268
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+8)) = v3268
	v3271 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3272 = *(*int32)(unsafe.Add(mBase, uint32(v3271)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+12)) = v3272
	v3276 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v3278 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, int32(277), v3276, v3278, l1)
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L49
	} else {
		goto L975
	}
L963:
	;
	goto L962
L964:
	;
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3221)+56))
	if v3222 == int32(0) {
		v3261 = v3216
		goto L963
	} else {
		goto L965
	}
L965:
	;
	v3225 = v3215 + v3222
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(v3221)+184))
	if base.Ui32(v3226) <= base.Ui32(v3225) {
		goto L967
	} else {
		goto L968
	}
L966:
	;
	v3236 = *(*int32)(unsafe.Add(mBase, uint32(v3221)+192))
	if base.B2i32(v3235 == int32(0))|base.B2i32(base.Ui32(v3225) <= base.Ui32(v3235)) != 0 {
		v3261 = v3236
		goto L963
	} else {
		goto L970
	}
L967:
	;
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v3221)+188))
	v3235 = v3228
	goto L966
L968:
	;
	goto L969
L969:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3221)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3221)+184)) = v3222
	v3233 = F_strchr(m, v3222, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3221)+188)) = v3233
	v3235 = v3233
	goto L966
L970:
	;
	v3241 = v3235
	v3244 = v3236
	goto L971
L971:
	;
	v3246 = int32(1)
	v3247 = v3244 + v3246
	*(*int32)(unsafe.Add(mBase, uint32(v3221)+192)) = v3247
	v3250 = v3241 + v3246
	*(*int32)(unsafe.Add(mBase, uint32(v3221)+184)) = v3250
	v3253 = F_strchr(m, v3250, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3221)+188)) = v3253
	if v3253 == int32(0) {
		v3261 = v3247
		goto L963
	} else {
		goto L973
	}
L972:
	;
	v3261 = v3247
	goto L963
L973:
	;
	if base.Ui32(v3253) < base.Ui32(v3225) {
		v3241 = v3253
		v3244 = v3247
		goto L971
	} else {
		goto L974
	}
L974:
	;
	goto L972
L975:
	;
	v3282 = int32(0)
	v3289 = F_read_sql_construct(m, int32(59), v3282, v3282, int32(_a_F_plpgsql_yyparse_16), v3205, v3282, int32(1), v3282, v3282, v3276, v3278, l1)
	mBase = m.M
	v3290 = m.ExcPending
	if v3290 != 0 {
		goto L49
	} else {
		goto L976
	}
L976:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+16)) = v3289
	v3293 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v3293)))
	if v3294 != 0 {
		goto L977
	} else {
		goto L978
	}
L977:
	;
	v3298 = int32(-1)
	v3299 = int32(0)
	goto L979
L978:
	;
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(v3293)+4))
	v3298 = v3296
	v3299 = int32(1)
	goto L979
L979:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3289)+20)) = uint8(v3299)
	*(*int32)(unsafe.Add(mBase, uint32(v3289)+16)) = v3298
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3211
	v9817 = v233
	goto L13
L980:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3304))) = int32(19)
	v3309 = v147 - int32(16)
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v3309)))
	v3311 = int32(0)
	if v3310 < v3311 {
		v3356 = v3311
		goto L982
	} else {
		goto L983
	}
L981:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3304)+4)) = v3356
	v3360 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v3361 = *(*int32)(unsafe.Add(mBase, uint32(v3360)+520))
	v3363 = v3361 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3360)+520)) = v3363
	*(*int32)(unsafe.Add(mBase, uint32(v3304)+8)) = v3363
	v3368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148-int32(48)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3304)+12)) = uint8(v3368)
	v3372 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v3304)+16)) = v3372
	if v3372 == int32(0) {
		goto L994
	} else {
		goto L995
	}
L982:
	;
	goto L981
L983:
	;
	v3316 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3317 = *(*int32)(unsafe.Add(mBase, uint32(v3316)+56))
	if v3317 == int32(0) {
		v3356 = v3311
		goto L982
	} else {
		goto L984
	}
L984:
	;
	v3320 = v3310 + v3317
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v3316)+184))
	if base.Ui32(v3321) <= base.Ui32(v3320) {
		goto L986
	} else {
		goto L987
	}
L985:
	;
	v3331 = *(*int32)(unsafe.Add(mBase, uint32(v3316)+192))
	if base.B2i32(v3330 == int32(0))|base.B2i32(base.Ui32(v3320) <= base.Ui32(v3330)) != 0 {
		v3356 = v3331
		goto L982
	} else {
		goto L989
	}
L986:
	;
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v3316)+188))
	v3330 = v3323
	goto L985
L987:
	;
	goto L988
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3316)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3316)+184)) = v3317
	v3328 = F_strchr(m, v3317, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3316)+188)) = v3328
	v3330 = v3328
	goto L985
L989:
	;
	v3336 = v3330
	v3339 = v3331
	goto L990
L990:
	;
	v3341 = int32(1)
	v3342 = v3339 + v3341
	*(*int32)(unsafe.Add(mBase, uint32(v3316)+192)) = v3342
	v3345 = v3336 + v3341
	*(*int32)(unsafe.Add(mBase, uint32(v3316)+184)) = v3345
	v3348 = F_strchr(m, v3345, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3316)+188)) = v3348
	if v3348 == int32(0) {
		v3356 = v3342
		goto L982
	} else {
		goto L992
	}
L991:
	;
	v3356 = v3342
	goto L982
L992:
	;
	if base.Ui32(v3348) < base.Ui32(v3320) {
		v3336 = v3348
		v3339 = v3342
		goto L990
	} else {
		goto L993
	}
L993:
	;
	goto L991
L994:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3304
	v9817 = v233
	goto L13
L995:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(v3372)+4))
	if v3376 <= int32(0) {
		goto L994
	} else {
		goto L996
	}
L996:
	;
	v3379 = int32(0)
	if v3379 < v3376 {
		goto L997
	} else {
		goto L998
	}
L997:
	;
	v3383 = v3376
	goto L999
L998:
	;
	v3383 = v3379
	goto L999
L999:
	;
	v3384 = *(*int32)(unsafe.Add(mBase, uint32(v3372)+12))
	v3389 = v3379
	goto L1000
L1000:
	;
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(v3384+v3389<<(uint(int32(2))%32))))
	v3414 = *(*int32)(unsafe.Add(mBase, uint32(v3413)))
	if base.Ui32(int32(10)) <= base.Ui32(v3414-int32(3)) {
		goto L1004
	} else {
		goto L1005
	}
L1001:
	;
	goto L994
L1002:
	;
	v3504 = v3389 + int32(1)
	if v3504 != v3383 {
		v3389 = v3504
		goto L1000
	} else {
		goto L1031
	}
L1003:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v3490 = m.ExcPending
	if v3490 != 0 {
		goto L49
	} else {
		goto L1028
	}
L1004:
	;
	switch v3414 {
	case 0, 1:
		goto L1007
	case 2:
		goto L1002
	default:
		goto L1003
	}
L1005:
	;
	goto L1006
L1006:
	;
	if v3368&int32(1) != 0 {
		goto L1002
	} else {
		goto L1018
	}
L1007:
	;
	if v3368&int32(1) == int32(0) {
		goto L1002
	} else {
		goto L1008
	}
L1008:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L49
	} else {
		goto L1009
	}
L1009:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3429 = m.ExcPending
	if v3429 != 0 {
		goto L49
	} else {
		goto L1010
	}
L1010:
	;
	v3430 = *(*int32)(unsafe.Add(mBase, uint32(v3413)))
	if base.Ui32(int32(12)) < base.Ui32(v3430) {
		goto L1012
	} else {
		goto L1013
	}
L1011:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+144)) = v3439
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_36), v28+int32(144))
	mBase = m.M
	v3445 = m.ExcPending
	if v3445 != 0 {
		goto L49
	} else {
		goto L1015
	}
L1012:
	;
	v3439 = int32(_a_F_plpgsql_yyparse_37)
	goto L1011
L1013:
	;
	goto L1014
L1014:
	;
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(v3430<<(uint(int32(2))%32))+uint32(_c_F_plpgsql_yyparse[27])))
	v3439 = v3438
	goto L1011
L1015:
	;
	v3446 = *(*int32)(unsafe.Add(mBase, uint32(v3309)))
	v3447 = F_plpgsql_scanner_errposition(m, v3446, l1)
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L49
	} else {
		goto L1016
	}
L1016:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1042), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v3453 = m.ExcPending
	if v3453 != 0 {
		goto L49
	} else {
		goto L1017
	}
L1017:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1018:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v3459 = m.ExcPending
	if v3459 != 0 {
		goto L49
	} else {
		goto L1019
	}
L1019:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		goto L49
	} else {
		goto L1020
	}
L1020:
	;
	v3463 = *(*int32)(unsafe.Add(mBase, uint32(v3413)))
	if base.Ui32(int32(12)) < base.Ui32(v3463) {
		goto L1022
	} else {
		goto L1023
	}
L1021:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v3472
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_38), v28+int32(160))
	mBase = m.M
	v3478 = m.ExcPending
	if v3478 != 0 {
		goto L49
	} else {
		goto L1025
	}
L1022:
	;
	v3472 = int32(_a_F_plpgsql_yyparse_37)
	goto L1021
L1023:
	;
	goto L1024
L1024:
	;
	v3471 = *(*int32)(unsafe.Add(mBase, uint32(v3463<<(uint(int32(2))%32))+uint32(_c_F_plpgsql_yyparse[27])))
	v3472 = v3471
	goto L1021
L1025:
	;
	v3479 = *(*int32)(unsafe.Add(mBase, uint32(v3309)))
	v3480 = F_plpgsql_scanner_errposition(m, v3479, l1)
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L49
	} else {
		goto L1026
	}
L1026:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1060), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		goto L49
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
	v3491 = *(*int32)(unsafe.Add(mBase, uint32(v3413)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v3491
	F_errmsg_internal(m, int32(_a_F_plpgsql_yyparse_39), v28+int32(128))
	mBase = m.M
	v3497 = m.ExcPending
	if v3497 != 0 {
		goto L49
	} else {
		goto L1029
	}
L1029:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1067), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v3502 = m.ExcPending
	if v3502 != 0 {
		goto L49
	} else {
		goto L1030
	}
L1030:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1031:
	;
	goto L1001
L1032:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3542
	v9817 = v233
	goto L13
L1033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3551
	v9817 = v233
	goto L13
L1034:
	;
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(v3559)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3555)+4)) = v3560
	v3562 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v3555))) = v3562
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v3555
	v9817 = v233
	goto L13
L1035:
	;
	F_plpgsql_yyerror(m, v28+int32(_a_F_plpgsql_yyparse_2), int32(0), l1, int32(_a_F_plpgsql_yyparse_40))
	mBase = m.M
	v3993 = m.ExcPending
	if v3993 != 0 {
		goto L49
	} else {
		goto L1175
	}
L1036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(6)
	v9817 = v233
	goto L13
L1037:
	;
	v3954 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v3954 == int32(0) {
		goto L1035
	} else {
		goto L1166
	}
L1038:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(12)
	v9817 = v233
	goto L13
L1039:
	;
	v3924 = int32(_a_F_plpgsql_yyparse_41)
	v3927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3891))))
	v3930 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[28])))
	if base.B2i32(v3927 == int32(0))|base.B2i32(v3927 != v3930) != 0 {
		v3948 = v3927
		v3949 = v3930
		goto L1159
	} else {
		goto L1160
	}
L1040:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(11)
	v9817 = v233
	goto L13
L1041:
	;
	v3891 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v3891 == int32(0) {
		goto L1035
	} else {
		goto L1149
	}
L1042:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(10)
	v9817 = v233
	goto L13
L1043:
	;
	v3861 = int32(_a_F_plpgsql_yyparse_42)
	v3864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3828))))
	v3867 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[29])))
	if base.B2i32(v3864 == int32(0))|base.B2i32(v3864 != v3867) != 0 {
		v3885 = v3864
		v3886 = v3867
		goto L1142
	} else {
		goto L1143
	}
L1044:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(9)
	v9817 = v233
	goto L13
L1045:
	;
	v3828 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v3828 == int32(0) {
		goto L1035
	} else {
		goto L1132
	}
L1046:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(8)
	v9817 = v233
	goto L13
L1047:
	;
	v3798 = int32(_a_F_plpgsql_yyparse_43)
	v3801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3765))))
	v3804 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[30])))
	if base.B2i32(v3801 == int32(0))|base.B2i32(v3801 != v3804) != 0 {
		v3822 = v3801
		v3823 = v3804
		goto L1125
	} else {
		goto L1126
	}
L1048:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(7)
	v9817 = v233
	goto L13
L1049:
	;
	v3765 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v3765 == int32(0) {
		goto L1035
	} else {
		goto L1115
	}
L1050:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(3)
	v9817 = v233
	goto L13
L1051:
	;
	v3735 = int32(_a_F_plpgsql_yyparse_44)
	v3738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3702))))
	v3741 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[31])))
	if base.B2i32(v3738 == int32(0))|base.B2i32(v3738 != v3741) != 0 {
		v3759 = v3738
		v3760 = v3741
		goto L1108
	} else {
		goto L1109
	}
L1052:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(5)
	v9817 = v233
	goto L13
L1053:
	;
	v3702 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v3702 == int32(0) {
		goto L1035
	} else {
		goto L1098
	}
L1054:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(4)
	v9817 = v233
	goto L13
L1055:
	;
	v3672 = int32(_a_F_plpgsql_yyparse_45)
	v3675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3639))))
	v3678 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[32])))
	if base.B2i32(v3675 == int32(0))|base.B2i32(v3675 != v3678) != 0 {
		v3696 = v3675
		v3697 = v3678
		goto L1091
	} else {
		goto L1092
	}
L1056:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(2)
	v9817 = v233
	goto L13
L1057:
	;
	v3639 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v3639 == int32(0) {
		goto L1035
	} else {
		goto L1081
	}
L1058:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(1)
	v9817 = v233
	goto L13
L1059:
	;
	v3609 = int32(_a_F_plpgsql_yyparse_46)
	v3612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3576))))
	v3615 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[33])))
	if base.B2i32(v3612 == int32(0))|base.B2i32(v3612 != v3615) != 0 {
		v3633 = v3612
		v3634 = v3615
		goto L1074
	} else {
		goto L1075
	}
L1060:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L1061:
	;
	v3573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v3573&int32(1) != 0 {
		goto L1035
	} else {
		goto L1063
	}
L1062:
	;
	switch v3569 - int32(277) {
	case 0:
		goto L1061
	default:
		goto L1035
	case 18:
		goto L1048
	case 22:
		goto L1046
	case 62:
		goto L1042
	case 73:
		goto L1056
	case 74:
		goto L1044
	case 75:
		goto L1050
	case 76:
		goto L1054
	case 77:
		goto L1052
	case 78:
		goto L1058
	case 85:
		goto L1036
	case 88:
		goto L1060
	case 91:
		goto L1038
	case 98:
		goto L1040
	}
L1063:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v3576 == int32(0) {
		goto L1035
	} else {
		goto L1064
	}
L1064:
	;
	v3579 = int32(_a_F_plpgsql_yyparse_47)
	v3582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3576))))
	v3585 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[34])))
	if base.B2i32(v3582 == int32(0))|base.B2i32(v3582 != v3585) != 0 {
		v3603 = v3582
		v3604 = v3585
		goto L1066
	} else {
		goto L1067
	}
L1065:
	;
	if v3603-v3604 != 0 {
		goto L1059
	} else {
		goto L1072
	}
L1066:
	;
	goto L1065
L1067:
	;
	v3588 = v3576
	v3589 = v3579
	goto L1068
L1068:
	;
	v3592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3589)+1)))
	v3593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3588)+1)))
	if v3593 == int32(0) {
		v3603 = v3593
		v3604 = v3592
		goto L1066
	} else {
		goto L1070
	}
L1069:
	;
	v3603 = v3593
	v3604 = v3592
	goto L1066
L1070:
	;
	v3596 = int32(1)
	if v3593 == v3592 {
		v3588 = v3588 + v3596
		v3589 = v3589 + v3596
		goto L1068
	} else {
		goto L1071
	}
L1071:
	;
	goto L1069
L1072:
	;
	goto L1060
L1073:
	;
	if v3633-v3634 != 0 {
		goto L1057
	} else {
		goto L1080
	}
L1074:
	;
	goto L1073
L1075:
	;
	v3618 = v3576
	v3619 = v3609
	goto L1076
L1076:
	;
	v3622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3619)+1)))
	v3623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3618)+1)))
	if v3623 == int32(0) {
		v3633 = v3623
		v3634 = v3622
		goto L1074
	} else {
		goto L1078
	}
L1077:
	;
	v3633 = v3623
	v3634 = v3622
	goto L1074
L1078:
	;
	v3626 = int32(1)
	if v3623 == v3622 {
		v3618 = v3618 + v3626
		v3619 = v3619 + v3626
		goto L1076
	} else {
		goto L1079
	}
L1079:
	;
	goto L1077
L1080:
	;
	goto L1058
L1081:
	;
	v3642 = int32(_a_F_plpgsql_yyparse_48)
	v3645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3639))))
	v3648 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[35])))
	if base.B2i32(v3645 == int32(0))|base.B2i32(v3645 != v3648) != 0 {
		v3666 = v3645
		v3667 = v3648
		goto L1083
	} else {
		goto L1084
	}
L1082:
	;
	if v3666-v3667 != 0 {
		goto L1055
	} else {
		goto L1089
	}
L1083:
	;
	goto L1082
L1084:
	;
	v3651 = v3639
	v3652 = v3642
	goto L1085
L1085:
	;
	v3655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3652)+1)))
	v3656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3651)+1)))
	if v3656 == int32(0) {
		v3666 = v3656
		v3667 = v3655
		goto L1083
	} else {
		goto L1087
	}
L1086:
	;
	v3666 = v3656
	v3667 = v3655
	goto L1083
L1087:
	;
	v3659 = int32(1)
	if v3656 == v3655 {
		v3651 = v3651 + v3659
		v3652 = v3652 + v3659
		goto L1085
	} else {
		goto L1088
	}
L1088:
	;
	goto L1086
L1089:
	;
	goto L1056
L1090:
	;
	if v3696-v3697 != 0 {
		goto L1053
	} else {
		goto L1097
	}
L1091:
	;
	goto L1090
L1092:
	;
	v3681 = v3639
	v3682 = v3672
	goto L1093
L1093:
	;
	v3685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3682)+1)))
	v3686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3681)+1)))
	if v3686 == int32(0) {
		v3696 = v3686
		v3697 = v3685
		goto L1091
	} else {
		goto L1095
	}
L1094:
	;
	v3696 = v3686
	v3697 = v3685
	goto L1091
L1095:
	;
	v3689 = int32(1)
	if v3686 == v3685 {
		v3681 = v3681 + v3689
		v3682 = v3682 + v3689
		goto L1093
	} else {
		goto L1096
	}
L1096:
	;
	goto L1094
L1097:
	;
	goto L1054
L1098:
	;
	v3705 = int32(_a_F_plpgsql_yyparse_49)
	v3708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3702))))
	v3711 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[36])))
	if base.B2i32(v3708 == int32(0))|base.B2i32(v3708 != v3711) != 0 {
		v3729 = v3708
		v3730 = v3711
		goto L1100
	} else {
		goto L1101
	}
L1099:
	;
	if v3729-v3730 != 0 {
		goto L1051
	} else {
		goto L1106
	}
L1100:
	;
	goto L1099
L1101:
	;
	v3714 = v3702
	v3715 = v3705
	goto L1102
L1102:
	;
	v3718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3715)+1)))
	v3719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3714)+1)))
	if v3719 == int32(0) {
		v3729 = v3719
		v3730 = v3718
		goto L1100
	} else {
		goto L1104
	}
L1103:
	;
	v3729 = v3719
	v3730 = v3718
	goto L1100
L1104:
	;
	v3722 = int32(1)
	if v3719 == v3718 {
		v3714 = v3714 + v3722
		v3715 = v3715 + v3722
		goto L1102
	} else {
		goto L1105
	}
L1105:
	;
	goto L1103
L1106:
	;
	goto L1052
L1107:
	;
	if v3759-v3760 != 0 {
		goto L1049
	} else {
		goto L1114
	}
L1108:
	;
	goto L1107
L1109:
	;
	v3744 = v3702
	v3745 = v3735
	goto L1110
L1110:
	;
	v3748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3745)+1)))
	v3749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3744)+1)))
	if v3749 == int32(0) {
		v3759 = v3749
		v3760 = v3748
		goto L1108
	} else {
		goto L1112
	}
L1111:
	;
	v3759 = v3749
	v3760 = v3748
	goto L1108
L1112:
	;
	v3752 = int32(1)
	if v3749 == v3748 {
		v3744 = v3744 + v3752
		v3745 = v3745 + v3752
		goto L1110
	} else {
		goto L1113
	}
L1113:
	;
	goto L1111
L1114:
	;
	goto L1050
L1115:
	;
	v3768 = int32(_a_F_plpgsql_yyparse_50)
	v3771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3765))))
	v3774 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[37])))
	if base.B2i32(v3771 == int32(0))|base.B2i32(v3771 != v3774) != 0 {
		v3792 = v3771
		v3793 = v3774
		goto L1117
	} else {
		goto L1118
	}
L1116:
	;
	if v3792-v3793 != 0 {
		goto L1047
	} else {
		goto L1123
	}
L1117:
	;
	goto L1116
L1118:
	;
	v3777 = v3765
	v3778 = v3768
	goto L1119
L1119:
	;
	v3781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3778)+1)))
	v3782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3777)+1)))
	if v3782 == int32(0) {
		v3792 = v3782
		v3793 = v3781
		goto L1117
	} else {
		goto L1121
	}
L1120:
	;
	v3792 = v3782
	v3793 = v3781
	goto L1117
L1121:
	;
	v3785 = int32(1)
	if v3782 == v3781 {
		v3777 = v3777 + v3785
		v3778 = v3778 + v3785
		goto L1119
	} else {
		goto L1122
	}
L1122:
	;
	goto L1120
L1123:
	;
	goto L1048
L1124:
	;
	if v3822-v3823 != 0 {
		goto L1045
	} else {
		goto L1131
	}
L1125:
	;
	goto L1124
L1126:
	;
	v3807 = v3765
	v3808 = v3798
	goto L1127
L1127:
	;
	v3811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3808)+1)))
	v3812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3807)+1)))
	if v3812 == int32(0) {
		v3822 = v3812
		v3823 = v3811
		goto L1125
	} else {
		goto L1129
	}
L1128:
	;
	v3822 = v3812
	v3823 = v3811
	goto L1125
L1129:
	;
	v3815 = int32(1)
	if v3812 == v3811 {
		v3807 = v3807 + v3815
		v3808 = v3808 + v3815
		goto L1127
	} else {
		goto L1130
	}
L1130:
	;
	goto L1128
L1131:
	;
	goto L1046
L1132:
	;
	v3831 = int32(_a_F_plpgsql_yyparse_51)
	v3834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3828))))
	v3837 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[38])))
	if base.B2i32(v3834 == int32(0))|base.B2i32(v3834 != v3837) != 0 {
		v3855 = v3834
		v3856 = v3837
		goto L1134
	} else {
		goto L1135
	}
L1133:
	;
	if v3855-v3856 != 0 {
		goto L1043
	} else {
		goto L1140
	}
L1134:
	;
	goto L1133
L1135:
	;
	v3840 = v3828
	v3841 = v3831
	goto L1136
L1136:
	;
	v3844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3841)+1)))
	v3845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3840)+1)))
	if v3845 == int32(0) {
		v3855 = v3845
		v3856 = v3844
		goto L1134
	} else {
		goto L1138
	}
L1137:
	;
	v3855 = v3845
	v3856 = v3844
	goto L1134
L1138:
	;
	v3848 = int32(1)
	if v3845 == v3844 {
		v3840 = v3840 + v3848
		v3841 = v3841 + v3848
		goto L1136
	} else {
		goto L1139
	}
L1139:
	;
	goto L1137
L1140:
	;
	goto L1044
L1141:
	;
	if v3885-v3886 != 0 {
		goto L1041
	} else {
		goto L1148
	}
L1142:
	;
	goto L1141
L1143:
	;
	v3870 = v3828
	v3871 = v3861
	goto L1144
L1144:
	;
	v3874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3871)+1)))
	v3875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3870)+1)))
	if v3875 == int32(0) {
		v3885 = v3875
		v3886 = v3874
		goto L1142
	} else {
		goto L1146
	}
L1145:
	;
	v3885 = v3875
	v3886 = v3874
	goto L1142
L1146:
	;
	v3878 = int32(1)
	if v3875 == v3874 {
		v3870 = v3870 + v3878
		v3871 = v3871 + v3878
		goto L1144
	} else {
		goto L1147
	}
L1147:
	;
	goto L1145
L1148:
	;
	goto L1042
L1149:
	;
	v3894 = int32(_a_F_plpgsql_yyparse_52)
	v3897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3891))))
	v3900 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[39])))
	if base.B2i32(v3897 == int32(0))|base.B2i32(v3897 != v3900) != 0 {
		v3918 = v3897
		v3919 = v3900
		goto L1151
	} else {
		goto L1152
	}
L1150:
	;
	if v3918-v3919 != 0 {
		goto L1039
	} else {
		goto L1157
	}
L1151:
	;
	goto L1150
L1152:
	;
	v3903 = v3891
	v3904 = v3894
	goto L1153
L1153:
	;
	v3907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3904)+1)))
	v3908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3903)+1)))
	if v3908 == int32(0) {
		v3918 = v3908
		v3919 = v3907
		goto L1151
	} else {
		goto L1155
	}
L1154:
	;
	v3918 = v3908
	v3919 = v3907
	goto L1151
L1155:
	;
	v3911 = int32(1)
	if v3908 == v3907 {
		v3903 = v3903 + v3911
		v3904 = v3904 + v3911
		goto L1153
	} else {
		goto L1156
	}
L1156:
	;
	goto L1154
L1157:
	;
	goto L1040
L1158:
	;
	if v3948-v3949 != 0 {
		goto L1037
	} else {
		goto L1165
	}
L1159:
	;
	goto L1158
L1160:
	;
	v3933 = v3891
	v3934 = v3924
	goto L1161
L1161:
	;
	v3937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3934)+1)))
	v3938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3933)+1)))
	if v3938 == int32(0) {
		v3948 = v3938
		v3949 = v3937
		goto L1159
	} else {
		goto L1163
	}
L1162:
	;
	v3948 = v3938
	v3949 = v3937
	goto L1159
L1163:
	;
	v3941 = int32(1)
	if v3938 == v3937 {
		v3933 = v3933 + v3941
		v3934 = v3934 + v3941
		goto L1161
	} else {
		goto L1164
	}
L1164:
	;
	goto L1162
L1165:
	;
	goto L1038
L1166:
	;
	v3957 = int32(_a_F_plpgsql_yyparse_53)
	v3960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3954))))
	v3963 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[40])))
	if base.B2i32(v3960 == int32(0))|base.B2i32(v3960 != v3963) != 0 {
		v3981 = v3960
		v3982 = v3963
		goto L1168
	} else {
		goto L1169
	}
L1167:
	;
	if v3981-v3982 != 0 {
		goto L1035
	} else {
		goto L1174
	}
L1168:
	;
	goto L1167
L1169:
	;
	v3966 = v3954
	v3967 = v3957
	goto L1170
L1170:
	;
	v3970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3967)+1)))
	v3971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3966)+1)))
	if v3971 == int32(0) {
		v3981 = v3971
		v3982 = v3970
		goto L1168
	} else {
		goto L1172
	}
L1171:
	;
	v3981 = v3971
	v3982 = v3970
	goto L1168
L1172:
	;
	v3974 = int32(1)
	if v3971 == v3970 {
		v3966 = v3966 + v3974
		v3967 = v3967 + v3974
		goto L1170
	} else {
		goto L1173
	}
L1173:
	;
	goto L1171
L1174:
	;
	goto L1036
L1175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1176:
	;
	v4000 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v4001 = m.ExcPending
	if v4001 != 0 {
		goto L49
	} else {
		goto L1177
	}
L1177:
	;
	if v4000 == int32(91) {
		goto L33
	} else {
		goto L1178
	}
L1178:
	;
	v4004 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v4005 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_check_assignable(m, v4004, v4005, l1)
	mBase = m.M
	v4007 = m.ExcPending
	if v4007 != 0 {
		goto L49
	} else {
		goto L1179
	}
L1179:
	;
	v4008 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4008
	v9817 = v233
	goto L13
L1180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4011))) = int32(2)
	v4017 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(28))))
	v4018 = int32(0)
	if v4017 < v4018 {
		v4063 = v4018
		goto L1182
	} else {
		goto L1183
	}
L1181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4011)+4)) = v4063
	v4067 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v4068 = *(*int32)(unsafe.Add(mBase, uint32(v4067)+520))
	v4070 = v4068 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4067)+520)) = v4070
	*(*int32)(unsafe.Add(mBase, uint32(v4011)+8)) = v4070
	v4075 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(96))))
	*(*int32)(unsafe.Add(mBase, uint32(v4011)+12)) = v4075
	v4079 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(80))))
	*(*int32)(unsafe.Add(mBase, uint32(v4011)+16)) = v4079
	v4083 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v4011)+20)) = v4083
	v4087 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	*(*int32)(unsafe.Add(mBase, uint32(v4011)+24)) = v4087
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4011
	v9817 = v233
	goto L13
L1182:
	;
	goto L1181
L1183:
	;
	v4023 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4024 = *(*int32)(unsafe.Add(mBase, uint32(v4023)+56))
	if v4024 == int32(0) {
		v4063 = v4018
		goto L1182
	} else {
		goto L1184
	}
L1184:
	;
	v4027 = v4017 + v4024
	v4028 = *(*int32)(unsafe.Add(mBase, uint32(v4023)+184))
	if base.Ui32(v4028) <= base.Ui32(v4027) {
		goto L1186
	} else {
		goto L1187
	}
L1185:
	;
	v4038 = *(*int32)(unsafe.Add(mBase, uint32(v4023)+192))
	if base.B2i32(v4037 == int32(0))|base.B2i32(base.Ui32(v4027) <= base.Ui32(v4037)) != 0 {
		v4063 = v4038
		goto L1182
	} else {
		goto L1189
	}
L1186:
	;
	v4030 = *(*int32)(unsafe.Add(mBase, uint32(v4023)+188))
	v4037 = v4030
	goto L1185
L1187:
	;
	goto L1188
L1188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4023)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4023)+184)) = v4024
	v4035 = F_strchr(m, v4024, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4023)+188)) = v4035
	v4037 = v4035
	goto L1185
L1189:
	;
	v4043 = v4037
	v4046 = v4038
	goto L1190
L1190:
	;
	v4048 = int32(1)
	v4049 = v4046 + v4048
	*(*int32)(unsafe.Add(mBase, uint32(v4023)+192)) = v4049
	v4052 = v4043 + v4048
	*(*int32)(unsafe.Add(mBase, uint32(v4023)+184)) = v4052
	v4055 = F_strchr(m, v4052, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4023)+188)) = v4055
	if v4055 == int32(0) {
		v4063 = v4049
		goto L1182
	} else {
		goto L1192
	}
L1191:
	;
	v4063 = v4049
	goto L1182
L1192:
	;
	if base.Ui32(v4055) < base.Ui32(v4027) {
		v4043 = v4055
		v4046 = v4049
		goto L1190
	} else {
		goto L1193
	}
L1193:
	;
	goto L1191
L1194:
	;
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4098 = int32(0)
	if v4097 < v4098 {
		v4143 = v4098
		goto L1196
	} else {
		goto L1197
	}
L1195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4093))) = v4143
	v4148 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4093)+4)) = v4148
	v4150 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4093)+8)) = v4150
	v4154 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(48))))
	v4155 = F_lappend(m, v4154, v4093)
	mBase = m.M
	v4156 = m.ExcPending
	if v4156 != 0 {
		goto L49
	} else {
		goto L1208
	}
L1196:
	;
	goto L1195
L1197:
	;
	v4103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4104 = *(*int32)(unsafe.Add(mBase, uint32(v4103)+56))
	if v4104 == int32(0) {
		v4143 = v4098
		goto L1196
	} else {
		goto L1198
	}
L1198:
	;
	v4107 = v4097 + v4104
	v4108 = *(*int32)(unsafe.Add(mBase, uint32(v4103)+184))
	if base.Ui32(v4108) <= base.Ui32(v4107) {
		goto L1200
	} else {
		goto L1201
	}
L1199:
	;
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(v4103)+192))
	if base.B2i32(v4117 == int32(0))|base.B2i32(base.Ui32(v4107) <= base.Ui32(v4117)) != 0 {
		v4143 = v4118
		goto L1196
	} else {
		goto L1203
	}
L1200:
	;
	v4110 = *(*int32)(unsafe.Add(mBase, uint32(v4103)+188))
	v4117 = v4110
	goto L1199
L1201:
	;
	goto L1202
L1202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4103)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4103)+184)) = v4104
	v4115 = F_strchr(m, v4104, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4103)+188)) = v4115
	v4117 = v4115
	goto L1199
L1203:
	;
	v4123 = v4117
	v4126 = v4118
	goto L1204
L1204:
	;
	v4128 = int32(1)
	v4129 = v4126 + v4128
	*(*int32)(unsafe.Add(mBase, uint32(v4103)+192)) = v4129
	v4132 = v4123 + v4128
	*(*int32)(unsafe.Add(mBase, uint32(v4103)+184)) = v4132
	v4135 = F_strchr(m, v4132, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4103)+188)) = v4135
	if v4135 == int32(0) {
		v4143 = v4129
		goto L1196
	} else {
		goto L1206
	}
L1205:
	;
	v4143 = v4129
	goto L1196
L1206:
	;
	if base.Ui32(v4135) < base.Ui32(v4107) {
		v4123 = v4135
		v4126 = v4129
		goto L1204
	} else {
		goto L1207
	}
L1207:
	;
	goto L1205
L1208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4155
	v9817 = v233
	goto L13
L1209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4179))) = int32(3)
	v4183 = int32(0)
	if v4164 < v4183 {
		v4228 = v4183
		goto L1211
	} else {
		goto L1212
	}
L1210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4179)+4)) = v4228
	v4232 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v4233 = *(*int32)(unsafe.Add(mBase, uint32(v4232)+520))
	v4235 = v4233 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4232)+520)) = v4235
	v4237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4179)+24)) = uint8(base.B2i32(v4173 != v4237))
	*(*int32)(unsafe.Add(mBase, uint32(v4179)+20)) = v4170
	*(*int32)(unsafe.Add(mBase, uint32(v4179)+16)) = v4237
	*(*int32)(unsafe.Add(mBase, uint32(v4179)+12)) = v4167
	*(*int32)(unsafe.Add(mBase, uint32(v4179)+8)) = v4235
	if v4173 == v4237 {
		v4253 = v4173
		goto L1223
	} else {
		goto L1224
	}
L1211:
	;
	goto L1210
L1212:
	;
	v4188 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+56))
	if v4189 == int32(0) {
		v4228 = v4183
		goto L1211
	} else {
		goto L1213
	}
L1213:
	;
	v4192 = v4164 + v4189
	v4193 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+184))
	if base.Ui32(v4193) <= base.Ui32(v4192) {
		goto L1215
	} else {
		goto L1216
	}
L1214:
	;
	v4203 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+192))
	if base.B2i32(v4202 == int32(0))|base.B2i32(base.Ui32(v4192) <= base.Ui32(v4202)) != 0 {
		v4228 = v4203
		goto L1211
	} else {
		goto L1218
	}
L1215:
	;
	v4195 = *(*int32)(unsafe.Add(mBase, uint32(v4188)+188))
	v4202 = v4195
	goto L1214
L1216:
	;
	goto L1217
L1217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4188)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4188)+184)) = v4189
	v4200 = F_strchr(m, v4189, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4188)+188)) = v4200
	v4202 = v4200
	goto L1214
L1218:
	;
	v4208 = v4202
	v4211 = v4203
	goto L1219
L1219:
	;
	v4213 = int32(1)
	v4214 = v4211 + v4213
	*(*int32)(unsafe.Add(mBase, uint32(v4188)+192)) = v4214
	v4217 = v4208 + v4213
	*(*int32)(unsafe.Add(mBase, uint32(v4188)+184)) = v4217
	v4220 = F_strchr(m, v4217, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4188)+188)) = v4220
	if v4220 == int32(0) {
		v4228 = v4214
		goto L1211
	} else {
		goto L1221
	}
L1220:
	;
	v4228 = v4214
	goto L1211
L1221:
	;
	if base.Ui32(v4220) < base.Ui32(v4192) {
		v4208 = v4220
		v4211 = v4214
		goto L1219
	} else {
		goto L1222
	}
L1222:
	;
	goto L1220
L1223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4179)+28)) = v4253
	if v4167 == int32(0) {
		goto L1227
	} else {
		goto L1228
	}
L1224:
	;
	v4247 = *(*int32)(unsafe.Add(mBase, uint32(v4173)+4))
	if v4247 != int32(1) {
		v4253 = v4173
		goto L1223
	} else {
		goto L1225
	}
L1225:
	;
	v4250 = *(*int32)(unsafe.Add(mBase, uint32(v4173)+12))
	v4251 = *(*int32)(unsafe.Add(mBase, uint32(v4250)))
	if v4251 != 0 {
		v4253 = v4173
		goto L1223
	} else {
		goto L1226
	}
L1226:
	;
	v4253 = int32(0)
	goto L1223
L1227:
	;
	m.G0 = v4176 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4179
	v9817 = v233
	goto L13
L1228:
	;
	v4258 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[41]))
	*(*int32)(unsafe.Add(mBase, uint32(v4176)+16)) = v4258
	v4261 = v4176 + int32(48)
	v4266 = F_pg_snprintf(m, v4261, int32(32), int32(_a_F_plpgsql_yyparse_54), v4176+int32(16))
	mBase = m.M
	v4267 = m.ExcPending
	if v4267 != 0 {
		goto L49
	} else {
		goto L1229
	}
L1229:
	;
	v4268 = int32(0)
	v4269 = *(*int32)(unsafe.Add(mBase, uint32(v4179)+4))
	v4274 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v4268, v4268)
	mBase = m.M
	v4275 = m.ExcPending
	if v4275 != 0 {
		goto L49
	} else {
		goto L1230
	}
L1230:
	;
	v4277 = F_plpgsql_build_variable(m, v4261, v4269, v4274, int32(1))
	mBase = m.M
	v4278 = m.ExcPending
	if v4278 != 0 {
		goto L49
	} else {
		goto L1231
	}
L1231:
	;
	v4279 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4179)+16)) = v4279
	if v4170 == int32(0) {
		goto L1227
	} else {
		goto L1232
	}
L1232:
	;
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(v4170)+4))
	if v4283 <= int32(0) {
		goto L1227
	} else {
		goto L1233
	}
L1233:
	;
	v4290 = v4268
	goto L1234
L1234:
	;
	v4311 = *(*int32)(unsafe.Add(mBase, uint32(v4170)+12))
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(v4311+v4290<<(uint(int32(2))%32))))
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(v4315)+4))
	v4318 = v4176 + int32(32)
	F_initStringInfo(m, v4318)
	mBase = m.M
	v4320 = m.ExcPending
	if v4320 != 0 {
		goto L49
	} else {
		goto L1236
	}
L1235:
	;
	goto L1227
L1236:
	;
	v4321 = *(*int32)(unsafe.Add(mBase, uint32(v4316)))
	*(*int32)(unsafe.Add(mBase, uint32(v4176)+4)) = v4321
	*(*int32)(unsafe.Add(mBase, uint32(v4176))) = v4176 + int32(48)
	F_appendStringInfo(m, v4318, int32(_a_F_plpgsql_yyparse_55), v4176)
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L49
	} else {
		goto L1237
	}
L1237:
	;
	v4329 = *(*int32)(unsafe.Add(mBase, uint32(v4316)))
	F_pfree(m, v4329)
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		goto L49
	} else {
		goto L1238
	}
L1238:
	;
	v4332 = *(*int32)(unsafe.Add(mBase, uint32(v4176)+32))
	v4333 = F_pstrdup(m, v4332)
	mBase = m.M
	v4334 = m.ExcPending
	if v4334 != 0 {
		goto L49
	} else {
		goto L1239
	}
L1239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4316))) = v4333
	v4337 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v4316)+12)) = v4337
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(v4176)+32))
	F_pfree(m, v4339)
	mBase = m.M
	v4341 = m.ExcPending
	if v4341 != 0 {
		goto L49
	} else {
		goto L1240
	}
L1240:
	;
	v4343 = v4290 + int32(1)
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(v4170)+4))
	if v4343 < v4344 {
		v4290 = v4343
		goto L1234
	} else {
		goto L1241
	}
L1241:
	;
	goto L1235
L1242:
	;
	if v4380 != int32(384) {
		goto L1243
	} else {
		goto L1244
	}
L1243:
	;
	F_plpgsql_push_back_token(m, v4380, v4377, v4379, l1)
	mBase = m.M
	v4385 = m.ExcPending
	if v4385 != 0 {
		goto L49
	} else {
		goto L1246
	}
L1244:
	;
	v4397 = int32(0)
	goto L1245
L1245:
	;
	F_plpgsql_push_back_token(m, int32(384), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4404 = m.ExcPending
	if v4404 != 0 {
		goto L49
	} else {
		goto L1248
	}
L1246:
	;
	v4387 = int32(0)
	v4391 = int32(1)
	v4395 = F_read_sql_construct(m, int32(384), v4387, v4387, int32(_a_F_plpgsql_yyparse_56), int32(2), v4391, v4391, v4387, v4387, v4377, v4379, l1)
	mBase = m.M
	v4396 = m.ExcPending
	if v4396 != 0 {
		goto L49
	} else {
		goto L1247
	}
L1247:
	;
	v4397 = v4395
	goto L1245
L1248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4397
	v9817 = v233
	goto L13
L1249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4410
	v9817 = v233
	goto L13
L1250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4419
	v9817 = v233
	goto L13
L1251:
	;
	v4427 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4428 = int32(0)
	if v4427 < v4428 {
		v4473 = v4428
		goto L1253
	} else {
		goto L1254
	}
L1252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4423))) = v4473
	v4478 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4423)+4)) = v4478
	v4480 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4423)+8)) = v4480
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4423
	v9817 = v233
	goto L13
L1253:
	;
	goto L1252
L1254:
	;
	v4433 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4434 = *(*int32)(unsafe.Add(mBase, uint32(v4433)+56))
	if v4434 == int32(0) {
		v4473 = v4428
		goto L1253
	} else {
		goto L1255
	}
L1255:
	;
	v4437 = v4427 + v4434
	v4438 = *(*int32)(unsafe.Add(mBase, uint32(v4433)+184))
	if base.Ui32(v4438) <= base.Ui32(v4437) {
		goto L1257
	} else {
		goto L1258
	}
L1256:
	;
	v4448 = *(*int32)(unsafe.Add(mBase, uint32(v4433)+192))
	if base.B2i32(v4447 == int32(0))|base.B2i32(base.Ui32(v4437) <= base.Ui32(v4447)) != 0 {
		v4473 = v4448
		goto L1253
	} else {
		goto L1260
	}
L1257:
	;
	v4440 = *(*int32)(unsafe.Add(mBase, uint32(v4433)+188))
	v4447 = v4440
	goto L1256
L1258:
	;
	goto L1259
L1259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4433)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4433)+184)) = v4434
	v4445 = F_strchr(m, v4434, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4433)+188)) = v4445
	v4447 = v4445
	goto L1256
L1260:
	;
	v4453 = v4447
	v4456 = v4448
	goto L1261
L1261:
	;
	v4458 = int32(1)
	v4459 = v4456 + v4458
	*(*int32)(unsafe.Add(mBase, uint32(v4433)+192)) = v4459
	v4462 = v4453 + v4458
	*(*int32)(unsafe.Add(mBase, uint32(v4433)+184)) = v4462
	v4465 = F_strchr(m, v4462, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4433)+188)) = v4465
	if v4465 == int32(0) {
		v4473 = v4459
		goto L1253
	} else {
		goto L1263
	}
L1262:
	;
	v4473 = v4459
	goto L1253
L1263:
	;
	if base.Ui32(v4465) < base.Ui32(v4437) {
		v4453 = v4465
		v4456 = v4459
		goto L1261
	} else {
		goto L1264
	}
L1264:
	;
	goto L1262
L1265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4485
	v9817 = v233
	goto L13
L1266:
	;
	goto L1267
L1267:
	;
	v4487 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+188)) = v4487
	*(*int32)(unsafe.Add(mBase, uint32(v28)+244)) = v4487
	v4494 = F_list_make1_impl(m, int32(1), v28+int32(188))
	mBase = m.M
	v4495 = m.ExcPending
	if v4495 != 0 {
		goto L49
	} else {
		goto L1268
	}
L1268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4494
	v9817 = v233
	goto L13
L1269:
	;
	v4500 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4498))) = v4500
	v4504 = *(*int32)(unsafe.Add(mBase, uint32(v147-v4500)))
	v4505 = int32(0)
	if v4504 < v4505 {
		v4550 = v4505
		goto L1271
	} else {
		goto L1272
	}
L1270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4498)+4)) = v4550
	v4554 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v4555 = *(*int32)(unsafe.Add(mBase, uint32(v4554)+520))
	v4557 = v4555 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4554)+520)) = v4557
	*(*int32)(unsafe.Add(mBase, uint32(v4498)+8)) = v4557
	v4561 = v148 - int32(32)
	v4562 = *(*int32)(unsafe.Add(mBase, uint32(v4561)))
	*(*int32)(unsafe.Add(mBase, uint32(v4498)+12)) = v4562
	v4564 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4498)+16)) = v4564
	v4566 = *(*int32)(unsafe.Add(mBase, uint32(v4561)))
	v4567 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v4568 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v4566, v4567, v4568, l1)
	mBase = m.M
	v4570 = m.ExcPending
	if v4570 != 0 {
		goto L49
	} else {
		goto L1283
	}
L1271:
	;
	goto L1270
L1272:
	;
	v4510 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4511 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+56))
	if v4511 == int32(0) {
		v4550 = v4505
		goto L1271
	} else {
		goto L1273
	}
L1273:
	;
	v4514 = v4504 + v4511
	v4515 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+184))
	if base.Ui32(v4515) <= base.Ui32(v4514) {
		goto L1275
	} else {
		goto L1276
	}
L1274:
	;
	v4525 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+192))
	if base.B2i32(v4524 == int32(0))|base.B2i32(base.Ui32(v4514) <= base.Ui32(v4524)) != 0 {
		v4550 = v4525
		goto L1271
	} else {
		goto L1278
	}
L1275:
	;
	v4517 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+188))
	v4524 = v4517
	goto L1274
L1276:
	;
	goto L1277
L1277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4510)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4510)+184)) = v4511
	v4522 = F_strchr(m, v4511, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4510)+188)) = v4522
	v4524 = v4522
	goto L1274
L1278:
	;
	v4530 = v4524
	v4533 = v4525
	goto L1279
L1279:
	;
	v4535 = int32(1)
	v4536 = v4533 + v4535
	*(*int32)(unsafe.Add(mBase, uint32(v4510)+192)) = v4536
	v4539 = v4530 + v4535
	*(*int32)(unsafe.Add(mBase, uint32(v4510)+184)) = v4539
	v4542 = F_strchr(m, v4539, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4510)+188)) = v4542
	if v4542 == int32(0) {
		v4550 = v4536
		goto L1271
	} else {
		goto L1281
	}
L1280:
	;
	v4550 = v4536
	goto L1271
L1281:
	;
	if base.Ui32(v4542) < base.Ui32(v4514) {
		v4530 = v4542
		v4533 = v4536
		goto L1279
	} else {
		goto L1282
	}
L1282:
	;
	goto L1280
L1283:
	;
	v4573 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v4574 = *(*int32)(unsafe.Add(mBase, uint32(v4573)))
	if v4574 != 0 {
		goto L1285
	} else {
		goto L1286
	}
L1284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4498
	v9817 = v233
	goto L13
L1285:
	;
	v4575 = v4573
	goto L1288
L1286:
	;
	v4578 = v4573
	goto L1287
L1287:
	;
	v4580 = *(*int32)(unsafe.Add(mBase, uint32(v4578)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = v4580
	goto L1284
L1288:
	;
	v4576 = *(*int32)(unsafe.Add(mBase, uint32(v4575)+8))
	v4577 = *(*int32)(unsafe.Add(mBase, uint32(v4576)))
	if v4577 != 0 {
		v4575 = v4576
		goto L1288
	} else {
		goto L1290
	}
L1289:
	;
	v4578 = v4576
	goto L1287
L1290:
	;
	goto L1289
L1291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4584))) = int32(5)
	v4590 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v4591 = int32(0)
	if v4590 < v4591 {
		v4636 = v4591
		goto L1293
	} else {
		goto L1294
	}
L1292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4584)+4)) = v4636
	v4640 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v4641 = *(*int32)(unsafe.Add(mBase, uint32(v4640)+520))
	v4643 = v4641 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4640)+520)) = v4643
	*(*int32)(unsafe.Add(mBase, uint32(v4584)+8)) = v4643
	v4647 = v148 - int32(48)
	v4648 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	*(*int32)(unsafe.Add(mBase, uint32(v4584)+12)) = v4648
	v4652 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4584)+16)) = v4652
	v4654 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4584)+20)) = v4654
	v4656 = *(*int32)(unsafe.Add(mBase, uint32(v4647)))
	v4657 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v4658 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v4656, v4657, v4658, l1)
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L49
	} else {
		goto L1305
	}
L1293:
	;
	goto L1292
L1294:
	;
	v4596 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4597 = *(*int32)(unsafe.Add(mBase, uint32(v4596)+56))
	if v4597 == int32(0) {
		v4636 = v4591
		goto L1293
	} else {
		goto L1295
	}
L1295:
	;
	v4600 = v4590 + v4597
	v4601 = *(*int32)(unsafe.Add(mBase, uint32(v4596)+184))
	if base.Ui32(v4601) <= base.Ui32(v4600) {
		goto L1297
	} else {
		goto L1298
	}
L1296:
	;
	v4611 = *(*int32)(unsafe.Add(mBase, uint32(v4596)+192))
	if base.B2i32(v4610 == int32(0))|base.B2i32(base.Ui32(v4600) <= base.Ui32(v4610)) != 0 {
		v4636 = v4611
		goto L1293
	} else {
		goto L1300
	}
L1297:
	;
	v4603 = *(*int32)(unsafe.Add(mBase, uint32(v4596)+188))
	v4610 = v4603
	goto L1296
L1298:
	;
	goto L1299
L1299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+184)) = v4597
	v4608 = F_strchr(m, v4597, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+188)) = v4608
	v4610 = v4608
	goto L1296
L1300:
	;
	v4616 = v4610
	v4619 = v4611
	goto L1301
L1301:
	;
	v4621 = int32(1)
	v4622 = v4619 + v4621
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+192)) = v4622
	v4625 = v4616 + v4621
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+184)) = v4625
	v4628 = F_strchr(m, v4625, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4596)+188)) = v4628
	if v4628 == int32(0) {
		v4636 = v4622
		goto L1293
	} else {
		goto L1303
	}
L1302:
	;
	v4636 = v4622
	goto L1293
L1303:
	;
	if base.Ui32(v4628) < base.Ui32(v4600) {
		v4616 = v4628
		v4619 = v4622
		goto L1301
	} else {
		goto L1304
	}
L1304:
	;
	goto L1302
L1305:
	;
	v4663 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v4664 = *(*int32)(unsafe.Add(mBase, uint32(v4663)))
	if v4664 != 0 {
		goto L1307
	} else {
		goto L1308
	}
L1306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4584
	v9817 = v233
	goto L13
L1307:
	;
	v4665 = v4663
	goto L1310
L1308:
	;
	v4668 = v4663
	goto L1309
L1309:
	;
	v4670 = *(*int32)(unsafe.Add(mBase, uint32(v4668)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = v4670
	goto L1306
L1310:
	;
	v4666 = *(*int32)(unsafe.Add(mBase, uint32(v4665)+8))
	v4667 = *(*int32)(unsafe.Add(mBase, uint32(v4666)))
	if v4667 != 0 {
		v4665 = v4666
		goto L1310
	} else {
		goto L1312
	}
L1311:
	;
	v4668 = v4666
	goto L1309
L1312:
	;
	goto L1311
L1313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+4)) = v4725
	v4729 = v148 - int32(48)
	v4730 = *(*int32)(unsafe.Add(mBase, uint32(v4729)))
	*(*int32)(unsafe.Add(mBase, uint32(v4675)+12)) = v4730
	if v4676 == int32(6) {
		goto L1326
	} else {
		goto L1327
	}
L1314:
	;
	goto L1313
L1315:
	;
	v4685 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4686 = *(*int32)(unsafe.Add(mBase, uint32(v4685)+56))
	if v4686 == int32(0) {
		v4725 = v4680
		goto L1314
	} else {
		goto L1316
	}
L1316:
	;
	v4689 = v4679 + v4686
	v4690 = *(*int32)(unsafe.Add(mBase, uint32(v4685)+184))
	if base.Ui32(v4690) <= base.Ui32(v4689) {
		goto L1318
	} else {
		goto L1319
	}
L1317:
	;
	v4700 = *(*int32)(unsafe.Add(mBase, uint32(v4685)+192))
	if base.B2i32(v4699 == int32(0))|base.B2i32(base.Ui32(v4689) <= base.Ui32(v4699)) != 0 {
		v4725 = v4700
		goto L1314
	} else {
		goto L1321
	}
L1318:
	;
	v4692 = *(*int32)(unsafe.Add(mBase, uint32(v4685)+188))
	v4699 = v4692
	goto L1317
L1319:
	;
	goto L1320
L1320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4685)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4685)+184)) = v4686
	v4697 = F_strchr(m, v4686, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4685)+188)) = v4697
	v4699 = v4697
	goto L1317
L1321:
	;
	v4705 = v4699
	v4708 = v4700
	goto L1322
L1322:
	;
	v4710 = int32(1)
	v4711 = v4708 + v4710
	*(*int32)(unsafe.Add(mBase, uint32(v4685)+192)) = v4711
	v4714 = v4705 + v4710
	*(*int32)(unsafe.Add(mBase, uint32(v4685)+184)) = v4714
	v4717 = F_strchr(m, v4714, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4685)+188)) = v4717
	if v4717 == int32(0) {
		v4725 = v4711
		goto L1314
	} else {
		goto L1324
	}
L1323:
	;
	v4725 = v4711
	goto L1314
L1324:
	;
	if base.Ui32(v4717) < base.Ui32(v4689) {
		v4705 = v4717
		v4708 = v4711
		goto L1322
	} else {
		goto L1325
	}
L1325:
	;
	goto L1323
L1326:
	;
	v4736 = int32(36)
	goto L1328
L1327:
	;
	v4736 = int32(20)
	goto L1328
L1328:
	;
	v4738 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v4675+v4736))) = v4738
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4675
	v4741 = *(*int32)(unsafe.Add(mBase, uint32(v4729)))
	v4742 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v4743 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v4741, v4742, v4743, l1)
	mBase = m.M
	v4745 = m.ExcPending
	if v4745 != 0 {
		goto L49
	} else {
		goto L1329
	}
L1329:
	;
	v4748 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v4749 = *(*int32)(unsafe.Add(mBase, uint32(v4748)))
	if v4749 != 0 {
		goto L1331
	} else {
		goto L1332
	}
L1330:
	;
	v9817 = v233
	goto L13
L1331:
	;
	v4750 = v4748
	goto L1334
L1332:
	;
	v4753 = v4748
	goto L1333
L1333:
	;
	v4755 = *(*int32)(unsafe.Add(mBase, uint32(v4753)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = v4755
	goto L1330
L1334:
	;
	v4751 = *(*int32)(unsafe.Add(mBase, uint32(v4750)+8))
	v4752 = *(*int32)(unsafe.Add(mBase, uint32(v4751)))
	if v4752 != 0 {
		v4750 = v4751
		goto L1334
	} else {
		goto L1336
	}
L1335:
	;
	v4753 = v4751
	goto L1333
L1336:
	;
	goto L1335
L1337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v4761
	if v4761 != int32(317) {
		goto L1341
	} else {
		goto L1342
	}
L1338:
	;
	if v4769&int32(1) == int32(0) {
		goto L18
	} else {
		goto L1383
	}
L1339:
	;
	if v4761 == int32(363) {
		goto L17
	} else {
		goto L1382
	}
L1340:
	;
	v4938 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v4939 = *(*int32)(unsafe.Add(mBase, uint32(v4938)))
	if v4939 != 0 {
		goto L1338
	} else {
		goto L1372
	}
L1341:
	;
	v4766 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	if v4761 != int32(277) {
		goto L1339
	} else {
		goto L1344
	}
L1342:
	;
	goto L1343
L1343:
	;
	v4807 = int32(0)
	v4810 = int32(1)
	v4819 = F_read_sql_construct(m, int32(336), int32(381), v4807, int32(_a_F_plpgsql_yyparse_57), int32(2), v4810, v4810, v4807, v28+int32(240), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4820 = m.ExcPending
	if v4820 != 0 {
		goto L49
	} else {
		goto L1355
	}
L1344:
	;
	v4769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v4769&int32(1) != 0 {
		goto L1340
	} else {
		goto L1345
	}
L1345:
	;
	v4772 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v4772 == int32(0) {
		goto L1340
	} else {
		goto L1346
	}
L1346:
	;
	v4775 = int32(_a_F_plpgsql_yyparse_58)
	v4778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4772))))
	v4781 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[42])))
	if base.B2i32(v4778 == int32(0))|base.B2i32(v4778 != v4781) != 0 {
		v4799 = v4778
		v4800 = v4781
		goto L1348
	} else {
		goto L1349
	}
L1347:
	;
	if v4799-v4800 != 0 {
		goto L1340
	} else {
		goto L1354
	}
L1348:
	;
	goto L1347
L1349:
	;
	v4784 = v4772
	v4785 = v4775
	goto L1350
L1350:
	;
	v4788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4785)+1)))
	v4789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4784)+1)))
	if v4789 == int32(0) {
		v4799 = v4789
		v4800 = v4788
		goto L1348
	} else {
		goto L1352
	}
L1351:
	;
	v4799 = v4789
	v4800 = v4788
	goto L1348
L1352:
	;
	v4792 = int32(1)
	if v4789 == v4788 {
		v4784 = v4784 + v4792
		v4785 = v4785 + v4792
		goto L1350
	} else {
		goto L1353
	}
L1353:
	;
	goto L1351
L1354:
	;
	goto L1343
L1355:
	;
	v4822 = F_palloc0(m, int32(32))
	mBase = m.M
	v4823 = m.ExcPending
	if v4823 != 0 {
		goto L49
	} else {
		goto L1356
	}
L1356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4822))) = int32(18)
	v4827 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v4828 = *(*int32)(unsafe.Add(mBase, uint32(v4827)+520))
	v4830 = v4828 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4827)+520)) = v4830
	*(*int32)(unsafe.Add(mBase, uint32(v4822)+8)) = v4830
	v4833 = int32(4)
	v4834 = v147 - v4833
	v4836 = v148 - v4833
	v4837 = *(*int32)(unsafe.Add(mBase, uint32(v4836)))
	if v4837 != 0 {
		goto L1358
	} else {
		goto L1359
	}
L1357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4822)+24)) = v4819
	v4860 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	if v4860 == int32(381) {
		goto L1364
	} else {
		goto L1365
	}
L1358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4822)+16)) = v4837
	v4839 = *(*int32)(unsafe.Add(mBase, uint32(v4836)))
	v4840 = *(*int32)(unsafe.Add(mBase, uint32(v4834)))
	F_check_assignable(m, v4839, v4840, l1)
	mBase = m.M
	v4842 = m.ExcPending
	if v4842 != 0 {
		goto L49
	} else {
		goto L1361
	}
L1359:
	;
	goto L1360
L1360:
	;
	v4845 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(8))))
	if v4845 == int32(0) {
		goto L32
	} else {
		goto L1362
	}
L1361:
	;
	goto L1357
L1362:
	;
	v4850 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4853 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v4854 = *(*int32)(unsafe.Add(mBase, uint32(v4834)))
	v4855 = F_make_scalar_list1(m, v4850, v4845, v4853, v4854, l1)
	mBase = m.M
	v4856 = m.ExcPending
	if v4856 != 0 {
		goto L49
	} else {
		goto L1363
	}
L1363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4822)+16)) = v4855
	goto L1357
L1364:
	;
	goto L1367
L1365:
	;
	goto L1366
L1366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4822
	v9817 = v233
	goto L13
L1367:
	;
	v4890 = int32(0)
	v4893 = int32(1)
	v4902 = F_read_sql_construct(m, int32(44), int32(336), v4890, int32(_a_F_plpgsql_yyparse_59), int32(2), v4893, v4893, v4890, v28+int32(240), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4903 = m.ExcPending
	if v4903 != 0 {
		goto L49
	} else {
		goto L1369
	}
L1368:
	;
	goto L1366
L1369:
	;
	v4904 = *(*int32)(unsafe.Add(mBase, uint32(v4822)+28))
	v4905 = F_lappend(m, v4904, v4902)
	mBase = m.M
	v4906 = m.ExcPending
	if v4906 != 0 {
		goto L49
	} else {
		goto L1370
	}
L1370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4822)+28)) = v4905
	v4908 = *(*int32)(unsafe.Add(mBase, uint32(v28)+240))
	if v4908 == int32(44) {
		goto L1367
	} else {
		goto L1371
	}
L1371:
	;
	goto L1368
L1372:
	;
	v4940 = *(*int32)(unsafe.Add(mBase, uint32(v4938)+24))
	v4941 = *(*int32)(unsafe.Add(mBase, uint32(v4940)+4))
	if v4941 != int32(1790) {
		goto L1338
	} else {
		goto L1373
	}
L1373:
	;
	v4945 = F_palloc0(m, int32(32))
	mBase = m.M
	v4946 = m.ExcPending
	if v4946 != 0 {
		goto L49
	} else {
		goto L1374
	}
L1374:
	;
	v4947 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4945))) = v4947
	v4950 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v4951 = *(*int32)(unsafe.Add(mBase, uint32(v4950)+520))
	v4953 = v4951 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4950)+520)) = v4953
	*(*int32)(unsafe.Add(mBase, uint32(v4945)+8)) = v4953
	v4956 = *(*int32)(unsafe.Add(mBase, uint32(v4938)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4945)+24)) = v4956
	v4960 = *(*int32)(unsafe.Add(mBase, uint32(v148-v4947)))
	if v4960 != 0 {
		goto L1375
	} else {
		goto L1376
	}
L1375:
	;
	v4963 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(4))))
	if v4963 != 0 {
		goto L31
	} else {
		goto L1378
	}
L1376:
	;
	goto L1377
L1377:
	;
	v4964 = *(*int32)(unsafe.Add(mBase, uint32(v4938)+28))
	if v4964 == int32(0) {
		goto L30
	} else {
		goto L1379
	}
L1378:
	;
	goto L1377
L1379:
	;
	v4972 = F_read_cursor_args(m, v4938, int32(336), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v4973 = m.ExcPending
	if v4973 != 0 {
		goto L49
	} else {
		goto L1380
	}
L1380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4945)+28)) = v4972
	v4977 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v4980 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v4984 = F_plpgsql_build_record(m, v4977, v4980, int32(0), int32(2249), int32(1))
	mBase = m.M
	v4985 = m.ExcPending
	if v4985 != 0 {
		goto L49
	} else {
		goto L1381
	}
L1381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4945)+16)) = v4984
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4945
	v9817 = v233
	goto L13
L1382:
	;
	goto L16
L1383:
	;
	goto L16
L1384:
	;
	v4998 = v4994
	goto L1386
L1385:
	;
	v4995 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v4996 = F_NameListToString(m, v4995)
	mBase = m.M
	v4997 = m.ExcPending
	if v4997 != 0 {
		goto L49
	} else {
		goto L1387
	}
L1386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v4998
	v5000 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5001 = int32(0)
	if v5000 < v5001 {
		v5046 = v5001
		goto L1389
	} else {
		goto L1390
	}
L1387:
	;
	v4998 = v4996
	goto L1386
L1388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v5046
	v5049 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v5050 = *(*int32)(unsafe.Add(mBase, uint32(v5049)))
	v5051 = int32(1)
	if base.Ui32(v5050-v5051) <= base.Ui32(v5051) {
		goto L1401
	} else {
		goto L1402
	}
L1389:
	;
	goto L1388
L1390:
	;
	v5006 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5007 = *(*int32)(unsafe.Add(mBase, uint32(v5006)+56))
	if v5007 == int32(0) {
		v5046 = v5001
		goto L1389
	} else {
		goto L1391
	}
L1391:
	;
	v5010 = v5000 + v5007
	v5011 = *(*int32)(unsafe.Add(mBase, uint32(v5006)+184))
	if base.Ui32(v5011) <= base.Ui32(v5010) {
		goto L1393
	} else {
		goto L1394
	}
L1392:
	;
	v5021 = *(*int32)(unsafe.Add(mBase, uint32(v5006)+192))
	if base.B2i32(v5020 == int32(0))|base.B2i32(base.Ui32(v5010) <= base.Ui32(v5020)) != 0 {
		v5046 = v5021
		goto L1389
	} else {
		goto L1396
	}
L1393:
	;
	v5013 = *(*int32)(unsafe.Add(mBase, uint32(v5006)+188))
	v5020 = v5013
	goto L1392
L1394:
	;
	goto L1395
L1395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+184)) = v5007
	v5018 = F_strchr(m, v5007, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+188)) = v5018
	v5020 = v5018
	goto L1392
L1396:
	;
	v5026 = v5020
	v5029 = v5021
	goto L1397
L1397:
	;
	v5031 = int32(1)
	v5032 = v5029 + v5031
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+192)) = v5032
	v5035 = v5026 + v5031
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+184)) = v5035
	v5038 = F_strchr(m, v5035, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5006)+188)) = v5038
	if v5038 == int32(0) {
		v5046 = v5032
		goto L1389
	} else {
		goto L1399
	}
L1398:
	;
	v5046 = v5032
	goto L1389
L1399:
	;
	if base.Ui32(v5038) < base.Ui32(v5010) {
		v5026 = v5038
		v5029 = v5032
		goto L1397
	} else {
		goto L1400
	}
L1400:
	;
	goto L1398
L1401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+280)) = int32(0)
	v5057 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+284)) = v5057
	v9817 = v233
	goto L13
L1402:
	;
	goto L1403
L1403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+284)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+280)) = v5049
	v5063 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v5065 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5066 = F_plpgsql_yylex(m, v5063, v5065, l1)
	mBase = m.M
	v5067 = m.ExcPending
	if v5067 != 0 {
		goto L49
	} else {
		goto L1404
	}
L1404:
	;
	F_plpgsql_push_back_token(m, v5066, v5063, v5065, l1)
	mBase = m.M
	v5069 = m.ExcPending
	if v5069 != 0 {
		goto L49
	} else {
		goto L1405
	}
L1405:
	;
	if v5066 != int32(44) {
		v9817 = v233
		goto L13
	} else {
		goto L1406
	}
L1406:
	;
	v5072 = *(*int32)(unsafe.Add(mBase, uint32(v28)+272))
	v5073 = *(*int32)(unsafe.Add(mBase, uint32(v28)+280))
	v5074 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5075 = F_read_into_scalar_list(m, v5072, v5073, v5074, v5063, v5065, l1)
	mBase = m.M
	v5076 = m.ExcPending
	if v5076 != 0 {
		goto L49
	} else {
		goto L1407
	}
L1407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+284)) = v5075
	v9817 = v233
	goto L13
L1408:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+280)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+276)) = v5126
	v5132 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v5134 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5135 = F_plpgsql_yylex(m, v5132, v5134, l1)
	mBase = m.M
	v5136 = m.ExcPending
	if v5136 != 0 {
		goto L49
	} else {
		goto L1421
	}
L1409:
	;
	goto L1408
L1410:
	;
	v5086 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5087 = *(*int32)(unsafe.Add(mBase, uint32(v5086)+56))
	if v5087 == int32(0) {
		v5126 = v5081
		goto L1409
	} else {
		goto L1411
	}
L1411:
	;
	v5090 = v5080 + v5087
	v5091 = *(*int32)(unsafe.Add(mBase, uint32(v5086)+184))
	if base.Ui32(v5091) <= base.Ui32(v5090) {
		goto L1413
	} else {
		goto L1414
	}
L1412:
	;
	v5101 = *(*int32)(unsafe.Add(mBase, uint32(v5086)+192))
	if base.B2i32(v5100 == int32(0))|base.B2i32(base.Ui32(v5090) <= base.Ui32(v5100)) != 0 {
		v5126 = v5101
		goto L1409
	} else {
		goto L1416
	}
L1413:
	;
	v5093 = *(*int32)(unsafe.Add(mBase, uint32(v5086)+188))
	v5100 = v5093
	goto L1412
L1414:
	;
	goto L1415
L1415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5086)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5086)+184)) = v5087
	v5098 = F_strchr(m, v5087, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5086)+188)) = v5098
	v5100 = v5098
	goto L1412
L1416:
	;
	v5106 = v5100
	v5109 = v5101
	goto L1417
L1417:
	;
	v5111 = int32(1)
	v5112 = v5109 + v5111
	*(*int32)(unsafe.Add(mBase, uint32(v5086)+192)) = v5112
	v5115 = v5106 + v5111
	*(*int32)(unsafe.Add(mBase, uint32(v5086)+184)) = v5115
	v5118 = F_strchr(m, v5115, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5086)+188)) = v5118
	if v5118 == int32(0) {
		v5126 = v5112
		goto L1409
	} else {
		goto L1419
	}
L1418:
	;
	v5126 = v5112
	goto L1409
L1419:
	;
	if base.Ui32(v5118) < base.Ui32(v5090) {
		v5106 = v5118
		v5109 = v5112
		goto L1417
	} else {
		goto L1420
	}
L1420:
	;
	goto L1418
L1421:
	;
	F_plpgsql_push_back_token(m, v5135, v5132, v5134, l1)
	mBase = m.M
	v5138 = m.ExcPending
	if v5138 != 0 {
		goto L49
	} else {
		goto L1422
	}
L1422:
	;
	if v5135 != int32(44) {
		v9817 = v233
		goto L13
	} else {
		goto L1423
	}
L1423:
	;
	goto L6
L1424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5142))) = int32(9)
	v5148 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(24))))
	v5149 = int32(0)
	if v5148 < v5149 {
		v5194 = v5149
		goto L1426
	} else {
		goto L1427
	}
L1425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5142)+4)) = v5194
	v5198 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5199 = *(*int32)(unsafe.Add(mBase, uint32(v5198)+520))
	v5201 = v5199 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5198)+520)) = v5201
	*(*int32)(unsafe.Add(mBase, uint32(v5142)+8)) = v5201
	v5205 = v148 - int32(112)
	v5206 = *(*int32)(unsafe.Add(mBase, uint32(v5205)))
	*(*int32)(unsafe.Add(mBase, uint32(v5142)+12)) = v5206
	v5210 = *(*int32)(unsafe.Add(mBase, uint32(v148+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v5142)+20)) = v5210
	v5214 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v5142)+24)) = v5214
	v5216 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v5142)+28)) = v5216
	v5219 = v148 - int32(68)
	v5220 = *(*int32)(unsafe.Add(mBase, uint32(v5219)))
	if v5220 == int32(0) {
		goto L1438
	} else {
		goto L1439
	}
L1426:
	;
	goto L1425
L1427:
	;
	v5154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5155 = *(*int32)(unsafe.Add(mBase, uint32(v5154)+56))
	if v5155 == int32(0) {
		v5194 = v5149
		goto L1426
	} else {
		goto L1428
	}
L1428:
	;
	v5158 = v5148 + v5155
	v5159 = *(*int32)(unsafe.Add(mBase, uint32(v5154)+184))
	if base.Ui32(v5159) <= base.Ui32(v5158) {
		goto L1430
	} else {
		goto L1431
	}
L1429:
	;
	v5169 = *(*int32)(unsafe.Add(mBase, uint32(v5154)+192))
	if base.B2i32(v5168 == int32(0))|base.B2i32(base.Ui32(v5158) <= base.Ui32(v5168)) != 0 {
		v5194 = v5169
		goto L1426
	} else {
		goto L1433
	}
L1430:
	;
	v5161 = *(*int32)(unsafe.Add(mBase, uint32(v5154)+188))
	v5168 = v5161
	goto L1429
L1431:
	;
	goto L1432
L1432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5154)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5154)+184)) = v5155
	v5166 = F_strchr(m, v5155, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5154)+188)) = v5166
	v5168 = v5166
	goto L1429
L1433:
	;
	v5174 = v5168
	v5177 = v5169
	goto L1434
L1434:
	;
	v5179 = int32(1)
	v5180 = v5177 + v5179
	*(*int32)(unsafe.Add(mBase, uint32(v5154)+192)) = v5180
	v5183 = v5174 + v5179
	*(*int32)(unsafe.Add(mBase, uint32(v5154)+184)) = v5183
	v5186 = F_strchr(m, v5183, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5154)+188)) = v5186
	if v5186 == int32(0) {
		v5194 = v5180
		goto L1426
	} else {
		goto L1436
	}
L1435:
	;
	v5194 = v5180
	goto L1426
L1436:
	;
	if base.Ui32(v5186) < base.Ui32(v5158) {
		v5174 = v5186
		v5177 = v5180
		goto L1434
	} else {
		goto L1437
	}
L1437:
	;
	goto L1435
L1438:
	;
	v5224 = v148 - int32(72)
	v5225 = *(*int32)(unsafe.Add(mBase, uint32(v5224)))
	if v5225 == int32(0) {
		goto L29
	} else {
		goto L1441
	}
L1439:
	;
	v5228 = v5220
	v5229 = v5219
	goto L1440
L1440:
	;
	v5230 = *(*int32)(unsafe.Add(mBase, uint32(v5228)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5142)+16)) = v5230
	v5232 = *(*int32)(unsafe.Add(mBase, uint32(v5229)))
	v5235 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(20))))
	F_check_assignable(m, v5232, v5235, l1)
	mBase = m.M
	v5237 = m.ExcPending
	if v5237 != 0 {
		goto L49
	} else {
		goto L1442
	}
L1441:
	;
	v5228 = v5225
	v5229 = v5224
	goto L1440
L1442:
	;
	v5238 = *(*int32)(unsafe.Add(mBase, uint32(v5205)))
	v5239 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v5240 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	F_check_labels(m, v5238, v5239, v5240, l1)
	mBase = m.M
	v5242 = m.ExcPending
	if v5242 != 0 {
		goto L49
	} else {
		goto L1443
	}
L1443:
	;
	v5245 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v5246 = *(*int32)(unsafe.Add(mBase, uint32(v5245)))
	if v5246 != 0 {
		goto L1445
	} else {
		goto L1446
	}
L1444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5142
	v9817 = v233
	goto L13
L1445:
	;
	v5247 = v5245
	goto L1448
L1446:
	;
	v5250 = v5245
	goto L1447
L1447:
	;
	v5252 = *(*int32)(unsafe.Add(mBase, uint32(v5250)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13])) = v5252
	goto L1444
L1448:
	;
	v5248 = *(*int32)(unsafe.Add(mBase, uint32(v5247)+8))
	v5249 = *(*int32)(unsafe.Add(mBase, uint32(v5248)))
	if v5249 != 0 {
		v5247 = v5248
		goto L1448
	} else {
		goto L1450
	}
L1449:
	;
	v5250 = v5248
	goto L1447
L1450:
	;
	goto L1449
L1451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5260))) = int32(10)
	v5265 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5266 = *(*int32)(unsafe.Add(mBase, uint32(v5265)+520))
	v5268 = v5266 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5265)+520)) = v5268
	*(*int32)(unsafe.Add(mBase, uint32(v5260)+8)) = v5268
	v5273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148-int32(32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5260)+12)) = uint8(v5273)
	v5276 = v147 - int32(8)
	v5277 = *(*int32)(unsafe.Add(mBase, uint32(v5276)))
	v5278 = int32(0)
	if v5277 < v5278 {
		v5323 = v5278
		goto L1453
	} else {
		goto L1454
	}
L1452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5260)+4)) = v5323
	v5327 = v148 - int32(16)
	v5328 = *(*int32)(unsafe.Add(mBase, uint32(v5327)))
	*(*int32)(unsafe.Add(mBase, uint32(v5260)+16)) = v5328
	v5330 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v5260)+20)) = v5330
	v5333 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[13]))
	v5334 = *(*int32)(unsafe.Add(mBase, uint32(v5327)))
	if v5334 != 0 {
		goto L1466
	} else {
		goto L1467
	}
L1453:
	;
	goto L1452
L1454:
	;
	v5283 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5284 = *(*int32)(unsafe.Add(mBase, uint32(v5283)+56))
	if v5284 == int32(0) {
		v5323 = v5278
		goto L1453
	} else {
		goto L1455
	}
L1455:
	;
	v5287 = v5277 + v5284
	v5288 = *(*int32)(unsafe.Add(mBase, uint32(v5283)+184))
	if base.Ui32(v5288) <= base.Ui32(v5287) {
		goto L1457
	} else {
		goto L1458
	}
L1456:
	;
	v5298 = *(*int32)(unsafe.Add(mBase, uint32(v5283)+192))
	if base.B2i32(v5297 == int32(0))|base.B2i32(base.Ui32(v5287) <= base.Ui32(v5297)) != 0 {
		v5323 = v5298
		goto L1453
	} else {
		goto L1460
	}
L1457:
	;
	v5290 = *(*int32)(unsafe.Add(mBase, uint32(v5283)+188))
	v5297 = v5290
	goto L1456
L1458:
	;
	goto L1459
L1459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5283)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5283)+184)) = v5284
	v5295 = F_strchr(m, v5284, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5283)+188)) = v5295
	v5297 = v5295
	goto L1456
L1460:
	;
	v5303 = v5297
	v5306 = v5298
	goto L1461
L1461:
	;
	v5308 = int32(1)
	v5309 = v5306 + v5308
	*(*int32)(unsafe.Add(mBase, uint32(v5283)+192)) = v5309
	v5312 = v5303 + v5308
	*(*int32)(unsafe.Add(mBase, uint32(v5283)+184)) = v5312
	v5315 = F_strchr(m, v5312, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5283)+188)) = v5315
	if v5315 == int32(0) {
		v5323 = v5309
		goto L1453
	} else {
		goto L1463
	}
L1462:
	;
	v5323 = v5309
	goto L1453
L1463:
	;
	if base.Ui32(v5315) < base.Ui32(v5287) {
		v5303 = v5315
		v5306 = v5309
		goto L1461
	} else {
		goto L1464
	}
L1464:
	;
	goto L1462
L1465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5260
	v9817 = v233
	goto L13
L1466:
	;
	v5335 = *(*int32)(unsafe.Add(mBase, uint32(v5327)))
	if v5333 != 0 {
		goto L1470
	} else {
		goto L1471
	}
L1467:
	;
	goto L1468
L1468:
	;
	if v5333 != 0 {
		goto L1488
	} else {
		goto L1489
	}
L1469:
	;
	if v5348 == int32(0) {
		goto L28
	} else {
		goto L1479
	}
L1470:
	;
	v5336 = v5333
	goto L1473
L1471:
	;
	goto L1472
L1472:
	;
	v5348 = int32(0)
	goto L1469
L1473:
	;
	v5338 = *(*int32)(unsafe.Add(mBase, uint32(v5336)))
	if v5338 != 0 {
		goto L1475
	} else {
		goto L1476
	}
L1474:
	;
	goto L1472
L1475:
	;
	v5342 = *(*int32)(unsafe.Add(mBase, uint32(v5336)+8))
	if v5342 != 0 {
		v5336 = v5342
		goto L1473
	} else {
		goto L1478
	}
L1476:
	;
	v5341 = F_strcmp(m, v5336+int32(12), v5335)
	mBase = m.M
	if v5341 != 0 {
		goto L1475
	} else {
		goto L1477
	}
L1477:
	;
	v5348 = v5336
	goto L1469
L1478:
	;
	goto L1474
L1479:
	;
	v5351 = *(*int32)(unsafe.Add(mBase, uint32(v5348)+4))
	if v5351 == int32(1) {
		goto L1465
	} else {
		goto L1480
	}
L1480:
	;
	v5354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5260)+12)))
	if v5354 != 0 {
		goto L1465
	} else {
		goto L1481
	}
L1481:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5358 = m.ExcPending
	if v5358 != 0 {
		goto L49
	} else {
		goto L1482
	}
L1482:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v5361 = m.ExcPending
	if v5361 != 0 {
		goto L49
	} else {
		goto L1483
	}
L1483:
	;
	v5362 = *(*int32)(unsafe.Add(mBase, uint32(v5327)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+208)) = v5362
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_60), v28+int32(208))
	mBase = m.M
	v5368 = m.ExcPending
	if v5368 != 0 {
		goto L49
	} else {
		goto L1484
	}
L1484:
	;
	v5371 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v5372 = F_plpgsql_scanner_errposition(m, v5371, l1)
	mBase = m.M
	v5373 = m.ExcPending
	if v5373 != 0 {
		goto L49
	} else {
		goto L1485
	}
L1485:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1753), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v5378 = m.ExcPending
	if v5378 != 0 {
		goto L49
	} else {
		goto L1486
	}
L1486:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1487:
	;
	if v5388 == int32(0) {
		goto L27
	} else {
		goto L1497
	}
L1488:
	;
	v5379 = v5333
	goto L1491
L1489:
	;
	goto L1490
L1490:
	;
	v5388 = int32(0)
	goto L1487
L1491:
	;
	v5380 = *(*int32)(unsafe.Add(mBase, uint32(v5379)))
	if v5380 != 0 {
		goto L1493
	} else {
		goto L1494
	}
L1492:
	;
	goto L1490
L1493:
	;
	v5384 = *(*int32)(unsafe.Add(mBase, uint32(v5379)+8))
	if v5384 != 0 {
		v5379 = v5384
		goto L1491
	} else {
		goto L1496
	}
L1494:
	;
	v5381 = *(*int32)(unsafe.Add(mBase, uint32(v5379)+4))
	if v5381 != int32(1) {
		goto L1493
	} else {
		goto L1495
	}
L1495:
	;
	v5388 = v5379
	goto L1487
L1496:
	;
	goto L1492
L1497:
	;
	goto L1465
L1498:
	;
	v5859 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v5861 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, v5401, v5859, v5861, l1)
	mBase = m.M
	v5863 = m.ExcPending
	if v5863 != 0 {
		goto L49
	} else {
		goto L1616
	}
L1499:
	;
	v5623 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5624 = m.G0
	v5626 = v5624 - int32(16)
	m.G0 = v5626
	v5629 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5629)+63)))
	if v5630 != 0 {
		goto L1569
	} else {
		goto L1570
	}
L1500:
	;
	v5595 = int32(_a_F_plpgsql_yyparse_61)
	v5598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5410))))
	v5601 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[43])))
	if base.B2i32(v5598 == int32(0))|base.B2i32(v5598 != v5601) != 0 {
		v5619 = v5598
		v5620 = v5601
		goto L1561
	} else {
		goto L1562
	}
L1501:
	;
	v5441 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5443 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v5445 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5447 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5447)+63)))
	if v5448 != 0 {
		goto L1520
	} else {
		goto L1521
	}
L1502:
	;
	if v5401 != int32(277) {
		goto L1504
	} else {
		goto L1505
	}
L1503:
	;
	switch v5401 - int32(341) {
	case 0:
		goto L1501
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		goto L1498
	case 17:
		goto L1499
	default:
		goto L1502
	}
L1504:
	;
	if v5401 != 0 {
		goto L1498
	} else {
		goto L1507
	}
L1505:
	;
	goto L1506
L1506:
	;
	v5407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v5407&int32(1) != 0 {
		goto L1498
	} else {
		goto L1508
	}
L1507:
	;
	goto L4
L1508:
	;
	v5410 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v5410 == int32(0) {
		goto L1498
	} else {
		goto L1509
	}
L1509:
	;
	v5413 = int32(_a_F_plpgsql_yyparse_62)
	v5416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5410))))
	v5419 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[44])))
	if base.B2i32(v5416 == int32(0))|base.B2i32(v5416 != v5419) != 0 {
		v5437 = v5416
		v5438 = v5419
		goto L1511
	} else {
		goto L1512
	}
L1510:
	;
	if v5437-v5438 != 0 {
		goto L1500
	} else {
		goto L1517
	}
L1511:
	;
	goto L1510
L1512:
	;
	v5422 = v5410
	v5423 = v5413
	goto L1513
L1513:
	;
	v5426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5423)+1)))
	v5427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5422)+1)))
	if v5427 == int32(0) {
		v5437 = v5427
		v5438 = v5426
		goto L1511
	} else {
		goto L1515
	}
L1514:
	;
	v5437 = v5427
	v5438 = v5426
	goto L1511
L1515:
	;
	v5430 = int32(1)
	if v5427 == v5426 {
		v5422 = v5422 + v5430
		v5423 = v5423 + v5430
		goto L1513
	} else {
		goto L1516
	}
L1516:
	;
	goto L1514
L1517:
	;
	goto L1501
L1518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5450
	v9817 = v233
	goto L13
L1519:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5576 = m.ExcPending
	if v5576 != 0 {
		goto L49
	} else {
		goto L1555
	}
L1520:
	;
	v5450 = F_palloc0(m, int32(20))
	mBase = m.M
	v5451 = m.ExcPending
	if v5451 != 0 {
		goto L49
	} else {
		goto L1523
	}
L1521:
	;
	goto L1522
L1522:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5558 = m.ExcPending
	if v5558 != 0 {
		goto L49
	} else {
		goto L1550
	}
L1523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5450))) = int32(12)
	v5454 = int32(0)
	if v5441 < v5454 {
		v5499 = v5454
		goto L1525
	} else {
		goto L1526
	}
L1524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5450)+4)) = v5499
	v5503 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5504 = *(*int32)(unsafe.Add(mBase, uint32(v5503)+520))
	v5506 = v5504 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5503)+520)) = v5506
	*(*int64)(unsafe.Add(mBase, uint32(v5450)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5450)+8)) = v5506
	v5511 = *(*int32)(unsafe.Add(mBase, uint32(v5503)+472))
	v5512 = F_plpgsql_yylex(m, v5443, v5445, l1)
	mBase = m.M
	v5513 = m.ExcPending
	if v5513 != 0 {
		goto L49
	} else {
		goto L1537
	}
L1525:
	;
	goto L1524
L1526:
	;
	v5459 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5460 = *(*int32)(unsafe.Add(mBase, uint32(v5459)+56))
	if v5460 == int32(0) {
		v5499 = v5454
		goto L1525
	} else {
		goto L1527
	}
L1527:
	;
	v5463 = v5441 + v5460
	v5464 = *(*int32)(unsafe.Add(mBase, uint32(v5459)+184))
	if base.Ui32(v5464) <= base.Ui32(v5463) {
		goto L1529
	} else {
		goto L1530
	}
L1528:
	;
	v5474 = *(*int32)(unsafe.Add(mBase, uint32(v5459)+192))
	if base.B2i32(v5473 == int32(0))|base.B2i32(base.Ui32(v5463) <= base.Ui32(v5473)) != 0 {
		v5499 = v5474
		goto L1525
	} else {
		goto L1532
	}
L1529:
	;
	v5466 = *(*int32)(unsafe.Add(mBase, uint32(v5459)+188))
	v5473 = v5466
	goto L1528
L1530:
	;
	goto L1531
L1531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5459)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5459)+184)) = v5460
	v5471 = F_strchr(m, v5460, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5459)+188)) = v5471
	v5473 = v5471
	goto L1528
L1532:
	;
	v5479 = v5473
	v5482 = v5474
	goto L1533
L1533:
	;
	v5484 = int32(1)
	v5485 = v5482 + v5484
	*(*int32)(unsafe.Add(mBase, uint32(v5459)+192)) = v5485
	v5488 = v5479 + v5484
	*(*int32)(unsafe.Add(mBase, uint32(v5459)+184)) = v5488
	v5491 = F_strchr(m, v5488, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5459)+188)) = v5491
	if v5491 == int32(0) {
		v5499 = v5485
		goto L1525
	} else {
		goto L1535
	}
L1534:
	;
	v5499 = v5485
	goto L1525
L1535:
	;
	if base.Ui32(v5491) < base.Ui32(v5463) {
		v5479 = v5491
		v5482 = v5485
		goto L1533
	} else {
		goto L1536
	}
L1536:
	;
	goto L1534
L1537:
	;
	if int32(0) <= v5511 {
		goto L1538
	} else {
		goto L1539
	}
L1538:
	;
	if v5512 != int32(59) {
		goto L1519
	} else {
		goto L1541
	}
L1539:
	;
	goto L1540
L1540:
	;
	if v5512 != int32(277) {
		goto L1542
	} else {
		goto L1543
	}
L1541:
	;
	v5519 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5520 = *(*int32)(unsafe.Add(mBase, uint32(v5519)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v5450)+16)) = v5520
	goto L1518
L1542:
	;
	F_plpgsql_push_back_token(m, v5512, v5443, v5445, l1)
	mBase = m.M
	v5542 = m.ExcPending
	if v5542 != 0 {
		goto L49
	} else {
		goto L1548
	}
L1543:
	;
	v5524 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v5525 = m.ExcPending
	if v5525 != 0 {
		goto L49
	} else {
		goto L1544
	}
L1544:
	;
	if v5524 != int32(59) {
		goto L1542
	} else {
		goto L1545
	}
L1545:
	;
	v5528 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v5529 = *(*int32)(unsafe.Add(mBase, uint32(v5528)))
	if base.B2i32(base.Ui32(int32(4)) < base.Ui32(v5529))|base.B2i32(v5529 == int32(3)) != 0 {
		goto L1542
	} else {
		goto L1546
	}
L1546:
	;
	v5535 = *(*int32)(unsafe.Add(mBase, uint32(v5528)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5450)+16)) = v5535
	v5537 = F_plpgsql_yylex(m, v5443, v5445, l1)
	mBase = m.M
	v5538 = m.ExcPending
	if v5538 != 0 {
		goto L49
	} else {
		goto L1547
	}
L1547:
	;
	goto L1518
L1548:
	;
	v5544 = int32(0)
	v5548 = int32(1)
	v5552 = F_read_sql_construct(m, int32(59), v5544, v5544, int32(_a_F_plpgsql_yyparse_16), int32(2), v5548, v5548, v5544, v5544, v5443, v5445, l1)
	mBase = m.M
	v5553 = m.ExcPending
	if v5553 != 0 {
		goto L49
	} else {
		goto L1549
	}
L1549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5450)+12)) = v5552
	goto L1518
L1550:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5561 = m.ExcPending
	if v5561 != 0 {
		goto L49
	} else {
		goto L1551
	}
L1551:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_63), int32(0))
	mBase = m.M
	v5565 = m.ExcPending
	if v5565 != 0 {
		goto L49
	} else {
		goto L1552
	}
L1552:
	;
	v5566 = F_plpgsql_scanner_errposition(m, v5441, l1)
	mBase = m.M
	v5567 = m.ExcPending
	if v5567 != 0 {
		goto L49
	} else {
		goto L1553
	}
L1553:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3449), int32(_a_F_plpgsql_yyparse_64))
	mBase = m.M
	v5572 = m.ExcPending
	if v5572 != 0 {
		goto L49
	} else {
		goto L1554
	}
L1554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1555:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5579 = m.ExcPending
	if v5579 != 0 {
		goto L49
	} else {
		goto L1556
	}
L1556:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_65), int32(0))
	mBase = m.M
	v5583 = m.ExcPending
	if v5583 != 0 {
		goto L49
	} else {
		goto L1557
	}
L1557:
	;
	v5584 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v5585 = F_plpgsql_scanner_errposition(m, v5584, l1)
	mBase = m.M
	v5586 = m.ExcPending
	if v5586 != 0 {
		goto L49
	} else {
		goto L1558
	}
L1558:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3464), int32(_a_F_plpgsql_yyparse_64))
	mBase = m.M
	v5591 = m.ExcPending
	if v5591 != 0 {
		goto L49
	} else {
		goto L1559
	}
L1559:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1560:
	;
	if v5619-v5620 != 0 {
		goto L1498
	} else {
		goto L1567
	}
L1561:
	;
	goto L1560
L1562:
	;
	v5604 = v5410
	v5605 = v5595
	goto L1563
L1563:
	;
	v5608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5605)+1)))
	v5609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5604)+1)))
	if v5609 == int32(0) {
		v5619 = v5609
		v5620 = v5608
		goto L1561
	} else {
		goto L1565
	}
L1564:
	;
	v5619 = v5609
	v5620 = v5608
	goto L1561
L1565:
	;
	v5612 = int32(1)
	if v5609 == v5608 {
		v5604 = v5604 + v5612
		v5605 = v5605 + v5612
		goto L1563
	} else {
		goto L1566
	}
L1566:
	;
	goto L1564
L1567:
	;
	goto L1499
L1568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5636
	v9817 = v233
	goto L13
L1569:
	;
	v5632 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v5634 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v5636 = F_palloc0(m, int32(24))
	mBase = m.M
	v5637 = m.ExcPending
	if v5637 != 0 {
		goto L49
	} else {
		goto L1572
	}
L1570:
	;
	goto L1571
L1571:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5841 = m.ExcPending
	if v5841 != 0 {
		goto L49
	} else {
		goto L1611
	}
L1572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5636))) = int32(13)
	v5640 = int32(0)
	if v5623 < v5640 {
		v5685 = v5640
		goto L1574
	} else {
		goto L1575
	}
L1573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5636)+4)) = v5685
	v5689 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5690 = *(*int32)(unsafe.Add(mBase, uint32(v5689)+520))
	v5692 = v5690 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5689)+520)) = v5692
	*(*int32)(unsafe.Add(mBase, uint32(v5636)+8)) = v5692
	v5695 = F_plpgsql_yylex(m, v5632, v5634, l1)
	mBase = m.M
	v5696 = m.ExcPending
	if v5696 != 0 {
		goto L49
	} else {
		goto L1588
	}
L1574:
	;
	goto L1573
L1575:
	;
	v5645 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5646 = *(*int32)(unsafe.Add(mBase, uint32(v5645)+56))
	if v5646 == int32(0) {
		v5685 = v5640
		goto L1574
	} else {
		goto L1576
	}
L1576:
	;
	v5649 = v5623 + v5646
	v5650 = *(*int32)(unsafe.Add(mBase, uint32(v5645)+184))
	if base.Ui32(v5650) <= base.Ui32(v5649) {
		goto L1578
	} else {
		goto L1579
	}
L1577:
	;
	v5660 = *(*int32)(unsafe.Add(mBase, uint32(v5645)+192))
	if base.B2i32(v5659 == int32(0))|base.B2i32(base.Ui32(v5649) <= base.Ui32(v5659)) != 0 {
		v5685 = v5660
		goto L1574
	} else {
		goto L1581
	}
L1578:
	;
	v5652 = *(*int32)(unsafe.Add(mBase, uint32(v5645)+188))
	v5659 = v5652
	goto L1577
L1579:
	;
	goto L1580
L1580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5645)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5645)+184)) = v5646
	v5657 = F_strchr(m, v5646, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5645)+188)) = v5657
	v5659 = v5657
	goto L1577
L1581:
	;
	v5665 = v5659
	v5668 = v5660
	goto L1582
L1582:
	;
	v5670 = int32(1)
	v5671 = v5668 + v5670
	*(*int32)(unsafe.Add(mBase, uint32(v5645)+192)) = v5671
	v5674 = v5665 + v5670
	*(*int32)(unsafe.Add(mBase, uint32(v5645)+184)) = v5674
	v5677 = F_strchr(m, v5674, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5645)+188)) = v5677
	if v5677 == int32(0) {
		v5685 = v5671
		goto L1574
	} else {
		goto L1584
	}
L1583:
	;
	v5685 = v5671
	goto L1574
L1584:
	;
	if base.Ui32(v5677) < base.Ui32(v5649) {
		v5665 = v5677
		v5668 = v5671
		goto L1582
	} else {
		goto L1585
	}
L1585:
	;
	goto L1583
L1586:
	;
	m.G0 = v5626 + int32(16)
	goto L1568
L1587:
	;
	v5752 = int32(0)
	v5755 = int32(1)
	v5760 = F_read_sql_construct(m, int32(59), int32(381), v5752, int32(_a_F_plpgsql_yyparse_66), int32(2), v5755, v5755, v5752, v5626+int32(12), v5632, v5634, l1)
	mBase = m.M
	v5761 = m.ExcPending
	if v5761 != 0 {
		goto L49
	} else {
		goto L1604
	}
L1588:
	;
	if v5695 == int32(317) {
		goto L1587
	} else {
		goto L1589
	}
L1589:
	;
	if v5695 != int32(277) {
		goto L1590
	} else {
		goto L1591
	}
L1590:
	;
	F_plpgsql_push_back_token(m, v5695, v5632, v5634, l1)
	mBase = m.M
	v5736 = m.ExcPending
	if v5736 != 0 {
		goto L49
	} else {
		goto L1602
	}
L1591:
	;
	v5701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v5701 != 0 {
		goto L1590
	} else {
		goto L1592
	}
L1592:
	;
	v5702 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v5702 == int32(0) {
		goto L1590
	} else {
		goto L1593
	}
L1593:
	;
	v5705 = int32(_a_F_plpgsql_yyparse_58)
	v5708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5702))))
	v5711 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[42])))
	if base.B2i32(v5708 == int32(0))|base.B2i32(v5708 != v5711) != 0 {
		v5729 = v5708
		v5730 = v5711
		goto L1595
	} else {
		goto L1596
	}
L1594:
	;
	if v5729-v5730 == int32(0) {
		goto L1587
	} else {
		goto L1601
	}
L1595:
	;
	goto L1594
L1596:
	;
	v5714 = v5702
	v5715 = v5705
	goto L1597
L1597:
	;
	v5718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5715)+1)))
	v5719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5714)+1)))
	if v5719 == int32(0) {
		v5729 = v5719
		v5730 = v5718
		goto L1595
	} else {
		goto L1599
	}
L1598:
	;
	v5729 = v5719
	v5730 = v5718
	goto L1595
L1599:
	;
	v5722 = int32(1)
	if v5719 == v5718 {
		v5714 = v5714 + v5722
		v5715 = v5715 + v5722
		goto L1597
	} else {
		goto L1600
	}
L1600:
	;
	goto L1598
L1601:
	;
	goto L1590
L1602:
	;
	v5738 = int32(0)
	v5746 = F_read_sql_construct(m, int32(59), v5738, v5738, int32(_a_F_plpgsql_yyparse_16), v5738, v5738, int32(1), v5738, v5738, v5632, v5634, l1)
	mBase = m.M
	v5747 = m.ExcPending
	if v5747 != 0 {
		goto L49
	} else {
		goto L1603
	}
L1603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5636)+12)) = v5746
	goto L1586
L1604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5636)+16)) = v5760
	v5763 = *(*int32)(unsafe.Add(mBase, uint32(v5626)+12))
	if v5763 != int32(381) {
		goto L1586
	} else {
		goto L1605
	}
L1605:
	;
	goto L1606
L1606:
	;
	v5793 = int32(0)
	v5796 = int32(1)
	v5801 = F_read_sql_construct(m, int32(44), int32(59), v5793, int32(_a_F_plpgsql_yyparse_67), int32(2), v5796, v5796, v5793, v5626+int32(12), v5632, v5634, l1)
	mBase = m.M
	v5802 = m.ExcPending
	if v5802 != 0 {
		goto L49
	} else {
		goto L1608
	}
L1607:
	;
	goto L1586
L1608:
	;
	v5803 = *(*int32)(unsafe.Add(mBase, uint32(v5636)+20))
	v5804 = F_lappend(m, v5803, v5801)
	mBase = m.M
	v5805 = m.ExcPending
	if v5805 != 0 {
		goto L49
	} else {
		goto L1609
	}
L1609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5636)+20)) = v5804
	v5807 = *(*int32)(unsafe.Add(mBase, uint32(v5626)+12))
	if v5807 == int32(44) {
		goto L1606
	} else {
		goto L1610
	}
L1610:
	;
	goto L1607
L1611:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5844 = m.ExcPending
	if v5844 != 0 {
		goto L49
	} else {
		goto L1612
	}
L1612:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_68), int32(0))
	mBase = m.M
	v5848 = m.ExcPending
	if v5848 != 0 {
		goto L49
	} else {
		goto L1613
	}
L1613:
	;
	v5849 = F_plpgsql_scanner_errposition(m, v5623, l1)
	mBase = m.M
	v5850 = m.ExcPending
	if v5850 != 0 {
		goto L49
	} else {
		goto L1614
	}
L1614:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3513), int32(_a_F_plpgsql_yyparse_69))
	mBase = m.M
	v5855 = m.ExcPending
	if v5855 != 0 {
		goto L49
	} else {
		goto L1615
	}
L1615:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1616:
	;
	v5864 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v5866 = F_palloc0(m, int32(20))
	mBase = m.M
	v5867 = m.ExcPending
	if v5867 != 0 {
		goto L49
	} else {
		goto L1617
	}
L1617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866))) = int32(11)
	v5870 = int32(0)
	if v5864 < v5870 {
		v5915 = v5870
		goto L1619
	} else {
		goto L1620
	}
L1618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+4)) = v5915
	v5919 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5920 = *(*int32)(unsafe.Add(mBase, uint32(v5919)+520))
	v5921 = int32(1)
	v5922 = v5920 + v5921
	*(*int32)(unsafe.Add(mBase, uint32(v5919)+520)) = v5922
	*(*int64)(unsafe.Add(mBase, uint32(v5866)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+8)) = v5922
	v5927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5919)+63)))
	if v5927 == v5921 {
		goto L1632
	} else {
		goto L1633
	}
L1619:
	;
	goto L1618
L1620:
	;
	v5875 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5876 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+56))
	if v5876 == int32(0) {
		v5915 = v5870
		goto L1619
	} else {
		goto L1621
	}
L1621:
	;
	v5879 = v5864 + v5876
	v5880 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+184))
	if base.Ui32(v5880) <= base.Ui32(v5879) {
		goto L1623
	} else {
		goto L1624
	}
L1622:
	;
	v5890 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+192))
	if base.B2i32(v5889 == int32(0))|base.B2i32(base.Ui32(v5879) <= base.Ui32(v5889)) != 0 {
		v5915 = v5890
		goto L1619
	} else {
		goto L1626
	}
L1623:
	;
	v5882 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+188))
	v5889 = v5882
	goto L1622
L1624:
	;
	goto L1625
L1625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+184)) = v5876
	v5887 = F_strchr(m, v5876, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+188)) = v5887
	v5889 = v5887
	goto L1622
L1626:
	;
	v5895 = v5889
	v5898 = v5890
	goto L1627
L1627:
	;
	v5900 = int32(1)
	v5901 = v5898 + v5900
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+192)) = v5901
	v5904 = v5895 + v5900
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+184)) = v5904
	v5907 = F_strchr(m, v5904, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+188)) = v5907
	if v5907 == int32(0) {
		v5915 = v5901
		goto L1619
	} else {
		goto L1629
	}
L1628:
	;
	v5915 = v5901
	goto L1619
L1629:
	;
	if base.Ui32(v5907) < base.Ui32(v5879) {
		v5895 = v5907
		v5898 = v5901
		goto L1627
	} else {
		goto L1630
	}
L1630:
	;
	goto L1628
L1631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v5866
	v9817 = v233
	goto L13
L1632:
	;
	v5930 = F_plpgsql_yylex(m, v5859, v5861, l1)
	mBase = m.M
	v5931 = m.ExcPending
	if v5931 != 0 {
		goto L49
	} else {
		goto L1635
	}
L1633:
	;
	goto L1634
L1634:
	;
	v5957 = *(*int32)(unsafe.Add(mBase, uint32(v5919)+52))
	if v5957 == int32(2278) {
		goto L1645
	} else {
		goto L1646
	}
L1635:
	;
	if v5930 == int32(59) {
		goto L1631
	} else {
		goto L1636
	}
L1636:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5937 = m.ExcPending
	if v5937 != 0 {
		goto L49
	} else {
		goto L1637
	}
L1637:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5940 = m.ExcPending
	if v5940 != 0 {
		goto L49
	} else {
		goto L1638
	}
L1638:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_70), int32(0))
	mBase = m.M
	v5944 = m.ExcPending
	if v5944 != 0 {
		goto L49
	} else {
		goto L1639
	}
L1639:
	;
	F_errhint(m, int32(_a_F_plpgsql_yyparse_71), int32(0))
	mBase = m.M
	v5948 = m.ExcPending
	if v5948 != 0 {
		goto L49
	} else {
		goto L1640
	}
L1640:
	;
	v5949 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v5950 = F_plpgsql_scanner_errposition(m, v5949, l1)
	mBase = m.M
	v5951 = m.ExcPending
	if v5951 != 0 {
		goto L49
	} else {
		goto L1641
	}
L1641:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3377), int32(_a_F_plpgsql_yyparse_72))
	mBase = m.M
	v5956 = m.ExcPending
	if v5956 != 0 {
		goto L49
	} else {
		goto L1642
	}
L1642:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1643:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v6050 = m.ExcPending
	if v6050 != 0 {
		goto L49
	} else {
		goto L1673
	}
L1644:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6034 = m.ExcPending
	if v6034 != 0 {
		goto L49
	} else {
		goto L1669
	}
L1645:
	;
	v5960 = F_plpgsql_yylex(m, v5859, v5861, l1)
	mBase = m.M
	v5961 = m.ExcPending
	if v5961 != 0 {
		goto L49
	} else {
		goto L1648
	}
L1646:
	;
	goto L1647
L1647:
	;
	v5988 = *(*int32)(unsafe.Add(mBase, uint32(v5919)+472))
	v5989 = F_plpgsql_yylex(m, v5859, v5861, l1)
	mBase = m.M
	v5990 = m.ExcPending
	if v5990 != 0 {
		goto L49
	} else {
		goto L1656
	}
L1648:
	;
	if v5960 == int32(59) {
		goto L1631
	} else {
		goto L1649
	}
L1649:
	;
	v5965 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5965)+65)))
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v5970 = m.ExcPending
	if v5970 != 0 {
		goto L49
	} else {
		goto L1650
	}
L1650:
	;
	if v5966 == int32(112) {
		goto L1644
	} else {
		goto L1651
	}
L1651:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v5975 = m.ExcPending
	if v5975 != 0 {
		goto L49
	} else {
		goto L1652
	}
L1652:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_73), int32(0))
	mBase = m.M
	v5979 = m.ExcPending
	if v5979 != 0 {
		goto L49
	} else {
		goto L1653
	}
L1653:
	;
	v5980 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v5981 = F_plpgsql_scanner_errposition(m, v5980, l1)
	mBase = m.M
	v5982 = m.ExcPending
	if v5982 != 0 {
		goto L49
	} else {
		goto L1654
	}
L1654:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3392), int32(_a_F_plpgsql_yyparse_72))
	mBase = m.M
	v5987 = m.ExcPending
	if v5987 != 0 {
		goto L49
	} else {
		goto L1655
	}
L1655:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1656:
	;
	if int32(0) <= v5988 {
		goto L1657
	} else {
		goto L1658
	}
L1657:
	;
	if v5989 != int32(59) {
		goto L1643
	} else {
		goto L1660
	}
L1658:
	;
	goto L1659
L1659:
	;
	if v5989 != int32(277) {
		goto L1661
	} else {
		goto L1662
	}
L1660:
	;
	v5996 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v5997 = *(*int32)(unsafe.Add(mBase, uint32(v5996)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+16)) = v5997
	goto L1631
L1661:
	;
	F_plpgsql_push_back_token(m, v5989, v5859, v5861, l1)
	mBase = m.M
	v6019 = m.ExcPending
	if v6019 != 0 {
		goto L49
	} else {
		goto L1667
	}
L1662:
	;
	v6001 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v6002 = m.ExcPending
	if v6002 != 0 {
		goto L49
	} else {
		goto L1663
	}
L1663:
	;
	if v6001 != int32(59) {
		goto L1661
	} else {
		goto L1664
	}
L1664:
	;
	v6005 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6006 = *(*int32)(unsafe.Add(mBase, uint32(v6005)))
	if base.B2i32(base.Ui32(int32(4)) < base.Ui32(v6006))|base.B2i32(v6006 == int32(3)) != 0 {
		goto L1661
	} else {
		goto L1665
	}
L1665:
	;
	v6012 = *(*int32)(unsafe.Add(mBase, uint32(v6005)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+16)) = v6012
	v6014 = F_plpgsql_yylex(m, v5859, v5861, l1)
	mBase = m.M
	v6015 = m.ExcPending
	if v6015 != 0 {
		goto L49
	} else {
		goto L1666
	}
L1666:
	;
	goto L1631
L1667:
	;
	v6021 = int32(0)
	v6025 = int32(1)
	v6029 = F_read_sql_construct(m, int32(59), v6021, v6021, int32(_a_F_plpgsql_yyparse_16), int32(2), v6025, v6025, v6021, v6021, v5859, v5861, l1)
	mBase = m.M
	v6030 = m.ExcPending
	if v6030 != 0 {
		goto L49
	} else {
		goto L1668
	}
L1668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+12)) = v6029
	goto L1631
L1669:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_74), int32(0))
	mBase = m.M
	v6038 = m.ExcPending
	if v6038 != 0 {
		goto L49
	} else {
		goto L1670
	}
L1670:
	;
	v6039 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v6040 = F_plpgsql_scanner_errposition(m, v6039, l1)
	mBase = m.M
	v6041 = m.ExcPending
	if v6041 != 0 {
		goto L49
	} else {
		goto L1671
	}
L1671:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3387), int32(_a_F_plpgsql_yyparse_72))
	mBase = m.M
	v6046 = m.ExcPending
	if v6046 != 0 {
		goto L49
	} else {
		goto L1672
	}
L1672:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1673:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v6053 = m.ExcPending
	if v6053 != 0 {
		goto L49
	} else {
		goto L1674
	}
L1674:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_75), int32(0))
	mBase = m.M
	v6057 = m.ExcPending
	if v6057 != 0 {
		goto L49
	} else {
		goto L1675
	}
L1675:
	;
	v6058 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	v6059 = F_plpgsql_scanner_errposition(m, v6058, l1)
	mBase = m.M
	v6060 = m.ExcPending
	if v6060 != 0 {
		goto L49
	} else {
		goto L1676
	}
L1676:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(3401), int32(_a_F_plpgsql_yyparse_72))
	mBase = m.M
	v6065 = m.ExcPending
	if v6065 != 0 {
		goto L49
	} else {
		goto L1677
	}
L1677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1678:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071))) = int32(14)
	v6075 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v6076 = int32(0)
	if v6075 < v6076 {
		v6121 = v6076
		goto L1680
	} else {
		goto L1681
	}
L1679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+4)) = v6121
	v6125 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v6126 = *(*int32)(unsafe.Add(mBase, uint32(v6125)+520))
	v6128 = v6126 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6125)+520)) = v6128
	v6130 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6071)+16)) = v6130
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+12)) = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+8)) = v6128
	*(*int64)(unsafe.Add(mBase, uint32(v6071)+24)) = v6130
	v6138 = v6071 + int32(16)
	v6142 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v6143 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v6142, l1)
	mBase = m.M
	v6144 = m.ExcPending
	if v6144 != 0 {
		goto L49
	} else {
		goto L1716
	}
L1680:
	;
	goto L1679
L1681:
	;
	v6081 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6082 = *(*int32)(unsafe.Add(mBase, uint32(v6081)+56))
	if v6082 == int32(0) {
		v6121 = v6076
		goto L1680
	} else {
		goto L1682
	}
L1682:
	;
	v6085 = v6075 + v6082
	v6086 = *(*int32)(unsafe.Add(mBase, uint32(v6081)+184))
	if base.Ui32(v6086) <= base.Ui32(v6085) {
		goto L1684
	} else {
		goto L1685
	}
L1683:
	;
	v6096 = *(*int32)(unsafe.Add(mBase, uint32(v6081)+192))
	if base.B2i32(v6095 == int32(0))|base.B2i32(base.Ui32(v6085) <= base.Ui32(v6095)) != 0 {
		v6121 = v6096
		goto L1680
	} else {
		goto L1687
	}
L1684:
	;
	v6088 = *(*int32)(unsafe.Add(mBase, uint32(v6081)+188))
	v6095 = v6088
	goto L1683
L1685:
	;
	goto L1686
L1686:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6081)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6081)+184)) = v6082
	v6093 = F_strchr(m, v6082, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6081)+188)) = v6093
	v6095 = v6093
	goto L1683
L1687:
	;
	v6101 = v6095
	v6104 = v6096
	goto L1688
L1688:
	;
	v6106 = int32(1)
	v6107 = v6104 + v6106
	*(*int32)(unsafe.Add(mBase, uint32(v6081)+192)) = v6107
	v6110 = v6101 + v6106
	*(*int32)(unsafe.Add(mBase, uint32(v6081)+184)) = v6110
	v6113 = F_strchr(m, v6110, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6081)+188)) = v6113
	if v6113 == int32(0) {
		v6121 = v6107
		goto L1680
	} else {
		goto L1690
	}
L1689:
	;
	v6121 = v6107
	goto L1680
L1690:
	;
	if base.Ui32(v6113) < base.Ui32(v6085) {
		v6101 = v6113
		v6104 = v6107
		goto L1688
	} else {
		goto L1691
	}
L1691:
	;
	goto L1689
L1692:
	;
	v7081 = *(*int32)(unsafe.Add(mBase, uint32(v6071)+20))
	if v7081 != 0 {
		goto L1959
	} else {
		goto L1960
	}
L1693:
	;
	v6681 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v6683 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v6685 = m.G0
	v6687 = v6685 - int32(16)
	m.G0 = v6687
	v6693 = int32(0)
	goto L1848
L1694:
	;
	if v6632 != int32(381) {
		goto L1692
	} else {
		goto L1843
	}
L1695:
	;
	goto L1838
L1696:
	;
	v6567 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v6568 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v6567, l1)
	mBase = m.M
	v6569 = m.ExcPending
	if v6569 != 0 {
		goto L49
	} else {
		goto L1835
	}
L1697:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6138))) = v6557
	v6560 = F_plpgsql_recognize_err_condition(m, v6557, int32(0))
	mBase = m.M
	v6561 = m.ExcPending
	if v6561 != 0 {
		goto L49
	} else {
		goto L1834
	}
L1698:
	;
	v6554 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6557 = v6554
	goto L1697
L1699:
	;
	v6536 = int32(0)
	goto L1824
L1700:
	;
	v6436 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6437 = m.ExcPending
	if v6437 != 0 {
		goto L49
	} else {
		goto L1800
	}
L1701:
	;
	v6396 = int32(277)
	v6397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v6397&int32(1) != 0 {
		v6531 = v6396
		goto L1699
	} else {
		goto L1790
	}
L1702:
	;
	switch v6377 - int32(261) {
	case 0:
		goto L1782
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15:
		v6531 = v6377
		goto L1699
	case 14:
		goto L1698
	case 16:
		goto L1701
	default:
		goto L1783
	}
L1703:
	;
	v6374 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6375 = m.ExcPending
	if v6375 != 0 {
		goto L49
	} else {
		goto L1781
	}
L1704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+12)) = int32(14)
	goto L1703
L1705:
	;
	v6339 = int32(_a_F_plpgsql_yyparse_76)
	v6342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6306))))
	v6345 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[45])))
	if base.B2i32(v6342 == int32(0))|base.B2i32(v6342 != v6345) != 0 {
		v6363 = v6342
		v6364 = v6345
		goto L1774
	} else {
		goto L1775
	}
L1706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+12)) = int32(15)
	goto L1703
L1707:
	;
	v6306 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6306 == int32(0) {
		goto L1701
	} else {
		goto L1764
	}
L1708:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+12)) = int32(17)
	goto L1703
L1709:
	;
	v6276 = int32(_a_F_plpgsql_yyparse_77)
	v6279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6237))))
	v6282 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[46])))
	if base.B2i32(v6279 == int32(0))|base.B2i32(v6279 != v6282) != 0 {
		v6300 = v6279
		v6301 = v6282
		goto L1757
	} else {
		goto L1758
	}
L1710:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+12)) = int32(18)
	v6274 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6275 = m.ExcPending
	if v6275 != 0 {
		goto L49
	} else {
		goto L1755
	}
L1711:
	;
	v6237 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6237 == int32(0) {
		goto L1701
	} else {
		goto L1746
	}
L1712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+12)) = int32(19)
	v6235 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6236 = m.ExcPending
	if v6236 != 0 {
		goto L49
	} else {
		goto L1745
	}
L1713:
	;
	v6201 = int32(_a_F_plpgsql_yyparse_78)
	v6204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6162))))
	v6207 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[47])))
	if base.B2i32(v6204 == int32(0))|base.B2i32(v6204 != v6207) != 0 {
		v6225 = v6204
		v6226 = v6207
		goto L1738
	} else {
		goto L1739
	}
L1714:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+12)) = int32(21)
	v6199 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6200 = m.ExcPending
	if v6200 != 0 {
		goto L49
	} else {
		goto L1736
	}
L1715:
	;
	v6159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v6159&int32(1) != 0 {
		goto L1701
	} else {
		goto L1726
	}
L1716:
	;
	if v6143 <= int32(303) {
		goto L1717
	} else {
		goto L1718
	}
L1717:
	;
	if v6143 == int32(59) {
		goto L1692
	} else {
		goto L1720
	}
L1718:
	;
	goto L1719
L1719:
	;
	switch v6143 - int32(304) {
	case 0:
		goto L1704
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30, 32, 33, 34, 35, 36, 37, 38, 39:
		v6377 = v6143
		goto L1702
	case 12:
		goto L1714
	case 26:
		goto L1708
	case 31:
		goto L1706
	case 40:
		goto L1710
	default:
		goto L1724
	}
L1720:
	;
	if v6143 == int32(277) {
		goto L1715
	} else {
		goto L1721
	}
L1721:
	;
	if v6143 != 0 {
		v6377 = v6143
		goto L1702
	} else {
		goto L1722
	}
L1722:
	;
	F_plpgsql_yyerror(m, v6142, int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v6154 = m.ExcPending
	if v6154 != 0 {
		goto L49
	} else {
		goto L1723
	}
L1723:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1724:
	;
	if v6143 == int32(383) {
		goto L1712
	} else {
		goto L1725
	}
L1725:
	;
	v6377 = v6143
	goto L1702
L1726:
	;
	v6162 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6162 == int32(0) {
		goto L1701
	} else {
		goto L1727
	}
L1727:
	;
	v6165 = int32(_a_F_plpgsql_yyparse_79)
	v6168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6162))))
	v6171 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[48])))
	if base.B2i32(v6168 == int32(0))|base.B2i32(v6168 != v6171) != 0 {
		v6189 = v6168
		v6190 = v6171
		goto L1729
	} else {
		goto L1730
	}
L1728:
	;
	if v6189-v6190 != 0 {
		goto L1713
	} else {
		goto L1735
	}
L1729:
	;
	goto L1728
L1730:
	;
	v6174 = v6162
	v6175 = v6165
	goto L1731
L1731:
	;
	v6178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6175)+1)))
	v6179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6174)+1)))
	if v6179 == int32(0) {
		v6189 = v6179
		v6190 = v6178
		goto L1729
	} else {
		goto L1733
	}
L1732:
	;
	v6189 = v6179
	v6190 = v6178
	goto L1729
L1733:
	;
	v6182 = int32(1)
	if v6179 == v6178 {
		v6174 = v6174 + v6182
		v6175 = v6175 + v6182
		goto L1731
	} else {
		goto L1734
	}
L1734:
	;
	goto L1732
L1735:
	;
	goto L1714
L1736:
	;
	v6377 = v6199
	goto L1702
L1737:
	;
	if v6225-v6226 != 0 {
		goto L1711
	} else {
		goto L1744
	}
L1738:
	;
	goto L1737
L1739:
	;
	v6210 = v6162
	v6211 = v6201
	goto L1740
L1740:
	;
	v6214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6211)+1)))
	v6215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6210)+1)))
	if v6215 == int32(0) {
		v6225 = v6215
		v6226 = v6214
		goto L1738
	} else {
		goto L1742
	}
L1741:
	;
	v6225 = v6215
	v6226 = v6214
	goto L1738
L1742:
	;
	v6218 = int32(1)
	if v6215 == v6214 {
		v6210 = v6210 + v6218
		v6211 = v6211 + v6218
		goto L1740
	} else {
		goto L1743
	}
L1743:
	;
	goto L1741
L1744:
	;
	goto L1712
L1745:
	;
	v6377 = v6235
	goto L1702
L1746:
	;
	v6240 = int32(_a_F_plpgsql_yyparse_80)
	v6243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6237))))
	v6246 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[49])))
	if base.B2i32(v6243 == int32(0))|base.B2i32(v6243 != v6246) != 0 {
		v6264 = v6243
		v6265 = v6246
		goto L1748
	} else {
		goto L1749
	}
L1747:
	;
	if v6264-v6265 != 0 {
		goto L1709
	} else {
		goto L1754
	}
L1748:
	;
	goto L1747
L1749:
	;
	v6249 = v6237
	v6250 = v6240
	goto L1750
L1750:
	;
	v6253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6250)+1)))
	v6254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6249)+1)))
	if v6254 == int32(0) {
		v6264 = v6254
		v6265 = v6253
		goto L1748
	} else {
		goto L1752
	}
L1751:
	;
	v6264 = v6254
	v6265 = v6253
	goto L1748
L1752:
	;
	v6257 = int32(1)
	if v6254 == v6253 {
		v6249 = v6249 + v6257
		v6250 = v6250 + v6257
		goto L1750
	} else {
		goto L1753
	}
L1753:
	;
	goto L1751
L1754:
	;
	goto L1710
L1755:
	;
	v6377 = v6274
	goto L1702
L1756:
	;
	if v6300-v6301 != 0 {
		goto L1707
	} else {
		goto L1763
	}
L1757:
	;
	goto L1756
L1758:
	;
	v6285 = v6237
	v6286 = v6276
	goto L1759
L1759:
	;
	v6289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6286)+1)))
	v6290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6285)+1)))
	if v6290 == int32(0) {
		v6300 = v6290
		v6301 = v6289
		goto L1757
	} else {
		goto L1761
	}
L1760:
	;
	v6300 = v6290
	v6301 = v6289
	goto L1757
L1761:
	;
	v6293 = int32(1)
	if v6290 == v6289 {
		v6285 = v6285 + v6293
		v6286 = v6286 + v6293
		goto L1759
	} else {
		goto L1762
	}
L1762:
	;
	goto L1760
L1763:
	;
	goto L1708
L1764:
	;
	v6309 = int32(_a_F_plpgsql_yyparse_81)
	v6312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6306))))
	v6315 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[50])))
	if base.B2i32(v6312 == int32(0))|base.B2i32(v6312 != v6315) != 0 {
		v6333 = v6312
		v6334 = v6315
		goto L1766
	} else {
		goto L1767
	}
L1765:
	;
	if v6333-v6334 != 0 {
		goto L1705
	} else {
		goto L1772
	}
L1766:
	;
	goto L1765
L1767:
	;
	v6318 = v6306
	v6319 = v6309
	goto L1768
L1768:
	;
	v6322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6319)+1)))
	v6323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6318)+1)))
	if v6323 == int32(0) {
		v6333 = v6323
		v6334 = v6322
		goto L1766
	} else {
		goto L1770
	}
L1769:
	;
	v6333 = v6323
	v6334 = v6322
	goto L1766
L1770:
	;
	v6326 = int32(1)
	if v6323 == v6322 {
		v6318 = v6318 + v6326
		v6319 = v6319 + v6326
		goto L1768
	} else {
		goto L1771
	}
L1771:
	;
	goto L1769
L1772:
	;
	goto L1706
L1773:
	;
	if v6363-v6364 != 0 {
		goto L1701
	} else {
		goto L1780
	}
L1774:
	;
	goto L1773
L1775:
	;
	v6348 = v6306
	v6349 = v6339
	goto L1776
L1776:
	;
	v6352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6349)+1)))
	v6353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6348)+1)))
	if v6353 == int32(0) {
		v6363 = v6353
		v6364 = v6352
		goto L1774
	} else {
		goto L1778
	}
L1777:
	;
	v6363 = v6353
	v6364 = v6352
	goto L1774
L1778:
	;
	v6356 = int32(1)
	if v6353 == v6352 {
		v6348 = v6348 + v6356
		v6349 = v6349 + v6356
		goto L1776
	} else {
		goto L1779
	}
L1779:
	;
	goto L1777
L1780:
	;
	goto L1704
L1781:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v6374
	v6377 = v6374
	goto L1702
L1782:
	;
	v6382 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+20)) = v6382
	v6388 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6389 = m.ExcPending
	if v6389 != 0 {
		goto L49
	} else {
		goto L1786
	}
L1783:
	;
	switch v6377 - int32(371) {
	case 0:
		goto L1700
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v6531 = v6377
		goto L1699
	case 10:
		goto L1693
	default:
		goto L1784
	}
L1784:
	;
	if v6377 != 0 {
		v6531 = v6377
		goto L1699
	} else {
		goto L1785
	}
L1785:
	;
	goto L4
L1786:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v6388
	switch v6388 - int32(44) {
	case 0:
		goto L1695
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L1787
	case 15:
		v6632 = v6388
		goto L1694
	default:
		goto L1788
	}
L1787:
	;
	goto L3
L1788:
	;
	if v6388 == int32(381) {
		v6632 = v6388
		goto L1694
	} else {
		goto L1789
	}
L1789:
	;
	goto L1787
L1790:
	;
	v6400 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6400 == int32(0) {
		v6531 = v6396
		goto L1699
	} else {
		goto L1791
	}
L1791:
	;
	v6403 = int32(_a_F_plpgsql_yyparse_17)
	v6406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6400))))
	v6409 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[12])))
	if base.B2i32(v6406 == int32(0))|base.B2i32(v6406 != v6409) != 0 {
		v6427 = v6406
		v6428 = v6409
		goto L1793
	} else {
		goto L1794
	}
L1792:
	;
	if v6427-v6428 != 0 {
		v6531 = v6396
		goto L1699
	} else {
		goto L1799
	}
L1793:
	;
	goto L1792
L1794:
	;
	v6412 = v6400
	v6413 = v6403
	goto L1795
L1795:
	;
	v6416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6413)+1)))
	v6417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6412)+1)))
	if v6417 == int32(0) {
		v6427 = v6417
		v6428 = v6416
		goto L1793
	} else {
		goto L1797
	}
L1796:
	;
	v6427 = v6417
	v6428 = v6416
	goto L1793
L1797:
	;
	v6420 = int32(1)
	if v6417 == v6416 {
		v6412 = v6412 + v6420
		v6413 = v6413 + v6420
		goto L1795
	} else {
		goto L1798
	}
L1798:
	;
	goto L1796
L1799:
	;
	goto L1700
L1800:
	;
	if v6436 != int32(261) {
		goto L3
	} else {
		goto L1801
	}
L1801:
	;
	v6440 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6441 = F_strlen(m, v6440)
	mBase = m.M
	if v6441 != int32(5) {
		goto L1
	} else {
		goto L1802
	}
L1802:
	;
	v6444 = int32(_a_F_plpgsql_yyparse_82)
	v6448 = m.G0
	v6450 = v6448 - int32(32)
	v6451 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6450)+24)) = v6451
	*(*int64)(unsafe.Add(mBase, uint32(v6450)+16)) = v6451
	*(*int64)(unsafe.Add(mBase, uint32(v6450)+8)) = v6451
	*(*int64)(unsafe.Add(mBase, uint32(v6450))) = v6451
	v6459 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[51])))
	if v6459 == int32(0) {
		goto L1804
	} else {
		goto L1805
	}
L1803:
	;
	if v6527 != int32(5) {
		goto L1
	} else {
		goto L1822
	}
L1804:
	;
	v6527 = int32(0)
	goto L1803
L1805:
	;
	goto L1806
L1806:
	;
	v6463 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[52])))
	if v6463 == int32(0) {
		goto L1807
	} else {
		goto L1808
	}
L1807:
	;
	v6467 = v6440
	goto L1810
L1808:
	;
	goto L1809
L1809:
	;
	v6477 = v6444
	v6478 = v6459
	goto L1813
L1810:
	;
	v6473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6467))))
	if v6473 == v6459 {
		v6467 = v6467 + int32(1)
		goto L1810
	} else {
		goto L1812
	}
L1811:
	;
	v6527 = v6467 - v6440
	goto L1803
L1812:
	;
	goto L1811
L1813:
	;
	v6485 = v6450 + int32(base.Ui32(v6478)>>(uint(int32(3))%32))&int32(28)
	v6486 = *(*int32)(unsafe.Add(mBase, uint32(v6485)))
	v6487 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6485))) = v6486 | v6487<<(uint(v6478)%32)
	v6491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6477)+1)))
	if v6491 != 0 {
		v6477 = v6477 + v6487
		v6478 = v6491
		goto L1813
	} else {
		goto L1815
	}
L1814:
	;
	v6494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6440))))
	if v6494 == int32(0) {
		v6517 = v6440
		goto L1816
	} else {
		goto L1817
	}
L1815:
	;
	goto L1814
L1816:
	;
	v6527 = v6517 - v6440
	goto L1803
L1817:
	;
	v6498 = v6440
	v6499 = v6494
	goto L1818
L1818:
	;
	v6507 = *(*int32)(unsafe.Add(mBase, uint32(v6450+int32(base.Ui32(v6499)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v6507)>>(uint(v6499)%32))&int32(1) == int32(0) {
		v6517 = v6498
		goto L1816
	} else {
		goto L1820
	}
L1819:
	;
	v6517 = v6515
	goto L1816
L1820:
	;
	v6513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6498)+1)))
	v6515 = v6498 + int32(1)
	if v6513 != 0 {
		v6498 = v6515
		v6499 = v6513
		goto L1818
	} else {
		goto L1821
	}
L1821:
	;
	goto L1819
L1822:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6138))) = v6440
	goto L1696
L1823:
	;
	if v6531 == v6542 {
		goto L1830
	} else {
		goto L1831
	}
L1824:
	;
	v6542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6536<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[20]))))
	v6543 = base.B2i32(v6531 == v6542)
	if v6543 == int32(0) {
		goto L1826
	} else {
		goto L1827
	}
L1825:
	;
	goto L1823
L1826:
	;
	v6547 = v6536 + int32(1)
	if v6547 != int32(85) {
		v6536 = v6547
		goto L1824
	} else {
		goto L1829
	}
L1827:
	;
	goto L1828
L1828:
	;
	goto L1825
L1829:
	;
	goto L1828
L1830:
	;
	v6551 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v6552 = F_pstrdup(m, v6551)
	mBase = m.M
	v6553 = m.ExcPending
	if v6553 != 0 {
		goto L49
	} else {
		goto L1833
	}
L1831:
	;
	goto L1832
L1832:
	;
	goto L3
L1833:
	;
	v6557 = v6552
	goto L1697
L1834:
	;
	goto L1696
L1835:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v6568
	if base.B2i32(v6568 == int32(59))|base.B2i32(v6568 == int32(381)) != 0 {
		v6632 = v6568
		goto L1694
	} else {
		goto L1836
	}
L1836:
	;
	F_plpgsql_yyerror(m, v6567, int32(0), l1, int32(_a_F_plpgsql_yyparse_5))
	mBase = m.M
	v6579 = m.ExcPending
	if v6579 != 0 {
		goto L49
	} else {
		goto L1837
	}
L1837:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1838:
	;
	v6610 = int32(1)
	v6619 = F_read_sql_construct(m, int32(44), int32(59), int32(381), int32(_a_F_plpgsql_yyparse_83), int32(2), v6610, v6610, int32(0), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v6620 = m.ExcPending
	if v6620 != 0 {
		goto L49
	} else {
		goto L1840
	}
L1839:
	;
	v6632 = v6625
	goto L1694
L1840:
	;
	v6621 = *(*int32)(unsafe.Add(mBase, uint32(v6071)+24))
	v6622 = F_lappend(m, v6621, v6619)
	mBase = m.M
	v6623 = m.ExcPending
	if v6623 != 0 {
		goto L49
	} else {
		goto L1841
	}
L1841:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+24)) = v6622
	v6625 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v6625 == int32(44) {
		goto L1838
	} else {
		goto L1842
	}
L1842:
	;
	goto L1839
L1843:
	;
	goto L1693
L1844:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6071)+28)) = v7032
	goto L1692
L1845:
	;
	F_plpgsql_yyerror(m, v6683, int32(0), l1, int32(_a_F_plpgsql_yyparse_84))
	mBase = m.M
	v7053 = m.ExcPending
	if v7053 != 0 {
		goto L49
	} else {
		goto L1955
	}
L1846:
	;
	F_plpgsql_yyerror(m, v6683, int32(0), l1, int32(_a_F_plpgsql_yyparse_85))
	mBase = m.M
	v7049 = m.ExcPending
	if v7049 != 0 {
		goto L49
	} else {
		goto L1954
	}
L1847:
	;
	F_plpgsql_yyerror(m, v6683, int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v7043 = m.ExcPending
	if v7043 != 0 {
		goto L49
	} else {
		goto L1953
	}
L1848:
	;
	v6714 = F_plpgsql_yylex(m, v6681, v6683, l1)
	mBase = m.M
	v6715 = m.ExcPending
	if v6715 != 0 {
		goto L49
	} else {
		goto L1850
	}
L1849:
	;
	m.G0 = v6687 + int32(16)
	goto L1844
L1850:
	;
	if v6714 == int32(0) {
		goto L1847
	} else {
		goto L1851
	}
L1851:
	;
	v6718 = int32(0)
	v6720 = F_palloc(m, int32(8))
	mBase = m.M
	v6721 = m.ExcPending
	if v6721 != 0 {
		goto L49
	} else {
		goto L1852
	}
L1852:
	;
	switch v6714 - int32(277) {
	case 0:
		goto L1869
	default:
		goto L1846
	case 17:
		goto L1862
	case 21:
		goto L1860
	case 26:
		goto L1858
	case 30:
		goto L1866
	case 37:
		v7008 = v6718
		goto L1853
	case 49:
		goto L1864
	case 61:
		goto L1868
	case 90:
		goto L1854
	case 97:
		goto L1856
	}
L1853:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6720))) = v7008
	v7011 = F_plpgsql_yylex(m, v6681, v6683, l1)
	mBase = m.M
	v7012 = m.ExcPending
	if v7012 != 0 {
		goto L49
	} else {
		goto L1948
	}
L1854:
	;
	v7008 = int32(8)
	goto L1853
L1855:
	;
	v6975 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6975 == int32(0) {
		goto L1846
	} else {
		goto L1939
	}
L1856:
	;
	v7008 = int32(7)
	goto L1853
L1857:
	;
	v6945 = int32(_a_F_plpgsql_yyparse_86)
	v6948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6912))))
	v6951 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[53])))
	if base.B2i32(v6948 == int32(0))|base.B2i32(v6948 != v6951) != 0 {
		v6969 = v6948
		v6970 = v6951
		goto L1932
	} else {
		goto L1933
	}
L1858:
	;
	v7008 = int32(6)
	goto L1853
L1859:
	;
	v6912 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6912 == int32(0) {
		goto L1846
	} else {
		goto L1922
	}
L1860:
	;
	v7008 = int32(5)
	goto L1853
L1861:
	;
	v6882 = int32(_a_F_plpgsql_yyparse_87)
	v6885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6849))))
	v6888 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[54])))
	if base.B2i32(v6885 == int32(0))|base.B2i32(v6885 != v6888) != 0 {
		v6906 = v6885
		v6907 = v6888
		goto L1915
	} else {
		goto L1916
	}
L1862:
	;
	v7008 = int32(4)
	goto L1853
L1863:
	;
	v6849 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6849 == int32(0) {
		goto L1846
	} else {
		goto L1905
	}
L1864:
	;
	v7008 = int32(3)
	goto L1853
L1865:
	;
	v6819 = int32(_a_F_plpgsql_yyparse_88)
	v6822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6786))))
	v6825 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[55])))
	if base.B2i32(v6822 == int32(0))|base.B2i32(v6822 != v6825) != 0 {
		v6843 = v6822
		v6844 = v6825
		goto L1898
	} else {
		goto L1899
	}
L1866:
	;
	v7008 = int32(2)
	goto L1853
L1867:
	;
	v6786 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6786 == int32(0) {
		goto L1846
	} else {
		goto L1888
	}
L1868:
	;
	v7008 = int32(1)
	goto L1853
L1869:
	;
	v6724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v6724 != 0 {
		goto L1846
	} else {
		goto L1870
	}
L1870:
	;
	v6725 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v6725 == int32(0) {
		goto L1846
	} else {
		goto L1871
	}
L1871:
	;
	v6728 = int32(_a_F_plpgsql_yyparse_89)
	v6731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6725))))
	v6734 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[56])))
	if base.B2i32(v6731 == int32(0))|base.B2i32(v6731 != v6734) != 0 {
		v6752 = v6731
		v6753 = v6734
		goto L1873
	} else {
		goto L1874
	}
L1872:
	;
	if v6752-v6753 == int32(0) {
		v7008 = v6718
		goto L1853
	} else {
		goto L1879
	}
L1873:
	;
	goto L1872
L1874:
	;
	v6737 = v6725
	v6738 = v6728
	goto L1875
L1875:
	;
	v6741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6738)+1)))
	v6742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6737)+1)))
	if v6742 == int32(0) {
		v6752 = v6742
		v6753 = v6741
		goto L1873
	} else {
		goto L1877
	}
L1876:
	;
	v6752 = v6742
	v6753 = v6741
	goto L1873
L1877:
	;
	v6745 = int32(1)
	if v6742 == v6741 {
		v6737 = v6737 + v6745
		v6738 = v6738 + v6745
		goto L1875
	} else {
		goto L1878
	}
L1878:
	;
	goto L1876
L1879:
	;
	v6757 = int32(_a_F_plpgsql_yyparse_90)
	v6760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6725))))
	v6763 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[57])))
	if base.B2i32(v6760 == int32(0))|base.B2i32(v6760 != v6763) != 0 {
		v6781 = v6760
		v6782 = v6763
		goto L1881
	} else {
		goto L1882
	}
L1880:
	;
	if v6781-v6782 != 0 {
		goto L1867
	} else {
		goto L1887
	}
L1881:
	;
	goto L1880
L1882:
	;
	v6766 = v6725
	v6767 = v6757
	goto L1883
L1883:
	;
	v6770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6767)+1)))
	v6771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6766)+1)))
	if v6771 == int32(0) {
		v6781 = v6771
		v6782 = v6770
		goto L1881
	} else {
		goto L1885
	}
L1884:
	;
	v6781 = v6771
	v6782 = v6770
	goto L1881
L1885:
	;
	v6774 = int32(1)
	if v6771 == v6770 {
		v6766 = v6766 + v6774
		v6767 = v6767 + v6774
		goto L1883
	} else {
		goto L1886
	}
L1886:
	;
	goto L1884
L1887:
	;
	goto L1868
L1888:
	;
	v6789 = int32(_a_F_plpgsql_yyparse_91)
	v6792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6786))))
	v6795 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[58])))
	if base.B2i32(v6792 == int32(0))|base.B2i32(v6792 != v6795) != 0 {
		v6813 = v6792
		v6814 = v6795
		goto L1890
	} else {
		goto L1891
	}
L1889:
	;
	if v6813-v6814 != 0 {
		goto L1865
	} else {
		goto L1896
	}
L1890:
	;
	goto L1889
L1891:
	;
	v6798 = v6786
	v6799 = v6789
	goto L1892
L1892:
	;
	v6802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6799)+1)))
	v6803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6798)+1)))
	if v6803 == int32(0) {
		v6813 = v6803
		v6814 = v6802
		goto L1890
	} else {
		goto L1894
	}
L1893:
	;
	v6813 = v6803
	v6814 = v6802
	goto L1890
L1894:
	;
	v6806 = int32(1)
	if v6803 == v6802 {
		v6798 = v6798 + v6806
		v6799 = v6799 + v6806
		goto L1892
	} else {
		goto L1895
	}
L1895:
	;
	goto L1893
L1896:
	;
	goto L1866
L1897:
	;
	if v6843-v6844 != 0 {
		goto L1863
	} else {
		goto L1904
	}
L1898:
	;
	goto L1897
L1899:
	;
	v6828 = v6786
	v6829 = v6819
	goto L1900
L1900:
	;
	v6832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6829)+1)))
	v6833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6828)+1)))
	if v6833 == int32(0) {
		v6843 = v6833
		v6844 = v6832
		goto L1898
	} else {
		goto L1902
	}
L1901:
	;
	v6843 = v6833
	v6844 = v6832
	goto L1898
L1902:
	;
	v6836 = int32(1)
	if v6833 == v6832 {
		v6828 = v6828 + v6836
		v6829 = v6829 + v6836
		goto L1900
	} else {
		goto L1903
	}
L1903:
	;
	goto L1901
L1904:
	;
	goto L1864
L1905:
	;
	v6852 = int32(_a_F_plpgsql_yyparse_92)
	v6855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6849))))
	v6858 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[59])))
	if base.B2i32(v6855 == int32(0))|base.B2i32(v6855 != v6858) != 0 {
		v6876 = v6855
		v6877 = v6858
		goto L1907
	} else {
		goto L1908
	}
L1906:
	;
	if v6876-v6877 != 0 {
		goto L1861
	} else {
		goto L1913
	}
L1907:
	;
	goto L1906
L1908:
	;
	v6861 = v6849
	v6862 = v6852
	goto L1909
L1909:
	;
	v6865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6862)+1)))
	v6866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6861)+1)))
	if v6866 == int32(0) {
		v6876 = v6866
		v6877 = v6865
		goto L1907
	} else {
		goto L1911
	}
L1910:
	;
	v6876 = v6866
	v6877 = v6865
	goto L1907
L1911:
	;
	v6869 = int32(1)
	if v6866 == v6865 {
		v6861 = v6861 + v6869
		v6862 = v6862 + v6869
		goto L1909
	} else {
		goto L1912
	}
L1912:
	;
	goto L1910
L1913:
	;
	goto L1862
L1914:
	;
	if v6906-v6907 != 0 {
		goto L1859
	} else {
		goto L1921
	}
L1915:
	;
	goto L1914
L1916:
	;
	v6891 = v6849
	v6892 = v6882
	goto L1917
L1917:
	;
	v6895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6892)+1)))
	v6896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6891)+1)))
	if v6896 == int32(0) {
		v6906 = v6896
		v6907 = v6895
		goto L1915
	} else {
		goto L1919
	}
L1918:
	;
	v6906 = v6896
	v6907 = v6895
	goto L1915
L1919:
	;
	v6899 = int32(1)
	if v6896 == v6895 {
		v6891 = v6891 + v6899
		v6892 = v6892 + v6899
		goto L1917
	} else {
		goto L1920
	}
L1920:
	;
	goto L1918
L1921:
	;
	goto L1860
L1922:
	;
	v6915 = int32(_a_F_plpgsql_yyparse_93)
	v6918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6912))))
	v6921 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[60])))
	if base.B2i32(v6918 == int32(0))|base.B2i32(v6918 != v6921) != 0 {
		v6939 = v6918
		v6940 = v6921
		goto L1924
	} else {
		goto L1925
	}
L1923:
	;
	if v6939-v6940 != 0 {
		goto L1857
	} else {
		goto L1930
	}
L1924:
	;
	goto L1923
L1925:
	;
	v6924 = v6912
	v6925 = v6915
	goto L1926
L1926:
	;
	v6928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6925)+1)))
	v6929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6924)+1)))
	if v6929 == int32(0) {
		v6939 = v6929
		v6940 = v6928
		goto L1924
	} else {
		goto L1928
	}
L1927:
	;
	v6939 = v6929
	v6940 = v6928
	goto L1924
L1928:
	;
	v6932 = int32(1)
	if v6929 == v6928 {
		v6924 = v6924 + v6932
		v6925 = v6925 + v6932
		goto L1926
	} else {
		goto L1929
	}
L1929:
	;
	goto L1927
L1930:
	;
	goto L1858
L1931:
	;
	if v6969-v6970 != 0 {
		goto L1855
	} else {
		goto L1938
	}
L1932:
	;
	goto L1931
L1933:
	;
	v6954 = v6912
	v6955 = v6945
	goto L1934
L1934:
	;
	v6958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6955)+1)))
	v6959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6954)+1)))
	if v6959 == int32(0) {
		v6969 = v6959
		v6970 = v6958
		goto L1932
	} else {
		goto L1936
	}
L1935:
	;
	v6969 = v6959
	v6970 = v6958
	goto L1932
L1936:
	;
	v6962 = int32(1)
	if v6959 == v6958 {
		v6954 = v6954 + v6962
		v6955 = v6955 + v6962
		goto L1934
	} else {
		goto L1937
	}
L1937:
	;
	goto L1935
L1938:
	;
	goto L1856
L1939:
	;
	v6978 = int32(_a_F_plpgsql_yyparse_94)
	v6981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6975))))
	v6984 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[61])))
	if base.B2i32(v6981 == int32(0))|base.B2i32(v6981 != v6984) != 0 {
		v7002 = v6981
		v7003 = v6984
		goto L1941
	} else {
		goto L1942
	}
L1940:
	;
	if v7002-v7003 != 0 {
		goto L1846
	} else {
		goto L1947
	}
L1941:
	;
	goto L1940
L1942:
	;
	v6987 = v6975
	v6988 = v6978
	goto L1943
L1943:
	;
	v6991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6988)+1)))
	v6992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6987)+1)))
	if v6992 == int32(0) {
		v7002 = v6992
		v7003 = v6991
		goto L1941
	} else {
		goto L1945
	}
L1944:
	;
	v7002 = v6992
	v7003 = v6991
	goto L1941
L1945:
	;
	v6995 = int32(1)
	if v6992 == v6991 {
		v6987 = v6987 + v6995
		v6988 = v6988 + v6995
		goto L1943
	} else {
		goto L1946
	}
L1946:
	;
	goto L1944
L1947:
	;
	goto L1854
L1948:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6687)+12)) = v7011
	if base.B2i32(v7011 != int32(61))&base.B2i32(v7011 != int32(270)) != 0 {
		goto L1845
	} else {
		goto L1949
	}
L1949:
	;
	v7021 = int32(0)
	v7024 = int32(1)
	v7029 = F_read_sql_construct(m, int32(44), int32(59), v7021, int32(_a_F_plpgsql_yyparse_67), int32(2), v7024, v7024, v7021, v6687+int32(12), v6681, v6683, l1)
	mBase = m.M
	v7030 = m.ExcPending
	if v7030 != 0 {
		goto L49
	} else {
		goto L1950
	}
L1950:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6720)+4)) = v7029
	v7032 = F_lappend(m, v6693, v6720)
	mBase = m.M
	v7033 = m.ExcPending
	if v7033 != 0 {
		goto L49
	} else {
		goto L1951
	}
L1951:
	;
	v7034 = *(*int32)(unsafe.Add(mBase, uint32(v6687)+12))
	if v7034 != int32(59) {
		v6693 = v7032
		goto L1848
	} else {
		goto L1952
	}
L1952:
	;
	goto L1849
L1953:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1954:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1955:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1956:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v6071
	v9817 = v233
	goto L13
L1957:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v7175 = m.ExcPending
	if v7175 != 0 {
		goto L49
	} else {
		goto L1982
	}
L1958:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v7158 = m.ExcPending
	if v7158 != 0 {
		goto L49
	} else {
		goto L1978
	}
L1959:
	;
	v7086 = int32(0)
	v7087 = v7081
	goto L1963
L1960:
	;
	goto L1961
L1961:
	;
	goto L1956
L1962:
	;
	if v7128 < v7086 {
		goto L1958
	} else {
		goto L1977
	}
L1963:
	;
	v7107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7087))))
	if v7107 != int32(37) {
		goto L1967
	} else {
		goto L1968
	}
L1964:
	;
	v7126 = *(*int32)(unsafe.Add(mBase, uint32(v7110)+4))
	if v7086 < v7126 {
		goto L1957
	} else {
		goto L1976
	}
L1965:
	;
	goto L1964
L1966:
	;
	v7086 = v7121
	v7087 = v7122 + int32(1)
	goto L1963
L1967:
	;
	if v7107 != 0 {
		v7121 = v7086
		v7122 = v7087
		goto L1966
	} else {
		goto L1970
	}
L1968:
	;
	goto L1969
L1969:
	;
	v7116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7087)+1)))
	v7118 = base.B2i32(v7116 != int32(37))
	if v7116 != int32(37) {
		goto L1973
	} else {
		goto L1974
	}
L1970:
	;
	v7110 = *(*int32)(unsafe.Add(mBase, uint32(v6071)+24))
	if v7110 != 0 {
		goto L1965
	} else {
		goto L1971
	}
L1971:
	;
	v7111 = int32(0)
	if v7111 <= v7086 {
		v7128 = v7111
		goto L1962
	} else {
		goto L1972
	}
L1972:
	;
	goto L1957
L1973:
	;
	v7119 = v7087
	goto L1975
L1974:
	;
	v7119 = v7087 + int32(1)
	goto L1975
L1975:
	;
	v7121 = v7086 + v7118
	v7122 = v7119
	goto L1966
L1976:
	;
	v7128 = v7126
	goto L1962
L1977:
	;
	goto L1961
L1978:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7161 = m.ExcPending
	if v7161 != 0 {
		goto L49
	} else {
		goto L1979
	}
L1979:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_95), int32(0))
	mBase = m.M
	v7165 = m.ExcPending
	if v7165 != 0 {
		goto L49
	} else {
		goto L1980
	}
L1980:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(_a_F_plpgsql_yyparse_96), int32(_a_F_plpgsql_yyparse_97))
	mBase = m.M
	v7170 = m.ExcPending
	if v7170 != 0 {
		goto L49
	} else {
		goto L1981
	}
L1981:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1982:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v7178 = m.ExcPending
	if v7178 != 0 {
		goto L49
	} else {
		goto L1983
	}
L1983:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_98), int32(0))
	mBase = m.M
	v7182 = m.ExcPending
	if v7182 != 0 {
		goto L49
	} else {
		goto L1984
	}
L1984:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(_a_F_plpgsql_yyparse_99), int32(_a_F_plpgsql_yyparse_97))
	mBase = m.M
	v7187 = m.ExcPending
	if v7187 != 0 {
		goto L49
	} else {
		goto L1985
	}
L1985:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1986:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7190))) = int32(15)
	v7194 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7195 = int32(0)
	if v7194 < v7195 {
		v7240 = v7195
		goto L1988
	} else {
		goto L1989
	}
L1987:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7190)+4)) = v7240
	v7244 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v7245 = *(*int32)(unsafe.Add(mBase, uint32(v7244)+520))
	v7246 = int32(1)
	v7247 = v7245 + v7246
	*(*int32)(unsafe.Add(mBase, uint32(v7244)+520)) = v7247
	*(*int32)(unsafe.Add(mBase, uint32(v7190)+8)) = v7247
	v7252 = int32(0)
	v7261 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v7263 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v7264 = F_read_sql_construct(m, int32(44), int32(59), v7252, int32(_a_F_plpgsql_yyparse_67), int32(2), v7246, v7246, v7252, v28+int32(256), v7261, v7263, l1)
	mBase = m.M
	v7265 = m.ExcPending
	if v7265 != 0 {
		goto L49
	} else {
		goto L2000
	}
L1988:
	;
	goto L1987
L1989:
	;
	v7200 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7201 = *(*int32)(unsafe.Add(mBase, uint32(v7200)+56))
	if v7201 == int32(0) {
		v7240 = v7195
		goto L1988
	} else {
		goto L1990
	}
L1990:
	;
	v7204 = v7194 + v7201
	v7205 = *(*int32)(unsafe.Add(mBase, uint32(v7200)+184))
	if base.Ui32(v7205) <= base.Ui32(v7204) {
		goto L1992
	} else {
		goto L1993
	}
L1991:
	;
	v7215 = *(*int32)(unsafe.Add(mBase, uint32(v7200)+192))
	if base.B2i32(v7214 == int32(0))|base.B2i32(base.Ui32(v7204) <= base.Ui32(v7214)) != 0 {
		v7240 = v7215
		goto L1988
	} else {
		goto L1995
	}
L1992:
	;
	v7207 = *(*int32)(unsafe.Add(mBase, uint32(v7200)+188))
	v7214 = v7207
	goto L1991
L1993:
	;
	goto L1994
L1994:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7200)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7200)+184)) = v7201
	v7212 = F_strchr(m, v7201, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7200)+188)) = v7212
	v7214 = v7212
	goto L1991
L1995:
	;
	v7220 = v7214
	v7223 = v7215
	goto L1996
L1996:
	;
	v7225 = int32(1)
	v7226 = v7223 + v7225
	*(*int32)(unsafe.Add(mBase, uint32(v7200)+192)) = v7226
	v7229 = v7220 + v7225
	*(*int32)(unsafe.Add(mBase, uint32(v7200)+184)) = v7229
	v7232 = F_strchr(m, v7229, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7200)+188)) = v7232
	if v7232 == int32(0) {
		v7240 = v7226
		goto L1988
	} else {
		goto L1998
	}
L1997:
	;
	v7240 = v7226
	goto L1988
L1998:
	;
	if base.Ui32(v7232) < base.Ui32(v7204) {
		v7220 = v7232
		v7223 = v7226
		goto L1996
	} else {
		goto L1999
	}
L1999:
	;
	goto L1997
L2000:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7190)+12)) = v7264
	v7267 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7267 == int32(44) {
		goto L2001
	} else {
		goto L2002
	}
L2001:
	;
	v7271 = int32(0)
	v7275 = int32(1)
	v7279 = F_read_sql_construct(m, int32(59), v7271, v7271, int32(_a_F_plpgsql_yyparse_16), int32(2), v7275, v7275, v7271, v7271, v7261, v7263, l1)
	mBase = m.M
	v7280 = m.ExcPending
	if v7280 != 0 {
		goto L49
	} else {
		goto L2004
	}
L2002:
	;
	v7282 = int32(0)
	goto L2003
L2003:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7190)+16)) = v7282
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7190
	v9817 = v233
	goto L13
L2004:
	;
	v7282 = v7279
	goto L2003
L2005:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7304
	v9817 = v233
	goto L13
L2006:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7314
	v9817 = v233
	goto L13
L2007:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7324
	v9817 = v233
	goto L13
L2008:
	;
	F_plpgsql_push_back_token(m, v7331, v7328, v7330, l1)
	mBase = m.M
	v7334 = m.ExcPending
	if v7334 != 0 {
		goto L49
	} else {
		goto L2009
	}
L2009:
	;
	v7335 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	switch v7331 - int32(46) {
	case 0, 15:
		goto L26
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2010
	default:
		goto L2011
	}
L2010:
	;
	v7348 = F_make_execsql_stmt(m, int32(275), v7335, v148, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7349 = m.ExcPending
	if v7349 != 0 {
		goto L49
	} else {
		goto L2013
	}
L2011:
	;
	if base.B2i32(v7331 == int32(270))|base.B2i32(v7331 == int32(91)) != 0 {
		goto L26
	} else {
		goto L2012
	}
L2012:
	;
	goto L2010
L2013:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7348
	v9817 = v233
	goto L13
L2014:
	;
	F_plpgsql_push_back_token(m, v7355, v7352, v7354, l1)
	mBase = m.M
	v7358 = m.ExcPending
	if v7358 != 0 {
		goto L49
	} else {
		goto L2015
	}
L2015:
	;
	v7359 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	switch v7355 - int32(46) {
	case 0, 15:
		goto L25
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2016
	default:
		goto L2017
	}
L2016:
	;
	v7373 = F_make_execsql_stmt(m, int32(276), v7359, int32(0), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7374 = m.ExcPending
	if v7374 != 0 {
		goto L49
	} else {
		goto L2019
	}
L2017:
	;
	if base.B2i32(v7355 == int32(270))|base.B2i32(v7355 == int32(91)) != 0 {
		goto L25
	} else {
		goto L2018
	}
L2018:
	;
	goto L2016
L2019:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7373
	v9817 = v233
	goto L13
L2020:
	;
	v7393 = F_palloc(m, int32(28))
	mBase = m.M
	v7394 = m.ExcPending
	if v7394 != 0 {
		goto L49
	} else {
		goto L2021
	}
L2021:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7393))) = int32(17)
	v7397 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v7398 = int32(0)
	if v7397 < v7398 {
		v7443 = v7398
		goto L2023
	} else {
		goto L2024
	}
L2022:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7393)+4)) = v7443
	v7447 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v7448 = *(*int32)(unsafe.Add(mBase, uint32(v7447)+520))
	v7450 = v7448 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7447)+520)) = v7450
	*(*int64)(unsafe.Add(mBase, uint32(v7393)+20)) = int64(0)
	v7454 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7393)+16)) = uint16(v7454)
	*(*int32)(unsafe.Add(mBase, uint32(v7393)+12)) = v7390
	*(*int32)(unsafe.Add(mBase, uint32(v7393)+8)) = v7450
	v7462 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	v7467 = v7462
	goto L2035
L2023:
	;
	goto L2022
L2024:
	;
	v7403 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7404 = *(*int32)(unsafe.Add(mBase, uint32(v7403)+56))
	if v7404 == int32(0) {
		v7443 = v7398
		goto L2023
	} else {
		goto L2025
	}
L2025:
	;
	v7407 = v7397 + v7404
	v7408 = *(*int32)(unsafe.Add(mBase, uint32(v7403)+184))
	if base.Ui32(v7408) <= base.Ui32(v7407) {
		goto L2027
	} else {
		goto L2028
	}
L2026:
	;
	v7418 = *(*int32)(unsafe.Add(mBase, uint32(v7403)+192))
	if base.B2i32(v7417 == int32(0))|base.B2i32(base.Ui32(v7407) <= base.Ui32(v7417)) != 0 {
		v7443 = v7418
		goto L2023
	} else {
		goto L2030
	}
L2027:
	;
	v7410 = *(*int32)(unsafe.Add(mBase, uint32(v7403)+188))
	v7417 = v7410
	goto L2026
L2028:
	;
	goto L2029
L2029:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7403)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7403)+184)) = v7404
	v7415 = F_strchr(m, v7404, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7403)+188)) = v7415
	v7417 = v7415
	goto L2026
L2030:
	;
	v7423 = v7417
	v7426 = v7418
	goto L2031
L2031:
	;
	v7428 = int32(1)
	v7429 = v7426 + v7428
	*(*int32)(unsafe.Add(mBase, uint32(v7403)+192)) = v7429
	v7432 = v7423 + v7428
	*(*int32)(unsafe.Add(mBase, uint32(v7403)+184)) = v7432
	v7435 = F_strchr(m, v7432, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7403)+188)) = v7435
	if v7435 == int32(0) {
		v7443 = v7429
		goto L2023
	} else {
		goto L2033
	}
L2032:
	;
	v7443 = v7429
	goto L2023
L2033:
	;
	if base.Ui32(v7435) < base.Ui32(v7407) {
		v7423 = v7435
		v7426 = v7429
		goto L2031
	} else {
		goto L2034
	}
L2034:
	;
	goto L2032
L2035:
	;
	if v7467 != int32(332) {
		goto L2039
	} else {
		goto L2040
	}
L2036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7393
	v9817 = v233
	goto L13
L2037:
	;
	goto L2036
L2038:
	;
	v7508 = *(*int32)(unsafe.Add(mBase, uint32(v7393)+24))
	if v7508 != 0 {
		goto L3
	} else {
		goto L2047
	}
L2039:
	;
	if v7467 == int32(381) {
		goto L2038
	} else {
		goto L2042
	}
L2040:
	;
	goto L2041
L2041:
	;
	v7494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7393)+16)))
	if v7494 == int32(1) {
		goto L3
	} else {
		goto L2044
	}
L2042:
	;
	if v7467 == int32(59) {
		goto L2037
	} else {
		goto L2043
	}
L2043:
	;
	goto L3
L2044:
	;
	v7497 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7393)+16)) = uint8(v7497)
	v7500 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v7502 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_read_into_target(m, v7393+int32(20), v7393+int32(17), v7500, v7502, l1)
	mBase = m.M
	v7504 = m.ExcPending
	if v7504 != 0 {
		goto L49
	} else {
		goto L2045
	}
L2045:
	;
	v7505 = F_plpgsql_yylex(m, v7500, v7502, l1)
	mBase = m.M
	v7506 = m.ExcPending
	if v7506 != 0 {
		goto L49
	} else {
		goto L2046
	}
L2046:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+256)) = v7505
	v7467 = v7505
	goto L2035
L2047:
	;
	goto L2048
L2048:
	;
	v7539 = int32(1)
	v7548 = F_read_sql_construct(m, int32(44), int32(59), int32(332), int32(_a_F_plpgsql_yyparse_100), int32(2), v7539, v7539, int32(0), v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7549 = m.ExcPending
	if v7549 != 0 {
		goto L49
	} else {
		goto L2050
	}
L2049:
	;
	v7467 = v7554
	goto L2035
L2050:
	;
	v7550 = *(*int32)(unsafe.Add(mBase, uint32(v7393)+24))
	v7551 = F_lappend(m, v7550, v7548)
	mBase = m.M
	v7552 = m.ExcPending
	if v7552 != 0 {
		goto L49
	} else {
		goto L2051
	}
L2051:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7393)+24)) = v7551
	v7554 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7554 == int32(44) {
		goto L2048
	} else {
		goto L2052
	}
L2052:
	;
	goto L2049
L2053:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7559))) = int32(20)
	v7565 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v7566 = int32(0)
	if v7565 < v7566 {
		v7611 = v7566
		goto L2055
	} else {
		goto L2056
	}
L2054:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+4)) = v7611
	v7615 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v7616 = *(*int32)(unsafe.Add(mBase, uint32(v7615)+520))
	v7618 = v7616 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7615)+520)) = v7618
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+8)) = v7618
	v7621 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v7622 = *(*int32)(unsafe.Add(mBase, uint32(v7621)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+16)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+12)) = v7622
	v7626 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v7627 = *(*int32)(unsafe.Add(mBase, uint32(v7626)+28))
	if v7627 == int32(0) {
		goto L2067
	} else {
		goto L2068
	}
L2055:
	;
	goto L2054
L2056:
	;
	v7571 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7572 = *(*int32)(unsafe.Add(mBase, uint32(v7571)+56))
	if v7572 == int32(0) {
		v7611 = v7566
		goto L2055
	} else {
		goto L2057
	}
L2057:
	;
	v7575 = v7565 + v7572
	v7576 = *(*int32)(unsafe.Add(mBase, uint32(v7571)+184))
	if base.Ui32(v7576) <= base.Ui32(v7575) {
		goto L2059
	} else {
		goto L2060
	}
L2058:
	;
	v7586 = *(*int32)(unsafe.Add(mBase, uint32(v7571)+192))
	if base.B2i32(v7585 == int32(0))|base.B2i32(base.Ui32(v7575) <= base.Ui32(v7585)) != 0 {
		v7611 = v7586
		goto L2055
	} else {
		goto L2062
	}
L2059:
	;
	v7578 = *(*int32)(unsafe.Add(mBase, uint32(v7571)+188))
	v7585 = v7578
	goto L2058
L2060:
	;
	goto L2061
L2061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7571)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7571)+184)) = v7572
	v7583 = F_strchr(m, v7572, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7571)+188)) = v7583
	v7585 = v7583
	goto L2058
L2062:
	;
	v7591 = v7585
	v7594 = v7586
	goto L2063
L2063:
	;
	v7596 = int32(1)
	v7597 = v7594 + v7596
	*(*int32)(unsafe.Add(mBase, uint32(v7571)+192)) = v7597
	v7600 = v7591 + v7596
	*(*int32)(unsafe.Add(mBase, uint32(v7571)+184)) = v7600
	v7603 = F_strchr(m, v7600, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7571)+188)) = v7603
	if v7603 == int32(0) {
		v7611 = v7597
		goto L2055
	} else {
		goto L2065
	}
L2064:
	;
	v7611 = v7597
	goto L2055
L2065:
	;
	if base.Ui32(v7603) < base.Ui32(v7575) {
		v7591 = v7603
		v7594 = v7597
		goto L2063
	} else {
		goto L2066
	}
L2066:
	;
	goto L2064
L2067:
	;
	v7630 = int32(2)
	v7635 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7636 = m.ExcPending
	if v7636 != 0 {
		goto L49
	} else {
		goto L2072
	}
L2068:
	;
	goto L2069
L2069:
	;
	v7903 = F_read_cursor_args(m, v7626, int32(59), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7904 = m.ExcPending
	if v7904 != 0 {
		goto L49
	} else {
		goto L2136
	}
L2070:
	;
	if v7745 != int32(321) {
		goto L19
	} else {
		goto L2108
	}
L2071:
	;
	v7736 = *(*int32)(unsafe.Add(mBase, uint32(v7559)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+16)) = v7736 | v7735
	v7743 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7744 = m.ExcPending
	if v7744 != 0 {
		goto L49
	} else {
		goto L2107
	}
L2072:
	;
	if v7635 == int32(369) {
		v7735 = v7630
		goto L2071
	} else {
		goto L2073
	}
L2073:
	;
	if v7635 != int32(342) {
		goto L2076
	} else {
		goto L2077
	}
L2074:
	;
	v7707 = int32(_a_F_plpgsql_yyparse_101)
	v7710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7706))))
	v7713 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[62])))
	if base.B2i32(v7710 == int32(0))|base.B2i32(v7710 != v7713) != 0 {
		v7731 = v7710
		v7732 = v7713
		goto L2100
	} else {
		goto L2101
	}
L2075:
	;
	v7703 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v7703 == int32(0) {
		goto L19
	} else {
		goto L2098
	}
L2076:
	;
	if v7635 != int32(277) {
		v7745 = v7635
		goto L2070
	} else {
		goto L2079
	}
L2077:
	;
	goto L2078
L2078:
	;
	v7657 = int32(4)
	v7662 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7663 = m.ExcPending
	if v7663 != 0 {
		goto L49
	} else {
		goto L2085
	}
L2079:
	;
	v7643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v7643&int32(1) != 0 {
		goto L19
	} else {
		goto L2080
	}
L2080:
	;
	v7646 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v7646 == int32(0) {
		goto L19
	} else {
		goto L2081
	}
L2081:
	;
	v7649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7646))))
	if v7649 != int32(110) {
		v7706 = v7646
		goto L2074
	} else {
		goto L2082
	}
L2082:
	;
	v7652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7646)+1)))
	if v7652 != int32(111) {
		goto L2075
	} else {
		goto L2083
	}
L2083:
	;
	v7655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7646)+2)))
	if v7655 != 0 {
		goto L2075
	} else {
		goto L2084
	}
L2084:
	;
	goto L2078
L2085:
	;
	if v7662 == int32(369) {
		v7735 = v7657
		goto L2071
	} else {
		goto L2086
	}
L2086:
	;
	if v7662 != int32(277) {
		v7745 = v7662
		goto L2070
	} else {
		goto L2087
	}
L2087:
	;
	v7668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v7668&int32(1) != 0 {
		goto L19
	} else {
		goto L2088
	}
L2088:
	;
	v7671 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v7671 == int32(0) {
		goto L19
	} else {
		goto L2089
	}
L2089:
	;
	v7674 = int32(_a_F_plpgsql_yyparse_101)
	v7677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7671))))
	v7680 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[62])))
	if base.B2i32(v7677 == int32(0))|base.B2i32(v7677 != v7680) != 0 {
		v7698 = v7677
		v7699 = v7680
		goto L2091
	} else {
		goto L2092
	}
L2090:
	;
	if v7698-v7699 == int32(0) {
		v7735 = v7657
		goto L2071
	} else {
		goto L2097
	}
L2091:
	;
	goto L2090
L2092:
	;
	v7683 = v7671
	v7684 = v7674
	goto L2093
L2093:
	;
	v7687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7684)+1)))
	v7688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7683)+1)))
	if v7688 == int32(0) {
		v7698 = v7688
		v7699 = v7687
		goto L2091
	} else {
		goto L2095
	}
L2094:
	;
	v7698 = v7688
	v7699 = v7687
	goto L2091
L2095:
	;
	v7691 = int32(1)
	if v7688 == v7687 {
		v7683 = v7683 + v7691
		v7684 = v7684 + v7691
		goto L2093
	} else {
		goto L2096
	}
L2096:
	;
	goto L2094
L2097:
	;
	goto L19
L2098:
	;
	v7706 = v7703
	goto L2074
L2099:
	;
	if v7731-v7732 != 0 {
		goto L19
	} else {
		goto L2106
	}
L2100:
	;
	goto L2099
L2101:
	;
	v7716 = v7706
	v7717 = v7707
	goto L2102
L2102:
	;
	v7720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7717)+1)))
	v7721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7716)+1)))
	if v7721 == int32(0) {
		v7731 = v7721
		v7732 = v7720
		goto L2100
	} else {
		goto L2104
	}
L2103:
	;
	v7731 = v7721
	v7732 = v7720
	goto L2100
L2104:
	;
	v7724 = int32(1)
	if v7721 == v7720 {
		v7716 = v7716 + v7724
		v7717 = v7717 + v7724
		goto L2102
	} else {
		goto L2105
	}
L2105:
	;
	goto L2103
L2106:
	;
	v7735 = v7630
	goto L2071
L2107:
	;
	v7745 = v7743
	goto L2070
L2108:
	;
	v7753 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7754 = m.ExcPending
	if v7754 != 0 {
		goto L49
	} else {
		goto L2110
	}
L2109:
	;
	v7889 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v7891 = v28 + int32(_a_F_plpgsql_yyparse_2)
	F_plpgsql_push_back_token(m, v7753, v7889, v7891, l1)
	mBase = m.M
	v7893 = m.ExcPending
	if v7893 != 0 {
		goto L49
	} else {
		goto L2134
	}
L2110:
	;
	if v7753 != int32(317) {
		goto L2111
	} else {
		goto L2112
	}
L2111:
	;
	if v7753 != int32(277) {
		goto L2109
	} else {
		goto L2114
	}
L2112:
	;
	goto L2113
L2113:
	;
	v7795 = int32(0)
	v7798 = int32(1)
	v7807 = F_read_sql_construct(m, int32(381), int32(59), v7795, int32(_a_F_plpgsql_yyparse_102), int32(2), v7798, v7798, v7795, v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7808 = m.ExcPending
	if v7808 != 0 {
		goto L49
	} else {
		goto L2125
	}
L2114:
	;
	v7759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v7759&int32(1) != 0 {
		goto L2109
	} else {
		goto L2115
	}
L2115:
	;
	v7762 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v7762 == int32(0) {
		goto L2109
	} else {
		goto L2116
	}
L2116:
	;
	v7765 = int32(_a_F_plpgsql_yyparse_58)
	v7768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7762))))
	v7771 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[42])))
	if base.B2i32(v7768 == int32(0))|base.B2i32(v7768 != v7771) != 0 {
		v7789 = v7768
		v7790 = v7771
		goto L2118
	} else {
		goto L2119
	}
L2117:
	;
	if v7789-v7790 != 0 {
		goto L2109
	} else {
		goto L2124
	}
L2118:
	;
	goto L2117
L2119:
	;
	v7774 = v7762
	v7775 = v7765
	goto L2120
L2120:
	;
	v7778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7775)+1)))
	v7779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7774)+1)))
	if v7779 == int32(0) {
		v7789 = v7779
		v7790 = v7778
		goto L2118
	} else {
		goto L2122
	}
L2121:
	;
	v7789 = v7779
	v7790 = v7778
	goto L2118
L2122:
	;
	v7782 = int32(1)
	if v7779 == v7778 {
		v7774 = v7774 + v7782
		v7775 = v7775 + v7782
		goto L2120
	} else {
		goto L2123
	}
L2123:
	;
	goto L2121
L2124:
	;
	goto L2113
L2125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+28)) = v7807
	v7810 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7810 == int32(381) {
		goto L2126
	} else {
		goto L2127
	}
L2126:
	;
	goto L2129
L2127:
	;
	goto L2128
L2128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7559
	v9817 = v233
	goto L13
L2129:
	;
	v7840 = int32(0)
	v7843 = int32(1)
	v7852 = F_read_sql_construct(m, int32(44), int32(59), v7840, int32(_a_F_plpgsql_yyparse_67), int32(2), v7843, v7843, v7840, v28+int32(256), v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v7853 = m.ExcPending
	if v7853 != 0 {
		goto L49
	} else {
		goto L2131
	}
L2130:
	;
	goto L2128
L2131:
	;
	v7854 = *(*int32)(unsafe.Add(mBase, uint32(v7559)+32))
	v7855 = F_lappend(m, v7854, v7852)
	mBase = m.M
	v7856 = m.ExcPending
	if v7856 != 0 {
		goto L49
	} else {
		goto L2132
	}
L2132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+32)) = v7855
	v7858 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v7858 == int32(44) {
		goto L2129
	} else {
		goto L2133
	}
L2133:
	;
	goto L2130
L2134:
	;
	v7894 = F_read_sql_stmt(m, v7889, v7891, l1)
	mBase = m.M
	v7895 = m.ExcPending
	if v7895 != 0 {
		goto L49
	} else {
		goto L2135
	}
L2135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+24)) = v7894
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7559
	v9817 = v233
	goto L13
L2136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7559)+20)) = v7903
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7559
	v9817 = v233
	goto L13
L2137:
	;
	v7919 = F_plpgsql_yylex(m, v7914, v7916, l1)
	mBase = m.M
	v7920 = m.ExcPending
	if v7920 != 0 {
		goto L49
	} else {
		goto L2138
	}
L2138:
	;
	if v7919 != int32(59) {
		goto L3
	} else {
		goto L2139
	}
L2139:
	;
	v7923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7909)+33)))
	if v7923 == int32(1) {
		goto L24
	} else {
		goto L2140
	}
L2140:
	;
	v7928 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v7929 = int32(0)
	if v7928 < v7929 {
		v7974 = v7929
		goto L2142
	} else {
		goto L2143
	}
L2141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7909)+4)) = v7974
	v7977 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	*(*int32)(unsafe.Add(mBase, uint32(v7909)+12)) = v7977
	v7981 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v7982 = *(*int32)(unsafe.Add(mBase, uint32(v7981)+4))
	v7983 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7909)+32)) = uint8(v7983)
	*(*int32)(unsafe.Add(mBase, uint32(v7909)+16)) = v7982
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7909
	v9817 = v233
	goto L13
L2142:
	;
	goto L2141
L2143:
	;
	v7934 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7935 = *(*int32)(unsafe.Add(mBase, uint32(v7934)+56))
	if v7935 == int32(0) {
		v7974 = v7929
		goto L2142
	} else {
		goto L2144
	}
L2144:
	;
	v7938 = v7928 + v7935
	v7939 = *(*int32)(unsafe.Add(mBase, uint32(v7934)+184))
	if base.Ui32(v7939) <= base.Ui32(v7938) {
		goto L2146
	} else {
		goto L2147
	}
L2145:
	;
	v7949 = *(*int32)(unsafe.Add(mBase, uint32(v7934)+192))
	if base.B2i32(v7948 == int32(0))|base.B2i32(base.Ui32(v7938) <= base.Ui32(v7948)) != 0 {
		v7974 = v7949
		goto L2142
	} else {
		goto L2149
	}
L2146:
	;
	v7941 = *(*int32)(unsafe.Add(mBase, uint32(v7934)+188))
	v7948 = v7941
	goto L2145
L2147:
	;
	goto L2148
L2148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7934)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7934)+184)) = v7935
	v7946 = F_strchr(m, v7935, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7934)+188)) = v7946
	v7948 = v7946
	goto L2145
L2149:
	;
	v7954 = v7948
	v7957 = v7949
	goto L2150
L2150:
	;
	v7959 = int32(1)
	v7960 = v7957 + v7959
	*(*int32)(unsafe.Add(mBase, uint32(v7934)+192)) = v7960
	v7963 = v7954 + v7959
	*(*int32)(unsafe.Add(mBase, uint32(v7934)+184)) = v7963
	v7966 = F_strchr(m, v7963, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7934)+188)) = v7966
	if v7966 == int32(0) {
		v7974 = v7960
		goto L2142
	} else {
		goto L2152
	}
L2151:
	;
	v7974 = v7960
	goto L2142
L2152:
	;
	if base.Ui32(v7966) < base.Ui32(v7938) {
		v7954 = v7966
		v7957 = v7960
		goto L2150
	} else {
		goto L2153
	}
L2153:
	;
	goto L2151
L2154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7989)+4)) = v8038
	v8043 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8044 = *(*int32)(unsafe.Add(mBase, uint32(v8043)+4))
	v8045 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7989)+32)) = uint8(v8045)
	*(*int32)(unsafe.Add(mBase, uint32(v7989)+16)) = v8044
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v7989
	v9817 = v233
	goto L13
L2155:
	;
	goto L2154
L2156:
	;
	v7998 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7999 = *(*int32)(unsafe.Add(mBase, uint32(v7998)+56))
	if v7999 == int32(0) {
		v8038 = v7993
		goto L2155
	} else {
		goto L2157
	}
L2157:
	;
	v8002 = v7992 + v7999
	v8003 = *(*int32)(unsafe.Add(mBase, uint32(v7998)+184))
	if base.Ui32(v8003) <= base.Ui32(v8002) {
		goto L2159
	} else {
		goto L2160
	}
L2158:
	;
	v8013 = *(*int32)(unsafe.Add(mBase, uint32(v7998)+192))
	if base.B2i32(v8012 == int32(0))|base.B2i32(base.Ui32(v8002) <= base.Ui32(v8012)) != 0 {
		v8038 = v8013
		goto L2155
	} else {
		goto L2162
	}
L2159:
	;
	v8005 = *(*int32)(unsafe.Add(mBase, uint32(v7998)+188))
	v8012 = v8005
	goto L2158
L2160:
	;
	goto L2161
L2161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7998)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7998)+184)) = v7999
	v8010 = F_strchr(m, v7999, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7998)+188)) = v8010
	v8012 = v8010
	goto L2158
L2162:
	;
	v8018 = v8012
	v8021 = v8013
	goto L2163
L2163:
	;
	v8023 = int32(1)
	v8024 = v8021 + v8023
	*(*int32)(unsafe.Add(mBase, uint32(v7998)+192)) = v8024
	v8027 = v8018 + v8023
	*(*int32)(unsafe.Add(mBase, uint32(v7998)+184)) = v8027
	v8030 = F_strchr(m, v8027, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7998)+188)) = v8030
	if v8030 == int32(0) {
		v8038 = v8024
		goto L2155
	} else {
		goto L2165
	}
L2164:
	;
	v8038 = v8024
	goto L2155
L2165:
	;
	if base.Ui32(v8030) < base.Ui32(v8002) {
		v8018 = v8030
		v8021 = v8024
		goto L2163
	} else {
		goto L2166
	}
L2166:
	;
	goto L2164
L2167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8060))) = int32(21)
	v8065 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v8066 = *(*int32)(unsafe.Add(mBase, uint32(v8065)+520))
	v8068 = v8066 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8065)+520)) = v8068
	v8070 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8060)+33)) = uint8(v8070)
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+28)) = v8070
	*(*int64)(unsafe.Add(mBase, uint32(v8060)+20)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+8)) = v8068
	v8077 = F_plpgsql_yylex(m, v8050, v8052, l1)
	mBase = m.M
	v8078 = m.ExcPending
	if v8078 != 0 {
		goto L49
	} else {
		goto L2180
	}
L2168:
	;
	v8411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8055)+15)))
	if v8411 != int32(1) {
		goto L2287
	} else {
		goto L2288
	}
L2169:
	;
	v8315 = int32(1)
	if v8077 == int32(282) {
		v8329 = v8315
		goto L2245
	} else {
		goto L2246
	}
L2170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+20)) = int32(3)
	v8299 = int32(0)
	v8302 = int32(1)
	v8306 = F_read_sql_construct(m, int32(324), int32(329), v8299, int32(_a_F_plpgsql_yyparse_103), int32(2), v8302, v8302, v8299, v8299, v8050, v8052, l1)
	mBase = m.M
	v8307 = m.ExcPending
	if v8307 != 0 {
		goto L49
	} else {
		goto L2243
	}
L2171:
	;
	v8264 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8264 == int32(0) {
		goto L2169
	} else {
		goto L2234
	}
L2172:
	;
	v8248 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+20)) = v8248
	v8252 = int32(0)
	v8255 = int32(1)
	v8259 = F_read_sql_construct(m, int32(324), int32(329), v8252, int32(_a_F_plpgsql_yyparse_103), v8248, v8255, v8255, v8252, v8252, v8050, v8052, l1)
	mBase = m.M
	v8260 = m.ExcPending
	if v8260 != 0 {
		goto L49
	} else {
		goto L2233
	}
L2173:
	;
	v8217 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8217 == int32(0) {
		goto L2171
	} else {
		goto L2224
	}
L2174:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8060)+20)) = int64(-4294967294)
	goto L2168
L2175:
	;
	v8184 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8184 == int32(0) {
		goto L2173
	} else {
		goto L2215
	}
L2176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+20)) = int32(2)
	goto L2168
L2177:
	;
	v8151 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8151 == int32(0) {
		goto L2175
	} else {
		goto L2206
	}
L2178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+20)) = int32(1)
	goto L2168
L2179:
	;
	switch v8077 - int32(277) {
	case 0:
		goto L2181
	case 1, 2:
		goto L2169
	case 3:
		goto L2172
	default:
		goto L2182
	}
L2180:
	;
	switch v8077 - int32(320) {
	case 0:
		goto L2176
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 38, 39:
		goto L2169
	case 14:
		goto L2174
	case 21:
		goto L2168
	case 37:
		goto L2178
	case 40:
		goto L2170
	default:
		goto L2179
	}
L2181:
	;
	v8087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8087 != 0 {
		goto L2169
	} else {
		goto L2185
	}
L2182:
	;
	if v8077 != 0 {
		goto L2169
	} else {
		goto L2183
	}
L2183:
	;
	F_plpgsql_yyerror(m, v8052, int32(0), l1, int32(_a_F_plpgsql_yyparse_6))
	mBase = m.M
	v8086 = m.ExcPending
	if v8086 != 0 {
		goto L49
	} else {
		goto L2184
	}
L2184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2185:
	;
	v8088 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8088 != 0 {
		goto L2186
	} else {
		goto L2187
	}
L2186:
	;
	v8089 = int32(_a_F_plpgsql_yyparse_62)
	v8092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8088))))
	v8095 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[44])))
	if base.B2i32(v8092 == int32(0))|base.B2i32(v8092 != v8095) != 0 {
		v8113 = v8092
		v8114 = v8095
		goto L2190
	} else {
		goto L2191
	}
L2187:
	;
	goto L2188
L2188:
	;
	v8118 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8118 == int32(0) {
		goto L2177
	} else {
		goto L2197
	}
L2189:
	;
	if v8113-v8114 == int32(0) {
		goto L2168
	} else {
		goto L2196
	}
L2190:
	;
	goto L2189
L2191:
	;
	v8098 = v8088
	v8099 = v8089
	goto L2192
L2192:
	;
	v8102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8099)+1)))
	v8103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8098)+1)))
	if v8103 == int32(0) {
		v8113 = v8103
		v8114 = v8102
		goto L2190
	} else {
		goto L2194
	}
L2193:
	;
	v8113 = v8103
	v8114 = v8102
	goto L2190
L2194:
	;
	v8106 = int32(1)
	if v8103 == v8102 {
		v8098 = v8098 + v8106
		v8099 = v8099 + v8106
		goto L2192
	} else {
		goto L2195
	}
L2195:
	;
	goto L2193
L2196:
	;
	goto L2188
L2197:
	;
	v8121 = int32(_a_F_plpgsql_yyparse_104)
	v8124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8118))))
	v8127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[63])))
	if base.B2i32(v8124 == int32(0))|base.B2i32(v8124 != v8127) != 0 {
		v8145 = v8124
		v8146 = v8127
		goto L2199
	} else {
		goto L2200
	}
L2198:
	;
	if v8145-v8146 != 0 {
		goto L2177
	} else {
		goto L2205
	}
L2199:
	;
	goto L2198
L2200:
	;
	v8130 = v8118
	v8131 = v8121
	goto L2201
L2201:
	;
	v8134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8131)+1)))
	v8135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8130)+1)))
	if v8135 == int32(0) {
		v8145 = v8135
		v8146 = v8134
		goto L2199
	} else {
		goto L2203
	}
L2202:
	;
	v8145 = v8135
	v8146 = v8134
	goto L2199
L2203:
	;
	v8138 = int32(1)
	if v8135 == v8134 {
		v8130 = v8130 + v8138
		v8131 = v8131 + v8138
		goto L2201
	} else {
		goto L2204
	}
L2204:
	;
	goto L2202
L2205:
	;
	goto L2178
L2206:
	;
	v8154 = int32(_a_F_plpgsql_yyparse_105)
	v8157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8151))))
	v8160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[64])))
	if base.B2i32(v8157 == int32(0))|base.B2i32(v8157 != v8160) != 0 {
		v8178 = v8157
		v8179 = v8160
		goto L2208
	} else {
		goto L2209
	}
L2207:
	;
	if v8178-v8179 != 0 {
		goto L2175
	} else {
		goto L2214
	}
L2208:
	;
	goto L2207
L2209:
	;
	v8163 = v8151
	v8164 = v8154
	goto L2210
L2210:
	;
	v8167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8164)+1)))
	v8168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8163)+1)))
	if v8168 == int32(0) {
		v8178 = v8168
		v8179 = v8167
		goto L2208
	} else {
		goto L2212
	}
L2211:
	;
	v8178 = v8168
	v8179 = v8167
	goto L2208
L2212:
	;
	v8171 = int32(1)
	if v8168 == v8167 {
		v8163 = v8163 + v8171
		v8164 = v8164 + v8171
		goto L2210
	} else {
		goto L2213
	}
L2213:
	;
	goto L2211
L2214:
	;
	goto L2176
L2215:
	;
	v8187 = int32(_a_F_plpgsql_yyparse_106)
	v8190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8184))))
	v8193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[65])))
	if base.B2i32(v8190 == int32(0))|base.B2i32(v8190 != v8193) != 0 {
		v8211 = v8190
		v8212 = v8193
		goto L2217
	} else {
		goto L2218
	}
L2216:
	;
	if v8211-v8212 != 0 {
		goto L2173
	} else {
		goto L2223
	}
L2217:
	;
	goto L2216
L2218:
	;
	v8196 = v8184
	v8197 = v8187
	goto L2219
L2219:
	;
	v8200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8197)+1)))
	v8201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8196)+1)))
	if v8201 == int32(0) {
		v8211 = v8201
		v8212 = v8200
		goto L2217
	} else {
		goto L2221
	}
L2220:
	;
	v8211 = v8201
	v8212 = v8200
	goto L2217
L2221:
	;
	v8204 = int32(1)
	if v8201 == v8200 {
		v8196 = v8196 + v8204
		v8197 = v8197 + v8204
		goto L2219
	} else {
		goto L2222
	}
L2222:
	;
	goto L2220
L2223:
	;
	goto L2174
L2224:
	;
	v8220 = int32(_a_F_plpgsql_yyparse_107)
	v8223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8217))))
	v8226 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[66])))
	if base.B2i32(v8223 == int32(0))|base.B2i32(v8223 != v8226) != 0 {
		v8244 = v8223
		v8245 = v8226
		goto L2226
	} else {
		goto L2227
	}
L2225:
	;
	if v8244-v8245 != 0 {
		goto L2171
	} else {
		goto L2232
	}
L2226:
	;
	goto L2225
L2227:
	;
	v8229 = v8217
	v8230 = v8220
	goto L2228
L2228:
	;
	v8233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8230)+1)))
	v8234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8229)+1)))
	if v8234 == int32(0) {
		v8244 = v8234
		v8245 = v8233
		goto L2226
	} else {
		goto L2230
	}
L2229:
	;
	v8244 = v8234
	v8245 = v8233
	goto L2226
L2230:
	;
	v8237 = int32(1)
	if v8234 == v8233 {
		v8229 = v8229 + v8237
		v8230 = v8230 + v8237
		goto L2228
	} else {
		goto L2231
	}
L2231:
	;
	goto L2229
L2232:
	;
	goto L2172
L2233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+28)) = v8259
	v8262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8055)+15)) = uint8(v8262)
	goto L2168
L2234:
	;
	v8267 = int32(_a_F_plpgsql_yyparse_108)
	v8270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8264))))
	v8273 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[67])))
	if base.B2i32(v8270 == int32(0))|base.B2i32(v8270 != v8273) != 0 {
		v8291 = v8270
		v8292 = v8273
		goto L2236
	} else {
		goto L2237
	}
L2235:
	;
	if v8291-v8292 != 0 {
		goto L2169
	} else {
		goto L2242
	}
L2236:
	;
	goto L2235
L2237:
	;
	v8276 = v8264
	v8277 = v8267
	goto L2238
L2238:
	;
	v8280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8277)+1)))
	v8281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8276)+1)))
	if v8281 == int32(0) {
		v8291 = v8281
		v8292 = v8280
		goto L2236
	} else {
		goto L2240
	}
L2239:
	;
	v8291 = v8281
	v8292 = v8280
	goto L2236
L2240:
	;
	v8284 = int32(1)
	if v8281 == v8280 {
		v8276 = v8276 + v8284
		v8277 = v8277 + v8284
		goto L2238
	} else {
		goto L2241
	}
L2241:
	;
	goto L2239
L2242:
	;
	goto L2170
L2243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+28)) = v8306
	v8309 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8055)+15)) = uint8(v8309)
	goto L2168
L2244:
	;
	if v8329 != 0 {
		goto L2252
	} else {
		goto L2253
	}
L2245:
	;
	goto L2244
L2246:
	;
	if v8077 != int32(277) {
		goto L2247
	} else {
		goto L2248
	}
L2247:
	;
	v8329 = int32(0)
	goto L2245
L2248:
	;
	v8319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8319 != 0 {
		goto L2247
	} else {
		goto L2249
	}
L2249:
	;
	v8320 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8320 == int32(0) {
		goto L2247
	} else {
		goto L2250
	}
L2250:
	;
	v8323 = F_strcmp(m, v8320, int32(_a_F_plpgsql_yyparse_109))
	mBase = m.M
	if v8323 == int32(0) {
		v8329 = v8315
		goto L2245
	} else {
		goto L2251
	}
L2251:
	;
	goto L2247
L2252:
	;
	v8330 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8060)+33)) = uint8(v8330)
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+24)) = int32(2147483647)
	goto L2168
L2253:
	;
	goto L2254
L2254:
	;
	v8337 = int32(1)
	if v8077 == int32(323) {
		v8351 = v8337
		goto L2256
	} else {
		goto L2257
	}
L2255:
	;
	if v8351 != 0 {
		goto L2263
	} else {
		goto L2264
	}
L2256:
	;
	goto L2255
L2257:
	;
	if v8077 != int32(277) {
		goto L2258
	} else {
		goto L2259
	}
L2258:
	;
	v8351 = int32(0)
	goto L2256
L2259:
	;
	v8341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8341 != 0 {
		goto L2258
	} else {
		goto L2260
	}
L2260:
	;
	v8342 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8342 == int32(0) {
		goto L2258
	} else {
		goto L2261
	}
L2261:
	;
	v8345 = F_strcmp(m, v8342, int32(_a_F_plpgsql_yyparse_110))
	mBase = m.M
	if v8345 == int32(0) {
		v8351 = v8337
		goto L2256
	} else {
		goto L2262
	}
L2262:
	;
	goto L2258
L2263:
	;
	F_complete_direction(m, v8060, v8055+int32(15), v8050, v8052, l1)
	mBase = m.M
	v8355 = m.ExcPending
	if v8355 != 0 {
		goto L49
	} else {
		goto L2266
	}
L2264:
	;
	goto L2265
L2265:
	;
	v8359 = int32(1)
	if v8077 == int32(286) {
		v8373 = v8359
		goto L2268
	} else {
		goto L2269
	}
L2266:
	;
	goto L2168
L2267:
	;
	if v8373 != 0 {
		goto L2275
	} else {
		goto L2276
	}
L2268:
	;
	goto L2267
L2269:
	;
	if v8077 != int32(277) {
		goto L2270
	} else {
		goto L2271
	}
L2270:
	;
	v8373 = int32(0)
	goto L2268
L2271:
	;
	v8363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0]))))
	if v8363 != 0 {
		goto L2270
	} else {
		goto L2272
	}
L2272:
	;
	v8364 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[5])))
	if v8364 == int32(0) {
		goto L2270
	} else {
		goto L2273
	}
L2273:
	;
	v8367 = F_strcmp(m, v8364, int32(_a_F_plpgsql_yyparse_111))
	mBase = m.M
	if v8367 == int32(0) {
		v8373 = v8359
		goto L2268
	} else {
		goto L2274
	}
L2274:
	;
	goto L2270
L2275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+20)) = int32(1)
	F_complete_direction(m, v8060, v8055+int32(15), v8050, v8052, l1)
	mBase = m.M
	v8379 = m.ExcPending
	if v8379 != 0 {
		goto L49
	} else {
		goto L2278
	}
L2276:
	;
	goto L2277
L2277:
	;
	switch v8077 - int32(324) {
	case 0, 5:
		goto L2282
	case 1, 2, 3, 4:
		goto L2280
	default:
		goto L2281
	}
L2278:
	;
	goto L2168
L2279:
	;
	F_plpgsql_push_back_token(m, int32(277), v8050, v8052, l1)
	mBase = m.M
	v8406 = m.ExcPending
	if v8406 != 0 {
		goto L49
	} else {
		goto L2286
	}
L2280:
	;
	F_plpgsql_push_back_token(m, v8077, v8050, v8052, l1)
	mBase = m.M
	v8387 = m.ExcPending
	if v8387 != 0 {
		goto L49
	} else {
		goto L2284
	}
L2281:
	;
	if v8077 == int32(277) {
		goto L2279
	} else {
		goto L2283
	}
L2282:
	;
	v8382 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8055)+15)) = uint8(v8382)
	goto L2168
L2283:
	;
	goto L2280
L2284:
	;
	v8390 = int32(0)
	v8393 = int32(1)
	v8397 = F_read_sql_construct(m, int32(324), int32(329), v8390, int32(_a_F_plpgsql_yyparse_103), int32(2), v8393, v8393, v8390, v8390, v8050, v8052, l1)
	mBase = m.M
	v8398 = m.ExcPending
	if v8398 != 0 {
		goto L49
	} else {
		goto L2285
	}
L2285:
	;
	v8399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8060)+33)) = uint8(v8399)
	*(*int32)(unsafe.Add(mBase, uint32(v8060)+28)) = v8397
	v8402 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8055)+15)) = uint8(v8402)
	goto L2168
L2286:
	;
	v8407 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8055)+15)) = uint8(v8407)
	goto L2168
L2287:
	;
	m.G0 = v8055 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8060
	v9817 = v233
	goto L13
L2288:
	;
	v8414 = F_plpgsql_yylex(m, v8050, v8052, l1)
	mBase = m.M
	v8415 = m.ExcPending
	if v8415 != 0 {
		goto L49
	} else {
		goto L2290
	}
L2289:
	;
	F_plpgsql_yyerror(m, v8052, int32(0), l1, int32(_a_F_plpgsql_yyparse_112))
	mBase = m.M
	v8421 = m.ExcPending
	if v8421 != 0 {
		goto L49
	} else {
		goto L2291
	}
L2290:
	;
	switch v8414 - int32(324) {
	case 0, 5:
		goto L2287
	default:
		goto L2289
	}
L2291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8427))) = int32(22)
	v8433 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v8434 = int32(0)
	if v8433 < v8434 {
		v8479 = v8434
		goto L2294
	} else {
		goto L2295
	}
L2293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8427)+4)) = v8479
	v8483 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v8484 = *(*int32)(unsafe.Add(mBase, uint32(v8483)+520))
	v8486 = v8484 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8483)+520)) = v8486
	*(*int32)(unsafe.Add(mBase, uint32(v8427)+8)) = v8486
	v8491 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v8492 = *(*int32)(unsafe.Add(mBase, uint32(v8491)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8427)+12)) = v8492
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8427
	v9817 = v233
	goto L13
L2294:
	;
	goto L2293
L2295:
	;
	v8439 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8440 = *(*int32)(unsafe.Add(mBase, uint32(v8439)+56))
	if v8440 == int32(0) {
		v8479 = v8434
		goto L2294
	} else {
		goto L2296
	}
L2296:
	;
	v8443 = v8433 + v8440
	v8444 = *(*int32)(unsafe.Add(mBase, uint32(v8439)+184))
	if base.Ui32(v8444) <= base.Ui32(v8443) {
		goto L2298
	} else {
		goto L2299
	}
L2297:
	;
	v8454 = *(*int32)(unsafe.Add(mBase, uint32(v8439)+192))
	if base.B2i32(v8453 == int32(0))|base.B2i32(base.Ui32(v8443) <= base.Ui32(v8453)) != 0 {
		v8479 = v8454
		goto L2294
	} else {
		goto L2301
	}
L2298:
	;
	v8446 = *(*int32)(unsafe.Add(mBase, uint32(v8439)+188))
	v8453 = v8446
	goto L2297
L2299:
	;
	goto L2300
L2300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8439)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8439)+184)) = v8440
	v8451 = F_strchr(m, v8440, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8439)+188)) = v8451
	v8453 = v8451
	goto L2297
L2301:
	;
	v8459 = v8453
	v8462 = v8454
	goto L2302
L2302:
	;
	v8464 = int32(1)
	v8465 = v8462 + v8464
	*(*int32)(unsafe.Add(mBase, uint32(v8439)+192)) = v8465
	v8468 = v8459 + v8464
	*(*int32)(unsafe.Add(mBase, uint32(v8439)+184)) = v8468
	v8471 = F_strchr(m, v8468, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8439)+188)) = v8471
	if v8471 == int32(0) {
		v8479 = v8465
		goto L2294
	} else {
		goto L2304
	}
L2303:
	;
	v8479 = v8465
	goto L2294
L2304:
	;
	if base.Ui32(v8471) < base.Ui32(v8443) {
		v8459 = v8471
		v8462 = v8465
		goto L2302
	} else {
		goto L2305
	}
L2305:
	;
	goto L2303
L2306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8498))) = int32(25)
	v8504 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v8505 = int32(0)
	if v8504 < v8505 {
		v8550 = v8505
		goto L2308
	} else {
		goto L2309
	}
L2307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8498)+4)) = v8550
	v8554 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v8555 = *(*int32)(unsafe.Add(mBase, uint32(v8554)+520))
	v8557 = v8555 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8554)+520)) = v8557
	*(*int32)(unsafe.Add(mBase, uint32(v8498)+8)) = v8557
	v8562 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8498)+12)) = uint8(base.B2i32(v8562 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8498
	v9817 = v233
	goto L13
L2308:
	;
	goto L2307
L2309:
	;
	v8510 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8511 = *(*int32)(unsafe.Add(mBase, uint32(v8510)+56))
	if v8511 == int32(0) {
		v8550 = v8505
		goto L2308
	} else {
		goto L2310
	}
L2310:
	;
	v8514 = v8504 + v8511
	v8515 = *(*int32)(unsafe.Add(mBase, uint32(v8510)+184))
	if base.Ui32(v8515) <= base.Ui32(v8514) {
		goto L2312
	} else {
		goto L2313
	}
L2311:
	;
	v8525 = *(*int32)(unsafe.Add(mBase, uint32(v8510)+192))
	if base.B2i32(v8524 == int32(0))|base.B2i32(base.Ui32(v8514) <= base.Ui32(v8524)) != 0 {
		v8550 = v8525
		goto L2308
	} else {
		goto L2315
	}
L2312:
	;
	v8517 = *(*int32)(unsafe.Add(mBase, uint32(v8510)+188))
	v8524 = v8517
	goto L2311
L2313:
	;
	goto L2314
L2314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8510)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8510)+184)) = v8511
	v8522 = F_strchr(m, v8511, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8510)+188)) = v8522
	v8524 = v8522
	goto L2311
L2315:
	;
	v8530 = v8524
	v8533 = v8525
	goto L2316
L2316:
	;
	v8535 = int32(1)
	v8536 = v8533 + v8535
	*(*int32)(unsafe.Add(mBase, uint32(v8510)+192)) = v8536
	v8539 = v8530 + v8535
	*(*int32)(unsafe.Add(mBase, uint32(v8510)+184)) = v8539
	v8542 = F_strchr(m, v8539, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8510)+188)) = v8542
	if v8542 == int32(0) {
		v8550 = v8536
		goto L2308
	} else {
		goto L2318
	}
L2317:
	;
	v8550 = v8536
	goto L2308
L2318:
	;
	if base.Ui32(v8542) < base.Ui32(v8514) {
		v8530 = v8542
		v8533 = v8536
		goto L2316
	} else {
		goto L2319
	}
L2319:
	;
	goto L2317
L2320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8568))) = int32(26)
	v8574 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v8575 = int32(0)
	if v8574 < v8575 {
		v8620 = v8575
		goto L2322
	} else {
		goto L2323
	}
L2321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8568)+4)) = v8620
	v8624 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v8625 = *(*int32)(unsafe.Add(mBase, uint32(v8624)+520))
	v8627 = v8625 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8624)+520)) = v8627
	*(*int32)(unsafe.Add(mBase, uint32(v8568)+8)) = v8627
	v8632 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8568)+12)) = uint8(base.B2i32(v8632 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8568
	v9817 = v233
	goto L13
L2322:
	;
	goto L2321
L2323:
	;
	v8580 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8581 = *(*int32)(unsafe.Add(mBase, uint32(v8580)+56))
	if v8581 == int32(0) {
		v8620 = v8575
		goto L2322
	} else {
		goto L2324
	}
L2324:
	;
	v8584 = v8574 + v8581
	v8585 = *(*int32)(unsafe.Add(mBase, uint32(v8580)+184))
	if base.Ui32(v8585) <= base.Ui32(v8584) {
		goto L2326
	} else {
		goto L2327
	}
L2325:
	;
	v8595 = *(*int32)(unsafe.Add(mBase, uint32(v8580)+192))
	if base.B2i32(v8594 == int32(0))|base.B2i32(base.Ui32(v8584) <= base.Ui32(v8594)) != 0 {
		v8620 = v8595
		goto L2322
	} else {
		goto L2329
	}
L2326:
	;
	v8587 = *(*int32)(unsafe.Add(mBase, uint32(v8580)+188))
	v8594 = v8587
	goto L2325
L2327:
	;
	goto L2328
L2328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8580)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8580)+184)) = v8581
	v8592 = F_strchr(m, v8581, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8580)+188)) = v8592
	v8594 = v8592
	goto L2325
L2329:
	;
	v8600 = v8594
	v8603 = v8595
	goto L2330
L2330:
	;
	v8605 = int32(1)
	v8606 = v8603 + v8605
	*(*int32)(unsafe.Add(mBase, uint32(v8580)+192)) = v8606
	v8609 = v8600 + v8605
	*(*int32)(unsafe.Add(mBase, uint32(v8580)+184)) = v8609
	v8612 = F_strchr(m, v8609, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8580)+188)) = v8612
	if v8612 == int32(0) {
		v8620 = v8606
		goto L2322
	} else {
		goto L2332
	}
L2331:
	;
	v8620 = v8606
	goto L2322
L2332:
	;
	if base.Ui32(v8612) < base.Ui32(v8584) {
		v8600 = v8612
		v8603 = v8606
		goto L2330
	} else {
		goto L2333
	}
L2333:
	;
	goto L2331
L2334:
	;
	v8645 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v8646 = m.ExcPending
	if v8646 != 0 {
		goto L49
	} else {
		goto L2335
	}
L2335:
	;
	if v8645 == int32(91) {
		goto L23
	} else {
		goto L2336
	}
L2336:
	;
	v8649 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v8650 = *(*int32)(unsafe.Add(mBase, uint32(v8649)+24))
	v8651 = *(*int32)(unsafe.Add(mBase, uint32(v8650)+4))
	if v8651 != int32(1790) {
		goto L22
	} else {
		goto L2337
	}
L2337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8649
	v9817 = v233
	goto L13
L2338:
	;
	v8706 = F_palloc(m, int32(12))
	mBase = m.M
	v8707 = m.ExcPending
	if v8707 != 0 {
		goto L49
	} else {
		goto L2351
	}
L2339:
	;
	goto L2338
L2340:
	;
	v8663 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8664 = *(*int32)(unsafe.Add(mBase, uint32(v8663)+56))
	if v8664 == int32(0) {
		v8703 = v8658
		goto L2339
	} else {
		goto L2341
	}
L2341:
	;
	v8667 = v8657 + v8664
	v8668 = *(*int32)(unsafe.Add(mBase, uint32(v8663)+184))
	if base.Ui32(v8668) <= base.Ui32(v8667) {
		goto L2343
	} else {
		goto L2344
	}
L2342:
	;
	v8678 = *(*int32)(unsafe.Add(mBase, uint32(v8663)+192))
	if base.B2i32(v8677 == int32(0))|base.B2i32(base.Ui32(v8667) <= base.Ui32(v8677)) != 0 {
		v8703 = v8678
		goto L2339
	} else {
		goto L2346
	}
L2343:
	;
	v8670 = *(*int32)(unsafe.Add(mBase, uint32(v8663)+188))
	v8677 = v8670
	goto L2342
L2344:
	;
	goto L2345
L2345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8663)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8663)+184)) = v8664
	v8675 = F_strchr(m, v8664, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8663)+188)) = v8675
	v8677 = v8675
	goto L2342
L2346:
	;
	v8683 = v8677
	v8686 = v8678
	goto L2347
L2347:
	;
	v8688 = int32(1)
	v8689 = v8686 + v8688
	*(*int32)(unsafe.Add(mBase, uint32(v8663)+192)) = v8689
	v8692 = v8683 + v8688
	*(*int32)(unsafe.Add(mBase, uint32(v8663)+184)) = v8692
	v8695 = F_strchr(m, v8692, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8663)+188)) = v8695
	if v8695 == int32(0) {
		v8703 = v8689
		goto L2339
	} else {
		goto L2349
	}
L2348:
	;
	v8703 = v8689
	goto L2339
L2349:
	;
	if base.Ui32(v8695) < base.Ui32(v8667) {
		v8683 = v8695
		v8686 = v8689
		goto L2347
	} else {
		goto L2350
	}
L2350:
	;
	goto L2348
L2351:
	;
	v8709 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v8710 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8709)+525)) = uint8(v8710)
	v8715 = *(*int32)(unsafe.Add(mBase, uint32(v8709)+44))
	v8717 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8715, int32(0))
	mBase = m.M
	v8718 = m.ExcPending
	if v8718 != 0 {
		goto L49
	} else {
		goto L2352
	}
L2352:
	;
	v8720 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_yyparse_17), v8703, v8717, int32(1))
	mBase = m.M
	v8721 = m.ExcPending
	if v8721 != 0 {
		goto L49
	} else {
		goto L2353
	}
L2353:
	;
	v8722 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8720)+16)) = uint8(v8722)
	v8724 = *(*int32)(unsafe.Add(mBase, uint32(v8720)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8706))) = v8724
	v8730 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v8731 = *(*int32)(unsafe.Add(mBase, uint32(v8730)+44))
	v8733 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8731, int32(0))
	mBase = m.M
	v8734 = m.ExcPending
	if v8734 != 0 {
		goto L49
	} else {
		goto L2354
	}
L2354:
	;
	v8736 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_yyparse_113), v8703, v8733, int32(1))
	mBase = m.M
	v8737 = m.ExcPending
	if v8737 != 0 {
		goto L49
	} else {
		goto L2355
	}
L2355:
	;
	v8738 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8736)+16)) = uint8(v8738)
	v8740 = *(*int32)(unsafe.Add(mBase, uint32(v8736)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8706)+4)) = v8740
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8706
	v9817 = v233
	goto L13
L2356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8753
	v9817 = v233
	goto L13
L2357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8762
	v9817 = v233
	goto L13
L2358:
	;
	v8770 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v8771 = int32(0)
	if v8770 < v8771 {
		v8816 = v8771
		goto L2360
	} else {
		goto L2361
	}
L2359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8766))) = v8816
	v8821 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8766)+4)) = v8821
	v8823 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v8766)+8)) = v8823
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8766
	v9817 = v233
	goto L13
L2360:
	;
	goto L2359
L2361:
	;
	v8776 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8777 = *(*int32)(unsafe.Add(mBase, uint32(v8776)+56))
	if v8777 == int32(0) {
		v8816 = v8771
		goto L2360
	} else {
		goto L2362
	}
L2362:
	;
	v8780 = v8770 + v8777
	v8781 = *(*int32)(unsafe.Add(mBase, uint32(v8776)+184))
	if base.Ui32(v8781) <= base.Ui32(v8780) {
		goto L2364
	} else {
		goto L2365
	}
L2363:
	;
	v8791 = *(*int32)(unsafe.Add(mBase, uint32(v8776)+192))
	if base.B2i32(v8790 == int32(0))|base.B2i32(base.Ui32(v8780) <= base.Ui32(v8790)) != 0 {
		v8816 = v8791
		goto L2360
	} else {
		goto L2367
	}
L2364:
	;
	v8783 = *(*int32)(unsafe.Add(mBase, uint32(v8776)+188))
	v8790 = v8783
	goto L2363
L2365:
	;
	goto L2366
L2366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8776)+192)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8776)+184)) = v8777
	v8788 = F_strchr(m, v8777, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8776)+188)) = v8788
	v8790 = v8788
	goto L2363
L2367:
	;
	v8796 = v8790
	v8799 = v8791
	goto L2368
L2368:
	;
	v8801 = int32(1)
	v8802 = v8799 + v8801
	*(*int32)(unsafe.Add(mBase, uint32(v8776)+192)) = v8802
	v8805 = v8796 + v8801
	*(*int32)(unsafe.Add(mBase, uint32(v8776)+184)) = v8805
	v8808 = F_strchr(m, v8805, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8776)+188)) = v8808
	if v8808 == int32(0) {
		v8816 = v8802
		goto L2360
	} else {
		goto L2370
	}
L2369:
	;
	v8816 = v8802
	goto L2360
L2370:
	;
	if base.Ui32(v8808) < base.Ui32(v8780) {
		v8796 = v8808
		v8799 = v8802
		goto L2368
	} else {
		goto L2371
	}
L2371:
	;
	goto L2369
L2372:
	;
	v8854 = *(*int32)(unsafe.Add(mBase, uint32(v8833)+8))
	if v8854 != 0 {
		v8833 = v8854
		goto L2372
	} else {
		goto L2374
	}
L2373:
	;
	v8855 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v8833)+8)) = v8855
	v8857 = *(*int32)(unsafe.Add(mBase, uint32(v8827)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8857
	v9817 = v233
	goto L13
L2374:
	;
	goto L2373
L2375:
	;
	if v8886-v8887 != 0 {
		goto L2382
	} else {
		goto L2383
	}
L2376:
	;
	goto L2375
L2377:
	;
	v8871 = v8861
	v8872 = v8862
	goto L2378
L2378:
	;
	v8875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8872)+1)))
	v8876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8871)+1)))
	if v8876 == int32(0) {
		v8886 = v8876
		v8887 = v8875
		goto L2376
	} else {
		goto L2380
	}
L2379:
	;
	v8886 = v8876
	v8887 = v8875
	goto L2376
L2380:
	;
	v8879 = int32(1)
	if v8876 == v8875 {
		v8871 = v8871 + v8879
		v8872 = v8872 + v8879
		goto L2378
	} else {
		goto L2381
	}
L2381:
	;
	goto L2379
L2382:
	;
	v8889 = F_plpgsql_parse_err_condition(m, v8861)
	mBase = m.M
	v8890 = m.ExcPending
	if v8890 != 0 {
		goto L49
	} else {
		goto L2385
	}
L2383:
	;
	goto L2384
L2384:
	;
	v8896 = F_plpgsql_yylex(m, v28+int32(_a_F_plpgsql_yyparse_8), v28+int32(_a_F_plpgsql_yyparse_2), l1)
	mBase = m.M
	v8897 = m.ExcPending
	if v8897 != 0 {
		goto L49
	} else {
		goto L2386
	}
L2385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8889
	v9817 = v233
	goto L13
L2386:
	;
	if v8896 != int32(261) {
		goto L3
	} else {
		goto L2387
	}
L2387:
	;
	v8900 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	v8901 = F_strlen(m, v8900)
	mBase = m.M
	if v8901 != int32(5) {
		goto L1
	} else {
		goto L2388
	}
L2388:
	;
	v8904 = int32(_a_F_plpgsql_yyparse_82)
	v8908 = m.G0
	v8910 = v8908 - int32(32)
	v8911 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8910)+24)) = v8911
	*(*int64)(unsafe.Add(mBase, uint32(v8910)+16)) = v8911
	*(*int64)(unsafe.Add(mBase, uint32(v8910)+8)) = v8911
	*(*int64)(unsafe.Add(mBase, uint32(v8910))) = v8911
	v8919 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[51])))
	if v8919 == int32(0) {
		goto L2390
	} else {
		goto L2391
	}
L2389:
	;
	if v8987 != int32(5) {
		goto L1
	} else {
		goto L2408
	}
L2390:
	;
	v8987 = int32(0)
	goto L2389
L2391:
	;
	goto L2392
L2392:
	;
	v8923 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[52])))
	if v8923 == int32(0) {
		goto L2393
	} else {
		goto L2394
	}
L2393:
	;
	v8927 = v8900
	goto L2396
L2394:
	;
	goto L2395
L2395:
	;
	v8937 = v8904
	v8938 = v8919
	goto L2399
L2396:
	;
	v8933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8927))))
	if v8933 == v8919 {
		v8927 = v8927 + int32(1)
		goto L2396
	} else {
		goto L2398
	}
L2397:
	;
	v8987 = v8927 - v8900
	goto L2389
L2398:
	;
	goto L2397
L2399:
	;
	v8945 = v8910 + int32(base.Ui32(v8938)>>(uint(int32(3))%32))&int32(28)
	v8946 = *(*int32)(unsafe.Add(mBase, uint32(v8945)))
	v8947 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8945))) = v8946 | v8947<<(uint(v8938)%32)
	v8951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8937)+1)))
	if v8951 != 0 {
		v8937 = v8937 + v8947
		v8938 = v8951
		goto L2399
	} else {
		goto L2401
	}
L2400:
	;
	v8954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8900))))
	if v8954 == int32(0) {
		v8977 = v8900
		goto L2402
	} else {
		goto L2403
	}
L2401:
	;
	goto L2400
L2402:
	;
	v8987 = v8977 - v8900
	goto L2389
L2403:
	;
	v8958 = v8900
	v8959 = v8954
	goto L2404
L2404:
	;
	v8967 = *(*int32)(unsafe.Add(mBase, uint32(v8910+int32(base.Ui32(v8959)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v8967)>>(uint(v8959)%32))&int32(1) == int32(0) {
		v8977 = v8958
		goto L2402
	} else {
		goto L2406
	}
L2405:
	;
	v8977 = v8975
	goto L2402
L2406:
	;
	v8973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8958)+1)))
	v8975 = v8958 + int32(1)
	if v8973 != 0 {
		v8958 = v8975
		v8959 = v8973
		goto L2404
	} else {
		goto L2407
	}
L2407:
	;
	goto L2405
L2408:
	;
	v8991 = F_palloc(m, int32(12))
	mBase = m.M
	v8992 = m.ExcPending
	if v8992 != 0 {
		goto L49
	} else {
		goto L2409
	}
L2409:
	;
	v8993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8900)+4)))
	v8994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8900)+3)))
	v8995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8900)+2)))
	v8996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8900)+1)))
	v8997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8900))))
	*(*int32)(unsafe.Add(mBase, uint32(v8991)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8991)+4)) = v8900
	v9001 = int32(16)
	v9003 = int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v8991))) = (v8997+v9001)&v9003 | (v8996+v9001)&v9003<<(uint(int32(6))%32) | (v8995+v9001)&v9003<<(uint(int32(12))%32) | (v8994+v9001)&v9003<<(uint(int32(18))%32) | (v8993+v9001)&v9003<<(uint(int32(24))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v8991
	v9817 = v233
	goto L13
L2410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9048
	v9817 = v233
	goto L13
L2411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9064
	v9817 = v233
	goto L13
L2412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9080
	v9817 = v233
	goto L13
L2413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L2414:
	;
	v9095 = *(*int32)(unsafe.Add(mBase, uint32(v9090)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9095
	v9817 = v233
	goto L13
L2415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = int32(0)
	v9817 = v233
	goto L13
L2416:
	;
	v9109 = *(*int32)(unsafe.Add(mBase, uint32(v9104)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9109
	v9817 = v233
	goto L13
L2417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9122
	v9817 = v233
	goto L13
L2418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9125
	v9817 = v233
	goto L13
L2419:
	;
	v9131 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[2])))
	if v67 != int32(3) {
		v9500 = v224
		v9502 = v9131
		goto L20
	} else {
		goto L2420
	}
L2420:
	;
	if int32(0) < v224 {
		goto L2421
	} else {
		goto L2422
	}
L2421:
	;
	v9500 = int32(-2)
	v9502 = v9131
	goto L20
L2422:
	;
	goto L2423
L2423:
	;
	if v224 != 0 {
		v9500 = v224
		v9502 = v9131
		goto L20
	} else {
		goto L2424
	}
L2424:
	;
	v9947 = int32(1)
	v9957 = v153
	goto L7
L2425:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2426:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9151 = m.ExcPending
	if v9151 != 0 {
		goto L49
	} else {
		goto L2427
	}
L2427:
	;
	v9152 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	v9153 = *(*int32)(unsafe.Add(mBase, uint32(v9152)+4))
	v9154 = F_format_type_be(m, v9153)
	mBase = m.M
	v9155 = m.ExcPending
	if v9155 != 0 {
		goto L49
	} else {
		goto L2428
	}
L2428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v9154
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_114), v28+int32(32))
	mBase = m.M
	v9161 = m.ExcPending
	if v9161 != 0 {
		goto L49
	} else {
		goto L2429
	}
L2429:
	;
	v9164 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(8))))
	v9165 = F_plpgsql_scanner_errposition(m, v9164, l1)
	mBase = m.M
	v9166 = m.ExcPending
	if v9166 != 0 {
		goto L49
	} else {
		goto L2430
	}
L2430:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(523), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9171 = m.ExcPending
	if v9171 != 0 {
		goto L49
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v9178 = m.ExcPending
	if v9178 != 0 {
		goto L49
	} else {
		goto L2433
	}
L2433:
	;
	v9179 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v9179
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_115), v28+int32(16))
	mBase = m.M
	v9185 = m.ExcPending
	if v9185 != 0 {
		goto L49
	} else {
		goto L2434
	}
L2434:
	;
	v9188 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9189 = F_plpgsql_scanner_errposition(m, v9188, l1)
	mBase = m.M
	v9190 = m.ExcPending
	if v9190 != 0 {
		goto L49
	} else {
		goto L2435
	}
L2435:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(542), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9195 = m.ExcPending
	if v9195 != 0 {
		goto L49
	} else {
		goto L2436
	}
L2436:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2437:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9202 = m.ExcPending
	if v9202 != 0 {
		goto L49
	} else {
		goto L2438
	}
L2438:
	;
	v9203 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v9203
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_116), v28+int32(48))
	mBase = m.M
	v9209 = m.ExcPending
	if v9209 != 0 {
		goto L49
	} else {
		goto L2439
	}
L2439:
	;
	v9210 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9211 = F_plpgsql_scanner_errposition(m, v9210, l1)
	mBase = m.M
	v9212 = m.ExcPending
	if v9212 != 0 {
		goto L49
	} else {
		goto L2440
	}
L2440:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(667), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9217 = m.ExcPending
	if v9217 != 0 {
		goto L49
	} else {
		goto L2441
	}
L2441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2442:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9224 = m.ExcPending
	if v9224 != 0 {
		goto L49
	} else {
		goto L2443
	}
L2443:
	;
	v9225 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v9225
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_116), v28-int32(-64))
	mBase = m.M
	v9231 = m.ExcPending
	if v9231 != 0 {
		goto L49
	} else {
		goto L2444
	}
L2444:
	;
	v9232 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9233 = F_plpgsql_scanner_errposition(m, v9232, l1)
	mBase = m.M
	v9234 = m.ExcPending
	if v9234 != 0 {
		goto L49
	} else {
		goto L2445
	}
L2445:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(682), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9239 = m.ExcPending
	if v9239 != 0 {
		goto L49
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
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9247 = m.ExcPending
	if v9247 != 0 {
		goto L49
	} else {
		goto L2448
	}
L2448:
	;
	v9248 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v9249 = F_NameListToString(m, v9248)
	mBase = m.M
	v9250 = m.ExcPending
	if v9250 != 0 {
		goto L49
	} else {
		goto L2449
	}
L2449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v9249
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_116), v28+int32(80))
	mBase = m.M
	v9256 = m.ExcPending
	if v9256 != 0 {
		goto L49
	} else {
		goto L2450
	}
L2450:
	;
	v9257 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9258 = F_plpgsql_scanner_errposition(m, v9257, l1)
	mBase = m.M
	v9259 = m.ExcPending
	if v9259 != 0 {
		goto L49
	} else {
		goto L2451
	}
L2451:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(708), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9264 = m.ExcPending
	if v9264 != 0 {
		goto L49
	} else {
		goto L2452
	}
L2452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2453:
	;
	F_errmsg_internal(m, int32(_a_F_plpgsql_yyparse_117), int32(0))
	mBase = m.M
	v9273 = m.ExcPending
	if v9273 != 0 {
		goto L49
	} else {
		goto L2454
	}
L2454:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(990), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9278 = m.ExcPending
	if v9278 != 0 {
		goto L49
	} else {
		goto L2455
	}
L2455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2456:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9285 = m.ExcPending
	if v9285 != 0 {
		goto L49
	} else {
		goto L2457
	}
L2457:
	;
	v9286 = F_NameOfDatum(m, v148)
	mBase = m.M
	v9287 = m.ExcPending
	if v9287 != 0 {
		goto L49
	} else {
		goto L2458
	}
L2458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = v9286
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_118), v28+int32(176))
	mBase = m.M
	v9293 = m.ExcPending
	if v9293 != 0 {
		goto L49
	} else {
		goto L2459
	}
L2459:
	;
	v9294 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9295 = F_plpgsql_scanner_errposition(m, v9294, l1)
	mBase = m.M
	v9296 = m.ExcPending
	if v9296 != 0 {
		goto L49
	} else {
		goto L2460
	}
L2460:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1174), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9301 = m.ExcPending
	if v9301 != 0 {
		goto L49
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9308 = m.ExcPending
	if v9308 != 0 {
		goto L49
	} else {
		goto L2463
	}
L2463:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_119), int32(0))
	mBase = m.M
	v9312 = m.ExcPending
	if v9312 != 0 {
		goto L49
	} else {
		goto L2464
	}
L2464:
	;
	v9313 = *(*int32)(unsafe.Add(mBase, uint32(v4834)))
	v9314 = F_plpgsql_scanner_errposition(m, v9313, l1)
	mBase = m.M
	v9315 = m.ExcPending
	if v9315 != 0 {
		goto L49
	} else {
		goto L2465
	}
L2465:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1404), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9320 = m.ExcPending
	if v9320 != 0 {
		goto L49
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
	v9327 = m.ExcPending
	if v9327 != 0 {
		goto L49
	} else {
		goto L2468
	}
L2468:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_120), int32(0))
	mBase = m.M
	v9331 = m.ExcPending
	if v9331 != 0 {
		goto L49
	} else {
		goto L2469
	}
L2469:
	;
	v9334 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9335 = F_plpgsql_scanner_errposition(m, v9334, l1)
	mBase = m.M
	v9336 = m.ExcPending
	if v9336 != 0 {
		goto L49
	} else {
		goto L2470
	}
L2470:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1439), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9341 = m.ExcPending
	if v9341 != 0 {
		goto L49
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
	v9348 = m.ExcPending
	if v9348 != 0 {
		goto L49
	} else {
		goto L2473
	}
L2473:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_121), int32(0))
	mBase = m.M
	v9352 = m.ExcPending
	if v9352 != 0 {
		goto L49
	} else {
		goto L2474
	}
L2474:
	;
	v9353 = F_plpgsql_scanner_errposition(m, v4766, l1)
	mBase = m.M
	v9354 = m.ExcPending
	if v9354 != 0 {
		goto L49
	} else {
		goto L2475
	}
L2475:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1446), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9359 = m.ExcPending
	if v9359 != 0 {
		goto L49
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
	v9366 = m.ExcPending
	if v9366 != 0 {
		goto L49
	} else {
		goto L2478
	}
L2478:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_122), int32(0))
	mBase = m.M
	v9370 = m.ExcPending
	if v9370 != 0 {
		goto L49
	} else {
		goto L2479
	}
L2479:
	;
	v9373 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(20))))
	v9374 = F_plpgsql_scanner_errposition(m, v9373, l1)
	mBase = m.M
	v9375 = m.ExcPending
	if v9375 != 0 {
		goto L49
	} else {
		goto L2480
	}
L2480:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1702), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9380 = m.ExcPending
	if v9380 != 0 {
		goto L49
	} else {
		goto L2481
	}
L2481:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2482:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9387 = m.ExcPending
	if v9387 != 0 {
		goto L49
	} else {
		goto L2483
	}
L2483:
	;
	v9388 = *(*int32)(unsafe.Add(mBase, uint32(v5327)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+192)) = v9388
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_123), v28+int32(192))
	mBase = m.M
	v9394 = m.ExcPending
	if v9394 != 0 {
		goto L49
	} else {
		goto L2484
	}
L2484:
	;
	v9397 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9398 = F_plpgsql_scanner_errposition(m, v9397, l1)
	mBase = m.M
	v9399 = m.ExcPending
	if v9399 != 0 {
		goto L49
	} else {
		goto L2485
	}
L2485:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1746), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9404 = m.ExcPending
	if v9404 != 0 {
		goto L49
	} else {
		goto L2486
	}
L2486:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2487:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9411 = m.ExcPending
	if v9411 != 0 {
		goto L49
	} else {
		goto L2488
	}
L2488:
	;
	v9414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5260)+12)))
	if v9414 != 0 {
		goto L2489
	} else {
		goto L2490
	}
L2489:
	;
	v9415 = int32(_a_F_plpgsql_yyparse_124)
	goto L2491
L2490:
	;
	v9415 = int32(_a_F_plpgsql_yyparse_125)
	goto L2491
L2491:
	;
	F_errmsg(m, v9415, int32(0))
	mBase = m.M
	v9418 = m.ExcPending
	if v9418 != 0 {
		goto L49
	} else {
		goto L2492
	}
L2492:
	;
	v9419 = *(*int32)(unsafe.Add(mBase, uint32(v5276)))
	v9420 = F_plpgsql_scanner_errposition(m, v9419, l1)
	mBase = m.M
	v9421 = m.ExcPending
	if v9421 != 0 {
		goto L49
	} else {
		goto L2493
	}
L2493:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1768), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9426 = m.ExcPending
	if v9426 != 0 {
		goto L49
	} else {
		goto L2494
	}
L2494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2496:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2497:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v9437 = m.ExcPending
	if v9437 != 0 {
		goto L49
	} else {
		goto L2498
	}
L2498:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_126), int32(0))
	mBase = m.M
	v9441 = m.ExcPending
	if v9441 != 0 {
		goto L49
	} else {
		goto L2499
	}
L2499:
	;
	v9444 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(12))))
	v9445 = F_plpgsql_scanner_errposition(m, v9444, l1)
	mBase = m.M
	v9446 = m.ExcPending
	if v9446 != 0 {
		goto L49
	} else {
		goto L2500
	}
L2500:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(2199), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9451 = m.ExcPending
	if v9451 != 0 {
		goto L49
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9458 = m.ExcPending
	if v9458 != 0 {
		goto L49
	} else {
		goto L2503
	}
L2503:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_127), int32(0))
	mBase = m.M
	v9462 = m.ExcPending
	if v9462 != 0 {
		goto L49
	} else {
		goto L2504
	}
L2504:
	;
	v9463 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9464 = F_plpgsql_scanner_errposition(m, v9463, l1)
	mBase = m.M
	v9465 = m.ExcPending
	if v9465 != 0 {
		goto L49
	} else {
		goto L2505
	}
L2505:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(2296), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9470 = m.ExcPending
	if v9470 != 0 {
		goto L49
	} else {
		goto L2506
	}
L2506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2507:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9477 = m.ExcPending
	if v9477 != 0 {
		goto L49
	} else {
		goto L2508
	}
L2508:
	;
	v9478 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v9479 = *(*int32)(unsafe.Add(mBase, uint32(v9478)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+224)) = v9479
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_128), v28+int32(224))
	mBase = m.M
	v9485 = m.ExcPending
	if v9485 != 0 {
		goto L49
	} else {
		goto L2509
	}
L2509:
	;
	v9486 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v9487 = F_plpgsql_scanner_errposition(m, v9486, l1)
	mBase = m.M
	v9488 = m.ExcPending
	if v9488 != 0 {
		goto L49
	} else {
		goto L2510
	}
L2510:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(2303), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9493 = m.ExcPending
	if v9493 != 0 {
		goto L49
	} else {
		goto L2511
	}
L2511:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2512:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2513:
	;
	if v9505&int32(_a_F_plpgsql_yyparse_129) == int32(_a_F_plpgsql_yyparse_14) {
		goto L2516
	} else {
		goto L2517
	}
L2514:
	;
	v9566 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v9509)+24)) = v9566
	v9568 = *(*int64)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_plpgsql_yyparse[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v9509)+16)) = v9568
	*(*int32)(unsafe.Add(mBase, uint32(v9507)+4)) = v9518
	v9881 = v9546
	v9882 = v9509 + int32(16)
	v9886 = v9500
	v9888 = v9507 + int32(4)
	v9889 = v9516
	v9895 = int32(3)
	goto L12
L2515:
	;
	goto L2514
L2516:
	;
	if v9516 == v153 {
		goto L2521
	} else {
		goto L2522
	}
L2517:
	;
	v9532 = base.I32_extend16_s(v9505)
	if v9532 < int32(-1) {
		goto L2516
	} else {
		goto L2518
	}
L2518:
	;
	v9535 = int32(1)
	v9538 = (v9532 + v9535) << (uint(v9535) % 32)
	v9541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9538)+uint32(_c_F_plpgsql_yyparse[10]))))
	if v9541 != v9535 {
		goto L2516
	} else {
		goto L2519
	}
L2519:
	;
	v9546 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9538)+uint32(_c_F_plpgsql_yyparse[11]))))
	if int32(0) < v9546 {
		goto L2515
	} else {
		goto L2520
	}
L2520:
	;
	goto L2516
L2521:
	;
	v9947 = int32(1)
	v9957 = v153
	goto L7
L2522:
	;
	v9554 = v9516 - int32(2)
	v9555 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9554))))
	v9560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9555<<(uint(int32(1))%32))+uint32(_c_F_plpgsql_yyparse[6]))))
	v9563 = *(*int32)(unsafe.Add(mBase, uint32(v9507)))
	v9505 = v9560
	v9507 = v9507 - int32(4)
	v9509 = v9509 - int32(16)
	v9516 = v9554
	v9518 = v9563
	goto L2513
L2524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2525:
	;
	v9587 = int32(_a_F_plpgsql_yyparse_130)
	v9590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9584))))
	v9593 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[68])))
	if base.B2i32(v9590 == int32(0))|base.B2i32(v9590 != v9593) != 0 {
		v9611 = v9590
		v9612 = v9593
		goto L2527
	} else {
		goto L2528
	}
L2526:
	;
	if v9611-v9612 != 0 {
		goto L16
	} else {
		goto L2533
	}
L2527:
	;
	goto L2526
L2528:
	;
	v9596 = v9584
	v9597 = v9587
	goto L2529
L2529:
	;
	v9600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9597)+1)))
	v9601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9596)+1)))
	if v9601 == int32(0) {
		v9611 = v9601
		v9612 = v9600
		goto L2527
	} else {
		goto L2531
	}
L2530:
	;
	v9611 = v9601
	v9612 = v9600
	goto L2527
L2531:
	;
	v9604 = int32(1)
	if v9601 == v9600 {
		v9596 = v9596 + v9604
		v9597 = v9597 + v9604
		goto L2529
	} else {
		goto L2532
	}
L2532:
	;
	goto L2530
L2533:
	;
	goto L17
L2534:
	;
	v9634 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v9634 == int32(269) {
		v9680 = v9632
		v9681 = v9616
		goto L15
	} else {
		goto L2535
	}
L2535:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yyparse_7))
	mBase = m.M
	v9640 = m.ExcPending
	if v9640 != 0 {
		goto L49
	} else {
		goto L2536
	}
L2536:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9643 = m.ExcPending
	if v9643 != 0 {
		goto L49
	} else {
		goto L2537
	}
L2537:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_131), int32(0))
	mBase = m.M
	v9647 = m.ExcPending
	if v9647 != 0 {
		goto L49
	} else {
		goto L2538
	}
L2538:
	;
	v9648 = F_plpgsql_scanner_errposition(m, v4766, l1)
	mBase = m.M
	v9649 = m.ExcPending
	if v9649 != 0 {
		goto L49
	} else {
		goto L2539
	}
L2539:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1570), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9654 = m.ExcPending
	if v9654 != 0 {
		goto L49
	} else {
		goto L2540
	}
L2540:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2541:
	;
	v9663 = int32(0)
	v9675 = F_read_sql_construct(m, int32(269), int32(336), v9663, int32(_a_F_plpgsql_yyparse_9), v9663, int32(1), v9663, v28+int32(240), v28+int32(256), v9658, v9660, l1)
	mBase = m.M
	v9676 = m.ExcPending
	if v9676 != 0 {
		goto L49
	} else {
		goto L2542
	}
L2542:
	;
	v9677 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v9677 != int32(269) {
		goto L14
	} else {
		goto L2543
	}
L2543:
	;
	v9680 = v9675
	v9681 = v9663
	goto L15
L2544:
	;
	v9691 = int32(0)
	v9697 = int32(1)
	v9703 = v28 + int32(_a_F_plpgsql_yyparse_8)
	v9705 = v28 + int32(_a_F_plpgsql_yyparse_2)
	v9706 = F_read_sql_construct(m, int32(336), int32(288), v9691, int32(_a_F_plpgsql_yyparse_9), int32(2), v9697, v9697, v9691, v28+int32(256), v9703, v9705, l1)
	mBase = m.M
	v9707 = m.ExcPending
	if v9707 != 0 {
		goto L49
	} else {
		goto L2545
	}
L2545:
	;
	v9708 = *(*int32)(unsafe.Add(mBase, uint32(v28)+256))
	if v9708 == int32(288) {
		goto L2546
	} else {
		goto L2547
	}
L2546:
	;
	v9712 = int32(0)
	v9716 = int32(1)
	v9720 = F_read_sql_construct(m, int32(336), v9712, v9712, int32(_a_F_plpgsql_yyparse_9), int32(2), v9716, v9716, v9712, v9712, v9703, v9705, l1)
	mBase = m.M
	v9721 = m.ExcPending
	if v9721 != 0 {
		goto L49
	} else {
		goto L2549
	}
L2547:
	;
	v9722 = v9691
	goto L2548
L2548:
	;
	v9725 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(8))))
	if v9725 != 0 {
		goto L2550
	} else {
		goto L2551
	}
L2549:
	;
	v9722 = v9720
	goto L2548
L2550:
	;
	v9728 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(4))))
	if v9728 != 0 {
		goto L11
	} else {
		goto L2553
	}
L2551:
	;
	goto L2552
L2552:
	;
	v9731 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v9734 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v9737 = int32(0)
	v9739 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v9737, v9737)
	mBase = m.M
	v9740 = m.ExcPending
	if v9740 != 0 {
		goto L49
	} else {
		goto L2554
	}
L2553:
	;
	goto L2552
L2554:
	;
	v9742 = F_plpgsql_build_variable(m, v9731, v9734, v9739, int32(1))
	mBase = m.M
	v9743 = m.ExcPending
	if v9743 != 0 {
		goto L49
	} else {
		goto L2555
	}
L2555:
	;
	v9745 = F_palloc0(m, int32(40))
	mBase = m.M
	v9746 = m.ExcPending
	if v9746 != 0 {
		goto L49
	} else {
		goto L2556
	}
L2556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9745))) = int32(6)
	v9750 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v9751 = *(*int32)(unsafe.Add(mBase, uint32(v9750)+520))
	v9753 = v9751 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9750)+520)) = v9753
	*(*int32)(unsafe.Add(mBase, uint32(v9745)+32)) = v9681
	*(*int32)(unsafe.Add(mBase, uint32(v9745)+16)) = v9742
	*(*int32)(unsafe.Add(mBase, uint32(v9745)+8)) = v9753
	*(*int32)(unsafe.Add(mBase, uint32(v9745)+28)) = v9722
	*(*int32)(unsafe.Add(mBase, uint32(v9745)+24)) = v9706
	*(*int32)(unsafe.Add(mBase, uint32(v9745)+20)) = v9680
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9745
	v9817 = v233
	goto L13
L2557:
	;
	v9768 = F_palloc0(m, int32(28))
	mBase = m.M
	v9769 = m.ExcPending
	if v9769 != 0 {
		goto L49
	} else {
		goto L2558
	}
L2558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9768))) = int32(7)
	v9773 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_yyparse[15]))
	v9774 = *(*int32)(unsafe.Add(mBase, uint32(v9773)+520))
	v9776 = v9774 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9773)+520)) = v9776
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+8)) = v9776
	v9779 = int32(4)
	v9780 = v147 - v9779
	v9782 = v148 - v9779
	v9783 = *(*int32)(unsafe.Add(mBase, uint32(v9782)))
	if v9783 != 0 {
		goto L2560
	} else {
		goto L2561
	}
L2559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+24)) = v9675
	*(*int32)(unsafe.Add(mBase, uint32(v28)+272)) = v9768
	v9817 = v233
	goto L13
L2560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+16)) = v9783
	v9785 = *(*int32)(unsafe.Add(mBase, uint32(v9782)))
	v9786 = *(*int32)(unsafe.Add(mBase, uint32(v9780)))
	F_check_assignable(m, v9785, v9786, l1)
	mBase = m.M
	v9788 = m.ExcPending
	if v9788 != 0 {
		goto L49
	} else {
		goto L2563
	}
L2561:
	;
	goto L2562
L2562:
	;
	v9791 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(8))))
	if v9791 == int32(0) {
		goto L8
	} else {
		goto L2564
	}
L2563:
	;
	goto L2559
L2564:
	;
	v9796 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(16))))
	v9799 = *(*int32)(unsafe.Add(mBase, uint32(v148-int32(12))))
	v9800 = *(*int32)(unsafe.Add(mBase, uint32(v9780)))
	v9801 = F_make_scalar_list1(m, v9796, v9791, v9799, v9800, l1)
	mBase = m.M
	v9802 = m.ExcPending
	if v9802 != 0 {
		goto L49
	} else {
		goto L2565
	}
L2565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+16)) = v9801
	goto L2559
L2566:
	;
	v9875 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9854)+uint32(_c_F_plpgsql_yyparse[69]))))
	v9881 = v9875
	v9882 = v9843
	v9886 = v9817
	v9888 = v252
	v9889 = v9846
	v9895 = v67
	goto L12
L2567:
	;
	v9862 = v9858 << (uint(int32(1)) % 32)
	v9865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9862)+uint32(_c_F_plpgsql_yyparse[10]))))
	if v9865 != v9847&int32(_a_F_plpgsql_yyparse_129) {
		goto L2566
	} else {
		goto L2568
	}
L2568:
	;
	v9871 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9862)+uint32(_c_F_plpgsql_yyparse[11]))))
	v9881 = v9871
	v9882 = v9843
	v9886 = v9817
	v9888 = v252
	v9889 = v9846
	v9895 = v67
	goto L12
L2569:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9909 = m.ExcPending
	if v9909 != 0 {
		goto L49
	} else {
		goto L2570
	}
L2570:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_132), int32(0))
	mBase = m.M
	v9913 = m.ExcPending
	if v9913 != 0 {
		goto L49
	} else {
		goto L2571
	}
L2571:
	;
	v9916 = *(*int32)(unsafe.Add(mBase, uint32(v147-int32(4))))
	v9917 = F_plpgsql_scanner_errposition(m, v9916, l1)
	mBase = m.M
	v9918 = m.ExcPending
	if v9918 != 0 {
		goto L49
	} else {
		goto L2572
	}
L2572:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1536), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9923 = m.ExcPending
	if v9923 != 0 {
		goto L49
	} else {
		goto L2573
	}
L2573:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2574:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9930 = m.ExcPending
	if v9930 != 0 {
		goto L49
	} else {
		goto L2575
	}
L2575:
	;
	F_errmsg(m, int32(_a_F_plpgsql_yyparse_119), int32(0))
	mBase = m.M
	v9934 = m.ExcPending
	if v9934 != 0 {
		goto L49
	} else {
		goto L2576
	}
L2576:
	;
	v9935 = *(*int32)(unsafe.Add(mBase, uint32(v9780)))
	v9936 = F_plpgsql_scanner_errposition(m, v9935, l1)
	mBase = m.M
	v9937 = m.ExcPending
	if v9937 != 0 {
		goto L49
	} else {
		goto L2577
	}
L2577:
	;
	F_errfinish(m, int32(_a_F_plpgsql_yyparse_21), int32(1597), int32(_a_F_plpgsql_yyparse_22))
	mBase = m.M
	v9942 = m.ExcPending
	if v9942 != 0 {
		goto L49
	} else {
		goto L2578
	}
L2578:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2579:
	;
	F_pfree(m, v9957)
	mBase = m.M
	v9972 = m.ExcPending
	if v9972 != 0 {
		goto L49
	} else {
		goto L2582
	}
L2580:
	;
	goto L2581
L2581:
	;
	m.G0 = v28 + int32(_a_F_plpgsql_yyparse_0)
	return v9947
L2582:
	;
	goto L2581
L2583:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2584:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2585:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2587:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2588:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
