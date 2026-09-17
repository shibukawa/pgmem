package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
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
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1045 int32
	_ = v1045
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1209 int32
	_ = v1209
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1346 int32
	_ = v1346
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1485 int32
	_ = v1485
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1621 int32
	_ = v1621
	var v1629 int32
	_ = v1629
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1745 int32
	_ = v1745
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1879 int32
	_ = v1879
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1915 int32
	_ = v1915
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1955 int32
	_ = v1955
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2043 int32
	_ = v2043
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2098 int32
	_ = v2098
	var v2103 int32
	_ = v2103
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2122 int32
	_ = v2122
	var v2127 int32
	_ = v2127
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2215 int32
	_ = v2215
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2240 int32
	_ = v2240
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2279 int32
	_ = v2279
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2298 int32
	_ = v2298
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2316 int32
	_ = v2316
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2352 int32
	_ = v2352
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2375 int32
	_ = v2375
	var v2377 int32
	_ = v2377
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2390 int32
	_ = v2390
	var v2393 int32
	_ = v2393
	var v2398 int32
	_ = v2398
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2422 int32
	_ = v2422
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2442 int32
	_ = v2442
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2474 int32
	_ = v2474
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2510 int32
	_ = v2510
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2527 int32
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2534 int32
	_ = v2534
	var v2545 int32
	_ = v2545
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2558 int32
	_ = v2558
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2569 int32
	_ = v2569
	var v2571 int32
	_ = v2571
	var v2573 int32
	_ = v2573
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2584 int32
	_ = v2584
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2600 int32
	_ = v2600
	var v2610 int32
	_ = v2610
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2619 int32
	_ = v2619
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2638 int32
	_ = v2638
	var v2646 int32
	_ = v2646
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2683 int32
	_ = v2683
	var v2687 int32
	_ = v2687
	var v2691 int32
	_ = v2691
	var v2696 int32
	_ = v2696
	var v2701 int32
	_ = v2701
	var v2704 int32
	_ = v2704
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2722 int32
	_ = v2722
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2730 int32
	_ = v2730
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2738 int32
	_ = v2738
	var v2748 int32
	_ = v2748
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2776 int32
	_ = v2776
	var v2784 int32
	_ = v2784
	var v2792 int32
	_ = v2792
	var v2794 int32
	_ = v2794
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2819 int32
	_ = v2819
	var v2823 int32
	_ = v2823
	var v2827 int32
	_ = v2827
	var v2832 int32
	_ = v2832
	var v2837 int32
	_ = v2837
	var v2840 int32
	_ = v2840
	var v2843 int32
	_ = v2843
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2866 int32
	_ = v2866
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2874 int32
	_ = v2874
	var v2884 int32
	_ = v2884
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2912 int32
	_ = v2912
	var v2920 int32
	_ = v2920
	var v2928 int32
	_ = v2928
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2943 int32
	_ = v2943
	var v2947 int32
	_ = v2947
	var v2951 int32
	_ = v2951
	var v2956 int32
	_ = v2956
	var v2961 int32
	_ = v2961
	var v2964 int32
	_ = v2964
	var v2967 int32
	_ = v2967
	var v2969 int32
	_ = v2969
	var v2971 int32
	_ = v2971
	var v2976 int32
	_ = v2976
	var v2978 int32
	_ = v2978
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2986 int32
	_ = v2986
	var v2990 int32
	_ = v2990
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2998 int32
	_ = v2998
	var v3008 int32
	_ = v3008
	var v3013 int32
	_ = v3013
	var v3015 int32
	_ = v3015
	var v3017 int32
	_ = v3017
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3036 int32
	_ = v3036
	var v3044 int32
	_ = v3044
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3055 int32
	_ = v3055
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3081 int32
	_ = v3081
	var v3085 int32
	_ = v3085
	var v3087 int32
	_ = v3087
	var v3091 int32
	_ = v3091
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3125 int32
	_ = v3125
	var v3129 int32
	_ = v3129
	var v3133 int32
	_ = v3133
	var v3138 int32
	_ = v3138
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3149 int32
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3153 int32
	_ = v3153
	var v3158 int32
	_ = v3158
	var v3160 int32
	_ = v3160
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3172 int32
	_ = v3172
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3180 int32
	_ = v3180
	var v3190 int32
	_ = v3190
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3199 int32
	_ = v3199
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3226 int32
	_ = v3226
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3243 int32
	_ = v3243
	var v3254 int32
	_ = v3254
	var v3258 int32
	_ = v3258
	var v3262 int32
	_ = v3262
	var v3267 int32
	_ = v3267
	var v3272 int32
	_ = v3272
	var v3275 int32
	_ = v3275
	var v3278 int32
	_ = v3278
	var v3280 int32
	_ = v3280
	var v3282 int32
	_ = v3282
	var v3287 int32
	_ = v3287
	var v3289 int32
	_ = v3289
	var v3293 int32
	_ = v3293
	var v3295 int32
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3301 int32
	_ = v3301
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3309 int32
	_ = v3309
	var v3319 int32
	_ = v3319
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3328 int32
	_ = v3328
	var v3333 int32
	_ = v3333
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3347 int32
	_ = v3347
	var v3355 int32
	_ = v3355
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3385 int32
	_ = v3385
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3394 int32
	_ = v3394
	var v3398 int32
	_ = v3398
	var v3408 int32
	_ = v3408
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3419 int32
	_ = v3419
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3427 int32
	_ = v3427
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3434 int32
	_ = v3434
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3441 int32
	_ = v3441
	var v3446 int32
	_ = v3446
	var v3449 int32
	_ = v3449
	var v3450 int32
	_ = v3450
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3458 int32
	_ = v3458
	var v3470 int32
	_ = v3470
	var v3471 int32
	_ = v3471
	var v3474 int32
	_ = v3474
	var v3475 int32
	_ = v3475
	var v3477 int32
	_ = v3477
	var v3489 int32
	_ = v3489
	var v3491 int32
	_ = v3491
	var v3492 int32
	_ = v3492
	var v3501 int32
	_ = v3501
	var v3502 int32
	_ = v3502
	var v3507 int32
	_ = v3507
	var v3508 int32
	_ = v3508
	var v3521 int32
	_ = v3521
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3537 int32
	_ = v3537
	var v3540 int32
	_ = v3540
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3547 int32
	_ = v3547
	var v3550 int32
	_ = v3550
	var v3553 int32
	_ = v3553
	var v3565 int32
	_ = v3565
	var v3567 int32
	_ = v3567
	var v3570 int32
	_ = v3570
	var v3571 int32
	_ = v3571
	var v3574 int32
	_ = v3574
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3603 int32
	_ = v3603
	var v3604 int32
	_ = v3604
	var v3609 int32
	_ = v3609
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3622 int32
	_ = v3622
	var v3625 int32
	_ = v3625
	var v3629 int32
	_ = v3629
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v15 {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	case 5:
		goto L4
	case 6:
		goto L3
	default:
		v3629 = v4
		goto L1
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v3629
L2:
	;
	v3622 = int32(1)
	F_MarkBufferDirtyHint(m, l2, v3622)
	mBase = m.M
	v3625 = m.ExcPending
	if v3625 != 0 {
		goto L64
	} else {
		goto L1296
	}
L3:
	;
	v3603 = F_HeapTupleSatisfiesVacuumHorizon(m, l0, l2, v13+int32(12))
	mBase = m.M
	v3604 = m.ExcPending
	if v3604 != 0 {
		goto L64
	} else {
		goto L1288
	}
L4:
	;
	v3416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v3417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3416)+20)))
	v3418 = int32(768)
	v3419 = v3417 & v3418
	if v3419 != v3418 {
		goto L1223
	} else {
		goto L1224
	}
L5:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v2390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	if v2390&int32(256) != 0 {
		goto L864
	} else {
		goto L865
	}
L6:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2094)+20)))
	if v2095&int32(256) != 0 {
		goto L754
	} else {
		goto L755
	}
L7:
	;
	v3629 = int32(1)
	goto L1
L8:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	if v1089&int32(256) != 0 {
		goto L397
	} else {
		goto L398
	}
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	v18 = base.I32_extend16_s(v17)
	if v17&int32(256) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v773 = int32(1)
	v774 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	if v774&int32(2048)|v774&int32(128)|base.B2i32(v774&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64)) != 0 {
		v3629 = v773
		goto L1
	} else {
		goto L283
	}
L11:
	;
	v767 = v156 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)) = uint16(v767)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L64
	} else {
		goto L282
	}
L12:
	;
	if v18&int32(512) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v756 = int32(768)
	if v18&v756 == v756 {
		goto L10
	} else {
		goto L279
	}
L15:
	;
	v3629 = int32(0)
	goto L1
L16:
	;
	goto L17
L17:
	;
	if v18&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if base.Ui32(v28) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	if v18 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L21:
	;
	if v148 != 0 {
		goto L61
	} else {
		goto L62
	}
L22:
	;
	v148 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v39 == v28 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v148 = int32(1)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v43 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v148 = v140
	goto L21
L29:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v47 == int32(0) {
		v140 = int32(0)
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v111 = int32(0)
	v113 = v43 - int32(1)
	goto L51
L32:
	;
	v52 = v47
	goto L33
L33:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if v57 == int32(4) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v140 = int32(0)
	goto L28
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v52)+80))
	if v104 != 0 {
		v52 = v104
		goto L33
	} else {
		goto L50
	}
L36:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v60 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v63 = int32(1)
	if v28 == v60 {
		v140 = v63
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v67 = v65 - int32(1)
	if v67 < int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v72 = int32(0)
	v74 = v67
	goto L40
L40:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v80 = int32(2)
	v81 = base.I32_div_s(v74-v72, v80)
	v82 = v81 + v72
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78+v82<<(uint(v80)%32))))
	if v86 == v28 {
		v140 = v63
		goto L28
	} else {
		goto L42
	}
L41:
	;
	goto L35
L42:
	;
	v90 = F_TransactionIdPrecedes(m, v86, v28)
	mBase = m.M
	if v90 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v91 = v82 + int32(1)
	goto L45
L44:
	;
	v91 = v72
	goto L45
L45:
	;
	if v90 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v94 = v74
	goto L48
L47:
	;
	v94 = v82 - int32(1)
	goto L48
L48:
	;
	if v91 <= v94 {
		v72 = v91
		v74 = v94
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	goto L34
L51:
	;
	v118 = int32(2)
	v119 = base.I32_div_s(v113-v111, v118)
	v120 = v119 + v111
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109+v120<<(uint(v118)%32))))
	v125 = base.B2i32(v124 == v28)
	if v124 == v28 {
		v140 = v125
		goto L28
	} else {
		goto L53
	}
L52:
	;
	v140 = v125
	goto L28
L53:
	;
	v128 = base.B2i32(base.Ui32(v124) < base.Ui32(v28))
	if base.Ui32(v124) < base.Ui32(v28) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v129 = v120 + int32(1)
	goto L56
L55:
	;
	v129 = v111
	goto L56
L56:
	;
	if base.Ui32(v124) < base.Ui32(v28) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v132 = v113
	goto L59
L58:
	;
	v132 = v120 - int32(1)
	goto L59
L59:
	;
	if v129 <= v132 {
		v111 = v129
		v113 = v132
		goto L51
	} else {
		goto L60
	}
L60:
	;
	goto L52
L61:
	;
	v3629 = int32(0)
	goto L1
L62:
	;
	goto L63
L63:
	;
	v150 = F_XidInMVCCSnapshot(m, v28, l1)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return int32(0)
L65:
	;
	if v150 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v154 = F_TransactionIdDidCommit(m, v28)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	if v154 == int32(0) {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v160 = v156 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)) = uint16(v160)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v3629 = int32(0)
	goto L1
L70:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if base.Ui32(v168) < base.Ui32(int32(3)) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if base.Ui32(v308) < base.Ui32(int32(3)) {
		goto L125
	} else {
		goto L126
	}
L73:
	;
	if v288 != 0 {
		goto L10
	} else {
		goto L113
	}
L74:
	;
	v288 = int32(0)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v179 == v168 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v288 = int32(1)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v183 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v288 = v280
	goto L73
L81:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v187 == int32(0) {
		v280 = int32(0)
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v251 = int32(0)
	v253 = v183 - int32(1)
	goto L103
L84:
	;
	v192 = v187
	goto L85
L85:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	if v197 == int32(4) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v280 = int32(0)
	goto L80
L87:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v192)+80))
	if v244 != 0 {
		v192 = v244
		goto L85
	} else {
		goto L102
	}
L88:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v200 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v203 = int32(1)
	if v168 == v200 {
		v280 = v203
		goto L80
	} else {
		goto L90
	}
L90:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v192)+52))
	v207 = v205 - int32(1)
	if v207 < int32(0) {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v212 = int32(0)
	v214 = v207
	goto L92
L92:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v192)+48))
	v220 = int32(2)
	v221 = base.I32_div_s(v214-v212, v220)
	v222 = v221 + v212
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v218+v222<<(uint(v220)%32))))
	if v226 == v168 {
		v280 = v203
		goto L80
	} else {
		goto L94
	}
L93:
	;
	goto L87
L94:
	;
	v230 = F_TransactionIdPrecedes(m, v226, v168)
	mBase = m.M
	if v230 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v231 = v222 + int32(1)
	goto L97
L96:
	;
	v231 = v212
	goto L97
L97:
	;
	if v230 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v234 = v214
	goto L100
L99:
	;
	v234 = v222 - int32(1)
	goto L100
L100:
	;
	if v231 <= v234 {
		v212 = v231
		v214 = v234
		goto L92
	} else {
		goto L101
	}
L101:
	;
	goto L93
L102:
	;
	goto L86
L103:
	;
	v258 = int32(2)
	v259 = base.I32_div_s(v253-v251, v258)
	v260 = v259 + v251
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v249+v260<<(uint(v258)%32))))
	v265 = base.B2i32(v264 == v168)
	if v264 == v168 {
		v280 = v265
		goto L80
	} else {
		goto L105
	}
L104:
	;
	v280 = v265
	goto L80
L105:
	;
	v268 = base.B2i32(base.Ui32(v264) < base.Ui32(v168))
	if base.Ui32(v264) < base.Ui32(v168) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v269 = v260 + int32(1)
	goto L108
L107:
	;
	v269 = v251
	goto L108
L108:
	;
	if base.Ui32(v264) < base.Ui32(v168) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v272 = v253
	goto L111
L110:
	;
	v272 = v260 - int32(1)
	goto L111
L111:
	;
	if v269 <= v272 {
		v251 = v269
		v253 = v272
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L104
L113:
	;
	v289 = F_XidInMVCCSnapshot(m, v168, l1)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L64
	} else {
		goto L114
	}
L114:
	;
	if v289 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v3629 = int32(0)
	goto L1
L116:
	;
	goto L117
L117:
	;
	v292 = F_TransactionIdDidCommit(m, v168)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L64
	} else {
		goto L118
	}
L118:
	;
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	if v292 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v296 = v294 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)) = uint16(v296)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L64
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v302 = v294 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)) = uint16(v302)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L64
	} else {
		goto L123
	}
L122:
	;
	goto L10
L123:
	;
	v3629 = int32(0)
	goto L1
L124:
	;
	if v428 != 0 {
		goto L164
	} else {
		goto L165
	}
L125:
	;
	v428 = int32(0)
	goto L124
L126:
	;
	goto L127
L127:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v319 == v308 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v428 = int32(1)
	goto L124
L129:
	;
	goto L130
L130:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v323 <= int32(0) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v428 = v420
	goto L124
L132:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v327 == int32(0) {
		v420 = int32(0)
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v391 = int32(0)
	v393 = v323 - int32(1)
	goto L154
L135:
	;
	v332 = v327
	goto L136
L136:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	if v337 == int32(4) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v420 = int32(0)
	goto L131
L138:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v332)+80))
	if v384 != 0 {
		v332 = v384
		goto L136
	} else {
		goto L153
	}
L139:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v340 == int32(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v343 = int32(1)
	if v308 == v340 {
		v420 = v343
		goto L131
	} else {
		goto L141
	}
L141:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v332)+52))
	v347 = v345 - int32(1)
	if v347 < int32(0) {
		goto L138
	} else {
		goto L142
	}
L142:
	;
	v352 = int32(0)
	v354 = v347
	goto L143
L143:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	v360 = int32(2)
	v361 = base.I32_div_s(v354-v352, v360)
	v362 = v361 + v352
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v358+v362<<(uint(v360)%32))))
	if v366 == v308 {
		v420 = v343
		goto L131
	} else {
		goto L145
	}
L144:
	;
	goto L138
L145:
	;
	v370 = F_TransactionIdPrecedes(m, v366, v308)
	mBase = m.M
	if v370 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v371 = v362 + int32(1)
	goto L148
L147:
	;
	v371 = v352
	goto L148
L148:
	;
	if v370 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v374 = v354
	goto L151
L150:
	;
	v374 = v362 - int32(1)
	goto L151
L151:
	;
	if v371 <= v374 {
		v352 = v371
		v354 = v374
		goto L143
	} else {
		goto L152
	}
L152:
	;
	goto L144
L153:
	;
	goto L137
L154:
	;
	v398 = int32(2)
	v399 = base.I32_div_s(v393-v391, v398)
	v400 = v399 + v391
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v389+v400<<(uint(v398)%32))))
	v405 = base.B2i32(v404 == v308)
	if v404 == v308 {
		v420 = v405
		goto L131
	} else {
		goto L156
	}
L155:
	;
	v420 = v405
	goto L131
L156:
	;
	v408 = base.B2i32(base.Ui32(v404) < base.Ui32(v308))
	if base.Ui32(v404) < base.Ui32(v308) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v409 = v400 + int32(1)
	goto L159
L158:
	;
	v409 = v391
	goto L159
L159:
	;
	if base.Ui32(v404) < base.Ui32(v308) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v412 = v393
	goto L162
L161:
	;
	v412 = v400 - int32(1)
	goto L162
L162:
	;
	if v409 <= v412 {
		v391 = v409
		v393 = v412
		goto L154
	} else {
		goto L163
	}
L163:
	;
	goto L155
L164:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v432&int32(32) != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	goto L166
L166:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v738 = F_XidInMVCCSnapshot(m, v737, l1)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L64
	} else {
		goto L269
	}
L167:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui32(v442) <= base.Ui32(v441) {
		v3629 = int32(0)
		goto L1
	} else {
		goto L171
	}
L168:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[4]))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v436+v431<<(uint(int32(3))%32))))
	v441 = v440
	goto L170
L169:
	;
	v441 = v431
	goto L170
L170:
	;
	goto L167
L171:
	;
	v444 = int32(1)
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	if v445&int32(2048)|v445&int32(128)|base.B2i32(v445&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64)) != 0 {
		v3629 = v444
		goto L1
	} else {
		goto L172
	}
L172:
	;
	if v445&int32(_a_F_HeapTupleSatisfiesVisibility_2) != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v458 = F_HeapTupleGetUpdateXid(m, v16)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L64
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if base.Ui32(v596) < base.Ui32(int32(3)) {
		goto L223
	} else {
		goto L224
	}
L176:
	;
	if base.Ui32(v458) < base.Ui32(int32(3)) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	if v579 == int32(0) {
		v3629 = v444
		goto L1
	} else {
		goto L217
	}
L178:
	;
	v579 = int32(0)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v470 == v458 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v579 = int32(1)
	goto L177
L182:
	;
	goto L183
L183:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v474 <= int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v579 = v571
	goto L177
L185:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v478 == int32(0) {
		v571 = int32(0)
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v542 = int32(0)
	v544 = v474 - int32(1)
	goto L207
L188:
	;
	v483 = v478
	goto L189
L189:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v483)+20))
	if v488 == int32(4) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v571 = int32(0)
	goto L184
L191:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v483)+80))
	if v535 != 0 {
		v483 = v535
		goto L189
	} else {
		goto L206
	}
L192:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	if v491 == int32(0) {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v494 = int32(1)
	if v458 == v491 {
		v571 = v494
		goto L184
	} else {
		goto L194
	}
L194:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v483)+52))
	v498 = v496 - int32(1)
	if v498 < int32(0) {
		goto L191
	} else {
		goto L195
	}
L195:
	;
	v503 = int32(0)
	v505 = v498
	goto L196
L196:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v483)+48))
	v511 = int32(2)
	v512 = base.I32_div_s(v505-v503, v511)
	v513 = v512 + v503
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v509+v513<<(uint(v511)%32))))
	if v517 == v458 {
		v571 = v494
		goto L184
	} else {
		goto L198
	}
L197:
	;
	goto L191
L198:
	;
	v521 = F_TransactionIdPrecedes(m, v517, v458)
	mBase = m.M
	if v521 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v522 = v513 + int32(1)
	goto L201
L200:
	;
	v522 = v503
	goto L201
L201:
	;
	if v521 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v525 = v505
	goto L204
L203:
	;
	v525 = v513 - int32(1)
	goto L204
L204:
	;
	if v522 <= v525 {
		v503 = v522
		v505 = v525
		goto L196
	} else {
		goto L205
	}
L205:
	;
	goto L197
L206:
	;
	goto L190
L207:
	;
	v549 = int32(2)
	v550 = base.I32_div_s(v544-v542, v549)
	v551 = v550 + v542
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v540+v551<<(uint(v549)%32))))
	v556 = base.B2i32(v555 == v458)
	if v555 == v458 {
		v571 = v556
		goto L184
	} else {
		goto L209
	}
L208:
	;
	v571 = v556
	goto L184
L209:
	;
	v559 = base.B2i32(base.Ui32(v555) < base.Ui32(v458))
	if base.Ui32(v555) < base.Ui32(v458) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v560 = v551 + int32(1)
	goto L212
L211:
	;
	v560 = v542
	goto L212
L212:
	;
	if base.Ui32(v555) < base.Ui32(v458) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v563 = v544
	goto L215
L214:
	;
	v563 = v551 - int32(1)
	goto L215
L215:
	;
	if v560 <= v563 {
		v542 = v560
		v544 = v563
		goto L207
	} else {
		goto L216
	}
L216:
	;
	goto L208
L217:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v584&int32(32) != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3629 = base.B2i32(base.Ui32(v594) <= base.Ui32(v593))
	goto L1
L219:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[4]))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v588+v583<<(uint(int32(3))%32))+4))
	v593 = v592
	goto L221
L220:
	;
	v593 = v583
	goto L221
L221:
	;
	goto L218
L222:
	;
	if v716 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L223:
	;
	v716 = int32(0)
	goto L222
L224:
	;
	goto L225
L225:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v607 == v596 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v716 = int32(1)
	goto L222
L227:
	;
	goto L228
L228:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v611 <= int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v716 = v708
	goto L222
L230:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v615 == int32(0) {
		v708 = int32(0)
		goto L229
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v679 = int32(0)
	v681 = v611 - int32(1)
	goto L252
L233:
	;
	v620 = v615
	goto L234
L234:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v620)+20))
	if v625 == int32(4) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v708 = int32(0)
	goto L229
L236:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v620)+80))
	if v672 != 0 {
		v620 = v672
		goto L234
	} else {
		goto L251
	}
L237:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	if v628 == int32(0) {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v631 = int32(1)
	if v596 == v628 {
		v708 = v631
		goto L229
	} else {
		goto L239
	}
L239:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v620)+52))
	v635 = v633 - int32(1)
	if v635 < int32(0) {
		goto L236
	} else {
		goto L240
	}
L240:
	;
	v640 = int32(0)
	v642 = v635
	goto L241
L241:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v620)+48))
	v648 = int32(2)
	v649 = base.I32_div_s(v642-v640, v648)
	v650 = v649 + v640
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v646+v650<<(uint(v648)%32))))
	if v654 == v596 {
		v708 = v631
		goto L229
	} else {
		goto L243
	}
L242:
	;
	goto L236
L243:
	;
	v658 = F_TransactionIdPrecedes(m, v654, v596)
	mBase = m.M
	if v658 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v659 = v650 + int32(1)
	goto L246
L245:
	;
	v659 = v640
	goto L246
L246:
	;
	if v658 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v662 = v642
	goto L249
L248:
	;
	v662 = v650 - int32(1)
	goto L249
L249:
	;
	if v659 <= v662 {
		v640 = v659
		v642 = v662
		goto L241
	} else {
		goto L250
	}
L250:
	;
	goto L242
L251:
	;
	goto L235
L252:
	;
	v686 = int32(2)
	v687 = base.I32_div_s(v681-v679, v686)
	v688 = v687 + v679
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v677+v688<<(uint(v686)%32))))
	v693 = base.B2i32(v692 == v596)
	if v692 == v596 {
		v708 = v693
		goto L229
	} else {
		goto L254
	}
L253:
	;
	v708 = v693
	goto L229
L254:
	;
	v696 = base.B2i32(base.Ui32(v692) < base.Ui32(v596))
	if base.Ui32(v692) < base.Ui32(v596) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v697 = v688 + int32(1)
	goto L257
L256:
	;
	v697 = v679
	goto L257
L257:
	;
	if base.Ui32(v692) < base.Ui32(v596) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v700 = v681
	goto L260
L259:
	;
	v700 = v688 - int32(1)
	goto L260
L260:
	;
	if v697 <= v700 {
		v679 = v697
		v681 = v700
		goto L252
	} else {
		goto L261
	}
L261:
	;
	goto L253
L262:
	;
	v719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	v721 = v719 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)) = uint16(v721)
	goto L2
L263:
	;
	goto L264
L264:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v725&int32(32) != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3629 = base.B2i32(base.Ui32(v735) <= base.Ui32(v734))
	goto L1
L266:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[4]))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v729+v724<<(uint(int32(3))%32))+4))
	v734 = v733
	goto L268
L267:
	;
	v734 = v724
	goto L268
L268:
	;
	goto L265
L269:
	;
	if v738 != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v3629 = int32(0)
	goto L1
L271:
	;
	goto L272
L272:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v742 = F_TransactionIdDidCommit(m, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L64
	} else {
		goto L273
	}
L273:
	;
	if v742 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_HeapTupleSetHintBits(m, v16, l2, int32(256), v745)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L64
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	v750 = v748 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)) = uint16(v750)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L64
	} else {
		goto L278
	}
L277:
	;
	goto L10
L278:
	;
	v3629 = int32(0)
	goto L1
L279:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v761 = F_XidInMVCCSnapshot(m, v760, l1)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L64
	} else {
		goto L280
	}
L280:
	;
	if v761 == int32(0) {
		goto L10
	} else {
		goto L281
	}
L281:
	;
	v3629 = int32(0)
	goto L1
L282:
	;
	goto L10
L283:
	;
	if v774&int32(_a_F_HeapTupleSatisfiesVisibility_2) != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v787 = F_HeapTupleGetUpdateXid(m, v16)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L64
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v774&int32(1024) == int32(0) {
		goto L338
	} else {
		goto L339
	}
L287:
	;
	if base.Ui32(v787) < base.Ui32(int32(3)) {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	if v908 != 0 {
		goto L328
	} else {
		goto L329
	}
L289:
	;
	v908 = int32(0)
	goto L288
L290:
	;
	goto L291
L291:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v799 == v787 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v908 = int32(1)
	goto L288
L293:
	;
	goto L294
L294:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v803 <= int32(0) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	v908 = v900
	goto L288
L296:
	;
	v807 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v807 == int32(0) {
		v900 = int32(0)
		goto L295
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v869 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v871 = int32(0)
	v873 = v803 - int32(1)
	goto L318
L299:
	;
	v812 = v807
	goto L300
L300:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v812)+20))
	if v817 == int32(4) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v900 = int32(0)
	goto L295
L302:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v812)+80))
	if v864 != 0 {
		v812 = v864
		goto L300
	} else {
		goto L317
	}
L303:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	if v820 == int32(0) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v823 = int32(1)
	if v787 == v820 {
		v900 = v823
		goto L295
	} else {
		goto L305
	}
L305:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v812)+52))
	v827 = v825 - int32(1)
	if v827 < int32(0) {
		goto L302
	} else {
		goto L306
	}
L306:
	;
	v832 = int32(0)
	v834 = v827
	goto L307
L307:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v812)+48))
	v840 = int32(2)
	v841 = base.I32_div_s(v834-v832, v840)
	v842 = v841 + v832
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v838+v842<<(uint(v840)%32))))
	if v846 == v787 {
		v900 = v823
		goto L295
	} else {
		goto L309
	}
L308:
	;
	goto L302
L309:
	;
	v850 = F_TransactionIdPrecedes(m, v846, v787)
	mBase = m.M
	if v850 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v851 = v842 + int32(1)
	goto L312
L311:
	;
	v851 = v832
	goto L312
L312:
	;
	if v850 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v854 = v834
	goto L315
L314:
	;
	v854 = v842 - int32(1)
	goto L315
L315:
	;
	if v851 <= v854 {
		v832 = v851
		v834 = v854
		goto L307
	} else {
		goto L316
	}
L316:
	;
	goto L308
L317:
	;
	goto L301
L318:
	;
	v878 = int32(2)
	v879 = base.I32_div_s(v873-v871, v878)
	v880 = v879 + v871
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v869+v880<<(uint(v878)%32))))
	v885 = base.B2i32(v884 == v787)
	if v884 == v787 {
		v900 = v885
		goto L295
	} else {
		goto L320
	}
L319:
	;
	v900 = v885
	goto L295
L320:
	;
	v888 = base.B2i32(base.Ui32(v884) < base.Ui32(v787))
	if base.Ui32(v884) < base.Ui32(v787) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v889 = v880 + int32(1)
	goto L323
L322:
	;
	v889 = v871
	goto L323
L323:
	;
	if base.Ui32(v884) < base.Ui32(v787) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v892 = v873
	goto L326
L325:
	;
	v892 = v880 - int32(1)
	goto L326
L326:
	;
	if v889 <= v892 {
		v871 = v889
		v873 = v892
		goto L318
	} else {
		goto L327
	}
L327:
	;
	goto L319
L328:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v911&int32(32) != 0 {
		goto L332
	} else {
		goto L333
	}
L329:
	;
	goto L330
L330:
	;
	v923 = F_XidInMVCCSnapshot(m, v787, l1)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L64
	} else {
		goto L335
	}
L331:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3629 = base.B2i32(base.Ui32(v921) <= base.Ui32(v920))
	goto L1
L332:
	;
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[4]))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v915+v910<<(uint(int32(3))%32))+4))
	v920 = v919
	goto L334
L333:
	;
	v920 = v910
	goto L334
L334:
	;
	goto L331
L335:
	;
	if v923 != 0 {
		v3629 = v773
		goto L1
	} else {
		goto L336
	}
L336:
	;
	v925 = F_TransactionIdDidCommit(m, v787)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L64
	} else {
		goto L337
	}
L337:
	;
	v3629 = v925 ^ int32(1)
	goto L1
L338:
	;
	if base.Ui32(v929) < base.Ui32(int32(3)) {
		goto L342
	} else {
		goto L343
	}
L339:
	;
	goto L340
L340:
	;
	v1085 = F_XidInMVCCSnapshot(m, v929, l1)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L64
	} else {
		goto L395
	}
L341:
	;
	if v1053 != 0 {
		goto L381
	} else {
		goto L382
	}
L342:
	;
	v1053 = int32(0)
	goto L341
L343:
	;
	goto L344
L344:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v944 == v929 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1053 = int32(1)
	goto L341
L346:
	;
	goto L347
L347:
	;
	v948 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v948 <= int32(0) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1053 = v1045
	goto L341
L349:
	;
	v952 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v952 == int32(0) {
		v1045 = int32(0)
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1016 = int32(0)
	v1018 = v948 - int32(1)
	goto L371
L352:
	;
	v957 = v952
	goto L353
L353:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v957)+20))
	if v962 == int32(4) {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1045 = int32(0)
	goto L348
L355:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v957)+80))
	if v1009 != 0 {
		v957 = v1009
		goto L353
	} else {
		goto L370
	}
L356:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v957)))
	if v965 == int32(0) {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v968 = int32(1)
	if v929 == v965 {
		v1045 = v968
		goto L348
	} else {
		goto L358
	}
L358:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v957)+52))
	v972 = v970 - int32(1)
	if v972 < int32(0) {
		goto L355
	} else {
		goto L359
	}
L359:
	;
	v977 = int32(0)
	v979 = v972
	goto L360
L360:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v957)+48))
	v985 = int32(2)
	v986 = base.I32_div_s(v979-v977, v985)
	v987 = v986 + v977
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v983+v987<<(uint(v985)%32))))
	if v991 == v929 {
		v1045 = v968
		goto L348
	} else {
		goto L362
	}
L361:
	;
	goto L355
L362:
	;
	v995 = F_TransactionIdPrecedes(m, v991, v929)
	mBase = m.M
	if v995 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v996 = v987 + int32(1)
	goto L365
L364:
	;
	v996 = v977
	goto L365
L365:
	;
	if v995 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v999 = v979
	goto L368
L367:
	;
	v999 = v987 - int32(1)
	goto L368
L368:
	;
	if v996 <= v999 {
		v977 = v996
		v979 = v999
		goto L360
	} else {
		goto L369
	}
L369:
	;
	goto L361
L370:
	;
	goto L354
L371:
	;
	v1023 = int32(2)
	v1024 = base.I32_div_s(v1018-v1016, v1023)
	v1025 = v1024 + v1016
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1014+v1025<<(uint(v1023)%32))))
	v1030 = base.B2i32(v1029 == v929)
	if v1029 == v929 {
		v1045 = v1030
		goto L348
	} else {
		goto L373
	}
L372:
	;
	v1045 = v1030
	goto L348
L373:
	;
	v1033 = base.B2i32(base.Ui32(v1029) < base.Ui32(v929))
	if base.Ui32(v1029) < base.Ui32(v929) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1034 = v1025 + int32(1)
	goto L376
L375:
	;
	v1034 = v1016
	goto L376
L376:
	;
	if base.Ui32(v1029) < base.Ui32(v929) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1037 = v1018
	goto L379
L378:
	;
	v1037 = v1025 - int32(1)
	goto L379
L379:
	;
	if v1034 <= v1037 {
		v1016 = v1034
		v1018 = v1037
		goto L371
	} else {
		goto L380
	}
L380:
	;
	goto L372
L381:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v1056&int32(32) != 0 {
		goto L385
	} else {
		goto L386
	}
L382:
	;
	goto L383
L383:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v1069 = F_XidInMVCCSnapshot(m, v1068, l1)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L64
	} else {
		goto L388
	}
L384:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3629 = base.B2i32(base.Ui32(v1066) <= base.Ui32(v1065))
	goto L1
L385:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[4]))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1060+v1055<<(uint(int32(3))%32))+4))
	v1065 = v1064
	goto L387
L386:
	;
	v1065 = v1055
	goto L387
L387:
	;
	goto L384
L388:
	;
	if v1069 != 0 {
		v3629 = v773
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v1072 = F_TransactionIdDidCommit(m, v1071)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L64
	} else {
		goto L390
	}
L390:
	;
	if v1072 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)))
	v1078 = v1076 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+20)) = uint16(v1078)
	goto L2
L392:
	;
	goto L393
L393:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_HeapTupleSetHintBits(m, v16, l2, int32(1024), v1081)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L64
	} else {
		goto L394
	}
L394:
	;
	v3629 = int32(0)
	goto L1
L395:
	;
	if v1085 != 0 {
		v3629 = v773
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v3629 = int32(0)
	goto L1
L397:
	;
	v1782 = int32(1)
	v1783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	if v1783&int32(2048) != 0 {
		v3629 = v1782
		goto L1
	} else {
		goto L648
	}
L398:
	;
	v1092 = base.I32_extend16_s(v1089)
	if v1092&int32(512) != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v3629 = v4
	goto L1
L400:
	;
	goto L401
L401:
	;
	if v1092&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L403
	} else {
		goto L404
	}
L402:
	;
	v1776 = v1222 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v1776)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L64
	} else {
		goto L647
	}
L403:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+8))
	if base.Ui32(v1097) < base.Ui32(int32(3)) {
		goto L407
	} else {
		goto L408
	}
L404:
	;
	goto L405
L405:
	;
	if v1092 < int32(0) {
		goto L454
	} else {
		goto L455
	}
L406:
	;
	if v1217 != 0 {
		goto L446
	} else {
		goto L447
	}
L407:
	;
	v1217 = int32(0)
	goto L406
L408:
	;
	goto L409
L409:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1108 == v1097 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1217 = int32(1)
	goto L406
L411:
	;
	goto L412
L412:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1112 <= int32(0) {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	v1217 = v1209
	goto L406
L414:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1116 == int32(0) {
		v1209 = int32(0)
		goto L413
	} else {
		goto L417
	}
L415:
	;
	goto L416
L416:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1180 = int32(0)
	v1182 = v1112 - int32(1)
	goto L436
L417:
	;
	v1121 = v1116
	goto L418
L418:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+20))
	if v1126 == int32(4) {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v1209 = int32(0)
	goto L413
L420:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+80))
	if v1173 != 0 {
		v1121 = v1173
		goto L418
	} else {
		goto L435
	}
L421:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	if v1129 == int32(0) {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	v1132 = int32(1)
	if v1097 == v1129 {
		v1209 = v1132
		goto L413
	} else {
		goto L423
	}
L423:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+52))
	v1136 = v1134 - int32(1)
	if v1136 < int32(0) {
		goto L420
	} else {
		goto L424
	}
L424:
	;
	v1141 = int32(0)
	v1143 = v1136
	goto L425
L425:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+48))
	v1149 = int32(2)
	v1150 = base.I32_div_s(v1143-v1141, v1149)
	v1151 = v1150 + v1141
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1147+v1151<<(uint(v1149)%32))))
	if v1155 == v1097 {
		v1209 = v1132
		goto L413
	} else {
		goto L427
	}
L426:
	;
	goto L420
L427:
	;
	v1159 = F_TransactionIdPrecedes(m, v1155, v1097)
	mBase = m.M
	if v1159 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1160 = v1151 + int32(1)
	goto L430
L429:
	;
	v1160 = v1141
	goto L430
L430:
	;
	if v1159 != 0 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1163 = v1143
	goto L433
L432:
	;
	v1163 = v1151 - int32(1)
	goto L433
L433:
	;
	if v1160 <= v1163 {
		v1141 = v1160
		v1143 = v1163
		goto L425
	} else {
		goto L434
	}
L434:
	;
	goto L426
L435:
	;
	goto L419
L436:
	;
	v1187 = int32(2)
	v1188 = base.I32_div_s(v1182-v1180, v1187)
	v1189 = v1188 + v1180
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1178+v1189<<(uint(v1187)%32))))
	v1194 = base.B2i32(v1193 == v1097)
	if v1193 == v1097 {
		v1209 = v1194
		goto L413
	} else {
		goto L438
	}
L437:
	;
	v1209 = v1194
	goto L413
L438:
	;
	v1197 = base.B2i32(base.Ui32(v1193) < base.Ui32(v1097))
	if base.Ui32(v1193) < base.Ui32(v1097) {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v1198 = v1189 + int32(1)
	goto L441
L440:
	;
	v1198 = v1180
	goto L441
L441:
	;
	if base.Ui32(v1193) < base.Ui32(v1097) {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1201 = v1182
	goto L444
L443:
	;
	v1201 = v1189 - int32(1)
	goto L444
L444:
	;
	if v1198 <= v1201 {
		v1180 = v1198
		v1182 = v1201
		goto L436
	} else {
		goto L445
	}
L445:
	;
	goto L437
L446:
	;
	v3629 = v4
	goto L1
L447:
	;
	goto L448
L448:
	;
	v1218 = F_TransactionIdIsInProgress(m, v1097)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L64
	} else {
		goto L449
	}
L449:
	;
	if v1218 != 0 {
		goto L397
	} else {
		goto L450
	}
L450:
	;
	v1220 = F_TransactionIdDidCommit(m, v1097)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L64
	} else {
		goto L451
	}
L451:
	;
	v1222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	if v1220 == int32(0) {
		goto L402
	} else {
		goto L452
	}
L452:
	;
	v1226 = v1222 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v1226)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L64
	} else {
		goto L453
	}
L453:
	;
	v3629 = int32(0)
	goto L1
L454:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+8))
	if base.Ui32(v1234) < base.Ui32(int32(3)) {
		goto L458
	} else {
		goto L459
	}
L455:
	;
	goto L456
L456:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	if base.Ui32(v1373) < base.Ui32(int32(3)) {
		goto L509
	} else {
		goto L510
	}
L457:
	;
	if v1354 != 0 {
		goto L397
	} else {
		goto L497
	}
L458:
	;
	v1354 = int32(0)
	goto L457
L459:
	;
	goto L460
L460:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1245 == v1234 {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v1354 = int32(1)
	goto L457
L462:
	;
	goto L463
L463:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1249 <= int32(0) {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	v1354 = v1346
	goto L457
L465:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1253 == int32(0) {
		v1346 = int32(0)
		goto L464
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1317 = int32(0)
	v1319 = v1249 - int32(1)
	goto L487
L468:
	;
	v1258 = v1253
	goto L469
L469:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+20))
	if v1263 == int32(4) {
		goto L471
	} else {
		goto L472
	}
L470:
	;
	v1346 = int32(0)
	goto L464
L471:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+80))
	if v1310 != 0 {
		v1258 = v1310
		goto L469
	} else {
		goto L486
	}
L472:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1258)))
	if v1266 == int32(0) {
		goto L471
	} else {
		goto L473
	}
L473:
	;
	v1269 = int32(1)
	if v1234 == v1266 {
		v1346 = v1269
		goto L464
	} else {
		goto L474
	}
L474:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+52))
	v1273 = v1271 - int32(1)
	if v1273 < int32(0) {
		goto L471
	} else {
		goto L475
	}
L475:
	;
	v1278 = int32(0)
	v1280 = v1273
	goto L476
L476:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+48))
	v1286 = int32(2)
	v1287 = base.I32_div_s(v1280-v1278, v1286)
	v1288 = v1287 + v1278
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1284+v1288<<(uint(v1286)%32))))
	if v1292 == v1234 {
		v1346 = v1269
		goto L464
	} else {
		goto L478
	}
L477:
	;
	goto L471
L478:
	;
	v1296 = F_TransactionIdPrecedes(m, v1292, v1234)
	mBase = m.M
	if v1296 != 0 {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v1297 = v1288 + int32(1)
	goto L481
L480:
	;
	v1297 = v1278
	goto L481
L481:
	;
	if v1296 != 0 {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v1300 = v1280
	goto L484
L483:
	;
	v1300 = v1288 - int32(1)
	goto L484
L484:
	;
	if v1297 <= v1300 {
		v1278 = v1297
		v1280 = v1300
		goto L476
	} else {
		goto L485
	}
L485:
	;
	goto L477
L486:
	;
	goto L470
L487:
	;
	v1324 = int32(2)
	v1325 = base.I32_div_s(v1319-v1317, v1324)
	v1326 = v1325 + v1317
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1315+v1326<<(uint(v1324)%32))))
	v1331 = base.B2i32(v1330 == v1234)
	if v1330 == v1234 {
		v1346 = v1331
		goto L464
	} else {
		goto L489
	}
L488:
	;
	v1346 = v1331
	goto L464
L489:
	;
	v1334 = base.B2i32(base.Ui32(v1330) < base.Ui32(v1234))
	if base.Ui32(v1330) < base.Ui32(v1234) {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v1335 = v1326 + int32(1)
	goto L492
L491:
	;
	v1335 = v1317
	goto L492
L492:
	;
	if base.Ui32(v1330) < base.Ui32(v1234) {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v1338 = v1319
	goto L495
L494:
	;
	v1338 = v1326 - int32(1)
	goto L495
L495:
	;
	if v1335 <= v1338 {
		v1317 = v1335
		v1319 = v1338
		goto L487
	} else {
		goto L496
	}
L496:
	;
	goto L488
L497:
	;
	v1355 = F_TransactionIdIsInProgress(m, v1234)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L64
	} else {
		goto L498
	}
L498:
	;
	if v1355 != 0 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v3629 = v4
	goto L1
L500:
	;
	goto L501
L501:
	;
	v1357 = F_TransactionIdDidCommit(m, v1234)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L64
	} else {
		goto L502
	}
L502:
	;
	v1359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	if v1357 != 0 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v1361 = v1359 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v1361)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L64
	} else {
		goto L506
	}
L504:
	;
	goto L505
L505:
	;
	v1367 = v1359 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v1367)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L64
	} else {
		goto L507
	}
L506:
	;
	goto L397
L507:
	;
	v3629 = int32(0)
	goto L1
L508:
	;
	if v1493 != 0 {
		goto L548
	} else {
		goto L549
	}
L509:
	;
	v1493 = int32(0)
	goto L508
L510:
	;
	goto L511
L511:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1384 == v1373 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v1493 = int32(1)
	goto L508
L513:
	;
	goto L514
L514:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1388 <= int32(0) {
		goto L516
	} else {
		goto L517
	}
L515:
	;
	v1493 = v1485
	goto L508
L516:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1392 == int32(0) {
		v1485 = int32(0)
		goto L515
	} else {
		goto L519
	}
L517:
	;
	goto L518
L518:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1456 = int32(0)
	v1458 = v1388 - int32(1)
	goto L538
L519:
	;
	v1397 = v1392
	goto L520
L520:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+20))
	if v1402 == int32(4) {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	v1485 = int32(0)
	goto L515
L522:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+80))
	if v1449 != 0 {
		v1397 = v1449
		goto L520
	} else {
		goto L537
	}
L523:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1397)))
	if v1405 == int32(0) {
		goto L522
	} else {
		goto L524
	}
L524:
	;
	v1408 = int32(1)
	if v1373 == v1405 {
		v1485 = v1408
		goto L515
	} else {
		goto L525
	}
L525:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+52))
	v1412 = v1410 - int32(1)
	if v1412 < int32(0) {
		goto L522
	} else {
		goto L526
	}
L526:
	;
	v1417 = int32(0)
	v1419 = v1412
	goto L527
L527:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+48))
	v1425 = int32(2)
	v1426 = base.I32_div_s(v1419-v1417, v1425)
	v1427 = v1426 + v1417
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1423+v1427<<(uint(v1425)%32))))
	if v1431 == v1373 {
		v1485 = v1408
		goto L515
	} else {
		goto L529
	}
L528:
	;
	goto L522
L529:
	;
	v1435 = F_TransactionIdPrecedes(m, v1431, v1373)
	mBase = m.M
	if v1435 != 0 {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v1436 = v1427 + int32(1)
	goto L532
L531:
	;
	v1436 = v1417
	goto L532
L532:
	;
	if v1435 != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v1439 = v1419
	goto L535
L534:
	;
	v1439 = v1427 - int32(1)
	goto L535
L535:
	;
	if v1436 <= v1439 {
		v1417 = v1436
		v1419 = v1439
		goto L527
	} else {
		goto L536
	}
L536:
	;
	goto L528
L537:
	;
	goto L521
L538:
	;
	v1463 = int32(2)
	v1464 = base.I32_div_s(v1458-v1456, v1463)
	v1465 = v1464 + v1456
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1454+v1465<<(uint(v1463)%32))))
	v1470 = base.B2i32(v1469 == v1373)
	if v1469 == v1373 {
		v1485 = v1470
		goto L515
	} else {
		goto L540
	}
L539:
	;
	v1485 = v1470
	goto L515
L540:
	;
	v1473 = base.B2i32(base.Ui32(v1469) < base.Ui32(v1373))
	if base.Ui32(v1469) < base.Ui32(v1373) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v1474 = v1465 + int32(1)
	goto L543
L542:
	;
	v1474 = v1456
	goto L543
L543:
	;
	if base.Ui32(v1469) < base.Ui32(v1373) {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v1477 = v1458
	goto L546
L545:
	;
	v1477 = v1465 - int32(1)
	goto L546
L546:
	;
	if v1474 <= v1477 {
		v1456 = v1474
		v1458 = v1477
		goto L538
	} else {
		goto L547
	}
L547:
	;
	goto L539
L548:
	;
	v1495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	if v1495&int32(2048)|v1495&int32(128)|base.B2i32(v1495&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64)) != 0 {
		v3629 = int32(1)
		goto L1
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	v1759 = F_TransactionIdIsInProgress(m, v1758)
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L64
	} else {
		goto L637
	}
L551:
	;
	if v1495&int32(_a_F_HeapTupleSatisfiesVisibility_2) != 0 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v1508 = F_HeapTupleGetUpdateXid(m, v1088)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L64
	} else {
		goto L555
	}
L553:
	;
	goto L554
L554:
	;
	v1632 = int32(0)
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	if base.Ui32(v1633) < base.Ui32(int32(3)) {
		goto L597
	} else {
		goto L598
	}
L555:
	;
	if base.Ui32(v1508) < base.Ui32(int32(3)) {
		goto L557
	} else {
		goto L558
	}
L556:
	;
	v3629 = v1629 ^ int32(1)
	goto L1
L557:
	;
	v1629 = int32(0)
	goto L556
L558:
	;
	goto L559
L559:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1520 == v1508 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v1629 = int32(1)
	goto L556
L561:
	;
	goto L562
L562:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1524 <= int32(0) {
		goto L564
	} else {
		goto L565
	}
L563:
	;
	v1629 = v1621
	goto L556
L564:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1528 == int32(0) {
		v1621 = int32(0)
		goto L563
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1592 = int32(0)
	v1594 = v1524 - int32(1)
	goto L586
L567:
	;
	v1533 = v1528
	goto L568
L568:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+20))
	if v1538 == int32(4) {
		goto L570
	} else {
		goto L571
	}
L569:
	;
	v1621 = int32(0)
	goto L563
L570:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+80))
	if v1585 != 0 {
		v1533 = v1585
		goto L568
	} else {
		goto L585
	}
L571:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1533)))
	if v1541 == int32(0) {
		goto L570
	} else {
		goto L572
	}
L572:
	;
	v1544 = int32(1)
	if v1508 == v1541 {
		v1621 = v1544
		goto L563
	} else {
		goto L573
	}
L573:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+52))
	v1548 = v1546 - int32(1)
	if v1548 < int32(0) {
		goto L570
	} else {
		goto L574
	}
L574:
	;
	v1553 = int32(0)
	v1555 = v1548
	goto L575
L575:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+48))
	v1561 = int32(2)
	v1562 = base.I32_div_s(v1555-v1553, v1561)
	v1563 = v1562 + v1553
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1559+v1563<<(uint(v1561)%32))))
	if v1567 == v1508 {
		v1621 = v1544
		goto L563
	} else {
		goto L577
	}
L576:
	;
	goto L570
L577:
	;
	v1571 = F_TransactionIdPrecedes(m, v1567, v1508)
	mBase = m.M
	if v1571 != 0 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v1572 = v1563 + int32(1)
	goto L580
L579:
	;
	v1572 = v1553
	goto L580
L580:
	;
	if v1571 != 0 {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v1575 = v1555
	goto L583
L582:
	;
	v1575 = v1563 - int32(1)
	goto L583
L583:
	;
	if v1572 <= v1575 {
		v1553 = v1572
		v1555 = v1575
		goto L575
	} else {
		goto L584
	}
L584:
	;
	goto L576
L585:
	;
	goto L569
L586:
	;
	v1599 = int32(2)
	v1600 = base.I32_div_s(v1594-v1592, v1599)
	v1601 = v1600 + v1592
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1590+v1601<<(uint(v1599)%32))))
	v1606 = base.B2i32(v1605 == v1508)
	if v1605 == v1508 {
		v1621 = v1606
		goto L563
	} else {
		goto L588
	}
L587:
	;
	v1621 = v1606
	goto L563
L588:
	;
	v1609 = base.B2i32(base.Ui32(v1605) < base.Ui32(v1508))
	if base.Ui32(v1605) < base.Ui32(v1508) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v1610 = v1601 + int32(1)
	goto L591
L590:
	;
	v1610 = v1592
	goto L591
L591:
	;
	if base.Ui32(v1605) < base.Ui32(v1508) {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v1613 = v1594
	goto L594
L593:
	;
	v1613 = v1601 - int32(1)
	goto L594
L594:
	;
	if v1610 <= v1613 {
		v1592 = v1610
		v1594 = v1613
		goto L586
	} else {
		goto L595
	}
L595:
	;
	goto L587
L596:
	;
	if v1753 != 0 {
		v3629 = v1632
		goto L1
	} else {
		goto L636
	}
L597:
	;
	v1753 = int32(0)
	goto L596
L598:
	;
	goto L599
L599:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1644 == v1633 {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v1753 = int32(1)
	goto L596
L601:
	;
	goto L602
L602:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1648 <= int32(0) {
		goto L604
	} else {
		goto L605
	}
L603:
	;
	v1753 = v1745
	goto L596
L604:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1652 == int32(0) {
		v1745 = v1632
		goto L603
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1716 = int32(0)
	v1718 = v1648 - int32(1)
	goto L626
L607:
	;
	v1657 = v1652
	goto L608
L608:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+20))
	if v1662 == int32(4) {
		goto L610
	} else {
		goto L611
	}
L609:
	;
	v1745 = int32(0)
	goto L603
L610:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+80))
	if v1709 != 0 {
		v1657 = v1709
		goto L608
	} else {
		goto L625
	}
L611:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1657)))
	if v1665 == int32(0) {
		goto L610
	} else {
		goto L612
	}
L612:
	;
	v1668 = int32(1)
	if v1633 == v1665 {
		v1745 = v1668
		goto L603
	} else {
		goto L613
	}
L613:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+52))
	v1672 = v1670 - int32(1)
	if v1672 < int32(0) {
		goto L610
	} else {
		goto L614
	}
L614:
	;
	v1677 = int32(0)
	v1679 = v1672
	goto L615
L615:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+48))
	v1685 = int32(2)
	v1686 = base.I32_div_s(v1679-v1677, v1685)
	v1687 = v1686 + v1677
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1683+v1687<<(uint(v1685)%32))))
	if v1691 == v1633 {
		v1745 = v1668
		goto L603
	} else {
		goto L617
	}
L616:
	;
	goto L610
L617:
	;
	v1695 = F_TransactionIdPrecedes(m, v1691, v1633)
	mBase = m.M
	if v1695 != 0 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v1696 = v1687 + int32(1)
	goto L620
L619:
	;
	v1696 = v1677
	goto L620
L620:
	;
	if v1695 != 0 {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v1699 = v1679
	goto L623
L622:
	;
	v1699 = v1687 - int32(1)
	goto L623
L623:
	;
	if v1696 <= v1699 {
		v1677 = v1696
		v1679 = v1699
		goto L615
	} else {
		goto L624
	}
L624:
	;
	goto L616
L625:
	;
	goto L609
L626:
	;
	v1723 = int32(2)
	v1724 = base.I32_div_s(v1718-v1716, v1723)
	v1725 = v1724 + v1716
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1714+v1725<<(uint(v1723)%32))))
	v1730 = base.B2i32(v1729 == v1633)
	if v1729 == v1633 {
		v1745 = v1730
		goto L603
	} else {
		goto L628
	}
L627:
	;
	v1745 = v1730
	goto L603
L628:
	;
	v1733 = base.B2i32(base.Ui32(v1729) < base.Ui32(v1633))
	if base.Ui32(v1729) < base.Ui32(v1633) {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	v1734 = v1725 + int32(1)
	goto L631
L630:
	;
	v1734 = v1716
	goto L631
L631:
	;
	if base.Ui32(v1729) < base.Ui32(v1633) {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v1737 = v1718
	goto L634
L633:
	;
	v1737 = v1725 - int32(1)
	goto L634
L634:
	;
	if v1734 <= v1737 {
		v1716 = v1734
		v1718 = v1737
		goto L626
	} else {
		goto L635
	}
L635:
	;
	goto L627
L636:
	;
	v1754 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	v1756 = v1754 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v1756)
	goto L2
L637:
	;
	if v1759 != 0 {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v3629 = v4
	goto L1
L639:
	;
	goto L640
L640:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	v1762 = F_TransactionIdDidCommit(m, v1761)
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L64
	} else {
		goto L641
	}
L641:
	;
	if v1762 != 0 {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	F_HeapTupleSetHintBits(m, v1088, l2, int32(256), v1765)
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L64
	} else {
		goto L645
	}
L643:
	;
	goto L644
L644:
	;
	v1768 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	v1770 = v1768 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v1770)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L64
	} else {
		goto L646
	}
L645:
	;
	goto L397
L646:
	;
	v3629 = v4
	goto L1
L647:
	;
	goto L397
L648:
	;
	if v1783&int32(1024) != 0 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	v3629 = int32(base.Ui32(v1783&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v1783&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L650:
	;
	goto L651
L651:
	;
	if v1783&int32(_a_F_HeapTupleSatisfiesVisibility_2) != 0 {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	if v1783&int32(128) != 0 {
		v3629 = v1782
		goto L1
	} else {
		goto L655
	}
L653:
	;
	goto L654
L654:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	if base.Ui32(v1931) < base.Ui32(int32(3)) {
		goto L702
	} else {
		goto L703
	}
L655:
	;
	v1802 = F_HeapTupleGetUpdateXid(m, v1088)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L64
	} else {
		goto L656
	}
L656:
	;
	if base.Ui32(v1802) < base.Ui32(int32(3)) {
		goto L658
	} else {
		goto L659
	}
L657:
	;
	if v1923 != 0 {
		v3629 = int32(0)
		goto L1
	} else {
		goto L697
	}
L658:
	;
	v1923 = int32(0)
	goto L657
L659:
	;
	goto L660
L660:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1814 == v1802 {
		goto L661
	} else {
		goto L662
	}
L661:
	;
	v1923 = int32(1)
	goto L657
L662:
	;
	goto L663
L663:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1818 <= int32(0) {
		goto L665
	} else {
		goto L666
	}
L664:
	;
	v1923 = v1915
	goto L657
L665:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1822 == int32(0) {
		v1915 = int32(0)
		goto L664
	} else {
		goto L668
	}
L666:
	;
	goto L667
L667:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1886 = int32(0)
	v1888 = v1818 - int32(1)
	goto L687
L668:
	;
	v1827 = v1822
	goto L669
L669:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+20))
	if v1832 == int32(4) {
		goto L671
	} else {
		goto L672
	}
L670:
	;
	v1915 = int32(0)
	goto L664
L671:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+80))
	if v1879 != 0 {
		v1827 = v1879
		goto L669
	} else {
		goto L686
	}
L672:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1827)))
	if v1835 == int32(0) {
		goto L671
	} else {
		goto L673
	}
L673:
	;
	v1838 = int32(1)
	if v1802 == v1835 {
		v1915 = v1838
		goto L664
	} else {
		goto L674
	}
L674:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+52))
	v1842 = v1840 - int32(1)
	if v1842 < int32(0) {
		goto L671
	} else {
		goto L675
	}
L675:
	;
	v1847 = int32(0)
	v1849 = v1842
	goto L676
L676:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+48))
	v1855 = int32(2)
	v1856 = base.I32_div_s(v1849-v1847, v1855)
	v1857 = v1856 + v1847
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v1853+v1857<<(uint(v1855)%32))))
	if v1861 == v1802 {
		v1915 = v1838
		goto L664
	} else {
		goto L678
	}
L677:
	;
	goto L671
L678:
	;
	v1865 = F_TransactionIdPrecedes(m, v1861, v1802)
	mBase = m.M
	if v1865 != 0 {
		goto L679
	} else {
		goto L680
	}
L679:
	;
	v1866 = v1857 + int32(1)
	goto L681
L680:
	;
	v1866 = v1847
	goto L681
L681:
	;
	if v1865 != 0 {
		goto L682
	} else {
		goto L683
	}
L682:
	;
	v1869 = v1849
	goto L684
L683:
	;
	v1869 = v1857 - int32(1)
	goto L684
L684:
	;
	if v1866 <= v1869 {
		v1847 = v1866
		v1849 = v1869
		goto L676
	} else {
		goto L685
	}
L685:
	;
	goto L677
L686:
	;
	goto L670
L687:
	;
	v1893 = int32(2)
	v1894 = base.I32_div_s(v1888-v1886, v1893)
	v1895 = v1894 + v1886
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1884+v1895<<(uint(v1893)%32))))
	v1900 = base.B2i32(v1899 == v1802)
	if v1899 == v1802 {
		v1915 = v1900
		goto L664
	} else {
		goto L689
	}
L688:
	;
	v1915 = v1900
	goto L664
L689:
	;
	v1903 = base.B2i32(base.Ui32(v1899) < base.Ui32(v1802))
	if base.Ui32(v1899) < base.Ui32(v1802) {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v1904 = v1895 + int32(1)
	goto L692
L691:
	;
	v1904 = v1886
	goto L692
L692:
	;
	if base.Ui32(v1899) < base.Ui32(v1802) {
		goto L693
	} else {
		goto L694
	}
L693:
	;
	v1907 = v1888
	goto L695
L694:
	;
	v1907 = v1895 - int32(1)
	goto L695
L695:
	;
	if v1904 <= v1907 {
		v1886 = v1904
		v1888 = v1907
		goto L687
	} else {
		goto L696
	}
L696:
	;
	goto L688
L697:
	;
	v1925 = F_TransactionIdIsInProgress(m, v1802)
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L64
	} else {
		goto L698
	}
L698:
	;
	if v1925 != 0 {
		v3629 = int32(1)
		goto L1
	} else {
		goto L699
	}
L699:
	;
	v1927 = F_TransactionIdDidCommit(m, v1802)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L64
	} else {
		goto L700
	}
L700:
	;
	v3629 = v1927 ^ int32(1)
	goto L1
L701:
	;
	if v2051 != 0 {
		goto L741
	} else {
		goto L742
	}
L702:
	;
	v2051 = int32(0)
	goto L701
L703:
	;
	goto L704
L704:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1942 == v1931 {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v2051 = int32(1)
	goto L701
L706:
	;
	goto L707
L707:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1946 <= int32(0) {
		goto L709
	} else {
		goto L710
	}
L708:
	;
	v2051 = v2043
	goto L701
L709:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1950 == int32(0) {
		v2043 = int32(0)
		goto L708
	} else {
		goto L712
	}
L710:
	;
	goto L711
L711:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v2014 = int32(0)
	v2016 = v1946 - int32(1)
	goto L731
L712:
	;
	v1955 = v1950
	goto L713
L713:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+20))
	if v1960 == int32(4) {
		goto L715
	} else {
		goto L716
	}
L714:
	;
	v2043 = int32(0)
	goto L708
L715:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+80))
	if v2007 != 0 {
		v1955 = v2007
		goto L713
	} else {
		goto L730
	}
L716:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1955)))
	if v1963 == int32(0) {
		goto L715
	} else {
		goto L717
	}
L717:
	;
	v1966 = int32(1)
	if v1931 == v1963 {
		v2043 = v1966
		goto L708
	} else {
		goto L718
	}
L718:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+52))
	v1970 = v1968 - int32(1)
	if v1970 < int32(0) {
		goto L715
	} else {
		goto L719
	}
L719:
	;
	v1975 = int32(0)
	v1977 = v1970
	goto L720
L720:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+48))
	v1983 = int32(2)
	v1984 = base.I32_div_s(v1977-v1975, v1983)
	v1985 = v1984 + v1975
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1981+v1985<<(uint(v1983)%32))))
	if v1989 == v1931 {
		v2043 = v1966
		goto L708
	} else {
		goto L722
	}
L721:
	;
	goto L715
L722:
	;
	v1993 = F_TransactionIdPrecedes(m, v1989, v1931)
	mBase = m.M
	if v1993 != 0 {
		goto L723
	} else {
		goto L724
	}
L723:
	;
	v1994 = v1985 + int32(1)
	goto L725
L724:
	;
	v1994 = v1975
	goto L725
L725:
	;
	if v1993 != 0 {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	v1997 = v1977
	goto L728
L727:
	;
	v1997 = v1985 - int32(1)
	goto L728
L728:
	;
	if v1994 <= v1997 {
		v1975 = v1994
		v1977 = v1997
		goto L720
	} else {
		goto L729
	}
L729:
	;
	goto L721
L730:
	;
	goto L714
L731:
	;
	v2021 = int32(2)
	v2022 = base.I32_div_s(v2016-v2014, v2021)
	v2023 = v2022 + v2014
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v2012+v2023<<(uint(v2021)%32))))
	v2028 = base.B2i32(v2027 == v1931)
	if v2027 == v1931 {
		v2043 = v2028
		goto L708
	} else {
		goto L733
	}
L732:
	;
	v2043 = v2028
	goto L708
L733:
	;
	v2031 = base.B2i32(base.Ui32(v2027) < base.Ui32(v1931))
	if base.Ui32(v2027) < base.Ui32(v1931) {
		goto L734
	} else {
		goto L735
	}
L734:
	;
	v2032 = v2023 + int32(1)
	goto L736
L735:
	;
	v2032 = v2014
	goto L736
L736:
	;
	if base.Ui32(v2027) < base.Ui32(v1931) {
		goto L737
	} else {
		goto L738
	}
L737:
	;
	v2035 = v2016
	goto L739
L738:
	;
	v2035 = v2023 - int32(1)
	goto L739
L739:
	;
	if v2032 <= v2035 {
		v2014 = v2032
		v2016 = v2035
		goto L731
	} else {
		goto L740
	}
L740:
	;
	goto L732
L741:
	;
	v2052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	v3629 = int32(base.Ui32(v2052&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v2052&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L742:
	;
	goto L743
L743:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	v2063 = F_TransactionIdIsInProgress(m, v2062)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L64
	} else {
		goto L744
	}
L744:
	;
	if v2063 != 0 {
		v3629 = v1782
		goto L1
	} else {
		goto L745
	}
L745:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	v2066 = F_TransactionIdDidCommit(m, v2065)
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L64
	} else {
		goto L746
	}
L746:
	;
	v2068 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)))
	if v2066 == int32(0) {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	v2072 = v2068 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v2072)
	goto L2
L748:
	;
	goto L749
L749:
	;
	v2076 = int32(0)
	if base.B2i32(v2068&int32(128) == v2076)&base.B2i32(v2068&int32(_a_F_HeapTupleSatisfiesVisibility_0) != int32(64)) == v2076 {
		goto L750
	} else {
		goto L751
	}
L750:
	;
	v2086 = v2068 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v1088)+20)) = uint16(v2086)
	goto L2
L751:
	;
	goto L752
L752:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	F_HeapTupleSetHintBits(m, v1088, l2, int32(1024), v2089)
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L64
	} else {
		goto L753
	}
L753:
	;
	v3629 = int32(0)
	goto L1
L754:
	;
	v3629 = int32(1)
	goto L1
L755:
	;
	v2098 = base.I32_extend16_s(v2095)
	if v2098&int32(512) != 0 {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v3629 = v4
	goto L1
L757:
	;
	goto L758
L758:
	;
	if v2098&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L760
	} else {
		goto L761
	}
L759:
	;
	v2377 = v2375 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2094)+20)) = uint16(v2377)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L64
	} else {
		goto L863
	}
L760:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+8))
	if base.Ui32(v2103) < base.Ui32(int32(3)) {
		goto L764
	} else {
		goto L765
	}
L761:
	;
	goto L762
L762:
	;
	if v2098 < int32(0) {
		goto L811
	} else {
		goto L812
	}
L763:
	;
	if v2223 != 0 {
		goto L803
	} else {
		goto L804
	}
L764:
	;
	v2223 = int32(0)
	goto L763
L765:
	;
	goto L766
L766:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v2114 == v2103 {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	v2223 = int32(1)
	goto L763
L768:
	;
	goto L769
L769:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v2118 <= int32(0) {
		goto L771
	} else {
		goto L772
	}
L770:
	;
	v2223 = v2215
	goto L763
L771:
	;
	v2122 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v2122 == int32(0) {
		v2215 = int32(0)
		goto L770
	} else {
		goto L774
	}
L772:
	;
	goto L773
L773:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v2186 = int32(0)
	v2188 = v2118 - int32(1)
	goto L793
L774:
	;
	v2127 = v2122
	goto L775
L775:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+20))
	if v2132 == int32(4) {
		goto L777
	} else {
		goto L778
	}
L776:
	;
	v2215 = int32(0)
	goto L770
L777:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+80))
	if v2179 != 0 {
		v2127 = v2179
		goto L775
	} else {
		goto L792
	}
L778:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2127)))
	if v2135 == int32(0) {
		goto L777
	} else {
		goto L779
	}
L779:
	;
	v2138 = int32(1)
	if v2103 == v2135 {
		v2215 = v2138
		goto L770
	} else {
		goto L780
	}
L780:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+52))
	v2142 = v2140 - int32(1)
	if v2142 < int32(0) {
		goto L777
	} else {
		goto L781
	}
L781:
	;
	v2147 = int32(0)
	v2149 = v2142
	goto L782
L782:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+48))
	v2155 = int32(2)
	v2156 = base.I32_div_s(v2149-v2147, v2155)
	v2157 = v2156 + v2147
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2153+v2157<<(uint(v2155)%32))))
	if v2161 == v2103 {
		v2215 = v2138
		goto L770
	} else {
		goto L784
	}
L783:
	;
	goto L777
L784:
	;
	v2165 = F_TransactionIdPrecedes(m, v2161, v2103)
	mBase = m.M
	if v2165 != 0 {
		goto L785
	} else {
		goto L786
	}
L785:
	;
	v2166 = v2157 + int32(1)
	goto L787
L786:
	;
	v2166 = v2147
	goto L787
L787:
	;
	if v2165 != 0 {
		goto L788
	} else {
		goto L789
	}
L788:
	;
	v2169 = v2149
	goto L790
L789:
	;
	v2169 = v2157 - int32(1)
	goto L790
L790:
	;
	if v2166 <= v2169 {
		v2147 = v2166
		v2149 = v2169
		goto L782
	} else {
		goto L791
	}
L791:
	;
	goto L783
L792:
	;
	goto L776
L793:
	;
	v2193 = int32(2)
	v2194 = base.I32_div_s(v2188-v2186, v2193)
	v2195 = v2194 + v2186
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2184+v2195<<(uint(v2193)%32))))
	v2200 = base.B2i32(v2199 == v2103)
	if v2199 == v2103 {
		v2215 = v2200
		goto L770
	} else {
		goto L795
	}
L794:
	;
	v2215 = v2200
	goto L770
L795:
	;
	v2203 = base.B2i32(base.Ui32(v2199) < base.Ui32(v2103))
	if base.Ui32(v2199) < base.Ui32(v2103) {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v2204 = v2195 + int32(1)
	goto L798
L797:
	;
	v2204 = v2186
	goto L798
L798:
	;
	if base.Ui32(v2199) < base.Ui32(v2103) {
		goto L799
	} else {
		goto L800
	}
L799:
	;
	v2207 = v2188
	goto L801
L800:
	;
	v2207 = v2195 - int32(1)
	goto L801
L801:
	;
	if v2204 <= v2207 {
		v2186 = v2204
		v2188 = v2207
		goto L793
	} else {
		goto L802
	}
L802:
	;
	goto L794
L803:
	;
	v3629 = v4
	goto L1
L804:
	;
	goto L805
L805:
	;
	v2224 = F_TransactionIdIsInProgress(m, v2103)
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L64
	} else {
		goto L806
	}
L806:
	;
	if v2224 != 0 {
		goto L754
	} else {
		goto L807
	}
L807:
	;
	v2226 = F_TransactionIdDidCommit(m, v2103)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L64
	} else {
		goto L808
	}
L808:
	;
	v2228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2094)+20)))
	if v2226 == int32(0) {
		v2375 = v2228
		goto L759
	} else {
		goto L809
	}
L809:
	;
	v2232 = v2228 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2094)+20)) = uint16(v2232)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L64
	} else {
		goto L810
	}
L810:
	;
	v3629 = int32(0)
	goto L1
L811:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+8))
	if base.Ui32(v2240) < base.Ui32(int32(3)) {
		goto L815
	} else {
		goto L816
	}
L812:
	;
	goto L813
L813:
	;
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2094)))
	if v2373 != 0 {
		goto L754
	} else {
		goto L862
	}
L814:
	;
	if v2360 != 0 {
		goto L754
	} else {
		goto L854
	}
L815:
	;
	v2360 = int32(0)
	goto L814
L816:
	;
	goto L817
L817:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v2251 == v2240 {
		goto L818
	} else {
		goto L819
	}
L818:
	;
	v2360 = int32(1)
	goto L814
L819:
	;
	goto L820
L820:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v2255 <= int32(0) {
		goto L822
	} else {
		goto L823
	}
L821:
	;
	v2360 = v2352
	goto L814
L822:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v2259 == int32(0) {
		v2352 = int32(0)
		goto L821
	} else {
		goto L825
	}
L823:
	;
	goto L824
L824:
	;
	v2321 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v2323 = int32(0)
	v2325 = v2255 - int32(1)
	goto L844
L825:
	;
	v2264 = v2259
	goto L826
L826:
	;
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+20))
	if v2269 == int32(4) {
		goto L828
	} else {
		goto L829
	}
L827:
	;
	v2352 = int32(0)
	goto L821
L828:
	;
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+80))
	if v2316 != 0 {
		v2264 = v2316
		goto L826
	} else {
		goto L843
	}
L829:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(v2264)))
	if v2272 == int32(0) {
		goto L828
	} else {
		goto L830
	}
L830:
	;
	v2275 = int32(1)
	if v2240 == v2272 {
		v2352 = v2275
		goto L821
	} else {
		goto L831
	}
L831:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+52))
	v2279 = v2277 - int32(1)
	if v2279 < int32(0) {
		goto L828
	} else {
		goto L832
	}
L832:
	;
	v2284 = int32(0)
	v2286 = v2279
	goto L833
L833:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+48))
	v2292 = int32(2)
	v2293 = base.I32_div_s(v2286-v2284, v2292)
	v2294 = v2293 + v2284
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v2290+v2294<<(uint(v2292)%32))))
	if v2298 == v2240 {
		v2352 = v2275
		goto L821
	} else {
		goto L835
	}
L834:
	;
	goto L828
L835:
	;
	v2302 = F_TransactionIdPrecedes(m, v2298, v2240)
	mBase = m.M
	if v2302 != 0 {
		goto L836
	} else {
		goto L837
	}
L836:
	;
	v2303 = v2294 + int32(1)
	goto L838
L837:
	;
	v2303 = v2284
	goto L838
L838:
	;
	if v2302 != 0 {
		goto L839
	} else {
		goto L840
	}
L839:
	;
	v2306 = v2286
	goto L841
L840:
	;
	v2306 = v2294 - int32(1)
	goto L841
L841:
	;
	if v2303 <= v2306 {
		v2284 = v2303
		v2286 = v2306
		goto L833
	} else {
		goto L842
	}
L842:
	;
	goto L834
L843:
	;
	goto L827
L844:
	;
	v2330 = int32(2)
	v2331 = base.I32_div_s(v2325-v2323, v2330)
	v2332 = v2331 + v2323
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v2321+v2332<<(uint(v2330)%32))))
	v2337 = base.B2i32(v2336 == v2240)
	if v2336 == v2240 {
		v2352 = v2337
		goto L821
	} else {
		goto L846
	}
L845:
	;
	v2352 = v2337
	goto L821
L846:
	;
	v2340 = base.B2i32(base.Ui32(v2336) < base.Ui32(v2240))
	if base.Ui32(v2336) < base.Ui32(v2240) {
		goto L847
	} else {
		goto L848
	}
L847:
	;
	v2341 = v2332 + int32(1)
	goto L849
L848:
	;
	v2341 = v2323
	goto L849
L849:
	;
	if base.Ui32(v2336) < base.Ui32(v2240) {
		goto L850
	} else {
		goto L851
	}
L850:
	;
	v2344 = v2325
	goto L852
L851:
	;
	v2344 = v2332 - int32(1)
	goto L852
L852:
	;
	if v2341 <= v2344 {
		v2323 = v2341
		v2325 = v2344
		goto L844
	} else {
		goto L853
	}
L853:
	;
	goto L845
L854:
	;
	v2361 = F_TransactionIdIsInProgress(m, v2240)
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L64
	} else {
		goto L855
	}
L855:
	;
	if v2361 != 0 {
		goto L856
	} else {
		goto L857
	}
L856:
	;
	v3629 = v4
	goto L1
L857:
	;
	goto L858
L858:
	;
	v2363 = F_TransactionIdDidCommit(m, v2240)
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L64
	} else {
		goto L859
	}
L859:
	;
	v2365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2094)+20)))
	if v2363 != 0 {
		v2375 = v2365
		goto L759
	} else {
		goto L860
	}
L860:
	;
	v2367 = v2365 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2094)+20)) = uint16(v2367)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L64
	} else {
		goto L861
	}
L861:
	;
	v3629 = int32(0)
	goto L1
L862:
	;
	v3629 = v4
	goto L1
L863:
	;
	goto L754
L864:
	;
	v3093 = int32(1)
	v3094 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	if v3094&int32(2048) != 0 {
		v3629 = v3093
		goto L1
	} else {
		goto L1112
	}
L865:
	;
	v2393 = base.I32_extend16_s(v2390)
	if v2393&int32(512) != 0 {
		v3629 = v4
		goto L1
	} else {
		goto L866
	}
L866:
	;
	if v2393&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L868
	} else {
		goto L869
	}
L867:
	;
	v3087 = v2523 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v3087)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L64
	} else {
		goto L1111
	}
L868:
	;
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+8))
	if base.Ui32(v2398) < base.Ui32(int32(3)) {
		goto L872
	} else {
		goto L873
	}
L869:
	;
	goto L870
L870:
	;
	if v2393 < int32(0) {
		goto L917
	} else {
		goto L918
	}
L871:
	;
	if v2518 != 0 {
		v3629 = v4
		goto L1
	} else {
		goto L911
	}
L872:
	;
	v2518 = int32(0)
	goto L871
L873:
	;
	goto L874
L874:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v2409 == v2398 {
		goto L875
	} else {
		goto L876
	}
L875:
	;
	v2518 = int32(1)
	goto L871
L876:
	;
	goto L877
L877:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v2413 <= int32(0) {
		goto L879
	} else {
		goto L880
	}
L878:
	;
	v2518 = v2510
	goto L871
L879:
	;
	v2417 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v2417 == int32(0) {
		v2510 = int32(0)
		goto L878
	} else {
		goto L882
	}
L880:
	;
	goto L881
L881:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v2481 = int32(0)
	v2483 = v2413 - int32(1)
	goto L901
L882:
	;
	v2422 = v2417
	goto L883
L883:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+20))
	if v2427 == int32(4) {
		goto L885
	} else {
		goto L886
	}
L884:
	;
	v2510 = int32(0)
	goto L878
L885:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+80))
	if v2474 != 0 {
		v2422 = v2474
		goto L883
	} else {
		goto L900
	}
L886:
	;
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v2422)))
	if v2430 == int32(0) {
		goto L885
	} else {
		goto L887
	}
L887:
	;
	v2433 = int32(1)
	if v2398 == v2430 {
		v2510 = v2433
		goto L878
	} else {
		goto L888
	}
L888:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+52))
	v2437 = v2435 - int32(1)
	if v2437 < int32(0) {
		goto L885
	} else {
		goto L889
	}
L889:
	;
	v2442 = int32(0)
	v2444 = v2437
	goto L890
L890:
	;
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v2422)+48))
	v2450 = int32(2)
	v2451 = base.I32_div_s(v2444-v2442, v2450)
	v2452 = v2451 + v2442
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2448+v2452<<(uint(v2450)%32))))
	if v2456 == v2398 {
		v2510 = v2433
		goto L878
	} else {
		goto L892
	}
L891:
	;
	goto L885
L892:
	;
	v2460 = F_TransactionIdPrecedes(m, v2456, v2398)
	mBase = m.M
	if v2460 != 0 {
		goto L893
	} else {
		goto L894
	}
L893:
	;
	v2461 = v2452 + int32(1)
	goto L895
L894:
	;
	v2461 = v2442
	goto L895
L895:
	;
	if v2460 != 0 {
		goto L896
	} else {
		goto L897
	}
L896:
	;
	v2464 = v2444
	goto L898
L897:
	;
	v2464 = v2452 - int32(1)
	goto L898
L898:
	;
	if v2461 <= v2464 {
		v2442 = v2461
		v2444 = v2464
		goto L890
	} else {
		goto L899
	}
L899:
	;
	goto L891
L900:
	;
	goto L884
L901:
	;
	v2488 = int32(2)
	v2489 = base.I32_div_s(v2483-v2481, v2488)
	v2490 = v2489 + v2481
	v2494 = *(*int32)(unsafe.Add(mBase, uint32(v2479+v2490<<(uint(v2488)%32))))
	v2495 = base.B2i32(v2494 == v2398)
	if v2494 == v2398 {
		v2510 = v2495
		goto L878
	} else {
		goto L903
	}
L902:
	;
	v2510 = v2495
	goto L878
L903:
	;
	v2498 = base.B2i32(base.Ui32(v2494) < base.Ui32(v2398))
	if base.Ui32(v2494) < base.Ui32(v2398) {
		goto L904
	} else {
		goto L905
	}
L904:
	;
	v2499 = v2490 + int32(1)
	goto L906
L905:
	;
	v2499 = v2481
	goto L906
L906:
	;
	if base.Ui32(v2494) < base.Ui32(v2398) {
		goto L907
	} else {
		goto L908
	}
L907:
	;
	v2502 = v2483
	goto L909
L908:
	;
	v2502 = v2490 - int32(1)
	goto L909
L909:
	;
	if v2499 <= v2502 {
		v2481 = v2499
		v2483 = v2502
		goto L901
	} else {
		goto L910
	}
L910:
	;
	goto L902
L911:
	;
	v2519 = F_TransactionIdIsInProgress(m, v2398)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L64
	} else {
		goto L912
	}
L912:
	;
	if v2519 != 0 {
		goto L864
	} else {
		goto L913
	}
L913:
	;
	v2521 = F_TransactionIdDidCommit(m, v2398)
	mBase = m.M
	v2522 = m.ExcPending
	if v2522 != 0 {
		goto L64
	} else {
		goto L914
	}
L914:
	;
	v2523 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	if v2521 == int32(0) {
		goto L867
	} else {
		goto L915
	}
L915:
	;
	v2527 = v2523 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v2527)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L64
	} else {
		goto L916
	}
L916:
	;
	v3629 = v4
	goto L1
L917:
	;
	v2534 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+8))
	if base.Ui32(v2534) < base.Ui32(int32(3)) {
		goto L921
	} else {
		goto L922
	}
L918:
	;
	goto L919
L919:
	;
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2385)))
	if base.Ui32(v2672) < base.Ui32(int32(3)) {
		goto L970
	} else {
		goto L971
	}
L920:
	;
	if v2654 != 0 {
		goto L864
	} else {
		goto L960
	}
L921:
	;
	v2654 = int32(0)
	goto L920
L922:
	;
	goto L923
L923:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v2545 == v2534 {
		goto L924
	} else {
		goto L925
	}
L924:
	;
	v2654 = int32(1)
	goto L920
L925:
	;
	goto L926
L926:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v2549 <= int32(0) {
		goto L928
	} else {
		goto L929
	}
L927:
	;
	v2654 = v2646
	goto L920
L928:
	;
	v2553 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v2553 == int32(0) {
		v2646 = int32(0)
		goto L927
	} else {
		goto L931
	}
L929:
	;
	goto L930
L930:
	;
	v2615 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v2617 = int32(0)
	v2619 = v2549 - int32(1)
	goto L950
L931:
	;
	v2558 = v2553
	goto L932
L932:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v2558)+20))
	if v2563 == int32(4) {
		goto L934
	} else {
		goto L935
	}
L933:
	;
	v2646 = int32(0)
	goto L927
L934:
	;
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v2558)+80))
	if v2610 != 0 {
		v2558 = v2610
		goto L932
	} else {
		goto L949
	}
L935:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v2558)))
	if v2566 == int32(0) {
		goto L934
	} else {
		goto L936
	}
L936:
	;
	v2569 = int32(1)
	if v2534 == v2566 {
		v2646 = v2569
		goto L927
	} else {
		goto L937
	}
L937:
	;
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(v2558)+52))
	v2573 = v2571 - int32(1)
	if v2573 < int32(0) {
		goto L934
	} else {
		goto L938
	}
L938:
	;
	v2578 = int32(0)
	v2580 = v2573
	goto L939
L939:
	;
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(v2558)+48))
	v2586 = int32(2)
	v2587 = base.I32_div_s(v2580-v2578, v2586)
	v2588 = v2587 + v2578
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v2584+v2588<<(uint(v2586)%32))))
	if v2592 == v2534 {
		v2646 = v2569
		goto L927
	} else {
		goto L941
	}
L940:
	;
	goto L934
L941:
	;
	v2596 = F_TransactionIdPrecedes(m, v2592, v2534)
	mBase = m.M
	if v2596 != 0 {
		goto L942
	} else {
		goto L943
	}
L942:
	;
	v2597 = v2588 + int32(1)
	goto L944
L943:
	;
	v2597 = v2578
	goto L944
L944:
	;
	if v2596 != 0 {
		goto L945
	} else {
		goto L946
	}
L945:
	;
	v2600 = v2580
	goto L947
L946:
	;
	v2600 = v2588 - int32(1)
	goto L947
L947:
	;
	if v2597 <= v2600 {
		v2578 = v2597
		v2580 = v2600
		goto L939
	} else {
		goto L948
	}
L948:
	;
	goto L940
L949:
	;
	goto L933
L950:
	;
	v2624 = int32(2)
	v2625 = base.I32_div_s(v2619-v2617, v2624)
	v2626 = v2625 + v2617
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2615+v2626<<(uint(v2624)%32))))
	v2631 = base.B2i32(v2630 == v2534)
	if v2630 == v2534 {
		v2646 = v2631
		goto L927
	} else {
		goto L952
	}
L951:
	;
	v2646 = v2631
	goto L927
L952:
	;
	v2634 = base.B2i32(base.Ui32(v2630) < base.Ui32(v2534))
	if base.Ui32(v2630) < base.Ui32(v2534) {
		goto L953
	} else {
		goto L954
	}
L953:
	;
	v2635 = v2626 + int32(1)
	goto L955
L954:
	;
	v2635 = v2617
	goto L955
L955:
	;
	if base.Ui32(v2630) < base.Ui32(v2534) {
		goto L956
	} else {
		goto L957
	}
L956:
	;
	v2638 = v2619
	goto L958
L957:
	;
	v2638 = v2626 - int32(1)
	goto L958
L958:
	;
	if v2635 <= v2638 {
		v2617 = v2635
		v2619 = v2638
		goto L950
	} else {
		goto L959
	}
L959:
	;
	goto L951
L960:
	;
	v2655 = F_TransactionIdIsInProgress(m, v2534)
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L64
	} else {
		goto L961
	}
L961:
	;
	if v2655 != 0 {
		v3629 = v4
		goto L1
	} else {
		goto L962
	}
L962:
	;
	v2657 = F_TransactionIdDidCommit(m, v2534)
	mBase = m.M
	v2658 = m.ExcPending
	if v2658 != 0 {
		goto L64
	} else {
		goto L963
	}
L963:
	;
	v2659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	if v2657 != 0 {
		goto L964
	} else {
		goto L965
	}
L964:
	;
	v2661 = v2659 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v2661)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L64
	} else {
		goto L967
	}
L965:
	;
	goto L966
L966:
	;
	v2667 = v2659 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v2667)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L64
	} else {
		goto L968
	}
L967:
	;
	goto L864
L968:
	;
	v3629 = v4
	goto L1
L969:
	;
	if v2792 != 0 {
		goto L1009
	} else {
		goto L1010
	}
L970:
	;
	v2792 = int32(0)
	goto L969
L971:
	;
	goto L972
L972:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v2683 == v2672 {
		goto L973
	} else {
		goto L974
	}
L973:
	;
	v2792 = int32(1)
	goto L969
L974:
	;
	goto L975
L975:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v2687 <= int32(0) {
		goto L977
	} else {
		goto L978
	}
L976:
	;
	v2792 = v2784
	goto L969
L977:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v2691 == int32(0) {
		v2784 = int32(0)
		goto L976
	} else {
		goto L980
	}
L978:
	;
	goto L979
L979:
	;
	v2753 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v2755 = int32(0)
	v2757 = v2687 - int32(1)
	goto L999
L980:
	;
	v2696 = v2691
	goto L981
L981:
	;
	v2701 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+20))
	if v2701 == int32(4) {
		goto L983
	} else {
		goto L984
	}
L982:
	;
	v2784 = int32(0)
	goto L976
L983:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+80))
	if v2748 != 0 {
		v2696 = v2748
		goto L981
	} else {
		goto L998
	}
L984:
	;
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v2696)))
	if v2704 == int32(0) {
		goto L983
	} else {
		goto L985
	}
L985:
	;
	v2707 = int32(1)
	if v2672 == v2704 {
		v2784 = v2707
		goto L976
	} else {
		goto L986
	}
L986:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+52))
	v2711 = v2709 - int32(1)
	if v2711 < int32(0) {
		goto L983
	} else {
		goto L987
	}
L987:
	;
	v2716 = int32(0)
	v2718 = v2711
	goto L988
L988:
	;
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+48))
	v2724 = int32(2)
	v2725 = base.I32_div_s(v2718-v2716, v2724)
	v2726 = v2725 + v2716
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2722+v2726<<(uint(v2724)%32))))
	if v2730 == v2672 {
		v2784 = v2707
		goto L976
	} else {
		goto L990
	}
L989:
	;
	goto L983
L990:
	;
	v2734 = F_TransactionIdPrecedes(m, v2730, v2672)
	mBase = m.M
	if v2734 != 0 {
		goto L991
	} else {
		goto L992
	}
L991:
	;
	v2735 = v2726 + int32(1)
	goto L993
L992:
	;
	v2735 = v2716
	goto L993
L993:
	;
	if v2734 != 0 {
		goto L994
	} else {
		goto L995
	}
L994:
	;
	v2738 = v2718
	goto L996
L995:
	;
	v2738 = v2726 - int32(1)
	goto L996
L996:
	;
	if v2735 <= v2738 {
		v2716 = v2735
		v2718 = v2738
		goto L988
	} else {
		goto L997
	}
L997:
	;
	goto L989
L998:
	;
	goto L982
L999:
	;
	v2762 = int32(2)
	v2763 = base.I32_div_s(v2757-v2755, v2762)
	v2764 = v2763 + v2755
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v2753+v2764<<(uint(v2762)%32))))
	v2769 = base.B2i32(v2768 == v2672)
	if v2768 == v2672 {
		v2784 = v2769
		goto L976
	} else {
		goto L1001
	}
L1000:
	;
	v2784 = v2769
	goto L976
L1001:
	;
	v2772 = base.B2i32(base.Ui32(v2768) < base.Ui32(v2672))
	if base.Ui32(v2768) < base.Ui32(v2672) {
		goto L1002
	} else {
		goto L1003
	}
L1002:
	;
	v2773 = v2764 + int32(1)
	goto L1004
L1003:
	;
	v2773 = v2755
	goto L1004
L1004:
	;
	if base.Ui32(v2768) < base.Ui32(v2672) {
		goto L1005
	} else {
		goto L1006
	}
L1005:
	;
	v2776 = v2757
	goto L1007
L1006:
	;
	v2776 = v2764 - int32(1)
	goto L1007
L1007:
	;
	if v2773 <= v2776 {
		v2755 = v2773
		v2757 = v2776
		goto L999
	} else {
		goto L1008
	}
L1008:
	;
	goto L1000
L1009:
	;
	v2794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	if v2794&int32(2048)|v2794&int32(128)|base.B2i32(v2794&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64)) != 0 {
		v3629 = int32(1)
		goto L1
	} else {
		goto L1012
	}
L1010:
	;
	goto L1011
L1011:
	;
	v3057 = *(*int32)(unsafe.Add(mBase, uint32(v2385)))
	v3058 = F_TransactionIdIsInProgress(m, v3057)
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L64
	} else {
		goto L1098
	}
L1012:
	;
	if v2794&int32(_a_F_HeapTupleSatisfiesVisibility_2) != 0 {
		goto L1013
	} else {
		goto L1014
	}
L1013:
	;
	v2807 = F_HeapTupleGetUpdateXid(m, v2385)
	mBase = m.M
	v2808 = m.ExcPending
	if v2808 != 0 {
		goto L64
	} else {
		goto L1016
	}
L1014:
	;
	goto L1015
L1015:
	;
	v2931 = int32(0)
	v2932 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	if base.Ui32(v2932) < base.Ui32(int32(3)) {
		goto L1058
	} else {
		goto L1059
	}
L1016:
	;
	if base.Ui32(v2807) < base.Ui32(int32(3)) {
		goto L1018
	} else {
		goto L1019
	}
L1017:
	;
	v3629 = v2928 ^ int32(1)
	goto L1
L1018:
	;
	v2928 = int32(0)
	goto L1017
L1019:
	;
	goto L1020
L1020:
	;
	v2819 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v2819 == v2807 {
		goto L1021
	} else {
		goto L1022
	}
L1021:
	;
	v2928 = int32(1)
	goto L1017
L1022:
	;
	goto L1023
L1023:
	;
	v2823 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v2823 <= int32(0) {
		goto L1025
	} else {
		goto L1026
	}
L1024:
	;
	v2928 = v2920
	goto L1017
L1025:
	;
	v2827 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v2827 == int32(0) {
		v2920 = int32(0)
		goto L1024
	} else {
		goto L1028
	}
L1026:
	;
	goto L1027
L1027:
	;
	v2889 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v2891 = int32(0)
	v2893 = v2823 - int32(1)
	goto L1047
L1028:
	;
	v2832 = v2827
	goto L1029
L1029:
	;
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+20))
	if v2837 == int32(4) {
		goto L1031
	} else {
		goto L1032
	}
L1030:
	;
	v2920 = int32(0)
	goto L1024
L1031:
	;
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+80))
	if v2884 != 0 {
		v2832 = v2884
		goto L1029
	} else {
		goto L1046
	}
L1032:
	;
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2832)))
	if v2840 == int32(0) {
		goto L1031
	} else {
		goto L1033
	}
L1033:
	;
	v2843 = int32(1)
	if v2807 == v2840 {
		v2920 = v2843
		goto L1024
	} else {
		goto L1034
	}
L1034:
	;
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+52))
	v2847 = v2845 - int32(1)
	if v2847 < int32(0) {
		goto L1031
	} else {
		goto L1035
	}
L1035:
	;
	v2852 = int32(0)
	v2854 = v2847
	goto L1036
L1036:
	;
	v2858 = *(*int32)(unsafe.Add(mBase, uint32(v2832)+48))
	v2860 = int32(2)
	v2861 = base.I32_div_s(v2854-v2852, v2860)
	v2862 = v2861 + v2852
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v2858+v2862<<(uint(v2860)%32))))
	if v2866 == v2807 {
		v2920 = v2843
		goto L1024
	} else {
		goto L1038
	}
L1037:
	;
	goto L1031
L1038:
	;
	v2870 = F_TransactionIdPrecedes(m, v2866, v2807)
	mBase = m.M
	if v2870 != 0 {
		goto L1039
	} else {
		goto L1040
	}
L1039:
	;
	v2871 = v2862 + int32(1)
	goto L1041
L1040:
	;
	v2871 = v2852
	goto L1041
L1041:
	;
	if v2870 != 0 {
		goto L1042
	} else {
		goto L1043
	}
L1042:
	;
	v2874 = v2854
	goto L1044
L1043:
	;
	v2874 = v2862 - int32(1)
	goto L1044
L1044:
	;
	if v2871 <= v2874 {
		v2852 = v2871
		v2854 = v2874
		goto L1036
	} else {
		goto L1045
	}
L1045:
	;
	goto L1037
L1046:
	;
	goto L1030
L1047:
	;
	v2898 = int32(2)
	v2899 = base.I32_div_s(v2893-v2891, v2898)
	v2900 = v2899 + v2891
	v2904 = *(*int32)(unsafe.Add(mBase, uint32(v2889+v2900<<(uint(v2898)%32))))
	v2905 = base.B2i32(v2904 == v2807)
	if v2904 == v2807 {
		v2920 = v2905
		goto L1024
	} else {
		goto L1049
	}
L1048:
	;
	v2920 = v2905
	goto L1024
L1049:
	;
	v2908 = base.B2i32(base.Ui32(v2904) < base.Ui32(v2807))
	if base.Ui32(v2904) < base.Ui32(v2807) {
		goto L1050
	} else {
		goto L1051
	}
L1050:
	;
	v2909 = v2900 + int32(1)
	goto L1052
L1051:
	;
	v2909 = v2891
	goto L1052
L1052:
	;
	if base.Ui32(v2904) < base.Ui32(v2807) {
		goto L1053
	} else {
		goto L1054
	}
L1053:
	;
	v2912 = v2893
	goto L1055
L1054:
	;
	v2912 = v2900 - int32(1)
	goto L1055
L1055:
	;
	if v2909 <= v2912 {
		v2891 = v2909
		v2893 = v2912
		goto L1047
	} else {
		goto L1056
	}
L1056:
	;
	goto L1048
L1057:
	;
	if v3052 != 0 {
		v3629 = v2931
		goto L1
	} else {
		goto L1097
	}
L1058:
	;
	v3052 = int32(0)
	goto L1057
L1059:
	;
	goto L1060
L1060:
	;
	v2943 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v2943 == v2932 {
		goto L1061
	} else {
		goto L1062
	}
L1061:
	;
	v3052 = int32(1)
	goto L1057
L1062:
	;
	goto L1063
L1063:
	;
	v2947 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v2947 <= int32(0) {
		goto L1065
	} else {
		goto L1066
	}
L1064:
	;
	v3052 = v3044
	goto L1057
L1065:
	;
	v2951 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v2951 == int32(0) {
		v3044 = v2931
		goto L1064
	} else {
		goto L1068
	}
L1066:
	;
	goto L1067
L1067:
	;
	v3013 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v3015 = int32(0)
	v3017 = v2947 - int32(1)
	goto L1087
L1068:
	;
	v2956 = v2951
	goto L1069
L1069:
	;
	v2961 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+20))
	if v2961 == int32(4) {
		goto L1071
	} else {
		goto L1072
	}
L1070:
	;
	v3044 = int32(0)
	goto L1064
L1071:
	;
	v3008 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+80))
	if v3008 != 0 {
		v2956 = v3008
		goto L1069
	} else {
		goto L1086
	}
L1072:
	;
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v2956)))
	if v2964 == int32(0) {
		goto L1071
	} else {
		goto L1073
	}
L1073:
	;
	v2967 = int32(1)
	if v2932 == v2964 {
		v3044 = v2967
		goto L1064
	} else {
		goto L1074
	}
L1074:
	;
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+52))
	v2971 = v2969 - int32(1)
	if v2971 < int32(0) {
		goto L1071
	} else {
		goto L1075
	}
L1075:
	;
	v2976 = int32(0)
	v2978 = v2971
	goto L1076
L1076:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+48))
	v2984 = int32(2)
	v2985 = base.I32_div_s(v2978-v2976, v2984)
	v2986 = v2985 + v2976
	v2990 = *(*int32)(unsafe.Add(mBase, uint32(v2982+v2986<<(uint(v2984)%32))))
	if v2990 == v2932 {
		v3044 = v2967
		goto L1064
	} else {
		goto L1078
	}
L1077:
	;
	goto L1071
L1078:
	;
	v2994 = F_TransactionIdPrecedes(m, v2990, v2932)
	mBase = m.M
	if v2994 != 0 {
		goto L1079
	} else {
		goto L1080
	}
L1079:
	;
	v2995 = v2986 + int32(1)
	goto L1081
L1080:
	;
	v2995 = v2976
	goto L1081
L1081:
	;
	if v2994 != 0 {
		goto L1082
	} else {
		goto L1083
	}
L1082:
	;
	v2998 = v2978
	goto L1084
L1083:
	;
	v2998 = v2986 - int32(1)
	goto L1084
L1084:
	;
	if v2995 <= v2998 {
		v2976 = v2995
		v2978 = v2998
		goto L1076
	} else {
		goto L1085
	}
L1085:
	;
	goto L1077
L1086:
	;
	goto L1070
L1087:
	;
	v3022 = int32(2)
	v3023 = base.I32_div_s(v3017-v3015, v3022)
	v3024 = v3023 + v3015
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v3013+v3024<<(uint(v3022)%32))))
	v3029 = base.B2i32(v3028 == v2932)
	if v3028 == v2932 {
		v3044 = v3029
		goto L1064
	} else {
		goto L1089
	}
L1088:
	;
	v3044 = v3029
	goto L1064
L1089:
	;
	v3032 = base.B2i32(base.Ui32(v3028) < base.Ui32(v2932))
	if base.Ui32(v3028) < base.Ui32(v2932) {
		goto L1090
	} else {
		goto L1091
	}
L1090:
	;
	v3033 = v3024 + int32(1)
	goto L1092
L1091:
	;
	v3033 = v3015
	goto L1092
L1092:
	;
	if base.Ui32(v3028) < base.Ui32(v2932) {
		goto L1093
	} else {
		goto L1094
	}
L1093:
	;
	v3036 = v3017
	goto L1095
L1094:
	;
	v3036 = v3024 - int32(1)
	goto L1095
L1095:
	;
	if v3033 <= v3036 {
		v3015 = v3033
		v3017 = v3036
		goto L1087
	} else {
		goto L1096
	}
L1096:
	;
	goto L1088
L1097:
	;
	v3053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	v3055 = v3053 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v3055)
	goto L2
L1098:
	;
	if v3058 != 0 {
		goto L1099
	} else {
		goto L1100
	}
L1099:
	;
	v3060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+16)))
	if v3060 == int32(_a_F_HeapTupleSatisfiesVisibility_3) {
		goto L1102
	} else {
		goto L1103
	}
L1100:
	;
	goto L1101
L1101:
	;
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v2385)))
	v3073 = F_TransactionIdDidCommit(m, v3072)
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L64
	} else {
		goto L1105
	}
L1102:
	;
	v3063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+14)))
	v3064 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v3063 | v3064<<(uint(int32(16))%32)
	goto L1104
L1103:
	;
	goto L1104
L1104:
	;
	v3069 = *(*int32)(unsafe.Add(mBase, uint32(v2385)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3069
	v3629 = int32(1)
	goto L1
L1105:
	;
	if v3073 != 0 {
		goto L1106
	} else {
		goto L1107
	}
L1106:
	;
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v2385)))
	F_HeapTupleSetHintBits(m, v2385, l2, int32(256), v3076)
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L64
	} else {
		goto L1109
	}
L1107:
	;
	goto L1108
L1108:
	;
	v3079 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	v3081 = v3079 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v3081)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		goto L64
	} else {
		goto L1110
	}
L1109:
	;
	goto L864
L1110:
	;
	v3629 = v4
	goto L1
L1111:
	;
	goto L864
L1112:
	;
	if v3094&int32(1024) != 0 {
		goto L1113
	} else {
		goto L1114
	}
L1113:
	;
	v3629 = int32(base.Ui32(v3094&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v3094&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L1114:
	;
	goto L1115
L1115:
	;
	if v3094&int32(_a_F_HeapTupleSatisfiesVisibility_2) != 0 {
		goto L1116
	} else {
		goto L1117
	}
L1116:
	;
	if v3094&int32(128) != 0 {
		v3629 = v3093
		goto L1
	} else {
		goto L1119
	}
L1117:
	;
	goto L1118
L1118:
	;
	v3243 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	if base.Ui32(v3243) < base.Ui32(int32(3)) {
		goto L1168
	} else {
		goto L1169
	}
L1119:
	;
	v3113 = F_HeapTupleGetUpdateXid(m, v2385)
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L64
	} else {
		goto L1120
	}
L1120:
	;
	if base.Ui32(v3113) < base.Ui32(int32(3)) {
		goto L1122
	} else {
		goto L1123
	}
L1121:
	;
	if v3234 != 0 {
		v3629 = int32(0)
		goto L1
	} else {
		goto L1161
	}
L1122:
	;
	v3234 = int32(0)
	goto L1121
L1123:
	;
	goto L1124
L1124:
	;
	v3125 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v3125 == v3113 {
		goto L1125
	} else {
		goto L1126
	}
L1125:
	;
	v3234 = int32(1)
	goto L1121
L1126:
	;
	goto L1127
L1127:
	;
	v3129 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v3129 <= int32(0) {
		goto L1129
	} else {
		goto L1130
	}
L1128:
	;
	v3234 = v3226
	goto L1121
L1129:
	;
	v3133 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v3133 == int32(0) {
		v3226 = int32(0)
		goto L1128
	} else {
		goto L1132
	}
L1130:
	;
	goto L1131
L1131:
	;
	v3195 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v3197 = int32(0)
	v3199 = v3129 - int32(1)
	goto L1151
L1132:
	;
	v3138 = v3133
	goto L1133
L1133:
	;
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(v3138)+20))
	if v3143 == int32(4) {
		goto L1135
	} else {
		goto L1136
	}
L1134:
	;
	v3226 = int32(0)
	goto L1128
L1135:
	;
	v3190 = *(*int32)(unsafe.Add(mBase, uint32(v3138)+80))
	if v3190 != 0 {
		v3138 = v3190
		goto L1133
	} else {
		goto L1150
	}
L1136:
	;
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v3138)))
	if v3146 == int32(0) {
		goto L1135
	} else {
		goto L1137
	}
L1137:
	;
	v3149 = int32(1)
	if v3113 == v3146 {
		v3226 = v3149
		goto L1128
	} else {
		goto L1138
	}
L1138:
	;
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v3138)+52))
	v3153 = v3151 - int32(1)
	if v3153 < int32(0) {
		goto L1135
	} else {
		goto L1139
	}
L1139:
	;
	v3158 = int32(0)
	v3160 = v3153
	goto L1140
L1140:
	;
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v3138)+48))
	v3166 = int32(2)
	v3167 = base.I32_div_s(v3160-v3158, v3166)
	v3168 = v3167 + v3158
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(v3164+v3168<<(uint(v3166)%32))))
	if v3172 == v3113 {
		v3226 = v3149
		goto L1128
	} else {
		goto L1142
	}
L1141:
	;
	goto L1135
L1142:
	;
	v3176 = F_TransactionIdPrecedes(m, v3172, v3113)
	mBase = m.M
	if v3176 != 0 {
		goto L1143
	} else {
		goto L1144
	}
L1143:
	;
	v3177 = v3168 + int32(1)
	goto L1145
L1144:
	;
	v3177 = v3158
	goto L1145
L1145:
	;
	if v3176 != 0 {
		goto L1146
	} else {
		goto L1147
	}
L1146:
	;
	v3180 = v3160
	goto L1148
L1147:
	;
	v3180 = v3168 - int32(1)
	goto L1148
L1148:
	;
	if v3177 <= v3180 {
		v3158 = v3177
		v3160 = v3180
		goto L1140
	} else {
		goto L1149
	}
L1149:
	;
	goto L1141
L1150:
	;
	goto L1134
L1151:
	;
	v3204 = int32(2)
	v3205 = base.I32_div_s(v3199-v3197, v3204)
	v3206 = v3205 + v3197
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v3195+v3206<<(uint(v3204)%32))))
	v3211 = base.B2i32(v3210 == v3113)
	if v3210 == v3113 {
		v3226 = v3211
		goto L1128
	} else {
		goto L1153
	}
L1152:
	;
	v3226 = v3211
	goto L1128
L1153:
	;
	v3214 = base.B2i32(base.Ui32(v3210) < base.Ui32(v3113))
	if base.Ui32(v3210) < base.Ui32(v3113) {
		goto L1154
	} else {
		goto L1155
	}
L1154:
	;
	v3215 = v3206 + int32(1)
	goto L1156
L1155:
	;
	v3215 = v3197
	goto L1156
L1156:
	;
	if base.Ui32(v3210) < base.Ui32(v3113) {
		goto L1157
	} else {
		goto L1158
	}
L1157:
	;
	v3218 = v3199
	goto L1159
L1158:
	;
	v3218 = v3206 - int32(1)
	goto L1159
L1159:
	;
	if v3215 <= v3218 {
		v3197 = v3215
		v3199 = v3218
		goto L1151
	} else {
		goto L1160
	}
L1160:
	;
	goto L1152
L1161:
	;
	v3235 = F_TransactionIdIsInProgress(m, v3113)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L64
	} else {
		goto L1162
	}
L1162:
	;
	if v3235 != 0 {
		goto L1163
	} else {
		goto L1164
	}
L1163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3113
	v3629 = int32(1)
	goto L1
L1164:
	;
	goto L1165
L1165:
	;
	v3239 = F_TransactionIdDidCommit(m, v3113)
	mBase = m.M
	v3240 = m.ExcPending
	if v3240 != 0 {
		goto L64
	} else {
		goto L1166
	}
L1166:
	;
	v3629 = v3239 ^ int32(1)
	goto L1
L1167:
	;
	if v3363 != 0 {
		goto L1207
	} else {
		goto L1208
	}
L1168:
	;
	v3363 = int32(0)
	goto L1167
L1169:
	;
	goto L1170
L1170:
	;
	v3254 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v3254 == v3243 {
		goto L1171
	} else {
		goto L1172
	}
L1171:
	;
	v3363 = int32(1)
	goto L1167
L1172:
	;
	goto L1173
L1173:
	;
	v3258 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v3258 <= int32(0) {
		goto L1175
	} else {
		goto L1176
	}
L1174:
	;
	v3363 = v3355
	goto L1167
L1175:
	;
	v3262 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v3262 == int32(0) {
		v3355 = int32(0)
		goto L1174
	} else {
		goto L1178
	}
L1176:
	;
	goto L1177
L1177:
	;
	v3324 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v3326 = int32(0)
	v3328 = v3258 - int32(1)
	goto L1197
L1178:
	;
	v3267 = v3262
	goto L1179
L1179:
	;
	v3272 = *(*int32)(unsafe.Add(mBase, uint32(v3267)+20))
	if v3272 == int32(4) {
		goto L1181
	} else {
		goto L1182
	}
L1180:
	;
	v3355 = int32(0)
	goto L1174
L1181:
	;
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v3267)+80))
	if v3319 != 0 {
		v3267 = v3319
		goto L1179
	} else {
		goto L1196
	}
L1182:
	;
	v3275 = *(*int32)(unsafe.Add(mBase, uint32(v3267)))
	if v3275 == int32(0) {
		goto L1181
	} else {
		goto L1183
	}
L1183:
	;
	v3278 = int32(1)
	if v3243 == v3275 {
		v3355 = v3278
		goto L1174
	} else {
		goto L1184
	}
L1184:
	;
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v3267)+52))
	v3282 = v3280 - int32(1)
	if v3282 < int32(0) {
		goto L1181
	} else {
		goto L1185
	}
L1185:
	;
	v3287 = int32(0)
	v3289 = v3282
	goto L1186
L1186:
	;
	v3293 = *(*int32)(unsafe.Add(mBase, uint32(v3267)+48))
	v3295 = int32(2)
	v3296 = base.I32_div_s(v3289-v3287, v3295)
	v3297 = v3296 + v3287
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3293+v3297<<(uint(v3295)%32))))
	if v3301 == v3243 {
		v3355 = v3278
		goto L1174
	} else {
		goto L1188
	}
L1187:
	;
	goto L1181
L1188:
	;
	v3305 = F_TransactionIdPrecedes(m, v3301, v3243)
	mBase = m.M
	if v3305 != 0 {
		goto L1189
	} else {
		goto L1190
	}
L1189:
	;
	v3306 = v3297 + int32(1)
	goto L1191
L1190:
	;
	v3306 = v3287
	goto L1191
L1191:
	;
	if v3305 != 0 {
		goto L1192
	} else {
		goto L1193
	}
L1192:
	;
	v3309 = v3289
	goto L1194
L1193:
	;
	v3309 = v3297 - int32(1)
	goto L1194
L1194:
	;
	if v3306 <= v3309 {
		v3287 = v3306
		v3289 = v3309
		goto L1186
	} else {
		goto L1195
	}
L1195:
	;
	goto L1187
L1196:
	;
	goto L1180
L1197:
	;
	v3333 = int32(2)
	v3334 = base.I32_div_s(v3328-v3326, v3333)
	v3335 = v3334 + v3326
	v3339 = *(*int32)(unsafe.Add(mBase, uint32(v3324+v3335<<(uint(v3333)%32))))
	v3340 = base.B2i32(v3339 == v3243)
	if v3339 == v3243 {
		v3355 = v3340
		goto L1174
	} else {
		goto L1199
	}
L1198:
	;
	v3355 = v3340
	goto L1174
L1199:
	;
	v3343 = base.B2i32(base.Ui32(v3339) < base.Ui32(v3243))
	if base.Ui32(v3339) < base.Ui32(v3243) {
		goto L1200
	} else {
		goto L1201
	}
L1200:
	;
	v3344 = v3335 + int32(1)
	goto L1202
L1201:
	;
	v3344 = v3326
	goto L1202
L1202:
	;
	if base.Ui32(v3339) < base.Ui32(v3243) {
		goto L1203
	} else {
		goto L1204
	}
L1203:
	;
	v3347 = v3328
	goto L1205
L1204:
	;
	v3347 = v3335 - int32(1)
	goto L1205
L1205:
	;
	if v3344 <= v3347 {
		v3326 = v3344
		v3328 = v3347
		goto L1197
	} else {
		goto L1206
	}
L1206:
	;
	goto L1198
L1207:
	;
	v3364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	v3629 = int32(base.Ui32(v3364&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v3364&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L1208:
	;
	goto L1209
L1209:
	;
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	v3375 = F_TransactionIdIsInProgress(m, v3374)
	mBase = m.M
	v3376 = m.ExcPending
	if v3376 != 0 {
		goto L64
	} else {
		goto L1210
	}
L1210:
	;
	if v3375 != 0 {
		goto L1211
	} else {
		goto L1212
	}
L1211:
	;
	v3377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	if v3377&int32(128)|base.B2i32(v3377&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64)) != 0 {
		v3629 = v3093
		goto L1
	} else {
		goto L1214
	}
L1212:
	;
	goto L1213
L1213:
	;
	v3387 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	v3388 = F_TransactionIdDidCommit(m, v3387)
	mBase = m.M
	v3389 = m.ExcPending
	if v3389 != 0 {
		goto L64
	} else {
		goto L1215
	}
L1214:
	;
	v3385 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3385
	v3629 = v3093
	goto L1
L1215:
	;
	v3390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)))
	if v3388 == int32(0) {
		goto L1216
	} else {
		goto L1217
	}
L1216:
	;
	v3394 = v3390 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v3394)
	goto L2
L1217:
	;
	goto L1218
L1218:
	;
	v3398 = int32(0)
	if base.B2i32(v3390&int32(128) == v3398)&base.B2i32(v3390&int32(_a_F_HeapTupleSatisfiesVisibility_0) != int32(64)) == v3398 {
		goto L1219
	} else {
		goto L1220
	}
L1219:
	;
	v3408 = v3390 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v2385)+20)) = uint16(v3408)
	goto L2
L1220:
	;
	goto L1221
L1221:
	;
	v3411 = *(*int32)(unsafe.Add(mBase, uint32(v2385)+4))
	F_HeapTupleSetHintBits(m, v2385, l2, int32(1024), v3411)
	mBase = m.M
	v3413 = m.ExcPending
	if v3413 != 0 {
		goto L64
	} else {
		goto L1222
	}
L1222:
	;
	v3629 = int32(0)
	goto L1
L1223:
	;
	if v3419 == int32(512) {
		goto L1226
	} else {
		goto L1227
	}
L1224:
	;
	v3426 = int32(2)
	goto L1225
L1225:
	;
	v3427 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v3428 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3429 = *(*int32)(unsafe.Add(mBase, uint32(v3416)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3426
	if v3427 == int32(0) {
		goto L1231
	} else {
		goto L1232
	}
L1226:
	;
	v3629 = int32(0)
	goto L1
L1227:
	;
	goto L1228
L1228:
	;
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v3416)))
	v3426 = v3425
	goto L1225
L1229:
	;
	v3629 = int32(0)
	goto L1
L1230:
	;
	v3507 = int32(1)
	v3508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3416)+20)))
	if v3508&int32(2048)|v3508&int32(128)|base.B2i32(v3508&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64)) != 0 {
		v3629 = v3507
		goto L1
	} else {
		goto L1260
	}
L1231:
	;
	v3458 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3458))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3426)) == int32(0) {
		goto L1239
	} else {
		goto L1240
	}
L1232:
	;
	v3434 = v13 + int32(12)
	v3437 = F_bsearch(m, v3434, v3428, v3427, int32(4), int32(185))
	mBase = m.M
	v3438 = m.ExcPending
	if v3438 != 0 {
		goto L64
	} else {
		goto L1233
	}
L1233:
	;
	if v3437 == int32(0) {
		goto L1231
	} else {
		goto L1234
	}
L1234:
	;
	v3441 = *(*int32)(unsafe.Add(mBase, uint32(v3416)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3441
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(-1)
	v3446 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[5]))
	v3449 = F_ResolveCminCmaxDuringDecoding(m, v3446, l1, l0, l2, v3434, v13+int32(8))
	mBase = m.M
	v3450 = m.ExcPending
	if v3450 != 0 {
		goto L64
	} else {
		goto L1235
	}
L1235:
	;
	if v3449 == int32(0) {
		goto L1229
	} else {
		goto L1236
	}
L1236:
	;
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v3454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui32(v3453) < base.Ui32(v3454) {
		goto L1230
	} else {
		goto L1237
	}
L1237:
	;
	v3629 = int32(0)
	goto L1
L1238:
	;
	if v3470 != 0 {
		goto L1242
	} else {
		goto L1243
	}
L1239:
	;
	v3470 = base.B2i32(base.Ui32(v3426) < base.Ui32(v3458))
	goto L1238
L1240:
	;
	goto L1241
L1241:
	;
	v3470 = int32(base.Ui32(v3426-v3458) >> (uint(int32(31)) % 32))
	goto L1238
L1242:
	;
	v3471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3416)+21)))
	if v3471&int32(1) != 0 {
		goto L1230
	} else {
		goto L1245
	}
L1243:
	;
	goto L1244
L1244:
	;
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3477))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3426)) == int32(0) {
		goto L1249
	} else {
		goto L1250
	}
L1245:
	;
	v3474 = F_TransactionIdDidCommit(m, v3426)
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L64
	} else {
		goto L1246
	}
L1246:
	;
	if v3474 != 0 {
		goto L1230
	} else {
		goto L1247
	}
L1247:
	;
	v3629 = int32(0)
	goto L1
L1248:
	;
	if v3489 != 0 {
		goto L1252
	} else {
		goto L1253
	}
L1249:
	;
	v3489 = base.B2i32(base.Ui32(v3477) <= base.Ui32(v3426))
	goto L1248
L1250:
	;
	goto L1251
L1251:
	;
	v3489 = base.B2i32(int32(0) <= v3426-v3477)
	goto L1248
L1252:
	;
	v3629 = int32(0)
	goto L1
L1253:
	;
	goto L1254
L1254:
	;
	v3491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3492 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3426
	if v3491 == int32(0) {
		goto L1255
	} else {
		goto L1256
	}
L1255:
	;
	v3629 = int32(0)
	goto L1
L1256:
	;
	goto L1257
L1257:
	;
	v3501 = F_bsearch(m, v13+int32(12), v3492, v3491, int32(4), int32(185))
	mBase = m.M
	v3502 = m.ExcPending
	if v3502 != 0 {
		goto L64
	} else {
		goto L1258
	}
L1258:
	;
	if v3501 != 0 {
		goto L1230
	} else {
		goto L1259
	}
L1259:
	;
	v3629 = int32(0)
	goto L1
L1260:
	;
	if v3508&int32(_a_F_HeapTupleSatisfiesVisibility_2) != 0 {
		goto L1261
	} else {
		goto L1262
	}
L1261:
	;
	v3521 = F_HeapTupleGetUpdateXid(m, v3416)
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L64
	} else {
		goto L1264
	}
L1262:
	;
	v3523 = v3429
	goto L1263
L1263:
	;
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3523
	if v3524 == int32(0) {
		goto L1265
	} else {
		goto L1266
	}
L1264:
	;
	v3523 = v3521
	goto L1263
L1265:
	;
	v3553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3553))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3523)) == int32(0) {
		goto L1273
	} else {
		goto L1274
	}
L1266:
	;
	v3530 = v13 + int32(12)
	v3533 = F_bsearch(m, v3530, v3525, v3524, int32(4), int32(185))
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L64
	} else {
		goto L1267
	}
L1267:
	;
	if v3533 == int32(0) {
		goto L1265
	} else {
		goto L1268
	}
L1268:
	;
	v3537 = *(*int32)(unsafe.Add(mBase, uint32(v3416)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v3537
	v3540 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[5]))
	v3543 = F_ResolveCminCmaxDuringDecoding(m, v3540, l1, l0, l2, v3530, v13+int32(8))
	mBase = m.M
	v3544 = m.ExcPending
	if v3544 != 0 {
		goto L64
	} else {
		goto L1269
	}
L1269:
	;
	if v3543 == int32(0) {
		v3629 = v3507
		goto L1
	} else {
		goto L1270
	}
L1270:
	;
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v3547 == int32(-1) {
		v3629 = v3507
		goto L1
	} else {
		goto L1271
	}
L1271:
	;
	v3550 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3629 = base.B2i32(base.Ui32(v3550) <= base.Ui32(v3547))
	goto L1
L1272:
	;
	if v3565 != 0 {
		goto L1276
	} else {
		goto L1277
	}
L1273:
	;
	v3565 = base.B2i32(base.Ui32(v3523) < base.Ui32(v3553))
	goto L1272
L1274:
	;
	goto L1275
L1275:
	;
	v3565 = int32(base.Ui32(v3523-v3553) >> (uint(int32(31)) % 32))
	goto L1272
L1276:
	;
	v3567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3416)+21)))
	if v3567&int32(4) != 0 {
		v3629 = int32(0)
		goto L1
	} else {
		goto L1279
	}
L1277:
	;
	goto L1278
L1278:
	;
	v3574 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3574))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3523)) == int32(0) {
		goto L1282
	} else {
		goto L1283
	}
L1279:
	;
	v3570 = F_TransactionIdDidCommit(m, v3523)
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L64
	} else {
		goto L1280
	}
L1280:
	;
	v3629 = v3570 ^ int32(1)
	goto L1
L1281:
	;
	if v3586 != 0 {
		v3629 = v3507
		goto L1
	} else {
		goto L1285
	}
L1282:
	;
	v3586 = base.B2i32(base.Ui32(v3574) <= base.Ui32(v3523))
	goto L1281
L1283:
	;
	goto L1284
L1284:
	;
	v3586 = base.B2i32(int32(0) <= v3523-v3574)
	goto L1281
L1285:
	;
	v3587 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3523
	if v3587 == int32(0) {
		v3629 = v3507
		goto L1
	} else {
		goto L1286
	}
L1286:
	;
	v3596 = F_bsearch(m, v13+int32(12), v3588, v3587, int32(4), int32(185))
	mBase = m.M
	v3597 = m.ExcPending
	if v3597 != 0 {
		goto L64
	} else {
		goto L1287
	}
L1287:
	;
	v3629 = base.B2i32(v3596 == int32(0))
	goto L1
L1288:
	;
	if v3603 == int32(2) {
		goto L1289
	} else {
		goto L1290
	}
L1289:
	;
	v3609 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v3610 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v3611 = F_GlobalVisTestIsRemovableXid(m, v3609, v3610)
	mBase = m.M
	v3612 = m.ExcPending
	if v3612 != 0 {
		goto L64
	} else {
		goto L1292
	}
L1290:
	;
	v3614 = v3603
	goto L1291
L1291:
	;
	v3629 = base.B2i32(v3614 != int32(0))
	goto L1
L1292:
	;
	if v3611 != 0 {
		goto L1293
	} else {
		goto L1294
	}
L1293:
	;
	v3613 = int32(0)
	goto L1295
L1294:
	;
	v3613 = int32(2)
	goto L1295
L1295:
	;
	v3614 = v3613
	goto L1291
L1296:
	;
	v3629 = v3622
	goto L1
}
