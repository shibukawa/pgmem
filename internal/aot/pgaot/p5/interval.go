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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v3) & v9
	v13 = base.I64_reinterpret_f64(v4) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v13) {
		v22 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v30 = int32(0) - v22&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
	} else {
		v18 = int32(1)
		if base.F64_gt(v3, v4) != 0 {
			v30 = v18
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
				v30 = v18
			} else {
				v22 = v18
				v30 = int32(0) - v22&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
			}
		}
	}
	return v30
}
func F_interval_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v9 int32
	_ = v9
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v173 int32
	_ = v173
	var v191 int32
	_ = v191
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v311 int64
	_ = v311
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v325 int64
	_ = v325
	var v332 int64
	_ = v332
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v349 int64
	_ = v349
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v364 int64
	_ = v364
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v406 int32
	_ = v406
	var v407 int64
	_ = v407
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v416 int64
	_ = v416
	var v419 int64
	_ = v419
	var v422 int64
	_ = v422
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v430 int64
	_ = v430
	var v437 int64
	_ = v437
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v453 int64
	_ = v453
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v468 int64
	_ = v468
	var v470 int64
	_ = v470
	var v471 int64
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v485 int64
	_ = v485
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v716 int32
	_ = v716
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v806 int32
	_ = v806
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v922 int32
	_ = v922
	var v952 int32
	_ = v952
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1026 int64
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 float64
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1061 int64
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1067 int64
	_ = v1067
	var v1070 int64
	_ = v1070
	var v1073 int64
	_ = v1073
	var v1076 int64
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1081 int64
	_ = v1081
	var v1088 int64
	_ = v1088
	var v1099 int64
	_ = v1099
	var v1100 int64
	_ = v1100
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int64
	_ = v1113
	var v1116 int64
	_ = v1116
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int64
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1191 int32
	_ = v1191
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1219 float64
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1229 float64
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 float64
	_ = v1235
	var v1236 int64
	_ = v1236
	var v1238 float64
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1245 int64
	_ = v1245
	var v1247 int64
	_ = v1247
	var v1253 int64
	_ = v1253
	var v1255 float64
	_ = v1255
	var v1259 int64
	_ = v1259
	var v1260 int64
	_ = v1260
	var v1272 int64
	_ = v1272
	var v1274 int64
	_ = v1274
	var v1276 float64
	_ = v1276
	var v1287 int64
	_ = v1287
	var v1288 int64
	_ = v1288
	var v1304 int32
	_ = v1304
	var v1307 int64
	_ = v1307
	var v1308 int64
	_ = v1308
	var v1313 int64
	_ = v1313
	var v1316 int64
	_ = v1316
	var v1319 int64
	_ = v1319
	var v1322 int64
	_ = v1322
	var v1323 int64
	_ = v1323
	var v1327 int64
	_ = v1327
	var v1334 int64
	_ = v1334
	var v1345 int64
	_ = v1345
	var v1346 int64
	_ = v1346
	var v1351 int64
	_ = v1351
	var v1352 int64
	_ = v1352
	var v1362 float64
	_ = v1362
	var v1366 int64
	_ = v1366
	var v1368 int64
	_ = v1368
	var v1370 float64
	_ = v1370
	var v1381 int64
	_ = v1381
	var v1382 int64
	_ = v1382
	var v1398 int32
	_ = v1398
	var v1401 int64
	_ = v1401
	var v1402 int64
	_ = v1402
	var v1407 int64
	_ = v1407
	var v1410 int64
	_ = v1410
	var v1413 int64
	_ = v1413
	var v1416 int64
	_ = v1416
	var v1417 int64
	_ = v1417
	var v1421 int64
	_ = v1421
	var v1428 int64
	_ = v1428
	var v1439 int64
	_ = v1439
	var v1440 int64
	_ = v1440
	var v1445 int64
	_ = v1445
	var v1446 int64
	_ = v1446
	var v1456 float64
	_ = v1456
	var v1460 int64
	_ = v1460
	var v1462 int64
	_ = v1462
	var v1464 float64
	_ = v1464
	var v1475 int64
	_ = v1475
	var v1476 int64
	_ = v1476
	var v1494 int32
	_ = v1494
	var v1497 int64
	_ = v1497
	var v1498 int64
	_ = v1498
	var v1503 int64
	_ = v1503
	var v1506 int64
	_ = v1506
	var v1509 int64
	_ = v1509
	var v1512 int64
	_ = v1512
	var v1513 int64
	_ = v1513
	var v1517 int64
	_ = v1517
	var v1524 int64
	_ = v1524
	var v1535 int64
	_ = v1535
	var v1536 int64
	_ = v1536
	var v1541 int64
	_ = v1541
	var v1542 int64
	_ = v1542
	var v1552 float64
	_ = v1552
	var v1556 int64
	_ = v1556
	var v1558 int64
	_ = v1558
	var v1560 float64
	_ = v1560
	var v1571 int64
	_ = v1571
	var v1572 int64
	_ = v1572
	var v1588 int32
	_ = v1588
	var v1591 int64
	_ = v1591
	var v1592 int64
	_ = v1592
	var v1597 int64
	_ = v1597
	var v1600 int64
	_ = v1600
	var v1603 int64
	_ = v1603
	var v1606 int64
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1611 int64
	_ = v1611
	var v1618 int64
	_ = v1618
	var v1629 int64
	_ = v1629
	var v1630 int64
	_ = v1630
	var v1635 int64
	_ = v1635
	var v1636 int64
	_ = v1636
	var v1646 float64
	_ = v1646
	var v1650 int64
	_ = v1650
	var v1652 int64
	_ = v1652
	var v1654 float64
	_ = v1654
	var v1665 int64
	_ = v1665
	var v1666 int64
	_ = v1666
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1696 float64
	_ = v1696
	var v1700 int64
	_ = v1700
	var v1702 int64
	_ = v1702
	var v1704 float64
	_ = v1704
	var v1715 int64
	_ = v1715
	var v1716 int64
	_ = v1716
	var v1717 int64
	_ = v1717
	var v1737 int64
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1757 float64
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1772 float64
	_ = v1772
	var v1776 float64
	_ = v1776
	var v1780 int64
	_ = v1780
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
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1831 float64
	_ = v1831
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1847 float64
	_ = v1847
	var v1851 float64
	_ = v1851
	var v1855 int64
	_ = v1855
	var v1857 int64
	_ = v1857
	var v1859 float64
	_ = v1859
	var v1870 int64
	_ = v1870
	var v1871 int64
	_ = v1871
	var v1872 int64
	_ = v1872
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1906 float64
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1931 int64
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1952 float64
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1977 int64
	_ = v1977
	var v1981 int32
	_ = v1981
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1998 float64
	_ = v1998
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2023 int64
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2044 float64
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2099 int32
	_ = v2099
	var v2109 int32
	_ = v2109
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2162 int64
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2195 int32
	_ = v2195
	var v2240 int32
	_ = v2240
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2276 int32
	_ = v2276
	var v2279 int64
	_ = v2279
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2303 int32
	_ = v2303
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2337 int32
	_ = v2337
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2359 float64
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2364 int32
	_ = v2364
	var v2365 float64
	_ = v2365
	var v2374 int32
	_ = v2374
	var v2382 float64
	_ = v2382
	var v2386 int64
	_ = v2386
	var v2388 int64
	_ = v2388
	var v2391 float64
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2409 int64
	_ = v2409
	var v2417 int32
	_ = v2417
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2431 int32
	_ = v2431
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2449 int32
	_ = v2449
	var v2452 int32
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2465 int32
	_ = v2465
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2477 int32
	_ = v2477
	var v2485 int32
	_ = v2485
	var v2492 int64
	_ = v2492
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2514 float64
	_ = v2514
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2547 int32
	_ = v2547
	var v2552 float64
	_ = v2552
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2560 int32
	_ = v2560
	var v2568 float64
	_ = v2568
	var v2573 float64
	_ = v2573
	var v2577 int64
	_ = v2577
	var v2579 int64
	_ = v2579
	var v2581 float64
	_ = v2581
	var v2592 int64
	_ = v2592
	var v2593 int64
	_ = v2593
	var v2594 int64
	_ = v2594
	var v2610 int64
	_ = v2610
	var v2614 int32
	_ = v2614
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2627 int32
	_ = v2627
	var v2632 float64
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2647 float64
	_ = v2647
	var v2652 float64
	_ = v2652
	var v2656 int64
	_ = v2656
	var v2658 int64
	_ = v2658
	var v2660 float64
	_ = v2660
	var v2671 int64
	_ = v2671
	var v2672 int64
	_ = v2672
	var v2673 int64
	_ = v2673
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2697 int32
	_ = v2697
	var v2702 float64
	_ = v2702
	var v2706 int64
	_ = v2706
	var v2708 int64
	_ = v2708
	var v2710 float64
	_ = v2710
	var v2721 int64
	_ = v2721
	var v2722 int64
	_ = v2722
	var v2723 int64
	_ = v2723
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2742 int32
	_ = v2742
	var v2743 int64
	_ = v2743
	var v2744 int64
	_ = v2744
	var v2746 int64
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2755 int32
	_ = v2755
	var v2757 int64
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2769 float64
	_ = v2769
	var v2773 int64
	_ = v2773
	var v2775 int64
	_ = v2775
	var v2777 float64
	_ = v2777
	var v2788 int64
	_ = v2788
	var v2789 int64
	_ = v2789
	var v2790 int64
	_ = v2790
	var v2802 int32
	_ = v2802
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2824 float64
	_ = v2824
	var v2828 int32
	_ = v2828
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2839 int32
	_ = v2839
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2853 int64
	_ = v2853
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2869 float64
	_ = v2869
	var v2873 float64
	_ = v2873
	var v2877 int32
	_ = v2877
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2883 int32
	_ = v2883
	var v2890 float64
	_ = v2890
	var v2894 float64
	_ = v2894
	var v2898 int64
	_ = v2898
	var v2900 int64
	_ = v2900
	var v2902 float64
	_ = v2902
	var v2913 int64
	_ = v2913
	var v2914 int64
	_ = v2914
	var v2915 int64
	_ = v2915
	var v2930 int32
	_ = v2930
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2944 int32
	_ = v2944
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2954 int64
	_ = v2954
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2969 float64
	_ = v2969
	var v2973 float64
	_ = v2973
	var v2977 int64
	_ = v2977
	var v2979 int64
	_ = v2979
	var v2981 float64
	_ = v2981
	var v2992 int64
	_ = v2992
	var v2993 int64
	_ = v2993
	var v2994 int64
	_ = v2994
	var v3007 int32
	_ = v3007
	var v3008 int32
	_ = v3008
	var v3016 int32
	_ = v3016
	var v3019 int64
	_ = v3019
	var v3020 int64
	_ = v3020
	var v3025 int64
	_ = v3025
	var v3028 int64
	_ = v3028
	var v3031 int64
	_ = v3031
	var v3034 int64
	_ = v3034
	var v3035 int64
	_ = v3035
	var v3039 int64
	_ = v3039
	var v3046 int64
	_ = v3046
	var v3057 int64
	_ = v3057
	var v3058 int64
	_ = v3058
	var v3063 int64
	_ = v3063
	var v3064 int64
	_ = v3064
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3076 float64
	_ = v3076
	var v3080 int64
	_ = v3080
	var v3082 int64
	_ = v3082
	var v3084 float64
	_ = v3084
	var v3095 int64
	_ = v3095
	var v3096 int64
	_ = v3096
	var v3106 int32
	_ = v3106
	var v3109 int64
	_ = v3109
	var v3110 int64
	_ = v3110
	var v3115 int64
	_ = v3115
	var v3118 int64
	_ = v3118
	var v3121 int64
	_ = v3121
	var v3124 int64
	_ = v3124
	var v3125 int64
	_ = v3125
	var v3129 int64
	_ = v3129
	var v3136 int64
	_ = v3136
	var v3147 int64
	_ = v3147
	var v3148 int64
	_ = v3148
	var v3153 int64
	_ = v3153
	var v3154 int64
	_ = v3154
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3166 float64
	_ = v3166
	var v3170 int64
	_ = v3170
	var v3172 int64
	_ = v3172
	var v3174 float64
	_ = v3174
	var v3185 int64
	_ = v3185
	var v3186 int64
	_ = v3186
	var v3196 int32
	_ = v3196
	var v3199 int64
	_ = v3199
	var v3200 int64
	_ = v3200
	var v3205 int64
	_ = v3205
	var v3208 int64
	_ = v3208
	var v3211 int64
	_ = v3211
	var v3214 int64
	_ = v3214
	var v3215 int64
	_ = v3215
	var v3219 int64
	_ = v3219
	var v3226 int64
	_ = v3226
	var v3237 int64
	_ = v3237
	var v3238 int64
	_ = v3238
	var v3243 int64
	_ = v3243
	var v3244 int64
	_ = v3244
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3256 float64
	_ = v3256
	var v3260 int64
	_ = v3260
	var v3262 int64
	_ = v3262
	var v3264 float64
	_ = v3264
	var v3275 int64
	_ = v3275
	var v3276 int64
	_ = v3276
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3295 int32
	_ = v3295
	var v3334 int32
	_ = v3334
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3342 int32
	_ = v3342
	var v3344 int32
	_ = v3344
	var v3345 int64
	_ = v3345
	var v3353 int32
	_ = v3353
	var v3357 int32
	_ = v3357
	var v3361 int32
	_ = v3361
	var v3367 int32
	_ = v3367
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3385 int32
	_ = v3385
	var v3388 int32
	_ = v3388
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3401 int32
	_ = v3401
	var v3407 int32
	_ = v3407
	var v3409 int32
	_ = v3409
	var v3413 int32
	_ = v3413
	var v3421 int32
	_ = v3421
	var v3428 int32
	_ = v3428
	var v3430 int64
	_ = v3430
	var v3433 int64
	_ = v3433
	var v3434 int64
	_ = v3434
	var v3439 int64
	_ = v3439
	var v3442 int64
	_ = v3442
	var v3445 int64
	_ = v3445
	var v3448 int64
	_ = v3448
	var v3449 int64
	_ = v3449
	var v3453 int64
	_ = v3453
	var v3460 int64
	_ = v3460
	var v3471 int32
	_ = v3471
	var v3472 int64
	_ = v3472
	var v3473 int64
	_ = v3473
	var v3477 int64
	_ = v3477
	var v3478 int64
	_ = v3478
	var v3484 int64
	_ = v3484
	var v3485 int64
	_ = v3485
	var v3487 int64
	_ = v3487
	var v3489 int64
	_ = v3489
	var v3490 int64
	_ = v3490
	var v3500 int64
	_ = v3500
	var v3501 int64
	_ = v3501
	var v3512 int64
	_ = v3512
	var v3514 int64
	_ = v3514
	var v3516 float64
	_ = v3516
	var v3527 int64
	_ = v3527
	var v3528 int64
	_ = v3528
	var v3541 int32
	_ = v3541
	var v3544 int64
	_ = v3544
	var v3545 int64
	_ = v3545
	var v3550 int64
	_ = v3550
	var v3553 int64
	_ = v3553
	var v3556 int64
	_ = v3556
	var v3559 int64
	_ = v3559
	var v3560 int64
	_ = v3560
	var v3564 int64
	_ = v3564
	var v3571 int64
	_ = v3571
	var v3582 int64
	_ = v3582
	var v3583 int64
	_ = v3583
	var v3588 int64
	_ = v3588
	var v3589 int64
	_ = v3589
	var v3599 float64
	_ = v3599
	var v3603 int64
	_ = v3603
	var v3605 int64
	_ = v3605
	var v3607 float64
	_ = v3607
	var v3618 int64
	_ = v3618
	var v3619 int64
	_ = v3619
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3642 int64
	_ = v3642
	var v3643 float64
	_ = v3643
	var v3644 int64
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3649 int32
	_ = v3649
	var v3651 int32
	_ = v3651
	var v3653 int64
	_ = v3653
	var v3658 int64
	_ = v3658
	var v3659 int64
	_ = v3659
	var v3663 int64
	_ = v3663
	var v3664 int64
	_ = v3664
	var v3674 float64
	_ = v3674
	var v3678 int64
	_ = v3678
	var v3680 int64
	_ = v3680
	var v3682 float64
	_ = v3682
	var v3693 int64
	_ = v3693
	var v3694 int64
	_ = v3694
	var v3705 int32
	_ = v3705
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3722 int32
	_ = v3722
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3733 int64
	_ = v3733
	var v3734 float64
	_ = v3734
	var v3735 int64
	_ = v3735
	var v3736 int32
	_ = v3736
	var v3740 int32
	_ = v3740
	var v3742 int32
	_ = v3742
	var v3744 int64
	_ = v3744
	var v3749 int64
	_ = v3749
	var v3750 int64
	_ = v3750
	var v3754 int64
	_ = v3754
	var v3755 int64
	_ = v3755
	var v3765 float64
	_ = v3765
	var v3769 int64
	_ = v3769
	var v3771 int64
	_ = v3771
	var v3773 float64
	_ = v3773
	var v3784 int64
	_ = v3784
	var v3785 int64
	_ = v3785
	var v3796 int32
	_ = v3796
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3808 int32
	_ = v3808
	var v3817 int32
	_ = v3817
	var v3834 int32
	_ = v3834
	var v3875 int32
	_ = v3875
	var v3902 int32
	_ = v3902
	var v3907 int32
	_ = v3907
	var v3908 int32
	_ = v3908
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3914 int32
	_ = v3914
	var v3917 int64
	_ = v3917
	var v3918 int64
	_ = v3918
	var v3921 int64
	_ = v3921
	var v3927 int32
	_ = v3927
	var v3929 int64
	_ = v3929
	var v3931 int32
	_ = v3931
	var v3932 int32
	_ = v3932
	var v3933 int32
	_ = v3933
	var v3938 int32
	_ = v3938
	var v3942 int32
	_ = v3942
	var v3947 int32
	_ = v3947
	var v3955 int32
	_ = v3955
	var v3957 int32
	_ = v3957
	var v3961 int32
	_ = v3961
	var v3966 int32
	_ = v3966
	var v3973 int32
	_ = v3973
	var v3975 int32
	_ = v3975
	v2 = int64(0)
	v9 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(528)
	m.G0 = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+512)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v40)+520)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v40)+504)) = v2
	v60 = F_ParseDateTime(m, v44, v40+int32(16), int32(256), v40+int32(384), v40+int32(272), v40+int32(496))
	mBase = m.M
	if v60 == v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v64 = v40 + int32(384)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40)+496))
	if v42 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v2240 = v60
	goto L3
L3:
	;
	if v2240 == int32(-1) {
		goto L482
	} else {
		goto L483
	}
L4:
	;
	v73 = int32(32767)
	goto L6
L5:
	;
	v73 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
	goto L6
L6:
	;
	v74 = m.G0
	v76 = v74 - int32(160)
	m.G0 = v76
	v79 = v40 + int32(500)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(17)
	v83 = v40 + int32(504)
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v84
	v86 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v83)+8)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v86
	if v67 <= v84 {
		v173 = v9
		goto L10
	} else {
		goto L11
	}
L7:
	;
	m.G0 = v76 + int32(160)
	v2240 = v2195
	goto L3
L8:
	;
	v257 = v9
	v258 = int32(8)
	v261 = v214
	v273 = v9
	v275 = v9
	goto L20
L9:
	;
	v214 = v67 - int32(1)
	v219 = base.B2i32(v97 == int32(45))
	goto L8
L10:
	;
	v191 = v67 - int32(1)
	if v191 < int32(0) {
		v2195 = int32(-1)
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[1114]))
	if v93 != int32(2) {
		v173 = v9
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97 != int32(45) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v67) < base.Ui32(int32(2)) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v111 = int32(1)
	goto L15
L15:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v64+v111<<(uint(int32(2))%32))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	switch v145 - int32(43) {
	case 0, 2:
		v173 = int32(0)
		goto L10
	default:
		goto L17
	}
L16:
	;
	v173 = v148
	goto L10
L17:
	;
	v148 = int32(1)
	v150 = v111 + v148
	if v150 != v67 {
		v111 = v150
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v214 = v191
	v219 = v173
	goto L8
L20:
	;
	v285 = int32(-1)
	v287 = v261 << (uint(int32(2)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(272)+v287)))
	switch v289 {
	case 0, 2:
		goto L27
	case 1, 6:
		goto L26
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		v2195 = v285
		goto L7
	}
L21:
	;
	v2154 = int32(0)
	v2157 = v2122 | base.B2i32(v2138 == v2154)
	v2158 = v2154 - v2157
	if v2157 != 0 {
		v2195 = v2158
		goto L7
	} else {
		goto L476
	}
L22:
	;
	if int32(0) < v261 {
		v257 = v2122
		v258 = v2123
		v261 = v261 - int32(1)
		v273 = v2138
		v275 = v2140
		goto L20
	} else {
		goto L475
	}
L23:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v76)+112))
	if v2109&v273 != 0 {
		goto L472
	} else {
		goto L473
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v1020 = v287 + v64
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1026 = F_strtox_2(m, v1021, v76+int32(116), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L174
L25:
	;
	v1016 = int32(18)
	goto L24
L26:
	;
	if v257&int32(1) != 0 {
		v2195 = v285
		goto L7
	} else {
		goto L72
	}
L27:
	;
	if v258 != int32(8) {
		v1016 = v258
		goto L24
	} else {
		goto L58
	}
L28:
	;
	v383 = v287 + v64
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v386 = v384 + int32(1)
	v387 = int32(58)
	v388 = F___strchrnul(m, v386, v387)
	mBase = m.M
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v390 == v387 {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v64)))
	v296 = F_DecodeTimeCommon(m, v291, v73, v76+int32(112), v76+int32(120))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	if v296 != 0 {
		v2195 = v296
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v300 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v300
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v76)+136))
	v305 = int64(3600000000)
	v306 = int64(0)
	v311 = int64(32)
	v314 = int64(base.Ui64(v302) >> (uint(v311) % 64))
	v317 = int64(4294967295)
	v320 = v302 & v317
	v321 = v305 * v320
	v325 = int64(base.Ui64(v321)>>(uint(v311)%64)) + v305*v314
	v332 = v320*v306 + v325&v317
	*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v302*v306 + v302>>(uint(int64(63))%64)*v305 + v306*v314 + int64(base.Ui64(v325)>>(uint(v311)%64)) + int64(base.Ui64(v332)>>(uint(v311)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = v321&v317 | v332<<(uint(v311)%64)
	goto L33
L33:
	;
	v343 = int32(-2)
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
	if v344 != v345>>(uint(int64(63))%64) {
		v2195 = v343
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v349 = v300 + v345
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v349
	if base.B2i32(v345 < int64(0))^base.B2i32(v349 < v300) != 0 {
		v2195 = v343
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v355 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+128)))
	v357 = v355 * int64(60000000)
	v358 = v349 + v357
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v358
	if base.B2i32(v357 < int64(0))^base.B2i32(v358 < v349) != 0 {
		v2195 = v343
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v364 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+124)))
	v366 = v364 * int64(1000000)
	v367 = v358 + v366
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v367
	if base.B2i32(v366 < int64(0))^base.B2i32(v367 < v358) != 0 {
		v2195 = v343
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v373 = int32(21)
	v374 = int32(0)
	if base.B2i32(v219 == v374)|base.B2i32(v367 <= int64(0)) != 0 {
		v2081 = v374
		v2082 = v373
		v2099 = v275
		goto L23
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0) - v367
	v2081 = v374
	v2082 = v373
	v2099 = v275
	goto L23
L39:
	;
	if v394 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L40:
	;
	v394 = v388
	goto L42
L41:
	;
	v394 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v401 = F_DecodeTimeCommon(m, v386, v73, v76+int32(112), v76+int32(120))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	if v401 != 0 {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v403 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v403
	v406 = v76 + int32(96)
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v76)+136))
	v410 = int64(3600000000)
	v411 = int64(0)
	v416 = int64(32)
	v419 = int64(base.Ui64(v407) >> (uint(v416) % 64))
	v422 = int64(4294967295)
	v425 = v407 & v422
	v426 = v410 * v425
	v430 = int64(base.Ui64(v426)>>(uint(v416)%64)) + v410*v419
	v437 = v425*v411 + v430&v422
	*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v407*v411 + v407>>(uint(int64(63))%64)*v410 + v411*v419 + int64(base.Ui64(v430)>>(uint(v416)%64)) + int64(base.Ui64(v437)>>(uint(v416)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v406))) = v426&v422 | v437<<(uint(v416)%64)
	goto L46
L46:
	;
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v76)+104))
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v76)+96))
	if v448 != v449>>(uint(int64(63))%64) {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v453 = v403 + v449
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v453
	if base.B2i32(v449 < int64(0))^base.B2i32(v453 < v403) != 0 {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	v459 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+128)))
	v461 = v459 * int64(60000000)
	v462 = v453 + v461
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v462
	if base.B2i32(v461 < int64(0))^base.B2i32(v462 < v453) != 0 {
		goto L27
	} else {
		goto L49
	}
L49:
	;
	v468 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+124)))
	v470 = v468 * int64(1000000)
	v471 = v462 + v470
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v471
	if base.B2i32(v470 < int64(0))^base.B2i32(v471 < v462) != 0 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if v478 == int32(45) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v471 == int64(-9223372036854775807-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v487 = v471
	goto L53
L53:
	;
	v488 = int32(21)
	v489 = int32(0)
	if v219&base.B2i32(int64(0) < v487) == v489 {
		v2081 = v489
		v2082 = v488
		v2099 = v275
		goto L23
	} else {
		goto L57
	}
L54:
	;
	v2195 = int32(-2)
	goto L7
L55:
	;
	goto L56
L56:
	;
	v485 = int64(0) - v471
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v485
	v487 = v485
	goto L53
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0) - v487
	v2081 = v489
	v2082 = v488
	v2099 = v275
	goto L23
L58:
	;
	if base.B2i32(int32(3071) < v73) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v1016 = int32(19)
	goto L24
L60:
	;
	if base.B2i32(v73 == int32(2048)) == int32(0) {
		goto L25
	} else {
		goto L71
	}
L61:
	;
	v1016 = int32(20)
	goto L24
L62:
	;
	v1016 = int32(21)
	goto L24
L63:
	;
	v1016 = int32(23)
	goto L24
L64:
	;
	switch v73 - int32(2) {
	case 0, 4:
		goto L63
	case 1, 3, 5:
		goto L25
	case 2:
		v1016 = int32(25)
		goto L24
	case 6:
		goto L62
	default:
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if base.B2i32(int32(6143) < v73) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	switch v73 - int32(1024) {
	case 0, 8:
		goto L61
	case 1, 2, 3, 4, 5, 6, 7:
		goto L25
	default:
		goto L60
	}
L68:
	;
	switch v73 - int32(3072) {
	case 0, 8:
		goto L59
	default:
		goto L25
	}
L69:
	;
	goto L70
L70:
	;
	goto L25
L71:
	;
	goto L59
L72:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v287+v64)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1115])))
	if v523 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v990 = int32(8)
	v993 = v967 & int32(255)
	if v993 == v990 {
		v2122 = int32(0)
		v2123 = v990
		v2138 = v273
		v2140 = v275
		goto L22
	} else {
		goto L168
	}
L74:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1042])))
	if v758 != 0 {
		goto L123
	} else {
		goto L124
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1115]))) = v686
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+11)))
	if v716 != int32(31) {
		v961 = v686
		v967 = v716
		goto L73
	} else {
		goto L121
	}
L76:
	;
	goto L81
L77:
	;
	goto L78
L78:
	;
	v571 = int32(*(*int8)(unsafe.Add(mBase, uint32(v520))))
	v583 = int32(1664608)
	v584 = int32(1663648)
	goto L94
L79:
	;
	if v560-v561 == int32(0) {
		v686 = v523
		goto L75
	} else {
		goto L93
	}
L81:
	;
	goto L82
L82:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v530 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v531 = v520
	v532 = v523
	v533 = int32(10)
	v534 = v530
	goto L87
L84:
	;
	v556 = v523
	v560 = int32(0)
	goto L85
L85:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	goto L79
L86:
	;
	v556 = v551
	v560 = v553
	goto L85
L87:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	if v534 != v536 {
		v551 = v532
		v553 = v534
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v551 = v545
	v553 = int32(0)
	goto L86
L89:
	;
	if v536 == int32(0) {
		v551 = v532
		v553 = v534
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v541 = v533 - int32(1)
	if v541 == int32(0) {
		v551 = v532
		v553 = v534
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v544 = int32(1)
	v545 = v532 + v544
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+1)))
	if v546 != 0 {
		v531 = v531 + v544
		v532 = v545
		v533 = v541
		v534 = v546
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	goto L78
L94:
	;
	v616 = v584 + (v583-v584)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v616))))
	v618 = v571 - v617
	if v618 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L74
L96:
	;
	goto L101
L97:
	;
	v668 = v618
	goto L98
L98:
	;
	v672 = base.B2i32(v668 < int32(0))
	if v668 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L99:
	;
	if v659 == int32(0) {
		v686 = v616
		goto L75
	} else {
		goto L113
	}
L101:
	;
	goto L102
L102:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v627 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v628 = v520
	v629 = v616
	v630 = int32(10)
	v631 = v627
	goto L107
L104:
	;
	v653 = v616
	v657 = int32(0)
	goto L105
L105:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653))))
	v659 = v657 - v658
	goto L99
L106:
	;
	v653 = v648
	v657 = v650
	goto L105
L107:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	if v631 != v633 {
		v648 = v629
		v650 = v631
		goto L106
	} else {
		goto L109
	}
L108:
	;
	v648 = v642
	v650 = int32(0)
	goto L106
L109:
	;
	if v633 == int32(0) {
		v648 = v629
		v650 = v631
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v638 = v630 - int32(1)
	if v638 == int32(0) {
		v648 = v629
		v650 = v631
		goto L106
	} else {
		goto L111
	}
L111:
	;
	v641 = int32(1)
	v642 = v629 + v641
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628)+1)))
	if v643 != 0 {
		v628 = v628 + v641
		v629 = v642
		v630 = v638
		v631 = v643
		goto L107
	} else {
		goto L112
	}
L112:
	;
	goto L108
L113:
	;
	v668 = v659
	goto L98
L114:
	;
	v673 = v616 - int32(16)
	goto L116
L115:
	;
	v673 = v583
	goto L116
L116:
	;
	if v668 < int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v676 = v584
	goto L119
L118:
	;
	v676 = v616 + int32(16)
	goto L119
L119:
	;
	if base.Ui32(v676) <= base.Ui32(v673) {
		v583 = v673
		v584 = v676
		goto L94
	} else {
		goto L120
	}
L120:
	;
	goto L95
L121:
	;
	goto L74
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1042]))) = v922
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922)+11)))
	v961 = v922
	v967 = v952
	goto L73
L123:
	;
	goto L128
L124:
	;
	goto L125
L125:
	;
	v806 = int32(*(*int8)(unsafe.Add(mBase, uint32(v520))))
	v818 = int32(1663632)
	v819 = int32(1662496)
	goto L141
L126:
	;
	if v795-v796 == int32(0) {
		v922 = v758
		goto L122
	} else {
		goto L140
	}
L128:
	;
	goto L129
L129:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v765 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v766 = v520
	v767 = v758
	v768 = int32(10)
	v769 = v765
	goto L134
L131:
	;
	v791 = v758
	v795 = int32(0)
	goto L132
L132:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791))))
	goto L126
L133:
	;
	v791 = v786
	v795 = v788
	goto L132
L134:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	if v769 != v771 {
		v786 = v767
		v788 = v769
		goto L133
	} else {
		goto L136
	}
L135:
	;
	v786 = v780
	v788 = int32(0)
	goto L133
L136:
	;
	if v771 == int32(0) {
		v786 = v767
		v788 = v769
		goto L133
	} else {
		goto L137
	}
L137:
	;
	v776 = v768 - int32(1)
	if v776 == int32(0) {
		v786 = v767
		v788 = v769
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v779 = int32(1)
	v780 = v767 + v779
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+1)))
	if v781 != 0 {
		v766 = v766 + v779
		v767 = v780
		v768 = v776
		v769 = v781
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
L140:
	;
	goto L125
L141:
	;
	v851 = v819 + (v818-v819)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v852 = int32(*(*int8)(unsafe.Add(mBase, uint32(v851))))
	v853 = v806 - v852
	if v853 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v2195 = int32(-1)
	goto L7
L143:
	;
	goto L148
L144:
	;
	v903 = v853
	goto L145
L145:
	;
	v907 = base.B2i32(v903 < int32(0))
	if v903 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L146:
	;
	if v894 == int32(0) {
		v922 = v851
		goto L122
	} else {
		goto L160
	}
L148:
	;
	goto L149
L149:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v862 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v863 = v520
	v864 = v851
	v865 = int32(10)
	v866 = v862
	goto L154
L151:
	;
	v888 = v851
	v892 = int32(0)
	goto L152
L152:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	v894 = v892 - v893
	goto L146
L153:
	;
	v888 = v883
	v892 = v885
	goto L152
L154:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864))))
	if v866 != v868 {
		v883 = v864
		v885 = v866
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v883 = v877
	v885 = int32(0)
	goto L153
L156:
	;
	if v868 == int32(0) {
		v883 = v864
		v885 = v866
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v873 = v865 - int32(1)
	if v873 == int32(0) {
		v883 = v864
		v885 = v866
		goto L153
	} else {
		goto L158
	}
L158:
	;
	v876 = int32(1)
	v877 = v864 + v876
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863)+1)))
	if v878 != 0 {
		v863 = v863 + v876
		v864 = v877
		v865 = v873
		v866 = v878
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L155
L160:
	;
	v903 = v894
	goto L145
L161:
	;
	v908 = v851 - int32(16)
	goto L163
L162:
	;
	v908 = v818
	goto L163
L163:
	;
	if v903 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v911 = v819
	goto L166
L165:
	;
	v911 = v851 + int32(16)
	goto L166
L166:
	;
	if base.Ui32(v911) <= base.Ui32(v908) {
		v818 = v908
		v819 = v911
		goto L141
	} else {
		goto L167
	}
L167:
	;
	goto L142
L168:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v961)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(0)
	v1000 = int32(-1)
	switch v993 {
	case 0:
		goto L169
	default:
		v2195 = v1000
		goto L7
	case 17:
		v2081 = int32(1)
		v2082 = v996
		v2099 = v275
		goto L23
	case 19:
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(31758)
	if v261 != v214 {
		v2195 = v1000
		goto L7
	} else {
		goto L172
	}
L170:
	;
	if v261 != v214 {
		v2195 = v1000
		goto L7
	} else {
		goto L171
	}
L171:
	;
	v2081 = int32(0)
	v2082 = v996
	v2099 = int32(1)
	goto L23
L172:
	;
	if base.Ui32(int32(1)) < base.Ui32(v996-int32(9)) {
		v2195 = v1000
		goto L7
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v996
	v1012 = int32(0)
	v2081 = v1012
	v2082 = v1012
	v2099 = v275
	goto L23
L174:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v1028 == int32(68) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v2195 = int32(-2)
	goto L7
L176:
	;
	goto L177
L177:
	;
	v1032 = float64(0)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v76)+116))
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	switch v1034 - int32(45) {
	case 0:
		goto L180
	case 1:
		goto L179
	default:
		goto L181
	}
L178:
	;
	if v219 == int32(0) {
		v1253 = v1236
		v1255 = v1238
		goto L232
	} else {
		goto L233
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+120)) = v1033
	v1124 = v1033 + int32(1)
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	if v1125 == int32(0) {
		v1229 = v1032
		goto L201
	} else {
		goto L202
	}
L180:
	;
	v1045 = F_strtol(m, v1033+int32(1), v76+int32(116), int32(10))
	mBase = m.M
	goto L183
L181:
	;
	if v1034 == int32(0) {
		v1236 = v1026
		v1238 = v1032
		v1240 = v1016
		goto L178
	} else {
		goto L182
	}
L182:
	;
	v2195 = int32(-1)
	goto L7
L183:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v1047 == int32(68) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v2195 = int32(-2)
	goto L7
L185:
	;
	goto L186
L186:
	;
	if base.Ui32(int32(11)) < base.Ui32(v1045) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v2195 = int32(-2)
	goto L7
L188:
	;
	goto L189
L189:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v76)+116))
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054))))
	if v1055 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v2195 = int32(-1)
	goto L7
L191:
	;
	goto L192
L192:
	;
	v1058 = v76 + int32(80)
	v1061 = int64(12)
	v1062 = int64(0)
	v1067 = int64(32)
	v1070 = int64(base.Ui64(v1026) >> (uint(v1067) % 64))
	v1073 = int64(4294967295)
	v1076 = v1026 & v1073
	v1077 = v1061 * v1076
	v1081 = int64(base.Ui64(v1077)>>(uint(v1067)%64)) + v1061*v1070
	v1088 = v1076*v1062 + v1081&v1073
	*(*int64)(unsafe.Add(mBase, uint32(v1058)+8)) = v1026*v1062 + v1026>>(uint(int64(63))%64)*v1061 + v1062*v1070 + int64(base.Ui64(v1081)>>(uint(v1067)%64)) + int64(base.Ui64(v1088)>>(uint(v1067)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1058))) = v1077&v1073 | v1088<<(uint(v1067)%64)
	goto L193
L193:
	;
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(v76)+88))
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v76)+80))
	if v1099 != v1100>>(uint(int64(63))%64) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v2195 = int32(-2)
	goto L7
L195:
	;
	goto L196
L196:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
	if v1109 == int32(45) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1112 = int32(0) - v1045
	goto L199
L198:
	;
	v1112 = v1045
	goto L199
L199:
	;
	v1113 = base.I64_extend_i32_s(v1112)
	v1116 = v1113 + v1100
	if base.B2i32(v1113 < int64(0))^base.B2i32(v1116 < v1100) == int32(0) {
		v1236 = v1116
		v1238 = v1032
		v1240 = int32(23)
		goto L178
	} else {
		goto L200
	}
L200:
	;
	v2195 = int32(-2)
	goto L7
L201:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231))))
	if v1232 == int32(45) {
		goto L229
	} else {
		goto L230
	}
L202:
	;
	v1128 = int32(553196)
	v1132 = m.G0
	v1134 = v1132 - int32(32)
	v1135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+24)) = v1135
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+16)) = v1135
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+8)) = v1135
	*(*int64)(unsafe.Add(mBase, uint32(v1134))) = v1135
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v1143 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v2195 = int32(-1)
	goto L7
L204:
	;
	v1212 = F_strlen(m, v1124)
	mBase = m.M
	if v1211 != v1212 {
		goto L203
	} else {
		goto L225
	}
L205:
	;
	v1211 = int32(0)
	goto L204
L206:
	;
	goto L207
L207:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v1147 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1151 = v1124
	goto L211
L209:
	;
	goto L210
L210:
	;
	v1161 = v1128
	v1162 = v1143
	goto L214
L211:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151))))
	if v1157 == v1143 {
		v1151 = v1151 + int32(1)
		goto L211
	} else {
		goto L213
	}
L212:
	;
	v1211 = v1151 - v1124
	goto L204
L213:
	;
	goto L212
L214:
	;
	v1169 = v1134 + int32(base.Ui32(v1162)>>(uint(int32(3))%32))&int32(28)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1169)))
	v1171 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1170 | v1171<<(uint(v1162)%32)
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1161)+1)))
	if v1175 != 0 {
		v1161 = v1161 + v1171
		v1162 = v1175
		goto L214
	} else {
		goto L216
	}
L215:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	if v1178 == int32(0) {
		v1203 = v1124
		goto L217
	} else {
		goto L218
	}
L216:
	;
	goto L215
L217:
	;
	v1211 = v1203 - v1124
	goto L204
L218:
	;
	v1182 = v1124
	v1183 = v1178
	goto L219
L219:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1134+int32(base.Ui32(v1183)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v1191)>>(uint(v1183)%32))&int32(1) == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v1203 = v1199
	goto L217
L221:
	;
	v1203 = v1182
	goto L217
L222:
	;
	goto L223
L223:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182)+1)))
	v1199 = v1182 + int32(1)
	if v1197 != 0 {
		v1182 = v1199
		v1183 = v1197
		goto L219
	} else {
		goto L224
	}
L224:
	;
	goto L220
L225:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v1219 = F_strtod(m, v1033, v76+int32(120))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L30
	} else {
		goto L226
	}
L226:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v76)+120))
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221))))
	if v1222 != 0 {
		goto L203
	} else {
		goto L227
	}
L227:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v1224 == int32(0) {
		v1229 = v1219
		goto L201
	} else {
		goto L228
	}
L228:
	;
	goto L203
L229:
	;
	v1235 = base.F64_neg(v1229)
	goto L231
L230:
	;
	v1235 = v1229
	goto L231
L231:
	;
	v1236 = v1026
	v1238 = v1235
	v1240 = v1016
	goto L178
L232:
	;
	switch v1240 - int32(18) {
	case 0:
		goto L245
	case 1:
		goto L244
	case 2:
		goto L243
	case 3:
		goto L242
	case 4:
		goto L241
	case 5:
		goto L240
	default:
		v2195 = int32(-1)
		goto L7
	case 7:
		goto L239
	case 8:
		goto L238
	case 9:
		goto L237
	case 10:
		goto L236
	case 11:
		goto L246
	case 12:
		goto L247
	}
L233:
	;
	v1245 = v1236 >> (uint(int64(63)) % 64)
	v1247 = v1245 - (v1236 ^ v1245)
	if base.F64_gt(v1238, float64(0)) == int32(0) {
		v1253 = v1247
		v1255 = v1238
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1253 = v1247
	v1255 = base.F64_neg(v1238)
	goto L232
L235:
	;
	v2081 = int32(0)
	v2082 = int32(21)
	v2099 = v275
	goto L23
L236:
	;
	if base.Ui64(v1253-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L458
	} else {
		goto L459
	}
L237:
	;
	if base.Ui64(v1253-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L444
	} else {
		goto L445
	}
L238:
	;
	if base.Ui64(v1253-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L430
	} else {
		goto L431
	}
L239:
	;
	if base.Ui64(v1253-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L419
	} else {
		goto L420
	}
L240:
	;
	if base.Ui64(v1253-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L393
	} else {
		goto L394
	}
L241:
	;
	if base.Ui64(v1253-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L364
	} else {
		goto L365
	}
L242:
	;
	if base.Ui64(v1253-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L345
	} else {
		goto L346
	}
L243:
	;
	v1588 = v76 - int32(-64)
	v1591 = int64(3600000000)
	v1592 = int64(0)
	v1597 = int64(32)
	v1600 = int64(base.Ui64(v1253) >> (uint(v1597) % 64))
	v1603 = int64(4294967295)
	v1606 = v1253 & v1603
	v1607 = v1591 * v1606
	v1611 = int64(base.Ui64(v1607)>>(uint(v1597)%64)) + v1591*v1600
	v1618 = v1606*v1592 + v1611&v1603
	*(*int64)(unsafe.Add(mBase, uint32(v1588)+8)) = v1253*v1592 + v1253>>(uint(int64(63))%64)*v1591 + v1592*v1600 + int64(base.Ui64(v1611)>>(uint(v1597)%64)) + int64(base.Ui64(v1618)>>(uint(v1597)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1588))) = v1607&v1603 | v1618<<(uint(v1597)%64)
	goto L325
L244:
	;
	v1494 = v76 + int32(48)
	v1497 = int64(60000000)
	v1498 = int64(0)
	v1503 = int64(32)
	v1506 = int64(base.Ui64(v1253) >> (uint(v1503) % 64))
	v1509 = int64(4294967295)
	v1512 = v1253 & v1509
	v1513 = v1497 * v1512
	v1517 = int64(base.Ui64(v1513)>>(uint(v1503)%64)) + v1497*v1506
	v1524 = v1512*v1498 + v1517&v1509
	*(*int64)(unsafe.Add(mBase, uint32(v1494)+8)) = v1253*v1498 + v1253>>(uint(int64(63))%64)*v1497 + v1498*v1506 + int64(base.Ui64(v1517)>>(uint(v1503)%64)) + int64(base.Ui64(v1524)>>(uint(v1503)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1494))) = v1513&v1509 | v1524<<(uint(v1503)%64)
	goto L305
L245:
	;
	v1398 = v76 + int32(32)
	v1401 = int64(1000000)
	v1402 = int64(0)
	v1407 = int64(32)
	v1410 = int64(base.Ui64(v1253) >> (uint(v1407) % 64))
	v1413 = int64(4294967295)
	v1416 = v1253 & v1413
	v1417 = v1401 * v1416
	v1421 = int64(base.Ui64(v1417)>>(uint(v1407)%64)) + v1401*v1410
	v1428 = v1416*v1402 + v1421&v1413
	*(*int64)(unsafe.Add(mBase, uint32(v1398)+8)) = v1253*v1402 + v1253>>(uint(int64(63))%64)*v1401 + v1402*v1410 + int64(base.Ui64(v1421)>>(uint(v1407)%64)) + int64(base.Ui64(v1428)>>(uint(v1407)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1398))) = v1417&v1413 | v1428<<(uint(v1407)%64)
	goto L284
L246:
	;
	v1304 = v76 + int32(16)
	v1307 = int64(1000)
	v1308 = int64(0)
	v1313 = int64(32)
	v1316 = int64(base.Ui64(v1253) >> (uint(v1313) % 64))
	v1319 = int64(4294967295)
	v1322 = v1253 & v1319
	v1323 = v1307 * v1322
	v1327 = int64(base.Ui64(v1323)>>(uint(v1313)%64)) + v1307*v1316
	v1334 = v1322*v1308 + v1327&v1319
	*(*int64)(unsafe.Add(mBase, uint32(v1304)+8)) = v1253*v1308 + v1253>>(uint(int64(63))%64)*v1307 + v1308*v1316 + int64(base.Ui64(v1327)>>(uint(v1313)%64)) + int64(base.Ui64(v1334)>>(uint(v1313)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1304))) = v1323&v1319 | v1334<<(uint(v1313)%64)
	goto L264
L247:
	;
	v1259 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1260 = v1259 + v1253
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1260
	if base.B2i32(v1253 < int64(0))^base.B2i32(v1260 < v1259) != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v2195 = int32(-2)
	goto L7
L249:
	;
	goto L250
L250:
	;
	if base.F64_ne(v1255, float64(0)) != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	if base.F64_lt(base.F64_abs(v1255), float64(9.223372036854776e+18)) != 0 {
		goto L256
	} else {
		goto L257
	}
L252:
	;
	goto L253
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(16384)
	v2081 = int32(0)
	v2082 = int32(30)
	v2099 = v275
	goto L23
L254:
	;
	v1288 = v1287 + v1260
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1288
	if base.B2i32(v1287 < int64(0))^base.B2i32(v1288 < v1260) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L263
	}
L255:
	;
	v1276 = base.F64_sub(v1255, base.F64_convert_i64_s(v1274))
	if base.F64_gt(v1276, float64(0.5)) != 0 {
		goto L259
	} else {
		goto L260
	}
L256:
	;
	v1272 = base.I64_trunc_f64_s(v1255)
	v1274 = v1272
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1274 = int64(-9223372036854775807 - 1)
	goto L255
L259:
	;
	v1287 = v1274 + int64(1)
	goto L254
L260:
	;
	goto L261
L261:
	;
	if base.F64_lt(v1276, float64(-0.5)) == int32(0) {
		v1287 = v1274
		goto L254
	} else {
		goto L262
	}
L262:
	;
	v1287 = v1274 - int64(1)
	goto L254
L263:
	;
	goto L253
L264:
	;
	v1345 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
	v1346 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
	if v1345 != v1346>>(uint(int64(63))%64) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v2195 = int32(-2)
	goto L7
L266:
	;
	goto L267
L267:
	;
	v1351 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1352 = v1351 + v1346
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1352
	if base.B2i32(v1346 < int64(0))^base.B2i32(v1352 < v1351) != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v2195 = int32(-2)
	goto L7
L269:
	;
	goto L270
L270:
	;
	if base.F64_ne(v1255, float64(0)) != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1362 = base.F64_mul(v1255, float64(1000))
	if base.F64_lt(base.F64_abs(v1362), float64(9.223372036854776e+18)) != 0 {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(8192)
	v2081 = int32(0)
	v2082 = int32(29)
	v2099 = v275
	goto L23
L274:
	;
	v1370 = base.F64_sub(v1362, base.F64_convert_i64_s(v1368))
	if base.F64_gt(v1370, float64(0.5)) != 0 {
		goto L279
	} else {
		goto L280
	}
L275:
	;
	v1366 = base.I64_trunc_f64_s(v1362)
	v1368 = v1366
	goto L274
L276:
	;
	goto L277
L277:
	;
	v1368 = int64(-9223372036854775807 - 1)
	goto L274
L278:
	;
	v1382 = v1381 + v1352
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1382
	if base.B2i32(v1381 < int64(0))^base.B2i32(v1382 < v1352) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L283
	}
L279:
	;
	v1381 = v1368 + int64(1)
	goto L278
L280:
	;
	goto L281
L281:
	;
	if base.F64_lt(v1370, float64(-0.5)) == int32(0) {
		v1381 = v1368
		goto L278
	} else {
		goto L282
	}
L282:
	;
	v1381 = v1368 - int64(1)
	goto L278
L283:
	;
	goto L273
L284:
	;
	v1439 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
	v1440 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
	if v1439 != v1440>>(uint(int64(63))%64) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v2195 = int32(-2)
	goto L7
L286:
	;
	goto L287
L287:
	;
	v1445 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1446 = v1445 + v1440
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1446
	if base.B2i32(v1440 < int64(0))^base.B2i32(v1446 < v1445) != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v2195 = int32(-2)
	goto L7
L289:
	;
	goto L290
L290:
	;
	if base.F64_ne(v1255, float64(0)) != 0 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v2081 = int32(0)
	v2082 = int32(18)
	v2099 = v275
	goto L23
L292:
	;
	v1456 = base.F64_mul(v1255, float64(1e+06))
	if base.F64_lt(base.F64_abs(v1456), float64(9.223372036854776e+18)) != 0 {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(4096)
	goto L291
L295:
	;
	v1464 = base.F64_sub(v1456, base.F64_convert_i64_s(v1462))
	if base.F64_gt(v1464, float64(0.5)) != 0 {
		goto L300
	} else {
		goto L301
	}
L296:
	;
	v1460 = base.I64_trunc_f64_s(v1456)
	v1462 = v1460
	goto L295
L297:
	;
	goto L298
L298:
	;
	v1462 = int64(-9223372036854775807 - 1)
	goto L295
L299:
	;
	v1476 = v1475 + v1446
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1476
	if base.B2i32(v1475 < int64(0))^base.B2i32(v1476 < v1446) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L304
	}
L300:
	;
	v1475 = v1462 + int64(1)
	goto L299
L301:
	;
	goto L302
L302:
	;
	if base.F64_lt(v1464, float64(-0.5)) == int32(0) {
		v1475 = v1462
		goto L299
	} else {
		goto L303
	}
L303:
	;
	v1475 = v1462 - int64(1)
	goto L299
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(28672)
	goto L291
L305:
	;
	v1535 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
	v1536 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
	if v1535 != v1536>>(uint(int64(63))%64) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v2195 = int32(-2)
	goto L7
L307:
	;
	goto L308
L308:
	;
	v1541 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1542 = v1541 + v1536
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1542
	if base.B2i32(v1536 < int64(0))^base.B2i32(v1542 < v1541) != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v2195 = int32(-2)
	goto L7
L310:
	;
	goto L311
L311:
	;
	if base.F64_ne(v1255, float64(0)) != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1552 = base.F64_mul(v1255, float64(6e+07))
	if base.F64_lt(base.F64_abs(v1552), float64(9.223372036854776e+18)) != 0 {
		goto L316
	} else {
		goto L317
	}
L313:
	;
	goto L314
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(2048)
	v2081 = int32(0)
	v2082 = int32(19)
	v2099 = v275
	goto L23
L315:
	;
	v1560 = base.F64_sub(v1552, base.F64_convert_i64_s(v1558))
	if base.F64_gt(v1560, float64(0.5)) != 0 {
		goto L320
	} else {
		goto L321
	}
L316:
	;
	v1556 = base.I64_trunc_f64_s(v1552)
	v1558 = v1556
	goto L315
L317:
	;
	goto L318
L318:
	;
	v1558 = int64(-9223372036854775807 - 1)
	goto L315
L319:
	;
	v1572 = v1571 + v1542
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1572
	if base.B2i32(v1571 < int64(0))^base.B2i32(v1572 < v1542) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L324
	}
L320:
	;
	v1571 = v1558 + int64(1)
	goto L319
L321:
	;
	goto L322
L322:
	;
	if base.F64_lt(v1560, float64(-0.5)) == int32(0) {
		v1571 = v1558
		goto L319
	} else {
		goto L323
	}
L323:
	;
	v1571 = v1558 - int64(1)
	goto L319
L324:
	;
	goto L314
L325:
	;
	v1629 = *(*int64)(unsafe.Add(mBase, uint32(v76)+72))
	v1630 = *(*int64)(unsafe.Add(mBase, uint32(v76)+64))
	if v1629 != v1630>>(uint(int64(63))%64) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2195 = int32(-2)
	goto L7
L327:
	;
	goto L328
L328:
	;
	v1635 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1636 = v1635 + v1630
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1636
	if base.B2i32(v1630 < int64(0))^base.B2i32(v1636 < v1635) != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v2195 = int32(-2)
	goto L7
L330:
	;
	goto L331
L331:
	;
	if base.F64_ne(v1255, float64(0)) != 0 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1646 = base.F64_mul(v1255, float64(3.6e+09))
	if base.F64_lt(base.F64_abs(v1646), float64(9.223372036854776e+18)) != 0 {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	goto L334
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(1024)
	goto L235
L335:
	;
	v1654 = base.F64_sub(v1646, base.F64_convert_i64_s(v1652))
	if base.F64_gt(v1654, float64(0.5)) != 0 {
		goto L340
	} else {
		goto L341
	}
L336:
	;
	v1650 = base.I64_trunc_f64_s(v1646)
	v1652 = v1650
	goto L335
L337:
	;
	goto L338
L338:
	;
	v1652 = int64(-9223372036854775807 - 1)
	goto L335
L339:
	;
	v1666 = v1665 + v1636
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1666
	if base.B2i32(v1665 < int64(0))^base.B2i32(v1666 < v1636) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L344
	}
L340:
	;
	v1665 = v1652 + int64(1)
	goto L339
L341:
	;
	goto L342
L342:
	;
	if base.F64_lt(v1654, float64(-0.5)) == int32(0) {
		v1665 = v1652
		goto L339
	} else {
		goto L343
	}
L343:
	;
	v1665 = v1652 - int64(1)
	goto L339
L344:
	;
	goto L334
L345:
	;
	v2195 = int32(-2)
	goto L7
L346:
	;
	goto L347
L347:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v1685 = base.I32_wrap_i64(v1253)
	v1686 = v1684 + v1685
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1686
	if base.B2i32(v1685 < int32(0))^base.B2i32(v1686 < v1684) != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2195 = int32(-2)
	goto L7
L349:
	;
	goto L350
L350:
	;
	if base.F64_ne(v1255, float64(0)) != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1696 = base.F64_mul(v1255, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v1696), float64(9.223372036854776e+18)) != 0 {
		goto L355
	} else {
		goto L356
	}
L352:
	;
	goto L353
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(8)
	goto L235
L354:
	;
	v1704 = base.F64_sub(v1696, base.F64_convert_i64_s(v1702))
	if base.F64_gt(v1704, float64(0.5)) != 0 {
		goto L359
	} else {
		goto L360
	}
L355:
	;
	v1700 = base.I64_trunc_f64_s(v1696)
	v1702 = v1700
	goto L354
L356:
	;
	goto L357
L357:
	;
	v1702 = int64(-9223372036854775807 - 1)
	goto L354
L358:
	;
	v1716 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1717 = v1716 + v1715
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1717
	if base.B2i32(v1715 < int64(0))^base.B2i32(v1717 < v1716) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L363
	}
L359:
	;
	v1715 = v1702 + int64(1)
	goto L358
L360:
	;
	goto L361
L361:
	;
	if base.F64_lt(v1704, float64(-0.5)) == int32(0) {
		v1715 = v1702
		goto L358
	} else {
		goto L362
	}
L362:
	;
	v1715 = v1702 - int64(1)
	goto L358
L363:
	;
	goto L353
L364:
	;
	v2195 = int32(-2)
	goto L7
L365:
	;
	goto L366
L366:
	;
	v1737 = v1253 * int64(7)
	v1741 = base.I32_wrap_i64(v1737)
	if base.I32_wrap_i64(int64(base.Ui64(v1737)>>(uint(int64(32))%64))) != v1741>>(uint(int32(31))%32) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v2195 = int32(-2)
	goto L7
L368:
	;
	goto L369
L369:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v1747 = v1746 + v1741
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1747
	if base.B2i32(v1741 < int32(0))^base.B2i32(v1747 < v1746) != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v2195 = int32(-2)
	goto L7
L371:
	;
	goto L372
L372:
	;
	if base.F64_eq(v1255, float64(0)) != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(16777216)
	v2081 = int32(0)
	v2082 = int32(22)
	v2099 = v275
	goto L23
L374:
	;
	v1757 = base.F64_mul(v1255, float64(7))
	if base.F64_lt(base.F64_abs(v1757), float64(2.147483648e+09)) != 0 {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1764 = v1763 + v1747
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1764
	if base.B2i32(v1763 < int32(0))^base.B2i32(v1764 < v1747) != 0 {
		goto L379
	} else {
		goto L380
	}
L376:
	;
	v1761 = base.I32_trunc_f64_s(v1757)
	v1763 = v1761
	goto L375
L377:
	;
	goto L378
L378:
	;
	v1763 = int32(-2147483648)
	goto L375
L379:
	;
	v2195 = int32(-2)
	goto L7
L380:
	;
	goto L381
L381:
	;
	v1772 = base.F64_sub(v1757, base.F64_convert_i32_s(v1763))
	if base.F64_eq(v1772, float64(0)) != 0 {
		goto L373
	} else {
		goto L382
	}
L382:
	;
	v1776 = base.F64_mul(v1772, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v1776), float64(9.223372036854776e+18)) != 0 {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	v1784 = base.F64_sub(v1776, base.F64_convert_i64_s(v1782))
	if base.F64_gt(v1784, float64(0.5)) != 0 {
		goto L388
	} else {
		goto L389
	}
L384:
	;
	v1780 = base.I64_trunc_f64_s(v1776)
	v1782 = v1780
	goto L383
L385:
	;
	goto L386
L386:
	;
	v1782 = int64(-9223372036854775807 - 1)
	goto L383
L387:
	;
	v1796 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1797 = v1796 + v1795
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1797
	if base.B2i32(v1795 < int64(0))^base.B2i32(v1797 < v1796) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L392
	}
L388:
	;
	v1795 = v1782 + int64(1)
	goto L387
L389:
	;
	goto L390
L390:
	;
	if base.F64_lt(v1784, float64(-0.5)) == int32(0) {
		v1795 = v1782
		goto L387
	} else {
		goto L391
	}
L391:
	;
	v1795 = v1782 - int64(1)
	goto L387
L392:
	;
	goto L373
L393:
	;
	v2195 = int32(-2)
	goto L7
L394:
	;
	goto L395
L395:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v1820 = base.I32_wrap_i64(v1253)
	v1821 = v1819 + v1820
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v1821
	if base.B2i32(v1820 < int32(0))^base.B2i32(v1821 < v1819) != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2195 = int32(-2)
	goto L7
L397:
	;
	goto L398
L398:
	;
	if base.F64_eq(v1255, float64(0)) != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(2)
	v2081 = int32(0)
	v2082 = int32(23)
	v2099 = v275
	goto L23
L400:
	;
	v1831 = base.F64_mul(v1255, float64(30))
	if base.F64_lt(base.F64_abs(v1831), float64(2.147483648e+09)) != 0 {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v1839 = v1837 + v1838
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1839
	if base.B2i32(v1837 < int32(0))^base.B2i32(v1839 < v1838) != 0 {
		goto L405
	} else {
		goto L406
	}
L402:
	;
	v1835 = base.I32_trunc_f64_s(v1831)
	v1837 = v1835
	goto L401
L403:
	;
	goto L404
L404:
	;
	v1837 = int32(-2147483648)
	goto L401
L405:
	;
	v2195 = int32(-2)
	goto L7
L406:
	;
	goto L407
L407:
	;
	v1847 = base.F64_sub(v1831, base.F64_convert_i32_s(v1837))
	if base.F64_eq(v1847, float64(0)) != 0 {
		goto L399
	} else {
		goto L408
	}
L408:
	;
	v1851 = base.F64_mul(v1847, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v1851), float64(9.223372036854776e+18)) != 0 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v1859 = base.F64_sub(v1851, base.F64_convert_i64_s(v1857))
	if base.F64_gt(v1859, float64(0.5)) != 0 {
		goto L414
	} else {
		goto L415
	}
L410:
	;
	v1855 = base.I64_trunc_f64_s(v1851)
	v1857 = v1855
	goto L409
L411:
	;
	goto L412
L412:
	;
	v1857 = int64(-9223372036854775807 - 1)
	goto L409
L413:
	;
	v1871 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1872 = v1871 + v1870
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1872
	if base.B2i32(v1870 < int64(0))^base.B2i32(v1872 < v1871) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L418
	}
L414:
	;
	v1870 = v1857 + int64(1)
	goto L413
L415:
	;
	goto L416
L416:
	;
	if base.F64_lt(v1859, float64(-0.5)) == int32(0) {
		v1870 = v1857
		goto L413
	} else {
		goto L417
	}
L417:
	;
	v1870 = v1857 - int64(1)
	goto L413
L418:
	;
	goto L399
L419:
	;
	v2195 = int32(-2)
	goto L7
L420:
	;
	goto L421
L421:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v1896 = base.I32_wrap_i64(v1253)
	v1897 = v1895 + v1896
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v1897
	if base.B2i32(v1896 < int32(0))^base.B2i32(v1897 < v1895) != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v2195 = int32(-2)
	goto L7
L423:
	;
	goto L424
L424:
	;
	v1906 = base.F64_nearest(base.F64_mul(v1255, float64(12)))
	if base.F64_lt(base.F64_abs(v1906), float64(2.147483648e+09)) != 0 {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v1914 = v1912 + v1913
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v1914
	if base.B2i32(v1912 < int32(0))^base.B2i32(v1914 < v1913) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L429
	}
L426:
	;
	v1910 = base.I32_trunc_f64_s(v1906)
	v1912 = v1910
	goto L425
L427:
	;
	goto L428
L428:
	;
	v1912 = int32(-2147483648)
	goto L425
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(4)
	v2081 = int32(0)
	v2082 = int32(25)
	v2099 = v275
	goto L23
L430:
	;
	v2195 = int32(-2)
	goto L7
L431:
	;
	goto L432
L432:
	;
	v1931 = v1253 * int64(10)
	v1935 = base.I32_wrap_i64(v1931)
	if base.I32_wrap_i64(int64(base.Ui64(v1931)>>(uint(int64(32))%64))) != v1935>>(uint(int32(31))%32) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2195 = int32(-2)
	goto L7
L434:
	;
	goto L435
L435:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v1941 = v1940 + v1935
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v1941
	if base.B2i32(v1935 < int32(0))^base.B2i32(v1941 < v1940) != 0 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v2195 = int32(-2)
	goto L7
L437:
	;
	goto L438
L438:
	;
	v1952 = base.F64_nearest(base.F64_mul(base.F64_mul(v1255, float64(10)), float64(12)))
	if base.F64_lt(base.F64_abs(v1952), float64(2.147483648e+09)) != 0 {
		goto L440
	} else {
		goto L441
	}
L439:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v1960 = v1958 + v1959
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v1960
	if base.B2i32(v1958 < int32(0))^base.B2i32(v1960 < v1959) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L443
	}
L440:
	;
	v1956 = base.I32_trunc_f64_s(v1952)
	v1958 = v1956
	goto L439
L441:
	;
	goto L442
L442:
	;
	v1958 = int32(-2147483648)
	goto L439
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(33554432)
	v2081 = int32(0)
	v2082 = int32(26)
	v2099 = v275
	goto L23
L444:
	;
	v2195 = int32(-2)
	goto L7
L445:
	;
	goto L446
L446:
	;
	v1977 = v1253 * int64(100)
	v1981 = base.I32_wrap_i64(v1977)
	if base.I32_wrap_i64(int64(base.Ui64(v1977)>>(uint(int64(32))%64))) != v1981>>(uint(int32(31))%32) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2195 = int32(-2)
	goto L7
L448:
	;
	goto L449
L449:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v1987 = v1986 + v1981
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v1987
	if base.B2i32(v1981 < int32(0))^base.B2i32(v1987 < v1986) != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2195 = int32(-2)
	goto L7
L451:
	;
	goto L452
L452:
	;
	v1998 = base.F64_nearest(base.F64_mul(base.F64_mul(v1255, float64(100)), float64(12)))
	if base.F64_lt(base.F64_abs(v1998), float64(2.147483648e+09)) != 0 {
		goto L454
	} else {
		goto L455
	}
L453:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v2006 = v2004 + v2005
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v2006
	if base.B2i32(v2004 < int32(0))^base.B2i32(v2006 < v2005) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L457
	}
L454:
	;
	v2002 = base.I32_trunc_f64_s(v1998)
	v2004 = v2002
	goto L453
L455:
	;
	goto L456
L456:
	;
	v2004 = int32(-2147483648)
	goto L453
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(67108864)
	v2081 = int32(0)
	v2082 = int32(27)
	v2099 = v275
	goto L23
L458:
	;
	v2195 = int32(-2)
	goto L7
L459:
	;
	goto L460
L460:
	;
	v2023 = v1253 * int64(1000)
	v2027 = base.I32_wrap_i64(v2023)
	if base.I32_wrap_i64(int64(base.Ui64(v2023)>>(uint(int64(32))%64))) != v2027>>(uint(int32(31))%32) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2195 = int32(-2)
	goto L7
L462:
	;
	goto L463
L463:
	;
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v2033 = v2032 + v2027
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v2033
	if base.B2i32(v2027 < int32(0))^base.B2i32(v2033 < v2032) != 0 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2195 = int32(-2)
	goto L7
L465:
	;
	goto L466
L466:
	;
	v2044 = base.F64_nearest(base.F64_mul(base.F64_mul(v1255, float64(1000)), float64(12)))
	if base.F64_lt(base.F64_abs(v2044), float64(2.147483648e+09)) != 0 {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v2052 = v2050 + v2051
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v2052
	if base.B2i32(v2050 < int32(0))^base.B2i32(v2052 < v2051) != 0 {
		v2195 = int32(-2)
		goto L7
	} else {
		goto L471
	}
L468:
	;
	v2048 = base.I32_trunc_f64_s(v2044)
	v2050 = v2048
	goto L467
L469:
	;
	goto L470
L470:
	;
	v2050 = int32(-2147483648)
	goto L467
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(134217728)
	v2081 = int32(0)
	v2082 = int32(28)
	v2099 = v275
	goto L23
L472:
	;
	v2195 = int32(-1)
	goto L7
L473:
	;
	goto L474
L474:
	;
	v2122 = v2081
	v2123 = v2082
	v2138 = v2109 | v273
	v2140 = v2099
	goto L22
L475:
	;
	goto L21
L476:
	;
	if v2140 == int32(0) {
		v2195 = v2158
		goto L7
	} else {
		goto L477
	}
L477:
	;
	v2161 = int32(-2)
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	if v2162 == int64(-9223372036854775807-1) {
		v2195 = v2161
		goto L7
	} else {
		goto L478
	}
L478:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v2165 == int32(-2147483648) {
		v2195 = v2161
		goto L7
	} else {
		goto L479
	}
L479:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v2168 == int32(-2147483648) {
		v2195 = v2161
		goto L7
	} else {
		goto L480
	}
L480:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if v2171 == int32(-2147483648) {
		v2195 = v2161
		goto L7
	} else {
		goto L481
	}
L481:
	;
	v2174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v2174 - v2171
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v2174 - v2168
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v2174 - v2165
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0) - v2162
	v2195 = v2174
	goto L7
L482:
	;
	v2266 = int32(0)
	v2267 = m.G0
	v2269 = v2267 - int32(112)
	m.G0 = v2269
	*(*int32)(unsafe.Add(mBase, uint32(v40+int32(500)))) = int32(17)
	v2276 = v40 + int32(504)
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+16)) = v2266
	v2279 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2276)+8)) = v2279
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v2279
	v2283 = int32(-1)
	v2284 = F_strlen(m, v44)
	mBase = m.M
	if base.Ui32(v2284) < base.Ui32(int32(2)) {
		v3834 = v2283
		goto L485
	} else {
		goto L486
	}
L483:
	;
	v3875 = v2240
	goto L484
L484:
	;
	if v3875 != 0 {
		goto L910
	} else {
		goto L911
	}
L485:
	;
	m.G0 = v2269 + int32(112)
	v3875 = v3834
	goto L484
L486:
	;
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v2287 != int32(80) {
		v3834 = v2283
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v2291 = v44 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+108)) = v2291
	v2293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2291))))
	if v2293 != 0 {
		goto L492
	} else {
		goto L493
	}
L488:
	;
	v3834 = v3817
	goto L485
L489:
	;
	v3817 = int32(0)
	goto L488
L490:
	;
	if v2310&int32(1) != 0 {
		v3834 = v2283
		goto L485
	} else {
		goto L843
	}
L491:
	;
	if v2393 != 0 {
		v3817 = v2341
		goto L488
	} else {
		goto L802
	}
L492:
	;
	v2303 = v2293
	v2309 = int32(1)
	v2310 = v2266
	v2311 = v2291
	goto L495
L493:
	;
	goto L494
L494:
	;
	v3834 = int32(0)
	goto L485
L495:
	;
	if v2303&int32(255) == int32(84) {
		goto L498
	} else {
		goto L499
	}
L496:
	;
	goto L494
L497:
	;
	v3295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3293))))
	if v3295 != 0 {
		v2303 = v3295
		v2309 = v3291
		v2310 = v3292
		v2311 = v3293
		goto L495
	} else {
		goto L801
	}
L498:
	;
	v2337 = v2311 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+108)) = v2337
	v2339 = int32(0)
	v3291 = v2339
	v3292 = v2339
	v3293 = v2337
	goto L497
L499:
	;
	goto L500
L500:
	;
	v2341 = int32(-1)
	if base.Ui32(int32(10)) <= base.Ui32((v2303-int32(48))&int32(255)) {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	if base.Ui32(int32(1)) < base.Ui32((v2303-int32(45))&int32(255)) {
		v3817 = v2341
		goto L488
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v2359 = F_strtod(m, v2311, v2269+int32(108))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L30
	} else {
		goto L505
	}
L504:
	;
	goto L503
L505:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+108))
	if v2361 == v2311 {
		v3817 = v2341
		goto L488
	} else {
		goto L506
	}
L506:
	;
	v2364 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v2364 != 0 {
		v3817 = v2341
		goto L488
	} else {
		goto L507
	}
L507:
	;
	v2365 = base.F64_abs(v2359)
	if base.F64_gt(v2365, float64(1e+15)) != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3817 = int32(-2)
	goto L488
L509:
	;
	goto L510
L510:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v2365)) {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v3817 = int32(-2)
	goto L488
L512:
	;
	goto L513
L513:
	;
	v2374 = v2361 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+108)) = v2374
	if base.F64_ge(v2359, float64(0)) != 0 {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2269)+96)) = v2388
	v2391 = base.F64_sub(v2359, base.F64_convert_i64_s(v2388))
	*(*float64)(unsafe.Add(mBase, uint32(v2269)+88)) = v2391
	v2393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2361))))
	if v2309&int32(1) != 0 {
		goto L521
	} else {
		goto L522
	}
L515:
	;
	v2382 = base.F64_floor(v2359)
	goto L517
L516:
	;
	v2382 = base.F64_neg(base.F64_floor(base.F64_neg(v2359)))
	goto L517
L517:
	;
	if base.F64_lt(base.F64_abs(v2382), float64(9.223372036854776e+18)) != 0 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2386 = base.I64_trunc_f64_s(v2382)
	v2388 = v2386
	goto L514
L519:
	;
	goto L520
L520:
	;
	v2388 = int64(-9223372036854775807 - 1)
	goto L514
L521:
	;
	switch v2393 - int32(45) {
	case 0:
		goto L524
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 40, 41, 43:
		v3817 = v2341
		goto L488
	case 23:
		goto L526
	case 32:
		goto L528
	case 39:
		goto L530
	case 42:
		goto L527
	case 44:
		goto L529
	default:
		goto L531
	}
L522:
	;
	goto L523
L523:
	;
	switch v2393 - int32(58) {
	case 0:
		goto L490
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 20, 21, 22, 23, 24:
		v3817 = v2341
		goto L488
	case 14:
		goto L746
	case 19:
		goto L745
	case 25:
		goto L744
	default:
		goto L491
	}
L524:
	;
	if v2310&int32(1) != 0 {
		v3817 = v2341
		goto L488
	} else {
		goto L670
	}
L525:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+16))
	v2734 = base.I32_wrap_i64(v2492)
	v2735 = v2733 + v2734
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+16)) = v2735
	if base.B2i32(v2734 < int32(0))^base.B2i32(v2735 < v2733) != 0 {
		goto L647
	} else {
		goto L648
	}
L526:
	;
	if base.Ui64(v2388-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L628
	} else {
		goto L629
	}
L527:
	;
	if base.Ui64(v2388-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L596
	} else {
		goto L597
	}
L528:
	;
	if base.Ui64(v2388-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L567
	} else {
		goto L568
	}
L529:
	;
	if base.Ui64(v2388-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L556
	} else {
		goto L557
	}
L530:
	;
	v2398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2311))))
	v2401 = v2311 + base.B2i32(v2398 == int32(45))
	v2402 = int32(553196)
	v2406 = m.G0
	v2408 = v2406 - int32(32)
	v2409 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2408)+24)) = v2409
	*(*int64)(unsafe.Add(mBase, uint32(v2408)+16)) = v2409
	*(*int64)(unsafe.Add(mBase, uint32(v2408)+8)) = v2409
	*(*int64)(unsafe.Add(mBase, uint32(v2408))) = v2409
	v2417 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v2417 == int32(0) {
		goto L534
	} else {
		goto L535
	}
L531:
	;
	if v2393 != 0 {
		v3817 = v2341
		goto L488
	} else {
		goto L532
	}
L532:
	;
	goto L530
L533:
	;
	if (v2310|base.B2i32(v2485 != int32(8)))&int32(1) != 0 {
		goto L524
	} else {
		goto L554
	}
L534:
	;
	v2485 = int32(0)
	goto L533
L535:
	;
	goto L536
L536:
	;
	v2421 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v2421 == int32(0) {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v2425 = v2401
	goto L540
L538:
	;
	goto L539
L539:
	;
	v2435 = v2402
	v2436 = v2417
	goto L543
L540:
	;
	v2431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2425))))
	if v2431 == v2417 {
		v2425 = v2425 + int32(1)
		goto L540
	} else {
		goto L542
	}
L541:
	;
	v2485 = v2425 - v2401
	goto L533
L542:
	;
	goto L541
L543:
	;
	v2443 = v2408 + int32(base.Ui32(v2436)>>(uint(int32(3))%32))&int32(28)
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2443)))
	v2445 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2443))) = v2444 | v2445<<(uint(v2436)%32)
	v2449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2435)+1)))
	if v2449 != 0 {
		v2435 = v2435 + v2445
		v2436 = v2449
		goto L543
	} else {
		goto L545
	}
L544:
	;
	v2452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2401))))
	if v2452 == int32(0) {
		v2477 = v2401
		goto L546
	} else {
		goto L547
	}
L545:
	;
	goto L544
L546:
	;
	v2485 = v2477 - v2401
	goto L533
L547:
	;
	v2456 = v2401
	v2457 = v2452
	goto L548
L548:
	;
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v2408+int32(base.Ui32(v2457)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v2465)>>(uint(v2457)%32))&int32(1) == int32(0) {
		goto L550
	} else {
		goto L551
	}
L549:
	;
	v2477 = v2473
	goto L546
L550:
	;
	v2477 = v2456
	goto L546
L551:
	;
	goto L552
L552:
	;
	v2471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2456)+1)))
	v2473 = v2456 + int32(1)
	if v2471 != 0 {
		v2456 = v2473
		v2457 = v2471
		goto L548
	} else {
		goto L553
	}
L553:
	;
	goto L549
L554:
	;
	v2492 = base.I64_div_s(v2388, int64(10000))
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v2492-int64(2147483648)) {
		goto L525
	} else {
		goto L555
	}
L555:
	;
	v3817 = int32(-2)
	goto L488
L556:
	;
	v3817 = int32(-2)
	goto L488
L557:
	;
	goto L558
L558:
	;
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+16))
	v2504 = base.I32_wrap_i64(v2388)
	v2505 = v2503 + v2504
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+16)) = v2505
	if base.B2i32(v2504 < int32(0))^base.B2i32(v2505 < v2503) != 0 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v3817 = int32(-2)
	goto L488
L560:
	;
	goto L561
L561:
	;
	v2514 = base.F64_nearest(base.F64_mul(v2391, float64(12)))
	if base.F64_lt(base.F64_abs(v2514), float64(2.147483648e+09)) != 0 {
		goto L563
	} else {
		goto L564
	}
L562:
	;
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	v2522 = v2520 + v2521
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+12)) = v2522
	v2524 = int32(1)
	v2527 = int32(0)
	if base.B2i32(v2520 < v2527)^base.B2i32(v2522 < v2521) == v2527 {
		v3291 = v2524
		v3292 = v2524
		v3293 = v2374
		goto L497
	} else {
		goto L566
	}
L563:
	;
	v2518 = base.I32_trunc_f64_s(v2514)
	v2520 = v2518
	goto L562
L564:
	;
	goto L565
L565:
	;
	v2520 = int32(-2147483648)
	goto L562
L566:
	;
	v3817 = int32(-2)
	goto L488
L567:
	;
	v3817 = int32(-2)
	goto L488
L568:
	;
	goto L569
L569:
	;
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	v2539 = base.I32_wrap_i64(v2388)
	v2540 = v2538 + v2539
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+12)) = v2540
	if base.B2i32(v2539 < int32(0))^base.B2i32(v2540 < v2538) != 0 {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v3817 = int32(-2)
	goto L488
L571:
	;
	goto L572
L572:
	;
	v2547 = int32(1)
	if base.F64_eq(v2391, float64(0)) != 0 {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v3291 = int32(1)
	v3292 = v2547
	v3293 = v2374
	goto L497
L574:
	;
	goto L575
L575:
	;
	v2552 = base.F64_mul(v2391, float64(30))
	if base.F64_lt(base.F64_abs(v2552), float64(2.147483648e+09)) != 0 {
		goto L577
	} else {
		goto L578
	}
L576:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	v2560 = v2558 + v2559
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2560
	if base.B2i32(v2558 < int32(0))^base.B2i32(v2560 < v2559) != 0 {
		goto L580
	} else {
		goto L581
	}
L577:
	;
	v2556 = base.I32_trunc_f64_s(v2552)
	v2558 = v2556
	goto L576
L578:
	;
	goto L579
L579:
	;
	v2558 = int32(-2147483648)
	goto L576
L580:
	;
	v3817 = int32(-2)
	goto L488
L581:
	;
	goto L582
L582:
	;
	v2568 = base.F64_sub(v2552, base.F64_convert_i32_s(v2558))
	if base.F64_eq(v2568, float64(0)) != 0 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v3291 = int32(1)
	v3292 = v2547
	v3293 = v2374
	goto L497
L584:
	;
	goto L585
L585:
	;
	v2573 = base.F64_mul(v2568, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2573), float64(9.223372036854776e+18)) != 0 {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	v2581 = base.F64_sub(v2573, base.F64_convert_i64_s(v2579))
	if base.F64_gt(v2581, float64(0.5)) != 0 {
		goto L591
	} else {
		goto L592
	}
L587:
	;
	v2577 = base.I64_trunc_f64_s(v2573)
	v2579 = v2577
	goto L586
L588:
	;
	goto L589
L589:
	;
	v2579 = int64(-9223372036854775807 - 1)
	goto L586
L590:
	;
	v2593 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v2594 = v2593 + v2592
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v2594
	if base.B2i32(v2592 < int64(0))^base.B2i32(v2594 < v2593) == int32(0) {
		v3291 = int32(1)
		v3292 = v2547
		v3293 = v2374
		goto L497
	} else {
		goto L595
	}
L591:
	;
	v2592 = v2579 + int64(1)
	goto L590
L592:
	;
	goto L593
L593:
	;
	if base.F64_lt(v2581, float64(-0.5)) == int32(0) {
		v2592 = v2579
		goto L590
	} else {
		goto L594
	}
L594:
	;
	v2592 = v2579 - int64(1)
	goto L590
L595:
	;
	v3817 = int32(-2)
	goto L488
L596:
	;
	v3817 = int32(-2)
	goto L488
L597:
	;
	goto L598
L598:
	;
	v2610 = v2388 * int64(7)
	v2614 = base.I32_wrap_i64(v2610)
	if base.I32_wrap_i64(int64(base.Ui64(v2610)>>(uint(int64(32))%64))) != v2614>>(uint(int32(31))%32) {
		goto L599
	} else {
		goto L600
	}
L599:
	;
	v3817 = int32(-2)
	goto L488
L600:
	;
	goto L601
L601:
	;
	v2619 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	v2620 = v2619 + v2614
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2620
	if base.B2i32(v2614 < int32(0))^base.B2i32(v2620 < v2619) != 0 {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v3817 = int32(-2)
	goto L488
L603:
	;
	goto L604
L604:
	;
	v2627 = int32(1)
	if base.F64_eq(v2391, float64(0)) != 0 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	v3291 = int32(1)
	v3292 = v2627
	v3293 = v2374
	goto L497
L606:
	;
	goto L607
L607:
	;
	v2632 = base.F64_mul(v2391, float64(7))
	if base.F64_lt(base.F64_abs(v2632), float64(2.147483648e+09)) != 0 {
		goto L609
	} else {
		goto L610
	}
L608:
	;
	v2639 = v2620 + v2638
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2639
	if base.B2i32(v2638 < int32(0))^base.B2i32(v2639 < v2620) != 0 {
		goto L612
	} else {
		goto L613
	}
L609:
	;
	v2636 = base.I32_trunc_f64_s(v2632)
	v2638 = v2636
	goto L608
L610:
	;
	goto L611
L611:
	;
	v2638 = int32(-2147483648)
	goto L608
L612:
	;
	v3817 = int32(-2)
	goto L488
L613:
	;
	goto L614
L614:
	;
	v2647 = base.F64_sub(v2632, base.F64_convert_i32_s(v2638))
	if base.F64_eq(v2647, float64(0)) != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3291 = int32(1)
	v3292 = v2627
	v3293 = v2374
	goto L497
L616:
	;
	goto L617
L617:
	;
	v2652 = base.F64_mul(v2647, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2652), float64(9.223372036854776e+18)) != 0 {
		goto L619
	} else {
		goto L620
	}
L618:
	;
	v2660 = base.F64_sub(v2652, base.F64_convert_i64_s(v2658))
	if base.F64_gt(v2660, float64(0.5)) != 0 {
		goto L623
	} else {
		goto L624
	}
L619:
	;
	v2656 = base.I64_trunc_f64_s(v2652)
	v2658 = v2656
	goto L618
L620:
	;
	goto L621
L621:
	;
	v2658 = int64(-9223372036854775807 - 1)
	goto L618
L622:
	;
	v2672 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v2673 = v2672 + v2671
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v2673
	if base.B2i32(v2671 < int64(0))^base.B2i32(v2673 < v2672) == int32(0) {
		v3291 = int32(1)
		v3292 = v2627
		v3293 = v2374
		goto L497
	} else {
		goto L627
	}
L623:
	;
	v2671 = v2658 + int64(1)
	goto L622
L624:
	;
	goto L625
L625:
	;
	if base.F64_lt(v2660, float64(-0.5)) == int32(0) {
		v2671 = v2658
		goto L622
	} else {
		goto L626
	}
L626:
	;
	v2671 = v2658 - int64(1)
	goto L622
L627:
	;
	v3817 = int32(-2)
	goto L488
L628:
	;
	v3817 = int32(-2)
	goto L488
L629:
	;
	goto L630
L630:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	v2689 = base.I32_wrap_i64(v2388)
	v2690 = v2688 + v2689
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2690
	if base.B2i32(v2689 < int32(0))^base.B2i32(v2690 < v2688) != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v3817 = int32(-2)
	goto L488
L632:
	;
	goto L633
L633:
	;
	v2697 = int32(1)
	if base.F64_eq(v2391, float64(0)) != 0 {
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v3291 = int32(1)
	v3292 = v2697
	v3293 = v2374
	goto L497
L635:
	;
	goto L636
L636:
	;
	v2702 = base.F64_mul(v2391, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2702), float64(9.223372036854776e+18)) != 0 {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	v2710 = base.F64_sub(v2702, base.F64_convert_i64_s(v2708))
	if base.F64_gt(v2710, float64(0.5)) != 0 {
		goto L642
	} else {
		goto L643
	}
L638:
	;
	v2706 = base.I64_trunc_f64_s(v2702)
	v2708 = v2706
	goto L637
L639:
	;
	goto L640
L640:
	;
	v2708 = int64(-9223372036854775807 - 1)
	goto L637
L641:
	;
	v2722 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v2723 = v2722 + v2721
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v2723
	if base.B2i32(v2721 < int64(0))^base.B2i32(v2723 < v2722) == int32(0) {
		v3291 = int32(1)
		v3292 = v2697
		v3293 = v2374
		goto L497
	} else {
		goto L646
	}
L642:
	;
	v2721 = v2708 + int64(1)
	goto L641
L643:
	;
	goto L644
L644:
	;
	if base.F64_lt(v2710, float64(-0.5)) == int32(0) {
		v2721 = v2708
		goto L641
	} else {
		goto L645
	}
L645:
	;
	v2721 = v2708 - int64(1)
	goto L641
L646:
	;
	v3817 = int32(-2)
	goto L488
L647:
	;
	v3817 = int32(-2)
	goto L488
L648:
	;
	goto L649
L649:
	;
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	v2743 = int64(100)
	v2744 = base.I64_div_s(v2388, v2743)
	v2746 = base.I64_rem_s(v2744, v2743)
	v2747 = base.I32_wrap_i64(v2746)
	v2748 = v2742 + v2747
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+12)) = v2748
	if base.B2i32(v2747 < int32(0))^base.B2i32(v2748 < v2742) != 0 {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	v3817 = int32(-2)
	goto L488
L651:
	;
	goto L652
L652:
	;
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	v2757 = base.I64_rem_s(v2388, int64(100))
	v2758 = base.I32_wrap_i64(v2757)
	v2759 = v2755 + v2758
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2759
	if base.B2i32(v2758 < int32(0))^base.B2i32(v2759 < v2755) != 0 {
		goto L653
	} else {
		goto L654
	}
L653:
	;
	v3817 = int32(-2)
	goto L488
L654:
	;
	goto L655
L655:
	;
	if base.F64_ne(v2391, float64(0)) != 0 {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	v2769 = base.F64_mul(v2391, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2769), float64(9.223372036854776e+18)) != 0 {
		goto L660
	} else {
		goto L661
	}
L657:
	;
	goto L658
L658:
	;
	v2802 = int32(0)
	if v2393 == v2802 {
		goto L489
	} else {
		goto L669
	}
L659:
	;
	v2777 = base.F64_sub(v2769, base.F64_convert_i64_s(v2775))
	if base.F64_gt(v2777, float64(0.5)) != 0 {
		goto L664
	} else {
		goto L665
	}
L660:
	;
	v2773 = base.I64_trunc_f64_s(v2769)
	v2775 = v2773
	goto L659
L661:
	;
	goto L662
L662:
	;
	v2775 = int64(-9223372036854775807 - 1)
	goto L659
L663:
	;
	v2789 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v2790 = v2789 + v2788
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v2790
	if base.B2i32(v2788 < int64(0))^base.B2i32(v2790 < v2789) != 0 {
		v3817 = int32(-2)
		goto L488
	} else {
		goto L668
	}
L664:
	;
	v2788 = v2775 + int64(1)
	goto L663
L665:
	;
	goto L666
L666:
	;
	if base.F64_lt(v2777, float64(-0.5)) == int32(0) {
		v2788 = v2775
		goto L663
	} else {
		goto L667
	}
L667:
	;
	v2788 = v2775 - int64(1)
	goto L663
L668:
	;
	goto L658
L669:
	;
	v3291 = int32(0)
	v3292 = v2802
	v3293 = v2374
	goto L497
L670:
	;
	if base.Ui64(v2388-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L671
	} else {
		goto L672
	}
L671:
	;
	v3817 = int32(-2)
	goto L488
L672:
	;
	goto L673
L673:
	;
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+16))
	v2814 = base.I32_wrap_i64(v2388)
	v2815 = v2813 + v2814
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+16)) = v2815
	if base.B2i32(v2814 < int32(0))^base.B2i32(v2815 < v2813) != 0 {
		goto L674
	} else {
		goto L675
	}
L674:
	;
	v3817 = int32(-2)
	goto L488
L675:
	;
	goto L676
L676:
	;
	v2824 = base.F64_nearest(base.F64_mul(v2391, float64(12)))
	if base.F64_lt(base.F64_abs(v2824), float64(2.147483648e+09)) != 0 {
		goto L678
	} else {
		goto L679
	}
L677:
	;
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	v2832 = v2830 + v2831
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+12)) = v2832
	if base.B2i32(v2830 < int32(0))^base.B2i32(v2832 < v2831) != 0 {
		goto L681
	} else {
		goto L682
	}
L678:
	;
	v2828 = base.I32_trunc_f64_s(v2824)
	v2830 = v2828
	goto L677
L679:
	;
	goto L680
L680:
	;
	v2830 = int32(-2147483648)
	goto L677
L681:
	;
	v3817 = int32(-2)
	goto L488
L682:
	;
	goto L683
L683:
	;
	v2839 = int32(0)
	if v2393 == int32(84) {
		goto L684
	} else {
		goto L685
	}
L684:
	;
	v3291 = int32(0)
	v3292 = v2839
	v3293 = v2374
	goto L497
L685:
	;
	goto L686
L686:
	;
	if v2393 == int32(0) {
		v3817 = v2393
		goto L488
	} else {
		goto L687
	}
L687:
	;
	v2851 = F_ParseISO8601Number(m, v2374, v2269+int32(108), v2269+int32(96), v2269+int32(88))
	mBase = m.M
	v2852 = m.ExcPending
	if v2852 != 0 {
		goto L30
	} else {
		goto L688
	}
L688:
	;
	if v2851 != 0 {
		v3817 = v2851
		goto L488
	} else {
		goto L689
	}
L689:
	;
	v2853 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+96))
	if base.Ui64(v2853-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v3817 = int32(-2)
	goto L488
L691:
	;
	goto L692
L692:
	;
	v2859 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	v2860 = base.I32_wrap_i64(v2853)
	v2861 = v2859 + v2860
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+12)) = v2861
	if base.B2i32(v2860 < int32(0))^base.B2i32(v2861 < v2859) != 0 {
		goto L693
	} else {
		goto L694
	}
L693:
	;
	v3817 = int32(-2)
	goto L488
L694:
	;
	goto L695
L695:
	;
	v2869 = *(*float64)(unsafe.Add(mBase, uint32(v2269)+88))
	if base.F64_eq(v2869, float64(0)) != 0 {
		v2930 = int32(1)
		goto L696
	} else {
		goto L697
	}
L696:
	;
	if v2930 == int32(0) {
		goto L713
	} else {
		goto L714
	}
L697:
	;
	v2873 = base.F64_mul(v2869, float64(30))
	if base.F64_lt(base.F64_abs(v2873), float64(2.147483648e+09)) != 0 {
		goto L699
	} else {
		goto L700
	}
L698:
	;
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	v2881 = v2879 + v2880
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2881
	v2883 = int32(0)
	if base.B2i32(v2879 < v2883)^base.B2i32(v2881 < v2880) != 0 {
		v2930 = v2883
		goto L696
	} else {
		goto L702
	}
L699:
	;
	v2877 = base.I32_trunc_f64_s(v2873)
	v2879 = v2877
	goto L698
L700:
	;
	goto L701
L701:
	;
	v2879 = int32(-2147483648)
	goto L698
L702:
	;
	v2890 = base.F64_sub(v2873, base.F64_convert_i32_s(v2879))
	if base.F64_eq(v2890, float64(0)) != 0 {
		v2930 = int32(1)
		goto L696
	} else {
		goto L703
	}
L703:
	;
	v2894 = base.F64_mul(v2890, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2894), float64(9.223372036854776e+18)) != 0 {
		goto L705
	} else {
		goto L706
	}
L704:
	;
	v2902 = base.F64_sub(v2894, base.F64_convert_i64_s(v2900))
	if base.F64_gt(v2902, float64(0.5)) != 0 {
		goto L709
	} else {
		goto L710
	}
L705:
	;
	v2898 = base.I64_trunc_f64_s(v2894)
	v2900 = v2898
	goto L704
L706:
	;
	goto L707
L707:
	;
	v2900 = int64(-9223372036854775807 - 1)
	goto L704
L708:
	;
	v2914 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v2915 = v2914 + v2913
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v2915
	v2930 = base.B2i32(base.B2i32(v2913 < int64(0))^base.B2i32(v2915 < v2914) == int32(0))
	goto L696
L709:
	;
	v2913 = v2900 + int64(1)
	goto L708
L710:
	;
	goto L711
L711:
	;
	if base.F64_lt(v2902, float64(-0.5)) == int32(0) {
		v2913 = v2900
		goto L708
	} else {
		goto L712
	}
L712:
	;
	v2913 = v2900 - int64(1)
	goto L708
L713:
	;
	v3817 = int32(-2)
	goto L488
L714:
	;
	goto L715
L715:
	;
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+108))
	v2935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2934))))
	if v2935 != int32(45) {
		goto L716
	} else {
		goto L717
	}
L716:
	;
	if v2935 == int32(84) {
		v3291 = int32(0)
		v3292 = v2839
		v3293 = v2934
		goto L497
	} else {
		goto L719
	}
L717:
	;
	goto L718
L718:
	;
	v2944 = v2934 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+108)) = v2944
	v2952 = F_ParseISO8601Number(m, v2944, v2269+int32(108), v2269+int32(96), v2269+int32(88))
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L30
	} else {
		goto L721
	}
L719:
	;
	if v2935 == int32(0) {
		v3817 = v2935
		goto L488
	} else {
		goto L720
	}
L720:
	;
	v3834 = v2283
	goto L485
L721:
	;
	if v2952 != 0 {
		v3817 = v2952
		goto L488
	} else {
		goto L722
	}
L722:
	;
	v2954 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+96))
	if base.Ui64(v2954-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L723
	} else {
		goto L724
	}
L723:
	;
	v3817 = int32(-2)
	goto L488
L724:
	;
	goto L725
L725:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	v2961 = base.I32_wrap_i64(v2954)
	v2962 = v2960 + v2961
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2962
	if base.B2i32(v2961 < int32(0))^base.B2i32(v2962 < v2960) != 0 {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	v3817 = int32(-2)
	goto L488
L727:
	;
	goto L728
L728:
	;
	v2969 = *(*float64)(unsafe.Add(mBase, uint32(v2269)+88))
	if base.F64_ne(v2969, float64(0)) != 0 {
		goto L729
	} else {
		goto L730
	}
L729:
	;
	v2973 = base.F64_mul(v2969, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2973), float64(9.223372036854776e+18)) != 0 {
		goto L733
	} else {
		goto L734
	}
L730:
	;
	goto L731
L731:
	;
	v3007 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+108))
	v3008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3007))))
	if v3008 == int32(84) {
		v3291 = int32(0)
		v3292 = v2839
		v3293 = v3007
		goto L497
	} else {
		goto L742
	}
L732:
	;
	v2981 = base.F64_sub(v2973, base.F64_convert_i64_s(v2979))
	if base.F64_gt(v2981, float64(0.5)) != 0 {
		goto L737
	} else {
		goto L738
	}
L733:
	;
	v2977 = base.I64_trunc_f64_s(v2973)
	v2979 = v2977
	goto L732
L734:
	;
	goto L735
L735:
	;
	v2979 = int64(-9223372036854775807 - 1)
	goto L732
L736:
	;
	v2993 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v2994 = v2993 + v2992
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v2994
	if base.B2i32(v2992 < int64(0))^base.B2i32(v2994 < v2993) != 0 {
		v3817 = int32(-2)
		goto L488
	} else {
		goto L741
	}
L737:
	;
	v2992 = v2979 + int64(1)
	goto L736
L738:
	;
	goto L739
L739:
	;
	if base.F64_lt(v2981, float64(-0.5)) == int32(0) {
		v2992 = v2979
		goto L736
	} else {
		goto L740
	}
L740:
	;
	v2992 = v2979 - int64(1)
	goto L736
L741:
	;
	goto L731
L742:
	;
	if v3008 == int32(0) {
		v3817 = v3008
		goto L488
	} else {
		goto L743
	}
L743:
	;
	v3834 = v2283
	goto L485
L744:
	;
	v3196 = v2269 + int32(40)
	v3199 = int64(1000000)
	v3200 = int64(0)
	v3205 = int64(32)
	v3208 = int64(base.Ui64(v2388) >> (uint(v3205) % 64))
	v3211 = int64(4294967295)
	v3214 = v2388 & v3211
	v3215 = v3199 * v3214
	v3219 = int64(base.Ui64(v3215)>>(uint(v3205)%64)) + v3199*v3208
	v3226 = v3214*v3200 + v3219&v3211
	*(*int64)(unsafe.Add(mBase, uint32(v3196)+8)) = v2388*v3200 + v2388>>(uint(int64(63))%64)*v3199 + v3200*v3208 + int64(base.Ui64(v3219)>>(uint(v3205)%64)) + int64(base.Ui64(v3226)>>(uint(v3205)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3196))) = v3215&v3211 | v3226<<(uint(v3205)%64)
	goto L783
L745:
	;
	v3106 = v2269 + int32(24)
	v3109 = int64(60000000)
	v3110 = int64(0)
	v3115 = int64(32)
	v3118 = int64(base.Ui64(v2388) >> (uint(v3115) % 64))
	v3121 = int64(4294967295)
	v3124 = v2388 & v3121
	v3125 = v3109 * v3124
	v3129 = int64(base.Ui64(v3125)>>(uint(v3115)%64)) + v3109*v3118
	v3136 = v3124*v3110 + v3129&v3121
	*(*int64)(unsafe.Add(mBase, uint32(v3106)+8)) = v2388*v3110 + v2388>>(uint(int64(63))%64)*v3109 + v3110*v3118 + int64(base.Ui64(v3129)>>(uint(v3115)%64)) + int64(base.Ui64(v3136)>>(uint(v3115)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3106))) = v3125&v3121 | v3136<<(uint(v3115)%64)
	goto L765
L746:
	;
	v3016 = v2269 + int32(8)
	v3019 = int64(3600000000)
	v3020 = int64(0)
	v3025 = int64(32)
	v3028 = int64(base.Ui64(v2388) >> (uint(v3025) % 64))
	v3031 = int64(4294967295)
	v3034 = v2388 & v3031
	v3035 = v3019 * v3034
	v3039 = int64(base.Ui64(v3035)>>(uint(v3025)%64)) + v3019*v3028
	v3046 = v3034*v3020 + v3039&v3031
	*(*int64)(unsafe.Add(mBase, uint32(v3016)+8)) = v2388*v3020 + v2388>>(uint(int64(63))%64)*v3019 + v3020*v3028 + int64(base.Ui64(v3039)>>(uint(v3025)%64)) + int64(base.Ui64(v3046)>>(uint(v3025)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3016))) = v3035&v3031 | v3046<<(uint(v3025)%64)
	goto L747
L747:
	;
	v3057 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+16))
	v3058 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+8))
	if v3057 != v3058>>(uint(int64(63))%64) {
		goto L748
	} else {
		goto L749
	}
L748:
	;
	v3817 = int32(-2)
	goto L488
L749:
	;
	goto L750
L750:
	;
	v3063 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v3064 = v3063 + v3058
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3064
	if base.B2i32(v3058 < int64(0))^base.B2i32(v3064 < v3063) != 0 {
		goto L751
	} else {
		goto L752
	}
L751:
	;
	v3817 = int32(-2)
	goto L488
L752:
	;
	goto L753
L753:
	;
	v3071 = int32(0)
	v3072 = int32(1)
	if base.F64_eq(v2391, float64(0)) != 0 {
		v3291 = v3071
		v3292 = v3072
		v3293 = v2374
		goto L497
	} else {
		goto L754
	}
L754:
	;
	v3076 = base.F64_mul(v2391, float64(3.6e+09))
	if base.F64_lt(base.F64_abs(v3076), float64(9.223372036854776e+18)) != 0 {
		goto L756
	} else {
		goto L757
	}
L755:
	;
	v3084 = base.F64_sub(v3076, base.F64_convert_i64_s(v3082))
	if base.F64_gt(v3084, float64(0.5)) != 0 {
		goto L760
	} else {
		goto L761
	}
L756:
	;
	v3080 = base.I64_trunc_f64_s(v3076)
	v3082 = v3080
	goto L755
L757:
	;
	goto L758
L758:
	;
	v3082 = int64(-9223372036854775807 - 1)
	goto L755
L759:
	;
	v3096 = v3095 + v3064
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3096
	if base.B2i32(v3095 < int64(0))^base.B2i32(v3096 < v3064) == int32(0) {
		v3291 = v3071
		v3292 = v3072
		v3293 = v2374
		goto L497
	} else {
		goto L764
	}
L760:
	;
	v3095 = v3082 + int64(1)
	goto L759
L761:
	;
	goto L762
L762:
	;
	if base.F64_lt(v3084, float64(-0.5)) == int32(0) {
		v3095 = v3082
		goto L759
	} else {
		goto L763
	}
L763:
	;
	v3095 = v3082 - int64(1)
	goto L759
L764:
	;
	v3817 = int32(-2)
	goto L488
L765:
	;
	v3147 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+32))
	v3148 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+24))
	if v3147 != v3148>>(uint(int64(63))%64) {
		goto L766
	} else {
		goto L767
	}
L766:
	;
	v3817 = int32(-2)
	goto L488
L767:
	;
	goto L768
L768:
	;
	v3153 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v3154 = v3153 + v3148
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3154
	if base.B2i32(v3148 < int64(0))^base.B2i32(v3154 < v3153) != 0 {
		goto L769
	} else {
		goto L770
	}
L769:
	;
	v3817 = int32(-2)
	goto L488
L770:
	;
	goto L771
L771:
	;
	v3161 = int32(0)
	v3162 = int32(1)
	if base.F64_eq(v2391, float64(0)) != 0 {
		v3291 = v3161
		v3292 = v3162
		v3293 = v2374
		goto L497
	} else {
		goto L772
	}
L772:
	;
	v3166 = base.F64_mul(v2391, float64(6e+07))
	if base.F64_lt(base.F64_abs(v3166), float64(9.223372036854776e+18)) != 0 {
		goto L774
	} else {
		goto L775
	}
L773:
	;
	v3174 = base.F64_sub(v3166, base.F64_convert_i64_s(v3172))
	if base.F64_gt(v3174, float64(0.5)) != 0 {
		goto L778
	} else {
		goto L779
	}
L774:
	;
	v3170 = base.I64_trunc_f64_s(v3166)
	v3172 = v3170
	goto L773
L775:
	;
	goto L776
L776:
	;
	v3172 = int64(-9223372036854775807 - 1)
	goto L773
L777:
	;
	v3186 = v3185 + v3154
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3186
	if base.B2i32(v3185 < int64(0))^base.B2i32(v3186 < v3154) == int32(0) {
		v3291 = v3161
		v3292 = v3162
		v3293 = v2374
		goto L497
	} else {
		goto L782
	}
L778:
	;
	v3185 = v3172 + int64(1)
	goto L777
L779:
	;
	goto L780
L780:
	;
	if base.F64_lt(v3174, float64(-0.5)) == int32(0) {
		v3185 = v3172
		goto L777
	} else {
		goto L781
	}
L781:
	;
	v3185 = v3172 - int64(1)
	goto L777
L782:
	;
	v3817 = int32(-2)
	goto L488
L783:
	;
	v3237 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+48))
	v3238 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+40))
	if v3237 != v3238>>(uint(int64(63))%64) {
		goto L784
	} else {
		goto L785
	}
L784:
	;
	v3817 = int32(-2)
	goto L488
L785:
	;
	goto L786
L786:
	;
	v3243 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v3244 = v3243 + v3238
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3244
	if base.B2i32(v3238 < int64(0))^base.B2i32(v3244 < v3243) != 0 {
		goto L787
	} else {
		goto L788
	}
L787:
	;
	v3817 = int32(-2)
	goto L488
L788:
	;
	goto L789
L789:
	;
	v3251 = int32(0)
	v3252 = int32(1)
	if base.F64_eq(v2391, float64(0)) != 0 {
		v3291 = v3251
		v3292 = v3252
		v3293 = v2374
		goto L497
	} else {
		goto L790
	}
L790:
	;
	v3256 = base.F64_mul(v2391, float64(1e+06))
	if base.F64_lt(base.F64_abs(v3256), float64(9.223372036854776e+18)) != 0 {
		goto L792
	} else {
		goto L793
	}
L791:
	;
	v3264 = base.F64_sub(v3256, base.F64_convert_i64_s(v3262))
	if base.F64_gt(v3264, float64(0.5)) != 0 {
		goto L796
	} else {
		goto L797
	}
L792:
	;
	v3260 = base.I64_trunc_f64_s(v3256)
	v3262 = v3260
	goto L791
L793:
	;
	goto L794
L794:
	;
	v3262 = int64(-9223372036854775807 - 1)
	goto L791
L795:
	;
	v3276 = v3275 + v3244
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3276
	if base.B2i32(v3275 < int64(0))^base.B2i32(v3276 < v3244) != 0 {
		v3817 = int32(-2)
		goto L488
	} else {
		goto L800
	}
L796:
	;
	v3275 = v3262 + int64(1)
	goto L795
L797:
	;
	goto L798
L798:
	;
	if base.F64_lt(v3264, float64(-0.5)) == int32(0) {
		v3275 = v3262
		goto L795
	} else {
		goto L799
	}
L799:
	;
	v3275 = v3262 - int64(1)
	goto L795
L800:
	;
	v3291 = v3251
	v3292 = v3252
	v3293 = v2374
	goto L497
L801:
	;
	goto L496
L802:
	;
	v3334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2311))))
	v3337 = v2311 + base.B2i32(v3334 == int32(45))
	v3338 = int32(553196)
	v3342 = m.G0
	v3344 = v3342 - int32(32)
	v3345 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3344)+24)) = v3345
	*(*int64)(unsafe.Add(mBase, uint32(v3344)+16)) = v3345
	*(*int64)(unsafe.Add(mBase, uint32(v3344)+8)) = v3345
	*(*int64)(unsafe.Add(mBase, uint32(v3344))) = v3345
	v3353 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v3353 == int32(0) {
		goto L804
	} else {
		goto L805
	}
L803:
	;
	if (v2310|base.B2i32(v3421 != int32(6)))&int32(1) != 0 {
		goto L490
	} else {
		goto L824
	}
L804:
	;
	v3421 = int32(0)
	goto L803
L805:
	;
	goto L806
L806:
	;
	v3357 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v3357 == int32(0) {
		goto L807
	} else {
		goto L808
	}
L807:
	;
	v3361 = v3337
	goto L810
L808:
	;
	goto L809
L809:
	;
	v3371 = v3338
	v3372 = v3353
	goto L813
L810:
	;
	v3367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3361))))
	if v3367 == v3353 {
		v3361 = v3361 + int32(1)
		goto L810
	} else {
		goto L812
	}
L811:
	;
	v3421 = v3361 - v3337
	goto L803
L812:
	;
	goto L811
L813:
	;
	v3379 = v3344 + int32(base.Ui32(v3372)>>(uint(int32(3))%32))&int32(28)
	v3380 = *(*int32)(unsafe.Add(mBase, uint32(v3379)))
	v3381 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3379))) = v3380 | v3381<<(uint(v3372)%32)
	v3385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3371)+1)))
	if v3385 != 0 {
		v3371 = v3371 + v3381
		v3372 = v3385
		goto L813
	} else {
		goto L815
	}
L814:
	;
	v3388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3337))))
	if v3388 == int32(0) {
		v3413 = v3337
		goto L816
	} else {
		goto L817
	}
L815:
	;
	goto L814
L816:
	;
	v3421 = v3413 - v3337
	goto L803
L817:
	;
	v3392 = v3337
	v3393 = v3388
	goto L818
L818:
	;
	v3401 = *(*int32)(unsafe.Add(mBase, uint32(v3344+int32(base.Ui32(v3393)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3401)>>(uint(v3393)%32))&int32(1) == int32(0) {
		goto L820
	} else {
		goto L821
	}
L819:
	;
	v3413 = v3409
	goto L816
L820:
	;
	v3413 = v3392
	goto L816
L821:
	;
	goto L822
L822:
	;
	v3407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3392)+1)))
	v3409 = v3392 + int32(1)
	if v3407 != 0 {
		v3392 = v3409
		v3393 = v3407
		goto L818
	} else {
		goto L823
	}
L823:
	;
	goto L819
L824:
	;
	v3428 = v2269 + int32(72)
	v3430 = base.I64_div_s(v2388, int64(10000))
	v3433 = int64(3600000000)
	v3434 = int64(0)
	v3439 = int64(32)
	v3442 = int64(base.Ui64(v3430) >> (uint(v3439) % 64))
	v3445 = int64(4294967295)
	v3448 = v3430 & v3445
	v3449 = v3433 * v3448
	v3453 = int64(base.Ui64(v3449)>>(uint(v3439)%64)) + v3433*v3442
	v3460 = v3448*v3434 + v3453&v3445
	*(*int64)(unsafe.Add(mBase, uint32(v3428)+8)) = v3430*v3434 + v3430>>(uint(int64(63))%64)*v3433 + v3434*v3442 + int64(base.Ui64(v3453)>>(uint(v3439)%64)) + int64(base.Ui64(v3460)>>(uint(v3439)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3428))) = v3449&v3445 | v3460<<(uint(v3439)%64)
	goto L825
L825:
	;
	v3471 = int32(-2)
	v3472 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+80))
	v3473 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+72))
	if v3472 != v3473>>(uint(int64(63))%64) {
		v3834 = v3471
		goto L485
	} else {
		goto L826
	}
L826:
	;
	v3477 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v3478 = v3477 + v3473
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3478
	if base.B2i32(v3473 < int64(0))^base.B2i32(v3478 < v3477) != 0 {
		v3834 = v3471
		goto L485
	} else {
		goto L827
	}
L827:
	;
	v3484 = int64(100)
	v3485 = base.I64_div_s(v2388, v3484)
	v3487 = base.I64_rem_s(v3485, v3484)
	v3489 = v3487 * int64(60000000)
	v3490 = v3478 + v3489
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3490
	if base.B2i32(v3489 < int64(0))^base.B2i32(v3490 < v3478) != 0 {
		v3834 = v3471
		goto L485
	} else {
		goto L828
	}
L828:
	;
	v3500 = (v2388 - v3485*int64(100)) * int64(1000000)
	v3501 = v3490 + v3500
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3501
	if base.B2i32(v3500 < int64(0))^base.B2i32(v3501 < v3490) != 0 {
		v3834 = v3471
		goto L485
	} else {
		goto L829
	}
L829:
	;
	if base.F64_ne(v2391, float64(0)) != 0 {
		goto L830
	} else {
		goto L831
	}
L830:
	;
	if base.F64_lt(base.F64_abs(v2391), float64(9.223372036854776e+18)) != 0 {
		goto L835
	} else {
		goto L836
	}
L831:
	;
	goto L832
L832:
	;
	v3834 = int32(0)
	goto L485
L833:
	;
	v3528 = v3501 + v3527
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3528
	if base.B2i32(v3527 < int64(0))^base.B2i32(v3528 < v3501) != 0 {
		v3834 = v3471
		goto L485
	} else {
		goto L842
	}
L834:
	;
	v3516 = base.F64_sub(v2391, base.F64_convert_i64_s(v3514))
	if base.F64_gt(v3516, float64(0.5)) != 0 {
		goto L838
	} else {
		goto L839
	}
L835:
	;
	v3512 = base.I64_trunc_f64_s(v2391)
	v3514 = v3512
	goto L834
L836:
	;
	goto L837
L837:
	;
	v3514 = int64(-9223372036854775807 - 1)
	goto L834
L838:
	;
	v3527 = v3514 + int64(1)
	goto L833
L839:
	;
	goto L840
L840:
	;
	if base.F64_lt(v3516, float64(-0.5)) == int32(0) {
		v3527 = v3514
		goto L833
	} else {
		goto L841
	}
L841:
	;
	v3527 = v3514 - int64(1)
	goto L833
L842:
	;
	goto L832
L843:
	;
	v3541 = v2269 + int32(56)
	v3544 = int64(3600000000)
	v3545 = int64(0)
	v3550 = int64(32)
	v3553 = int64(base.Ui64(v2388) >> (uint(v3550) % 64))
	v3556 = int64(4294967295)
	v3559 = v2388 & v3556
	v3560 = v3544 * v3559
	v3564 = int64(base.Ui64(v3560)>>(uint(v3550)%64)) + v3544*v3553
	v3571 = v3559*v3545 + v3564&v3556
	*(*int64)(unsafe.Add(mBase, uint32(v3541)+8)) = v2388*v3545 + v2388>>(uint(int64(63))%64)*v3544 + v3545*v3553 + int64(base.Ui64(v3564)>>(uint(v3550)%64)) + int64(base.Ui64(v3571)>>(uint(v3550)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3541))) = v3560&v3556 | v3571<<(uint(v3550)%64)
	goto L844
L844:
	;
	v3582 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+64))
	v3583 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+56))
	if v3582 != v3583>>(uint(int64(63))%64) {
		goto L845
	} else {
		goto L846
	}
L845:
	;
	v3834 = int32(-2)
	goto L485
L846:
	;
	goto L847
L847:
	;
	v3588 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v3589 = v3588 + v3583
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3589
	if base.B2i32(v3583 < int64(0))^base.B2i32(v3589 < v3588) != 0 {
		goto L848
	} else {
		goto L849
	}
L848:
	;
	v3834 = int32(-2)
	goto L485
L849:
	;
	goto L850
L850:
	;
	if base.F64_eq(v2391, float64(0)) != 0 {
		goto L851
	} else {
		goto L852
	}
L851:
	;
	if v2393 == int32(0) {
		goto L863
	} else {
		goto L864
	}
L852:
	;
	v3599 = base.F64_mul(v2391, float64(3.6e+09))
	if base.F64_lt(base.F64_abs(v3599), float64(9.223372036854776e+18)) != 0 {
		goto L854
	} else {
		goto L855
	}
L853:
	;
	v3607 = base.F64_sub(v3599, base.F64_convert_i64_s(v3605))
	if base.F64_gt(v3607, float64(0.5)) != 0 {
		goto L858
	} else {
		goto L859
	}
L854:
	;
	v3603 = base.I64_trunc_f64_s(v3599)
	v3605 = v3603
	goto L853
L855:
	;
	goto L856
L856:
	;
	v3605 = int64(-9223372036854775807 - 1)
	goto L853
L857:
	;
	v3619 = v3618 + v3589
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3619
	if base.B2i32(v3618 < int64(0))^base.B2i32(v3619 < v3589) == int32(0) {
		goto L851
	} else {
		goto L862
	}
L858:
	;
	v3618 = v3605 + int64(1)
	goto L857
L859:
	;
	goto L860
L860:
	;
	if base.F64_lt(v3607, float64(-0.5)) == int32(0) {
		v3618 = v3605
		goto L857
	} else {
		goto L861
	}
L861:
	;
	v3618 = v3605 - int64(1)
	goto L857
L862:
	;
	v3834 = int32(-2)
	goto L485
L863:
	;
	v3834 = int32(0)
	goto L485
L864:
	;
	goto L865
L865:
	;
	v3640 = F_ParseISO8601Number(m, v2374, v2269+int32(108), v2269+int32(96), v2269+int32(88))
	mBase = m.M
	v3641 = m.ExcPending
	if v3641 != 0 {
		goto L30
	} else {
		goto L866
	}
L866:
	;
	if v3640 != 0 {
		v3834 = v3640
		goto L485
	} else {
		goto L867
	}
L867:
	;
	v3642 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+96))
	v3643 = *(*float64)(unsafe.Add(mBase, uint32(v2269)+88))
	v3644 = int64(60000000)
	v3645 = int32(0)
	v3649 = m.G0
	v3651 = v3649 - int32(16)
	m.G0 = v3651
	v3653 = int64(63)
	F___multi3(m, v3651, v3642, v3642>>(uint(v3653)%64), v3644, int64(0))
	mBase = m.M
	v3658 = *(*int64)(unsafe.Add(mBase, uint32(v3651)+8))
	v3659 = *(*int64)(unsafe.Add(mBase, uint32(v3651)))
	if v3658 != v3659>>(uint(v3653)%64) {
		v3705 = v3645
		goto L869
	} else {
		goto L870
	}
L868:
	;
	if v3705 == int32(0) {
		goto L882
	} else {
		goto L883
	}
L869:
	;
	m.G0 = v3651 + int32(16)
	goto L868
L870:
	;
	v3663 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v3664 = v3663 + v3659
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3664
	if base.B2i32(v3659 < int64(0))^base.B2i32(v3664 < v3663) != 0 {
		v3705 = v3645
		goto L869
	} else {
		goto L871
	}
L871:
	;
	if base.F64_eq(v3643, float64(0)) != 0 {
		v3705 = int32(1)
		goto L869
	} else {
		goto L872
	}
L872:
	;
	v3674 = base.F64_mul(v3643, base.F64_convert_i64_u(v3644))
	if base.F64_lt(base.F64_abs(v3674), float64(9.223372036854776e+18)) != 0 {
		goto L874
	} else {
		goto L875
	}
L873:
	;
	v3682 = base.F64_sub(v3674, base.F64_convert_i64_s(v3680))
	if base.F64_gt(v3682, float64(0.5)) != 0 {
		goto L878
	} else {
		goto L879
	}
L874:
	;
	v3678 = base.I64_trunc_f64_s(v3674)
	v3680 = v3678
	goto L873
L875:
	;
	goto L876
L876:
	;
	v3680 = int64(-9223372036854775807 - 1)
	goto L873
L877:
	;
	v3694 = v3693 + v3664
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3694
	v3705 = base.B2i32(base.B2i32(v3693 < int64(0))^base.B2i32(v3694 < v3664) == int32(0))
	goto L869
L878:
	;
	v3693 = v3680 + int64(1)
	goto L877
L879:
	;
	goto L880
L880:
	;
	if base.F64_lt(v3682, float64(-0.5)) == int32(0) {
		v3693 = v3680
		goto L877
	} else {
		goto L881
	}
L881:
	;
	v3693 = v3680 - int64(1)
	goto L877
L882:
	;
	v3834 = int32(-2)
	goto L485
L883:
	;
	goto L884
L884:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+108))
	v3715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3714))))
	if v3715 == int32(0) {
		v3834 = v3715
		goto L485
	} else {
		goto L885
	}
L885:
	;
	if v3715 != int32(58) {
		goto L886
	} else {
		goto L887
	}
L886:
	;
	v3834 = int32(-1)
	goto L485
L887:
	;
	goto L888
L888:
	;
	v3722 = v3714 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2269)+108)) = v3722
	v3730 = F_ParseISO8601Number(m, v3722, v2269+int32(108), v2269+int32(96), v2269+int32(88))
	mBase = m.M
	v3731 = m.ExcPending
	if v3731 != 0 {
		goto L30
	} else {
		goto L889
	}
L889:
	;
	if v3730 != 0 {
		v3834 = v3730
		goto L485
	} else {
		goto L890
	}
L890:
	;
	v3733 = *(*int64)(unsafe.Add(mBase, uint32(v2269)+96))
	v3734 = *(*float64)(unsafe.Add(mBase, uint32(v2269)+88))
	v3735 = int64(1000000)
	v3736 = int32(0)
	v3740 = m.G0
	v3742 = v3740 - int32(16)
	m.G0 = v3742
	v3744 = int64(63)
	F___multi3(m, v3742, v3733, v3733>>(uint(v3744)%64), v3735, int64(0))
	mBase = m.M
	v3749 = *(*int64)(unsafe.Add(mBase, uint32(v3742)+8))
	v3750 = *(*int64)(unsafe.Add(mBase, uint32(v3742)))
	if v3749 != v3750>>(uint(v3744)%64) {
		v3796 = v3736
		goto L892
	} else {
		goto L893
	}
L891:
	;
	if v3796 == int32(0) {
		v3834 = int32(-2)
		goto L485
	} else {
		goto L905
	}
L892:
	;
	m.G0 = v3742 + int32(16)
	goto L891
L893:
	;
	v3754 = *(*int64)(unsafe.Add(mBase, uint32(v2276)))
	v3755 = v3754 + v3750
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3755
	if base.B2i32(v3750 < int64(0))^base.B2i32(v3755 < v3754) != 0 {
		v3796 = v3736
		goto L892
	} else {
		goto L894
	}
L894:
	;
	if base.F64_eq(v3734, float64(0)) != 0 {
		v3796 = int32(1)
		goto L892
	} else {
		goto L895
	}
L895:
	;
	v3765 = base.F64_mul(v3734, base.F64_convert_i64_u(v3735))
	if base.F64_lt(base.F64_abs(v3765), float64(9.223372036854776e+18)) != 0 {
		goto L897
	} else {
		goto L898
	}
L896:
	;
	v3773 = base.F64_sub(v3765, base.F64_convert_i64_s(v3771))
	if base.F64_gt(v3773, float64(0.5)) != 0 {
		goto L901
	} else {
		goto L902
	}
L897:
	;
	v3769 = base.I64_trunc_f64_s(v3765)
	v3771 = v3769
	goto L896
L898:
	;
	goto L899
L899:
	;
	v3771 = int64(-9223372036854775807 - 1)
	goto L896
L900:
	;
	v3785 = v3784 + v3755
	*(*int64)(unsafe.Add(mBase, uint32(v2276))) = v3785
	v3796 = base.B2i32(base.B2i32(v3784 < int64(0))^base.B2i32(v3785 < v3755) == int32(0))
	goto L892
L901:
	;
	v3784 = v3771 + int64(1)
	goto L900
L902:
	;
	goto L903
L903:
	;
	if base.F64_lt(v3773, float64(-0.5)) == int32(0) {
		v3784 = v3771
		goto L900
	} else {
		goto L904
	}
L904:
	;
	v3784 = v3771 - int64(1)
	goto L900
L905:
	;
	v3806 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+108))
	v3807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3806))))
	if v3807 != 0 {
		goto L906
	} else {
		goto L907
	}
L906:
	;
	v3808 = int32(-1)
	goto L908
L907:
	;
	v3808 = int32(0)
	goto L908
L908:
	;
	v3834 = v3808
	goto L485
L909:
	;
	m.G0 = v40 + int32(528)
	return v3975
L910:
	;
	if v3875 == int32(-2) {
		goto L913
	} else {
		goto L914
	}
L911:
	;
	goto L912
L912:
	;
	v3912 = F_palloc(m, int32(16))
	mBase = m.M
	v3913 = m.ExcPending
	if v3913 != 0 {
		goto L30
	} else {
		goto L917
	}
L913:
	;
	v3902 = int32(-4)
	goto L915
L914:
	;
	v3902 = v3875
	goto L915
L915:
	;
	F_DateTimeParseError(m, v3902, v40+int32(8), v44, int32(310296), v43)
	mBase = m.M
	v3907 = m.ExcPending
	if v3907 != 0 {
		goto L30
	} else {
		goto L916
	}
L916:
	;
	v3908 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v3908)
	v3975 = int32(0)
	goto L909
L917:
	;
	v3914 = *(*int32)(unsafe.Add(mBase, uint32(v40)+500))
	switch v3914 - int32(9) {
	case 0:
		goto L921
	case 1:
		goto L919
	default:
		goto L920
	case 8:
		goto L922
	}
L918:
	;
	F_AdjustIntervalForTypmod(m, v3912, v42, v43)
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		goto L30
	} else {
		goto L934
	}
L919:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3912)+8)) = int64(9223372034707292159)
	*(*int64)(unsafe.Add(mBase, uint32(v3912))) = int64(9223372036854775807)
	goto L918
L920:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3955 = m.ExcPending
	if v3955 != 0 {
		goto L30
	} else {
		goto L931
	}
L921:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3912)+8)) = int64(-9223372034707292160)
	*(*int64)(unsafe.Add(mBase, uint32(v3912))) = int64(-9223372036854775807 - 1)
	goto L918
L922:
	;
	v3917 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+516)))
	v3918 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+520)))
	v3921 = v3917 + v3918*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v3921-int64(2147483648)) {
		goto L923
	} else {
		goto L924
	}
L923:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3912)+12)) = uint32(v3921)
	v3927 = *(*int32)(unsafe.Add(mBase, uint32(v40)+512))
	*(*int32)(unsafe.Add(mBase, uint32(v3912)+8)) = v3927
	v3929 = *(*int64)(unsafe.Add(mBase, uint32(v40)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v3912))) = v3929
	goto L918
L924:
	;
	goto L925
L925:
	;
	v3931 = int32(0)
	v3932 = F_errsave_start(m, v43)
	mBase = m.M
	v3933 = m.ExcPending
	if v3933 != 0 {
		goto L30
	} else {
		goto L926
	}
L926:
	;
	if v3932 == int32(0) {
		v3975 = v3931
		goto L909
	} else {
		goto L927
	}
L927:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v3938 = m.ExcPending
	if v3938 != 0 {
		goto L30
	} else {
		goto L928
	}
L928:
	;
	F_errmsg(m, int32(403451), int32(0))
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		goto L30
	} else {
		goto L929
	}
L929:
	;
	F_errsave_finish(m, v43, int32(497032), int32(948), int32(280217))
	mBase = m.M
	v3947 = m.ExcPending
	if v3947 != 0 {
		goto L30
	} else {
		goto L930
	}
L930:
	;
	v3975 = v3931
	goto L909
L931:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v44
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v40)+500))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v3957
	F_errmsg_internal(m, int32(713858), v40)
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L30
	} else {
		goto L932
	}
L932:
	;
	F_errfinish(m, int32(497032), int32(961), int32(280217))
	mBase = m.M
	v3966 = m.ExcPending
	if v3966 != 0 {
		goto L30
	} else {
		goto L933
	}
L933:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L934:
	;
	v3975 = v3912
	goto L909
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
	v16 = F_DirectFunctionCall2Coll(m, int32(1476), v4, l1, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v22 = F_DirectFunctionCall2Coll(m, int32(1475), v4, v16, v7+int32(8))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_DirectFunctionCall2Coll(m, int32(1474), v4, v22, l0)
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v110 + v112
	v120 = v113 + v115
	v124 = int64(63)
	return base.B2i32(v119^v120|(base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v112)))+(v111+v110>>(uint(v124)%64))^(base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v115)))+(v114+v113>>(uint(v124)%64)))) != int64(0))
}
func F_interval_part_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v224 int64
	_ = v224
	var v228 float64
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v321 int64
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 float64
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v436 int64
	_ = v436
	var v441 int64
	_ = v441
	var v444 int64
	_ = v444
	var v445 int64
	_ = v445
	var v450 int64
	_ = v450
	var v453 int64
	_ = v453
	var v456 int64
	_ = v456
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v464 int64
	_ = v464
	var v471 int64
	_ = v471
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v484 int64
	_ = v484
	var v490 int64
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v520 int64
	_ = v520
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v556 int64
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = int32(1)
	v27 = v22 + v26
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v32 = v30 & v26
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = v27
	goto L5
L4:
	;
	v33 = v22 + int32(4)
	goto L5
L5:
	;
	if v30 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28))))
	v64 = F_downcase_truncate_identifier(m, v33, v61, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v36 = int32(4)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v38&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v51 = int32(1)
	if v32 != 0 {
		v61 = int32(base.Ui32(v30)>>(uint(v51)%32)) - v51
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v47 = v36
	goto L12
L11:
	;
	v47 = base.B2i32(v38 == int32(18)) << (uint(v36) % 32)
	goto L12
L12:
	;
	if v38 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = v36
	goto L15
L14:
	;
	v50 = v47
	goto L15
L15:
	;
	v61 = v50
	goto L6
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v67 = v17 + int32(76)
	v74 = *(*int32)(unsafe.Add(mBase, _consts[1115]))
	if v74 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v135 == int32(31) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1115])) = v118
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118)+11)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v126
	v135 = v125
	goto L18
L20:
	;
	v76 = F_strncmp(m, v64, v74, int32(10))
	mBase = m.M
	if v76 == int32(0) {
		v118 = v74
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64))))
	v85 = int32(1663648)
	v87 = int32(1664608)
	goto L24
L23:
	;
	goto L22
L24:
	;
	v94 = v85 + (v87-v85)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v95 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94))))
	v96 = v79 - v95
	if v96 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(0)
	v135 = int32(31)
	goto L18
L26:
	;
	v100 = F_strncmp(m, v64, v94, int32(10))
	mBase = m.M
	if v100 == int32(0) {
		v118 = v94
		goto L19
	} else {
		goto L29
	}
L27:
	;
	v103 = v96
	goto L28
L28:
	;
	v107 = base.B2i32(v103 < int32(0))
	if v103 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v103 = v100
	goto L28
L30:
	;
	v108 = v94 - int32(16)
	goto L32
L31:
	;
	v108 = v87
	goto L32
L32:
	;
	if v103 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v111 = v85
	goto L35
L34:
	;
	v111 = v94 + int32(16)
	goto L35
L35:
	;
	if base.Ui32(v111) <= base.Ui32(v108) {
		v85 = v111
		v87 = v108
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v139 = v17 + int32(76)
	v146 = *(*int32)(unsafe.Add(mBase, _consts[1042]))
	if v146 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v208 = v135
	goto L39
L39:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if v209 != int32(2147483647) {
		goto L62
	} else {
		goto L63
	}
L40:
	;
	v208 = v207
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1042])) = v190
	v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v190)+11)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v198
	v207 = v197
	goto L40
L42:
	;
	v148 = F_strncmp(m, v64, v146, int32(10))
	mBase = m.M
	if v148 == int32(0) {
		v190 = v146
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v151 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64))))
	v157 = int32(1662496)
	v159 = int32(1663632)
	goto L46
L45:
	;
	goto L44
L46:
	;
	v166 = v157 + (v159-v157)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v167 = int32(*(*int8)(unsafe.Add(mBase, uint32(v166))))
	v168 = v151 - v167
	if v168 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(0)
	v207 = int32(31)
	goto L40
L48:
	;
	v172 = F_strncmp(m, v64, v166, int32(10))
	mBase = m.M
	if v172 == int32(0) {
		v190 = v166
		goto L41
	} else {
		goto L51
	}
L49:
	;
	v175 = v168
	goto L50
L50:
	;
	v179 = base.B2i32(v175 < int32(0))
	if v175 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v175 = v172
	goto L50
L52:
	;
	v180 = v166 - int32(16)
	goto L54
L53:
	;
	v180 = v159
	goto L54
L54:
	;
	if v175 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v183 = v157
	goto L57
L56:
	;
	v183 = v166 + int32(16)
	goto L57
L57:
	;
	if base.Ui32(v183) <= base.Ui32(v180) {
		v157 = v183
		v159 = v180
		goto L46
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	m.G0 = v17 + int32(80)
	return v571
L60:
	;
	if v208 == int32(17) {
		goto L100
	} else {
		goto L101
	}
L61:
	;
	switch v208 {
	case 0, 17:
		goto L70
	default:
		goto L71
	}
L62:
	;
	if v209 != int32(-2147483648) {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v221 != int32(2147483647) {
		goto L60
	} else {
		goto L68
	}
L65:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v214 != int32(-2147483648) {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	if v217 != int64(-9223372036854775807-1) {
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v228 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L61
L68:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	if v224 != int64(9223372036854775807) {
		goto L60
	} else {
		goto L69
	}
L69:
	;
	v228 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L61
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if base.Ui32(v251) <= base.Ui32(int32(30)) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v237 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v64
	F_errmsg(m, int32(190460), v17+int32(48))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(497032), int32(6058), int32(83481))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v301 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v301)
	v571 = int32(0)
	goto L59
L78:
	;
	if l1 != 0 {
		goto L89
	} else {
		goto L90
	}
L79:
	;
	v255 = int32(1) << (uint(v251) % 32)
	if v255&int32(506464256) != 0 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	if v255&int32(1640759296) != 0 {
		goto L77
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v269 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v64
	F_errmsg(m, int32(190497), v17-int32(-64))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(497032), int32(6089), int32(83481))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	if base.F64_lt(v228, float64(0)) != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v299 = F_Float8GetDatum(m, v228)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L97
	}
L92:
	;
	v286 = int32(0)
	v290 = F_DirectFunctionCall3Coll(m, int32(408), v286, int32(11446), v286, int32(-1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v293 = int32(0)
	v297 = F_DirectFunctionCall3Coll(m, int32(408), v293, int32(11457), v293, int32(-1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v571 = v290
	goto L59
L96:
	;
	v571 = v297
	goto L59
L97:
	;
	v571 = v299
	goto L59
L98:
	;
	if l1 != 0 {
		goto L153
	} else {
		goto L154
	}
L99:
	;
	v556 = base.I64_extend32_s(v321) + base.I64_extend32_s(v318)*int64(1000000)
	goto L98
L100:
	;
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v308 = base.I64_div_s(v306, int64(3600000000))
	v311 = v308*int64(-3600000000) + v306
	v313 = base.I64_div_s(v311, int64(60000000))
	v316 = v313*int64(-60000000) + v311
	v318 = base.I64_div_s(v316, int64(1000000))
	v321 = v318*int64(4293967296) + v316
	v322 = base.I32_wrap_i64(v321)
	v323 = base.I32_wrap_i64(v318)
	v325 = base.I32_rem_s(v209, int32(12))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	switch v327 - int32(18) {
	case 0:
		goto L113
	case 1:
		goto L112
	case 2:
		v556 = v308
		goto L98
	case 3:
		goto L111
	case 4:
		goto L110
	case 5:
		goto L109
	case 6:
		goto L108
	case 7:
		goto L107
	case 8:
		goto L106
	case 9:
		goto L105
	case 10:
		goto L104
	case 11:
		goto L114
	case 12:
		goto L99
	default:
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v208 != 0 {
		goto L133
	} else {
		goto L134
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L128
	}
L104:
	;
	v398 = base.I32_div_s(v209, int32(12000))
	v556 = base.I64_extend_i32_s(v398)
	goto L98
L105:
	;
	v395 = base.I32_div_s(v209, int32(1200))
	v556 = base.I64_extend_i32_s(v395)
	goto L98
L106:
	;
	v392 = base.I32_div_s(v209, int32(120))
	v556 = base.I64_extend_i32_s(v392)
	goto L98
L107:
	;
	v389 = base.I32_div_s(v209, int32(12))
	v556 = base.I64_extend_i32_s(v389)
	goto L98
L108:
	;
	if int32(0) <= v209 {
		goto L125
	} else {
		goto L126
	}
L109:
	;
	v556 = base.I64_extend_i32_s(v325)
	goto L98
L110:
	;
	v365 = base.I32_div_s(v326, int32(7))
	v556 = base.I64_extend_i32_s(v365)
	goto L98
L111:
	;
	v556 = base.I64_extend_i32_s(v326)
	goto L98
L112:
	;
	v556 = base.I64_extend32_s(v313)
	goto L98
L113:
	;
	if l1 != 0 {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	if l1 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v336 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v321)+base.I64_extend32_s(v318)*int64(1000000), int32(3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v339 = float64(1000)
	v345 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v323), v339), base.F64_div(base.F64_convert_i32_s(v322), v339)))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	v571 = v336
	goto L59
L119:
	;
	v571 = v345
	goto L59
L120:
	;
	v353 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v321)+base.I64_extend32_s(v318)*int64(1000000), int32(6))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v360 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v322), float64(1e+06)), base.F64_convert_i32_s(v323)))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	v571 = v353
	goto L59
L124:
	;
	v571 = v360
	goto L59
L125:
	;
	v373 = base.I32_div_u_s(v325&int32(255), int32(3))
	v556 = base.I64_extend_i32_u(v373 + int32(1))
	goto L98
L126:
	;
	goto L127
L127:
	;
	v380 = base.I32_rem_s(int32(0)-v209, int32(12))
	v383 = base.I32_div_s(base.I32_extend8_s(v380), int32(-3))
	v556 = base.I64_extend8_s(base.I64_extend_i32_u(v383 - int32(1)))
	goto L98
L128:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v408 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v64
	F_errmsg(m, int32(190497), v17)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(497032), int32(6233), int32(246191))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L148
	}
L134:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v420 != int32(11) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	if l1 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v424 = v17 + int32(16)
	v425 = int32(12)
	v426 = base.I32_rem_s(v209, v425)
	v431 = base.I32_div_s(v209, v425)
	v436 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+8)))
	v441 = (base.I64_extend_i32_s(v426*int32(120)) + base.I64_extend_i32_s(v431)*int64(1461) + v436<<(uint(int64(2))%64)) * int64(21600)
	v444 = int64(1000000)
	v445 = int64(0)
	v450 = int64(32)
	v453 = int64(base.Ui64(v441) >> (uint(v450) % 64))
	v456 = int64(4294967295)
	v459 = v441 & v456
	v460 = v444 * v459
	v464 = int64(base.Ui64(v460)>>(uint(v450)%64)) + v444*v453
	v471 = v459*v445 + v464&v456
	*(*int64)(unsafe.Add(mBase, uint32(v424)+8)) = v441*v445 + v441>>(uint(int64(63))%64)*v444 + v445*v453 + int64(base.Ui64(v464)>>(uint(v450)%64)) + int64(base.Ui64(v471)>>(uint(v450)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v424))) = v460&v456 | v471<<(uint(v450)%64)
	goto L139
L137:
	;
	goto L138
L138:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v509 = int32(12)
	v510 = base.I32_div_s(v209, v509)
	v520 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v527 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v505), float64(86400)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v209-v510*v509), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v510), float64(3.15576e+07)), base.F64_div(base.F64_convert_i64_s(v520), float64(1e+06))))))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L147
	}
L139:
	;
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	if v483 != v484>>(uint(int64(63))%64) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v498 = F_int64_div_fast_to_numeric(m, v482, int32(6))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L144
	}
L141:
	;
	v490 = v482 + v484
	if base.B2i32(v482 < int64(0))^base.B2i32(v490 < v484) != 0 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v494 = F_int64_div_fast_to_numeric(m, v490, int32(6))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v571 = v494
	goto L59
L144:
	;
	v500 = F_int64_to_numeric(m, v441)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v503 = F_numeric_add_opt_error(m, v498, v500, int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v571 = v503
	goto L59
L147:
	;
	v571 = v527
	goto L59
L148:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v537 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v64
	F_errmsg(m, int32(190460), v17+int32(32))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(497032), int32(6294), int32(246191))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	v557 = F_int64_to_numeric(m, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v560 = F_Float8GetDatum(m, base.F64_convert_i64_s(v556))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v571 = v557
	goto L59
L157:
	;
	v571 = v560
	goto L59
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
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = v51
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v53
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
							F_errmsg(m, int32(403451), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(497032), int32(4285), int32(287329))
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
	var v49 int32
	_ = v49
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != int32(457) {
		v49 = v2
		return v49
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 != int32(7) {
			v49 = v2
			return v49
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
			if v17 != 0 {
				v49 = v2
				return v49
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
						v49 = v44
						return v49
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
									v49 = v2
									return v49
								} else {
									if v26 != 0 {
										v44 = F_relabel_to_typmod(m, v18, v19)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											v49 = v44
											return v49
										}
									} else {
										v32 = v19 & int32(65535)
										if base.Ui32(int32(5)) < base.Ui32(v32) {
											v44 = F_relabel_to_typmod(m, v18, v19)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												v49 = v44
												return v49
											}
										} else {
											if v22 < int32(0) {
												v38 = int32(-1)
											} else {
												v38 = v22
											}
											if base.Ui32(v32) < base.Ui32(v38&int32(65535)) {
												v49 = v2
												return v49
											} else {
												v44 = F_relabel_to_typmod(m, v18, v19)
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return int32(0)
												} else {
													v49 = v44
													return v49
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
