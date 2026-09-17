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
	var v334 int32
	_ = v334
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v386 int32
	_ = v386
	var v400 int32
	_ = v400
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v433 int32
	_ = v433
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v484 int32
	_ = v484
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v533 int32
	_ = v533
	var v547 int32
	_ = v547
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v581 int32
	_ = v581
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v667 int32
	_ = v667
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v808 int32
	_ = v808
	var v821 int32
	_ = v821
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v875 int32
	_ = v875
	var v888 int32
	_ = v888
	var v902 int32
	_ = v902
	var v914 int32
	_ = v914
	var v925 int32
	_ = v925
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v979 int32
	_ = v979
	var v992 int32
	_ = v992
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
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
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1078 int32
	_ = v1078
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1105 int32
	_ = v1105
	var v1120 int32
	_ = v1120
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1149 int32
	_ = v1149
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1185 int32
	_ = v1185
	var v1196 int32
	_ = v1196
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1235 int32
	_ = v1235
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int64
	_ = v1250
	var v1261 int32
	_ = v1261
	var v1272 int32
	_ = v1272
	var v1283 int32
	_ = v1283
	var v1294 int32
	_ = v1294
	var v1298 int64
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
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
	var v1354 int64
	_ = v1354
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1477 int32
	_ = v1477
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1508 int32
	_ = v1508
	var v1523 int32
	_ = v1523
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1550 int32
	_ = v1550
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1580 int32
	_ = v1580
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1615 int32
	_ = v1615
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1644 int32
	_ = v1644
	var v1654 int32
	_ = v1654
	var v1660 int32
	_ = v1660
	var v1674 int32
	_ = v1674
	var v1676 int64
	_ = v1676
	var v1691 int32
	_ = v1691
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1753 int32
	_ = v1753
	var v1767 int32
	_ = v1767
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1815 int32
	_ = v1815
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1842 int32
	_ = v1842
	var v1852 int32
	_ = v1852
	var v1858 int32
	_ = v1858
	var v1872 int32
	_ = v1872
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1921 int32
	_ = v1921
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1948 int32
	_ = v1948
	var v1958 int32
	_ = v1958
	var v1964 int32
	_ = v1964
	var v1978 int32
	_ = v1978
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2064 int32
	_ = v2064
	var v2075 int32
	_ = v2075
	var v2089 int32
	_ = v2089
	var v2093 int32
	_ = v2093
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2150 int32
	_ = v2150
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2223 int32
	_ = v2223
	var v2233 int32
	_ = v2233
	var v2247 int32
	_ = v2247
	var v2254 int32
	_ = v2254
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2291 int32
	_ = v2291
	var v2306 int32
	_ = v2306
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2336 int32
	_ = v2336
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2363 int32
	_ = v2363
	var v2380 int32
	_ = v2380
	var v2394 int32
	_ = v2394
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2416 int32
	_ = v2416
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2490 int32
	_ = v2490
	var v2515 int32
	_ = v2515
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2536 int32
	_ = v2536
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2564 int32
	_ = v2564
	var v2575 int32
	_ = v2575
	var v2586 int32
	_ = v2586
	var v2597 int32
	_ = v2597
	var v2613 int32
	_ = v2613
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2640 int32
	_ = v2640
	var v2650 int32
	_ = v2650
	var v2656 int32
	_ = v2656
	var v2670 int32
	_ = v2670
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2703 int32
	_ = v2703
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2730 int32
	_ = v2730
	var v2740 int32
	_ = v2740
	var v2746 int32
	_ = v2746
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2800 int32
	_ = v2800
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2812 int32
	_ = v2812
	var v2827 int32
	_ = v2827
	var v2837 int32
	_ = v2837
	var v2843 int32
	_ = v2843
	var v2857 int32
	_ = v2857
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2891 int32
	_ = v2891
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2918 int32
	_ = v2918
	var v2928 int32
	_ = v2928
	var v2934 int32
	_ = v2934
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2951 int32
	_ = v2951
	var v2966 int32
	_ = v2966
	var v2978 int32
	_ = v2978
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v3012 int32
	_ = v3012
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3039 int32
	_ = v3039
	var v3049 int32
	_ = v3049
	var v3055 int32
	_ = v3055
	var v3069 int32
	_ = v3069
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3095 int32
	_ = v3095
	var v3108 int32
	_ = v3108
	var v3122 int32
	_ = v3122
	var v3135 int32
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3166 int64
	_ = v3166
	var v3179 int32
	_ = v3179
	var v3191 int32
	_ = v3191
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3215 int32
	_ = v3215
	var v3221 int32
	_ = v3221
	var v3233 int32
	_ = v3233
	var v3287 int32
	_ = v3287
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3315 int32
	_ = v3315
	var v3329 int32
	_ = v3329
	var v3371 int32
	_ = v3371
	var v3372 int64
	_ = v3372
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3382 int32
	_ = v3382
	var v3384 int32
	_ = v3384
	var v3386 int32
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3392 int64
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3395 int32
	_ = v3395
	var v3397 int32
	_ = v3397
	v8 = int32(0)
	v42 = m.G0
	v44 = v42 - int32(752)
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
	v3371 = int32(m.ExcTag)
	v3372 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v3371 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		goto L12
	} else {
		goto L341
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_list_free(m, v3221)
	mBase = m.M
	v3233 = m.ExcPending
	if v3233 != 0 {
		goto L12
	} else {
		goto L340
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v3139
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v3138
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v3140
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v3166
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v3136
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v3141
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v3142
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v3137
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v3135
	F_relation_close(m, v3135, int32(0))
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L12
	} else {
		goto L331
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_list_free(m, v2038)
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L12
	} else {
		goto L276
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v2547 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[0]))
	v2548 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L12
	} else {
		goto L271
	}
L18:
	;
	if v1331 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L19:
	;
	v1323 = v61
	v1324 = v62
	v1325 = v63
	v1326 = v64
	v1327 = v65
	v1328 = v66
	v1329 = v67
	v1330 = v68
	v1331 = v69
	v1354 = v93
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v61
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v44+int32(672)))) = v124
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v44+int32(668)))) = v127
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v44)+668))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v138 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v109
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v155 = int32(_a_F_RefreshMatViewByOid_3)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[3]))
	v159 = v157 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[3])) = v159
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(202), int32(_a_F_RefreshMatViewByOid_6))
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
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+124)))
	if v349 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(208), int32(_a_F_RefreshMatViewByOid_6))
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+388)) = int32(_a_F_RefreshMatViewByOid_8)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+384)) = int32(_a_F_RefreshMatViewByOid_9)
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_10), v44+int32(384))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(215), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	goto L9
L48:
	;
	if v351 != int32(1) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v106)+68))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	if int32(0) < v351 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L12
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+400)) = v369 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_11), v44+int32(400))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(225), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	goto L9
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L12
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v450 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+480)) = v416 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_12), v44+int32(480))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(230), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	goto L9
L62:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	if v499 != 0 {
		goto L71
	} else {
		goto L72
	}
L63:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+17)))
	if v453 != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L12
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+464)) = v467 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_13), v44+int32(464))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(236), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
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
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v500 == int32(1) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L12
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+416)) = v516 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_14), v44+int32(416))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(242), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	goto L9
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v557 = F_RelationGetIndexList(m, v106)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L12
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v967)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_CheckTableNotInUse(m, v106, v48)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L12
	} else {
		goto L112
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_relation_close(m, v621, int32(1))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L12
	} else {
		goto L110
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_list_free(m, v557)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L12
	} else {
		goto L102
	}
L83:
	;
	if v557 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v561 = int32(0)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	if v562 <= v561 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v581 = v561
	goto L86
L86:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v557)+12))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v606+v581<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v621 = F_index_open(m, v610, int32(1))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L12
	} else {
		goto L89
	}
L87:
	;
	goto L82
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_relation_close(m, v621, int32(1))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L12
	} else {
		goto L100
	}
L89:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v621)+192))
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+12)))
	if v624 != int32(1) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+16)))
	if v627 != int32(1) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+18)))
	if v630 != int32(1) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v642 = F_RelationGetIndexPredicate(m, v621)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	if v642 != 0 {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	v644 = int32(*(*int16)(unsafe.Add(mBase, uint32(v623)+8)))
	if v644 <= int32(0) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v667 = int32(0)
	goto L96
L96:
	;
	v694 = int32(*(*int16)(unsafe.Add(mBase, uint32(v623+int32(48)+v667<<(uint(int32(1))%32)))))
	if v694 <= int32(0) {
		goto L88
	} else {
		goto L98
	}
L97:
	;
	goto L81
L98:
	;
	v698 = v667 + int32(1)
	if v644 != v698 {
		v667 = v698
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v754 = v581 + int32(1)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	if v754 < v755 {
		v581 = v754
		goto L86
	} else {
		goto L101
	}
L101:
	;
	goto L87
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errcode(m, int32(325))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v845 = F_get_namespace_name(m, v835)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v859 = F_quote_qualified_identifier(m, v845, v847+int32(4))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+432)) = v859
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_15), v44+int32(432))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errhint(m, int32(_a_F_RefreshMatViewByOid_16), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(276), int32(_a_F_RefreshMatViewByOid_6))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	goto L9
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_list_free(m, v557)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L12
	} else {
		goto L111
	}
L111:
	;
	goto L80
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_SetMatViewPopulatedState(m, v106, l3^int32(1))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
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
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1025 = base.I32_extend8_s(v1012)
	v1027 = F_make_new_heap(m, l1, v1014, v1015, v1025, int32(7))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L12
	} else {
		goto L119
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1004 = F_GetDefaultTablespace(m, int32(116), int32(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L12
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008)+118)))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+92))
	v1011 = v67
	v1012 = v1009
	v1013 = v1008
	v1014 = v1010
	goto L114
L118:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v106)+48))
	v1011 = v1004
	v1012 = int32(116)
	v1013 = v1006
	v1014 = v1004
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1042 = F_palloc0(m, int32(40))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L12
	} else {
		goto L123
	}
L121:
	;
	v1298 = int64(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+20)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+16)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+12)) = int32(562)
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+8)) = int32(563)
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+4)) = int32(564)
	*(*int32)(unsafe.Add(mBase, uint32(v1042))) = int32(565)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1064 = F_copyObjectImpl(m, v968)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_AcquireRewriteLocks(m, v1064, int32(1), int32(0))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1088 = F_QueryRewrite(m, v1064)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L12
	} else {
		goto L127
	}
L126:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+12))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)))
	v1138 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[4]))
	if v1138 != 0 {
		goto L135
	} else {
		goto L136
	}
L127:
	;
	if v1088 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	if v1090 == int32(1) {
		goto L126
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L12
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v44)+448)) = v51
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_17), v44+int32(448))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L12
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(422), int32(_a_F_RefreshMatViewByOid_18))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L12
	} else {
		goto L134
	}
L134:
	;
	goto L9
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_ProcessInterrupts(m)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L12
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1161 = F_pg_plan_query(m, v1136, l5, int32(2048), int32(0))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L12
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1173 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[5]))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1173)))
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_PushCopiedSnapshot(m, v1174)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L12
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1207 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[5]))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1207)))
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v1218 = int32(0)
	v1222 = F_CreateQueryDesc(m, v1161, l5, v1208, v1218, v1042, v1218, v1218, v1218)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L12
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_ExecutorStart(m, v1222, int32(0))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_ExecutorRun(m, v1222, int32(1), int64(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+40))
	v1250 = *(*int64)(unsafe.Add(mBase, uint32(v1249)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_ExecutorFinish(m, v1222)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_ExecutorEnd(m, v1222)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_FreeQueryDesc(m, v1222)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L12
	} else {
		goto L150
	}
L150:
	;
	v1298 = v1250
	goto L122
L151:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[6]))
	v1305 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7]))
	v1307 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8]))
	goto L152
L152:
	;
	v1309 = v44 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v1309)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1309))) = v44 + int32(508)
	goto L155
L153:
	;
	v1323 = v106
	v1324 = v1027
	v1325 = v109
	v1326 = v1303
	v1327 = v1305
	v1328 = v1307
	v1329 = v1011
	v1330 = v159
	v1331 = int32(0)
	v1354 = v1298
	goto L18
L155:
	;
	goto L153
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7])) = v44 + int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v44)+668))
	v1373 = v44 + int32(676)
	F_initStringInfo(m, v1373)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L12
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[6])) = v1326
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7])) = v1327
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8])) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_pg_re_throw(m)
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L12
	} else {
		goto L270
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1386 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L12
	} else {
		goto L160
	}
L160:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+48))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1399 = F_get_namespace_name(m, v1389)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L12
	} else {
		goto L161
	}
L161:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1413 = F_quote_qualified_identifier(m, v1399, v1401+int32(4))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L12
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1425 = F_table_open(m, v1324, int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L12
	} else {
		goto L163
	}
L163:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+48))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1438 = F_get_namespace_name(m, v1428)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L12
	} else {
		goto L164
	}
L164:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1452 = F_quote_qualified_identifier(m, v1438, v1440+int32(4))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L12
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1464 = v44 + int32(692)
	F_initStringInfo(m, v1464)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_appendStringInfoString(m, v1464, v1452)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+368)) = int32(2)
	F_appendStringInfo(m, v1464, int32(_a_F_RefreshMatViewByOid_19), v44+int32(368))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L12
	} else {
		goto L168
	}
L168:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v44)+692))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+48))
	v1496 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1495)+120)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L12
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+352)) = v1452
	F_appendStringInfo(m, v1373, int32(_a_F_RefreshMatViewByOid_20), v44+int32(352))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L12
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1534 = F_SPI_exec(m, v1533)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L12
	} else {
		goto L171
	}
L171:
	;
	if v1534 != int32(4) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L12
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1591 = v44 + int32(676)
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1591)))
	v1593 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1592))) = uint8(v1593)
	*(*int32)(unsafe.Add(mBase, uint32(v1591)+12)) = v1593
	*(*int32)(unsafe.Add(mBase, uint32(v1591)+4)) = v1593
	goto L178
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+336)) = v1560
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(336))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L12
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(647), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L12
	} else {
		goto L177
	}
L177:
	;
	goto L9
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+328)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v44)+324)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v44)+320)) = v1452
	F_appendStringInfo(m, v1591, int32(_a_F_RefreshMatViewByOid_23), v44+int32(320))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L12
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1628 = F_SPI_execute(m, v1625, int32(0), int32(1))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L12
	} else {
		goto L180
	}
L180:
	;
	if v1628 != int32(5) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L12
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v1676 = *(*int64)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[9]))
	if v1676 != int64(0) {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+304)) = v1654
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(304))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L12
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(670), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L12
	} else {
		goto L186
	}
L186:
	;
	goto L9
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L12
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v1371 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v1325
	goto L196
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errcode(m, int32(66))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L12
	} else {
		goto L191
	}
L191:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+288)) = v1704 + int32(4)
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_24), v44+int32(288))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L12
	} else {
		goto L192
	}
L192:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[10]))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1723)))
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1723)+4))
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1725)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1737 = F_SPI_getvalue(m, v1726, v1724, int32(1))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L12
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v1737
	F_errdetail(m, int32(_a_F_RefreshMatViewByOid_25), v44+int32(272))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L12
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(685), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L12
	} else {
		goto L195
	}
L195:
	;
	goto L9
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1793 = v44 + int32(676)
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1793)))
	v1795 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1794))) = uint8(v1795)
	*(*int32)(unsafe.Add(mBase, uint32(v1793)+12)) = v1795
	*(*int32)(unsafe.Add(mBase, uint32(v1793)+4)) = v1795
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+256)) = v1494
	F_appendStringInfo(m, v1793, int32(_a_F_RefreshMatViewByOid_26), v44+int32(256))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L12
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1826 = F_SPI_exec(m, v1825)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L12
	} else {
		goto L199
	}
L199:
	;
	if v1826 != int32(4) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L12
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v1371 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v1325
	goto L206
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+240)) = v1852
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(240))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L12
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(703), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L12
	} else {
		goto L205
	}
L205:
	;
	goto L9
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1898 = v44 + int32(676)
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1898)))
	v1900 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1899))) = uint8(v1900)
	*(*int32)(unsafe.Add(mBase, uint32(v1898)+12)) = v1900
	*(*int32)(unsafe.Add(mBase, uint32(v1898)+4)) = v1900
	goto L207
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+228)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v44)+224)) = v1494
	F_appendStringInfo(m, v1898, int32(_a_F_RefreshMatViewByOid_27), v44+int32(224))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L12
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v1932 = F_SPI_exec(m, v1931)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L12
	} else {
		goto L209
	}
L209:
	;
	if v1932 != int32(4) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L12
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1989 = v44 + int32(676)
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1989)))
	v1991 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1990))) = uint8(v1991)
	*(*int32)(unsafe.Add(mBase, uint32(v1989)+12)) = v1991
	*(*int32)(unsafe.Add(mBase, uint32(v1989)+4)) = v1991
	goto L216
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+208)) = v1958
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(208))
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L12
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(711), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L12
	} else {
		goto L215
	}
L215:
	;
	goto L9
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+204)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v44)+200)) = v1413
	*(*int32)(unsafe.Add(mBase, uint32(v44)+196)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v44)+192)) = v1494
	F_appendStringInfo(m, v1989, int32(_a_F_RefreshMatViewByOid_28), v44+int32(192))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L12
	} else {
		goto L217
	}
L217:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2027 = F_palloc0(m, v1496<<(uint(int32(2))%32))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L12
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2038 = F_RelationGetIndexList(m, v1386)
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L12
	} else {
		goto L219
	}
L219:
	;
	if v2038 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v3221 = int32(0)
	goto L14
L221:
	;
	goto L222
L222:
	;
	v2043 = int32(0)
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+4))
	if v2045 <= v2043 {
		v3221 = v2038
		goto L14
	} else {
		goto L223
	}
L223:
	;
	v2064 = v2043
	v2075 = v2043
	goto L224
L224:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+12))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2089+v2064<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2104 = F_index_open(m, v2093, int32(3))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L12
	} else {
		goto L227
	}
L225:
	;
	goto L16
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_relation_close(m, v2104, int32(0))
	mBase = m.M
	v2515 = m.ExcPending
	if v2515 != 0 {
		goto L12
	} else {
		goto L268
	}
L227:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+192))
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2106)+12)))
	if v2107 != int32(1) {
		v2490 = v2075
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v2110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2106)+16)))
	if v2110 != int32(1) {
		v2490 = v2075
		goto L226
	} else {
		goto L229
	}
L229:
	;
	v2113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2106)+18)))
	if v2113 != int32(1) {
		v2490 = v2075
		goto L226
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2125 = F_RelationGetIndexPredicate(m, v2104)
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L12
	} else {
		goto L231
	}
L231:
	;
	if v2125 != 0 {
		v2490 = v2075
		goto L226
	} else {
		goto L232
	}
L232:
	;
	v2127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2106)+8)))
	if v2127 <= int32(0) {
		v2490 = v2075
		goto L226
	} else {
		goto L233
	}
L233:
	;
	v2150 = int32(0)
	goto L234
L234:
	;
	v2177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2106+int32(48)+v2150<<(uint(int32(1))%32)))))
	if v2177 <= int32(0) {
		v2490 = v2075
		goto L226
	} else {
		goto L236
	}
L235:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+192))
	v2184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2183)+10)))
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2197 = F_SysCacheGetAttrNotNull(m, int32(34), v2185, int32(18))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L12
	} else {
		goto L238
	}
L236:
	;
	v2181 = v2150 + int32(1)
	if v2181 != v2127 {
		v2150 = v2181
		goto L234
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	if v2184 <= int32(0) {
		v2490 = v2075
		goto L226
	} else {
		goto L239
	}
L239:
	;
	v2223 = int32(0)
	v2233 = v2075
	goto L240
L240:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v2015)))
	v2254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2183+int32(48)+v2223<<(uint(int32(1))%32)))))
	v2259 = v2015 + v2247<<(uint(int32(4))%32) + v2254*int32(100) - int32(80)
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(v2259)+68))
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v2197+int32(24)+v2223<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2275 = F_SearchSysCache1(m, int32(14), v2264)
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L12
	} else {
		goto L242
	}
L241:
	;
	v2490 = v2459
	goto L226
L242:
	;
	if v2275 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
		goto L12
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2275)+16))
	v2322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2321)+22)))
	v2323 = v2321 + v2322
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+84))
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_ReleaseCatCache(m, v2275)
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L12
	} else {
		goto L249
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v2264
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_29), v44+int32(16))
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L12
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(774), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L12
	} else {
		goto L248
	}
L248:
	;
	goto L9
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2347 = F_get_opfamily_member_for_cmptype(m, v2325, v2324, v2324, int32(3))
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L12
	} else {
		goto L250
	}
L250:
	;
	if v2347 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L12
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v2399 = v2027 + (v2254-int32(1))<<(uint(int32(2))%32)
	v2400 = *(*int32)(unsafe.Add(mBase, uint32(v2399)))
	if v2347 != v2400 {
		goto L257
	} else {
		goto L258
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+40)) = v2325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v2324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v2324
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_30), v44+int32(32))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L12
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(783), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L12
	} else {
		goto L256
	}
L256:
	;
	goto L9
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2399))) = v2347
	if v2233 != 0 {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	v2459 = v2233
	goto L259
L259:
	;
	v2461 = v2223 + int32(1)
	if v2461 != v2184 {
		v2223 = v2461
		v2233 = v2459
		goto L240
	} else {
		goto L267
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_appendStringInfoString(m, v44+int32(676), int32(_a_F_RefreshMatViewByOid_31))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L12
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2428 = v2259 + int32(4)
	v2429 = F_quote_qualified_identifier(m, int32(_a_F_RefreshMatViewByOid_32), v2428)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L12
	} else {
		goto L264
	}
L263:
	;
	goto L262
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2441 = F_quote_qualified_identifier(m, int32(_a_F_RefreshMatViewByOid_33), v2428)
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L12
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_generate_operator_clause(m, v44+int32(676), v2429, v2260, v2347, v2441, v2260)
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L12
	} else {
		goto L266
	}
L266:
	;
	v2459 = int32(1)
	goto L259
L267:
	;
	goto L241
L268:
	;
	v2517 = v2064 + int32(1)
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(v2038)+4))
	if v2517 < v2518 {
		v2064 = v2517
		v2075 = v2490
		goto L224
	} else {
		goto L269
	}
L269:
	;
	goto L225
L270:
	;
	goto L9
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	v2559 = int32(0)
	v2561 = int32(1)
	F_finish_heap_swap(m, l1, v1027, v2559, v2559, v2561, v2561, v2547, v2548, v1025)
	mBase = m.M
	v2564 = m.ExcPending
	if v2564 != 0 {
		goto L12
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_pgstat_count_truncate(m, v106)
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L12
	} else {
		goto L273
	}
L273:
	;
	if l3 != 0 {
		v3135 = v106
		v3136 = v1027
		v3137 = v109
		v3138 = v64
		v3139 = v65
		v3140 = v66
		v3141 = v1011
		v3142 = v159
		v3166 = v1298
		goto L15
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1011
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v106
	F_pgstat_count_heap_insert(m, v106, v1298)
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L12
	} else {
		goto L275
	}
L275:
	;
	v3135 = v106
	v3136 = v1027
	v3137 = v109
	v3138 = v64
	v3139 = v65
	v3140 = v66
	v3141 = v1011
	v3142 = v159
	v3166 = v1298
	goto L15
L276:
	;
	if v2490 == int32(0) {
		goto L13
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_appendStringInfoString(m, v44+int32(676), int32(_a_F_RefreshMatViewByOid_34))
	mBase = m.M
	v2613 = m.ExcPending
	if v2613 != 0 {
		goto L12
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2624 = F_SPI_exec(m, v2623)
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L12
	} else {
		goto L279
	}
L279:
	;
	if v2624 != int32(7) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L12
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2681 = v44 + int32(676)
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(v2681)))
	v2683 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2682))) = uint8(v2683)
	*(*int32)(unsafe.Add(mBase, uint32(v2681)+12)) = v2683
	*(*int32)(unsafe.Add(mBase, uint32(v2681)+4)) = v2683
	goto L286
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+176)) = v2650
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(176))
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L12
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(848), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L12
	} else {
		goto L285
	}
L285:
	;
	goto L9
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+160)) = v1494
	F_appendStringInfo(m, v2681, int32(_a_F_RefreshMatViewByOid_20), v44+int32(160))
	mBase = m.M
	v2703 = m.ExcPending
	if v2703 != 0 {
		goto L12
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2714 = F_SPI_exec(m, v2713)
	mBase = m.M
	v2715 = m.ExcPending
	if v2715 != 0 {
		goto L12
	} else {
		goto L288
	}
L288:
	;
	if v2714 != int32(4) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L12
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v2761 = int32(_a_F_RefreshMatViewByOid_35)
	v2763 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8])) = v2763 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2777 = v44 + int32(676)
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v2777)))
	v2779 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2778))) = uint8(v2779)
	*(*int32)(unsafe.Add(mBase, uint32(v2777)+12)) = v2779
	*(*int32)(unsafe.Add(mBase, uint32(v2777)+4)) = v2779
	goto L295
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+144)) = v2740
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(144))
	mBase = m.M
	v2746 = m.ExcPending
	if v2746 != 0 {
		goto L12
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(859), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		goto L12
	} else {
		goto L294
	}
L294:
	;
	goto L9
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+132)) = v1494
	*(*int32)(unsafe.Add(mBase, uint32(v44)+128)) = v1413
	F_appendStringInfo(m, v2777, int32(_a_F_RefreshMatViewByOid_36), v44+int32(128))
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L12
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2811 = F_SPI_exec(m, v2810)
	mBase = m.M
	v2812 = m.ExcPending
	if v2812 != 0 {
		goto L12
	} else {
		goto L297
	}
L297:
	;
	if v2811 != int32(8) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		goto L12
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2868 = v44 + int32(676)
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2868)))
	v2870 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2869))) = uint8(v2870)
	*(*int32)(unsafe.Add(mBase, uint32(v2868)+12)) = v2870
	*(*int32)(unsafe.Add(mBase, uint32(v2868)+4)) = v2870
	goto L304
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+112)) = v2837
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(112))
	mBase = m.M
	v2843 = m.ExcPending
	if v2843 != 0 {
		goto L12
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(872), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v2857 = m.ExcPending
	if v2857 != 0 {
		goto L12
	} else {
		goto L303
	}
L303:
	;
	goto L9
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+100)) = v1494
	*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v1413
	F_appendStringInfo(m, v2868, int32(_a_F_RefreshMatViewByOid_37), v44+int32(96))
	mBase = m.M
	v2891 = m.ExcPending
	if v2891 != 0 {
		goto L12
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v2902 = F_SPI_exec(m, v2901)
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L12
	} else {
		goto L306
	}
L306:
	;
	if v2902 != int32(7) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L12
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v2949 = int32(_a_F_RefreshMatViewByOid_35)
	v2951 = *(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[8])) = v2951 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_relation_close(m, v1425, int32(0))
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		goto L12
	} else {
		goto L313
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2928 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v2928
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(80))
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		goto L12
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(881), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v2948 = m.ExcPending
	if v2948 != 0 {
		goto L12
	} else {
		goto L312
	}
L312:
	;
	goto L9
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_relation_close(m, v1386, int32(0))
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L12
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v2989 = v44 + int32(676)
	v2990 = *(*int32)(unsafe.Add(mBase, uint32(v2989)))
	v2991 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2990))) = uint8(v2991)
	*(*int32)(unsafe.Add(mBase, uint32(v2989)+12)) = v2991
	*(*int32)(unsafe.Add(mBase, uint32(v2989)+4)) = v2991
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44)+68)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v44)+64)) = v1494
	F_appendStringInfo(m, v2989, int32(_a_F_RefreshMatViewByOid_38), v44-int32(-64))
	mBase = m.M
	v3012 = m.ExcPending
	if v3012 != 0 {
		goto L12
	} else {
		goto L316
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	v3023 = F_SPI_exec(m, v3022)
	mBase = m.M
	v3024 = m.ExcPending
	if v3024 != 0 {
		goto L12
	} else {
		goto L317
	}
L317:
	;
	if v3023 != int32(4) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L12
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v3079 = F_SPI_finish(m)
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L12
	} else {
		goto L324
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(v44)+676))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v3049
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_21), v44+int32(48))
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L12
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(892), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v3069 = m.ExcPending
	if v3069 != 0 {
		goto L12
	} else {
		goto L323
	}
L323:
	;
	goto L9
L324:
	;
	if v3079 != int32(2) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3095 = m.ExcPending
	if v3095 != 0 {
		goto L12
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[7])) = v1327
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[6])) = v1326
	v3135 = v1323
	v3136 = v1324
	v3137 = v1325
	v3138 = v1326
	v3139 = v1327
	v3140 = v1328
	v3141 = v1329
	v3142 = v1330
	v3166 = v1354
	goto L15
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errmsg_internal(m, int32(_a_F_RefreshMatViewByOid_39), int32(0))
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L12
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(896), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v3122 = m.ExcPending
	if v3122 != 0 {
		goto L12
	} else {
		goto L330
	}
L330:
	;
	goto L9
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v3139
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v3138
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v3140
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v3166
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v3136
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v3141
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v3142
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v3137
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v3135
	F_AtEOXact_GUC(m, int32(0), v3142)
	mBase = m.M
	v3191 = m.ExcPending
	if v3191 != 0 {
		goto L12
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v3138
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v3139
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v3140
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v3166
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v3136
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v3141
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v3142
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v3137
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v3135
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v44)+672))
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(v44)+668))
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[2])) = v3202
	*(*int32)(unsafe.Add(mBase, _c_F_RefreshMatViewByOid[1])) = v3201
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	if l6 != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6)+8)) = v3166
	if l2 != 0 {
		goto L337
	} else {
		goto L338
	}
L335:
	;
	goto L336
L336:
	;
	m.G0 = v44 + int32(752)
	return
L337:
	;
	v3215 = int32(179)
	goto L339
L338:
	;
	v3215 = int32(169)
	goto L339
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v3215
	goto L336
L340:
	;
	goto L13
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errcode(m, int32(1088))
	mBase = m.M
	v3299 = m.ExcPending
	if v3299 != 0 {
		goto L12
	} else {
		goto L342
	}
L342:
	;
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v3300 + int32(4)
	F_errmsg(m, int32(_a_F_RefreshMatViewByOid_40), v44)
	mBase = m.M
	v3315 = m.ExcPending
	if v3315 != 0 {
		goto L12
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v1327
	*(*int32)(unsafe.Add(mBase, uint32(v44)+708)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v44)+716)) = v1328
	*(*int64)(unsafe.Add(mBase, uint32(v44)+720)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v44)+732)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v44)+748)) = v1323
	F_errfinish(m, int32(_a_F_RefreshMatViewByOid_5), int32(839), int32(_a_F_RefreshMatViewByOid_22))
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L12
	} else {
		goto L344
	}
L344:
	;
	goto L9
L345:
	;
	v3376 = int32(v3372)
	m.G0 = v44
	v3378 = *(*int32)(unsafe.Add(mBase, uint32(v3376)+4))
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3376)))
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(v3379)))
	if v44+int32(508) == v3382 {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	m.ExcPending = 1
	goto L354
L347:
	;
	if v3386 != 0 {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	v3384 = *(*int32)(unsafe.Add(mBase, uint32(v3379)+4))
	v3386 = v3384
	goto L350
L349:
	;
	v3386 = int32(0)
	goto L350
L350:
	;
	goto L347
L351:
	;
	v3387 = *(*int32)(unsafe.Add(mBase, uint32(v44)+748))
	v3388 = *(*int32)(unsafe.Add(mBase, uint32(v44)+744))
	v3389 = *(*int32)(unsafe.Add(mBase, uint32(v44)+740))
	v3390 = *(*int32)(unsafe.Add(mBase, uint32(v44)+736))
	v3391 = *(*int32)(unsafe.Add(mBase, uint32(v44)+732))
	v3392 = *(*int64)(unsafe.Add(mBase, uint32(v44)+720))
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(v44)+716))
	v3394 = *(*int32)(unsafe.Add(mBase, uint32(v44)+712))
	v3395 = *(*int32)(unsafe.Add(mBase, uint32(v44)+708))
	v61 = v3387
	v62 = v3391
	v63 = v3388
	v64 = v3395
	v65 = v3394
	v66 = v3393
	v67 = v3390
	v68 = v3389
	v69 = v3378
	v70 = v3386
	v93 = v3392
	goto L7
L352:
	;
	goto L353
L353:
	;
	F___wasm_longjmp(m, v3379, v3378)
	mBase = m.M
	v3397 = m.ExcPending
	if v3397 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	return
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
