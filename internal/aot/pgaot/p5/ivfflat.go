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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v3 = int32(_a_F_IvfflatInitPage_0)
	v5 = int32(0)
	if v5|(l1&int32(3)|int32(1)) == v5 {
		v21 = l1 + v3
		v23 = l1 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l1^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l1, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l1, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(_a_F_IvfflatInitPage_1)
	v44 = int32(_a_F_IvfflatInitPage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v44)
	v50 = int32(_a_F_IvfflatInitPage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v54 = l1 + v53
	v55 = int32(_a_F_IvfflatInitPage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(-1)
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
	var v80 int32
	_ = v80
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v180 int32
	_ = v180
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v206 float64
	_ = v206
	var v210 int32
	_ = v210
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
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
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
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
	var v417 int32
	_ = v417
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
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v513 int32
	_ = v513
	var v514 int64
	_ = v514
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 float64
	_ = v735
	var v746 int32
	_ = v746
	var v785 float64
	_ = v785
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 float64
	_ = v806
	var v809 float64
	_ = v809
	var v812 int32
	_ = v812
	var v813 float32
	_ = v813
	var v814 float64
	_ = v814
	var v816 float32
	_ = v816
	var v820 float64
	_ = v820
	var v821 float64
	_ = v821
	var v823 int32
	_ = v823
	var v868 float64
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int64
	_ = v878
	var v879 int64
	_ = v879
	var v880 int64
	_ = v880
	var v901 float64
	_ = v901
	var v909 int32
	_ = v909
	var v948 float64
	_ = v948
	var v954 float32
	_ = v954
	var v956 float64
	_ = v956
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1099 int32
	_ = v1099
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 float32
	_ = v1133
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1178 float32
	_ = v1178
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 float32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1197 float32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 float32
	_ = v1203
	var v1207 float32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 float32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 float32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 float32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 float32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1265 float32
	_ = v1265
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1311 float32
	_ = v1311
	var v1322 float32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 float32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1337 int32
	_ = v1337
	var v1369 float32
	_ = v1369
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1455 int32
	_ = v1455
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1503 int32
	_ = v1503
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1562 int32
	_ = v1562
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 float64
	_ = v1615
	var v1618 float32
	_ = v1618
	var v1626 int32
	_ = v1626
	var v1681 int32
	_ = v1681
	var v1725 int32
	_ = v1725
	var v1726 float32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1768 float32
	_ = v1768
	var v1780 float32
	_ = v1780
	var v1782 float32
	_ = v1782
	var v1783 float32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1791 float32
	_ = v1791
	var v1793 float32
	_ = v1793
	var v1794 float32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1811 int32
	_ = v1811
	var v1841 float32
	_ = v1841
	var v1853 float32
	_ = v1853
	var v1855 float32
	_ = v1855
	var v1894 float32
	_ = v1894
	var v1907 int32
	_ = v1907
	var v1955 int32
	_ = v1955
	var v1958 int32
	_ = v1958
	var v1969 int32
	_ = v1969
	var v1980 int32
	_ = v1980
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 float32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2018 float32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2044 int32
	_ = v2044
	var v2071 int32
	_ = v2071
	var v2073 float32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 float32
	_ = v2077
	var v2084 float32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2108 float64
	_ = v2108
	var v2109 float32
	_ = v2109
	var v2112 float32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2115 float32
	_ = v2115
	var v2116 float32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2126 float32
	_ = v2126
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 float64
	_ = v2139
	var v2140 float32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2159 int32
	_ = v2159
	var v2180 int32
	_ = v2180
	var v2208 int32
	_ = v2208
	var v2232 int32
	_ = v2232
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2280 int32
	_ = v2280
	var v2284 int32
	_ = v2284
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2349 int32
	_ = v2349
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2365 int32
	_ = v2365
	var v2407 int32
	_ = v2407
	var v2469 int32
	_ = v2469
	var v2475 int32
	_ = v2475
	var v2516 int32
	_ = v2516
	var v2523 int32
	_ = v2523
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
	var v2648 int32
	_ = v2648
	var v2652 int32
	_ = v2652
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
	var v2736 int32
	_ = v2736
	var v2782 int32
	_ = v2782
	var v2784 int32
	_ = v2784
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
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2903 int32
	_ = v2903
	var v2940 int32
	_ = v2940
	var v2942 int32
	_ = v2942
	var v2946 int32
	_ = v2946
	var v2949 int32
	_ = v2949
	var v2956 int32
	_ = v2956
	var v3001 int32
	_ = v3001
	var v3004 int64
	_ = v3004
	var v3005 int64
	_ = v3005
	var v3006 int64
	_ = v3006
	var v3027 float64
	_ = v3027
	var v3031 int32
	_ = v3031
	var v3035 int32
	_ = v3035
	var v3041 int32
	_ = v3041
	var v3045 int32
	_ = v3045
	var v3085 int32
	_ = v3085
	var v3086 float32
	_ = v3086
	var v3094 float32
	_ = v3094
	var v3096 float32
	_ = v3096
	var v3104 float32
	_ = v3104
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3109 int32
	_ = v3109
	var v3117 int32
	_ = v3117
	var v3161 int32
	_ = v3161
	var v3162 float32
	_ = v3162
	var v3170 float32
	_ = v3170
	var v3218 float32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3267 int32
	_ = v3267
	var v3269 int32
	_ = v3269
	var v3270 float32
	_ = v3270
	var v3273 float32
	_ = v3273
	var v3277 int32
	_ = v3277
	var v3279 int32
	_ = v3279
	var v3289 int32
	_ = v3289
	var v3331 int32
	_ = v3331
	var v3332 float32
	_ = v3332
	var v3382 int32
	_ = v3382
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3438 int32
	_ = v3438
	var v3480 int32
	_ = v3480
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3486 int32
	_ = v3486
	var v3491 int32
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3545 int32
	_ = v3545
	var v3550 int32
	_ = v3550
	var v3551 int32
	_ = v3551
	var v3553 int32
	_ = v3553
	var v3555 int32
	_ = v3555
	var v3557 int32
	_ = v3557
	var v3564 int32
	_ = v3564
	var v3606 int32
	_ = v3606
	var v3608 int32
	_ = v3608
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3623 float64
	_ = v3623
	var v3627 int32
	_ = v3627
	var v3675 int32
	_ = v3675
	var v3684 int32
	_ = v3684
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3736 int32
	_ = v3736
	var v3740 int32
	_ = v3740
	var v3779 int32
	_ = v3779
	var v3780 int32
	_ = v3780
	var v3781 float32
	_ = v3781
	var v3782 float32
	_ = v3782
	var v3784 float32
	_ = v3784
	var v3785 float32
	_ = v3785
	var v3788 float32
	_ = v3788
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3793 float32
	_ = v3793
	var v3794 float32
	_ = v3794
	var v3796 float32
	_ = v3796
	var v3797 float32
	_ = v3797
	var v3800 float32
	_ = v3800
	var v3802 int32
	_ = v3802
	var v3803 int32
	_ = v3803
	var v3805 int32
	_ = v3805
	var v3813 int32
	_ = v3813
	var v3856 int32
	_ = v3856
	var v3857 int32
	_ = v3857
	var v3858 float32
	_ = v3858
	var v3859 float32
	_ = v3859
	var v3861 float32
	_ = v3861
	var v3862 float32
	_ = v3862
	var v3865 float32
	_ = v3865
	var v3914 int32
	_ = v3914
	var v3916 int32
	_ = v3916
	var v3924 int32
	_ = v3924
	var v3926 int32
	_ = v3926
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3968 int32
	_ = v3968
	var v3969 float32
	_ = v3969
	var v3971 int32
	_ = v3971
	var v3975 float32
	_ = v3975
	var v3979 int32
	_ = v3979
	var v3980 int32
	_ = v3980
	var v3981 float32
	_ = v3981
	var v3983 int32
	_ = v3983
	var v3987 float32
	_ = v3987
	var v3991 int32
	_ = v3991
	var v3993 int32
	_ = v3993
	var v4001 int32
	_ = v4001
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4045 int32
	_ = v4045
	var v4046 float32
	_ = v4046
	var v4048 int32
	_ = v4048
	var v4052 float32
	_ = v4052
	var v4101 int32
	_ = v4101
	var v4108 int32
	_ = v4108
	var v4150 int32
	_ = v4150
	var v4152 int32
	_ = v4152
	var v4153 int32
	_ = v4153
	var v4157 int32
	_ = v4157
	var v4159 int32
	_ = v4159
	var v4211 int32
	_ = v4211
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4268 int32
	_ = v4268
	var v4269 int32
	_ = v4269
	var v4272 int32
	_ = v4272
	var v4275 int32
	_ = v4275
	var v4280 int32
	_ = v4280
	var v4323 int32
	_ = v4323
	var v4329 int32
	_ = v4329
	var v4331 int32
	_ = v4331
	var v4332 int32
	_ = v4332
	var v4335 int32
	_ = v4335
	var v4337 int32
	_ = v4337
	var v4338 int32
	_ = v4338
	var v4339 int32
	_ = v4339
	var v4346 int32
	_ = v4346
	var v4391 float32
	_ = v4391
	var v4403 int32
	_ = v4403
	var v4407 int32
	_ = v4407
	var v4412 int32
	_ = v4412
	var v4414 int32
	_ = v4414
	var v4463 int32
	_ = v4463
	var v4464 int32
	_ = v4464
	var v4513 int32
	_ = v4513
	var v4514 int32
	_ = v4514
	var v4517 int32
	_ = v4517
	var v4520 int32
	_ = v4520
	var v4521 int32
	_ = v4521
	var v4527 int32
	_ = v4527
	var v4569 int32
	_ = v4569
	var v4571 int32
	_ = v4571
	var v4572 int32
	_ = v4572
	var v4575 int32
	_ = v4575
	var v4576 int32
	_ = v4576
	var v4577 float64
	_ = v4577
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4587 int32
	_ = v4587
	var v4591 int32
	_ = v4591
	var v4596 int32
	_ = v4596
	var v4646 int32
	_ = v4646
	var v4650 int32
	_ = v4650
	var v4654 int32
	_ = v4654
	var v4659 int32
	_ = v4659
	var v4663 int32
	_ = v4663
	var v4667 int32
	_ = v4667
	var v4672 int32
	_ = v4672
	var v4676 int32
	_ = v4676
	var v4680 int32
	_ = v4680
	var v4685 int32
	_ = v4685
	var v4735 int32
	_ = v4735
	var v4739 int32
	_ = v4739
	var v4744 int32
	_ = v4744
	v6 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	v53 = F_AllocSetContextCreateInternal(m, v48, int32(_a_F_IvfflatKmeans_0), v6, int32(_a_F_IvfflatKmeans_1), int32(_a_F_IvfflatKmeans_2))
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
	v55 = int32(_a_F_IvfflatKmeans_3)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0])) = v53
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v60 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4735 = m.ExcPending
	if v4735 != 0 {
		goto L1
	} else {
		goto L453
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4676 = m.ExcPending
	if v4676 != 0 {
		goto L1
	} else {
		goto L450
	}
L5:
	;
	v4260 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v4261 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4260 == v4261 {
		goto L399
	} else {
		goto L400
	}
L6:
	;
	v64 = F_HnswOptionalProcInfo(m, l0, int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v334 = F_mul_size(m, v328, (v329+int32(7))&int32(-8))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L33
	}
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v69 = F_mul_size(m, int32(4), v59)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v71 = F_palloc(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v73 < v74 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v80 = v73
	goto L15
L13:
	;
	goto L14
L14:
	;
	if v64 == int32(0) {
		goto L5
	} else {
		goto L29
	}
L15:
	;
	if int32(0) <= v80 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L14
L17:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	m.T0[v258].(func(*base.Module, int32, int32, int32))(m, v124+v125*v80, v59, v71)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L27
	}
L18:
	;
	v135 = v128
	goto L23
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v128 = int32(0)
	if v128 < v59 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L3
L22:
	;
	goto L17
L23:
	;
	v180 = int32(_a_F_IvfflatKmeans_4)
	v183 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1]))
	v184 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2]))
	v185 = v183 ^ v184
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2])) = base.I64_rotl(v185, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1])) = v185<<(uint(int64(16))%64) ^ base.I64_rotl(v183, int64(24)) ^ v185
	v206 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v183*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L25
L24:
	;
	goto L17
L25:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v71+v135<<(uint(int32(2))%32)))) = base.F32_demote_f64(v206)
	v210 = v135 + int32(1)
	if v210 != v59 {
		v135 = v210
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v263 = v261 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v263 < v265 {
		v80 = v263
		goto L15
	} else {
		goto L28
	}
L28:
	;
	goto L16
L29:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	v321 = F_AllocSetContextCreateInternal(m, v316, int32(_a_F_IvfflatKmeans_5), int32(0), int32(_a_F_IvfflatKmeans_1), int32(_a_F_IvfflatKmeans_2))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_IvfflatNormVectors(m, l3, v67, l2, v321)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_MemoryContextDelete(m, v321)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L5
L33:
	;
	v336 = F_add_size(m, int32(20), v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v339 = F_mul_size(m, v328, v59)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v341 = F_mul_size(m, int32(4), v339)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v344 = F_mul_size(m, int32(4), v328)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v347 = F_mul_size(m, int32(4), v60)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v350 = F_mul_size(m, v60, v328)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v352 = F_mul_size(m, int32(4), v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v355 = F_mul_size(m, int32(4), v60)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v358 = F_mul_size(m, int32(4), v328)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v361 = F_mul_size(m, v328, v328)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v363 = F_mul_size(m, int32(4), v361)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v366 = F_mul_size(m, int32(4), v328)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v368 = F_add_size(m, l4, v336)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v370 = F_add_size(m, v368, v341)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v372 = F_add_size(m, v370, v344)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v374 = F_add_size(m, v372, v347)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v376 = F_add_size(m, v374, v352)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v378 = F_add_size(m, v376, v355)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v380 = F_add_size(m, v378, v358)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v382 = F_add_size(m, v380, v363)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v384 = F_add_size(m, v382, v366)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_IvfflatCheckMemoryUsage(m, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v389 = base.I32_div_s(int32(2147483647), v328)
	if v389 < v328 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v393 = F_index_getprocinfo(m, l0, int32(1), int32(3))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v396 = F_HnswOptionalProcInfo(m, l0, int32(4))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v400 = F_palloc(m, v341)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v402 = F_palloc(m, v344)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v404 = F_palloc(m, v347)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v407 = F_palloc_extended(m, v352, int32(1))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v409 = F_palloc(m, v355)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v411 = F_palloc(m, v358)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v414 = F_palloc_extended(m, v363, int32(1))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v416 = F_palloc(m, v366)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v419 = F_VectorArrayInit(m, v328, v59, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v328
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v424 = F_mul_size(m, int32(4), v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v426 = F_palloc(m, v424)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v432 = F_index_getprocinfo(m, l0, int32(1), int32(3))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v437 = Fn13966(m, int64(32))
	mBase = m.M
	goto L71
L71:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v439 = base.I32_rem_u_s(v437, v438)
	if v439 < int32(0) {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v442 <= v439 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v444 = int32(0)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_VectorArraySet(m, l2, v444, v446+v447*v439)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v452 + int32(1)
	if v428 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v429 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L76:
	;
	v459 = v428 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v428) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v469 = v444
	v471 = int32(0)
	goto L80
L78:
	;
	v533 = v444
	goto L79
L79:
	;
	v580 = v533
	v581 = int32(0)
	goto L84
L80:
	;
	v513 = v426 + v469<<(uint(int32(2))%32)
	v514 = int64(9187343237679939583)
	*(*int64)(unsafe.Add(mBase, uint32(v513)+24)) = v514
	*(*int64)(unsafe.Add(mBase, uint32(v513)+16)) = v514
	*(*int64)(unsafe.Add(mBase, uint32(v513)+8)) = v514
	*(*int64)(unsafe.Add(mBase, uint32(v513))) = v514
	v522 = int32(8)
	v523 = v469 + v522
	v525 = v471 + v522
	if v525 != v428&int32(2147483640) {
		v469 = v523
		v471 = v525
		goto L80
	} else {
		goto L82
	}
L81:
	;
	if v459 == int32(0) {
		goto L75
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v533 = v523
	goto L79
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426+v580<<(uint(int32(2))%32)))) = int32(2139095039)
	v627 = int32(1)
	v630 = v581 + v627
	if v630 != v459 {
		v580 = v580 + v627
		v581 = v630
		goto L84
	} else {
		goto L86
	}
L85:
	;
	goto L75
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v426)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L117
	}
L88:
	;
	v681 = v428 - int32(1)
	v682 = int32(0)
	v690 = v682
	goto L89
L89:
	;
	v732 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[3]))
	if v732 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v735 = float64(0)
	if base.B2i32(v428 <= v682) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	v746 = int32(0)
	v785 = v735
	goto L98
L96:
	;
	v868 = v735
	goto L97
L97:
	;
	v872 = v690 + int32(1)
	if v872 == v429 {
		goto L87
	} else {
		goto L107
	}
L98:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v788 <= v746 {
		goto L3
	} else {
		goto L100
	}
L99:
	;
	v868 = v821
	goto L97
L100:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v790 <= v690 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v804 = F_FunctionCall2Coll(m, v432, v435, v796+v797*v746, v800+v801*v690)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v806 = *(*float64)(unsafe.Add(mBase, uint32(v804)))
	*(*float32)(unsafe.Add(mBase, uint32(v407+v690<<(uint(int32(2))%32)+v746*v429<<(uint(int32(2))%32)))) = base.F32_demote_f64(v806)
	v809 = base.F64_mul(v806, v806)
	v812 = v426 + v746<<(uint(int32(2))%32)
	v813 = *(*float32)(unsafe.Add(mBase, uint32(v812)))
	v814 = base.F64_promote_f32(v813)
	if base.F64_lt(v809, v814) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v816 = base.F32_demote_f64(v809)
	*(*float32)(unsafe.Add(mBase, uint32(v812))) = v816
	v820 = base.F64_promote_f32(v816)
	goto L105
L104:
	;
	v820 = v814
	goto L105
L105:
	;
	v821 = base.F64_add(v820, v785)
	v823 = v746 + int32(1)
	if v823 != v428 {
		v746 = v823
		v785 = v821
		goto L98
	} else {
		goto L106
	}
L106:
	;
	goto L99
L107:
	;
	v874 = int32(0)
	v875 = int32(_a_F_IvfflatKmeans_4)
	v878 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1]))
	v879 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2]))
	v880 = v878 ^ v879
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2])) = base.I64_rotl(v880, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1])) = v880<<(uint(int64(16))%64) ^ base.I64_rotl(v878, int64(24)) ^ v880
	v901 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v878*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L108
L108:
	;
	if v428 < int32(2) {
		v966 = v874
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1008 <= v966 {
		goto L3
	} else {
		goto L115
	}
L110:
	;
	v909 = v874
	v948 = base.F64_mul(v901, v868)
	goto L111
L111:
	;
	v954 = *(*float32)(unsafe.Add(mBase, uint32(v426+v909<<(uint(int32(2))%32))))
	v956 = base.F64_sub(v948, base.F64_promote_f32(v954))
	if base.F64_le(v956, float64(0)) != 0 {
		v966 = v909
		goto L109
	} else {
		goto L113
	}
L112:
	;
	v966 = v681
	goto L109
L113:
	;
	v960 = v909 + int32(1)
	if v960 != v681 {
		v909 = v960
		v948 = v956
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_VectorArraySet(m, l2, v872, v1010+v1011*v966)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1016 + int32(1)
	v690 = v872
	goto L89
L117:
	;
	if int32(0) < v60 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v1072 = int32(3)
	v1073 = v328 & v1072
	v1099 = v6
	goto L121
L119:
	;
	goto L120
L120:
	;
	v1432 = int32(2147483646)
	v1434 = int32(1)
	v1437 = v328 & v1432
	v1439 = v328 & v1434
	v1441 = v328 - v1434
	v1455 = int32(0)
	goto L168
L121:
	;
	if v328 <= int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L120
L123:
	;
	v1378 = v1099 << (uint(int32(2)) % 32)
	*(*float32)(unsafe.Add(mBase, uint32(v409+v1378))) = v1369
	*(*int32)(unsafe.Add(mBase, uint32(v1378+v404))) = v1337
	v1384 = v1099 + int32(1)
	if v1384 != v60 {
		v1099 = v1384
		goto L121
	} else {
		goto L167
	}
L124:
	;
	v1337 = int32(0)
	v1369 = float32(3.4028235e+38)
	goto L123
L125:
	;
	goto L126
L126:
	;
	v1131 = v407 + v328*v1099<<(uint(int32(2))%32)
	v1132 = int32(0)
	v1133 = float32(3.4028235e+38)
	if base.B2i32(base.Ui32(v328-int32(1)) < base.Ui32(v1072)) == v1132 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1144 = v1132
	v1146 = v1132
	v1147 = v1132
	v1178 = v1133
	goto L130
L128:
	;
	v1232 = v1132
	v1233 = v1132
	v1265 = v1133
	goto L129
L129:
	;
	v1278 = v1232
	v1279 = v1233
	v1281 = v1132
	v1311 = v1265
	goto L158
L130:
	;
	v1187 = v1144 | int32(3)
	v1188 = int32(2)
	v1191 = *(*float32)(unsafe.Add(mBase, uint32(v1131+v1187<<(uint(v1188)%32))))
	v1193 = v1144 | v1188
	v1197 = *(*float32)(unsafe.Add(mBase, uint32(v1131+v1193<<(uint(v1188)%32))))
	v1199 = v1144 | int32(1)
	v1203 = *(*float32)(unsafe.Add(mBase, uint32(v1131+v1199<<(uint(v1188)%32))))
	v1207 = *(*float32)(unsafe.Add(mBase, uint32(v1131+v1144<<(uint(v1188)%32))))
	v1208 = base.F32_gt(v1178, v1207)
	if v1208 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v1073 == int32(0) {
		v1337 = v1219
		v1369 = v1215
		goto L123
	} else {
		goto L157
	}
L132:
	;
	v1209 = v1207
	goto L134
L133:
	;
	v1209 = v1178
	goto L134
L134:
	;
	v1210 = base.F32_gt(v1209, v1203)
	if v1210 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v1211 = v1203
	goto L137
L136:
	;
	v1211 = v1209
	goto L137
L137:
	;
	v1212 = base.F32_gt(v1211, v1197)
	if v1212 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v1213 = v1197
	goto L140
L139:
	;
	v1213 = v1211
	goto L140
L140:
	;
	v1214 = base.F32_gt(v1213, v1191)
	if v1214 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v1215 = v1191
	goto L143
L142:
	;
	v1215 = v1213
	goto L143
L143:
	;
	if v1208 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v1216 = v1144
	goto L146
L145:
	;
	v1216 = v1146
	goto L146
L146:
	;
	if v1210 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v1217 = v1199
	goto L149
L148:
	;
	v1217 = v1216
	goto L149
L149:
	;
	if v1212 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v1218 = v1193
	goto L152
L151:
	;
	v1218 = v1217
	goto L152
L152:
	;
	if v1214 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1219 = v1187
	goto L155
L154:
	;
	v1219 = v1218
	goto L155
L155:
	;
	v1220 = int32(4)
	v1221 = v1144 + v1220
	v1223 = v1147 + v1220
	if v1223 != v328&int32(2147483644) {
		v1144 = v1221
		v1146 = v1219
		v1147 = v1223
		v1178 = v1215
		goto L130
	} else {
		goto L156
	}
L156:
	;
	goto L131
L157:
	;
	v1232 = v1221
	v1233 = v1219
	v1265 = v1215
	goto L129
L158:
	;
	v1322 = *(*float32)(unsafe.Add(mBase, uint32(v1131+v1278<<(uint(int32(2))%32))))
	v1323 = base.F32_gt(v1311, v1322)
	if v1323 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v1337 = v1325
	v1369 = v1324
	goto L123
L160:
	;
	v1324 = v1322
	goto L162
L161:
	;
	v1324 = v1311
	goto L162
L162:
	;
	if v1323 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1325 = v1278
	goto L165
L164:
	;
	v1325 = v1279
	goto L165
L165:
	;
	v1326 = int32(1)
	v1329 = v1281 + v1326
	if v1329 != v1073 {
		v1278 = v1278 + v1326
		v1279 = v1325
		v1281 = v1329
		v1311 = v1324
		goto L158
	} else {
		goto L166
	}
L166:
	;
	goto L159
L167:
	;
	goto L122
L168:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[3]))
	if v1490 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L5
L170:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v1493 = int32(0)
	if v1493 < v328 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	goto L172
L174:
	;
	v1503 = v1493
	goto L177
L175:
	;
	goto L176
L176:
	;
	v1955 = int32(0)
	if v1955 < v60 {
		goto L218
	} else {
		goto L219
	}
L177:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1503 < v1542 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v1681 = int32(0)
	goto L190
L179:
	;
	if v1545 != v328 {
		v1503 = v1545
		goto L177
	} else {
		goto L189
	}
L180:
	;
	v1545 = v1503 + int32(1)
	if v328 <= v1545 {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	goto L3
L183:
	;
	v1547 = int32(2)
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1562 = v1545
	goto L184
L184:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1604 <= v1562 {
		goto L3
	} else {
		goto L186
	}
L185:
	;
	goto L179
L186:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1613 = F_FunctionCall2Coll(m, v393, v399, v1550+v1551*v1503, v1609+v1610*v1562)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v1615 = *(*float64)(unsafe.Add(mBase, uint32(v1613)))
	v1618 = base.F32_demote_f64(base.F64_mul(v1615, float64(0.5)))
	*(*float32)(unsafe.Add(mBase, uint32(v414+v1503*v328<<(uint(v1547)%32)+v1562<<(uint(int32(2))%32)))) = v1618
	*(*float32)(unsafe.Add(mBase, uint32(v414+v1503<<(uint(v1547)%32)+v1562*v328<<(uint(int32(2))%32)))) = v1618
	v1626 = v1562 + int32(1)
	if v328 != v1626 {
		v1562 = v1626
		goto L184
	} else {
		goto L188
	}
L188:
	;
	goto L185
L189:
	;
	goto L178
L190:
	;
	v1725 = v414 + v1681*v328<<(uint(int32(2))%32)
	v1726 = float32(3.4028235e+38)
	v1727 = int32(0)
	if v1441 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	goto L176
L192:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v411+v1681<<(uint(int32(2))%32)))) = v1894
	v1907 = v1681 + int32(1)
	if v1907 != v328 {
		v1681 = v1907
		goto L190
	} else {
		goto L216
	}
L193:
	;
	v1734 = v1727
	v1736 = v1727
	v1768 = v1726
	goto L196
L194:
	;
	v1811 = v1727
	v1841 = v1726
	goto L195
L195:
	;
	if v1681 == v1811 {
		v1894 = v1841
		goto L192
	} else {
		goto L212
	}
L196:
	;
	if v1734 != v1681 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v1439 == int32(0) {
		v1894 = v1794
		goto L192
	} else {
		goto L211
	}
L198:
	;
	v1780 = *(*float32)(unsafe.Add(mBase, uint32(v1725+v1734<<(uint(int32(2))%32))))
	if base.F32_gt(v1768, v1780) != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1783 = v1768
	goto L200
L200:
	;
	v1786 = v1734 | int32(1)
	if v1786 != v1681 {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	v1782 = v1780
	goto L203
L202:
	;
	v1782 = v1768
	goto L203
L203:
	;
	v1783 = v1782
	goto L200
L204:
	;
	v1791 = *(*float32)(unsafe.Add(mBase, uint32(v1725+v1786<<(uint(int32(2))%32))))
	if base.F32_gt(v1783, v1791) != 0 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v1794 = v1783
	goto L206
L206:
	;
	v1796 = int32(2)
	v1797 = v1734 + v1796
	v1799 = v1736 + v1796
	if v1799 != v1437 {
		v1734 = v1797
		v1736 = v1799
		v1768 = v1794
		goto L196
	} else {
		goto L210
	}
L207:
	;
	v1793 = v1791
	goto L209
L208:
	;
	v1793 = v1783
	goto L209
L209:
	;
	v1794 = v1793
	goto L206
L210:
	;
	goto L197
L211:
	;
	v1811 = v1797
	v1841 = v1794
	goto L195
L212:
	;
	v1853 = *(*float32)(unsafe.Add(mBase, uint32(v1725+v1811<<(uint(int32(2))%32))))
	if base.F32_gt(v1841, v1853) != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1855 = v1853
	goto L215
L214:
	;
	v1855 = v1841
	goto L215
L215:
	;
	v1894 = v1855
	goto L192
L216:
	;
	goto L191
L217:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v2262 = int32(0)
	v2263 = base.B2i32(v2261 <= v2262)
	if v2263 == v2262 {
		goto L248
	} else {
		goto L249
	}
L218:
	;
	v1958 = int32(0)
	v1969 = v1955
	v1980 = v1958
	goto L221
L219:
	;
	goto L220
L220:
	;
	v2232 = int32(1)
	goto L217
L221:
	;
	if v328 <= int32(0) {
		v2180 = v1980
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v2232 = base.B2i32(v2180 == int32(0))
	goto L217
L223:
	;
	v2208 = v1969 + int32(1)
	if v2208 != v60 {
		v1969 = v2208
		v1980 = v2180
		goto L221
	} else {
		goto L247
	}
L224:
	;
	v2009 = int32(2)
	v2010 = v1969 << (uint(v2009) % 32)
	v2011 = v409 + v2010
	v2012 = *(*float32)(unsafe.Add(mBase, uint32(v2011)))
	v2013 = v2010 + v404
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2013)))
	v2018 = *(*float32)(unsafe.Add(mBase, uint32(v411+v2014<<(uint(v2009)%32))))
	if base.F32_le(v2012, v2018) != 0 {
		v2180 = v1980
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v2023 = v407 + v1969*v328<<(uint(int32(2))%32)
	v2029 = int32(0)
	v2031 = base.B2i32(v1455 != v1958)
	v2044 = v1980
	goto L226
L226:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2013)))
	if v2029 == v2071 {
		v2151 = v2031
		v2153 = v2044
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v2180 = v2153
	goto L223
L228:
	;
	v2159 = v2029 + int32(1)
	if v2159 != v328 {
		v2029 = v2159
		v2031 = v2151
		v2044 = v2153
		goto L226
	} else {
		goto L246
	}
L229:
	;
	v2073 = *(*float32)(unsafe.Add(mBase, uint32(v2011)))
	v2075 = v2029 << (uint(int32(2)) % 32)
	v2076 = v2023 + v2075
	v2077 = *(*float32)(unsafe.Add(mBase, uint32(v2076)))
	if base.F32_le(v2073, v2077) != 0 {
		v2151 = v2031
		v2153 = v2044
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v2084 = *(*float32)(unsafe.Add(mBase, uint32(v414+v2071*v328<<(uint(int32(2))%32)+v2075)))
	if base.F32_le(v2073, v2084) != 0 {
		v2151 = v2031
		v2153 = v2044
		goto L228
	} else {
		goto L231
	}
L231:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2086 <= v1969 {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2091 = v2088 + v2089*v1969
	if v2031&int32(1) != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v2071 < int32(0) {
		goto L3
	} else {
		goto L236
	}
L234:
	;
	v2113 = v2071
	v2115 = v2073
	v2116 = v2077
	goto L235
L235:
	;
	if base.F32_gt(v2115, v2116) == int32(0) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2096 <= v2071 {
		goto L3
	} else {
		goto L237
	}
L237:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2102 = F_FunctionCall2Coll(m, v393, v399, v2091, v2098+v2099*v2071)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v2013)))
	v2108 = *(*float64)(unsafe.Add(mBase, uint32(v2102)))
	v2109 = base.F32_demote_f64(v2108)
	*(*float32)(unsafe.Add(mBase, uint32(v2023+v2104<<(uint(int32(2))%32)))) = v2109
	*(*float32)(unsafe.Add(mBase, uint32(v2011))) = v2109
	v2112 = *(*float32)(unsafe.Add(mBase, uint32(v2076)))
	v2113 = v2104
	v2115 = v2109
	v2116 = v2112
	goto L235
L239:
	;
	v2120 = int32(0)
	v2126 = *(*float32)(unsafe.Add(mBase, uint32(v414+v2113*v328<<(uint(int32(2))%32)+v2075)))
	if base.F32_gt(v2115, v2126) == v2120 {
		v2151 = v2120
		v2153 = v2044
		goto L228
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2131 <= v2029 {
		goto L3
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2137 = F_FunctionCall2Coll(m, v393, v399, v2091, v2133+v2134*v2029)
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v2139 = *(*float64)(unsafe.Add(mBase, uint32(v2137)))
	v2140 = base.F32_demote_f64(v2139)
	*(*float32)(unsafe.Add(mBase, uint32(v2076))) = v2140
	v2142 = int32(0)
	if base.F32_gt(v2115, v2140) == v2142 {
		v2151 = v2142
		v2153 = v2044
		goto L228
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2013))) = v2029
	*(*float32)(unsafe.Add(mBase, uint32(v2011))) = v2140
	v2151 = v2142
	v2153 = v2044 + int32(1)
	goto L228
L246:
	;
	goto L227
L247:
	;
	goto L222
L248:
	;
	v2267 = v2259 << (uint(int32(2)) % 32)
	v2268 = int32(0)
	if v2261 != int32(1) {
		goto L252
	} else {
		goto L253
	}
L249:
	;
	v2475 = v2260
	goto L250
L250:
	;
	v2516 = int32(0)
	if v2516 < v2475 {
		goto L269
	} else {
		goto L270
	}
L251:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2475 = v2469
	goto L250
L252:
	;
	v2280 = v2268
	v2284 = int32(0)
	goto L255
L253:
	;
	v2365 = v2268
	goto L254
L254:
	;
	v2407 = int32(0)
	if base.B2i32(v2267 == v2407)|base.B2i32(v2259 <= v2407) == v2407 {
		goto L265
	} else {
		goto L266
	}
L255:
	;
	v2322 = int32(0)
	v2323 = base.B2i32(v2259 <= v2322)
	if v2323|base.B2i32(v2267 == v2322) == v2322 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v2261&int32(1) == int32(0) {
		goto L251
	} else {
		goto L264
	}
L257:
	;
	base.MemoryFill(m, v400+v2280*v2267, int32(0), v2267)
	goto L259
L258:
	;
	goto L259
L259:
	;
	v2336 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v402+v2280<<(uint(int32(2))%32)))) = v2336
	v2339 = v2280 | int32(1)
	if v2323|base.B2i32(v2267 == v2336) == v2336 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	base.MemoryFill(m, v400+v2339*v2267, int32(0), v2267)
	goto L262
L261:
	;
	goto L262
L262:
	;
	v2349 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v402+v2339<<(uint(v2349)%32)))) = int32(0)
	v2355 = v2280 + v2349
	v2357 = v2284 + v2349
	if v2357 != v2261&int32(2147483646) {
		v2280 = v2355
		v2284 = v2357
		goto L255
	} else {
		goto L263
	}
L263:
	;
	goto L256
L264:
	;
	v2365 = v2355
	goto L254
L265:
	;
	base.MemoryFill(m, v400+v2365*v2267, int32(0), v2267)
	goto L267
L266:
	;
	goto L267
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v402+v2365<<(uint(int32(2))%32)))) = int32(0)
	goto L251
L268:
	;
	v3430 = int32(0)
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	if v3430 < v3431 {
		goto L335
	} else {
		goto L336
	}
L269:
	;
	v2523 = v2516
	goto L272
L270:
	;
	goto L271
L271:
	;
	if v2260 <= int32(0) {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2565 <= v2523 {
		goto L3
	} else {
		goto L274
	}
L273:
	;
	goto L271
L274:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2572 = int32(2)
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v404+v2523<<(uint(v2572)%32))))
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	m.T0[v2580].(func(*base.Module, int32, int32))(m, v2567+v2568*v2523, v400+v2571*v2575<<(uint(v2572)%32))
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v2584 = v2523 + int32(1)
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2584 < v2585 {
		v2523 = v2584
		goto L272
	} else {
		goto L276
	}
L276:
	;
	goto L273
L277:
	;
	if v2261 <= v2262 {
		goto L268
	} else {
		goto L289
	}
L278:
	;
	v2636 = v2260 & int32(3)
	v2637 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v2260) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v2648 = v2637
	v2652 = int32(0)
	goto L282
L280:
	;
	v2736 = v2637
	goto L281
L281:
	;
	v2782 = v2736
	v2784 = v2637
	goto L286
L282:
	;
	v2690 = int32(2)
	v2692 = v404 + v2648<<(uint(v2690)%32)
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(v2692)))
	v2696 = v402 + v2693<<(uint(v2690)%32)
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v2696)))
	v2698 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2696))) = v2697 + v2698
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(v2692)+4))
	v2704 = v402 + v2701<<(uint(v2690)%32)
	v2705 = *(*int32)(unsafe.Add(mBase, uint32(v2704)))
	*(*int32)(unsafe.Add(mBase, uint32(v2704))) = v2705 + v2698
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2692)+8))
	v2712 = v402 + v2709<<(uint(v2690)%32)
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v2712)))
	*(*int32)(unsafe.Add(mBase, uint32(v2712))) = v2713 + v2698
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2692)+12))
	v2720 = v402 + v2717<<(uint(v2690)%32)
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v2720)))
	*(*int32)(unsafe.Add(mBase, uint32(v2720))) = v2721 + v2698
	v2725 = int32(4)
	v2726 = v2648 + v2725
	v2728 = v2652 + v2725
	if v2728 != v2260&int32(2147483644) {
		v2648 = v2726
		v2652 = v2728
		goto L282
	} else {
		goto L284
	}
L283:
	;
	if v2636 == int32(0) {
		goto L277
	} else {
		goto L285
	}
L284:
	;
	goto L283
L285:
	;
	v2736 = v2726
	goto L281
L286:
	;
	v2824 = int32(2)
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v404+v2782<<(uint(v2824)%32))))
	v2830 = v402 + v2827<<(uint(v2824)%32)
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2830)))
	v2832 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2830))) = v2831 + v2832
	v2838 = v2784 + v2832
	if v2838 != v2636 {
		v2782 = v2782 + v2832
		v2784 = v2838
		goto L286
	} else {
		goto L288
	}
L287:
	;
	goto L277
L288:
	;
	goto L287
L289:
	;
	v2887 = v2259 & int32(2147483646)
	v2888 = int32(1)
	v2889 = v2259 & v2888
	v2891 = v2259 - v2888
	v2903 = int32(0)
	goto L290
L290:
	;
	v2940 = int32(2)
	v2942 = v400 + v2259*v2903<<(uint(v2940)%32)
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v402+v2903<<(uint(v2940)%32))))
	if v2946 <= int32(0) {
		goto L293
	} else {
		goto L294
	}
L291:
	;
	goto L268
L292:
	;
	v3382 = v2903 + int32(1)
	if v3382 != v2261 {
		v2903 = v3382
		goto L290
	} else {
		goto L333
	}
L293:
	;
	v2949 = int32(0)
	if v2259 <= v2949 {
		goto L292
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	if v2259 <= int32(0) {
		goto L292
	} else {
		goto L301
	}
L296:
	;
	v2956 = v2949
	goto L297
L297:
	;
	v3001 = int32(_a_F_IvfflatKmeans_4)
	v3004 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1]))
	v3005 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2]))
	v3006 = v3004 ^ v3005
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2])) = base.I64_rotl(v3006, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1])) = v3006<<(uint(int64(16))%64) ^ base.I64_rotl(v3004, int64(24)) ^ v3006
	v3027 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v3004*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L299
L298:
	;
	goto L292
L299:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v2942+v2956<<(uint(int32(2))%32)))) = base.F32_demote_f64(v3027)
	v3031 = v2956 + int32(1)
	if v3031 != v2259 {
		v2956 = v3031
		goto L297
	} else {
		goto L300
	}
L300:
	;
	goto L298
L301:
	;
	v3035 = int32(0)
	if v2891 != 0 {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	v3218 = base.F32_convert_i32_u(v2946)
	v3219 = int32(0)
	if v2891 != 0 {
		goto L326
	} else {
		goto L327
	}
L303:
	;
	v3041 = v3035
	v3045 = v3035
	goto L306
L304:
	;
	v3117 = v3035
	goto L305
L305:
	;
	v3161 = v2942 + v3117<<(uint(int32(2))%32)
	v3162 = *(*float32)(unsafe.Add(mBase, uint32(v3161)))
	if base.F32_ne(base.F32_abs(v3162), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L302
	} else {
		goto L322
	}
L306:
	;
	v3085 = v2942 + v3041<<(uint(int32(2))%32)
	v3086 = *(*float32)(unsafe.Add(mBase, uint32(v3085)))
	if base.F32_eq(base.F32_abs(v3086), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	if v2889 == int32(0) {
		goto L302
	} else {
		goto L321
	}
L308:
	;
	if base.F32_gt(v3086, float32(0)) != 0 {
		goto L311
	} else {
		goto L312
	}
L309:
	;
	goto L310
L310:
	;
	v3096 = *(*float32)(unsafe.Add(mBase, uint32(v3085)+4))
	if base.F32_eq(base.F32_abs(v3096), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L314
	} else {
		goto L315
	}
L311:
	;
	v3094 = float32(3.4028235e+38)
	goto L313
L312:
	;
	v3094 = float32(-3.4028235e+38)
	goto L313
L313:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3085))) = v3094
	goto L310
L314:
	;
	if base.F32_gt(v3096, float32(0)) != 0 {
		goto L317
	} else {
		goto L318
	}
L315:
	;
	goto L316
L316:
	;
	v3106 = int32(2)
	v3107 = v3041 + v3106
	v3109 = v3045 + v3106
	if v3109 != v2887 {
		v3041 = v3107
		v3045 = v3109
		goto L306
	} else {
		goto L320
	}
L317:
	;
	v3104 = float32(3.4028235e+38)
	goto L319
L318:
	;
	v3104 = float32(-3.4028235e+38)
	goto L319
L319:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3085)+4)) = v3104
	goto L316
L320:
	;
	goto L307
L321:
	;
	v3117 = v3107
	goto L305
L322:
	;
	if base.F32_gt(v3162, float32(0)) != 0 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v3170 = float32(3.4028235e+38)
	goto L325
L324:
	;
	v3170 = float32(-3.4028235e+38)
	goto L325
L325:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3161))) = v3170
	goto L302
L326:
	;
	v3227 = v3219
	v3229 = v3219
	goto L329
L327:
	;
	v3289 = v3219
	goto L328
L328:
	;
	v3331 = v2942 + v3289<<(uint(int32(2))%32)
	v3332 = *(*float32)(unsafe.Add(mBase, uint32(v3331)))
	*(*float32)(unsafe.Add(mBase, uint32(v3331))) = base.F32_div(v3332, v3218)
	goto L292
L329:
	;
	v3267 = int32(2)
	v3269 = v2942 + v3227<<(uint(v3267)%32)
	v3270 = *(*float32)(unsafe.Add(mBase, uint32(v3269)))
	*(*float32)(unsafe.Add(mBase, uint32(v3269))) = base.F32_div(v3270, v3218)
	v3273 = *(*float32)(unsafe.Add(mBase, uint32(v3269)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v3269)+4)) = base.F32_div(v3273, v3218)
	v3277 = v3227 + v3267
	v3279 = v3229 + v3267
	if v3279 != v2887 {
		v3227 = v3277
		v3229 = v3279
		goto L329
	} else {
		goto L331
	}
L330:
	;
	if v2889 == int32(0) {
		goto L292
	} else {
		goto L332
	}
L331:
	;
	goto L330
L332:
	;
	v3289 = v3277
	goto L328
L333:
	;
	goto L291
L334:
	;
	v3675 = int32(0)
	if v60 <= v3675 {
		goto L356
	} else {
		goto L357
	}
L335:
	;
	v3438 = v3430
	goto L338
L336:
	;
	goto L337
L337:
	;
	if v396 != 0 {
		goto L343
	} else {
		goto L344
	}
L338:
	;
	v3480 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v3480 <= v3438 {
		goto L3
	} else {
		goto L340
	}
L339:
	;
	goto L337
L340:
	;
	v3482 = *(*int32)(unsafe.Add(mBase, uint32(v419)+16))
	v3483 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v3486 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v3491 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	m.T0[v3491].(func(*base.Module, int32, int32, int32))(m, v3482+v3483*v3438, v3486, v400+v3438*v3486<<(uint(int32(2))%32))
	mBase = m.M
	v3493 = m.ExcPending
	if v3493 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v3495 = v3438 + int32(1)
	v3496 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	if v3495 < v3496 {
		v3438 = v3495
		goto L338
	} else {
		goto L342
	}
L342:
	;
	goto L339
L343:
	;
	v3545 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	v3550 = F_AllocSetContextCreateInternal(m, v3545, int32(_a_F_IvfflatKmeans_5), int32(0), int32(_a_F_IvfflatKmeans_1), int32(_a_F_IvfflatKmeans_2))
	mBase = m.M
	v3551 = m.ExcPending
	if v3551 != 0 {
		goto L1
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	v3557 = int32(0)
	if v328 <= v3557 {
		goto L334
	} else {
		goto L349
	}
L346:
	;
	F_IvfflatNormVectors(m, l3, v399, v419, v3550)
	mBase = m.M
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	F_MemoryContextDelete(m, v3550)
	mBase = m.M
	v3555 = m.ExcPending
	if v3555 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	goto L345
L349:
	;
	v3564 = v3557
	goto L350
L350:
	;
	v3606 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3606 <= v3564 {
		goto L3
	} else {
		goto L352
	}
L351:
	;
	goto L334
L352:
	;
	v3608 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v3608 <= v3564 {
		goto L3
	} else {
		goto L353
	}
L353:
	;
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3617 = *(*int32)(unsafe.Add(mBase, uint32(v419)+16))
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v3621 = F_FunctionCall2Coll(m, v393, v399, v3613+v3614*v3564, v3617+v3618*v3564)
	mBase = m.M
	v3622 = m.ExcPending
	if v3622 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	v3623 = *(*float64)(unsafe.Add(mBase, uint32(v3621)))
	*(*float32)(unsafe.Add(mBase, uint32(v416+v3564<<(uint(int32(2))%32)))) = base.F32_demote_f64(v3623)
	v3627 = v3564 + int32(1)
	if v328 != v3627 {
		v3564 = v3627
		goto L350
	} else {
		goto L355
	}
L355:
	;
	goto L351
L356:
	;
	v4101 = int32(0)
	if v328 <= v4101 {
		goto L386
	} else {
		goto L387
	}
L357:
	;
	v3684 = v3675
	goto L358
L358:
	;
	if v328 <= int32(0) {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	v3916 = int32(0)
	if v60 != int32(1) {
		goto L379
	} else {
		goto L380
	}
L360:
	;
	v3914 = v3684 + int32(1)
	if v3914 != v60 {
		v3684 = v3914
		goto L358
	} else {
		goto L378
	}
L361:
	;
	v3729 = v407 + v3684*v328<<(uint(int32(2))%32)
	v3730 = int32(0)
	if v1441 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v3736 = v3730
	v3740 = v3730
	goto L365
L363:
	;
	v3813 = v3730
	goto L364
L364:
	;
	v3856 = v3813 << (uint(int32(2)) % 32)
	v3857 = v3729 + v3856
	v3858 = float32(0)
	v3859 = *(*float32)(unsafe.Add(mBase, uint32(v3857)))
	v3861 = *(*float32)(unsafe.Add(mBase, uint32(v3856+v416)))
	v3862 = base.F32_sub(v3859, v3861)
	if base.F32_lt(v3862, v3858) != 0 {
		goto L375
	} else {
		goto L376
	}
L365:
	;
	v3779 = v3736 << (uint(int32(2)) % 32)
	v3780 = v3729 + v3779
	v3781 = float32(0)
	v3782 = *(*float32)(unsafe.Add(mBase, uint32(v3780)))
	v3784 = *(*float32)(unsafe.Add(mBase, uint32(v3779+v416)))
	v3785 = base.F32_sub(v3782, v3784)
	if base.F32_lt(v3785, v3781) != 0 {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	if v1439 == int32(0) {
		goto L360
	} else {
		goto L374
	}
L367:
	;
	v3788 = v3781
	goto L369
L368:
	;
	v3788 = v3785
	goto L369
L369:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3780))) = v3788
	v3791 = v3779 | int32(4)
	v3792 = v3729 + v3791
	v3793 = float32(0)
	v3794 = *(*float32)(unsafe.Add(mBase, uint32(v3792)))
	v3796 = *(*float32)(unsafe.Add(mBase, uint32(v3791+v416)))
	v3797 = base.F32_sub(v3794, v3796)
	if base.F32_lt(v3797, v3793) != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v3800 = v3793
	goto L372
L371:
	;
	v3800 = v3797
	goto L372
L372:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3792))) = v3800
	v3802 = int32(2)
	v3803 = v3736 + v3802
	v3805 = v3740 + v3802
	if v3805 != v1437 {
		v3736 = v3803
		v3740 = v3805
		goto L365
	} else {
		goto L373
	}
L373:
	;
	goto L366
L374:
	;
	v3813 = v3803
	goto L364
L375:
	;
	v3865 = v3858
	goto L377
L376:
	;
	v3865 = v3862
	goto L377
L377:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3857))) = v3865
	goto L360
L378:
	;
	goto L359
L379:
	;
	v3924 = v3916
	v3926 = v3916
	goto L382
L380:
	;
	v4001 = v3916
	goto L381
L381:
	;
	v4043 = int32(2)
	v4044 = v4001 << (uint(v4043) % 32)
	v4045 = v409 + v4044
	v4046 = *(*float32)(unsafe.Add(mBase, uint32(v4045)))
	v4048 = *(*int32)(unsafe.Add(mBase, uint32(v4044+v404)))
	v4052 = *(*float32)(unsafe.Add(mBase, uint32(v416+v4048<<(uint(v4043)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v4045))) = base.F32_add(v4046, v4052)
	goto L356
L382:
	;
	v3966 = int32(2)
	v3967 = v3924 << (uint(v3966) % 32)
	v3968 = v409 + v3967
	v3969 = *(*float32)(unsafe.Add(mBase, uint32(v3968)))
	v3971 = *(*int32)(unsafe.Add(mBase, uint32(v3967+v404)))
	v3975 = *(*float32)(unsafe.Add(mBase, uint32(v416+v3971<<(uint(v3966)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v3968))) = base.F32_add(v3969, v3975)
	v3979 = v3967 | int32(4)
	v3980 = v409 + v3979
	v3981 = *(*float32)(unsafe.Add(mBase, uint32(v3980)))
	v3983 = *(*int32)(unsafe.Add(mBase, uint32(v3979+v404)))
	v3987 = *(*float32)(unsafe.Add(mBase, uint32(v416+v3983<<(uint(v3966)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v3980))) = base.F32_add(v3981, v3987)
	v3991 = v3924 + v3966
	v3993 = v3926 + v3966
	if v3993 != v60&v1432 {
		v3924 = v3991
		v3926 = v3993
		goto L382
	} else {
		goto L384
	}
L383:
	;
	if v60&v1434 == int32(0) {
		goto L356
	} else {
		goto L385
	}
L384:
	;
	goto L383
L385:
	;
	v4001 = v3991
	goto L381
L386:
	;
	if base.B2i32(v1455 != int32(0))&v2232 != 0 {
		goto L5
	} else {
		goto L395
	}
L387:
	;
	v4108 = v4101
	goto L388
L388:
	;
	v4150 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v4108 < v4150 {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	goto L3
L390:
	;
	v4152 = *(*int32)(unsafe.Add(mBase, uint32(v419)+16))
	v4153 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	F_VectorArraySet(m, l2, v4108, v4152+v4153*v4108)
	mBase = m.M
	v4157 = m.ExcPending
	if v4157 != 0 {
		goto L1
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	goto L389
L393:
	;
	v4159 = v4108 + int32(1)
	if v328 != v4159 {
		v4108 = v4159
		goto L388
	} else {
		goto L394
	}
L394:
	;
	goto L386
L395:
	;
	v4211 = v1455 + int32(1)
	if v4211 != int32(500) {
		v1455 = v4211
		goto L168
	} else {
		goto L396
	}
L396:
	;
	goto L169
L397:
	;
	goto L3
L398:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4663 = m.ExcPending
	if v4663 != 0 {
		goto L1
	} else {
		goto L447
	}
L399:
	;
	v4264 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4265 = F_mul_size(m, int32(4), v4264)
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L1
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4650 = m.ExcPending
	if v4650 != 0 {
		goto L1
	} else {
		goto L444
	}
L402:
	;
	v4267 = F_palloc(m, v4265)
	mBase = m.M
	v4268 = m.ExcPending
	if v4268 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v4269 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v4269 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v4272 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4275 = v4272
	v4280 = int32(0)
	goto L407
L405:
	;
	goto L406
L406:
	;
	v4513 = F_HnswOptionalProcInfo(m, l0, int32(2))
	mBase = m.M
	v4514 = m.ExcPending
	if v4514 != 0 {
		goto L1
	} else {
		goto L429
	}
L407:
	;
	if v4275 <= int32(0) {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	goto L406
L409:
	;
	v4329 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4329 <= v4280 {
		goto L3
	} else {
		goto L412
	}
L410:
	;
	v4323 = v4275 << (uint(int32(2)) % 32)
	if v4323 == int32(0) {
		goto L409
	} else {
		goto L411
	}
L411:
	;
	base.MemoryFill(m, v4267, int32(0), v4323)
	goto L409
L412:
	;
	v4331 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v4332 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v4335 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	m.T0[v4335].(func(*base.Module, int32, int32))(m, v4331+v4332*v4280, v4267)
	mBase = m.M
	v4337 = m.ExcPending
	if v4337 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	v4338 = int32(0)
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4338 < v4339 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v4346 = v4338
	goto L417
L415:
	;
	goto L416
L416:
	;
	v4463 = v4280 + int32(1)
	v4464 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4463 < v4464 {
		v4275 = v4339
		v4280 = v4463
		goto L407
	} else {
		goto L427
	}
L417:
	;
	v4391 = *(*float32)(unsafe.Add(mBase, uint32(v4267+v4346<<(uint(int32(2))%32))))
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v4391)&int32(2147483647)) {
		goto L398
	} else {
		goto L419
	}
L418:
	;
	goto L416
L419:
	;
	if base.F32_eq(base.F32_abs(v4391), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4403 = m.ExcPending
	if v4403 != 0 {
		goto L1
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v4414 = v4346 + int32(1)
	if v4414 != v4339 {
		v4346 = v4414
		goto L417
	} else {
		goto L426
	}
L423:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_6), int32(0))
	mBase = m.M
	v4407 = m.ExcPending
	if v4407 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(509), int32(_a_F_IvfflatKmeans_8))
	mBase = m.M
	v4412 = m.ExcPending
	if v4412 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L426:
	;
	goto L418
L427:
	;
	goto L408
L428:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0])) = v56
	F_MemoryContextDelete(m, v53)
	mBase = m.M
	v4646 = m.ExcPending
	if v4646 != 0 {
		goto L1
	} else {
		goto L443
	}
L429:
	;
	if v4513 == int32(0) {
		goto L428
	} else {
		goto L430
	}
L430:
	;
	v4517 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4517 <= int32(0) {
		goto L428
	} else {
		goto L431
	}
L431:
	;
	v4520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v4521 = *(*int32)(unsafe.Add(mBase, uint32(v4520)))
	v4527 = int32(0)
	goto L432
L432:
	;
	v4569 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4569 <= v4527 {
		goto L397
	} else {
		goto L434
	}
L433:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4587 = m.ExcPending
	if v4587 != 0 {
		goto L1
	} else {
		goto L440
	}
L434:
	;
	v4571 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v4572 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v4575 = F_FunctionCall1Coll(m, v4513, v4521, v4571+v4572*v4527)
	mBase = m.M
	v4576 = m.ExcPending
	if v4576 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	v4577 = *(*float64)(unsafe.Add(mBase, uint32(v4575)))
	if base.F64_ne(v4577, float64(0)) != 0 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v4581 = v4527 + int32(1)
	v4582 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4582 <= v4581 {
		goto L428
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	goto L433
L439:
	;
	v4527 = v4581
	goto L432
L440:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_9), int32(0))
	mBase = m.M
	v4591 = m.ExcPending
	if v4591 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(532), int32(_a_F_IvfflatKmeans_10))
	mBase = m.M
	v4596 = m.ExcPending
	if v4596 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L443:
	;
	return
L444:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_11), int32(0))
	mBase = m.M
	v4654 = m.ExcPending
	if v4654 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(543), int32(_a_F_IvfflatKmeans_12))
	mBase = m.M
	v4659 = m.ExcPending
	if v4659 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L447:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_13), int32(0))
	mBase = m.M
	v4667 = m.ExcPending
	if v4667 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(506), int32(_a_F_IvfflatKmeans_8))
	mBase = m.M
	v4672 = m.ExcPending
	if v4672 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_14), int32(0))
	mBase = m.M
	v4680 = m.ExcPending
	if v4680 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(294), int32(_a_F_IvfflatKmeans_15))
	mBase = m.M
	v4685 = m.ExcPending
	if v4685 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_16), int32(0))
	mBase = m.M
	v4739 = m.ExcPending
	if v4739 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_17), int32(326), int32(_a_F_IvfflatKmeans_18))
	mBase = m.M
	v4744 = m.ExcPending
	if v4744 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
