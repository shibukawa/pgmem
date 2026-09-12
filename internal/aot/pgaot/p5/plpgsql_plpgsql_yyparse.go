package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_plpgsql_yyparse(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v194 int64
	_ = v194
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v614 int32
	_ = v614
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v777 int32
	_ = v777
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v850 int32
	_ = v850
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v928 int32
	_ = v928
	var v970 int32
	_ = v970
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v998 int32
	_ = v998
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1076 int32
	_ = v1076
	var v1118 int32
	_ = v1118
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
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
	var v1156 int32
	_ = v1156
	var v1165 int32
	_ = v1165
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1276 int32
	_ = v1276
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1306 int32
	_ = v1306
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1342 int32
	_ = v1342
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1426 int32
	_ = v1426
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1626 int32
	_ = v1626
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1644 int32
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1663 int32
	_ = v1663
	var v1672 int32
	_ = v1672
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1741 int32
	_ = v1741
	var v1783 int32
	_ = v1783
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1874 int32
	_ = v1874
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1901 int32
	_ = v1901
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1913 int32
	_ = v1913
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v2012 int32
	_ = v2012
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2049 int32
	_ = v2049
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2070 int32
	_ = v2070
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2085 int32
	_ = v2085
	var v2127 int32
	_ = v2127
	var v2169 int32
	_ = v2169
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2228 int32
	_ = v2228
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2326 int32
	_ = v2326
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2372 int32
	_ = v2372
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2458 int32
	_ = v2458
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2509 int32
	_ = v2509
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2540 int32
	_ = v2540
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2559 int32
	_ = v2559
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2617 int32
	_ = v2617
	var v2618 int32
	_ = v2618
	var v2624 int32
	_ = v2624
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2692 int32
	_ = v2692
	var v2697 int32
	_ = v2697
	var v2701 int32
	_ = v2701
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2734 int32
	_ = v2734
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2740 int32
	_ = v2740
	var v2742 int32
	_ = v2742
	var v2744 int32
	_ = v2744
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2762 int32
	_ = v2762
	var v2764 int32
	_ = v2764
	var v2766 int32
	_ = v2766
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2772 int32
	_ = v2772
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2797 int32
	_ = v2797
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2804 int32
	_ = v2804
	var v2806 int32
	_ = v2806
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2823 int32
	_ = v2823
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2831 int32
	_ = v2831
	var v2840 int32
	_ = v2840
	var v2842 int32
	_ = v2842
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2873 int32
	_ = v2873
	var v2880 int32
	_ = v2880
	var v2883 int32
	_ = v2883
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2903 int32
	_ = v2903
	var v2910 int32
	_ = v2910
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2919 int32
	_ = v2919
	var v2925 int32
	_ = v2925
	var v2932 int32
	_ = v2932
	var v2936 int32
	_ = v2936
	var v2939 int32
	_ = v2939
	var v2945 int32
	_ = v2945
	var v2952 int32
	_ = v2952
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2969 int32
	_ = v2969
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2978 int32
	_ = v2978
	var v2980 int32
	_ = v2980
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2986 int32
	_ = v2986
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2995 int32
	_ = v2995
	var v2997 int32
	_ = v2997
	var v3002 int32
	_ = v3002
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3026 int32
	_ = v3026
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3088 int32
	_ = v3088
	var v3090 int32
	_ = v3090
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3097 int32
	_ = v3097
	var v3100 int32
	_ = v3100
	var v3107 int32
	_ = v3107
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3144 int32
	_ = v3144
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3163 int32
	_ = v3163
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3178 int32
	_ = v3178
	var v3180 int32
	_ = v3180
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3187 int32
	_ = v3187
	var v3190 int32
	_ = v3190
	var v3197 int32
	_ = v3197
	var v3202 int32
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3205 int32
	_ = v3205
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3245 int32
	_ = v3245
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3257 int32
	_ = v3257
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3269 int32
	_ = v3269
	var v3270 int32
	_ = v3270
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3288 int32
	_ = v3288
	var v3290 int32
	_ = v3290
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3300 int32
	_ = v3300
	var v3307 int32
	_ = v3307
	var v3312 int32
	_ = v3312
	var v3313 int32
	_ = v3313
	var v3315 int32
	_ = v3315
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3327 int32
	_ = v3327
	var v3329 int32
	_ = v3329
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3347 int32
	_ = v3347
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3362 int32
	_ = v3362
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3374 int32
	_ = v3374
	var v3379 int32
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3382 int32
	_ = v3382
	var v3386 int32
	_ = v3386
	var v3388 int32
	_ = v3388
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3395 int32
	_ = v3395
	var v3398 int32
	_ = v3398
	var v3405 int32
	_ = v3405
	var v3410 int32
	_ = v3410
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3418 int32
	_ = v3418
	var v3422 int32
	_ = v3422
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3438 int32
	_ = v3438
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3475 int32
	_ = v3475
	var v3478 int32
	_ = v3478
	var v3479 int32
	_ = v3479
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3502 int32
	_ = v3502
	var v3508 int32
	_ = v3508
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3520 int32
	_ = v3520
	var v3521 int32
	_ = v3521
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3535 int32
	_ = v3535
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3546 int32
	_ = v3546
	var v3551 int32
	_ = v3551
	var v3553 int32
	_ = v3553
	var v3580 int32
	_ = v3580
	var v3582 int32
	_ = v3582
	var v3584 int32
	_ = v3584
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3591 int32
	_ = v3591
	var v3593 int32
	_ = v3593
	var v3599 int32
	_ = v3599
	var v3600 int32
	_ = v3600
	var v3603 int32
	_ = v3603
	var v3604 int32
	_ = v3604
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3610 int32
	_ = v3610
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3625 int32
	_ = v3625
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3633 int32
	_ = v3633
	var v3634 int32
	_ = v3634
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3641 int32
	_ = v3641
	var v3648 int32
	_ = v3648
	var v3649 int32
	_ = v3649
	var v3654 int32
	_ = v3654
	var v3657 int32
	_ = v3657
	var v3658 int32
	_ = v3658
	var v3662 int32
	_ = v3662
	var v3663 int32
	_ = v3663
	var v3666 int32
	_ = v3666
	var v3667 int32
	_ = v3667
	var v3670 int32
	_ = v3670
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3683 int32
	_ = v3683
	var v3686 int32
	_ = v3686
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3702 int32
	_ = v3702
	var v3709 int32
	_ = v3709
	var v3710 int32
	_ = v3710
	var v3715 int32
	_ = v3715
	var v3718 int32
	_ = v3718
	var v3719 int32
	_ = v3719
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3727 int32
	_ = v3727
	var v3728 int32
	_ = v3728
	var v3731 int32
	_ = v3731
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3744 int32
	_ = v3744
	var v3747 int32
	_ = v3747
	var v3750 int32
	_ = v3750
	var v3751 int32
	_ = v3751
	var v3755 int32
	_ = v3755
	var v3756 int32
	_ = v3756
	var v3759 int32
	_ = v3759
	var v3760 int32
	_ = v3760
	var v3763 int32
	_ = v3763
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3780 int32
	_ = v3780
	var v3784 int32
	_ = v3784
	var v3785 int32
	_ = v3785
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3792 int32
	_ = v3792
	var v3799 int32
	_ = v3799
	var v3800 int32
	_ = v3800
	var v3805 int32
	_ = v3805
	var v3808 int32
	_ = v3808
	var v3811 int32
	_ = v3811
	var v3812 int32
	_ = v3812
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3824 int32
	_ = v3824
	var v3831 int32
	_ = v3831
	var v3832 int32
	_ = v3832
	var v3837 int32
	_ = v3837
	var v3840 int32
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3845 int32
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3853 int32
	_ = v3853
	var v3860 int32
	_ = v3860
	var v3861 int32
	_ = v3861
	var v3866 int32
	_ = v3866
	var v3869 int32
	_ = v3869
	var v3872 int32
	_ = v3872
	var v3873 int32
	_ = v3873
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3881 int32
	_ = v3881
	var v3882 int32
	_ = v3882
	var v3885 int32
	_ = v3885
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3898 int32
	_ = v3898
	var v3901 int32
	_ = v3901
	var v3902 int32
	_ = v3902
	var v3906 int32
	_ = v3906
	var v3907 int32
	_ = v3907
	var v3910 int32
	_ = v3910
	var v3911 int32
	_ = v3911
	var v3914 int32
	_ = v3914
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3927 int32
	_ = v3927
	var v3930 int32
	_ = v3930
	var v3933 int32
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3938 int32
	_ = v3938
	var v3939 int32
	_ = v3939
	var v3942 int32
	_ = v3942
	var v3943 int32
	_ = v3943
	var v3946 int32
	_ = v3946
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3959 int32
	_ = v3959
	var v3962 int32
	_ = v3962
	var v3963 int32
	_ = v3963
	var v3967 int32
	_ = v3967
	var v3968 int32
	_ = v3968
	var v3971 int32
	_ = v3971
	var v3972 int32
	_ = v3972
	var v3975 int32
	_ = v3975
	var v3982 int32
	_ = v3982
	var v3983 int32
	_ = v3983
	var v3988 int32
	_ = v3988
	var v3991 int32
	_ = v3991
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3999 int32
	_ = v3999
	var v4000 int32
	_ = v4000
	var v4003 int32
	_ = v4003
	var v4004 int32
	_ = v4004
	var v4007 int32
	_ = v4007
	var v4014 int32
	_ = v4014
	var v4015 int32
	_ = v4015
	var v4026 int32
	_ = v4026
	var v4027 int32
	_ = v4027
	var v4028 int32
	_ = v4028
	var v4033 int32
	_ = v4033
	var v4034 int32
	_ = v4034
	var v4037 int32
	_ = v4037
	var v4038 int32
	_ = v4038
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4045 int32
	_ = v4045
	var v4046 int32
	_ = v4046
	var v4048 int32
	_ = v4048
	var v4050 int32
	_ = v4050
	var v4051 int32
	_ = v4051
	var v4056 int32
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4066 int32
	_ = v4066
	var v4067 int32
	_ = v4067
	var v4069 int32
	_ = v4069
	var v4074 int32
	_ = v4074
	var v4076 int32
	_ = v4076
	var v4077 int32
	_ = v4077
	var v4081 int32
	_ = v4081
	var v4083 int32
	_ = v4083
	var v4086 int32
	_ = v4086
	var v4087 int32
	_ = v4087
	var v4090 int32
	_ = v4090
	var v4093 int32
	_ = v4093
	var v4100 int32
	_ = v4100
	var v4105 int32
	_ = v4105
	var v4106 int32
	_ = v4106
	var v4108 int32
	_ = v4108
	var v4113 int32
	_ = v4113
	var v4117 int32
	_ = v4117
	var v4121 int32
	_ = v4121
	var v4125 int32
	_ = v4125
	var v4131 int32
	_ = v4131
	var v4132 int32
	_ = v4132
	var v4135 int32
	_ = v4135
	var v4136 int32
	_ = v4136
	var v4141 int32
	_ = v4141
	var v4142 int32
	_ = v4142
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4148 int32
	_ = v4148
	var v4153 int32
	_ = v4153
	var v4155 int32
	_ = v4155
	var v4156 int32
	_ = v4156
	var v4160 int32
	_ = v4160
	var v4162 int32
	_ = v4162
	var v4165 int32
	_ = v4165
	var v4166 int32
	_ = v4166
	var v4169 int32
	_ = v4169
	var v4172 int32
	_ = v4172
	var v4179 int32
	_ = v4179
	var v4185 int32
	_ = v4185
	var v4187 int32
	_ = v4187
	var v4191 int32
	_ = v4191
	var v4192 int32
	_ = v4192
	var v4193 int32
	_ = v4193
	var v4197 int32
	_ = v4197
	var v4201 int32
	_ = v4201
	var v4204 int32
	_ = v4204
	var v4207 int32
	_ = v4207
	var v4210 int32
	_ = v4210
	var v4212 int32
	_ = v4212
	var v4213 int32
	_ = v4213
	var v4216 int32
	_ = v4216
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
	var v4233 int32
	_ = v4233
	var v4235 int32
	_ = v4235
	var v4236 int32
	_ = v4236
	var v4240 int32
	_ = v4240
	var v4242 int32
	_ = v4242
	var v4245 int32
	_ = v4245
	var v4246 int32
	_ = v4246
	var v4249 int32
	_ = v4249
	var v4252 int32
	_ = v4252
	var v4259 int32
	_ = v4259
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4267 int32
	_ = v4267
	var v4269 int32
	_ = v4269
	var v4279 int32
	_ = v4279
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4285 int32
	_ = v4285
	var v4290 int32
	_ = v4290
	var v4298 int32
	_ = v4298
	var v4299 int32
	_ = v4299
	var v4300 int32
	_ = v4300
	var v4303 int32
	_ = v4303
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4311 int32
	_ = v4311
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4317 int32
	_ = v4317
	var v4324 int32
	_ = v4324
	var v4344 int32
	_ = v4344
	var v4348 int32
	_ = v4348
	var v4349 int32
	_ = v4349
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4365 int32
	_ = v4365
	var v4366 int32
	_ = v4366
	var v4368 int32
	_ = v4368
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4371 int32
	_ = v4371
	var v4374 int32
	_ = v4374
	var v4376 int32
	_ = v4376
	var v4378 int32
	_ = v4378
	var v4380 int32
	_ = v4380
	var v4381 int32
	_ = v4381
	var v4413 int32
	_ = v4413
	var v4414 int32
	_ = v4414
	var v4422 int32
	_ = v4422
	var v4424 int32
	_ = v4424
	var v4428 int32
	_ = v4428
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4438 int32
	_ = v4438
	var v4445 int32
	_ = v4445
	var v4449 int32
	_ = v4449
	var v4450 int32
	_ = v4450
	var v4451 int32
	_ = v4451
	var v4452 int32
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4460 int32
	_ = v4460
	var v4461 int32
	_ = v4461
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4468 int32
	_ = v4468
	var v4469 int32
	_ = v4469
	var v4474 int32
	_ = v4474
	var v4475 int32
	_ = v4475
	var v4478 int32
	_ = v4478
	var v4479 int32
	_ = v4479
	var v4481 int32
	_ = v4481
	var v4486 int32
	_ = v4486
	var v4488 int32
	_ = v4488
	var v4489 int32
	_ = v4489
	var v4493 int32
	_ = v4493
	var v4495 int32
	_ = v4495
	var v4498 int32
	_ = v4498
	var v4499 int32
	_ = v4499
	var v4502 int32
	_ = v4502
	var v4505 int32
	_ = v4505
	var v4512 int32
	_ = v4512
	var v4518 int32
	_ = v4518
	var v4520 int32
	_ = v4520
	var v4525 int32
	_ = v4525
	var v4527 int32
	_ = v4527
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4538 int32
	_ = v4538
	var v4539 int32
	_ = v4539
	var v4540 int32
	_ = v4540
	var v4544 int32
	_ = v4544
	var v4545 int32
	_ = v4545
	var v4550 int32
	_ = v4550
	var v4551 int32
	_ = v4551
	var v4554 int32
	_ = v4554
	var v4555 int32
	_ = v4555
	var v4557 int32
	_ = v4557
	var v4562 int32
	_ = v4562
	var v4564 int32
	_ = v4564
	var v4565 int32
	_ = v4565
	var v4569 int32
	_ = v4569
	var v4571 int32
	_ = v4571
	var v4574 int32
	_ = v4574
	var v4575 int32
	_ = v4575
	var v4578 int32
	_ = v4578
	var v4581 int32
	_ = v4581
	var v4588 int32
	_ = v4588
	var v4593 int32
	_ = v4593
	var v4594 int32
	_ = v4594
	var v4596 int32
	_ = v4596
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4603 int32
	_ = v4603
	var v4605 int32
	_ = v4605
	var v4606 int32
	_ = v4606
	var v4607 int32
	_ = v4607
	var v4609 int32
	_ = v4609
	var v4612 int32
	_ = v4612
	var v4613 int32
	_ = v4613
	var v4614 int32
	_ = v4614
	var v4615 int32
	_ = v4615
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4619 int32
	_ = v4619
	var v4623 int32
	_ = v4623
	var v4624 int32
	_ = v4624
	var v4629 int32
	_ = v4629
	var v4630 int32
	_ = v4630
	var v4635 int32
	_ = v4635
	var v4636 int32
	_ = v4636
	var v4639 int32
	_ = v4639
	var v4640 int32
	_ = v4640
	var v4642 int32
	_ = v4642
	var v4647 int32
	_ = v4647
	var v4649 int32
	_ = v4649
	var v4650 int32
	_ = v4650
	var v4654 int32
	_ = v4654
	var v4656 int32
	_ = v4656
	var v4659 int32
	_ = v4659
	var v4660 int32
	_ = v4660
	var v4663 int32
	_ = v4663
	var v4666 int32
	_ = v4666
	var v4673 int32
	_ = v4673
	var v4678 int32
	_ = v4678
	var v4679 int32
	_ = v4679
	var v4681 int32
	_ = v4681
	var v4685 int32
	_ = v4685
	var v4686 int32
	_ = v4686
	var v4690 int32
	_ = v4690
	var v4692 int32
	_ = v4692
	var v4694 int32
	_ = v4694
	var v4695 int32
	_ = v4695
	var v4696 int32
	_ = v4696
	var v4698 int32
	_ = v4698
	var v4701 int32
	_ = v4701
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4704 int32
	_ = v4704
	var v4705 int32
	_ = v4705
	var v4706 int32
	_ = v4706
	var v4708 int32
	_ = v4708
	var v4713 int32
	_ = v4713
	var v4714 int32
	_ = v4714
	var v4717 int32
	_ = v4717
	var v4718 int32
	_ = v4718
	var v4723 int32
	_ = v4723
	var v4724 int32
	_ = v4724
	var v4727 int32
	_ = v4727
	var v4728 int32
	_ = v4728
	var v4730 int32
	_ = v4730
	var v4735 int32
	_ = v4735
	var v4737 int32
	_ = v4737
	var v4738 int32
	_ = v4738
	var v4742 int32
	_ = v4742
	var v4744 int32
	_ = v4744
	var v4747 int32
	_ = v4747
	var v4748 int32
	_ = v4748
	var v4751 int32
	_ = v4751
	var v4754 int32
	_ = v4754
	var v4761 int32
	_ = v4761
	var v4766 int32
	_ = v4766
	var v4767 int32
	_ = v4767
	var v4773 int32
	_ = v4773
	var v4775 int32
	_ = v4775
	var v4778 int32
	_ = v4778
	var v4779 int32
	_ = v4779
	var v4780 int32
	_ = v4780
	var v4782 int32
	_ = v4782
	var v4785 int32
	_ = v4785
	var v4786 int32
	_ = v4786
	var v4787 int32
	_ = v4787
	var v4788 int32
	_ = v4788
	var v4789 int32
	_ = v4789
	var v4790 int32
	_ = v4790
	var v4792 int32
	_ = v4792
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4801 int32
	_ = v4801
	var v4808 int32
	_ = v4808
	var v4811 int32
	_ = v4811
	var v4820 int32
	_ = v4820
	var v4821 int32
	_ = v4821
	var v4823 int32
	_ = v4823
	var v4824 int32
	_ = v4824
	var v4828 int32
	_ = v4828
	var v4829 int32
	_ = v4829
	var v4831 int32
	_ = v4831
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4838 int32
	_ = v4838
	var v4841 int32
	_ = v4841
	var v4843 int32
	_ = v4843
	var v4846 int32
	_ = v4846
	var v4851 int32
	_ = v4851
	var v4854 int32
	_ = v4854
	var v4857 int32
	_ = v4857
	var v4859 int32
	_ = v4859
	var v4861 int32
	_ = v4861
	var v4862 int32
	_ = v4862
	var v4871 int32
	_ = v4871
	var v4872 int32
	_ = v4872
	var v4875 int32
	_ = v4875
	var v4876 int32
	_ = v4876
	var v4878 int32
	_ = v4878
	var v4880 int32
	_ = v4880
	var v4883 int32
	_ = v4883
	var v4890 int32
	_ = v4890
	var v4919 int32
	_ = v4919
	var v4922 int32
	_ = v4922
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4933 int32
	_ = v4933
	var v4934 int32
	_ = v4934
	var v4935 int32
	_ = v4935
	var v4937 int32
	_ = v4937
	var v4965 int32
	_ = v4965
	var v4966 int32
	_ = v4966
	var v4967 int32
	_ = v4967
	var v4968 int32
	_ = v4968
	var v4972 int32
	_ = v4972
	var v4973 int32
	_ = v4973
	var v4974 int32
	_ = v4974
	var v4977 int32
	_ = v4977
	var v4978 int32
	_ = v4978
	var v4980 int32
	_ = v4980
	var v4983 int32
	_ = v4983
	var v4987 int32
	_ = v4987
	var v4990 int32
	_ = v4990
	var v4991 int32
	_ = v4991
	var v4999 int32
	_ = v4999
	var v5000 int32
	_ = v5000
	var v5004 int32
	_ = v5004
	var v5007 int32
	_ = v5007
	var v5011 int32
	_ = v5011
	var v5012 int32
	_ = v5012
	var v5015 int32
	_ = v5015
	var v5016 int32
	_ = v5016
	var v5019 int32
	_ = v5019
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5027 int32
	_ = v5027
	var v5028 int32
	_ = v5028
	var v5031 int32
	_ = v5031
	var v5032 int32
	_ = v5032
	var v5035 int32
	_ = v5035
	var v5042 int32
	_ = v5042
	var v5043 int32
	_ = v5043
	var v5053 int32
	_ = v5053
	var v5054 int32
	_ = v5054
	var v5070 int32
	_ = v5070
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5076 int32
	_ = v5076
	var v5079 int32
	_ = v5079
	var v5080 int32
	_ = v5080
	var v5081 int32
	_ = v5081
	var v5082 int32
	_ = v5082
	var v5083 int32
	_ = v5083
	var v5086 int32
	_ = v5086
	var v5092 int32
	_ = v5092
	var v5093 int32
	_ = v5093
	var v5102 int32
	_ = v5102
	var v5103 int32
	_ = v5103
	var v5107 int32
	_ = v5107
	var v5115 int32
	_ = v5115
	var v5116 int32
	_ = v5116
	var v5120 int32
	_ = v5120
	var v5121 int32
	_ = v5121
	var v5123 int32
	_ = v5123
	var v5127 int32
	_ = v5127
	var v5128 int32
	_ = v5128
	var v5130 int32
	_ = v5130
	var v5133 int32
	_ = v5133
	var v5135 int32
	_ = v5135
	var v5138 int32
	_ = v5138
	var v5143 int32
	_ = v5143
	var v5146 int32
	_ = v5146
	var v5149 int32
	_ = v5149
	var v5151 int32
	_ = v5151
	var v5153 int32
	_ = v5153
	var v5154 int32
	_ = v5154
	var v5163 int32
	_ = v5163
	var v5164 int32
	_ = v5164
	var v5167 int32
	_ = v5167
	var v5168 int32
	_ = v5168
	var v5170 int32
	_ = v5170
	var v5172 int32
	_ = v5172
	var v5175 int32
	_ = v5175
	var v5183 int32
	_ = v5183
	var v5184 int32
	_ = v5184
	var v5185 int32
	_ = v5185
	var v5186 int32
	_ = v5186
	var v5187 int32
	_ = v5187
	var v5189 int32
	_ = v5189
	var v5190 int32
	_ = v5190
	var v5195 int32
	_ = v5195
	var v5196 int32
	_ = v5196
	var v5199 int32
	_ = v5199
	var v5200 int32
	_ = v5200
	var v5202 int32
	_ = v5202
	var v5207 int32
	_ = v5207
	var v5209 int32
	_ = v5209
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
	var v5223 int32
	_ = v5223
	var v5226 int32
	_ = v5226
	var v5233 int32
	_ = v5233
	var v5237 int32
	_ = v5237
	var v5238 int32
	_ = v5238
	var v5239 int32
	_ = v5239
	var v5245 int32
	_ = v5245
	var v5254 int32
	_ = v5254
	var v5255 int32
	_ = v5255
	var v5261 int32
	_ = v5261
	var v5264 int32
	_ = v5264
	var v5265 int32
	_ = v5265
	var v5266 int32
	_ = v5266
	var v5271 int32
	_ = v5271
	var v5272 int32
	_ = v5272
	var v5274 int32
	_ = v5274
	var v5276 int32
	_ = v5276
	var v5277 int32
	_ = v5277
	var v5282 int32
	_ = v5282
	var v5283 int32
	_ = v5283
	var v5286 int32
	_ = v5286
	var v5287 int32
	_ = v5287
	var v5289 int32
	_ = v5289
	var v5294 int32
	_ = v5294
	var v5296 int32
	_ = v5296
	var v5297 int32
	_ = v5297
	var v5301 int32
	_ = v5301
	var v5303 int32
	_ = v5303
	var v5306 int32
	_ = v5306
	var v5307 int32
	_ = v5307
	var v5310 int32
	_ = v5310
	var v5313 int32
	_ = v5313
	var v5320 int32
	_ = v5320
	var v5330 int32
	_ = v5330
	var v5331 int32
	_ = v5331
	var v5337 int32
	_ = v5337
	var v5340 int32
	_ = v5340
	var v5342 int32
	_ = v5342
	var v5343 int32
	_ = v5343
	var v5345 int32
	_ = v5345
	var v5347 int32
	_ = v5347
	var v5348 int32
	_ = v5348
	var v5353 int32
	_ = v5353
	var v5354 int32
	_ = v5354
	var v5359 int32
	_ = v5359
	var v5360 int32
	_ = v5360
	var v5363 int32
	_ = v5363
	var v5364 int32
	_ = v5364
	var v5366 int32
	_ = v5366
	var v5371 int32
	_ = v5371
	var v5373 int32
	_ = v5373
	var v5374 int32
	_ = v5374
	var v5378 int32
	_ = v5378
	var v5380 int32
	_ = v5380
	var v5383 int32
	_ = v5383
	var v5384 int32
	_ = v5384
	var v5387 int32
	_ = v5387
	var v5390 int32
	_ = v5390
	var v5397 int32
	_ = v5397
	var v5402 int32
	_ = v5402
	var v5403 int32
	_ = v5403
	var v5405 int32
	_ = v5405
	var v5409 int32
	_ = v5409
	var v5410 int32
	_ = v5410
	var v5414 int32
	_ = v5414
	var v5418 int32
	_ = v5418
	var v5420 int32
	_ = v5420
	var v5423 int32
	_ = v5423
	var v5424 int32
	_ = v5424
	var v5428 int32
	_ = v5428
	var v5429 int32
	_ = v5429
	var v5432 int32
	_ = v5432
	var v5433 int32
	_ = v5433
	var v5434 int32
	_ = v5434
	var v5436 int32
	_ = v5436
	var v5439 int32
	_ = v5439
	var v5441 int32
	_ = v5441
	var v5442 int32
	_ = v5442
	var v5443 int32
	_ = v5443
	var v5444 int32
	_ = v5444
	var v5446 int32
	_ = v5446
	var v5449 int32
	_ = v5449
	var v5450 int32
	_ = v5450
	var v5451 int32
	_ = v5451
	var v5452 int32
	_ = v5452
	var v5453 int32
	_ = v5453
	var v5454 int32
	_ = v5454
	var v5456 int32
	_ = v5456
	var v5461 int32
	_ = v5461
	var v5464 int32
	_ = v5464
	var v5465 int32
	_ = v5465
	var v5469 int32
	_ = v5469
	var v5470 int32
	_ = v5470
	var v5472 int32
	_ = v5472
	var v5477 int32
	_ = v5477
	var v5480 int32
	_ = v5480
	var v5481 int32
	_ = v5481
	var v5482 int32
	_ = v5482
	var v5487 int32
	_ = v5487
	var v5488 int32
	_ = v5488
	var v5491 int32
	_ = v5491
	var v5492 int32
	_ = v5492
	var v5494 int32
	_ = v5494
	var v5499 int32
	_ = v5499
	var v5501 int32
	_ = v5501
	var v5502 int32
	_ = v5502
	var v5506 int32
	_ = v5506
	var v5508 int32
	_ = v5508
	var v5511 int32
	_ = v5511
	var v5512 int32
	_ = v5512
	var v5515 int32
	_ = v5515
	var v5518 int32
	_ = v5518
	var v5525 int32
	_ = v5525
	var v5530 int32
	_ = v5530
	var v5531 int32
	_ = v5531
	var v5533 int32
	_ = v5533
	var v5536 int32
	_ = v5536
	var v5537 int32
	_ = v5537
	var v5538 int32
	_ = v5538
	var v5539 int32
	_ = v5539
	var v5541 int32
	_ = v5541
	var v5544 int32
	_ = v5544
	var v5545 int32
	_ = v5545
	var v5551 int32
	_ = v5551
	var v5554 int32
	_ = v5554
	var v5557 int32
	_ = v5557
	var v5561 int32
	_ = v5561
	var v5564 int32
	_ = v5564
	var v5565 int32
	_ = v5565
	var v5571 int32
	_ = v5571
	var v5574 int32
	_ = v5574
	var v5575 int32
	_ = v5575
	var v5576 int32
	_ = v5576
	var v5581 int32
	_ = v5581
	var v5582 int32
	_ = v5582
	var v5583 int32
	_ = v5583
	var v5584 int32
	_ = v5584
	var v5587 int32
	_ = v5587
	var v5591 int32
	_ = v5591
	var v5596 int32
	_ = v5596
	var v5598 int32
	_ = v5598
	var v5604 int32
	_ = v5604
	var v5605 int32
	_ = v5605
	var v5615 int32
	_ = v5615
	var v5616 int32
	_ = v5616
	var v5617 int32
	_ = v5617
	var v5620 int32
	_ = v5620
	var v5623 int32
	_ = v5623
	var v5624 int32
	_ = v5624
	var v5628 int32
	_ = v5628
	var v5629 int32
	_ = v5629
	var v5632 int32
	_ = v5632
	var v5633 int32
	_ = v5633
	var v5636 int32
	_ = v5636
	var v5643 int32
	_ = v5643
	var v5644 int32
	_ = v5644
	var v5647 int32
	_ = v5647
	var v5649 int32
	_ = v5649
	var v5650 int32
	_ = v5650
	var v5654 int32
	_ = v5654
	var v5655 int32
	_ = v5655
	var v5658 int32
	_ = v5658
	var v5663 int32
	_ = v5663
	var v5664 int32
	_ = v5664
	var v5667 int32
	_ = v5667
	var v5668 int32
	_ = v5668
	var v5670 int32
	_ = v5670
	var v5675 int32
	_ = v5675
	var v5677 int32
	_ = v5677
	var v5678 int32
	_ = v5678
	var v5682 int32
	_ = v5682
	var v5684 int32
	_ = v5684
	var v5687 int32
	_ = v5687
	var v5688 int32
	_ = v5688
	var v5691 int32
	_ = v5691
	var v5694 int32
	_ = v5694
	var v5701 int32
	_ = v5701
	var v5706 int32
	_ = v5706
	var v5707 int32
	_ = v5707
	var v5709 int32
	_ = v5709
	var v5714 int32
	_ = v5714
	var v5719 int32
	_ = v5719
	var v5720 int32
	_ = v5720
	var v5726 int32
	_ = v5726
	var v5727 int32
	_ = v5727
	var v5732 int32
	_ = v5732
	var v5733 int32
	_ = v5733
	var v5736 int32
	_ = v5736
	var v5737 int32
	_ = v5737
	var v5742 int32
	_ = v5742
	var v5748 int32
	_ = v5748
	var v5749 int32
	_ = v5749
	var v5758 int32
	_ = v5758
	var v5760 int32
	_ = v5760
	var v5764 int32
	_ = v5764
	var v5772 int32
	_ = v5772
	var v5773 int32
	_ = v5773
	var v5776 int32
	_ = v5776
	var v5779 int32
	_ = v5779
	var v5780 int32
	_ = v5780
	var v5784 int32
	_ = v5784
	var v5785 int32
	_ = v5785
	var v5788 int32
	_ = v5788
	var v5789 int32
	_ = v5789
	var v5792 int32
	_ = v5792
	var v5799 int32
	_ = v5799
	var v5800 int32
	_ = v5800
	var v5803 int32
	_ = v5803
	var v5805 int32
	_ = v5805
	var v5806 int32
	_ = v5806
	var v5810 int32
	_ = v5810
	var v5811 int32
	_ = v5811
	var v5814 int32
	_ = v5814
	var v5819 int32
	_ = v5819
	var v5820 int32
	_ = v5820
	var v5823 int32
	_ = v5823
	var v5824 int32
	_ = v5824
	var v5826 int32
	_ = v5826
	var v5831 int32
	_ = v5831
	var v5833 int32
	_ = v5833
	var v5834 int32
	_ = v5834
	var v5838 int32
	_ = v5838
	var v5840 int32
	_ = v5840
	var v5843 int32
	_ = v5843
	var v5844 int32
	_ = v5844
	var v5847 int32
	_ = v5847
	var v5850 int32
	_ = v5850
	var v5857 int32
	_ = v5857
	var v5862 int32
	_ = v5862
	var v5863 int32
	_ = v5863
	var v5865 int32
	_ = v5865
	var v5872 int32
	_ = v5872
	var v5873 int32
	_ = v5873
	var v5881 int32
	_ = v5881
	var v5883 int32
	_ = v5883
	var v5895 int32
	_ = v5895
	var v5896 int32
	_ = v5896
	var v5901 int32
	_ = v5901
	var v5904 int32
	_ = v5904
	var v5913 int32
	_ = v5913
	var v5914 int32
	_ = v5914
	var v5916 int32
	_ = v5916
	var v5945 int32
	_ = v5945
	var v5948 int32
	_ = v5948
	var v5957 int32
	_ = v5957
	var v5958 int32
	_ = v5958
	var v5959 int32
	_ = v5959
	var v5960 int32
	_ = v5960
	var v5961 int32
	_ = v5961
	var v5963 int32
	_ = v5963
	var v5997 int32
	_ = v5997
	var v5998 int32
	_ = v5998
	var v6000 int32
	_ = v6000
	var v6001 int32
	_ = v6001
	var v6004 int32
	_ = v6004
	var v6009 int32
	_ = v6009
	var v6010 int32
	_ = v6010
	var v6013 int32
	_ = v6013
	var v6014 int32
	_ = v6014
	var v6016 int32
	_ = v6016
	var v6021 int32
	_ = v6021
	var v6023 int32
	_ = v6023
	var v6024 int32
	_ = v6024
	var v6028 int32
	_ = v6028
	var v6030 int32
	_ = v6030
	var v6033 int32
	_ = v6033
	var v6034 int32
	_ = v6034
	var v6037 int32
	_ = v6037
	var v6040 int32
	_ = v6040
	var v6047 int32
	_ = v6047
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
	var v6067 int32
	_ = v6067
	var v6068 int32
	_ = v6068
	var v6074 int32
	_ = v6074
	var v6077 int32
	_ = v6077
	var v6081 int32
	_ = v6081
	var v6085 int32
	_ = v6085
	var v6086 int32
	_ = v6086
	var v6087 int32
	_ = v6087
	var v6088 int32
	_ = v6088
	var v6093 int32
	_ = v6093
	var v6094 int32
	_ = v6094
	var v6101 int32
	_ = v6101
	var v6102 int32
	_ = v6102
	var v6106 int32
	_ = v6106
	var v6107 int32
	_ = v6107
	var v6111 int32
	_ = v6111
	var v6116 int32
	_ = v6116
	var v6120 int32
	_ = v6120
	var v6121 int32
	_ = v6121
	var v6122 int32
	_ = v6122
	var v6123 int32
	_ = v6123
	var v6128 int32
	_ = v6128
	var v6129 int32
	_ = v6129
	var v6134 int32
	_ = v6134
	var v6135 int32
	_ = v6135
	var v6141 int32
	_ = v6141
	var v6142 int32
	_ = v6142
	var v6147 int32
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6151 int32
	_ = v6151
	var v6152 int32
	_ = v6152
	var v6157 int32
	_ = v6157
	var v6163 int32
	_ = v6163
	var v6164 int32
	_ = v6164
	var v6173 int32
	_ = v6173
	var v6175 int32
	_ = v6175
	var v6179 int32
	_ = v6179
	var v6187 int32
	_ = v6187
	var v6188 int32
	_ = v6188
	var v6195 int32
	_ = v6195
	var v6196 int32
	_ = v6196
	var v6199 int32
	_ = v6199
	var v6200 int32
	_ = v6200
	var v6205 int32
	_ = v6205
	var v6206 int32
	_ = v6206
	var v6209 int32
	_ = v6209
	var v6210 int32
	_ = v6210
	var v6212 int32
	_ = v6212
	var v6217 int32
	_ = v6217
	var v6219 int32
	_ = v6219
	var v6220 int32
	_ = v6220
	var v6224 int32
	_ = v6224
	var v6226 int32
	_ = v6226
	var v6229 int32
	_ = v6229
	var v6230 int32
	_ = v6230
	var v6233 int32
	_ = v6233
	var v6236 int32
	_ = v6236
	var v6243 int32
	_ = v6243
	var v6248 int32
	_ = v6248
	var v6249 int32
	_ = v6249
	var v6251 int32
	_ = v6251
	var v6253 int64
	_ = v6253
	var v6255 int32
	_ = v6255
	var v6265 int32
	_ = v6265
	var v6266 int32
	_ = v6266
	var v6269 int32
	_ = v6269
	var v6281 int32
	_ = v6281
	var v6286 int32
	_ = v6286
	var v6287 int32
	_ = v6287
	var v6290 int32
	_ = v6290
	var v6293 int32
	_ = v6293
	var v6294 int32
	_ = v6294
	var v6298 int32
	_ = v6298
	var v6299 int32
	_ = v6299
	var v6302 int32
	_ = v6302
	var v6303 int32
	_ = v6303
	var v6306 int32
	_ = v6306
	var v6313 int32
	_ = v6313
	var v6314 int32
	_ = v6314
	var v6318 int32
	_ = v6318
	var v6321 int32
	_ = v6321
	var v6322 int32
	_ = v6322
	var v6326 int32
	_ = v6326
	var v6327 int32
	_ = v6327
	var v6330 int32
	_ = v6330
	var v6331 int32
	_ = v6331
	var v6334 int32
	_ = v6334
	var v6341 int32
	_ = v6341
	var v6342 int32
	_ = v6342
	var v6346 int32
	_ = v6346
	var v6349 int32
	_ = v6349
	var v6352 int32
	_ = v6352
	var v6353 int32
	_ = v6353
	var v6357 int32
	_ = v6357
	var v6358 int32
	_ = v6358
	var v6361 int32
	_ = v6361
	var v6362 int32
	_ = v6362
	var v6365 int32
	_ = v6365
	var v6372 int32
	_ = v6372
	var v6373 int32
	_ = v6373
	var v6377 int32
	_ = v6377
	var v6380 int32
	_ = v6380
	var v6381 int32
	_ = v6381
	var v6385 int32
	_ = v6385
	var v6386 int32
	_ = v6386
	var v6389 int32
	_ = v6389
	var v6390 int32
	_ = v6390
	var v6393 int32
	_ = v6393
	var v6400 int32
	_ = v6400
	var v6401 int32
	_ = v6401
	var v6405 int32
	_ = v6405
	var v6408 int32
	_ = v6408
	var v6411 int32
	_ = v6411
	var v6412 int32
	_ = v6412
	var v6416 int32
	_ = v6416
	var v6417 int32
	_ = v6417
	var v6420 int32
	_ = v6420
	var v6421 int32
	_ = v6421
	var v6424 int32
	_ = v6424
	var v6431 int32
	_ = v6431
	var v6432 int32
	_ = v6432
	var v6436 int32
	_ = v6436
	var v6439 int32
	_ = v6439
	var v6440 int32
	_ = v6440
	var v6444 int32
	_ = v6444
	var v6445 int32
	_ = v6445
	var v6448 int32
	_ = v6448
	var v6449 int32
	_ = v6449
	var v6452 int32
	_ = v6452
	var v6459 int32
	_ = v6459
	var v6460 int32
	_ = v6460
	var v6465 int32
	_ = v6465
	var v6471 int32
	_ = v6471
	var v6472 int32
	_ = v6472
	var v6474 int32
	_ = v6474
	var v6485 int32
	_ = v6485
	var v6486 int32
	_ = v6486
	var v6492 int32
	_ = v6492
	var v6493 int32
	_ = v6493
	var v6504 int32
	_ = v6504
	var v6507 int32
	_ = v6507
	var v6508 int32
	_ = v6508
	var v6509 int32
	_ = v6509
	var v6512 int32
	_ = v6512
	var v6515 int32
	_ = v6515
	var v6516 int32
	_ = v6516
	var v6520 int32
	_ = v6520
	var v6521 int32
	_ = v6521
	var v6524 int32
	_ = v6524
	var v6525 int32
	_ = v6525
	var v6528 int32
	_ = v6528
	var v6535 int32
	_ = v6535
	var v6536 int32
	_ = v6536
	var v6544 int32
	_ = v6544
	var v6545 int32
	_ = v6545
	var v6548 int32
	_ = v6548
	var v6549 int32
	_ = v6549
	var v6552 int32
	_ = v6552
	var v6556 int32
	_ = v6556
	var v6558 int32
	_ = v6558
	var v6559 int64
	_ = v6559
	var v6567 int32
	_ = v6567
	var v6571 int32
	_ = v6571
	var v6575 int32
	_ = v6575
	var v6581 int32
	_ = v6581
	var v6585 int32
	_ = v6585
	var v6586 int32
	_ = v6586
	var v6593 int32
	_ = v6593
	var v6594 int32
	_ = v6594
	var v6595 int32
	_ = v6595
	var v6599 int32
	_ = v6599
	var v6602 int32
	_ = v6602
	var v6606 int32
	_ = v6606
	var v6607 int32
	_ = v6607
	var v6615 int32
	_ = v6615
	var v6621 int32
	_ = v6621
	var v6623 int32
	_ = v6623
	var v6627 int32
	_ = v6627
	var v6635 int32
	_ = v6635
	var v6639 int32
	_ = v6639
	var v6644 int32
	_ = v6644
	var v6650 int32
	_ = v6650
	var v6651 int32
	_ = v6651
	var v6655 int32
	_ = v6655
	var v6659 int32
	_ = v6659
	var v6660 int32
	_ = v6660
	var v6661 int32
	_ = v6661
	var v6667 int32
	_ = v6667
	var v6668 int32
	_ = v6668
	var v6671 int32
	_ = v6671
	var v6674 int32
	_ = v6674
	var v6675 int32
	_ = v6675
	var v6682 int32
	_ = v6682
	var v6683 int32
	_ = v6683
	var v6694 int32
	_ = v6694
	var v6724 int32
	_ = v6724
	var v6733 int32
	_ = v6733
	var v6734 int32
	_ = v6734
	var v6735 int32
	_ = v6735
	var v6736 int32
	_ = v6736
	var v6737 int32
	_ = v6737
	var v6739 int32
	_ = v6739
	var v6745 int32
	_ = v6745
	var v6799 int32
	_ = v6799
	var v6821 int32
	_ = v6821
	var v6822 int32
	_ = v6822
	var v6827 int32
	_ = v6827
	var v6828 int32
	_ = v6828
	var v6829 int32
	_ = v6829
	var v6830 int32
	_ = v6830
	var v6833 int32
	_ = v6833
	var v6834 int32
	_ = v6834
	var v6837 int32
	_ = v6837
	var v6840 int32
	_ = v6840
	var v6841 int32
	_ = v6841
	var v6845 int32
	_ = v6845
	var v6846 int32
	_ = v6846
	var v6849 int32
	_ = v6849
	var v6850 int32
	_ = v6850
	var v6853 int32
	_ = v6853
	var v6860 int32
	_ = v6860
	var v6861 int32
	_ = v6861
	var v6865 int32
	_ = v6865
	var v6868 int32
	_ = v6868
	var v6869 int32
	_ = v6869
	var v6873 int32
	_ = v6873
	var v6874 int32
	_ = v6874
	var v6877 int32
	_ = v6877
	var v6878 int32
	_ = v6878
	var v6881 int32
	_ = v6881
	var v6888 int32
	_ = v6888
	var v6889 int32
	_ = v6889
	var v6893 int32
	_ = v6893
	var v6896 int32
	_ = v6896
	var v6897 int32
	_ = v6897
	var v6901 int32
	_ = v6901
	var v6902 int32
	_ = v6902
	var v6905 int32
	_ = v6905
	var v6906 int32
	_ = v6906
	var v6909 int32
	_ = v6909
	var v6916 int32
	_ = v6916
	var v6917 int32
	_ = v6917
	var v6921 int32
	_ = v6921
	var v6924 int32
	_ = v6924
	var v6925 int32
	_ = v6925
	var v6929 int32
	_ = v6929
	var v6930 int32
	_ = v6930
	var v6933 int32
	_ = v6933
	var v6934 int32
	_ = v6934
	var v6937 int32
	_ = v6937
	var v6944 int32
	_ = v6944
	var v6945 int32
	_ = v6945
	var v6949 int32
	_ = v6949
	var v6952 int32
	_ = v6952
	var v6953 int32
	_ = v6953
	var v6957 int32
	_ = v6957
	var v6958 int32
	_ = v6958
	var v6961 int32
	_ = v6961
	var v6962 int32
	_ = v6962
	var v6965 int32
	_ = v6965
	var v6972 int32
	_ = v6972
	var v6973 int32
	_ = v6973
	var v6977 int32
	_ = v6977
	var v6980 int32
	_ = v6980
	var v6981 int32
	_ = v6981
	var v6985 int32
	_ = v6985
	var v6986 int32
	_ = v6986
	var v6989 int32
	_ = v6989
	var v6990 int32
	_ = v6990
	var v6993 int32
	_ = v6993
	var v7000 int32
	_ = v7000
	var v7001 int32
	_ = v7001
	var v7005 int32
	_ = v7005
	var v7008 int32
	_ = v7008
	var v7009 int32
	_ = v7009
	var v7013 int32
	_ = v7013
	var v7014 int32
	_ = v7014
	var v7017 int32
	_ = v7017
	var v7018 int32
	_ = v7018
	var v7021 int32
	_ = v7021
	var v7028 int32
	_ = v7028
	var v7029 int32
	_ = v7029
	var v7033 int32
	_ = v7033
	var v7036 int32
	_ = v7036
	var v7037 int32
	_ = v7037
	var v7041 int32
	_ = v7041
	var v7042 int32
	_ = v7042
	var v7045 int32
	_ = v7045
	var v7046 int32
	_ = v7046
	var v7049 int32
	_ = v7049
	var v7056 int32
	_ = v7056
	var v7057 int32
	_ = v7057
	var v7061 int32
	_ = v7061
	var v7064 int32
	_ = v7064
	var v7065 int32
	_ = v7065
	var v7069 int32
	_ = v7069
	var v7070 int32
	_ = v7070
	var v7073 int32
	_ = v7073
	var v7074 int32
	_ = v7074
	var v7077 int32
	_ = v7077
	var v7084 int32
	_ = v7084
	var v7085 int32
	_ = v7085
	var v7089 int32
	_ = v7089
	var v7096 int32
	_ = v7096
	var v7097 int32
	_ = v7097
	var v7106 int32
	_ = v7106
	var v7109 int32
	_ = v7109
	var v7118 int32
	_ = v7118
	var v7119 int32
	_ = v7119
	var v7121 int32
	_ = v7121
	var v7122 int32
	_ = v7122
	var v7123 int32
	_ = v7123
	var v7152 int32
	_ = v7152
	var v7156 int32
	_ = v7156
	var v7159 int32
	_ = v7159
	var v7177 int32
	_ = v7177
	var v7180 int32
	_ = v7180
	var v7181 int32
	_ = v7181
	var v7186 int32
	_ = v7186
	var v7188 int32
	_ = v7188
	var v7189 int32
	_ = v7189
	var v7191 int32
	_ = v7191
	var v7193 int32
	_ = v7193
	var v7196 int32
	_ = v7196
	var v7198 int32
	_ = v7198
	var v7226 int32
	_ = v7226
	var v7227 int32
	_ = v7227
	var v7230 int32
	_ = v7230
	var v7231 int32
	_ = v7231
	var v7236 int32
	_ = v7236
	var v7237 int32
	_ = v7237
	var v7240 int32
	_ = v7240
	var v7241 int32
	_ = v7241
	var v7243 int32
	_ = v7243
	var v7248 int32
	_ = v7248
	var v7250 int32
	_ = v7250
	var v7251 int32
	_ = v7251
	var v7255 int32
	_ = v7255
	var v7257 int32
	_ = v7257
	var v7260 int32
	_ = v7260
	var v7261 int32
	_ = v7261
	var v7264 int32
	_ = v7264
	var v7267 int32
	_ = v7267
	var v7274 int32
	_ = v7274
	var v7279 int32
	_ = v7279
	var v7280 int32
	_ = v7280
	var v7281 int32
	_ = v7281
	var v7282 int32
	_ = v7282
	var v7287 int32
	_ = v7287
	var v7299 int32
	_ = v7299
	var v7300 int32
	_ = v7300
	var v7302 int32
	_ = v7302
	var v7306 int32
	_ = v7306
	var v7310 int32
	_ = v7310
	var v7318 int32
	_ = v7318
	var v7319 int32
	_ = v7319
	var v7321 int32
	_ = v7321
	var v7326 int32
	_ = v7326
	var v7330 int32
	_ = v7330
	var v7334 int32
	_ = v7334
	var v7337 int32
	_ = v7337
	var v7343 int32
	_ = v7343
	var v7344 int32
	_ = v7344
	var v7347 int32
	_ = v7347
	var v7353 int32
	_ = v7353
	var v7354 int32
	_ = v7354
	var v7357 int32
	_ = v7357
	var v7363 int32
	_ = v7363
	var v7364 int32
	_ = v7364
	var v7370 int32
	_ = v7370
	var v7371 int32
	_ = v7371
	var v7377 int32
	_ = v7377
	var v7385 int32
	_ = v7385
	var v7390 int32
	_ = v7390
	var v7391 int32
	_ = v7391
	var v7397 int32
	_ = v7397
	var v7398 int32
	_ = v7398
	var v7404 int32
	_ = v7404
	var v7412 int32
	_ = v7412
	var v7418 int32
	_ = v7418
	var v7419 int32
	_ = v7419
	var v7426 int32
	_ = v7426
	var v7435 int32
	_ = v7435
	var v7436 int32
	_ = v7436
	var v7438 int32
	_ = v7438
	var v7439 int32
	_ = v7439
	var v7442 int32
	_ = v7442
	var v7443 int32
	_ = v7443
	var v7448 int32
	_ = v7448
	var v7449 int32
	_ = v7449
	var v7452 int32
	_ = v7452
	var v7453 int32
	_ = v7453
	var v7455 int32
	_ = v7455
	var v7460 int32
	_ = v7460
	var v7462 int32
	_ = v7462
	var v7463 int32
	_ = v7463
	var v7467 int32
	_ = v7467
	var v7469 int32
	_ = v7469
	var v7472 int32
	_ = v7472
	var v7473 int32
	_ = v7473
	var v7476 int32
	_ = v7476
	var v7479 int32
	_ = v7479
	var v7486 int32
	_ = v7486
	var v7491 int32
	_ = v7491
	var v7492 int32
	_ = v7492
	var v7494 int32
	_ = v7494
	var v7498 int32
	_ = v7498
	var v7506 int32
	_ = v7506
	var v7511 int32
	_ = v7511
	var v7542 int32
	_ = v7542
	var v7543 int32
	_ = v7543
	var v7546 int32
	_ = v7546
	var v7553 int32
	_ = v7553
	var v7558 int32
	_ = v7558
	var v7559 int32
	_ = v7559
	var v7561 int32
	_ = v7561
	var v7591 int32
	_ = v7591
	var v7600 int32
	_ = v7600
	var v7601 int32
	_ = v7601
	var v7602 int32
	_ = v7602
	var v7603 int32
	_ = v7603
	var v7604 int32
	_ = v7604
	var v7606 int32
	_ = v7606
	var v7611 int32
	_ = v7611
	var v7612 int32
	_ = v7612
	var v7617 int32
	_ = v7617
	var v7618 int32
	_ = v7618
	var v7623 int32
	_ = v7623
	var v7624 int32
	_ = v7624
	var v7627 int32
	_ = v7627
	var v7628 int32
	_ = v7628
	var v7630 int32
	_ = v7630
	var v7635 int32
	_ = v7635
	var v7637 int32
	_ = v7637
	var v7638 int32
	_ = v7638
	var v7642 int32
	_ = v7642
	var v7644 int32
	_ = v7644
	var v7647 int32
	_ = v7647
	var v7648 int32
	_ = v7648
	var v7651 int32
	_ = v7651
	var v7654 int32
	_ = v7654
	var v7661 int32
	_ = v7661
	var v7666 int32
	_ = v7666
	var v7667 int32
	_ = v7667
	var v7669 int32
	_ = v7669
	var v7672 int32
	_ = v7672
	var v7673 int32
	_ = v7673
	var v7677 int32
	_ = v7677
	var v7678 int32
	_ = v7678
	var v7681 int32
	_ = v7681
	var v7686 int32
	_ = v7686
	var v7687 int32
	_ = v7687
	var v7694 int32
	_ = v7694
	var v7695 int32
	_ = v7695
	var v7698 int32
	_ = v7698
	var v7701 int32
	_ = v7701
	var v7704 int32
	_ = v7704
	var v7706 int32
	_ = v7706
	var v7711 int32
	_ = v7711
	var v7712 int32
	_ = v7712
	var v7717 int32
	_ = v7717
	var v7718 int32
	_ = v7718
	var v7721 int32
	_ = v7721
	var v7724 int32
	_ = v7724
	var v7725 int32
	_ = v7725
	var v7729 int32
	_ = v7729
	var v7730 int32
	_ = v7730
	var v7733 int32
	_ = v7733
	var v7734 int32
	_ = v7734
	var v7737 int32
	_ = v7737
	var v7744 int32
	_ = v7744
	var v7745 int32
	_ = v7745
	var v7749 int32
	_ = v7749
	var v7752 int32
	_ = v7752
	var v7753 int32
	_ = v7753
	var v7756 int32
	_ = v7756
	var v7757 int32
	_ = v7757
	var v7761 int32
	_ = v7761
	var v7762 int32
	_ = v7762
	var v7765 int32
	_ = v7765
	var v7766 int32
	_ = v7766
	var v7769 int32
	_ = v7769
	var v7776 int32
	_ = v7776
	var v7777 int32
	_ = v7777
	var v7780 int32
	_ = v7780
	var v7781 int32
	_ = v7781
	var v7788 int32
	_ = v7788
	var v7789 int32
	_ = v7789
	var v7790 int32
	_ = v7790
	var v7798 int32
	_ = v7798
	var v7799 int32
	_ = v7799
	var v7804 int32
	_ = v7804
	var v7807 int32
	_ = v7807
	var v7816 int32
	_ = v7816
	var v7817 int32
	_ = v7817
	var v7819 int32
	_ = v7819
	var v7848 int32
	_ = v7848
	var v7851 int32
	_ = v7851
	var v7860 int32
	_ = v7860
	var v7861 int32
	_ = v7861
	var v7862 int32
	_ = v7862
	var v7863 int32
	_ = v7863
	var v7864 int32
	_ = v7864
	var v7866 int32
	_ = v7866
	var v7899 int32
	_ = v7899
	var v7901 int32
	_ = v7901
	var v7913 int32
	_ = v7913
	var v7914 int32
	_ = v7914
	var v7922 int32
	_ = v7922
	var v7923 int32
	_ = v7923
	var v7928 int32
	_ = v7928
	var v7937 int32
	_ = v7937
	var v7942 int32
	_ = v7942
	var v7943 int32
	_ = v7943
	var v7946 int32
	_ = v7946
	var v7951 int32
	_ = v7951
	var v7952 int32
	_ = v7952
	var v7957 int32
	_ = v7957
	var v7958 int32
	_ = v7958
	var v7961 int32
	_ = v7961
	var v7962 int32
	_ = v7962
	var v7964 int32
	_ = v7964
	var v7969 int32
	_ = v7969
	var v7971 int32
	_ = v7971
	var v7972 int32
	_ = v7972
	var v7976 int32
	_ = v7976
	var v7978 int32
	_ = v7978
	var v7981 int32
	_ = v7981
	var v7982 int32
	_ = v7982
	var v7985 int32
	_ = v7985
	var v7988 int32
	_ = v7988
	var v7995 int32
	_ = v7995
	var v7999 int32
	_ = v7999
	var v8003 int32
	_ = v8003
	var v8004 int32
	_ = v8004
	var v8005 int32
	_ = v8005
	var v8011 int32
	_ = v8011
	var v8014 int32
	_ = v8014
	var v8015 int32
	_ = v8015
	var v8020 int32
	_ = v8020
	var v8021 int32
	_ = v8021
	var v8024 int32
	_ = v8024
	var v8025 int32
	_ = v8025
	var v8027 int32
	_ = v8027
	var v8032 int32
	_ = v8032
	var v8034 int32
	_ = v8034
	var v8035 int32
	_ = v8035
	var v8039 int32
	_ = v8039
	var v8041 int32
	_ = v8041
	var v8044 int32
	_ = v8044
	var v8045 int32
	_ = v8045
	var v8048 int32
	_ = v8048
	var v8051 int32
	_ = v8051
	var v8058 int32
	_ = v8058
	var v8064 int32
	_ = v8064
	var v8065 int32
	_ = v8065
	var v8066 int32
	_ = v8066
	var v8070 int32
	_ = v8070
	var v8073 int32
	_ = v8073
	var v8074 int32
	_ = v8074
	var v8078 int32
	_ = v8078
	var v8079 int32
	_ = v8079
	var v8081 int32
	_ = v8081
	var v8083 int32
	_ = v8083
	var v8094 int32
	_ = v8094
	var v8095 int32
	_ = v8095
	var v8103 int32
	_ = v8103
	var v8104 int32
	_ = v8104
	var v8105 int32
	_ = v8105
	var v8108 int32
	_ = v8108
	var v8111 int32
	_ = v8111
	var v8112 int32
	_ = v8112
	var v8116 int32
	_ = v8116
	var v8117 int32
	_ = v8117
	var v8120 int32
	_ = v8120
	var v8121 int32
	_ = v8121
	var v8124 int32
	_ = v8124
	var v8131 int32
	_ = v8131
	var v8132 int32
	_ = v8132
	var v8136 int32
	_ = v8136
	var v8139 int32
	_ = v8139
	var v8140 int32
	_ = v8140
	var v8144 int32
	_ = v8144
	var v8145 int32
	_ = v8145
	var v8148 int32
	_ = v8148
	var v8149 int32
	_ = v8149
	var v8152 int32
	_ = v8152
	var v8159 int32
	_ = v8159
	var v8160 int32
	_ = v8160
	var v8165 int32
	_ = v8165
	var v8168 int32
	_ = v8168
	var v8169 int32
	_ = v8169
	var v8173 int32
	_ = v8173
	var v8174 int32
	_ = v8174
	var v8177 int32
	_ = v8177
	var v8178 int32
	_ = v8178
	var v8181 int32
	_ = v8181
	var v8188 int32
	_ = v8188
	var v8189 int32
	_ = v8189
	var v8194 int32
	_ = v8194
	var v8197 int32
	_ = v8197
	var v8198 int32
	_ = v8198
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
	var v8223 int32
	_ = v8223
	var v8226 int32
	_ = v8226
	var v8227 int32
	_ = v8227
	var v8231 int32
	_ = v8231
	var v8232 int32
	_ = v8232
	var v8235 int32
	_ = v8235
	var v8236 int32
	_ = v8236
	var v8239 int32
	_ = v8239
	var v8246 int32
	_ = v8246
	var v8247 int32
	_ = v8247
	var v8250 int32
	_ = v8250
	var v8254 int32
	_ = v8254
	var v8257 int32
	_ = v8257
	var v8265 int32
	_ = v8265
	var v8266 int32
	_ = v8266
	var v8268 int32
	_ = v8268
	var v8271 int32
	_ = v8271
	var v8272 int32
	_ = v8272
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
	var v8310 int32
	_ = v8310
	var v8311 int32
	_ = v8311
	var v8313 int32
	_ = v8313
	var v8316 int32
	_ = v8316
	var v8319 int32
	_ = v8319
	var v8320 int32
	_ = v8320
	var v8324 int32
	_ = v8324
	var v8325 int32
	_ = v8325
	var v8328 int32
	_ = v8328
	var v8329 int32
	_ = v8329
	var v8332 int32
	_ = v8332
	var v8339 int32
	_ = v8339
	var v8340 int32
	_ = v8340
	var v8343 int32
	_ = v8343
	var v8347 int32
	_ = v8347
	var v8350 int32
	_ = v8350
	var v8351 int32
	_ = v8351
	var v8355 int32
	_ = v8355
	var v8356 int32
	_ = v8356
	var v8359 int32
	_ = v8359
	var v8360 int32
	_ = v8360
	var v8363 int32
	_ = v8363
	var v8370 int32
	_ = v8370
	var v8371 int32
	_ = v8371
	var v8381 int32
	_ = v8381
	var v8382 int32
	_ = v8382
	var v8385 int32
	_ = v8385
	var v8388 int32
	_ = v8388
	var v8389 int32
	_ = v8389
	var v8393 int32
	_ = v8393
	var v8394 int32
	_ = v8394
	var v8397 int32
	_ = v8397
	var v8398 int32
	_ = v8398
	var v8401 int32
	_ = v8401
	var v8408 int32
	_ = v8408
	var v8409 int32
	_ = v8409
	var v8421 int32
	_ = v8421
	var v8429 int32
	_ = v8429
	var v8430 int32
	_ = v8430
	var v8437 int32
	_ = v8437
	var v8440 int32
	_ = v8440
	var v8443 int32
	_ = v8443
	var v8451 int32
	_ = v8451
	var v8452 int32
	_ = v8452
	var v8453 int32
	_ = v8453
	var v8457 int32
	_ = v8457
	var v8460 int32
	_ = v8460
	var v8467 int32
	_ = v8467
	var v8468 int32
	_ = v8468
	var v8476 int32
	_ = v8476
	var v8479 int32
	_ = v8479
	var v8480 int32
	_ = v8480
	var v8485 int32
	_ = v8485
	var v8486 int32
	_ = v8486
	var v8491 int32
	_ = v8491
	var v8492 int32
	_ = v8492
	var v8495 int32
	_ = v8495
	var v8496 int32
	_ = v8496
	var v8498 int32
	_ = v8498
	var v8503 int32
	_ = v8503
	var v8505 int32
	_ = v8505
	var v8506 int32
	_ = v8506
	var v8510 int32
	_ = v8510
	var v8512 int32
	_ = v8512
	var v8515 int32
	_ = v8515
	var v8516 int32
	_ = v8516
	var v8519 int32
	_ = v8519
	var v8522 int32
	_ = v8522
	var v8529 int32
	_ = v8529
	var v8534 int32
	_ = v8534
	var v8535 int32
	_ = v8535
	var v8537 int32
	_ = v8537
	var v8542 int32
	_ = v8542
	var v8543 int32
	_ = v8543
	var v8549 int32
	_ = v8549
	var v8550 int32
	_ = v8550
	var v8555 int32
	_ = v8555
	var v8556 int32
	_ = v8556
	var v8561 int32
	_ = v8561
	var v8562 int32
	_ = v8562
	var v8565 int32
	_ = v8565
	var v8566 int32
	_ = v8566
	var v8568 int32
	_ = v8568
	var v8573 int32
	_ = v8573
	var v8575 int32
	_ = v8575
	var v8576 int32
	_ = v8576
	var v8580 int32
	_ = v8580
	var v8582 int32
	_ = v8582
	var v8585 int32
	_ = v8585
	var v8586 int32
	_ = v8586
	var v8589 int32
	_ = v8589
	var v8592 int32
	_ = v8592
	var v8599 int32
	_ = v8599
	var v8604 int32
	_ = v8604
	var v8605 int32
	_ = v8605
	var v8607 int32
	_ = v8607
	var v8612 int32
	_ = v8612
	var v8618 int32
	_ = v8618
	var v8619 int32
	_ = v8619
	var v8624 int32
	_ = v8624
	var v8625 int32
	_ = v8625
	var v8630 int32
	_ = v8630
	var v8631 int32
	_ = v8631
	var v8634 int32
	_ = v8634
	var v8635 int32
	_ = v8635
	var v8637 int32
	_ = v8637
	var v8642 int32
	_ = v8642
	var v8644 int32
	_ = v8644
	var v8645 int32
	_ = v8645
	var v8649 int32
	_ = v8649
	var v8651 int32
	_ = v8651
	var v8654 int32
	_ = v8654
	var v8655 int32
	_ = v8655
	var v8658 int32
	_ = v8658
	var v8661 int32
	_ = v8661
	var v8668 int32
	_ = v8668
	var v8673 int32
	_ = v8673
	var v8674 int32
	_ = v8674
	var v8676 int32
	_ = v8676
	var v8681 int32
	_ = v8681
	var v8692 int32
	_ = v8692
	var v8693 int32
	_ = v8693
	var v8694 int32
	_ = v8694
	var v8695 int32
	_ = v8695
	var v8698 int32
	_ = v8698
	var v8699 int32
	_ = v8699
	var v8700 int32
	_ = v8700
	var v8704 int32
	_ = v8704
	var v8706 int32
	_ = v8706
	var v8707 int32
	_ = v8707
	var v8709 int32
	_ = v8709
	var v8712 int32
	_ = v8712
	var v8713 int32
	_ = v8713
	var v8718 int32
	_ = v8718
	var v8719 int32
	_ = v8719
	var v8722 int32
	_ = v8722
	var v8723 int32
	_ = v8723
	var v8725 int32
	_ = v8725
	var v8730 int32
	_ = v8730
	var v8732 int32
	_ = v8732
	var v8733 int32
	_ = v8733
	var v8737 int32
	_ = v8737
	var v8739 int32
	_ = v8739
	var v8742 int32
	_ = v8742
	var v8743 int32
	_ = v8743
	var v8746 int32
	_ = v8746
	var v8749 int32
	_ = v8749
	var v8756 int32
	_ = v8756
	var v8760 int32
	_ = v8760
	var v8761 int32
	_ = v8761
	var v8763 int32
	_ = v8763
	var v8764 int32
	_ = v8764
	var v8769 int32
	_ = v8769
	var v8771 int32
	_ = v8771
	var v8772 int32
	_ = v8772
	var v8774 int32
	_ = v8774
	var v8775 int32
	_ = v8775
	var v8776 int32
	_ = v8776
	var v8778 int32
	_ = v8778
	var v8784 int32
	_ = v8784
	var v8785 int32
	_ = v8785
	var v8787 int32
	_ = v8787
	var v8788 int32
	_ = v8788
	var v8790 int32
	_ = v8790
	var v8791 int32
	_ = v8791
	var v8792 int32
	_ = v8792
	var v8794 int32
	_ = v8794
	var v8799 int32
	_ = v8799
	var v8800 int32
	_ = v8800
	var v8805 int32
	_ = v8805
	var v8806 int32
	_ = v8806
	var v8807 int32
	_ = v8807
	var v8808 int32
	_ = v8808
	var v8810 int32
	_ = v8810
	var v8816 int32
	_ = v8816
	var v8817 int32
	_ = v8817
	var v8820 int32
	_ = v8820
	var v8821 int32
	_ = v8821
	var v8824 int32
	_ = v8824
	var v8825 int32
	_ = v8825
	var v8830 int32
	_ = v8830
	var v8831 int32
	_ = v8831
	var v8834 int32
	_ = v8834
	var v8835 int32
	_ = v8835
	var v8837 int32
	_ = v8837
	var v8842 int32
	_ = v8842
	var v8844 int32
	_ = v8844
	var v8845 int32
	_ = v8845
	var v8849 int32
	_ = v8849
	var v8851 int32
	_ = v8851
	var v8854 int32
	_ = v8854
	var v8855 int32
	_ = v8855
	var v8858 int32
	_ = v8858
	var v8861 int32
	_ = v8861
	var v8868 int32
	_ = v8868
	var v8874 int32
	_ = v8874
	var v8876 int32
	_ = v8876
	var v8880 int32
	_ = v8880
	var v8881 int32
	_ = v8881
	var v8885 int32
	_ = v8885
	var v8906 int32
	_ = v8906
	var v8907 int32
	_ = v8907
	var v8909 int32
	_ = v8909
	var v8911 int32
	_ = v8911
	var v8913 int32
	_ = v8913
	var v8914 int32
	_ = v8914
	var v8917 int32
	_ = v8917
	var v8918 int32
	_ = v8918
	var v8922 int32
	_ = v8922
	var v8923 int32
	_ = v8923
	var v8926 int32
	_ = v8926
	var v8927 int32
	_ = v8927
	var v8930 int32
	_ = v8930
	var v8937 int32
	_ = v8937
	var v8938 int32
	_ = v8938
	var v8940 int32
	_ = v8940
	var v8941 int32
	_ = v8941
	var v8947 int32
	_ = v8947
	var v8948 int32
	_ = v8948
	var v8951 int32
	_ = v8951
	var v8952 int32
	_ = v8952
	var v8955 int32
	_ = v8955
	var v8959 int32
	_ = v8959
	var v8961 int32
	_ = v8961
	var v8962 int64
	_ = v8962
	var v8970 int32
	_ = v8970
	var v8974 int32
	_ = v8974
	var v8978 int32
	_ = v8978
	var v8984 int32
	_ = v8984
	var v8988 int32
	_ = v8988
	var v8989 int32
	_ = v8989
	var v8996 int32
	_ = v8996
	var v8997 int32
	_ = v8997
	var v8998 int32
	_ = v8998
	var v9002 int32
	_ = v9002
	var v9005 int32
	_ = v9005
	var v9009 int32
	_ = v9009
	var v9010 int32
	_ = v9010
	var v9018 int32
	_ = v9018
	var v9024 int32
	_ = v9024
	var v9026 int32
	_ = v9026
	var v9030 int32
	_ = v9030
	var v9038 int32
	_ = v9038
	var v9042 int32
	_ = v9042
	var v9043 int32
	_ = v9043
	var v9044 int32
	_ = v9044
	var v9045 int32
	_ = v9045
	var v9046 int32
	_ = v9046
	var v9047 int32
	_ = v9047
	var v9048 int32
	_ = v9048
	var v9052 int32
	_ = v9052
	var v9054 int32
	_ = v9054
	var v9087 int32
	_ = v9087
	var v9091 int32
	_ = v9091
	var v9099 int32
	_ = v9099
	var v9100 int32
	_ = v9100
	var v9103 int32
	_ = v9103
	var v9107 int32
	_ = v9107
	var v9115 int32
	_ = v9115
	var v9116 int32
	_ = v9116
	var v9119 int32
	_ = v9119
	var v9123 int32
	_ = v9123
	var v9131 int32
	_ = v9131
	var v9132 int32
	_ = v9132
	var v9134 int32
	_ = v9134
	var v9137 int32
	_ = v9137
	var v9141 int32
	_ = v9141
	var v9142 int32
	_ = v9142
	var v9145 int32
	_ = v9145
	var v9146 int32
	_ = v9146
	var v9151 int32
	_ = v9151
	var v9155 int32
	_ = v9155
	var v9156 int32
	_ = v9156
	var v9159 int32
	_ = v9159
	var v9160 int32
	_ = v9160
	var v9164 int32
	_ = v9164
	var v9168 int32
	_ = v9168
	var v9170 int32
	_ = v9170
	var v9172 int32
	_ = v9172
	var v9173 int32
	_ = v9173
	var v9174 int32
	_ = v9174
	var v9176 int32
	_ = v9176
	var v9188 int32
	_ = v9188
	var v9195 int32
	_ = v9195
	var v9199 int32
	_ = v9199
	var v9202 int32
	_ = v9202
	var v9203 int32
	_ = v9203
	var v9204 int32
	_ = v9204
	var v9205 int32
	_ = v9205
	var v9206 int32
	_ = v9206
	var v9212 int32
	_ = v9212
	var v9215 int32
	_ = v9215
	var v9216 int32
	_ = v9216
	var v9217 int32
	_ = v9217
	var v9222 int32
	_ = v9222
	var v9226 int32
	_ = v9226
	var v9229 int32
	_ = v9229
	var v9230 int32
	_ = v9230
	var v9236 int32
	_ = v9236
	var v9239 int32
	_ = v9239
	var v9240 int32
	_ = v9240
	var v9241 int32
	_ = v9241
	var v9246 int32
	_ = v9246
	var v9250 int32
	_ = v9250
	var v9253 int32
	_ = v9253
	var v9254 int32
	_ = v9254
	var v9260 int32
	_ = v9260
	var v9261 int32
	_ = v9261
	var v9262 int32
	_ = v9262
	var v9263 int32
	_ = v9263
	var v9268 int32
	_ = v9268
	var v9272 int32
	_ = v9272
	var v9275 int32
	_ = v9275
	var v9276 int32
	_ = v9276
	var v9282 int32
	_ = v9282
	var v9283 int32
	_ = v9283
	var v9284 int32
	_ = v9284
	var v9285 int32
	_ = v9285
	var v9290 int32
	_ = v9290
	var v9295 int32
	_ = v9295
	var v9298 int32
	_ = v9298
	var v9299 int32
	_ = v9299
	var v9300 int32
	_ = v9300
	var v9301 int32
	_ = v9301
	var v9307 int32
	_ = v9307
	var v9308 int32
	_ = v9308
	var v9309 int32
	_ = v9309
	var v9310 int32
	_ = v9310
	var v9315 int32
	_ = v9315
	var v9321 int32
	_ = v9321
	var v9327 int32
	_ = v9327
	var v9333 int32
	_ = v9333
	var v9338 int32
	_ = v9338
	var v9342 int32
	_ = v9342
	var v9347 int32
	_ = v9347
	var v9351 int32
	_ = v9351
	var v9354 int32
	_ = v9354
	var v9355 int32
	_ = v9355
	var v9356 int32
	_ = v9356
	var v9362 int32
	_ = v9362
	var v9363 int32
	_ = v9363
	var v9364 int32
	_ = v9364
	var v9365 int32
	_ = v9365
	var v9370 int32
	_ = v9370
	var v9374 int32
	_ = v9374
	var v9377 int32
	_ = v9377
	var v9381 int32
	_ = v9381
	var v9384 int32
	_ = v9384
	var v9385 int32
	_ = v9385
	var v9386 int32
	_ = v9386
	var v9391 int32
	_ = v9391
	var v9395 int32
	_ = v9395
	var v9398 int32
	_ = v9398
	var v9402 int32
	_ = v9402
	var v9405 int32
	_ = v9405
	var v9406 int32
	_ = v9406
	var v9407 int32
	_ = v9407
	var v9412 int32
	_ = v9412
	var v9416 int32
	_ = v9416
	var v9419 int32
	_ = v9419
	var v9423 int32
	_ = v9423
	var v9424 int32
	_ = v9424
	var v9425 int32
	_ = v9425
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
	var v9465 int32
	_ = v9465
	var v9466 int32
	_ = v9466
	var v9467 int32
	_ = v9467
	var v9472 int32
	_ = v9472
	var v9476 int32
	_ = v9476
	var v9479 int32
	_ = v9479
	var v9480 int32
	_ = v9480
	var v9486 int32
	_ = v9486
	var v9489 int32
	_ = v9489
	var v9490 int32
	_ = v9490
	var v9491 int32
	_ = v9491
	var v9496 int32
	_ = v9496
	var v9500 int32
	_ = v9500
	var v9503 int32
	_ = v9503
	var v9506 int32
	_ = v9506
	var v9507 int32
	_ = v9507
	var v9510 int32
	_ = v9510
	var v9511 int32
	_ = v9511
	var v9512 int32
	_ = v9512
	var v9513 int32
	_ = v9513
	var v9518 int32
	_ = v9518
	var v9522 int32
	_ = v9522
	var v9525 int32
	_ = v9525
	var v9529 int32
	_ = v9529
	var v9530 int32
	_ = v9530
	var v9531 int32
	_ = v9531
	var v9536 int32
	_ = v9536
	var v9540 int32
	_ = v9540
	var v9543 int32
	_ = v9543
	var v9547 int32
	_ = v9547
	var v9548 int32
	_ = v9548
	var v9549 int32
	_ = v9549
	var v9550 int32
	_ = v9550
	var v9555 int32
	_ = v9555
	var v9559 int32
	_ = v9559
	var v9562 int32
	_ = v9562
	var v9566 int32
	_ = v9566
	var v9567 int32
	_ = v9567
	var v9568 int32
	_ = v9568
	var v9573 int32
	_ = v9573
	var v9576 int32
	_ = v9576
	var v9580 int32
	_ = v9580
	var v9581 int32
	_ = v9581
	var v9582 int32
	_ = v9582
	var v9583 int32
	_ = v9583
	var v9588 int32
	_ = v9588
	var v9592 int32
	_ = v9592
	var v9595 int32
	_ = v9595
	var v9599 int32
	_ = v9599
	var v9600 int32
	_ = v9600
	var v9601 int32
	_ = v9601
	var v9602 int32
	_ = v9602
	var v9607 int32
	_ = v9607
	var v9613 int32
	_ = v9613
	var v9619 int32
	_ = v9619
	var v9625 int32
	_ = v9625
	var v9631 int32
	_ = v9631
	var v9638 int32
	_ = v9638
	var v9644 int32
	_ = v9644
	var v9648 int32
	_ = v9648
	var v9651 int32
	_ = v9651
	var v9655 int32
	_ = v9655
	var v9660 int32
	_ = v9660
	var v9661 int32
	_ = v9661
	var v9663 int32
	_ = v9663
	var v9664 int32
	_ = v9664
	var v9666 int32
	_ = v9666
	var v9672 int32
	_ = v9672
	var v9678 int32
	_ = v9678
	var v9684 int32
	_ = v9684
	var v9688 int32
	_ = v9688
	var v9691 int32
	_ = v9691
	var v9695 int32
	_ = v9695
	var v9698 int32
	_ = v9698
	var v9699 int32
	_ = v9699
	var v9700 int32
	_ = v9700
	var v9705 int32
	_ = v9705
	var v9709 int32
	_ = v9709
	var v9712 int32
	_ = v9712
	var v9716 int32
	_ = v9716
	var v9717 int32
	_ = v9717
	var v9718 int32
	_ = v9718
	var v9719 int32
	_ = v9719
	var v9724 int32
	_ = v9724
	var v9728 int32
	_ = v9728
	var v9731 int32
	_ = v9731
	var v9732 int32
	_ = v9732
	var v9733 int32
	_ = v9733
	var v9739 int32
	_ = v9739
	var v9740 int32
	_ = v9740
	var v9741 int32
	_ = v9741
	var v9742 int32
	_ = v9742
	var v9747 int32
	_ = v9747
	var v9753 int32
	_ = v9753
	var v9759 int32
	_ = v9759
	var v9765 int32
	_ = v9765
	var v9771 int32
	_ = v9771
	var v9779 int32
	_ = v9779
	var v9784 int32
	_ = v9784
	var v9787 int32
	_ = v9787
	var v9791 int32
	_ = v9791
	var v9796 int32
	_ = v9796
	var v9798 int32
	_ = v9798
	var v9801 int32
	_ = v9801
	var v9814 int32
	_ = v9814
	var v9815 int32
	_ = v9815
	var v9816 int32
	_ = v9816
	var v9822 int32
	_ = v9822
	var v9825 int32
	_ = v9825
	var v9829 int32
	_ = v9829
	var v9830 int32
	_ = v9830
	var v9831 int32
	_ = v9831
	var v9836 int32
	_ = v9836
	var v9837 int32
	_ = v9837
	var v9838 int32
	_ = v9838
	var v9842 int32
	_ = v9842
	var v9845 int32
	_ = v9845
	var v9846 int32
	_ = v9846
	var v9847 int32
	_ = v9847
	var v9848 int32
	_ = v9848
	var v9851 int32
	_ = v9851
	var v9857 int32
	_ = v9857
	var v9858 int32
	_ = v9858
	var v9868 int32
	_ = v9868
	var v9869 int32
	_ = v9869
	var v9873 int32
	_ = v9873
	var v9879 int32
	_ = v9879
	var v9885 int32
	_ = v9885
	var v9894 int32
	_ = v9894
	var v9895 int32
	_ = v9895
	var v9896 int32
	_ = v9896
	var v9900 int32
	_ = v9900
	var v9904 int32
	_ = v9904
	var v9912 int32
	_ = v9912
	var v9913 int32
	_ = v9913
	var v9914 int32
	_ = v9914
	var v9917 int32
	_ = v9917
	var v9920 int32
	_ = v9920
	var v9923 int32
	_ = v9923
	var v9926 int32
	_ = v9926
	var v9929 int32
	_ = v9929
	var v9931 int32
	_ = v9931
	var v9932 int32
	_ = v9932
	var v9934 int32
	_ = v9934
	var v9935 int32
	_ = v9935
	var v9937 int32
	_ = v9937
	var v9938 int32
	_ = v9938
	var v9942 int32
	_ = v9942
	var v9943 int32
	_ = v9943
	var v9945 int32
	_ = v9945
	var v9957 int32
	_ = v9957
	var v9960 int32
	_ = v9960
	var v9964 int32
	_ = v9964
	var v9967 int32
	_ = v9967
	var v9968 int32
	_ = v9968
	var v9969 int32
	_ = v9969
	var v9974 int32
	_ = v9974
	var v9980 int32
	_ = v9980
	var v9984 int32
	_ = v9984
	var v9989 int32
	_ = v9989
	var v9995 int32
	_ = v9995
	var v10025 int32
	_ = v10025
	var v10033 int32
	_ = v10033
	var v10047 int32
	_ = v10047
	var v10050 int32
	_ = v10050
	var v10056 int32
	_ = v10056
	var v10057 int32
	_ = v10057
	var v10061 int32
	_ = v10061
	var v10064 int32
	_ = v10064
	var v10066 int32
	_ = v10066
	var v10067 int32
	_ = v10067
	var v10068 int32
	_ = v10068
	var v10071 int32
	_ = v10071
	var v10072 int32
	_ = v10072
	var v10085 int32
	_ = v10085
	var v10087 int32
	_ = v10087
	var v10088 int32
	_ = v10088
	var v10094 int32
	_ = v10094
	var v10096 int32
	_ = v10096
	var v10098 int32
	_ = v10098
	var v10099 int32
	_ = v10099
	var v10101 int32
	_ = v10101
	var v10102 int32
	_ = v10102
	var v10103 int32
	_ = v10103
	var v10104 int32
	_ = v10104
	var v10105 int32
	_ = v10105
	var v10107 int32
	_ = v10107
	var v10113 int32
	_ = v10113
	var v10120 int32
	_ = v10120
	var v10149 int32
	_ = v10149
	var v10165 int32
	_ = v10165
	var v10168 int32
	_ = v10168
	var v10169 int64
	_ = v10169
	var v10171 int64
	_ = v10171
	var v10175 int32
	_ = v10175
	var v10176 int32
	_ = v10176
	var v10178 int32
	_ = v10178
	var v10179 int32
	_ = v10179
	var v10182 int32
	_ = v10182
	var v10186 int32
	_ = v10186
	var v10189 int32
	_ = v10189
	var v10190 int32
	_ = v10190
	var v10197 int32
	_ = v10197
	var v10205 int32
	_ = v10205
	var v10208 int32
	_ = v10208
	var v10212 int32
	_ = v10212
	var v10214 int32
	_ = v10214
	var v10218 int32
	_ = v10218
	var v10223 int32
	_ = v10223
	var v10226 int32
	_ = v10226
	var v10240 int32
	_ = v10240
	v25 = m.G0
	v27 = v25 - int32(4816)
	m.G0 = v27
	v30 = v27 + int32(312)
	v33 = v27 + int32(320)
	v35 = v27 + int32(1120)
	v37 = v27 + int32(4320)
	v42 = int32(0)
	v44 = v35
	v48 = int32(-2)
	v53 = v37
	v55 = v37
	v56 = v33
	v57 = int32(200)
	v60 = v33
	v61 = v35
	goto L1
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v42)
	v65 = v57 << (uint(int32(1)) % 32)
	if base.Ui32(v53) < base.Ui32(v55+v65-int32(2)) {
		goto L59
	} else {
		goto L60
	}
L2:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(261077))
	mBase = m.M
	v10240 = m.ExcPending
	if v10240 != 0 {
		goto L66
	} else {
		goto L2697
	}
L3:
	;
	goto L2
L4:
	;
	v42 = v10212
	v44 = v10214
	v48 = v10218
	v53 = v10223 + int32(2)
	v55 = v136
	v56 = v10226
	v57 = v138
	v60 = v140
	v61 = v141
	goto L1
L5:
	;
	v10165 = int32(0) - v247
	v10168 = v132 + v10165<<(uint(int32(4))%32)
	v10169 = *(*int64)(unsafe.Add(mBase, uint32(v27)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v10168)+16)) = v10169
	v10171 = *(*int64)(unsafe.Add(mBase, uint32(v27)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v10168)+24)) = v10171
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v262
	v10175 = v10168 + int32(16)
	v10176 = int32(1)
	v10178 = v134 + v10165<<(uint(v10176)%32)
	v10179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10178))))
	v10182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[1354]))))
	v10186 = (v10182 - int32(137)) << (uint(v10176) % 32)
	v10189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10186)+uint32(_consts[1355]))))
	v10190 = v10179 + v10189
	if base.Ui32(int32(1293)) < base.Ui32(v10190) {
		goto L2694
	} else {
		goto L2695
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v10120
	v10149 = int32(-2)
	goto L5
L7:
	;
	v9989 = int32(0)
	v9995 = v9984
	goto L2661
L8:
	;
	F_plpgsql_push_back_token(m, v2516, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9980 = m.ExcPending
	if v9980 != 0 {
		goto L66
	} else {
		goto L2660
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9837)+4)) = int32(2)
	v9842 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1356])))
	if v9842 == int32(1) {
		goto L2638
	} else {
		goto L2639
	}
L10:
	;
	v9798 = int32(1)
	v9801 = int32(0)
	v9814 = F_read_sql_construct(m, int32(269), int32(336), v9801, int32(526588), v9801, v9798, v9801, v27+int32(4764), v27+int32(4752), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9815 = m.ExcPending
	if v9815 != 0 {
		goto L66
	} else {
		goto L2631
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9784 = m.ExcPending
	if v9784 != 0 {
		goto L66
	} else {
		goto L2627
	}
L12:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(730103))
	mBase = m.M
	v9779 = m.ExcPending
	if v9779 != 0 {
		goto L66
	} else {
		goto L2626
	}
L13:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v9771 = m.ExcPending
	if v9771 != 0 {
		goto L66
	} else {
		goto L2625
	}
L14:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(413776))
	mBase = m.M
	v9765 = m.ExcPending
	if v9765 != 0 {
		goto L66
	} else {
		goto L2624
	}
L15:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(413776))
	mBase = m.M
	v9759 = m.ExcPending
	if v9759 != 0 {
		goto L66
	} else {
		goto L2623
	}
L16:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v9753 = m.ExcPending
	if v9753 != 0 {
		goto L66
	} else {
		goto L2622
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9728 = m.ExcPending
	if v9728 != 0 {
		goto L66
	} else {
		goto L2617
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9709 = m.ExcPending
	if v9709 != 0 {
		goto L66
	} else {
		goto L2612
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9688 = m.ExcPending
	if v9688 != 0 {
		goto L66
	} else {
		goto L2607
	}
L20:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v9684 = m.ExcPending
	if v9684 != 0 {
		goto L66
	} else {
		goto L2606
	}
L21:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v9678 = m.ExcPending
	if v9678 != 0 {
		goto L66
	} else {
		goto L2605
	}
L22:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v9672 = m.ExcPending
	if v9672 != 0 {
		goto L66
	} else {
		goto L2604
	}
L23:
	;
	v9664 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v9664, l1)
	mBase = m.M
	v9666 = m.ExcPending
	if v9666 != 0 {
		goto L66
	} else {
		goto L2603
	}
L24:
	;
	v9661 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v9661, l1)
	mBase = m.M
	v9663 = m.ExcPending
	if v9663 != 0 {
		goto L66
	} else {
		goto L2602
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9648 = m.ExcPending
	if v9648 != 0 {
		goto L66
	} else {
		goto L2598
	}
L26:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(730220))
	mBase = m.M
	v9644 = m.ExcPending
	if v9644 != 0 {
		goto L66
	} else {
		goto L2597
	}
L27:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(246502))
	mBase = m.M
	v9638 = m.ExcPending
	if v9638 != 0 {
		goto L66
	} else {
		goto L2596
	}
L28:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(250277))
	mBase = m.M
	v9631 = m.ExcPending
	if v9631 != 0 {
		goto L66
	} else {
		goto L2595
	}
L29:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(413776))
	mBase = m.M
	v9625 = m.ExcPending
	if v9625 != 0 {
		goto L66
	} else {
		goto L2594
	}
L30:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(413776))
	mBase = m.M
	v9619 = m.ExcPending
	if v9619 != 0 {
		goto L66
	} else {
		goto L2593
	}
L31:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v9613 = m.ExcPending
	if v9613 != 0 {
		goto L66
	} else {
		goto L2592
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9592 = m.ExcPending
	if v9592 != 0 {
		goto L66
	} else {
		goto L2587
	}
L33:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9576 = m.ExcPending
	if v9576 != 0 {
		goto L66
	} else {
		goto L2583
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9559 = m.ExcPending
	if v9559 != 0 {
		goto L66
	} else {
		goto L2578
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9540 = m.ExcPending
	if v9540 != 0 {
		goto L66
	} else {
		goto L2573
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9522 = m.ExcPending
	if v9522 != 0 {
		goto L66
	} else {
		goto L2568
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9500 = m.ExcPending
	if v9500 != 0 {
		goto L66
	} else {
		goto L2560
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9476 = m.ExcPending
	if v9476 != 0 {
		goto L66
	} else {
		goto L2555
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9455 = m.ExcPending
	if v9455 != 0 {
		goto L66
	} else {
		goto L2550
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9434 = m.ExcPending
	if v9434 != 0 {
		goto L66
	} else {
		goto L2545
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9416 = m.ExcPending
	if v9416 != 0 {
		goto L66
	} else {
		goto L2540
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9395 = m.ExcPending
	if v9395 != 0 {
		goto L66
	} else {
		goto L2535
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9374 = m.ExcPending
	if v9374 != 0 {
		goto L66
	} else {
		goto L2530
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9351 = m.ExcPending
	if v9351 != 0 {
		goto L66
	} else {
		goto L2524
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9338 = m.ExcPending
	if v9338 != 0 {
		goto L66
	} else {
		goto L2521
	}
L46:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(729856))
	mBase = m.M
	v9333 = m.ExcPending
	if v9333 != 0 {
		goto L66
	} else {
		goto L2520
	}
L47:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(261055))
	mBase = m.M
	v9327 = m.ExcPending
	if v9327 != 0 {
		goto L66
	} else {
		goto L2519
	}
L48:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(261055))
	mBase = m.M
	v9321 = m.ExcPending
	if v9321 != 0 {
		goto L66
	} else {
		goto L2518
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9295 = m.ExcPending
	if v9295 != 0 {
		goto L66
	} else {
		goto L2512
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9272 = m.ExcPending
	if v9272 != 0 {
		goto L66
	} else {
		goto L2507
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9250 = m.ExcPending
	if v9250 != 0 {
		goto L66
	} else {
		goto L2502
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9226 = m.ExcPending
	if v9226 != 0 {
		goto L66
	} else {
		goto L2497
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9199 = m.ExcPending
	if v9199 != 0 {
		goto L66
	} else {
		goto L2491
	}
L54:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), l0, l1, int32(440744))
	mBase = m.M
	v9195 = m.ExcPending
	if v9195 != 0 {
		goto L66
	} else {
		goto L2490
	}
L55:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), l0, l1, int32(212048))
	mBase = m.M
	v9188 = m.ExcPending
	if v9188 != 0 {
		goto L66
	} else {
		goto L2489
	}
L56:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[1357]))))
	v249 = int32(4)
	v251 = v132 + (int32(1)-v247)<<(uint(v249)%32)
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v251)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+304)) = v254
	v260 = v142 - v247<<(uint(int32(2))%32) + v249
	if v247 != 0 {
		goto L116
	} else {
		goto L117
	}
L57:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[1358]))))
	if v237 == int32(0) {
		goto L55
	} else {
		goto L115
	}
L58:
	;
	if v27+int32(4320) != v217 {
		goto L111
	} else {
		goto L112
	}
L59:
	;
	v132 = v44
	v134 = v53
	v136 = v55
	v138 = v57
	v140 = v60
	v141 = v61
	v142 = v56
	goto L61
L60:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v57) {
		goto L54
	} else {
		goto L62
	}
L61:
	;
	v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42<<(uint(int32(1))%32))+uint32(_consts[1359]))))
	if v147 == int32(-249) {
		v232 = v48
		goto L57
	} else {
		goto L88
	}
L62:
	;
	v72 = int32(10000)
	if base.Ui32(v72) <= base.Ui32(v65) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v75 = v72
	goto L65
L64:
	;
	v75 = v65
	goto L65
L65:
	;
	v80 = F_palloc(m, v75*int32(22)+int32(30))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	return int32(0)
L67:
	;
	if v80 == int32(0) {
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v87 = int32(1)
	v90 = (v53-v55)>>(uint(v87)%32) + v87
	v92 = v90 << (uint(v87) % 32)
	if v92 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v101 = v94 + (v75<<(uint(int32(1))%32)+int32(15))&int32(2147483632)
	v103 = v90 << (uint(int32(4)) % 32)
	if v103 != 0 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v93 = F__emscripten_memcpy_bulkmem(m, v80, v55, v92)
	mBase = m.M
	v94 = v93
	goto L72
L71:
	;
	v94 = v80
	goto L72
L72:
	;
	goto L69
L73:
	;
	v108 = v105 + v75<<(uint(int32(4))%32)
	v110 = v90 << (uint(int32(2)) % 32)
	if v110 != 0 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v104 = F__emscripten_memcpy_bulkmem(m, v101, v61, v103)
	mBase = m.M
	v105 = v104
	goto L76
L75:
	;
	v105 = v101
	goto L76
L76:
	;
	goto L73
L77:
	;
	if v27+int32(4320) != v55 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v111 = F__emscripten_memcpy_bulkmem(m, v108, v60, v110)
	mBase = m.M
	v112 = v111
	goto L80
L79:
	;
	v112 = v108
	goto L80
L80:
	;
	goto L77
L81:
	;
	F_pfree(m, v55)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L66
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(v75) <= base.Ui32(v90) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	v217 = v94
	v220 = int32(1)
	goto L58
L86:
	;
	goto L87
L87:
	;
	v132 = v103 + v105 - int32(16)
	v134 = v94 + v90<<(uint(int32(1))%32) - int32(2)
	v136 = v94
	v138 = v75
	v140 = v112
	v141 = v105
	v142 = v110 + v112 - int32(4)
	goto L61
L88:
	;
	if v48 == int32(-2) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v171 = v170 + v147
	if base.Ui32(int32(1293)) < base.Ui32(v171) {
		v232 = v169
		goto L57
	} else {
		goto L98
	}
L90:
	;
	v156 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L66
	} else {
		goto L93
	}
L91:
	;
	v158 = v48
	goto L92
L92:
	;
	if v158 <= int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v158 = v156
	goto L92
L94:
	;
	v161 = int32(0)
	v169 = v161
	v170 = v161
	goto L89
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(int32(385)) < base.Ui32(v158) {
		v169 = v158
		v170 = int32(2)
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+uint32(_consts[1360]))))
	v169 = v158
	v170 = v168
	goto L89
L98:
	;
	v175 = v171 << (uint(int32(1)) % 32)
	v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[1361]))))
	if v170 != v178 {
		v232 = v169
		goto L57
	} else {
		goto L99
	}
L99:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[1362]))))
	v183 = base.I32_extend16_s(v182)
	if v183 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if v182 == int32(0) {
		goto L55
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v183 != int32(3) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if v182 == int32(65373) {
		goto L55
	} else {
		goto L104
	}
L104:
	;
	v241 = v169
	v242 = int32(0) - v183
	goto L56
L105:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	*(*int64)(unsafe.Add(mBase, uint32(v132)+16)) = v194
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	*(*int64)(unsafe.Add(mBase, uint32(v132)+24)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v200
	if v169 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	v217 = v136
	v220 = int32(0)
	goto L58
L108:
	;
	v204 = int32(-2)
	goto L110
L109:
	;
	v204 = int32(0)
	goto L110
L110:
	;
	v10212 = v183
	v10214 = v132 + int32(16)
	v10218 = v204
	v10223 = v134
	v10226 = v142 + int32(4)
	goto L4
L111:
	;
	F_pfree(m, v217)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L66
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	m.G0 = v27 + int32(4816)
	return v220
L114:
	;
	goto L113
L115:
	;
	v241 = v232
	v242 = v237
	goto L56
L116:
	;
	v261 = v260
	goto L118
L117:
	;
	v261 = v142
	goto L118
L118:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	switch v242 - int32(2) {
	case 0:
		goto L274
	default:
		v10149 = v241
		goto L5
	case 3:
		goto L273
	case 4:
		goto L272
	case 5:
		goto L271
	case 6:
		goto L270
	case 7:
		goto L269
	case 8:
		goto L268
	case 9:
		goto L267
	case 12:
		goto L266
	case 13:
		goto L265
	case 14:
		goto L264
	case 15:
		goto L263
	case 16:
		goto L262
	case 21:
		goto L261
	case 22:
		goto L260
	case 23:
		goto L259
	case 24:
		goto L258
	case 25:
		goto L257
	case 26:
		goto L256
	case 27:
		goto L255
	case 28:
		goto L254
	case 29:
		goto L253
	case 30:
		goto L252
	case 31:
		goto L251
	case 32:
		goto L250
	case 33:
		goto L249
	case 34:
		goto L248
	case 37:
		goto L247
	case 38:
		goto L246
	case 39:
		goto L245
	case 40:
		goto L244
	case 41:
		goto L243
	case 42:
		goto L242
	case 43:
		goto L241
	case 44:
		goto L240
	case 45:
		goto L239
	case 46:
		goto L238
	case 47:
		goto L237
	case 48:
		goto L236
	case 49:
		goto L235
	case 50:
		goto L234
	case 51:
		goto L233
	case 52:
		goto L232
	case 57:
		goto L231
	case 58:
		goto L230
	case 59:
		goto L229
	case 60:
		goto L228
	case 61:
		goto L227
	case 62:
		goto L226
	case 63:
		goto L225
	case 64:
		goto L224
	case 65:
		goto L223
	case 66:
		goto L222
	case 67:
		goto L221
	case 68:
		goto L220
	case 69:
		goto L219
	case 70:
		goto L218
	case 71:
		goto L217
	case 72:
		goto L216
	case 73:
		goto L215
	case 74:
		goto L214
	case 75:
		goto L213
	case 76:
		goto L212
	case 77:
		goto L211
	case 78:
		goto L210
	case 79:
		goto L209
	case 80:
		goto L208
	case 81:
		goto L207
	case 82:
		goto L206
	case 83:
		goto L205
	case 84:
		goto L204
	case 85:
		goto L203
	case 86:
		goto L202
	case 87:
		goto L201
	case 88:
		goto L200
	case 89:
		goto L199
	case 90:
		goto L198
	case 91:
		goto L197
	case 92:
		goto L196
	case 93:
		goto L195
	case 94:
		goto L194
	case 95:
		goto L193
	case 96:
		goto L192
	case 97:
		goto L191
	case 98:
		goto L190
	case 99:
		goto L189
	case 100:
		goto L188
	case 101:
		goto L187
	case 102:
		goto L186
	case 103:
		goto L185
	case 104:
		goto L184
	case 105:
		goto L183
	case 106:
		goto L182
	case 107:
		goto L181
	case 108:
		goto L180
	case 109:
		goto L179
	case 110:
		goto L178
	case 111:
		goto L177
	case 112:
		goto L176
	case 113:
		goto L175
	case 114:
		goto L174
	case 115:
		goto L173
	case 116:
		goto L172
	case 117:
		goto L171
	case 118:
		goto L170
	case 119:
		goto L169
	case 120:
		goto L168
	case 121:
		goto L167
	case 122:
		goto L166
	case 123:
		goto L165
	case 124:
		goto L164
	case 125:
		goto L163
	case 126:
		goto L162
	case 127:
		goto L161
	case 128:
		goto L160
	case 129:
		goto L159
	case 130:
		goto L158
	case 131:
		goto L157
	case 132:
		goto L156
	case 133:
		goto L155
	case 134:
		goto L154
	case 135:
		goto L153
	case 136:
		goto L152
	case 137:
		goto L151
	case 138:
		goto L150
	case 139:
		goto L149
	case 140:
		goto L148
	case 141:
		goto L147
	case 142:
		goto L146
	case 143:
		goto L145
	case 144:
		goto L144
	case 145:
		goto L143
	case 146:
		goto L142
	case 147:
		goto L141
	case 148:
		goto L140
	case 149:
		goto L139
	case 150:
		goto L138
	case 151:
		goto L137
	case 152:
		goto L136
	case 153:
		goto L135
	case 154:
		goto L134
	case 155:
		goto L133
	case 156:
		goto L132
	case 157:
		goto L131
	case 158:
		goto L130
	case 159:
		goto L129
	case 160:
		goto L128
	case 161:
		goto L127
	case 162:
		goto L126
	case 163:
		goto L125
	case 164:
		goto L124
	case 165:
		goto L123
	case 166:
		goto L122
	case 167:
		goto L121
	case 168:
		goto L120
	case 169:
		goto L119
	}
L119:
	;
	v9176 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v9176 == int32(0) {
		goto L13
	} else {
		goto L2488
	}
L120:
	;
	v9172 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v9173 = F_pstrdup(m, v9172)
	mBase = m.M
	v9174 = m.ExcPending
	if v9174 != 0 {
		goto L66
	} else {
		goto L2487
	}
L121:
	;
	v9170 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9170
	v10149 = v241
	goto L5
L122:
	;
	v9168 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9168
	v10149 = v241
	goto L5
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L124:
	;
	v9164 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9164
	v10149 = v241
	goto L5
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L126:
	;
	v9155 = v132 - int32(16)
	v9156 = *(*int32)(unsafe.Add(mBase, uint32(v9155)))
	F_plpgsql_ns_push(m, v9156, int32(1))
	mBase = m.M
	v9159 = m.ExcPending
	if v9159 != 0 {
		goto L66
	} else {
		goto L2486
	}
L127:
	;
	F_plpgsql_ns_push(m, int32(0), int32(1))
	mBase = m.M
	v9151 = m.ExcPending
	if v9151 != 0 {
		goto L66
	} else {
		goto L2485
	}
L128:
	;
	v9141 = v132 - int32(16)
	v9142 = *(*int32)(unsafe.Add(mBase, uint32(v9141)))
	F_plpgsql_ns_push(m, v9142, int32(0))
	mBase = m.M
	v9145 = m.ExcPending
	if v9145 != 0 {
		goto L66
	} else {
		goto L2484
	}
L129:
	;
	v9134 = int32(0)
	F_plpgsql_ns_push(m, v9134, v9134)
	mBase = m.M
	v9137 = m.ExcPending
	if v9137 != 0 {
		goto L66
	} else {
		goto L2483
	}
L130:
	;
	v9119 = int32(0)
	v9123 = int32(1)
	v9131 = F_read_sql_construct(m, int32(336), v9119, v9119, int32(526588), int32(2), v9123, v9123, v9119, v9119, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9132 = m.ExcPending
	if v9132 != 0 {
		goto L66
	} else {
		goto L2482
	}
L131:
	;
	v9103 = int32(0)
	v9107 = int32(1)
	v9115 = F_read_sql_construct(m, int32(376), v9103, v9103, int32(530605), int32(2), v9107, v9107, v9103, v9103, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9116 = m.ExcPending
	if v9116 != 0 {
		goto L66
	} else {
		goto L2481
	}
L132:
	;
	v9087 = int32(0)
	v9091 = int32(1)
	v9099 = F_read_sql_construct(m, int32(59), v9087, v9087, int32(546849), int32(2), v9091, v9091, v9087, v9087, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9100 = m.ExcPending
	if v9100 != 0 {
		goto L66
	} else {
		goto L2480
	}
L133:
	;
	v8913 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v8914 = int32(351311)
	v8917 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1366])))
	v8918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8913))))
	if v8918 == int32(0) {
		v8937 = v8917
		v8938 = v8918
		goto L2443
	} else {
		goto L2444
	}
L134:
	;
	v8911 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8911
	v10149 = v241
	goto L5
L135:
	;
	v8880 = v132 - int32(32)
	v8881 = *(*int32)(unsafe.Add(mBase, uint32(v8880)))
	v8885 = v8881
	goto L2439
L136:
	;
	v8820 = F_palloc0(m, int32(12))
	mBase = m.M
	v8821 = m.ExcPending
	if v8821 != 0 {
		goto L66
	} else {
		goto L2424
	}
L137:
	;
	v8810 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+272)) = v8810
	*(*int32)(unsafe.Add(mBase, uint32(v27)+276)) = v8810
	v8816 = F_list_make1_impl(m, int32(1), v27+int32(272))
	mBase = m.M
	v8817 = m.ExcPending
	if v8817 != 0 {
		goto L66
	} else {
		goto L2423
	}
L138:
	;
	v8805 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8806 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v8807 = F_lappend(m, v8805, v8806)
	mBase = m.M
	v8808 = m.ExcPending
	if v8808 != 0 {
		goto L66
	} else {
		goto L2422
	}
L139:
	;
	v8799 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8800 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v8799)+8)) = v8800
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8799
	v10149 = v241
	goto L5
L140:
	;
	v8712 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v8713 = int32(0)
	if v8712 < v8713 {
		v8756 = v8713
		goto L2404
	} else {
		goto L2405
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L142:
	;
	v8707 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v8707, l1)
	mBase = m.M
	v8709 = m.ExcPending
	if v8709 != 0 {
		goto L66
	} else {
		goto L2402
	}
L143:
	;
	v8704 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v8704, l1)
	mBase = m.M
	v8706 = m.ExcPending
	if v8706 != 0 {
		goto L66
	} else {
		goto L2401
	}
L144:
	;
	v8692 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v8693 = *(*int32)(unsafe.Add(mBase, uint32(v8692)))
	if v8693 != 0 {
		goto L18
	} else {
		goto L2397
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(1)
	v10149 = v241
	goto L5
L148:
	;
	v8618 = F_palloc(m, int32(16))
	mBase = m.M
	v8619 = m.ExcPending
	if v8619 != 0 {
		goto L66
	} else {
		goto L2382
	}
L149:
	;
	v8549 = F_palloc(m, int32(16))
	mBase = m.M
	v8550 = m.ExcPending
	if v8550 != 0 {
		goto L66
	} else {
		goto L2367
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L151:
	;
	v8479 = F_palloc(m, int32(16))
	mBase = m.M
	v8480 = m.ExcPending
	if v8480 != 0 {
		goto L66
	} else {
		goto L2352
	}
L152:
	;
	v8070 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = uint8(v8070)
	v8073 = F_palloc0(m, int32(36))
	mBase = m.M
	v8074 = m.ExcPending
	if v8074 != 0 {
		goto L66
	} else {
		goto L2230
	}
L153:
	;
	v8011 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	v8014 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v8015 = int32(0)
	if v8014 < v8015 {
		v8058 = v8015
		goto L2217
	} else {
		goto L2218
	}
L154:
	;
	v7928 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	F_read_into_target(m, v27+int32(4784), int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7937 = m.ExcPending
	if v7937 != 0 {
		goto L66
	} else {
		goto L2198
	}
L155:
	;
	v7611 = F_palloc0(m, int32(36))
	mBase = m.M
	v7612 = m.ExcPending
	if v7612 != 0 {
		goto L66
	} else {
		goto L2123
	}
L156:
	;
	v7426 = int32(1)
	v7435 = F_read_sql_construct(m, int32(332), int32(381), int32(59), int32(546825), int32(2), v7426, v7426, int32(0), v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7436 = m.ExcPending
	if v7436 != 0 {
		goto L66
	} else {
		goto L2088
	}
L157:
	;
	v7397 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7398 = m.ExcPending
	if v7398 != 0 {
		goto L66
	} else {
		goto L2081
	}
L158:
	;
	v7370 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7371 = m.ExcPending
	if v7371 != 0 {
		goto L66
	} else {
		goto L2074
	}
L159:
	;
	v7357 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7363 = F_make_execsql_stmt(m, int32(337), v7357, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7364 = m.ExcPending
	if v7364 != 0 {
		goto L66
	} else {
		goto L2073
	}
L160:
	;
	v7347 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7353 = F_make_execsql_stmt(m, int32(331), v7347, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7354 = m.ExcPending
	if v7354 != 0 {
		goto L66
	} else {
		goto L2072
	}
L161:
	;
	v7337 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7343 = F_make_execsql_stmt(m, int32(328), v7337, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7344 = m.ExcPending
	if v7344 != 0 {
		goto L66
	} else {
		goto L2071
	}
L162:
	;
	v7326 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7326
	v7330 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v7330
	v7334 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v7334
	v10149 = v241
	goto L5
L163:
	;
	v7226 = F_palloc(m, int32(20))
	mBase = m.M
	v7227 = m.ExcPending
	if v7227 != 0 {
		goto L66
	} else {
		goto L2051
	}
L164:
	;
	v6195 = F_palloc(m, int32(32))
	mBase = m.M
	v6196 = m.ExcPending
	if v6196 != 0 {
		goto L66
	} else {
		goto L1747
	}
L165:
	;
	v5604 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5605 = m.ExcPending
	if v5605 != 0 {
		goto L66
	} else {
		goto L1607
	}
L166:
	;
	v5598 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v5598)
	v10149 = v241
	goto L5
L167:
	;
	v5596 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v5596)
	v10149 = v241
	goto L5
L168:
	;
	v5464 = F_palloc0(m, int32(24))
	mBase = m.M
	v5465 = m.ExcPending
	if v5465 != 0 {
		goto L66
	} else {
		goto L1554
	}
L169:
	;
	v5461 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5461
	v10149 = v241
	goto L5
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L171:
	;
	v5347 = F_palloc0(m, int32(32))
	mBase = m.M
	v5348 = m.ExcPending
	if v5348 != 0 {
		goto L66
	} else {
		goto L1526
	}
L172:
	;
	v5343 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v5343, l1)
	mBase = m.M
	v5345 = m.ExcPending
	if v5345 != 0 {
		goto L66
	} else {
		goto L1525
	}
L173:
	;
	v5274 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5274
	v5276 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5277 = int32(0)
	if v5276 < v5277 {
		v5320 = v5277
		goto L1508
	} else {
		goto L1509
	}
L174:
	;
	v5183 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v5183 != 0 {
		goto L1482
	} else {
		goto L1483
	}
L175:
	;
	v4798 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4799 = m.ExcPending
	if v4799 != 0 {
		goto L66
	} else {
		goto L1412
	}
L176:
	;
	v4713 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v4714 = *(*int32)(unsafe.Add(mBase, uint32(v4713)))
	v4717 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4718 = int32(0)
	if v4717 < v4718 {
		v4761 = v4718
		goto L1388
	} else {
		goto L1389
	}
L177:
	;
	v4623 = F_palloc0(m, int32(24))
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		goto L66
	} else {
		goto L1364
	}
L178:
	;
	v4538 = F_palloc0(m, int32(20))
	mBase = m.M
	v4539 = m.ExcPending
	if v4539 != 0 {
		goto L66
	} else {
		goto L1341
	}
L179:
	;
	v4525 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v4525 != 0 {
		goto L1337
	} else {
		goto L1338
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L181:
	;
	v4464 = F_palloc(m, int32(12))
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L66
	} else {
		goto L1322
	}
L182:
	;
	v4454 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v4454
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v4454
	v4460 = F_list_make1_impl(m, int32(1), v27+int32(216))
	mBase = m.M
	v4461 = m.ExcPending
	if v4461 != 0 {
		goto L66
	} else {
		goto L1321
	}
L183:
	;
	v4449 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v4450 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v4451 = F_lappend(m, v4449, v4450)
	mBase = m.M
	v4452 = m.ExcPending
	if v4452 != 0 {
		goto L66
	} else {
		goto L1320
	}
L184:
	;
	v4413 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4414 = m.ExcPending
	if v4414 != 0 {
		goto L66
	} else {
		goto L1313
	}
L185:
	;
	v4201 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(80))))
	v4204 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	v4207 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	v4210 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(24))))
	v4212 = F_palloc(m, int32(32))
	mBase = m.M
	v4213 = m.ExcPending
	if v4213 != 0 {
		goto L66
	} else {
		goto L1279
	}
L186:
	;
	v4197 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4197
	v10149 = v241
	goto L5
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L188:
	;
	v4131 = F_palloc0(m, int32(12))
	mBase = m.M
	v4132 = m.ExcPending
	if v4132 != 0 {
		goto L66
	} else {
		goto L1263
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L190:
	;
	v4050 = F_palloc0(m, int32(28))
	mBase = m.M
	v4051 = m.ExcPending
	if v4051 != 0 {
		goto L66
	} else {
		goto L1248
	}
L191:
	;
	v4046 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v4046, l1)
	mBase = m.M
	v4048 = m.ExcPending
	if v4048 != 0 {
		goto L66
	} else {
		goto L1247
	}
L192:
	;
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v4043, l1)
	mBase = m.M
	v4045 = m.ExcPending
	if v4045 != 0 {
		goto L66
	} else {
		goto L1246
	}
L193:
	;
	v4027 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v4028 = *(*int32)(unsafe.Add(mBase, uint32(v4027)))
	if base.Ui32(v4028-int32(1)) < base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L1242
	}
L194:
	;
	v3617 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3618 = m.ExcPending
	if v3618 != 0 {
		goto L66
	} else {
		goto L1115
	}
L195:
	;
	v3603 = F_palloc(m, int32(8))
	mBase = m.M
	v3604 = m.ExcPending
	if v3604 != 0 {
		goto L66
	} else {
		goto L1087
	}
L196:
	;
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+172)) = v3593
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v3593
	v3599 = F_list_make1_impl(m, int32(1), v27+int32(172))
	mBase = m.M
	v3600 = m.ExcPending
	if v3600 != 0 {
		goto L66
	} else {
		goto L1086
	}
L197:
	;
	v3588 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3590 = F_lappend(m, v3588, v3589)
	mBase = m.M
	v3591 = m.ExcPending
	if v3591 != 0 {
		goto L66
	} else {
		goto L1085
	}
L198:
	;
	v3584 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v3584)
	v10149 = v241
	goto L5
L199:
	;
	v3582 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v3582)
	v10149 = v241
	goto L5
L200:
	;
	v3580 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v3580)
	v10149 = v241
	goto L5
L201:
	;
	v3355 = F_palloc0(m, int32(20))
	mBase = m.M
	v3356 = m.ExcPending
	if v3356 != 0 {
		goto L66
	} else {
		goto L1032
	}
L202:
	;
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v3239 == int32(0) {
		goto L1006
	} else {
		goto L1007
	}
L203:
	;
	v3149 = F_palloc0(m, int32(24))
	mBase = m.M
	v3150 = m.ExcPending
	if v3150 != 0 {
		goto L66
	} else {
		goto L989
	}
L204:
	;
	v3059 = F_palloc0(m, int32(24))
	mBase = m.M
	v3060 = m.ExcPending
	if v3060 != 0 {
		goto L66
	} else {
		goto L972
	}
L205:
	;
	v2775 = F_palloc0(m, int32(16))
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
		goto L66
	} else {
		goto L905
	}
L206:
	;
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2772
	v10149 = v241
	goto L5
L207:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2770
	v10149 = v241
	goto L5
L208:
	;
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2768
	v10149 = v241
	goto L5
L209:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2766
	v10149 = v241
	goto L5
L210:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2764
	v10149 = v241
	goto L5
L211:
	;
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2762
	v10149 = v241
	goto L5
L212:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2760
	v10149 = v241
	goto L5
L213:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2758
	v10149 = v241
	goto L5
L214:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2756
	v10149 = v241
	goto L5
L215:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2754
	v10149 = v241
	goto L5
L216:
	;
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2752
	v10149 = v241
	goto L5
L217:
	;
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2750
	v10149 = v241
	goto L5
L218:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2748
	v10149 = v241
	goto L5
L219:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2746
	v10149 = v241
	goto L5
L220:
	;
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2744
	v10149 = v241
	goto L5
L221:
	;
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2742
	v10149 = v241
	goto L5
L222:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2740
	v10149 = v241
	goto L5
L223:
	;
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2738
	v10149 = v241
	goto L5
L224:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2736
	v10149 = v241
	goto L5
L225:
	;
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2734
	v10149 = v241
	goto L5
L226:
	;
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2732
	v10149 = v241
	goto L5
L227:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2730
	v10149 = v241
	goto L5
L228:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2728
	v10149 = v241
	goto L5
L229:
	;
	v2726 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2726
	v10149 = v241
	goto L5
L230:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v2717 == int32(0) {
		goto L901
	} else {
		goto L902
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L232:
	;
	v2697 = int32(0)
	v2701 = int32(1)
	v2709 = F_read_sql_construct(m, int32(59), v2697, v2697, int32(546849), int32(2), v2701, v2701, v2697, v2697, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		goto L66
	} else {
		goto L900
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L234:
	;
	v2692 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v2692)
	v10149 = v241
	goto L5
L235:
	;
	v2690 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v2690)
	v10149 = v241
	goto L5
L236:
	;
	v2685 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v2687 = F_get_collation_oid(m, v2685, int32(0))
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L66
	} else {
		goto L899
	}
L237:
	;
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v2670 = F_pstrdup(m, v2669)
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L66
	} else {
		goto L895
	}
L238:
	;
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v2656 = F_makeString(m, v2655)
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L66
	} else {
		goto L892
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L240:
	;
	if v241 == int32(-2) {
		goto L740
	} else {
		goto L741
	}
L241:
	;
	v2214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v2214)
	v10149 = v241
	goto L5
L242:
	;
	v2212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v2212)
	v10149 = v241
	goto L5
L243:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v1827 = F_pstrdup(m, v1826)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L66
	} else {
		goto L639
	}
L244:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1442
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v1445 = int32(0)
	if v1444 < v1445 {
		v1488 = v1445
		goto L540
	} else {
		goto L541
	}
L245:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v1132 == int32(0) {
		goto L49
	} else {
		goto L460
	}
L246:
	;
	v985 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v985 == int32(0) {
		goto L426
	} else {
		goto L427
	}
L247:
	;
	v837 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v837 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L248:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v833 = F_plpgsql_build_variable(m, v827, v830, v831, int32(1))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L66
	} else {
		goto L383
	}
L249:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v822 = F_lappend(m, v820, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L66
	} else {
		goto L382
	}
L250:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v809
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v809
	v815 = F_list_make1_impl(m, int32(1), v27+int32(44))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L66
	} else {
		goto L381
	}
L251:
	;
	v632 = F_palloc0(m, int32(40))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L66
	} else {
		goto L350
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L253:
	;
	v614 = int32(0)
	v626 = F_read_sql_construct(m, int32(59), v614, v614, int32(546849), v614, v614, int32(1), v614, v614, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L66
	} else {
		goto L349
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(2)
	v10149 = v241
	goto L5
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(4)
	v10149 = v241
	goto L5
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L257:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if v569 != 0 {
		goto L338
	} else {
		goto L339
	}
L258:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	F_plpgsql_ns_push(m, v562, int32(2))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L66
	} else {
		goto L336
	}
L259:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	F_plpgsql_ns_additem(m, v553, v554, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L66
	} else {
		goto L335
	}
L260:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	if v502 != 0 {
		goto L324
	} else {
		goto L325
	}
L261:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L66
	} else {
		goto L319
	}
L262:
	;
	v474 = F_plpgsql_add_initdatums(m, int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L66
	} else {
		goto L318
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = int32(0)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v468
	v470 = F_plpgsql_add_initdatums(m, v30)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L66
	} else {
		goto L317
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+308)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v459
	v10149 = v241
	goto L5
L265:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+308)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v450
	v10149 = v241
	goto L5
L266:
	;
	v349 = F_palloc0(m, int32(32))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L66
	} else {
		goto L294
	}
L267:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v345 = F_pstrdup(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L66
	} else {
		goto L293
	}
L268:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v342
	v10149 = v241
	goto L5
L269:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+488)) = int32(2)
	v10149 = v241
	goto L5
L270:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+488)) = int32(1)
	v10149 = v241
	goto L5
L271:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+488)) = int32(0)
	v10149 = v241
	goto L5
L272:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v273 != int32(111) {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1369])) = uint8(v270)
	v10149 = v241
	goto L5
L274:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v267
	v10149 = v241
	goto L5
L275:
	;
	v284 = int32(338833)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v288 == int32(0) {
		v307 = v287
		v308 = v288
		goto L280
	} else {
		goto L281
	}
L276:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	if v276 != int32(110) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+2)))
	if v279 != 0 {
		goto L275
	} else {
		goto L278
	}
L278:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+492)) = uint8(v282)
	v10149 = v241
	goto L5
L279:
	;
	if v308-v307 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L280:
	;
	goto L279
L281:
	;
	if v287 != v288 {
		v307 = v287
		v308 = v288
		goto L280
	} else {
		goto L282
	}
L282:
	;
	v292 = v272
	v293 = v284
	goto L283
L283:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+1)))
	if v297 == int32(0) {
		v307 = v296
		v308 = v297
		goto L280
	} else {
		goto L285
	}
L284:
	;
	v307 = v296
	v308 = v297
	goto L280
L285:
	;
	v300 = int32(1)
	if v296 == v297 {
		v292 = v292 + v300
		v293 = v293 + v300
		goto L283
	} else {
		goto L286
	}
L286:
	;
	goto L284
L287:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v313)+492)) = uint8(v314)
	v10149 = v241
	goto L5
L288:
	;
	goto L289
L289:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L66
	} else {
		goto L290
	}
L290:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v320
	F_errmsg_internal(m, int32(183310), v27)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L66
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(26959), int32(396), int32(360630))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L66
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v345
	v10149 = v241
	goto L5
L294:
	;
	v351 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = v351
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(16))))
	if v355 < v351 {
		v399 = v351
		goto L296
	} else {
		goto L297
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+4)) = v399
	v404 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+520))
	v407 = v405 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v404)+520)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v407
	v411 = v132 - int32(80)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+12)) = v412
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(76))))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+20)) = v416
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(72))))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+24)) = v420
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+16)) = v424
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+28)) = v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_check_labels(m, v430, v431, v432, l1)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L66
	} else {
		goto L309
	}
L296:
	;
	goto L295
L297:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+60))
	if v362 == int32(0) {
		v399 = v351
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v365 = v355 + v362
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v361)+188))
	if base.Ui32(v366) <= base.Ui32(v365) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v361)+196))
	if v375 == int32(0) {
		v399 = v376
		goto L296
	} else {
		goto L303
	}
L300:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v361)+192))
	v375 = v368
	goto L299
L301:
	;
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v361)+188)) = v362
	v373 = F_strchr(m, v362, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v361)+192)) = v373
	v375 = v373
	goto L299
L303:
	;
	if base.Ui32(v365) <= base.Ui32(v375) {
		v399 = v376
		goto L296
	} else {
		goto L304
	}
L304:
	;
	v380 = v375
	v382 = v376
	goto L305
L305:
	;
	v385 = int32(1)
	v386 = v382 + v385
	*(*int32)(unsafe.Add(mBase, uint32(v361)+196)) = v386
	v389 = v380 + v385
	*(*int32)(unsafe.Add(mBase, uint32(v361)+188)) = v389
	v392 = F_strchr(m, v389, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v361)+192)) = v392
	if v392 == int32(0) {
		v399 = v386
		goto L296
	} else {
		goto L307
	}
L306:
	;
	v399 = v386
	goto L296
L307:
	;
	if base.Ui32(v392) < base.Ui32(v365) {
		v380 = v392
		v382 = v386
		goto L305
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if v438 != 0 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v349
	v10149 = v241
	goto L5
L311:
	;
	v439 = v437
	goto L314
L312:
	;
	v442 = v437
	goto L313
L313:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v442)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1350])) = v444
	goto L310
L314:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if v441 != 0 {
		v439 = v440
		goto L314
	} else {
		goto L316
	}
L315:
	;
	v442 = v440
	goto L313
L316:
	;
	goto L315
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v470
	v10149 = v241
	goto L5
L318:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = int32(1)
	v10149 = v241
	goto L5
L319:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L66
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(216159), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L66
	} else {
		goto L321
	}
L321:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v493 = F_plpgsql_scanner_errposition(m, v492, l1)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L66
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(26959), int32(502), int32(360630))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L66
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	v504 = v132 - int32(48)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+16))
	if v506 == int32(0) {
		goto L53
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(80))))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(76))))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	v522 = F_plpgsql_build_variable(m, v514, v517, v520, int32(1))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L66
	} else {
		goto L328
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v505)+16)) = v502
	goto L326
L328:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+int32(-64)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+16)) = uint8(v526)
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132-int32(16)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+17)) = uint8(v530)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+20)) = v532
	if base.B2i32(v532 == int32(0))&base.B2i32(v530 == int32(1)) != 0 {
		goto L52
	} else {
		goto L329
	}
L329:
	;
	if v532 == int32(0) {
		v10149 = v241
		goto L5
	} else {
		goto L330
	}
L330:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if v541 != 0 {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v532)+20)) = uint8(v547)
	*(*int32)(unsafe.Add(mBase, uint32(v532)+16)) = v546
	v10149 = v241
	goto L5
L332:
	;
	v546 = int32(-1)
	v547 = int32(0)
	goto L331
L333:
	;
	goto L334
L334:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	v546 = v544
	v547 = int32(1)
	goto L331
L335:
	;
	v10149 = v241
	goto L5
L336:
	;
	v10149 = v241
	goto L5
L337:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(96))))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(92))))
	v585 = int32(0)
	v587 = F_plpgsql_build_datatype(m, int32(1790), int32(-1), v585, v585)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L66
	} else {
		goto L344
	}
L338:
	;
	v570 = v568
	goto L341
L339:
	;
	v573 = v568
	goto L340
L340:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v573)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1350])) = v575
	goto L337
L341:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)+8))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	if v572 != 0 {
		v570 = v571
		goto L341
	} else {
		goto L343
	}
L342:
	;
	v573 = v571
	goto L340
L343:
	;
	goto L342
L344:
	;
	v590 = F_plpgsql_build_variable(m, v579, v582, v587, int32(1))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L66
	} else {
		goto L345
	}
L345:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v590)+28)) = v592
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	if v596 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	v599 = v597
	goto L348
L347:
	;
	v599 = int32(-1)
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v590)+32)) = v599
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(80))))
	*(*int32)(unsafe.Add(mBase, uint32(v590)+36)) = v603 | int32(256)
	v10149 = v241
	goto L5
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v626
	v10149 = v241
	goto L5
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+8)) = int32(670471)
	*(*int32)(unsafe.Add(mBase, uint32(v632))) = int32(1)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v641 = int32(0)
	if v640 < v641 {
		v684 = v641
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v687 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v632)+24)) = v687
	*(*int32)(unsafe.Add(mBase, uint32(v632)+12)) = v684
	v692 = v132 - int32(16)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	if v693 != 0 {
		goto L365
	} else {
		goto L366
	}
L352:
	;
	goto L351
L353:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)+60))
	if v647 == int32(0) {
		v684 = v641
		goto L352
	} else {
		goto L354
	}
L354:
	;
	v650 = v640 + v647
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v646)+188))
	if base.Ui32(v651) <= base.Ui32(v650) {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v646)+196))
	if v660 == int32(0) {
		v684 = v661
		goto L352
	} else {
		goto L359
	}
L356:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v646)+192))
	v660 = v653
	goto L355
L357:
	;
	goto L358
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v646)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v646)+188)) = v647
	v658 = F_strchr(m, v647, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v646)+192)) = v658
	v660 = v658
	goto L355
L359:
	;
	if base.Ui32(v650) <= base.Ui32(v660) {
		v684 = v661
		goto L352
	} else {
		goto L360
	}
L360:
	;
	v665 = v660
	v667 = v661
	goto L361
L361:
	;
	v670 = int32(1)
	v671 = v667 + v670
	*(*int32)(unsafe.Add(mBase, uint32(v646)+196)) = v671
	v674 = v665 + v670
	*(*int32)(unsafe.Add(mBase, uint32(v646)+188)) = v674
	v677 = F_strchr(m, v674, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v646)+192)) = v677
	if v677 == int32(0) {
		v684 = v671
		goto L352
	} else {
		goto L363
	}
L362:
	;
	v684 = v671
	goto L352
L363:
	;
	if base.Ui32(v677) < base.Ui32(v650) {
		v665 = v677
		v667 = v671
		goto L361
	} else {
		goto L364
	}
L364:
	;
	goto L362
L365:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+4))
	v695 = v694
	goto L367
L366:
	;
	v695 = v687
	goto L367
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+28)) = v695
	v699 = F_palloc(m, v695<<(uint(int32(2))%32))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L66
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+32)) = v699
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v632)+28))
	v705 = F_palloc(m, v702<<(uint(int32(2))%32))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L66
	} else {
		goto L369
	}
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+36)) = v705
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	if v708 != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if int32(0) < v709 {
		goto L373
	} else {
		goto L374
	}
L371:
	;
	v803 = int32(0)
	goto L372
L372:
	;
	F_list_free(m, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L66
	} else {
		goto L379
	}
L373:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v632)+32))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v708)+12))
	v718 = int32(0)
	goto L376
L374:
	;
	goto L375
L375:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	v803 = v777
	goto L372
L376:
	;
	v740 = v718 << (uint(int32(2)) % 32)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v740+v713)))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v712+v740))) = v744
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v740+v705))) = v747
	v750 = v718 + int32(1)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if v750 < v751 {
		v718 = v750
		goto L376
	} else {
		goto L378
	}
L377:
	;
	goto L375
L378:
	;
	goto L377
L379:
	;
	F_plpgsql_adddatum(m, v632)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L66
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v632
	v10149 = v241
	goto L5
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v815
	v10149 = v241
	goto L5
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v822
	v10149 = v241
	goto L5
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v833
	v10149 = v241
	goto L5
L384:
	;
	if v970 == int32(0) {
		goto L51
	} else {
		goto L421
	}
L385:
	;
	goto L384
L386:
	;
	v970 = v864
	goto L385
L388:
	;
	v970 = int32(0)
	goto L385
L389:
	;
	v850 = v837
	goto L390
L390:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v850)))
	if v859 != 0 {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	goto L388
L392:
	;
	v864 = v850
	v866 = v859
	goto L395
L393:
	;
	v886 = v850
	goto L394
L394:
	;
	goto L402
L395:
	;
	v871 = F_strcmp(m, v864+int32(12), v839)
	mBase = m.M
	if v871 != 0 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v886 = v880
	goto L394
L397:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v864)+8))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v880)))
	if v881 != 0 {
		v864 = v880
		v866 = v881
		goto L395
	} else {
		goto L401
	}
L398:
	;
	if base.B2i32(v866 == int32(1))&int32(0) != 0 {
		goto L397
	} else {
		goto L399
	}
L399:
	;
	goto L386
L401:
	;
	goto L396
L402:
	;
	goto L418
L418:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v886)+8))
	if v928 != 0 {
		v850 = v928
		goto L390
	} else {
		goto L419
	}
L419:
	;
	goto L391
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v970
	v10149 = v241
	goto L5
L422:
	;
	if v1118 == int32(0) {
		goto L50
	} else {
		goto L459
	}
L423:
	;
	goto L422
L424:
	;
	v1118 = v1012
	goto L423
L426:
	;
	v1118 = int32(0)
	goto L423
L427:
	;
	v998 = v985
	goto L428
L428:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	if v1007 != 0 {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	goto L426
L430:
	;
	v1012 = v998
	v1014 = v1007
	goto L433
L431:
	;
	v1034 = v998
	goto L432
L432:
	;
	goto L440
L433:
	;
	v1019 = F_strcmp(m, v1012+int32(12), v987)
	mBase = m.M
	if v1019 != 0 {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	v1034 = v1028
	goto L432
L435:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+8))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1028)))
	if v1029 != 0 {
		v1012 = v1028
		v1014 = v1029
		goto L433
	} else {
		goto L439
	}
L436:
	;
	if base.B2i32(v1014 == int32(1))&int32(0) != 0 {
		goto L435
	} else {
		goto L437
	}
L437:
	;
	goto L424
L439:
	;
	goto L434
L440:
	;
	goto L456
L456:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+8))
	if v1076 != 0 {
		v998 = v1076
		goto L428
	} else {
		goto L457
	}
L457:
	;
	goto L429
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1118
	v10149 = v241
	goto L5
L460:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+4))
	switch v1135 - int32(2) {
	case 0:
		goto L463
	case 1:
		goto L462
	default:
		goto L49
	}
L461:
	;
	if v1438 == int32(0) {
		goto L49
	} else {
		goto L538
	}
L462:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+12))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+4))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+4))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+8))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+4))
	if v1288 == int32(0) {
		goto L505
	} else {
		goto L506
	}
L463:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1141)+12))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+4))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+4))
	if v1139 == int32(0) {
		goto L468
	} else {
		goto L469
	}
L464:
	;
	v1438 = v1286
	goto L461
L465:
	;
	v1286 = v1276
	goto L464
L466:
	;
	v1276 = v1170
	goto L465
L468:
	;
	v1276 = int32(0)
	goto L465
L469:
	;
	v1156 = v1139
	goto L470
L470:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1156)))
	if v1165 != 0 {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	goto L468
L472:
	;
	v1170 = v1156
	v1172 = v1165
	goto L475
L473:
	;
	v1192 = v1156
	goto L474
L474:
	;
	if v1146 == int32(0) {
		v1229 = v1192
		goto L482
	} else {
		goto L483
	}
L475:
	;
	v1177 = F_strcmp(m, v1170+int32(12), v1144)
	mBase = m.M
	if v1177 != 0 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v1192 = v1186
	goto L474
L477:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+8))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)))
	if v1187 != 0 {
		v1170 = v1186
		v1172 = v1187
		goto L475
	} else {
		goto L481
	}
L478:
	;
	if base.B2i32(v1172 == int32(1))&base.B2i32(v1146 != int32(0)) != 0 {
		goto L477
	} else {
		goto L479
	}
L479:
	;
	goto L466
L481:
	;
	goto L476
L482:
	;
	goto L498
L483:
	;
	v1201 = F_strcmp(m, v1192+int32(12), v1144)
	mBase = m.M
	if v1201 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1202 = v1192
	goto L486
L485:
	;
	v1202 = v1156
	goto L486
L486:
	;
	if v1201 != 0 {
		v1229 = v1202
		goto L482
	} else {
		goto L487
	}
L487:
	;
	if v1165 == int32(0) {
		v1229 = v1202
		goto L482
	} else {
		goto L488
	}
L488:
	;
	v1205 = v1156
	v1212 = v1165
	goto L489
L489:
	;
	v1216 = F_strcmp(m, v1205+int32(12), v1146)
	mBase = m.M
	if v1216 != 0 {
		goto L491
	} else {
		goto L492
	}
L490:
	;
	v1229 = v1223
	goto L482
L491:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+8))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1223)))
	if v1224 != 0 {
		v1205 = v1223
		v1212 = v1224
		goto L489
	} else {
		goto L497
	}
L492:
	;
	if int32(0)&base.B2i32(v1212 == int32(1)) != 0 {
		goto L491
	} else {
		goto L493
	}
L493:
	;
	goto L494
L494:
	;
	v1286 = v1205
	goto L464
L497:
	;
	goto L490
L498:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+8))
	if v1234 != 0 {
		v1156 = v1234
		goto L470
	} else {
		goto L499
	}
L499:
	;
	goto L471
L501:
	;
	v1438 = v1436
	goto L461
L502:
	;
	v1436 = v1426
	goto L501
L503:
	;
	v1426 = v1320
	goto L502
L505:
	;
	v1426 = int32(0)
	goto L502
L506:
	;
	v1306 = v1288
	goto L507
L507:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1306)))
	if v1315 != 0 {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	goto L505
L509:
	;
	v1320 = v1306
	v1322 = v1315
	goto L512
L510:
	;
	v1342 = v1306
	goto L511
L511:
	;
	if v1295 == int32(0) {
		v1379 = v1342
		goto L519
	} else {
		goto L520
	}
L512:
	;
	v1327 = F_strcmp(m, v1320+int32(12), v1293)
	mBase = m.M
	if v1327 != 0 {
		goto L514
	} else {
		goto L515
	}
L513:
	;
	v1342 = v1336
	goto L511
L514:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+8))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1336)))
	if v1337 != 0 {
		v1320 = v1336
		v1322 = v1337
		goto L512
	} else {
		goto L518
	}
L515:
	;
	if base.B2i32(v1322 == int32(1))&base.B2i32(v1295 != int32(0)) != 0 {
		goto L514
	} else {
		goto L516
	}
L516:
	;
	goto L503
L518:
	;
	goto L513
L519:
	;
	goto L535
L520:
	;
	v1351 = F_strcmp(m, v1342+int32(12), v1293)
	mBase = m.M
	if v1351 != 0 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v1352 = v1342
	goto L523
L522:
	;
	v1352 = v1306
	goto L523
L523:
	;
	if v1351 != 0 {
		v1379 = v1352
		goto L519
	} else {
		goto L524
	}
L524:
	;
	if v1315 == int32(0) {
		v1379 = v1352
		goto L519
	} else {
		goto L525
	}
L525:
	;
	v1355 = v1306
	v1362 = v1315
	goto L526
L526:
	;
	v1366 = F_strcmp(m, v1355+int32(12), v1295)
	mBase = m.M
	if v1366 != 0 {
		goto L528
	} else {
		goto L529
	}
L527:
	;
	v1379 = v1373
	goto L519
L528:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+8))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	if v1374 != 0 {
		v1355 = v1373
		v1362 = v1374
		goto L526
	} else {
		goto L534
	}
L529:
	;
	if base.B2i32(v1297 != int32(0))&base.B2i32(v1362 == int32(1)) != 0 {
		goto L528
	} else {
		goto L530
	}
L530:
	;
	goto L531
L531:
	;
	v1436 = v1355
	goto L501
L534:
	;
	goto L527
L535:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1379)+8))
	if v1384 != 0 {
		v1306 = v1384
		goto L507
	} else {
		goto L536
	}
L536:
	;
	goto L508
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1438
	v10149 = v241
	goto L5
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1488
	v1493 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v1493 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L540:
	;
	goto L539
L541:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+60))
	if v1451 == int32(0) {
		v1488 = v1445
		goto L540
	} else {
		goto L542
	}
L542:
	;
	v1454 = v1444 + v1451
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+188))
	if base.Ui32(v1455) <= base.Ui32(v1454) {
		goto L544
	} else {
		goto L545
	}
L543:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+196))
	if v1464 == int32(0) {
		v1488 = v1465
		goto L540
	} else {
		goto L547
	}
L544:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+192))
	v1464 = v1457
	goto L543
L545:
	;
	goto L546
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+188)) = v1451
	v1462 = F_strchr(m, v1451, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+192)) = v1462
	v1464 = v1462
	goto L543
L547:
	;
	if base.Ui32(v1454) <= base.Ui32(v1464) {
		v1488 = v1465
		goto L540
	} else {
		goto L548
	}
L548:
	;
	v1469 = v1464
	v1471 = v1465
	goto L549
L549:
	;
	v1474 = int32(1)
	v1475 = v1471 + v1474
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+196)) = v1475
	v1478 = v1469 + v1474
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+188)) = v1478
	v1481 = F_strchr(m, v1478, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+192)) = v1481
	if v1481 == int32(0) {
		v1488 = v1475
		goto L540
	} else {
		goto L551
	}
L550:
	;
	v1488 = v1475
	goto L540
L551:
	;
	if base.Ui32(v1481) < base.Ui32(v1454) {
		v1469 = v1481
		v1471 = v1475
		goto L549
	} else {
		goto L552
	}
L552:
	;
	goto L550
L553:
	;
	if v1626 != 0 {
		goto L48
	} else {
		goto L590
	}
L554:
	;
	goto L553
L555:
	;
	v1626 = v1520
	goto L554
L557:
	;
	v1626 = int32(0)
	goto L554
L558:
	;
	goto L559
L559:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1493)))
	if v1515 != 0 {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v1520 = v1493
	v1522 = v1515
	goto L564
L562:
	;
	goto L563
L563:
	;
	goto L571
L564:
	;
	v1527 = F_strcmp(m, v1520+int32(12), v1495)
	mBase = m.M
	if v1527 != 0 {
		goto L566
	} else {
		goto L567
	}
L565:
	;
	goto L563
L566:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+8))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	if v1537 != 0 {
		v1520 = v1536
		v1522 = v1537
		goto L564
	} else {
		goto L570
	}
L567:
	;
	if base.B2i32(v1522 == int32(1))&int32(0) != 0 {
		goto L566
	} else {
		goto L568
	}
L568:
	;
	goto L555
L570:
	;
	goto L565
L571:
	;
	goto L557
L590:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v1639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638)+496)))
	if v1639&int32(2) == int32(0) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638)+500)))
	if v1644&int32(2) == int32(0) {
		v10149 = v241
		goto L5
	} else {
		goto L594
	}
L592:
	;
	goto L593
L593:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v1650 == int32(0) {
		goto L599
	} else {
		goto L600
	}
L594:
	;
	goto L593
L595:
	;
	if v1783 == int32(0) {
		v10149 = v241
		goto L5
	} else {
		goto L632
	}
L596:
	;
	goto L595
L597:
	;
	v1783 = v1677
	goto L596
L599:
	;
	v1783 = int32(0)
	goto L596
L600:
	;
	v1663 = v1650
	goto L601
L601:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1663)))
	if v1672 != 0 {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	goto L599
L603:
	;
	v1677 = v1663
	v1679 = v1672
	goto L606
L604:
	;
	v1699 = v1663
	goto L605
L605:
	;
	goto L613
L606:
	;
	v1684 = F_strcmp(m, v1677+int32(12), v1652)
	mBase = m.M
	if v1684 != 0 {
		goto L608
	} else {
		goto L609
	}
L607:
	;
	v1699 = v1693
	goto L605
L608:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+8))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1693)))
	if v1694 != 0 {
		v1677 = v1693
		v1679 = v1694
		goto L606
	} else {
		goto L612
	}
L609:
	;
	if base.B2i32(v1679 == int32(1))&int32(0) != 0 {
		goto L608
	} else {
		goto L610
	}
L610:
	;
	goto L597
L612:
	;
	goto L607
L613:
	;
	goto L629
L629:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1699)+8))
	if v1741 != 0 {
		v1663 = v1741
		goto L601
	} else {
		goto L630
	}
L630:
	;
	goto L602
L632:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1797)+500))
	v1804 = F_errstart(m, v1798&int32(2)+int32(19), int32(556047))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L66
	} else {
		goto L633
	}
L633:
	;
	if v1804 == int32(0) {
		v10149 = v241
		goto L5
	} else {
		goto L634
	}
L634:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L66
	} else {
		goto L635
	}
L635:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v1811
	F_errmsg(m, int32(396619), v27+int32(96))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L66
	} else {
		goto L636
	}
L636:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v1819 = F_plpgsql_scanner_errposition(m, v1818, l1)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L66
	} else {
		goto L637
	}
L637:
	;
	F_errfinish(m, int32(26959), int32(737), int32(360630))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L66
	} else {
		goto L638
	}
L638:
	;
	v10149 = v241
	goto L5
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1827
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v1831 = int32(0)
	if v1830 < v1831 {
		v1874 = v1831
		goto L641
	} else {
		goto L642
	}
L640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1874
	v1879 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v1879 == int32(0) {
		goto L658
	} else {
		goto L659
	}
L641:
	;
	goto L640
L642:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+60))
	if v1837 == int32(0) {
		v1874 = v1831
		goto L641
	} else {
		goto L643
	}
L643:
	;
	v1840 = v1830 + v1837
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+188))
	if base.Ui32(v1841) <= base.Ui32(v1840) {
		goto L645
	} else {
		goto L646
	}
L644:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+196))
	if v1850 == int32(0) {
		v1874 = v1851
		goto L641
	} else {
		goto L648
	}
L645:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+192))
	v1850 = v1843
	goto L644
L646:
	;
	goto L647
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1836)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1836)+188)) = v1837
	v1848 = F_strchr(m, v1837, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1836)+192)) = v1848
	v1850 = v1848
	goto L644
L648:
	;
	if base.Ui32(v1840) <= base.Ui32(v1850) {
		v1874 = v1851
		goto L641
	} else {
		goto L649
	}
L649:
	;
	v1855 = v1850
	v1857 = v1851
	goto L650
L650:
	;
	v1860 = int32(1)
	v1861 = v1857 + v1860
	*(*int32)(unsafe.Add(mBase, uint32(v1836)+196)) = v1861
	v1864 = v1855 + v1860
	*(*int32)(unsafe.Add(mBase, uint32(v1836)+188)) = v1864
	v1867 = F_strchr(m, v1864, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1836)+192)) = v1867
	if v1867 == int32(0) {
		v1874 = v1861
		goto L641
	} else {
		goto L652
	}
L651:
	;
	v1874 = v1861
	goto L641
L652:
	;
	if base.Ui32(v1867) < base.Ui32(v1840) {
		v1855 = v1867
		v1857 = v1861
		goto L650
	} else {
		goto L653
	}
L653:
	;
	goto L651
L654:
	;
	if v2012 != 0 {
		goto L47
	} else {
		goto L691
	}
L655:
	;
	goto L654
L656:
	;
	v2012 = v1906
	goto L655
L658:
	;
	v2012 = int32(0)
	goto L655
L659:
	;
	goto L660
L660:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1879)))
	if v1901 != 0 {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v1906 = v1879
	v1908 = v1901
	goto L665
L663:
	;
	goto L664
L664:
	;
	goto L672
L665:
	;
	v1913 = F_strcmp(m, v1906+int32(12), v1881)
	mBase = m.M
	if v1913 != 0 {
		goto L667
	} else {
		goto L668
	}
L666:
	;
	goto L664
L667:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+8))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1922)))
	if v1923 != 0 {
		v1906 = v1922
		v1908 = v1923
		goto L665
	} else {
		goto L671
	}
L668:
	;
	if base.B2i32(v1908 == int32(1))&int32(0) != 0 {
		goto L667
	} else {
		goto L669
	}
L669:
	;
	goto L656
L671:
	;
	goto L666
L672:
	;
	goto L658
L691:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v2025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2024)+496)))
	if v2025&int32(2) == int32(0) {
		goto L692
	} else {
		goto L693
	}
L692:
	;
	v2030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2024)+500)))
	if v2030&int32(2) == int32(0) {
		v10149 = v241
		goto L5
	} else {
		goto L695
	}
L693:
	;
	goto L694
L694:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v2036 == int32(0) {
		goto L700
	} else {
		goto L701
	}
L695:
	;
	goto L694
L696:
	;
	if v2169 == int32(0) {
		v10149 = v241
		goto L5
	} else {
		goto L733
	}
L697:
	;
	goto L696
L698:
	;
	v2169 = v2063
	goto L697
L700:
	;
	v2169 = int32(0)
	goto L697
L701:
	;
	v2049 = v2036
	goto L702
L702:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v2049)))
	if v2058 != 0 {
		goto L704
	} else {
		goto L705
	}
L703:
	;
	goto L700
L704:
	;
	v2063 = v2049
	v2065 = v2058
	goto L707
L705:
	;
	v2085 = v2049
	goto L706
L706:
	;
	goto L714
L707:
	;
	v2070 = F_strcmp(m, v2063+int32(12), v2038)
	mBase = m.M
	if v2070 != 0 {
		goto L709
	} else {
		goto L710
	}
L708:
	;
	v2085 = v2079
	goto L706
L709:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2063)+8))
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2079)))
	if v2080 != 0 {
		v2063 = v2079
		v2065 = v2080
		goto L707
	} else {
		goto L713
	}
L710:
	;
	if base.B2i32(v2065 == int32(1))&int32(0) != 0 {
		goto L709
	} else {
		goto L711
	}
L711:
	;
	goto L698
L713:
	;
	goto L708
L714:
	;
	goto L730
L730:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2085)+8))
	if v2127 != 0 {
		v2049 = v2127
		goto L702
	} else {
		goto L731
	}
L731:
	;
	goto L703
L733:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+500))
	v2190 = F_errstart(m, v2184&int32(2)+int32(19), int32(556047))
	mBase = m.M
	v2191 = m.ExcPending
	if v2191 != 0 {
		goto L66
	} else {
		goto L734
	}
L734:
	;
	if v2190 == int32(0) {
		v10149 = v241
		goto L5
	} else {
		goto L735
	}
L735:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L66
	} else {
		goto L736
	}
L736:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v2197
	F_errmsg(m, int32(396619), v27+int32(112))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L66
	} else {
		goto L737
	}
L737:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v2205 = F_plpgsql_scanner_errposition(m, v2204, l1)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L66
	} else {
		goto L738
	}
L738:
	;
	F_errfinish(m, int32(26959), int32(765), int32(360630))
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L66
	} else {
		goto L739
	}
L739:
	;
	v10149 = v241
	goto L5
L740:
	;
	v2222 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L66
	} else {
		goto L743
	}
L741:
	;
	v2224 = v241
	goto L742
L742:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	if v2224 == int32(275) {
		goto L745
	} else {
		goto L746
	}
L743:
	;
	v2224 = v2222
	goto L742
L744:
	;
	if v2506 == int32(0) {
		v9984 = v2509
		goto L7
	} else {
		goto L857
	}
L745:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v2233 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L66
	} else {
		goto L748
	}
L746:
	;
	goto L747
L747:
	;
	v2315 = int32(0)
	goto L781
L748:
	;
	if v2233 != int32(37) {
		v9984 = v2233
		goto L7
	} else {
		goto L749
	}
L749:
	;
	v2241 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L66
	} else {
		goto L750
	}
L750:
	;
	switch v2241 - int32(366) {
	case 0:
		goto L751
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v9984 = v2241
		goto L7
	case 12:
		goto L753
	default:
		goto L754
	}
L751:
	;
	v2310 = F_plpgsql_parse_wordrowtype(m, v2228)
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L66
	} else {
		goto L779
	}
L752:
	;
	v2282 = int32(365467)
	v2285 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1370])))
	v2286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2249))))
	if v2286 == int32(0) {
		v2305 = v2285
		v2306 = v2286
		goto L771
	} else {
		goto L772
	}
L753:
	;
	v2280 = F_plpgsql_parse_wordtype(m, v2228)
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L66
	} else {
		goto L769
	}
L754:
	;
	v2245 = int32(277)
	if v2241 != v2245 {
		goto L755
	} else {
		goto L756
	}
L755:
	;
	v9984 = v2241
	goto L7
L756:
	;
	goto L757
L757:
	;
	v2248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v2248 != 0 {
		v9984 = v2245
		goto L7
	} else {
		goto L758
	}
L758:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v2249 == int32(0) {
		v9984 = v2245
		goto L7
	} else {
		goto L759
	}
L759:
	;
	v2252 = int32(371227)
	v2255 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1372])))
	v2256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2249))))
	if v2256 == int32(0) {
		v2275 = v2255
		v2276 = v2256
		goto L761
	} else {
		goto L762
	}
L760:
	;
	if v2276-v2275 != 0 {
		goto L752
	} else {
		goto L768
	}
L761:
	;
	goto L760
L762:
	;
	if v2255 != v2256 {
		v2275 = v2255
		v2276 = v2256
		goto L761
	} else {
		goto L763
	}
L763:
	;
	v2260 = v2249
	v2261 = v2252
	goto L764
L764:
	;
	v2264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2261)+1)))
	v2265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2260)+1)))
	if v2265 == int32(0) {
		v2275 = v2264
		v2276 = v2265
		goto L761
	} else {
		goto L766
	}
L765:
	;
	v2275 = v2264
	v2276 = v2265
	goto L761
L766:
	;
	v2268 = int32(1)
	if v2264 == v2265 {
		v2260 = v2260 + v2268
		v2261 = v2261 + v2268
		goto L764
	} else {
		goto L767
	}
L767:
	;
	goto L765
L768:
	;
	goto L753
L769:
	;
	v2506 = v2280
	v2509 = v2241
	goto L744
L770:
	;
	if v2306-v2305 != 0 {
		v9984 = v2245
		goto L7
	} else {
		goto L778
	}
L771:
	;
	goto L770
L772:
	;
	if v2285 != v2286 {
		v2305 = v2285
		v2306 = v2286
		goto L771
	} else {
		goto L773
	}
L773:
	;
	v2290 = v2249
	v2291 = v2282
	goto L774
L774:
	;
	v2294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2291)+1)))
	v2295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2290)+1)))
	if v2295 == int32(0) {
		v2305 = v2294
		v2306 = v2295
		goto L771
	} else {
		goto L776
	}
L775:
	;
	v2305 = v2294
	v2306 = v2295
	goto L771
L776:
	;
	v2298 = int32(1)
	if v2294 == v2295 {
		v2290 = v2290 + v2298
		v2291 = v2291 + v2298
		goto L774
	} else {
		goto L777
	}
L777:
	;
	goto L775
L778:
	;
	goto L751
L779:
	;
	v2506 = v2310
	v2509 = v2241
	goto L744
L780:
	;
	if v2224 == v2321 {
		goto L787
	} else {
		goto L788
	}
L781:
	;
	v2321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2315<<(uint(int32(1))%32))+uint32(_consts[1373]))))
	v2322 = base.B2i32(v2224 == v2321)
	if v2322 == int32(0) {
		goto L783
	} else {
		goto L784
	}
L782:
	;
	goto L780
L783:
	;
	v2326 = v2315 + int32(1)
	if v2326 != int32(83) {
		v2315 = v2326
		goto L781
	} else {
		goto L786
	}
L784:
	;
	goto L785
L785:
	;
	goto L782
L786:
	;
	goto L785
L787:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v2331 = F_pstrdup(m, v2330)
	mBase = m.M
	v2332 = m.ExcPending
	if v2332 != 0 {
		goto L66
	} else {
		goto L790
	}
L788:
	;
	goto L789
L789:
	;
	if v2224 != int32(276) {
		v9984 = v2224
		goto L7
	} else {
		goto L823
	}
L790:
	;
	v2337 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L66
	} else {
		goto L791
	}
L791:
	;
	if v2337 != int32(37) {
		v9984 = v2337
		goto L7
	} else {
		goto L792
	}
L792:
	;
	v2345 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L66
	} else {
		goto L793
	}
L793:
	;
	switch v2345 - int32(366) {
	case 0:
		goto L794
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v9984 = v2345
		goto L7
	case 12:
		goto L796
	default:
		goto L797
	}
L794:
	;
	v2414 = F_plpgsql_parse_wordrowtype(m, v2331)
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L66
	} else {
		goto L822
	}
L795:
	;
	v2386 = int32(365467)
	v2389 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1370])))
	v2390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353))))
	if v2390 == int32(0) {
		v2409 = v2389
		v2410 = v2390
		goto L814
	} else {
		goto L815
	}
L796:
	;
	v2384 = F_plpgsql_parse_wordtype(m, v2331)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L66
	} else {
		goto L812
	}
L797:
	;
	v2349 = int32(277)
	if v2345 != v2349 {
		goto L798
	} else {
		goto L799
	}
L798:
	;
	v9984 = v2345
	goto L7
L799:
	;
	goto L800
L800:
	;
	v2352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v2352 != 0 {
		v9984 = v2349
		goto L7
	} else {
		goto L801
	}
L801:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v2353 == int32(0) {
		v9984 = v2349
		goto L7
	} else {
		goto L802
	}
L802:
	;
	v2356 = int32(371227)
	v2359 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1372])))
	v2360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353))))
	if v2360 == int32(0) {
		v2379 = v2359
		v2380 = v2360
		goto L804
	} else {
		goto L805
	}
L803:
	;
	if v2380-v2379 != 0 {
		goto L795
	} else {
		goto L811
	}
L804:
	;
	goto L803
L805:
	;
	if v2359 != v2360 {
		v2379 = v2359
		v2380 = v2360
		goto L804
	} else {
		goto L806
	}
L806:
	;
	v2364 = v2353
	v2365 = v2356
	goto L807
L807:
	;
	v2368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2365)+1)))
	v2369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2364)+1)))
	if v2369 == int32(0) {
		v2379 = v2368
		v2380 = v2369
		goto L804
	} else {
		goto L809
	}
L808:
	;
	v2379 = v2368
	v2380 = v2369
	goto L804
L809:
	;
	v2372 = int32(1)
	if v2368 == v2369 {
		v2364 = v2364 + v2372
		v2365 = v2365 + v2372
		goto L807
	} else {
		goto L810
	}
L810:
	;
	goto L808
L811:
	;
	goto L796
L812:
	;
	v2506 = v2384
	v2509 = v2345
	goto L744
L813:
	;
	if v2410-v2409 != 0 {
		v9984 = v2349
		goto L7
	} else {
		goto L821
	}
L814:
	;
	goto L813
L815:
	;
	if v2389 != v2390 {
		v2409 = v2389
		v2410 = v2390
		goto L814
	} else {
		goto L816
	}
L816:
	;
	v2394 = v2353
	v2395 = v2386
	goto L817
L817:
	;
	v2398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2395)+1)))
	v2399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2394)+1)))
	if v2399 == int32(0) {
		v2409 = v2398
		v2410 = v2399
		goto L814
	} else {
		goto L819
	}
L818:
	;
	v2409 = v2398
	v2410 = v2399
	goto L814
L819:
	;
	v2402 = int32(1)
	if v2398 == v2399 {
		v2394 = v2394 + v2402
		v2395 = v2395 + v2402
		goto L817
	} else {
		goto L820
	}
L820:
	;
	goto L818
L821:
	;
	goto L794
L822:
	;
	v2506 = v2414
	v2509 = v2345
	goto L744
L823:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v2423 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L66
	} else {
		goto L824
	}
L824:
	;
	if v2423 != int32(37) {
		v9984 = v2423
		goto L7
	} else {
		goto L825
	}
L825:
	;
	v2431 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L66
	} else {
		goto L826
	}
L826:
	;
	switch v2431 - int32(366) {
	case 0:
		goto L828
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v9984 = v2431
		goto L7
	case 12:
		goto L830
	default:
		goto L831
	}
L827:
	;
	v2506 = v2504
	v2509 = v2431
	goto L744
L828:
	;
	v2500 = F_plpgsql_parse_cwordrowtype(m, v2418)
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L66
	} else {
		goto L856
	}
L829:
	;
	v2472 = int32(365467)
	v2475 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1370])))
	v2476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2439))))
	if v2476 == int32(0) {
		v2495 = v2475
		v2496 = v2476
		goto L848
	} else {
		goto L849
	}
L830:
	;
	v2470 = F_plpgsql_parse_cwordtype(m, v2418)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L66
	} else {
		goto L846
	}
L831:
	;
	v2435 = int32(277)
	if v2431 != v2435 {
		goto L832
	} else {
		goto L833
	}
L832:
	;
	v9984 = v2431
	goto L7
L833:
	;
	goto L834
L834:
	;
	v2438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v2438 != 0 {
		v9984 = v2435
		goto L7
	} else {
		goto L835
	}
L835:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v2439 == int32(0) {
		v9984 = v2435
		goto L7
	} else {
		goto L836
	}
L836:
	;
	v2442 = int32(371227)
	v2445 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1372])))
	v2446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2439))))
	if v2446 == int32(0) {
		v2465 = v2445
		v2466 = v2446
		goto L838
	} else {
		goto L839
	}
L837:
	;
	if v2466-v2465 != 0 {
		goto L829
	} else {
		goto L845
	}
L838:
	;
	goto L837
L839:
	;
	if v2445 != v2446 {
		v2465 = v2445
		v2466 = v2446
		goto L838
	} else {
		goto L840
	}
L840:
	;
	v2450 = v2439
	v2451 = v2442
	goto L841
L841:
	;
	v2454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2451)+1)))
	v2455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2450)+1)))
	if v2455 == int32(0) {
		v2465 = v2454
		v2466 = v2455
		goto L838
	} else {
		goto L843
	}
L842:
	;
	v2465 = v2454
	v2466 = v2455
	goto L838
L843:
	;
	v2458 = int32(1)
	if v2454 == v2455 {
		v2450 = v2450 + v2458
		v2451 = v2451 + v2458
		goto L841
	} else {
		goto L844
	}
L844:
	;
	goto L842
L845:
	;
	goto L830
L846:
	;
	v2504 = v2470
	goto L827
L847:
	;
	if v2496-v2495 != 0 {
		v9984 = v2435
		goto L7
	} else {
		goto L855
	}
L848:
	;
	goto L847
L849:
	;
	if v2475 != v2476 {
		v2495 = v2475
		v2496 = v2476
		goto L848
	} else {
		goto L850
	}
L850:
	;
	v2480 = v2439
	v2481 = v2472
	goto L851
L851:
	;
	v2484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2481)+1)))
	v2485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2480)+1)))
	if v2485 == int32(0) {
		v2495 = v2484
		v2496 = v2485
		goto L848
	} else {
		goto L853
	}
L852:
	;
	v2495 = v2484
	v2496 = v2485
	goto L848
L853:
	;
	v2488 = int32(1)
	if v2484 == v2485 {
		v2480 = v2480 + v2488
		v2481 = v2481 + v2488
		goto L851
	} else {
		goto L854
	}
L854:
	;
	goto L852
L855:
	;
	goto L828
L856:
	;
	v2504 = v2500
	goto L827
L857:
	;
	v2516 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L66
	} else {
		goto L863
	}
L858:
	;
	F_plpgsql_push_back_token(m, v2624, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L66
	} else {
		goto L890
	}
L859:
	;
	goto L880
L860:
	;
	if v2516 != int32(91) {
		goto L8
	} else {
		goto L879
	}
L861:
	;
	v2565 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2566 = m.ExcPending
	if v2566 != 0 {
		goto L66
	} else {
		goto L877
	}
L862:
	;
	v2520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v2520 != 0 {
		goto L864
	} else {
		goto L865
	}
L863:
	;
	switch v2516 - int32(277) {
	case 0:
		goto L862
	case 1, 2, 3, 4, 5, 6:
		goto L8
	case 7:
		goto L861
	default:
		goto L860
	}
L864:
	;
	F_plpgsql_push_back_token(m, int32(277), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2559 = m.ExcPending
	if v2559 != 0 {
		goto L66
	} else {
		goto L876
	}
L865:
	;
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v2521 == int32(0) {
		goto L864
	} else {
		goto L866
	}
L866:
	;
	v2524 = int32(26214)
	v2527 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1374])))
	v2528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2521))))
	if v2528 == int32(0) {
		v2547 = v2527
		v2548 = v2528
		goto L868
	} else {
		goto L869
	}
L867:
	;
	if v2548-v2547 == int32(0) {
		goto L861
	} else {
		goto L875
	}
L868:
	;
	goto L867
L869:
	;
	if v2527 != v2528 {
		v2547 = v2527
		v2548 = v2528
		goto L868
	} else {
		goto L870
	}
L870:
	;
	v2532 = v2521
	v2533 = v2524
	goto L871
L871:
	;
	v2536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2533)+1)))
	v2537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2532)+1)))
	if v2537 == int32(0) {
		v2547 = v2536
		v2548 = v2537
		goto L868
	} else {
		goto L873
	}
L872:
	;
	v2547 = v2536
	v2548 = v2537
	goto L868
L873:
	;
	v2540 = int32(1)
	if v2536 == v2537 {
		v2532 = v2532 + v2540
		v2533 = v2533 + v2540
		goto L871
	} else {
		goto L874
	}
L874:
	;
	goto L872
L875:
	;
	goto L864
L876:
	;
	v10120 = v2506
	goto L6
L877:
	;
	if v2565 == int32(91) {
		goto L859
	} else {
		goto L878
	}
L878:
	;
	v2624 = v2565
	goto L858
L879:
	;
	goto L859
L880:
	;
	v2600 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2601 = m.ExcPending
	if v2601 != 0 {
		goto L66
	} else {
		goto L882
	}
L881:
	;
	v2624 = v2617
	goto L858
L882:
	;
	if v2600 == int32(266) {
		goto L883
	} else {
		goto L884
	}
L883:
	;
	v2608 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L66
	} else {
		goto L886
	}
L884:
	;
	v2610 = v2600
	goto L885
L885:
	;
	if v2610 != int32(93) {
		goto L46
	} else {
		goto L887
	}
L886:
	;
	v2610 = v2608
	goto L885
L887:
	;
	v2617 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L66
	} else {
		goto L888
	}
L888:
	;
	if v2617 == int32(91) {
		goto L880
	} else {
		goto L889
	}
L889:
	;
	goto L881
L890:
	;
	v2651 = F_plpgsql_build_datatype_arrayof(m, v2506)
	mBase = m.M
	v2652 = m.ExcPending
	if v2652 != 0 {
		goto L66
	} else {
		goto L891
	}
L891:
	;
	v10120 = v2651
	goto L6
L892:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+120)) = v2656
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v2656
	v2663 = F_list_make1_impl(m, int32(1), v27+int32(120))
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L66
	} else {
		goto L893
	}
L893:
	;
	v2666 = F_get_collation_oid(m, v2663, int32(0))
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L66
	} else {
		goto L894
	}
L894:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2666
	v10149 = v241
	goto L5
L895:
	;
	v2672 = F_makeString(m, v2670)
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L66
	} else {
		goto L896
	}
L896:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+124)) = v2672
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v2672
	v2679 = F_list_make1_impl(m, int32(1), v27+int32(124))
	mBase = m.M
	v2680 = m.ExcPending
	if v2680 != 0 {
		goto L66
	} else {
		goto L897
	}
L897:
	;
	v2682 = F_get_collation_oid(m, v2679, int32(0))
	mBase = m.M
	v2683 = m.ExcPending
	if v2683 != 0 {
		goto L66
	} else {
		goto L898
	}
L898:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2682
	v10149 = v241
	goto L5
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2687
	v10149 = v241
	goto L5
L900:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2709
	v10149 = v241
	goto L5
L901:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2716
	v10149 = v241
	goto L5
L902:
	;
	goto L903
L903:
	;
	v2721 = F_lappend(m, v2716, v2717)
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L66
	} else {
		goto L904
	}
L904:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2721
	v10149 = v241
	goto L5
L905:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2775))) = int32(23)
	v2779 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v2780 = int32(0)
	if v2779 < v2780 {
		v2823 = v2780
		goto L907
	} else {
		goto L908
	}
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2775)+4)) = v2823
	v2828 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v2828)+520))
	v2831 = v2829 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2828)+520)) = v2831
	*(*int32)(unsafe.Add(mBase, uint32(v2775)+8)) = v2831
	F_plpgsql_push_back_token(m, int32(349), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L66
	} else {
		goto L920
	}
L907:
	;
	goto L906
L908:
	;
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2785)+60))
	if v2786 == int32(0) {
		v2823 = v2780
		goto L907
	} else {
		goto L909
	}
L909:
	;
	v2789 = v2779 + v2786
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v2785)+188))
	if base.Ui32(v2790) <= base.Ui32(v2789) {
		goto L911
	} else {
		goto L912
	}
L910:
	;
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v2785)+196))
	if v2799 == int32(0) {
		v2823 = v2800
		goto L907
	} else {
		goto L914
	}
L911:
	;
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v2785)+192))
	v2799 = v2792
	goto L910
L912:
	;
	goto L913
L913:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2785)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2785)+188)) = v2786
	v2797 = F_strchr(m, v2786, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2785)+192)) = v2797
	v2799 = v2797
	goto L910
L914:
	;
	if base.Ui32(v2789) <= base.Ui32(v2799) {
		v2823 = v2800
		goto L907
	} else {
		goto L915
	}
L915:
	;
	v2804 = v2799
	v2806 = v2800
	goto L916
L916:
	;
	v2809 = int32(1)
	v2810 = v2806 + v2809
	*(*int32)(unsafe.Add(mBase, uint32(v2785)+196)) = v2810
	v2813 = v2804 + v2809
	*(*int32)(unsafe.Add(mBase, uint32(v2785)+188)) = v2813
	v2816 = F_strchr(m, v2813, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2785)+192)) = v2816
	if v2816 == int32(0) {
		v2823 = v2810
		goto L907
	} else {
		goto L918
	}
L917:
	;
	v2823 = v2810
	goto L907
L918:
	;
	if base.Ui32(v2816) < base.Ui32(v2789) {
		v2804 = v2816
		v2806 = v2810
		goto L916
	} else {
		goto L919
	}
L919:
	;
	goto L917
L920:
	;
	v2842 = int32(0)
	v2855 = F_read_sql_construct(m, int32(59), v2842, v2842, int32(546849), v2842, v2842, v2842, v27+int32(4752), v2842, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v2856 = m.ExcPending
	if v2856 != 0 {
		goto L66
	} else {
		goto L921
	}
L921:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2775)+12)) = v2855
	v2858 = *(*int32)(unsafe.Add(mBase, uint32(v2855)))
	v2860 = *(*int32)(unsafe.Add(mBase, _consts[1375]))
	*(*int32)(unsafe.Add(mBase, uint32(v2858))) = v2860
	v2863 = *(*int32)(unsafe.Add(mBase, _consts[1376]))
	*(*int32)(unsafe.Add(mBase, uint32(v2858)+3)) = v2863
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(v2775)+12))
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v2865)))
	v2868 = v2866 + int32(1)
	v2869 = F_strlen(m, v2866)
	mBase = m.M
	if v2866 == v2868 {
		goto L923
	} else {
		goto L924
	}
L922:
	;
	v3015 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1356])))
	if v3015 == int32(1) {
		goto L968
	} else {
		goto L969
	}
L923:
	;
	goto L922
L924:
	;
	v2873 = v2866 + v2869
	if base.Ui32(v2868-v2873) <= base.Ui32(int32(0)-v2869<<(uint(int32(1))%32)) {
		goto L925
	} else {
		goto L926
	}
L925:
	;
	v2880 = F___memcpy(m, v2866, v2868, v2869)
	mBase = m.M
	goto L922
L926:
	;
	goto L927
L927:
	;
	v2883 = (v2866 ^ v2868) & int32(3)
	if base.Ui32(v2866) < base.Ui32(v2868) {
		goto L930
	} else {
		goto L931
	}
L928:
	;
	if v2985 == int32(0) {
		goto L923
	} else {
		goto L964
	}
L929:
	;
	if base.Ui32(v2963) <= base.Ui32(int32(3)) {
		v2984 = v2962
		v2985 = v2963
		v2986 = v2964
		goto L928
	} else {
		goto L960
	}
L930:
	;
	if v2883 != 0 {
		goto L933
	} else {
		goto L934
	}
L931:
	;
	goto L932
L932:
	;
	if v2883 != 0 {
		v2945 = v2869
		goto L943
	} else {
		goto L944
	}
L933:
	;
	v2984 = v2868
	v2985 = v2869
	v2986 = v2866
	goto L928
L934:
	;
	goto L935
L935:
	;
	if v2866&int32(3) == int32(0) {
		goto L936
	} else {
		goto L937
	}
L936:
	;
	v2962 = v2868
	v2963 = v2869
	v2964 = v2866
	goto L929
L937:
	;
	goto L938
L938:
	;
	v2890 = v2868
	v2891 = v2869
	v2892 = v2866
	goto L939
L939:
	;
	if v2891 == int32(0) {
		goto L923
	} else {
		goto L941
	}
L940:
	;
	v2962 = v2899
	v2963 = v2901
	v2964 = v2903
	goto L929
L941:
	;
	v2896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2890))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2892))) = uint8(v2896)
	v2898 = int32(1)
	v2899 = v2890 + v2898
	v2901 = v2891 - v2898
	v2903 = v2892 + v2898
	if v2903&int32(3) != 0 {
		v2890 = v2899
		v2891 = v2901
		v2892 = v2903
		goto L939
	} else {
		goto L942
	}
L942:
	;
	goto L940
L943:
	;
	if v2945 == int32(0) {
		goto L923
	} else {
		goto L956
	}
L944:
	;
	if v2873&int32(3) != 0 {
		goto L945
	} else {
		goto L946
	}
L945:
	;
	v2910 = v2869
	goto L948
L946:
	;
	v2925 = v2869
	goto L947
L947:
	;
	if base.Ui32(v2925) <= base.Ui32(int32(3)) {
		v2945 = v2925
		goto L943
	} else {
		goto L952
	}
L948:
	;
	if v2910 == int32(0) {
		goto L923
	} else {
		goto L950
	}
L949:
	;
	v2925 = v2916
	goto L947
L950:
	;
	v2916 = v2910 - int32(1)
	v2917 = v2866 + v2916
	v2919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2868+v2916))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2917))) = uint8(v2919)
	if v2917&int32(3) != 0 {
		v2910 = v2916
		goto L948
	} else {
		goto L951
	}
L951:
	;
	goto L949
L952:
	;
	v2932 = v2925
	goto L953
L953:
	;
	v2936 = v2932 - int32(4)
	v2939 = *(*int32)(unsafe.Add(mBase, uint32(v2868+v2936)))
	*(*int32)(unsafe.Add(mBase, uint32(v2866+v2936))) = v2939
	if base.Ui32(int32(3)) < base.Ui32(v2936) {
		v2932 = v2936
		goto L953
	} else {
		goto L955
	}
L954:
	;
	v2945 = v2936
	goto L943
L955:
	;
	goto L954
L956:
	;
	v2952 = v2945
	goto L957
L957:
	;
	v2956 = v2952 - int32(1)
	v2959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2868+v2956))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2866+v2956))) = uint8(v2959)
	if v2956 != 0 {
		v2952 = v2956
		goto L957
	} else {
		goto L959
	}
L958:
	;
	goto L923
L959:
	;
	goto L958
L960:
	;
	v2969 = v2962
	v2970 = v2963
	v2971 = v2964
	goto L961
L961:
	;
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v2969)))
	*(*int32)(unsafe.Add(mBase, uint32(v2971))) = v2973
	v2975 = int32(4)
	v2976 = v2969 + v2975
	v2978 = v2971 + v2975
	v2980 = v2970 - v2975
	if base.Ui32(int32(3)) < base.Ui32(v2980) {
		v2969 = v2976
		v2970 = v2980
		v2971 = v2978
		goto L961
	} else {
		goto L963
	}
L962:
	;
	v2984 = v2976
	v2985 = v2980
	v2986 = v2978
	goto L928
L963:
	;
	goto L962
L964:
	;
	v2991 = v2984
	v2992 = v2985
	v2993 = v2986
	goto L965
L965:
	;
	v2995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2991))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2993))) = uint8(v2995)
	v2997 = int32(1)
	v3002 = v2992 - v2997
	if v3002 != 0 {
		v2991 = v2991 + v2997
		v2992 = v3002
		v2993 = v2993 + v2997
		goto L965
	} else {
		goto L967
	}
L966:
	;
	goto L923
L967:
	;
	goto L966
L968:
	;
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v2775)+12))
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v3018)+4))
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v3018)))
	v3021 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377])))
	v3022 = int32(4515248)
	v3023 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v3026 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3026
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v3021 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1380]))) = int32(6810)
	v3034 = int32(4508152)
	v3035 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4784)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = v3035
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1381]))) = v27 + int32(4768)
	v3044 = F_raw_parser(m, v3020, v3019)
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L66
	} else {
		goto L971
	}
L969:
	;
	goto L970
L970:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2775
	v10149 = v241
	goto L5
L971:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3023
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v3049
	goto L970
L972:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3059))) = int32(24)
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v3064 = int32(0)
	if v3063 < v3064 {
		v3107 = v3064
		goto L974
	} else {
		goto L975
	}
L973:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3059)+4)) = v3107
	v3112 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v3113 = *(*int32)(unsafe.Add(mBase, uint32(v3112)+520))
	v3115 = v3113 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3112)+520)) = v3115
	*(*int32)(unsafe.Add(mBase, uint32(v3059)+8)) = v3115
	F_plpgsql_push_back_token(m, int32(289), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L66
	} else {
		goto L987
	}
L974:
	;
	goto L973
L975:
	;
	v3069 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3070 = *(*int32)(unsafe.Add(mBase, uint32(v3069)+60))
	if v3070 == int32(0) {
		v3107 = v3064
		goto L974
	} else {
		goto L976
	}
L976:
	;
	v3073 = v3063 + v3070
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v3069)+188))
	if base.Ui32(v3074) <= base.Ui32(v3073) {
		goto L978
	} else {
		goto L979
	}
L977:
	;
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v3069)+196))
	if v3083 == int32(0) {
		v3107 = v3084
		goto L974
	} else {
		goto L981
	}
L978:
	;
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v3069)+192))
	v3083 = v3076
	goto L977
L979:
	;
	goto L980
L980:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3069)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3069)+188)) = v3070
	v3081 = F_strchr(m, v3070, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3069)+192)) = v3081
	v3083 = v3081
	goto L977
L981:
	;
	if base.Ui32(v3073) <= base.Ui32(v3083) {
		v3107 = v3084
		goto L974
	} else {
		goto L982
	}
L982:
	;
	v3088 = v3083
	v3090 = v3084
	goto L983
L983:
	;
	v3093 = int32(1)
	v3094 = v3090 + v3093
	*(*int32)(unsafe.Add(mBase, uint32(v3069)+196)) = v3094
	v3097 = v3088 + v3093
	*(*int32)(unsafe.Add(mBase, uint32(v3069)+188)) = v3097
	v3100 = F_strchr(m, v3097, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3069)+192)) = v3100
	if v3100 == int32(0) {
		v3107 = v3094
		goto L974
	} else {
		goto L985
	}
L984:
	;
	v3107 = v3094
	goto L974
L985:
	;
	if base.Ui32(v3100) < base.Ui32(v3073) {
		v3088 = v3100
		v3090 = v3094
		goto L983
	} else {
		goto L986
	}
L986:
	;
	goto L984
L987:
	;
	v3126 = int32(0)
	v3138 = F_read_sql_construct(m, int32(59), v3126, v3126, int32(546849), v3126, v3126, int32(1), v3126, v3126, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		goto L66
	} else {
		goto L988
	}
L988:
	;
	v3140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3059)+16)) = uint8(v3140)
	*(*int32)(unsafe.Add(mBase, uint32(v3059)+12)) = v3138
	v3144 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3144)+524)) = uint8(v3140)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3059
	v10149 = v241
	goto L5
L989:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3149))) = int32(24)
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v3154 = int32(0)
	if v3153 < v3154 {
		v3197 = v3154
		goto L991
	} else {
		goto L992
	}
L990:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3149)+4)) = v3197
	v3202 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v3203 = *(*int32)(unsafe.Add(mBase, uint32(v3202)+520))
	v3205 = v3203 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3202)+520)) = v3205
	*(*int32)(unsafe.Add(mBase, uint32(v3149)+8)) = v3205
	F_plpgsql_push_back_token(m, int32(309), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3214 = m.ExcPending
	if v3214 != 0 {
		goto L66
	} else {
		goto L1004
	}
L991:
	;
	goto L990
L992:
	;
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3160 = *(*int32)(unsafe.Add(mBase, uint32(v3159)+60))
	if v3160 == int32(0) {
		v3197 = v3154
		goto L991
	} else {
		goto L993
	}
L993:
	;
	v3163 = v3153 + v3160
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v3159)+188))
	if base.Ui32(v3164) <= base.Ui32(v3163) {
		goto L995
	} else {
		goto L996
	}
L994:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v3159)+196))
	if v3173 == int32(0) {
		v3197 = v3174
		goto L991
	} else {
		goto L998
	}
L995:
	;
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v3159)+192))
	v3173 = v3166
	goto L994
L996:
	;
	goto L997
L997:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3159)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3159)+188)) = v3160
	v3171 = F_strchr(m, v3160, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3159)+192)) = v3171
	v3173 = v3171
	goto L994
L998:
	;
	if base.Ui32(v3163) <= base.Ui32(v3173) {
		v3197 = v3174
		goto L991
	} else {
		goto L999
	}
L999:
	;
	v3178 = v3173
	v3180 = v3174
	goto L1000
L1000:
	;
	v3183 = int32(1)
	v3184 = v3180 + v3183
	*(*int32)(unsafe.Add(mBase, uint32(v3159)+196)) = v3184
	v3187 = v3178 + v3183
	*(*int32)(unsafe.Add(mBase, uint32(v3159)+188)) = v3187
	v3190 = F_strchr(m, v3187, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3159)+192)) = v3190
	if v3190 == int32(0) {
		v3197 = v3184
		goto L991
	} else {
		goto L1002
	}
L1001:
	;
	v3197 = v3184
	goto L991
L1002:
	;
	if base.Ui32(v3190) < base.Ui32(v3163) {
		v3178 = v3190
		v3180 = v3184
		goto L1000
	} else {
		goto L1003
	}
L1003:
	;
	goto L1001
L1004:
	;
	v3216 = int32(0)
	v3228 = F_read_sql_construct(m, int32(59), v3216, v3216, int32(546849), v3216, v3216, int32(1), v3216, v3216, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		goto L66
	} else {
		goto L1005
	}
L1005:
	;
	v3230 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3149)+16)) = uint8(v3230)
	*(*int32)(unsafe.Add(mBase, uint32(v3149)+12)) = v3228
	v3234 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v3235 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3234)+524)) = uint8(v3235)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3149
	v10149 = v241
	goto L5
L1006:
	;
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	if v3242 == int32(0) {
		goto L45
	} else {
		goto L1009
	}
L1007:
	;
	v3253 = int32(3)
	goto L1008
L1008:
	;
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_check_assignable(m, v3254, v3255, l1)
	mBase = m.M
	v3257 = m.ExcPending
	if v3257 != 0 {
		goto L66
	} else {
		goto L1011
	}
L1009:
	;
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(v3242)+4))
	if base.Ui32(int32(3)) <= base.Ui32(v3245-int32(1)) {
		goto L45
	} else {
		goto L1010
	}
L1010:
	;
	v3253 = v3245 + int32(2)
	goto L1008
L1011:
	;
	v3259 = F_palloc0(m, int32(20))
	mBase = m.M
	v3260 = m.ExcPending
	if v3260 != 0 {
		goto L66
	} else {
		goto L1012
	}
L1012:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3259))) = int32(1)
	v3263 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v3264 = int32(0)
	if v3263 < v3264 {
		v3307 = v3264
		goto L1014
	} else {
		goto L1015
	}
L1013:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3259)+4)) = v3307
	v3312 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(v3312)+520))
	v3315 = v3313 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3312)+520)) = v3315
	*(*int32)(unsafe.Add(mBase, uint32(v3259)+8)) = v3315
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v3318)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3259)+12)) = v3319
	F_plpgsql_push_back_token(m, int32(277), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3327 = m.ExcPending
	if v3327 != 0 {
		goto L66
	} else {
		goto L1027
	}
L1014:
	;
	goto L1013
L1015:
	;
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(v3269)+60))
	if v3270 == int32(0) {
		v3307 = v3264
		goto L1014
	} else {
		goto L1016
	}
L1016:
	;
	v3273 = v3263 + v3270
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v3269)+188))
	if base.Ui32(v3274) <= base.Ui32(v3273) {
		goto L1018
	} else {
		goto L1019
	}
L1017:
	;
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v3269)+196))
	if v3283 == int32(0) {
		v3307 = v3284
		goto L1014
	} else {
		goto L1021
	}
L1018:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3269)+192))
	v3283 = v3276
	goto L1017
L1019:
	;
	goto L1020
L1020:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+188)) = v3270
	v3281 = F_strchr(m, v3270, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+192)) = v3281
	v3283 = v3281
	goto L1017
L1021:
	;
	if base.Ui32(v3273) <= base.Ui32(v3283) {
		v3307 = v3284
		goto L1014
	} else {
		goto L1022
	}
L1022:
	;
	v3288 = v3283
	v3290 = v3284
	goto L1023
L1023:
	;
	v3293 = int32(1)
	v3294 = v3290 + v3293
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+196)) = v3294
	v3297 = v3288 + v3293
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+188)) = v3297
	v3300 = F_strchr(m, v3297, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+192)) = v3300
	if v3300 == int32(0) {
		v3307 = v3294
		goto L1014
	} else {
		goto L1025
	}
L1024:
	;
	v3307 = v3294
	goto L1014
L1025:
	;
	if base.Ui32(v3300) < base.Ui32(v3273) {
		v3288 = v3300
		v3290 = v3294
		goto L1023
	} else {
		goto L1026
	}
L1026:
	;
	goto L1024
L1027:
	;
	v3329 = int32(0)
	v3340 = F_read_sql_construct(m, int32(59), v3329, v3329, int32(546849), v3253, v3329, int32(1), v3329, v3329, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3341 = m.ExcPending
	if v3341 != 0 {
		goto L66
	} else {
		goto L1028
	}
L1028:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3259)+16)) = v3340
	v3344 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(v3344)))
	if v3345 != 0 {
		goto L1029
	} else {
		goto L1030
	}
L1029:
	;
	v3349 = int32(-1)
	v3350 = int32(0)
	goto L1031
L1030:
	;
	v3347 = *(*int32)(unsafe.Add(mBase, uint32(v3344)+4))
	v3349 = v3347
	v3350 = int32(1)
	goto L1031
L1031:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3340)+20)) = uint8(v3350)
	*(*int32)(unsafe.Add(mBase, uint32(v3340)+16)) = v3349
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3259
	v10149 = v241
	goto L5
L1032:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3355))) = int32(19)
	v3360 = v142 - int32(16)
	v3361 = *(*int32)(unsafe.Add(mBase, uint32(v3360)))
	v3362 = int32(0)
	if v3361 < v3362 {
		v3405 = v3362
		goto L1034
	} else {
		goto L1035
	}
L1033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3355)+4)) = v3405
	v3410 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v3411 = *(*int32)(unsafe.Add(mBase, uint32(v3410)+520))
	v3413 = v3411 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3410)+520)) = v3413
	*(*int32)(unsafe.Add(mBase, uint32(v3355)+8)) = v3413
	v3418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132-int32(48)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3355)+12)) = uint8(v3418)
	v3422 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v3355)+16)) = v3422
	if v3422 == int32(0) {
		goto L1047
	} else {
		goto L1048
	}
L1034:
	;
	goto L1033
L1035:
	;
	v3367 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(v3367)+60))
	if v3368 == int32(0) {
		v3405 = v3362
		goto L1034
	} else {
		goto L1036
	}
L1036:
	;
	v3371 = v3361 + v3368
	v3372 = *(*int32)(unsafe.Add(mBase, uint32(v3367)+188))
	if base.Ui32(v3372) <= base.Ui32(v3371) {
		goto L1038
	} else {
		goto L1039
	}
L1037:
	;
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(v3367)+196))
	if v3381 == int32(0) {
		v3405 = v3382
		goto L1034
	} else {
		goto L1041
	}
L1038:
	;
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(v3367)+192))
	v3381 = v3374
	goto L1037
L1039:
	;
	goto L1040
L1040:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3367)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3367)+188)) = v3368
	v3379 = F_strchr(m, v3368, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3367)+192)) = v3379
	v3381 = v3379
	goto L1037
L1041:
	;
	if base.Ui32(v3371) <= base.Ui32(v3381) {
		v3405 = v3382
		goto L1034
	} else {
		goto L1042
	}
L1042:
	;
	v3386 = v3381
	v3388 = v3382
	goto L1043
L1043:
	;
	v3391 = int32(1)
	v3392 = v3388 + v3391
	*(*int32)(unsafe.Add(mBase, uint32(v3367)+196)) = v3392
	v3395 = v3386 + v3391
	*(*int32)(unsafe.Add(mBase, uint32(v3367)+188)) = v3395
	v3398 = F_strchr(m, v3395, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3367)+192)) = v3398
	if v3398 == int32(0) {
		v3405 = v3392
		goto L1034
	} else {
		goto L1045
	}
L1044:
	;
	v3405 = v3392
	goto L1034
L1045:
	;
	if base.Ui32(v3398) < base.Ui32(v3371) {
		v3386 = v3398
		v3388 = v3392
		goto L1043
	} else {
		goto L1046
	}
L1046:
	;
	goto L1044
L1047:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3355
	v10149 = v241
	goto L5
L1048:
	;
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(v3422)+4))
	if v3426 <= int32(0) {
		goto L1047
	} else {
		goto L1049
	}
L1049:
	;
	v3429 = int32(0)
	if v3429 < v3426 {
		goto L1050
	} else {
		goto L1051
	}
L1050:
	;
	v3433 = v3426
	goto L1052
L1051:
	;
	v3433 = v3429
	goto L1052
L1052:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, uint32(v3422)+12))
	v3438 = v3429
	goto L1053
L1053:
	;
	v3462 = *(*int32)(unsafe.Add(mBase, uint32(v3434+v3438<<(uint(int32(2))%32))))
	v3463 = *(*int32)(unsafe.Add(mBase, uint32(v3462)))
	if base.Ui32(int32(10)) <= base.Ui32(v3463-int32(3)) {
		goto L1057
	} else {
		goto L1058
	}
L1054:
	;
	goto L1047
L1055:
	;
	v3553 = v3438 + int32(1)
	if v3553 != v3433 {
		v3438 = v3553
		goto L1053
	} else {
		goto L1084
	}
L1056:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v3539 = m.ExcPending
	if v3539 != 0 {
		goto L66
	} else {
		goto L1081
	}
L1057:
	;
	switch v3463 {
	case 0, 1:
		goto L1060
	case 2:
		goto L1055
	default:
		goto L1056
	}
L1058:
	;
	goto L1059
L1059:
	;
	if v3418&int32(1) != 0 {
		goto L1055
	} else {
		goto L1071
	}
L1060:
	;
	if v3418&int32(1) == int32(0) {
		goto L1055
	} else {
		goto L1061
	}
L1061:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L66
	} else {
		goto L1062
	}
L1062:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3478 = m.ExcPending
	if v3478 != 0 {
		goto L66
	} else {
		goto L1063
	}
L1063:
	;
	v3479 = *(*int32)(unsafe.Add(mBase, uint32(v3462)))
	if base.Ui32(int32(12)) < base.Ui32(v3479) {
		goto L1065
	} else {
		goto L1066
	}
L1064:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v3488
	F_errmsg(m, int32(523589), v27+int32(144))
	mBase = m.M
	v3494 = m.ExcPending
	if v3494 != 0 {
		goto L66
	} else {
		goto L1068
	}
L1065:
	;
	v3488 = int32(243550)
	goto L1064
L1066:
	;
	goto L1067
L1067:
	;
	v3487 = *(*int32)(unsafe.Add(mBase, uint32(v3479<<(uint(int32(2))%32))+uint32(_consts[1382])))
	v3488 = v3487
	goto L1064
L1068:
	;
	v3495 = *(*int32)(unsafe.Add(mBase, uint32(v3360)))
	v3496 = F_plpgsql_scanner_errposition(m, v3495, l1)
	mBase = m.M
	v3497 = m.ExcPending
	if v3497 != 0 {
		goto L66
	} else {
		goto L1069
	}
L1069:
	;
	F_errfinish(m, int32(26959), int32(1042), int32(360630))
	mBase = m.M
	v3502 = m.ExcPending
	if v3502 != 0 {
		goto L66
	} else {
		goto L1070
	}
L1070:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1071:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v3508 = m.ExcPending
	if v3508 != 0 {
		goto L66
	} else {
		goto L1072
	}
L1072:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		goto L66
	} else {
		goto L1073
	}
L1073:
	;
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v3462)))
	if base.Ui32(int32(12)) < base.Ui32(v3512) {
		goto L1075
	} else {
		goto L1076
	}
L1074:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v3521
	F_errmsg(m, int32(523511), v27+int32(160))
	mBase = m.M
	v3527 = m.ExcPending
	if v3527 != 0 {
		goto L66
	} else {
		goto L1078
	}
L1075:
	;
	v3521 = int32(243550)
	goto L1074
L1076:
	;
	goto L1077
L1077:
	;
	v3520 = *(*int32)(unsafe.Add(mBase, uint32(v3512<<(uint(int32(2))%32))+uint32(_consts[1382])))
	v3521 = v3520
	goto L1074
L1078:
	;
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3360)))
	v3529 = F_plpgsql_scanner_errposition(m, v3528, l1)
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L66
	} else {
		goto L1079
	}
L1079:
	;
	F_errfinish(m, int32(26959), int32(1060), int32(360630))
	mBase = m.M
	v3535 = m.ExcPending
	if v3535 != 0 {
		goto L66
	} else {
		goto L1080
	}
L1080:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1081:
	;
	v3540 = *(*int32)(unsafe.Add(mBase, uint32(v3462)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v3540
	F_errmsg_internal(m, int32(487206), v27+int32(128))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L66
	} else {
		goto L1082
	}
L1082:
	;
	F_errfinish(m, int32(26959), int32(1067), int32(360630))
	mBase = m.M
	v3551 = m.ExcPending
	if v3551 != 0 {
		goto L66
	} else {
		goto L1083
	}
L1083:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1084:
	;
	goto L1054
L1085:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3590
	v10149 = v241
	goto L5
L1086:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3599
	v10149 = v241
	goto L5
L1087:
	;
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	v3608 = *(*int32)(unsafe.Add(mBase, uint32(v3607)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3603)+4)) = v3608
	v3610 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v3603))) = v3610
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3603
	v10149 = v241
	goto L5
L1088:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(291096))
	mBase = m.M
	v4026 = m.ExcPending
	if v4026 != 0 {
		goto L66
	} else {
		goto L1241
	}
L1089:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(6)
	v10149 = v241
	goto L5
L1090:
	;
	v3988 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v3988 == int32(0) {
		goto L1088
	} else {
		goto L1231
	}
L1091:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(12)
	v10149 = v241
	goto L5
L1092:
	;
	v3959 = int32(380851)
	v3962 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1383])))
	v3963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927))))
	if v3963 == int32(0) {
		v3982 = v3962
		v3983 = v3963
		goto L1223
	} else {
		goto L1224
	}
L1093:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(11)
	v10149 = v241
	goto L5
L1094:
	;
	v3927 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v3927 == int32(0) {
		goto L1088
	} else {
		goto L1212
	}
L1095:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(10)
	v10149 = v241
	goto L5
L1096:
	;
	v3898 = int32(63226)
	v3901 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1384])))
	v3902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3866))))
	if v3902 == int32(0) {
		v3921 = v3901
		v3922 = v3902
		goto L1204
	} else {
		goto L1205
	}
L1097:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(9)
	v10149 = v241
	goto L5
L1098:
	;
	v3866 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v3866 == int32(0) {
		goto L1088
	} else {
		goto L1193
	}
L1099:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(8)
	v10149 = v241
	goto L5
L1100:
	;
	v3837 = int32(378374)
	v3840 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1385])))
	v3841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3805))))
	if v3841 == int32(0) {
		v3860 = v3840
		v3861 = v3841
		goto L1185
	} else {
		goto L1186
	}
L1101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(7)
	v10149 = v241
	goto L5
L1102:
	;
	v3805 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v3805 == int32(0) {
		goto L1088
	} else {
		goto L1174
	}
L1103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(3)
	v10149 = v241
	goto L5
L1104:
	;
	v3776 = int32(59460)
	v3779 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1386])))
	v3780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3744))))
	if v3780 == int32(0) {
		v3799 = v3779
		v3800 = v3780
		goto L1166
	} else {
		goto L1167
	}
L1105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(5)
	v10149 = v241
	goto L5
L1106:
	;
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v3744 == int32(0) {
		goto L1088
	} else {
		goto L1155
	}
L1107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(4)
	v10149 = v241
	goto L5
L1108:
	;
	v3715 = int32(305455)
	v3718 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1387])))
	v3719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3683))))
	if v3719 == int32(0) {
		v3738 = v3718
		v3739 = v3719
		goto L1147
	} else {
		goto L1148
	}
L1109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(2)
	v10149 = v241
	goto L5
L1110:
	;
	v3683 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v3683 == int32(0) {
		goto L1088
	} else {
		goto L1136
	}
L1111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(1)
	v10149 = v241
	goto L5
L1112:
	;
	v3654 = int32(434738)
	v3657 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1388])))
	v3658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3622))))
	if v3658 == int32(0) {
		v3677 = v3657
		v3678 = v3658
		goto L1128
	} else {
		goto L1129
	}
L1113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L1114:
	;
	v3621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v3621 != 0 {
		goto L1088
	} else {
		goto L1116
	}
L1115:
	;
	switch v3617 - int32(277) {
	case 0:
		goto L1114
	default:
		goto L1088
	case 18:
		goto L1101
	case 22:
		goto L1099
	case 62:
		goto L1095
	case 73:
		goto L1109
	case 74:
		goto L1097
	case 75:
		goto L1103
	case 76:
		goto L1107
	case 77:
		goto L1105
	case 78:
		goto L1111
	case 85:
		goto L1089
	case 88:
		goto L1113
	case 91:
		goto L1091
	case 98:
		goto L1093
	}
L1116:
	;
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v3622 == int32(0) {
		goto L1088
	} else {
		goto L1117
	}
L1117:
	;
	v3625 = int32(87383)
	v3628 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1389])))
	v3629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3622))))
	if v3629 == int32(0) {
		v3648 = v3628
		v3649 = v3629
		goto L1119
	} else {
		goto L1120
	}
L1118:
	;
	if v3649-v3648 != 0 {
		goto L1112
	} else {
		goto L1126
	}
L1119:
	;
	goto L1118
L1120:
	;
	if v3628 != v3629 {
		v3648 = v3628
		v3649 = v3629
		goto L1119
	} else {
		goto L1121
	}
L1121:
	;
	v3633 = v3622
	v3634 = v3625
	goto L1122
L1122:
	;
	v3637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3634)+1)))
	v3638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3633)+1)))
	if v3638 == int32(0) {
		v3648 = v3637
		v3649 = v3638
		goto L1119
	} else {
		goto L1124
	}
L1123:
	;
	v3648 = v3637
	v3649 = v3638
	goto L1119
L1124:
	;
	v3641 = int32(1)
	if v3637 == v3638 {
		v3633 = v3633 + v3641
		v3634 = v3634 + v3641
		goto L1122
	} else {
		goto L1125
	}
L1125:
	;
	goto L1123
L1126:
	;
	goto L1113
L1127:
	;
	if v3678-v3677 != 0 {
		goto L1110
	} else {
		goto L1135
	}
L1128:
	;
	goto L1127
L1129:
	;
	if v3657 != v3658 {
		v3677 = v3657
		v3678 = v3658
		goto L1128
	} else {
		goto L1130
	}
L1130:
	;
	v3662 = v3622
	v3663 = v3654
	goto L1131
L1131:
	;
	v3666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3663)+1)))
	v3667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3662)+1)))
	if v3667 == int32(0) {
		v3677 = v3666
		v3678 = v3667
		goto L1128
	} else {
		goto L1133
	}
L1132:
	;
	v3677 = v3666
	v3678 = v3667
	goto L1128
L1133:
	;
	v3670 = int32(1)
	if v3666 == v3667 {
		v3662 = v3662 + v3670
		v3663 = v3663 + v3670
		goto L1131
	} else {
		goto L1134
	}
L1134:
	;
	goto L1132
L1135:
	;
	goto L1111
L1136:
	;
	v3686 = int32(59481)
	v3689 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1390])))
	v3690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3683))))
	if v3690 == int32(0) {
		v3709 = v3689
		v3710 = v3690
		goto L1138
	} else {
		goto L1139
	}
L1137:
	;
	if v3710-v3709 != 0 {
		goto L1108
	} else {
		goto L1145
	}
L1138:
	;
	goto L1137
L1139:
	;
	if v3689 != v3690 {
		v3709 = v3689
		v3710 = v3690
		goto L1138
	} else {
		goto L1140
	}
L1140:
	;
	v3694 = v3683
	v3695 = v3686
	goto L1141
L1141:
	;
	v3698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3695)+1)))
	v3699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3694)+1)))
	if v3699 == int32(0) {
		v3709 = v3698
		v3710 = v3699
		goto L1138
	} else {
		goto L1143
	}
L1142:
	;
	v3709 = v3698
	v3710 = v3699
	goto L1138
L1143:
	;
	v3702 = int32(1)
	if v3698 == v3699 {
		v3694 = v3694 + v3702
		v3695 = v3695 + v3702
		goto L1141
	} else {
		goto L1144
	}
L1144:
	;
	goto L1142
L1145:
	;
	goto L1109
L1146:
	;
	if v3739-v3738 != 0 {
		goto L1106
	} else {
		goto L1154
	}
L1147:
	;
	goto L1146
L1148:
	;
	if v3718 != v3719 {
		v3738 = v3718
		v3739 = v3719
		goto L1147
	} else {
		goto L1149
	}
L1149:
	;
	v3723 = v3683
	v3724 = v3715
	goto L1150
L1150:
	;
	v3727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3724)+1)))
	v3728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3723)+1)))
	if v3728 == int32(0) {
		v3738 = v3727
		v3739 = v3728
		goto L1147
	} else {
		goto L1152
	}
L1151:
	;
	v3738 = v3727
	v3739 = v3728
	goto L1147
L1152:
	;
	v3731 = int32(1)
	if v3727 == v3728 {
		v3723 = v3723 + v3731
		v3724 = v3724 + v3731
		goto L1150
	} else {
		goto L1153
	}
L1153:
	;
	goto L1151
L1154:
	;
	goto L1107
L1155:
	;
	v3747 = int32(89106)
	v3750 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1391])))
	v3751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3744))))
	if v3751 == int32(0) {
		v3770 = v3750
		v3771 = v3751
		goto L1157
	} else {
		goto L1158
	}
L1156:
	;
	if v3771-v3770 != 0 {
		goto L1104
	} else {
		goto L1164
	}
L1157:
	;
	goto L1156
L1158:
	;
	if v3750 != v3751 {
		v3770 = v3750
		v3771 = v3751
		goto L1157
	} else {
		goto L1159
	}
L1159:
	;
	v3755 = v3744
	v3756 = v3747
	goto L1160
L1160:
	;
	v3759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3756)+1)))
	v3760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3755)+1)))
	if v3760 == int32(0) {
		v3770 = v3759
		v3771 = v3760
		goto L1157
	} else {
		goto L1162
	}
L1161:
	;
	v3770 = v3759
	v3771 = v3760
	goto L1157
L1162:
	;
	v3763 = int32(1)
	if v3759 == v3760 {
		v3755 = v3755 + v3763
		v3756 = v3756 + v3763
		goto L1160
	} else {
		goto L1163
	}
L1163:
	;
	goto L1161
L1164:
	;
	goto L1105
L1165:
	;
	if v3800-v3799 != 0 {
		goto L1102
	} else {
		goto L1173
	}
L1166:
	;
	goto L1165
L1167:
	;
	if v3779 != v3780 {
		v3799 = v3779
		v3800 = v3780
		goto L1166
	} else {
		goto L1168
	}
L1168:
	;
	v3784 = v3744
	v3785 = v3776
	goto L1169
L1169:
	;
	v3788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3785)+1)))
	v3789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3784)+1)))
	if v3789 == int32(0) {
		v3799 = v3788
		v3800 = v3789
		goto L1166
	} else {
		goto L1171
	}
L1170:
	;
	v3799 = v3788
	v3800 = v3789
	goto L1166
L1171:
	;
	v3792 = int32(1)
	if v3788 == v3789 {
		v3784 = v3784 + v3792
		v3785 = v3785 + v3792
		goto L1169
	} else {
		goto L1172
	}
L1172:
	;
	goto L1170
L1173:
	;
	goto L1103
L1174:
	;
	v3808 = int32(378934)
	v3811 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1392])))
	v3812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3805))))
	if v3812 == int32(0) {
		v3831 = v3811
		v3832 = v3812
		goto L1176
	} else {
		goto L1177
	}
L1175:
	;
	if v3832-v3831 != 0 {
		goto L1100
	} else {
		goto L1183
	}
L1176:
	;
	goto L1175
L1177:
	;
	if v3811 != v3812 {
		v3831 = v3811
		v3832 = v3812
		goto L1176
	} else {
		goto L1178
	}
L1178:
	;
	v3816 = v3805
	v3817 = v3808
	goto L1179
L1179:
	;
	v3820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3817)+1)))
	v3821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3816)+1)))
	if v3821 == int32(0) {
		v3831 = v3820
		v3832 = v3821
		goto L1176
	} else {
		goto L1181
	}
L1180:
	;
	v3831 = v3820
	v3832 = v3821
	goto L1176
L1181:
	;
	v3824 = int32(1)
	if v3820 == v3821 {
		v3816 = v3816 + v3824
		v3817 = v3817 + v3824
		goto L1179
	} else {
		goto L1182
	}
L1182:
	;
	goto L1180
L1183:
	;
	goto L1101
L1184:
	;
	if v3861-v3860 != 0 {
		goto L1098
	} else {
		goto L1192
	}
L1185:
	;
	goto L1184
L1186:
	;
	if v3840 != v3841 {
		v3860 = v3840
		v3861 = v3841
		goto L1185
	} else {
		goto L1187
	}
L1187:
	;
	v3845 = v3805
	v3846 = v3837
	goto L1188
L1188:
	;
	v3849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3846)+1)))
	v3850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3845)+1)))
	if v3850 == int32(0) {
		v3860 = v3849
		v3861 = v3850
		goto L1185
	} else {
		goto L1190
	}
L1189:
	;
	v3860 = v3849
	v3861 = v3850
	goto L1185
L1190:
	;
	v3853 = int32(1)
	if v3849 == v3850 {
		v3845 = v3845 + v3853
		v3846 = v3846 + v3853
		goto L1188
	} else {
		goto L1191
	}
L1191:
	;
	goto L1189
L1192:
	;
	goto L1099
L1193:
	;
	v3869 = int32(379177)
	v3872 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1393])))
	v3873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3866))))
	if v3873 == int32(0) {
		v3892 = v3872
		v3893 = v3873
		goto L1195
	} else {
		goto L1196
	}
L1194:
	;
	if v3893-v3892 != 0 {
		goto L1096
	} else {
		goto L1202
	}
L1195:
	;
	goto L1194
L1196:
	;
	if v3872 != v3873 {
		v3892 = v3872
		v3893 = v3873
		goto L1195
	} else {
		goto L1197
	}
L1197:
	;
	v3877 = v3866
	v3878 = v3869
	goto L1198
L1198:
	;
	v3881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3878)+1)))
	v3882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3877)+1)))
	if v3882 == int32(0) {
		v3892 = v3881
		v3893 = v3882
		goto L1195
	} else {
		goto L1200
	}
L1199:
	;
	v3892 = v3881
	v3893 = v3882
	goto L1195
L1200:
	;
	v3885 = int32(1)
	if v3881 == v3882 {
		v3877 = v3877 + v3885
		v3878 = v3878 + v3885
		goto L1198
	} else {
		goto L1201
	}
L1201:
	;
	goto L1199
L1202:
	;
	goto L1097
L1203:
	;
	if v3922-v3921 != 0 {
		goto L1094
	} else {
		goto L1211
	}
L1204:
	;
	goto L1203
L1205:
	;
	if v3901 != v3902 {
		v3921 = v3901
		v3922 = v3902
		goto L1204
	} else {
		goto L1206
	}
L1206:
	;
	v3906 = v3866
	v3907 = v3898
	goto L1207
L1207:
	;
	v3910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3907)+1)))
	v3911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3906)+1)))
	if v3911 == int32(0) {
		v3921 = v3910
		v3922 = v3911
		goto L1204
	} else {
		goto L1209
	}
L1208:
	;
	v3921 = v3910
	v3922 = v3911
	goto L1204
L1209:
	;
	v3914 = int32(1)
	if v3910 == v3911 {
		v3906 = v3906 + v3914
		v3907 = v3907 + v3914
		goto L1207
	} else {
		goto L1210
	}
L1210:
	;
	goto L1208
L1211:
	;
	goto L1095
L1212:
	;
	v3930 = int32(379845)
	v3933 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
	v3934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927))))
	if v3934 == int32(0) {
		v3953 = v3933
		v3954 = v3934
		goto L1214
	} else {
		goto L1215
	}
L1213:
	;
	if v3954-v3953 != 0 {
		goto L1092
	} else {
		goto L1221
	}
L1214:
	;
	goto L1213
L1215:
	;
	if v3933 != v3934 {
		v3953 = v3933
		v3954 = v3934
		goto L1214
	} else {
		goto L1216
	}
L1216:
	;
	v3938 = v3927
	v3939 = v3930
	goto L1217
L1217:
	;
	v3942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3939)+1)))
	v3943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3938)+1)))
	if v3943 == int32(0) {
		v3953 = v3942
		v3954 = v3943
		goto L1214
	} else {
		goto L1219
	}
L1218:
	;
	v3953 = v3942
	v3954 = v3943
	goto L1214
L1219:
	;
	v3946 = int32(1)
	if v3942 == v3943 {
		v3938 = v3938 + v3946
		v3939 = v3939 + v3946
		goto L1217
	} else {
		goto L1220
	}
L1220:
	;
	goto L1218
L1221:
	;
	goto L1093
L1222:
	;
	if v3983-v3982 != 0 {
		goto L1090
	} else {
		goto L1230
	}
L1223:
	;
	goto L1222
L1224:
	;
	if v3962 != v3963 {
		v3982 = v3962
		v3983 = v3963
		goto L1223
	} else {
		goto L1225
	}
L1225:
	;
	v3967 = v3927
	v3968 = v3959
	goto L1226
L1226:
	;
	v3971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3968)+1)))
	v3972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3967)+1)))
	if v3972 == int32(0) {
		v3982 = v3971
		v3983 = v3972
		goto L1223
	} else {
		goto L1228
	}
L1227:
	;
	v3982 = v3971
	v3983 = v3972
	goto L1223
L1228:
	;
	v3975 = int32(1)
	if v3971 == v3972 {
		v3967 = v3967 + v3975
		v3968 = v3968 + v3975
		goto L1226
	} else {
		goto L1229
	}
L1229:
	;
	goto L1227
L1230:
	;
	goto L1091
L1231:
	;
	v3991 = int32(351302)
	v3994 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1395])))
	v3995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3988))))
	if v3995 == int32(0) {
		v4014 = v3994
		v4015 = v3995
		goto L1233
	} else {
		goto L1234
	}
L1232:
	;
	if v4015-v4014 != 0 {
		goto L1088
	} else {
		goto L1240
	}
L1233:
	;
	goto L1232
L1234:
	;
	if v3994 != v3995 {
		v4014 = v3994
		v4015 = v3995
		goto L1233
	} else {
		goto L1235
	}
L1235:
	;
	v3999 = v3988
	v4000 = v3991
	goto L1236
L1236:
	;
	v4003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4000)+1)))
	v4004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3999)+1)))
	if v4004 == int32(0) {
		v4014 = v4003
		v4015 = v4004
		goto L1233
	} else {
		goto L1238
	}
L1237:
	;
	v4014 = v4003
	v4015 = v4004
	goto L1233
L1238:
	;
	v4007 = int32(1)
	if v4003 == v4004 {
		v3999 = v3999 + v4007
		v4000 = v4000 + v4007
		goto L1236
	} else {
		goto L1239
	}
L1239:
	;
	goto L1237
L1240:
	;
	goto L1089
L1241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1242:
	;
	v4033 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v4034 = m.ExcPending
	if v4034 != 0 {
		goto L66
	} else {
		goto L1243
	}
L1243:
	;
	if v4033 == int32(91) {
		goto L44
	} else {
		goto L1244
	}
L1244:
	;
	v4037 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v4038 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_check_assignable(m, v4037, v4038, l1)
	mBase = m.M
	v4040 = m.ExcPending
	if v4040 != 0 {
		goto L66
	} else {
		goto L1245
	}
L1245:
	;
	v4041 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4041
	v10149 = v241
	goto L5
L1246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4050))) = int32(2)
	v4056 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(28))))
	v4057 = int32(0)
	if v4056 < v4057 {
		v4100 = v4057
		goto L1250
	} else {
		goto L1251
	}
L1249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4050)+4)) = v4100
	v4105 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v4106 = *(*int32)(unsafe.Add(mBase, uint32(v4105)+520))
	v4108 = v4106 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4105)+520)) = v4108
	*(*int32)(unsafe.Add(mBase, uint32(v4050)+8)) = v4108
	v4113 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(96))))
	*(*int32)(unsafe.Add(mBase, uint32(v4050)+12)) = v4113
	v4117 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(80))))
	*(*int32)(unsafe.Add(mBase, uint32(v4050)+16)) = v4117
	v4121 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v4050)+20)) = v4121
	v4125 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	*(*int32)(unsafe.Add(mBase, uint32(v4050)+24)) = v4125
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4050
	v10149 = v241
	goto L5
L1250:
	;
	goto L1249
L1251:
	;
	v4062 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4063 = *(*int32)(unsafe.Add(mBase, uint32(v4062)+60))
	if v4063 == int32(0) {
		v4100 = v4057
		goto L1250
	} else {
		goto L1252
	}
L1252:
	;
	v4066 = v4056 + v4063
	v4067 = *(*int32)(unsafe.Add(mBase, uint32(v4062)+188))
	if base.Ui32(v4067) <= base.Ui32(v4066) {
		goto L1254
	} else {
		goto L1255
	}
L1253:
	;
	v4077 = *(*int32)(unsafe.Add(mBase, uint32(v4062)+196))
	if v4076 == int32(0) {
		v4100 = v4077
		goto L1250
	} else {
		goto L1257
	}
L1254:
	;
	v4069 = *(*int32)(unsafe.Add(mBase, uint32(v4062)+192))
	v4076 = v4069
	goto L1253
L1255:
	;
	goto L1256
L1256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4062)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4062)+188)) = v4063
	v4074 = F_strchr(m, v4063, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4062)+192)) = v4074
	v4076 = v4074
	goto L1253
L1257:
	;
	if base.Ui32(v4066) <= base.Ui32(v4076) {
		v4100 = v4077
		goto L1250
	} else {
		goto L1258
	}
L1258:
	;
	v4081 = v4076
	v4083 = v4077
	goto L1259
L1259:
	;
	v4086 = int32(1)
	v4087 = v4083 + v4086
	*(*int32)(unsafe.Add(mBase, uint32(v4062)+196)) = v4087
	v4090 = v4081 + v4086
	*(*int32)(unsafe.Add(mBase, uint32(v4062)+188)) = v4090
	v4093 = F_strchr(m, v4090, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4062)+192)) = v4093
	if v4093 == int32(0) {
		v4100 = v4087
		goto L1250
	} else {
		goto L1261
	}
L1260:
	;
	v4100 = v4087
	goto L1250
L1261:
	;
	if base.Ui32(v4093) < base.Ui32(v4066) {
		v4081 = v4093
		v4083 = v4087
		goto L1259
	} else {
		goto L1262
	}
L1262:
	;
	goto L1260
L1263:
	;
	v4135 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4136 = int32(0)
	if v4135 < v4136 {
		v4179 = v4136
		goto L1265
	} else {
		goto L1266
	}
L1264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4131))) = v4179
	v4185 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4131)+4)) = v4185
	v4187 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4131)+8)) = v4187
	v4191 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	v4192 = F_lappend(m, v4191, v4131)
	mBase = m.M
	v4193 = m.ExcPending
	if v4193 != 0 {
		goto L66
	} else {
		goto L1278
	}
L1265:
	;
	goto L1264
L1266:
	;
	v4141 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4142 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+60))
	if v4142 == int32(0) {
		v4179 = v4136
		goto L1265
	} else {
		goto L1267
	}
L1267:
	;
	v4145 = v4135 + v4142
	v4146 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+188))
	if base.Ui32(v4146) <= base.Ui32(v4145) {
		goto L1269
	} else {
		goto L1270
	}
L1268:
	;
	v4156 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+196))
	if v4155 == int32(0) {
		v4179 = v4156
		goto L1265
	} else {
		goto L1272
	}
L1269:
	;
	v4148 = *(*int32)(unsafe.Add(mBase, uint32(v4141)+192))
	v4155 = v4148
	goto L1268
L1270:
	;
	goto L1271
L1271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+188)) = v4142
	v4153 = F_strchr(m, v4142, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+192)) = v4153
	v4155 = v4153
	goto L1268
L1272:
	;
	if base.Ui32(v4145) <= base.Ui32(v4155) {
		v4179 = v4156
		goto L1265
	} else {
		goto L1273
	}
L1273:
	;
	v4160 = v4155
	v4162 = v4156
	goto L1274
L1274:
	;
	v4165 = int32(1)
	v4166 = v4162 + v4165
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+196)) = v4166
	v4169 = v4160 + v4165
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+188)) = v4169
	v4172 = F_strchr(m, v4169, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4141)+192)) = v4172
	if v4172 == int32(0) {
		v4179 = v4166
		goto L1265
	} else {
		goto L1276
	}
L1275:
	;
	v4179 = v4166
	goto L1265
L1276:
	;
	if base.Ui32(v4172) < base.Ui32(v4145) {
		v4160 = v4172
		v4162 = v4166
		goto L1274
	} else {
		goto L1277
	}
L1277:
	;
	goto L1275
L1278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4192
	v10149 = v241
	goto L5
L1279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4212))) = int32(3)
	v4216 = int32(0)
	if v4210 < v4216 {
		v4259 = v4216
		goto L1281
	} else {
		goto L1282
	}
L1280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4212)+4)) = v4259
	v4264 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4264)+520))
	v4267 = v4265 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4264)+520)) = v4267
	v4269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4212)+24)) = uint8(base.B2i32(v4207 != v4269))
	*(*int32)(unsafe.Add(mBase, uint32(v4212)+20)) = v4204
	*(*int32)(unsafe.Add(mBase, uint32(v4212)+16)) = v4269
	*(*int32)(unsafe.Add(mBase, uint32(v4212)+12)) = v4201
	*(*int32)(unsafe.Add(mBase, uint32(v4212)+8)) = v4267
	if v4207 == v4269 {
		v4285 = v4207
		goto L1294
	} else {
		goto L1295
	}
L1281:
	;
	goto L1280
L1282:
	;
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4222 = *(*int32)(unsafe.Add(mBase, uint32(v4221)+60))
	if v4222 == int32(0) {
		v4259 = v4216
		goto L1281
	} else {
		goto L1283
	}
L1283:
	;
	v4225 = v4210 + v4222
	v4226 = *(*int32)(unsafe.Add(mBase, uint32(v4221)+188))
	if base.Ui32(v4226) <= base.Ui32(v4225) {
		goto L1285
	} else {
		goto L1286
	}
L1284:
	;
	v4236 = *(*int32)(unsafe.Add(mBase, uint32(v4221)+196))
	if v4235 == int32(0) {
		v4259 = v4236
		goto L1281
	} else {
		goto L1288
	}
L1285:
	;
	v4228 = *(*int32)(unsafe.Add(mBase, uint32(v4221)+192))
	v4235 = v4228
	goto L1284
L1286:
	;
	goto L1287
L1287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4221)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4221)+188)) = v4222
	v4233 = F_strchr(m, v4222, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4221)+192)) = v4233
	v4235 = v4233
	goto L1284
L1288:
	;
	if base.Ui32(v4225) <= base.Ui32(v4235) {
		v4259 = v4236
		goto L1281
	} else {
		goto L1289
	}
L1289:
	;
	v4240 = v4235
	v4242 = v4236
	goto L1290
L1290:
	;
	v4245 = int32(1)
	v4246 = v4242 + v4245
	*(*int32)(unsafe.Add(mBase, uint32(v4221)+196)) = v4246
	v4249 = v4240 + v4245
	*(*int32)(unsafe.Add(mBase, uint32(v4221)+188)) = v4249
	v4252 = F_strchr(m, v4249, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4221)+192)) = v4252
	if v4252 == int32(0) {
		v4259 = v4246
		goto L1281
	} else {
		goto L1292
	}
L1291:
	;
	v4259 = v4246
	goto L1281
L1292:
	;
	if base.Ui32(v4252) < base.Ui32(v4225) {
		v4240 = v4252
		v4242 = v4246
		goto L1290
	} else {
		goto L1293
	}
L1293:
	;
	goto L1291
L1294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4212)+28)) = v4285
	if v4201 == int32(0) {
		goto L1298
	} else {
		goto L1299
	}
L1295:
	;
	v4279 = *(*int32)(unsafe.Add(mBase, uint32(v4207)+4))
	if v4279 != int32(1) {
		v4285 = v4207
		goto L1294
	} else {
		goto L1296
	}
L1296:
	;
	v4282 = *(*int32)(unsafe.Add(mBase, uint32(v4207)+12))
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(v4282)))
	if v4283 != 0 {
		v4285 = v4207
		goto L1294
	} else {
		goto L1297
	}
L1297:
	;
	v4285 = int32(0)
	goto L1294
L1298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4212
	v10149 = v241
	goto L5
L1299:
	;
	v4290 = *(*int32)(unsafe.Add(mBase, _consts[1347]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v4290
	v4298 = F_pg_snprintf(m, v27+int32(4784), int32(32), int32(507344), v27+int32(208))
	mBase = m.M
	v4299 = m.ExcPending
	if v4299 != 0 {
		goto L66
	} else {
		goto L1300
	}
L1300:
	;
	v4300 = int32(0)
	v4303 = *(*int32)(unsafe.Add(mBase, uint32(v4212)+4))
	v4308 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v4300, v4300)
	mBase = m.M
	v4309 = m.ExcPending
	if v4309 != 0 {
		goto L66
	} else {
		goto L1301
	}
L1301:
	;
	v4311 = F_plpgsql_build_variable(m, v27+int32(4784), v4303, v4308, int32(1))
	mBase = m.M
	v4312 = m.ExcPending
	if v4312 != 0 {
		goto L66
	} else {
		goto L1302
	}
L1302:
	;
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(v4311)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4212)+16)) = v4313
	if v4204 == int32(0) {
		goto L1298
	} else {
		goto L1303
	}
L1303:
	;
	v4317 = *(*int32)(unsafe.Add(mBase, uint32(v4204)+4))
	if v4317 <= int32(0) {
		goto L1298
	} else {
		goto L1304
	}
L1304:
	;
	v4324 = v4300
	goto L1305
L1305:
	;
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(v4204)+12))
	v4348 = *(*int32)(unsafe.Add(mBase, uint32(v4344+v4324<<(uint(int32(2))%32))))
	v4349 = *(*int32)(unsafe.Add(mBase, uint32(v4348)+4))
	F_initStringInfo(m, v27+int32(4768))
	mBase = m.M
	v4353 = m.ExcPending
	if v4353 != 0 {
		goto L66
	} else {
		goto L1307
	}
L1306:
	;
	goto L1298
L1307:
	;
	v4354 = *(*int32)(unsafe.Add(mBase, uint32(v4349)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+196)) = v4354
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v27 + int32(4784)
	F_appendStringInfo(m, v27+int32(4768), int32(675364), v27+int32(192))
	mBase = m.M
	v4365 = m.ExcPending
	if v4365 != 0 {
		goto L66
	} else {
		goto L1308
	}
L1308:
	;
	v4366 = *(*int32)(unsafe.Add(mBase, uint32(v4349)))
	F_pfree(m, v4366)
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		goto L66
	} else {
		goto L1309
	}
L1309:
	;
	v4369 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379])))
	v4370 = F_pstrdup(m, v4369)
	mBase = m.M
	v4371 = m.ExcPending
	if v4371 != 0 {
		goto L66
	} else {
		goto L1310
	}
L1310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4349))) = v4370
	v4374 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	*(*int32)(unsafe.Add(mBase, uint32(v4349)+12)) = v4374
	v4376 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379])))
	F_pfree(m, v4376)
	mBase = m.M
	v4378 = m.ExcPending
	if v4378 != 0 {
		goto L66
	} else {
		goto L1311
	}
L1311:
	;
	v4380 = v4324 + int32(1)
	v4381 = *(*int32)(unsafe.Add(mBase, uint32(v4204)+4))
	if v4380 < v4381 {
		v4324 = v4380
		goto L1305
	} else {
		goto L1312
	}
L1312:
	;
	goto L1306
L1313:
	;
	if v4413 != int32(384) {
		goto L1314
	} else {
		goto L1315
	}
L1314:
	;
	F_plpgsql_push_back_token(m, v4413, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4422 = m.ExcPending
	if v4422 != 0 {
		goto L66
	} else {
		goto L1317
	}
L1315:
	;
	v4438 = int32(0)
	goto L1316
L1316:
	;
	F_plpgsql_push_back_token(m, int32(384), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4445 = m.ExcPending
	if v4445 != 0 {
		goto L66
	} else {
		goto L1319
	}
L1317:
	;
	v4424 = int32(0)
	v4428 = int32(1)
	v4436 = F_read_sql_construct(m, int32(384), v4424, v4424, int32(530600), int32(2), v4428, v4428, v4424, v4424, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4437 = m.ExcPending
	if v4437 != 0 {
		goto L66
	} else {
		goto L1318
	}
L1318:
	;
	v4438 = v4436
	goto L1316
L1319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4438
	v10149 = v241
	goto L5
L1320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4451
	v10149 = v241
	goto L5
L1321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4460
	v10149 = v241
	goto L5
L1322:
	;
	v4468 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4469 = int32(0)
	if v4468 < v4469 {
		v4512 = v4469
		goto L1324
	} else {
		goto L1325
	}
L1323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4464))) = v4512
	v4518 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4464)+4)) = v4518
	v4520 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4464)+8)) = v4520
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4464
	v10149 = v241
	goto L5
L1324:
	;
	goto L1323
L1325:
	;
	v4474 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4475 = *(*int32)(unsafe.Add(mBase, uint32(v4474)+60))
	if v4475 == int32(0) {
		v4512 = v4469
		goto L1324
	} else {
		goto L1326
	}
L1326:
	;
	v4478 = v4468 + v4475
	v4479 = *(*int32)(unsafe.Add(mBase, uint32(v4474)+188))
	if base.Ui32(v4479) <= base.Ui32(v4478) {
		goto L1328
	} else {
		goto L1329
	}
L1327:
	;
	v4489 = *(*int32)(unsafe.Add(mBase, uint32(v4474)+196))
	if v4488 == int32(0) {
		v4512 = v4489
		goto L1324
	} else {
		goto L1331
	}
L1328:
	;
	v4481 = *(*int32)(unsafe.Add(mBase, uint32(v4474)+192))
	v4488 = v4481
	goto L1327
L1329:
	;
	goto L1330
L1330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4474)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4474)+188)) = v4475
	v4486 = F_strchr(m, v4475, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4474)+192)) = v4486
	v4488 = v4486
	goto L1327
L1331:
	;
	if base.Ui32(v4478) <= base.Ui32(v4488) {
		v4512 = v4489
		goto L1324
	} else {
		goto L1332
	}
L1332:
	;
	v4493 = v4488
	v4495 = v4489
	goto L1333
L1333:
	;
	v4498 = int32(1)
	v4499 = v4495 + v4498
	*(*int32)(unsafe.Add(mBase, uint32(v4474)+196)) = v4499
	v4502 = v4493 + v4498
	*(*int32)(unsafe.Add(mBase, uint32(v4474)+188)) = v4502
	v4505 = F_strchr(m, v4502, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4474)+192)) = v4505
	if v4505 == int32(0) {
		v4512 = v4499
		goto L1324
	} else {
		goto L1335
	}
L1334:
	;
	v4512 = v4499
	goto L1324
L1335:
	;
	if base.Ui32(v4505) < base.Ui32(v4478) {
		v4493 = v4505
		v4495 = v4499
		goto L1333
	} else {
		goto L1336
	}
L1336:
	;
	goto L1334
L1337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4525
	v10149 = v241
	goto L5
L1338:
	;
	goto L1339
L1339:
	;
	v4527 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v4527
	*(*int32)(unsafe.Add(mBase, uint32(v27)+280)) = v4527
	v4534 = F_list_make1_impl(m, int32(1), v27+int32(220))
	mBase = m.M
	v4535 = m.ExcPending
	if v4535 != 0 {
		goto L66
	} else {
		goto L1340
	}
L1340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4534
	v10149 = v241
	goto L5
L1341:
	;
	v4540 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4538))) = v4540
	v4544 = *(*int32)(unsafe.Add(mBase, uint32(v142-v4540)))
	v4545 = int32(0)
	if v4544 < v4545 {
		v4588 = v4545
		goto L1343
	} else {
		goto L1344
	}
L1342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4538)+4)) = v4588
	v4593 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v4594 = *(*int32)(unsafe.Add(mBase, uint32(v4593)+520))
	v4596 = v4594 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4593)+520)) = v4596
	*(*int32)(unsafe.Add(mBase, uint32(v4538)+8)) = v4596
	v4600 = v132 - int32(32)
	v4601 = *(*int32)(unsafe.Add(mBase, uint32(v4600)))
	*(*int32)(unsafe.Add(mBase, uint32(v4538)+12)) = v4601
	v4603 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4538)+16)) = v4603
	v4605 = *(*int32)(unsafe.Add(mBase, uint32(v4600)))
	v4606 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v4607 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v4605, v4606, v4607, l1)
	mBase = m.M
	v4609 = m.ExcPending
	if v4609 != 0 {
		goto L66
	} else {
		goto L1356
	}
L1343:
	;
	goto L1342
L1344:
	;
	v4550 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4551 = *(*int32)(unsafe.Add(mBase, uint32(v4550)+60))
	if v4551 == int32(0) {
		v4588 = v4545
		goto L1343
	} else {
		goto L1345
	}
L1345:
	;
	v4554 = v4544 + v4551
	v4555 = *(*int32)(unsafe.Add(mBase, uint32(v4550)+188))
	if base.Ui32(v4555) <= base.Ui32(v4554) {
		goto L1347
	} else {
		goto L1348
	}
L1346:
	;
	v4565 = *(*int32)(unsafe.Add(mBase, uint32(v4550)+196))
	if v4564 == int32(0) {
		v4588 = v4565
		goto L1343
	} else {
		goto L1350
	}
L1347:
	;
	v4557 = *(*int32)(unsafe.Add(mBase, uint32(v4550)+192))
	v4564 = v4557
	goto L1346
L1348:
	;
	goto L1349
L1349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4550)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4550)+188)) = v4551
	v4562 = F_strchr(m, v4551, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4550)+192)) = v4562
	v4564 = v4562
	goto L1346
L1350:
	;
	if base.Ui32(v4554) <= base.Ui32(v4564) {
		v4588 = v4565
		goto L1343
	} else {
		goto L1351
	}
L1351:
	;
	v4569 = v4564
	v4571 = v4565
	goto L1352
L1352:
	;
	v4574 = int32(1)
	v4575 = v4571 + v4574
	*(*int32)(unsafe.Add(mBase, uint32(v4550)+196)) = v4575
	v4578 = v4569 + v4574
	*(*int32)(unsafe.Add(mBase, uint32(v4550)+188)) = v4578
	v4581 = F_strchr(m, v4578, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4550)+192)) = v4581
	if v4581 == int32(0) {
		v4588 = v4575
		goto L1343
	} else {
		goto L1354
	}
L1353:
	;
	v4588 = v4575
	goto L1343
L1354:
	;
	if base.Ui32(v4581) < base.Ui32(v4554) {
		v4569 = v4581
		v4571 = v4575
		goto L1352
	} else {
		goto L1355
	}
L1355:
	;
	goto L1353
L1356:
	;
	v4612 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v4613 = *(*int32)(unsafe.Add(mBase, uint32(v4612)))
	if v4613 != 0 {
		goto L1358
	} else {
		goto L1359
	}
L1357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4538
	v10149 = v241
	goto L5
L1358:
	;
	v4614 = v4612
	goto L1361
L1359:
	;
	v4617 = v4612
	goto L1360
L1360:
	;
	v4619 = *(*int32)(unsafe.Add(mBase, uint32(v4617)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1350])) = v4619
	goto L1357
L1361:
	;
	v4615 = *(*int32)(unsafe.Add(mBase, uint32(v4614)+8))
	v4616 = *(*int32)(unsafe.Add(mBase, uint32(v4615)))
	if v4616 != 0 {
		v4614 = v4615
		goto L1361
	} else {
		goto L1363
	}
L1362:
	;
	v4617 = v4615
	goto L1360
L1363:
	;
	goto L1362
L1364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4623))) = int32(5)
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4630 = int32(0)
	if v4629 < v4630 {
		v4673 = v4630
		goto L1366
	} else {
		goto L1367
	}
L1365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4623)+4)) = v4673
	v4678 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v4679 = *(*int32)(unsafe.Add(mBase, uint32(v4678)+520))
	v4681 = v4679 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4678)+520)) = v4681
	*(*int32)(unsafe.Add(mBase, uint32(v4623)+8)) = v4681
	v4685 = v132 - int32(48)
	v4686 = *(*int32)(unsafe.Add(mBase, uint32(v4685)))
	*(*int32)(unsafe.Add(mBase, uint32(v4623)+12)) = v4686
	v4690 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4623)+16)) = v4690
	v4692 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4623)+20)) = v4692
	v4694 = *(*int32)(unsafe.Add(mBase, uint32(v4685)))
	v4695 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v4696 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v4694, v4695, v4696, l1)
	mBase = m.M
	v4698 = m.ExcPending
	if v4698 != 0 {
		goto L66
	} else {
		goto L1379
	}
L1366:
	;
	goto L1365
L1367:
	;
	v4635 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4636 = *(*int32)(unsafe.Add(mBase, uint32(v4635)+60))
	if v4636 == int32(0) {
		v4673 = v4630
		goto L1366
	} else {
		goto L1368
	}
L1368:
	;
	v4639 = v4629 + v4636
	v4640 = *(*int32)(unsafe.Add(mBase, uint32(v4635)+188))
	if base.Ui32(v4640) <= base.Ui32(v4639) {
		goto L1370
	} else {
		goto L1371
	}
L1369:
	;
	v4650 = *(*int32)(unsafe.Add(mBase, uint32(v4635)+196))
	if v4649 == int32(0) {
		v4673 = v4650
		goto L1366
	} else {
		goto L1373
	}
L1370:
	;
	v4642 = *(*int32)(unsafe.Add(mBase, uint32(v4635)+192))
	v4649 = v4642
	goto L1369
L1371:
	;
	goto L1372
L1372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4635)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4635)+188)) = v4636
	v4647 = F_strchr(m, v4636, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4635)+192)) = v4647
	v4649 = v4647
	goto L1369
L1373:
	;
	if base.Ui32(v4639) <= base.Ui32(v4649) {
		v4673 = v4650
		goto L1366
	} else {
		goto L1374
	}
L1374:
	;
	v4654 = v4649
	v4656 = v4650
	goto L1375
L1375:
	;
	v4659 = int32(1)
	v4660 = v4656 + v4659
	*(*int32)(unsafe.Add(mBase, uint32(v4635)+196)) = v4660
	v4663 = v4654 + v4659
	*(*int32)(unsafe.Add(mBase, uint32(v4635)+188)) = v4663
	v4666 = F_strchr(m, v4663, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4635)+192)) = v4666
	if v4666 == int32(0) {
		v4673 = v4660
		goto L1366
	} else {
		goto L1377
	}
L1376:
	;
	v4673 = v4660
	goto L1366
L1377:
	;
	if base.Ui32(v4666) < base.Ui32(v4639) {
		v4654 = v4666
		v4656 = v4660
		goto L1375
	} else {
		goto L1378
	}
L1378:
	;
	goto L1376
L1379:
	;
	v4701 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v4702 = *(*int32)(unsafe.Add(mBase, uint32(v4701)))
	if v4702 != 0 {
		goto L1381
	} else {
		goto L1382
	}
L1380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4623
	v10149 = v241
	goto L5
L1381:
	;
	v4703 = v4701
	goto L1384
L1382:
	;
	v4706 = v4701
	goto L1383
L1383:
	;
	v4708 = *(*int32)(unsafe.Add(mBase, uint32(v4706)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1350])) = v4708
	goto L1380
L1384:
	;
	v4704 = *(*int32)(unsafe.Add(mBase, uint32(v4703)+8))
	v4705 = *(*int32)(unsafe.Add(mBase, uint32(v4704)))
	if v4705 != 0 {
		v4703 = v4704
		goto L1384
	} else {
		goto L1386
	}
L1385:
	;
	v4706 = v4704
	goto L1383
L1386:
	;
	goto L1385
L1387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4713)+4)) = v4761
	v4766 = v132 - int32(48)
	v4767 = *(*int32)(unsafe.Add(mBase, uint32(v4766)))
	*(*int32)(unsafe.Add(mBase, uint32(v4713)+12)) = v4767
	if v4714 == int32(6) {
		goto L1401
	} else {
		goto L1402
	}
L1388:
	;
	goto L1387
L1389:
	;
	v4723 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4724 = *(*int32)(unsafe.Add(mBase, uint32(v4723)+60))
	if v4724 == int32(0) {
		v4761 = v4718
		goto L1388
	} else {
		goto L1390
	}
L1390:
	;
	v4727 = v4717 + v4724
	v4728 = *(*int32)(unsafe.Add(mBase, uint32(v4723)+188))
	if base.Ui32(v4728) <= base.Ui32(v4727) {
		goto L1392
	} else {
		goto L1393
	}
L1391:
	;
	v4738 = *(*int32)(unsafe.Add(mBase, uint32(v4723)+196))
	if v4737 == int32(0) {
		v4761 = v4738
		goto L1388
	} else {
		goto L1395
	}
L1392:
	;
	v4730 = *(*int32)(unsafe.Add(mBase, uint32(v4723)+192))
	v4737 = v4730
	goto L1391
L1393:
	;
	goto L1394
L1394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4723)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4723)+188)) = v4724
	v4735 = F_strchr(m, v4724, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4723)+192)) = v4735
	v4737 = v4735
	goto L1391
L1395:
	;
	if base.Ui32(v4727) <= base.Ui32(v4737) {
		v4761 = v4738
		goto L1388
	} else {
		goto L1396
	}
L1396:
	;
	v4742 = v4737
	v4744 = v4738
	goto L1397
L1397:
	;
	v4747 = int32(1)
	v4748 = v4744 + v4747
	*(*int32)(unsafe.Add(mBase, uint32(v4723)+196)) = v4748
	v4751 = v4742 + v4747
	*(*int32)(unsafe.Add(mBase, uint32(v4723)+188)) = v4751
	v4754 = F_strchr(m, v4751, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4723)+192)) = v4754
	if v4754 == int32(0) {
		v4761 = v4748
		goto L1388
	} else {
		goto L1399
	}
L1398:
	;
	v4761 = v4748
	goto L1388
L1399:
	;
	if base.Ui32(v4754) < base.Ui32(v4727) {
		v4742 = v4754
		v4744 = v4748
		goto L1397
	} else {
		goto L1400
	}
L1400:
	;
	goto L1398
L1401:
	;
	v4773 = int32(36)
	goto L1403
L1402:
	;
	v4773 = int32(20)
	goto L1403
L1403:
	;
	v4775 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4713+v4773))) = v4775
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4713
	v4778 = *(*int32)(unsafe.Add(mBase, uint32(v4766)))
	v4779 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v4780 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v4778, v4779, v4780, l1)
	mBase = m.M
	v4782 = m.ExcPending
	if v4782 != 0 {
		goto L66
	} else {
		goto L1404
	}
L1404:
	;
	v4785 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v4786 = *(*int32)(unsafe.Add(mBase, uint32(v4785)))
	if v4786 != 0 {
		goto L1406
	} else {
		goto L1407
	}
L1405:
	;
	v10149 = v241
	goto L5
L1406:
	;
	v4787 = v4785
	goto L1409
L1407:
	;
	v4790 = v4785
	goto L1408
L1408:
	;
	v4792 = *(*int32)(unsafe.Add(mBase, uint32(v4790)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1350])) = v4792
	goto L1405
L1409:
	;
	v4788 = *(*int32)(unsafe.Add(mBase, uint32(v4787)+8))
	v4789 = *(*int32)(unsafe.Add(mBase, uint32(v4788)))
	if v4789 != 0 {
		v4787 = v4788
		goto L1409
	} else {
		goto L1411
	}
L1410:
	;
	v4790 = v4788
	goto L1408
L1411:
	;
	goto L1410
L1412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377]))) = v4798
	v4801 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	if v4798 != int32(277) {
		goto L1415
	} else {
		goto L1416
	}
L1413:
	;
	F_plpgsql_push_back_token(m, v4798, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5053 = m.ExcPending
	if v5053 != 0 {
		goto L66
	} else {
		goto L1463
	}
L1414:
	;
	if v4798 == int32(363) {
		goto L10
	} else {
		goto L1462
	}
L1415:
	;
	if v4798 != int32(317) {
		goto L1414
	} else {
		goto L1418
	}
L1416:
	;
	goto L1417
L1417:
	;
	v4965 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v4966 = *(*int32)(unsafe.Add(mBase, uint32(v4965)))
	if v4966 != 0 {
		goto L1440
	} else {
		goto L1441
	}
L1418:
	;
	v4808 = int32(0)
	v4811 = int32(1)
	v4820 = F_read_sql_construct(m, int32(336), int32(381), v4808, int32(536074), int32(2), v4811, v4811, v4808, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4821 = m.ExcPending
	if v4821 != 0 {
		goto L66
	} else {
		goto L1419
	}
L1419:
	;
	v4823 = F_palloc0(m, int32(32))
	mBase = m.M
	v4824 = m.ExcPending
	if v4824 != 0 {
		goto L66
	} else {
		goto L1420
	}
L1420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4823))) = int32(18)
	v4828 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v4829 = *(*int32)(unsafe.Add(mBase, uint32(v4828)+520))
	v4831 = v4829 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4828)+520)) = v4831
	*(*int32)(unsafe.Add(mBase, uint32(v4823)+8)) = v4831
	v4835 = v132 - int32(4)
	v4836 = *(*int32)(unsafe.Add(mBase, uint32(v4835)))
	if v4836 != 0 {
		goto L1422
	} else {
		goto L1423
	}
L1421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4823)+24)) = v4820
	v4890 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v4890 == int32(381) {
		goto L1432
	} else {
		goto L1433
	}
L1422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4823)+16)) = v4836
	v4838 = *(*int32)(unsafe.Add(mBase, uint32(v4835)))
	v4841 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v4838, v4841, l1)
	mBase = m.M
	v4843 = m.ExcPending
	if v4843 != 0 {
		goto L66
	} else {
		goto L1425
	}
L1423:
	;
	goto L1424
L1424:
	;
	v4846 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(8))))
	if v4846 == int32(0) {
		goto L43
	} else {
		goto L1426
	}
L1425:
	;
	goto L1421
L1426:
	;
	v4851 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v4854 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v4857 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v4846, v4857, l1)
	mBase = m.M
	v4859 = m.ExcPending
	if v4859 != 0 {
		goto L66
	} else {
		goto L1427
	}
L1427:
	;
	v4861 = F_palloc0(m, int32(40))
	mBase = m.M
	v4862 = m.ExcPending
	if v4862 != 0 {
		goto L66
	} else {
		goto L1428
	}
L1428:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4861)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v4861)+12)) = v4851
	*(*int32)(unsafe.Add(mBase, uint32(v4861)+8)) = int32(670471)
	*(*int32)(unsafe.Add(mBase, uint32(v4861))) = int32(1)
	v4871 = F_palloc(m, int32(4))
	mBase = m.M
	v4872 = m.ExcPending
	if v4872 != 0 {
		goto L66
	} else {
		goto L1429
	}
L1429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4861)+32)) = v4871
	v4875 = F_palloc(m, int32(4))
	mBase = m.M
	v4876 = m.ExcPending
	if v4876 != 0 {
		goto L66
	} else {
		goto L1430
	}
L1430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4861)+36)) = v4875
	v4878 = *(*int32)(unsafe.Add(mBase, uint32(v4861)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4878))) = v4854
	v4880 = *(*int32)(unsafe.Add(mBase, uint32(v4846)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4875))) = v4880
	F_plpgsql_adddatum(m, v4861)
	mBase = m.M
	v4883 = m.ExcPending
	if v4883 != 0 {
		goto L66
	} else {
		goto L1431
	}
L1431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4823)+16)) = v4861
	goto L1421
L1432:
	;
	goto L1435
L1433:
	;
	goto L1434
L1434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4823
	v10149 = v241
	goto L5
L1435:
	;
	v4919 = int32(0)
	v4922 = int32(1)
	v4931 = F_read_sql_construct(m, int32(44), int32(336), v4919, int32(526583), int32(2), v4922, v4922, v4919, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4932 = m.ExcPending
	if v4932 != 0 {
		goto L66
	} else {
		goto L1437
	}
L1436:
	;
	goto L1434
L1437:
	;
	v4933 = *(*int32)(unsafe.Add(mBase, uint32(v4823)+28))
	v4934 = F_lappend(m, v4933, v4931)
	mBase = m.M
	v4935 = m.ExcPending
	if v4935 != 0 {
		goto L66
	} else {
		goto L1438
	}
L1438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4823)+28)) = v4934
	v4937 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v4937 == int32(44) {
		goto L1435
	} else {
		goto L1439
	}
L1439:
	;
	goto L1436
L1440:
	;
	v5015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v5015 != 0 {
		goto L1413
	} else {
		goto L1451
	}
L1441:
	;
	v4967 = *(*int32)(unsafe.Add(mBase, uint32(v4965)+24))
	v4968 = *(*int32)(unsafe.Add(mBase, uint32(v4967)+4))
	if v4968 != int32(1790) {
		goto L1440
	} else {
		goto L1442
	}
L1442:
	;
	v4972 = F_palloc0(m, int32(32))
	mBase = m.M
	v4973 = m.ExcPending
	if v4973 != 0 {
		goto L66
	} else {
		goto L1443
	}
L1443:
	;
	v4974 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4972))) = v4974
	v4977 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v4978 = *(*int32)(unsafe.Add(mBase, uint32(v4977)+520))
	v4980 = v4978 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4977)+520)) = v4980
	*(*int32)(unsafe.Add(mBase, uint32(v4972)+8)) = v4980
	v4983 = *(*int32)(unsafe.Add(mBase, uint32(v4965)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4972)+24)) = v4983
	v4987 = *(*int32)(unsafe.Add(mBase, uint32(v132-v4974)))
	if v4987 != 0 {
		goto L1444
	} else {
		goto L1445
	}
L1444:
	;
	v4990 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
	if v4990 != 0 {
		goto L42
	} else {
		goto L1447
	}
L1445:
	;
	goto L1446
L1446:
	;
	v4991 = *(*int32)(unsafe.Add(mBase, uint32(v4965)+28))
	if v4991 == int32(0) {
		goto L41
	} else {
		goto L1448
	}
L1447:
	;
	goto L1446
L1448:
	;
	v4999 = F_read_cursor_args(m, v4965, int32(336), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5000 = m.ExcPending
	if v5000 != 0 {
		goto L66
	} else {
		goto L1449
	}
L1449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4972)+28)) = v4999
	v5004 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v5007 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v5011 = F_plpgsql_build_record(m, v5004, v5007, int32(0), int32(2249), int32(1))
	mBase = m.M
	v5012 = m.ExcPending
	if v5012 != 0 {
		goto L66
	} else {
		goto L1450
	}
L1450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4972)+16)) = v5011
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4972
	v10149 = v241
	goto L5
L1451:
	;
	v5016 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v5016 == int32(0) {
		goto L1413
	} else {
		goto L1452
	}
L1452:
	;
	v5019 = int32(360583)
	v5022 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1396])))
	v5023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5016))))
	if v5023 == int32(0) {
		v5042 = v5022
		v5043 = v5023
		goto L1454
	} else {
		goto L1455
	}
L1453:
	;
	if v5043-v5042 != 0 {
		goto L1413
	} else {
		goto L1461
	}
L1454:
	;
	goto L1453
L1455:
	;
	if v5022 != v5023 {
		v5042 = v5022
		v5043 = v5023
		goto L1454
	} else {
		goto L1456
	}
L1456:
	;
	v5027 = v5016
	v5028 = v5019
	goto L1457
L1457:
	;
	v5031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5028)+1)))
	v5032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5027)+1)))
	if v5032 == int32(0) {
		v5042 = v5031
		v5043 = v5032
		goto L1454
	} else {
		goto L1459
	}
L1458:
	;
	v5042 = v5031
	v5043 = v5032
	goto L1454
L1459:
	;
	v5035 = int32(1)
	if v5031 == v5032 {
		v5027 = v5027 + v5035
		v5028 = v5028 + v5035
		goto L1457
	} else {
		goto L1460
	}
L1460:
	;
	goto L1458
L1461:
	;
	goto L10
L1462:
	;
	goto L1413
L1463:
	;
	v5054 = int32(0)
	v5070 = F_read_sql_construct(m, int32(269), int32(336), v5054, int32(526588), v5054, int32(1), v5054, v27+int32(4764), v27+int32(4752), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5071 = m.ExcPending
	if v5071 != 0 {
		goto L66
	} else {
		goto L1464
	}
L1464:
	;
	v5072 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377])))
	if v5072 == int32(269) {
		v9837 = v5070
		v9838 = v5054
		goto L9
	} else {
		goto L1465
	}
L1465:
	;
	v5076 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1356])))
	if v5076 == int32(1) {
		goto L1466
	} else {
		goto L1467
	}
L1466:
	;
	v5079 = *(*int32)(unsafe.Add(mBase, uint32(v5070)+4))
	v5080 = *(*int32)(unsafe.Add(mBase, uint32(v5070)))
	v5081 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1397])))
	v5082 = int32(4515248)
	v5083 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v5086 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v5086
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v5081
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1380]))) = int32(6810)
	v5092 = int32(4508152)
	v5093 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4784)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = v5093
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1381]))) = v27 + int32(4768)
	v5102 = F_raw_parser(m, v5080, v5079)
	mBase = m.M
	v5103 = m.ExcPending
	if v5103 != 0 {
		goto L66
	} else {
		goto L1469
	}
L1467:
	;
	goto L1468
L1468:
	;
	v5115 = F_palloc0(m, int32(28))
	mBase = m.M
	v5116 = m.ExcPending
	if v5116 != 0 {
		goto L66
	} else {
		goto L1470
	}
L1469:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v5083
	v5107 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v5107
	goto L1468
L1470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5115))) = int32(7)
	v5120 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5121 = *(*int32)(unsafe.Add(mBase, uint32(v5120)+520))
	v5123 = v5121 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5120)+520)) = v5123
	*(*int32)(unsafe.Add(mBase, uint32(v5115)+8)) = v5123
	v5127 = v132 - int32(4)
	v5128 = *(*int32)(unsafe.Add(mBase, uint32(v5127)))
	if v5128 != 0 {
		goto L1472
	} else {
		goto L1473
	}
L1471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5115)+24)) = v5070
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5115
	v10149 = v241
	goto L5
L1472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5115)+16)) = v5128
	v5130 = *(*int32)(unsafe.Add(mBase, uint32(v5127)))
	v5133 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v5130, v5133, l1)
	mBase = m.M
	v5135 = m.ExcPending
	if v5135 != 0 {
		goto L66
	} else {
		goto L1475
	}
L1473:
	;
	goto L1474
L1474:
	;
	v5138 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(8))))
	if v5138 == int32(0) {
		goto L40
	} else {
		goto L1476
	}
L1475:
	;
	goto L1471
L1476:
	;
	v5143 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v5146 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v5149 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v5138, v5149, l1)
	mBase = m.M
	v5151 = m.ExcPending
	if v5151 != 0 {
		goto L66
	} else {
		goto L1477
	}
L1477:
	;
	v5153 = F_palloc0(m, int32(40))
	mBase = m.M
	v5154 = m.ExcPending
	if v5154 != 0 {
		goto L66
	} else {
		goto L1478
	}
L1478:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5153)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5153)+12)) = v5143
	*(*int32)(unsafe.Add(mBase, uint32(v5153)+8)) = int32(670471)
	*(*int32)(unsafe.Add(mBase, uint32(v5153))) = int32(1)
	v5163 = F_palloc(m, int32(4))
	mBase = m.M
	v5164 = m.ExcPending
	if v5164 != 0 {
		goto L66
	} else {
		goto L1479
	}
L1479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5153)+32)) = v5163
	v5167 = F_palloc(m, int32(4))
	mBase = m.M
	v5168 = m.ExcPending
	if v5168 != 0 {
		goto L66
	} else {
		goto L1480
	}
L1480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5153)+36)) = v5167
	v5170 = *(*int32)(unsafe.Add(mBase, uint32(v5153)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5170))) = v5146
	v5172 = *(*int32)(unsafe.Add(mBase, uint32(v5138)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5167))) = v5172
	F_plpgsql_adddatum(m, v5153)
	mBase = m.M
	v5175 = m.ExcPending
	if v5175 != 0 {
		goto L66
	} else {
		goto L1481
	}
L1481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5115)+16)) = v5153
	goto L1471
L1482:
	;
	v5187 = v5183
	goto L1484
L1483:
	;
	v5184 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v5185 = F_NameListToString(m, v5184)
	mBase = m.M
	v5186 = m.ExcPending
	if v5186 != 0 {
		goto L66
	} else {
		goto L1485
	}
L1484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5187
	v5189 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5190 = int32(0)
	if v5189 < v5190 {
		v5233 = v5190
		goto L1487
	} else {
		goto L1488
	}
L1485:
	;
	v5187 = v5185
	goto L1484
L1486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v5233
	v5237 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v5238 = *(*int32)(unsafe.Add(mBase, uint32(v5237)))
	v5239 = int32(1)
	if base.Ui32(v5238-v5239) <= base.Ui32(v5239) {
		goto L1500
	} else {
		goto L1501
	}
L1487:
	;
	goto L1486
L1488:
	;
	v5195 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5196 = *(*int32)(unsafe.Add(mBase, uint32(v5195)+60))
	if v5196 == int32(0) {
		v5233 = v5190
		goto L1487
	} else {
		goto L1489
	}
L1489:
	;
	v5199 = v5189 + v5196
	v5200 = *(*int32)(unsafe.Add(mBase, uint32(v5195)+188))
	if base.Ui32(v5200) <= base.Ui32(v5199) {
		goto L1491
	} else {
		goto L1492
	}
L1490:
	;
	v5210 = *(*int32)(unsafe.Add(mBase, uint32(v5195)+196))
	if v5209 == int32(0) {
		v5233 = v5210
		goto L1487
	} else {
		goto L1494
	}
L1491:
	;
	v5202 = *(*int32)(unsafe.Add(mBase, uint32(v5195)+192))
	v5209 = v5202
	goto L1490
L1492:
	;
	goto L1493
L1493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5195)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5195)+188)) = v5196
	v5207 = F_strchr(m, v5196, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5195)+192)) = v5207
	v5209 = v5207
	goto L1490
L1494:
	;
	if base.Ui32(v5199) <= base.Ui32(v5209) {
		v5233 = v5210
		goto L1487
	} else {
		goto L1495
	}
L1495:
	;
	v5214 = v5209
	v5216 = v5210
	goto L1496
L1496:
	;
	v5219 = int32(1)
	v5220 = v5216 + v5219
	*(*int32)(unsafe.Add(mBase, uint32(v5195)+196)) = v5220
	v5223 = v5214 + v5219
	*(*int32)(unsafe.Add(mBase, uint32(v5195)+188)) = v5223
	v5226 = F_strchr(m, v5223, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5195)+192)) = v5226
	if v5226 == int32(0) {
		v5233 = v5220
		goto L1487
	} else {
		goto L1498
	}
L1497:
	;
	v5233 = v5220
	goto L1487
L1498:
	;
	if base.Ui32(v5226) < base.Ui32(v5199) {
		v5214 = v5226
		v5216 = v5220
		goto L1496
	} else {
		goto L1499
	}
L1499:
	;
	goto L1497
L1500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = int32(0)
	v5245 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v5245
	v10149 = v241
	goto L5
L1501:
	;
	goto L1502
L1502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v5237
	v5254 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5255 = m.ExcPending
	if v5255 != 0 {
		goto L66
	} else {
		goto L1503
	}
L1503:
	;
	F_plpgsql_push_back_token(m, v5254, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5261 = m.ExcPending
	if v5261 != 0 {
		goto L66
	} else {
		goto L1504
	}
L1504:
	;
	if v5254 != int32(44) {
		v10149 = v241
		goto L5
	} else {
		goto L1505
	}
L1505:
	;
	v5264 = *(*int32)(unsafe.Add(mBase, uint32(v27)+304))
	v5265 = *(*int32)(unsafe.Add(mBase, uint32(v27)+312))
	v5266 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5271 = F_read_into_scalar_list(m, v5264, v5265, v5266, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5272 = m.ExcPending
	if v5272 != 0 {
		goto L66
	} else {
		goto L1506
	}
L1506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v5271
	v10149 = v241
	goto L5
L1507:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+312)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v5320
	v5330 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5331 = m.ExcPending
	if v5331 != 0 {
		goto L66
	} else {
		goto L1521
	}
L1508:
	;
	goto L1507
L1509:
	;
	v5282 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5283 = *(*int32)(unsafe.Add(mBase, uint32(v5282)+60))
	if v5283 == int32(0) {
		v5320 = v5277
		goto L1508
	} else {
		goto L1510
	}
L1510:
	;
	v5286 = v5276 + v5283
	v5287 = *(*int32)(unsafe.Add(mBase, uint32(v5282)+188))
	if base.Ui32(v5287) <= base.Ui32(v5286) {
		goto L1512
	} else {
		goto L1513
	}
L1511:
	;
	v5297 = *(*int32)(unsafe.Add(mBase, uint32(v5282)+196))
	if v5296 == int32(0) {
		v5320 = v5297
		goto L1508
	} else {
		goto L1515
	}
L1512:
	;
	v5289 = *(*int32)(unsafe.Add(mBase, uint32(v5282)+192))
	v5296 = v5289
	goto L1511
L1513:
	;
	goto L1514
L1514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+188)) = v5283
	v5294 = F_strchr(m, v5283, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+192)) = v5294
	v5296 = v5294
	goto L1511
L1515:
	;
	if base.Ui32(v5286) <= base.Ui32(v5296) {
		v5320 = v5297
		goto L1508
	} else {
		goto L1516
	}
L1516:
	;
	v5301 = v5296
	v5303 = v5297
	goto L1517
L1517:
	;
	v5306 = int32(1)
	v5307 = v5303 + v5306
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+196)) = v5307
	v5310 = v5301 + v5306
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+188)) = v5310
	v5313 = F_strchr(m, v5310, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5282)+192)) = v5313
	if v5313 == int32(0) {
		v5320 = v5307
		goto L1508
	} else {
		goto L1519
	}
L1518:
	;
	v5320 = v5307
	goto L1508
L1519:
	;
	if base.Ui32(v5313) < base.Ui32(v5286) {
		v5301 = v5313
		v5303 = v5307
		goto L1517
	} else {
		goto L1520
	}
L1520:
	;
	goto L1518
L1521:
	;
	F_plpgsql_push_back_token(m, v5330, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5337 = m.ExcPending
	if v5337 != 0 {
		goto L66
	} else {
		goto L1522
	}
L1522:
	;
	if v5330 != int32(44) {
		v10149 = v241
		goto L5
	} else {
		goto L1523
	}
L1523:
	;
	v5340 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v5340, l1)
	mBase = m.M
	v5342 = m.ExcPending
	if v5342 != 0 {
		goto L66
	} else {
		goto L1524
	}
L1524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1525:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5347))) = int32(9)
	v5353 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(24))))
	v5354 = int32(0)
	if v5353 < v5354 {
		v5397 = v5354
		goto L1528
	} else {
		goto L1529
	}
L1527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5347)+4)) = v5397
	v5402 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5403 = *(*int32)(unsafe.Add(mBase, uint32(v5402)+520))
	v5405 = v5403 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5402)+520)) = v5405
	*(*int32)(unsafe.Add(mBase, uint32(v5347)+8)) = v5405
	v5409 = v132 - int32(112)
	v5410 = *(*int32)(unsafe.Add(mBase, uint32(v5409)))
	*(*int32)(unsafe.Add(mBase, uint32(v5347)+12)) = v5410
	v5414 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v5347)+20)) = v5414
	v5418 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v5347)+24)) = v5418
	v5420 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v5347)+28)) = v5420
	v5423 = v132 - int32(68)
	v5424 = *(*int32)(unsafe.Add(mBase, uint32(v5423)))
	if v5424 == int32(0) {
		goto L1541
	} else {
		goto L1542
	}
L1528:
	;
	goto L1527
L1529:
	;
	v5359 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5360 = *(*int32)(unsafe.Add(mBase, uint32(v5359)+60))
	if v5360 == int32(0) {
		v5397 = v5354
		goto L1528
	} else {
		goto L1530
	}
L1530:
	;
	v5363 = v5353 + v5360
	v5364 = *(*int32)(unsafe.Add(mBase, uint32(v5359)+188))
	if base.Ui32(v5364) <= base.Ui32(v5363) {
		goto L1532
	} else {
		goto L1533
	}
L1531:
	;
	v5374 = *(*int32)(unsafe.Add(mBase, uint32(v5359)+196))
	if v5373 == int32(0) {
		v5397 = v5374
		goto L1528
	} else {
		goto L1535
	}
L1532:
	;
	v5366 = *(*int32)(unsafe.Add(mBase, uint32(v5359)+192))
	v5373 = v5366
	goto L1531
L1533:
	;
	goto L1534
L1534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5359)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5359)+188)) = v5360
	v5371 = F_strchr(m, v5360, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5359)+192)) = v5371
	v5373 = v5371
	goto L1531
L1535:
	;
	if base.Ui32(v5363) <= base.Ui32(v5373) {
		v5397 = v5374
		goto L1528
	} else {
		goto L1536
	}
L1536:
	;
	v5378 = v5373
	v5380 = v5374
	goto L1537
L1537:
	;
	v5383 = int32(1)
	v5384 = v5380 + v5383
	*(*int32)(unsafe.Add(mBase, uint32(v5359)+196)) = v5384
	v5387 = v5378 + v5383
	*(*int32)(unsafe.Add(mBase, uint32(v5359)+188)) = v5387
	v5390 = F_strchr(m, v5387, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5359)+192)) = v5390
	if v5390 == int32(0) {
		v5397 = v5384
		goto L1528
	} else {
		goto L1539
	}
L1538:
	;
	v5397 = v5384
	goto L1528
L1539:
	;
	if base.Ui32(v5390) < base.Ui32(v5363) {
		v5378 = v5390
		v5380 = v5384
		goto L1537
	} else {
		goto L1540
	}
L1540:
	;
	goto L1538
L1541:
	;
	v5428 = v132 - int32(72)
	v5429 = *(*int32)(unsafe.Add(mBase, uint32(v5428)))
	if v5429 == int32(0) {
		goto L39
	} else {
		goto L1544
	}
L1542:
	;
	v5432 = v5423
	v5433 = v5424
	goto L1543
L1543:
	;
	v5434 = *(*int32)(unsafe.Add(mBase, uint32(v5433)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5347)+16)) = v5434
	v5436 = *(*int32)(unsafe.Add(mBase, uint32(v5432)))
	v5439 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(20))))
	F_check_assignable(m, v5436, v5439, l1)
	mBase = m.M
	v5441 = m.ExcPending
	if v5441 != 0 {
		goto L66
	} else {
		goto L1545
	}
L1544:
	;
	v5432 = v5428
	v5433 = v5429
	goto L1543
L1545:
	;
	v5442 = *(*int32)(unsafe.Add(mBase, uint32(v5409)))
	v5443 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v5444 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v5442, v5443, v5444, l1)
	mBase = m.M
	v5446 = m.ExcPending
	if v5446 != 0 {
		goto L66
	} else {
		goto L1546
	}
L1546:
	;
	v5449 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v5450 = *(*int32)(unsafe.Add(mBase, uint32(v5449)))
	if v5450 != 0 {
		goto L1548
	} else {
		goto L1549
	}
L1547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5347
	v10149 = v241
	goto L5
L1548:
	;
	v5451 = v5449
	goto L1551
L1549:
	;
	v5454 = v5449
	goto L1550
L1550:
	;
	v5456 = *(*int32)(unsafe.Add(mBase, uint32(v5454)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1350])) = v5456
	goto L1547
L1551:
	;
	v5452 = *(*int32)(unsafe.Add(mBase, uint32(v5451)+8))
	v5453 = *(*int32)(unsafe.Add(mBase, uint32(v5452)))
	if v5453 != 0 {
		v5451 = v5452
		goto L1551
	} else {
		goto L1553
	}
L1552:
	;
	v5454 = v5452
	goto L1550
L1553:
	;
	goto L1552
L1554:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5464))) = int32(10)
	v5469 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5470 = *(*int32)(unsafe.Add(mBase, uint32(v5469)+520))
	v5472 = v5470 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5469)+520)) = v5472
	*(*int32)(unsafe.Add(mBase, uint32(v5464)+8)) = v5472
	v5477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132-int32(32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5464)+12)) = uint8(v5477)
	v5480 = v142 - int32(8)
	v5481 = *(*int32)(unsafe.Add(mBase, uint32(v5480)))
	v5482 = int32(0)
	if v5481 < v5482 {
		v5525 = v5482
		goto L1556
	} else {
		goto L1557
	}
L1555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5464)+4)) = v5525
	v5530 = v132 - int32(16)
	v5531 = *(*int32)(unsafe.Add(mBase, uint32(v5530)))
	*(*int32)(unsafe.Add(mBase, uint32(v5464)+16)) = v5531
	v5533 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v5464)+20)) = v5533
	v5536 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v5537 = *(*int32)(unsafe.Add(mBase, uint32(v5530)))
	if v5537 != 0 {
		goto L1570
	} else {
		goto L1571
	}
L1556:
	;
	goto L1555
L1557:
	;
	v5487 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5488 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+60))
	if v5488 == int32(0) {
		v5525 = v5482
		goto L1556
	} else {
		goto L1558
	}
L1558:
	;
	v5491 = v5481 + v5488
	v5492 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+188))
	if base.Ui32(v5492) <= base.Ui32(v5491) {
		goto L1560
	} else {
		goto L1561
	}
L1559:
	;
	v5502 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+196))
	if v5501 == int32(0) {
		v5525 = v5502
		goto L1556
	} else {
		goto L1563
	}
L1560:
	;
	v5494 = *(*int32)(unsafe.Add(mBase, uint32(v5487)+192))
	v5501 = v5494
	goto L1559
L1561:
	;
	goto L1562
L1562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5487)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5487)+188)) = v5488
	v5499 = F_strchr(m, v5488, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5487)+192)) = v5499
	v5501 = v5499
	goto L1559
L1563:
	;
	if base.Ui32(v5491) <= base.Ui32(v5501) {
		v5525 = v5502
		goto L1556
	} else {
		goto L1564
	}
L1564:
	;
	v5506 = v5501
	v5508 = v5502
	goto L1565
L1565:
	;
	v5511 = int32(1)
	v5512 = v5508 + v5511
	*(*int32)(unsafe.Add(mBase, uint32(v5487)+196)) = v5512
	v5515 = v5506 + v5511
	*(*int32)(unsafe.Add(mBase, uint32(v5487)+188)) = v5515
	v5518 = F_strchr(m, v5515, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5487)+192)) = v5518
	if v5518 == int32(0) {
		v5525 = v5512
		goto L1556
	} else {
		goto L1567
	}
L1566:
	;
	v5525 = v5512
	goto L1556
L1567:
	;
	if base.Ui32(v5518) < base.Ui32(v5491) {
		v5506 = v5518
		v5508 = v5512
		goto L1565
	} else {
		goto L1568
	}
L1568:
	;
	goto L1566
L1569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5464
	v10149 = v241
	goto L5
L1570:
	;
	v5538 = *(*int32)(unsafe.Add(mBase, uint32(v5530)))
	if v5536 != 0 {
		goto L1574
	} else {
		goto L1575
	}
L1571:
	;
	goto L1572
L1572:
	;
	if v5536 != 0 {
		goto L1592
	} else {
		goto L1593
	}
L1573:
	;
	if v5551 == int32(0) {
		goto L38
	} else {
		goto L1583
	}
L1574:
	;
	v5539 = v5536
	goto L1577
L1575:
	;
	goto L1576
L1576:
	;
	v5551 = int32(0)
	goto L1573
L1577:
	;
	v5541 = *(*int32)(unsafe.Add(mBase, uint32(v5539)))
	if v5541 != 0 {
		goto L1579
	} else {
		goto L1580
	}
L1578:
	;
	goto L1576
L1579:
	;
	v5545 = *(*int32)(unsafe.Add(mBase, uint32(v5539)+8))
	if v5545 != 0 {
		v5539 = v5545
		goto L1577
	} else {
		goto L1582
	}
L1580:
	;
	v5544 = F_strcmp(m, v5539+int32(12), v5538)
	mBase = m.M
	if v5544 != 0 {
		goto L1579
	} else {
		goto L1581
	}
L1581:
	;
	v5551 = v5539
	goto L1573
L1582:
	;
	goto L1578
L1583:
	;
	v5554 = *(*int32)(unsafe.Add(mBase, uint32(v5551)+4))
	if v5554 == int32(1) {
		goto L1569
	} else {
		goto L1584
	}
L1584:
	;
	v5557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5464)+12)))
	if v5557 != 0 {
		goto L1569
	} else {
		goto L1585
	}
L1585:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v5561 = m.ExcPending
	if v5561 != 0 {
		goto L66
	} else {
		goto L1586
	}
L1586:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v5564 = m.ExcPending
	if v5564 != 0 {
		goto L66
	} else {
		goto L1587
	}
L1587:
	;
	v5565 = *(*int32)(unsafe.Add(mBase, uint32(v5530)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+240)) = v5565
	F_errmsg(m, int32(537348), v27+int32(240))
	mBase = m.M
	v5571 = m.ExcPending
	if v5571 != 0 {
		goto L66
	} else {
		goto L1588
	}
L1588:
	;
	v5574 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v5575 = F_plpgsql_scanner_errposition(m, v5574, l1)
	mBase = m.M
	v5576 = m.ExcPending
	if v5576 != 0 {
		goto L66
	} else {
		goto L1589
	}
L1589:
	;
	F_errfinish(m, int32(26959), int32(1752), int32(360630))
	mBase = m.M
	v5581 = m.ExcPending
	if v5581 != 0 {
		goto L66
	} else {
		goto L1590
	}
L1590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1591:
	;
	if v5591 == int32(0) {
		goto L37
	} else {
		goto L1601
	}
L1592:
	;
	v5582 = v5536
	goto L1595
L1593:
	;
	goto L1594
L1594:
	;
	v5591 = int32(0)
	goto L1591
L1595:
	;
	v5583 = *(*int32)(unsafe.Add(mBase, uint32(v5582)))
	if v5583 != 0 {
		goto L1597
	} else {
		goto L1598
	}
L1596:
	;
	goto L1594
L1597:
	;
	v5587 = *(*int32)(unsafe.Add(mBase, uint32(v5582)+8))
	if v5587 != 0 {
		v5582 = v5587
		goto L1595
	} else {
		goto L1600
	}
L1598:
	;
	v5584 = *(*int32)(unsafe.Add(mBase, uint32(v5582)+4))
	if v5584 != int32(1) {
		goto L1597
	} else {
		goto L1599
	}
L1599:
	;
	v5591 = v5582
	goto L1591
L1600:
	;
	goto L1596
L1601:
	;
	goto L1569
L1602:
	;
	F_plpgsql_push_back_token(m, v5604, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5997 = m.ExcPending
	if v5997 != 0 {
		goto L66
	} else {
		goto L1694
	}
L1603:
	;
	v5803 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5805 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5805)+63)))
	if v5806 == int32(0) {
		goto L34
	} else {
		goto L1663
	}
L1604:
	;
	v5776 = int32(17126)
	v5779 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1398])))
	v5780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5617))))
	if v5780 == int32(0) {
		v5799 = v5779
		v5800 = v5780
		goto L1655
	} else {
		goto L1656
	}
L1605:
	;
	v5647 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5649 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5649)+63)))
	if v5650 == int32(0) {
		goto L36
	} else {
		goto L1624
	}
L1606:
	;
	if v5604 != int32(277) {
		goto L1608
	} else {
		goto L1609
	}
L1607:
	;
	switch v5604 - int32(341) {
	case 0:
		goto L1605
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		goto L1602
	case 17:
		goto L1603
	default:
		goto L1606
	}
L1608:
	;
	if v5604 != 0 {
		goto L1602
	} else {
		goto L1611
	}
L1609:
	;
	goto L1610
L1610:
	;
	v5616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v5616 != 0 {
		goto L1602
	} else {
		goto L1613
	}
L1611:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(250277))
	mBase = m.M
	v5615 = m.ExcPending
	if v5615 != 0 {
		goto L66
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
	v5617 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v5617 == int32(0) {
		goto L1602
	} else {
		goto L1614
	}
L1614:
	;
	v5620 = int32(63501)
	v5623 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1399])))
	v5624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5617))))
	if v5624 == int32(0) {
		v5643 = v5623
		v5644 = v5624
		goto L1616
	} else {
		goto L1617
	}
L1615:
	;
	if v5644-v5643 != 0 {
		goto L1604
	} else {
		goto L1623
	}
L1616:
	;
	goto L1615
L1617:
	;
	if v5623 != v5624 {
		v5643 = v5623
		v5644 = v5624
		goto L1616
	} else {
		goto L1618
	}
L1618:
	;
	v5628 = v5617
	v5629 = v5620
	goto L1619
L1619:
	;
	v5632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5629)+1)))
	v5633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5628)+1)))
	if v5633 == int32(0) {
		v5643 = v5632
		v5644 = v5633
		goto L1616
	} else {
		goto L1621
	}
L1620:
	;
	v5643 = v5632
	v5644 = v5633
	goto L1616
L1621:
	;
	v5636 = int32(1)
	if v5632 == v5633 {
		v5628 = v5628 + v5636
		v5629 = v5629 + v5636
		goto L1619
	} else {
		goto L1622
	}
L1622:
	;
	goto L1620
L1623:
	;
	goto L1605
L1624:
	;
	v5654 = F_palloc0(m, int32(20))
	mBase = m.M
	v5655 = m.ExcPending
	if v5655 != 0 {
		goto L66
	} else {
		goto L1625
	}
L1625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5654))) = int32(12)
	v5658 = int32(0)
	if v5647 < v5658 {
		v5701 = v5658
		goto L1627
	} else {
		goto L1628
	}
L1626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+4)) = v5701
	v5706 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5707 = *(*int32)(unsafe.Add(mBase, uint32(v5706)+520))
	v5709 = v5707 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5706)+520)) = v5709
	*(*int64)(unsafe.Add(mBase, uint32(v5654)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+8)) = v5709
	v5714 = *(*int32)(unsafe.Add(mBase, uint32(v5706)+472))
	v5719 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5720 = m.ExcPending
	if v5720 != 0 {
		goto L66
	} else {
		goto L1640
	}
L1627:
	;
	goto L1626
L1628:
	;
	v5663 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5664 = *(*int32)(unsafe.Add(mBase, uint32(v5663)+60))
	if v5664 == int32(0) {
		v5701 = v5658
		goto L1627
	} else {
		goto L1629
	}
L1629:
	;
	v5667 = v5647 + v5664
	v5668 = *(*int32)(unsafe.Add(mBase, uint32(v5663)+188))
	if base.Ui32(v5668) <= base.Ui32(v5667) {
		goto L1631
	} else {
		goto L1632
	}
L1630:
	;
	v5678 = *(*int32)(unsafe.Add(mBase, uint32(v5663)+196))
	if v5677 == int32(0) {
		v5701 = v5678
		goto L1627
	} else {
		goto L1634
	}
L1631:
	;
	v5670 = *(*int32)(unsafe.Add(mBase, uint32(v5663)+192))
	v5677 = v5670
	goto L1630
L1632:
	;
	goto L1633
L1633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5663)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5663)+188)) = v5664
	v5675 = F_strchr(m, v5664, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5663)+192)) = v5675
	v5677 = v5675
	goto L1630
L1634:
	;
	if base.Ui32(v5667) <= base.Ui32(v5677) {
		v5701 = v5678
		goto L1627
	} else {
		goto L1635
	}
L1635:
	;
	v5682 = v5677
	v5684 = v5678
	goto L1636
L1636:
	;
	v5687 = int32(1)
	v5688 = v5684 + v5687
	*(*int32)(unsafe.Add(mBase, uint32(v5663)+196)) = v5688
	v5691 = v5682 + v5687
	*(*int32)(unsafe.Add(mBase, uint32(v5663)+188)) = v5691
	v5694 = F_strchr(m, v5691, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5663)+192)) = v5694
	if v5694 == int32(0) {
		v5701 = v5688
		goto L1627
	} else {
		goto L1638
	}
L1637:
	;
	v5701 = v5688
	goto L1627
L1638:
	;
	if base.Ui32(v5694) < base.Ui32(v5667) {
		v5682 = v5694
		v5684 = v5688
		goto L1636
	} else {
		goto L1639
	}
L1639:
	;
	goto L1637
L1640:
	;
	if int32(0) <= v5714 {
		goto L1641
	} else {
		goto L1642
	}
L1641:
	;
	if v5719 != int32(59) {
		goto L35
	} else {
		goto L1644
	}
L1642:
	;
	goto L1643
L1643:
	;
	if v5719 != int32(277) {
		goto L1645
	} else {
		goto L1646
	}
L1644:
	;
	v5726 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5727 = *(*int32)(unsafe.Add(mBase, uint32(v5726)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+16)) = v5727
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5654
	v10149 = v241
	goto L5
L1645:
	;
	F_plpgsql_push_back_token(m, v5719, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5758 = m.ExcPending
	if v5758 != 0 {
		goto L66
	} else {
		goto L1652
	}
L1646:
	;
	v5732 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v5733 = m.ExcPending
	if v5733 != 0 {
		goto L66
	} else {
		goto L1647
	}
L1647:
	;
	if v5732 != int32(59) {
		goto L1645
	} else {
		goto L1648
	}
L1648:
	;
	v5736 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v5737 = *(*int32)(unsafe.Add(mBase, uint32(v5736)))
	if base.Ui32(int32(4)) < base.Ui32(v5737) {
		goto L1645
	} else {
		goto L1649
	}
L1649:
	;
	if v5737 == int32(3) {
		goto L1645
	} else {
		goto L1650
	}
L1650:
	;
	v5742 = *(*int32)(unsafe.Add(mBase, uint32(v5736)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+16)) = v5742
	v5748 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5749 = m.ExcPending
	if v5749 != 0 {
		goto L66
	} else {
		goto L1651
	}
L1651:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5654
	v10149 = v241
	goto L5
L1652:
	;
	v5760 = int32(0)
	v5764 = int32(1)
	v5772 = F_read_sql_construct(m, int32(59), v5760, v5760, int32(546849), int32(2), v5764, v5764, v5760, v5760, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5773 = m.ExcPending
	if v5773 != 0 {
		goto L66
	} else {
		goto L1653
	}
L1653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5654)+12)) = v5772
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5654
	v10149 = v241
	goto L5
L1654:
	;
	if v5800-v5799 != 0 {
		goto L1602
	} else {
		goto L1662
	}
L1655:
	;
	goto L1654
L1656:
	;
	if v5779 != v5780 {
		v5799 = v5779
		v5800 = v5780
		goto L1655
	} else {
		goto L1657
	}
L1657:
	;
	v5784 = v5617
	v5785 = v5776
	goto L1658
L1658:
	;
	v5788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5785)+1)))
	v5789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5784)+1)))
	if v5789 == int32(0) {
		v5799 = v5788
		v5800 = v5789
		goto L1655
	} else {
		goto L1660
	}
L1659:
	;
	v5799 = v5788
	v5800 = v5789
	goto L1655
L1660:
	;
	v5792 = int32(1)
	if v5788 == v5789 {
		v5784 = v5784 + v5792
		v5785 = v5785 + v5792
		goto L1658
	} else {
		goto L1661
	}
L1661:
	;
	goto L1659
L1662:
	;
	goto L1603
L1663:
	;
	v5810 = F_palloc0(m, int32(24))
	mBase = m.M
	v5811 = m.ExcPending
	if v5811 != 0 {
		goto L66
	} else {
		goto L1664
	}
L1664:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5810))) = int32(13)
	v5814 = int32(0)
	if v5803 < v5814 {
		v5857 = v5814
		goto L1666
	} else {
		goto L1667
	}
L1665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5810)+4)) = v5857
	v5862 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v5863 = *(*int32)(unsafe.Add(mBase, uint32(v5862)+520))
	v5865 = v5863 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5862)+520)) = v5865
	*(*int32)(unsafe.Add(mBase, uint32(v5810)+8)) = v5865
	v5872 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5873 = m.ExcPending
	if v5873 != 0 {
		goto L66
	} else {
		goto L1679
	}
L1666:
	;
	goto L1665
L1667:
	;
	v5819 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5820 = *(*int32)(unsafe.Add(mBase, uint32(v5819)+60))
	if v5820 == int32(0) {
		v5857 = v5814
		goto L1666
	} else {
		goto L1668
	}
L1668:
	;
	v5823 = v5803 + v5820
	v5824 = *(*int32)(unsafe.Add(mBase, uint32(v5819)+188))
	if base.Ui32(v5824) <= base.Ui32(v5823) {
		goto L1670
	} else {
		goto L1671
	}
L1669:
	;
	v5834 = *(*int32)(unsafe.Add(mBase, uint32(v5819)+196))
	if v5833 == int32(0) {
		v5857 = v5834
		goto L1666
	} else {
		goto L1673
	}
L1670:
	;
	v5826 = *(*int32)(unsafe.Add(mBase, uint32(v5819)+192))
	v5833 = v5826
	goto L1669
L1671:
	;
	goto L1672
L1672:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+188)) = v5820
	v5831 = F_strchr(m, v5820, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+192)) = v5831
	v5833 = v5831
	goto L1669
L1673:
	;
	if base.Ui32(v5823) <= base.Ui32(v5833) {
		v5857 = v5834
		goto L1666
	} else {
		goto L1674
	}
L1674:
	;
	v5838 = v5833
	v5840 = v5834
	goto L1675
L1675:
	;
	v5843 = int32(1)
	v5844 = v5840 + v5843
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+196)) = v5844
	v5847 = v5838 + v5843
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+188)) = v5847
	v5850 = F_strchr(m, v5847, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5819)+192)) = v5850
	if v5850 == int32(0) {
		v5857 = v5844
		goto L1666
	} else {
		goto L1677
	}
L1676:
	;
	v5857 = v5844
	goto L1666
L1677:
	;
	if base.Ui32(v5850) < base.Ui32(v5823) {
		v5838 = v5850
		v5840 = v5844
		goto L1675
	} else {
		goto L1678
	}
L1678:
	;
	goto L1676
L1679:
	;
	if v5872 != int32(317) {
		goto L1680
	} else {
		goto L1681
	}
L1680:
	;
	F_plpgsql_push_back_token(m, v5872, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5881 = m.ExcPending
	if v5881 != 0 {
		goto L66
	} else {
		goto L1683
	}
L1681:
	;
	goto L1682
L1682:
	;
	v5901 = int32(0)
	v5904 = int32(1)
	v5913 = F_read_sql_construct(m, int32(59), int32(381), v5901, int32(536093), int32(2), v5904, v5904, v5901, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5914 = m.ExcPending
	if v5914 != 0 {
		goto L66
	} else {
		goto L1685
	}
L1683:
	;
	v5883 = int32(0)
	v5895 = F_read_sql_construct(m, int32(59), v5883, v5883, int32(546849), v5883, v5883, int32(1), v5883, v5883, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5896 = m.ExcPending
	if v5896 != 0 {
		goto L66
	} else {
		goto L1684
	}
L1684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5810)+12)) = v5895
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5810
	v10149 = v241
	goto L5
L1685:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5810)+16)) = v5913
	v5916 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v5916 == int32(381) {
		goto L1686
	} else {
		goto L1687
	}
L1686:
	;
	goto L1689
L1687:
	;
	goto L1688
L1688:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5810
	v10149 = v241
	goto L5
L1689:
	;
	v5945 = int32(0)
	v5948 = int32(1)
	v5957 = F_read_sql_construct(m, int32(44), int32(59), v5945, int32(546844), int32(2), v5948, v5948, v5945, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5958 = m.ExcPending
	if v5958 != 0 {
		goto L66
	} else {
		goto L1691
	}
L1690:
	;
	goto L1688
L1691:
	;
	v5959 = *(*int32)(unsafe.Add(mBase, uint32(v5810)+20))
	v5960 = F_lappend(m, v5959, v5957)
	mBase = m.M
	v5961 = m.ExcPending
	if v5961 != 0 {
		goto L66
	} else {
		goto L1692
	}
L1692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5810)+20)) = v5960
	v5963 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v5963 == int32(44) {
		goto L1689
	} else {
		goto L1693
	}
L1693:
	;
	goto L1690
L1694:
	;
	v5998 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v6000 = F_palloc0(m, int32(20))
	mBase = m.M
	v6001 = m.ExcPending
	if v6001 != 0 {
		goto L66
	} else {
		goto L1695
	}
L1695:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6000))) = int32(11)
	v6004 = int32(0)
	if v5998 < v6004 {
		v6047 = v6004
		goto L1697
	} else {
		goto L1698
	}
L1696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6000)+4)) = v6047
	v6052 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v6053 = *(*int32)(unsafe.Add(mBase, uint32(v6052)+520))
	v6054 = int32(1)
	v6055 = v6053 + v6054
	*(*int32)(unsafe.Add(mBase, uint32(v6052)+520)) = v6055
	*(*int64)(unsafe.Add(mBase, uint32(v6000)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v6000)+8)) = v6055
	v6060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6052)+63)))
	if v6060 == v6054 {
		goto L1711
	} else {
		goto L1712
	}
L1697:
	;
	goto L1696
L1698:
	;
	v6009 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6010 = *(*int32)(unsafe.Add(mBase, uint32(v6009)+60))
	if v6010 == int32(0) {
		v6047 = v6004
		goto L1697
	} else {
		goto L1699
	}
L1699:
	;
	v6013 = v5998 + v6010
	v6014 = *(*int32)(unsafe.Add(mBase, uint32(v6009)+188))
	if base.Ui32(v6014) <= base.Ui32(v6013) {
		goto L1701
	} else {
		goto L1702
	}
L1700:
	;
	v6024 = *(*int32)(unsafe.Add(mBase, uint32(v6009)+196))
	if v6023 == int32(0) {
		v6047 = v6024
		goto L1697
	} else {
		goto L1704
	}
L1701:
	;
	v6016 = *(*int32)(unsafe.Add(mBase, uint32(v6009)+192))
	v6023 = v6016
	goto L1700
L1702:
	;
	goto L1703
L1703:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6009)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6009)+188)) = v6010
	v6021 = F_strchr(m, v6010, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6009)+192)) = v6021
	v6023 = v6021
	goto L1700
L1704:
	;
	if base.Ui32(v6013) <= base.Ui32(v6023) {
		v6047 = v6024
		goto L1697
	} else {
		goto L1705
	}
L1705:
	;
	v6028 = v6023
	v6030 = v6024
	goto L1706
L1706:
	;
	v6033 = int32(1)
	v6034 = v6030 + v6033
	*(*int32)(unsafe.Add(mBase, uint32(v6009)+196)) = v6034
	v6037 = v6028 + v6033
	*(*int32)(unsafe.Add(mBase, uint32(v6009)+188)) = v6037
	v6040 = F_strchr(m, v6037, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6009)+192)) = v6040
	if v6040 == int32(0) {
		v6047 = v6034
		goto L1697
	} else {
		goto L1708
	}
L1707:
	;
	v6047 = v6034
	goto L1697
L1708:
	;
	if base.Ui32(v6040) < base.Ui32(v6013) {
		v6028 = v6040
		v6030 = v6034
		goto L1706
	} else {
		goto L1709
	}
L1709:
	;
	goto L1707
L1710:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6000
	v10149 = v241
	goto L5
L1711:
	;
	v6067 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6068 = m.ExcPending
	if v6068 != 0 {
		goto L66
	} else {
		goto L1714
	}
L1712:
	;
	goto L1713
L1713:
	;
	v6094 = *(*int32)(unsafe.Add(mBase, uint32(v6052)+52))
	if v6094 == int32(2278) {
		goto L1722
	} else {
		goto L1723
	}
L1714:
	;
	if v6067 == int32(59) {
		goto L1710
	} else {
		goto L1715
	}
L1715:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v6074 = m.ExcPending
	if v6074 != 0 {
		goto L66
	} else {
		goto L1716
	}
L1716:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v6077 = m.ExcPending
	if v6077 != 0 {
		goto L66
	} else {
		goto L1717
	}
L1717:
	;
	F_errmsg(m, int32(106458), int32(0))
	mBase = m.M
	v6081 = m.ExcPending
	if v6081 != 0 {
		goto L66
	} else {
		goto L1718
	}
L1718:
	;
	F_errhint(m, int32(655643), int32(0))
	mBase = m.M
	v6085 = m.ExcPending
	if v6085 != 0 {
		goto L66
	} else {
		goto L1719
	}
L1719:
	;
	v6086 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	v6087 = F_plpgsql_scanner_errposition(m, v6086, l1)
	mBase = m.M
	v6088 = m.ExcPending
	if v6088 != 0 {
		goto L66
	} else {
		goto L1720
	}
L1720:
	;
	F_errfinish(m, int32(26959), int32(3373), int32(97047))
	mBase = m.M
	v6093 = m.ExcPending
	if v6093 != 0 {
		goto L66
	} else {
		goto L1721
	}
L1721:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1722:
	;
	v6101 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6102 = m.ExcPending
	if v6102 != 0 {
		goto L66
	} else {
		goto L1725
	}
L1723:
	;
	goto L1724
L1724:
	;
	v6129 = *(*int32)(unsafe.Add(mBase, uint32(v6052)+472))
	v6134 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6135 = m.ExcPending
	if v6135 != 0 {
		goto L66
	} else {
		goto L1733
	}
L1725:
	;
	if v6101 == int32(59) {
		goto L1710
	} else {
		goto L1726
	}
L1726:
	;
	v6106 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v6107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6106)+65)))
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v6111 = m.ExcPending
	if v6111 != 0 {
		goto L66
	} else {
		goto L1727
	}
L1727:
	;
	if v6107 == int32(112) {
		goto L33
	} else {
		goto L1728
	}
L1728:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v6116 = m.ExcPending
	if v6116 != 0 {
		goto L66
	} else {
		goto L1729
	}
L1729:
	;
	F_errmsg(m, int32(433319), int32(0))
	mBase = m.M
	v6120 = m.ExcPending
	if v6120 != 0 {
		goto L66
	} else {
		goto L1730
	}
L1730:
	;
	v6121 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	v6122 = F_plpgsql_scanner_errposition(m, v6121, l1)
	mBase = m.M
	v6123 = m.ExcPending
	if v6123 != 0 {
		goto L66
	} else {
		goto L1731
	}
L1731:
	;
	F_errfinish(m, int32(26959), int32(3388), int32(97047))
	mBase = m.M
	v6128 = m.ExcPending
	if v6128 != 0 {
		goto L66
	} else {
		goto L1732
	}
L1732:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1733:
	;
	if int32(0) <= v6129 {
		goto L1734
	} else {
		goto L1735
	}
L1734:
	;
	if v6134 != int32(59) {
		goto L32
	} else {
		goto L1737
	}
L1735:
	;
	goto L1736
L1736:
	;
	if v6134 != int32(277) {
		goto L1738
	} else {
		goto L1739
	}
L1737:
	;
	v6141 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v6142 = *(*int32)(unsafe.Add(mBase, uint32(v6141)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v6000)+16)) = v6142
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6000
	v10149 = v241
	goto L5
L1738:
	;
	F_plpgsql_push_back_token(m, v6134, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6173 = m.ExcPending
	if v6173 != 0 {
		goto L66
	} else {
		goto L1745
	}
L1739:
	;
	v6147 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v6148 = m.ExcPending
	if v6148 != 0 {
		goto L66
	} else {
		goto L1740
	}
L1740:
	;
	if v6147 != int32(59) {
		goto L1738
	} else {
		goto L1741
	}
L1741:
	;
	v6151 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v6152 = *(*int32)(unsafe.Add(mBase, uint32(v6151)))
	if base.Ui32(int32(4)) < base.Ui32(v6152) {
		goto L1738
	} else {
		goto L1742
	}
L1742:
	;
	if v6152 == int32(3) {
		goto L1738
	} else {
		goto L1743
	}
L1743:
	;
	v6157 = *(*int32)(unsafe.Add(mBase, uint32(v6151)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6000)+16)) = v6157
	v6163 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6164 = m.ExcPending
	if v6164 != 0 {
		goto L66
	} else {
		goto L1744
	}
L1744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6000
	v10149 = v241
	goto L5
L1745:
	;
	v6175 = int32(0)
	v6179 = int32(1)
	v6187 = F_read_sql_construct(m, int32(59), v6175, v6175, int32(546849), int32(2), v6179, v6179, v6175, v6175, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6188 = m.ExcPending
	if v6188 != 0 {
		goto L66
	} else {
		goto L1746
	}
L1746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6000)+12)) = v6187
	goto L1710
L1747:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6195))) = int32(14)
	v6199 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v6200 = int32(0)
	if v6199 < v6200 {
		v6243 = v6200
		goto L1749
	} else {
		goto L1750
	}
L1748:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6195)+4)) = v6243
	v6248 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v6249 = *(*int32)(unsafe.Add(mBase, uint32(v6248)+520))
	v6251 = v6249 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6248)+520)) = v6251
	v6253 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6195)+16)) = v6253
	v6255 = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v6195)+12)) = v6255
	*(*int32)(unsafe.Add(mBase, uint32(v6195)+8)) = v6251
	*(*int64)(unsafe.Add(mBase, uint32(v6195)+24)) = v6253
	v6265 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6266 = m.ExcPending
	if v6266 != 0 {
		goto L66
	} else {
		goto L1762
	}
L1749:
	;
	goto L1748
L1750:
	;
	v6205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6206 = *(*int32)(unsafe.Add(mBase, uint32(v6205)+60))
	if v6206 == int32(0) {
		v6243 = v6200
		goto L1749
	} else {
		goto L1751
	}
L1751:
	;
	v6209 = v6199 + v6206
	v6210 = *(*int32)(unsafe.Add(mBase, uint32(v6205)+188))
	if base.Ui32(v6210) <= base.Ui32(v6209) {
		goto L1753
	} else {
		goto L1754
	}
L1752:
	;
	v6220 = *(*int32)(unsafe.Add(mBase, uint32(v6205)+196))
	if v6219 == int32(0) {
		v6243 = v6220
		goto L1749
	} else {
		goto L1756
	}
L1753:
	;
	v6212 = *(*int32)(unsafe.Add(mBase, uint32(v6205)+192))
	v6219 = v6212
	goto L1752
L1754:
	;
	goto L1755
L1755:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6205)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6205)+188)) = v6206
	v6217 = F_strchr(m, v6206, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6205)+192)) = v6217
	v6219 = v6217
	goto L1752
L1756:
	;
	if base.Ui32(v6209) <= base.Ui32(v6219) {
		v6243 = v6220
		goto L1749
	} else {
		goto L1757
	}
L1757:
	;
	v6224 = v6219
	v6226 = v6220
	goto L1758
L1758:
	;
	v6229 = int32(1)
	v6230 = v6226 + v6229
	*(*int32)(unsafe.Add(mBase, uint32(v6205)+196)) = v6230
	v6233 = v6224 + v6229
	*(*int32)(unsafe.Add(mBase, uint32(v6205)+188)) = v6233
	v6236 = F_strchr(m, v6233, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6205)+192)) = v6236
	if v6236 == int32(0) {
		v6243 = v6230
		goto L1749
	} else {
		goto L1760
	}
L1759:
	;
	v6243 = v6230
	goto L1749
L1760:
	;
	if base.Ui32(v6236) < base.Ui32(v6209) {
		v6224 = v6236
		v6226 = v6230
		goto L1758
	} else {
		goto L1761
	}
L1761:
	;
	goto L1759
L1762:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v6265
	v6269 = v6195 + int32(16)
	if v6265 <= int32(303) {
		goto L1785
	} else {
		goto L1786
	}
L1763:
	;
	v7152 = *(*int32)(unsafe.Add(mBase, uint32(v6195)+20))
	if v7152 != 0 {
		goto L2032
	} else {
		goto L2033
	}
L1764:
	;
	v6799 = int32(0)
	goto L1922
L1765:
	;
	if v6745 != int32(381) {
		goto L1763
	} else {
		goto L1921
	}
L1766:
	;
	goto L1916
L1767:
	;
	v6682 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6683 = m.ExcPending
	if v6683 != 0 {
		goto L66
	} else {
		goto L1912
	}
L1768:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6269))) = v6671
	v6674 = F_plpgsql_recognize_err_condition(m, v6671, int32(0))
	mBase = m.M
	v6675 = m.ExcPending
	if v6675 != 0 {
		goto L66
	} else {
		goto L1911
	}
L1769:
	;
	v6668 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v6671 = v6668
	goto L1768
L1770:
	;
	v6644 = int32(0)
	goto L1900
L1771:
	;
	v6544 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6545 = m.ExcPending
	if v6545 != 0 {
		goto L66
	} else {
		goto L1874
	}
L1772:
	;
	v6507 = int32(277)
	v6508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v6508 != 0 {
		v6639 = v6507
		goto L1770
	} else {
		goto L1863
	}
L1773:
	;
	switch v6474 - int32(261) {
	case 0:
		goto L1853
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15:
		v6639 = v6474
		goto L1770
	case 14:
		goto L1769
	case 16:
		goto L1772
	default:
		goto L1854
	}
L1774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6195)+12)) = v6465
	v6471 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6472 = m.ExcPending
	if v6472 != 0 {
		goto L66
	} else {
		goto L1852
	}
L1775:
	;
	v6465 = int32(14)
	goto L1774
L1776:
	;
	v6436 = int32(326716)
	v6439 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1400])))
	v6440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6405))))
	if v6440 == int32(0) {
		v6459 = v6439
		v6460 = v6440
		goto L1844
	} else {
		goto L1845
	}
L1777:
	;
	v6465 = int32(15)
	goto L1774
L1778:
	;
	v6405 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v6405 == int32(0) {
		goto L1772
	} else {
		goto L1833
	}
L1779:
	;
	v6465 = int32(17)
	goto L1774
L1780:
	;
	v6377 = int32(241915)
	v6380 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1401])))
	v6381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6346))))
	if v6381 == int32(0) {
		v6400 = v6380
		v6401 = v6381
		goto L1825
	} else {
		goto L1826
	}
L1781:
	;
	v6465 = int32(18)
	goto L1774
L1782:
	;
	v6346 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v6346 == int32(0) {
		goto L1772
	} else {
		goto L1814
	}
L1783:
	;
	v6465 = int32(19)
	goto L1774
L1784:
	;
	v6286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v6286 != 0 {
		goto L1772
	} else {
		goto L1794
	}
L1785:
	;
	if v6265 == int32(59) {
		goto L1763
	} else {
		goto L1788
	}
L1786:
	;
	goto L1787
L1787:
	;
	switch v6265 - int32(304) {
	case 0:
		goto L1775
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30, 32, 33, 34, 35, 36, 37, 38, 39:
		v6474 = v6265
		goto L1773
	case 12:
		v6465 = v6255
		goto L1774
	case 26:
		goto L1779
	case 31:
		goto L1777
	case 40:
		goto L1781
	default:
		goto L1792
	}
L1788:
	;
	if v6265 == int32(277) {
		goto L1784
	} else {
		goto L1789
	}
L1789:
	;
	if v6265 != 0 {
		v6474 = v6265
		goto L1773
	} else {
		goto L1790
	}
L1790:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(250277))
	mBase = m.M
	v6281 = m.ExcPending
	if v6281 != 0 {
		goto L66
	} else {
		goto L1791
	}
L1791:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1792:
	;
	if v6265 == int32(383) {
		goto L1783
	} else {
		goto L1793
	}
L1793:
	;
	v6474 = v6265
	goto L1773
L1794:
	;
	v6287 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v6287 == int32(0) {
		goto L1772
	} else {
		goto L1795
	}
L1795:
	;
	v6290 = int32(247690)
	v6293 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1402])))
	v6294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6287))))
	if v6294 == int32(0) {
		v6313 = v6293
		v6314 = v6294
		goto L1797
	} else {
		goto L1798
	}
L1796:
	;
	if v6314-v6313 == int32(0) {
		v6465 = v6255
		goto L1774
	} else {
		goto L1804
	}
L1797:
	;
	goto L1796
L1798:
	;
	if v6293 != v6294 {
		v6313 = v6293
		v6314 = v6294
		goto L1797
	} else {
		goto L1799
	}
L1799:
	;
	v6298 = v6287
	v6299 = v6290
	goto L1800
L1800:
	;
	v6302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6299)+1)))
	v6303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6298)+1)))
	if v6303 == int32(0) {
		v6313 = v6302
		v6314 = v6303
		goto L1797
	} else {
		goto L1802
	}
L1801:
	;
	v6313 = v6302
	v6314 = v6303
	goto L1797
L1802:
	;
	v6306 = int32(1)
	if v6302 == v6303 {
		v6298 = v6298 + v6306
		v6299 = v6299 + v6306
		goto L1800
	} else {
		goto L1803
	}
L1803:
	;
	goto L1801
L1804:
	;
	v6318 = int32(334569)
	v6321 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1403])))
	v6322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6287))))
	if v6322 == int32(0) {
		v6341 = v6321
		v6342 = v6322
		goto L1806
	} else {
		goto L1807
	}
L1805:
	;
	if v6342-v6341 != 0 {
		goto L1782
	} else {
		goto L1813
	}
L1806:
	;
	goto L1805
L1807:
	;
	if v6321 != v6322 {
		v6341 = v6321
		v6342 = v6322
		goto L1806
	} else {
		goto L1808
	}
L1808:
	;
	v6326 = v6287
	v6327 = v6318
	goto L1809
L1809:
	;
	v6330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6327)+1)))
	v6331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6326)+1)))
	if v6331 == int32(0) {
		v6341 = v6330
		v6342 = v6331
		goto L1806
	} else {
		goto L1811
	}
L1810:
	;
	v6341 = v6330
	v6342 = v6331
	goto L1806
L1811:
	;
	v6334 = int32(1)
	if v6330 == v6331 {
		v6326 = v6326 + v6334
		v6327 = v6327 + v6334
		goto L1809
	} else {
		goto L1812
	}
L1812:
	;
	goto L1810
L1813:
	;
	goto L1783
L1814:
	;
	v6349 = int32(417885)
	v6352 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1404])))
	v6353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6346))))
	if v6353 == int32(0) {
		v6372 = v6352
		v6373 = v6353
		goto L1816
	} else {
		goto L1817
	}
L1815:
	;
	if v6373-v6372 != 0 {
		goto L1780
	} else {
		goto L1823
	}
L1816:
	;
	goto L1815
L1817:
	;
	if v6352 != v6353 {
		v6372 = v6352
		v6373 = v6353
		goto L1816
	} else {
		goto L1818
	}
L1818:
	;
	v6357 = v6346
	v6358 = v6349
	goto L1819
L1819:
	;
	v6361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6358)+1)))
	v6362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6357)+1)))
	if v6362 == int32(0) {
		v6372 = v6361
		v6373 = v6362
		goto L1816
	} else {
		goto L1821
	}
L1820:
	;
	v6372 = v6361
	v6373 = v6362
	goto L1816
L1821:
	;
	v6365 = int32(1)
	if v6361 == v6362 {
		v6357 = v6357 + v6365
		v6358 = v6358 + v6365
		goto L1819
	} else {
		goto L1822
	}
L1822:
	;
	goto L1820
L1823:
	;
	goto L1781
L1824:
	;
	if v6401-v6400 != 0 {
		goto L1778
	} else {
		goto L1832
	}
L1825:
	;
	goto L1824
L1826:
	;
	if v6380 != v6381 {
		v6400 = v6380
		v6401 = v6381
		goto L1825
	} else {
		goto L1827
	}
L1827:
	;
	v6385 = v6346
	v6386 = v6377
	goto L1828
L1828:
	;
	v6389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6386)+1)))
	v6390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6385)+1)))
	if v6390 == int32(0) {
		v6400 = v6389
		v6401 = v6390
		goto L1825
	} else {
		goto L1830
	}
L1829:
	;
	v6400 = v6389
	v6401 = v6390
	goto L1825
L1830:
	;
	v6393 = int32(1)
	if v6389 == v6390 {
		v6385 = v6385 + v6393
		v6386 = v6386 + v6393
		goto L1828
	} else {
		goto L1831
	}
L1831:
	;
	goto L1829
L1832:
	;
	goto L1779
L1833:
	;
	v6408 = int32(327313)
	v6411 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1405])))
	v6412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6405))))
	if v6412 == int32(0) {
		v6431 = v6411
		v6432 = v6412
		goto L1835
	} else {
		goto L1836
	}
L1834:
	;
	if v6432-v6431 != 0 {
		goto L1776
	} else {
		goto L1842
	}
L1835:
	;
	goto L1834
L1836:
	;
	if v6411 != v6412 {
		v6431 = v6411
		v6432 = v6412
		goto L1835
	} else {
		goto L1837
	}
L1837:
	;
	v6416 = v6405
	v6417 = v6408
	goto L1838
L1838:
	;
	v6420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6417)+1)))
	v6421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6416)+1)))
	if v6421 == int32(0) {
		v6431 = v6420
		v6432 = v6421
		goto L1835
	} else {
		goto L1840
	}
L1839:
	;
	v6431 = v6420
	v6432 = v6421
	goto L1835
L1840:
	;
	v6424 = int32(1)
	if v6420 == v6421 {
		v6416 = v6416 + v6424
		v6417 = v6417 + v6424
		goto L1838
	} else {
		goto L1841
	}
L1841:
	;
	goto L1839
L1842:
	;
	goto L1777
L1843:
	;
	if v6460-v6459 != 0 {
		goto L1772
	} else {
		goto L1851
	}
L1844:
	;
	goto L1843
L1845:
	;
	if v6439 != v6440 {
		v6459 = v6439
		v6460 = v6440
		goto L1844
	} else {
		goto L1846
	}
L1846:
	;
	v6444 = v6405
	v6445 = v6436
	goto L1847
L1847:
	;
	v6448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6445)+1)))
	v6449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6444)+1)))
	if v6449 == int32(0) {
		v6459 = v6448
		v6460 = v6449
		goto L1844
	} else {
		goto L1849
	}
L1848:
	;
	v6459 = v6448
	v6460 = v6449
	goto L1844
L1849:
	;
	v6452 = int32(1)
	if v6448 == v6449 {
		v6444 = v6444 + v6452
		v6445 = v6445 + v6452
		goto L1847
	} else {
		goto L1850
	}
L1850:
	;
	goto L1848
L1851:
	;
	goto L1775
L1852:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v6471
	v6474 = v6471
	goto L1773
L1853:
	;
	v6486 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	*(*int32)(unsafe.Add(mBase, uint32(v6195)+20)) = v6486
	v6492 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6493 = m.ExcPending
	if v6493 != 0 {
		goto L66
	} else {
		goto L1858
	}
L1854:
	;
	switch v6474 - int32(371) {
	case 0:
		goto L1771
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v6639 = v6474
		goto L1770
	case 10:
		goto L1764
	default:
		goto L1855
	}
L1855:
	;
	if v6474 != 0 {
		v6639 = v6474
		goto L1770
	} else {
		goto L1856
	}
L1856:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(250277))
	mBase = m.M
	v6485 = m.ExcPending
	if v6485 != 0 {
		goto L66
	} else {
		goto L1857
	}
L1857:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1858:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v6492
	switch v6492 - int32(44) {
	case 0:
		goto L1766
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L1859
	case 15:
		v6745 = v6492
		goto L1765
	default:
		goto L1860
	}
L1859:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v6504 = m.ExcPending
	if v6504 != 0 {
		goto L66
	} else {
		goto L1862
	}
L1860:
	;
	if v6492 == int32(381) {
		v6745 = v6492
		goto L1765
	} else {
		goto L1861
	}
L1861:
	;
	goto L1859
L1862:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1863:
	;
	v6509 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v6509 == int32(0) {
		v6639 = v6507
		goto L1770
	} else {
		goto L1864
	}
L1864:
	;
	v6512 = int32(351311)
	v6515 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1366])))
	v6516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6509))))
	if v6516 == int32(0) {
		v6535 = v6515
		v6536 = v6516
		goto L1866
	} else {
		goto L1867
	}
L1865:
	;
	if v6536-v6535 != 0 {
		v6639 = v6507
		goto L1770
	} else {
		goto L1873
	}
L1866:
	;
	goto L1865
L1867:
	;
	if v6515 != v6516 {
		v6535 = v6515
		v6536 = v6516
		goto L1866
	} else {
		goto L1868
	}
L1868:
	;
	v6520 = v6509
	v6521 = v6512
	goto L1869
L1869:
	;
	v6524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6521)+1)))
	v6525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6520)+1)))
	if v6525 == int32(0) {
		v6535 = v6524
		v6536 = v6525
		goto L1866
	} else {
		goto L1871
	}
L1870:
	;
	v6535 = v6524
	v6536 = v6525
	goto L1866
L1871:
	;
	v6528 = int32(1)
	if v6524 == v6525 {
		v6520 = v6520 + v6528
		v6521 = v6521 + v6528
		goto L1869
	} else {
		goto L1872
	}
L1872:
	;
	goto L1870
L1873:
	;
	goto L1771
L1874:
	;
	if v6544 != int32(261) {
		goto L31
	} else {
		goto L1875
	}
L1875:
	;
	v6548 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v6549 = F_strlen(m, v6548)
	mBase = m.M
	if v6549 != int32(5) {
		goto L30
	} else {
		goto L1876
	}
L1876:
	;
	v6552 = int32(507831)
	v6556 = m.G0
	v6558 = v6556 - int32(32)
	v6559 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6558)+24)) = v6559
	*(*int64)(unsafe.Add(mBase, uint32(v6558)+16)) = v6559
	*(*int64)(unsafe.Add(mBase, uint32(v6558)+8)) = v6559
	*(*int64)(unsafe.Add(mBase, uint32(v6558))) = v6559
	v6567 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1406])))
	if v6567 == int32(0) {
		goto L1878
	} else {
		goto L1879
	}
L1877:
	;
	if v6635 != int32(5) {
		goto L29
	} else {
		goto L1898
	}
L1878:
	;
	v6635 = int32(0)
	goto L1877
L1879:
	;
	goto L1880
L1880:
	;
	v6571 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1407])))
	if v6571 == int32(0) {
		goto L1881
	} else {
		goto L1882
	}
L1881:
	;
	v6575 = v6548
	goto L1884
L1882:
	;
	goto L1883
L1883:
	;
	v6585 = v6552
	v6586 = v6567
	goto L1887
L1884:
	;
	v6581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6575))))
	if v6581 == v6567 {
		v6575 = v6575 + int32(1)
		goto L1884
	} else {
		goto L1886
	}
L1885:
	;
	v6635 = v6575 - v6548
	goto L1877
L1886:
	;
	goto L1885
L1887:
	;
	v6593 = v6558 + int32(base.Ui32(v6586)>>(uint(int32(3))%32))&int32(28)
	v6594 = *(*int32)(unsafe.Add(mBase, uint32(v6593)))
	v6595 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6593))) = v6594 | v6595<<(uint(v6586)%32)
	v6599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6585)+1)))
	if v6599 != 0 {
		v6585 = v6585 + v6595
		v6586 = v6599
		goto L1887
	} else {
		goto L1889
	}
L1888:
	;
	v6602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6548))))
	if v6602 == int32(0) {
		v6627 = v6548
		goto L1890
	} else {
		goto L1891
	}
L1889:
	;
	goto L1888
L1890:
	;
	v6635 = v6627 - v6548
	goto L1877
L1891:
	;
	v6606 = v6548
	v6607 = v6602
	goto L1892
L1892:
	;
	v6615 = *(*int32)(unsafe.Add(mBase, uint32(v6558+int32(base.Ui32(v6607)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v6615)>>(uint(v6607)%32))&int32(1) == int32(0) {
		goto L1894
	} else {
		goto L1895
	}
L1893:
	;
	v6627 = v6623
	goto L1890
L1894:
	;
	v6627 = v6606
	goto L1890
L1895:
	;
	goto L1896
L1896:
	;
	v6621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6606)+1)))
	v6623 = v6606 + int32(1)
	if v6621 != 0 {
		v6606 = v6623
		v6607 = v6621
		goto L1892
	} else {
		goto L1897
	}
L1897:
	;
	goto L1893
L1898:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6269))) = v6548
	goto L1767
L1899:
	;
	if v6639 == v6650 {
		goto L1906
	} else {
		goto L1907
	}
L1900:
	;
	v6650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6644<<(uint(int32(1))%32))+uint32(_consts[1373]))))
	v6651 = base.B2i32(v6639 == v6650)
	if v6651 == int32(0) {
		goto L1902
	} else {
		goto L1903
	}
L1901:
	;
	goto L1899
L1902:
	;
	v6655 = v6644 + int32(1)
	if v6655 != int32(83) {
		v6644 = v6655
		goto L1900
	} else {
		goto L1905
	}
L1903:
	;
	goto L1904
L1904:
	;
	goto L1901
L1905:
	;
	goto L1904
L1906:
	;
	v6659 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v6660 = F_pstrdup(m, v6659)
	mBase = m.M
	v6661 = m.ExcPending
	if v6661 != 0 {
		goto L66
	} else {
		goto L1909
	}
L1907:
	;
	goto L1908
L1908:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v6667 = m.ExcPending
	if v6667 != 0 {
		goto L66
	} else {
		goto L1910
	}
L1909:
	;
	v6671 = v6660
	goto L1768
L1910:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1911:
	;
	goto L1767
L1912:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v6682
	if v6682 == int32(59) {
		v6745 = v6682
		goto L1765
	} else {
		goto L1913
	}
L1913:
	;
	if v6682 == int32(381) {
		v6745 = v6682
		goto L1765
	} else {
		goto L1914
	}
L1914:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v6694 = m.ExcPending
	if v6694 != 0 {
		goto L66
	} else {
		goto L1915
	}
L1915:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1916:
	;
	v6724 = int32(1)
	v6733 = F_read_sql_construct(m, int32(44), int32(59), int32(381), int32(536088), int32(2), v6724, v6724, int32(0), v27+int32(4768), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6734 = m.ExcPending
	if v6734 != 0 {
		goto L66
	} else {
		goto L1918
	}
L1917:
	;
	v6745 = v6739
	goto L1765
L1918:
	;
	v6735 = *(*int32)(unsafe.Add(mBase, uint32(v6195)+24))
	v6736 = F_lappend(m, v6735, v6733)
	mBase = m.M
	v6737 = m.ExcPending
	if v6737 != 0 {
		goto L66
	} else {
		goto L1919
	}
L1919:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6195)+24)) = v6736
	v6739 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379])))
	if v6739 == int32(44) {
		goto L1916
	} else {
		goto L1920
	}
L1920:
	;
	goto L1917
L1921:
	;
	goto L1764
L1922:
	;
	v6821 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6822 = m.ExcPending
	if v6822 != 0 {
		goto L66
	} else {
		goto L1924
	}
L1923:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6195)+28)) = v7121
	goto L1763
L1924:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = v6821
	if v6821 == int32(0) {
		goto L28
	} else {
		goto L1925
	}
L1925:
	;
	v6827 = F_palloc(m, int32(8))
	mBase = m.M
	v6828 = m.ExcPending
	if v6828 != 0 {
		goto L66
	} else {
		goto L1926
	}
L1926:
	;
	v6829 = int32(0)
	v6830 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	switch v6830 - int32(277) {
	case 0:
		goto L1943
	default:
		goto L27
	case 17:
		goto L1936
	case 21:
		goto L1934
	case 26:
		goto L1932
	case 30:
		goto L1940
	case 37:
		v7089 = v6829
		goto L1927
	case 49:
		goto L1938
	case 61:
		goto L1942
	case 90:
		goto L1928
	case 97:
		goto L1930
	}
L1927:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6827))) = v7089
	v7096 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7097 = m.ExcPending
	if v7097 != 0 {
		goto L66
	} else {
		goto L2027
	}
L1928:
	;
	v7089 = int32(8)
	goto L1927
L1929:
	;
	v7061 = int32(506420)
	v7064 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1408])))
	v7065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v7065 == int32(0) {
		v7084 = v7064
		v7085 = v7065
		goto L2019
	} else {
		goto L2020
	}
L1930:
	;
	v7089 = int32(7)
	goto L1927
L1931:
	;
	v7033 = int32(395078)
	v7036 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1409])))
	v7037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v7037 == int32(0) {
		v7056 = v7036
		v7057 = v7037
		goto L2010
	} else {
		goto L2011
	}
L1932:
	;
	v7089 = int32(6)
	goto L1927
L1933:
	;
	v7005 = int32(366016)
	v7008 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1410])))
	v7009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v7009 == int32(0) {
		v7028 = v7008
		v7029 = v7009
		goto L2001
	} else {
		goto L2002
	}
L1934:
	;
	v7089 = int32(5)
	goto L1927
L1935:
	;
	v6977 = int32(90324)
	v6980 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1411])))
	v6981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v6981 == int32(0) {
		v7000 = v6980
		v7001 = v6981
		goto L1992
	} else {
		goto L1993
	}
L1936:
	;
	v7089 = int32(4)
	goto L1927
L1937:
	;
	v6949 = int32(274808)
	v6952 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1412])))
	v6953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v6953 == int32(0) {
		v6972 = v6952
		v6973 = v6953
		goto L1983
	} else {
		goto L1984
	}
L1938:
	;
	v7089 = int32(3)
	goto L1927
L1939:
	;
	v6921 = int32(89119)
	v6924 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1413])))
	v6925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v6925 == int32(0) {
		v6944 = v6924
		v6945 = v6925
		goto L1974
	} else {
		goto L1975
	}
L1940:
	;
	v7089 = int32(2)
	goto L1927
L1941:
	;
	v6893 = int32(305588)
	v6896 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1414])))
	v6897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v6897 == int32(0) {
		v6916 = v6896
		v6917 = v6897
		goto L1965
	} else {
		goto L1966
	}
L1942:
	;
	v7089 = int32(1)
	goto L1927
L1943:
	;
	v6833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v6833 != 0 {
		goto L27
	} else {
		goto L1944
	}
L1944:
	;
	v6834 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v6834 == int32(0) {
		goto L27
	} else {
		goto L1945
	}
L1945:
	;
	v6837 = int32(413442)
	v6840 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1415])))
	v6841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v6841 == int32(0) {
		v6860 = v6840
		v6861 = v6841
		goto L1947
	} else {
		goto L1948
	}
L1946:
	;
	if v6861-v6860 == int32(0) {
		v7089 = v6829
		goto L1927
	} else {
		goto L1954
	}
L1947:
	;
	goto L1946
L1948:
	;
	if v6840 != v6841 {
		v6860 = v6840
		v6861 = v6841
		goto L1947
	} else {
		goto L1949
	}
L1949:
	;
	v6845 = v6834
	v6846 = v6837
	goto L1950
L1950:
	;
	v6849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6846)+1)))
	v6850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6845)+1)))
	if v6850 == int32(0) {
		v6860 = v6849
		v6861 = v6850
		goto L1947
	} else {
		goto L1952
	}
L1951:
	;
	v6860 = v6849
	v6861 = v6850
	goto L1947
L1952:
	;
	v6853 = int32(1)
	if v6849 == v6850 {
		v6845 = v6845 + v6853
		v6846 = v6846 + v6853
		goto L1950
	} else {
		goto L1953
	}
L1953:
	;
	goto L1951
L1954:
	;
	v6865 = int32(405063)
	v6868 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1416])))
	v6869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6834))))
	if v6869 == int32(0) {
		v6888 = v6868
		v6889 = v6869
		goto L1956
	} else {
		goto L1957
	}
L1955:
	;
	if v6889-v6888 != 0 {
		goto L1941
	} else {
		goto L1963
	}
L1956:
	;
	goto L1955
L1957:
	;
	if v6868 != v6869 {
		v6888 = v6868
		v6889 = v6869
		goto L1956
	} else {
		goto L1958
	}
L1958:
	;
	v6873 = v6834
	v6874 = v6865
	goto L1959
L1959:
	;
	v6877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6874)+1)))
	v6878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6873)+1)))
	if v6878 == int32(0) {
		v6888 = v6877
		v6889 = v6878
		goto L1956
	} else {
		goto L1961
	}
L1960:
	;
	v6888 = v6877
	v6889 = v6878
	goto L1956
L1961:
	;
	v6881 = int32(1)
	if v6877 == v6878 {
		v6873 = v6873 + v6881
		v6874 = v6874 + v6881
		goto L1959
	} else {
		goto L1962
	}
L1962:
	;
	goto L1960
L1963:
	;
	goto L1942
L1964:
	;
	if v6917-v6916 != 0 {
		goto L1939
	} else {
		goto L1972
	}
L1965:
	;
	goto L1964
L1966:
	;
	if v6896 != v6897 {
		v6916 = v6896
		v6917 = v6897
		goto L1965
	} else {
		goto L1967
	}
L1967:
	;
	v6901 = v6834
	v6902 = v6893
	goto L1968
L1968:
	;
	v6905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6902)+1)))
	v6906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6901)+1)))
	if v6906 == int32(0) {
		v6916 = v6905
		v6917 = v6906
		goto L1965
	} else {
		goto L1970
	}
L1969:
	;
	v6916 = v6905
	v6917 = v6906
	goto L1965
L1970:
	;
	v6909 = int32(1)
	if v6905 == v6906 {
		v6901 = v6901 + v6909
		v6902 = v6902 + v6909
		goto L1968
	} else {
		goto L1971
	}
L1971:
	;
	goto L1969
L1972:
	;
	goto L1940
L1973:
	;
	if v6945-v6944 != 0 {
		goto L1937
	} else {
		goto L1981
	}
L1974:
	;
	goto L1973
L1975:
	;
	if v6924 != v6925 {
		v6944 = v6924
		v6945 = v6925
		goto L1974
	} else {
		goto L1976
	}
L1976:
	;
	v6929 = v6834
	v6930 = v6921
	goto L1977
L1977:
	;
	v6933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6930)+1)))
	v6934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6929)+1)))
	if v6934 == int32(0) {
		v6944 = v6933
		v6945 = v6934
		goto L1974
	} else {
		goto L1979
	}
L1978:
	;
	v6944 = v6933
	v6945 = v6934
	goto L1974
L1979:
	;
	v6937 = int32(1)
	if v6933 == v6934 {
		v6929 = v6929 + v6937
		v6930 = v6930 + v6937
		goto L1977
	} else {
		goto L1980
	}
L1980:
	;
	goto L1978
L1981:
	;
	goto L1938
L1982:
	;
	if v6973-v6972 != 0 {
		goto L1935
	} else {
		goto L1990
	}
L1983:
	;
	goto L1982
L1984:
	;
	if v6952 != v6953 {
		v6972 = v6952
		v6973 = v6953
		goto L1983
	} else {
		goto L1985
	}
L1985:
	;
	v6957 = v6834
	v6958 = v6949
	goto L1986
L1986:
	;
	v6961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6958)+1)))
	v6962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6957)+1)))
	if v6962 == int32(0) {
		v6972 = v6961
		v6973 = v6962
		goto L1983
	} else {
		goto L1988
	}
L1987:
	;
	v6972 = v6961
	v6973 = v6962
	goto L1983
L1988:
	;
	v6965 = int32(1)
	if v6961 == v6962 {
		v6957 = v6957 + v6965
		v6958 = v6958 + v6965
		goto L1986
	} else {
		goto L1989
	}
L1989:
	;
	goto L1987
L1990:
	;
	goto L1936
L1991:
	;
	if v7001-v7000 != 0 {
		goto L1933
	} else {
		goto L1999
	}
L1992:
	;
	goto L1991
L1993:
	;
	if v6980 != v6981 {
		v7000 = v6980
		v7001 = v6981
		goto L1992
	} else {
		goto L1994
	}
L1994:
	;
	v6985 = v6834
	v6986 = v6977
	goto L1995
L1995:
	;
	v6989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6986)+1)))
	v6990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6985)+1)))
	if v6990 == int32(0) {
		v7000 = v6989
		v7001 = v6990
		goto L1992
	} else {
		goto L1997
	}
L1996:
	;
	v7000 = v6989
	v7001 = v6990
	goto L1992
L1997:
	;
	v6993 = int32(1)
	if v6989 == v6990 {
		v6985 = v6985 + v6993
		v6986 = v6986 + v6993
		goto L1995
	} else {
		goto L1998
	}
L1998:
	;
	goto L1996
L1999:
	;
	goto L1934
L2000:
	;
	if v7029-v7028 != 0 {
		goto L1931
	} else {
		goto L2008
	}
L2001:
	;
	goto L2000
L2002:
	;
	if v7008 != v7009 {
		v7028 = v7008
		v7029 = v7009
		goto L2001
	} else {
		goto L2003
	}
L2003:
	;
	v7013 = v6834
	v7014 = v7005
	goto L2004
L2004:
	;
	v7017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7014)+1)))
	v7018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7013)+1)))
	if v7018 == int32(0) {
		v7028 = v7017
		v7029 = v7018
		goto L2001
	} else {
		goto L2006
	}
L2005:
	;
	v7028 = v7017
	v7029 = v7018
	goto L2001
L2006:
	;
	v7021 = int32(1)
	if v7017 == v7018 {
		v7013 = v7013 + v7021
		v7014 = v7014 + v7021
		goto L2004
	} else {
		goto L2007
	}
L2007:
	;
	goto L2005
L2008:
	;
	goto L1932
L2009:
	;
	if v7057-v7056 != 0 {
		goto L1929
	} else {
		goto L2017
	}
L2010:
	;
	goto L2009
L2011:
	;
	if v7036 != v7037 {
		v7056 = v7036
		v7057 = v7037
		goto L2010
	} else {
		goto L2012
	}
L2012:
	;
	v7041 = v6834
	v7042 = v7033
	goto L2013
L2013:
	;
	v7045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7042)+1)))
	v7046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7041)+1)))
	if v7046 == int32(0) {
		v7056 = v7045
		v7057 = v7046
		goto L2010
	} else {
		goto L2015
	}
L2014:
	;
	v7056 = v7045
	v7057 = v7046
	goto L2010
L2015:
	;
	v7049 = int32(1)
	if v7045 == v7046 {
		v7041 = v7041 + v7049
		v7042 = v7042 + v7049
		goto L2013
	} else {
		goto L2016
	}
L2016:
	;
	goto L2014
L2017:
	;
	goto L1930
L2018:
	;
	if v7085-v7084 != 0 {
		goto L27
	} else {
		goto L2026
	}
L2019:
	;
	goto L2018
L2020:
	;
	if v7064 != v7065 {
		v7084 = v7064
		v7085 = v7065
		goto L2019
	} else {
		goto L2021
	}
L2021:
	;
	v7069 = v6834
	v7070 = v7061
	goto L2022
L2022:
	;
	v7073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7070)+1)))
	v7074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7069)+1)))
	if v7074 == int32(0) {
		v7084 = v7073
		v7085 = v7074
		goto L2019
	} else {
		goto L2024
	}
L2023:
	;
	v7084 = v7073
	v7085 = v7074
	goto L2019
L2024:
	;
	v7077 = int32(1)
	if v7073 == v7074 {
		v7069 = v7069 + v7077
		v7070 = v7070 + v7077
		goto L2022
	} else {
		goto L2025
	}
L2025:
	;
	goto L2023
L2026:
	;
	goto L1928
L2027:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = v7096
	if base.B2i32(v7096 != int32(61))&base.B2i32(v7096 != int32(270)) != 0 {
		goto L26
	} else {
		goto L2028
	}
L2028:
	;
	v7106 = int32(0)
	v7109 = int32(1)
	v7118 = F_read_sql_construct(m, int32(44), int32(59), v7106, int32(546844), int32(2), v7109, v7109, v7106, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7119 = m.ExcPending
	if v7119 != 0 {
		goto L66
	} else {
		goto L2029
	}
L2029:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6827)+4)) = v7118
	v7121 = F_lappend(m, v6799, v6827)
	mBase = m.M
	v7122 = m.ExcPending
	if v7122 != 0 {
		goto L66
	} else {
		goto L2030
	}
L2030:
	;
	v7123 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v7123 != int32(59) {
		v6799 = v7121
		goto L1922
	} else {
		goto L2031
	}
L2031:
	;
	goto L1923
L2032:
	;
	v7156 = v7152
	v7159 = int32(0)
	goto L2036
L2033:
	;
	goto L2034
L2034:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6195
	v10149 = v241
	goto L5
L2035:
	;
	if v7198 < v7159 {
		goto L25
	} else {
		goto L2050
	}
L2036:
	;
	v7177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7156))))
	if v7177 != int32(37) {
		goto L2040
	} else {
		goto L2041
	}
L2037:
	;
	v7196 = *(*int32)(unsafe.Add(mBase, uint32(v7180)+4))
	if v7159 < v7196 {
		goto L11
	} else {
		goto L2049
	}
L2038:
	;
	goto L2037
L2039:
	;
	v7156 = v7191 + int32(1)
	v7159 = v7193
	goto L2036
L2040:
	;
	if v7177 != 0 {
		v7191 = v7156
		v7193 = v7159
		goto L2039
	} else {
		goto L2043
	}
L2041:
	;
	goto L2042
L2042:
	;
	v7186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7156)+1)))
	v7188 = base.B2i32(v7186 != int32(37))
	if v7186 != int32(37) {
		goto L2046
	} else {
		goto L2047
	}
L2043:
	;
	v7180 = *(*int32)(unsafe.Add(mBase, uint32(v6195)+24))
	if v7180 != 0 {
		goto L2038
	} else {
		goto L2044
	}
L2044:
	;
	v7181 = int32(0)
	if v7181 <= v7159 {
		v7198 = v7181
		goto L2035
	} else {
		goto L2045
	}
L2045:
	;
	goto L11
L2046:
	;
	v7189 = v7156
	goto L2048
L2047:
	;
	v7189 = v7156 + int32(1)
	goto L2048
L2048:
	;
	v7191 = v7189
	v7193 = v7188 + v7159
	goto L2039
L2049:
	;
	v7198 = v7196
	goto L2035
L2050:
	;
	goto L2034
L2051:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7226))) = int32(15)
	v7230 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7231 = int32(0)
	if v7230 < v7231 {
		v7274 = v7231
		goto L2053
	} else {
		goto L2054
	}
L2052:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7226)+4)) = v7274
	v7279 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v7280 = *(*int32)(unsafe.Add(mBase, uint32(v7279)+520))
	v7281 = int32(1)
	v7282 = v7280 + v7281
	*(*int32)(unsafe.Add(mBase, uint32(v7279)+520)) = v7282
	*(*int32)(unsafe.Add(mBase, uint32(v7226)+8)) = v7282
	v7287 = int32(0)
	v7299 = F_read_sql_construct(m, int32(44), int32(59), v7287, int32(546844), int32(2), v7281, v7281, v7287, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7300 = m.ExcPending
	if v7300 != 0 {
		goto L66
	} else {
		goto L2066
	}
L2053:
	;
	goto L2052
L2054:
	;
	v7236 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7237 = *(*int32)(unsafe.Add(mBase, uint32(v7236)+60))
	if v7237 == int32(0) {
		v7274 = v7231
		goto L2053
	} else {
		goto L2055
	}
L2055:
	;
	v7240 = v7230 + v7237
	v7241 = *(*int32)(unsafe.Add(mBase, uint32(v7236)+188))
	if base.Ui32(v7241) <= base.Ui32(v7240) {
		goto L2057
	} else {
		goto L2058
	}
L2056:
	;
	v7251 = *(*int32)(unsafe.Add(mBase, uint32(v7236)+196))
	if v7250 == int32(0) {
		v7274 = v7251
		goto L2053
	} else {
		goto L2060
	}
L2057:
	;
	v7243 = *(*int32)(unsafe.Add(mBase, uint32(v7236)+192))
	v7250 = v7243
	goto L2056
L2058:
	;
	goto L2059
L2059:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7236)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7236)+188)) = v7237
	v7248 = F_strchr(m, v7237, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7236)+192)) = v7248
	v7250 = v7248
	goto L2056
L2060:
	;
	if base.Ui32(v7240) <= base.Ui32(v7250) {
		v7274 = v7251
		goto L2053
	} else {
		goto L2061
	}
L2061:
	;
	v7255 = v7250
	v7257 = v7251
	goto L2062
L2062:
	;
	v7260 = int32(1)
	v7261 = v7257 + v7260
	*(*int32)(unsafe.Add(mBase, uint32(v7236)+196)) = v7261
	v7264 = v7255 + v7260
	*(*int32)(unsafe.Add(mBase, uint32(v7236)+188)) = v7264
	v7267 = F_strchr(m, v7264, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7236)+192)) = v7267
	if v7267 == int32(0) {
		v7274 = v7261
		goto L2053
	} else {
		goto L2064
	}
L2063:
	;
	v7274 = v7261
	goto L2053
L2064:
	;
	if base.Ui32(v7267) < base.Ui32(v7240) {
		v7255 = v7267
		v7257 = v7261
		goto L2062
	} else {
		goto L2065
	}
L2065:
	;
	goto L2063
L2066:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7226)+12)) = v7299
	v7302 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v7302 == int32(44) {
		goto L2067
	} else {
		goto L2068
	}
L2067:
	;
	v7306 = int32(0)
	v7310 = int32(1)
	v7318 = F_read_sql_construct(m, int32(59), v7306, v7306, int32(546849), int32(2), v7310, v7310, v7306, v7306, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7319 = m.ExcPending
	if v7319 != 0 {
		goto L66
	} else {
		goto L2070
	}
L2068:
	;
	v7321 = int32(0)
	goto L2069
L2069:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7226)+16)) = v7321
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7226
	v10149 = v241
	goto L5
L2070:
	;
	v7321 = v7318
	goto L2069
L2071:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7343
	v10149 = v241
	goto L5
L2072:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7353
	v10149 = v241
	goto L5
L2073:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7363
	v10149 = v241
	goto L5
L2074:
	;
	F_plpgsql_push_back_token(m, v7370, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7377 = m.ExcPending
	if v7377 != 0 {
		goto L66
	} else {
		goto L2075
	}
L2075:
	;
	switch v7370 - int32(46) {
	case 0, 15:
		goto L24
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2076
	default:
		goto L2077
	}
L2076:
	;
	v7385 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7390 = F_make_execsql_stmt(m, int32(275), v7385, v132, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7391 = m.ExcPending
	if v7391 != 0 {
		goto L66
	} else {
		goto L2080
	}
L2077:
	;
	if v7370 == int32(270) {
		goto L24
	} else {
		goto L2078
	}
L2078:
	;
	if v7370 == int32(91) {
		goto L24
	} else {
		goto L2079
	}
L2079:
	;
	goto L2076
L2080:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7390
	v10149 = v241
	goto L5
L2081:
	;
	F_plpgsql_push_back_token(m, v7397, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7404 = m.ExcPending
	if v7404 != 0 {
		goto L66
	} else {
		goto L2082
	}
L2082:
	;
	switch v7397 - int32(46) {
	case 0, 15:
		goto L23
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2083
	default:
		goto L2084
	}
L2083:
	;
	v7412 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7418 = F_make_execsql_stmt(m, int32(276), v7412, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7419 = m.ExcPending
	if v7419 != 0 {
		goto L66
	} else {
		goto L2087
	}
L2084:
	;
	if v7397 == int32(270) {
		goto L23
	} else {
		goto L2085
	}
L2085:
	;
	if v7397 == int32(91) {
		goto L23
	} else {
		goto L2086
	}
L2086:
	;
	goto L2083
L2087:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7418
	v10149 = v241
	goto L5
L2088:
	;
	v7438 = F_palloc(m, int32(28))
	mBase = m.M
	v7439 = m.ExcPending
	if v7439 != 0 {
		goto L66
	} else {
		goto L2089
	}
L2089:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7438))) = int32(17)
	v7442 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7443 = int32(0)
	if v7442 < v7443 {
		v7486 = v7443
		goto L2091
	} else {
		goto L2092
	}
L2090:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7438)+4)) = v7486
	v7491 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v7492 = *(*int32)(unsafe.Add(mBase, uint32(v7491)+520))
	v7494 = v7492 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7491)+520)) = v7494
	*(*int64)(unsafe.Add(mBase, uint32(v7438)+20)) = int64(0)
	v7498 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7438)+16)) = uint16(v7498)
	*(*int32)(unsafe.Add(mBase, uint32(v7438)+12)) = v7435
	*(*int32)(unsafe.Add(mBase, uint32(v7438)+8)) = v7494
	v7506 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	v7511 = v7506
	goto L2104
L2091:
	;
	goto L2090
L2092:
	;
	v7448 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7449 = *(*int32)(unsafe.Add(mBase, uint32(v7448)+60))
	if v7449 == int32(0) {
		v7486 = v7443
		goto L2091
	} else {
		goto L2093
	}
L2093:
	;
	v7452 = v7442 + v7449
	v7453 = *(*int32)(unsafe.Add(mBase, uint32(v7448)+188))
	if base.Ui32(v7453) <= base.Ui32(v7452) {
		goto L2095
	} else {
		goto L2096
	}
L2094:
	;
	v7463 = *(*int32)(unsafe.Add(mBase, uint32(v7448)+196))
	if v7462 == int32(0) {
		v7486 = v7463
		goto L2091
	} else {
		goto L2098
	}
L2095:
	;
	v7455 = *(*int32)(unsafe.Add(mBase, uint32(v7448)+192))
	v7462 = v7455
	goto L2094
L2096:
	;
	goto L2097
L2097:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7448)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7448)+188)) = v7449
	v7460 = F_strchr(m, v7449, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7448)+192)) = v7460
	v7462 = v7460
	goto L2094
L2098:
	;
	if base.Ui32(v7452) <= base.Ui32(v7462) {
		v7486 = v7463
		goto L2091
	} else {
		goto L2099
	}
L2099:
	;
	v7467 = v7462
	v7469 = v7463
	goto L2100
L2100:
	;
	v7472 = int32(1)
	v7473 = v7469 + v7472
	*(*int32)(unsafe.Add(mBase, uint32(v7448)+196)) = v7473
	v7476 = v7467 + v7472
	*(*int32)(unsafe.Add(mBase, uint32(v7448)+188)) = v7476
	v7479 = F_strchr(m, v7476, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7448)+192)) = v7479
	if v7479 == int32(0) {
		v7486 = v7473
		goto L2091
	} else {
		goto L2102
	}
L2101:
	;
	v7486 = v7473
	goto L2091
L2102:
	;
	if base.Ui32(v7479) < base.Ui32(v7452) {
		v7467 = v7479
		v7469 = v7473
		goto L2100
	} else {
		goto L2103
	}
L2103:
	;
	goto L2101
L2104:
	;
	if v7511 != int32(332) {
		goto L2108
	} else {
		goto L2109
	}
L2105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7438
	v10149 = v241
	goto L5
L2106:
	;
	goto L2105
L2107:
	;
	v7561 = *(*int32)(unsafe.Add(mBase, uint32(v7438)+24))
	if v7561 != 0 {
		goto L21
	} else {
		goto L2117
	}
L2108:
	;
	if v7511 == int32(381) {
		goto L2107
	} else {
		goto L2111
	}
L2109:
	;
	goto L2110
L2110:
	;
	v7543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7438)+16)))
	if v7543 == int32(1) {
		goto L22
	} else {
		goto L2114
	}
L2111:
	;
	if v7511 == int32(59) {
		goto L2106
	} else {
		goto L2112
	}
L2112:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(212048))
	mBase = m.M
	v7542 = m.ExcPending
	if v7542 != 0 {
		goto L66
	} else {
		goto L2113
	}
L2113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2114:
	;
	v7546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7438)+16)) = uint8(v7546)
	F_read_into_target(m, v7438+int32(20), v7438+int32(17), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7553 = m.ExcPending
	if v7553 != 0 {
		goto L66
	} else {
		goto L2115
	}
L2115:
	;
	v7558 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7559 = m.ExcPending
	if v7559 != 0 {
		goto L66
	} else {
		goto L2116
	}
L2116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = v7558
	v7511 = v7558
	goto L2104
L2117:
	;
	goto L2118
L2118:
	;
	v7591 = int32(1)
	v7600 = F_read_sql_construct(m, int32(44), int32(59), int32(332), int32(526684), int32(2), v7591, v7591, int32(0), v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7601 = m.ExcPending
	if v7601 != 0 {
		goto L66
	} else {
		goto L2120
	}
L2119:
	;
	v7511 = v7606
	goto L2104
L2120:
	;
	v7602 = *(*int32)(unsafe.Add(mBase, uint32(v7438)+24))
	v7603 = F_lappend(m, v7602, v7600)
	mBase = m.M
	v7604 = m.ExcPending
	if v7604 != 0 {
		goto L66
	} else {
		goto L2121
	}
L2121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7438)+24)) = v7603
	v7606 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v7606 == int32(44) {
		goto L2118
	} else {
		goto L2122
	}
L2122:
	;
	goto L2119
L2123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7611))) = int32(20)
	v7617 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v7618 = int32(0)
	if v7617 < v7618 {
		v7661 = v7618
		goto L2125
	} else {
		goto L2126
	}
L2124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+4)) = v7661
	v7666 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v7667 = *(*int32)(unsafe.Add(mBase, uint32(v7666)+520))
	v7669 = v7667 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7666)+520)) = v7669
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+8)) = v7669
	v7672 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v7673 = *(*int32)(unsafe.Add(mBase, uint32(v7672)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+16)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+12)) = v7673
	v7677 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v7678 = *(*int32)(unsafe.Add(mBase, uint32(v7677)+28))
	if v7678 == int32(0) {
		goto L2138
	} else {
		goto L2139
	}
L2125:
	;
	goto L2124
L2126:
	;
	v7623 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7624 = *(*int32)(unsafe.Add(mBase, uint32(v7623)+60))
	if v7624 == int32(0) {
		v7661 = v7618
		goto L2125
	} else {
		goto L2127
	}
L2127:
	;
	v7627 = v7617 + v7624
	v7628 = *(*int32)(unsafe.Add(mBase, uint32(v7623)+188))
	if base.Ui32(v7628) <= base.Ui32(v7627) {
		goto L2129
	} else {
		goto L2130
	}
L2128:
	;
	v7638 = *(*int32)(unsafe.Add(mBase, uint32(v7623)+196))
	if v7637 == int32(0) {
		v7661 = v7638
		goto L2125
	} else {
		goto L2132
	}
L2129:
	;
	v7630 = *(*int32)(unsafe.Add(mBase, uint32(v7623)+192))
	v7637 = v7630
	goto L2128
L2130:
	;
	goto L2131
L2131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7623)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7623)+188)) = v7624
	v7635 = F_strchr(m, v7624, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7623)+192)) = v7635
	v7637 = v7635
	goto L2128
L2132:
	;
	if base.Ui32(v7627) <= base.Ui32(v7637) {
		v7661 = v7638
		goto L2125
	} else {
		goto L2133
	}
L2133:
	;
	v7642 = v7637
	v7644 = v7638
	goto L2134
L2134:
	;
	v7647 = int32(1)
	v7648 = v7644 + v7647
	*(*int32)(unsafe.Add(mBase, uint32(v7623)+196)) = v7648
	v7651 = v7642 + v7647
	*(*int32)(unsafe.Add(mBase, uint32(v7623)+188)) = v7651
	v7654 = F_strchr(m, v7651, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7623)+192)) = v7654
	if v7654 == int32(0) {
		v7661 = v7648
		goto L2125
	} else {
		goto L2136
	}
L2135:
	;
	v7661 = v7648
	goto L2125
L2136:
	;
	if base.Ui32(v7654) < base.Ui32(v7627) {
		v7642 = v7654
		v7644 = v7648
		goto L2134
	} else {
		goto L2137
	}
L2137:
	;
	goto L2135
L2138:
	;
	v7681 = int32(2)
	v7686 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7687 = m.ExcPending
	if v7687 != 0 {
		goto L66
	} else {
		goto L2143
	}
L2139:
	;
	goto L2140
L2140:
	;
	v7922 = F_read_cursor_args(m, v7677, int32(59), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7923 = m.ExcPending
	if v7923 != 0 {
		goto L66
	} else {
		goto L2197
	}
L2141:
	;
	if v7790 != int32(321) {
		goto L12
	} else {
		goto L2181
	}
L2142:
	;
	v7781 = *(*int32)(unsafe.Add(mBase, uint32(v7611)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+16)) = v7781 | v7780
	v7788 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7789 = m.ExcPending
	if v7789 != 0 {
		goto L66
	} else {
		goto L2180
	}
L2143:
	;
	if v7686 == int32(369) {
		v7780 = v7681
		goto L2142
	} else {
		goto L2144
	}
L2144:
	;
	if v7686 != int32(342) {
		goto L2147
	} else {
		goto L2148
	}
L2145:
	;
	v7753 = int32(303759)
	v7756 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1417])))
	v7757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7752))))
	if v7757 == int32(0) {
		v7776 = v7756
		v7777 = v7757
		goto L2172
	} else {
		goto L2173
	}
L2146:
	;
	v7749 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v7749 == int32(0) {
		goto L12
	} else {
		goto L2170
	}
L2147:
	;
	if v7686 != int32(277) {
		v7790 = v7686
		goto L2141
	} else {
		goto L2150
	}
L2148:
	;
	goto L2149
L2149:
	;
	v7706 = int32(4)
	v7711 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7712 = m.ExcPending
	if v7712 != 0 {
		goto L66
	} else {
		goto L2156
	}
L2150:
	;
	v7694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v7694 != 0 {
		goto L12
	} else {
		goto L2151
	}
L2151:
	;
	v7695 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v7695 == int32(0) {
		goto L12
	} else {
		goto L2152
	}
L2152:
	;
	v7698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7695))))
	if v7698 != int32(110) {
		v7752 = v7695
		goto L2145
	} else {
		goto L2153
	}
L2153:
	;
	v7701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7695)+1)))
	if v7701 != int32(111) {
		goto L2146
	} else {
		goto L2154
	}
L2154:
	;
	v7704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7695)+2)))
	if v7704 != 0 {
		goto L2146
	} else {
		goto L2155
	}
L2155:
	;
	goto L2149
L2156:
	;
	if v7711 == int32(369) {
		v7780 = v7706
		goto L2142
	} else {
		goto L2157
	}
L2157:
	;
	if v7711 != int32(277) {
		v7790 = v7711
		goto L2141
	} else {
		goto L2158
	}
L2158:
	;
	v7717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v7717 != 0 {
		goto L12
	} else {
		goto L2159
	}
L2159:
	;
	v7718 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v7718 == int32(0) {
		goto L12
	} else {
		goto L2160
	}
L2160:
	;
	v7721 = int32(303759)
	v7724 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1417])))
	v7725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7718))))
	if v7725 == int32(0) {
		v7744 = v7724
		v7745 = v7725
		goto L2162
	} else {
		goto L2163
	}
L2161:
	;
	if v7745-v7744 == int32(0) {
		v7780 = v7706
		goto L2142
	} else {
		goto L2169
	}
L2162:
	;
	goto L2161
L2163:
	;
	if v7724 != v7725 {
		v7744 = v7724
		v7745 = v7725
		goto L2162
	} else {
		goto L2164
	}
L2164:
	;
	v7729 = v7718
	v7730 = v7721
	goto L2165
L2165:
	;
	v7733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7730)+1)))
	v7734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7729)+1)))
	if v7734 == int32(0) {
		v7744 = v7733
		v7745 = v7734
		goto L2162
	} else {
		goto L2167
	}
L2166:
	;
	v7744 = v7733
	v7745 = v7734
	goto L2162
L2167:
	;
	v7737 = int32(1)
	if v7733 == v7734 {
		v7729 = v7729 + v7737
		v7730 = v7730 + v7737
		goto L2165
	} else {
		goto L2168
	}
L2168:
	;
	goto L2166
L2169:
	;
	goto L12
L2170:
	;
	v7752 = v7749
	goto L2145
L2171:
	;
	if v7777-v7776 != 0 {
		goto L12
	} else {
		goto L2179
	}
L2172:
	;
	goto L2171
L2173:
	;
	if v7756 != v7757 {
		v7776 = v7756
		v7777 = v7757
		goto L2172
	} else {
		goto L2174
	}
L2174:
	;
	v7761 = v7752
	v7762 = v7753
	goto L2175
L2175:
	;
	v7765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7762)+1)))
	v7766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7761)+1)))
	if v7766 == int32(0) {
		v7776 = v7765
		v7777 = v7766
		goto L2172
	} else {
		goto L2177
	}
L2176:
	;
	v7776 = v7765
	v7777 = v7766
	goto L2172
L2177:
	;
	v7769 = int32(1)
	if v7765 == v7766 {
		v7761 = v7761 + v7769
		v7762 = v7762 + v7769
		goto L2175
	} else {
		goto L2178
	}
L2178:
	;
	goto L2176
L2179:
	;
	v7780 = v7681
	goto L2142
L2180:
	;
	v7790 = v7788
	goto L2141
L2181:
	;
	v7798 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7799 = m.ExcPending
	if v7799 != 0 {
		goto L66
	} else {
		goto L2182
	}
L2182:
	;
	if v7798 == int32(317) {
		goto L2183
	} else {
		goto L2184
	}
L2183:
	;
	v7804 = int32(0)
	v7807 = int32(1)
	v7816 = F_read_sql_construct(m, int32(381), int32(59), v7804, int32(546833), int32(2), v7807, v7807, v7804, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7817 = m.ExcPending
	if v7817 != 0 {
		goto L66
	} else {
		goto L2186
	}
L2184:
	;
	goto L2185
L2185:
	;
	F_plpgsql_push_back_token(m, v7798, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7899 = m.ExcPending
	if v7899 != 0 {
		goto L66
	} else {
		goto L2195
	}
L2186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+28)) = v7816
	v7819 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v7819 == int32(381) {
		goto L2187
	} else {
		goto L2188
	}
L2187:
	;
	goto L2190
L2188:
	;
	goto L2189
L2189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7611
	v10149 = v241
	goto L5
L2190:
	;
	v7848 = int32(0)
	v7851 = int32(1)
	v7860 = F_read_sql_construct(m, int32(44), int32(59), v7848, int32(546844), int32(2), v7851, v7851, v7848, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7861 = m.ExcPending
	if v7861 != 0 {
		goto L66
	} else {
		goto L2192
	}
L2191:
	;
	goto L2189
L2192:
	;
	v7862 = *(*int32)(unsafe.Add(mBase, uint32(v7611)+32))
	v7863 = F_lappend(m, v7862, v7860)
	mBase = m.M
	v7864 = m.ExcPending
	if v7864 != 0 {
		goto L66
	} else {
		goto L2193
	}
L2193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+32)) = v7863
	v7866 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	if v7866 == int32(44) {
		goto L2190
	} else {
		goto L2194
	}
L2194:
	;
	goto L2191
L2195:
	;
	v7901 = int32(0)
	v7913 = F_read_sql_construct(m, int32(59), v7901, v7901, int32(546849), v7901, v7901, int32(1), v7901, v7901, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7914 = m.ExcPending
	if v7914 != 0 {
		goto L66
	} else {
		goto L2196
	}
L2196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+24)) = v7913
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7611
	v10149 = v241
	goto L5
L2197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7611)+20)) = v7922
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7611
	v10149 = v241
	goto L5
L2198:
	;
	v7942 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7943 = m.ExcPending
	if v7943 != 0 {
		goto L66
	} else {
		goto L2199
	}
L2199:
	;
	if v7942 != int32(59) {
		goto L20
	} else {
		goto L2200
	}
L2200:
	;
	v7946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7928)+33)))
	if v7946 == int32(1) {
		goto L19
	} else {
		goto L2201
	}
L2201:
	;
	v7951 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v7952 = int32(0)
	if v7951 < v7952 {
		v7995 = v7952
		goto L2203
	} else {
		goto L2204
	}
L2202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7928)+4)) = v7995
	v7999 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	*(*int32)(unsafe.Add(mBase, uint32(v7928)+12)) = v7999
	v8003 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8004 = *(*int32)(unsafe.Add(mBase, uint32(v8003)+4))
	v8005 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7928)+32)) = uint8(v8005)
	*(*int32)(unsafe.Add(mBase, uint32(v7928)+16)) = v8004
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7928
	v10149 = v241
	goto L5
L2203:
	;
	goto L2202
L2204:
	;
	v7957 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7958 = *(*int32)(unsafe.Add(mBase, uint32(v7957)+60))
	if v7958 == int32(0) {
		v7995 = v7952
		goto L2203
	} else {
		goto L2205
	}
L2205:
	;
	v7961 = v7951 + v7958
	v7962 = *(*int32)(unsafe.Add(mBase, uint32(v7957)+188))
	if base.Ui32(v7962) <= base.Ui32(v7961) {
		goto L2207
	} else {
		goto L2208
	}
L2206:
	;
	v7972 = *(*int32)(unsafe.Add(mBase, uint32(v7957)+196))
	if v7971 == int32(0) {
		v7995 = v7972
		goto L2203
	} else {
		goto L2210
	}
L2207:
	;
	v7964 = *(*int32)(unsafe.Add(mBase, uint32(v7957)+192))
	v7971 = v7964
	goto L2206
L2208:
	;
	goto L2209
L2209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7957)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7957)+188)) = v7958
	v7969 = F_strchr(m, v7958, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7957)+192)) = v7969
	v7971 = v7969
	goto L2206
L2210:
	;
	if base.Ui32(v7961) <= base.Ui32(v7971) {
		v7995 = v7972
		goto L2203
	} else {
		goto L2211
	}
L2211:
	;
	v7976 = v7971
	v7978 = v7972
	goto L2212
L2212:
	;
	v7981 = int32(1)
	v7982 = v7978 + v7981
	*(*int32)(unsafe.Add(mBase, uint32(v7957)+196)) = v7982
	v7985 = v7976 + v7981
	*(*int32)(unsafe.Add(mBase, uint32(v7957)+188)) = v7985
	v7988 = F_strchr(m, v7985, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7957)+192)) = v7988
	if v7988 == int32(0) {
		v7995 = v7982
		goto L2203
	} else {
		goto L2214
	}
L2213:
	;
	v7995 = v7982
	goto L2203
L2214:
	;
	if base.Ui32(v7988) < base.Ui32(v7961) {
		v7976 = v7988
		v7978 = v7982
		goto L2212
	} else {
		goto L2215
	}
L2215:
	;
	goto L2213
L2216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8011)+4)) = v8058
	v8064 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8065 = *(*int32)(unsafe.Add(mBase, uint32(v8064)+4))
	v8066 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8011)+32)) = uint8(v8066)
	*(*int32)(unsafe.Add(mBase, uint32(v8011)+16)) = v8065
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8011
	v10149 = v241
	goto L5
L2217:
	;
	goto L2216
L2218:
	;
	v8020 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8021 = *(*int32)(unsafe.Add(mBase, uint32(v8020)+60))
	if v8021 == int32(0) {
		v8058 = v8015
		goto L2217
	} else {
		goto L2219
	}
L2219:
	;
	v8024 = v8014 + v8021
	v8025 = *(*int32)(unsafe.Add(mBase, uint32(v8020)+188))
	if base.Ui32(v8025) <= base.Ui32(v8024) {
		goto L2221
	} else {
		goto L2222
	}
L2220:
	;
	v8035 = *(*int32)(unsafe.Add(mBase, uint32(v8020)+196))
	if v8034 == int32(0) {
		v8058 = v8035
		goto L2217
	} else {
		goto L2224
	}
L2221:
	;
	v8027 = *(*int32)(unsafe.Add(mBase, uint32(v8020)+192))
	v8034 = v8027
	goto L2220
L2222:
	;
	goto L2223
L2223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8020)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8020)+188)) = v8021
	v8032 = F_strchr(m, v8021, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8020)+192)) = v8032
	v8034 = v8032
	goto L2220
L2224:
	;
	if base.Ui32(v8024) <= base.Ui32(v8034) {
		v8058 = v8035
		goto L2217
	} else {
		goto L2225
	}
L2225:
	;
	v8039 = v8034
	v8041 = v8035
	goto L2226
L2226:
	;
	v8044 = int32(1)
	v8045 = v8041 + v8044
	*(*int32)(unsafe.Add(mBase, uint32(v8020)+196)) = v8045
	v8048 = v8039 + v8044
	*(*int32)(unsafe.Add(mBase, uint32(v8020)+188)) = v8048
	v8051 = F_strchr(m, v8048, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8020)+192)) = v8051
	if v8051 == int32(0) {
		v8058 = v8045
		goto L2217
	} else {
		goto L2228
	}
L2227:
	;
	v8058 = v8045
	goto L2217
L2228:
	;
	if base.Ui32(v8051) < base.Ui32(v8024) {
		v8039 = v8051
		v8041 = v8045
		goto L2226
	} else {
		goto L2229
	}
L2229:
	;
	goto L2227
L2230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8073))) = int32(21)
	v8078 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v8079 = *(*int32)(unsafe.Add(mBase, uint32(v8078)+520))
	v8081 = v8079 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8078)+520)) = v8081
	v8083 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8073)+33)) = uint8(v8083)
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+28)) = v8083
	*(*int64)(unsafe.Add(mBase, uint32(v8073)+20)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+8)) = v8081
	v8094 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8095 = m.ExcPending
	if v8095 != 0 {
		goto L66
	} else {
		goto L2252
	}
L2231:
	;
	v8460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))))
	if v8460 != int32(1) {
		goto L2347
	} else {
		goto L2348
	}
L2232:
	;
	v8457 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = uint8(v8457)
	goto L2231
L2233:
	;
	F_plpgsql_push_back_token(m, v8094, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8437 = m.ExcPending
	if v8437 != 0 {
		goto L66
	} else {
		goto L2345
	}
L2234:
	;
	F_plpgsql_push_back_token(m, int32(277), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8429 = m.ExcPending
	if v8429 != 0 {
		goto L66
	} else {
		goto L2344
	}
L2235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+20)) = int32(1)
	F_complete_direction(m, v8073, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8421 = m.ExcPending
	if v8421 != 0 {
		goto L66
	} else {
		goto L2343
	}
L2236:
	;
	v8382 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v8382 == int32(0) {
		goto L2234
	} else {
		goto L2333
	}
L2237:
	;
	F_complete_direction(m, v8073, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8381 = m.ExcPending
	if v8381 != 0 {
		goto L66
	} else {
		goto L2332
	}
L2238:
	;
	v8347 = int32(421678)
	v8350 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1418])))
	v8351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8313))))
	if v8351 == int32(0) {
		v8370 = v8350
		v8371 = v8351
		goto L2324
	} else {
		goto L2325
	}
L2239:
	;
	v8343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8073)+33)) = uint8(v8343)
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+24)) = int32(2147483647)
	goto L2231
L2240:
	;
	v8313 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v8313 == int32(0) {
		goto L2234
	} else {
		goto L2313
	}
L2241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+20)) = int32(3)
	v8299 = int32(0)
	v8302 = int32(1)
	v8310 = F_read_sql_construct(m, int32(324), int32(329), v8299, int32(530486), int32(2), v8302, v8302, v8299, v8299, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8311 = m.ExcPending
	if v8311 != 0 {
		goto L66
	} else {
		goto L2312
	}
L2242:
	;
	v8268 = int32(343001)
	v8271 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1419])))
	v8272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8105))))
	if v8272 == int32(0) {
		v8291 = v8271
		v8292 = v8272
		goto L2304
	} else {
		goto L2305
	}
L2243:
	;
	v8250 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+20)) = v8250
	v8254 = int32(0)
	v8257 = int32(1)
	v8265 = F_read_sql_construct(m, int32(324), int32(329), v8254, int32(530486), v8250, v8257, v8257, v8254, v8254, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8266 = m.ExcPending
	if v8266 != 0 {
		goto L66
	} else {
		goto L2302
	}
L2244:
	;
	v8223 = int32(348487)
	v8226 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1420])))
	v8227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8105))))
	if v8227 == int32(0) {
		v8246 = v8226
		v8247 = v8227
		goto L2294
	} else {
		goto L2295
	}
L2245:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8073)+20)) = int64(-4294967294)
	goto L2231
L2246:
	;
	v8194 = int32(78017)
	v8197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1421])))
	v8198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8105))))
	if v8198 == int32(0) {
		v8217 = v8197
		v8218 = v8198
		goto L2285
	} else {
		goto L2286
	}
L2247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+20)) = int32(2)
	goto L2231
L2248:
	;
	v8165 = int32(68001)
	v8168 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1422])))
	v8169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8105))))
	if v8169 == int32(0) {
		v8188 = v8168
		v8189 = v8169
		goto L2276
	} else {
		goto L2277
	}
L2249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+20)) = int32(1)
	goto L2231
L2250:
	;
	v8104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364]))))
	if v8104 != 0 {
		goto L2234
	} else {
		goto L2255
	}
L2251:
	;
	if v8094 != 0 {
		goto L2233
	} else {
		goto L2253
	}
L2252:
	;
	switch v8094 - int32(277) {
	case 0:
		goto L2250
	case 1, 2, 4, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 44, 45, 48, 49, 50, 51, 53, 54, 55, 56, 58, 59, 60, 61, 62, 63, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 81, 82:
		goto L2233
	case 3:
		goto L2243
	case 5:
		goto L2239
	case 9:
		goto L2235
	case 43:
		goto L2247
	case 46:
		goto L2237
	case 47, 52:
		goto L2232
	case 57:
		goto L2245
	case 64:
		goto L2231
	case 80:
		goto L2249
	case 83:
		goto L2241
	default:
		goto L2251
	}
L2253:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(250277))
	mBase = m.M
	v8103 = m.ExcPending
	if v8103 != 0 {
		goto L66
	} else {
		goto L2254
	}
L2254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2255:
	;
	v8105 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1371])))
	if v8105 == int32(0) {
		goto L2234
	} else {
		goto L2256
	}
L2256:
	;
	v8108 = int32(63501)
	v8111 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1399])))
	v8112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8105))))
	if v8112 == int32(0) {
		v8131 = v8111
		v8132 = v8112
		goto L2258
	} else {
		goto L2259
	}
L2257:
	;
	if v8132-v8131 == int32(0) {
		goto L2231
	} else {
		goto L2265
	}
L2258:
	;
	goto L2257
L2259:
	;
	if v8111 != v8112 {
		v8131 = v8111
		v8132 = v8112
		goto L2258
	} else {
		goto L2260
	}
L2260:
	;
	v8116 = v8105
	v8117 = v8108
	goto L2261
L2261:
	;
	v8120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8117)+1)))
	v8121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8116)+1)))
	if v8121 == int32(0) {
		v8131 = v8120
		v8132 = v8121
		goto L2258
	} else {
		goto L2263
	}
L2262:
	;
	v8131 = v8120
	v8132 = v8121
	goto L2258
L2263:
	;
	v8124 = int32(1)
	if v8120 == v8121 {
		v8116 = v8116 + v8124
		v8117 = v8117 + v8124
		goto L2261
	} else {
		goto L2264
	}
L2264:
	;
	goto L2262
L2265:
	;
	v8136 = int32(212835)
	v8139 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1423])))
	v8140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8105))))
	if v8140 == int32(0) {
		v8159 = v8139
		v8160 = v8140
		goto L2267
	} else {
		goto L2268
	}
L2266:
	;
	if v8160-v8159 != 0 {
		goto L2248
	} else {
		goto L2274
	}
L2267:
	;
	goto L2266
L2268:
	;
	if v8139 != v8140 {
		v8159 = v8139
		v8160 = v8140
		goto L2267
	} else {
		goto L2269
	}
L2269:
	;
	v8144 = v8105
	v8145 = v8136
	goto L2270
L2270:
	;
	v8148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8145)+1)))
	v8149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8144)+1)))
	if v8149 == int32(0) {
		v8159 = v8148
		v8160 = v8149
		goto L2267
	} else {
		goto L2272
	}
L2271:
	;
	v8159 = v8148
	v8160 = v8149
	goto L2267
L2272:
	;
	v8152 = int32(1)
	if v8148 == v8149 {
		v8144 = v8144 + v8152
		v8145 = v8145 + v8152
		goto L2270
	} else {
		goto L2273
	}
L2273:
	;
	goto L2271
L2274:
	;
	goto L2249
L2275:
	;
	if v8189-v8188 != 0 {
		goto L2246
	} else {
		goto L2283
	}
L2276:
	;
	goto L2275
L2277:
	;
	if v8168 != v8169 {
		v8188 = v8168
		v8189 = v8169
		goto L2276
	} else {
		goto L2278
	}
L2278:
	;
	v8173 = v8105
	v8174 = v8165
	goto L2279
L2279:
	;
	v8177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8174)+1)))
	v8178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8173)+1)))
	if v8178 == int32(0) {
		v8188 = v8177
		v8189 = v8178
		goto L2276
	} else {
		goto L2281
	}
L2280:
	;
	v8188 = v8177
	v8189 = v8178
	goto L2276
L2281:
	;
	v8181 = int32(1)
	if v8177 == v8178 {
		v8173 = v8173 + v8181
		v8174 = v8174 + v8181
		goto L2279
	} else {
		goto L2282
	}
L2282:
	;
	goto L2280
L2283:
	;
	goto L2247
L2284:
	;
	if v8218-v8217 != 0 {
		goto L2244
	} else {
		goto L2292
	}
L2285:
	;
	goto L2284
L2286:
	;
	if v8197 != v8198 {
		v8217 = v8197
		v8218 = v8198
		goto L2285
	} else {
		goto L2287
	}
L2287:
	;
	v8202 = v8105
	v8203 = v8194
	goto L2288
L2288:
	;
	v8206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8203)+1)))
	v8207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8202)+1)))
	if v8207 == int32(0) {
		v8217 = v8206
		v8218 = v8207
		goto L2285
	} else {
		goto L2290
	}
L2289:
	;
	v8217 = v8206
	v8218 = v8207
	goto L2285
L2290:
	;
	v8210 = int32(1)
	if v8206 == v8207 {
		v8202 = v8202 + v8210
		v8203 = v8203 + v8210
		goto L2288
	} else {
		goto L2291
	}
L2291:
	;
	goto L2289
L2292:
	;
	goto L2245
L2293:
	;
	if v8247-v8246 != 0 {
		goto L2242
	} else {
		goto L2301
	}
L2294:
	;
	goto L2293
L2295:
	;
	if v8226 != v8227 {
		v8246 = v8226
		v8247 = v8227
		goto L2294
	} else {
		goto L2296
	}
L2296:
	;
	v8231 = v8105
	v8232 = v8223
	goto L2297
L2297:
	;
	v8235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+1)))
	v8236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8231)+1)))
	if v8236 == int32(0) {
		v8246 = v8235
		v8247 = v8236
		goto L2294
	} else {
		goto L2299
	}
L2298:
	;
	v8246 = v8235
	v8247 = v8236
	goto L2294
L2299:
	;
	v8239 = int32(1)
	if v8235 == v8236 {
		v8231 = v8231 + v8239
		v8232 = v8232 + v8239
		goto L2297
	} else {
		goto L2300
	}
L2300:
	;
	goto L2298
L2301:
	;
	goto L2243
L2302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+28)) = v8265
	goto L2232
L2303:
	;
	if v8292-v8291 != 0 {
		goto L2240
	} else {
		goto L2311
	}
L2304:
	;
	goto L2303
L2305:
	;
	if v8271 != v8272 {
		v8291 = v8271
		v8292 = v8272
		goto L2304
	} else {
		goto L2306
	}
L2306:
	;
	v8276 = v8105
	v8277 = v8268
	goto L2307
L2307:
	;
	v8280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8277)+1)))
	v8281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8276)+1)))
	if v8281 == int32(0) {
		v8291 = v8280
		v8292 = v8281
		goto L2304
	} else {
		goto L2309
	}
L2308:
	;
	v8291 = v8280
	v8292 = v8281
	goto L2304
L2309:
	;
	v8284 = int32(1)
	if v8280 == v8281 {
		v8276 = v8276 + v8284
		v8277 = v8277 + v8284
		goto L2307
	} else {
		goto L2310
	}
L2310:
	;
	goto L2308
L2311:
	;
	goto L2241
L2312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+28)) = v8310
	goto L2232
L2313:
	;
	v8316 = int32(305221)
	v8319 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1245])))
	v8320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8313))))
	if v8320 == int32(0) {
		v8339 = v8319
		v8340 = v8320
		goto L2315
	} else {
		goto L2316
	}
L2314:
	;
	if v8340-v8339 != 0 {
		goto L2238
	} else {
		goto L2322
	}
L2315:
	;
	goto L2314
L2316:
	;
	if v8319 != v8320 {
		v8339 = v8319
		v8340 = v8320
		goto L2315
	} else {
		goto L2317
	}
L2317:
	;
	v8324 = v8313
	v8325 = v8316
	goto L2318
L2318:
	;
	v8328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8325)+1)))
	v8329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8324)+1)))
	if v8329 == int32(0) {
		v8339 = v8328
		v8340 = v8329
		goto L2315
	} else {
		goto L2320
	}
L2319:
	;
	v8339 = v8328
	v8340 = v8329
	goto L2315
L2320:
	;
	v8332 = int32(1)
	if v8328 == v8329 {
		v8324 = v8324 + v8332
		v8325 = v8325 + v8332
		goto L2318
	} else {
		goto L2321
	}
L2321:
	;
	goto L2319
L2322:
	;
	goto L2239
L2323:
	;
	if v8371-v8370 != 0 {
		goto L2236
	} else {
		goto L2331
	}
L2324:
	;
	goto L2323
L2325:
	;
	if v8350 != v8351 {
		v8370 = v8350
		v8371 = v8351
		goto L2324
	} else {
		goto L2326
	}
L2326:
	;
	v8355 = v8313
	v8356 = v8347
	goto L2327
L2327:
	;
	v8359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8356)+1)))
	v8360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8355)+1)))
	if v8360 == int32(0) {
		v8370 = v8359
		v8371 = v8360
		goto L2324
	} else {
		goto L2329
	}
L2328:
	;
	v8370 = v8359
	v8371 = v8360
	goto L2324
L2329:
	;
	v8363 = int32(1)
	if v8359 == v8360 {
		v8355 = v8355 + v8363
		v8356 = v8356 + v8363
		goto L2327
	} else {
		goto L2330
	}
L2330:
	;
	goto L2328
L2331:
	;
	goto L2237
L2332:
	;
	goto L2231
L2333:
	;
	v8385 = int32(421768)
	v8388 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1424])))
	v8389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8382))))
	if v8389 == int32(0) {
		v8408 = v8388
		v8409 = v8389
		goto L2335
	} else {
		goto L2336
	}
L2334:
	;
	if v8409-v8408 != 0 {
		goto L2234
	} else {
		goto L2342
	}
L2335:
	;
	goto L2334
L2336:
	;
	if v8388 != v8389 {
		v8408 = v8388
		v8409 = v8389
		goto L2335
	} else {
		goto L2337
	}
L2337:
	;
	v8393 = v8382
	v8394 = v8385
	goto L2338
L2338:
	;
	v8397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8394)+1)))
	v8398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8393)+1)))
	if v8398 == int32(0) {
		v8408 = v8397
		v8409 = v8398
		goto L2335
	} else {
		goto L2340
	}
L2339:
	;
	v8408 = v8397
	v8409 = v8398
	goto L2335
L2340:
	;
	v8401 = int32(1)
	if v8397 == v8398 {
		v8393 = v8393 + v8401
		v8394 = v8394 + v8401
		goto L2338
	} else {
		goto L2341
	}
L2341:
	;
	goto L2339
L2342:
	;
	goto L2235
L2343:
	;
	goto L2231
L2344:
	;
	v8430 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = uint8(v8430)
	goto L2231
L2345:
	;
	v8440 = int32(0)
	v8443 = int32(1)
	v8451 = F_read_sql_construct(m, int32(324), int32(329), v8440, int32(530486), int32(2), v8443, v8443, v8440, v8440, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8452 = m.ExcPending
	if v8452 != 0 {
		goto L66
	} else {
		goto L2346
	}
L2346:
	;
	v8453 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8073)+33)) = uint8(v8453)
	*(*int32)(unsafe.Add(mBase, uint32(v8073)+28)) = v8451
	goto L2232
L2347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8073
	v10149 = v241
	goto L5
L2348:
	;
	v8467 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8468 = m.ExcPending
	if v8468 != 0 {
		goto L66
	} else {
		goto L2350
	}
L2349:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(530477))
	mBase = m.M
	v8476 = m.ExcPending
	if v8476 != 0 {
		goto L66
	} else {
		goto L2351
	}
L2350:
	;
	switch v8467 - int32(324) {
	case 0, 5:
		goto L2347
	default:
		goto L2349
	}
L2351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8479))) = int32(22)
	v8485 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v8486 = int32(0)
	if v8485 < v8486 {
		v8529 = v8486
		goto L2354
	} else {
		goto L2355
	}
L2353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8479)+4)) = v8529
	v8534 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v8535 = *(*int32)(unsafe.Add(mBase, uint32(v8534)+520))
	v8537 = v8535 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8534)+520)) = v8537
	*(*int32)(unsafe.Add(mBase, uint32(v8479)+8)) = v8537
	v8542 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8543 = *(*int32)(unsafe.Add(mBase, uint32(v8542)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8479)+12)) = v8543
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8479
	v10149 = v241
	goto L5
L2354:
	;
	goto L2353
L2355:
	;
	v8491 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8492 = *(*int32)(unsafe.Add(mBase, uint32(v8491)+60))
	if v8492 == int32(0) {
		v8529 = v8486
		goto L2354
	} else {
		goto L2356
	}
L2356:
	;
	v8495 = v8485 + v8492
	v8496 = *(*int32)(unsafe.Add(mBase, uint32(v8491)+188))
	if base.Ui32(v8496) <= base.Ui32(v8495) {
		goto L2358
	} else {
		goto L2359
	}
L2357:
	;
	v8506 = *(*int32)(unsafe.Add(mBase, uint32(v8491)+196))
	if v8505 == int32(0) {
		v8529 = v8506
		goto L2354
	} else {
		goto L2361
	}
L2358:
	;
	v8498 = *(*int32)(unsafe.Add(mBase, uint32(v8491)+192))
	v8505 = v8498
	goto L2357
L2359:
	;
	goto L2360
L2360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8491)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8491)+188)) = v8492
	v8503 = F_strchr(m, v8492, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8491)+192)) = v8503
	v8505 = v8503
	goto L2357
L2361:
	;
	if base.Ui32(v8495) <= base.Ui32(v8505) {
		v8529 = v8506
		goto L2354
	} else {
		goto L2362
	}
L2362:
	;
	v8510 = v8505
	v8512 = v8506
	goto L2363
L2363:
	;
	v8515 = int32(1)
	v8516 = v8512 + v8515
	*(*int32)(unsafe.Add(mBase, uint32(v8491)+196)) = v8516
	v8519 = v8510 + v8515
	*(*int32)(unsafe.Add(mBase, uint32(v8491)+188)) = v8519
	v8522 = F_strchr(m, v8519, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8491)+192)) = v8522
	if v8522 == int32(0) {
		v8529 = v8516
		goto L2354
	} else {
		goto L2365
	}
L2364:
	;
	v8529 = v8516
	goto L2354
L2365:
	;
	if base.Ui32(v8522) < base.Ui32(v8495) {
		v8510 = v8522
		v8512 = v8516
		goto L2363
	} else {
		goto L2366
	}
L2366:
	;
	goto L2364
L2367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8549))) = int32(25)
	v8555 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v8556 = int32(0)
	if v8555 < v8556 {
		v8599 = v8556
		goto L2369
	} else {
		goto L2370
	}
L2368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8549)+4)) = v8599
	v8604 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v8605 = *(*int32)(unsafe.Add(mBase, uint32(v8604)+520))
	v8607 = v8605 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8604)+520)) = v8607
	*(*int32)(unsafe.Add(mBase, uint32(v8549)+8)) = v8607
	v8612 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8549)+12)) = uint8(base.B2i32(v8612 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8549
	v10149 = v241
	goto L5
L2369:
	;
	goto L2368
L2370:
	;
	v8561 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8562 = *(*int32)(unsafe.Add(mBase, uint32(v8561)+60))
	if v8562 == int32(0) {
		v8599 = v8556
		goto L2369
	} else {
		goto L2371
	}
L2371:
	;
	v8565 = v8555 + v8562
	v8566 = *(*int32)(unsafe.Add(mBase, uint32(v8561)+188))
	if base.Ui32(v8566) <= base.Ui32(v8565) {
		goto L2373
	} else {
		goto L2374
	}
L2372:
	;
	v8576 = *(*int32)(unsafe.Add(mBase, uint32(v8561)+196))
	if v8575 == int32(0) {
		v8599 = v8576
		goto L2369
	} else {
		goto L2376
	}
L2373:
	;
	v8568 = *(*int32)(unsafe.Add(mBase, uint32(v8561)+192))
	v8575 = v8568
	goto L2372
L2374:
	;
	goto L2375
L2375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8561)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8561)+188)) = v8562
	v8573 = F_strchr(m, v8562, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8561)+192)) = v8573
	v8575 = v8573
	goto L2372
L2376:
	;
	if base.Ui32(v8565) <= base.Ui32(v8575) {
		v8599 = v8576
		goto L2369
	} else {
		goto L2377
	}
L2377:
	;
	v8580 = v8575
	v8582 = v8576
	goto L2378
L2378:
	;
	v8585 = int32(1)
	v8586 = v8582 + v8585
	*(*int32)(unsafe.Add(mBase, uint32(v8561)+196)) = v8586
	v8589 = v8580 + v8585
	*(*int32)(unsafe.Add(mBase, uint32(v8561)+188)) = v8589
	v8592 = F_strchr(m, v8589, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8561)+192)) = v8592
	if v8592 == int32(0) {
		v8599 = v8586
		goto L2369
	} else {
		goto L2380
	}
L2379:
	;
	v8599 = v8586
	goto L2369
L2380:
	;
	if base.Ui32(v8592) < base.Ui32(v8565) {
		v8580 = v8592
		v8582 = v8586
		goto L2378
	} else {
		goto L2381
	}
L2381:
	;
	goto L2379
L2382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8618))) = int32(26)
	v8624 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v8625 = int32(0)
	if v8624 < v8625 {
		v8668 = v8625
		goto L2384
	} else {
		goto L2385
	}
L2383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8618)+4)) = v8668
	v8673 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v8674 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+520))
	v8676 = v8674 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8673)+520)) = v8676
	*(*int32)(unsafe.Add(mBase, uint32(v8618)+8)) = v8676
	v8681 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8618)+12)) = uint8(base.B2i32(v8681 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8618
	v10149 = v241
	goto L5
L2384:
	;
	goto L2383
L2385:
	;
	v8630 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8631 = *(*int32)(unsafe.Add(mBase, uint32(v8630)+60))
	if v8631 == int32(0) {
		v8668 = v8625
		goto L2384
	} else {
		goto L2386
	}
L2386:
	;
	v8634 = v8624 + v8631
	v8635 = *(*int32)(unsafe.Add(mBase, uint32(v8630)+188))
	if base.Ui32(v8635) <= base.Ui32(v8634) {
		goto L2388
	} else {
		goto L2389
	}
L2387:
	;
	v8645 = *(*int32)(unsafe.Add(mBase, uint32(v8630)+196))
	if v8644 == int32(0) {
		v8668 = v8645
		goto L2384
	} else {
		goto L2391
	}
L2388:
	;
	v8637 = *(*int32)(unsafe.Add(mBase, uint32(v8630)+192))
	v8644 = v8637
	goto L2387
L2389:
	;
	goto L2390
L2390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8630)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8630)+188)) = v8631
	v8642 = F_strchr(m, v8631, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8630)+192)) = v8642
	v8644 = v8642
	goto L2387
L2391:
	;
	if base.Ui32(v8634) <= base.Ui32(v8644) {
		v8668 = v8645
		goto L2384
	} else {
		goto L2392
	}
L2392:
	;
	v8649 = v8644
	v8651 = v8645
	goto L2393
L2393:
	;
	v8654 = int32(1)
	v8655 = v8651 + v8654
	*(*int32)(unsafe.Add(mBase, uint32(v8630)+196)) = v8655
	v8658 = v8649 + v8654
	*(*int32)(unsafe.Add(mBase, uint32(v8630)+188)) = v8658
	v8661 = F_strchr(m, v8658, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8630)+192)) = v8661
	if v8661 == int32(0) {
		v8668 = v8655
		goto L2384
	} else {
		goto L2395
	}
L2394:
	;
	v8668 = v8655
	goto L2384
L2395:
	;
	if base.Ui32(v8661) < base.Ui32(v8634) {
		v8649 = v8661
		v8651 = v8655
		goto L2393
	} else {
		goto L2396
	}
L2396:
	;
	goto L2394
L2397:
	;
	v8694 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v8695 = m.ExcPending
	if v8695 != 0 {
		goto L66
	} else {
		goto L2398
	}
L2398:
	;
	if v8694 == int32(91) {
		goto L18
	} else {
		goto L2399
	}
L2399:
	;
	v8698 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v8699 = *(*int32)(unsafe.Add(mBase, uint32(v8698)+24))
	v8700 = *(*int32)(unsafe.Add(mBase, uint32(v8699)+4))
	if v8700 != int32(1790) {
		goto L17
	} else {
		goto L2400
	}
L2400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8698
	v10149 = v241
	goto L5
L2401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2403:
	;
	v8760 = F_palloc(m, int32(12))
	mBase = m.M
	v8761 = m.ExcPending
	if v8761 != 0 {
		goto L66
	} else {
		goto L2417
	}
L2404:
	;
	goto L2403
L2405:
	;
	v8718 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8719 = *(*int32)(unsafe.Add(mBase, uint32(v8718)+60))
	if v8719 == int32(0) {
		v8756 = v8713
		goto L2404
	} else {
		goto L2406
	}
L2406:
	;
	v8722 = v8712 + v8719
	v8723 = *(*int32)(unsafe.Add(mBase, uint32(v8718)+188))
	if base.Ui32(v8723) <= base.Ui32(v8722) {
		goto L2408
	} else {
		goto L2409
	}
L2407:
	;
	v8733 = *(*int32)(unsafe.Add(mBase, uint32(v8718)+196))
	if v8732 == int32(0) {
		v8756 = v8733
		goto L2404
	} else {
		goto L2411
	}
L2408:
	;
	v8725 = *(*int32)(unsafe.Add(mBase, uint32(v8718)+192))
	v8732 = v8725
	goto L2407
L2409:
	;
	goto L2410
L2410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8718)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8718)+188)) = v8719
	v8730 = F_strchr(m, v8719, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8718)+192)) = v8730
	v8732 = v8730
	goto L2407
L2411:
	;
	if base.Ui32(v8722) <= base.Ui32(v8732) {
		v8756 = v8733
		goto L2404
	} else {
		goto L2412
	}
L2412:
	;
	v8737 = v8732
	v8739 = v8733
	goto L2413
L2413:
	;
	v8742 = int32(1)
	v8743 = v8739 + v8742
	*(*int32)(unsafe.Add(mBase, uint32(v8718)+196)) = v8743
	v8746 = v8737 + v8742
	*(*int32)(unsafe.Add(mBase, uint32(v8718)+188)) = v8746
	v8749 = F_strchr(m, v8746, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8718)+192)) = v8749
	if v8749 == int32(0) {
		v8756 = v8743
		goto L2404
	} else {
		goto L2415
	}
L2414:
	;
	v8756 = v8743
	goto L2404
L2415:
	;
	if base.Ui32(v8749) < base.Ui32(v8722) {
		v8737 = v8749
		v8739 = v8743
		goto L2413
	} else {
		goto L2416
	}
L2416:
	;
	goto L2414
L2417:
	;
	v8763 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v8764 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8763)+525)) = uint8(v8764)
	v8769 = *(*int32)(unsafe.Add(mBase, uint32(v8763)+44))
	v8771 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8769, int32(0))
	mBase = m.M
	v8772 = m.ExcPending
	if v8772 != 0 {
		goto L66
	} else {
		goto L2418
	}
L2418:
	;
	v8774 = F_plpgsql_build_variable(m, int32(351311), v8756, v8771, int32(1))
	mBase = m.M
	v8775 = m.ExcPending
	if v8775 != 0 {
		goto L66
	} else {
		goto L2419
	}
L2419:
	;
	v8776 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8774)+16)) = uint8(v8776)
	v8778 = *(*int32)(unsafe.Add(mBase, uint32(v8774)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8760))) = v8778
	v8784 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v8785 = *(*int32)(unsafe.Add(mBase, uint32(v8784)+44))
	v8787 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8785, int32(0))
	mBase = m.M
	v8788 = m.ExcPending
	if v8788 != 0 {
		goto L66
	} else {
		goto L2420
	}
L2420:
	;
	v8790 = F_plpgsql_build_variable(m, int32(287602), v8756, v8787, int32(1))
	mBase = m.M
	v8791 = m.ExcPending
	if v8791 != 0 {
		goto L66
	} else {
		goto L2421
	}
L2421:
	;
	v8792 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8790)+16)) = uint8(v8792)
	v8794 = *(*int32)(unsafe.Add(mBase, uint32(v8790)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8760)+4)) = v8794
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8760
	v10149 = v241
	goto L5
L2422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8807
	v10149 = v241
	goto L5
L2423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8816
	v10149 = v241
	goto L5
L2424:
	;
	v8824 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v8825 = int32(0)
	if v8824 < v8825 {
		v8868 = v8825
		goto L2426
	} else {
		goto L2427
	}
L2425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8820))) = v8868
	v8874 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8820)+4)) = v8874
	v8876 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v8820)+8)) = v8876
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8820
	v10149 = v241
	goto L5
L2426:
	;
	goto L2425
L2427:
	;
	v8830 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8831 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+60))
	if v8831 == int32(0) {
		v8868 = v8825
		goto L2426
	} else {
		goto L2428
	}
L2428:
	;
	v8834 = v8824 + v8831
	v8835 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+188))
	if base.Ui32(v8835) <= base.Ui32(v8834) {
		goto L2430
	} else {
		goto L2431
	}
L2429:
	;
	v8845 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+196))
	if v8844 == int32(0) {
		v8868 = v8845
		goto L2426
	} else {
		goto L2433
	}
L2430:
	;
	v8837 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+192))
	v8844 = v8837
	goto L2429
L2431:
	;
	goto L2432
L2432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+188)) = v8831
	v8842 = F_strchr(m, v8831, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+192)) = v8842
	v8844 = v8842
	goto L2429
L2433:
	;
	if base.Ui32(v8834) <= base.Ui32(v8844) {
		v8868 = v8845
		goto L2426
	} else {
		goto L2434
	}
L2434:
	;
	v8849 = v8844
	v8851 = v8845
	goto L2435
L2435:
	;
	v8854 = int32(1)
	v8855 = v8851 + v8854
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+196)) = v8855
	v8858 = v8849 + v8854
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+188)) = v8858
	v8861 = F_strchr(m, v8858, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+192)) = v8861
	if v8861 == int32(0) {
		v8868 = v8855
		goto L2426
	} else {
		goto L2437
	}
L2436:
	;
	v8868 = v8855
	goto L2426
L2437:
	;
	if base.Ui32(v8861) < base.Ui32(v8834) {
		v8849 = v8861
		v8851 = v8855
		goto L2435
	} else {
		goto L2438
	}
L2438:
	;
	goto L2436
L2439:
	;
	v8906 = *(*int32)(unsafe.Add(mBase, uint32(v8885)+8))
	if v8906 != 0 {
		v8885 = v8906
		goto L2439
	} else {
		goto L2441
	}
L2440:
	;
	v8907 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v8885)+8)) = v8907
	v8909 = *(*int32)(unsafe.Add(mBase, uint32(v8880)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8909
	v10149 = v241
	goto L5
L2441:
	;
	goto L2440
L2442:
	;
	if v8938-v8937 != 0 {
		goto L2450
	} else {
		goto L2451
	}
L2443:
	;
	goto L2442
L2444:
	;
	if v8917 != v8918 {
		v8937 = v8917
		v8938 = v8918
		goto L2443
	} else {
		goto L2445
	}
L2445:
	;
	v8922 = v8913
	v8923 = v8914
	goto L2446
L2446:
	;
	v8926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8923)+1)))
	v8927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8922)+1)))
	if v8927 == int32(0) {
		v8937 = v8926
		v8938 = v8927
		goto L2443
	} else {
		goto L2448
	}
L2447:
	;
	v8937 = v8926
	v8938 = v8927
	goto L2443
L2448:
	;
	v8930 = int32(1)
	if v8926 == v8927 {
		v8922 = v8922 + v8930
		v8923 = v8923 + v8930
		goto L2446
	} else {
		goto L2449
	}
L2449:
	;
	goto L2447
L2450:
	;
	v8940 = F_plpgsql_parse_err_condition(m, v8913)
	mBase = m.M
	v8941 = m.ExcPending
	if v8941 != 0 {
		goto L66
	} else {
		goto L2453
	}
L2451:
	;
	goto L2452
L2452:
	;
	v8947 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8948 = m.ExcPending
	if v8948 != 0 {
		goto L66
	} else {
		goto L2454
	}
L2453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8940
	v10149 = v241
	goto L5
L2454:
	;
	if v8947 != int32(261) {
		goto L16
	} else {
		goto L2455
	}
L2455:
	;
	v8951 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	v8952 = F_strlen(m, v8951)
	mBase = m.M
	if v8952 != int32(5) {
		goto L15
	} else {
		goto L2456
	}
L2456:
	;
	v8955 = int32(507831)
	v8959 = m.G0
	v8961 = v8959 - int32(32)
	v8962 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8961)+24)) = v8962
	*(*int64)(unsafe.Add(mBase, uint32(v8961)+16)) = v8962
	*(*int64)(unsafe.Add(mBase, uint32(v8961)+8)) = v8962
	*(*int64)(unsafe.Add(mBase, uint32(v8961))) = v8962
	v8970 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1406])))
	if v8970 == int32(0) {
		goto L2458
	} else {
		goto L2459
	}
L2457:
	;
	if v9038 != int32(5) {
		goto L14
	} else {
		goto L2478
	}
L2458:
	;
	v9038 = int32(0)
	goto L2457
L2459:
	;
	goto L2460
L2460:
	;
	v8974 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1407])))
	if v8974 == int32(0) {
		goto L2461
	} else {
		goto L2462
	}
L2461:
	;
	v8978 = v8951
	goto L2464
L2462:
	;
	goto L2463
L2463:
	;
	v8988 = v8955
	v8989 = v8970
	goto L2467
L2464:
	;
	v8984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8978))))
	if v8984 == v8970 {
		v8978 = v8978 + int32(1)
		goto L2464
	} else {
		goto L2466
	}
L2465:
	;
	v9038 = v8978 - v8951
	goto L2457
L2466:
	;
	goto L2465
L2467:
	;
	v8996 = v8961 + int32(base.Ui32(v8989)>>(uint(int32(3))%32))&int32(28)
	v8997 = *(*int32)(unsafe.Add(mBase, uint32(v8996)))
	v8998 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8996))) = v8997 | v8998<<(uint(v8989)%32)
	v9002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8988)+1)))
	if v9002 != 0 {
		v8988 = v8988 + v8998
		v8989 = v9002
		goto L2467
	} else {
		goto L2469
	}
L2468:
	;
	v9005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8951))))
	if v9005 == int32(0) {
		v9030 = v8951
		goto L2470
	} else {
		goto L2471
	}
L2469:
	;
	goto L2468
L2470:
	;
	v9038 = v9030 - v8951
	goto L2457
L2471:
	;
	v9009 = v8951
	v9010 = v9005
	goto L2472
L2472:
	;
	v9018 = *(*int32)(unsafe.Add(mBase, uint32(v8961+int32(base.Ui32(v9010)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v9018)>>(uint(v9010)%32))&int32(1) == int32(0) {
		goto L2474
	} else {
		goto L2475
	}
L2473:
	;
	v9030 = v9026
	goto L2470
L2474:
	;
	v9030 = v9009
	goto L2470
L2475:
	;
	goto L2476
L2476:
	;
	v9024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9009)+1)))
	v9026 = v9009 + int32(1)
	if v9024 != 0 {
		v9009 = v9026
		v9010 = v9024
		goto L2472
	} else {
		goto L2477
	}
L2477:
	;
	goto L2473
L2478:
	;
	v9042 = F_palloc(m, int32(12))
	mBase = m.M
	v9043 = m.ExcPending
	if v9043 != 0 {
		goto L66
	} else {
		goto L2479
	}
L2479:
	;
	v9044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8951)+4)))
	v9045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8951)+3)))
	v9046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8951)+2)))
	v9047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8951)+1)))
	v9048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8951))))
	*(*int32)(unsafe.Add(mBase, uint32(v9042)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9042)+4)) = v8951
	v9052 = int32(16)
	v9054 = int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v9042))) = (v9048+v9052)&v9054 | (v9047+v9052)&v9054<<(uint(int32(6))%32) | (v9046+v9052)&v9054<<(uint(int32(12))%32) | (v9045+v9052)&v9054<<(uint(int32(18))%32) | (v9044+v9052)&v9054<<(uint(int32(24))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9042
	v10149 = v241
	goto L5
L2480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9099
	v10149 = v241
	goto L5
L2481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9115
	v10149 = v241
	goto L5
L2482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9131
	v10149 = v241
	goto L5
L2483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L2484:
	;
	v9146 = *(*int32)(unsafe.Add(mBase, uint32(v9141)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9146
	v10149 = v241
	goto L5
L2485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10149 = v241
	goto L5
L2486:
	;
	v9160 = *(*int32)(unsafe.Add(mBase, uint32(v9155)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9160
	v10149 = v241
	goto L5
L2487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9173
	v10149 = v241
	goto L5
L2488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9176
	v10149 = v241
	goto L5
L2489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2490:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2491:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9202 = m.ExcPending
	if v9202 != 0 {
		goto L66
	} else {
		goto L2492
	}
L2492:
	;
	v9203 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	v9204 = *(*int32)(unsafe.Add(mBase, uint32(v9203)+4))
	v9205 = F_format_type_be(m, v9204)
	mBase = m.M
	v9206 = m.ExcPending
	if v9206 != 0 {
		goto L66
	} else {
		goto L2493
	}
L2493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v9205
	F_errmsg(m, int32(188243), v27+int32(32))
	mBase = m.M
	v9212 = m.ExcPending
	if v9212 != 0 {
		goto L66
	} else {
		goto L2494
	}
L2494:
	;
	v9215 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v9216 = F_plpgsql_scanner_errposition(m, v9215, l1)
	mBase = m.M
	v9217 = m.ExcPending
	if v9217 != 0 {
		goto L66
	} else {
		goto L2495
	}
L2495:
	;
	F_errfinish(m, int32(26959), int32(523), int32(360630))
	mBase = m.M
	v9222 = m.ExcPending
	if v9222 != 0 {
		goto L66
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v9229 = m.ExcPending
	if v9229 != 0 {
		goto L66
	} else {
		goto L2498
	}
L2498:
	;
	v9230 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v9230
	F_errmsg(m, int32(532422), v27+int32(16))
	mBase = m.M
	v9236 = m.ExcPending
	if v9236 != 0 {
		goto L66
	} else {
		goto L2499
	}
L2499:
	;
	v9239 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9240 = F_plpgsql_scanner_errposition(m, v9239, l1)
	mBase = m.M
	v9241 = m.ExcPending
	if v9241 != 0 {
		goto L66
	} else {
		goto L2500
	}
L2500:
	;
	F_errfinish(m, int32(26959), int32(542), int32(360630))
	mBase = m.M
	v9246 = m.ExcPending
	if v9246 != 0 {
		goto L66
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
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9253 = m.ExcPending
	if v9253 != 0 {
		goto L66
	} else {
		goto L2503
	}
L2503:
	;
	v9254 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v9254
	F_errmsg(m, int32(72478), v27+int32(48))
	mBase = m.M
	v9260 = m.ExcPending
	if v9260 != 0 {
		goto L66
	} else {
		goto L2504
	}
L2504:
	;
	v9261 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9262 = F_plpgsql_scanner_errposition(m, v9261, l1)
	mBase = m.M
	v9263 = m.ExcPending
	if v9263 != 0 {
		goto L66
	} else {
		goto L2505
	}
L2505:
	;
	F_errfinish(m, int32(26959), int32(667), int32(360630))
	mBase = m.M
	v9268 = m.ExcPending
	if v9268 != 0 {
		goto L66
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
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9275 = m.ExcPending
	if v9275 != 0 {
		goto L66
	} else {
		goto L2508
	}
L2508:
	;
	v9276 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v9276
	F_errmsg(m, int32(72478), v27-int32(-64))
	mBase = m.M
	v9282 = m.ExcPending
	if v9282 != 0 {
		goto L66
	} else {
		goto L2509
	}
L2509:
	;
	v9283 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9284 = F_plpgsql_scanner_errposition(m, v9283, l1)
	mBase = m.M
	v9285 = m.ExcPending
	if v9285 != 0 {
		goto L66
	} else {
		goto L2510
	}
L2510:
	;
	F_errfinish(m, int32(26959), int32(682), int32(360630))
	mBase = m.M
	v9290 = m.ExcPending
	if v9290 != 0 {
		goto L66
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
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9298 = m.ExcPending
	if v9298 != 0 {
		goto L66
	} else {
		goto L2513
	}
L2513:
	;
	v9299 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v9300 = F_NameListToString(m, v9299)
	mBase = m.M
	v9301 = m.ExcPending
	if v9301 != 0 {
		goto L66
	} else {
		goto L2514
	}
L2514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v9300
	F_errmsg(m, int32(72478), v27+int32(80))
	mBase = m.M
	v9307 = m.ExcPending
	if v9307 != 0 {
		goto L66
	} else {
		goto L2515
	}
L2515:
	;
	v9308 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9309 = F_plpgsql_scanner_errposition(m, v9308, l1)
	mBase = m.M
	v9310 = m.ExcPending
	if v9310 != 0 {
		goto L66
	} else {
		goto L2516
	}
L2516:
	;
	F_errfinish(m, int32(26959), int32(708), int32(360630))
	mBase = m.M
	v9315 = m.ExcPending
	if v9315 != 0 {
		goto L66
	} else {
		goto L2517
	}
L2517:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2520:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2521:
	;
	F_errmsg_internal(m, int32(163453), int32(0))
	mBase = m.M
	v9342 = m.ExcPending
	if v9342 != 0 {
		goto L66
	} else {
		goto L2522
	}
L2522:
	;
	F_errfinish(m, int32(26959), int32(990), int32(360630))
	mBase = m.M
	v9347 = m.ExcPending
	if v9347 != 0 {
		goto L66
	} else {
		goto L2523
	}
L2523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2524:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9354 = m.ExcPending
	if v9354 != 0 {
		goto L66
	} else {
		goto L2525
	}
L2525:
	;
	v9355 = F_NameOfDatum(m, v132)
	mBase = m.M
	v9356 = m.ExcPending
	if v9356 != 0 {
		goto L66
	} else {
		goto L2526
	}
L2526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+176)) = v9355
	F_errmsg(m, int32(396432), v27+int32(176))
	mBase = m.M
	v9362 = m.ExcPending
	if v9362 != 0 {
		goto L66
	} else {
		goto L2527
	}
L2527:
	;
	v9363 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9364 = F_plpgsql_scanner_errposition(m, v9363, l1)
	mBase = m.M
	v9365 = m.ExcPending
	if v9365 != 0 {
		goto L66
	} else {
		goto L2528
	}
L2528:
	;
	F_errfinish(m, int32(26959), int32(1174), int32(360630))
	mBase = m.M
	v9370 = m.ExcPending
	if v9370 != 0 {
		goto L66
	} else {
		goto L2529
	}
L2529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2530:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9377 = m.ExcPending
	if v9377 != 0 {
		goto L66
	} else {
		goto L2531
	}
L2531:
	;
	F_errmsg(m, int32(166699), int32(0))
	mBase = m.M
	v9381 = m.ExcPending
	if v9381 != 0 {
		goto L66
	} else {
		goto L2532
	}
L2532:
	;
	v9384 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9385 = F_plpgsql_scanner_errposition(m, v9384, l1)
	mBase = m.M
	v9386 = m.ExcPending
	if v9386 != 0 {
		goto L66
	} else {
		goto L2533
	}
L2533:
	;
	F_errfinish(m, int32(26959), int32(1403), int32(360630))
	mBase = m.M
	v9391 = m.ExcPending
	if v9391 != 0 {
		goto L66
	} else {
		goto L2534
	}
L2534:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2535:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9398 = m.ExcPending
	if v9398 != 0 {
		goto L66
	} else {
		goto L2536
	}
L2536:
	;
	F_errmsg(m, int32(396210), int32(0))
	mBase = m.M
	v9402 = m.ExcPending
	if v9402 != 0 {
		goto L66
	} else {
		goto L2537
	}
L2537:
	;
	v9405 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9406 = F_plpgsql_scanner_errposition(m, v9405, l1)
	mBase = m.M
	v9407 = m.ExcPending
	if v9407 != 0 {
		goto L66
	} else {
		goto L2538
	}
L2538:
	;
	F_errfinish(m, int32(26959), int32(1438), int32(360630))
	mBase = m.M
	v9412 = m.ExcPending
	if v9412 != 0 {
		goto L66
	} else {
		goto L2539
	}
L2539:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2540:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9419 = m.ExcPending
	if v9419 != 0 {
		goto L66
	} else {
		goto L2541
	}
L2541:
	;
	F_errmsg(m, int32(396358), int32(0))
	mBase = m.M
	v9423 = m.ExcPending
	if v9423 != 0 {
		goto L66
	} else {
		goto L2542
	}
L2542:
	;
	v9424 = F_plpgsql_scanner_errposition(m, v4801, l1)
	mBase = m.M
	v9425 = m.ExcPending
	if v9425 != 0 {
		goto L66
	} else {
		goto L2543
	}
L2543:
	;
	F_errfinish(m, int32(26959), int32(1445), int32(360630))
	mBase = m.M
	v9430 = m.ExcPending
	if v9430 != 0 {
		goto L66
	} else {
		goto L2544
	}
L2544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2545:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9437 = m.ExcPending
	if v9437 != 0 {
		goto L66
	} else {
		goto L2546
	}
L2546:
	;
	F_errmsg(m, int32(166699), int32(0))
	mBase = m.M
	v9441 = m.ExcPending
	if v9441 != 0 {
		goto L66
	} else {
		goto L2547
	}
L2547:
	;
	v9444 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9445 = F_plpgsql_scanner_errposition(m, v9444, l1)
	mBase = m.M
	v9446 = m.ExcPending
	if v9446 != 0 {
		goto L66
	} else {
		goto L2548
	}
L2548:
	;
	F_errfinish(m, int32(26959), int32(1596), int32(360630))
	mBase = m.M
	v9451 = m.ExcPending
	if v9451 != 0 {
		goto L66
	} else {
		goto L2549
	}
L2549:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2550:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9458 = m.ExcPending
	if v9458 != 0 {
		goto L66
	} else {
		goto L2551
	}
L2551:
	;
	F_errmsg(m, int32(166877), int32(0))
	mBase = m.M
	v9462 = m.ExcPending
	if v9462 != 0 {
		goto L66
	} else {
		goto L2552
	}
L2552:
	;
	v9465 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(20))))
	v9466 = F_plpgsql_scanner_errposition(m, v9465, l1)
	mBase = m.M
	v9467 = m.ExcPending
	if v9467 != 0 {
		goto L66
	} else {
		goto L2553
	}
L2553:
	;
	F_errfinish(m, int32(26959), int32(1701), int32(360630))
	mBase = m.M
	v9472 = m.ExcPending
	if v9472 != 0 {
		goto L66
	} else {
		goto L2554
	}
L2554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2555:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9479 = m.ExcPending
	if v9479 != 0 {
		goto L66
	} else {
		goto L2556
	}
L2556:
	;
	v9480 = *(*int32)(unsafe.Add(mBase, uint32(v5530)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+224)) = v9480
	F_errmsg(m, int32(95529), v27+int32(224))
	mBase = m.M
	v9486 = m.ExcPending
	if v9486 != 0 {
		goto L66
	} else {
		goto L2557
	}
L2557:
	;
	v9489 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9490 = F_plpgsql_scanner_errposition(m, v9489, l1)
	mBase = m.M
	v9491 = m.ExcPending
	if v9491 != 0 {
		goto L66
	} else {
		goto L2558
	}
L2558:
	;
	F_errfinish(m, int32(26959), int32(1745), int32(360630))
	mBase = m.M
	v9496 = m.ExcPending
	if v9496 != 0 {
		goto L66
	} else {
		goto L2559
	}
L2559:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2560:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9503 = m.ExcPending
	if v9503 != 0 {
		goto L66
	} else {
		goto L2561
	}
L2561:
	;
	v9506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5464)+12)))
	if v9506 != 0 {
		goto L2562
	} else {
		goto L2563
	}
L2562:
	;
	v9507 = int32(308035)
	goto L2564
L2563:
	;
	v9507 = int32(234515)
	goto L2564
L2564:
	;
	F_errmsg(m, v9507, int32(0))
	mBase = m.M
	v9510 = m.ExcPending
	if v9510 != 0 {
		goto L66
	} else {
		goto L2565
	}
L2565:
	;
	v9511 = *(*int32)(unsafe.Add(mBase, uint32(v5480)))
	v9512 = F_plpgsql_scanner_errposition(m, v9511, l1)
	mBase = m.M
	v9513 = m.ExcPending
	if v9513 != 0 {
		goto L66
	} else {
		goto L2566
	}
L2566:
	;
	F_errfinish(m, int32(26959), int32(1767), int32(360630))
	mBase = m.M
	v9518 = m.ExcPending
	if v9518 != 0 {
		goto L66
	} else {
		goto L2567
	}
L2567:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2568:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9525 = m.ExcPending
	if v9525 != 0 {
		goto L66
	} else {
		goto L2569
	}
L2569:
	;
	F_errmsg(m, int32(254149), int32(0))
	mBase = m.M
	v9529 = m.ExcPending
	if v9529 != 0 {
		goto L66
	} else {
		goto L2570
	}
L2570:
	;
	v9530 = F_plpgsql_scanner_errposition(m, v5647, l1)
	mBase = m.M
	v9531 = m.ExcPending
	if v9531 != 0 {
		goto L66
	} else {
		goto L2571
	}
L2571:
	;
	F_errfinish(m, int32(26959), int32(3445), int32(97015))
	mBase = m.M
	v9536 = m.ExcPending
	if v9536 != 0 {
		goto L66
	} else {
		goto L2572
	}
L2572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2573:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9543 = m.ExcPending
	if v9543 != 0 {
		goto L66
	} else {
		goto L2574
	}
L2574:
	;
	F_errmsg(m, int32(132689), int32(0))
	mBase = m.M
	v9547 = m.ExcPending
	if v9547 != 0 {
		goto L66
	} else {
		goto L2575
	}
L2575:
	;
	v9548 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	v9549 = F_plpgsql_scanner_errposition(m, v9548, l1)
	mBase = m.M
	v9550 = m.ExcPending
	if v9550 != 0 {
		goto L66
	} else {
		goto L2576
	}
L2576:
	;
	F_errfinish(m, int32(26959), int32(3460), int32(97015))
	mBase = m.M
	v9555 = m.ExcPending
	if v9555 != 0 {
		goto L66
	} else {
		goto L2577
	}
L2577:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2578:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9562 = m.ExcPending
	if v9562 != 0 {
		goto L66
	} else {
		goto L2579
	}
L2579:
	;
	F_errmsg(m, int32(254101), int32(0))
	mBase = m.M
	v9566 = m.ExcPending
	if v9566 != 0 {
		goto L66
	} else {
		goto L2580
	}
L2580:
	;
	v9567 = F_plpgsql_scanner_errposition(m, v5803, l1)
	mBase = m.M
	v9568 = m.ExcPending
	if v9568 != 0 {
		goto L66
	} else {
		goto L2581
	}
L2581:
	;
	F_errfinish(m, int32(26959), int32(3509), int32(96992))
	mBase = m.M
	v9573 = m.ExcPending
	if v9573 != 0 {
		goto L66
	} else {
		goto L2582
	}
L2582:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2583:
	;
	F_errmsg(m, int32(363855), int32(0))
	mBase = m.M
	v9580 = m.ExcPending
	if v9580 != 0 {
		goto L66
	} else {
		goto L2584
	}
L2584:
	;
	v9581 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	v9582 = F_plpgsql_scanner_errposition(m, v9581, l1)
	mBase = m.M
	v9583 = m.ExcPending
	if v9583 != 0 {
		goto L66
	} else {
		goto L2585
	}
L2585:
	;
	F_errfinish(m, int32(26959), int32(3383), int32(97047))
	mBase = m.M
	v9588 = m.ExcPending
	if v9588 != 0 {
		goto L66
	} else {
		goto L2586
	}
L2586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2587:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9595 = m.ExcPending
	if v9595 != 0 {
		goto L66
	} else {
		goto L2588
	}
L2588:
	;
	F_errmsg(m, int32(132757), int32(0))
	mBase = m.M
	v9599 = m.ExcPending
	if v9599 != 0 {
		goto L66
	} else {
		goto L2589
	}
L2589:
	;
	v9600 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	v9601 = F_plpgsql_scanner_errposition(m, v9600, l1)
	mBase = m.M
	v9602 = m.ExcPending
	if v9602 != 0 {
		goto L66
	} else {
		goto L2590
	}
L2590:
	;
	F_errfinish(m, int32(26959), int32(3397), int32(97047))
	mBase = m.M
	v9607 = m.ExcPending
	if v9607 != 0 {
		goto L66
	} else {
		goto L2591
	}
L2591:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2592:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2594:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2596:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2597:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2598:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9651 = m.ExcPending
	if v9651 != 0 {
		goto L66
	} else {
		goto L2599
	}
L2599:
	;
	F_errmsg(m, int32(539131), int32(0))
	mBase = m.M
	v9655 = m.ExcPending
	if v9655 != 0 {
		goto L66
	} else {
		goto L2600
	}
L2600:
	;
	F_errfinish(m, int32(26959), int32(4162), int32(132225))
	mBase = m.M
	v9660 = m.ExcPending
	if v9660 != 0 {
		goto L66
	} else {
		goto L2601
	}
L2601:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2602:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2603:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2604:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2605:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2606:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2607:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v9691 = m.ExcPending
	if v9691 != 0 {
		goto L66
	} else {
		goto L2608
	}
L2608:
	;
	F_errmsg(m, int32(113841), int32(0))
	mBase = m.M
	v9695 = m.ExcPending
	if v9695 != 0 {
		goto L66
	} else {
		goto L2609
	}
L2609:
	;
	v9698 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v9699 = F_plpgsql_scanner_errposition(m, v9698, l1)
	mBase = m.M
	v9700 = m.ExcPending
	if v9700 != 0 {
		goto L66
	} else {
		goto L2610
	}
L2610:
	;
	F_errfinish(m, int32(26959), int32(2197), int32(360630))
	mBase = m.M
	v9705 = m.ExcPending
	if v9705 != 0 {
		goto L66
	} else {
		goto L2611
	}
L2611:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2612:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9712 = m.ExcPending
	if v9712 != 0 {
		goto L66
	} else {
		goto L2613
	}
L2613:
	;
	F_errmsg(m, int32(396522), int32(0))
	mBase = m.M
	v9716 = m.ExcPending
	if v9716 != 0 {
		goto L66
	} else {
		goto L2614
	}
L2614:
	;
	v9717 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9718 = F_plpgsql_scanner_errposition(m, v9717, l1)
	mBase = m.M
	v9719 = m.ExcPending
	if v9719 != 0 {
		goto L66
	} else {
		goto L2615
	}
L2615:
	;
	F_errfinish(m, int32(26959), int32(2294), int32(360630))
	mBase = m.M
	v9724 = m.ExcPending
	if v9724 != 0 {
		goto L66
	} else {
		goto L2616
	}
L2616:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2617:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9731 = m.ExcPending
	if v9731 != 0 {
		goto L66
	} else {
		goto L2618
	}
L2618:
	;
	v9732 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v9733 = *(*int32)(unsafe.Add(mBase, uint32(v9732)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v9733
	F_errmsg(m, int32(209831), v27+int32(256))
	mBase = m.M
	v9739 = m.ExcPending
	if v9739 != 0 {
		goto L66
	} else {
		goto L2619
	}
L2619:
	;
	v9740 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9741 = F_plpgsql_scanner_errposition(m, v9740, l1)
	mBase = m.M
	v9742 = m.ExcPending
	if v9742 != 0 {
		goto L66
	} else {
		goto L2620
	}
L2620:
	;
	F_errfinish(m, int32(26959), int32(2301), int32(360630))
	mBase = m.M
	v9747 = m.ExcPending
	if v9747 != 0 {
		goto L66
	} else {
		goto L2621
	}
L2621:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2622:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2623:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2624:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2625:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2626:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2627:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9787 = m.ExcPending
	if v9787 != 0 {
		goto L66
	} else {
		goto L2628
	}
L2628:
	;
	F_errmsg(m, int32(539091), int32(0))
	mBase = m.M
	v9791 = m.ExcPending
	if v9791 != 0 {
		goto L66
	} else {
		goto L2629
	}
L2629:
	;
	F_errfinish(m, int32(26959), int32(4158), int32(132225))
	mBase = m.M
	v9796 = m.ExcPending
	if v9796 != 0 {
		goto L66
	} else {
		goto L2630
	}
L2630:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2631:
	;
	v9816 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377])))
	if v9816 == int32(269) {
		v9837 = v9814
		v9838 = v9798
		goto L9
	} else {
		goto L2632
	}
L2632:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9822 = m.ExcPending
	if v9822 != 0 {
		goto L66
	} else {
		goto L2633
	}
L2633:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9825 = m.ExcPending
	if v9825 != 0 {
		goto L66
	} else {
		goto L2634
	}
L2634:
	;
	F_errmsg(m, int32(234554), int32(0))
	mBase = m.M
	v9829 = m.ExcPending
	if v9829 != 0 {
		goto L66
	} else {
		goto L2635
	}
L2635:
	;
	v9830 = F_plpgsql_scanner_errposition(m, v4801, l1)
	mBase = m.M
	v9831 = m.ExcPending
	if v9831 != 0 {
		goto L66
	} else {
		goto L2636
	}
L2636:
	;
	F_errfinish(m, int32(26959), int32(1569), int32(360630))
	mBase = m.M
	v9836 = m.ExcPending
	if v9836 != 0 {
		goto L66
	} else {
		goto L2637
	}
L2637:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2638:
	;
	v9845 = *(*int32)(unsafe.Add(mBase, uint32(v9837)))
	v9846 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1397])))
	v9847 = int32(4515248)
	v9848 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v9851 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v9851
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v9846
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1380]))) = int32(6810)
	v9857 = int32(4508152)
	v9858 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4784)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367]))) = v9858
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1381]))) = v27 + int32(4768)
	v9868 = F_raw_parser(m, v9845, int32(2))
	mBase = m.M
	v9869 = m.ExcPending
	if v9869 != 0 {
		goto L66
	} else {
		goto L2641
	}
L2639:
	;
	goto L2640
L2640:
	;
	v9879 = int32(0)
	v9885 = int32(1)
	v9894 = F_read_sql_construct(m, int32(336), int32(288), v9879, int32(526588), int32(2), v9885, v9885, v9879, v27+int32(4752), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9895 = m.ExcPending
	if v9895 != 0 {
		goto L66
	} else {
		goto L2642
	}
L2641:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v9848
	v9873 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v9873
	goto L2640
L2642:
	;
	v9896 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377])))
	if v9896 == int32(288) {
		goto L2643
	} else {
		goto L2644
	}
L2643:
	;
	v9900 = int32(0)
	v9904 = int32(1)
	v9912 = F_read_sql_construct(m, int32(336), v9900, v9900, int32(526588), int32(2), v9904, v9904, v9900, v9900, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9913 = m.ExcPending
	if v9913 != 0 {
		goto L66
	} else {
		goto L2646
	}
L2644:
	;
	v9914 = v9879
	goto L2645
L2645:
	;
	v9917 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(8))))
	if v9917 != 0 {
		goto L2648
	} else {
		goto L2649
	}
L2646:
	;
	v9914 = v9912
	goto L2645
L2647:
	;
	F_errstart_cold(m, int32(21), int32(556047))
	mBase = m.M
	v9957 = m.ExcPending
	if v9957 != 0 {
		goto L66
	} else {
		goto L2655
	}
L2648:
	;
	v9920 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
	if v9920 != 0 {
		goto L2647
	} else {
		goto L2651
	}
L2649:
	;
	goto L2650
L2650:
	;
	v9923 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v9926 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v9929 = int32(0)
	v9931 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v9929, v9929)
	mBase = m.M
	v9932 = m.ExcPending
	if v9932 != 0 {
		goto L66
	} else {
		goto L2652
	}
L2651:
	;
	goto L2650
L2652:
	;
	v9934 = F_plpgsql_build_variable(m, v9923, v9926, v9931, int32(1))
	mBase = m.M
	v9935 = m.ExcPending
	if v9935 != 0 {
		goto L66
	} else {
		goto L2653
	}
L2653:
	;
	v9937 = F_palloc0(m, int32(40))
	mBase = m.M
	v9938 = m.ExcPending
	if v9938 != 0 {
		goto L66
	} else {
		goto L2654
	}
L2654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9937))) = int32(6)
	v9942 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v9943 = *(*int32)(unsafe.Add(mBase, uint32(v9942)+520))
	v9945 = v9943 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9942)+520)) = v9945
	*(*int32)(unsafe.Add(mBase, uint32(v9937)+32)) = v9838
	*(*int32)(unsafe.Add(mBase, uint32(v9937)+16)) = v9934
	*(*int32)(unsafe.Add(mBase, uint32(v9937)+8)) = v9945
	*(*int32)(unsafe.Add(mBase, uint32(v9937)+28)) = v9914
	*(*int32)(unsafe.Add(mBase, uint32(v9937)+24)) = v9894
	*(*int32)(unsafe.Add(mBase, uint32(v9937)+20)) = v9837
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9937
	v10149 = v241
	goto L5
L2655:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9960 = m.ExcPending
	if v9960 != 0 {
		goto L66
	} else {
		goto L2656
	}
L2656:
	;
	F_errmsg(m, int32(396261), int32(0))
	mBase = m.M
	v9964 = m.ExcPending
	if v9964 != 0 {
		goto L66
	} else {
		goto L2657
	}
L2657:
	;
	v9967 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9968 = F_plpgsql_scanner_errposition(m, v9967, l1)
	mBase = m.M
	v9969 = m.ExcPending
	if v9969 != 0 {
		goto L66
	} else {
		goto L2658
	}
L2658:
	;
	F_errfinish(m, int32(26959), int32(1535), int32(360630))
	mBase = m.M
	v9974 = m.ExcPending
	if v9974 != 0 {
		goto L66
	} else {
		goto L2659
	}
L2659:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2660:
	;
	v10120 = v2506
	goto L6
L2661:
	;
	if v9995 <= int32(292) {
		goto L2668
	} else {
		goto L2669
	}
L2662:
	;
	F_initStringInfo(m, v27+int32(4784))
	mBase = m.M
	v10061 = m.ExcPending
	if v10061 != 0 {
		goto L66
	} else {
		goto L2686
	}
L2663:
	;
	goto L2662
L2664:
	;
	if base.B2i32(v9995 != int32(44))&base.B2i32(v9995 != int32(41))|v9989 == int32(0) {
		goto L2663
	} else {
		goto L2678
	}
L2665:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(161213))
	mBase = m.M
	v10033 = m.ExcPending
	if v10033 != 0 {
		goto L66
	} else {
		goto L2677
	}
L2666:
	;
	if v9995 != int32(343) {
		goto L2664
	} else {
		goto L2676
	}
L2667:
	;
	if v9989 != 0 {
		goto L2665
	} else {
		goto L2674
	}
L2668:
	;
	switch v9995 - int32(59) {
	case 0, 2:
		goto L2663
	case 1:
		goto L2664
	default:
		goto L2671
	}
L2669:
	;
	goto L2670
L2670:
	;
	switch v9995 - int32(293) {
	case 0, 13:
		goto L2663
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L2664
	default:
		goto L2666
	}
L2671:
	;
	if v9995 == int32(270) {
		goto L2663
	} else {
		goto L2672
	}
L2672:
	;
	if v9995 == int32(0) {
		goto L2667
	} else {
		goto L2673
	}
L2673:
	;
	goto L2664
L2674:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(261107))
	mBase = m.M
	v10025 = m.ExcPending
	if v10025 != 0 {
		goto L66
	} else {
		goto L2675
	}
L2675:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2676:
	;
	goto L2663
L2677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2678:
	;
	if v9995 == int32(41) {
		goto L2679
	} else {
		goto L2680
	}
L2679:
	;
	v10047 = int32(-1)
	goto L2681
L2680:
	;
	v10047 = int32(0)
	goto L2681
L2681:
	;
	if v9995 == int32(40) {
		goto L2682
	} else {
		goto L2683
	}
L2682:
	;
	v10050 = int32(1)
	goto L2684
L2683:
	;
	v10050 = v10047
	goto L2684
L2684:
	;
	v10056 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v10057 = m.ExcPending
	if v10057 != 0 {
		goto L66
	} else {
		goto L2685
	}
L2685:
	;
	v9989 = v10050 + v9989
	v9995 = v10056
	goto L2661
L2686:
	;
	v10064 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1365])))
	F_plpgsql_append_source_text(m, v27+int32(4784), v2225, v10064, l1)
	mBase = m.M
	v10066 = m.ExcPending
	if v10066 != 0 {
		goto L66
	} else {
		goto L2687
	}
L2687:
	;
	v10067 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	v10068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10067))))
	if v10068 == int32(0) {
		goto L3
	} else {
		goto L2688
	}
L2688:
	;
	v10071 = int32(4508152)
	v10072 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4768)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1425]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377]))) = v2225
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = int32(6810)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = v10072
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1426]))) = v27 + int32(4752)
	v10085 = int32(0)
	v10087 = F_typeStringToTypeName(m, v10067, v10085)
	mBase = m.M
	v10088 = m.ExcPending
	if v10088 != 0 {
		goto L66
	} else {
		goto L2689
	}
L2689:
	;
	F_typenameTypeIdAndMod(m, v10085, v10087, v27+int32(4764), v27+int32(4760))
	mBase = m.M
	v10094 = m.ExcPending
	if v10094 != 0 {
		goto L66
	} else {
		goto L2690
	}
L2690:
	;
	v10096 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v10096
	v10098 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1397])))
	v10099 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1427])))
	v10101 = *(*int32)(unsafe.Add(mBase, _consts[1368]))
	v10102 = *(*int32)(unsafe.Add(mBase, uint32(v10101)+44))
	v10103 = F_plpgsql_build_datatype(m, v10098, v10099, v10102, v10087)
	mBase = m.M
	v10104 = m.ExcPending
	if v10104 != 0 {
		goto L66
	} else {
		goto L2691
	}
L2691:
	;
	v10105 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1367])))
	F_pfree(m, v10105)
	mBase = m.M
	v10107 = m.ExcPending
	if v10107 != 0 {
		goto L66
	} else {
		goto L2692
	}
L2692:
	;
	F_plpgsql_push_back_token(m, v9995, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v10113 = m.ExcPending
	if v10113 != 0 {
		goto L66
	} else {
		goto L2693
	}
L2693:
	;
	v10120 = v10103
	goto L6
L2694:
	;
	v10208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10186)+uint32(_consts[1428]))))
	v10212 = v10208
	v10214 = v10175
	v10218 = v10149
	v10223 = v10178
	v10226 = v260
	goto L4
L2695:
	;
	v10197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10190<<(uint(int32(1))%32))+uint32(_consts[1361]))))
	if v10197 != v10179&int32(65535) {
		goto L2694
	} else {
		goto L2696
	}
L2696:
	;
	v10205 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10190<<(uint(int32(1))%32))+uint32(_consts[1362]))))
	v10212 = v10205
	v10214 = v10175
	v10218 = v10149
	v10223 = v10178
	v10226 = v260
	goto L4
L2697:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
