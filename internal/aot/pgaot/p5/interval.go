package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_interval_cmp_lower_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_range_cmp_bounds(m, l2, l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_interval_cmp_upper_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v4 float64
	_ = v4
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v3) & v9
	v13 = base.I64_reinterpret_f64(v4) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v13) {
		v23 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v32 = int32(0) - v23&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
	} else {
		v18 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v10))|base.F64_gt(v3, v4) != 0 {
			v32 = v18
		} else {
			v23 = v18
			v32 = int32(0) - v23&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
		}
	}
	return v32
}
func F_interval_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v8 int32
	_ = v8
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v106 int32
	_ = v106
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v168 int32
	_ = v168
	var v185 int32
	_ = v185
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v301 int64
	_ = v301
	var v304 int64
	_ = v304
	var v307 int64
	_ = v307
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v315 int64
	_ = v315
	var v322 int64
	_ = v322
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v335 int64
	_ = v335
	var v339 int64
	_ = v339
	var v345 int64
	_ = v345
	var v347 int64
	_ = v347
	var v348 int64
	_ = v348
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v400 int64
	_ = v400
	var v401 int64
	_ = v401
	var v406 int64
	_ = v406
	var v409 int64
	_ = v409
	var v412 int64
	_ = v412
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v420 int64
	_ = v420
	var v427 int64
	_ = v427
	var v438 int64
	_ = v438
	var v439 int64
	_ = v439
	var v443 int64
	_ = v443
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v458 int64
	_ = v458
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v475 int64
	_ = v475
	var v477 int64
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v556 int32
	_ = v556
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v702 int32
	_ = v702
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v790 int32
	_ = v790
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v905 int32
	_ = v905
	var v937 int32
	_ = v937
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1010 int64
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1016 float64
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1045 int64
	_ = v1045
	var v1046 int64
	_ = v1046
	var v1051 int64
	_ = v1051
	var v1054 int64
	_ = v1054
	var v1057 int64
	_ = v1057
	var v1060 int64
	_ = v1060
	var v1061 int64
	_ = v1061
	var v1065 int64
	_ = v1065
	var v1072 int64
	_ = v1072
	var v1083 int64
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int64
	_ = v1097
	var v1100 int64
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int64
	_ = v1119
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1203 float64
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1214 float64
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 float64
	_ = v1220
	var v1221 int64
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 float64
	_ = v1226
	var v1230 int64
	_ = v1230
	var v1232 int64
	_ = v1232
	var v1238 int64
	_ = v1238
	var v1240 float64
	_ = v1240
	var v1244 int64
	_ = v1244
	var v1245 int64
	_ = v1245
	var v1254 int64
	_ = v1254
	var v1256 float64
	_ = v1256
	var v1267 int64
	_ = v1267
	var v1268 int64
	_ = v1268
	var v1284 int32
	_ = v1284
	var v1287 int64
	_ = v1287
	var v1288 int64
	_ = v1288
	var v1293 int64
	_ = v1293
	var v1296 int64
	_ = v1296
	var v1299 int64
	_ = v1299
	var v1302 int64
	_ = v1302
	var v1303 int64
	_ = v1303
	var v1307 int64
	_ = v1307
	var v1314 int64
	_ = v1314
	var v1325 int64
	_ = v1325
	var v1326 int64
	_ = v1326
	var v1331 int64
	_ = v1331
	var v1332 int64
	_ = v1332
	var v1342 float64
	_ = v1342
	var v1343 int64
	_ = v1343
	var v1345 float64
	_ = v1345
	var v1356 int64
	_ = v1356
	var v1357 int64
	_ = v1357
	var v1373 int32
	_ = v1373
	var v1376 int64
	_ = v1376
	var v1377 int64
	_ = v1377
	var v1382 int64
	_ = v1382
	var v1385 int64
	_ = v1385
	var v1388 int64
	_ = v1388
	var v1391 int64
	_ = v1391
	var v1392 int64
	_ = v1392
	var v1396 int64
	_ = v1396
	var v1403 int64
	_ = v1403
	var v1414 int64
	_ = v1414
	var v1415 int64
	_ = v1415
	var v1420 int64
	_ = v1420
	var v1421 int64
	_ = v1421
	var v1431 float64
	_ = v1431
	var v1432 int64
	_ = v1432
	var v1434 float64
	_ = v1434
	var v1445 int64
	_ = v1445
	var v1446 int64
	_ = v1446
	var v1464 int32
	_ = v1464
	var v1467 int64
	_ = v1467
	var v1468 int64
	_ = v1468
	var v1473 int64
	_ = v1473
	var v1476 int64
	_ = v1476
	var v1479 int64
	_ = v1479
	var v1482 int64
	_ = v1482
	var v1483 int64
	_ = v1483
	var v1487 int64
	_ = v1487
	var v1494 int64
	_ = v1494
	var v1505 int64
	_ = v1505
	var v1506 int64
	_ = v1506
	var v1511 int64
	_ = v1511
	var v1512 int64
	_ = v1512
	var v1522 float64
	_ = v1522
	var v1523 int64
	_ = v1523
	var v1525 float64
	_ = v1525
	var v1536 int64
	_ = v1536
	var v1537 int64
	_ = v1537
	var v1553 int32
	_ = v1553
	var v1556 int64
	_ = v1556
	var v1557 int64
	_ = v1557
	var v1562 int64
	_ = v1562
	var v1565 int64
	_ = v1565
	var v1568 int64
	_ = v1568
	var v1571 int64
	_ = v1571
	var v1572 int64
	_ = v1572
	var v1576 int64
	_ = v1576
	var v1583 int64
	_ = v1583
	var v1594 int64
	_ = v1594
	var v1595 int64
	_ = v1595
	var v1600 int64
	_ = v1600
	var v1601 int64
	_ = v1601
	var v1611 float64
	_ = v1611
	var v1612 int64
	_ = v1612
	var v1614 float64
	_ = v1614
	var v1625 int64
	_ = v1625
	var v1626 int64
	_ = v1626
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1656 float64
	_ = v1656
	var v1657 int64
	_ = v1657
	var v1659 float64
	_ = v1659
	var v1670 int64
	_ = v1670
	var v1671 int64
	_ = v1671
	var v1672 int64
	_ = v1672
	var v1692 int64
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1712 float64
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1722 float64
	_ = v1722
	var v1726 float64
	_ = v1726
	var v1727 int64
	_ = v1727
	var v1729 float64
	_ = v1729
	var v1740 int64
	_ = v1740
	var v1741 int64
	_ = v1741
	var v1742 int64
	_ = v1742
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1776 int32
	_ = v1776
	var v1778 float64
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1788 float64
	_ = v1788
	var v1792 float64
	_ = v1792
	var v1793 int64
	_ = v1793
	var v1795 float64
	_ = v1795
	var v1806 int64
	_ = v1806
	var v1807 int64
	_ = v1807
	var v1808 int64
	_ = v1808
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1862 int64
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1903 int64
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1920 int32
	_ = v1920
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1944 int64
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1961 int32
	_ = v1961
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2052 int32
	_ = v2052
	var v2060 int32
	_ = v2060
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2075 int32
	_ = v2075
	var v2076 int64
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2108 int32
	_ = v2108
	var v2153 int32
	_ = v2153
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2188 int32
	_ = v2188
	var v2191 int64
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2273 float64
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2279 float64
	_ = v2279
	var v2288 int32
	_ = v2288
	var v2296 float64
	_ = v2296
	var v2297 int64
	_ = v2297
	var v2300 float64
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2315 int32
	_ = v2315
	var v2317 int32
	_ = v2317
	var v2318 int64
	_ = v2318
	var v2326 int32
	_ = v2326
	var v2330 int32
	_ = v2330
	var v2334 int32
	_ = v2334
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2374 int32
	_ = v2374
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2394 int32
	_ = v2394
	var v2399 int64
	_ = v2399
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2441 float64
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2451 float64
	_ = v2451
	var v2456 float64
	_ = v2456
	var v2457 int64
	_ = v2457
	var v2459 float64
	_ = v2459
	var v2470 int64
	_ = v2470
	var v2471 int64
	_ = v2471
	var v2472 int64
	_ = v2472
	var v2486 int64
	_ = v2486
	var v2490 int32
	_ = v2490
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2503 int32
	_ = v2503
	var v2508 float64
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2518 float64
	_ = v2518
	var v2523 float64
	_ = v2523
	var v2524 int64
	_ = v2524
	var v2526 float64
	_ = v2526
	var v2537 int64
	_ = v2537
	var v2538 int64
	_ = v2538
	var v2539 int64
	_ = v2539
	var v2552 int32
	_ = v2552
	var v2554 int32
	_ = v2554
	var v2561 int32
	_ = v2561
	var v2566 float64
	_ = v2566
	var v2567 int64
	_ = v2567
	var v2569 float64
	_ = v2569
	var v2580 int64
	_ = v2580
	var v2581 int64
	_ = v2581
	var v2582 int64
	_ = v2582
	var v2590 int32
	_ = v2590
	var v2592 int32
	_ = v2592
	var v2599 int32
	_ = v2599
	var v2600 int64
	_ = v2600
	var v2601 int64
	_ = v2601
	var v2603 int64
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2612 int32
	_ = v2612
	var v2614 int64
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2626 float64
	_ = v2626
	var v2627 int64
	_ = v2627
	var v2629 float64
	_ = v2629
	var v2640 int64
	_ = v2640
	var v2641 int64
	_ = v2641
	var v2642 int64
	_ = v2642
	var v2654 int32
	_ = v2654
	var v2663 int32
	_ = v2663
	var v2665 int32
	_ = v2665
	var v2672 int32
	_ = v2672
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2684 int32
	_ = v2684
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2698 int64
	_ = v2698
	var v2704 int32
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2714 float64
	_ = v2714
	var v2717 int32
	_ = v2717
	var v2719 float64
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2723 int32
	_ = v2723
	var v2730 float64
	_ = v2730
	var v2734 float64
	_ = v2734
	var v2735 int64
	_ = v2735
	var v2737 float64
	_ = v2737
	var v2748 int64
	_ = v2748
	var v2749 int64
	_ = v2749
	var v2750 int64
	_ = v2750
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2779 int32
	_ = v2779
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int64
	_ = v2789
	var v2795 int32
	_ = v2795
	var v2797 int32
	_ = v2797
	var v2804 float64
	_ = v2804
	var v2808 float64
	_ = v2808
	var v2809 int64
	_ = v2809
	var v2811 float64
	_ = v2811
	var v2822 int64
	_ = v2822
	var v2823 int64
	_ = v2823
	var v2824 int64
	_ = v2824
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2847 int64
	_ = v2847
	var v2848 int64
	_ = v2848
	var v2853 int64
	_ = v2853
	var v2856 int64
	_ = v2856
	var v2859 int64
	_ = v2859
	var v2862 int64
	_ = v2862
	var v2863 int64
	_ = v2863
	var v2867 int64
	_ = v2867
	var v2874 int64
	_ = v2874
	var v2885 int64
	_ = v2885
	var v2886 int64
	_ = v2886
	var v2891 int64
	_ = v2891
	var v2892 int64
	_ = v2892
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2904 float64
	_ = v2904
	var v2905 int64
	_ = v2905
	var v2907 float64
	_ = v2907
	var v2918 int64
	_ = v2918
	var v2919 int64
	_ = v2919
	var v2927 int32
	_ = v2927
	var v2930 int64
	_ = v2930
	var v2931 int64
	_ = v2931
	var v2936 int64
	_ = v2936
	var v2939 int64
	_ = v2939
	var v2942 int64
	_ = v2942
	var v2945 int64
	_ = v2945
	var v2946 int64
	_ = v2946
	var v2950 int64
	_ = v2950
	var v2957 int64
	_ = v2957
	var v2968 int64
	_ = v2968
	var v2969 int64
	_ = v2969
	var v2974 int64
	_ = v2974
	var v2975 int64
	_ = v2975
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2987 float64
	_ = v2987
	var v2988 int64
	_ = v2988
	var v2990 float64
	_ = v2990
	var v3001 int64
	_ = v3001
	var v3002 int64
	_ = v3002
	var v3010 int32
	_ = v3010
	var v3013 int64
	_ = v3013
	var v3014 int64
	_ = v3014
	var v3019 int64
	_ = v3019
	var v3022 int64
	_ = v3022
	var v3025 int64
	_ = v3025
	var v3028 int64
	_ = v3028
	var v3029 int64
	_ = v3029
	var v3033 int64
	_ = v3033
	var v3040 int64
	_ = v3040
	var v3051 int64
	_ = v3051
	var v3052 int64
	_ = v3052
	var v3057 int64
	_ = v3057
	var v3058 int64
	_ = v3058
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3070 float64
	_ = v3070
	var v3071 int64
	_ = v3071
	var v3073 float64
	_ = v3073
	var v3084 int64
	_ = v3084
	var v3085 int64
	_ = v3085
	var v3092 int32
	_ = v3092
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3100 int32
	_ = v3100
	var v3102 int32
	_ = v3102
	var v3103 int64
	_ = v3103
	var v3111 int32
	_ = v3111
	var v3115 int32
	_ = v3115
	var v3119 int32
	_ = v3119
	var v3125 int32
	_ = v3125
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3159 int32
	_ = v3159
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3169 int32
	_ = v3169
	var v3179 int32
	_ = v3179
	var v3184 int32
	_ = v3184
	var v3186 int64
	_ = v3186
	var v3189 int64
	_ = v3189
	var v3190 int64
	_ = v3190
	var v3195 int64
	_ = v3195
	var v3198 int64
	_ = v3198
	var v3201 int64
	_ = v3201
	var v3204 int64
	_ = v3204
	var v3205 int64
	_ = v3205
	var v3209 int64
	_ = v3209
	var v3216 int64
	_ = v3216
	var v3227 int32
	_ = v3227
	var v3228 int64
	_ = v3228
	var v3229 int64
	_ = v3229
	var v3233 int64
	_ = v3233
	var v3234 int64
	_ = v3234
	var v3240 int64
	_ = v3240
	var v3241 int64
	_ = v3241
	var v3243 int64
	_ = v3243
	var v3245 int64
	_ = v3245
	var v3246 int64
	_ = v3246
	var v3256 int64
	_ = v3256
	var v3257 int64
	_ = v3257
	var v3265 int64
	_ = v3265
	var v3267 float64
	_ = v3267
	var v3278 int64
	_ = v3278
	var v3279 int64
	_ = v3279
	var v3288 int32
	_ = v3288
	var v3291 int64
	_ = v3291
	var v3292 int64
	_ = v3292
	var v3297 int64
	_ = v3297
	var v3300 int64
	_ = v3300
	var v3303 int64
	_ = v3303
	var v3306 int64
	_ = v3306
	var v3307 int64
	_ = v3307
	var v3311 int64
	_ = v3311
	var v3318 int64
	_ = v3318
	var v3329 int64
	_ = v3329
	var v3330 int64
	_ = v3330
	var v3335 int64
	_ = v3335
	var v3336 int64
	_ = v3336
	var v3346 float64
	_ = v3346
	var v3347 int64
	_ = v3347
	var v3349 float64
	_ = v3349
	var v3360 int64
	_ = v3360
	var v3361 int64
	_ = v3361
	var v3381 int32
	_ = v3381
	var v3382 int32
	_ = v3382
	var v3383 int64
	_ = v3383
	var v3384 float64
	_ = v3384
	var v3385 int64
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3394 int64
	_ = v3394
	var v3399 int64
	_ = v3399
	var v3400 int64
	_ = v3400
	var v3404 int64
	_ = v3404
	var v3405 int64
	_ = v3405
	var v3415 float64
	_ = v3415
	var v3416 int64
	_ = v3416
	var v3418 float64
	_ = v3418
	var v3429 int64
	_ = v3429
	var v3430 int64
	_ = v3430
	var v3441 int32
	_ = v3441
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3458 int32
	_ = v3458
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3469 int64
	_ = v3469
	var v3470 float64
	_ = v3470
	var v3471 int64
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3476 int32
	_ = v3476
	var v3478 int32
	_ = v3478
	var v3480 int64
	_ = v3480
	var v3485 int64
	_ = v3485
	var v3486 int64
	_ = v3486
	var v3490 int64
	_ = v3490
	var v3491 int64
	_ = v3491
	var v3501 float64
	_ = v3501
	var v3502 int64
	_ = v3502
	var v3504 float64
	_ = v3504
	var v3515 int64
	_ = v3515
	var v3516 int64
	_ = v3516
	var v3527 int32
	_ = v3527
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3541 int32
	_ = v3541
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3548 int32
	_ = v3548
	var v3559 int32
	_ = v3559
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3567 int32
	_ = v3567
	var v3610 int32
	_ = v3610
	var v3624 int32
	_ = v3624
	var v3669 int32
	_ = v3669
	var v3695 int32
	_ = v3695
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3704 int32
	_ = v3704
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3709 int64
	_ = v3709
	var v3710 int64
	_ = v3710
	var v3713 int64
	_ = v3713
	var v3719 int32
	_ = v3719
	var v3721 int64
	_ = v3721
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3729 int32
	_ = v3729
	var v3733 int32
	_ = v3733
	var v3738 int32
	_ = v3738
	var v3746 int32
	_ = v3746
	var v3748 int32
	_ = v3748
	var v3752 int32
	_ = v3752
	var v3757 int32
	_ = v3757
	var v3763 int32
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3768 int64
	_ = v3768
	v2 = int64(0)
	v8 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(528)
	m.G0 = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+504)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v39)+512)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v39)+520)) = v8
	v54 = v39 + int32(384)
	v56 = v39 + int32(272)
	v59 = F_ParseDateTime(m, v43, v39+int32(16), int32(256), v54, v56, v39+int32(496))
	mBase = m.M
	if v59 == v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+496))
	if v42 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v2153 = v59
	goto L3
L3:
	;
	if v2153 == int32(-1) {
		goto L414
	} else {
		goto L415
	}
L4:
	;
	v68 = int32(_a_F_interval_in_0)
	goto L6
L5:
	;
	v68 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
	goto L6
L6:
	;
	v69 = m.G0
	v71 = v69 - int32(160)
	m.G0 = v71
	v74 = v39 + int32(500)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(17)
	v78 = v39 + int32(504)
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v79
	v81 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v81
	if v62 <= v79 {
		v168 = v8
		goto L10
	} else {
		goto L11
	}
L7:
	;
	m.G0 = v71 + int32(160)
	v2153 = v2108
	goto L3
L8:
	;
	v247 = int32(8)
	v248 = v8
	v252 = v204
	v264 = v8
	v272 = v8
	goto L19
L9:
	;
	v204 = v62 - int32(1)
	v213 = base.B2i32(v92 == int32(45))
	goto L8
L10:
	;
	v185 = v62 - int32(1)
	if v185 < int32(0) {
		v2108 = int32(-1)
		goto L7
	} else {
		goto L18
	}
L11:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[0]))
	if v88 != int32(2) {
		v168 = v8
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if base.B2i32(v92 != int32(45))|base.B2i32(base.Ui32(v62) < base.Ui32(int32(2))) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v106 = int32(1)
	goto L14
L14:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v54+v106<<(uint(int32(2))%32))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	switch v140 - int32(43) {
	case 0, 2:
		v168 = int32(0)
		goto L10
	default:
		goto L16
	}
L15:
	;
	v168 = v143
	goto L10
L16:
	;
	v143 = int32(1)
	v145 = v106 + v143
	if v145 != v62 {
		v106 = v145
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v204 = v185
	v213 = v168
	goto L8
L19:
	;
	v275 = int32(-1)
	v277 = v252 << (uint(int32(2)) % 32)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v56+v277)))
	switch v279 {
	case 0, 2:
		goto L26
	case 1, 6:
		goto L25
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		v2108 = v275
		goto L7
	}
L20:
	;
	v2067 = int32(0)
	v2070 = v2036 | base.B2i32(v2052 == v2067)
	if v2070|base.B2i32(v2060 == v2067) != 0 {
		v2108 = v2067 - v2070
		goto L7
	} else {
		goto L409
	}
L21:
	;
	if int32(0) < v252 {
		v247 = v2035
		v248 = v2036
		v252 = v252 - int32(1)
		v264 = v2052
		v272 = v2060
		goto L19
	} else {
		goto L408
	}
L22:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v71)+112))
	if v2023&v264 != 0 {
		goto L405
	} else {
		goto L406
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_interval_in[1])) = int32(0)
	v1004 = v277 + v54
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1010 = F_strtox_2(m, v1005, v71+int32(116), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L165
L24:
	;
	v1000 = int32(19)
	goto L23
L25:
	;
	if v248&int32(1) != 0 {
		v2108 = v275
		goto L7
	} else {
		goto L68
	}
L26:
	;
	if v247 != int32(8) {
		v1000 = v247
		goto L23
	} else {
		goto L57
	}
L27:
	;
	v373 = v277 + v54
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v376 = v374 + int32(1)
	v377 = int32(58)
	v378 = F___strchrnul(m, v376, v377)
	mBase = m.M
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v380 == v377 {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277+v54)))
	v286 = F_DecodeTimeCommon(m, v281, v68, v71+int32(112), v71+int32(120))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	if v286 != 0 {
		v2108 = v286
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v290 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v290
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v71)+136))
	v295 = int64(3600000000)
	v296 = int64(0)
	v301 = int64(32)
	v304 = int64(base.Ui64(v292) >> (uint(v301) % 64))
	v307 = int64(4294967295)
	v310 = v292 & v307
	v311 = v295 * v310
	v315 = int64(base.Ui64(v311)>>(uint(v301)%64)) + v295*v304
	v322 = v310*v296 + v315&v307
	*(*int64)(unsafe.Add(mBase, uint32(v71)+8)) = v292*v296 + v292>>(uint(int64(63))%64)*v295 + v296*v304 + int64(base.Ui64(v315)>>(uint(v301)%64)) + int64(base.Ui64(v322)>>(uint(v301)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v71))) = v311&v307 | v322<<(uint(v301)%64)
	goto L32
L32:
	;
	v333 = int32(-2)
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	if v334 != v335>>(uint(int64(63))%64) {
		v2108 = v333
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v339 = v290 + v335
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v339
	if base.B2i32(v335 < int64(0))^base.B2i32(v339 < v290) != 0 {
		v2108 = v333
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v345 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+128)))
	v347 = v345 * int64(60000000)
	v348 = v339 + v347
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v348
	if base.B2i32(v347 < int64(0))^base.B2i32(v348 < v339) != 0 {
		v2108 = v333
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v354 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+124)))
	v356 = v354 * int64(1000000)
	v357 = v348 + v356
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v357
	if base.B2i32(v356 < int64(0))^base.B2i32(v357 < v348) != 0 {
		v2108 = v333
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v363 = int32(21)
	v364 = int32(0)
	if base.B2i32(v213 == v364)|base.B2i32(v357 <= int64(0)) != 0 {
		v1995 = v363
		v1996 = v364
		v2020 = v272
		goto L22
	} else {
		goto L37
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = int64(0) - v357
	v1995 = v363
	v1996 = v364
	v2020 = v272
	goto L22
L38:
	;
	if v384 == int32(0) {
		goto L26
	} else {
		goto L42
	}
L39:
	;
	v384 = v378
	goto L41
L40:
	;
	v384 = int32(0)
	goto L41
L41:
	;
	goto L38
L42:
	;
	v391 = F_DecodeTimeCommon(m, v376, v68, v71+int32(112), v71+int32(120))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	if v391 != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v393 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v393
	v396 = v71 + int32(96)
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v71)+136))
	v400 = int64(3600000000)
	v401 = int64(0)
	v406 = int64(32)
	v409 = int64(base.Ui64(v397) >> (uint(v406) % 64))
	v412 = int64(4294967295)
	v415 = v397 & v412
	v416 = v400 * v415
	v420 = int64(base.Ui64(v416)>>(uint(v406)%64)) + v400*v409
	v427 = v415*v401 + v420&v412
	*(*int64)(unsafe.Add(mBase, uint32(v396)+8)) = v397*v401 + v397>>(uint(int64(63))%64)*v400 + v401*v409 + int64(base.Ui64(v420)>>(uint(v406)%64)) + int64(base.Ui64(v427)>>(uint(v406)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v396))) = v416&v412 | v427<<(uint(v406)%64)
	goto L45
L45:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v71)+104))
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v71)+96))
	if v438 != v439>>(uint(int64(63))%64) {
		goto L26
	} else {
		goto L46
	}
L46:
	;
	v443 = v393 + v439
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v443
	if base.B2i32(v439 < int64(0))^base.B2i32(v443 < v393) != 0 {
		goto L26
	} else {
		goto L47
	}
L47:
	;
	v449 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+128)))
	v451 = v449 * int64(60000000)
	v452 = v443 + v451
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v452
	if base.B2i32(v451 < int64(0))^base.B2i32(v452 < v443) != 0 {
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v458 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+124)))
	v460 = v458 * int64(1000000)
	v461 = v452 + v460
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v461
	if base.B2i32(v460 < int64(0))^base.B2i32(v461 < v452) != 0 {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	if v468 == int32(45) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v461 == int64(-9223372036854775807-1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v477 = v461
	goto L52
L52:
	;
	v478 = int32(21)
	v479 = int32(0)
	if v213&base.B2i32(int64(0) < v477) == v479 {
		v1995 = v478
		v1996 = v479
		v2020 = v272
		goto L22
	} else {
		goto L56
	}
L53:
	;
	v2108 = int32(-2)
	goto L7
L54:
	;
	goto L55
L55:
	;
	v475 = int64(0) - v461
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v475
	v477 = v475
	goto L52
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = int64(0) - v477
	v1995 = v478
	v1996 = v479
	v2020 = v272
	goto L22
L57:
	;
	if base.B2i32(int32(1023) < v68) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v1000 = int32(18)
	goto L23
L59:
	;
	if v68 == int32(2048) {
		goto L24
	} else {
		goto L67
	}
L60:
	;
	v1000 = int32(20)
	goto L23
L61:
	;
	v1000 = int32(21)
	goto L23
L62:
	;
	v1000 = int32(23)
	goto L23
L63:
	;
	switch v68 - int32(2) {
	case 0, 4:
		goto L62
	default:
		goto L58
	case 2:
		v1000 = int32(25)
		goto L23
	case 6:
		goto L61
	}
L64:
	;
	goto L65
L65:
	;
	switch v68 - int32(1024) {
	case 0, 8:
		goto L60
	case 1, 2, 3, 4, 5, 6, 7:
		goto L58
	default:
		goto L66
	}
L66:
	;
	switch v68 - int32(3072) {
	case 0, 8:
		goto L24
	case 1, 2, 3, 4, 5, 6, 7:
		goto L58
	default:
		goto L59
	}
L67:
	;
	goto L58
L68:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v277+v54)))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v277)+uint32(_c_F_interval_in[2])))
	if v507 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v974 = int32(8)
	v977 = v952 & int32(255)
	if v977 == v974 {
		v2035 = v974
		v2036 = int32(0)
		v2052 = v264
		v2060 = v272
		goto L21
	} else {
		goto L160
	}
L70:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v277)+uint32(_c_F_interval_in[3])))
	if v741 != 0 {
		goto L117
	} else {
		goto L118
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+uint32(_c_F_interval_in[2]))) = v670
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+11)))
	if v702 != int32(31) {
		v945 = v670
		v952 = v702
		goto L69
	} else {
		goto L115
	}
L72:
	;
	goto L77
L73:
	;
	goto L74
L74:
	;
	v556 = int32(*(*int8)(unsafe.Add(mBase, uint32(v506))))
	v567 = int32(_a_F_interval_in_1)
	v568 = int32(_a_F_interval_in_2)
	goto L89
L75:
	;
	if v545-v546 == int32(0) {
		v670 = v507
		goto L71
	} else {
		goto L88
	}
L77:
	;
	goto L78
L78:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v514 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v515 = v506
	v516 = v507
	v517 = int32(10)
	v518 = v514
	goto L83
L80:
	;
	v541 = v507
	v545 = int32(0)
	goto L81
L81:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	goto L75
L82:
	;
	v541 = v536
	v545 = v538
	goto L81
L83:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
	if base.B2i32(v518 != v520)|base.B2i32(v520 == int32(0)) != 0 {
		v536 = v516
		v538 = v518
		goto L82
	} else {
		goto L85
	}
L84:
	;
	v536 = v530
	v538 = int32(0)
	goto L82
L85:
	;
	v526 = v517 - int32(1)
	if v526 == int32(0) {
		v536 = v516
		v538 = v518
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v529 = int32(1)
	v530 = v516 + v529
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	if v531 != 0 {
		v515 = v515 + v529
		v516 = v530
		v517 = v526
		v518 = v531
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	goto L74
L89:
	;
	v600 = v567 + (v568-v567)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v601 = int32(*(*int8)(unsafe.Add(mBase, uint32(v600))))
	v602 = v556 - v601
	if v602 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L70
L91:
	;
	goto L96
L92:
	;
	v653 = v602
	goto L93
L93:
	;
	v657 = base.B2i32(v653 < int32(0))
	if v653 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L94:
	;
	if v644 == int32(0) {
		v670 = v600
		goto L71
	} else {
		goto L107
	}
L96:
	;
	goto L97
L97:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v611 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v612 = v506
	v613 = v600
	v614 = int32(10)
	v615 = v611
	goto L102
L99:
	;
	v638 = v600
	v642 = int32(0)
	goto L100
L100:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	v644 = v642 - v643
	goto L94
L101:
	;
	v638 = v633
	v642 = v635
	goto L100
L102:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	if base.B2i32(v615 != v617)|base.B2i32(v617 == int32(0)) != 0 {
		v633 = v613
		v635 = v615
		goto L101
	} else {
		goto L104
	}
L103:
	;
	v633 = v627
	v635 = int32(0)
	goto L101
L104:
	;
	v623 = v614 - int32(1)
	if v623 == int32(0) {
		v633 = v613
		v635 = v615
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v626 = int32(1)
	v627 = v613 + v626
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	if v628 != 0 {
		v612 = v612 + v626
		v613 = v627
		v614 = v623
		v615 = v628
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v653 = v644
	goto L93
L108:
	;
	v658 = v600 - int32(16)
	goto L110
L109:
	;
	v658 = v568
	goto L110
L110:
	;
	if v653 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v661 = v567
	goto L113
L112:
	;
	v661 = v600 + int32(16)
	goto L113
L113:
	;
	if base.Ui32(v661) <= base.Ui32(v658) {
		v567 = v661
		v568 = v658
		goto L89
	} else {
		goto L114
	}
L114:
	;
	goto L90
L115:
	;
	goto L70
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+uint32(_c_F_interval_in[3]))) = v905
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905)+11)))
	v945 = v905
	v952 = v937
	goto L69
L117:
	;
	goto L122
L118:
	;
	goto L119
L119:
	;
	v790 = int32(*(*int8)(unsafe.Add(mBase, uint32(v506))))
	v801 = int32(_a_F_interval_in_3)
	v802 = int32(_a_F_interval_in_4)
	goto L134
L120:
	;
	if v779-v780 == int32(0) {
		v905 = v741
		goto L116
	} else {
		goto L133
	}
L122:
	;
	goto L123
L123:
	;
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v748 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v749 = v506
	v750 = v741
	v751 = int32(10)
	v752 = v748
	goto L128
L125:
	;
	v775 = v741
	v779 = int32(0)
	goto L126
L126:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775))))
	goto L120
L127:
	;
	v775 = v770
	v779 = v772
	goto L126
L128:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	if base.B2i32(v752 != v754)|base.B2i32(v754 == int32(0)) != 0 {
		v770 = v750
		v772 = v752
		goto L127
	} else {
		goto L130
	}
L129:
	;
	v770 = v764
	v772 = int32(0)
	goto L127
L130:
	;
	v760 = v751 - int32(1)
	if v760 == int32(0) {
		v770 = v750
		v772 = v752
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v763 = int32(1)
	v764 = v750 + v763
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+1)))
	if v765 != 0 {
		v749 = v749 + v763
		v750 = v764
		v751 = v760
		v752 = v765
		goto L128
	} else {
		goto L132
	}
L132:
	;
	goto L129
L133:
	;
	goto L119
L134:
	;
	v834 = v801 + (v802-v801)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v835 = int32(*(*int8)(unsafe.Add(mBase, uint32(v834))))
	v836 = v790 - v835
	if v836 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v2108 = int32(-1)
	goto L7
L136:
	;
	goto L141
L137:
	;
	v887 = v836
	goto L138
L138:
	;
	v891 = base.B2i32(v887 < int32(0))
	if v887 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L139:
	;
	if v878 == int32(0) {
		v905 = v834
		goto L116
	} else {
		goto L152
	}
L141:
	;
	goto L142
L142:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v845 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v846 = v506
	v847 = v834
	v848 = int32(10)
	v849 = v845
	goto L147
L144:
	;
	v872 = v834
	v876 = int32(0)
	goto L145
L145:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v872))))
	v878 = v876 - v877
	goto L139
L146:
	;
	v872 = v867
	v876 = v869
	goto L145
L147:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	if base.B2i32(v849 != v851)|base.B2i32(v851 == int32(0)) != 0 {
		v867 = v847
		v869 = v849
		goto L146
	} else {
		goto L149
	}
L148:
	;
	v867 = v861
	v869 = int32(0)
	goto L146
L149:
	;
	v857 = v848 - int32(1)
	if v857 == int32(0) {
		v867 = v847
		v869 = v849
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v860 = int32(1)
	v861 = v847 + v860
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846)+1)))
	if v862 != 0 {
		v846 = v846 + v860
		v847 = v861
		v848 = v857
		v849 = v862
		goto L147
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v887 = v878
	goto L138
L153:
	;
	v892 = v834 - int32(16)
	goto L155
L154:
	;
	v892 = v802
	goto L155
L155:
	;
	if v887 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v895 = v801
	goto L158
L157:
	;
	v895 = v834 + int32(16)
	goto L158
L158:
	;
	if base.Ui32(v895) <= base.Ui32(v892) {
		v801 = v895
		v802 = v892
		goto L134
	} else {
		goto L159
	}
L159:
	;
	goto L135
L160:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v945)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(0)
	v984 = int32(-1)
	switch v977 {
	case 0:
		goto L161
	default:
		v2108 = v984
		goto L7
	case 17:
		v1995 = v980
		v1996 = int32(1)
		v2020 = v272
		goto L22
	case 19:
		goto L162
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(_a_F_interval_in_5)
	if base.B2i32(v204 != v252)|base.B2i32(base.Ui32(int32(1)) < base.Ui32(v980-int32(9))) != 0 {
		v2108 = v984
		goto L7
	} else {
		goto L164
	}
L162:
	;
	if v204 != v252 {
		v2108 = v984
		goto L7
	} else {
		goto L163
	}
L163:
	;
	v1995 = v980
	v1996 = int32(0)
	v2020 = int32(1)
	goto L22
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v980
	v997 = int32(0)
	v1995 = v997
	v1996 = v997
	v2020 = v272
	goto L22
L165:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v1012 == int32(68) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v2108 = int32(-2)
	goto L7
L167:
	;
	goto L168
L168:
	;
	v1016 = float64(0)
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v71)+116))
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1017))))
	switch v1018 - int32(45) {
	case 0:
		goto L171
	case 1:
		goto L170
	default:
		goto L172
	}
L169:
	;
	if v213 == int32(0) {
		v1238 = v1221
		v1240 = v1226
		goto L221
	} else {
		goto L222
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+120)) = v1017
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1017)+1)))
	if v1107 == int32(0) {
		v1214 = v1016
		goto L192
	} else {
		goto L193
	}
L171:
	;
	v1029 = F_strtol(m, v1017+int32(1), v71+int32(116), int32(10))
	mBase = m.M
	goto L174
L172:
	;
	if v1018 == int32(0) {
		v1221 = v1010
		v1224 = v1000
		v1226 = v1016
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v2108 = int32(-1)
	goto L7
L174:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v1031 == int32(68) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v2108 = int32(-2)
	goto L7
L176:
	;
	goto L177
L177:
	;
	if base.Ui32(int32(11)) < base.Ui32(v1029) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v2108 = int32(-2)
	goto L7
L179:
	;
	goto L180
L180:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v71)+116))
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1038))))
	if v1039 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v2108 = int32(-1)
	goto L7
L182:
	;
	goto L183
L183:
	;
	v1042 = v71 + int32(80)
	v1045 = int64(12)
	v1046 = int64(0)
	v1051 = int64(32)
	v1054 = int64(base.Ui64(v1010) >> (uint(v1051) % 64))
	v1057 = int64(4294967295)
	v1060 = v1010 & v1057
	v1061 = v1045 * v1060
	v1065 = int64(base.Ui64(v1061)>>(uint(v1051)%64)) + v1045*v1054
	v1072 = v1060*v1046 + v1065&v1057
	*(*int64)(unsafe.Add(mBase, uint32(v1042)+8)) = v1010*v1046 + v1010>>(uint(int64(63))%64)*v1045 + v1046*v1054 + int64(base.Ui64(v1065)>>(uint(v1051)%64)) + int64(base.Ui64(v1072)>>(uint(v1051)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1042))) = v1061&v1057 | v1072<<(uint(v1051)%64)
	goto L184
L184:
	;
	v1083 = *(*int64)(unsafe.Add(mBase, uint32(v71)+88))
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(v71)+80))
	if v1083 != v1084>>(uint(int64(63))%64) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v2108 = int32(-2)
	goto L7
L186:
	;
	goto L187
L187:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092))))
	if v1093 == int32(45) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1096 = int32(0) - v1029
	goto L190
L189:
	;
	v1096 = v1029
	goto L190
L190:
	;
	v1097 = base.I64_extend_i32_s(v1096)
	v1100 = v1084 + v1097
	if base.B2i32(v1097 < int64(0))^base.B2i32(v1100 < v1084) == int32(0) {
		v1221 = v1100
		v1224 = int32(23)
		v1226 = v1016
		goto L169
	} else {
		goto L191
	}
L191:
	;
	v2108 = int32(-2)
	goto L7
L192:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1216))))
	if v1217 == int32(45) {
		goto L218
	} else {
		goto L219
	}
L193:
	;
	v1111 = v1017 + int32(1)
	v1112 = int32(_a_F_interval_in_6)
	v1116 = m.G0
	v1118 = v1116 - int32(32)
	v1119 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1118)+24)) = v1119
	*(*int64)(unsafe.Add(mBase, uint32(v1118)+16)) = v1119
	*(*int64)(unsafe.Add(mBase, uint32(v1118)+8)) = v1119
	*(*int64)(unsafe.Add(mBase, uint32(v1118))) = v1119
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[4])))
	if v1127 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v2108 = int32(-1)
	goto L7
L195:
	;
	v1196 = F_strlen(m, v1111)
	mBase = m.M
	if v1195 != v1196 {
		goto L194
	} else {
		goto L214
	}
L196:
	;
	v1195 = int32(0)
	goto L195
L197:
	;
	goto L198
L198:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[5])))
	if v1131 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1135 = v1111
	goto L202
L200:
	;
	goto L201
L201:
	;
	v1145 = v1112
	v1146 = v1127
	goto L205
L202:
	;
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135))))
	if v1141 == v1127 {
		v1135 = v1135 + int32(1)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v1195 = v1135 - v1111
	goto L195
L204:
	;
	goto L203
L205:
	;
	v1153 = v1118 + int32(base.Ui32(v1146)>>(uint(int32(3))%32))&int32(28)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1153)))
	v1155 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1153))) = v1154 | v1155<<(uint(v1146)%32)
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+1)))
	if v1159 != 0 {
		v1145 = v1145 + v1155
		v1146 = v1159
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	if v1162 == int32(0) {
		v1185 = v1111
		goto L208
	} else {
		goto L209
	}
L207:
	;
	goto L206
L208:
	;
	v1195 = v1185 - v1111
	goto L195
L209:
	;
	v1166 = v1111
	v1167 = v1162
	goto L210
L210:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1118+int32(base.Ui32(v1167)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v1175)>>(uint(v1167)%32))&int32(1) == int32(0) {
		v1185 = v1166
		goto L208
	} else {
		goto L212
	}
L211:
	;
	v1185 = v1183
	goto L208
L212:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1166)+1)))
	v1183 = v1166 + int32(1)
	if v1181 != 0 {
		v1166 = v1183
		v1167 = v1181
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_interval_in[1])) = int32(0)
	v1203 = F_strtod(m, v1017, v71+int32(120))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L29
	} else {
		goto L215
	}
L215:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v71)+120))
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205))))
	if v1206 != 0 {
		goto L194
	} else {
		goto L216
	}
L216:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v1208 == int32(0) {
		v1214 = v1203
		goto L192
	} else {
		goto L217
	}
L217:
	;
	goto L194
L218:
	;
	v1220 = base.F64_neg(v1214)
	goto L220
L219:
	;
	v1220 = v1214
	goto L220
L220:
	;
	v1221 = v1010
	v1224 = v1000
	v1226 = v1220
	goto L169
L221:
	;
	switch v1224 - int32(18) {
	case 0:
		goto L234
	case 1:
		goto L233
	case 2:
		goto L232
	case 3:
		goto L231
	case 4:
		goto L230
	case 5:
		goto L229
	default:
		v2108 = int32(-1)
		goto L7
	case 7:
		goto L228
	case 8:
		goto L227
	case 9:
		goto L226
	case 10:
		goto L225
	case 11:
		goto L235
	case 12:
		goto L236
	}
L222:
	;
	v1230 = v1221 >> (uint(int64(63)) % 64)
	v1232 = v1230 - (v1221 ^ v1230)
	if base.F64_gt(v1226, float64(0)) == int32(0) {
		v1238 = v1232
		v1240 = v1226
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1238 = v1232
	v1240 = base.F64_neg(v1226)
	goto L221
L224:
	;
	v1995 = int32(21)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L225:
	;
	if base.Ui64(v1238-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L395
	} else {
		goto L396
	}
L226:
	;
	if base.Ui64(v1238-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L385
	} else {
		goto L386
	}
L227:
	;
	if base.Ui64(v1238-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L375
	} else {
		goto L376
	}
L228:
	;
	if base.Ui64(v1238-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L368
	} else {
		goto L369
	}
L229:
	;
	if base.Ui64(v1238-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L350
	} else {
		goto L351
	}
L230:
	;
	if base.Ui64(v1238-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L329
	} else {
		goto L330
	}
L231:
	;
	if base.Ui64(v1238-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L314
	} else {
		goto L315
	}
L232:
	;
	v1553 = v71 - int32(-64)
	v1556 = int64(3600000000)
	v1557 = int64(0)
	v1562 = int64(32)
	v1565 = int64(base.Ui64(v1238) >> (uint(v1562) % 64))
	v1568 = int64(4294967295)
	v1571 = v1238 & v1568
	v1572 = v1556 * v1571
	v1576 = int64(base.Ui64(v1572)>>(uint(v1562)%64)) + v1556*v1565
	v1583 = v1571*v1557 + v1576&v1568
	*(*int64)(unsafe.Add(mBase, uint32(v1553)+8)) = v1238*v1557 + v1238>>(uint(int64(63))%64)*v1556 + v1557*v1565 + int64(base.Ui64(v1576)>>(uint(v1562)%64)) + int64(base.Ui64(v1583)>>(uint(v1562)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1553))) = v1572&v1568 | v1583<<(uint(v1562)%64)
	goto L298
L233:
	;
	v1464 = v71 + int32(48)
	v1467 = int64(60000000)
	v1468 = int64(0)
	v1473 = int64(32)
	v1476 = int64(base.Ui64(v1238) >> (uint(v1473) % 64))
	v1479 = int64(4294967295)
	v1482 = v1238 & v1479
	v1483 = v1467 * v1482
	v1487 = int64(base.Ui64(v1483)>>(uint(v1473)%64)) + v1467*v1476
	v1494 = v1482*v1468 + v1487&v1479
	*(*int64)(unsafe.Add(mBase, uint32(v1464)+8)) = v1238*v1468 + v1238>>(uint(int64(63))%64)*v1467 + v1468*v1476 + int64(base.Ui64(v1487)>>(uint(v1473)%64)) + int64(base.Ui64(v1494)>>(uint(v1473)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1464))) = v1483&v1479 | v1494<<(uint(v1473)%64)
	goto L282
L234:
	;
	v1373 = v71 + int32(32)
	v1376 = int64(1000000)
	v1377 = int64(0)
	v1382 = int64(32)
	v1385 = int64(base.Ui64(v1238) >> (uint(v1382) % 64))
	v1388 = int64(4294967295)
	v1391 = v1238 & v1388
	v1392 = v1376 * v1391
	v1396 = int64(base.Ui64(v1392)>>(uint(v1382)%64)) + v1376*v1385
	v1403 = v1391*v1377 + v1396&v1388
	*(*int64)(unsafe.Add(mBase, uint32(v1373)+8)) = v1238*v1377 + v1238>>(uint(int64(63))%64)*v1376 + v1377*v1385 + int64(base.Ui64(v1396)>>(uint(v1382)%64)) + int64(base.Ui64(v1403)>>(uint(v1382)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1373))) = v1392&v1388 | v1403<<(uint(v1382)%64)
	goto L265
L235:
	;
	v1284 = v71 + int32(16)
	v1287 = int64(1000)
	v1288 = int64(0)
	v1293 = int64(32)
	v1296 = int64(base.Ui64(v1238) >> (uint(v1293) % 64))
	v1299 = int64(4294967295)
	v1302 = v1238 & v1299
	v1303 = v1287 * v1302
	v1307 = int64(base.Ui64(v1303)>>(uint(v1293)%64)) + v1287*v1296
	v1314 = v1302*v1288 + v1307&v1299
	*(*int64)(unsafe.Add(mBase, uint32(v1284)+8)) = v1238*v1288 + v1238>>(uint(int64(63))%64)*v1287 + v1288*v1296 + int64(base.Ui64(v1307)>>(uint(v1293)%64)) + int64(base.Ui64(v1314)>>(uint(v1293)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1284))) = v1303&v1299 | v1314<<(uint(v1293)%64)
	goto L249
L236:
	;
	v1244 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1245 = v1244 + v1238
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1245
	if base.B2i32(v1238 < int64(0))^base.B2i32(v1245 < v1244) != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v2108 = int32(-2)
	goto L7
L238:
	;
	goto L239
L239:
	;
	if base.F64_ne(v1240, float64(0)) != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1254 = base.I64_trunc_sat_f64_s(v1240)
	v1256 = base.F64_sub(v1240, base.F64_convert_i64_s(v1254))
	if base.F64_gt(v1256, float64(0.5)) != 0 {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(_a_F_interval_in_7)
	v1995 = int32(30)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L243:
	;
	v1268 = v1267 + v1245
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1268
	if base.B2i32(v1267 < int64(0))^base.B2i32(v1268 < v1245) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L248
	}
L244:
	;
	v1267 = v1254 + int64(1)
	goto L243
L245:
	;
	goto L246
L246:
	;
	if base.F64_lt(v1256, float64(-0.5)) == int32(0) {
		v1267 = v1254
		goto L243
	} else {
		goto L247
	}
L247:
	;
	v1267 = v1254 - int64(1)
	goto L243
L248:
	;
	goto L242
L249:
	;
	v1325 = *(*int64)(unsafe.Add(mBase, uint32(v71)+24))
	v1326 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
	if v1325 != v1326>>(uint(int64(63))%64) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v2108 = int32(-2)
	goto L7
L251:
	;
	goto L252
L252:
	;
	v1331 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1332 = v1331 + v1326
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1332
	if base.B2i32(v1326 < int64(0))^base.B2i32(v1332 < v1331) != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v2108 = int32(-2)
	goto L7
L254:
	;
	goto L255
L255:
	;
	if base.F64_ne(v1240, float64(0)) != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1342 = base.F64_mul(v1240, float64(1000))
	v1343 = base.I64_trunc_sat_f64_s(v1342)
	v1345 = base.F64_sub(v1342, base.F64_convert_i64_s(v1343))
	if base.F64_gt(v1345, float64(0.5)) != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(_a_F_interval_in_8)
	v1995 = int32(29)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L259:
	;
	v1357 = v1356 + v1332
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1357
	if base.B2i32(v1356 < int64(0))^base.B2i32(v1357 < v1332) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L264
	}
L260:
	;
	v1356 = v1343 + int64(1)
	goto L259
L261:
	;
	goto L262
L262:
	;
	if base.F64_lt(v1345, float64(-0.5)) == int32(0) {
		v1356 = v1343
		goto L259
	} else {
		goto L263
	}
L263:
	;
	v1356 = v1343 - int64(1)
	goto L259
L264:
	;
	goto L258
L265:
	;
	v1414 = *(*int64)(unsafe.Add(mBase, uint32(v71)+40))
	v1415 = *(*int64)(unsafe.Add(mBase, uint32(v71)+32))
	if v1414 != v1415>>(uint(int64(63))%64) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v2108 = int32(-2)
	goto L7
L267:
	;
	goto L268
L268:
	;
	v1420 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1421 = v1420 + v1415
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1421
	if base.B2i32(v1415 < int64(0))^base.B2i32(v1421 < v1420) != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v2108 = int32(-2)
	goto L7
L270:
	;
	goto L271
L271:
	;
	if base.F64_ne(v1240, float64(0)) != 0 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1995 = int32(18)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L273:
	;
	v1431 = base.F64_mul(v1240, float64(1e+06))
	v1432 = base.I64_trunc_sat_f64_s(v1431)
	v1434 = base.F64_sub(v1431, base.F64_convert_i64_s(v1432))
	if base.F64_gt(v1434, float64(0.5)) != 0 {
		goto L277
	} else {
		goto L278
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(_a_F_interval_in_9)
	goto L272
L276:
	;
	v1446 = v1445 + v1421
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1446
	if base.B2i32(v1445 < int64(0))^base.B2i32(v1446 < v1421) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L281
	}
L277:
	;
	v1445 = v1432 + int64(1)
	goto L276
L278:
	;
	goto L279
L279:
	;
	if base.F64_lt(v1434, float64(-0.5)) == int32(0) {
		v1445 = v1432
		goto L276
	} else {
		goto L280
	}
L280:
	;
	v1445 = v1432 - int64(1)
	goto L276
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(_a_F_interval_in_10)
	goto L272
L282:
	;
	v1505 = *(*int64)(unsafe.Add(mBase, uint32(v71)+56))
	v1506 = *(*int64)(unsafe.Add(mBase, uint32(v71)+48))
	if v1505 != v1506>>(uint(int64(63))%64) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v2108 = int32(-2)
	goto L7
L284:
	;
	goto L285
L285:
	;
	v1511 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1512 = v1511 + v1506
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1512
	if base.B2i32(v1506 < int64(0))^base.B2i32(v1512 < v1511) != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v2108 = int32(-2)
	goto L7
L287:
	;
	goto L288
L288:
	;
	if base.F64_ne(v1240, float64(0)) != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1522 = base.F64_mul(v1240, float64(6e+07))
	v1523 = base.I64_trunc_sat_f64_s(v1522)
	v1525 = base.F64_sub(v1522, base.F64_convert_i64_s(v1523))
	if base.F64_gt(v1525, float64(0.5)) != 0 {
		goto L293
	} else {
		goto L294
	}
L290:
	;
	goto L291
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(2048)
	v1995 = int32(19)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L292:
	;
	v1537 = v1536 + v1512
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1537
	if base.B2i32(v1536 < int64(0))^base.B2i32(v1537 < v1512) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L297
	}
L293:
	;
	v1536 = v1523 + int64(1)
	goto L292
L294:
	;
	goto L295
L295:
	;
	if base.F64_lt(v1525, float64(-0.5)) == int32(0) {
		v1536 = v1523
		goto L292
	} else {
		goto L296
	}
L296:
	;
	v1536 = v1523 - int64(1)
	goto L292
L297:
	;
	goto L291
L298:
	;
	v1594 = *(*int64)(unsafe.Add(mBase, uint32(v71)+72))
	v1595 = *(*int64)(unsafe.Add(mBase, uint32(v71)+64))
	if v1594 != v1595>>(uint(int64(63))%64) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v2108 = int32(-2)
	goto L7
L300:
	;
	goto L301
L301:
	;
	v1600 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1601 = v1600 + v1595
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1601
	if base.B2i32(v1595 < int64(0))^base.B2i32(v1601 < v1600) != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v2108 = int32(-2)
	goto L7
L303:
	;
	goto L304
L304:
	;
	if base.F64_ne(v1240, float64(0)) != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1611 = base.F64_mul(v1240, float64(3.6e+09))
	v1612 = base.I64_trunc_sat_f64_s(v1611)
	v1614 = base.F64_sub(v1611, base.F64_convert_i64_s(v1612))
	if base.F64_gt(v1614, float64(0.5)) != 0 {
		goto L309
	} else {
		goto L310
	}
L306:
	;
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(1024)
	goto L224
L308:
	;
	v1626 = v1625 + v1601
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1626
	if base.B2i32(v1625 < int64(0))^base.B2i32(v1626 < v1601) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L313
	}
L309:
	;
	v1625 = v1612 + int64(1)
	goto L308
L310:
	;
	goto L311
L311:
	;
	if base.F64_lt(v1614, float64(-0.5)) == int32(0) {
		v1625 = v1612
		goto L308
	} else {
		goto L312
	}
L312:
	;
	v1625 = v1612 - int64(1)
	goto L308
L313:
	;
	goto L307
L314:
	;
	v2108 = int32(-2)
	goto L7
L315:
	;
	goto L316
L316:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v1646 = v1644 + base.I32_wrap_i64(v1238)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v1646
	if base.B2i32(v1646 < v1644)^base.B2i32(v1238 < int64(0)) != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v2108 = int32(-2)
	goto L7
L318:
	;
	goto L319
L319:
	;
	if base.F64_ne(v1240, float64(0)) != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1656 = base.F64_mul(v1240, float64(8.64e+10))
	v1657 = base.I64_trunc_sat_f64_s(v1656)
	v1659 = base.F64_sub(v1656, base.F64_convert_i64_s(v1657))
	if base.F64_gt(v1659, float64(0.5)) != 0 {
		goto L324
	} else {
		goto L325
	}
L321:
	;
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(8)
	goto L224
L323:
	;
	v1671 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1672 = v1671 + v1670
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1672
	if base.B2i32(v1670 < int64(0))^base.B2i32(v1672 < v1671) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L328
	}
L324:
	;
	v1670 = v1657 + int64(1)
	goto L323
L325:
	;
	goto L326
L326:
	;
	if base.F64_lt(v1659, float64(-0.5)) == int32(0) {
		v1670 = v1657
		goto L323
	} else {
		goto L327
	}
L327:
	;
	v1670 = v1657 - int64(1)
	goto L323
L328:
	;
	goto L322
L329:
	;
	v2108 = int32(-2)
	goto L7
L330:
	;
	goto L331
L331:
	;
	v1692 = v1238 * int64(7)
	v1696 = base.I32_wrap_i64(v1692)
	if base.I32_wrap_i64(int64(base.Ui64(v1692)>>(uint(int64(32))%64))) != v1696>>(uint(int32(31))%32) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v2108 = int32(-2)
	goto L7
L333:
	;
	goto L334
L334:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v1702 = v1701 + v1696
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v1702
	if base.B2i32(v1696 < int32(0))^base.B2i32(v1702 < v1701) != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v2108 = int32(-2)
	goto L7
L336:
	;
	goto L337
L337:
	;
	if base.F64_eq(v1240, float64(0)) != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(16777216)
	v1995 = int32(22)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L339:
	;
	v1712 = base.F64_mul(v1240, float64(7))
	v1713 = base.I32_trunc_sat_f64_s(v1712)
	v1714 = v1702 + v1713
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v1714
	if base.B2i32(v1713 < int32(0))^base.B2i32(v1714 < v1702) != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2108 = int32(-2)
	goto L7
L341:
	;
	goto L342
L342:
	;
	v1722 = base.F64_sub(v1712, base.F64_convert_i32_s(v1713))
	if base.F64_eq(v1722, float64(0)) != 0 {
		goto L338
	} else {
		goto L343
	}
L343:
	;
	v1726 = base.F64_mul(v1722, float64(8.64e+10))
	v1727 = base.I64_trunc_sat_f64_s(v1726)
	v1729 = base.F64_sub(v1726, base.F64_convert_i64_s(v1727))
	if base.F64_gt(v1729, float64(0.5)) != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1741 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1742 = v1741 + v1740
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1742
	if base.B2i32(v1740 < int64(0))^base.B2i32(v1742 < v1741) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L349
	}
L345:
	;
	v1740 = v1727 + int64(1)
	goto L344
L346:
	;
	goto L347
L347:
	;
	if base.F64_lt(v1729, float64(-0.5)) == int32(0) {
		v1740 = v1727
		goto L344
	} else {
		goto L348
	}
L348:
	;
	v1740 = v1727 - int64(1)
	goto L344
L349:
	;
	goto L338
L350:
	;
	v2108 = int32(-2)
	goto L7
L351:
	;
	goto L352
L352:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v1767 = v1765 + base.I32_wrap_i64(v1238)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v1767
	if base.B2i32(v1767 < v1765)^base.B2i32(v1238 < int64(0)) != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v2108 = int32(-2)
	goto L7
L354:
	;
	goto L355
L355:
	;
	if base.F64_eq(v1240, float64(0)) != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(2)
	v1995 = int32(23)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L357:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v1778 = base.F64_mul(v1240, float64(30))
	v1779 = base.I32_trunc_sat_f64_s(v1778)
	v1780 = v1776 + v1779
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v1780
	if base.B2i32(v1779 < int32(0))^base.B2i32(v1780 < v1776) != 0 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v2108 = int32(-2)
	goto L7
L359:
	;
	goto L360
L360:
	;
	v1788 = base.F64_sub(v1778, base.F64_convert_i32_s(v1779))
	if base.F64_eq(v1788, float64(0)) != 0 {
		goto L356
	} else {
		goto L361
	}
L361:
	;
	v1792 = base.F64_mul(v1788, float64(8.64e+10))
	v1793 = base.I64_trunc_sat_f64_s(v1792)
	v1795 = base.F64_sub(v1792, base.F64_convert_i64_s(v1793))
	if base.F64_gt(v1795, float64(0.5)) != 0 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v1807 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	v1808 = v1807 + v1806
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v1808
	if base.B2i32(v1806 < int64(0))^base.B2i32(v1808 < v1807) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L367
	}
L363:
	;
	v1806 = v1793 + int64(1)
	goto L362
L364:
	;
	goto L365
L365:
	;
	if base.F64_lt(v1795, float64(-0.5)) == int32(0) {
		v1806 = v1793
		goto L362
	} else {
		goto L366
	}
L366:
	;
	v1806 = v1793 - int64(1)
	goto L362
L367:
	;
	goto L356
L368:
	;
	v2108 = int32(-2)
	goto L7
L369:
	;
	goto L370
L370:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v1833 = v1831 + base.I32_wrap_i64(v1238)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v1833
	if base.B2i32(v1833 < v1831)^base.B2i32(v1238 < int64(0)) != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v2108 = int32(-2)
	goto L7
L372:
	;
	goto L373
L373:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v1844 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v1240, float64(12))))
	v1845 = v1840 + v1844
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v1845
	if base.B2i32(v1844 < int32(0))^base.B2i32(v1845 < v1840) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(4)
	v1995 = int32(25)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L375:
	;
	v2108 = int32(-2)
	goto L7
L376:
	;
	goto L377
L377:
	;
	v1862 = v1238 * int64(10)
	v1866 = base.I32_wrap_i64(v1862)
	if base.I32_wrap_i64(int64(base.Ui64(v1862)>>(uint(int64(32))%64))) != v1866>>(uint(int32(31))%32) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v2108 = int32(-2)
	goto L7
L379:
	;
	goto L380
L380:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v1872 = v1871 + v1866
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v1872
	if base.B2i32(v1866 < int32(0))^base.B2i32(v1872 < v1871) != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v2108 = int32(-2)
	goto L7
L382:
	;
	goto L383
L383:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v1885 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_mul(v1240, float64(10)), float64(12))))
	v1886 = v1879 + v1885
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v1886
	if base.B2i32(v1885 < int32(0))^base.B2i32(v1886 < v1879) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L384
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(33554432)
	v1995 = int32(26)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L385:
	;
	v2108 = int32(-2)
	goto L7
L386:
	;
	goto L387
L387:
	;
	v1903 = v1238 * int64(100)
	v1907 = base.I32_wrap_i64(v1903)
	if base.I32_wrap_i64(int64(base.Ui64(v1903)>>(uint(int64(32))%64))) != v1907>>(uint(int32(31))%32) {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2108 = int32(-2)
	goto L7
L389:
	;
	goto L390
L390:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v1913 = v1912 + v1907
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v1913
	if base.B2i32(v1907 < int32(0))^base.B2i32(v1913 < v1912) != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2108 = int32(-2)
	goto L7
L392:
	;
	goto L393
L393:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v1926 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_mul(v1240, float64(100)), float64(12))))
	v1927 = v1920 + v1926
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v1927
	if base.B2i32(v1926 < int32(0))^base.B2i32(v1927 < v1920) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L394
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(67108864)
	v1995 = int32(27)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L395:
	;
	v2108 = int32(-2)
	goto L7
L396:
	;
	goto L397
L397:
	;
	v1944 = v1238 * int64(1000)
	v1948 = base.I32_wrap_i64(v1944)
	if base.I32_wrap_i64(int64(base.Ui64(v1944)>>(uint(int64(32))%64))) != v1948>>(uint(int32(31))%32) {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v2108 = int32(-2)
	goto L7
L399:
	;
	goto L400
L400:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v1954 = v1953 + v1948
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v1954
	if base.B2i32(v1948 < int32(0))^base.B2i32(v1954 < v1953) != 0 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v2108 = int32(-2)
	goto L7
L402:
	;
	goto L403
L403:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v1967 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_mul(v1240, float64(1000)), float64(12))))
	v1968 = v1961 + v1967
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v1968
	if base.B2i32(v1967 < int32(0))^base.B2i32(v1968 < v1961) != 0 {
		v2108 = int32(-2)
		goto L7
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+112)) = int32(134217728)
	v1995 = int32(28)
	v1996 = int32(0)
	v2020 = v272
	goto L22
L405:
	;
	v2108 = int32(-1)
	goto L7
L406:
	;
	goto L407
L407:
	;
	v2035 = v1995
	v2036 = v1996
	v2052 = v2023 | v264
	v2060 = v2020
	goto L21
L408:
	;
	goto L20
L409:
	;
	v2075 = int32(-2)
	v2076 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	if v2076 == int64(-9223372036854775807-1) {
		v2108 = v2075
		goto L7
	} else {
		goto L410
	}
L410:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v2079 == int32(-2147483648) {
		v2108 = v2075
		goto L7
	} else {
		goto L411
	}
L411:
	;
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v2082 == int32(-2147483648) {
		v2108 = v2075
		goto L7
	} else {
		goto L412
	}
L412:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v2085 == int32(-2147483648) {
		v2108 = v2075
		goto L7
	} else {
		goto L413
	}
L413:
	;
	v2088 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v2088 - v2085
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v2088 - v2082
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v2088 - v2079
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = int64(0) - v2076
	v2108 = v2088
	goto L7
L414:
	;
	v2178 = int32(0)
	v2179 = m.G0
	v2181 = v2179 - int32(112)
	m.G0 = v2181
	*(*int32)(unsafe.Add(mBase, uint32(v39+int32(500)))) = int32(17)
	v2188 = v39 + int32(504)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+16)) = v2178
	v2191 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2188)+8)) = v2191
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2191
	v2195 = int32(-1)
	v2196 = F_strlen(m, v43)
	mBase = m.M
	if base.Ui32(v2196) < base.Ui32(int32(2)) {
		v3624 = v2195
		goto L417
	} else {
		goto L418
	}
L415:
	;
	v3669 = v2153
	goto L416
L416:
	;
	if v3669 != 0 {
		goto L753
	} else {
		goto L754
	}
L417:
	;
	m.G0 = v2181 + int32(112)
	v3669 = v3624
	goto L416
L418:
	;
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v2199 != int32(80) {
		v3624 = v2195
		goto L417
	} else {
		goto L419
	}
L419:
	;
	v2203 = v43 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+108)) = v2203
	v2205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v2205 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	v3624 = v3610
	goto L417
L421:
	;
	v3624 = int32(0)
	goto L417
L422:
	;
	v2220 = v2203
	v2222 = v2205
	v2223 = int32(1)
	v2225 = v2178
	goto L423
L423:
	;
	if v2222&int32(255) == int32(84) {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	goto L421
L425:
	;
	v3567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3559))))
	if v3567 != 0 {
		v2220 = v3559
		v2222 = v3567
		v2223 = v3562
		v2225 = v3563
		goto L423
	} else {
		goto L751
	}
L426:
	;
	v2250 = v2220 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+108)) = v2250
	v2252 = int32(0)
	v3559 = v2250
	v3562 = v2252
	v3563 = v2252
	goto L425
L427:
	;
	goto L428
L428:
	;
	v2254 = int32(-1)
	v2257 = int32(255)
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32((v2222-int32(45))&v2257))&base.B2i32(base.Ui32(int32(10)) <= base.Ui32((v2222-int32(48))&v2257)) != 0 {
		v3610 = v2254
		goto L420
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_interval_in[1])) = int32(0)
	v2273 = F_strtod(m, v2220, v2181+int32(108))
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L29
	} else {
		goto L430
	}
L430:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2181)+108))
	if v2275 == v2220 {
		v3610 = v2254
		goto L420
	} else {
		goto L431
	}
L431:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v2278 != 0 {
		v3610 = v2254
		goto L420
	} else {
		goto L432
	}
L432:
	;
	v2279 = base.F64_abs(v2273)
	if base.F64_gt(v2279, float64(1e+15)) != 0 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v3624 = int32(-2)
	goto L417
L434:
	;
	goto L435
L435:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v2279)) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v3624 = int32(-2)
	goto L417
L437:
	;
	goto L438
L438:
	;
	v2288 = v2275 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+108)) = v2288
	if base.F64_ge(v2273, float64(0)) != 0 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v2296 = base.F64_floor(v2273)
	goto L441
L440:
	;
	v2296 = base.F64_neg(base.F64_floor(base.F64_neg(v2273)))
	goto L441
L441:
	;
	v2297 = base.I64_trunc_sat_f64_s(v2296)
	*(*int64)(unsafe.Add(mBase, uint32(v2181)+96)) = v2297
	v2300 = base.F64_sub(v2273, base.F64_convert_i64_s(v2297))
	*(*float64)(unsafe.Add(mBase, uint32(v2181)+88)) = v2300
	v2302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275))))
	if v2223&int32(1) != 0 {
		goto L444
	} else {
		goto L445
	}
L442:
	;
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+12))
	v3545 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v2300, float64(12))))
	v3546 = v3541 + v3545
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+12)) = v3546
	v3548 = int32(1)
	if base.B2i32(v3545 < int32(0))^base.B2i32(v3546 < v3541) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L750
	}
L443:
	;
	v3610 = int32(0)
	goto L420
L444:
	;
	switch v2302 - int32(45) {
	case 0:
		goto L447
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 40, 41, 43:
		v3610 = v2254
		goto L420
	case 23:
		goto L449
	case 32:
		goto L451
	case 39:
		goto L453
	case 42:
		goto L450
	case 44:
		goto L452
	default:
		goto L454
	}
L445:
	;
	goto L446
L446:
	;
	switch v2302 - int32(58) {
	case 0:
		goto L618
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 20, 21, 22, 23, 24:
		v3610 = v2254
		goto L420
	case 14:
		goto L622
	case 19:
		goto L621
	case 25:
		goto L620
	default:
		goto L619
	}
L447:
	;
	if v2225 != 0 {
		v3610 = v2254
		goto L420
	} else {
		goto L560
	}
L448:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+16))
	v2592 = v2590 + base.I32_wrap_i64(v2399)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+16)) = v2592
	if base.B2i32(v2592 < v2590)^base.B2i32(v2399 < int64(0)) != 0 {
		goto L541
	} else {
		goto L542
	}
L449:
	;
	if base.Ui64(v2297-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L526
	} else {
		goto L527
	}
L450:
	;
	if base.Ui64(v2297-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L502
	} else {
		goto L503
	}
L451:
	;
	if base.Ui64(v2297-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L481
	} else {
		goto L482
	}
L452:
	;
	if base.Ui64(v2297-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L477
	} else {
		goto L478
	}
L453:
	;
	v2307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2220))))
	v2310 = v2220 + base.B2i32(v2307 == int32(45))
	v2311 = int32(_a_F_interval_in_6)
	v2315 = m.G0
	v2317 = v2315 - int32(32)
	v2318 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2317)+24)) = v2318
	*(*int64)(unsafe.Add(mBase, uint32(v2317)+16)) = v2318
	*(*int64)(unsafe.Add(mBase, uint32(v2317)+8)) = v2318
	*(*int64)(unsafe.Add(mBase, uint32(v2317))) = v2318
	v2326 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[4])))
	if v2326 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L454:
	;
	if v2302 != 0 {
		v3610 = v2254
		goto L420
	} else {
		goto L455
	}
L455:
	;
	goto L453
L456:
	;
	if base.B2i32(v2394 != int32(8))|v2225 != 0 {
		goto L447
	} else {
		goto L475
	}
L457:
	;
	v2394 = int32(0)
	goto L456
L458:
	;
	goto L459
L459:
	;
	v2330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[5])))
	if v2330 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v2334 = v2310
	goto L463
L461:
	;
	goto L462
L462:
	;
	v2344 = v2311
	v2345 = v2326
	goto L466
L463:
	;
	v2340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2334))))
	if v2340 == v2326 {
		v2334 = v2334 + int32(1)
		goto L463
	} else {
		goto L465
	}
L464:
	;
	v2394 = v2334 - v2310
	goto L456
L465:
	;
	goto L464
L466:
	;
	v2352 = v2317 + int32(base.Ui32(v2345)>>(uint(int32(3))%32))&int32(28)
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v2352)))
	v2354 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2352))) = v2353 | v2354<<(uint(v2345)%32)
	v2358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2344)+1)))
	if v2358 != 0 {
		v2344 = v2344 + v2354
		v2345 = v2358
		goto L466
	} else {
		goto L468
	}
L467:
	;
	v2361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2310))))
	if v2361 == int32(0) {
		v2384 = v2310
		goto L469
	} else {
		goto L470
	}
L468:
	;
	goto L467
L469:
	;
	v2394 = v2384 - v2310
	goto L456
L470:
	;
	v2365 = v2310
	v2366 = v2361
	goto L471
L471:
	;
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(v2317+int32(base.Ui32(v2366)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v2374)>>(uint(v2366)%32))&int32(1) == int32(0) {
		v2384 = v2365
		goto L469
	} else {
		goto L473
	}
L472:
	;
	v2384 = v2382
	goto L469
L473:
	;
	v2380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2365)+1)))
	v2382 = v2365 + int32(1)
	if v2380 != 0 {
		v2365 = v2382
		v2366 = v2380
		goto L471
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	v2399 = base.I64_div_s(v2297, int64(10000))
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v2399-int64(2147483648)) {
		goto L448
	} else {
		goto L476
	}
L476:
	;
	v3624 = int32(-2)
	goto L417
L477:
	;
	v3624 = int32(-2)
	goto L417
L478:
	;
	goto L479
L479:
	;
	v2410 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+16))
	v2412 = v2410 + base.I32_wrap_i64(v2297)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+16)) = v2412
	if base.B2i32(v2412 < v2410)^base.B2i32(v2297 < int64(0)) == int32(0) {
		goto L442
	} else {
		goto L480
	}
L480:
	;
	v3624 = int32(-2)
	goto L417
L481:
	;
	v3624 = int32(-2)
	goto L417
L482:
	;
	goto L483
L483:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+12))
	v2428 = v2426 + base.I32_wrap_i64(v2297)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+12)) = v2428
	if base.B2i32(v2428 < v2426)^base.B2i32(v2297 < int64(0)) != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v3624 = int32(-2)
	goto L417
L485:
	;
	goto L486
L486:
	;
	v2435 = int32(1)
	if base.F64_eq(v2300, float64(0)) != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2435
	goto L425
L488:
	;
	goto L489
L489:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+8))
	v2441 = base.F64_mul(v2300, float64(30))
	v2442 = base.I32_trunc_sat_f64_s(v2441)
	v2443 = v2439 + v2442
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+8)) = v2443
	if base.B2i32(v2442 < int32(0))^base.B2i32(v2443 < v2439) != 0 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v3624 = int32(-2)
	goto L417
L491:
	;
	goto L492
L492:
	;
	v2451 = base.F64_sub(v2441, base.F64_convert_i32_s(v2442))
	if base.F64_eq(v2451, float64(0)) != 0 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2435
	goto L425
L494:
	;
	goto L495
L495:
	;
	v2456 = base.F64_mul(v2451, float64(8.64e+10))
	v2457 = base.I64_trunc_sat_f64_s(v2456)
	v2459 = base.F64_sub(v2456, base.F64_convert_i64_s(v2457))
	if base.F64_gt(v2459, float64(0.5)) != 0 {
		goto L497
	} else {
		goto L498
	}
L496:
	;
	v2471 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2472 = v2471 + v2470
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2472
	if base.B2i32(v2470 < int64(0))^base.B2i32(v2472 < v2471) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L501
	}
L497:
	;
	v2470 = v2457 + int64(1)
	goto L496
L498:
	;
	goto L499
L499:
	;
	if base.F64_lt(v2459, float64(-0.5)) == int32(0) {
		v2470 = v2457
		goto L496
	} else {
		goto L500
	}
L500:
	;
	v2470 = v2457 - int64(1)
	goto L496
L501:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2435
	goto L425
L502:
	;
	v3624 = int32(-2)
	goto L417
L503:
	;
	goto L504
L504:
	;
	v2486 = v2297 * int64(7)
	v2490 = base.I32_wrap_i64(v2486)
	if base.I32_wrap_i64(int64(base.Ui64(v2486)>>(uint(int64(32))%64))) != v2490>>(uint(int32(31))%32) {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v3624 = int32(-2)
	goto L417
L506:
	;
	goto L507
L507:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+8))
	v2496 = v2495 + v2490
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+8)) = v2496
	if base.B2i32(v2490 < int32(0))^base.B2i32(v2496 < v2495) != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3624 = int32(-2)
	goto L417
L509:
	;
	goto L510
L510:
	;
	v2503 = int32(1)
	if base.F64_eq(v2300, float64(0)) != 0 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2503
	goto L425
L512:
	;
	goto L513
L513:
	;
	v2508 = base.F64_mul(v2300, float64(7))
	v2509 = base.I32_trunc_sat_f64_s(v2508)
	v2510 = v2496 + v2509
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+8)) = v2510
	if base.B2i32(v2509 < int32(0))^base.B2i32(v2510 < v2496) != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v3624 = int32(-2)
	goto L417
L515:
	;
	goto L516
L516:
	;
	v2518 = base.F64_sub(v2508, base.F64_convert_i32_s(v2509))
	if base.F64_eq(v2518, float64(0)) != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2503
	goto L425
L518:
	;
	goto L519
L519:
	;
	v2523 = base.F64_mul(v2518, float64(8.64e+10))
	v2524 = base.I64_trunc_sat_f64_s(v2523)
	v2526 = base.F64_sub(v2523, base.F64_convert_i64_s(v2524))
	if base.F64_gt(v2526, float64(0.5)) != 0 {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	v2538 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2539 = v2538 + v2537
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2539
	if base.B2i32(v2537 < int64(0))^base.B2i32(v2539 < v2538) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L525
	}
L521:
	;
	v2537 = v2524 + int64(1)
	goto L520
L522:
	;
	goto L523
L523:
	;
	if base.F64_lt(v2526, float64(-0.5)) == int32(0) {
		v2537 = v2524
		goto L520
	} else {
		goto L524
	}
L524:
	;
	v2537 = v2524 - int64(1)
	goto L520
L525:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2503
	goto L425
L526:
	;
	v3624 = int32(-2)
	goto L417
L527:
	;
	goto L528
L528:
	;
	v2552 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+8))
	v2554 = v2552 + base.I32_wrap_i64(v2297)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+8)) = v2554
	if base.B2i32(v2554 < v2552)^base.B2i32(v2297 < int64(0)) != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v3624 = int32(-2)
	goto L417
L530:
	;
	goto L531
L531:
	;
	v2561 = int32(1)
	if base.F64_eq(v2300, float64(0)) != 0 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2561
	goto L425
L533:
	;
	goto L534
L534:
	;
	v2566 = base.F64_mul(v2300, float64(8.64e+10))
	v2567 = base.I64_trunc_sat_f64_s(v2566)
	v2569 = base.F64_sub(v2566, base.F64_convert_i64_s(v2567))
	if base.F64_gt(v2569, float64(0.5)) != 0 {
		goto L536
	} else {
		goto L537
	}
L535:
	;
	v2581 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2582 = v2581 + v2580
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2582
	if base.B2i32(v2580 < int64(0))^base.B2i32(v2582 < v2581) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L540
	}
L536:
	;
	v2580 = v2567 + int64(1)
	goto L535
L537:
	;
	goto L538
L538:
	;
	if base.F64_lt(v2569, float64(-0.5)) == int32(0) {
		v2580 = v2567
		goto L535
	} else {
		goto L539
	}
L539:
	;
	v2580 = v2567 - int64(1)
	goto L535
L540:
	;
	v3559 = v2288
	v3562 = int32(1)
	v3563 = v2561
	goto L425
L541:
	;
	v3624 = int32(-2)
	goto L417
L542:
	;
	goto L543
L543:
	;
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+12))
	v2600 = int64(100)
	v2601 = base.I64_div_s(v2297, v2600)
	v2603 = base.I64_rem_s(v2601, v2600)
	v2605 = v2599 + base.I32_wrap_i64(v2603)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+12)) = v2605
	if base.B2i32(v2605 < v2599)^base.B2i32(v2603 < int64(0)) != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v3624 = int32(-2)
	goto L417
L545:
	;
	goto L546
L546:
	;
	v2612 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+8))
	v2614 = base.I64_rem_s(v2297, int64(100))
	v2616 = v2612 + base.I32_wrap_i64(v2614)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+8)) = v2616
	if base.B2i32(v2616 < v2612)^base.B2i32(v2614 < int64(0)) != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v3624 = int32(-2)
	goto L417
L548:
	;
	goto L549
L549:
	;
	if base.F64_ne(v2300, float64(0)) != 0 {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	v2626 = base.F64_mul(v2300, float64(8.64e+10))
	v2627 = base.I64_trunc_sat_f64_s(v2626)
	v2629 = base.F64_sub(v2626, base.F64_convert_i64_s(v2627))
	if base.F64_gt(v2629, float64(0.5)) != 0 {
		goto L554
	} else {
		goto L555
	}
L551:
	;
	goto L552
L552:
	;
	v2654 = int32(0)
	if v2302 == v2654 {
		goto L443
	} else {
		goto L559
	}
L553:
	;
	v2641 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2642 = v2641 + v2640
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2642
	if base.B2i32(v2640 < int64(0))^base.B2i32(v2642 < v2641) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L558
	}
L554:
	;
	v2640 = v2627 + int64(1)
	goto L553
L555:
	;
	goto L556
L556:
	;
	if base.F64_lt(v2629, float64(-0.5)) == int32(0) {
		v2640 = v2627
		goto L553
	} else {
		goto L557
	}
L557:
	;
	v2640 = v2627 - int64(1)
	goto L553
L558:
	;
	goto L552
L559:
	;
	v3559 = v2288
	v3562 = int32(0)
	v3563 = v2654
	goto L425
L560:
	;
	if base.Ui64(v2297-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v3624 = int32(-2)
	goto L417
L562:
	;
	goto L563
L563:
	;
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+16))
	v2665 = v2663 + base.I32_wrap_i64(v2297)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+16)) = v2665
	if base.B2i32(v2665 < v2663)^base.B2i32(v2297 < int64(0)) != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v3624 = int32(-2)
	goto L417
L565:
	;
	goto L566
L566:
	;
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+12))
	v2676 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v2300, float64(12))))
	v2677 = v2672 + v2676
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+12)) = v2677
	if base.B2i32(v2676 < int32(0))^base.B2i32(v2677 < v2672) != 0 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v3624 = int32(-2)
	goto L417
L568:
	;
	goto L569
L569:
	;
	v2684 = int32(0)
	if v2302 == int32(84) {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v3559 = v2288
	v3562 = int32(0)
	v3563 = v2684
	goto L425
L571:
	;
	goto L572
L572:
	;
	if v2302 == int32(0) {
		v3610 = v2302
		goto L420
	} else {
		goto L573
	}
L573:
	;
	v2696 = F_ParseISO8601Number(m, v2288, v2181+int32(108), v2181+int32(96), v2181+int32(88))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L29
	} else {
		goto L574
	}
L574:
	;
	if v2696 != 0 {
		v3610 = v2696
		goto L420
	} else {
		goto L575
	}
L575:
	;
	v2698 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+96))
	if base.Ui64(v2698-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v3624 = int32(-2)
	goto L417
L577:
	;
	goto L578
L578:
	;
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+12))
	v2706 = v2704 + base.I32_wrap_i64(v2698)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+12)) = v2706
	if base.B2i32(v2706 < v2704)^base.B2i32(v2698 < int64(0)) != 0 {
		goto L579
	} else {
		goto L580
	}
L579:
	;
	v3624 = int32(-2)
	goto L417
L580:
	;
	goto L581
L581:
	;
	v2714 = *(*float64)(unsafe.Add(mBase, uint32(v2181)+88))
	if base.F64_eq(v2714, float64(0)) != 0 {
		v2765 = int32(1)
		goto L582
	} else {
		goto L583
	}
L582:
	;
	if v2765 == int32(0) {
		goto L591
	} else {
		goto L592
	}
L583:
	;
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+8))
	v2719 = base.F64_mul(v2714, float64(30))
	v2720 = base.I32_trunc_sat_f64_s(v2719)
	v2721 = v2717 + v2720
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+8)) = v2721
	v2723 = int32(0)
	if base.B2i32(v2720 < v2723)^base.B2i32(v2721 < v2717) != 0 {
		v2765 = v2723
		goto L582
	} else {
		goto L584
	}
L584:
	;
	v2730 = base.F64_sub(v2719, base.F64_convert_i32_s(v2720))
	if base.F64_eq(v2730, float64(0)) != 0 {
		v2765 = int32(1)
		goto L582
	} else {
		goto L585
	}
L585:
	;
	v2734 = base.F64_mul(v2730, float64(8.64e+10))
	v2735 = base.I64_trunc_sat_f64_s(v2734)
	v2737 = base.F64_sub(v2734, base.F64_convert_i64_s(v2735))
	if base.F64_gt(v2737, float64(0.5)) != 0 {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	v2749 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2750 = v2749 + v2748
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2750
	v2765 = base.B2i32(base.B2i32(v2748 < int64(0))^base.B2i32(v2750 < v2749) == int32(0))
	goto L582
L587:
	;
	v2748 = v2735 + int64(1)
	goto L586
L588:
	;
	goto L589
L589:
	;
	if base.F64_lt(v2737, float64(-0.5)) == int32(0) {
		v2748 = v2735
		goto L586
	} else {
		goto L590
	}
L590:
	;
	v2748 = v2735 - int64(1)
	goto L586
L591:
	;
	v3624 = int32(-2)
	goto L417
L592:
	;
	goto L593
L593:
	;
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v2181)+108))
	v2770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2769))))
	if v2770 != int32(45) {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	if v2770 == int32(84) {
		v3559 = v2769
		v3562 = int32(0)
		v3563 = v2684
		goto L425
	} else {
		goto L597
	}
L595:
	;
	goto L596
L596:
	;
	v2779 = v2769 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+108)) = v2779
	v2787 = F_ParseISO8601Number(m, v2779, v2181+int32(108), v2181+int32(96), v2181+int32(88))
	mBase = m.M
	v2788 = m.ExcPending
	if v2788 != 0 {
		goto L29
	} else {
		goto L599
	}
L597:
	;
	if v2770 == int32(0) {
		v3610 = v2770
		goto L420
	} else {
		goto L598
	}
L598:
	;
	v3624 = v2195
	goto L417
L599:
	;
	if v2787 != 0 {
		v3610 = v2787
		goto L420
	} else {
		goto L600
	}
L600:
	;
	v2789 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+96))
	if base.Ui64(v2789-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v3624 = int32(-2)
	goto L417
L602:
	;
	goto L603
L603:
	;
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v2188)+8))
	v2797 = v2795 + base.I32_wrap_i64(v2789)
	*(*int32)(unsafe.Add(mBase, uint32(v2188)+8)) = v2797
	if base.B2i32(v2797 < v2795)^base.B2i32(v2789 < int64(0)) != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v3624 = int32(-2)
	goto L417
L605:
	;
	goto L606
L606:
	;
	v2804 = *(*float64)(unsafe.Add(mBase, uint32(v2181)+88))
	if base.F64_ne(v2804, float64(0)) != 0 {
		goto L607
	} else {
		goto L608
	}
L607:
	;
	v2808 = base.F64_mul(v2804, float64(8.64e+10))
	v2809 = base.I64_trunc_sat_f64_s(v2808)
	v2811 = base.F64_sub(v2808, base.F64_convert_i64_s(v2809))
	if base.F64_gt(v2811, float64(0.5)) != 0 {
		goto L611
	} else {
		goto L612
	}
L608:
	;
	goto L609
L609:
	;
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2181)+108))
	v2838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2837))))
	if v2838 == int32(84) {
		v3559 = v2837
		v3562 = int32(0)
		v3563 = v2684
		goto L425
	} else {
		goto L616
	}
L610:
	;
	v2823 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2824 = v2823 + v2822
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2824
	if base.B2i32(v2822 < int64(0))^base.B2i32(v2824 < v2823) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L615
	}
L611:
	;
	v2822 = v2809 + int64(1)
	goto L610
L612:
	;
	goto L613
L613:
	;
	if base.F64_lt(v2811, float64(-0.5)) == int32(0) {
		v2822 = v2809
		goto L610
	} else {
		goto L614
	}
L614:
	;
	v2822 = v2809 - int64(1)
	goto L610
L615:
	;
	goto L609
L616:
	;
	if v2838 == int32(0) {
		v3610 = v2838
		goto L420
	} else {
		goto L617
	}
L617:
	;
	v3624 = v2195
	goto L417
L618:
	;
	if v2225 != 0 {
		v3624 = v2195
		goto L417
	} else {
		goto L698
	}
L619:
	;
	if v2302 != 0 {
		v3610 = v2254
		goto L420
	} else {
		goto L665
	}
L620:
	;
	v3010 = v2181 + int32(32)
	v3013 = int64(1000000)
	v3014 = int64(0)
	v3019 = int64(32)
	v3022 = int64(base.Ui64(v2297) >> (uint(v3019) % 64))
	v3025 = int64(4294967295)
	v3028 = v2297 & v3025
	v3029 = v3013 * v3028
	v3033 = int64(base.Ui64(v3029)>>(uint(v3019)%64)) + v3013*v3022
	v3040 = v3028*v3014 + v3033&v3025
	*(*int64)(unsafe.Add(mBase, uint32(v3010)+8)) = v2297*v3014 + v2297>>(uint(int64(63))%64)*v3013 + v3014*v3022 + int64(base.Ui64(v3033)>>(uint(v3019)%64)) + int64(base.Ui64(v3040)>>(uint(v3019)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3010))) = v3029&v3025 | v3040<<(uint(v3019)%64)
	goto L651
L621:
	;
	v2927 = v2181 + int32(16)
	v2930 = int64(60000000)
	v2931 = int64(0)
	v2936 = int64(32)
	v2939 = int64(base.Ui64(v2297) >> (uint(v2936) % 64))
	v2942 = int64(4294967295)
	v2945 = v2297 & v2942
	v2946 = v2930 * v2945
	v2950 = int64(base.Ui64(v2946)>>(uint(v2936)%64)) + v2930*v2939
	v2957 = v2945*v2931 + v2950&v2942
	*(*int64)(unsafe.Add(mBase, uint32(v2927)+8)) = v2297*v2931 + v2297>>(uint(int64(63))%64)*v2930 + v2931*v2939 + int64(base.Ui64(v2950)>>(uint(v2936)%64)) + int64(base.Ui64(v2957)>>(uint(v2936)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2927))) = v2946&v2942 | v2957<<(uint(v2936)%64)
	goto L637
L622:
	;
	v2847 = int64(3600000000)
	v2848 = int64(0)
	v2853 = int64(32)
	v2856 = int64(base.Ui64(v2297) >> (uint(v2853) % 64))
	v2859 = int64(4294967295)
	v2862 = v2297 & v2859
	v2863 = v2847 * v2862
	v2867 = int64(base.Ui64(v2863)>>(uint(v2853)%64)) + v2847*v2856
	v2874 = v2862*v2848 + v2867&v2859
	*(*int64)(unsafe.Add(mBase, uint32(v2181)+8)) = v2297*v2848 + v2297>>(uint(int64(63))%64)*v2847 + v2848*v2856 + int64(base.Ui64(v2867)>>(uint(v2853)%64)) + int64(base.Ui64(v2874)>>(uint(v2853)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2181))) = v2863&v2859 | v2874<<(uint(v2853)%64)
	goto L623
L623:
	;
	v2885 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+8))
	v2886 = *(*int64)(unsafe.Add(mBase, uint32(v2181)))
	if v2885 != v2886>>(uint(int64(63))%64) {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v3624 = int32(-2)
	goto L417
L625:
	;
	goto L626
L626:
	;
	v2891 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2892 = v2891 + v2886
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2892
	if base.B2i32(v2886 < int64(0))^base.B2i32(v2892 < v2891) != 0 {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v3624 = int32(-2)
	goto L417
L628:
	;
	goto L629
L629:
	;
	v2899 = int32(0)
	v2900 = int32(1)
	if base.F64_eq(v2300, float64(0)) != 0 {
		v3559 = v2288
		v3562 = v2899
		v3563 = v2900
		goto L425
	} else {
		goto L630
	}
L630:
	;
	v2904 = base.F64_mul(v2300, float64(3.6e+09))
	v2905 = base.I64_trunc_sat_f64_s(v2904)
	v2907 = base.F64_sub(v2904, base.F64_convert_i64_s(v2905))
	if base.F64_gt(v2907, float64(0.5)) != 0 {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	v2919 = v2918 + v2892
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2919
	if base.B2i32(v2918 < int64(0))^base.B2i32(v2919 < v2892) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L636
	}
L632:
	;
	v2918 = v2905 + int64(1)
	goto L631
L633:
	;
	goto L634
L634:
	;
	if base.F64_lt(v2907, float64(-0.5)) == int32(0) {
		v2918 = v2905
		goto L631
	} else {
		goto L635
	}
L635:
	;
	v2918 = v2905 - int64(1)
	goto L631
L636:
	;
	v3559 = v2288
	v3562 = v2899
	v3563 = v2900
	goto L425
L637:
	;
	v2968 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+24))
	v2969 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+16))
	if v2968 != v2969>>(uint(int64(63))%64) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v3624 = int32(-2)
	goto L417
L639:
	;
	goto L640
L640:
	;
	v2974 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v2975 = v2974 + v2969
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v2975
	if base.B2i32(v2969 < int64(0))^base.B2i32(v2975 < v2974) != 0 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	v3624 = int32(-2)
	goto L417
L642:
	;
	goto L643
L643:
	;
	v2982 = int32(0)
	v2983 = int32(1)
	if base.F64_eq(v2300, float64(0)) != 0 {
		v3559 = v2288
		v3562 = v2982
		v3563 = v2983
		goto L425
	} else {
		goto L644
	}
L644:
	;
	v2987 = base.F64_mul(v2300, float64(6e+07))
	v2988 = base.I64_trunc_sat_f64_s(v2987)
	v2990 = base.F64_sub(v2987, base.F64_convert_i64_s(v2988))
	if base.F64_gt(v2990, float64(0.5)) != 0 {
		goto L646
	} else {
		goto L647
	}
L645:
	;
	v3002 = v3001 + v2975
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3002
	if base.B2i32(v3001 < int64(0))^base.B2i32(v3002 < v2975) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L650
	}
L646:
	;
	v3001 = v2988 + int64(1)
	goto L645
L647:
	;
	goto L648
L648:
	;
	if base.F64_lt(v2990, float64(-0.5)) == int32(0) {
		v3001 = v2988
		goto L645
	} else {
		goto L649
	}
L649:
	;
	v3001 = v2988 - int64(1)
	goto L645
L650:
	;
	v3559 = v2288
	v3562 = v2982
	v3563 = v2983
	goto L425
L651:
	;
	v3051 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+40))
	v3052 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+32))
	if v3051 != v3052>>(uint(int64(63))%64) {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v3624 = int32(-2)
	goto L417
L653:
	;
	goto L654
L654:
	;
	v3057 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v3058 = v3057 + v3052
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3058
	if base.B2i32(v3052 < int64(0))^base.B2i32(v3058 < v3057) != 0 {
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v3624 = int32(-2)
	goto L417
L656:
	;
	goto L657
L657:
	;
	v3065 = int32(0)
	v3066 = int32(1)
	if base.F64_eq(v2300, float64(0)) != 0 {
		v3559 = v2288
		v3562 = v3065
		v3563 = v3066
		goto L425
	} else {
		goto L658
	}
L658:
	;
	v3070 = base.F64_mul(v2300, float64(1e+06))
	v3071 = base.I64_trunc_sat_f64_s(v3070)
	v3073 = base.F64_sub(v3070, base.F64_convert_i64_s(v3071))
	if base.F64_gt(v3073, float64(0.5)) != 0 {
		goto L660
	} else {
		goto L661
	}
L659:
	;
	v3085 = v3084 + v3058
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3085
	if base.B2i32(v3084 < int64(0))^base.B2i32(v3085 < v3058) != 0 {
		v3610 = int32(-2)
		goto L420
	} else {
		goto L664
	}
L660:
	;
	v3084 = v3071 + int64(1)
	goto L659
L661:
	;
	goto L662
L662:
	;
	if base.F64_lt(v3073, float64(-0.5)) == int32(0) {
		v3084 = v3071
		goto L659
	} else {
		goto L663
	}
L663:
	;
	v3084 = v3071 - int64(1)
	goto L659
L664:
	;
	v3559 = v2288
	v3562 = v3065
	v3563 = v3066
	goto L425
L665:
	;
	v3092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2220))))
	v3095 = v2220 + base.B2i32(v3092 == int32(45))
	v3096 = int32(_a_F_interval_in_6)
	v3100 = m.G0
	v3102 = v3100 - int32(32)
	v3103 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3102)+24)) = v3103
	*(*int64)(unsafe.Add(mBase, uint32(v3102)+16)) = v3103
	*(*int64)(unsafe.Add(mBase, uint32(v3102)+8)) = v3103
	*(*int64)(unsafe.Add(mBase, uint32(v3102))) = v3103
	v3111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[4])))
	if v3111 == int32(0) {
		goto L667
	} else {
		goto L668
	}
L666:
	;
	if base.B2i32(v3179 != int32(6))|v2225 != 0 {
		goto L618
	} else {
		goto L685
	}
L667:
	;
	v3179 = int32(0)
	goto L666
L668:
	;
	goto L669
L669:
	;
	v3115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[5])))
	if v3115 == int32(0) {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v3119 = v3095
	goto L673
L671:
	;
	goto L672
L672:
	;
	v3129 = v3096
	v3130 = v3111
	goto L676
L673:
	;
	v3125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3119))))
	if v3125 == v3111 {
		v3119 = v3119 + int32(1)
		goto L673
	} else {
		goto L675
	}
L674:
	;
	v3179 = v3119 - v3095
	goto L666
L675:
	;
	goto L674
L676:
	;
	v3137 = v3102 + int32(base.Ui32(v3130)>>(uint(int32(3))%32))&int32(28)
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v3137)))
	v3139 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3137))) = v3138 | v3139<<(uint(v3130)%32)
	v3143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3129)+1)))
	if v3143 != 0 {
		v3129 = v3129 + v3139
		v3130 = v3143
		goto L676
	} else {
		goto L678
	}
L677:
	;
	v3146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3095))))
	if v3146 == int32(0) {
		v3169 = v3095
		goto L679
	} else {
		goto L680
	}
L678:
	;
	goto L677
L679:
	;
	v3179 = v3169 - v3095
	goto L666
L680:
	;
	v3150 = v3095
	v3151 = v3146
	goto L681
L681:
	;
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(v3102+int32(base.Ui32(v3151)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3159)>>(uint(v3151)%32))&int32(1) == int32(0) {
		v3169 = v3150
		goto L679
	} else {
		goto L683
	}
L682:
	;
	v3169 = v3167
	goto L679
L683:
	;
	v3165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3150)+1)))
	v3167 = v3150 + int32(1)
	if v3165 != 0 {
		v3150 = v3167
		v3151 = v3165
		goto L681
	} else {
		goto L684
	}
L684:
	;
	goto L682
L685:
	;
	v3184 = v2181 - int32(-64)
	v3186 = base.I64_div_s(v2297, int64(10000))
	v3189 = int64(3600000000)
	v3190 = int64(0)
	v3195 = int64(32)
	v3198 = int64(base.Ui64(v3186) >> (uint(v3195) % 64))
	v3201 = int64(4294967295)
	v3204 = v3186 & v3201
	v3205 = v3189 * v3204
	v3209 = int64(base.Ui64(v3205)>>(uint(v3195)%64)) + v3189*v3198
	v3216 = v3204*v3190 + v3209&v3201
	*(*int64)(unsafe.Add(mBase, uint32(v3184)+8)) = v3186*v3190 + v3186>>(uint(int64(63))%64)*v3189 + v3190*v3198 + int64(base.Ui64(v3209)>>(uint(v3195)%64)) + int64(base.Ui64(v3216)>>(uint(v3195)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3184))) = v3205&v3201 | v3216<<(uint(v3195)%64)
	goto L686
L686:
	;
	v3227 = int32(-2)
	v3228 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+72))
	v3229 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+64))
	if v3228 != v3229>>(uint(int64(63))%64) {
		v3624 = v3227
		goto L417
	} else {
		goto L687
	}
L687:
	;
	v3233 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v3234 = v3233 + v3229
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3234
	if base.B2i32(v3229 < int64(0))^base.B2i32(v3234 < v3233) != 0 {
		v3624 = v3227
		goto L417
	} else {
		goto L688
	}
L688:
	;
	v3240 = int64(100)
	v3241 = base.I64_div_s(v2297, v3240)
	v3243 = base.I64_rem_s(v3241, v3240)
	v3245 = v3243 * int64(60000000)
	v3246 = v3234 + v3245
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3246
	if base.B2i32(v3245 < int64(0))^base.B2i32(v3246 < v3234) != 0 {
		v3624 = v3227
		goto L417
	} else {
		goto L689
	}
L689:
	;
	v3256 = (v2297 - v3241*int64(100)) * int64(1000000)
	v3257 = v3246 + v3256
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3257
	if base.B2i32(v3256 < int64(0))^base.B2i32(v3257 < v3246) != 0 {
		v3624 = v3227
		goto L417
	} else {
		goto L690
	}
L690:
	;
	if base.F64_eq(v2300, float64(0)) != 0 {
		goto L421
	} else {
		goto L691
	}
L691:
	;
	v3265 = base.I64_trunc_sat_f64_s(v2300)
	v3267 = base.F64_sub(v2300, base.F64_convert_i64_s(v3265))
	if base.F64_gt(v3267, float64(0.5)) != 0 {
		goto L693
	} else {
		goto L694
	}
L692:
	;
	v3279 = v3278 + v3257
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3279
	if base.B2i32(v3278 < int64(0))^base.B2i32(v3279 < v3257) == int32(0) {
		goto L421
	} else {
		goto L697
	}
L693:
	;
	v3278 = v3265 + int64(1)
	goto L692
L694:
	;
	goto L695
L695:
	;
	if base.F64_lt(v3267, float64(-0.5)) == int32(0) {
		v3278 = v3265
		goto L692
	} else {
		goto L696
	}
L696:
	;
	v3278 = v3265 - int64(1)
	goto L692
L697:
	;
	v3624 = v3227
	goto L417
L698:
	;
	v3288 = v2181 + int32(48)
	v3291 = int64(3600000000)
	v3292 = int64(0)
	v3297 = int64(32)
	v3300 = int64(base.Ui64(v2297) >> (uint(v3297) % 64))
	v3303 = int64(4294967295)
	v3306 = v2297 & v3303
	v3307 = v3291 * v3306
	v3311 = int64(base.Ui64(v3307)>>(uint(v3297)%64)) + v3291*v3300
	v3318 = v3306*v3292 + v3311&v3303
	*(*int64)(unsafe.Add(mBase, uint32(v3288)+8)) = v2297*v3292 + v2297>>(uint(int64(63))%64)*v3291 + v3292*v3300 + int64(base.Ui64(v3311)>>(uint(v3297)%64)) + int64(base.Ui64(v3318)>>(uint(v3297)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3288))) = v3307&v3303 | v3318<<(uint(v3297)%64)
	goto L699
L699:
	;
	v3329 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+56))
	v3330 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+48))
	if v3329 != v3330>>(uint(int64(63))%64) {
		goto L700
	} else {
		goto L701
	}
L700:
	;
	v3624 = int32(-2)
	goto L417
L701:
	;
	goto L702
L702:
	;
	v3335 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v3336 = v3335 + v3330
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3336
	if base.B2i32(v3330 < int64(0))^base.B2i32(v3336 < v3335) != 0 {
		goto L703
	} else {
		goto L704
	}
L703:
	;
	v3624 = int32(-2)
	goto L417
L704:
	;
	goto L705
L705:
	;
	if base.F64_eq(v2300, float64(0)) != 0 {
		goto L706
	} else {
		goto L707
	}
L706:
	;
	if v2302 == int32(0) {
		goto L421
	} else {
		goto L714
	}
L707:
	;
	v3346 = base.F64_mul(v2300, float64(3.6e+09))
	v3347 = base.I64_trunc_sat_f64_s(v3346)
	v3349 = base.F64_sub(v3346, base.F64_convert_i64_s(v3347))
	if base.F64_gt(v3349, float64(0.5)) != 0 {
		goto L709
	} else {
		goto L710
	}
L708:
	;
	v3361 = v3360 + v3336
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3361
	if base.B2i32(v3360 < int64(0))^base.B2i32(v3361 < v3336) == int32(0) {
		goto L706
	} else {
		goto L713
	}
L709:
	;
	v3360 = v3347 + int64(1)
	goto L708
L710:
	;
	goto L711
L711:
	;
	if base.F64_lt(v3349, float64(-0.5)) == int32(0) {
		v3360 = v3347
		goto L708
	} else {
		goto L712
	}
L712:
	;
	v3360 = v3347 - int64(1)
	goto L708
L713:
	;
	v3624 = int32(-2)
	goto L417
L714:
	;
	v3381 = F_ParseISO8601Number(m, v2288, v2181+int32(108), v2181+int32(96), v2181+int32(88))
	mBase = m.M
	v3382 = m.ExcPending
	if v3382 != 0 {
		goto L29
	} else {
		goto L715
	}
L715:
	;
	if v3381 != 0 {
		v3624 = v3381
		goto L417
	} else {
		goto L716
	}
L716:
	;
	v3383 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+96))
	v3384 = *(*float64)(unsafe.Add(mBase, uint32(v2181)+88))
	v3385 = int64(60000000)
	v3386 = int32(0)
	v3390 = m.G0
	v3392 = v3390 - int32(16)
	m.G0 = v3392
	v3394 = int64(63)
	F___multi3(m, v3392, v3383, v3383>>(uint(v3394)%64), v3385, int64(0))
	mBase = m.M
	v3399 = *(*int64)(unsafe.Add(mBase, uint32(v3392)+8))
	v3400 = *(*int64)(unsafe.Add(mBase, uint32(v3392)))
	if v3399 != v3400>>(uint(v3394)%64) {
		v3441 = v3386
		goto L718
	} else {
		goto L719
	}
L717:
	;
	if v3441 == int32(0) {
		goto L727
	} else {
		goto L728
	}
L718:
	;
	m.G0 = v3392 + int32(16)
	goto L717
L719:
	;
	v3404 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v3405 = v3404 + v3400
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3405
	if base.B2i32(v3400 < int64(0))^base.B2i32(v3405 < v3404) != 0 {
		v3441 = v3386
		goto L718
	} else {
		goto L720
	}
L720:
	;
	if base.F64_eq(v3384, float64(0)) != 0 {
		v3441 = int32(1)
		goto L718
	} else {
		goto L721
	}
L721:
	;
	v3415 = base.F64_mul(v3384, base.F64_convert_i64_u(v3385))
	v3416 = base.I64_trunc_sat_f64_s(v3415)
	v3418 = base.F64_sub(v3415, base.F64_convert_i64_s(v3416))
	if base.F64_gt(v3418, float64(0.5)) != 0 {
		goto L723
	} else {
		goto L724
	}
L722:
	;
	v3430 = v3405 + v3429
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3430
	v3441 = base.B2i32(base.B2i32(v3429 < int64(0))^base.B2i32(v3430 < v3405) == int32(0))
	goto L718
L723:
	;
	v3429 = v3416 + int64(1)
	goto L722
L724:
	;
	goto L725
L725:
	;
	if base.F64_lt(v3418, float64(-0.5)) == int32(0) {
		v3429 = v3416
		goto L722
	} else {
		goto L726
	}
L726:
	;
	v3429 = v3416 - int64(1)
	goto L722
L727:
	;
	v3624 = int32(-2)
	goto L417
L728:
	;
	goto L729
L729:
	;
	v3450 = *(*int32)(unsafe.Add(mBase, uint32(v2181)+108))
	v3451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3450))))
	if v3451 == int32(0) {
		v3624 = v3451
		goto L417
	} else {
		goto L730
	}
L730:
	;
	if v3451 != int32(58) {
		goto L731
	} else {
		goto L732
	}
L731:
	;
	v3624 = int32(-1)
	goto L417
L732:
	;
	goto L733
L733:
	;
	v3458 = v3450 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+108)) = v3458
	v3466 = F_ParseISO8601Number(m, v3458, v2181+int32(108), v2181+int32(96), v2181+int32(88))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L29
	} else {
		goto L734
	}
L734:
	;
	if v3466 != 0 {
		v3624 = v3466
		goto L417
	} else {
		goto L735
	}
L735:
	;
	v3469 = *(*int64)(unsafe.Add(mBase, uint32(v2181)+96))
	v3470 = *(*float64)(unsafe.Add(mBase, uint32(v2181)+88))
	v3471 = int64(1000000)
	v3472 = int32(0)
	v3476 = m.G0
	v3478 = v3476 - int32(16)
	m.G0 = v3478
	v3480 = int64(63)
	F___multi3(m, v3478, v3469, v3469>>(uint(v3480)%64), v3471, int64(0))
	mBase = m.M
	v3485 = *(*int64)(unsafe.Add(mBase, uint32(v3478)+8))
	v3486 = *(*int64)(unsafe.Add(mBase, uint32(v3478)))
	if v3485 != v3486>>(uint(v3480)%64) {
		v3527 = v3472
		goto L737
	} else {
		goto L738
	}
L736:
	;
	if v3527 == int32(0) {
		v3624 = int32(-2)
		goto L417
	} else {
		goto L746
	}
L737:
	;
	m.G0 = v3478 + int32(16)
	goto L736
L738:
	;
	v3490 = *(*int64)(unsafe.Add(mBase, uint32(v2188)))
	v3491 = v3490 + v3486
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3491
	if base.B2i32(v3486 < int64(0))^base.B2i32(v3491 < v3490) != 0 {
		v3527 = v3472
		goto L737
	} else {
		goto L739
	}
L739:
	;
	if base.F64_eq(v3470, float64(0)) != 0 {
		v3527 = int32(1)
		goto L737
	} else {
		goto L740
	}
L740:
	;
	v3501 = base.F64_mul(v3470, base.F64_convert_i64_u(v3471))
	v3502 = base.I64_trunc_sat_f64_s(v3501)
	v3504 = base.F64_sub(v3501, base.F64_convert_i64_s(v3502))
	if base.F64_gt(v3504, float64(0.5)) != 0 {
		goto L742
	} else {
		goto L743
	}
L741:
	;
	v3516 = v3491 + v3515
	*(*int64)(unsafe.Add(mBase, uint32(v2188))) = v3516
	v3527 = base.B2i32(base.B2i32(v3515 < int64(0))^base.B2i32(v3516 < v3491) == int32(0))
	goto L737
L742:
	;
	v3515 = v3502 + int64(1)
	goto L741
L743:
	;
	goto L744
L744:
	;
	if base.F64_lt(v3504, float64(-0.5)) == int32(0) {
		v3515 = v3502
		goto L741
	} else {
		goto L745
	}
L745:
	;
	v3515 = v3502 - int64(1)
	goto L741
L746:
	;
	v3537 = *(*int32)(unsafe.Add(mBase, uint32(v2181)+108))
	v3538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3537))))
	if v3538 != 0 {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	v3539 = int32(-1)
	goto L749
L748:
	;
	v3539 = int32(0)
	goto L749
L749:
	;
	v3624 = v3539
	goto L417
L750:
	;
	v3559 = v2288
	v3562 = v3548
	v3563 = v3548
	goto L425
L751:
	;
	goto L424
L752:
	;
	m.G0 = v39 + int32(528)
	return v3768
L753:
	;
	if v3669 == int32(-2) {
		goto L756
	} else {
		goto L757
	}
L754:
	;
	goto L755
L755:
	;
	v3704 = F_palloc(m, int32(16))
	mBase = m.M
	v3705 = m.ExcPending
	if v3705 != 0 {
		goto L29
	} else {
		goto L760
	}
L756:
	;
	v3695 = int32(-4)
	goto L758
L757:
	;
	v3695 = v3669
	goto L758
L758:
	;
	F_DateTimeParseError(m, v3695, v39+int32(8), v43, int32(_a_F_interval_in_11), v41)
	mBase = m.M
	v3700 = m.ExcPending
	if v3700 != 0 {
		goto L29
	} else {
		goto L759
	}
L759:
	;
	v3701 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v3701)
	v3768 = v2
	goto L752
L760:
	;
	v3706 = *(*int32)(unsafe.Add(mBase, uint32(v39)+500))
	switch v3706 - int32(9) {
	case 0:
		goto L764
	case 1:
		goto L762
	default:
		goto L763
	case 8:
		goto L765
	}
L761:
	;
	v3763 = F_AdjustIntervalForTypmod(m, v3704, v42, v41)
	mBase = m.M
	v3764 = m.ExcPending
	if v3764 != 0 {
		goto L29
	} else {
		goto L777
	}
L762:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3704)+8)) = int64(9223372034707292159)
	*(*int64)(unsafe.Add(mBase, uint32(v3704))) = int64(9223372036854775807)
	goto L761
L763:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3746 = m.ExcPending
	if v3746 != 0 {
		goto L29
	} else {
		goto L774
	}
L764:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3704)+8)) = int64(-9223372034707292160)
	*(*int64)(unsafe.Add(mBase, uint32(v3704))) = int64(-9223372036854775807 - 1)
	goto L761
L765:
	;
	v3709 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+516)))
	v3710 = int64(*(*int32)(unsafe.Add(mBase, uint32(v39)+520)))
	v3713 = v3709 + v3710*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v3713-int64(2147483648)) {
		goto L766
	} else {
		goto L767
	}
L766:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3704)+12)) = uint32(v3713)
	v3719 = *(*int32)(unsafe.Add(mBase, uint32(v39)+512))
	*(*int32)(unsafe.Add(mBase, uint32(v3704)+8)) = v3719
	v3721 = *(*int64)(unsafe.Add(mBase, uint32(v39)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v3704))) = v3721
	goto L761
L767:
	;
	goto L768
L768:
	;
	v3723 = F_errsave_start(m, v41)
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L29
	} else {
		goto L769
	}
L769:
	;
	if v3723 == int32(0) {
		v3768 = v2
		goto L752
	} else {
		goto L770
	}
L770:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v3729 = m.ExcPending
	if v3729 != 0 {
		goto L29
	} else {
		goto L771
	}
L771:
	;
	F_errmsg(m, int32(_a_F_interval_in_12), int32(0))
	mBase = m.M
	v3733 = m.ExcPending
	if v3733 != 0 {
		goto L29
	} else {
		goto L772
	}
L772:
	;
	F_errsave_finish(m, v41, int32(_a_F_interval_in_13), int32(948), int32(_a_F_interval_in_14))
	mBase = m.M
	v3738 = m.ExcPending
	if v3738 != 0 {
		goto L29
	} else {
		goto L773
	}
L773:
	;
	v3768 = v2
	goto L752
L774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v43
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v39)+500))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v3748
	F_errmsg_internal(m, int32(_a_F_interval_in_15), v39)
	mBase = m.M
	v3752 = m.ExcPending
	if v3752 != 0 {
		goto L29
	} else {
		goto L775
	}
L775:
	;
	F_errfinish(m, int32(_a_F_interval_in_13), int32(961), int32(_a_F_interval_in_14))
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L29
	} else {
		goto L776
	}
L776:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L777:
	;
	v3768 = base.I64_extend_i32_u(v3704)
	goto L752
}
func F_interval_lerp(m *base.Module, l0 int64, l1 int64, l2 float64) int64 {
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	v5 = int32(0)
	v10 = F_DirectFunctionCall2Coll(m, int32(1606), v5, l1, l0)
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v15 = F_DirectFunctionCall2Coll(m, int32(1605), v5, v10, base.I64_reinterpret_f64(l2))
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = F_DirectFunctionCall2Coll(m, int32(1604), v5, v15, l0)
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				return v17
			}
		}
	}
}
func F_interval_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v51 int64
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+12)))
	v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+8)))
	v21 = v17*int64(30) + v20
	v30 = int64(32)
	v31 = int64(20)
	v33 = int64(base.Ui64(v21) >> (uint(v30) % 64))
	v36 = int64(4294967295)
	v37 = int64(500654080)
	v39 = v21 & v36
	v40 = v37 * v39
	v44 = int64(base.Ui64(v40)>>(uint(v30)%64)) + v37*v33
	v51 = v39*v31 + v44&v36
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v21*int64(0) + v21>>(uint(int64(63))%64)*int64(86400000000) + v31*v33 + int64(base.Ui64(v44)>>(uint(v30)%64)) + int64(base.Ui64(v51)>>(uint(v30)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v40&v36 | v51<<(uint(v30)%64)
	v63 = v14 + int32(16)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+12)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	v69 = v65*int64(30) + v68
	v78 = int64(32)
	v79 = int64(20)
	v81 = int64(base.Ui64(v69) >> (uint(v78) % 64))
	v84 = int64(4294967295)
	v85 = int64(500654080)
	v87 = v69 & v84
	v88 = v85 * v87
	v92 = int64(base.Ui64(v88)>>(uint(v78)%64)) + v85*v81
	v99 = v87*v79 + v92&v84
	*(*int64)(unsafe.Add(mBase, uint32(v63)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v63))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	return base.I64_extend_i32_u(base.B2i32(v119^v120|(base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115)))+(v114+v113>>(uint(v124)%64))^(base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112)))+(v111+v110>>(uint(v124)%64)))) != int64(0)))
}
func F_interval_part_common(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 float64
	_ = v211
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int64
	_ = v307
	var v312 int64
	_ = v312
	var v315 int64
	_ = v315
	var v316 int64
	_ = v316
	var v321 int64
	_ = v321
	var v324 int64
	_ = v324
	var v327 int64
	_ = v327
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v342 int64
	_ = v342
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v361 int64
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v393 int64
	_ = v393
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v428 int64
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v444 int64
	_ = v444
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v23 = int32(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v27 = v25 & v23
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = v23
	goto L5
L4:
	;
	v28 = int32(4)
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v25 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v59 = F_downcase_truncate_identifier(m, v19+v28, v57, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v36 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v47 = int32(1)
	if v27 != 0 {
		v57 = int32(base.Ui32(v25)>>(uint(v47)%32)) - v47
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v39 = int32(16)
	goto L12
L11:
	;
	v39 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = int32(4)
	goto L15
L14:
	;
	v46 = v39
	goto L15
L15:
	;
	v57 = v46
	goto L6
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v62 = v16 + int32(76)
	v66 = Fn14210(m, v59, v62, int32(_a_F_interval_part_common_0), int32(_a_F_interval_part_common_1), int32(_a_F_interval_part_common_2))
	mBase = m.M
	goto L18
L18:
	;
	if v66 == int32(31) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = Fn14210(m, v59, v62, int32(_a_F_interval_part_common_3), int32(_a_F_interval_part_common_4), int32(_a_F_interval_part_common_5))
	mBase = m.M
	goto L22
L20:
	;
	v73 = v66
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v74 != int32(2147483647) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v73 = v72
	goto L21
L23:
	;
	m.G0 = v16 + int32(80)
	return v444
L24:
	;
	if v73 == int32(17) {
		goto L64
	} else {
		goto L65
	}
L25:
	;
	v94 = int32(0)
	if base.B2i32(v73 == v94)|base.B2i32(v73 == int32(17)) == v94 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	if v74 != int32(-2147483648) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v86 != int32(2147483647) {
		goto L24
	} else {
		goto L32
	}
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v79 != int32(-2147483648) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	if v82 != int64(-9223372036854775807-1) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v93 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L25
L32:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	if v89 != int64(9223372036854775807) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v93 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L25
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if base.Ui32(v123) <= base.Ui32(int32(30)) {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v109 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v59
	F_errmsg(m, int32(_a_F_interval_part_common_11), v16+int32(48))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_13), int32(_a_F_interval_part_common_14))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
	v444 = int64(0)
	goto L23
L43:
	;
	if l1 != 0 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v127 = int32(1) << (uint(v123) % 32)
	if v127&int32(506464256) != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	if v127&int32(1640759296) != 0 {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v141 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v59
	F_errmsg(m, int32(_a_F_interval_part_common_7), v16-int32(-64))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_15), int32(_a_F_interval_part_common_14))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	if base.F64_lt(v93, float64(0)) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v444 = base.I64_reinterpret_f64(v93)
	goto L23
L57:
	;
	v162 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11926), int64(0), int64(-1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v169 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11937), int64(0), int64(-1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v444 = v162
	goto L23
L61:
	;
	v444 = v169
	goto L23
L62:
	;
	if l1 != 0 {
		goto L114
	} else {
		goto L115
	}
L63:
	;
	v428 = base.I64_extend32_s(v192) + base.I64_extend32_s(v189)*int64(1000000)
	goto L62
L64:
	;
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	v179 = base.I64_div_s(v177, int64(3600000000))
	v182 = v179*int64(-3600000000) + v177
	v184 = base.I64_div_s(v182, int64(60000000))
	v187 = v184*int64(-60000000) + v182
	v189 = base.I64_div_s(v187, int64(1000000))
	v192 = v189*int64(4293967296) + v187
	v193 = base.I32_wrap_i64(v192)
	v194 = base.I32_wrap_i64(v189)
	v196 = base.I32_rem_s(v74, int32(12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	switch v198 - int32(18) {
	case 0:
		goto L77
	case 1:
		goto L76
	case 2:
		v428 = v179
		goto L62
	case 3:
		goto L75
	case 4:
		goto L74
	case 5:
		goto L73
	case 6:
		goto L72
	case 7:
		goto L71
	case 8:
		goto L70
	case 9:
		goto L69
	case 10:
		goto L68
	case 11:
		goto L78
	case 12:
		goto L63
	default:
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v73 != 0 {
		goto L95
	} else {
		goto L96
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L90
	}
L68:
	;
	v269 = base.I32_div_s(v74, int32(_a_F_interval_part_common_6))
	v428 = base.I64_extend_i32_s(v269)
	goto L62
L69:
	;
	v266 = base.I32_div_s(v74, int32(1200))
	v428 = base.I64_extend_i32_s(v266)
	goto L62
L70:
	;
	v263 = base.I32_div_s(v74, int32(120))
	v428 = base.I64_extend_i32_s(v263)
	goto L62
L71:
	;
	v260 = base.I32_div_s(v74, int32(12))
	v428 = base.I64_extend_i32_s(v260)
	goto L62
L72:
	;
	if int32(0) <= v74 {
		goto L87
	} else {
		goto L88
	}
L73:
	;
	v428 = base.I64_extend_i32_s(v196)
	goto L62
L74:
	;
	v236 = base.I32_div_s(v197, int32(7))
	v428 = base.I64_extend_i32_s(v236)
	goto L62
L75:
	;
	v428 = base.I64_extend_i32_s(v197)
	goto L62
L76:
	;
	v428 = base.I64_extend32_s(v184)
	goto L62
L77:
	;
	if l1 != 0 {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	if l1 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v207 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v192)+base.I64_extend32_s(v189)*int64(1000000), int32(3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v211 = float64(1000)
	v444 = base.I64_reinterpret_f64(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v194), v211), base.F64_div(base.F64_convert_i32_s(v193), v211)))
	goto L23
L82:
	;
	v444 = base.I64_extend_i32_u(v207)
	goto L23
L83:
	;
	v224 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v192)+base.I64_extend32_s(v189)*int64(1000000), int32(6))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v444 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i32_s(v193), float64(1e+06)), base.F64_convert_i32_s(v194)))
	goto L23
L86:
	;
	v444 = base.I64_extend_i32_u(v224)
	goto L23
L87:
	;
	v244 = base.I32_div_u_s(v196&int32(255), int32(3))
	v428 = base.I64_extend_i32_u(v244 + int32(1))
	goto L62
L88:
	;
	goto L89
L89:
	;
	v251 = base.I32_rem_s(int32(0)-v74, int32(12))
	v254 = base.I32_div_s(base.I32_extend8_s(v251), int32(-3))
	v428 = base.I64_extend8_s(base.I64_extend_i32_u(v254 - int32(1)))
	goto L62
L90:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v279 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v59
	F_errmsg(m, int32(_a_F_interval_part_common_7), v16)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_9), int32(_a_F_interval_part_common_10))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L109
	}
L96:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if v291 != int32(11) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if l1 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v295 = v16 + int32(16)
	v296 = int32(12)
	v297 = base.I32_rem_s(v74, v296)
	v302 = base.I32_div_s(v74, v296)
	v307 = int64(*(*int32)(unsafe.Add(mBase, uint32(v30)+8)))
	v312 = (base.I64_extend_i32_s(v297*int32(120)) + base.I64_extend_i32_s(v302)*int64(1461) + v307<<(uint(int64(2))%64)) * int64(21600)
	v315 = int64(1000000)
	v316 = int64(0)
	v321 = int64(32)
	v324 = int64(base.Ui64(v312) >> (uint(v321) % 64))
	v327 = int64(4294967295)
	v330 = v312 & v327
	v331 = v315 * v330
	v335 = int64(base.Ui64(v331)>>(uint(v321)%64)) + v315*v324
	v342 = v330*v316 + v335&v327
	*(*int64)(unsafe.Add(mBase, uint32(v295)+8)) = v312*v316 + v312>>(uint(int64(63))%64)*v315 + v316*v324 + int64(base.Ui64(v335)>>(uint(v321)%64)) + int64(base.Ui64(v342)>>(uint(v321)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v331&v327 | v342<<(uint(v321)%64)
	goto L101
L99:
	;
	goto L100
L100:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v382 = int32(12)
	v383 = base.I32_div_s(v74, v382)
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	v444 = base.I64_reinterpret_f64(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v378), float64(86400)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v74-v383*v382), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v383), float64(3.15576e+07)), base.F64_div(base.F64_convert_i64_s(v393), float64(1e+06))))))
	goto L23
L101:
	;
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	if v354 != v355>>(uint(int64(63))%64) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v370 = F_int64_div_fast_to_numeric(m, v353, int32(6))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L106
	}
L103:
	;
	v361 = v353 + v355
	if base.B2i32(v353 < int64(0))^base.B2i32(v361 < v355) != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v365 = F_int64_div_fast_to_numeric(m, v361, int32(6))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v444 = base.I64_extend_i32_u(v365)
	goto L23
L106:
	;
	v372 = F_int64_to_numeric(m, v312)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v375 = F_numeric_add_safe(m, v370, v372, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v444 = base.I64_extend_i32_u(v375)
	goto L23
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v409 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v59
	F_errmsg(m, int32(_a_F_interval_part_common_11), v16+int32(32))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_12), int32(_a_F_interval_part_common_10))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v429 = F_int64_to_numeric(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v444 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v428))
	goto L23
L117:
	;
	v444 = base.I64_extend_i32_u(v429)
	goto L23
}
func F_interval_sum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v44 int64
	_ = v44
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v5 != 0 {
		v19 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
		return int64(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v6 == int32(0) {
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			return int64(0)
		} else {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v10 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
			if v9+v10 != int64(0)-v13 {
				v23 = int64(0)
				if base.B2i32(v23 < v9)&base.B2i32(v23 < v13) == int32(0) {
					v31 = F_palloc(m, int32(16))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
						if int64(0) < v35 {
							*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(9223372034707292159)
							*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(9223372036854775807)
							return base.I64_extend_i32_u(v31)
						} else {
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
							if int64(0) < v44 {
								*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(-9223372034707292160)
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(-9223372036854775807 - 1)
								return base.I64_extend_i32_u(v31)
							} else {
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v53
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = v55
								return base.I64_extend_i32_u(v31)
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_sum_0), int32(0))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_sum_1), int32(_a_F_interval_sum_2), int32(_a_F_interval_sum_3))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				return int64(0)
			}
		}
	}
}
func F_interval_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int64
	_ = v51
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != int32(463) {
		v51 = v5
		return v51
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 != int32(7) {
			v51 = v5
			return v51
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+32)))
			if v17 != 0 {
				v51 = v5
				return v51
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
				if v19 < int32(0) {
					v44 = F_relabel_to_typmod(m, v18, v19)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						v51 = base.I64_extend_i32_u(v44)
						return v51
					}
				} else {
					v22 = F_exprTypmod(m, v18)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = F_intervaltypmodleastfield(m, v22)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int64(0)
						} else {
							v28 = F_intervaltypmodleastfield(m, v19)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int64(0)
							} else {
								if base.Ui32(v26) < base.Ui32(v28) {
									v51 = v5
									return v51
								} else {
									if v26 != 0 {
										v44 = F_relabel_to_typmod(m, v18, v19)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int64(0)
										} else {
											v51 = base.I64_extend_i32_u(v44)
											return v51
										}
									} else {
										v32 = v19 & int32(_a_F_interval_support_0)
										if base.Ui32(int32(5)) < base.Ui32(v32) {
											v44 = F_relabel_to_typmod(m, v18, v19)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int64(0)
											} else {
												v51 = base.I64_extend_i32_u(v44)
												return v51
											}
										} else {
											if v22 < int32(0) {
												v38 = int32(-1)
											} else {
												v38 = v22
											}
											if base.Ui32(v32) < base.Ui32(v38&int32(_a_F_interval_support_0)) {
												v51 = v5
												return v51
											} else {
												v44 = F_relabel_to_typmod(m, v18, v19)
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return int64(0)
												} else {
													v51 = base.I64_extend_i32_u(v44)
													return v51
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
