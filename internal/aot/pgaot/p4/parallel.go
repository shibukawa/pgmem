package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ExecParallelHashJoin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
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
	var v370 int32
	_ = v370
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int64
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v510 int32
	_ = v510
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v832 int64
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int64
	_ = v921
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v957 int64
	_ = v957
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 float64
	_ = v970
	var v974 int32
	_ = v974
	var v977 float64
	_ = v977
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int64
	_ = v996
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 float64
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int64
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1202 int64
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1216 float64
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1260 int64
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1283 int64
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 float64
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int64
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1408 int32
	_ = v1408
	var v1409 int64
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1421 float64
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1556 int32
	_ = v1556
	var v1564 int32
	_ = v1564
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1606 int32
	_ = v1606
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1829 int32
	_ = v1829
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1959 int32
	_ = v1959
	var v1964 int32
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1986 int32
	_ = v1986
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2015 int32
	_ = v2015
	var v2020 int32
	_ = v2020
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(48)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+136))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v44 = v34 + int32(56)
	v55 = v29
	goto L4
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L407
	}
L4:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
	if v70 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v27 + int32(48)
	return v1986
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	switch v73 - int32(1) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	case 5:
		goto L13
	case 6:
		goto L12
	case 7:
		goto L11
	default:
		goto L3
	}
L9:
	;
	goto L8
L10:
	;
	goto L5
L11:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+144))
	if v1465 == int32(0) {
		v1986 = v2
		goto L10
	} else {
		goto L336
	}
L12:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1348 = F_tuplestore_gettupleslot(m, v1344, int32(1), int32(0), v1347)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L1
	} else {
		goto L312
	}
L13:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v1222 = F_tuplestore_gettupleslot_force(m, v1220, v1221)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L284
	}
L14:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v1016 = v1013
	goto L243
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(2)
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+173)))
	if v983 != 0 {
		goto L4
	} else {
		goto L232
	}
L16:
	;
	v764 = m.G0
	v766 = v764 - int32(16)
	m.G0 = v766
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v771 != 0 {
		goto L177
	} else {
		goto L178
	}
L17:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+44))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+48))
	if v353|base.B2i32(v352 != int32(1)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	v78 = F_ExecHashTableCreate(m, v33)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v78
	v82 = F_MultiExecProcNode(m, v33)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v78)+64))
	if base.F64_ne(v84, float64(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+56)) = v122
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+174)) = uint8(v124)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v126 == int32(3) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	if v87 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v88 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if int32(3) < v89 {
		v1986 = v2
		goto L10
	} else {
		goto L25
	}
L25:
	;
	goto L26
L26:
	;
	v117 = F_BarrierArriveAndWait(m, v44, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v1986 = v2
	goto L10
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v119 < int32(4) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	if v129 < int32(2) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+48)) = int32(-1)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v343 != 0 {
		goto L73
	} else {
		goto L74
	}
L33:
	;
	v315 = F_BarrierArriveAndWait(m, v44, int32(134217748))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L72
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L35
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v134)+52))
	if v159 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v251 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
	if v252 <= v251 {
		goto L33
	} else {
		goto L67
	}
L37:
	;
	F_ExecReScan(m, v134)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v163 = m.T0[v162].(func(*base.Module, int32) int32)(m, v134)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	goto L36
L42:
	;
	if v163 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
	if v167&int32(2) != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v163
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v133)+20))
	F_MemoryContextReset(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v174 = int32(_a_F_ExecParallelHashJoin_0)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v133)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v178
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+24))
	v183 = m.T0[v182].(func(*base.Module, int32, int32, int32) int64)(m, v176, v133, v27+int32(43))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v175
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+44)) = uint32(v183)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+43)))
	if v188 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
	if v246 == int32(0) {
		goto L35
	} else {
		goto L65
	}
L48:
	;
	v193 = F_ExecFetchSlotMinimalTuple(m, v163, v27+int32(31))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v232 != int32(1) {
		goto L47
	} else {
		goto L59
	}
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(32)))) = (v202 - int32(1)) & v195
	if base.Ui32(int32(2)) <= base.Ui32(v201) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v132)+144))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217+v218*int32(36))+32))
	F_sts_puttuple(m, v222, v27+int32(44), v193)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L56
	}
L53:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v215 = (v201 - int32(1)) & base.I32_rotr(v195, v211)
	goto L55
L54:
	;
	v215 = int32(0)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(36)))) = v215
	goto L52
L56:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+31)))
	if v227 != int32(1) {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v193)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v235 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v238 = F_ExecHashBuildNullTupleStore(m, v132)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v241 = v235
	goto L62
L62:
	;
	F_tuplestore_puttupleslot(m, v241, v163)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v238
	v241 = v238
	goto L62
L64:
	;
	goto L47
L65:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L35
L67:
	;
	v256 = v251
	goto L68
L68:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v132)+144))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v256*int32(36))+32))
	F_sts_end_write(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L33
L70:
	;
	v287 = v256 + int32(1)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
	if v287 < v288 {
		v256 = v287
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L32
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(6)
	v55 = v78
	goto L4
L74:
	;
	goto L75
L75:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	if v346 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(7)
	v55 = v78
	goto L4
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(8)
	v55 = v78
	goto L4
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v489
	v620 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+173)) = uint8(v620)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v622
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v55)+44))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(132)))) = (v628 - int32(1)) & v622
	if base.Ui32(int32(2)) <= base.Ui32(v627) {
		goto L150
	} else {
		goto L151
	}
L80:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v539 != 0 {
		goto L120
	} else {
		goto L121
	}
L81:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+4)))
	if v510&int32(2) == int32(0) {
		goto L79
	} else {
		goto L119
	}
L82:
	;
	F_ExecForceStoreMinimalTuple(m, v443, v445, int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L117
	}
L83:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v351)+144))
	v478 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v474+v353*int32(36))+25)) = uint8(v478)
	goto L80
L84:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
	if v359 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	if v352 <= v353 {
		goto L83
	} else {
		goto L113
	}
L87:
	;
	F_ExecReScan(m, v30)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v363 = m.T0[v362].(func(*base.Module, int32) int32)(m, v30)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	if v363 == int32(0) {
		goto L83
	} else {
		goto L92
	}
L92:
	;
	v370 = v363
	goto L93
L93:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+4)))
	if v391&int32(2) != 0 {
		goto L83
	} else {
		goto L95
	}
L94:
	;
	goto L83
L95:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+12)) = v370
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v394)+20))
	F_MemoryContextReset(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v399 = int32(_a_F_ExecParallelHashJoin_0)
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v394)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v403
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401)+24))
	v408 = m.T0[v407].(func(*base.Module, int32, int32, int32) int64)(m, v401, v394, v27+int32(44))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v400
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+24)) = uint32(v408)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+44)))
	if v413 == int32(0) {
		v489 = v370
		goto L81
	} else {
		goto L98
	}
L98:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v416 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v419 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
	if v429 != 0 {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v422 = F_ExecHashBuildNullTupleStore(m, v351)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	v425 = v419
	goto L104
L104:
	;
	F_tuplestore_puttupleslot(m, v425, v370)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v422
	v425 = v422
	goto L104
L106:
	;
	goto L101
L107:
	;
	F_ExecReScan(m, v30)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v433 = m.T0[v432].(func(*base.Module, int32) int32)(m, v30)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	if v433 != 0 {
		v370 = v433
		goto L93
	} else {
		goto L112
	}
L112:
	;
	goto L94
L113:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v351)+144))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v436+v353*int32(36))+32))
	v443 = F_sts_parallel_scan_next(m, v440, v27+int32(24))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v443 != 0 {
		goto L82
	} else {
		goto L115
	}
L115:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+8))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	m.T0[v447].(func(*base.Module, int32))(m, v445)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	goto L83
L117:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v483 == int32(0) {
		goto L80
	} else {
		goto L118
	}
L118:
	;
	v489 = v483
	goto L81
L119:
	;
	goto L80
L120:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+144))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v540)+48))
	v544 = v542 * int32(36)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v541+v544)))
	v549 = F_BarrierArriveAndDetachExceptLast(m, v546+int32(4))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v614 != 0 {
		goto L146
	} else {
		goto L147
	}
L123:
	;
	if v603 != 0 {
		goto L137
	} else {
		goto L138
	}
L124:
	;
	if v549 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v540)+144))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v540)+48))
	v558 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v553+v554*int32(36))+26)) = uint8(v558)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v540)+144))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v560+v544)+28))
	F_sts_end_parallel_scan(m, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+61)))
	if v582 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L128:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v540)+144))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v565+v544)+32))
	F_sts_end_parallel_scan(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v546)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v540)+48)) = int32(-1)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v540)+104))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v577 = v570 + v574<<(uint(int32(2))%32)
	if base.Ui32(v577) < base.Ui32(v573) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v579 = v573
	goto L132
L131:
	;
	v579 = v577
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v540)+104)) = v579
	v603 = int32(0)
	goto L123
L133:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v540)+144))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v540)+48))
	v590 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v585+v586*int32(36))+26)) = uint8(v590)
	F_ExecHashTableDetachBatch(m, v540)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v603 = int32(1)
	goto L123
L136:
	;
	v603 = int32(0)
	goto L123
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(5)
	goto L4
L138:
	;
	goto L139
L139:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v606 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(6)
	goto L4
L141:
	;
	goto L142
L142:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	if v609 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(7)
	goto L4
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(8)
	goto L4
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(6)
	goto L4
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(8)
	goto L4
L149:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+24)))
	if v647 != int32(1) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v641 = (v627 - int32(1)) & base.I32_rotr(v622, v637)
	goto L152
L151:
	;
	v641 = int32(0)
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(20)))) = v641
	goto L149
L153:
	;
	v687 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v687
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v686
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v55)+48))
	if base.B2i32(v690 == v691)|base.B2i32(v686 != int32(-1)) == v687 {
		goto L163
	} else {
		goto L164
	}
L154:
	;
	v686 = int32(-1)
	goto L153
L155:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
	v653 = v651 - int32(1)
	v654 = v643 & v653
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v650+v654<<(uint(int32(2))%32))))
	if v658 == int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v661 = v654
	v663 = v658
	goto L157
L157:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	if v643 == v666 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L154
L159:
	;
	v686 = v661
	goto L153
L160:
	;
	goto L161
L161:
	;
	v670 = (v661 + int32(1)) & v653
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v650+v670<<(uint(int32(2))%32))))
	if v674 != 0 {
		v661 = v670
		v663 = v674
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	v700 = F_ExecFetchSlotMinimalTuple(m, v489, v27+int32(36))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(3)
	goto L16
L166:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v55)+92))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v704
	v708 = v702 + v703<<(uint(int32(2))%32)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v708)))
	if v709 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v712 = int32(_a_F_ExecParallelHashJoin_0)
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v55)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v715
	v718 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	v723 = v709
	goto L169
L169:
	;
	F_BufFileWrite(m, v723, v27+int32(44), int32(4))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708))) = v718
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v713
	v723 = v718
	goto L169
L171:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	F_BufFileWrite(m, v723, v700, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+36)))
	if v733 != int32(1) {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	F_pfree(m, v700)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	goto L4
L175:
	;
	m.G0 = v766 + int32(16)
	if v898 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L176:
	;
	if v785 != 0 {
		goto L182
	} else {
		goto L183
	}
L177:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v769)+136))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v771)))
	v774 = F_dsa_get_address(m, v772, v773)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v769)+136))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v769)+20))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v777+v778<<(uint(int32(2))%32))))
	v783 = F_dsa_get_address(m, v776, v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L181
	}
L180:
	;
	v785 = v774
	goto L176
L181:
	;
	v785 = v783
	goto L176
L182:
	;
	v787 = v785
	goto L185
L183:
	;
	goto L184
L184:
	;
	v898 = int32(0)
	goto L175
L185:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v787)+4))
	if v810 != v768 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	goto L184
L187:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v769)+136))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v787)))
	v847 = F_dsa_get_address(m, v845, v846)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L198
	}
L188:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v816 = F_ExecStoreMinimalTuple(m, v787+int32(8), v814, int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v816
	if v770 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v787
	v898 = int32(1)
	goto L175
L191:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v824 = int32(_a_F_ExecParallelHashJoin_0)
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v827
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v770)+24))
	v832 = m.T0[v831].(func(*base.Module, int32, int32, int32) int64)(m, v770, v35, v766+int32(15))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L195
	}
L194:
	;
	goto L190
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v825
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	if v832 == int64(0) {
		goto L187
	} else {
		goto L197
	}
L197:
	;
	goto L190
L198:
	;
	if v847 != 0 {
		v787 = v847
		goto L185
	} else {
		goto L199
	}
L199:
	;
	goto L186
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(4)
	goto L4
L201:
	;
	goto L202
L202:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v906 == int32(6) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v910 = int32(*(*int16)(unsafe.Add(mBase, uint32(v909)+18)))
	if v910 < int32(0) {
		goto L4
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	if v32 != 0 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	goto L205
L207:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v974 == int32(0) {
		goto L4
	} else {
		goto L231
	}
L208:
	;
	v913 = int32(_a_F_ExecParallelHashJoin_0)
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v916
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v921 = m.T0[v920].(func(*base.Module, int32, int32, int32) int64)(m, v32, v35, v27+int32(44))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v928 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+173)) = uint8(v928)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v931 = int32(*(*int16)(unsafe.Add(mBase, uint32(v930)+18)))
	if int32(0) <= v931 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v914
	if v921 == int64(0) {
		goto L207
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	v935 = v931 | int32(_a_F_ExecParallelHashJoin_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v930)+18)) = uint16(v935)
	goto L215
L214:
	;
	goto L215
L215:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v937 == int32(5) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(2)
	goto L4
L217:
	;
	goto L218
L218:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v942 == int32(1) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(2)
	goto L221
L220:
	;
	goto L221
L221:
	;
	if v937 == int32(7) {
		goto L4
	} else {
		goto L222
	}
L222:
	;
	if v31 != 0 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v967 == int32(0) {
		goto L4
	} else {
		goto L230
	}
L224:
	;
	v949 = int32(_a_F_ExecParallelHashJoin_0)
	v950 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v952
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v957 = m.T0[v956].(func(*base.Module, int32, int32, int32) int64)(m, v31, v35, v27+int32(44))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v965 = F_ExecProject(m, v964)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v950
	if v957 == int64(0) {
		goto L223
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1986 = v965
	goto L10
L230:
	;
	v970 = *(*float64)(unsafe.Add(mBase, uint32(v967)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v967)+432)) = base.F64_add(v970, float64(1))
	goto L4
L231:
	;
	v977 = *(*float64)(unsafe.Add(mBase, uint32(v974)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v974)+424)) = base.F64_add(v977, float64(1))
	goto L4
L232:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v984 == int32(0) {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v984
	if v31 != 0 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1006 == int32(0) {
		goto L4
	} else {
		goto L241
	}
L235:
	;
	v988 = int32(_a_F_ExecParallelHashJoin_0)
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v991
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v996 = m.T0[v995].(func(*base.Module, int32, int32, int32) int64)(m, v31, v35, v27+int32(44))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1004 = F_ExecProject(m, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v989
	if v996 == int64(0) {
		goto L234
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v1986 = v1004
	goto L10
L241:
	;
	v1009 = *(*float64)(unsafe.Add(mBase, uint32(v1006)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v1006)+432)) = base.F64_add(v1009, float64(1))
	goto L4
L242:
	;
	if v1158 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L243:
	;
	if v1016 != 0 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	if v1059 != 0 {
		goto L252
	} else {
		goto L253
	}
L246:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+136))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	v1041 = F_dsa_get_address(m, v1039, v1040)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	if v1045 <= v1044 {
		v1158 = int32(0)
		goto L242
	} else {
		goto L250
	}
L249:
	;
	v1059 = v1041
	goto L245
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v1044 + int32(1)
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+136))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+20))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1051+v1044<<(uint(int32(2))%32))))
	v1056 = F_dsa_get_address(m, v1050, v1055)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1059 = v1056
	goto L245
L252:
	;
	v1061 = v1059
	goto L255
L253:
	;
	goto L254
L254:
	;
	v1127 = int32(0)
	v1129 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
	if v1129 == v1127 {
		v1016 = v1127
		goto L243
	} else {
		goto L264
	}
L255:
	;
	v1084 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1061)+18)))
	if int32(0) <= v1084 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	goto L254
L257:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1091 = F_ExecStoreMinimalTuple(m, v1061+int32(8), v1089, int32(0))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+136))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1061)))
	v1101 = F_dsa_get_address(m, v1099, v1100)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v1091
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v1094)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v1061
	v1158 = int32(1)
	goto L242
L262:
	;
	if v1101 != 0 {
		v1061 = v1101
		goto L255
	} else {
		goto L263
	}
L263:
	;
	goto L256
L264:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v1016 = v1127
	goto L243
L266:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v1161 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v1169
	if v31 != 0 {
		goto L276
	} else {
		goto L277
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(6)
	goto L4
L270:
	;
	goto L271
L271:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	if v1164 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(7)
	goto L4
L273:
	;
	goto L274
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(8)
	goto L4
L275:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1213 == int32(0) {
		goto L4
	} else {
		goto L283
	}
L276:
	;
	v1171 = int32(_a_F_ExecParallelHashJoin_0)
	v1172 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1174
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v1179 = m.T0[v1178].(func(*base.Module, int32, int32, int32) int64)(m, v31, v35, v27+int32(44))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+80))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+24))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+8))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+12))
	m.T0[v1190].(func(*base.Module, int32))(m, v1188)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L281
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1172
	if v1179 == int64(0) {
		goto L275
	} else {
		goto L280
	}
L280:
	;
	goto L278
L281:
	;
	v1193 = int32(_a_F_ExecParallelHashJoin_0)
	v1194 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1196
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+32))
	v1202 = m.T0[v1201].(func(*base.Module, int32, int32, int32) int64)(m, v1186+int32(8), v1187, int32(0))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1194
	v1206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1188)+4)))
	v1208 = v1206 & int32(_a_F_ExecParallelHashJoin_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1188)+4)) = uint16(v1208)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+12))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1210)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1188)+6)) = uint16(v1211)
	v1986 = v1188
	goto L10
L283:
	;
	v1216 = *(*float64)(unsafe.Add(mBase, uint32(v1213)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v1213)+432)) = base.F64_add(v1216, float64(1))
	goto L4
L284:
	;
	if v1222 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	goto L288
L286:
	;
	goto L287
L287:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	F_tuplestore_end(m, v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L308
	}
L288:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v1248
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v1250
	if v31 != 0 {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	goto L287
L290:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1294 != 0 {
		goto L298
	} else {
		goto L299
	}
L291:
	;
	v1252 = int32(_a_F_ExecParallelHashJoin_0)
	v1253 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1255
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v1260 = m.T0[v1259].(func(*base.Module, int32, int32, int32) int64)(m, v31, v35, v27+int32(44))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L1
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+80))
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+24))
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+8))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+12))
	m.T0[v1271].(func(*base.Module, int32))(m, v1269)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L296
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1253
	if v1260 == int64(0) {
		goto L290
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	v1274 = int32(_a_F_ExecParallelHashJoin_0)
	v1275 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1277
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1267)+32))
	v1283 = m.T0[v1282].(func(*base.Module, int32, int32, int32) int64)(m, v1267+int32(8), v1268, int32(0))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1275
	v1287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1269)+4)))
	v1289 = v1287 & int32(_a_F_ExecParallelHashJoin_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1269)+4)) = uint16(v1289)
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+12))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1269)+6)) = uint16(v1292)
	v1986 = v1269
	goto L10
L298:
	;
	v1295 = *(*float64)(unsafe.Add(mBase, uint32(v1294)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v1294)+432)) = base.F64_add(v1295, float64(1))
	goto L300
L299:
	;
	goto L300
L300:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v1299)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
	if v1303 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L1
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v1308 = F_tuplestore_gettupleslot_force(m, v1306, v1307)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L1
	} else {
		goto L306
	}
L305:
	;
	goto L304
L306:
	;
	if v1308 != 0 {
		goto L288
	} else {
		goto L307
	}
L307:
	;
	goto L289
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	if v1339 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(7)
	goto L4
L310:
	;
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(8)
	goto L4
L312:
	;
	if v1348 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	goto L316
L314:
	;
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(8)
	goto L4
L316:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v1374
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v1376
	if v31 != 0 {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	goto L315
L318:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1420 != 0 {
		goto L326
	} else {
		goto L327
	}
L319:
	;
	v1378 = int32(_a_F_ExecParallelHashJoin_0)
	v1379 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1381
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v1386 = m.T0[v1385].(func(*base.Module, int32, int32, int32) int64)(m, v31, v35, v27+int32(44))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+80))
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+24))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+8))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+12))
	m.T0[v1397].(func(*base.Module, int32))(m, v1395)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L324
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1379
	if v1386 == int64(0) {
		goto L318
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	v1400 = int32(_a_F_ExecParallelHashJoin_0)
	v1401 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1403
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+32))
	v1409 = m.T0[v1408].(func(*base.Module, int32, int32, int32) int64)(m, v1393+int32(8), v1394, int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1401
	v1413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1395)+4)))
	v1415 = v1413 & int32(_a_F_ExecParallelHashJoin_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1395)+4)) = uint16(v1415)
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+12))
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1417)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1395)+6)) = uint16(v1418)
	v1986 = v1395
	goto L10
L326:
	;
	v1421 = *(*float64)(unsafe.Add(mBase, uint32(v1420)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v1420)+432)) = base.F64_add(v1421, float64(1))
	goto L328
L327:
	;
	goto L328
L328:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v1425)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
	if v1429 != 0 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1436 = F_tuplestore_gettupleslot(m, v1432, int32(1), int32(0), v1435)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L334
	}
L333:
	;
	goto L332
L334:
	;
	if v1436 != 0 {
		goto L316
	} else {
		goto L335
	}
L335:
	;
	goto L317
L336:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+48))
	if int32(0) <= v1468 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1465+v1468*int32(36))+26)) = uint8(v1474)
	F_ExecHashTableDetachBatch(m, v1464)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+140))
	v1481 = base.AtomicRmwAdd32(m, v1478, int32(164), int32(1))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+44))
	v1483 = base.I32_rem_u_s(v1481, v1482)
	v1487 = v1483
	goto L341
L340:
	;
	goto L339
L341:
	;
	v1509 = v1487 * int32(36)
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+144))
	v1511 = v1509 + v1510
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511)+26)))
	if v1512 != 0 {
		goto L343
	} else {
		goto L344
	}
L342:
	;
	v1986 = v2
	goto L10
L343:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+44))
	v1977 = base.I32_rem_s(v1487+int32(1), v1976)
	if v1977 != v1483 {
		v1487 = v1977
		goto L341
	} else {
		goto L406
	}
L344:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	v1515 = v1513 + int32(4)
	v1516 = F_BarrierAttach(m, v1515)
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L352
	}
L345:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1464, v1487)
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L1
	} else {
		goto L404
	}
L346:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L1
	} else {
		goto L401
	}
L347:
	;
	F_BarrierDetach(m, v1515)
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L1
	} else {
		goto L400
	}
L348:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1464, v1487)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L398
	}
L349:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1464, v1487)
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L1
	} else {
		goto L370
	}
L350:
	;
	v1708 = F_BarrierArriveAndWait(m, v1515, int32(134217742))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L1
	} else {
		goto L369
	}
L351:
	;
	v1519 = F_BarrierArriveAndWait(m, v1515, int32(134217743))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L353
	}
L352:
	;
	switch v1516 {
	case 0:
		goto L351
	case 1:
		goto L350
	case 2:
		goto L349
	case 3:
		goto L348
	case 4:
		goto L345
	case 5:
		goto L347
	default:
		goto L346
	}
L353:
	;
	if v1519 == int32(0) {
		goto L350
	} else {
		goto L354
	}
L354:
	;
	v1523 = int32(0)
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+144))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1524+v1487*int32(36))))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+136))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+140))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+16))
	v1535 = F_dsa_allocate_extended(m, v1529, v1531<<(uint(int32(2))%32), v1523)
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1528))) = v1535
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+136))
	v1539 = F_dsa_get_address(m, v1538, v1535)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	if v1531 <= int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	goto L350
L358:
	;
	v1544 = v1531 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v1531) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1556 = v1523
	v1564 = int32(0)
	goto L362
L360:
	;
	v1606 = v1523
	goto L361
L361:
	;
	v1627 = int32(0)
	v1631 = v1606
	goto L366
L362:
	;
	v1576 = v1539 + v1556<<(uint(int32(2))%32)
	v1577 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1576))) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1576)+4)) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1576)+8)) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1576)+12)) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1576)+16)) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1576)+20)) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1576)+24)) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1576)+28)) = v1577
	v1593 = int32(8)
	v1594 = v1556 + v1593
	v1596 = v1564 + v1593
	if v1596 != v1531&int32(-8) {
		v1556 = v1594
		v1564 = v1596
		goto L362
	} else {
		goto L364
	}
L363:
	;
	if v1544 == int32(0) {
		goto L357
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	v1606 = v1594
	goto L361
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1539+v1631<<(uint(int32(2))%32)))) = int32(0)
	v1654 = int32(1)
	v1657 = v1627 + v1654
	if v1657 != v1544 {
		v1627 = v1657
		v1631 = v1631 + v1654
		goto L366
	} else {
		goto L368
	}
L367:
	;
	goto L357
L368:
	;
	goto L367
L369:
	;
	goto L349
L370:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+144))
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1736+v1509)+28))
	F_sts_begin_parallel_scan(m, v1738)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	v1743 = F_sts_parallel_scan_next(m, v1738, v27+int32(44))
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	if v1743 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1747 = v1743
	goto L376
L374:
	;
	goto L375
L375:
	;
	F_sts_end_parallel_scan(m, v1738)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L1
	} else {
		goto L396
	}
L376:
	;
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	F_ExecForceStoreMinimalTuple(m, v1747, v1769, int32(0))
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L1
	} else {
		goto L378
	}
L377:
	;
	goto L375
L378:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	v1775 = m.G0
	v1777 = v1775 - int32(16)
	m.G0 = v1777
	v1781 = F_ExecFetchSlotMinimalTuple(m, v1773, v1777+int32(15))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1464)))
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1781)))
	v1785 = int32(8)
	v1789 = F_ExecParallelHashTupleAlloc(m, v1464, v1784+v1785, v1777+v1785)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1789)+4)) = v1774
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1781)))
	if v1792 != 0 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	base.MemoryCopy(m, v1789+int32(8), v1781, v1792)
	goto L383
L382:
	;
	goto L383
L383:
	;
	v1796 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1789)+18)))
	v1798 = v1796 & int32(_a_F_ExecParallelHashJoin_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v1789)+18)) = uint16(v1798)
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+8))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+20))
	v1807 = v1801 + (v1783-int32(1))&v1774<<(uint(int32(2))%32)
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	*(*int32)(unsafe.Add(mBase, uint32(v1789))) = v1808
	v1811 = base.AtomicRmwCmpxchg32(m, v1807, int32(0), v1808, v1800)
	if v1808 != v1811 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1829 = v1811
	goto L387
L385:
	;
	goto L386
L386:
	;
	v1867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1777)+15)))
	if v1867 == int32(1) {
		goto L390
	} else {
		goto L391
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1789))) = v1829
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	*(*int32)(unsafe.Add(mBase, uint32(v1789))) = v1838
	v1841 = base.AtomicRmwCmpxchg32(m, v1807, int32(0), v1838, v1800)
	if v1838 != v1841 {
		v1829 = v1841
		goto L387
	} else {
		goto L389
	}
L388:
	;
	goto L386
L389:
	;
	goto L388
L390:
	;
	F_pfree(m, v1781)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L1
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	m.G0 = v1777 + int32(16)
	v1877 = F_sts_parallel_scan_next(m, v1738, v27+int32(44))
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L1
	} else {
		goto L394
	}
L393:
	;
	goto L392
L394:
	;
	if v1877 != 0 {
		v1747 = v1877
		goto L376
	} else {
		goto L395
	}
L395:
	;
	goto L377
L396:
	;
	v1906 = F_BarrierArriveAndWait(m, v1515, int32(134217744))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	goto L348
L398:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+144))
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1934+v1509)+32))
	F_sts_begin_parallel_scan(m, v1936)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(2)
	goto L4
L400:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+144))
	v1945 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1943+v1509)+26)) = uint8(v1945)
	*(*int32)(unsafe.Add(mBase, uint32(v1464)+48)) = int32(-1)
	goto L343
L401:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1953
	F_errmsg_internal(m, int32(_a_F_ExecParallelHashJoin_4), v27+int32(16))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	F_errfinish(m, int32(_a_F_ExecParallelHashJoin_5), int32(1549), int32(_a_F_ExecParallelHashJoin_6))
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L404:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1464)+144))
	v1969 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1967+v1509)+26)) = uint8(v1969)
	F_ExecHashTableDetachBatch(m, v1464)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	goto L343
L406:
	;
	goto L342
L407:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v2011
	F_errmsg_internal(m, int32(_a_F_ExecParallelHashJoin_7), v27)
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	F_errfinish(m, int32(_a_F_ExecParallelHashJoin_5), int32(790), int32(_a_F_ExecParallelHashJoin_8))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecParallelInitializeWorker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	if l0 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v10 - int32(403) {
		case 0:
			v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+36)))
			if v202 != int32(1) {
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			} else {
				v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v207 = int64(*(*int32)(unsafe.Add(mBase, uint32(v206)+40)))
				v209 = F_shm_toc_lookup(m, v205, v207, int32(0))
				mBase = m.M
				v210 = m.ExcPending
				if v210 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(743)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v209
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		default:
			v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
			mBase = m.M
			v325 = m.ExcPending
			if v325 != 0 {
				return int32(0)
			} else {
				return v324
			}
		case 6:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+36)))
			if v14 == int32(1) {
				v17 = F_ScanRelIsReadOnly(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+132))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+40)))
					v27 = F_shm_toc_lookup(m, v23, v25, int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						if v17 != 0 {
							v36 = int32(1024)
						} else {
							v36 = int32(0)
						}
						v38 = F_table_beginscan_parallel(m, v29, v27, v22<<(uint(int32(7))%32)&int32(2048)|v36)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v38
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+132)))
							if v45&int32(16) != 0 {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49)+40)))
								v54 = F_shm_toc_lookup(m, v48, v50-int64(3458764513820540928), int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v54
									v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
									mBase = m.M
									v325 = m.ExcPending
									if v325 != 0 {
										return int32(0)
									} else {
										return v324
									}
								}
							} else {
								v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
								mBase = m.M
								v325 = m.ExcPending
								if v325 != 0 {
									return int32(0)
								} else {
									return v324
								}
							}
						}
					}
				}
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+132)))
				if v45&int32(16) != 0 {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49)+40)))
					v54 = F_shm_toc_lookup(m, v48, v50-int64(3458764513820540928), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v54
						v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							return v324
						}
					}
				} else {
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		case 8:
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
			if v58 == int32(1) {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+40)))
				v65 = F_shm_toc_lookup(m, v61, v63, int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
					v74 = F_ScanRelIsReadOnly(m, l0)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						if v74 != 0 {
							v76 = int32(1024)
						} else {
							v76 = int32(0)
						}
						v77 = F_index_beginscan_parallel(m, v67, v68, v69, v70, v71, v65, v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v77
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							if v80 != 0 {
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
								if v81 != int32(1) {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v91 != 0 {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v93)+40)))
										v98 = F_shm_toc_lookup(m, v92, v94-int64(3458764513820540928), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v98
											v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
											mBase = m.M
											v325 = m.ExcPending
											if v325 != 0 {
												return int32(0)
											} else {
												return v324
											}
										}
									} else {
										v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
										mBase = m.M
										v325 = m.ExcPending
										if v325 != 0 {
											return int32(0)
										} else {
											return v324
										}
									}
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									F_index_rescan(m, v77, v84, v85, v86, v87)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v91 != 0 {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v93)+40)))
											v98 = F_shm_toc_lookup(m, v92, v94-int64(3458764513820540928), int32(0))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v98
												v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
												mBase = m.M
												v325 = m.ExcPending
												if v325 != 0 {
													return int32(0)
												} else {
													return v324
												}
											}
										} else {
											v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
											mBase = m.M
											v325 = m.ExcPending
											if v325 != 0 {
												return int32(0)
											} else {
												return v324
											}
										}
									}
								}
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								F_index_rescan(m, v77, v84, v85, v86, v87)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v91 != 0 {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v93)+40)))
										v98 = F_shm_toc_lookup(m, v92, v94-int64(3458764513820540928), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v98
											v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
											mBase = m.M
											v325 = m.ExcPending
											if v325 != 0 {
												return int32(0)
											} else {
												return v324
											}
										}
									} else {
										v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
										mBase = m.M
										v325 = m.ExcPending
										if v325 != 0 {
											return int32(0)
										} else {
											return v324
										}
									}
								}
							}
						}
					}
				}
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v91 != 0 {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v93)+40)))
					v98 = F_shm_toc_lookup(m, v92, v94-int64(3458764513820540928), int32(0))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v98
						v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							return v324
						}
					}
				} else {
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		case 9:
			v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+36)))
			if v102 == int32(1) {
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v106)+40)))
				v109 = F_shm_toc_lookup(m, v105, v107, int32(0))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					v118 = F_ScanRelIsReadOnly(m, l0)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						if v118 != 0 {
							v120 = int32(1024)
						} else {
							v120 = int32(0)
						}
						v121 = F_index_beginscan_parallel(m, v111, v112, v113, v114, v115, v109, v120)
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v121
							v124 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v121)+28)) = uint8(v124)
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
							if v126 != 0 {
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
								if v127 != int32(1) {
									F_ExecBitmapIndexScanInitializeWorker(m, l0, l1)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int32(0)
									} else {
										v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
										mBase = m.M
										v325 = m.ExcPending
										if v325 != 0 {
											return int32(0)
										} else {
											return v324
										}
									}
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
									F_index_rescan(m, v130, v131, v132, v133, v134)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										F_ExecBitmapIndexScanInitializeWorker(m, l0, l1)
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int32(0)
										} else {
											v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
											mBase = m.M
											v325 = m.ExcPending
											if v325 != 0 {
												return int32(0)
											} else {
												return v324
											}
										}
									}
								}
							} else {
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								F_index_rescan(m, v130, v131, v132, v133, v134)
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
									return int32(0)
								} else {
									F_ExecBitmapIndexScanInitializeWorker(m, l0, l1)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int32(0)
									} else {
										v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
										mBase = m.M
										v325 = m.ExcPending
										if v325 != 0 {
											return int32(0)
										} else {
											return v324
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_ExecBitmapIndexScanInitializeWorker(m, l0, l1)
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return int32(0)
				} else {
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		case 10:
			F_ExecBitmapIndexScanInitializeWorker(m, l0, l1)
			mBase = m.M
			v141 = m.ExcPending
			if v141 != 0 {
				return int32(0)
			} else {
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			}
		case 11:
			v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+36)))
			if v232 == int32(1) {
				v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v237 = int64(*(*int32)(unsafe.Add(mBase, uint32(v236)+40)))
				v239 = F_shm_toc_lookup(m, v235, v237, int32(0))
				mBase = m.M
				v240 = m.ExcPending
				if v240 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v239
					v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v242 != 0 {
						v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v245 = int64(*(*int32)(unsafe.Add(mBase, uint32(v244)+40)))
						v249 = F_shm_toc_lookup(m, v243, v245-int64(3458764513820540928), int32(0))
						mBase = m.M
						v250 = m.ExcPending
						if v250 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v249
							v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
							mBase = m.M
							v325 = m.ExcPending
							if v325 != 0 {
								return int32(0)
							} else {
								return v324
							}
						}
					} else {
						v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							return v324
						}
					}
				}
			} else {
				v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v242 != 0 {
					v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v245 = int64(*(*int32)(unsafe.Add(mBase, uint32(v244)+40)))
					v249 = F_shm_toc_lookup(m, v243, v245-int64(3458764513820540928), int32(0))
					mBase = m.M
					v250 = m.ExcPending
					if v250 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v249
						v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							return v324
						}
					}
				} else {
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		case 13:
			v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+36)))
			if v160 == int32(1) {
				v163 = F_ScanRelIsReadOnly(m, l0)
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return int32(0)
				} else {
					v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+132))
					v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v169 = int64(*(*int32)(unsafe.Add(mBase, uint32(v168)+40)))
					v171 = F_shm_toc_lookup(m, v167, v169, int32(0))
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						if v163 != 0 {
							v180 = int32(1024)
						} else {
							v180 = int32(0)
						}
						v182 = F_table_beginscan_parallel_tidrange(m, v173, v171, v166<<(uint(int32(7))%32)&int32(2048)|v180)
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v182
							v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+132)))
							if v189&int32(16) != 0 {
								v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v194 = int64(*(*int32)(unsafe.Add(mBase, uint32(v193)+40)))
								v198 = F_shm_toc_lookup(m, v192, v194-int64(3458764513820540928), int32(0))
								mBase = m.M
								v199 = m.ExcPending
								if v199 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v198
									v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
									mBase = m.M
									v325 = m.ExcPending
									if v325 != 0 {
										return int32(0)
									} else {
										return v324
									}
								}
							} else {
								v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
								mBase = m.M
								v325 = m.ExcPending
								if v325 != 0 {
									return int32(0)
								} else {
									return v324
								}
							}
						}
					}
				}
			} else {
				v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+132)))
				if v189&int32(16) != 0 {
					v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v194 = int64(*(*int32)(unsafe.Add(mBase, uint32(v193)+40)))
					v198 = F_shm_toc_lookup(m, v192, v194-int64(3458764513820540928), int32(0))
					mBase = m.M
					v199 = m.ExcPending
					if v199 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v198
						v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							return v324
						}
					}
				} else {
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		case 21:
			v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+36)))
			if v143 != int32(1) {
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			} else {
				v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+160))
				if v147 != 0 {
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v150 = int64(*(*int32)(unsafe.Add(mBase, uint32(v149)+40)))
					v152 = F_shm_toc_lookup(m, v148, v150, int32(0))
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return int32(0)
					} else {
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+160))
						m.T0[v155].(func(*base.Module, int32, int32, int32))(m, l0, v154, v152)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
							mBase = m.M
							v325 = m.ExcPending
							if v325 != 0 {
								return int32(0)
							} else {
								return v324
							}
						}
					}
				} else {
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		case 22:
			v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+36)))
			if v215 != int32(1) {
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			} else {
				v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+40))
				if v219 != 0 {
					v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v222 = int64(*(*int32)(unsafe.Add(mBase, uint32(v221)+40)))
					v224 = F_shm_toc_lookup(m, v220, v222, int32(0))
					mBase = m.M
					v225 = m.ExcPending
					if v225 != 0 {
						return int32(0)
					} else {
						v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v227 = *(*int32)(unsafe.Add(mBase, uint32(v218)+40))
						m.T0[v227].(func(*base.Module, int32, int32, int32))(m, l0, v226, v224)
						mBase = m.M
						v229 = m.ExcPending
						if v229 != 0 {
							return int32(0)
						} else {
							v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
							mBase = m.M
							v325 = m.ExcPending
							if v325 != 0 {
								return int32(0)
							} else {
								return v324
							}
						}
					}
				} else {
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			}
		case 26:
			v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+36)))
			if v253 != int32(1) {
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			} else {
				v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v258 = int64(*(*int32)(unsafe.Add(mBase, uint32(v257)+40)))
				v260 = F_shm_toc_lookup(m, v256, v258, int32(0))
				mBase = m.M
				v261 = m.ExcPending
				if v261 != 0 {
					return int32(0)
				} else {
					v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_SharedFileSetAttach(m, v260+int32(168), v264)
					mBase = m.M
					v266 = m.ExcPending
					if v266 != 0 {
						return int32(0)
					} else {
						v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v267)+136)) = v260
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(678)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(766)
						v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							return v324
						}
					}
				}
			}
		case 28:
			v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v315 = int64(*(*int32)(unsafe.Add(mBase, uint32(v314)+40)))
			v317 = F_shm_toc_lookup(m, v313, v315, int32(1))
			mBase = m.M
			v318 = m.ExcPending
			if v318 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v317
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			}
		case 29:
			v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v290 = int64(*(*int32)(unsafe.Add(mBase, uint32(v289)+40)))
			v292 = F_shm_toc_lookup(m, v288, v290, int32(1))
			mBase = m.M
			v293 = m.ExcPending
			if v293 != 0 {
				return int32(0)
			} else {
				v294 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v294)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v292
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			}
		case 30:
			v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v299 = int64(*(*int32)(unsafe.Add(mBase, uint32(v298)+40)))
			v301 = F_shm_toc_lookup(m, v297, v299, int32(1))
			mBase = m.M
			v302 = m.ExcPending
			if v302 != 0 {
				return int32(0)
			} else {
				v303 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)) = uint8(v303)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v301
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			}
		case 32:
			v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v308 = int64(*(*int32)(unsafe.Add(mBase, uint32(v307)+40)))
			v310 = F_shm_toc_lookup(m, v306, v308, int32(1))
			mBase = m.M
			v311 = m.ExcPending
			if v311 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v310
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			}
		case 37:
			v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v273 != 0 {
				v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v276 = int64(*(*int32)(unsafe.Add(mBase, uint32(v275)+40)))
				v278 = F_shm_toc_lookup(m, v274, v276, int32(0))
				mBase = m.M
				v279 = m.ExcPending
				if v279 != 0 {
					return int32(0)
				} else {
					v281 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelInitializeWorker[0]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v278 + v281*int32(20) + int32(4)
					v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						return v324
					}
				}
			} else {
				v324 = F_planstate_tree_walker_impl(m, l0, int32(676), l1)
				mBase = m.M
				v325 = m.ExcPending
				if v325 != 0 {
					return int32(0)
				} else {
					return v324
				}
			}
		}
	}
}
func F_ExecParallelReportInstrumentation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_InstrEndLoop(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v19 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v63 = int32(440)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelReportInstrumentation[0]))
	v70 = l1 + v59 + v61*v27*v63 + v67*v63
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v70)+392))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v71)+392))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+392)) = v72 + v73
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v70)+184))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v71)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+184)) = v76 + v77
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v71)+400))
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v70)+400))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+400)) = base.F64_add(v80, v81)
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v71)+408))
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v70)+408))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+408)) = base.F64_add(v84, v85)
	v88 = *(*float64)(unsafe.Add(mBase, uint32(v71)+416))
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v70)+416))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+416)) = base.F64_add(v88, v89)
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v71)+424))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v70)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+424)) = base.F64_add(v92, v93)
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v71)+432))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(v70)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+432)) = base.F64_add(v96, v97)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v100 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	v27 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v27<<(uint(int32(2))%32))))
	if v34 == v13 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v37 = v27 + int32(1)
	if v37 != v19 {
		v27 = v37
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	F_errmsg_internal(m, int32(_a_F_ExecParallelReportInstrumentation_0), v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_ExecParallelReportInstrumentation_1), int32(1377), int32(_a_F_ExecParallelReportInstrumentation_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v191 = F_planstate_tree_walker_impl(m, l0, int32(677), l1)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L21
	}
L15:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v70)+192))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v71)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+192)) = v103 + v104
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v70)+200))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v71)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+200)) = v107 + v108
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v70)+208))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v71)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+208)) = v111 + v112
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v70)+216))
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v71)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+216)) = v115 + v116
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v70)+224))
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v71)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+224)) = v119 + v120
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v70)+232))
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v71)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+232)) = v123 + v124
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v70)+240))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v71)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+240)) = v127 + v128
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v70)+248))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v71)+248))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+248)) = v131 + v132
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v70)+256))
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v71)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+256)) = v135 + v136
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v70)+264))
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v71)+264))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+264)) = v139 + v140
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v70)+272))
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v71)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+272)) = v143 + v144
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v70)+280))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v71)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+280)) = v147 + v148
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v70)+288))
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v71)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+288)) = v151 + v152
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v70)+296))
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v71)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+296)) = v155 + v156
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v70)+304))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v71)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+304)) = v159 + v160
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v70)+312))
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v71)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+312)) = v163 + v164
	goto L17
L16:
	;
	goto L17
L17:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+2)))
	if v167 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v70)+336))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v71)+336))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+336)) = v170 + v171
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v70)+320))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v71)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+320)) = v174 + v175
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v70)+328))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v71)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+328)) = v178 + v179
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v70)+344))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v71)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+344)) = v182 + v183
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v70)+352))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v71)+352))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+352)) = v186 + v187
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L14
L21:
	;
	m.G0 = v10 + int32(16)
	return v191
}
func F_ParallelQueryMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
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
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v339 int64
	_ = v339
	var v343 int64
	_ = v343
	var v347 int64
	_ = v347
	var v351 int64
	_ = v351
	var v355 int64
	_ = v355
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v362 int64
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int64
	_ = v388
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v405 int64
	_ = v405
	var v407 int64
	_ = v407
	var v411 int64
	_ = v411
	var v413 int64
	_ = v413
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
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
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v25 = F_shm_toc_lookup(m, l1, int64(-2305843009213693951), v3)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_shm_toc_lookup(m, l1, int64(-2305843009213693947), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[0]))
	v35 = v29 + v32<<(uint(int32(16))%32)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[1]))
	F_shm_mq_set_sender(m, v35, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = F_shm_mq_attach(m, v35, l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = F_CreateTupleQueueDestReceiver(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v46 = F_shm_toc_lookup(m, l1, int64(-2305843009213693946), int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v49 = v48
	goto L10
L9:
	;
	v49 = v3
	goto L10
L10:
	;
	v52 = F_shm_toc_lookup(m, l1, int64(-2305843009213693943), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v56 = F_shm_toc_lookup(m, l1, int64(-2305843009213693944), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = F_shm_toc_lookup(m, l1, int64(-2305843009213693950), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v62 = F_stringToNode(m, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v66 = F_shm_toc_lookup(m, l1, int64(-2305843009213693949), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v66
	v70 = v21 + int32(8)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v71 + v73
	v80 = F_palloc(m, v72<<(uint(v73)%32)+int32(32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v72
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = int32(862)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v80
	if v87 < v72 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v105 = v3
	goto L20
L18:
	;
	goto L19
L19:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[2]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	goto L24
L20:
	;
	v114 = int32(4)
	v116 = v80 + int32(32) + v105<<(uint(v114)%32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v120 + v114
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+10)) = uint16(v124)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v126 + int32(2)
	v132 = F_datumRestore(m, v70, v116+int32(8))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v132
	v136 = v105 + int32(1)
	if v136 != v72 {
		v105 = v136
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v160 = int32(0)
	v162 = F_CreateQueryDesc(m, v62, v56, v159, v160, v42, v80, v160, v49)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[3])) = v164
	F_pgstat_report_activity(m, int32(3), v164)
	mBase = m.M
	v170 = F_shm_toc_lookup(m, l1, int64(-2305843009213693945), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v172 = F_dsa_attach_in_place(m, v170, l0)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+36)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	F_ExecutorStart(m, v162, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v162)+48))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+172)) = v172
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v183 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v162)+48))
	v258 = F_ExecParallelInitializeWorker(m, v255, v21+int32(8))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	v186 = F_dsa_get_address(m, v172, v183)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v162)+44))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v186 + int32(4)
	if v189 <= int32(0) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v200 = int32(0)
	goto L33
L33:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v214 + int32(4)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v188)+92))
	v224 = v221 + v215*int32(24)
	v227 = F_datumRestore(m, v21+int32(8), v224+int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v224)+8)) = v227
	v233 = v200 + int32(1)
	if v233 != v189 {
		v200 = v233
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v162)+48))
	v264 = v261
	goto L40
L38:
	;
	base.MemoryCopy(m, int32(_a_F_ParallelQueryMain_0), int32(_a_F_ParallelQueryMain_1), int32(128))
	v339 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[5])) = v339
	v343 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[7])) = v343
	v347 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[9])) = v347
	v351 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[11])) = v351
	v355 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[13])) = v355
	goto L65
L39:
	;
	goto L38
L40:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	switch v266 - int32(400) {
	case 0:
		goto L46
	default:
		goto L39
	case 3:
		goto L49
	case 4:
		goto L50
	case 17:
		goto L45
	case 32:
		goto L48
	case 33:
		goto L47
	case 38, 39:
		goto L44
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v264)+120)) = v260
	v325 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+117)) = uint8(v325)
	goto L39
L42:
	;
	goto L41
L43:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v264+v321)))
	v264 = v323
	goto L40
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v264)+112)) = v260
	v321 = int32(36)
	goto L43
L45:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	if v316 == int32(0) {
		v321 = int32(116)
		goto L43
	} else {
		goto L64
	}
L46:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	if v314 != 0 {
		v264 = v314
		goto L40
	} else {
		goto L63
	}
L47:
	;
	if v260 < int64(0) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	if int64(0) <= v260 {
		goto L42
	} else {
		goto L59
	}
L49:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v286 <= int32(0) {
		goto L39
	} else {
		goto L55
	}
L50:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v269 <= int32(0) {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	v275 = int32(0)
	goto L52
L52:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v275<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v260, v280)
	mBase = m.M
	v283 = v275 + int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v283 < v284 {
		v275 = v283
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L39
L54:
	;
	goto L53
L55:
	;
	v292 = int32(0)
	goto L56
L56:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293+v292<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v260, v297)
	mBase = m.M
	v300 = v292 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v300 < v301 {
		v292 = v300
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L39
L58:
	;
	goto L57
L59:
	;
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+117)) = uint8(v305)
	goto L38
L60:
	;
	v309 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+116)) = uint8(v309)
	goto L38
L61:
	;
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v264)+120)) = v260
	v312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+116)) = uint8(v312)
	goto L38
L63:
	;
	goto L39
L64:
	;
	goto L39
L65:
	;
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v359 = int64(0)
	if v359 < v358 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v362 = v358
	goto L68
L67:
	;
	v362 = v359
	goto L68
L68:
	;
	F_ExecutorRun(m, v162, int32(1), v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_ExecutorFinish(m, v162)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v369 = F_shm_toc_lookup(m, l1, int64(-2305843009213693948), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v373 = F_shm_toc_lookup(m, l1, int64(-2305843009213693942), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[0]))
	v379 = v369 + v376<<(uint(int32(7))%32)
	v382 = v373 + v376*int32(40)
	base.MemoryFill(m, v379, int32(0), int32(128))
	F_BufferUsageAccumDiff(m, v379, int32(_a_F_ParallelQueryMain_0))
	mBase = m.M
	v388 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v382)+32)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v382)+24)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v382)+16)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v382)+8)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v382))) = v388
	v399 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[8]))
	v401 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v382)+16)) = v399 - v401
	v405 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[12]))
	v407 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v382))) = v405 - v407
	v411 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[10]))
	v413 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v382)+8)) = v411 - v413
	v417 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[6]))
	v419 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v382)+24)) = v417 - v419
	v423 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[4]))
	v425 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v382)+32)) = v423 - v425
	goto L73
L73:
	;
	if v46 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v162)+48))
	v429 = F_ExecParallelReportInstrumentation(m, v428, v46)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v162)+44))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+180))
	v433 = int32(0)
	if base.B2i32(v432 == v433)|base.B2i32(v52 == v433) == v433 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[0]))
	v444 = v52 + v441*int32(48)
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v432)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v444)+48)) = v445
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v432)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v444)+40)) = v447
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v432)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v444)+32)) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v432)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v444)+24)) = v451
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v432)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v444)+16)) = v453
	v455 = *(*int64)(unsafe.Add(mBase, uint32(v432)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v444)+8)) = v455
	goto L80
L79:
	;
	goto L80
L80:
	;
	F_ExecutorEnd(m, v162)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_dsa_detach(m, v172)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_FreeQueryDesc(m, v162)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	m.T0[v464].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	m.G0 = v21 + int32(16)
	return
}
func F_ReinitializeParallelDSM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v6 = int32(_a_F_ReinitializeParallelDSM_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[0]))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[0])) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = F_shm_toc_lookup(m, v29, int64(-65535), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L8
	}
L2:
	;
	F_WaitForParallelWorkersToFinish(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	F_WaitForParallelWorkersToExit(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v21 == v19 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_pfree(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = int64(0)
	goto L1
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = int64(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v36 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[0])) = v7
	return
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v42 = F_shm_toc_lookup(m, v39, int64(-65534), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v44 <= int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(0)
	goto L13
L13:
	;
	v55 = v42 + v49<<(uint(int32(14))%32)
	v57 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v55))), uint32(v57))
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v55)+4)) = v60
	v64 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v55)+36)) = uint16(v64)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = int32(_a_F_ReinitializeParallelDSM_1)
	goto L15
L14:
	;
	goto L9
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[2]))
	F_shm_mq_set_receiver(m, v55, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v78 = F_shm_mq_attach(m, v55, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v80+v49<<(uint(int32(3))%32))+4)) = v78
	v86 = v49 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v86 < v87 {
		v49 = v86
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
}
func F_parallel_vacuum_process_all_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v340 int64
	_ = v340
	var v343 int32
	_ = v343
	var v345 int64
	_ = v345
	var v346 int64
	_ = v346
	var v349 int32
	_ = v349
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v355 int32
	_ = v355
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v361 int32
	_ = v361
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v367 int32
	_ = v367
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v373 int32
	_ = v373
	var v375 int64
	_ = v375
	var v376 int64
	_ = v376
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v385 int32
	_ = v385
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v391 int32
	_ = v391
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v397 int32
	_ = v397
	var v399 int64
	_ = v399
	var v400 int64
	_ = v400
	var v403 int32
	_ = v403
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v409 int32
	_ = v409
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v415 int32
	_ = v415
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v421 int32
	_ = v421
	var v423 int64
	_ = v423
	var v424 int64
	_ = v424
	var v427 int32
	_ = v427
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v433 int32
	_ = v433
	var v435 int64
	_ = v435
	var v436 int64
	_ = v436
	var v439 int32
	_ = v439
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v445 int32
	_ = v445
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v454 int64
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v26 = v24 - int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v26 < v28 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = int32(1)
	v24 = v17
	goto L1
L3:
	;
	goto L4
L4:
	;
	v18 = int32(2)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if l1 != 0 {
		v23 = v18
		v24 = v19
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v23 = v18
	v24 = v20 + v19
	goto L1
L6:
	;
	v30 = v26
	goto L8
L7:
	;
	v30 = v28
	goto L8
L8:
	;
	v31 = int32(0)
	if base.B2i32(l3 == int32(0))|base.B2i32(v30 <= v31) == v31 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v36 + v30
	goto L11
L10:
	;
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v39 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+44)) = v101
	if v101 < v30 {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v55 = v52 + v47*int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v23
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v47))))
	if v60 != int32(1) {
		v85 = int32(0)
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+4)) = uint8(v85)
	v88 = v47 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v88 < v89 {
		v47 = v88
		goto L15
	} else {
		goto L21
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v47<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+204))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+29)))
	if l2 != 0 {
		v85 = v69 & int32(1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v72 = int32(0)
	if v69&int32(6) == v72 {
		v85 = v72
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v79 = int32(0)
	v85 = base.B2i32(v69&int32(2) == v79) | base.B2i32(l1 <= v79)
	goto L17
L21:
	;
	goto L16
L22:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v583 {
		goto L118
	} else {
		goto L119
	}
L23:
	;
	if int32(0) < l1 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v463 != 0 {
		goto L90
	} else {
		goto L91
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ReinitializeParallelDSM(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+36)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+40)) = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	if v118 < v30 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	return
L30:
	;
	goto L28
L31:
	;
	v120 = v118
	goto L33
L32:
	;
	v120 = v30
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LaunchParallelWorkers(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
	if v126 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v155 = F_errstart(m, v153, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L29
	} else {
		goto L38
	}
L36:
	;
	v130 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[2])) = v130
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[1])) = v130
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0])) = v136 + int32(40)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[3])) = v136 + int32(36)
	if l3 == v130 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v146 + v148
	goto L35
L38:
	;
	if l2 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v197 != 0 {
		goto L55
	} else {
		goto L56
	}
L40:
	;
	F_errfinish(m, int32(_a_F_parallel_vacuum_process_all_indexes_0), v191, int32(_a_F_parallel_vacuum_process_all_indexes_1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L29
	} else {
		goto L54
	}
L41:
	;
	if v155 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v155 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v161
	if v161 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v168 = int32(_a_F_parallel_vacuum_process_all_indexes_2)
	goto L47
L46:
	;
	v168 = int32(_a_F_parallel_vacuum_process_all_indexes_3)
	goto L47
L47:
	;
	F_errmsg(m, v168, v12+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v191 = int32(920)
	goto L40
L49:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v177
	if v177 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v184 = int32(_a_F_parallel_vacuum_process_all_indexes_4)
	goto L52
L51:
	;
	v184 = int32(_a_F_parallel_vacuum_process_all_indexes_5)
	goto L52
L52:
	;
	F_errmsg(m, v184, v12+int32(32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L29
	} else {
		goto L53
	}
L53:
	;
	v191 = int32(926)
	goto L40
L54:
	;
	goto L39
L55:
	;
	v200 = base.AtomicRmwAdd32(m, v197, int32(0), int32(1))
	goto L57
L56:
	;
	goto L57
L57:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v201 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v206 = v201
	v209 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v243 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v217 = v214 + v209*int32(48)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+4)))
	if v218 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L60
L63:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221+v209<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v225, v217)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L29
	} else {
		goto L66
	}
L64:
	;
	v229 = v206
	goto L65
L65:
	;
	v231 = v209 + int32(1)
	if v231 < v229 {
		v206 = v229
		v209 = v231
		goto L61
	} else {
		goto L67
	}
L66:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v229 = v228
	goto L65
L67:
	;
	goto L62
L68:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v260 = base.AtomicRmwAdd32(m, v257, int32(44), int32(1))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v260 < v261 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v247 = int32(0)
	v248 = base.AtomicRmwSub32(m, v243, v247, int32(1))
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v250 == v247 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v255 = base.AtomicRmwAdd32(m, v250, int32(0), int32(1))
	goto L68
L71:
	;
	v267 = v260
	goto L74
L72:
	;
	goto L73
L73:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v302 != 0 {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v275 = v272 + v267*int32(48)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+4)))
	if v276 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L73
L76:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v267<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v283, v275)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L29
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v289 = base.AtomicRmwAdd32(m, v286, int32(44), int32(1))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v289 < v290 {
		v267 = v289
		goto L74
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	goto L75
L81:
	;
	v305 = base.AtomicRmwSub32(m, v302, int32(0), int32(1))
	goto L83
L82:
	;
	goto L83
L83:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_WaitForParallelWorkersToFinish(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L29
	} else {
		goto L84
	}
L84:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	if v310 <= int32(0) {
		goto L22
	} else {
		goto L85
	}
L85:
	;
	v318 = int32(0)
	goto L86
L86:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v326 = v323 + v318<<(uint(int32(7))%32)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v330 = v327 + v318*int32(40)
	v331 = int32(_a_F_parallel_vacuum_process_all_indexes_6)
	v333 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[4]))
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v326)))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[4])) = v333 + v334
	v337 = int32(_a_F_parallel_vacuum_process_all_indexes_7)
	v339 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[5]))
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v326)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[5])) = v339 + v340
	v343 = int32(_a_F_parallel_vacuum_process_all_indexes_8)
	v345 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[6]))
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v326)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[6])) = v345 + v346
	v349 = int32(_a_F_parallel_vacuum_process_all_indexes_9)
	v351 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[7]))
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v326)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[7])) = v351 + v352
	v355 = int32(_a_F_parallel_vacuum_process_all_indexes_10)
	v357 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[8]))
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v326)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[8])) = v357 + v358
	v361 = int32(_a_F_parallel_vacuum_process_all_indexes_11)
	v363 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[9]))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v326)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[9])) = v363 + v364
	v367 = int32(_a_F_parallel_vacuum_process_all_indexes_12)
	v369 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[10]))
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v326)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[10])) = v369 + v370
	v373 = int32(_a_F_parallel_vacuum_process_all_indexes_13)
	v375 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[11]))
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v326)+56))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[11])) = v375 + v376
	v379 = int32(_a_F_parallel_vacuum_process_all_indexes_14)
	v381 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[12]))
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v326)+64))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[12])) = v381 + v382
	v385 = int32(_a_F_parallel_vacuum_process_all_indexes_15)
	v387 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[13]))
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v326)+72))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[13])) = v387 + v388
	v391 = int32(_a_F_parallel_vacuum_process_all_indexes_16)
	v393 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[14]))
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v326)+80))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[14])) = v393 + v394
	v397 = int32(_a_F_parallel_vacuum_process_all_indexes_17)
	v399 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[15]))
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v326)+88))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[15])) = v399 + v400
	v403 = int32(_a_F_parallel_vacuum_process_all_indexes_18)
	v405 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[16]))
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v326)+96))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[16])) = v405 + v406
	v409 = int32(_a_F_parallel_vacuum_process_all_indexes_19)
	v411 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[17]))
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v326)+104))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[17])) = v411 + v412
	v415 = int32(_a_F_parallel_vacuum_process_all_indexes_20)
	v417 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[18]))
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v326)+112))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[18])) = v417 + v418
	v421 = int32(_a_F_parallel_vacuum_process_all_indexes_21)
	v423 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[19]))
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v326)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[19])) = v423 + v424
	v427 = int32(_a_F_parallel_vacuum_process_all_indexes_22)
	v429 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[20]))
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v330)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[20])) = v429 + v430
	v433 = int32(_a_F_parallel_vacuum_process_all_indexes_23)
	v435 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[21]))
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v330)))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[21])) = v435 + v436
	v439 = int32(_a_F_parallel_vacuum_process_all_indexes_24)
	v441 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[22]))
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v330)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[22])) = v441 + v442
	v445 = int32(_a_F_parallel_vacuum_process_all_indexes_25)
	v447 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[23]))
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v330)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[23])) = v447 + v448
	v451 = int32(_a_F_parallel_vacuum_process_all_indexes_26)
	v453 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[24]))
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v330)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[24])) = v453 + v454
	goto L88
L87:
	;
	goto L22
L88:
	;
	v458 = v318 + int32(1)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+20))
	if v458 < v460 {
		v318 = v458
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v466 = base.AtomicRmwAdd32(m, v463, int32(0), int32(1))
	goto L92
L91:
	;
	goto L92
L92:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v467 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v472 = v467
	v475 = int32(0)
	goto L96
L94:
	;
	goto L95
L95:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v509 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v483 = v480 + v475*int32(48)
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
	if v484 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L95
L98:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487+v475<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v491, v483)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L29
	} else {
		goto L101
	}
L99:
	;
	v495 = v472
	goto L100
L100:
	;
	v497 = v475 + int32(1)
	if v497 < v495 {
		v472 = v495
		v475 = v497
		goto L96
	} else {
		goto L102
	}
L101:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v495 = v494
	goto L100
L102:
	;
	goto L97
L103:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v526 = base.AtomicRmwAdd32(m, v523, int32(44), int32(1))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v526 < v527 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v513 = int32(0)
	v514 = base.AtomicRmwSub32(m, v509, v513, int32(1))
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v516 == v513 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v521 = base.AtomicRmwAdd32(m, v516, int32(0), int32(1))
	goto L103
L106:
	;
	v533 = v526
	goto L109
L107:
	;
	goto L108
L108:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v568 == int32(0) {
		goto L22
	} else {
		goto L116
	}
L109:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v541 = v538 + v533*int32(48)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+4)))
	if v542 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L108
L111:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v545+v533<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v549, v541)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L29
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v555 = base.AtomicRmwAdd32(m, v552, int32(44), int32(1))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v555 < v556 {
		v533 = v555
		goto L109
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	goto L110
L116:
	;
	v573 = base.AtomicRmwSub32(m, v568, int32(0), int32(1))
	goto L22
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L29
	} else {
		goto L128
	}
L118:
	;
	v591 = int32(0)
	goto L121
L119:
	;
	goto L120
L120:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[3]))
	if v619 != 0 {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v599 = v596 + v591*int32(48)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	if v600 != int32(3) {
		goto L117
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v599))) = int32(0)
	v606 = v591 + int32(1)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v606 < v607 {
		v591 = v606
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[1])) = v621
	v624 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[3])) = v624
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0])) = v624
	goto L127
L126:
	;
	goto L127
L127:
	;
	m.G0 = v12 + int32(48)
	return
L128:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v636+v591<<(uint(int32(2))%32))))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v641 + int32(4)
	F_errmsg_internal(m, int32(_a_F_parallel_vacuum_process_all_indexes_27), v12)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L29
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_parallel_vacuum_process_all_indexes_0), int32(961), int32(_a_F_parallel_vacuum_process_all_indexes_1))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L29
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
