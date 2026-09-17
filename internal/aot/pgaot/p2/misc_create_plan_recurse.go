package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_plan_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v155 int32
	_ = v155
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v299 float64
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 float64
	_ = v306
	var v308 float64
	_ = v308
	var v310 float64
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
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
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 float64
	_ = v362
	var v363 float64
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 float64
	_ = v377
	var v380 int32
	_ = v380
	var v381 float64
	_ = v381
	var v387 float64
	_ = v387
	var v393 float64
	_ = v393
	var v395 float64
	_ = v395
	var v397 float64
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
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
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 float64
	_ = v466
	var v467 float64
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 float64
	_ = v481
	var v484 int32
	_ = v484
	var v485 float64
	_ = v485
	var v491 float64
	_ = v491
	var v497 float64
	_ = v497
	var v499 float64
	_ = v499
	var v501 float64
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v541 float64
	_ = v541
	var v543 float64
	_ = v543
	var v545 float64
	_ = v545
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 float64
	_ = v554
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
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
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v738 int32
	_ = v738
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v792 int32
	_ = v792
	var v829 int32
	_ = v829
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1209 float64
	_ = v1209
	var v1211 float64
	_ = v1211
	var v1213 float64
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1276 int32
	_ = v1276
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
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
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1385 int64
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1407 float64
	_ = v1407
	var v1409 float64
	_ = v1409
	var v1411 float64
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1503 int64
	_ = v1503
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1590 int32
	_ = v1590
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1703 int32
	_ = v1703
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1982 int32
	_ = v1982
	var v2015 int32
	_ = v2015
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2109 int32
	_ = v2109
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2156 int32
	_ = v2156
	var v2163 int32
	_ = v2163
	var v2169 int32
	_ = v2169
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2233 int32
	_ = v2233
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2389 int32
	_ = v2389
	var v2421 int32
	_ = v2421
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2440 int32
	_ = v2440
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2478 int32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2485 int32
	_ = v2485
	var v2490 int32
	_ = v2490
	var v2494 int32
	_ = v2494
	var v2497 int32
	_ = v2497
	var v2501 int32
	_ = v2501
	var v2506 int32
	_ = v2506
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2517 int32
	_ = v2517
	var v2522 int32
	_ = v2522
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2533 int32
	_ = v2533
	var v2538 int32
	_ = v2538
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2598 int32
	_ = v2598
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2648 int32
	_ = v2648
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2657 int32
	_ = v2657
	var v2667 int32
	_ = v2667
	var v2671 int32
	_ = v2671
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2681 int32
	_ = v2681
	var v2707 int32
	_ = v2707
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2714 int32
	_ = v2714
	var v2718 int32
	_ = v2718
	var v2721 int32
	_ = v2721
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2743 int32
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2757 int32
	_ = v2757
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2774 int32
	_ = v2774
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2789 int32
	_ = v2789
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2798 int32
	_ = v2798
	var v2801 int32
	_ = v2801
	var v2804 int32
	_ = v2804
	var v2808 int32
	_ = v2808
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2815 int32
	_ = v2815
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2856 int32
	_ = v2856
	var v2859 int32
	_ = v2859
	var v2866 int32
	_ = v2866
	var v2870 int32
	_ = v2870
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2901 int32
	_ = v2901
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2913 int32
	_ = v2913
	var v2918 int32
	_ = v2918
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2937 int32
	_ = v2937
	var v2948 int32
	_ = v2948
	var v2950 int32
	_ = v2950
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2970 int32
	_ = v2970
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2979 int32
	_ = v2979
	var v2982 int32
	_ = v2982
	var v2985 int32
	_ = v2985
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2995 int32
	_ = v2995
	var v2997 int32
	_ = v2997
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3016 int32
	_ = v3016
	var v3020 int32
	_ = v3020
	var v3028 int32
	_ = v3028
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3041 int32
	_ = v3041
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3058 int32
	_ = v3058
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3066 int32
	_ = v3066
	var v3069 int32
	_ = v3069
	var v3073 int32
	_ = v3073
	var v3078 int32
	_ = v3078
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3085 int32
	_ = v3085
	var v3087 int32
	_ = v3087
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3099 int32
	_ = v3099
	var v3100 int32
	_ = v3100
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3148 int32
	_ = v3148
	var v3150 float64
	_ = v3150
	var v3152 float64
	_ = v3152
	var v3154 float64
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3159 int32
	_ = v3159
	var v3161 int32
	_ = v3161
	var v3163 int32
	_ = v3163
	var v3164 int32
	_ = v3164
	var v3165 int32
	_ = v3165
	var v3166 int64
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3172 int32
	_ = v3172
	var v3174 int32
	_ = v3174
	var v3180 int32
	_ = v3180
	var v3182 float64
	_ = v3182
	var v3184 float64
	_ = v3184
	var v3186 float64
	_ = v3186
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3191 int32
	_ = v3191
	var v3193 int32
	_ = v3193
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3202 int32
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3218 int32
	_ = v3218
	var v3257 int32
	_ = v3257
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3267 int32
	_ = v3267
	var v3269 int32
	_ = v3269
	var v3270 int32
	_ = v3270
	var v3276 int32
	_ = v3276
	var v3280 int32
	_ = v3280
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3288 int32
	_ = v3288
	var v3330 float64
	_ = v3330
	var v3340 float64
	_ = v3340
	var v3343 float64
	_ = v3343
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3359 int32
	_ = v3359
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3376 int32
	_ = v3376
	var v3380 int32
	_ = v3380
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3438 int32
	_ = v3438
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3541 int32
	_ = v3541
	var v3543 float64
	_ = v3543
	var v3545 float64
	_ = v3545
	var v3547 float64
	_ = v3547
	var v3549 int32
	_ = v3549
	var v3550 int32
	_ = v3550
	var v3552 int32
	_ = v3552
	var v3554 int32
	_ = v3554
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3561 int32
	_ = v3561
	var v3564 int32
	_ = v3564
	var v3568 int32
	_ = v3568
	var v3570 int32
	_ = v3570
	var v3573 int32
	_ = v3573
	var v3609 int32
	_ = v3609
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3619 int32
	_ = v3619
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3628 int32
	_ = v3628
	var v3632 int32
	_ = v3632
	var v3633 int32
	_ = v3633
	var v3635 int32
	_ = v3635
	var v3636 int32
	_ = v3636
	var v3646 int32
	_ = v3646
	var v3682 int32
	_ = v3682
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3689 int32
	_ = v3689
	var v3690 float64
	_ = v3690
	var v3700 float64
	_ = v3700
	var v3703 float64
	_ = v3703
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3708 int32
	_ = v3708
	var v3710 int32
	_ = v3710
	var v3711 int32
	_ = v3711
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3726 int32
	_ = v3726
	var v3727 int32
	_ = v3727
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3735 int32
	_ = v3735
	var v3742 int32
	_ = v3742
	var v3744 int32
	_ = v3744
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3795 int32
	_ = v3795
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3803 int32
	_ = v3803
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3808 int32
	_ = v3808
	var v3811 int32
	_ = v3811
	var v3814 int32
	_ = v3814
	var v3815 int32
	_ = v3815
	var v3869 int32
	_ = v3869
	var v3871 float64
	_ = v3871
	var v3873 float64
	_ = v3873
	var v3875 float64
	_ = v3875
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3880 int32
	_ = v3880
	var v3882 int32
	_ = v3882
	var v3884 int32
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3893 int32
	_ = v3893
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3896 int32
	_ = v3896
	var v3900 int32
	_ = v3900
	var v3903 int32
	_ = v3903
	var v3907 int32
	_ = v3907
	var v3908 int32
	_ = v3908
	var v3913 int32
	_ = v3913
	var v3949 int32
	_ = v3949
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3955 int32
	_ = v3955
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3959 int32
	_ = v3959
	var v3961 int32
	_ = v3961
	var v3962 int32
	_ = v3962
	var v3968 int32
	_ = v3968
	var v3972 int32
	_ = v3972
	var v3973 int32
	_ = v3973
	var v3975 int32
	_ = v3975
	var v3976 int32
	_ = v3976
	var v3986 int32
	_ = v3986
	var v4024 int32
	_ = v4024
	var v4025 int32
	_ = v4025
	var v4027 int32
	_ = v4027
	var v4028 int32
	_ = v4028
	var v4029 int32
	_ = v4029
	var v4030 int32
	_ = v4030
	var v4031 int32
	_ = v4031
	var v4032 int32
	_ = v4032
	var v4035 int32
	_ = v4035
	var v4042 int32
	_ = v4042
	var v4086 int32
	_ = v4086
	var v4087 int32
	_ = v4087
	var v4089 int32
	_ = v4089
	var v4090 int32
	_ = v4090
	var v4091 int32
	_ = v4091
	var v4092 int32
	_ = v4092
	var v4093 int32
	_ = v4093
	var v4096 int32
	_ = v4096
	var v4099 int32
	_ = v4099
	var v4100 int32
	_ = v4100
	var v4101 int32
	_ = v4101
	var v4104 int32
	_ = v4104
	var v4105 int32
	_ = v4105
	var v4111 int32
	_ = v4111
	var v4153 int32
	_ = v4153
	var v4154 int32
	_ = v4154
	var v4156 int32
	_ = v4156
	var v4157 int32
	_ = v4157
	var v4158 int32
	_ = v4158
	var v4159 int32
	_ = v4159
	var v4160 int32
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4162 int32
	_ = v4162
	var v4165 int32
	_ = v4165
	var v4168 int32
	_ = v4168
	var v4216 int32
	_ = v4216
	var v4217 int32
	_ = v4217
	var v4219 int32
	_ = v4219
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4226 int32
	_ = v4226
	var v4229 int32
	_ = v4229
	var v4230 int32
	_ = v4230
	var v4231 int32
	_ = v4231
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4237 int32
	_ = v4237
	var v4281 int32
	_ = v4281
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4285 int32
	_ = v4285
	var v4286 int32
	_ = v4286
	var v4289 int32
	_ = v4289
	var v4291 int32
	_ = v4291
	var v4301 int32
	_ = v4301
	var v4303 int32
	_ = v4303
	var v4305 int32
	_ = v4305
	var v4309 int32
	_ = v4309
	var v4311 int32
	_ = v4311
	var v4313 int32
	_ = v4313
	var v4315 int32
	_ = v4315
	var v4317 int32
	_ = v4317
	var v4325 int32
	_ = v4325
	var v4327 float64
	_ = v4327
	var v4329 float64
	_ = v4329
	var v4331 float64
	_ = v4331
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4336 int32
	_ = v4336
	var v4338 int32
	_ = v4338
	var v4340 int32
	_ = v4340
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4346 int32
	_ = v4346
	var v4347 int32
	_ = v4347
	var v4348 int32
	_ = v4348
	var v4349 int32
	_ = v4349
	var v4352 int32
	_ = v4352
	var v4355 int32
	_ = v4355
	var v4358 int32
	_ = v4358
	var v4360 int32
	_ = v4360
	var v4361 int32
	_ = v4361
	var v4366 int32
	_ = v4366
	var v4370 int32
	_ = v4370
	var v4372 int32
	_ = v4372
	var v4373 int32
	_ = v4373
	var v4414 int32
	_ = v4414
	var v4415 int32
	_ = v4415
	var v4416 int32
	_ = v4416
	var v4417 int32
	_ = v4417
	var v4418 int32
	_ = v4418
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4421 int32
	_ = v4421
	var v4422 int32
	_ = v4422
	var v4424 int32
	_ = v4424
	var v4426 int32
	_ = v4426
	var v4428 int32
	_ = v4428
	var v4430 int32
	_ = v4430
	var v4431 int32
	_ = v4431
	var v4432 int32
	_ = v4432
	var v4434 int32
	_ = v4434
	var v4440 int32
	_ = v4440
	var v4443 int32
	_ = v4443
	var v4482 int32
	_ = v4482
	var v4486 int32
	_ = v4486
	var v4487 int32
	_ = v4487
	var v4489 int32
	_ = v4489
	var v4531 int32
	_ = v4531
	var v4532 int32
	_ = v4532
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4538 int32
	_ = v4538
	var v4542 int32
	_ = v4542
	var v4590 int32
	_ = v4590
	var v4632 int32
	_ = v4632
	var v4633 int32
	_ = v4633
	var v4634 int32
	_ = v4634
	var v4637 int32
	_ = v4637
	var v4643 int32
	_ = v4643
	var v4685 int32
	_ = v4685
	var v4689 int32
	_ = v4689
	var v4690 int32
	_ = v4690
	var v4691 int32
	_ = v4691
	var v4692 int32
	_ = v4692
	var v4693 int32
	_ = v4693
	var v4694 int32
	_ = v4694
	var v4697 int32
	_ = v4697
	var v4700 int32
	_ = v4700
	var v4701 int32
	_ = v4701
	var v4750 int32
	_ = v4750
	var v4753 int32
	_ = v4753
	var v4754 int32
	_ = v4754
	var v4755 int32
	_ = v4755
	var v4764 int32
	_ = v4764
	var v4770 int32
	_ = v4770
	var v4772 int32
	_ = v4772
	var v4801 int32
	_ = v4801
	var v4805 int32
	_ = v4805
	var v4806 int32
	_ = v4806
	var v4810 int32
	_ = v4810
	var v4811 int32
	_ = v4811
	var v4812 int32
	_ = v4812
	var v4813 int32
	_ = v4813
	var v4816 int32
	_ = v4816
	var v4817 int32
	_ = v4817
	var v4818 int32
	_ = v4818
	var v4819 int32
	_ = v4819
	var v4824 int32
	_ = v4824
	var v4866 int32
	_ = v4866
	var v4869 int32
	_ = v4869
	var v4873 int32
	_ = v4873
	var v4874 int32
	_ = v4874
	var v4878 int32
	_ = v4878
	var v4881 int32
	_ = v4881
	var v4882 int32
	_ = v4882
	var v4887 int32
	_ = v4887
	var v4928 int32
	_ = v4928
	var v4934 int32
	_ = v4934
	var v4936 int32
	_ = v4936
	var v4937 int32
	_ = v4937
	var v4938 int32
	_ = v4938
	var v4939 int32
	_ = v4939
	var v4942 int32
	_ = v4942
	var v4943 int32
	_ = v4943
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4947 int32
	_ = v4947
	var v4948 int32
	_ = v4948
	var v4949 int32
	_ = v4949
	var v4950 int32
	_ = v4950
	var v4951 int32
	_ = v4951
	var v4954 int32
	_ = v4954
	var v4959 int32
	_ = v4959
	var v5002 int32
	_ = v5002
	var v5003 int32
	_ = v5003
	var v5005 int32
	_ = v5005
	var v5007 int32
	_ = v5007
	var v5009 int32
	_ = v5009
	var v5013 int32
	_ = v5013
	var v5016 int32
	_ = v5016
	var v5019 int32
	_ = v5019
	var v5020 int32
	_ = v5020
	var v5024 int32
	_ = v5024
	var v5032 int32
	_ = v5032
	var v5033 int32
	_ = v5033
	var v5036 int32
	_ = v5036
	var v5047 int32
	_ = v5047
	var v5052 int32
	_ = v5052
	var v5055 int32
	_ = v5055
	var v5058 int32
	_ = v5058
	var v5059 int32
	_ = v5059
	var v5060 int32
	_ = v5060
	var v5063 int32
	_ = v5063
	var v5066 int32
	_ = v5066
	var v5067 int32
	_ = v5067
	var v5071 int32
	_ = v5071
	var v5114 int32
	_ = v5114
	var v5115 int32
	_ = v5115
	var v5118 int32
	_ = v5118
	var v5120 int32
	_ = v5120
	var v5122 int32
	_ = v5122
	var v5128 int32
	_ = v5128
	var v5137 int32
	_ = v5137
	var v5138 int32
	_ = v5138
	var v5148 int32
	_ = v5148
	var v5187 int32
	_ = v5187
	var v5188 int32
	_ = v5188
	var v5189 int32
	_ = v5189
	var v5194 int32
	_ = v5194
	var v5198 int32
	_ = v5198
	var v5203 int32
	_ = v5203
	var v5209 int32
	_ = v5209
	var v5248 int32
	_ = v5248
	var v5249 int32
	_ = v5249
	var v5250 int32
	_ = v5250
	var v5251 int32
	_ = v5251
	var v5256 int32
	_ = v5256
	var v5259 int32
	_ = v5259
	var v5269 int32
	_ = v5269
	var v5298 int32
	_ = v5298
	var v5299 int32
	_ = v5299
	var v5301 int32
	_ = v5301
	var v5302 int32
	_ = v5302
	var v5303 int32
	_ = v5303
	var v5304 int32
	_ = v5304
	var v5305 int32
	_ = v5305
	var v5306 int32
	_ = v5306
	var v5307 int32
	_ = v5307
	var v5308 int32
	_ = v5308
	var v5309 int64
	_ = v5309
	var v5310 int32
	_ = v5310
	var v5311 float64
	_ = v5311
	var v5313 int32
	_ = v5313
	var v5314 int32
	_ = v5314
	var v5326 float64
	_ = v5326
	var v5329 float64
	_ = v5329
	var v5331 int32
	_ = v5331
	var v5332 int32
	_ = v5332
	var v5351 int32
	_ = v5351
	var v5355 int32
	_ = v5355
	var v5356 int32
	_ = v5356
	var v5358 int32
	_ = v5358
	var v5359 int32
	_ = v5359
	var v5361 int32
	_ = v5361
	var v5367 int32
	_ = v5367
	var v5375 int32
	_ = v5375
	var v5406 int32
	_ = v5406
	var v5407 int32
	_ = v5407
	var v5408 int32
	_ = v5408
	var v5412 int32
	_ = v5412
	var v5413 int32
	_ = v5413
	var v5414 int32
	_ = v5414
	var v5417 int32
	_ = v5417
	var v5418 int32
	_ = v5418
	var v5419 int32
	_ = v5419
	var v5425 int32
	_ = v5425
	var v5467 int32
	_ = v5467
	var v5470 int32
	_ = v5470
	var v5474 int32
	_ = v5474
	var v5475 int32
	_ = v5475
	var v5479 int32
	_ = v5479
	var v5482 int32
	_ = v5482
	var v5483 int32
	_ = v5483
	var v5494 int32
	_ = v5494
	var v5529 int32
	_ = v5529
	var v5531 int32
	_ = v5531
	var v5532 int32
	_ = v5532
	var v5533 int32
	_ = v5533
	var v5534 int32
	_ = v5534
	var v5535 int32
	_ = v5535
	var v5536 int32
	_ = v5536
	var v5537 int32
	_ = v5537
	var v5541 int32
	_ = v5541
	var v5544 int32
	_ = v5544
	var v5548 int32
	_ = v5548
	var v5549 int32
	_ = v5549
	var v5550 int32
	_ = v5550
	var v5590 int32
	_ = v5590
	var v5594 int32
	_ = v5594
	var v5595 int32
	_ = v5595
	var v5596 int32
	_ = v5596
	var v5597 int32
	_ = v5597
	var v5598 int32
	_ = v5598
	var v5600 int32
	_ = v5600
	var v5602 int32
	_ = v5602
	var v5603 int32
	_ = v5603
	var v5609 int32
	_ = v5609
	var v5613 int32
	_ = v5613
	var v5614 int32
	_ = v5614
	var v5616 int32
	_ = v5616
	var v5617 int32
	_ = v5617
	var v5623 int32
	_ = v5623
	var v5663 int32
	_ = v5663
	var v5664 int32
	_ = v5664
	var v5665 int32
	_ = v5665
	var v5666 int32
	_ = v5666
	var v5667 int32
	_ = v5667
	var v5668 int32
	_ = v5668
	var v5669 int32
	_ = v5669
	var v5670 int32
	_ = v5670
	var v5671 int32
	_ = v5671
	var v5672 int64
	_ = v5672
	var v5673 int32
	_ = v5673
	var v5674 float64
	_ = v5674
	var v5676 int32
	_ = v5676
	var v5677 int32
	_ = v5677
	var v5689 float64
	_ = v5689
	var v5692 float64
	_ = v5692
	var v5694 int32
	_ = v5694
	var v5697 int32
	_ = v5697
	var v5713 int32
	_ = v5713
	var v5715 float64
	_ = v5715
	var v5717 float64
	_ = v5717
	var v5719 float64
	_ = v5719
	var v5721 int32
	_ = v5721
	var v5722 int32
	_ = v5722
	var v5724 int32
	_ = v5724
	var v5726 int32
	_ = v5726
	var v5728 int32
	_ = v5728
	var v5730 int32
	_ = v5730
	var v5731 int32
	_ = v5731
	var v5732 int32
	_ = v5732
	var v5733 int32
	_ = v5733
	var v5737 int32
	_ = v5737
	var v5740 int32
	_ = v5740
	var v5743 int32
	_ = v5743
	var v5745 int32
	_ = v5745
	var v5747 int32
	_ = v5747
	var v5785 int32
	_ = v5785
	var v5789 int32
	_ = v5789
	var v5790 int32
	_ = v5790
	var v5791 int32
	_ = v5791
	var v5792 int32
	_ = v5792
	var v5793 int32
	_ = v5793
	var v5795 int32
	_ = v5795
	var v5797 int32
	_ = v5797
	var v5798 int32
	_ = v5798
	var v5804 int32
	_ = v5804
	var v5808 int32
	_ = v5808
	var v5809 int32
	_ = v5809
	var v5811 int32
	_ = v5811
	var v5812 int32
	_ = v5812
	var v5820 int32
	_ = v5820
	var v5858 int32
	_ = v5858
	var v5859 int32
	_ = v5859
	var v5860 int32
	_ = v5860
	var v5861 int32
	_ = v5861
	var v5862 int32
	_ = v5862
	var v5863 int32
	_ = v5863
	var v5864 int32
	_ = v5864
	var v5865 int32
	_ = v5865
	var v5866 int32
	_ = v5866
	var v5867 int32
	_ = v5867
	var v5868 int32
	_ = v5868
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5871 int32
	_ = v5871
	var v5872 int32
	_ = v5872
	var v5873 int32
	_ = v5873
	var v5874 int32
	_ = v5874
	var v5875 int32
	_ = v5875
	var v5876 int64
	_ = v5876
	var v5877 float64
	_ = v5877
	var v5879 int32
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5892 float64
	_ = v5892
	var v5895 float64
	_ = v5895
	var v5897 int32
	_ = v5897
	var v5898 int32
	_ = v5898
	var v5915 int32
	_ = v5915
	var v5917 float64
	_ = v5917
	var v5919 float64
	_ = v5919
	var v5921 float64
	_ = v5921
	var v5923 int32
	_ = v5923
	var v5924 int32
	_ = v5924
	var v5926 int32
	_ = v5926
	var v5928 int32
	_ = v5928
	var v5930 int32
	_ = v5930
	var v5932 int32
	_ = v5932
	var v5933 int32
	_ = v5933
	var v5934 int32
	_ = v5934
	var v5935 int32
	_ = v5935
	var v5936 int32
	_ = v5936
	var v5940 int32
	_ = v5940
	var v5943 int32
	_ = v5943
	var v5946 int32
	_ = v5946
	var v5947 int32
	_ = v5947
	var v5949 int32
	_ = v5949
	var v5988 int32
	_ = v5988
	var v5992 int32
	_ = v5992
	var v5993 int32
	_ = v5993
	var v5994 int32
	_ = v5994
	var v5995 int32
	_ = v5995
	var v5996 int32
	_ = v5996
	var v5998 int32
	_ = v5998
	var v6000 int32
	_ = v6000
	var v6001 int32
	_ = v6001
	var v6007 int32
	_ = v6007
	var v6011 int32
	_ = v6011
	var v6012 int32
	_ = v6012
	var v6014 int32
	_ = v6014
	var v6015 int32
	_ = v6015
	var v6019 int32
	_ = v6019
	var v6061 int32
	_ = v6061
	var v6062 int32
	_ = v6062
	var v6063 int32
	_ = v6063
	var v6064 int32
	_ = v6064
	var v6065 int32
	_ = v6065
	var v6066 int32
	_ = v6066
	var v6067 int32
	_ = v6067
	var v6068 int32
	_ = v6068
	var v6069 int32
	_ = v6069
	var v6070 int32
	_ = v6070
	var v6071 int32
	_ = v6071
	var v6072 int32
	_ = v6072
	var v6073 int32
	_ = v6073
	var v6074 int32
	_ = v6074
	var v6075 int32
	_ = v6075
	var v6076 int32
	_ = v6076
	var v6078 int32
	_ = v6078
	var v6079 int32
	_ = v6079
	var v6091 int32
	_ = v6091
	var v6093 float64
	_ = v6093
	var v6095 float64
	_ = v6095
	var v6097 float64
	_ = v6097
	var v6099 int32
	_ = v6099
	var v6100 int32
	_ = v6100
	var v6102 int32
	_ = v6102
	var v6104 int32
	_ = v6104
	var v6106 int32
	_ = v6106
	var v6109 int32
	_ = v6109
	var v6110 int32
	_ = v6110
	var v6111 int32
	_ = v6111
	var v6112 int32
	_ = v6112
	var v6114 int32
	_ = v6114
	var v6115 int32
	_ = v6115
	var v6116 int32
	_ = v6116
	var v6119 int32
	_ = v6119
	var v6126 int32
	_ = v6126
	var v6127 int32
	_ = v6127
	var v6128 int32
	_ = v6128
	var v6129 int32
	_ = v6129
	var v6141 int32
	_ = v6141
	var v6142 int32
	_ = v6142
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6145 int32
	_ = v6145
	var v6146 int32
	_ = v6146
	var v6147 int32
	_ = v6147
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6153 int32
	_ = v6153
	var v6155 int32
	_ = v6155
	var v6166 int32
	_ = v6166
	var v6168 float64
	_ = v6168
	var v6170 float64
	_ = v6170
	var v6172 float64
	_ = v6172
	var v6174 int32
	_ = v6174
	var v6175 int32
	_ = v6175
	var v6177 int32
	_ = v6177
	var v6179 int32
	_ = v6179
	var v6181 int32
	_ = v6181
	var v6184 int32
	_ = v6184
	var v6185 int32
	_ = v6185
	var v6186 int32
	_ = v6186
	var v6188 int32
	_ = v6188
	var v6189 int32
	_ = v6189
	var v6190 int32
	_ = v6190
	var v6193 int32
	_ = v6193
	var v6200 int32
	_ = v6200
	var v6201 int32
	_ = v6201
	var v6202 int32
	_ = v6202
	var v6203 int32
	_ = v6203
	var v6215 int32
	_ = v6215
	var v6216 int32
	_ = v6216
	var v6217 int32
	_ = v6217
	var v6218 int32
	_ = v6218
	var v6219 int32
	_ = v6219
	var v6220 int32
	_ = v6220
	var v6221 int32
	_ = v6221
	var v6223 int32
	_ = v6223
	var v6224 int32
	_ = v6224
	var v6227 int32
	_ = v6227
	var v6229 int32
	_ = v6229
	var v6231 int32
	_ = v6231
	var v6237 int32
	_ = v6237
	var v6246 int32
	_ = v6246
	var v6248 float64
	_ = v6248
	var v6250 float64
	_ = v6250
	var v6252 float64
	_ = v6252
	var v6254 int32
	_ = v6254
	var v6255 int32
	_ = v6255
	var v6257 int32
	_ = v6257
	var v6259 int32
	_ = v6259
	var v6261 int32
	_ = v6261
	var v6262 int32
	_ = v6262
	var v6264 int32
	_ = v6264
	var v6265 int32
	_ = v6265
	var v6266 int32
	_ = v6266
	var v6267 int32
	_ = v6267
	var v6268 int32
	_ = v6268
	var v6271 int32
	_ = v6271
	var v6274 int32
	_ = v6274
	var v6277 int32
	_ = v6277
	var v6278 int32
	_ = v6278
	var v6280 int32
	_ = v6280
	var v6319 int32
	_ = v6319
	var v6323 int32
	_ = v6323
	var v6324 int32
	_ = v6324
	var v6325 int32
	_ = v6325
	var v6326 int32
	_ = v6326
	var v6327 int32
	_ = v6327
	var v6329 int32
	_ = v6329
	var v6331 int32
	_ = v6331
	var v6332 int32
	_ = v6332
	var v6338 int32
	_ = v6338
	var v6342 int32
	_ = v6342
	var v6343 int32
	_ = v6343
	var v6345 int32
	_ = v6345
	var v6346 int32
	_ = v6346
	var v6350 int32
	_ = v6350
	var v6392 int32
	_ = v6392
	var v6393 int32
	_ = v6393
	var v6394 int32
	_ = v6394
	var v6395 int32
	_ = v6395
	var v6397 int32
	_ = v6397
	var v6398 int32
	_ = v6398
	var v6399 int32
	_ = v6399
	var v6414 int32
	_ = v6414
	var v6416 float64
	_ = v6416
	var v6418 float64
	_ = v6418
	var v6420 float64
	_ = v6420
	var v6422 int32
	_ = v6422
	var v6423 int32
	_ = v6423
	var v6425 int32
	_ = v6425
	var v6427 int32
	_ = v6427
	var v6429 int32
	_ = v6429
	var v6430 int32
	_ = v6430
	var v6432 int32
	_ = v6432
	var v6435 int32
	_ = v6435
	var v6437 int32
	_ = v6437
	var v6439 int32
	_ = v6439
	var v6442 int32
	_ = v6442
	var v6443 int32
	_ = v6443
	var v6444 int32
	_ = v6444
	var v6445 int32
	_ = v6445
	var v6447 int32
	_ = v6447
	var v6448 int32
	_ = v6448
	var v6451 int32
	_ = v6451
	var v6452 int32
	_ = v6452
	var v6460 int32
	_ = v6460
	var v6461 int32
	_ = v6461
	var v6463 int32
	_ = v6463
	var v6464 int32
	_ = v6464
	var v6465 int32
	_ = v6465
	var v6466 int32
	_ = v6466
	var v6467 int32
	_ = v6467
	var v6470 int32
	_ = v6470
	var v6473 int32
	_ = v6473
	var v6476 int32
	_ = v6476
	var v6482 int32
	_ = v6482
	var v6523 int32
	_ = v6523
	var v6524 int32
	_ = v6524
	var v6526 int32
	_ = v6526
	var v6527 int32
	_ = v6527
	var v6528 int32
	_ = v6528
	var v6531 int32
	_ = v6531
	var v6535 int32
	_ = v6535
	var v6539 int32
	_ = v6539
	var v6544 int32
	_ = v6544
	var v6545 int32
	_ = v6545
	var v6548 int32
	_ = v6548
	var v6549 int32
	_ = v6549
	var v6554 int32
	_ = v6554
	var v6596 int32
	_ = v6596
	var v6600 int32
	_ = v6600
	var v6601 int32
	_ = v6601
	var v6603 int32
	_ = v6603
	var v6604 int32
	_ = v6604
	var v6606 int32
	_ = v6606
	var v6607 int32
	_ = v6607
	var v6609 int32
	_ = v6609
	var v6610 int32
	_ = v6610
	var v6611 int32
	_ = v6611
	var v6614 int32
	_ = v6614
	var v6615 int32
	_ = v6615
	var v6616 int32
	_ = v6616
	var v6624 int32
	_ = v6624
	var v6629 int32
	_ = v6629
	var v6661 int32
	_ = v6661
	var v6662 int32
	_ = v6662
	var v6664 int32
	_ = v6664
	var v6665 int32
	_ = v6665
	var v6668 int32
	_ = v6668
	var v6671 int32
	_ = v6671
	var v6676 int32
	_ = v6676
	var v6679 int32
	_ = v6679
	var v6680 int32
	_ = v6680
	var v6730 int32
	_ = v6730
	var v6732 float64
	_ = v6732
	var v6734 float64
	_ = v6734
	var v6736 float64
	_ = v6736
	var v6738 int32
	_ = v6738
	var v6739 int32
	_ = v6739
	var v6741 int32
	_ = v6741
	var v6743 int32
	_ = v6743
	var v6751 int32
	_ = v6751
	var v6752 int32
	_ = v6752
	var v6760 int32
	_ = v6760
	var v6765 int32
	_ = v6765
	var v6813 int32
	_ = v6813
	var v6817 int32
	_ = v6817
	var v6822 int32
	_ = v6822
	var v6823 int32
	_ = v6823
	var v6825 int32
	_ = v6825
	var v6827 int32
	_ = v6827
	var v6828 int32
	_ = v6828
	var v6829 int32
	_ = v6829
	var v6830 int32
	_ = v6830
	var v6833 int32
	_ = v6833
	var v6834 int32
	_ = v6834
	var v6835 int32
	_ = v6835
	var v6836 int32
	_ = v6836
	var v6839 int32
	_ = v6839
	var v6840 int32
	_ = v6840
	var v6843 int32
	_ = v6843
	var v6846 int32
	_ = v6846
	var v6849 int32
	_ = v6849
	var v6853 int32
	_ = v6853
	var v6888 int32
	_ = v6888
	var v6892 int32
	_ = v6892
	var v6893 int32
	_ = v6893
	var v6894 int32
	_ = v6894
	var v6895 int32
	_ = v6895
	var v6896 int32
	_ = v6896
	var v6898 int32
	_ = v6898
	var v6900 int32
	_ = v6900
	var v6901 int32
	_ = v6901
	var v6907 int32
	_ = v6907
	var v6911 int32
	_ = v6911
	var v6912 int32
	_ = v6912
	var v6914 int32
	_ = v6914
	var v6915 int32
	_ = v6915
	var v6919 int32
	_ = v6919
	var v6973 int32
	_ = v6973
	var v7012 int32
	_ = v7012
	var v7013 int32
	_ = v7013
	var v7014 int32
	_ = v7014
	var v7020 int32
	_ = v7020
	var v7021 int32
	_ = v7021
	var v7023 int32
	_ = v7023
	var v7024 int32
	_ = v7024
	var v7062 int32
	_ = v7062
	var v7066 int32
	_ = v7066
	var v7067 int32
	_ = v7067
	var v7068 int32
	_ = v7068
	var v7073 int32
	_ = v7073
	var v7075 int32
	_ = v7075
	var v7076 int32
	_ = v7076
	var v7077 int32
	_ = v7077
	var v7078 int32
	_ = v7078
	var v7081 int32
	_ = v7081
	var v7082 int32
	_ = v7082
	var v7083 int32
	_ = v7083
	var v7085 int32
	_ = v7085
	var v7086 int32
	_ = v7086
	var v7093 int32
	_ = v7093
	var v7094 int32
	_ = v7094
	var v7139 int32
	_ = v7139
	var v7178 int32
	_ = v7178
	var v7186 int32
	_ = v7186
	var v7225 int32
	_ = v7225
	var v7226 int32
	_ = v7226
	var v7227 int32
	_ = v7227
	var v7228 int32
	_ = v7228
	var v7272 int32
	_ = v7272
	var v7273 int32
	_ = v7273
	var v7275 int32
	_ = v7275
	var v7278 int32
	_ = v7278
	var v7279 int32
	_ = v7279
	var v7281 int32
	_ = v7281
	var v7282 int32
	_ = v7282
	var v7283 int32
	_ = v7283
	var v7286 int32
	_ = v7286
	var v7292 int32
	_ = v7292
	var v7335 int32
	_ = v7335
	var v7336 int32
	_ = v7336
	var v7338 int32
	_ = v7338
	var v7339 int32
	_ = v7339
	var v7340 int32
	_ = v7340
	var v7346 int32
	_ = v7346
	var v7349 int32
	_ = v7349
	var v7350 int32
	_ = v7350
	var v7351 int32
	_ = v7351
	var v7354 int32
	_ = v7354
	var v7355 int32
	_ = v7355
	var v7401 int32
	_ = v7401
	var v7404 int32
	_ = v7404
	var v7408 int32
	_ = v7408
	var v7412 int32
	_ = v7412
	var v7413 int32
	_ = v7413
	var v7416 int32
	_ = v7416
	var v7422 int32
	_ = v7422
	var v7465 int32
	_ = v7465
	var v7466 int32
	_ = v7466
	var v7468 int32
	_ = v7468
	var v7471 int32
	_ = v7471
	var v7472 int32
	_ = v7472
	var v7476 int32
	_ = v7476
	var v7479 int32
	_ = v7479
	var v7480 int32
	_ = v7480
	var v7526 int32
	_ = v7526
	var v7527 int32
	_ = v7527
	var v7528 int32
	_ = v7528
	var v7532 int32
	_ = v7532
	var v7535 int32
	_ = v7535
	var v7539 int32
	_ = v7539
	var v7540 int32
	_ = v7540
	var v7542 int32
	_ = v7542
	var v7581 int32
	_ = v7581
	var v7585 int32
	_ = v7585
	var v7586 int32
	_ = v7586
	var v7587 int32
	_ = v7587
	var v7588 int32
	_ = v7588
	var v7589 int32
	_ = v7589
	var v7591 int32
	_ = v7591
	var v7593 int32
	_ = v7593
	var v7594 int32
	_ = v7594
	var v7600 int32
	_ = v7600
	var v7604 int32
	_ = v7604
	var v7605 int32
	_ = v7605
	var v7607 int32
	_ = v7607
	var v7608 int32
	_ = v7608
	var v7613 int32
	_ = v7613
	var v7654 float64
	_ = v7654
	var v7656 int32
	_ = v7656
	var v7657 int32
	_ = v7657
	var v7669 float64
	_ = v7669
	var v7672 float64
	_ = v7672
	var v7674 int32
	_ = v7674
	var v7682 int64
	_ = v7682
	var v7684 int32
	_ = v7684
	var v7698 int32
	_ = v7698
	var v7699 int32
	_ = v7699
	var v7738 int32
	_ = v7738
	var v7742 int32
	_ = v7742
	var v7743 int32
	_ = v7743
	var v7744 int32
	_ = v7744
	var v7746 int32
	_ = v7746
	var v7747 int32
	_ = v7747
	var v7750 int32
	_ = v7750
	var v7754 int32
	_ = v7754
	var v7758 int32
	_ = v7758
	var v7761 int32
	_ = v7761
	var v7764 int32
	_ = v7764
	var v7765 int32
	_ = v7765
	var v7769 int32
	_ = v7769
	var v7777 int32
	_ = v7777
	var v7778 int32
	_ = v7778
	var v7781 int32
	_ = v7781
	var v7792 int32
	_ = v7792
	var v7795 int32
	_ = v7795
	var v7796 int32
	_ = v7796
	var v7799 int32
	_ = v7799
	var v7800 int32
	_ = v7800
	var v7809 int32
	_ = v7809
	var v7816 int32
	_ = v7816
	var v7819 int32
	_ = v7819
	var v7822 int32
	_ = v7822
	var v7824 int32
	_ = v7824
	var v7825 int32
	_ = v7825
	var v7830 int32
	_ = v7830
	var v7834 int32
	_ = v7834
	var v7835 int32
	_ = v7835
	var v7836 int32
	_ = v7836
	var v7845 int32
	_ = v7845
	var v7846 int32
	_ = v7846
	var v7847 int32
	_ = v7847
	var v7848 int32
	_ = v7848
	var v7849 int32
	_ = v7849
	var v7850 int32
	_ = v7850
	var v7851 int32
	_ = v7851
	var v7852 int32
	_ = v7852
	var v7853 int32
	_ = v7853
	var v7855 int32
	_ = v7855
	var v7857 int32
	_ = v7857
	var v7859 int32
	_ = v7859
	var v7861 int32
	_ = v7861
	var v7862 int32
	_ = v7862
	var v7863 int32
	_ = v7863
	var v7865 int32
	_ = v7865
	var v7871 int32
	_ = v7871
	var v7873 int32
	_ = v7873
	var v7880 int32
	_ = v7880
	var v7882 int32
	_ = v7882
	var v7884 int32
	_ = v7884
	var v7886 int32
	_ = v7886
	var v7896 int32
	_ = v7896
	var v7897 int32
	_ = v7897
	var v7899 int32
	_ = v7899
	var v7900 int32
	_ = v7900
	var v7903 int32
	_ = v7903
	var v7907 int32
	_ = v7907
	var v7929 int32
	_ = v7929
	var v7933 int32
	_ = v7933
	var v7942 int32
	_ = v7942
	var v7949 int32
	_ = v7949
	var v7950 int32
	_ = v7950
	var v7952 int32
	_ = v7952
	var v7953 int32
	_ = v7953
	var v7958 int32
	_ = v7958
	var v7964 int32
	_ = v7964
	var v7969 int32
	_ = v7969
	var v7973 int32
	_ = v7973
	var v7979 int32
	_ = v7979
	var v7984 int32
	_ = v7984
	var v7988 int32
	_ = v7988
	var v7992 int32
	_ = v7992
	var v7997 int32
	_ = v7997
	var v8001 int32
	_ = v8001
	var v8005 int32
	_ = v8005
	var v8010 int32
	_ = v8010
	var v8015 int32
	_ = v8015
	var v8055 int32
	_ = v8055
	var v8057 int32
	_ = v8057
	var v8058 int32
	_ = v8058
	var v8059 int32
	_ = v8059
	var v8062 int32
	_ = v8062
	var v8063 int32
	_ = v8063
	var v8065 int32
	_ = v8065
	var v8066 int32
	_ = v8066
	var v8067 int32
	_ = v8067
	var v8068 int32
	_ = v8068
	var v8069 int32
	_ = v8069
	var v8070 int32
	_ = v8070
	var v8071 int32
	_ = v8071
	var v8074 int32
	_ = v8074
	var v8083 int32
	_ = v8083
	var v8125 int32
	_ = v8125
	var v8126 int32
	_ = v8126
	var v8128 int32
	_ = v8128
	var v8129 int32
	_ = v8129
	var v8130 int32
	_ = v8130
	var v8131 int32
	_ = v8131
	var v8134 int32
	_ = v8134
	var v8137 int32
	_ = v8137
	var v8138 int32
	_ = v8138
	var v8139 int32
	_ = v8139
	var v8142 int32
	_ = v8142
	var v8145 int32
	_ = v8145
	var v8146 int32
	_ = v8146
	var v8154 int32
	_ = v8154
	var v8193 int32
	_ = v8193
	var v8194 int32
	_ = v8194
	var v8197 int32
	_ = v8197
	var v8199 int32
	_ = v8199
	var v8200 int32
	_ = v8200
	var v8201 int32
	_ = v8201
	var v8207 int32
	_ = v8207
	var v8212 int32
	_ = v8212
	var v8214 int32
	_ = v8214
	var v8217 int32
	_ = v8217
	var v8220 float64
	_ = v8220
	var v8221 float64
	_ = v8221
	var v8222 int32
	_ = v8222
	var v8225 int32
	_ = v8225
	var v8228 int32
	_ = v8228
	var v8229 int32
	_ = v8229
	var v8230 int32
	_ = v8230
	var v8235 float64
	_ = v8235
	var v8238 int32
	_ = v8238
	var v8239 float64
	_ = v8239
	var v8245 float64
	_ = v8245
	var v8251 float64
	_ = v8251
	var v8253 float64
	_ = v8253
	var v8255 float64
	_ = v8255
	var v8257 int32
	_ = v8257
	var v8258 int32
	_ = v8258
	var v8261 int32
	_ = v8261
	var v8264 int32
	_ = v8264
	var v8265 int32
	_ = v8265
	var v8268 int32
	_ = v8268
	var v8269 int32
	_ = v8269
	var v8270 int32
	_ = v8270
	var v8271 int32
	_ = v8271
	var v8280 int32
	_ = v8280
	var v8281 int32
	_ = v8281
	var v8283 int32
	_ = v8283
	var v8284 int32
	_ = v8284
	var v8285 int32
	_ = v8285
	var v8286 int32
	_ = v8286
	var v8287 int32
	_ = v8287
	var v8290 int32
	_ = v8290
	var v8298 int32
	_ = v8298
	var v8341 int32
	_ = v8341
	var v8342 int32
	_ = v8342
	var v8344 int32
	_ = v8344
	var v8345 int32
	_ = v8345
	var v8346 int32
	_ = v8346
	var v8347 int32
	_ = v8347
	var v8348 int32
	_ = v8348
	var v8351 int32
	_ = v8351
	var v8354 int32
	_ = v8354
	var v8355 int32
	_ = v8355
	var v8356 int32
	_ = v8356
	var v8359 int32
	_ = v8359
	var v8360 int32
	_ = v8360
	var v8412 int32
	_ = v8412
	var v8454 int32
	_ = v8454
	var v8456 float64
	_ = v8456
	var v8458 float64
	_ = v8458
	var v8460 float64
	_ = v8460
	var v8462 int32
	_ = v8462
	var v8463 int32
	_ = v8463
	var v8465 int32
	_ = v8465
	var v8467 int32
	_ = v8467
	var v8471 int32
	_ = v8471
	var v8516 int32
	_ = v8516
	var v8519 int32
	_ = v8519
	var v8520 int32
	_ = v8520
	var v8521 int32
	_ = v8521
	var v8522 int32
	_ = v8522
	var v8523 int32
	_ = v8523
	var v8524 int32
	_ = v8524
	var v8527 int32
	_ = v8527
	var v8528 int32
	_ = v8528
	var v8529 int32
	_ = v8529
	var v8530 int32
	_ = v8530
	var v8531 int32
	_ = v8531
	var v8532 int32
	_ = v8532
	var v8536 int32
	_ = v8536
	var v8577 int32
	_ = v8577
	var v8580 int32
	_ = v8580
	var v8582 int32
	_ = v8582
	var v8586 int32
	_ = v8586
	var v8591 int32
	_ = v8591
	var v8594 int32
	_ = v8594
	var v8597 int32
	_ = v8597
	var v8599 int32
	_ = v8599
	var v8602 int32
	_ = v8602
	var v8605 int32
	_ = v8605
	var v8606 int32
	_ = v8606
	var v8611 int32
	_ = v8611
	var v8613 int32
	_ = v8613
	var v8615 int32
	_ = v8615
	var v8620 int32
	_ = v8620
	var v8624 int32
	_ = v8624
	var v8625 int32
	_ = v8625
	var v8626 int32
	_ = v8626
	var v8630 int32
	_ = v8630
	var v8631 int32
	_ = v8631
	var v8632 int32
	_ = v8632
	var v8633 int32
	_ = v8633
	var v8637 int32
	_ = v8637
	var v8638 int32
	_ = v8638
	var v8639 int32
	_ = v8639
	var v8641 int32
	_ = v8641
	var v8642 int32
	_ = v8642
	var v8645 int32
	_ = v8645
	var v8646 int32
	_ = v8646
	var v8653 int32
	_ = v8653
	var v8654 int32
	_ = v8654
	var v8663 int32
	_ = v8663
	var v8665 float64
	_ = v8665
	var v8667 float64
	_ = v8667
	var v8669 float64
	_ = v8669
	var v8671 int32
	_ = v8671
	var v8672 int32
	_ = v8672
	var v8674 int32
	_ = v8674
	var v8676 int32
	_ = v8676
	var v8678 int32
	_ = v8678
	var v8681 int32
	_ = v8681
	var v8682 int32
	_ = v8682
	var v8684 int32
	_ = v8684
	var v8685 int32
	_ = v8685
	var v8688 int32
	_ = v8688
	var v8689 int32
	_ = v8689
	var v8695 int32
	_ = v8695
	var v8697 float64
	_ = v8697
	var v8699 float64
	_ = v8699
	var v8701 float64
	_ = v8701
	var v8703 int32
	_ = v8703
	var v8704 int32
	_ = v8704
	var v8706 int32
	_ = v8706
	var v8708 int32
	_ = v8708
	var v8710 int32
	_ = v8710
	var v8712 int32
	_ = v8712
	var v8713 int32
	_ = v8713
	var v8714 int32
	_ = v8714
	var v8715 int32
	_ = v8715
	var v8716 int32
	_ = v8716
	var v8720 int32
	_ = v8720
	var v8723 int32
	_ = v8723
	var v8726 int32
	_ = v8726
	var v8727 int32
	_ = v8727
	var v8729 int32
	_ = v8729
	var v8768 int32
	_ = v8768
	var v8772 int32
	_ = v8772
	var v8773 int32
	_ = v8773
	var v8774 int32
	_ = v8774
	var v8775 int32
	_ = v8775
	var v8776 int32
	_ = v8776
	var v8778 int32
	_ = v8778
	var v8780 int32
	_ = v8780
	var v8781 int32
	_ = v8781
	var v8787 int32
	_ = v8787
	var v8791 int32
	_ = v8791
	var v8792 int32
	_ = v8792
	var v8794 int32
	_ = v8794
	var v8795 int32
	_ = v8795
	var v8799 int32
	_ = v8799
	var v8842 int32
	_ = v8842
	var v8843 int32
	_ = v8843
	var v8844 int32
	_ = v8844
	var v8852 int32
	_ = v8852
	var v8854 float64
	_ = v8854
	var v8856 float64
	_ = v8856
	var v8858 float64
	_ = v8858
	var v8860 int32
	_ = v8860
	var v8861 int32
	_ = v8861
	var v8863 int32
	_ = v8863
	var v8865 int32
	_ = v8865
	var v8867 int32
	_ = v8867
	var v8870 int32
	_ = v8870
	var v8871 int32
	_ = v8871
	var v8872 int32
	_ = v8872
	var v8874 int32
	_ = v8874
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
	var v8885 int32
	_ = v8885
	var v8888 int32
	_ = v8888
	var v8893 int32
	_ = v8893
	var v8896 int32
	_ = v8896
	var v8900 int32
	_ = v8900
	var v8902 int32
	_ = v8902
	var v8904 int32
	_ = v8904
	var v8905 int32
	_ = v8905
	var v8906 int32
	_ = v8906
	var v8907 int32
	_ = v8907
	var v8911 int32
	_ = v8911
	var v8914 int32
	_ = v8914
	var v8917 int32
	_ = v8917
	var v8920 int32
	_ = v8920
	var v8921 int32
	_ = v8921
	var v8959 int32
	_ = v8959
	var v8963 int32
	_ = v8963
	var v8964 int32
	_ = v8964
	var v8965 int32
	_ = v8965
	var v8966 int32
	_ = v8966
	var v8967 int32
	_ = v8967
	var v8969 int32
	_ = v8969
	var v8971 int32
	_ = v8971
	var v8972 int32
	_ = v8972
	var v8978 int32
	_ = v8978
	var v8982 int32
	_ = v8982
	var v8983 int32
	_ = v8983
	var v8985 int32
	_ = v8985
	var v8986 int32
	_ = v8986
	var v8989 int32
	_ = v8989
	var v8990 int32
	_ = v8990
	var v8991 int32
	_ = v8991
	var v8992 int32
	_ = v8992
	var v8996 int32
	_ = v8996
	var v8999 int32
	_ = v8999
	var v9002 int32
	_ = v9002
	var v9005 int32
	_ = v9005
	var v9006 int32
	_ = v9006
	var v9044 int32
	_ = v9044
	var v9048 int32
	_ = v9048
	var v9049 int32
	_ = v9049
	var v9050 int32
	_ = v9050
	var v9051 int32
	_ = v9051
	var v9052 int32
	_ = v9052
	var v9054 int32
	_ = v9054
	var v9056 int32
	_ = v9056
	var v9057 int32
	_ = v9057
	var v9063 int32
	_ = v9063
	var v9067 int32
	_ = v9067
	var v9068 int32
	_ = v9068
	var v9070 int32
	_ = v9070
	var v9071 int32
	_ = v9071
	var v9078 int32
	_ = v9078
	var v9117 int32
	_ = v9117
	var v9118 int32
	_ = v9118
	var v9119 int32
	_ = v9119
	var v9122 int32
	_ = v9122
	var v9127 int32
	_ = v9127
	var v9167 float64
	_ = v9167
	var v9169 float64
	_ = v9169
	var v9171 float64
	_ = v9171
	var v9173 int32
	_ = v9173
	var v9174 int32
	_ = v9174
	var v9177 int32
	_ = v9177
	var v9178 int32
	_ = v9178
	var v9179 int32
	_ = v9179
	var v9189 int32
	_ = v9189
	var v9191 float64
	_ = v9191
	var v9193 float64
	_ = v9193
	var v9195 float64
	_ = v9195
	var v9197 int32
	_ = v9197
	var v9198 int32
	_ = v9198
	var v9200 int32
	_ = v9200
	var v9202 int32
	_ = v9202
	var v9246 int32
	_ = v9246
	var v9248 int32
	_ = v9248
	var v9251 int32
	_ = v9251
	var v9258 int32
	_ = v9258
	var v9298 int32
	_ = v9298
	var v9302 int32
	_ = v9302
	var v9303 int32
	_ = v9303
	var v9304 int32
	_ = v9304
	var v9305 int32
	_ = v9305
	var v9306 int32
	_ = v9306
	var v9307 int32
	_ = v9307
	var v9308 int64
	_ = v9308
	var v9309 int32
	_ = v9309
	var v9311 int32
	_ = v9311
	var v9312 int32
	_ = v9312
	var v9315 int32
	_ = v9315
	var v9316 int64
	_ = v9316
	var v9320 int32
	_ = v9320
	var v9328 int32
	_ = v9328
	var v9329 int32
	_ = v9329
	var v9331 int32
	_ = v9331
	var v9332 float64
	_ = v9332
	var v9334 float64
	_ = v9334
	var v9338 int32
	_ = v9338
	var v9339 int32
	_ = v9339
	var v9340 int32
	_ = v9340
	var v9344 int32
	_ = v9344
	var v9345 int32
	_ = v9345
	var v9347 int32
	_ = v9347
	var v9349 int32
	_ = v9349
	var v9351 int32
	_ = v9351
	var v9353 int32
	_ = v9353
	var v9354 int32
	_ = v9354
	var v9355 int32
	_ = v9355
	var v9356 int32
	_ = v9356
	var v9357 int32
	_ = v9357
	var v9359 int32
	_ = v9359
	var v9360 int32
	_ = v9360
	var v9362 int32
	_ = v9362
	var v9363 int32
	_ = v9363
	var v9364 int32
	_ = v9364
	var v9366 int32
	_ = v9366
	var v9367 int32
	_ = v9367
	var v9368 int32
	_ = v9368
	var v9369 int32
	_ = v9369
	var v9370 int32
	_ = v9370
	var v9373 int32
	_ = v9373
	var v9374 int32
	_ = v9374
	var v9377 int32
	_ = v9377
	var v9378 int32
	_ = v9378
	var v9379 int32
	_ = v9379
	var v9380 int32
	_ = v9380
	var v9386 int32
	_ = v9386
	var v9387 int32
	_ = v9387
	var v9389 int32
	_ = v9389
	var v9392 int32
	_ = v9392
	var v9393 int32
	_ = v9393
	var v9394 int32
	_ = v9394
	var v9395 int32
	_ = v9395
	var v9396 int32
	_ = v9396
	var v9397 int32
	_ = v9397
	var v9399 int32
	_ = v9399
	var v9400 int32
	_ = v9400
	var v9401 int32
	_ = v9401
	var v9403 int32
	_ = v9403
	var v9404 int32
	_ = v9404
	var v9405 int32
	_ = v9405
	var v9411 int32
	_ = v9411
	var v9413 int32
	_ = v9413
	var v9415 int32
	_ = v9415
	var v9421 int32
	_ = v9421
	var v9422 int32
	_ = v9422
	var v9424 int32
	_ = v9424
	var v9425 int32
	_ = v9425
	var v9426 int32
	_ = v9426
	var v9429 int32
	_ = v9429
	var v9434 int32
	_ = v9434
	var v9435 int32
	_ = v9435
	var v9481 int32
	_ = v9481
	var v9482 int32
	_ = v9482
	var v9483 int32
	_ = v9483
	var v9487 int32
	_ = v9487
	var v9490 int32
	_ = v9490
	var v9494 int32
	_ = v9494
	var v9496 int32
	_ = v9496
	var v9497 int32
	_ = v9497
	var v9536 int32
	_ = v9536
	var v9540 int32
	_ = v9540
	var v9541 int32
	_ = v9541
	var v9542 int32
	_ = v9542
	var v9543 int32
	_ = v9543
	var v9544 int32
	_ = v9544
	var v9546 int32
	_ = v9546
	var v9548 int32
	_ = v9548
	var v9549 int32
	_ = v9549
	var v9555 int32
	_ = v9555
	var v9559 int32
	_ = v9559
	var v9560 int32
	_ = v9560
	var v9562 int32
	_ = v9562
	var v9563 int32
	_ = v9563
	var v9570 int32
	_ = v9570
	var v9609 int32
	_ = v9609
	var v9611 int32
	_ = v9611
	var v9612 int32
	_ = v9612
	var v9621 int32
	_ = v9621
	var v9623 float64
	_ = v9623
	var v9625 float64
	_ = v9625
	var v9627 float64
	_ = v9627
	var v9629 int32
	_ = v9629
	var v9630 int32
	_ = v9630
	var v9632 int32
	_ = v9632
	var v9634 int32
	_ = v9634
	var v9636 int32
	_ = v9636
	var v9638 int32
	_ = v9638
	var v9639 int32
	_ = v9639
	var v9643 int32
	_ = v9643
	var v9646 int32
	_ = v9646
	var v9649 int32
	_ = v9649
	var v9650 int32
	_ = v9650
	var v9651 int32
	_ = v9651
	var v9691 int32
	_ = v9691
	var v9695 int32
	_ = v9695
	var v9696 int32
	_ = v9696
	var v9697 int32
	_ = v9697
	var v9698 int32
	_ = v9698
	var v9699 int32
	_ = v9699
	var v9701 int32
	_ = v9701
	var v9703 int32
	_ = v9703
	var v9704 int32
	_ = v9704
	var v9710 int32
	_ = v9710
	var v9714 int32
	_ = v9714
	var v9715 int32
	_ = v9715
	var v9717 int32
	_ = v9717
	var v9718 int32
	_ = v9718
	var v9723 int32
	_ = v9723
	var v9764 int32
	_ = v9764
	var v9765 int32
	_ = v9765
	var v9766 int32
	_ = v9766
	var v9768 int32
	_ = v9768
	var v9769 int32
	_ = v9769
	var v9778 int32
	_ = v9778
	var v9780 float64
	_ = v9780
	var v9782 float64
	_ = v9782
	var v9784 float64
	_ = v9784
	var v9786 int32
	_ = v9786
	var v9787 int32
	_ = v9787
	var v9789 int32
	_ = v9789
	var v9791 int32
	_ = v9791
	var v9793 int32
	_ = v9793
	var v9794 int32
	_ = v9794
	var v9796 int32
	_ = v9796
	var v9797 int32
	_ = v9797
	var v9800 int32
	_ = v9800
	var v9801 int32
	_ = v9801
	var v9805 int32
	_ = v9805
	var v9806 int32
	_ = v9806
	var v9809 int32
	_ = v9809
	var v9814 int32
	_ = v9814
	var v9818 int32
	_ = v9818
	var v9821 int32
	_ = v9821
	var v9854 int32
	_ = v9854
	var v9858 int32
	_ = v9858
	var v9859 int32
	_ = v9859
	var v9860 int32
	_ = v9860
	var v9861 int32
	_ = v9861
	var v9862 int32
	_ = v9862
	var v9864 int32
	_ = v9864
	var v9866 int32
	_ = v9866
	var v9867 int32
	_ = v9867
	var v9873 int32
	_ = v9873
	var v9877 int32
	_ = v9877
	var v9878 int32
	_ = v9878
	var v9880 int32
	_ = v9880
	var v9881 int32
	_ = v9881
	var v9885 int32
	_ = v9885
	var v9930 int32
	_ = v9930
	var v9943 int32
	_ = v9943
	var v9976 int32
	_ = v9976
	var v9977 int32
	_ = v9977
	var v9978 int32
	_ = v9978
	var v9979 int32
	_ = v9979
	var v9981 float64
	_ = v9981
	var v9983 float64
	_ = v9983
	var v9985 float64
	_ = v9985
	var v9987 int32
	_ = v9987
	var v9988 int32
	_ = v9988
	var v9990 int32
	_ = v9990
	var v9992 int32
	_ = v9992
	var v9993 int32
	_ = v9993
	var v9999 int32
	_ = v9999
	var v10001 int32
	_ = v10001
	var v10002 int32
	_ = v10002
	var v10008 int32
	_ = v10008
	var v10015 int32
	_ = v10015
	var v10016 int32
	_ = v10016
	var v10017 int32
	_ = v10017
	var v10018 int32
	_ = v10018
	var v10019 int32
	_ = v10019
	var v10020 int32
	_ = v10020
	var v10021 int32
	_ = v10021
	var v10024 int32
	_ = v10024
	var v10036 int32
	_ = v10036
	var v10039 int32
	_ = v10039
	var v10072 int32
	_ = v10072
	var v10076 int32
	_ = v10076
	var v10078 int32
	_ = v10078
	var v10079 int32
	_ = v10079
	var v10080 int32
	_ = v10080
	var v10081 int32
	_ = v10081
	var v10082 int32
	_ = v10082
	var v10094 int32
	_ = v10094
	var v10095 int32
	_ = v10095
	var v10096 int32
	_ = v10096
	var v10097 int32
	_ = v10097
	var v10098 int32
	_ = v10098
	var v10100 int32
	_ = v10100
	var v10108 int32
	_ = v10108
	var v10109 int32
	_ = v10109
	var v10110 int32
	_ = v10110
	var v10113 int32
	_ = v10113
	var v10114 int32
	_ = v10114
	var v10116 int32
	_ = v10116
	var v10117 int32
	_ = v10117
	var v10119 int32
	_ = v10119
	var v10121 int32
	_ = v10121
	var v10124 int32
	_ = v10124
	var v10125 int32
	_ = v10125
	var v10126 int32
	_ = v10126
	var v10131 int32
	_ = v10131
	var v10132 int32
	_ = v10132
	var v10133 int32
	_ = v10133
	var v10136 int32
	_ = v10136
	var v10137 int32
	_ = v10137
	var v10138 int32
	_ = v10138
	var v10141 int32
	_ = v10141
	var v10142 int32
	_ = v10142
	var v10144 int32
	_ = v10144
	var v10149 int32
	_ = v10149
	var v10162 int32
	_ = v10162
	var v10163 int32
	_ = v10163
	var v10172 int32
	_ = v10172
	var v10176 int32
	_ = v10176
	var v10180 int32
	_ = v10180
	var v10182 int32
	_ = v10182
	var v10186 int32
	_ = v10186
	var v10187 int32
	_ = v10187
	var v10192 int32
	_ = v10192
	var v10195 int32
	_ = v10195
	var v10202 int32
	_ = v10202
	var v10204 int32
	_ = v10204
	var v10208 int32
	_ = v10208
	var v10216 int32
	_ = v10216
	var v10217 int32
	_ = v10217
	var v10218 int32
	_ = v10218
	var v10219 int32
	_ = v10219
	var v10221 int32
	_ = v10221
	var v10222 int32
	_ = v10222
	var v10225 int32
	_ = v10225
	var v10227 int32
	_ = v10227
	var v10228 int32
	_ = v10228
	var v10229 int32
	_ = v10229
	var v10235 int32
	_ = v10235
	var v10240 int32
	_ = v10240
	var v10242 int32
	_ = v10242
	var v10245 int32
	_ = v10245
	var v10246 float64
	_ = v10246
	var v10247 float64
	_ = v10247
	var v10248 int32
	_ = v10248
	var v10251 int32
	_ = v10251
	var v10252 float64
	_ = v10252
	var v10254 int32
	_ = v10254
	var v10255 int32
	_ = v10255
	var v10256 int32
	_ = v10256
	var v10261 float64
	_ = v10261
	var v10264 int32
	_ = v10264
	var v10265 float64
	_ = v10265
	var v10271 float64
	_ = v10271
	var v10277 float64
	_ = v10277
	var v10279 float64
	_ = v10279
	var v10281 float64
	_ = v10281
	var v10283 int32
	_ = v10283
	var v10284 int32
	_ = v10284
	var v10287 int32
	_ = v10287
	var v10289 int32
	_ = v10289
	var v10296 int32
	_ = v10296
	var v10297 int32
	_ = v10297
	var v10299 int32
	_ = v10299
	var v10300 int32
	_ = v10300
	var v10313 int32
	_ = v10313
	var v10349 int32
	_ = v10349
	var v10352 int32
	_ = v10352
	var v10354 int32
	_ = v10354
	var v10355 int32
	_ = v10355
	var v10358 int32
	_ = v10358
	var v10359 int32
	_ = v10359
	var v10360 int32
	_ = v10360
	var v10370 int32
	_ = v10370
	var v10371 int32
	_ = v10371
	var v10372 int32
	_ = v10372
	var v10373 int32
	_ = v10373
	var v10374 int32
	_ = v10374
	var v10375 int32
	_ = v10375
	var v10379 int32
	_ = v10379
	var v10383 int32
	_ = v10383
	var v10388 int32
	_ = v10388
	var v10389 int32
	_ = v10389
	var v10390 int32
	_ = v10390
	var v10394 int32
	_ = v10394
	var v10395 int32
	_ = v10395
	var v10398 int32
	_ = v10398
	var v10402 int32
	_ = v10402
	var v10404 int32
	_ = v10404
	var v10407 int32
	_ = v10407
	var v10443 int32
	_ = v10443
	var v10447 int32
	_ = v10447
	var v10448 int32
	_ = v10448
	var v10449 int32
	_ = v10449
	var v10450 int32
	_ = v10450
	var v10451 int32
	_ = v10451
	var v10453 int32
	_ = v10453
	var v10455 int32
	_ = v10455
	var v10456 int32
	_ = v10456
	var v10462 int32
	_ = v10462
	var v10466 int32
	_ = v10466
	var v10467 int32
	_ = v10467
	var v10469 int32
	_ = v10469
	var v10470 int32
	_ = v10470
	var v10474 int32
	_ = v10474
	var v10519 int32
	_ = v10519
	var v10529 int32
	_ = v10529
	var v10565 int32
	_ = v10565
	var v10566 int32
	_ = v10566
	var v10567 int32
	_ = v10567
	var v10568 int32
	_ = v10568
	var v10578 int32
	_ = v10578
	var v10581 int32
	_ = v10581
	var v10583 int32
	_ = v10583
	var v10584 int32
	_ = v10584
	var v10590 int32
	_ = v10590
	var v10591 int32
	_ = v10591
	var v10593 int32
	_ = v10593
	var v10594 int32
	_ = v10594
	var v10603 int32
	_ = v10603
	var v10605 float64
	_ = v10605
	var v10607 float64
	_ = v10607
	var v10609 float64
	_ = v10609
	var v10611 int32
	_ = v10611
	var v10612 int32
	_ = v10612
	var v10614 int32
	_ = v10614
	var v10616 int32
	_ = v10616
	var v10619 int32
	_ = v10619
	var v10620 int32
	_ = v10620
	var v10628 int32
	_ = v10628
	var v10630 int32
	_ = v10630
	var v10631 int32
	_ = v10631
	var v10644 int32
	_ = v10644
	var v10645 int32
	_ = v10645
	var v10646 int32
	_ = v10646
	var v10647 int32
	_ = v10647
	var v10649 int32
	_ = v10649
	var v10651 int32
	_ = v10651
	var v10653 int32
	_ = v10653
	var v10656 int32
	_ = v10656
	var v10657 int32
	_ = v10657
	var v10660 int32
	_ = v10660
	var v10664 int32
	_ = v10664
	var v10665 int32
	_ = v10665
	var v10666 int32
	_ = v10666
	var v10669 int32
	_ = v10669
	var v10681 int32
	_ = v10681
	var v10682 int32
	_ = v10682
	var v10685 int32
	_ = v10685
	var v10717 int32
	_ = v10717
	var v10721 int32
	_ = v10721
	var v10723 int32
	_ = v10723
	var v10724 int32
	_ = v10724
	var v10727 int32
	_ = v10727
	var v10728 int32
	_ = v10728
	var v10729 int32
	_ = v10729
	var v10741 int32
	_ = v10741
	var v10742 int32
	_ = v10742
	var v10743 int32
	_ = v10743
	var v10744 int32
	_ = v10744
	var v10746 int32
	_ = v10746
	var v10754 int32
	_ = v10754
	var v10755 int32
	_ = v10755
	var v10756 int32
	_ = v10756
	var v10759 int32
	_ = v10759
	var v10760 int32
	_ = v10760
	var v10762 int32
	_ = v10762
	var v10763 int32
	_ = v10763
	var v10765 int32
	_ = v10765
	var v10767 int32
	_ = v10767
	var v10770 int32
	_ = v10770
	var v10771 int32
	_ = v10771
	var v10772 int32
	_ = v10772
	var v10777 int32
	_ = v10777
	var v10778 int32
	_ = v10778
	var v10779 int32
	_ = v10779
	var v10782 int32
	_ = v10782
	var v10783 int32
	_ = v10783
	var v10784 int32
	_ = v10784
	var v10787 int32
	_ = v10787
	var v10788 int32
	_ = v10788
	var v10790 int32
	_ = v10790
	var v10795 int32
	_ = v10795
	var v10808 int32
	_ = v10808
	var v10809 int32
	_ = v10809
	var v10818 int32
	_ = v10818
	var v10822 int32
	_ = v10822
	var v10826 int32
	_ = v10826
	var v10828 int32
	_ = v10828
	var v10832 int32
	_ = v10832
	var v10833 int32
	_ = v10833
	var v10838 int32
	_ = v10838
	var v10841 int32
	_ = v10841
	var v10848 int32
	_ = v10848
	var v10850 int32
	_ = v10850
	var v10854 int32
	_ = v10854
	var v10862 int32
	_ = v10862
	var v10863 int32
	_ = v10863
	var v10864 int32
	_ = v10864
	var v10865 int32
	_ = v10865
	var v10867 int32
	_ = v10867
	var v10868 int32
	_ = v10868
	var v10871 int32
	_ = v10871
	var v10873 int32
	_ = v10873
	var v10874 int32
	_ = v10874
	var v10875 int32
	_ = v10875
	var v10881 int32
	_ = v10881
	var v10886 int32
	_ = v10886
	var v10888 int32
	_ = v10888
	var v10891 int32
	_ = v10891
	var v10892 float64
	_ = v10892
	var v10893 float64
	_ = v10893
	var v10894 int32
	_ = v10894
	var v10897 int32
	_ = v10897
	var v10898 float64
	_ = v10898
	var v10900 int32
	_ = v10900
	var v10901 int32
	_ = v10901
	var v10902 int32
	_ = v10902
	var v10907 float64
	_ = v10907
	var v10910 int32
	_ = v10910
	var v10911 float64
	_ = v10911
	var v10917 float64
	_ = v10917
	var v10923 float64
	_ = v10923
	var v10925 float64
	_ = v10925
	var v10927 float64
	_ = v10927
	var v10929 int32
	_ = v10929
	var v10930 int32
	_ = v10930
	var v10933 int32
	_ = v10933
	var v10935 int32
	_ = v10935
	var v10944 int32
	_ = v10944
	var v10945 int32
	_ = v10945
	var v10947 int32
	_ = v10947
	var v10948 int32
	_ = v10948
	var v10949 int32
	_ = v10949
	var v10951 int32
	_ = v10951
	var v10952 int32
	_ = v10952
	var v10957 int32
	_ = v10957
	var v10961 int32
	_ = v10961
	var v10966 int32
	_ = v10966
	var v10976 int32
	_ = v10976
	var v10979 int32
	_ = v10979
	var v10983 int32
	_ = v10983
	var v11014 int32
	_ = v11014
	var v11017 int32
	_ = v11017
	var v11019 int32
	_ = v11019
	var v11020 int32
	_ = v11020
	var v11021 int32
	_ = v11021
	var v11022 int32
	_ = v11022
	var v11024 int32
	_ = v11024
	var v11025 int32
	_ = v11025
	var v11026 int32
	_ = v11026
	var v11027 int32
	_ = v11027
	var v11028 int32
	_ = v11028
	var v11029 int32
	_ = v11029
	var v11030 int32
	_ = v11030
	var v11033 int32
	_ = v11033
	var v11034 int32
	_ = v11034
	var v11035 int32
	_ = v11035
	var v11041 int32
	_ = v11041
	var v11043 int32
	_ = v11043
	var v11045 float64
	_ = v11045
	var v11047 float64
	_ = v11047
	var v11049 float64
	_ = v11049
	var v11051 int32
	_ = v11051
	var v11052 int32
	_ = v11052
	var v11054 int32
	_ = v11054
	var v11056 int32
	_ = v11056
	var v11063 int32
	_ = v11063
	var v11064 int32
	_ = v11064
	var v11065 int32
	_ = v11065
	var v11066 int32
	_ = v11066
	var v11067 int32
	_ = v11067
	var v11068 int32
	_ = v11068
	var v11069 int32
	_ = v11069
	var v11070 int32
	_ = v11070
	var v11071 int32
	_ = v11071
	var v11075 int32
	_ = v11075
	var v11078 int32
	_ = v11078
	var v11081 int32
	_ = v11081
	var v11082 int32
	_ = v11082
	var v11084 int32
	_ = v11084
	var v11123 int32
	_ = v11123
	var v11127 int32
	_ = v11127
	var v11128 int32
	_ = v11128
	var v11129 int32
	_ = v11129
	var v11130 int32
	_ = v11130
	var v11131 int32
	_ = v11131
	var v11133 int32
	_ = v11133
	var v11135 int32
	_ = v11135
	var v11136 int32
	_ = v11136
	var v11142 int32
	_ = v11142
	var v11146 int32
	_ = v11146
	var v11147 int32
	_ = v11147
	var v11149 int32
	_ = v11149
	var v11150 int32
	_ = v11150
	var v11154 int32
	_ = v11154
	var v11196 int32
	_ = v11196
	var v11197 int32
	_ = v11197
	var v11198 int32
	_ = v11198
	var v11199 int32
	_ = v11199
	var v11200 int32
	_ = v11200
	var v11201 int32
	_ = v11201
	var v11202 int32
	_ = v11202
	var v11204 int32
	_ = v11204
	var v11206 int32
	_ = v11206
	var v11207 int32
	_ = v11207
	var v11208 int32
	_ = v11208
	var v11209 int32
	_ = v11209
	var v11210 int32
	_ = v11210
	var v11211 int32
	_ = v11211
	var v11212 int32
	_ = v11212
	var v11213 int32
	_ = v11213
	var v11215 int32
	_ = v11215
	var v11217 int32
	_ = v11217
	var v11218 int32
	_ = v11218
	var v11219 int32
	_ = v11219
	var v11221 int32
	_ = v11221
	var v11223 int32
	_ = v11223
	var v11224 int32
	_ = v11224
	var v11226 int32
	_ = v11226
	var v11230 int32
	_ = v11230
	var v11231 int32
	_ = v11231
	var v11237 int32
	_ = v11237
	var v11239 int32
	_ = v11239
	var v11240 int32
	_ = v11240
	var v11245 int32
	_ = v11245
	var v11246 int32
	_ = v11246
	var v11249 int32
	_ = v11249
	var v11250 int32
	_ = v11250
	var v11251 int32
	_ = v11251
	var v11253 int32
	_ = v11253
	var v11254 int32
	_ = v11254
	var v11255 int32
	_ = v11255
	var v11257 int32
	_ = v11257
	var v11260 int32
	_ = v11260
	var v11261 int32
	_ = v11261
	var v11263 int32
	_ = v11263
	var v11264 int32
	_ = v11264
	var v11265 int32
	_ = v11265
	var v11266 int32
	_ = v11266
	var v11267 int32
	_ = v11267
	var v11274 int32
	_ = v11274
	var v11277 int32
	_ = v11277
	var v11280 int32
	_ = v11280
	var v11315 int32
	_ = v11315
	var v11317 int32
	_ = v11317
	var v11321 int32
	_ = v11321
	var v11322 int32
	_ = v11322
	var v11323 int32
	_ = v11323
	var v11326 int32
	_ = v11326
	var v11327 int32
	_ = v11327
	var v11328 int32
	_ = v11328
	var v11329 int32
	_ = v11329
	var v11330 int32
	_ = v11330
	var v11331 int32
	_ = v11331
	var v11332 int32
	_ = v11332
	var v11335 int32
	_ = v11335
	var v11336 int32
	_ = v11336
	var v11337 int32
	_ = v11337
	var v11338 int32
	_ = v11338
	var v11347 int32
	_ = v11347
	var v11348 int32
	_ = v11348
	var v11350 int32
	_ = v11350
	var v11353 int32
	_ = v11353
	var v11354 int32
	_ = v11354
	var v11359 int32
	_ = v11359
	var v11366 int32
	_ = v11366
	var v11368 int32
	_ = v11368
	var v11370 int32
	_ = v11370
	var v11373 int32
	_ = v11373
	var v11375 int32
	_ = v11375
	var v11377 int32
	_ = v11377
	var v11384 int32
	_ = v11384
	var v11391 int32
	_ = v11391
	var v11394 int32
	_ = v11394
	var v11404 int32
	_ = v11404
	var v11405 int32
	_ = v11405
	var v11407 int32
	_ = v11407
	var v11410 int32
	_ = v11410
	var v11411 int32
	_ = v11411
	var v11416 int32
	_ = v11416
	var v11423 int32
	_ = v11423
	var v11425 int32
	_ = v11425
	var v11427 int32
	_ = v11427
	var v11428 int32
	_ = v11428
	var v11430 int32
	_ = v11430
	var v11432 int32
	_ = v11432
	var v11439 int32
	_ = v11439
	var v11442 int32
	_ = v11442
	var v11443 int32
	_ = v11443
	var v11444 int32
	_ = v11444
	var v11446 int32
	_ = v11446
	var v11447 int32
	_ = v11447
	var v11450 int32
	_ = v11450
	var v11451 int32
	_ = v11451
	var v11452 int32
	_ = v11452
	var v11454 int32
	_ = v11454
	var v11455 int32
	_ = v11455
	var v11456 int32
	_ = v11456
	var v11466 int32
	_ = v11466
	var v11467 int32
	_ = v11467
	var v11470 int32
	_ = v11470
	var v11474 int32
	_ = v11474
	var v11477 int32
	_ = v11477
	var v11479 int32
	_ = v11479
	var v11482 int32
	_ = v11482
	var v11489 int32
	_ = v11489
	var v11491 int32
	_ = v11491
	var v11499 int32
	_ = v11499
	var v11500 int32
	_ = v11500
	var v11513 int32
	_ = v11513
	var v11520 int32
	_ = v11520
	var v11521 int32
	_ = v11521
	var v11560 int32
	_ = v11560
	var v11562 int32
	_ = v11562
	var v11566 int32
	_ = v11566
	var v11569 int32
	_ = v11569
	var v11570 int32
	_ = v11570
	var v11571 int32
	_ = v11571
	var v11572 int32
	_ = v11572
	var v11574 int32
	_ = v11574
	var v11581 int32
	_ = v11581
	var v11583 int32
	_ = v11583
	var v11584 int32
	_ = v11584
	var v11587 int32
	_ = v11587
	var v11591 int32
	_ = v11591
	var v11594 int32
	_ = v11594
	var v11596 int32
	_ = v11596
	var v11599 int32
	_ = v11599
	var v11606 int32
	_ = v11606
	var v11608 int32
	_ = v11608
	var v11616 int32
	_ = v11616
	var v11617 int32
	_ = v11617
	var v11630 int32
	_ = v11630
	var v11638 int32
	_ = v11638
	var v11677 int32
	_ = v11677
	var v11678 int32
	_ = v11678
	var v11679 int32
	_ = v11679
	var v11680 int32
	_ = v11680
	var v11681 int32
	_ = v11681
	var v11683 int32
	_ = v11683
	var v11684 int32
	_ = v11684
	var v11688 int32
	_ = v11688
	var v11689 int32
	_ = v11689
	var v11690 int32
	_ = v11690
	var v11691 int32
	_ = v11691
	var v11693 int32
	_ = v11693
	var v11694 int32
	_ = v11694
	var v11695 int32
	_ = v11695
	var v11703 int32
	_ = v11703
	var v11743 int32
	_ = v11743
	var v11744 int32
	_ = v11744
	var v11748 int32
	_ = v11748
	var v11751 int32
	_ = v11751
	var v11754 int32
	_ = v11754
	var v11800 int32
	_ = v11800
	var v11879 int32
	_ = v11879
	var v11882 int32
	_ = v11882
	var v11883 int32
	_ = v11883
	var v11884 int32
	_ = v11884
	var v11891 int32
	_ = v11891
	var v11895 int32
	_ = v11895
	var v11896 int32
	_ = v11896
	var v11932 int32
	_ = v11932
	var v11936 int32
	_ = v11936
	var v11937 int32
	_ = v11937
	var v11938 int32
	_ = v11938
	var v11941 int32
	_ = v11941
	var v11942 int32
	_ = v11942
	var v11943 int32
	_ = v11943
	var v11944 int32
	_ = v11944
	var v11945 int32
	_ = v11945
	var v11947 int32
	_ = v11947
	var v11949 int32
	_ = v11949
	var v11950 int32
	_ = v11950
	var v11951 int32
	_ = v11951
	var v11952 int32
	_ = v11952
	var v11953 int32
	_ = v11953
	var v11957 int32
	_ = v11957
	var v11961 int32
	_ = v11961
	var v11965 int32
	_ = v11965
	var v11966 int32
	_ = v11966
	var v11967 int32
	_ = v11967
	var v11968 int32
	_ = v11968
	var v11971 int32
	_ = v11971
	var v11972 int32
	_ = v11972
	var v11974 int32
	_ = v11974
	var v11975 int32
	_ = v11975
	var v11977 int32
	_ = v11977
	var v11978 int32
	_ = v11978
	var v11987 int32
	_ = v11987
	var v11988 int32
	_ = v11988
	var v12024 int32
	_ = v12024
	var v12028 int32
	_ = v12028
	var v12029 int32
	_ = v12029
	var v12041 int32
	_ = v12041
	var v12074 int32
	_ = v12074
	var v12075 int32
	_ = v12075
	var v12076 int32
	_ = v12076
	var v12077 int32
	_ = v12077
	var v12079 int32
	_ = v12079
	var v12080 int32
	_ = v12080
	var v12091 int32
	_ = v12091
	var v12092 int32
	_ = v12092
	var v12093 int32
	_ = v12093
	var v12097 int32
	_ = v12097
	var v12100 int32
	_ = v12100
	var v12103 int32
	_ = v12103
	var v12104 int32
	_ = v12104
	var v12106 int32
	_ = v12106
	var v12145 int32
	_ = v12145
	var v12149 int32
	_ = v12149
	var v12150 int32
	_ = v12150
	var v12151 int32
	_ = v12151
	var v12152 int32
	_ = v12152
	var v12153 int32
	_ = v12153
	var v12155 int32
	_ = v12155
	var v12157 int32
	_ = v12157
	var v12158 int32
	_ = v12158
	var v12164 int32
	_ = v12164
	var v12168 int32
	_ = v12168
	var v12169 int32
	_ = v12169
	var v12171 int32
	_ = v12171
	var v12172 int32
	_ = v12172
	var v12176 int32
	_ = v12176
	var v12218 int32
	_ = v12218
	var v12219 int32
	_ = v12219
	var v12221 int32
	_ = v12221
	var v12224 int32
	_ = v12224
	var v12225 int32
	_ = v12225
	var v12226 int32
	_ = v12226
	var v12227 int32
	_ = v12227
	var v12229 int32
	_ = v12229
	var v12230 int32
	_ = v12230
	var v12231 int32
	_ = v12231
	var v12232 int32
	_ = v12232
	var v12233 int32
	_ = v12233
	var v12236 int32
	_ = v12236
	var v12240 int32
	_ = v12240
	var v12241 int32
	_ = v12241
	var v12247 int32
	_ = v12247
	var v12249 int32
	_ = v12249
	var v12250 int32
	_ = v12250
	var v12255 int32
	_ = v12255
	var v12256 int32
	_ = v12256
	var v12257 int32
	_ = v12257
	var v12258 int32
	_ = v12258
	var v12259 int32
	_ = v12259
	var v12260 int32
	_ = v12260
	var v12262 int32
	_ = v12262
	var v12263 int32
	_ = v12263
	var v12264 int32
	_ = v12264
	var v12266 int32
	_ = v12266
	var v12267 int32
	_ = v12267
	var v12268 int32
	_ = v12268
	var v12270 int32
	_ = v12270
	var v12271 int32
	_ = v12271
	var v12272 int32
	_ = v12272
	var v12273 int32
	_ = v12273
	var v12274 int32
	_ = v12274
	var v12275 int32
	_ = v12275
	var v12279 int32
	_ = v12279
	var v12280 int32
	_ = v12280
	var v12283 int32
	_ = v12283
	var v12284 int32
	_ = v12284
	var v12285 int32
	_ = v12285
	var v12286 int32
	_ = v12286
	var v12287 int32
	_ = v12287
	var v12288 int32
	_ = v12288
	var v12291 int32
	_ = v12291
	var v12292 int32
	_ = v12292
	var v12293 int32
	_ = v12293
	var v12294 int32
	_ = v12294
	var v12297 int32
	_ = v12297
	var v12298 int32
	_ = v12298
	var v12302 int32
	_ = v12302
	var v12303 int32
	_ = v12303
	var v12304 int32
	_ = v12304
	var v12305 int32
	_ = v12305
	var v12306 int32
	_ = v12306
	var v12312 int32
	_ = v12312
	var v12313 int32
	_ = v12313
	var v12314 int32
	_ = v12314
	var v12315 int32
	_ = v12315
	var v12320 int32
	_ = v12320
	var v12321 int32
	_ = v12321
	var v12322 int32
	_ = v12322
	var v12323 int32
	_ = v12323
	var v12329 int32
	_ = v12329
	var v12332 int32
	_ = v12332
	var v12333 int32
	_ = v12333
	var v12336 int32
	_ = v12336
	var v12337 int32
	_ = v12337
	var v12369 int32
	_ = v12369
	var v12373 int32
	_ = v12373
	var v12374 int32
	_ = v12374
	var v12375 int32
	_ = v12375
	var v12376 int32
	_ = v12376
	var v12377 int32
	_ = v12377
	var v12378 int32
	_ = v12378
	var v12379 int32
	_ = v12379
	var v12380 int32
	_ = v12380
	var v12381 int32
	_ = v12381
	var v12382 int32
	_ = v12382
	var v12383 int32
	_ = v12383
	var v12384 int32
	_ = v12384
	var v12385 int32
	_ = v12385
	var v12386 int32
	_ = v12386
	var v12387 int32
	_ = v12387
	var v12388 int32
	_ = v12388
	var v12389 int32
	_ = v12389
	var v12391 int32
	_ = v12391
	var v12392 int32
	_ = v12392
	var v12401 int32
	_ = v12401
	var v12402 int32
	_ = v12402
	var v12403 int32
	_ = v12403
	var v12405 int32
	_ = v12405
	var v12406 int32
	_ = v12406
	var v12407 int32
	_ = v12407
	var v12408 int32
	_ = v12408
	var v12439 int32
	_ = v12439
	var v12440 int32
	_ = v12440
	var v12443 int32
	_ = v12443
	var v12444 int32
	_ = v12444
	var v12445 int32
	_ = v12445
	var v12450 int32
	_ = v12450
	var v12456 int32
	_ = v12456
	var v12458 float64
	_ = v12458
	var v12460 float64
	_ = v12460
	var v12462 float64
	_ = v12462
	var v12464 int32
	_ = v12464
	var v12468 int32
	_ = v12468
	var v12471 int32
	_ = v12471
	var v12474 int32
	_ = v12474
	var v12476 float64
	_ = v12476
	var v12478 int32
	_ = v12478
	var v12479 int32
	_ = v12479
	var v12480 int32
	_ = v12480
	var v12481 int32
	_ = v12481
	var v12483 int32
	_ = v12483
	var v12484 int32
	_ = v12484
	var v12501 int32
	_ = v12501
	var v12542 int32
	_ = v12542
	var v12544 float64
	_ = v12544
	var v12546 float64
	_ = v12546
	var v12548 float64
	_ = v12548
	var v12550 int32
	_ = v12550
	var v12551 int32
	_ = v12551
	var v12553 int32
	_ = v12553
	var v12555 int32
	_ = v12555
	var v12557 int32
	_ = v12557
	var v12560 int32
	_ = v12560
	var v12561 int32
	_ = v12561
	var v12562 int32
	_ = v12562
	var v12564 int32
	_ = v12564
	var v12565 int32
	_ = v12565
	var v12568 int32
	_ = v12568
	var v12569 int32
	_ = v12569
	var v12573 int32
	_ = v12573
	var v12580 int32
	_ = v12580
	v4 = int32(0)
	v45 = m.G0
	v47 = v45 - int32(144)
	m.G0 = v47
	F_check_stack_depth(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v53 - int32(331) {
	case 0:
		goto L9
	case 1:
		goto L10
	case 2:
		goto L23
	case 3:
		goto L7
	case 4:
		goto L8
	case 5:
		goto L21
	default:
		goto L26
	case 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24:
		goto L28
	case 25:
		goto L6
	case 27:
		goto L27
	case 28:
		goto L5
	case 29:
		goto L11
	case 30:
		goto L12
	case 31:
		goto L15
	case 32:
		goto L16
	case 33:
		goto L17
	case 34:
		goto L18
	case 35:
		goto L19
	case 36:
		goto L13
	case 37:
		goto L14
	case 38:
		goto L25
	case 40:
		goto L20
	case 41:
		goto L22
	case 42:
		goto L24
	}
L3:
	;
	m.G0 = v12580 + int32(144)
	return v12573
L4:
	;
	v12542 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12501)+4)) = v12542
	v12544 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v12501)+8)) = v12544
	v12546 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v12501)+16)) = v12546
	v12548 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v12501)+24)) = v12548
	v12550 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12551 = *(*int32)(unsafe.Add(mBase, uint32(v12550)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12501)+32)) = v12551
	v12553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12501)+36)) = uint8(v12553)
	v12555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12501)+37)) = uint8(v12555)
	v12557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)))
	if v12557 != int32(1) {
		v12573 = v12501
		v12580 = v47
		goto L3
	} else {
		goto L1753
	}
L5:
	;
	v12091 = int32(0)
	v12092 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12093 = *(*int32)(unsafe.Add(mBase, uint32(v12092)+4))
	if v12093 == v12091 {
		v12176 = v12091
		goto L1691
	} else {
		goto L1692
	}
L6:
	;
	v11069 = int32(0)
	v11070 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11071 = *(*int32)(unsafe.Add(mBase, uint32(v11070)+4))
	if v11071 == v11069 {
		v11154 = v11069
		goto L1538
	} else {
		goto L1539
	}
L7:
	;
	v10389 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10390 = *(*int32)(unsafe.Add(mBase, uint32(v10389)+4))
	if v10390 == int32(0) {
		goto L1429
	} else {
		goto L1430
	}
L8:
	;
	v9796 = F_palloc0(m, int32(104))
	mBase = m.M
	v9797 = m.ExcPending
	if v9797 != 0 {
		goto L1
	} else {
		goto L1342
	}
L9:
	;
	v8867 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v8867 - int32(292) {
	case 0:
		goto L1224
	default:
		goto L1223
	case 9:
		goto L1226
	case 19:
		goto L1225
	}
L10:
	;
	v8710 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v8712 = F_create_plan_recurse(m, l0, v8710, int32(0))
	mBase = m.M
	v8713 = m.ExcPending
	if v8713 != 0 {
		goto L1
	} else {
		goto L1206
	}
L11:
	;
	v8678 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v8681 = F_create_plan_recurse(m, l0, v8678, l2|int32(2))
	mBase = m.M
	v8682 = m.ExcPending
	if v8682 != 0 {
		goto L1
	} else {
		goto L1204
	}
L12:
	;
	v8516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v8519 = F_create_plan_recurse(m, l0, v8516, l2|int32(2))
	mBase = m.M
	v8520 = m.ExcPending
	if v8520 != 0 {
		goto L1
	} else {
		goto L1176
	}
L13:
	;
	v6432 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6432 == int32(306) {
		goto L923
	} else {
		goto L924
	}
L14:
	;
	v6261 = int32(1)
	v6262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6264 = F_create_plan_recurse(m, l0, v6262, v6261)
	mBase = m.M
	v6265 = m.ExcPending
	if v6265 != 0 {
		goto L1
	} else {
		goto L905
	}
L15:
	;
	v6181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6184 = F_create_plan_recurse(m, l0, v6181, l2|int32(2))
	mBase = m.M
	v6185 = m.ExcPending
	if v6185 != 0 {
		goto L1
	} else {
		goto L899
	}
L16:
	;
	v6106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6109 = F_create_plan_recurse(m, l0, v6106, l2|int32(2))
	mBase = m.M
	v6110 = m.ExcPending
	if v6110 != 0 {
		goto L1
	} else {
		goto L893
	}
L17:
	;
	v5930 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v5932 = F_create_plan_recurse(m, l0, v5930, int32(4))
	mBase = m.M
	v5933 = m.ExcPending
	if v5933 != 0 {
		goto L1
	} else {
		goto L869
	}
L18:
	;
	v4340 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4340 == int32(310) {
		goto L673
	} else {
		goto L674
	}
L19:
	;
	v3884 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v3885 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+12))
	if v3885 != 0 {
		goto L628
	} else {
		goto L629
	}
L20:
	;
	v3556 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v3556)+4))
	if v3557 == int32(0) {
		v3646 = v4
		goto L582
	} else {
		goto L583
	}
L21:
	;
	v3195 = int32(1)
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3198 = F_create_plan_recurse(m, l0, v3196, v3195)
	mBase = m.M
	v3199 = m.ExcPending
	if v3199 != 0 {
		goto L1
	} else {
		goto L537
	}
L22:
	;
	v3163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3164 = F_create_plan_recurse(m, l0, v3163, l2)
	mBase = m.M
	v3165 = m.ExcPending
	if v3165 != 0 {
		goto L1
	} else {
		goto L535
	}
L23:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1422 = F_create_plan_recurse(m, l0, v1420, int32(1))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L1
	} else {
		goto L215
	}
L24:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1247 = F_create_plan_recurse(m, l0, v1246, l2)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L1
	} else {
		goto L198
	}
L25:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1068 = int32(0)
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+4))
	if v1070 == v1068 {
		v1153 = v1068
		goto L179
	} else {
		goto L180
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L176
	}
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v59 == int32(0) {
		v155 = v4
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v56 = F_create_scan_plan(m, l0, l1, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v12573 = v56
	v12580 = v47
	goto L3
L30:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v188 != 0 {
		goto L45
	} else {
		goto L46
	}
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 <= int32(0) {
		v155 = v4
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v70 = int32(1)
	v72 = v4
	v82 = v4
	goto L33
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v72<<(uint(int32(2))%32))))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v116 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v155 = v134
	goto L30
L35:
	;
	v117 = F_replace_nestloop_params_mutator(m, v115, l0)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v119 = v115
	goto L37
L37:
	;
	v121 = int32(0)
	v123 = F_makeTargetEntry(m, v119, base.I32_extend16_s(v70), v121, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v119 = v117
	goto L37
L39:
	;
	if v66 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v66+v70<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = v130
	goto L42
L41:
	;
	goto L42
L42:
	;
	v134 = F_lappend(m, v82, v123)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v137 = v72 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v137 < v138 {
		v70 = v70 + int32(1)
		v72 = v137
		v82 = v134
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	v189 = int32(2)
	goto L47
L46:
	;
	v189 = int32(0)
	goto L47
L47:
	;
	v190 = F_create_plan_recurse(m, l0, v185, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v195 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v196 = int32(2)
	goto L51
L50:
	;
	v196 = int32(0)
	goto L51
L51:
	;
	v197 = F_create_plan_recurse(m, l0, v192, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v200 = F_order_qual_clauses(m, l0, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v200
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v204)%32)&int32(174) != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v224 = F_get_actual_clauses(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L60
	}
L55:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	F_extract_actual_join_clauses(m, v200, v209, v47+int32(52), v47+int32(48))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v217 = F_extract_actual_clauses(m, v200, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v217
	goto L54
L60:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v227 = F_list_difference(m, v226, v224)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v227
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v230 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v231 = F_replace_nestloop_params_mutator(m, v227, l0)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v239 = l1 + int32(104)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	v244 = F_get_switched_clauses(m, v240, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v231
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v235 = F_replace_nestloop_params_mutator(m, v234, l0)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v235
	goto L64
L67:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v246 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v420 != 0 {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v248 = l1 + int32(100)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[0])))
	if v252 != int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v408 = v405 - int32(-64)
	v412 = v190
	goto L68
L72:
	;
	v319 = int32(0)
	v322 = v47 + int32(56)
	v331 = F_prepare_sort_from_pathkeys(m, v190, v246, v250, v319, v319, v322, v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L78
	}
L73:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v255 <= int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v258 = int32(0)
	v261 = v47 + int32(56)
	v270 = F_prepare_sort_from_pathkeys(m, v190, v246, v250, v258, v258, v261, v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v278 = F_palloc0(m, int32(104))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = int32(363)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v270)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+96)) = v255
	v284 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+56)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v278)+52)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v278)+48)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v278)+44)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v278)+88)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v278)+84)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v278)+80)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v278)+76)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v278)+72)) = v272
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	v297 = *(*float64)(unsafe.Add(mBase, uint32(v270)+8))
	v298 = *(*float64)(unsafe.Add(mBase, uint32(v270)+16))
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v270)+24))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v270)+32))
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	F_cost_incremental_sort(m, v261, l0, v295, v255, v296, v297, v298, v299, v300, v302, float64(-1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v278)+8)) = v306
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v278)+16)) = v308
	v310 = *(*float64)(unsafe.Add(mBase, uint32(v270)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v278)+24)) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+32))
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v278)+36)) = uint8(v313)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+32)) = v312
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v278)+37)) = uint8(v316)
	v408 = v248
	v412 = v278
	goto L68
L78:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v339 = F_palloc0(m, int32(96))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339))) = int32(362)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v331)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+44)) = v343
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	v346 = int32(_a_F_create_plan_recurse_0)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+88)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v339)+84)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v339)+80)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v339)+76)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v339)+72)) = v333
	v353 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v339)+56)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v339)+52)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v339)+48)) = v353
	v358 = int32(1)
	v360 = v345 + (v347 ^ v358)
	*(*int32)(unsafe.Add(mBase, uint32(v339)+4)) = v360
	v362 = *(*float64)(unsafe.Add(mBase, uint32(v331)+16))
	v363 = *(*float64)(unsafe.Add(mBase, uint32(v331)+24))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v331)+32))
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v370 = m.G0
	v371 = int32(16)
	v372 = v370 - v371
	m.G0 = v372
	F_cost_tuplesort(m, v372+int32(8), v372, v363, v364, float64(0), v367, float64(-1))
	mBase = m.M
	v377 = *(*float64)(unsafe.Add(mBase, uint32(v372)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v322)+32)) = v363
	v380 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v381 = base.F64_add(v362, v377)
	*(*float64)(unsafe.Add(mBase, uint32(v322)+48)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v322)+40)) = v360 + (v380 ^ v358)
	v387 = *(*float64)(unsafe.Add(mBase, uint32(v372)))
	*(*float64)(unsafe.Add(mBase, uint32(v322)+56)) = base.F64_add(v381, v387)
	m.G0 = v372 + v371
	goto L80
L80:
	;
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v339)+8)) = v393
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v339)+16)) = v395
	v397 = *(*float64)(unsafe.Add(mBase, uint32(v331)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v339)+24)) = v397
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v331)+32))
	v400 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+36)) = uint8(v400)
	*(*int32)(unsafe.Add(mBase, uint32(v339)+32)) = v399
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+37)) = uint8(v403)
	v408 = v248
	v412 = v339
	goto L68
L81:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+113)))
	if v524 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+8))
	v423 = int32(0)
	v426 = v47 + int32(56)
	v435 = F_prepare_sort_from_pathkeys(m, v197, v420, v422, v423, v423, v426, v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v513 = v197
	v515 = v509 - int32(-64)
	goto L81
L85:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v443 = F_palloc0(m, int32(96))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443))) = int32(362)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v435)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v443)+44)) = v447
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	v450 = int32(_a_F_create_plan_recurse_0)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v443)+88)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v443)+84)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v443)+80)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v443)+76)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v443)+72)) = v437
	v457 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v443)+56)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v443)+52)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v443)+48)) = v457
	v462 = int32(1)
	v464 = v449 + (v451 ^ v462)
	*(*int32)(unsafe.Add(mBase, uint32(v443)+4)) = v464
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v435)+16))
	v467 = *(*float64)(unsafe.Add(mBase, uint32(v435)+24))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v435)+32))
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v474 = m.G0
	v475 = int32(16)
	v476 = v474 - v475
	m.G0 = v476
	F_cost_tuplesort(m, v476+int32(8), v476, v467, v468, float64(0), v471, float64(-1))
	mBase = m.M
	v481 = *(*float64)(unsafe.Add(mBase, uint32(v476)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v426)+32)) = v467
	v484 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v485 = base.F64_add(v466, v481)
	*(*float64)(unsafe.Add(mBase, uint32(v426)+48)) = v485
	*(*int32)(unsafe.Add(mBase, uint32(v426)+40)) = v464 + (v484 ^ v462)
	v491 = *(*float64)(unsafe.Add(mBase, uint32(v476)))
	*(*float64)(unsafe.Add(mBase, uint32(v426)+56)) = base.F64_add(v485, v491)
	m.G0 = v476 + v475
	goto L87
L87:
	;
	v497 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v443)+8)) = v497
	v499 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v443)+16)) = v499
	v501 = *(*float64)(unsafe.Add(mBase, uint32(v435)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v443)+24)) = v501
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v435)+32))
	v504 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v443)+36)) = uint8(v504)
	*(*int32)(unsafe.Add(mBase, uint32(v443)+32)) = v503
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v443)+37)) = uint8(v507)
	v513 = v443
	v515 = v239
	goto L81
L88:
	;
	v562 = int32(0)
	if v244 != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	v559 = v513
	goto L88
L90:
	;
	goto L91
L91:
	;
	v528 = F_palloc0(m, int32(72))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = int32(360)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v513)+44))
	v533 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v528)+56)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v528)+52)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v528)+48)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v528)+44)) = v532
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v528)+4)) = v539
	v541 = *(*float64)(unsafe.Add(mBase, uint32(v513)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v528)+8)) = v541
	v543 = *(*float64)(unsafe.Add(mBase, uint32(v513)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v528)+16)) = v543
	v545 = *(*float64)(unsafe.Add(mBase, uint32(v513)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v528)+24)) = v545
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v513)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v528)+36)) = uint8(v533)
	*(*int32)(unsafe.Add(mBase, uint32(v528)+32)) = v547
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v528)+37)) = uint8(v551)
	v554 = *(*float64)(unsafe.Add(mBase, _c_F_create_plan_recurse[3]))
	*(*float64)(unsafe.Add(mBase, uint32(v528)+16)) = base.F64_add(v543, base.F64_mul(v545, v554))
	v559 = v528
	goto L88
L93:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v565 = v564
	goto L95
L94:
	;
	v565 = v562
	goto L95
L95:
	;
	v567 = v565 << (uint(int32(2)) % 32)
	v568 = F_palloc(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v570 = F_palloc(m, v567)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v572 = F_palloc(m, v565)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v574 = F_palloc(m, v565)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v419 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v577 = v576
	goto L102
L101:
	;
	v577 = v562
	goto L102
L102:
	;
	if v523 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	v580 = v578
	goto L105
L104:
	;
	v580 = int32(0)
	goto L105
L105:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	if v581 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L173
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L170
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L1
	} else {
		goto L167
	}
L109:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+112)))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v997 = F_palloc0(m, int32(112))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L166
	}
L110:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	if v584 <= int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v587 = int32(0)
	v594 = v587
	v596 = v587
	v597 = v4
	v600 = v580
	v602 = v577
	goto L112
L112:
	;
	v634 = v596 << (uint(int32(2)) % 32)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v581)+12))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v634+v635)))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+120)))
	if v640 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L109
L114:
	;
	v641 = int32(104)
	goto L116
L115:
	;
	v641 = int32(100)
	goto L116
L116:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v637+v641)))
	if v640 != 0 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v666)+8))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v874)+8))
	if v915 != v916 {
		goto L107
	} else {
		goto L158
	}
L118:
	;
	if v523 == int32(0) {
		v784 = v713
		v792 = v714
		goto L146
	} else {
		goto L147
	}
L119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L142
	}
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L139
	}
L121:
	;
	v646 = int32(100)
	goto L123
L122:
	;
	v646 = int32(104)
	goto L123
L123:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v637+v646)))
	if v597 != v648 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if v602 == int32(0) {
		goto L120
	} else {
		goto L127
	}
L125:
	;
	v666 = v594
	v667 = v597
	v668 = v602
	goto L126
L126:
	;
	if v600 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	if v648 != v653 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	v656 = v602 + int32(4)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if base.Ui32(v656) < base.Ui32(v658+v659<<(uint(int32(2))%32)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v664 = v656
	goto L131
L130:
	;
	v664 = int32(0)
	goto L131
L131:
	;
	v666 = v652
	v667 = v648
	v668 = v664
	goto L126
L132:
	;
	v671 = int32(0)
	v713 = v671
	v714 = v671
	goto L118
L133:
	;
	goto L134
L134:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	if v643 != v674 {
		v713 = v673
		v714 = v674
		goto L118
	} else {
		goto L135
	}
L135:
	;
	v677 = v600 + int32(4)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if base.Ui32(v677) < base.Ui32(v679+v680<<(uint(int32(2))%32)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v685 = v677
	goto L138
L137:
	;
	v685 = int32(0)
	goto L138
L138:
	;
	v874 = v673
	v881 = v685
	v914 = int32(1)
	goto L117
L139:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_1), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_3), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_1), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_5), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	v874 = v829
	v881 = v600
	v914 = int32(0)
	goto L117
L146:
	;
	if v792 != v643 {
		goto L108
	} else {
		goto L157
	}
L147:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if v717 <= int32(0) {
		v784 = v713
		v792 = v714
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v720 = int32(0)
	if v720 < v717 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v723 = v717
	goto L151
L150:
	;
	v723 = v720
	goto L151
L151:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v523)+12))
	v729 = int32(0)
	v730 = v713
	v738 = v714
	goto L152
L152:
	;
	v772 = v724 + v729<<(uint(int32(2))%32)
	if v772 == v600 {
		v784 = v730
		v792 = v738
		goto L146
	} else {
		goto L154
	}
L153:
	;
	v784 = v774
	v792 = v775
	goto L146
L154:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)+4))
	if v643 == v775 {
		v829 = v774
		goto L145
	} else {
		goto L155
	}
L155:
	;
	v778 = v729 + int32(1)
	if v778 != v723 {
		v729 = v778
		v730 = v774
		v738 = v775
		goto L152
	} else {
		goto L156
	}
L156:
	;
	goto L153
L157:
	;
	v829 = v784
	goto L145
L158:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v666)+4))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v918)+8))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v920)+8))
	if v919 != v921 {
		goto L107
	} else {
		goto L159
	}
L159:
	;
	if v914 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v666)+12))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v874)+12))
	if v923 != v924 {
		goto L106
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v568+v634))) = v915
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v666)+4))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v932)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v570+v634))) = v933
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v666)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v596+v572))) = uint8(base.B2i32(v936 == int32(5)))
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v596+v574))) = uint8(v941)
	v944 = v596 + int32(1)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	if v944 < v945 {
		v594 = v666
		v596 = v944
		v597 = v667
		v600 = v881
		v602 = v668
		goto L112
	} else {
		goto L165
	}
L163:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+16)))
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+16)))
	if v926 != v927 {
		goto L106
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	goto L113
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v997)+108)) = v574
	*(*int32)(unsafe.Add(mBase, uint32(v997)+104)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v997)+100)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v997)+96)) = v568
	*(*int32)(unsafe.Add(mBase, uint32(v997)+92)) = v244
	*(*uint8)(unsafe.Add(mBase, uint32(v997)+88)) = uint8(v993)
	*(*int32)(unsafe.Add(mBase, uint32(v997)+56)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v997)+52)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v997)+48)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v997)+44)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v997))) = int32(358)
	*(*int32)(unsafe.Add(mBase, uint32(v997)+80)) = v994
	*(*uint8)(unsafe.Add(mBase, uint32(v997)+76)) = uint8(v992)
	*(*int32)(unsafe.Add(mBase, uint32(v997)+72)) = v991
	v12501 = v997
	goto L4
L167:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_6), int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_7), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_8), int32(0))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_9), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_8), int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_10), int32(_a_F_create_plan_recurse_4))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v1057
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_11), v47)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(546), int32(_a_F_create_plan_recurse_12))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1197 = F_create_plan_recurse(m, l0, v1195, int32(1))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L1
	} else {
		goto L194
	}
L180:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	if v1074 <= int32(0) {
		v1153 = v1068
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+8))
	v1080 = v1068
	v1081 = int32(1)
	v1083 = v4
	goto L182
L182:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+12))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1122+v1083<<(uint(int32(2))%32))))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1127 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v1153 = v1145
	goto L179
L184:
	;
	v1128 = F_replace_nestloop_params_mutator(m, v1126, l0)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	v1130 = v1126
	goto L186
L186:
	;
	v1132 = int32(0)
	v1134 = F_makeTargetEntry(m, v1130, base.I32_extend16_s(v1081), v1132, v1132)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L1
	} else {
		goto L188
	}
L187:
	;
	v1130 = v1128
	goto L186
L188:
	;
	if v1077 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1077+v1081<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+16)) = v1141
	goto L191
L190:
	;
	goto L191
L191:
	;
	v1145 = F_lappend(m, v1080, v1134)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v1148 = v1083 + int32(1)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	if v1148 < v1149 {
		v1080 = v1145
		v1081 = v1081 + int32(1)
		v1083 = v1148
		goto L182
	} else {
		goto L193
	}
L193:
	;
	goto L183
L194:
	;
	v1200 = F_palloc0(m, int32(104))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1200)+44)) = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v1200))) = int32(369)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1200)+72)) = v1205
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1200)+4)) = v1207
	v1209 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1200)+8)) = v1209
	v1211 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1200)+16)) = v1211
	v1213 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1200)+24)) = v1213
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1215)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1200)+32)) = v1216
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1200)+36)) = uint8(v1218)
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1200)+37)) = uint8(v1220)
	v1222 = F_assign_special_exec_param(m, l0)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1200)+76)) = v1222
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+8))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+8))
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1200)+84))
	v1240 = F_prepare_sort_from_pathkeys(m, v1197, v1067, v1227, v1228, int32(0), v1200+int32(80), v1200+int32(84), v1200+int32(88), v1200+int32(92), v1200+int32(96))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1200)+52)) = v1240
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1244 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1243)+83)) = uint8(v1244)
	v12573 = v1200
	v12580 = v47
	goto L3
L198:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v1249 != int32(1) {
		v1345 = v4
		v1348 = v4
		v1349 = v4
		v1352 = v4
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1385 = *(*int64)(unsafe.Add(mBase, uint32(l1)+76))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v1388 = F_palloc0(m, int32(104))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L214
	}
L200:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+124))
	if v1253 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+4))
	v1255 = v1254
	goto L203
L202:
	;
	v1255 = v4
	goto L203
L203:
	;
	v1258 = F_palloc(m, v1255<<(uint(int32(1))%32))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v1261 = v1255 << (uint(int32(2)) % 32)
	v1262 = F_palloc(m, v1261)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1264 = F_palloc(m, v1261)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+124))
	if v1266 == int32(0) {
		v1345 = v4
		v1348 = v1258
		v1349 = v1264
		v1352 = v1262
		goto L199
	} else {
		goto L207
	}
L207:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	if v1269 <= int32(0) {
		v1345 = v4
		v1348 = v1258
		v1349 = v1264
		v1352 = v1262
		goto L199
	} else {
		goto L208
	}
L208:
	;
	v1276 = v4
	goto L209
L209:
	;
	v1320 = v1276 << (uint(int32(2)) % 32)
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+12))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1320+v1321)))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+76))
	v1325 = F_get_sortgroupclause_tle(m, v1323, v1324)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L1
	} else {
		goto L211
	}
L210:
	;
	v1345 = v1338
	v1348 = v1258
	v1349 = v1264
	v1352 = v1262
	goto L199
L211:
	;
	v1327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1325)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1258+v1276<<(uint(int32(1))%32)))) = uint16(v1327)
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1320+v1262))) = v1330
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1325)+4))
	v1334 = F_exprCollation(m, v1333)
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1320+v1264))) = v1334
	v1338 = v1276 + int32(1)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1266)+4))
	if v1338 < v1339 {
		v1276 = v1338
		goto L209
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1388))) = int32(373)
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+96)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+92)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+88)) = v1348
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+84)) = v1345
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+80)) = v1386
	*(*int64)(unsafe.Add(mBase, uint32(v1388)+72)) = v1385
	v1399 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+56)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+52)) = v1247
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+48)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+44)) = v1392
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+4)) = v1405
	v1407 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v1388)+8)) = v1407
	v1409 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1388)+16)) = v1409
	v1411 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1388)+24)) = v1411
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+32)) = v1414
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1388)+36)) = uint8(v1416)
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1388)+37)) = uint8(v1418)
	v12573 = v1388
	v12580 = v47
	goto L3
L215:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+44))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v1433 = int32(0)
	goto L217
L216:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+92)))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v1486 = F_palloc0(m, int32(168))
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L227
	}
L217:
	;
	v1434 = int32(0)
	if v1424 == v1434 {
		v1444 = v1434
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if v1425 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+4))
	if v1438 <= v1433 {
		v1444 = int32(0)
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+12))
	v1444 = v1440 + v1433<<(uint(int32(2))%32)
	goto L219
L222:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1444)))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1452+v1433<<(uint(int32(2))%32))))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1454)+12)) = v1459
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1454)+16)) = v1461
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1454)+20)) = v1463
	v1465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1458)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1454)+24)) = uint16(v1465)
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1454)+26)) = uint8(v1467)
	v1433 = v1433 + int32(1)
	goto L217
L223:
	;
	goto L216
L224:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+4))
	if base.B2i32(v1444 == int32(0))|base.B2i32(v1449 <= v1433) != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+12))
	if v1452 != 0 {
		goto L222
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	v1488 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+56)) = v1488
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+52)) = v1422
	*(*int32)(unsafe.Add(mBase, uint32(v1486))) = int32(333)
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+92)) = v1484
	*(*uint8)(unsafe.Add(mBase, uint32(v1486)+88)) = uint8(v1483)
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+84)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+80)) = v1481
	*(*uint8)(unsafe.Add(mBase, uint32(v1486)+76)) = uint8(v1480)
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+72)) = v1479
	*(*int64)(unsafe.Add(mBase, uint32(v1486)+44)) = int64(0)
	if v1474 == v1488 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+100)) = v1477
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+96)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+156)) = v2593
	v2635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v2635)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+104)) = v2636
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2638)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+164)) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+160)) = v1473
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+124)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+112)) = v1476
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+108)) = v2639
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+128)) = v1471
	if v1484 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L229:
	;
	v1503 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1486)+148)) = v1503
	*(*int64)(unsafe.Add(mBase, uint32(v1486)+140)) = v1503
	*(*int64)(unsafe.Add(mBase, uint32(v1486)+132)) = v1503
	v2591 = v1486
	v2593 = v4
	v2598 = v47
	goto L228
L230:
	;
	goto L231
L231:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+132)) = v1509
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+140)) = v1511
	v1513 = int32(0)
	if v1511 == v1513 {
		v1672 = v1513
		goto L232
	} else {
		goto L233
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+144)) = v1672
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+148)) = v1674
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1676)+84))
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+8))
	if v1678 != 0 {
		goto L245
	} else {
		goto L246
	}
L233:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+4))
	if int32(0) < v1518 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1523 = int32(1)
	v1527 = v4
	v1529 = v1513
	goto L237
L235:
	;
	v1590 = v4
	goto L236
L236:
	;
	v1672 = v1590
	goto L232
L237:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+12))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1565+v1529<<(uint(int32(2))%32))))
	v1570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569)+26)))
	if v1570 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1590 = v1576
	goto L236
L239:
	;
	v1573 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1569)+8)))
	v1574 = F_lappend_int(m, v1527, v1573)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	v1576 = v1527
	goto L241
L241:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1569)+8)) = uint16(v1523)
	v1578 = int32(1)
	v1581 = v1529 + v1578
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+4))
	if v1581 < v1582 {
		v1523 = v1523 + v1578
		v1527 = v1576
		v1529 = v1581
		goto L237
	} else {
		goto L243
	}
L242:
	;
	v1576 = v1574
	goto L241
L243:
	;
	goto L238
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+136)) = v2583
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1486)+152)) = v2585
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+32))
	v2591 = v1486
	v2593 = v2587
	v2598 = v47
	goto L228
L245:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1676)+52))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+12))
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1676)+32))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1682+v1683<<(uint(int32(2))%32)-int32(4))))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+16))
	v1692 = F_table_open(m, v1690, int32(0))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+16))
	if v1679 != 0 {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v2583 = int32(0)
	goto L244
L248:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+8))
	if v1694 == int32(0) {
		v1795 = v4
		v1796 = v4
		goto L250
	} else {
		goto L251
	}
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L1
	} else {
		goto L390
	}
L250:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+16))
	if v1814 != 0 {
		goto L264
	} else {
		goto L265
	}
L251:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+4))
	if v1697 <= int32(0) {
		v1795 = v4
		v1796 = v4
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v1703 = int32(0)
	v1726 = v4
	v1727 = v4
	goto L253
L253:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+12))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1745+v1703<<(uint(int32(2))%32))))
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+4))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1750)))
	if v1751 != int32(6) {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	v1795 = v1764
	v1796 = v1765
	goto L250
L255:
	;
	v1767 = v1703 + int32(1)
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+4))
	if v1767 < v1768 {
		v1703 = v1767
		v1726 = v1764
		v1727 = v1765
		goto L253
	} else {
		goto L262
	}
L256:
	;
	v1754 = F_lappend(m, v1726, v1750)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1756 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1750)+8)))
	if v1756 == int32(0) {
		goto L249
	} else {
		goto L260
	}
L259:
	;
	v1764 = v1754
	v1765 = v1727
	goto L255
L260:
	;
	v1761 = F_bms_add_member(m, v1727, v1756+int32(7))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	v1764 = v1726
	v1765 = v1761
	goto L255
L262:
	;
	goto L254
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L1
	} else {
		goto L386
	}
L264:
	;
	v1815 = F_get_constraint_index(m, v1814)
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	v1819 = v4
	goto L266
L266:
	;
	v1820 = F_RelationGetIndexList(m, v1692)
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L1
	} else {
		goto L271
	}
L267:
	;
	if v1815 == int32(0) {
		goto L263
	} else {
		goto L268
	}
L268:
	;
	v1819 = v1815
	goto L266
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L1
	} else {
		goto L382
	}
L270:
	;
	F_list_free(m, v1820)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L375
	}
L271:
	;
	if v1820 == int32(0) {
		v2440 = v4
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+4))
	if v1824 <= int32(0) {
		v2440 = v4
		goto L270
	} else {
		goto L273
	}
L273:
	;
	v1837 = int32(0)
	v1842 = v4
	goto L274
L274:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+12))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1872+v1837<<(uint(int32(2))%32))))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+24))
	v1878 = F_index_open(m, v1876, v1877)
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	v2440 = v2389
	goto L270
L276:
	;
	F_relation_close(m, v1878, int32(0))
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L1
	} else {
		goto L373
	}
L277:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+192))
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1880)+18)))
	if v1881 != int32(1) {
		v2389 = v1842
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1880)))
	if v1884 == v1819 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1880)+15)))
	if v1886 == int32(1) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	if v1819 != 0 {
		v2389 = v1842
		goto L276
	} else {
		goto L290
	}
L282:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+4))
	if v1889 == int32(2) {
		goto L269
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1892 = F_lappend_oid(m, v1842, v1819)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L1
	} else {
		goto L286
	}
L285:
	;
	goto L284
L286:
	;
	F_list_free(m, v1820)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_relation_close(m, v1878, int32(0))
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_relation_close(m, v1692, int32(0))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v2583 = v1892
	goto L244
L290:
	;
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1880)+12)))
	if v1902 != int32(1) {
		v2389 = v1842
		goto L276
	} else {
		goto L291
	}
L291:
	;
	v1905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1880)+15)))
	if v1905 != 0 {
		v2389 = v1842
		goto L276
	} else {
		goto L292
	}
L292:
	;
	v1906 = int32(0)
	v1908 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1880)+10)))
	if v1906 < v1908 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1916 = v1906
	v1917 = v1908
	v1922 = v1906
	goto L296
L294:
	;
	v1982 = v1906
	goto L295
L295:
	;
	v2015 = int32(0)
	if base.B2i32(v1982 == v2015)|base.B2i32(v1796 == v2015) != 0 {
		v2061 = base.B2i32(v1982|v1796 == v2015)
		goto L304
	} else {
		goto L305
	}
L296:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+192))
	v1959 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1955+v1916<<(uint(int32(1))%32))+48)))
	if v1959 != 0 {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1982 = v1966
	goto L295
L298:
	;
	v1962 = F_bms_add_member(m, v1922, v1959+int32(7))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L1
	} else {
		goto L301
	}
L299:
	;
	v1965 = v1917
	v1966 = v1922
	goto L300
L300:
	;
	v1968 = v1916 + int32(1)
	if v1968 < base.I32_extend16_s(v1965) {
		v1916 = v1968
		v1917 = v1965
		v1922 = v1966
		goto L296
	} else {
		goto L302
	}
L301:
	;
	v1964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1880)+10)))
	v1965 = v1964
	v1966 = v1962
	goto L300
L302:
	;
	goto L297
L303:
	;
	if v2061 == int32(0) {
		v2389 = v1842
		goto L276
	} else {
		goto L314
	}
L304:
	;
	goto L303
L305:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v1982)+4))
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v1796)+4))
	if v2029 != v2030 {
		v2061 = int32(0)
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v2032 = int32(1)
	if v2029 <= v2032 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v2035 = v2032
	goto L309
L308:
	;
	v2035 = v2029
	goto L309
L309:
	;
	v2036 = int32(8)
	v2041 = int32(0)
	goto L310
L310:
	;
	v2049 = v2041 << (uint(int32(2)) % 32)
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v1982+v2036+v2049)))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1796+v2036+v2049)))
	v2054 = base.B2i32(v2051 == v2053)
	if v2051 != v2053 {
		v2061 = v2054
		goto L304
	} else {
		goto L312
	}
L311:
	;
	v2061 = v2054
	goto L304
L312:
	;
	v2057 = v2041 + int32(1)
	if v2057 != v2035 {
		v2041 = v2057
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	v2068 = F_RelationGetIndexExpressions(m, v1878)
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v2071 = base.B2i32(v1683 == int32(1))
	v2072 = int32(0)
	if v2071|base.B2i32(v2068 == v2072) == v2072 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	F_ChangeVarNodes(m, v2068, int32(1), v1683)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+8))
	if v2080 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	goto L318
L320:
	;
	v2354 = F_list_difference(m, v2068, v1795)
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L1
	} else {
		goto L363
	}
L321:
	;
	v2083 = int32(0)
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+4))
	if v2084 <= v2083 {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v2109 = v2083
	goto L323
L323:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+12))
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2131+v2109<<(uint(int32(2))%32))))
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+8))
	if v2137 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L324:
	;
	goto L320
L325:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+4))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2296)))
	if v2297 == int32(6) {
		goto L356
	} else {
		goto L357
	}
L326:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+52))
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v2151)))
	if v2152 <= int32(0) {
		v2389 = v1842
		goto L276
	} else {
		goto L335
	}
L327:
	;
	v2144 = F_get_opclass_family(m, v2136)
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L1
	} else {
		goto L333
	}
L328:
	;
	if v2136 == int32(0) {
		goto L325
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	if v2136 != 0 {
		goto L327
	} else {
		goto L332
	}
L331:
	;
	goto L327
L332:
	;
	v2142 = int32(0)
	v2149 = v2142
	v2150 = v2142
	goto L326
L333:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	v2147 = F_get_opclass_input_type(m, v2146)
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v2149 = v2144
	v2150 = v2147
	goto L326
L335:
	;
	v2156 = int32(1)
	v2163 = v2156
	v2169 = int32(0)
	v2186 = v2152
	v2187 = v2156
	goto L336
L336:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+192))
	v2203 = int32(1)
	v2204 = v2163 - v2203
	v2208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2202+v2204<<(uint(v2203)%32))+48)))
	v2210 = v2204 << (uint(int32(2)) % 32)
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+248))
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2210+v2211)))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	if v2214 != 0 {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	v2389 = v1842
	goto L276
L338:
	;
	v2249 = v2187 + int32(1)
	v2250 = base.I32_extend16_s(v2249)
	if v2250 <= v2244 {
		v2163 = v2250
		v2169 = v2169 + base.B2i32(v2208 != int32(0))
		v2186 = v2244
		v2187 = v2249
		goto L336
	} else {
		goto L355
	}
L339:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+208))
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2215+v2210)))
	if v2149 != v2217 {
		v2244 = v2186
		goto L338
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+8))
	if v2223 != v2213 {
		goto L344
	} else {
		goto L345
	}
L342:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+212))
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2219+v2210)))
	if v2150 != v2221 {
		v2244 = v2186
		goto L338
	} else {
		goto L343
	}
L343:
	;
	goto L341
L344:
	;
	v2226 = v2223
	goto L346
L345:
	;
	v2226 = int32(0)
	goto L346
L346:
	;
	if v2226 != 0 {
		v2244 = v2186
		goto L338
	} else {
		goto L347
	}
L347:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+4))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2227)))
	if v2228 == int32(6) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2227)+8)))
	if v2231 != v2208 {
		v2244 = v2186
		goto L338
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	if v2208 != 0 {
		v2244 = v2186
		goto L338
	} else {
		goto L352
	}
L351:
	;
	goto L325
L352:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+12))
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v2233+(v2204-v2169)<<(uint(int32(2))%32))))
	v2239 = F_equal(m, v2227, v2238)
	mBase = m.M
	v2240 = m.ExcPending
	if v2240 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	if v2239 != 0 {
		goto L325
	} else {
		goto L354
	}
L354:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+52))
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2241)))
	v2244 = v2242
	goto L338
L355:
	;
	goto L337
L356:
	;
	v2307 = v2109 + int32(1)
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2080)+4))
	if v2307 < v2308 {
		v2109 = v2307
		goto L323
	} else {
		goto L362
	}
L357:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+8))
	if v2300 != 0 {
		goto L356
	} else {
		goto L358
	}
L358:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+12))
	if v2301 != 0 {
		goto L356
	} else {
		goto L359
	}
L359:
	;
	v2302 = F_list_member(m, v2068, v2296)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	if v2302 == int32(0) {
		v2389 = v1842
		goto L276
	} else {
		goto L361
	}
L361:
	;
	goto L356
L362:
	;
	goto L324
L363:
	;
	if v2354 != 0 {
		v2389 = v1842
		goto L276
	} else {
		goto L364
	}
L364:
	;
	v2356 = F_RelationGetIndexPredicate(m, v1878)
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v2358 = int32(0)
	if v2071|base.B2i32(v2356 == v2358) == v2358 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	F_ChangeVarNodes(m, v2356, int32(1), v1683)
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L1
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+12))
	v2368 = F_predicate_implied_by(m, v2356, v2366, int32(0))
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L1
	} else {
		goto L370
	}
L369:
	;
	goto L368
L370:
	;
	if v2368 == int32(0) {
		v2389 = v1842
		goto L276
	} else {
		goto L371
	}
L371:
	;
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v1880)))
	v2373 = F_lappend_oid(m, v1842, v2372)
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	v2389 = v2373
	goto L276
L373:
	;
	v2423 = v1837 + int32(1)
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+4))
	if v2423 < v2424 {
		v1837 = v2423
		v1842 = v2389
		goto L274
	} else {
		goto L374
	}
L374:
	;
	goto L275
L375:
	;
	F_relation_close(m, v1692, int32(0))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	if v2440 != 0 {
		v2583 = v2440
		goto L244
	} else {
		goto L377
	}
L377:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	F_errcode(m, int32(_a_F_create_plan_recurse_13))
	mBase = m.M
	v2481 = m.ExcPending
	if v2481 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_14), int32(0))
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_15), int32(960), int32(_a_F_create_plan_recurse_16))
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L382:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2497 = m.ExcPending
	if v2497 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_17), int32(0))
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_15), int32(843), int32(_a_F_create_plan_recurse_16))
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L386:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_18), int32(0))
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_15), int32(793), int32(_a_F_create_plan_recurse_16))
	mBase = m.M
	v2522 = m.ExcPending
	if v2522 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_19), int32(0))
	mBase = m.M
	v2533 = m.ExcPending
	if v2533 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_15), int32(776), int32(_a_F_create_plan_recurse_16))
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+120)) = v3115
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+116)) = v3113
	v3148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+4)) = v3148
	v3150 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v2591)+8)) = v3150
	v3152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v2591)+16)) = v3152
	v3154 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v2591)+24)) = v3154
	v3156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(v3156)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2591)+32)) = v3157
	v3159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2591)+36)) = uint8(v3159)
	v3161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2591)+37)) = uint8(v3161)
	v12573 = v2591
	v12580 = v2598
	goto L3
L395:
	;
	v2648 = int32(0)
	v3113 = v2648
	v3115 = v2648
	goto L394
L396:
	;
	goto L397
L397:
	;
	v2650 = int32(0)
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v1484)+4))
	if v2651 <= v2650 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v3113 = int32(0)
	v3115 = v2650
	goto L394
L399:
	;
	goto L400
L400:
	;
	v2657 = int32(0)
	v2667 = v2657
	v2671 = v2657
	v2674 = v2657
	v2676 = v2650
	v2677 = v2657
	v2678 = v2657
	v2681 = v2657
	goto L401
L401:
	;
	v2707 = *(*int32)(unsafe.Add(mBase, uint32(v1484)+12))
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2707+v2667<<(uint(int32(2))%32))))
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v2712) <= base.Ui32(v2711) {
		goto L407
	} else {
		goto L408
	}
L402:
	;
	v3113 = v3096
	v3115 = v3090
	goto L394
L403:
	;
	v3096 = F_lappend(m, v2674, v3085)
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L1
	} else {
		goto L533
	}
L404:
	;
	v3081 = F_bms_add_member(m, v2676, v2667)
	mBase = m.M
	v3082 = m.ExcPending
	if v3082 != 0 {
		goto L1
	} else {
		goto L532
	}
L405:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3066 = m.ExcPending
	if v3066 != 0 {
		goto L1
	} else {
		goto L528
	}
L406:
	;
	v2752 = int32(0)
	if base.B2i32(v1479 != int32(5))|base.B2i32(v2751 == v2752) == v2752 {
		goto L418
	} else {
		goto L419
	}
L407:
	;
	v2723 = int32(0)
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2724 != 0 {
		goto L411
	} else {
		goto L412
	}
L408:
	;
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(v2714+v2711<<(uint(int32(2))%32))))
	if v2718 == int32(0) {
		goto L407
	} else {
		goto L409
	}
L409:
	;
	v2721 = *(*int32)(unsafe.Add(mBase, uint32(v2718)+168))
	v2751 = v2721
	goto L406
L410:
	;
	v2737 = *(*int32)(unsafe.Add(mBase, uint32(v2736)))
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2737)+12))
	if v2738 != 0 {
		v3085 = v2723
		v3087 = v2671
		v3090 = v2676
		v3091 = v2677
		v3092 = v2678
		v3093 = v2681
		goto L403
	} else {
		goto L414
	}
L411:
	;
	v2736 = v2724 + v2711<<(uint(int32(2))%32)
	goto L410
L412:
	;
	goto L413
L413:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2729 = *(*int32)(unsafe.Add(mBase, uint32(v2728)+52))
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2729)+12))
	v2736 = v2730 + v2711<<(uint(int32(2))%32) - int32(4)
	goto L410
L414:
	;
	v2739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2737)+21)))
	if v2739 != int32(102) {
		v3085 = v2723
		v3087 = v2671
		v3090 = v2676
		v3091 = v2677
		v3092 = v2678
		v3093 = v2681
		goto L403
	} else {
		goto L415
	}
L415:
	;
	v2743 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[4])))
	if v2743&int32(2) != 0 {
		goto L405
	} else {
		goto L416
	}
L416:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2737)+16))
	v2747 = F_GetFdwRoutineByRelId(m, v2746)
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L1
	} else {
		goto L417
	}
L417:
	;
	v2751 = v2747
	goto L406
L418:
	;
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2757 != 0 {
		goto L422
	} else {
		goto L423
	}
L419:
	;
	goto L420
L420:
	;
	v2795 = int32(0)
	if v2751 == v2795 {
		v3085 = v2795
		v3087 = v2671
		v3090 = v2676
		v3091 = v2677
		v3092 = v2678
		v3093 = v2681
		goto L403
	} else {
		goto L431
	}
L421:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2769)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		goto L1
	} else {
		goto L425
	}
L422:
	;
	v2769 = v2757 + v2711<<(uint(int32(2))%32)
	goto L421
L423:
	;
	goto L424
L424:
	;
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v2761)+52))
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v2762)+12))
	v2769 = v2763 + v2711<<(uint(int32(2))%32) - int32(4)
	goto L421
L425:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v2770)+16))
	v2779 = F_get_rel_name(m, v2778)
	mBase = m.M
	v2780 = m.ExcPending
	if v2780 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2598)+16)) = v2779
	F_errmsg(m, int32(_a_F_create_plan_recurse_20), v2598+int32(16))
	mBase = m.M
	v2786 = m.ExcPending
	if v2786 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	v2787 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2770)+21)))
	F_errdetail_relkind_not_supported(m, v2787)
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_21), int32(_a_F_create_plan_recurse_22))
	mBase = m.M
	v2794 = m.ExcPending
	if v2794 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L431:
	;
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+88))
	if v2798 == int32(0) {
		v3051 = v2671
		v3053 = v2677
		v3054 = v2678
		v3055 = v2681
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v3058 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+44))
	if v3058 == int32(0) {
		v3085 = v2795
		v3087 = v3051
		v3090 = v2676
		v3091 = v3053
		v3092 = v3054
		v3093 = v3055
		goto L403
	} else {
		goto L526
	}
L433:
	;
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+92))
	if v2801 == int32(0) {
		v3051 = v2671
		v3053 = v2677
		v3054 = v2678
		v3055 = v2681
		goto L432
	} else {
		goto L434
	}
L434:
	;
	v2804 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+96))
	if base.B2i32(v2804 == int32(0))|v1477 != 0 {
		v3051 = v2671
		v3053 = v2677
		v3054 = v2678
		v3055 = v2681
		goto L432
	} else {
		goto L435
	}
L435:
	;
	v2808 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+100))
	if v2808 == int32(0) {
		v3051 = v2671
		v3053 = v2677
		v3054 = v2678
		v3055 = v2681
		goto L432
	} else {
		goto L436
	}
L436:
	;
	v2811 = m.G0
	v2813 = v2811 - int32(16)
	m.G0 = v2813
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2815 != 0 {
		goto L439
	} else {
		goto L440
	}
L437:
	;
	if v2856 != 0 {
		v3051 = v2671
		v3053 = v2677
		v3054 = v2678
		v3055 = v2681
		goto L432
	} else {
		goto L462
	}
L438:
	;
	v2828 = *(*int32)(unsafe.Add(mBase, uint32(v2827)))
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v2828)+16))
	v2831 = F_table_open(m, v2829, int32(0))
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L1
	} else {
		goto L442
	}
L439:
	;
	v2827 = v2815 + v2711<<(uint(int32(2))%32)
	goto L438
L440:
	;
	goto L441
L441:
	;
	v2819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v2819)+52))
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+12))
	v2827 = v2821 + v2711<<(uint(int32(2))%32) - int32(4)
	goto L438
L442:
	;
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v2831)+76))
	v2834 = int32(0)
	switch v1479 - int32(2) {
	case 0:
		goto L447
	case 1:
		goto L448
	case 2:
		goto L446
	case 3:
		v2856 = v2834
		goto L444
	default:
		goto L443
	}
L443:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2866 = m.ExcPending
	if v2866 != 0 {
		goto L1
	} else {
		goto L459
	}
L444:
	;
	F_relation_close(m, v2831, int32(0))
	mBase = m.M
	v2859 = m.ExcPending
	if v2859 != 0 {
		goto L1
	} else {
		goto L458
	}
L445:
	;
	v2856 = int32(1)
	goto L444
L446:
	;
	if v2833 == int32(0) {
		v2856 = v2834
		goto L444
	} else {
		goto L455
	}
L447:
	;
	if v2833 == int32(0) {
		v2856 = v2834
		goto L444
	} else {
		goto L452
	}
L448:
	;
	if v2833 == int32(0) {
		v2856 = v2834
		goto L444
	} else {
		goto L449
	}
L449:
	;
	v2839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2833)+9)))
	if v2839 != 0 {
		goto L445
	} else {
		goto L450
	}
L450:
	;
	v2840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2833)+8)))
	if v2840 == int32(1) {
		goto L445
	} else {
		goto L451
	}
L451:
	;
	v2856 = v2834
	goto L444
L452:
	;
	v2845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2833)+14)))
	if v2845 != 0 {
		goto L445
	} else {
		goto L453
	}
L453:
	;
	v2846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2833)+13)))
	if v2846 == int32(1) {
		goto L445
	} else {
		goto L454
	}
L454:
	;
	v2856 = v2834
	goto L444
L455:
	;
	v2851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2833)+19)))
	if v2851 != 0 {
		goto L445
	} else {
		goto L456
	}
L456:
	;
	v2852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2833)+18)))
	if v2852 != int32(1) {
		v2856 = v2834
		goto L444
	} else {
		goto L457
	}
L457:
	;
	goto L445
L458:
	;
	m.G0 = v2813 + int32(16)
	goto L437
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2813))) = v1479
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_23), v2813)
	mBase = m.M
	v2870 = m.ExcPending
	if v2870 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_15), int32(2312), int32(_a_F_create_plan_recurse_24))
	mBase = m.M
	v2875 = m.ExcPending
	if v2875 != 0 {
		goto L1
	} else {
		goto L461
	}
L461:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L462:
	;
	v2877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2877 != 0 {
		goto L464
	} else {
		goto L465
	}
L463:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v2889)))
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v2890)+16))
	v2893 = F_table_open(m, v2891, int32(0))
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
		goto L1
	} else {
		goto L467
	}
L464:
	;
	v2889 = v2877 + v2711<<(uint(int32(2))%32)
	goto L463
L465:
	;
	goto L466
L466:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2882 = *(*int32)(unsafe.Add(mBase, uint32(v2881)+52))
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2882)+12))
	v2889 = v2883 + v2711<<(uint(int32(2))%32) - int32(4)
	goto L463
L467:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v2893)+52))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2895)+16))
	if v2896 != 0 {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v2897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2896)+17)))
	v2898 = v2897
	goto L470
L469:
	;
	v2898 = int32(0)
	goto L470
L470:
	;
	F_relation_close(m, v2893, int32(0))
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	if v2898&int32(1) != 0 {
		v3051 = v2671
		v3053 = v2677
		v3054 = v2678
		v3055 = v2681
		goto L432
	} else {
		goto L472
	}
L472:
	;
	if v2677&int32(1) == int32(0) {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	if v2681 == int32(0) {
		goto L494
	} else {
		goto L495
	}
L474:
	;
	v2908 = int32(0)
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2910 = *(*int32)(unsafe.Add(mBase, uint32(v2909)+96))
	if v2910 == v2908 {
		v2932 = v2908
		goto L477
	} else {
		goto L478
	}
L475:
	;
	goto L476
L476:
	;
	v2937 = int32(1)
	if v2678&v2937 == int32(0) {
		goto L473
	} else {
		goto L490
	}
L477:
	;
	if v2932 == int32(0) {
		goto L473
	} else {
		goto L489
	}
L478:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v2910)))
	if v2913 != int32(61) {
		goto L480
	} else {
		goto L481
	}
L479:
	;
	v2929 = F_expression_tree_walker_impl(m, v2910, int32(901), int32(0))
	mBase = m.M
	v2930 = m.ExcPending
	if v2930 != 0 {
		goto L1
	} else {
		goto L488
	}
L480:
	;
	if v2913 != int32(6) {
		goto L479
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2910)+4))
	v2932 = base.B2i32(v2924 == int32(0))
	goto L477
L483:
	;
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v2910)+28))
	if v2918 == int32(0) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2922 = *(*int32)(unsafe.Add(mBase, uint32(v2910)+32))
	if v2922 != 0 {
		v2932 = int32(1)
		goto L477
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	v2932 = int32(0)
	goto L477
L487:
	;
	goto L486
L488:
	;
	v2932 = v2929
	goto L477
L489:
	;
	v2935 = int32(1)
	v3051 = v2671
	v3053 = v2935
	v3054 = v2935
	v3055 = v2681
	goto L432
L490:
	;
	v3051 = v2671
	v3053 = v2937
	v3054 = int32(1)
	v3055 = v2681
	goto L432
L491:
	;
	v3051 = v3048
	v3053 = v3046
	v3054 = v3047
	v3055 = int32(1)
	goto L432
L492:
	;
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+88))
	v3038 = m.T0[v3037].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v2591, v2711, v2667)
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L1
	} else {
		goto L524
	}
L493:
	;
	v3046 = v3028
	v3047 = int32(0)
	v3048 = int32(1)
	goto L491
L494:
	;
	v2948 = m.G0
	v2950 = v2948 - int32(16)
	m.G0 = v2950
	v2952 = int32(0)
	v2953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2953 != 0 {
		goto L500
	} else {
		goto L501
	}
L495:
	;
	goto L496
L496:
	;
	v3020 = int32(1)
	if v2671&v3020 == int32(0) {
		goto L492
	} else {
		goto L523
	}
L497:
	;
	if v2997&int32(1) == int32(0) {
		goto L492
	} else {
		goto L522
	}
L498:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3007 = m.ExcPending
	if v3007 != 0 {
		goto L1
	} else {
		goto L519
	}
L499:
	;
	v2966 = *(*int32)(unsafe.Add(mBase, uint32(v2965)))
	v2967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2966)+21)))
	if v2967 != int32(102) {
		goto L503
	} else {
		goto L504
	}
L500:
	;
	v2965 = v2953 + v1481<<(uint(int32(2))%32)
	goto L499
L501:
	;
	goto L502
L502:
	;
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v2957)+52))
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2958)+12))
	v2965 = v2959 + v1481<<(uint(int32(2))%32) - int32(4)
	goto L499
L503:
	;
	v2970 = *(*int32)(unsafe.Add(mBase, uint32(v2966)+16))
	v2972 = F_table_open(m, v2970, int32(0))
	mBase = m.M
	v2973 = m.ExcPending
	if v2973 != 0 {
		goto L1
	} else {
		goto L506
	}
L504:
	;
	v2997 = v2952
	goto L505
L505:
	;
	m.G0 = v2950 + int32(16)
	goto L497
L506:
	;
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(v2972)+76))
	switch v1479 - int32(2) {
	case 0:
		goto L509
	case 1:
		goto L510
	case 2:
		goto L508
	case 3:
		v2992 = v2952
		goto L507
	default:
		goto L498
	}
L507:
	;
	F_relation_close(m, v2972, int32(0))
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L1
	} else {
		goto L518
	}
L508:
	;
	if v2974 == int32(0) {
		v2992 = v2952
		goto L507
	} else {
		goto L517
	}
L509:
	;
	if v2974 == int32(0) {
		v2992 = v2952
		goto L507
	} else {
		goto L512
	}
L510:
	;
	if v2974 == int32(0) {
		v2992 = v2952
		goto L507
	} else {
		goto L511
	}
L511:
	;
	v2979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2974)+25)))
	v2992 = v2979
	goto L507
L512:
	;
	v2982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2974)+26)))
	if v2982 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v2985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2974)+27)))
	if v2985 != int32(1) {
		v2992 = v2952
		goto L507
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	v2992 = int32(1)
	goto L507
L516:
	;
	goto L515
L517:
	;
	v2991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2974)+28)))
	v2992 = v2991
	goto L507
L518:
	;
	v2997 = v2992
	goto L505
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2950))) = v1479
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_23), v2950)
	mBase = m.M
	v3011 = m.ExcPending
	if v3011 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_15), int32(2366), int32(_a_F_create_plan_recurse_25))
	mBase = m.M
	v3016 = m.ExcPending
	if v3016 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L522:
	;
	v3028 = int32(1)
	goto L493
L523:
	;
	v3028 = v3020
	goto L493
L524:
	;
	if v3038 != 0 {
		goto L404
	} else {
		goto L525
	}
L525:
	;
	v3041 = int32(0)
	v3046 = int32(1)
	v3047 = v3041
	v3048 = v3041
	goto L491
L526:
	;
	v3061 = m.T0[v3058].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v2591, v2711, v2667)
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	v3085 = v3061
	v3087 = v3051
	v3090 = v2676
	v3091 = v3053
	v3092 = v3054
	v3093 = v3055
	goto L403
L528:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v3069 = m.ExcPending
	if v3069 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	F_errmsg(m, int32(_a_F_create_plan_recurse_26), int32(0))
	mBase = m.M
	v3073 = m.ExcPending
	if v3073 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_27), int32(_a_F_create_plan_recurse_22))
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L532:
	;
	v3085 = v2795
	v3087 = int32(0)
	v3090 = v3081
	v3091 = int32(1)
	v3092 = int32(0)
	v3093 = int32(1)
	goto L403
L533:
	;
	v3099 = v2667 + int32(1)
	v3100 = *(*int32)(unsafe.Add(mBase, uint32(v1484)+4))
	if v3099 < v3100 {
		v2667 = v3099
		v2671 = v3087
		v2674 = v3096
		v2676 = v3090
		v2677 = v3091
		v2678 = v3092
		v2681 = v3093
		goto L401
	} else {
		goto L534
	}
L534:
	;
	goto L402
L535:
	;
	v3166 = *(*int64)(unsafe.Add(mBase, uint32(l1)+76))
	v3168 = F_palloc0(m, int32(80))
	mBase = m.M
	v3169 = m.ExcPending
	if v3169 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3168))) = int32(372)
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(v3164)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v3168)+72)) = v3166
	v3174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+56)) = v3174
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+52)) = v3164
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+48)) = v3174
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+44)) = v3172
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+4)) = v3180
	v3182 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3168)+8)) = v3182
	v3184 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3168)+16)) = v3184
	v3186 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3168)+24)) = v3186
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(v3188)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3168)+32)) = v3189
	v3191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3168)+36)) = uint8(v3191)
	v3193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3168)+37)) = uint8(v3193)
	v12573 = v3168
	v12580 = v47
	goto L3
L537:
	;
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v3202 = F_create_plan_recurse(m, l0, v3200, int32(1))
	mBase = m.M
	v3203 = m.ExcPending
	if v3203 != 0 {
		goto L1
	} else {
		goto L538
	}
L538:
	;
	v3204 = int32(0)
	v3205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+4))
	if v3206 == v3204 {
		v3288 = v3204
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v3330 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v3330)&int64(9223372036854775807)) {
		goto L555
	} else {
		goto L556
	}
L540:
	;
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3206)+4))
	if v3209 <= int32(0) {
		v3288 = v3204
		goto L539
	} else {
		goto L541
	}
L541:
	;
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+8))
	v3215 = v3204
	v3216 = v3195
	v3218 = v4
	goto L542
L542:
	;
	v3257 = *(*int32)(unsafe.Add(mBase, uint32(v3206)+12))
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(v3257+v3218<<(uint(int32(2))%32))))
	v3262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3262 != 0 {
		goto L544
	} else {
		goto L545
	}
L543:
	;
	v3288 = v3280
	goto L539
L544:
	;
	v3263 = F_replace_nestloop_params_mutator(m, v3261, l0)
	mBase = m.M
	v3264 = m.ExcPending
	if v3264 != 0 {
		goto L1
	} else {
		goto L547
	}
L545:
	;
	v3265 = v3261
	goto L546
L546:
	;
	v3267 = int32(0)
	v3269 = F_makeTargetEntry(m, v3265, base.I32_extend16_s(v3216), v3267, v3267)
	mBase = m.M
	v3270 = m.ExcPending
	if v3270 != 0 {
		goto L1
	} else {
		goto L548
	}
L547:
	;
	v3265 = v3263
	goto L546
L548:
	;
	if v3212 != 0 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3212+v3216<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+16)) = v3276
	goto L551
L550:
	;
	goto L551
L551:
	;
	v3280 = F_lappend(m, v3215, v3269)
	mBase = m.M
	v3281 = m.ExcPending
	if v3281 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	v3283 = v3218 + int32(1)
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v3206)+4))
	if v3283 < v3284 {
		v3215 = v3280
		v3216 = v3216 + int32(1)
		v3218 = v3283
		goto L542
	} else {
		goto L553
	}
L553:
	;
	goto L543
L554:
	;
	v3346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v3347 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v3349 = F_palloc0(m, int32(96))
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		goto L1
	} else {
		goto L564
	}
L555:
	;
	v3345 = int32(2147483647)
	goto L554
L556:
	;
	goto L557
L557:
	;
	if base.F64_le(v3330, float64(0)) != 0 {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v3345 = int32(0)
	goto L554
L559:
	;
	goto L560
L560:
	;
	v3340 = float64(2.147483647e+09)
	if base.F64_lt(v3330, v3340) != 0 {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v3343 = v3330
	goto L563
L562:
	;
	v3343 = v3340
	goto L563
L563:
	;
	v3345 = base.I32_trunc_sat_f64_s(v3343)
	goto L554
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3349))) = int32(336)
	if v3346 != 0 {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	v3353 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+4))
	v3354 = v3353
	goto L567
L566:
	;
	v3354 = v4
	goto L567
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+76)) = v3354
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+72)) = v3347
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+56)) = v3202
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+52)) = v3198
	v3359 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+48)) = v3359
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+44)) = v3288
	if v3359 < v3354 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v3366 = F_palloc(m, v3354<<(uint(int32(1))%32))
	mBase = m.M
	v3367 = m.ExcPending
	if v3367 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+92)) = v3345
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+4)) = v3541
	v3543 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3349)+8)) = v3543
	v3545 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3349)+16)) = v3545
	v3547 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3349)+24)) = v3547
	v3549 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3550 = *(*int32)(unsafe.Add(mBase, uint32(v3549)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+32)) = v3550
	v3552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3349)+36)) = uint8(v3552)
	v3554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3349)+37)) = uint8(v3554)
	v12573 = v3349
	v12580 = v47
	goto L3
L571:
	;
	v3369 = v3354 << (uint(int32(2)) % 32)
	v3370 = F_palloc(m, v3369)
	mBase = m.M
	v3371 = m.ExcPending
	if v3371 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	v3372 = F_palloc(m, v3369)
	mBase = m.M
	v3373 = m.ExcPending
	if v3373 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	if v3346 == int32(0) {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+88)) = v3372
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+84)) = v3370
	*(*int32)(unsafe.Add(mBase, uint32(v3349)+80)) = v3366
	goto L570
L575:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+4))
	if v3376 <= int32(0) {
		goto L574
	} else {
		goto L576
	}
L576:
	;
	v3380 = int32(0)
	goto L577
L577:
	;
	v3428 = v3380 << (uint(int32(2)) % 32)
	v3429 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+12))
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v3428+v3429)))
	v3432 = *(*int32)(unsafe.Add(mBase, uint32(v3349)+44))
	v3433 = F_get_sortgroupclause_tle(m, v3431, v3432)
	mBase = m.M
	v3434 = m.ExcPending
	if v3434 != 0 {
		goto L1
	} else {
		goto L579
	}
L578:
	;
	goto L574
L579:
	;
	v3435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3433)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3366+v3380<<(uint(int32(1))%32)))) = uint16(v3435)
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(v3431)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3370+v3428))) = v3438
	v3441 = *(*int32)(unsafe.Add(mBase, uint32(v3433)+4))
	v3442 = F_exprCollation(m, v3441)
	mBase = m.M
	v3443 = m.ExcPending
	if v3443 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3372+v3428))) = v3442
	v3446 = v3380 + int32(1)
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(v3346)+4))
	if v3446 < v3447 {
		v3380 = v3446
		goto L577
	} else {
		goto L581
	}
L581:
	;
	goto L578
L582:
	;
	v3682 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3684 = l2 | int32(4)
	v3685 = F_create_plan_recurse(m, l0, v3682, v3684)
	mBase = m.M
	v3686 = m.ExcPending
	if v3686 != 0 {
		goto L1
	} else {
		goto L597
	}
L583:
	;
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v3557)+4))
	if v3561 <= int32(0) {
		v3646 = v4
		goto L582
	} else {
		goto L584
	}
L584:
	;
	v3564 = *(*int32)(unsafe.Add(mBase, uint32(v3556)+8))
	v3568 = int32(1)
	v3570 = v4
	v3573 = v4
	goto L585
L585:
	;
	v3609 = *(*int32)(unsafe.Add(mBase, uint32(v3557)+12))
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v3609+v3570<<(uint(int32(2))%32))))
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3614 != 0 {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	v3646 = v3632
	goto L582
L587:
	;
	v3615 = F_replace_nestloop_params_mutator(m, v3613, l0)
	mBase = m.M
	v3616 = m.ExcPending
	if v3616 != 0 {
		goto L1
	} else {
		goto L590
	}
L588:
	;
	v3617 = v3613
	goto L589
L589:
	;
	v3619 = int32(0)
	v3621 = F_makeTargetEntry(m, v3617, base.I32_extend16_s(v3568), v3619, v3619)
	mBase = m.M
	v3622 = m.ExcPending
	if v3622 != 0 {
		goto L1
	} else {
		goto L591
	}
L590:
	;
	v3617 = v3615
	goto L589
L591:
	;
	if v3564 != 0 {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(v3564+v3568<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v3621)+16)) = v3628
	goto L594
L593:
	;
	goto L594
L594:
	;
	v3632 = F_lappend(m, v3573, v3621)
	mBase = m.M
	v3633 = m.ExcPending
	if v3633 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	v3635 = v3570 + int32(1)
	v3636 = *(*int32)(unsafe.Add(mBase, uint32(v3557)+4))
	if v3635 < v3636 {
		v3568 = v3568 + int32(1)
		v3570 = v3635
		v3573 = v3632
		goto L585
	} else {
		goto L596
	}
L596:
	;
	goto L586
L597:
	;
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v3688 = F_create_plan_recurse(m, l0, v3687, v3684)
	mBase = m.M
	v3689 = m.ExcPending
	if v3689 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	v3690 = *(*float64)(unsafe.Add(mBase, uint32(l1)+96))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v3690)&int64(9223372036854775807)) {
		goto L600
	} else {
		goto L601
	}
L599:
	;
	v3706 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v3707 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v3710 = F_palloc0(m, int32(104))
	mBase = m.M
	v3711 = m.ExcPending
	if v3711 != 0 {
		goto L1
	} else {
		goto L609
	}
L600:
	;
	v3705 = int32(2147483647)
	goto L599
L601:
	;
	goto L602
L602:
	;
	if base.F64_le(v3690, float64(0)) != 0 {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v3705 = int32(0)
	goto L599
L604:
	;
	goto L605
L605:
	;
	v3700 = float64(2.147483647e+09)
	if base.F64_lt(v3690, v3700) != 0 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v3703 = v3690
	goto L608
L607:
	;
	v3703 = v3700
	goto L608
L608:
	;
	v3705 = base.I32_trunc_sat_f64_s(v3703)
	goto L599
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3710))) = int32(371)
	if v3706 != 0 {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v3706)+4))
	v3715 = v3714
	goto L612
L611:
	;
	v3715 = v4
	goto L612
L612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+56)) = v3688
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+52)) = v3685
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+44)) = v3646
	v3723 = F_palloc(m, v3715<<(uint(int32(1))%32))
	mBase = m.M
	v3724 = m.ExcPending
	if v3724 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v3726 = v3715 << (uint(int32(2)) % 32)
	v3727 = F_palloc(m, v3726)
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	v3729 = F_palloc(m, v3726)
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	v3731 = F_palloc(m, v3715)
	mBase = m.M
	v3732 = m.ExcPending
	if v3732 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	if v3706 == int32(0) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+100)) = v3705
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+96)) = v3731
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+92)) = v3729
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+88)) = v3727
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+84)) = v3723
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+80)) = v3715
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+76)) = v3707
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+72)) = v3708
	v3869 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+4)) = v3869
	v3871 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v3710)+8)) = v3871
	v3873 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v3710)+16)) = v3873
	v3875 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v3710)+24)) = v3875
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3877)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3710)+32)) = v3878
	v3880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3710)+36)) = uint8(v3880)
	v3882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3710)+37)) = uint8(v3882)
	v12573 = v3710
	v12580 = v47
	goto L3
L618:
	;
	v3735 = *(*int32)(unsafe.Add(mBase, uint32(v3706)+4))
	if v3735 <= int32(0) {
		goto L617
	} else {
		goto L619
	}
L619:
	;
	if v3707 == int32(1) {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v3742 = int32(8)
	goto L622
L621:
	;
	v3742 = int32(12)
	goto L622
L622:
	;
	v3744 = int32(0)
	goto L623
L623:
	;
	v3792 = v3744 << (uint(int32(2)) % 32)
	v3793 = *(*int32)(unsafe.Add(mBase, uint32(v3706)+12))
	v3795 = *(*int32)(unsafe.Add(mBase, uint32(v3792+v3793)))
	v3796 = *(*int32)(unsafe.Add(mBase, uint32(v3710)+44))
	v3797 = F_get_sortgroupclause_tle(m, v3795, v3796)
	mBase = m.M
	v3798 = m.ExcPending
	if v3798 != 0 {
		goto L1
	} else {
		goto L625
	}
L624:
	;
	goto L617
L625:
	;
	v3799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3797)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3723+v3744<<(uint(int32(1))%32)))) = uint16(v3799)
	v3803 = *(*int32)(unsafe.Add(mBase, uint32(v3795+v3742)))
	*(*int32)(unsafe.Add(mBase, uint32(v3727+v3792))) = v3803
	v3806 = *(*int32)(unsafe.Add(mBase, uint32(v3797)+4))
	v3807 = F_exprCollation(m, v3806)
	mBase = m.M
	v3808 = m.ExcPending
	if v3808 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3729+v3792))) = v3807
	v3811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3795)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3744+v3731))) = uint8(v3811)
	v3814 = v3744 + int32(1)
	v3815 = *(*int32)(unsafe.Add(mBase, uint32(v3706)+4))
	if v3814 < v3815 {
		v3744 = v3814
		goto L623
	} else {
		goto L627
	}
L627:
	;
	goto L624
L628:
	;
	v3886 = *(*int32)(unsafe.Add(mBase, uint32(v3885)+4))
	v3887 = v3886
	goto L630
L629:
	;
	v3887 = v4
	goto L630
L630:
	;
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+16))
	if v3888 != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v3889 = *(*int32)(unsafe.Add(mBase, uint32(v3888)+4))
	v3890 = v3889
	goto L633
L632:
	;
	v3890 = v4
	goto L633
L633:
	;
	v3891 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v3893 = F_create_plan_recurse(m, l0, v3891, int32(6))
	mBase = m.M
	v3894 = m.ExcPending
	if v3894 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	v3895 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3896 = *(*int32)(unsafe.Add(mBase, uint32(v3895)+4))
	if v3896 == int32(0) {
		v3986 = v4
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v4024 = F_palloc(m, v3887<<(uint(int32(1))%32))
	mBase = m.M
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L1
	} else {
		goto L650
	}
L636:
	;
	v3900 = *(*int32)(unsafe.Add(mBase, uint32(v3896)+4))
	if v3900 <= int32(0) {
		v3986 = v4
		goto L635
	} else {
		goto L637
	}
L637:
	;
	v3903 = *(*int32)(unsafe.Add(mBase, uint32(v3895)+8))
	v3907 = int32(0)
	v3908 = int32(1)
	v3913 = v4
	goto L638
L638:
	;
	v3949 = *(*int32)(unsafe.Add(mBase, uint32(v3896)+12))
	v3953 = *(*int32)(unsafe.Add(mBase, uint32(v3949+v3907<<(uint(int32(2))%32))))
	v3954 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3954 != 0 {
		goto L640
	} else {
		goto L641
	}
L639:
	;
	v3986 = v3972
	goto L635
L640:
	;
	v3955 = F_replace_nestloop_params_mutator(m, v3953, l0)
	mBase = m.M
	v3956 = m.ExcPending
	if v3956 != 0 {
		goto L1
	} else {
		goto L643
	}
L641:
	;
	v3957 = v3953
	goto L642
L642:
	;
	v3959 = int32(0)
	v3961 = F_makeTargetEntry(m, v3957, base.I32_extend16_s(v3908), v3959, v3959)
	mBase = m.M
	v3962 = m.ExcPending
	if v3962 != 0 {
		goto L1
	} else {
		goto L644
	}
L643:
	;
	v3957 = v3955
	goto L642
L644:
	;
	if v3903 != 0 {
		goto L645
	} else {
		goto L646
	}
L645:
	;
	v3968 = *(*int32)(unsafe.Add(mBase, uint32(v3903+v3908<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v3961)+16)) = v3968
	goto L647
L646:
	;
	goto L647
L647:
	;
	v3972 = F_lappend(m, v3913, v3961)
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	v3975 = v3907 + int32(1)
	v3976 = *(*int32)(unsafe.Add(mBase, uint32(v3896)+4))
	if v3975 < v3976 {
		v3907 = v3975
		v3908 = v3908 + int32(1)
		v3913 = v3972
		goto L638
	} else {
		goto L649
	}
L649:
	;
	goto L639
L650:
	;
	v4027 = v3887 << (uint(int32(2)) % 32)
	v4028 = F_palloc(m, v4027)
	mBase = m.M
	v4029 = m.ExcPending
	if v4029 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	v4030 = F_palloc(m, v4027)
	mBase = m.M
	v4031 = m.ExcPending
	if v4031 != 0 {
		goto L1
	} else {
		goto L652
	}
L652:
	;
	v4032 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+12))
	if v4032 == int32(0) {
		v4111 = v4
		goto L653
	} else {
		goto L654
	}
L653:
	;
	v4153 = F_palloc(m, v3890<<(uint(int32(1))%32))
	mBase = m.M
	v4154 = m.ExcPending
	if v4154 != 0 {
		goto L1
	} else {
		goto L661
	}
L654:
	;
	v4035 = *(*int32)(unsafe.Add(mBase, uint32(v4032)+4))
	if v4035 <= int32(0) {
		v4111 = v4
		goto L653
	} else {
		goto L655
	}
L655:
	;
	v4042 = v4
	goto L656
L656:
	;
	v4086 = v4042 << (uint(int32(2)) % 32)
	v4087 = *(*int32)(unsafe.Add(mBase, uint32(v4032)+12))
	v4089 = *(*int32)(unsafe.Add(mBase, uint32(v4086+v4087)))
	v4090 = *(*int32)(unsafe.Add(mBase, uint32(v3893)+44))
	v4091 = F_get_sortgroupclause_tle(m, v4089, v4090)
	mBase = m.M
	v4092 = m.ExcPending
	if v4092 != 0 {
		goto L1
	} else {
		goto L658
	}
L657:
	;
	v4111 = v4104
	goto L653
L658:
	;
	v4093 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4091)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4024+v4042<<(uint(int32(1))%32)))) = uint16(v4093)
	v4096 = *(*int32)(unsafe.Add(mBase, uint32(v4089)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4086+v4028))) = v4096
	v4099 = *(*int32)(unsafe.Add(mBase, uint32(v4091)+4))
	v4100 = F_exprCollation(m, v4099)
	mBase = m.M
	v4101 = m.ExcPending
	if v4101 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4086+v4030))) = v4100
	v4104 = v4042 + int32(1)
	v4105 = *(*int32)(unsafe.Add(mBase, uint32(v4032)+4))
	if v4104 < v4105 {
		v4042 = v4104
		goto L656
	} else {
		goto L660
	}
L660:
	;
	goto L657
L661:
	;
	v4156 = v3890 << (uint(int32(2)) % 32)
	v4157 = F_palloc(m, v4156)
	mBase = m.M
	v4158 = m.ExcPending
	if v4158 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	v4159 = F_palloc(m, v4156)
	mBase = m.M
	v4160 = m.ExcPending
	if v4160 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	v4161 = int32(0)
	v4162 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+16))
	if v4162 == v4161 {
		v4237 = v4161
		goto L664
	} else {
		goto L665
	}
L664:
	;
	v4281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v4282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+88)))
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v4285 = F_palloc0(m, int32(152))
	mBase = m.M
	v4286 = m.ExcPending
	if v4286 != 0 {
		goto L1
	} else {
		goto L672
	}
L665:
	;
	v4165 = *(*int32)(unsafe.Add(mBase, uint32(v4162)+4))
	if v4165 <= int32(0) {
		v4237 = v4161
		goto L664
	} else {
		goto L666
	}
L666:
	;
	v4168 = v4161
	goto L667
L667:
	;
	v4216 = v4168 << (uint(int32(2)) % 32)
	v4217 = *(*int32)(unsafe.Add(mBase, uint32(v4162)+12))
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v4216+v4217)))
	v4220 = *(*int32)(unsafe.Add(mBase, uint32(v3893)+44))
	v4221 = F_get_sortgroupclause_tle(m, v4219, v4220)
	mBase = m.M
	v4222 = m.ExcPending
	if v4222 != 0 {
		goto L1
	} else {
		goto L669
	}
L668:
	;
	v4237 = v4234
	goto L664
L669:
	;
	v4223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4221)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4153+v4168<<(uint(int32(1))%32)))) = uint16(v4223)
	v4226 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4216+v4157))) = v4226
	v4229 = *(*int32)(unsafe.Add(mBase, uint32(v4221)+4))
	v4230 = F_exprCollation(m, v4229)
	mBase = m.M
	v4231 = m.ExcPending
	if v4231 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4216+v4159))) = v4230
	v4234 = v4168 + int32(1)
	v4235 = *(*int32)(unsafe.Add(mBase, uint32(v4162)+4))
	if v4234 < v4235 {
		v4168 = v4234
		goto L667
	} else {
		goto L671
	}
L671:
	;
	goto L668
L672:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4285))) = int32(366)
	v4289 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+72)) = v4289
	v4291 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+108)) = v4159
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+104)) = v4157
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+100)) = v4153
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+96)) = v4237
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+92)) = v4030
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+88)) = v4028
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+84)) = v4024
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+80)) = v4111
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+76)) = v4291
	v4301 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+112)) = v4301
	v4303 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+116)) = v4303
	v4305 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+128)) = v4283
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+124)) = v4283
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+120)) = v4305
	v4309 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+132)) = v4309
	v4311 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+136)) = v4311
	v4313 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+140)) = v4313
	v4315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3884)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4285)+144)) = uint8(v4315)
	v4317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3884)+45)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4285)+146)) = uint8(v4282)
	*(*uint8)(unsafe.Add(mBase, uint32(v4285)+145)) = uint8(v4317)
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+52)) = v3893
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+44)) = v3986
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+48)) = v4281
	v4325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+4)) = v4325
	v4327 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v4285)+8)) = v4327
	v4329 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v4285)+16)) = v4329
	v4331 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v4285)+24)) = v4331
	v4333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v4334 = *(*int32)(unsafe.Add(mBase, uint32(v4333)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4285)+32)) = v4334
	v4336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4285)+36)) = uint8(v4336)
	v4338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4285)+37)) = uint8(v4338)
	v12573 = v4285
	v12580 = v47
	goto L3
L673:
	;
	v4343 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v4346 = F_create_plan_recurse(m, l0, v4344, int32(4))
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		goto L1
	} else {
		goto L676
	}
L674:
	;
	goto L675
L675:
	;
	v5728 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v5730 = F_create_plan_recurse(m, l0, v5728, int32(4))
	mBase = m.M
	v5731 = m.ExcPending
	if v5731 != 0 {
		goto L1
	} else {
		goto L835
	}
L676:
	;
	v4348 = int32(2)
	v4349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v4349 == int32(0) {
		v4590 = v4348
		goto L677
	} else {
		goto L678
	}
L677:
	;
	v4632 = F_palloc0(m, v4590)
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		goto L1
	} else {
		goto L709
	}
L678:
	;
	v4352 = *(*int32)(unsafe.Add(mBase, uint32(v4349)+4))
	if v4352 <= int32(0) {
		v4590 = v4348
		goto L677
	} else {
		goto L679
	}
L679:
	;
	v4355 = int32(0)
	if v4355 < v4352 {
		goto L680
	} else {
		goto L681
	}
L680:
	;
	v4358 = v4352
	goto L682
L681:
	;
	v4358 = v4355
	goto L682
L682:
	;
	v4360 = v4358 & int32(3)
	v4361 = int32(0)
	if int32(4) <= v4352 {
		goto L684
	} else {
		goto L685
	}
L683:
	;
	v4590 = v4542<<(uint(int32(1))%32) + int32(2)
	goto L677
L684:
	;
	v4366 = *(*int32)(unsafe.Add(mBase, uint32(v4349)+12))
	v4370 = v4361
	v4372 = int32(0)
	v4373 = v4
	goto L687
L685:
	;
	v4440 = v4361
	v4443 = v4
	goto L686
L686:
	;
	v4482 = *(*int32)(unsafe.Add(mBase, uint32(v4349)+12))
	v4486 = v4440
	v4487 = int32(0)
	v4489 = v4443
	goto L703
L687:
	;
	v4414 = v4366 + v4373<<(uint(int32(2))%32)
	v4415 = *(*int32)(unsafe.Add(mBase, uint32(v4414)+12))
	v4416 = *(*int32)(unsafe.Add(mBase, uint32(v4415)+4))
	v4417 = *(*int32)(unsafe.Add(mBase, uint32(v4414)+8))
	v4418 = *(*int32)(unsafe.Add(mBase, uint32(v4417)+4))
	v4419 = *(*int32)(unsafe.Add(mBase, uint32(v4414)+4))
	v4420 = *(*int32)(unsafe.Add(mBase, uint32(v4419)+4))
	v4421 = *(*int32)(unsafe.Add(mBase, uint32(v4414)))
	v4422 = *(*int32)(unsafe.Add(mBase, uint32(v4421)+4))
	if base.Ui32(v4370) < base.Ui32(v4422) {
		goto L689
	} else {
		goto L690
	}
L688:
	;
	if v4360 == int32(0) {
		v4542 = v4430
		goto L683
	} else {
		goto L702
	}
L689:
	;
	v4424 = v4422
	goto L691
L690:
	;
	v4424 = v4370
	goto L691
L691:
	;
	if base.Ui32(v4424) < base.Ui32(v4420) {
		goto L692
	} else {
		goto L693
	}
L692:
	;
	v4426 = v4420
	goto L694
L693:
	;
	v4426 = v4424
	goto L694
L694:
	;
	if base.Ui32(v4426) < base.Ui32(v4418) {
		goto L695
	} else {
		goto L696
	}
L695:
	;
	v4428 = v4418
	goto L697
L696:
	;
	v4428 = v4426
	goto L697
L697:
	;
	if base.Ui32(v4428) < base.Ui32(v4416) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v4430 = v4416
	goto L700
L699:
	;
	v4430 = v4428
	goto L700
L700:
	;
	v4431 = int32(4)
	v4432 = v4373 + v4431
	v4434 = v4372 + v4431
	if v4434 != v4358&int32(2147483644) {
		v4370 = v4430
		v4372 = v4434
		v4373 = v4432
		goto L687
	} else {
		goto L701
	}
L701:
	;
	goto L688
L702:
	;
	v4440 = v4430
	v4443 = v4432
	goto L686
L703:
	;
	v4531 = *(*int32)(unsafe.Add(mBase, uint32(v4482+v4489<<(uint(int32(2))%32))))
	v4532 = *(*int32)(unsafe.Add(mBase, uint32(v4531)+4))
	if base.Ui32(v4486) < base.Ui32(v4532) {
		goto L705
	} else {
		goto L706
	}
L704:
	;
	v4542 = v4534
	goto L683
L705:
	;
	v4534 = v4532
	goto L707
L706:
	;
	v4534 = v4486
	goto L707
L707:
	;
	v4535 = int32(1)
	v4538 = v4487 + v4535
	if v4538 != v4360 {
		v4486 = v4534
		v4487 = v4538
		v4489 = v4489 + v4535
		goto L703
	} else {
		goto L708
	}
L708:
	;
	goto L704
L709:
	;
	v4634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v4634 == int32(0) {
		goto L710
	} else {
		goto L711
	}
L710:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v4632
	if v4343 == int32(0) {
		v5367 = v4632
		v5375 = v4
		goto L717
	} else {
		goto L718
	}
L711:
	;
	v4637 = *(*int32)(unsafe.Add(mBase, uint32(v4634)+4))
	if v4637 <= int32(0) {
		goto L710
	} else {
		goto L712
	}
L712:
	;
	v4643 = int32(0)
	goto L713
L713:
	;
	v4685 = *(*int32)(unsafe.Add(mBase, uint32(v4634)+12))
	v4689 = *(*int32)(unsafe.Add(mBase, uint32(v4685+v4643<<(uint(int32(2))%32))))
	v4690 = *(*int32)(unsafe.Add(mBase, uint32(v4346)+44))
	v4691 = F_get_sortgroupclause_tle(m, v4689, v4690)
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
		goto L1
	} else {
		goto L715
	}
L714:
	;
	goto L710
L715:
	;
	v4693 = *(*int32)(unsafe.Add(mBase, uint32(v4689)+4))
	v4694 = int32(1)
	v4697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4691)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4632+v4693<<(uint(v4694)%32)))) = uint16(v4697)
	v4700 = v4643 + v4694
	v4701 = *(*int32)(unsafe.Add(mBase, uint32(v4634)+4))
	if v4700 < v4701 {
		v4643 = v4700
		goto L713
	} else {
		goto L716
	}
L716:
	;
	goto L714
L717:
	;
	v5406 = *(*int32)(unsafe.Add(mBase, uint32(v4343)+12))
	v5407 = *(*int32)(unsafe.Add(mBase, uint32(v5406)))
	v5408 = *(*int32)(unsafe.Add(mBase, uint32(v5407)+4))
	if v5408 == int32(0) {
		goto L795
	} else {
		goto L796
	}
L718:
	;
	v4750 = *(*int32)(unsafe.Add(mBase, uint32(v4343)+4))
	if v4750 < int32(2) {
		v5367 = v4632
		v5375 = v4
		goto L717
	} else {
		goto L719
	}
L719:
	;
	v4753 = *(*int32)(unsafe.Add(mBase, uint32(v4343)+12))
	v4754 = *(*int32)(unsafe.Add(mBase, uint32(v4753)))
	v4755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4754)+25)))
	v4764 = int32(1)
	v4770 = v4
	v4772 = v4755
	goto L720
L720:
	;
	v4801 = *(*int32)(unsafe.Add(mBase, uint32(v4343)+12))
	v4805 = *(*int32)(unsafe.Add(mBase, uint32(v4801+v4764<<(uint(int32(2))%32))))
	v4806 = *(*int32)(unsafe.Add(mBase, uint32(v4805)+4))
	if v4806 == int32(0) {
		goto L723
	} else {
		goto L724
	}
L721:
	;
	v5361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v5367 = v5361
	v5375 = v5355
	goto L717
L722:
	;
	v4928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4805)+25)))
	if (v4928|v4772)&int32(1) == int32(0) {
		goto L736
	} else {
		goto L737
	}
L723:
	;
	v4810 = F_palloc0(m, int32(0))
	mBase = m.M
	v4811 = m.ExcPending
	if v4811 != 0 {
		goto L1
	} else {
		goto L726
	}
L724:
	;
	goto L725
L725:
	;
	v4812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v4813 = *(*int32)(unsafe.Add(mBase, uint32(v4806)+4))
	v4816 = F_palloc0(m, v4813<<(uint(int32(1))%32))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L1
	} else {
		goto L727
	}
L726:
	;
	v4887 = v4810
	goto L722
L727:
	;
	v4818 = int32(0)
	v4819 = *(*int32)(unsafe.Add(mBase, uint32(v4806)+4))
	if v4819 <= v4818 {
		v4887 = v4816
		goto L722
	} else {
		goto L728
	}
L728:
	;
	v4824 = v4818
	goto L729
L729:
	;
	v4866 = int32(1)
	v4869 = *(*int32)(unsafe.Add(mBase, uint32(v4806)+12))
	v4873 = *(*int32)(unsafe.Add(mBase, uint32(v4869+v4824<<(uint(int32(2))%32))))
	v4874 = *(*int32)(unsafe.Add(mBase, uint32(v4873)+4))
	v4878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4812+v4874<<(uint(v4866)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v4816+v4824<<(uint(v4866)%32)))) = uint16(v4878)
	v4881 = v4824 + v4866
	v4882 = *(*int32)(unsafe.Add(mBase, uint32(v4806)+4))
	if v4881 < v4882 {
		v4824 = v4881
		goto L729
	} else {
		goto L731
	}
L730:
	;
	v4887 = v4816
	goto L722
L731:
	;
	goto L730
L732:
	;
	if v5256 != 0 {
		goto L773
	} else {
		goto L774
	}
L733:
	;
	v5248 = int32(0)
	v5249 = *(*int32)(unsafe.Add(mBase, uint32(v4805)+8))
	v5250 = *(*int32)(unsafe.Add(mBase, uint32(v5249)+12))
	v5251 = *(*int32)(unsafe.Add(mBase, uint32(v5250)))
	v5256 = v5251
	v5259 = v5209
	v5269 = v5248
	v5298 = base.B2i32(v5251 != v5248)
	goto L732
L734:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5194 = m.ExcPending
	if v5194 != 0 {
		goto L1
	} else {
		goto L770
	}
L735:
	;
	v5187 = *(*int32)(unsafe.Add(mBase, uint32(v4805)+8))
	v5188 = *(*int32)(unsafe.Add(mBase, uint32(v5187)+12))
	v5189 = *(*int32)(unsafe.Add(mBase, uint32(v5188)))
	v5256 = v5189
	v5259 = v5148
	v5269 = v4772
	v5298 = int32(2)
	goto L732
L736:
	;
	v4934 = int32(0)
	v4936 = *(*int32)(unsafe.Add(mBase, uint32(v4346)+44))
	v4937 = *(*int32)(unsafe.Add(mBase, uint32(v4805)+4))
	if v4937 != 0 {
		goto L739
	} else {
		goto L740
	}
L737:
	;
	goto L738
L738:
	;
	v5138 = int32(0)
	if v4928&int32(1) == v5138 {
		v5209 = v5138
		goto L733
	} else {
		goto L769
	}
L739:
	;
	v4938 = *(*int32)(unsafe.Add(mBase, uint32(v4937)+4))
	v4939 = v4938
	goto L741
L740:
	;
	v4939 = v4934
	goto L741
L741:
	;
	v4942 = F_palloc(m, v4939<<(uint(int32(1))%32))
	mBase = m.M
	v4943 = m.ExcPending
	if v4943 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	v4945 = v4939 << (uint(int32(2)) % 32)
	v4946 = F_palloc(m, v4945)
	mBase = m.M
	v4947 = m.ExcPending
	if v4947 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	v4948 = F_palloc(m, v4945)
	mBase = m.M
	v4949 = m.ExcPending
	if v4949 != 0 {
		goto L1
	} else {
		goto L744
	}
L744:
	;
	v4950 = F_palloc(m, v4939)
	mBase = m.M
	v4951 = m.ExcPending
	if v4951 != 0 {
		goto L1
	} else {
		goto L745
	}
L745:
	;
	if v4937 == int32(0) {
		v5071 = v4934
		goto L746
	} else {
		goto L747
	}
L746:
	;
	v5114 = F_palloc0(m, int32(96))
	mBase = m.M
	v5115 = m.ExcPending
	if v5115 != 0 {
		goto L1
	} else {
		goto L767
	}
L747:
	;
	v4954 = *(*int32)(unsafe.Add(mBase, uint32(v4937)+4))
	if v4954 <= int32(0) {
		v5071 = v4934
		goto L746
	} else {
		goto L748
	}
L748:
	;
	v4959 = v4934
	goto L749
L749:
	;
	v5002 = v4959 << (uint(int32(2)) % 32)
	v5003 = *(*int32)(unsafe.Add(mBase, uint32(v4937)+12))
	v5005 = *(*int32)(unsafe.Add(mBase, uint32(v5002+v5003)))
	v5007 = v4959 << (uint(int32(1)) % 32)
	v5009 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4887+v5007))))
	if v4936 != 0 {
		goto L753
	} else {
		goto L754
	}
L750:
	;
	v5071 = v5066
	goto L746
L751:
	;
	if v5047 == int32(0) {
		goto L734
	} else {
		goto L764
	}
L752:
	;
	goto L751
L753:
	;
	v5013 = *(*int32)(unsafe.Add(mBase, uint32(v4936)+4))
	if v5013 <= int32(0) {
		v5047 = int32(0)
		goto L752
	} else {
		goto L756
	}
L754:
	;
	goto L755
L755:
	;
	v5047 = int32(0)
	goto L752
L756:
	;
	v5016 = int32(0)
	if v5016 < v5013 {
		goto L757
	} else {
		goto L758
	}
L757:
	;
	v5019 = v5013
	goto L759
L758:
	;
	v5019 = v5016
	goto L759
L759:
	;
	v5020 = *(*int32)(unsafe.Add(mBase, uint32(v4936)+12))
	v5024 = int32(0)
	goto L760
L760:
	;
	v5032 = *(*int32)(unsafe.Add(mBase, uint32(v5020+v5024<<(uint(int32(2))%32))))
	v5033 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5032)+8)))
	if v5033 == v5009&int32(_a_F_create_plan_recurse_28) {
		v5047 = v5032
		goto L752
	} else {
		goto L762
	}
L761:
	;
	goto L755
L762:
	;
	v5036 = v5024 + int32(1)
	if v5036 != v5019 {
		v5024 = v5036
		goto L760
	} else {
		goto L763
	}
L763:
	;
	goto L761
L764:
	;
	v5052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5047)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v4942+v5007))) = uint16(v5052)
	v5055 = *(*int32)(unsafe.Add(mBase, uint32(v5005)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5002+v4946))) = v5055
	v5058 = *(*int32)(unsafe.Add(mBase, uint32(v5047)+4))
	v5059 = F_exprCollation(m, v5058)
	mBase = m.M
	v5060 = m.ExcPending
	if v5060 != 0 {
		goto L1
	} else {
		goto L765
	}
L765:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5002+v4948))) = v5059
	v5063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5005)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4959+v4950))) = uint8(v5063)
	v5066 = v4959 + int32(1)
	v5067 = *(*int32)(unsafe.Add(mBase, uint32(v4937)+4))
	if v5066 < v5067 {
		v4959 = v5066
		goto L749
	} else {
		goto L766
	}
L766:
	;
	goto L750
L767:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5114))) = int32(362)
	v5118 = *(*int32)(unsafe.Add(mBase, uint32(v4346)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+44)) = v5118
	v5120 = *(*int32)(unsafe.Add(mBase, uint32(v4346)+4))
	v5122 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+88)) = v4950
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+84)) = v4948
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+80)) = v4946
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+76)) = v4942
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+72)) = v5071
	v5128 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+56)) = v5128
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+52)) = v4346
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+48)) = v5128
	*(*int32)(unsafe.Add(mBase, uint32(v5114)+4)) = v5120 + (v5122 ^ int32(1))
	v5137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4805)+25)))
	if v5137 != 0 {
		v5148 = v5114
		goto L735
	} else {
		goto L768
	}
L768:
	;
	v5209 = v5114
	goto L733
L769:
	;
	v5148 = v5138
	goto L735
L770:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_29), int32(0))
	mBase = m.M
	v5198 = m.ExcPending
	if v5198 != 0 {
		goto L1
	} else {
		goto L771
	}
L771:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_30), int32(_a_F_create_plan_recurse_31))
	mBase = m.M
	v5203 = m.ExcPending
	if v5203 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L773:
	;
	v5299 = *(*int32)(unsafe.Add(mBase, uint32(v5256)+4))
	v5301 = v5299
	goto L775
L774:
	;
	v5301 = int32(0)
	goto L775
L775:
	;
	v5302 = *(*int32)(unsafe.Add(mBase, uint32(v4805)+4))
	v5303 = F_extract_grouping_ops(m, v5302)
	mBase = m.M
	v5304 = m.ExcPending
	if v5304 != 0 {
		goto L1
	} else {
		goto L776
	}
L776:
	;
	v5305 = *(*int32)(unsafe.Add(mBase, uint32(v4805)+4))
	v5306 = *(*int32)(unsafe.Add(mBase, uint32(v4346)+44))
	v5307 = F_extract_grouping_collations(m, v5305, v5306)
	mBase = m.M
	v5308 = m.ExcPending
	if v5308 != 0 {
		goto L1
	} else {
		goto L777
	}
L777:
	;
	v5309 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+88)))
	v5310 = *(*int32)(unsafe.Add(mBase, uint32(v4805)+8))
	v5311 = *(*float64)(unsafe.Add(mBase, uint32(v4805)+16))
	v5313 = F_palloc0(m, int32(128))
	mBase = m.M
	v5314 = m.ExcPending
	if v5314 != 0 {
		goto L1
	} else {
		goto L778
	}
L778:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5313))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5311)&int64(9223372036854775807)) {
		goto L780
	} else {
		goto L781
	}
L779:
	;
	v5332 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+120)) = v5332
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+116)) = v5310
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+112)) = v5332
	*(*int64)(unsafe.Add(mBase, uint32(v5313)+104)) = v5309
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+96)) = v5331
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+92)) = v5307
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+88)) = v5303
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+84)) = v4887
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+80)) = v5301
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+76)) = v5332
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+72)) = v5298
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+56)) = v5332
	*(*int32)(unsafe.Add(mBase, uint32(v5313)+52)) = v5259
	*(*int64)(unsafe.Add(mBase, uint32(v5313)+44)) = int64(0)
	if v5259 != 0 {
		goto L789
	} else {
		goto L790
	}
L780:
	;
	v5331 = int32(2147483647)
	goto L779
L781:
	;
	goto L782
L782:
	;
	if base.F64_le(v5311, float64(0)) != 0 {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	v5331 = int32(0)
	goto L779
L784:
	;
	goto L785
L785:
	;
	v5326 = float64(2.147483647e+09)
	if base.F64_lt(v5311, v5326) != 0 {
		goto L786
	} else {
		goto L787
	}
L786:
	;
	v5329 = v5311
	goto L788
L787:
	;
	v5329 = v5326
	goto L788
L788:
	;
	v5331 = base.I32_trunc_sat_f64_s(v5329)
	goto L779
L789:
	;
	v5351 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5259)+52)) = v5351
	*(*int32)(unsafe.Add(mBase, uint32(v5259)+44)) = v5351
	goto L791
L790:
	;
	goto L791
L791:
	;
	v5355 = F_lappend(m, v4770, v5313)
	mBase = m.M
	v5356 = m.ExcPending
	if v5356 != 0 {
		goto L1
	} else {
		goto L792
	}
L792:
	;
	v5358 = v4764 + int32(1)
	v5359 = *(*int32)(unsafe.Add(mBase, uint32(v4343)+4))
	if v5358 < v5359 {
		v4764 = v5358
		v4770 = v5355
		v4772 = v5269
		goto L720
	} else {
		goto L793
	}
L793:
	;
	goto L721
L794:
	;
	v5529 = int32(0)
	v5531 = *(*int32)(unsafe.Add(mBase, uint32(v5407)+8))
	v5532 = *(*int32)(unsafe.Add(mBase, uint32(v5531)+12))
	v5533 = *(*int32)(unsafe.Add(mBase, uint32(v5532)))
	if v5533 != 0 {
		goto L804
	} else {
		goto L805
	}
L795:
	;
	v5412 = F_palloc0(m, int32(0))
	mBase = m.M
	v5413 = m.ExcPending
	if v5413 != 0 {
		goto L1
	} else {
		goto L798
	}
L796:
	;
	goto L797
L797:
	;
	v5414 = *(*int32)(unsafe.Add(mBase, uint32(v5408)+4))
	v5417 = F_palloc0(m, v5414<<(uint(int32(1))%32))
	mBase = m.M
	v5418 = m.ExcPending
	if v5418 != 0 {
		goto L1
	} else {
		goto L799
	}
L798:
	;
	v5494 = v5412
	goto L794
L799:
	;
	v5419 = *(*int32)(unsafe.Add(mBase, uint32(v5408)+4))
	if v5419 <= int32(0) {
		v5494 = v5417
		goto L794
	} else {
		goto L800
	}
L800:
	;
	v5425 = int32(0)
	goto L801
L801:
	;
	v5467 = int32(1)
	v5470 = *(*int32)(unsafe.Add(mBase, uint32(v5408)+12))
	v5474 = *(*int32)(unsafe.Add(mBase, uint32(v5470+v5425<<(uint(int32(2))%32))))
	v5475 = *(*int32)(unsafe.Add(mBase, uint32(v5474)+4))
	v5479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5367+v5475<<(uint(v5467)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5417+v5425<<(uint(v5467)%32)))) = uint16(v5479)
	v5482 = v5425 + v5467
	v5483 = *(*int32)(unsafe.Add(mBase, uint32(v5408)+4))
	if v5482 < v5483 {
		v5425 = v5482
		goto L801
	} else {
		goto L803
	}
L802:
	;
	v5494 = v5417
	goto L794
L803:
	;
	goto L802
L804:
	;
	v5534 = *(*int32)(unsafe.Add(mBase, uint32(v5533)+4))
	v5535 = v5534
	goto L806
L805:
	;
	v5535 = v5529
	goto L806
L806:
	;
	v5536 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5537 = *(*int32)(unsafe.Add(mBase, uint32(v5536)+4))
	if v5537 == int32(0) {
		v5623 = v5529
		goto L807
	} else {
		goto L808
	}
L807:
	;
	v5663 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v5664 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v5665 = *(*int32)(unsafe.Add(mBase, uint32(v5407)+4))
	v5666 = F_extract_grouping_ops(m, v5665)
	mBase = m.M
	v5667 = m.ExcPending
	if v5667 != 0 {
		goto L1
	} else {
		goto L822
	}
L808:
	;
	v5541 = *(*int32)(unsafe.Add(mBase, uint32(v5537)+4))
	if v5541 <= int32(0) {
		v5623 = v5529
		goto L807
	} else {
		goto L809
	}
L809:
	;
	v5544 = *(*int32)(unsafe.Add(mBase, uint32(v5536)+8))
	v5548 = int32(1)
	v5549 = int32(0)
	v5550 = v5529
	goto L810
L810:
	;
	v5590 = *(*int32)(unsafe.Add(mBase, uint32(v5537)+12))
	v5594 = *(*int32)(unsafe.Add(mBase, uint32(v5590+v5549<<(uint(int32(2))%32))))
	v5595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5595 != 0 {
		goto L812
	} else {
		goto L813
	}
L811:
	;
	v5623 = v5613
	goto L807
L812:
	;
	v5596 = F_replace_nestloop_params_mutator(m, v5594, l0)
	mBase = m.M
	v5597 = m.ExcPending
	if v5597 != 0 {
		goto L1
	} else {
		goto L815
	}
L813:
	;
	v5598 = v5594
	goto L814
L814:
	;
	v5600 = int32(0)
	v5602 = F_makeTargetEntry(m, v5598, base.I32_extend16_s(v5548), v5600, v5600)
	mBase = m.M
	v5603 = m.ExcPending
	if v5603 != 0 {
		goto L1
	} else {
		goto L816
	}
L815:
	;
	v5598 = v5596
	goto L814
L816:
	;
	if v5544 != 0 {
		goto L817
	} else {
		goto L818
	}
L817:
	;
	v5609 = *(*int32)(unsafe.Add(mBase, uint32(v5544+v5548<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v5602)+16)) = v5609
	goto L819
L818:
	;
	goto L819
L819:
	;
	v5613 = F_lappend(m, v5550, v5602)
	mBase = m.M
	v5614 = m.ExcPending
	if v5614 != 0 {
		goto L1
	} else {
		goto L820
	}
L820:
	;
	v5616 = v5549 + int32(1)
	v5617 = *(*int32)(unsafe.Add(mBase, uint32(v5537)+4))
	if v5616 < v5617 {
		v5548 = v5548 + int32(1)
		v5549 = v5616
		v5550 = v5613
		goto L810
	} else {
		goto L821
	}
L821:
	;
	goto L811
L822:
	;
	v5668 = *(*int32)(unsafe.Add(mBase, uint32(v5407)+4))
	v5669 = *(*int32)(unsafe.Add(mBase, uint32(v4346)+44))
	v5670 = F_extract_grouping_collations(m, v5668, v5669)
	mBase = m.M
	v5671 = m.ExcPending
	if v5671 != 0 {
		goto L1
	} else {
		goto L823
	}
L823:
	;
	v5672 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+88)))
	v5673 = *(*int32)(unsafe.Add(mBase, uint32(v5407)+8))
	v5674 = *(*float64)(unsafe.Add(mBase, uint32(v5407)+16))
	v5676 = F_palloc0(m, int32(128))
	mBase = m.M
	v5677 = m.ExcPending
	if v5677 != 0 {
		goto L1
	} else {
		goto L824
	}
L824:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5676))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5674)&int64(9223372036854775807)) {
		goto L826
	} else {
		goto L827
	}
L825:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+120)) = v5375
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+116)) = v5673
	v5697 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+112)) = v5697
	*(*int64)(unsafe.Add(mBase, uint32(v5676)+104)) = v5672
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+96)) = v5694
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+92)) = v5670
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+88)) = v5666
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+84)) = v5494
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+80)) = v5535
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+76)) = v5697
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+72)) = v5664
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+48)) = v5663
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+56)) = v5697
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+52)) = v4346
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+44)) = v5623
	v5713 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+4)) = v5713
	v5715 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v5676)+8)) = v5715
	v5717 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v5676)+16)) = v5717
	v5719 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v5676)+24)) = v5719
	v5721 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5722 = *(*int32)(unsafe.Add(mBase, uint32(v5721)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5676)+32)) = v5722
	v5724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5676)+36)) = uint8(v5724)
	v5726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5676)+37)) = uint8(v5726)
	v12573 = v5676
	v12580 = v47
	goto L3
L826:
	;
	v5694 = int32(2147483647)
	goto L825
L827:
	;
	goto L828
L828:
	;
	if base.F64_le(v5674, float64(0)) != 0 {
		goto L829
	} else {
		goto L830
	}
L829:
	;
	v5694 = int32(0)
	goto L825
L830:
	;
	goto L831
L831:
	;
	v5689 = float64(2.147483647e+09)
	if base.F64_lt(v5674, v5689) != 0 {
		goto L832
	} else {
		goto L833
	}
L832:
	;
	v5692 = v5674
	goto L834
L833:
	;
	v5692 = v5689
	goto L834
L834:
	;
	v5694 = base.I32_trunc_sat_f64_s(v5692)
	goto L825
L835:
	;
	v5732 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5733 = *(*int32)(unsafe.Add(mBase, uint32(v5732)+4))
	if v5733 == int32(0) {
		v5820 = v4
		goto L836
	} else {
		goto L837
	}
L836:
	;
	v5858 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v5859 = F_order_qual_clauses(m, l0, v5858)
	mBase = m.M
	v5860 = m.ExcPending
	if v5860 != 0 {
		goto L1
	} else {
		goto L851
	}
L837:
	;
	v5737 = *(*int32)(unsafe.Add(mBase, uint32(v5733)+4))
	if v5737 <= int32(0) {
		v5820 = v4
		goto L836
	} else {
		goto L838
	}
L838:
	;
	v5740 = *(*int32)(unsafe.Add(mBase, uint32(v5732)+8))
	v5743 = int32(1)
	v5745 = v4
	v5747 = v4
	goto L839
L839:
	;
	v5785 = *(*int32)(unsafe.Add(mBase, uint32(v5733)+12))
	v5789 = *(*int32)(unsafe.Add(mBase, uint32(v5785+v5745<<(uint(int32(2))%32))))
	v5790 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5790 != 0 {
		goto L841
	} else {
		goto L842
	}
L840:
	;
	v5820 = v5808
	goto L836
L841:
	;
	v5791 = F_replace_nestloop_params_mutator(m, v5789, l0)
	mBase = m.M
	v5792 = m.ExcPending
	if v5792 != 0 {
		goto L1
	} else {
		goto L844
	}
L842:
	;
	v5793 = v5789
	goto L843
L843:
	;
	v5795 = int32(0)
	v5797 = F_makeTargetEntry(m, v5793, base.I32_extend16_s(v5743), v5795, v5795)
	mBase = m.M
	v5798 = m.ExcPending
	if v5798 != 0 {
		goto L1
	} else {
		goto L845
	}
L844:
	;
	v5793 = v5791
	goto L843
L845:
	;
	if v5740 != 0 {
		goto L846
	} else {
		goto L847
	}
L846:
	;
	v5804 = *(*int32)(unsafe.Add(mBase, uint32(v5740+v5743<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v5797)+16)) = v5804
	goto L848
L847:
	;
	goto L848
L848:
	;
	v5808 = F_lappend(m, v5747, v5797)
	mBase = m.M
	v5809 = m.ExcPending
	if v5809 != 0 {
		goto L1
	} else {
		goto L849
	}
L849:
	;
	v5811 = v5745 + int32(1)
	v5812 = *(*int32)(unsafe.Add(mBase, uint32(v5733)+4))
	if v5811 < v5812 {
		v5743 = v5743 + int32(1)
		v5745 = v5811
		v5747 = v5808
		goto L839
	} else {
		goto L850
	}
L850:
	;
	goto L840
L851:
	;
	v5861 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v5862 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v5863 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v5863 != 0 {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	v5864 = *(*int32)(unsafe.Add(mBase, uint32(v5863)+4))
	v5865 = v5864
	goto L854
L853:
	;
	v5865 = v4
	goto L854
L854:
	;
	v5866 = *(*int32)(unsafe.Add(mBase, uint32(v5730)+44))
	v5867 = F_extract_grouping_cols(m, v5863, v5866)
	mBase = m.M
	v5868 = m.ExcPending
	if v5868 != 0 {
		goto L1
	} else {
		goto L855
	}
L855:
	;
	v5869 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v5870 = F_extract_grouping_ops(m, v5869)
	mBase = m.M
	v5871 = m.ExcPending
	if v5871 != 0 {
		goto L1
	} else {
		goto L856
	}
L856:
	;
	v5872 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v5873 = *(*int32)(unsafe.Add(mBase, uint32(v5730)+44))
	v5874 = F_extract_grouping_collations(m, v5872, v5873)
	mBase = m.M
	v5875 = m.ExcPending
	if v5875 != 0 {
		goto L1
	} else {
		goto L857
	}
L857:
	;
	v5876 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+96)))
	v5877 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	v5879 = F_palloc0(m, int32(128))
	mBase = m.M
	v5880 = m.ExcPending
	if v5880 != 0 {
		goto L1
	} else {
		goto L858
	}
L858:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5879))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5877)&int64(9223372036854775807)) {
		goto L860
	} else {
		goto L861
	}
L859:
	;
	v5898 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+120)) = v5898
	*(*int64)(unsafe.Add(mBase, uint32(v5879)+112)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5879)+104)) = v5876
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+96)) = v5897
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+92)) = v5874
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+88)) = v5870
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+84)) = v5867
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+80)) = v5865
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+76)) = v5861
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+72)) = v5862
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+48)) = v5859
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+56)) = v5898
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+52)) = v5730
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+44)) = v5820
	v5915 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+4)) = v5915
	v5917 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v5879)+8)) = v5917
	v5919 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v5879)+16)) = v5919
	v5921 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v5879)+24)) = v5921
	v5923 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5924 = *(*int32)(unsafe.Add(mBase, uint32(v5923)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v5879)+32)) = v5924
	v5926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5879)+36)) = uint8(v5926)
	v5928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5879)+37)) = uint8(v5928)
	v12573 = v5879
	v12580 = v47
	goto L3
L860:
	;
	v5897 = int32(2147483647)
	goto L859
L861:
	;
	goto L862
L862:
	;
	if base.F64_le(v5877, float64(0)) != 0 {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	v5897 = int32(0)
	goto L859
L864:
	;
	goto L865
L865:
	;
	v5892 = float64(2.147483647e+09)
	if base.F64_lt(v5877, v5892) != 0 {
		goto L866
	} else {
		goto L867
	}
L866:
	;
	v5895 = v5877
	goto L868
L867:
	;
	v5895 = v5892
	goto L868
L868:
	;
	v5897 = base.I32_trunc_sat_f64_s(v5895)
	goto L859
L869:
	;
	v5934 = int32(0)
	v5935 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v5936 = *(*int32)(unsafe.Add(mBase, uint32(v5935)+4))
	if v5936 == v5934 {
		v6019 = v5934
		goto L870
	} else {
		goto L871
	}
L870:
	;
	v6061 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6062 = F_order_qual_clauses(m, l0, v6061)
	mBase = m.M
	v6063 = m.ExcPending
	if v6063 != 0 {
		goto L1
	} else {
		goto L885
	}
L871:
	;
	v5940 = *(*int32)(unsafe.Add(mBase, uint32(v5936)+4))
	if v5940 <= int32(0) {
		v6019 = v5934
		goto L870
	} else {
		goto L872
	}
L872:
	;
	v5943 = *(*int32)(unsafe.Add(mBase, uint32(v5935)+8))
	v5946 = v5934
	v5947 = int32(1)
	v5949 = v4
	goto L873
L873:
	;
	v5988 = *(*int32)(unsafe.Add(mBase, uint32(v5936)+12))
	v5992 = *(*int32)(unsafe.Add(mBase, uint32(v5988+v5949<<(uint(int32(2))%32))))
	v5993 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5993 != 0 {
		goto L875
	} else {
		goto L876
	}
L874:
	;
	v6019 = v6011
	goto L870
L875:
	;
	v5994 = F_replace_nestloop_params_mutator(m, v5992, l0)
	mBase = m.M
	v5995 = m.ExcPending
	if v5995 != 0 {
		goto L1
	} else {
		goto L878
	}
L876:
	;
	v5996 = v5992
	goto L877
L877:
	;
	v5998 = int32(0)
	v6000 = F_makeTargetEntry(m, v5996, base.I32_extend16_s(v5947), v5998, v5998)
	mBase = m.M
	v6001 = m.ExcPending
	if v6001 != 0 {
		goto L1
	} else {
		goto L879
	}
L878:
	;
	v5996 = v5994
	goto L877
L879:
	;
	if v5943 != 0 {
		goto L880
	} else {
		goto L881
	}
L880:
	;
	v6007 = *(*int32)(unsafe.Add(mBase, uint32(v5943+v5947<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v6000)+16)) = v6007
	goto L882
L881:
	;
	goto L882
L882:
	;
	v6011 = F_lappend(m, v5946, v6000)
	mBase = m.M
	v6012 = m.ExcPending
	if v6012 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	v6014 = v5949 + int32(1)
	v6015 = *(*int32)(unsafe.Add(mBase, uint32(v5936)+4))
	if v6014 < v6015 {
		v5946 = v6011
		v5947 = v5947 + int32(1)
		v5949 = v6014
		goto L873
	} else {
		goto L884
	}
L884:
	;
	goto L874
L885:
	;
	v6064 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v6064 != 0 {
		goto L886
	} else {
		goto L887
	}
L886:
	;
	v6065 = *(*int32)(unsafe.Add(mBase, uint32(v6064)+4))
	v6066 = v6065
	goto L888
L887:
	;
	v6066 = v4
	goto L888
L888:
	;
	v6067 = *(*int32)(unsafe.Add(mBase, uint32(v5932)+44))
	v6068 = F_extract_grouping_cols(m, v6064, v6067)
	mBase = m.M
	v6069 = m.ExcPending
	if v6069 != 0 {
		goto L1
	} else {
		goto L889
	}
L889:
	;
	v6070 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6071 = F_extract_grouping_ops(m, v6070)
	mBase = m.M
	v6072 = m.ExcPending
	if v6072 != 0 {
		goto L1
	} else {
		goto L890
	}
L890:
	;
	v6073 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6074 = *(*int32)(unsafe.Add(mBase, uint32(v5932)+44))
	v6075 = F_extract_grouping_collations(m, v6073, v6074)
	mBase = m.M
	v6076 = m.ExcPending
	if v6076 != 0 {
		goto L1
	} else {
		goto L891
	}
L891:
	;
	v6078 = F_palloc0(m, int32(88))
	mBase = m.M
	v6079 = m.ExcPending
	if v6079 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+84)) = v6075
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+80)) = v6071
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+76)) = v6068
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+72)) = v6066
	*(*int32)(unsafe.Add(mBase, uint32(v6078))) = int32(364)
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+48)) = v6062
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+52)) = v5932
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+44)) = v6019
	v6091 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+4)) = v6091
	v6093 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6078)+8)) = v6093
	v6095 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6078)+16)) = v6095
	v6097 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6078)+24)) = v6097
	v6099 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6100 = *(*int32)(unsafe.Add(mBase, uint32(v6099)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6078)+32)) = v6100
	v6102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6078)+36)) = uint8(v6102)
	v6104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6078)+37)) = uint8(v6104)
	v12573 = v6078
	v12580 = v47
	goto L3
L893:
	;
	v6111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v6114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6115 = *(*int32)(unsafe.Add(mBase, uint32(v6114)+8))
	v6116 = *(*int32)(unsafe.Add(mBase, uint32(v6115)+4))
	if base.Ui32(int32(5)) < base.Ui32(v6116) {
		v6128 = int32(0)
		goto L894
	} else {
		goto L895
	}
L894:
	;
	v6129 = int32(0)
	v6141 = F_prepare_sort_from_pathkeys(m, v6109, v6112, v6128, v6129, v6129, v47+int32(56), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v6142 = m.ExcPending
	if v6142 != 0 {
		goto L1
	} else {
		goto L897
	}
L895:
	;
	v6119 = int32(0)
	if int32(1)<<(uint(v6116)%32)&int32(44) == v6119 {
		v6128 = v6119
		goto L894
	} else {
		goto L896
	}
L896:
	;
	v6126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6127 = *(*int32)(unsafe.Add(mBase, uint32(v6126)+8))
	v6128 = v6127
	goto L894
L897:
	;
	v6143 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v6144 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v6145 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v6146 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v6147 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v6149 = F_palloc0(m, int32(104))
	mBase = m.M
	v6150 = m.ExcPending
	if v6150 != 0 {
		goto L1
	} else {
		goto L898
	}
L898:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6149))) = int32(363)
	v6153 = *(*int32)(unsafe.Add(mBase, uint32(v6141)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+96)) = v6111
	v6155 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+56)) = v6155
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+52)) = v6141
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+48)) = v6155
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+44)) = v6153
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+88)) = v6147
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+84)) = v6146
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+80)) = v6145
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+76)) = v6144
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+72)) = v6143
	v6166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+4)) = v6166
	v6168 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6149)+8)) = v6168
	v6170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6149)+16)) = v6170
	v6172 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6149)+24)) = v6172
	v6174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6175 = *(*int32)(unsafe.Add(mBase, uint32(v6174)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6149)+32)) = v6175
	v6177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6149)+36)) = uint8(v6177)
	v6179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6149)+37)) = uint8(v6179)
	v12573 = v6149
	v12580 = v47
	goto L3
L899:
	;
	v6186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v6188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6189 = *(*int32)(unsafe.Add(mBase, uint32(v6188)+8))
	v6190 = *(*int32)(unsafe.Add(mBase, uint32(v6189)+4))
	if base.Ui32(int32(5)) < base.Ui32(v6190) {
		v6202 = int32(0)
		goto L900
	} else {
		goto L901
	}
L900:
	;
	v6203 = int32(0)
	v6215 = F_prepare_sort_from_pathkeys(m, v6184, v6186, v6202, v6203, v6203, v47+int32(56), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128))
	mBase = m.M
	v6216 = m.ExcPending
	if v6216 != 0 {
		goto L1
	} else {
		goto L903
	}
L901:
	;
	v6193 = int32(0)
	if int32(1)<<(uint(v6190)%32)&int32(44) == v6193 {
		v6202 = v6193
		goto L900
	} else {
		goto L902
	}
L902:
	;
	v6200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6201 = *(*int32)(unsafe.Add(mBase, uint32(v6200)+8))
	v6202 = v6201
	goto L900
L903:
	;
	v6217 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v6218 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v6219 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v6220 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v6221 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v6223 = F_palloc0(m, int32(96))
	mBase = m.M
	v6224 = m.ExcPending
	if v6224 != 0 {
		goto L1
	} else {
		goto L904
	}
L904:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6223))) = int32(362)
	v6227 = *(*int32)(unsafe.Add(mBase, uint32(v6215)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+44)) = v6227
	v6229 = *(*int32)(unsafe.Add(mBase, uint32(v6215)+4))
	v6231 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+88)) = v6221
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+84)) = v6220
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+80)) = v6219
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+76)) = v6218
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+72)) = v6217
	v6237 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+56)) = v6237
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+52)) = v6215
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+48)) = v6237
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+4)) = v6229 + (v6231 ^ int32(1))
	v6246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+4)) = v6246
	v6248 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6223)+8)) = v6248
	v6250 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6223)+16)) = v6250
	v6252 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6223)+24)) = v6252
	v6254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6255 = *(*int32)(unsafe.Add(mBase, uint32(v6254)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6223)+32)) = v6255
	v6257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6223)+36)) = uint8(v6257)
	v6259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6223)+37)) = uint8(v6259)
	v12573 = v6223
	v12580 = v47
	goto L3
L905:
	;
	v6266 = int32(0)
	v6267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6268 = *(*int32)(unsafe.Add(mBase, uint32(v6267)+4))
	if v6268 == v6266 {
		v6350 = v6266
		goto L906
	} else {
		goto L907
	}
L906:
	;
	v6392 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6393 = F_assign_special_exec_param(m, l0)
	mBase = m.M
	v6394 = m.ExcPending
	if v6394 != 0 {
		goto L1
	} else {
		goto L921
	}
L907:
	;
	v6271 = *(*int32)(unsafe.Add(mBase, uint32(v6268)+4))
	if v6271 <= int32(0) {
		v6350 = v6266
		goto L906
	} else {
		goto L908
	}
L908:
	;
	v6274 = *(*int32)(unsafe.Add(mBase, uint32(v6267)+8))
	v6277 = v6266
	v6278 = v6261
	v6280 = v4
	goto L909
L909:
	;
	v6319 = *(*int32)(unsafe.Add(mBase, uint32(v6268)+12))
	v6323 = *(*int32)(unsafe.Add(mBase, uint32(v6319+v6280<<(uint(int32(2))%32))))
	v6324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6324 != 0 {
		goto L911
	} else {
		goto L912
	}
L910:
	;
	v6350 = v6342
	goto L906
L911:
	;
	v6325 = F_replace_nestloop_params_mutator(m, v6323, l0)
	mBase = m.M
	v6326 = m.ExcPending
	if v6326 != 0 {
		goto L1
	} else {
		goto L914
	}
L912:
	;
	v6327 = v6323
	goto L913
L913:
	;
	v6329 = int32(0)
	v6331 = F_makeTargetEntry(m, v6327, base.I32_extend16_s(v6278), v6329, v6329)
	mBase = m.M
	v6332 = m.ExcPending
	if v6332 != 0 {
		goto L1
	} else {
		goto L915
	}
L914:
	;
	v6327 = v6325
	goto L913
L915:
	;
	if v6274 != 0 {
		goto L916
	} else {
		goto L917
	}
L916:
	;
	v6338 = *(*int32)(unsafe.Add(mBase, uint32(v6274+v6278<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v6331)+16)) = v6338
	goto L918
L917:
	;
	goto L918
L918:
	;
	v6342 = F_lappend(m, v6277, v6331)
	mBase = m.M
	v6343 = m.ExcPending
	if v6343 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	v6345 = v6280 + int32(1)
	v6346 = *(*int32)(unsafe.Add(mBase, uint32(v6268)+4))
	if v6345 < v6346 {
		v6277 = v6342
		v6278 = v6278 + int32(1)
		v6280 = v6345
		goto L909
	} else {
		goto L920
	}
L920:
	;
	goto L910
L921:
	;
	v6395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v6397 = F_palloc0(m, int32(88))
	mBase = m.M
	v6398 = m.ExcPending
	if v6398 != 0 {
		goto L1
	} else {
		goto L922
	}
L922:
	;
	v6399 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+84)) = v6399
	*(*uint8)(unsafe.Add(mBase, uint32(v6397)+81)) = uint8(v6399)
	*(*uint8)(unsafe.Add(mBase, uint32(v6397)+80)) = uint8(v6395)
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+76)) = v6393
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+72)) = v6392
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+56)) = v6399
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+52)) = v6264
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+48)) = v6399
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+44)) = v6350
	*(*int32)(unsafe.Add(mBase, uint32(v6397))) = int32(368)
	v6414 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+4)) = v6414
	v6416 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6397)+8)) = v6416
	v6418 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6397)+16)) = v6418
	v6420 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6397)+24)) = v6420
	v6422 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6423 = *(*int32)(unsafe.Add(mBase, uint32(v6422)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6397)+32)) = v6423
	v6425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6397)+36)) = uint8(v6425)
	v6427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6397)+37)) = uint8(v6427)
	v6429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6430 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6429)+83)) = uint8(v6430)
	v12573 = v6397
	v12580 = v47
	goto L3
L923:
	;
	v6435 = m.G0
	v6437 = v6435 - int32(16)
	m.G0 = v6437
	v6439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6442 = F_create_plan_recurse(m, l0, v6439, l2|int32(4))
	mBase = m.M
	v6443 = m.ExcPending
	if v6443 != 0 {
		goto L1
	} else {
		goto L927
	}
L924:
	;
	goto L925
L925:
	;
	v6823 = m.G0
	v6825 = v6823 - int32(112)
	m.G0 = v6825
	v6827 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v6828 = F_create_plan_recurse(m, l0, v6827, l2)
	mBase = m.M
	v6829 = m.ExcPending
	if v6829 != 0 {
		goto L1
	} else {
		goto L970
	}
L926:
	;
	v12573 = v6447
	v12580 = v47
	goto L3
L927:
	;
	v6444 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v6445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v6447 = F_palloc0(m, int32(88))
	mBase = m.M
	v6448 = m.ExcPending
	if v6448 != 0 {
		goto L1
	} else {
		goto L928
	}
L928:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6447))) = int32(367)
	v6451 = *(*int32)(unsafe.Add(mBase, uint32(v6442)+44))
	v6452 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+56)) = v6452
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+52)) = v6442
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+48)) = v6452
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+44)) = v6451
	v6460 = F_palloc(m, v6445<<(uint(int32(1))%32))
	mBase = m.M
	v6461 = m.ExcPending
	if v6461 != 0 {
		goto L1
	} else {
		goto L929
	}
L929:
	;
	v6463 = v6445 << (uint(int32(2)) % 32)
	v6464 = F_palloc(m, v6463)
	mBase = m.M
	v6465 = m.ExcPending
	if v6465 != 0 {
		goto L1
	} else {
		goto L930
	}
L930:
	;
	v6466 = F_palloc(m, v6463)
	mBase = m.M
	v6467 = m.ExcPending
	if v6467 != 0 {
		goto L1
	} else {
		goto L931
	}
L931:
	;
	if v6444 == int32(0) {
		goto L934
	} else {
		goto L935
	}
L932:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6813 = m.ExcPending
	if v6813 != 0 {
		goto L1
	} else {
		goto L967
	}
L933:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6751 = m.ExcPending
	if v6751 != 0 {
		goto L1
	} else {
		goto L964
	}
L934:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+84)) = v6466
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+80)) = v6464
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+76)) = v6460
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+72)) = v6445
	v6730 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+4)) = v6730
	v6732 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v6447)+8)) = v6732
	v6734 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v6447)+16)) = v6734
	v6736 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v6447)+24)) = v6736
	v6738 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6739 = *(*int32)(unsafe.Add(mBase, uint32(v6738)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v6447)+32)) = v6739
	v6741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6447)+36)) = uint8(v6741)
	v6743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6447)+37)) = uint8(v6743)
	m.G0 = v6437 + int32(16)
	goto L926
L935:
	;
	v6470 = *(*int32)(unsafe.Add(mBase, uint32(v6444)+4))
	if v6470 <= int32(0) {
		goto L934
	} else {
		goto L936
	}
L936:
	;
	v6473 = int32(0)
	if v6473 < v6445 {
		goto L937
	} else {
		goto L938
	}
L937:
	;
	v6476 = v6445
	goto L939
L938:
	;
	v6476 = v6473
	goto L939
L939:
	;
	v6482 = v4
	goto L940
L940:
	;
	if v6482 == v6476 {
		goto L934
	} else {
		goto L942
	}
L941:
	;
	goto L934
L942:
	;
	v6523 = v6482 << (uint(int32(2)) % 32)
	v6524 = *(*int32)(unsafe.Add(mBase, uint32(v6444)+12))
	v6526 = *(*int32)(unsafe.Add(mBase, uint32(v6523+v6524)))
	v6527 = *(*int32)(unsafe.Add(mBase, uint32(v6526)+4))
	v6528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6527)+41)))
	if v6528 == int32(1) {
		goto L945
	} else {
		goto L946
	}
L943:
	;
	v6661 = *(*int32)(unsafe.Add(mBase, uint32(v6526)+8))
	v6662 = *(*int32)(unsafe.Add(mBase, uint32(v6624)+16))
	v6664 = F_get_opfamily_member_for_cmptype(m, v6661, v6662, v6662, int32(3))
	mBase = m.M
	v6665 = m.ExcPending
	if v6665 != 0 {
		goto L1
	} else {
		goto L961
	}
L944:
	;
	v6609 = *(*int32)(unsafe.Add(mBase, uint32(v6447)+44))
	v6610 = F_get_sortgroupref_tle(m, v6531, v6609)
	mBase = m.M
	v6611 = m.ExcPending
	if v6611 != 0 {
		goto L1
	} else {
		goto L959
	}
L945:
	;
	v6531 = *(*int32)(unsafe.Add(mBase, uint32(v6527)+44))
	if v6531 != 0 {
		goto L944
	} else {
		goto L948
	}
L946:
	;
	goto L947
L947:
	;
	v6545 = *(*int32)(unsafe.Add(mBase, uint32(v6447)+44))
	if v6545 == int32(0) {
		goto L932
	} else {
		goto L952
	}
L948:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6535 = m.ExcPending
	if v6535 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_32), int32(0))
	mBase = m.M
	v6539 = m.ExcPending
	if v6539 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_33), int32(_a_F_create_plan_recurse_34))
	mBase = m.M
	v6544 = m.ExcPending
	if v6544 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L952:
	;
	v6548 = int32(0)
	v6549 = *(*int32)(unsafe.Add(mBase, uint32(v6545)+4))
	if v6549 <= v6548 {
		goto L932
	} else {
		goto L953
	}
L953:
	;
	v6554 = v6548
	goto L954
L954:
	;
	v6596 = *(*int32)(unsafe.Add(mBase, uint32(v6545)+12))
	v6600 = *(*int32)(unsafe.Add(mBase, uint32(v6596+v6554<<(uint(int32(2))%32))))
	v6601 = *(*int32)(unsafe.Add(mBase, uint32(v6600)+4))
	v6603 = F_find_ec_member_matching_expr(m, v6527, v6601, int32(0))
	mBase = m.M
	v6604 = m.ExcPending
	if v6604 != 0 {
		goto L1
	} else {
		goto L956
	}
L955:
	;
	goto L932
L956:
	;
	if v6603 != 0 {
		v6624 = v6603
		v6629 = v6600
		goto L943
	} else {
		goto L957
	}
L957:
	;
	v6606 = v6554 + int32(1)
	v6607 = *(*int32)(unsafe.Add(mBase, uint32(v6545)+4))
	if v6606 < v6607 {
		v6554 = v6606
		goto L954
	} else {
		goto L958
	}
L958:
	;
	goto L955
L959:
	;
	if v6610 == int32(0) {
		goto L932
	} else {
		goto L960
	}
L960:
	;
	v6614 = *(*int32)(unsafe.Add(mBase, uint32(v6527)+16))
	v6615 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+12))
	v6616 = *(*int32)(unsafe.Add(mBase, uint32(v6615)))
	v6624 = v6616
	v6629 = v6610
	goto L943
L961:
	;
	if v6664 == int32(0) {
		goto L933
	} else {
		goto L962
	}
L962:
	;
	v6668 = int32(1)
	v6671 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6629)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v6460+v6482<<(uint(v6668)%32)))) = uint16(v6671)
	*(*int32)(unsafe.Add(mBase, uint32(v6464+v6523))) = v6664
	v6676 = *(*int32)(unsafe.Add(mBase, uint32(v6527)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6523+v6466))) = v6676
	v6679 = v6482 + v6668
	v6680 = *(*int32)(unsafe.Add(mBase, uint32(v6444)+4))
	if v6679 < v6680 {
		v6482 = v6679
		goto L940
	} else {
		goto L963
	}
L963:
	;
	goto L941
L964:
	;
	v6752 = *(*int32)(unsafe.Add(mBase, uint32(v6526)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6437)+12)) = v6752
	*(*int32)(unsafe.Add(mBase, uint32(v6437)+8)) = v6662
	*(*int32)(unsafe.Add(mBase, uint32(v6437)+4)) = v6662
	*(*int32)(unsafe.Add(mBase, uint32(v6437))) = int32(3)
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_35), v6437)
	mBase = m.M
	v6760 = m.ExcPending
	if v6760 != 0 {
		goto L1
	} else {
		goto L965
	}
L965:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_36), int32(_a_F_create_plan_recurse_34))
	mBase = m.M
	v6765 = m.ExcPending
	if v6765 != 0 {
		goto L1
	} else {
		goto L966
	}
L966:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L967:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_37), int32(0))
	mBase = m.M
	v6817 = m.ExcPending
	if v6817 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(_a_F_create_plan_recurse_38), int32(_a_F_create_plan_recurse_34))
	mBase = m.M
	v6822 = m.ExcPending
	if v6822 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L970:
	;
	v6830 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v6830 == int32(0) {
		goto L972
	} else {
		goto L973
	}
L971:
	;
	m.G0 = v6825 + int32(112)
	v12573 = v8471
	v12580 = v47
	goto L3
L972:
	;
	v8471 = v6828
	goto L971
L973:
	;
	goto L974
L974:
	;
	v6833 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v6834 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v6835 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v6836 = *(*int32)(unsafe.Add(mBase, uint32(v6835)+4))
	if v6836 == int32(0) {
		goto L976
	} else {
		goto L977
	}
L975:
	;
	if v6833 != 0 {
		goto L994
	} else {
		goto L995
	}
L976:
	;
	v6973 = int32(0)
	v7012 = int32(1)
	goto L975
L977:
	;
	v6839 = int32(1)
	v6840 = *(*int32)(unsafe.Add(mBase, uint32(v6836)+4))
	if v6840 <= int32(0) {
		v6973 = v4
		v7012 = v6839
		goto L975
	} else {
		goto L978
	}
L978:
	;
	v6843 = *(*int32)(unsafe.Add(mBase, uint32(v6835)+8))
	v6846 = v6839
	v6849 = v4
	v6853 = v4
	goto L979
L979:
	;
	v6888 = *(*int32)(unsafe.Add(mBase, uint32(v6836)+12))
	v6892 = *(*int32)(unsafe.Add(mBase, uint32(v6888+v6853<<(uint(int32(2))%32))))
	v6893 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v6893 != 0 {
		goto L981
	} else {
		goto L982
	}
L980:
	;
	if v6911 == int32(0) {
		goto L976
	} else {
		goto L991
	}
L981:
	;
	v6894 = F_replace_nestloop_params_mutator(m, v6892, l0)
	mBase = m.M
	v6895 = m.ExcPending
	if v6895 != 0 {
		goto L1
	} else {
		goto L984
	}
L982:
	;
	v6896 = v6892
	goto L983
L983:
	;
	v6898 = int32(0)
	v6900 = F_makeTargetEntry(m, v6896, base.I32_extend16_s(v6846), v6898, v6898)
	mBase = m.M
	v6901 = m.ExcPending
	if v6901 != 0 {
		goto L1
	} else {
		goto L985
	}
L984:
	;
	v6896 = v6894
	goto L983
L985:
	;
	if v6843 != 0 {
		goto L986
	} else {
		goto L987
	}
L986:
	;
	v6907 = *(*int32)(unsafe.Add(mBase, uint32(v6843+v6846<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v6900)+16)) = v6907
	goto L988
L987:
	;
	goto L988
L988:
	;
	v6911 = F_lappend(m, v6849, v6900)
	mBase = m.M
	v6912 = m.ExcPending
	if v6912 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	v6914 = v6853 + int32(1)
	v6915 = *(*int32)(unsafe.Add(mBase, uint32(v6836)+4))
	if v6914 < v6915 {
		v6846 = v6846 + int32(1)
		v6849 = v6911
		v6853 = v6914
		goto L979
	} else {
		goto L990
	}
L990:
	;
	goto L980
L991:
	;
	v6919 = *(*int32)(unsafe.Add(mBase, uint32(v6911)+4))
	v6973 = v6911
	v7012 = v6919 + int32(1)
	goto L975
L992:
	;
	v7272 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+44))
	if v6833 != 0 {
		goto L1012
	} else {
		goto L1013
	}
L993:
	;
	v7225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v7226 = F_change_plan_targetlist(m, v6828, v7186, v7225)
	mBase = m.M
	v7227 = m.ExcPending
	if v7227 != 0 {
		goto L1
	} else {
		goto L1011
	}
L994:
	;
	v7013 = int32(0)
	v7014 = *(*int32)(unsafe.Add(mBase, uint32(v6833)+4))
	if v7013 < v7014 {
		goto L997
	} else {
		goto L998
	}
L995:
	;
	v7139 = v6973
	goto L996
L996:
	;
	v7178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v7178 != int32(2) {
		v7228 = v6828
		goto L992
	} else {
		goto L1010
	}
L997:
	;
	v7020 = v7012
	v7021 = int32(0)
	v7023 = v6973
	v7024 = v7013
	goto L1000
L998:
	;
	v7093 = v6973
	v7094 = v7013
	goto L999
L999:
	;
	if v7094&int32(1) != 0 {
		v7186 = v7093
		goto L993
	} else {
		goto L1009
	}
L1000:
	;
	v7062 = *(*int32)(unsafe.Add(mBase, uint32(v6833)+12))
	v7066 = *(*int32)(unsafe.Add(mBase, uint32(v7062+v7021<<(uint(int32(2))%32))))
	v7067 = F_tlist_member(m, v7066, v7023)
	mBase = m.M
	v7068 = m.ExcPending
	if v7068 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1001:
	;
	v7093 = v7082
	v7094 = v7083
	goto L999
L1002:
	;
	if v7067 == int32(0) {
		goto L1003
	} else {
		goto L1004
	}
L1003:
	;
	v7073 = int32(0)
	v7075 = F_makeTargetEntry(m, v7066, base.I32_extend16_s(v7020), v7073, v7073)
	mBase = m.M
	v7076 = m.ExcPending
	if v7076 != 0 {
		goto L1
	} else {
		goto L1006
	}
L1004:
	;
	v7081 = v7020
	v7082 = v7023
	v7083 = v7024
	goto L1005
L1005:
	;
	v7085 = v7021 + int32(1)
	v7086 = *(*int32)(unsafe.Add(mBase, uint32(v6833)+4))
	if v7085 < v7086 {
		v7020 = v7081
		v7021 = v7085
		v7023 = v7082
		v7024 = v7083
		goto L1000
	} else {
		goto L1008
	}
L1006:
	;
	v7077 = F_lappend(m, v7023, v7075)
	mBase = m.M
	v7078 = m.ExcPending
	if v7078 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1007:
	;
	v7081 = v7020 + int32(1)
	v7082 = v7077
	v7083 = int32(1)
	goto L1005
L1008:
	;
	goto L1001
L1009:
	;
	v7139 = v7093
	goto L996
L1010:
	;
	v7186 = v7139
	goto L993
L1011:
	;
	v7228 = v7226
	goto L992
L1012:
	;
	v7273 = *(*int32)(unsafe.Add(mBase, uint32(v6833)+4))
	v7275 = v7273
	goto L1014
L1013:
	;
	v7275 = int32(0)
	goto L1014
L1014:
	;
	v7278 = F_palloc(m, v7275<<(uint(int32(1))%32))
	mBase = m.M
	v7279 = m.ExcPending
	if v7279 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	v7281 = v7275 << (uint(int32(2)) % 32)
	v7282 = F_palloc(m, v7281)
	mBase = m.M
	v7283 = m.ExcPending
	if v7283 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	if v6833 == int32(0) {
		goto L1021
	} else {
		goto L1022
	}
L1017:
	;
	v8454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+4)) = v8454
	v8456 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8412)+8)) = v8456
	v8458 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8412)+16)) = v8458
	v8460 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8412)+24)) = v8460
	v8462 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8463 = *(*int32)(unsafe.Add(mBase, uint32(v8462)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8412)+32)) = v8463
	v8465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8412)+36)) = uint8(v8465)
	v8467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8412)+37)) = uint8(v8467)
	v8471 = v8412
	goto L971
L1018:
	;
	v8055 = int32(0)
	v8057 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+44))
	if v8015 != 0 {
		goto L1144
	} else {
		goto L1145
	}
L1019:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8001 = m.ExcPending
	if v8001 != 0 {
		goto L1
	} else {
		goto L1141
	}
L1020:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7988 = m.ExcPending
	if v7988 != 0 {
		goto L1
	} else {
		goto L1138
	}
L1021:
	;
	v7401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v7401 != int32(1) {
		goto L1031
	} else {
		goto L1032
	}
L1022:
	;
	v7286 = *(*int32)(unsafe.Add(mBase, uint32(v6833)+4))
	if v7286 <= int32(0) {
		goto L1021
	} else {
		goto L1023
	}
L1023:
	;
	v7292 = int32(0)
	goto L1024
L1024:
	;
	v7335 = v7292 << (uint(int32(2)) % 32)
	v7336 = *(*int32)(unsafe.Add(mBase, uint32(v6833)+12))
	v7338 = *(*int32)(unsafe.Add(mBase, uint32(v7335+v7336)))
	v7339 = F_tlist_member(m, v7338, v7272)
	mBase = m.M
	v7340 = m.ExcPending
	if v7340 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1025:
	;
	goto L1021
L1026:
	;
	if v7339 == int32(0) {
		goto L1020
	} else {
		goto L1027
	}
L1027:
	;
	v7346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7339)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v7278+v7292<<(uint(int32(1))%32)))) = uint16(v7346)
	v7349 = *(*int32)(unsafe.Add(mBase, uint32(v7339)+4))
	v7350 = F_exprCollation(m, v7349)
	mBase = m.M
	v7351 = m.ExcPending
	if v7351 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1028:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7282+v7335))) = v7350
	v7354 = v7292 + int32(1)
	v7355 = *(*int32)(unsafe.Add(mBase, uint32(v6833)+4))
	if v7354 < v7355 {
		v7292 = v7354
		goto L1024
	} else {
		goto L1029
	}
L1029:
	;
	goto L1025
L1030:
	;
	v7698 = v7404
	v7699 = int32(0)
	goto L1072
L1031:
	;
	v7404 = int32(0)
	if v6834 == v7404 {
		v8015 = v7404
		goto L1018
	} else {
		goto L1034
	}
L1032:
	;
	goto L1033
L1033:
	;
	v7412 = F_palloc(m, v7281)
	mBase = m.M
	v7413 = m.ExcPending
	if v7413 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1034:
	;
	v7408 = *(*int32)(unsafe.Add(mBase, uint32(v6834)+4))
	if v7408 <= int32(0) {
		v8015 = v7404
		goto L1018
	} else {
		goto L1035
	}
L1035:
	;
	goto L1030
L1036:
	;
	if v6834 == int32(0) {
		goto L1037
	} else {
		goto L1038
	}
L1037:
	;
	v7526 = int32(0)
	v7527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7528 = *(*int32)(unsafe.Add(mBase, uint32(v7527)+4))
	if v7528 == v7526 {
		v7613 = v7526
		goto L1045
	} else {
		goto L1046
	}
L1038:
	;
	v7416 = *(*int32)(unsafe.Add(mBase, uint32(v6834)+4))
	if v7416 <= int32(0) {
		goto L1037
	} else {
		goto L1039
	}
L1039:
	;
	v7422 = int32(0)
	goto L1040
L1040:
	;
	v7465 = v7422 << (uint(int32(2)) % 32)
	v7466 = *(*int32)(unsafe.Add(mBase, uint32(v6834)+12))
	v7468 = *(*int32)(unsafe.Add(mBase, uint32(v7465+v7466)))
	v7471 = F_get_compatible_hash_operators(m, v7468, v6825+int32(40))
	mBase = m.M
	v7472 = m.ExcPending
	if v7472 != 0 {
		goto L1
	} else {
		goto L1042
	}
L1041:
	;
	goto L1037
L1042:
	;
	if v7471 == int32(0) {
		goto L1019
	} else {
		goto L1043
	}
L1043:
	;
	v7476 = *(*int32)(unsafe.Add(mBase, uint32(v6825)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v7465+v7412))) = v7476
	v7479 = v7422 + int32(1)
	v7480 = *(*int32)(unsafe.Add(mBase, uint32(v6834)+4))
	if v7479 < v7480 {
		v7422 = v7479
		goto L1040
	} else {
		goto L1044
	}
L1044:
	;
	goto L1041
L1045:
	;
	v7654 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v7656 = F_palloc0(m, int32(128))
	mBase = m.M
	v7657 = m.ExcPending
	if v7657 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1046:
	;
	v7532 = *(*int32)(unsafe.Add(mBase, uint32(v7528)+4))
	if v7532 <= int32(0) {
		v7613 = v7526
		goto L1045
	} else {
		goto L1047
	}
L1047:
	;
	v7535 = *(*int32)(unsafe.Add(mBase, uint32(v7527)+8))
	v7539 = int32(1)
	v7540 = v7526
	v7542 = int32(0)
	goto L1048
L1048:
	;
	v7581 = *(*int32)(unsafe.Add(mBase, uint32(v7528)+12))
	v7585 = *(*int32)(unsafe.Add(mBase, uint32(v7581+v7542<<(uint(int32(2))%32))))
	v7586 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v7586 != 0 {
		goto L1050
	} else {
		goto L1051
	}
L1049:
	;
	v7613 = v7604
	goto L1045
L1050:
	;
	v7587 = F_replace_nestloop_params_mutator(m, v7585, l0)
	mBase = m.M
	v7588 = m.ExcPending
	if v7588 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1051:
	;
	v7589 = v7585
	goto L1052
L1052:
	;
	v7591 = int32(0)
	v7593 = F_makeTargetEntry(m, v7589, base.I32_extend16_s(v7539), v7591, v7591)
	mBase = m.M
	v7594 = m.ExcPending
	if v7594 != 0 {
		goto L1
	} else {
		goto L1054
	}
L1053:
	;
	v7589 = v7587
	goto L1052
L1054:
	;
	if v7535 != 0 {
		goto L1055
	} else {
		goto L1056
	}
L1055:
	;
	v7600 = *(*int32)(unsafe.Add(mBase, uint32(v7535+v7539<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v7593)+16)) = v7600
	goto L1057
L1056:
	;
	goto L1057
L1057:
	;
	v7604 = F_lappend(m, v7540, v7593)
	mBase = m.M
	v7605 = m.ExcPending
	if v7605 != 0 {
		goto L1
	} else {
		goto L1058
	}
L1058:
	;
	v7607 = v7542 + int32(1)
	v7608 = *(*int32)(unsafe.Add(mBase, uint32(v7528)+4))
	if v7607 < v7608 {
		v7539 = v7539 + int32(1)
		v7540 = v7604
		v7542 = v7607
		goto L1048
	} else {
		goto L1059
	}
L1059:
	;
	goto L1049
L1060:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7656))) = int32(365)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v7654)&int64(9223372036854775807)) {
		goto L1062
	} else {
		goto L1063
	}
L1061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+96)) = v7674
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+92)) = v7282
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+88)) = v7412
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+84)) = v7278
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+80)) = v7275
	*(*int64)(unsafe.Add(mBase, uint32(v7656)+72)) = int64(2)
	v7682 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7656)+104)) = v7682
	v7684 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+48)) = v7684
	*(*int64)(unsafe.Add(mBase, uint32(v7656)+112)) = v7682
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+120)) = v7684
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+56)) = v7684
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+52)) = v7228
	*(*int32)(unsafe.Add(mBase, uint32(v7656)+44)) = v7613
	v8412 = v7656
	goto L1017
L1062:
	;
	v7674 = int32(2147483647)
	goto L1061
L1063:
	;
	goto L1064
L1064:
	;
	if base.F64_le(v7654, float64(0)) != 0 {
		goto L1065
	} else {
		goto L1066
	}
L1065:
	;
	v7674 = int32(0)
	goto L1061
L1066:
	;
	goto L1067
L1067:
	;
	v7669 = float64(2.147483647e+09)
	if base.F64_lt(v7654, v7669) != 0 {
		goto L1068
	} else {
		goto L1069
	}
L1068:
	;
	v7672 = v7654
	goto L1070
L1069:
	;
	v7672 = v7669
	goto L1070
L1070:
	;
	v7674 = base.I32_trunc_sat_f64_s(v7672)
	goto L1061
L1071:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7973 = m.ExcPending
	if v7973 != 0 {
		goto L1
	} else {
		goto L1135
	}
L1072:
	;
	v7738 = *(*int32)(unsafe.Add(mBase, uint32(v6834)+12))
	v7742 = *(*int32)(unsafe.Add(mBase, uint32(v7738+v7699<<(uint(int32(2))%32))))
	v7743 = F_get_ordering_op_for_equality_op(m, v7742)
	mBase = m.M
	v7744 = m.ExcPending
	if v7744 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1073:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7958 = m.ExcPending
	if v7958 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1074:
	;
	if v7743 != 0 {
		goto L1075
	} else {
		goto L1076
	}
L1075:
	;
	v7746 = F_get_equality_op_for_ordering_op(m, v7743, int32(0))
	mBase = m.M
	v7747 = m.ExcPending
	if v7747 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1076:
	;
	goto L1077
L1077:
	;
	goto L1073
L1078:
	;
	if v7746 == int32(0) {
		goto L1071
	} else {
		goto L1079
	}
L1079:
	;
	v7750 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+44))
	v7754 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7278+v7699<<(uint(int32(1))%32)))))
	if v7750 != 0 {
		goto L1082
	} else {
		goto L1083
	}
L1080:
	;
	v7795 = F_palloc0(m, int32(20))
	mBase = m.M
	v7796 = m.ExcPending
	if v7796 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1081:
	;
	goto L1080
L1082:
	;
	v7758 = *(*int32)(unsafe.Add(mBase, uint32(v7750)+4))
	if v7758 <= int32(0) {
		v7792 = int32(0)
		goto L1081
	} else {
		goto L1085
	}
L1083:
	;
	goto L1084
L1084:
	;
	v7792 = int32(0)
	goto L1081
L1085:
	;
	v7761 = int32(0)
	if v7761 < v7758 {
		goto L1086
	} else {
		goto L1087
	}
L1086:
	;
	v7764 = v7758
	goto L1088
L1087:
	;
	v7764 = v7761
	goto L1088
L1088:
	;
	v7765 = *(*int32)(unsafe.Add(mBase, uint32(v7750)+12))
	v7769 = int32(0)
	goto L1089
L1089:
	;
	v7777 = *(*int32)(unsafe.Add(mBase, uint32(v7765+v7769<<(uint(int32(2))%32))))
	v7778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7777)+8)))
	if v7778 == v7754&int32(_a_F_create_plan_recurse_28) {
		v7792 = v7777
		goto L1081
	} else {
		goto L1091
	}
L1090:
	;
	goto L1084
L1091:
	;
	v7781 = v7769 + int32(1)
	if v7781 != v7764 {
		v7769 = v7781
		goto L1089
	} else {
		goto L1092
	}
L1092:
	;
	goto L1090
L1093:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7795))) = int32(106)
	v7799 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+44))
	v7800 = int32(0)
	v7809 = *(*int32)(unsafe.Add(mBase, uint32(v7792)+16))
	if v7809 == v7800 {
		goto L1095
	} else {
		goto L1096
	}
L1094:
	;
	v7942 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7795)+18)) = uint8(v7942)
	*(*uint16)(unsafe.Add(mBase, uint32(v7795)+16)) = uint16(v7942)
	*(*int32)(unsafe.Add(mBase, uint32(v7795)+12)) = v7743
	*(*int32)(unsafe.Add(mBase, uint32(v7795)+8)) = v7746
	*(*int32)(unsafe.Add(mBase, uint32(v7795)+4)) = v7933
	v7949 = F_lappend(m, v7698, v7795)
	mBase = m.M
	v7950 = m.ExcPending
	if v7950 != 0 {
		goto L1
	} else {
		goto L1130
	}
L1095:
	;
	if v7799 == int32(0) {
		v7929 = int32(1)
		goto L1098
	} else {
		goto L1099
	}
L1096:
	;
	v7933 = v7809
	goto L1097
L1097:
	;
	goto L1094
L1098:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7792)+16)) = v7929
	v7933 = v7929
	goto L1097
L1099:
	;
	v7816 = *(*int32)(unsafe.Add(mBase, uint32(v7799)+4))
	if v7816 <= int32(0) {
		v7929 = int32(1)
		goto L1098
	} else {
		goto L1100
	}
L1100:
	;
	v7819 = int32(0)
	if v7819 < v7816 {
		goto L1101
	} else {
		goto L1102
	}
L1101:
	;
	v7822 = v7816
	goto L1103
L1102:
	;
	v7822 = v7819
	goto L1103
L1103:
	;
	v7824 = v7822 & int32(3)
	v7825 = int32(0)
	if int32(4) <= v7816 {
		goto L1105
	} else {
		goto L1106
	}
L1104:
	;
	v7929 = v7907 + int32(1)
	goto L1098
L1105:
	;
	v7830 = *(*int32)(unsafe.Add(mBase, uint32(v7799)+12))
	v7834 = v7825
	v7835 = int32(0)
	v7836 = v7800
	goto L1108
L1106:
	;
	v7871 = v7825
	v7873 = v7800
	goto L1107
L1107:
	;
	v7880 = *(*int32)(unsafe.Add(mBase, uint32(v7799)+12))
	v7882 = int32(0)
	v7884 = v7871
	v7886 = v7873
	goto L1124
L1108:
	;
	v7845 = v7830 + v7836<<(uint(int32(2))%32)
	v7846 = *(*int32)(unsafe.Add(mBase, uint32(v7845)+12))
	v7847 = *(*int32)(unsafe.Add(mBase, uint32(v7846)+16))
	v7848 = *(*int32)(unsafe.Add(mBase, uint32(v7845)+8))
	v7849 = *(*int32)(unsafe.Add(mBase, uint32(v7848)+16))
	v7850 = *(*int32)(unsafe.Add(mBase, uint32(v7845)+4))
	v7851 = *(*int32)(unsafe.Add(mBase, uint32(v7850)+16))
	v7852 = *(*int32)(unsafe.Add(mBase, uint32(v7845)))
	v7853 = *(*int32)(unsafe.Add(mBase, uint32(v7852)+16))
	if base.Ui32(v7834) < base.Ui32(v7853) {
		goto L1110
	} else {
		goto L1111
	}
L1109:
	;
	if v7824 == int32(0) {
		v7907 = v7861
		goto L1104
	} else {
		goto L1123
	}
L1110:
	;
	v7855 = v7853
	goto L1112
L1111:
	;
	v7855 = v7834
	goto L1112
L1112:
	;
	if base.Ui32(v7855) < base.Ui32(v7851) {
		goto L1113
	} else {
		goto L1114
	}
L1113:
	;
	v7857 = v7851
	goto L1115
L1114:
	;
	v7857 = v7855
	goto L1115
L1115:
	;
	if base.Ui32(v7857) < base.Ui32(v7849) {
		goto L1116
	} else {
		goto L1117
	}
L1116:
	;
	v7859 = v7849
	goto L1118
L1117:
	;
	v7859 = v7857
	goto L1118
L1118:
	;
	if base.Ui32(v7859) < base.Ui32(v7847) {
		goto L1119
	} else {
		goto L1120
	}
L1119:
	;
	v7861 = v7847
	goto L1121
L1120:
	;
	v7861 = v7859
	goto L1121
L1121:
	;
	v7862 = int32(4)
	v7863 = v7836 + v7862
	v7865 = v7835 + v7862
	if v7865 != v7822&int32(2147483644) {
		v7834 = v7861
		v7835 = v7865
		v7836 = v7863
		goto L1108
	} else {
		goto L1122
	}
L1122:
	;
	goto L1109
L1123:
	;
	v7871 = v7861
	v7873 = v7863
	goto L1107
L1124:
	;
	v7896 = *(*int32)(unsafe.Add(mBase, uint32(v7880+v7886<<(uint(int32(2))%32))))
	v7897 = *(*int32)(unsafe.Add(mBase, uint32(v7896)+16))
	if base.Ui32(v7884) < base.Ui32(v7897) {
		goto L1126
	} else {
		goto L1127
	}
L1125:
	;
	v7907 = v7899
	goto L1104
L1126:
	;
	v7899 = v7897
	goto L1128
L1127:
	;
	v7899 = v7884
	goto L1128
L1128:
	;
	v7900 = int32(1)
	v7903 = v7882 + v7900
	if v7903 != v7824 {
		v7882 = v7903
		v7884 = v7899
		v7886 = v7886 + v7900
		goto L1124
	} else {
		goto L1129
	}
L1129:
	;
	goto L1125
L1130:
	;
	v7952 = v7699 + int32(1)
	v7953 = *(*int32)(unsafe.Add(mBase, uint32(v6834)+4))
	if v7952 < v7953 {
		v7698 = v7949
		v7699 = v7952
		goto L1072
	} else {
		goto L1131
	}
L1131:
	;
	v8015 = v7949
	goto L1018
L1132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6825)+16)) = v7742
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_39), v6825+int32(16))
	mBase = m.M
	v7964 = m.ExcPending
	if v7964 != 0 {
		goto L1
	} else {
		goto L1133
	}
L1133:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1875), int32(_a_F_create_plan_recurse_40))
	mBase = m.M
	v7969 = m.ExcPending
	if v7969 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6825)+32)) = v7743
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_41), v6825+int32(32))
	mBase = m.M
	v7979 = m.ExcPending
	if v7979 != 0 {
		goto L1
	} else {
		goto L1136
	}
L1136:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1886), int32(_a_F_create_plan_recurse_40))
	mBase = m.M
	v7984 = m.ExcPending
	if v7984 != 0 {
		goto L1
	} else {
		goto L1137
	}
L1137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1138:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_42), int32(0))
	mBase = m.M
	v7992 = m.ExcPending
	if v7992 != 0 {
		goto L1
	} else {
		goto L1139
	}
L1139:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1809), int32(_a_F_create_plan_recurse_40))
	mBase = m.M
	v7997 = m.ExcPending
	if v7997 != 0 {
		goto L1
	} else {
		goto L1140
	}
L1140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6825))) = v7468
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_43), v6825)
	mBase = m.M
	v8005 = m.ExcPending
	if v8005 != 0 {
		goto L1
	} else {
		goto L1142
	}
L1142:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1834), int32(_a_F_create_plan_recurse_40))
	mBase = m.M
	v8010 = m.ExcPending
	if v8010 != 0 {
		goto L1
	} else {
		goto L1143
	}
L1143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1144:
	;
	v8058 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+4))
	v8059 = v8058
	goto L1146
L1145:
	;
	v8059 = v8055
	goto L1146
L1146:
	;
	v8062 = F_palloc(m, v8059<<(uint(int32(1))%32))
	mBase = m.M
	v8063 = m.ExcPending
	if v8063 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1147:
	;
	v8065 = v8059 << (uint(int32(2)) % 32)
	v8066 = F_palloc(m, v8065)
	mBase = m.M
	v8067 = m.ExcPending
	if v8067 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1148:
	;
	v8068 = F_palloc(m, v8065)
	mBase = m.M
	v8069 = m.ExcPending
	if v8069 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1149:
	;
	v8070 = F_palloc(m, v8059)
	mBase = m.M
	v8071 = m.ExcPending
	if v8071 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1150:
	;
	if v8015 == int32(0) {
		v8154 = v8055
		goto L1151
	} else {
		goto L1152
	}
L1151:
	;
	v8193 = F_palloc0(m, int32(96))
	mBase = m.M
	v8194 = m.ExcPending
	if v8194 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1152:
	;
	v8074 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+4))
	if v8074 <= int32(0) {
		v8154 = v8055
		goto L1151
	} else {
		goto L1153
	}
L1153:
	;
	v8083 = v8055
	goto L1154
L1154:
	;
	v8125 = v8083 << (uint(int32(2)) % 32)
	v8126 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+12))
	v8128 = *(*int32)(unsafe.Add(mBase, uint32(v8125+v8126)))
	v8129 = F_get_sortgroupclause_tle(m, v8128, v8057)
	mBase = m.M
	v8130 = m.ExcPending
	if v8130 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1155:
	;
	v8154 = v8145
	goto L1151
L1156:
	;
	v8131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8129)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v8062+v8083<<(uint(int32(1))%32)))) = uint16(v8131)
	v8134 = *(*int32)(unsafe.Add(mBase, uint32(v8128)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8125+v8066))) = v8134
	v8137 = *(*int32)(unsafe.Add(mBase, uint32(v8129)+4))
	v8138 = F_exprCollation(m, v8137)
	mBase = m.M
	v8139 = m.ExcPending
	if v8139 != 0 {
		goto L1
	} else {
		goto L1157
	}
L1157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8125+v8068))) = v8138
	v8142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8128)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8083+v8070))) = uint8(v8142)
	v8145 = v8083 + int32(1)
	v8146 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+4))
	if v8145 < v8146 {
		v8083 = v8145
		goto L1154
	} else {
		goto L1158
	}
L1158:
	;
	goto L1155
L1159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8193))) = int32(362)
	v8197 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+44)) = v8197
	v8199 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+4))
	v8200 = int32(_a_F_create_plan_recurse_0)
	v8201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+88)) = v8070
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+84)) = v8068
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+80)) = v8066
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+76)) = v8062
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+72)) = v8154
	v8207 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+56)) = v8207
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+52)) = v7228
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+48)) = v8207
	v8212 = int32(1)
	v8214 = v8199 + (v8201 ^ v8212)
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+4)) = v8214
	v8217 = v6825 + int32(40)
	v8220 = *(*float64)(unsafe.Add(mBase, uint32(v7228)+16))
	v8221 = *(*float64)(unsafe.Add(mBase, uint32(v7228)+24))
	v8222 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+32))
	v8225 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v8228 = m.G0
	v8229 = int32(16)
	v8230 = v8228 - v8229
	m.G0 = v8230
	F_cost_tuplesort(m, v8230+int32(8), v8230, v8221, v8222, float64(0), v8225, float64(-1))
	mBase = m.M
	v8235 = *(*float64)(unsafe.Add(mBase, uint32(v8230)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v8217)+32)) = v8221
	v8238 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v8239 = base.F64_add(v8220, v8235)
	*(*float64)(unsafe.Add(mBase, uint32(v8217)+48)) = v8239
	*(*int32)(unsafe.Add(mBase, uint32(v8217)+40)) = v8214 + (v8238 ^ v8212)
	v8245 = *(*float64)(unsafe.Add(mBase, uint32(v8230)))
	*(*float64)(unsafe.Add(mBase, uint32(v8217)+56)) = base.F64_add(v8239, v8245)
	m.G0 = v8230 + v8229
	goto L1160
L1160:
	;
	v8251 = *(*float64)(unsafe.Add(mBase, uint32(v6825)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v8193)+8)) = v8251
	v8253 = *(*float64)(unsafe.Add(mBase, uint32(v6825)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v8193)+16)) = v8253
	v8255 = *(*float64)(unsafe.Add(mBase, uint32(v7228)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v8193)+24)) = v8255
	v8257 = *(*int32)(unsafe.Add(mBase, uint32(v7228)+32))
	v8258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8193)+36)) = uint8(v8258)
	*(*int32)(unsafe.Add(mBase, uint32(v8193)+32)) = v8257
	v8261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7228)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8193)+37)) = uint8(v8261)
	v8264 = F_palloc0(m, int32(88))
	mBase = m.M
	v8265 = m.ExcPending
	if v8265 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8264))) = int32(367)
	if v8015 != 0 {
		goto L1162
	} else {
		goto L1163
	}
L1162:
	;
	v8268 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+4))
	v8269 = v8268
	goto L1164
L1163:
	;
	v8269 = v7404
	goto L1164
L1164:
	;
	v8270 = *(*int32)(unsafe.Add(mBase, uint32(v8193)+44))
	v8271 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+56)) = v8271
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+52)) = v8193
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+48)) = v8271
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+44)) = v8270
	v8280 = F_palloc(m, v8269<<(uint(int32(1))%32))
	mBase = m.M
	v8281 = m.ExcPending
	if v8281 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	v8283 = v8269 << (uint(int32(2)) % 32)
	v8284 = F_palloc(m, v8283)
	mBase = m.M
	v8285 = m.ExcPending
	if v8285 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1166:
	;
	v8286 = F_palloc(m, v8283)
	mBase = m.M
	v8287 = m.ExcPending
	if v8287 != 0 {
		goto L1
	} else {
		goto L1167
	}
L1167:
	;
	if v8015 == int32(0) {
		goto L1168
	} else {
		goto L1169
	}
L1168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+84)) = v8286
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+80)) = v8284
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+76)) = v8280
	*(*int32)(unsafe.Add(mBase, uint32(v8264)+72)) = v8269
	v8412 = v8264
	goto L1017
L1169:
	;
	v8290 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+4))
	if v8290 <= int32(0) {
		goto L1168
	} else {
		goto L1170
	}
L1170:
	;
	v8298 = v8271
	goto L1171
L1171:
	;
	v8341 = v8298 << (uint(int32(2)) % 32)
	v8342 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+12))
	v8344 = *(*int32)(unsafe.Add(mBase, uint32(v8341+v8342)))
	v8345 = *(*int32)(unsafe.Add(mBase, uint32(v8264)+44))
	v8346 = F_get_sortgroupclause_tle(m, v8344, v8345)
	mBase = m.M
	v8347 = m.ExcPending
	if v8347 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1172:
	;
	goto L1168
L1173:
	;
	v8348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8346)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v8280+v8298<<(uint(int32(1))%32)))) = uint16(v8348)
	v8351 = *(*int32)(unsafe.Add(mBase, uint32(v8344)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8284+v8341))) = v8351
	v8354 = *(*int32)(unsafe.Add(mBase, uint32(v8346)+4))
	v8355 = F_exprCollation(m, v8354)
	mBase = m.M
	v8356 = m.ExcPending
	if v8356 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8286+v8341))) = v8355
	v8359 = v8298 + int32(1)
	v8360 = *(*int32)(unsafe.Add(mBase, uint32(v8015)+4))
	if v8359 < v8360 {
		v8298 = v8359
		goto L1171
	} else {
		goto L1175
	}
L1175:
	;
	goto L1172
L1176:
	;
	v8521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v8522 = F_replace_nestloop_params_mutator(m, v8521, l0)
	mBase = m.M
	v8523 = m.ExcPending
	if v8523 != 0 {
		goto L1
	} else {
		goto L1177
	}
L1177:
	;
	if v8522 != 0 {
		goto L1178
	} else {
		goto L1179
	}
L1178:
	;
	v8524 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+4))
	v8527 = v8524 << (uint(int32(2)) % 32)
	goto L1180
L1179:
	;
	v8527 = v4
	goto L1180
L1180:
	;
	v8528 = F_palloc(m, v8527)
	mBase = m.M
	v8529 = m.ExcPending
	if v8529 != 0 {
		goto L1
	} else {
		goto L1181
	}
L1181:
	;
	v8530 = F_palloc(m, v8527)
	mBase = m.M
	v8531 = m.ExcPending
	if v8531 != 0 {
		goto L1
	} else {
		goto L1182
	}
L1182:
	;
	v8532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v8536 = v4
	goto L1183
L1183:
	;
	v8577 = int32(0)
	if v8522 == v8577 {
		v8586 = v8577
		goto L1185
	} else {
		goto L1186
	}
L1184:
	;
	v8611 = m.G0
	v8613 = v8611 - int32(16)
	m.G0 = v8613
	v8615 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8613)+12)) = v8615
	if v8522 == v8615 {
		v8633 = v8615
		goto L1193
	} else {
		goto L1194
	}
L1185:
	;
	if v8532 == int32(0) {
		goto L1188
	} else {
		goto L1189
	}
L1186:
	;
	v8580 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+4))
	if v8580 <= v8536 {
		v8586 = v8577
		goto L1185
	} else {
		goto L1187
	}
L1187:
	;
	v8582 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+12))
	v8586 = v8582 + v8536<<(uint(int32(2))%32)
	goto L1185
L1188:
	;
	goto L1184
L1189:
	;
	v8591 = *(*int32)(unsafe.Add(mBase, uint32(v8532)+4))
	if base.B2i32(v8586 == int32(0))|base.B2i32(v8591 <= v8536) != 0 {
		goto L1188
	} else {
		goto L1190
	}
L1190:
	;
	v8594 = *(*int32)(unsafe.Add(mBase, uint32(v8532)+12))
	if v8594 == int32(0) {
		goto L1188
	} else {
		goto L1191
	}
L1191:
	;
	v8597 = *(*int32)(unsafe.Add(mBase, uint32(v8586)))
	v8599 = v8536 << (uint(int32(2)) % 32)
	v8602 = *(*int32)(unsafe.Add(mBase, uint32(v8599+v8594)))
	*(*int32)(unsafe.Add(mBase, uint32(v8528+v8599))) = v8602
	v8605 = F_exprCollation(m, v8597)
	mBase = m.M
	v8606 = m.ExcPending
	if v8606 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8599+v8530))) = v8605
	v8536 = v8536 + int32(1)
	goto L1183
L1193:
	;
	m.G0 = v8613 + int32(16)
	v8637 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v8638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+85)))
	v8639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	v8641 = F_palloc0(m, int32(104))
	mBase = m.M
	v8642 = m.ExcPending
	if v8642 != 0 {
		goto L1
	} else {
		goto L1200
	}
L1194:
	;
	v8620 = *(*int32)(unsafe.Add(mBase, uint32(v8522)))
	if v8620 == int32(8) {
		goto L1195
	} else {
		goto L1196
	}
L1195:
	;
	v8624 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+8))
	v8625 = F_bms_add_member(m, int32(0), v8624)
	mBase = m.M
	v8626 = m.ExcPending
	if v8626 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1196:
	;
	goto L1197
L1197:
	;
	v8630 = F_expression_tree_walker_impl(m, v8522, int32(876), v8613+int32(12))
	mBase = m.M
	v8631 = m.ExcPending
	if v8631 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1198:
	;
	v8633 = v8625
	goto L1193
L1199:
	;
	v8632 = *(*int32)(unsafe.Add(mBase, uint32(v8613)+12))
	v8633 = v8632
	goto L1193
L1200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8641))) = int32(361)
	v8645 = *(*int32)(unsafe.Add(mBase, uint32(v8519)+44))
	v8646 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+56)) = v8646
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+52)) = v8519
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+48)) = v8646
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+44)) = v8645
	if v8522 != 0 {
		goto L1201
	} else {
		goto L1202
	}
L1201:
	;
	v8653 = *(*int32)(unsafe.Add(mBase, uint32(v8522)+4))
	v8654 = v8653
	goto L1203
L1202:
	;
	v8654 = v8646
	goto L1203
L1203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+96)) = v8633
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+92)) = v8637
	*(*uint8)(unsafe.Add(mBase, uint32(v8641)+89)) = uint8(v8638)
	*(*uint8)(unsafe.Add(mBase, uint32(v8641)+88)) = uint8(v8639)
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+84)) = v8522
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+80)) = v8530
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+76)) = v8528
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+72)) = v8654
	v8663 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+4)) = v8663
	v8665 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8641)+8)) = v8665
	v8667 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8641)+16)) = v8667
	v8669 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8641)+24)) = v8669
	v8671 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8672 = *(*int32)(unsafe.Add(mBase, uint32(v8671)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8641)+32)) = v8672
	v8674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8641)+36)) = uint8(v8674)
	v8676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8641)+37)) = uint8(v8676)
	v12573 = v8641
	v12580 = v47
	goto L3
L1204:
	;
	v8684 = F_palloc0(m, int32(72))
	mBase = m.M
	v8685 = m.ExcPending
	if v8685 != 0 {
		goto L1
	} else {
		goto L1205
	}
L1205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8684))) = int32(360)
	v8688 = *(*int32)(unsafe.Add(mBase, uint32(v8681)+44))
	v8689 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8684)+56)) = v8689
	*(*int32)(unsafe.Add(mBase, uint32(v8684)+52)) = v8681
	*(*int32)(unsafe.Add(mBase, uint32(v8684)+48)) = v8689
	*(*int32)(unsafe.Add(mBase, uint32(v8684)+44)) = v8688
	v8695 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8684)+4)) = v8695
	v8697 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8684)+8)) = v8697
	v8699 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8684)+16)) = v8699
	v8701 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8684)+24)) = v8701
	v8703 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8704 = *(*int32)(unsafe.Add(mBase, uint32(v8703)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8684)+32)) = v8704
	v8706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8684)+36)) = uint8(v8706)
	v8708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8684)+37)) = uint8(v8708)
	v12573 = v8684
	v12580 = v47
	goto L3
L1206:
	;
	v8714 = int32(0)
	v8715 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8716 = *(*int32)(unsafe.Add(mBase, uint32(v8715)+4))
	if v8716 == v8714 {
		v8799 = v8714
		goto L1207
	} else {
		goto L1208
	}
L1207:
	;
	v8842 = F_palloc0(m, int32(72))
	mBase = m.M
	v8843 = m.ExcPending
	if v8843 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1208:
	;
	v8720 = *(*int32)(unsafe.Add(mBase, uint32(v8716)+4))
	if v8720 <= int32(0) {
		v8799 = v8714
		goto L1207
	} else {
		goto L1209
	}
L1209:
	;
	v8723 = *(*int32)(unsafe.Add(mBase, uint32(v8715)+8))
	v8726 = v8714
	v8727 = int32(1)
	v8729 = v4
	goto L1210
L1210:
	;
	v8768 = *(*int32)(unsafe.Add(mBase, uint32(v8716)+12))
	v8772 = *(*int32)(unsafe.Add(mBase, uint32(v8768+v8729<<(uint(int32(2))%32))))
	v8773 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8773 != 0 {
		goto L1212
	} else {
		goto L1213
	}
L1211:
	;
	v8799 = v8791
	goto L1207
L1212:
	;
	v8774 = F_replace_nestloop_params_mutator(m, v8772, l0)
	mBase = m.M
	v8775 = m.ExcPending
	if v8775 != 0 {
		goto L1
	} else {
		goto L1215
	}
L1213:
	;
	v8776 = v8772
	goto L1214
L1214:
	;
	v8778 = int32(0)
	v8780 = F_makeTargetEntry(m, v8776, base.I32_extend16_s(v8727), v8778, v8778)
	mBase = m.M
	v8781 = m.ExcPending
	if v8781 != 0 {
		goto L1
	} else {
		goto L1216
	}
L1215:
	;
	v8776 = v8774
	goto L1214
L1216:
	;
	if v8723 != 0 {
		goto L1217
	} else {
		goto L1218
	}
L1217:
	;
	v8787 = *(*int32)(unsafe.Add(mBase, uint32(v8723+v8727<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v8780)+16)) = v8787
	goto L1219
L1218:
	;
	goto L1219
L1219:
	;
	v8791 = F_lappend(m, v8726, v8780)
	mBase = m.M
	v8792 = m.ExcPending
	if v8792 != 0 {
		goto L1
	} else {
		goto L1220
	}
L1220:
	;
	v8794 = v8729 + int32(1)
	v8795 = *(*int32)(unsafe.Add(mBase, uint32(v8716)+4))
	if v8794 < v8795 {
		v8726 = v8791
		v8727 = v8727 + int32(1)
		v8729 = v8794
		goto L1210
	} else {
		goto L1221
	}
L1221:
	;
	goto L1211
L1222:
	;
	v8844 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8842)+56)) = v8844
	*(*int32)(unsafe.Add(mBase, uint32(v8842)+52)) = v8712
	*(*int32)(unsafe.Add(mBase, uint32(v8842)+48)) = v8844
	*(*int32)(unsafe.Add(mBase, uint32(v8842)+44)) = v8799
	*(*int32)(unsafe.Add(mBase, uint32(v8842))) = int32(332)
	v8852 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8842)+4)) = v8852
	v8854 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v8842)+8)) = v8854
	v8856 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v8842)+16)) = v8856
	v8858 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v8842)+24)) = v8858
	v8860 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8861 = *(*int32)(unsafe.Add(mBase, uint32(v8860)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8842)+32)) = v8861
	v8863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8842)+36)) = uint8(v8863)
	v8865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v8842)+37)) = uint8(v8865)
	v12573 = v8842
	v12580 = v47
	goto L3
L1223:
	;
	v9793 = F_create_scan_plan(m, l0, l1, l2)
	mBase = m.M
	v9794 = m.ExcPending
	if v9794 != 0 {
		goto L1
	} else {
		goto L1341
	}
L1224:
	;
	v9638 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9639 = *(*int32)(unsafe.Add(mBase, uint32(v9638)+4))
	if v9639 == int32(0) {
		v9723 = v4
		goto L1324
	} else {
		goto L1325
	}
L1225:
	;
	v9248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v9248 == int32(0) {
		goto L1282
	} else {
		goto L1283
	}
L1226:
	;
	v8870 = F_use_physical_tlist(m, l0, l1, l2)
	mBase = m.M
	v8871 = m.ExcPending
	if v8871 != 0 {
		goto L1
	} else {
		goto L1227
	}
L1227:
	;
	v8872 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v8870 != 0 {
		goto L1231
	} else {
		goto L1232
	}
L1228:
	;
	v9246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9202)+37)) = uint8(v9246)
	v12573 = v9202
	v12580 = v47
	goto L3
L1229:
	;
	v9177 = F_palloc0(m, int32(80))
	mBase = m.M
	v9178 = m.ExcPending
	if v9178 != 0 {
		goto L1
	} else {
		goto L1281
	}
L1230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9122)+44)) = v9127
	v9167 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9122)+8)) = v9167
	v9169 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9122)+16)) = v9169
	v9171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9122)+24)) = v9171
	v9173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9174 = *(*int32)(unsafe.Add(mBase, uint32(v9173)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9122)+32)) = v9174
	v9202 = v9122
	goto L1228
L1231:
	;
	v8874 = F_create_plan_recurse(m, l0, v8872, int32(0))
	mBase = m.M
	v8875 = m.ExcPending
	if v8875 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1232:
	;
	goto L1233
L1233:
	;
	v8884 = int32(0)
	v8885 = *(*int32)(unsafe.Add(mBase, uint32(v8872)+4))
	switch v8885 - int32(332) {
	case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
		v8900 = v8884
		goto L1238
	case 2:
		goto L1240
	default:
		goto L1239
	case 23:
		goto L1241
	}
L1234:
	;
	v8876 = *(*int32)(unsafe.Add(mBase, uint32(v8874)+44))
	if l2&int32(4) == int32(0) {
		v9122 = v8874
		v9127 = v8876
		goto L1230
	} else {
		goto L1235
	}
L1235:
	;
	v8881 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_apply_pathtarget_labeling_to_tlist(m, v8876, v8881)
	mBase = m.M
	v8883 = m.ExcPending
	if v8883 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1236:
	;
	v9122 = v8874
	v9127 = v8876
	goto L1230
L1237:
	;
	if v8902 != 0 {
		goto L1243
	} else {
		goto L1244
	}
L1238:
	;
	v8902 = v8900
	goto L1237
L1239:
	;
	v8900 = int32(1)
	goto L1238
L1240:
	;
	v8893 = *(*int32)(unsafe.Add(mBase, uint32(v8872)))
	if v8893 != int32(290) {
		v8900 = v8884
		goto L1238
	} else {
		goto L1242
	}
L1241:
	;
	v8888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8872)+72)))
	v8902 = int32(base.Ui32(v8888&int32(4)) >> (uint(int32(2)) % 32))
	goto L1237
L1242:
	;
	v8896 = *(*int32)(unsafe.Add(mBase, uint32(v8872)+72))
	v8902 = base.B2i32(v8896 == int32(0))
	goto L1237
L1243:
	;
	v8904 = F_create_plan_recurse(m, l0, v8872, int32(8))
	mBase = m.M
	v8905 = m.ExcPending
	if v8905 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1244:
	;
	goto L1245
L1245:
	;
	v8989 = F_create_plan_recurse(m, l0, v8872, int32(0))
	mBase = m.M
	v8990 = m.ExcPending
	if v8990 != 0 {
		goto L1
	} else {
		goto L1263
	}
L1246:
	;
	v8906 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8907 = *(*int32)(unsafe.Add(mBase, uint32(v8906)+4))
	if v8907 == int32(0) {
		goto L1247
	} else {
		goto L1248
	}
L1247:
	;
	v9122 = v8904
	v9127 = v4
	goto L1230
L1248:
	;
	goto L1249
L1249:
	;
	v8911 = *(*int32)(unsafe.Add(mBase, uint32(v8907)+4))
	if v8911 <= int32(0) {
		v9122 = v8904
		v9127 = v4
		goto L1230
	} else {
		goto L1250
	}
L1250:
	;
	v8914 = *(*int32)(unsafe.Add(mBase, uint32(v8906)+8))
	v8917 = int32(1)
	v8920 = v4
	v8921 = v4
	goto L1251
L1251:
	;
	v8959 = *(*int32)(unsafe.Add(mBase, uint32(v8907)+12))
	v8963 = *(*int32)(unsafe.Add(mBase, uint32(v8959+v8921<<(uint(int32(2))%32))))
	v8964 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v8964 != 0 {
		goto L1253
	} else {
		goto L1254
	}
L1252:
	;
	v9122 = v8904
	v9127 = v8982
	goto L1230
L1253:
	;
	v8965 = F_replace_nestloop_params_mutator(m, v8963, l0)
	mBase = m.M
	v8966 = m.ExcPending
	if v8966 != 0 {
		goto L1
	} else {
		goto L1256
	}
L1254:
	;
	v8967 = v8963
	goto L1255
L1255:
	;
	v8969 = int32(0)
	v8971 = F_makeTargetEntry(m, v8967, base.I32_extend16_s(v8917), v8969, v8969)
	mBase = m.M
	v8972 = m.ExcPending
	if v8972 != 0 {
		goto L1
	} else {
		goto L1257
	}
L1256:
	;
	v8967 = v8965
	goto L1255
L1257:
	;
	if v8914 != 0 {
		goto L1258
	} else {
		goto L1259
	}
L1258:
	;
	v8978 = *(*int32)(unsafe.Add(mBase, uint32(v8914+v8917<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v8971)+16)) = v8978
	goto L1260
L1259:
	;
	goto L1260
L1260:
	;
	v8982 = F_lappend(m, v8920, v8971)
	mBase = m.M
	v8983 = m.ExcPending
	if v8983 != 0 {
		goto L1
	} else {
		goto L1261
	}
L1261:
	;
	v8985 = v8921 + int32(1)
	v8986 = *(*int32)(unsafe.Add(mBase, uint32(v8907)+4))
	if v8985 < v8986 {
		v8917 = v8917 + int32(1)
		v8920 = v8982
		v8921 = v8985
		goto L1251
	} else {
		goto L1262
	}
L1262:
	;
	goto L1252
L1263:
	;
	v8991 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8992 = *(*int32)(unsafe.Add(mBase, uint32(v8991)+4))
	if v8992 == int32(0) {
		v9078 = v4
		goto L1264
	} else {
		goto L1265
	}
L1264:
	;
	v9117 = *(*int32)(unsafe.Add(mBase, uint32(v8989)+44))
	v9118 = F_tlist_same_exprs(m, v9078, v9117)
	mBase = m.M
	v9119 = m.ExcPending
	if v9119 != 0 {
		goto L1
	} else {
		goto L1279
	}
L1265:
	;
	v8996 = *(*int32)(unsafe.Add(mBase, uint32(v8992)+4))
	if v8996 <= int32(0) {
		v9078 = v4
		goto L1264
	} else {
		goto L1266
	}
L1266:
	;
	v8999 = *(*int32)(unsafe.Add(mBase, uint32(v8991)+8))
	v9002 = int32(1)
	v9005 = v4
	v9006 = v4
	goto L1267
L1267:
	;
	v9044 = *(*int32)(unsafe.Add(mBase, uint32(v8992)+12))
	v9048 = *(*int32)(unsafe.Add(mBase, uint32(v9044+v9006<<(uint(int32(2))%32))))
	v9049 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9049 != 0 {
		goto L1269
	} else {
		goto L1270
	}
L1268:
	;
	v9078 = v9067
	goto L1264
L1269:
	;
	v9050 = F_replace_nestloop_params_mutator(m, v9048, l0)
	mBase = m.M
	v9051 = m.ExcPending
	if v9051 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1270:
	;
	v9052 = v9048
	goto L1271
L1271:
	;
	v9054 = int32(0)
	v9056 = F_makeTargetEntry(m, v9052, base.I32_extend16_s(v9002), v9054, v9054)
	mBase = m.M
	v9057 = m.ExcPending
	if v9057 != 0 {
		goto L1
	} else {
		goto L1273
	}
L1272:
	;
	v9052 = v9050
	goto L1271
L1273:
	;
	if v8999 != 0 {
		goto L1274
	} else {
		goto L1275
	}
L1274:
	;
	v9063 = *(*int32)(unsafe.Add(mBase, uint32(v8999+v9002<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v9056)+16)) = v9063
	goto L1276
L1275:
	;
	goto L1276
L1276:
	;
	v9067 = F_lappend(m, v9005, v9056)
	mBase = m.M
	v9068 = m.ExcPending
	if v9068 != 0 {
		goto L1
	} else {
		goto L1277
	}
L1277:
	;
	v9070 = v9006 + int32(1)
	v9071 = *(*int32)(unsafe.Add(mBase, uint32(v8992)+4))
	if v9070 < v9071 {
		v9002 = v9002 + int32(1)
		v9005 = v9067
		v9006 = v9070
		goto L1267
	} else {
		goto L1278
	}
L1278:
	;
	goto L1268
L1279:
	;
	if v9118 == int32(0) {
		goto L1229
	} else {
		goto L1280
	}
L1280:
	;
	v9122 = v8989
	v9127 = v9078
	goto L1230
L1281:
	;
	v9179 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9177)+72)) = v9179
	*(*int32)(unsafe.Add(mBase, uint32(v9177)+56)) = v9179
	*(*int32)(unsafe.Add(mBase, uint32(v9177)+52)) = v8989
	*(*int32)(unsafe.Add(mBase, uint32(v9177)+48)) = v9179
	*(*int32)(unsafe.Add(mBase, uint32(v9177)+44)) = v9078
	*(*int32)(unsafe.Add(mBase, uint32(v9177))) = int32(331)
	v9189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9177)+4)) = v9189
	v9191 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9177)+8)) = v9191
	v9193 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9177)+16)) = v9193
	v9195 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9177)+24)) = v9195
	v9197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9198 = *(*int32)(unsafe.Add(mBase, uint32(v9197)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9177)+32)) = v9198
	v9200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9177)+36)) = uint8(v9200)
	v9202 = v9177
	goto L1228
L1282:
	;
	v9481 = int32(0)
	v9482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9483 = *(*int32)(unsafe.Add(mBase, uint32(v9482)+4))
	if v9483 == v9481 {
		v9570 = v9481
		goto L1308
	} else {
		goto L1309
	}
L1283:
	;
	v9251 = *(*int32)(unsafe.Add(mBase, uint32(v9248)+4))
	if v9251 <= int32(0) {
		goto L1282
	} else {
		goto L1284
	}
L1284:
	;
	v9258 = v4
	goto L1285
L1285:
	;
	v9298 = *(*int32)(unsafe.Add(mBase, uint32(v9248)+12))
	v9302 = *(*int32)(unsafe.Add(mBase, uint32(v9298+v9258<<(uint(int32(2))%32))))
	v9303 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+16))
	v9304 = *(*int32)(unsafe.Add(mBase, uint32(v9303)+4))
	v9305 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+20))
	v9306 = F_create_plan(m, v9303, v9305)
	mBase = m.M
	v9307 = m.ExcPending
	if v9307 != 0 {
		goto L1
	} else {
		goto L1287
	}
L1286:
	;
	goto L1282
L1287:
	;
	v9308 = *(*int64)(unsafe.Add(mBase, uint32(v9304)+128))
	v9309 = *(*int32)(unsafe.Add(mBase, uint32(v9304)+136))
	v9311 = F_palloc0(m, int32(104))
	mBase = m.M
	v9312 = m.ExcPending
	if v9312 != 0 {
		goto L1
	} else {
		goto L1288
	}
L1288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9311))) = int32(373)
	v9315 = *(*int32)(unsafe.Add(mBase, uint32(v9306)+44))
	v9316 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9311)+84)) = v9316
	*(*int32)(unsafe.Add(mBase, uint32(v9311)+80)) = v9309
	*(*int64)(unsafe.Add(mBase, uint32(v9311)+72)) = v9308
	v9320 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9311)+56)) = v9320
	*(*int32)(unsafe.Add(mBase, uint32(v9311)+52)) = v9306
	*(*int32)(unsafe.Add(mBase, uint32(v9311)+48)) = v9320
	*(*int32)(unsafe.Add(mBase, uint32(v9311)+44)) = v9315
	*(*int64)(unsafe.Add(mBase, uint32(v9311)+92)) = v9316
	v9328 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+20))
	v9329 = *(*int32)(unsafe.Add(mBase, uint32(v9328)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9311)+4)) = v9329
	v9331 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+20))
	v9332 = *(*float64)(unsafe.Add(mBase, uint32(v9331)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9311)+8)) = v9332
	v9334 = *(*float64)(unsafe.Add(mBase, uint32(v9302)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v9311)+24)) = int64(4607182418800017408)
	*(*float64)(unsafe.Add(mBase, uint32(v9311)+16)) = v9334
	v9338 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+20))
	v9339 = *(*int32)(unsafe.Add(mBase, uint32(v9338)+12))
	v9340 = *(*int32)(unsafe.Add(mBase, uint32(v9339)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v9311)+36)) = uint8(v9320)
	*(*int32)(unsafe.Add(mBase, uint32(v9311)+32)) = v9340
	v9344 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+20))
	v9345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9344)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9311)+37)) = uint8(v9345)
	v9347 = *(*int32)(unsafe.Add(mBase, uint32(v9302)+32))
	v9349 = m.G0
	v9351 = v9349 - int32(32)
	m.G0 = v9351
	v9353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9354 = *(*int32)(unsafe.Add(mBase, uint32(v9353)+8))
	v9355 = F_lappend(m, v9354, v9311)
	mBase = m.M
	v9356 = m.ExcPending
	if v9356 != 0 {
		goto L1
	} else {
		goto L1289
	}
L1289:
	;
	v9357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9357)+8)) = v9355
	v9359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9360 = *(*int32)(unsafe.Add(mBase, uint32(v9359)+12))
	v9362 = F_lappend(m, v9360, int32(0))
	mBase = m.M
	v9363 = m.ExcPending
	if v9363 != 0 {
		goto L1
	} else {
		goto L1290
	}
L1290:
	;
	v9364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9364)+12)) = v9362
	v9366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9367 = *(*int32)(unsafe.Add(mBase, uint32(v9366)+16))
	v9368 = F_lappend(m, v9367, v9303)
	mBase = m.M
	v9369 = m.ExcPending
	if v9369 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1291:
	;
	v9370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9370)+16)) = v9368
	v9373 = F_palloc0(m, int32(72))
	mBase = m.M
	v9374 = m.ExcPending
	if v9374 != 0 {
		goto L1
	} else {
		goto L1292
	}
L1292:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9373))) = int64(17179869207)
	v9377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9378 = *(*int32)(unsafe.Add(mBase, uint32(v9377)+8))
	if v9378 != 0 {
		goto L1293
	} else {
		goto L1294
	}
L1293:
	;
	v9379 = *(*int32)(unsafe.Add(mBase, uint32(v9378)+4))
	v9380 = v9379
	goto L1295
L1294:
	;
	v9380 = v9320
	goto L1295
L1295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9373)+16)) = v9380
	*(*int32)(unsafe.Add(mBase, uint32(v9351)+16)) = v9380
	v9386 = F_psprintf(m, int32(_a_F_create_plan_recurse_44), v9351+int32(16))
	mBase = m.M
	v9387 = m.ExcPending
	if v9387 != 0 {
		goto L1
	} else {
		goto L1296
	}
L1296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9373)+20)) = v9386
	v9389 = *(*int32)(unsafe.Add(mBase, uint32(v9311)+44))
	if v9389 == int32(0) {
		goto L1298
	} else {
		goto L1299
	}
L1297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9373)+32)) = v9411
	v9413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9311)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9373)+38)) = uint8(v9413)
	v9415 = *(*int32)(unsafe.Add(mBase, uint32(v9347)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9351)+12)) = v9415
	*(*int32)(unsafe.Add(mBase, uint32(v9351)+28)) = v9415
	v9421 = F_list_make1_impl(m, int32(471), v9351+int32(12))
	mBase = m.M
	v9422 = m.ExcPending
	if v9422 != 0 {
		goto L1
	} else {
		goto L1304
	}
L1298:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9373)+24)) = int64(-4294965018)
	v9411 = int32(0)
	goto L1297
L1299:
	;
	v9392 = *(*int32)(unsafe.Add(mBase, uint32(v9389)+12))
	v9393 = *(*int32)(unsafe.Add(mBase, uint32(v9392)))
	v9394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9393)+26)))
	if v9394 != 0 {
		goto L1298
	} else {
		goto L1300
	}
L1300:
	;
	v9395 = *(*int32)(unsafe.Add(mBase, uint32(v9393)+4))
	v9396 = F_exprType(m, v9395)
	mBase = m.M
	v9397 = m.ExcPending
	if v9397 != 0 {
		goto L1
	} else {
		goto L1301
	}
L1301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9373)+24)) = v9396
	v9399 = *(*int32)(unsafe.Add(mBase, uint32(v9393)+4))
	v9400 = F_exprTypmod(m, v9399)
	mBase = m.M
	v9401 = m.ExcPending
	if v9401 != 0 {
		goto L1
	} else {
		goto L1302
	}
L1302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9373)+28)) = v9400
	v9403 = *(*int32)(unsafe.Add(mBase, uint32(v9393)+4))
	v9404 = F_exprCollation(m, v9403)
	mBase = m.M
	v9405 = m.ExcPending
	if v9405 != 0 {
		goto L1
	} else {
		goto L1303
	}
L1303:
	;
	v9411 = v9404
	goto L1297
L1304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9373)+40)) = v9421
	v9424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v9425 = F_lappend(m, v9424, v9373)
	mBase = m.M
	v9426 = m.ExcPending
	if v9426 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v9425
	F_cost_subplan(m, v9373, v9311)
	mBase = m.M
	v9429 = m.ExcPending
	if v9429 != 0 {
		goto L1
	} else {
		goto L1306
	}
L1306:
	;
	m.G0 = v9351 + int32(32)
	v9434 = v9258 + int32(1)
	v9435 = *(*int32)(unsafe.Add(mBase, uint32(v9248)+4))
	if v9434 < v9435 {
		v9258 = v9434
		goto L1285
	} else {
		goto L1307
	}
L1307:
	;
	goto L1286
L1308:
	;
	v9609 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v9611 = F_palloc0(m, int32(80))
	mBase = m.M
	v9612 = m.ExcPending
	if v9612 != 0 {
		goto L1
	} else {
		goto L1323
	}
L1309:
	;
	v9487 = *(*int32)(unsafe.Add(mBase, uint32(v9483)+4))
	if v9487 <= int32(0) {
		v9570 = v9481
		goto L1308
	} else {
		goto L1310
	}
L1310:
	;
	v9490 = *(*int32)(unsafe.Add(mBase, uint32(v9482)+8))
	v9494 = int32(1)
	v9496 = int32(0)
	v9497 = v9481
	goto L1311
L1311:
	;
	v9536 = *(*int32)(unsafe.Add(mBase, uint32(v9483)+12))
	v9540 = *(*int32)(unsafe.Add(mBase, uint32(v9536+v9496<<(uint(int32(2))%32))))
	v9541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9541 != 0 {
		goto L1313
	} else {
		goto L1314
	}
L1312:
	;
	v9570 = v9559
	goto L1308
L1313:
	;
	v9542 = F_replace_nestloop_params_mutator(m, v9540, l0)
	mBase = m.M
	v9543 = m.ExcPending
	if v9543 != 0 {
		goto L1
	} else {
		goto L1316
	}
L1314:
	;
	v9544 = v9540
	goto L1315
L1315:
	;
	v9546 = int32(0)
	v9548 = F_makeTargetEntry(m, v9544, base.I32_extend16_s(v9494), v9546, v9546)
	mBase = m.M
	v9549 = m.ExcPending
	if v9549 != 0 {
		goto L1
	} else {
		goto L1317
	}
L1316:
	;
	v9544 = v9542
	goto L1315
L1317:
	;
	if v9490 != 0 {
		goto L1318
	} else {
		goto L1319
	}
L1318:
	;
	v9555 = *(*int32)(unsafe.Add(mBase, uint32(v9490+v9494<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v9548)+16)) = v9555
	goto L1320
L1319:
	;
	goto L1320
L1320:
	;
	v9559 = F_lappend(m, v9497, v9548)
	mBase = m.M
	v9560 = m.ExcPending
	if v9560 != 0 {
		goto L1
	} else {
		goto L1321
	}
L1321:
	;
	v9562 = v9496 + int32(1)
	v9563 = *(*int32)(unsafe.Add(mBase, uint32(v9483)+4))
	if v9562 < v9563 {
		v9494 = v9494 + int32(1)
		v9496 = v9562
		v9497 = v9559
		goto L1311
	} else {
		goto L1322
	}
L1322:
	;
	goto L1312
L1323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9611)+72)) = v9609
	*(*int32)(unsafe.Add(mBase, uint32(v9611)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9611)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9611)+44)) = v9570
	*(*int32)(unsafe.Add(mBase, uint32(v9611))) = int32(331)
	v9621 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9611)+4)) = v9621
	v9623 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9611)+8)) = v9623
	v9625 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9611)+16)) = v9625
	v9627 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9611)+24)) = v9627
	v9629 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9630 = *(*int32)(unsafe.Add(mBase, uint32(v9629)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9611)+32)) = v9630
	v9632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9611)+36)) = uint8(v9632)
	v9634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9611)+37)) = uint8(v9634)
	v9636 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v9636
	v12573 = v9611
	v12580 = v47
	goto L3
L1324:
	;
	v9764 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v9765 = F_order_qual_clauses(m, l0, v9764)
	mBase = m.M
	v9766 = m.ExcPending
	if v9766 != 0 {
		goto L1
	} else {
		goto L1339
	}
L1325:
	;
	v9643 = *(*int32)(unsafe.Add(mBase, uint32(v9639)+4))
	if v9643 <= int32(0) {
		v9723 = v4
		goto L1324
	} else {
		goto L1326
	}
L1326:
	;
	v9646 = *(*int32)(unsafe.Add(mBase, uint32(v9638)+8))
	v9649 = int32(1)
	v9650 = v4
	v9651 = v4
	goto L1327
L1327:
	;
	v9691 = *(*int32)(unsafe.Add(mBase, uint32(v9639)+12))
	v9695 = *(*int32)(unsafe.Add(mBase, uint32(v9691+v9651<<(uint(int32(2))%32))))
	v9696 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9696 != 0 {
		goto L1329
	} else {
		goto L1330
	}
L1328:
	;
	v9723 = v9714
	goto L1324
L1329:
	;
	v9697 = F_replace_nestloop_params_mutator(m, v9695, l0)
	mBase = m.M
	v9698 = m.ExcPending
	if v9698 != 0 {
		goto L1
	} else {
		goto L1332
	}
L1330:
	;
	v9699 = v9695
	goto L1331
L1331:
	;
	v9701 = int32(0)
	v9703 = F_makeTargetEntry(m, v9699, base.I32_extend16_s(v9649), v9701, v9701)
	mBase = m.M
	v9704 = m.ExcPending
	if v9704 != 0 {
		goto L1
	} else {
		goto L1333
	}
L1332:
	;
	v9699 = v9697
	goto L1331
L1333:
	;
	if v9646 != 0 {
		goto L1334
	} else {
		goto L1335
	}
L1334:
	;
	v9710 = *(*int32)(unsafe.Add(mBase, uint32(v9646+v9649<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v9703)+16)) = v9710
	goto L1336
L1335:
	;
	goto L1336
L1336:
	;
	v9714 = F_lappend(m, v9650, v9703)
	mBase = m.M
	v9715 = m.ExcPending
	if v9715 != 0 {
		goto L1
	} else {
		goto L1337
	}
L1337:
	;
	v9717 = v9651 + int32(1)
	v9718 = *(*int32)(unsafe.Add(mBase, uint32(v9639)+4))
	if v9717 < v9718 {
		v9649 = v9649 + int32(1)
		v9650 = v9714
		v9651 = v9717
		goto L1327
	} else {
		goto L1338
	}
L1338:
	;
	goto L1328
L1339:
	;
	v9768 = F_palloc0(m, int32(80))
	mBase = m.M
	v9769 = m.ExcPending
	if v9769 != 0 {
		goto L1
	} else {
		goto L1340
	}
L1340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+72)) = v9765
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9768)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+44)) = v9723
	*(*int32)(unsafe.Add(mBase, uint32(v9768))) = int32(331)
	v9778 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+4)) = v9778
	v9780 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9768)+8)) = v9780
	v9782 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9768)+16)) = v9782
	v9784 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9768)+24)) = v9784
	v9786 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9787 = *(*int32)(unsafe.Add(mBase, uint32(v9786)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9768)+32)) = v9787
	v9789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9768)+36)) = uint8(v9789)
	v9791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9768)+37)) = uint8(v9791)
	v12573 = v9768
	v12580 = v47
	goto L3
L1341:
	;
	v12573 = v9793
	v12580 = v47
	goto L3
L1342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9796))) = int32(335)
	v9800 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9801 = *(*int32)(unsafe.Add(mBase, uint32(v9800)+4))
	if v9801 == int32(0) {
		goto L1344
	} else {
		goto L1345
	}
L1343:
	;
	v9977 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v9978 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9979 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+4)) = v9979
	v9981 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v9796)+8)) = v9981
	v9983 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v9796)+16)) = v9983
	v9985 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v9796)+24)) = v9985
	v9987 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9988 = *(*int32)(unsafe.Add(mBase, uint32(v9987)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+32)) = v9988
	v9990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9796)+36)) = uint8(v9990)
	v9992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	v9993 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+56)) = v9993
	*(*int64)(unsafe.Add(mBase, uint32(v9796)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+44)) = v9943
	*(*uint8)(unsafe.Add(mBase, uint32(v9796)+37)) = uint8(v9992)
	v9999 = *(*int32)(unsafe.Add(mBase, uint32(v9978)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+72)) = v9999
	v10001 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10002 = *(*int32)(unsafe.Add(mBase, uint32(v10001)+8))
	v10008 = v9796 + int32(84)
	v10015 = F_prepare_sort_from_pathkeys(m, v9796, v9977, v10002, v9993, int32(1), v9796+int32(80), v10008, v9796+int32(88), v9796+int32(92), v9796+int32(96))
	mBase = m.M
	v10016 = m.ExcPending
	if v10016 != 0 {
		goto L1
	} else {
		goto L1360
	}
L1344:
	;
	v9930 = int32(0)
	v9943 = v9930
	v9976 = v9930
	goto L1343
L1345:
	;
	v9805 = int32(0)
	v9806 = *(*int32)(unsafe.Add(mBase, uint32(v9801)+4))
	if v9806 <= v9805 {
		v9943 = v4
		v9976 = v9805
		goto L1343
	} else {
		goto L1346
	}
L1346:
	;
	v9809 = *(*int32)(unsafe.Add(mBase, uint32(v9800)+8))
	v9814 = int32(1)
	v9818 = v4
	v9821 = v4
	goto L1347
L1347:
	;
	v9854 = *(*int32)(unsafe.Add(mBase, uint32(v9801)+12))
	v9858 = *(*int32)(unsafe.Add(mBase, uint32(v9854+v9818<<(uint(int32(2))%32))))
	v9859 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9859 != 0 {
		goto L1349
	} else {
		goto L1350
	}
L1348:
	;
	if v9877 == int32(0) {
		goto L1344
	} else {
		goto L1359
	}
L1349:
	;
	v9860 = F_replace_nestloop_params_mutator(m, v9858, l0)
	mBase = m.M
	v9861 = m.ExcPending
	if v9861 != 0 {
		goto L1
	} else {
		goto L1352
	}
L1350:
	;
	v9862 = v9858
	goto L1351
L1351:
	;
	v9864 = int32(0)
	v9866 = F_makeTargetEntry(m, v9862, base.I32_extend16_s(v9814), v9864, v9864)
	mBase = m.M
	v9867 = m.ExcPending
	if v9867 != 0 {
		goto L1
	} else {
		goto L1353
	}
L1352:
	;
	v9862 = v9860
	goto L1351
L1353:
	;
	if v9809 != 0 {
		goto L1354
	} else {
		goto L1355
	}
L1354:
	;
	v9873 = *(*int32)(unsafe.Add(mBase, uint32(v9809+v9814<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v9866)+16)) = v9873
	goto L1356
L1355:
	;
	goto L1356
L1356:
	;
	v9877 = F_lappend(m, v9821, v9866)
	mBase = m.M
	v9878 = m.ExcPending
	if v9878 != 0 {
		goto L1
	} else {
		goto L1357
	}
L1357:
	;
	v9880 = v9818 + int32(1)
	v9881 = *(*int32)(unsafe.Add(mBase, uint32(v9801)+4))
	if v9880 < v9881 {
		v9814 = v9814 + int32(1)
		v9818 = v9880
		v9821 = v9877
		goto L1347
	} else {
		goto L1358
	}
L1358:
	;
	goto L1348
L1359:
	;
	v9885 = *(*int32)(unsafe.Add(mBase, uint32(v9877)+4))
	v9943 = v9877
	v9976 = v9885
	goto L1343
L1360:
	;
	v10017 = *(*int32)(unsafe.Add(mBase, uint32(v9796)+44))
	if v10017 != 0 {
		goto L1361
	} else {
		goto L1362
	}
L1361:
	;
	v10018 = *(*int32)(unsafe.Add(mBase, uint32(v10017)+4))
	v10019 = v10018
	goto L1363
L1362:
	;
	v10019 = v4
	goto L1363
L1363:
	;
	v10020 = int32(0)
	v10021 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10021 == v10020 {
		v10313 = v10020
		goto L1365
	} else {
		goto L1366
	}
L1364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10379 = m.ExcPending
	if v10379 != 0 {
		goto L1
	} else {
		goto L1425
	}
L1365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+100)) = int32(-1)
	v10349 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[5])))
	if v10349 == int32(0) {
		goto L1417
	} else {
		goto L1418
	}
L1366:
	;
	v10024 = *(*int32)(unsafe.Add(mBase, uint32(v10021)+4))
	if v10024 <= int32(0) {
		v10313 = v10020
		goto L1365
	} else {
		goto L1367
	}
L1367:
	;
	v10036 = int32(0)
	v10039 = v10020
	goto L1368
L1368:
	;
	v10072 = *(*int32)(unsafe.Add(mBase, uint32(v10021)+12))
	v10076 = *(*int32)(unsafe.Add(mBase, uint32(v10072+v10036<<(uint(int32(2))%32))))
	v10078 = F_create_plan_recurse(m, l0, v10076, int32(1))
	mBase = m.M
	v10079 = m.ExcPending
	if v10079 != 0 {
		goto L1
	} else {
		goto L1370
	}
L1369:
	;
	v10313 = v10296
	goto L1365
L1370:
	;
	v10080 = *(*int32)(unsafe.Add(mBase, uint32(v10076)+8))
	v10081 = *(*int32)(unsafe.Add(mBase, uint32(v10080)+8))
	v10082 = *(*int32)(unsafe.Add(mBase, uint32(v10008)))
	v10094 = F_prepare_sort_from_pathkeys(m, v10078, v9977, v10081, v10082, int32(0), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128), v47+int32(52))
	mBase = m.M
	v10095 = m.ExcPending
	if v10095 != 0 {
		goto L1
	} else {
		goto L1371
	}
L1371:
	;
	v10096 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v10097 = *(*int32)(unsafe.Add(mBase, uint32(v10008)))
	v10098 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v10100 = v10098 << (uint(int32(1)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v10100) {
		goto L1375
	} else {
		goto L1376
	}
L1372:
	;
	if v10162 != 0 {
		goto L1364
	} else {
		goto L1390
	}
L1373:
	;
	v10162 = int32(0)
	goto L1372
L1374:
	;
	v10136 = v10131
	v10137 = v10132
	v10138 = v10133
	goto L1384
L1375:
	;
	if (v10096|v10097)&int32(3) != 0 {
		v10131 = v10096
		v10132 = v10097
		v10133 = v10100
		goto L1374
	} else {
		goto L1378
	}
L1376:
	;
	v10124 = v10096
	v10125 = v10097
	v10126 = v10100
	goto L1377
L1377:
	;
	if v10126 == int32(0) {
		goto L1373
	} else {
		goto L1383
	}
L1378:
	;
	v10108 = v10096
	v10109 = v10097
	v10110 = v10100
	goto L1379
L1379:
	;
	v10113 = *(*int32)(unsafe.Add(mBase, uint32(v10108)))
	v10114 = *(*int32)(unsafe.Add(mBase, uint32(v10109)))
	if v10113 != v10114 {
		v10131 = v10108
		v10132 = v10109
		v10133 = v10110
		goto L1374
	} else {
		goto L1381
	}
L1380:
	;
	v10124 = v10119
	v10125 = v10117
	v10126 = v10121
	goto L1377
L1381:
	;
	v10116 = int32(4)
	v10117 = v10109 + v10116
	v10119 = v10108 + v10116
	v10121 = v10110 - v10116
	if base.Ui32(int32(3)) < base.Ui32(v10121) {
		v10108 = v10119
		v10109 = v10117
		v10110 = v10121
		goto L1379
	} else {
		goto L1382
	}
L1382:
	;
	goto L1380
L1383:
	;
	v10131 = v10124
	v10132 = v10125
	v10133 = v10126
	goto L1374
L1384:
	;
	v10141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10136))))
	v10142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10137))))
	if v10141 == v10142 {
		goto L1386
	} else {
		goto L1387
	}
L1385:
	;
	v10162 = v10141 - v10142
	goto L1372
L1386:
	;
	v10144 = int32(1)
	v10149 = v10138 - v10144
	if v10149 != 0 {
		v10136 = v10136 + v10144
		v10137 = v10137 + v10144
		v10138 = v10149
		goto L1384
	} else {
		goto L1389
	}
L1387:
	;
	goto L1388
L1388:
	;
	goto L1385
L1389:
	;
	goto L1373
L1390:
	;
	v10163 = *(*int32)(unsafe.Add(mBase, uint32(v10076)+64))
	if v9977 == v10163 {
		goto L1393
	} else {
		goto L1394
	}
L1391:
	;
	v10296 = F_lappend(m, v10039, v10289)
	mBase = m.M
	v10297 = m.ExcPending
	if v10297 != 0 {
		goto L1
	} else {
		goto L1415
	}
L1392:
	;
	if v10216 != 0 {
		goto L1410
	} else {
		goto L1411
	}
L1393:
	;
	v10216 = int32(1)
	goto L1392
L1394:
	;
	goto L1395
L1395:
	;
	v10172 = int32(0)
	goto L1397
L1396:
	;
	v10216 = v10208
	goto L1392
L1397:
	;
	v10176 = int32(0)
	if v9977 == v10176 {
		v10186 = v10176
		goto L1399
	} else {
		goto L1400
	}
L1398:
	;
	v10208 = int32(0)
	goto L1396
L1399:
	;
	if v10163 != 0 {
		goto L1403
	} else {
		goto L1404
	}
L1400:
	;
	v10180 = *(*int32)(unsafe.Add(mBase, uint32(v9977)+4))
	if v10180 <= v10172 {
		v10186 = int32(0)
		goto L1399
	} else {
		goto L1401
	}
L1401:
	;
	v10182 = *(*int32)(unsafe.Add(mBase, uint32(v9977)+12))
	v10186 = v10182 + v10172<<(uint(int32(2))%32)
	goto L1399
L1402:
	;
	v10192 = base.B2i32(v10186 == int32(0))
	if v10186 == int32(0) {
		v10208 = v10192
		goto L1396
	} else {
		goto L1407
	}
L1403:
	;
	v10187 = *(*int32)(unsafe.Add(mBase, uint32(v10163)+4))
	if v10172 < v10187 {
		goto L1402
	} else {
		goto L1406
	}
L1404:
	;
	goto L1405
L1405:
	;
	v10216 = base.B2i32(v10186 == int32(0))
	goto L1392
L1406:
	;
	goto L1405
L1407:
	;
	v10195 = *(*int32)(unsafe.Add(mBase, uint32(v10163)+12))
	if v10195 == int32(0) {
		v10208 = v10192
		goto L1396
	} else {
		goto L1408
	}
L1408:
	;
	v10202 = *(*int32)(unsafe.Add(mBase, uint32(v10186)))
	v10204 = *(*int32)(unsafe.Add(mBase, uint32(v10172<<(uint(int32(2))%32)+v10195)))
	if v10202 == v10204 {
		v10172 = v10172 + int32(1)
		goto L1397
	} else {
		goto L1409
	}
L1409:
	;
	goto L1398
L1410:
	;
	v10289 = v10094
	goto L1391
L1411:
	;
	goto L1412
L1412:
	;
	v10217 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	v10218 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v10219 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v10221 = F_palloc0(m, int32(96))
	mBase = m.M
	v10222 = m.ExcPending
	if v10222 != 0 {
		goto L1
	} else {
		goto L1413
	}
L1413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10221))) = int32(362)
	v10225 = *(*int32)(unsafe.Add(mBase, uint32(v10094)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+44)) = v10225
	v10227 = *(*int32)(unsafe.Add(mBase, uint32(v10094)+4))
	v10228 = int32(_a_F_create_plan_recurse_0)
	v10229 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+88)) = v10219
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+84)) = v10218
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+80)) = v10217
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+76)) = v10096
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+72)) = v10098
	v10235 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+56)) = v10235
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+52)) = v10094
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+48)) = v10235
	v10240 = int32(1)
	v10242 = v10227 + (v10229 ^ v10240)
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+4)) = v10242
	v10245 = v47 + int32(56)
	v10246 = *(*float64)(unsafe.Add(mBase, uint32(v10094)+16))
	v10247 = *(*float64)(unsafe.Add(mBase, uint32(v10094)+24))
	v10248 = *(*int32)(unsafe.Add(mBase, uint32(v10094)+32))
	v10251 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v10252 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v10254 = m.G0
	v10255 = int32(16)
	v10256 = v10254 - v10255
	m.G0 = v10256
	F_cost_tuplesort(m, v10256+int32(8), v10256, v10247, v10248, float64(0), v10251, v10252)
	mBase = m.M
	v10261 = *(*float64)(unsafe.Add(mBase, uint32(v10256)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v10245)+32)) = v10247
	v10264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v10265 = base.F64_add(v10246, v10261)
	*(*float64)(unsafe.Add(mBase, uint32(v10245)+48)) = v10265
	*(*int32)(unsafe.Add(mBase, uint32(v10245)+40)) = v10242 + (v10264 ^ v10240)
	v10271 = *(*float64)(unsafe.Add(mBase, uint32(v10256)))
	*(*float64)(unsafe.Add(mBase, uint32(v10245)+56)) = base.F64_add(v10265, v10271)
	m.G0 = v10256 + v10255
	goto L1414
L1414:
	;
	v10277 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v10221)+8)) = v10277
	v10279 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v10221)+16)) = v10279
	v10281 = *(*float64)(unsafe.Add(mBase, uint32(v10094)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v10221)+24)) = v10281
	v10283 = *(*int32)(unsafe.Add(mBase, uint32(v10094)+32))
	v10284 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10221)+36)) = uint8(v10284)
	*(*int32)(unsafe.Add(mBase, uint32(v10221)+32)) = v10283
	v10287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10094)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10221)+37)) = uint8(v10287)
	v10289 = v10221
	goto L1391
L1415:
	;
	v10299 = v10036 + int32(1)
	v10300 = *(*int32)(unsafe.Add(mBase, uint32(v10021)+4))
	if v10299 < v10300 {
		v10036 = v10299
		v10039 = v10296
		goto L1368
	} else {
		goto L1416
	}
L1416:
	;
	goto L1369
L1417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+76)) = v10313
	if base.B2i32(l2&int32(3) == int32(0))|base.B2i32(v9976 == v10019) != 0 {
		v12573 = v9796
		v12580 = v47
		goto L3
	} else {
		goto L1422
	}
L1418:
	;
	v10352 = *(*int32)(unsafe.Add(mBase, uint32(v9978)+184))
	v10354 = F_extract_actual_clauses(m, v10352, int32(0))
	mBase = m.M
	v10355 = m.ExcPending
	if v10355 != 0 {
		goto L1
	} else {
		goto L1419
	}
L1419:
	;
	if v10354 == int32(0) {
		goto L1417
	} else {
		goto L1420
	}
L1420:
	;
	v10358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v10359 = F_make_partition_pruneinfo(m, l0, v9978, v10358, v10354)
	mBase = m.M
	v10360 = m.ExcPending
	if v10360 != 0 {
		goto L1
	} else {
		goto L1421
	}
L1421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9796)+100)) = v10359
	goto L1417
L1422:
	;
	v10370 = *(*int32)(unsafe.Add(mBase, uint32(v9796)+44))
	v10371 = F_list_copy_head(m, v10370, v9976)
	mBase = m.M
	v10372 = m.ExcPending
	if v10372 != 0 {
		goto L1
	} else {
		goto L1423
	}
L1423:
	;
	v10373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9796)+37)))
	v10374 = F_inject_projection_plan(m, v9796, v10371, v10373)
	mBase = m.M
	v10375 = m.ExcPending
	if v10375 != 0 {
		goto L1
	} else {
		goto L1424
	}
L1424:
	;
	v12573 = v10374
	v12580 = v47
	goto L3
L1425:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_45), int32(0))
	mBase = m.M
	v10383 = m.ExcPending
	if v10383 != 0 {
		goto L1
	} else {
		goto L1426
	}
L1426:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1519), int32(_a_F_create_plan_recurse_46))
	mBase = m.M
	v10388 = m.ExcPending
	if v10388 != 0 {
		goto L1
	} else {
		goto L1427
	}
L1427:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1428:
	;
	v10566 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10567 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v10568 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = v10568
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v10568
	*(*int32)(unsafe.Add(mBase, uint32(v47)+132)) = v10568
	*(*int32)(unsafe.Add(mBase, uint32(v47)+128)) = v10568
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v10568
	v10578 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10578 == v10568 {
		goto L1445
	} else {
		goto L1446
	}
L1429:
	;
	v10519 = int32(0)
	v10529 = v10519
	v10565 = v10519
	goto L1428
L1430:
	;
	v10394 = int32(0)
	v10395 = *(*int32)(unsafe.Add(mBase, uint32(v10390)+4))
	if v10395 <= v10394 {
		v10529 = v4
		v10565 = v10394
		goto L1428
	} else {
		goto L1431
	}
L1431:
	;
	v10398 = *(*int32)(unsafe.Add(mBase, uint32(v10389)+8))
	v10402 = int32(1)
	v10404 = v4
	v10407 = v4
	goto L1432
L1432:
	;
	v10443 = *(*int32)(unsafe.Add(mBase, uint32(v10390)+12))
	v10447 = *(*int32)(unsafe.Add(mBase, uint32(v10443+v10404<<(uint(int32(2))%32))))
	v10448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v10448 != 0 {
		goto L1434
	} else {
		goto L1435
	}
L1433:
	;
	if v10466 == int32(0) {
		goto L1429
	} else {
		goto L1444
	}
L1434:
	;
	v10449 = F_replace_nestloop_params_mutator(m, v10447, l0)
	mBase = m.M
	v10450 = m.ExcPending
	if v10450 != 0 {
		goto L1
	} else {
		goto L1437
	}
L1435:
	;
	v10451 = v10447
	goto L1436
L1436:
	;
	v10453 = int32(0)
	v10455 = F_makeTargetEntry(m, v10451, base.I32_extend16_s(v10402), v10453, v10453)
	mBase = m.M
	v10456 = m.ExcPending
	if v10456 != 0 {
		goto L1
	} else {
		goto L1438
	}
L1437:
	;
	v10451 = v10449
	goto L1436
L1438:
	;
	if v10398 != 0 {
		goto L1439
	} else {
		goto L1440
	}
L1439:
	;
	v10462 = *(*int32)(unsafe.Add(mBase, uint32(v10398+v10402<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v10455)+16)) = v10462
	goto L1441
L1440:
	;
	goto L1441
L1441:
	;
	v10466 = F_lappend(m, v10407, v10455)
	mBase = m.M
	v10467 = m.ExcPending
	if v10467 != 0 {
		goto L1
	} else {
		goto L1442
	}
L1442:
	;
	v10469 = v10404 + int32(1)
	v10470 = *(*int32)(unsafe.Add(mBase, uint32(v10390)+4))
	if v10469 < v10470 {
		v10402 = v10402 + int32(1)
		v10404 = v10469
		v10407 = v10466
		goto L1432
	} else {
		goto L1443
	}
L1443:
	;
	goto L1433
L1444:
	;
	v10474 = *(*int32)(unsafe.Add(mBase, uint32(v10466)+4))
	v10529 = v10466
	v10565 = v10474
	goto L1428
L1445:
	;
	v10581 = int32(0)
	v10583 = F_makeBoolConst(m, v10581, v10581)
	mBase = m.M
	v10584 = m.ExcPending
	if v10584 != 0 {
		goto L1
	} else {
		goto L1448
	}
L1446:
	;
	goto L1447
L1447:
	;
	v10619 = F_palloc0(m, int32(96))
	mBase = m.M
	v10620 = m.ExcPending
	if v10620 != 0 {
		goto L1
	} else {
		goto L1451
	}
L1448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v10583
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v10583
	v10590 = F_list_make1_impl(m, int32(1), v47+int32(28))
	mBase = m.M
	v10591 = m.ExcPending
	if v10591 != 0 {
		goto L1
	} else {
		goto L1449
	}
L1449:
	;
	v10593 = F_palloc0(m, int32(80))
	mBase = m.M
	v10594 = m.ExcPending
	if v10594 != 0 {
		goto L1
	} else {
		goto L1450
	}
L1450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10593)+72)) = v10590
	*(*int32)(unsafe.Add(mBase, uint32(v10593)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10593)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10593)+44)) = v10529
	*(*int32)(unsafe.Add(mBase, uint32(v10593))) = int32(331)
	v10603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10593)+4)) = v10603
	v10605 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v10593)+8)) = v10605
	v10607 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v10593)+16)) = v10607
	v10609 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v10593)+24)) = v10609
	v10611 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10612 = *(*int32)(unsafe.Add(mBase, uint32(v10611)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10593)+32)) = v10612
	v10614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10593)+36)) = uint8(v10614)
	v10616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10593)+37)) = uint8(v10616)
	v12573 = v10593
	v12580 = v47
	goto L3
L1451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10619)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+44)) = v10529
	*(*int32)(unsafe.Add(mBase, uint32(v10619))) = int32(334)
	v10628 = *(*int32)(unsafe.Add(mBase, uint32(v10566)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+72)) = v10628
	if v10567 != 0 {
		goto L1454
	} else {
		goto L1455
	}
L1452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+88)) = int32(-1)
	v11014 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[5])))
	if v11014 == int32(0) {
		goto L1524
	} else {
		goto L1525
	}
L1453:
	;
	v10666 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10666 == int32(0) {
		v10976 = v4
		v10979 = v4
		v10983 = v10665
		goto L1452
	} else {
		goto L1464
	}
L1454:
	;
	v10630 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10631 = *(*int32)(unsafe.Add(mBase, uint32(v10630)+8))
	v10644 = F_prepare_sort_from_pathkeys(m, v10619, v10567, v10631, int32(0), int32(1), v47+int32(140), v47+int32(136), v47+int32(132), v47+int32(128), v47+int32(52))
	mBase = m.M
	v10645 = m.ExcPending
	if v10645 != 0 {
		goto L1
	} else {
		goto L1457
	}
L1455:
	;
	goto L1456
L1456:
	;
	v10651 = int32(1)
	v10653 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[6])))
	if v10653 != v10651 {
		v10664 = v4
		v10665 = v10651
		goto L1453
	} else {
		goto L1461
	}
L1457:
	;
	v10646 = *(*int32)(unsafe.Add(mBase, uint32(v10619)+44))
	if v10646 != 0 {
		goto L1458
	} else {
		goto L1459
	}
L1458:
	;
	v10647 = *(*int32)(unsafe.Add(mBase, uint32(v10646)+4))
	v10649 = v10647
	goto L1460
L1459:
	;
	v10649 = int32(0)
	goto L1460
L1460:
	;
	v10664 = v4
	v10665 = base.B2i32(v10649 == v10565)
	goto L1453
L1461:
	;
	v10656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v10656 != 0 {
		v10664 = v4
		v10665 = v10651
		goto L1453
	} else {
		goto L1462
	}
L1462:
	;
	v10657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v10657 == int32(0) {
		v10976 = v4
		v10979 = v4
		v10983 = v10651
		goto L1452
	} else {
		goto L1463
	}
L1463:
	;
	v10660 = *(*int32)(unsafe.Add(mBase, uint32(v10657)+4))
	v10664 = base.B2i32(int32(1) < v10660)
	v10665 = v10651
	goto L1453
L1464:
	;
	v10669 = *(*int32)(unsafe.Add(mBase, uint32(v10666)+4))
	if v10669 <= int32(0) {
		v10976 = v4
		v10979 = v4
		v10983 = v10665
		goto L1452
	} else {
		goto L1465
	}
L1465:
	;
	v10681 = int32(0)
	v10682 = v4
	v10685 = v4
	goto L1466
L1466:
	;
	v10717 = *(*int32)(unsafe.Add(mBase, uint32(v10666)+12))
	v10721 = *(*int32)(unsafe.Add(mBase, uint32(v10717+v10681<<(uint(int32(2))%32))))
	v10723 = F_create_plan_recurse(m, l0, v10721, int32(1))
	mBase = m.M
	v10724 = m.ExcPending
	if v10724 != 0 {
		goto L1
	} else {
		goto L1469
	}
L1467:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10957 = m.ExcPending
	if v10957 != 0 {
		goto L1
	} else {
		goto L1521
	}
L1468:
	;
	goto L1467
L1469:
	;
	if v10567 == int32(0) {
		v10935 = v10723
		goto L1470
	} else {
		goto L1471
	}
L1470:
	;
	if v10664 != 0 {
		goto L1515
	} else {
		goto L1516
	}
L1471:
	;
	v10727 = *(*int32)(unsafe.Add(mBase, uint32(v10721)+8))
	v10728 = *(*int32)(unsafe.Add(mBase, uint32(v10727)+8))
	v10729 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v10741 = F_prepare_sort_from_pathkeys(m, v10723, v10567, v10728, v10729, int32(0), v47+int32(48), v47+int32(44), v47+int32(40), v47+int32(36), v47+int32(32))
	mBase = m.M
	v10742 = m.ExcPending
	if v10742 != 0 {
		goto L1
	} else {
		goto L1472
	}
L1472:
	;
	v10743 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v10744 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	v10746 = v10744 << (uint(int32(1)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v10746) {
		goto L1476
	} else {
		goto L1477
	}
L1473:
	;
	if v10808 != 0 {
		goto L1468
	} else {
		goto L1491
	}
L1474:
	;
	v10808 = int32(0)
	goto L1473
L1475:
	;
	v10782 = v10777
	v10783 = v10778
	v10784 = v10779
	goto L1485
L1476:
	;
	if (v10743|v10729)&int32(3) != 0 {
		v10777 = v10743
		v10778 = v10729
		v10779 = v10746
		goto L1475
	} else {
		goto L1479
	}
L1477:
	;
	v10770 = v10743
	v10771 = v10729
	v10772 = v10746
	goto L1478
L1478:
	;
	if v10772 == int32(0) {
		goto L1474
	} else {
		goto L1484
	}
L1479:
	;
	v10754 = v10743
	v10755 = v10729
	v10756 = v10746
	goto L1480
L1480:
	;
	v10759 = *(*int32)(unsafe.Add(mBase, uint32(v10754)))
	v10760 = *(*int32)(unsafe.Add(mBase, uint32(v10755)))
	if v10759 != v10760 {
		v10777 = v10754
		v10778 = v10755
		v10779 = v10756
		goto L1475
	} else {
		goto L1482
	}
L1481:
	;
	v10770 = v10765
	v10771 = v10763
	v10772 = v10767
	goto L1478
L1482:
	;
	v10762 = int32(4)
	v10763 = v10755 + v10762
	v10765 = v10754 + v10762
	v10767 = v10756 - v10762
	if base.Ui32(int32(3)) < base.Ui32(v10767) {
		v10754 = v10765
		v10755 = v10763
		v10756 = v10767
		goto L1480
	} else {
		goto L1483
	}
L1483:
	;
	goto L1481
L1484:
	;
	v10777 = v10770
	v10778 = v10771
	v10779 = v10772
	goto L1475
L1485:
	;
	v10787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10782))))
	v10788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10783))))
	if v10787 == v10788 {
		goto L1487
	} else {
		goto L1488
	}
L1486:
	;
	v10808 = v10787 - v10788
	goto L1473
L1487:
	;
	v10790 = int32(1)
	v10795 = v10784 - v10790
	if v10795 != 0 {
		v10782 = v10782 + v10790
		v10783 = v10783 + v10790
		v10784 = v10795
		goto L1485
	} else {
		goto L1490
	}
L1488:
	;
	goto L1489
L1489:
	;
	goto L1486
L1490:
	;
	goto L1474
L1491:
	;
	v10809 = *(*int32)(unsafe.Add(mBase, uint32(v10721)+64))
	if v10567 == v10809 {
		goto L1493
	} else {
		goto L1494
	}
L1492:
	;
	if v10862 != 0 {
		goto L1510
	} else {
		goto L1511
	}
L1493:
	;
	v10862 = int32(1)
	goto L1492
L1494:
	;
	goto L1495
L1495:
	;
	v10818 = int32(0)
	goto L1497
L1496:
	;
	v10862 = v10854
	goto L1492
L1497:
	;
	v10822 = int32(0)
	if v10567 == v10822 {
		v10832 = v10822
		goto L1499
	} else {
		goto L1500
	}
L1498:
	;
	v10854 = int32(0)
	goto L1496
L1499:
	;
	if v10809 != 0 {
		goto L1503
	} else {
		goto L1504
	}
L1500:
	;
	v10826 = *(*int32)(unsafe.Add(mBase, uint32(v10567)+4))
	if v10826 <= v10818 {
		v10832 = int32(0)
		goto L1499
	} else {
		goto L1501
	}
L1501:
	;
	v10828 = *(*int32)(unsafe.Add(mBase, uint32(v10567)+12))
	v10832 = v10828 + v10818<<(uint(int32(2))%32)
	goto L1499
L1502:
	;
	v10838 = base.B2i32(v10832 == int32(0))
	if v10832 == int32(0) {
		v10854 = v10838
		goto L1496
	} else {
		goto L1507
	}
L1503:
	;
	v10833 = *(*int32)(unsafe.Add(mBase, uint32(v10809)+4))
	if v10818 < v10833 {
		goto L1502
	} else {
		goto L1506
	}
L1504:
	;
	goto L1505
L1505:
	;
	v10862 = base.B2i32(v10832 == int32(0))
	goto L1492
L1506:
	;
	goto L1505
L1507:
	;
	v10841 = *(*int32)(unsafe.Add(mBase, uint32(v10809)+12))
	if v10841 == int32(0) {
		v10854 = v10838
		goto L1496
	} else {
		goto L1508
	}
L1508:
	;
	v10848 = *(*int32)(unsafe.Add(mBase, uint32(v10832)))
	v10850 = *(*int32)(unsafe.Add(mBase, uint32(v10818<<(uint(int32(2))%32)+v10841)))
	if v10848 == v10850 {
		v10818 = v10818 + int32(1)
		goto L1497
	} else {
		goto L1509
	}
L1509:
	;
	goto L1498
L1510:
	;
	v10935 = v10741
	goto L1470
L1511:
	;
	goto L1512
L1512:
	;
	v10863 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	v10864 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v10865 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	v10867 = F_palloc0(m, int32(96))
	mBase = m.M
	v10868 = m.ExcPending
	if v10868 != 0 {
		goto L1
	} else {
		goto L1513
	}
L1513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10867))) = int32(362)
	v10871 = *(*int32)(unsafe.Add(mBase, uint32(v10741)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+44)) = v10871
	v10873 = *(*int32)(unsafe.Add(mBase, uint32(v10741)+4))
	v10874 = int32(_a_F_create_plan_recurse_0)
	v10875 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+88)) = v10865
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+84)) = v10864
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+80)) = v10863
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+76)) = v10743
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+72)) = v10744
	v10881 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+56)) = v10881
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+52)) = v10741
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+48)) = v10881
	v10886 = int32(1)
	v10888 = v10873 + (v10875 ^ v10886)
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+4)) = v10888
	v10891 = v47 + int32(56)
	v10892 = *(*float64)(unsafe.Add(mBase, uint32(v10741)+16))
	v10893 = *(*float64)(unsafe.Add(mBase, uint32(v10741)+24))
	v10894 = *(*int32)(unsafe.Add(mBase, uint32(v10741)+32))
	v10897 = *(*int32)(unsafe.Add(mBase, _c_F_create_plan_recurse[1]))
	v10898 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v10900 = m.G0
	v10901 = int32(16)
	v10902 = v10900 - v10901
	m.G0 = v10902
	F_cost_tuplesort(m, v10902+int32(8), v10902, v10893, v10894, float64(0), v10897, v10898)
	mBase = m.M
	v10907 = *(*float64)(unsafe.Add(mBase, uint32(v10902)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v10891)+32)) = v10893
	v10910 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_plan_recurse[2])))
	v10911 = base.F64_add(v10892, v10907)
	*(*float64)(unsafe.Add(mBase, uint32(v10891)+48)) = v10911
	*(*int32)(unsafe.Add(mBase, uint32(v10891)+40)) = v10888 + (v10910 ^ v10886)
	v10917 = *(*float64)(unsafe.Add(mBase, uint32(v10902)))
	*(*float64)(unsafe.Add(mBase, uint32(v10891)+56)) = base.F64_add(v10911, v10917)
	m.G0 = v10902 + v10901
	goto L1514
L1514:
	;
	v10923 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v10867)+8)) = v10923
	v10925 = *(*float64)(unsafe.Add(mBase, uint32(v47)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v10867)+16)) = v10925
	v10927 = *(*float64)(unsafe.Add(mBase, uint32(v10741)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v10867)+24)) = v10927
	v10929 = *(*int32)(unsafe.Add(mBase, uint32(v10741)+32))
	v10930 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10867)+36)) = uint8(v10930)
	*(*int32)(unsafe.Add(mBase, uint32(v10867)+32)) = v10929
	v10933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10741)+37)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10867)+37)) = uint8(v10933)
	v10935 = v10867
	goto L1470
L1515:
	;
	v10944 = F_mark_async_capable_plan(m, v10935, v10721)
	mBase = m.M
	v10945 = m.ExcPending
	if v10945 != 0 {
		goto L1
	} else {
		goto L1518
	}
L1516:
	;
	v10947 = v10682
	goto L1517
L1517:
	;
	v10948 = F_lappend(m, v10685, v10935)
	mBase = m.M
	v10949 = m.ExcPending
	if v10949 != 0 {
		goto L1
	} else {
		goto L1519
	}
L1518:
	;
	v10947 = v10944 + v10682
	goto L1517
L1519:
	;
	v10951 = v10681 + int32(1)
	v10952 = *(*int32)(unsafe.Add(mBase, uint32(v10666)+4))
	if v10951 < v10952 {
		v10681 = v10951
		v10682 = v10947
		v10685 = v10948
		goto L1466
	} else {
		goto L1520
	}
L1520:
	;
	v10976 = v10947
	v10979 = v10948
	v10983 = v10665
	goto L1452
L1521:
	;
	F_errmsg_internal(m, int32(_a_F_create_plan_recurse_47), int32(0))
	mBase = m.M
	v10961 = m.ExcPending
	if v10961 != 0 {
		goto L1
	} else {
		goto L1522
	}
L1522:
	;
	F_errfinish(m, int32(_a_F_create_plan_recurse_2), int32(1347), int32(_a_F_create_plan_recurse_48))
	mBase = m.M
	v10966 = m.ExcPending
	if v10966 != 0 {
		goto L1
	} else {
		goto L1523
	}
L1523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+80)) = v10976
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+76)) = v10979
	v11041 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+84)) = v11041
	v11043 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+4)) = v11043
	v11045 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v10619)+8)) = v11045
	v11047 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v10619)+16)) = v11047
	v11049 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v10619)+24)) = v11049
	v11051 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11052 = *(*int32)(unsafe.Add(mBase, uint32(v11051)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+32)) = v11052
	v11054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10619)+36)) = uint8(v11054)
	v11056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10619)+37)) = uint8(v11056)
	if base.B2i32(l2&int32(3) == int32(0))|v10983 != 0 {
		v12573 = v10619
		v12580 = v47
		goto L3
	} else {
		goto L1535
	}
L1525:
	;
	v11017 = *(*int32)(unsafe.Add(mBase, uint32(v10566)+184))
	v11019 = F_extract_actual_clauses(m, v11017, int32(0))
	mBase = m.M
	v11020 = m.ExcPending
	if v11020 != 0 {
		goto L1
	} else {
		goto L1526
	}
L1526:
	;
	v11021 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11021 != 0 {
		goto L1527
	} else {
		goto L1528
	}
L1527:
	;
	v11022 = *(*int32)(unsafe.Add(mBase, uint32(v11021)+16))
	v11024 = F_extract_actual_clauses(m, v11022, int32(0))
	mBase = m.M
	v11025 = m.ExcPending
	if v11025 != 0 {
		goto L1
	} else {
		goto L1530
	}
L1528:
	;
	v11030 = v11019
	goto L1529
L1529:
	;
	if v11030 == int32(0) {
		goto L1524
	} else {
		goto L1533
	}
L1530:
	;
	v11026 = F_replace_nestloop_params_mutator(m, v11024, l0)
	mBase = m.M
	v11027 = m.ExcPending
	if v11027 != 0 {
		goto L1
	} else {
		goto L1531
	}
L1531:
	;
	v11028 = F_list_concat(m, v11019, v11026)
	mBase = m.M
	v11029 = m.ExcPending
	if v11029 != 0 {
		goto L1
	} else {
		goto L1532
	}
L1532:
	;
	v11030 = v11028
	goto L1529
L1533:
	;
	v11033 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v11034 = F_make_partition_pruneinfo(m, l0, v10566, v11033, v11030)
	mBase = m.M
	v11035 = m.ExcPending
	if v11035 != 0 {
		goto L1
	} else {
		goto L1534
	}
L1534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10619)+88)) = v11034
	goto L1524
L1535:
	;
	v11063 = *(*int32)(unsafe.Add(mBase, uint32(v10619)+44))
	v11064 = F_list_copy_head(m, v11063, v10565)
	mBase = m.M
	v11065 = m.ExcPending
	if v11065 != 0 {
		goto L1
	} else {
		goto L1536
	}
L1536:
	;
	v11066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10619)+37)))
	v11067 = F_inject_projection_plan(m, v10619, v11064, v11066)
	mBase = m.M
	v11068 = m.ExcPending
	if v11068 != 0 {
		goto L1
	} else {
		goto L1537
	}
L1537:
	;
	v12573 = v11067
	v12580 = v47
	goto L3
L1538:
	;
	v11196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v11197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v11198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v11199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v11200 = *(*int32)(unsafe.Add(mBase, uint32(v11199)+8))
	v11201 = F_reparameterize_path_by_child(m, l0, v11198, v11200)
	mBase = m.M
	v11202 = m.ExcPending
	if v11202 != 0 {
		goto L1
	} else {
		goto L1553
	}
L1539:
	;
	v11075 = *(*int32)(unsafe.Add(mBase, uint32(v11071)+4))
	if v11075 <= int32(0) {
		v11154 = v11069
		goto L1538
	} else {
		goto L1540
	}
L1540:
	;
	v11078 = *(*int32)(unsafe.Add(mBase, uint32(v11070)+8))
	v11081 = v11069
	v11082 = int32(1)
	v11084 = v4
	goto L1541
L1541:
	;
	v11123 = *(*int32)(unsafe.Add(mBase, uint32(v11071)+12))
	v11127 = *(*int32)(unsafe.Add(mBase, uint32(v11123+v11084<<(uint(int32(2))%32))))
	v11128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11128 != 0 {
		goto L1543
	} else {
		goto L1544
	}
L1542:
	;
	v11154 = v11146
	goto L1538
L1543:
	;
	v11129 = F_replace_nestloop_params_mutator(m, v11127, l0)
	mBase = m.M
	v11130 = m.ExcPending
	if v11130 != 0 {
		goto L1
	} else {
		goto L1546
	}
L1544:
	;
	v11131 = v11127
	goto L1545
L1545:
	;
	v11133 = int32(0)
	v11135 = F_makeTargetEntry(m, v11131, base.I32_extend16_s(v11082), v11133, v11133)
	mBase = m.M
	v11136 = m.ExcPending
	if v11136 != 0 {
		goto L1
	} else {
		goto L1547
	}
L1546:
	;
	v11131 = v11129
	goto L1545
L1547:
	;
	if v11078 != 0 {
		goto L1548
	} else {
		goto L1549
	}
L1548:
	;
	v11142 = *(*int32)(unsafe.Add(mBase, uint32(v11078+v11082<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v11135)+16)) = v11142
	goto L1550
L1549:
	;
	goto L1550
L1550:
	;
	v11146 = F_lappend(m, v11081, v11135)
	mBase = m.M
	v11147 = m.ExcPending
	if v11147 != 0 {
		goto L1
	} else {
		goto L1551
	}
L1551:
	;
	v11149 = v11084 + int32(1)
	v11150 = *(*int32)(unsafe.Add(mBase, uint32(v11071)+4))
	if v11149 < v11150 {
		v11081 = v11146
		v11082 = v11082 + int32(1)
		v11084 = v11149
		goto L1541
	} else {
		goto L1552
	}
L1552:
	;
	goto L1542
L1553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v11201
	v11204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v11206 = F_create_plan_recurse(m, l0, v11204, int32(0))
	mBase = m.M
	v11207 = m.ExcPending
	if v11207 != 0 {
		goto L1
	} else {
		goto L1554
	}
L1554:
	;
	v11208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	v11209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v11210 = *(*int32)(unsafe.Add(mBase, uint32(v11209)+8))
	v11211 = *(*int32)(unsafe.Add(mBase, uint32(v11210)+8))
	v11212 = F_bms_union(m, v11208, v11211)
	mBase = m.M
	v11213 = m.ExcPending
	if v11213 != 0 {
		goto L1
	} else {
		goto L1555
	}
L1555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v11212
	v11215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v11217 = F_create_plan_recurse(m, l0, v11215, int32(0))
	mBase = m.M
	v11218 = m.ExcPending
	if v11218 != 0 {
		goto L1
	} else {
		goto L1556
	}
L1556:
	;
	v11219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	F_bms_free(m, v11219)
	mBase = m.M
	v11221 = m.ExcPending
	if v11221 != 0 {
		goto L1
	} else {
		goto L1557
	}
L1557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v11197
	v11223 = F_order_qual_clauses(m, l0, v11196)
	mBase = m.M
	v11224 = m.ExcPending
	if v11224 != 0 {
		goto L1
	} else {
		goto L1558
	}
L1558:
	;
	v11226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v11226)%32)&int32(174) != 0 {
		goto L1560
	} else {
		goto L1561
	}
L1559:
	;
	v11245 = int32(0)
	v11246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11246 == v11245 {
		v11261 = v11245
		goto L1565
	} else {
		goto L1566
	}
L1560:
	;
	v11230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11231 = *(*int32)(unsafe.Add(mBase, uint32(v11230)+8))
	F_extract_actual_join_clauses(m, v11223, v11231, v47+int32(56), v47+int32(140))
	mBase = m.M
	v11237 = m.ExcPending
	if v11237 != 0 {
		goto L1
	} else {
		goto L1563
	}
L1561:
	;
	goto L1562
L1562:
	;
	v11239 = F_extract_actual_clauses(m, v11223, int32(0))
	mBase = m.M
	v11240 = m.ExcPending
	if v11240 != 0 {
		goto L1
	} else {
		goto L1564
	}
L1563:
	;
	goto L1559
L1564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v11239
	goto L1559
L1565:
	;
	if v11261 != 0 {
		goto L1572
	} else {
		goto L1573
	}
L1566:
	;
	v11249 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v11250 = F_replace_nestloop_params_mutator(m, v11249, l0)
	mBase = m.M
	v11251 = m.ExcPending
	if v11251 != 0 {
		goto L1
	} else {
		goto L1567
	}
L1567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v11250
	v11253 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v11254 = F_replace_nestloop_params_mutator(m, v11253, l0)
	mBase = m.M
	v11255 = m.ExcPending
	if v11255 != 0 {
		goto L1
	} else {
		goto L1568
	}
L1568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = v11254
	v11257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11257 == int32(0) {
		v11261 = v11245
		goto L1565
	} else {
		goto L1569
	}
L1569:
	;
	v11260 = *(*int32)(unsafe.Add(mBase, uint32(v11257)+4))
	v11261 = v11260
	goto L1565
L1570:
	;
	v12074 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v12075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v12076 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v12077 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v12079 = F_palloc0(m, int32(96))
	mBase = m.M
	v12080 = m.ExcPending
	if v12080 != 0 {
		goto L1
	} else {
		goto L1690
	}
L1571:
	;
	if v11879 == int32(0) {
		v12041 = v11206
		goto L1570
	} else {
		goto L1664
	}
L1572:
	;
	v11263 = F_bms_union(m, v11211, v11261)
	mBase = m.M
	v11264 = m.ExcPending
	if v11264 != 0 {
		goto L1
	} else {
		goto L1575
	}
L1573:
	;
	v11265 = v11211
	goto L1574
L1574:
	;
	v11266 = int32(0)
	v11267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v11267 == v11266 {
		v11879 = v11266
		goto L1571
	} else {
		goto L1576
	}
L1575:
	;
	v11265 = v11263
	goto L1574
L1576:
	;
	v11274 = int32(0)
	v11277 = v11267
	v11280 = v4
	goto L1577
L1577:
	;
	v11315 = *(*int32)(unsafe.Add(mBase, uint32(v11277)+4))
	if v11274 < v11315 {
		goto L1579
	} else {
		goto L1580
	}
L1578:
	;
	v11879 = v11800
	goto L1571
L1579:
	;
	v11317 = *(*int32)(unsafe.Add(mBase, uint32(v11277)+12))
	v11321 = *(*int32)(unsafe.Add(mBase, uint32(v11317+v11274<<(uint(int32(2))%32))))
	v11322 = *(*int32)(unsafe.Add(mBase, uint32(v11321)+8))
	v11323 = *(*int32)(unsafe.Add(mBase, uint32(v11322)))
	if v11323 == int32(6) {
		goto L1585
	} else {
		goto L1586
	}
L1580:
	;
	v11800 = v11280
	goto L1581
L1581:
	;
	goto L1578
L1582:
	;
	if v11751 != 0 {
		v11274 = v11748 + int32(1)
		v11277 = v11751
		v11280 = v11754
		goto L1577
	} else {
		goto L1663
	}
L1583:
	;
	v11743 = F_lappend(m, v11280, v11321)
	mBase = m.M
	v11744 = m.ExcPending
	if v11744 != 0 {
		goto L1
	} else {
		goto L1662
	}
L1584:
	;
	v11683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11684 = *(*int32)(unsafe.Add(mBase, uint32(v11329)+4))
	v11688 = *(*int32)(unsafe.Add(mBase, uint32(v11683+v11684<<(uint(int32(2))%32))))
	v11689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	v11690 = F_list_delete_nth_cell(m, v11689, v11274)
	mBase = m.M
	v11691 = m.ExcPending
	if v11691 != 0 {
		goto L1
	} else {
		goto L1660
	}
L1585:
	;
	v11326 = *(*int32)(unsafe.Add(mBase, uint32(v11322)+4))
	v11327 = F_bms_is_member(m, v11326, v11211)
	mBase = m.M
	v11328 = m.ExcPending
	if v11328 != 0 {
		goto L1
	} else {
		goto L1588
	}
L1586:
	;
	v11331 = v11322
	v11332 = v11323
	goto L1587
L1587:
	;
	if v11332 != int32(319) {
		v11748 = v11274
		v11751 = v11277
		v11754 = v11280
		goto L1582
	} else {
		goto L1590
	}
L1588:
	;
	v11329 = *(*int32)(unsafe.Add(mBase, uint32(v11321)+8))
	if v11327 != 0 {
		goto L1584
	} else {
		goto L1589
	}
L1589:
	;
	v11330 = *(*int32)(unsafe.Add(mBase, uint32(v11329)))
	v11331 = v11329
	v11332 = v11330
	goto L1587
L1590:
	;
	v11335 = F_find_placeholder_info(m, l0, v11331)
	mBase = m.M
	v11336 = m.ExcPending
	if v11336 != 0 {
		goto L1
	} else {
		goto L1591
	}
L1591:
	;
	v11337 = *(*int32)(unsafe.Add(mBase, uint32(v11335)+12))
	v11338 = int32(0)
	if v11337 == v11338 {
		goto L1593
	} else {
		goto L1594
	}
L1592:
	;
	if v11391 == int32(0) {
		v11748 = v11274
		v11751 = v11277
		v11754 = v11280
		goto L1582
	} else {
		goto L1606
	}
L1593:
	;
	v11391 = int32(1)
	goto L1592
L1594:
	;
	goto L1595
L1595:
	;
	if v11265 == int32(0) {
		v11384 = v11338
		goto L1596
	} else {
		goto L1597
	}
L1596:
	;
	v11391 = v11384
	goto L1592
L1597:
	;
	v11347 = *(*int32)(unsafe.Add(mBase, uint32(v11337)+4))
	v11348 = *(*int32)(unsafe.Add(mBase, uint32(v11265)+4))
	if v11348 < v11347 {
		v11384 = v11338
		goto L1596
	} else {
		goto L1598
	}
L1598:
	;
	v11350 = int32(1)
	if v11347 <= v11350 {
		goto L1599
	} else {
		goto L1600
	}
L1599:
	;
	v11353 = v11350
	goto L1601
L1600:
	;
	v11353 = v11347
	goto L1601
L1601:
	;
	v11354 = int32(8)
	v11359 = int32(0)
	goto L1602
L1602:
	;
	v11366 = v11359 << (uint(int32(2)) % 32)
	v11368 = *(*int32)(unsafe.Add(mBase, uint32(v11337+v11354+v11366)))
	v11370 = *(*int32)(unsafe.Add(mBase, uint32(v11265+v11354+v11366)))
	v11373 = v11368 & (v11370 ^ int32(-1))
	v11375 = base.B2i32(v11373 == int32(0))
	if v11373 != 0 {
		v11384 = v11375
		goto L1596
	} else {
		goto L1604
	}
L1603:
	;
	v11384 = v11375
	goto L1596
L1604:
	;
	v11377 = v11359 + int32(1)
	if v11377 != v11353 {
		v11359 = v11377
		goto L1602
	} else {
		goto L1605
	}
L1605:
	;
	goto L1603
L1606:
	;
	v11394 = int32(0)
	if base.B2i32(v11337 == v11394)|base.B2i32(v11211 == v11394) != 0 {
		v11439 = v11394
		goto L1608
	} else {
		goto L1609
	}
L1607:
	;
	if v11439 == int32(0) {
		v11748 = v11274
		v11751 = v11277
		v11754 = v11280
		goto L1582
	} else {
		goto L1620
	}
L1608:
	;
	goto L1607
L1609:
	;
	v11404 = *(*int32)(unsafe.Add(mBase, uint32(v11337)+4))
	v11405 = *(*int32)(unsafe.Add(mBase, uint32(v11211)+4))
	if v11404 < v11405 {
		goto L1610
	} else {
		goto L1611
	}
L1610:
	;
	v11407 = v11404
	goto L1612
L1611:
	;
	v11407 = v11405
	goto L1612
L1612:
	;
	if v11407 <= int32(1) {
		goto L1613
	} else {
		goto L1614
	}
L1613:
	;
	v11410 = int32(1)
	goto L1615
L1614:
	;
	v11410 = v11407
	goto L1615
L1615:
	;
	v11411 = int32(8)
	v11416 = int32(0)
	goto L1616
L1616:
	;
	v11423 = v11416 << (uint(int32(2)) % 32)
	v11425 = *(*int32)(unsafe.Add(mBase, uint32(v11211+v11411+v11423)))
	v11427 = *(*int32)(unsafe.Add(mBase, uint32(v11337+v11411+v11423)))
	v11428 = v11425 & v11427
	v11430 = base.B2i32(v11428 != int32(0))
	if v11428 != 0 {
		v11439 = v11430
		goto L1608
	} else {
		goto L1618
	}
L1617:
	;
	v11439 = v11430
	goto L1608
L1618:
	;
	v11432 = v11416 + int32(1)
	if v11432 != v11410 {
		v11416 = v11432
		goto L1616
	} else {
		goto L1619
	}
L1619:
	;
	goto L1617
L1620:
	;
	v11442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	v11443 = F_list_delete_nth_cell(m, v11442, v11274)
	mBase = m.M
	v11444 = m.ExcPending
	if v11444 != 0 {
		goto L1
	} else {
		goto L1621
	}
L1621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v11443
	v11446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11446)+39)))
	if v11447 == int32(1) {
		goto L1622
	} else {
		goto L1623
	}
L1622:
	;
	v11450 = *(*int32)(unsafe.Add(mBase, uint32(v11335)+8))
	v11451 = F_copyObjectImpl(m, v11450)
	mBase = m.M
	v11452 = m.ExcPending
	if v11452 != 0 {
		goto L1
	} else {
		goto L1625
	}
L1623:
	;
	v11454 = v11331
	goto L1624
L1624:
	;
	v11455 = int32(0)
	v11456 = *(*int32)(unsafe.Add(mBase, uint32(v11335)+12))
	if v11456 == v11455 {
		goto L1628
	} else {
		goto L1629
	}
L1625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11321)+8)) = v11451
	v11454 = v11451
	goto L1624
L1626:
	;
	if int32(0) < v11513 {
		goto L1637
	} else {
		goto L1638
	}
L1627:
	;
	v11513 = base.I32_ctz(v11499) | v11500<<(uint(int32(5))%32)
	goto L1626
L1628:
	;
	v11513 = int32(-2)
	goto L1626
L1629:
	;
	v11466 = base.I32_div_s(int32(0), int32(32))
	v11467 = *(*int32)(unsafe.Add(mBase, uint32(v11456)+4))
	if v11467 <= v11466 {
		goto L1628
	} else {
		goto L1630
	}
L1630:
	;
	v11470 = v11456 + int32(8)
	v11474 = *(*int32)(unsafe.Add(mBase, uint32(v11470+v11466<<(uint(int32(2))%32))))
	v11477 = v11474 & int32(-1)
	if v11477 != 0 {
		v11499 = v11477
		v11500 = v11466
		goto L1627
	} else {
		goto L1631
	}
L1631:
	;
	v11479 = v11466 + int32(1)
	if v11479 == v11467 {
		goto L1628
	} else {
		goto L1632
	}
L1632:
	;
	v11482 = v11479
	goto L1633
L1633:
	;
	v11489 = *(*int32)(unsafe.Add(mBase, uint32(v11470+v11482<<(uint(int32(2))%32))))
	if v11489 != 0 {
		v11499 = v11489
		v11500 = v11482
		goto L1627
	} else {
		goto L1635
	}
L1634:
	;
	goto L1628
L1635:
	;
	v11491 = v11482 + int32(1)
	if v11491 != v11467 {
		v11482 = v11491
		goto L1633
	} else {
		goto L1636
	}
L1636:
	;
	goto L1634
L1637:
	;
	v11520 = v11513
	v11521 = v11455
	goto L1640
L1638:
	;
	v11638 = v11455
	goto L1639
L1639:
	;
	v11677 = *(*int32)(unsafe.Add(mBase, uint32(v11335)+12))
	v11678 = F_bms_del_members(m, v11638, v11677)
	mBase = m.M
	v11679 = m.ExcPending
	if v11679 != 0 {
		goto L1
	} else {
		goto L1658
	}
L1640:
	;
	v11560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v11520 == v11560 {
		v11572 = v11521
		goto L1642
	} else {
		goto L1643
	}
L1641:
	;
	v11638 = v11572
	goto L1639
L1642:
	;
	v11574 = *(*int32)(unsafe.Add(mBase, uint32(v11335)+12))
	if v11574 == int32(0) {
		goto L1648
	} else {
		goto L1649
	}
L1643:
	;
	v11562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11566 = *(*int32)(unsafe.Add(mBase, uint32(v11562+v11520<<(uint(int32(2))%32))))
	if v11566 == int32(0) {
		v11572 = v11521
		goto L1642
	} else {
		goto L1644
	}
L1644:
	;
	v11569 = *(*int32)(unsafe.Add(mBase, uint32(v11566)+96))
	v11570 = F_bms_add_members(m, v11521, v11569)
	mBase = m.M
	v11571 = m.ExcPending
	if v11571 != 0 {
		goto L1
	} else {
		goto L1645
	}
L1645:
	;
	v11572 = v11570
	goto L1642
L1646:
	;
	if int32(0) < v11630 {
		v11520 = v11630
		v11521 = v11572
		goto L1640
	} else {
		goto L1657
	}
L1647:
	;
	v11630 = base.I32_ctz(v11616) | v11617<<(uint(int32(5))%32)
	goto L1646
L1648:
	;
	v11630 = int32(-2)
	goto L1646
L1649:
	;
	v11581 = v11520 + int32(1)
	v11583 = base.I32_div_s(v11581, int32(32))
	v11584 = *(*int32)(unsafe.Add(mBase, uint32(v11574)+4))
	if v11584 <= v11583 {
		goto L1648
	} else {
		goto L1650
	}
L1650:
	;
	v11587 = v11574 + int32(8)
	v11591 = *(*int32)(unsafe.Add(mBase, uint32(v11587+v11583<<(uint(int32(2))%32))))
	v11594 = v11591 & (int32(-1) << (uint(v11581) % 32))
	if v11594 != 0 {
		v11616 = v11594
		v11617 = v11583
		goto L1647
	} else {
		goto L1651
	}
L1651:
	;
	v11596 = v11583 + int32(1)
	if v11596 == v11584 {
		goto L1648
	} else {
		goto L1652
	}
L1652:
	;
	v11599 = v11596
	goto L1653
L1653:
	;
	v11606 = *(*int32)(unsafe.Add(mBase, uint32(v11587+v11599<<(uint(int32(2))%32))))
	if v11606 != 0 {
		v11616 = v11606
		v11617 = v11599
		goto L1647
	} else {
		goto L1655
	}
L1654:
	;
	goto L1648
L1655:
	;
	v11608 = v11599 + int32(1)
	if v11608 != v11584 {
		v11599 = v11608
		goto L1653
	} else {
		goto L1656
	}
L1656:
	;
	goto L1654
L1657:
	;
	goto L1641
L1658:
	;
	v11680 = F_bms_intersect(m, v11678, v11211)
	mBase = m.M
	v11681 = m.ExcPending
	if v11681 != 0 {
		goto L1
	} else {
		goto L1659
	}
L1659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11454)+12)) = v11680
	v11703 = v11443
	goto L1583
L1660:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v11690
	v11693 = *(*int32)(unsafe.Add(mBase, uint32(v11688)+96))
	v11694 = F_bms_intersect(m, v11693, v11211)
	mBase = m.M
	v11695 = m.ExcPending
	if v11695 != 0 {
		goto L1
	} else {
		goto L1661
	}
L1661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11329)+24)) = v11694
	v11703 = v11690
	goto L1583
L1662:
	;
	v11748 = v11274 - int32(1)
	v11751 = v11703
	v11754 = v11743
	goto L1582
L1663:
	;
	v11800 = v11754
	goto L1581
L1664:
	;
	v11882 = *(*int32)(unsafe.Add(mBase, uint32(v11206)+44))
	v11883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11206)+37)))
	v11884 = *(*int32)(unsafe.Add(mBase, uint32(v11879)+4))
	if int32(0) < v11884 {
		goto L1665
	} else {
		goto L1666
	}
L1665:
	;
	v11891 = int32(0)
	v11895 = v11883
	v11896 = v11882
	goto L1668
L1666:
	;
	v11987 = v11883
	v11988 = v11882
	goto L1667
L1667:
	;
	v12024 = *(*int32)(unsafe.Add(mBase, uint32(v11206)+44))
	if v11988 == v12024 {
		v12041 = v11206
		goto L1570
	} else {
		goto L1688
	}
L1668:
	;
	v11932 = *(*int32)(unsafe.Add(mBase, uint32(v11879)+12))
	v11936 = *(*int32)(unsafe.Add(mBase, uint32(v11932+v11891<<(uint(int32(2))%32))))
	v11937 = *(*int32)(unsafe.Add(mBase, uint32(v11936)+8))
	v11938 = *(*int32)(unsafe.Add(mBase, uint32(v11937)))
	if v11938 == int32(6) {
		v11974 = v11895
		v11975 = v11896
		goto L1670
	} else {
		goto L1671
	}
L1669:
	;
	v11987 = v11974
	v11988 = v11975
	goto L1667
L1670:
	;
	v11977 = v11891 + int32(1)
	v11978 = *(*int32)(unsafe.Add(mBase, uint32(v11879)+4))
	if v11977 < v11978 {
		v11891 = v11977
		v11895 = v11974
		v11896 = v11975
		goto L1668
	} else {
		goto L1687
	}
L1671:
	;
	v11941 = F_tlist_member(m, v11937, v11896)
	mBase = m.M
	v11942 = m.ExcPending
	if v11942 != 0 {
		goto L1
	} else {
		goto L1672
	}
L1672:
	;
	if v11941 != 0 {
		v11974 = v11895
		v11975 = v11896
		goto L1670
	} else {
		goto L1673
	}
L1673:
	;
	v11943 = *(*int32)(unsafe.Add(mBase, uint32(v11937)+4))
	v11944 = F_replace_nestloop_params_mutator(m, v11943, l0)
	mBase = m.M
	v11945 = m.ExcPending
	if v11945 != 0 {
		goto L1
	} else {
		goto L1674
	}
L1674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11937)+4)) = v11944
	v11947 = *(*int32)(unsafe.Add(mBase, uint32(v11206)+44))
	if v11947 == v11896 {
		goto L1675
	} else {
		goto L1676
	}
L1675:
	;
	v11949 = F_list_copy(m, v11896)
	mBase = m.M
	v11950 = m.ExcPending
	if v11950 != 0 {
		goto L1
	} else {
		goto L1678
	}
L1676:
	;
	v11951 = v11896
	goto L1677
L1677:
	;
	v11952 = F_copyObjectImpl(m, v11937)
	mBase = m.M
	v11953 = m.ExcPending
	if v11953 != 0 {
		goto L1
	} else {
		goto L1679
	}
L1678:
	;
	v11951 = v11949
	goto L1677
L1679:
	;
	if v11951 != 0 {
		goto L1680
	} else {
		goto L1681
	}
L1680:
	;
	v11957 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11951)+4)))
	v11961 = v11957 + int32(1)
	goto L1682
L1681:
	;
	v11961 = int32(1)
	goto L1682
L1682:
	;
	v11965 = F_makeTargetEntry(m, v11952, base.I32_extend16_s(v11961), int32(0), int32(1))
	mBase = m.M
	v11966 = m.ExcPending
	if v11966 != 0 {
		goto L1
	} else {
		goto L1683
	}
L1683:
	;
	v11967 = F_lappend(m, v11951, v11965)
	mBase = m.M
	v11968 = m.ExcPending
	if v11968 != 0 {
		goto L1
	} else {
		goto L1684
	}
L1684:
	;
	if v11895&int32(1) == int32(0) {
		v11974 = int32(0)
		v11975 = v11967
		goto L1670
	} else {
		goto L1685
	}
L1685:
	;
	v11971 = F_is_parallel_safe(m, l0, v11937)
	mBase = m.M
	v11972 = m.ExcPending
	if v11972 != 0 {
		goto L1
	} else {
		goto L1686
	}
L1686:
	;
	v11974 = v11971
	v11975 = v11967
	goto L1670
L1687:
	;
	goto L1669
L1688:
	;
	v12028 = F_change_plan_targetlist(m, v11206, v11988, v11987&int32(1))
	mBase = m.M
	v12029 = m.ExcPending
	if v12029 != 0 {
		goto L1
	} else {
		goto L1689
	}
L1689:
	;
	v12041 = v12028
	goto L1570
L1690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12079)+88)) = v11879
	*(*int32)(unsafe.Add(mBase, uint32(v12079)+80)) = v12077
	*(*uint8)(unsafe.Add(mBase, uint32(v12079)+76)) = uint8(v12075)
	*(*int32)(unsafe.Add(mBase, uint32(v12079)+72)) = v12074
	*(*int32)(unsafe.Add(mBase, uint32(v12079)+56)) = v11217
	*(*int32)(unsafe.Add(mBase, uint32(v12079)+52)) = v12041
	*(*int32)(unsafe.Add(mBase, uint32(v12079)+48)) = v12076
	*(*int32)(unsafe.Add(mBase, uint32(v12079)+44)) = v11154
	*(*int32)(unsafe.Add(mBase, uint32(v12079))) = int32(356)
	v12501 = v12079
	goto L4
L1691:
	;
	v12218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v12219 = int32(2)
	v12221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v12219 <= v12221 {
		goto L1706
	} else {
		goto L1707
	}
L1692:
	;
	v12097 = *(*int32)(unsafe.Add(mBase, uint32(v12093)+4))
	if v12097 <= int32(0) {
		v12176 = v12091
		goto L1691
	} else {
		goto L1693
	}
L1693:
	;
	v12100 = *(*int32)(unsafe.Add(mBase, uint32(v12092)+8))
	v12103 = v12091
	v12104 = int32(1)
	v12106 = v4
	goto L1694
L1694:
	;
	v12145 = *(*int32)(unsafe.Add(mBase, uint32(v12093)+12))
	v12149 = *(*int32)(unsafe.Add(mBase, uint32(v12145+v12106<<(uint(int32(2))%32))))
	v12150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v12150 != 0 {
		goto L1696
	} else {
		goto L1697
	}
L1695:
	;
	v12176 = v12168
	goto L1691
L1696:
	;
	v12151 = F_replace_nestloop_params_mutator(m, v12149, l0)
	mBase = m.M
	v12152 = m.ExcPending
	if v12152 != 0 {
		goto L1
	} else {
		goto L1699
	}
L1697:
	;
	v12153 = v12149
	goto L1698
L1698:
	;
	v12155 = int32(0)
	v12157 = F_makeTargetEntry(m, v12153, base.I32_extend16_s(v12104), v12155, v12155)
	mBase = m.M
	v12158 = m.ExcPending
	if v12158 != 0 {
		goto L1
	} else {
		goto L1700
	}
L1699:
	;
	v12153 = v12151
	goto L1698
L1700:
	;
	if v12100 != 0 {
		goto L1701
	} else {
		goto L1702
	}
L1701:
	;
	v12164 = *(*int32)(unsafe.Add(mBase, uint32(v12100+v12104<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v12157)+16)) = v12164
	goto L1703
L1702:
	;
	goto L1703
L1703:
	;
	v12168 = F_lappend(m, v12103, v12157)
	mBase = m.M
	v12169 = m.ExcPending
	if v12169 != 0 {
		goto L1
	} else {
		goto L1704
	}
L1704:
	;
	v12171 = v12106 + int32(1)
	v12172 = *(*int32)(unsafe.Add(mBase, uint32(v12093)+4))
	if v12171 < v12172 {
		v12103 = v12168
		v12104 = v12104 + int32(1)
		v12106 = v12171
		goto L1694
	} else {
		goto L1705
	}
L1705:
	;
	goto L1695
L1706:
	;
	v12224 = v12219
	goto L1708
L1707:
	;
	v12224 = int32(0)
	goto L1708
L1708:
	;
	v12225 = F_create_plan_recurse(m, l0, v12218, v12224)
	mBase = m.M
	v12226 = m.ExcPending
	if v12226 != 0 {
		goto L1
	} else {
		goto L1709
	}
L1709:
	;
	v12227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v12229 = F_create_plan_recurse(m, l0, v12227, int32(2))
	mBase = m.M
	v12230 = m.ExcPending
	if v12230 != 0 {
		goto L1
	} else {
		goto L1710
	}
L1710:
	;
	v12231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v12232 = F_order_qual_clauses(m, l0, v12231)
	mBase = m.M
	v12233 = m.ExcPending
	if v12233 != 0 {
		goto L1
	} else {
		goto L1711
	}
L1711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12232
	v12236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if int32(1)<<(uint(v12236)%32)&int32(174) != 0 {
		goto L1713
	} else {
		goto L1714
	}
L1712:
	;
	v12255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v12256 = F_get_actual_clauses(m, v12255)
	mBase = m.M
	v12257 = m.ExcPending
	if v12257 != 0 {
		goto L1
	} else {
		goto L1718
	}
L1713:
	;
	v12240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12241 = *(*int32)(unsafe.Add(mBase, uint32(v12240)+8))
	F_extract_actual_join_clauses(m, v12232, v12241, v47+int32(56), v47+int32(140))
	mBase = m.M
	v12247 = m.ExcPending
	if v12247 != 0 {
		goto L1
	} else {
		goto L1716
	}
L1714:
	;
	goto L1715
L1715:
	;
	v12249 = F_extract_actual_clauses(m, v12232, int32(0))
	mBase = m.M
	v12250 = m.ExcPending
	if v12250 != 0 {
		goto L1
	} else {
		goto L1717
	}
L1716:
	;
	goto L1712
L1717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12249
	goto L1712
L1718:
	;
	v12258 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v12259 = F_list_difference(m, v12258, v12256)
	mBase = m.M
	v12260 = m.ExcPending
	if v12260 != 0 {
		goto L1
	} else {
		goto L1719
	}
L1719:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12259
	v12262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v12262 != 0 {
		goto L1720
	} else {
		goto L1721
	}
L1720:
	;
	v12263 = F_replace_nestloop_params_mutator(m, v12259, l0)
	mBase = m.M
	v12264 = m.ExcPending
	if v12264 != 0 {
		goto L1
	} else {
		goto L1723
	}
L1721:
	;
	goto L1722
L1722:
	;
	v12270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v12271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v12272 = *(*int32)(unsafe.Add(mBase, uint32(v12271)+8))
	v12273 = *(*int32)(unsafe.Add(mBase, uint32(v12272)+8))
	v12274 = F_get_switched_clauses(m, v12270, v12273)
	mBase = m.M
	v12275 = m.ExcPending
	if v12275 != 0 {
		goto L1
	} else {
		goto L1728
	}
L1723:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+56)) = v12263
	v12266 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v12267 = F_replace_nestloop_params_mutator(m, v12266, l0)
	mBase = m.M
	v12268 = m.ExcPending
	if v12268 != 0 {
		goto L1
	} else {
		goto L1724
	}
L1724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+140)) = v12267
	goto L1722
L1725:
	;
	v12439 = F_palloc0(m, int32(96))
	mBase = m.M
	v12440 = m.ExcPending
	if v12440 != 0 {
		goto L1
	} else {
		goto L1748
	}
L1726:
	;
	v12329 = int32(0)
	v12332 = v4
	v12333 = v12279
	v12336 = v4
	v12337 = v4
	goto L1741
L1727:
	;
	v12401 = v4
	v12402 = v12320
	v12403 = v12321
	v12405 = v4
	v12406 = v4
	v12407 = v12322
	v12408 = v12323
	goto L1725
L1728:
	;
	if v12274 == int32(0) {
		goto L1729
	} else {
		goto L1730
	}
L1729:
	;
	v12320 = int32(0)
	v12321 = v4
	v12322 = v4
	v12323 = v4
	goto L1727
L1730:
	;
	goto L1731
L1731:
	;
	v12279 = int32(0)
	v12280 = *(*int32)(unsafe.Add(mBase, uint32(v12274)+4))
	if v12280 != int32(1) {
		goto L1733
	} else {
		goto L1734
	}
L1732:
	;
	v12315 = *(*int32)(unsafe.Add(mBase, uint32(v12274)+4))
	if int32(0) < v12315 {
		goto L1726
	} else {
		goto L1740
	}
L1733:
	;
	v12312 = v4
	v12313 = v4
	v12314 = int32(0)
	goto L1732
L1734:
	;
	v12283 = *(*int32)(unsafe.Add(mBase, uint32(v12274)+12))
	v12284 = *(*int32)(unsafe.Add(mBase, uint32(v12283)))
	v12285 = *(*int32)(unsafe.Add(mBase, uint32(v12284)+28))
	v12286 = *(*int32)(unsafe.Add(mBase, uint32(v12285)+12))
	v12287 = *(*int32)(unsafe.Add(mBase, uint32(v12286)))
	v12288 = *(*int32)(unsafe.Add(mBase, uint32(v12287)))
	if v12288 == int32(27) {
		goto L1735
	} else {
		goto L1736
	}
L1735:
	;
	v12291 = *(*int32)(unsafe.Add(mBase, uint32(v12287)+4))
	v12292 = *(*int32)(unsafe.Add(mBase, uint32(v12291)))
	v12293 = v12291
	v12294 = v12292
	goto L1737
L1736:
	;
	v12293 = v12287
	v12294 = v12288
	goto L1737
L1737:
	;
	if v12294 != int32(6) {
		goto L1733
	} else {
		goto L1738
	}
L1738:
	;
	v12297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12298 = *(*int32)(unsafe.Add(mBase, uint32(v12293)+4))
	v12302 = *(*int32)(unsafe.Add(mBase, uint32(v12297+v12298<<(uint(int32(2))%32))))
	v12303 = *(*int32)(unsafe.Add(mBase, uint32(v12302)+12))
	if v12303 != 0 {
		goto L1733
	} else {
		goto L1739
	}
L1739:
	;
	v12304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12302)+20)))
	v12305 = *(*int32)(unsafe.Add(mBase, uint32(v12302)+16))
	v12306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12293)+8)))
	v12312 = v12304
	v12313 = v12305
	v12314 = v12306
	goto L1732
L1740:
	;
	v12320 = v12279
	v12321 = v12312
	v12322 = v12314
	v12323 = v12313
	goto L1727
L1741:
	;
	v12369 = *(*int32)(unsafe.Add(mBase, uint32(v12274)+12))
	v12373 = *(*int32)(unsafe.Add(mBase, uint32(v12369+v12329<<(uint(int32(2))%32))))
	v12374 = *(*int32)(unsafe.Add(mBase, uint32(v12373)+4))
	v12375 = F_lappend_oid(m, v12337, v12374)
	mBase = m.M
	v12376 = m.ExcPending
	if v12376 != 0 {
		goto L1
	} else {
		goto L1743
	}
L1742:
	;
	v12401 = v12378
	v12402 = v12383
	v12403 = v12312
	v12405 = v12388
	v12406 = v12375
	v12407 = v12314
	v12408 = v12313
	goto L1725
L1743:
	;
	v12377 = *(*int32)(unsafe.Add(mBase, uint32(v12373)+24))
	v12378 = F_lappend_oid(m, v12332, v12377)
	mBase = m.M
	v12379 = m.ExcPending
	if v12379 != 0 {
		goto L1
	} else {
		goto L1744
	}
L1744:
	;
	v12380 = *(*int32)(unsafe.Add(mBase, uint32(v12373)+28))
	v12381 = *(*int32)(unsafe.Add(mBase, uint32(v12380)+12))
	v12382 = *(*int32)(unsafe.Add(mBase, uint32(v12381)))
	v12383 = F_lappend(m, v12333, v12382)
	mBase = m.M
	v12384 = m.ExcPending
	if v12384 != 0 {
		goto L1
	} else {
		goto L1745
	}
L1745:
	;
	v12385 = *(*int32)(unsafe.Add(mBase, uint32(v12373)+28))
	v12386 = *(*int32)(unsafe.Add(mBase, uint32(v12385)+12))
	v12387 = *(*int32)(unsafe.Add(mBase, uint32(v12386)+4))
	v12388 = F_lappend(m, v12336, v12387)
	mBase = m.M
	v12389 = m.ExcPending
	if v12389 != 0 {
		goto L1
	} else {
		goto L1746
	}
L1746:
	;
	v12391 = v12329 + int32(1)
	v12392 = *(*int32)(unsafe.Add(mBase, uint32(v12274)+4))
	if v12391 < v12392 {
		v12329 = v12391
		v12332 = v12378
		v12333 = v12383
		v12336 = v12388
		v12337 = v12375
		goto L1741
	} else {
		goto L1747
	}
L1747:
	;
	goto L1742
L1748:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12439))) = int32(370)
	v12443 = *(*int32)(unsafe.Add(mBase, uint32(v12229)+44))
	v12444 = int32(1)
	v12445 = v12403 & v12444
	*(*uint8)(unsafe.Add(mBase, uint32(v12439)+82)) = uint8(v12445)
	*(*uint16)(unsafe.Add(mBase, uint32(v12439)+80)) = uint16(v12407)
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+76)) = v12408
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+72)) = v12405
	v12450 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+56)) = v12450
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+52)) = v12229
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+48)) = v12450
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+44)) = v12443
	v12456 = *(*int32)(unsafe.Add(mBase, uint32(v12229)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+4)) = v12456
	v12458 = *(*float64)(unsafe.Add(mBase, uint32(v12229)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12439)+8)) = v12458
	v12460 = *(*float64)(unsafe.Add(mBase, uint32(v12229)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v12439)+16)) = v12460
	v12462 = *(*float64)(unsafe.Add(mBase, uint32(v12229)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v12439)+24)) = v12462
	v12464 = *(*int32)(unsafe.Add(mBase, uint32(v12229)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v12439)+36)) = uint8(v12450)
	*(*int32)(unsafe.Add(mBase, uint32(v12439)+32)) = v12464
	v12468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12229)+37)))
	*(*float64)(unsafe.Add(mBase, uint32(v12439)+8)) = v12460
	*(*uint8)(unsafe.Add(mBase, uint32(v12439)+37)) = uint8(v12468)
	v12471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v12471 == v12444 {
		goto L1749
	} else {
		goto L1750
	}
L1749:
	;
	v12474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12439)+36)) = uint8(v12474)
	v12476 = *(*float64)(unsafe.Add(mBase, uint32(l1)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v12439)+88)) = v12476
	goto L1751
L1750:
	;
	goto L1751
L1751:
	;
	v12478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v12479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+76)))
	v12480 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	v12481 = *(*int32)(unsafe.Add(mBase, uint32(v47)+140))
	v12483 = F_palloc0(m, int32(104))
	mBase = m.M
	v12484 = m.ExcPending
	if v12484 != 0 {
		goto L1
	} else {
		goto L1752
	}
L1752:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+100)) = v12402
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+96)) = v12401
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+92)) = v12406
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+88)) = v12274
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+56)) = v12439
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+52)) = v12225
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+48)) = v12481
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+44)) = v12176
	*(*int32)(unsafe.Add(mBase, uint32(v12483))) = int32(359)
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+80)) = v12480
	*(*uint8)(unsafe.Add(mBase, uint32(v12483)+76)) = uint8(v12479)
	*(*int32)(unsafe.Add(mBase, uint32(v12483)+72)) = v12478
	v12501 = v12483
	goto L4
L1753:
	;
	v12560 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v12561 = F_order_qual_clauses(m, l0, v12560)
	mBase = m.M
	v12562 = m.ExcPending
	if v12562 != 0 {
		goto L1
	} else {
		goto L1754
	}
L1754:
	;
	v12564 = F_extract_actual_clauses(m, v12561, int32(1))
	mBase = m.M
	v12565 = m.ExcPending
	if v12565 != 0 {
		goto L1
	} else {
		goto L1755
	}
L1755:
	;
	if v12564 == int32(0) {
		v12573 = v12501
		v12580 = v47
		goto L3
	} else {
		goto L1756
	}
L1756:
	;
	v12568 = F_create_gating_plan(m, l0, l1, v12501, v12564)
	mBase = m.M
	v12569 = m.ExcPending
	if v12569 != 0 {
		goto L1
	} else {
		goto L1757
	}
L1757:
	;
	v12573 = v12568
	v12580 = v47
	goto L3
}
