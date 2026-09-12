package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_IvfflatCommitBuffer(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	F_GenericXLogFinish(m, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_UnlockReleaseBuffer(m, l0)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_IvfflatGetLists(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v2 == int32(0) {
		return int32(100)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		return v7
	}
}
func F_IvfflatInitPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	if l1&int32(3) != 0 {
	} else {
	}
	v29 = F___memset(m, l1, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(1572864)
	v35 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v35)
	v41 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v45 = l1 + v44
	v46 = int32(65412)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)) = uint16(v46)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(-1)
	return
}
func F_IvfflatKmeans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v193 int32
	_ = v193
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v219 float64
	_ = v219
	var v223 int32
	_ = v223
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
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
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v456 int64
	_ = v456
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v552 int32
	_ = v552
	var v553 int64
	_ = v553
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v730 int32
	_ = v730
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 float64
	_ = v774
	var v786 int32
	_ = v786
	var v825 float64
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 float64
	_ = v845
	var v848 float64
	_ = v848
	var v851 int32
	_ = v851
	var v852 float32
	_ = v852
	var v853 float64
	_ = v853
	var v855 float32
	_ = v855
	var v859 float64
	_ = v859
	var v860 float64
	_ = v860
	var v862 int32
	_ = v862
	var v908 float64
	_ = v908
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int64
	_ = v917
	var v918 int64
	_ = v918
	var v919 int64
	_ = v919
	var v940 float64
	_ = v940
	var v949 int32
	_ = v949
	var v987 float64
	_ = v987
	var v993 float32
	_ = v993
	var v995 float64
	_ = v995
	var v999 int32
	_ = v999
	var v1006 int32
	_ = v1006
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1106 int32
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1139 int32
	_ = v1139
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 float32
	_ = v1170
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1214 float32
	_ = v1214
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1227 float32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1233 float32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1239 float32
	_ = v1239
	var v1243 float32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 float32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 float32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 float32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 float32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
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
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1299 float32
	_ = v1299
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1324 int32
	_ = v1324
	var v1347 float32
	_ = v1347
	var v1358 float32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 float32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1374 int32
	_ = v1374
	var v1405 float32
	_ = v1405
	var v1414 int32
	_ = v1414
	var v1420 int32
	_ = v1420
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1505 int32
	_ = v1505
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1542 int32
	_ = v1542
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1599 int32
	_ = v1599
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 float64
	_ = v1651
	var v1654 float32
	_ = v1654
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1744 int32
	_ = v1744
	var v1787 int32
	_ = v1787
	var v1788 float32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1829 float32
	_ = v1829
	var v1841 float32
	_ = v1841
	var v1843 float32
	_ = v1843
	var v1844 float32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1852 float32
	_ = v1852
	var v1854 float32
	_ = v1854
	var v1855 float32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1867 int32
	_ = v1867
	var v1900 float32
	_ = v1900
	var v1914 float32
	_ = v1914
	var v1916 float32
	_ = v1916
	var v1917 float32
	_ = v1917
	var v1924 int32
	_ = v1924
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1990 int32
	_ = v1990
	var v1998 int32
	_ = v1998
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2029 float32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2035 float32
	_ = v2035
	var v2040 int32
	_ = v2040
	var v2047 int32
	_ = v2047
	var v2056 int32
	_ = v2056
	var v2062 int32
	_ = v2062
	var v2088 int32
	_ = v2088
	var v2090 float32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 float32
	_ = v2094
	var v2101 float32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2125 float64
	_ = v2125
	var v2126 float32
	_ = v2126
	var v2129 float32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2132 float32
	_ = v2132
	var v2133 float32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2143 float32
	_ = v2143
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2156 float64
	_ = v2156
	var v2157 float32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2176 int32
	_ = v2176
	var v2198 int32
	_ = v2198
	var v2225 int32
	_ = v2225
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2263 int32
	_ = v2263
	var v2268 int32
	_ = v2268
	var v2291 int32
	_ = v2291
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2337 int32
	_ = v2337
	var v2341 int32
	_ = v2341
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2393 int32
	_ = v2393
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2416 int32
	_ = v2416
	var v2463 int32
	_ = v2463
	var v2469 int32
	_ = v2469
	var v2476 int32
	_ = v2476
	var v2516 int32
	_ = v2516
	var v2524 int32
	_ = v2524
	var v2565 int32
	_ = v2565
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2575 int32
	_ = v2575
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2649 int32
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2690 int32
	_ = v2690
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2701 int32
	_ = v2701
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	var v2735 int32
	_ = v2735
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2824 int32
	_ = v2824
	var v2827 int32
	_ = v2827
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2838 int32
	_ = v2838
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2911 int32
	_ = v2911
	var v2942 int32
	_ = v2942
	var v2944 int32
	_ = v2944
	var v2948 int32
	_ = v2948
	var v2951 int32
	_ = v2951
	var v2960 int32
	_ = v2960
	var v3003 int32
	_ = v3003
	var v3006 int64
	_ = v3006
	var v3007 int64
	_ = v3007
	var v3008 int64
	_ = v3008
	var v3029 float64
	_ = v3029
	var v3033 int32
	_ = v3033
	var v3037 int32
	_ = v3037
	var v3045 int32
	_ = v3045
	var v3048 int32
	_ = v3048
	var v3087 int32
	_ = v3087
	var v3088 float32
	_ = v3088
	var v3096 float32
	_ = v3096
	var v3099 int32
	_ = v3099
	var v3100 float32
	_ = v3100
	var v3108 float32
	_ = v3108
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3113 int32
	_ = v3113
	var v3121 int32
	_ = v3121
	var v3165 int32
	_ = v3165
	var v3166 float32
	_ = v3166
	var v3174 float32
	_ = v3174
	var v3178 float32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3187 int32
	_ = v3187
	var v3190 int32
	_ = v3190
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3230 float32
	_ = v3230
	var v3234 int32
	_ = v3234
	var v3235 float32
	_ = v3235
	var v3239 int32
	_ = v3239
	var v3241 int32
	_ = v3241
	var v3249 int32
	_ = v3249
	var v3293 int32
	_ = v3293
	var v3294 float32
	_ = v3294
	var v3344 int32
	_ = v3344
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3401 int32
	_ = v3401
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3445 int32
	_ = v3445
	var v3448 int32
	_ = v3448
	var v3453 int32
	_ = v3453
	var v3455 int32
	_ = v3455
	var v3457 int32
	_ = v3457
	var v3458 int32
	_ = v3458
	var v3507 int32
	_ = v3507
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3515 int32
	_ = v3515
	var v3517 int32
	_ = v3517
	var v3519 int32
	_ = v3519
	var v3527 int32
	_ = v3527
	var v3568 int32
	_ = v3568
	var v3570 int32
	_ = v3570
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3585 float64
	_ = v3585
	var v3589 int32
	_ = v3589
	var v3594 int32
	_ = v3594
	var v3598 int32
	_ = v3598
	var v3603 int32
	_ = v3603
	var v3607 int32
	_ = v3607
	var v3611 int32
	_ = v3611
	var v3616 int32
	_ = v3616
	var v3620 int32
	_ = v3620
	var v3624 int32
	_ = v3624
	var v3629 int32
	_ = v3629
	var v3633 int32
	_ = v3633
	var v3637 int32
	_ = v3637
	var v3642 int32
	_ = v3642
	var v3689 int32
	_ = v3689
	var v3703 int32
	_ = v3703
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3751 int32
	_ = v3751
	var v3755 int32
	_ = v3755
	var v3793 int32
	_ = v3793
	var v3794 int32
	_ = v3794
	var v3795 float32
	_ = v3795
	var v3796 float32
	_ = v3796
	var v3798 float32
	_ = v3798
	var v3799 float32
	_ = v3799
	var v3802 float32
	_ = v3802
	var v3805 int32
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3807 float32
	_ = v3807
	var v3808 float32
	_ = v3808
	var v3810 float32
	_ = v3810
	var v3811 float32
	_ = v3811
	var v3814 float32
	_ = v3814
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3819 int32
	_ = v3819
	var v3826 int32
	_ = v3826
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3872 float32
	_ = v3872
	var v3873 float32
	_ = v3873
	var v3875 float32
	_ = v3875
	var v3876 float32
	_ = v3876
	var v3879 float32
	_ = v3879
	var v3928 int32
	_ = v3928
	var v3930 int32
	_ = v3930
	var v3939 int32
	_ = v3939
	var v3941 int32
	_ = v3941
	var v3980 int32
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3983 float32
	_ = v3983
	var v3985 int32
	_ = v3985
	var v3989 float32
	_ = v3989
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v3995 float32
	_ = v3995
	var v3997 int32
	_ = v3997
	var v4001 float32
	_ = v4001
	var v4005 int32
	_ = v4005
	var v4007 int32
	_ = v4007
	var v4014 int32
	_ = v4014
	var v4057 int32
	_ = v4057
	var v4058 int32
	_ = v4058
	var v4059 int32
	_ = v4059
	var v4060 float32
	_ = v4060
	var v4062 int32
	_ = v4062
	var v4066 float32
	_ = v4066
	var v4115 int32
	_ = v4115
	var v4123 int32
	_ = v4123
	var v4164 int32
	_ = v4164
	var v4166 int32
	_ = v4166
	var v4167 int32
	_ = v4167
	var v4171 int32
	_ = v4171
	var v4173 int32
	_ = v4173
	var v4178 int32
	_ = v4178
	var v4182 int32
	_ = v4182
	var v4187 int32
	_ = v4187
	var v4238 int32
	_ = v4238
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4291 int32
	_ = v4291
	var v4292 int32
	_ = v4292
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4295 int32
	_ = v4295
	var v4296 int32
	_ = v4296
	var v4299 int32
	_ = v4299
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4356 int32
	_ = v4356
	var v4357 int32
	_ = v4357
	var v4360 int32
	_ = v4360
	var v4362 int32
	_ = v4362
	var v4363 int32
	_ = v4363
	var v4364 int32
	_ = v4364
	var v4372 int32
	_ = v4372
	var v4416 float32
	_ = v4416
	var v4428 int32
	_ = v4428
	var v4432 int32
	_ = v4432
	var v4437 int32
	_ = v4437
	var v4439 int32
	_ = v4439
	var v4488 int32
	_ = v4488
	var v4489 int32
	_ = v4489
	var v4538 int32
	_ = v4538
	var v4539 int32
	_ = v4539
	var v4542 int32
	_ = v4542
	var v4545 int32
	_ = v4545
	var v4546 int32
	_ = v4546
	var v4553 int32
	_ = v4553
	var v4594 int32
	_ = v4594
	var v4596 int32
	_ = v4596
	var v4597 int32
	_ = v4597
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4602 float64
	_ = v4602
	var v4606 int32
	_ = v4606
	var v4607 int32
	_ = v4607
	var v4612 int32
	_ = v4612
	var v4616 int32
	_ = v4616
	var v4621 int32
	_ = v4621
	var v4671 int32
	_ = v4671
	var v4675 int32
	_ = v4675
	var v4679 int32
	_ = v4679
	var v4684 int32
	_ = v4684
	var v4688 int32
	_ = v4688
	var v4692 int32
	_ = v4692
	var v4697 int32
	_ = v4697
	var v4701 int32
	_ = v4701
	var v4705 int32
	_ = v4705
	var v4710 int32
	_ = v4710
	var v4714 int32
	_ = v4714
	var v4718 int32
	_ = v4718
	var v4723 int32
	_ = v4723
	var v4727 int32
	_ = v4727
	var v4731 int32
	_ = v4731
	var v4736 int32
	_ = v4736
	var v4740 int32
	_ = v4740
	var v4744 int32
	_ = v4744
	var v4749 int32
	_ = v4749
	var v4753 int32
	_ = v4753
	var v4757 int32
	_ = v4757
	var v4762 int32
	_ = v4762
	var v4766 int32
	_ = v4766
	var v4770 int32
	_ = v4770
	var v4775 int32
	_ = v4775
	var v4779 int32
	_ = v4779
	var v4783 int32
	_ = v4783
	var v4788 int32
	_ = v4788
	v6 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v53 = F_AllocSetContextCreateInternal(m, v48, int32(65362), v6, int32(8192), int32(8388608))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v55 = int32(4562096)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v53
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v60 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4779 = m.ExcPending
	if v4779 != 0 {
		goto L1
	} else {
		goto L519
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4766 = m.ExcPending
	if v4766 != 0 {
		goto L1
	} else {
		goto L516
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4753 = m.ExcPending
	if v4753 != 0 {
		goto L1
	} else {
		goto L513
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4740 = m.ExcPending
	if v4740 != 0 {
		goto L1
	} else {
		goto L510
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4727 = m.ExcPending
	if v4727 != 0 {
		goto L1
	} else {
		goto L507
	}
L8:
	;
	v4287 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v4288 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4287 == v4288 {
		goto L449
	} else {
		goto L450
	}
L9:
	;
	v64 = F_HnswOptionalProcInfo(m, l0, int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v347 = F_mul_size(m, v341, (v342+int32(7))&int32(-8))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L39
	}
L12:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v69 = F_mul_size(m, int32(4), v59)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v71 = F_palloc(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v73 < v74 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v81 = v73
	goto L18
L16:
	;
	goto L17
L17:
	;
	if v64 == int32(0) {
		goto L8
	} else {
		goto L35
	}
L18:
	;
	if int32(0) <= v81 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L17
L20:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	m.T0[v271].(func(*base.Module, int32, int32, int32))(m, v124+v125*v81, v59, v71)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L33
	}
L21:
	;
	v149 = v128
	goto L29
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v128 = int32(0)
	if v128 < v59 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L20
L26:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v193 = int32(4645600)
	v196 = *(*int64)(unsafe.Add(mBase, _consts[58]))
	v197 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	v198 = v196 ^ v197
	*(*int64)(unsafe.Add(mBase, _consts[59])) = base.I64_rotl(v198, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[58])) = v198<<(uint(int64(16))%64) ^ base.I64_rotl(v196, int64(24)) ^ v198
	v219 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v196*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L31
L30:
	;
	goto L20
L31:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v71+v149<<(uint(int32(2))%32)))) = base.F32_demote_f64(v219)
	v223 = v149 + int32(1)
	if v223 != v59 {
		v149 = v223
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v276 = v274 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v276 < v278 {
		v81 = v276
		goto L18
	} else {
		goto L34
	}
L34:
	;
	goto L19
L35:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v334 = F_AllocSetContextCreateInternal(m, v329, int32(65577), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_IvfflatNormVectors(m, l3, v67, l2, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_MemoryContextDelete(m, v334)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L8
L39:
	;
	v349 = F_add_size(m, int32(20), v347)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v352 = F_mul_size(m, v341, v59)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v354 = F_mul_size(m, int32(4), v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v357 = F_mul_size(m, int32(4), v341)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v360 = F_mul_size(m, int32(4), v60)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v363 = F_mul_size(m, v60, v341)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v365 = F_mul_size(m, int32(4), v363)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v368 = F_mul_size(m, int32(4), v60)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v371 = F_mul_size(m, int32(4), v341)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v374 = F_mul_size(m, v341, v341)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v376 = F_mul_size(m, int32(4), v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v379 = F_mul_size(m, int32(4), v341)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v381 = F_add_size(m, l4, v349)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v383 = F_add_size(m, v381, v354)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v385 = F_add_size(m, v383, v357)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v387 = F_add_size(m, v385, v360)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v389 = F_add_size(m, v387, v365)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v391 = F_add_size(m, v389, v368)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v393 = F_add_size(m, v391, v371)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v395 = F_add_size(m, v393, v376)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v397 = F_add_size(m, v395, v379)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_IvfflatCheckMemoryUsage(m, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v402 = base.I32_div_s(int32(2147483647), v341)
	if v402 < v341 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	v406 = F_index_getprocinfo(m, l0, int32(1), int32(3))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v409 = F_HnswOptionalProcInfo(m, l0, int32(4))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v413 = F_palloc(m, v354)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v415 = F_palloc(m, v357)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v417 = F_palloc(m, v360)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v420 = F_palloc_extended(m, v365, int32(1))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v422 = F_palloc(m, v368)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v424 = F_palloc(m, v371)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v427 = F_palloc_extended(m, v376, int32(1))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v429 = F_palloc(m, v379)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v432 = F_VectorArrayInit(m, v341, v59, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432))) = v341
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v437 = F_mul_size(m, int32(4), v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v439 = F_palloc(m, v437)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v445 = F_index_getprocinfo(m, l0, int32(1), int32(3))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v451 = int32(4645608)
	v452 = int32(4645600)
	v453 = *(*int64)(unsafe.Add(mBase, _consts[58]))
	v455 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	v456 = v453 ^ v455
	*(*int64)(unsafe.Add(mBase, _consts[59])) = base.I64_rotl(v456, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[58])) = v456<<(uint(int64(16))%64) ^ base.I64_rotl(v453, int64(24)) ^ v456
	goto L77
L77:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v478 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v453*int64(5), int64(7))*int64(9))>>(uint(int64(32))%64))), v477)
	if v478 < int32(0) {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v481 <= v478 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v483 = int32(0)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_VectorArraySet(m, l2, v483, v485+v486*v478)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v491 + int32(1)
	if v441 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v442 <= int32(0) {
		goto L93
	} else {
		goto L94
	}
L82:
	;
	v498 = v441 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v441) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v509 = v483
	v511 = int32(0)
	goto L86
L84:
	;
	v571 = v483
	goto L85
L85:
	;
	if v498 == int32(0) {
		goto L81
	} else {
		goto L89
	}
L86:
	;
	v552 = v439 + v509<<(uint(int32(2))%32)
	v553 = int64(9187343237679939583)
	*(*int64)(unsafe.Add(mBase, uint32(v552))) = v553
	*(*int64)(unsafe.Add(mBase, uint32(v552)+8)) = v553
	*(*int64)(unsafe.Add(mBase, uint32(v552)+16)) = v553
	*(*int64)(unsafe.Add(mBase, uint32(v552)+24)) = v553
	v561 = int32(8)
	v562 = v509 + v561
	v564 = v511 + v561
	if v564 != v441&int32(2147483640) {
		v509 = v562
		v511 = v564
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v571 = v562
	goto L85
L88:
	;
	goto L87
L89:
	;
	v620 = v571
	v621 = int32(0)
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439+v620<<(uint(int32(2))%32)))) = int32(2139095039)
	v666 = int32(1)
	v669 = v621 + v666
	if v669 != v498 {
		v620 = v620 + v666
		v621 = v669
		goto L90
	} else {
		goto L92
	}
L91:
	;
	goto L81
L92:
	;
	goto L91
L93:
	;
	F_pfree(m, v439)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L123
	}
L94:
	;
	v720 = v441 - int32(1)
	v721 = int32(0)
	v730 = v721
	goto L95
L95:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v771 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v774 = float64(0)
	if base.B2i32(v441 <= v721) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v786 = int32(0)
	v825 = v774
	goto L104
L102:
	;
	v908 = v774
	goto L103
L103:
	;
	v911 = v730 + int32(1)
	if v911 == v442 {
		goto L93
	} else {
		goto L113
	}
L104:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v827 <= v786 {
		goto L5
	} else {
		goto L106
	}
L105:
	;
	v908 = v860
	goto L103
L106:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v829 <= v730 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v843 = F_FunctionCall2Coll(m, v445, v448, v835+v836*v786, v839+v840*v730)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v845 = *(*float64)(unsafe.Add(mBase, uint32(v843)))
	*(*float32)(unsafe.Add(mBase, uint32(v420+v730<<(uint(int32(2))%32)+v786*v442<<(uint(int32(2))%32)))) = base.F32_demote_f64(v845)
	v848 = base.F64_mul(v845, v845)
	v851 = v439 + v786<<(uint(int32(2))%32)
	v852 = *(*float32)(unsafe.Add(mBase, uint32(v851)))
	v853 = base.F64_promote_f32(v852)
	if base.F64_lt(v848, v853) != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v855 = base.F32_demote_f64(v848)
	*(*float32)(unsafe.Add(mBase, uint32(v851))) = v855
	v859 = base.F64_promote_f32(v855)
	goto L111
L110:
	;
	v859 = v853
	goto L111
L111:
	;
	v860 = base.F64_add(v859, v825)
	v862 = v786 + int32(1)
	if v862 != v441 {
		v786 = v862
		v825 = v860
		goto L104
	} else {
		goto L112
	}
L112:
	;
	goto L105
L113:
	;
	v913 = int32(0)
	v914 = int32(4645600)
	v917 = *(*int64)(unsafe.Add(mBase, _consts[58]))
	v918 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	v919 = v917 ^ v918
	*(*int64)(unsafe.Add(mBase, _consts[59])) = base.I64_rotl(v919, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[58])) = v919<<(uint(int64(16))%64) ^ base.I64_rotl(v917, int64(24)) ^ v919
	v940 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v917*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L114
L114:
	;
	if v441 < int32(2) {
		v1006 = v913
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1047 <= v1006 {
		goto L7
	} else {
		goto L121
	}
L116:
	;
	v949 = v913
	v987 = base.F64_mul(v940, v908)
	goto L117
L117:
	;
	v993 = *(*float32)(unsafe.Add(mBase, uint32(v439+v949<<(uint(int32(2))%32))))
	v995 = base.F64_sub(v987, base.F64_promote_f32(v993))
	if base.F64_le(v995, float64(0)) != 0 {
		v1006 = v949
		goto L115
	} else {
		goto L119
	}
L118:
	;
	v1006 = v720
	goto L115
L119:
	;
	v999 = v949 + int32(1)
	if v999 != v720 {
		v949 = v999
		v987 = v995
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_VectorArraySet(m, l2, v911, v1049+v1050*v1006)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1055 + int32(1)
	v730 = v911
	goto L95
L123:
	;
	if int32(0) < v60 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v1112 = v341 & int32(3)
	v1139 = v6
	goto L127
L125:
	;
	goto L126
L126:
	;
	v1468 = int32(2147483646)
	v1470 = int32(1)
	v1473 = v341 & v1468
	v1475 = v341 & v1470
	v1477 = v341 - v1470
	v1505 = int32(0)
	goto L174
L127:
	;
	if v341 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L126
L129:
	;
	v1414 = v1139 << (uint(int32(2)) % 32)
	*(*float32)(unsafe.Add(mBase, uint32(v422+v1414))) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v1414+v417))) = v1374
	v1420 = v1139 + int32(1)
	if v1420 != v60 {
		v1139 = v1420
		goto L127
	} else {
		goto L173
	}
L130:
	;
	v1374 = int32(0)
	v1405 = float32(3.4028235e+38)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v1168 = v420 + v341*v1139<<(uint(int32(2))%32)
	v1169 = int32(0)
	v1170 = float32(3.4028235e+38)
	if base.B2i32(base.Ui32(v341) < base.Ui32(int32(4))) == v1169 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v1181 = v1169
	v1183 = v1169
	v1185 = v1169
	v1214 = v1170
	goto L136
L134:
	;
	v1266 = v1169
	v1268 = v1169
	v1299 = v1170
	goto L135
L135:
	;
	if v1112 == int32(0) {
		v1374 = v1268
		v1405 = v1299
		goto L129
	} else {
		goto L163
	}
L136:
	;
	v1223 = v1181 | int32(3)
	v1224 = int32(2)
	v1227 = *(*float32)(unsafe.Add(mBase, uint32(v1168+v1223<<(uint(v1224)%32))))
	v1229 = v1181 | v1224
	v1233 = *(*float32)(unsafe.Add(mBase, uint32(v1168+v1229<<(uint(v1224)%32))))
	v1235 = v1181 | int32(1)
	v1239 = *(*float32)(unsafe.Add(mBase, uint32(v1168+v1235<<(uint(v1224)%32))))
	v1243 = *(*float32)(unsafe.Add(mBase, uint32(v1168+v1181<<(uint(v1224)%32))))
	v1244 = base.F32_gt(v1214, v1243)
	if v1244 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v1266 = v1257
	v1268 = v1255
	v1299 = v1251
	goto L135
L138:
	;
	v1245 = v1243
	goto L140
L139:
	;
	v1245 = v1214
	goto L140
L140:
	;
	v1246 = base.F32_gt(v1245, v1239)
	if v1246 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v1247 = v1239
	goto L143
L142:
	;
	v1247 = v1245
	goto L143
L143:
	;
	v1248 = base.F32_gt(v1247, v1233)
	if v1248 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v1249 = v1233
	goto L146
L145:
	;
	v1249 = v1247
	goto L146
L146:
	;
	v1250 = base.F32_gt(v1249, v1227)
	if v1250 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v1251 = v1227
	goto L149
L148:
	;
	v1251 = v1249
	goto L149
L149:
	;
	if v1244 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v1252 = v1181
	goto L152
L151:
	;
	v1252 = v1183
	goto L152
L152:
	;
	if v1246 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1253 = v1235
	goto L155
L154:
	;
	v1253 = v1252
	goto L155
L155:
	;
	if v1248 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v1254 = v1229
	goto L158
L157:
	;
	v1254 = v1253
	goto L158
L158:
	;
	if v1250 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1255 = v1223
	goto L161
L160:
	;
	v1255 = v1254
	goto L161
L161:
	;
	v1256 = int32(4)
	v1257 = v1181 + v1256
	v1259 = v1185 + v1256
	if v1259 != v341&int32(2147483644) {
		v1181 = v1257
		v1183 = v1255
		v1185 = v1259
		v1214 = v1251
		goto L136
	} else {
		goto L162
	}
L162:
	;
	goto L137
L163:
	;
	v1314 = v1266
	v1316 = v1268
	v1324 = v1169
	v1347 = v1299
	goto L164
L164:
	;
	v1358 = *(*float32)(unsafe.Add(mBase, uint32(v1168+v1314<<(uint(int32(2))%32))))
	v1359 = base.F32_gt(v1347, v1358)
	if v1359 != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v1374 = v1361
	v1405 = v1360
	goto L129
L166:
	;
	v1360 = v1358
	goto L168
L167:
	;
	v1360 = v1347
	goto L168
L168:
	;
	if v1359 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1361 = v1314
	goto L171
L170:
	;
	v1361 = v1316
	goto L171
L171:
	;
	v1362 = int32(1)
	v1365 = v1324 + v1362
	if v1365 != v1112 {
		v1314 = v1314 + v1362
		v1316 = v1361
		v1324 = v1365
		v1347 = v1360
		goto L164
	} else {
		goto L172
	}
L172:
	;
	goto L165
L173:
	;
	goto L128
L174:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v1526 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	goto L8
L176:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v1529 = int32(0)
	if v1529 < v341 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	v1542 = v1529
	goto L183
L181:
	;
	goto L182
L182:
	;
	v1972 = int32(0)
	if v1972 < v60 {
		goto L234
	} else {
		goto L235
	}
L183:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1542 < v1578 {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	v1744 = int32(0)
	goto L203
L185:
	;
	if v341 != v1581 {
		v1542 = v1581
		goto L183
	} else {
		goto L202
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L1
	} else {
		goto L199
	}
L187:
	;
	v1581 = v1542 + int32(1)
	if v341 <= v1581 {
		goto L185
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L1
	} else {
		goto L196
	}
L190:
	;
	v1583 = int32(2)
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1599 = v1581
	goto L191
L191:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1640 <= v1599 {
		goto L186
	} else {
		goto L193
	}
L192:
	;
	goto L185
L193:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1649 = F_FunctionCall2Coll(m, v406, v412, v1586+v1587*v1542, v1645+v1646*v1599)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v1651 = *(*float64)(unsafe.Add(mBase, uint32(v1649)))
	v1654 = base.F32_demote_f64(base.F64_mul(v1651, float64(0.5)))
	*(*float32)(unsafe.Add(mBase, uint32(v427+v1542*v341<<(uint(v1583)%32)+v1599<<(uint(int32(2))%32)))) = v1654
	*(*float32)(unsafe.Add(mBase, uint32(v427+v1542<<(uint(v1583)%32)+v1599*v341<<(uint(int32(2))%32)))) = v1654
	v1662 = v1599 + int32(1)
	if v341 != v1662 {
		v1599 = v1662
		goto L191
	} else {
		goto L195
	}
L195:
	;
	goto L192
L196:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	goto L184
L203:
	;
	v1787 = v427 + v1744*v341<<(uint(int32(2))%32)
	v1788 = float32(3.4028235e+38)
	v1789 = int32(0)
	if v1477 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L182
L205:
	;
	v1796 = v1789
	v1798 = v1789
	v1829 = v1788
	goto L208
L206:
	;
	v1867 = v1789
	v1900 = v1788
	goto L207
L207:
	;
	if v1475 == int32(0) {
		v1917 = v1900
		goto L223
	} else {
		goto L224
	}
L208:
	;
	if v1796 != v1744 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v1867 = v1858
	v1900 = v1855
	goto L207
L210:
	;
	v1841 = *(*float32)(unsafe.Add(mBase, uint32(v1787+v1796<<(uint(int32(2))%32))))
	if base.F32_gt(v1829, v1841) != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1844 = v1829
	goto L212
L212:
	;
	v1847 = v1796 | int32(1)
	if v1847 != v1744 {
		goto L216
	} else {
		goto L217
	}
L213:
	;
	v1843 = v1841
	goto L215
L214:
	;
	v1843 = v1829
	goto L215
L215:
	;
	v1844 = v1843
	goto L212
L216:
	;
	v1852 = *(*float32)(unsafe.Add(mBase, uint32(v1787+v1847<<(uint(int32(2))%32))))
	if base.F32_gt(v1844, v1852) != 0 {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v1855 = v1844
	goto L218
L218:
	;
	v1857 = int32(2)
	v1858 = v1796 + v1857
	v1860 = v1798 + v1857
	if v1860 != v1473 {
		v1796 = v1858
		v1798 = v1860
		v1829 = v1855
		goto L208
	} else {
		goto L222
	}
L219:
	;
	v1854 = v1852
	goto L221
L220:
	;
	v1854 = v1844
	goto L221
L221:
	;
	v1855 = v1854
	goto L218
L222:
	;
	goto L209
L223:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v424+v1744<<(uint(int32(2))%32)))) = v1917
	v1924 = v1744 + int32(1)
	if v1924 != v341 {
		v1744 = v1924
		goto L203
	} else {
		goto L229
	}
L224:
	;
	if v1867 == v1744 {
		v1917 = v1900
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1914 = *(*float32)(unsafe.Add(mBase, uint32(v1787+v1867<<(uint(int32(2))%32))))
	if base.F32_gt(v1900, v1914) != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1916 = v1914
	goto L228
L227:
	;
	v1916 = v1900
	goto L228
L228:
	;
	v1917 = v1916
	goto L223
L229:
	;
	goto L204
L230:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	v2318 = int32(0)
	v2319 = base.B2i32(v2317 <= v2318)
	if v2319 == v2318 {
		goto L273
	} else {
		goto L274
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L1
	} else {
		goto L270
	}
L232:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L1
	} else {
		goto L267
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L1
	} else {
		goto L264
	}
L234:
	;
	v1975 = int32(0)
	v1990 = v1972
	v1998 = v1975
	goto L237
L235:
	;
	goto L236
L236:
	;
	v2291 = int32(1)
	goto L230
L237:
	;
	if v341 <= int32(0) {
		v2198 = v1998
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v2291 = base.B2i32(v2198 == int32(0))
	goto L230
L239:
	;
	v2225 = v1990 + int32(1)
	if v2225 != v60 {
		v1990 = v2225
		v1998 = v2198
		goto L237
	} else {
		goto L263
	}
L240:
	;
	v2026 = int32(2)
	v2027 = v1990 << (uint(v2026) % 32)
	v2028 = v422 + v2027
	v2029 = *(*float32)(unsafe.Add(mBase, uint32(v2028)))
	v2030 = v2027 + v417
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2030)))
	v2035 = *(*float32)(unsafe.Add(mBase, uint32(v424+v2031<<(uint(v2026)%32))))
	if base.F32_le(v2029, v2035) != 0 {
		v2198 = v1998
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v2040 = v420 + v341*v1990<<(uint(int32(2))%32)
	v2047 = int32(0)
	v2056 = base.B2i32(v1505 != v1975)
	v2062 = v1998
	goto L242
L242:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v2030)))
	if v2047 == v2088 {
		v2170 = v2056
		v2172 = v2062
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v2198 = v2172
	goto L239
L244:
	;
	v2176 = v2047 + int32(1)
	if v2176 != v341 {
		v2047 = v2176
		v2056 = v2170
		v2062 = v2172
		goto L242
	} else {
		goto L262
	}
L245:
	;
	v2090 = *(*float32)(unsafe.Add(mBase, uint32(v2028)))
	v2092 = v2047 << (uint(int32(2)) % 32)
	v2093 = v2040 + v2092
	v2094 = *(*float32)(unsafe.Add(mBase, uint32(v2093)))
	if base.F32_le(v2090, v2094) != 0 {
		v2170 = v2056
		v2172 = v2062
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v2101 = *(*float32)(unsafe.Add(mBase, uint32(v427+v2088*v341<<(uint(int32(2))%32)+v2092)))
	if base.F32_le(v2090, v2101) != 0 {
		v2170 = v2056
		v2172 = v2062
		goto L244
	} else {
		goto L247
	}
L247:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2103 <= v1990 {
		goto L233
	} else {
		goto L248
	}
L248:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2108 = v2105 + v2106*v1990
	if v2056&int32(1) != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	if v2088 < int32(0) {
		goto L232
	} else {
		goto L252
	}
L250:
	;
	v2130 = v2088
	v2132 = v2090
	v2133 = v2094
	goto L251
L251:
	;
	if base.F32_gt(v2132, v2133) == int32(0) {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2113 <= v2088 {
		goto L232
	} else {
		goto L253
	}
L253:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2119 = F_FunctionCall2Coll(m, v406, v412, v2108, v2115+v2116*v2088)
	mBase = m.M
	v2120 = m.ExcPending
	if v2120 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2030)))
	v2125 = *(*float64)(unsafe.Add(mBase, uint32(v2119)))
	v2126 = base.F32_demote_f64(v2125)
	*(*float32)(unsafe.Add(mBase, uint32(v2040+v2121<<(uint(int32(2))%32)))) = v2126
	*(*float32)(unsafe.Add(mBase, uint32(v2028))) = v2126
	v2129 = *(*float32)(unsafe.Add(mBase, uint32(v2093)))
	v2130 = v2121
	v2132 = v2126
	v2133 = v2129
	goto L251
L255:
	;
	v2137 = int32(0)
	v2143 = *(*float32)(unsafe.Add(mBase, uint32(v427+v2130*v341<<(uint(int32(2))%32)+v2092)))
	if base.F32_gt(v2132, v2143) == v2137 {
		v2170 = v2137
		v2172 = v2062
		goto L244
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2148 <= v2047 {
		goto L231
	} else {
		goto L259
	}
L258:
	;
	goto L257
L259:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2154 = F_FunctionCall2Coll(m, v406, v412, v2108, v2150+v2151*v2047)
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v2156 = *(*float64)(unsafe.Add(mBase, uint32(v2154)))
	v2157 = base.F32_demote_f64(v2156)
	*(*float32)(unsafe.Add(mBase, uint32(v2093))) = v2157
	v2159 = int32(0)
	if base.F32_gt(v2132, v2157) == v2159 {
		v2170 = v2159
		v2172 = v2062
		goto L244
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2030))) = v2047
	*(*float32)(unsafe.Add(mBase, uint32(v2028))) = v2157
	v2170 = v2159
	v2172 = v2062 + int32(1)
	goto L244
L262:
	;
	goto L243
L263:
	;
	goto L238
L264:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	v2322 = int32(1)
	v2325 = v2315 << (uint(int32(2)) % 32)
	v2326 = int32(0)
	if v2317 != v2322 {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v2476 = v2316
	goto L275
L275:
	;
	v2516 = int32(0)
	if v2516 < v2476 {
		goto L299
	} else {
		goto L300
	}
L276:
	;
	v2337 = v2326
	v2341 = int32(0)
	goto L279
L277:
	;
	v2416 = v2326
	goto L278
L278:
	;
	if v2317&v2322 != 0 {
		goto L290
	} else {
		goto L291
	}
L279:
	;
	v2378 = int32(0)
	v2379 = base.B2i32(v2315 <= v2378)
	if v2379 == v2378 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v2416 = v2407
	goto L278
L281:
	;
	v2386 = F__emscripten_memset_bulkmem(m, v413+v2337*v2325, base.I32_extend8_s(int32(0)), v2325)
	mBase = m.M
	goto L284
L282:
	;
	goto L283
L283:
	;
	v2390 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v415+v2337<<(uint(int32(2))%32)))) = v2390
	v2393 = v2337 | int32(1)
	if v2379 == v2390 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	goto L283
L285:
	;
	v2400 = F__emscripten_memset_bulkmem(m, v413+v2393*v2325, base.I32_extend8_s(int32(0)), v2325)
	mBase = m.M
	goto L288
L286:
	;
	goto L287
L287:
	;
	v2401 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v415+v2393<<(uint(v2401)%32)))) = int32(0)
	v2407 = v2337 + v2401
	v2409 = v2341 + v2401
	if v2409 != v2317&int32(2147483646) {
		v2337 = v2407
		v2341 = v2409
		goto L279
	} else {
		goto L289
	}
L288:
	;
	goto L287
L289:
	;
	goto L280
L290:
	;
	if int32(0) < v2315 {
		goto L293
	} else {
		goto L294
	}
L291:
	;
	goto L292
L292:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2476 = v2469
	goto L275
L293:
	;
	v2463 = F__emscripten_memset_bulkmem(m, v413+v2416*v2325, base.I32_extend8_s(int32(0)), v2325)
	mBase = m.M
	goto L296
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415+v2416<<(uint(int32(2))%32)))) = int32(0)
	goto L292
L296:
	;
	goto L295
L297:
	;
	v3689 = int32(0)
	if v60 <= v3689 {
		goto L402
	} else {
		goto L403
	}
L298:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3633 = m.ExcPending
	if v3633 != 0 {
		goto L1
	} else {
		goto L399
	}
L299:
	;
	v2524 = v2516
	goto L302
L300:
	;
	goto L301
L301:
	;
	if v2316 <= int32(0) {
		goto L307
	} else {
		goto L308
	}
L302:
	;
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2565 <= v2524 {
		goto L298
	} else {
		goto L304
	}
L303:
	;
	goto L301
L304:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2572 = int32(2)
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v417+v2524<<(uint(v2572)%32))))
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	m.T0[v2580].(func(*base.Module, int32, int32))(m, v2567+v2568*v2524, v413+v2571*v2575<<(uint(v2572)%32))
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v2584 = v2524 + int32(1)
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2584 < v2585 {
		v2524 = v2584
		goto L302
	} else {
		goto L306
	}
L306:
	;
	goto L303
L307:
	;
	if v2319 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L308:
	;
	v2636 = v2316 & int32(3)
	v2637 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v2316) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v2649 = v2637
	v2653 = int32(0)
	goto L312
L310:
	;
	v2735 = v2637
	goto L311
L311:
	;
	if v2636 == int32(0) {
		goto L307
	} else {
		goto L315
	}
L312:
	;
	v2690 = int32(2)
	v2692 = v417 + v2649<<(uint(v2690)%32)
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(v2692)))
	v2696 = v415 + v2693<<(uint(v2690)%32)
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v2696)))
	v2698 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2696))) = v2697 + v2698
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(v2692)+4))
	v2704 = v415 + v2701<<(uint(v2690)%32)
	v2705 = *(*int32)(unsafe.Add(mBase, uint32(v2704)))
	*(*int32)(unsafe.Add(mBase, uint32(v2704))) = v2705 + v2698
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2692)+8))
	v2712 = v415 + v2709<<(uint(v2690)%32)
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v2712)))
	*(*int32)(unsafe.Add(mBase, uint32(v2712))) = v2713 + v2698
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2692)+12))
	v2720 = v415 + v2717<<(uint(v2690)%32)
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v2720)))
	*(*int32)(unsafe.Add(mBase, uint32(v2720))) = v2721 + v2698
	v2725 = int32(4)
	v2726 = v2649 + v2725
	v2728 = v2653 + v2725
	if v2728 != v2316&int32(2147483644) {
		v2649 = v2726
		v2653 = v2728
		goto L312
	} else {
		goto L314
	}
L313:
	;
	v2735 = v2726
	goto L311
L314:
	;
	goto L313
L315:
	;
	v2783 = v2735
	v2785 = v2637
	goto L316
L316:
	;
	v2824 = int32(2)
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v417+v2783<<(uint(v2824)%32))))
	v2830 = v415 + v2827<<(uint(v2824)%32)
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2830)))
	v2832 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2830))) = v2831 + v2832
	v2838 = v2785 + v2832
	if v2838 != v2636 {
		v2783 = v2783 + v2832
		v2785 = v2838
		goto L316
	} else {
		goto L318
	}
L317:
	;
	goto L307
L318:
	;
	goto L317
L319:
	;
	v2889 = v2315 & int32(2147483646)
	v2890 = int32(1)
	v2891 = v2315 & v2890
	v2893 = v2315 - v2890
	v2911 = int32(0)
	goto L322
L320:
	;
	goto L321
L321:
	;
	v3392 = int32(0)
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	if v3392 < v3393 {
		goto L369
	} else {
		goto L370
	}
L322:
	;
	v2942 = int32(2)
	v2944 = v413 + v2315*v2911<<(uint(v2942)%32)
	v2948 = *(*int32)(unsafe.Add(mBase, uint32(v415+v2911<<(uint(v2942)%32))))
	if v2948 <= int32(0) {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	goto L321
L324:
	;
	v3344 = v2911 + int32(1)
	if v3344 != v2317 {
		v2911 = v3344
		goto L322
	} else {
		goto L365
	}
L325:
	;
	v2951 = int32(0)
	if v2315 <= v2951 {
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	if v2315 <= int32(0) {
		goto L324
	} else {
		goto L333
	}
L328:
	;
	v2960 = v2951
	goto L329
L329:
	;
	v3003 = int32(4645600)
	v3006 = *(*int64)(unsafe.Add(mBase, _consts[58]))
	v3007 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	v3008 = v3006 ^ v3007
	*(*int64)(unsafe.Add(mBase, _consts[59])) = base.I64_rotl(v3008, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[58])) = v3008<<(uint(int64(16))%64) ^ base.I64_rotl(v3006, int64(24)) ^ v3008
	v3029 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v3006*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L331
L330:
	;
	goto L324
L331:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v2944+v2960<<(uint(int32(2))%32)))) = base.F32_demote_f64(v3029)
	v3033 = v2960 + int32(1)
	if v3033 != v2315 {
		v2960 = v3033
		goto L329
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	v3037 = int32(0)
	if v2893 != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v3045 = v3037
	v3048 = v3037
	goto L337
L335:
	;
	v3121 = v3037
	goto L336
L336:
	;
	if v2891 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L337:
	;
	v3087 = v2944 + v3045<<(uint(int32(2))%32)
	v3088 = *(*float32)(unsafe.Add(mBase, uint32(v3087)))
	if base.F32_eq(base.F32_abs(v3088), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	v3121 = v3111
	goto L336
L339:
	;
	if base.F32_gt(v3088, float32(0)) != 0 {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	goto L341
L341:
	;
	v3099 = v3087 + int32(4)
	v3100 = *(*float32)(unsafe.Add(mBase, uint32(v3099)))
	if base.F32_eq(base.F32_abs(v3100), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L345
	} else {
		goto L346
	}
L342:
	;
	v3096 = float32(3.4028235e+38)
	goto L344
L343:
	;
	v3096 = float32(-3.4028235e+38)
	goto L344
L344:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3087))) = v3096
	goto L341
L345:
	;
	if base.F32_gt(v3100, float32(0)) != 0 {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	goto L347
L347:
	;
	v3110 = int32(2)
	v3111 = v3045 + v3110
	v3113 = v3048 + v3110
	if v3113 != v2889 {
		v3045 = v3111
		v3048 = v3113
		goto L337
	} else {
		goto L351
	}
L348:
	;
	v3108 = float32(3.4028235e+38)
	goto L350
L349:
	;
	v3108 = float32(-3.4028235e+38)
	goto L350
L350:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3099))) = v3108
	goto L347
L351:
	;
	goto L338
L352:
	;
	v3178 = base.F32_convert_i32_s(v2948)
	v3179 = int32(0)
	if v2893 != 0 {
		goto L358
	} else {
		goto L359
	}
L353:
	;
	v3165 = v2944 + v3121<<(uint(int32(2))%32)
	v3166 = *(*float32)(unsafe.Add(mBase, uint32(v3165)))
	if base.F32_ne(base.F32_abs(v3166), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L352
	} else {
		goto L354
	}
L354:
	;
	if base.F32_gt(v3166, float32(0)) != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v3174 = float32(3.4028235e+38)
	goto L357
L356:
	;
	v3174 = float32(-3.4028235e+38)
	goto L357
L357:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3165))) = v3174
	goto L352
L358:
	;
	v3187 = v3179
	v3190 = v3179
	goto L361
L359:
	;
	v3249 = v3179
	goto L360
L360:
	;
	if v2891 == int32(0) {
		goto L324
	} else {
		goto L364
	}
L361:
	;
	v3227 = int32(2)
	v3229 = v2944 + v3187<<(uint(v3227)%32)
	v3230 = *(*float32)(unsafe.Add(mBase, uint32(v3229)))
	*(*float32)(unsafe.Add(mBase, uint32(v3229))) = base.F32_div(v3230, v3178)
	v3234 = v3229 + int32(4)
	v3235 = *(*float32)(unsafe.Add(mBase, uint32(v3234)))
	*(*float32)(unsafe.Add(mBase, uint32(v3234))) = base.F32_div(v3235, v3178)
	v3239 = v3187 + v3227
	v3241 = v3190 + v3227
	if v3241 != v2889 {
		v3187 = v3239
		v3190 = v3241
		goto L361
	} else {
		goto L363
	}
L362:
	;
	v3249 = v3239
	goto L360
L363:
	;
	goto L362
L364:
	;
	v3293 = v2944 + v3249<<(uint(int32(2))%32)
	v3294 = *(*float32)(unsafe.Add(mBase, uint32(v3293)))
	*(*float32)(unsafe.Add(mBase, uint32(v3293))) = base.F32_div(v3294, v3178)
	goto L324
L365:
	;
	goto L323
L366:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3620 = m.ExcPending
	if v3620 != 0 {
		goto L1
	} else {
		goto L396
	}
L367:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3607 = m.ExcPending
	if v3607 != 0 {
		goto L1
	} else {
		goto L393
	}
L368:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L1
	} else {
		goto L390
	}
L369:
	;
	v3401 = v3392
	goto L372
L370:
	;
	goto L371
L371:
	;
	if v409 != 0 {
		goto L377
	} else {
		goto L378
	}
L372:
	;
	v3442 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v3442 <= v3401 {
		goto L368
	} else {
		goto L374
	}
L373:
	;
	goto L371
L374:
	;
	v3444 = *(*int32)(unsafe.Add(mBase, uint32(v432)+16))
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v3448 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	m.T0[v3453].(func(*base.Module, int32, int32, int32))(m, v3444+v3445*v3401, v3448, v413+v3401*v3448<<(uint(int32(2))%32))
	mBase = m.M
	v3455 = m.ExcPending
	if v3455 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v3457 = v3401 + int32(1)
	v3458 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	if v3457 < v3458 {
		v3401 = v3457
		goto L372
	} else {
		goto L376
	}
L376:
	;
	goto L373
L377:
	;
	v3507 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v3512 = F_AllocSetContextCreateInternal(m, v3507, int32(65577), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v3513 = m.ExcPending
	if v3513 != 0 {
		goto L1
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v3519 = int32(0)
	if v341 <= v3519 {
		goto L297
	} else {
		goto L383
	}
L380:
	;
	F_IvfflatNormVectors(m, l3, v412, v432, v3512)
	mBase = m.M
	v3515 = m.ExcPending
	if v3515 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	F_MemoryContextDelete(m, v3512)
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	goto L379
L383:
	;
	v3527 = v3519
	goto L384
L384:
	;
	v3568 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3568 <= v3527 {
		goto L367
	} else {
		goto L386
	}
L385:
	;
	goto L297
L386:
	;
	v3570 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v3570 <= v3527 {
		goto L366
	} else {
		goto L387
	}
L387:
	;
	v3575 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3579 = *(*int32)(unsafe.Add(mBase, uint32(v432)+16))
	v3580 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v3583 = F_FunctionCall2Coll(m, v406, v412, v3575+v3576*v3527, v3579+v3580*v3527)
	mBase = m.M
	v3584 = m.ExcPending
	if v3584 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	v3585 = *(*float64)(unsafe.Add(mBase, uint32(v3583)))
	*(*float32)(unsafe.Add(mBase, uint32(v429+v3527<<(uint(int32(2))%32)))) = base.F32_demote_f64(v3585)
	v3589 = v3527 + int32(1)
	if v341 != v3589 {
		v3527 = v3589
		goto L384
	} else {
		goto L389
	}
L389:
	;
	goto L385
L390:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v3598 = m.ExcPending
	if v3598 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v3603 = m.ExcPending
	if v3603 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v3616 = m.ExcPending
	if v3616 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L396:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v3624 = m.ExcPending
	if v3624 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L399:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v3637 = m.ExcPending
	if v3637 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v3642 = m.ExcPending
	if v3642 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	v4115 = int32(0)
	if v341 <= v4115 {
		goto L432
	} else {
		goto L433
	}
L403:
	;
	v3703 = v3689
	goto L404
L404:
	;
	if v341 <= int32(0) {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	v3930 = int32(0)
	if v60 != int32(1) {
		goto L425
	} else {
		goto L426
	}
L406:
	;
	v3928 = v3703 + int32(1)
	if v3928 != v60 {
		v3703 = v3928
		goto L404
	} else {
		goto L424
	}
L407:
	;
	v3743 = v420 + v341*v3703<<(uint(int32(2))%32)
	v3744 = int32(0)
	if v1477 != 0 {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v3751 = v3744
	v3755 = v3744
	goto L411
L409:
	;
	v3826 = v3744
	goto L410
L410:
	;
	if v1475 == int32(0) {
		goto L406
	} else {
		goto L420
	}
L411:
	;
	v3793 = v3751 << (uint(int32(2)) % 32)
	v3794 = v3743 + v3793
	v3795 = float32(0)
	v3796 = *(*float32)(unsafe.Add(mBase, uint32(v3794)))
	v3798 = *(*float32)(unsafe.Add(mBase, uint32(v3793+v429)))
	v3799 = base.F32_sub(v3796, v3798)
	if base.F32_lt(v3799, v3795) != 0 {
		goto L413
	} else {
		goto L414
	}
L412:
	;
	v3826 = v3817
	goto L410
L413:
	;
	v3802 = v3795
	goto L415
L414:
	;
	v3802 = v3799
	goto L415
L415:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3794))) = v3802
	v3805 = v3793 | int32(4)
	v3806 = v3743 + v3805
	v3807 = float32(0)
	v3808 = *(*float32)(unsafe.Add(mBase, uint32(v3806)))
	v3810 = *(*float32)(unsafe.Add(mBase, uint32(v3805+v429)))
	v3811 = base.F32_sub(v3808, v3810)
	if base.F32_lt(v3811, v3807) != 0 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v3814 = v3807
	goto L418
L417:
	;
	v3814 = v3811
	goto L418
L418:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3806))) = v3814
	v3816 = int32(2)
	v3817 = v3751 + v3816
	v3819 = v3755 + v3816
	if v3819 != v1473 {
		v3751 = v3817
		v3755 = v3819
		goto L411
	} else {
		goto L419
	}
L419:
	;
	goto L412
L420:
	;
	v3870 = v3826 << (uint(int32(2)) % 32)
	v3871 = v3743 + v3870
	v3872 = float32(0)
	v3873 = *(*float32)(unsafe.Add(mBase, uint32(v3871)))
	v3875 = *(*float32)(unsafe.Add(mBase, uint32(v3870+v429)))
	v3876 = base.F32_sub(v3873, v3875)
	if base.F32_lt(v3876, v3872) != 0 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v3879 = v3872
	goto L423
L422:
	;
	v3879 = v3876
	goto L423
L423:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3871))) = v3879
	goto L406
L424:
	;
	goto L405
L425:
	;
	v3939 = v3930
	v3941 = v3930
	goto L428
L426:
	;
	v4014 = v3930
	goto L427
L427:
	;
	if v60&v1470 == int32(0) {
		goto L402
	} else {
		goto L431
	}
L428:
	;
	v3980 = int32(2)
	v3981 = v3939 << (uint(v3980) % 32)
	v3982 = v422 + v3981
	v3983 = *(*float32)(unsafe.Add(mBase, uint32(v3982)))
	v3985 = *(*int32)(unsafe.Add(mBase, uint32(v3981+v417)))
	v3989 = *(*float32)(unsafe.Add(mBase, uint32(v429+v3985<<(uint(v3980)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v3982))) = base.F32_add(v3983, v3989)
	v3993 = v3981 | int32(4)
	v3994 = v422 + v3993
	v3995 = *(*float32)(unsafe.Add(mBase, uint32(v3994)))
	v3997 = *(*int32)(unsafe.Add(mBase, uint32(v3993+v417)))
	v4001 = *(*float32)(unsafe.Add(mBase, uint32(v429+v3997<<(uint(v3980)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v3994))) = base.F32_add(v3995, v4001)
	v4005 = v3939 + v3980
	v4007 = v3941 + v3980
	if v4007 != v60&v1468 {
		v3939 = v4005
		v3941 = v4007
		goto L428
	} else {
		goto L430
	}
L429:
	;
	v4014 = v4005
	goto L427
L430:
	;
	goto L429
L431:
	;
	v4057 = int32(2)
	v4058 = v4014 << (uint(v4057) % 32)
	v4059 = v422 + v4058
	v4060 = *(*float32)(unsafe.Add(mBase, uint32(v4059)))
	v4062 = *(*int32)(unsafe.Add(mBase, uint32(v4058+v417)))
	v4066 = *(*float32)(unsafe.Add(mBase, uint32(v429+v4062<<(uint(v4057)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v4059))) = base.F32_add(v4060, v4066)
	goto L402
L432:
	;
	if base.B2i32(v1505 != int32(0))&v2291 != 0 {
		goto L8
	} else {
		goto L444
	}
L433:
	;
	v4123 = v4115
	goto L434
L434:
	;
	v4164 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v4123 < v4164 {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
		goto L1
	} else {
		goto L441
	}
L436:
	;
	v4166 = *(*int32)(unsafe.Add(mBase, uint32(v432)+16))
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	F_VectorArraySet(m, l2, v4123, v4166+v4167*v4123)
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L1
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	goto L435
L439:
	;
	v4173 = v4123 + int32(1)
	if v341 != v4173 {
		v4123 = v4173
		goto L434
	} else {
		goto L440
	}
L440:
	;
	goto L432
L441:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v4182 = m.ExcPending
	if v4182 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L444:
	;
	v4238 = v1505 + int32(1)
	if v4238 != int32(500) {
		v1505 = v4238
		goto L174
	} else {
		goto L445
	}
L445:
	;
	goto L175
L446:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4714 = m.ExcPending
	if v4714 != 0 {
		goto L1
	} else {
		goto L504
	}
L447:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4701 = m.ExcPending
	if v4701 != 0 {
		goto L1
	} else {
		goto L501
	}
L448:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4688 = m.ExcPending
	if v4688 != 0 {
		goto L1
	} else {
		goto L498
	}
L449:
	;
	v4291 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4292 = F_mul_size(m, int32(4), v4291)
	mBase = m.M
	v4293 = m.ExcPending
	if v4293 != 0 {
		goto L1
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4675 = m.ExcPending
	if v4675 != 0 {
		goto L1
	} else {
		goto L495
	}
L452:
	;
	v4294 = F_palloc(m, v4292)
	mBase = m.M
	v4295 = m.ExcPending
	if v4295 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v4296 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v4296 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v4299 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4308 = int32(0)
	v4309 = v4299
	goto L457
L455:
	;
	goto L456
L456:
	;
	v4538 = F_HnswOptionalProcInfo(m, l0, int32(2))
	mBase = m.M
	v4539 = m.ExcPending
	if v4539 != 0 {
		goto L1
	} else {
		goto L480
	}
L457:
	;
	if int32(0) < v4309 {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	goto L456
L459:
	;
	v4353 = F__emscripten_memset_bulkmem(m, v4294, base.I32_extend8_s(int32(0)), v4309<<(uint(int32(2))%32))
	mBase = m.M
	goto L462
L460:
	;
	goto L461
L461:
	;
	v4354 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4354 <= v4308 {
		goto L448
	} else {
		goto L463
	}
L462:
	;
	goto L461
L463:
	;
	v4356 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v4357 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v4360 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	m.T0[v4360].(func(*base.Module, int32, int32))(m, v4356+v4357*v4308, v4294)
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	v4363 = int32(0)
	v4364 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4363 < v4364 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v4372 = v4363
	goto L468
L466:
	;
	goto L467
L467:
	;
	v4488 = v4308 + int32(1)
	v4489 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4488 < v4489 {
		v4308 = v4488
		v4309 = v4364
		goto L457
	} else {
		goto L478
	}
L468:
	;
	v4416 = *(*float32)(unsafe.Add(mBase, uint32(v4294+v4372<<(uint(int32(2))%32))))
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v4416)&int32(2147483647)) {
		goto L447
	} else {
		goto L470
	}
L469:
	;
	goto L467
L470:
	;
	if base.F32_eq(base.F32_abs(v4416), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4428 = m.ExcPending
	if v4428 != 0 {
		goto L1
	} else {
		goto L474
	}
L472:
	;
	goto L473
L473:
	;
	v4439 = v4372 + int32(1)
	if v4439 != v4364 {
		v4372 = v4439
		goto L468
	} else {
		goto L477
	}
L474:
	;
	F_errmsg_internal(m, int32(655658), int32(0))
	mBase = m.M
	v4432 = m.ExcPending
	if v4432 != 0 {
		goto L1
	} else {
		goto L475
	}
L475:
	;
	F_errfinish(m, int32(518587), int32(509), int32(131669))
	mBase = m.M
	v4437 = m.ExcPending
	if v4437 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L477:
	;
	goto L469
L478:
	;
	goto L458
L479:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v56
	F_MemoryContextDelete(m, v53)
	mBase = m.M
	v4671 = m.ExcPending
	if v4671 != 0 {
		goto L1
	} else {
		goto L494
	}
L480:
	;
	if v4538 == int32(0) {
		goto L479
	} else {
		goto L481
	}
L481:
	;
	v4542 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4542 <= int32(0) {
		goto L479
	} else {
		goto L482
	}
L482:
	;
	v4545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v4546 = *(*int32)(unsafe.Add(mBase, uint32(v4545)))
	v4553 = int32(0)
	goto L483
L483:
	;
	v4594 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4594 <= v4553 {
		goto L446
	} else {
		goto L485
	}
L484:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4612 = m.ExcPending
	if v4612 != 0 {
		goto L1
	} else {
		goto L491
	}
L485:
	;
	v4596 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v4597 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v4600 = F_FunctionCall1Coll(m, v4538, v4546, v4596+v4597*v4553)
	mBase = m.M
	v4601 = m.ExcPending
	if v4601 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	v4602 = *(*float64)(unsafe.Add(mBase, uint32(v4600)))
	if base.F64_ne(v4602, float64(0)) != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v4606 = v4553 + int32(1)
	v4607 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4607 <= v4606 {
		goto L479
	} else {
		goto L490
	}
L488:
	;
	goto L489
L489:
	;
	goto L484
L490:
	;
	v4553 = v4606
	goto L483
L491:
	;
	F_errmsg_internal(m, int32(655617), int32(0))
	mBase = m.M
	v4616 = m.ExcPending
	if v4616 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	F_errfinish(m, int32(518587), int32(532), int32(160376))
	mBase = m.M
	v4621 = m.ExcPending
	if v4621 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	return
L495:
	;
	F_errmsg_internal(m, int32(655527), int32(0))
	mBase = m.M
	v4679 = m.ExcPending
	if v4679 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(518587), int32(543), int32(140954))
	mBase = m.M
	v4684 = m.ExcPending
	if v4684 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v4697 = m.ExcPending
	if v4697 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L501:
	;
	F_errmsg_internal(m, int32(655704), int32(0))
	mBase = m.M
	v4705 = m.ExcPending
	if v4705 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	F_errfinish(m, int32(518587), int32(506), int32(131669))
	mBase = m.M
	v4710 = m.ExcPending
	if v4710 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v4718 = m.ExcPending
	if v4718 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v4723 = m.ExcPending
	if v4723 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v4731 = m.ExcPending
	if v4731 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v4744 = m.ExcPending
	if v4744 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v4749 = m.ExcPending
	if v4749 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v4762 = m.ExcPending
	if v4762 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(476713), int32(0))
	mBase = m.M
	v4770 = m.ExcPending
	if v4770 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	F_errfinish(m, int32(343466), int32(326), int32(116085))
	mBase = m.M
	v4775 = m.ExcPending
	if v4775 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L519:
	;
	F_errmsg_internal(m, int32(655568), int32(0))
	mBase = m.M
	v4783 = m.ExcPending
	if v4783 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	F_errfinish(m, int32(518587), int32(294), int32(159395))
	mBase = m.M
	v4788 = m.ExcPending
	if v4788 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
