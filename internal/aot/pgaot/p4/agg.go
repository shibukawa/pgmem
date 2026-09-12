package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAgg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
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
	var v39 int32
	_ = v39
	var v56 int32
	_ = v56
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v213 int64
	_ = v213
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int64
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v848 int64
	_ = v848
	var v850 int64
	_ = v850
	var v851 int64
	_ = v851
	var v855 int64
	_ = v855
	var v867 int32
	_ = v867
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1030 float64
	_ = v1030
	var v1052 int32
	_ = v1052
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1126 int32
	_ = v1126
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1213 int64
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
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
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1369 float64
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1391 float64
	_ = v1391
	var v1392 float64
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1400 int32
	_ = v1400
	var v1401 float64
	_ = v1401
	var v1407 float64
	_ = v1407
	var v1415 int64
	_ = v1415
	var v1421 int32
	_ = v1421
	var v1422 float64
	_ = v1422
	var v1428 float64
	_ = v1428
	var v1434 float64
	_ = v1434
	var v1436 float64
	_ = v1436
	var v1439 float64
	_ = v1439
	var v1442 float64
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1466 float64
	_ = v1466
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 float64
	_ = v1478
	var v1482 float64
	_ = v1482
	var v1490 int64
	_ = v1490
	var v1507 int64
	_ = v1507
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1622 int32
	_ = v1622
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1817 int32
	_ = v1817
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1895 int64
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1960 int32
	_ = v1960
	var v1961 float64
	_ = v1961
	var v1962 float64
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1977 int32
	_ = v1977
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1990 int32
	_ = v1990
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2025 int32
	_ = v2025
	var v2028 int64
	_ = v2028
	var v2030 int64
	_ = v2030
	var v2031 int64
	_ = v2031
	var v2035 int64
	_ = v2035
	var v2047 int32
	_ = v2047
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
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2066 int64
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2111 int32
	_ = v2111
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2133 int32
	_ = v2133
	var v2138 int32
	_ = v2138
	var v2141 int32
	_ = v2141
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2181 int32
	_ = v2181
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)))
	if v29 != 0 {
		v2162 = v21
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v2181 + int32(16)
	return v2166
L7:
	;
	v2166 = int32(0)
	v2181 = v2162
	goto L6
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	switch v31 {
	case 0, 1:
		goto L10
	case 2:
		goto L11
	case 3:
		v1071 = v21
		goto L9
	default:
		v2162 = v21
		goto L7
	}
L9:
	;
	v1073 = m.G0
	v1075 = v1073 - int32(80)
	m.G0 = v1075
	v1078 = l0 + int32(288)
	v1080 = l0 + int32(280)
	goto L222
L10:
	;
	v271 = int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v272 <= v271 {
		goto L57
	} else {
		goto L58
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)))
	if v32 != 0 {
		v1071 = v21
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v34 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v101 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	if v34 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v39 = v34
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	if v56&int32(2) != 0 {
		goto L13
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v39
	F_lookup_hash_entries(m, l0)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v62 = int32(4470752)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int32)(m, v65, v67, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v63
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	F_MemoryContextReset(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v80 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v80 != 0 {
		v39 = v80
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if int32(0) < v102 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v164 = v100
	goto L26
L26:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v180&int32(-2) != int32(2) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v107 = int32(0)
	v110 = v100
	goto L30
L28:
	;
	v141 = v100
	v155 = v101
	goto L29
L29:
	;
	F_pfree(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L34
	}
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v127 = v124 + v107*int32(24)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	F_hashagg_spill_finish(m, l0, v127, v107)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v141 = v131
	v155 = v136
	goto L29
L32:
	;
	v131 = v110 + v128
	v133 = v107 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v133 < v134 {
		v107 = v133
		v110 = v131
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	v164 = v141
	goto L26
L35:
	;
	v225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)) = uint8(v225)
	v227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v227)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v227
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v237 = v233 + int32(4)
	v241 = int32(-1)
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v235)))
	if v242 == int64(0) {
		v264 = v241
		goto L49
	} else {
		goto L50
	}
L36:
	;
	goto L35
L37:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v186 = F_MemoryContextMemAllocated(m, v185)
	mBase = m.M
	goto L39
L39:
	;
	goto L40
L40:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v194 = F_MemoryContextMemAllocated(m, v193)
	mBase = m.M
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	v198 = F_MemoryContextMemAllocated(m, v197)
	mBase = m.M
	v199 = v186 + v164<<(uint(int32(13))%32) + v194 + v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v200) < base.Ui32(v199) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v199
	goto L43
L42:
	;
	goto L43
L43:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v203 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v213 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v213 == int64(0) {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	v206 = F_LogicalTapeSetBlocks(m, v203)
	mBase = m.M
	v208 = v206 << (uint(int64(3)) % 64)
	v209 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v208) <= base.Ui64(v209) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v208
	goto L44
L47:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v198), base.F64_convert_i64_u(v213)), float64(12))
	goto L36
L48:
	;
	v1071 = v21
	goto L9
L49:
	;
	v267 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v237)+8)) = uint8(v267)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v264
	goto L48
L50:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	v247 = int32(0)
	goto L51
L51:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v245+v247*int32(12))+4))
	if v255 != int32(1) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v264 = v241
	goto L49
L53:
	;
	v264 = v247
	goto L49
L54:
	;
	goto L55
L55:
	;
	v259 = v247 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v259)) < base.Ui64(v242) {
		v247 = v259
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	v275 = v271
	goto L59
L58:
	;
	v275 = v272
	goto L59
L59:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v289 = v276
	v295 = v275
	goto L60
L60:
	;
	F_ReScanExprContext(m, v281)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v2162 = v21
	goto L7
L62:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v303 < v295 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v307 = v303 + int32(1)
	goto L65
L64:
	;
	v307 = v295
	goto L65
L65:
	;
	if int32(0) <= v303 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v310 = v307
	goto L68
L67:
	;
	v310 = v295
	goto L68
L68:
	;
	if int32(0) < v310 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v314 = int32(0)
	goto L72
L70:
	;
	v346 = v303
	goto L71
L71:
	;
	v360 = int32(1)
	v361 = v295 - v360
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)))
	if v362 != v360 {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331+v314<<(uint(int32(2))%32))))
	F_ReScanExprContext(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L74
	}
L73:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v346 = v341
	goto L71
L74:
	;
	v339 = v314 + int32(1)
	if v339 != v310 {
		v314 = v339
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v468
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)))
	if v470 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L77:
	;
	v452 = int32(0)
	if v346 < v452 {
		v462 = v361
		v463 = v452
		v465 = v289
		v466 = v295
		v467 = v310
		goto L76
	} else {
		goto L107
	}
L78:
	;
	if v346 < v361 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v366 < v367-int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_initialize_phase(m, l0, v366+int32(1))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v389 == int32(3) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(-1)
	v377 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v377)
	v380 = int32(1)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v382 <= v380 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v385 = v380
	goto L86
L85:
	;
	v385 = v382
	goto L86
L86:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v381)+20))
	v462 = v385 - int32(1)
	v463 = v377
	v465 = v388
	v466 = v385
	v467 = v385
	goto L76
L87:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v392 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v450 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v450)
	v2162 = v21
	goto L7
L90:
	;
	F_tuplesort_end(m, v392)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v397 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = int32(0)
	goto L92
L94:
	;
	F_tuplesort_end(m, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v402 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)) = uint8(v402)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v412 = v408 + int32(4)
	v416 = int32(-1)
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v410)))
	if v417 == int64(0) {
		v439 = v416
		goto L99
	} else {
		goto L100
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = int32(0)
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = int32(0)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v448
	v1071 = v21
	goto L9
L99:
	;
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v412)+8)) = uint8(v442)
	*(*int32)(unsafe.Add(mBase, uint32(v412)+4)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = v439
	goto L98
L100:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v410)+20))
	v422 = int32(0)
	goto L101
L101:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v420+v422*int32(12))+4))
	if v430 != int32(1) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v439 = v416
	goto L99
L103:
	;
	v439 = v422
	goto L99
L104:
	;
	goto L105
L105:
	;
	v434 = v422 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v434)) < base.Ui64(v417) {
		v422 = v434
		goto L101
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	if v361 <= v346 {
		v462 = v361
		v463 = v452
		v465 = v289
		v466 = v295
		v467 = v310
		goto L76
	} else {
		goto L108
	}
L108:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+8))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457+v346<<(uint(int32(2))%32))+4))
	v462 = v361
	v463 = v461
	v465 = v289
	v466 = v295
	v467 = v310
	goto L76
L109:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)))
	if v1052 == int32(0) {
		v289 = v465
		v295 = v466
		goto L60
	} else {
		goto L219
	}
L110:
	;
	F_prepare_projection_slot(m, l0, v967, v950)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L4
	} else {
		goto L208
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(0)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v521 != 0 {
		goto L126
	} else {
		goto L127
	}
L112:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v465)+72))
	if v473 == int32(0) {
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v514 = v512 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v514
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	v950 = v514
	v967 = v516
	goto L110
L115:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v476 == int32(-1) {
		goto L111
	} else {
		goto L116
	}
L116:
	;
	if v462 <= v476 {
		goto L111
	} else {
		goto L117
	}
L117:
	;
	if v463 <= int32(0) {
		goto L111
	} else {
		goto L118
	}
L118:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+16))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v483+v463<<(uint(int32(2))%32)-int32(4))))
	if v489 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	F_MemoryContextReset(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v495 = int32(4470752)
	v496 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v498
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v489)+20))
	v503 = m.T0[v502].(func(*base.Module, int32, int32, int32) int32)(m, v489, v280, v21+int32(13))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L123
	}
L122:
	;
	goto L111
L123:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v496
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	F_MemoryContextReset(m, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	if v503 != 0 {
		goto L111
	} else {
		goto L125
	}
L125:
	;
	goto L114
L126:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v607 = int32(0)
	goto L144
L127:
	;
	v522 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L129
	}
L128:
	;
	if int32(0) < v272 {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	if v522 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+4)))
	if v526&int32(2) != 0 {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+44))
	v531 = m.T0[v530].(func(*base.Module, int32) int32)(m, v522)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v531
	goto L126
L133:
	;
	v536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v536)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v541 = v538
	goto L136
L134:
	;
	goto L135
L135:
	;
	v571 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v571)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v465)+72))
	if v573 != 0 {
		v2162 = v21
		goto L7
	} else {
		goto L143
	}
L136:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v539)+8))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v558+v541<<(uint(int32(2))%32))))
	if int32(0) < v562 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v569 < v466 {
		goto L126
	} else {
		goto L142
	}
L138:
	;
	v566 = v541 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v566
	if v566 < v466 {
		v541 = v566
		goto L136
	} else {
		goto L141
	}
L139:
	;
	v569 = v541
	goto L140
L140:
	;
	goto L137
L141:
	;
	v569 = v566
	goto L140
L142:
	;
	goto L109
L143:
	;
	goto L126
L144:
	;
	v614 = v607 << (uint(int32(2)) % 32)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v278+v614)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v617+v614)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v607
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v619
	v622 = int32(0)
	if v622 < v593 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v675 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L146:
	;
	v626 = v622
	goto L149
L147:
	;
	goto L148
L148:
	;
	v673 = v607 + int32(1)
	if v673 != v467 {
		v607 = v673
		goto L144
	} else {
		goto L153
	}
L149:
	;
	F_initialize_aggregate(m, l0, v592+v626*int32(224), v616+v626<<(uint(int32(3))%32))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L4
	} else {
		goto L151
	}
L150:
	;
	goto L148
L151:
	;
	v652 = v626 + int32(1)
	if v652 != v593 {
		v626 = v652
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	goto L145
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = v277
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v950 = v948
	v967 = v277
	goto L110
L155:
	;
	F_ExecForceStoreHeapTuple(m, v675, v277, int32(1))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v277
	goto L157
L157:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v702 != int32(3) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+44))
	v926 = m.T0[v925].(func(*base.Module, int32) int32)(m, v727)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L4
	} else {
		goto L207
	}
L159:
	;
	v710 = int32(4470752)
	v711 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+28))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v716
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v713)+20))
	v720 = m.T0[v719].(func(*base.Module, int32, int32, int32) int32)(m, v713, v715, int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L163
	}
L160:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v705 != int32(1) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	F_lookup_hash_entries(m, l0)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	goto L159
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v711
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	F_MemoryContextReset(m, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v727 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v727
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v465)+72))
	if v894 == int32(0) {
		goto L157
	} else {
		goto L202
	}
L166:
	;
	if v727 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+4)))
	if v729&int32(2) == int32(0) {
		goto L165
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v734 != int32(3) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	goto L169
L171:
	;
	if int32(0) < v272 {
		goto L199
	} else {
		goto L200
	}
L172:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v737 != int32(1) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v740 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v822&int32(-2) != int32(2) {
		goto L187
	} else {
		goto L188
	}
L175:
	;
	v803 = int32(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v744 = int32(0)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v744 < v746 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v750 = v744
	v753 = v744
	goto L181
L179:
	;
	v781 = v744
	v783 = v740
	goto L180
L180:
	;
	F_pfree(m, v783)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L185
	}
L181:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v770 = v767 + v750*int32(24)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	F_hashagg_spill_finish(m, l0, v770, v750)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L183
	}
L182:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v781 = v774
	v783 = v779
	goto L180
L183:
	;
	v774 = v753 + v771
	v776 = v750 + int32(1)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v776 < v777 {
		v750 = v776
		v753 = v774
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	v803 = v781
	goto L174
L186:
	;
	v867 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v867)
	goto L171
L187:
	;
	goto L186
L188:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v828 = F_MemoryContextMemAllocated(m, v827)
	mBase = m.M
	goto L190
L190:
	;
	goto L191
L191:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v836 = F_MemoryContextMemAllocated(m, v835)
	mBase = m.M
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)+20))
	v840 = F_MemoryContextMemAllocated(m, v839)
	mBase = m.M
	v841 = v828 + v803<<(uint(int32(13))%32) + v836 + v840
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v842) < base.Ui32(v841) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v841
	goto L194
L193:
	;
	goto L194
L194:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v845 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v855 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v855 == int64(0) {
		goto L187
	} else {
		goto L198
	}
L196:
	;
	v848 = F_LogicalTapeSetBlocks(m, v845)
	mBase = m.M
	v850 = v848 << (uint(int64(3)) % 64)
	v851 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v850) <= base.Ui64(v851) {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v850
	goto L195
L198:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v840), base.F64_convert_i64_u(v855)), float64(12))
	goto L187
L199:
	;
	v889 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v889)
	goto L154
L200:
	;
	goto L201
L201:
	;
	v891 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v891)
	goto L154
L202:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v465)+80))
	if v897 <= int32(0) {
		goto L157
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v277
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v901)+16))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v465)+80))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v902+v903<<(uint(int32(2))%32)-int32(4))))
	if v909 == int32(0) {
		goto L157
	} else {
		goto L204
	}
L204:
	;
	v912 = int32(4470752)
	v913 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v280)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v915
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v909)+20))
	v920 = m.T0[v919].(func(*base.Module, int32, int32, int32) int32)(m, v909, v280, v21+int32(14))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v913
	if v920 != 0 {
		goto L157
	} else {
		goto L206
	}
L206:
	;
	goto L158
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v926
	goto L154
L208:
	;
	v971 = v950 << (uint(int32(2)) % 32)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v971+v972)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v974
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v971+v278)))
	F_finalize_aggregates(m, l0, v279, v978)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L4
	} else {
		goto L209
	}
L209:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v981 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1027 == int32(0) {
		goto L109
	} else {
		goto L218
	}
L211:
	;
	v982 = int32(4470752)
	v983 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v985)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v986
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v981)+20))
	v991 = m.T0[v990].(func(*base.Module, int32, int32, int32) int32)(m, v981, v985, v21+int32(15))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L4
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+72))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+16))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+8))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+12))
	m.T0[v1004].(func(*base.Module, int32))(m, v1002)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L4
	} else {
		goto L216
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v983
	if v991 == int32(0) {
		goto L210
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v1007 = int32(4470752)
	v1008 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1010
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+24))
	v1016 = m.T0[v1015].(func(*base.Module, int32, int32, int32) int32)(m, v1000+int32(4), v1001, int32(0))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1008
	v1020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1002)+4)))
	v1022 = v1020 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v1002)+4)) = uint16(v1022)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+12))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1002)+6)) = uint16(v1025)
	v2166 = v1002
	v2181 = v21
	goto L6
L218:
	;
	v1030 = *(*float64)(unsafe.Add(mBase, uint32(v1027)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v1027)+240)) = base.F64_add(v1030, float64(1))
	goto L109
L219:
	;
	goto L61
L220:
	;
	if v1622 == int32(0) {
		v2162 = v1071
		goto L7
	} else {
		goto L456
	}
L221:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2120 = m.ExcPending
	if v2120 != 0 {
		goto L4
	} else {
		goto L452
	}
L222:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1113 = v1099 + v1100*int32(52)
	goto L227
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L4
	} else {
		goto L448
	}
L224:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = int64(0)
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1661)))
	if v1662 != int32(3) {
		goto L344
	} else {
		goto L345
	}
L225:
	;
	m.G0 = v1075 + int32(80)
	goto L220
L226:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+72))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+16))
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+8))
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+12))
	m.T0[v1597].(func(*base.Module, int32))(m, v1595)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L4
	} else {
		goto L341
	}
L227:
	;
	v1126 = v1113 + int32(4)
	goto L229
L228:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v1373 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L229:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1113)))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+16))
	v1148 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v1148 != 0 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L228
L231:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L4
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1145)))
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126)+8)))
	v1158 = v1155
	goto L237
L234:
	;
	goto L233
L235:
	;
	goto L230
L236:
	;
	if v1190 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L237:
	;
	if v1158&int32(1) != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1190 = v1173
	goto L236
L239:
	;
	v1190 = int32(0)
	goto L236
L240:
	;
	goto L241
L241:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+20))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	v1169 = v1165 & (v1166 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v1126))) = v1169
	v1173 = v1164 + v1166*int32(12)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+4))
	v1177 = v1174 & (v1175 ^ v1169)
	if v1177 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1180 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1126)+8)) = uint8(v1180)
	goto L244
L243:
	;
	goto L244
L244:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+4))
	if v1184 != int32(1) {
		v1158 = base.B2i32(v1177 == int32(0))
		goto L237
	} else {
		goto L245
	}
L245:
	;
	goto L238
L246:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1195 = v1193 + int32(1)
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1196 <= v1195 {
		goto L235
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+20))
	F_MemoryContextReset(m, v1242)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L4
	} else {
		goto L259
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1195
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1199
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1204 = v1201 + v1195*int32(52)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	v1208 = v1204 + int32(4)
	v1212 = int32(-1)
	v1213 = *(*int64)(unsafe.Add(mBase, uint32(v1206)))
	if v1213 == int64(0) {
		v1235 = v1212
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1113 = v1204
	goto L227
L251:
	;
	v1238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1208)+8)) = uint8(v1238)
	*(*int32)(unsafe.Add(mBase, uint32(v1208)+4)) = v1235
	*(*int32)(unsafe.Add(mBase, uint32(v1208))) = v1235
	goto L250
L252:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+20))
	v1218 = int32(0)
	goto L253
L253:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1216+v1218*int32(12))+4))
	if v1226 != int32(1) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1235 = v1212
	goto L251
L255:
	;
	v1235 = v1218
	goto L251
L256:
	;
	goto L257
L257:
	;
	v1230 = v1218 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1230)) < base.Ui64(v1213) {
		v1218 = v1230
		goto L253
	} else {
		goto L258
	}
L258:
	;
	goto L254
L259:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	v1247 = F_ExecStoreMinimalTuple(m, v1245, v1146, int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+12))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1249)))
	v1251 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1146)+6)))
	if v1251 < v1250 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	F_slot_getsomeattrs_int(m, v1146, v1250)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L4
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+8))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+12))
	m.T0[v1256].(func(*base.Module, int32))(m, v1104)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L4
	} else {
		goto L265
	}
L264:
	;
	goto L263
L265:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+20))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1261)))
	v1264 = F__emscripten_memset_bulkmem(m, v1259, base.I32_extend8_s(int32(1)), v1262)
	mBase = m.M
	goto L266
L266:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+32))
	if int32(0) < v1265 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1271 = int32(0)
	goto L270
L268:
	;
	goto L269
L269:
	;
	v1332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1104)+4)))
	v1334 = v1332 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v1104)+4)) = uint16(v1334)
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1336)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1104)+6)) = uint16(v1337)
	goto L273
L270:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+16))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+40))
	v1289 = int32(1)
	v1292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1288+v1271<<(uint(v1289)%32)))))
	v1294 = v1292 - v1289
	v1295 = int32(2)
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+16))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1298+v1271<<(uint(v1295)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1287+v1294<<(uint(v1295)%32)))) = v1302
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+20))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+20))
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306+v1271))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1304+v1294))) = uint8(v1308)
	v1311 = v1271 + v1289
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+32))
	if v1311 < v1312 {
		v1271 = v1311
		goto L270
	} else {
		goto L272
	}
L271:
	;
	goto L269
L272:
	;
	goto L271
L273:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+32))
	if v1339 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	v1343 = v1340 - v1339
	goto L276
L275:
	;
	v1343 = int32(0)
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+12)) = v1104
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	F_prepare_projection_slot(m, l0, v1104, v1345)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	F_finalize_aggregates(m, l0, v1105, v1343)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1350 == int32(0) {
		goto L226
	} else {
		goto L279
	}
L279:
	;
	v1353 = int32(4470752)
	v1354 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1357
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+20))
	v1362 = m.T0[v1361].(func(*base.Module, int32, int32, int32) int32)(m, v1350, v1356, v1075+int32(48))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1354
	if v1362 != 0 {
		goto L226
	} else {
		goto L281
	}
L281:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1366 == int32(0) {
		goto L229
	} else {
		goto L282
	}
L282:
	;
	v1369 = *(*float64)(unsafe.Add(mBase, uint32(v1366)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v1366)+240)) = base.F64_add(v1369, float64(1))
	goto L229
L283:
	;
	v1376 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v1376)
	v1622 = int32(0)
	goto L225
L284:
	;
	goto L285
L285:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+12))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+4))
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1380+v1381<<(uint(int32(2))%32)-int32(4))))
	v1388 = F_list_delete_last(m, v1373)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L4
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v1388
	v1391 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v1392 = *(*float64)(unsafe.Add(mBase, uint32(v1387)+24))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+4))
	v1400 = F_get_hash_memory_limit(m)
	mBase = m.M
	v1401 = base.F64_convert_i32_u(v1400)
	if base.F64_ge(v1401, base.F64_mul(v1391, v1392)) != 0 {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v1518 = v1516 << (uint(int32(2)) % 32)
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v1519&int32(3) != 0 {
		v1537 = v1518
		goto L326
	} else {
		goto L327
	}
L288:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1078))) = v1507
	goto L287
L289:
	;
	v1507 = int64(0)
	goto L288
L290:
	;
	goto L294
L291:
	;
	goto L292
L292:
	;
	v1421 = F_get_hash_memory_limit(m)
	mBase = m.M
	v1422 = base.F64_convert_i32_u(v1421)
	v1428 = base.F64_mul(base.F64_add(base.F64_mul(v1422, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v1434 = base.F64_add(base.F64_div(base.F64_mul(v1391, base.F64_mul(v1392, float64(1.5))), v1422), float64(1))
	if base.F64_gt(v1434, v1428) != 0 {
		goto L298
	} else {
		goto L299
	}
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080))) = v1400
	v1407 = base.F64_div(v1401, v1391)
	if base.F64_lt(v1407, float64(1.8446744073709552e+19))&base.F64_ge(v1407, float64(0)) == int32(0) {
		goto L289
	} else {
		goto L296
	}
L296:
	;
	v1415 = base.I64_trunc_f64_u(v1407)
	*(*int64)(unsafe.Add(mBase, uint32(v1078))) = v1415
	goto L287
L297:
	;
	v1449 = F_my_log2(m, v1448)
	mBase = m.M
	if int32(31) < v1393+v1449 {
		goto L310
	} else {
		goto L311
	}
L298:
	;
	v1436 = v1428
	goto L300
L299:
	;
	v1436 = v1434
	goto L300
L300:
	;
	if base.F64_lt(v1436, float64(4)) != 0 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1439 = float64(4)
	goto L303
L302:
	;
	v1439 = v1436
	goto L303
L303:
	;
	if base.F64_gt(v1439, float64(1024)) != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1442 = float64(1024)
	goto L306
L305:
	;
	v1442 = v1439
	goto L306
L306:
	;
	if base.F64_lt(base.F64_abs(v1442), float64(2.147483648e+09)) != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1446 = base.I32_trunc_f64_s(v1442)
	v1448 = v1446
	goto L297
L308:
	;
	goto L309
L309:
	;
	v1448 = int32(-2147483648)
	goto L297
L310:
	;
	v1453 = int32(32) - v1393
	goto L312
L311:
	;
	v1453 = v1449
	goto L312
L312:
	;
	goto L314
L314:
	;
	goto L315
L315:
	;
	v1460 = int32(8192)<<(uint(v1453)%32) - int32(-8192)
	v1466 = base.F64_mul(v1401, float64(0.75))
	if base.F64_lt(v1466, float64(4.294967296e+09))&base.F64_ge(v1466, float64(0)) != 0 {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	if base.Ui32(v1460<<(uint(int32(2))%32)) < base.Ui32(v1400) {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	v1472 = base.I32_trunc_f64_u(v1466)
	v1474 = v1472
	goto L316
L318:
	;
	goto L319
L319:
	;
	v1474 = int32(0)
	goto L316
L320:
	;
	v1475 = v1400 - v1460
	goto L322
L321:
	;
	v1475 = v1474
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1080))) = v1475
	v1478 = base.F64_convert_i32_u(v1475)
	if base.F64_lt(v1391, v1478) == int32(0) {
		v1507 = int64(1)
		goto L288
	} else {
		goto L323
	}
L323:
	;
	v1482 = base.F64_div(v1478, v1391)
	if base.F64_lt(v1482, float64(1.8446744073709552e+19))&base.F64_ge(v1482, float64(0)) == int32(0) {
		goto L289
	} else {
		goto L324
	}
L324:
	;
	v1490 = base.I64_trunc_f64_u(v1482)
	*(*int64)(unsafe.Add(mBase, uint32(v1078))) = v1490
	goto L287
L325:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_ReScanExprContext(m, v1544)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L4
	} else {
		goto L334
	}
L326:
	;
	v1541 = F__emscripten_memset_bulkmem(m, v1519, base.I32_extend8_s(int32(0)), v1537)
	mBase = m.M
	goto L333
L327:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v1518) {
		v1537 = v1518
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1524 = v1518 + v1519
	if base.Ui32(v1524) <= base.Ui32(v1519) {
		goto L325
	} else {
		goto L329
	}
L329:
	;
	v1529 = v1519 + int32(4)
	if base.Ui32(v1529) < base.Ui32(v1524) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1531 = v1524
	goto L332
L331:
	;
	v1531 = v1529
	goto L332
L332:
	;
	v1537 = (v1519^int32(-1)+v1531)&int32(-4) + int32(4)
	goto L326
L333:
	;
	goto L325
L334:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	F_MemoryContextReset(m, v1547)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L4
	} else {
		goto L335
	}
L335:
	;
	v1550 = int32(0)
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1551 <= v1550 {
		goto L224
	} else {
		goto L336
	}
L336:
	;
	v1556 = v1550
	goto L337
L337:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1572+v1556*int32(52))))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1576)))
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+20))
	v1579 = int32(0)
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	v1583 = F___memset(m, v1578, v1579, v1580*int32(12))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1577)+8)) = v1579
	goto L339
L338:
	;
	goto L224
L339:
	;
	v1587 = v1556 + int32(1)
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1587 < v1588 {
		v1556 = v1587
		goto L337
	} else {
		goto L340
	}
L340:
	;
	goto L338
L341:
	;
	v1600 = int32(4470752)
	v1601 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1603
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+24))
	v1609 = m.T0[v1608].(func(*base.Module, int32, int32, int32) int32)(m, v1593+int32(4), v1594, int32(0))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L4
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1601
	v1613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1595)+4)))
	v1615 = v1613 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v1595)+4)) = uint16(v1615)
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+12))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1617)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1595)+6)) = uint16(v1618)
	v1622 = v1595
	goto L225
L343:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1673
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1675
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1682 != int32(2) {
		goto L347
	} else {
		goto L348
	}
L344:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v1672 = v1665
	goto L343
L345:
	;
	goto L346
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(1)
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1668 + int32(48)
	v1672 = v1668
	goto L343
L347:
	;
	v1685 = int32(48)
	goto L349
L348:
	;
	v1685 = int32(0)
	goto L349
L349:
	;
	v1686 = v1672 + v1685
	v1688 = v1686 + int32(44)
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1688)))
	if v1689 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)))
	v1693 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v1693)
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1593852)
	v1701 = F_ExecBuildAggTrans(m, l0, v1686, int32(0), v1693, v1693)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L4
	} else {
		goto L353
	}
L351:
	;
	v1707 = v1689
	goto L352
L352:
	;
	v1709 = v1679 + v1673*int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v1686)+28)) = v1707
	v1713 = int32(0)
	goto L355
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1688))) = v1701
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v1692)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1695
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1688)))
	v1707 = v1706
	goto L352
L354:
	;
	goto L223
L355:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1709)))
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+16))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v1733 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1075)+47)) = uint8(v1733)
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)))
	v1737 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v1737 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+8))
	F_LogicalTapeClose(m, v1985)
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L4
	} else {
		goto L420
	}
L357:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L4
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+8))
	v1744 = F_LogicalTapeRead(m, v1740, v1075+int32(72), int32(4))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L4
	} else {
		goto L362
	}
L360:
	;
	goto L359
L361:
	;
	goto L356
L362:
	;
	if v1744 != int32(4) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	if v1744 == int32(0) {
		goto L361
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+72))
	v1772 = F_LogicalTapeRead(m, v1740, v1075+int32(76), int32(4))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L4
	} else {
		goto L371
	}
L366:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L4
	} else {
		goto L367
	}
L367:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+8)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+4)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1075))) = v1740
	F_errmsg_internal(m, int32(156701), v1075)
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L4
	} else {
		goto L369
	}
L369:
	;
	F_errfinish(m, int32(490595), int32(3131), int32(457230))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L4
	} else {
		goto L370
	}
L370:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L371:
	;
	if v1772 != int32(4) {
		goto L354
	} else {
		goto L372
	}
L372:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+76))
	v1777 = F_palloc(m, v1776)
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L4
	} else {
		goto L373
	}
L373:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1777))) = v1779
	v1781 = int32(4)
	v1785 = F_LogicalTapeRead(m, v1740, v1777+v1781, v1779-v1781)
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L4
	} else {
		goto L374
	}
L374:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+76))
	if v1785 != v1787-int32(4) {
		goto L221
	} else {
		goto L375
	}
L375:
	;
	v1792 = F_ExecStoreMinimalTuple(m, v1777, v1732, int32(1))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v1794)+12)) = v1732
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+36))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1797)+12))
	v1799 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1798)+6)))
	if v1799 < v1796 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	F_slot_getsomeattrs_int(m, v1798, v1796)
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L4
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	if v1735 != 0 {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	goto L379
L381:
	;
	v1806 = int32(0)
	goto L383
L382:
	;
	v1806 = v1075 + int32(47)
	goto L383
L383:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+8))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+12))
	m.T0[v1808].(func(*base.Module, int32))(m, v1731)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+32))
	if int32(0) < v1811 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1817 = int32(0)
	goto L388
L386:
	;
	goto L387
L387:
	;
	v1878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1731)+4)))
	v1880 = v1878 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v1731)+4)) = uint16(v1880)
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+12))
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1882)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1731)+6)) = uint16(v1883)
	goto L391
L388:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+16))
	v1834 = int32(2)
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+16))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+40))
	v1839 = int32(1)
	v1842 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1838+v1817<<(uint(v1839)%32)))))
	v1844 = v1842 - v1839
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1837+v1844<<(uint(v1834)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1833+v1817<<(uint(v1834)%32)))) = v1848
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+20))
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1798)+20))
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1852+v1844))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1850+v1817))) = uint8(v1854)
	v1857 = v1817 + v1839
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+32))
	if v1857 < v1858 {
		v1817 = v1857
		goto L388
	} else {
		goto L390
	}
L389:
	;
	goto L387
L390:
	;
	goto L389
L391:
	;
	v1885 = m.G0
	v1887 = v1885 - int32(16)
	m.G0 = v1887
	v1889 = int32(4470752)
	v1890 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+28))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1892
	*(*int32)(unsafe.Add(mBase, uint32(v1730)+40)) = v1731
	v1895 = *(*int64)(unsafe.Add(mBase, uint32(v1730)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1730)+44)) = v1895
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1730)))
	if v1806 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1890
	m.G0 = v1887 + int32(16)
	if v1920 != 0 {
		goto L403
	} else {
		goto L404
	}
L393:
	;
	v1900 = F_tuplehash_insert_hash_internal(m, v1897, v1768, v1887+int32(15))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L4
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1918 = F_tuplehash_lookup_hash_internal(m, v1897, v1768)
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L4
	} else {
		goto L401
	}
L396:
	;
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1887)+15)))
	if v1902 == int32(1) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1905 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1806))) = uint8(v1905)
	v1920 = v1900
	goto L392
L398:
	;
	goto L399
L399:
	;
	v1907 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1806))) = uint8(v1907)
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+24))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1910
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+32))
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1731)+8))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+48))
	v1915 = m.T0[v1914].(func(*base.Module, int32, int32) int32)(m, v1731, v1912)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L4
	} else {
		goto L400
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1900))) = v1915
	v1920 = v1900
	goto L392
L401:
	;
	v1920 = v1918
	goto L392
L402:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+20))
	F_MemoryContextReset(m, v1982)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L4
	} else {
		goto L419
	}
L403:
	;
	v1926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+47)))
	if v1926 == int32(1) {
		goto L406
	} else {
		goto L407
	}
L404:
	;
	goto L405
L405:
	;
	if v1713 == int32(0) {
		goto L414
	} else {
		goto L415
	}
L406:
	;
	F_initialize_hash_entry(m, l0, v1730, v1920)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L4
	} else {
		goto L409
	}
L407:
	;
	goto L408
L408:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+32))
	if v1936 != 0 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L408
L410:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1920)))
	v1940 = v1937 - v1936
	goto L412
L411:
	;
	v1940 = int32(0)
	goto L412
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1931+v1932<<(uint(int32(2))%32)))) = v1940
	v1942 = int32(4470752)
	v1943 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1944)+28))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1947)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1948
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+20))
	v1952 = m.T0[v1951].(func(*base.Module, int32, int32, int32) int32)(m, v1945, v1947, int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L4
	} else {
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1943
	v1977 = v1713
	goto L402
L414:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+4))
	v1961 = *(*float64)(unsafe.Add(mBase, uint32(v1387)+24))
	v1962 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	F_hashagg_spill_init(m, v1075+int32(48), v1379, v1960, v1961, v1962)
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L4
	} else {
		goto L417
	}
L415:
	;
	goto L416
L416:
	;
	F_hashagg_spill_tuple(m, l0, v1075+int32(48), v1732, v1768)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L4
	} else {
		goto L418
	}
L417:
	;
	goto L416
L418:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	*(*int32)(unsafe.Add(mBase, uint32(v1969+v1970<<(uint(int32(2))%32)))) = int32(0)
	v1977 = int32(1)
	goto L402
L419:
	;
	v1713 = v1977
	goto L355
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1990
	if v1713 != 0 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	F_hashagg_spill_finish(m, l0, v1075+int32(48), v1995)
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L4
	} else {
		goto L424
	}
L422:
	;
	v2000 = int32(0)
	goto L423
L423:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2002&int32(-2) != int32(2) {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+48))
	v2000 = v1998
	goto L423
L425:
	;
	v2047 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v2047)
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v2049
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2051
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v2057 = v2053 + v2054*int32(52)
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v2057)))
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2058)))
	v2061 = v2057 + int32(4)
	v2065 = int32(-1)
	v2066 = *(*int64)(unsafe.Add(mBase, uint32(v2059)))
	if v2066 == int64(0) {
		v2088 = v2065
		goto L439
	} else {
		goto L440
	}
L426:
	;
	goto L425
L427:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v2008 = F_MemoryContextMemAllocated(m, v2007)
	mBase = m.M
	goto L428
L428:
	;
	goto L430
L430:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v2016 = F_MemoryContextMemAllocated(m, v2015)
	mBase = m.M
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+20))
	v2020 = F_MemoryContextMemAllocated(m, v2019)
	mBase = m.M
	v2021 = v2008 + (v2000<<(uint(int32(13))%32) - int32(-8192)) + v2016 + v2020
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v2022) < base.Ui32(v2021) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v2021
	goto L433
L432:
	;
	goto L433
L433:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v2025 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v2035 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v2035 == int64(0) {
		goto L426
	} else {
		goto L437
	}
L435:
	;
	v2028 = F_LogicalTapeSetBlocks(m, v2025)
	mBase = m.M
	v2030 = v2028 << (uint(int64(3)) % 64)
	v2031 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v2030) <= base.Ui64(v2031) {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v2030
	goto L434
L437:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v2020), base.F64_convert_i64_u(v2035)), float64(12))
	goto L426
L438:
	;
	F_pfree(m, v1387)
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L4
	} else {
		goto L447
	}
L439:
	;
	v2091 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2061)+8)) = uint8(v2091)
	*(*int32)(unsafe.Add(mBase, uint32(v2061)+4)) = v2088
	*(*int32)(unsafe.Add(mBase, uint32(v2061))) = v2088
	goto L438
L440:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2059)+20))
	v2071 = int32(0)
	goto L441
L441:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2069+v2071*int32(12))+4))
	if v2079 != int32(1) {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	v2088 = v2065
	goto L439
L443:
	;
	v2088 = v2071
	goto L439
L444:
	;
	goto L445
L445:
	;
	v2083 = v2071 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v2083)) < base.Ui64(v2066) {
		v2071 = v2083
		goto L441
	} else {
		goto L446
	}
L446:
	;
	goto L442
L447:
	;
	goto L222
L448:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+40)) = v1772
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+36)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+32)) = v1740
	F_errmsg_internal(m, int32(156701), v1075+int32(32))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L4
	} else {
		goto L450
	}
L450:
	;
	F_errfinish(m, int32(490595), int32(3140), int32(457230))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L4
	} else {
		goto L451
	}
L451:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L452:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L4
	} else {
		goto L453
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+16)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+24)) = v1785
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+20)) = v2125 - int32(4)
	F_errmsg_internal(m, int32(156701), v1075+int32(16))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L4
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(490595), int32(3152), int32(457230))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L4
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L456:
	;
	v2141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1622)+4)))
	if v2141&int32(2) == int32(0) {
		v2166 = v1622
		v2181 = v1071
		goto L6
	} else {
		goto L457
	}
L457:
	;
	v2162 = v1071
	goto L7
}
