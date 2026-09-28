package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RefreshMatViewByOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v93 int64
	_ = v93
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v218 int32
	_ = v218
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v250 int32
	_ = v250
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v316 int32
	_ = v316
	var v336 int32
	_ = v336
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v388 int32
	_ = v388
	var v402 int32
	_ = v402
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v435 int32
	_ = v435
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v486 int32
	_ = v486
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v535 int32
	_ = v535
	var v549 int32
	_ = v549
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v583 int32
	_ = v583
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v669 int32
	_ = v669
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v810 int32
	_ = v810
	var v823 int32
	_ = v823
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v877 int32
	_ = v877
	var v890 int32
	_ = v890
	var v904 int32
	_ = v904
	var v916 int32
	_ = v916
	var v927 int32
	_ = v927
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v981 int32
	_ = v981
	var v994 int32
	_ = v994
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1080 int32
	_ = v1080
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1107 int32
	_ = v1107
	var v1122 int32
	_ = v1122
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1151 int32
	_ = v1151
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1188 int32
	_ = v1188
	var v1199 int32
	_ = v1199
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1238 int32
	_ = v1238
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int64
	_ = v1253
	var v1264 int32
	_ = v1264
	var v1275 int32
	_ = v1275
	var v1286 int32
	_ = v1286
	var v1297 int32
	_ = v1297
	var v1301 int64
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1357 int64
	_ = v1357
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1496 int32
	_ = v1496
	var v1511 int32
	_ = v1511
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1538 int32
	_ = v1538
	var v1548 int32
	_ = v1548
	var v1554 int32
	_ = v1554
	var v1568 int32
	_ = v1568
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1603 int32
	_ = v1603
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1632 int32
	_ = v1632
	var v1642 int32
	_ = v1642
	var v1648 int32
	_ = v1648
	var v1662 int32
	_ = v1662
	var v1664 int64
	_ = v1664
	var v1679 int32
	_ = v1679
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1755 int32
	_ = v1755
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1803 int32
	_ = v1803
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1830 int32
	_ = v1830
	var v1840 int32
	_ = v1840
	var v1846 int32
	_ = v1846
	var v1860 int32
	_ = v1860
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1909 int32
	_ = v1909
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1936 int32
	_ = v1936
	var v1946 int32
	_ = v1946
	var v1952 int32
	_ = v1952
	var v1966 int32
	_ = v1966
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2051 int32
	_ = v2051
	var v2062 int32
	_ = v2062
	var v2076 int32
	_ = v2076
	var v2080 int32
	_ = v2080
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2137 int32
	_ = v2137
	var v2164 int32
	_ = v2164
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2184 int64
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2211 int32
	_ = v2211
	var v2221 int32
	_ = v2221
	var v2235 int32
	_ = v2235
	var v2242 int32
	_ = v2242
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2280 int32
	_ = v2280
	var v2295 int32
	_ = v2295
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2325 int32
	_ = v2325
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2352 int32
	_ = v2352
	var v2369 int32
	_ = v2369
	var v2383 int32
	_ = v2383
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2405 int32
	_ = v2405
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2479 int32
	_ = v2479
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2525 int32
	_ = v2525
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2548 int32
	_ = v2548
	var v2550 int32
	_ = v2550
	var v2554 int32
	_ = v2554
	var v2565 int32
	_ = v2565
	var v2576 int32
	_ = v2576
	var v2587 int32
	_ = v2587
	var v2603 int32
	_ = v2603
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2630 int32
	_ = v2630
	var v2640 int32
	_ = v2640
	var v2646 int32
	_ = v2646
	var v2660 int32
	_ = v2660
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2693 int32
	_ = v2693
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2720 int32
	_ = v2720
	var v2730 int32
	_ = v2730
	var v2736 int32
	_ = v2736
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2753 int32
	_ = v2753
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2790 int32
	_ = v2790
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2817 int32
	_ = v2817
	var v2827 int32
	_ = v2827
	var v2833 int32
	_ = v2833
	var v2847 int32
	_ = v2847
	var v2858 int32
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2881 int32
	_ = v2881
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2908 int32
	_ = v2908
	var v2918 int32
	_ = v2918
	var v2924 int32
	_ = v2924
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2941 int32
	_ = v2941
	var v2956 int32
	_ = v2956
	var v2968 int32
	_ = v2968
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2981 int32
	_ = v2981
	var v3002 int32
	_ = v3002
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3029 int32
	_ = v3029
	var v3039 int32
	_ = v3039
	var v3045 int32
	_ = v3045
	var v3059 int32
	_ = v3059
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3085 int32
	_ = v3085
	var v3098 int32
	_ = v3098
	var v3112 int32
	_ = v3112
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3156 int64
	_ = v3156
	var v3169 int32
	_ = v3169
	var v3181 int32
	_ = v3181
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3205 int32
	_ = v3205
	var v3211 int32
	_ = v3211
	var v3223 int32
	_ = v3223
	var v3277 int32
	_ = v3277
	var v3289 int32
	_ = v3289
	var v3290 int32
	_ = v3290
	var v3305 int32
	_ = v3305
	var v3319 int32
	_ = v3319
	var v3361 int32
	_ = v3361
	var v3362 int64
	_ = v3362
	var v3366 int32
	_ = v3366
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3372 int32
	_ = v3372
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3382 int64
	_ = v3382
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3387 int32
	_ = v3387
	v8 = int32(0)
	v42 = m.G0
	v44 = v42 - int32(736)
	m.G0 = v44
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = int32(_a_F_RefreshMatViewByOid_0)
	goto L3
L2:
	;
	v48 = int32(_a_F_RefreshMatViewByOid_1)
	goto L3
L3:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v51 = int32(_a_F_RefreshMatViewByOid_2)
	goto L6
L5:
	;
	v51 = int32(_a_F_RefreshMatViewByOid_1)
	goto L6
L6:
	;
	v61 = v8
	v62 = v8
	v63 = v8
	v64 = v8
	v65 = v8
	v66 = v8
	v67 = v8
	v68 = v8
	v69 = v8
	v70 = int32(-1)
	v93 = int64(0)
	goto L7
L7:
	;
	goto L10
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	goto L8
L10:
	;
	if v70 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v3361 = int32(m.ExcTag)
	v3362 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v3361 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L12
	} else {
		goto L340
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_list_free(m, v3211)
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L12
	} else {
		goto L339
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v3129
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v3128
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v3130
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v3156
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v3126
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v3131
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v3132
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v3127
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v3125
	F_relation_close(m, v3125, int32(0))
	mBase = m.M
	v3169 = m.ExcPending
	if v3169 != 0 {
		goto L12
	} else {
		goto L330
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_list_free(m, v2025)
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L12
	} else {
		goto L275
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v2536 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[0]))
	v2537 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L12
	} else {
		goto L270
	}
L18:
	;
	if v1334 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L19:
	;
	v1326 = v61
	v1327 = v62
	v1328 = v63
	v1329 = v64
	v1330 = v65
	v1331 = v66
	v1332 = v67
	v1333 = v68
	v1334 = v69
	v1357 = v93
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v61
	v106 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v44+int32(672)))) = v124
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v44+int32(668)))) = v127
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v44)+668))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v138 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v109
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v155 = int32(_a_F_RefreshMatViewByOid_3)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[3]))
	v159 = v157 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[3])) = v159
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_RestrictSearchPath(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+119)))
	if v173 != int32(109) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l4 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errcode(m, int32(1088))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+496)) = v201 + int32(4)
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_4), v44+int32(496))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(201), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	goto L9
L34:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+124)))
	if v351 != 0 {
		goto L49
	} else {
		goto L50
	}
L35:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+129)))
	if v235 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L12
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if l3 == int32(0) {
		goto L34
	} else {
		goto L43
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errcode(m, int32(1088))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_7), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(207), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	goto L9
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errcode(m, int32(16801924))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+392)) = int32(_a_F_RefreshMatViewByOid_8)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+388)) = int32(_a_F_RefreshMatViewByOid_9)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+384)) = int32(_a_F_RefreshMatViewByOid_10)
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_11), v44+int32(384))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(214), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	goto L9
L48:
	;
	if v353 != int32(1) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v106)+68))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	if int32(0) < v353 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L12
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+400)) = v371 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_12), v44+int32(400))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(224), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	goto L9
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L12
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	if v452 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+480)) = v418 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_13), v44+int32(480))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(229), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	goto L9
L62:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	if v501 != 0 {
		goto L71
	} else {
		goto L72
	}
L63:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+17)))
	if v455 != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L12
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+464)) = v469 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_14), v44+int32(464))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(235), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	goto L9
L70:
	;
	if l4 != 0 {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)+4))
	if v502 == int32(1) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+416)) = v518 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_15), v44+int32(416))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(241), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	goto L9
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v559 = F_RelationGetIndexList(m, v106)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v501)+12))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v969)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_CheckTableNotInUse(m, v106, v48)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L12
	} else {
		goto L112
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_relation_close(m, v623, int32(1))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L12
	} else {
		goto L110
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_list_free(m, v559)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L12
	} else {
		goto L102
	}
L83:
	;
	if v559 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v563 = int32(0)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	if v564 <= v563 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v583 = v563
	goto L86
L86:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v608+v583<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v623 = F_index_open(m, v612, int32(1))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L12
	} else {
		goto L89
	}
L87:
	;
	goto L82
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_relation_close(m, v623, int32(1))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L12
	} else {
		goto L100
	}
L89:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v623)+192))
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+12)))
	if v626 != int32(1) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+16)))
	if v629 != int32(1) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+18)))
	if v632 != int32(1) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v644 = F_RelationGetIndexPredicate(m, v623)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	if v644 != 0 {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	v646 = int32(*(*int16)(unsafe.Add(mBase, uint32(v625)+8)))
	if v646 <= int32(0) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v669 = int32(0)
	goto L96
L96:
	;
	v696 = int32(*(*int16)(unsafe.Add(mBase, uint32(v625+int32(48)+v669<<(uint(int32(1))%32)))))
	if v696 <= int32(0) {
		goto L88
	} else {
		goto L98
	}
L97:
	;
	goto L81
L98:
	;
	v700 = v669 + int32(1)
	if v646 != v700 {
		v669 = v700
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v756 = v583 + int32(1)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	if v756 < v757 {
		v583 = v756
		goto L86
	} else {
		goto L101
	}
L101:
	;
	goto L87
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errcode(m, int32(325))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v836)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v847 = F_get_namespace_name(m, v837)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v861 = F_quote_qualified_identifier(m, v847, v849+int32(4))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+432)) = v861
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_16), v44+int32(432))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errhint(m, int32(_a_F_RefreshMatViewByOid_17), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(275), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	goto L9
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_list_free(m, v559)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L12
	} else {
		goto L111
	}
L111:
	;
	goto L80
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_SetMatViewPopulatedState(m, v106, l3^int32(1))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	if l4 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1027 = base.I32_extend8_s(v1014)
	v1029 = F_make_new_heap(m, l1, v1016, v1017, v1027, int32(7))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L12
	} else {
		goto L119
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1006 = F_GetDefaultTablespace(m, int32(116), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L12
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1010)+118)))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+92))
	v1013 = v67
	v1014 = v1011
	v1015 = v1010
	v1016 = v1012
	goto L114
L118:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v1013 = v1006
	v1014 = int32(116)
	v1015 = v1008
	v1016 = v1006
	goto L114
L119:
	;
	if l3 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1044 = F_palloc0(m, int32(40))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L12
	} else {
		goto L123
	}
L121:
	;
	v1301 = int64(0)
	goto L122
L122:
	;
	if l4 == int32(0) {
		goto L17
	} else {
		goto L151
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+20)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+16)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+12)) = int32(604)
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+8)) = int32(605)
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+4)) = int32(606)
	*(*int32)(unsafe.Add(mBase, uint32(v1044))) = int32(607)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1066 = F_copyObjectImpl(m, v970)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_AcquireRewriteLocks(m, v1066, int32(1), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1090 = F_QueryRewrite(m, v1066)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L12
	} else {
		goto L127
	}
L126:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+12))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1137)))
	v1140 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[4]))
	if v1140 != 0 {
		goto L135
	} else {
		goto L136
	}
L127:
	;
	if v1090 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	if v1092 == int32(1) {
		goto L126
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L12
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+448)) = v51
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_18), v44+int32(448))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L12
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(421), int32(_a_F_RefreshMatViewByOid_19))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L12
	} else {
		goto L134
	}
L134:
	;
	goto L9
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_ProcessInterrupts(m)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L12
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1162 = int32(0)
	v1164 = F_pg_plan_query(m, v1138, l5, int32(2048), v1162, v1162)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L12
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1176 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[5]))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1176)))
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_PushCopiedSnapshot(m, v1177)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L12
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1210 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[5]))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1210)))
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v1221 = int32(0)
	v1225 = F_CreateQueryDesc(m, v1164, l5, v1211, v1221, v1044, v1221, v1221, v1221)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L12
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_ExecutorStart(m, v1225, int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_ExecutorRun(m, v1225, int32(1), int64(0))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+44))
	v1253 = *(*int64)(unsafe.Add(mBase, uint32(v1252)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_ExecutorFinish(m, v1225)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_ExecutorEnd(m, v1225)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_FreeQueryDesc(m, v1225)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L12
	} else {
		goto L150
	}
L150:
	;
	v1301 = v1253
	goto L122
L151:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[6]))
	v1308 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7]))
	v1310 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8]))
	goto L152
L152:
	;
	v1312 = v44 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v1312)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1312))) = v44 + int32(508)
	goto L155
L153:
	;
	v1326 = v106
	v1327 = v1029
	v1328 = v109
	v1329 = v1306
	v1330 = v1308
	v1331 = v1310
	v1332 = v1013
	v1333 = v159
	v1334 = int32(0)
	v1357 = v1301
	goto L18
L155:
	;
	goto L153
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7])) = v44 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v44)+668))
	v1376 = v44 + int32(676)
	F_initStringInfo(m, v1376)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L12
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[6])) = v1329
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7])) = v1330
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8])) = v1331
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_pg_re_throw(m)
	mBase = m.M
	v2525 = m.ExcPending
	if v2525 != 0 {
		goto L12
	} else {
		goto L269
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1389 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+48))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1391)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1402 = F_get_namespace_name(m, v1392)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1416 = F_quote_qualified_identifier(m, v1402, v1404+int32(4))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L12
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1428 = F_table_open(m, v1327, int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L12
	} else {
		goto L163
	}
L163:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+48))
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1441 = F_get_namespace_name(m, v1431)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L12
	} else {
		goto L164
	}
L164:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1454 = v1443 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+368)) = v1454
	v1459 = F_psprintf(m, int32(_a_F_RefreshMatViewByOid_20), v44+int32(368))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L12
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1470 = F_quote_qualified_identifier(m, v1441, v1454)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1481 = F_quote_qualified_identifier(m, v1441, v1459)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+48))
	v1484 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1483)+120)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L12
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+352)) = v1470
	F_appendStringInfo(m, v1376, int32(_a_F_RefreshMatViewByOid_21), v44+int32(352))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L12
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1522 = F_SPI_exec(m, v1521)
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L12
	} else {
		goto L170
	}
L170:
	;
	if v1522 != int32(4) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L12
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1579 = v44 + int32(676)
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1579)))
	v1581 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1580))) = uint8(v1581)
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+12)) = v1581
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+4)) = v1581
	goto L177
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+336)) = v1548
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(336))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(635), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L12
	} else {
		goto L176
	}
L176:
	;
	goto L9
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+328)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v44)+324)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v44)+320)) = v1470
	F_appendStringInfo(m, v1579, int32(_a_F_RefreshMatViewByOid_24), v44+int32(320))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L12
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1616 = F_SPI_execute(m, v1613, int32(0), int32(1))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L12
	} else {
		goto L179
	}
L179:
	;
	if v1616 != int32(5) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L12
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v1664 = *(*int64)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[9]))
	if v1664 != int64(0) {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+304)) = v1642
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(304))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L12
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(658), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L12
	} else {
		goto L185
	}
L185:
	;
	goto L9
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L12
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v1374 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v1328
	goto L195
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errcode(m, int32(66))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L12
	} else {
		goto L190
	}
L190:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+288)) = v1692 + int32(4)
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_25), v44+int32(288))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L12
	} else {
		goto L191
	}
L191:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[10]))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+4))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1713)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1725 = F_SPI_getvalue(m, v1714, v1712, int32(1))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L12
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v1725
	v1740 = F_errdetail(m, int32(_a_F_RefreshMatViewByOid_26), v44+int32(272))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L12
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(673), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L12
	} else {
		goto L194
	}
L194:
	;
	goto L9
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1781 = v44 + int32(676)
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1781)))
	v1783 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1782))) = uint8(v1783)
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+12)) = v1783
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+4)) = v1783
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+256)) = v1481
	F_appendStringInfo(m, v1781, int32(_a_F_RefreshMatViewByOid_27), v44+int32(256))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L12
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1814 = F_SPI_exec(m, v1813)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L12
	} else {
		goto L198
	}
L198:
	;
	if v1814 != int32(4) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L12
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v1374 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v1328
	goto L205
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+240)) = v1840
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(240))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L12
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(691), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L12
	} else {
		goto L204
	}
L204:
	;
	goto L9
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1886 = v44 + int32(676)
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1886)))
	v1888 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1887))) = uint8(v1888)
	*(*int32)(unsafe.Add(mBase, uint32(v1886)+12)) = v1888
	*(*int32)(unsafe.Add(mBase, uint32(v1886)+4)) = v1888
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+228)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v44)+224)) = v1481
	F_appendStringInfo(m, v1886, int32(_a_F_RefreshMatViewByOid_28), v44+int32(224))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L12
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1920 = F_SPI_exec(m, v1919)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L12
	} else {
		goto L208
	}
L208:
	;
	if v1920 != int32(4) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L12
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1977 = v44 + int32(676)
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1977)))
	v1979 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1978))) = uint8(v1979)
	*(*int32)(unsafe.Add(mBase, uint32(v1977)+12)) = v1979
	*(*int32)(unsafe.Add(mBase, uint32(v1977)+4)) = v1979
	goto L215
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+208)) = v1946
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(208))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L12
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(699), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L12
	} else {
		goto L214
	}
L214:
	;
	goto L9
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+204)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v44)+200)) = v1416
	*(*int32)(unsafe.Add(mBase, uint32(v44)+196)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v44)+192)) = v1481
	F_appendStringInfo(m, v1977, int32(_a_F_RefreshMatViewByOid_29), v44+int32(192))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L12
	} else {
		goto L216
	}
L216:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2014 = F_palloc0_mul(m, int32(4), v1484)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L12
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2025 = F_RelationGetIndexList(m, v1389)
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L12
	} else {
		goto L218
	}
L218:
	;
	if v2025 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v3211 = int32(0)
	goto L14
L220:
	;
	goto L221
L221:
	;
	v2030 = int32(0)
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2025)+4))
	if v2032 <= v2030 {
		v3211 = v2025
		goto L14
	} else {
		goto L222
	}
L222:
	;
	v2051 = v2030
	v2062 = v2030
	goto L223
L223:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2025)+12))
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2076+v2051<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2091 = F_index_open(m, v2080, int32(3))
	mBase = m.M
	v2092 = m.ExcPending
	if v2092 != 0 {
		goto L12
	} else {
		goto L226
	}
L224:
	;
	goto L16
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_relation_close(m, v2091, int32(0))
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L12
	} else {
		goto L267
	}
L226:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+192))
	v2094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2093)+12)))
	if v2094 != int32(1) {
		v2479 = v2062
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v2097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2093)+16)))
	if v2097 != int32(1) {
		v2479 = v2062
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2093)+18)))
	if v2100 != int32(1) {
		v2479 = v2062
		goto L225
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2112 = F_RelationGetIndexPredicate(m, v2091)
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L12
	} else {
		goto L230
	}
L230:
	;
	if v2112 != 0 {
		v2479 = v2062
		goto L225
	} else {
		goto L231
	}
L231:
	;
	v2114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2093)+8)))
	if v2114 <= int32(0) {
		v2479 = v2062
		goto L225
	} else {
		goto L232
	}
L232:
	;
	v2137 = int32(0)
	goto L233
L233:
	;
	v2164 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2093+int32(48)+v2137<<(uint(int32(1))%32)))))
	if v2164 <= int32(0) {
		v2479 = v2062
		goto L225
	} else {
		goto L235
	}
L234:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+192))
	v2171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2170)+10)))
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2184 = F_SysCacheGetAttrNotNull(m, int32(34), v2172, int32(18))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L12
	} else {
		goto L237
	}
L235:
	;
	v2168 = v2137 + int32(1)
	if v2168 != v2114 {
		v2137 = v2168
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	if v2171 <= int32(0) {
		v2479 = v2062
		goto L225
	} else {
		goto L238
	}
L238:
	;
	v2211 = int32(0)
	v2221 = v2062
	goto L239
L239:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v2003)))
	v2242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2170+int32(48)+v2211<<(uint(int32(1))%32)))))
	v2247 = v2003 + v2235<<(uint(int32(3))%32) + v2242*int32(100) - int32(72)
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v2247)+68))
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v2184)+int32(24)+v2211<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2264 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(v2252))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L12
	} else {
		goto L241
	}
L240:
	;
	v2479 = v2448
	goto L225
L241:
	;
	if v2264 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L12
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+16))
	v2311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2310)+22)))
	v2312 = v2310 + v2311
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+84))
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_ReleaseCatCache(m, v2264)
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L12
	} else {
		goto L248
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v2252
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_30), v44+int32(16))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L12
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(762), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L12
	} else {
		goto L247
	}
L247:
	;
	goto L9
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2336 = F_get_opfamily_member_for_cmptype(m, v2314, v2313, v2313, int32(3))
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L12
	} else {
		goto L249
	}
L249:
	;
	if v2336 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L12
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v2388 = v2014 + (v2242-int32(1))<<(uint(int32(2))%32)
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v2388)))
	if v2336 != v2389 {
		goto L256
	} else {
		goto L257
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+40)) = v2314
	*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v2313
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v2313
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_31), v44+int32(32))
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L12
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(771), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L12
	} else {
		goto L255
	}
L255:
	;
	goto L9
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2388))) = v2336
	if v2221 != 0 {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	v2448 = v2221
	goto L258
L258:
	;
	v2450 = v2211 + int32(1)
	if v2450 != v2171 {
		v2211 = v2450
		v2221 = v2448
		goto L239
	} else {
		goto L266
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_appendStringInfoString(m, v44+int32(676), int32(_a_F_RefreshMatViewByOid_32))
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L12
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2417 = v2247 + int32(4)
	v2418 = F_quote_qualified_identifier(m, int32(_a_F_RefreshMatViewByOid_33), v2417)
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L12
	} else {
		goto L263
	}
L262:
	;
	goto L261
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2430 = F_quote_qualified_identifier(m, int32(_a_F_RefreshMatViewByOid_34), v2417)
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L12
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_generate_operator_clause(m, v44+int32(676), v2418, v2248, v2336, v2430, v2248)
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L12
	} else {
		goto L265
	}
L265:
	;
	v2448 = int32(1)
	goto L258
L266:
	;
	goto L240
L267:
	;
	v2506 = v2051 + int32(1)
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2025)+4))
	if v2506 < v2507 {
		v2051 = v2506
		v2062 = v2479
		goto L223
	} else {
		goto L268
	}
L268:
	;
	goto L224
L269:
	;
	goto L9
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	v2548 = int32(0)
	v2550 = int32(1)
	F_finish_heap_swap(m, l1, v1029, v2548, v2548, v2550, v2550, v2550, v2536, v2537, v1027)
	mBase = m.M
	v2554 = m.ExcPending
	if v2554 != 0 {
		goto L12
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_pgstat_count_truncate(m, v106)
	mBase = m.M
	v2565 = m.ExcPending
	if v2565 != 0 {
		goto L12
	} else {
		goto L272
	}
L272:
	;
	if l3 != 0 {
		v3125 = v106
		v3126 = v1029
		v3127 = v109
		v3128 = v64
		v3129 = v65
		v3130 = v66
		v3131 = v1013
		v3132 = v159
		v3156 = v1301
		goto L15
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v106
	F_pgstat_count_heap_insert(m, v106, v1301)
	mBase = m.M
	v2576 = m.ExcPending
	if v2576 != 0 {
		goto L12
	} else {
		goto L274
	}
L274:
	;
	v3125 = v106
	v3126 = v1029
	v3127 = v109
	v3128 = v64
	v3129 = v65
	v3130 = v66
	v3131 = v1013
	v3132 = v159
	v3156 = v1301
	goto L15
L275:
	;
	if v2479 == int32(0) {
		goto L13
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_appendStringInfoString(m, v44+int32(676), int32(_a_F_RefreshMatViewByOid_35))
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L12
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2614 = F_SPI_exec(m, v2613)
	mBase = m.M
	v2615 = m.ExcPending
	if v2615 != 0 {
		goto L12
	} else {
		goto L278
	}
L278:
	;
	if v2614 != int32(7) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L12
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2671 = v44 + int32(676)
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2671)))
	v2673 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2672))) = uint8(v2673)
	*(*int32)(unsafe.Add(mBase, uint32(v2671)+12)) = v2673
	*(*int32)(unsafe.Add(mBase, uint32(v2671)+4)) = v2673
	goto L285
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2640 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+176)) = v2640
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(176))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L12
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(836), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L12
	} else {
		goto L284
	}
L284:
	;
	goto L9
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+160)) = v1481
	F_appendStringInfo(m, v2671, int32(_a_F_RefreshMatViewByOid_21), v44+int32(160))
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L12
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2704 = F_SPI_exec(m, v2703)
	mBase = m.M
	v2705 = m.ExcPending
	if v2705 != 0 {
		goto L12
	} else {
		goto L287
	}
L287:
	;
	if v2704 != int32(4) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2720 = m.ExcPending
	if v2720 != 0 {
		goto L12
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v2751 = int32(_a_F_RefreshMatViewByOid_36)
	v2753 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8])) = v2753 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2767 = v44 + int32(676)
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v2767)))
	v2769 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2768))) = uint8(v2769)
	*(*int32)(unsafe.Add(mBase, uint32(v2767)+12)) = v2769
	*(*int32)(unsafe.Add(mBase, uint32(v2767)+4)) = v2769
	goto L294
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+144)) = v2730
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(144))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L12
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(847), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v2750 = m.ExcPending
	if v2750 != 0 {
		goto L12
	} else {
		goto L293
	}
L293:
	;
	goto L9
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+132)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+128)) = v1416
	F_appendStringInfo(m, v2767, int32(_a_F_RefreshMatViewByOid_37), v44+int32(128))
	mBase = m.M
	v2790 = m.ExcPending
	if v2790 != 0 {
		goto L12
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2801 = F_SPI_exec(m, v2800)
	mBase = m.M
	v2802 = m.ExcPending
	if v2802 != 0 {
		goto L12
	} else {
		goto L296
	}
L296:
	;
	if v2801 != int32(8) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L12
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2858 = v44 + int32(676)
	v2859 = *(*int32)(unsafe.Add(mBase, uint32(v2858)))
	v2860 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2859))) = uint8(v2860)
	*(*int32)(unsafe.Add(mBase, uint32(v2858)+12)) = v2860
	*(*int32)(unsafe.Add(mBase, uint32(v2858)+4)) = v2860
	goto L303
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+112)) = v2827
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(112))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L12
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(860), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v2847 = m.ExcPending
	if v2847 != 0 {
		goto L12
	} else {
		goto L302
	}
L302:
	;
	goto L9
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+100)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v1416
	F_appendStringInfo(m, v2858, int32(_a_F_RefreshMatViewByOid_38), v44+int32(96))
	mBase = m.M
	v2881 = m.ExcPending
	if v2881 != 0 {
		goto L12
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2892 = F_SPI_exec(m, v2891)
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L12
	} else {
		goto L305
	}
L305:
	;
	if v2892 != int32(7) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L12
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v2939 = int32(_a_F_RefreshMatViewByOid_36)
	v2941 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8])) = v2941 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_relation_close(m, v1428, int32(0))
	mBase = m.M
	v2956 = m.ExcPending
	if v2956 != 0 {
		goto L12
	} else {
		goto L312
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v2918
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(80))
	mBase = m.M
	v2924 = m.ExcPending
	if v2924 != 0 {
		goto L12
	} else {
		goto L310
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(869), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L12
	} else {
		goto L311
	}
L311:
	;
	goto L9
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_relation_close(m, v1389, int32(0))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L12
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v2979 = v44 + int32(676)
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(v2979)))
	v2981 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2980))) = uint8(v2981)
	*(*int32)(unsafe.Add(mBase, uint32(v2979)+12)) = v2981
	*(*int32)(unsafe.Add(mBase, uint32(v2979)+4)) = v2981
	goto L314
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+68)) = v1470
	*(*int32)(unsafe.Add(mBase, uint32(v44)+64)) = v1481
	F_appendStringInfo(m, v2979, int32(_a_F_RefreshMatViewByOid_39), v44-int32(-64))
	mBase = m.M
	v3002 = m.ExcPending
	if v3002 != 0 {
		goto L12
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v3012 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v3013 = F_SPI_exec(m, v3012)
	mBase = m.M
	v3014 = m.ExcPending
	if v3014 != 0 {
		goto L12
	} else {
		goto L316
	}
L316:
	;
	if v3013 != int32(4) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L12
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v3069 = F_SPI_finish(m)
	mBase = m.M
	v3070 = m.ExcPending
	if v3070 != 0 {
		goto L12
	} else {
		goto L323
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	v3039 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v3039
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_22), v44+int32(48))
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L12
	} else {
		goto L321
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(880), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L12
	} else {
		goto L322
	}
L322:
	;
	goto L9
L323:
	;
	if v3069 != int32(2) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		goto L12
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7])) = v1330
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[6])) = v1329
	v3125 = v1326
	v3126 = v1327
	v3127 = v1328
	v3128 = v1329
	v3129 = v1330
	v3130 = v1331
	v3131 = v1332
	v3132 = v1333
	v3156 = v1357
	goto L15
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_40), int32(0))
	mBase = m.M
	v3098 = m.ExcPending
	if v3098 != 0 {
		goto L12
	} else {
		goto L328
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(884), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v3112 = m.ExcPending
	if v3112 != 0 {
		goto L12
	} else {
		goto L329
	}
L329:
	;
	goto L9
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v3129
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v3128
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v3130
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v3156
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v3126
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v3131
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v3132
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v3127
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v3125
	F_AtEOXact_GUC(m, int32(0), v3132)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L12
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v3128
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v3129
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v3130
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v3156
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v3126
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v3131
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v3132
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v3127
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v3125
	v3191 = *(*int32)(unsafe.Add(mBase, uint32(v44)+672))
	v3192 = *(*int32)(unsafe.Add(mBase, uint32(v44)+668))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v3192
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v3191
	goto L332
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	if l6 != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6)+8)) = v3156
	if l2 != 0 {
		goto L336
	} else {
		goto L337
	}
L334:
	;
	goto L335
L335:
	;
	m.G0 = v44 + int32(736)
	return
L336:
	;
	v3205 = int32(180)
	goto L338
L337:
	;
	v3205 = int32(169)
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v3205
	goto L335
L339:
	;
	goto L13
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errcode(m, int32(1088))
	mBase = m.M
	v3289 = m.ExcPending
	if v3289 != 0 {
		goto L12
	} else {
		goto L341
	}
L341:
	;
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v3290 + int32(4)
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_41), v44)
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L12
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+696)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+692)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+700)) = v1331
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+720)) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v44)+724)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v44)+728)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1326
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(827), int32(_a_F_RefreshMatViewByOid_23))
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L12
	} else {
		goto L343
	}
L343:
	;
	goto L9
L344:
	;
	v3366 = int32(v3362)
	m.G0 = v44
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(v3366)+4))
	v3369 = *(*int32)(unsafe.Add(mBase, uint32(v3366)))
	v3372 = *(*int32)(unsafe.Add(mBase, uint32(v3369)))
	if v44+int32(508) == v3372 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	m.ExcPending = 1
	goto L353
L346:
	;
	if v3376 != 0 {
		goto L350
	} else {
		goto L351
	}
L347:
	;
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(v3369)+4))
	v3376 = v3374
	goto L349
L348:
	;
	v3376 = int32(0)
	goto L349
L349:
	;
	goto L346
L350:
	;
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(v44)+732))
	v3378 = *(*int32)(unsafe.Add(mBase, uint32(v44)+728))
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v44)+724))
	v3380 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v3381 = *(*int32)(unsafe.Add(mBase, uint32(v44)+716))
	v3382 = *(*int64)(unsafe.Add(mBase, uint32(v44)+704))
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v44)+700))
	v3384 = *(*int32)(unsafe.Add(mBase, uint32(v44)+696))
	v3385 = *(*int32)(unsafe.Add(mBase, uint32(v44)+692))
	v61 = v3377
	v62 = v3381
	v63 = v3378
	v64 = v3385
	v65 = v3384
	v66 = v3383
	v67 = v3380
	v68 = v3379
	v69 = v3368
	v70 = v3376
	v93 = v3382
	goto L7
L351:
	;
	goto L352
L352:
	;
	F___wasm_longjmp(m, v3369, v3368)
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	return
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
