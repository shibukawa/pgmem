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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v184 int32
	_ = v184
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v210 float64
	_ = v210
	var v214 int32
	_ = v214
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
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
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
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
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
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
	var v413 int32
	_ = v413
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
	var v422 int32
	_ = v422
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
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v704 int32
	_ = v704
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 float64
	_ = v751
	var v762 int32
	_ = v762
	var v801 float64
	_ = v801
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int64
	_ = v824
	var v825 int32
	_ = v825
	var v826 float64
	_ = v826
	var v829 float64
	_ = v829
	var v832 int32
	_ = v832
	var v833 float32
	_ = v833
	var v834 float64
	_ = v834
	var v836 float32
	_ = v836
	var v840 float64
	_ = v840
	var v841 float64
	_ = v841
	var v843 int32
	_ = v843
	var v888 float64
	_ = v888
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int64
	_ = v900
	var v901 int64
	_ = v901
	var v902 int64
	_ = v902
	var v923 float64
	_ = v923
	var v931 int32
	_ = v931
	var v970 float64
	_ = v970
	var v978 float32
	_ = v978
	var v980 float64
	_ = v980
	var v984 int32
	_ = v984
	var v990 int32
	_ = v990
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1124 int32
	_ = v1124
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 float32
	_ = v1163
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1207 float32
	_ = v1207
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 float32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1228 float32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1234 float32
	_ = v1234
	var v1238 float32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 float32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 float32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 float32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 float32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1296 float32
	_ = v1296
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1344 float32
	_ = v1344
	var v1357 float32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 float32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1373 int32
	_ = v1373
	var v1404 float32
	_ = v1404
	var v1415 int32
	_ = v1415
	var v1421 int32
	_ = v1421
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1494 int32
	_ = v1494
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1543 int32
	_ = v1543
	var v1585 int32
	_ = v1585
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1606 int32
	_ = v1606
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1660 int64
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1665 float32
	_ = v1665
	var v1673 int32
	_ = v1673
	var v1730 int32
	_ = v1730
	var v1776 int32
	_ = v1776
	var v1777 float32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1785 int32
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1819 float32
	_ = v1819
	var v1833 float32
	_ = v1833
	var v1835 float32
	_ = v1835
	var v1836 float32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1844 float32
	_ = v1844
	var v1846 float32
	_ = v1846
	var v1847 float32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1862 int32
	_ = v1862
	var v1894 float32
	_ = v1894
	var v1908 float32
	_ = v1908
	var v1910 float32
	_ = v1910
	var v1949 float32
	_ = v1949
	var v1964 int32
	_ = v1964
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 float32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2079 float32
	_ = v2079
	var v2084 int32
	_ = v2084
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2096 int32
	_ = v2096
	var v2134 int32
	_ = v2134
	var v2136 float32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 float32
	_ = v2140
	var v2147 float32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2155 int64
	_ = v2155
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2167 int64
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2174 float32
	_ = v2174
	var v2177 float32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 float32
	_ = v2179
	var v2180 float32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2191 float32
	_ = v2191
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2203 int64
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2206 float32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2226 int32
	_ = v2226
	var v2238 int32
	_ = v2238
	var v2277 int32
	_ = v2277
	var v2296 int32
	_ = v2296
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2351 int32
	_ = v2351
	var v2353 int32
	_ = v2353
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2422 int32
	_ = v2422
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2438 int32
	_ = v2438
	var v2482 int32
	_ = v2482
	var v2546 int32
	_ = v2546
	var v2552 int32
	_ = v2552
	var v2595 int32
	_ = v2595
	var v2602 int32
	_ = v2602
	var v2646 int32
	_ = v2646
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2656 int32
	_ = v2656
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2731 int32
	_ = v2731
	var v2733 int32
	_ = v2733
	var v2775 int32
	_ = v2775
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2786 int32
	_ = v2786
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2794 int32
	_ = v2794
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2802 int32
	_ = v2802
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2821 int32
	_ = v2821
	var v2869 int32
	_ = v2869
	var v2872 int32
	_ = v2872
	var v2913 int32
	_ = v2913
	var v2916 int32
	_ = v2916
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2927 int32
	_ = v2927
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2982 int32
	_ = v2982
	var v2994 int32
	_ = v2994
	var v3033 int32
	_ = v3033
	var v3035 int32
	_ = v3035
	var v3039 int32
	_ = v3039
	var v3042 int32
	_ = v3042
	var v3049 int32
	_ = v3049
	var v3096 int32
	_ = v3096
	var v3099 int64
	_ = v3099
	var v3100 int64
	_ = v3100
	var v3101 int64
	_ = v3101
	var v3122 float64
	_ = v3122
	var v3126 int32
	_ = v3126
	var v3130 int32
	_ = v3130
	var v3136 int32
	_ = v3136
	var v3138 int32
	_ = v3138
	var v3182 int32
	_ = v3182
	var v3183 float32
	_ = v3183
	var v3191 float32
	_ = v3191
	var v3193 float32
	_ = v3193
	var v3201 float32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3206 int32
	_ = v3206
	var v3214 int32
	_ = v3214
	var v3260 int32
	_ = v3260
	var v3261 float32
	_ = v3261
	var v3269 float32
	_ = v3269
	var v3319 float32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3370 int32
	_ = v3370
	var v3372 int32
	_ = v3372
	var v3373 float32
	_ = v3373
	var v3376 float32
	_ = v3376
	var v3380 int32
	_ = v3380
	var v3382 int32
	_ = v3382
	var v3393 int32
	_ = v3393
	var v3436 int32
	_ = v3436
	var v3437 float32
	_ = v3437
	var v3489 int32
	_ = v3489
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3547 int32
	_ = v3547
	var v3591 int32
	_ = v3591
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3597 int32
	_ = v3597
	var v3602 int32
	_ = v3602
	var v3604 int32
	_ = v3604
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3658 int32
	_ = v3658
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3666 int32
	_ = v3666
	var v3668 int32
	_ = v3668
	var v3670 int32
	_ = v3670
	var v3677 int32
	_ = v3677
	var v3721 int32
	_ = v3721
	var v3723 int32
	_ = v3723
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3733 int32
	_ = v3733
	var v3734 int32
	_ = v3734
	var v3738 int64
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3744 int32
	_ = v3744
	var v3794 int32
	_ = v3794
	var v3804 int32
	_ = v3804
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3857 int32
	_ = v3857
	var v3859 int32
	_ = v3859
	var v3902 int32
	_ = v3902
	var v3903 int32
	_ = v3903
	var v3904 float32
	_ = v3904
	var v3905 float32
	_ = v3905
	var v3907 float32
	_ = v3907
	var v3908 float32
	_ = v3908
	var v3911 float32
	_ = v3911
	var v3914 int32
	_ = v3914
	var v3915 int32
	_ = v3915
	var v3916 float32
	_ = v3916
	var v3917 float32
	_ = v3917
	var v3919 float32
	_ = v3919
	var v3920 float32
	_ = v3920
	var v3923 float32
	_ = v3923
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3928 int32
	_ = v3928
	var v3936 int32
	_ = v3936
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3983 float32
	_ = v3983
	var v3984 float32
	_ = v3984
	var v3986 float32
	_ = v3986
	var v3987 float32
	_ = v3987
	var v3990 float32
	_ = v3990
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4051 int32
	_ = v4051
	var v4054 int32
	_ = v4054
	var v4095 int32
	_ = v4095
	var v4096 int32
	_ = v4096
	var v4097 int32
	_ = v4097
	var v4098 float32
	_ = v4098
	var v4100 int32
	_ = v4100
	var v4104 float32
	_ = v4104
	var v4108 int32
	_ = v4108
	var v4109 int32
	_ = v4109
	var v4110 float32
	_ = v4110
	var v4112 int32
	_ = v4112
	var v4116 float32
	_ = v4116
	var v4120 int32
	_ = v4120
	var v4122 int32
	_ = v4122
	var v4130 int32
	_ = v4130
	var v4174 int32
	_ = v4174
	var v4175 int32
	_ = v4175
	var v4176 int32
	_ = v4176
	var v4177 float32
	_ = v4177
	var v4179 int32
	_ = v4179
	var v4183 float32
	_ = v4183
	var v4234 int32
	_ = v4234
	var v4241 int32
	_ = v4241
	var v4285 int32
	_ = v4285
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4292 int32
	_ = v4292
	var v4294 int32
	_ = v4294
	var v4348 int32
	_ = v4348
	var v4399 int32
	_ = v4399
	var v4400 int32
	_ = v4400
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4406 int32
	_ = v4406
	var v4409 int32
	_ = v4409
	var v4412 int32
	_ = v4412
	var v4418 int32
	_ = v4418
	var v4462 int32
	_ = v4462
	var v4468 int32
	_ = v4468
	var v4470 int32
	_ = v4470
	var v4471 int32
	_ = v4471
	var v4474 int32
	_ = v4474
	var v4476 int32
	_ = v4476
	var v4477 int32
	_ = v4477
	var v4478 int32
	_ = v4478
	var v4485 int32
	_ = v4485
	var v4532 float32
	_ = v4532
	var v4544 int32
	_ = v4544
	var v4548 int32
	_ = v4548
	var v4553 int32
	_ = v4553
	var v4555 int32
	_ = v4555
	var v4606 int32
	_ = v4606
	var v4607 int32
	_ = v4607
	var v4658 int32
	_ = v4658
	var v4659 int32
	_ = v4659
	var v4662 int32
	_ = v4662
	var v4665 int32
	_ = v4665
	var v4666 int32
	_ = v4666
	var v4672 int32
	_ = v4672
	var v4716 int32
	_ = v4716
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4723 int64
	_ = v4723
	var v4724 int32
	_ = v4724
	var v4730 int32
	_ = v4730
	var v4731 int32
	_ = v4731
	var v4736 int32
	_ = v4736
	var v4740 int32
	_ = v4740
	var v4745 int32
	_ = v4745
	var v4797 int32
	_ = v4797
	var v4801 int32
	_ = v4801
	var v4805 int32
	_ = v4805
	var v4810 int32
	_ = v4810
	var v4814 int32
	_ = v4814
	var v4818 int32
	_ = v4818
	var v4823 int32
	_ = v4823
	var v4827 int32
	_ = v4827
	var v4831 int32
	_ = v4831
	var v4836 int32
	_ = v4836
	var v4888 int32
	_ = v4888
	var v4892 int32
	_ = v4892
	var v4897 int32
	_ = v4897
	v6 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	v55 = F_AllocSetContextCreateInternal(m, v50, int32(_a_F_IvfflatKmeans_0), v6, int32(_a_F_IvfflatKmeans_1), int32(_a_F_IvfflatKmeans_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v57 = int32(_a_F_IvfflatKmeans_3)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0])) = v55
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v62 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4888 = m.ExcPending
	if v4888 != 0 {
		goto L1
	} else {
		goto L450
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4827 = m.ExcPending
	if v4827 != 0 {
		goto L1
	} else {
		goto L447
	}
L5:
	;
	v4399 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v4400 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4399 == v4400 {
		goto L397
	} else {
		goto L398
	}
L6:
	;
	v66 = F_HnswOptionalProcInfo(m, l0, int32(4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v342 = F_mul_size(m, v336, (v337+int32(7))&int32(-8))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L32
	}
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v71 = F_palloc_mul(m, int32(4), v61)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v73 < v74 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v80 = v73
	goto L14
L12:
	;
	goto L13
L13:
	;
	if v66 == int32(0) {
		goto L5
	} else {
		goto L28
	}
L14:
	;
	if int32(0) <= v80 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L13
L16:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	m.T0[v264].(func(*base.Module, int32, int32, int32))(m, v126+v127*v80, v61, v71)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L26
	}
L17:
	;
	v137 = v130
	goto L22
L18:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v130 = int32(0)
	if v130 < v61 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L3
L21:
	;
	goto L16
L22:
	;
	v184 = int32(_a_F_IvfflatKmeans_4)
	v187 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1]))
	v188 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2]))
	v189 = v187 ^ v188
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2])) = base.I64_rotl(v189, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1])) = v189<<(uint(int64(16))%64) ^ base.I64_rotl(v187, int64(24)) ^ v189
	v210 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v187*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L24
L23:
	;
	goto L16
L24:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v71+v137<<(uint(int32(2))%32)))) = base.F32_demote_f64(v210)
	v214 = v137 + int32(1)
	if v214 != v61 {
		v137 = v214
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v269 = v267 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v269 < v271 {
		v80 = v269
		goto L14
	} else {
		goto L27
	}
L27:
	;
	goto L15
L28:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	v329 = F_AllocSetContextCreateInternal(m, v324, int32(_a_F_IvfflatKmeans_5), int32(0), int32(_a_F_IvfflatKmeans_1), int32(_a_F_IvfflatKmeans_2))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_IvfflatNormVectors(m, l3, v69, l2, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_MemoryContextDelete(m, v329)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L5
L32:
	;
	v344 = F_add_size(m, int32(20), v342)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v347 = F_mul_size(m, v336, v61)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v349 = F_mul_size(m, int32(4), v347)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v352 = F_mul_size(m, int32(4), v336)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v355 = F_mul_size(m, int32(4), v62)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v358 = F_mul_size(m, v62, v336)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v360 = F_mul_size(m, int32(4), v358)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v363 = F_mul_size(m, int32(4), v62)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v366 = F_mul_size(m, int32(4), v336)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v369 = F_mul_size(m, v336, v336)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v371 = F_mul_size(m, int32(4), v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v374 = F_mul_size(m, int32(4), v336)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v376 = F_add_size(m, l4, v344)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v378 = F_add_size(m, v376, v349)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v380 = F_add_size(m, v378, v352)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v382 = F_add_size(m, v380, v355)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v384 = F_add_size(m, v382, v360)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v386 = F_add_size(m, v384, v363)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v388 = F_add_size(m, v386, v366)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v390 = F_add_size(m, v388, v371)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v392 = F_add_size(m, v390, v374)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_IvfflatCheckMemoryUsage(m, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v397 = base.I32_div_s(int32(2147483647), v336)
	if v397 < v336 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v401 = F_index_getprocinfo(m, l0, int32(1), int32(3))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v404 = F_HnswOptionalProcInfo(m, l0, int32(4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v408 = F_palloc(m, v349)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v410 = F_palloc(m, v352)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v412 = F_palloc(m, v355)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v415 = F_palloc_extended(m, v360, int32(1))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v417 = F_palloc(m, v363)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v419 = F_palloc(m, v366)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v422 = F_palloc_extended(m, v371, int32(1))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v424 = F_palloc(m, v374)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v427 = F_VectorArrayInit(m, v336, v61, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v336
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v432 = F_palloc_mul(m, int32(4), v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v438 = F_index_getprocinfo(m, l0, int32(1), int32(3))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v443 = Fn14349(m, int64(32))
	mBase = m.M
	goto L69
L69:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v445 = base.I32_rem_u_s(v443, v444)
	if v445 < int32(0) {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v448 <= v445 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v450 = int32(0)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_VectorArraySet(m, l2, v450, v452+v453*v445)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v458 + int32(1)
	if v434 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v435 <= int32(0) {
		goto L85
	} else {
		goto L86
	}
L74:
	;
	v465 = v434 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v434) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v475 = v450
	v478 = int32(0)
	goto L78
L76:
	;
	v541 = v450
	goto L77
L77:
	;
	v590 = v541
	v591 = int32(0)
	goto L82
L78:
	;
	v521 = v432 + v475<<(uint(int32(2))%32)
	v522 = int64(9187343237679939583)
	*(*int64)(unsafe.Add(mBase, uint32(v521)+24)) = v522
	*(*int64)(unsafe.Add(mBase, uint32(v521)+16)) = v522
	*(*int64)(unsafe.Add(mBase, uint32(v521)+8)) = v522
	*(*int64)(unsafe.Add(mBase, uint32(v521))) = v522
	v530 = int32(8)
	v531 = v475 + v530
	v533 = v478 + v530
	if v533 != v434&int32(2147483640) {
		v475 = v531
		v478 = v533
		goto L78
	} else {
		goto L80
	}
L79:
	;
	if v465 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v541 = v531
	goto L77
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432+v590<<(uint(int32(2))%32)))) = int32(2139095039)
	v639 = int32(1)
	v642 = v591 + v639
	if v642 != v465 {
		v590 = v590 + v639
		v591 = v642
		goto L82
	} else {
		goto L84
	}
L83:
	;
	goto L73
L84:
	;
	goto L83
L85:
	;
	F_pfree(m, v432)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L115
	}
L86:
	;
	v695 = v434 - int32(1)
	v696 = int32(0)
	v704 = v696
	goto L87
L87:
	;
	v748 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[3]))
	if v748 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v751 = float64(0)
	if base.B2i32(v434 <= v696) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	v762 = int32(0)
	v801 = v751
	goto L96
L94:
	;
	v888 = v751
	goto L95
L95:
	;
	v894 = v704 + int32(1)
	if v894 == v435 {
		goto L85
	} else {
		goto L105
	}
L96:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v806 <= v762 {
		goto L3
	} else {
		goto L98
	}
L97:
	;
	v888 = v841
	goto L95
L98:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v808 <= v704 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v824 = F_FunctionCall2Coll(m, v438, v441, base.I64_extend_i32_u(v814+v815*v762), base.I64_extend_i32_u(v819+v820*v704))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v826 = base.F64_reinterpret_i64(v824)
	*(*float32)(unsafe.Add(mBase, uint32(v415+v704<<(uint(int32(2))%32)+v762*v435<<(uint(int32(2))%32)))) = base.F32_demote_f64(v826)
	v829 = base.F64_mul(v826, v826)
	v832 = v432 + v762<<(uint(int32(2))%32)
	v833 = *(*float32)(unsafe.Add(mBase, uint32(v832)))
	v834 = base.F64_promote_f32(v833)
	if base.F64_lt(v829, v834) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v836 = base.F32_demote_f64(v829)
	*(*float32)(unsafe.Add(mBase, uint32(v832))) = v836
	v840 = base.F64_promote_f32(v836)
	goto L103
L102:
	;
	v840 = v834
	goto L103
L103:
	;
	v841 = base.F64_add(v840, v801)
	v843 = v762 + int32(1)
	if v843 != v434 {
		v762 = v843
		v801 = v841
		goto L96
	} else {
		goto L104
	}
L104:
	;
	goto L97
L105:
	;
	v896 = int32(0)
	v897 = int32(_a_F_IvfflatKmeans_4)
	v900 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1]))
	v901 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2]))
	v902 = v900 ^ v901
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2])) = base.I64_rotl(v902, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1])) = v902<<(uint(int64(16))%64) ^ base.I64_rotl(v900, int64(24)) ^ v902
	v923 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v900*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L106
L106:
	;
	if v434 < int32(2) {
		v990 = v896
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1034 <= v990 {
		goto L3
	} else {
		goto L113
	}
L108:
	;
	v931 = v896
	v970 = base.F64_mul(v923, v888)
	goto L109
L109:
	;
	v978 = *(*float32)(unsafe.Add(mBase, uint32(v432+v931<<(uint(int32(2))%32))))
	v980 = base.F64_sub(v970, base.F64_promote_f32(v978))
	if base.F64_le(v980, float64(0)) != 0 {
		v990 = v931
		goto L107
	} else {
		goto L111
	}
L110:
	;
	v990 = v695
	goto L107
L111:
	;
	v984 = v931 + int32(1)
	if v984 != v695 {
		v931 = v984
		v970 = v980
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_VectorArraySet(m, l2, v894, v1036+v1037*v990)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1042 + int32(1)
	v704 = v894
	goto L87
L115:
	;
	if int32(0) < v62 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v1100 = int32(3)
	v1101 = v336 & v1100
	v1124 = v6
	goto L119
L117:
	;
	goto L118
L118:
	;
	v1471 = int32(2147483646)
	v1473 = int32(1)
	v1476 = v336 & v1471
	v1478 = v336 & v1473
	v1480 = v336 - v1473
	v1494 = int32(0)
	goto L166
L119:
	;
	if v336 <= int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L118
L121:
	;
	v1415 = v1124 << (uint(int32(2)) % 32)
	*(*float32)(unsafe.Add(mBase, uint32(v417+v1415))) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v1415+v412))) = v1373
	v1421 = v1124 + int32(1)
	if v1421 != v62 {
		v1124 = v1421
		goto L119
	} else {
		goto L165
	}
L122:
	;
	v1373 = int32(0)
	v1404 = float32(3.4028235e+38)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v1161 = v415 + v336*v1124<<(uint(int32(2))%32)
	v1162 = int32(0)
	v1163 = float32(3.4028235e+38)
	if base.B2i32(base.Ui32(v336-int32(1)) < base.Ui32(v1100)) == v1162 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v1173 = v1162
	v1174 = v1162
	v1176 = v1162
	v1207 = v1163
	goto L128
L126:
	;
	v1263 = v1162
	v1265 = v1162
	v1296 = v1163
	goto L127
L127:
	;
	v1311 = v1263
	v1312 = v1162
	v1313 = v1265
	v1344 = v1296
	goto L156
L128:
	;
	v1218 = v1173 | int32(3)
	v1219 = int32(2)
	v1222 = *(*float32)(unsafe.Add(mBase, uint32(v1161+v1218<<(uint(v1219)%32))))
	v1224 = v1173 | v1219
	v1228 = *(*float32)(unsafe.Add(mBase, uint32(v1161+v1224<<(uint(v1219)%32))))
	v1230 = v1173 | int32(1)
	v1234 = *(*float32)(unsafe.Add(mBase, uint32(v1161+v1230<<(uint(v1219)%32))))
	v1238 = *(*float32)(unsafe.Add(mBase, uint32(v1161+v1173<<(uint(v1219)%32))))
	v1239 = base.F32_gt(v1207, v1238)
	if v1239 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v1101 == int32(0) {
		v1373 = v1250
		v1404 = v1246
		goto L121
	} else {
		goto L155
	}
L130:
	;
	v1240 = v1238
	goto L132
L131:
	;
	v1240 = v1207
	goto L132
L132:
	;
	v1241 = base.F32_gt(v1240, v1234)
	if v1241 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v1242 = v1234
	goto L135
L134:
	;
	v1242 = v1240
	goto L135
L135:
	;
	v1243 = base.F32_gt(v1242, v1228)
	if v1243 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v1244 = v1228
	goto L138
L137:
	;
	v1244 = v1242
	goto L138
L138:
	;
	v1245 = base.F32_gt(v1244, v1222)
	if v1245 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v1246 = v1222
	goto L141
L140:
	;
	v1246 = v1244
	goto L141
L141:
	;
	if v1239 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v1247 = v1173
	goto L144
L143:
	;
	v1247 = v1176
	goto L144
L144:
	;
	if v1241 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v1248 = v1230
	goto L147
L146:
	;
	v1248 = v1247
	goto L147
L147:
	;
	if v1243 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1249 = v1224
	goto L150
L149:
	;
	v1249 = v1248
	goto L150
L150:
	;
	if v1245 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v1250 = v1218
	goto L153
L152:
	;
	v1250 = v1249
	goto L153
L153:
	;
	v1251 = int32(4)
	v1252 = v1173 + v1251
	v1254 = v1174 + v1251
	if v1254 != v336&int32(2147483644) {
		v1173 = v1252
		v1174 = v1254
		v1176 = v1250
		v1207 = v1246
		goto L128
	} else {
		goto L154
	}
L154:
	;
	goto L129
L155:
	;
	v1263 = v1252
	v1265 = v1250
	v1296 = v1246
	goto L127
L156:
	;
	v1357 = *(*float32)(unsafe.Add(mBase, uint32(v1161+v1311<<(uint(int32(2))%32))))
	v1358 = base.F32_gt(v1344, v1357)
	if v1358 != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v1373 = v1360
	v1404 = v1359
	goto L121
L158:
	;
	v1359 = v1357
	goto L160
L159:
	;
	v1359 = v1344
	goto L160
L160:
	;
	if v1358 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1360 = v1311
	goto L163
L162:
	;
	v1360 = v1313
	goto L163
L163:
	;
	v1361 = int32(1)
	v1364 = v1312 + v1361
	if v1364 != v1101 {
		v1311 = v1311 + v1361
		v1312 = v1364
		v1313 = v1360
		v1344 = v1359
		goto L156
	} else {
		goto L164
	}
L164:
	;
	goto L157
L165:
	;
	goto L120
L166:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[3]))
	if v1531 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L5
L168:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v1534 = int32(0)
	if v1534 < v336 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	goto L170
L172:
	;
	v1543 = v1534
	goto L175
L173:
	;
	goto L174
L174:
	;
	v2014 = int32(0)
	if v2014 < v62 {
		goto L216
	} else {
		goto L217
	}
L175:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1543 < v1585 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v1730 = int32(0)
	goto L188
L177:
	;
	if v1588 != v336 {
		v1543 = v1588
		goto L175
	} else {
		goto L187
	}
L178:
	;
	v1588 = v1543 + int32(1)
	if v336 <= v1588 {
		goto L177
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	goto L3
L181:
	;
	v1590 = int32(2)
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1606 = v1588
	goto L182
L182:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1650 <= v1606 {
		goto L3
	} else {
		goto L184
	}
L183:
	;
	goto L177
L184:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1660 = F_FunctionCall2Coll(m, v401, v407, base.I64_extend_i32_u(v1593+v1594*v1543), base.I64_extend_i32_u(v1655+v1656*v1606))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v1665 = base.F32_demote_f64(base.F64_mul(base.F64_reinterpret_i64(v1660), float64(0.5)))
	*(*float32)(unsafe.Add(mBase, uint32(v422+v1543*v336<<(uint(v1590)%32)+v1606<<(uint(int32(2))%32)))) = v1665
	*(*float32)(unsafe.Add(mBase, uint32(v422+v1543<<(uint(v1590)%32)+v1606*v336<<(uint(int32(2))%32)))) = v1665
	v1673 = v1606 + int32(1)
	if v336 != v1673 {
		v1606 = v1673
		goto L182
	} else {
		goto L186
	}
L186:
	;
	goto L183
L187:
	;
	goto L176
L188:
	;
	v1776 = v422 + v1730*v336<<(uint(int32(2))%32)
	v1777 = float32(3.4028235e+38)
	v1778 = int32(0)
	if v1480 != 0 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	goto L174
L190:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v419+v1730<<(uint(int32(2))%32)))) = v1949
	v1964 = v1730 + int32(1)
	if v1964 != v336 {
		v1730 = v1964
		goto L188
	} else {
		goto L214
	}
L191:
	;
	v1785 = v1778
	v1788 = v1778
	v1819 = v1777
	goto L194
L192:
	;
	v1862 = v1778
	v1894 = v1777
	goto L193
L193:
	;
	if v1730 == v1862 {
		v1949 = v1894
		goto L190
	} else {
		goto L210
	}
L194:
	;
	if v1785 != v1730 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v1478 == int32(0) {
		v1949 = v1847
		goto L190
	} else {
		goto L209
	}
L196:
	;
	v1833 = *(*float32)(unsafe.Add(mBase, uint32(v1776+v1785<<(uint(int32(2))%32))))
	if base.F32_gt(v1819, v1833) != 0 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v1836 = v1819
	goto L198
L198:
	;
	v1839 = v1785 | int32(1)
	if v1839 != v1730 {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v1835 = v1833
	goto L201
L200:
	;
	v1835 = v1819
	goto L201
L201:
	;
	v1836 = v1835
	goto L198
L202:
	;
	v1844 = *(*float32)(unsafe.Add(mBase, uint32(v1776+v1839<<(uint(int32(2))%32))))
	if base.F32_gt(v1836, v1844) != 0 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v1847 = v1836
	goto L204
L204:
	;
	v1849 = int32(2)
	v1850 = v1785 + v1849
	v1852 = v1788 + v1849
	if v1852 != v1476 {
		v1785 = v1850
		v1788 = v1852
		v1819 = v1847
		goto L194
	} else {
		goto L208
	}
L205:
	;
	v1846 = v1844
	goto L207
L206:
	;
	v1846 = v1836
	goto L207
L207:
	;
	v1847 = v1846
	goto L204
L208:
	;
	goto L195
L209:
	;
	v1862 = v1850
	v1894 = v1847
	goto L193
L210:
	;
	v1908 = *(*float32)(unsafe.Add(mBase, uint32(v1776+v1862<<(uint(int32(2))%32))))
	if base.F32_gt(v1894, v1908) != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1910 = v1908
	goto L213
L212:
	;
	v1910 = v1894
	goto L213
L213:
	;
	v1949 = v1910
	goto L190
L214:
	;
	goto L189
L215:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v2333 = int32(0)
	v2334 = base.B2i32(v2332 <= v2333)
	if v2334 == v2333 {
		goto L246
	} else {
		goto L247
	}
L216:
	;
	v2017 = int32(0)
	v2028 = v2014
	v2030 = v2017
	goto L219
L217:
	;
	goto L218
L218:
	;
	v2296 = int32(1)
	goto L215
L219:
	;
	if v336 <= int32(0) {
		v2238 = v2030
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v2296 = base.B2i32(v2238 == int32(0))
	goto L215
L221:
	;
	v2277 = v2028 + int32(1)
	if v2277 != v62 {
		v2028 = v2277
		v2030 = v2238
		goto L219
	} else {
		goto L245
	}
L222:
	;
	v2070 = int32(2)
	v2071 = v2028 << (uint(v2070) % 32)
	v2072 = v417 + v2071
	v2073 = *(*float32)(unsafe.Add(mBase, uint32(v2072)))
	v2074 = v2071 + v412
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2074)))
	v2079 = *(*float32)(unsafe.Add(mBase, uint32(v419+v2075<<(uint(v2070)%32))))
	if base.F32_le(v2073, v2079) != 0 {
		v2238 = v2030
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v2084 = v415 + v2028*v336<<(uint(int32(2))%32)
	v2090 = int32(0)
	v2092 = base.B2i32(v1494 != v2017)
	v2096 = v2030
	goto L224
L224:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2074)))
	if v2090 == v2134 {
		v2217 = v2092
		v2218 = v2096
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v2238 = v2218
	goto L221
L226:
	;
	v2226 = v2090 + int32(1)
	if v2226 != v336 {
		v2090 = v2226
		v2092 = v2217
		v2096 = v2218
		goto L224
	} else {
		goto L244
	}
L227:
	;
	v2136 = *(*float32)(unsafe.Add(mBase, uint32(v2072)))
	v2138 = v2090 << (uint(int32(2)) % 32)
	v2139 = v2084 + v2138
	v2140 = *(*float32)(unsafe.Add(mBase, uint32(v2139)))
	if base.F32_le(v2136, v2140) != 0 {
		v2217 = v2092
		v2218 = v2096
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v2147 = *(*float32)(unsafe.Add(mBase, uint32(v422+v2134*v336<<(uint(int32(2))%32)+v2138)))
	if base.F32_le(v2136, v2147) != 0 {
		v2217 = v2092
		v2218 = v2096
		goto L226
	} else {
		goto L229
	}
L229:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2149 <= v2028 {
		goto L3
	} else {
		goto L230
	}
L230:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2155 = base.I64_extend_i32_u(v2151 + v2152*v2028)
	if v2092&int32(1) != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	if v2134 < int32(0) {
		goto L3
	} else {
		goto L234
	}
L232:
	;
	v2178 = v2134
	v2179 = v2136
	v2180 = v2140
	goto L233
L233:
	;
	if base.F32_gt(v2179, v2180) == int32(0) {
		goto L237
	} else {
		goto L238
	}
L234:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2160 <= v2134 {
		goto L3
	} else {
		goto L235
	}
L235:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2167 = F_FunctionCall2Coll(m, v401, v407, v2155, base.I64_extend_i32_u(v2162+v2163*v2134))
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v2074)))
	v2174 = base.F32_demote_f64(base.F64_reinterpret_i64(v2167))
	*(*float32)(unsafe.Add(mBase, uint32(v2084+v2169<<(uint(int32(2))%32)))) = v2174
	*(*float32)(unsafe.Add(mBase, uint32(v2072))) = v2174
	v2177 = *(*float32)(unsafe.Add(mBase, uint32(v2139)))
	v2178 = v2169
	v2179 = v2174
	v2180 = v2177
	goto L233
L237:
	;
	v2185 = int32(0)
	v2191 = *(*float32)(unsafe.Add(mBase, uint32(v422+v2178*v336<<(uint(int32(2))%32)+v2138)))
	if base.F32_gt(v2179, v2191) == v2185 {
		v2217 = v2185
		v2218 = v2096
		goto L226
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2196 <= v2090 {
		goto L3
	} else {
		goto L241
	}
L240:
	;
	goto L239
L241:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2203 = F_FunctionCall2Coll(m, v401, v407, v2155, base.I64_extend_i32_u(v2198+v2199*v2090))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v2206 = base.F32_demote_f64(base.F64_reinterpret_i64(v2203))
	*(*float32)(unsafe.Add(mBase, uint32(v2139))) = v2206
	v2208 = int32(0)
	if base.F32_gt(v2179, v2206) == v2208 {
		v2217 = v2208
		v2218 = v2096
		goto L226
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2074))) = v2090
	*(*float32)(unsafe.Add(mBase, uint32(v2072))) = v2206
	v2217 = v2208
	v2218 = v2096 + int32(1)
	goto L226
L244:
	;
	goto L225
L245:
	;
	goto L220
L246:
	;
	v2338 = v2330 << (uint(int32(2)) % 32)
	v2339 = int32(0)
	if v2332 != int32(1) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v2552 = v2331
	goto L248
L248:
	;
	v2595 = int32(0)
	if v2595 < v2552 {
		goto L267
	} else {
		goto L268
	}
L249:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2552 = v2546
	goto L248
L250:
	;
	v2351 = v2339
	v2353 = int32(0)
	goto L253
L251:
	;
	v2438 = v2339
	goto L252
L252:
	;
	v2482 = int32(0)
	if base.B2i32(v2338 == v2482)|base.B2i32(v2330 <= v2482) == v2482 {
		goto L263
	} else {
		goto L264
	}
L253:
	;
	v2395 = int32(0)
	v2396 = base.B2i32(v2330 <= v2395)
	if v2396|base.B2i32(v2338 == v2395) == v2395 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	if v2332&int32(1) == int32(0) {
		goto L249
	} else {
		goto L262
	}
L255:
	;
	base.MemoryFill(m, v408+v2351*v2338, int32(0), v2338)
	goto L257
L256:
	;
	goto L257
L257:
	;
	v2409 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v410+v2351<<(uint(int32(2))%32)))) = v2409
	v2412 = v2351 | int32(1)
	if v2396|base.B2i32(v2338 == v2409) == v2409 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	base.MemoryFill(m, v408+v2412*v2338, int32(0), v2338)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v2422 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v410+v2412<<(uint(v2422)%32)))) = int32(0)
	v2428 = v2351 + v2422
	v2430 = v2353 + v2422
	if v2430 != v2332&int32(2147483646) {
		v2351 = v2428
		v2353 = v2430
		goto L253
	} else {
		goto L261
	}
L261:
	;
	goto L254
L262:
	;
	v2438 = v2428
	goto L252
L263:
	;
	base.MemoryFill(m, v408+v2438*v2338, int32(0), v2338)
	goto L265
L264:
	;
	goto L265
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410+v2438<<(uint(int32(2))%32)))) = int32(0)
	goto L249
L266:
	;
	v3539 = int32(0)
	v3540 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	if v3539 < v3540 {
		goto L333
	} else {
		goto L334
	}
L267:
	;
	v2602 = v2595
	goto L270
L268:
	;
	goto L269
L269:
	;
	if v2331 <= int32(0) {
		goto L275
	} else {
		goto L276
	}
L270:
	;
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2646 <= v2602 {
		goto L3
	} else {
		goto L272
	}
L271:
	;
	goto L269
L272:
	;
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2653 = int32(2)
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v412+v2602<<(uint(v2653)%32))))
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	m.T0[v2661].(func(*base.Module, int32, int32))(m, v2648+v2649*v2602, v408+v2652*v2656<<(uint(v2653)%32))
	mBase = m.M
	v2663 = m.ExcPending
	if v2663 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v2665 = v2602 + int32(1)
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2665 < v2666 {
		v2602 = v2665
		goto L270
	} else {
		goto L274
	}
L274:
	;
	goto L271
L275:
	;
	if v2332 <= v2333 {
		goto L266
	} else {
		goto L287
	}
L276:
	;
	v2719 = v2331 & int32(3)
	v2720 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v2331) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v2731 = v2720
	v2733 = int32(0)
	goto L280
L278:
	;
	v2821 = v2720
	goto L279
L279:
	;
	v2869 = v2821
	v2872 = v2720
	goto L284
L280:
	;
	v2775 = int32(2)
	v2777 = v412 + v2731<<(uint(v2775)%32)
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v2777)))
	v2781 = v410 + v2778<<(uint(v2775)%32)
	v2782 = *(*int32)(unsafe.Add(mBase, uint32(v2781)))
	v2783 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2781))) = v2782 + v2783
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2777)+4))
	v2789 = v410 + v2786<<(uint(v2775)%32)
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v2789)))
	*(*int32)(unsafe.Add(mBase, uint32(v2789))) = v2790 + v2783
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2777)+8))
	v2797 = v410 + v2794<<(uint(v2775)%32)
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(v2797)))
	*(*int32)(unsafe.Add(mBase, uint32(v2797))) = v2798 + v2783
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(v2777)+12))
	v2805 = v410 + v2802<<(uint(v2775)%32)
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2805)))
	*(*int32)(unsafe.Add(mBase, uint32(v2805))) = v2806 + v2783
	v2810 = int32(4)
	v2811 = v2731 + v2810
	v2813 = v2733 + v2810
	if v2813 != v2331&int32(2147483644) {
		v2731 = v2811
		v2733 = v2813
		goto L280
	} else {
		goto L282
	}
L281:
	;
	if v2719 == int32(0) {
		goto L275
	} else {
		goto L283
	}
L282:
	;
	goto L281
L283:
	;
	v2821 = v2811
	goto L279
L284:
	;
	v2913 = int32(2)
	v2916 = *(*int32)(unsafe.Add(mBase, uint32(v412+v2869<<(uint(v2913)%32))))
	v2919 = v410 + v2916<<(uint(v2913)%32)
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2919)))
	v2921 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2919))) = v2920 + v2921
	v2927 = v2872 + v2921
	if v2927 != v2719 {
		v2869 = v2869 + v2921
		v2872 = v2927
		goto L284
	} else {
		goto L286
	}
L285:
	;
	goto L275
L286:
	;
	goto L285
L287:
	;
	v2978 = v2330 & int32(2147483646)
	v2979 = int32(1)
	v2980 = v2330 & v2979
	v2982 = v2330 - v2979
	v2994 = int32(0)
	goto L288
L288:
	;
	v3033 = int32(2)
	v3035 = v408 + v2330*v2994<<(uint(v3033)%32)
	v3039 = *(*int32)(unsafe.Add(mBase, uint32(v410+v2994<<(uint(v3033)%32))))
	if v3039 <= int32(0) {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	goto L266
L290:
	;
	v3489 = v2994 + int32(1)
	if v3489 != v2332 {
		v2994 = v3489
		goto L288
	} else {
		goto L331
	}
L291:
	;
	v3042 = int32(0)
	if v2330 <= v3042 {
		goto L290
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	if v2330 <= int32(0) {
		goto L290
	} else {
		goto L299
	}
L294:
	;
	v3049 = v3042
	goto L295
L295:
	;
	v3096 = int32(_a_F_IvfflatKmeans_4)
	v3099 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1]))
	v3100 = *(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2]))
	v3101 = v3099 ^ v3100
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[2])) = base.I64_rotl(v3101, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_IvfflatKmeans[1])) = v3101<<(uint(int64(16))%64) ^ base.I64_rotl(v3099, int64(24)) ^ v3101
	v3122 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v3099*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L297
L296:
	;
	goto L290
L297:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3035+v3049<<(uint(int32(2))%32)))) = base.F32_demote_f64(v3122)
	v3126 = v3049 + int32(1)
	if v3126 != v2330 {
		v3049 = v3126
		goto L295
	} else {
		goto L298
	}
L298:
	;
	goto L296
L299:
	;
	v3130 = int32(0)
	if v2982 != 0 {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v3319 = base.F32_convert_i32_u(v3039)
	v3320 = int32(0)
	if v2982 != 0 {
		goto L324
	} else {
		goto L325
	}
L301:
	;
	v3136 = v3130
	v3138 = v3130
	goto L304
L302:
	;
	v3214 = v3130
	goto L303
L303:
	;
	v3260 = v3035 + v3214<<(uint(int32(2))%32)
	v3261 = *(*float32)(unsafe.Add(mBase, uint32(v3260)))
	if base.F32_ne(base.F32_abs(v3261), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L300
	} else {
		goto L320
	}
L304:
	;
	v3182 = v3035 + v3136<<(uint(int32(2))%32)
	v3183 = *(*float32)(unsafe.Add(mBase, uint32(v3182)))
	if base.F32_eq(base.F32_abs(v3183), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	if v2980 == int32(0) {
		goto L300
	} else {
		goto L319
	}
L306:
	;
	if base.F32_gt(v3183, float32(0)) != 0 {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L308
L308:
	;
	v3193 = *(*float32)(unsafe.Add(mBase, uint32(v3182)+4))
	if base.F32_eq(base.F32_abs(v3193), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L312
	} else {
		goto L313
	}
L309:
	;
	v3191 = float32(3.4028235e+38)
	goto L311
L310:
	;
	v3191 = float32(-3.4028235e+38)
	goto L311
L311:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3182))) = v3191
	goto L308
L312:
	;
	if base.F32_gt(v3193, float32(0)) != 0 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	goto L314
L314:
	;
	v3203 = int32(2)
	v3204 = v3136 + v3203
	v3206 = v3138 + v3203
	if v3206 != v2978 {
		v3136 = v3204
		v3138 = v3206
		goto L304
	} else {
		goto L318
	}
L315:
	;
	v3201 = float32(3.4028235e+38)
	goto L317
L316:
	;
	v3201 = float32(-3.4028235e+38)
	goto L317
L317:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3182)+4)) = v3201
	goto L314
L318:
	;
	goto L305
L319:
	;
	v3214 = v3204
	goto L303
L320:
	;
	if base.F32_gt(v3261, float32(0)) != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v3269 = float32(3.4028235e+38)
	goto L323
L322:
	;
	v3269 = float32(-3.4028235e+38)
	goto L323
L323:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3260))) = v3269
	goto L300
L324:
	;
	v3328 = v3320
	v3329 = v3320
	goto L327
L325:
	;
	v3393 = v3320
	goto L326
L326:
	;
	v3436 = v3035 + v3393<<(uint(int32(2))%32)
	v3437 = *(*float32)(unsafe.Add(mBase, uint32(v3436)))
	*(*float32)(unsafe.Add(mBase, uint32(v3436))) = base.F32_div(v3437, v3319)
	goto L290
L327:
	;
	v3370 = int32(2)
	v3372 = v3035 + v3329<<(uint(v3370)%32)
	v3373 = *(*float32)(unsafe.Add(mBase, uint32(v3372)))
	*(*float32)(unsafe.Add(mBase, uint32(v3372))) = base.F32_div(v3373, v3319)
	v3376 = *(*float32)(unsafe.Add(mBase, uint32(v3372)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v3372)+4)) = base.F32_div(v3376, v3319)
	v3380 = v3329 + v3370
	v3382 = v3328 + v3370
	if v3382 != v2978 {
		v3328 = v3382
		v3329 = v3380
		goto L327
	} else {
		goto L329
	}
L328:
	;
	if v2980 == int32(0) {
		goto L290
	} else {
		goto L330
	}
L329:
	;
	goto L328
L330:
	;
	v3393 = v3380
	goto L326
L331:
	;
	goto L289
L332:
	;
	v3794 = int32(0)
	if v62 <= v3794 {
		goto L354
	} else {
		goto L355
	}
L333:
	;
	v3547 = v3539
	goto L336
L334:
	;
	goto L335
L335:
	;
	if v404 != 0 {
		goto L341
	} else {
		goto L342
	}
L336:
	;
	v3591 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v3591 <= v3547 {
		goto L3
	} else {
		goto L338
	}
L337:
	;
	goto L335
L338:
	;
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v427)+16))
	v3594 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v3597 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
	v3602 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	m.T0[v3602].(func(*base.Module, int32, int32, int32))(m, v3593+v3594*v3547, v3597, v408+v3547*v3597<<(uint(int32(2))%32))
	mBase = m.M
	v3604 = m.ExcPending
	if v3604 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	v3606 = v3547 + int32(1)
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	if v3606 < v3607 {
		v3547 = v3606
		goto L336
	} else {
		goto L340
	}
L340:
	;
	goto L337
L341:
	;
	v3658 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0]))
	v3663 = F_AllocSetContextCreateInternal(m, v3658, int32(_a_F_IvfflatKmeans_5), int32(0), int32(_a_F_IvfflatKmeans_1), int32(_a_F_IvfflatKmeans_2))
	mBase = m.M
	v3664 = m.ExcPending
	if v3664 != 0 {
		goto L1
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v3670 = int32(0)
	if v336 <= v3670 {
		goto L332
	} else {
		goto L347
	}
L344:
	;
	F_IvfflatNormVectors(m, l3, v407, v427, v3663)
	mBase = m.M
	v3666 = m.ExcPending
	if v3666 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	F_MemoryContextDelete(m, v3663)
	mBase = m.M
	v3668 = m.ExcPending
	if v3668 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	goto L343
L347:
	;
	v3677 = v3670
	goto L348
L348:
	;
	v3721 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3721 <= v3677 {
		goto L3
	} else {
		goto L350
	}
L349:
	;
	goto L332
L350:
	;
	v3723 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v3723 <= v3677 {
		goto L3
	} else {
		goto L351
	}
L351:
	;
	v3728 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v3729 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v3733 = *(*int32)(unsafe.Add(mBase, uint32(v427)+16))
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v3738 = F_FunctionCall2Coll(m, v401, v407, base.I64_extend_i32_u(v3728+v3729*v3677), base.I64_extend_i32_u(v3733+v3734*v3677))
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v424+v3677<<(uint(int32(2))%32)))) = base.F32_demote_f64(base.F64_reinterpret_i64(v3738))
	v3744 = v3677 + int32(1)
	if v336 != v3744 {
		v3677 = v3744
		goto L348
	} else {
		goto L353
	}
L353:
	;
	goto L349
L354:
	;
	v4234 = int32(0)
	if v336 <= v4234 {
		goto L384
	} else {
		goto L385
	}
L355:
	;
	v3804 = v3794
	goto L356
L356:
	;
	if v336 <= int32(0) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v4043 = int32(0)
	if v62 != int32(1) {
		goto L377
	} else {
		goto L378
	}
L358:
	;
	v4041 = v3804 + int32(1)
	if v4041 != v62 {
		v3804 = v4041
		goto L356
	} else {
		goto L376
	}
L359:
	;
	v3850 = v415 + v3804*v336<<(uint(int32(2))%32)
	v3851 = int32(0)
	if v1480 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v3857 = v3851
	v3859 = v3851
	goto L363
L361:
	;
	v3936 = v3851
	goto L362
L362:
	;
	v3981 = v3936 << (uint(int32(2)) % 32)
	v3982 = v3850 + v3981
	v3983 = float32(0)
	v3984 = *(*float32)(unsafe.Add(mBase, uint32(v3982)))
	v3986 = *(*float32)(unsafe.Add(mBase, uint32(v3981+v424)))
	v3987 = base.F32_sub(v3984, v3986)
	if base.F32_lt(v3987, v3983) != 0 {
		goto L373
	} else {
		goto L374
	}
L363:
	;
	v3902 = v3857 << (uint(int32(2)) % 32)
	v3903 = v3850 + v3902
	v3904 = float32(0)
	v3905 = *(*float32)(unsafe.Add(mBase, uint32(v3903)))
	v3907 = *(*float32)(unsafe.Add(mBase, uint32(v3902+v424)))
	v3908 = base.F32_sub(v3905, v3907)
	if base.F32_lt(v3908, v3904) != 0 {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	if v1478 == int32(0) {
		goto L358
	} else {
		goto L372
	}
L365:
	;
	v3911 = v3904
	goto L367
L366:
	;
	v3911 = v3908
	goto L367
L367:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3903))) = v3911
	v3914 = v3902 | int32(4)
	v3915 = v3850 + v3914
	v3916 = float32(0)
	v3917 = *(*float32)(unsafe.Add(mBase, uint32(v3915)))
	v3919 = *(*float32)(unsafe.Add(mBase, uint32(v3914+v424)))
	v3920 = base.F32_sub(v3917, v3919)
	if base.F32_lt(v3920, v3916) != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v3923 = v3916
	goto L370
L369:
	;
	v3923 = v3920
	goto L370
L370:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3915))) = v3923
	v3925 = int32(2)
	v3926 = v3857 + v3925
	v3928 = v3859 + v3925
	if v3928 != v1476 {
		v3857 = v3926
		v3859 = v3928
		goto L363
	} else {
		goto L371
	}
L371:
	;
	goto L364
L372:
	;
	v3936 = v3926
	goto L362
L373:
	;
	v3990 = v3983
	goto L375
L374:
	;
	v3990 = v3987
	goto L375
L375:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v3982))) = v3990
	goto L358
L376:
	;
	goto L357
L377:
	;
	v4051 = v4043
	v4054 = v4043
	goto L380
L378:
	;
	v4130 = v4043
	goto L379
L379:
	;
	v4174 = int32(2)
	v4175 = v4130 << (uint(v4174) % 32)
	v4176 = v417 + v4175
	v4177 = *(*float32)(unsafe.Add(mBase, uint32(v4176)))
	v4179 = *(*int32)(unsafe.Add(mBase, uint32(v4175+v412)))
	v4183 = *(*float32)(unsafe.Add(mBase, uint32(v424+v4179<<(uint(v4174)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v4176))) = base.F32_add(v4177, v4183)
	goto L354
L380:
	;
	v4095 = int32(2)
	v4096 = v4051 << (uint(v4095) % 32)
	v4097 = v417 + v4096
	v4098 = *(*float32)(unsafe.Add(mBase, uint32(v4097)))
	v4100 = *(*int32)(unsafe.Add(mBase, uint32(v4096+v412)))
	v4104 = *(*float32)(unsafe.Add(mBase, uint32(v424+v4100<<(uint(v4095)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v4097))) = base.F32_add(v4098, v4104)
	v4108 = v4096 | int32(4)
	v4109 = v417 + v4108
	v4110 = *(*float32)(unsafe.Add(mBase, uint32(v4109)))
	v4112 = *(*int32)(unsafe.Add(mBase, uint32(v4108+v412)))
	v4116 = *(*float32)(unsafe.Add(mBase, uint32(v424+v4112<<(uint(v4095)%32))))
	*(*float32)(unsafe.Add(mBase, uint32(v4109))) = base.F32_add(v4110, v4116)
	v4120 = v4051 + v4095
	v4122 = v4054 + v4095
	if v4122 != v62&v1471 {
		v4051 = v4120
		v4054 = v4122
		goto L380
	} else {
		goto L382
	}
L381:
	;
	if v62&v1473 == int32(0) {
		goto L354
	} else {
		goto L383
	}
L382:
	;
	goto L381
L383:
	;
	v4130 = v4120
	goto L379
L384:
	;
	if base.B2i32(v1494 != int32(0))&v2296 != 0 {
		goto L5
	} else {
		goto L393
	}
L385:
	;
	v4241 = v4234
	goto L386
L386:
	;
	v4285 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v4241 < v4285 {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	goto L3
L388:
	;
	v4287 = *(*int32)(unsafe.Add(mBase, uint32(v427)+16))
	v4288 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	F_VectorArraySet(m, l2, v4241, v4287+v4288*v4241)
	mBase = m.M
	v4292 = m.ExcPending
	if v4292 != 0 {
		goto L1
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	goto L387
L391:
	;
	v4294 = v4241 + int32(1)
	if v336 != v4294 {
		v4241 = v4294
		goto L386
	} else {
		goto L392
	}
L392:
	;
	goto L384
L393:
	;
	v4348 = v1494 + int32(1)
	if v4348 != int32(500) {
		v1494 = v4348
		goto L166
	} else {
		goto L394
	}
L394:
	;
	goto L167
L395:
	;
	goto L3
L396:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4814 = m.ExcPending
	if v4814 != 0 {
		goto L1
	} else {
		goto L444
	}
L397:
	;
	v4403 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4404 = F_palloc_mul(m, int32(4), v4403)
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
		goto L1
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4801 = m.ExcPending
	if v4801 != 0 {
		goto L1
	} else {
		goto L441
	}
L400:
	;
	v4406 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v4406 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v4409 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4412 = v4409
	v4418 = int32(0)
	goto L404
L402:
	;
	goto L403
L403:
	;
	v4658 = F_HnswOptionalProcInfo(m, l0, int32(2))
	mBase = m.M
	v4659 = m.ExcPending
	if v4659 != 0 {
		goto L1
	} else {
		goto L426
	}
L404:
	;
	if v4412 <= int32(0) {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	goto L403
L406:
	;
	v4468 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4468 <= v4418 {
		goto L3
	} else {
		goto L409
	}
L407:
	;
	v4462 = v4412 << (uint(int32(2)) % 32)
	if v4462 == int32(0) {
		goto L406
	} else {
		goto L408
	}
L408:
	;
	base.MemoryFill(m, v4404, int32(0), v4462)
	goto L406
L409:
	;
	v4470 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v4471 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v4474 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	m.T0[v4474].(func(*base.Module, int32, int32))(m, v4470+v4471*v4418, v4404)
	mBase = m.M
	v4476 = m.ExcPending
	if v4476 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	v4477 = int32(0)
	v4478 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v4477 < v4478 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v4485 = v4477
	goto L414
L412:
	;
	goto L413
L413:
	;
	v4606 = v4418 + int32(1)
	v4607 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4606 < v4607 {
		v4412 = v4478
		v4418 = v4606
		goto L404
	} else {
		goto L424
	}
L414:
	;
	v4532 = *(*float32)(unsafe.Add(mBase, uint32(v4404+v4485<<(uint(int32(2))%32))))
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v4532)&int32(2147483647)) {
		goto L396
	} else {
		goto L416
	}
L415:
	;
	goto L413
L416:
	;
	if base.F32_eq(base.F32_abs(v4532), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4544 = m.ExcPending
	if v4544 != 0 {
		goto L1
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	v4555 = v4485 + int32(1)
	if v4555 != v4478 {
		v4485 = v4555
		goto L414
	} else {
		goto L423
	}
L420:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_6), int32(0))
	mBase = m.M
	v4548 = m.ExcPending
	if v4548 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(509), int32(_a_F_IvfflatKmeans_8))
	mBase = m.M
	v4553 = m.ExcPending
	if v4553 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L423:
	;
	goto L415
L424:
	;
	goto L405
L425:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IvfflatKmeans[0])) = v58
	F_MemoryContextDelete(m, v55)
	mBase = m.M
	v4797 = m.ExcPending
	if v4797 != 0 {
		goto L1
	} else {
		goto L440
	}
L426:
	;
	if v4658 == int32(0) {
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v4662 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4662 <= int32(0) {
		goto L425
	} else {
		goto L428
	}
L428:
	;
	v4665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v4666 = *(*int32)(unsafe.Add(mBase, uint32(v4665)))
	v4672 = int32(0)
	goto L429
L429:
	;
	v4716 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v4716 <= v4672 {
		goto L395
	} else {
		goto L431
	}
L430:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		goto L1
	} else {
		goto L437
	}
L431:
	;
	v4718 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v4719 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v4723 = F_FunctionCall1Coll(m, v4658, v4666, base.I64_extend_i32_u(v4718+v4719*v4672))
	mBase = m.M
	v4724 = m.ExcPending
	if v4724 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	if v4723&int64(9223372036854775807) != int64(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v4730 = v4672 + int32(1)
	v4731 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v4731 <= v4730 {
		goto L425
	} else {
		goto L436
	}
L434:
	;
	goto L435
L435:
	;
	goto L430
L436:
	;
	v4672 = v4730
	goto L429
L437:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_9), int32(0))
	mBase = m.M
	v4740 = m.ExcPending
	if v4740 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(532), int32(_a_F_IvfflatKmeans_10))
	mBase = m.M
	v4745 = m.ExcPending
	if v4745 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L440:
	;
	return
L441:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_11), int32(0))
	mBase = m.M
	v4805 = m.ExcPending
	if v4805 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(543), int32(_a_F_IvfflatKmeans_12))
	mBase = m.M
	v4810 = m.ExcPending
	if v4810 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_13), int32(0))
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(506), int32(_a_F_IvfflatKmeans_8))
	mBase = m.M
	v4823 = m.ExcPending
	if v4823 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_14), int32(0))
	mBase = m.M
	v4831 = m.ExcPending
	if v4831 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_7), int32(294), int32(_a_F_IvfflatKmeans_15))
	mBase = m.M
	v4836 = m.ExcPending
	if v4836 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_IvfflatKmeans_16), int32(0))
	mBase = m.M
	v4892 = m.ExcPending
	if v4892 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(_a_F_IvfflatKmeans_17), int32(326), int32(_a_F_IvfflatKmeans_18))
	mBase = m.M
	v4897 = m.ExcPending
	if v4897 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
