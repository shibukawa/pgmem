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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
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
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
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
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
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
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
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
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1817 int32
	_ = v1817
	var v1822 int32
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v2001 int32
	_ = v2001
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2054 int32
	_ = v2054
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2105 int32
	_ = v2105
	var v2110 int32
	_ = v2110
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2225 int32
	_ = v2225
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2264 int32
	_ = v2264
	var v2269 int32
	_ = v2269
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2338 int32
	_ = v2338
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2355 int32
	_ = v2355
	var v2360 int32
	_ = v2360
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2377 int32
	_ = v2377
	var v2381 int32
	_ = v2381
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2429 int32
	_ = v2429
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2446 int32
	_ = v2446
	var v2451 int32
	_ = v2451
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2476 int32
	_ = v2476
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2506 int32
	_ = v2506
	var v2509 int32
	_ = v2509
	var v2520 int32
	_ = v2520
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2537 int32
	_ = v2537
	var v2542 int32
	_ = v2542
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2553 int32
	_ = v2553
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2559 int32
	_ = v2559
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2578 int32
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2600 int32
	_ = v2600
	var v2611 int32
	_ = v2611
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2623 int32
	_ = v2623
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2631 int32
	_ = v2631
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2650 int32
	_ = v2650
	var v2655 int32
	_ = v2655
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2672 int32
	_ = v2672
	var v2676 int32
	_ = v2676
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2713 int32
	_ = v2713
	var v2724 int32
	_ = v2724
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2742 int32
	_ = v2742
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2762 int32
	_ = v2762
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2774 int32
	_ = v2774
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2786 int32
	_ = v2786
	var v2791 int32
	_ = v2791
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2808 int32
	_ = v2808
	var v2812 int32
	_ = v2812
	var v2816 int32
	_ = v2816
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2846 int32
	_ = v2846
	var v2849 int32
	_ = v2849
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2875 int32
	_ = v2875
	var v2880 int32
	_ = v2880
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2897 int32
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2905 int32
	_ = v2905
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2912 int32
	_ = v2912
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2938 int32
	_ = v2938
	var v2949 int32
	_ = v2949
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2963 int32
	_ = v2963
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2996 int32
	_ = v2996
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3008 int32
	_ = v3008
	var v3013 int32
	_ = v3013
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3024 int32
	_ = v3024
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3034 int32
	_ = v3034
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3068 int32
	_ = v3068
	var v3071 int32
	_ = v3071
	var v3082 int32
	_ = v3082
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3097 int32
	_ = v3097
	var v3102 int32
	_ = v3102
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3119 int32
	_ = v3119
	var v3123 int32
	_ = v3123
	var v3127 int32
	_ = v3127
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3134 int32
	_ = v3134
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3157 int32
	_ = v3157
	var v3160 int32
	_ = v3160
	var v3171 int32
	_ = v3171
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3188 int32
	_ = v3188
	var v3193 int32
	_ = v3193
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3204 int32
	_ = v3204
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3210 int32
	_ = v3210
	var v3214 int32
	_ = v3214
	var v3218 int32
	_ = v3218
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3225 int32
	_ = v3225
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3248 int32
	_ = v3248
	var v3251 int32
	_ = v3251
	var v3262 int32
	_ = v3262
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3289 int32
	_ = v3289
	var v3292 int32
	_ = v3292
	var v3295 int32
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3304 int32
	_ = v3304
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3316 int32
	_ = v3316
	var v3321 int32
	_ = v3321
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3332 int32
	_ = v3332
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3338 int32
	_ = v3338
	var v3342 int32
	_ = v3342
	var v3346 int32
	_ = v3346
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3353 int32
	_ = v3353
	var v3357 int32
	_ = v3357
	var v3358 int32
	_ = v3358
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3376 int32
	_ = v3376
	var v3379 int32
	_ = v3379
	var v3390 int32
	_ = v3390
	var v3395 int32
	_ = v3395
	var v3396 int32
	_ = v3396
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3407 int32
	_ = v3407
	var v3412 int32
	_ = v3412
	var v3419 int32
	_ = v3419
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3433 int32
	_ = v3433
	var v3437 int32
	_ = v3437
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3448 int32
	_ = v3448
	var v3449 int32
	_ = v3449
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3467 int32
	_ = v3467
	var v3470 int32
	_ = v3470
	var v3481 int32
	_ = v3481
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3498 int32
	_ = v3498
	var v3503 int32
	_ = v3503
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3514 int32
	_ = v3514
	var v3516 int32
	_ = v3516
	var v3517 int32
	_ = v3517
	var v3520 int32
	_ = v3520
	var v3524 int32
	_ = v3524
	var v3528 int32
	_ = v3528
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3535 int32
	_ = v3535
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3550 int32
	_ = v3550
	var v3551 int32
	_ = v3551
	var v3554 int32
	_ = v3554
	var v3555 int32
	_ = v3555
	var v3558 int32
	_ = v3558
	var v3561 int32
	_ = v3561
	var v3572 int32
	_ = v3572
	var v3577 int32
	_ = v3577
	var v3578 int32
	_ = v3578
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3589 int32
	_ = v3589
	var v3594 int32
	_ = v3594
	var v3601 int32
	_ = v3601
	var v3602 int32
	_ = v3602
	var v3605 int32
	_ = v3605
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3611 int32
	_ = v3611
	var v3615 int32
	_ = v3615
	var v3619 int32
	_ = v3619
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3624 int32
	_ = v3624
	var v3626 int32
	_ = v3626
	var v3630 int32
	_ = v3630
	var v3631 int32
	_ = v3631
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3641 int32
	_ = v3641
	var v3642 int32
	_ = v3642
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3649 int32
	_ = v3649
	var v3652 int32
	_ = v3652
	var v3663 int32
	_ = v3663
	var v3668 int32
	_ = v3668
	var v3669 int32
	_ = v3669
	var v3674 int32
	_ = v3674
	var v3680 int32
	_ = v3680
	var v3681 int32
	_ = v3681
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3687 int32
	_ = v3687
	var v3689 int32
	_ = v3689
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3698 int32
	_ = v3698
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3708 int32
	_ = v3708
	var v3713 int32
	_ = v3713
	var v3720 int32
	_ = v3720
	var v3721 int32
	_ = v3721
	var v3724 int32
	_ = v3724
	var v3726 int32
	_ = v3726
	var v3727 int32
	_ = v3727
	var v3730 int32
	_ = v3730
	var v3734 int32
	_ = v3734
	var v3738 int32
	_ = v3738
	var v3741 int32
	_ = v3741
	var v3742 int32
	_ = v3742
	var v3743 int32
	_ = v3743
	var v3745 int32
	_ = v3745
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3755 int32
	_ = v3755
	var v3756 int32
	_ = v3756
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3764 int32
	_ = v3764
	var v3765 int32
	_ = v3765
	var v3768 int32
	_ = v3768
	var v3771 int32
	_ = v3771
	var v3782 int32
	_ = v3782
	var v3785 int32
	_ = v3785
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3795 int32
	_ = v3795
	var v3797 int32
	_ = v3797
	var v3802 int32
	_ = v3802
	var v3804 int32
	_ = v3804
	var v3811 int32
	_ = v3811
	var v3815 int32
	_ = v3815
	var v3817 int32
	_ = v3817
	var v3819 int32
	_ = v3819
	var v3833 int32
	_ = v3833
	var v3834 int32
	_ = v3834
	var v3837 int32
	_ = v3837
	var v3841 int32
	_ = v3841
	var v3842 int32
	_ = v3842
	var v3848 int32
	_ = v3848
	var v3855 int32
	_ = v3855
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3861 int32
	_ = v3861
	var v3875 int32
	_ = v3875
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3878 int32
	_ = v3878
	var v3880 int32
	_ = v3880
	var v3881 int32
	_ = v3881
	var v3889 int32
	_ = v3889
	var v3891 int32
	_ = v3891
	var v3896 int32
	_ = v3896
	var v3898 int32
	_ = v3898
	var v3905 int32
	_ = v3905
	var v3908 int32
	_ = v3908
	var v3912 int32
	_ = v3912
	var v3919 int32
	_ = v3919
	var v3920 int32
	_ = v3920
	var v3934 int32
	_ = v3934
	var v3939 int32
	_ = v3939
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
	var v3951 int32
	_ = v3951
	var v3952 int32
	_ = v3952
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3962 int32
	_ = v3962
	var v3966 int32
	_ = v3966
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(4294967296)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v13
	v16 = v13 + int32(3)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 <= v16 {
		v220 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13
	v223 = v13
	goto L45
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v16))))
	if base.B2i32(v21 != int32(167))&base.B2i32(v21 != int32(132)) != 0 {
		v220 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_0), int32(4))
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
		v220 = v2
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
		v220 = v37
		goto L1
	}
L7:
	;
	v213 = int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+8)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v214))) = int64(1)
	v220 = v213
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
	v220 = v40
	goto L1
L27:
	;
	if v209 < int32(4) {
		v220 = v126
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
	return v3966
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v223
	v232 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_1), int32(144))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v603
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v606 != 0 {
		goto L234
	} else {
		goto L235
	}
L47:
	;
	goto L46
L48:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v223 = v601
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L208
L53:
	;
	v538 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_2))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L204
	}
L54:
	;
	v532 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_3))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L202
	}
L55:
	;
	v526 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_4))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L4
	} else {
		goto L200
	}
L56:
	;
	v520 = F_slice_from_s(m, l0, int32(4), int32(_a_F_arabic_UTF_8_stem_5))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L198
	}
L57:
	;
	v514 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_6))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L196
	}
L58:
	;
	v508 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_7))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L194
	}
L59:
	;
	v502 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_8))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L192
	}
L60:
	;
	v496 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_9))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L190
	}
L61:
	;
	v490 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_10))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L188
	}
L62:
	;
	v484 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_11))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L186
	}
L63:
	;
	v478 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_12))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L184
	}
L64:
	;
	v472 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_13))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L182
	}
L65:
	;
	v466 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_14))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L180
	}
L66:
	;
	v460 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_15))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L178
	}
L67:
	;
	v454 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_16))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L176
	}
L68:
	;
	v448 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_17))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L174
	}
L69:
	;
	v442 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_18))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L172
	}
L70:
	;
	v436 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_19))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L170
	}
L71:
	;
	v430 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_20))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L168
	}
L72:
	;
	v424 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_21))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L166
	}
L73:
	;
	v418 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_22))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L164
	}
L74:
	;
	v412 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_23))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L162
	}
L75:
	;
	v406 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_24))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L160
	}
L76:
	;
	v400 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_25))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L158
	}
L77:
	;
	v394 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_26))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L156
	}
L78:
	;
	v388 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_27))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L154
	}
L79:
	;
	v382 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_28))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L152
	}
L80:
	;
	v376 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_29))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L150
	}
L81:
	;
	v370 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_30))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L148
	}
L82:
	;
	v364 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_31))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L146
	}
L83:
	;
	v358 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_32))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L144
	}
L84:
	;
	v352 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_33))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L142
	}
L85:
	;
	v346 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_34))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L140
	}
L86:
	;
	v340 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_35))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L138
	}
L87:
	;
	v334 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_36))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L136
	}
L88:
	;
	v328 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_37))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L134
	}
L89:
	;
	v322 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_38))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L132
	}
L90:
	;
	v316 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_39))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L130
	}
L91:
	;
	v310 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_40))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L128
	}
L92:
	;
	v304 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_41))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L126
	}
L93:
	;
	v298 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_42))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L124
	}
L94:
	;
	v292 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_43))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L122
	}
L95:
	;
	v286 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_44))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L120
	}
L96:
	;
	v280 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_45))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L118
	}
L97:
	;
	v274 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_46))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L116
	}
L98:
	;
	v268 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_47))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L114
	}
L99:
	;
	v262 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_48))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L112
	}
L100:
	;
	v256 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_49))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L110
	}
L101:
	;
	v250 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_50))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L108
	}
L102:
	;
	v244 = F_slice_from_s(m, l0, int32(1), int32(_a_F_arabic_UTF_8_stem_51))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L106
	}
L103:
	;
	v238 = F_slice_del(m, l0)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v238 {
		goto L48
	} else {
		goto L105
	}
L105:
	;
	v3966 = v238
	goto L44
L106:
	;
	if int32(0) <= v244 {
		goto L48
	} else {
		goto L107
	}
L107:
	;
	v3966 = v244
	goto L44
L108:
	;
	if int32(0) <= v250 {
		goto L48
	} else {
		goto L109
	}
L109:
	;
	v3966 = v250
	goto L44
L110:
	;
	if int32(0) <= v256 {
		goto L48
	} else {
		goto L111
	}
L111:
	;
	v3966 = v256
	goto L44
L112:
	;
	if int32(0) <= v262 {
		goto L48
	} else {
		goto L113
	}
L113:
	;
	v3966 = v262
	goto L44
L114:
	;
	if int32(0) <= v268 {
		goto L48
	} else {
		goto L115
	}
L115:
	;
	v3966 = v268
	goto L44
L116:
	;
	if int32(0) <= v274 {
		goto L48
	} else {
		goto L117
	}
L117:
	;
	v3966 = v274
	goto L44
L118:
	;
	if int32(0) <= v280 {
		goto L48
	} else {
		goto L119
	}
L119:
	;
	v3966 = v280
	goto L44
L120:
	;
	if int32(0) <= v286 {
		goto L48
	} else {
		goto L121
	}
L121:
	;
	v3966 = v286
	goto L44
L122:
	;
	if int32(0) <= v292 {
		goto L48
	} else {
		goto L123
	}
L123:
	;
	v3966 = v292
	goto L44
L124:
	;
	if int32(0) <= v298 {
		goto L48
	} else {
		goto L125
	}
L125:
	;
	v3966 = v298
	goto L44
L126:
	;
	if int32(0) <= v304 {
		goto L48
	} else {
		goto L127
	}
L127:
	;
	v3966 = v304
	goto L44
L128:
	;
	if int32(0) <= v310 {
		goto L48
	} else {
		goto L129
	}
L129:
	;
	v3966 = v310
	goto L44
L130:
	;
	if int32(0) <= v316 {
		goto L48
	} else {
		goto L131
	}
L131:
	;
	v3966 = v316
	goto L44
L132:
	;
	if int32(0) <= v322 {
		goto L48
	} else {
		goto L133
	}
L133:
	;
	v3966 = v322
	goto L44
L134:
	;
	if int32(0) <= v328 {
		goto L48
	} else {
		goto L135
	}
L135:
	;
	v3966 = v328
	goto L44
L136:
	;
	if int32(0) <= v334 {
		goto L48
	} else {
		goto L137
	}
L137:
	;
	v3966 = v334
	goto L44
L138:
	;
	if int32(0) <= v340 {
		goto L48
	} else {
		goto L139
	}
L139:
	;
	v3966 = v340
	goto L44
L140:
	;
	if int32(0) <= v346 {
		goto L48
	} else {
		goto L141
	}
L141:
	;
	v3966 = v346
	goto L44
L142:
	;
	if int32(0) <= v352 {
		goto L48
	} else {
		goto L143
	}
L143:
	;
	v3966 = v352
	goto L44
L144:
	;
	if int32(0) <= v358 {
		goto L48
	} else {
		goto L145
	}
L145:
	;
	v3966 = v358
	goto L44
L146:
	;
	if int32(0) <= v364 {
		goto L48
	} else {
		goto L147
	}
L147:
	;
	v3966 = v364
	goto L44
L148:
	;
	if int32(0) <= v370 {
		goto L48
	} else {
		goto L149
	}
L149:
	;
	v3966 = v370
	goto L44
L150:
	;
	if int32(0) <= v376 {
		goto L48
	} else {
		goto L151
	}
L151:
	;
	v3966 = v376
	goto L44
L152:
	;
	if int32(0) <= v382 {
		goto L48
	} else {
		goto L153
	}
L153:
	;
	v3966 = v382
	goto L44
L154:
	;
	if int32(0) <= v388 {
		goto L48
	} else {
		goto L155
	}
L155:
	;
	v3966 = v388
	goto L44
L156:
	;
	if int32(0) <= v394 {
		goto L48
	} else {
		goto L157
	}
L157:
	;
	v3966 = v394
	goto L44
L158:
	;
	if int32(0) <= v400 {
		goto L48
	} else {
		goto L159
	}
L159:
	;
	v3966 = v400
	goto L44
L160:
	;
	if int32(0) <= v406 {
		goto L48
	} else {
		goto L161
	}
L161:
	;
	v3966 = v406
	goto L44
L162:
	;
	if int32(0) <= v412 {
		goto L48
	} else {
		goto L163
	}
L163:
	;
	v3966 = v412
	goto L44
L164:
	;
	if int32(0) <= v418 {
		goto L48
	} else {
		goto L165
	}
L165:
	;
	v3966 = v418
	goto L44
L166:
	;
	if int32(0) <= v424 {
		goto L48
	} else {
		goto L167
	}
L167:
	;
	v3966 = v424
	goto L44
L168:
	;
	if int32(0) <= v430 {
		goto L48
	} else {
		goto L169
	}
L169:
	;
	v3966 = v430
	goto L44
L170:
	;
	if int32(0) <= v436 {
		goto L48
	} else {
		goto L171
	}
L171:
	;
	v3966 = v436
	goto L44
L172:
	;
	if int32(0) <= v442 {
		goto L48
	} else {
		goto L173
	}
L173:
	;
	v3966 = v442
	goto L44
L174:
	;
	if int32(0) <= v448 {
		goto L48
	} else {
		goto L175
	}
L175:
	;
	v3966 = v448
	goto L44
L176:
	;
	if int32(0) <= v454 {
		goto L48
	} else {
		goto L177
	}
L177:
	;
	v3966 = v454
	goto L44
L178:
	;
	if int32(0) <= v460 {
		goto L48
	} else {
		goto L179
	}
L179:
	;
	v3966 = v460
	goto L44
L180:
	;
	if int32(0) <= v466 {
		goto L48
	} else {
		goto L181
	}
L181:
	;
	v3966 = v466
	goto L44
L182:
	;
	if int32(0) <= v472 {
		goto L48
	} else {
		goto L183
	}
L183:
	;
	v3966 = v472
	goto L44
L184:
	;
	if int32(0) <= v478 {
		goto L48
	} else {
		goto L185
	}
L185:
	;
	v3966 = v478
	goto L44
L186:
	;
	if int32(0) <= v484 {
		goto L48
	} else {
		goto L187
	}
L187:
	;
	v3966 = v484
	goto L44
L188:
	;
	if int32(0) <= v490 {
		goto L48
	} else {
		goto L189
	}
L189:
	;
	v3966 = v490
	goto L44
L190:
	;
	if int32(0) <= v496 {
		goto L48
	} else {
		goto L191
	}
L191:
	;
	v3966 = v496
	goto L44
L192:
	;
	if int32(0) <= v502 {
		goto L48
	} else {
		goto L193
	}
L193:
	;
	v3966 = v502
	goto L44
L194:
	;
	if int32(0) <= v508 {
		goto L48
	} else {
		goto L195
	}
L195:
	;
	v3966 = v508
	goto L44
L196:
	;
	if int32(0) <= v514 {
		goto L48
	} else {
		goto L197
	}
L197:
	;
	v3966 = v514
	goto L44
L198:
	;
	if int32(0) <= v520 {
		goto L48
	} else {
		goto L199
	}
L199:
	;
	v3966 = v520
	goto L44
L200:
	;
	if int32(0) <= v526 {
		goto L48
	} else {
		goto L201
	}
L201:
	;
	v3966 = v526
	goto L44
L202:
	;
	if int32(0) <= v532 {
		goto L48
	} else {
		goto L203
	}
L203:
	;
	v3966 = v532
	goto L44
L204:
	;
	if int32(0) <= v538 {
		goto L48
	} else {
		goto L205
	}
L205:
	;
	v3966 = v538
	goto L44
L206:
	;
	if v596 < int32(0) {
		goto L47
	} else {
		goto L226
	}
L208:
	;
	goto L209
L209:
	;
	goto L210
L210:
	;
	v551 = v223
	v553 = int32(1)
	goto L213
L212:
	;
	v596 = v581
	goto L206
L213:
	;
	if v544 <= v551 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L212
L215:
	;
	v596 = int32(-1)
	goto L206
L216:
	;
	goto L217
L217:
	;
	v558 = v551 + int32(1)
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v551))))
	if base.Ui32(v560) < base.Ui32(int32(192)) {
		v581 = v558
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v582 = int32(1)
	if v582 < v553 {
		v551 = v581
		v553 = v553 - v582
		goto L213
	} else {
		goto L225
	}
L219:
	;
	if v544 <= v558 {
		v581 = v558
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v567 = v558
	goto L221
L221:
	;
	v570 = int32(*(*int8)(unsafe.Add(mBase, uint32(v543+v567))))
	if int32(-65) < v570 {
		v581 = v567
		goto L218
	} else {
		goto L223
	}
L222:
	;
	v581 = v544
	goto L218
L223:
	;
	v574 = v567 + int32(1)
	if v574 != v544 {
		v567 = v574
		goto L221
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	goto L214
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
	goto L48
L227:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2225
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2225
	v2229 = v2225 + int32(3)
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2230 <= v2229 {
		goto L699
	} else {
		goto L700
	}
L228:
	;
	v2213 = F_slice_del(m, l0)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L4
	} else {
		goto L693
	}
L229:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2187
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2187
	v2191 = v2187 - int32(1)
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2191 <= v2192 {
		v2220 = v2183
		goto L227
	} else {
		goto L687
	}
L230:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2078
	v2081 = v2078 - int32(1)
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2081 <= v2082 {
		goto L665
	} else {
		goto L666
	}
L231:
	;
	if v2071 != 0 {
		v3966 = v2069
		goto L44
	} else {
		goto L664
	}
L232:
	;
	if int32(0) <= v1316 {
		v2220 = v1195
		goto L227
	} else {
		goto L663
	}
L233:
	;
	if int32(0) <= v922 {
		v2220 = v927
		goto L227
	} else {
		goto L662
	}
L234:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v607
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v610 = int32(1)
	v613 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_52), int32(12))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L4
	} else {
		goto L237
	}
L235:
	;
	v1321 = v603
	v1322 = v220
	v1325 = v605
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1321
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1325)+8))
	if v1328 == int32(0) {
		v2183 = v1322
		goto L229
	} else {
		goto L415
	}
L237:
	;
	if v613 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v616 = v613
	v618 = v607
	v619 = v609
	v620 = v610
	goto L241
L239:
	;
	v912 = v607
	v913 = v609
	v914 = v610
	goto L240
L240:
	;
	v916 = v912 - v913
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v916 + v917
	if v914 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L241:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v622
	switch v616 - int32(1) {
	case 0:
		goto L247
	case 1:
		goto L246
	case 2:
		goto L245
	default:
		goto L244
	}
L242:
	;
	v912 = v904
	v913 = v905
	v914 = base.B2i32(int32(0) < v906)
	goto L240
L243:
	;
	goto L242
L244:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v894
	v897 = v620 - int32(1)
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v901 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_52), int32(12))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L4
	} else {
		goto L305
	}
L245:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v805 = int32(0)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v804-int32(4))))
	if v812 == v805 {
		goto L287
	} else {
		goto L288
	}
L246:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v716 = int32(0)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v715-int32(4))))
	if v723 == v716 {
		goto L268
	} else {
		goto L269
	}
L247:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v627 = int32(0)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v626-int32(4))))
	if v634 == v627 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	if v708 < int32(4) {
		v904 = v618
		v905 = v619
		v906 = v620
		goto L243
	} else {
		goto L264
	}
L249:
	;
	v708 = int32(0)
	goto L248
L250:
	;
	goto L251
L251:
	;
	v639 = v634 & int32(3)
	if base.Ui32(v634) < base.Ui32(int32(4)) {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	v708 = v697
	goto L248
L253:
	;
	v681 = v675
	v682 = v676
	v686 = v627
	goto L261
L254:
	;
	v675 = v626
	v676 = int32(0)
	goto L253
L255:
	;
	goto L256
L256:
	;
	v646 = v626
	v647 = int32(0)
	v650 = v627
	goto L257
L257:
	;
	v652 = int32(*(*int8)(unsafe.Add(mBase, uint32(v646))))
	v653 = int32(-65)
	v656 = int32(*(*int8)(unsafe.Add(mBase, uint32(v646)+1)))
	v660 = int32(*(*int8)(unsafe.Add(mBase, uint32(v646)+2)))
	v664 = int32(*(*int8)(unsafe.Add(mBase, uint32(v646)+3)))
	v667 = v647 + base.B2i32(v653 < v652) + base.B2i32(v653 < v656) + base.B2i32(v653 < v660) + base.B2i32(v653 < v664)
	v668 = int32(4)
	v669 = v646 + v668
	v671 = v650 + v668
	if v671 != v634&int32(-4) {
		v646 = v669
		v647 = v667
		v650 = v671
		goto L257
	} else {
		goto L259
	}
L258:
	;
	if v639 == int32(0) {
		v697 = v667
		goto L252
	} else {
		goto L260
	}
L259:
	;
	goto L258
L260:
	;
	v675 = v669
	v676 = v667
	goto L253
L261:
	;
	v687 = int32(*(*int8)(unsafe.Add(mBase, uint32(v681))))
	v690 = v682 + base.B2i32(int32(-65) < v687)
	v691 = int32(1)
	v694 = v686 + v691
	if v694 != v639 {
		v681 = v681 + v691
		v682 = v690
		v686 = v694
		goto L261
	} else {
		goto L263
	}
L262:
	;
	v697 = v690
	goto L252
L263:
	;
	goto L262
L264:
	;
	v711 = F_slice_del(m, l0)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	if int32(0) <= v711 {
		goto L244
	} else {
		goto L266
	}
L266:
	;
	v3966 = v711
	goto L44
L267:
	;
	if v797 < int32(5) {
		v904 = v618
		v905 = v619
		v906 = v620
		goto L243
	} else {
		goto L283
	}
L268:
	;
	v797 = int32(0)
	goto L267
L269:
	;
	goto L270
L270:
	;
	v728 = v723 & int32(3)
	if base.Ui32(v723) < base.Ui32(int32(4)) {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	v797 = v786
	goto L267
L272:
	;
	v770 = v764
	v771 = v765
	v775 = v716
	goto L280
L273:
	;
	v764 = v715
	v765 = int32(0)
	goto L272
L274:
	;
	goto L275
L275:
	;
	v735 = v715
	v736 = int32(0)
	v739 = v716
	goto L276
L276:
	;
	v741 = int32(*(*int8)(unsafe.Add(mBase, uint32(v735))))
	v742 = int32(-65)
	v745 = int32(*(*int8)(unsafe.Add(mBase, uint32(v735)+1)))
	v749 = int32(*(*int8)(unsafe.Add(mBase, uint32(v735)+2)))
	v753 = int32(*(*int8)(unsafe.Add(mBase, uint32(v735)+3)))
	v756 = v736 + base.B2i32(v742 < v741) + base.B2i32(v742 < v745) + base.B2i32(v742 < v749) + base.B2i32(v742 < v753)
	v757 = int32(4)
	v758 = v735 + v757
	v760 = v739 + v757
	if v760 != v723&int32(-4) {
		v735 = v758
		v736 = v756
		v739 = v760
		goto L276
	} else {
		goto L278
	}
L277:
	;
	if v728 == int32(0) {
		v786 = v756
		goto L271
	} else {
		goto L279
	}
L278:
	;
	goto L277
L279:
	;
	v764 = v758
	v765 = v756
	goto L272
L280:
	;
	v776 = int32(*(*int8)(unsafe.Add(mBase, uint32(v770))))
	v779 = v771 + base.B2i32(int32(-65) < v776)
	v780 = int32(1)
	v783 = v775 + v780
	if v783 != v728 {
		v770 = v770 + v780
		v771 = v779
		v775 = v783
		goto L280
	} else {
		goto L282
	}
L281:
	;
	v786 = v779
	goto L271
L282:
	;
	goto L281
L283:
	;
	v800 = F_slice_del(m, l0)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	if int32(0) <= v800 {
		goto L244
	} else {
		goto L285
	}
L285:
	;
	v3966 = v800
	goto L44
L286:
	;
	if v886 < int32(6) {
		v904 = v618
		v905 = v619
		v906 = v620
		goto L243
	} else {
		goto L302
	}
L287:
	;
	v886 = int32(0)
	goto L286
L288:
	;
	goto L289
L289:
	;
	v817 = v812 & int32(3)
	if base.Ui32(v812) < base.Ui32(int32(4)) {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	v886 = v875
	goto L286
L291:
	;
	v859 = v853
	v860 = v854
	v864 = v805
	goto L299
L292:
	;
	v853 = v804
	v854 = int32(0)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v824 = v804
	v825 = int32(0)
	v828 = v805
	goto L295
L295:
	;
	v830 = int32(*(*int8)(unsafe.Add(mBase, uint32(v824))))
	v831 = int32(-65)
	v834 = int32(*(*int8)(unsafe.Add(mBase, uint32(v824)+1)))
	v838 = int32(*(*int8)(unsafe.Add(mBase, uint32(v824)+2)))
	v842 = int32(*(*int8)(unsafe.Add(mBase, uint32(v824)+3)))
	v845 = v825 + base.B2i32(v831 < v830) + base.B2i32(v831 < v834) + base.B2i32(v831 < v838) + base.B2i32(v831 < v842)
	v846 = int32(4)
	v847 = v824 + v846
	v849 = v828 + v846
	if v849 != v812&int32(-4) {
		v824 = v847
		v825 = v845
		v828 = v849
		goto L295
	} else {
		goto L297
	}
L296:
	;
	if v817 == int32(0) {
		v875 = v845
		goto L290
	} else {
		goto L298
	}
L297:
	;
	goto L296
L298:
	;
	v853 = v847
	v854 = v845
	goto L291
L299:
	;
	v865 = int32(*(*int8)(unsafe.Add(mBase, uint32(v859))))
	v868 = v860 + base.B2i32(int32(-65) < v865)
	v869 = int32(1)
	v872 = v864 + v869
	if v872 != v817 {
		v859 = v859 + v869
		v860 = v868
		v864 = v872
		goto L299
	} else {
		goto L301
	}
L300:
	;
	v875 = v868
	goto L290
L301:
	;
	goto L300
L302:
	;
	v889 = F_slice_del(m, l0)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L4
	} else {
		goto L303
	}
L303:
	;
	if v889 < int32(0) {
		v3966 = v889
		goto L44
	} else {
		goto L304
	}
L304:
	;
	goto L244
L305:
	;
	if v901 != 0 {
		v616 = v901
		v618 = v894
		v619 = v898
		v620 = v897
		goto L241
	} else {
		goto L306
	}
L306:
	;
	v904 = v894
	v905 = v898
	v906 = v897
	goto L243
L307:
	;
	v922 = F_r_Suffix_Verb_Step2a(m, l0)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L4
	} else {
		goto L310
	}
L308:
	;
	v1194 = v917
	v1195 = v220
	goto L309
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1194
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1194
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1194-int32(3) <= v1199 {
		goto L386
	} else {
		goto L387
	}
L310:
	;
	if v922 < int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v926 = v922
	goto L313
L312:
	;
	v926 = v220
	goto L313
L313:
	;
	if v922 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v927 = v926
	goto L316
L315:
	;
	v927 = v220
	goto L316
L316:
	;
	if v922 != 0 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v933 = v932 + v916
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v933
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v933
	v937 = v933 - int32(1)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v937 <= v938 {
		goto L321
	} else {
		goto L322
	}
L318:
	;
	v931 = int32(base.Ui32(v922) >> (uint(int32(31)) % 32))
	goto L320
L319:
	;
	v931 = int32(7)
	goto L320
L320:
	;
	switch v931 {
	case 0:
		v2220 = v927
		goto L227
	default:
		goto L233
	case 7:
		goto L317
	}
L321:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1135 = v1134 + v916
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1135
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L368
L322:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940+v937))))
	if v942 != int32(136) {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v947 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_53), int32(2))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L4
	} else {
		goto L324
	}
L324:
	;
	if v947 == int32(0) {
		goto L321
	} else {
		goto L325
	}
L325:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v951
	switch v947 - int32(1) {
	case 0:
		goto L327
	case 1:
		goto L326
	default:
		v2220 = v927
		goto L227
	}
L326:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1045 = int32(0)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1044-int32(4))))
	if v1052 == v1045 {
		goto L348
	} else {
		goto L349
	}
L327:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v956 = int32(0)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v955-int32(4))))
	if v963 == v956 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	if v1037 < int32(4) {
		goto L321
	} else {
		goto L344
	}
L329:
	;
	v1037 = int32(0)
	goto L328
L330:
	;
	goto L331
L331:
	;
	v968 = v963 & int32(3)
	if base.Ui32(v963) < base.Ui32(int32(4)) {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	v1037 = v1026
	goto L328
L333:
	;
	v1010 = v1004
	v1011 = v1005
	v1015 = v956
	goto L341
L334:
	;
	v1004 = v955
	v1005 = int32(0)
	goto L333
L335:
	;
	goto L336
L336:
	;
	v975 = v955
	v976 = int32(0)
	v979 = v956
	goto L337
L337:
	;
	v981 = int32(*(*int8)(unsafe.Add(mBase, uint32(v975))))
	v982 = int32(-65)
	v985 = int32(*(*int8)(unsafe.Add(mBase, uint32(v975)+1)))
	v989 = int32(*(*int8)(unsafe.Add(mBase, uint32(v975)+2)))
	v993 = int32(*(*int8)(unsafe.Add(mBase, uint32(v975)+3)))
	v996 = v976 + base.B2i32(v982 < v981) + base.B2i32(v982 < v985) + base.B2i32(v982 < v989) + base.B2i32(v982 < v993)
	v997 = int32(4)
	v998 = v975 + v997
	v1000 = v979 + v997
	if v1000 != v963&int32(-4) {
		v975 = v998
		v976 = v996
		v979 = v1000
		goto L337
	} else {
		goto L339
	}
L338:
	;
	if v968 == int32(0) {
		v1026 = v996
		goto L332
	} else {
		goto L340
	}
L339:
	;
	goto L338
L340:
	;
	v1004 = v998
	v1005 = v996
	goto L333
L341:
	;
	v1016 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1010))))
	v1019 = v1011 + base.B2i32(int32(-65) < v1016)
	v1020 = int32(1)
	v1023 = v1015 + v1020
	if v1023 != v968 {
		v1010 = v1010 + v1020
		v1011 = v1019
		v1015 = v1023
		goto L341
	} else {
		goto L343
	}
L342:
	;
	v1026 = v1019
	goto L332
L343:
	;
	goto L342
L344:
	;
	v1040 = F_slice_del(m, l0)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L4
	} else {
		goto L345
	}
L345:
	;
	if int32(0) <= v1040 {
		v2220 = v927
		goto L227
	} else {
		goto L346
	}
L346:
	;
	v3966 = v1040
	goto L44
L347:
	;
	if v1126 < int32(6) {
		goto L321
	} else {
		goto L363
	}
L348:
	;
	v1126 = int32(0)
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1057 = v1052 & int32(3)
	if base.Ui32(v1052) < base.Ui32(int32(4)) {
		goto L353
	} else {
		goto L354
	}
L351:
	;
	v1126 = v1115
	goto L347
L352:
	;
	v1099 = v1093
	v1100 = v1094
	v1104 = v1045
	goto L360
L353:
	;
	v1093 = v1044
	v1094 = int32(0)
	goto L352
L354:
	;
	goto L355
L355:
	;
	v1064 = v1044
	v1065 = int32(0)
	v1068 = v1045
	goto L356
L356:
	;
	v1070 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1064))))
	v1071 = int32(-65)
	v1074 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1064)+1)))
	v1078 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1064)+2)))
	v1082 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1064)+3)))
	v1085 = v1065 + base.B2i32(v1071 < v1070) + base.B2i32(v1071 < v1074) + base.B2i32(v1071 < v1078) + base.B2i32(v1071 < v1082)
	v1086 = int32(4)
	v1087 = v1064 + v1086
	v1089 = v1068 + v1086
	if v1089 != v1052&int32(-4) {
		v1064 = v1087
		v1065 = v1085
		v1068 = v1089
		goto L356
	} else {
		goto L358
	}
L357:
	;
	if v1057 == int32(0) {
		v1115 = v1085
		goto L351
	} else {
		goto L359
	}
L358:
	;
	goto L357
L359:
	;
	v1093 = v1087
	v1094 = v1085
	goto L352
L360:
	;
	v1105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1099))))
	v1108 = v1100 + base.B2i32(int32(-65) < v1105)
	v1109 = int32(1)
	v1112 = v1104 + v1109
	if v1112 != v1057 {
		v1099 = v1099 + v1109
		v1100 = v1108
		v1104 = v1112
		goto L360
	} else {
		goto L362
	}
L361:
	;
	v1115 = v1108
	goto L351
L362:
	;
	goto L361
L363:
	;
	v1129 = F_slice_del(m, l0)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L4
	} else {
		goto L364
	}
L364:
	;
	if int32(0) <= v1129 {
		v2220 = v927
		goto L227
	} else {
		goto L365
	}
L365:
	;
	v3966 = v1129
	goto L44
L366:
	;
	if int32(0) <= v1190 {
		v2220 = v927
		goto L227
	} else {
		goto L385
	}
L368:
	;
	goto L369
L369:
	;
	goto L370
L370:
	;
	v1145 = v1135
	v1147 = int32(1)
	goto L373
L372:
	;
	v1190 = v1172
	goto L366
L373:
	;
	if v1145 <= v1138 {
		goto L375
	} else {
		goto L376
	}
L374:
	;
	goto L372
L375:
	;
	v1190 = int32(-1)
	goto L366
L376:
	;
	goto L377
L377:
	;
	v1152 = v1145 - int32(1)
	v1154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1137+v1152))))
	if base.B2i32(int32(0) <= v1154)|base.B2i32(v1152 <= v1138) != 0 {
		v1172 = v1152
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1176 = int32(1)
	if v1176 < v1147 {
		v1145 = v1172
		v1147 = v1147 - v1176
		goto L373
	} else {
		goto L384
	}
L379:
	;
	v1160 = v1152
	goto L380
L380:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+v1160))))
	if base.Ui32(int32(191)) < base.Ui32(v1165) {
		v1172 = v1160
		goto L378
	} else {
		goto L382
	}
L381:
	;
	v1172 = v1138
	goto L378
L382:
	;
	v1169 = v1160 - int32(1)
	if v1138 < v1169 {
		v1160 = v1169
		goto L380
	} else {
		goto L383
	}
L383:
	;
	goto L381
L384:
	;
	goto L374
L385:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1194 = v1193
	v1195 = v927
	goto L309
L386:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1314
	v1316 = F_r_Suffix_Verb_Step2a(m, l0)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L4
	} else {
		goto L413
	}
L387:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203+v1194-int32(1)))))
	if base.B2i32(v1207 != int32(167))&base.B2i32(v1207 != int32(133)) != 0 {
		goto L386
	} else {
		goto L388
	}
L388:
	;
	v1215 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_54), int32(2))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L4
	} else {
		goto L389
	}
L389:
	;
	if v1215 == int32(0) {
		goto L386
	} else {
		goto L390
	}
L390:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1219
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1222 = int32(0)
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1221-int32(4))))
	if v1229 == v1222 {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	if v1303 < int32(5) {
		goto L386
	} else {
		goto L407
	}
L392:
	;
	v1303 = int32(0)
	goto L391
L393:
	;
	goto L394
L394:
	;
	v1234 = v1229 & int32(3)
	if base.Ui32(v1229) < base.Ui32(int32(4)) {
		goto L397
	} else {
		goto L398
	}
L395:
	;
	v1303 = v1292
	goto L391
L396:
	;
	v1276 = v1270
	v1277 = v1271
	v1281 = v1222
	goto L404
L397:
	;
	v1270 = v1221
	v1271 = int32(0)
	goto L396
L398:
	;
	goto L399
L399:
	;
	v1241 = v1221
	v1242 = int32(0)
	v1245 = v1222
	goto L400
L400:
	;
	v1247 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1241))))
	v1248 = int32(-65)
	v1251 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1241)+1)))
	v1255 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1241)+2)))
	v1259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1241)+3)))
	v1262 = v1242 + base.B2i32(v1248 < v1247) + base.B2i32(v1248 < v1251) + base.B2i32(v1248 < v1255) + base.B2i32(v1248 < v1259)
	v1263 = int32(4)
	v1264 = v1241 + v1263
	v1266 = v1245 + v1263
	if v1266 != v1229&int32(-4) {
		v1241 = v1264
		v1242 = v1262
		v1245 = v1266
		goto L400
	} else {
		goto L402
	}
L401:
	;
	if v1234 == int32(0) {
		v1292 = v1262
		goto L395
	} else {
		goto L403
	}
L402:
	;
	goto L401
L403:
	;
	v1270 = v1264
	v1271 = v1262
	goto L396
L404:
	;
	v1282 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1276))))
	v1285 = v1277 + base.B2i32(int32(-65) < v1282)
	v1286 = int32(1)
	v1289 = v1281 + v1286
	if v1289 != v1234 {
		v1276 = v1276 + v1286
		v1277 = v1285
		v1281 = v1289
		goto L404
	} else {
		goto L406
	}
L405:
	;
	v1292 = v1285
	goto L395
L406:
	;
	goto L405
L407:
	;
	v1306 = F_slice_del(m, l0)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L4
	} else {
		goto L408
	}
L408:
	;
	if v1306 < int32(0) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1310 = v1306
	goto L411
L410:
	;
	v1310 = v1195
	goto L411
L411:
	;
	if int32(0) <= v1306 {
		v2220 = v1310
		goto L227
	} else {
		goto L412
	}
L412:
	;
	v3966 = v1310
	goto L44
L413:
	;
	if v1316 != 0 {
		goto L232
	} else {
		goto L414
	}
L414:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1321 = v1319
	v1322 = v1195
	v1325 = v1318
	goto L236
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1321
	v1333 = v1321 - int32(1)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1333 <= v1334 {
		goto L417
	} else {
		goto L418
	}
L416:
	;
	if v2062 != 0 {
		v3966 = v2060
		goto L44
	} else {
		goto L661
	}
L417:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1443
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1445)))
	if v1446 != 0 {
		goto L447
	} else {
		goto L448
	}
L418:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336+v1333))))
	if v1338 != int32(169) {
		goto L417
	} else {
		goto L419
	}
L419:
	;
	v1343 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_55), int32(1))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L4
	} else {
		goto L420
	}
L420:
	;
	if v1343 == int32(0) {
		goto L417
	} else {
		goto L421
	}
L421:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1347
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1350 = int32(0)
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1349-int32(4))))
	if v1357 == v1350 {
		goto L423
	} else {
		goto L424
	}
L422:
	;
	if v1431 < int32(4) {
		goto L417
	} else {
		goto L438
	}
L423:
	;
	v1431 = int32(0)
	goto L422
L424:
	;
	goto L425
L425:
	;
	v1362 = v1357 & int32(3)
	if base.Ui32(v1357) < base.Ui32(int32(4)) {
		goto L428
	} else {
		goto L429
	}
L426:
	;
	v1431 = v1420
	goto L422
L427:
	;
	v1404 = v1398
	v1405 = v1399
	v1409 = v1350
	goto L435
L428:
	;
	v1398 = v1349
	v1399 = int32(0)
	goto L427
L429:
	;
	goto L430
L430:
	;
	v1369 = v1349
	v1370 = int32(0)
	v1373 = v1350
	goto L431
L431:
	;
	v1375 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1369))))
	v1376 = int32(-65)
	v1379 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1369)+1)))
	v1383 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1369)+2)))
	v1387 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1369)+3)))
	v1390 = v1370 + base.B2i32(v1376 < v1375) + base.B2i32(v1376 < v1379) + base.B2i32(v1376 < v1383) + base.B2i32(v1376 < v1387)
	v1391 = int32(4)
	v1392 = v1369 + v1391
	v1394 = v1373 + v1391
	if v1394 != v1357&int32(-4) {
		v1369 = v1392
		v1370 = v1390
		v1373 = v1394
		goto L431
	} else {
		goto L433
	}
L432:
	;
	if v1362 == int32(0) {
		v1420 = v1390
		goto L426
	} else {
		goto L434
	}
L433:
	;
	goto L432
L434:
	;
	v1398 = v1392
	v1399 = v1390
	goto L427
L435:
	;
	v1410 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1404))))
	v1413 = v1405 + base.B2i32(int32(-65) < v1410)
	v1414 = int32(1)
	v1417 = v1409 + v1414
	if v1417 != v1362 {
		v1404 = v1404 + v1414
		v1405 = v1413
		v1409 = v1417
		goto L435
	} else {
		goto L437
	}
L436:
	;
	v1420 = v1413
	goto L426
L437:
	;
	goto L436
L438:
	;
	v1434 = F_slice_del(m, l0)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L4
	} else {
		goto L439
	}
L439:
	;
	if v1434 < int32(0) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1438 = v1434
	goto L442
L441:
	;
	v1438 = v1322
	goto L442
L442:
	;
	if int32(0) <= v1434 {
		v2073 = v1438
		goto L230
	} else {
		goto L443
	}
L443:
	;
	v2060 = v1438
	v2062 = int32(base.Ui32(v1434) >> (uint(int32(31)) % 32))
	goto L416
L444:
	;
	if int32(0) <= v2047 {
		v2073 = v2042
		goto L230
	} else {
		goto L660
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2054
	v2073 = v2050
	goto L230
L446:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1855
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1855
	v1859 = v1855 - int32(1)
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1859 <= v1860 {
		v2018 = v1850
		goto L585
	} else {
		goto L586
	}
L447:
	;
	v1850 = v1322
	goto L446
L448:
	;
	goto L449
L449:
	;
	v1447 = int32(0)
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1448
	v1452 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_56), int32(10))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L4
	} else {
		goto L451
	}
L450:
	;
	if v1729 < int32(0) {
		goto L514
	} else {
		goto L515
	}
L451:
	;
	if v1452 == int32(0) {
		v1729 = v1447
		goto L450
	} else {
		goto L452
	}
L452:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1456
	switch v1452 - int32(1) {
	case 0:
		goto L456
	case 1:
		goto L455
	case 2:
		goto L454
	default:
		goto L453
	}
L453:
	;
	v1729 = int32(1)
	goto L450
L454:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1639 = int32(0)
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1638-int32(4))))
	if v1646 == v1639 {
		goto L496
	} else {
		goto L497
	}
L455:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1550 = int32(0)
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1549-int32(4))))
	if v1557 == v1550 {
		goto L477
	} else {
		goto L478
	}
L456:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1461 = int32(0)
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1460-int32(4))))
	if v1468 == v1461 {
		goto L458
	} else {
		goto L459
	}
L457:
	;
	if v1542 < int32(4) {
		v1729 = v1447
		goto L450
	} else {
		goto L473
	}
L458:
	;
	v1542 = int32(0)
	goto L457
L459:
	;
	goto L460
L460:
	;
	v1473 = v1468 & int32(3)
	if base.Ui32(v1468) < base.Ui32(int32(4)) {
		goto L463
	} else {
		goto L464
	}
L461:
	;
	v1542 = v1531
	goto L457
L462:
	;
	v1515 = v1509
	v1516 = v1510
	v1520 = v1461
	goto L470
L463:
	;
	v1509 = v1460
	v1510 = int32(0)
	goto L462
L464:
	;
	goto L465
L465:
	;
	v1480 = v1460
	v1481 = int32(0)
	v1484 = v1461
	goto L466
L466:
	;
	v1486 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1480))))
	v1487 = int32(-65)
	v1490 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1480)+1)))
	v1494 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1480)+2)))
	v1498 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1480)+3)))
	v1501 = v1481 + base.B2i32(v1487 < v1486) + base.B2i32(v1487 < v1490) + base.B2i32(v1487 < v1494) + base.B2i32(v1487 < v1498)
	v1502 = int32(4)
	v1503 = v1480 + v1502
	v1505 = v1484 + v1502
	if v1505 != v1468&int32(-4) {
		v1480 = v1503
		v1481 = v1501
		v1484 = v1505
		goto L466
	} else {
		goto L468
	}
L467:
	;
	if v1473 == int32(0) {
		v1531 = v1501
		goto L461
	} else {
		goto L469
	}
L468:
	;
	goto L467
L469:
	;
	v1509 = v1503
	v1510 = v1501
	goto L462
L470:
	;
	v1521 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1515))))
	v1524 = v1516 + base.B2i32(int32(-65) < v1521)
	v1525 = int32(1)
	v1528 = v1520 + v1525
	if v1528 != v1473 {
		v1515 = v1515 + v1525
		v1516 = v1524
		v1520 = v1528
		goto L470
	} else {
		goto L472
	}
L471:
	;
	v1531 = v1524
	goto L461
L472:
	;
	goto L471
L473:
	;
	v1545 = F_slice_del(m, l0)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L4
	} else {
		goto L474
	}
L474:
	;
	if int32(0) <= v1545 {
		goto L453
	} else {
		goto L475
	}
L475:
	;
	v1729 = v1545
	goto L450
L476:
	;
	if v1631 < int32(5) {
		v1729 = v1447
		goto L450
	} else {
		goto L492
	}
L477:
	;
	v1631 = int32(0)
	goto L476
L478:
	;
	goto L479
L479:
	;
	v1562 = v1557 & int32(3)
	if base.Ui32(v1557) < base.Ui32(int32(4)) {
		goto L482
	} else {
		goto L483
	}
L480:
	;
	v1631 = v1620
	goto L476
L481:
	;
	v1604 = v1598
	v1605 = v1599
	v1609 = v1550
	goto L489
L482:
	;
	v1598 = v1549
	v1599 = int32(0)
	goto L481
L483:
	;
	goto L484
L484:
	;
	v1569 = v1549
	v1570 = int32(0)
	v1573 = v1550
	goto L485
L485:
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
		goto L485
	} else {
		goto L487
	}
L486:
	;
	if v1562 == int32(0) {
		v1620 = v1590
		goto L480
	} else {
		goto L488
	}
L487:
	;
	goto L486
L488:
	;
	v1598 = v1592
	v1599 = v1590
	goto L481
L489:
	;
	v1610 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1604))))
	v1613 = v1605 + base.B2i32(int32(-65) < v1610)
	v1614 = int32(1)
	v1617 = v1609 + v1614
	if v1617 != v1562 {
		v1604 = v1604 + v1614
		v1605 = v1613
		v1609 = v1617
		goto L489
	} else {
		goto L491
	}
L490:
	;
	v1620 = v1613
	goto L480
L491:
	;
	goto L490
L492:
	;
	v1634 = F_slice_del(m, l0)
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L4
	} else {
		goto L493
	}
L493:
	;
	if int32(0) <= v1634 {
		goto L453
	} else {
		goto L494
	}
L494:
	;
	v1729 = v1634
	goto L450
L495:
	;
	if v1720 < int32(6) {
		v1729 = v1447
		goto L450
	} else {
		goto L511
	}
L496:
	;
	v1720 = int32(0)
	goto L495
L497:
	;
	goto L498
L498:
	;
	v1651 = v1646 & int32(3)
	if base.Ui32(v1646) < base.Ui32(int32(4)) {
		goto L501
	} else {
		goto L502
	}
L499:
	;
	v1720 = v1709
	goto L495
L500:
	;
	v1693 = v1687
	v1694 = v1688
	v1698 = v1639
	goto L508
L501:
	;
	v1687 = v1638
	v1688 = int32(0)
	goto L500
L502:
	;
	goto L503
L503:
	;
	v1658 = v1638
	v1659 = int32(0)
	v1662 = v1639
	goto L504
L504:
	;
	v1664 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1658))))
	v1665 = int32(-65)
	v1668 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1658)+1)))
	v1672 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1658)+2)))
	v1676 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1658)+3)))
	v1679 = v1659 + base.B2i32(v1665 < v1664) + base.B2i32(v1665 < v1668) + base.B2i32(v1665 < v1672) + base.B2i32(v1665 < v1676)
	v1680 = int32(4)
	v1681 = v1658 + v1680
	v1683 = v1662 + v1680
	if v1683 != v1646&int32(-4) {
		v1658 = v1681
		v1659 = v1679
		v1662 = v1683
		goto L504
	} else {
		goto L506
	}
L505:
	;
	if v1651 == int32(0) {
		v1709 = v1679
		goto L499
	} else {
		goto L507
	}
L506:
	;
	goto L505
L507:
	;
	v1687 = v1681
	v1688 = v1679
	goto L500
L508:
	;
	v1699 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1693))))
	v1702 = v1694 + base.B2i32(int32(-65) < v1699)
	v1703 = int32(1)
	v1706 = v1698 + v1703
	if v1706 != v1651 {
		v1693 = v1693 + v1703
		v1694 = v1702
		v1698 = v1706
		goto L508
	} else {
		goto L510
	}
L509:
	;
	v1709 = v1702
	goto L499
L510:
	;
	goto L509
L511:
	;
	v1723 = F_slice_del(m, l0)
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L4
	} else {
		goto L512
	}
L512:
	;
	if v1723 < int32(0) {
		v1729 = v1723
		goto L450
	} else {
		goto L513
	}
L513:
	;
	goto L453
L514:
	;
	v1732 = v1729
	goto L516
L515:
	;
	v1732 = v1322
	goto L516
L516:
	;
	if v1729 != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v1733 = v1732
	goto L519
L518:
	;
	v1733 = v1322
	goto L519
L519:
	;
	v1735 = int32(base.Ui32(v1729) >> (uint(int32(31)) % 32))
	if v1729 != 0 {
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v1737 = v1735
	goto L522
L521:
	;
	v1737 = int32(17)
	goto L522
L522:
	;
	if v1737 != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	if v1737 == int32(17) {
		goto L526
	} else {
		goto L527
	}
L524:
	;
	goto L525
L525:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1742 = F_r_Suffix_Noun_Step2a(m, l0)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L4
	} else {
		goto L529
	}
L526:
	;
	v1850 = v1733
	goto L446
L527:
	;
	v2060 = v1733
	v2062 = v1735
	goto L416
L529:
	;
	if v1742 < int32(0) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v1746 = v1742
	goto L532
L531:
	;
	v1746 = v1733
	goto L532
L532:
	;
	if v1742 != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v1747 = v1746
	goto L535
L534:
	;
	v1747 = v1733
	goto L535
L535:
	;
	v1749 = int32(base.Ui32(v1742) >> (uint(int32(31)) % 32))
	if v1742 != 0 {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	v1751 = v1749
	goto L538
L537:
	;
	v1751 = int32(18)
	goto L538
L538:
	;
	if v1751 == int32(0) {
		v2073 = v1747
		goto L230
	} else {
		goto L539
	}
L539:
	;
	if v1751 != int32(18) {
		v2069 = v1747
		v2071 = v1749
		goto L231
	} else {
		goto L540
	}
L540:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1757 = v1741 - v1740
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1756 - v1757
	v1760 = F_r_Suffix_Noun_Step2b(m, l0)
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L4
	} else {
		goto L541
	}
L541:
	;
	if v1760 < int32(0) {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v1764 = v1760
	goto L544
L543:
	;
	v1764 = v1747
	goto L544
L544:
	;
	if v1760 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v1765 = v1764
	goto L547
L546:
	;
	v1765 = v1747
	goto L547
L547:
	;
	v1767 = int32(base.Ui32(v1760) >> (uint(int32(31)) % 32))
	if v1760 != 0 {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	v1769 = v1767
	goto L550
L549:
	;
	v1769 = int32(20)
	goto L550
L550:
	;
	if v1769 == int32(0) {
		v2073 = v1765
		goto L230
	} else {
		goto L551
	}
L551:
	;
	if v1769 != int32(20) {
		v2069 = v1765
		v2071 = v1767
		goto L231
	} else {
		goto L552
	}
L552:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1774 - v1757
	v1777 = F_r_Suffix_Noun_Step2c1(m, l0)
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L4
	} else {
		goto L553
	}
L553:
	;
	if v1777 < int32(0) {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v1781 = v1777
	goto L556
L555:
	;
	v1781 = v1765
	goto L556
L556:
	;
	if v1777 != 0 {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v1782 = v1781
	goto L559
L558:
	;
	v1782 = v1765
	goto L559
L559:
	;
	v1784 = int32(base.Ui32(v1777) >> (uint(int32(31)) % 32))
	if v1777 != 0 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v1786 = v1784
	goto L562
L561:
	;
	v1786 = int32(21)
	goto L562
L562:
	;
	if v1786 == int32(0) {
		v2073 = v1782
		goto L230
	} else {
		goto L563
	}
L563:
	;
	if v1786 != int32(21) {
		v2069 = v1782
		v2071 = v1784
		goto L231
	} else {
		goto L564
	}
L564:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1792 = v1791 - v1757
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1792
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L567
L565:
	;
	if int32(0) <= v1847 {
		v2050 = v1782
		v2054 = v1847
		goto L445
	} else {
		goto L584
	}
L567:
	;
	goto L568
L568:
	;
	goto L569
L569:
	;
	v1802 = v1792
	v1804 = int32(1)
	goto L572
L571:
	;
	v1847 = v1829
	goto L565
L572:
	;
	if v1802 <= v1795 {
		goto L574
	} else {
		goto L575
	}
L573:
	;
	goto L571
L574:
	;
	v1847 = int32(-1)
	goto L565
L575:
	;
	goto L576
L576:
	;
	v1809 = v1802 - int32(1)
	v1811 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1794+v1809))))
	if base.B2i32(int32(0) <= v1811)|base.B2i32(v1809 <= v1795) != 0 {
		v1829 = v1809
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v1833 = int32(1)
	if v1833 < v1804 {
		v1802 = v1829
		v1804 = v1804 - v1833
		goto L572
	} else {
		goto L583
	}
L578:
	;
	v1817 = v1809
	goto L579
L579:
	;
	v1822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1794+v1817))))
	if base.Ui32(int32(191)) < base.Ui32(v1822) {
		v1829 = v1817
		goto L577
	} else {
		goto L581
	}
L580:
	;
	v1829 = v1795
	goto L577
L581:
	;
	v1826 = v1817 - int32(1)
	if v1795 < v1826 {
		v1817 = v1826
		goto L579
	} else {
		goto L582
	}
L582:
	;
	goto L580
L583:
	;
	goto L573
L584:
	;
	v1850 = v1782
	goto L446
L585:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2023
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v2025)))
	if v2026 != 0 {
		goto L643
	} else {
		goto L644
	}
L586:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1862+v1859))))
	if v1864 != int32(134) {
		v2018 = v1850
		goto L585
	} else {
		goto L587
	}
L587:
	;
	v1869 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_57), int32(1))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L4
	} else {
		goto L588
	}
L588:
	;
	if v1869 == int32(0) {
		v2018 = v1850
		goto L585
	} else {
		goto L589
	}
L589:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1873
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1876 = int32(0)
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1875-int32(4))))
	if v1883 == v1876 {
		goto L591
	} else {
		goto L592
	}
L590:
	;
	if v1957 < int32(6) {
		v2018 = v1850
		goto L585
	} else {
		goto L606
	}
L591:
	;
	v1957 = int32(0)
	goto L590
L592:
	;
	goto L593
L593:
	;
	v1888 = v1883 & int32(3)
	if base.Ui32(v1883) < base.Ui32(int32(4)) {
		goto L596
	} else {
		goto L597
	}
L594:
	;
	v1957 = v1946
	goto L590
L595:
	;
	v1930 = v1924
	v1931 = v1925
	v1935 = v1876
	goto L603
L596:
	;
	v1924 = v1875
	v1925 = int32(0)
	goto L595
L597:
	;
	goto L598
L598:
	;
	v1895 = v1875
	v1896 = int32(0)
	v1899 = v1876
	goto L599
L599:
	;
	v1901 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1895))))
	v1902 = int32(-65)
	v1905 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1895)+1)))
	v1909 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1895)+2)))
	v1913 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1895)+3)))
	v1916 = v1896 + base.B2i32(v1902 < v1901) + base.B2i32(v1902 < v1905) + base.B2i32(v1902 < v1909) + base.B2i32(v1902 < v1913)
	v1917 = int32(4)
	v1918 = v1895 + v1917
	v1920 = v1899 + v1917
	if v1920 != v1883&int32(-4) {
		v1895 = v1918
		v1896 = v1916
		v1899 = v1920
		goto L599
	} else {
		goto L601
	}
L600:
	;
	if v1888 == int32(0) {
		v1946 = v1916
		goto L594
	} else {
		goto L602
	}
L601:
	;
	goto L600
L602:
	;
	v1924 = v1918
	v1925 = v1916
	goto L595
L603:
	;
	v1936 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1930))))
	v1939 = v1931 + base.B2i32(int32(-65) < v1936)
	v1940 = int32(1)
	v1943 = v1935 + v1940
	if v1943 != v1888 {
		v1930 = v1930 + v1940
		v1931 = v1939
		v1935 = v1943
		goto L603
	} else {
		goto L605
	}
L604:
	;
	v1946 = v1939
	goto L594
L605:
	;
	goto L604
L606:
	;
	v1960 = F_slice_del(m, l0)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L4
	} else {
		goto L607
	}
L607:
	;
	v1963 = base.B2i32(v1960 < int32(0))
	if v1960 < int32(0) {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	if v1960 < int32(0) {
		goto L611
	} else {
		goto L612
	}
L609:
	;
	goto L610
L610:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1969 = F_r_Suffix_Noun_Step2a(m, l0)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L4
	} else {
		goto L614
	}
L611:
	;
	v1964 = v1960
	goto L613
L612:
	;
	v1964 = v1850
	goto L613
L613:
	;
	v2060 = v1964
	v2062 = int32(base.Ui32(v1960) >> (uint(int32(31)) % 32))
	goto L416
L614:
	;
	if v1969 < int32(0) {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v1973 = v1969
	goto L617
L616:
	;
	v1973 = v1850
	goto L617
L617:
	;
	if v1969 != 0 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v1974 = v1973
	goto L620
L619:
	;
	v1974 = v1850
	goto L620
L620:
	;
	v1976 = int32(base.Ui32(v1969) >> (uint(int32(31)) % 32))
	if v1969 != 0 {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v1978 = v1976
	goto L623
L622:
	;
	v1978 = int32(23)
	goto L623
L623:
	;
	if v1978 == int32(0) {
		v2073 = v1974
		goto L230
	} else {
		goto L624
	}
L624:
	;
	if v1978 != int32(23) {
		v2012 = v1974
		v2014 = v1976
		goto L625
	} else {
		goto L626
	}
L625:
	;
	if v2014 == int32(0) {
		v2073 = v2012
		goto L230
	} else {
		goto L642
	}
L626:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1984 = v1968 - v1967
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1983 - v1984
	v1987 = F_r_Suffix_Noun_Step2b(m, l0)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L4
	} else {
		goto L627
	}
L627:
	;
	if v1987 < int32(0) {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v1991 = v1987
	goto L630
L629:
	;
	v1991 = v1974
	goto L630
L630:
	;
	if v1987 != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v1992 = v1991
	goto L633
L632:
	;
	v1992 = v1974
	goto L633
L633:
	;
	v1994 = int32(base.Ui32(v1987) >> (uint(int32(31)) % 32))
	if v1987 != 0 {
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v1996 = v1994
	goto L636
L635:
	;
	v1996 = int32(25)
	goto L636
L636:
	;
	if v1996 == int32(0) {
		v2073 = v1992
		goto L230
	} else {
		goto L637
	}
L637:
	;
	if v1996 != int32(25) {
		v2012 = v1992
		v2014 = v1994
		goto L625
	} else {
		goto L638
	}
L638:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2001 - v1984
	v2004 = F_r_Suffix_Noun_Step2c1(m, l0)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L4
	} else {
		goto L639
	}
L639:
	;
	if v2004 == int32(0) {
		v2018 = v1992
		goto L585
	} else {
		goto L640
	}
L640:
	;
	if int32(0) <= v2004 {
		v2073 = v1992
		goto L230
	} else {
		goto L641
	}
L641:
	;
	v2012 = v2004
	v2014 = int32(base.Ui32(v2004) >> (uint(int32(31)) % 32))
	goto L625
L642:
	;
	v3966 = v2012
	goto L44
L643:
	;
	v2042 = v2018
	v2045 = v2023
	goto L645
L644:
	;
	v2027 = F_r_Suffix_Noun_Step2a(m, l0)
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L4
	} else {
		goto L646
	}
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2045
	v2047 = F_r_Suffix_Noun_Step2b(m, l0)
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L4
	} else {
		goto L658
	}
L646:
	;
	if v2027 < int32(0) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v2031 = v2027
	goto L649
L648:
	;
	v2031 = v2018
	goto L649
L649:
	;
	if v2027 != 0 {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	v2032 = v2031
	goto L652
L651:
	;
	v2032 = v2018
	goto L652
L652:
	;
	v2034 = int32(base.Ui32(v2027) >> (uint(int32(31)) % 32))
	if v2027 != 0 {
		goto L653
	} else {
		goto L654
	}
L653:
	;
	v2036 = v2034
	goto L655
L654:
	;
	v2036 = int32(27)
	goto L655
L655:
	;
	if v2036 == int32(0) {
		v2073 = v2032
		goto L230
	} else {
		goto L656
	}
L656:
	;
	if v2036 != int32(27) {
		v2060 = v2032
		v2062 = v2034
		goto L416
	} else {
		goto L657
	}
L657:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2042 = v2032
	v2045 = v2041
	goto L645
L658:
	;
	if v2047 != 0 {
		goto L444
	} else {
		goto L659
	}
L659:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2050 = v2042
	v2054 = v2049
	goto L445
L660:
	;
	v2060 = v2047
	v2062 = int32(base.Ui32(v2047) >> (uint(int32(31)) % 32))
	goto L416
L661:
	;
	v2073 = v2060
	goto L230
L662:
	;
	v3966 = v927
	goto L44
L663:
	;
	v3966 = v1316
	goto L44
L664:
	;
	v2073 = v2069
	goto L230
L665:
	;
	v2183 = v2073
	goto L229
L666:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084+v2081))))
	if v2086 != int32(138) {
		goto L665
	} else {
		goto L667
	}
L667:
	;
	v2091 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_58), int32(1))
	mBase = m.M
	v2092 = m.ExcPending
	if v2092 != 0 {
		goto L4
	} else {
		goto L668
	}
L668:
	;
	if v2091 == int32(0) {
		goto L665
	} else {
		goto L669
	}
L669:
	;
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2095
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2098 = int32(0)
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v2097-int32(4))))
	if v2105 == v2098 {
		goto L671
	} else {
		goto L672
	}
L670:
	;
	if int32(3) <= v2179 {
		goto L228
	} else {
		goto L686
	}
L671:
	;
	v2179 = int32(0)
	goto L670
L672:
	;
	goto L673
L673:
	;
	v2110 = v2105 & int32(3)
	if base.Ui32(v2105) < base.Ui32(int32(4)) {
		goto L676
	} else {
		goto L677
	}
L674:
	;
	v2179 = v2168
	goto L670
L675:
	;
	v2152 = v2146
	v2153 = v2147
	v2157 = v2098
	goto L683
L676:
	;
	v2146 = v2097
	v2147 = int32(0)
	goto L675
L677:
	;
	goto L678
L678:
	;
	v2117 = v2097
	v2118 = int32(0)
	v2121 = v2098
	goto L679
L679:
	;
	v2123 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2117))))
	v2124 = int32(-65)
	v2127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2117)+1)))
	v2131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2117)+2)))
	v2135 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2117)+3)))
	v2138 = v2118 + base.B2i32(v2124 < v2123) + base.B2i32(v2124 < v2127) + base.B2i32(v2124 < v2131) + base.B2i32(v2124 < v2135)
	v2139 = int32(4)
	v2140 = v2117 + v2139
	v2142 = v2121 + v2139
	if v2142 != v2105&int32(-4) {
		v2117 = v2140
		v2118 = v2138
		v2121 = v2142
		goto L679
	} else {
		goto L681
	}
L680:
	;
	if v2110 == int32(0) {
		v2168 = v2138
		goto L674
	} else {
		goto L682
	}
L681:
	;
	goto L680
L682:
	;
	v2146 = v2140
	v2147 = v2138
	goto L675
L683:
	;
	v2158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2152))))
	v2161 = v2153 + base.B2i32(int32(-65) < v2158)
	v2162 = int32(1)
	v2165 = v2157 + v2162
	if v2165 != v2110 {
		v2152 = v2152 + v2162
		v2153 = v2161
		v2157 = v2165
		goto L683
	} else {
		goto L685
	}
L684:
	;
	v2168 = v2161
	goto L674
L685:
	;
	goto L684
L686:
	;
	goto L665
L687:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194+v2191))))
	if v2196 != int32(137) {
		v2220 = v2183
		goto L227
	} else {
		goto L688
	}
L688:
	;
	v2201 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_59), int32(1))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L4
	} else {
		goto L689
	}
L689:
	;
	if v2201 == int32(0) {
		v2220 = v2183
		goto L227
	} else {
		goto L690
	}
L690:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2205
	v2209 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_60))
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L4
	} else {
		goto L691
	}
L691:
	;
	if v2209 < int32(0) {
		v3966 = v2209
		goto L44
	} else {
		goto L692
	}
L692:
	;
	v2220 = v2183
	goto L227
L693:
	;
	v2216 = base.B2i32(v2213 < int32(0))
	if v2213 < int32(0) {
		goto L694
	} else {
		goto L695
	}
L694:
	;
	v2217 = v2213
	goto L696
L695:
	;
	v2217 = v2073
	goto L696
L696:
	;
	if v2213 < int32(0) {
		v3966 = v2217
		goto L44
	} else {
		goto L697
	}
L697:
	;
	v2220 = v2217
	goto L227
L698:
	;
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2623
	v2626 = v2623 + int32(1)
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2627 <= v2626 {
		goto L785
	} else {
		goto L786
	}
L699:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2225
	goto L698
L700:
	;
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2232+v2229))))
	if base.B2i32(v2234&int32(224) != int32(160))|base.B2i32(int32(1)<<(uint(v2234)%32)&int32(188) == int32(0)) != 0 {
		goto L699
	} else {
		goto L701
	}
L701:
	;
	v2248 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_61), int32(5))
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L4
	} else {
		goto L702
	}
L702:
	;
	if v2248 == int32(0) {
		goto L699
	} else {
		goto L703
	}
L703:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2252
	switch v2248 - int32(1) {
	case 0:
		goto L707
	case 1:
		goto L706
	case 2:
		goto L705
	case 3:
		goto L704
	default:
		goto L698
	}
L704:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2530 = int32(0)
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(v2529-int32(4))))
	if v2537 == v2530 {
		goto L766
	} else {
		goto L767
	}
L705:
	;
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2439 = int32(0)
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v2438-int32(4))))
	if v2446 == v2439 {
		goto L747
	} else {
		goto L748
	}
L706:
	;
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2348 = int32(0)
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2347-int32(4))))
	if v2355 == v2348 {
		goto L728
	} else {
		goto L729
	}
L707:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2257 = int32(0)
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v2256-int32(4))))
	if v2264 == v2257 {
		goto L709
	} else {
		goto L710
	}
L708:
	;
	if v2338 < int32(4) {
		goto L699
	} else {
		goto L724
	}
L709:
	;
	v2338 = int32(0)
	goto L708
L710:
	;
	goto L711
L711:
	;
	v2269 = v2264 & int32(3)
	if base.Ui32(v2264) < base.Ui32(int32(4)) {
		goto L714
	} else {
		goto L715
	}
L712:
	;
	v2338 = v2327
	goto L708
L713:
	;
	v2311 = v2305
	v2312 = v2306
	v2316 = v2257
	goto L721
L714:
	;
	v2305 = v2256
	v2306 = int32(0)
	goto L713
L715:
	;
	goto L716
L716:
	;
	v2276 = v2256
	v2277 = int32(0)
	v2280 = v2257
	goto L717
L717:
	;
	v2282 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2276))))
	v2283 = int32(-65)
	v2286 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2276)+1)))
	v2290 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2276)+2)))
	v2294 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2276)+3)))
	v2297 = v2277 + base.B2i32(v2283 < v2282) + base.B2i32(v2283 < v2286) + base.B2i32(v2283 < v2290) + base.B2i32(v2283 < v2294)
	v2298 = int32(4)
	v2299 = v2276 + v2298
	v2301 = v2280 + v2298
	if v2301 != v2264&int32(-4) {
		v2276 = v2299
		v2277 = v2297
		v2280 = v2301
		goto L717
	} else {
		goto L719
	}
L718:
	;
	if v2269 == int32(0) {
		v2327 = v2297
		goto L712
	} else {
		goto L720
	}
L719:
	;
	goto L718
L720:
	;
	v2305 = v2299
	v2306 = v2297
	goto L713
L721:
	;
	v2317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2311))))
	v2320 = v2312 + base.B2i32(int32(-65) < v2317)
	v2321 = int32(1)
	v2324 = v2316 + v2321
	if v2324 != v2269 {
		v2311 = v2311 + v2321
		v2312 = v2320
		v2316 = v2324
		goto L721
	} else {
		goto L723
	}
L722:
	;
	v2327 = v2320
	goto L712
L723:
	;
	goto L722
L724:
	;
	v2343 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_62))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L4
	} else {
		goto L725
	}
L725:
	;
	if int32(0) <= v2343 {
		goto L698
	} else {
		goto L726
	}
L726:
	;
	v3966 = v2343
	goto L44
L727:
	;
	if v2429 < int32(4) {
		goto L699
	} else {
		goto L743
	}
L728:
	;
	v2429 = int32(0)
	goto L727
L729:
	;
	goto L730
L730:
	;
	v2360 = v2355 & int32(3)
	if base.Ui32(v2355) < base.Ui32(int32(4)) {
		goto L733
	} else {
		goto L734
	}
L731:
	;
	v2429 = v2418
	goto L727
L732:
	;
	v2402 = v2396
	v2403 = v2397
	v2407 = v2348
	goto L740
L733:
	;
	v2396 = v2347
	v2397 = int32(0)
	goto L732
L734:
	;
	goto L735
L735:
	;
	v2367 = v2347
	v2368 = int32(0)
	v2371 = v2348
	goto L736
L736:
	;
	v2373 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2367))))
	v2374 = int32(-65)
	v2377 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2367)+1)))
	v2381 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2367)+2)))
	v2385 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2367)+3)))
	v2388 = v2368 + base.B2i32(v2374 < v2373) + base.B2i32(v2374 < v2377) + base.B2i32(v2374 < v2381) + base.B2i32(v2374 < v2385)
	v2389 = int32(4)
	v2390 = v2367 + v2389
	v2392 = v2371 + v2389
	if v2392 != v2355&int32(-4) {
		v2367 = v2390
		v2368 = v2388
		v2371 = v2392
		goto L736
	} else {
		goto L738
	}
L737:
	;
	if v2360 == int32(0) {
		v2418 = v2388
		goto L731
	} else {
		goto L739
	}
L738:
	;
	goto L737
L739:
	;
	v2396 = v2390
	v2397 = v2388
	goto L732
L740:
	;
	v2408 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2402))))
	v2411 = v2403 + base.B2i32(int32(-65) < v2408)
	v2412 = int32(1)
	v2415 = v2407 + v2412
	if v2415 != v2360 {
		v2402 = v2402 + v2412
		v2403 = v2411
		v2407 = v2415
		goto L740
	} else {
		goto L742
	}
L741:
	;
	v2418 = v2411
	goto L731
L742:
	;
	goto L741
L743:
	;
	v2434 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_63))
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L4
	} else {
		goto L744
	}
L744:
	;
	if int32(0) <= v2434 {
		goto L698
	} else {
		goto L745
	}
L745:
	;
	v3966 = v2434
	goto L44
L746:
	;
	if v2520 < int32(4) {
		goto L699
	} else {
		goto L762
	}
L747:
	;
	v2520 = int32(0)
	goto L746
L748:
	;
	goto L749
L749:
	;
	v2451 = v2446 & int32(3)
	if base.Ui32(v2446) < base.Ui32(int32(4)) {
		goto L752
	} else {
		goto L753
	}
L750:
	;
	v2520 = v2509
	goto L746
L751:
	;
	v2493 = v2487
	v2494 = v2488
	v2498 = v2439
	goto L759
L752:
	;
	v2487 = v2438
	v2488 = int32(0)
	goto L751
L753:
	;
	goto L754
L754:
	;
	v2458 = v2438
	v2459 = int32(0)
	v2462 = v2439
	goto L755
L755:
	;
	v2464 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2458))))
	v2465 = int32(-65)
	v2468 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2458)+1)))
	v2472 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2458)+2)))
	v2476 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2458)+3)))
	v2479 = v2459 + base.B2i32(v2465 < v2464) + base.B2i32(v2465 < v2468) + base.B2i32(v2465 < v2472) + base.B2i32(v2465 < v2476)
	v2480 = int32(4)
	v2481 = v2458 + v2480
	v2483 = v2462 + v2480
	if v2483 != v2446&int32(-4) {
		v2458 = v2481
		v2459 = v2479
		v2462 = v2483
		goto L755
	} else {
		goto L757
	}
L756:
	;
	if v2451 == int32(0) {
		v2509 = v2479
		goto L750
	} else {
		goto L758
	}
L757:
	;
	goto L756
L758:
	;
	v2487 = v2481
	v2488 = v2479
	goto L751
L759:
	;
	v2499 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2493))))
	v2502 = v2494 + base.B2i32(int32(-65) < v2499)
	v2503 = int32(1)
	v2506 = v2498 + v2503
	if v2506 != v2451 {
		v2493 = v2493 + v2503
		v2494 = v2502
		v2498 = v2506
		goto L759
	} else {
		goto L761
	}
L760:
	;
	v2509 = v2502
	goto L750
L761:
	;
	goto L760
L762:
	;
	v2525 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_64))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L4
	} else {
		goto L763
	}
L763:
	;
	if int32(0) <= v2525 {
		goto L698
	} else {
		goto L764
	}
L764:
	;
	v3966 = v2525
	goto L44
L765:
	;
	if v2611 < int32(4) {
		goto L699
	} else {
		goto L781
	}
L766:
	;
	v2611 = int32(0)
	goto L765
L767:
	;
	goto L768
L768:
	;
	v2542 = v2537 & int32(3)
	if base.Ui32(v2537) < base.Ui32(int32(4)) {
		goto L771
	} else {
		goto L772
	}
L769:
	;
	v2611 = v2600
	goto L765
L770:
	;
	v2584 = v2578
	v2585 = v2579
	v2589 = v2530
	goto L778
L771:
	;
	v2578 = v2529
	v2579 = int32(0)
	goto L770
L772:
	;
	goto L773
L773:
	;
	v2549 = v2529
	v2550 = int32(0)
	v2553 = v2530
	goto L774
L774:
	;
	v2555 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2549))))
	v2556 = int32(-65)
	v2559 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2549)+1)))
	v2563 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2549)+2)))
	v2567 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2549)+3)))
	v2570 = v2550 + base.B2i32(v2556 < v2555) + base.B2i32(v2556 < v2559) + base.B2i32(v2556 < v2563) + base.B2i32(v2556 < v2567)
	v2571 = int32(4)
	v2572 = v2549 + v2571
	v2574 = v2553 + v2571
	if v2574 != v2537&int32(-4) {
		v2549 = v2572
		v2550 = v2570
		v2553 = v2574
		goto L774
	} else {
		goto L776
	}
L775:
	;
	if v2542 == int32(0) {
		v2600 = v2570
		goto L769
	} else {
		goto L777
	}
L776:
	;
	goto L775
L777:
	;
	v2578 = v2572
	v2579 = v2570
	goto L770
L778:
	;
	v2590 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2584))))
	v2593 = v2585 + base.B2i32(int32(-65) < v2590)
	v2594 = int32(1)
	v2597 = v2589 + v2594
	if v2597 != v2542 {
		v2584 = v2584 + v2594
		v2585 = v2593
		v2589 = v2597
		goto L778
	} else {
		goto L780
	}
L779:
	;
	v2600 = v2593
	goto L769
L780:
	;
	goto L779
L781:
	;
	v2616 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_65))
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		goto L4
	} else {
		goto L782
	}
L782:
	;
	if int32(0) <= v2616 {
		goto L698
	} else {
		goto L783
	}
L783:
	;
	v3966 = v2616
	goto L44
L784:
	;
	v2753 = int32(0)
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2754
	v2757 = v2754 + int32(3)
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2758 <= v2757 {
		v2959 = v2753
		goto L814
	} else {
		goto L815
	}
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2623
	v2752 = v2623
	goto L784
L786:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2629+v2626))))
	switch v2631 - int32(129) {
	case 0, 7:
		goto L787
	default:
		goto L785
	}
L787:
	;
	v2636 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_66), int32(2))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		goto L4
	} else {
		goto L788
	}
L788:
	;
	if v2636 == int32(0) {
		goto L785
	} else {
		goto L789
	}
L789:
	;
	v2640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2640
	v2642 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2643 = int32(0)
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v2642-int32(4))))
	if v2650 == v2643 {
		goto L791
	} else {
		goto L792
	}
L790:
	;
	if v2724 < int32(4) {
		goto L785
	} else {
		goto L806
	}
L791:
	;
	v2724 = int32(0)
	goto L790
L792:
	;
	goto L793
L793:
	;
	v2655 = v2650 & int32(3)
	if base.Ui32(v2650) < base.Ui32(int32(4)) {
		goto L796
	} else {
		goto L797
	}
L794:
	;
	v2724 = v2713
	goto L790
L795:
	;
	v2697 = v2691
	v2698 = v2692
	v2702 = v2643
	goto L803
L796:
	;
	v2691 = v2642
	v2692 = int32(0)
	goto L795
L797:
	;
	goto L798
L798:
	;
	v2662 = v2642
	v2663 = int32(0)
	v2666 = v2643
	goto L799
L799:
	;
	v2668 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2662))))
	v2669 = int32(-65)
	v2672 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2662)+1)))
	v2676 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2662)+2)))
	v2680 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2662)+3)))
	v2683 = v2663 + base.B2i32(v2669 < v2668) + base.B2i32(v2669 < v2672) + base.B2i32(v2669 < v2676) + base.B2i32(v2669 < v2680)
	v2684 = int32(4)
	v2685 = v2662 + v2684
	v2687 = v2666 + v2684
	if v2687 != v2650&int32(-4) {
		v2662 = v2685
		v2663 = v2683
		v2666 = v2687
		goto L799
	} else {
		goto L801
	}
L800:
	;
	if v2655 == int32(0) {
		v2713 = v2683
		goto L794
	} else {
		goto L802
	}
L801:
	;
	goto L800
L802:
	;
	v2691 = v2685
	v2692 = v2683
	goto L795
L803:
	;
	v2703 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2697))))
	v2706 = v2698 + base.B2i32(int32(-65) < v2703)
	v2707 = int32(1)
	v2710 = v2702 + v2707
	if v2710 != v2655 {
		v2697 = v2697 + v2707
		v2698 = v2706
		v2702 = v2710
		goto L803
	} else {
		goto L805
	}
L804:
	;
	v2713 = v2706
	goto L794
L805:
	;
	goto L804
L806:
	;
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2728 = int32(2)
	v2730 = int32(0)
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2732-v2727 < v2728 {
		v2742 = v2730
		goto L808
	} else {
		goto L809
	}
L807:
	;
	if v2742 != 0 {
		goto L785
	} else {
		goto L811
	}
L808:
	;
	goto L807
L809:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2738 = F_memcmp(m, v2736+v2727, int32(_a_F_arabic_UTF_8_stem_67), v2728)
	mBase = m.M
	if v2738 != 0 {
		v2742 = v2730
		goto L808
	} else {
		goto L810
	}
L810:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2728 + v2727
	v2742 = int32(1)
	goto L808
L811:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2727
	v2744 = F_slice_del(m, l0)
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		goto L4
	} else {
		goto L812
	}
L812:
	;
	if v2744 < int32(0) {
		v3966 = v2744
		goto L44
	} else {
		goto L813
	}
L813:
	;
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2752 = v2748
	goto L784
L814:
	;
	v2961 = int32(base.Ui32(v2959) >> (uint(int32(31)) % 32))
	if v2959 != 0 {
		goto L861
	} else {
		goto L862
	}
L815:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2760+v2757))))
	if base.B2i32(v2762 != int32(167))&base.B2i32(v2762 != int32(132)) != 0 {
		v2959 = v2753
		goto L814
	} else {
		goto L816
	}
L816:
	;
	v2770 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_68), int32(4))
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		goto L4
	} else {
		goto L817
	}
L817:
	;
	if v2770 == int32(0) {
		v2959 = v2753
		goto L814
	} else {
		goto L818
	}
L818:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2774
	switch v2770 - int32(1) {
	case 0:
		goto L821
	case 1:
		goto L820
	default:
		goto L819
	}
L819:
	;
	v2959 = int32(1)
	goto L814
L820:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2868 = int32(0)
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v2867-int32(4))))
	if v2875 == v2868 {
		goto L842
	} else {
		goto L843
	}
L821:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2779 = int32(0)
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2778-int32(4))))
	if v2786 == v2779 {
		goto L823
	} else {
		goto L824
	}
L822:
	;
	if v2860 < int32(6) {
		v2959 = v2753
		goto L814
	} else {
		goto L838
	}
L823:
	;
	v2860 = int32(0)
	goto L822
L824:
	;
	goto L825
L825:
	;
	v2791 = v2786 & int32(3)
	if base.Ui32(v2786) < base.Ui32(int32(4)) {
		goto L828
	} else {
		goto L829
	}
L826:
	;
	v2860 = v2849
	goto L822
L827:
	;
	v2833 = v2827
	v2834 = v2828
	v2838 = v2779
	goto L835
L828:
	;
	v2827 = v2778
	v2828 = int32(0)
	goto L827
L829:
	;
	goto L830
L830:
	;
	v2798 = v2778
	v2799 = int32(0)
	v2802 = v2779
	goto L831
L831:
	;
	v2804 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2798))))
	v2805 = int32(-65)
	v2808 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2798)+1)))
	v2812 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2798)+2)))
	v2816 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2798)+3)))
	v2819 = v2799 + base.B2i32(v2805 < v2804) + base.B2i32(v2805 < v2808) + base.B2i32(v2805 < v2812) + base.B2i32(v2805 < v2816)
	v2820 = int32(4)
	v2821 = v2798 + v2820
	v2823 = v2802 + v2820
	if v2823 != v2786&int32(-4) {
		v2798 = v2821
		v2799 = v2819
		v2802 = v2823
		goto L831
	} else {
		goto L833
	}
L832:
	;
	if v2791 == int32(0) {
		v2849 = v2819
		goto L826
	} else {
		goto L834
	}
L833:
	;
	goto L832
L834:
	;
	v2827 = v2821
	v2828 = v2819
	goto L827
L835:
	;
	v2839 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2833))))
	v2842 = v2834 + base.B2i32(int32(-65) < v2839)
	v2843 = int32(1)
	v2846 = v2838 + v2843
	if v2846 != v2791 {
		v2833 = v2833 + v2843
		v2834 = v2842
		v2838 = v2846
		goto L835
	} else {
		goto L837
	}
L836:
	;
	v2849 = v2842
	goto L826
L837:
	;
	goto L836
L838:
	;
	v2863 = F_slice_del(m, l0)
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L4
	} else {
		goto L839
	}
L839:
	;
	if int32(0) <= v2863 {
		goto L819
	} else {
		goto L840
	}
L840:
	;
	v2959 = v2863
	goto L814
L841:
	;
	if v2949 < int32(5) {
		v2959 = v2753
		goto L814
	} else {
		goto L857
	}
L842:
	;
	v2949 = int32(0)
	goto L841
L843:
	;
	goto L844
L844:
	;
	v2880 = v2875 & int32(3)
	if base.Ui32(v2875) < base.Ui32(int32(4)) {
		goto L847
	} else {
		goto L848
	}
L845:
	;
	v2949 = v2938
	goto L841
L846:
	;
	v2922 = v2916
	v2923 = v2917
	v2927 = v2868
	goto L854
L847:
	;
	v2916 = v2867
	v2917 = int32(0)
	goto L846
L848:
	;
	goto L849
L849:
	;
	v2887 = v2867
	v2888 = int32(0)
	v2891 = v2868
	goto L850
L850:
	;
	v2893 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2887))))
	v2894 = int32(-65)
	v2897 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2887)+1)))
	v2901 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2887)+2)))
	v2905 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2887)+3)))
	v2908 = v2888 + base.B2i32(v2894 < v2893) + base.B2i32(v2894 < v2897) + base.B2i32(v2894 < v2901) + base.B2i32(v2894 < v2905)
	v2909 = int32(4)
	v2910 = v2887 + v2909
	v2912 = v2891 + v2909
	if v2912 != v2875&int32(-4) {
		v2887 = v2910
		v2888 = v2908
		v2891 = v2912
		goto L850
	} else {
		goto L852
	}
L851:
	;
	if v2880 == int32(0) {
		v2938 = v2908
		goto L845
	} else {
		goto L853
	}
L852:
	;
	goto L851
L853:
	;
	v2916 = v2910
	v2917 = v2908
	goto L846
L854:
	;
	v2928 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2922))))
	v2931 = v2923 + base.B2i32(int32(-65) < v2928)
	v2932 = int32(1)
	v2935 = v2927 + v2932
	if v2935 != v2880 {
		v2922 = v2922 + v2932
		v2923 = v2931
		v2927 = v2935
		goto L854
	} else {
		goto L856
	}
L855:
	;
	v2938 = v2931
	goto L845
L856:
	;
	goto L855
L857:
	;
	v2952 = F_slice_del(m, l0)
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L4
	} else {
		goto L858
	}
L858:
	;
	if v2952 < int32(0) {
		v2959 = v2952
		goto L814
	} else {
		goto L859
	}
L859:
	;
	goto L819
L860:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2225
	v3811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3811
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3811
	v3815 = v3811 - int32(1)
	if v3815 <= v2225 {
		goto L1073
	} else {
		goto L1074
	}
L861:
	;
	v2963 = v2961
	goto L863
L862:
	;
	v2963 = int32(32)
	goto L863
L863:
	;
	if v2963 == int32(0) {
		goto L860
	} else {
		goto L864
	}
L864:
	;
	if v2959 < int32(0) {
		goto L865
	} else {
		goto L866
	}
L865:
	;
	v2968 = v2959
	goto L867
L866:
	;
	v2968 = v2220
	goto L867
L867:
	;
	if v2959 != 0 {
		goto L868
	} else {
		goto L869
	}
L868:
	;
	v2969 = v2968
	goto L870
L869:
	;
	v2969 = v2220
	goto L870
L870:
	;
	if v2963 != int32(32) {
		v3802 = v2969
		v3804 = v2961
		goto L871
	} else {
		goto L872
	}
L871:
	;
	if v3804 != 0 {
		v3966 = v3802
		goto L44
	} else {
		goto L1072
	}
L872:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2752
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(v2973)+8))
	if v2974 != 0 {
		goto L873
	} else {
		goto L874
	}
L873:
	;
	v2975 = int32(0)
	v2976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2976
	v2979 = v2976 + int32(1)
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2980 <= v2979 {
		v3274 = v2975
		goto L876
	} else {
		goto L877
	}
L874:
	;
	v3289 = v2973
	goto L875
L875:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2752
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3289)+4))
	if v3292 == int32(0) {
		goto L860
	} else {
		goto L955
	}
L876:
	;
	v3276 = int32(base.Ui32(v3274) >> (uint(int32(31)) % 32))
	if v3274 != 0 {
		goto L942
	} else {
		goto L943
	}
L877:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2982+v2979))))
	if base.B2i32(v2984 != int32(168))&base.B2i32(v2984 != int32(131)) != 0 {
		v3274 = v2975
		goto L876
	} else {
		goto L878
	}
L878:
	;
	v2992 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_69), int32(4))
	mBase = m.M
	v2993 = m.ExcPending
	if v2993 != 0 {
		goto L4
	} else {
		goto L879
	}
L879:
	;
	if v2992 == int32(0) {
		v3274 = v2975
		goto L876
	} else {
		goto L880
	}
L880:
	;
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2996
	switch v2992 - int32(1) {
	case 0:
		goto L884
	case 1:
		goto L883
	case 2:
		goto L882
	default:
		goto L881
	}
L881:
	;
	v3274 = int32(1)
	goto L876
L882:
	;
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3181 = int32(0)
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v3180-int32(4))))
	if v3188 == v3181 {
		goto L924
	} else {
		goto L925
	}
L883:
	;
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3090 = int32(0)
	v3097 = *(*int32)(unsafe.Add(mBase, uint32(v3089-int32(4))))
	if v3097 == v3090 {
		goto L905
	} else {
		goto L906
	}
L884:
	;
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3001 = int32(0)
	v3008 = *(*int32)(unsafe.Add(mBase, uint32(v3000-int32(4))))
	if v3008 == v3001 {
		goto L886
	} else {
		goto L887
	}
L885:
	;
	if v3082 < int32(4) {
		v3274 = v2975
		goto L876
	} else {
		goto L901
	}
L886:
	;
	v3082 = int32(0)
	goto L885
L887:
	;
	goto L888
L888:
	;
	v3013 = v3008 & int32(3)
	if base.Ui32(v3008) < base.Ui32(int32(4)) {
		goto L891
	} else {
		goto L892
	}
L889:
	;
	v3082 = v3071
	goto L885
L890:
	;
	v3055 = v3049
	v3056 = v3050
	v3060 = v3001
	goto L898
L891:
	;
	v3049 = v3000
	v3050 = int32(0)
	goto L890
L892:
	;
	goto L893
L893:
	;
	v3020 = v3000
	v3021 = int32(0)
	v3024 = v3001
	goto L894
L894:
	;
	v3026 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3020))))
	v3027 = int32(-65)
	v3030 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3020)+1)))
	v3034 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3020)+2)))
	v3038 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3020)+3)))
	v3041 = v3021 + base.B2i32(v3027 < v3026) + base.B2i32(v3027 < v3030) + base.B2i32(v3027 < v3034) + base.B2i32(v3027 < v3038)
	v3042 = int32(4)
	v3043 = v3020 + v3042
	v3045 = v3024 + v3042
	if v3045 != v3008&int32(-4) {
		v3020 = v3043
		v3021 = v3041
		v3024 = v3045
		goto L894
	} else {
		goto L896
	}
L895:
	;
	if v3013 == int32(0) {
		v3071 = v3041
		goto L889
	} else {
		goto L897
	}
L896:
	;
	goto L895
L897:
	;
	v3049 = v3043
	v3050 = v3041
	goto L890
L898:
	;
	v3061 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3055))))
	v3064 = v3056 + base.B2i32(int32(-65) < v3061)
	v3065 = int32(1)
	v3068 = v3060 + v3065
	if v3068 != v3013 {
		v3055 = v3055 + v3065
		v3056 = v3064
		v3060 = v3068
		goto L898
	} else {
		goto L900
	}
L899:
	;
	v3071 = v3064
	goto L889
L900:
	;
	goto L899
L901:
	;
	v3085 = F_slice_del(m, l0)
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L4
	} else {
		goto L902
	}
L902:
	;
	if int32(0) <= v3085 {
		goto L881
	} else {
		goto L903
	}
L903:
	;
	v3274 = v3085
	goto L876
L904:
	;
	if v3171 < int32(4) {
		v3274 = v2975
		goto L876
	} else {
		goto L920
	}
L905:
	;
	v3171 = int32(0)
	goto L904
L906:
	;
	goto L907
L907:
	;
	v3102 = v3097 & int32(3)
	if base.Ui32(v3097) < base.Ui32(int32(4)) {
		goto L910
	} else {
		goto L911
	}
L908:
	;
	v3171 = v3160
	goto L904
L909:
	;
	v3144 = v3138
	v3145 = v3139
	v3149 = v3090
	goto L917
L910:
	;
	v3138 = v3089
	v3139 = int32(0)
	goto L909
L911:
	;
	goto L912
L912:
	;
	v3109 = v3089
	v3110 = int32(0)
	v3113 = v3090
	goto L913
L913:
	;
	v3115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3109))))
	v3116 = int32(-65)
	v3119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3109)+1)))
	v3123 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3109)+2)))
	v3127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3109)+3)))
	v3130 = v3110 + base.B2i32(v3116 < v3115) + base.B2i32(v3116 < v3119) + base.B2i32(v3116 < v3123) + base.B2i32(v3116 < v3127)
	v3131 = int32(4)
	v3132 = v3109 + v3131
	v3134 = v3113 + v3131
	if v3134 != v3097&int32(-4) {
		v3109 = v3132
		v3110 = v3130
		v3113 = v3134
		goto L913
	} else {
		goto L915
	}
L914:
	;
	if v3102 == int32(0) {
		v3160 = v3130
		goto L908
	} else {
		goto L916
	}
L915:
	;
	goto L914
L916:
	;
	v3138 = v3132
	v3139 = v3130
	goto L909
L917:
	;
	v3150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3144))))
	v3153 = v3145 + base.B2i32(int32(-65) < v3150)
	v3154 = int32(1)
	v3157 = v3149 + v3154
	if v3157 != v3102 {
		v3144 = v3144 + v3154
		v3145 = v3153
		v3149 = v3157
		goto L917
	} else {
		goto L919
	}
L918:
	;
	v3160 = v3153
	goto L908
L919:
	;
	goto L918
L920:
	;
	v3176 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_70))
	mBase = m.M
	v3177 = m.ExcPending
	if v3177 != 0 {
		goto L4
	} else {
		goto L921
	}
L921:
	;
	if int32(0) <= v3176 {
		goto L881
	} else {
		goto L922
	}
L922:
	;
	v3274 = v3176
	goto L876
L923:
	;
	if v3262 < int32(4) {
		v3274 = v2975
		goto L876
	} else {
		goto L939
	}
L924:
	;
	v3262 = int32(0)
	goto L923
L925:
	;
	goto L926
L926:
	;
	v3193 = v3188 & int32(3)
	if base.Ui32(v3188) < base.Ui32(int32(4)) {
		goto L929
	} else {
		goto L930
	}
L927:
	;
	v3262 = v3251
	goto L923
L928:
	;
	v3235 = v3229
	v3236 = v3230
	v3240 = v3181
	goto L936
L929:
	;
	v3229 = v3180
	v3230 = int32(0)
	goto L928
L930:
	;
	goto L931
L931:
	;
	v3200 = v3180
	v3201 = int32(0)
	v3204 = v3181
	goto L932
L932:
	;
	v3206 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3200))))
	v3207 = int32(-65)
	v3210 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3200)+1)))
	v3214 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3200)+2)))
	v3218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3200)+3)))
	v3221 = v3201 + base.B2i32(v3207 < v3206) + base.B2i32(v3207 < v3210) + base.B2i32(v3207 < v3214) + base.B2i32(v3207 < v3218)
	v3222 = int32(4)
	v3223 = v3200 + v3222
	v3225 = v3204 + v3222
	if v3225 != v3188&int32(-4) {
		v3200 = v3223
		v3201 = v3221
		v3204 = v3225
		goto L932
	} else {
		goto L934
	}
L933:
	;
	if v3193 == int32(0) {
		v3251 = v3221
		goto L927
	} else {
		goto L935
	}
L934:
	;
	goto L933
L935:
	;
	v3229 = v3223
	v3230 = v3221
	goto L928
L936:
	;
	v3241 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3235))))
	v3244 = v3236 + base.B2i32(int32(-65) < v3241)
	v3245 = int32(1)
	v3248 = v3240 + v3245
	if v3248 != v3193 {
		v3235 = v3235 + v3245
		v3236 = v3244
		v3240 = v3248
		goto L936
	} else {
		goto L938
	}
L937:
	;
	v3251 = v3244
	goto L927
L938:
	;
	goto L937
L939:
	;
	v3267 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_71))
	mBase = m.M
	v3268 = m.ExcPending
	if v3268 != 0 {
		goto L4
	} else {
		goto L940
	}
L940:
	;
	if v3267 < int32(0) {
		v3274 = v3267
		goto L876
	} else {
		goto L941
	}
L941:
	;
	goto L881
L942:
	;
	v3278 = v3276
	goto L944
L943:
	;
	v3278 = int32(34)
	goto L944
L944:
	;
	if v3278 == int32(0) {
		goto L860
	} else {
		goto L945
	}
L945:
	;
	if v3278 != int32(34) {
		goto L946
	} else {
		goto L947
	}
L946:
	;
	if v3274 < int32(0) {
		goto L949
	} else {
		goto L950
	}
L947:
	;
	goto L948
L948:
	;
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3289 = v3287
	goto L875
L949:
	;
	v3285 = v3274
	goto L951
L950:
	;
	v3285 = v2969
	goto L951
L951:
	;
	if v3274 != 0 {
		goto L952
	} else {
		goto L953
	}
L952:
	;
	v3286 = v3285
	goto L954
L953:
	;
	v3286 = v2969
	goto L954
L954:
	;
	v3802 = v3286
	v3804 = v3276
	goto L871
L955:
	;
	v3295 = int32(0)
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3296
	v3300 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_72), int32(4))
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L4
	} else {
		goto L957
	}
L956:
	;
	if v3674 == int32(0) {
		goto L1041
	} else {
		goto L1042
	}
L957:
	;
	if v3300 == int32(0) {
		v3674 = v3295
		goto L956
	} else {
		goto L958
	}
L958:
	;
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3304
	switch v3300 - int32(1) {
	case 0:
		goto L963
	case 1:
		goto L962
	case 2:
		goto L961
	case 3:
		goto L960
	default:
		goto L959
	}
L959:
	;
	v3674 = int32(1)
	goto L956
L960:
	;
	v3581 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3582 = int32(0)
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(v3581-int32(4))))
	if v3589 == v3582 {
		goto L1022
	} else {
		goto L1023
	}
L961:
	;
	v3490 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3491 = int32(0)
	v3498 = *(*int32)(unsafe.Add(mBase, uint32(v3490-int32(4))))
	if v3498 == v3491 {
		goto L1003
	} else {
		goto L1004
	}
L962:
	;
	v3399 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3400 = int32(0)
	v3407 = *(*int32)(unsafe.Add(mBase, uint32(v3399-int32(4))))
	if v3407 == v3400 {
		goto L984
	} else {
		goto L985
	}
L963:
	;
	v3308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3309 = int32(0)
	v3316 = *(*int32)(unsafe.Add(mBase, uint32(v3308-int32(4))))
	if v3316 == v3309 {
		goto L965
	} else {
		goto L966
	}
L964:
	;
	if v3390 < int32(5) {
		v3674 = v3295
		goto L956
	} else {
		goto L980
	}
L965:
	;
	v3390 = int32(0)
	goto L964
L966:
	;
	goto L967
L967:
	;
	v3321 = v3316 & int32(3)
	if base.Ui32(v3316) < base.Ui32(int32(4)) {
		goto L970
	} else {
		goto L971
	}
L968:
	;
	v3390 = v3379
	goto L964
L969:
	;
	v3363 = v3357
	v3364 = v3358
	v3368 = v3309
	goto L977
L970:
	;
	v3357 = v3308
	v3358 = int32(0)
	goto L969
L971:
	;
	goto L972
L972:
	;
	v3328 = v3308
	v3329 = int32(0)
	v3332 = v3309
	goto L973
L973:
	;
	v3334 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3328))))
	v3335 = int32(-65)
	v3338 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3328)+1)))
	v3342 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3328)+2)))
	v3346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3328)+3)))
	v3349 = v3329 + base.B2i32(v3335 < v3334) + base.B2i32(v3335 < v3338) + base.B2i32(v3335 < v3342) + base.B2i32(v3335 < v3346)
	v3350 = int32(4)
	v3351 = v3328 + v3350
	v3353 = v3332 + v3350
	if v3353 != v3316&int32(-4) {
		v3328 = v3351
		v3329 = v3349
		v3332 = v3353
		goto L973
	} else {
		goto L975
	}
L974:
	;
	if v3321 == int32(0) {
		v3379 = v3349
		goto L968
	} else {
		goto L976
	}
L975:
	;
	goto L974
L976:
	;
	v3357 = v3351
	v3358 = v3349
	goto L969
L977:
	;
	v3369 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3363))))
	v3372 = v3364 + base.B2i32(int32(-65) < v3369)
	v3373 = int32(1)
	v3376 = v3368 + v3373
	if v3376 != v3321 {
		v3363 = v3363 + v3373
		v3364 = v3372
		v3368 = v3376
		goto L977
	} else {
		goto L979
	}
L978:
	;
	v3379 = v3372
	goto L968
L979:
	;
	goto L978
L980:
	;
	v3395 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_73))
	mBase = m.M
	v3396 = m.ExcPending
	if v3396 != 0 {
		goto L4
	} else {
		goto L981
	}
L981:
	;
	if int32(0) <= v3395 {
		goto L959
	} else {
		goto L982
	}
L982:
	;
	v3674 = v3395
	goto L956
L983:
	;
	if v3481 < int32(5) {
		v3674 = v3295
		goto L956
	} else {
		goto L999
	}
L984:
	;
	v3481 = int32(0)
	goto L983
L985:
	;
	goto L986
L986:
	;
	v3412 = v3407 & int32(3)
	if base.Ui32(v3407) < base.Ui32(int32(4)) {
		goto L989
	} else {
		goto L990
	}
L987:
	;
	v3481 = v3470
	goto L983
L988:
	;
	v3454 = v3448
	v3455 = v3449
	v3459 = v3400
	goto L996
L989:
	;
	v3448 = v3399
	v3449 = int32(0)
	goto L988
L990:
	;
	goto L991
L991:
	;
	v3419 = v3399
	v3420 = int32(0)
	v3423 = v3400
	goto L992
L992:
	;
	v3425 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3419))))
	v3426 = int32(-65)
	v3429 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3419)+1)))
	v3433 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3419)+2)))
	v3437 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3419)+3)))
	v3440 = v3420 + base.B2i32(v3426 < v3425) + base.B2i32(v3426 < v3429) + base.B2i32(v3426 < v3433) + base.B2i32(v3426 < v3437)
	v3441 = int32(4)
	v3442 = v3419 + v3441
	v3444 = v3423 + v3441
	if v3444 != v3407&int32(-4) {
		v3419 = v3442
		v3420 = v3440
		v3423 = v3444
		goto L992
	} else {
		goto L994
	}
L993:
	;
	if v3412 == int32(0) {
		v3470 = v3440
		goto L987
	} else {
		goto L995
	}
L994:
	;
	goto L993
L995:
	;
	v3448 = v3442
	v3449 = v3440
	goto L988
L996:
	;
	v3460 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3454))))
	v3463 = v3455 + base.B2i32(int32(-65) < v3460)
	v3464 = int32(1)
	v3467 = v3459 + v3464
	if v3467 != v3412 {
		v3454 = v3454 + v3464
		v3455 = v3463
		v3459 = v3467
		goto L996
	} else {
		goto L998
	}
L997:
	;
	v3470 = v3463
	goto L987
L998:
	;
	goto L997
L999:
	;
	v3486 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_74))
	mBase = m.M
	v3487 = m.ExcPending
	if v3487 != 0 {
		goto L4
	} else {
		goto L1000
	}
L1000:
	;
	if int32(0) <= v3486 {
		goto L959
	} else {
		goto L1001
	}
L1001:
	;
	v3674 = v3486
	goto L956
L1002:
	;
	if v3572 < int32(5) {
		v3674 = v3295
		goto L956
	} else {
		goto L1018
	}
L1003:
	;
	v3572 = int32(0)
	goto L1002
L1004:
	;
	goto L1005
L1005:
	;
	v3503 = v3498 & int32(3)
	if base.Ui32(v3498) < base.Ui32(int32(4)) {
		goto L1008
	} else {
		goto L1009
	}
L1006:
	;
	v3572 = v3561
	goto L1002
L1007:
	;
	v3545 = v3539
	v3546 = v3540
	v3550 = v3491
	goto L1015
L1008:
	;
	v3539 = v3490
	v3540 = int32(0)
	goto L1007
L1009:
	;
	goto L1010
L1010:
	;
	v3510 = v3490
	v3511 = int32(0)
	v3514 = v3491
	goto L1011
L1011:
	;
	v3516 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3510))))
	v3517 = int32(-65)
	v3520 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3510)+1)))
	v3524 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3510)+2)))
	v3528 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3510)+3)))
	v3531 = v3511 + base.B2i32(v3517 < v3516) + base.B2i32(v3517 < v3520) + base.B2i32(v3517 < v3524) + base.B2i32(v3517 < v3528)
	v3532 = int32(4)
	v3533 = v3510 + v3532
	v3535 = v3514 + v3532
	if v3535 != v3498&int32(-4) {
		v3510 = v3533
		v3511 = v3531
		v3514 = v3535
		goto L1011
	} else {
		goto L1013
	}
L1012:
	;
	if v3503 == int32(0) {
		v3561 = v3531
		goto L1006
	} else {
		goto L1014
	}
L1013:
	;
	goto L1012
L1014:
	;
	v3539 = v3533
	v3540 = v3531
	goto L1007
L1015:
	;
	v3551 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3545))))
	v3554 = v3546 + base.B2i32(int32(-65) < v3551)
	v3555 = int32(1)
	v3558 = v3550 + v3555
	if v3558 != v3503 {
		v3545 = v3545 + v3555
		v3546 = v3554
		v3550 = v3558
		goto L1015
	} else {
		goto L1017
	}
L1016:
	;
	v3561 = v3554
	goto L1006
L1017:
	;
	goto L1016
L1018:
	;
	v3577 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_75))
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L4
	} else {
		goto L1019
	}
L1019:
	;
	if int32(0) <= v3577 {
		goto L959
	} else {
		goto L1020
	}
L1020:
	;
	v3674 = v3577
	goto L956
L1021:
	;
	if v3663 < int32(5) {
		v3674 = v3295
		goto L956
	} else {
		goto L1037
	}
L1022:
	;
	v3663 = int32(0)
	goto L1021
L1023:
	;
	goto L1024
L1024:
	;
	v3594 = v3589 & int32(3)
	if base.Ui32(v3589) < base.Ui32(int32(4)) {
		goto L1027
	} else {
		goto L1028
	}
L1025:
	;
	v3663 = v3652
	goto L1021
L1026:
	;
	v3636 = v3630
	v3637 = v3631
	v3641 = v3582
	goto L1034
L1027:
	;
	v3630 = v3581
	v3631 = int32(0)
	goto L1026
L1028:
	;
	goto L1029
L1029:
	;
	v3601 = v3581
	v3602 = int32(0)
	v3605 = v3582
	goto L1030
L1030:
	;
	v3607 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3601))))
	v3608 = int32(-65)
	v3611 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3601)+1)))
	v3615 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3601)+2)))
	v3619 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3601)+3)))
	v3622 = v3602 + base.B2i32(v3608 < v3607) + base.B2i32(v3608 < v3611) + base.B2i32(v3608 < v3615) + base.B2i32(v3608 < v3619)
	v3623 = int32(4)
	v3624 = v3601 + v3623
	v3626 = v3605 + v3623
	if v3626 != v3589&int32(-4) {
		v3601 = v3624
		v3602 = v3622
		v3605 = v3626
		goto L1030
	} else {
		goto L1032
	}
L1031:
	;
	if v3594 == int32(0) {
		v3652 = v3622
		goto L1025
	} else {
		goto L1033
	}
L1032:
	;
	goto L1031
L1033:
	;
	v3630 = v3624
	v3631 = v3622
	goto L1026
L1034:
	;
	v3642 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3636))))
	v3645 = v3637 + base.B2i32(int32(-65) < v3642)
	v3646 = int32(1)
	v3649 = v3641 + v3646
	if v3649 != v3594 {
		v3636 = v3636 + v3646
		v3637 = v3645
		v3641 = v3649
		goto L1034
	} else {
		goto L1036
	}
L1035:
	;
	v3652 = v3645
	goto L1025
L1036:
	;
	goto L1035
L1037:
	;
	v3668 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_76))
	mBase = m.M
	v3669 = m.ExcPending
	if v3669 != 0 {
		goto L4
	} else {
		goto L1038
	}
L1038:
	;
	if v3668 < int32(0) {
		v3674 = v3668
		goto L956
	} else {
		goto L1039
	}
L1039:
	;
	goto L959
L1040:
	;
	v3680 = int32(0)
	v3681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3681
	v3684 = v3681 + int32(5)
	v3685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3685 <= v3684 {
		v3797 = v3680
		goto L1045
	} else {
		goto L1046
	}
L1041:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2752
	goto L1040
L1042:
	;
	goto L1043
L1043:
	;
	if v3674 < int32(0) {
		v3966 = v3674
		goto L44
	} else {
		goto L1044
	}
L1044:
	;
	goto L1040
L1045:
	;
	if int32(0) <= v3797 {
		goto L860
	} else {
		goto L1071
	}
L1046:
	;
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3687+v3684))))
	if v3689 != int32(170) {
		v3797 = v3680
		goto L1045
	} else {
		goto L1047
	}
L1047:
	;
	v3694 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_77), int32(3))
	mBase = m.M
	v3695 = m.ExcPending
	if v3695 != 0 {
		goto L4
	} else {
		goto L1048
	}
L1048:
	;
	if v3694 == int32(0) {
		v3797 = v3680
		goto L1045
	} else {
		goto L1049
	}
L1049:
	;
	v3698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3698
	v3700 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3701 = int32(0)
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3700-int32(4))))
	if v3708 == v3701 {
		goto L1051
	} else {
		goto L1052
	}
L1050:
	;
	if v3782 < int32(5) {
		v3797 = v3680
		goto L1045
	} else {
		goto L1066
	}
L1051:
	;
	v3782 = int32(0)
	goto L1050
L1052:
	;
	goto L1053
L1053:
	;
	v3713 = v3708 & int32(3)
	if base.Ui32(v3708) < base.Ui32(int32(4)) {
		goto L1056
	} else {
		goto L1057
	}
L1054:
	;
	v3782 = v3771
	goto L1050
L1055:
	;
	v3755 = v3749
	v3756 = v3750
	v3760 = v3701
	goto L1063
L1056:
	;
	v3749 = v3700
	v3750 = int32(0)
	goto L1055
L1057:
	;
	goto L1058
L1058:
	;
	v3720 = v3700
	v3721 = int32(0)
	v3724 = v3701
	goto L1059
L1059:
	;
	v3726 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3720))))
	v3727 = int32(-65)
	v3730 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3720)+1)))
	v3734 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3720)+2)))
	v3738 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3720)+3)))
	v3741 = v3721 + base.B2i32(v3727 < v3726) + base.B2i32(v3727 < v3730) + base.B2i32(v3727 < v3734) + base.B2i32(v3727 < v3738)
	v3742 = int32(4)
	v3743 = v3720 + v3742
	v3745 = v3724 + v3742
	if v3745 != v3708&int32(-4) {
		v3720 = v3743
		v3721 = v3741
		v3724 = v3745
		goto L1059
	} else {
		goto L1061
	}
L1060:
	;
	if v3713 == int32(0) {
		v3771 = v3741
		goto L1054
	} else {
		goto L1062
	}
L1061:
	;
	goto L1060
L1062:
	;
	v3749 = v3743
	v3750 = v3741
	goto L1055
L1063:
	;
	v3761 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3755))))
	v3764 = v3756 + base.B2i32(int32(-65) < v3761)
	v3765 = int32(1)
	v3768 = v3760 + v3765
	if v3768 != v3713 {
		v3755 = v3755 + v3765
		v3756 = v3764
		v3760 = v3768
		goto L1063
	} else {
		goto L1065
	}
L1064:
	;
	v3771 = v3764
	goto L1054
L1065:
	;
	goto L1064
L1066:
	;
	v3785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v3785)+4)) = int64(1)
	v3791 = F_slice_from_s(m, l0, int32(6), int32(_a_F_arabic_UTF_8_stem_78))
	mBase = m.M
	v3792 = m.ExcPending
	if v3792 != 0 {
		goto L4
	} else {
		goto L1067
	}
L1067:
	;
	if int32(0) <= v3791 {
		goto L1068
	} else {
		goto L1069
	}
L1068:
	;
	v3795 = int32(1)
	goto L1070
L1069:
	;
	v3795 = v3791
	goto L1070
L1070:
	;
	v3797 = v3795
	goto L1045
L1071:
	;
	v3802 = v3797
	v3804 = int32(base.Ui32(v3797) >> (uint(int32(31)) % 32))
	goto L871
L1072:
	;
	goto L860
L1073:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2225
	v3848 = v2225
	goto L1080
L1074:
	;
	v3817 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3817+v3815))))
	if base.B2i32(v3819&int32(224) != int32(160))|base.B2i32(int32(1)<<(uint(v3819)%32)&int32(124) == int32(0)) != 0 {
		goto L1073
	} else {
		goto L1075
	}
L1075:
	;
	v3833 = F_find_among_b(m, l0, int32(_a_F_arabic_UTF_8_stem_79), int32(5))
	mBase = m.M
	v3834 = m.ExcPending
	if v3834 != 0 {
		goto L4
	} else {
		goto L1076
	}
L1076:
	;
	if v3833 == int32(0) {
		goto L1073
	} else {
		goto L1077
	}
L1077:
	;
	v3837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3837
	v3841 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_80))
	mBase = m.M
	v3842 = m.ExcPending
	if v3842 != 0 {
		goto L4
	} else {
		goto L1078
	}
L1078:
	;
	if v3841 < int32(0) {
		v3966 = v3841
		goto L44
	} else {
		goto L1079
	}
L1079:
	;
	goto L1073
L1080:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3848
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3857 = v3848 + int32(1)
	v3858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3858 <= v3857 {
		v3880 = v3858
		v3881 = v3855
		goto L1084
	} else {
		goto L1085
	}
L1081:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2225
	v3966 = int32(1)
	goto L44
L1082:
	;
	goto L1081
L1083:
	;
	v3939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3939
	switch v3875 - int32(1) {
	case 0:
		goto L1113
	case 1:
		goto L1112
	case 2:
		goto L1111
	default:
		goto L1110
	}
L1084:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3848
	goto L1091
L1085:
	;
	v3861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3857+v3855))))
	if base.B2i32(v3861&int32(224) != int32(160))|base.B2i32(int32(1)<<(uint(v3861)%32)&int32(124) == int32(0)) != 0 {
		v3880 = v3858
		v3881 = v3855
		goto L1084
	} else {
		goto L1086
	}
L1086:
	;
	v3875 = F_find_among(m, l0, int32(_a_F_arabic_UTF_8_stem_81), int32(5))
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L4
	} else {
		goto L1087
	}
L1087:
	;
	if v3875 != 0 {
		goto L1083
	} else {
		goto L1088
	}
L1088:
	;
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3880 = v3877
	v3881 = v3878
	goto L1084
L1089:
	;
	if v3934 < int32(0) {
		goto L1082
	} else {
		goto L1109
	}
L1091:
	;
	goto L1092
L1092:
	;
	goto L1093
L1093:
	;
	v3889 = v3848
	v3891 = int32(1)
	goto L1096
L1095:
	;
	v3934 = v3919
	goto L1089
L1096:
	;
	if v3880 <= v3889 {
		goto L1098
	} else {
		goto L1099
	}
L1097:
	;
	goto L1095
L1098:
	;
	v3934 = int32(-1)
	goto L1089
L1099:
	;
	goto L1100
L1100:
	;
	v3896 = v3889 + int32(1)
	v3898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3881+v3889))))
	if base.Ui32(v3898) < base.Ui32(int32(192)) {
		v3919 = v3896
		goto L1101
	} else {
		goto L1102
	}
L1101:
	;
	v3920 = int32(1)
	if v3920 < v3891 {
		v3889 = v3919
		v3891 = v3891 - v3920
		goto L1096
	} else {
		goto L1108
	}
L1102:
	;
	if v3880 <= v3896 {
		v3919 = v3896
		goto L1101
	} else {
		goto L1103
	}
L1103:
	;
	v3905 = v3896
	goto L1104
L1104:
	;
	v3908 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3881+v3905))))
	if int32(-65) < v3908 {
		v3919 = v3905
		goto L1101
	} else {
		goto L1106
	}
L1105:
	;
	v3919 = v3880
	goto L1101
L1106:
	;
	v3912 = v3905 + int32(1)
	if v3912 != v3880 {
		v3905 = v3912
		goto L1104
	} else {
		goto L1107
	}
L1107:
	;
	goto L1105
L1108:
	;
	goto L1097
L1109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3934
	v3848 = v3934
	goto L1080
L1110:
	;
	v3962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3848 = v3962
	goto L1080
L1111:
	;
	v3957 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_82))
	mBase = m.M
	v3958 = m.ExcPending
	if v3958 != 0 {
		goto L4
	} else {
		goto L1118
	}
L1112:
	;
	v3951 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_83))
	mBase = m.M
	v3952 = m.ExcPending
	if v3952 != 0 {
		goto L4
	} else {
		goto L1116
	}
L1113:
	;
	v3945 = F_slice_from_s(m, l0, int32(2), int32(_a_F_arabic_UTF_8_stem_84))
	mBase = m.M
	v3946 = m.ExcPending
	if v3946 != 0 {
		goto L4
	} else {
		goto L1114
	}
L1114:
	;
	if int32(0) <= v3945 {
		goto L1110
	} else {
		goto L1115
	}
L1115:
	;
	v3966 = v3945
	goto L44
L1116:
	;
	if int32(0) <= v3951 {
		goto L1110
	} else {
		goto L1117
	}
L1117:
	;
	v3966 = v3951
	goto L44
L1118:
	;
	if v3957 < int32(0) {
		v3966 = v3957
		goto L44
	} else {
		goto L1119
	}
L1119:
	;
	goto L1110
}
