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
	var v2876 int32
	_ = v2876
	var v2881 int32
	_ = v2881
	var v2885 int32
	_ = v2885
	var v2890 int32
	_ = v2890
	var v2892 int32
	_ = v2892
	var v2896 int32
	_ = v2896
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2911 int32
	_ = v2911
	var v2915 int32
	_ = v2915
	var v2917 int32
	_ = v2917
	var v2925 int32
	_ = v2925
	var v2929 int32
	_ = v2929
	var v2936 int32
	_ = v2936
	var v2939 int32
	_ = v2939
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2952 int32
	_ = v2952
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2957 int32
	_ = v2957
	var v2959 int32
	_ = v2959
	var v2966 int32
	_ = v2966
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2981 int32
	_ = v2981
	var v2988 int32
	_ = v2988
	var v2992 int32
	_ = v2992
	var v2995 int32
	_ = v2995
	var v3001 int32
	_ = v3001
	var v3008 int32
	_ = v3008
	var v3012 int32
	_ = v3012
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3029 int32
	_ = v3029
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3034 int32
	_ = v3034
	var v3036 int32
	_ = v3036
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3051 int32
	_ = v3051
	var v3053 int32
	_ = v3053
	var v3058 int32
	_ = v3058
	var v3071 int32
	_ = v3071
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3105 int32
	_ = v3105
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3132 int32
	_ = v3132
	var v3137 int32
	_ = v3137
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3144 int32
	_ = v3144
	var v3146 int32
	_ = v3146
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3153 int32
	_ = v3153
	var v3156 int32
	_ = v3156
	var v3163 int32
	_ = v3163
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3180 int32
	_ = v3180
	var v3182 int32
	_ = v3182
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3200 int32
	_ = v3200
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
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3243 int32
	_ = v3243
	var v3246 int32
	_ = v3246
	var v3253 int32
	_ = v3253
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3261 int32
	_ = v3261
	var v3270 int32
	_ = v3270
	var v3272 int32
	_ = v3272
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3295 int32
	_ = v3295
	var v3298 int32
	_ = v3298
	var v3301 int32
	_ = v3301
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3313 int32
	_ = v3313
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3337 int32
	_ = v3337
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3344 int32
	_ = v3344
	var v3346 int32
	_ = v3346
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3353 int32
	_ = v3353
	var v3356 int32
	_ = v3356
	var v3363 int32
	_ = v3363
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3371 int32
	_ = v3371
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3400 int32
	_ = v3400
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3405 int32
	_ = v3405
	var v3406 int32
	_ = v3406
	var v3411 int32
	_ = v3411
	var v3412 int32
	_ = v3412
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3427 int32
	_ = v3427
	var v3428 int32
	_ = v3428
	var v3430 int32
	_ = v3430
	var v3435 int32
	_ = v3435
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3451 int32
	_ = v3451
	var v3454 int32
	_ = v3454
	var v3461 int32
	_ = v3461
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3474 int32
	_ = v3474
	var v3478 int32
	_ = v3478
	var v3482 int32
	_ = v3482
	var v3485 int32
	_ = v3485
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3494 int32
	_ = v3494
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3531 int32
	_ = v3531
	var v3534 int32
	_ = v3534
	var v3535 int32
	_ = v3535
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3550 int32
	_ = v3550
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3558 int32
	_ = v3558
	var v3564 int32
	_ = v3564
	var v3567 int32
	_ = v3567
	var v3568 int32
	_ = v3568
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3591 int32
	_ = v3591
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3602 int32
	_ = v3602
	var v3607 int32
	_ = v3607
	var v3609 int32
	_ = v3609
	var v3636 int32
	_ = v3636
	var v3638 int32
	_ = v3638
	var v3640 int32
	_ = v3640
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3649 int32
	_ = v3649
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3659 int32
	_ = v3659
	var v3660 int32
	_ = v3660
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3666 int32
	_ = v3666
	var v3673 int32
	_ = v3673
	var v3674 int32
	_ = v3674
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3693 int32
	_ = v3693
	var v3694 int32
	_ = v3694
	var v3697 int32
	_ = v3697
	var v3704 int32
	_ = v3704
	var v3705 int32
	_ = v3705
	var v3710 int32
	_ = v3710
	var v3713 int32
	_ = v3713
	var v3714 int32
	_ = v3714
	var v3718 int32
	_ = v3718
	var v3719 int32
	_ = v3719
	var v3722 int32
	_ = v3722
	var v3723 int32
	_ = v3723
	var v3726 int32
	_ = v3726
	var v3733 int32
	_ = v3733
	var v3734 int32
	_ = v3734
	var v3739 int32
	_ = v3739
	var v3742 int32
	_ = v3742
	var v3745 int32
	_ = v3745
	var v3746 int32
	_ = v3746
	var v3750 int32
	_ = v3750
	var v3751 int32
	_ = v3751
	var v3754 int32
	_ = v3754
	var v3755 int32
	_ = v3755
	var v3758 int32
	_ = v3758
	var v3765 int32
	_ = v3765
	var v3766 int32
	_ = v3766
	var v3771 int32
	_ = v3771
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3779 int32
	_ = v3779
	var v3780 int32
	_ = v3780
	var v3783 int32
	_ = v3783
	var v3784 int32
	_ = v3784
	var v3787 int32
	_ = v3787
	var v3794 int32
	_ = v3794
	var v3795 int32
	_ = v3795
	var v3800 int32
	_ = v3800
	var v3803 int32
	_ = v3803
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3811 int32
	_ = v3811
	var v3812 int32
	_ = v3812
	var v3815 int32
	_ = v3815
	var v3816 int32
	_ = v3816
	var v3819 int32
	_ = v3819
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3832 int32
	_ = v3832
	var v3835 int32
	_ = v3835
	var v3836 int32
	_ = v3836
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
	var v3868 int32
	_ = v3868
	var v3872 int32
	_ = v3872
	var v3873 int32
	_ = v3873
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3880 int32
	_ = v3880
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3893 int32
	_ = v3893
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3901 int32
	_ = v3901
	var v3902 int32
	_ = v3902
	var v3905 int32
	_ = v3905
	var v3906 int32
	_ = v3906
	var v3909 int32
	_ = v3909
	var v3916 int32
	_ = v3916
	var v3917 int32
	_ = v3917
	var v3922 int32
	_ = v3922
	var v3925 int32
	_ = v3925
	var v3928 int32
	_ = v3928
	var v3929 int32
	_ = v3929
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
	var v3958 int32
	_ = v3958
	var v3962 int32
	_ = v3962
	var v3963 int32
	_ = v3963
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3970 int32
	_ = v3970
	var v3977 int32
	_ = v3977
	var v3978 int32
	_ = v3978
	var v3983 int32
	_ = v3983
	var v3986 int32
	_ = v3986
	var v3989 int32
	_ = v3989
	var v3990 int32
	_ = v3990
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3998 int32
	_ = v3998
	var v3999 int32
	_ = v3999
	var v4002 int32
	_ = v4002
	var v4009 int32
	_ = v4009
	var v4010 int32
	_ = v4010
	var v4015 int32
	_ = v4015
	var v4018 int32
	_ = v4018
	var v4019 int32
	_ = v4019
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4027 int32
	_ = v4027
	var v4028 int32
	_ = v4028
	var v4031 int32
	_ = v4031
	var v4038 int32
	_ = v4038
	var v4039 int32
	_ = v4039
	var v4044 int32
	_ = v4044
	var v4047 int32
	_ = v4047
	var v4050 int32
	_ = v4050
	var v4051 int32
	_ = v4051
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4059 int32
	_ = v4059
	var v4060 int32
	_ = v4060
	var v4063 int32
	_ = v4063
	var v4070 int32
	_ = v4070
	var v4071 int32
	_ = v4071
	var v4082 int32
	_ = v4082
	var v4083 int32
	_ = v4083
	var v4084 int32
	_ = v4084
	var v4089 int32
	_ = v4089
	var v4090 int32
	_ = v4090
	var v4093 int32
	_ = v4093
	var v4094 int32
	_ = v4094
	var v4096 int32
	_ = v4096
	var v4097 int32
	_ = v4097
	var v4099 int32
	_ = v4099
	var v4101 int32
	_ = v4101
	var v4102 int32
	_ = v4102
	var v4104 int32
	_ = v4104
	var v4106 int32
	_ = v4106
	var v4107 int32
	_ = v4107
	var v4112 int32
	_ = v4112
	var v4113 int32
	_ = v4113
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4122 int32
	_ = v4122
	var v4123 int32
	_ = v4123
	var v4125 int32
	_ = v4125
	var v4130 int32
	_ = v4130
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4137 int32
	_ = v4137
	var v4139 int32
	_ = v4139
	var v4142 int32
	_ = v4142
	var v4143 int32
	_ = v4143
	var v4146 int32
	_ = v4146
	var v4149 int32
	_ = v4149
	var v4156 int32
	_ = v4156
	var v4161 int32
	_ = v4161
	var v4162 int32
	_ = v4162
	var v4164 int32
	_ = v4164
	var v4169 int32
	_ = v4169
	var v4173 int32
	_ = v4173
	var v4177 int32
	_ = v4177
	var v4181 int32
	_ = v4181
	var v4187 int32
	_ = v4187
	var v4188 int32
	_ = v4188
	var v4191 int32
	_ = v4191
	var v4192 int32
	_ = v4192
	var v4197 int32
	_ = v4197
	var v4198 int32
	_ = v4198
	var v4201 int32
	_ = v4201
	var v4202 int32
	_ = v4202
	var v4204 int32
	_ = v4204
	var v4209 int32
	_ = v4209
	var v4211 int32
	_ = v4211
	var v4212 int32
	_ = v4212
	var v4216 int32
	_ = v4216
	var v4218 int32
	_ = v4218
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4225 int32
	_ = v4225
	var v4228 int32
	_ = v4228
	var v4235 int32
	_ = v4235
	var v4241 int32
	_ = v4241
	var v4243 int32
	_ = v4243
	var v4247 int32
	_ = v4247
	var v4248 int32
	_ = v4248
	var v4249 int32
	_ = v4249
	var v4253 int32
	_ = v4253
	var v4257 int32
	_ = v4257
	var v4260 int32
	_ = v4260
	var v4263 int32
	_ = v4263
	var v4266 int32
	_ = v4266
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4272 int32
	_ = v4272
	var v4277 int32
	_ = v4277
	var v4278 int32
	_ = v4278
	var v4281 int32
	_ = v4281
	var v4282 int32
	_ = v4282
	var v4284 int32
	_ = v4284
	var v4289 int32
	_ = v4289
	var v4291 int32
	_ = v4291
	var v4292 int32
	_ = v4292
	var v4296 int32
	_ = v4296
	var v4298 int32
	_ = v4298
	var v4301 int32
	_ = v4301
	var v4302 int32
	_ = v4302
	var v4305 int32
	_ = v4305
	var v4308 int32
	_ = v4308
	var v4315 int32
	_ = v4315
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4323 int32
	_ = v4323
	var v4325 int32
	_ = v4325
	var v4335 int32
	_ = v4335
	var v4338 int32
	_ = v4338
	var v4339 int32
	_ = v4339
	var v4341 int32
	_ = v4341
	var v4346 int32
	_ = v4346
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4356 int32
	_ = v4356
	var v4359 int32
	_ = v4359
	var v4364 int32
	_ = v4364
	var v4365 int32
	_ = v4365
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4369 int32
	_ = v4369
	var v4373 int32
	_ = v4373
	var v4380 int32
	_ = v4380
	var v4400 int32
	_ = v4400
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4421 int32
	_ = v4421
	var v4422 int32
	_ = v4422
	var v4424 int32
	_ = v4424
	var v4425 int32
	_ = v4425
	var v4426 int32
	_ = v4426
	var v4427 int32
	_ = v4427
	var v4430 int32
	_ = v4430
	var v4432 int32
	_ = v4432
	var v4434 int32
	_ = v4434
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4469 int32
	_ = v4469
	var v4470 int32
	_ = v4470
	var v4478 int32
	_ = v4478
	var v4480 int32
	_ = v4480
	var v4484 int32
	_ = v4484
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4494 int32
	_ = v4494
	var v4501 int32
	_ = v4501
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4507 int32
	_ = v4507
	var v4508 int32
	_ = v4508
	var v4510 int32
	_ = v4510
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4520 int32
	_ = v4520
	var v4521 int32
	_ = v4521
	var v4524 int32
	_ = v4524
	var v4525 int32
	_ = v4525
	var v4530 int32
	_ = v4530
	var v4531 int32
	_ = v4531
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4537 int32
	_ = v4537
	var v4542 int32
	_ = v4542
	var v4544 int32
	_ = v4544
	var v4545 int32
	_ = v4545
	var v4549 int32
	_ = v4549
	var v4551 int32
	_ = v4551
	var v4554 int32
	_ = v4554
	var v4555 int32
	_ = v4555
	var v4558 int32
	_ = v4558
	var v4561 int32
	_ = v4561
	var v4568 int32
	_ = v4568
	var v4574 int32
	_ = v4574
	var v4576 int32
	_ = v4576
	var v4581 int32
	_ = v4581
	var v4583 int32
	_ = v4583
	var v4590 int32
	_ = v4590
	var v4591 int32
	_ = v4591
	var v4594 int32
	_ = v4594
	var v4595 int32
	_ = v4595
	var v4596 int32
	_ = v4596
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4606 int32
	_ = v4606
	var v4607 int32
	_ = v4607
	var v4610 int32
	_ = v4610
	var v4611 int32
	_ = v4611
	var v4613 int32
	_ = v4613
	var v4618 int32
	_ = v4618
	var v4620 int32
	_ = v4620
	var v4621 int32
	_ = v4621
	var v4625 int32
	_ = v4625
	var v4627 int32
	_ = v4627
	var v4630 int32
	_ = v4630
	var v4631 int32
	_ = v4631
	var v4634 int32
	_ = v4634
	var v4637 int32
	_ = v4637
	var v4644 int32
	_ = v4644
	var v4649 int32
	_ = v4649
	var v4650 int32
	_ = v4650
	var v4652 int32
	_ = v4652
	var v4656 int32
	_ = v4656
	var v4657 int32
	_ = v4657
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
	var v4668 int32
	_ = v4668
	var v4669 int32
	_ = v4669
	var v4670 int32
	_ = v4670
	var v4671 int32
	_ = v4671
	var v4672 int32
	_ = v4672
	var v4673 int32
	_ = v4673
	var v4675 int32
	_ = v4675
	var v4679 int32
	_ = v4679
	var v4680 int32
	_ = v4680
	var v4685 int32
	_ = v4685
	var v4686 int32
	_ = v4686
	var v4691 int32
	_ = v4691
	var v4692 int32
	_ = v4692
	var v4695 int32
	_ = v4695
	var v4696 int32
	_ = v4696
	var v4698 int32
	_ = v4698
	var v4703 int32
	_ = v4703
	var v4705 int32
	_ = v4705
	var v4706 int32
	_ = v4706
	var v4710 int32
	_ = v4710
	var v4712 int32
	_ = v4712
	var v4715 int32
	_ = v4715
	var v4716 int32
	_ = v4716
	var v4719 int32
	_ = v4719
	var v4722 int32
	_ = v4722
	var v4729 int32
	_ = v4729
	var v4734 int32
	_ = v4734
	var v4735 int32
	_ = v4735
	var v4737 int32
	_ = v4737
	var v4741 int32
	_ = v4741
	var v4742 int32
	_ = v4742
	var v4746 int32
	_ = v4746
	var v4748 int32
	_ = v4748
	var v4750 int32
	_ = v4750
	var v4751 int32
	_ = v4751
	var v4752 int32
	_ = v4752
	var v4754 int32
	_ = v4754
	var v4757 int32
	_ = v4757
	var v4758 int32
	_ = v4758
	var v4759 int32
	_ = v4759
	var v4760 int32
	_ = v4760
	var v4761 int32
	_ = v4761
	var v4762 int32
	_ = v4762
	var v4764 int32
	_ = v4764
	var v4769 int32
	_ = v4769
	var v4770 int32
	_ = v4770
	var v4773 int32
	_ = v4773
	var v4774 int32
	_ = v4774
	var v4779 int32
	_ = v4779
	var v4780 int32
	_ = v4780
	var v4783 int32
	_ = v4783
	var v4784 int32
	_ = v4784
	var v4786 int32
	_ = v4786
	var v4791 int32
	_ = v4791
	var v4793 int32
	_ = v4793
	var v4794 int32
	_ = v4794
	var v4798 int32
	_ = v4798
	var v4800 int32
	_ = v4800
	var v4803 int32
	_ = v4803
	var v4804 int32
	_ = v4804
	var v4807 int32
	_ = v4807
	var v4810 int32
	_ = v4810
	var v4817 int32
	_ = v4817
	var v4822 int32
	_ = v4822
	var v4823 int32
	_ = v4823
	var v4829 int32
	_ = v4829
	var v4831 int32
	_ = v4831
	var v4834 int32
	_ = v4834
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4838 int32
	_ = v4838
	var v4841 int32
	_ = v4841
	var v4842 int32
	_ = v4842
	var v4843 int32
	_ = v4843
	var v4844 int32
	_ = v4844
	var v4845 int32
	_ = v4845
	var v4846 int32
	_ = v4846
	var v4848 int32
	_ = v4848
	var v4854 int32
	_ = v4854
	var v4855 int32
	_ = v4855
	var v4857 int32
	_ = v4857
	var v4864 int32
	_ = v4864
	var v4867 int32
	_ = v4867
	var v4876 int32
	_ = v4876
	var v4877 int32
	_ = v4877
	var v4879 int32
	_ = v4879
	var v4880 int32
	_ = v4880
	var v4884 int32
	_ = v4884
	var v4885 int32
	_ = v4885
	var v4887 int32
	_ = v4887
	var v4891 int32
	_ = v4891
	var v4892 int32
	_ = v4892
	var v4894 int32
	_ = v4894
	var v4897 int32
	_ = v4897
	var v4899 int32
	_ = v4899
	var v4902 int32
	_ = v4902
	var v4907 int32
	_ = v4907
	var v4910 int32
	_ = v4910
	var v4913 int32
	_ = v4913
	var v4915 int32
	_ = v4915
	var v4917 int32
	_ = v4917
	var v4918 int32
	_ = v4918
	var v4927 int32
	_ = v4927
	var v4928 int32
	_ = v4928
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4934 int32
	_ = v4934
	var v4936 int32
	_ = v4936
	var v4939 int32
	_ = v4939
	var v4946 int32
	_ = v4946
	var v4975 int32
	_ = v4975
	var v4978 int32
	_ = v4978
	var v4987 int32
	_ = v4987
	var v4988 int32
	_ = v4988
	var v4989 int32
	_ = v4989
	var v4990 int32
	_ = v4990
	var v4991 int32
	_ = v4991
	var v4993 int32
	_ = v4993
	var v5021 int32
	_ = v5021
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5024 int32
	_ = v5024
	var v5028 int32
	_ = v5028
	var v5029 int32
	_ = v5029
	var v5030 int32
	_ = v5030
	var v5033 int32
	_ = v5033
	var v5034 int32
	_ = v5034
	var v5036 int32
	_ = v5036
	var v5039 int32
	_ = v5039
	var v5043 int32
	_ = v5043
	var v5046 int32
	_ = v5046
	var v5047 int32
	_ = v5047
	var v5055 int32
	_ = v5055
	var v5056 int32
	_ = v5056
	var v5060 int32
	_ = v5060
	var v5063 int32
	_ = v5063
	var v5067 int32
	_ = v5067
	var v5068 int32
	_ = v5068
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5075 int32
	_ = v5075
	var v5078 int32
	_ = v5078
	var v5079 int32
	_ = v5079
	var v5083 int32
	_ = v5083
	var v5084 int32
	_ = v5084
	var v5087 int32
	_ = v5087
	var v5088 int32
	_ = v5088
	var v5091 int32
	_ = v5091
	var v5098 int32
	_ = v5098
	var v5099 int32
	_ = v5099
	var v5109 int32
	_ = v5109
	var v5110 int32
	_ = v5110
	var v5126 int32
	_ = v5126
	var v5127 int32
	_ = v5127
	var v5128 int32
	_ = v5128
	var v5132 int32
	_ = v5132
	var v5135 int32
	_ = v5135
	var v5136 int32
	_ = v5136
	var v5137 int32
	_ = v5137
	var v5138 int32
	_ = v5138
	var v5139 int32
	_ = v5139
	var v5142 int32
	_ = v5142
	var v5148 int32
	_ = v5148
	var v5149 int32
	_ = v5149
	var v5158 int32
	_ = v5158
	var v5159 int32
	_ = v5159
	var v5163 int32
	_ = v5163
	var v5171 int32
	_ = v5171
	var v5172 int32
	_ = v5172
	var v5176 int32
	_ = v5176
	var v5177 int32
	_ = v5177
	var v5179 int32
	_ = v5179
	var v5183 int32
	_ = v5183
	var v5184 int32
	_ = v5184
	var v5186 int32
	_ = v5186
	var v5189 int32
	_ = v5189
	var v5191 int32
	_ = v5191
	var v5194 int32
	_ = v5194
	var v5199 int32
	_ = v5199
	var v5202 int32
	_ = v5202
	var v5205 int32
	_ = v5205
	var v5207 int32
	_ = v5207
	var v5209 int32
	_ = v5209
	var v5210 int32
	_ = v5210
	var v5219 int32
	_ = v5219
	var v5220 int32
	_ = v5220
	var v5223 int32
	_ = v5223
	var v5224 int32
	_ = v5224
	var v5226 int32
	_ = v5226
	var v5228 int32
	_ = v5228
	var v5231 int32
	_ = v5231
	var v5239 int32
	_ = v5239
	var v5240 int32
	_ = v5240
	var v5241 int32
	_ = v5241
	var v5242 int32
	_ = v5242
	var v5243 int32
	_ = v5243
	var v5245 int32
	_ = v5245
	var v5246 int32
	_ = v5246
	var v5251 int32
	_ = v5251
	var v5252 int32
	_ = v5252
	var v5255 int32
	_ = v5255
	var v5256 int32
	_ = v5256
	var v5258 int32
	_ = v5258
	var v5263 int32
	_ = v5263
	var v5265 int32
	_ = v5265
	var v5266 int32
	_ = v5266
	var v5270 int32
	_ = v5270
	var v5272 int32
	_ = v5272
	var v5275 int32
	_ = v5275
	var v5276 int32
	_ = v5276
	var v5279 int32
	_ = v5279
	var v5282 int32
	_ = v5282
	var v5289 int32
	_ = v5289
	var v5293 int32
	_ = v5293
	var v5294 int32
	_ = v5294
	var v5295 int32
	_ = v5295
	var v5301 int32
	_ = v5301
	var v5310 int32
	_ = v5310
	var v5311 int32
	_ = v5311
	var v5317 int32
	_ = v5317
	var v5320 int32
	_ = v5320
	var v5321 int32
	_ = v5321
	var v5322 int32
	_ = v5322
	var v5327 int32
	_ = v5327
	var v5328 int32
	_ = v5328
	var v5330 int32
	_ = v5330
	var v5332 int32
	_ = v5332
	var v5333 int32
	_ = v5333
	var v5338 int32
	_ = v5338
	var v5339 int32
	_ = v5339
	var v5342 int32
	_ = v5342
	var v5343 int32
	_ = v5343
	var v5345 int32
	_ = v5345
	var v5350 int32
	_ = v5350
	var v5352 int32
	_ = v5352
	var v5353 int32
	_ = v5353
	var v5357 int32
	_ = v5357
	var v5359 int32
	_ = v5359
	var v5362 int32
	_ = v5362
	var v5363 int32
	_ = v5363
	var v5366 int32
	_ = v5366
	var v5369 int32
	_ = v5369
	var v5376 int32
	_ = v5376
	var v5386 int32
	_ = v5386
	var v5387 int32
	_ = v5387
	var v5393 int32
	_ = v5393
	var v5396 int32
	_ = v5396
	var v5398 int32
	_ = v5398
	var v5399 int32
	_ = v5399
	var v5401 int32
	_ = v5401
	var v5403 int32
	_ = v5403
	var v5404 int32
	_ = v5404
	var v5409 int32
	_ = v5409
	var v5410 int32
	_ = v5410
	var v5415 int32
	_ = v5415
	var v5416 int32
	_ = v5416
	var v5419 int32
	_ = v5419
	var v5420 int32
	_ = v5420
	var v5422 int32
	_ = v5422
	var v5427 int32
	_ = v5427
	var v5429 int32
	_ = v5429
	var v5430 int32
	_ = v5430
	var v5434 int32
	_ = v5434
	var v5436 int32
	_ = v5436
	var v5439 int32
	_ = v5439
	var v5440 int32
	_ = v5440
	var v5443 int32
	_ = v5443
	var v5446 int32
	_ = v5446
	var v5453 int32
	_ = v5453
	var v5458 int32
	_ = v5458
	var v5459 int32
	_ = v5459
	var v5461 int32
	_ = v5461
	var v5465 int32
	_ = v5465
	var v5466 int32
	_ = v5466
	var v5470 int32
	_ = v5470
	var v5474 int32
	_ = v5474
	var v5476 int32
	_ = v5476
	var v5479 int32
	_ = v5479
	var v5480 int32
	_ = v5480
	var v5484 int32
	_ = v5484
	var v5485 int32
	_ = v5485
	var v5488 int32
	_ = v5488
	var v5489 int32
	_ = v5489
	var v5490 int32
	_ = v5490
	var v5492 int32
	_ = v5492
	var v5495 int32
	_ = v5495
	var v5497 int32
	_ = v5497
	var v5498 int32
	_ = v5498
	var v5499 int32
	_ = v5499
	var v5500 int32
	_ = v5500
	var v5502 int32
	_ = v5502
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
	var v5512 int32
	_ = v5512
	var v5517 int32
	_ = v5517
	var v5520 int32
	_ = v5520
	var v5521 int32
	_ = v5521
	var v5525 int32
	_ = v5525
	var v5526 int32
	_ = v5526
	var v5528 int32
	_ = v5528
	var v5533 int32
	_ = v5533
	var v5536 int32
	_ = v5536
	var v5537 int32
	_ = v5537
	var v5538 int32
	_ = v5538
	var v5543 int32
	_ = v5543
	var v5544 int32
	_ = v5544
	var v5547 int32
	_ = v5547
	var v5548 int32
	_ = v5548
	var v5550 int32
	_ = v5550
	var v5555 int32
	_ = v5555
	var v5557 int32
	_ = v5557
	var v5558 int32
	_ = v5558
	var v5562 int32
	_ = v5562
	var v5564 int32
	_ = v5564
	var v5567 int32
	_ = v5567
	var v5568 int32
	_ = v5568
	var v5571 int32
	_ = v5571
	var v5574 int32
	_ = v5574
	var v5581 int32
	_ = v5581
	var v5586 int32
	_ = v5586
	var v5587 int32
	_ = v5587
	var v5589 int32
	_ = v5589
	var v5592 int32
	_ = v5592
	var v5593 int32
	_ = v5593
	var v5594 int32
	_ = v5594
	var v5595 int32
	_ = v5595
	var v5597 int32
	_ = v5597
	var v5600 int32
	_ = v5600
	var v5601 int32
	_ = v5601
	var v5607 int32
	_ = v5607
	var v5610 int32
	_ = v5610
	var v5613 int32
	_ = v5613
	var v5617 int32
	_ = v5617
	var v5620 int32
	_ = v5620
	var v5621 int32
	_ = v5621
	var v5627 int32
	_ = v5627
	var v5630 int32
	_ = v5630
	var v5631 int32
	_ = v5631
	var v5632 int32
	_ = v5632
	var v5637 int32
	_ = v5637
	var v5638 int32
	_ = v5638
	var v5639 int32
	_ = v5639
	var v5640 int32
	_ = v5640
	var v5643 int32
	_ = v5643
	var v5647 int32
	_ = v5647
	var v5652 int32
	_ = v5652
	var v5654 int32
	_ = v5654
	var v5660 int32
	_ = v5660
	var v5661 int32
	_ = v5661
	var v5671 int32
	_ = v5671
	var v5672 int32
	_ = v5672
	var v5673 int32
	_ = v5673
	var v5676 int32
	_ = v5676
	var v5679 int32
	_ = v5679
	var v5680 int32
	_ = v5680
	var v5684 int32
	_ = v5684
	var v5685 int32
	_ = v5685
	var v5688 int32
	_ = v5688
	var v5689 int32
	_ = v5689
	var v5692 int32
	_ = v5692
	var v5699 int32
	_ = v5699
	var v5700 int32
	_ = v5700
	var v5703 int32
	_ = v5703
	var v5705 int32
	_ = v5705
	var v5706 int32
	_ = v5706
	var v5710 int32
	_ = v5710
	var v5711 int32
	_ = v5711
	var v5714 int32
	_ = v5714
	var v5719 int32
	_ = v5719
	var v5720 int32
	_ = v5720
	var v5723 int32
	_ = v5723
	var v5724 int32
	_ = v5724
	var v5726 int32
	_ = v5726
	var v5731 int32
	_ = v5731
	var v5733 int32
	_ = v5733
	var v5734 int32
	_ = v5734
	var v5738 int32
	_ = v5738
	var v5740 int32
	_ = v5740
	var v5743 int32
	_ = v5743
	var v5744 int32
	_ = v5744
	var v5747 int32
	_ = v5747
	var v5750 int32
	_ = v5750
	var v5757 int32
	_ = v5757
	var v5762 int32
	_ = v5762
	var v5763 int32
	_ = v5763
	var v5765 int32
	_ = v5765
	var v5770 int32
	_ = v5770
	var v5775 int32
	_ = v5775
	var v5776 int32
	_ = v5776
	var v5782 int32
	_ = v5782
	var v5783 int32
	_ = v5783
	var v5788 int32
	_ = v5788
	var v5789 int32
	_ = v5789
	var v5792 int32
	_ = v5792
	var v5793 int32
	_ = v5793
	var v5798 int32
	_ = v5798
	var v5804 int32
	_ = v5804
	var v5805 int32
	_ = v5805
	var v5814 int32
	_ = v5814
	var v5816 int32
	_ = v5816
	var v5820 int32
	_ = v5820
	var v5828 int32
	_ = v5828
	var v5829 int32
	_ = v5829
	var v5832 int32
	_ = v5832
	var v5835 int32
	_ = v5835
	var v5836 int32
	_ = v5836
	var v5840 int32
	_ = v5840
	var v5841 int32
	_ = v5841
	var v5844 int32
	_ = v5844
	var v5845 int32
	_ = v5845
	var v5848 int32
	_ = v5848
	var v5855 int32
	_ = v5855
	var v5856 int32
	_ = v5856
	var v5859 int32
	_ = v5859
	var v5861 int32
	_ = v5861
	var v5862 int32
	_ = v5862
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
	var v5894 int32
	_ = v5894
	var v5896 int32
	_ = v5896
	var v5899 int32
	_ = v5899
	var v5900 int32
	_ = v5900
	var v5903 int32
	_ = v5903
	var v5906 int32
	_ = v5906
	var v5913 int32
	_ = v5913
	var v5918 int32
	_ = v5918
	var v5919 int32
	_ = v5919
	var v5921 int32
	_ = v5921
	var v5928 int32
	_ = v5928
	var v5929 int32
	_ = v5929
	var v5937 int32
	_ = v5937
	var v5939 int32
	_ = v5939
	var v5951 int32
	_ = v5951
	var v5952 int32
	_ = v5952
	var v5957 int32
	_ = v5957
	var v5960 int32
	_ = v5960
	var v5969 int32
	_ = v5969
	var v5970 int32
	_ = v5970
	var v5972 int32
	_ = v5972
	var v6001 int32
	_ = v6001
	var v6004 int32
	_ = v6004
	var v6013 int32
	_ = v6013
	var v6014 int32
	_ = v6014
	var v6015 int32
	_ = v6015
	var v6016 int32
	_ = v6016
	var v6017 int32
	_ = v6017
	var v6019 int32
	_ = v6019
	var v6053 int32
	_ = v6053
	var v6054 int32
	_ = v6054
	var v6056 int32
	_ = v6056
	var v6057 int32
	_ = v6057
	var v6060 int32
	_ = v6060
	var v6065 int32
	_ = v6065
	var v6066 int32
	_ = v6066
	var v6069 int32
	_ = v6069
	var v6070 int32
	_ = v6070
	var v6072 int32
	_ = v6072
	var v6077 int32
	_ = v6077
	var v6079 int32
	_ = v6079
	var v6080 int32
	_ = v6080
	var v6084 int32
	_ = v6084
	var v6086 int32
	_ = v6086
	var v6089 int32
	_ = v6089
	var v6090 int32
	_ = v6090
	var v6093 int32
	_ = v6093
	var v6096 int32
	_ = v6096
	var v6103 int32
	_ = v6103
	var v6108 int32
	_ = v6108
	var v6109 int32
	_ = v6109
	var v6110 int32
	_ = v6110
	var v6111 int32
	_ = v6111
	var v6116 int32
	_ = v6116
	var v6123 int32
	_ = v6123
	var v6124 int32
	_ = v6124
	var v6130 int32
	_ = v6130
	var v6133 int32
	_ = v6133
	var v6137 int32
	_ = v6137
	var v6141 int32
	_ = v6141
	var v6142 int32
	_ = v6142
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6157 int32
	_ = v6157
	var v6158 int32
	_ = v6158
	var v6162 int32
	_ = v6162
	var v6163 int32
	_ = v6163
	var v6167 int32
	_ = v6167
	var v6172 int32
	_ = v6172
	var v6176 int32
	_ = v6176
	var v6177 int32
	_ = v6177
	var v6178 int32
	_ = v6178
	var v6179 int32
	_ = v6179
	var v6184 int32
	_ = v6184
	var v6185 int32
	_ = v6185
	var v6190 int32
	_ = v6190
	var v6191 int32
	_ = v6191
	var v6197 int32
	_ = v6197
	var v6198 int32
	_ = v6198
	var v6203 int32
	_ = v6203
	var v6204 int32
	_ = v6204
	var v6207 int32
	_ = v6207
	var v6208 int32
	_ = v6208
	var v6213 int32
	_ = v6213
	var v6219 int32
	_ = v6219
	var v6220 int32
	_ = v6220
	var v6229 int32
	_ = v6229
	var v6231 int32
	_ = v6231
	var v6235 int32
	_ = v6235
	var v6243 int32
	_ = v6243
	var v6244 int32
	_ = v6244
	var v6251 int32
	_ = v6251
	var v6252 int32
	_ = v6252
	var v6255 int32
	_ = v6255
	var v6256 int32
	_ = v6256
	var v6261 int32
	_ = v6261
	var v6262 int32
	_ = v6262
	var v6265 int32
	_ = v6265
	var v6266 int32
	_ = v6266
	var v6268 int32
	_ = v6268
	var v6273 int32
	_ = v6273
	var v6275 int32
	_ = v6275
	var v6276 int32
	_ = v6276
	var v6280 int32
	_ = v6280
	var v6282 int32
	_ = v6282
	var v6285 int32
	_ = v6285
	var v6286 int32
	_ = v6286
	var v6289 int32
	_ = v6289
	var v6292 int32
	_ = v6292
	var v6299 int32
	_ = v6299
	var v6304 int32
	_ = v6304
	var v6305 int32
	_ = v6305
	var v6307 int32
	_ = v6307
	var v6309 int64
	_ = v6309
	var v6311 int32
	_ = v6311
	var v6321 int32
	_ = v6321
	var v6322 int32
	_ = v6322
	var v6325 int32
	_ = v6325
	var v6337 int32
	_ = v6337
	var v6342 int32
	_ = v6342
	var v6343 int32
	_ = v6343
	var v6346 int32
	_ = v6346
	var v6349 int32
	_ = v6349
	var v6350 int32
	_ = v6350
	var v6354 int32
	_ = v6354
	var v6355 int32
	_ = v6355
	var v6358 int32
	_ = v6358
	var v6359 int32
	_ = v6359
	var v6362 int32
	_ = v6362
	var v6369 int32
	_ = v6369
	var v6370 int32
	_ = v6370
	var v6374 int32
	_ = v6374
	var v6377 int32
	_ = v6377
	var v6378 int32
	_ = v6378
	var v6382 int32
	_ = v6382
	var v6383 int32
	_ = v6383
	var v6386 int32
	_ = v6386
	var v6387 int32
	_ = v6387
	var v6390 int32
	_ = v6390
	var v6397 int32
	_ = v6397
	var v6398 int32
	_ = v6398
	var v6402 int32
	_ = v6402
	var v6405 int32
	_ = v6405
	var v6408 int32
	_ = v6408
	var v6409 int32
	_ = v6409
	var v6413 int32
	_ = v6413
	var v6414 int32
	_ = v6414
	var v6417 int32
	_ = v6417
	var v6418 int32
	_ = v6418
	var v6421 int32
	_ = v6421
	var v6428 int32
	_ = v6428
	var v6429 int32
	_ = v6429
	var v6433 int32
	_ = v6433
	var v6436 int32
	_ = v6436
	var v6437 int32
	_ = v6437
	var v6441 int32
	_ = v6441
	var v6442 int32
	_ = v6442
	var v6445 int32
	_ = v6445
	var v6446 int32
	_ = v6446
	var v6449 int32
	_ = v6449
	var v6456 int32
	_ = v6456
	var v6457 int32
	_ = v6457
	var v6461 int32
	_ = v6461
	var v6464 int32
	_ = v6464
	var v6467 int32
	_ = v6467
	var v6468 int32
	_ = v6468
	var v6472 int32
	_ = v6472
	var v6473 int32
	_ = v6473
	var v6476 int32
	_ = v6476
	var v6477 int32
	_ = v6477
	var v6480 int32
	_ = v6480
	var v6487 int32
	_ = v6487
	var v6488 int32
	_ = v6488
	var v6492 int32
	_ = v6492
	var v6495 int32
	_ = v6495
	var v6496 int32
	_ = v6496
	var v6500 int32
	_ = v6500
	var v6501 int32
	_ = v6501
	var v6504 int32
	_ = v6504
	var v6505 int32
	_ = v6505
	var v6508 int32
	_ = v6508
	var v6515 int32
	_ = v6515
	var v6516 int32
	_ = v6516
	var v6521 int32
	_ = v6521
	var v6527 int32
	_ = v6527
	var v6528 int32
	_ = v6528
	var v6530 int32
	_ = v6530
	var v6541 int32
	_ = v6541
	var v6542 int32
	_ = v6542
	var v6548 int32
	_ = v6548
	var v6549 int32
	_ = v6549
	var v6560 int32
	_ = v6560
	var v6563 int32
	_ = v6563
	var v6564 int32
	_ = v6564
	var v6565 int32
	_ = v6565
	var v6568 int32
	_ = v6568
	var v6571 int32
	_ = v6571
	var v6572 int32
	_ = v6572
	var v6576 int32
	_ = v6576
	var v6577 int32
	_ = v6577
	var v6580 int32
	_ = v6580
	var v6581 int32
	_ = v6581
	var v6584 int32
	_ = v6584
	var v6591 int32
	_ = v6591
	var v6592 int32
	_ = v6592
	var v6600 int32
	_ = v6600
	var v6601 int32
	_ = v6601
	var v6604 int32
	_ = v6604
	var v6612 int32
	_ = v6612
	var v6617 int32
	_ = v6617
	var v6621 int32
	_ = v6621
	var v6626 int32
	_ = v6626
	var v6628 int32
	_ = v6628
	var v6632 int32
	_ = v6632
	var v6638 int32
	_ = v6638
	var v6641 int32
	_ = v6641
	var v6647 int32
	_ = v6647
	var v6651 int32
	_ = v6651
	var v6653 int32
	_ = v6653
	var v6661 int32
	_ = v6661
	var v6664 int32
	_ = v6664
	var v6668 int32
	_ = v6668
	var v6670 int32
	_ = v6670
	var v6671 int64
	_ = v6671
	var v6679 int32
	_ = v6679
	var v6683 int32
	_ = v6683
	var v6687 int32
	_ = v6687
	var v6693 int32
	_ = v6693
	var v6697 int32
	_ = v6697
	var v6698 int32
	_ = v6698
	var v6705 int32
	_ = v6705
	var v6706 int32
	_ = v6706
	var v6707 int32
	_ = v6707
	var v6711 int32
	_ = v6711
	var v6714 int32
	_ = v6714
	var v6718 int32
	_ = v6718
	var v6719 int32
	_ = v6719
	var v6727 int32
	_ = v6727
	var v6733 int32
	_ = v6733
	var v6735 int32
	_ = v6735
	var v6739 int32
	_ = v6739
	var v6747 int32
	_ = v6747
	var v6751 int32
	_ = v6751
	var v6756 int32
	_ = v6756
	var v6762 int32
	_ = v6762
	var v6763 int32
	_ = v6763
	var v6767 int32
	_ = v6767
	var v6771 int32
	_ = v6771
	var v6772 int32
	_ = v6772
	var v6773 int32
	_ = v6773
	var v6779 int32
	_ = v6779
	var v6780 int32
	_ = v6780
	var v6783 int32
	_ = v6783
	var v6786 int32
	_ = v6786
	var v6787 int32
	_ = v6787
	var v6794 int32
	_ = v6794
	var v6795 int32
	_ = v6795
	var v6806 int32
	_ = v6806
	var v6836 int32
	_ = v6836
	var v6845 int32
	_ = v6845
	var v6846 int32
	_ = v6846
	var v6847 int32
	_ = v6847
	var v6848 int32
	_ = v6848
	var v6849 int32
	_ = v6849
	var v6851 int32
	_ = v6851
	var v6857 int32
	_ = v6857
	var v6911 int32
	_ = v6911
	var v6933 int32
	_ = v6933
	var v6934 int32
	_ = v6934
	var v6939 int32
	_ = v6939
	var v6940 int32
	_ = v6940
	var v6941 int32
	_ = v6941
	var v6942 int32
	_ = v6942
	var v6945 int32
	_ = v6945
	var v6946 int32
	_ = v6946
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
	var v7092 int32
	_ = v7092
	var v7093 int32
	_ = v7093
	var v7097 int32
	_ = v7097
	var v7098 int32
	_ = v7098
	var v7101 int32
	_ = v7101
	var v7102 int32
	_ = v7102
	var v7105 int32
	_ = v7105
	var v7112 int32
	_ = v7112
	var v7113 int32
	_ = v7113
	var v7117 int32
	_ = v7117
	var v7120 int32
	_ = v7120
	var v7121 int32
	_ = v7121
	var v7125 int32
	_ = v7125
	var v7126 int32
	_ = v7126
	var v7129 int32
	_ = v7129
	var v7130 int32
	_ = v7130
	var v7133 int32
	_ = v7133
	var v7140 int32
	_ = v7140
	var v7141 int32
	_ = v7141
	var v7145 int32
	_ = v7145
	var v7148 int32
	_ = v7148
	var v7149 int32
	_ = v7149
	var v7153 int32
	_ = v7153
	var v7154 int32
	_ = v7154
	var v7157 int32
	_ = v7157
	var v7158 int32
	_ = v7158
	var v7161 int32
	_ = v7161
	var v7168 int32
	_ = v7168
	var v7169 int32
	_ = v7169
	var v7173 int32
	_ = v7173
	var v7176 int32
	_ = v7176
	var v7177 int32
	_ = v7177
	var v7181 int32
	_ = v7181
	var v7182 int32
	_ = v7182
	var v7185 int32
	_ = v7185
	var v7186 int32
	_ = v7186
	var v7189 int32
	_ = v7189
	var v7196 int32
	_ = v7196
	var v7197 int32
	_ = v7197
	var v7201 int32
	_ = v7201
	var v7208 int32
	_ = v7208
	var v7209 int32
	_ = v7209
	var v7218 int32
	_ = v7218
	var v7221 int32
	_ = v7221
	var v7230 int32
	_ = v7230
	var v7231 int32
	_ = v7231
	var v7233 int32
	_ = v7233
	var v7234 int32
	_ = v7234
	var v7235 int32
	_ = v7235
	var v7264 int32
	_ = v7264
	var v7268 int32
	_ = v7268
	var v7271 int32
	_ = v7271
	var v7289 int32
	_ = v7289
	var v7292 int32
	_ = v7292
	var v7293 int32
	_ = v7293
	var v7298 int32
	_ = v7298
	var v7300 int32
	_ = v7300
	var v7301 int32
	_ = v7301
	var v7303 int32
	_ = v7303
	var v7305 int32
	_ = v7305
	var v7308 int32
	_ = v7308
	var v7310 int32
	_ = v7310
	var v7338 int32
	_ = v7338
	var v7339 int32
	_ = v7339
	var v7342 int32
	_ = v7342
	var v7343 int32
	_ = v7343
	var v7348 int32
	_ = v7348
	var v7349 int32
	_ = v7349
	var v7352 int32
	_ = v7352
	var v7353 int32
	_ = v7353
	var v7355 int32
	_ = v7355
	var v7360 int32
	_ = v7360
	var v7362 int32
	_ = v7362
	var v7363 int32
	_ = v7363
	var v7367 int32
	_ = v7367
	var v7369 int32
	_ = v7369
	var v7372 int32
	_ = v7372
	var v7373 int32
	_ = v7373
	var v7376 int32
	_ = v7376
	var v7379 int32
	_ = v7379
	var v7386 int32
	_ = v7386
	var v7391 int32
	_ = v7391
	var v7392 int32
	_ = v7392
	var v7393 int32
	_ = v7393
	var v7394 int32
	_ = v7394
	var v7399 int32
	_ = v7399
	var v7411 int32
	_ = v7411
	var v7412 int32
	_ = v7412
	var v7414 int32
	_ = v7414
	var v7418 int32
	_ = v7418
	var v7422 int32
	_ = v7422
	var v7430 int32
	_ = v7430
	var v7431 int32
	_ = v7431
	var v7433 int32
	_ = v7433
	var v7438 int32
	_ = v7438
	var v7442 int32
	_ = v7442
	var v7446 int32
	_ = v7446
	var v7449 int32
	_ = v7449
	var v7455 int32
	_ = v7455
	var v7456 int32
	_ = v7456
	var v7459 int32
	_ = v7459
	var v7465 int32
	_ = v7465
	var v7466 int32
	_ = v7466
	var v7469 int32
	_ = v7469
	var v7475 int32
	_ = v7475
	var v7476 int32
	_ = v7476
	var v7482 int32
	_ = v7482
	var v7483 int32
	_ = v7483
	var v7489 int32
	_ = v7489
	var v7497 int32
	_ = v7497
	var v7502 int32
	_ = v7502
	var v7503 int32
	_ = v7503
	var v7509 int32
	_ = v7509
	var v7510 int32
	_ = v7510
	var v7516 int32
	_ = v7516
	var v7524 int32
	_ = v7524
	var v7530 int32
	_ = v7530
	var v7531 int32
	_ = v7531
	var v7538 int32
	_ = v7538
	var v7547 int32
	_ = v7547
	var v7548 int32
	_ = v7548
	var v7550 int32
	_ = v7550
	var v7551 int32
	_ = v7551
	var v7554 int32
	_ = v7554
	var v7555 int32
	_ = v7555
	var v7560 int32
	_ = v7560
	var v7561 int32
	_ = v7561
	var v7564 int32
	_ = v7564
	var v7565 int32
	_ = v7565
	var v7567 int32
	_ = v7567
	var v7572 int32
	_ = v7572
	var v7574 int32
	_ = v7574
	var v7575 int32
	_ = v7575
	var v7579 int32
	_ = v7579
	var v7581 int32
	_ = v7581
	var v7584 int32
	_ = v7584
	var v7585 int32
	_ = v7585
	var v7588 int32
	_ = v7588
	var v7591 int32
	_ = v7591
	var v7598 int32
	_ = v7598
	var v7603 int32
	_ = v7603
	var v7604 int32
	_ = v7604
	var v7606 int32
	_ = v7606
	var v7610 int32
	_ = v7610
	var v7618 int32
	_ = v7618
	var v7623 int32
	_ = v7623
	var v7654 int32
	_ = v7654
	var v7655 int32
	_ = v7655
	var v7658 int32
	_ = v7658
	var v7665 int32
	_ = v7665
	var v7670 int32
	_ = v7670
	var v7671 int32
	_ = v7671
	var v7673 int32
	_ = v7673
	var v7703 int32
	_ = v7703
	var v7712 int32
	_ = v7712
	var v7713 int32
	_ = v7713
	var v7714 int32
	_ = v7714
	var v7715 int32
	_ = v7715
	var v7716 int32
	_ = v7716
	var v7718 int32
	_ = v7718
	var v7723 int32
	_ = v7723
	var v7724 int32
	_ = v7724
	var v7729 int32
	_ = v7729
	var v7730 int32
	_ = v7730
	var v7735 int32
	_ = v7735
	var v7736 int32
	_ = v7736
	var v7739 int32
	_ = v7739
	var v7740 int32
	_ = v7740
	var v7742 int32
	_ = v7742
	var v7747 int32
	_ = v7747
	var v7749 int32
	_ = v7749
	var v7750 int32
	_ = v7750
	var v7754 int32
	_ = v7754
	var v7756 int32
	_ = v7756
	var v7759 int32
	_ = v7759
	var v7760 int32
	_ = v7760
	var v7763 int32
	_ = v7763
	var v7766 int32
	_ = v7766
	var v7773 int32
	_ = v7773
	var v7778 int32
	_ = v7778
	var v7779 int32
	_ = v7779
	var v7781 int32
	_ = v7781
	var v7784 int32
	_ = v7784
	var v7785 int32
	_ = v7785
	var v7789 int32
	_ = v7789
	var v7790 int32
	_ = v7790
	var v7793 int32
	_ = v7793
	var v7798 int32
	_ = v7798
	var v7799 int32
	_ = v7799
	var v7806 int32
	_ = v7806
	var v7807 int32
	_ = v7807
	var v7810 int32
	_ = v7810
	var v7813 int32
	_ = v7813
	var v7816 int32
	_ = v7816
	var v7818 int32
	_ = v7818
	var v7823 int32
	_ = v7823
	var v7824 int32
	_ = v7824
	var v7829 int32
	_ = v7829
	var v7830 int32
	_ = v7830
	var v7833 int32
	_ = v7833
	var v7836 int32
	_ = v7836
	var v7837 int32
	_ = v7837
	var v7841 int32
	_ = v7841
	var v7842 int32
	_ = v7842
	var v7845 int32
	_ = v7845
	var v7846 int32
	_ = v7846
	var v7849 int32
	_ = v7849
	var v7856 int32
	_ = v7856
	var v7857 int32
	_ = v7857
	var v7861 int32
	_ = v7861
	var v7864 int32
	_ = v7864
	var v7865 int32
	_ = v7865
	var v7868 int32
	_ = v7868
	var v7869 int32
	_ = v7869
	var v7873 int32
	_ = v7873
	var v7874 int32
	_ = v7874
	var v7877 int32
	_ = v7877
	var v7878 int32
	_ = v7878
	var v7881 int32
	_ = v7881
	var v7888 int32
	_ = v7888
	var v7889 int32
	_ = v7889
	var v7892 int32
	_ = v7892
	var v7893 int32
	_ = v7893
	var v7900 int32
	_ = v7900
	var v7901 int32
	_ = v7901
	var v7902 int32
	_ = v7902
	var v7910 int32
	_ = v7910
	var v7911 int32
	_ = v7911
	var v7916 int32
	_ = v7916
	var v7919 int32
	_ = v7919
	var v7928 int32
	_ = v7928
	var v7929 int32
	_ = v7929
	var v7931 int32
	_ = v7931
	var v7960 int32
	_ = v7960
	var v7963 int32
	_ = v7963
	var v7972 int32
	_ = v7972
	var v7973 int32
	_ = v7973
	var v7974 int32
	_ = v7974
	var v7975 int32
	_ = v7975
	var v7976 int32
	_ = v7976
	var v7978 int32
	_ = v7978
	var v8011 int32
	_ = v8011
	var v8013 int32
	_ = v8013
	var v8025 int32
	_ = v8025
	var v8026 int32
	_ = v8026
	var v8034 int32
	_ = v8034
	var v8035 int32
	_ = v8035
	var v8040 int32
	_ = v8040
	var v8049 int32
	_ = v8049
	var v8054 int32
	_ = v8054
	var v8055 int32
	_ = v8055
	var v8058 int32
	_ = v8058
	var v8063 int32
	_ = v8063
	var v8064 int32
	_ = v8064
	var v8069 int32
	_ = v8069
	var v8070 int32
	_ = v8070
	var v8073 int32
	_ = v8073
	var v8074 int32
	_ = v8074
	var v8076 int32
	_ = v8076
	var v8081 int32
	_ = v8081
	var v8083 int32
	_ = v8083
	var v8084 int32
	_ = v8084
	var v8088 int32
	_ = v8088
	var v8090 int32
	_ = v8090
	var v8093 int32
	_ = v8093
	var v8094 int32
	_ = v8094
	var v8097 int32
	_ = v8097
	var v8100 int32
	_ = v8100
	var v8107 int32
	_ = v8107
	var v8111 int32
	_ = v8111
	var v8115 int32
	_ = v8115
	var v8116 int32
	_ = v8116
	var v8117 int32
	_ = v8117
	var v8123 int32
	_ = v8123
	var v8126 int32
	_ = v8126
	var v8127 int32
	_ = v8127
	var v8132 int32
	_ = v8132
	var v8133 int32
	_ = v8133
	var v8136 int32
	_ = v8136
	var v8137 int32
	_ = v8137
	var v8139 int32
	_ = v8139
	var v8144 int32
	_ = v8144
	var v8146 int32
	_ = v8146
	var v8147 int32
	_ = v8147
	var v8151 int32
	_ = v8151
	var v8153 int32
	_ = v8153
	var v8156 int32
	_ = v8156
	var v8157 int32
	_ = v8157
	var v8160 int32
	_ = v8160
	var v8163 int32
	_ = v8163
	var v8170 int32
	_ = v8170
	var v8176 int32
	_ = v8176
	var v8177 int32
	_ = v8177
	var v8178 int32
	_ = v8178
	var v8182 int32
	_ = v8182
	var v8185 int32
	_ = v8185
	var v8186 int32
	_ = v8186
	var v8190 int32
	_ = v8190
	var v8191 int32
	_ = v8191
	var v8193 int32
	_ = v8193
	var v8195 int32
	_ = v8195
	var v8206 int32
	_ = v8206
	var v8207 int32
	_ = v8207
	var v8215 int32
	_ = v8215
	var v8216 int32
	_ = v8216
	var v8217 int32
	_ = v8217
	var v8220 int32
	_ = v8220
	var v8223 int32
	_ = v8223
	var v8224 int32
	_ = v8224
	var v8228 int32
	_ = v8228
	var v8229 int32
	_ = v8229
	var v8232 int32
	_ = v8232
	var v8233 int32
	_ = v8233
	var v8236 int32
	_ = v8236
	var v8243 int32
	_ = v8243
	var v8244 int32
	_ = v8244
	var v8248 int32
	_ = v8248
	var v8251 int32
	_ = v8251
	var v8252 int32
	_ = v8252
	var v8256 int32
	_ = v8256
	var v8257 int32
	_ = v8257
	var v8260 int32
	_ = v8260
	var v8261 int32
	_ = v8261
	var v8264 int32
	_ = v8264
	var v8271 int32
	_ = v8271
	var v8272 int32
	_ = v8272
	var v8277 int32
	_ = v8277
	var v8280 int32
	_ = v8280
	var v8281 int32
	_ = v8281
	var v8285 int32
	_ = v8285
	var v8286 int32
	_ = v8286
	var v8289 int32
	_ = v8289
	var v8290 int32
	_ = v8290
	var v8293 int32
	_ = v8293
	var v8300 int32
	_ = v8300
	var v8301 int32
	_ = v8301
	var v8306 int32
	_ = v8306
	var v8309 int32
	_ = v8309
	var v8310 int32
	_ = v8310
	var v8314 int32
	_ = v8314
	var v8315 int32
	_ = v8315
	var v8318 int32
	_ = v8318
	var v8319 int32
	_ = v8319
	var v8322 int32
	_ = v8322
	var v8329 int32
	_ = v8329
	var v8330 int32
	_ = v8330
	var v8335 int32
	_ = v8335
	var v8338 int32
	_ = v8338
	var v8339 int32
	_ = v8339
	var v8343 int32
	_ = v8343
	var v8344 int32
	_ = v8344
	var v8347 int32
	_ = v8347
	var v8348 int32
	_ = v8348
	var v8351 int32
	_ = v8351
	var v8358 int32
	_ = v8358
	var v8359 int32
	_ = v8359
	var v8362 int32
	_ = v8362
	var v8366 int32
	_ = v8366
	var v8369 int32
	_ = v8369
	var v8377 int32
	_ = v8377
	var v8378 int32
	_ = v8378
	var v8380 int32
	_ = v8380
	var v8383 int32
	_ = v8383
	var v8384 int32
	_ = v8384
	var v8388 int32
	_ = v8388
	var v8389 int32
	_ = v8389
	var v8392 int32
	_ = v8392
	var v8393 int32
	_ = v8393
	var v8396 int32
	_ = v8396
	var v8403 int32
	_ = v8403
	var v8404 int32
	_ = v8404
	var v8411 int32
	_ = v8411
	var v8414 int32
	_ = v8414
	var v8422 int32
	_ = v8422
	var v8423 int32
	_ = v8423
	var v8425 int32
	_ = v8425
	var v8428 int32
	_ = v8428
	var v8431 int32
	_ = v8431
	var v8432 int32
	_ = v8432
	var v8436 int32
	_ = v8436
	var v8437 int32
	_ = v8437
	var v8440 int32
	_ = v8440
	var v8441 int32
	_ = v8441
	var v8444 int32
	_ = v8444
	var v8451 int32
	_ = v8451
	var v8452 int32
	_ = v8452
	var v8455 int32
	_ = v8455
	var v8459 int32
	_ = v8459
	var v8462 int32
	_ = v8462
	var v8463 int32
	_ = v8463
	var v8467 int32
	_ = v8467
	var v8468 int32
	_ = v8468
	var v8471 int32
	_ = v8471
	var v8472 int32
	_ = v8472
	var v8475 int32
	_ = v8475
	var v8482 int32
	_ = v8482
	var v8483 int32
	_ = v8483
	var v8493 int32
	_ = v8493
	var v8494 int32
	_ = v8494
	var v8497 int32
	_ = v8497
	var v8500 int32
	_ = v8500
	var v8501 int32
	_ = v8501
	var v8505 int32
	_ = v8505
	var v8506 int32
	_ = v8506
	var v8509 int32
	_ = v8509
	var v8510 int32
	_ = v8510
	var v8513 int32
	_ = v8513
	var v8520 int32
	_ = v8520
	var v8521 int32
	_ = v8521
	var v8533 int32
	_ = v8533
	var v8541 int32
	_ = v8541
	var v8542 int32
	_ = v8542
	var v8549 int32
	_ = v8549
	var v8552 int32
	_ = v8552
	var v8555 int32
	_ = v8555
	var v8563 int32
	_ = v8563
	var v8564 int32
	_ = v8564
	var v8565 int32
	_ = v8565
	var v8569 int32
	_ = v8569
	var v8572 int32
	_ = v8572
	var v8579 int32
	_ = v8579
	var v8580 int32
	_ = v8580
	var v8588 int32
	_ = v8588
	var v8591 int32
	_ = v8591
	var v8592 int32
	_ = v8592
	var v8597 int32
	_ = v8597
	var v8598 int32
	_ = v8598
	var v8603 int32
	_ = v8603
	var v8604 int32
	_ = v8604
	var v8607 int32
	_ = v8607
	var v8608 int32
	_ = v8608
	var v8610 int32
	_ = v8610
	var v8615 int32
	_ = v8615
	var v8617 int32
	_ = v8617
	var v8618 int32
	_ = v8618
	var v8622 int32
	_ = v8622
	var v8624 int32
	_ = v8624
	var v8627 int32
	_ = v8627
	var v8628 int32
	_ = v8628
	var v8631 int32
	_ = v8631
	var v8634 int32
	_ = v8634
	var v8641 int32
	_ = v8641
	var v8646 int32
	_ = v8646
	var v8647 int32
	_ = v8647
	var v8649 int32
	_ = v8649
	var v8654 int32
	_ = v8654
	var v8655 int32
	_ = v8655
	var v8661 int32
	_ = v8661
	var v8662 int32
	_ = v8662
	var v8667 int32
	_ = v8667
	var v8668 int32
	_ = v8668
	var v8673 int32
	_ = v8673
	var v8674 int32
	_ = v8674
	var v8677 int32
	_ = v8677
	var v8678 int32
	_ = v8678
	var v8680 int32
	_ = v8680
	var v8685 int32
	_ = v8685
	var v8687 int32
	_ = v8687
	var v8688 int32
	_ = v8688
	var v8692 int32
	_ = v8692
	var v8694 int32
	_ = v8694
	var v8697 int32
	_ = v8697
	var v8698 int32
	_ = v8698
	var v8701 int32
	_ = v8701
	var v8704 int32
	_ = v8704
	var v8711 int32
	_ = v8711
	var v8716 int32
	_ = v8716
	var v8717 int32
	_ = v8717
	var v8719 int32
	_ = v8719
	var v8724 int32
	_ = v8724
	var v8730 int32
	_ = v8730
	var v8731 int32
	_ = v8731
	var v8736 int32
	_ = v8736
	var v8737 int32
	_ = v8737
	var v8742 int32
	_ = v8742
	var v8743 int32
	_ = v8743
	var v8746 int32
	_ = v8746
	var v8747 int32
	_ = v8747
	var v8749 int32
	_ = v8749
	var v8754 int32
	_ = v8754
	var v8756 int32
	_ = v8756
	var v8757 int32
	_ = v8757
	var v8761 int32
	_ = v8761
	var v8763 int32
	_ = v8763
	var v8766 int32
	_ = v8766
	var v8767 int32
	_ = v8767
	var v8770 int32
	_ = v8770
	var v8773 int32
	_ = v8773
	var v8780 int32
	_ = v8780
	var v8785 int32
	_ = v8785
	var v8786 int32
	_ = v8786
	var v8788 int32
	_ = v8788
	var v8793 int32
	_ = v8793
	var v8804 int32
	_ = v8804
	var v8805 int32
	_ = v8805
	var v8806 int32
	_ = v8806
	var v8807 int32
	_ = v8807
	var v8810 int32
	_ = v8810
	var v8811 int32
	_ = v8811
	var v8812 int32
	_ = v8812
	var v8816 int32
	_ = v8816
	var v8818 int32
	_ = v8818
	var v8819 int32
	_ = v8819
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
	var v8872 int32
	_ = v8872
	var v8873 int32
	_ = v8873
	var v8875 int32
	_ = v8875
	var v8876 int32
	_ = v8876
	var v8881 int32
	_ = v8881
	var v8883 int32
	_ = v8883
	var v8884 int32
	_ = v8884
	var v8886 int32
	_ = v8886
	var v8887 int32
	_ = v8887
	var v8888 int32
	_ = v8888
	var v8890 int32
	_ = v8890
	var v8896 int32
	_ = v8896
	var v8897 int32
	_ = v8897
	var v8899 int32
	_ = v8899
	var v8900 int32
	_ = v8900
	var v8902 int32
	_ = v8902
	var v8903 int32
	_ = v8903
	var v8904 int32
	_ = v8904
	var v8906 int32
	_ = v8906
	var v8911 int32
	_ = v8911
	var v8912 int32
	_ = v8912
	var v8917 int32
	_ = v8917
	var v8918 int32
	_ = v8918
	var v8919 int32
	_ = v8919
	var v8920 int32
	_ = v8920
	var v8922 int32
	_ = v8922
	var v8928 int32
	_ = v8928
	var v8929 int32
	_ = v8929
	var v8932 int32
	_ = v8932
	var v8933 int32
	_ = v8933
	var v8936 int32
	_ = v8936
	var v8937 int32
	_ = v8937
	var v8942 int32
	_ = v8942
	var v8943 int32
	_ = v8943
	var v8946 int32
	_ = v8946
	var v8947 int32
	_ = v8947
	var v8949 int32
	_ = v8949
	var v8954 int32
	_ = v8954
	var v8956 int32
	_ = v8956
	var v8957 int32
	_ = v8957
	var v8961 int32
	_ = v8961
	var v8963 int32
	_ = v8963
	var v8966 int32
	_ = v8966
	var v8967 int32
	_ = v8967
	var v8970 int32
	_ = v8970
	var v8973 int32
	_ = v8973
	var v8980 int32
	_ = v8980
	var v8986 int32
	_ = v8986
	var v8988 int32
	_ = v8988
	var v8992 int32
	_ = v8992
	var v8993 int32
	_ = v8993
	var v8997 int32
	_ = v8997
	var v9018 int32
	_ = v9018
	var v9019 int32
	_ = v9019
	var v9021 int32
	_ = v9021
	var v9023 int32
	_ = v9023
	var v9025 int32
	_ = v9025
	var v9026 int32
	_ = v9026
	var v9029 int32
	_ = v9029
	var v9030 int32
	_ = v9030
	var v9034 int32
	_ = v9034
	var v9035 int32
	_ = v9035
	var v9038 int32
	_ = v9038
	var v9039 int32
	_ = v9039
	var v9042 int32
	_ = v9042
	var v9049 int32
	_ = v9049
	var v9050 int32
	_ = v9050
	var v9052 int32
	_ = v9052
	var v9053 int32
	_ = v9053
	var v9059 int32
	_ = v9059
	var v9060 int32
	_ = v9060
	var v9063 int32
	_ = v9063
	var v9071 int32
	_ = v9071
	var v9076 int32
	_ = v9076
	var v9080 int32
	_ = v9080
	var v9085 int32
	_ = v9085
	var v9087 int32
	_ = v9087
	var v9091 int32
	_ = v9091
	var v9097 int32
	_ = v9097
	var v9100 int32
	_ = v9100
	var v9106 int32
	_ = v9106
	var v9110 int32
	_ = v9110
	var v9112 int32
	_ = v9112
	var v9120 int32
	_ = v9120
	var v9123 int32
	_ = v9123
	var v9127 int32
	_ = v9127
	var v9129 int32
	_ = v9129
	var v9130 int64
	_ = v9130
	var v9138 int32
	_ = v9138
	var v9142 int32
	_ = v9142
	var v9146 int32
	_ = v9146
	var v9152 int32
	_ = v9152
	var v9156 int32
	_ = v9156
	var v9157 int32
	_ = v9157
	var v9164 int32
	_ = v9164
	var v9165 int32
	_ = v9165
	var v9166 int32
	_ = v9166
	var v9170 int32
	_ = v9170
	var v9173 int32
	_ = v9173
	var v9177 int32
	_ = v9177
	var v9178 int32
	_ = v9178
	var v9186 int32
	_ = v9186
	var v9192 int32
	_ = v9192
	var v9194 int32
	_ = v9194
	var v9198 int32
	_ = v9198
	var v9206 int32
	_ = v9206
	var v9210 int32
	_ = v9210
	var v9211 int32
	_ = v9211
	var v9212 int32
	_ = v9212
	var v9213 int32
	_ = v9213
	var v9214 int32
	_ = v9214
	var v9215 int32
	_ = v9215
	var v9216 int32
	_ = v9216
	var v9220 int32
	_ = v9220
	var v9222 int32
	_ = v9222
	var v9255 int32
	_ = v9255
	var v9259 int32
	_ = v9259
	var v9267 int32
	_ = v9267
	var v9268 int32
	_ = v9268
	var v9271 int32
	_ = v9271
	var v9275 int32
	_ = v9275
	var v9283 int32
	_ = v9283
	var v9284 int32
	_ = v9284
	var v9287 int32
	_ = v9287
	var v9291 int32
	_ = v9291
	var v9299 int32
	_ = v9299
	var v9300 int32
	_ = v9300
	var v9302 int32
	_ = v9302
	var v9305 int32
	_ = v9305
	var v9309 int32
	_ = v9309
	var v9310 int32
	_ = v9310
	var v9313 int32
	_ = v9313
	var v9314 int32
	_ = v9314
	var v9319 int32
	_ = v9319
	var v9323 int32
	_ = v9323
	var v9324 int32
	_ = v9324
	var v9327 int32
	_ = v9327
	var v9328 int32
	_ = v9328
	var v9332 int32
	_ = v9332
	var v9336 int32
	_ = v9336
	var v9338 int32
	_ = v9338
	var v9340 int32
	_ = v9340
	var v9341 int32
	_ = v9341
	var v9342 int32
	_ = v9342
	var v9344 int32
	_ = v9344
	var v9356 int32
	_ = v9356
	var v9363 int32
	_ = v9363
	var v9367 int32
	_ = v9367
	var v9370 int32
	_ = v9370
	var v9371 int32
	_ = v9371
	var v9372 int32
	_ = v9372
	var v9373 int32
	_ = v9373
	var v9374 int32
	_ = v9374
	var v9380 int32
	_ = v9380
	var v9383 int32
	_ = v9383
	var v9384 int32
	_ = v9384
	var v9385 int32
	_ = v9385
	var v9390 int32
	_ = v9390
	var v9394 int32
	_ = v9394
	var v9397 int32
	_ = v9397
	var v9398 int32
	_ = v9398
	var v9404 int32
	_ = v9404
	var v9407 int32
	_ = v9407
	var v9408 int32
	_ = v9408
	var v9409 int32
	_ = v9409
	var v9414 int32
	_ = v9414
	var v9418 int32
	_ = v9418
	var v9421 int32
	_ = v9421
	var v9422 int32
	_ = v9422
	var v9428 int32
	_ = v9428
	var v9429 int32
	_ = v9429
	var v9430 int32
	_ = v9430
	var v9431 int32
	_ = v9431
	var v9436 int32
	_ = v9436
	var v9440 int32
	_ = v9440
	var v9443 int32
	_ = v9443
	var v9444 int32
	_ = v9444
	var v9450 int32
	_ = v9450
	var v9451 int32
	_ = v9451
	var v9452 int32
	_ = v9452
	var v9453 int32
	_ = v9453
	var v9458 int32
	_ = v9458
	var v9463 int32
	_ = v9463
	var v9466 int32
	_ = v9466
	var v9467 int32
	_ = v9467
	var v9468 int32
	_ = v9468
	var v9469 int32
	_ = v9469
	var v9475 int32
	_ = v9475
	var v9476 int32
	_ = v9476
	var v9477 int32
	_ = v9477
	var v9478 int32
	_ = v9478
	var v9483 int32
	_ = v9483
	var v9489 int32
	_ = v9489
	var v9495 int32
	_ = v9495
	var v9501 int32
	_ = v9501
	var v9506 int32
	_ = v9506
	var v9510 int32
	_ = v9510
	var v9515 int32
	_ = v9515
	var v9519 int32
	_ = v9519
	var v9522 int32
	_ = v9522
	var v9523 int32
	_ = v9523
	var v9524 int32
	_ = v9524
	var v9530 int32
	_ = v9530
	var v9531 int32
	_ = v9531
	var v9532 int32
	_ = v9532
	var v9533 int32
	_ = v9533
	var v9538 int32
	_ = v9538
	var v9542 int32
	_ = v9542
	var v9545 int32
	_ = v9545
	var v9549 int32
	_ = v9549
	var v9552 int32
	_ = v9552
	var v9553 int32
	_ = v9553
	var v9554 int32
	_ = v9554
	var v9559 int32
	_ = v9559
	var v9563 int32
	_ = v9563
	var v9566 int32
	_ = v9566
	var v9570 int32
	_ = v9570
	var v9573 int32
	_ = v9573
	var v9574 int32
	_ = v9574
	var v9575 int32
	_ = v9575
	var v9580 int32
	_ = v9580
	var v9584 int32
	_ = v9584
	var v9587 int32
	_ = v9587
	var v9591 int32
	_ = v9591
	var v9592 int32
	_ = v9592
	var v9593 int32
	_ = v9593
	var v9598 int32
	_ = v9598
	var v9602 int32
	_ = v9602
	var v9605 int32
	_ = v9605
	var v9609 int32
	_ = v9609
	var v9612 int32
	_ = v9612
	var v9613 int32
	_ = v9613
	var v9614 int32
	_ = v9614
	var v9619 int32
	_ = v9619
	var v9623 int32
	_ = v9623
	var v9626 int32
	_ = v9626
	var v9630 int32
	_ = v9630
	var v9633 int32
	_ = v9633
	var v9634 int32
	_ = v9634
	var v9635 int32
	_ = v9635
	var v9640 int32
	_ = v9640
	var v9644 int32
	_ = v9644
	var v9647 int32
	_ = v9647
	var v9648 int32
	_ = v9648
	var v9654 int32
	_ = v9654
	var v9657 int32
	_ = v9657
	var v9658 int32
	_ = v9658
	var v9659 int32
	_ = v9659
	var v9664 int32
	_ = v9664
	var v9668 int32
	_ = v9668
	var v9671 int32
	_ = v9671
	var v9674 int32
	_ = v9674
	var v9675 int32
	_ = v9675
	var v9678 int32
	_ = v9678
	var v9679 int32
	_ = v9679
	var v9680 int32
	_ = v9680
	var v9681 int32
	_ = v9681
	var v9686 int32
	_ = v9686
	var v9690 int32
	_ = v9690
	var v9693 int32
	_ = v9693
	var v9697 int32
	_ = v9697
	var v9698 int32
	_ = v9698
	var v9699 int32
	_ = v9699
	var v9704 int32
	_ = v9704
	var v9708 int32
	_ = v9708
	var v9711 int32
	_ = v9711
	var v9715 int32
	_ = v9715
	var v9716 int32
	_ = v9716
	var v9717 int32
	_ = v9717
	var v9718 int32
	_ = v9718
	var v9723 int32
	_ = v9723
	var v9727 int32
	_ = v9727
	var v9730 int32
	_ = v9730
	var v9734 int32
	_ = v9734
	var v9735 int32
	_ = v9735
	var v9736 int32
	_ = v9736
	var v9741 int32
	_ = v9741
	var v9744 int32
	_ = v9744
	var v9748 int32
	_ = v9748
	var v9749 int32
	_ = v9749
	var v9750 int32
	_ = v9750
	var v9751 int32
	_ = v9751
	var v9756 int32
	_ = v9756
	var v9760 int32
	_ = v9760
	var v9763 int32
	_ = v9763
	var v9767 int32
	_ = v9767
	var v9768 int32
	_ = v9768
	var v9769 int32
	_ = v9769
	var v9770 int32
	_ = v9770
	var v9775 int32
	_ = v9775
	var v9781 int32
	_ = v9781
	var v9787 int32
	_ = v9787
	var v9793 int32
	_ = v9793
	var v9799 int32
	_ = v9799
	var v9806 int32
	_ = v9806
	var v9812 int32
	_ = v9812
	var v9816 int32
	_ = v9816
	var v9819 int32
	_ = v9819
	var v9823 int32
	_ = v9823
	var v9828 int32
	_ = v9828
	var v9829 int32
	_ = v9829
	var v9831 int32
	_ = v9831
	var v9832 int32
	_ = v9832
	var v9834 int32
	_ = v9834
	var v9840 int32
	_ = v9840
	var v9846 int32
	_ = v9846
	var v9852 int32
	_ = v9852
	var v9856 int32
	_ = v9856
	var v9859 int32
	_ = v9859
	var v9863 int32
	_ = v9863
	var v9866 int32
	_ = v9866
	var v9867 int32
	_ = v9867
	var v9868 int32
	_ = v9868
	var v9873 int32
	_ = v9873
	var v9877 int32
	_ = v9877
	var v9880 int32
	_ = v9880
	var v9884 int32
	_ = v9884
	var v9885 int32
	_ = v9885
	var v9886 int32
	_ = v9886
	var v9887 int32
	_ = v9887
	var v9892 int32
	_ = v9892
	var v9896 int32
	_ = v9896
	var v9899 int32
	_ = v9899
	var v9900 int32
	_ = v9900
	var v9901 int32
	_ = v9901
	var v9907 int32
	_ = v9907
	var v9908 int32
	_ = v9908
	var v9909 int32
	_ = v9909
	var v9910 int32
	_ = v9910
	var v9915 int32
	_ = v9915
	var v9921 int32
	_ = v9921
	var v9927 int32
	_ = v9927
	var v9933 int32
	_ = v9933
	var v9939 int32
	_ = v9939
	var v9947 int32
	_ = v9947
	var v9952 int32
	_ = v9952
	var v9955 int32
	_ = v9955
	var v9959 int32
	_ = v9959
	var v9964 int32
	_ = v9964
	var v9966 int32
	_ = v9966
	var v9969 int32
	_ = v9969
	var v9982 int32
	_ = v9982
	var v9983 int32
	_ = v9983
	var v9984 int32
	_ = v9984
	var v9990 int32
	_ = v9990
	var v9993 int32
	_ = v9993
	var v9997 int32
	_ = v9997
	var v9998 int32
	_ = v9998
	var v9999 int32
	_ = v9999
	var v10004 int32
	_ = v10004
	var v10005 int32
	_ = v10005
	var v10006 int32
	_ = v10006
	var v10010 int32
	_ = v10010
	var v10013 int32
	_ = v10013
	var v10014 int32
	_ = v10014
	var v10015 int32
	_ = v10015
	var v10016 int32
	_ = v10016
	var v10019 int32
	_ = v10019
	var v10025 int32
	_ = v10025
	var v10026 int32
	_ = v10026
	var v10036 int32
	_ = v10036
	var v10037 int32
	_ = v10037
	var v10041 int32
	_ = v10041
	var v10047 int32
	_ = v10047
	var v10053 int32
	_ = v10053
	var v10062 int32
	_ = v10062
	var v10063 int32
	_ = v10063
	var v10064 int32
	_ = v10064
	var v10068 int32
	_ = v10068
	var v10072 int32
	_ = v10072
	var v10080 int32
	_ = v10080
	var v10081 int32
	_ = v10081
	var v10082 int32
	_ = v10082
	var v10085 int32
	_ = v10085
	var v10088 int32
	_ = v10088
	var v10091 int32
	_ = v10091
	var v10094 int32
	_ = v10094
	var v10097 int32
	_ = v10097
	var v10099 int32
	_ = v10099
	var v10100 int32
	_ = v10100
	var v10102 int32
	_ = v10102
	var v10103 int32
	_ = v10103
	var v10105 int32
	_ = v10105
	var v10106 int32
	_ = v10106
	var v10110 int32
	_ = v10110
	var v10111 int32
	_ = v10111
	var v10113 int32
	_ = v10113
	var v10125 int32
	_ = v10125
	var v10128 int32
	_ = v10128
	var v10132 int32
	_ = v10132
	var v10135 int32
	_ = v10135
	var v10136 int32
	_ = v10136
	var v10137 int32
	_ = v10137
	var v10142 int32
	_ = v10142
	var v10148 int32
	_ = v10148
	var v10152 int32
	_ = v10152
	var v10157 int32
	_ = v10157
	var v10163 int32
	_ = v10163
	var v10193 int32
	_ = v10193
	var v10201 int32
	_ = v10201
	var v10215 int32
	_ = v10215
	var v10218 int32
	_ = v10218
	var v10224 int32
	_ = v10224
	var v10225 int32
	_ = v10225
	var v10229 int32
	_ = v10229
	var v10232 int32
	_ = v10232
	var v10234 int32
	_ = v10234
	var v10235 int32
	_ = v10235
	var v10236 int32
	_ = v10236
	var v10239 int32
	_ = v10239
	var v10240 int32
	_ = v10240
	var v10253 int32
	_ = v10253
	var v10255 int32
	_ = v10255
	var v10256 int32
	_ = v10256
	var v10262 int32
	_ = v10262
	var v10264 int32
	_ = v10264
	var v10266 int32
	_ = v10266
	var v10267 int32
	_ = v10267
	var v10269 int32
	_ = v10269
	var v10270 int32
	_ = v10270
	var v10271 int32
	_ = v10271
	var v10272 int32
	_ = v10272
	var v10273 int32
	_ = v10273
	var v10275 int32
	_ = v10275
	var v10281 int32
	_ = v10281
	var v10288 int32
	_ = v10288
	var v10317 int32
	_ = v10317
	var v10333 int32
	_ = v10333
	var v10336 int32
	_ = v10336
	var v10337 int64
	_ = v10337
	var v10339 int64
	_ = v10339
	var v10343 int32
	_ = v10343
	var v10344 int32
	_ = v10344
	var v10346 int32
	_ = v10346
	var v10347 int32
	_ = v10347
	var v10350 int32
	_ = v10350
	var v10354 int32
	_ = v10354
	var v10357 int32
	_ = v10357
	var v10358 int32
	_ = v10358
	var v10365 int32
	_ = v10365
	var v10373 int32
	_ = v10373
	var v10376 int32
	_ = v10376
	var v10380 int32
	_ = v10380
	var v10382 int32
	_ = v10382
	var v10386 int32
	_ = v10386
	var v10391 int32
	_ = v10391
	var v10394 int32
	_ = v10394
	var v10408 int32
	_ = v10408
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
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(257572))
	mBase = m.M
	v10408 = m.ExcPending
	if v10408 != 0 {
		goto L66
	} else {
		goto L2748
	}
L3:
	;
	goto L2
L4:
	;
	v42 = v10380
	v44 = v10382
	v48 = v10386
	v53 = v10391 + int32(2)
	v55 = v136
	v56 = v10394
	v57 = v138
	v60 = v140
	v61 = v141
	goto L1
L5:
	;
	v10333 = int32(0) - v247
	v10336 = v132 + v10333<<(uint(int32(4))%32)
	v10337 = *(*int64)(unsafe.Add(mBase, uint32(v27)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v10336)+16)) = v10337
	v10339 = *(*int64)(unsafe.Add(mBase, uint32(v27)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v10336)+24)) = v10339
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v262
	v10343 = v10336 + int32(16)
	v10344 = int32(1)
	v10346 = v134 + v10333<<(uint(v10344)%32)
	v10347 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10346))))
	v10350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[1353]))))
	v10354 = (v10350 - int32(137)) << (uint(v10344) % 32)
	v10357 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10354)+uint32(_consts[1354]))))
	v10358 = v10347 + v10357
	if base.Ui32(int32(1293)) < base.Ui32(v10358) {
		goto L2745
	} else {
		goto L2746
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v10288
	v10317 = int32(-2)
	goto L5
L7:
	;
	v10157 = int32(0)
	v10163 = v10152
	goto L2712
L8:
	;
	F_plpgsql_push_back_token(m, v2516, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v10148 = m.ExcPending
	if v10148 != 0 {
		goto L66
	} else {
		goto L2711
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10005)+4)) = int32(2)
	v10010 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1355])))
	if v10010 == int32(1) {
		goto L2689
	} else {
		goto L2690
	}
L10:
	;
	v9966 = int32(1)
	v9969 = int32(0)
	v9982 = F_read_sql_construct(m, int32(269), int32(336), v9969, int32(518719), v9969, v9966, v9969, v27+int32(4764), v27+int32(4752), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9983 = m.ExcPending
	if v9983 != 0 {
		goto L66
	} else {
		goto L2682
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9952 = m.ExcPending
	if v9952 != 0 {
		goto L66
	} else {
		goto L2678
	}
L12:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(706339))
	mBase = m.M
	v9947 = m.ExcPending
	if v9947 != 0 {
		goto L66
	} else {
		goto L2677
	}
L13:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v9939 = m.ExcPending
	if v9939 != 0 {
		goto L66
	} else {
		goto L2676
	}
L14:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(407821))
	mBase = m.M
	v9933 = m.ExcPending
	if v9933 != 0 {
		goto L66
	} else {
		goto L2675
	}
L15:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(407821))
	mBase = m.M
	v9927 = m.ExcPending
	if v9927 != 0 {
		goto L66
	} else {
		goto L2674
	}
L16:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v9921 = m.ExcPending
	if v9921 != 0 {
		goto L66
	} else {
		goto L2673
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9896 = m.ExcPending
	if v9896 != 0 {
		goto L66
	} else {
		goto L2668
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9877 = m.ExcPending
	if v9877 != 0 {
		goto L66
	} else {
		goto L2663
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9856 = m.ExcPending
	if v9856 != 0 {
		goto L66
	} else {
		goto L2658
	}
L20:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v9852 = m.ExcPending
	if v9852 != 0 {
		goto L66
	} else {
		goto L2657
	}
L21:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v9846 = m.ExcPending
	if v9846 != 0 {
		goto L66
	} else {
		goto L2656
	}
L22:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v9840 = m.ExcPending
	if v9840 != 0 {
		goto L66
	} else {
		goto L2655
	}
L23:
	;
	v9832 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v9832, l1)
	mBase = m.M
	v9834 = m.ExcPending
	if v9834 != 0 {
		goto L66
	} else {
		goto L2654
	}
L24:
	;
	v9829 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v9829, l1)
	mBase = m.M
	v9831 = m.ExcPending
	if v9831 != 0 {
		goto L66
	} else {
		goto L2653
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9816 = m.ExcPending
	if v9816 != 0 {
		goto L66
	} else {
		goto L2649
	}
L26:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(706415))
	mBase = m.M
	v9812 = m.ExcPending
	if v9812 != 0 {
		goto L66
	} else {
		goto L2648
	}
L27:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(243097))
	mBase = m.M
	v9806 = m.ExcPending
	if v9806 != 0 {
		goto L66
	} else {
		goto L2647
	}
L28:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(246872))
	mBase = m.M
	v9799 = m.ExcPending
	if v9799 != 0 {
		goto L66
	} else {
		goto L2646
	}
L29:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(407821))
	mBase = m.M
	v9793 = m.ExcPending
	if v9793 != 0 {
		goto L66
	} else {
		goto L2645
	}
L30:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(407821))
	mBase = m.M
	v9787 = m.ExcPending
	if v9787 != 0 {
		goto L66
	} else {
		goto L2644
	}
L31:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v9781 = m.ExcPending
	if v9781 != 0 {
		goto L66
	} else {
		goto L2643
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9760 = m.ExcPending
	if v9760 != 0 {
		goto L66
	} else {
		goto L2638
	}
L33:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9744 = m.ExcPending
	if v9744 != 0 {
		goto L66
	} else {
		goto L2634
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9727 = m.ExcPending
	if v9727 != 0 {
		goto L66
	} else {
		goto L2629
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9708 = m.ExcPending
	if v9708 != 0 {
		goto L66
	} else {
		goto L2624
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9690 = m.ExcPending
	if v9690 != 0 {
		goto L66
	} else {
		goto L2619
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9668 = m.ExcPending
	if v9668 != 0 {
		goto L66
	} else {
		goto L2611
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9644 = m.ExcPending
	if v9644 != 0 {
		goto L66
	} else {
		goto L2606
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9623 = m.ExcPending
	if v9623 != 0 {
		goto L66
	} else {
		goto L2601
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9602 = m.ExcPending
	if v9602 != 0 {
		goto L66
	} else {
		goto L2596
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9584 = m.ExcPending
	if v9584 != 0 {
		goto L66
	} else {
		goto L2591
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9563 = m.ExcPending
	if v9563 != 0 {
		goto L66
	} else {
		goto L2586
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9542 = m.ExcPending
	if v9542 != 0 {
		goto L66
	} else {
		goto L2581
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9519 = m.ExcPending
	if v9519 != 0 {
		goto L66
	} else {
		goto L2575
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9506 = m.ExcPending
	if v9506 != 0 {
		goto L66
	} else {
		goto L2572
	}
L46:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(706092))
	mBase = m.M
	v9501 = m.ExcPending
	if v9501 != 0 {
		goto L66
	} else {
		goto L2571
	}
L47:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(257550))
	mBase = m.M
	v9495 = m.ExcPending
	if v9495 != 0 {
		goto L66
	} else {
		goto L2570
	}
L48:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(257550))
	mBase = m.M
	v9489 = m.ExcPending
	if v9489 != 0 {
		goto L66
	} else {
		goto L2569
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9463 = m.ExcPending
	if v9463 != 0 {
		goto L66
	} else {
		goto L2563
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9440 = m.ExcPending
	if v9440 != 0 {
		goto L66
	} else {
		goto L2558
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9418 = m.ExcPending
	if v9418 != 0 {
		goto L66
	} else {
		goto L2553
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9394 = m.ExcPending
	if v9394 != 0 {
		goto L66
	} else {
		goto L2548
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9367 = m.ExcPending
	if v9367 != 0 {
		goto L66
	} else {
		goto L2542
	}
L54:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), l0, l1, int32(434459))
	mBase = m.M
	v9363 = m.ExcPending
	if v9363 != 0 {
		goto L66
	} else {
		goto L2541
	}
L55:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), l0, l1, int32(209437))
	mBase = m.M
	v9356 = m.ExcPending
	if v9356 != 0 {
		goto L66
	} else {
		goto L2540
	}
L56:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_consts[1356]))))
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
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[1357]))))
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
	v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42<<(uint(int32(1))%32))+uint32(_consts[1358]))))
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
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+uint32(_consts[1359]))))
	v169 = v158
	v170 = v168
	goto L89
L98:
	;
	v175 = v171 << (uint(int32(1)) % 32)
	v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[1360]))))
	if v170 != v178 {
		v232 = v169
		goto L57
	} else {
		goto L99
	}
L99:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[1361]))))
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
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	*(*int64)(unsafe.Add(mBase, uint32(v132)+16)) = v194
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363])))
	*(*int64)(unsafe.Add(mBase, uint32(v132)+24)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
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
	v10380 = v183
	v10382 = v132 + int32(16)
	v10386 = v204
	v10391 = v134
	v10394 = v142 + int32(4)
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
		v10317 = v241
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
	v9344 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v9344 == int32(0) {
		goto L13
	} else {
		goto L2539
	}
L120:
	;
	v9340 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v9341 = F_pstrdup(m, v9340)
	mBase = m.M
	v9342 = m.ExcPending
	if v9342 != 0 {
		goto L66
	} else {
		goto L2538
	}
L121:
	;
	v9338 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9338
	v10317 = v241
	goto L5
L122:
	;
	v9336 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9336
	v10317 = v241
	goto L5
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L124:
	;
	v9332 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9332
	v10317 = v241
	goto L5
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L126:
	;
	v9323 = v132 - int32(16)
	v9324 = *(*int32)(unsafe.Add(mBase, uint32(v9323)))
	F_plpgsql_ns_push(m, v9324, int32(1))
	mBase = m.M
	v9327 = m.ExcPending
	if v9327 != 0 {
		goto L66
	} else {
		goto L2537
	}
L127:
	;
	F_plpgsql_ns_push(m, int32(0), int32(1))
	mBase = m.M
	v9319 = m.ExcPending
	if v9319 != 0 {
		goto L66
	} else {
		goto L2536
	}
L128:
	;
	v9309 = v132 - int32(16)
	v9310 = *(*int32)(unsafe.Add(mBase, uint32(v9309)))
	F_plpgsql_ns_push(m, v9310, int32(0))
	mBase = m.M
	v9313 = m.ExcPending
	if v9313 != 0 {
		goto L66
	} else {
		goto L2535
	}
L129:
	;
	v9302 = int32(0)
	F_plpgsql_ns_push(m, v9302, v9302)
	mBase = m.M
	v9305 = m.ExcPending
	if v9305 != 0 {
		goto L66
	} else {
		goto L2534
	}
L130:
	;
	v9287 = int32(0)
	v9291 = int32(1)
	v9299 = F_read_sql_construct(m, int32(336), v9287, v9287, int32(518719), int32(2), v9291, v9291, v9287, v9287, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9300 = m.ExcPending
	if v9300 != 0 {
		goto L66
	} else {
		goto L2533
	}
L131:
	;
	v9271 = int32(0)
	v9275 = int32(1)
	v9283 = F_read_sql_construct(m, int32(376), v9271, v9271, int32(522683), int32(2), v9275, v9275, v9271, v9271, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9284 = m.ExcPending
	if v9284 != 0 {
		goto L66
	} else {
		goto L2532
	}
L132:
	;
	v9255 = int32(0)
	v9259 = int32(1)
	v9267 = F_read_sql_construct(m, int32(59), v9255, v9255, int32(538734), int32(2), v9259, v9259, v9255, v9255, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9268 = m.ExcPending
	if v9268 != 0 {
		goto L66
	} else {
		goto L2531
	}
L133:
	;
	v9025 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v9026 = int32(346115)
	v9029 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1365])))
	v9030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9025))))
	if v9030 == int32(0) {
		v9049 = v9029
		v9050 = v9030
		goto L2477
	} else {
		goto L2478
	}
L134:
	;
	v9023 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9023
	v10317 = v241
	goto L5
L135:
	;
	v8992 = v132 - int32(32)
	v8993 = *(*int32)(unsafe.Add(mBase, uint32(v8992)))
	v8997 = v8993
	goto L2473
L136:
	;
	v8932 = F_palloc0(m, int32(12))
	mBase = m.M
	v8933 = m.ExcPending
	if v8933 != 0 {
		goto L66
	} else {
		goto L2458
	}
L137:
	;
	v8922 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+272)) = v8922
	*(*int32)(unsafe.Add(mBase, uint32(v27)+276)) = v8922
	v8928 = F_list_make1_impl(m, int32(1), v27+int32(272))
	mBase = m.M
	v8929 = m.ExcPending
	if v8929 != 0 {
		goto L66
	} else {
		goto L2457
	}
L138:
	;
	v8917 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8918 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v8919 = F_lappend(m, v8917, v8918)
	mBase = m.M
	v8920 = m.ExcPending
	if v8920 != 0 {
		goto L66
	} else {
		goto L2456
	}
L139:
	;
	v8911 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8912 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v8911)+8)) = v8912
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8911
	v10317 = v241
	goto L5
L140:
	;
	v8824 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v8825 = int32(0)
	if v8824 < v8825 {
		v8868 = v8825
		goto L2438
	} else {
		goto L2439
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L142:
	;
	v8819 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v8819, l1)
	mBase = m.M
	v8821 = m.ExcPending
	if v8821 != 0 {
		goto L66
	} else {
		goto L2436
	}
L143:
	;
	v8816 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v8816, l1)
	mBase = m.M
	v8818 = m.ExcPending
	if v8818 != 0 {
		goto L66
	} else {
		goto L2435
	}
L144:
	;
	v8804 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v8805 = *(*int32)(unsafe.Add(mBase, uint32(v8804)))
	if v8805 != 0 {
		goto L18
	} else {
		goto L2431
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(1)
	v10317 = v241
	goto L5
L148:
	;
	v8730 = F_palloc(m, int32(16))
	mBase = m.M
	v8731 = m.ExcPending
	if v8731 != 0 {
		goto L66
	} else {
		goto L2416
	}
L149:
	;
	v8661 = F_palloc(m, int32(16))
	mBase = m.M
	v8662 = m.ExcPending
	if v8662 != 0 {
		goto L66
	} else {
		goto L2401
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L151:
	;
	v8591 = F_palloc(m, int32(16))
	mBase = m.M
	v8592 = m.ExcPending
	if v8592 != 0 {
		goto L66
	} else {
		goto L2386
	}
L152:
	;
	v8182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = uint8(v8182)
	v8185 = F_palloc0(m, int32(36))
	mBase = m.M
	v8186 = m.ExcPending
	if v8186 != 0 {
		goto L66
	} else {
		goto L2264
	}
L153:
	;
	v8123 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	v8126 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v8127 = int32(0)
	if v8126 < v8127 {
		v8170 = v8127
		goto L2251
	} else {
		goto L2252
	}
L154:
	;
	v8040 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	F_read_into_target(m, v27+int32(4784), int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8049 = m.ExcPending
	if v8049 != 0 {
		goto L66
	} else {
		goto L2232
	}
L155:
	;
	v7723 = F_palloc0(m, int32(36))
	mBase = m.M
	v7724 = m.ExcPending
	if v7724 != 0 {
		goto L66
	} else {
		goto L2157
	}
L156:
	;
	v7538 = int32(1)
	v7547 = F_read_sql_construct(m, int32(332), int32(381), int32(59), int32(538710), int32(2), v7538, v7538, int32(0), v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7548 = m.ExcPending
	if v7548 != 0 {
		goto L66
	} else {
		goto L2122
	}
L157:
	;
	v7509 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7510 = m.ExcPending
	if v7510 != 0 {
		goto L66
	} else {
		goto L2115
	}
L158:
	;
	v7482 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7483 = m.ExcPending
	if v7483 != 0 {
		goto L66
	} else {
		goto L2108
	}
L159:
	;
	v7469 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7475 = F_make_execsql_stmt(m, int32(337), v7469, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7476 = m.ExcPending
	if v7476 != 0 {
		goto L66
	} else {
		goto L2107
	}
L160:
	;
	v7459 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7465 = F_make_execsql_stmt(m, int32(331), v7459, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7466 = m.ExcPending
	if v7466 != 0 {
		goto L66
	} else {
		goto L2106
	}
L161:
	;
	v7449 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7455 = F_make_execsql_stmt(m, int32(328), v7449, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7456 = m.ExcPending
	if v7456 != 0 {
		goto L66
	} else {
		goto L2105
	}
L162:
	;
	v7438 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7438
	v7442 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v7442
	v7446 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v7446
	v10317 = v241
	goto L5
L163:
	;
	v7338 = F_palloc(m, int32(20))
	mBase = m.M
	v7339 = m.ExcPending
	if v7339 != 0 {
		goto L66
	} else {
		goto L2085
	}
L164:
	;
	v6251 = F_palloc(m, int32(32))
	mBase = m.M
	v6252 = m.ExcPending
	if v6252 != 0 {
		goto L66
	} else {
		goto L1764
	}
L165:
	;
	v5660 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5661 = m.ExcPending
	if v5661 != 0 {
		goto L66
	} else {
		goto L1624
	}
L166:
	;
	v5654 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v5654)
	v10317 = v241
	goto L5
L167:
	;
	v5652 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v5652)
	v10317 = v241
	goto L5
L168:
	;
	v5520 = F_palloc0(m, int32(24))
	mBase = m.M
	v5521 = m.ExcPending
	if v5521 != 0 {
		goto L66
	} else {
		goto L1571
	}
L169:
	;
	v5517 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5517
	v10317 = v241
	goto L5
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L171:
	;
	v5403 = F_palloc0(m, int32(32))
	mBase = m.M
	v5404 = m.ExcPending
	if v5404 != 0 {
		goto L66
	} else {
		goto L1543
	}
L172:
	;
	v5399 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v5399, l1)
	mBase = m.M
	v5401 = m.ExcPending
	if v5401 != 0 {
		goto L66
	} else {
		goto L1542
	}
L173:
	;
	v5330 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5330
	v5332 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5333 = int32(0)
	if v5332 < v5333 {
		v5376 = v5333
		goto L1525
	} else {
		goto L1526
	}
L174:
	;
	v5239 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v5239 != 0 {
		goto L1499
	} else {
		goto L1500
	}
L175:
	;
	v4854 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4855 = m.ExcPending
	if v4855 != 0 {
		goto L66
	} else {
		goto L1429
	}
L176:
	;
	v4769 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v4770 = *(*int32)(unsafe.Add(mBase, uint32(v4769)))
	v4773 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4774 = int32(0)
	if v4773 < v4774 {
		v4817 = v4774
		goto L1405
	} else {
		goto L1406
	}
L177:
	;
	v4679 = F_palloc0(m, int32(24))
	mBase = m.M
	v4680 = m.ExcPending
	if v4680 != 0 {
		goto L66
	} else {
		goto L1381
	}
L178:
	;
	v4594 = F_palloc0(m, int32(20))
	mBase = m.M
	v4595 = m.ExcPending
	if v4595 != 0 {
		goto L66
	} else {
		goto L1358
	}
L179:
	;
	v4581 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v4581 != 0 {
		goto L1354
	} else {
		goto L1355
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L181:
	;
	v4520 = F_palloc(m, int32(12))
	mBase = m.M
	v4521 = m.ExcPending
	if v4521 != 0 {
		goto L66
	} else {
		goto L1339
	}
L182:
	;
	v4510 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v4510
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v4510
	v4516 = F_list_make1_impl(m, int32(1), v27+int32(216))
	mBase = m.M
	v4517 = m.ExcPending
	if v4517 != 0 {
		goto L66
	} else {
		goto L1338
	}
L183:
	;
	v4505 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v4506 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v4507 = F_lappend(m, v4505, v4506)
	mBase = m.M
	v4508 = m.ExcPending
	if v4508 != 0 {
		goto L66
	} else {
		goto L1337
	}
L184:
	;
	v4469 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		goto L66
	} else {
		goto L1330
	}
L185:
	;
	v4257 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(80))))
	v4260 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	v4263 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	v4266 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(24))))
	v4268 = F_palloc(m, int32(32))
	mBase = m.M
	v4269 = m.ExcPending
	if v4269 != 0 {
		goto L66
	} else {
		goto L1296
	}
L186:
	;
	v4253 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4253
	v10317 = v241
	goto L5
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L188:
	;
	v4187 = F_palloc0(m, int32(12))
	mBase = m.M
	v4188 = m.ExcPending
	if v4188 != 0 {
		goto L66
	} else {
		goto L1280
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L190:
	;
	v4106 = F_palloc0(m, int32(28))
	mBase = m.M
	v4107 = m.ExcPending
	if v4107 != 0 {
		goto L66
	} else {
		goto L1265
	}
L191:
	;
	v4102 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_cword_is_not_variable(m, v132, v4102, l1)
	mBase = m.M
	v4104 = m.ExcPending
	if v4104 != 0 {
		goto L66
	} else {
		goto L1264
	}
L192:
	;
	v4099 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v4099, l1)
	mBase = m.M
	v4101 = m.ExcPending
	if v4101 != 0 {
		goto L66
	} else {
		goto L1263
	}
L193:
	;
	v4083 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v4084 = *(*int32)(unsafe.Add(mBase, uint32(v4083)))
	if base.Ui32(v4084-int32(1)) < base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L1259
	}
L194:
	;
	v3673 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L66
	} else {
		goto L1132
	}
L195:
	;
	v3659 = F_palloc(m, int32(8))
	mBase = m.M
	v3660 = m.ExcPending
	if v3660 != 0 {
		goto L66
	} else {
		goto L1104
	}
L196:
	;
	v3649 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+172)) = v3649
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v3649
	v3655 = F_list_make1_impl(m, int32(1), v27+int32(172))
	mBase = m.M
	v3656 = m.ExcPending
	if v3656 != 0 {
		goto L66
	} else {
		goto L1103
	}
L197:
	;
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	v3645 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3646 = F_lappend(m, v3644, v3645)
	mBase = m.M
	v3647 = m.ExcPending
	if v3647 != 0 {
		goto L66
	} else {
		goto L1102
	}
L198:
	;
	v3640 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v3640)
	v10317 = v241
	goto L5
L199:
	;
	v3638 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v3638)
	v10317 = v241
	goto L5
L200:
	;
	v3636 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v3636)
	v10317 = v241
	goto L5
L201:
	;
	v3411 = F_palloc0(m, int32(20))
	mBase = m.M
	v3412 = m.ExcPending
	if v3412 != 0 {
		goto L66
	} else {
		goto L1049
	}
L202:
	;
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v3295 == int32(0) {
		goto L1023
	} else {
		goto L1024
	}
L203:
	;
	v3205 = F_palloc0(m, int32(24))
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L66
	} else {
		goto L1006
	}
L204:
	;
	v3115 = F_palloc0(m, int32(24))
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		goto L66
	} else {
		goto L989
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
	v10317 = v241
	goto L5
L207:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2770
	v10317 = v241
	goto L5
L208:
	;
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2768
	v10317 = v241
	goto L5
L209:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2766
	v10317 = v241
	goto L5
L210:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2764
	v10317 = v241
	goto L5
L211:
	;
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2762
	v10317 = v241
	goto L5
L212:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2760
	v10317 = v241
	goto L5
L213:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2758
	v10317 = v241
	goto L5
L214:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2756
	v10317 = v241
	goto L5
L215:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2754
	v10317 = v241
	goto L5
L216:
	;
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2752
	v10317 = v241
	goto L5
L217:
	;
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2750
	v10317 = v241
	goto L5
L218:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2748
	v10317 = v241
	goto L5
L219:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2746
	v10317 = v241
	goto L5
L220:
	;
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2744
	v10317 = v241
	goto L5
L221:
	;
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2742
	v10317 = v241
	goto L5
L222:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2740
	v10317 = v241
	goto L5
L223:
	;
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2738
	v10317 = v241
	goto L5
L224:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2736
	v10317 = v241
	goto L5
L225:
	;
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2734
	v10317 = v241
	goto L5
L226:
	;
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2732
	v10317 = v241
	goto L5
L227:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2730
	v10317 = v241
	goto L5
L228:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2728
	v10317 = v241
	goto L5
L229:
	;
	v2726 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2726
	v10317 = v241
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
	v10317 = v241
	goto L5
L232:
	;
	v2697 = int32(0)
	v2701 = int32(1)
	v2709 = F_read_sql_construct(m, int32(59), v2697, v2697, int32(538734), int32(2), v2701, v2701, v2697, v2697, v27+int32(4736), v27+int32(4732), l1)
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
	v10317 = v241
	goto L5
L234:
	;
	v2692 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v2692)
	v10317 = v241
	goto L5
L235:
	;
	v2690 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v2690)
	v10317 = v241
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
	v10317 = v241
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
	v10317 = v241
	goto L5
L242:
	;
	v2212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+304)) = uint8(v2212)
	v10317 = v241
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
	v985 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v985 == int32(0) {
		goto L426
	} else {
		goto L427
	}
L247:
	;
	v837 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
	v10317 = v241
	goto L5
L253:
	;
	v614 = int32(0)
	v626 = F_read_sql_construct(m, int32(59), v614, v614, int32(538734), v614, v614, int32(1), v614, v614, v27+int32(4736), v27+int32(4732), l1)
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
	v10317 = v241
	goto L5
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(4)
	v10317 = v241
	goto L5
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L257:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
	F_errstart_cold(m, int32(21), int32(541439))
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
	*(*int32)(unsafe.Add(mBase, _consts[1348])) = int32(0)
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
	*(*int32)(unsafe.Add(mBase, _consts[1348])) = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+308)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v459
	v10317 = v241
	goto L5
L265:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1348])) = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+308)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v450
	v10317 = v241
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
	v10317 = v241
	goto L5
L269:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+488)) = int32(2)
	v10317 = v241
	goto L5
L270:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+488)) = int32(1)
	v10317 = v241
	goto L5
L271:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+488)) = int32(0)
	v10317 = v241
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
	*(*uint8)(unsafe.Add(mBase, _consts[1368])) = uint8(v270)
	v10317 = v241
	goto L5
L274:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v267
	v10317 = v241
	goto L5
L275:
	;
	v284 = int32(333774)
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
	v281 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+492)) = uint8(v282)
	v10317 = v241
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
	v313 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v313)+492)) = uint8(v314)
	v10317 = v241
	goto L5
L288:
	;
	goto L289
L289:
	;
	F_errstart_cold(m, int32(21), int32(541439))
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
	F_errmsg_internal(m, int32(180749), v27)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L66
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(26734), int32(396), int32(355297))
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
	v10317 = v241
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
	v404 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
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
	v437 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if v438 != 0 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v349
	v10317 = v241
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
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = v444
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
	v10317 = v241
	goto L5
L318:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1348])) = int32(1)
	v10317 = v241
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
	F_errmsg(m, int32(213434), int32(0))
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
	F_errfinish(m, int32(26734), int32(502), int32(355297))
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
		v10317 = v241
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
	v10317 = v241
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
	v10317 = v241
	goto L5
L336:
	;
	v10317 = v241
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
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = v575
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
	v10317 = v241
	goto L5
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v626
	v10317 = v241
	goto L5
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+8)) = int32(646955)
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
	v10317 = v241
	goto L5
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v815
	v10317 = v241
	goto L5
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v822
	v10317 = v241
	goto L5
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v833
	v10317 = v241
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
	v10317 = v241
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
	v10317 = v241
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
	v1288 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
	v1139 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
	v10317 = v241
	goto L5
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1488
	v1493 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
	v1638 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
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
		v10317 = v241
		goto L5
	} else {
		goto L594
	}
L592:
	;
	goto L593
L593:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
		v10317 = v241
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
	v1797 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1797)+500))
	v1804 = F_errstart(m, v1798&int32(2)+int32(19), int32(541439))
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
		v10317 = v241
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
	F_errmsg(m, int32(390916), v27+int32(96))
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
	F_errfinish(m, int32(26734), int32(737), int32(355297))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L66
	} else {
		goto L638
	}
L638:
	;
	v10317 = v241
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
	v1879 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
	v2024 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
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
		v10317 = v241
		goto L5
	} else {
		goto L695
	}
L693:
	;
	goto L694
L694:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
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
		v10317 = v241
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
	v2183 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2183)+500))
	v2190 = F_errstart(m, v2184&int32(2)+int32(19), int32(541439))
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
		v10317 = v241
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
	F_errmsg(m, int32(390916), v27+int32(112))
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
	F_errfinish(m, int32(26734), int32(765), int32(355297))
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L66
	} else {
		goto L739
	}
L739:
	;
	v10317 = v241
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
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
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
		v10152 = v2509
		goto L7
	} else {
		goto L857
	}
L745:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
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
		v10152 = v2233
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
		v10152 = v2241
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
	v2282 = int32(360049)
	v2285 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1369])))
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
	v10152 = v2241
	goto L7
L756:
	;
	goto L757
L757:
	;
	v2248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v2248 != 0 {
		v10152 = v2245
		goto L7
	} else {
		goto L758
	}
L758:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v2249 == int32(0) {
		v10152 = v2245
		goto L7
	} else {
		goto L759
	}
L759:
	;
	v2252 = int32(365770)
	v2255 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1371])))
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
		v10152 = v2245
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
	v2321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2315<<(uint(int32(1))%32))+uint32(_consts[1372]))))
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
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
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
		v10152 = v2224
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
		v10152 = v2337
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
		v10152 = v2345
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
	v2386 = int32(360049)
	v2389 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1369])))
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
	v10152 = v2345
	goto L7
L799:
	;
	goto L800
L800:
	;
	v2352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v2352 != 0 {
		v10152 = v2349
		goto L7
	} else {
		goto L801
	}
L801:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v2353 == int32(0) {
		v10152 = v2349
		goto L7
	} else {
		goto L802
	}
L802:
	;
	v2356 = int32(365770)
	v2359 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1371])))
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
		v10152 = v2349
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
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
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
		v10152 = v2423
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
		v10152 = v2431
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
	v2472 = int32(360049)
	v2475 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1369])))
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
	v10152 = v2431
	goto L7
L833:
	;
	goto L834
L834:
	;
	v2438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v2438 != 0 {
		v10152 = v2435
		goto L7
	} else {
		goto L835
	}
L835:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v2439 == int32(0) {
		v10152 = v2435
		goto L7
	} else {
		goto L836
	}
L836:
	;
	v2442 = int32(365770)
	v2445 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1371])))
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
		v10152 = v2435
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
	v2520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
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
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v2521 == int32(0) {
		goto L864
	} else {
		goto L866
	}
L866:
	;
	v2524 = int32(25989)
	v2527 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1373])))
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
	v10288 = v2506
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
	v10288 = v2651
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
	v10317 = v241
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
	v10317 = v241
	goto L5
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2687
	v10317 = v241
	goto L5
L900:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2709
	v10317 = v241
	goto L5
L901:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2716
	v10317 = v241
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
	v10317 = v241
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
	v2828 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
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
	v2855 = F_read_sql_construct(m, int32(59), v2842, v2842, int32(538734), v2842, v2842, v2842, v27+int32(4752), v2842, v27+int32(4736), v27+int32(4732), l1)
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
	v2860 = *(*int32)(unsafe.Add(mBase, _consts[1374]))
	*(*int32)(unsafe.Add(mBase, uint32(v2858))) = v2860
	v2863 = *(*int32)(unsafe.Add(mBase, _consts[1375]))
	*(*int32)(unsafe.Add(mBase, uint32(v2858)+3)) = v2863
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(v2775)+12))
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v2865)))
	v2868 = v2866 + int32(1)
	if v2866&int32(3) == int32(0) {
		v2892 = v2866
		goto L924
	} else {
		goto L925
	}
L922:
	;
	if v2866 == v2868 {
		goto L940
	} else {
		goto L941
	}
L923:
	;
	v2925 = v2917 - v2866
	goto L922
L924:
	;
	v2896 = v2892
	goto L933
L925:
	;
	v2876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2866))))
	if v2876 == int32(0) {
		goto L926
	} else {
		goto L927
	}
L926:
	;
	v2925 = int32(0)
	goto L922
L927:
	;
	goto L928
L928:
	;
	v2881 = v2866
	goto L929
L929:
	;
	v2885 = v2881 + int32(1)
	if v2885&int32(3) == int32(0) {
		v2892 = v2885
		goto L924
	} else {
		goto L931
	}
L930:
	;
	v2917 = v2885
	goto L923
L931:
	;
	v2890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2885))))
	if v2890 != 0 {
		v2881 = v2885
		goto L929
	} else {
		goto L932
	}
L932:
	;
	goto L930
L933:
	;
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2896)))
	v2905 = int32(-2139062144)
	if (int32(16843008)-v2902|v2902)&v2905 == v2905 {
		v2896 = v2896 + int32(4)
		goto L933
	} else {
		goto L935
	}
L934:
	;
	v2911 = v2896
	goto L936
L935:
	;
	goto L934
L936:
	;
	v2915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2911))))
	if v2915 != 0 {
		v2911 = v2911 + int32(1)
		goto L936
	} else {
		goto L938
	}
L937:
	;
	v2917 = v2911
	goto L923
L938:
	;
	goto L937
L939:
	;
	v3071 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1355])))
	if v3071 == int32(1) {
		goto L985
	} else {
		goto L986
	}
L940:
	;
	goto L939
L941:
	;
	v2929 = v2866 + v2925
	if base.Ui32(v2868-v2929) <= base.Ui32(int32(0)-v2925<<(uint(int32(1))%32)) {
		goto L942
	} else {
		goto L943
	}
L942:
	;
	v2936 = F___memcpy(m, v2866, v2868, v2925)
	mBase = m.M
	goto L939
L943:
	;
	goto L944
L944:
	;
	v2939 = (v2866 ^ v2868) & int32(3)
	if base.Ui32(v2866) < base.Ui32(v2868) {
		goto L947
	} else {
		goto L948
	}
L945:
	;
	if v3041 == int32(0) {
		goto L940
	} else {
		goto L981
	}
L946:
	;
	if base.Ui32(v3019) <= base.Ui32(int32(3)) {
		v3040 = v3018
		v3041 = v3019
		v3042 = v3020
		goto L945
	} else {
		goto L977
	}
L947:
	;
	if v2939 != 0 {
		goto L950
	} else {
		goto L951
	}
L948:
	;
	goto L949
L949:
	;
	if v2939 != 0 {
		v3001 = v2925
		goto L960
	} else {
		goto L961
	}
L950:
	;
	v3040 = v2868
	v3041 = v2925
	v3042 = v2866
	goto L945
L951:
	;
	goto L952
L952:
	;
	if v2866&int32(3) == int32(0) {
		goto L953
	} else {
		goto L954
	}
L953:
	;
	v3018 = v2868
	v3019 = v2925
	v3020 = v2866
	goto L946
L954:
	;
	goto L955
L955:
	;
	v2946 = v2868
	v2947 = v2925
	v2948 = v2866
	goto L956
L956:
	;
	if v2947 == int32(0) {
		goto L940
	} else {
		goto L958
	}
L957:
	;
	v3018 = v2955
	v3019 = v2957
	v3020 = v2959
	goto L946
L958:
	;
	v2952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2946))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2948))) = uint8(v2952)
	v2954 = int32(1)
	v2955 = v2946 + v2954
	v2957 = v2947 - v2954
	v2959 = v2948 + v2954
	if v2959&int32(3) != 0 {
		v2946 = v2955
		v2947 = v2957
		v2948 = v2959
		goto L956
	} else {
		goto L959
	}
L959:
	;
	goto L957
L960:
	;
	if v3001 == int32(0) {
		goto L940
	} else {
		goto L973
	}
L961:
	;
	if v2929&int32(3) != 0 {
		goto L962
	} else {
		goto L963
	}
L962:
	;
	v2966 = v2925
	goto L965
L963:
	;
	v2981 = v2925
	goto L964
L964:
	;
	if base.Ui32(v2981) <= base.Ui32(int32(3)) {
		v3001 = v2981
		goto L960
	} else {
		goto L969
	}
L965:
	;
	if v2966 == int32(0) {
		goto L940
	} else {
		goto L967
	}
L966:
	;
	v2981 = v2972
	goto L964
L967:
	;
	v2972 = v2966 - int32(1)
	v2973 = v2866 + v2972
	v2975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2868+v2972))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2973))) = uint8(v2975)
	if v2973&int32(3) != 0 {
		v2966 = v2972
		goto L965
	} else {
		goto L968
	}
L968:
	;
	goto L966
L969:
	;
	v2988 = v2981
	goto L970
L970:
	;
	v2992 = v2988 - int32(4)
	v2995 = *(*int32)(unsafe.Add(mBase, uint32(v2868+v2992)))
	*(*int32)(unsafe.Add(mBase, uint32(v2866+v2992))) = v2995
	if base.Ui32(int32(3)) < base.Ui32(v2992) {
		v2988 = v2992
		goto L970
	} else {
		goto L972
	}
L971:
	;
	v3001 = v2992
	goto L960
L972:
	;
	goto L971
L973:
	;
	v3008 = v3001
	goto L974
L974:
	;
	v3012 = v3008 - int32(1)
	v3015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2868+v3012))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2866+v3012))) = uint8(v3015)
	if v3012 != 0 {
		v3008 = v3012
		goto L974
	} else {
		goto L976
	}
L975:
	;
	goto L940
L976:
	;
	goto L975
L977:
	;
	v3025 = v3018
	v3026 = v3019
	v3027 = v3020
	goto L978
L978:
	;
	v3029 = *(*int32)(unsafe.Add(mBase, uint32(v3025)))
	*(*int32)(unsafe.Add(mBase, uint32(v3027))) = v3029
	v3031 = int32(4)
	v3032 = v3025 + v3031
	v3034 = v3027 + v3031
	v3036 = v3026 - v3031
	if base.Ui32(int32(3)) < base.Ui32(v3036) {
		v3025 = v3032
		v3026 = v3036
		v3027 = v3034
		goto L978
	} else {
		goto L980
	}
L979:
	;
	v3040 = v3032
	v3041 = v3036
	v3042 = v3034
	goto L945
L980:
	;
	goto L979
L981:
	;
	v3047 = v3040
	v3048 = v3041
	v3049 = v3042
	goto L982
L982:
	;
	v3051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3047))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3049))) = uint8(v3051)
	v3053 = int32(1)
	v3058 = v3048 - v3053
	if v3058 != 0 {
		v3047 = v3047 + v3053
		v3048 = v3058
		v3049 = v3049 + v3053
		goto L982
	} else {
		goto L984
	}
L983:
	;
	goto L940
L984:
	;
	goto L983
L985:
	;
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v2775)+12))
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(v3074)+4))
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v3074)))
	v3077 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1376])))
	v3078 = int32(4470752)
	v3079 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v3082 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3082
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v3077 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = int32(6372)
	v3090 = int32(4463656)
	v3091 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4784)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = v3091
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1380]))) = v27 + int32(4768)
	v3100 = F_raw_parser(m, v3076, v3075)
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L66
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v2775
	v10317 = v241
	goto L5
L988:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3079
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v3105
	goto L987
L989:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3115))) = int32(24)
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v3120 = int32(0)
	if v3119 < v3120 {
		v3163 = v3120
		goto L991
	} else {
		goto L992
	}
L990:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3115)+4)) = v3163
	v3168 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v3168)+520))
	v3171 = v3169 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+520)) = v3171
	*(*int32)(unsafe.Add(mBase, uint32(v3115)+8)) = v3171
	F_plpgsql_push_back_token(m, int32(289), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3180 = m.ExcPending
	if v3180 != 0 {
		goto L66
	} else {
		goto L1004
	}
L991:
	;
	goto L990
L992:
	;
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+60))
	if v3126 == int32(0) {
		v3163 = v3120
		goto L991
	} else {
		goto L993
	}
L993:
	;
	v3129 = v3119 + v3126
	v3130 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+188))
	if base.Ui32(v3130) <= base.Ui32(v3129) {
		goto L995
	} else {
		goto L996
	}
L994:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+196))
	if v3139 == int32(0) {
		v3163 = v3140
		goto L991
	} else {
		goto L998
	}
L995:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+192))
	v3139 = v3132
	goto L994
L996:
	;
	goto L997
L997:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3125)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3125)+188)) = v3126
	v3137 = F_strchr(m, v3126, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3125)+192)) = v3137
	v3139 = v3137
	goto L994
L998:
	;
	if base.Ui32(v3129) <= base.Ui32(v3139) {
		v3163 = v3140
		goto L991
	} else {
		goto L999
	}
L999:
	;
	v3144 = v3139
	v3146 = v3140
	goto L1000
L1000:
	;
	v3149 = int32(1)
	v3150 = v3146 + v3149
	*(*int32)(unsafe.Add(mBase, uint32(v3125)+196)) = v3150
	v3153 = v3144 + v3149
	*(*int32)(unsafe.Add(mBase, uint32(v3125)+188)) = v3153
	v3156 = F_strchr(m, v3153, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3125)+192)) = v3156
	if v3156 == int32(0) {
		v3163 = v3150
		goto L991
	} else {
		goto L1002
	}
L1001:
	;
	v3163 = v3150
	goto L991
L1002:
	;
	if base.Ui32(v3156) < base.Ui32(v3129) {
		v3144 = v3156
		v3146 = v3150
		goto L1000
	} else {
		goto L1003
	}
L1003:
	;
	goto L1001
L1004:
	;
	v3182 = int32(0)
	v3194 = F_read_sql_construct(m, int32(59), v3182, v3182, int32(538734), v3182, v3182, int32(1), v3182, v3182, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		goto L66
	} else {
		goto L1005
	}
L1005:
	;
	v3196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3115)+16)) = uint8(v3196)
	*(*int32)(unsafe.Add(mBase, uint32(v3115)+12)) = v3194
	v3200 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3200)+524)) = uint8(v3196)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3115
	v10317 = v241
	goto L5
L1006:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3205))) = int32(24)
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v3210 = int32(0)
	if v3209 < v3210 {
		v3253 = v3210
		goto L1008
	} else {
		goto L1009
	}
L1007:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+4)) = v3253
	v3258 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v3258)+520))
	v3261 = v3259 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3258)+520)) = v3261
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+8)) = v3261
	F_plpgsql_push_back_token(m, int32(309), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3270 = m.ExcPending
	if v3270 != 0 {
		goto L66
	} else {
		goto L1021
	}
L1008:
	;
	goto L1007
L1009:
	;
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+60))
	if v3216 == int32(0) {
		v3253 = v3210
		goto L1008
	} else {
		goto L1010
	}
L1010:
	;
	v3219 = v3209 + v3216
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+188))
	if base.Ui32(v3220) <= base.Ui32(v3219) {
		goto L1012
	} else {
		goto L1013
	}
L1011:
	;
	v3230 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+196))
	if v3229 == int32(0) {
		v3253 = v3230
		goto L1008
	} else {
		goto L1015
	}
L1012:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+192))
	v3229 = v3222
	goto L1011
L1013:
	;
	goto L1014
L1014:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+188)) = v3216
	v3227 = F_strchr(m, v3216, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+192)) = v3227
	v3229 = v3227
	goto L1011
L1015:
	;
	if base.Ui32(v3219) <= base.Ui32(v3229) {
		v3253 = v3230
		goto L1008
	} else {
		goto L1016
	}
L1016:
	;
	v3234 = v3229
	v3236 = v3230
	goto L1017
L1017:
	;
	v3239 = int32(1)
	v3240 = v3236 + v3239
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+196)) = v3240
	v3243 = v3234 + v3239
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+188)) = v3243
	v3246 = F_strchr(m, v3243, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3215)+192)) = v3246
	if v3246 == int32(0) {
		v3253 = v3240
		goto L1008
	} else {
		goto L1019
	}
L1018:
	;
	v3253 = v3240
	goto L1008
L1019:
	;
	if base.Ui32(v3246) < base.Ui32(v3219) {
		v3234 = v3246
		v3236 = v3240
		goto L1017
	} else {
		goto L1020
	}
L1020:
	;
	goto L1018
L1021:
	;
	v3272 = int32(0)
	v3284 = F_read_sql_construct(m, int32(59), v3272, v3272, int32(538734), v3272, v3272, int32(1), v3272, v3272, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3285 = m.ExcPending
	if v3285 != 0 {
		goto L66
	} else {
		goto L1022
	}
L1022:
	;
	v3286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3205)+16)) = uint8(v3286)
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+12)) = v3284
	v3290 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v3291 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3290)+524)) = uint8(v3291)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3205
	v10317 = v241
	goto L5
L1023:
	;
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	if v3298 == int32(0) {
		goto L45
	} else {
		goto L1026
	}
L1024:
	;
	v3309 = int32(3)
	goto L1025
L1025:
	;
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_check_assignable(m, v3310, v3311, l1)
	mBase = m.M
	v3313 = m.ExcPending
	if v3313 != 0 {
		goto L66
	} else {
		goto L1028
	}
L1026:
	;
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3298)+4))
	if base.Ui32(int32(3)) <= base.Ui32(v3301-int32(1)) {
		goto L45
	} else {
		goto L1027
	}
L1027:
	;
	v3309 = v3301 + int32(2)
	goto L1025
L1028:
	;
	v3315 = F_palloc0(m, int32(20))
	mBase = m.M
	v3316 = m.ExcPending
	if v3316 != 0 {
		goto L66
	} else {
		goto L1029
	}
L1029:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3315))) = int32(1)
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v3320 = int32(0)
	if v3319 < v3320 {
		v3363 = v3320
		goto L1031
	} else {
		goto L1032
	}
L1030:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3315)+4)) = v3363
	v3368 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v3369 = *(*int32)(unsafe.Add(mBase, uint32(v3368)+520))
	v3371 = v3369 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3368)+520)) = v3371
	*(*int32)(unsafe.Add(mBase, uint32(v3315)+8)) = v3371
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3375 = *(*int32)(unsafe.Add(mBase, uint32(v3374)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3315)+12)) = v3375
	F_plpgsql_push_back_token(m, int32(277), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L66
	} else {
		goto L1044
	}
L1031:
	;
	goto L1030
L1032:
	;
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3326 = *(*int32)(unsafe.Add(mBase, uint32(v3325)+60))
	if v3326 == int32(0) {
		v3363 = v3320
		goto L1031
	} else {
		goto L1033
	}
L1033:
	;
	v3329 = v3319 + v3326
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v3325)+188))
	if base.Ui32(v3330) <= base.Ui32(v3329) {
		goto L1035
	} else {
		goto L1036
	}
L1034:
	;
	v3340 = *(*int32)(unsafe.Add(mBase, uint32(v3325)+196))
	if v3339 == int32(0) {
		v3363 = v3340
		goto L1031
	} else {
		goto L1038
	}
L1035:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, uint32(v3325)+192))
	v3339 = v3332
	goto L1034
L1036:
	;
	goto L1037
L1037:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3325)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3325)+188)) = v3326
	v3337 = F_strchr(m, v3326, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3325)+192)) = v3337
	v3339 = v3337
	goto L1034
L1038:
	;
	if base.Ui32(v3329) <= base.Ui32(v3339) {
		v3363 = v3340
		goto L1031
	} else {
		goto L1039
	}
L1039:
	;
	v3344 = v3339
	v3346 = v3340
	goto L1040
L1040:
	;
	v3349 = int32(1)
	v3350 = v3346 + v3349
	*(*int32)(unsafe.Add(mBase, uint32(v3325)+196)) = v3350
	v3353 = v3344 + v3349
	*(*int32)(unsafe.Add(mBase, uint32(v3325)+188)) = v3353
	v3356 = F_strchr(m, v3353, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3325)+192)) = v3356
	if v3356 == int32(0) {
		v3363 = v3350
		goto L1031
	} else {
		goto L1042
	}
L1041:
	;
	v3363 = v3350
	goto L1031
L1042:
	;
	if base.Ui32(v3356) < base.Ui32(v3329) {
		v3344 = v3356
		v3346 = v3350
		goto L1040
	} else {
		goto L1043
	}
L1043:
	;
	goto L1041
L1044:
	;
	v3385 = int32(0)
	v3396 = F_read_sql_construct(m, int32(59), v3385, v3385, int32(538734), v3309, v3385, int32(1), v3385, v3385, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v3397 = m.ExcPending
	if v3397 != 0 {
		goto L66
	} else {
		goto L1045
	}
L1045:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3315)+16)) = v3396
	v3400 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v3401 = *(*int32)(unsafe.Add(mBase, uint32(v3400)))
	if v3401 != 0 {
		goto L1046
	} else {
		goto L1047
	}
L1046:
	;
	v3405 = int32(-1)
	v3406 = int32(0)
	goto L1048
L1047:
	;
	v3403 = *(*int32)(unsafe.Add(mBase, uint32(v3400)+4))
	v3405 = v3403
	v3406 = int32(1)
	goto L1048
L1048:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3396)+20)) = uint8(v3406)
	*(*int32)(unsafe.Add(mBase, uint32(v3396)+16)) = v3405
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3315
	v10317 = v241
	goto L5
L1049:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3411))) = int32(19)
	v3416 = v142 - int32(16)
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(v3416)))
	v3418 = int32(0)
	if v3417 < v3418 {
		v3461 = v3418
		goto L1051
	} else {
		goto L1052
	}
L1050:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3411)+4)) = v3461
	v3466 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v3467 = *(*int32)(unsafe.Add(mBase, uint32(v3466)+520))
	v3469 = v3467 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3466)+520)) = v3469
	*(*int32)(unsafe.Add(mBase, uint32(v3411)+8)) = v3469
	v3474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132-int32(48)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3411)+12)) = uint8(v3474)
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v3411)+16)) = v3478
	if v3478 == int32(0) {
		goto L1064
	} else {
		goto L1065
	}
L1051:
	;
	goto L1050
L1052:
	;
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v3423)+60))
	if v3424 == int32(0) {
		v3461 = v3418
		goto L1051
	} else {
		goto L1053
	}
L1053:
	;
	v3427 = v3417 + v3424
	v3428 = *(*int32)(unsafe.Add(mBase, uint32(v3423)+188))
	if base.Ui32(v3428) <= base.Ui32(v3427) {
		goto L1055
	} else {
		goto L1056
	}
L1054:
	;
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(v3423)+196))
	if v3437 == int32(0) {
		v3461 = v3438
		goto L1051
	} else {
		goto L1058
	}
L1055:
	;
	v3430 = *(*int32)(unsafe.Add(mBase, uint32(v3423)+192))
	v3437 = v3430
	goto L1054
L1056:
	;
	goto L1057
L1057:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3423)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3423)+188)) = v3424
	v3435 = F_strchr(m, v3424, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3423)+192)) = v3435
	v3437 = v3435
	goto L1054
L1058:
	;
	if base.Ui32(v3427) <= base.Ui32(v3437) {
		v3461 = v3438
		goto L1051
	} else {
		goto L1059
	}
L1059:
	;
	v3442 = v3437
	v3444 = v3438
	goto L1060
L1060:
	;
	v3447 = int32(1)
	v3448 = v3444 + v3447
	*(*int32)(unsafe.Add(mBase, uint32(v3423)+196)) = v3448
	v3451 = v3442 + v3447
	*(*int32)(unsafe.Add(mBase, uint32(v3423)+188)) = v3451
	v3454 = F_strchr(m, v3451, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3423)+192)) = v3454
	if v3454 == int32(0) {
		v3461 = v3448
		goto L1051
	} else {
		goto L1062
	}
L1061:
	;
	v3461 = v3448
	goto L1051
L1062:
	;
	if base.Ui32(v3454) < base.Ui32(v3427) {
		v3442 = v3454
		v3444 = v3448
		goto L1060
	} else {
		goto L1063
	}
L1063:
	;
	goto L1061
L1064:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3411
	v10317 = v241
	goto L5
L1065:
	;
	v3482 = *(*int32)(unsafe.Add(mBase, uint32(v3478)+4))
	if v3482 <= int32(0) {
		goto L1064
	} else {
		goto L1066
	}
L1066:
	;
	v3485 = int32(0)
	if v3485 < v3482 {
		goto L1067
	} else {
		goto L1068
	}
L1067:
	;
	v3489 = v3482
	goto L1069
L1068:
	;
	v3489 = v3485
	goto L1069
L1069:
	;
	v3490 = *(*int32)(unsafe.Add(mBase, uint32(v3478)+12))
	v3494 = v3485
	goto L1070
L1070:
	;
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v3490+v3494<<(uint(int32(2))%32))))
	v3519 = *(*int32)(unsafe.Add(mBase, uint32(v3518)))
	if base.Ui32(int32(10)) <= base.Ui32(v3519-int32(3)) {
		goto L1074
	} else {
		goto L1075
	}
L1071:
	;
	goto L1064
L1072:
	;
	v3609 = v3494 + int32(1)
	if v3609 != v3489 {
		v3494 = v3609
		goto L1070
	} else {
		goto L1101
	}
L1073:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v3595 = m.ExcPending
	if v3595 != 0 {
		goto L66
	} else {
		goto L1098
	}
L1074:
	;
	switch v3519 {
	case 0, 1:
		goto L1077
	case 2:
		goto L1072
	default:
		goto L1073
	}
L1075:
	;
	goto L1076
L1076:
	;
	if v3474&int32(1) != 0 {
		goto L1072
	} else {
		goto L1088
	}
L1077:
	;
	if v3474&int32(1) == int32(0) {
		goto L1072
	} else {
		goto L1078
	}
L1078:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v3531 = m.ExcPending
	if v3531 != 0 {
		goto L66
	} else {
		goto L1079
	}
L1079:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L66
	} else {
		goto L1080
	}
L1080:
	;
	v3535 = *(*int32)(unsafe.Add(mBase, uint32(v3518)))
	if base.Ui32(int32(12)) < base.Ui32(v3535) {
		goto L1082
	} else {
		goto L1083
	}
L1081:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v3544
	F_errmsg(m, int32(515781), v27+int32(144))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		goto L66
	} else {
		goto L1085
	}
L1082:
	;
	v3544 = int32(240184)
	goto L1081
L1083:
	;
	goto L1084
L1084:
	;
	v3543 = *(*int32)(unsafe.Add(mBase, uint32(v3535<<(uint(int32(2))%32))+uint32(_consts[1381])))
	v3544 = v3543
	goto L1081
L1085:
	;
	v3551 = *(*int32)(unsafe.Add(mBase, uint32(v3416)))
	v3552 = F_plpgsql_scanner_errposition(m, v3551, l1)
	mBase = m.M
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L66
	} else {
		goto L1086
	}
L1086:
	;
	F_errfinish(m, int32(26734), int32(1042), int32(355297))
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
		goto L66
	} else {
		goto L1087
	}
L1087:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1088:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L66
	} else {
		goto L1089
	}
L1089:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3567 = m.ExcPending
	if v3567 != 0 {
		goto L66
	} else {
		goto L1090
	}
L1090:
	;
	v3568 = *(*int32)(unsafe.Add(mBase, uint32(v3518)))
	if base.Ui32(int32(12)) < base.Ui32(v3568) {
		goto L1092
	} else {
		goto L1093
	}
L1091:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v3577
	F_errmsg(m, int32(515703), v27+int32(160))
	mBase = m.M
	v3583 = m.ExcPending
	if v3583 != 0 {
		goto L66
	} else {
		goto L1095
	}
L1092:
	;
	v3577 = int32(240184)
	goto L1091
L1093:
	;
	goto L1094
L1094:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(v3568<<(uint(int32(2))%32))+uint32(_consts[1381])))
	v3577 = v3576
	goto L1091
L1095:
	;
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v3416)))
	v3585 = F_plpgsql_scanner_errposition(m, v3584, l1)
	mBase = m.M
	v3586 = m.ExcPending
	if v3586 != 0 {
		goto L66
	} else {
		goto L1096
	}
L1096:
	;
	F_errfinish(m, int32(26734), int32(1060), int32(355297))
	mBase = m.M
	v3591 = m.ExcPending
	if v3591 != 0 {
		goto L66
	} else {
		goto L1097
	}
L1097:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1098:
	;
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v3518)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v3596
	F_errmsg_internal(m, int32(479949), v27+int32(128))
	mBase = m.M
	v3602 = m.ExcPending
	if v3602 != 0 {
		goto L66
	} else {
		goto L1099
	}
L1099:
	;
	F_errfinish(m, int32(26734), int32(1067), int32(355297))
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		goto L66
	} else {
		goto L1100
	}
L1100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1101:
	;
	goto L1071
L1102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3646
	v10317 = v241
	goto L5
L1103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3655
	v10317 = v241
	goto L5
L1104:
	;
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	v3664 = *(*int32)(unsafe.Add(mBase, uint32(v3663)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3659)+4)) = v3664
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v3659))) = v3666
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v3659
	v10317 = v241
	goto L5
L1105:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(286899))
	mBase = m.M
	v4082 = m.ExcPending
	if v4082 != 0 {
		goto L66
	} else {
		goto L1258
	}
L1106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(6)
	v10317 = v241
	goto L5
L1107:
	;
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v4044 == int32(0) {
		goto L1105
	} else {
		goto L1248
	}
L1108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(12)
	v10317 = v241
	goto L5
L1109:
	;
	v4015 = int32(375163)
	v4018 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1382])))
	v4019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3983))))
	if v4019 == int32(0) {
		v4038 = v4018
		v4039 = v4019
		goto L1240
	} else {
		goto L1241
	}
L1110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(11)
	v10317 = v241
	goto L5
L1111:
	;
	v3983 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v3983 == int32(0) {
		goto L1105
	} else {
		goto L1229
	}
L1112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(10)
	v10317 = v241
	goto L5
L1113:
	;
	v3954 = int32(62665)
	v3957 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1383])))
	v3958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3922))))
	if v3958 == int32(0) {
		v3977 = v3957
		v3978 = v3958
		goto L1221
	} else {
		goto L1222
	}
L1114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(9)
	v10317 = v241
	goto L5
L1115:
	;
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v3922 == int32(0) {
		goto L1105
	} else {
		goto L1210
	}
L1116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(8)
	v10317 = v241
	goto L5
L1117:
	;
	v3893 = int32(372686)
	v3896 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1384])))
	v3897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3861))))
	if v3897 == int32(0) {
		v3916 = v3896
		v3917 = v3897
		goto L1202
	} else {
		goto L1203
	}
L1118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(7)
	v10317 = v241
	goto L5
L1119:
	;
	v3861 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v3861 == int32(0) {
		goto L1105
	} else {
		goto L1191
	}
L1120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(3)
	v10317 = v241
	goto L5
L1121:
	;
	v3832 = int32(59025)
	v3835 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1385])))
	v3836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3800))))
	if v3836 == int32(0) {
		v3855 = v3835
		v3856 = v3836
		goto L1183
	} else {
		goto L1184
	}
L1122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(5)
	v10317 = v241
	goto L5
L1123:
	;
	v3800 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v3800 == int32(0) {
		goto L1105
	} else {
		goto L1172
	}
L1124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(4)
	v10317 = v241
	goto L5
L1125:
	;
	v3771 = int32(301141)
	v3774 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1386])))
	v3775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3739))))
	if v3775 == int32(0) {
		v3794 = v3774
		v3795 = v3775
		goto L1164
	} else {
		goto L1165
	}
L1126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(2)
	v10317 = v241
	goto L5
L1127:
	;
	v3739 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v3739 == int32(0) {
		goto L1105
	} else {
		goto L1153
	}
L1128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(1)
	v10317 = v241
	goto L5
L1129:
	;
	v3710 = int32(428491)
	v3713 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1387])))
	v3714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3678))))
	if v3714 == int32(0) {
		v3733 = v3713
		v3734 = v3714
		goto L1145
	} else {
		goto L1146
	}
L1130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L1131:
	;
	v3677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v3677 != 0 {
		goto L1105
	} else {
		goto L1133
	}
L1132:
	;
	switch v3673 - int32(277) {
	case 0:
		goto L1131
	default:
		goto L1105
	case 18:
		goto L1118
	case 22:
		goto L1116
	case 62:
		goto L1112
	case 73:
		goto L1126
	case 74:
		goto L1114
	case 75:
		goto L1120
	case 76:
		goto L1124
	case 77:
		goto L1122
	case 78:
		goto L1128
	case 85:
		goto L1106
	case 88:
		goto L1130
	case 91:
		goto L1108
	case 98:
		goto L1110
	}
L1133:
	;
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v3678 == int32(0) {
		goto L1105
	} else {
		goto L1134
	}
L1134:
	;
	v3681 = int32(86519)
	v3684 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1388])))
	v3685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3678))))
	if v3685 == int32(0) {
		v3704 = v3684
		v3705 = v3685
		goto L1136
	} else {
		goto L1137
	}
L1135:
	;
	if v3705-v3704 != 0 {
		goto L1129
	} else {
		goto L1143
	}
L1136:
	;
	goto L1135
L1137:
	;
	if v3684 != v3685 {
		v3704 = v3684
		v3705 = v3685
		goto L1136
	} else {
		goto L1138
	}
L1138:
	;
	v3689 = v3678
	v3690 = v3681
	goto L1139
L1139:
	;
	v3693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3690)+1)))
	v3694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3689)+1)))
	if v3694 == int32(0) {
		v3704 = v3693
		v3705 = v3694
		goto L1136
	} else {
		goto L1141
	}
L1140:
	;
	v3704 = v3693
	v3705 = v3694
	goto L1136
L1141:
	;
	v3697 = int32(1)
	if v3693 == v3694 {
		v3689 = v3689 + v3697
		v3690 = v3690 + v3697
		goto L1139
	} else {
		goto L1142
	}
L1142:
	;
	goto L1140
L1143:
	;
	goto L1130
L1144:
	;
	if v3734-v3733 != 0 {
		goto L1127
	} else {
		goto L1152
	}
L1145:
	;
	goto L1144
L1146:
	;
	if v3713 != v3714 {
		v3733 = v3713
		v3734 = v3714
		goto L1145
	} else {
		goto L1147
	}
L1147:
	;
	v3718 = v3678
	v3719 = v3710
	goto L1148
L1148:
	;
	v3722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3719)+1)))
	v3723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3718)+1)))
	if v3723 == int32(0) {
		v3733 = v3722
		v3734 = v3723
		goto L1145
	} else {
		goto L1150
	}
L1149:
	;
	v3733 = v3722
	v3734 = v3723
	goto L1145
L1150:
	;
	v3726 = int32(1)
	if v3722 == v3723 {
		v3718 = v3718 + v3726
		v3719 = v3719 + v3726
		goto L1148
	} else {
		goto L1151
	}
L1151:
	;
	goto L1149
L1152:
	;
	goto L1128
L1153:
	;
	v3742 = int32(59046)
	v3745 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1389])))
	v3746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3739))))
	if v3746 == int32(0) {
		v3765 = v3745
		v3766 = v3746
		goto L1155
	} else {
		goto L1156
	}
L1154:
	;
	if v3766-v3765 != 0 {
		goto L1125
	} else {
		goto L1162
	}
L1155:
	;
	goto L1154
L1156:
	;
	if v3745 != v3746 {
		v3765 = v3745
		v3766 = v3746
		goto L1155
	} else {
		goto L1157
	}
L1157:
	;
	v3750 = v3739
	v3751 = v3742
	goto L1158
L1158:
	;
	v3754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3751)+1)))
	v3755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3750)+1)))
	if v3755 == int32(0) {
		v3765 = v3754
		v3766 = v3755
		goto L1155
	} else {
		goto L1160
	}
L1159:
	;
	v3765 = v3754
	v3766 = v3755
	goto L1155
L1160:
	;
	v3758 = int32(1)
	if v3754 == v3755 {
		v3750 = v3750 + v3758
		v3751 = v3751 + v3758
		goto L1158
	} else {
		goto L1161
	}
L1161:
	;
	goto L1159
L1162:
	;
	goto L1126
L1163:
	;
	if v3795-v3794 != 0 {
		goto L1123
	} else {
		goto L1171
	}
L1164:
	;
	goto L1163
L1165:
	;
	if v3774 != v3775 {
		v3794 = v3774
		v3795 = v3775
		goto L1164
	} else {
		goto L1166
	}
L1166:
	;
	v3779 = v3739
	v3780 = v3771
	goto L1167
L1167:
	;
	v3783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3780)+1)))
	v3784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3779)+1)))
	if v3784 == int32(0) {
		v3794 = v3783
		v3795 = v3784
		goto L1164
	} else {
		goto L1169
	}
L1168:
	;
	v3794 = v3783
	v3795 = v3784
	goto L1164
L1169:
	;
	v3787 = int32(1)
	if v3783 == v3784 {
		v3779 = v3779 + v3787
		v3780 = v3780 + v3787
		goto L1167
	} else {
		goto L1170
	}
L1170:
	;
	goto L1168
L1171:
	;
	goto L1124
L1172:
	;
	v3803 = int32(88219)
	v3806 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1390])))
	v3807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3800))))
	if v3807 == int32(0) {
		v3826 = v3806
		v3827 = v3807
		goto L1174
	} else {
		goto L1175
	}
L1173:
	;
	if v3827-v3826 != 0 {
		goto L1121
	} else {
		goto L1181
	}
L1174:
	;
	goto L1173
L1175:
	;
	if v3806 != v3807 {
		v3826 = v3806
		v3827 = v3807
		goto L1174
	} else {
		goto L1176
	}
L1176:
	;
	v3811 = v3800
	v3812 = v3803
	goto L1177
L1177:
	;
	v3815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3812)+1)))
	v3816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3811)+1)))
	if v3816 == int32(0) {
		v3826 = v3815
		v3827 = v3816
		goto L1174
	} else {
		goto L1179
	}
L1178:
	;
	v3826 = v3815
	v3827 = v3816
	goto L1174
L1179:
	;
	v3819 = int32(1)
	if v3815 == v3816 {
		v3811 = v3811 + v3819
		v3812 = v3812 + v3819
		goto L1177
	} else {
		goto L1180
	}
L1180:
	;
	goto L1178
L1181:
	;
	goto L1122
L1182:
	;
	if v3856-v3855 != 0 {
		goto L1119
	} else {
		goto L1190
	}
L1183:
	;
	goto L1182
L1184:
	;
	if v3835 != v3836 {
		v3855 = v3835
		v3856 = v3836
		goto L1183
	} else {
		goto L1185
	}
L1185:
	;
	v3840 = v3800
	v3841 = v3832
	goto L1186
L1186:
	;
	v3844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3841)+1)))
	v3845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3840)+1)))
	if v3845 == int32(0) {
		v3855 = v3844
		v3856 = v3845
		goto L1183
	} else {
		goto L1188
	}
L1187:
	;
	v3855 = v3844
	v3856 = v3845
	goto L1183
L1188:
	;
	v3848 = int32(1)
	if v3844 == v3845 {
		v3840 = v3840 + v3848
		v3841 = v3841 + v3848
		goto L1186
	} else {
		goto L1189
	}
L1189:
	;
	goto L1187
L1190:
	;
	goto L1120
L1191:
	;
	v3864 = int32(373246)
	v3867 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1391])))
	v3868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3861))))
	if v3868 == int32(0) {
		v3887 = v3867
		v3888 = v3868
		goto L1193
	} else {
		goto L1194
	}
L1192:
	;
	if v3888-v3887 != 0 {
		goto L1117
	} else {
		goto L1200
	}
L1193:
	;
	goto L1192
L1194:
	;
	if v3867 != v3868 {
		v3887 = v3867
		v3888 = v3868
		goto L1193
	} else {
		goto L1195
	}
L1195:
	;
	v3872 = v3861
	v3873 = v3864
	goto L1196
L1196:
	;
	v3876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3873)+1)))
	v3877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3872)+1)))
	if v3877 == int32(0) {
		v3887 = v3876
		v3888 = v3877
		goto L1193
	} else {
		goto L1198
	}
L1197:
	;
	v3887 = v3876
	v3888 = v3877
	goto L1193
L1198:
	;
	v3880 = int32(1)
	if v3876 == v3877 {
		v3872 = v3872 + v3880
		v3873 = v3873 + v3880
		goto L1196
	} else {
		goto L1199
	}
L1199:
	;
	goto L1197
L1200:
	;
	goto L1118
L1201:
	;
	if v3917-v3916 != 0 {
		goto L1115
	} else {
		goto L1209
	}
L1202:
	;
	goto L1201
L1203:
	;
	if v3896 != v3897 {
		v3916 = v3896
		v3917 = v3897
		goto L1202
	} else {
		goto L1204
	}
L1204:
	;
	v3901 = v3861
	v3902 = v3893
	goto L1205
L1205:
	;
	v3905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3902)+1)))
	v3906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3901)+1)))
	if v3906 == int32(0) {
		v3916 = v3905
		v3917 = v3906
		goto L1202
	} else {
		goto L1207
	}
L1206:
	;
	v3916 = v3905
	v3917 = v3906
	goto L1202
L1207:
	;
	v3909 = int32(1)
	if v3905 == v3906 {
		v3901 = v3901 + v3909
		v3902 = v3902 + v3909
		goto L1205
	} else {
		goto L1208
	}
L1208:
	;
	goto L1206
L1209:
	;
	goto L1116
L1210:
	;
	v3925 = int32(373489)
	v3928 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1392])))
	v3929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3922))))
	if v3929 == int32(0) {
		v3948 = v3928
		v3949 = v3929
		goto L1212
	} else {
		goto L1213
	}
L1211:
	;
	if v3949-v3948 != 0 {
		goto L1113
	} else {
		goto L1219
	}
L1212:
	;
	goto L1211
L1213:
	;
	if v3928 != v3929 {
		v3948 = v3928
		v3949 = v3929
		goto L1212
	} else {
		goto L1214
	}
L1214:
	;
	v3933 = v3922
	v3934 = v3925
	goto L1215
L1215:
	;
	v3937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3934)+1)))
	v3938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3933)+1)))
	if v3938 == int32(0) {
		v3948 = v3937
		v3949 = v3938
		goto L1212
	} else {
		goto L1217
	}
L1216:
	;
	v3948 = v3937
	v3949 = v3938
	goto L1212
L1217:
	;
	v3941 = int32(1)
	if v3937 == v3938 {
		v3933 = v3933 + v3941
		v3934 = v3934 + v3941
		goto L1215
	} else {
		goto L1218
	}
L1218:
	;
	goto L1216
L1219:
	;
	goto L1114
L1220:
	;
	if v3978-v3977 != 0 {
		goto L1111
	} else {
		goto L1228
	}
L1221:
	;
	goto L1220
L1222:
	;
	if v3957 != v3958 {
		v3977 = v3957
		v3978 = v3958
		goto L1221
	} else {
		goto L1223
	}
L1223:
	;
	v3962 = v3922
	v3963 = v3954
	goto L1224
L1224:
	;
	v3966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3963)+1)))
	v3967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3962)+1)))
	if v3967 == int32(0) {
		v3977 = v3966
		v3978 = v3967
		goto L1221
	} else {
		goto L1226
	}
L1225:
	;
	v3977 = v3966
	v3978 = v3967
	goto L1221
L1226:
	;
	v3970 = int32(1)
	if v3966 == v3967 {
		v3962 = v3962 + v3970
		v3963 = v3963 + v3970
		goto L1224
	} else {
		goto L1227
	}
L1227:
	;
	goto L1225
L1228:
	;
	goto L1112
L1229:
	;
	v3986 = int32(374157)
	v3989 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1393])))
	v3990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3983))))
	if v3990 == int32(0) {
		v4009 = v3989
		v4010 = v3990
		goto L1231
	} else {
		goto L1232
	}
L1230:
	;
	if v4010-v4009 != 0 {
		goto L1109
	} else {
		goto L1238
	}
L1231:
	;
	goto L1230
L1232:
	;
	if v3989 != v3990 {
		v4009 = v3989
		v4010 = v3990
		goto L1231
	} else {
		goto L1233
	}
L1233:
	;
	v3994 = v3983
	v3995 = v3986
	goto L1234
L1234:
	;
	v3998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3995)+1)))
	v3999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3994)+1)))
	if v3999 == int32(0) {
		v4009 = v3998
		v4010 = v3999
		goto L1231
	} else {
		goto L1236
	}
L1235:
	;
	v4009 = v3998
	v4010 = v3999
	goto L1231
L1236:
	;
	v4002 = int32(1)
	if v3998 == v3999 {
		v3994 = v3994 + v4002
		v3995 = v3995 + v4002
		goto L1234
	} else {
		goto L1237
	}
L1237:
	;
	goto L1235
L1238:
	;
	goto L1110
L1239:
	;
	if v4039-v4038 != 0 {
		goto L1107
	} else {
		goto L1247
	}
L1240:
	;
	goto L1239
L1241:
	;
	if v4018 != v4019 {
		v4038 = v4018
		v4039 = v4019
		goto L1240
	} else {
		goto L1242
	}
L1242:
	;
	v4023 = v3983
	v4024 = v4015
	goto L1243
L1243:
	;
	v4027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4024)+1)))
	v4028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4023)+1)))
	if v4028 == int32(0) {
		v4038 = v4027
		v4039 = v4028
		goto L1240
	} else {
		goto L1245
	}
L1244:
	;
	v4038 = v4027
	v4039 = v4028
	goto L1240
L1245:
	;
	v4031 = int32(1)
	if v4027 == v4028 {
		v4023 = v4023 + v4031
		v4024 = v4024 + v4031
		goto L1243
	} else {
		goto L1246
	}
L1246:
	;
	goto L1244
L1247:
	;
	goto L1108
L1248:
	;
	v4047 = int32(346106)
	v4050 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1394])))
	v4051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4044))))
	if v4051 == int32(0) {
		v4070 = v4050
		v4071 = v4051
		goto L1250
	} else {
		goto L1251
	}
L1249:
	;
	if v4071-v4070 != 0 {
		goto L1105
	} else {
		goto L1257
	}
L1250:
	;
	goto L1249
L1251:
	;
	if v4050 != v4051 {
		v4070 = v4050
		v4071 = v4051
		goto L1250
	} else {
		goto L1252
	}
L1252:
	;
	v4055 = v4044
	v4056 = v4047
	goto L1253
L1253:
	;
	v4059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4056)+1)))
	v4060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4055)+1)))
	if v4060 == int32(0) {
		v4070 = v4059
		v4071 = v4060
		goto L1250
	} else {
		goto L1255
	}
L1254:
	;
	v4070 = v4059
	v4071 = v4060
	goto L1250
L1255:
	;
	v4063 = int32(1)
	if v4059 == v4060 {
		v4055 = v4055 + v4063
		v4056 = v4056 + v4063
		goto L1253
	} else {
		goto L1256
	}
L1256:
	;
	goto L1254
L1257:
	;
	goto L1106
L1258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1259:
	;
	v4089 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v4090 = m.ExcPending
	if v4090 != 0 {
		goto L66
	} else {
		goto L1260
	}
L1260:
	;
	if v4089 == int32(91) {
		goto L44
	} else {
		goto L1261
	}
L1261:
	;
	v4093 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v4094 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_check_assignable(m, v4093, v4094, l1)
	mBase = m.M
	v4096 = m.ExcPending
	if v4096 != 0 {
		goto L66
	} else {
		goto L1262
	}
L1262:
	;
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4097
	v10317 = v241
	goto L5
L1263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4106))) = int32(2)
	v4112 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(28))))
	v4113 = int32(0)
	if v4112 < v4113 {
		v4156 = v4113
		goto L1267
	} else {
		goto L1268
	}
L1266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4106)+4)) = v4156
	v4161 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v4162 = *(*int32)(unsafe.Add(mBase, uint32(v4161)+520))
	v4164 = v4162 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4161)+520)) = v4164
	*(*int32)(unsafe.Add(mBase, uint32(v4106)+8)) = v4164
	v4169 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(96))))
	*(*int32)(unsafe.Add(mBase, uint32(v4106)+12)) = v4169
	v4173 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(80))))
	*(*int32)(unsafe.Add(mBase, uint32(v4106)+16)) = v4173
	v4177 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v4106)+20)) = v4177
	v4181 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	*(*int32)(unsafe.Add(mBase, uint32(v4106)+24)) = v4181
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4106
	v10317 = v241
	goto L5
L1267:
	;
	goto L1266
L1268:
	;
	v4118 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4119 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+60))
	if v4119 == int32(0) {
		v4156 = v4113
		goto L1267
	} else {
		goto L1269
	}
L1269:
	;
	v4122 = v4112 + v4119
	v4123 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+188))
	if base.Ui32(v4123) <= base.Ui32(v4122) {
		goto L1271
	} else {
		goto L1272
	}
L1270:
	;
	v4133 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+196))
	if v4132 == int32(0) {
		v4156 = v4133
		goto L1267
	} else {
		goto L1274
	}
L1271:
	;
	v4125 = *(*int32)(unsafe.Add(mBase, uint32(v4118)+192))
	v4132 = v4125
	goto L1270
L1272:
	;
	goto L1273
L1273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4118)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4118)+188)) = v4119
	v4130 = F_strchr(m, v4119, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4118)+192)) = v4130
	v4132 = v4130
	goto L1270
L1274:
	;
	if base.Ui32(v4122) <= base.Ui32(v4132) {
		v4156 = v4133
		goto L1267
	} else {
		goto L1275
	}
L1275:
	;
	v4137 = v4132
	v4139 = v4133
	goto L1276
L1276:
	;
	v4142 = int32(1)
	v4143 = v4139 + v4142
	*(*int32)(unsafe.Add(mBase, uint32(v4118)+196)) = v4143
	v4146 = v4137 + v4142
	*(*int32)(unsafe.Add(mBase, uint32(v4118)+188)) = v4146
	v4149 = F_strchr(m, v4146, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4118)+192)) = v4149
	if v4149 == int32(0) {
		v4156 = v4143
		goto L1267
	} else {
		goto L1278
	}
L1277:
	;
	v4156 = v4143
	goto L1267
L1278:
	;
	if base.Ui32(v4149) < base.Ui32(v4122) {
		v4137 = v4149
		v4139 = v4143
		goto L1276
	} else {
		goto L1279
	}
L1279:
	;
	goto L1277
L1280:
	;
	v4191 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4192 = int32(0)
	if v4191 < v4192 {
		v4235 = v4192
		goto L1282
	} else {
		goto L1283
	}
L1281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4187))) = v4235
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+4)) = v4241
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4187)+8)) = v4243
	v4247 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(48))))
	v4248 = F_lappend(m, v4247, v4187)
	mBase = m.M
	v4249 = m.ExcPending
	if v4249 != 0 {
		goto L66
	} else {
		goto L1295
	}
L1282:
	;
	goto L1281
L1283:
	;
	v4197 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4198 = *(*int32)(unsafe.Add(mBase, uint32(v4197)+60))
	if v4198 == int32(0) {
		v4235 = v4192
		goto L1282
	} else {
		goto L1284
	}
L1284:
	;
	v4201 = v4191 + v4198
	v4202 = *(*int32)(unsafe.Add(mBase, uint32(v4197)+188))
	if base.Ui32(v4202) <= base.Ui32(v4201) {
		goto L1286
	} else {
		goto L1287
	}
L1285:
	;
	v4212 = *(*int32)(unsafe.Add(mBase, uint32(v4197)+196))
	if v4211 == int32(0) {
		v4235 = v4212
		goto L1282
	} else {
		goto L1289
	}
L1286:
	;
	v4204 = *(*int32)(unsafe.Add(mBase, uint32(v4197)+192))
	v4211 = v4204
	goto L1285
L1287:
	;
	goto L1288
L1288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4197)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4197)+188)) = v4198
	v4209 = F_strchr(m, v4198, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4197)+192)) = v4209
	v4211 = v4209
	goto L1285
L1289:
	;
	if base.Ui32(v4201) <= base.Ui32(v4211) {
		v4235 = v4212
		goto L1282
	} else {
		goto L1290
	}
L1290:
	;
	v4216 = v4211
	v4218 = v4212
	goto L1291
L1291:
	;
	v4221 = int32(1)
	v4222 = v4218 + v4221
	*(*int32)(unsafe.Add(mBase, uint32(v4197)+196)) = v4222
	v4225 = v4216 + v4221
	*(*int32)(unsafe.Add(mBase, uint32(v4197)+188)) = v4225
	v4228 = F_strchr(m, v4225, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4197)+192)) = v4228
	if v4228 == int32(0) {
		v4235 = v4222
		goto L1282
	} else {
		goto L1293
	}
L1292:
	;
	v4235 = v4222
	goto L1282
L1293:
	;
	if base.Ui32(v4228) < base.Ui32(v4201) {
		v4216 = v4228
		v4218 = v4222
		goto L1291
	} else {
		goto L1294
	}
L1294:
	;
	goto L1292
L1295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4248
	v10317 = v241
	goto L5
L1296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4268))) = int32(3)
	v4272 = int32(0)
	if v4266 < v4272 {
		v4315 = v4272
		goto L1298
	} else {
		goto L1299
	}
L1297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4268)+4)) = v4315
	v4320 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v4321 = *(*int32)(unsafe.Add(mBase, uint32(v4320)+520))
	v4323 = v4321 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4320)+520)) = v4323
	v4325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4268)+24)) = uint8(base.B2i32(v4263 != v4325))
	*(*int32)(unsafe.Add(mBase, uint32(v4268)+20)) = v4260
	*(*int32)(unsafe.Add(mBase, uint32(v4268)+16)) = v4325
	*(*int32)(unsafe.Add(mBase, uint32(v4268)+12)) = v4257
	*(*int32)(unsafe.Add(mBase, uint32(v4268)+8)) = v4323
	if v4263 == v4325 {
		v4341 = v4263
		goto L1311
	} else {
		goto L1312
	}
L1298:
	;
	goto L1297
L1299:
	;
	v4277 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4278 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+60))
	if v4278 == int32(0) {
		v4315 = v4272
		goto L1298
	} else {
		goto L1300
	}
L1300:
	;
	v4281 = v4266 + v4278
	v4282 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+188))
	if base.Ui32(v4282) <= base.Ui32(v4281) {
		goto L1302
	} else {
		goto L1303
	}
L1301:
	;
	v4292 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+196))
	if v4291 == int32(0) {
		v4315 = v4292
		goto L1298
	} else {
		goto L1305
	}
L1302:
	;
	v4284 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+192))
	v4291 = v4284
	goto L1301
L1303:
	;
	goto L1304
L1304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4277)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4277)+188)) = v4278
	v4289 = F_strchr(m, v4278, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4277)+192)) = v4289
	v4291 = v4289
	goto L1301
L1305:
	;
	if base.Ui32(v4281) <= base.Ui32(v4291) {
		v4315 = v4292
		goto L1298
	} else {
		goto L1306
	}
L1306:
	;
	v4296 = v4291
	v4298 = v4292
	goto L1307
L1307:
	;
	v4301 = int32(1)
	v4302 = v4298 + v4301
	*(*int32)(unsafe.Add(mBase, uint32(v4277)+196)) = v4302
	v4305 = v4296 + v4301
	*(*int32)(unsafe.Add(mBase, uint32(v4277)+188)) = v4305
	v4308 = F_strchr(m, v4305, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4277)+192)) = v4308
	if v4308 == int32(0) {
		v4315 = v4302
		goto L1298
	} else {
		goto L1309
	}
L1308:
	;
	v4315 = v4302
	goto L1298
L1309:
	;
	if base.Ui32(v4308) < base.Ui32(v4281) {
		v4296 = v4308
		v4298 = v4302
		goto L1307
	} else {
		goto L1310
	}
L1310:
	;
	goto L1308
L1311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4268)+28)) = v4341
	if v4257 == int32(0) {
		goto L1315
	} else {
		goto L1316
	}
L1312:
	;
	v4335 = *(*int32)(unsafe.Add(mBase, uint32(v4263)+4))
	if v4335 != int32(1) {
		v4341 = v4263
		goto L1311
	} else {
		goto L1313
	}
L1313:
	;
	v4338 = *(*int32)(unsafe.Add(mBase, uint32(v4263)+12))
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(v4338)))
	if v4339 != 0 {
		v4341 = v4263
		goto L1311
	} else {
		goto L1314
	}
L1314:
	;
	v4341 = int32(0)
	goto L1311
L1315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4268
	v10317 = v241
	goto L5
L1316:
	;
	v4346 = *(*int32)(unsafe.Add(mBase, _consts[1346]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v4346
	v4354 = F_pg_snprintf(m, v27+int32(4784), int32(32), int32(499623), v27+int32(208))
	mBase = m.M
	v4355 = m.ExcPending
	if v4355 != 0 {
		goto L66
	} else {
		goto L1317
	}
L1317:
	;
	v4356 = int32(0)
	v4359 = *(*int32)(unsafe.Add(mBase, uint32(v4268)+4))
	v4364 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v4356, v4356)
	mBase = m.M
	v4365 = m.ExcPending
	if v4365 != 0 {
		goto L66
	} else {
		goto L1318
	}
L1318:
	;
	v4367 = F_plpgsql_build_variable(m, v27+int32(4784), v4359, v4364, int32(1))
	mBase = m.M
	v4368 = m.ExcPending
	if v4368 != 0 {
		goto L66
	} else {
		goto L1319
	}
L1319:
	;
	v4369 = *(*int32)(unsafe.Add(mBase, uint32(v4367)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4268)+16)) = v4369
	if v4260 == int32(0) {
		goto L1315
	} else {
		goto L1320
	}
L1320:
	;
	v4373 = *(*int32)(unsafe.Add(mBase, uint32(v4260)+4))
	if v4373 <= int32(0) {
		goto L1315
	} else {
		goto L1321
	}
L1321:
	;
	v4380 = v4356
	goto L1322
L1322:
	;
	v4400 = *(*int32)(unsafe.Add(mBase, uint32(v4260)+12))
	v4404 = *(*int32)(unsafe.Add(mBase, uint32(v4400+v4380<<(uint(int32(2))%32))))
	v4405 = *(*int32)(unsafe.Add(mBase, uint32(v4404)+4))
	F_initStringInfo(m, v27+int32(4768))
	mBase = m.M
	v4409 = m.ExcPending
	if v4409 != 0 {
		goto L66
	} else {
		goto L1324
	}
L1323:
	;
	goto L1315
L1324:
	;
	v4410 = *(*int32)(unsafe.Add(mBase, uint32(v4405)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+196)) = v4410
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v27 + int32(4784)
	F_appendStringInfo(m, v27+int32(4768), int32(651848), v27+int32(192))
	mBase = m.M
	v4421 = m.ExcPending
	if v4421 != 0 {
		goto L66
	} else {
		goto L1325
	}
L1325:
	;
	v4422 = *(*int32)(unsafe.Add(mBase, uint32(v4405)))
	F_pfree(m, v4422)
	mBase = m.M
	v4424 = m.ExcPending
	if v4424 != 0 {
		goto L66
	} else {
		goto L1326
	}
L1326:
	;
	v4425 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378])))
	v4426 = F_pstrdup(m, v4425)
	mBase = m.M
	v4427 = m.ExcPending
	if v4427 != 0 {
		goto L66
	} else {
		goto L1327
	}
L1327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4405))) = v4426
	v4430 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	*(*int32)(unsafe.Add(mBase, uint32(v4405)+12)) = v4430
	v4432 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378])))
	F_pfree(m, v4432)
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L66
	} else {
		goto L1328
	}
L1328:
	;
	v4436 = v4380 + int32(1)
	v4437 = *(*int32)(unsafe.Add(mBase, uint32(v4260)+4))
	if v4436 < v4437 {
		v4380 = v4436
		goto L1322
	} else {
		goto L1329
	}
L1329:
	;
	goto L1323
L1330:
	;
	if v4469 != int32(384) {
		goto L1331
	} else {
		goto L1332
	}
L1331:
	;
	F_plpgsql_push_back_token(m, v4469, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4478 = m.ExcPending
	if v4478 != 0 {
		goto L66
	} else {
		goto L1334
	}
L1332:
	;
	v4494 = int32(0)
	goto L1333
L1333:
	;
	F_plpgsql_push_back_token(m, int32(384), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4501 = m.ExcPending
	if v4501 != 0 {
		goto L66
	} else {
		goto L1336
	}
L1334:
	;
	v4480 = int32(0)
	v4484 = int32(1)
	v4492 = F_read_sql_construct(m, int32(384), v4480, v4480, int32(522678), int32(2), v4484, v4484, v4480, v4480, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4493 = m.ExcPending
	if v4493 != 0 {
		goto L66
	} else {
		goto L1335
	}
L1335:
	;
	v4494 = v4492
	goto L1333
L1336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4494
	v10317 = v241
	goto L5
L1337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4507
	v10317 = v241
	goto L5
L1338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4516
	v10317 = v241
	goto L5
L1339:
	;
	v4524 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4525 = int32(0)
	if v4524 < v4525 {
		v4568 = v4525
		goto L1341
	} else {
		goto L1342
	}
L1340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4520))) = v4568
	v4574 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4520)+4)) = v4574
	v4576 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4520)+8)) = v4576
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4520
	v10317 = v241
	goto L5
L1341:
	;
	goto L1340
L1342:
	;
	v4530 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4531 = *(*int32)(unsafe.Add(mBase, uint32(v4530)+60))
	if v4531 == int32(0) {
		v4568 = v4525
		goto L1341
	} else {
		goto L1343
	}
L1343:
	;
	v4534 = v4524 + v4531
	v4535 = *(*int32)(unsafe.Add(mBase, uint32(v4530)+188))
	if base.Ui32(v4535) <= base.Ui32(v4534) {
		goto L1345
	} else {
		goto L1346
	}
L1344:
	;
	v4545 = *(*int32)(unsafe.Add(mBase, uint32(v4530)+196))
	if v4544 == int32(0) {
		v4568 = v4545
		goto L1341
	} else {
		goto L1348
	}
L1345:
	;
	v4537 = *(*int32)(unsafe.Add(mBase, uint32(v4530)+192))
	v4544 = v4537
	goto L1344
L1346:
	;
	goto L1347
L1347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4530)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4530)+188)) = v4531
	v4542 = F_strchr(m, v4531, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4530)+192)) = v4542
	v4544 = v4542
	goto L1344
L1348:
	;
	if base.Ui32(v4534) <= base.Ui32(v4544) {
		v4568 = v4545
		goto L1341
	} else {
		goto L1349
	}
L1349:
	;
	v4549 = v4544
	v4551 = v4545
	goto L1350
L1350:
	;
	v4554 = int32(1)
	v4555 = v4551 + v4554
	*(*int32)(unsafe.Add(mBase, uint32(v4530)+196)) = v4555
	v4558 = v4549 + v4554
	*(*int32)(unsafe.Add(mBase, uint32(v4530)+188)) = v4558
	v4561 = F_strchr(m, v4558, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4530)+192)) = v4561
	if v4561 == int32(0) {
		v4568 = v4555
		goto L1341
	} else {
		goto L1352
	}
L1351:
	;
	v4568 = v4555
	goto L1341
L1352:
	;
	if base.Ui32(v4561) < base.Ui32(v4534) {
		v4549 = v4561
		v4551 = v4555
		goto L1350
	} else {
		goto L1353
	}
L1353:
	;
	goto L1351
L1354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4581
	v10317 = v241
	goto L5
L1355:
	;
	goto L1356
L1356:
	;
	v4583 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v4583
	*(*int32)(unsafe.Add(mBase, uint32(v27)+280)) = v4583
	v4590 = F_list_make1_impl(m, int32(1), v27+int32(220))
	mBase = m.M
	v4591 = m.ExcPending
	if v4591 != 0 {
		goto L66
	} else {
		goto L1357
	}
L1357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4590
	v10317 = v241
	goto L5
L1358:
	;
	v4596 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v4594))) = v4596
	v4600 = *(*int32)(unsafe.Add(mBase, uint32(v142-v4596)))
	v4601 = int32(0)
	if v4600 < v4601 {
		v4644 = v4601
		goto L1360
	} else {
		goto L1361
	}
L1359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4594)+4)) = v4644
	v4649 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v4650 = *(*int32)(unsafe.Add(mBase, uint32(v4649)+520))
	v4652 = v4650 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4649)+520)) = v4652
	*(*int32)(unsafe.Add(mBase, uint32(v4594)+8)) = v4652
	v4656 = v132 - int32(32)
	v4657 = *(*int32)(unsafe.Add(mBase, uint32(v4656)))
	*(*int32)(unsafe.Add(mBase, uint32(v4594)+12)) = v4657
	v4659 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4594)+16)) = v4659
	v4661 = *(*int32)(unsafe.Add(mBase, uint32(v4656)))
	v4662 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v4663 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v4661, v4662, v4663, l1)
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		goto L66
	} else {
		goto L1373
	}
L1360:
	;
	goto L1359
L1361:
	;
	v4606 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4607 = *(*int32)(unsafe.Add(mBase, uint32(v4606)+60))
	if v4607 == int32(0) {
		v4644 = v4601
		goto L1360
	} else {
		goto L1362
	}
L1362:
	;
	v4610 = v4600 + v4607
	v4611 = *(*int32)(unsafe.Add(mBase, uint32(v4606)+188))
	if base.Ui32(v4611) <= base.Ui32(v4610) {
		goto L1364
	} else {
		goto L1365
	}
L1363:
	;
	v4621 = *(*int32)(unsafe.Add(mBase, uint32(v4606)+196))
	if v4620 == int32(0) {
		v4644 = v4621
		goto L1360
	} else {
		goto L1367
	}
L1364:
	;
	v4613 = *(*int32)(unsafe.Add(mBase, uint32(v4606)+192))
	v4620 = v4613
	goto L1363
L1365:
	;
	goto L1366
L1366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4606)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4606)+188)) = v4607
	v4618 = F_strchr(m, v4607, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4606)+192)) = v4618
	v4620 = v4618
	goto L1363
L1367:
	;
	if base.Ui32(v4610) <= base.Ui32(v4620) {
		v4644 = v4621
		goto L1360
	} else {
		goto L1368
	}
L1368:
	;
	v4625 = v4620
	v4627 = v4621
	goto L1369
L1369:
	;
	v4630 = int32(1)
	v4631 = v4627 + v4630
	*(*int32)(unsafe.Add(mBase, uint32(v4606)+196)) = v4631
	v4634 = v4625 + v4630
	*(*int32)(unsafe.Add(mBase, uint32(v4606)+188)) = v4634
	v4637 = F_strchr(m, v4634, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4606)+192)) = v4637
	if v4637 == int32(0) {
		v4644 = v4631
		goto L1360
	} else {
		goto L1371
	}
L1370:
	;
	v4644 = v4631
	goto L1360
L1371:
	;
	if base.Ui32(v4637) < base.Ui32(v4610) {
		v4625 = v4637
		v4627 = v4631
		goto L1369
	} else {
		goto L1372
	}
L1372:
	;
	goto L1370
L1373:
	;
	v4668 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	v4669 = *(*int32)(unsafe.Add(mBase, uint32(v4668)))
	if v4669 != 0 {
		goto L1375
	} else {
		goto L1376
	}
L1374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4594
	v10317 = v241
	goto L5
L1375:
	;
	v4670 = v4668
	goto L1378
L1376:
	;
	v4673 = v4668
	goto L1377
L1377:
	;
	v4675 = *(*int32)(unsafe.Add(mBase, uint32(v4673)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = v4675
	goto L1374
L1378:
	;
	v4671 = *(*int32)(unsafe.Add(mBase, uint32(v4670)+8))
	v4672 = *(*int32)(unsafe.Add(mBase, uint32(v4671)))
	if v4672 != 0 {
		v4670 = v4671
		goto L1378
	} else {
		goto L1380
	}
L1379:
	;
	v4673 = v4671
	goto L1377
L1380:
	;
	goto L1379
L1381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4679))) = int32(5)
	v4685 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v4686 = int32(0)
	if v4685 < v4686 {
		v4729 = v4686
		goto L1383
	} else {
		goto L1384
	}
L1382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+4)) = v4729
	v4734 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v4735 = *(*int32)(unsafe.Add(mBase, uint32(v4734)+520))
	v4737 = v4735 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4734)+520)) = v4737
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+8)) = v4737
	v4741 = v132 - int32(48)
	v4742 = *(*int32)(unsafe.Add(mBase, uint32(v4741)))
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+12)) = v4742
	v4746 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+16)) = v4746
	v4748 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4679)+20)) = v4748
	v4750 = *(*int32)(unsafe.Add(mBase, uint32(v4741)))
	v4751 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v4752 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v4750, v4751, v4752, l1)
	mBase = m.M
	v4754 = m.ExcPending
	if v4754 != 0 {
		goto L66
	} else {
		goto L1396
	}
L1383:
	;
	goto L1382
L1384:
	;
	v4691 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4692 = *(*int32)(unsafe.Add(mBase, uint32(v4691)+60))
	if v4692 == int32(0) {
		v4729 = v4686
		goto L1383
	} else {
		goto L1385
	}
L1385:
	;
	v4695 = v4685 + v4692
	v4696 = *(*int32)(unsafe.Add(mBase, uint32(v4691)+188))
	if base.Ui32(v4696) <= base.Ui32(v4695) {
		goto L1387
	} else {
		goto L1388
	}
L1386:
	;
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v4691)+196))
	if v4705 == int32(0) {
		v4729 = v4706
		goto L1383
	} else {
		goto L1390
	}
L1387:
	;
	v4698 = *(*int32)(unsafe.Add(mBase, uint32(v4691)+192))
	v4705 = v4698
	goto L1386
L1388:
	;
	goto L1389
L1389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4691)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4691)+188)) = v4692
	v4703 = F_strchr(m, v4692, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4691)+192)) = v4703
	v4705 = v4703
	goto L1386
L1390:
	;
	if base.Ui32(v4695) <= base.Ui32(v4705) {
		v4729 = v4706
		goto L1383
	} else {
		goto L1391
	}
L1391:
	;
	v4710 = v4705
	v4712 = v4706
	goto L1392
L1392:
	;
	v4715 = int32(1)
	v4716 = v4712 + v4715
	*(*int32)(unsafe.Add(mBase, uint32(v4691)+196)) = v4716
	v4719 = v4710 + v4715
	*(*int32)(unsafe.Add(mBase, uint32(v4691)+188)) = v4719
	v4722 = F_strchr(m, v4719, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4691)+192)) = v4722
	if v4722 == int32(0) {
		v4729 = v4716
		goto L1383
	} else {
		goto L1394
	}
L1393:
	;
	v4729 = v4716
	goto L1383
L1394:
	;
	if base.Ui32(v4722) < base.Ui32(v4695) {
		v4710 = v4722
		v4712 = v4716
		goto L1392
	} else {
		goto L1395
	}
L1395:
	;
	goto L1393
L1396:
	;
	v4757 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	v4758 = *(*int32)(unsafe.Add(mBase, uint32(v4757)))
	if v4758 != 0 {
		goto L1398
	} else {
		goto L1399
	}
L1397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4679
	v10317 = v241
	goto L5
L1398:
	;
	v4759 = v4757
	goto L1401
L1399:
	;
	v4762 = v4757
	goto L1400
L1400:
	;
	v4764 = *(*int32)(unsafe.Add(mBase, uint32(v4762)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = v4764
	goto L1397
L1401:
	;
	v4760 = *(*int32)(unsafe.Add(mBase, uint32(v4759)+8))
	v4761 = *(*int32)(unsafe.Add(mBase, uint32(v4760)))
	if v4761 != 0 {
		v4759 = v4760
		goto L1401
	} else {
		goto L1403
	}
L1402:
	;
	v4762 = v4760
	goto L1400
L1403:
	;
	goto L1402
L1404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4769)+4)) = v4817
	v4822 = v132 - int32(48)
	v4823 = *(*int32)(unsafe.Add(mBase, uint32(v4822)))
	*(*int32)(unsafe.Add(mBase, uint32(v4769)+12)) = v4823
	if v4770 == int32(6) {
		goto L1418
	} else {
		goto L1419
	}
L1405:
	;
	goto L1404
L1406:
	;
	v4779 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4780 = *(*int32)(unsafe.Add(mBase, uint32(v4779)+60))
	if v4780 == int32(0) {
		v4817 = v4774
		goto L1405
	} else {
		goto L1407
	}
L1407:
	;
	v4783 = v4773 + v4780
	v4784 = *(*int32)(unsafe.Add(mBase, uint32(v4779)+188))
	if base.Ui32(v4784) <= base.Ui32(v4783) {
		goto L1409
	} else {
		goto L1410
	}
L1408:
	;
	v4794 = *(*int32)(unsafe.Add(mBase, uint32(v4779)+196))
	if v4793 == int32(0) {
		v4817 = v4794
		goto L1405
	} else {
		goto L1412
	}
L1409:
	;
	v4786 = *(*int32)(unsafe.Add(mBase, uint32(v4779)+192))
	v4793 = v4786
	goto L1408
L1410:
	;
	goto L1411
L1411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4779)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4779)+188)) = v4780
	v4791 = F_strchr(m, v4780, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4779)+192)) = v4791
	v4793 = v4791
	goto L1408
L1412:
	;
	if base.Ui32(v4783) <= base.Ui32(v4793) {
		v4817 = v4794
		goto L1405
	} else {
		goto L1413
	}
L1413:
	;
	v4798 = v4793
	v4800 = v4794
	goto L1414
L1414:
	;
	v4803 = int32(1)
	v4804 = v4800 + v4803
	*(*int32)(unsafe.Add(mBase, uint32(v4779)+196)) = v4804
	v4807 = v4798 + v4803
	*(*int32)(unsafe.Add(mBase, uint32(v4779)+188)) = v4807
	v4810 = F_strchr(m, v4807, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v4779)+192)) = v4810
	if v4810 == int32(0) {
		v4817 = v4804
		goto L1405
	} else {
		goto L1416
	}
L1415:
	;
	v4817 = v4804
	goto L1405
L1416:
	;
	if base.Ui32(v4810) < base.Ui32(v4783) {
		v4798 = v4810
		v4800 = v4804
		goto L1414
	} else {
		goto L1417
	}
L1417:
	;
	goto L1415
L1418:
	;
	v4829 = int32(36)
	goto L1420
L1419:
	;
	v4829 = int32(20)
	goto L1420
L1420:
	;
	v4831 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v4769+v4829))) = v4831
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4769
	v4834 = *(*int32)(unsafe.Add(mBase, uint32(v4822)))
	v4835 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v4836 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v4834, v4835, v4836, l1)
	mBase = m.M
	v4838 = m.ExcPending
	if v4838 != 0 {
		goto L66
	} else {
		goto L1421
	}
L1421:
	;
	v4841 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	v4842 = *(*int32)(unsafe.Add(mBase, uint32(v4841)))
	if v4842 != 0 {
		goto L1423
	} else {
		goto L1424
	}
L1422:
	;
	v10317 = v241
	goto L5
L1423:
	;
	v4843 = v4841
	goto L1426
L1424:
	;
	v4846 = v4841
	goto L1425
L1425:
	;
	v4848 = *(*int32)(unsafe.Add(mBase, uint32(v4846)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = v4848
	goto L1422
L1426:
	;
	v4844 = *(*int32)(unsafe.Add(mBase, uint32(v4843)+8))
	v4845 = *(*int32)(unsafe.Add(mBase, uint32(v4844)))
	if v4845 != 0 {
		v4843 = v4844
		goto L1426
	} else {
		goto L1428
	}
L1427:
	;
	v4846 = v4844
	goto L1425
L1428:
	;
	goto L1427
L1429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1376]))) = v4854
	v4857 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	if v4854 != int32(277) {
		goto L1432
	} else {
		goto L1433
	}
L1430:
	;
	F_plpgsql_push_back_token(m, v4854, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5109 = m.ExcPending
	if v5109 != 0 {
		goto L66
	} else {
		goto L1480
	}
L1431:
	;
	if v4854 == int32(363) {
		goto L10
	} else {
		goto L1479
	}
L1432:
	;
	if v4854 != int32(317) {
		goto L1431
	} else {
		goto L1435
	}
L1433:
	;
	goto L1434
L1434:
	;
	v5021 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	v5022 = *(*int32)(unsafe.Add(mBase, uint32(v5021)))
	if v5022 != 0 {
		goto L1457
	} else {
		goto L1458
	}
L1435:
	;
	v4864 = int32(0)
	v4867 = int32(1)
	v4876 = F_read_sql_construct(m, int32(336), int32(381), v4864, int32(528054), int32(2), v4867, v4867, v4864, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4877 = m.ExcPending
	if v4877 != 0 {
		goto L66
	} else {
		goto L1436
	}
L1436:
	;
	v4879 = F_palloc0(m, int32(32))
	mBase = m.M
	v4880 = m.ExcPending
	if v4880 != 0 {
		goto L66
	} else {
		goto L1437
	}
L1437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4879))) = int32(18)
	v4884 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v4885 = *(*int32)(unsafe.Add(mBase, uint32(v4884)+520))
	v4887 = v4885 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4884)+520)) = v4887
	*(*int32)(unsafe.Add(mBase, uint32(v4879)+8)) = v4887
	v4891 = v132 - int32(4)
	v4892 = *(*int32)(unsafe.Add(mBase, uint32(v4891)))
	if v4892 != 0 {
		goto L1439
	} else {
		goto L1440
	}
L1438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4879)+24)) = v4876
	v4946 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v4946 == int32(381) {
		goto L1449
	} else {
		goto L1450
	}
L1439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4879)+16)) = v4892
	v4894 = *(*int32)(unsafe.Add(mBase, uint32(v4891)))
	v4897 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v4894, v4897, l1)
	mBase = m.M
	v4899 = m.ExcPending
	if v4899 != 0 {
		goto L66
	} else {
		goto L1442
	}
L1440:
	;
	goto L1441
L1441:
	;
	v4902 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(8))))
	if v4902 == int32(0) {
		goto L43
	} else {
		goto L1443
	}
L1442:
	;
	goto L1438
L1443:
	;
	v4907 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v4910 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v4913 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v4902, v4913, l1)
	mBase = m.M
	v4915 = m.ExcPending
	if v4915 != 0 {
		goto L66
	} else {
		goto L1444
	}
L1444:
	;
	v4917 = F_palloc0(m, int32(40))
	mBase = m.M
	v4918 = m.ExcPending
	if v4918 != 0 {
		goto L66
	} else {
		goto L1445
	}
L1445:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4917)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v4917)+12)) = v4907
	*(*int32)(unsafe.Add(mBase, uint32(v4917)+8)) = int32(646955)
	*(*int32)(unsafe.Add(mBase, uint32(v4917))) = int32(1)
	v4927 = F_palloc(m, int32(4))
	mBase = m.M
	v4928 = m.ExcPending
	if v4928 != 0 {
		goto L66
	} else {
		goto L1446
	}
L1446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4917)+32)) = v4927
	v4931 = F_palloc(m, int32(4))
	mBase = m.M
	v4932 = m.ExcPending
	if v4932 != 0 {
		goto L66
	} else {
		goto L1447
	}
L1447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4917)+36)) = v4931
	v4934 = *(*int32)(unsafe.Add(mBase, uint32(v4917)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4934))) = v4910
	v4936 = *(*int32)(unsafe.Add(mBase, uint32(v4902)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4931))) = v4936
	F_plpgsql_adddatum(m, v4917)
	mBase = m.M
	v4939 = m.ExcPending
	if v4939 != 0 {
		goto L66
	} else {
		goto L1448
	}
L1448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4879)+16)) = v4917
	goto L1438
L1449:
	;
	goto L1452
L1450:
	;
	goto L1451
L1451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v4879
	v10317 = v241
	goto L5
L1452:
	;
	v4975 = int32(0)
	v4978 = int32(1)
	v4987 = F_read_sql_construct(m, int32(44), int32(336), v4975, int32(518714), int32(2), v4978, v4978, v4975, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v4988 = m.ExcPending
	if v4988 != 0 {
		goto L66
	} else {
		goto L1454
	}
L1453:
	;
	goto L1451
L1454:
	;
	v4989 = *(*int32)(unsafe.Add(mBase, uint32(v4879)+28))
	v4990 = F_lappend(m, v4989, v4987)
	mBase = m.M
	v4991 = m.ExcPending
	if v4991 != 0 {
		goto L66
	} else {
		goto L1455
	}
L1455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4879)+28)) = v4990
	v4993 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v4993 == int32(44) {
		goto L1452
	} else {
		goto L1456
	}
L1456:
	;
	goto L1453
L1457:
	;
	v5071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v5071 != 0 {
		goto L1430
	} else {
		goto L1468
	}
L1458:
	;
	v5023 = *(*int32)(unsafe.Add(mBase, uint32(v5021)+24))
	v5024 = *(*int32)(unsafe.Add(mBase, uint32(v5023)+4))
	if v5024 != int32(1790) {
		goto L1457
	} else {
		goto L1459
	}
L1459:
	;
	v5028 = F_palloc0(m, int32(32))
	mBase = m.M
	v5029 = m.ExcPending
	if v5029 != 0 {
		goto L66
	} else {
		goto L1460
	}
L1460:
	;
	v5030 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v5028))) = v5030
	v5033 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5034 = *(*int32)(unsafe.Add(mBase, uint32(v5033)+520))
	v5036 = v5034 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5033)+520)) = v5036
	*(*int32)(unsafe.Add(mBase, uint32(v5028)+8)) = v5036
	v5039 = *(*int32)(unsafe.Add(mBase, uint32(v5021)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5028)+24)) = v5039
	v5043 = *(*int32)(unsafe.Add(mBase, uint32(v132-v5030)))
	if v5043 != 0 {
		goto L1461
	} else {
		goto L1462
	}
L1461:
	;
	v5046 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
	if v5046 != 0 {
		goto L42
	} else {
		goto L1464
	}
L1462:
	;
	goto L1463
L1463:
	;
	v5047 = *(*int32)(unsafe.Add(mBase, uint32(v5021)+28))
	if v5047 == int32(0) {
		goto L41
	} else {
		goto L1465
	}
L1464:
	;
	goto L1463
L1465:
	;
	v5055 = F_read_cursor_args(m, v5021, int32(336), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5056 = m.ExcPending
	if v5056 != 0 {
		goto L66
	} else {
		goto L1466
	}
L1466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5028)+28)) = v5055
	v5060 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v5063 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v5067 = F_plpgsql_build_record(m, v5060, v5063, int32(0), int32(2249), int32(1))
	mBase = m.M
	v5068 = m.ExcPending
	if v5068 != 0 {
		goto L66
	} else {
		goto L1467
	}
L1467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5028)+16)) = v5067
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5028
	v10317 = v241
	goto L5
L1468:
	;
	v5072 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v5072 == int32(0) {
		goto L1430
	} else {
		goto L1469
	}
L1469:
	;
	v5075 = int32(355250)
	v5078 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1395])))
	v5079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5072))))
	if v5079 == int32(0) {
		v5098 = v5078
		v5099 = v5079
		goto L1471
	} else {
		goto L1472
	}
L1470:
	;
	if v5099-v5098 != 0 {
		goto L1430
	} else {
		goto L1478
	}
L1471:
	;
	goto L1470
L1472:
	;
	if v5078 != v5079 {
		v5098 = v5078
		v5099 = v5079
		goto L1471
	} else {
		goto L1473
	}
L1473:
	;
	v5083 = v5072
	v5084 = v5075
	goto L1474
L1474:
	;
	v5087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5084)+1)))
	v5088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5083)+1)))
	if v5088 == int32(0) {
		v5098 = v5087
		v5099 = v5088
		goto L1471
	} else {
		goto L1476
	}
L1475:
	;
	v5098 = v5087
	v5099 = v5088
	goto L1471
L1476:
	;
	v5091 = int32(1)
	if v5087 == v5088 {
		v5083 = v5083 + v5091
		v5084 = v5084 + v5091
		goto L1474
	} else {
		goto L1477
	}
L1477:
	;
	goto L1475
L1478:
	;
	goto L10
L1479:
	;
	goto L1430
L1480:
	;
	v5110 = int32(0)
	v5126 = F_read_sql_construct(m, int32(269), int32(336), v5110, int32(518719), v5110, int32(1), v5110, v27+int32(4764), v27+int32(4752), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5127 = m.ExcPending
	if v5127 != 0 {
		goto L66
	} else {
		goto L1481
	}
L1481:
	;
	v5128 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1376])))
	if v5128 == int32(269) {
		v10005 = v5126
		v10006 = v5110
		goto L9
	} else {
		goto L1482
	}
L1482:
	;
	v5132 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1355])))
	if v5132 == int32(1) {
		goto L1483
	} else {
		goto L1484
	}
L1483:
	;
	v5135 = *(*int32)(unsafe.Add(mBase, uint32(v5126)+4))
	v5136 = *(*int32)(unsafe.Add(mBase, uint32(v5126)))
	v5137 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1396])))
	v5138 = int32(4470752)
	v5139 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v5142 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v5142
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v5137
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = int32(6372)
	v5148 = int32(4463656)
	v5149 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4784)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = v5149
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1380]))) = v27 + int32(4768)
	v5158 = F_raw_parser(m, v5136, v5135)
	mBase = m.M
	v5159 = m.ExcPending
	if v5159 != 0 {
		goto L66
	} else {
		goto L1486
	}
L1484:
	;
	goto L1485
L1485:
	;
	v5171 = F_palloc0(m, int32(28))
	mBase = m.M
	v5172 = m.ExcPending
	if v5172 != 0 {
		goto L66
	} else {
		goto L1487
	}
L1486:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v5139
	v5163 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v5163
	goto L1485
L1487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5171))) = int32(7)
	v5176 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5177 = *(*int32)(unsafe.Add(mBase, uint32(v5176)+520))
	v5179 = v5177 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5176)+520)) = v5179
	*(*int32)(unsafe.Add(mBase, uint32(v5171)+8)) = v5179
	v5183 = v132 - int32(4)
	v5184 = *(*int32)(unsafe.Add(mBase, uint32(v5183)))
	if v5184 != 0 {
		goto L1489
	} else {
		goto L1490
	}
L1488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5171)+24)) = v5126
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5171
	v10317 = v241
	goto L5
L1489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5171)+16)) = v5184
	v5186 = *(*int32)(unsafe.Add(mBase, uint32(v5183)))
	v5189 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v5186, v5189, l1)
	mBase = m.M
	v5191 = m.ExcPending
	if v5191 != 0 {
		goto L66
	} else {
		goto L1492
	}
L1490:
	;
	goto L1491
L1491:
	;
	v5194 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(8))))
	if v5194 == int32(0) {
		goto L40
	} else {
		goto L1493
	}
L1492:
	;
	goto L1488
L1493:
	;
	v5199 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v5202 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v5205 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	F_check_assignable(m, v5194, v5205, l1)
	mBase = m.M
	v5207 = m.ExcPending
	if v5207 != 0 {
		goto L66
	} else {
		goto L1494
	}
L1494:
	;
	v5209 = F_palloc0(m, int32(40))
	mBase = m.M
	v5210 = m.ExcPending
	if v5210 != 0 {
		goto L66
	} else {
		goto L1495
	}
L1495:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5209)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5209)+12)) = v5199
	*(*int32)(unsafe.Add(mBase, uint32(v5209)+8)) = int32(646955)
	*(*int32)(unsafe.Add(mBase, uint32(v5209))) = int32(1)
	v5219 = F_palloc(m, int32(4))
	mBase = m.M
	v5220 = m.ExcPending
	if v5220 != 0 {
		goto L66
	} else {
		goto L1496
	}
L1496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5209)+32)) = v5219
	v5223 = F_palloc(m, int32(4))
	mBase = m.M
	v5224 = m.ExcPending
	if v5224 != 0 {
		goto L66
	} else {
		goto L1497
	}
L1497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5209)+36)) = v5223
	v5226 = *(*int32)(unsafe.Add(mBase, uint32(v5209)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5226))) = v5202
	v5228 = *(*int32)(unsafe.Add(mBase, uint32(v5194)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5223))) = v5228
	F_plpgsql_adddatum(m, v5209)
	mBase = m.M
	v5231 = m.ExcPending
	if v5231 != 0 {
		goto L66
	} else {
		goto L1498
	}
L1498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5171)+16)) = v5209
	goto L1488
L1499:
	;
	v5243 = v5239
	goto L1501
L1500:
	;
	v5240 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v5241 = F_NameListToString(m, v5240)
	mBase = m.M
	v5242 = m.ExcPending
	if v5242 != 0 {
		goto L66
	} else {
		goto L1502
	}
L1501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5243
	v5245 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5246 = int32(0)
	if v5245 < v5246 {
		v5289 = v5246
		goto L1504
	} else {
		goto L1505
	}
L1502:
	;
	v5243 = v5241
	goto L1501
L1503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v5289
	v5293 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v5294 = *(*int32)(unsafe.Add(mBase, uint32(v5293)))
	v5295 = int32(1)
	if base.Ui32(v5294-v5295) <= base.Ui32(v5295) {
		goto L1517
	} else {
		goto L1518
	}
L1504:
	;
	goto L1503
L1505:
	;
	v5251 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5252 = *(*int32)(unsafe.Add(mBase, uint32(v5251)+60))
	if v5252 == int32(0) {
		v5289 = v5246
		goto L1504
	} else {
		goto L1506
	}
L1506:
	;
	v5255 = v5245 + v5252
	v5256 = *(*int32)(unsafe.Add(mBase, uint32(v5251)+188))
	if base.Ui32(v5256) <= base.Ui32(v5255) {
		goto L1508
	} else {
		goto L1509
	}
L1507:
	;
	v5266 = *(*int32)(unsafe.Add(mBase, uint32(v5251)+196))
	if v5265 == int32(0) {
		v5289 = v5266
		goto L1504
	} else {
		goto L1511
	}
L1508:
	;
	v5258 = *(*int32)(unsafe.Add(mBase, uint32(v5251)+192))
	v5265 = v5258
	goto L1507
L1509:
	;
	goto L1510
L1510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5251)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5251)+188)) = v5252
	v5263 = F_strchr(m, v5252, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5251)+192)) = v5263
	v5265 = v5263
	goto L1507
L1511:
	;
	if base.Ui32(v5255) <= base.Ui32(v5265) {
		v5289 = v5266
		goto L1504
	} else {
		goto L1512
	}
L1512:
	;
	v5270 = v5265
	v5272 = v5266
	goto L1513
L1513:
	;
	v5275 = int32(1)
	v5276 = v5272 + v5275
	*(*int32)(unsafe.Add(mBase, uint32(v5251)+196)) = v5276
	v5279 = v5270 + v5275
	*(*int32)(unsafe.Add(mBase, uint32(v5251)+188)) = v5279
	v5282 = F_strchr(m, v5279, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5251)+192)) = v5282
	if v5282 == int32(0) {
		v5289 = v5276
		goto L1504
	} else {
		goto L1515
	}
L1514:
	;
	v5289 = v5276
	goto L1504
L1515:
	;
	if base.Ui32(v5282) < base.Ui32(v5255) {
		v5270 = v5282
		v5272 = v5276
		goto L1513
	} else {
		goto L1516
	}
L1516:
	;
	goto L1514
L1517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = int32(0)
	v5301 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v5301
	v10317 = v241
	goto L5
L1518:
	;
	goto L1519
L1519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v5293
	v5310 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5311 = m.ExcPending
	if v5311 != 0 {
		goto L66
	} else {
		goto L1520
	}
L1520:
	;
	F_plpgsql_push_back_token(m, v5310, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5317 = m.ExcPending
	if v5317 != 0 {
		goto L66
	} else {
		goto L1521
	}
L1521:
	;
	if v5310 != int32(44) {
		v10317 = v241
		goto L5
	} else {
		goto L1522
	}
L1522:
	;
	v5320 = *(*int32)(unsafe.Add(mBase, uint32(v27)+304))
	v5321 = *(*int32)(unsafe.Add(mBase, uint32(v27)+312))
	v5322 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5327 = F_read_into_scalar_list(m, v5320, v5321, v5322, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5328 = m.ExcPending
	if v5328 != 0 {
		goto L66
	} else {
		goto L1523
	}
L1523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v5327
	v10317 = v241
	goto L5
L1524:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+312)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v5376
	v5386 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5387 = m.ExcPending
	if v5387 != 0 {
		goto L66
	} else {
		goto L1538
	}
L1525:
	;
	goto L1524
L1526:
	;
	v5338 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5339 = *(*int32)(unsafe.Add(mBase, uint32(v5338)+60))
	if v5339 == int32(0) {
		v5376 = v5333
		goto L1525
	} else {
		goto L1527
	}
L1527:
	;
	v5342 = v5332 + v5339
	v5343 = *(*int32)(unsafe.Add(mBase, uint32(v5338)+188))
	if base.Ui32(v5343) <= base.Ui32(v5342) {
		goto L1529
	} else {
		goto L1530
	}
L1528:
	;
	v5353 = *(*int32)(unsafe.Add(mBase, uint32(v5338)+196))
	if v5352 == int32(0) {
		v5376 = v5353
		goto L1525
	} else {
		goto L1532
	}
L1529:
	;
	v5345 = *(*int32)(unsafe.Add(mBase, uint32(v5338)+192))
	v5352 = v5345
	goto L1528
L1530:
	;
	goto L1531
L1531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5338)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5338)+188)) = v5339
	v5350 = F_strchr(m, v5339, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5338)+192)) = v5350
	v5352 = v5350
	goto L1528
L1532:
	;
	if base.Ui32(v5342) <= base.Ui32(v5352) {
		v5376 = v5353
		goto L1525
	} else {
		goto L1533
	}
L1533:
	;
	v5357 = v5352
	v5359 = v5353
	goto L1534
L1534:
	;
	v5362 = int32(1)
	v5363 = v5359 + v5362
	*(*int32)(unsafe.Add(mBase, uint32(v5338)+196)) = v5363
	v5366 = v5357 + v5362
	*(*int32)(unsafe.Add(mBase, uint32(v5338)+188)) = v5366
	v5369 = F_strchr(m, v5366, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5338)+192)) = v5369
	if v5369 == int32(0) {
		v5376 = v5363
		goto L1525
	} else {
		goto L1536
	}
L1535:
	;
	v5376 = v5363
	goto L1525
L1536:
	;
	if base.Ui32(v5369) < base.Ui32(v5342) {
		v5357 = v5369
		v5359 = v5363
		goto L1534
	} else {
		goto L1537
	}
L1537:
	;
	goto L1535
L1538:
	;
	F_plpgsql_push_back_token(m, v5386, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5393 = m.ExcPending
	if v5393 != 0 {
		goto L66
	} else {
		goto L1539
	}
L1539:
	;
	if v5386 != int32(44) {
		v10317 = v241
		goto L5
	} else {
		goto L1540
	}
L1540:
	;
	v5396 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	F_word_is_not_variable(m, v132, v5396, l1)
	mBase = m.M
	v5398 = m.ExcPending
	if v5398 != 0 {
		goto L66
	} else {
		goto L1541
	}
L1541:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1542:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5403))) = int32(9)
	v5409 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(24))))
	v5410 = int32(0)
	if v5409 < v5410 {
		v5453 = v5410
		goto L1545
	} else {
		goto L1546
	}
L1544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+4)) = v5453
	v5458 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5459 = *(*int32)(unsafe.Add(mBase, uint32(v5458)+520))
	v5461 = v5459 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5458)+520)) = v5461
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+8)) = v5461
	v5465 = v132 - int32(112)
	v5466 = *(*int32)(unsafe.Add(mBase, uint32(v5465)))
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+12)) = v5466
	v5470 = *(*int32)(unsafe.Add(mBase, uint32(v132+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+20)) = v5470
	v5474 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+24)) = v5474
	v5476 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+28)) = v5476
	v5479 = v132 - int32(68)
	v5480 = *(*int32)(unsafe.Add(mBase, uint32(v5479)))
	if v5480 == int32(0) {
		goto L1558
	} else {
		goto L1559
	}
L1545:
	;
	goto L1544
L1546:
	;
	v5415 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5416 = *(*int32)(unsafe.Add(mBase, uint32(v5415)+60))
	if v5416 == int32(0) {
		v5453 = v5410
		goto L1545
	} else {
		goto L1547
	}
L1547:
	;
	v5419 = v5409 + v5416
	v5420 = *(*int32)(unsafe.Add(mBase, uint32(v5415)+188))
	if base.Ui32(v5420) <= base.Ui32(v5419) {
		goto L1549
	} else {
		goto L1550
	}
L1548:
	;
	v5430 = *(*int32)(unsafe.Add(mBase, uint32(v5415)+196))
	if v5429 == int32(0) {
		v5453 = v5430
		goto L1545
	} else {
		goto L1552
	}
L1549:
	;
	v5422 = *(*int32)(unsafe.Add(mBase, uint32(v5415)+192))
	v5429 = v5422
	goto L1548
L1550:
	;
	goto L1551
L1551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5415)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5415)+188)) = v5416
	v5427 = F_strchr(m, v5416, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5415)+192)) = v5427
	v5429 = v5427
	goto L1548
L1552:
	;
	if base.Ui32(v5419) <= base.Ui32(v5429) {
		v5453 = v5430
		goto L1545
	} else {
		goto L1553
	}
L1553:
	;
	v5434 = v5429
	v5436 = v5430
	goto L1554
L1554:
	;
	v5439 = int32(1)
	v5440 = v5436 + v5439
	*(*int32)(unsafe.Add(mBase, uint32(v5415)+196)) = v5440
	v5443 = v5434 + v5439
	*(*int32)(unsafe.Add(mBase, uint32(v5415)+188)) = v5443
	v5446 = F_strchr(m, v5443, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5415)+192)) = v5446
	if v5446 == int32(0) {
		v5453 = v5440
		goto L1545
	} else {
		goto L1556
	}
L1555:
	;
	v5453 = v5440
	goto L1545
L1556:
	;
	if base.Ui32(v5446) < base.Ui32(v5419) {
		v5434 = v5446
		v5436 = v5440
		goto L1554
	} else {
		goto L1557
	}
L1557:
	;
	goto L1555
L1558:
	;
	v5484 = v132 - int32(72)
	v5485 = *(*int32)(unsafe.Add(mBase, uint32(v5484)))
	if v5485 == int32(0) {
		goto L39
	} else {
		goto L1561
	}
L1559:
	;
	v5488 = v5479
	v5489 = v5480
	goto L1560
L1560:
	;
	v5490 = *(*int32)(unsafe.Add(mBase, uint32(v5489)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5403)+16)) = v5490
	v5492 = *(*int32)(unsafe.Add(mBase, uint32(v5488)))
	v5495 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(20))))
	F_check_assignable(m, v5492, v5495, l1)
	mBase = m.M
	v5497 = m.ExcPending
	if v5497 != 0 {
		goto L66
	} else {
		goto L1562
	}
L1561:
	;
	v5488 = v5484
	v5489 = v5485
	goto L1560
L1562:
	;
	v5498 = *(*int32)(unsafe.Add(mBase, uint32(v5465)))
	v5499 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v5500 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	F_check_labels(m, v5498, v5499, v5500, l1)
	mBase = m.M
	v5502 = m.ExcPending
	if v5502 != 0 {
		goto L66
	} else {
		goto L1563
	}
L1563:
	;
	v5505 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	v5506 = *(*int32)(unsafe.Add(mBase, uint32(v5505)))
	if v5506 != 0 {
		goto L1565
	} else {
		goto L1566
	}
L1564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5403
	v10317 = v241
	goto L5
L1565:
	;
	v5507 = v5505
	goto L1568
L1566:
	;
	v5510 = v5505
	goto L1567
L1567:
	;
	v5512 = *(*int32)(unsafe.Add(mBase, uint32(v5510)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1349])) = v5512
	goto L1564
L1568:
	;
	v5508 = *(*int32)(unsafe.Add(mBase, uint32(v5507)+8))
	v5509 = *(*int32)(unsafe.Add(mBase, uint32(v5508)))
	if v5509 != 0 {
		v5507 = v5508
		goto L1568
	} else {
		goto L1570
	}
L1569:
	;
	v5510 = v5508
	goto L1567
L1570:
	;
	goto L1569
L1571:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5520))) = int32(10)
	v5525 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5526 = *(*int32)(unsafe.Add(mBase, uint32(v5525)+520))
	v5528 = v5526 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5525)+520)) = v5528
	*(*int32)(unsafe.Add(mBase, uint32(v5520)+8)) = v5528
	v5533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132-int32(32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5520)+12)) = uint8(v5533)
	v5536 = v142 - int32(8)
	v5537 = *(*int32)(unsafe.Add(mBase, uint32(v5536)))
	v5538 = int32(0)
	if v5537 < v5538 {
		v5581 = v5538
		goto L1573
	} else {
		goto L1574
	}
L1572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5520)+4)) = v5581
	v5586 = v132 - int32(16)
	v5587 = *(*int32)(unsafe.Add(mBase, uint32(v5586)))
	*(*int32)(unsafe.Add(mBase, uint32(v5520)+16)) = v5587
	v5589 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v5520)+20)) = v5589
	v5592 = *(*int32)(unsafe.Add(mBase, _consts[1349]))
	v5593 = *(*int32)(unsafe.Add(mBase, uint32(v5586)))
	if v5593 != 0 {
		goto L1587
	} else {
		goto L1588
	}
L1573:
	;
	goto L1572
L1574:
	;
	v5543 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5544 = *(*int32)(unsafe.Add(mBase, uint32(v5543)+60))
	if v5544 == int32(0) {
		v5581 = v5538
		goto L1573
	} else {
		goto L1575
	}
L1575:
	;
	v5547 = v5537 + v5544
	v5548 = *(*int32)(unsafe.Add(mBase, uint32(v5543)+188))
	if base.Ui32(v5548) <= base.Ui32(v5547) {
		goto L1577
	} else {
		goto L1578
	}
L1576:
	;
	v5558 = *(*int32)(unsafe.Add(mBase, uint32(v5543)+196))
	if v5557 == int32(0) {
		v5581 = v5558
		goto L1573
	} else {
		goto L1580
	}
L1577:
	;
	v5550 = *(*int32)(unsafe.Add(mBase, uint32(v5543)+192))
	v5557 = v5550
	goto L1576
L1578:
	;
	goto L1579
L1579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5543)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5543)+188)) = v5544
	v5555 = F_strchr(m, v5544, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5543)+192)) = v5555
	v5557 = v5555
	goto L1576
L1580:
	;
	if base.Ui32(v5547) <= base.Ui32(v5557) {
		v5581 = v5558
		goto L1573
	} else {
		goto L1581
	}
L1581:
	;
	v5562 = v5557
	v5564 = v5558
	goto L1582
L1582:
	;
	v5567 = int32(1)
	v5568 = v5564 + v5567
	*(*int32)(unsafe.Add(mBase, uint32(v5543)+196)) = v5568
	v5571 = v5562 + v5567
	*(*int32)(unsafe.Add(mBase, uint32(v5543)+188)) = v5571
	v5574 = F_strchr(m, v5571, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5543)+192)) = v5574
	if v5574 == int32(0) {
		v5581 = v5568
		goto L1573
	} else {
		goto L1584
	}
L1583:
	;
	v5581 = v5568
	goto L1573
L1584:
	;
	if base.Ui32(v5574) < base.Ui32(v5547) {
		v5562 = v5574
		v5564 = v5568
		goto L1582
	} else {
		goto L1585
	}
L1585:
	;
	goto L1583
L1586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5520
	v10317 = v241
	goto L5
L1587:
	;
	v5594 = *(*int32)(unsafe.Add(mBase, uint32(v5586)))
	if v5592 != 0 {
		goto L1591
	} else {
		goto L1592
	}
L1588:
	;
	goto L1589
L1589:
	;
	if v5592 != 0 {
		goto L1609
	} else {
		goto L1610
	}
L1590:
	;
	if v5607 == int32(0) {
		goto L38
	} else {
		goto L1600
	}
L1591:
	;
	v5595 = v5592
	goto L1594
L1592:
	;
	goto L1593
L1593:
	;
	v5607 = int32(0)
	goto L1590
L1594:
	;
	v5597 = *(*int32)(unsafe.Add(mBase, uint32(v5595)))
	if v5597 != 0 {
		goto L1596
	} else {
		goto L1597
	}
L1595:
	;
	goto L1593
L1596:
	;
	v5601 = *(*int32)(unsafe.Add(mBase, uint32(v5595)+8))
	if v5601 != 0 {
		v5595 = v5601
		goto L1594
	} else {
		goto L1599
	}
L1597:
	;
	v5600 = F_strcmp(m, v5595+int32(12), v5594)
	mBase = m.M
	if v5600 != 0 {
		goto L1596
	} else {
		goto L1598
	}
L1598:
	;
	v5607 = v5595
	goto L1590
L1599:
	;
	goto L1595
L1600:
	;
	v5610 = *(*int32)(unsafe.Add(mBase, uint32(v5607)+4))
	if v5610 == int32(1) {
		goto L1586
	} else {
		goto L1601
	}
L1601:
	;
	v5613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5520)+12)))
	if v5613 != 0 {
		goto L1586
	} else {
		goto L1602
	}
L1602:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v5617 = m.ExcPending
	if v5617 != 0 {
		goto L66
	} else {
		goto L1603
	}
L1603:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v5620 = m.ExcPending
	if v5620 != 0 {
		goto L66
	} else {
		goto L1604
	}
L1604:
	;
	v5621 = *(*int32)(unsafe.Add(mBase, uint32(v5586)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+240)) = v5621
	F_errmsg(m, int32(529319), v27+int32(240))
	mBase = m.M
	v5627 = m.ExcPending
	if v5627 != 0 {
		goto L66
	} else {
		goto L1605
	}
L1605:
	;
	v5630 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v5631 = F_plpgsql_scanner_errposition(m, v5630, l1)
	mBase = m.M
	v5632 = m.ExcPending
	if v5632 != 0 {
		goto L66
	} else {
		goto L1606
	}
L1606:
	;
	F_errfinish(m, int32(26734), int32(1752), int32(355297))
	mBase = m.M
	v5637 = m.ExcPending
	if v5637 != 0 {
		goto L66
	} else {
		goto L1607
	}
L1607:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1608:
	;
	if v5647 == int32(0) {
		goto L37
	} else {
		goto L1618
	}
L1609:
	;
	v5638 = v5592
	goto L1612
L1610:
	;
	goto L1611
L1611:
	;
	v5647 = int32(0)
	goto L1608
L1612:
	;
	v5639 = *(*int32)(unsafe.Add(mBase, uint32(v5638)))
	if v5639 != 0 {
		goto L1614
	} else {
		goto L1615
	}
L1613:
	;
	goto L1611
L1614:
	;
	v5643 = *(*int32)(unsafe.Add(mBase, uint32(v5638)+8))
	if v5643 != 0 {
		v5638 = v5643
		goto L1612
	} else {
		goto L1617
	}
L1615:
	;
	v5640 = *(*int32)(unsafe.Add(mBase, uint32(v5638)+4))
	if v5640 != int32(1) {
		goto L1614
	} else {
		goto L1616
	}
L1616:
	;
	v5647 = v5638
	goto L1608
L1617:
	;
	goto L1613
L1618:
	;
	goto L1586
L1619:
	;
	F_plpgsql_push_back_token(m, v5660, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6053 = m.ExcPending
	if v6053 != 0 {
		goto L66
	} else {
		goto L1711
	}
L1620:
	;
	v5859 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5861 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5861)+63)))
	if v5862 == int32(0) {
		goto L34
	} else {
		goto L1680
	}
L1621:
	;
	v5832 = int32(17006)
	v5835 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1397])))
	v5836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5673))))
	if v5836 == int32(0) {
		v5855 = v5835
		v5856 = v5836
		goto L1672
	} else {
		goto L1673
	}
L1622:
	;
	v5703 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v5705 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5705)+63)))
	if v5706 == int32(0) {
		goto L36
	} else {
		goto L1641
	}
L1623:
	;
	if v5660 != int32(277) {
		goto L1625
	} else {
		goto L1626
	}
L1624:
	;
	switch v5660 - int32(341) {
	case 0:
		goto L1622
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
		goto L1619
	case 17:
		goto L1620
	default:
		goto L1623
	}
L1625:
	;
	if v5660 != 0 {
		goto L1619
	} else {
		goto L1628
	}
L1626:
	;
	goto L1627
L1627:
	;
	v5672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v5672 != 0 {
		goto L1619
	} else {
		goto L1630
	}
L1628:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(246872))
	mBase = m.M
	v5671 = m.ExcPending
	if v5671 != 0 {
		goto L66
	} else {
		goto L1629
	}
L1629:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1630:
	;
	v5673 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v5673 == int32(0) {
		goto L1619
	} else {
		goto L1631
	}
L1631:
	;
	v5676 = int32(62940)
	v5679 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1398])))
	v5680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5673))))
	if v5680 == int32(0) {
		v5699 = v5679
		v5700 = v5680
		goto L1633
	} else {
		goto L1634
	}
L1632:
	;
	if v5700-v5699 != 0 {
		goto L1621
	} else {
		goto L1640
	}
L1633:
	;
	goto L1632
L1634:
	;
	if v5679 != v5680 {
		v5699 = v5679
		v5700 = v5680
		goto L1633
	} else {
		goto L1635
	}
L1635:
	;
	v5684 = v5673
	v5685 = v5676
	goto L1636
L1636:
	;
	v5688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5685)+1)))
	v5689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5684)+1)))
	if v5689 == int32(0) {
		v5699 = v5688
		v5700 = v5689
		goto L1633
	} else {
		goto L1638
	}
L1637:
	;
	v5699 = v5688
	v5700 = v5689
	goto L1633
L1638:
	;
	v5692 = int32(1)
	if v5688 == v5689 {
		v5684 = v5684 + v5692
		v5685 = v5685 + v5692
		goto L1636
	} else {
		goto L1639
	}
L1639:
	;
	goto L1637
L1640:
	;
	goto L1622
L1641:
	;
	v5710 = F_palloc0(m, int32(20))
	mBase = m.M
	v5711 = m.ExcPending
	if v5711 != 0 {
		goto L66
	} else {
		goto L1642
	}
L1642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5710))) = int32(12)
	v5714 = int32(0)
	if v5703 < v5714 {
		v5757 = v5714
		goto L1644
	} else {
		goto L1645
	}
L1643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5710)+4)) = v5757
	v5762 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5763 = *(*int32)(unsafe.Add(mBase, uint32(v5762)+520))
	v5765 = v5763 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5762)+520)) = v5765
	*(*int64)(unsafe.Add(mBase, uint32(v5710)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v5710)+8)) = v5765
	v5770 = *(*int32)(unsafe.Add(mBase, uint32(v5762)+472))
	v5775 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5776 = m.ExcPending
	if v5776 != 0 {
		goto L66
	} else {
		goto L1657
	}
L1644:
	;
	goto L1643
L1645:
	;
	v5719 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5720 = *(*int32)(unsafe.Add(mBase, uint32(v5719)+60))
	if v5720 == int32(0) {
		v5757 = v5714
		goto L1644
	} else {
		goto L1646
	}
L1646:
	;
	v5723 = v5703 + v5720
	v5724 = *(*int32)(unsafe.Add(mBase, uint32(v5719)+188))
	if base.Ui32(v5724) <= base.Ui32(v5723) {
		goto L1648
	} else {
		goto L1649
	}
L1647:
	;
	v5734 = *(*int32)(unsafe.Add(mBase, uint32(v5719)+196))
	if v5733 == int32(0) {
		v5757 = v5734
		goto L1644
	} else {
		goto L1651
	}
L1648:
	;
	v5726 = *(*int32)(unsafe.Add(mBase, uint32(v5719)+192))
	v5733 = v5726
	goto L1647
L1649:
	;
	goto L1650
L1650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5719)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5719)+188)) = v5720
	v5731 = F_strchr(m, v5720, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5719)+192)) = v5731
	v5733 = v5731
	goto L1647
L1651:
	;
	if base.Ui32(v5723) <= base.Ui32(v5733) {
		v5757 = v5734
		goto L1644
	} else {
		goto L1652
	}
L1652:
	;
	v5738 = v5733
	v5740 = v5734
	goto L1653
L1653:
	;
	v5743 = int32(1)
	v5744 = v5740 + v5743
	*(*int32)(unsafe.Add(mBase, uint32(v5719)+196)) = v5744
	v5747 = v5738 + v5743
	*(*int32)(unsafe.Add(mBase, uint32(v5719)+188)) = v5747
	v5750 = F_strchr(m, v5747, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5719)+192)) = v5750
	if v5750 == int32(0) {
		v5757 = v5744
		goto L1644
	} else {
		goto L1655
	}
L1654:
	;
	v5757 = v5744
	goto L1644
L1655:
	;
	if base.Ui32(v5750) < base.Ui32(v5723) {
		v5738 = v5750
		v5740 = v5744
		goto L1653
	} else {
		goto L1656
	}
L1656:
	;
	goto L1654
L1657:
	;
	if int32(0) <= v5770 {
		goto L1658
	} else {
		goto L1659
	}
L1658:
	;
	if v5775 != int32(59) {
		goto L35
	} else {
		goto L1661
	}
L1659:
	;
	goto L1660
L1660:
	;
	if v5775 != int32(277) {
		goto L1662
	} else {
		goto L1663
	}
L1661:
	;
	v5782 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5783 = *(*int32)(unsafe.Add(mBase, uint32(v5782)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v5710)+16)) = v5783
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5710
	v10317 = v241
	goto L5
L1662:
	;
	F_plpgsql_push_back_token(m, v5775, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5814 = m.ExcPending
	if v5814 != 0 {
		goto L66
	} else {
		goto L1669
	}
L1663:
	;
	v5788 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v5789 = m.ExcPending
	if v5789 != 0 {
		goto L66
	} else {
		goto L1664
	}
L1664:
	;
	if v5788 != int32(59) {
		goto L1662
	} else {
		goto L1665
	}
L1665:
	;
	v5792 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	v5793 = *(*int32)(unsafe.Add(mBase, uint32(v5792)))
	if base.Ui32(int32(4)) < base.Ui32(v5793) {
		goto L1662
	} else {
		goto L1666
	}
L1666:
	;
	if v5793 == int32(3) {
		goto L1662
	} else {
		goto L1667
	}
L1667:
	;
	v5798 = *(*int32)(unsafe.Add(mBase, uint32(v5792)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5710)+16)) = v5798
	v5804 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5805 = m.ExcPending
	if v5805 != 0 {
		goto L66
	} else {
		goto L1668
	}
L1668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5710
	v10317 = v241
	goto L5
L1669:
	;
	v5816 = int32(0)
	v5820 = int32(1)
	v5828 = F_read_sql_construct(m, int32(59), v5816, v5816, int32(538734), int32(2), v5820, v5820, v5816, v5816, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5829 = m.ExcPending
	if v5829 != 0 {
		goto L66
	} else {
		goto L1670
	}
L1670:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5710)+12)) = v5828
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5710
	v10317 = v241
	goto L5
L1671:
	;
	if v5856-v5855 != 0 {
		goto L1619
	} else {
		goto L1679
	}
L1672:
	;
	goto L1671
L1673:
	;
	if v5835 != v5836 {
		v5855 = v5835
		v5856 = v5836
		goto L1672
	} else {
		goto L1674
	}
L1674:
	;
	v5840 = v5673
	v5841 = v5832
	goto L1675
L1675:
	;
	v5844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5841)+1)))
	v5845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5840)+1)))
	if v5845 == int32(0) {
		v5855 = v5844
		v5856 = v5845
		goto L1672
	} else {
		goto L1677
	}
L1676:
	;
	v5855 = v5844
	v5856 = v5845
	goto L1672
L1677:
	;
	v5848 = int32(1)
	if v5844 == v5845 {
		v5840 = v5840 + v5848
		v5841 = v5841 + v5848
		goto L1675
	} else {
		goto L1678
	}
L1678:
	;
	goto L1676
L1679:
	;
	goto L1620
L1680:
	;
	v5866 = F_palloc0(m, int32(24))
	mBase = m.M
	v5867 = m.ExcPending
	if v5867 != 0 {
		goto L66
	} else {
		goto L1681
	}
L1681:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866))) = int32(13)
	v5870 = int32(0)
	if v5859 < v5870 {
		v5913 = v5870
		goto L1683
	} else {
		goto L1684
	}
L1682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+4)) = v5913
	v5918 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v5919 = *(*int32)(unsafe.Add(mBase, uint32(v5918)+520))
	v5921 = v5919 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5918)+520)) = v5921
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+8)) = v5921
	v5928 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5929 = m.ExcPending
	if v5929 != 0 {
		goto L66
	} else {
		goto L1696
	}
L1683:
	;
	goto L1682
L1684:
	;
	v5875 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5876 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+60))
	if v5876 == int32(0) {
		v5913 = v5870
		goto L1683
	} else {
		goto L1685
	}
L1685:
	;
	v5879 = v5859 + v5876
	v5880 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+188))
	if base.Ui32(v5880) <= base.Ui32(v5879) {
		goto L1687
	} else {
		goto L1688
	}
L1686:
	;
	v5890 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+196))
	if v5889 == int32(0) {
		v5913 = v5890
		goto L1683
	} else {
		goto L1690
	}
L1687:
	;
	v5882 = *(*int32)(unsafe.Add(mBase, uint32(v5875)+192))
	v5889 = v5882
	goto L1686
L1688:
	;
	goto L1689
L1689:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+188)) = v5876
	v5887 = F_strchr(m, v5876, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+192)) = v5887
	v5889 = v5887
	goto L1686
L1690:
	;
	if base.Ui32(v5879) <= base.Ui32(v5889) {
		v5913 = v5890
		goto L1683
	} else {
		goto L1691
	}
L1691:
	;
	v5894 = v5889
	v5896 = v5890
	goto L1692
L1692:
	;
	v5899 = int32(1)
	v5900 = v5896 + v5899
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+196)) = v5900
	v5903 = v5894 + v5899
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+188)) = v5903
	v5906 = F_strchr(m, v5903, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v5875)+192)) = v5906
	if v5906 == int32(0) {
		v5913 = v5900
		goto L1683
	} else {
		goto L1694
	}
L1693:
	;
	v5913 = v5900
	goto L1683
L1694:
	;
	if base.Ui32(v5906) < base.Ui32(v5879) {
		v5894 = v5906
		v5896 = v5900
		goto L1692
	} else {
		goto L1695
	}
L1695:
	;
	goto L1693
L1696:
	;
	if v5928 != int32(317) {
		goto L1697
	} else {
		goto L1698
	}
L1697:
	;
	F_plpgsql_push_back_token(m, v5928, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5937 = m.ExcPending
	if v5937 != 0 {
		goto L66
	} else {
		goto L1700
	}
L1698:
	;
	goto L1699
L1699:
	;
	v5957 = int32(0)
	v5960 = int32(1)
	v5969 = F_read_sql_construct(m, int32(59), int32(381), v5957, int32(528073), int32(2), v5960, v5960, v5957, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5970 = m.ExcPending
	if v5970 != 0 {
		goto L66
	} else {
		goto L1702
	}
L1700:
	;
	v5939 = int32(0)
	v5951 = F_read_sql_construct(m, int32(59), v5939, v5939, int32(538734), v5939, v5939, int32(1), v5939, v5939, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v5952 = m.ExcPending
	if v5952 != 0 {
		goto L66
	} else {
		goto L1701
	}
L1701:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+12)) = v5951
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5866
	v10317 = v241
	goto L5
L1702:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+16)) = v5969
	v5972 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v5972 == int32(381) {
		goto L1703
	} else {
		goto L1704
	}
L1703:
	;
	goto L1706
L1704:
	;
	goto L1705
L1705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v5866
	v10317 = v241
	goto L5
L1706:
	;
	v6001 = int32(0)
	v6004 = int32(1)
	v6013 = F_read_sql_construct(m, int32(44), int32(59), v6001, int32(538729), int32(2), v6004, v6004, v6001, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6014 = m.ExcPending
	if v6014 != 0 {
		goto L66
	} else {
		goto L1708
	}
L1707:
	;
	goto L1705
L1708:
	;
	v6015 = *(*int32)(unsafe.Add(mBase, uint32(v5866)+20))
	v6016 = F_lappend(m, v6015, v6013)
	mBase = m.M
	v6017 = m.ExcPending
	if v6017 != 0 {
		goto L66
	} else {
		goto L1709
	}
L1709:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5866)+20)) = v6016
	v6019 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v6019 == int32(44) {
		goto L1706
	} else {
		goto L1710
	}
L1710:
	;
	goto L1707
L1711:
	;
	v6054 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v6056 = F_palloc0(m, int32(20))
	mBase = m.M
	v6057 = m.ExcPending
	if v6057 != 0 {
		goto L66
	} else {
		goto L1712
	}
L1712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6056))) = int32(11)
	v6060 = int32(0)
	if v6054 < v6060 {
		v6103 = v6060
		goto L1714
	} else {
		goto L1715
	}
L1713:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6056)+4)) = v6103
	v6108 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v6109 = *(*int32)(unsafe.Add(mBase, uint32(v6108)+520))
	v6110 = int32(1)
	v6111 = v6109 + v6110
	*(*int32)(unsafe.Add(mBase, uint32(v6108)+520)) = v6111
	*(*int64)(unsafe.Add(mBase, uint32(v6056)+12)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v6056)+8)) = v6111
	v6116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6108)+63)))
	if v6116 == v6110 {
		goto L1728
	} else {
		goto L1729
	}
L1714:
	;
	goto L1713
L1715:
	;
	v6065 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6066 = *(*int32)(unsafe.Add(mBase, uint32(v6065)+60))
	if v6066 == int32(0) {
		v6103 = v6060
		goto L1714
	} else {
		goto L1716
	}
L1716:
	;
	v6069 = v6054 + v6066
	v6070 = *(*int32)(unsafe.Add(mBase, uint32(v6065)+188))
	if base.Ui32(v6070) <= base.Ui32(v6069) {
		goto L1718
	} else {
		goto L1719
	}
L1717:
	;
	v6080 = *(*int32)(unsafe.Add(mBase, uint32(v6065)+196))
	if v6079 == int32(0) {
		v6103 = v6080
		goto L1714
	} else {
		goto L1721
	}
L1718:
	;
	v6072 = *(*int32)(unsafe.Add(mBase, uint32(v6065)+192))
	v6079 = v6072
	goto L1717
L1719:
	;
	goto L1720
L1720:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6065)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6065)+188)) = v6066
	v6077 = F_strchr(m, v6066, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6065)+192)) = v6077
	v6079 = v6077
	goto L1717
L1721:
	;
	if base.Ui32(v6069) <= base.Ui32(v6079) {
		v6103 = v6080
		goto L1714
	} else {
		goto L1722
	}
L1722:
	;
	v6084 = v6079
	v6086 = v6080
	goto L1723
L1723:
	;
	v6089 = int32(1)
	v6090 = v6086 + v6089
	*(*int32)(unsafe.Add(mBase, uint32(v6065)+196)) = v6090
	v6093 = v6084 + v6089
	*(*int32)(unsafe.Add(mBase, uint32(v6065)+188)) = v6093
	v6096 = F_strchr(m, v6093, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6065)+192)) = v6096
	if v6096 == int32(0) {
		v6103 = v6090
		goto L1714
	} else {
		goto L1725
	}
L1724:
	;
	v6103 = v6090
	goto L1714
L1725:
	;
	if base.Ui32(v6096) < base.Ui32(v6069) {
		v6084 = v6096
		v6086 = v6090
		goto L1723
	} else {
		goto L1726
	}
L1726:
	;
	goto L1724
L1727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6056
	v10317 = v241
	goto L5
L1728:
	;
	v6123 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6124 = m.ExcPending
	if v6124 != 0 {
		goto L66
	} else {
		goto L1731
	}
L1729:
	;
	goto L1730
L1730:
	;
	v6150 = *(*int32)(unsafe.Add(mBase, uint32(v6108)+52))
	if v6150 == int32(2278) {
		goto L1739
	} else {
		goto L1740
	}
L1731:
	;
	if v6123 == int32(59) {
		goto L1727
	} else {
		goto L1732
	}
L1732:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v6130 = m.ExcPending
	if v6130 != 0 {
		goto L66
	} else {
		goto L1733
	}
L1733:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v6133 = m.ExcPending
	if v6133 != 0 {
		goto L66
	} else {
		goto L1734
	}
L1734:
	;
	F_errmsg(m, int32(104820), int32(0))
	mBase = m.M
	v6137 = m.ExcPending
	if v6137 != 0 {
		goto L66
	} else {
		goto L1735
	}
L1735:
	;
	F_errhint(m, int32(632187), int32(0))
	mBase = m.M
	v6141 = m.ExcPending
	if v6141 != 0 {
		goto L66
	} else {
		goto L1736
	}
L1736:
	;
	v6142 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	v6143 = F_plpgsql_scanner_errposition(m, v6142, l1)
	mBase = m.M
	v6144 = m.ExcPending
	if v6144 != 0 {
		goto L66
	} else {
		goto L1737
	}
L1737:
	;
	F_errfinish(m, int32(26734), int32(3373), int32(95826))
	mBase = m.M
	v6149 = m.ExcPending
	if v6149 != 0 {
		goto L66
	} else {
		goto L1738
	}
L1738:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1739:
	;
	v6157 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6158 = m.ExcPending
	if v6158 != 0 {
		goto L66
	} else {
		goto L1742
	}
L1740:
	;
	goto L1741
L1741:
	;
	v6185 = *(*int32)(unsafe.Add(mBase, uint32(v6108)+472))
	v6190 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6191 = m.ExcPending
	if v6191 != 0 {
		goto L66
	} else {
		goto L1750
	}
L1742:
	;
	if v6157 == int32(59) {
		goto L1727
	} else {
		goto L1743
	}
L1743:
	;
	v6162 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v6163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6162)+65)))
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v6167 = m.ExcPending
	if v6167 != 0 {
		goto L66
	} else {
		goto L1744
	}
L1744:
	;
	if v6163 == int32(112) {
		goto L33
	} else {
		goto L1745
	}
L1745:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v6172 = m.ExcPending
	if v6172 != 0 {
		goto L66
	} else {
		goto L1746
	}
L1746:
	;
	F_errmsg(m, int32(427072), int32(0))
	mBase = m.M
	v6176 = m.ExcPending
	if v6176 != 0 {
		goto L66
	} else {
		goto L1747
	}
L1747:
	;
	v6177 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	v6178 = F_plpgsql_scanner_errposition(m, v6177, l1)
	mBase = m.M
	v6179 = m.ExcPending
	if v6179 != 0 {
		goto L66
	} else {
		goto L1748
	}
L1748:
	;
	F_errfinish(m, int32(26734), int32(3388), int32(95826))
	mBase = m.M
	v6184 = m.ExcPending
	if v6184 != 0 {
		goto L66
	} else {
		goto L1749
	}
L1749:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1750:
	;
	if int32(0) <= v6185 {
		goto L1751
	} else {
		goto L1752
	}
L1751:
	;
	if v6190 != int32(59) {
		goto L32
	} else {
		goto L1754
	}
L1752:
	;
	goto L1753
L1753:
	;
	if v6190 != int32(277) {
		goto L1755
	} else {
		goto L1756
	}
L1754:
	;
	v6197 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v6198 = *(*int32)(unsafe.Add(mBase, uint32(v6197)+472))
	*(*int32)(unsafe.Add(mBase, uint32(v6056)+16)) = v6198
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6056
	v10317 = v241
	goto L5
L1755:
	;
	F_plpgsql_push_back_token(m, v6190, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6229 = m.ExcPending
	if v6229 != 0 {
		goto L66
	} else {
		goto L1762
	}
L1756:
	;
	v6203 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v6204 = m.ExcPending
	if v6204 != 0 {
		goto L66
	} else {
		goto L1757
	}
L1757:
	;
	if v6203 != int32(59) {
		goto L1755
	} else {
		goto L1758
	}
L1758:
	;
	v6207 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	v6208 = *(*int32)(unsafe.Add(mBase, uint32(v6207)))
	if base.Ui32(int32(4)) < base.Ui32(v6208) {
		goto L1755
	} else {
		goto L1759
	}
L1759:
	;
	if v6208 == int32(3) {
		goto L1755
	} else {
		goto L1760
	}
L1760:
	;
	v6213 = *(*int32)(unsafe.Add(mBase, uint32(v6207)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6056)+16)) = v6213
	v6219 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6220 = m.ExcPending
	if v6220 != 0 {
		goto L66
	} else {
		goto L1761
	}
L1761:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6056
	v10317 = v241
	goto L5
L1762:
	;
	v6231 = int32(0)
	v6235 = int32(1)
	v6243 = F_read_sql_construct(m, int32(59), v6231, v6231, int32(538734), int32(2), v6235, v6235, v6231, v6231, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6244 = m.ExcPending
	if v6244 != 0 {
		goto L66
	} else {
		goto L1763
	}
L1763:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6056)+12)) = v6243
	goto L1727
L1764:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6251))) = int32(14)
	v6255 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v6256 = int32(0)
	if v6255 < v6256 {
		v6299 = v6256
		goto L1766
	} else {
		goto L1767
	}
L1765:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6251)+4)) = v6299
	v6304 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v6305 = *(*int32)(unsafe.Add(mBase, uint32(v6304)+520))
	v6307 = v6305 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6304)+520)) = v6307
	v6309 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6251)+16)) = v6309
	v6311 = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v6251)+12)) = v6311
	*(*int32)(unsafe.Add(mBase, uint32(v6251)+8)) = v6307
	*(*int64)(unsafe.Add(mBase, uint32(v6251)+24)) = v6309
	v6321 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6322 = m.ExcPending
	if v6322 != 0 {
		goto L66
	} else {
		goto L1779
	}
L1766:
	;
	goto L1765
L1767:
	;
	v6261 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6262 = *(*int32)(unsafe.Add(mBase, uint32(v6261)+60))
	if v6262 == int32(0) {
		v6299 = v6256
		goto L1766
	} else {
		goto L1768
	}
L1768:
	;
	v6265 = v6255 + v6262
	v6266 = *(*int32)(unsafe.Add(mBase, uint32(v6261)+188))
	if base.Ui32(v6266) <= base.Ui32(v6265) {
		goto L1770
	} else {
		goto L1771
	}
L1769:
	;
	v6276 = *(*int32)(unsafe.Add(mBase, uint32(v6261)+196))
	if v6275 == int32(0) {
		v6299 = v6276
		goto L1766
	} else {
		goto L1773
	}
L1770:
	;
	v6268 = *(*int32)(unsafe.Add(mBase, uint32(v6261)+192))
	v6275 = v6268
	goto L1769
L1771:
	;
	goto L1772
L1772:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6261)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6261)+188)) = v6262
	v6273 = F_strchr(m, v6262, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6261)+192)) = v6273
	v6275 = v6273
	goto L1769
L1773:
	;
	if base.Ui32(v6265) <= base.Ui32(v6275) {
		v6299 = v6276
		goto L1766
	} else {
		goto L1774
	}
L1774:
	;
	v6280 = v6275
	v6282 = v6276
	goto L1775
L1775:
	;
	v6285 = int32(1)
	v6286 = v6282 + v6285
	*(*int32)(unsafe.Add(mBase, uint32(v6261)+196)) = v6286
	v6289 = v6280 + v6285
	*(*int32)(unsafe.Add(mBase, uint32(v6261)+188)) = v6289
	v6292 = F_strchr(m, v6289, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v6261)+192)) = v6292
	if v6292 == int32(0) {
		v6299 = v6286
		goto L1766
	} else {
		goto L1777
	}
L1776:
	;
	v6299 = v6286
	goto L1766
L1777:
	;
	if base.Ui32(v6292) < base.Ui32(v6265) {
		v6280 = v6292
		v6282 = v6286
		goto L1775
	} else {
		goto L1778
	}
L1778:
	;
	goto L1776
L1779:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v6321
	v6325 = v6251 + int32(16)
	if v6321 <= int32(303) {
		goto L1802
	} else {
		goto L1803
	}
L1780:
	;
	v7264 = *(*int32)(unsafe.Add(mBase, uint32(v6251)+20))
	if v7264 != 0 {
		goto L2066
	} else {
		goto L2067
	}
L1781:
	;
	v6911 = int32(0)
	goto L1956
L1782:
	;
	if v6857 != int32(381) {
		goto L1780
	} else {
		goto L1955
	}
L1783:
	;
	goto L1950
L1784:
	;
	v6794 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6795 = m.ExcPending
	if v6795 != 0 {
		goto L66
	} else {
		goto L1946
	}
L1785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6325))) = v6783
	v6786 = F_plpgsql_recognize_err_condition(m, v6783, int32(0))
	mBase = m.M
	v6787 = m.ExcPending
	if v6787 != 0 {
		goto L66
	} else {
		goto L1945
	}
L1786:
	;
	v6780 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	v6783 = v6780
	goto L1785
L1787:
	;
	v6756 = int32(0)
	goto L1934
L1788:
	;
	v6600 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6601 = m.ExcPending
	if v6601 != 0 {
		goto L66
	} else {
		goto L1891
	}
L1789:
	;
	v6563 = int32(277)
	v6564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v6564 != 0 {
		v6751 = v6563
		goto L1787
	} else {
		goto L1880
	}
L1790:
	;
	switch v6530 - int32(261) {
	case 0:
		goto L1870
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15:
		v6751 = v6530
		goto L1787
	case 14:
		goto L1786
	case 16:
		goto L1789
	default:
		goto L1871
	}
L1791:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6251)+12)) = v6521
	v6527 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6528 = m.ExcPending
	if v6528 != 0 {
		goto L66
	} else {
		goto L1869
	}
L1792:
	;
	v6521 = int32(14)
	goto L1791
L1793:
	;
	v6492 = int32(321882)
	v6495 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1399])))
	v6496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6461))))
	if v6496 == int32(0) {
		v6515 = v6495
		v6516 = v6496
		goto L1861
	} else {
		goto L1862
	}
L1794:
	;
	v6521 = int32(15)
	goto L1791
L1795:
	;
	v6461 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v6461 == int32(0) {
		goto L1789
	} else {
		goto L1850
	}
L1796:
	;
	v6521 = int32(17)
	goto L1791
L1797:
	;
	v6433 = int32(238549)
	v6436 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1400])))
	v6437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6402))))
	if v6437 == int32(0) {
		v6456 = v6436
		v6457 = v6437
		goto L1842
	} else {
		goto L1843
	}
L1798:
	;
	v6521 = int32(18)
	goto L1791
L1799:
	;
	v6402 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v6402 == int32(0) {
		goto L1789
	} else {
		goto L1831
	}
L1800:
	;
	v6521 = int32(19)
	goto L1791
L1801:
	;
	v6342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v6342 != 0 {
		goto L1789
	} else {
		goto L1811
	}
L1802:
	;
	if v6321 == int32(59) {
		goto L1780
	} else {
		goto L1805
	}
L1803:
	;
	goto L1804
L1804:
	;
	switch v6321 - int32(304) {
	case 0:
		goto L1792
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30, 32, 33, 34, 35, 36, 37, 38, 39:
		v6530 = v6321
		goto L1790
	case 12:
		v6521 = v6311
		goto L1791
	case 26:
		goto L1796
	case 31:
		goto L1794
	case 40:
		goto L1798
	default:
		goto L1809
	}
L1805:
	;
	if v6321 == int32(277) {
		goto L1801
	} else {
		goto L1806
	}
L1806:
	;
	if v6321 != 0 {
		v6530 = v6321
		goto L1790
	} else {
		goto L1807
	}
L1807:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(246872))
	mBase = m.M
	v6337 = m.ExcPending
	if v6337 != 0 {
		goto L66
	} else {
		goto L1808
	}
L1808:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1809:
	;
	if v6321 == int32(383) {
		goto L1800
	} else {
		goto L1810
	}
L1810:
	;
	v6530 = v6321
	goto L1790
L1811:
	;
	v6343 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v6343 == int32(0) {
		goto L1789
	} else {
		goto L1812
	}
L1812:
	;
	v6346 = int32(244285)
	v6349 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1401])))
	v6350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6343))))
	if v6350 == int32(0) {
		v6369 = v6349
		v6370 = v6350
		goto L1814
	} else {
		goto L1815
	}
L1813:
	;
	if v6370-v6369 == int32(0) {
		v6521 = v6311
		goto L1791
	} else {
		goto L1821
	}
L1814:
	;
	goto L1813
L1815:
	;
	if v6349 != v6350 {
		v6369 = v6349
		v6370 = v6350
		goto L1814
	} else {
		goto L1816
	}
L1816:
	;
	v6354 = v6343
	v6355 = v6346
	goto L1817
L1817:
	;
	v6358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6355)+1)))
	v6359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6354)+1)))
	if v6359 == int32(0) {
		v6369 = v6358
		v6370 = v6359
		goto L1814
	} else {
		goto L1819
	}
L1818:
	;
	v6369 = v6358
	v6370 = v6359
	goto L1814
L1819:
	;
	v6362 = int32(1)
	if v6358 == v6359 {
		v6354 = v6354 + v6362
		v6355 = v6355 + v6362
		goto L1817
	} else {
		goto L1820
	}
L1820:
	;
	goto L1818
L1821:
	;
	v6374 = int32(329626)
	v6377 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1402])))
	v6378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6343))))
	if v6378 == int32(0) {
		v6397 = v6377
		v6398 = v6378
		goto L1823
	} else {
		goto L1824
	}
L1822:
	;
	if v6398-v6397 != 0 {
		goto L1799
	} else {
		goto L1830
	}
L1823:
	;
	goto L1822
L1824:
	;
	if v6377 != v6378 {
		v6397 = v6377
		v6398 = v6378
		goto L1823
	} else {
		goto L1825
	}
L1825:
	;
	v6382 = v6343
	v6383 = v6374
	goto L1826
L1826:
	;
	v6386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6383)+1)))
	v6387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6382)+1)))
	if v6387 == int32(0) {
		v6397 = v6386
		v6398 = v6387
		goto L1823
	} else {
		goto L1828
	}
L1827:
	;
	v6397 = v6386
	v6398 = v6387
	goto L1823
L1828:
	;
	v6390 = int32(1)
	if v6386 == v6387 {
		v6382 = v6382 + v6390
		v6383 = v6383 + v6390
		goto L1826
	} else {
		goto L1829
	}
L1829:
	;
	goto L1827
L1830:
	;
	goto L1800
L1831:
	;
	v6405 = int32(411812)
	v6408 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1403])))
	v6409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6402))))
	if v6409 == int32(0) {
		v6428 = v6408
		v6429 = v6409
		goto L1833
	} else {
		goto L1834
	}
L1832:
	;
	if v6429-v6428 != 0 {
		goto L1797
	} else {
		goto L1840
	}
L1833:
	;
	goto L1832
L1834:
	;
	if v6408 != v6409 {
		v6428 = v6408
		v6429 = v6409
		goto L1833
	} else {
		goto L1835
	}
L1835:
	;
	v6413 = v6402
	v6414 = v6405
	goto L1836
L1836:
	;
	v6417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6414)+1)))
	v6418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6413)+1)))
	if v6418 == int32(0) {
		v6428 = v6417
		v6429 = v6418
		goto L1833
	} else {
		goto L1838
	}
L1837:
	;
	v6428 = v6417
	v6429 = v6418
	goto L1833
L1838:
	;
	v6421 = int32(1)
	if v6417 == v6418 {
		v6413 = v6413 + v6421
		v6414 = v6414 + v6421
		goto L1836
	} else {
		goto L1839
	}
L1839:
	;
	goto L1837
L1840:
	;
	goto L1798
L1841:
	;
	if v6457-v6456 != 0 {
		goto L1795
	} else {
		goto L1849
	}
L1842:
	;
	goto L1841
L1843:
	;
	if v6436 != v6437 {
		v6456 = v6436
		v6457 = v6437
		goto L1842
	} else {
		goto L1844
	}
L1844:
	;
	v6441 = v6402
	v6442 = v6433
	goto L1845
L1845:
	;
	v6445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6442)+1)))
	v6446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6441)+1)))
	if v6446 == int32(0) {
		v6456 = v6445
		v6457 = v6446
		goto L1842
	} else {
		goto L1847
	}
L1846:
	;
	v6456 = v6445
	v6457 = v6446
	goto L1842
L1847:
	;
	v6449 = int32(1)
	if v6445 == v6446 {
		v6441 = v6441 + v6449
		v6442 = v6442 + v6449
		goto L1845
	} else {
		goto L1848
	}
L1848:
	;
	goto L1846
L1849:
	;
	goto L1796
L1850:
	;
	v6464 = int32(322479)
	v6467 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1404])))
	v6468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6461))))
	if v6468 == int32(0) {
		v6487 = v6467
		v6488 = v6468
		goto L1852
	} else {
		goto L1853
	}
L1851:
	;
	if v6488-v6487 != 0 {
		goto L1793
	} else {
		goto L1859
	}
L1852:
	;
	goto L1851
L1853:
	;
	if v6467 != v6468 {
		v6487 = v6467
		v6488 = v6468
		goto L1852
	} else {
		goto L1854
	}
L1854:
	;
	v6472 = v6461
	v6473 = v6464
	goto L1855
L1855:
	;
	v6476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6473)+1)))
	v6477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6472)+1)))
	if v6477 == int32(0) {
		v6487 = v6476
		v6488 = v6477
		goto L1852
	} else {
		goto L1857
	}
L1856:
	;
	v6487 = v6476
	v6488 = v6477
	goto L1852
L1857:
	;
	v6480 = int32(1)
	if v6476 == v6477 {
		v6472 = v6472 + v6480
		v6473 = v6473 + v6480
		goto L1855
	} else {
		goto L1858
	}
L1858:
	;
	goto L1856
L1859:
	;
	goto L1794
L1860:
	;
	if v6516-v6515 != 0 {
		goto L1789
	} else {
		goto L1868
	}
L1861:
	;
	goto L1860
L1862:
	;
	if v6495 != v6496 {
		v6515 = v6495
		v6516 = v6496
		goto L1861
	} else {
		goto L1863
	}
L1863:
	;
	v6500 = v6461
	v6501 = v6492
	goto L1864
L1864:
	;
	v6504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6501)+1)))
	v6505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6500)+1)))
	if v6505 == int32(0) {
		v6515 = v6504
		v6516 = v6505
		goto L1861
	} else {
		goto L1866
	}
L1865:
	;
	v6515 = v6504
	v6516 = v6505
	goto L1861
L1866:
	;
	v6508 = int32(1)
	if v6504 == v6505 {
		v6500 = v6500 + v6508
		v6501 = v6501 + v6508
		goto L1864
	} else {
		goto L1867
	}
L1867:
	;
	goto L1865
L1868:
	;
	goto L1792
L1869:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v6527
	v6530 = v6527
	goto L1790
L1870:
	;
	v6542 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	*(*int32)(unsafe.Add(mBase, uint32(v6251)+20)) = v6542
	v6548 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6549 = m.ExcPending
	if v6549 != 0 {
		goto L66
	} else {
		goto L1875
	}
L1871:
	;
	switch v6530 - int32(371) {
	case 0:
		goto L1788
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v6751 = v6530
		goto L1787
	case 10:
		goto L1781
	default:
		goto L1872
	}
L1872:
	;
	if v6530 != 0 {
		v6751 = v6530
		goto L1787
	} else {
		goto L1873
	}
L1873:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(246872))
	mBase = m.M
	v6541 = m.ExcPending
	if v6541 != 0 {
		goto L66
	} else {
		goto L1874
	}
L1874:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1875:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v6548
	switch v6548 - int32(44) {
	case 0:
		goto L1783
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L1876
	case 15:
		v6857 = v6548
		goto L1782
	default:
		goto L1877
	}
L1876:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v6560 = m.ExcPending
	if v6560 != 0 {
		goto L66
	} else {
		goto L1879
	}
L1877:
	;
	if v6548 == int32(381) {
		v6857 = v6548
		goto L1782
	} else {
		goto L1878
	}
L1878:
	;
	goto L1876
L1879:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1880:
	;
	v6565 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v6565 == int32(0) {
		v6751 = v6563
		goto L1787
	} else {
		goto L1881
	}
L1881:
	;
	v6568 = int32(346115)
	v6571 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1365])))
	v6572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6565))))
	if v6572 == int32(0) {
		v6591 = v6571
		v6592 = v6572
		goto L1883
	} else {
		goto L1884
	}
L1882:
	;
	if v6592-v6591 != 0 {
		v6751 = v6563
		goto L1787
	} else {
		goto L1890
	}
L1883:
	;
	goto L1882
L1884:
	;
	if v6571 != v6572 {
		v6591 = v6571
		v6592 = v6572
		goto L1883
	} else {
		goto L1885
	}
L1885:
	;
	v6576 = v6565
	v6577 = v6568
	goto L1886
L1886:
	;
	v6580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6577)+1)))
	v6581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6576)+1)))
	if v6581 == int32(0) {
		v6591 = v6580
		v6592 = v6581
		goto L1883
	} else {
		goto L1888
	}
L1887:
	;
	v6591 = v6580
	v6592 = v6581
	goto L1883
L1888:
	;
	v6584 = int32(1)
	if v6580 == v6581 {
		v6576 = v6576 + v6584
		v6577 = v6577 + v6584
		goto L1886
	} else {
		goto L1889
	}
L1889:
	;
	goto L1887
L1890:
	;
	goto L1788
L1891:
	;
	if v6600 != int32(261) {
		goto L31
	} else {
		goto L1892
	}
L1892:
	;
	v6604 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	if v6604&int32(3) == int32(0) {
		v6628 = v6604
		goto L1895
	} else {
		goto L1896
	}
L1893:
	;
	if v6661 != int32(5) {
		goto L30
	} else {
		goto L1910
	}
L1894:
	;
	v6661 = v6653 - v6604
	goto L1893
L1895:
	;
	v6632 = v6628
	goto L1904
L1896:
	;
	v6612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6604))))
	if v6612 == int32(0) {
		goto L1897
	} else {
		goto L1898
	}
L1897:
	;
	v6661 = int32(0)
	goto L1893
L1898:
	;
	goto L1899
L1899:
	;
	v6617 = v6604
	goto L1900
L1900:
	;
	v6621 = v6617 + int32(1)
	if v6621&int32(3) == int32(0) {
		v6628 = v6621
		goto L1895
	} else {
		goto L1902
	}
L1901:
	;
	v6653 = v6621
	goto L1894
L1902:
	;
	v6626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6621))))
	if v6626 != 0 {
		v6617 = v6621
		goto L1900
	} else {
		goto L1903
	}
L1903:
	;
	goto L1901
L1904:
	;
	v6638 = *(*int32)(unsafe.Add(mBase, uint32(v6632)))
	v6641 = int32(-2139062144)
	if (int32(16843008)-v6638|v6638)&v6641 == v6641 {
		v6632 = v6632 + int32(4)
		goto L1904
	} else {
		goto L1906
	}
L1905:
	;
	v6647 = v6632
	goto L1907
L1906:
	;
	goto L1905
L1907:
	;
	v6651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6647))))
	if v6651 != 0 {
		v6647 = v6647 + int32(1)
		goto L1907
	} else {
		goto L1909
	}
L1908:
	;
	v6653 = v6647
	goto L1894
L1909:
	;
	goto L1908
L1910:
	;
	v6664 = int32(500110)
	v6668 = m.G0
	v6670 = v6668 - int32(32)
	v6671 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6670)+24)) = v6671
	*(*int64)(unsafe.Add(mBase, uint32(v6670)+16)) = v6671
	*(*int64)(unsafe.Add(mBase, uint32(v6670)+8)) = v6671
	*(*int64)(unsafe.Add(mBase, uint32(v6670))) = v6671
	v6679 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1405])))
	if v6679 == int32(0) {
		goto L1912
	} else {
		goto L1913
	}
L1911:
	;
	if v6747 != int32(5) {
		goto L29
	} else {
		goto L1932
	}
L1912:
	;
	v6747 = int32(0)
	goto L1911
L1913:
	;
	goto L1914
L1914:
	;
	v6683 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1406])))
	if v6683 == int32(0) {
		goto L1915
	} else {
		goto L1916
	}
L1915:
	;
	v6687 = v6604
	goto L1918
L1916:
	;
	goto L1917
L1917:
	;
	v6697 = v6664
	v6698 = v6679
	goto L1921
L1918:
	;
	v6693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6687))))
	if v6693 == v6679 {
		v6687 = v6687 + int32(1)
		goto L1918
	} else {
		goto L1920
	}
L1919:
	;
	v6747 = v6687 - v6604
	goto L1911
L1920:
	;
	goto L1919
L1921:
	;
	v6705 = v6670 + int32(base.Ui32(v6698)>>(uint(int32(3))%32))&int32(28)
	v6706 = *(*int32)(unsafe.Add(mBase, uint32(v6705)))
	v6707 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6705))) = v6706 | v6707<<(uint(v6698)%32)
	v6711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6697)+1)))
	if v6711 != 0 {
		v6697 = v6697 + v6707
		v6698 = v6711
		goto L1921
	} else {
		goto L1923
	}
L1922:
	;
	v6714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6604))))
	if v6714 == int32(0) {
		v6739 = v6604
		goto L1924
	} else {
		goto L1925
	}
L1923:
	;
	goto L1922
L1924:
	;
	v6747 = v6739 - v6604
	goto L1911
L1925:
	;
	v6718 = v6604
	v6719 = v6714
	goto L1926
L1926:
	;
	v6727 = *(*int32)(unsafe.Add(mBase, uint32(v6670+int32(base.Ui32(v6719)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v6727)>>(uint(v6719)%32))&int32(1) == int32(0) {
		goto L1928
	} else {
		goto L1929
	}
L1927:
	;
	v6739 = v6735
	goto L1924
L1928:
	;
	v6739 = v6718
	goto L1924
L1929:
	;
	goto L1930
L1930:
	;
	v6733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6718)+1)))
	v6735 = v6718 + int32(1)
	if v6733 != 0 {
		v6718 = v6735
		v6719 = v6733
		goto L1926
	} else {
		goto L1931
	}
L1931:
	;
	goto L1927
L1932:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6325))) = v6604
	goto L1784
L1933:
	;
	if v6751 == v6762 {
		goto L1940
	} else {
		goto L1941
	}
L1934:
	;
	v6762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6756<<(uint(int32(1))%32))+uint32(_consts[1372]))))
	v6763 = base.B2i32(v6751 == v6762)
	if v6763 == int32(0) {
		goto L1936
	} else {
		goto L1937
	}
L1935:
	;
	goto L1933
L1936:
	;
	v6767 = v6756 + int32(1)
	if v6767 != int32(83) {
		v6756 = v6767
		goto L1934
	} else {
		goto L1939
	}
L1937:
	;
	goto L1938
L1938:
	;
	goto L1935
L1939:
	;
	goto L1938
L1940:
	;
	v6771 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	v6772 = F_pstrdup(m, v6771)
	mBase = m.M
	v6773 = m.ExcPending
	if v6773 != 0 {
		goto L66
	} else {
		goto L1943
	}
L1941:
	;
	goto L1942
L1942:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v6779 = m.ExcPending
	if v6779 != 0 {
		goto L66
	} else {
		goto L1944
	}
L1943:
	;
	v6783 = v6772
	goto L1785
L1944:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1945:
	;
	goto L1784
L1946:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v6794
	if v6794 == int32(59) {
		v6857 = v6794
		goto L1782
	} else {
		goto L1947
	}
L1947:
	;
	if v6794 == int32(381) {
		v6857 = v6794
		goto L1782
	} else {
		goto L1948
	}
L1948:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v6806 = m.ExcPending
	if v6806 != 0 {
		goto L66
	} else {
		goto L1949
	}
L1949:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1950:
	;
	v6836 = int32(1)
	v6845 = F_read_sql_construct(m, int32(44), int32(59), int32(381), int32(528068), int32(2), v6836, v6836, int32(0), v27+int32(4768), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6846 = m.ExcPending
	if v6846 != 0 {
		goto L66
	} else {
		goto L1952
	}
L1951:
	;
	v6857 = v6851
	goto L1782
L1952:
	;
	v6847 = *(*int32)(unsafe.Add(mBase, uint32(v6251)+24))
	v6848 = F_lappend(m, v6847, v6845)
	mBase = m.M
	v6849 = m.ExcPending
	if v6849 != 0 {
		goto L66
	} else {
		goto L1953
	}
L1953:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6251)+24)) = v6848
	v6851 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378])))
	if v6851 == int32(44) {
		goto L1950
	} else {
		goto L1954
	}
L1954:
	;
	goto L1951
L1955:
	;
	goto L1781
L1956:
	;
	v6933 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v6934 = m.ExcPending
	if v6934 != 0 {
		goto L66
	} else {
		goto L1958
	}
L1957:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6251)+28)) = v7233
	goto L1780
L1958:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = v6933
	if v6933 == int32(0) {
		goto L28
	} else {
		goto L1959
	}
L1959:
	;
	v6939 = F_palloc(m, int32(8))
	mBase = m.M
	v6940 = m.ExcPending
	if v6940 != 0 {
		goto L66
	} else {
		goto L1960
	}
L1960:
	;
	v6941 = int32(0)
	v6942 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	switch v6942 - int32(277) {
	case 0:
		goto L1977
	default:
		goto L27
	case 17:
		goto L1970
	case 21:
		goto L1968
	case 26:
		goto L1966
	case 30:
		goto L1974
	case 37:
		v7201 = v6941
		goto L1961
	case 49:
		goto L1972
	case 61:
		goto L1976
	case 90:
		goto L1962
	case 97:
		goto L1964
	}
L1961:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6939))) = v7201
	v7208 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7209 = m.ExcPending
	if v7209 != 0 {
		goto L66
	} else {
		goto L2061
	}
L1962:
	;
	v7201 = int32(8)
	goto L1961
L1963:
	;
	v7173 = int32(498699)
	v7176 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1407])))
	v7177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v7177 == int32(0) {
		v7196 = v7176
		v7197 = v7177
		goto L2053
	} else {
		goto L2054
	}
L1964:
	;
	v7201 = int32(7)
	goto L1961
L1965:
	;
	v7145 = int32(389375)
	v7148 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1408])))
	v7149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v7149 == int32(0) {
		v7168 = v7148
		v7169 = v7149
		goto L2044
	} else {
		goto L2045
	}
L1966:
	;
	v7201 = int32(6)
	goto L1961
L1967:
	;
	v7117 = int32(360598)
	v7120 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1409])))
	v7121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v7121 == int32(0) {
		v7140 = v7120
		v7141 = v7121
		goto L2035
	} else {
		goto L2036
	}
L1968:
	;
	v7201 = int32(5)
	goto L1961
L1969:
	;
	v7089 = int32(89437)
	v7092 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1410])))
	v7093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v7093 == int32(0) {
		v7112 = v7092
		v7113 = v7093
		goto L2026
	} else {
		goto L2027
	}
L1970:
	;
	v7201 = int32(4)
	goto L1961
L1971:
	;
	v7061 = int32(270971)
	v7064 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1411])))
	v7065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v7065 == int32(0) {
		v7084 = v7064
		v7085 = v7065
		goto L2017
	} else {
		goto L2018
	}
L1972:
	;
	v7201 = int32(3)
	goto L1961
L1973:
	;
	v7033 = int32(88232)
	v7036 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1412])))
	v7037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v7037 == int32(0) {
		v7056 = v7036
		v7057 = v7037
		goto L2008
	} else {
		goto L2009
	}
L1974:
	;
	v7201 = int32(2)
	goto L1961
L1975:
	;
	v7005 = int32(301274)
	v7008 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1413])))
	v7009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v7009 == int32(0) {
		v7028 = v7008
		v7029 = v7009
		goto L1999
	} else {
		goto L2000
	}
L1976:
	;
	v7201 = int32(1)
	goto L1961
L1977:
	;
	v6945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v6945 != 0 {
		goto L27
	} else {
		goto L1978
	}
L1978:
	;
	v6946 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v6946 == int32(0) {
		goto L27
	} else {
		goto L1979
	}
L1979:
	;
	v6949 = int32(407487)
	v6952 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1414])))
	v6953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v6953 == int32(0) {
		v6972 = v6952
		v6973 = v6953
		goto L1981
	} else {
		goto L1982
	}
L1980:
	;
	if v6973-v6972 == int32(0) {
		v7201 = v6941
		goto L1961
	} else {
		goto L1988
	}
L1981:
	;
	goto L1980
L1982:
	;
	if v6952 != v6953 {
		v6972 = v6952
		v6973 = v6953
		goto L1981
	} else {
		goto L1983
	}
L1983:
	;
	v6957 = v6946
	v6958 = v6949
	goto L1984
L1984:
	;
	v6961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6958)+1)))
	v6962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6957)+1)))
	if v6962 == int32(0) {
		v6972 = v6961
		v6973 = v6962
		goto L1981
	} else {
		goto L1986
	}
L1985:
	;
	v6972 = v6961
	v6973 = v6962
	goto L1981
L1986:
	;
	v6965 = int32(1)
	if v6961 == v6962 {
		v6957 = v6957 + v6965
		v6958 = v6958 + v6965
		goto L1984
	} else {
		goto L1987
	}
L1987:
	;
	goto L1985
L1988:
	;
	v6977 = int32(399277)
	v6980 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1415])))
	v6981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6946))))
	if v6981 == int32(0) {
		v7000 = v6980
		v7001 = v6981
		goto L1990
	} else {
		goto L1991
	}
L1989:
	;
	if v7001-v7000 != 0 {
		goto L1975
	} else {
		goto L1997
	}
L1990:
	;
	goto L1989
L1991:
	;
	if v6980 != v6981 {
		v7000 = v6980
		v7001 = v6981
		goto L1990
	} else {
		goto L1992
	}
L1992:
	;
	v6985 = v6946
	v6986 = v6977
	goto L1993
L1993:
	;
	v6989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6986)+1)))
	v6990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6985)+1)))
	if v6990 == int32(0) {
		v7000 = v6989
		v7001 = v6990
		goto L1990
	} else {
		goto L1995
	}
L1994:
	;
	v7000 = v6989
	v7001 = v6990
	goto L1990
L1995:
	;
	v6993 = int32(1)
	if v6989 == v6990 {
		v6985 = v6985 + v6993
		v6986 = v6986 + v6993
		goto L1993
	} else {
		goto L1996
	}
L1996:
	;
	goto L1994
L1997:
	;
	goto L1976
L1998:
	;
	if v7029-v7028 != 0 {
		goto L1973
	} else {
		goto L2006
	}
L1999:
	;
	goto L1998
L2000:
	;
	if v7008 != v7009 {
		v7028 = v7008
		v7029 = v7009
		goto L1999
	} else {
		goto L2001
	}
L2001:
	;
	v7013 = v6946
	v7014 = v7005
	goto L2002
L2002:
	;
	v7017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7014)+1)))
	v7018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7013)+1)))
	if v7018 == int32(0) {
		v7028 = v7017
		v7029 = v7018
		goto L1999
	} else {
		goto L2004
	}
L2003:
	;
	v7028 = v7017
	v7029 = v7018
	goto L1999
L2004:
	;
	v7021 = int32(1)
	if v7017 == v7018 {
		v7013 = v7013 + v7021
		v7014 = v7014 + v7021
		goto L2002
	} else {
		goto L2005
	}
L2005:
	;
	goto L2003
L2006:
	;
	goto L1974
L2007:
	;
	if v7057-v7056 != 0 {
		goto L1971
	} else {
		goto L2015
	}
L2008:
	;
	goto L2007
L2009:
	;
	if v7036 != v7037 {
		v7056 = v7036
		v7057 = v7037
		goto L2008
	} else {
		goto L2010
	}
L2010:
	;
	v7041 = v6946
	v7042 = v7033
	goto L2011
L2011:
	;
	v7045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7042)+1)))
	v7046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7041)+1)))
	if v7046 == int32(0) {
		v7056 = v7045
		v7057 = v7046
		goto L2008
	} else {
		goto L2013
	}
L2012:
	;
	v7056 = v7045
	v7057 = v7046
	goto L2008
L2013:
	;
	v7049 = int32(1)
	if v7045 == v7046 {
		v7041 = v7041 + v7049
		v7042 = v7042 + v7049
		goto L2011
	} else {
		goto L2014
	}
L2014:
	;
	goto L2012
L2015:
	;
	goto L1972
L2016:
	;
	if v7085-v7084 != 0 {
		goto L1969
	} else {
		goto L2024
	}
L2017:
	;
	goto L2016
L2018:
	;
	if v7064 != v7065 {
		v7084 = v7064
		v7085 = v7065
		goto L2017
	} else {
		goto L2019
	}
L2019:
	;
	v7069 = v6946
	v7070 = v7061
	goto L2020
L2020:
	;
	v7073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7070)+1)))
	v7074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7069)+1)))
	if v7074 == int32(0) {
		v7084 = v7073
		v7085 = v7074
		goto L2017
	} else {
		goto L2022
	}
L2021:
	;
	v7084 = v7073
	v7085 = v7074
	goto L2017
L2022:
	;
	v7077 = int32(1)
	if v7073 == v7074 {
		v7069 = v7069 + v7077
		v7070 = v7070 + v7077
		goto L2020
	} else {
		goto L2023
	}
L2023:
	;
	goto L2021
L2024:
	;
	goto L1970
L2025:
	;
	if v7113-v7112 != 0 {
		goto L1967
	} else {
		goto L2033
	}
L2026:
	;
	goto L2025
L2027:
	;
	if v7092 != v7093 {
		v7112 = v7092
		v7113 = v7093
		goto L2026
	} else {
		goto L2028
	}
L2028:
	;
	v7097 = v6946
	v7098 = v7089
	goto L2029
L2029:
	;
	v7101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7098)+1)))
	v7102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7097)+1)))
	if v7102 == int32(0) {
		v7112 = v7101
		v7113 = v7102
		goto L2026
	} else {
		goto L2031
	}
L2030:
	;
	v7112 = v7101
	v7113 = v7102
	goto L2026
L2031:
	;
	v7105 = int32(1)
	if v7101 == v7102 {
		v7097 = v7097 + v7105
		v7098 = v7098 + v7105
		goto L2029
	} else {
		goto L2032
	}
L2032:
	;
	goto L2030
L2033:
	;
	goto L1968
L2034:
	;
	if v7141-v7140 != 0 {
		goto L1965
	} else {
		goto L2042
	}
L2035:
	;
	goto L2034
L2036:
	;
	if v7120 != v7121 {
		v7140 = v7120
		v7141 = v7121
		goto L2035
	} else {
		goto L2037
	}
L2037:
	;
	v7125 = v6946
	v7126 = v7117
	goto L2038
L2038:
	;
	v7129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7126)+1)))
	v7130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7125)+1)))
	if v7130 == int32(0) {
		v7140 = v7129
		v7141 = v7130
		goto L2035
	} else {
		goto L2040
	}
L2039:
	;
	v7140 = v7129
	v7141 = v7130
	goto L2035
L2040:
	;
	v7133 = int32(1)
	if v7129 == v7130 {
		v7125 = v7125 + v7133
		v7126 = v7126 + v7133
		goto L2038
	} else {
		goto L2041
	}
L2041:
	;
	goto L2039
L2042:
	;
	goto L1966
L2043:
	;
	if v7169-v7168 != 0 {
		goto L1963
	} else {
		goto L2051
	}
L2044:
	;
	goto L2043
L2045:
	;
	if v7148 != v7149 {
		v7168 = v7148
		v7169 = v7149
		goto L2044
	} else {
		goto L2046
	}
L2046:
	;
	v7153 = v6946
	v7154 = v7145
	goto L2047
L2047:
	;
	v7157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7154)+1)))
	v7158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7153)+1)))
	if v7158 == int32(0) {
		v7168 = v7157
		v7169 = v7158
		goto L2044
	} else {
		goto L2049
	}
L2048:
	;
	v7168 = v7157
	v7169 = v7158
	goto L2044
L2049:
	;
	v7161 = int32(1)
	if v7157 == v7158 {
		v7153 = v7153 + v7161
		v7154 = v7154 + v7161
		goto L2047
	} else {
		goto L2050
	}
L2050:
	;
	goto L2048
L2051:
	;
	goto L1964
L2052:
	;
	if v7197-v7196 != 0 {
		goto L27
	} else {
		goto L2060
	}
L2053:
	;
	goto L2052
L2054:
	;
	if v7176 != v7177 {
		v7196 = v7176
		v7197 = v7177
		goto L2053
	} else {
		goto L2055
	}
L2055:
	;
	v7181 = v6946
	v7182 = v7173
	goto L2056
L2056:
	;
	v7185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7182)+1)))
	v7186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7181)+1)))
	if v7186 == int32(0) {
		v7196 = v7185
		v7197 = v7186
		goto L2053
	} else {
		goto L2058
	}
L2057:
	;
	v7196 = v7185
	v7197 = v7186
	goto L2053
L2058:
	;
	v7189 = int32(1)
	if v7185 == v7186 {
		v7181 = v7181 + v7189
		v7182 = v7182 + v7189
		goto L2056
	} else {
		goto L2059
	}
L2059:
	;
	goto L2057
L2060:
	;
	goto L1962
L2061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = v7208
	if base.B2i32(v7208 != int32(61))&base.B2i32(v7208 != int32(270)) != 0 {
		goto L26
	} else {
		goto L2062
	}
L2062:
	;
	v7218 = int32(0)
	v7221 = int32(1)
	v7230 = F_read_sql_construct(m, int32(44), int32(59), v7218, int32(538729), int32(2), v7221, v7221, v7218, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7231 = m.ExcPending
	if v7231 != 0 {
		goto L66
	} else {
		goto L2063
	}
L2063:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6939)+4)) = v7230
	v7233 = F_lappend(m, v6911, v6939)
	mBase = m.M
	v7234 = m.ExcPending
	if v7234 != 0 {
		goto L66
	} else {
		goto L2064
	}
L2064:
	;
	v7235 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v7235 != int32(59) {
		v6911 = v7233
		goto L1956
	} else {
		goto L2065
	}
L2065:
	;
	goto L1957
L2066:
	;
	v7268 = v7264
	v7271 = int32(0)
	goto L2070
L2067:
	;
	goto L2068
L2068:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v6251
	v10317 = v241
	goto L5
L2069:
	;
	if v7310 < v7271 {
		goto L25
	} else {
		goto L2084
	}
L2070:
	;
	v7289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7268))))
	if v7289 != int32(37) {
		goto L2074
	} else {
		goto L2075
	}
L2071:
	;
	v7308 = *(*int32)(unsafe.Add(mBase, uint32(v7292)+4))
	if v7271 < v7308 {
		goto L11
	} else {
		goto L2083
	}
L2072:
	;
	goto L2071
L2073:
	;
	v7268 = v7303 + int32(1)
	v7271 = v7305
	goto L2070
L2074:
	;
	if v7289 != 0 {
		v7303 = v7268
		v7305 = v7271
		goto L2073
	} else {
		goto L2077
	}
L2075:
	;
	goto L2076
L2076:
	;
	v7298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7268)+1)))
	v7300 = base.B2i32(v7298 != int32(37))
	if v7298 != int32(37) {
		goto L2080
	} else {
		goto L2081
	}
L2077:
	;
	v7292 = *(*int32)(unsafe.Add(mBase, uint32(v6251)+24))
	if v7292 != 0 {
		goto L2072
	} else {
		goto L2078
	}
L2078:
	;
	v7293 = int32(0)
	if v7293 <= v7271 {
		v7310 = v7293
		goto L2069
	} else {
		goto L2079
	}
L2079:
	;
	goto L11
L2080:
	;
	v7301 = v7268
	goto L2082
L2081:
	;
	v7301 = v7268 + int32(1)
	goto L2082
L2082:
	;
	v7303 = v7301
	v7305 = v7300 + v7271
	goto L2073
L2083:
	;
	v7310 = v7308
	goto L2069
L2084:
	;
	goto L2068
L2085:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7338))) = int32(15)
	v7342 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7343 = int32(0)
	if v7342 < v7343 {
		v7386 = v7343
		goto L2087
	} else {
		goto L2088
	}
L2086:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7338)+4)) = v7386
	v7391 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v7392 = *(*int32)(unsafe.Add(mBase, uint32(v7391)+520))
	v7393 = int32(1)
	v7394 = v7392 + v7393
	*(*int32)(unsafe.Add(mBase, uint32(v7391)+520)) = v7394
	*(*int32)(unsafe.Add(mBase, uint32(v7338)+8)) = v7394
	v7399 = int32(0)
	v7411 = F_read_sql_construct(m, int32(44), int32(59), v7399, int32(538729), int32(2), v7393, v7393, v7399, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7412 = m.ExcPending
	if v7412 != 0 {
		goto L66
	} else {
		goto L2100
	}
L2087:
	;
	goto L2086
L2088:
	;
	v7348 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7349 = *(*int32)(unsafe.Add(mBase, uint32(v7348)+60))
	if v7349 == int32(0) {
		v7386 = v7343
		goto L2087
	} else {
		goto L2089
	}
L2089:
	;
	v7352 = v7342 + v7349
	v7353 = *(*int32)(unsafe.Add(mBase, uint32(v7348)+188))
	if base.Ui32(v7353) <= base.Ui32(v7352) {
		goto L2091
	} else {
		goto L2092
	}
L2090:
	;
	v7363 = *(*int32)(unsafe.Add(mBase, uint32(v7348)+196))
	if v7362 == int32(0) {
		v7386 = v7363
		goto L2087
	} else {
		goto L2094
	}
L2091:
	;
	v7355 = *(*int32)(unsafe.Add(mBase, uint32(v7348)+192))
	v7362 = v7355
	goto L2090
L2092:
	;
	goto L2093
L2093:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7348)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7348)+188)) = v7349
	v7360 = F_strchr(m, v7349, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7348)+192)) = v7360
	v7362 = v7360
	goto L2090
L2094:
	;
	if base.Ui32(v7352) <= base.Ui32(v7362) {
		v7386 = v7363
		goto L2087
	} else {
		goto L2095
	}
L2095:
	;
	v7367 = v7362
	v7369 = v7363
	goto L2096
L2096:
	;
	v7372 = int32(1)
	v7373 = v7369 + v7372
	*(*int32)(unsafe.Add(mBase, uint32(v7348)+196)) = v7373
	v7376 = v7367 + v7372
	*(*int32)(unsafe.Add(mBase, uint32(v7348)+188)) = v7376
	v7379 = F_strchr(m, v7376, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7348)+192)) = v7379
	if v7379 == int32(0) {
		v7386 = v7373
		goto L2087
	} else {
		goto L2098
	}
L2097:
	;
	v7386 = v7373
	goto L2087
L2098:
	;
	if base.Ui32(v7379) < base.Ui32(v7352) {
		v7367 = v7379
		v7369 = v7373
		goto L2096
	} else {
		goto L2099
	}
L2099:
	;
	goto L2097
L2100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7338)+12)) = v7411
	v7414 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v7414 == int32(44) {
		goto L2101
	} else {
		goto L2102
	}
L2101:
	;
	v7418 = int32(0)
	v7422 = int32(1)
	v7430 = F_read_sql_construct(m, int32(59), v7418, v7418, int32(538734), int32(2), v7422, v7422, v7418, v7418, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7431 = m.ExcPending
	if v7431 != 0 {
		goto L66
	} else {
		goto L2104
	}
L2102:
	;
	v7433 = int32(0)
	goto L2103
L2103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7338)+16)) = v7433
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7338
	v10317 = v241
	goto L5
L2104:
	;
	v7433 = v7430
	goto L2103
L2105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7455
	v10317 = v241
	goto L5
L2106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7465
	v10317 = v241
	goto L5
L2107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7475
	v10317 = v241
	goto L5
L2108:
	;
	F_plpgsql_push_back_token(m, v7482, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7489 = m.ExcPending
	if v7489 != 0 {
		goto L66
	} else {
		goto L2109
	}
L2109:
	;
	switch v7482 - int32(46) {
	case 0, 15:
		goto L24
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2110
	default:
		goto L2111
	}
L2110:
	;
	v7497 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7502 = F_make_execsql_stmt(m, int32(275), v7497, v132, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7503 = m.ExcPending
	if v7503 != 0 {
		goto L66
	} else {
		goto L2114
	}
L2111:
	;
	if v7482 == int32(270) {
		goto L24
	} else {
		goto L2112
	}
L2112:
	;
	if v7482 == int32(91) {
		goto L24
	} else {
		goto L2113
	}
L2113:
	;
	goto L2110
L2114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7502
	v10317 = v241
	goto L5
L2115:
	;
	F_plpgsql_push_back_token(m, v7509, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7516 = m.ExcPending
	if v7516 != 0 {
		goto L66
	} else {
		goto L2116
	}
L2116:
	;
	switch v7509 - int32(46) {
	case 0, 15:
		goto L23
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L2117
	default:
		goto L2118
	}
L2117:
	;
	v7524 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7530 = F_make_execsql_stmt(m, int32(276), v7524, int32(0), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7531 = m.ExcPending
	if v7531 != 0 {
		goto L66
	} else {
		goto L2121
	}
L2118:
	;
	if v7509 == int32(270) {
		goto L23
	} else {
		goto L2119
	}
L2119:
	;
	if v7509 == int32(91) {
		goto L23
	} else {
		goto L2120
	}
L2120:
	;
	goto L2117
L2121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7530
	v10317 = v241
	goto L5
L2122:
	;
	v7550 = F_palloc(m, int32(28))
	mBase = m.M
	v7551 = m.ExcPending
	if v7551 != 0 {
		goto L66
	} else {
		goto L2123
	}
L2123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7550))) = int32(17)
	v7554 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v7555 = int32(0)
	if v7554 < v7555 {
		v7598 = v7555
		goto L2125
	} else {
		goto L2126
	}
L2124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7550)+4)) = v7598
	v7603 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v7604 = *(*int32)(unsafe.Add(mBase, uint32(v7603)+520))
	v7606 = v7604 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7603)+520)) = v7606
	*(*int64)(unsafe.Add(mBase, uint32(v7550)+20)) = int64(0)
	v7610 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7550)+16)) = uint16(v7610)
	*(*int32)(unsafe.Add(mBase, uint32(v7550)+12)) = v7547
	*(*int32)(unsafe.Add(mBase, uint32(v7550)+8)) = v7606
	v7618 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	v7623 = v7618
	goto L2138
L2125:
	;
	goto L2124
L2126:
	;
	v7560 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7561 = *(*int32)(unsafe.Add(mBase, uint32(v7560)+60))
	if v7561 == int32(0) {
		v7598 = v7555
		goto L2125
	} else {
		goto L2127
	}
L2127:
	;
	v7564 = v7554 + v7561
	v7565 = *(*int32)(unsafe.Add(mBase, uint32(v7560)+188))
	if base.Ui32(v7565) <= base.Ui32(v7564) {
		goto L2129
	} else {
		goto L2130
	}
L2128:
	;
	v7575 = *(*int32)(unsafe.Add(mBase, uint32(v7560)+196))
	if v7574 == int32(0) {
		v7598 = v7575
		goto L2125
	} else {
		goto L2132
	}
L2129:
	;
	v7567 = *(*int32)(unsafe.Add(mBase, uint32(v7560)+192))
	v7574 = v7567
	goto L2128
L2130:
	;
	goto L2131
L2131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7560)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7560)+188)) = v7561
	v7572 = F_strchr(m, v7561, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7560)+192)) = v7572
	v7574 = v7572
	goto L2128
L2132:
	;
	if base.Ui32(v7564) <= base.Ui32(v7574) {
		v7598 = v7575
		goto L2125
	} else {
		goto L2133
	}
L2133:
	;
	v7579 = v7574
	v7581 = v7575
	goto L2134
L2134:
	;
	v7584 = int32(1)
	v7585 = v7581 + v7584
	*(*int32)(unsafe.Add(mBase, uint32(v7560)+196)) = v7585
	v7588 = v7579 + v7584
	*(*int32)(unsafe.Add(mBase, uint32(v7560)+188)) = v7588
	v7591 = F_strchr(m, v7588, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7560)+192)) = v7591
	if v7591 == int32(0) {
		v7598 = v7585
		goto L2125
	} else {
		goto L2136
	}
L2135:
	;
	v7598 = v7585
	goto L2125
L2136:
	;
	if base.Ui32(v7591) < base.Ui32(v7564) {
		v7579 = v7591
		v7581 = v7585
		goto L2134
	} else {
		goto L2137
	}
L2137:
	;
	goto L2135
L2138:
	;
	if v7623 != int32(332) {
		goto L2142
	} else {
		goto L2143
	}
L2139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7550
	v10317 = v241
	goto L5
L2140:
	;
	goto L2139
L2141:
	;
	v7673 = *(*int32)(unsafe.Add(mBase, uint32(v7550)+24))
	if v7673 != 0 {
		goto L21
	} else {
		goto L2151
	}
L2142:
	;
	if v7623 == int32(381) {
		goto L2141
	} else {
		goto L2145
	}
L2143:
	;
	goto L2144
L2144:
	;
	v7655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7550)+16)))
	if v7655 == int32(1) {
		goto L22
	} else {
		goto L2148
	}
L2145:
	;
	if v7623 == int32(59) {
		goto L2140
	} else {
		goto L2146
	}
L2146:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(209437))
	mBase = m.M
	v7654 = m.ExcPending
	if v7654 != 0 {
		goto L66
	} else {
		goto L2147
	}
L2147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2148:
	;
	v7658 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7550)+16)) = uint8(v7658)
	F_read_into_target(m, v7550+int32(20), v7550+int32(17), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7665 = m.ExcPending
	if v7665 != 0 {
		goto L66
	} else {
		goto L2149
	}
L2149:
	;
	v7670 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7671 = m.ExcPending
	if v7671 != 0 {
		goto L66
	} else {
		goto L2150
	}
L2150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = v7670
	v7623 = v7670
	goto L2138
L2151:
	;
	goto L2152
L2152:
	;
	v7703 = int32(1)
	v7712 = F_read_sql_construct(m, int32(44), int32(59), int32(332), int32(518809), int32(2), v7703, v7703, int32(0), v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7713 = m.ExcPending
	if v7713 != 0 {
		goto L66
	} else {
		goto L2154
	}
L2153:
	;
	v7623 = v7718
	goto L2138
L2154:
	;
	v7714 = *(*int32)(unsafe.Add(mBase, uint32(v7550)+24))
	v7715 = F_lappend(m, v7714, v7712)
	mBase = m.M
	v7716 = m.ExcPending
	if v7716 != 0 {
		goto L66
	} else {
		goto L2155
	}
L2155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7550)+24)) = v7715
	v7718 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v7718 == int32(44) {
		goto L2152
	} else {
		goto L2156
	}
L2156:
	;
	goto L2153
L2157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7723))) = int32(20)
	v7729 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v7730 = int32(0)
	if v7729 < v7730 {
		v7773 = v7730
		goto L2159
	} else {
		goto L2160
	}
L2158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+4)) = v7773
	v7778 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v7779 = *(*int32)(unsafe.Add(mBase, uint32(v7778)+520))
	v7781 = v7779 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7778)+520)) = v7781
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+8)) = v7781
	v7784 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v7785 = *(*int32)(unsafe.Add(mBase, uint32(v7784)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+16)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+12)) = v7785
	v7789 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v7790 = *(*int32)(unsafe.Add(mBase, uint32(v7789)+28))
	if v7790 == int32(0) {
		goto L2172
	} else {
		goto L2173
	}
L2159:
	;
	goto L2158
L2160:
	;
	v7735 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7736 = *(*int32)(unsafe.Add(mBase, uint32(v7735)+60))
	if v7736 == int32(0) {
		v7773 = v7730
		goto L2159
	} else {
		goto L2161
	}
L2161:
	;
	v7739 = v7729 + v7736
	v7740 = *(*int32)(unsafe.Add(mBase, uint32(v7735)+188))
	if base.Ui32(v7740) <= base.Ui32(v7739) {
		goto L2163
	} else {
		goto L2164
	}
L2162:
	;
	v7750 = *(*int32)(unsafe.Add(mBase, uint32(v7735)+196))
	if v7749 == int32(0) {
		v7773 = v7750
		goto L2159
	} else {
		goto L2166
	}
L2163:
	;
	v7742 = *(*int32)(unsafe.Add(mBase, uint32(v7735)+192))
	v7749 = v7742
	goto L2162
L2164:
	;
	goto L2165
L2165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7735)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7735)+188)) = v7736
	v7747 = F_strchr(m, v7736, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7735)+192)) = v7747
	v7749 = v7747
	goto L2162
L2166:
	;
	if base.Ui32(v7739) <= base.Ui32(v7749) {
		v7773 = v7750
		goto L2159
	} else {
		goto L2167
	}
L2167:
	;
	v7754 = v7749
	v7756 = v7750
	goto L2168
L2168:
	;
	v7759 = int32(1)
	v7760 = v7756 + v7759
	*(*int32)(unsafe.Add(mBase, uint32(v7735)+196)) = v7760
	v7763 = v7754 + v7759
	*(*int32)(unsafe.Add(mBase, uint32(v7735)+188)) = v7763
	v7766 = F_strchr(m, v7763, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7735)+192)) = v7766
	if v7766 == int32(0) {
		v7773 = v7760
		goto L2159
	} else {
		goto L2170
	}
L2169:
	;
	v7773 = v7760
	goto L2159
L2170:
	;
	if base.Ui32(v7766) < base.Ui32(v7739) {
		v7754 = v7766
		v7756 = v7760
		goto L2168
	} else {
		goto L2171
	}
L2171:
	;
	goto L2169
L2172:
	;
	v7793 = int32(2)
	v7798 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7799 = m.ExcPending
	if v7799 != 0 {
		goto L66
	} else {
		goto L2177
	}
L2173:
	;
	goto L2174
L2174:
	;
	v8034 = F_read_cursor_args(m, v7789, int32(59), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8035 = m.ExcPending
	if v8035 != 0 {
		goto L66
	} else {
		goto L2231
	}
L2175:
	;
	if v7902 != int32(321) {
		goto L12
	} else {
		goto L2215
	}
L2176:
	;
	v7893 = *(*int32)(unsafe.Add(mBase, uint32(v7723)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+16)) = v7893 | v7892
	v7900 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7901 = m.ExcPending
	if v7901 != 0 {
		goto L66
	} else {
		goto L2214
	}
L2177:
	;
	if v7798 == int32(369) {
		v7892 = v7793
		goto L2176
	} else {
		goto L2178
	}
L2178:
	;
	if v7798 != int32(342) {
		goto L2181
	} else {
		goto L2182
	}
L2179:
	;
	v7865 = int32(299445)
	v7868 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1416])))
	v7869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7864))))
	if v7869 == int32(0) {
		v7888 = v7868
		v7889 = v7869
		goto L2206
	} else {
		goto L2207
	}
L2180:
	;
	v7861 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v7861 == int32(0) {
		goto L12
	} else {
		goto L2204
	}
L2181:
	;
	if v7798 != int32(277) {
		v7902 = v7798
		goto L2175
	} else {
		goto L2184
	}
L2182:
	;
	goto L2183
L2183:
	;
	v7818 = int32(4)
	v7823 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7824 = m.ExcPending
	if v7824 != 0 {
		goto L66
	} else {
		goto L2190
	}
L2184:
	;
	v7806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v7806 != 0 {
		goto L12
	} else {
		goto L2185
	}
L2185:
	;
	v7807 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v7807 == int32(0) {
		goto L12
	} else {
		goto L2186
	}
L2186:
	;
	v7810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7807))))
	if v7810 != int32(110) {
		v7864 = v7807
		goto L2179
	} else {
		goto L2187
	}
L2187:
	;
	v7813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7807)+1)))
	if v7813 != int32(111) {
		goto L2180
	} else {
		goto L2188
	}
L2188:
	;
	v7816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7807)+2)))
	if v7816 != 0 {
		goto L2180
	} else {
		goto L2189
	}
L2189:
	;
	goto L2183
L2190:
	;
	if v7823 == int32(369) {
		v7892 = v7818
		goto L2176
	} else {
		goto L2191
	}
L2191:
	;
	if v7823 != int32(277) {
		v7902 = v7823
		goto L2175
	} else {
		goto L2192
	}
L2192:
	;
	v7829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v7829 != 0 {
		goto L12
	} else {
		goto L2193
	}
L2193:
	;
	v7830 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v7830 == int32(0) {
		goto L12
	} else {
		goto L2194
	}
L2194:
	;
	v7833 = int32(299445)
	v7836 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1416])))
	v7837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7830))))
	if v7837 == int32(0) {
		v7856 = v7836
		v7857 = v7837
		goto L2196
	} else {
		goto L2197
	}
L2195:
	;
	if v7857-v7856 == int32(0) {
		v7892 = v7818
		goto L2176
	} else {
		goto L2203
	}
L2196:
	;
	goto L2195
L2197:
	;
	if v7836 != v7837 {
		v7856 = v7836
		v7857 = v7837
		goto L2196
	} else {
		goto L2198
	}
L2198:
	;
	v7841 = v7830
	v7842 = v7833
	goto L2199
L2199:
	;
	v7845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7842)+1)))
	v7846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7841)+1)))
	if v7846 == int32(0) {
		v7856 = v7845
		v7857 = v7846
		goto L2196
	} else {
		goto L2201
	}
L2200:
	;
	v7856 = v7845
	v7857 = v7846
	goto L2196
L2201:
	;
	v7849 = int32(1)
	if v7845 == v7846 {
		v7841 = v7841 + v7849
		v7842 = v7842 + v7849
		goto L2199
	} else {
		goto L2202
	}
L2202:
	;
	goto L2200
L2203:
	;
	goto L12
L2204:
	;
	v7864 = v7861
	goto L2179
L2205:
	;
	if v7889-v7888 != 0 {
		goto L12
	} else {
		goto L2213
	}
L2206:
	;
	goto L2205
L2207:
	;
	if v7868 != v7869 {
		v7888 = v7868
		v7889 = v7869
		goto L2206
	} else {
		goto L2208
	}
L2208:
	;
	v7873 = v7864
	v7874 = v7865
	goto L2209
L2209:
	;
	v7877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7874)+1)))
	v7878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7873)+1)))
	if v7878 == int32(0) {
		v7888 = v7877
		v7889 = v7878
		goto L2206
	} else {
		goto L2211
	}
L2210:
	;
	v7888 = v7877
	v7889 = v7878
	goto L2206
L2211:
	;
	v7881 = int32(1)
	if v7877 == v7878 {
		v7873 = v7873 + v7881
		v7874 = v7874 + v7881
		goto L2209
	} else {
		goto L2212
	}
L2212:
	;
	goto L2210
L2213:
	;
	v7892 = v7793
	goto L2176
L2214:
	;
	v7902 = v7900
	goto L2175
L2215:
	;
	v7910 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7911 = m.ExcPending
	if v7911 != 0 {
		goto L66
	} else {
		goto L2216
	}
L2216:
	;
	if v7910 == int32(317) {
		goto L2217
	} else {
		goto L2218
	}
L2217:
	;
	v7916 = int32(0)
	v7919 = int32(1)
	v7928 = F_read_sql_construct(m, int32(381), int32(59), v7916, int32(538718), int32(2), v7919, v7919, v7916, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7929 = m.ExcPending
	if v7929 != 0 {
		goto L66
	} else {
		goto L2220
	}
L2218:
	;
	goto L2219
L2219:
	;
	F_plpgsql_push_back_token(m, v7910, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8011 = m.ExcPending
	if v8011 != 0 {
		goto L66
	} else {
		goto L2229
	}
L2220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+28)) = v7928
	v7931 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v7931 == int32(381) {
		goto L2221
	} else {
		goto L2222
	}
L2221:
	;
	goto L2224
L2222:
	;
	goto L2223
L2223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7723
	v10317 = v241
	goto L5
L2224:
	;
	v7960 = int32(0)
	v7963 = int32(1)
	v7972 = F_read_sql_construct(m, int32(44), int32(59), v7960, int32(538729), int32(2), v7963, v7963, v7960, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v7973 = m.ExcPending
	if v7973 != 0 {
		goto L66
	} else {
		goto L2226
	}
L2225:
	;
	goto L2223
L2226:
	;
	v7974 = *(*int32)(unsafe.Add(mBase, uint32(v7723)+32))
	v7975 = F_lappend(m, v7974, v7972)
	mBase = m.M
	v7976 = m.ExcPending
	if v7976 != 0 {
		goto L66
	} else {
		goto L2227
	}
L2227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+32)) = v7975
	v7978 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	if v7978 == int32(44) {
		goto L2224
	} else {
		goto L2228
	}
L2228:
	;
	goto L2225
L2229:
	;
	v8013 = int32(0)
	v8025 = F_read_sql_construct(m, int32(59), v8013, v8013, int32(538734), v8013, v8013, int32(1), v8013, v8013, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8026 = m.ExcPending
	if v8026 != 0 {
		goto L66
	} else {
		goto L2230
	}
L2230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+24)) = v8025
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7723
	v10317 = v241
	goto L5
L2231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7723)+20)) = v8034
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v7723
	v10317 = v241
	goto L5
L2232:
	;
	v8054 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8055 = m.ExcPending
	if v8055 != 0 {
		goto L66
	} else {
		goto L2233
	}
L2233:
	;
	if v8054 != int32(59) {
		goto L20
	} else {
		goto L2234
	}
L2234:
	;
	v8058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8040)+33)))
	if v8058 == int32(1) {
		goto L19
	} else {
		goto L2235
	}
L2235:
	;
	v8063 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v8064 = int32(0)
	if v8063 < v8064 {
		v8107 = v8064
		goto L2237
	} else {
		goto L2238
	}
L2236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8040)+4)) = v8107
	v8111 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	*(*int32)(unsafe.Add(mBase, uint32(v8040)+12)) = v8111
	v8115 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8116 = *(*int32)(unsafe.Add(mBase, uint32(v8115)+4))
	v8117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8040)+32)) = uint8(v8117)
	*(*int32)(unsafe.Add(mBase, uint32(v8040)+16)) = v8116
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8040
	v10317 = v241
	goto L5
L2237:
	;
	goto L2236
L2238:
	;
	v8069 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8070 = *(*int32)(unsafe.Add(mBase, uint32(v8069)+60))
	if v8070 == int32(0) {
		v8107 = v8064
		goto L2237
	} else {
		goto L2239
	}
L2239:
	;
	v8073 = v8063 + v8070
	v8074 = *(*int32)(unsafe.Add(mBase, uint32(v8069)+188))
	if base.Ui32(v8074) <= base.Ui32(v8073) {
		goto L2241
	} else {
		goto L2242
	}
L2240:
	;
	v8084 = *(*int32)(unsafe.Add(mBase, uint32(v8069)+196))
	if v8083 == int32(0) {
		v8107 = v8084
		goto L2237
	} else {
		goto L2244
	}
L2241:
	;
	v8076 = *(*int32)(unsafe.Add(mBase, uint32(v8069)+192))
	v8083 = v8076
	goto L2240
L2242:
	;
	goto L2243
L2243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8069)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8069)+188)) = v8070
	v8081 = F_strchr(m, v8070, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8069)+192)) = v8081
	v8083 = v8081
	goto L2240
L2244:
	;
	if base.Ui32(v8073) <= base.Ui32(v8083) {
		v8107 = v8084
		goto L2237
	} else {
		goto L2245
	}
L2245:
	;
	v8088 = v8083
	v8090 = v8084
	goto L2246
L2246:
	;
	v8093 = int32(1)
	v8094 = v8090 + v8093
	*(*int32)(unsafe.Add(mBase, uint32(v8069)+196)) = v8094
	v8097 = v8088 + v8093
	*(*int32)(unsafe.Add(mBase, uint32(v8069)+188)) = v8097
	v8100 = F_strchr(m, v8097, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8069)+192)) = v8100
	if v8100 == int32(0) {
		v8107 = v8094
		goto L2237
	} else {
		goto L2248
	}
L2247:
	;
	v8107 = v8094
	goto L2237
L2248:
	;
	if base.Ui32(v8100) < base.Ui32(v8073) {
		v8088 = v8100
		v8090 = v8094
		goto L2246
	} else {
		goto L2249
	}
L2249:
	;
	goto L2247
L2250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8123)+4)) = v8170
	v8176 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8177 = *(*int32)(unsafe.Add(mBase, uint32(v8176)+4))
	v8178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8123)+32)) = uint8(v8178)
	*(*int32)(unsafe.Add(mBase, uint32(v8123)+16)) = v8177
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8123
	v10317 = v241
	goto L5
L2251:
	;
	goto L2250
L2252:
	;
	v8132 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8133 = *(*int32)(unsafe.Add(mBase, uint32(v8132)+60))
	if v8133 == int32(0) {
		v8170 = v8127
		goto L2251
	} else {
		goto L2253
	}
L2253:
	;
	v8136 = v8126 + v8133
	v8137 = *(*int32)(unsafe.Add(mBase, uint32(v8132)+188))
	if base.Ui32(v8137) <= base.Ui32(v8136) {
		goto L2255
	} else {
		goto L2256
	}
L2254:
	;
	v8147 = *(*int32)(unsafe.Add(mBase, uint32(v8132)+196))
	if v8146 == int32(0) {
		v8170 = v8147
		goto L2251
	} else {
		goto L2258
	}
L2255:
	;
	v8139 = *(*int32)(unsafe.Add(mBase, uint32(v8132)+192))
	v8146 = v8139
	goto L2254
L2256:
	;
	goto L2257
L2257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8132)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8132)+188)) = v8133
	v8144 = F_strchr(m, v8133, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8132)+192)) = v8144
	v8146 = v8144
	goto L2254
L2258:
	;
	if base.Ui32(v8136) <= base.Ui32(v8146) {
		v8170 = v8147
		goto L2251
	} else {
		goto L2259
	}
L2259:
	;
	v8151 = v8146
	v8153 = v8147
	goto L2260
L2260:
	;
	v8156 = int32(1)
	v8157 = v8153 + v8156
	*(*int32)(unsafe.Add(mBase, uint32(v8132)+196)) = v8157
	v8160 = v8151 + v8156
	*(*int32)(unsafe.Add(mBase, uint32(v8132)+188)) = v8160
	v8163 = F_strchr(m, v8160, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8132)+192)) = v8163
	if v8163 == int32(0) {
		v8170 = v8157
		goto L2251
	} else {
		goto L2262
	}
L2261:
	;
	v8170 = v8157
	goto L2251
L2262:
	;
	if base.Ui32(v8163) < base.Ui32(v8136) {
		v8151 = v8163
		v8153 = v8157
		goto L2260
	} else {
		goto L2263
	}
L2263:
	;
	goto L2261
L2264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8185))) = int32(21)
	v8190 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v8191 = *(*int32)(unsafe.Add(mBase, uint32(v8190)+520))
	v8193 = v8191 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8190)+520)) = v8193
	v8195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8185)+33)) = uint8(v8195)
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+28)) = v8195
	*(*int64)(unsafe.Add(mBase, uint32(v8185)+20)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+8)) = v8193
	v8206 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8207 = m.ExcPending
	if v8207 != 0 {
		goto L66
	} else {
		goto L2286
	}
L2265:
	;
	v8572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))))
	if v8572 != int32(1) {
		goto L2381
	} else {
		goto L2382
	}
L2266:
	;
	v8569 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = uint8(v8569)
	goto L2265
L2267:
	;
	F_plpgsql_push_back_token(m, v8206, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8549 = m.ExcPending
	if v8549 != 0 {
		goto L66
	} else {
		goto L2379
	}
L2268:
	;
	F_plpgsql_push_back_token(m, int32(277), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8541 = m.ExcPending
	if v8541 != 0 {
		goto L66
	} else {
		goto L2378
	}
L2269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+20)) = int32(1)
	F_complete_direction(m, v8185, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8533 = m.ExcPending
	if v8533 != 0 {
		goto L66
	} else {
		goto L2377
	}
L2270:
	;
	v8494 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v8494 == int32(0) {
		goto L2268
	} else {
		goto L2367
	}
L2271:
	;
	F_complete_direction(m, v8185, v27+int32(4784), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8493 = m.ExcPending
	if v8493 != 0 {
		goto L66
	} else {
		goto L2366
	}
L2272:
	;
	v8459 = int32(415477)
	v8462 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1417])))
	v8463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8425))))
	if v8463 == int32(0) {
		v8482 = v8462
		v8483 = v8463
		goto L2358
	} else {
		goto L2359
	}
L2273:
	;
	v8455 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8185)+33)) = uint8(v8455)
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+24)) = int32(2147483647)
	goto L2265
L2274:
	;
	v8425 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v8425 == int32(0) {
		goto L2268
	} else {
		goto L2347
	}
L2275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+20)) = int32(3)
	v8411 = int32(0)
	v8414 = int32(1)
	v8422 = F_read_sql_construct(m, int32(324), int32(329), v8411, int32(522564), int32(2), v8414, v8414, v8411, v8411, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8423 = m.ExcPending
	if v8423 != 0 {
		goto L66
	} else {
		goto L2346
	}
L2276:
	;
	v8380 = int32(337840)
	v8383 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1418])))
	v8384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8217))))
	if v8384 == int32(0) {
		v8403 = v8383
		v8404 = v8384
		goto L2338
	} else {
		goto L2339
	}
L2277:
	;
	v8362 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+20)) = v8362
	v8366 = int32(0)
	v8369 = int32(1)
	v8377 = F_read_sql_construct(m, int32(324), int32(329), v8366, int32(522564), v8362, v8369, v8369, v8366, v8366, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8378 = m.ExcPending
	if v8378 != 0 {
		goto L66
	} else {
		goto L2336
	}
L2278:
	;
	v8335 = int32(343304)
	v8338 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1419])))
	v8339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8217))))
	if v8339 == int32(0) {
		v8358 = v8338
		v8359 = v8339
		goto L2328
	} else {
		goto L2329
	}
L2279:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8185)+20)) = int64(-4294967294)
	goto L2265
L2280:
	;
	v8306 = int32(77227)
	v8309 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1420])))
	v8310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8217))))
	if v8310 == int32(0) {
		v8329 = v8309
		v8330 = v8310
		goto L2319
	} else {
		goto L2320
	}
L2281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+20)) = int32(2)
	goto L2265
L2282:
	;
	v8277 = int32(67211)
	v8280 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1421])))
	v8281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8217))))
	if v8281 == int32(0) {
		v8300 = v8280
		v8301 = v8281
		goto L2310
	} else {
		goto L2311
	}
L2283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+20)) = int32(1)
	goto L2265
L2284:
	;
	v8216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1363]))))
	if v8216 != 0 {
		goto L2268
	} else {
		goto L2289
	}
L2285:
	;
	if v8206 != 0 {
		goto L2267
	} else {
		goto L2287
	}
L2286:
	;
	switch v8206 - int32(277) {
	case 0:
		goto L2284
	case 1, 2, 4, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 44, 45, 48, 49, 50, 51, 53, 54, 55, 56, 58, 59, 60, 61, 62, 63, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 81, 82:
		goto L2267
	case 3:
		goto L2277
	case 5:
		goto L2273
	case 9:
		goto L2269
	case 43:
		goto L2281
	case 46:
		goto L2271
	case 47, 52:
		goto L2266
	case 57:
		goto L2279
	case 64:
		goto L2265
	case 80:
		goto L2283
	case 83:
		goto L2275
	default:
		goto L2285
	}
L2287:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(246872))
	mBase = m.M
	v8215 = m.ExcPending
	if v8215 != 0 {
		goto L66
	} else {
		goto L2288
	}
L2288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2289:
	;
	v8217 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1370])))
	if v8217 == int32(0) {
		goto L2268
	} else {
		goto L2290
	}
L2290:
	;
	v8220 = int32(62940)
	v8223 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1398])))
	v8224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8217))))
	if v8224 == int32(0) {
		v8243 = v8223
		v8244 = v8224
		goto L2292
	} else {
		goto L2293
	}
L2291:
	;
	if v8244-v8243 == int32(0) {
		goto L2265
	} else {
		goto L2299
	}
L2292:
	;
	goto L2291
L2293:
	;
	if v8223 != v8224 {
		v8243 = v8223
		v8244 = v8224
		goto L2292
	} else {
		goto L2294
	}
L2294:
	;
	v8228 = v8217
	v8229 = v8220
	goto L2295
L2295:
	;
	v8232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8229)+1)))
	v8233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8228)+1)))
	if v8233 == int32(0) {
		v8243 = v8232
		v8244 = v8233
		goto L2292
	} else {
		goto L2297
	}
L2296:
	;
	v8243 = v8232
	v8244 = v8233
	goto L2292
L2297:
	;
	v8236 = int32(1)
	if v8232 == v8233 {
		v8228 = v8228 + v8236
		v8229 = v8229 + v8236
		goto L2295
	} else {
		goto L2298
	}
L2298:
	;
	goto L2296
L2299:
	;
	v8248 = int32(210224)
	v8251 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1422])))
	v8252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8217))))
	if v8252 == int32(0) {
		v8271 = v8251
		v8272 = v8252
		goto L2301
	} else {
		goto L2302
	}
L2300:
	;
	if v8272-v8271 != 0 {
		goto L2282
	} else {
		goto L2308
	}
L2301:
	;
	goto L2300
L2302:
	;
	if v8251 != v8252 {
		v8271 = v8251
		v8272 = v8252
		goto L2301
	} else {
		goto L2303
	}
L2303:
	;
	v8256 = v8217
	v8257 = v8248
	goto L2304
L2304:
	;
	v8260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8257)+1)))
	v8261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8256)+1)))
	if v8261 == int32(0) {
		v8271 = v8260
		v8272 = v8261
		goto L2301
	} else {
		goto L2306
	}
L2305:
	;
	v8271 = v8260
	v8272 = v8261
	goto L2301
L2306:
	;
	v8264 = int32(1)
	if v8260 == v8261 {
		v8256 = v8256 + v8264
		v8257 = v8257 + v8264
		goto L2304
	} else {
		goto L2307
	}
L2307:
	;
	goto L2305
L2308:
	;
	goto L2283
L2309:
	;
	if v8301-v8300 != 0 {
		goto L2280
	} else {
		goto L2317
	}
L2310:
	;
	goto L2309
L2311:
	;
	if v8280 != v8281 {
		v8300 = v8280
		v8301 = v8281
		goto L2310
	} else {
		goto L2312
	}
L2312:
	;
	v8285 = v8217
	v8286 = v8277
	goto L2313
L2313:
	;
	v8289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8286)+1)))
	v8290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8285)+1)))
	if v8290 == int32(0) {
		v8300 = v8289
		v8301 = v8290
		goto L2310
	} else {
		goto L2315
	}
L2314:
	;
	v8300 = v8289
	v8301 = v8290
	goto L2310
L2315:
	;
	v8293 = int32(1)
	if v8289 == v8290 {
		v8285 = v8285 + v8293
		v8286 = v8286 + v8293
		goto L2313
	} else {
		goto L2316
	}
L2316:
	;
	goto L2314
L2317:
	;
	goto L2281
L2318:
	;
	if v8330-v8329 != 0 {
		goto L2278
	} else {
		goto L2326
	}
L2319:
	;
	goto L2318
L2320:
	;
	if v8309 != v8310 {
		v8329 = v8309
		v8330 = v8310
		goto L2319
	} else {
		goto L2321
	}
L2321:
	;
	v8314 = v8217
	v8315 = v8306
	goto L2322
L2322:
	;
	v8318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8315)+1)))
	v8319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8314)+1)))
	if v8319 == int32(0) {
		v8329 = v8318
		v8330 = v8319
		goto L2319
	} else {
		goto L2324
	}
L2323:
	;
	v8329 = v8318
	v8330 = v8319
	goto L2319
L2324:
	;
	v8322 = int32(1)
	if v8318 == v8319 {
		v8314 = v8314 + v8322
		v8315 = v8315 + v8322
		goto L2322
	} else {
		goto L2325
	}
L2325:
	;
	goto L2323
L2326:
	;
	goto L2279
L2327:
	;
	if v8359-v8358 != 0 {
		goto L2276
	} else {
		goto L2335
	}
L2328:
	;
	goto L2327
L2329:
	;
	if v8338 != v8339 {
		v8358 = v8338
		v8359 = v8339
		goto L2328
	} else {
		goto L2330
	}
L2330:
	;
	v8343 = v8217
	v8344 = v8335
	goto L2331
L2331:
	;
	v8347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8344)+1)))
	v8348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8343)+1)))
	if v8348 == int32(0) {
		v8358 = v8347
		v8359 = v8348
		goto L2328
	} else {
		goto L2333
	}
L2332:
	;
	v8358 = v8347
	v8359 = v8348
	goto L2328
L2333:
	;
	v8351 = int32(1)
	if v8347 == v8348 {
		v8343 = v8343 + v8351
		v8344 = v8344 + v8351
		goto L2331
	} else {
		goto L2334
	}
L2334:
	;
	goto L2332
L2335:
	;
	goto L2277
L2336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+28)) = v8377
	goto L2266
L2337:
	;
	if v8404-v8403 != 0 {
		goto L2274
	} else {
		goto L2345
	}
L2338:
	;
	goto L2337
L2339:
	;
	if v8383 != v8384 {
		v8403 = v8383
		v8404 = v8384
		goto L2338
	} else {
		goto L2340
	}
L2340:
	;
	v8388 = v8217
	v8389 = v8380
	goto L2341
L2341:
	;
	v8392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8389)+1)))
	v8393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8388)+1)))
	if v8393 == int32(0) {
		v8403 = v8392
		v8404 = v8393
		goto L2338
	} else {
		goto L2343
	}
L2342:
	;
	v8403 = v8392
	v8404 = v8393
	goto L2338
L2343:
	;
	v8396 = int32(1)
	if v8392 == v8393 {
		v8388 = v8388 + v8396
		v8389 = v8389 + v8396
		goto L2341
	} else {
		goto L2344
	}
L2344:
	;
	goto L2342
L2345:
	;
	goto L2275
L2346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+28)) = v8422
	goto L2266
L2347:
	;
	v8428 = int32(300907)
	v8431 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1245])))
	v8432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8425))))
	if v8432 == int32(0) {
		v8451 = v8431
		v8452 = v8432
		goto L2349
	} else {
		goto L2350
	}
L2348:
	;
	if v8452-v8451 != 0 {
		goto L2272
	} else {
		goto L2356
	}
L2349:
	;
	goto L2348
L2350:
	;
	if v8431 != v8432 {
		v8451 = v8431
		v8452 = v8432
		goto L2349
	} else {
		goto L2351
	}
L2351:
	;
	v8436 = v8425
	v8437 = v8428
	goto L2352
L2352:
	;
	v8440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8437)+1)))
	v8441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8436)+1)))
	if v8441 == int32(0) {
		v8451 = v8440
		v8452 = v8441
		goto L2349
	} else {
		goto L2354
	}
L2353:
	;
	v8451 = v8440
	v8452 = v8441
	goto L2349
L2354:
	;
	v8444 = int32(1)
	if v8440 == v8441 {
		v8436 = v8436 + v8444
		v8437 = v8437 + v8444
		goto L2352
	} else {
		goto L2355
	}
L2355:
	;
	goto L2353
L2356:
	;
	goto L2273
L2357:
	;
	if v8483-v8482 != 0 {
		goto L2270
	} else {
		goto L2365
	}
L2358:
	;
	goto L2357
L2359:
	;
	if v8462 != v8463 {
		v8482 = v8462
		v8483 = v8463
		goto L2358
	} else {
		goto L2360
	}
L2360:
	;
	v8467 = v8425
	v8468 = v8459
	goto L2361
L2361:
	;
	v8471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8468)+1)))
	v8472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8467)+1)))
	if v8472 == int32(0) {
		v8482 = v8471
		v8483 = v8472
		goto L2358
	} else {
		goto L2363
	}
L2362:
	;
	v8482 = v8471
	v8483 = v8472
	goto L2358
L2363:
	;
	v8475 = int32(1)
	if v8471 == v8472 {
		v8467 = v8467 + v8475
		v8468 = v8468 + v8475
		goto L2361
	} else {
		goto L2364
	}
L2364:
	;
	goto L2362
L2365:
	;
	goto L2271
L2366:
	;
	goto L2265
L2367:
	;
	v8497 = int32(415567)
	v8500 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1423])))
	v8501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8494))))
	if v8501 == int32(0) {
		v8520 = v8500
		v8521 = v8501
		goto L2369
	} else {
		goto L2370
	}
L2368:
	;
	if v8521-v8520 != 0 {
		goto L2268
	} else {
		goto L2376
	}
L2369:
	;
	goto L2368
L2370:
	;
	if v8500 != v8501 {
		v8520 = v8500
		v8521 = v8501
		goto L2369
	} else {
		goto L2371
	}
L2371:
	;
	v8505 = v8494
	v8506 = v8497
	goto L2372
L2372:
	;
	v8509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8506)+1)))
	v8510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8505)+1)))
	if v8510 == int32(0) {
		v8520 = v8509
		v8521 = v8510
		goto L2369
	} else {
		goto L2374
	}
L2373:
	;
	v8520 = v8509
	v8521 = v8510
	goto L2369
L2374:
	;
	v8513 = int32(1)
	if v8509 == v8510 {
		v8505 = v8505 + v8513
		v8506 = v8506 + v8513
		goto L2372
	} else {
		goto L2375
	}
L2375:
	;
	goto L2373
L2376:
	;
	goto L2269
L2377:
	;
	goto L2265
L2378:
	;
	v8542 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = uint8(v8542)
	goto L2265
L2379:
	;
	v8552 = int32(0)
	v8555 = int32(1)
	v8563 = F_read_sql_construct(m, int32(324), int32(329), v8552, int32(522564), int32(2), v8555, v8555, v8552, v8552, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8564 = m.ExcPending
	if v8564 != 0 {
		goto L66
	} else {
		goto L2380
	}
L2380:
	;
	v8565 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8185)+33)) = uint8(v8565)
	*(*int32)(unsafe.Add(mBase, uint32(v8185)+28)) = v8563
	goto L2266
L2381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8185
	v10317 = v241
	goto L5
L2382:
	;
	v8579 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v8580 = m.ExcPending
	if v8580 != 0 {
		goto L66
	} else {
		goto L2384
	}
L2383:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(522555))
	mBase = m.M
	v8588 = m.ExcPending
	if v8588 != 0 {
		goto L66
	} else {
		goto L2385
	}
L2384:
	;
	switch v8579 - int32(324) {
	case 0, 5:
		goto L2381
	default:
		goto L2383
	}
L2385:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8591))) = int32(22)
	v8597 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v8598 = int32(0)
	if v8597 < v8598 {
		v8641 = v8598
		goto L2388
	} else {
		goto L2389
	}
L2387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8591)+4)) = v8641
	v8646 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v8647 = *(*int32)(unsafe.Add(mBase, uint32(v8646)+520))
	v8649 = v8647 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8646)+520)) = v8649
	*(*int32)(unsafe.Add(mBase, uint32(v8591)+8)) = v8649
	v8654 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v8655 = *(*int32)(unsafe.Add(mBase, uint32(v8654)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8591)+12)) = v8655
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8591
	v10317 = v241
	goto L5
L2388:
	;
	goto L2387
L2389:
	;
	v8603 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8604 = *(*int32)(unsafe.Add(mBase, uint32(v8603)+60))
	if v8604 == int32(0) {
		v8641 = v8598
		goto L2388
	} else {
		goto L2390
	}
L2390:
	;
	v8607 = v8597 + v8604
	v8608 = *(*int32)(unsafe.Add(mBase, uint32(v8603)+188))
	if base.Ui32(v8608) <= base.Ui32(v8607) {
		goto L2392
	} else {
		goto L2393
	}
L2391:
	;
	v8618 = *(*int32)(unsafe.Add(mBase, uint32(v8603)+196))
	if v8617 == int32(0) {
		v8641 = v8618
		goto L2388
	} else {
		goto L2395
	}
L2392:
	;
	v8610 = *(*int32)(unsafe.Add(mBase, uint32(v8603)+192))
	v8617 = v8610
	goto L2391
L2393:
	;
	goto L2394
L2394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8603)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8603)+188)) = v8604
	v8615 = F_strchr(m, v8604, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8603)+192)) = v8615
	v8617 = v8615
	goto L2391
L2395:
	;
	if base.Ui32(v8607) <= base.Ui32(v8617) {
		v8641 = v8618
		goto L2388
	} else {
		goto L2396
	}
L2396:
	;
	v8622 = v8617
	v8624 = v8618
	goto L2397
L2397:
	;
	v8627 = int32(1)
	v8628 = v8624 + v8627
	*(*int32)(unsafe.Add(mBase, uint32(v8603)+196)) = v8628
	v8631 = v8622 + v8627
	*(*int32)(unsafe.Add(mBase, uint32(v8603)+188)) = v8631
	v8634 = F_strchr(m, v8631, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8603)+192)) = v8634
	if v8634 == int32(0) {
		v8641 = v8628
		goto L2388
	} else {
		goto L2399
	}
L2398:
	;
	v8641 = v8628
	goto L2388
L2399:
	;
	if base.Ui32(v8634) < base.Ui32(v8607) {
		v8622 = v8634
		v8624 = v8628
		goto L2397
	} else {
		goto L2400
	}
L2400:
	;
	goto L2398
L2401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8661))) = int32(25)
	v8667 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v8668 = int32(0)
	if v8667 < v8668 {
		v8711 = v8668
		goto L2403
	} else {
		goto L2404
	}
L2402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8661)+4)) = v8711
	v8716 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v8717 = *(*int32)(unsafe.Add(mBase, uint32(v8716)+520))
	v8719 = v8717 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8716)+520)) = v8719
	*(*int32)(unsafe.Add(mBase, uint32(v8661)+8)) = v8719
	v8724 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8661)+12)) = uint8(base.B2i32(v8724 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8661
	v10317 = v241
	goto L5
L2403:
	;
	goto L2402
L2404:
	;
	v8673 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8674 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+60))
	if v8674 == int32(0) {
		v8711 = v8668
		goto L2403
	} else {
		goto L2405
	}
L2405:
	;
	v8677 = v8667 + v8674
	v8678 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+188))
	if base.Ui32(v8678) <= base.Ui32(v8677) {
		goto L2407
	} else {
		goto L2408
	}
L2406:
	;
	v8688 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+196))
	if v8687 == int32(0) {
		v8711 = v8688
		goto L2403
	} else {
		goto L2410
	}
L2407:
	;
	v8680 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+192))
	v8687 = v8680
	goto L2406
L2408:
	;
	goto L2409
L2409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8673)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8673)+188)) = v8674
	v8685 = F_strchr(m, v8674, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8673)+192)) = v8685
	v8687 = v8685
	goto L2406
L2410:
	;
	if base.Ui32(v8677) <= base.Ui32(v8687) {
		v8711 = v8688
		goto L2403
	} else {
		goto L2411
	}
L2411:
	;
	v8692 = v8687
	v8694 = v8688
	goto L2412
L2412:
	;
	v8697 = int32(1)
	v8698 = v8694 + v8697
	*(*int32)(unsafe.Add(mBase, uint32(v8673)+196)) = v8698
	v8701 = v8692 + v8697
	*(*int32)(unsafe.Add(mBase, uint32(v8673)+188)) = v8701
	v8704 = F_strchr(m, v8701, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8673)+192)) = v8704
	if v8704 == int32(0) {
		v8711 = v8698
		goto L2403
	} else {
		goto L2414
	}
L2413:
	;
	v8711 = v8698
	goto L2403
L2414:
	;
	if base.Ui32(v8704) < base.Ui32(v8677) {
		v8692 = v8704
		v8694 = v8698
		goto L2412
	} else {
		goto L2415
	}
L2415:
	;
	goto L2413
L2416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8730))) = int32(26)
	v8736 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v8737 = int32(0)
	if v8736 < v8737 {
		v8780 = v8737
		goto L2418
	} else {
		goto L2419
	}
L2417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8730)+4)) = v8780
	v8785 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v8786 = *(*int32)(unsafe.Add(mBase, uint32(v8785)+520))
	v8788 = v8786 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8785)+520)) = v8788
	*(*int32)(unsafe.Add(mBase, uint32(v8730)+8)) = v8788
	v8793 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	*(*uint8)(unsafe.Add(mBase, uint32(v8730)+12)) = uint8(base.B2i32(v8793 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8730
	v10317 = v241
	goto L5
L2418:
	;
	goto L2417
L2419:
	;
	v8742 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8743 = *(*int32)(unsafe.Add(mBase, uint32(v8742)+60))
	if v8743 == int32(0) {
		v8780 = v8737
		goto L2418
	} else {
		goto L2420
	}
L2420:
	;
	v8746 = v8736 + v8743
	v8747 = *(*int32)(unsafe.Add(mBase, uint32(v8742)+188))
	if base.Ui32(v8747) <= base.Ui32(v8746) {
		goto L2422
	} else {
		goto L2423
	}
L2421:
	;
	v8757 = *(*int32)(unsafe.Add(mBase, uint32(v8742)+196))
	if v8756 == int32(0) {
		v8780 = v8757
		goto L2418
	} else {
		goto L2425
	}
L2422:
	;
	v8749 = *(*int32)(unsafe.Add(mBase, uint32(v8742)+192))
	v8756 = v8749
	goto L2421
L2423:
	;
	goto L2424
L2424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8742)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8742)+188)) = v8743
	v8754 = F_strchr(m, v8743, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8742)+192)) = v8754
	v8756 = v8754
	goto L2421
L2425:
	;
	if base.Ui32(v8746) <= base.Ui32(v8756) {
		v8780 = v8757
		goto L2418
	} else {
		goto L2426
	}
L2426:
	;
	v8761 = v8756
	v8763 = v8757
	goto L2427
L2427:
	;
	v8766 = int32(1)
	v8767 = v8763 + v8766
	*(*int32)(unsafe.Add(mBase, uint32(v8742)+196)) = v8767
	v8770 = v8761 + v8766
	*(*int32)(unsafe.Add(mBase, uint32(v8742)+188)) = v8770
	v8773 = F_strchr(m, v8770, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8742)+192)) = v8773
	if v8773 == int32(0) {
		v8780 = v8767
		goto L2418
	} else {
		goto L2429
	}
L2428:
	;
	v8780 = v8767
	goto L2418
L2429:
	;
	if base.Ui32(v8773) < base.Ui32(v8746) {
		v8761 = v8773
		v8763 = v8767
		goto L2427
	} else {
		goto L2430
	}
L2430:
	;
	goto L2428
L2431:
	;
	v8806 = F_plpgsql_peek(m, l1)
	mBase = m.M
	v8807 = m.ExcPending
	if v8807 != 0 {
		goto L66
	} else {
		goto L2432
	}
L2432:
	;
	if v8806 == int32(91) {
		goto L18
	} else {
		goto L2433
	}
L2433:
	;
	v8810 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v8811 = *(*int32)(unsafe.Add(mBase, uint32(v8810)+24))
	v8812 = *(*int32)(unsafe.Add(mBase, uint32(v8811)+4))
	if v8812 != int32(1790) {
		goto L17
	} else {
		goto L2434
	}
L2434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8810
	v10317 = v241
	goto L5
L2435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2436:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2437:
	;
	v8872 = F_palloc(m, int32(12))
	mBase = m.M
	v8873 = m.ExcPending
	if v8873 != 0 {
		goto L66
	} else {
		goto L2451
	}
L2438:
	;
	goto L2437
L2439:
	;
	v8830 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8831 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+60))
	if v8831 == int32(0) {
		v8868 = v8825
		goto L2438
	} else {
		goto L2440
	}
L2440:
	;
	v8834 = v8824 + v8831
	v8835 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+188))
	if base.Ui32(v8835) <= base.Ui32(v8834) {
		goto L2442
	} else {
		goto L2443
	}
L2441:
	;
	v8845 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+196))
	if v8844 == int32(0) {
		v8868 = v8845
		goto L2438
	} else {
		goto L2445
	}
L2442:
	;
	v8837 = *(*int32)(unsafe.Add(mBase, uint32(v8830)+192))
	v8844 = v8837
	goto L2441
L2443:
	;
	goto L2444
L2444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+188)) = v8831
	v8842 = F_strchr(m, v8831, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8830)+192)) = v8842
	v8844 = v8842
	goto L2441
L2445:
	;
	if base.Ui32(v8834) <= base.Ui32(v8844) {
		v8868 = v8845
		goto L2438
	} else {
		goto L2446
	}
L2446:
	;
	v8849 = v8844
	v8851 = v8845
	goto L2447
L2447:
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
		goto L2438
	} else {
		goto L2449
	}
L2448:
	;
	v8868 = v8855
	goto L2438
L2449:
	;
	if base.Ui32(v8861) < base.Ui32(v8834) {
		v8849 = v8861
		v8851 = v8855
		goto L2447
	} else {
		goto L2450
	}
L2450:
	;
	goto L2448
L2451:
	;
	v8875 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v8876 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8875)+525)) = uint8(v8876)
	v8881 = *(*int32)(unsafe.Add(mBase, uint32(v8875)+44))
	v8883 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8881, int32(0))
	mBase = m.M
	v8884 = m.ExcPending
	if v8884 != 0 {
		goto L66
	} else {
		goto L2452
	}
L2452:
	;
	v8886 = F_plpgsql_build_variable(m, int32(346115), v8868, v8883, int32(1))
	mBase = m.M
	v8887 = m.ExcPending
	if v8887 != 0 {
		goto L66
	} else {
		goto L2453
	}
L2453:
	;
	v8888 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8886)+16)) = uint8(v8888)
	v8890 = *(*int32)(unsafe.Add(mBase, uint32(v8886)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8872))) = v8890
	v8896 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v8897 = *(*int32)(unsafe.Add(mBase, uint32(v8896)+44))
	v8899 = F_plpgsql_build_datatype(m, int32(25), int32(-1), v8897, int32(0))
	mBase = m.M
	v8900 = m.ExcPending
	if v8900 != 0 {
		goto L66
	} else {
		goto L2454
	}
L2454:
	;
	v8902 = F_plpgsql_build_variable(m, int32(283458), v8868, v8899, int32(1))
	mBase = m.M
	v8903 = m.ExcPending
	if v8903 != 0 {
		goto L66
	} else {
		goto L2455
	}
L2455:
	;
	v8904 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8902)+16)) = uint8(v8904)
	v8906 = *(*int32)(unsafe.Add(mBase, uint32(v8902)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8872)+4)) = v8906
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8872
	v10317 = v241
	goto L5
L2456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8919
	v10317 = v241
	goto L5
L2457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8928
	v10317 = v241
	goto L5
L2458:
	;
	v8936 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v8937 = int32(0)
	if v8936 < v8937 {
		v8980 = v8937
		goto L2460
	} else {
		goto L2461
	}
L2459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8932))) = v8980
	v8986 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8932)+4)) = v8986
	v8988 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v8932)+8)) = v8988
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v8932
	v10317 = v241
	goto L5
L2460:
	;
	goto L2459
L2461:
	;
	v8942 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8943 = *(*int32)(unsafe.Add(mBase, uint32(v8942)+60))
	if v8943 == int32(0) {
		v8980 = v8937
		goto L2460
	} else {
		goto L2462
	}
L2462:
	;
	v8946 = v8936 + v8943
	v8947 = *(*int32)(unsafe.Add(mBase, uint32(v8942)+188))
	if base.Ui32(v8947) <= base.Ui32(v8946) {
		goto L2464
	} else {
		goto L2465
	}
L2463:
	;
	v8957 = *(*int32)(unsafe.Add(mBase, uint32(v8942)+196))
	if v8956 == int32(0) {
		v8980 = v8957
		goto L2460
	} else {
		goto L2467
	}
L2464:
	;
	v8949 = *(*int32)(unsafe.Add(mBase, uint32(v8942)+192))
	v8956 = v8949
	goto L2463
L2465:
	;
	goto L2466
L2466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8942)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8942)+188)) = v8943
	v8954 = F_strchr(m, v8943, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8942)+192)) = v8954
	v8956 = v8954
	goto L2463
L2467:
	;
	if base.Ui32(v8946) <= base.Ui32(v8956) {
		v8980 = v8957
		goto L2460
	} else {
		goto L2468
	}
L2468:
	;
	v8961 = v8956
	v8963 = v8957
	goto L2469
L2469:
	;
	v8966 = int32(1)
	v8967 = v8963 + v8966
	*(*int32)(unsafe.Add(mBase, uint32(v8942)+196)) = v8967
	v8970 = v8961 + v8966
	*(*int32)(unsafe.Add(mBase, uint32(v8942)+188)) = v8970
	v8973 = F_strchr(m, v8970, int32(10))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8942)+192)) = v8973
	if v8973 == int32(0) {
		v8980 = v8967
		goto L2460
	} else {
		goto L2471
	}
L2470:
	;
	v8980 = v8967
	goto L2460
L2471:
	;
	if base.Ui32(v8973) < base.Ui32(v8946) {
		v8961 = v8973
		v8963 = v8967
		goto L2469
	} else {
		goto L2472
	}
L2472:
	;
	goto L2470
L2473:
	;
	v9018 = *(*int32)(unsafe.Add(mBase, uint32(v8997)+8))
	if v9018 != 0 {
		v8997 = v9018
		goto L2473
	} else {
		goto L2475
	}
L2474:
	;
	v9019 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v8997)+8)) = v9019
	v9021 = *(*int32)(unsafe.Add(mBase, uint32(v8992)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9021
	v10317 = v241
	goto L5
L2475:
	;
	goto L2474
L2476:
	;
	if v9050-v9049 != 0 {
		goto L2484
	} else {
		goto L2485
	}
L2477:
	;
	goto L2476
L2478:
	;
	if v9029 != v9030 {
		v9049 = v9029
		v9050 = v9030
		goto L2477
	} else {
		goto L2479
	}
L2479:
	;
	v9034 = v9025
	v9035 = v9026
	goto L2480
L2480:
	;
	v9038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9035)+1)))
	v9039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9034)+1)))
	if v9039 == int32(0) {
		v9049 = v9038
		v9050 = v9039
		goto L2477
	} else {
		goto L2482
	}
L2481:
	;
	v9049 = v9038
	v9050 = v9039
	goto L2477
L2482:
	;
	v9042 = int32(1)
	if v9038 == v9039 {
		v9034 = v9034 + v9042
		v9035 = v9035 + v9042
		goto L2480
	} else {
		goto L2483
	}
L2483:
	;
	goto L2481
L2484:
	;
	v9052 = F_plpgsql_parse_err_condition(m, v9025)
	mBase = m.M
	v9053 = m.ExcPending
	if v9053 != 0 {
		goto L66
	} else {
		goto L2487
	}
L2485:
	;
	goto L2486
L2486:
	;
	v9059 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v9060 = m.ExcPending
	if v9060 != 0 {
		goto L66
	} else {
		goto L2488
	}
L2487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9052
	v10317 = v241
	goto L5
L2488:
	;
	if v9059 != int32(261) {
		goto L16
	} else {
		goto L2489
	}
L2489:
	;
	v9063 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1362])))
	if v9063&int32(3) == int32(0) {
		v9087 = v9063
		goto L2492
	} else {
		goto L2493
	}
L2490:
	;
	if v9120 != int32(5) {
		goto L15
	} else {
		goto L2507
	}
L2491:
	;
	v9120 = v9112 - v9063
	goto L2490
L2492:
	;
	v9091 = v9087
	goto L2501
L2493:
	;
	v9071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063))))
	if v9071 == int32(0) {
		goto L2494
	} else {
		goto L2495
	}
L2494:
	;
	v9120 = int32(0)
	goto L2490
L2495:
	;
	goto L2496
L2496:
	;
	v9076 = v9063
	goto L2497
L2497:
	;
	v9080 = v9076 + int32(1)
	if v9080&int32(3) == int32(0) {
		v9087 = v9080
		goto L2492
	} else {
		goto L2499
	}
L2498:
	;
	v9112 = v9080
	goto L2491
L2499:
	;
	v9085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9080))))
	if v9085 != 0 {
		v9076 = v9080
		goto L2497
	} else {
		goto L2500
	}
L2500:
	;
	goto L2498
L2501:
	;
	v9097 = *(*int32)(unsafe.Add(mBase, uint32(v9091)))
	v9100 = int32(-2139062144)
	if (int32(16843008)-v9097|v9097)&v9100 == v9100 {
		v9091 = v9091 + int32(4)
		goto L2501
	} else {
		goto L2503
	}
L2502:
	;
	v9106 = v9091
	goto L2504
L2503:
	;
	goto L2502
L2504:
	;
	v9110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9106))))
	if v9110 != 0 {
		v9106 = v9106 + int32(1)
		goto L2504
	} else {
		goto L2506
	}
L2505:
	;
	v9112 = v9106
	goto L2491
L2506:
	;
	goto L2505
L2507:
	;
	v9123 = int32(500110)
	v9127 = m.G0
	v9129 = v9127 - int32(32)
	v9130 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9129)+24)) = v9130
	*(*int64)(unsafe.Add(mBase, uint32(v9129)+16)) = v9130
	*(*int64)(unsafe.Add(mBase, uint32(v9129)+8)) = v9130
	*(*int64)(unsafe.Add(mBase, uint32(v9129))) = v9130
	v9138 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1405])))
	if v9138 == int32(0) {
		goto L2509
	} else {
		goto L2510
	}
L2508:
	;
	if v9206 != int32(5) {
		goto L14
	} else {
		goto L2529
	}
L2509:
	;
	v9206 = int32(0)
	goto L2508
L2510:
	;
	goto L2511
L2511:
	;
	v9142 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1406])))
	if v9142 == int32(0) {
		goto L2512
	} else {
		goto L2513
	}
L2512:
	;
	v9146 = v9063
	goto L2515
L2513:
	;
	goto L2514
L2514:
	;
	v9156 = v9123
	v9157 = v9138
	goto L2518
L2515:
	;
	v9152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9146))))
	if v9152 == v9138 {
		v9146 = v9146 + int32(1)
		goto L2515
	} else {
		goto L2517
	}
L2516:
	;
	v9206 = v9146 - v9063
	goto L2508
L2517:
	;
	goto L2516
L2518:
	;
	v9164 = v9129 + int32(base.Ui32(v9157)>>(uint(int32(3))%32))&int32(28)
	v9165 = *(*int32)(unsafe.Add(mBase, uint32(v9164)))
	v9166 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9164))) = v9165 | v9166<<(uint(v9157)%32)
	v9170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9156)+1)))
	if v9170 != 0 {
		v9156 = v9156 + v9166
		v9157 = v9170
		goto L2518
	} else {
		goto L2520
	}
L2519:
	;
	v9173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063))))
	if v9173 == int32(0) {
		v9198 = v9063
		goto L2521
	} else {
		goto L2522
	}
L2520:
	;
	goto L2519
L2521:
	;
	v9206 = v9198 - v9063
	goto L2508
L2522:
	;
	v9177 = v9063
	v9178 = v9173
	goto L2523
L2523:
	;
	v9186 = *(*int32)(unsafe.Add(mBase, uint32(v9129+int32(base.Ui32(v9178)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v9186)>>(uint(v9178)%32))&int32(1) == int32(0) {
		goto L2525
	} else {
		goto L2526
	}
L2524:
	;
	v9198 = v9194
	goto L2521
L2525:
	;
	v9198 = v9177
	goto L2521
L2526:
	;
	goto L2527
L2527:
	;
	v9192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9177)+1)))
	v9194 = v9177 + int32(1)
	if v9192 != 0 {
		v9177 = v9194
		v9178 = v9192
		goto L2523
	} else {
		goto L2528
	}
L2528:
	;
	goto L2524
L2529:
	;
	v9210 = F_palloc(m, int32(12))
	mBase = m.M
	v9211 = m.ExcPending
	if v9211 != 0 {
		goto L66
	} else {
		goto L2530
	}
L2530:
	;
	v9212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063)+4)))
	v9213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063)+3)))
	v9214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063)+2)))
	v9215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063)+1)))
	v9216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9063))))
	*(*int32)(unsafe.Add(mBase, uint32(v9210)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9210)+4)) = v9063
	v9220 = int32(16)
	v9222 = int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v9210))) = (v9216+v9220)&v9222 | (v9215+v9220)&v9222<<(uint(int32(6))%32) | (v9214+v9220)&v9222<<(uint(int32(12))%32) | (v9213+v9220)&v9222<<(uint(int32(18))%32) | (v9212+v9220)&v9222<<(uint(int32(24))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9210
	v10317 = v241
	goto L5
L2531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9267
	v10317 = v241
	goto L5
L2532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9283
	v10317 = v241
	goto L5
L2533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9299
	v10317 = v241
	goto L5
L2534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L2535:
	;
	v9314 = *(*int32)(unsafe.Add(mBase, uint32(v9309)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9314
	v10317 = v241
	goto L5
L2536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = int32(0)
	v10317 = v241
	goto L5
L2537:
	;
	v9328 = *(*int32)(unsafe.Add(mBase, uint32(v9323)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9328
	v10317 = v241
	goto L5
L2538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9341
	v10317 = v241
	goto L5
L2539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v9344
	v10317 = v241
	goto L5
L2540:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2541:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2542:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9370 = m.ExcPending
	if v9370 != 0 {
		goto L66
	} else {
		goto L2543
	}
L2543:
	;
	v9371 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	v9372 = *(*int32)(unsafe.Add(mBase, uint32(v9371)+4))
	v9373 = F_format_type_be(m, v9372)
	mBase = m.M
	v9374 = m.ExcPending
	if v9374 != 0 {
		goto L66
	} else {
		goto L2544
	}
L2544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v9373
	F_errmsg(m, int32(185682), v27+int32(32))
	mBase = m.M
	v9380 = m.ExcPending
	if v9380 != 0 {
		goto L66
	} else {
		goto L2545
	}
L2545:
	;
	v9383 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(8))))
	v9384 = F_plpgsql_scanner_errposition(m, v9383, l1)
	mBase = m.M
	v9385 = m.ExcPending
	if v9385 != 0 {
		goto L66
	} else {
		goto L2546
	}
L2546:
	;
	F_errfinish(m, int32(26734), int32(523), int32(355297))
	mBase = m.M
	v9390 = m.ExcPending
	if v9390 != 0 {
		goto L66
	} else {
		goto L2547
	}
L2547:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2548:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v9397 = m.ExcPending
	if v9397 != 0 {
		goto L66
	} else {
		goto L2549
	}
L2549:
	;
	v9398 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v9398
	F_errmsg(m, int32(524464), v27+int32(16))
	mBase = m.M
	v9404 = m.ExcPending
	if v9404 != 0 {
		goto L66
	} else {
		goto L2550
	}
L2550:
	;
	v9407 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9408 = F_plpgsql_scanner_errposition(m, v9407, l1)
	mBase = m.M
	v9409 = m.ExcPending
	if v9409 != 0 {
		goto L66
	} else {
		goto L2551
	}
L2551:
	;
	F_errfinish(m, int32(26734), int32(542), int32(355297))
	mBase = m.M
	v9414 = m.ExcPending
	if v9414 != 0 {
		goto L66
	} else {
		goto L2552
	}
L2552:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2553:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9421 = m.ExcPending
	if v9421 != 0 {
		goto L66
	} else {
		goto L2554
	}
L2554:
	;
	v9422 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v9422
	F_errmsg(m, int32(71688), v27+int32(48))
	mBase = m.M
	v9428 = m.ExcPending
	if v9428 != 0 {
		goto L66
	} else {
		goto L2555
	}
L2555:
	;
	v9429 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9430 = F_plpgsql_scanner_errposition(m, v9429, l1)
	mBase = m.M
	v9431 = m.ExcPending
	if v9431 != 0 {
		goto L66
	} else {
		goto L2556
	}
L2556:
	;
	F_errfinish(m, int32(26734), int32(667), int32(355297))
	mBase = m.M
	v9436 = m.ExcPending
	if v9436 != 0 {
		goto L66
	} else {
		goto L2557
	}
L2557:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2558:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9443 = m.ExcPending
	if v9443 != 0 {
		goto L66
	} else {
		goto L2559
	}
L2559:
	;
	v9444 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v9444
	F_errmsg(m, int32(71688), v27-int32(-64))
	mBase = m.M
	v9450 = m.ExcPending
	if v9450 != 0 {
		goto L66
	} else {
		goto L2560
	}
L2560:
	;
	v9451 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9452 = F_plpgsql_scanner_errposition(m, v9451, l1)
	mBase = m.M
	v9453 = m.ExcPending
	if v9453 != 0 {
		goto L66
	} else {
		goto L2561
	}
L2561:
	;
	F_errfinish(m, int32(26734), int32(682), int32(355297))
	mBase = m.M
	v9458 = m.ExcPending
	if v9458 != 0 {
		goto L66
	} else {
		goto L2562
	}
L2562:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2563:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v9466 = m.ExcPending
	if v9466 != 0 {
		goto L66
	} else {
		goto L2564
	}
L2564:
	;
	v9467 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v9468 = F_NameListToString(m, v9467)
	mBase = m.M
	v9469 = m.ExcPending
	if v9469 != 0 {
		goto L66
	} else {
		goto L2565
	}
L2565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v9468
	F_errmsg(m, int32(71688), v27+int32(80))
	mBase = m.M
	v9475 = m.ExcPending
	if v9475 != 0 {
		goto L66
	} else {
		goto L2566
	}
L2566:
	;
	v9476 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9477 = F_plpgsql_scanner_errposition(m, v9476, l1)
	mBase = m.M
	v9478 = m.ExcPending
	if v9478 != 0 {
		goto L66
	} else {
		goto L2567
	}
L2567:
	;
	F_errfinish(m, int32(26734), int32(708), int32(355297))
	mBase = m.M
	v9483 = m.ExcPending
	if v9483 != 0 {
		goto L66
	} else {
		goto L2568
	}
L2568:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2569:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2570:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2571:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2572:
	;
	F_errmsg_internal(m, int32(161068), int32(0))
	mBase = m.M
	v9510 = m.ExcPending
	if v9510 != 0 {
		goto L66
	} else {
		goto L2573
	}
L2573:
	;
	F_errfinish(m, int32(26734), int32(990), int32(355297))
	mBase = m.M
	v9515 = m.ExcPending
	if v9515 != 0 {
		goto L66
	} else {
		goto L2574
	}
L2574:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2575:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9522 = m.ExcPending
	if v9522 != 0 {
		goto L66
	} else {
		goto L2576
	}
L2576:
	;
	v9523 = F_NameOfDatum(m, v132)
	mBase = m.M
	v9524 = m.ExcPending
	if v9524 != 0 {
		goto L66
	} else {
		goto L2577
	}
L2577:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+176)) = v9523
	F_errmsg(m, int32(390729), v27+int32(176))
	mBase = m.M
	v9530 = m.ExcPending
	if v9530 != 0 {
		goto L66
	} else {
		goto L2578
	}
L2578:
	;
	v9531 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9532 = F_plpgsql_scanner_errposition(m, v9531, l1)
	mBase = m.M
	v9533 = m.ExcPending
	if v9533 != 0 {
		goto L66
	} else {
		goto L2579
	}
L2579:
	;
	F_errfinish(m, int32(26734), int32(1174), int32(355297))
	mBase = m.M
	v9538 = m.ExcPending
	if v9538 != 0 {
		goto L66
	} else {
		goto L2580
	}
L2580:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2581:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9545 = m.ExcPending
	if v9545 != 0 {
		goto L66
	} else {
		goto L2582
	}
L2582:
	;
	F_errmsg(m, int32(164314), int32(0))
	mBase = m.M
	v9549 = m.ExcPending
	if v9549 != 0 {
		goto L66
	} else {
		goto L2583
	}
L2583:
	;
	v9552 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9553 = F_plpgsql_scanner_errposition(m, v9552, l1)
	mBase = m.M
	v9554 = m.ExcPending
	if v9554 != 0 {
		goto L66
	} else {
		goto L2584
	}
L2584:
	;
	F_errfinish(m, int32(26734), int32(1403), int32(355297))
	mBase = m.M
	v9559 = m.ExcPending
	if v9559 != 0 {
		goto L66
	} else {
		goto L2585
	}
L2585:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2586:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9566 = m.ExcPending
	if v9566 != 0 {
		goto L66
	} else {
		goto L2587
	}
L2587:
	;
	F_errmsg(m, int32(390507), int32(0))
	mBase = m.M
	v9570 = m.ExcPending
	if v9570 != 0 {
		goto L66
	} else {
		goto L2588
	}
L2588:
	;
	v9573 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9574 = F_plpgsql_scanner_errposition(m, v9573, l1)
	mBase = m.M
	v9575 = m.ExcPending
	if v9575 != 0 {
		goto L66
	} else {
		goto L2589
	}
L2589:
	;
	F_errfinish(m, int32(26734), int32(1438), int32(355297))
	mBase = m.M
	v9580 = m.ExcPending
	if v9580 != 0 {
		goto L66
	} else {
		goto L2590
	}
L2590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2591:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9587 = m.ExcPending
	if v9587 != 0 {
		goto L66
	} else {
		goto L2592
	}
L2592:
	;
	F_errmsg(m, int32(390655), int32(0))
	mBase = m.M
	v9591 = m.ExcPending
	if v9591 != 0 {
		goto L66
	} else {
		goto L2593
	}
L2593:
	;
	v9592 = F_plpgsql_scanner_errposition(m, v4857, l1)
	mBase = m.M
	v9593 = m.ExcPending
	if v9593 != 0 {
		goto L66
	} else {
		goto L2594
	}
L2594:
	;
	F_errfinish(m, int32(26734), int32(1445), int32(355297))
	mBase = m.M
	v9598 = m.ExcPending
	if v9598 != 0 {
		goto L66
	} else {
		goto L2595
	}
L2595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2596:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9605 = m.ExcPending
	if v9605 != 0 {
		goto L66
	} else {
		goto L2597
	}
L2597:
	;
	F_errmsg(m, int32(164314), int32(0))
	mBase = m.M
	v9609 = m.ExcPending
	if v9609 != 0 {
		goto L66
	} else {
		goto L2598
	}
L2598:
	;
	v9612 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9613 = F_plpgsql_scanner_errposition(m, v9612, l1)
	mBase = m.M
	v9614 = m.ExcPending
	if v9614 != 0 {
		goto L66
	} else {
		goto L2599
	}
L2599:
	;
	F_errfinish(m, int32(26734), int32(1596), int32(355297))
	mBase = m.M
	v9619 = m.ExcPending
	if v9619 != 0 {
		goto L66
	} else {
		goto L2600
	}
L2600:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2601:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9626 = m.ExcPending
	if v9626 != 0 {
		goto L66
	} else {
		goto L2602
	}
L2602:
	;
	F_errmsg(m, int32(164492), int32(0))
	mBase = m.M
	v9630 = m.ExcPending
	if v9630 != 0 {
		goto L66
	} else {
		goto L2603
	}
L2603:
	;
	v9633 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(20))))
	v9634 = F_plpgsql_scanner_errposition(m, v9633, l1)
	mBase = m.M
	v9635 = m.ExcPending
	if v9635 != 0 {
		goto L66
	} else {
		goto L2604
	}
L2604:
	;
	F_errfinish(m, int32(26734), int32(1701), int32(355297))
	mBase = m.M
	v9640 = m.ExcPending
	if v9640 != 0 {
		goto L66
	} else {
		goto L2605
	}
L2605:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2606:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9647 = m.ExcPending
	if v9647 != 0 {
		goto L66
	} else {
		goto L2607
	}
L2607:
	;
	v9648 = *(*int32)(unsafe.Add(mBase, uint32(v5586)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+224)) = v9648
	F_errmsg(m, int32(94308), v27+int32(224))
	mBase = m.M
	v9654 = m.ExcPending
	if v9654 != 0 {
		goto L66
	} else {
		goto L2608
	}
L2608:
	;
	v9657 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v9658 = F_plpgsql_scanner_errposition(m, v9657, l1)
	mBase = m.M
	v9659 = m.ExcPending
	if v9659 != 0 {
		goto L66
	} else {
		goto L2609
	}
L2609:
	;
	F_errfinish(m, int32(26734), int32(1745), int32(355297))
	mBase = m.M
	v9664 = m.ExcPending
	if v9664 != 0 {
		goto L66
	} else {
		goto L2610
	}
L2610:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2611:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9671 = m.ExcPending
	if v9671 != 0 {
		goto L66
	} else {
		goto L2612
	}
L2612:
	;
	v9674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5520)+12)))
	if v9674 != 0 {
		goto L2613
	} else {
		goto L2614
	}
L2613:
	;
	v9675 = int32(303524)
	goto L2615
L2614:
	;
	v9675 = int32(231398)
	goto L2615
L2615:
	;
	F_errmsg(m, v9675, int32(0))
	mBase = m.M
	v9678 = m.ExcPending
	if v9678 != 0 {
		goto L66
	} else {
		goto L2616
	}
L2616:
	;
	v9679 = *(*int32)(unsafe.Add(mBase, uint32(v5536)))
	v9680 = F_plpgsql_scanner_errposition(m, v9679, l1)
	mBase = m.M
	v9681 = m.ExcPending
	if v9681 != 0 {
		goto L66
	} else {
		goto L2617
	}
L2617:
	;
	F_errfinish(m, int32(26734), int32(1767), int32(355297))
	mBase = m.M
	v9686 = m.ExcPending
	if v9686 != 0 {
		goto L66
	} else {
		goto L2618
	}
L2618:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2619:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9693 = m.ExcPending
	if v9693 != 0 {
		goto L66
	} else {
		goto L2620
	}
L2620:
	;
	F_errmsg(m, int32(250744), int32(0))
	mBase = m.M
	v9697 = m.ExcPending
	if v9697 != 0 {
		goto L66
	} else {
		goto L2621
	}
L2621:
	;
	v9698 = F_plpgsql_scanner_errposition(m, v5703, l1)
	mBase = m.M
	v9699 = m.ExcPending
	if v9699 != 0 {
		goto L66
	} else {
		goto L2622
	}
L2622:
	;
	F_errfinish(m, int32(26734), int32(3445), int32(95794))
	mBase = m.M
	v9704 = m.ExcPending
	if v9704 != 0 {
		goto L66
	} else {
		goto L2623
	}
L2623:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2624:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9711 = m.ExcPending
	if v9711 != 0 {
		goto L66
	} else {
		goto L2625
	}
L2625:
	;
	F_errmsg(m, int32(130624), int32(0))
	mBase = m.M
	v9715 = m.ExcPending
	if v9715 != 0 {
		goto L66
	} else {
		goto L2626
	}
L2626:
	;
	v9716 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	v9717 = F_plpgsql_scanner_errposition(m, v9716, l1)
	mBase = m.M
	v9718 = m.ExcPending
	if v9718 != 0 {
		goto L66
	} else {
		goto L2627
	}
L2627:
	;
	F_errfinish(m, int32(26734), int32(3460), int32(95794))
	mBase = m.M
	v9723 = m.ExcPending
	if v9723 != 0 {
		goto L66
	} else {
		goto L2628
	}
L2628:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2629:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9730 = m.ExcPending
	if v9730 != 0 {
		goto L66
	} else {
		goto L2630
	}
L2630:
	;
	F_errmsg(m, int32(250696), int32(0))
	mBase = m.M
	v9734 = m.ExcPending
	if v9734 != 0 {
		goto L66
	} else {
		goto L2631
	}
L2631:
	;
	v9735 = F_plpgsql_scanner_errposition(m, v5859, l1)
	mBase = m.M
	v9736 = m.ExcPending
	if v9736 != 0 {
		goto L66
	} else {
		goto L2632
	}
L2632:
	;
	F_errfinish(m, int32(26734), int32(3509), int32(95771))
	mBase = m.M
	v9741 = m.ExcPending
	if v9741 != 0 {
		goto L66
	} else {
		goto L2633
	}
L2633:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2634:
	;
	F_errmsg(m, int32(358437), int32(0))
	mBase = m.M
	v9748 = m.ExcPending
	if v9748 != 0 {
		goto L66
	} else {
		goto L2635
	}
L2635:
	;
	v9749 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	v9750 = F_plpgsql_scanner_errposition(m, v9749, l1)
	mBase = m.M
	v9751 = m.ExcPending
	if v9751 != 0 {
		goto L66
	} else {
		goto L2636
	}
L2636:
	;
	F_errfinish(m, int32(26734), int32(3383), int32(95826))
	mBase = m.M
	v9756 = m.ExcPending
	if v9756 != 0 {
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9763 = m.ExcPending
	if v9763 != 0 {
		goto L66
	} else {
		goto L2639
	}
L2639:
	;
	F_errmsg(m, int32(130692), int32(0))
	mBase = m.M
	v9767 = m.ExcPending
	if v9767 != 0 {
		goto L66
	} else {
		goto L2640
	}
L2640:
	;
	v9768 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	v9769 = F_plpgsql_scanner_errposition(m, v9768, l1)
	mBase = m.M
	v9770 = m.ExcPending
	if v9770 != 0 {
		goto L66
	} else {
		goto L2641
	}
L2641:
	;
	F_errfinish(m, int32(26734), int32(3397), int32(95826))
	mBase = m.M
	v9775 = m.ExcPending
	if v9775 != 0 {
		goto L66
	} else {
		goto L2642
	}
L2642:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2643:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2644:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2645:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2646:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2647:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2648:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2649:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9819 = m.ExcPending
	if v9819 != 0 {
		goto L66
	} else {
		goto L2650
	}
L2650:
	;
	F_errmsg(m, int32(531097), int32(0))
	mBase = m.M
	v9823 = m.ExcPending
	if v9823 != 0 {
		goto L66
	} else {
		goto L2651
	}
L2651:
	;
	F_errfinish(m, int32(26734), int32(4162), int32(130160))
	mBase = m.M
	v9828 = m.ExcPending
	if v9828 != 0 {
		goto L66
	} else {
		goto L2652
	}
L2652:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2653:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2654:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2655:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2656:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2657:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2658:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v9859 = m.ExcPending
	if v9859 != 0 {
		goto L66
	} else {
		goto L2659
	}
L2659:
	;
	F_errmsg(m, int32(112125), int32(0))
	mBase = m.M
	v9863 = m.ExcPending
	if v9863 != 0 {
		goto L66
	} else {
		goto L2660
	}
L2660:
	;
	v9866 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(12))))
	v9867 = F_plpgsql_scanner_errposition(m, v9866, l1)
	mBase = m.M
	v9868 = m.ExcPending
	if v9868 != 0 {
		goto L66
	} else {
		goto L2661
	}
L2661:
	;
	F_errfinish(m, int32(26734), int32(2197), int32(355297))
	mBase = m.M
	v9873 = m.ExcPending
	if v9873 != 0 {
		goto L66
	} else {
		goto L2662
	}
L2662:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2663:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9880 = m.ExcPending
	if v9880 != 0 {
		goto L66
	} else {
		goto L2664
	}
L2664:
	;
	F_errmsg(m, int32(390819), int32(0))
	mBase = m.M
	v9884 = m.ExcPending
	if v9884 != 0 {
		goto L66
	} else {
		goto L2665
	}
L2665:
	;
	v9885 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9886 = F_plpgsql_scanner_errposition(m, v9885, l1)
	mBase = m.M
	v9887 = m.ExcPending
	if v9887 != 0 {
		goto L66
	} else {
		goto L2666
	}
L2666:
	;
	F_errfinish(m, int32(26734), int32(2294), int32(355297))
	mBase = m.M
	v9892 = m.ExcPending
	if v9892 != 0 {
		goto L66
	} else {
		goto L2667
	}
L2667:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2668:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v9899 = m.ExcPending
	if v9899 != 0 {
		goto L66
	} else {
		goto L2669
	}
L2669:
	;
	v9900 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v9901 = *(*int32)(unsafe.Add(mBase, uint32(v9900)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v9901
	F_errmsg(m, int32(207245), v27+int32(256))
	mBase = m.M
	v9907 = m.ExcPending
	if v9907 != 0 {
		goto L66
	} else {
		goto L2670
	}
L2670:
	;
	v9908 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v9909 = F_plpgsql_scanner_errposition(m, v9908, l1)
	mBase = m.M
	v9910 = m.ExcPending
	if v9910 != 0 {
		goto L66
	} else {
		goto L2671
	}
L2671:
	;
	F_errfinish(m, int32(26734), int32(2301), int32(355297))
	mBase = m.M
	v9915 = m.ExcPending
	if v9915 != 0 {
		goto L66
	} else {
		goto L2672
	}
L2672:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2673:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2674:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2675:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2676:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2678:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9955 = m.ExcPending
	if v9955 != 0 {
		goto L66
	} else {
		goto L2679
	}
L2679:
	;
	F_errmsg(m, int32(531057), int32(0))
	mBase = m.M
	v9959 = m.ExcPending
	if v9959 != 0 {
		goto L66
	} else {
		goto L2680
	}
L2680:
	;
	F_errfinish(m, int32(26734), int32(4158), int32(130160))
	mBase = m.M
	v9964 = m.ExcPending
	if v9964 != 0 {
		goto L66
	} else {
		goto L2681
	}
L2681:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2682:
	;
	v9984 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1376])))
	if v9984 == int32(269) {
		v10005 = v9982
		v10006 = v9966
		goto L9
	} else {
		goto L2683
	}
L2683:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v9990 = m.ExcPending
	if v9990 != 0 {
		goto L66
	} else {
		goto L2684
	}
L2684:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v9993 = m.ExcPending
	if v9993 != 0 {
		goto L66
	} else {
		goto L2685
	}
L2685:
	;
	F_errmsg(m, int32(231437), int32(0))
	mBase = m.M
	v9997 = m.ExcPending
	if v9997 != 0 {
		goto L66
	} else {
		goto L2686
	}
L2686:
	;
	v9998 = F_plpgsql_scanner_errposition(m, v4857, l1)
	mBase = m.M
	v9999 = m.ExcPending
	if v9999 != 0 {
		goto L66
	} else {
		goto L2687
	}
L2687:
	;
	F_errfinish(m, int32(26734), int32(1569), int32(355297))
	mBase = m.M
	v10004 = m.ExcPending
	if v10004 != 0 {
		goto L66
	} else {
		goto L2688
	}
L2688:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2689:
	;
	v10013 = *(*int32)(unsafe.Add(mBase, uint32(v10005)))
	v10014 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1396])))
	v10015 = int32(4470752)
	v10016 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v10019 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v10019
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v10014
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1379]))) = int32(6372)
	v10025 = int32(4463656)
	v10026 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4784)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366]))) = v10026
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1380]))) = v27 + int32(4768)
	v10036 = F_raw_parser(m, v10013, int32(2))
	mBase = m.M
	v10037 = m.ExcPending
	if v10037 != 0 {
		goto L66
	} else {
		goto L2692
	}
L2690:
	;
	goto L2691
L2691:
	;
	v10047 = int32(0)
	v10053 = int32(1)
	v10062 = F_read_sql_construct(m, int32(336), int32(288), v10047, int32(518719), int32(2), v10053, v10053, v10047, v27+int32(4752), v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v10063 = m.ExcPending
	if v10063 != 0 {
		goto L66
	} else {
		goto L2693
	}
L2692:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v10016
	v10041 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v10041
	goto L2691
L2693:
	;
	v10064 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1376])))
	if v10064 == int32(288) {
		goto L2694
	} else {
		goto L2695
	}
L2694:
	;
	v10068 = int32(0)
	v10072 = int32(1)
	v10080 = F_read_sql_construct(m, int32(336), v10068, v10068, int32(518719), int32(2), v10072, v10072, v10068, v10068, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v10081 = m.ExcPending
	if v10081 != 0 {
		goto L66
	} else {
		goto L2697
	}
L2695:
	;
	v10082 = v10047
	goto L2696
L2696:
	;
	v10085 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(8))))
	if v10085 != 0 {
		goto L2699
	} else {
		goto L2700
	}
L2697:
	;
	v10082 = v10080
	goto L2696
L2698:
	;
	F_errstart_cold(m, int32(21), int32(541439))
	mBase = m.M
	v10125 = m.ExcPending
	if v10125 != 0 {
		goto L66
	} else {
		goto L2706
	}
L2699:
	;
	v10088 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(4))))
	if v10088 != 0 {
		goto L2698
	} else {
		goto L2702
	}
L2700:
	;
	goto L2701
L2701:
	;
	v10091 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(16))))
	v10094 = *(*int32)(unsafe.Add(mBase, uint32(v132-int32(12))))
	v10097 = int32(0)
	v10099 = F_plpgsql_build_datatype(m, int32(23), int32(-1), v10097, v10097)
	mBase = m.M
	v10100 = m.ExcPending
	if v10100 != 0 {
		goto L66
	} else {
		goto L2703
	}
L2702:
	;
	goto L2701
L2703:
	;
	v10102 = F_plpgsql_build_variable(m, v10091, v10094, v10099, int32(1))
	mBase = m.M
	v10103 = m.ExcPending
	if v10103 != 0 {
		goto L66
	} else {
		goto L2704
	}
L2704:
	;
	v10105 = F_palloc0(m, int32(40))
	mBase = m.M
	v10106 = m.ExcPending
	if v10106 != 0 {
		goto L66
	} else {
		goto L2705
	}
L2705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10105))) = int32(6)
	v10110 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v10111 = *(*int32)(unsafe.Add(mBase, uint32(v10110)+520))
	v10113 = v10111 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10110)+520)) = v10113
	*(*int32)(unsafe.Add(mBase, uint32(v10105)+32)) = v10006
	*(*int32)(unsafe.Add(mBase, uint32(v10105)+16)) = v10102
	*(*int32)(unsafe.Add(mBase, uint32(v10105)+8)) = v10113
	*(*int32)(unsafe.Add(mBase, uint32(v10105)+28)) = v10082
	*(*int32)(unsafe.Add(mBase, uint32(v10105)+24)) = v10062
	*(*int32)(unsafe.Add(mBase, uint32(v10105)+20)) = v10005
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v10105
	v10317 = v241
	goto L5
L2706:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v10128 = m.ExcPending
	if v10128 != 0 {
		goto L66
	} else {
		goto L2707
	}
L2707:
	;
	F_errmsg(m, int32(390558), int32(0))
	mBase = m.M
	v10132 = m.ExcPending
	if v10132 != 0 {
		goto L66
	} else {
		goto L2708
	}
L2708:
	;
	v10135 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v10136 = F_plpgsql_scanner_errposition(m, v10135, l1)
	mBase = m.M
	v10137 = m.ExcPending
	if v10137 != 0 {
		goto L66
	} else {
		goto L2709
	}
L2709:
	;
	F_errfinish(m, int32(26734), int32(1535), int32(355297))
	mBase = m.M
	v10142 = m.ExcPending
	if v10142 != 0 {
		goto L66
	} else {
		goto L2710
	}
L2710:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2711:
	;
	v10288 = v2506
	goto L6
L2712:
	;
	if v10163 <= int32(292) {
		goto L2719
	} else {
		goto L2720
	}
L2713:
	;
	F_initStringInfo(m, v27+int32(4784))
	mBase = m.M
	v10229 = m.ExcPending
	if v10229 != 0 {
		goto L66
	} else {
		goto L2737
	}
L2714:
	;
	goto L2713
L2715:
	;
	if base.B2i32(v10163 != int32(44))&base.B2i32(v10163 != int32(41))|v10157 == int32(0) {
		goto L2714
	} else {
		goto L2729
	}
L2716:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(158828))
	mBase = m.M
	v10201 = m.ExcPending
	if v10201 != 0 {
		goto L66
	} else {
		goto L2728
	}
L2717:
	;
	if v10163 != int32(343) {
		goto L2715
	} else {
		goto L2727
	}
L2718:
	;
	if v10157 != 0 {
		goto L2716
	} else {
		goto L2725
	}
L2719:
	;
	switch v10163 - int32(59) {
	case 0, 2:
		goto L2714
	case 1:
		goto L2715
	default:
		goto L2722
	}
L2720:
	;
	goto L2721
L2721:
	;
	switch v10163 - int32(293) {
	case 0, 13:
		goto L2714
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		goto L2715
	default:
		goto L2717
	}
L2722:
	;
	if v10163 == int32(270) {
		goto L2714
	} else {
		goto L2723
	}
L2723:
	;
	if v10163 == int32(0) {
		goto L2718
	} else {
		goto L2724
	}
L2724:
	;
	goto L2715
L2725:
	;
	F_plpgsql_yyerror(m, v27+int32(4732), int32(0), l1, int32(257602))
	mBase = m.M
	v10193 = m.ExcPending
	if v10193 != 0 {
		goto L66
	} else {
		goto L2726
	}
L2726:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2727:
	;
	goto L2714
L2728:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2729:
	;
	if v10163 == int32(41) {
		goto L2730
	} else {
		goto L2731
	}
L2730:
	;
	v10215 = int32(-1)
	goto L2732
L2731:
	;
	v10215 = int32(0)
	goto L2732
L2732:
	;
	if v10163 == int32(40) {
		goto L2733
	} else {
		goto L2734
	}
L2733:
	;
	v10218 = int32(1)
	goto L2735
L2734:
	;
	v10218 = v10215
	goto L2735
L2735:
	;
	v10224 = F_plpgsql_yylex(m, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v10225 = m.ExcPending
	if v10225 != 0 {
		goto L66
	} else {
		goto L2736
	}
L2736:
	;
	v10157 = v10218 + v10157
	v10163 = v10224
	goto L2712
L2737:
	;
	v10232 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1364])))
	F_plpgsql_append_source_text(m, v27+int32(4784), v2225, v10232, l1)
	mBase = m.M
	v10234 = m.ExcPending
	if v10234 != 0 {
		goto L66
	} else {
		goto L2738
	}
L2738:
	;
	v10235 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	v10236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10235))))
	if v10236 == int32(0) {
		goto L3
	} else {
		goto L2739
	}
L2739:
	;
	v10239 = int32(4463656)
	v10240 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v27 + int32(4768)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1424]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1376]))) = v2225
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1377]))) = int32(6372)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378]))) = v10240
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1425]))) = v27 + int32(4752)
	v10253 = int32(0)
	v10255 = F_typeStringToTypeName(m, v10235, v10253)
	mBase = m.M
	v10256 = m.ExcPending
	if v10256 != 0 {
		goto L66
	} else {
		goto L2740
	}
L2740:
	;
	F_typenameTypeIdAndMod(m, v10253, v10255, v27+int32(4764), v27+int32(4760))
	mBase = m.M
	v10262 = m.ExcPending
	if v10262 != 0 {
		goto L66
	} else {
		goto L2741
	}
L2741:
	;
	v10264 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1378])))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v10264
	v10266 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1396])))
	v10267 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1426])))
	v10269 = *(*int32)(unsafe.Add(mBase, _consts[1367]))
	v10270 = *(*int32)(unsafe.Add(mBase, uint32(v10269)+44))
	v10271 = F_plpgsql_build_datatype(m, v10266, v10267, v10270, v10255)
	mBase = m.M
	v10272 = m.ExcPending
	if v10272 != 0 {
		goto L66
	} else {
		goto L2742
	}
L2742:
	;
	v10273 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1366])))
	F_pfree(m, v10273)
	mBase = m.M
	v10275 = m.ExcPending
	if v10275 != 0 {
		goto L66
	} else {
		goto L2743
	}
L2743:
	;
	F_plpgsql_push_back_token(m, v10163, v27+int32(4736), v27+int32(4732), l1)
	mBase = m.M
	v10281 = m.ExcPending
	if v10281 != 0 {
		goto L66
	} else {
		goto L2744
	}
L2744:
	;
	v10288 = v10271
	goto L6
L2745:
	;
	v10376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10354)+uint32(_consts[1427]))))
	v10380 = v10376
	v10382 = v10343
	v10386 = v10317
	v10391 = v10346
	v10394 = v260
	goto L4
L2746:
	;
	v10365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10358<<(uint(int32(1))%32))+uint32(_consts[1360]))))
	if v10365 != v10347&int32(65535) {
		goto L2745
	} else {
		goto L2747
	}
L2747:
	;
	v10373 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10358<<(uint(int32(1))%32))+uint32(_consts[1361]))))
	v10380 = v10373
	v10382 = v10343
	v10386 = v10317
	v10391 = v10346
	v10394 = v260
	goto L4
L2748:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
