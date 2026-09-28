package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_arabic_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1835 int32
	_ = v1835
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1914 int32
	_ = v1914
	var v1918 int32
	_ = v1918
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
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1962 int32
	_ = v1962
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1994 int32
	_ = v1994
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
	var v2013 int32
	_ = v2013
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2063 int32
	_ = v2063
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2115 int32
	_ = v2115
	var v2120 int32
	_ = v2120
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2189 int32
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2237 int32
	_ = v2237
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2246 int32
	_ = v2246
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2277 int32
	_ = v2277
	var v2282 int32
	_ = v2282
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2299 int32
	_ = v2299
	var v2303 int32
	_ = v2303
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2351 int32
	_ = v2351
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2368 int32
	_ = v2368
	var v2373 int32
	_ = v2373
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2442 int32
	_ = v2442
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2459 int32
	_ = v2459
	var v2464 int32
	_ = v2464
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2496 int32
	_ = v2496
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2533 int32
	_ = v2533
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2550 int32
	_ = v2550
	var v2555 int32
	_ = v2555
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2576 int32
	_ = v2576
	var v2580 int32
	_ = v2580
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2610 int32
	_ = v2610
	var v2613 int32
	_ = v2613
	var v2624 int32
	_ = v2624
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2636 int32
	_ = v2636
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2654 int32
	_ = v2654
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2664 int32
	_ = v2664
	var v2669 int32
	_ = v2669
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2690 int32
	_ = v2690
	var v2694 int32
	_ = v2694
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2738 int32
	_ = v2738
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2744 int32
	_ = v2744
	var v2746 int32
	_ = v2746
	var v2750 int32
	_ = v2750
	var v2752 int32
	_ = v2752
	var v2756 int32
	_ = v2756
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2773 int32
	_ = v2773
	var v2775 int32
	_ = v2775
	var v2784 int32
	_ = v2784
	var v2785 int32
	_ = v2785
	var v2788 int32
	_ = v2788
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2800 int32
	_ = v2800
	var v2805 int32
	_ = v2805
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2822 int32
	_ = v2822
	var v2826 int32
	_ = v2826
	var v2830 int32
	_ = v2830
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2837 int32
	_ = v2837
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2874 int32
	_ = v2874
	var v2877 int32
	_ = v2877
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2888 int32
	_ = v2888
	var v2893 int32
	_ = v2893
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2910 int32
	_ = v2910
	var v2914 int32
	_ = v2914
	var v2918 int32
	_ = v2918
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2925 int32
	_ = v2925
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2948 int32
	_ = v2948
	var v2951 int32
	_ = v2951
	var v2962 int32
	_ = v2962
	var v2965 int32
	_ = v2965
	var v2971 int32
	_ = v2971
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2980 int32
	_ = v2980
	var v2981 int32
	_ = v2981
	var v2985 int32
	_ = v2985
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2995 int32
	_ = v2995
	var v2997 int32
	_ = v2997
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3010 int32
	_ = v3010
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3022 int32
	_ = v3022
	var v3027 int32
	_ = v3027
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3038 int32
	_ = v3038
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3048 int32
	_ = v3048
	var v3052 int32
	_ = v3052
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3085 int32
	_ = v3085
	var v3096 int32
	_ = v3096
	var v3099 int32
	_ = v3099
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3110 int32
	_ = v3110
	var v3115 int32
	_ = v3115
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3126 int32
	_ = v3126
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3132 int32
	_ = v3132
	var v3136 int32
	_ = v3136
	var v3140 int32
	_ = v3140
	var v3143 int32
	_ = v3143
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3170 int32
	_ = v3170
	var v3173 int32
	_ = v3173
	var v3184 int32
	_ = v3184
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3201 int32
	_ = v3201
	var v3206 int32
	_ = v3206
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3227 int32
	_ = v3227
	var v3231 int32
	_ = v3231
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3238 int32
	_ = v3238
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3264 int32
	_ = v3264
	var v3275 int32
	_ = v3275
	var v3280 int32
	_ = v3280
	var v3281 int32
	_ = v3281
	var v3287 int32
	_ = v3287
	var v3289 int32
	_ = v3289
	var v3291 int32
	_ = v3291
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3304 int32
	_ = v3304
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3317 int32
	_ = v3317
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3329 int32
	_ = v3329
	var v3334 int32
	_ = v3334
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3345 int32
	_ = v3345
	var v3347 int32
	_ = v3347
	var v3348 int32
	_ = v3348
	var v3351 int32
	_ = v3351
	var v3355 int32
	_ = v3355
	var v3359 int32
	_ = v3359
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3366 int32
	_ = v3366
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3381 int32
	_ = v3381
	var v3382 int32
	_ = v3382
	var v3385 int32
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3389 int32
	_ = v3389
	var v3392 int32
	_ = v3392
	var v3403 int32
	_ = v3403
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3420 int32
	_ = v3420
	var v3425 int32
	_ = v3425
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3436 int32
	_ = v3436
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3442 int32
	_ = v3442
	var v3446 int32
	_ = v3446
	var v3450 int32
	_ = v3450
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3457 int32
	_ = v3457
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3467 int32
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3480 int32
	_ = v3480
	var v3483 int32
	_ = v3483
	var v3494 int32
	_ = v3494
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3511 int32
	_ = v3511
	var v3516 int32
	_ = v3516
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3527 int32
	_ = v3527
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3541 int32
	_ = v3541
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3548 int32
	_ = v3548
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3567 int32
	_ = v3567
	var v3568 int32
	_ = v3568
	var v3571 int32
	_ = v3571
	var v3574 int32
	_ = v3574
	var v3585 int32
	_ = v3585
	var v3590 int32
	_ = v3590
	var v3591 int32
	_ = v3591
	var v3594 int32
	_ = v3594
	var v3595 int32
	_ = v3595
	var v3602 int32
	_ = v3602
	var v3607 int32
	_ = v3607
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3618 int32
	_ = v3618
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3624 int32
	_ = v3624
	var v3628 int32
	_ = v3628
	var v3632 int32
	_ = v3632
	var v3635 int32
	_ = v3635
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3639 int32
	_ = v3639
	var v3643 int32
	_ = v3643
	var v3644 int32
	_ = v3644
	var v3649 int32
	_ = v3649
	var v3650 int32
	_ = v3650
	var v3654 int32
	_ = v3654
	var v3655 int32
	_ = v3655
	var v3658 int32
	_ = v3658
	var v3659 int32
	_ = v3659
	var v3662 int32
	_ = v3662
	var v3665 int32
	_ = v3665
	var v3676 int32
	_ = v3676
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3687 int32
	_ = v3687
	var v3693 int32
	_ = v3693
	var v3694 int32
	_ = v3694
	var v3697 int32
	_ = v3697
	var v3698 int32
	_ = v3698
	var v3700 int32
	_ = v3700
	var v3702 int32
	_ = v3702
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3712 int32
	_ = v3712
	var v3714 int32
	_ = v3714
	var v3715 int32
	_ = v3715
	var v3722 int32
	_ = v3722
	var v3727 int32
	_ = v3727
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3738 int32
	_ = v3738
	var v3740 int32
	_ = v3740
	var v3741 int32
	_ = v3741
	var v3744 int32
	_ = v3744
	var v3748 int32
	_ = v3748
	var v3752 int32
	_ = v3752
	var v3755 int32
	_ = v3755
	var v3756 int32
	_ = v3756
	var v3757 int32
	_ = v3757
	var v3759 int32
	_ = v3759
	var v3763 int32
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3769 int32
	_ = v3769
	var v3770 int32
	_ = v3770
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3778 int32
	_ = v3778
	var v3779 int32
	_ = v3779
	var v3782 int32
	_ = v3782
	var v3785 int32
	_ = v3785
	var v3796 int32
	_ = v3796
	var v3799 int32
	_ = v3799
	var v3804 int32
	_ = v3804
	var v3805 int32
	_ = v3805
	var v3808 int32
	_ = v3808
	var v3810 int32
	_ = v3810
	var v3815 int32
	_ = v3815
	var v3817 int32
	_ = v3817
	var v3824 int32
	_ = v3824
	var v3828 int32
	_ = v3828
	var v3830 int32
	_ = v3830
	var v3832 int32
	_ = v3832
	var v3847 int32
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3851 int32
	_ = v3851
	var v3855 int32
	_ = v3855
	var v3856 int32
	_ = v3856
	var v3862 int32
	_ = v3862
	var v3869 int32
	_ = v3869
	var v3871 int32
	_ = v3871
	var v3872 int32
	_ = v3872
	var v3875 int32
	_ = v3875
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3895 int32
	_ = v3895
	var v3896 int32
	_ = v3896
	var v3904 int32
	_ = v3904
	var v3906 int32
	_ = v3906
	var v3911 int32
	_ = v3911
	var v3913 int32
	_ = v3913
	var v3920 int32
	_ = v3920
	var v3923 int32
	_ = v3923
	var v3927 int32
	_ = v3927
	var v3934 int32
	_ = v3934
	var v3935 int32
	_ = v3935
	var v3949 int32
	_ = v3949
	var v3954 int32
	_ = v3954
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3972 int32
	_ = v3972
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3981 int32
	_ = v3981
	v2 = int32(0)
	v8 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v8)
	v10 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v12
	v15 = v12 + int32(3)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v16 <= v15 {
		v219 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v222 = v12
	goto L45
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v15))))
	if base.B2i32(v20 != int32(167))&base.B2i32(v20 != int32(132)) != 0 {
		v219 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_0), int32(4), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v29 == int32(0) {
		v219 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35
	v37 = int32(1)
	switch v29 - v37 {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		v219 = v37
		goto L1
	}
L7:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v213)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v213)
	v219 = v213
	goto L1
L8:
	;
	v126 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127-int32(4))))
	if v135 == v126 {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(4))))
	if v49 == v40 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if int32(5) <= v123 {
		goto L7
	} else {
		goto L26
	}
L11:
	;
	v123 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v54 = v49 & int32(3)
	if base.Ui32(v49) < base.Ui32(int32(4)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v123 = v112
	goto L10
L15:
	;
	v96 = v90
	v97 = v91
	v101 = v40
	goto L23
L16:
	;
	v90 = v41
	v91 = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v61 = v41
	v62 = int32(0)
	v65 = v40
	goto L19
L19:
	;
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	v68 = int32(-65)
	v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+1)))
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+2)))
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+3)))
	v82 = v62 + base.B2i32(v68 < v67) + base.B2i32(v68 < v71) + base.B2i32(v68 < v75) + base.B2i32(v68 < v79)
	v83 = int32(4)
	v84 = v61 + v83
	v86 = v65 + v83
	if v86 != v49&int32(-4) {
		v61 = v84
		v62 = v82
		v65 = v86
		goto L19
	} else {
		goto L21
	}
L20:
	;
	if v54 == int32(0) {
		v112 = v82
		goto L14
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v90 = v84
	v91 = v82
	goto L15
L23:
	;
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	v105 = v97 + base.B2i32(int32(-65) < v102)
	v106 = int32(1)
	v109 = v101 + v106
	if v109 != v54 {
		v96 = v96 + v106
		v97 = v105
		v101 = v109
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v112 = v105
	goto L14
L25:
	;
	goto L24
L26:
	;
	v219 = v40
	goto L1
L27:
	;
	if v209 < int32(4) {
		v219 = v126
		goto L1
	} else {
		goto L43
	}
L28:
	;
	v209 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v140 = v135 & int32(3)
	if base.Ui32(v135) < base.Ui32(int32(4)) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v209 = v198
	goto L27
L32:
	;
	v182 = v176
	v183 = v177
	v187 = v126
	goto L40
L33:
	;
	v176 = v127
	v177 = int32(0)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v147 = v127
	v148 = int32(0)
	v151 = v126
	goto L36
L36:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147))))
	v154 = int32(-65)
	v157 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+1)))
	v161 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+2)))
	v165 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+3)))
	v168 = v148 + base.B2i32(v154 < v153) + base.B2i32(v154 < v157) + base.B2i32(v154 < v161) + base.B2i32(v154 < v165)
	v169 = int32(4)
	v170 = v147 + v169
	v172 = v151 + v169
	if v172 != v135&int32(-4) {
		v147 = v170
		v148 = v168
		v151 = v172
		goto L36
	} else {
		goto L38
	}
L37:
	;
	if v140 == int32(0) {
		v198 = v168
		goto L31
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v176 = v170
	v177 = v168
	goto L32
L40:
	;
	v188 = int32(*(*int8)(unsafe.Add(mBase, uint32(v182))))
	v191 = v183 + base.B2i32(int32(-65) < v188)
	v192 = int32(1)
	v195 = v187 + v192
	if v195 != v140 {
		v182 = v182 + v192
		v183 = v191
		v187 = v195
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v198 = v191
	goto L31
L42:
	;
	goto L41
L43:
	;
	goto L7
L44:
	;
	return v3981
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v222
	v232 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_1), int32(144), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v604 != 0 {
		goto L233
	} else {
		goto L234
	}
L47:
	;
	goto L46
L48:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = v600
	goto L45
L49:
	;
	if v232 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v234
	switch v232 - int32(1) {
	case 0:
		goto L103
	case 1:
		goto L102
	case 2:
		goto L101
	case 3:
		goto L100
	case 4:
		goto L99
	case 5:
		goto L98
	case 6:
		goto L97
	case 7:
		goto L96
	case 8:
		goto L95
	case 9:
		goto L94
	case 10:
		goto L93
	case 11:
		goto L92
	case 12:
		goto L91
	case 13:
		goto L90
	case 14:
		goto L89
	case 15:
		goto L88
	case 16:
		goto L87
	case 17:
		goto L86
	case 18:
		goto L85
	case 19:
		goto L84
	case 20:
		goto L83
	case 21:
		goto L82
	case 22:
		goto L81
	case 23:
		goto L80
	case 24:
		goto L79
	case 25:
		goto L78
	case 26:
		goto L77
	case 27:
		goto L76
	case 28:
		goto L75
	case 29:
		goto L74
	case 30:
		goto L73
	case 31:
		goto L72
	case 32:
		goto L71
	case 33:
		goto L70
	case 34:
		goto L69
	case 35:
		goto L68
	case 36:
		goto L67
	case 37:
		goto L66
	case 38:
		goto L65
	case 39:
		goto L64
	case 40:
		goto L63
	case 41:
		goto L62
	case 42:
		goto L61
	case 43:
		goto L60
	case 44:
		goto L59
	case 45:
		goto L58
	case 46:
		goto L57
	case 47:
		goto L56
	case 48:
		goto L55
	case 49:
		goto L54
	case 50:
		goto L53
	default:
		goto L48
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L207
L53:
	;
	v537 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_2))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L203
	}
L54:
	;
	v531 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_3))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L201
	}
L55:
	;
	v525 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_4))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L4
	} else {
		goto L199
	}
L56:
	;
	v519 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_5))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L197
	}
L57:
	;
	v513 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_6))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L195
	}
L58:
	;
	v507 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_7))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L193
	}
L59:
	;
	v501 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_8))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L191
	}
L60:
	;
	v495 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_9))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L189
	}
L61:
	;
	v489 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_10))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L187
	}
L62:
	;
	v483 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_11))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L185
	}
L63:
	;
	v477 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_12))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L183
	}
L64:
	;
	v471 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_13))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L181
	}
L65:
	;
	v465 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_14))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L179
	}
L66:
	;
	v459 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_15))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L177
	}
L67:
	;
	v453 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_16))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L175
	}
L68:
	;
	v447 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_17))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L173
	}
L69:
	;
	v441 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_18))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L171
	}
L70:
	;
	v435 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_19))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L169
	}
L71:
	;
	v429 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_20))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L167
	}
L72:
	;
	v423 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_21))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L165
	}
L73:
	;
	v417 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_22))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L163
	}
L74:
	;
	v411 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_23))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L161
	}
L75:
	;
	v405 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_24))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L159
	}
L76:
	;
	v399 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_25))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L157
	}
L77:
	;
	v393 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_26))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L155
	}
L78:
	;
	v387 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_27))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L153
	}
L79:
	;
	v381 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_28))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L151
	}
L80:
	;
	v375 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_29))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L149
	}
L81:
	;
	v369 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_30))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L147
	}
L82:
	;
	v363 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_31))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L145
	}
L83:
	;
	v357 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_32))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L143
	}
L84:
	;
	v351 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_33))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L141
	}
L85:
	;
	v345 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_34))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L139
	}
L86:
	;
	v339 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_35))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L137
	}
L87:
	;
	v333 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_36))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L135
	}
L88:
	;
	v327 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_37))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L133
	}
L89:
	;
	v321 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_38))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L131
	}
L90:
	;
	v315 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_39))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L129
	}
L91:
	;
	v309 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_40))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L127
	}
L92:
	;
	v303 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_41))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L125
	}
L93:
	;
	v297 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_42))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L123
	}
L94:
	;
	v291 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_43))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L121
	}
L95:
	;
	v285 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_44))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L119
	}
L96:
	;
	v279 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_45))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L117
	}
L97:
	;
	v273 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_46))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L115
	}
L98:
	;
	v267 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_47))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L113
	}
L99:
	;
	v261 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_48))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L111
	}
L100:
	;
	v255 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_49))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L109
	}
L101:
	;
	v249 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_50))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L107
	}
L102:
	;
	v243 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_51))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	v238 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v238 {
		goto L48
	} else {
		goto L104
	}
L104:
	;
	v3981 = v238
	goto L44
L105:
	;
	if int32(0) <= v243 {
		goto L48
	} else {
		goto L106
	}
L106:
	;
	v3981 = v243
	goto L44
L107:
	;
	if int32(0) <= v249 {
		goto L48
	} else {
		goto L108
	}
L108:
	;
	v3981 = v249
	goto L44
L109:
	;
	if int32(0) <= v255 {
		goto L48
	} else {
		goto L110
	}
L110:
	;
	v3981 = v255
	goto L44
L111:
	;
	if int32(0) <= v261 {
		goto L48
	} else {
		goto L112
	}
L112:
	;
	v3981 = v261
	goto L44
L113:
	;
	if int32(0) <= v267 {
		goto L48
	} else {
		goto L114
	}
L114:
	;
	v3981 = v267
	goto L44
L115:
	;
	if int32(0) <= v273 {
		goto L48
	} else {
		goto L116
	}
L116:
	;
	v3981 = v273
	goto L44
L117:
	;
	if int32(0) <= v279 {
		goto L48
	} else {
		goto L118
	}
L118:
	;
	v3981 = v279
	goto L44
L119:
	;
	if int32(0) <= v285 {
		goto L48
	} else {
		goto L120
	}
L120:
	;
	v3981 = v285
	goto L44
L121:
	;
	if int32(0) <= v291 {
		goto L48
	} else {
		goto L122
	}
L122:
	;
	v3981 = v291
	goto L44
L123:
	;
	if int32(0) <= v297 {
		goto L48
	} else {
		goto L124
	}
L124:
	;
	v3981 = v297
	goto L44
L125:
	;
	if int32(0) <= v303 {
		goto L48
	} else {
		goto L126
	}
L126:
	;
	v3981 = v303
	goto L44
L127:
	;
	if int32(0) <= v309 {
		goto L48
	} else {
		goto L128
	}
L128:
	;
	v3981 = v309
	goto L44
L129:
	;
	if int32(0) <= v315 {
		goto L48
	} else {
		goto L130
	}
L130:
	;
	v3981 = v315
	goto L44
L131:
	;
	if int32(0) <= v321 {
		goto L48
	} else {
		goto L132
	}
L132:
	;
	v3981 = v321
	goto L44
L133:
	;
	if int32(0) <= v327 {
		goto L48
	} else {
		goto L134
	}
L134:
	;
	v3981 = v327
	goto L44
L135:
	;
	if int32(0) <= v333 {
		goto L48
	} else {
		goto L136
	}
L136:
	;
	v3981 = v333
	goto L44
L137:
	;
	if int32(0) <= v339 {
		goto L48
	} else {
		goto L138
	}
L138:
	;
	v3981 = v339
	goto L44
L139:
	;
	if int32(0) <= v345 {
		goto L48
	} else {
		goto L140
	}
L140:
	;
	v3981 = v345
	goto L44
L141:
	;
	if int32(0) <= v351 {
		goto L48
	} else {
		goto L142
	}
L142:
	;
	v3981 = v351
	goto L44
L143:
	;
	if int32(0) <= v357 {
		goto L48
	} else {
		goto L144
	}
L144:
	;
	v3981 = v357
	goto L44
L145:
	;
	if int32(0) <= v363 {
		goto L48
	} else {
		goto L146
	}
L146:
	;
	v3981 = v363
	goto L44
L147:
	;
	if int32(0) <= v369 {
		goto L48
	} else {
		goto L148
	}
L148:
	;
	v3981 = v369
	goto L44
L149:
	;
	if int32(0) <= v375 {
		goto L48
	} else {
		goto L150
	}
L150:
	;
	v3981 = v375
	goto L44
L151:
	;
	if int32(0) <= v381 {
		goto L48
	} else {
		goto L152
	}
L152:
	;
	v3981 = v381
	goto L44
L153:
	;
	if int32(0) <= v387 {
		goto L48
	} else {
		goto L154
	}
L154:
	;
	v3981 = v387
	goto L44
L155:
	;
	if int32(0) <= v393 {
		goto L48
	} else {
		goto L156
	}
L156:
	;
	v3981 = v393
	goto L44
L157:
	;
	if int32(0) <= v399 {
		goto L48
	} else {
		goto L158
	}
L158:
	;
	v3981 = v399
	goto L44
L159:
	;
	if int32(0) <= v405 {
		goto L48
	} else {
		goto L160
	}
L160:
	;
	v3981 = v405
	goto L44
L161:
	;
	if int32(0) <= v411 {
		goto L48
	} else {
		goto L162
	}
L162:
	;
	v3981 = v411
	goto L44
L163:
	;
	if int32(0) <= v417 {
		goto L48
	} else {
		goto L164
	}
L164:
	;
	v3981 = v417
	goto L44
L165:
	;
	if int32(0) <= v423 {
		goto L48
	} else {
		goto L166
	}
L166:
	;
	v3981 = v423
	goto L44
L167:
	;
	if int32(0) <= v429 {
		goto L48
	} else {
		goto L168
	}
L168:
	;
	v3981 = v429
	goto L44
L169:
	;
	if int32(0) <= v435 {
		goto L48
	} else {
		goto L170
	}
L170:
	;
	v3981 = v435
	goto L44
L171:
	;
	if int32(0) <= v441 {
		goto L48
	} else {
		goto L172
	}
L172:
	;
	v3981 = v441
	goto L44
L173:
	;
	if int32(0) <= v447 {
		goto L48
	} else {
		goto L174
	}
L174:
	;
	v3981 = v447
	goto L44
L175:
	;
	if int32(0) <= v453 {
		goto L48
	} else {
		goto L176
	}
L176:
	;
	v3981 = v453
	goto L44
L177:
	;
	if int32(0) <= v459 {
		goto L48
	} else {
		goto L178
	}
L178:
	;
	v3981 = v459
	goto L44
L179:
	;
	if int32(0) <= v465 {
		goto L48
	} else {
		goto L180
	}
L180:
	;
	v3981 = v465
	goto L44
L181:
	;
	if int32(0) <= v471 {
		goto L48
	} else {
		goto L182
	}
L182:
	;
	v3981 = v471
	goto L44
L183:
	;
	if int32(0) <= v477 {
		goto L48
	} else {
		goto L184
	}
L184:
	;
	v3981 = v477
	goto L44
L185:
	;
	if int32(0) <= v483 {
		goto L48
	} else {
		goto L186
	}
L186:
	;
	v3981 = v483
	goto L44
L187:
	;
	if int32(0) <= v489 {
		goto L48
	} else {
		goto L188
	}
L188:
	;
	v3981 = v489
	goto L44
L189:
	;
	if int32(0) <= v495 {
		goto L48
	} else {
		goto L190
	}
L190:
	;
	v3981 = v495
	goto L44
L191:
	;
	if int32(0) <= v501 {
		goto L48
	} else {
		goto L192
	}
L192:
	;
	v3981 = v501
	goto L44
L193:
	;
	if int32(0) <= v507 {
		goto L48
	} else {
		goto L194
	}
L194:
	;
	v3981 = v507
	goto L44
L195:
	;
	if int32(0) <= v513 {
		goto L48
	} else {
		goto L196
	}
L196:
	;
	v3981 = v513
	goto L44
L197:
	;
	if int32(0) <= v519 {
		goto L48
	} else {
		goto L198
	}
L198:
	;
	v3981 = v519
	goto L44
L199:
	;
	if int32(0) <= v525 {
		goto L48
	} else {
		goto L200
	}
L200:
	;
	v3981 = v525
	goto L44
L201:
	;
	if int32(0) <= v531 {
		goto L48
	} else {
		goto L202
	}
L202:
	;
	v3981 = v531
	goto L44
L203:
	;
	if int32(0) <= v537 {
		goto L48
	} else {
		goto L204
	}
L204:
	;
	v3981 = v537
	goto L44
L205:
	;
	if v595 < int32(0) {
		goto L47
	} else {
		goto L225
	}
L207:
	;
	goto L208
L208:
	;
	goto L209
L209:
	;
	v550 = v222
	v552 = int32(1)
	goto L212
L211:
	;
	v595 = v580
	goto L205
L212:
	;
	if v543 <= v550 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L211
L214:
	;
	v595 = int32(-1)
	goto L205
L215:
	;
	goto L216
L216:
	;
	v557 = v550 + int32(1)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542+v550))))
	if base.Ui32(v559) < base.Ui32(int32(192)) {
		v580 = v557
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v581 = int32(1)
	if v581 < v552 {
		v550 = v580
		v552 = v552 - v581
		goto L212
	} else {
		goto L224
	}
L218:
	;
	if v543 <= v557 {
		v580 = v557
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v566 = v557
	goto L220
L220:
	;
	v569 = int32(*(*int8)(unsafe.Add(mBase, uint32(v542+v566))))
	if int32(-65) < v569 {
		v580 = v566
		goto L217
	} else {
		goto L222
	}
L221:
	;
	v580 = v543
	goto L217
L222:
	;
	v573 = v566 + int32(1)
	if v573 != v543 {
		v566 = v573
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	goto L213
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v595
	goto L48
L226:
	;
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2237
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2237
	v2241 = v2237 + int32(3)
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2242 <= v2241 {
		goto L688
	} else {
		goto L689
	}
L227:
	;
	v2226 = F_slice_del(m, l0)
	mBase = m.M
	v2228 = base.B2i32(v2226 < int32(0))
	if v2226 < int32(0) {
		goto L683
	} else {
		goto L684
	}
L228:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2197
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2197
	v2200 = int32(2)
	v2202 = int32(0)
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2197-v2205 < v2200 {
		v2215 = v2202
		goto L677
	} else {
		goto L678
	}
L229:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2087
	v2089 = int32(2)
	v2091 = int32(0)
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2087-v2094 < v2089 {
		v2104 = v2091
		goto L653
	} else {
		goto L654
	}
L230:
	;
	if v2080 != 0 {
		v3981 = v2078
		goto L44
	} else {
		goto L651
	}
L231:
	;
	if int32(0) <= v1316 {
		v2232 = v1195
		goto L226
	} else {
		goto L650
	}
L232:
	;
	if int32(0) <= v919 {
		v2232 = v924
		goto L226
	} else {
		goto L649
	}
L233:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v605
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v608 = int32(1)
	v612 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_52), int32(12), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L4
	} else {
		goto L236
	}
L234:
	;
	v1320 = v602
	v1321 = v219
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1320
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v1327 == int32(0) {
		v2193 = v1321
		goto L228
	} else {
		goto L409
	}
L236:
	;
	if v612 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v615 = v612
	v617 = v605
	v618 = v607
	v619 = v608
	goto L240
L238:
	;
	v909 = v605
	v910 = v607
	v911 = v608
	goto L239
L239:
	;
	v913 = v909 - v910
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v913 + v914
	if v911 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L240:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v621
	switch v615 - int32(1) {
	case 0:
		goto L246
	case 1:
		goto L245
	case 2:
		goto L244
	default:
		goto L243
	}
L241:
	;
	v909 = v901
	v910 = v902
	v911 = base.B2i32(int32(0) < v903)
	goto L239
L242:
	;
	goto L241
L243:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v890
	v893 = v619 - int32(1)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v898 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_52), int32(12), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L4
	} else {
		goto L301
	}
L244:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v802 = int32(0)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v801-int32(4))))
	if v809 == v802 {
		goto L284
	} else {
		goto L285
	}
L245:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v714 = int32(0)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v713-int32(4))))
	if v721 == v714 {
		goto L266
	} else {
		goto L267
	}
L246:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v626 = int32(0)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v625-int32(4))))
	if v633 == v626 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	if v707 < int32(4) {
		v901 = v617
		v902 = v618
		v903 = v619
		goto L242
	} else {
		goto L263
	}
L248:
	;
	v707 = int32(0)
	goto L247
L249:
	;
	goto L250
L250:
	;
	v638 = v633 & int32(3)
	if base.Ui32(v633) < base.Ui32(int32(4)) {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	v707 = v696
	goto L247
L252:
	;
	v680 = v674
	v681 = v675
	v685 = v626
	goto L260
L253:
	;
	v674 = v625
	v675 = int32(0)
	goto L252
L254:
	;
	goto L255
L255:
	;
	v645 = v625
	v646 = int32(0)
	v649 = v626
	goto L256
L256:
	;
	v651 = int32(*(*int8)(unsafe.Add(mBase, uint32(v645))))
	v652 = int32(-65)
	v655 = int32(*(*int8)(unsafe.Add(mBase, uint32(v645)+1)))
	v659 = int32(*(*int8)(unsafe.Add(mBase, uint32(v645)+2)))
	v663 = int32(*(*int8)(unsafe.Add(mBase, uint32(v645)+3)))
	v666 = v646 + base.B2i32(v652 < v651) + base.B2i32(v652 < v655) + base.B2i32(v652 < v659) + base.B2i32(v652 < v663)
	v667 = int32(4)
	v668 = v645 + v667
	v670 = v649 + v667
	if v670 != v633&int32(-4) {
		v645 = v668
		v646 = v666
		v649 = v670
		goto L256
	} else {
		goto L258
	}
L257:
	;
	if v638 == int32(0) {
		v696 = v666
		goto L251
	} else {
		goto L259
	}
L258:
	;
	goto L257
L259:
	;
	v674 = v668
	v675 = v666
	goto L252
L260:
	;
	v686 = int32(*(*int8)(unsafe.Add(mBase, uint32(v680))))
	v689 = v681 + base.B2i32(int32(-65) < v686)
	v690 = int32(1)
	v693 = v685 + v690
	if v693 != v638 {
		v680 = v680 + v690
		v681 = v689
		v685 = v693
		goto L260
	} else {
		goto L262
	}
L261:
	;
	v696 = v689
	goto L251
L262:
	;
	goto L261
L263:
	;
	v710 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v710 {
		goto L243
	} else {
		goto L264
	}
L264:
	;
	v3981 = v710
	goto L44
L265:
	;
	if v795 < int32(5) {
		v901 = v617
		v902 = v618
		v903 = v619
		goto L242
	} else {
		goto L281
	}
L266:
	;
	v795 = int32(0)
	goto L265
L267:
	;
	goto L268
L268:
	;
	v726 = v721 & int32(3)
	if base.Ui32(v721) < base.Ui32(int32(4)) {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	v795 = v784
	goto L265
L270:
	;
	v768 = v762
	v769 = v763
	v773 = v714
	goto L278
L271:
	;
	v762 = v713
	v763 = int32(0)
	goto L270
L272:
	;
	goto L273
L273:
	;
	v733 = v713
	v734 = int32(0)
	v737 = v714
	goto L274
L274:
	;
	v739 = int32(*(*int8)(unsafe.Add(mBase, uint32(v733))))
	v740 = int32(-65)
	v743 = int32(*(*int8)(unsafe.Add(mBase, uint32(v733)+1)))
	v747 = int32(*(*int8)(unsafe.Add(mBase, uint32(v733)+2)))
	v751 = int32(*(*int8)(unsafe.Add(mBase, uint32(v733)+3)))
	v754 = v734 + base.B2i32(v740 < v739) + base.B2i32(v740 < v743) + base.B2i32(v740 < v747) + base.B2i32(v740 < v751)
	v755 = int32(4)
	v756 = v733 + v755
	v758 = v737 + v755
	if v758 != v721&int32(-4) {
		v733 = v756
		v734 = v754
		v737 = v758
		goto L274
	} else {
		goto L276
	}
L275:
	;
	if v726 == int32(0) {
		v784 = v754
		goto L269
	} else {
		goto L277
	}
L276:
	;
	goto L275
L277:
	;
	v762 = v756
	v763 = v754
	goto L270
L278:
	;
	v774 = int32(*(*int8)(unsafe.Add(mBase, uint32(v768))))
	v777 = v769 + base.B2i32(int32(-65) < v774)
	v778 = int32(1)
	v781 = v773 + v778
	if v781 != v726 {
		v768 = v768 + v778
		v769 = v777
		v773 = v781
		goto L278
	} else {
		goto L280
	}
L279:
	;
	v784 = v777
	goto L269
L280:
	;
	goto L279
L281:
	;
	v798 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v798 {
		goto L243
	} else {
		goto L282
	}
L282:
	;
	v3981 = v798
	goto L44
L283:
	;
	if v883 < int32(6) {
		v901 = v617
		v902 = v618
		v903 = v619
		goto L242
	} else {
		goto L299
	}
L284:
	;
	v883 = int32(0)
	goto L283
L285:
	;
	goto L286
L286:
	;
	v814 = v809 & int32(3)
	if base.Ui32(v809) < base.Ui32(int32(4)) {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v883 = v872
	goto L283
L288:
	;
	v856 = v850
	v857 = v851
	v861 = v802
	goto L296
L289:
	;
	v850 = v801
	v851 = int32(0)
	goto L288
L290:
	;
	goto L291
L291:
	;
	v821 = v801
	v822 = int32(0)
	v825 = v802
	goto L292
L292:
	;
	v827 = int32(*(*int8)(unsafe.Add(mBase, uint32(v821))))
	v828 = int32(-65)
	v831 = int32(*(*int8)(unsafe.Add(mBase, uint32(v821)+1)))
	v835 = int32(*(*int8)(unsafe.Add(mBase, uint32(v821)+2)))
	v839 = int32(*(*int8)(unsafe.Add(mBase, uint32(v821)+3)))
	v842 = v822 + base.B2i32(v828 < v827) + base.B2i32(v828 < v831) + base.B2i32(v828 < v835) + base.B2i32(v828 < v839)
	v843 = int32(4)
	v844 = v821 + v843
	v846 = v825 + v843
	if v846 != v809&int32(-4) {
		v821 = v844
		v822 = v842
		v825 = v846
		goto L292
	} else {
		goto L294
	}
L293:
	;
	if v814 == int32(0) {
		v872 = v842
		goto L287
	} else {
		goto L295
	}
L294:
	;
	goto L293
L295:
	;
	v850 = v844
	v851 = v842
	goto L288
L296:
	;
	v862 = int32(*(*int8)(unsafe.Add(mBase, uint32(v856))))
	v865 = v857 + base.B2i32(int32(-65) < v862)
	v866 = int32(1)
	v869 = v861 + v866
	if v869 != v814 {
		v856 = v856 + v866
		v857 = v865
		v861 = v869
		goto L296
	} else {
		goto L298
	}
L297:
	;
	v872 = v865
	goto L287
L298:
	;
	goto L297
L299:
	;
	v886 = F_slice_del(m, l0)
	mBase = m.M
	if v886 < int32(0) {
		v3981 = v886
		goto L44
	} else {
		goto L300
	}
L300:
	;
	goto L243
L301:
	;
	if v898 != 0 {
		v615 = v898
		v617 = v890
		v618 = v894
		v619 = v893
		goto L240
	} else {
		goto L302
	}
L302:
	;
	v901 = v890
	v902 = v894
	v903 = v893
	goto L242
L303:
	;
	v919 = F_r_Suffix_Verb_Step2a(m, l0)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L4
	} else {
		goto L306
	}
L304:
	;
	v1194 = v914
	v1195 = v219
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1194
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1194
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1194-int32(3) <= v1199 {
		goto L381
	} else {
		goto L382
	}
L306:
	;
	if v919 < int32(0) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v923 = v919
	goto L309
L308:
	;
	v923 = v219
	goto L309
L309:
	;
	if v919 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v924 = v923
	goto L312
L311:
	;
	v924 = v219
	goto L312
L312:
	;
	if v919 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v928 = int32(base.Ui32(v919) >> (uint(int32(31)) % 32))
	goto L315
L314:
	;
	v928 = int32(13)
	goto L315
L315:
	;
	if v928 == int32(0) {
		v2232 = v924
		goto L226
	} else {
		goto L316
	}
L316:
	;
	if v928 != int32(13) {
		goto L232
	} else {
		goto L317
	}
L317:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v934 = v933 + v913
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v934
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v934
	v938 = v934 - int32(1)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v938 <= v939 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1135 = v1134 + v913
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1135
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L363
L319:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941+v938))))
	if v943 != int32(136) {
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v949 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_53), int32(2), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L4
	} else {
		goto L321
	}
L321:
	;
	if v949 == int32(0) {
		goto L318
	} else {
		goto L322
	}
L322:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v953
	switch v949 - int32(1) {
	case 0:
		goto L324
	case 1:
		goto L323
	default:
		v2232 = v924
		goto L226
	}
L323:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1046 = int32(0)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1045-int32(4))))
	if v1053 == v1046 {
		goto L344
	} else {
		goto L345
	}
L324:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v958 = int32(0)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v957-int32(4))))
	if v965 == v958 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	if v1039 < int32(4) {
		goto L318
	} else {
		goto L341
	}
L326:
	;
	v1039 = int32(0)
	goto L325
L327:
	;
	goto L328
L328:
	;
	v970 = v965 & int32(3)
	if base.Ui32(v965) < base.Ui32(int32(4)) {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	v1039 = v1028
	goto L325
L330:
	;
	v1012 = v1006
	v1013 = v1007
	v1017 = v958
	goto L338
L331:
	;
	v1006 = v957
	v1007 = int32(0)
	goto L330
L332:
	;
	goto L333
L333:
	;
	v977 = v957
	v978 = int32(0)
	v981 = v958
	goto L334
L334:
	;
	v983 = int32(*(*int8)(unsafe.Add(mBase, uint32(v977))))
	v984 = int32(-65)
	v987 = int32(*(*int8)(unsafe.Add(mBase, uint32(v977)+1)))
	v991 = int32(*(*int8)(unsafe.Add(mBase, uint32(v977)+2)))
	v995 = int32(*(*int8)(unsafe.Add(mBase, uint32(v977)+3)))
	v998 = v978 + base.B2i32(v984 < v983) + base.B2i32(v984 < v987) + base.B2i32(v984 < v991) + base.B2i32(v984 < v995)
	v999 = int32(4)
	v1000 = v977 + v999
	v1002 = v981 + v999
	if v1002 != v965&int32(-4) {
		v977 = v1000
		v978 = v998
		v981 = v1002
		goto L334
	} else {
		goto L336
	}
L335:
	;
	if v970 == int32(0) {
		v1028 = v998
		goto L329
	} else {
		goto L337
	}
L336:
	;
	goto L335
L337:
	;
	v1006 = v1000
	v1007 = v998
	goto L330
L338:
	;
	v1018 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1012))))
	v1021 = v1013 + base.B2i32(int32(-65) < v1018)
	v1022 = int32(1)
	v1025 = v1017 + v1022
	if v1025 != v970 {
		v1012 = v1012 + v1022
		v1013 = v1021
		v1017 = v1025
		goto L338
	} else {
		goto L340
	}
L339:
	;
	v1028 = v1021
	goto L329
L340:
	;
	goto L339
L341:
	;
	v1042 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1042 {
		v2232 = v924
		goto L226
	} else {
		goto L342
	}
L342:
	;
	v3981 = v1042
	goto L44
L343:
	;
	if v1127 < int32(6) {
		goto L318
	} else {
		goto L359
	}
L344:
	;
	v1127 = int32(0)
	goto L343
L345:
	;
	goto L346
L346:
	;
	v1058 = v1053 & int32(3)
	if base.Ui32(v1053) < base.Ui32(int32(4)) {
		goto L349
	} else {
		goto L350
	}
L347:
	;
	v1127 = v1116
	goto L343
L348:
	;
	v1100 = v1094
	v1101 = v1095
	v1105 = v1046
	goto L356
L349:
	;
	v1094 = v1045
	v1095 = int32(0)
	goto L348
L350:
	;
	goto L351
L351:
	;
	v1065 = v1045
	v1066 = int32(0)
	v1069 = v1046
	goto L352
L352:
	;
	v1071 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1065))))
	v1072 = int32(-65)
	v1075 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1065)+1)))
	v1079 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1065)+2)))
	v1083 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1065)+3)))
	v1086 = v1066 + base.B2i32(v1072 < v1071) + base.B2i32(v1072 < v1075) + base.B2i32(v1072 < v1079) + base.B2i32(v1072 < v1083)
	v1087 = int32(4)
	v1088 = v1065 + v1087
	v1090 = v1069 + v1087
	if v1090 != v1053&int32(-4) {
		v1065 = v1088
		v1066 = v1086
		v1069 = v1090
		goto L352
	} else {
		goto L354
	}
L353:
	;
	if v1058 == int32(0) {
		v1116 = v1086
		goto L347
	} else {
		goto L355
	}
L354:
	;
	goto L353
L355:
	;
	v1094 = v1088
	v1095 = v1086
	goto L348
L356:
	;
	v1106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1100))))
	v1109 = v1101 + base.B2i32(int32(-65) < v1106)
	v1110 = int32(1)
	v1113 = v1105 + v1110
	if v1113 != v1058 {
		v1100 = v1100 + v1110
		v1101 = v1109
		v1105 = v1113
		goto L356
	} else {
		goto L358
	}
L357:
	;
	v1116 = v1109
	goto L347
L358:
	;
	goto L357
L359:
	;
	v1130 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1130 {
		v2232 = v924
		goto L226
	} else {
		goto L360
	}
L360:
	;
	v3981 = v1130
	goto L44
L361:
	;
	if int32(0) <= v1190 {
		v2232 = v924
		goto L226
	} else {
		goto L380
	}
L363:
	;
	goto L364
L364:
	;
	goto L365
L365:
	;
	v1145 = v1135
	v1147 = int32(1)
	goto L368
L367:
	;
	v1190 = v1172
	goto L361
L368:
	;
	if v1145 <= v1138 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	goto L367
L370:
	;
	v1190 = int32(-1)
	goto L361
L371:
	;
	goto L372
L372:
	;
	v1152 = v1145 - int32(1)
	v1154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1137+v1152))))
	if base.B2i32(int32(0) <= v1154)|base.B2i32(v1152 <= v1138) != 0 {
		v1172 = v1152
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1176 = int32(1)
	if v1176 < v1147 {
		v1145 = v1172
		v1147 = v1147 - v1176
		goto L368
	} else {
		goto L379
	}
L374:
	;
	v1160 = v1152
	goto L375
L375:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+v1160))))
	if base.Ui32(int32(191)) < base.Ui32(v1165) {
		v1172 = v1160
		goto L373
	} else {
		goto L377
	}
L376:
	;
	v1172 = v1138
	goto L373
L377:
	;
	v1169 = v1160 - int32(1)
	if v1138 < v1169 {
		v1160 = v1169
		goto L375
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	goto L369
L380:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1194 = v1193
	v1195 = v924
	goto L305
L381:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1314
	v1316 = F_r_Suffix_Verb_Step2a(m, l0)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L4
	} else {
		goto L407
	}
L382:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203+v1194-int32(1)))))
	if base.B2i32(v1207 != int32(167))&base.B2i32(v1207 != int32(133)) != 0 {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1216 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_54), int32(2), int32(0))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	if v1216 == int32(0) {
		goto L381
	} else {
		goto L385
	}
L385:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1220
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1223 = int32(0)
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1222-int32(4))))
	if v1230 == v1223 {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	if v1304 < int32(5) {
		goto L381
	} else {
		goto L402
	}
L387:
	;
	v1304 = int32(0)
	goto L386
L388:
	;
	goto L389
L389:
	;
	v1235 = v1230 & int32(3)
	if base.Ui32(v1230) < base.Ui32(int32(4)) {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	v1304 = v1293
	goto L386
L391:
	;
	v1277 = v1271
	v1278 = v1272
	v1282 = v1223
	goto L399
L392:
	;
	v1271 = v1222
	v1272 = int32(0)
	goto L391
L393:
	;
	goto L394
L394:
	;
	v1242 = v1222
	v1243 = int32(0)
	v1246 = v1223
	goto L395
L395:
	;
	v1248 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1242))))
	v1249 = int32(-65)
	v1252 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1242)+1)))
	v1256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1242)+2)))
	v1260 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1242)+3)))
	v1263 = v1243 + base.B2i32(v1249 < v1248) + base.B2i32(v1249 < v1252) + base.B2i32(v1249 < v1256) + base.B2i32(v1249 < v1260)
	v1264 = int32(4)
	v1265 = v1242 + v1264
	v1267 = v1246 + v1264
	if v1267 != v1230&int32(-4) {
		v1242 = v1265
		v1243 = v1263
		v1246 = v1267
		goto L395
	} else {
		goto L397
	}
L396:
	;
	if v1235 == int32(0) {
		v1293 = v1263
		goto L390
	} else {
		goto L398
	}
L397:
	;
	goto L396
L398:
	;
	v1271 = v1265
	v1272 = v1263
	goto L391
L399:
	;
	v1283 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1277))))
	v1286 = v1278 + base.B2i32(int32(-65) < v1283)
	v1287 = int32(1)
	v1290 = v1282 + v1287
	if v1290 != v1235 {
		v1277 = v1277 + v1287
		v1278 = v1286
		v1282 = v1290
		goto L399
	} else {
		goto L401
	}
L400:
	;
	v1293 = v1286
	goto L390
L401:
	;
	goto L400
L402:
	;
	v1307 = F_slice_del(m, l0)
	mBase = m.M
	if v1307 < int32(0) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1310 = v1307
	goto L405
L404:
	;
	v1310 = v1195
	goto L405
L405:
	;
	if int32(0) <= v1307 {
		v2232 = v1310
		goto L226
	} else {
		goto L406
	}
L406:
	;
	v3981 = v1310
	goto L44
L407:
	;
	if v1316 != 0 {
		goto L231
	} else {
		goto L408
	}
L408:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1320 = v1318
	v1321 = v1195
	goto L235
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1320
	v1331 = int32(2)
	v1333 = int32(0)
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1335-v1336 < v1331 {
		v1346 = v1333
		goto L413
	} else {
		goto L414
	}
L410:
	;
	if v2071 != 0 {
		v3981 = v2069
		goto L44
	} else {
		goto L648
	}
L411:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1444
	v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v1446 != 0 {
		goto L441
	} else {
		goto L442
	}
L412:
	;
	if v1346 == int32(0) {
		goto L411
	} else {
		goto L416
	}
L413:
	;
	goto L412
L414:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1342 = F_memcmp(m, v1339+v1335-v1331, int32(_a_F_arabic_UTF_8_stem_55), v1331)
	mBase = m.M
	if v1342 != 0 {
		v1346 = v1333
		goto L413
	} else {
		goto L415
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1335 - v1331
	v1346 = int32(1)
	goto L413
L416:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1349
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1352 = int32(0)
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1351-int32(4))))
	if v1359 == v1352 {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	if v1433 < int32(4) {
		goto L411
	} else {
		goto L433
	}
L418:
	;
	v1433 = int32(0)
	goto L417
L419:
	;
	goto L420
L420:
	;
	v1364 = v1359 & int32(3)
	if base.Ui32(v1359) < base.Ui32(int32(4)) {
		goto L423
	} else {
		goto L424
	}
L421:
	;
	v1433 = v1422
	goto L417
L422:
	;
	v1406 = v1400
	v1407 = v1401
	v1411 = v1352
	goto L430
L423:
	;
	v1400 = v1351
	v1401 = int32(0)
	goto L422
L424:
	;
	goto L425
L425:
	;
	v1371 = v1351
	v1372 = int32(0)
	v1375 = v1352
	goto L426
L426:
	;
	v1377 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1371))))
	v1378 = int32(-65)
	v1381 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1371)+1)))
	v1385 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1371)+2)))
	v1389 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1371)+3)))
	v1392 = v1372 + base.B2i32(v1378 < v1377) + base.B2i32(v1378 < v1381) + base.B2i32(v1378 < v1385) + base.B2i32(v1378 < v1389)
	v1393 = int32(4)
	v1394 = v1371 + v1393
	v1396 = v1375 + v1393
	if v1396 != v1359&int32(-4) {
		v1371 = v1394
		v1372 = v1392
		v1375 = v1396
		goto L426
	} else {
		goto L428
	}
L427:
	;
	if v1364 == int32(0) {
		v1422 = v1392
		goto L421
	} else {
		goto L429
	}
L428:
	;
	goto L427
L429:
	;
	v1400 = v1394
	v1401 = v1392
	goto L422
L430:
	;
	v1412 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1406))))
	v1415 = v1407 + base.B2i32(int32(-65) < v1412)
	v1416 = int32(1)
	v1419 = v1411 + v1416
	if v1419 != v1364 {
		v1406 = v1406 + v1416
		v1407 = v1415
		v1411 = v1419
		goto L430
	} else {
		goto L432
	}
L431:
	;
	v1422 = v1415
	goto L421
L432:
	;
	goto L431
L433:
	;
	v1436 = F_slice_del(m, l0)
	mBase = m.M
	if v1436 < int32(0) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1439 = v1436
	goto L436
L435:
	;
	v1439 = v1321
	goto L436
L436:
	;
	if int32(0) <= v1436 {
		v2082 = v1439
		goto L229
	} else {
		goto L437
	}
L437:
	;
	v2069 = v1439
	v2071 = int32(base.Ui32(v1436) >> (uint(int32(31)) % 32))
	goto L410
L438:
	;
	if int32(0) <= v2057 {
		v2082 = v2049
		goto L229
	} else {
		goto L647
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2063
	v2082 = v2059
	goto L229
L440:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1857
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1857
	v1860 = int32(2)
	v1862 = int32(0)
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1857-v1865 < v1860 {
		v1875 = v1862
		goto L574
	} else {
		goto L575
	}
L441:
	;
	v1852 = v1321
	goto L440
L442:
	;
	goto L443
L443:
	;
	v1447 = int32(0)
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1448
	v1453 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_56), int32(10), v1447)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L4
	} else {
		goto L445
	}
L444:
	;
	if v1727 < int32(0) {
		goto L505
	} else {
		goto L506
	}
L445:
	;
	if v1453 == int32(0) {
		v1727 = v1447
		goto L444
	} else {
		goto L446
	}
L446:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1457
	switch v1453 - int32(1) {
	case 0:
		goto L450
	case 1:
		goto L449
	case 2:
		goto L448
	default:
		goto L447
	}
L447:
	;
	v1727 = int32(1)
	goto L444
L448:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1638 = int32(0)
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1637-int32(4))))
	if v1645 == v1638 {
		goto L488
	} else {
		goto L489
	}
L449:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1550 = int32(0)
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1549-int32(4))))
	if v1557 == v1550 {
		goto L470
	} else {
		goto L471
	}
L450:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1462 = int32(0)
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1461-int32(4))))
	if v1469 == v1462 {
		goto L452
	} else {
		goto L453
	}
L451:
	;
	if v1543 < int32(4) {
		v1727 = v1447
		goto L444
	} else {
		goto L467
	}
L452:
	;
	v1543 = int32(0)
	goto L451
L453:
	;
	goto L454
L454:
	;
	v1474 = v1469 & int32(3)
	if base.Ui32(v1469) < base.Ui32(int32(4)) {
		goto L457
	} else {
		goto L458
	}
L455:
	;
	v1543 = v1532
	goto L451
L456:
	;
	v1516 = v1510
	v1517 = v1511
	v1521 = v1462
	goto L464
L457:
	;
	v1510 = v1461
	v1511 = int32(0)
	goto L456
L458:
	;
	goto L459
L459:
	;
	v1481 = v1461
	v1482 = int32(0)
	v1485 = v1462
	goto L460
L460:
	;
	v1487 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1481))))
	v1488 = int32(-65)
	v1491 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1481)+1)))
	v1495 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1481)+2)))
	v1499 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1481)+3)))
	v1502 = v1482 + base.B2i32(v1488 < v1487) + base.B2i32(v1488 < v1491) + base.B2i32(v1488 < v1495) + base.B2i32(v1488 < v1499)
	v1503 = int32(4)
	v1504 = v1481 + v1503
	v1506 = v1485 + v1503
	if v1506 != v1469&int32(-4) {
		v1481 = v1504
		v1482 = v1502
		v1485 = v1506
		goto L460
	} else {
		goto L462
	}
L461:
	;
	if v1474 == int32(0) {
		v1532 = v1502
		goto L455
	} else {
		goto L463
	}
L462:
	;
	goto L461
L463:
	;
	v1510 = v1504
	v1511 = v1502
	goto L456
L464:
	;
	v1522 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1516))))
	v1525 = v1517 + base.B2i32(int32(-65) < v1522)
	v1526 = int32(1)
	v1529 = v1521 + v1526
	if v1529 != v1474 {
		v1516 = v1516 + v1526
		v1517 = v1525
		v1521 = v1529
		goto L464
	} else {
		goto L466
	}
L465:
	;
	v1532 = v1525
	goto L455
L466:
	;
	goto L465
L467:
	;
	v1546 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1546 {
		goto L447
	} else {
		goto L468
	}
L468:
	;
	v1727 = v1546
	goto L444
L469:
	;
	if v1631 < int32(5) {
		v1727 = v1447
		goto L444
	} else {
		goto L485
	}
L470:
	;
	v1631 = int32(0)
	goto L469
L471:
	;
	goto L472
L472:
	;
	v1562 = v1557 & int32(3)
	if base.Ui32(v1557) < base.Ui32(int32(4)) {
		goto L475
	} else {
		goto L476
	}
L473:
	;
	v1631 = v1620
	goto L469
L474:
	;
	v1604 = v1598
	v1605 = v1599
	v1609 = v1550
	goto L482
L475:
	;
	v1598 = v1549
	v1599 = int32(0)
	goto L474
L476:
	;
	goto L477
L477:
	;
	v1569 = v1549
	v1570 = int32(0)
	v1573 = v1550
	goto L478
L478:
	;
	v1575 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1569))))
	v1576 = int32(-65)
	v1579 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1569)+1)))
	v1583 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1569)+2)))
	v1587 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1569)+3)))
	v1590 = v1570 + base.B2i32(v1576 < v1575) + base.B2i32(v1576 < v1579) + base.B2i32(v1576 < v1583) + base.B2i32(v1576 < v1587)
	v1591 = int32(4)
	v1592 = v1569 + v1591
	v1594 = v1573 + v1591
	if v1594 != v1557&int32(-4) {
		v1569 = v1592
		v1570 = v1590
		v1573 = v1594
		goto L478
	} else {
		goto L480
	}
L479:
	;
	if v1562 == int32(0) {
		v1620 = v1590
		goto L473
	} else {
		goto L481
	}
L480:
	;
	goto L479
L481:
	;
	v1598 = v1592
	v1599 = v1590
	goto L474
L482:
	;
	v1610 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1604))))
	v1613 = v1605 + base.B2i32(int32(-65) < v1610)
	v1614 = int32(1)
	v1617 = v1609 + v1614
	if v1617 != v1562 {
		v1604 = v1604 + v1614
		v1605 = v1613
		v1609 = v1617
		goto L482
	} else {
		goto L484
	}
L483:
	;
	v1620 = v1613
	goto L473
L484:
	;
	goto L483
L485:
	;
	v1634 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1634 {
		goto L447
	} else {
		goto L486
	}
L486:
	;
	v1727 = v1634
	goto L444
L487:
	;
	if v1719 < int32(6) {
		v1727 = v1447
		goto L444
	} else {
		goto L503
	}
L488:
	;
	v1719 = int32(0)
	goto L487
L489:
	;
	goto L490
L490:
	;
	v1650 = v1645 & int32(3)
	if base.Ui32(v1645) < base.Ui32(int32(4)) {
		goto L493
	} else {
		goto L494
	}
L491:
	;
	v1719 = v1708
	goto L487
L492:
	;
	v1692 = v1686
	v1693 = v1687
	v1697 = v1638
	goto L500
L493:
	;
	v1686 = v1637
	v1687 = int32(0)
	goto L492
L494:
	;
	goto L495
L495:
	;
	v1657 = v1637
	v1658 = int32(0)
	v1661 = v1638
	goto L496
L496:
	;
	v1663 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1657))))
	v1664 = int32(-65)
	v1667 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1657)+1)))
	v1671 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1657)+2)))
	v1675 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1657)+3)))
	v1678 = v1658 + base.B2i32(v1664 < v1663) + base.B2i32(v1664 < v1667) + base.B2i32(v1664 < v1671) + base.B2i32(v1664 < v1675)
	v1679 = int32(4)
	v1680 = v1657 + v1679
	v1682 = v1661 + v1679
	if v1682 != v1645&int32(-4) {
		v1657 = v1680
		v1658 = v1678
		v1661 = v1682
		goto L496
	} else {
		goto L498
	}
L497:
	;
	if v1650 == int32(0) {
		v1708 = v1678
		goto L491
	} else {
		goto L499
	}
L498:
	;
	goto L497
L499:
	;
	v1686 = v1680
	v1687 = v1678
	goto L492
L500:
	;
	v1698 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1692))))
	v1701 = v1693 + base.B2i32(int32(-65) < v1698)
	v1702 = int32(1)
	v1705 = v1697 + v1702
	if v1705 != v1650 {
		v1692 = v1692 + v1702
		v1693 = v1701
		v1697 = v1705
		goto L500
	} else {
		goto L502
	}
L501:
	;
	v1708 = v1701
	goto L491
L502:
	;
	goto L501
L503:
	;
	v1722 = F_slice_del(m, l0)
	mBase = m.M
	if v1722 < int32(0) {
		v1727 = v1722
		goto L444
	} else {
		goto L504
	}
L504:
	;
	goto L447
L505:
	;
	v1730 = v1727
	goto L507
L506:
	;
	v1730 = v1321
	goto L507
L507:
	;
	if v1727 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v1731 = v1730
	goto L510
L509:
	;
	v1731 = v1321
	goto L510
L510:
	;
	v1733 = int32(base.Ui32(v1727) >> (uint(int32(31)) % 32))
	if v1727 != 0 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v1735 = v1733
	goto L513
L512:
	;
	v1735 = int32(20)
	goto L513
L513:
	;
	if v1735 == int32(20) {
		v1852 = v1731
		goto L440
	} else {
		goto L514
	}
L514:
	;
	if v1735 != 0 {
		v2069 = v1731
		v2071 = v1733
		goto L410
	} else {
		goto L515
	}
L515:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1740 = F_r_Suffix_Noun_Step2a(m, l0)
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L4
	} else {
		goto L516
	}
L516:
	;
	if v1740 < int32(0) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v1744 = v1740
	goto L519
L518:
	;
	v1744 = v1731
	goto L519
L519:
	;
	if v1740 != 0 {
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v1745 = v1744
	goto L522
L521:
	;
	v1745 = v1731
	goto L522
L522:
	;
	v1747 = int32(base.Ui32(v1740) >> (uint(int32(31)) % 32))
	if v1740 != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v1749 = v1747
	goto L525
L524:
	;
	v1749 = int32(23)
	goto L525
L525:
	;
	if v1749 == int32(0) {
		v2082 = v1745
		goto L229
	} else {
		goto L526
	}
L526:
	;
	if v1749 != int32(23) {
		v2078 = v1745
		v2080 = v1747
		goto L230
	} else {
		goto L527
	}
L527:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1755 = v1739 - v1738
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1754 - v1755
	v1761 = Fn14362(m, l0, int32(5), int32(_a_F_arabic_UTF_8_stem_57), int32(4))
	mBase = m.M
	goto L528
L528:
	;
	if v1761 < int32(0) {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v1764 = v1761
	goto L531
L530:
	;
	v1764 = v1745
	goto L531
L531:
	;
	if v1761 != 0 {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v1765 = v1764
	goto L534
L533:
	;
	v1765 = v1745
	goto L534
L534:
	;
	v1767 = int32(base.Ui32(v1761) >> (uint(int32(31)) % 32))
	if v1761 != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v1769 = v1767
	goto L537
L536:
	;
	v1769 = int32(24)
	goto L537
L537:
	;
	if v1769 == int32(0) {
		v2082 = v1765
		goto L229
	} else {
		goto L538
	}
L538:
	;
	if v1769 != int32(24) {
		v2078 = v1765
		v2080 = v1767
		goto L230
	} else {
		goto L539
	}
L539:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1774 - v1755
	v1780 = Fn14362(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_58), int32(2))
	mBase = m.M
	goto L540
L540:
	;
	if v1780 < int32(0) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v1783 = v1780
	goto L543
L542:
	;
	v1783 = v1765
	goto L543
L543:
	;
	if v1780 != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v1784 = v1783
	goto L546
L545:
	;
	v1784 = v1765
	goto L546
L546:
	;
	v1786 = int32(base.Ui32(v1780) >> (uint(int32(31)) % 32))
	if v1780 != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v1788 = v1786
	goto L549
L548:
	;
	v1788 = int32(25)
	goto L549
L549:
	;
	if v1788 == int32(0) {
		v2082 = v1784
		goto L229
	} else {
		goto L550
	}
L550:
	;
	if v1788 != int32(25) {
		v2078 = v1784
		v2080 = v1786
		goto L230
	} else {
		goto L551
	}
L551:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1794 = v1793 - v1755
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1794
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L554
L552:
	;
	if int32(0) <= v1849 {
		v2059 = v1784
		v2063 = v1849
		goto L439
	} else {
		goto L571
	}
L554:
	;
	goto L555
L555:
	;
	goto L556
L556:
	;
	v1804 = v1794
	v1806 = int32(1)
	goto L559
L558:
	;
	v1849 = v1831
	goto L552
L559:
	;
	if v1804 <= v1797 {
		goto L561
	} else {
		goto L562
	}
L560:
	;
	goto L558
L561:
	;
	v1849 = int32(-1)
	goto L552
L562:
	;
	goto L563
L563:
	;
	v1811 = v1804 - int32(1)
	v1813 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1796+v1811))))
	if base.B2i32(int32(0) <= v1813)|base.B2i32(v1811 <= v1797) != 0 {
		v1831 = v1811
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v1835 = int32(1)
	if v1835 < v1806 {
		v1804 = v1831
		v1806 = v1806 - v1835
		goto L559
	} else {
		goto L570
	}
L565:
	;
	v1819 = v1811
	goto L566
L566:
	;
	v1824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1796+v1819))))
	if base.Ui32(int32(191)) < base.Ui32(v1824) {
		v1831 = v1819
		goto L564
	} else {
		goto L568
	}
L567:
	;
	v1831 = v1797
	goto L564
L568:
	;
	v1828 = v1819 - int32(1)
	if v1797 < v1828 {
		v1819 = v1828
		goto L566
	} else {
		goto L569
	}
L569:
	;
	goto L567
L570:
	;
	goto L560
L571:
	;
	v1852 = v1784
	goto L440
L572:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2031
	v2033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2033 != 0 {
		goto L630
	} else {
		goto L631
	}
L573:
	;
	if v1875 == int32(0) {
		v2026 = v1852
		goto L572
	} else {
		goto L577
	}
L574:
	;
	goto L573
L575:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1871 = F_memcmp(m, v1868+v1857-v1860, int32(_a_F_arabic_UTF_8_stem_59), v1860)
	mBase = m.M
	if v1871 != 0 {
		v1875 = v1862
		goto L574
	} else {
		goto L576
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1857 - v1860
	v1875 = int32(1)
	goto L574
L577:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1878
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1881 = int32(0)
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1880-int32(4))))
	if v1888 == v1881 {
		goto L579
	} else {
		goto L580
	}
L578:
	;
	if v1962 < int32(6) {
		v2026 = v1852
		goto L572
	} else {
		goto L594
	}
L579:
	;
	v1962 = int32(0)
	goto L578
L580:
	;
	goto L581
L581:
	;
	v1893 = v1888 & int32(3)
	if base.Ui32(v1888) < base.Ui32(int32(4)) {
		goto L584
	} else {
		goto L585
	}
L582:
	;
	v1962 = v1951
	goto L578
L583:
	;
	v1935 = v1929
	v1936 = v1930
	v1940 = v1881
	goto L591
L584:
	;
	v1929 = v1880
	v1930 = int32(0)
	goto L583
L585:
	;
	goto L586
L586:
	;
	v1900 = v1880
	v1901 = int32(0)
	v1904 = v1881
	goto L587
L587:
	;
	v1906 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1900))))
	v1907 = int32(-65)
	v1910 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1900)+1)))
	v1914 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1900)+2)))
	v1918 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1900)+3)))
	v1921 = v1901 + base.B2i32(v1907 < v1906) + base.B2i32(v1907 < v1910) + base.B2i32(v1907 < v1914) + base.B2i32(v1907 < v1918)
	v1922 = int32(4)
	v1923 = v1900 + v1922
	v1925 = v1904 + v1922
	if v1925 != v1888&int32(-4) {
		v1900 = v1923
		v1901 = v1921
		v1904 = v1925
		goto L587
	} else {
		goto L589
	}
L588:
	;
	if v1893 == int32(0) {
		v1951 = v1921
		goto L582
	} else {
		goto L590
	}
L589:
	;
	goto L588
L590:
	;
	v1929 = v1923
	v1930 = v1921
	goto L583
L591:
	;
	v1941 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1935))))
	v1944 = v1936 + base.B2i32(int32(-65) < v1941)
	v1945 = int32(1)
	v1948 = v1940 + v1945
	if v1948 != v1893 {
		v1935 = v1935 + v1945
		v1936 = v1944
		v1940 = v1948
		goto L591
	} else {
		goto L593
	}
L592:
	;
	v1951 = v1944
	goto L582
L593:
	;
	goto L592
L594:
	;
	v1965 = F_slice_del(m, l0)
	mBase = m.M
	v1967 = base.B2i32(v1965 < int32(0))
	if v1965 < int32(0) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	if v1965 < int32(0) {
		goto L598
	} else {
		goto L599
	}
L596:
	;
	goto L597
L597:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1973 = F_r_Suffix_Noun_Step2a(m, l0)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L4
	} else {
		goto L601
	}
L598:
	;
	v1968 = v1965
	goto L600
L599:
	;
	v1968 = v1852
	goto L600
L600:
	;
	v2069 = v1968
	v2071 = int32(base.Ui32(v1965) >> (uint(int32(31)) % 32))
	goto L410
L601:
	;
	if v1973 < int32(0) {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v1977 = v1973
	goto L604
L603:
	;
	v1977 = v1852
	goto L604
L604:
	;
	if v1973 != 0 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	v1978 = v1977
	goto L607
L606:
	;
	v1978 = v1852
	goto L607
L607:
	;
	v1980 = int32(base.Ui32(v1973) >> (uint(int32(31)) % 32))
	if v1973 != 0 {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v1982 = v1980
	goto L610
L609:
	;
	v1982 = int32(29)
	goto L610
L610:
	;
	if v1982 == int32(0) {
		v2082 = v1978
		goto L229
	} else {
		goto L611
	}
L611:
	;
	if v1982 != int32(29) {
		v2020 = v1978
		v2022 = v1980
		goto L612
	} else {
		goto L613
	}
L612:
	;
	if v2022 == int32(0) {
		v2082 = v2020
		goto L229
	} else {
		goto L629
	}
L613:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1988 = v1972 - v1971
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1987 - v1988
	v1994 = Fn14362(m, l0, int32(5), int32(_a_F_arabic_UTF_8_stem_57), int32(4))
	mBase = m.M
	goto L614
L614:
	;
	if v1994 < int32(0) {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v1997 = v1994
	goto L617
L616:
	;
	v1997 = v1978
	goto L617
L617:
	;
	if v1994 != 0 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v1998 = v1997
	goto L620
L619:
	;
	v1998 = v1978
	goto L620
L620:
	;
	v2000 = int32(base.Ui32(v1994) >> (uint(int32(31)) % 32))
	if v1994 != 0 {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v2002 = v2000
	goto L623
L622:
	;
	v2002 = int32(30)
	goto L623
L623:
	;
	if v2002 == int32(0) {
		v2082 = v1998
		goto L229
	} else {
		goto L624
	}
L624:
	;
	if v2002 != int32(30) {
		v2020 = v1998
		v2022 = v2000
		goto L612
	} else {
		goto L625
	}
L625:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2007 - v1988
	v2013 = Fn14362(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_58), int32(2))
	mBase = m.M
	goto L626
L626:
	;
	if v2013 == int32(0) {
		v2026 = v1998
		goto L572
	} else {
		goto L627
	}
L627:
	;
	if int32(0) <= v2013 {
		v2082 = v1998
		goto L229
	} else {
		goto L628
	}
L628:
	;
	v2020 = v2013
	v2022 = int32(base.Ui32(v2013) >> (uint(int32(31)) % 32))
	goto L612
L629:
	;
	v3981 = v2020
	goto L44
L630:
	;
	v2049 = v2026
	v2052 = v2031
	goto L632
L631:
	;
	v2034 = F_r_Suffix_Noun_Step2a(m, l0)
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L4
	} else {
		goto L633
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2052
	v2057 = Fn14362(m, l0, int32(5), int32(_a_F_arabic_UTF_8_stem_57), int32(4))
	mBase = m.M
	goto L645
L633:
	;
	if v2034 < int32(0) {
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v2038 = v2034
	goto L636
L635:
	;
	v2038 = v2026
	goto L636
L636:
	;
	if v2034 != 0 {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v2039 = v2038
	goto L639
L638:
	;
	v2039 = v2026
	goto L639
L639:
	;
	v2040 = int32(31)
	v2041 = int32(base.Ui32(v2034) >> (uint(v2040) % 32))
	if v2034 != 0 {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v2043 = v2041
	goto L642
L641:
	;
	v2043 = v2040
	goto L642
L642:
	;
	if v2043 == int32(0) {
		v2082 = v2039
		goto L229
	} else {
		goto L643
	}
L643:
	;
	if v2043 != int32(31) {
		v2069 = v2039
		v2071 = v2041
		goto L410
	} else {
		goto L644
	}
L644:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2049 = v2039
	v2052 = v2048
	goto L632
L645:
	;
	if v2057 != 0 {
		goto L438
	} else {
		goto L646
	}
L646:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2059 = v2049
	v2063 = v2058
	goto L439
L647:
	;
	v2069 = v2057
	v2071 = int32(base.Ui32(v2057) >> (uint(int32(31)) % 32))
	goto L410
L648:
	;
	v2082 = v2069
	goto L229
L649:
	;
	v3981 = v924
	goto L44
L650:
	;
	v3981 = v1316
	goto L44
L651:
	;
	v2082 = v2078
	goto L229
L652:
	;
	if v2104 != 0 {
		goto L656
	} else {
		goto L657
	}
L653:
	;
	goto L652
L654:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2100 = F_memcmp(m, v2097+v2087-v2089, int32(_a_F_arabic_UTF_8_stem_60), v2089)
	mBase = m.M
	if v2100 != 0 {
		v2104 = v2091
		goto L653
	} else {
		goto L655
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2087 - v2089
	v2104 = int32(1)
	goto L653
L656:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2105
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2108 = int32(0)
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2107-int32(4))))
	if v2115 == v2108 {
		goto L660
	} else {
		goto L661
	}
L657:
	;
	goto L658
L658:
	;
	v2193 = v2082
	goto L228
L659:
	;
	if int32(3) <= v2189 {
		goto L227
	} else {
		goto L675
	}
L660:
	;
	v2189 = int32(0)
	goto L659
L661:
	;
	goto L662
L662:
	;
	v2120 = v2115 & int32(3)
	if base.Ui32(v2115) < base.Ui32(int32(4)) {
		goto L665
	} else {
		goto L666
	}
L663:
	;
	v2189 = v2178
	goto L659
L664:
	;
	v2162 = v2156
	v2163 = v2157
	v2167 = v2108
	goto L672
L665:
	;
	v2156 = v2107
	v2157 = int32(0)
	goto L664
L666:
	;
	goto L667
L667:
	;
	v2127 = v2107
	v2128 = int32(0)
	v2131 = v2108
	goto L668
L668:
	;
	v2133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2127))))
	v2134 = int32(-65)
	v2137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2127)+1)))
	v2141 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2127)+2)))
	v2145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2127)+3)))
	v2148 = v2128 + base.B2i32(v2134 < v2133) + base.B2i32(v2134 < v2137) + base.B2i32(v2134 < v2141) + base.B2i32(v2134 < v2145)
	v2149 = int32(4)
	v2150 = v2127 + v2149
	v2152 = v2131 + v2149
	if v2152 != v2115&int32(-4) {
		v2127 = v2150
		v2128 = v2148
		v2131 = v2152
		goto L668
	} else {
		goto L670
	}
L669:
	;
	if v2120 == int32(0) {
		v2178 = v2148
		goto L663
	} else {
		goto L671
	}
L670:
	;
	goto L669
L671:
	;
	v2156 = v2150
	v2157 = v2148
	goto L664
L672:
	;
	v2168 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2162))))
	v2171 = v2163 + base.B2i32(int32(-65) < v2168)
	v2172 = int32(1)
	v2175 = v2167 + v2172
	if v2175 != v2120 {
		v2162 = v2162 + v2172
		v2163 = v2171
		v2167 = v2175
		goto L672
	} else {
		goto L674
	}
L673:
	;
	v2178 = v2171
	goto L663
L674:
	;
	goto L673
L675:
	;
	goto L658
L676:
	;
	if v2215 == int32(0) {
		v2232 = v2193
		goto L226
	} else {
		goto L680
	}
L677:
	;
	goto L676
L678:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2211 = F_memcmp(m, v2208+v2197-v2200, int32(_a_F_arabic_UTF_8_stem_61), v2200)
	mBase = m.M
	if v2211 != 0 {
		v2215 = v2202
		goto L677
	} else {
		goto L679
	}
L679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2197 - v2200
	v2215 = int32(1)
	goto L677
L680:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2218
	v2222 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_62))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L4
	} else {
		goto L681
	}
L681:
	;
	if v2222 < int32(0) {
		v3981 = v2222
		goto L44
	} else {
		goto L682
	}
L682:
	;
	v2232 = v2193
	goto L226
L683:
	;
	v2229 = v2226
	goto L685
L684:
	;
	v2229 = v2082
	goto L685
L685:
	;
	if v2226 < int32(0) {
		v3981 = v2229
		goto L44
	} else {
		goto L686
	}
L686:
	;
	v2232 = v2229
	goto L226
L687:
	;
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2636
	v2639 = v2636 + int32(1)
	v2640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2640 <= v2639 {
		goto L774
	} else {
		goto L775
	}
L688:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2237
	goto L687
L689:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2244+v2241))))
	if base.B2i32(v2246&int32(224) != int32(160))|base.B2i32(int32(1)<<(uint(v2246)%32)&int32(188) == int32(0)) != 0 {
		goto L688
	} else {
		goto L690
	}
L690:
	;
	v2261 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_63), int32(5), int32(0))
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L4
	} else {
		goto L691
	}
L691:
	;
	if v2261 == int32(0) {
		goto L688
	} else {
		goto L692
	}
L692:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2265
	switch v2261 - int32(1) {
	case 0:
		goto L696
	case 1:
		goto L695
	case 2:
		goto L694
	case 3:
		goto L693
	default:
		goto L687
	}
L693:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2543 = int32(0)
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2542-int32(4))))
	if v2550 == v2543 {
		goto L755
	} else {
		goto L756
	}
L694:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2452 = int32(0)
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2451-int32(4))))
	if v2459 == v2452 {
		goto L736
	} else {
		goto L737
	}
L695:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2361 = int32(0)
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v2360-int32(4))))
	if v2368 == v2361 {
		goto L717
	} else {
		goto L718
	}
L696:
	;
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2270 = int32(0)
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v2269-int32(4))))
	if v2277 == v2270 {
		goto L698
	} else {
		goto L699
	}
L697:
	;
	if v2351 < int32(4) {
		goto L688
	} else {
		goto L713
	}
L698:
	;
	v2351 = int32(0)
	goto L697
L699:
	;
	goto L700
L700:
	;
	v2282 = v2277 & int32(3)
	if base.Ui32(v2277) < base.Ui32(int32(4)) {
		goto L703
	} else {
		goto L704
	}
L701:
	;
	v2351 = v2340
	goto L697
L702:
	;
	v2324 = v2318
	v2325 = v2319
	v2329 = v2270
	goto L710
L703:
	;
	v2318 = v2269
	v2319 = int32(0)
	goto L702
L704:
	;
	goto L705
L705:
	;
	v2289 = v2269
	v2290 = int32(0)
	v2293 = v2270
	goto L706
L706:
	;
	v2295 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2289))))
	v2296 = int32(-65)
	v2299 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2289)+1)))
	v2303 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2289)+2)))
	v2307 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2289)+3)))
	v2310 = v2290 + base.B2i32(v2296 < v2295) + base.B2i32(v2296 < v2299) + base.B2i32(v2296 < v2303) + base.B2i32(v2296 < v2307)
	v2311 = int32(4)
	v2312 = v2289 + v2311
	v2314 = v2293 + v2311
	if v2314 != v2277&int32(-4) {
		v2289 = v2312
		v2290 = v2310
		v2293 = v2314
		goto L706
	} else {
		goto L708
	}
L707:
	;
	if v2282 == int32(0) {
		v2340 = v2310
		goto L701
	} else {
		goto L709
	}
L708:
	;
	goto L707
L709:
	;
	v2318 = v2312
	v2319 = v2310
	goto L702
L710:
	;
	v2330 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2324))))
	v2333 = v2325 + base.B2i32(int32(-65) < v2330)
	v2334 = int32(1)
	v2337 = v2329 + v2334
	if v2337 != v2282 {
		v2324 = v2324 + v2334
		v2325 = v2333
		v2329 = v2337
		goto L710
	} else {
		goto L712
	}
L711:
	;
	v2340 = v2333
	goto L701
L712:
	;
	goto L711
L713:
	;
	v2356 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_64))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L4
	} else {
		goto L714
	}
L714:
	;
	if int32(0) <= v2356 {
		goto L687
	} else {
		goto L715
	}
L715:
	;
	v3981 = v2356
	goto L44
L716:
	;
	if v2442 < int32(4) {
		goto L688
	} else {
		goto L732
	}
L717:
	;
	v2442 = int32(0)
	goto L716
L718:
	;
	goto L719
L719:
	;
	v2373 = v2368 & int32(3)
	if base.Ui32(v2368) < base.Ui32(int32(4)) {
		goto L722
	} else {
		goto L723
	}
L720:
	;
	v2442 = v2431
	goto L716
L721:
	;
	v2415 = v2409
	v2416 = v2410
	v2420 = v2361
	goto L729
L722:
	;
	v2409 = v2360
	v2410 = int32(0)
	goto L721
L723:
	;
	goto L724
L724:
	;
	v2380 = v2360
	v2381 = int32(0)
	v2384 = v2361
	goto L725
L725:
	;
	v2386 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2380))))
	v2387 = int32(-65)
	v2390 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2380)+1)))
	v2394 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2380)+2)))
	v2398 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2380)+3)))
	v2401 = v2381 + base.B2i32(v2387 < v2386) + base.B2i32(v2387 < v2390) + base.B2i32(v2387 < v2394) + base.B2i32(v2387 < v2398)
	v2402 = int32(4)
	v2403 = v2380 + v2402
	v2405 = v2384 + v2402
	if v2405 != v2368&int32(-4) {
		v2380 = v2403
		v2381 = v2401
		v2384 = v2405
		goto L725
	} else {
		goto L727
	}
L726:
	;
	if v2373 == int32(0) {
		v2431 = v2401
		goto L720
	} else {
		goto L728
	}
L727:
	;
	goto L726
L728:
	;
	v2409 = v2403
	v2410 = v2401
	goto L721
L729:
	;
	v2421 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2415))))
	v2424 = v2416 + base.B2i32(int32(-65) < v2421)
	v2425 = int32(1)
	v2428 = v2420 + v2425
	if v2428 != v2373 {
		v2415 = v2415 + v2425
		v2416 = v2424
		v2420 = v2428
		goto L729
	} else {
		goto L731
	}
L730:
	;
	v2431 = v2424
	goto L720
L731:
	;
	goto L730
L732:
	;
	v2447 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_65))
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L4
	} else {
		goto L733
	}
L733:
	;
	if int32(0) <= v2447 {
		goto L687
	} else {
		goto L734
	}
L734:
	;
	v3981 = v2447
	goto L44
L735:
	;
	if v2533 < int32(4) {
		goto L688
	} else {
		goto L751
	}
L736:
	;
	v2533 = int32(0)
	goto L735
L737:
	;
	goto L738
L738:
	;
	v2464 = v2459 & int32(3)
	if base.Ui32(v2459) < base.Ui32(int32(4)) {
		goto L741
	} else {
		goto L742
	}
L739:
	;
	v2533 = v2522
	goto L735
L740:
	;
	v2506 = v2500
	v2507 = v2501
	v2511 = v2452
	goto L748
L741:
	;
	v2500 = v2451
	v2501 = int32(0)
	goto L740
L742:
	;
	goto L743
L743:
	;
	v2471 = v2451
	v2472 = int32(0)
	v2475 = v2452
	goto L744
L744:
	;
	v2477 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2471))))
	v2478 = int32(-65)
	v2481 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2471)+1)))
	v2485 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2471)+2)))
	v2489 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2471)+3)))
	v2492 = v2472 + base.B2i32(v2478 < v2477) + base.B2i32(v2478 < v2481) + base.B2i32(v2478 < v2485) + base.B2i32(v2478 < v2489)
	v2493 = int32(4)
	v2494 = v2471 + v2493
	v2496 = v2475 + v2493
	if v2496 != v2459&int32(-4) {
		v2471 = v2494
		v2472 = v2492
		v2475 = v2496
		goto L744
	} else {
		goto L746
	}
L745:
	;
	if v2464 == int32(0) {
		v2522 = v2492
		goto L739
	} else {
		goto L747
	}
L746:
	;
	goto L745
L747:
	;
	v2500 = v2494
	v2501 = v2492
	goto L740
L748:
	;
	v2512 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2506))))
	v2515 = v2507 + base.B2i32(int32(-65) < v2512)
	v2516 = int32(1)
	v2519 = v2511 + v2516
	if v2519 != v2464 {
		v2506 = v2506 + v2516
		v2507 = v2515
		v2511 = v2519
		goto L748
	} else {
		goto L750
	}
L749:
	;
	v2522 = v2515
	goto L739
L750:
	;
	goto L749
L751:
	;
	v2538 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_66))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L4
	} else {
		goto L752
	}
L752:
	;
	if int32(0) <= v2538 {
		goto L687
	} else {
		goto L753
	}
L753:
	;
	v3981 = v2538
	goto L44
L754:
	;
	if v2624 < int32(4) {
		goto L688
	} else {
		goto L770
	}
L755:
	;
	v2624 = int32(0)
	goto L754
L756:
	;
	goto L757
L757:
	;
	v2555 = v2550 & int32(3)
	if base.Ui32(v2550) < base.Ui32(int32(4)) {
		goto L760
	} else {
		goto L761
	}
L758:
	;
	v2624 = v2613
	goto L754
L759:
	;
	v2597 = v2591
	v2598 = v2592
	v2602 = v2543
	goto L767
L760:
	;
	v2591 = v2542
	v2592 = int32(0)
	goto L759
L761:
	;
	goto L762
L762:
	;
	v2562 = v2542
	v2563 = int32(0)
	v2566 = v2543
	goto L763
L763:
	;
	v2568 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2562))))
	v2569 = int32(-65)
	v2572 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2562)+1)))
	v2576 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2562)+2)))
	v2580 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2562)+3)))
	v2583 = v2563 + base.B2i32(v2569 < v2568) + base.B2i32(v2569 < v2572) + base.B2i32(v2569 < v2576) + base.B2i32(v2569 < v2580)
	v2584 = int32(4)
	v2585 = v2562 + v2584
	v2587 = v2566 + v2584
	if v2587 != v2550&int32(-4) {
		v2562 = v2585
		v2563 = v2583
		v2566 = v2587
		goto L763
	} else {
		goto L765
	}
L764:
	;
	if v2555 == int32(0) {
		v2613 = v2583
		goto L758
	} else {
		goto L766
	}
L765:
	;
	goto L764
L766:
	;
	v2591 = v2585
	v2592 = v2583
	goto L759
L767:
	;
	v2603 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2597))))
	v2606 = v2598 + base.B2i32(int32(-65) < v2603)
	v2607 = int32(1)
	v2610 = v2602 + v2607
	if v2610 != v2555 {
		v2597 = v2597 + v2607
		v2598 = v2606
		v2602 = v2610
		goto L767
	} else {
		goto L769
	}
L768:
	;
	v2613 = v2606
	goto L758
L769:
	;
	goto L768
L770:
	;
	v2629 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_67))
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L4
	} else {
		goto L771
	}
L771:
	;
	if int32(0) <= v2629 {
		goto L687
	} else {
		goto L772
	}
L772:
	;
	v3981 = v2629
	goto L44
L773:
	;
	v2766 = int32(0)
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2767
	v2770 = v2767 + int32(3)
	v2771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2771 <= v2770 {
		v2971 = v2766
		goto L802
	} else {
		goto L803
	}
L774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2636
	v2765 = v2636
	goto L773
L775:
	;
	v2642 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2642+v2639))))
	switch v2644 - int32(129) {
	case 0, 7:
		goto L776
	default:
		goto L774
	}
L776:
	;
	v2650 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_68), int32(2), int32(0))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L4
	} else {
		goto L777
	}
L777:
	;
	if v2650 == int32(0) {
		goto L774
	} else {
		goto L778
	}
L778:
	;
	v2654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2654
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2657 = int32(0)
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v2656-int32(4))))
	if v2664 == v2657 {
		goto L780
	} else {
		goto L781
	}
L779:
	;
	if v2738 < int32(4) {
		goto L774
	} else {
		goto L795
	}
L780:
	;
	v2738 = int32(0)
	goto L779
L781:
	;
	goto L782
L782:
	;
	v2669 = v2664 & int32(3)
	if base.Ui32(v2664) < base.Ui32(int32(4)) {
		goto L785
	} else {
		goto L786
	}
L783:
	;
	v2738 = v2727
	goto L779
L784:
	;
	v2711 = v2705
	v2712 = v2706
	v2716 = v2657
	goto L792
L785:
	;
	v2705 = v2656
	v2706 = int32(0)
	goto L784
L786:
	;
	goto L787
L787:
	;
	v2676 = v2656
	v2677 = int32(0)
	v2680 = v2657
	goto L788
L788:
	;
	v2682 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2676))))
	v2683 = int32(-65)
	v2686 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2676)+1)))
	v2690 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2676)+2)))
	v2694 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2676)+3)))
	v2697 = v2677 + base.B2i32(v2683 < v2682) + base.B2i32(v2683 < v2686) + base.B2i32(v2683 < v2690) + base.B2i32(v2683 < v2694)
	v2698 = int32(4)
	v2699 = v2676 + v2698
	v2701 = v2680 + v2698
	if v2701 != v2664&int32(-4) {
		v2676 = v2699
		v2677 = v2697
		v2680 = v2701
		goto L788
	} else {
		goto L790
	}
L789:
	;
	if v2669 == int32(0) {
		v2727 = v2697
		goto L783
	} else {
		goto L791
	}
L790:
	;
	goto L789
L791:
	;
	v2705 = v2699
	v2706 = v2697
	goto L784
L792:
	;
	v2717 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2711))))
	v2720 = v2712 + base.B2i32(int32(-65) < v2717)
	v2721 = int32(1)
	v2724 = v2716 + v2721
	if v2724 != v2669 {
		v2711 = v2711 + v2721
		v2712 = v2720
		v2716 = v2724
		goto L792
	} else {
		goto L794
	}
L793:
	;
	v2727 = v2720
	goto L783
L794:
	;
	goto L793
L795:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2742 = int32(2)
	v2744 = int32(0)
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2746-v2741 < v2742 {
		v2756 = v2744
		goto L797
	} else {
		goto L798
	}
L796:
	;
	if v2756 != 0 {
		goto L774
	} else {
		goto L800
	}
L797:
	;
	goto L796
L798:
	;
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2752 = F_memcmp(m, v2750+v2741, int32(_a_F_arabic_UTF_8_stem_69), v2742)
	mBase = m.M
	if v2752 != 0 {
		v2756 = v2744
		goto L797
	} else {
		goto L799
	}
L799:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2742 + v2741
	v2756 = int32(1)
	goto L797
L800:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2741
	v2758 = F_slice_del(m, l0)
	mBase = m.M
	if v2758 < int32(0) {
		v3981 = v2758
		goto L44
	} else {
		goto L801
	}
L801:
	;
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2765 = v2761
	goto L773
L802:
	;
	v2973 = int32(base.Ui32(v2971) >> (uint(int32(31)) % 32))
	if v2971 != 0 {
		goto L847
	} else {
		goto L848
	}
L803:
	;
	v2773 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2773+v2770))))
	if base.B2i32(v2775 != int32(167))&base.B2i32(v2775 != int32(132)) != 0 {
		v2971 = v2766
		goto L802
	} else {
		goto L804
	}
L804:
	;
	v2784 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_70), int32(4), int32(0))
	mBase = m.M
	v2785 = m.ExcPending
	if v2785 != 0 {
		goto L4
	} else {
		goto L805
	}
L805:
	;
	if v2784 == int32(0) {
		v2971 = v2766
		goto L802
	} else {
		goto L806
	}
L806:
	;
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2788
	switch v2784 - int32(1) {
	case 0:
		goto L809
	case 1:
		goto L808
	default:
		goto L807
	}
L807:
	;
	v2971 = int32(1)
	goto L802
L808:
	;
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2881 = int32(0)
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v2880-int32(4))))
	if v2888 == v2881 {
		goto L829
	} else {
		goto L830
	}
L809:
	;
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2793 = int32(0)
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v2792-int32(4))))
	if v2800 == v2793 {
		goto L811
	} else {
		goto L812
	}
L810:
	;
	if v2874 < int32(6) {
		v2971 = v2766
		goto L802
	} else {
		goto L826
	}
L811:
	;
	v2874 = int32(0)
	goto L810
L812:
	;
	goto L813
L813:
	;
	v2805 = v2800 & int32(3)
	if base.Ui32(v2800) < base.Ui32(int32(4)) {
		goto L816
	} else {
		goto L817
	}
L814:
	;
	v2874 = v2863
	goto L810
L815:
	;
	v2847 = v2841
	v2848 = v2842
	v2852 = v2793
	goto L823
L816:
	;
	v2841 = v2792
	v2842 = int32(0)
	goto L815
L817:
	;
	goto L818
L818:
	;
	v2812 = v2792
	v2813 = int32(0)
	v2816 = v2793
	goto L819
L819:
	;
	v2818 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2812))))
	v2819 = int32(-65)
	v2822 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2812)+1)))
	v2826 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2812)+2)))
	v2830 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2812)+3)))
	v2833 = v2813 + base.B2i32(v2819 < v2818) + base.B2i32(v2819 < v2822) + base.B2i32(v2819 < v2826) + base.B2i32(v2819 < v2830)
	v2834 = int32(4)
	v2835 = v2812 + v2834
	v2837 = v2816 + v2834
	if v2837 != v2800&int32(-4) {
		v2812 = v2835
		v2813 = v2833
		v2816 = v2837
		goto L819
	} else {
		goto L821
	}
L820:
	;
	if v2805 == int32(0) {
		v2863 = v2833
		goto L814
	} else {
		goto L822
	}
L821:
	;
	goto L820
L822:
	;
	v2841 = v2835
	v2842 = v2833
	goto L815
L823:
	;
	v2853 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2847))))
	v2856 = v2848 + base.B2i32(int32(-65) < v2853)
	v2857 = int32(1)
	v2860 = v2852 + v2857
	if v2860 != v2805 {
		v2847 = v2847 + v2857
		v2848 = v2856
		v2852 = v2860
		goto L823
	} else {
		goto L825
	}
L824:
	;
	v2863 = v2856
	goto L814
L825:
	;
	goto L824
L826:
	;
	v2877 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2877 {
		goto L807
	} else {
		goto L827
	}
L827:
	;
	v2971 = v2877
	goto L802
L828:
	;
	if v2962 < int32(5) {
		v2971 = v2766
		goto L802
	} else {
		goto L844
	}
L829:
	;
	v2962 = int32(0)
	goto L828
L830:
	;
	goto L831
L831:
	;
	v2893 = v2888 & int32(3)
	if base.Ui32(v2888) < base.Ui32(int32(4)) {
		goto L834
	} else {
		goto L835
	}
L832:
	;
	v2962 = v2951
	goto L828
L833:
	;
	v2935 = v2929
	v2936 = v2930
	v2940 = v2881
	goto L841
L834:
	;
	v2929 = v2880
	v2930 = int32(0)
	goto L833
L835:
	;
	goto L836
L836:
	;
	v2900 = v2880
	v2901 = int32(0)
	v2904 = v2881
	goto L837
L837:
	;
	v2906 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2900))))
	v2907 = int32(-65)
	v2910 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2900)+1)))
	v2914 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2900)+2)))
	v2918 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2900)+3)))
	v2921 = v2901 + base.B2i32(v2907 < v2906) + base.B2i32(v2907 < v2910) + base.B2i32(v2907 < v2914) + base.B2i32(v2907 < v2918)
	v2922 = int32(4)
	v2923 = v2900 + v2922
	v2925 = v2904 + v2922
	if v2925 != v2888&int32(-4) {
		v2900 = v2923
		v2901 = v2921
		v2904 = v2925
		goto L837
	} else {
		goto L839
	}
L838:
	;
	if v2893 == int32(0) {
		v2951 = v2921
		goto L832
	} else {
		goto L840
	}
L839:
	;
	goto L838
L840:
	;
	v2929 = v2923
	v2930 = v2921
	goto L833
L841:
	;
	v2941 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2935))))
	v2944 = v2936 + base.B2i32(int32(-65) < v2941)
	v2945 = int32(1)
	v2948 = v2940 + v2945
	if v2948 != v2893 {
		v2935 = v2935 + v2945
		v2936 = v2944
		v2940 = v2948
		goto L841
	} else {
		goto L843
	}
L842:
	;
	v2951 = v2944
	goto L832
L843:
	;
	goto L842
L844:
	;
	v2965 = F_slice_del(m, l0)
	mBase = m.M
	if v2965 < int32(0) {
		v2971 = v2965
		goto L802
	} else {
		goto L845
	}
L845:
	;
	goto L807
L846:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2237
	v3824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3824
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3824
	v3828 = v3824 - int32(1)
	if v3828 <= v2237 {
		goto L1055
	} else {
		goto L1056
	}
L847:
	;
	v2975 = v2973
	goto L849
L848:
	;
	v2975 = int32(38)
	goto L849
L849:
	;
	if v2975 == int32(0) {
		goto L846
	} else {
		goto L850
	}
L850:
	;
	if v2971 < int32(0) {
		goto L851
	} else {
		goto L852
	}
L851:
	;
	v2980 = v2971
	goto L853
L852:
	;
	v2980 = v2232
	goto L853
L853:
	;
	if v2971 != 0 {
		goto L854
	} else {
		goto L855
	}
L854:
	;
	v2981 = v2980
	goto L856
L855:
	;
	v2981 = v2232
	goto L856
L856:
	;
	if v2975 != int32(38) {
		v3815 = v2981
		v3817 = v2973
		goto L857
	} else {
		goto L858
	}
L857:
	;
	if v3817 != 0 {
		v3981 = v3815
		goto L44
	} else {
		goto L1054
	}
L858:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2765
	v2985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v2985 == int32(0) {
		goto L859
	} else {
		goto L860
	}
L859:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2765
	v3304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v3304 == int32(0) {
		goto L846
	} else {
		goto L937
	}
L860:
	;
	v2988 = int32(0)
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2989
	v2992 = v2989 + int32(1)
	v2993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2993 <= v2992 {
		v3287 = v2988
		goto L861
	} else {
		goto L862
	}
L861:
	;
	v3289 = int32(base.Ui32(v3287) >> (uint(int32(31)) % 32))
	if v3287 != 0 {
		goto L926
	} else {
		goto L927
	}
L862:
	;
	v2995 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2995+v2992))))
	if base.B2i32(v2997 != int32(168))&base.B2i32(v2997 != int32(131)) != 0 {
		v3287 = v2988
		goto L861
	} else {
		goto L863
	}
L863:
	;
	v3006 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_71), int32(4), int32(0))
	mBase = m.M
	v3007 = m.ExcPending
	if v3007 != 0 {
		goto L4
	} else {
		goto L864
	}
L864:
	;
	if v3006 == int32(0) {
		v3287 = v2988
		goto L861
	} else {
		goto L865
	}
L865:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3010
	switch v3006 - int32(1) {
	case 0:
		goto L869
	case 1:
		goto L868
	case 2:
		goto L867
	default:
		goto L866
	}
L866:
	;
	v3287 = int32(1)
	goto L861
L867:
	;
	v3193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3194 = int32(0)
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v3193-int32(4))))
	if v3201 == v3194 {
		goto L908
	} else {
		goto L909
	}
L868:
	;
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3103 = int32(0)
	v3110 = *(*int32)(unsafe.Add(mBase, uint32(v3102-int32(4))))
	if v3110 == v3103 {
		goto L889
	} else {
		goto L890
	}
L869:
	;
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3015 = int32(0)
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v3014-int32(4))))
	if v3022 == v3015 {
		goto L871
	} else {
		goto L872
	}
L870:
	;
	if v3096 < int32(4) {
		v3287 = v2988
		goto L861
	} else {
		goto L886
	}
L871:
	;
	v3096 = int32(0)
	goto L870
L872:
	;
	goto L873
L873:
	;
	v3027 = v3022 & int32(3)
	if base.Ui32(v3022) < base.Ui32(int32(4)) {
		goto L876
	} else {
		goto L877
	}
L874:
	;
	v3096 = v3085
	goto L870
L875:
	;
	v3069 = v3063
	v3070 = v3064
	v3074 = v3015
	goto L883
L876:
	;
	v3063 = v3014
	v3064 = int32(0)
	goto L875
L877:
	;
	goto L878
L878:
	;
	v3034 = v3014
	v3035 = int32(0)
	v3038 = v3015
	goto L879
L879:
	;
	v3040 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3034))))
	v3041 = int32(-65)
	v3044 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3034)+1)))
	v3048 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3034)+2)))
	v3052 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3034)+3)))
	v3055 = v3035 + base.B2i32(v3041 < v3040) + base.B2i32(v3041 < v3044) + base.B2i32(v3041 < v3048) + base.B2i32(v3041 < v3052)
	v3056 = int32(4)
	v3057 = v3034 + v3056
	v3059 = v3038 + v3056
	if v3059 != v3022&int32(-4) {
		v3034 = v3057
		v3035 = v3055
		v3038 = v3059
		goto L879
	} else {
		goto L881
	}
L880:
	;
	if v3027 == int32(0) {
		v3085 = v3055
		goto L874
	} else {
		goto L882
	}
L881:
	;
	goto L880
L882:
	;
	v3063 = v3057
	v3064 = v3055
	goto L875
L883:
	;
	v3075 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3069))))
	v3078 = v3070 + base.B2i32(int32(-65) < v3075)
	v3079 = int32(1)
	v3082 = v3074 + v3079
	if v3082 != v3027 {
		v3069 = v3069 + v3079
		v3070 = v3078
		v3074 = v3082
		goto L883
	} else {
		goto L885
	}
L884:
	;
	v3085 = v3078
	goto L874
L885:
	;
	goto L884
L886:
	;
	v3099 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v3099 {
		goto L866
	} else {
		goto L887
	}
L887:
	;
	v3287 = v3099
	goto L861
L888:
	;
	if v3184 < int32(4) {
		v3287 = v2988
		goto L861
	} else {
		goto L904
	}
L889:
	;
	v3184 = int32(0)
	goto L888
L890:
	;
	goto L891
L891:
	;
	v3115 = v3110 & int32(3)
	if base.Ui32(v3110) < base.Ui32(int32(4)) {
		goto L894
	} else {
		goto L895
	}
L892:
	;
	v3184 = v3173
	goto L888
L893:
	;
	v3157 = v3151
	v3158 = v3152
	v3162 = v3103
	goto L901
L894:
	;
	v3151 = v3102
	v3152 = int32(0)
	goto L893
L895:
	;
	goto L896
L896:
	;
	v3122 = v3102
	v3123 = int32(0)
	v3126 = v3103
	goto L897
L897:
	;
	v3128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3122))))
	v3129 = int32(-65)
	v3132 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3122)+1)))
	v3136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3122)+2)))
	v3140 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3122)+3)))
	v3143 = v3123 + base.B2i32(v3129 < v3128) + base.B2i32(v3129 < v3132) + base.B2i32(v3129 < v3136) + base.B2i32(v3129 < v3140)
	v3144 = int32(4)
	v3145 = v3122 + v3144
	v3147 = v3126 + v3144
	if v3147 != v3110&int32(-4) {
		v3122 = v3145
		v3123 = v3143
		v3126 = v3147
		goto L897
	} else {
		goto L899
	}
L898:
	;
	if v3115 == int32(0) {
		v3173 = v3143
		goto L892
	} else {
		goto L900
	}
L899:
	;
	goto L898
L900:
	;
	v3151 = v3145
	v3152 = v3143
	goto L893
L901:
	;
	v3163 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3157))))
	v3166 = v3158 + base.B2i32(int32(-65) < v3163)
	v3167 = int32(1)
	v3170 = v3162 + v3167
	if v3170 != v3115 {
		v3157 = v3157 + v3167
		v3158 = v3166
		v3162 = v3170
		goto L901
	} else {
		goto L903
	}
L902:
	;
	v3173 = v3166
	goto L892
L903:
	;
	goto L902
L904:
	;
	v3189 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_72))
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L4
	} else {
		goto L905
	}
L905:
	;
	if int32(0) <= v3189 {
		goto L866
	} else {
		goto L906
	}
L906:
	;
	v3287 = v3189
	goto L861
L907:
	;
	if v3275 < int32(4) {
		v3287 = v2988
		goto L861
	} else {
		goto L923
	}
L908:
	;
	v3275 = int32(0)
	goto L907
L909:
	;
	goto L910
L910:
	;
	v3206 = v3201 & int32(3)
	if base.Ui32(v3201) < base.Ui32(int32(4)) {
		goto L913
	} else {
		goto L914
	}
L911:
	;
	v3275 = v3264
	goto L907
L912:
	;
	v3248 = v3242
	v3249 = v3243
	v3253 = v3194
	goto L920
L913:
	;
	v3242 = v3193
	v3243 = int32(0)
	goto L912
L914:
	;
	goto L915
L915:
	;
	v3213 = v3193
	v3214 = int32(0)
	v3217 = v3194
	goto L916
L916:
	;
	v3219 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3213))))
	v3220 = int32(-65)
	v3223 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3213)+1)))
	v3227 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3213)+2)))
	v3231 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3213)+3)))
	v3234 = v3214 + base.B2i32(v3220 < v3219) + base.B2i32(v3220 < v3223) + base.B2i32(v3220 < v3227) + base.B2i32(v3220 < v3231)
	v3235 = int32(4)
	v3236 = v3213 + v3235
	v3238 = v3217 + v3235
	if v3238 != v3201&int32(-4) {
		v3213 = v3236
		v3214 = v3234
		v3217 = v3238
		goto L916
	} else {
		goto L918
	}
L917:
	;
	if v3206 == int32(0) {
		v3264 = v3234
		goto L911
	} else {
		goto L919
	}
L918:
	;
	goto L917
L919:
	;
	v3242 = v3236
	v3243 = v3234
	goto L912
L920:
	;
	v3254 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3248))))
	v3257 = v3249 + base.B2i32(int32(-65) < v3254)
	v3258 = int32(1)
	v3261 = v3253 + v3258
	if v3261 != v3206 {
		v3248 = v3248 + v3258
		v3249 = v3257
		v3253 = v3261
		goto L920
	} else {
		goto L922
	}
L921:
	;
	v3264 = v3257
	goto L911
L922:
	;
	goto L921
L923:
	;
	v3280 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_73))
	mBase = m.M
	v3281 = m.ExcPending
	if v3281 != 0 {
		goto L4
	} else {
		goto L924
	}
L924:
	;
	if v3280 < int32(0) {
		v3287 = v3280
		goto L861
	} else {
		goto L925
	}
L925:
	;
	goto L866
L926:
	;
	v3291 = v3289
	goto L928
L927:
	;
	v3291 = int32(39)
	goto L928
L928:
	;
	if v3291 == int32(0) {
		goto L846
	} else {
		goto L929
	}
L929:
	;
	if v3291 == int32(39) {
		goto L859
	} else {
		goto L930
	}
L930:
	;
	if v3287 < int32(0) {
		goto L931
	} else {
		goto L932
	}
L931:
	;
	v3298 = v3287
	goto L933
L932:
	;
	v3298 = v2981
	goto L933
L933:
	;
	if v3287 != 0 {
		goto L934
	} else {
		goto L935
	}
L934:
	;
	v3299 = v3298
	goto L936
L935:
	;
	v3299 = v2981
	goto L936
L936:
	;
	v3815 = v3299
	v3817 = v3289
	goto L857
L937:
	;
	v3307 = int32(0)
	v3308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3308
	v3313 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_74), int32(4), v3307)
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L4
	} else {
		goto L939
	}
L938:
	;
	if v3687 == int32(0) {
		goto L1023
	} else {
		goto L1024
	}
L939:
	;
	if v3313 == int32(0) {
		v3687 = v3307
		goto L938
	} else {
		goto L940
	}
L940:
	;
	v3317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3317
	switch v3313 - int32(1) {
	case 0:
		goto L945
	case 1:
		goto L944
	case 2:
		goto L943
	case 3:
		goto L942
	default:
		goto L941
	}
L941:
	;
	v3687 = int32(1)
	goto L938
L942:
	;
	v3594 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3595 = int32(0)
	v3602 = *(*int32)(unsafe.Add(mBase, uint32(v3594-int32(4))))
	if v3602 == v3595 {
		goto L1004
	} else {
		goto L1005
	}
L943:
	;
	v3503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3504 = int32(0)
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v3503-int32(4))))
	if v3511 == v3504 {
		goto L985
	} else {
		goto L986
	}
L944:
	;
	v3412 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3413 = int32(0)
	v3420 = *(*int32)(unsafe.Add(mBase, uint32(v3412-int32(4))))
	if v3420 == v3413 {
		goto L966
	} else {
		goto L967
	}
L945:
	;
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3322 = int32(0)
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(v3321-int32(4))))
	if v3329 == v3322 {
		goto L947
	} else {
		goto L948
	}
L946:
	;
	if v3403 < int32(5) {
		v3687 = v3307
		goto L938
	} else {
		goto L962
	}
L947:
	;
	v3403 = int32(0)
	goto L946
L948:
	;
	goto L949
L949:
	;
	v3334 = v3329 & int32(3)
	if base.Ui32(v3329) < base.Ui32(int32(4)) {
		goto L952
	} else {
		goto L953
	}
L950:
	;
	v3403 = v3392
	goto L946
L951:
	;
	v3376 = v3370
	v3377 = v3371
	v3381 = v3322
	goto L959
L952:
	;
	v3370 = v3321
	v3371 = int32(0)
	goto L951
L953:
	;
	goto L954
L954:
	;
	v3341 = v3321
	v3342 = int32(0)
	v3345 = v3322
	goto L955
L955:
	;
	v3347 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3341))))
	v3348 = int32(-65)
	v3351 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3341)+1)))
	v3355 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3341)+2)))
	v3359 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3341)+3)))
	v3362 = v3342 + base.B2i32(v3348 < v3347) + base.B2i32(v3348 < v3351) + base.B2i32(v3348 < v3355) + base.B2i32(v3348 < v3359)
	v3363 = int32(4)
	v3364 = v3341 + v3363
	v3366 = v3345 + v3363
	if v3366 != v3329&int32(-4) {
		v3341 = v3364
		v3342 = v3362
		v3345 = v3366
		goto L955
	} else {
		goto L957
	}
L956:
	;
	if v3334 == int32(0) {
		v3392 = v3362
		goto L950
	} else {
		goto L958
	}
L957:
	;
	goto L956
L958:
	;
	v3370 = v3364
	v3371 = v3362
	goto L951
L959:
	;
	v3382 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3376))))
	v3385 = v3377 + base.B2i32(int32(-65) < v3382)
	v3386 = int32(1)
	v3389 = v3381 + v3386
	if v3389 != v3334 {
		v3376 = v3376 + v3386
		v3377 = v3385
		v3381 = v3389
		goto L959
	} else {
		goto L961
	}
L960:
	;
	v3392 = v3385
	goto L950
L961:
	;
	goto L960
L962:
	;
	v3408 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_75))
	mBase = m.M
	v3409 = m.ExcPending
	if v3409 != 0 {
		goto L4
	} else {
		goto L963
	}
L963:
	;
	if int32(0) <= v3408 {
		goto L941
	} else {
		goto L964
	}
L964:
	;
	v3687 = v3408
	goto L938
L965:
	;
	if v3494 < int32(5) {
		v3687 = v3307
		goto L938
	} else {
		goto L981
	}
L966:
	;
	v3494 = int32(0)
	goto L965
L967:
	;
	goto L968
L968:
	;
	v3425 = v3420 & int32(3)
	if base.Ui32(v3420) < base.Ui32(int32(4)) {
		goto L971
	} else {
		goto L972
	}
L969:
	;
	v3494 = v3483
	goto L965
L970:
	;
	v3467 = v3461
	v3468 = v3462
	v3472 = v3413
	goto L978
L971:
	;
	v3461 = v3412
	v3462 = int32(0)
	goto L970
L972:
	;
	goto L973
L973:
	;
	v3432 = v3412
	v3433 = int32(0)
	v3436 = v3413
	goto L974
L974:
	;
	v3438 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3432))))
	v3439 = int32(-65)
	v3442 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3432)+1)))
	v3446 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3432)+2)))
	v3450 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3432)+3)))
	v3453 = v3433 + base.B2i32(v3439 < v3438) + base.B2i32(v3439 < v3442) + base.B2i32(v3439 < v3446) + base.B2i32(v3439 < v3450)
	v3454 = int32(4)
	v3455 = v3432 + v3454
	v3457 = v3436 + v3454
	if v3457 != v3420&int32(-4) {
		v3432 = v3455
		v3433 = v3453
		v3436 = v3457
		goto L974
	} else {
		goto L976
	}
L975:
	;
	if v3425 == int32(0) {
		v3483 = v3453
		goto L969
	} else {
		goto L977
	}
L976:
	;
	goto L975
L977:
	;
	v3461 = v3455
	v3462 = v3453
	goto L970
L978:
	;
	v3473 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3467))))
	v3476 = v3468 + base.B2i32(int32(-65) < v3473)
	v3477 = int32(1)
	v3480 = v3472 + v3477
	if v3480 != v3425 {
		v3467 = v3467 + v3477
		v3468 = v3476
		v3472 = v3480
		goto L978
	} else {
		goto L980
	}
L979:
	;
	v3483 = v3476
	goto L969
L980:
	;
	goto L979
L981:
	;
	v3499 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_76))
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L4
	} else {
		goto L982
	}
L982:
	;
	if int32(0) <= v3499 {
		goto L941
	} else {
		goto L983
	}
L983:
	;
	v3687 = v3499
	goto L938
L984:
	;
	if v3585 < int32(5) {
		v3687 = v3307
		goto L938
	} else {
		goto L1000
	}
L985:
	;
	v3585 = int32(0)
	goto L984
L986:
	;
	goto L987
L987:
	;
	v3516 = v3511 & int32(3)
	if base.Ui32(v3511) < base.Ui32(int32(4)) {
		goto L990
	} else {
		goto L991
	}
L988:
	;
	v3585 = v3574
	goto L984
L989:
	;
	v3558 = v3552
	v3559 = v3553
	v3563 = v3504
	goto L997
L990:
	;
	v3552 = v3503
	v3553 = int32(0)
	goto L989
L991:
	;
	goto L992
L992:
	;
	v3523 = v3503
	v3524 = int32(0)
	v3527 = v3504
	goto L993
L993:
	;
	v3529 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3523))))
	v3530 = int32(-65)
	v3533 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3523)+1)))
	v3537 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3523)+2)))
	v3541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3523)+3)))
	v3544 = v3524 + base.B2i32(v3530 < v3529) + base.B2i32(v3530 < v3533) + base.B2i32(v3530 < v3537) + base.B2i32(v3530 < v3541)
	v3545 = int32(4)
	v3546 = v3523 + v3545
	v3548 = v3527 + v3545
	if v3548 != v3511&int32(-4) {
		v3523 = v3546
		v3524 = v3544
		v3527 = v3548
		goto L993
	} else {
		goto L995
	}
L994:
	;
	if v3516 == int32(0) {
		v3574 = v3544
		goto L988
	} else {
		goto L996
	}
L995:
	;
	goto L994
L996:
	;
	v3552 = v3546
	v3553 = v3544
	goto L989
L997:
	;
	v3564 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3558))))
	v3567 = v3559 + base.B2i32(int32(-65) < v3564)
	v3568 = int32(1)
	v3571 = v3563 + v3568
	if v3571 != v3516 {
		v3558 = v3558 + v3568
		v3559 = v3567
		v3563 = v3571
		goto L997
	} else {
		goto L999
	}
L998:
	;
	v3574 = v3567
	goto L988
L999:
	;
	goto L998
L1000:
	;
	v3590 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_77))
	mBase = m.M
	v3591 = m.ExcPending
	if v3591 != 0 {
		goto L4
	} else {
		goto L1001
	}
L1001:
	;
	if int32(0) <= v3590 {
		goto L941
	} else {
		goto L1002
	}
L1002:
	;
	v3687 = v3590
	goto L938
L1003:
	;
	if v3676 < int32(5) {
		v3687 = v3307
		goto L938
	} else {
		goto L1019
	}
L1004:
	;
	v3676 = int32(0)
	goto L1003
L1005:
	;
	goto L1006
L1006:
	;
	v3607 = v3602 & int32(3)
	if base.Ui32(v3602) < base.Ui32(int32(4)) {
		goto L1009
	} else {
		goto L1010
	}
L1007:
	;
	v3676 = v3665
	goto L1003
L1008:
	;
	v3649 = v3643
	v3650 = v3644
	v3654 = v3595
	goto L1016
L1009:
	;
	v3643 = v3594
	v3644 = int32(0)
	goto L1008
L1010:
	;
	goto L1011
L1011:
	;
	v3614 = v3594
	v3615 = int32(0)
	v3618 = v3595
	goto L1012
L1012:
	;
	v3620 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3614))))
	v3621 = int32(-65)
	v3624 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3614)+1)))
	v3628 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3614)+2)))
	v3632 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3614)+3)))
	v3635 = v3615 + base.B2i32(v3621 < v3620) + base.B2i32(v3621 < v3624) + base.B2i32(v3621 < v3628) + base.B2i32(v3621 < v3632)
	v3636 = int32(4)
	v3637 = v3614 + v3636
	v3639 = v3618 + v3636
	if v3639 != v3602&int32(-4) {
		v3614 = v3637
		v3615 = v3635
		v3618 = v3639
		goto L1012
	} else {
		goto L1014
	}
L1013:
	;
	if v3607 == int32(0) {
		v3665 = v3635
		goto L1007
	} else {
		goto L1015
	}
L1014:
	;
	goto L1013
L1015:
	;
	v3643 = v3637
	v3644 = v3635
	goto L1008
L1016:
	;
	v3655 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3649))))
	v3658 = v3650 + base.B2i32(int32(-65) < v3655)
	v3659 = int32(1)
	v3662 = v3654 + v3659
	if v3662 != v3607 {
		v3649 = v3649 + v3659
		v3650 = v3658
		v3654 = v3662
		goto L1016
	} else {
		goto L1018
	}
L1017:
	;
	v3665 = v3658
	goto L1007
L1018:
	;
	goto L1017
L1019:
	;
	v3681 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_78))
	mBase = m.M
	v3682 = m.ExcPending
	if v3682 != 0 {
		goto L4
	} else {
		goto L1020
	}
L1020:
	;
	if v3681 < int32(0) {
		v3687 = v3681
		goto L938
	} else {
		goto L1021
	}
L1021:
	;
	goto L941
L1022:
	;
	v3693 = int32(0)
	v3694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3694
	v3697 = v3694 + int32(5)
	v3698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3698 <= v3697 {
		v3810 = v3693
		goto L1027
	} else {
		goto L1028
	}
L1023:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2765
	goto L1022
L1024:
	;
	goto L1025
L1025:
	;
	if v3687 < int32(0) {
		v3981 = v3687
		goto L44
	} else {
		goto L1026
	}
L1026:
	;
	goto L1022
L1027:
	;
	if int32(0) <= v3810 {
		goto L846
	} else {
		goto L1053
	}
L1028:
	;
	v3700 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3700+v3697))))
	if v3702 != int32(170) {
		v3810 = v3693
		goto L1027
	} else {
		goto L1029
	}
L1029:
	;
	v3708 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_79), int32(3), int32(0))
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L4
	} else {
		goto L1030
	}
L1030:
	;
	if v3708 == int32(0) {
		v3810 = v3693
		goto L1027
	} else {
		goto L1031
	}
L1031:
	;
	v3712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3712
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3715 = int32(0)
	v3722 = *(*int32)(unsafe.Add(mBase, uint32(v3714-int32(4))))
	if v3722 == v3715 {
		goto L1033
	} else {
		goto L1034
	}
L1032:
	;
	if v3796 < int32(5) {
		v3810 = v3693
		goto L1027
	} else {
		goto L1048
	}
L1033:
	;
	v3796 = int32(0)
	goto L1032
L1034:
	;
	goto L1035
L1035:
	;
	v3727 = v3722 & int32(3)
	if base.Ui32(v3722) < base.Ui32(int32(4)) {
		goto L1038
	} else {
		goto L1039
	}
L1036:
	;
	v3796 = v3785
	goto L1032
L1037:
	;
	v3769 = v3763
	v3770 = v3764
	v3774 = v3715
	goto L1045
L1038:
	;
	v3763 = v3714
	v3764 = int32(0)
	goto L1037
L1039:
	;
	goto L1040
L1040:
	;
	v3734 = v3714
	v3735 = int32(0)
	v3738 = v3715
	goto L1041
L1041:
	;
	v3740 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3734))))
	v3741 = int32(-65)
	v3744 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3734)+1)))
	v3748 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3734)+2)))
	v3752 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3734)+3)))
	v3755 = v3735 + base.B2i32(v3741 < v3740) + base.B2i32(v3741 < v3744) + base.B2i32(v3741 < v3748) + base.B2i32(v3741 < v3752)
	v3756 = int32(4)
	v3757 = v3734 + v3756
	v3759 = v3738 + v3756
	if v3759 != v3722&int32(-4) {
		v3734 = v3757
		v3735 = v3755
		v3738 = v3759
		goto L1041
	} else {
		goto L1043
	}
L1042:
	;
	if v3727 == int32(0) {
		v3785 = v3755
		goto L1036
	} else {
		goto L1044
	}
L1043:
	;
	goto L1042
L1044:
	;
	v3763 = v3757
	v3764 = v3755
	goto L1037
L1045:
	;
	v3775 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3769))))
	v3778 = v3770 + base.B2i32(int32(-65) < v3775)
	v3779 = int32(1)
	v3782 = v3774 + v3779
	if v3782 != v3727 {
		v3769 = v3769 + v3779
		v3770 = v3778
		v3774 = v3782
		goto L1045
	} else {
		goto L1047
	}
L1046:
	;
	v3785 = v3778
	goto L1036
L1047:
	;
	goto L1046
L1048:
	;
	v3799 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+29)) = uint16(v3799)
	v3804 = F_slice_from_s(m, l0, int32(6), int32(_a_F_arabic_UTF_8_stem_80))
	mBase = m.M
	v3805 = m.ExcPending
	if v3805 != 0 {
		goto L4
	} else {
		goto L1049
	}
L1049:
	;
	if int32(0) <= v3804 {
		goto L1050
	} else {
		goto L1051
	}
L1050:
	;
	v3808 = v3799
	goto L1052
L1051:
	;
	v3808 = v3804
	goto L1052
L1052:
	;
	v3810 = v3808
	goto L1027
L1053:
	;
	v3815 = v3810
	v3817 = int32(base.Ui32(v3810) >> (uint(int32(31)) % 32))
	goto L857
L1054:
	;
	goto L846
L1055:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2237
	v3862 = v2237
	goto L1062
L1056:
	;
	v3830 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3830+v3828))))
	if base.B2i32(v3832&int32(224) != int32(160))|base.B2i32(int32(1)<<(uint(v3832)%32)&int32(124) == int32(0)) != 0 {
		goto L1055
	} else {
		goto L1057
	}
L1057:
	;
	v3847 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_81), int32(5), int32(0))
	mBase = m.M
	v3848 = m.ExcPending
	if v3848 != 0 {
		goto L4
	} else {
		goto L1058
	}
L1058:
	;
	if v3847 == int32(0) {
		goto L1055
	} else {
		goto L1059
	}
L1059:
	;
	v3851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3851
	v3855 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_82))
	mBase = m.M
	v3856 = m.ExcPending
	if v3856 != 0 {
		goto L4
	} else {
		goto L1060
	}
L1060:
	;
	if v3855 < int32(0) {
		v3981 = v3855
		goto L44
	} else {
		goto L1061
	}
L1061:
	;
	goto L1055
L1062:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3862
	v3869 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3871 = v3862 + int32(1)
	v3872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3872 <= v3871 {
		v3895 = v3872
		v3896 = v3869
		goto L1066
	} else {
		goto L1067
	}
L1063:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2237
	v3981 = int32(1)
	goto L44
L1064:
	;
	goto L1063
L1065:
	;
	v3954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3954
	switch v3890 - int32(1) {
	case 0:
		goto L1095
	case 1:
		goto L1094
	case 2:
		goto L1093
	default:
		goto L1092
	}
L1066:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3862
	goto L1073
L1067:
	;
	v3875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3871+v3869))))
	if base.B2i32(v3875&int32(224) != int32(160))|base.B2i32(int32(1)<<(uint(v3875)%32)&int32(124) == int32(0)) != 0 {
		v3895 = v3872
		v3896 = v3869
		goto L1066
	} else {
		goto L1068
	}
L1068:
	;
	v3890 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_83), int32(5), int32(0))
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L4
	} else {
		goto L1069
	}
L1069:
	;
	if v3890 != 0 {
		goto L1065
	} else {
		goto L1070
	}
L1070:
	;
	v3892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3893 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3895 = v3892
	v3896 = v3893
	goto L1066
L1071:
	;
	if v3949 < int32(0) {
		goto L1064
	} else {
		goto L1091
	}
L1073:
	;
	goto L1074
L1074:
	;
	goto L1075
L1075:
	;
	v3904 = v3862
	v3906 = int32(1)
	goto L1078
L1077:
	;
	v3949 = v3934
	goto L1071
L1078:
	;
	if v3895 <= v3904 {
		goto L1080
	} else {
		goto L1081
	}
L1079:
	;
	goto L1077
L1080:
	;
	v3949 = int32(-1)
	goto L1071
L1081:
	;
	goto L1082
L1082:
	;
	v3911 = v3904 + int32(1)
	v3913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3896+v3904))))
	if base.Ui32(v3913) < base.Ui32(int32(192)) {
		v3934 = v3911
		goto L1083
	} else {
		goto L1084
	}
L1083:
	;
	v3935 = int32(1)
	if v3935 < v3906 {
		v3904 = v3934
		v3906 = v3906 - v3935
		goto L1078
	} else {
		goto L1090
	}
L1084:
	;
	if v3895 <= v3911 {
		v3934 = v3911
		goto L1083
	} else {
		goto L1085
	}
L1085:
	;
	v3920 = v3911
	goto L1086
L1086:
	;
	v3923 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3896+v3920))))
	if int32(-65) < v3923 {
		v3934 = v3920
		goto L1083
	} else {
		goto L1088
	}
L1087:
	;
	v3934 = v3895
	goto L1083
L1088:
	;
	v3927 = v3920 + int32(1)
	if v3927 != v3895 {
		v3920 = v3927
		goto L1086
	} else {
		goto L1089
	}
L1089:
	;
	goto L1087
L1090:
	;
	goto L1079
L1091:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3949
	v3862 = v3949
	goto L1062
L1092:
	;
	v3977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3862 = v3977
	goto L1062
L1093:
	;
	v3972 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_84))
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		goto L4
	} else {
		goto L1100
	}
L1094:
	;
	v3966 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_85))
	mBase = m.M
	v3967 = m.ExcPending
	if v3967 != 0 {
		goto L4
	} else {
		goto L1098
	}
L1095:
	;
	v3960 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_86))
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L4
	} else {
		goto L1096
	}
L1096:
	;
	if int32(0) <= v3960 {
		goto L1092
	} else {
		goto L1097
	}
L1097:
	;
	v3981 = v3960
	goto L44
L1098:
	;
	if int32(0) <= v3966 {
		goto L1092
	} else {
		goto L1099
	}
L1099:
	;
	v3981 = v3966
	goto L44
L1100:
	;
	if v3972 < int32(0) {
		v3981 = v3972
		goto L44
	} else {
		goto L1101
	}
L1101:
	;
	goto L1092
}
