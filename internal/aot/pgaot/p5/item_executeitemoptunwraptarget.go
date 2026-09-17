package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_executeItemOptUnwrapTarget(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
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
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int64
	_ = v515
	var v517 int64
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v574 int32
	_ = v574
	var v592 int32
	_ = v592
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v826 int64
	_ = v826
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v973 int64
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1094 int64
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1338 int64
	_ = v1338
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1475 int64
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1521 int64
	_ = v1521
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1584 int32
	_ = v1584
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1682 int32
	_ = v1682
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1734 int64
	_ = v1734
	var v1735 int64
	_ = v1735
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1783 int32
	_ = v1783
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1863 int32
	_ = v1863
	var v1869 int32
	_ = v1869
	var v1871 int64
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1888 int64
	_ = v1888
	var v1890 int64
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1950 int32
	_ = v1950
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2023 int32
	_ = v2023
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2046 int32
	_ = v2046
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2065 int32
	_ = v2065
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2085 int64
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2120 int32
	_ = v2120
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2147 int64
	_ = v2147
	var v2151 int32
	_ = v2151
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2234 int32
	_ = v2234
	var v2238 int32
	_ = v2238
	var v2243 int32
	_ = v2243
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2285 int32
	_ = v2285
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2298 int32
	_ = v2298
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
	var v2313 int32
	_ = v2313
	var v2319 int32
	_ = v2319
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2358 int32
	_ = v2358
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2378 int32
	_ = v2378
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2388 int32
	_ = v2388
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2412 int32
	_ = v2412
	var v2419 int32
	_ = v2419
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2432 int32
	_ = v2432
	var v2437 int32
	_ = v2437
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2452 int32
	_ = v2452
	var v2457 int32
	_ = v2457
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2479 int64
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2489 int32
	_ = v2489
	var v2490 int64
	_ = v2490
	var v2491 int64
	_ = v2491
	var v2492 int64
	_ = v2492
	var v2495 int64
	_ = v2495
	var v2496 int64
	_ = v2496
	var v2498 int64
	_ = v2498
	var v2499 int64
	_ = v2499
	var v2502 int64
	_ = v2502
	var v2508 int64
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2538 int32
	_ = v2538
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2548 int32
	_ = v2548
	var v2550 int32
	_ = v2550
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2566 int32
	_ = v2566
	var v2567 int64
	_ = v2567
	var v2568 int64
	_ = v2568
	var v2569 int64
	_ = v2569
	var v2572 int64
	_ = v2572
	var v2573 int64
	_ = v2573
	var v2575 int64
	_ = v2575
	var v2576 int64
	_ = v2576
	var v2579 int64
	_ = v2579
	var v2589 int32
	_ = v2589
	var v2600 int32
	_ = v2600
	var v2607 int32
	_ = v2607
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2620 int32
	_ = v2620
	var v2625 int32
	_ = v2625
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2639 int32
	_ = v2639
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2660 int32
	_ = v2660
	var v2663 int64
	_ = v2663
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2668 int64
	_ = v2668
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2698 int32
	_ = v2698
	var v2703 int32
	_ = v2703
	var v2704 int64
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2721 int32
	_ = v2721
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2740 int32
	_ = v2740
	var v2745 int32
	_ = v2745
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2758 int32
	_ = v2758
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2768 int32
	_ = v2768
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2789 int32
	_ = v2789
	var v2794 int32
	_ = v2794
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2810 int32
	_ = v2810
	var v2814 int32
	_ = v2814
	var v2815 int32
	_ = v2815
	var v2817 int32
	_ = v2817
	var v2825 int32
	_ = v2825
	var v2829 int32
	_ = v2829
	var v2839 int32
	_ = v2839
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2858 int32
	_ = v2858
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int64
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2884 int32
	_ = v2884
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2892 int32
	_ = v2892
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2902 int32
	_ = v2902
	var v2906 int32
	_ = v2906
	var v2909 int64
	_ = v2909
	var v2912 int32
	_ = v2912
	var v2913 int32
	_ = v2913
	var v2914 int64
	_ = v2914
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2925 int32
	_ = v2925
	var v2929 int64
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2934 int32
	_ = v2934
	var v2938 int32
	_ = v2938
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2965 int32
	_ = v2965
	var v2967 int32
	_ = v2967
	var v2969 int32
	_ = v2969
	var v2973 int32
	_ = v2973
	var v2974 int32
	_ = v2974
	var v2978 int32
	_ = v2978
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2990 int32
	_ = v2990
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v3003 int32
	_ = v3003
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3008 int32
	_ = v3008
	var v3016 int32
	_ = v3016
	var v3021 int32
	_ = v3021
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3033 int32
	_ = v3033
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3057 int32
	_ = v3057
	var v3060 int64
	_ = v3060
	var v3066 float64
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3071 int32
	_ = v3071
	var v3077 int32
	_ = v3077
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3092 int32
	_ = v3092
	var v3097 int32
	_ = v3097
	var v3098 float64
	_ = v3098
	var v3105 int32
	_ = v3105
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3118 int64
	_ = v3118
	var v3124 float64
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3129 int32
	_ = v3129
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3150 int32
	_ = v3150
	var v3155 int32
	_ = v3155
	var v3156 float64
	_ = v3156
	var v3163 int32
	_ = v3163
	var v3169 int32
	_ = v3169
	var v3175 int32
	_ = v3175
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3180 int32
	_ = v3180
	var v3181 int32
	_ = v3181
	var v3187 int32
	_ = v3187
	var v3192 int32
	_ = v3192
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3206 int32
	_ = v3206
	var v3212 int32
	_ = v3212
	var v3213 int32
	_ = v3213
	var v3217 int32
	_ = v3217
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3229 int32
	_ = v3229
	var v3234 int32
	_ = v3234
	var v3238 int32
	_ = v3238
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3250 int32
	_ = v3250
	var v3255 int32
	_ = v3255
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3260 int32
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3277 int32
	_ = v3277
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3287 int32
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3299 int32
	_ = v3299
	var v3304 int32
	_ = v3304
	var v3306 int32
	_ = v3306
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3316 int32
	_ = v3316
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3341 int32
	_ = v3341
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3347 int32
	_ = v3347
	var v3348 int32
	_ = v3348
	var v3350 int32
	_ = v3350
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3360 int64
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3377 int32
	_ = v3377
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3382 int32
	_ = v3382
	var v3385 int32
	_ = v3385
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3394 int32
	_ = v3394
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3404 int32
	_ = v3404
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3421 int32
	_ = v3421
	var v3422 int32
	_ = v3422
	var v3423 int32
	_ = v3423
	var v3425 int32
	_ = v3425
	var v3432 int32
	_ = v3432
	var v3435 int32
	_ = v3435
	var v3439 int32
	_ = v3439
	var v3444 int32
	_ = v3444
	var v3448 int32
	_ = v3448
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3461 int32
	_ = v3461
	var v3466 int32
	_ = v3466
	var v3468 int32
	_ = v3468
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3474 int32
	_ = v3474
	var v3477 int32
	_ = v3477
	var v3478 int32
	_ = v3478
	var v3482 int32
	_ = v3482
	var v3484 int32
	_ = v3484
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3492 int32
	_ = v3492
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3507 int32
	_ = v3507
	var v3510 int32
	_ = v3510
	var v3514 int32
	_ = v3514
	var v3519 int32
	_ = v3519
	var v3524 int32
	_ = v3524
	v6 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(992)
	m.G0 = v24
	F_check_stack_depth(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[0]))
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v34 {
	case 0, 1, 2, 3, 28:
		goto L42
	case 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 30, 41, 42:
		goto L41
	case 14:
		goto L29
	case 15:
		goto L40
	case 16:
		goto L39
	case 17:
		goto L38
	case 18:
		goto L37
	case 19:
		goto L36
	case 20:
		goto L35
	case 21:
		goto L34
	case 22:
		goto L33
	case 23:
		goto L32
	case 24:
		goto L9
	case 25:
		goto L10
	case 26:
		goto L11
	case 27:
		goto L12
	case 29:
		goto L13
	case 31:
		goto L14
	case 32:
		goto L15
	case 33:
		goto L16
	case 34:
		goto L17
	case 35:
		goto L18
	case 36:
		goto L19
	case 37, 45, 50, 51, 52, 53:
		goto L20
	case 38:
		goto L21
	default:
		goto L28
	case 40:
		goto L22
	case 43:
		goto L23
	case 44:
		goto L24
	case 46, 48:
		goto L25
	case 47:
		goto L26
	case 49:
		goto L27
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v24 + int32(992)
	return v3524
L8:
	;
	v3499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3499 != 0 {
		v3524 = v328
		goto L7
	} else {
		goto L1107
	}
L9:
	;
	v3468 = v24 + int32(752)
	v3469 = F_jspGetNext(m, l1, v3468)
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L1
	} else {
		goto L1097
	}
L10:
	;
	v3381 = F_JsonbType(m, l2)
	mBase = m.M
	v3382 = m.ExcPending
	if v3382 != 0 {
		goto L1
	} else {
		goto L1057
	}
L11:
	;
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3379 = F_executeNextItem(m, l0, l1, int32(0), v3377, l3, int32(1))
	mBase = m.M
	v3380 = m.ExcPending
	if v3380 != 0 {
		goto L1
	} else {
		goto L1055
	}
L12:
	;
	v3360 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v3362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3363 = *(*int32)(unsafe.Add(mBase, uint32(v3362)))
	if v3363 == int32(18) {
		goto L1051
	} else {
		goto L1052
	}
L13:
	;
	if l4 != 0 {
		goto L1041
	} else {
		goto L1042
	}
L14:
	;
	v3321 = F_palloc(m, int32(20))
	mBase = m.M
	v3322 = m.ExcPending
	if v3322 != 0 {
		goto L1
	} else {
		goto L1036
	}
L15:
	;
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3265 != int32(18) {
		goto L1020
	} else {
		goto L1021
	}
L16:
	;
	v3263 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1276), l3)
	mBase = m.M
	v3264 = m.ExcPending
	if v3264 != 0 {
		goto L1
	} else {
		goto L1018
	}
L17:
	;
	v3260 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1422), l3)
	mBase = m.M
	v3261 = m.ExcPending
	if v3261 != 0 {
		goto L1
	} else {
		goto L1017
	}
L18:
	;
	v3257 = F_executeNumericItemMethod(m, l0, l1, l2, l4, int32(1421), l3)
	mBase = m.M
	v3258 = m.ExcPending
	if v3258 != 0 {
		goto L1
	} else {
		goto L1016
	}
L19:
	;
	if l4 == int32(0) {
		goto L956
	} else {
		goto L957
	}
L20:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l4 == int32(0) {
		goto L648
	} else {
		goto L649
	}
L21:
	;
	if l4 == int32(0) {
		goto L562
	} else {
		goto L563
	}
L22:
	;
	v1612 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L548
	}
L23:
	;
	if l4 == int32(0) {
		goto L502
	} else {
		goto L503
	}
L24:
	;
	if l4 == int32(0) {
		goto L455
	} else {
		goto L456
	}
L25:
	;
	if l4 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L26:
	;
	if l4 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L27:
	;
	if l4 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L262
	}
L29:
	;
	v674 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1407), l3)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L261
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L256
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L253
	}
L32:
	;
	v324 = F_JsonbType(m, l2)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L160
	}
L33:
	;
	v272 = F_JsonbType(m, l2)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L135
	}
L34:
	;
	v230 = F_JsonbType(m, l2)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L114
	}
L35:
	;
	v228 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(1406), l3)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L113
	}
L36:
	;
	v225 = F_executeUnaryArithmExpr(m, l0, l1, l2, int32(0), l3)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L112
	}
L37:
	;
	v222 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1405), l3)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L111
	}
L38:
	;
	v219 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1404), l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L110
	}
L39:
	;
	v216 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1403), l3)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L109
	}
L40:
	;
	v213 = F_executeBinaryArithmExpr(m, l0, l1, l2, int32(1402), l3)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L108
	}
L41:
	;
	v184 = F_executeBoolItem(m, l0, l1, l2, int32(1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L98
	}
L42:
	;
	v37 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v37|l3 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v43 != int32(28) {
		v3524 = int32(0)
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v37 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v52 = F_palloc(m, int32(20))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v54 = v24 + int32(908)
	goto L50
L50:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v56 {
	case 0:
		goto L53
	case 1:
		goto L56
	case 2:
		goto L57
	case 3:
		goto L58
	default:
		goto L54
	case 28:
		goto L55
	}
L51:
	;
	v54 = v52
	goto L50
L52:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v129 {
		goto L77
	} else {
		goto L78
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(0)
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L74
	}
L55:
	;
	v77 = v24 + int32(936)
	if v77 != 0 {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(1)
	v71 = v54 + int32(4)
	if v71 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(2)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v66
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(3)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)) = uint8(base.B2i32(v60 != int32(0)))
	goto L52
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v74
	goto L52
L60:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72
	goto L62
L61:
	;
	goto L62
L62:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L59
L63:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v81 == int32(0) {
		goto L30
	} else {
		goto L67
	}
L64:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v78
	goto L66
L65:
	;
	goto L66
L66:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L63
L67:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = m.T0[v89].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v81, v80, v84, v24+int32(952), v24+int32(888))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v90 == int32(0) {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v24)+888))
	if v94 <= int32(0) {
		goto L52
	} else {
		goto L70
	}
L70:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v90)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+960))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v94
	if v104 == int32(18) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v109 = v103
	goto L73
L72:
	;
	v109 = int32(0)
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v109
	goto L52
L74:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_0), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2982), int32(_a_F_executeItemOptUnwrapTarget_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
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
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v135 = F_executeItemOptUnwrapTarget(m, l0, v24+int32(752), v54, l3, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if l3 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v55
	v3524 = v135
	goto L7
L81:
	;
	if v37 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	goto L83
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v55
	v3524 = int32(0)
	goto L7
L84:
	;
	v151 = m.G0
	v153 = v151 - int32(16)
	m.G0 = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v155 != 0 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v150 = v54
	goto L84
L86:
	;
	goto L87
L87:
	;
	v142 = F_palloc(m, int32(20))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = v144
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v142)+8)) = v146
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v148
	v150 = v142
	goto L84
L89:
	;
	m.G0 = v153 + int32(16)
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v150
	v162 = F_list_make2_impl(m, v153+int32(4), v153)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v167 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v162
	goto L89
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v150
	goto L89
L95:
	;
	goto L96
L96:
	;
	v171 = F_lappend(m, v167, v150)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v171
	goto L89
L98:
	;
	v188 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if l3 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v192 = int32(0)
	if v188 == v192 {
		v3524 = v192
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v184 != int32(2) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)) = uint8(base.B2i32(v184 == int32(1)))
	v203 = int32(3)
	goto L106
L105:
	;
	v203 = int32(0)
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = v203
	v210 = F_executeNextItem(m, l0, l1, v24+int32(752), v24+int32(952), l3, int32(1))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v3524 = v210
	goto L7
L108:
	;
	v3524 = v213
	goto L7
L109:
	;
	v3524 = v216
	goto L7
L110:
	;
	v3524 = v219
	goto L7
L111:
	;
	v3524 = v222
	goto L7
L112:
	;
	v3524 = v225
	goto L7
L113:
	;
	v3524 = v228
	goto L7
L114:
	;
	if v230 == int32(16) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v235 = v24 + int32(752)
	v237 = F_jspGetNext(m, l1, v235)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v243 = int32(1)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v244 == v243 {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	if v237 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v239 = v235
	goto L121
L120:
	;
	v239 = int32(0)
	goto L121
L121:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v241 = F_executeItemUnwrapTargetArray(m, l0, v239, l2, l3, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v3524 = v241
	goto L7
L123:
	;
	v249 = F_executeNextItem(m, l0, l1, int32(0), l2, l3, int32(1))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v251 != 0 {
		v3524 = v243
		goto L7
	} else {
		goto L127
	}
L126:
	;
	v3524 = v249
	goto L7
L127:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v252 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v3524 = int32(2)
	goto L7
L129:
	;
	goto L130
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_3), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(849), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	if v272 == int32(17) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v277 = v24 + int32(752)
	v278 = F_jspGetNext(m, l1, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	if l4 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v280 != int32(18) {
		goto L31
	} else {
		goto L140
	}
L140:
	;
	if v278 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v284 = v277
	goto L143
L142:
	;
	v284 = int32(0)
	goto L143
L143:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v286 = int32(1)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v291 = F_executeAnyItem(m, l0, v284, v285, l3, v286, v286, v286, int32(0), v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v3524 = v291
	goto L7
L145:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v302 != 0 {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	v295 = F_JsonbType(m, l2)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v295 != int32(16) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v300 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v3524 = v300
	goto L7
L150:
	;
	v3524 = int32(1)
	goto L7
L151:
	;
	goto L152
L152:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v304 != int32(1) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v3524 = int32(2)
	goto L7
L154:
	;
	goto L155
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_5), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(872), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	if v324 != int32(16) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v328 = int32(1)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v329 != v328 {
		goto L8
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v333 != int32(18) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L163
L165:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = F_jspGetNext(m, l1, v24+int32(752))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L169
	}
L166:
	;
	v347 = int32(-1)
	goto L165
L167:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v337&int32(1342177280) != int32(1073741824) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v347 = v337 & int32(268435455)
	goto L165
L169:
	;
	v356 = base.B2i32(v347 < int32(0))
	if v347 < int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v357 = int32(1)
	goto L172
L171:
	;
	v357 = v347
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v359 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v363 = v357 - int32(1)
	v381 = v6
	goto L176
L174:
	;
	v617 = int32(1)
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v348
	v3524 = v617
	goto L7
L176:
	;
	v388 = int32(2)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v395 = v381 << (uint(int32(3)) % 32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395+v396)))
	F_jspInitByBuffer(m, v24+int32(952), v393, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	v617 = v592
	goto L175
L178:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v401+v395)+4))
	if v403 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v24+int32(908), v404, v403)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v413 = F_getArrayIndex(m, l0, v24+int32(952), l2, v24+int32(900))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L185
	}
L182:
	;
	goto L181
L183:
	;
	v610 = v381 + int32(1)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v610 < v611 {
		v381 = v610
		goto L176
	} else {
		goto L252
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v348
	v3524 = v574
	goto L7
L185:
	;
	if v413 == int32(2) {
		v574 = v388
		goto L184
	} else {
		goto L186
	}
L186:
	;
	if v403 != int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v429 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L188:
	;
	v421 = F_getArrayIndex(m, l0, v24+int32(908), l2, v24+int32(988))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v426
	v428 = v426
	goto L187
L191:
	;
	if v421 == int32(2) {
		v574 = v388
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	v428 = v425
	goto L187
L193:
	;
	if v459 < v363 {
		goto L208
	} else {
		goto L209
	}
L194:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v459 = v432
	goto L193
L195:
	;
	goto L196
L196:
	;
	if v428 < int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v439 != int32(1) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	if v435 < v428 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	if v435 < v357 {
		v459 = v435
		goto L193
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	v3524 = int32(2)
	goto L7
L202:
	;
	goto L203
L203:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errcode(m, int32(51118210))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_6), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(921), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v461 = v459
	goto L210
L209:
	;
	v461 = v363
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v461
	v463 = int32(0)
	if v463 < v428 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v466 = v428
	goto L213
L212:
	;
	v466 = v463
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = v466
	v468 = int32(1)
	if v461 < v466 {
		v592 = v468
		goto L183
	} else {
		goto L214
	}
L214:
	;
	v474 = v468
	v477 = v466
	v479 = v461
	goto L215
L215:
	;
	if v356 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	if l3|v560 != 0 {
		v592 = v560
		goto L183
	} else {
		goto L251
	}
L217:
	;
	goto L216
L218:
	;
	v558 = v477 + int32(1)
	if v558 <= v556 {
		v474 = v554
		v477 = v558
		v479 = v556
		goto L215
	} else {
		goto L250
	}
L219:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v494 = F_getIthJsonbValueFromContainer(m, v493, v477)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v498 = l2
	goto L221
L221:
	;
	if base.B2i32(l3 != int32(0))|v352 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	if v494 == int32(0) {
		v554 = v474
		v556 = v479
		goto L218
	} else {
		goto L223
	}
L223:
	;
	v498 = v494
	goto L221
L224:
	;
	v3524 = int32(0)
	goto L7
L225:
	;
	goto L226
L226:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v502 <= int32(0) {
		goto L230
	} else {
		goto L231
	}
L227:
	;
	v554 = int32(0)
	v556 = v552
	goto L218
L228:
	;
	v546 = F_lappend(m, v534, v519)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L249
	}
L229:
	;
	if l3|v543 != 0 {
		v554 = v543
		v556 = v479
		goto L218
	} else {
		goto L248
	}
L230:
	;
	if l3 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	goto L232
L232:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v539 = F_executeItemOptUnwrapTarget(m, l0, v24+int32(752), v498, l3, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L246
	}
L233:
	;
	v543 = int32(0)
	goto L229
L234:
	;
	goto L235
L235:
	;
	if int32(0) <= v347 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v520 != 0 {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	v519 = v498
	goto L236
L238:
	;
	goto L239
L239:
	;
	v511 = F_palloc(m, int32(20))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v498)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v511)+16)) = v513
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v498)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v511)+8)) = v515
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v498)))
	*(*int64)(unsafe.Add(mBase, uint32(v511))) = v517
	v519 = v511
	goto L236
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+888)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v24)+936)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v519
	v529 = F_list_make2_impl(m, v24+int32(44), v24+int32(40))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v534 != 0 {
		goto L228
	} else {
		goto L245
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v529
	v552 = v479
	goto L227
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v519
	v552 = v479
	goto L227
L246:
	;
	if v539 == int32(2) {
		v574 = v388
		goto L184
	} else {
		goto L247
	}
L247:
	;
	v543 = v539
	goto L229
L248:
	;
	v560 = int32(0)
	goto L217
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v546
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v552 = v549
	goto L227
L250:
	;
	v560 = v554
	goto L217
L251:
	;
	v574 = int32(0)
	goto L184
L252:
	;
	goto L177
L253:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v639
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_7), v24+int32(32))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(858), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v660 = F_pnstrdup(m, v80, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v660
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_8), v24+int32(16))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(3158), int32(_a_F_executeItemOptUnwrapTarget_9))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	v3524 = v674
	goto L7
L262:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v680
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_10), v24)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1663), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	v699 = F_JsonbType(m, l2)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L277
	}
L266:
	;
	v692 = F_JsonbType(m, l2)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	if v692 != int32(16) {
		goto L265
	} else {
		goto L268
	}
L268:
	;
	v697 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v3524 = v697
	goto L7
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v753
	v755 = F_strlen(m, v753)
	mBase = m.M
	v756 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v24)+956)) = v755
	v763 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(952), l3, v756)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L291
	}
L271:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v750 = F_pnstrdup(m, v748, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L290
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v723 != int32(1) {
		v3524 = int32(2)
		goto L7
	} else {
		goto L284
	}
L274:
	;
	v713 = v24 + int32(752)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v718 = F_JsonEncodeDateTime(m, v713, v714, v715, l2+int32(16))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L282
	}
L275:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v710 != 0 {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v706 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L278
	}
L277:
	;
	switch v699 - int32(1) {
	case 0:
		goto L271
	case 1:
		goto L276
	case 2:
		goto L275
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
		goto L272
	default:
		goto L273
	case 31:
		goto L274
	}
L278:
	;
	v753 = v706
	goto L270
L279:
	;
	v711 = int32(_a_F_executeItemOptUnwrapTarget_11)
	goto L281
L280:
	;
	v711 = int32(_a_F_executeItemOptUnwrapTarget_12)
	goto L281
L281:
	;
	v753 = v711
	goto L270
L282:
	;
	v720 = F_pstrdup(m, v713)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v753 = v720
	goto L270
L284:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v734 = F_jspOperationName(m, v733)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+736)) = v734
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_13), v24+int32(736))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1648), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	goto L272
L290:
	;
	v753 = v750
	goto L270
L291:
	;
	v3524 = v763
	goto L7
L292:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v774 - int32(1) {
	case 0:
		goto L300
	case 1:
		goto L301
	default:
		goto L299
	}
L293:
	;
	v767 = F_JsonbType(m, l2)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	if v767 != int32(16) {
		goto L292
	} else {
		goto L295
	}
L295:
	;
	v772 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v3524 = v772
	goto L7
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v902 = F_DirectFunctionCall1Coll(m, int32(1398), int32(0), v897)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L334
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+908)) = v780
	v897 = v780
	goto L297
L299:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v871 != int32(1) {
		v3524 = int32(2)
		goto L7
	} else {
		goto L328
	}
L300:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v820 = F_pnstrdup(m, v818, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L313
	}
L301:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v780 = F_numeric_int4_opt_error(m, v777, v24+int32(952))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+952)))
	if v782 != int32(1) {
		goto L298
	} else {
		goto L303
	}
L303:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v785 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v3524 = int32(2)
	goto L7
L305:
	;
	goto L306
L306:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v799 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v802 = F_jspOperationName(m, v801)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+712)) = int32(_a_F_executeItemOptUnwrapTarget_14)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+708)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v24)+704)) = v799
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(704))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1563), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v823
	v826 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v826
	v834 = F_DirectInputFunctionCallSafe(m, int32(1408), v820, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L315
	}
L314:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v897 = v869
	goto L297
L315:
	;
	if v834 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v836&int32(1) == int32(0) {
		goto L314
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v841 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	goto L318
L320:
	;
	v3524 = int32(2)
	goto L7
L321:
	;
	goto L322
L322:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L323
	}
L323:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v853 = F_jspOperationName(m, v852)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+728)) = int32(_a_F_executeItemOptUnwrapTarget_14)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+724)) = v853
	*(*int32)(unsafe.Add(mBase, uint32(v24)+720)) = v820
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(720))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1585), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L328:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v882 = F_jspOperationName(m, v881)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+688)) = v882
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_16), v24+int32(688))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1593), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	v904 = F_pg_detoast_datum(m, v902)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v904
	v911 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	v3524 = v911
	goto L7
L337:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v922 - int32(1) {
	case 0:
		goto L349
	case 1:
		goto L350
	default:
		goto L348
	}
L338:
	;
	v915 = F_JsonbType(m, l2)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	if v915 != int32(16) {
		goto L337
	} else {
		goto L340
	}
L340:
	;
	v920 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v3524 = v920
	goto L7
L342:
	;
	v3524 = int32(2)
	goto L7
L343:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L1
	} else {
		goto L452
	}
L344:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L449
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+956)) = v1273
	*(*int32)(unsafe.Add(mBase, uint32(v24)+952)) = int32(2)
	v1284 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(952), l3, int32(1))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L448
	}
L346:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1084 != int32(46) {
		v1273 = v1082
		goto L345
	} else {
		goto L398
	}
L347:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L393
	}
L348:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1037 != int32(1) {
		goto L342
	} else {
		goto L387
	}
L349:
	;
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v970
	v973 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v973
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v978 = F_pnstrdup(m, v976, v977)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L367
	}
L350:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v925)+4)))
	goto L352
L351:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v962 != int32(46) {
		v1273 = v925
		goto L345
	} else {
		goto L364
	}
L352:
	;
	if base.B2i32(v926 == int32(_a_F_executeItemOptUnwrapTarget_17)) == int32(0) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v931 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v925)+4)))
	goto L356
L354:
	;
	goto L355
L355:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v938 != int32(1) {
		goto L342
	} else {
		goto L358
	}
L356:
	;
	if base.B2i32(v931&int32(_a_F_executeItemOptUnwrapTarget_18) == int32(_a_F_executeItemOptUnwrapTarget_19)) == int32(0) {
		goto L351
	} else {
		goto L357
	}
L357:
	;
	goto L355
L358:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v949 = F_jspOperationName(m, v948)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+592)) = v949
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_20), v24+int32(592))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1414), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	v967 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v925)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1082 = v925
	v1083 = v967
	goto L346
L366:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1020 = F_pg_detoast_datum(m, v1019)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L379
	}
L367:
	;
	v985 = F_DirectInputFunctionCallSafe(m, int32(408), v978, int32(-1), v24+int32(752), v24+int32(908))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	if v985 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v987&int32(1) == int32(0) {
		goto L366
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v992 != int32(1) {
		goto L342
	} else {
		goto L373
	}
L372:
	;
	goto L371
L373:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1003 = F_jspOperationName(m, v1002)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+664)) = int32(_a_F_executeItemOptUnwrapTarget_21)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+660)) = v1003
	*(*int32)(unsafe.Add(mBase, uint32(v24)+656)) = v978
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(656))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1439), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L379:
	;
	v1022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1020)+4)))
	goto L380
L380:
	;
	if base.B2i32(v1022 == int32(_a_F_executeItemOptUnwrapTarget_17)) == int32(0) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1027 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1020)+4)))
	goto L384
L382:
	;
	goto L383
L383:
	;
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1034 == int32(1) {
		goto L347
	} else {
		goto L386
	}
L384:
	;
	if base.B2i32(v1027&int32(_a_F_executeItemOptUnwrapTarget_18) == int32(_a_F_executeItemOptUnwrapTarget_19)) == int32(0) {
		v1082 = v1020
		v1083 = v978
		goto L346
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	goto L342
L387:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1048 = F_jspOperationName(m, v1047)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+576)) = v1048
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_16), v24+int32(576))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1455), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1069 = F_jspOperationName(m, v1068)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+672)) = v1069
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_20), v24+int32(672))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1446), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L398:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v1087 == int32(0) {
		v1273 = v1082
		goto L345
	} else {
		goto L399
	}
L399:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+896)) = v1091
	v1094 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+888)) = v1094
	v1097 = v24 + int32(752)
	F_jspGetArg(m, l1, v1097)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+752))
	if v1100 != int32(2) {
		goto L344
	} else {
		goto L401
	}
L401:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+12))
	v1106 = F_numeric_int4_opt_error(m, v1103, v24+int32(984))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+984)))
	if v1108 == int32(1) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1111 != int32(1) {
		goto L342
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1135 == int32(0) {
		v1178 = v6
		goto L412
	} else {
		goto L413
	}
L406:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1122 = F_jspOperationName(m, v1121)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+608)) = v1122
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_22), v24+int32(608))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1487), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L412:
	;
	v1180 = v24 + int32(908)
	if int32(0) <= v1106 {
		goto L425
	} else {
		goto L426
	}
L413:
	;
	v1139 = v24 + int32(752)
	F_jspGetRightArg(m, l1, v1139)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v24)+752))
	if v1142 != int32(2) {
		goto L343
	} else {
		goto L415
	}
L415:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+12))
	v1148 = F_numeric_int4_opt_error(m, v1145, v24+int32(984))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+984)))
	if v1150 != int32(1) {
		v1178 = v1148
		goto L412
	} else {
		goto L417
	}
L417:
	;
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1153 != int32(1) {
		goto L342
	} else {
		goto L418
	}
L418:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L419
	}
L419:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L420
	}
L420:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1164 = F_jspOperationName(m, v1163)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+640)) = v1164
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_23), v24+int32(640))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1501), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = v1180
	v1200 = v24 + int32(936)
	if int32(0) <= v1178 {
		goto L429
	} else {
		goto L430
	}
L425:
	;
	v1190 = v1106
	v1191 = int32(0)
	goto L427
L426:
	;
	v1185 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1180))) = uint8(v1185)
	v1190 = int32(0) - v1106
	v1191 = int32(1)
	goto L427
L427:
	;
	v1193 = F_pg_ultoa_n(m, v1190, v1180+v1191)
	mBase = m.M
	v1196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1180+(v1193+v1191)))) = uint8(v1196)
	goto L424
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+904)) = v1200
	v1226 = F_construct_array_builtin(m, v24+int32(900), int32(2), int32(2275))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L433
	}
L429:
	;
	v1210 = v1178
	v1211 = int32(0)
	goto L431
L430:
	;
	v1205 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1200))) = uint8(v1205)
	v1210 = int32(0) - v1178
	v1211 = int32(1)
	goto L431
L431:
	;
	v1213 = F_pg_ultoa_n(m, v1210, v1200+v1211)
	mBase = m.M
	v1216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1200+(v1213+v1211)))) = uint8(v1216)
	goto L428
L432:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	v1269 = F_pg_detoast_datum(m, v1268)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L446
	}
L433:
	;
	v1228 = F_DirectFunctionCall1Coll(m, int32(1409), int32(0), v1226)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	v1234 = F_DirectInputFunctionCallSafe(m, int32(408), v1083, v1228, v24+int32(888), v24+int32(988))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	if v1234 != 0 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+892)))
	if v1236&int32(1) == int32(0) {
		goto L432
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1241 != int32(1) {
		goto L342
	} else {
		goto L440
	}
L439:
	;
	goto L438
L440:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1252 = F_jspOperationName(m, v1251)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+632)) = int32(_a_F_executeItemOptUnwrapTarget_21)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+628)) = v1252
	*(*int32)(unsafe.Add(mBase, uint32(v24)+624)) = v1083
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(624))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1528), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	F_pfree(m, v1226)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	v1273 = v1269
	goto L345
L448:
	;
	v3524 = v1284
	goto L7
L449:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_24), int32(0))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1479), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_25), int32(0))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1493), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L455:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1326 - int32(1) {
	case 0:
		goto L465
	case 1:
		goto L466
	case 2:
		goto L463
	default:
		goto L464
	}
L456:
	;
	v1319 = F_JsonbType(m, l2)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	if v1319 != int32(16) {
		goto L455
	} else {
		goto L458
	}
L458:
	;
	v1324 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L1
	} else {
		goto L459
	}
L459:
	;
	v3524 = v1324
	goto L7
L460:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L497
	}
L461:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)) = uint8(v1429)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(3)
	v1437 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L496
	}
L462:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+936)) = uint8(v1427)
	v1429 = v1427
	goto L461
L463:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v1427 = v1425
	goto L462
L464:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1421 == int32(1) {
		goto L460
	} else {
		goto L495
	}
L465:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1386 = F_pnstrdup(m, v1384, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L482
	}
L466:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1332 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v1331)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v1335
	v1338 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v1338
	v1346 = F_DirectInputFunctionCallSafe(m, int32(1408), v1332, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L469
	}
L468:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1427 = base.B2i32(v1381 != int32(0))
	goto L462
L469:
	;
	if v1346 != 0 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v1348&int32(1) == int32(0) {
		goto L468
	} else {
		goto L473
	}
L471:
	;
	goto L472
L472:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1353 == int32(0) {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	goto L472
L474:
	;
	v3524 = int32(2)
	goto L7
L475:
	;
	goto L476
L476:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1365 = F_jspOperationName(m, v1364)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L479
	}
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+552)) = int32(_a_F_executeItemOptUnwrapTarget_26)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+548)) = v1365
	*(*int32)(unsafe.Add(mBase, uint32(v24)+544)) = v1332
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(544))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1357), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L482:
	;
	v1390 = F_strlen(m, v1386)
	mBase = m.M
	v1391 = F_parse_bool_with_len(m, v1386, v1390, v24+int32(936))
	mBase = m.M
	goto L483
L483:
	;
	if v1391 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+936)))
	v1429 = v1392
	goto L461
L485:
	;
	goto L486
L486:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1393 != int32(1) {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v3524 = int32(2)
	goto L7
L488:
	;
	goto L489
L489:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L491
	}
L491:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1405 = F_jspOperationName(m, v1404)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+568)) = int32(_a_F_executeItemOptUnwrapTarget_26)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+564)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v24)+560)) = v1386
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(560))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1377), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	v3524 = int32(2)
	goto L7
L496:
	;
	v3524 = v1437
	goto L7
L497:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1447 = F_jspOperationName(m, v1446)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+528)) = v1447
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_27), v24+int32(528))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1386), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L502:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1469 - int32(1) {
	case 0:
		goto L510
	case 1:
		goto L511
	default:
		goto L509
	}
L503:
	;
	v1462 = F_JsonbType(m, l2)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	if v1462 != int32(16) {
		goto L502
	} else {
		goto L505
	}
L505:
	;
	v1467 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	v3524 = v1467
	goto L7
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v1599 = F_DirectFunctionCall1Coll(m, int32(1399), int32(0), v1593)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L1
	} else {
		goto L545
	}
L508:
	;
	v1590 = F_Int64GetDatum(m, v1475)
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L1
	} else {
		goto L544
	}
L509:
	;
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1566 != int32(1) {
		v3524 = int32(2)
		goto L7
	} else {
		goto L538
	}
L510:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1515 = F_pnstrdup(m, v1513, v1514)
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L1
	} else {
		goto L523
	}
L511:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1475 = F_numeric_int8_opt_error(m, v1472, v24+int32(952))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+952)))
	if v1477 != int32(1) {
		goto L508
	} else {
		goto L513
	}
L513:
	;
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1480 == int32(0) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v3524 = int32(2)
	goto L7
L515:
	;
	goto L516
L516:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1494 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v1493)
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1497 = F_jspOperationName(m, v1496)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+504)) = int32(_a_F_executeItemOptUnwrapTarget_28)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+500)) = v1497
	*(*int32)(unsafe.Add(mBase, uint32(v24)+496)) = v1494
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(496))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1283), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L523:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v1518
	v1521 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v1521
	v1529 = F_DirectInputFunctionCallSafe(m, int32(546), v1515, int32(-1), v24+int32(952), v24+int32(908))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L525
	}
L524:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v24)+908))
	v1593 = v1564
	goto L507
L525:
	;
	if v1529 != 0 {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v1531&int32(1) == int32(0) {
		goto L524
	} else {
		goto L529
	}
L527:
	;
	goto L528
L528:
	;
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1536 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L529:
	;
	goto L528
L530:
	;
	v3524 = int32(2)
	goto L7
L531:
	;
	goto L532
L532:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1548 = F_jspOperationName(m, v1547)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+520)) = int32(_a_F_executeItemOptUnwrapTarget_28)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+516)) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v24)+512)) = v1515
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(512))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1305), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L538:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L1
	} else {
		goto L539
	}
L539:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1577 = F_jspOperationName(m, v1576)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+480)) = v1577
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_16), v24+int32(480))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1313), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+908)) = v1590
	v1593 = v1590
	goto L507
L545:
	;
	v1601 = F_pg_detoast_datum(m, v1599)
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v1601
	v1608 = F_executeNextItem(m, l0, l1, int32(0), v24+int32(752), l3, int32(1))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L1
	} else {
		goto L547
	}
L547:
	;
	v3524 = v1608
	goto L7
L548:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(0) <= v1614 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v1617 = int32(0)
	if base.B2i32(l3 == v1617)&(v1612^int32(1)) != 0 {
		v3524 = v1617
		goto L7
	} else {
		goto L552
	}
L550:
	;
	goto L551
L551:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L1
	} else {
		goto L559
	}
L552:
	;
	if v1612 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v1628 = F_palloc(m, int32(20))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L1
	} else {
		goto L556
	}
L554:
	;
	v1630 = v24 + int32(952)
	goto L555
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1630))) = int32(2)
	v1636 = F_int64_to_numeric(m, base.I64_extend_i32_s(v1614-int32(1)))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L557
	}
L556:
	;
	v1630 = v1628
	goto L555
L557:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1630)+4)) = v1636
	v1641 = F_executeNextItem(m, l0, l1, v24+int32(752), v1630, l3, v1612)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L1
	} else {
		goto L558
	}
L558:
	;
	v3524 = v1641
	goto L7
L559:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_29), int32(0))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L560
	}
L560:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1241), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L562:
	;
	v1665 = m.G0
	v1667 = v1665 - int32(224)
	m.G0 = v1667
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1669 == int32(18) {
		goto L571
	} else {
		goto L572
	}
L563:
	;
	v1658 = F_JsonbType(m, l2)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	if v1658 != int32(16) {
		goto L562
	} else {
		goto L565
	}
L565:
	;
	v1663 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	v3524 = v1663
	goto L7
L567:
	;
	v3524 = v1950
	goto L7
L568:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L643
	}
L569:
	;
	m.G0 = v1667 + int32(224)
	goto L567
L570:
	;
	if v1673&int32(268435455) == int32(0) {
		goto L584
	} else {
		goto L585
	}
L571:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1672)))
	if v1673&int32(536870912) != 0 {
		goto L570
	} else {
		goto L574
	}
L572:
	;
	goto L573
L573:
	;
	v1682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v1682 != int32(1) {
		goto L576
	} else {
		goto L577
	}
L574:
	;
	if v1673&int32(1073741824) == int32(0) {
		goto L568
	} else {
		goto L575
	}
L575:
	;
	goto L573
L576:
	;
	v1950 = int32(2)
	goto L569
L577:
	;
	goto L578
L578:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	F_errcode(m, int32(319553666))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1694 = F_jspOperationName(m, v1693)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+16)) = v1694
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_30), v1667+int32(16))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2840), int32(_a_F_executeItemOptUnwrapTarget_31))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L584:
	;
	v1950 = int32(1)
	goto L569
L585:
	;
	goto L586
L586:
	;
	v1714 = F_jspGetNext(m, l1, v1667+int32(188))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+116)) = int32(_a_F_executeItemOptUnwrapTarget_32)
	*(*int64)(unsafe.Add(mBase, uint32(v1667)+108)) = int64(12884901889)
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+96)) = int32(_a_F_executeItemOptUnwrapTarget_33)
	*(*int64)(unsafe.Add(mBase, uint32(v1667)+88)) = int64(21474836481)
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+76)) = int32(_a_F_executeItemOptUnwrapTarget_34)
	*(*int64)(unsafe.Add(mBase, uint32(v1667)+68)) = int64(8589934593)
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1728 == int32(18) {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1734 = base.I64_extend_i32_s(v1672 - v1731)
	goto L590
L589:
	;
	v1734 = int64(0)
	goto L590
L590:
	;
	v1735 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+128)) = int32(2)
	v1741 = F_int64_to_numeric(m, v1735*int64(10000000000)+v1734)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+132)) = v1741
	v1744 = F_JsonbIteratorInit(m, v1672)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+64)) = v1744
	v1747 = int32(1)
	v1753 = F_JsonbIteratorNext(m, v1667-int32(-64), v1667+int32(168), v1747)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	if v1753 == int32(0) {
		v1950 = v1747
		goto L569
	} else {
		goto L594
	}
L594:
	;
	v1762 = v1747
	v1764 = v1753
	goto L595
L595:
	;
	if v1764 != int32(1) {
		v1933 = v1762
		goto L598
	} else {
		goto L599
	}
L596:
	;
	v1950 = v1945
	goto L569
L597:
	;
	goto L596
L598:
	;
	v1942 = F_JsonbIteratorNext(m, v1667-int32(-64), v1667+int32(168), int32(1))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L641
	}
L599:
	;
	v1783 = int32(0)
	if base.B2i32(l3 != int32(0))|v1714 == v1783 {
		v1945 = v1783
		goto L597
	} else {
		goto L600
	}
L600:
	;
	v1789 = v1667 + int32(148)
	v1791 = F_JsonbIteratorNext(m, v1667-int32(-64), v1789, int32(1))
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	v1793 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+40)) = v1793
	v1796 = v1667 + int32(40)
	v1799 = F_pushJsonbValue(m, v1796, int32(6), v1793)
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	v1804 = F_pushJsonbValue(m, v1796, int32(1), v1667+int32(108))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L1
	} else {
		goto L603
	}
L603:
	;
	v1809 = F_pushJsonbValue(m, v1796, int32(2), v1667+int32(168))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	v1814 = F_pushJsonbValue(m, v1796, int32(1), v1667+int32(88))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	v1817 = F_pushJsonbValue(m, v1796, int32(2), v1789)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	v1822 = F_pushJsonbValue(m, v1796, int32(1), v1667+int32(68))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	v1827 = F_pushJsonbValue(m, v1796, int32(2), v1667+int32(128))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	v1831 = F_pushJsonbValue(m, v1796, int32(7), int32(0))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	v1833 = F_JsonbValueToJsonb(m, v1831)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+44)) = int32(18)
	v1838 = v1833 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+52)) = v1838
	v1840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1833))))
	if v1840 == int32(1) {
		goto L612
	} else {
		goto L613
	}
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+48)) = v1869
	v1871 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1838
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1873
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1873 + int32(1)
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1878 <= int32(0) {
		goto L623
	} else {
		goto L624
	}
L612:
	;
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1833)+1)))
	if v1846 == int32(18) {
		goto L615
	} else {
		goto L616
	}
L613:
	;
	goto L614
L614:
	;
	v1857 = int32(1)
	if v1840&v1857 != 0 {
		v1869 = int32(base.Ui32(v1840)>>(uint(v1857)%32)) - v1857
		goto L611
	} else {
		goto L621
	}
L615:
	;
	v1849 = int32(16)
	goto L617
L616:
	;
	v1849 = int32(0)
	goto L617
L617:
	;
	if base.Ui32((v1846-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v1856 = int32(4)
	goto L620
L619:
	;
	v1856 = v1849
	goto L620
L620:
	;
	v1869 = v1856
	goto L611
L621:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1833)))
	v1869 = int32(base.Ui32(v1863)>>(uint(int32(2))%32)) - int32(4)
	goto L611
L622:
	;
	if l3 != 0 {
		v1933 = v1928
		goto L598
	} else {
		goto L639
	}
L623:
	;
	if l3 == int32(0) {
		goto L626
	} else {
		goto L627
	}
L624:
	;
	goto L625
L625:
	;
	v1921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v1922 = F_executeItemOptUnwrapTarget(m, l0, v1667+int32(188), v1667+int32(44), l3, v1921)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L1
	} else {
		goto L637
	}
L626:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1871
	v1928 = int32(0)
	goto L622
L627:
	;
	v1884 = F_palloc(m, int32(20))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1667)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1884)+16)) = v1886
	v1888 = *(*int64)(unsafe.Add(mBase, uint32(v1667)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v1884)+8)) = v1888
	v1890 = *(*int64)(unsafe.Add(mBase, uint32(v1667)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v1884))) = v1890
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1892 != 0 {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+216)) = v1884
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+220)) = v1892
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+36)) = v1892
	*(*int32)(unsafe.Add(mBase, uint32(v1667)+32)) = v1884
	v1901 = F_list_make2_impl(m, v1667+int32(36), v1667+int32(32))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L1
	} else {
		goto L632
	}
L630:
	;
	goto L631
L631:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1906 == int32(0) {
		goto L633
	} else {
		goto L634
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v1901
	goto L626
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1884
	goto L626
L634:
	;
	goto L635
L635:
	;
	v1910 = F_lappend(m, v1906, v1884)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v1910
	goto L626
L637:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v1871
	v1925 = int32(2)
	if v1922 == v1925 {
		v1945 = v1925
		goto L597
	} else {
		goto L638
	}
L638:
	;
	v1928 = v1922
	goto L622
L639:
	;
	v1930 = int32(0)
	if v1928 == v1930 {
		v1945 = v1930
		goto L597
	} else {
		goto L640
	}
L640:
	;
	v1933 = v1928
	goto L598
L641:
	;
	if v1942 != 0 {
		v1762 = v1933
		v1764 = v1942
		goto L595
	} else {
		goto L642
	}
L642:
	;
	v1950 = v1933
	goto L569
L643:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1672)))
	*(*int32)(unsafe.Add(mBase, uint32(v1667))) = v1976
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_35), v1667)
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(3629), int32(_a_F_executeItemOptUnwrapTarget_36))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L646:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2073 = F_cstring_to_text_with_len(m, v2071, v2072)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L1
	} else {
		goto L671
	}
L647:
	;
	v2046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2046 != int32(1) {
		goto L663
	} else {
		goto L664
	}
L648:
	;
	v2036 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2036
	*(*int32)(unsafe.Add(mBase, uint32(v24)+984)) = int32(0)
	if v1986 == int32(1) {
		goto L646
	} else {
		goto L662
	}
L649:
	;
	switch v1986 - int32(16) {
	case 0:
		goto L651
	default:
		goto L648
	case 2:
		goto L652
	}
L650:
	;
	v2029 = int32(1)
	v2032 = int32(0)
	v2034 = F_executeAnyItem(m, l0, l1, v1991, l3, v2029, v2029, v2029, v2032, v2032)
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L661
	}
L651:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L1
	} else {
		goto L658
	}
L652:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1991)))
	if v1992&int32(536870912) != 0 {
		goto L647
	} else {
		goto L653
	}
L653:
	;
	if v1992&int32(1073741824) != 0 {
		goto L650
	} else {
		goto L654
	}
L654:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1991)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+448)) = v2001
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_35), v24+int32(448))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(3629), int32(_a_F_executeItemOptUnwrapTarget_36))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L658:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+464)) = v2017
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_37), v24+int32(464))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1680), int32(_a_F_executeItemOptUnwrapTarget_38))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L661:
	;
	v3524 = v2034
	goto L7
L662:
	;
	goto L647
L663:
	;
	v3524 = int32(2)
	goto L7
L664:
	;
	goto L665
L665:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L666
	}
L666:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L1
	} else {
		goto L667
	}
L667:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2058 = F_jspOperationName(m, v2057)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+432)) = v2058
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_39), v24+int32(432))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2357), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L671:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2075 - int32(37) {
	case 0:
		goto L678
	default:
		goto L677
	case 8:
		v2143 = v2036
		goto L676
	}
L672:
	;
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v2346 - int32(37) {
	case 0:
		v2938 = v2332
		goto L745
	default:
		goto L753
	case 8:
		goto L758
	case 13:
		goto L757
	case 14:
		goto L756
	case 15:
		goto L755
	case 16:
		goto L754
	}
L673:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2307 = m.ExcPending
	if v2307 != 0 {
		goto L1
	} else {
		goto L737
	}
L674:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L1
	} else {
		goto L733
	}
L675:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L1
	} else {
		goto L730
	}
L676:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	v2147 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	v2151 = int32(0)
	goto L698
L677:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2120 == int32(0) {
		v2143 = v2036
		goto L676
	} else {
		goto L691
	}
L678:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2078 == int32(0) {
		v2143 = v2036
		goto L676
	} else {
		goto L679
	}
L679:
	;
	v2082 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v2082
	v2085 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2085
	v2088 = v24 + int32(952)
	F_jspGetArg(m, l1, v2088)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	if v2091 != int32(1) {
		goto L675
	} else {
		goto L681
	}
L681:
	;
	v2095 = v24 + int32(936)
	if v2095 != 0 {
		goto L683
	} else {
		goto L684
	}
L682:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v24)+936))
	v2100 = F_cstring_to_text_with_len(m, v2098, v2099)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L1
	} else {
		goto L686
	}
L683:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2088)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2095))) = v2096
	goto L685
L684:
	;
	goto L685
L685:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2088)+12))
	goto L682
L686:
	;
	v2111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2111 != 0 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v2112 = int32(0)
	goto L689
L688:
	;
	v2112 = v24 + int32(752)
	goto L689
L689:
	;
	v2113 = F_parse_datetime(m, v2073, v2100, v24+int32(900), v24+int32(988), v24+int32(984), v2112)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	v2115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	v2327 = v2115 << (uint(int32(1)) % 32) & int32(2)
	v2332 = v2113
	v2333 = v2036
	goto L672
L691:
	;
	v2124 = v24 + int32(952)
	F_jspGetArg(m, l1, v2124)
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L1
	} else {
		goto L692
	}
L692:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v24)+952))
	if v2127 != int32(2) {
		goto L674
	} else {
		goto L693
	}
L693:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2124)+12))
	v2133 = F_numeric_int4_opt_error(m, v2130, v24+int32(752))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	v2135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+752)))
	if v2135 != int32(1) {
		v2143 = v2133
		goto L676
	} else {
		goto L695
	}
L695:
	;
	v2138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2138 == int32(1) {
		goto L673
	} else {
		goto L696
	}
L696:
	;
	v3524 = int32(2)
	goto L7
L697:
	;
	v2209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v2210 == int32(37) {
		goto L709
	} else {
		goto L710
	}
L698:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v2145
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2147
	v2173 = v2151 << (uint(int32(2)) % 32)
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+uint32(_c_F_executeItemOptUnwrapTarget[3])))
	if v2174 == int32(0) {
		goto L700
	} else {
		goto L701
	}
L699:
	;
	v2327 = int32(0)
	v2332 = v2201
	v2333 = v2143
	goto L672
L700:
	;
	v2177 = int32(_a_F_executeItemOptUnwrapTarget_41)
	v2178 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[4]))
	v2181 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[4])) = v2181
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+uint32(_c_F_executeItemOptUnwrapTarget[6])))
	v2186 = F_cstring_to_text(m, v2185)
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L1
	} else {
		goto L703
	}
L701:
	;
	v2191 = v2174
	goto L702
L702:
	;
	v2201 = F_parse_datetime(m, v2073, v2191, v24+int32(900), v24+int32(988), v24+int32(984), v24+int32(752))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L1
	} else {
		goto L704
	}
L703:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2173)+uint32(_c_F_executeItemOptUnwrapTarget[3]))) = v2186
	*(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[4])) = v2178
	v2191 = v2186
	goto L702
L704:
	;
	v2203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v2203 != 0 {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v2205 = v2151 + int32(1)
	if v2205 == int32(13) {
		goto L697
	} else {
		goto L708
	}
L706:
	;
	goto L707
L707:
	;
	goto L699
L708:
	;
	v2151 = v2205
	goto L698
L709:
	;
	if v2209&int32(1) == int32(0) {
		goto L712
	} else {
		goto L713
	}
L710:
	;
	goto L711
L711:
	;
	if v2209&int32(1) == int32(0) {
		goto L721
	} else {
		goto L722
	}
L712:
	;
	v3524 = int32(2)
	goto L7
L713:
	;
	goto L714
L714:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	v2225 = F_text_to_cstring(m, v2073)
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v2225
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = int32(_a_F_executeItemOptUnwrapTarget_42)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_43), v24+int32(160))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	F_errhint(m, int32(_a_F_executeItemOptUnwrapTarget_44), int32(0))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L1
	} else {
		goto L719
	}
L719:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2485), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L1
	} else {
		goto L720
	}
L720:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L721:
	;
	v3524 = int32(2)
	goto L7
L722:
	;
	goto L723
L723:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L1
	} else {
		goto L724
	}
L724:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L1
	} else {
		goto L725
	}
L725:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2257 = F_jspOperationName(m, v2256)
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L1
	} else {
		goto L726
	}
L726:
	;
	v2259 = F_text_to_cstring(m, v2073)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L1
	} else {
		goto L727
	}
L727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+180)) = v2259
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v2257
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_43), v24+int32(176))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L1
	} else {
		goto L728
	}
L728:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2490), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L1
	} else {
		goto L729
	}
L729:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L730:
	;
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_45), int32(0))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L1
	} else {
		goto L731
	}
L731:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2383), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2285 = m.ExcPending
	if v2285 != 0 {
		goto L1
	} else {
		goto L732
	}
L732:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L733:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2291 = F_jspOperationName(m, v2290)
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+416)) = v2291
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_46), v24+int32(416))
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2442), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
	} else {
		goto L736
	}
L736:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L737:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L1
	} else {
		goto L738
	}
L738:
	;
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2312 = F_jspOperationName(m, v2311)
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+400)) = v2312
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_47), v24+int32(400))
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2450), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L742:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3026 = m.ExcPending
	if v3026 != 0 {
		goto L1
	} else {
		goto L953
	}
L743:
	;
	v2996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2996 != int32(1) {
		goto L945
	} else {
		goto L946
	}
L744:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L1
	} else {
		goto L940
	}
L745:
	;
	F_pfree(m, v2073)
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L1
	} else {
		goto L928
	}
L746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1184)
	v2938 = v2934
	goto L745
L747:
	;
	v2929 = *(*int64)(unsafe.Add(mBase, uint32(v24)+888))
	v2930 = F_Int64GetDatum(m, v2929)
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L1
	} else {
		goto L927
	}
L748:
	;
	v3524 = int32(2)
	goto L7
L749:
	;
	if v2333 == int32(-1) {
		v2934 = v2902
		goto L746
	} else {
		goto L922
	}
L750:
	;
	v2899 = F_DirectFunctionCall1Coll(m, v2897, int32(0), v2332)
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L1
	} else {
		goto L921
	}
L751:
	;
	v2884 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[7]))
	v2886 = m.G0
	v2887 = int32(16)
	v2888 = v2886 - v2887
	m.G0 = v2888
	v2892 = F_DetermineTimeZoneOffsetInternal(m, v24+int32(752), v2884, v2888+int32(8))
	mBase = m.M
	m.G0 = v2888 + v2887
	goto L920
L752:
	;
	v2864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2864, int32(_a_F_executeItemOptUnwrapTarget_48), int32(_a_F_executeItemOptUnwrapTarget_49))
	mBase = m.M
	v2868 = m.ExcPending
	if v2868 != 0 {
		goto L1
	} else {
		goto L917
	}
L753:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2851 = m.ExcPending
	if v2851 != 0 {
		goto L1
	} else {
		goto L914
	}
L754:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2712 <= int32(1183) {
		goto L893
	} else {
		goto L894
	}
L755:
	;
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2589 <= int32(1183) {
		goto L853
	} else {
		goto L854
	}
L756:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2516 <= int32(1183) {
		goto L824
	} else {
		goto L825
	}
L757:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2401 <= int32(1183) {
		goto L787
	} else {
		goto L788
	}
L758:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	if v2349 <= int32(1183) {
		goto L764
	} else {
		goto L765
	}
L759:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1082)
	v2938 = v2397
	goto L745
L760:
	;
	v2395 = F_DirectFunctionCall1Coll(m, v2393, int32(0), v2332)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L1
	} else {
		goto L779
	}
L761:
	;
	if v2349 != int32(1114) {
		goto L742
	} else {
		goto L778
	}
L762:
	;
	v2384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2384, int32(_a_F_executeItemOptUnwrapTarget_49), int32(_a_F_executeItemOptUnwrapTarget_50))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L1
	} else {
		goto L777
	}
L763:
	;
	v2358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2358 != int32(1) {
		goto L769
	} else {
		goto L770
	}
L764:
	;
	switch v2349 - int32(1082) {
	case 0:
		v2397 = v2332
		goto L759
	case 1:
		goto L763
	default:
		goto L761
	}
L765:
	;
	goto L766
L766:
	;
	if v2349 == int32(1184) {
		goto L762
	} else {
		goto L767
	}
L767:
	;
	if v2349 != int32(1266) {
		goto L742
	} else {
		goto L768
	}
L768:
	;
	goto L763
L769:
	;
	v3524 = int32(2)
	goto L7
L770:
	;
	goto L771
L771:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L1
	} else {
		goto L772
	}
L772:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L773
	}
L773:
	;
	v2369 = F_text_to_cstring(m, v2073)
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L1
	} else {
		goto L774
	}
L774:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+228)) = v2369
	*(*int32)(unsafe.Add(mBase, uint32(v24)+224)) = int32(_a_F_executeItemOptUnwrapTarget_50)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_43), v24+int32(224))
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L1
	} else {
		goto L775
	}
L775:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2517), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L1
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
	v2393 = int32(1410)
	goto L760
L778:
	;
	v2393 = int32(1411)
	goto L760
L779:
	;
	v2397 = v2395
	goto L759
L780:
	;
	if v2333 != int32(-1) {
		goto L806
	} else {
		goto L807
	}
L781:
	;
	v2469 = F_DirectFunctionCall1Coll(m, v2467, int32(0), v2332)
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		goto L1
	} else {
		goto L805
	}
L782:
	;
	v2462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2462, v2461, int32(_a_F_executeItemOptUnwrapTarget_51))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L1
	} else {
		goto L804
	}
L783:
	;
	v2460 = int32(1414)
	v2461 = int32(_a_F_executeItemOptUnwrapTarget_49)
	goto L782
L784:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L1
	} else {
		goto L801
	}
L785:
	;
	if v2401 == int32(1114) {
		v2467 = int32(1413)
		goto L781
	} else {
		goto L800
	}
L786:
	;
	v2412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2412 != int32(1) {
		goto L792
	} else {
		goto L793
	}
L787:
	;
	switch v2401 - int32(1082) {
	case 0:
		goto L786
	case 1:
		v2472 = v2332
		goto L780
	default:
		goto L785
	}
L788:
	;
	goto L789
L789:
	;
	if v2401 == int32(1184) {
		goto L783
	} else {
		goto L790
	}
L790:
	;
	if v2401 != int32(1266) {
		goto L784
	} else {
		goto L791
	}
L791:
	;
	v2460 = int32(1412)
	v2461 = int32(_a_F_executeItemOptUnwrapTarget_52)
	goto L782
L792:
	;
	v3524 = int32(2)
	goto L7
L793:
	;
	goto L794
L794:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L1
	} else {
		goto L795
	}
L795:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L1
	} else {
		goto L796
	}
L796:
	;
	v2423 = F_text_to_cstring(m, v2073)
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L1
	} else {
		goto L797
	}
L797:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+260)) = v2423
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = int32(_a_F_executeItemOptUnwrapTarget_51)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_43), v24+int32(256))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L1
	} else {
		goto L798
	}
L798:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2545), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L1
	} else {
		goto L799
	}
L799:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L800:
	;
	goto L784
L801:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+240)) = v2446
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_53), v24+int32(240))
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L1
	} else {
		goto L802
	}
L802:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2566), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L1
	} else {
		goto L803
	}
L803:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L804:
	;
	v2467 = v2460
	goto L781
L805:
	;
	v2472 = v2469
	goto L780
L806:
	;
	v2477 = F_anytime_typmod_check(m, int32(0), v2333)
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L1
	} else {
		goto L809
	}
L807:
	;
	v2513 = v2472
	goto L808
L808:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1083)
	v2938 = v2513
	goto L745
L809:
	;
	v2479 = *(*int64)(unsafe.Add(mBase, uint32(v2472)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2479
	v2482 = v24 + int32(752)
	if base.Ui32(v2477) <= base.Ui32(int32(6)) {
		goto L811
	} else {
		goto L812
	}
L810:
	;
	v2508 = *(*int64)(unsafe.Add(mBase, uint32(v24)+752))
	v2509 = F_Int64GetDatum(m, v2508)
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L1
	} else {
		goto L818
	}
L811:
	;
	v2489 = v2477 << (uint(int32(3)) % 32)
	v2490 = *(*int64)(unsafe.Add(mBase, uint32(v2489)+uint32(_c_F_executeItemOptUnwrapTarget[8])))
	v2491 = *(*int64)(unsafe.Add(mBase, uint32(v2489)+uint32(_c_F_executeItemOptUnwrapTarget[9])))
	v2492 = *(*int64)(unsafe.Add(mBase, uint32(v2482)))
	if int64(0) <= v2492 {
		goto L815
	} else {
		goto L816
	}
L812:
	;
	goto L813
L813:
	;
	goto L810
L814:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2482))) = v2502
	goto L813
L815:
	;
	v2495 = v2491 + v2492
	v2496 = base.I64_rem_s(v2495, v2490)
	v2502 = v2495 - v2496
	goto L814
L816:
	;
	goto L817
L817:
	;
	v2498 = v2491 - v2492
	v2499 = base.I64_rem_s(v2498, v2490)
	v2502 = v2499 - v2498
	goto L814
L818:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2477
	v2513 = v2509
	goto L808
L819:
	;
	if v2333 != int32(-1) {
		goto L835
	} else {
		goto L836
	}
L820:
	;
	v2552 = F_DirectFunctionCall1Coll(m, v2550, int32(0), v2332)
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L1
	} else {
		goto L834
	}
L821:
	;
	v2544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2544, int32(_a_F_executeItemOptUnwrapTarget_51), int32(_a_F_executeItemOptUnwrapTarget_52))
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		goto L1
	} else {
		goto L833
	}
L822:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L830
	}
L823:
	;
	if v2516 == int32(1114) {
		goto L743
	} else {
		goto L829
	}
L824:
	;
	switch v2516 - int32(1082) {
	case 0:
		goto L743
	case 1:
		goto L821
	default:
		goto L823
	}
L825:
	;
	goto L826
L826:
	;
	if v2516 == int32(1184) {
		v2550 = int32(1415)
		goto L820
	} else {
		goto L827
	}
L827:
	;
	if v2516 != int32(1266) {
		goto L822
	} else {
		goto L828
	}
L828:
	;
	v2554 = v2332
	goto L819
L829:
	;
	goto L822
L830:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+272)) = v2532
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_53), v24+int32(272))
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2613), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L1
	} else {
		goto L832
	}
L832:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L833:
	;
	v2550 = int32(1416)
	goto L820
L834:
	;
	v2554 = v2552
	goto L819
L835:
	;
	v2558 = F_anytime_typmod_check(m, int32(1), v2333)
	mBase = m.M
	v2559 = m.ExcPending
	if v2559 != 0 {
		goto L1
	} else {
		goto L838
	}
L836:
	;
	goto L837
L837:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1266)
	v2938 = v2554
	goto L745
L838:
	;
	if base.Ui32(v2558) <= base.Ui32(int32(6)) {
		goto L840
	} else {
		goto L841
	}
L839:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2558
	goto L837
L840:
	;
	v2566 = v2558 << (uint(int32(3)) % 32)
	v2567 = *(*int64)(unsafe.Add(mBase, uint32(v2566)+uint32(_c_F_executeItemOptUnwrapTarget[8])))
	v2568 = *(*int64)(unsafe.Add(mBase, uint32(v2566)+uint32(_c_F_executeItemOptUnwrapTarget[9])))
	v2569 = *(*int64)(unsafe.Add(mBase, uint32(v2554)))
	if int64(0) <= v2569 {
		goto L844
	} else {
		goto L845
	}
L841:
	;
	goto L842
L842:
	;
	goto L839
L843:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2554))) = v2579
	goto L842
L844:
	;
	v2572 = v2568 + v2569
	v2573 = base.I64_rem_s(v2572, v2567)
	v2579 = v2572 - v2573
	goto L843
L845:
	;
	goto L846
L846:
	;
	v2575 = v2568 - v2569
	v2576 = base.I64_rem_s(v2575, v2567)
	v2579 = v2576 - v2575
	goto L843
L847:
	;
	if v2333 != int32(-1) {
		goto L872
	} else {
		goto L873
	}
L848:
	;
	v2653 = F_DirectFunctionCall1Coll(m, v2651, int32(0), v2332)
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L1
	} else {
		goto L871
	}
L849:
	;
	v2645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2645, int32(_a_F_executeItemOptUnwrapTarget_49), int32(_a_F_executeItemOptUnwrapTarget_48))
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L1
	} else {
		goto L870
	}
L850:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2632 = m.ExcPending
	if v2632 != 0 {
		goto L1
	} else {
		goto L867
	}
L851:
	;
	if v2589 == int32(1114) {
		v2655 = v2332
		goto L847
	} else {
		goto L866
	}
L852:
	;
	v2600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2600 != int32(1) {
		goto L858
	} else {
		goto L859
	}
L853:
	;
	switch v2589 - int32(1082) {
	case 0:
		v2651 = int32(1417)
		goto L848
	case 1:
		goto L852
	default:
		goto L851
	}
L854:
	;
	goto L855
L855:
	;
	if v2589 == int32(1184) {
		goto L849
	} else {
		goto L856
	}
L856:
	;
	if v2589 != int32(1266) {
		goto L850
	} else {
		goto L857
	}
L857:
	;
	goto L852
L858:
	;
	v3524 = int32(2)
	goto L7
L859:
	;
	goto L860
L860:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L1
	} else {
		goto L861
	}
L861:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L1
	} else {
		goto L862
	}
L862:
	;
	v2611 = F_text_to_cstring(m, v2073)
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+340)) = v2611
	*(*int32)(unsafe.Add(mBase, uint32(v24)+336)) = int32(_a_F_executeItemOptUnwrapTarget_48)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_43), v24+int32(336))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2649), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L1
	} else {
		goto L865
	}
L865:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L866:
	;
	goto L850
L867:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+304)) = v2633
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_53), v24+int32(304))
	mBase = m.M
	v2639 = m.ExcPending
	if v2639 != 0 {
		goto L1
	} else {
		goto L868
	}
L868:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2660), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L1
	} else {
		goto L869
	}
L869:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L870:
	;
	v2651 = int32(1418)
	goto L848
L871:
	;
	v2655 = v2653
	goto L847
L872:
	;
	v2660 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v2660
	v2663 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = v2663
	v2666 = F_anytimestamp_typmod_check(m, int32(0), v2333)
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L1
	} else {
		goto L875
	}
L873:
	;
	v2709 = v2655
	goto L874
L874:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+900)) = int32(1114)
	v2938 = v2709
	goto L745
L875:
	;
	v2668 = *(*int64)(unsafe.Add(mBase, uint32(v2655)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+936)) = v2668
	F_AdjustTimestampForTypmod(m, v24+int32(936), v2666, v24+int32(752))
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L1
	} else {
		goto L876
	}
L876:
	;
	v2676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+756)))
	if v2676 == int32(1) {
		goto L877
	} else {
		goto L878
	}
L877:
	;
	v2679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2679 == int32(0) {
		goto L880
	} else {
		goto L881
	}
L878:
	;
	goto L879
L879:
	;
	v2704 = *(*int64)(unsafe.Add(mBase, uint32(v24)+936))
	v2705 = F_Int64GetDatum(m, v2704)
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L1
	} else {
		goto L888
	}
L880:
	;
	v3524 = int32(2)
	goto L7
L881:
	;
	goto L882
L882:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2686 = m.ExcPending
	if v2686 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	v2690 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2691 = F_jspOperationName(m, v2690)
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+320)) = v2691
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_54), v24+int32(320))
	mBase = m.M
	v2698 = m.ExcPending
	if v2698 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2679), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2703 = m.ExcPending
	if v2703 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L888:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2666
	v2709 = v2705
	goto L874
L889:
	;
	v2764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+35)))
	F_checkTimezoneIsUsedForCast(m, v2764, int32(_a_F_executeItemOptUnwrapTarget_50), int32(_a_F_executeItemOptUnwrapTarget_49))
	mBase = m.M
	v2768 = m.ExcPending
	if v2768 != 0 {
		goto L1
	} else {
		goto L908
	}
L890:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L1
	} else {
		goto L905
	}
L891:
	;
	if v2712 == int32(1114) {
		goto L752
	} else {
		goto L904
	}
L892:
	;
	v2721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2721 != int32(1) {
		goto L748
	} else {
		goto L898
	}
L893:
	;
	switch v2712 - int32(1082) {
	case 0:
		goto L889
	case 1:
		goto L892
	default:
		goto L891
	}
L894:
	;
	goto L895
L895:
	;
	if v2712 == int32(1184) {
		v2902 = v2332
		goto L749
	} else {
		goto L896
	}
L896:
	;
	if v2712 != int32(1266) {
		goto L890
	} else {
		goto L897
	}
L897:
	;
	goto L892
L898:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L1
	} else {
		goto L899
	}
L899:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	v2731 = F_text_to_cstring(m, v2073)
	mBase = m.M
	v2732 = m.ExcPending
	if v2732 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+388)) = v2731
	*(*int32)(unsafe.Add(mBase, uint32(v24)+384)) = int32(_a_F_executeItemOptUnwrapTarget_55)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_43), v24+int32(384))
	mBase = m.M
	v2740 = m.ExcPending
	if v2740 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2720), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		goto L1
	} else {
		goto L903
	}
L903:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L904:
	;
	goto L890
L905:
	;
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+352)) = v2752
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_53), v24+int32(352))
	mBase = m.M
	v2758 = m.ExcPending
	if v2758 != 0 {
		goto L1
	} else {
		goto L906
	}
L906:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2741), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2763 = m.ExcPending
	if v2763 != 0 {
		goto L1
	} else {
		goto L907
	}
L907:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L908:
	;
	v2780 = v2332 + int32(_a_F_executeItemOptUnwrapTarget_56)
	v2781 = int32(_a_F_executeItemOptUnwrapTarget_57)
	v2782 = base.I32_div_u_s(v2780, v2781)
	v2783 = int32(3)
	v2789 = int32(2)
	v2794 = base.I32_div_u_s((v2782*int32(1073595727)+v2780)<<(uint(v2789)%32)|v2783, v2781)
	v2797 = v2332 + int32(_a_F_executeItemOptUnwrapTarget_58) + v2782*v2783 + v2794 + int32(_a_F_executeItemOptUnwrapTarget_59)
	v2798 = int32(1461)
	v2799 = base.I32_div_u_s(v2797, v2798)
	v2802 = v2799*int32(-1461) + v2797
	v2804 = v2802 << (uint(v2789) % 32)
	if base.Ui32(v2798) <= base.Ui32(v2804) {
		goto L911
	} else {
		goto L912
	}
L909:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+752)) = int64(0)
	v2880 = int32(1419)
	goto L751
L910:
	;
	v2817 = base.I32_div_u_s(v2804, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(772)))) = v2817 + v2799<<(uint(int32(2))%32) - int32(_a_F_executeItemOptUnwrapTarget_60)
	v2825 = v2815 + int32(123)
	v2829 = int32(base.Ui32(v2825*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(764)))) = v2825 - int32(base.Ui32(v2829*int32(_a_F_executeItemOptUnwrapTarget_61))>>(uint(int32(8))%32))
	v2839 = base.I32_rem_u_s(v2829+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(768)))) = v2839 + int32(1)
	goto L909
L911:
	;
	v2810 = base.I32_rem_u_s(v2802+int32(305), int32(365))
	v2815 = v2810
	goto L910
L912:
	;
	goto L913
L913:
	;
	v2814 = base.I32_rem_u_s(v2802+int32(306), int32(366))
	v2815 = v2814
	goto L910
L914:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = v2852
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_10), v24+int32(192))
	mBase = m.M
	v2858 = m.ExcPending
	if v2858 != 0 {
		goto L1
	} else {
		goto L915
	}
L915:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2771), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2863 = m.ExcPending
	if v2863 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L917:
	;
	v2869 = int32(1420)
	v2870 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v2871 = int32(0)
	v2878 = F_timestamp2tm(m, v2870, v2871, v24+int32(752), v24+int32(948), v2871, v2871)
	mBase = m.M
	v2879 = m.ExcPending
	if v2879 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	if v2878 != 0 {
		v2897 = v2869
		goto L750
	} else {
		goto L919
	}
L919:
	;
	v2880 = v2869
	goto L751
L920:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+984)) = v2892
	v2897 = v2880
	goto L750
L921:
	;
	v2902 = v2899
	goto L749
L922:
	;
	v2906 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+944)) = v2906
	v2909 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+936)) = v2909
	v2912 = F_anytimestamp_typmod_check(m, int32(1), v2333)
	mBase = m.M
	v2913 = m.ExcPending
	if v2913 != 0 {
		goto L1
	} else {
		goto L923
	}
L923:
	;
	v2914 = *(*int64)(unsafe.Add(mBase, uint32(v2902)))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+888)) = v2914
	F_AdjustTimestampForTypmod(m, v24+int32(888), v2912, v24+int32(936))
	mBase = m.M
	v2921 = m.ExcPending
	if v2921 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	v2922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+940)))
	if v2922 != int32(1) {
		goto L747
	} else {
		goto L925
	}
L925:
	;
	v2925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v2925 != 0 {
		goto L744
	} else {
		goto L926
	}
L926:
	;
	goto L748
L927:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+988)) = v2912
	v2934 = v2930
	goto L746
L928:
	;
	v2942 = int32(2)
	if v2327 == v2942 {
		v3524 = v2942
		goto L7
	} else {
		goto L929
	}
L929:
	;
	v2947 = F_jspGetNext(m, l1, v24+int32(952))
	mBase = m.M
	v2948 = m.ExcPending
	if v2948 != 0 {
		goto L1
	} else {
		goto L930
	}
L930:
	;
	if l3 == int32(0) {
		goto L931
	} else {
		goto L932
	}
L931:
	;
	if v2947 == int32(0) {
		v3524 = v2327
		goto L7
	} else {
		goto L934
	}
L932:
	;
	goto L933
L933:
	;
	if v2947 == int32(0) {
		goto L935
	} else {
		goto L936
	}
L934:
	;
	goto L933
L935:
	;
	v2959 = F_palloc(m, int32(20))
	mBase = m.M
	v2960 = m.ExcPending
	if v2960 != 0 {
		goto L1
	} else {
		goto L938
	}
L936:
	;
	v2961 = v24 + int32(908)
	goto L937
L937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2961)+4)) = v2938
	*(*int32)(unsafe.Add(mBase, uint32(v2961))) = int32(32)
	v2965 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v2961)+8)) = v2965
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v24)+988))
	*(*int32)(unsafe.Add(mBase, uint32(v2961)+12)) = v2967
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(v24)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v2961)+16)) = v2969
	v2973 = F_executeNextItem(m, l0, l1, v24+int32(952), v2961, l3, v2947)
	mBase = m.M
	v2974 = m.ExcPending
	if v2974 != 0 {
		goto L1
	} else {
		goto L939
	}
L938:
	;
	v2961 = v2959
	goto L937
L939:
	;
	v3524 = v2973
	goto L7
L940:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v2981 = m.ExcPending
	if v2981 != 0 {
		goto L1
	} else {
		goto L941
	}
L941:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v2983 = F_jspOperationName(m, v2982)
	mBase = m.M
	v2984 = m.ExcPending
	if v2984 != 0 {
		goto L1
	} else {
		goto L942
	}
L942:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v2983
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_54), v24+int32(368))
	mBase = m.M
	v2990 = m.ExcPending
	if v2990 != 0 {
		goto L1
	} else {
		goto L943
	}
L943:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2760), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L1
	} else {
		goto L944
	}
L944:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L945:
	;
	v3524 = int32(2)
	goto L7
L946:
	;
	goto L947
L947:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3003 = m.ExcPending
	if v3003 != 0 {
		goto L1
	} else {
		goto L948
	}
L948:
	;
	F_errcode(m, int32(17563778))
	mBase = m.M
	v3006 = m.ExcPending
	if v3006 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	v3007 = F_text_to_cstring(m, v2073)
	mBase = m.M
	v3008 = m.ExcPending
	if v3008 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+292)) = v3007
	*(*int32)(unsafe.Add(mBase, uint32(v24)+288)) = int32(_a_F_executeItemOptUnwrapTarget_62)
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_43), v24+int32(288))
	mBase = m.M
	v3016 = m.ExcPending
	if v3016 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2598), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v3021 = m.ExcPending
	if v3021 != 0 {
		goto L1
	} else {
		goto L952
	}
L952:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L953:
	;
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(v24)+900))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v3027
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapTarget_53), v24+int32(208))
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L954
	}
L954:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(2530), int32(_a_F_executeItemOptUnwrapTarget_40))
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L1
	} else {
		goto L955
	}
L955:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L956:
	;
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v3048 - int32(1) {
	case 0:
		goto L966
	case 1:
		goto L967
	default:
		goto L965
	}
L957:
	;
	v3041 = F_JsonbType(m, l2)
	mBase = m.M
	v3042 = m.ExcPending
	if v3042 != 0 {
		goto L1
	} else {
		goto L958
	}
L958:
	;
	if v3041 != int32(16) {
		goto L956
	} else {
		goto L959
	}
L959:
	;
	v3046 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3047 = m.ExcPending
	if v3047 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	v3524 = v3046
	goto L7
L961:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3238 = m.ExcPending
	if v3238 != 0 {
		goto L1
	} else {
		goto L1011
	}
L962:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3217 = m.ExcPending
	if v3217 != 0 {
		goto L1
	} else {
		goto L1006
	}
L963:
	;
	v3212 = F_executeNextItem(m, l0, l1, int32(0), v3206, l3, int32(1))
	mBase = m.M
	v3213 = m.ExcPending
	if v3213 != 0 {
		goto L1
	} else {
		goto L1005
	}
L964:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = int32(2)
	v3197 = F_Float8GetDatum(m, v3124)
	mBase = m.M
	v3198 = m.ExcPending
	if v3198 != 0 {
		goto L1
	} else {
		goto L1002
	}
L965:
	;
	v3169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3169 != int32(1) {
		v3524 = int32(2)
		goto L7
	} else {
		goto L996
	}
L966:
	;
	v3110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3112 = F_pnstrdup(m, v3110, v3111)
	mBase = m.M
	v3113 = m.ExcPending
	if v3113 != 0 {
		goto L1
	} else {
		goto L982
	}
L967:
	;
	v3053 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3054 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v3053)
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	v3057 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v3057
	v3060 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v3060
	v3066 = F_float8in_internal(m, v3054, int32(0), int32(_a_F_executeItemOptUnwrapTarget_63), v3054, v24+int32(952))
	mBase = m.M
	v3067 = m.ExcPending
	if v3067 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	v3068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v3068 == int32(1) {
		goto L971
	} else {
		goto L972
	}
L970:
	;
	v3524 = int32(2)
	goto L7
L971:
	;
	v3071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3071 != int32(1) {
		goto L970
	} else {
		goto L974
	}
L972:
	;
	goto L973
L973:
	;
	v3098 = base.F64_abs(v3066)
	if base.F64_ne(v3098, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3098)) < base.Ui64(int64(9218868437227405313))) != 0 {
		v3206 = l2
		goto L963
	} else {
		goto L980
	}
L974:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L1
	} else {
		goto L975
	}
L975:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L1
	} else {
		goto L976
	}
L976:
	;
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3082 = F_jspOperationName(m, v3081)
	mBase = m.M
	v3083 = m.ExcPending
	if v3083 != 0 {
		goto L1
	} else {
		goto L977
	}
L977:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = int32(_a_F_executeItemOptUnwrapTarget_63)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v3082
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v3054
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(96))
	mBase = m.M
	v3092 = m.ExcPending
	if v3092 != 0 {
		goto L1
	} else {
		goto L978
	}
L978:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1166), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L1
	} else {
		goto L979
	}
L979:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L980:
	;
	v3105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3105 == int32(1) {
		goto L962
	} else {
		goto L981
	}
L981:
	;
	goto L970
L982:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+960)) = v3115
	v3118 = *(*int64)(unsafe.Add(mBase, _c_F_executeItemOptUnwrapTarget[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+952)) = v3118
	v3124 = F_float8in_internal(m, v3112, int32(0), int32(_a_F_executeItemOptUnwrapTarget_63), v3112, v24+int32(952))
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L1
	} else {
		goto L983
	}
L983:
	;
	v3126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+956)))
	if v3126 == int32(1) {
		goto L985
	} else {
		goto L986
	}
L984:
	;
	v3524 = int32(2)
	goto L7
L985:
	;
	v3129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3129 != int32(1) {
		goto L984
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	v3156 = base.F64_abs(v3124)
	if base.F64_ne(v3156, math.Float64frombits(uint64(0x7ff0000000000000)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v3156)) < base.Ui64(int64(9218868437227405313))) != 0 {
		goto L964
	} else {
		goto L994
	}
L988:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3135 = m.ExcPending
	if v3135 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3138 = m.ExcPending
	if v3138 != 0 {
		goto L1
	} else {
		goto L990
	}
L990:
	;
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3140 = F_jspOperationName(m, v3139)
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = int32(_a_F_executeItemOptUnwrapTarget_63)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v3140
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v3112
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_15), v24+int32(128))
	mBase = m.M
	v3150 = m.ExcPending
	if v3150 != 0 {
		goto L1
	} else {
		goto L992
	}
L992:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1192), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3155 = m.ExcPending
	if v3155 != 0 {
		goto L1
	} else {
		goto L993
	}
L993:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L994:
	;
	v3163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3163 == int32(1) {
		goto L961
	} else {
		goto L995
	}
L995:
	;
	goto L984
L996:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3175 = m.ExcPending
	if v3175 != 0 {
		goto L1
	} else {
		goto L997
	}
L997:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3178 = m.ExcPending
	if v3178 != 0 {
		goto L1
	} else {
		goto L998
	}
L998:
	;
	v3179 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3180 = F_jspOperationName(m, v3179)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L1
	} else {
		goto L999
	}
L999:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v3180
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_16), v24+int32(80))
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L1
	} else {
		goto L1000
	}
L1000:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1210), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3192 = m.ExcPending
	if v3192 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1001:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1002:
	;
	v3199 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v3197)
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1003:
	;
	v3201 = F_pg_detoast_datum(m, v3199)
	mBase = m.M
	v3202 = m.ExcPending
	if v3202 != 0 {
		goto L1
	} else {
		goto L1004
	}
L1004:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+756)) = v3201
	v3206 = v24 + int32(752)
	goto L963
L1005:
	;
	v3524 = v3212
	goto L7
L1006:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3220 = m.ExcPending
	if v3220 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1007:
	;
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3222 = F_jspOperationName(m, v3221)
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1008:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v3222
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_20), v24+int32(112))
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1171), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3234 = m.ExcPending
	if v3234 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1010:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1011:
	;
	F_errcode(m, int32(101449858))
	mBase = m.M
	v3241 = m.ExcPending
	if v3241 != 0 {
		goto L1
	} else {
		goto L1012
	}
L1012:
	;
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3243 = F_jspOperationName(m, v3242)
	mBase = m.M
	v3244 = m.ExcPending
	if v3244 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1013:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v3243
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_20), v24+int32(144))
	mBase = m.M
	v3250 = m.ExcPending
	if v3250 != 0 {
		goto L1
	} else {
		goto L1014
	}
L1014:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1197), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3255 = m.ExcPending
	if v3255 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1016:
	;
	v3524 = v3257
	goto L7
L1017:
	;
	v3524 = v3260
	goto L7
L1018:
	;
	v3524 = v3263
	goto L7
L1019:
	;
	v3308 = F_palloc(m, int32(20))
	mBase = m.M
	v3309 = m.ExcPending
	if v3309 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1020:
	;
	v3277 = int32(1)
	v3278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3278 != 0 {
		v3306 = v3277
		goto L1019
	} else {
		goto L1023
	}
L1021:
	;
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v3268)))
	if v3269&int32(1342177280) != int32(1073741824) {
		goto L1020
	} else {
		goto L1022
	}
L1022:
	;
	v3306 = v3269 & int32(268435455)
	goto L1019
L1023:
	;
	v3279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3279 != 0 {
		v3524 = v3277
		goto L7
	} else {
		goto L1024
	}
L1024:
	;
	v3280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3280 != int32(1) {
		goto L1025
	} else {
		goto L1026
	}
L1025:
	;
	v3524 = int32(2)
	goto L7
L1026:
	;
	goto L1027
L1027:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1028:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3290 = m.ExcPending
	if v3290 != 0 {
		goto L1
	} else {
		goto L1029
	}
L1029:
	;
	v3291 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3292 = F_jspOperationName(m, v3291)
	mBase = m.M
	v3293 = m.ExcPending
	if v3293 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v3292
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_64), v24-int32(-64))
	mBase = m.M
	v3299 = m.ExcPending
	if v3299 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1113), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3304 = m.ExcPending
	if v3304 != 0 {
		goto L1
	} else {
		goto L1032
	}
L1032:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3308))) = int32(2)
	v3313 = F_int64_to_numeric(m, base.I64_extend_i32_u(v3306))
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L1
	} else {
		goto L1034
	}
L1034:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3308)+4)) = v3313
	v3316 = int32(0)
	v3318 = F_executeNextItem(m, l0, l1, v3316, v3308, l3, v3316)
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	v3524 = v3318
	goto L7
L1036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3321))) = int32(1)
	v3325 = F_JsonbTypeName(m, l2)
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1037:
	;
	v3327 = F_pstrdup(m, v3325)
	mBase = m.M
	v3328 = m.ExcPending
	if v3328 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1038:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3321)+8)) = v3327
	v3330 = F_strlen(m, v3327)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3321)+4)) = v3330
	v3332 = int32(0)
	v3334 = F_executeNextItem(m, l0, l1, v3332, v3321, l3, v3332)
	mBase = m.M
	v3335 = m.ExcPending
	if v3335 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1039:
	;
	v3524 = v3334
	goto L7
L1040:
	;
	v3358 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3359 = m.ExcPending
	if v3359 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1041:
	;
	v3336 = F_JsonbType(m, l2)
	mBase = m.M
	v3337 = m.ExcPending
	if v3337 != 0 {
		goto L1
	} else {
		goto L1044
	}
L1042:
	;
	goto L1043
L1043:
	;
	v3341 = v24 + int32(752)
	F_jspGetArg(m, l1, v3341)
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L1
	} else {
		goto L1046
	}
L1044:
	;
	if v3336 == int32(16) {
		goto L1040
	} else {
		goto L1045
	}
L1045:
	;
	goto L1043
L1046:
	;
	v3344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
	v3347 = F_executeBoolItem(m, l0, v3341, l2, int32(0))
	mBase = m.M
	v3348 = m.ExcPending
	if v3348 != 0 {
		goto L1
	} else {
		goto L1047
	}
L1047:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3344
	v3350 = int32(1)
	if v3347 != v3350 {
		v3524 = v3350
		goto L7
	} else {
		goto L1048
	}
L1048:
	;
	v3355 = F_executeNextItem(m, l0, l1, int32(0), l2, l3, int32(1))
	mBase = m.M
	v3356 = m.ExcPending
	if v3356 != 0 {
		goto L1
	} else {
		goto L1049
	}
L1049:
	;
	v3524 = v3355
	goto L7
L1050:
	;
	v3524 = v3358
	goto L7
L1051:
	;
	v3366 = *(*int32)(unsafe.Add(mBase, uint32(v3362)+8))
	v3367 = v3366
	goto L1053
L1052:
	;
	v3367 = int32(0)
	goto L1053
L1053:
	;
	v3368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v3368
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3367
	v3373 = F_executeNextItem(m, l0, l1, v3368, v3362, l3, int32(1))
	mBase = m.M
	v3374 = m.ExcPending
	if v3374 != 0 {
		goto L1
	} else {
		goto L1054
	}
L1054:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v3360
	v3524 = v3373
	goto L7
L1055:
	;
	v3524 = v3379
	goto L7
L1056:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1057:
	;
	if v3381 == int32(17) {
		goto L1058
	} else {
		goto L1059
	}
L1058:
	;
	v3385 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+752)) = v3385
	v3389 = v24 + int32(756)
	if v3389 != 0 {
		goto L1062
	} else {
		goto L1063
	}
L1059:
	;
	goto L1060
L1060:
	;
	if l4 == int32(0) {
		goto L1077
	} else {
		goto L1078
	}
L1061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+760)) = v3392
	v3394 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3398 = F_findJsonbValueFromContainer(m, v3394, int32(536870912), v24+int32(752))
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1062:
	;
	v3390 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3389))) = v3390
	goto L1064
L1063:
	;
	goto L1064
L1064:
	;
	v3392 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	goto L1061
L1065:
	;
	if v3398 != 0 {
		goto L1066
	} else {
		goto L1067
	}
L1066:
	;
	v3400 = int32(0)
	v3402 = F_executeNextItem(m, l0, l1, v3400, v3398, l3, v3400)
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1067:
	;
	goto L1068
L1068:
	;
	v3409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3409 != 0 {
		v3524 = v3385
		goto L7
	} else {
		goto L1075
	}
L1069:
	;
	if l3 != 0 {
		goto L1070
	} else {
		goto L1071
	}
L1070:
	;
	v3404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v3404 <= int32(0) {
		v3524 = v3402
		goto L7
	} else {
		goto L1073
	}
L1071:
	;
	goto L1072
L1072:
	;
	F_pfree(m, v3398)
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L1
	} else {
		goto L1074
	}
L1073:
	;
	goto L1072
L1074:
	;
	v3524 = v3402
	goto L7
L1075:
	;
	v3410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3410 == int32(1) {
		goto L1056
	} else {
		goto L1076
	}
L1076:
	;
	v3524 = int32(2)
	goto L7
L1077:
	;
	v3423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v3423 != 0 {
		goto L1082
	} else {
		goto L1083
	}
L1078:
	;
	v3416 = F_JsonbType(m, l2)
	mBase = m.M
	v3417 = m.ExcPending
	if v3417 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1079:
	;
	if v3416 != int32(16) {
		goto L1077
	} else {
		goto L1080
	}
L1080:
	;
	v3421 = F_executeItemUnwrapTargetArray(m, l0, l1, l2, l3, int32(0))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1081:
	;
	v3524 = v3421
	goto L7
L1082:
	;
	v3524 = int32(1)
	goto L7
L1083:
	;
	goto L1084
L1084:
	;
	v3425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3425 != int32(1) {
		goto L1085
	} else {
		goto L1086
	}
L1085:
	;
	v3524 = int32(2)
	goto L7
L1086:
	;
	goto L1087
L1087:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3432 = m.ExcPending
	if v3432 != 0 {
		goto L1
	} else {
		goto L1088
	}
L1088:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3435 = m.ExcPending
	if v3435 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1089:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_65), int32(0))
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1090:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1054), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3444 = m.ExcPending
	if v3444 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1092:
	;
	F_errcode(m, int32(285999234))
	mBase = m.M
	v3451 = m.ExcPending
	if v3451 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1093:
	;
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v24)+760))
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v24)+756))
	v3454 = F_pnstrdup(m, v3452, v3453)
	mBase = m.M
	v3455 = m.ExcPending
	if v3455 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1094:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v3454
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_66), v24+int32(48))
	mBase = m.M
	v3461 = m.ExcPending
	if v3461 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(1044), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1096:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1097:
	;
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v3472 != 0 {
		v3482 = int32(1)
		goto L1098
	} else {
		goto L1099
	}
L1098:
	;
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v3484 != int32(18) {
		v3524 = v3482
		goto L7
	} else {
		goto L1102
	}
L1099:
	;
	v3473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v3474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3474)
	v3477 = F_executeNextItem(m, l0, l1, v3468, l2, l3, v3474)
	mBase = m.M
	v3478 = m.ExcPending
	if v3478 != 0 {
		goto L1
	} else {
		goto L1100
	}
L1100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v3473)
	if l3|v3477 != 0 {
		v3482 = v3477
		goto L1098
	} else {
		goto L1101
	}
L1101:
	;
	v3524 = int32(0)
	goto L7
L1102:
	;
	if v3469 != 0 {
		goto L1103
	} else {
		goto L1104
	}
L1103:
	;
	v3490 = v24 + int32(752)
	goto L1105
L1104:
	;
	v3490 = int32(0)
	goto L1105
L1105:
	;
	v3491 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3492 = int32(1)
	v3493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v3497 = F_executeAnyItem(m, l0, v3490, v3491, l3, v3492, v3493, v3494, v3492, v3496)
	mBase = m.M
	v3498 = m.ExcPending
	if v3498 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1106:
	;
	v3524 = v3497
	goto L7
L1107:
	;
	v3500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v3500 != int32(1) {
		goto L1108
	} else {
		goto L1109
	}
L1108:
	;
	v3524 = int32(2)
	goto L7
L1109:
	;
	goto L1110
L1110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L1
	} else {
		goto L1111
	}
L1111:
	;
	F_errcode(m, int32(151781506))
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1112:
	;
	F_errmsg(m, int32(_a_F_executeItemOptUnwrapTarget_67), int32(0))
	mBase = m.M
	v3514 = m.ExcPending
	if v3514 != 0 {
		goto L1
	} else {
		goto L1113
	}
L1113:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapTarget_1), int32(978), int32(_a_F_executeItemOptUnwrapTarget_4))
	mBase = m.M
	v3519 = m.ExcPending
	if v3519 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
