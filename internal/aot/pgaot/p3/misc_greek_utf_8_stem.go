package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_greek_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
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
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
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
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
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
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1605 int32
	_ = v1605
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1758 int32
	_ = v1758
	var v1774 int32
	_ = v1774
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1996 int32
	_ = v1996
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2040 int32
	_ = v2040
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2158 int32
	_ = v2158
	var v2174 int32
	_ = v2174
	var v2181 int32
	_ = v2181
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2232 int32
	_ = v2232
	var v2236 int32
	_ = v2236
	var v2240 int32
	_ = v2240
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2270 int32
	_ = v2270
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2289 int32
	_ = v2289
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2401 int32
	_ = v2401
	var v2417 int32
	_ = v2417
	var v2424 int32
	_ = v2424
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2517 int32
	_ = v2517
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2536 int32
	_ = v2536
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2562 int32
	_ = v2562
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2572 int32
	_ = v2572
	var v2578 int32
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2582 int32
	_ = v2582
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2598 int32
	_ = v2598
	var v2602 int32
	_ = v2602
	var v2606 int32
	_ = v2606
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2628 int32
	_ = v2628
	var v2633 int32
	_ = v2633
	var v2636 int32
	_ = v2636
	var v2639 int32
	_ = v2639
	var v2643 int32
	_ = v2643
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2671 int32
	_ = v2671
	var v2673 int32
	_ = v2673
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2682 int32
	_ = v2682
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2705 int32
	_ = v2705
	var v2707 int32
	_ = v2707
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2737 int32
	_ = v2737
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2747 int32
	_ = v2747
	var v2750 int32
	_ = v2750
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2759 int32
	_ = v2759
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2783 int32
	_ = v2783
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2806 int32
	_ = v2806
	var v2808 int32
	_ = v2808
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2818 int32
	_ = v2818
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2827 int32
	_ = v2827
	var v2830 int32
	_ = v2830
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
	var v2845 int32
	_ = v2845
	var v2849 int32
	_ = v2849
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2856 int32
	_ = v2856
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2903 int32
	_ = v2903
	var v2906 int32
	_ = v2906
	var v2909 int32
	_ = v2909
	var v2912 int32
	_ = v2912
	var v2913 int32
	_ = v2913
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2924 int32
	_ = v2924
	var v2928 int32
	_ = v2928
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2950 int32
	_ = v2950
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2965 int32
	_ = v2965
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2974 int32
	_ = v2974
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2983 int32
	_ = v2983
	var v2985 int32
	_ = v2985
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3006 int32
	_ = v3006
	var v3008 int32
	_ = v3008
	var v3011 int32
	_ = v3011
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3026 int32
	_ = v3026
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3035 int32
	_ = v3035
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3044 int32
	_ = v3044
	var v3047 int32
	_ = v3047
	var v3050 int32
	_ = v3050
	var v3054 int32
	_ = v3054
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3085 int32
	_ = v3085
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3097 int32
	_ = v3097
	var v3098 int32
	_ = v3098
	var v3101 int32
	_ = v3101
	var v3104 int32
	_ = v3104
	var v3108 int32
	_ = v3108
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3114 int32
	_ = v3114
	var v3117 int32
	_ = v3117
	var v3121 int32
	_ = v3121
	var v3125 int32
	_ = v3125
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3134 int32
	_ = v3134
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3169 int32
	_ = v3169
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3175 int32
	_ = v3175
	var v3178 int32
	_ = v3178
	var v3182 int32
	_ = v3182
	var v3186 int32
	_ = v3186
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3204 int32
	_ = v3204
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3230 int32
	_ = v3230
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3243 int32
	_ = v3243
	var v3247 int32
	_ = v3247
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3256 int32
	_ = v3256
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3262 int32
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3265 int32
	_ = v3265
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3276 int32
	_ = v3276
	var v3277 int32
	_ = v3277
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3291 int32
	_ = v3291
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3313 int32
	_ = v3313
	var v3315 int32
	_ = v3315
	var v3317 int32
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3329 int32
	_ = v3329
	var v3332 int32
	_ = v3332
	var v3333 int32
	_ = v3333
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3344 int32
	_ = v3344
	var v3347 int32
	_ = v3347
	var v3351 int32
	_ = v3351
	var v3355 int32
	_ = v3355
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3364 int32
	_ = v3364
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3375 int32
	_ = v3375
	var v3379 int32
	_ = v3379
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v13 = v9
	goto L2
L1:
	;
	return v3379
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v13
	v20 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_0), int32(46))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v236 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v235-int32(4))))
	if v243 == v236 {
		goto L102
	} else {
		goto L103
	}
L4:
	;
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v24
	switch v20 - int32(1) {
	case 0:
		goto L32
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	case 4:
		goto L28
	case 5:
		goto L27
	case 6:
		goto L26
	case 7:
		goto L25
	case 8:
		goto L24
	case 9:
		goto L23
	case 10:
		goto L22
	case 11:
		goto L21
	case 12:
		goto L20
	case 13:
		goto L19
	case 14:
		goto L18
	case 15:
		goto L17
	case 16:
		goto L16
	case 17:
		goto L15
	case 18:
		goto L14
	case 19:
		goto L13
	case 20:
		goto L12
	case 21:
		goto L11
	case 22:
		goto L10
	case 23:
		goto L9
	case 24:
		goto L8
	default:
		goto L7
	}
L7:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v232
	goto L2
L8:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L83
L9:
	;
	v168 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L79
	}
L10:
	;
	v162 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_2))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L77
	}
L11:
	;
	v156 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L75
	}
L12:
	;
	v150 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L73
	}
L13:
	;
	v144 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_5))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L71
	}
L14:
	;
	v138 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_6))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L69
	}
L15:
	;
	v132 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_7))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L67
	}
L16:
	;
	v126 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_8))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L65
	}
L17:
	;
	v120 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_9))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L63
	}
L18:
	;
	v114 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_10))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L61
	}
L19:
	;
	v108 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_11))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L59
	}
L20:
	;
	v102 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_12))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L57
	}
L21:
	;
	v96 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_13))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L55
	}
L22:
	;
	v90 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_14))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L53
	}
L23:
	;
	v84 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_15))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L51
	}
L24:
	;
	v78 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L49
	}
L25:
	;
	v72 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_17))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L47
	}
L26:
	;
	v66 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_18))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L45
	}
L27:
	;
	v60 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_19))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L43
	}
L28:
	;
	v54 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_20))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L41
	}
L29:
	;
	v48 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_21))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L39
	}
L30:
	;
	v42 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_22))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L37
	}
L31:
	;
	v36 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_23))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L35
	}
L32:
	;
	v30 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_24))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	if int32(0) <= v30 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v3379 = v30
	goto L1
L35:
	;
	if int32(0) <= v36 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v3379 = v36
	goto L1
L37:
	;
	if int32(0) <= v42 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v3379 = v42
	goto L1
L39:
	;
	if int32(0) <= v48 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v3379 = v48
	goto L1
L41:
	;
	if int32(0) <= v54 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v3379 = v54
	goto L1
L43:
	;
	if int32(0) <= v60 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v3379 = v60
	goto L1
L45:
	;
	if int32(0) <= v66 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v3379 = v66
	goto L1
L47:
	;
	if int32(0) <= v72 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v3379 = v72
	goto L1
L49:
	;
	if int32(0) <= v78 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v3379 = v78
	goto L1
L51:
	;
	if int32(0) <= v84 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v3379 = v84
	goto L1
L53:
	;
	if int32(0) <= v90 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v3379 = v90
	goto L1
L55:
	;
	if int32(0) <= v96 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v3379 = v96
	goto L1
L57:
	;
	if int32(0) <= v102 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v3379 = v102
	goto L1
L59:
	;
	if int32(0) <= v108 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v3379 = v108
	goto L1
L61:
	;
	if int32(0) <= v114 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v3379 = v114
	goto L1
L63:
	;
	if int32(0) <= v120 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v3379 = v120
	goto L1
L65:
	;
	if int32(0) <= v126 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v3379 = v126
	goto L1
L67:
	;
	if int32(0) <= v132 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	v3379 = v132
	goto L1
L69:
	;
	if int32(0) <= v138 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	v3379 = v138
	goto L1
L71:
	;
	if int32(0) <= v144 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v3379 = v144
	goto L1
L73:
	;
	if int32(0) <= v150 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v3379 = v150
	goto L1
L75:
	;
	if int32(0) <= v156 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v3379 = v156
	goto L1
L77:
	;
	if int32(0) <= v162 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	v3379 = v162
	goto L1
L79:
	;
	if int32(0) <= v168 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v3379 = v168
	goto L1
L81:
	;
	if v226 < int32(0) {
		goto L4
	} else {
		goto L100
	}
L83:
	;
	goto L84
L84:
	;
	goto L85
L85:
	;
	v181 = v24
	v183 = int32(1)
	goto L88
L87:
	;
	v226 = v208
	goto L81
L88:
	;
	if v181 <= v174 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L87
L90:
	;
	v226 = int32(-1)
	goto L81
L91:
	;
	goto L92
L92:
	;
	v188 = v181 - int32(1)
	v190 = int32(*(*int8)(unsafe.Add(mBase, uint32(v173+v188))))
	if base.B2i32(int32(0) <= v190)|base.B2i32(v188 <= v174) != 0 {
		v208 = v188
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v212 = int32(1)
	if v212 < v183 {
		v181 = v208
		v183 = v183 - v212
		goto L88
	} else {
		goto L99
	}
L94:
	;
	v196 = v188
	goto L95
L95:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v196))))
	if base.Ui32(int32(191)) < base.Ui32(v201) {
		v208 = v196
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v208 = v174
	goto L93
L97:
	;
	v205 = v196 - int32(1)
	if v174 < v205 {
		v196 = v205
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	goto L89
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v226
	goto L7
L101:
	;
	if v317 < int32(3) {
		v3379 = int32(0)
		goto L1
	} else {
		goto L117
	}
L102:
	;
	v317 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v248 = v243 & int32(3)
	if base.Ui32(v243) < base.Ui32(int32(4)) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v317 = v306
	goto L101
L106:
	;
	v290 = v284
	v291 = v285
	v295 = v236
	goto L114
L107:
	;
	v284 = v235
	v285 = int32(0)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v255 = v235
	v256 = int32(0)
	v259 = v236
	goto L110
L110:
	;
	v261 = int32(*(*int8)(unsafe.Add(mBase, uint32(v255))))
	v262 = int32(-65)
	v265 = int32(*(*int8)(unsafe.Add(mBase, uint32(v255)+1)))
	v269 = int32(*(*int8)(unsafe.Add(mBase, uint32(v255)+2)))
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v255)+3)))
	v276 = v256 + base.B2i32(v262 < v261) + base.B2i32(v262 < v265) + base.B2i32(v262 < v269) + base.B2i32(v262 < v273)
	v277 = int32(4)
	v278 = v255 + v277
	v280 = v259 + v277
	if v280 != v243&int32(-4) {
		v255 = v278
		v256 = v276
		v259 = v280
		goto L110
	} else {
		goto L112
	}
L111:
	;
	if v248 == int32(0) {
		v306 = v276
		goto L105
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v284 = v278
	v285 = v276
	goto L106
L114:
	;
	v296 = int32(*(*int8)(unsafe.Add(mBase, uint32(v290))))
	v299 = v291 + base.B2i32(int32(-65) < v296)
	v300 = int32(1)
	v303 = v295 + v300
	if v303 != v248 {
		v290 = v290 + v300
		v291 = v299
		v295 = v303
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v306 = v299
	goto L105
L116:
	;
	goto L115
L117:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v328 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_25), int32(40))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	if v328 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v330
	switch v328 - int32(1) {
	case 0:
		goto L133
	case 1:
		goto L132
	case 2:
		goto L131
	case 3:
		goto L130
	case 4:
		goto L129
	case 5:
		goto L128
	case 6:
		goto L127
	case 7:
		goto L126
	case 8:
		goto L125
	case 9:
		goto L124
	case 10:
		goto L123
	default:
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v405 = v323 - v325
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v407 = v405 + v406
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v407
	v412 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_26), int32(14))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L157
	}
L122:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v401))) = int32(0)
	goto L121
L123:
	;
	v396 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_27))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L154
	}
L124:
	;
	v390 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_28))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L152
	}
L125:
	;
	v384 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_29))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L150
	}
L126:
	;
	v378 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_30))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L148
	}
L127:
	;
	v372 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_31))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L146
	}
L128:
	;
	v366 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_32))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L144
	}
L129:
	;
	v360 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_33))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L142
	}
L130:
	;
	v354 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_34))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L5
	} else {
		goto L140
	}
L131:
	;
	v348 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_35))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L138
	}
L132:
	;
	v342 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_36))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L136
	}
L133:
	;
	v336 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_37))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	if int32(0) <= v336 {
		goto L122
	} else {
		goto L135
	}
L135:
	;
	v3379 = v336
	goto L1
L136:
	;
	if int32(0) <= v342 {
		goto L122
	} else {
		goto L137
	}
L137:
	;
	v3379 = v342
	goto L1
L138:
	;
	if int32(0) <= v348 {
		goto L122
	} else {
		goto L139
	}
L139:
	;
	v3379 = v348
	goto L1
L140:
	;
	if int32(0) <= v354 {
		goto L122
	} else {
		goto L141
	}
L141:
	;
	v3379 = v354
	goto L1
L142:
	;
	if int32(0) <= v360 {
		goto L122
	} else {
		goto L143
	}
L143:
	;
	v3379 = v360
	goto L1
L144:
	;
	if int32(0) <= v366 {
		goto L122
	} else {
		goto L145
	}
L145:
	;
	v3379 = v366
	goto L1
L146:
	;
	if int32(0) <= v372 {
		goto L122
	} else {
		goto L147
	}
L147:
	;
	v3379 = v372
	goto L1
L148:
	;
	if int32(0) <= v378 {
		goto L122
	} else {
		goto L149
	}
L149:
	;
	v3379 = v378
	goto L1
L150:
	;
	if int32(0) <= v384 {
		goto L122
	} else {
		goto L151
	}
L151:
	;
	v3379 = v384
	goto L1
L152:
	;
	if int32(0) <= v390 {
		goto L122
	} else {
		goto L153
	}
L153:
	;
	v3379 = v390
	goto L1
L154:
	;
	if v396 < int32(0) {
		v3379 = v396
		goto L1
	} else {
		goto L155
	}
L155:
	;
	goto L122
L156:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v454 = v453 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v454
	v461 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_38), int32(7))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L171
	}
L157:
	;
	if v412 == int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v416
	v418 = F_slice_del(m, l0)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L159
	}
L159:
	;
	if v418 < int32(0) {
		v3379 = v418
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = int32(0)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v425
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v425
	v430 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_39), int32(31))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	if v430 == int32(0) {
		goto L156
	} else {
		goto L162
	}
L162:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v435 < v434 {
		goto L156
	} else {
		goto L163
	}
L163:
	;
	switch v430 - int32(1) {
	case 0:
		goto L165
	case 1:
		goto L164
	default:
		goto L156
	}
L164:
	;
	v447 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_40))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L168
	}
L165:
	;
	v441 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_41))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L5
	} else {
		goto L166
	}
L166:
	;
	if int32(0) <= v441 {
		goto L156
	} else {
		goto L167
	}
L167:
	;
	v3379 = v441
	goto L1
L168:
	;
	if v447 < int32(0) {
		v3379 = v447
		goto L1
	} else {
		goto L169
	}
L169:
	;
	goto L156
L170:
	;
	if v500 < int32(0) {
		v3379 = v500
		goto L1
	} else {
		goto L183
	}
L171:
	;
	if v461 == int32(0) {
		v500 = int32(0)
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v465
	v467 = F_slice_del(m, l0)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L5
	} else {
		goto L174
	}
L173:
	;
	v500 = v497
	goto L170
L174:
	;
	if v467 < int32(0) {
		v497 = v467
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v472 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v472
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v474
	v480 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_42), int32(8))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L5
	} else {
		goto L176
	}
L176:
	;
	if v480 == int32(0) {
		v500 = v472
		goto L170
	} else {
		goto L177
	}
L177:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v486 < v485 {
		v497 = int32(0)
		goto L173
	} else {
		goto L178
	}
L178:
	;
	v491 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_43))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L5
	} else {
		goto L179
	}
L179:
	;
	if int32(0) <= v491 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v495 = int32(1)
	goto L182
L181:
	;
	v495 = v491
	goto L182
L182:
	;
	v497 = v495
	goto L173
L183:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v504 = v503 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v504
	v509 = int32(6)
	v511 = int32(0)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v504-v514 < v509 {
		v524 = v511
		goto L189
	} else {
		goto L190
	}
L184:
	;
	if v598 < int32(0) {
		v3379 = v598
		goto L1
	} else {
		goto L210
	}
L185:
	;
	v598 = v593
	goto L184
L186:
	;
	v549 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_44), int32(7))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L5
	} else {
		goto L196
	}
L187:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v540 = v538 + (v504 - v503)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v540
	goto L186
L188:
	;
	if v524 == int32(0) {
		goto L187
	} else {
		goto L192
	}
L189:
	;
	goto L188
L190:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v520 = F_memcmp(m, v517+v504-v509, int32(_a_F_greek_UTF_8_stem_45), v509)
	mBase = m.M
	if v520 != 0 {
		v524 = v511
		goto L189
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504 - v509
	v524 = int32(1)
	goto L189
L192:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v527
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v529 < v527 {
		goto L187
	} else {
		goto L193
	}
L193:
	;
	v533 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_46))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	if int32(0) <= v533 {
		goto L186
	} else {
		goto L195
	}
L195:
	;
	v593 = v533
	goto L185
L196:
	;
	if v549 == int32(0) {
		v598 = int32(0)
		goto L184
	} else {
		goto L197
	}
L197:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v553
	v555 = F_slice_del(m, l0)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	if v555 < int32(0) {
		v593 = v555
		goto L185
	} else {
		goto L199
	}
L199:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v560 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = v560
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v562
	v568 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_47), int32(32))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	if v568 == int32(0) {
		v598 = v560
		goto L184
	} else {
		goto L201
	}
L201:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v574 < v573 {
		v593 = int32(0)
		goto L185
	} else {
		goto L202
	}
L202:
	;
	switch v568 - int32(1) {
	case 0:
		goto L205
	case 1:
		goto L204
	default:
		goto L203
	}
L203:
	;
	v593 = int32(1)
	goto L185
L204:
	;
	v586 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_48))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L208
	}
L205:
	;
	v580 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_49))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	if int32(0) <= v580 {
		goto L203
	} else {
		goto L207
	}
L207:
	;
	v593 = v580
	goto L185
L208:
	;
	if v586 < int32(0) {
		v593 = v586
		goto L185
	} else {
		goto L209
	}
L209:
	;
	goto L203
L210:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v602 = v601 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v602
	v609 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_50), int32(7))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L5
	} else {
		goto L212
	}
L211:
	;
	if v670 < int32(0) {
		v3379 = v670
		goto L1
	} else {
		goto L227
	}
L212:
	;
	if v609 == int32(0) {
		v670 = int32(0)
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v613
	v615 = F_slice_del(m, l0)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L5
	} else {
		goto L215
	}
L214:
	;
	v670 = v667
	goto L211
L215:
	;
	if v615 < int32(0) {
		v667 = v615
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v620 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = v620
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v622
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v622
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v622-int32(3) <= v626 {
		v670 = v620
		goto L211
	} else {
		goto L217
	}
L217:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631+v622-int32(1)))))
	if v635&int32(224) != int32(160) {
		v670 = int32(0)
		goto L211
	} else {
		goto L218
	}
L218:
	;
	v640 = int32(0)
	if int32(1)<<(uint(v635)%32)&int32(-2145255424) == v640 {
		v670 = v640
		goto L211
	} else {
		goto L219
	}
L219:
	;
	v650 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_51), int32(19))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	if v650 == int32(0) {
		v670 = int32(0)
		goto L211
	} else {
		goto L221
	}
L221:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v656 < v655 {
		v667 = int32(0)
		goto L214
	} else {
		goto L222
	}
L222:
	;
	v661 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_52))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	if int32(0) <= v661 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v665 = int32(1)
	goto L226
L225:
	;
	v665 = v661
	goto L226
L226:
	;
	v667 = v665
	goto L214
L227:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v674 = v673 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v674
	v681 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_53), int32(11))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L5
	} else {
		goto L229
	}
L228:
	;
	if v728 < int32(0) {
		v3379 = v728
		goto L1
	} else {
		goto L244
	}
L229:
	;
	if v681 == int32(0) {
		v728 = int32(0)
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v685
	v687 = F_slice_del(m, l0)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L5
	} else {
		goto L232
	}
L231:
	;
	v728 = v725
	goto L228
L232:
	;
	if v687 < int32(0) {
		v725 = v687
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v692 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v691))) = v692
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v694
	v700 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_54), int32(40))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L234
	}
L234:
	;
	if v700 == int32(0) {
		v728 = v692
		goto L228
	} else {
		goto L235
	}
L235:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v706 < v705 {
		v725 = int32(0)
		goto L231
	} else {
		goto L236
	}
L236:
	;
	switch v700 - int32(1) {
	case 0:
		goto L239
	case 1:
		goto L238
	default:
		goto L237
	}
L237:
	;
	v725 = int32(1)
	goto L231
L238:
	;
	v718 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_55))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L5
	} else {
		goto L242
	}
L239:
	;
	v712 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_56))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L5
	} else {
		goto L240
	}
L240:
	;
	if int32(0) <= v712 {
		goto L237
	} else {
		goto L241
	}
L241:
	;
	v725 = v712
	goto L231
L242:
	;
	if v718 < int32(0) {
		v725 = v718
		goto L231
	} else {
		goto L243
	}
L243:
	;
	goto L237
L244:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v732 = v731 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v732
	v739 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_57), int32(6))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L5
	} else {
		goto L246
	}
L245:
	;
	if v891 < int32(0) {
		v3379 = v891
		goto L1
	} else {
		goto L298
	}
L246:
	;
	if v739 == int32(0) {
		v891 = int32(0)
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v743
	v745 = F_slice_del(m, l0)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L5
	} else {
		goto L249
	}
L248:
	;
	v891 = v883
	goto L245
L249:
	;
	if v745 < int32(0) {
		v883 = v745
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = int32(0)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v752
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v752-int32(3) <= v756 {
		v792 = v756
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v883 = v881
	goto L248
L252:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v795 = v793 + (v752 - v755)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v795
	v798 = int32(0)
	if v795-int32(9) <= v792 {
		v881 = v798
		goto L251
	} else {
		goto L264
	}
L253:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760+v752-int32(1)))))
	if v764 != int32(181) {
		v792 = v756
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v769 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_58), int32(7))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v769 == int32(0) {
		v792 = v771
		goto L252
	} else {
		goto L256
	}
L256:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v771 < v774 {
		v792 = v771
		goto L252
	} else {
		goto L257
	}
L257:
	;
	v776 = int32(1)
	switch v769 - v776 {
	case 0:
		goto L259
	case 1:
		goto L258
	default:
		v883 = v776
		goto L248
	}
L258:
	;
	v787 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_59))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L262
	}
L259:
	;
	v781 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_60))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	if v781 < int32(0) {
		v881 = v781
		goto L251
	} else {
		goto L261
	}
L261:
	;
	v883 = v776
	goto L248
L262:
	;
	if v787 < int32(0) {
		v881 = v787
		goto L251
	} else {
		goto L263
	}
L263:
	;
	v883 = v776
	goto L248
L264:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802+v795-int32(1)))))
	switch v806 - int32(186) {
	case 0, 3:
		goto L265
	default:
		v881 = v798
		goto L251
	}
L265:
	;
	v811 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_61), int32(10))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L266
	}
L266:
	;
	if v811 == int32(0) {
		v881 = v798
		goto L251
	} else {
		goto L267
	}
L267:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v815
	v817 = int32(1)
	switch v811 - v817 {
	case 0:
		goto L277
	case 1:
		goto L276
	case 2:
		goto L275
	case 3:
		goto L274
	case 4:
		goto L273
	case 5:
		goto L272
	case 6:
		goto L271
	case 7:
		goto L270
	case 8:
		goto L269
	case 9:
		goto L268
	default:
		v883 = v817
		goto L248
	}
L268:
	;
	v876 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_62))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L5
	} else {
		goto L296
	}
L269:
	;
	v870 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_63))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L5
	} else {
		goto L294
	}
L270:
	;
	v864 = F_slice_from_s(m, l0, int32(16), int32(_a_F_greek_UTF_8_stem_64))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L5
	} else {
		goto L292
	}
L271:
	;
	v858 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_65))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L5
	} else {
		goto L290
	}
L272:
	;
	v852 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_66))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L5
	} else {
		goto L288
	}
L273:
	;
	v846 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_67))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L5
	} else {
		goto L286
	}
L274:
	;
	v840 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_68))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L5
	} else {
		goto L284
	}
L275:
	;
	v834 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_69))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L5
	} else {
		goto L282
	}
L276:
	;
	v828 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_70))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L280
	}
L277:
	;
	v822 = F_slice_from_s(m, l0, int32(12), int32(_a_F_greek_UTF_8_stem_71))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	if v822 < int32(0) {
		v881 = v822
		goto L251
	} else {
		goto L279
	}
L279:
	;
	v883 = v817
	goto L248
L280:
	;
	if v828 < int32(0) {
		v881 = v828
		goto L251
	} else {
		goto L281
	}
L281:
	;
	v883 = v817
	goto L248
L282:
	;
	if v834 < int32(0) {
		v881 = v834
		goto L251
	} else {
		goto L283
	}
L283:
	;
	v883 = v817
	goto L248
L284:
	;
	if v840 < int32(0) {
		v881 = v840
		goto L251
	} else {
		goto L285
	}
L285:
	;
	v883 = v817
	goto L248
L286:
	;
	if v846 < int32(0) {
		v881 = v846
		goto L251
	} else {
		goto L287
	}
L287:
	;
	v883 = v817
	goto L248
L288:
	;
	if v852 < int32(0) {
		v881 = v852
		goto L251
	} else {
		goto L289
	}
L289:
	;
	v883 = v817
	goto L248
L290:
	;
	if v858 < int32(0) {
		v881 = v858
		goto L251
	} else {
		goto L291
	}
L291:
	;
	v883 = v817
	goto L248
L292:
	;
	if v864 < int32(0) {
		v881 = v864
		goto L251
	} else {
		goto L293
	}
L293:
	;
	v883 = v817
	goto L248
L294:
	;
	if v870 < int32(0) {
		v881 = v870
		goto L251
	} else {
		goto L295
	}
L295:
	;
	v883 = v817
	goto L248
L296:
	;
	if int32(0) <= v876 {
		v883 = v817
		goto L248
	} else {
		goto L297
	}
L297:
	;
	v881 = v876
	goto L251
L298:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v895 = v894 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v895
	v897 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v895
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v895-int32(9) <= v900 {
		v959 = v897
		goto L300
	} else {
		goto L301
	}
L299:
	;
	if v962 < int32(0) {
		v3379 = v962
		goto L1
	} else {
		goto L316
	}
L300:
	;
	v962 = v959
	goto L299
L301:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904+v895-int32(1)))))
	switch v908 - int32(177) {
	case 0, 8:
		goto L302
	default:
		v959 = v897
		goto L300
	}
L302:
	;
	v913 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_72), int32(4))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L5
	} else {
		goto L303
	}
L303:
	;
	if v913 == int32(0) {
		v959 = v897
		goto L300
	} else {
		goto L304
	}
L304:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v917
	v919 = F_slice_del(m, l0)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L5
	} else {
		goto L305
	}
L305:
	;
	if v919 < int32(0) {
		v959 = v919
		goto L300
	} else {
		goto L306
	}
L306:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v924 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v923))) = v924
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v926
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v926
	v931 = v926 - int32(1)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v931 <= v932 {
		v962 = v924
		goto L299
	} else {
		goto L307
	}
L307:
	;
	v934 = int32(0)
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935+v931))))
	switch v937 - int32(131) {
	case 0, 4:
		goto L308
	default:
		v959 = v934
		goto L300
	}
L308:
	;
	v943 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_73), int32(2))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L5
	} else {
		goto L309
	}
L309:
	;
	if v943 == int32(0) {
		v962 = int32(0)
		goto L299
	} else {
		goto L310
	}
L310:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v948 < v947 {
		v959 = v934
		goto L300
	} else {
		goto L311
	}
L311:
	;
	v953 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_74))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L5
	} else {
		goto L312
	}
L312:
	;
	if int32(0) <= v953 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v957 = int32(1)
	goto L315
L314:
	;
	v957 = v953
	goto L315
L315:
	;
	v959 = v957
	goto L300
L316:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v966 = v965 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v966
	v973 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_75), int32(8))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L5
	} else {
		goto L318
	}
L317:
	;
	if v1054 < int32(0) {
		v3379 = v1054
		goto L1
	} else {
		goto L343
	}
L318:
	;
	if v973 == int32(0) {
		v1054 = int32(0)
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v977
	v979 = F_slice_del(m, l0)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L5
	} else {
		goto L321
	}
L320:
	;
	v1054 = v1048
	goto L317
L321:
	;
	if v979 < int32(0) {
		v1048 = v979
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v983))) = int32(0)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v986
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v992 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_76), int32(46))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L5
	} else {
		goto L326
	}
L323:
	;
	v1048 = v1047
	goto L320
L324:
	;
	v1042 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_77))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L5
	} else {
		goto L341
	}
L325:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1010 = v1008 + (v986 - v989)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1010
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1010
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1010
	v1014 = int32(6)
	v1016 = int32(0)
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1010-v1019 < v1014 {
		v1029 = v1016
		goto L333
	} else {
		goto L334
	}
L326:
	;
	if v992 == int32(0) {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v997 < v996 {
		goto L325
	} else {
		goto L328
	}
L328:
	;
	v999 = int32(1)
	switch v992 - v999 {
	case 0:
		goto L324
	case 1:
		goto L329
	default:
		v1048 = v999
		goto L320
	}
L329:
	;
	v1004 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_78))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	if v1004 < int32(0) {
		v1047 = v1004
		goto L323
	} else {
		goto L331
	}
L331:
	;
	v1048 = v999
	goto L320
L332:
	;
	if v1029 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	goto L332
L334:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1025 = F_memcmp(m, v1022+v1010-v1014, int32(_a_F_greek_UTF_8_stem_79), v1014)
	mBase = m.M
	if v1025 != 0 {
		v1029 = v1016
		goto L333
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1010 - v1014
	v1029 = int32(1)
	goto L333
L336:
	;
	v1047 = int32(0)
	goto L323
L337:
	;
	goto L338
L338:
	;
	v1035 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_80))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	if v1035 < int32(0) {
		v1047 = v1035
		goto L323
	} else {
		goto L340
	}
L340:
	;
	v1054 = int32(1)
	goto L317
L341:
	;
	if int32(0) <= v1042 {
		v1048 = v999
		goto L320
	} else {
		goto L342
	}
L342:
	;
	v1047 = v1042
	goto L323
L343:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1058 = v1057 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1058
	v1060 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1058
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1058-int32(7) <= v1063 {
		v1150 = v1060
		goto L345
	} else {
		goto L346
	}
L344:
	;
	if v1156 < int32(0) {
		v3379 = v1156
		goto L1
	} else {
		goto L368
	}
L345:
	;
	v1156 = v1150
	goto L344
L346:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1069 = int32(1)
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067+v1058-v1069))))
	if base.B2i32(v1071&int32(224) != int32(160))|base.B2i32(v1069<<(uint(v1071)%32)&int32(-1610481664) == int32(0)) != 0 {
		v1150 = v1060
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1085 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_81), int32(3))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L5
	} else {
		goto L348
	}
L348:
	;
	if v1085 == int32(0) {
		v1150 = v1060
		goto L345
	} else {
		goto L349
	}
L349:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1089
	v1091 = F_slice_del(m, l0)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L5
	} else {
		goto L350
	}
L350:
	;
	if v1091 < int32(0) {
		v1150 = v1091
		goto L345
	} else {
		goto L351
	}
L351:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1095))) = int32(0)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1098
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1104 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_82), int32(4))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L5
	} else {
		goto L352
	}
L352:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1104 == int32(0) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1121 = v1119 + (v1098 - v1101)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1121
	v1125 = int32(0)
	v1127 = v1121 - int32(1)
	if v1127 <= v1106 {
		v1150 = v1125
		goto L345
	} else {
		goto L360
	}
L354:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1106 < v1109 {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1114 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_83))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L5
	} else {
		goto L356
	}
L356:
	;
	if int32(0) <= v1114 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1118 = int32(1)
	goto L359
L358:
	;
	v1118 = v1114
	goto L359
L359:
	;
	v1156 = v1118
	goto L344
L360:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129+v1127))))
	switch v1131 - int32(181) {
	case 0, 8:
		goto L361
	default:
		v1150 = v1125
		goto L345
	}
L361:
	;
	v1136 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_84), int32(2))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L362
	}
L362:
	;
	if v1136 == int32(0) {
		v1150 = v1125
		goto L345
	} else {
		goto L363
	}
L363:
	;
	v1143 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_85))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L5
	} else {
		goto L364
	}
L364:
	;
	if int32(0) <= v1143 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1147 = int32(1)
	goto L367
L366:
	;
	v1147 = v1143
	goto L367
L367:
	;
	v1150 = v1147
	goto L345
L368:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1160 = v1159 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1160
	v1167 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_86), int32(4))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L5
	} else {
		goto L370
	}
L369:
	;
	if v1206 < int32(0) {
		v3379 = v1206
		goto L1
	} else {
		goto L382
	}
L370:
	;
	if v1167 == int32(0) {
		v1206 = int32(0)
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1171
	v1173 = F_slice_del(m, l0)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L5
	} else {
		goto L373
	}
L372:
	;
	v1206 = v1203
	goto L369
L373:
	;
	if v1173 < int32(0) {
		v1203 = v1173
		goto L372
	} else {
		goto L374
	}
L374:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1177))) = v1178
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1180
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1180
	v1186 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_87), int32(7))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L5
	} else {
		goto L375
	}
L375:
	;
	if v1186 == int32(0) {
		v1206 = v1178
		goto L369
	} else {
		goto L376
	}
L376:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1192 < v1191 {
		v1203 = int32(0)
		goto L372
	} else {
		goto L377
	}
L377:
	;
	v1197 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_88))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L5
	} else {
		goto L378
	}
L378:
	;
	if int32(0) <= v1197 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1201 = int32(1)
	goto L381
L380:
	;
	v1201 = v1197
	goto L381
L381:
	;
	v1203 = v1201
	goto L372
L382:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1210 = v1209 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1210
	v1212 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1210
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1210-int32(7) <= v1215 {
		v1261 = v1212
		goto L383
	} else {
		goto L384
	}
L383:
	;
	if v1261 < int32(0) {
		v3379 = v1261
		goto L1
	} else {
		goto L396
	}
L384:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219+v1210-int32(1)))))
	if base.B2i32(v1223 != int32(189))&base.B2i32(v1223 != int32(131)) != 0 {
		v1261 = v1212
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v1231 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_89), int32(2))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L5
	} else {
		goto L386
	}
L386:
	;
	if v1231 == int32(0) {
		v1261 = v1212
		goto L383
	} else {
		goto L387
	}
L387:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1235
	v1237 = F_slice_del(m, l0)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L5
	} else {
		goto L388
	}
L388:
	;
	if v1237 < int32(0) {
		v1261 = v1237
		goto L383
	} else {
		goto L389
	}
L389:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1246 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_90), int32(10))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L5
	} else {
		goto L390
	}
L390:
	;
	if v1246 != 0 {
		v1261 = int32(0)
		goto L383
	} else {
		goto L391
	}
L391:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1250 = v1248 + (v1241 - v1242)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1250
	v1254 = F_insert_s(m, l0, v1250, v1250, int32(4), int32(_a_F_greek_UTF_8_stem_91))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L5
	} else {
		goto L392
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1250
	if int32(0) <= v1254 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1260 = int32(1)
	goto L395
L394:
	;
	v1260 = v1254
	goto L395
L395:
	;
	v1261 = v1260
	goto L383
L396:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1267 = v1266 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1267
	v1269 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1267
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1267-int32(7) <= v1272 {
		v1335 = v1269
		goto L398
	} else {
		goto L399
	}
L397:
	;
	if v1338 < int32(0) {
		v3379 = v1338
		goto L1
	} else {
		goto L413
	}
L398:
	;
	v1338 = v1335
	goto L397
L399:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1276+v1267-int32(1)))))
	if base.B2i32(v1280 != int32(189))&base.B2i32(v1280 != int32(131)) != 0 {
		v1335 = v1269
		goto L398
	} else {
		goto L400
	}
L400:
	;
	v1288 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_92), int32(2))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L5
	} else {
		goto L401
	}
L401:
	;
	if v1288 == int32(0) {
		v1335 = v1269
		goto L398
	} else {
		goto L402
	}
L402:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1292
	v1294 = F_slice_del(m, l0)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	if v1294 < int32(0) {
		v1335 = v1294
		goto L398
	} else {
		goto L404
	}
L404:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1298
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1298-int32(3) <= v1302 {
		v1338 = int32(0)
		goto L397
	} else {
		goto L405
	}
L405:
	;
	v1306 = int32(0)
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1307+v1298-int32(1)))))
	if base.B2i32(v1311 == int32(187))|base.B2i32(v1311 == int32(128)) == v1306 {
		v1338 = v1306
		goto L397
	} else {
		goto L406
	}
L406:
	;
	v1322 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_93), int32(8))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L5
	} else {
		goto L407
	}
L407:
	;
	if v1322 == int32(0) {
		v1335 = int32(0)
		goto L398
	} else {
		goto L408
	}
L408:
	;
	v1329 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_94))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L5
	} else {
		goto L409
	}
L409:
	;
	if int32(0) <= v1329 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1333 = int32(1)
	goto L412
L411:
	;
	v1333 = v1329
	goto L412
L412:
	;
	v1335 = v1333
	goto L398
L413:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1342 = v1341 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1342
	v1344 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1342
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1342-int32(9) <= v1347 {
		v1391 = v1344
		goto L414
	} else {
		goto L415
	}
L414:
	;
	if v1391 < int32(0) {
		v3379 = v1391
		goto L1
	} else {
		goto L427
	}
L415:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1351+v1342-int32(1)))))
	if base.B2i32(v1355 != int32(189))&base.B2i32(v1355 != int32(131)) != 0 {
		v1391 = v1344
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v1363 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_95), int32(2))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L5
	} else {
		goto L417
	}
L417:
	;
	if v1363 == int32(0) {
		v1391 = v1344
		goto L414
	} else {
		goto L418
	}
L418:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1367
	v1369 = F_slice_del(m, l0)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L5
	} else {
		goto L419
	}
L419:
	;
	if v1369 < int32(0) {
		v1391 = v1369
		goto L414
	} else {
		goto L420
	}
L420:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1373
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1373
	v1379 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_96), int32(15))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L5
	} else {
		goto L421
	}
L421:
	;
	if v1379 == int32(0) {
		v1391 = int32(0)
		goto L414
	} else {
		goto L422
	}
L422:
	;
	v1386 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_97))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L5
	} else {
		goto L423
	}
L423:
	;
	if int32(0) <= v1386 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1390 = int32(1)
	goto L426
L425:
	;
	v1390 = v1386
	goto L426
L426:
	;
	v1391 = v1390
	goto L414
L427:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1396 = v1395 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1396
	v1398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1396
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1396-int32(5) <= v1401 {
		v1452 = v1398
		goto L429
	} else {
		goto L430
	}
L428:
	;
	if v1456 < int32(0) {
		v3379 = v1456
		goto L1
	} else {
		goto L443
	}
L429:
	;
	v1456 = v1452
	goto L428
L430:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1405+v1396-int32(1)))))
	if base.B2i32(v1409 != int32(189))&base.B2i32(v1409 != int32(131)) != 0 {
		v1452 = v1398
		goto L429
	} else {
		goto L431
	}
L431:
	;
	v1417 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_98), int32(2))
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L5
	} else {
		goto L432
	}
L432:
	;
	if v1417 == int32(0) {
		v1452 = v1398
		goto L429
	} else {
		goto L433
	}
L433:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1421
	v1423 = F_slice_del(m, l0)
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L5
	} else {
		goto L434
	}
L434:
	;
	if v1423 < int32(0) {
		v1452 = v1423
		goto L429
	} else {
		goto L435
	}
L435:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1428 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1427))) = v1428
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1430
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1430
	v1436 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_99), int32(8))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L5
	} else {
		goto L436
	}
L436:
	;
	if v1436 == int32(0) {
		v1456 = v1428
		goto L428
	} else {
		goto L437
	}
L437:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1442 < v1441 {
		v1452 = int32(0)
		goto L429
	} else {
		goto L438
	}
L438:
	;
	v1447 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_100))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L5
	} else {
		goto L439
	}
L439:
	;
	if int32(0) <= v1447 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1451 = int32(1)
	goto L442
L441:
	;
	v1451 = v1447
	goto L442
L442:
	;
	v1452 = v1451
	goto L429
L443:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1460 = v1459 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1460
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1460
	v1467 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_101), int32(3))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L5
	} else {
		goto L445
	}
L444:
	;
	if v1625 < int32(0) {
		v3379 = v1625
		goto L1
	} else {
		goto L478
	}
L445:
	;
	if v1467 == int32(0) {
		v1625 = int32(0)
		goto L444
	} else {
		goto L446
	}
L446:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1471
	v1473 = F_slice_del(m, l0)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L5
	} else {
		goto L448
	}
L447:
	;
	v1625 = v1622
	goto L444
L448:
	;
	if v1473 < int32(0) {
		v1622 = v1473
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1478 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1477))) = v1478
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1480
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L452
L450:
	;
	if v1612 != 0 {
		v1622 = v1478
		goto L447
	} else {
		goto L473
	}
L451:
	;
	v1612 = v1605
	goto L450
L452:
	;
	if v1480 <= v1497 {
		v1605 = int32(-1)
		goto L451
	} else {
		goto L454
	}
L453:
	;
	v1605 = int32(0)
	goto L451
L454:
	;
	v1514 = int32(1)
	v1515 = v1480 - v1514
	v1517 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1498+v1515))))
	v1519 = v1517 & int32(255)
	if base.B2i32(v1515 == v1497)|base.B2i32(int32(0) <= v1517) != 0 {
		v1577 = v1519
		v1581 = v1514
		goto L455
	} else {
		goto L456
	}
L455:
	;
	if int32(969) < v1577 {
		goto L463
	} else {
		goto L464
	}
L456:
	;
	v1526 = v1519 & int32(63)
	v1528 = v1480 - int32(2)
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498+v1528))))
	v1532 = v1530 << (uint(int32(6)) % 32)
	if base.B2i32(v1528 != v1497)&base.B2i32(base.Ui32(v1530) < base.Ui32(int32(192))) == int32(0) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v1577 = v1532&int32(1984) | v1526
	v1581 = int32(2)
	goto L455
L458:
	;
	goto L459
L459:
	;
	v1545 = v1532&int32(4032) | v1526
	v1547 = v1480 - int32(3)
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498+v1547))))
	if base.B2i32(v1547 != v1497)&base.B2i32(base.Ui32(v1549) < base.Ui32(int32(224))) == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v1577 = v1549<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v1545
	v1581 = int32(3)
	goto L455
L461:
	;
	goto L462
L462:
	;
	v1567 = int32(4)
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1480+v1498-v1567))))
	v1577 = v1549<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v1569&int32(7)<<(uint(int32(18))%32) | v1545
	v1581 = v1567
	goto L455
L463:
	;
	v1612 = v1581
	goto L450
L464:
	;
	goto L465
L465:
	;
	v1583 = v1577 - int32(945)
	if v1583 < int32(0) {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v1612 = v1581
	goto L450
L467:
	;
	goto L468
L468:
	;
	v1589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1583)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[0]))))
	if int32(base.Ui32(v1589)>>(uint(v1583&int32(7))%32))&int32(1) == int32(0) {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v1612 = v1581
	goto L450
L470:
	;
	goto L471
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1480 - v1581
	goto L472
L472:
	;
	goto L453
L473:
	;
	v1616 = F_slice_from_s(m, l0, int32(2), int32(_a_F_greek_UTF_8_stem_104))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L5
	} else {
		goto L474
	}
L474:
	;
	if int32(0) <= v1616 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v1620 = int32(1)
	goto L477
L476:
	;
	v1620 = v1616
	goto L477
L477:
	;
	v1622 = v1620
	goto L447
L478:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1629 = v1628 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1629
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1629
	v1636 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_105), int32(4))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L5
	} else {
		goto L480
	}
L479:
	;
	if v1820 < int32(0) {
		v3379 = v1820
		goto L1
	} else {
		goto L521
	}
L480:
	;
	if v1636 == int32(0) {
		v1820 = int32(0)
		goto L479
	} else {
		goto L481
	}
L481:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1640
	v1642 = F_slice_del(m, l0)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L5
	} else {
		goto L483
	}
L482:
	;
	v1820 = v1816
	goto L479
L483:
	;
	if v1642 < int32(0) {
		v1816 = v1642
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1646))) = int32(0)
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1649
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1649
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L488
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1796
	v1798 = int32(0)
	v1801 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_106), int32(36))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L5
	} else {
		goto L514
	}
L486:
	;
	if v1781 == int32(0) {
		goto L509
	} else {
		goto L510
	}
L487:
	;
	v1781 = v1774
	goto L486
L488:
	;
	if v1649 <= v1666 {
		v1774 = int32(-1)
		goto L487
	} else {
		goto L490
	}
L489:
	;
	v1774 = int32(0)
	goto L487
L490:
	;
	v1683 = int32(1)
	v1684 = v1649 - v1683
	v1686 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1667+v1684))))
	v1688 = v1686 & int32(255)
	if base.B2i32(v1684 == v1666)|base.B2i32(int32(0) <= v1686) != 0 {
		v1746 = v1688
		v1750 = v1683
		goto L491
	} else {
		goto L492
	}
L491:
	;
	if int32(969) < v1746 {
		goto L499
	} else {
		goto L500
	}
L492:
	;
	v1695 = v1688 & int32(63)
	v1697 = v1649 - int32(2)
	v1699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1667+v1697))))
	v1701 = v1699 << (uint(int32(6)) % 32)
	if base.B2i32(v1697 != v1666)&base.B2i32(base.Ui32(v1699) < base.Ui32(int32(192))) == int32(0) {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v1746 = v1701&int32(1984) | v1695
	v1750 = int32(2)
	goto L491
L494:
	;
	goto L495
L495:
	;
	v1714 = v1701&int32(4032) | v1695
	v1716 = v1649 - int32(3)
	v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1667+v1716))))
	if base.B2i32(v1716 != v1666)&base.B2i32(base.Ui32(v1718) < base.Ui32(int32(224))) == int32(0) {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v1746 = v1718<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v1714
	v1750 = int32(3)
	goto L491
L497:
	;
	goto L498
L498:
	;
	v1736 = int32(4)
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1649+v1667-v1736))))
	v1746 = v1718<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v1738&int32(7)<<(uint(int32(18))%32) | v1714
	v1750 = v1736
	goto L491
L499:
	;
	v1781 = v1750
	goto L486
L500:
	;
	goto L501
L501:
	;
	v1752 = v1746 - int32(945)
	if v1752 < int32(0) {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v1781 = v1750
	goto L486
L503:
	;
	goto L504
L504:
	;
	v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1752)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[0]))))
	if int32(base.Ui32(v1758)>>(uint(v1752&int32(7))%32))&int32(1) == int32(0) {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v1781 = v1750
	goto L486
L506:
	;
	goto L507
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1649 - v1750
	goto L508
L508:
	;
	goto L489
L509:
	;
	v1786 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_107))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L5
	} else {
		goto L512
	}
L510:
	;
	goto L511
L511:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1793 = v1791 + (v1649 - v1652)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1793
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1793
	v1796 = v1793
	goto L485
L512:
	;
	if v1786 < int32(0) {
		v1816 = v1786
		goto L482
	} else {
		goto L513
	}
L513:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1796 = v1790
	goto L485
L514:
	;
	if v1801 == int32(0) {
		v1816 = v1798
		goto L482
	} else {
		goto L515
	}
L515:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1806 < v1805 {
		v1816 = v1798
		goto L482
	} else {
		goto L516
	}
L516:
	;
	v1811 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_108))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L5
	} else {
		goto L517
	}
L517:
	;
	if int32(0) <= v1811 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v1815 = int32(1)
	goto L520
L519:
	;
	v1815 = v1811
	goto L520
L520:
	;
	v1816 = v1815
	goto L482
L521:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1824 = v1823 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1824
	v1829 = int32(10)
	v1831 = int32(0)
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1824-v1834 < v1829 {
		v1844 = v1831
		goto L526
	} else {
		goto L527
	}
L522:
	;
	if v1950 < int32(0) {
		v3379 = v1950
		goto L1
	} else {
		goto L554
	}
L523:
	;
	v1950 = v1946
	goto L522
L524:
	;
	v1858 = v1824 - v1823
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1860 = v1858 + v1859
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1860
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1860
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1860-int32(9) <= v1863 {
		goto L533
	} else {
		goto L534
	}
L525:
	;
	if v1844 == int32(0) {
		goto L524
	} else {
		goto L529
	}
L526:
	;
	goto L525
L527:
	;
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1840 = F_memcmp(m, v1837+v1824-v1829, int32(_a_F_greek_UTF_8_stem_109), v1829)
	mBase = m.M
	if v1840 != 0 {
		v1844 = v1831
		goto L526
	} else {
		goto L528
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1824 - v1829
	v1844 = int32(1)
	goto L526
L529:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1847
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1849 < v1847 {
		goto L524
	} else {
		goto L530
	}
L530:
	;
	v1853 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_110))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L5
	} else {
		goto L531
	}
L531:
	;
	if v1853 < int32(0) {
		v1946 = v1853
		goto L523
	} else {
		goto L532
	}
L532:
	;
	goto L524
L533:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1891 = v1890 + v1858
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1891
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1891
	v1894 = int32(0)
	v1895 = int32(6)
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1891-v1900 < v1895 {
		v1910 = v1894
		goto L541
	} else {
		goto L542
	}
L534:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1867+v1860-int32(1)))))
	if v1871 != int32(181) {
		goto L533
	} else {
		goto L535
	}
L535:
	;
	v1876 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_111), int32(5))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L5
	} else {
		goto L536
	}
L536:
	;
	if v1876 == int32(0) {
		goto L533
	} else {
		goto L537
	}
L537:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1880
	v1882 = F_slice_del(m, l0)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L5
	} else {
		goto L538
	}
L538:
	;
	if v1882 < int32(0) {
		v1946 = v1882
		goto L523
	} else {
		goto L539
	}
L539:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1886))) = int32(0)
	goto L533
L540:
	;
	if v1910 == int32(0) {
		v1950 = v1894
		goto L522
	} else {
		goto L544
	}
L541:
	;
	goto L540
L542:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1906 = F_memcmp(m, v1903+v1891-v1895, int32(_a_F_greek_UTF_8_stem_112), v1895)
	mBase = m.M
	if v1906 != 0 {
		v1910 = v1894
		goto L541
	} else {
		goto L543
	}
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1891 - v1895
	v1910 = int32(1)
	goto L541
L544:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1913
	v1915 = F_slice_del(m, l0)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L5
	} else {
		goto L545
	}
L545:
	;
	if v1915 < int32(0) {
		v1946 = v1915
		goto L523
	} else {
		goto L546
	}
L546:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1920 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1919))) = v1920
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1922
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1922
	v1928 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_113), int32(12))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L5
	} else {
		goto L547
	}
L547:
	;
	if v1928 == int32(0) {
		v1950 = v1920
		goto L522
	} else {
		goto L548
	}
L548:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1934 < v1933 {
		v1946 = int32(0)
		goto L523
	} else {
		goto L549
	}
L549:
	;
	v1939 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_114))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L5
	} else {
		goto L550
	}
L550:
	;
	if int32(0) <= v1939 {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	v1943 = int32(1)
	goto L553
L552:
	;
	v1943 = v1939
	goto L553
L553:
	;
	v1946 = v1943
	goto L523
L554:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1954 = v1953 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1954
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1954
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1954-int32(9) <= v1959 {
		goto L557
	} else {
		goto L558
	}
L555:
	;
	if v2223 < int32(0) {
		v3379 = v2223
		goto L1
	} else {
		goto L614
	}
L556:
	;
	v2223 = v2219
	goto L555
L557:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2018 = v2016 + (v1954 - v1953)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2018
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2018
	v2021 = int32(0)
	v2022 = int32(6)
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2018-v2027 < v2022 {
		v2037 = v2021
		goto L572
	} else {
		goto L573
	}
L558:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1963+v1954-int32(1)))))
	if v1967 != int32(181) {
		goto L557
	} else {
		goto L559
	}
L559:
	;
	v1972 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_115), int32(11))
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L5
	} else {
		goto L560
	}
L560:
	;
	if v1972 == int32(0) {
		goto L557
	} else {
		goto L561
	}
L561:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1976
	v1978 = F_slice_del(m, l0)
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L5
	} else {
		goto L562
	}
L562:
	;
	if v1978 < int32(0) {
		v2219 = v1978
		goto L556
	} else {
		goto L563
	}
L563:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1982))) = int32(0)
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1985
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1985
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1985-int32(3) <= v1988 {
		goto L557
	} else {
		goto L564
	}
L564:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1992+v1985-int32(1)))))
	switch v1996 - int32(129) {
	case 0, 2:
		goto L565
	default:
		goto L557
	}
L565:
	;
	v2001 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_116), int32(2))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L5
	} else {
		goto L566
	}
L566:
	;
	if v2001 == int32(0) {
		goto L557
	} else {
		goto L567
	}
L567:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2006 < v2005 {
		goto L557
	} else {
		goto L568
	}
L568:
	;
	v2010 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_117))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L5
	} else {
		goto L569
	}
L569:
	;
	if v2010 < int32(0) {
		v2219 = v2010
		goto L556
	} else {
		goto L570
	}
L570:
	;
	goto L557
L571:
	;
	if v2037 == int32(0) {
		v2223 = v2021
		goto L555
	} else {
		goto L575
	}
L572:
	;
	goto L571
L573:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2033 = F_memcmp(m, v2030+v2018-v2022, int32(_a_F_greek_UTF_8_stem_118), v2022)
	mBase = m.M
	if v2033 != 0 {
		v2037 = v2021
		goto L572
	} else {
		goto L574
	}
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2018 - v2022
	v2037 = int32(1)
	goto L572
L575:
	;
	v2040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2040
	v2042 = F_slice_del(m, l0)
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L5
	} else {
		goto L576
	}
L576:
	;
	if v2042 < int32(0) {
		v2219 = v2042
		goto L556
	} else {
		goto L577
	}
L577:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2046))) = int32(0)
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2049
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2049
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L581
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2196
	v2199 = int32(0)
	v2202 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_119), int32(95))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L5
	} else {
		goto L607
	}
L579:
	;
	if v2181 == int32(0) {
		goto L602
	} else {
		goto L603
	}
L580:
	;
	v2181 = v2174
	goto L579
L581:
	;
	if v2049 <= v2066 {
		v2174 = int32(-1)
		goto L580
	} else {
		goto L583
	}
L582:
	;
	v2174 = int32(0)
	goto L580
L583:
	;
	v2083 = int32(1)
	v2084 = v2049 - v2083
	v2086 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2067+v2084))))
	v2088 = v2086 & int32(255)
	if base.B2i32(v2084 == v2066)|base.B2i32(int32(0) <= v2086) != 0 {
		v2146 = v2088
		v2150 = v2083
		goto L584
	} else {
		goto L585
	}
L584:
	;
	if int32(969) < v2146 {
		goto L592
	} else {
		goto L593
	}
L585:
	;
	v2095 = v2088 & int32(63)
	v2097 = v2049 - int32(2)
	v2099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2067+v2097))))
	v2101 = v2099 << (uint(int32(6)) % 32)
	if base.B2i32(v2097 != v2066)&base.B2i32(base.Ui32(v2099) < base.Ui32(int32(192))) == int32(0) {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	v2146 = v2101&int32(1984) | v2095
	v2150 = int32(2)
	goto L584
L587:
	;
	goto L588
L588:
	;
	v2114 = v2101&int32(4032) | v2095
	v2116 = v2049 - int32(3)
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2067+v2116))))
	if base.B2i32(v2116 != v2066)&base.B2i32(base.Ui32(v2118) < base.Ui32(int32(224))) == int32(0) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v2146 = v2118<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v2114
	v2150 = int32(3)
	goto L584
L590:
	;
	goto L591
L591:
	;
	v2136 = int32(4)
	v2138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2049+v2067-v2136))))
	v2146 = v2118<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v2138&int32(7)<<(uint(int32(18))%32) | v2114
	v2150 = v2136
	goto L584
L592:
	;
	v2181 = v2150
	goto L579
L593:
	;
	goto L594
L594:
	;
	v2152 = v2146 - int32(945)
	if v2152 < int32(0) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v2181 = v2150
	goto L579
L596:
	;
	goto L597
L597:
	;
	v2158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2152)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[1]))))
	if int32(base.Ui32(v2158)>>(uint(v2152&int32(7))%32))&int32(1) == int32(0) {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v2181 = v2150
	goto L579
L599:
	;
	goto L600
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2049 - v2150
	goto L601
L601:
	;
	goto L582
L602:
	;
	v2186 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_120))
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L5
	} else {
		goto L605
	}
L603:
	;
	goto L604
L604:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2193 = v2191 + (v2049 - v2052)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2193
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2193
	v2196 = v2193
	goto L578
L605:
	;
	if v2186 < int32(0) {
		v2219 = v2186
		goto L556
	} else {
		goto L606
	}
L606:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2196 = v2190
	goto L578
L607:
	;
	if v2202 == int32(0) {
		v2219 = v2199
		goto L556
	} else {
		goto L608
	}
L608:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2207 < v2206 {
		v2219 = v2199
		goto L556
	} else {
		goto L609
	}
L609:
	;
	v2212 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_121))
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L5
	} else {
		goto L610
	}
L610:
	;
	if int32(0) <= v2212 {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v2216 = int32(1)
	goto L613
L612:
	;
	v2216 = v2212
	goto L613
L613:
	;
	v2219 = v2216
	goto L556
L614:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2227 = v2226 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2227
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2227
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2227-int32(9) <= v2232 {
		goto L617
	} else {
		goto L618
	}
L615:
	;
	if v2482 < int32(0) {
		v3379 = v2482
		goto L1
	} else {
		goto L673
	}
L616:
	;
	v2482 = v2478
	goto L615
L617:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2261 = v2259 + (v2227 - v2226)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2261
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2261
	v2264 = int32(0)
	v2265 = int32(6)
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2261-v2270 < v2265 {
		v2280 = v2264
		goto L625
	} else {
		goto L626
	}
L618:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236+v2227-int32(1)))))
	if v2240 != int32(181) {
		goto L617
	} else {
		goto L619
	}
L619:
	;
	v2245 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_122), int32(1))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L5
	} else {
		goto L620
	}
L620:
	;
	if v2245 == int32(0) {
		goto L617
	} else {
		goto L621
	}
L621:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2249
	v2251 = F_slice_del(m, l0)
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L5
	} else {
		goto L622
	}
L622:
	;
	if v2251 < int32(0) {
		v2478 = v2251
		goto L616
	} else {
		goto L623
	}
L623:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2255))) = int32(0)
	goto L617
L624:
	;
	if v2280 == int32(0) {
		v2482 = v2264
		goto L615
	} else {
		goto L628
	}
L625:
	;
	goto L624
L626:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2276 = F_memcmp(m, v2273+v2261-v2265, int32(_a_F_greek_UTF_8_stem_123), v2265)
	mBase = m.M
	if v2276 != 0 {
		v2280 = v2264
		goto L625
	} else {
		goto L627
	}
L627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2261 - v2265
	v2280 = int32(1)
	goto L625
L628:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2283
	v2285 = F_slice_del(m, l0)
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L5
	} else {
		goto L629
	}
L629:
	;
	if v2285 < int32(0) {
		v2478 = v2285
		goto L616
	} else {
		goto L630
	}
L630:
	;
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2289))) = int32(0)
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2292
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L634
L631:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2456
	v2458 = int32(0)
	v2461 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_124), int32(25))
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L5
	} else {
		goto L666
	}
L632:
	;
	if v2424 == int32(0) {
		goto L655
	} else {
		goto L656
	}
L633:
	;
	v2424 = v2417
	goto L632
L634:
	;
	if v2292 <= v2309 {
		v2417 = int32(-1)
		goto L633
	} else {
		goto L636
	}
L635:
	;
	v2417 = int32(0)
	goto L633
L636:
	;
	v2326 = int32(1)
	v2327 = v2292 - v2326
	v2329 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2310+v2327))))
	v2331 = v2329 & int32(255)
	if base.B2i32(v2327 == v2309)|base.B2i32(int32(0) <= v2329) != 0 {
		v2389 = v2331
		v2393 = v2326
		goto L637
	} else {
		goto L638
	}
L637:
	;
	if int32(969) < v2389 {
		goto L645
	} else {
		goto L646
	}
L638:
	;
	v2338 = v2331 & int32(63)
	v2340 = v2292 - int32(2)
	v2342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2310+v2340))))
	v2344 = v2342 << (uint(int32(6)) % 32)
	if base.B2i32(v2340 != v2309)&base.B2i32(base.Ui32(v2342) < base.Ui32(int32(192))) == int32(0) {
		goto L639
	} else {
		goto L640
	}
L639:
	;
	v2389 = v2344&int32(1984) | v2338
	v2393 = int32(2)
	goto L637
L640:
	;
	goto L641
L641:
	;
	v2357 = v2344&int32(4032) | v2338
	v2359 = v2292 - int32(3)
	v2361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2310+v2359))))
	if base.B2i32(v2359 != v2309)&base.B2i32(base.Ui32(v2361) < base.Ui32(int32(224))) == int32(0) {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	v2389 = v2361<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_102) | v2357
	v2393 = int32(3)
	goto L637
L643:
	;
	goto L644
L644:
	;
	v2379 = int32(4)
	v2381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2292+v2310-v2379))))
	v2389 = v2361<<(uint(int32(12))%32)&int32(_a_F_greek_UTF_8_stem_103) | v2381&int32(7)<<(uint(int32(18))%32) | v2357
	v2393 = v2379
	goto L637
L645:
	;
	v2424 = v2393
	goto L632
L646:
	;
	goto L647
L647:
	;
	v2395 = v2389 - int32(945)
	if v2395 < int32(0) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	v2424 = v2393
	goto L632
L649:
	;
	goto L650
L650:
	;
	v2401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2395)>>(uint(int32(3))%32)))+uint32(_c_F_greek_UTF_8_stem[1]))))
	if int32(base.Ui32(v2401)>>(uint(v2395&int32(7))%32))&int32(1) == int32(0) {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	v2424 = v2393
	goto L632
L652:
	;
	goto L653
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2292 - v2393
	goto L654
L654:
	;
	goto L635
L655:
	;
	v2429 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_125))
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L5
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2434 = v2295 - v2292
	v2435 = v2433 - v2434
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2435
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2435
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2435
	v2441 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_126), int32(31))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L5
	} else {
		goto L660
	}
L658:
	;
	if int32(0) <= v2429 {
		goto L631
	} else {
		goto L659
	}
L659:
	;
	v2478 = v2429
	goto L616
L660:
	;
	if v2441 != 0 {
		goto L661
	} else {
		goto L662
	}
L661:
	;
	v2445 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_127))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L5
	} else {
		goto L664
	}
L662:
	;
	goto L663
L663:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2450 = v2449 - v2434
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2450
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2450
	goto L631
L664:
	;
	if int32(0) <= v2445 {
		goto L631
	} else {
		goto L665
	}
L665:
	;
	v2478 = v2445
	goto L616
L666:
	;
	if v2461 == int32(0) {
		v2478 = v2458
		goto L616
	} else {
		goto L667
	}
L667:
	;
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2466 < v2465 {
		v2478 = v2458
		goto L616
	} else {
		goto L668
	}
L668:
	;
	v2471 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_128))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L5
	} else {
		goto L669
	}
L669:
	;
	if int32(0) <= v2471 {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v2475 = int32(1)
	goto L672
L671:
	;
	v2475 = v2471
	goto L672
L672:
	;
	v2478 = v2475
	goto L616
L673:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2486 = v2485 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2486
	v2488 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2486
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2486-int32(9) <= v2491 {
		v2585 = v2488
		goto L675
	} else {
		goto L676
	}
L674:
	;
	if v2589 < int32(0) {
		v3379 = v2589
		goto L1
	} else {
		goto L702
	}
L675:
	;
	v2589 = v2585
	goto L674
L676:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2495+v2486-int32(1)))))
	if v2499 != int32(131) {
		v2585 = v2488
		goto L675
	} else {
		goto L677
	}
L677:
	;
	v2504 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_129), int32(2))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L5
	} else {
		goto L678
	}
L678:
	;
	if v2504 == int32(0) {
		v2585 = v2488
		goto L675
	} else {
		goto L679
	}
L679:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2508
	v2510 = F_slice_del(m, l0)
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L5
	} else {
		goto L680
	}
L680:
	;
	if v2510 < int32(0) {
		v2585 = v2510
		goto L675
	} else {
		goto L681
	}
L681:
	;
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2515 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2514))) = v2515
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2517
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2517
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2521 = int32(6)
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2517-v2526 < v2521 {
		v2536 = v2515
		goto L684
	} else {
		goto L685
	}
L682:
	;
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2552 = v2550 + (v2517 - v2520)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2552
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2552
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2552
	v2556 = int32(0)
	v2557 = int32(6)
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2552-v2562 < v2557 {
		v2572 = v2556
		goto L694
	} else {
		goto L695
	}
L683:
	;
	if v2536 == int32(0) {
		goto L682
	} else {
		goto L687
	}
L684:
	;
	goto L683
L685:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2532 = F_memcmp(m, v2529+v2517-v2521, int32(_a_F_greek_UTF_8_stem_130), v2521)
	mBase = m.M
	if v2532 != 0 {
		v2536 = v2515
		goto L684
	} else {
		goto L686
	}
L686:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2517 - v2521
	v2536 = int32(1)
	goto L684
L687:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2540 < v2539 {
		goto L682
	} else {
		goto L688
	}
L688:
	;
	v2545 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_131))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L5
	} else {
		goto L689
	}
L689:
	;
	if int32(0) <= v2545 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v2549 = int32(1)
	goto L692
L691:
	;
	v2549 = v2545
	goto L692
L692:
	;
	v2589 = v2549
	goto L674
L693:
	;
	if v2572 == int32(0) {
		v2589 = v2556
		goto L674
	} else {
		goto L697
	}
L694:
	;
	goto L693
L695:
	;
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2568 = F_memcmp(m, v2565+v2552-v2557, int32(_a_F_greek_UTF_8_stem_132), v2557)
	mBase = m.M
	if v2568 != 0 {
		v2572 = v2556
		goto L694
	} else {
		goto L696
	}
L696:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2552 - v2557
	v2572 = int32(1)
	goto L694
L697:
	;
	v2578 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_133))
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L5
	} else {
		goto L698
	}
L698:
	;
	if int32(0) <= v2578 {
		goto L699
	} else {
		goto L700
	}
L699:
	;
	v2582 = int32(1)
	goto L701
L700:
	;
	v2582 = v2578
	goto L701
L701:
	;
	v2585 = v2582
	goto L675
L702:
	;
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2593 = v2592 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2593
	v2595 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2593
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2593-int32(11) <= v2598 {
		v2659 = v2595
		goto L704
	} else {
		goto L705
	}
L703:
	;
	if v2662 < int32(0) {
		v3379 = v2662
		goto L1
	} else {
		goto L721
	}
L704:
	;
	v2662 = v2659
	goto L703
L705:
	;
	v2602 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2602+v2593-int32(1)))))
	if v2606 != int32(181) {
		v2659 = v2595
		goto L704
	} else {
		goto L706
	}
L706:
	;
	v2611 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_134), int32(2))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L5
	} else {
		goto L707
	}
L707:
	;
	if v2611 == int32(0) {
		v2659 = v2595
		goto L704
	} else {
		goto L708
	}
L708:
	;
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2615
	v2617 = F_slice_del(m, l0)
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L5
	} else {
		goto L709
	}
L709:
	;
	if v2617 < int32(0) {
		v2659 = v2617
		goto L704
	} else {
		goto L710
	}
L710:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2622 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2621))) = v2622
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2624
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2624
	v2628 = int32(4)
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2624-v2633 < v2628 {
		v2643 = v2622
		goto L712
	} else {
		goto L713
	}
L711:
	;
	if v2643 == int32(0) {
		v2662 = v2622
		goto L703
	} else {
		goto L715
	}
L712:
	;
	goto L711
L713:
	;
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2639 = F_memcmp(m, v2636+v2624-v2628, int32(_a_F_greek_UTF_8_stem_135), v2628)
	mBase = m.M
	if v2639 != 0 {
		v2643 = v2622
		goto L712
	} else {
		goto L714
	}
L714:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2624 - v2628
	v2643 = int32(1)
	goto L712
L715:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2648 < v2647 {
		v2659 = int32(0)
		goto L704
	} else {
		goto L716
	}
L716:
	;
	v2653 = F_slice_from_s(m, l0, int32(10), int32(_a_F_greek_UTF_8_stem_136))
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L5
	} else {
		goto L717
	}
L717:
	;
	if int32(0) <= v2653 {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	v2657 = int32(1)
	goto L720
L719:
	;
	v2657 = v2653
	goto L720
L720:
	;
	v2659 = v2657
	goto L704
L721:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2666 = v2665 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2666
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2666
	v2671 = int32(10)
	v2673 = int32(0)
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2666-v2676 < v2671 {
		v2686 = v2673
		goto L726
	} else {
		goto L727
	}
L722:
	;
	if v2783 < int32(0) {
		v3379 = v2783
		goto L1
	} else {
		goto L753
	}
L723:
	;
	v2783 = v2781
	goto L722
L724:
	;
	v2726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2728 = v2726 + (v2666 - v2665)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2728
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2728
	v2731 = int32(0)
	v2732 = int32(8)
	v2737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2728-v2737 < v2732 {
		v2747 = v2731
		goto L740
	} else {
		goto L741
	}
L725:
	;
	if v2686 == int32(0) {
		goto L724
	} else {
		goto L729
	}
L726:
	;
	goto L725
L727:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2682 = F_memcmp(m, v2679+v2666-v2671, int32(_a_F_greek_UTF_8_stem_137), v2671)
	mBase = m.M
	if v2682 != 0 {
		v2686 = v2673
		goto L726
	} else {
		goto L728
	}
L728:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2666 - v2671
	v2686 = int32(1)
	goto L726
L729:
	;
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2689
	v2691 = F_slice_del(m, l0)
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L5
	} else {
		goto L730
	}
L730:
	;
	if v2691 < int32(0) {
		v2781 = v2691
		goto L723
	} else {
		goto L731
	}
L731:
	;
	v2695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2695))) = int32(0)
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2698
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2698
	v2702 = v2698 - int32(1)
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2702 <= v2703 {
		goto L724
	} else {
		goto L732
	}
L732:
	;
	v2705 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2705+v2702))))
	switch v2707 - int32(128) {
	case 0, 6:
		goto L733
	default:
		goto L724
	}
L733:
	;
	v2712 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_138), int32(6))
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L5
	} else {
		goto L734
	}
L734:
	;
	if v2712 == int32(0) {
		goto L724
	} else {
		goto L735
	}
L735:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2717 < v2716 {
		goto L724
	} else {
		goto L736
	}
L736:
	;
	v2721 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_139))
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L5
	} else {
		goto L737
	}
L737:
	;
	if v2721 < int32(0) {
		v2781 = v2721
		goto L723
	} else {
		goto L738
	}
L738:
	;
	goto L724
L739:
	;
	if v2747 == int32(0) {
		v2783 = v2731
		goto L722
	} else {
		goto L743
	}
L740:
	;
	goto L739
L741:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2743 = F_memcmp(m, v2740+v2728-v2732, int32(_a_F_greek_UTF_8_stem_140), v2732)
	mBase = m.M
	if v2743 != 0 {
		v2747 = v2731
		goto L740
	} else {
		goto L742
	}
L742:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2728 - v2732
	v2747 = int32(1)
	goto L740
L743:
	;
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2750
	v2752 = F_slice_del(m, l0)
	mBase = m.M
	v2753 = m.ExcPending
	if v2753 != 0 {
		goto L5
	} else {
		goto L744
	}
L744:
	;
	if v2752 < int32(0) {
		v2781 = v2752
		goto L723
	} else {
		goto L745
	}
L745:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2757 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2756))) = v2757
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2759
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2759
	v2765 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_141), int32(9))
	mBase = m.M
	v2766 = m.ExcPending
	if v2766 != 0 {
		goto L5
	} else {
		goto L746
	}
L746:
	;
	if v2765 == int32(0) {
		v2783 = v2757
		goto L722
	} else {
		goto L747
	}
L747:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2771 < v2770 {
		v2781 = int32(0)
		goto L723
	} else {
		goto L748
	}
L748:
	;
	v2776 = F_slice_from_s(m, l0, int32(8), int32(_a_F_greek_UTF_8_stem_142))
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L5
	} else {
		goto L749
	}
L749:
	;
	if int32(0) <= v2776 {
		goto L750
	} else {
		goto L751
	}
L750:
	;
	v2780 = int32(1)
	goto L752
L751:
	;
	v2780 = v2776
	goto L752
L752:
	;
	v2781 = v2780
	goto L723
L753:
	;
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2787 = v2786 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2787
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2787
	v2794 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_143), int32(3))
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L5
	} else {
		goto L756
	}
L754:
	;
	if v2882 < int32(0) {
		v3379 = v2882
		goto L1
	} else {
		goto L783
	}
L755:
	;
	v2882 = v2878
	goto L754
L756:
	;
	if v2794 != 0 {
		goto L757
	} else {
		goto L758
	}
L757:
	;
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2796
	v2798 = F_slice_del(m, l0)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L5
	} else {
		goto L760
	}
L758:
	;
	goto L759
L759:
	;
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2808 = v2806 + (v2787 - v2786)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2808
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2808
	v2814 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_144), int32(3))
	mBase = m.M
	v2815 = m.ExcPending
	if v2815 != 0 {
		goto L5
	} else {
		goto L762
	}
L760:
	;
	if v2798 < int32(0) {
		v2878 = v2798
		goto L755
	} else {
		goto L761
	}
L761:
	;
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2802))) = int32(0)
	goto L759
L762:
	;
	if v2814 == int32(0) {
		v2882 = int32(0)
		goto L754
	} else {
		goto L763
	}
L763:
	;
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2818
	v2820 = F_slice_del(m, l0)
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L5
	} else {
		goto L764
	}
L764:
	;
	if v2820 < int32(0) {
		v2878 = v2820
		goto L755
	} else {
		goto L765
	}
L765:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2824))) = int32(0)
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2827
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2827
	v2830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2833 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_145), int32(6))
	mBase = m.M
	v2834 = m.ExcPending
	if v2834 != 0 {
		goto L5
	} else {
		goto L766
	}
L766:
	;
	if v2833 != 0 {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	v2838 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_146))
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L5
	} else {
		goto L770
	}
L768:
	;
	goto L769
L769:
	;
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2845 = v2843 + (v2827 - v2830)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2845
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2845
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2845
	v2849 = int32(0)
	v2851 = v2845 - int32(1)
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2851 <= v2852 {
		v2878 = v2849
		goto L755
	} else {
		goto L774
	}
L770:
	;
	if int32(0) <= v2838 {
		goto L771
	} else {
		goto L772
	}
L771:
	;
	v2842 = int32(1)
	goto L773
L772:
	;
	v2842 = v2838
	goto L773
L773:
	;
	v2882 = v2842
	goto L754
L774:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2854+v2851))))
	if v2856 != int32(184) {
		v2878 = v2849
		goto L755
	} else {
		goto L775
	}
L775:
	;
	v2861 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_147), int32(5))
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L5
	} else {
		goto L776
	}
L776:
	;
	if v2861 == int32(0) {
		v2878 = v2849
		goto L755
	} else {
		goto L777
	}
L777:
	;
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2866 < v2865 {
		v2878 = v2849
		goto L755
	} else {
		goto L778
	}
L778:
	;
	v2871 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_148))
	mBase = m.M
	v2872 = m.ExcPending
	if v2872 != 0 {
		goto L5
	} else {
		goto L779
	}
L779:
	;
	if int32(0) <= v2871 {
		goto L780
	} else {
		goto L781
	}
L780:
	;
	v2875 = int32(1)
	goto L782
L781:
	;
	v2875 = v2871
	goto L782
L782:
	;
	v2878 = v2875
	goto L755
L783:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2886 = v2885 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2886
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2886
	v2893 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_149), int32(3))
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
		goto L5
	} else {
		goto L785
	}
L784:
	;
	if v2950 < int32(0) {
		v3379 = v2950
		goto L1
	} else {
		goto L805
	}
L785:
	;
	if v2893 == int32(0) {
		v2950 = int32(0)
		goto L784
	} else {
		goto L786
	}
L786:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2897
	v2899 = F_slice_del(m, l0)
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L5
	} else {
		goto L788
	}
L787:
	;
	v2950 = v2946
	goto L784
L788:
	;
	if v2899 < int32(0) {
		v2946 = v2899
		goto L787
	} else {
		goto L789
	}
L789:
	;
	v2903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2903))) = int32(0)
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2906
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2906
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2912 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_150), int32(12))
	mBase = m.M
	v2913 = m.ExcPending
	if v2913 != 0 {
		goto L5
	} else {
		goto L790
	}
L790:
	;
	if v2912 != 0 {
		goto L791
	} else {
		goto L792
	}
L791:
	;
	v2917 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_151))
	mBase = m.M
	v2918 = m.ExcPending
	if v2918 != 0 {
		goto L5
	} else {
		goto L794
	}
L792:
	;
	goto L793
L793:
	;
	v2922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2924 = v2922 + (v2906 - v2909)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2924
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2924
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2924
	v2928 = int32(0)
	v2931 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_152), int32(25))
	mBase = m.M
	v2932 = m.ExcPending
	if v2932 != 0 {
		goto L5
	} else {
		goto L798
	}
L794:
	;
	if int32(0) <= v2917 {
		goto L795
	} else {
		goto L796
	}
L795:
	;
	v2921 = int32(1)
	goto L797
L796:
	;
	v2921 = v2917
	goto L797
L797:
	;
	v2950 = v2921
	goto L784
L798:
	;
	if v2931 == int32(0) {
		v2946 = v2928
		goto L787
	} else {
		goto L799
	}
L799:
	;
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2936 < v2935 {
		v2946 = v2928
		goto L787
	} else {
		goto L800
	}
L800:
	;
	v2941 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_153))
	mBase = m.M
	v2942 = m.ExcPending
	if v2942 != 0 {
		goto L5
	} else {
		goto L801
	}
L801:
	;
	if int32(0) <= v2941 {
		goto L802
	} else {
		goto L803
	}
L802:
	;
	v2945 = int32(1)
	goto L804
L803:
	;
	v2945 = v2941
	goto L804
L804:
	;
	v2946 = v2945
	goto L787
L805:
	;
	v2953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2954 = v2953 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2954
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2954
	v2961 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_154), int32(3))
	mBase = m.M
	v2962 = m.ExcPending
	if v2962 != 0 {
		goto L5
	} else {
		goto L807
	}
L806:
	;
	if v3011 < int32(0) {
		v3379 = v3011
		goto L1
	} else {
		goto L821
	}
L807:
	;
	if v2961 == int32(0) {
		v3011 = int32(0)
		goto L806
	} else {
		goto L808
	}
L808:
	;
	v2965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2965
	v2967 = F_slice_del(m, l0)
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L5
	} else {
		goto L810
	}
L809:
	;
	v3011 = v3008
	goto L806
L810:
	;
	if v2967 < int32(0) {
		v3008 = v2967
		goto L809
	} else {
		goto L811
	}
L811:
	;
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2972 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2971))) = v2972
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2974
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2974
	v2979 = v2974 - int32(1)
	v2980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2979 <= v2980 {
		v3011 = v2972
		goto L806
	} else {
		goto L812
	}
L812:
	;
	v2983 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2983+v2979))))
	if v2985 != int32(189) {
		v3011 = int32(0)
		goto L806
	} else {
		goto L813
	}
L813:
	;
	v2991 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_155), int32(6))
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L5
	} else {
		goto L814
	}
L814:
	;
	if v2991 == int32(0) {
		v3011 = int32(0)
		goto L806
	} else {
		goto L815
	}
L815:
	;
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2997 < v2996 {
		v3008 = int32(0)
		goto L809
	} else {
		goto L816
	}
L816:
	;
	v3002 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_156))
	mBase = m.M
	v3003 = m.ExcPending
	if v3003 != 0 {
		goto L5
	} else {
		goto L817
	}
L817:
	;
	if int32(0) <= v3002 {
		goto L818
	} else {
		goto L819
	}
L818:
	;
	v3006 = int32(1)
	goto L820
L819:
	;
	v3006 = v3002
	goto L820
L820:
	;
	v3008 = v3006
	goto L809
L821:
	;
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3015 = v3014 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3015
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3015
	v3022 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_157), int32(3))
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L5
	} else {
		goto L823
	}
L822:
	;
	if v3108 < int32(0) {
		v3379 = v3108
		goto L1
	} else {
		goto L850
	}
L823:
	;
	if v3022 == int32(0) {
		v3108 = int32(0)
		goto L822
	} else {
		goto L824
	}
L824:
	;
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3026
	v3028 = F_slice_del(m, l0)
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L5
	} else {
		goto L826
	}
L825:
	;
	v3108 = v3104
	goto L822
L826:
	;
	if v3028 < int32(0) {
		v3104 = v3028
		goto L825
	} else {
		goto L827
	}
L827:
	;
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3033 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3032))) = v3033
	v3035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3035
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3035
	v3038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3039 = int32(8)
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3035-v3044 < v3039 {
		v3054 = v3033
		goto L829
	} else {
		goto L830
	}
L828:
	;
	if v3054 != 0 {
		goto L832
	} else {
		goto L833
	}
L829:
	;
	goto L828
L830:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3050 = F_memcmp(m, v3047+v3035-v3039, int32(_a_F_greek_UTF_8_stem_158), v3039)
	mBase = m.M
	if v3050 != 0 {
		v3054 = v3033
		goto L829
	} else {
		goto L831
	}
L831:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3035 - v3039
	v3054 = int32(1)
	goto L829
L832:
	;
	v3058 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_159))
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L5
	} else {
		goto L835
	}
L833:
	;
	goto L834
L834:
	;
	v3063 = v3035 - v3038
	v3064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3065 = v3063 + v3064
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3065
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3065
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3065
	v3069 = int32(1)
	v3072 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_160), int32(12))
	mBase = m.M
	v3073 = m.ExcPending
	if v3073 != 0 {
		goto L5
	} else {
		goto L842
	}
L835:
	;
	if int32(0) <= v3058 {
		goto L836
	} else {
		goto L837
	}
L836:
	;
	v3062 = int32(1)
	goto L838
L837:
	;
	v3062 = v3058
	goto L838
L838:
	;
	v3108 = v3062
	goto L822
L839:
	;
	v3104 = v3101
	goto L825
L840:
	;
	v3080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3081 = v3080 + v3063
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3081
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3081
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3081
	v3085 = int32(0)
	v3088 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_161), int32(44))
	mBase = m.M
	v3089 = m.ExcPending
	if v3089 != 0 {
		goto L5
	} else {
		goto L845
	}
L841:
	;
	v3076 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_162))
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L5
	} else {
		goto L843
	}
L842:
	;
	switch v3072 {
	case 0:
		goto L840
	case 1:
		goto L841
	default:
		v3104 = v3069
		goto L825
	}
L843:
	;
	if v3076 < int32(0) {
		v3101 = v3076
		goto L839
	} else {
		goto L844
	}
L844:
	;
	v3104 = v3069
	goto L825
L845:
	;
	if v3088 == int32(0) {
		v3101 = v3085
		goto L839
	} else {
		goto L846
	}
L846:
	;
	v3092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3093 < v3092 {
		v3101 = v3085
		goto L839
	} else {
		goto L847
	}
L847:
	;
	v3097 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_163))
	mBase = m.M
	v3098 = m.ExcPending
	if v3098 != 0 {
		goto L5
	} else {
		goto L848
	}
L848:
	;
	if int32(0) <= v3097 {
		v3104 = v3069
		goto L825
	} else {
		goto L849
	}
L849:
	;
	v3101 = v3097
	goto L839
L850:
	;
	v3111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3112 = v3111 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3112
	v3114 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3112
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3112-int32(7) <= v3117 {
		v3166 = v3114
		goto L852
	} else {
		goto L853
	}
L851:
	;
	if v3169 < int32(0) {
		v3379 = v3169
		goto L1
	} else {
		goto L866
	}
L852:
	;
	v3169 = v3166
	goto L851
L853:
	;
	v3121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3121+v3112-int32(1)))))
	if v3125 != int32(181) {
		v3166 = v3114
		goto L852
	} else {
		goto L854
	}
L854:
	;
	v3130 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_164), int32(1))
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L5
	} else {
		goto L855
	}
L855:
	;
	if v3130 == int32(0) {
		v3166 = v3114
		goto L852
	} else {
		goto L856
	}
L856:
	;
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3134
	v3136 = F_slice_del(m, l0)
	mBase = m.M
	v3137 = m.ExcPending
	if v3137 != 0 {
		goto L5
	} else {
		goto L857
	}
L857:
	;
	if v3136 < int32(0) {
		v3166 = v3136
		goto L852
	} else {
		goto L858
	}
L858:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3141 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3140))) = v3141
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3143
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3143
	v3149 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_165), int32(10))
	mBase = m.M
	v3150 = m.ExcPending
	if v3150 != 0 {
		goto L5
	} else {
		goto L859
	}
L859:
	;
	if v3149 == int32(0) {
		v3169 = v3141
		goto L851
	} else {
		goto L860
	}
L860:
	;
	v3154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3155 < v3154 {
		v3166 = int32(0)
		goto L852
	} else {
		goto L861
	}
L861:
	;
	v3160 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_166))
	mBase = m.M
	v3161 = m.ExcPending
	if v3161 != 0 {
		goto L5
	} else {
		goto L862
	}
L862:
	;
	if int32(0) <= v3160 {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	v3164 = int32(1)
	goto L865
L864:
	;
	v3164 = v3160
	goto L865
L865:
	;
	v3166 = v3164
	goto L852
L866:
	;
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3173 = v3172 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3173
	v3175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3173
	v3178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3173-int32(7) <= v3178 {
		v3227 = v3175
		goto L868
	} else {
		goto L869
	}
L867:
	;
	if v3230 < int32(0) {
		v3379 = v3230
		goto L1
	} else {
		goto L882
	}
L868:
	;
	v3230 = v3227
	goto L867
L869:
	;
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3182+v3173-int32(1)))))
	if v3186 != int32(181) {
		v3227 = v3175
		goto L868
	} else {
		goto L870
	}
L870:
	;
	v3191 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_167), int32(3))
	mBase = m.M
	v3192 = m.ExcPending
	if v3192 != 0 {
		goto L5
	} else {
		goto L871
	}
L871:
	;
	if v3191 == int32(0) {
		v3227 = v3175
		goto L868
	} else {
		goto L872
	}
L872:
	;
	v3195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3195
	v3197 = F_slice_del(m, l0)
	mBase = m.M
	v3198 = m.ExcPending
	if v3198 != 0 {
		goto L5
	} else {
		goto L873
	}
L873:
	;
	if v3197 < int32(0) {
		v3227 = v3197
		goto L868
	} else {
		goto L874
	}
L874:
	;
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3202 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3201))) = v3202
	v3204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3204
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3204
	v3210 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_168), int32(6))
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L5
	} else {
		goto L875
	}
L875:
	;
	if v3210 == int32(0) {
		v3230 = v3202
		goto L867
	} else {
		goto L876
	}
L876:
	;
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3216 < v3215 {
		v3227 = int32(0)
		goto L868
	} else {
		goto L877
	}
L877:
	;
	v3221 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_169))
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L5
	} else {
		goto L878
	}
L878:
	;
	if int32(0) <= v3221 {
		goto L879
	} else {
		goto L880
	}
L879:
	;
	v3225 = int32(1)
	goto L881
L880:
	;
	v3225 = v3221
	goto L881
L881:
	;
	v3227 = v3225
	goto L868
L882:
	;
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3234 = v3233 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3234
	v3236 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3234
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3234-int32(7) <= v3239 {
		v3288 = v3236
		goto L884
	} else {
		goto L885
	}
L883:
	;
	if v3291 < int32(0) {
		v3379 = v3291
		goto L1
	} else {
		goto L898
	}
L884:
	;
	v3291 = v3288
	goto L883
L885:
	;
	v3243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3243+v3234-int32(1)))))
	if v3247 != int32(181) {
		v3288 = v3236
		goto L884
	} else {
		goto L886
	}
L886:
	;
	v3252 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_170), int32(3))
	mBase = m.M
	v3253 = m.ExcPending
	if v3253 != 0 {
		goto L5
	} else {
		goto L887
	}
L887:
	;
	if v3252 == int32(0) {
		v3288 = v3236
		goto L884
	} else {
		goto L888
	}
L888:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3256
	v3258 = F_slice_del(m, l0)
	mBase = m.M
	v3259 = m.ExcPending
	if v3259 != 0 {
		goto L5
	} else {
		goto L889
	}
L889:
	;
	if v3258 < int32(0) {
		v3288 = v3258
		goto L884
	} else {
		goto L890
	}
L890:
	;
	v3262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3263 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3262))) = v3263
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3265
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3265
	v3271 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_171), int32(7))
	mBase = m.M
	v3272 = m.ExcPending
	if v3272 != 0 {
		goto L5
	} else {
		goto L891
	}
L891:
	;
	if v3271 == int32(0) {
		v3291 = v3263
		goto L883
	} else {
		goto L892
	}
L892:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3277 < v3276 {
		v3288 = int32(0)
		goto L884
	} else {
		goto L893
	}
L893:
	;
	v3282 = F_slice_from_s(m, l0, int32(6), int32(_a_F_greek_UTF_8_stem_172))
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		goto L5
	} else {
		goto L894
	}
L894:
	;
	if int32(0) <= v3282 {
		goto L895
	} else {
		goto L896
	}
L895:
	;
	v3286 = int32(1)
	goto L897
L896:
	;
	v3286 = v3282
	goto L897
L897:
	;
	v3288 = v3286
	goto L884
L898:
	;
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3295 = v3294 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3295
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3295
	v3302 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_173), int32(3))
	mBase = m.M
	v3303 = m.ExcPending
	if v3303 != 0 {
		goto L5
	} else {
		goto L900
	}
L899:
	;
	if v3337 < int32(0) {
		v3379 = v3337
		goto L1
	} else {
		goto L913
	}
L900:
	;
	if v3302 != 0 {
		goto L901
	} else {
		goto L902
	}
L901:
	;
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3304
	v3308 = F_slice_from_s(m, l0, int32(4), int32(_a_F_greek_UTF_8_stem_174))
	mBase = m.M
	v3309 = m.ExcPending
	if v3309 != 0 {
		goto L5
	} else {
		goto L904
	}
L902:
	;
	goto L903
L903:
	;
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3315 = v3313 + (v3295 - v3294)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3315
	v3317 = int32(0)
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v3318)))
	if v3319 == v3317 {
		v3337 = v3317
		goto L899
	} else {
		goto L906
	}
L904:
	;
	if v3308 < int32(0) {
		v3337 = v3308
		goto L899
	} else {
		goto L905
	}
L905:
	;
	goto L903
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3315
	v3325 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_175), int32(84))
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L5
	} else {
		goto L907
	}
L907:
	;
	if v3325 == int32(0) {
		v3337 = v3317
		goto L899
	} else {
		goto L908
	}
L908:
	;
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3329
	v3332 = F_slice_del(m, l0)
	mBase = m.M
	v3333 = m.ExcPending
	if v3333 != 0 {
		goto L5
	} else {
		goto L909
	}
L909:
	;
	if int32(0) <= v3332 {
		goto L910
	} else {
		goto L911
	}
L910:
	;
	v3336 = int32(1)
	goto L912
L911:
	;
	v3336 = v3332
	goto L912
L912:
	;
	v3337 = v3336
	goto L899
L913:
	;
	v3341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3342 = v3341 + v405
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3342
	v3344 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3342
	v3347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3342-int32(7) <= v3347 {
		v3372 = v3344
		goto L914
	} else {
		goto L915
	}
L914:
	;
	if v3372 < int32(0) {
		v3379 = v3372
		goto L1
	} else {
		goto L923
	}
L915:
	;
	v3351 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3351+v3342-int32(1)))))
	switch v3355 - int32(129) {
	case 0, 3:
		goto L916
	default:
		v3372 = v3344
		goto L914
	}
L916:
	;
	v3360 = F_find_among_b(m, l0, int32(_a_F_greek_UTF_8_stem_176), int32(8))
	mBase = m.M
	v3361 = m.ExcPending
	if v3361 != 0 {
		goto L5
	} else {
		goto L917
	}
L917:
	;
	if v3360 == int32(0) {
		v3372 = v3344
		goto L914
	} else {
		goto L918
	}
L918:
	;
	v3364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3364
	v3367 = F_slice_del(m, l0)
	mBase = m.M
	v3368 = m.ExcPending
	if v3368 != 0 {
		goto L5
	} else {
		goto L919
	}
L919:
	;
	if int32(0) <= v3367 {
		goto L920
	} else {
		goto L921
	}
L920:
	;
	v3371 = int32(1)
	goto L922
L921:
	;
	v3371 = v3367
	goto L922
L922:
	;
	v3372 = v3371
	goto L914
L923:
	;
	v3375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3375
	v3379 = int32(1)
	goto L1
}
