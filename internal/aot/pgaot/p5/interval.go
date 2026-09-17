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
func F_interval_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v7 int32
	_ = v7
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v165 int32
	_ = v165
	var v182 int32
	_ = v182
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v302 int64
	_ = v302
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v310 int64
	_ = v310
	var v317 int64
	_ = v317
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v334 int64
	_ = v334
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v391 int32
	_ = v391
	var v392 int64
	_ = v392
	var v395 int64
	_ = v395
	var v396 int64
	_ = v396
	var v401 int64
	_ = v401
	var v404 int64
	_ = v404
	var v407 int64
	_ = v407
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v415 int64
	_ = v415
	var v422 int64
	_ = v422
	var v433 int64
	_ = v433
	var v434 int64
	_ = v434
	var v438 int64
	_ = v438
	var v444 int64
	_ = v444
	var v446 int64
	_ = v446
	var v447 int64
	_ = v447
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v456 int64
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int64
	_ = v470
	var v472 int64
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v663 int32
	_ = v663
	var v695 int32
	_ = v695
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v782 int32
	_ = v782
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v895 int32
	_ = v895
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v999 int64
	_ = v999
	var v1001 int32
	_ = v1001
	var v1005 float64
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1034 int64
	_ = v1034
	var v1035 int64
	_ = v1035
	var v1040 int64
	_ = v1040
	var v1043 int64
	_ = v1043
	var v1046 int64
	_ = v1046
	var v1049 int64
	_ = v1049
	var v1050 int64
	_ = v1050
	var v1054 int64
	_ = v1054
	var v1061 int64
	_ = v1061
	var v1072 int64
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1089 int64
	_ = v1089
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1108 int64
	_ = v1108
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1192 float64
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1203 float64
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 float64
	_ = v1209
	var v1210 int64
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1215 float64
	_ = v1215
	var v1219 int64
	_ = v1219
	var v1221 int64
	_ = v1221
	var v1227 int64
	_ = v1227
	var v1229 float64
	_ = v1229
	var v1233 int64
	_ = v1233
	var v1234 int64
	_ = v1234
	var v1243 int64
	_ = v1243
	var v1245 float64
	_ = v1245
	var v1256 int64
	_ = v1256
	var v1257 int64
	_ = v1257
	var v1273 int32
	_ = v1273
	var v1276 int64
	_ = v1276
	var v1277 int64
	_ = v1277
	var v1282 int64
	_ = v1282
	var v1285 int64
	_ = v1285
	var v1288 int64
	_ = v1288
	var v1291 int64
	_ = v1291
	var v1292 int64
	_ = v1292
	var v1296 int64
	_ = v1296
	var v1303 int64
	_ = v1303
	var v1314 int64
	_ = v1314
	var v1315 int64
	_ = v1315
	var v1320 int64
	_ = v1320
	var v1321 int64
	_ = v1321
	var v1331 float64
	_ = v1331
	var v1332 int64
	_ = v1332
	var v1334 float64
	_ = v1334
	var v1345 int64
	_ = v1345
	var v1346 int64
	_ = v1346
	var v1362 int32
	_ = v1362
	var v1365 int64
	_ = v1365
	var v1366 int64
	_ = v1366
	var v1371 int64
	_ = v1371
	var v1374 int64
	_ = v1374
	var v1377 int64
	_ = v1377
	var v1380 int64
	_ = v1380
	var v1381 int64
	_ = v1381
	var v1385 int64
	_ = v1385
	var v1392 int64
	_ = v1392
	var v1403 int64
	_ = v1403
	var v1404 int64
	_ = v1404
	var v1409 int64
	_ = v1409
	var v1410 int64
	_ = v1410
	var v1420 float64
	_ = v1420
	var v1421 int64
	_ = v1421
	var v1423 float64
	_ = v1423
	var v1434 int64
	_ = v1434
	var v1435 int64
	_ = v1435
	var v1453 int32
	_ = v1453
	var v1456 int64
	_ = v1456
	var v1457 int64
	_ = v1457
	var v1462 int64
	_ = v1462
	var v1465 int64
	_ = v1465
	var v1468 int64
	_ = v1468
	var v1471 int64
	_ = v1471
	var v1472 int64
	_ = v1472
	var v1476 int64
	_ = v1476
	var v1483 int64
	_ = v1483
	var v1494 int64
	_ = v1494
	var v1495 int64
	_ = v1495
	var v1500 int64
	_ = v1500
	var v1501 int64
	_ = v1501
	var v1511 float64
	_ = v1511
	var v1512 int64
	_ = v1512
	var v1514 float64
	_ = v1514
	var v1525 int64
	_ = v1525
	var v1526 int64
	_ = v1526
	var v1542 int32
	_ = v1542
	var v1545 int64
	_ = v1545
	var v1546 int64
	_ = v1546
	var v1551 int64
	_ = v1551
	var v1554 int64
	_ = v1554
	var v1557 int64
	_ = v1557
	var v1560 int64
	_ = v1560
	var v1561 int64
	_ = v1561
	var v1565 int64
	_ = v1565
	var v1572 int64
	_ = v1572
	var v1583 int64
	_ = v1583
	var v1584 int64
	_ = v1584
	var v1589 int64
	_ = v1589
	var v1590 int64
	_ = v1590
	var v1600 float64
	_ = v1600
	var v1601 int64
	_ = v1601
	var v1603 float64
	_ = v1603
	var v1614 int64
	_ = v1614
	var v1615 int64
	_ = v1615
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1645 float64
	_ = v1645
	var v1646 int64
	_ = v1646
	var v1648 float64
	_ = v1648
	var v1659 int64
	_ = v1659
	var v1660 int64
	_ = v1660
	var v1661 int64
	_ = v1661
	var v1681 int64
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1701 float64
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1711 float64
	_ = v1711
	var v1715 float64
	_ = v1715
	var v1716 int64
	_ = v1716
	var v1718 float64
	_ = v1718
	var v1729 int64
	_ = v1729
	var v1730 int64
	_ = v1730
	var v1731 int64
	_ = v1731
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1765 int32
	_ = v1765
	var v1767 float64
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1777 float64
	_ = v1777
	var v1781 float64
	_ = v1781
	var v1782 int64
	_ = v1782
	var v1784 float64
	_ = v1784
	var v1795 int64
	_ = v1795
	var v1796 int64
	_ = v1796
	var v1797 int64
	_ = v1797
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1851 int64
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1868 int32
	_ = v1868
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1892 int64
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1909 int32
	_ = v1909
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1933 int64
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1950 int32
	_ = v1950
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v2008 int32
	_ = v2008
	var v2011 int32
	_ = v2011
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2039 int32
	_ = v2039
	var v2047 int32
	_ = v2047
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2062 int32
	_ = v2062
	var v2063 int64
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2094 int32
	_ = v2094
	var v2137 int32
	_ = v2137
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2173 int32
	_ = v2173
	var v2176 int64
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2184 int32
	_ = v2184
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2241 int32
	_ = v2241
	var v2257 float64
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2262 int32
	_ = v2262
	var v2263 float64
	_ = v2263
	var v2272 int32
	_ = v2272
	var v2280 float64
	_ = v2280
	var v2281 int64
	_ = v2281
	var v2284 float64
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2302 int64
	_ = v2302
	var v2310 int32
	_ = v2310
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2358 int32
	_ = v2358
	var v2364 int32
	_ = v2364
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2378 int32
	_ = v2378
	var v2383 int64
	_ = v2383
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2425 float64
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2435 float64
	_ = v2435
	var v2440 float64
	_ = v2440
	var v2441 int64
	_ = v2441
	var v2443 float64
	_ = v2443
	var v2454 int64
	_ = v2454
	var v2455 int64
	_ = v2455
	var v2456 int64
	_ = v2456
	var v2470 int64
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2487 int32
	_ = v2487
	var v2492 float64
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2502 float64
	_ = v2502
	var v2507 float64
	_ = v2507
	var v2508 int64
	_ = v2508
	var v2510 float64
	_ = v2510
	var v2521 int64
	_ = v2521
	var v2522 int64
	_ = v2522
	var v2523 int64
	_ = v2523
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2545 int32
	_ = v2545
	var v2550 float64
	_ = v2550
	var v2551 int64
	_ = v2551
	var v2553 float64
	_ = v2553
	var v2564 int64
	_ = v2564
	var v2565 int64
	_ = v2565
	var v2566 int64
	_ = v2566
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2583 int32
	_ = v2583
	var v2584 int64
	_ = v2584
	var v2585 int64
	_ = v2585
	var v2587 int64
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2596 int32
	_ = v2596
	var v2598 int64
	_ = v2598
	var v2600 int32
	_ = v2600
	var v2610 float64
	_ = v2610
	var v2611 int64
	_ = v2611
	var v2613 float64
	_ = v2613
	var v2624 int64
	_ = v2624
	var v2625 int64
	_ = v2625
	var v2626 int64
	_ = v2626
	var v2638 int32
	_ = v2638
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2656 int32
	_ = v2656
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2668 int32
	_ = v2668
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2682 int64
	_ = v2682
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2698 float64
	_ = v2698
	var v2701 int32
	_ = v2701
	var v2703 float64
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2707 int32
	_ = v2707
	var v2714 float64
	_ = v2714
	var v2718 float64
	_ = v2718
	var v2719 int64
	_ = v2719
	var v2721 float64
	_ = v2721
	var v2732 int64
	_ = v2732
	var v2733 int64
	_ = v2733
	var v2734 int64
	_ = v2734
	var v2749 int32
	_ = v2749
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2763 int32
	_ = v2763
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2773 int64
	_ = v2773
	var v2779 int32
	_ = v2779
	var v2781 int32
	_ = v2781
	var v2788 float64
	_ = v2788
	var v2792 float64
	_ = v2792
	var v2793 int64
	_ = v2793
	var v2795 float64
	_ = v2795
	var v2806 int64
	_ = v2806
	var v2807 int64
	_ = v2807
	var v2808 int64
	_ = v2808
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2831 int64
	_ = v2831
	var v2832 int64
	_ = v2832
	var v2837 int64
	_ = v2837
	var v2840 int64
	_ = v2840
	var v2843 int64
	_ = v2843
	var v2846 int64
	_ = v2846
	var v2847 int64
	_ = v2847
	var v2851 int64
	_ = v2851
	var v2858 int64
	_ = v2858
	var v2869 int64
	_ = v2869
	var v2870 int64
	_ = v2870
	var v2875 int64
	_ = v2875
	var v2876 int64
	_ = v2876
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2888 float64
	_ = v2888
	var v2889 int64
	_ = v2889
	var v2891 float64
	_ = v2891
	var v2902 int64
	_ = v2902
	var v2903 int64
	_ = v2903
	var v2911 int32
	_ = v2911
	var v2914 int64
	_ = v2914
	var v2915 int64
	_ = v2915
	var v2920 int64
	_ = v2920
	var v2923 int64
	_ = v2923
	var v2926 int64
	_ = v2926
	var v2929 int64
	_ = v2929
	var v2930 int64
	_ = v2930
	var v2934 int64
	_ = v2934
	var v2941 int64
	_ = v2941
	var v2952 int64
	_ = v2952
	var v2953 int64
	_ = v2953
	var v2958 int64
	_ = v2958
	var v2959 int64
	_ = v2959
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2971 float64
	_ = v2971
	var v2972 int64
	_ = v2972
	var v2974 float64
	_ = v2974
	var v2985 int64
	_ = v2985
	var v2986 int64
	_ = v2986
	var v2994 int32
	_ = v2994
	var v2997 int64
	_ = v2997
	var v2998 int64
	_ = v2998
	var v3003 int64
	_ = v3003
	var v3006 int64
	_ = v3006
	var v3009 int64
	_ = v3009
	var v3012 int64
	_ = v3012
	var v3013 int64
	_ = v3013
	var v3017 int64
	_ = v3017
	var v3024 int64
	_ = v3024
	var v3035 int64
	_ = v3035
	var v3036 int64
	_ = v3036
	var v3041 int64
	_ = v3041
	var v3042 int64
	_ = v3042
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3054 float64
	_ = v3054
	var v3055 int64
	_ = v3055
	var v3057 float64
	_ = v3057
	var v3068 int64
	_ = v3068
	var v3069 int64
	_ = v3069
	var v3076 int32
	_ = v3076
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3087 int64
	_ = v3087
	var v3095 int32
	_ = v3095
	var v3099 int32
	_ = v3099
	var v3103 int32
	_ = v3103
	var v3109 int32
	_ = v3109
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3127 int32
	_ = v3127
	var v3130 int32
	_ = v3130
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3143 int32
	_ = v3143
	var v3149 int32
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3153 int32
	_ = v3153
	var v3163 int32
	_ = v3163
	var v3168 int32
	_ = v3168
	var v3170 int64
	_ = v3170
	var v3173 int64
	_ = v3173
	var v3174 int64
	_ = v3174
	var v3179 int64
	_ = v3179
	var v3182 int64
	_ = v3182
	var v3185 int64
	_ = v3185
	var v3188 int64
	_ = v3188
	var v3189 int64
	_ = v3189
	var v3193 int64
	_ = v3193
	var v3200 int64
	_ = v3200
	var v3211 int32
	_ = v3211
	var v3212 int64
	_ = v3212
	var v3213 int64
	_ = v3213
	var v3217 int64
	_ = v3217
	var v3218 int64
	_ = v3218
	var v3224 int64
	_ = v3224
	var v3225 int64
	_ = v3225
	var v3227 int64
	_ = v3227
	var v3229 int64
	_ = v3229
	var v3230 int64
	_ = v3230
	var v3240 int64
	_ = v3240
	var v3241 int64
	_ = v3241
	var v3249 int64
	_ = v3249
	var v3251 float64
	_ = v3251
	var v3262 int64
	_ = v3262
	var v3263 int64
	_ = v3263
	var v3272 int32
	_ = v3272
	var v3275 int64
	_ = v3275
	var v3276 int64
	_ = v3276
	var v3281 int64
	_ = v3281
	var v3284 int64
	_ = v3284
	var v3287 int64
	_ = v3287
	var v3290 int64
	_ = v3290
	var v3291 int64
	_ = v3291
	var v3295 int64
	_ = v3295
	var v3302 int64
	_ = v3302
	var v3313 int64
	_ = v3313
	var v3314 int64
	_ = v3314
	var v3319 int64
	_ = v3319
	var v3320 int64
	_ = v3320
	var v3330 float64
	_ = v3330
	var v3331 int64
	_ = v3331
	var v3333 float64
	_ = v3333
	var v3344 int64
	_ = v3344
	var v3345 int64
	_ = v3345
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3367 int64
	_ = v3367
	var v3368 float64
	_ = v3368
	var v3369 int64
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3378 int64
	_ = v3378
	var v3383 int64
	_ = v3383
	var v3384 int64
	_ = v3384
	var v3388 int64
	_ = v3388
	var v3389 int64
	_ = v3389
	var v3399 float64
	_ = v3399
	var v3400 int64
	_ = v3400
	var v3402 float64
	_ = v3402
	var v3413 int64
	_ = v3413
	var v3414 int64
	_ = v3414
	var v3425 int32
	_ = v3425
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3442 int32
	_ = v3442
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3453 int64
	_ = v3453
	var v3454 float64
	_ = v3454
	var v3455 int64
	_ = v3455
	var v3456 int32
	_ = v3456
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3464 int64
	_ = v3464
	var v3469 int64
	_ = v3469
	var v3470 int64
	_ = v3470
	var v3474 int64
	_ = v3474
	var v3475 int64
	_ = v3475
	var v3485 float64
	_ = v3485
	var v3486 int64
	_ = v3486
	var v3488 float64
	_ = v3488
	var v3499 int64
	_ = v3499
	var v3500 int64
	_ = v3500
	var v3511 int32
	_ = v3511
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3525 int32
	_ = v3525
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3532 int32
	_ = v3532
	var v3543 int32
	_ = v3543
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3551 int32
	_ = v3551
	var v3594 int32
	_ = v3594
	var v3606 int32
	_ = v3606
	var v3649 int32
	_ = v3649
	var v3676 int32
	_ = v3676
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3686 int32
	_ = v3686
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3691 int64
	_ = v3691
	var v3692 int64
	_ = v3692
	var v3695 int64
	_ = v3695
	var v3701 int32
	_ = v3701
	var v3703 int64
	_ = v3703
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3712 int32
	_ = v3712
	var v3716 int32
	_ = v3716
	var v3721 int32
	_ = v3721
	var v3729 int32
	_ = v3729
	var v3731 int32
	_ = v3731
	var v3735 int32
	_ = v3735
	var v3740 int32
	_ = v3740
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	v2 = int64(0)
	v7 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(528)
	m.G0 = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+504)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v38)+512)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v38)+520)) = v7
	v53 = v38 + int32(384)
	v55 = v38 + int32(272)
	v58 = F_ParseDateTime(m, v42, v38+int32(16), int32(256), v53, v55, v38+int32(496))
	mBase = m.M
	if v58 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+496))
	if v41 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v2137 = v58
	goto L3
L3:
	;
	if v2137 == int32(-1) {
		goto L414
	} else {
		goto L415
	}
L4:
	;
	v67 = int32(_a_F_interval_in_0)
	goto L6
L5:
	;
	v67 = int32(base.Ui32(v41) >> (uint(int32(16)) % 32))
	goto L6
L6:
	;
	v68 = m.G0
	v70 = v68 - int32(160)
	m.G0 = v70
	v73 = v38 + int32(500)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(17)
	v77 = v38 + int32(504)
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v78
	v80 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v80
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v80
	if v61 <= v78 {
		v165 = v7
		goto L10
	} else {
		goto L11
	}
L7:
	;
	m.G0 = v70 + int32(160)
	v2137 = v2094
	goto L3
L8:
	;
	v242 = int32(8)
	v243 = v7
	v246 = v201
	v259 = v7
	v267 = v7
	goto L19
L9:
	;
	v201 = v61 - int32(1)
	v209 = base.B2i32(v91 == int32(45))
	goto L8
L10:
	;
	v182 = v61 - int32(1)
	if v182 < int32(0) {
		v2094 = int32(-1)
		goto L7
	} else {
		goto L18
	}
L11:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[0]))
	if v87 != int32(2) {
		v165 = v7
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if base.B2i32(v91 != int32(45))|base.B2i32(base.Ui32(v61) < base.Ui32(int32(2))) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v104 = int32(1)
	goto L14
L14:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v53+v104<<(uint(int32(2))%32))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	switch v138 - int32(43) {
	case 0, 2:
		v165 = int32(0)
		goto L10
	default:
		goto L16
	}
L15:
	;
	v165 = v141
	goto L10
L16:
	;
	v141 = int32(1)
	v143 = v104 + v141
	if v143 != v61 {
		v104 = v143
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v201 = v182
	v209 = v165
	goto L8
L19:
	;
	v270 = int32(-1)
	v272 = v246 << (uint(int32(2)) % 32)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v55+v272)))
	switch v274 {
	case 0, 2:
		goto L26
	case 1, 6:
		goto L25
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		v2094 = v270
		goto L7
	}
L20:
	;
	v2054 = int32(0)
	v2057 = v2023 | base.B2i32(v2039 == v2054)
	if v2057|base.B2i32(v2047 == v2054) != 0 {
		v2094 = v2054 - v2057
		goto L7
	} else {
		goto L409
	}
L21:
	;
	if int32(0) < v246 {
		v242 = v2022
		v243 = v2023
		v246 = v246 - int32(1)
		v259 = v2039
		v267 = v2047
		goto L19
	} else {
		goto L408
	}
L22:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v70)+112))
	if v2011&v259 != 0 {
		goto L405
	} else {
		goto L406
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_interval_in[1])) = int32(0)
	v993 = v272 + v53
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	v999 = F_strtox_2(m, v994, v70+int32(116), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L165
L24:
	;
	v989 = int32(19)
	goto L23
L25:
	;
	if v243&int32(1) != 0 {
		v2094 = v270
		goto L7
	} else {
		goto L68
	}
L26:
	;
	if v242 != int32(8) {
		v989 = v242
		goto L23
	} else {
		goto L57
	}
L27:
	;
	v368 = v272 + v53
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v371 = v369 + int32(1)
	v372 = int32(58)
	v373 = F___strchrnul(m, v371, v372)
	mBase = m.M
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v375 == v372 {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v272+v53)))
	v281 = F_DecodeTimeCommon(m, v276, v67, v70+int32(112), v70+int32(120))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	if v281 != 0 {
		v2094 = v281
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v285 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v70)+136))
	v290 = int64(3600000000)
	v291 = int64(0)
	v296 = int64(32)
	v299 = int64(base.Ui64(v287) >> (uint(v296) % 64))
	v302 = int64(4294967295)
	v305 = v287 & v302
	v306 = v290 * v305
	v310 = int64(base.Ui64(v306)>>(uint(v296)%64)) + v290*v299
	v317 = v305*v291 + v310&v302
	*(*int64)(unsafe.Add(mBase, uint32(v70)+8)) = v287*v291 + v287>>(uint(int64(63))%64)*v290 + v291*v299 + int64(base.Ui64(v310)>>(uint(v296)%64)) + int64(base.Ui64(v317)>>(uint(v296)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = v306&v302 | v317<<(uint(v296)%64)
	goto L32
L32:
	;
	v328 = int32(-2)
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	if v329 != v330>>(uint(int64(63))%64) {
		v2094 = v328
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v334 = v285 + v330
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v334
	if base.B2i32(v330 < int64(0))^base.B2i32(v334 < v285) != 0 {
		v2094 = v328
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v340 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+128)))
	v342 = v340 * int64(60000000)
	v343 = v334 + v342
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v343
	if base.B2i32(v342 < int64(0))^base.B2i32(v343 < v334) != 0 {
		v2094 = v328
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v349 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+124)))
	v351 = v349 * int64(1000000)
	v352 = v343 + v351
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v352
	if base.B2i32(v351 < int64(0))^base.B2i32(v352 < v343) != 0 {
		v2094 = v328
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v358 = int32(21)
	v359 = int32(0)
	if base.B2i32(v209 == v359)|base.B2i32(v352 <= int64(0)) != 0 {
		v1983 = v358
		v1984 = v359
		v2008 = v267
		goto L22
	} else {
		goto L37
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = int64(0) - v352
	v1983 = v358
	v1984 = v359
	v2008 = v267
	goto L22
L38:
	;
	if v379 == int32(0) {
		goto L26
	} else {
		goto L42
	}
L39:
	;
	v379 = v373
	goto L41
L40:
	;
	v379 = int32(0)
	goto L41
L41:
	;
	goto L38
L42:
	;
	v386 = F_DecodeTimeCommon(m, v371, v67, v70+int32(112), v70+int32(120))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	if v386 != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v388 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v388
	v391 = v70 + int32(96)
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v70)+136))
	v395 = int64(3600000000)
	v396 = int64(0)
	v401 = int64(32)
	v404 = int64(base.Ui64(v392) >> (uint(v401) % 64))
	v407 = int64(4294967295)
	v410 = v392 & v407
	v411 = v395 * v410
	v415 = int64(base.Ui64(v411)>>(uint(v401)%64)) + v395*v404
	v422 = v410*v396 + v415&v407
	*(*int64)(unsafe.Add(mBase, uint32(v391)+8)) = v392*v396 + v392>>(uint(int64(63))%64)*v395 + v396*v404 + int64(base.Ui64(v415)>>(uint(v401)%64)) + int64(base.Ui64(v422)>>(uint(v401)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v391))) = v411&v407 | v422<<(uint(v401)%64)
	goto L45
L45:
	;
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v70)+104))
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v70)+96))
	if v433 != v434>>(uint(int64(63))%64) {
		goto L26
	} else {
		goto L46
	}
L46:
	;
	v438 = v388 + v434
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v438
	if base.B2i32(v434 < int64(0))^base.B2i32(v438 < v388) != 0 {
		goto L26
	} else {
		goto L47
	}
L47:
	;
	v444 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+128)))
	v446 = v444 * int64(60000000)
	v447 = v438 + v446
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v447
	if base.B2i32(v446 < int64(0))^base.B2i32(v447 < v438) != 0 {
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v453 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+124)))
	v455 = v453 * int64(1000000)
	v456 = v447 + v455
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v456
	if base.B2i32(v455 < int64(0))^base.B2i32(v456 < v447) != 0 {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v463 == int32(45) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v456 == int64(-9223372036854775807-1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v472 = v456
	goto L52
L52:
	;
	v473 = int32(21)
	v474 = int32(0)
	if v209&base.B2i32(int64(0) < v472) == v474 {
		v1983 = v473
		v1984 = v474
		v2008 = v267
		goto L22
	} else {
		goto L56
	}
L53:
	;
	v2094 = int32(-2)
	goto L7
L54:
	;
	goto L55
L55:
	;
	v470 = int64(0) - v456
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v470
	v472 = v470
	goto L52
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = int64(0) - v472
	v1983 = v473
	v1984 = v474
	v2008 = v267
	goto L22
L57:
	;
	if base.B2i32(int32(1023) < v67) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v989 = int32(18)
	goto L23
L59:
	;
	if v67 == int32(2048) {
		goto L24
	} else {
		goto L67
	}
L60:
	;
	v989 = int32(20)
	goto L23
L61:
	;
	v989 = int32(21)
	goto L23
L62:
	;
	v989 = int32(23)
	goto L23
L63:
	;
	switch v67 - int32(2) {
	case 0, 4:
		goto L62
	default:
		goto L58
	case 2:
		v989 = int32(25)
		goto L23
	case 6:
		goto L61
	}
L64:
	;
	goto L65
L65:
	;
	switch v67 - int32(1024) {
	case 0, 8:
		goto L60
	case 1, 2, 3, 4, 5, 6, 7:
		goto L58
	default:
		goto L66
	}
L66:
	;
	switch v67 - int32(3072) {
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
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v272+v53)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v272)+uint32(_c_F_interval_in[2])))
	if v502 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v963 = int32(8)
	v966 = v941 & int32(255)
	if v966 == v963 {
		v2022 = v963
		v2023 = int32(0)
		v2039 = v259
		v2047 = v267
		goto L21
	} else {
		goto L160
	}
L70:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v272)+uint32(_c_F_interval_in[3])))
	if v733 != 0 {
		goto L117
	} else {
		goto L118
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+uint32(_c_F_interval_in[2]))) = v663
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+11)))
	if v695 != int32(31) {
		v934 = v663
		v941 = v695
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
	v551 = int32(*(*int8)(unsafe.Add(mBase, uint32(v501))))
	v561 = int32(_a_F_interval_in_1)
	v562 = int32(_a_F_interval_in_2)
	goto L89
L75:
	;
	if v540-v541 == int32(0) {
		v663 = v502
		goto L71
	} else {
		goto L88
	}
L77:
	;
	goto L78
L78:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v509 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v510 = v501
	v511 = v502
	v512 = int32(10)
	v513 = v509
	goto L83
L80:
	;
	v536 = v502
	v540 = int32(0)
	goto L81
L81:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
	goto L75
L82:
	;
	v536 = v531
	v540 = v533
	goto L81
L83:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511))))
	if base.B2i32(v513 != v515)|base.B2i32(v515 == int32(0)) != 0 {
		v531 = v511
		v533 = v513
		goto L82
	} else {
		goto L85
	}
L84:
	;
	v531 = v525
	v533 = int32(0)
	goto L82
L85:
	;
	v521 = v512 - int32(1)
	if v521 == int32(0) {
		v531 = v511
		v533 = v513
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v524 = int32(1)
	v525 = v511 + v524
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)))
	if v526 != 0 {
		v510 = v510 + v524
		v511 = v525
		v512 = v521
		v513 = v526
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
	v594 = v561 + (v562-v561)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v595 = int32(*(*int8)(unsafe.Add(mBase, uint32(v594))))
	v596 = v551 - v595
	if v596 == int32(0) {
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
	v647 = v596
	goto L93
L93:
	;
	v651 = base.B2i32(v647 < int32(0))
	if v647 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L94:
	;
	if v638 == int32(0) {
		v663 = v594
		goto L71
	} else {
		goto L107
	}
L96:
	;
	goto L97
L97:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v605 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v606 = v501
	v607 = v594
	v608 = int32(10)
	v609 = v605
	goto L102
L99:
	;
	v632 = v594
	v636 = int32(0)
	goto L100
L100:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	v638 = v636 - v637
	goto L94
L101:
	;
	v632 = v627
	v636 = v629
	goto L100
L102:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	if base.B2i32(v609 != v611)|base.B2i32(v611 == int32(0)) != 0 {
		v627 = v607
		v629 = v609
		goto L101
	} else {
		goto L104
	}
L103:
	;
	v627 = v621
	v629 = int32(0)
	goto L101
L104:
	;
	v617 = v608 - int32(1)
	if v617 == int32(0) {
		v627 = v607
		v629 = v609
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v620 = int32(1)
	v621 = v607 + v620
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606)+1)))
	if v622 != 0 {
		v606 = v606 + v620
		v607 = v621
		v608 = v617
		v609 = v622
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v647 = v638
	goto L93
L108:
	;
	v652 = v594 - int32(16)
	goto L110
L109:
	;
	v652 = v562
	goto L110
L110:
	;
	if v647 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v655 = v561
	goto L113
L112:
	;
	v655 = v594 + int32(16)
	goto L113
L113:
	;
	if base.Ui32(v655) <= base.Ui32(v652) {
		v561 = v655
		v562 = v652
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
	*(*int32)(unsafe.Add(mBase, uint32(v272)+uint32(_c_F_interval_in[3]))) = v895
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+11)))
	v934 = v895
	v941 = v927
	goto L69
L117:
	;
	goto L122
L118:
	;
	goto L119
L119:
	;
	v782 = int32(*(*int8)(unsafe.Add(mBase, uint32(v501))))
	v792 = int32(_a_F_interval_in_3)
	v793 = int32(_a_F_interval_in_4)
	goto L134
L120:
	;
	if v771-v772 == int32(0) {
		v895 = v733
		goto L116
	} else {
		goto L133
	}
L122:
	;
	goto L123
L123:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v740 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v741 = v501
	v742 = v733
	v743 = int32(10)
	v744 = v740
	goto L128
L125:
	;
	v767 = v733
	v771 = int32(0)
	goto L126
L126:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	goto L120
L127:
	;
	v767 = v762
	v771 = v764
	goto L126
L128:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742))))
	if base.B2i32(v744 != v746)|base.B2i32(v746 == int32(0)) != 0 {
		v762 = v742
		v764 = v744
		goto L127
	} else {
		goto L130
	}
L129:
	;
	v762 = v756
	v764 = int32(0)
	goto L127
L130:
	;
	v752 = v743 - int32(1)
	if v752 == int32(0) {
		v762 = v742
		v764 = v744
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v755 = int32(1)
	v756 = v742 + v755
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741)+1)))
	if v757 != 0 {
		v741 = v741 + v755
		v742 = v756
		v743 = v752
		v744 = v757
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
	v825 = v792 + (v793-v792)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v826 = int32(*(*int8)(unsafe.Add(mBase, uint32(v825))))
	v827 = v782 - v826
	if v827 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v2094 = int32(-1)
	goto L7
L136:
	;
	goto L141
L137:
	;
	v878 = v827
	goto L138
L138:
	;
	v882 = base.B2i32(v878 < int32(0))
	if v878 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L139:
	;
	if v869 == int32(0) {
		v895 = v825
		goto L116
	} else {
		goto L152
	}
L141:
	;
	goto L142
L142:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v836 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v837 = v501
	v838 = v825
	v839 = int32(10)
	v840 = v836
	goto L147
L144:
	;
	v863 = v825
	v867 = int32(0)
	goto L145
L145:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	v869 = v867 - v868
	goto L139
L146:
	;
	v863 = v858
	v867 = v860
	goto L145
L147:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838))))
	if base.B2i32(v840 != v842)|base.B2i32(v842 == int32(0)) != 0 {
		v858 = v838
		v860 = v840
		goto L146
	} else {
		goto L149
	}
L148:
	;
	v858 = v852
	v860 = int32(0)
	goto L146
L149:
	;
	v848 = v839 - int32(1)
	if v848 == int32(0) {
		v858 = v838
		v860 = v840
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v851 = int32(1)
	v852 = v838 + v851
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837)+1)))
	if v853 != 0 {
		v837 = v837 + v851
		v838 = v852
		v839 = v848
		v840 = v853
		goto L147
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v878 = v869
	goto L138
L153:
	;
	v883 = v825 - int32(16)
	goto L155
L154:
	;
	v883 = v793
	goto L155
L155:
	;
	if v878 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v886 = v792
	goto L158
L157:
	;
	v886 = v825 + int32(16)
	goto L158
L158:
	;
	if base.Ui32(v886) <= base.Ui32(v883) {
		v792 = v886
		v793 = v883
		goto L134
	} else {
		goto L159
	}
L159:
	;
	goto L135
L160:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v934)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(0)
	v973 = int32(-1)
	switch v966 {
	case 0:
		goto L161
	default:
		v2094 = v973
		goto L7
	case 17:
		v1983 = v969
		v1984 = int32(1)
		v2008 = v267
		goto L22
	case 19:
		goto L162
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(_a_F_interval_in_5)
	if base.B2i32(v246 != v201)|base.B2i32(base.Ui32(int32(1)) < base.Ui32(v969-int32(9))) != 0 {
		v2094 = v973
		goto L7
	} else {
		goto L164
	}
L162:
	;
	if v246 != v201 {
		v2094 = v973
		goto L7
	} else {
		goto L163
	}
L163:
	;
	v1983 = v969
	v1984 = int32(0)
	v2008 = int32(1)
	goto L22
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v969
	v986 = int32(0)
	v1983 = v986
	v1984 = v986
	v2008 = v267
	goto L22
L165:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v1001 == int32(68) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v2094 = int32(-2)
	goto L7
L167:
	;
	goto L168
L168:
	;
	v1005 = float64(0)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v70)+116))
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006))))
	switch v1007 - int32(45) {
	case 0:
		goto L171
	case 1:
		goto L170
	default:
		goto L172
	}
L169:
	;
	if v209 == int32(0) {
		v1227 = v1210
		v1229 = v1215
		goto L221
	} else {
		goto L222
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+120)) = v1006
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+1)))
	if v1096 == int32(0) {
		v1203 = v1005
		goto L192
	} else {
		goto L193
	}
L171:
	;
	v1018 = F_strtol(m, v1006+int32(1), v70+int32(116), int32(10))
	mBase = m.M
	goto L174
L172:
	;
	if v1007 == int32(0) {
		v1210 = v999
		v1213 = v989
		v1215 = v1005
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v2094 = int32(-1)
	goto L7
L174:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v1020 == int32(68) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v2094 = int32(-2)
	goto L7
L176:
	;
	goto L177
L177:
	;
	if base.Ui32(int32(11)) < base.Ui32(v1018) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v2094 = int32(-2)
	goto L7
L179:
	;
	goto L180
L180:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v70)+116))
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027))))
	if v1028 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v2094 = int32(-1)
	goto L7
L182:
	;
	goto L183
L183:
	;
	v1031 = v70 + int32(80)
	v1034 = int64(12)
	v1035 = int64(0)
	v1040 = int64(32)
	v1043 = int64(base.Ui64(v999) >> (uint(v1040) % 64))
	v1046 = int64(4294967295)
	v1049 = v999 & v1046
	v1050 = v1034 * v1049
	v1054 = int64(base.Ui64(v1050)>>(uint(v1040)%64)) + v1034*v1043
	v1061 = v1049*v1035 + v1054&v1046
	*(*int64)(unsafe.Add(mBase, uint32(v1031)+8)) = v999*v1035 + v999>>(uint(int64(63))%64)*v1034 + v1035*v1043 + int64(base.Ui64(v1054)>>(uint(v1040)%64)) + int64(base.Ui64(v1061)>>(uint(v1040)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1031))) = v1050&v1046 | v1061<<(uint(v1040)%64)
	goto L184
L184:
	;
	v1072 = *(*int64)(unsafe.Add(mBase, uint32(v70)+88))
	v1073 = *(*int64)(unsafe.Add(mBase, uint32(v70)+80))
	if v1072 != v1073>>(uint(int64(63))%64) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v2094 = int32(-2)
	goto L7
L186:
	;
	goto L187
L187:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1081))))
	if v1082 == int32(45) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1085 = int32(0) - v1018
	goto L190
L189:
	;
	v1085 = v1018
	goto L190
L190:
	;
	v1086 = base.I64_extend_i32_s(v1085)
	v1089 = v1073 + v1086
	if base.B2i32(v1086 < int64(0))^base.B2i32(v1089 < v1073) == int32(0) {
		v1210 = v1089
		v1213 = int32(23)
		v1215 = v1005
		goto L169
	} else {
		goto L191
	}
L191:
	;
	v2094 = int32(-2)
	goto L7
L192:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205))))
	if v1206 == int32(45) {
		goto L218
	} else {
		goto L219
	}
L193:
	;
	v1100 = v1006 + int32(1)
	v1101 = int32(_a_F_interval_in_6)
	v1105 = m.G0
	v1107 = v1105 - int32(32)
	v1108 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1107)+24)) = v1108
	*(*int64)(unsafe.Add(mBase, uint32(v1107)+16)) = v1108
	*(*int64)(unsafe.Add(mBase, uint32(v1107)+8)) = v1108
	*(*int64)(unsafe.Add(mBase, uint32(v1107))) = v1108
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[4])))
	if v1116 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v2094 = int32(-1)
	goto L7
L195:
	;
	v1185 = F_strlen(m, v1100)
	mBase = m.M
	if v1184 != v1185 {
		goto L194
	} else {
		goto L214
	}
L196:
	;
	v1184 = int32(0)
	goto L195
L197:
	;
	goto L198
L198:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[5])))
	if v1120 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1124 = v1100
	goto L202
L200:
	;
	goto L201
L201:
	;
	v1134 = v1101
	v1135 = v1116
	goto L205
L202:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	if v1130 == v1116 {
		v1124 = v1124 + int32(1)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v1184 = v1124 - v1100
	goto L195
L204:
	;
	goto L203
L205:
	;
	v1142 = v1107 + int32(base.Ui32(v1135)>>(uint(int32(3))%32))&int32(28)
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)))
	v1144 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1142))) = v1143 | v1144<<(uint(v1135)%32)
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134)+1)))
	if v1148 != 0 {
		v1134 = v1134 + v1144
		v1135 = v1148
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1100))))
	if v1151 == int32(0) {
		v1174 = v1100
		goto L208
	} else {
		goto L209
	}
L207:
	;
	goto L206
L208:
	;
	v1184 = v1174 - v1100
	goto L195
L209:
	;
	v1155 = v1100
	v1156 = v1151
	goto L210
L210:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1107+int32(base.Ui32(v1156)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v1164)>>(uint(v1156)%32))&int32(1) == int32(0) {
		v1174 = v1155
		goto L208
	} else {
		goto L212
	}
L211:
	;
	v1174 = v1172
	goto L208
L212:
	;
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155)+1)))
	v1172 = v1155 + int32(1)
	if v1170 != 0 {
		v1155 = v1172
		v1156 = v1170
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
	v1192 = F_strtod(m, v1006, v70+int32(120))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L29
	} else {
		goto L215
	}
L215:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v70)+120))
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1194))))
	if v1195 != 0 {
		goto L194
	} else {
		goto L216
	}
L216:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v1197 == int32(0) {
		v1203 = v1192
		goto L192
	} else {
		goto L217
	}
L217:
	;
	goto L194
L218:
	;
	v1209 = base.F64_neg(v1203)
	goto L220
L219:
	;
	v1209 = v1203
	goto L220
L220:
	;
	v1210 = v999
	v1213 = v989
	v1215 = v1209
	goto L169
L221:
	;
	switch v1213 - int32(18) {
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
		v2094 = int32(-1)
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
	v1219 = v1210 >> (uint(int64(63)) % 64)
	v1221 = v1219 - (v1210 ^ v1219)
	if base.F64_gt(v1215, float64(0)) == int32(0) {
		v1227 = v1221
		v1229 = v1215
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1227 = v1221
	v1229 = base.F64_neg(v1215)
	goto L221
L224:
	;
	v1983 = int32(21)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L225:
	;
	if base.Ui64(v1227-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L395
	} else {
		goto L396
	}
L226:
	;
	if base.Ui64(v1227-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L385
	} else {
		goto L386
	}
L227:
	;
	if base.Ui64(v1227-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L375
	} else {
		goto L376
	}
L228:
	;
	if base.Ui64(v1227-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L368
	} else {
		goto L369
	}
L229:
	;
	if base.Ui64(v1227-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L350
	} else {
		goto L351
	}
L230:
	;
	if base.Ui64(v1227-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L329
	} else {
		goto L330
	}
L231:
	;
	if base.Ui64(v1227-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L314
	} else {
		goto L315
	}
L232:
	;
	v1542 = v70 - int32(-64)
	v1545 = int64(3600000000)
	v1546 = int64(0)
	v1551 = int64(32)
	v1554 = int64(base.Ui64(v1227) >> (uint(v1551) % 64))
	v1557 = int64(4294967295)
	v1560 = v1227 & v1557
	v1561 = v1545 * v1560
	v1565 = int64(base.Ui64(v1561)>>(uint(v1551)%64)) + v1545*v1554
	v1572 = v1560*v1546 + v1565&v1557
	*(*int64)(unsafe.Add(mBase, uint32(v1542)+8)) = v1227*v1546 + v1227>>(uint(int64(63))%64)*v1545 + v1546*v1554 + int64(base.Ui64(v1565)>>(uint(v1551)%64)) + int64(base.Ui64(v1572)>>(uint(v1551)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1542))) = v1561&v1557 | v1572<<(uint(v1551)%64)
	goto L298
L233:
	;
	v1453 = v70 + int32(48)
	v1456 = int64(60000000)
	v1457 = int64(0)
	v1462 = int64(32)
	v1465 = int64(base.Ui64(v1227) >> (uint(v1462) % 64))
	v1468 = int64(4294967295)
	v1471 = v1227 & v1468
	v1472 = v1456 * v1471
	v1476 = int64(base.Ui64(v1472)>>(uint(v1462)%64)) + v1456*v1465
	v1483 = v1471*v1457 + v1476&v1468
	*(*int64)(unsafe.Add(mBase, uint32(v1453)+8)) = v1227*v1457 + v1227>>(uint(int64(63))%64)*v1456 + v1457*v1465 + int64(base.Ui64(v1476)>>(uint(v1462)%64)) + int64(base.Ui64(v1483)>>(uint(v1462)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1453))) = v1472&v1468 | v1483<<(uint(v1462)%64)
	goto L282
L234:
	;
	v1362 = v70 + int32(32)
	v1365 = int64(1000000)
	v1366 = int64(0)
	v1371 = int64(32)
	v1374 = int64(base.Ui64(v1227) >> (uint(v1371) % 64))
	v1377 = int64(4294967295)
	v1380 = v1227 & v1377
	v1381 = v1365 * v1380
	v1385 = int64(base.Ui64(v1381)>>(uint(v1371)%64)) + v1365*v1374
	v1392 = v1380*v1366 + v1385&v1377
	*(*int64)(unsafe.Add(mBase, uint32(v1362)+8)) = v1227*v1366 + v1227>>(uint(int64(63))%64)*v1365 + v1366*v1374 + int64(base.Ui64(v1385)>>(uint(v1371)%64)) + int64(base.Ui64(v1392)>>(uint(v1371)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1362))) = v1381&v1377 | v1392<<(uint(v1371)%64)
	goto L265
L235:
	;
	v1273 = v70 + int32(16)
	v1276 = int64(1000)
	v1277 = int64(0)
	v1282 = int64(32)
	v1285 = int64(base.Ui64(v1227) >> (uint(v1282) % 64))
	v1288 = int64(4294967295)
	v1291 = v1227 & v1288
	v1292 = v1276 * v1291
	v1296 = int64(base.Ui64(v1292)>>(uint(v1282)%64)) + v1276*v1285
	v1303 = v1291*v1277 + v1296&v1288
	*(*int64)(unsafe.Add(mBase, uint32(v1273)+8)) = v1227*v1277 + v1227>>(uint(int64(63))%64)*v1276 + v1277*v1285 + int64(base.Ui64(v1296)>>(uint(v1282)%64)) + int64(base.Ui64(v1303)>>(uint(v1282)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1273))) = v1292&v1288 | v1303<<(uint(v1282)%64)
	goto L249
L236:
	;
	v1233 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1234 = v1233 + v1227
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1234
	if base.B2i32(v1227 < int64(0))^base.B2i32(v1234 < v1233) != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v2094 = int32(-2)
	goto L7
L238:
	;
	goto L239
L239:
	;
	if base.F64_ne(v1229, float64(0)) != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1243 = base.I64_trunc_sat_f64_s(v1229)
	v1245 = base.F64_sub(v1229, base.F64_convert_i64_s(v1243))
	if base.F64_gt(v1245, float64(0.5)) != 0 {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(_a_F_interval_in_7)
	v1983 = int32(30)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L243:
	;
	v1257 = v1256 + v1234
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1257
	if base.B2i32(v1256 < int64(0))^base.B2i32(v1257 < v1234) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L248
	}
L244:
	;
	v1256 = v1243 + int64(1)
	goto L243
L245:
	;
	goto L246
L246:
	;
	if base.F64_lt(v1245, float64(-0.5)) == int32(0) {
		v1256 = v1243
		goto L243
	} else {
		goto L247
	}
L247:
	;
	v1256 = v1243 - int64(1)
	goto L243
L248:
	;
	goto L242
L249:
	;
	v1314 = *(*int64)(unsafe.Add(mBase, uint32(v70)+24))
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(v70)+16))
	if v1314 != v1315>>(uint(int64(63))%64) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v2094 = int32(-2)
	goto L7
L251:
	;
	goto L252
L252:
	;
	v1320 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1321 = v1320 + v1315
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1321
	if base.B2i32(v1315 < int64(0))^base.B2i32(v1321 < v1320) != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v2094 = int32(-2)
	goto L7
L254:
	;
	goto L255
L255:
	;
	if base.F64_ne(v1229, float64(0)) != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1331 = base.F64_mul(v1229, float64(1000))
	v1332 = base.I64_trunc_sat_f64_s(v1331)
	v1334 = base.F64_sub(v1331, base.F64_convert_i64_s(v1332))
	if base.F64_gt(v1334, float64(0.5)) != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(_a_F_interval_in_8)
	v1983 = int32(29)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L259:
	;
	v1346 = v1345 + v1321
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1346
	if base.B2i32(v1345 < int64(0))^base.B2i32(v1346 < v1321) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L264
	}
L260:
	;
	v1345 = v1332 + int64(1)
	goto L259
L261:
	;
	goto L262
L262:
	;
	if base.F64_lt(v1334, float64(-0.5)) == int32(0) {
		v1345 = v1332
		goto L259
	} else {
		goto L263
	}
L263:
	;
	v1345 = v1332 - int64(1)
	goto L259
L264:
	;
	goto L258
L265:
	;
	v1403 = *(*int64)(unsafe.Add(mBase, uint32(v70)+40))
	v1404 = *(*int64)(unsafe.Add(mBase, uint32(v70)+32))
	if v1403 != v1404>>(uint(int64(63))%64) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v2094 = int32(-2)
	goto L7
L267:
	;
	goto L268
L268:
	;
	v1409 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1410 = v1409 + v1404
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1410
	if base.B2i32(v1404 < int64(0))^base.B2i32(v1410 < v1409) != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v2094 = int32(-2)
	goto L7
L270:
	;
	goto L271
L271:
	;
	if base.F64_ne(v1229, float64(0)) != 0 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1983 = int32(18)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L273:
	;
	v1420 = base.F64_mul(v1229, float64(1e+06))
	v1421 = base.I64_trunc_sat_f64_s(v1420)
	v1423 = base.F64_sub(v1420, base.F64_convert_i64_s(v1421))
	if base.F64_gt(v1423, float64(0.5)) != 0 {
		goto L277
	} else {
		goto L278
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(_a_F_interval_in_9)
	goto L272
L276:
	;
	v1435 = v1434 + v1410
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1435
	if base.B2i32(v1434 < int64(0))^base.B2i32(v1435 < v1410) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L281
	}
L277:
	;
	v1434 = v1421 + int64(1)
	goto L276
L278:
	;
	goto L279
L279:
	;
	if base.F64_lt(v1423, float64(-0.5)) == int32(0) {
		v1434 = v1421
		goto L276
	} else {
		goto L280
	}
L280:
	;
	v1434 = v1421 - int64(1)
	goto L276
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(_a_F_interval_in_10)
	goto L272
L282:
	;
	v1494 = *(*int64)(unsafe.Add(mBase, uint32(v70)+56))
	v1495 = *(*int64)(unsafe.Add(mBase, uint32(v70)+48))
	if v1494 != v1495>>(uint(int64(63))%64) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v2094 = int32(-2)
	goto L7
L284:
	;
	goto L285
L285:
	;
	v1500 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1501 = v1500 + v1495
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1501
	if base.B2i32(v1495 < int64(0))^base.B2i32(v1501 < v1500) != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v2094 = int32(-2)
	goto L7
L287:
	;
	goto L288
L288:
	;
	if base.F64_ne(v1229, float64(0)) != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1511 = base.F64_mul(v1229, float64(6e+07))
	v1512 = base.I64_trunc_sat_f64_s(v1511)
	v1514 = base.F64_sub(v1511, base.F64_convert_i64_s(v1512))
	if base.F64_gt(v1514, float64(0.5)) != 0 {
		goto L293
	} else {
		goto L294
	}
L290:
	;
	goto L291
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(2048)
	v1983 = int32(19)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L292:
	;
	v1526 = v1525 + v1501
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1526
	if base.B2i32(v1525 < int64(0))^base.B2i32(v1526 < v1501) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L297
	}
L293:
	;
	v1525 = v1512 + int64(1)
	goto L292
L294:
	;
	goto L295
L295:
	;
	if base.F64_lt(v1514, float64(-0.5)) == int32(0) {
		v1525 = v1512
		goto L292
	} else {
		goto L296
	}
L296:
	;
	v1525 = v1512 - int64(1)
	goto L292
L297:
	;
	goto L291
L298:
	;
	v1583 = *(*int64)(unsafe.Add(mBase, uint32(v70)+72))
	v1584 = *(*int64)(unsafe.Add(mBase, uint32(v70)+64))
	if v1583 != v1584>>(uint(int64(63))%64) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v2094 = int32(-2)
	goto L7
L300:
	;
	goto L301
L301:
	;
	v1589 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1590 = v1589 + v1584
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1590
	if base.B2i32(v1584 < int64(0))^base.B2i32(v1590 < v1589) != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v2094 = int32(-2)
	goto L7
L303:
	;
	goto L304
L304:
	;
	if base.F64_ne(v1229, float64(0)) != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1600 = base.F64_mul(v1229, float64(3.6e+09))
	v1601 = base.I64_trunc_sat_f64_s(v1600)
	v1603 = base.F64_sub(v1600, base.F64_convert_i64_s(v1601))
	if base.F64_gt(v1603, float64(0.5)) != 0 {
		goto L309
	} else {
		goto L310
	}
L306:
	;
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(1024)
	goto L224
L308:
	;
	v1615 = v1614 + v1590
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1615
	if base.B2i32(v1614 < int64(0))^base.B2i32(v1615 < v1590) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L313
	}
L309:
	;
	v1614 = v1601 + int64(1)
	goto L308
L310:
	;
	goto L311
L311:
	;
	if base.F64_lt(v1603, float64(-0.5)) == int32(0) {
		v1614 = v1601
		goto L308
	} else {
		goto L312
	}
L312:
	;
	v1614 = v1601 - int64(1)
	goto L308
L313:
	;
	goto L307
L314:
	;
	v2094 = int32(-2)
	goto L7
L315:
	;
	goto L316
L316:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v1635 = v1633 + base.I32_wrap_i64(v1227)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v1635
	if base.B2i32(v1635 < v1633)^base.B2i32(v1227 < int64(0)) != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v2094 = int32(-2)
	goto L7
L318:
	;
	goto L319
L319:
	;
	if base.F64_ne(v1229, float64(0)) != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1645 = base.F64_mul(v1229, float64(8.64e+10))
	v1646 = base.I64_trunc_sat_f64_s(v1645)
	v1648 = base.F64_sub(v1645, base.F64_convert_i64_s(v1646))
	if base.F64_gt(v1648, float64(0.5)) != 0 {
		goto L324
	} else {
		goto L325
	}
L321:
	;
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(8)
	goto L224
L323:
	;
	v1660 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1661 = v1660 + v1659
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1661
	if base.B2i32(v1659 < int64(0))^base.B2i32(v1661 < v1660) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L328
	}
L324:
	;
	v1659 = v1646 + int64(1)
	goto L323
L325:
	;
	goto L326
L326:
	;
	if base.F64_lt(v1648, float64(-0.5)) == int32(0) {
		v1659 = v1646
		goto L323
	} else {
		goto L327
	}
L327:
	;
	v1659 = v1646 - int64(1)
	goto L323
L328:
	;
	goto L322
L329:
	;
	v2094 = int32(-2)
	goto L7
L330:
	;
	goto L331
L331:
	;
	v1681 = v1227 * int64(7)
	v1685 = base.I32_wrap_i64(v1681)
	if base.I32_wrap_i64(int64(base.Ui64(v1681)>>(uint(int64(32))%64))) != v1685>>(uint(int32(31))%32) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v2094 = int32(-2)
	goto L7
L333:
	;
	goto L334
L334:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v1691 = v1690 + v1685
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v1691
	if base.B2i32(v1685 < int32(0))^base.B2i32(v1691 < v1690) != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v2094 = int32(-2)
	goto L7
L336:
	;
	goto L337
L337:
	;
	if base.F64_eq(v1229, float64(0)) != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(16777216)
	v1983 = int32(22)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L339:
	;
	v1701 = base.F64_mul(v1229, float64(7))
	v1702 = base.I32_trunc_sat_f64_s(v1701)
	v1703 = v1691 + v1702
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v1703
	if base.B2i32(v1702 < int32(0))^base.B2i32(v1703 < v1691) != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2094 = int32(-2)
	goto L7
L341:
	;
	goto L342
L342:
	;
	v1711 = base.F64_sub(v1701, base.F64_convert_i32_s(v1702))
	if base.F64_eq(v1711, float64(0)) != 0 {
		goto L338
	} else {
		goto L343
	}
L343:
	;
	v1715 = base.F64_mul(v1711, float64(8.64e+10))
	v1716 = base.I64_trunc_sat_f64_s(v1715)
	v1718 = base.F64_sub(v1715, base.F64_convert_i64_s(v1716))
	if base.F64_gt(v1718, float64(0.5)) != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1730 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1731 = v1730 + v1729
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1731
	if base.B2i32(v1729 < int64(0))^base.B2i32(v1731 < v1730) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L349
	}
L345:
	;
	v1729 = v1716 + int64(1)
	goto L344
L346:
	;
	goto L347
L347:
	;
	if base.F64_lt(v1718, float64(-0.5)) == int32(0) {
		v1729 = v1716
		goto L344
	} else {
		goto L348
	}
L348:
	;
	v1729 = v1716 - int64(1)
	goto L344
L349:
	;
	goto L338
L350:
	;
	v2094 = int32(-2)
	goto L7
L351:
	;
	goto L352
L352:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v1756 = v1754 + base.I32_wrap_i64(v1227)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v1756
	if base.B2i32(v1756 < v1754)^base.B2i32(v1227 < int64(0)) != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v2094 = int32(-2)
	goto L7
L354:
	;
	goto L355
L355:
	;
	if base.F64_eq(v1229, float64(0)) != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(2)
	v1983 = int32(23)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L357:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v1767 = base.F64_mul(v1229, float64(30))
	v1768 = base.I32_trunc_sat_f64_s(v1767)
	v1769 = v1765 + v1768
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v1769
	if base.B2i32(v1768 < int32(0))^base.B2i32(v1769 < v1765) != 0 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v2094 = int32(-2)
	goto L7
L359:
	;
	goto L360
L360:
	;
	v1777 = base.F64_sub(v1767, base.F64_convert_i32_s(v1768))
	if base.F64_eq(v1777, float64(0)) != 0 {
		goto L356
	} else {
		goto L361
	}
L361:
	;
	v1781 = base.F64_mul(v1777, float64(8.64e+10))
	v1782 = base.I64_trunc_sat_f64_s(v1781)
	v1784 = base.F64_sub(v1781, base.F64_convert_i64_s(v1782))
	if base.F64_gt(v1784, float64(0.5)) != 0 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v1796 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v1797 = v1796 + v1795
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v1797
	if base.B2i32(v1795 < int64(0))^base.B2i32(v1797 < v1796) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L367
	}
L363:
	;
	v1795 = v1782 + int64(1)
	goto L362
L364:
	;
	goto L365
L365:
	;
	if base.F64_lt(v1784, float64(-0.5)) == int32(0) {
		v1795 = v1782
		goto L362
	} else {
		goto L366
	}
L366:
	;
	v1795 = v1782 - int64(1)
	goto L362
L367:
	;
	goto L356
L368:
	;
	v2094 = int32(-2)
	goto L7
L369:
	;
	goto L370
L370:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v1822 = v1820 + base.I32_wrap_i64(v1227)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v1822
	if base.B2i32(v1822 < v1820)^base.B2i32(v1227 < int64(0)) != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v2094 = int32(-2)
	goto L7
L372:
	;
	goto L373
L373:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v1833 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v1229, float64(12))))
	v1834 = v1829 + v1833
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v1834
	if base.B2i32(v1833 < int32(0))^base.B2i32(v1834 < v1829) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(4)
	v1983 = int32(25)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L375:
	;
	v2094 = int32(-2)
	goto L7
L376:
	;
	goto L377
L377:
	;
	v1851 = v1227 * int64(10)
	v1855 = base.I32_wrap_i64(v1851)
	if base.I32_wrap_i64(int64(base.Ui64(v1851)>>(uint(int64(32))%64))) != v1855>>(uint(int32(31))%32) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v2094 = int32(-2)
	goto L7
L379:
	;
	goto L380
L380:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v1861 = v1860 + v1855
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v1861
	if base.B2i32(v1855 < int32(0))^base.B2i32(v1861 < v1860) != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v2094 = int32(-2)
	goto L7
L382:
	;
	goto L383
L383:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v1874 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_mul(v1229, float64(10)), float64(12))))
	v1875 = v1868 + v1874
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v1875
	if base.B2i32(v1874 < int32(0))^base.B2i32(v1875 < v1868) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L384
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(33554432)
	v1983 = int32(26)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L385:
	;
	v2094 = int32(-2)
	goto L7
L386:
	;
	goto L387
L387:
	;
	v1892 = v1227 * int64(100)
	v1896 = base.I32_wrap_i64(v1892)
	if base.I32_wrap_i64(int64(base.Ui64(v1892)>>(uint(int64(32))%64))) != v1896>>(uint(int32(31))%32) {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2094 = int32(-2)
	goto L7
L389:
	;
	goto L390
L390:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v1902 = v1901 + v1896
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v1902
	if base.B2i32(v1896 < int32(0))^base.B2i32(v1902 < v1901) != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2094 = int32(-2)
	goto L7
L392:
	;
	goto L393
L393:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v1915 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_mul(v1229, float64(100)), float64(12))))
	v1916 = v1909 + v1915
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v1916
	if base.B2i32(v1915 < int32(0))^base.B2i32(v1916 < v1909) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L394
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(67108864)
	v1983 = int32(27)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L395:
	;
	v2094 = int32(-2)
	goto L7
L396:
	;
	goto L397
L397:
	;
	v1933 = v1227 * int64(1000)
	v1937 = base.I32_wrap_i64(v1933)
	if base.I32_wrap_i64(int64(base.Ui64(v1933)>>(uint(int64(32))%64))) != v1937>>(uint(int32(31))%32) {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v2094 = int32(-2)
	goto L7
L399:
	;
	goto L400
L400:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v1943 = v1942 + v1937
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v1943
	if base.B2i32(v1937 < int32(0))^base.B2i32(v1943 < v1942) != 0 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v2094 = int32(-2)
	goto L7
L402:
	;
	goto L403
L403:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v1956 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_mul(v1229, float64(1000)), float64(12))))
	v1957 = v1950 + v1956
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v1957
	if base.B2i32(v1956 < int32(0))^base.B2i32(v1957 < v1950) != 0 {
		v2094 = int32(-2)
		goto L7
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+112)) = int32(134217728)
	v1983 = int32(28)
	v1984 = int32(0)
	v2008 = v267
	goto L22
L405:
	;
	v2094 = int32(-1)
	goto L7
L406:
	;
	goto L407
L407:
	;
	v2022 = v1983
	v2023 = v1984
	v2039 = v2011 | v259
	v2047 = v2008
	goto L21
L408:
	;
	goto L20
L409:
	;
	v2062 = int32(-2)
	v2063 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	if v2063 == int64(-9223372036854775807-1) {
		v2094 = v2062
		goto L7
	} else {
		goto L410
	}
L410:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	if v2066 == int32(-2147483648) {
		v2094 = v2062
		goto L7
	} else {
		goto L411
	}
L411:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v2069 == int32(-2147483648) {
		v2094 = v2062
		goto L7
	} else {
		goto L412
	}
L412:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v2072 == int32(-2147483648) {
		v2094 = v2062
		goto L7
	} else {
		goto L413
	}
L413:
	;
	v2075 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v2075 - v2072
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v2075 - v2069
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v2075 - v2066
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = int64(0) - v2063
	v2094 = v2075
	goto L7
L414:
	;
	v2163 = int32(0)
	v2164 = m.G0
	v2166 = v2164 - int32(112)
	m.G0 = v2166
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(500)))) = int32(17)
	v2173 = v38 + int32(504)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+16)) = v2163
	v2176 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2173)+8)) = v2176
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2176
	v2180 = int32(-1)
	v2181 = F_strlen(m, v42)
	mBase = m.M
	if base.Ui32(v2181) < base.Ui32(int32(2)) {
		v3606 = v2180
		goto L417
	} else {
		goto L418
	}
L415:
	;
	v3649 = v2137
	goto L416
L416:
	;
	if v3649 != 0 {
		goto L753
	} else {
		goto L754
	}
L417:
	;
	m.G0 = v2166 + int32(112)
	v3649 = v3606
	goto L416
L418:
	;
	v2184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v2184 != int32(80) {
		v3606 = v2180
		goto L417
	} else {
		goto L419
	}
L419:
	;
	v2188 = v42 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2166)+108)) = v2188
	v2190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v2190 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	v3606 = v3594
	goto L417
L421:
	;
	v3606 = int32(0)
	goto L417
L422:
	;
	v2204 = v2188
	v2205 = v2190
	v2207 = int32(1)
	v2209 = v2163
	goto L423
L423:
	;
	if v2205&int32(255) == int32(84) {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	goto L421
L425:
	;
	v3551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3543))))
	if v3551 != 0 {
		v2204 = v3543
		v2205 = v3551
		v2207 = v3546
		v2209 = v3547
		goto L423
	} else {
		goto L751
	}
L426:
	;
	v2234 = v2204 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2166)+108)) = v2234
	v2236 = int32(0)
	v3543 = v2234
	v3546 = v2236
	v3547 = v2236
	goto L425
L427:
	;
	goto L428
L428:
	;
	v2238 = int32(-1)
	v2241 = int32(255)
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32((v2205-int32(45))&v2241))&base.B2i32(base.Ui32(int32(10)) <= base.Ui32((v2205-int32(48))&v2241)) != 0 {
		v3594 = v2238
		goto L420
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_interval_in[1])) = int32(0)
	v2257 = F_strtod(m, v2204, v2166+int32(108))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L29
	} else {
		goto L430
	}
L430:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+108))
	if v2259 == v2204 {
		v3594 = v2238
		goto L420
	} else {
		goto L431
	}
L431:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, _c_F_interval_in[1]))
	if v2262 != 0 {
		v3594 = v2238
		goto L420
	} else {
		goto L432
	}
L432:
	;
	v2263 = base.F64_abs(v2257)
	if base.F64_gt(v2263, float64(1e+15)) != 0 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v3606 = int32(-2)
	goto L417
L434:
	;
	goto L435
L435:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v2263)) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v3606 = int32(-2)
	goto L417
L437:
	;
	goto L438
L438:
	;
	v2272 = v2259 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2166)+108)) = v2272
	if base.F64_ge(v2257, float64(0)) != 0 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v2280 = base.F64_floor(v2257)
	goto L441
L440:
	;
	v2280 = base.F64_neg(base.F64_floor(base.F64_neg(v2257)))
	goto L441
L441:
	;
	v2281 = base.I64_trunc_sat_f64_s(v2280)
	*(*int64)(unsafe.Add(mBase, uint32(v2166)+96)) = v2281
	v2284 = base.F64_sub(v2257, base.F64_convert_i64_s(v2281))
	*(*float64)(unsafe.Add(mBase, uint32(v2166)+88)) = v2284
	v2286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2259))))
	if v2207&int32(1) != 0 {
		goto L444
	} else {
		goto L445
	}
L442:
	;
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+12))
	v3529 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v2284, float64(12))))
	v3530 = v3525 + v3529
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+12)) = v3530
	v3532 = int32(1)
	if base.B2i32(v3529 < int32(0))^base.B2i32(v3530 < v3525) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L750
	}
L443:
	;
	v3594 = int32(0)
	goto L420
L444:
	;
	switch v2286 - int32(45) {
	case 0:
		goto L447
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 40, 41, 43:
		v3594 = v2238
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
	switch v2286 - int32(58) {
	case 0:
		goto L618
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 20, 21, 22, 23, 24:
		v3594 = v2238
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
	if v2209 != 0 {
		v3594 = v2238
		goto L420
	} else {
		goto L560
	}
L448:
	;
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+16))
	v2576 = v2574 + base.I32_wrap_i64(v2383)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+16)) = v2576
	if base.B2i32(v2576 < v2574)^base.B2i32(v2383 < int64(0)) != 0 {
		goto L541
	} else {
		goto L542
	}
L449:
	;
	if base.Ui64(v2281-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L526
	} else {
		goto L527
	}
L450:
	;
	if base.Ui64(v2281-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L502
	} else {
		goto L503
	}
L451:
	;
	if base.Ui64(v2281-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L481
	} else {
		goto L482
	}
L452:
	;
	if base.Ui64(v2281-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L477
	} else {
		goto L478
	}
L453:
	;
	v2291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2204))))
	v2294 = v2204 + base.B2i32(v2291 == int32(45))
	v2295 = int32(_a_F_interval_in_6)
	v2299 = m.G0
	v2301 = v2299 - int32(32)
	v2302 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2301)+24)) = v2302
	*(*int64)(unsafe.Add(mBase, uint32(v2301)+16)) = v2302
	*(*int64)(unsafe.Add(mBase, uint32(v2301)+8)) = v2302
	*(*int64)(unsafe.Add(mBase, uint32(v2301))) = v2302
	v2310 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[4])))
	if v2310 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L454:
	;
	if v2286 != 0 {
		v3594 = v2238
		goto L420
	} else {
		goto L455
	}
L455:
	;
	goto L453
L456:
	;
	if base.B2i32(v2378 != int32(8))|v2209 != 0 {
		goto L447
	} else {
		goto L475
	}
L457:
	;
	v2378 = int32(0)
	goto L456
L458:
	;
	goto L459
L459:
	;
	v2314 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[5])))
	if v2314 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v2318 = v2294
	goto L463
L461:
	;
	goto L462
L462:
	;
	v2328 = v2295
	v2329 = v2310
	goto L466
L463:
	;
	v2324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2318))))
	if v2324 == v2310 {
		v2318 = v2318 + int32(1)
		goto L463
	} else {
		goto L465
	}
L464:
	;
	v2378 = v2318 - v2294
	goto L456
L465:
	;
	goto L464
L466:
	;
	v2336 = v2301 + int32(base.Ui32(v2329)>>(uint(int32(3))%32))&int32(28)
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v2336)))
	v2338 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2336))) = v2337 | v2338<<(uint(v2329)%32)
	v2342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328)+1)))
	if v2342 != 0 {
		v2328 = v2328 + v2338
		v2329 = v2342
		goto L466
	} else {
		goto L468
	}
L467:
	;
	v2345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2294))))
	if v2345 == int32(0) {
		v2368 = v2294
		goto L469
	} else {
		goto L470
	}
L468:
	;
	goto L467
L469:
	;
	v2378 = v2368 - v2294
	goto L456
L470:
	;
	v2349 = v2294
	v2350 = v2345
	goto L471
L471:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2301+int32(base.Ui32(v2350)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v2358)>>(uint(v2350)%32))&int32(1) == int32(0) {
		v2368 = v2349
		goto L469
	} else {
		goto L473
	}
L472:
	;
	v2368 = v2366
	goto L469
L473:
	;
	v2364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2349)+1)))
	v2366 = v2349 + int32(1)
	if v2364 != 0 {
		v2349 = v2366
		v2350 = v2364
		goto L471
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	v2383 = base.I64_div_s(v2281, int64(10000))
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v2383-int64(2147483648)) {
		goto L448
	} else {
		goto L476
	}
L476:
	;
	v3606 = int32(-2)
	goto L417
L477:
	;
	v3606 = int32(-2)
	goto L417
L478:
	;
	goto L479
L479:
	;
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+16))
	v2396 = v2394 + base.I32_wrap_i64(v2281)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+16)) = v2396
	if base.B2i32(v2396 < v2394)^base.B2i32(v2281 < int64(0)) == int32(0) {
		goto L442
	} else {
		goto L480
	}
L480:
	;
	v3606 = int32(-2)
	goto L417
L481:
	;
	v3606 = int32(-2)
	goto L417
L482:
	;
	goto L483
L483:
	;
	v2410 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+12))
	v2412 = v2410 + base.I32_wrap_i64(v2281)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+12)) = v2412
	if base.B2i32(v2412 < v2410)^base.B2i32(v2281 < int64(0)) != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v3606 = int32(-2)
	goto L417
L485:
	;
	goto L486
L486:
	;
	v2419 = int32(1)
	if base.F64_eq(v2284, float64(0)) != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2419
	goto L425
L488:
	;
	goto L489
L489:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+8))
	v2425 = base.F64_mul(v2284, float64(30))
	v2426 = base.I32_trunc_sat_f64_s(v2425)
	v2427 = v2423 + v2426
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2427
	if base.B2i32(v2426 < int32(0))^base.B2i32(v2427 < v2423) != 0 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v3606 = int32(-2)
	goto L417
L491:
	;
	goto L492
L492:
	;
	v2435 = base.F64_sub(v2425, base.F64_convert_i32_s(v2426))
	if base.F64_eq(v2435, float64(0)) != 0 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2419
	goto L425
L494:
	;
	goto L495
L495:
	;
	v2440 = base.F64_mul(v2435, float64(8.64e+10))
	v2441 = base.I64_trunc_sat_f64_s(v2440)
	v2443 = base.F64_sub(v2440, base.F64_convert_i64_s(v2441))
	if base.F64_gt(v2443, float64(0.5)) != 0 {
		goto L497
	} else {
		goto L498
	}
L496:
	;
	v2455 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2456 = v2455 + v2454
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2456
	if base.B2i32(v2454 < int64(0))^base.B2i32(v2456 < v2455) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L501
	}
L497:
	;
	v2454 = v2441 + int64(1)
	goto L496
L498:
	;
	goto L499
L499:
	;
	if base.F64_lt(v2443, float64(-0.5)) == int32(0) {
		v2454 = v2441
		goto L496
	} else {
		goto L500
	}
L500:
	;
	v2454 = v2441 - int64(1)
	goto L496
L501:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2419
	goto L425
L502:
	;
	v3606 = int32(-2)
	goto L417
L503:
	;
	goto L504
L504:
	;
	v2470 = v2281 * int64(7)
	v2474 = base.I32_wrap_i64(v2470)
	if base.I32_wrap_i64(int64(base.Ui64(v2470)>>(uint(int64(32))%64))) != v2474>>(uint(int32(31))%32) {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v3606 = int32(-2)
	goto L417
L506:
	;
	goto L507
L507:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+8))
	v2480 = v2479 + v2474
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2480
	if base.B2i32(v2474 < int32(0))^base.B2i32(v2480 < v2479) != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3606 = int32(-2)
	goto L417
L509:
	;
	goto L510
L510:
	;
	v2487 = int32(1)
	if base.F64_eq(v2284, float64(0)) != 0 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2487
	goto L425
L512:
	;
	goto L513
L513:
	;
	v2492 = base.F64_mul(v2284, float64(7))
	v2493 = base.I32_trunc_sat_f64_s(v2492)
	v2494 = v2480 + v2493
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2494
	if base.B2i32(v2493 < int32(0))^base.B2i32(v2494 < v2480) != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v3606 = int32(-2)
	goto L417
L515:
	;
	goto L516
L516:
	;
	v2502 = base.F64_sub(v2492, base.F64_convert_i32_s(v2493))
	if base.F64_eq(v2502, float64(0)) != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2487
	goto L425
L518:
	;
	goto L519
L519:
	;
	v2507 = base.F64_mul(v2502, float64(8.64e+10))
	v2508 = base.I64_trunc_sat_f64_s(v2507)
	v2510 = base.F64_sub(v2507, base.F64_convert_i64_s(v2508))
	if base.F64_gt(v2510, float64(0.5)) != 0 {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	v2522 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2523 = v2522 + v2521
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2523
	if base.B2i32(v2521 < int64(0))^base.B2i32(v2523 < v2522) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L525
	}
L521:
	;
	v2521 = v2508 + int64(1)
	goto L520
L522:
	;
	goto L523
L523:
	;
	if base.F64_lt(v2510, float64(-0.5)) == int32(0) {
		v2521 = v2508
		goto L520
	} else {
		goto L524
	}
L524:
	;
	v2521 = v2508 - int64(1)
	goto L520
L525:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2487
	goto L425
L526:
	;
	v3606 = int32(-2)
	goto L417
L527:
	;
	goto L528
L528:
	;
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+8))
	v2538 = v2536 + base.I32_wrap_i64(v2281)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2538
	if base.B2i32(v2538 < v2536)^base.B2i32(v2281 < int64(0)) != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v3606 = int32(-2)
	goto L417
L530:
	;
	goto L531
L531:
	;
	v2545 = int32(1)
	if base.F64_eq(v2284, float64(0)) != 0 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2545
	goto L425
L533:
	;
	goto L534
L534:
	;
	v2550 = base.F64_mul(v2284, float64(8.64e+10))
	v2551 = base.I64_trunc_sat_f64_s(v2550)
	v2553 = base.F64_sub(v2550, base.F64_convert_i64_s(v2551))
	if base.F64_gt(v2553, float64(0.5)) != 0 {
		goto L536
	} else {
		goto L537
	}
L535:
	;
	v2565 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2566 = v2565 + v2564
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2566
	if base.B2i32(v2564 < int64(0))^base.B2i32(v2566 < v2565) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L540
	}
L536:
	;
	v2564 = v2551 + int64(1)
	goto L535
L537:
	;
	goto L538
L538:
	;
	if base.F64_lt(v2553, float64(-0.5)) == int32(0) {
		v2564 = v2551
		goto L535
	} else {
		goto L539
	}
L539:
	;
	v2564 = v2551 - int64(1)
	goto L535
L540:
	;
	v3543 = v2272
	v3546 = int32(1)
	v3547 = v2545
	goto L425
L541:
	;
	v3606 = int32(-2)
	goto L417
L542:
	;
	goto L543
L543:
	;
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+12))
	v2584 = int64(100)
	v2585 = base.I64_div_s(v2281, v2584)
	v2587 = base.I64_rem_s(v2585, v2584)
	v2589 = v2583 + base.I32_wrap_i64(v2587)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+12)) = v2589
	if base.B2i32(v2589 < v2583)^base.B2i32(v2587 < int64(0)) != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v3606 = int32(-2)
	goto L417
L545:
	;
	goto L546
L546:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+8))
	v2598 = base.I64_rem_s(v2281, int64(100))
	v2600 = v2596 + base.I32_wrap_i64(v2598)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2600
	if base.B2i32(v2600 < v2596)^base.B2i32(v2598 < int64(0)) != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v3606 = int32(-2)
	goto L417
L548:
	;
	goto L549
L549:
	;
	if base.F64_ne(v2284, float64(0)) != 0 {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	v2610 = base.F64_mul(v2284, float64(8.64e+10))
	v2611 = base.I64_trunc_sat_f64_s(v2610)
	v2613 = base.F64_sub(v2610, base.F64_convert_i64_s(v2611))
	if base.F64_gt(v2613, float64(0.5)) != 0 {
		goto L554
	} else {
		goto L555
	}
L551:
	;
	goto L552
L552:
	;
	v2638 = int32(0)
	if v2286 == v2638 {
		goto L443
	} else {
		goto L559
	}
L553:
	;
	v2625 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2626 = v2625 + v2624
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2626
	if base.B2i32(v2624 < int64(0))^base.B2i32(v2626 < v2625) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L558
	}
L554:
	;
	v2624 = v2611 + int64(1)
	goto L553
L555:
	;
	goto L556
L556:
	;
	if base.F64_lt(v2613, float64(-0.5)) == int32(0) {
		v2624 = v2611
		goto L553
	} else {
		goto L557
	}
L557:
	;
	v2624 = v2611 - int64(1)
	goto L553
L558:
	;
	goto L552
L559:
	;
	v3543 = v2272
	v3546 = int32(0)
	v3547 = v2638
	goto L425
L560:
	;
	if base.Ui64(v2281-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v3606 = int32(-2)
	goto L417
L562:
	;
	goto L563
L563:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+16))
	v2649 = v2647 + base.I32_wrap_i64(v2281)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+16)) = v2649
	if base.B2i32(v2649 < v2647)^base.B2i32(v2281 < int64(0)) != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v3606 = int32(-2)
	goto L417
L565:
	;
	goto L566
L566:
	;
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+12))
	v2660 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v2284, float64(12))))
	v2661 = v2656 + v2660
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+12)) = v2661
	if base.B2i32(v2660 < int32(0))^base.B2i32(v2661 < v2656) != 0 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v3606 = int32(-2)
	goto L417
L568:
	;
	goto L569
L569:
	;
	v2668 = int32(0)
	if v2286 == int32(84) {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v3543 = v2272
	v3546 = int32(0)
	v3547 = v2668
	goto L425
L571:
	;
	goto L572
L572:
	;
	if v2286 == int32(0) {
		v3594 = v2286
		goto L420
	} else {
		goto L573
	}
L573:
	;
	v2680 = F_ParseISO8601Number(m, v2272, v2166+int32(108), v2166+int32(96), v2166+int32(88))
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L29
	} else {
		goto L574
	}
L574:
	;
	if v2680 != 0 {
		v3594 = v2680
		goto L420
	} else {
		goto L575
	}
L575:
	;
	v2682 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+96))
	if base.Ui64(v2682-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v3606 = int32(-2)
	goto L417
L577:
	;
	goto L578
L578:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+12))
	v2690 = v2688 + base.I32_wrap_i64(v2682)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+12)) = v2690
	if base.B2i32(v2690 < v2688)^base.B2i32(v2682 < int64(0)) != 0 {
		goto L579
	} else {
		goto L580
	}
L579:
	;
	v3606 = int32(-2)
	goto L417
L580:
	;
	goto L581
L581:
	;
	v2698 = *(*float64)(unsafe.Add(mBase, uint32(v2166)+88))
	if base.F64_eq(v2698, float64(0)) != 0 {
		v2749 = int32(1)
		goto L582
	} else {
		goto L583
	}
L582:
	;
	if v2749 == int32(0) {
		goto L591
	} else {
		goto L592
	}
L583:
	;
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+8))
	v2703 = base.F64_mul(v2698, float64(30))
	v2704 = base.I32_trunc_sat_f64_s(v2703)
	v2705 = v2701 + v2704
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2705
	v2707 = int32(0)
	if base.B2i32(v2704 < v2707)^base.B2i32(v2705 < v2701) != 0 {
		v2749 = v2707
		goto L582
	} else {
		goto L584
	}
L584:
	;
	v2714 = base.F64_sub(v2703, base.F64_convert_i32_s(v2704))
	if base.F64_eq(v2714, float64(0)) != 0 {
		v2749 = int32(1)
		goto L582
	} else {
		goto L585
	}
L585:
	;
	v2718 = base.F64_mul(v2714, float64(8.64e+10))
	v2719 = base.I64_trunc_sat_f64_s(v2718)
	v2721 = base.F64_sub(v2718, base.F64_convert_i64_s(v2719))
	if base.F64_gt(v2721, float64(0.5)) != 0 {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	v2733 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2734 = v2733 + v2732
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2734
	v2749 = base.B2i32(base.B2i32(v2732 < int64(0))^base.B2i32(v2734 < v2733) == int32(0))
	goto L582
L587:
	;
	v2732 = v2719 + int64(1)
	goto L586
L588:
	;
	goto L589
L589:
	;
	if base.F64_lt(v2721, float64(-0.5)) == int32(0) {
		v2732 = v2719
		goto L586
	} else {
		goto L590
	}
L590:
	;
	v2732 = v2719 - int64(1)
	goto L586
L591:
	;
	v3606 = int32(-2)
	goto L417
L592:
	;
	goto L593
L593:
	;
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+108))
	v2754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2753))))
	if v2754 != int32(45) {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	if v2754 == int32(84) {
		v3543 = v2753
		v3546 = int32(0)
		v3547 = v2668
		goto L425
	} else {
		goto L597
	}
L595:
	;
	goto L596
L596:
	;
	v2763 = v2753 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2166)+108)) = v2763
	v2771 = F_ParseISO8601Number(m, v2763, v2166+int32(108), v2166+int32(96), v2166+int32(88))
	mBase = m.M
	v2772 = m.ExcPending
	if v2772 != 0 {
		goto L29
	} else {
		goto L599
	}
L597:
	;
	if v2754 == int32(0) {
		v3594 = v2754
		goto L420
	} else {
		goto L598
	}
L598:
	;
	v3606 = v2180
	goto L417
L599:
	;
	if v2771 != 0 {
		v3594 = v2771
		goto L420
	} else {
		goto L600
	}
L600:
	;
	v2773 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+96))
	if base.Ui64(v2773-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v3606 = int32(-2)
	goto L417
L602:
	;
	goto L603
L603:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+8))
	v2781 = v2779 + base.I32_wrap_i64(v2773)
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+8)) = v2781
	if base.B2i32(v2781 < v2779)^base.B2i32(v2773 < int64(0)) != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v3606 = int32(-2)
	goto L417
L605:
	;
	goto L606
L606:
	;
	v2788 = *(*float64)(unsafe.Add(mBase, uint32(v2166)+88))
	if base.F64_ne(v2788, float64(0)) != 0 {
		goto L607
	} else {
		goto L608
	}
L607:
	;
	v2792 = base.F64_mul(v2788, float64(8.64e+10))
	v2793 = base.I64_trunc_sat_f64_s(v2792)
	v2795 = base.F64_sub(v2792, base.F64_convert_i64_s(v2793))
	if base.F64_gt(v2795, float64(0.5)) != 0 {
		goto L611
	} else {
		goto L612
	}
L608:
	;
	goto L609
L609:
	;
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+108))
	v2822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2821))))
	if v2822 == int32(84) {
		v3543 = v2821
		v3546 = int32(0)
		v3547 = v2668
		goto L425
	} else {
		goto L616
	}
L610:
	;
	v2807 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2808 = v2807 + v2806
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2808
	if base.B2i32(v2806 < int64(0))^base.B2i32(v2808 < v2807) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L615
	}
L611:
	;
	v2806 = v2793 + int64(1)
	goto L610
L612:
	;
	goto L613
L613:
	;
	if base.F64_lt(v2795, float64(-0.5)) == int32(0) {
		v2806 = v2793
		goto L610
	} else {
		goto L614
	}
L614:
	;
	v2806 = v2793 - int64(1)
	goto L610
L615:
	;
	goto L609
L616:
	;
	if v2822 == int32(0) {
		v3594 = v2822
		goto L420
	} else {
		goto L617
	}
L617:
	;
	v3606 = v2180
	goto L417
L618:
	;
	if v2209 != 0 {
		v3606 = v2180
		goto L417
	} else {
		goto L698
	}
L619:
	;
	if v2286 != 0 {
		v3594 = v2238
		goto L420
	} else {
		goto L665
	}
L620:
	;
	v2994 = v2166 + int32(32)
	v2997 = int64(1000000)
	v2998 = int64(0)
	v3003 = int64(32)
	v3006 = int64(base.Ui64(v2281) >> (uint(v3003) % 64))
	v3009 = int64(4294967295)
	v3012 = v2281 & v3009
	v3013 = v2997 * v3012
	v3017 = int64(base.Ui64(v3013)>>(uint(v3003)%64)) + v2997*v3006
	v3024 = v3012*v2998 + v3017&v3009
	*(*int64)(unsafe.Add(mBase, uint32(v2994)+8)) = v2281*v2998 + v2281>>(uint(int64(63))%64)*v2997 + v2998*v3006 + int64(base.Ui64(v3017)>>(uint(v3003)%64)) + int64(base.Ui64(v3024)>>(uint(v3003)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2994))) = v3013&v3009 | v3024<<(uint(v3003)%64)
	goto L651
L621:
	;
	v2911 = v2166 + int32(16)
	v2914 = int64(60000000)
	v2915 = int64(0)
	v2920 = int64(32)
	v2923 = int64(base.Ui64(v2281) >> (uint(v2920) % 64))
	v2926 = int64(4294967295)
	v2929 = v2281 & v2926
	v2930 = v2914 * v2929
	v2934 = int64(base.Ui64(v2930)>>(uint(v2920)%64)) + v2914*v2923
	v2941 = v2929*v2915 + v2934&v2926
	*(*int64)(unsafe.Add(mBase, uint32(v2911)+8)) = v2281*v2915 + v2281>>(uint(int64(63))%64)*v2914 + v2915*v2923 + int64(base.Ui64(v2934)>>(uint(v2920)%64)) + int64(base.Ui64(v2941)>>(uint(v2920)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2911))) = v2930&v2926 | v2941<<(uint(v2920)%64)
	goto L637
L622:
	;
	v2831 = int64(3600000000)
	v2832 = int64(0)
	v2837 = int64(32)
	v2840 = int64(base.Ui64(v2281) >> (uint(v2837) % 64))
	v2843 = int64(4294967295)
	v2846 = v2281 & v2843
	v2847 = v2831 * v2846
	v2851 = int64(base.Ui64(v2847)>>(uint(v2837)%64)) + v2831*v2840
	v2858 = v2846*v2832 + v2851&v2843
	*(*int64)(unsafe.Add(mBase, uint32(v2166)+8)) = v2281*v2832 + v2281>>(uint(int64(63))%64)*v2831 + v2832*v2840 + int64(base.Ui64(v2851)>>(uint(v2837)%64)) + int64(base.Ui64(v2858)>>(uint(v2837)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2166))) = v2847&v2843 | v2858<<(uint(v2837)%64)
	goto L623
L623:
	;
	v2869 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+8))
	v2870 = *(*int64)(unsafe.Add(mBase, uint32(v2166)))
	if v2869 != v2870>>(uint(int64(63))%64) {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v3606 = int32(-2)
	goto L417
L625:
	;
	goto L626
L626:
	;
	v2875 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2876 = v2875 + v2870
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2876
	if base.B2i32(v2870 < int64(0))^base.B2i32(v2876 < v2875) != 0 {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v3606 = int32(-2)
	goto L417
L628:
	;
	goto L629
L629:
	;
	v2883 = int32(0)
	v2884 = int32(1)
	if base.F64_eq(v2284, float64(0)) != 0 {
		v3543 = v2272
		v3546 = v2883
		v3547 = v2884
		goto L425
	} else {
		goto L630
	}
L630:
	;
	v2888 = base.F64_mul(v2284, float64(3.6e+09))
	v2889 = base.I64_trunc_sat_f64_s(v2888)
	v2891 = base.F64_sub(v2888, base.F64_convert_i64_s(v2889))
	if base.F64_gt(v2891, float64(0.5)) != 0 {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	v2903 = v2902 + v2876
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2903
	if base.B2i32(v2902 < int64(0))^base.B2i32(v2903 < v2876) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L636
	}
L632:
	;
	v2902 = v2889 + int64(1)
	goto L631
L633:
	;
	goto L634
L634:
	;
	if base.F64_lt(v2891, float64(-0.5)) == int32(0) {
		v2902 = v2889
		goto L631
	} else {
		goto L635
	}
L635:
	;
	v2902 = v2889 - int64(1)
	goto L631
L636:
	;
	v3543 = v2272
	v3546 = v2883
	v3547 = v2884
	goto L425
L637:
	;
	v2952 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+24))
	v2953 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+16))
	if v2952 != v2953>>(uint(int64(63))%64) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v3606 = int32(-2)
	goto L417
L639:
	;
	goto L640
L640:
	;
	v2958 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v2959 = v2958 + v2953
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2959
	if base.B2i32(v2953 < int64(0))^base.B2i32(v2959 < v2958) != 0 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	v3606 = int32(-2)
	goto L417
L642:
	;
	goto L643
L643:
	;
	v2966 = int32(0)
	v2967 = int32(1)
	if base.F64_eq(v2284, float64(0)) != 0 {
		v3543 = v2272
		v3546 = v2966
		v3547 = v2967
		goto L425
	} else {
		goto L644
	}
L644:
	;
	v2971 = base.F64_mul(v2284, float64(6e+07))
	v2972 = base.I64_trunc_sat_f64_s(v2971)
	v2974 = base.F64_sub(v2971, base.F64_convert_i64_s(v2972))
	if base.F64_gt(v2974, float64(0.5)) != 0 {
		goto L646
	} else {
		goto L647
	}
L645:
	;
	v2986 = v2985 + v2959
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v2986
	if base.B2i32(v2985 < int64(0))^base.B2i32(v2986 < v2959) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L650
	}
L646:
	;
	v2985 = v2972 + int64(1)
	goto L645
L647:
	;
	goto L648
L648:
	;
	if base.F64_lt(v2974, float64(-0.5)) == int32(0) {
		v2985 = v2972
		goto L645
	} else {
		goto L649
	}
L649:
	;
	v2985 = v2972 - int64(1)
	goto L645
L650:
	;
	v3543 = v2272
	v3546 = v2966
	v3547 = v2967
	goto L425
L651:
	;
	v3035 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+40))
	v3036 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+32))
	if v3035 != v3036>>(uint(int64(63))%64) {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v3606 = int32(-2)
	goto L417
L653:
	;
	goto L654
L654:
	;
	v3041 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v3042 = v3041 + v3036
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3042
	if base.B2i32(v3036 < int64(0))^base.B2i32(v3042 < v3041) != 0 {
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v3606 = int32(-2)
	goto L417
L656:
	;
	goto L657
L657:
	;
	v3049 = int32(0)
	v3050 = int32(1)
	if base.F64_eq(v2284, float64(0)) != 0 {
		v3543 = v2272
		v3546 = v3049
		v3547 = v3050
		goto L425
	} else {
		goto L658
	}
L658:
	;
	v3054 = base.F64_mul(v2284, float64(1e+06))
	v3055 = base.I64_trunc_sat_f64_s(v3054)
	v3057 = base.F64_sub(v3054, base.F64_convert_i64_s(v3055))
	if base.F64_gt(v3057, float64(0.5)) != 0 {
		goto L660
	} else {
		goto L661
	}
L659:
	;
	v3069 = v3068 + v3042
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3069
	if base.B2i32(v3068 < int64(0))^base.B2i32(v3069 < v3042) != 0 {
		v3594 = int32(-2)
		goto L420
	} else {
		goto L664
	}
L660:
	;
	v3068 = v3055 + int64(1)
	goto L659
L661:
	;
	goto L662
L662:
	;
	if base.F64_lt(v3057, float64(-0.5)) == int32(0) {
		v3068 = v3055
		goto L659
	} else {
		goto L663
	}
L663:
	;
	v3068 = v3055 - int64(1)
	goto L659
L664:
	;
	v3543 = v2272
	v3546 = v3049
	v3547 = v3050
	goto L425
L665:
	;
	v3076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2204))))
	v3079 = v2204 + base.B2i32(v3076 == int32(45))
	v3080 = int32(_a_F_interval_in_6)
	v3084 = m.G0
	v3086 = v3084 - int32(32)
	v3087 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3086)+24)) = v3087
	*(*int64)(unsafe.Add(mBase, uint32(v3086)+16)) = v3087
	*(*int64)(unsafe.Add(mBase, uint32(v3086)+8)) = v3087
	*(*int64)(unsafe.Add(mBase, uint32(v3086))) = v3087
	v3095 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[4])))
	if v3095 == int32(0) {
		goto L667
	} else {
		goto L668
	}
L666:
	;
	if base.B2i32(v3163 != int32(6))|v2209 != 0 {
		goto L618
	} else {
		goto L685
	}
L667:
	;
	v3163 = int32(0)
	goto L666
L668:
	;
	goto L669
L669:
	;
	v3099 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_in[5])))
	if v3099 == int32(0) {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v3103 = v3079
	goto L673
L671:
	;
	goto L672
L672:
	;
	v3113 = v3080
	v3114 = v3095
	goto L676
L673:
	;
	v3109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3103))))
	if v3109 == v3095 {
		v3103 = v3103 + int32(1)
		goto L673
	} else {
		goto L675
	}
L674:
	;
	v3163 = v3103 - v3079
	goto L666
L675:
	;
	goto L674
L676:
	;
	v3121 = v3086 + int32(base.Ui32(v3114)>>(uint(int32(3))%32))&int32(28)
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v3121)))
	v3123 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3121))) = v3122 | v3123<<(uint(v3114)%32)
	v3127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3113)+1)))
	if v3127 != 0 {
		v3113 = v3113 + v3123
		v3114 = v3127
		goto L676
	} else {
		goto L678
	}
L677:
	;
	v3130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3079))))
	if v3130 == int32(0) {
		v3153 = v3079
		goto L679
	} else {
		goto L680
	}
L678:
	;
	goto L677
L679:
	;
	v3163 = v3153 - v3079
	goto L666
L680:
	;
	v3134 = v3079
	v3135 = v3130
	goto L681
L681:
	;
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(v3086+int32(base.Ui32(v3135)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3143)>>(uint(v3135)%32))&int32(1) == int32(0) {
		v3153 = v3134
		goto L679
	} else {
		goto L683
	}
L682:
	;
	v3153 = v3151
	goto L679
L683:
	;
	v3149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3134)+1)))
	v3151 = v3134 + int32(1)
	if v3149 != 0 {
		v3134 = v3151
		v3135 = v3149
		goto L681
	} else {
		goto L684
	}
L684:
	;
	goto L682
L685:
	;
	v3168 = v2166 - int32(-64)
	v3170 = base.I64_div_s(v2281, int64(10000))
	v3173 = int64(3600000000)
	v3174 = int64(0)
	v3179 = int64(32)
	v3182 = int64(base.Ui64(v3170) >> (uint(v3179) % 64))
	v3185 = int64(4294967295)
	v3188 = v3170 & v3185
	v3189 = v3173 * v3188
	v3193 = int64(base.Ui64(v3189)>>(uint(v3179)%64)) + v3173*v3182
	v3200 = v3188*v3174 + v3193&v3185
	*(*int64)(unsafe.Add(mBase, uint32(v3168)+8)) = v3170*v3174 + v3170>>(uint(int64(63))%64)*v3173 + v3174*v3182 + int64(base.Ui64(v3193)>>(uint(v3179)%64)) + int64(base.Ui64(v3200)>>(uint(v3179)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3168))) = v3189&v3185 | v3200<<(uint(v3179)%64)
	goto L686
L686:
	;
	v3211 = int32(-2)
	v3212 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+72))
	v3213 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+64))
	if v3212 != v3213>>(uint(int64(63))%64) {
		v3606 = v3211
		goto L417
	} else {
		goto L687
	}
L687:
	;
	v3217 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v3218 = v3217 + v3213
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3218
	if base.B2i32(v3213 < int64(0))^base.B2i32(v3218 < v3217) != 0 {
		v3606 = v3211
		goto L417
	} else {
		goto L688
	}
L688:
	;
	v3224 = int64(100)
	v3225 = base.I64_div_s(v2281, v3224)
	v3227 = base.I64_rem_s(v3225, v3224)
	v3229 = v3227 * int64(60000000)
	v3230 = v3218 + v3229
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3230
	if base.B2i32(v3229 < int64(0))^base.B2i32(v3230 < v3218) != 0 {
		v3606 = v3211
		goto L417
	} else {
		goto L689
	}
L689:
	;
	v3240 = (v2281 - v3225*int64(100)) * int64(1000000)
	v3241 = v3230 + v3240
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3241
	if base.B2i32(v3240 < int64(0))^base.B2i32(v3241 < v3230) != 0 {
		v3606 = v3211
		goto L417
	} else {
		goto L690
	}
L690:
	;
	if base.F64_eq(v2284, float64(0)) != 0 {
		goto L421
	} else {
		goto L691
	}
L691:
	;
	v3249 = base.I64_trunc_sat_f64_s(v2284)
	v3251 = base.F64_sub(v2284, base.F64_convert_i64_s(v3249))
	if base.F64_gt(v3251, float64(0.5)) != 0 {
		goto L693
	} else {
		goto L694
	}
L692:
	;
	v3263 = v3262 + v3241
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3263
	if base.B2i32(v3262 < int64(0))^base.B2i32(v3263 < v3241) == int32(0) {
		goto L421
	} else {
		goto L697
	}
L693:
	;
	v3262 = v3249 + int64(1)
	goto L692
L694:
	;
	goto L695
L695:
	;
	if base.F64_lt(v3251, float64(-0.5)) == int32(0) {
		v3262 = v3249
		goto L692
	} else {
		goto L696
	}
L696:
	;
	v3262 = v3249 - int64(1)
	goto L692
L697:
	;
	v3606 = v3211
	goto L417
L698:
	;
	v3272 = v2166 + int32(48)
	v3275 = int64(3600000000)
	v3276 = int64(0)
	v3281 = int64(32)
	v3284 = int64(base.Ui64(v2281) >> (uint(v3281) % 64))
	v3287 = int64(4294967295)
	v3290 = v2281 & v3287
	v3291 = v3275 * v3290
	v3295 = int64(base.Ui64(v3291)>>(uint(v3281)%64)) + v3275*v3284
	v3302 = v3290*v3276 + v3295&v3287
	*(*int64)(unsafe.Add(mBase, uint32(v3272)+8)) = v2281*v3276 + v2281>>(uint(int64(63))%64)*v3275 + v3276*v3284 + int64(base.Ui64(v3295)>>(uint(v3281)%64)) + int64(base.Ui64(v3302)>>(uint(v3281)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3272))) = v3291&v3287 | v3302<<(uint(v3281)%64)
	goto L699
L699:
	;
	v3313 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+56))
	v3314 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+48))
	if v3313 != v3314>>(uint(int64(63))%64) {
		goto L700
	} else {
		goto L701
	}
L700:
	;
	v3606 = int32(-2)
	goto L417
L701:
	;
	goto L702
L702:
	;
	v3319 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v3320 = v3319 + v3314
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3320
	if base.B2i32(v3314 < int64(0))^base.B2i32(v3320 < v3319) != 0 {
		goto L703
	} else {
		goto L704
	}
L703:
	;
	v3606 = int32(-2)
	goto L417
L704:
	;
	goto L705
L705:
	;
	if base.F64_eq(v2284, float64(0)) != 0 {
		goto L706
	} else {
		goto L707
	}
L706:
	;
	if v2286 == int32(0) {
		goto L421
	} else {
		goto L714
	}
L707:
	;
	v3330 = base.F64_mul(v2284, float64(3.6e+09))
	v3331 = base.I64_trunc_sat_f64_s(v3330)
	v3333 = base.F64_sub(v3330, base.F64_convert_i64_s(v3331))
	if base.F64_gt(v3333, float64(0.5)) != 0 {
		goto L709
	} else {
		goto L710
	}
L708:
	;
	v3345 = v3344 + v3320
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3345
	if base.B2i32(v3344 < int64(0))^base.B2i32(v3345 < v3320) == int32(0) {
		goto L706
	} else {
		goto L713
	}
L709:
	;
	v3344 = v3331 + int64(1)
	goto L708
L710:
	;
	goto L711
L711:
	;
	if base.F64_lt(v3333, float64(-0.5)) == int32(0) {
		v3344 = v3331
		goto L708
	} else {
		goto L712
	}
L712:
	;
	v3344 = v3331 - int64(1)
	goto L708
L713:
	;
	v3606 = int32(-2)
	goto L417
L714:
	;
	v3365 = F_ParseISO8601Number(m, v2272, v2166+int32(108), v2166+int32(96), v2166+int32(88))
	mBase = m.M
	v3366 = m.ExcPending
	if v3366 != 0 {
		goto L29
	} else {
		goto L715
	}
L715:
	;
	if v3365 != 0 {
		v3606 = v3365
		goto L417
	} else {
		goto L716
	}
L716:
	;
	v3367 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+96))
	v3368 = *(*float64)(unsafe.Add(mBase, uint32(v2166)+88))
	v3369 = int64(60000000)
	v3370 = int32(0)
	v3374 = m.G0
	v3376 = v3374 - int32(16)
	m.G0 = v3376
	v3378 = int64(63)
	F___multi3(m, v3376, v3367, v3367>>(uint(v3378)%64), v3369, int64(0))
	mBase = m.M
	v3383 = *(*int64)(unsafe.Add(mBase, uint32(v3376)+8))
	v3384 = *(*int64)(unsafe.Add(mBase, uint32(v3376)))
	if v3383 != v3384>>(uint(v3378)%64) {
		v3425 = v3370
		goto L718
	} else {
		goto L719
	}
L717:
	;
	if v3425 == int32(0) {
		goto L727
	} else {
		goto L728
	}
L718:
	;
	m.G0 = v3376 + int32(16)
	goto L717
L719:
	;
	v3388 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v3389 = v3388 + v3384
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3389
	if base.B2i32(v3384 < int64(0))^base.B2i32(v3389 < v3388) != 0 {
		v3425 = v3370
		goto L718
	} else {
		goto L720
	}
L720:
	;
	if base.F64_eq(v3368, float64(0)) != 0 {
		v3425 = int32(1)
		goto L718
	} else {
		goto L721
	}
L721:
	;
	v3399 = base.F64_mul(v3368, base.F64_convert_i64_u(v3369))
	v3400 = base.I64_trunc_sat_f64_s(v3399)
	v3402 = base.F64_sub(v3399, base.F64_convert_i64_s(v3400))
	if base.F64_gt(v3402, float64(0.5)) != 0 {
		goto L723
	} else {
		goto L724
	}
L722:
	;
	v3414 = v3389 + v3413
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3414
	v3425 = base.B2i32(base.B2i32(v3413 < int64(0))^base.B2i32(v3414 < v3389) == int32(0))
	goto L718
L723:
	;
	v3413 = v3400 + int64(1)
	goto L722
L724:
	;
	goto L725
L725:
	;
	if base.F64_lt(v3402, float64(-0.5)) == int32(0) {
		v3413 = v3400
		goto L722
	} else {
		goto L726
	}
L726:
	;
	v3413 = v3400 - int64(1)
	goto L722
L727:
	;
	v3606 = int32(-2)
	goto L417
L728:
	;
	goto L729
L729:
	;
	v3434 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+108))
	v3435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3434))))
	if v3435 == int32(0) {
		v3606 = v3435
		goto L417
	} else {
		goto L730
	}
L730:
	;
	if v3435 != int32(58) {
		goto L731
	} else {
		goto L732
	}
L731:
	;
	v3606 = int32(-1)
	goto L417
L732:
	;
	goto L733
L733:
	;
	v3442 = v3434 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2166)+108)) = v3442
	v3450 = F_ParseISO8601Number(m, v3442, v2166+int32(108), v2166+int32(96), v2166+int32(88))
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L29
	} else {
		goto L734
	}
L734:
	;
	if v3450 != 0 {
		v3606 = v3450
		goto L417
	} else {
		goto L735
	}
L735:
	;
	v3453 = *(*int64)(unsafe.Add(mBase, uint32(v2166)+96))
	v3454 = *(*float64)(unsafe.Add(mBase, uint32(v2166)+88))
	v3455 = int64(1000000)
	v3456 = int32(0)
	v3460 = m.G0
	v3462 = v3460 - int32(16)
	m.G0 = v3462
	v3464 = int64(63)
	F___multi3(m, v3462, v3453, v3453>>(uint(v3464)%64), v3455, int64(0))
	mBase = m.M
	v3469 = *(*int64)(unsafe.Add(mBase, uint32(v3462)+8))
	v3470 = *(*int64)(unsafe.Add(mBase, uint32(v3462)))
	if v3469 != v3470>>(uint(v3464)%64) {
		v3511 = v3456
		goto L737
	} else {
		goto L738
	}
L736:
	;
	if v3511 == int32(0) {
		v3606 = int32(-2)
		goto L417
	} else {
		goto L746
	}
L737:
	;
	m.G0 = v3462 + int32(16)
	goto L736
L738:
	;
	v3474 = *(*int64)(unsafe.Add(mBase, uint32(v2173)))
	v3475 = v3474 + v3470
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3475
	if base.B2i32(v3470 < int64(0))^base.B2i32(v3475 < v3474) != 0 {
		v3511 = v3456
		goto L737
	} else {
		goto L739
	}
L739:
	;
	if base.F64_eq(v3454, float64(0)) != 0 {
		v3511 = int32(1)
		goto L737
	} else {
		goto L740
	}
L740:
	;
	v3485 = base.F64_mul(v3454, base.F64_convert_i64_u(v3455))
	v3486 = base.I64_trunc_sat_f64_s(v3485)
	v3488 = base.F64_sub(v3485, base.F64_convert_i64_s(v3486))
	if base.F64_gt(v3488, float64(0.5)) != 0 {
		goto L742
	} else {
		goto L743
	}
L741:
	;
	v3500 = v3475 + v3499
	*(*int64)(unsafe.Add(mBase, uint32(v2173))) = v3500
	v3511 = base.B2i32(base.B2i32(v3499 < int64(0))^base.B2i32(v3500 < v3475) == int32(0))
	goto L737
L742:
	;
	v3499 = v3486 + int64(1)
	goto L741
L743:
	;
	goto L744
L744:
	;
	if base.F64_lt(v3488, float64(-0.5)) == int32(0) {
		v3499 = v3486
		goto L741
	} else {
		goto L745
	}
L745:
	;
	v3499 = v3486 - int64(1)
	goto L741
L746:
	;
	v3521 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+108))
	v3522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3521))))
	if v3522 != 0 {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	v3523 = int32(-1)
	goto L749
L748:
	;
	v3523 = int32(0)
	goto L749
L749:
	;
	v3606 = v3523
	goto L417
L750:
	;
	v3543 = v2272
	v3546 = v3532
	v3547 = v3532
	goto L425
L751:
	;
	goto L424
L752:
	;
	m.G0 = v38 + int32(528)
	return v3749
L753:
	;
	if v3649 == int32(-2) {
		goto L756
	} else {
		goto L757
	}
L754:
	;
	goto L755
L755:
	;
	v3686 = F_palloc(m, int32(16))
	mBase = m.M
	v3687 = m.ExcPending
	if v3687 != 0 {
		goto L29
	} else {
		goto L760
	}
L756:
	;
	v3676 = int32(-4)
	goto L758
L757:
	;
	v3676 = v3649
	goto L758
L758:
	;
	F_DateTimeParseError(m, v3676, v38+int32(8), v42, int32(_a_F_interval_in_11), v40)
	mBase = m.M
	v3681 = m.ExcPending
	if v3681 != 0 {
		goto L29
	} else {
		goto L759
	}
L759:
	;
	v3682 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v3682)
	v3749 = int32(0)
	goto L752
L760:
	;
	v3688 = *(*int32)(unsafe.Add(mBase, uint32(v38)+500))
	switch v3688 - int32(9) {
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
	F_AdjustIntervalForTypmod(m, v3686, v41, v40)
	mBase = m.M
	v3747 = m.ExcPending
	if v3747 != 0 {
		goto L29
	} else {
		goto L777
	}
L762:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3686)+8)) = int64(9223372034707292159)
	*(*int64)(unsafe.Add(mBase, uint32(v3686))) = int64(9223372036854775807)
	goto L761
L763:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3729 = m.ExcPending
	if v3729 != 0 {
		goto L29
	} else {
		goto L774
	}
L764:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3686)+8)) = int64(-9223372034707292160)
	*(*int64)(unsafe.Add(mBase, uint32(v3686))) = int64(-9223372036854775807 - 1)
	goto L761
L765:
	;
	v3691 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+516)))
	v3692 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+520)))
	v3695 = v3691 + v3692*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v3695-int64(2147483648)) {
		goto L766
	} else {
		goto L767
	}
L766:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3686)+12)) = uint32(v3695)
	v3701 = *(*int32)(unsafe.Add(mBase, uint32(v38)+512))
	*(*int32)(unsafe.Add(mBase, uint32(v3686)+8)) = v3701
	v3703 = *(*int64)(unsafe.Add(mBase, uint32(v38)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v3686))) = v3703
	goto L761
L767:
	;
	goto L768
L768:
	;
	v3705 = int32(0)
	v3706 = F_errsave_start(m, v40)
	mBase = m.M
	v3707 = m.ExcPending
	if v3707 != 0 {
		goto L29
	} else {
		goto L769
	}
L769:
	;
	if v3706 == int32(0) {
		v3749 = v3705
		goto L752
	} else {
		goto L770
	}
L770:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v3712 = m.ExcPending
	if v3712 != 0 {
		goto L29
	} else {
		goto L771
	}
L771:
	;
	F_errmsg(m, int32(_a_F_interval_in_12), int32(0))
	mBase = m.M
	v3716 = m.ExcPending
	if v3716 != 0 {
		goto L29
	} else {
		goto L772
	}
L772:
	;
	F_errsave_finish(m, v40, int32(_a_F_interval_in_13), int32(948), int32(_a_F_interval_in_14))
	mBase = m.M
	v3721 = m.ExcPending
	if v3721 != 0 {
		goto L29
	} else {
		goto L773
	}
L773:
	;
	v3749 = v3705
	goto L752
L774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v42
	v3731 = *(*int32)(unsafe.Add(mBase, uint32(v38)+500))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v3731
	F_errmsg_internal(m, int32(_a_F_interval_in_15), v38)
	mBase = m.M
	v3735 = m.ExcPending
	if v3735 != 0 {
		goto L29
	} else {
		goto L775
	}
L775:
	;
	F_errfinish(m, int32(_a_F_interval_in_13), int32(961), int32(_a_F_interval_in_14))
	mBase = m.M
	v3740 = m.ExcPending
	if v3740 != 0 {
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
	v3749 = v3686
	goto L752
}
func F_interval_lerp(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	v16 = F_DirectFunctionCall2Coll(m, int32(1460), v4, l1, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v22 = F_DirectFunctionCall2Coll(m, int32(1459), v4, v16, v7+int32(8))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_DirectFunctionCall2Coll(m, int32(1458), v4, v22, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v24
			}
		}
	}
}
func F_interval_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v53 int64
	_ = v53
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
	v17 = v14 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+12)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	v23 = v19*int64(30) + v22
	v32 = int64(32)
	v33 = int64(20)
	v35 = int64(base.Ui64(v23) >> (uint(v32) % 64))
	v38 = int64(4294967295)
	v39 = int64(500654080)
	v41 = v23 & v38
	v42 = v39 * v41
	v46 = int64(base.Ui64(v42)>>(uint(v32)%64)) + v39*v35
	v53 = v41*v33 + v46&v38
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v23*int64(0) + v23>>(uint(int64(63))%64)*int64(86400000000) + v33*v35 + int64(base.Ui64(v46)>>(uint(v32)%64)) + int64(base.Ui64(v53)>>(uint(v32)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v42&v38 | v53<<(uint(v32)%64)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	return base.B2i32(v119^v120|(base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115)))+(v114+v113>>(uint(v124)%64))^(base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112)))+(v111+v110>>(uint(v124)%64)))) != int64(0))
}
func F_interval_part_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 float64
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v310 int64
	_ = v310
	var v315 int64
	_ = v315
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v324 int64
	_ = v324
	var v327 int64
	_ = v327
	var v330 int64
	_ = v330
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v338 int64
	_ = v338
	var v345 int64
	_ = v345
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v364 int64
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v394 int64
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v430 int64
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v446 int32
	_ = v446
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(1)
	v25 = v20 + v24
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v30 = v28 & v24
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = v25
	goto L5
L4:
	;
	v31 = v20 + int32(4)
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v28 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v61 = F_downcase_truncate_identifier(m, v31, v59, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v38 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v49 = int32(1)
	if v30 != 0 {
		v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v41 = int32(16)
	goto L12
L11:
	;
	v41 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v38-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = int32(4)
	goto L15
L14:
	;
	v48 = v41
	goto L15
L15:
	;
	v59 = v48
	goto L6
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v64 = v17 + int32(76)
	v68 = Fn13825(m, v61, v64, int32(_a_F_interval_part_common_0), int32(_a_F_interval_part_common_1), int32(_a_F_interval_part_common_2))
	mBase = m.M
	goto L18
L18:
	;
	if v68 == int32(31) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v74 = Fn13825(m, v61, v64, int32(_a_F_interval_part_common_3), int32(_a_F_interval_part_common_4), int32(_a_F_interval_part_common_5))
	mBase = m.M
	goto L22
L20:
	;
	v75 = v68
	goto L21
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v76 != int32(2147483647) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v75 = v74
	goto L21
L23:
	;
	m.G0 = v17 + int32(80)
	return v446
L24:
	;
	if v75 == int32(17) {
		goto L65
	} else {
		goto L66
	}
L25:
	;
	v96 = int32(0)
	if base.B2i32(v75 == v96)|base.B2i32(v75 == int32(17)) == v96 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	if v76 != int32(-2147483648) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v88 != int32(2147483647) {
		goto L24
	} else {
		goto L32
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v81 != int32(-2147483648) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	if v84 != int64(-9223372036854775807-1) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v95 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L25
L32:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	if v91 != int64(9223372036854775807) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v95 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L25
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if base.Ui32(v125) <= base.Ui32(int32(30)) {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v111 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v61
	F_errmsg(m, int32(_a_F_interval_part_common_11), v17+int32(48))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_13), int32(_a_F_interval_part_common_14))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
	v446 = int32(0)
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
	v129 = int32(1) << (uint(v125) % 32)
	if v129&int32(506464256) != 0 {
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
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	if v129&int32(1640759296) != 0 {
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
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v143 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v61
	F_errmsg(m, int32(_a_F_interval_part_common_7), v17-int32(-64))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_17), int32(_a_F_interval_part_common_14))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	if base.F64_lt(v95, float64(0)) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v173 = F_Float8GetDatum(m, v95)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L62
	}
L57:
	;
	v160 = int32(0)
	v164 = F_DirectFunctionCall3Coll(m, int32(408), v160, int32(_a_F_interval_part_common_15), v160, int32(-1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v167 = int32(0)
	v171 = F_DirectFunctionCall3Coll(m, int32(408), v167, int32(_a_F_interval_part_common_16), v167, int32(-1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v446 = v164
	goto L23
L61:
	;
	v446 = v171
	goto L23
L62:
	;
	v446 = v173
	goto L23
L63:
	;
	if l1 != 0 {
		goto L118
	} else {
		goto L119
	}
L64:
	;
	v430 = base.I64_extend32_s(v195) + base.I64_extend32_s(v192)*int64(1000000)
	goto L63
L65:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v182 = base.I64_div_s(v180, int64(3600000000))
	v185 = v182*int64(-3600000000) + v180
	v187 = base.I64_div_s(v185, int64(60000000))
	v190 = v187*int64(-60000000) + v185
	v192 = base.I64_div_s(v190, int64(1000000))
	v195 = v192*int64(4293967296) + v190
	v196 = base.I32_wrap_i64(v195)
	v197 = base.I32_wrap_i64(v192)
	v199 = base.I32_rem_s(v76, int32(12))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	switch v201 - int32(18) {
	case 0:
		goto L78
	case 1:
		goto L77
	case 2:
		v430 = v182
		goto L63
	case 3:
		goto L76
	case 4:
		goto L75
	case 5:
		goto L74
	case 6:
		goto L73
	case 7:
		goto L72
	case 8:
		goto L71
	case 9:
		goto L70
	case 10:
		goto L69
	case 11:
		goto L79
	case 12:
		goto L64
	default:
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if v75 != 0 {
		goto L98
	} else {
		goto L99
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L93
	}
L69:
	;
	v272 = base.I32_div_s(v76, int32(_a_F_interval_part_common_6))
	v430 = base.I64_extend_i32_s(v272)
	goto L63
L70:
	;
	v269 = base.I32_div_s(v76, int32(1200))
	v430 = base.I64_extend_i32_s(v269)
	goto L63
L71:
	;
	v266 = base.I32_div_s(v76, int32(120))
	v430 = base.I64_extend_i32_s(v266)
	goto L63
L72:
	;
	v263 = base.I32_div_s(v76, int32(12))
	v430 = base.I64_extend_i32_s(v263)
	goto L63
L73:
	;
	if int32(0) <= v76 {
		goto L90
	} else {
		goto L91
	}
L74:
	;
	v430 = base.I64_extend_i32_s(v199)
	goto L63
L75:
	;
	v239 = base.I32_div_s(v200, int32(7))
	v430 = base.I64_extend_i32_s(v239)
	goto L63
L76:
	;
	v430 = base.I64_extend_i32_s(v200)
	goto L63
L77:
	;
	v430 = base.I64_extend32_s(v187)
	goto L63
L78:
	;
	if l1 != 0 {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	if l1 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v210 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v195)+base.I64_extend32_s(v192)*int64(1000000), int32(3))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v213 = float64(1000)
	v219 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v197), v213), base.F64_div(base.F64_convert_i32_s(v196), v213)))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	v446 = v210
	goto L23
L84:
	;
	v446 = v219
	goto L23
L85:
	;
	v227 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v195)+base.I64_extend32_s(v192)*int64(1000000), int32(6))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v234 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v196), float64(1e+06)), base.F64_convert_i32_s(v197)))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	v446 = v227
	goto L23
L89:
	;
	v446 = v234
	goto L23
L90:
	;
	v247 = base.I32_div_u_s(v199&int32(255), int32(3))
	v430 = base.I64_extend_i32_u(v247 + int32(1))
	goto L63
L91:
	;
	goto L92
L92:
	;
	v254 = base.I32_rem_s(int32(0)-v76, int32(12))
	v257 = base.I32_div_s(base.I32_extend8_s(v254), int32(-3))
	v430 = base.I64_extend8_s(base.I64_extend_i32_u(v257 - int32(1)))
	goto L63
L93:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v282 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v61
	F_errmsg(m, int32(_a_F_interval_part_common_7), v17)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_9), int32(_a_F_interval_part_common_10))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L113
	}
L99:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v294 != int32(11) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	if l1 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v298 = v17 + int32(16)
	v299 = int32(12)
	v300 = base.I32_rem_s(v76, v299)
	v305 = base.I32_div_s(v76, v299)
	v310 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32)+8)))
	v315 = (base.I64_extend_i32_s(v300*int32(120)) + base.I64_extend_i32_s(v305)*int64(1461) + v310<<(uint(int64(2))%64)) * int64(21600)
	v318 = int64(1000000)
	v319 = int64(0)
	v324 = int64(32)
	v327 = int64(base.Ui64(v315) >> (uint(v324) % 64))
	v330 = int64(4294967295)
	v333 = v315 & v330
	v334 = v318 * v333
	v338 = int64(base.Ui64(v334)>>(uint(v324)%64)) + v318*v327
	v345 = v333*v319 + v338&v330
	*(*int64)(unsafe.Add(mBase, uint32(v298)+8)) = v315*v319 + v315>>(uint(int64(63))%64)*v318 + v319*v327 + int64(base.Ui64(v338)>>(uint(v324)%64)) + int64(base.Ui64(v345)>>(uint(v324)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v334&v330 | v345<<(uint(v324)%64)
	goto L104
L102:
	;
	goto L103
L103:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v383 = int32(12)
	v384 = base.I32_div_s(v76, v383)
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v401 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v379), float64(86400)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v76-v384*v383), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v384), float64(3.15576e+07)), base.F64_div(base.F64_convert_i64_s(v394), float64(1e+06))))))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L112
	}
L104:
	;
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	if v357 != v358>>(uint(int64(63))%64) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v372 = F_int64_div_fast_to_numeric(m, v356, int32(6))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L109
	}
L106:
	;
	v364 = v356 + v358
	if base.B2i32(v356 < int64(0))^base.B2i32(v364 < v358) != 0 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v368 = F_int64_div_fast_to_numeric(m, v364, int32(6))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v446 = v368
	goto L23
L109:
	;
	v374 = F_int64_to_numeric(m, v315)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v377 = F_numeric_add_opt_error(m, v372, v374, int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v446 = v377
	goto L23
L112:
	;
	v446 = v401
	goto L23
L113:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v411 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v61
	F_errmsg(m, int32(_a_F_interval_part_common_11), v17+int32(32))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_interval_part_common_8), int32(_a_F_interval_part_common_12), int32(_a_F_interval_part_common_10))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v431 = F_int64_to_numeric(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v434 = F_Float8GetDatum(m, base.F64_convert_i64_s(v430))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	v446 = v431
	goto L23
L122:
	;
	v446 = v434
	goto L23
}
func F_interval_sum(m *base.Module, l0 int32) int32 {
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
	var v43 int64
	_ = v43
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 != 0 {
		v19 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v6 == int32(0) {
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			return int32(0)
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
						return int32(0)
					} else {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
						if int64(0) < v35 {
							*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(9223372034707292159)
							*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(9223372036854775807)
							return v31
						} else {
							v43 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
							if int64(0) < v43 {
								*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(-9223372034707292160)
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(-9223372036854775807 - 1)
								return v31
							} else {
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v51
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = v53
								return v31
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_interval_sum_0), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_interval_sum_1), int32(_a_F_interval_sum_2), int32(_a_F_interval_sum_3))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
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
				return int32(0)
			}
		}
	}
}
func F_interval_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v50 int32
	_ = v50
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != int32(457) {
		v50 = v2
		return v50
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 != int32(7) {
			v50 = v2
			return v50
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
			if v17 != 0 {
				v50 = v2
				return v50
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
				if v19 < int32(0) {
					v44 = F_relabel_to_typmod(m, v18, v19)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v50 = v44
						return v50
					}
				} else {
					v22 = F_exprTypmod(m, v18)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_intervaltypmodleastfield(m, v22)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = F_intervaltypmodleastfield(m, v19)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v26) < base.Ui32(v28) {
									v50 = v2
									return v50
								} else {
									if v26 != 0 {
										v44 = F_relabel_to_typmod(m, v18, v19)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											v50 = v44
											return v50
										}
									} else {
										v32 = v19 & int32(_a_F_interval_support_0)
										if base.Ui32(int32(5)) < base.Ui32(v32) {
											v44 = F_relabel_to_typmod(m, v18, v19)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												v50 = v44
												return v50
											}
										} else {
											if v22 < int32(0) {
												v38 = int32(-1)
											} else {
												v38 = v22
											}
											if base.Ui32(v32) < base.Ui32(v38&int32(_a_F_interval_support_0)) {
												v50 = v2
												return v50
											} else {
												v44 = F_relabel_to_typmod(m, v18, v19)
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return int32(0)
												} else {
													v50 = v44
													return v50
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
