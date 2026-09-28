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
	var v40 int32
	_ = v40
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
	var v72 int64
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v216 int64
	_ = v216
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
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
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int64
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int64
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v862 int64
	_ = v862
	var v864 int64
	_ = v864
	var v865 int64
	_ = v865
	var v869 int64
	_ = v869
	var v881 int32
	_ = v881
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int64
	_ = v934
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int64
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 float64
	_ = v1131
	var v1153 int32
	_ = v1153
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1227 int32
	_ = v1227
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1314 int64
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1404 int64
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1550 int64
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1559 float64
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1581 float64
	_ = v1581
	var v1582 float64
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1590 int32
	_ = v1590
	var v1591 float64
	_ = v1591
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	var v1606 float64
	_ = v1606
	var v1612 float64
	_ = v1612
	var v1618 float64
	_ = v1618
	var v1620 float64
	_ = v1620
	var v1623 float64
	_ = v1623
	var v1626 float64
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1646 int32
	_ = v1646
	var v1654 int32
	_ = v1654
	var v1656 float64
	_ = v1656
	var v1661 int64
	_ = v1661
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1693 int32
	_ = v1693
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1754 int32
	_ = v1754
	var v1755 int64
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1866 int32
	_ = v1866
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1906 int32
	_ = v1906
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1962 int32
	_ = v1962
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1994 int64
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2041 int64
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2098 int64
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2106 int32
	_ = v2106
	var v2107 float64
	_ = v2107
	var v2108 float64
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2177 int64
	_ = v2177
	var v2179 int64
	_ = v2179
	var v2180 int64
	_ = v2180
	var v2184 int64
	_ = v2184
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2214 int32
	_ = v2214
	var v2215 int64
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2228 int32
	_ = v2228
	var v2232 int32
	_ = v2232
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2274 int32
	_ = v2274
	var v2282 int32
	_ = v2282
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2310 int32
	_ = v2310
	var v2316 int32
	_ = v2316
	var v2329 int32
	_ = v2329
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[0]))
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
		v2310 = v21
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
	m.G0 = v2329 + int32(16)
	return v2316
L7:
	;
	v2316 = int32(0)
	v2329 = v2310
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
		v1171 = v21
		goto L9
	default:
		v2310 = v21
		goto L7
	}
L9:
	;
	v1174 = m.G0
	v1176 = v1174 - int32(80)
	m.G0 = v1176
	v1179 = l0 + int32(288)
	v1181 = l0 + int32(280)
	goto L238
L10:
	;
	v274 = int32(1)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v275 <= v274 {
		goto L57
	} else {
		goto L58
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)))
	if v32 != 0 {
		v1171 = v21
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
	v40 = v34
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v40
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
	v62 = int32(_a_F_ExecAgg_0)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int64)(m, v65, v67, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v63
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
		v40 = v80
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
	v163 = v100
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
	v108 = int32(0)
	v109 = v100
	goto L30
L28:
	;
	v140 = v100
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
	v127 = v124 + v108*int32(24)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	F_hashagg_spill_finish(m, l0, v127, v108)
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
	v140 = v131
	v155 = v136
	goto L29
L32:
	;
	v131 = v109 + v128
	v133 = v108 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v133 < v134 {
		v108 = v133
		v109 = v131
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
	v163 = v140
	goto L26
L35:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)) = uint8(v228)
	v230 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v230
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v240 = v236 + int32(4)
	v244 = int32(-1)
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v238)))
	if v245 == int64(0) {
		v267 = v244
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
	v187 = F_MemoryContextMemAllocated(m, v185, int32(1))
	mBase = m.M
	goto L39
L39:
	;
	goto L40
L40:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v195 = int32(1)
	v196 = F_MemoryContextMemAllocated(m, v194, v195)
	mBase = m.M
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	v201 = F_MemoryContextMemAllocated(m, v199, v195)
	mBase = m.M
	v202 = v187 + v163<<(uint(int32(13))%32) + v196 + v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v203) < base.Ui32(v202) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v202
	goto L43
L42:
	;
	goto L43
L43:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v206 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v216 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v216 == int64(0) {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	v209 = F_LogicalTapeSetBlocks(m, v206)
	mBase = m.M
	v211 = v209 << (uint(int64(3)) % 64)
	v212 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v211) <= base.Ui64(v212) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v211
	goto L44
L47:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v201), base.F64_convert_i64_u(v216)), float64(12))
	goto L36
L48:
	;
	v1171 = v21
	goto L9
L49:
	;
	v270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+8)) = uint8(v270)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v267
	goto L48
L50:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	v250 = int32(0)
	goto L51
L51:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v248+v250*int32(12))+4))
	if v258 != int32(1) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v267 = v244
	goto L49
L53:
	;
	v267 = v250
	goto L49
L54:
	;
	goto L55
L55:
	;
	v262 = v250 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v262)) < base.Ui64(v245) {
		v250 = v262
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	v278 = v274
	goto L59
L58:
	;
	v278 = v275
	goto L59
L59:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v289 = v278
	v293 = v279
	goto L60
L60:
	;
	F_ReScanExprContext(m, v284)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v2310 = v21
	goto L7
L62:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v306 < v289 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v310 = v306 + int32(1)
	goto L65
L64:
	;
	v310 = v289
	goto L65
L65:
	;
	if int32(0) <= v306 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v313 = v310
	goto L68
L67:
	;
	v313 = v289
	goto L68
L68:
	;
	if int32(0) < v313 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v318 = int32(0)
	goto L72
L70:
	;
	v348 = v306
	goto L71
L71:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)))
	v364 = int32(1)
	v367 = v289 - v364
	if base.B2i32(v363 != v364)|base.B2i32(v348 < v367) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v334+v318<<(uint(int32(2))%32))))
	F_ReScanExprContext(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L74
	}
L73:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v348 = v344
	goto L71
L74:
	;
	v342 = v318 + int32(1)
	if v342 != v313 {
		v318 = v342
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v283)+8)) = v474
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)))
	if v476 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L77:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v372 < v373-int32(1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v458 = int32(0)
	if base.B2i32(v348 < v458)|base.B2i32(v367 <= v348) != 0 {
		v469 = v313
		v470 = v367
		v471 = v289
		v472 = v293
		v473 = v458
		goto L76
	} else {
		goto L107
	}
L80:
	;
	F_initialize_phase(m, l0, v372+int32(1))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v395 == int32(3) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(-1)
	v383 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v383)
	v385 = int32(1)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v387 <= v385 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v390 = v385
	goto L86
L85:
	;
	v390 = v387
	goto L86
L86:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v386)+20))
	v469 = v390
	v470 = v390 - int32(1)
	v471 = v390
	v472 = v393
	v473 = int32(0)
	goto L76
L87:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v398 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v456 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v456)
	v2310 = v21
	goto L7
L90:
	;
	F_tuplesort_end(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v403 != 0 {
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
	F_tuplesort_end(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)) = uint8(v408)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v412
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v418 = v414 + int32(4)
	v422 = int32(-1)
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v416)))
	if v423 == int64(0) {
		v445 = v422
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
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v454
	v1171 = v21
	goto L9
L99:
	;
	v448 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+8)) = uint8(v448)
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v445
	goto L98
L100:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v416)+20))
	v428 = int32(0)
	goto L101
L101:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v426+v428*int32(12))+4))
	if v436 != int32(1) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v445 = v422
	goto L99
L103:
	;
	v445 = v428
	goto L99
L104:
	;
	goto L105
L105:
	;
	v440 = v428 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v440)) < base.Ui64(v423) {
		v428 = v440
		goto L101
	} else {
		goto L106
	}
L106:
	;
	goto L102
L107:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v464+v348<<(uint(int32(2))%32))+4))
	v469 = v313
	v470 = v367
	v471 = v289
	v472 = v293
	v473 = v468
	goto L76
L108:
	;
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)))
	if v1153 == int32(0) {
		v289 = v471
		v293 = v472
		goto L60
	} else {
		goto L235
	}
L109:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)+12))
	if v985 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(0)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v531 != 0 {
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v472)+72))
	if v479 == int32(0) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v524 = v522 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v524
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	v966 = v524
	v983 = v526
	goto L109
L114:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if base.B2i32(v482 == int32(-1))|base.B2i32(v470 <= v482)|base.B2i32(v473 <= int32(0)) != 0 {
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+16))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v491+v473<<(uint(int32(2))%32)-int32(4))))
	if v497 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v283)+20))
	F_MemoryContextReset(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v503 = int32(_a_F_ExecAgg_0)
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v283)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v506
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v497)+24))
	v511 = m.T0[v510].(func(*base.Module, int32, int32, int32) int64)(m, v497, v283, v21+int32(13))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L120
	}
L119:
	;
	goto L110
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v504
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v283)+20))
	F_MemoryContextReset(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	if v511 != int64(0) {
		goto L110
	} else {
		goto L122
	}
L122:
	;
	goto L113
L123:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v608 = int32(0)
	goto L141
L124:
	;
	v532 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L126
	}
L125:
	;
	if int32(0) < v275 {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	if v532 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+4)))
	if v536&int32(2) != 0 {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v532)+8))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+44))
	v541 = m.T0[v540].(func(*base.Module, int32) int32)(m, v532)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v541
	goto L123
L130:
	;
	v546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v546)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v552 = v548
	goto L133
L131:
	;
	goto L132
L132:
	;
	v581 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v581)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v472)+72))
	if v583 != 0 {
		v2310 = v21
		goto L7
	} else {
		goto L140
	}
L133:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v549)+8))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v568+v552<<(uint(int32(2))%32))))
	if int32(0) < v572 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v579 < v471 {
		goto L123
	} else {
		goto L139
	}
L135:
	;
	v576 = v552 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v576
	if v576 < v471 {
		v552 = v576
		goto L133
	} else {
		goto L138
	}
L136:
	;
	v579 = v552
	goto L137
L137:
	;
	goto L134
L138:
	;
	v579 = v576
	goto L137
L139:
	;
	goto L108
L140:
	;
	goto L123
L141:
	;
	v624 = v608 << (uint(int32(2)) % 32)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v281+v624)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v627+v624)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v608
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v629
	v632 = int32(0)
	if v632 < v603 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v685 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L143:
	;
	v637 = v632
	goto L146
L144:
	;
	goto L145
L145:
	;
	v683 = v608 + int32(1)
	if v683 != v469 {
		v608 = v683
		goto L141
	} else {
		goto L150
	}
L146:
	;
	F_initialize_aggregate(m, l0, v602+v637*int32(240), v626+v637<<(uint(int32(4))%32))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L4
	} else {
		goto L148
	}
L147:
	;
	goto L145
L148:
	;
	v662 = v637 + int32(1)
	if v662 != v603 {
		v637 = v662
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	goto L142
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = v280
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v966 = v964
	v983 = v280
	goto L109
L152:
	;
	F_ExecForceStoreHeapTuple(m, v685, v280, int32(1))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v283)+12)) = v280
	goto L154
L154:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v712 != int32(3) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v737)+8))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)+44))
	v942 = m.T0[v941].(func(*base.Module, int32) int32)(m, v737)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L4
	} else {
		goto L204
	}
L156:
	;
	v720 = int32(_a_F_ExecAgg_0)
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+28))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v725)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v726
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v723)+24))
	v730 = m.T0[v729].(func(*base.Module, int32, int32, int32) int64)(m, v723, v725, int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L160
	}
L157:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v715 != int32(1) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	F_lookup_hash_entries(m, l0)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v721
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v283)+20))
	F_MemoryContextReset(m, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v737 = F_fetch_input_tuple(m, l0)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L4
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+12)) = v737
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v472)+72))
	if v908 == int32(0) {
		goto L154
	} else {
		goto L199
	}
L163:
	;
	if v737 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+4)))
	if v739&int32(2) == int32(0) {
		goto L162
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v744 != int32(3) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L166
L168:
	;
	if int32(0) < v275 {
		goto L196
	} else {
		goto L197
	}
L169:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v747 != int32(1) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v750 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v833&int32(-2) != int32(2) {
		goto L184
	} else {
		goto L185
	}
L172:
	;
	v816 = int32(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v754 = int32(0)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v754 < v756 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v761 = v754
	v762 = v754
	goto L178
L176:
	;
	v793 = v754
	v808 = v750
	goto L177
L177:
	;
	F_pfree(m, v808)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L4
	} else {
		goto L182
	}
L178:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v780 = v777 + v761*int32(24)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	F_hashagg_spill_finish(m, l0, v780, v761)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L4
	} else {
		goto L180
	}
L179:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v793 = v784
	v808 = v789
	goto L177
L180:
	;
	v784 = v762 + v781
	v786 = v761 + int32(1)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v786 < v787 {
		v761 = v786
		v762 = v784
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	v816 = v793
	goto L171
L183:
	;
	v881 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v881)
	goto L168
L184:
	;
	goto L183
L185:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v840 = F_MemoryContextMemAllocated(m, v838, int32(1))
	mBase = m.M
	goto L187
L187:
	;
	goto L188
L188:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v848 = int32(1)
	v849 = F_MemoryContextMemAllocated(m, v847, v848)
	mBase = m.M
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v851)+20))
	v854 = F_MemoryContextMemAllocated(m, v852, v848)
	mBase = m.M
	v855 = v840 + v816<<(uint(int32(13))%32) + v849 + v854
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v856) < base.Ui32(v855) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v855
	goto L191
L190:
	;
	goto L191
L191:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v859 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v869 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v869 == int64(0) {
		goto L184
	} else {
		goto L195
	}
L193:
	;
	v862 = F_LogicalTapeSetBlocks(m, v859)
	mBase = m.M
	v864 = v862 << (uint(int64(3)) % 64)
	v865 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v864) <= base.Ui64(v865) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v864
	goto L192
L195:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v854), base.F64_convert_i64_u(v869)), float64(12))
	goto L184
L196:
	;
	v903 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v903)
	goto L151
L197:
	;
	goto L198
L198:
	;
	v905 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v905)
	goto L151
L199:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v472)+80))
	if v911 <= int32(0) {
		goto L154
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+8)) = v280
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v915)+16))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v472)+80))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v916+v917<<(uint(int32(2))%32)-int32(4))))
	if v923 == int32(0) {
		goto L154
	} else {
		goto L201
	}
L201:
	;
	v926 = int32(_a_F_ExecAgg_0)
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v283)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v929
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v923)+24))
	v934 = m.T0[v933].(func(*base.Module, int32, int32, int32) int64)(m, v923, v283, v21+int32(14))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v927
	if v934 != int64(0) {
		goto L154
	} else {
		goto L203
	}
L203:
	;
	goto L155
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v942
	goto L151
L205:
	;
	v1073 = v966 << (uint(int32(2)) % 32)
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1073+v1074)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1076
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1073+v281)))
	F_finalize_aggregates(m, l0, v282, v1080)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L4
	} else {
		goto L225
	}
L206:
	;
	v988 = int32(2)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v985+v966<<(uint(v988)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v991
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+4)))
	if v993&v988 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	F_ExecStoreAllNullTuple(m, v983)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L4
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v998 == int32(0) {
		goto L205
	} else {
		goto L211
	}
L210:
	;
	goto L205
L211:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v998)+12))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	v1003 = int32(*(*int16)(unsafe.Add(mBase, uint32(v983)+6)))
	if v1003 < v1002 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v983)+8))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1005)+16))
	m.T0[v1006].(func(*base.Module, int32, int32))(m, v983, v1002)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L4
	} else {
		goto L215
	}
L213:
	;
	v1012 = v998
	goto L214
L214:
	;
	v1013 = int32(0)
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+4))
	if v1014 <= v1013 {
		goto L205
	} else {
		goto L217
	}
L215:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1009 == int32(0) {
		goto L205
	} else {
		goto L216
	}
L216:
	;
	v1012 = v1009
	goto L214
L217:
	;
	v1019 = v1013
	goto L218
L218:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+12))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1035+v1019<<(uint(int32(2))%32))))
	v1040 = F_bms_is_member(m, v1039, v991)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L4
	} else {
		goto L220
	}
L219:
	;
	goto L205
L220:
	;
	if v1040 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v983)+20))
	v1046 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1044+v1039-v1046))) = uint8(v1046)
	goto L223
L222:
	;
	goto L223
L223:
	;
	v1051 = v1019 + int32(1)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+4))
	if v1051 < v1052 {
		v1019 = v1051
		goto L218
	} else {
		goto L224
	}
L224:
	;
	goto L219
L225:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1083 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1128 == int32(0) {
		goto L108
	} else {
		goto L234
	}
L227:
	;
	v1084 = int32(_a_F_ExecAgg_0)
	v1085 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1088
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+24))
	v1093 = m.T0[v1092].(func(*base.Module, int32, int32, int32) int64)(m, v1083, v1087, v21+int32(15))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L4
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+80))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+24))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1103)+8))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	m.T0[v1105].(func(*base.Module, int32))(m, v1103)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L4
	} else {
		goto L232
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1085
	if v1093 == int64(0) {
		goto L226
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v1108 = int32(_a_F_ExecAgg_0)
	v1109 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1111
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+32))
	v1117 = m.T0[v1116].(func(*base.Module, int32, int32, int32) int64)(m, v1101+int32(8), v1102, int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1109
	v1121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1103)+4)))
	v1123 = v1121 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1103)+4)) = uint16(v1123)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1103)+12))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1125)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1103)+6)) = uint16(v1126)
	v2316 = v1103
	v2329 = v21
	goto L6
L234:
	;
	v1131 = *(*float64)(unsafe.Add(mBase, uint32(v1128)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v1128)+424)) = base.F64_add(v1131, float64(1))
	goto L108
L235:
	;
	goto L61
L236:
	;
	if v1767 == int32(0) {
		v2310 = v1171
		goto L7
	} else {
		goto L485
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L4
	} else {
		goto L481
	}
L238:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1211 = v1200 + v1201*int32(52)
	goto L243
L239:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L4
	} else {
		goto L477
	}
L240:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = int64(0)
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	if v1808 != int32(3) {
		goto L373
	} else {
		goto L374
	}
L241:
	;
	m.G0 = v1176 + int32(80)
	goto L236
L242:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+80))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+24))
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1741)+8))
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+12))
	m.T0[v1743].(func(*base.Module, int32))(m, v1741)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L4
	} else {
		goto L370
	}
L243:
	;
	v1227 = v1211 + int32(4)
	goto L245
L244:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v1563 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L245:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1211)))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+16))
	v1249 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[0]))
	if v1249 != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	goto L244
L247:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L4
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1246)))
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+8)))
	v1259 = v1256
	goto L253
L250:
	;
	goto L249
L251:
	;
	goto L246
L252:
	;
	if v1291 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L253:
	;
	if v1259&int32(1) != 0 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1291 = v1274
	goto L252
L255:
	;
	v1291 = int32(0)
	goto L252
L256:
	;
	goto L257
L257:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+20))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+12))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1227)))
	v1270 = v1266 & (v1267 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v1227))) = v1270
	v1274 = v1265 + v1267*int32(12)
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+12))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	v1278 = v1275 & (v1276 ^ v1270)
	if v1278 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1227)+8)) = uint8(v1281)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	if v1285 != int32(1) {
		v1259 = base.B2i32(v1278 == int32(0))
		goto L253
	} else {
		goto L261
	}
L261:
	;
	goto L254
L262:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1296 = v1294 + int32(1)
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1297 <= v1296 {
		goto L251
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+20))
	F_MemoryContextReset(m, v1343)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L4
	} else {
		goto L275
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1296
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1300
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1305 = v1302 + v1296*int32(52)
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1306)))
	v1309 = v1305 + int32(4)
	v1313 = int32(-1)
	v1314 = *(*int64)(unsafe.Add(mBase, uint32(v1307)))
	if v1314 == int64(0) {
		v1336 = v1313
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1211 = v1305
	goto L243
L267:
	;
	v1339 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1309)+8)) = uint8(v1339)
	*(*int32)(unsafe.Add(mBase, uint32(v1309)+4)) = v1336
	*(*int32)(unsafe.Add(mBase, uint32(v1309))) = v1336
	goto L266
L268:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+20))
	v1319 = int32(0)
	goto L269
L269:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1317+v1319*int32(12))+4))
	if v1327 != int32(1) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1336 = v1313
	goto L267
L271:
	;
	v1336 = v1319
	goto L267
L272:
	;
	goto L273
L273:
	;
	v1331 = v1319 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1331)) < base.Ui64(v1314) {
		v1319 = v1331
		goto L269
	} else {
		goto L274
	}
L274:
	;
	goto L270
L275:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1348 = F_ExecStoreMinimalTuple(m, v1346, v1247, int32(0))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L4
	} else {
		goto L276
	}
L276:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+12))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)))
	v1352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1247)+6)))
	if v1352 < v1351 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+8))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+16))
	m.T0[v1355].(func(*base.Module, int32, int32))(m, v1247, v1351)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L4
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+8))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+12))
	m.T0[v1359].(func(*base.Module, int32))(m, v1205)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L4
	} else {
		goto L281
	}
L280:
	;
	goto L279
L281:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+12))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1362)))
	if v1363 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+20))
	base.MemoryFill(m, v1364, int32(1), v1363)
	goto L284
L283:
	;
	goto L284
L284:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+32))
	if int32(0) < v1367 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1372 = int32(0)
	goto L288
L286:
	;
	goto L287
L287:
	;
	v1434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1205)+4)))
	v1436 = v1434 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1205)+4)) = uint16(v1436)
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+12))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1438)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1205)+6)) = uint16(v1439)
	goto L291
L288:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+16))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+40))
	v1391 = int32(1)
	v1394 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1390+v1372<<(uint(v1391)%32)))))
	v1396 = v1394 - v1391
	v1397 = int32(3)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+16))
	v1404 = *(*int64)(unsafe.Add(mBase, uint32(v1400+v1372<<(uint(v1397)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1389+v1396<<(uint(v1397)%32)))) = v1404
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+20))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+20))
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1408+v1372))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1406+v1396))) = uint8(v1410)
	v1413 = v1372 + v1391
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+32))
	if v1413 < v1414 {
		v1372 = v1413
		goto L288
	} else {
		goto L290
	}
L289:
	;
	goto L287
L290:
	;
	goto L289
L291:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+32))
	if v1441 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1291)))
	v1445 = v1442 - v1441
	goto L294
L293:
	;
	v1445 = int32(0)
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1207)+12)) = v1205
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+12))
	if v1448 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	F_finalize_aggregates(m, l0, v1206, v1445)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L4
	} else {
		goto L315
	}
L296:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v1452 = int32(2)
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1448+v1451<<(uint(v1452)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v1455
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205)+4)))
	if v1457&v1452 != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	F_ExecStoreAllNullTuple(m, v1205)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L4
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1462 == int32(0) {
		goto L295
	} else {
		goto L301
	}
L300:
	;
	goto L295
L301:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1462)+12))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1465)))
	v1467 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1205)+6)))
	if v1467 < v1466 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+8))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+16))
	m.T0[v1470].(func(*base.Module, int32, int32))(m, v1205, v1466)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L4
	} else {
		goto L305
	}
L303:
	;
	v1476 = v1462
	goto L304
L304:
	;
	v1477 = int32(0)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+4))
	if v1478 <= v1477 {
		goto L295
	} else {
		goto L307
	}
L305:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1473 == int32(0) {
		goto L295
	} else {
		goto L306
	}
L306:
	;
	v1476 = v1473
	goto L304
L307:
	;
	v1482 = v1477
	goto L308
L308:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+12))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1499+v1482<<(uint(int32(2))%32))))
	v1504 = F_bms_is_member(m, v1503, v1455)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L4
	} else {
		goto L310
	}
L309:
	;
	goto L295
L310:
	;
	if v1504 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+20))
	v1510 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1508+v1503-v1510))) = uint8(v1510)
	goto L313
L312:
	;
	goto L313
L313:
	;
	v1515 = v1482 + int32(1)
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1476)+4))
	if v1515 < v1516 {
		v1482 = v1515
		goto L308
	} else {
		goto L314
	}
L314:
	;
	goto L309
L315:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1538 == int32(0) {
		goto L242
	} else {
		goto L316
	}
L316:
	;
	v1541 = int32(_a_F_ExecAgg_0)
	v1542 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1545
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+24))
	v1550 = m.T0[v1549].(func(*base.Module, int32, int32, int32) int64)(m, v1538, v1544, v1176+int32(48))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L4
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1542
	if v1550 != int64(0) {
		goto L242
	} else {
		goto L318
	}
L318:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1556 == int32(0) {
		goto L245
	} else {
		goto L319
	}
L319:
	;
	v1559 = *(*float64)(unsafe.Add(mBase, uint32(v1556)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v1556)+424)) = base.F64_add(v1559, float64(1))
	goto L245
L320:
	;
	v1566 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+181)) = uint8(v1566)
	v1767 = int32(0)
	goto L241
L321:
	;
	goto L322
L322:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1563)+12))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1563)+4))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1570+v1571<<(uint(int32(2))%32)-int32(4))))
	v1578 = F_list_delete_last(m, v1563)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L4
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v1578
	v1581 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	v1582 = *(*float64)(unsafe.Add(mBase, uint32(v1577)+24))
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+4))
	v1590 = F_get_hash_memory_limit(m)
	mBase = m.M
	v1591 = base.F64_convert_i32_u(v1590)
	if base.F64_ge(v1591, base.F64_mul(v1581, v1582)) != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v1673 = v1671 << (uint(int32(2)) % 32)
	if v1668&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v1673)) == int32(0) {
		goto L356
	} else {
		goto L357
	}
L325:
	;
	goto L329
L326:
	;
	goto L327
L327:
	;
	v1600 = int32(32)
	v1605 = F_get_hash_memory_limit(m)
	mBase = m.M
	v1606 = base.F64_convert_i32_u(v1605)
	v1612 = base.F64_mul(base.F64_add(base.F64_mul(v1606, float64(0.25)), float64(-8192)), float64(0.0001220703125))
	v1618 = base.F64_add(base.F64_div(base.F64_mul(v1581, base.F64_mul(v1582, float64(1.5))), v1606), float64(1))
	if base.F64_gt(v1618, v1612) != 0 {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	goto L330
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1181))) = v1590
	*(*int64)(unsafe.Add(mBase, uint32(v1179))) = base.I64_trunc_sat_f64_u(base.F64_div(v1591, v1581))
	goto L324
L331:
	;
	v1620 = v1612
	goto L333
L332:
	;
	v1620 = v1618
	goto L333
L333:
	;
	if base.F64_lt(v1620, float64(4)) != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1623 = float64(4)
	goto L336
L335:
	;
	v1623 = v1620
	goto L336
L336:
	;
	if base.F64_gt(v1623, float64(1024)) != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1626 = float64(1024)
	goto L339
L338:
	;
	v1626 = v1623
	goto L339
L339:
	;
	v1627 = base.I32_trunc_sat_f64_s(v1626)
	if base.Ui32(int32(2)) <= base.Ui32(v1627) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1635 = v1600 - base.I32_clz(v1627-int32(1))
	goto L342
L341:
	;
	v1635 = int32(0)
	goto L342
L342:
	;
	if int32(31) < v1583+v1635 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1639 = v1600 - v1583
	goto L345
L344:
	;
	v1639 = v1635
	goto L345
L345:
	;
	goto L347
L347:
	;
	goto L348
L348:
	;
	v1646 = int32(_a_F_ExecAgg_2)<<(uint(v1639)%32) - int32(-8192)
	if base.Ui32(v1646<<(uint(int32(2))%32)) < base.Ui32(v1590) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1654 = v1590 - v1646
	goto L351
L350:
	;
	v1654 = base.I32_trunc_sat_f64_u(base.F64_mul(v1591, float64(0.75)))
	goto L351
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1181))) = v1654
	v1656 = base.F64_convert_i32_u(v1654)
	if base.F64_gt(v1656, v1581) != 0 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1661 = base.I64_trunc_sat_f64_u(base.F64_div(v1656, v1581))
	goto L354
L353:
	;
	v1661 = int64(1)
	goto L354
L354:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1179))) = v1661
	goto L324
L355:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_ReScanExprContext(m, v1701)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L4
	} else {
		goto L364
	}
L356:
	;
	if v1673 == int32(0) {
		goto L355
	} else {
		goto L359
	}
L357:
	;
	v1693 = v1673
	goto L358
L358:
	;
	if v1693 == int32(0) {
		goto L355
	} else {
		goto L363
	}
L359:
	;
	v1683 = v1673 + v1668
	v1685 = v1668 + int32(4)
	if base.Ui32(v1685) < base.Ui32(v1683) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1687 = v1683
	goto L362
L361:
	;
	v1687 = v1685
	goto L362
L362:
	;
	v1693 = (v1668^int32(-1)+v1687)&int32(-4) + int32(4)
	goto L358
L363:
	;
	base.MemoryFill(m, v1668, int32(0), v1693)
	goto L355
L364:
	;
	v1704 = int32(0)
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1705 <= v1704 {
		goto L240
	} else {
		goto L365
	}
L365:
	;
	v1709 = v1704
	goto L366
L366:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1726+v1709*int32(52))))
	F_ResetTupleHashTable(m, v1730)
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L4
	} else {
		goto L368
	}
L367:
	;
	goto L240
L368:
	;
	v1734 = v1709 + int32(1)
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v1734 < v1735 {
		v1709 = v1734
		goto L366
	} else {
		goto L369
	}
L369:
	;
	goto L367
L370:
	;
	v1746 = int32(_a_F_ExecAgg_0)
	v1747 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1740)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1749
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+32))
	v1755 = m.T0[v1754].(func(*base.Module, int32, int32, int32) int64)(m, v1739+int32(8), v1740, int32(0))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L4
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v1747
	v1759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1741)+4)))
	v1761 = v1759 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1741)+4)) = uint16(v1761)
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1741)+12))
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1763)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1741)+6)) = uint16(v1764)
	v1767 = v1741
	goto L241
L372:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v1819
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v1821
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1826 != int32(2) {
		goto L376
	} else {
		goto L377
	}
L373:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v1818 = v1811
	goto L372
L374:
	;
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(1)
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1814 + int32(48)
	v1818 = v1814
	goto L372
L376:
	;
	v1829 = int32(48)
	goto L378
L377:
	;
	v1829 = int32(0)
	goto L378
L378:
	;
	v1830 = v1818 + v1829
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1830)+44))
	if v1831 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)))
	v1835 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v1835)
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(_a_F_ExecAgg_3)
	v1843 = F_ExecBuildAggTrans(m, l0, v1830, int32(0), v1835, v1835)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L4
	} else {
		goto L382
	}
L380:
	;
	v1849 = v1831
	goto L381
L381:
	;
	v1853 = v1823 + v1819*int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v1830)+28)) = v1849
	v1866 = int32(0)
	goto L384
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1830)+44)) = v1843
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v1834)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1837
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1830)+44))
	v1849 = v1848
	goto L381
L383:
	;
	goto L239
L384:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+16))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v1877 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1176)+47)) = uint8(v1877)
	v1879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)))
	v1881 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[0]))
	if v1881 != 0 {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+8))
	F_LogicalTapeClose(m, v2131)
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L4
	} else {
		goto L449
	}
L386:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L4
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+8))
	v1888 = F_LogicalTapeRead(m, v1884, v1176+int32(72), int32(4))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L4
	} else {
		goto L391
	}
L389:
	;
	goto L388
L390:
	;
	goto L385
L391:
	;
	if v1888 != int32(4) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	if v1888 == int32(0) {
		goto L390
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+72))
	v1916 = F_LogicalTapeRead(m, v1884, v1176+int32(76), int32(4))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L4
	} else {
		goto L400
	}
L395:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L4
	} else {
		goto L396
	}
L396:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L4
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+8)) = v1888
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+4)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1176))) = v1884
	F_errmsg_internal(m, int32(_a_F_ExecAgg_4), v1176)
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L4
	} else {
		goto L398
	}
L398:
	;
	F_errfinish(m, int32(_a_F_ExecAgg_5), int32(3133), int32(_a_F_ExecAgg_6))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L4
	} else {
		goto L399
	}
L399:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L400:
	;
	if v1916 != int32(4) {
		goto L383
	} else {
		goto L401
	}
L401:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+76))
	v1921 = F_palloc(m, v1920)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L4
	} else {
		goto L402
	}
L402:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1921))) = v1923
	v1925 = int32(4)
	v1929 = F_LogicalTapeRead(m, v1884, v1921+v1925, v1923-v1925)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L4
	} else {
		goto L403
	}
L403:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+76))
	if v1929 != v1931-int32(4) {
		goto L237
	} else {
		goto L404
	}
L404:
	;
	v1936 = F_ExecStoreMinimalTuple(m, v1921, v1876, int32(1))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L4
	} else {
		goto L405
	}
L405:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v1938)+12)) = v1876
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+36))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1941)+12))
	v1943 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1942)+6)))
	if v1943 < v1940 {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1942)+8))
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+16))
	m.T0[v1946].(func(*base.Module, int32, int32))(m, v1942, v1940)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L4
	} else {
		goto L409
	}
L407:
	;
	goto L408
L408:
	;
	if v1879 != 0 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L408
L410:
	;
	v1952 = int32(0)
	goto L412
L411:
	;
	v1952 = v1176 + int32(47)
	goto L412
L412:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+8))
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1953)+12))
	m.T0[v1954].(func(*base.Module, int32))(m, v1875)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L4
	} else {
		goto L413
	}
L413:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+32))
	if int32(0) < v1957 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1962 = int32(0)
	goto L417
L415:
	;
	goto L416
L416:
	;
	v2024 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1875)+4)))
	v2026 = v2024 & int32(_a_F_ExecAgg_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1875)+4)) = uint16(v2026)
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+12))
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v2028)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1875)+6)) = uint16(v2029)
	goto L420
L417:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+16))
	v1980 = int32(3)
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1942)+16))
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+40))
	v1985 = int32(1)
	v1988 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1984+v1962<<(uint(v1985)%32)))))
	v1990 = v1988 - v1985
	v1994 = *(*int64)(unsafe.Add(mBase, uint32(v1983+v1990<<(uint(v1980)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1979+v1962<<(uint(v1980)%32)))) = v1994
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+20))
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1942)+20))
	v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1998+v1990))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1996+v1962))) = uint8(v2000)
	v2003 = v1962 + v1985
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+32))
	if v2003 < v2004 {
		v1962 = v2003
		goto L417
	} else {
		goto L419
	}
L418:
	;
	goto L416
L419:
	;
	goto L418
L420:
	;
	v2031 = m.G0
	v2033 = v2031 - int32(16)
	m.G0 = v2033
	v2035 = int32(_a_F_ExecAgg_0)
	v2036 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v1874)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v2038
	*(*int32)(unsafe.Add(mBase, uint32(v1874)+40)) = v1875
	v2041 = *(*int64)(unsafe.Add(mBase, uint32(v1874)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1874)+44)) = v2041
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1874)))
	if v1952 != 0 {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v2036
	m.G0 = v2033 + int32(16)
	if v2066 != 0 {
		goto L432
	} else {
		goto L433
	}
L422:
	;
	v2046 = F_tuplehash_insert_hash_internal(m, v2043, v1912, v2033+int32(15))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L4
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	v2064 = F_tuplehash_lookup_hash_internal(m, v2043, v1912)
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L4
	} else {
		goto L430
	}
L425:
	;
	v2048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2033)+15)))
	if v2048 == int32(1) {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2051 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1952))) = uint8(v2051)
	v2066 = v2046
	goto L421
L427:
	;
	goto L428
L428:
	;
	v2053 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1952))) = uint8(v2053)
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v1874)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v2056
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v1874)+32))
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+8))
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v2059)+48))
	v2061 = m.T0[v2060].(func(*base.Module, int32, int32) int32)(m, v1875, v2058)
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L4
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2046))) = v2061
	v2066 = v2046
	goto L421
L430:
	;
	v2066 = v2064
	goto L421
L431:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+20))
	F_MemoryContextReset(m, v2128)
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L4
	} else {
		goto L448
	}
L432:
	;
	v2072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176)+47)))
	if v2072 == int32(1) {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	goto L434
L434:
	;
	if v1866 == int32(0) {
		goto L443
	} else {
		goto L444
	}
L435:
	;
	F_initialize_hash_entry(m, l0, v1874, v2066)
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L4
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v1874)+32))
	if v2082 != 0 {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	goto L437
L439:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v2066)))
	v2086 = v2083 - v2082
	goto L441
L440:
	;
	v2086 = int32(0)
	goto L441
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2077+v2078<<(uint(int32(2))%32)))) = v2086
	v2088 = int32(_a_F_ExecAgg_0)
	v2089 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1]))
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2090)+28))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v2093)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v2094
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+24))
	v2098 = m.T0[v2097].(func(*base.Module, int32, int32, int32) int64)(m, v2091, v2093, int32(0))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L4
	} else {
		goto L442
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAgg[1])) = v2089
	v2126 = v1866
	goto L431
L443:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+4))
	v2107 = *(*float64)(unsafe.Add(mBase, uint32(v1577)+24))
	v2108 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	F_hashagg_spill_init(m, v1176+int32(48), v1569, v2106, v2107, v2108)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L4
	} else {
		goto L446
	}
L444:
	;
	goto L445
L445:
	;
	F_hashagg_spill_tuple(m, l0, v1176+int32(48), v1876, v1912)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L4
	} else {
		goto L447
	}
L446:
	;
	goto L445
L447:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	*(*int32)(unsafe.Add(mBase, uint32(v2115+v2116<<(uint(int32(2))%32)))) = int32(0)
	v2126 = int32(1)
	goto L431
L448:
	;
	v1866 = v2126
	goto L384
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v2136
	if v1866 != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	F_hashagg_spill_finish(m, l0, v1176+int32(48), v2141)
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L4
	} else {
		goto L453
	}
L451:
	;
	v2146 = int32(0)
	goto L452
L452:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2148&int32(-2) != int32(2) {
		goto L455
	} else {
		goto L456
	}
L453:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+48))
	v2146 = v2144
	goto L452
L454:
	;
	v2196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v2196)
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v2198
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v2200
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	v2206 = v2202 + v2203*int32(52)
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2206)))
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v2207)))
	v2210 = v2206 + int32(4)
	v2214 = int32(-1)
	v2215 = *(*int64)(unsafe.Add(mBase, uint32(v2208)))
	if v2215 == int64(0) {
		v2237 = v2214
		goto L468
	} else {
		goto L469
	}
L455:
	;
	goto L454
L456:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v2155 = F_MemoryContextMemAllocated(m, v2153, int32(1))
	mBase = m.M
	goto L457
L457:
	;
	goto L459
L459:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v2163 = int32(1)
	v2164 = F_MemoryContextMemAllocated(m, v2162, v2163)
	mBase = m.M
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+20))
	v2169 = F_MemoryContextMemAllocated(m, v2167, v2163)
	mBase = m.M
	v2170 = v2155 + (v2146<<(uint(int32(13))%32) - int32(-8192)) + v2164 + v2169
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	if base.Ui32(v2171) < base.Ui32(v2170) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v2170
	goto L462
L461:
	;
	goto L462
L462:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v2174 == int32(0) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v2184 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v2184 == int64(0) {
		goto L455
	} else {
		goto L466
	}
L464:
	;
	v2177 = F_LogicalTapeSetBlocks(m, v2174)
	mBase = m.M
	v2179 = v2177 << (uint(int64(3)) % 64)
	v2180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	if base.Ui64(v2179) <= base.Ui64(v2180) {
		goto L463
	} else {
		goto L465
	}
L465:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v2179
	goto L463
L466:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+304)) = base.F64_add(base.F64_div(base.F64_convert_i32_u(v2169), base.F64_convert_i64_u(v2184)), float64(12))
	goto L455
L467:
	;
	F_pfree(m, v1577)
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L4
	} else {
		goto L476
	}
L468:
	;
	v2240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2210)+8)) = uint8(v2240)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+4)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(v2210))) = v2237
	goto L467
L469:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2208)+20))
	v2220 = int32(0)
	goto L470
L470:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2218+v2220*int32(12))+4))
	if v2228 != int32(1) {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	v2237 = v2214
	goto L468
L472:
	;
	v2237 = v2220
	goto L468
L473:
	;
	goto L474
L474:
	;
	v2232 = v2220 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v2232)) < base.Ui64(v2215) {
		v2220 = v2232
		goto L470
	} else {
		goto L475
	}
L475:
	;
	goto L471
L476:
	;
	goto L238
L477:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L4
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+40)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+36)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+32)) = v1884
	F_errmsg_internal(m, int32(_a_F_ExecAgg_4), v1176+int32(32))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L4
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(_a_F_ExecAgg_5), int32(3142), int32(_a_F_ExecAgg_6))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L4
	} else {
		goto L480
	}
L480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L481:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L4
	} else {
		goto L482
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+16)) = v1884
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+24)) = v1929
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+20)) = v2274 - int32(4)
	F_errmsg_internal(m, int32(_a_F_ExecAgg_4), v1176+int32(16))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L4
	} else {
		goto L483
	}
L483:
	;
	F_errfinish(m, int32(_a_F_ExecAgg_5), int32(3154), int32(_a_F_ExecAgg_6))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L4
	} else {
		goto L484
	}
L484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L485:
	;
	v2290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1767)+4)))
	if v2290&int32(2) == int32(0) {
		v2316 = v1767
		v2329 = v1171
		goto L6
	} else {
		goto L486
	}
L486:
	;
	v2310 = v1171
	goto L7
}
